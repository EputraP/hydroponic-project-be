package routes

import "github.com/gin-gonic/gin"

func RoutesGrowing(srv *gin.Engine, h HandlersGrowing, middlewares Middlewares) {
	plantGrowth := srv.Group("/plant-growth")
	plantGrowth.POST("/create", h.PlantGrowth.CreatePlantGrowth)
	plantGrowth.GET("/", h.PlantGrowth.GetPlantGrowths)

	unhealthyPlantTreatment := srv.Group("/unhealthy-plant-treatment")
	unhealthyPlantTreatment.POST("/create", h.UnhealthyPlantTreatment.CreateUnhealthyPlantTreatment)
	unhealthyPlantTreatment.GET("/", h.UnhealthyPlantTreatment.GetUnhealthyPlantTreatment)
}
