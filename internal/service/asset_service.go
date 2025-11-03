package service

import (
	"hydroponic-be/internal/dto"
	errs "hydroponic-be/internal/errors"
	"hydroponic-be/internal/model"
	"hydroponic-be/internal/repository"
	"hydroponic-be/internal/util/logger"
)

type AssetService interface {
	CreateAsset(input *dto.Asset) (*dto.Asset, error)
}

type assetService struct {
	assetRepo repository.AssetRepository
}

type AssetServiceConfig struct {
	AssetRepo repository.AssetRepository
}

func NewAssetService(config AssetServiceConfig) AssetService {
	return &assetService{
		assetRepo: config.AssetRepo,
	}
}

func (s *assetService) CreateAsset(input *dto.Asset) (*dto.Asset, error) {
	logger.Info("assetService", "Creating new asset", map[string]string{
		"name": input.AssetName,
	})

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
