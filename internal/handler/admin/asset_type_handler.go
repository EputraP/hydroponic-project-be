package handler

import (
	service "hydroponic-be/internal/service/admin"
	"hydroponic-be/internal/util/logger"
	"hydroponic-be/internal/util/response"

	"github.com/gin-gonic/gin"
)

type AssetTypeHandler struct {
	assetTypeService service.AssetTypeService
}

type AssetTypeHandlerConfig struct {
	AssetTypeService service.AssetTypeService
}

func NewAssetTypeHandler(config AssetTypeHandlerConfig) *AssetTypeHandler {
	return &AssetTypeHandler{
		assetTypeService: config.AssetTypeService,
	}
}

func (h *AssetTypeHandler) GetAssetTypes(c *gin.Context) {

	logger.Info("AssetTypeHandler", "Init GetAssetTypes handler", nil)

	resp, err := h.assetTypeService.GetAssetTypes()
	if err != nil {
		logger.Error("AssetTypeHandler", "Failed to execute GetAssetTypes Service", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, err.Error())
		return
	}

	response.JSON(c, 200, "GetProcesses Success", resp)
}
