package routes

import (
	"hydroponic-be/internal/handler"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Plant     *handler.PlantHandler
	Process   *handler.ProcessHandler
	Remark    *handler.RemarkHandler
	Uom       *handler.UomHandler
	AssetType *handler.AssetTypeHandler
}

type Middlewares struct {
}

func Build(srv *gin.Engine, h Handlers, middlewares Middlewares) {

	plant := srv.Group("/plant")
	plant.POST("/create", h.Plant.CreatePlant)
	plant.GET("/", h.Plant.GetPlants)
	plant.DELETE("/:plantId", h.Plant.DeletePlant)

	process := srv.Group("/process")
	process.GET("/", h.Process.GetProcesses)
	process.GET("/modules", h.Process.GetModules)

	remark := srv.Group("/remark")
	remark.GET("/", h.Remark.GetRemarks)
	remark.GET("/by-process/:processId", h.Remark.GetRemarksByProcessId)

	uom := srv.Group("/uom")
	uom.GET("/", h.Uom.GetUoms)

	asset_type := srv.Group("/asset-type")
	asset_type.GET("/", h.AssetType.GetAssetTypes)
}
