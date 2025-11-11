package setup

import (
	handler "hydroponic-be/internal/handler/transaction"
	"hydroponic-be/internal/routes"
	service "hydroponic-be/internal/service/transaction"
)

func prepareTransaction(db *DBs) (handlers routes.HandlersTransaction) {

	assetOpsTransactionRepo := *db.Transaction.AssetOpsTransactionRepo
	assetRepo := *db.Admin.AssetRepo

	assetOpsTransactionService := service.NewAssetOpsTransactionService(service.AssetOpsTransactionServiceConfig{
		AssetOpsTransactionRepo: assetOpsTransactionRepo,
		AssetRepo:               assetRepo,
	})

	assetOpsTransactionHandler := handler.NewAssetOpsTransactionHandler(handler.AssetOpsTransactionHandlerConfig{
		AssetOpsTransactionService: assetOpsTransactionService,
	})

	handlers = routes.HandlersTransaction{
		AssetOpsTransaction: assetOpsTransactionHandler,
	}

	return
}
