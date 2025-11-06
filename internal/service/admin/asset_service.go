package service

import (
	dto "hydroponic-be/internal/dto/admin"
	errs "hydroponic-be/internal/errors"
	model "hydroponic-be/internal/model/admin"
	repository "hydroponic-be/internal/repository/admin"
	"hydroponic-be/internal/util/logger"
)

type AssetService interface {
	CreateAsset(input *dto.Asset) (*dto.Asset, error)
	GetAssets() (*[]model.Asset, error)
}

type assetService struct {
	assetRepo repository.AssetRepository
	plantRepo repository.PlantRepository
	uomRepo   repository.UomRepository
	assetType repository.AssetTypeRepository
}

type AssetServiceConfig struct {
	AssetRepo repository.AssetRepository
	PlantRepo repository.PlantRepository
	UomRepo   repository.UomRepository
	AssetType repository.AssetTypeRepository
}

func NewAssetService(config AssetServiceConfig) AssetService {
	return &assetService{
		assetRepo: config.AssetRepo,
		plantRepo: config.PlantRepo,
		uomRepo:   config.UomRepo,
		assetType: config.AssetType,
	}
}

func (s *assetService) CreateAsset(input *dto.Asset) (*dto.Asset, error) {
	logger.Info("assetService", "Creating new asset", map[string]string{
		"name": input.AssetName,
	})

	plantById, err := s.plantRepo.GetPlantById(&input.PlantID)
	if err != nil || len(*plantById) == 0 {
		logger.Error("assetService", "Invalid PlantID provided", map[string]string{
			"error": errs.InvalidPlantId.Error(),
		})
		return nil, errs.InvalidPlantId
	}

	uomById, err := s.uomRepo.GetUoMById(&input.UOMID)
	if err != nil || len(*uomById) == 0 {
		logger.Error("assetService", "Invalid UOMID provided", map[string]string{
			"error": errs.InvalidUoMId.Error(),
		})
		return nil, errs.InvalidUoMId
	}

	assetTypeById, err := s.assetType.GetAssetTypeById(&input.AssetTypeID)
	if err != nil || len(*assetTypeById) == 0 {
		logger.Error("assetService", "Invalid AssetTypeID provided", map[string]string{
			"error": errs.InvalidAssetTypeId.Error(),
		})
		return nil, errs.InvalidAssetTypeId
	}

	createdAsset, err := s.assetRepo.CreateAsset(&model.Asset{
		UOMID:       input.UOMID,
		PlantID:     input.PlantID,
		AssetTypeID: input.AssetTypeID,
		AssetName:   input.AssetName,
		Price:       input.Price,
		Value:       input.Value,
		Cycle:       input.Cycle,
	})
	if err != nil {
		logger.Error("assetService", "Error creating asset", map[string]string{
			"error": err.Error(),
		})
		return nil, errs.ErrorOnCreatingNewAsset
	}

	logger.Info("assetService", "Asset created successfully", map[string]string{
		"name": createdAsset.AssetName,
	})
	return &dto.Asset{
		UOMID:       createdAsset.UOMID,
		PlantID:     createdAsset.PlantID,
		AssetTypeID: createdAsset.AssetTypeID,
		AssetName:   createdAsset.AssetName,
		Price:       createdAsset.Price,
		Value:       createdAsset.Value,
		Cycle:       createdAsset.Cycle,
	}, nil
}

func (s *assetService) GetAssets() (*[]model.Asset, error) {
	logger.Info("assetService", "Init GetAssets Service", nil)

	res, err := s.assetRepo.GetAssets()
	if err != nil {
		logger.Error("assetService", "Failed to fetch GetAssets Repo", map[string]string{
			"error": err.Error(),
		})
		return nil, err
	}

	logger.Info("assetService", "Finished GetAssets Service", map[string]string{})
	return res, nil
}
