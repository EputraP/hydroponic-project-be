package routes

import "github.com/gin-gonic/gin"

func RoutesAdmin(srv *gin.Engine, h HandlersAdmin, middlewares Middlewares) {
	plant := srv.Group("/plant")
	plant.POST("/create", h.Plant.CreatePlant)
	plant.GET("/", h.Plant.GetPlants)
	plant.DELETE("/:plantId", h.Plant.DeletePlant)

	process := srv.Group("/process")
	process.GET("/", h.Process.GetProcesses)
	process.GET("/modules", h.Process.GetModules)
	process.GET("/modules/:processId", h.Process.GetSubModules)

	remark := srv.Group("/remark")
	remark.GET("/", h.Remark.GetRemarks)
	remark.GET("/by-process/:processId", h.Remark.GetRemarksByProcessId)

	uom := srv.Group("/uom")
	uom.GET("/", h.Uom.GetUoms)

	asset_type := srv.Group("/asset-type")
	asset_type.GET("/", h.AssetType.GetAssetTypes)

	asset := srv.Group("/asset")
	asset.POST("/create", h.Asset.CreateAsset)
	asset.GET("/", h.Asset.GetAssets)

}
