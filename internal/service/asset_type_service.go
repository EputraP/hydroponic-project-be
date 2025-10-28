package service

import (
	"hydroponic-be/internal/model"
	"hydroponic-be/internal/repository"
	"hydroponic-be/internal/util/logger"
)

type AssetTypeService interface {
	GetAssetTypes() (*[]model.AssetType, error)
}

type assetTypeService struct {
	assetTypeRepo repository.AssetTypeRepository
}

type AssetTypeServiceConfig struct {
	AssetTypeRepo repository.AssetTypeRepository
}

func NewAssetTypeService(config AssetTypeServiceConfig) AssetTypeService {
	return &assetTypeService{
		assetTypeRepo: config.AssetTypeRepo,
	}
}

func (s *assetTypeService) GetAssetTypes() (*[]model.AssetType, error) {
	logger.Info("assetTypeService", "Init GetAssetTypes Service", nil)

	res, err := s.assetTypeRepo.GetAssetTypes()
	if err != nil {
		logger.Error("assetTypeService", "Failed to fetch GetAssetTypes Repo", map[string]string{
			"error": err.Error(),
		})
		return nil, err
	}

	logger.Info("assetTypeService", "Finished GetAssetTypes Service", map[string]string{})
	return res, nil
}
