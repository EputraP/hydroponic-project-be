package handler

import (
	service "hydroponic-be/internal/service/transaction"
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
