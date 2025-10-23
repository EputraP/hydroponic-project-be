package handler

import (
	"hydroponic-be/internal/util/response"

	"github.com/gin-gonic/gin"
)

type PlantHandler struct {
	// farmService      service.FarmService
	// systemLogService service.SystemLogService
}

type PlantHandlerConfig struct {
	// FarmService      service.FarmService
	// SystemLogService service.SystemLogService
}

func NewPlantHandler(config PlantHandlerConfig) *PlantHandler {
	return &PlantHandler{
		// farmService:      config.FarmService,
		// systemLogService: config.SystemLogService,
	}
}

func (h *PlantHandler) CreatePlant(c *gin.Context) {

	response.JSON(c, 201, "Create Plant Success", "")
}

func (h *PlantHandler) GetPlants(c *gin.Context) {

	response.JSON(c, 201, "Get Plant Success", "")
}

func (h *PlantHandler) DeletePlant(c *gin.Context) {

	response.JSON(c, 201, "Delete Plant Success", "")
}
