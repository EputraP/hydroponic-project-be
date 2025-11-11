package handler

import (
	dto "hydroponic-be/internal/dto/transaction"
	errs "hydroponic-be/internal/errors"
	service "hydroponic-be/internal/service/transaction"
	"hydroponic-be/internal/util/logger"
	"hydroponic-be/internal/util/response"

	"github.com/gin-gonic/gin"
)

type AssetOpsTransactionHandler struct {
	assetOpsTransactionService service.AssetOpsTransactionService
}

type AssetOpsTransactionHandlerConfig struct {
	AssetOpsTransactionService service.AssetOpsTransactionService
}

func NewAssetOpsTransactionHandler(config AssetOpsTransactionHandlerConfig) *AssetOpsTransactionHandler {
	return &AssetOpsTransactionHandler{
		assetOpsTransactionService: config.AssetOpsTransactionService,
	}
}

func (h *AssetOpsTransactionHandler) CreateAssetOpsTransaction(c *gin.Context) {
	var createAssetOps *dto.AssetOpsTransaction

	if err := c.ShouldBindJSON(&createAssetOps); err != nil {
		logger.Error("AssetOpsTransactionHandler", "Invalid request body", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, errs.InvalidRequestBody.Error())
		return
	}

	logger.Info("AssetOpsTransactionHandler", "Starting asset ops creation process", map[string]string{})

	resp, err := h.assetOpsTransactionService.CreateAssetOpsTransaction(createAssetOps)
	if err != nil {
		logger.Error("AssetOpsTransactionHandler", "Failed to create asset ops record", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, err.Error())
		return
	}

	logger.Info("AssetOpsTransactionHandler", "Asset Ops record created successfully", map[string]string{})

	response.JSON(c, 201, "Create Asset Ops Record Success", resp)
}

func (h *AssetOpsTransactionHandler) GetAssetOpsTransaction(c *gin.Context) {

	logger.Info("AssetOpsTransactionHandler", "Init GetAssetOpsTransaction handler", nil)

	resp, err := h.assetOpsTransactionService.GetAssetOpsTransaction()
	if err != nil {
		logger.Error("AssetOpsTransactionHandler", "Failed to execute GetAssetOpsTransaction Service", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, err.Error())
		return
	}

	response.JSON(c, 200, "GetAssetOpsTransaction Success", resp)
}
