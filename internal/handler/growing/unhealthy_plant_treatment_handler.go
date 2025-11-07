package handler

import (
	service "hydroponic-be/internal/service/growing"
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
