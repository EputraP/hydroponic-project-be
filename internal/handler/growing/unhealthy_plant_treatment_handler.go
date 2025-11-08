package handler

import (
	dto "hydroponic-be/internal/dto/growing"
	errs "hydroponic-be/internal/errors"
	service "hydroponic-be/internal/service/growing"
	"hydroponic-be/internal/util/logger"
	"hydroponic-be/internal/util/response"

	"github.com/gin-gonic/gin"
)

type UnhealthyPlantTreatmentHandler struct {
	unhealthyPlantTreatmentService service.UnhealthyPlantTreatmentService
}

type UnhealthyPlantTreatmentHandlerConfig struct {
	UnhealthyPlantTreatmentService service.UnhealthyPlantTreatmentService
}

func NewUnhealthyPlantTreatmentHandler(config UnhealthyPlantTreatmentHandlerConfig) *UnhealthyPlantTreatmentHandler {
	return &UnhealthyPlantTreatmentHandler{
		unhealthyPlantTreatmentService: config.UnhealthyPlantTreatmentService,
	}
}

func (h *UnhealthyPlantTreatmentHandler) CreateUnhealthyPlantTreatment(c *gin.Context) {
	var createUnhealthyPlantTreatment *dto.UnhealthyPlantTreatment

	if err := c.ShouldBindJSON(&createUnhealthyPlantTreatment); err != nil {
		logger.Error("UnhealthyPlantTreatmentHandler", "Invalid request body", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, errs.InvalidRequestBody.Error())
		return
	}

	logger.Info("UnhealthyPlantTreatmentHandler", "Starting unhealthy plant treatment record creation process", map[string]string{})

	resp, err := h.unhealthyPlantTreatmentService.CreateUnhealthyPlantTreatment(createUnhealthyPlantTreatment)
	if err != nil {
		logger.Error("UnhealthyPlantTreatmentHandler", "Failed to create unhealthy plant treatment record", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, err.Error())
		return
	}

	logger.Info("UnhealthyPlantTreatmentHandler", "unhealthy plant treatment record created successfully", map[string]string{})

	response.JSON(c, 201, "Create Unhealthy Plant Treatment Record Success", resp)
}
