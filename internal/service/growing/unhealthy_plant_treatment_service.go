package service

import (
	repository "hydroponic-be/internal/repository/growing"
)

type UnhealthyPlantTreatmentService interface {
}

type unhealthyPlantTreatmentService struct {
	plantGrowthRepo repository.UnhealthyPlantTreatmentRepository
}

type UnhealthyPlantTreatmentServiceConfig struct {
	PlantGrowthRepo repository.UnhealthyPlantTreatmentRepository
}

func NewUnhealthyPlantTreatmentService(config UnhealthyPlantTreatmentServiceConfig) UnhealthyPlantTreatmentService {
	return &unhealthyPlantTreatmentService{
		plantGrowthRepo: config.PlantGrowthRepo,
	}
}
