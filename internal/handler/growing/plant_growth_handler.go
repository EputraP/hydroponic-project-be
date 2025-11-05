package handler

import (
	service "hydroponic-be/internal/service/growing"
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
