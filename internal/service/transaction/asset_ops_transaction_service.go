package service

import repository "hydroponic-be/internal/repository/transaction"

type AssetOpsTransactionService interface {
}

type assetOpsTransactionService struct {
	assetOpsTransactionRepo repository.AssetOpsTransactionRepository
}

type AssetOpsTransactionServiceConfig struct {
	AssetOpsTransactionRepo repository.AssetOpsTransactionRepository
}

func NewAssetOpsTransactionService(config AssetOpsTransactionServiceConfig) AssetOpsTransactionService {
	return &assetOpsTransactionService{
		assetOpsTransactionRepo: config.AssetOpsTransactionRepo,
	}
}
