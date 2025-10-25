package handler

import (
	"hydroponic-be/internal/service"
	"hydroponic-be/internal/util/response"

	"github.com/gin-gonic/gin"
)

type PlantHandler struct {
	plantService service.PlantService
}

type PlantHandlerConfig struct {
	PlantService service.PlantService
}

func NewPlantHandler(config PlantHandlerConfig) *PlantHandler {
	return &PlantHandler{
		plantService: config.PlantService,
	}
}

func (h *PlantHandler) CreatePlant(c *gin.Context) {

	response.JSON(c, 201, "Create Plant Success", h.plantService.CreatePlant())
}

func (h *PlantHandler) GetPlants(c *gin.Context) {

	response.JSON(c, 200, "Get Plant Success", "")
}

func (h *PlantHandler) DeletePlant(c *gin.Context) {

	response.JSON(c, 201, "Delete Plant Success", "")
}
