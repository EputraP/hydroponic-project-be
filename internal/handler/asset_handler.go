package handler

import (
	"hydroponic-be/internal/service"
)

type AssetHandler struct {
	assetService service.AssetService
}

type AssetHandlerConfig struct {
	AssetService service.AssetService
}

func NewAssetHandler(config AssetHandlerConfig) *AssetHandler {
	return &AssetHandler{
		assetService: config.AssetService,
	}
}
