package routes

import "github.com/gin-gonic/gin"

func RoutesTransaction(srv *gin.Engine, h HandlersTransaction, middlewares Middlewares) {
	assetOpsTransaction := srv.Group("/asset-ops-transaction")
	assetOpsTransaction.POST("/create", h.AssetOpsTransaction.CreateAssetOpsTransaction)

}
