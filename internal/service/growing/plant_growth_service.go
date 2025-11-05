package service

import repository "hydroponic-be/internal/repository/growing"

type PlantGrowthService interface {
}

type plantGrowthService struct {
	plantGrowthRepo repository.PlantGrowthRepository
}

type PlantGrowthServiceConfig struct {
	PlantGrowthRepo repository.PlantGrowthRepository
}

func NewPlantGrowthService(config PlantGrowthServiceConfig) PlantGrowthService {
	return &plantGrowthService{
		plantGrowthRepo: config.PlantGrowthRepo,
	}
}
