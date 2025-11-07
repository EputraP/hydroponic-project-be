package routes

import (
	handlerAdmin "hydroponic-be/internal/handler/admin"
	handlerGrowing "hydroponic-be/internal/handler/growing"

	"github.com/gin-gonic/gin"
)

type HandlersAdmin struct {
	Plant     *handlerAdmin.PlantHandler
	Process   *handlerAdmin.ProcessHandler
	Remark    *handlerAdmin.RemarkHandler
	Uom       *handlerAdmin.UomHandler
	AssetType *handlerAdmin.AssetTypeHandler
	Asset     *handlerAdmin.AssetHandler
}

type HandlersGrowing struct {
	PlantGrowth *handlerGrowing.PlantGrowthHandler
}
type Handlers struct {
	Admin   HandlersAdmin
	Growing HandlersGrowing
}

type Middlewares struct {
}

func Build(srv *gin.Engine, h Handlers, middlewares Middlewares) {

	RoutesAdmin(srv, h.Admin, middlewares)
	RoutesGrowing(srv, h.Growing, middlewares)
}
