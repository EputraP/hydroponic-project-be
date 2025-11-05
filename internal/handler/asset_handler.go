package handler

import (
	"hydroponic-be/internal/dto"
	errs "hydroponic-be/internal/errors"
	"hydroponic-be/internal/service"
	"hydroponic-be/internal/util/logger"
	"hydroponic-be/internal/util/response"

	"github.com/gin-gonic/gin"
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

func (h *AssetHandler) CreateAsset(c *gin.Context) {
	var createAsset *dto.Asset

	if err := c.ShouldBindJSON(&createAsset); err != nil {
		logger.Error("AssetHandler", "Invalid request body", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, errs.InvalidRequestBody.Error())
		return
	}

	logger.Info("AssetHandler", "Starting asset creation process", map[string]string{
		"name": createAsset.AssetName,
	})

	resp, err := h.assetService.CreateAsset(createAsset)
	if err != nil {
		logger.Error("AssetHandler", "Failed to create asset", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, err.Error())
		return
	}

	logger.Info("AssetHandler", "Asset created successfully", map[string]string{
		"name": resp.AssetName,
	})

	response.JSON(c, 201, "Create Asset Success", resp)
}

func (h *AssetHandler) GetAssets(c *gin.Context) {

	logger.Info("assetHanlder", "Init GetAssets handler", nil)

	resp, err := h.assetService.GetAssets()
	if err != nil {
		logger.Error("assetHanlder", "Failed to execute GetAssets Service", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, err.Error())
		return
	}

	response.JSON(c, 200, "GetAssets Success", resp)
}
