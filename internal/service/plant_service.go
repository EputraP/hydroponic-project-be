package service

import (
	"hydroponic-be/internal/dto"
	"hydroponic-be/internal/repository"
	"hydroponic-be/internal/util/logger"
)

type PlantService interface {
	CreatePlant() string
	GetPlants() (*[]dto.GetPlants, error)
}

type plantService struct {
	plantRepo repository.PlantRepository
}

type PlantServiceConfig struct {
	PlantRepo repository.PlantRepository
}

func NewPlantService(config PlantServiceConfig) PlantService {
	return &plantService{
		plantRepo: config.PlantRepo,
	}
}

func (s *plantService) CreatePlant() string {
	return s.plantRepo.CreateFarm()

}

func (s *plantService) GetPlants() (*[]dto.GetPlants, error) {
	logger.Info("plant Service", "Init GetPlants Service", nil)

	res, err := s.plantRepo.GetPlants()
	if err != nil {
		logger.Error("plantService", "Failed to fetch GetPlants", map[string]string{
			"error": err.Error(),
		})
		return nil, err
	}

	logger.Info("plantService", "Finished GetPlants Service", map[string]string{})
	return res, nil
}

func (s *plantService) DeletePlant() {

}
