package handler

import (
	errs "hydroponic-be/internal/errors"
	"hydroponic-be/internal/service"
	"hydroponic-be/internal/util/logger"
	"hydroponic-be/internal/util/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	logger.Info("plantHandler", "Init GetPlants handler", nil)

	resp, err := h.plantService.GetPlants()
	if err != nil {
		logger.Error("plantHandler", "Failed to execute GetPlants Service", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, err.Error())
		return
	}

	response.JSON(c, 200, "Get Plant Success", resp)
}

func (h *PlantHandler) DeletePlant(c *gin.Context) {
	plantId := c.Param("plantId")
	id, paramErr := uuid.Parse(plantId)
	if paramErr != nil {
		logger.Error("plantHandler", "Invalid plant ID parameter", map[string]string{
			"plantId": plantId,
		})
		response.Error(c, 400, errs.InvalidParam.Error())
		return
	}

	logger.Info("plantHandler", "Init DeletePlant handler", map[string]string{
		"plantId": plantId,
	})

	resp, err := h.plantService.DeletePlant(&id)
	if err != nil {
		logger.Error("plantHandler", "Failed to execute DeletePlant service", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, err.Error())
		return
	}

	response.JSON(c, 201, "Delete Plant Success", resp)
}
