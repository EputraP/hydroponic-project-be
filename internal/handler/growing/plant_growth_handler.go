package handler

import (
	dto "hydroponic-be/internal/dto/growing"
	errs "hydroponic-be/internal/errors"
	service "hydroponic-be/internal/service/growing"
	"hydroponic-be/internal/util/logger"
	"hydroponic-be/internal/util/response"

	"github.com/gin-gonic/gin"
)

type PlantGrowthHandler struct {
	plantGrowthService service.PlantGrowthService
}

type PlantGrowthHandlerConfig struct {
	PlantGrowthService service.PlantGrowthService
}

func NewPlantGrowthHandler(config PlantGrowthHandlerConfig) *PlantGrowthHandler {
	return &PlantGrowthHandler{
		plantGrowthService: config.PlantGrowthService,
	}
}

func (h *PlantGrowthHandler) CreatePlantGrowth(c *gin.Context) {
	var createPlantGrowth *dto.PlantGrowth

	if err := c.ShouldBindJSON(&createPlantGrowth); err != nil {
		logger.Error("PlantGrowthHandler", "Invalid request body", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, errs.InvalidRequestBody.Error())
		return
	}

	logger.Info("PlantGrowthHandler", "Starting plant growth record creation process", map[string]string{})

	resp, err := h.plantGrowthService.CreatePlantGrowth(createPlantGrowth)
	if err != nil {
		logger.Error("PlantGrowthHandler", "Failed to create plant growth record", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, err.Error())
		return
	}

	logger.Info("PlantGrowthHandler", "plant growth record created successfully", map[string]string{})

	response.JSON(c, 201, "Create Plant Growth Record Success", resp)
}
