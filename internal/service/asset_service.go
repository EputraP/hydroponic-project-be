package service

import (
	"hydroponic-be/internal/repository"
)

type AssetService interface {
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
