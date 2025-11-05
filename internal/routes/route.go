package routes

import (
	handler "hydroponic-be/internal/handler/admin"

	"github.com/gin-gonic/gin"
)

type HandlersAdmin struct {
	Plant     *handler.PlantHandler
	Process   *handler.ProcessHandler
	Remark    *handler.RemarkHandler
	Uom       *handler.UomHandler
	AssetType *handler.AssetTypeHandler
	Asset     *handler.AssetHandler
}

type Middlewares struct {
}

func Build(srv *gin.Engine, ha HandlersAdmin, middlewares Middlewares) {

	RoutesAdmin(srv, ha, middlewares)
}
