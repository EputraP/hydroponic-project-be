package routes

import (
	"hydroponic-be/internal/handler"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Plant   *handler.PlantHandler
	Process *handler.ProcessHandler
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

}
