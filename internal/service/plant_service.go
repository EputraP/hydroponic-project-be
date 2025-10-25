package service

import "hydroponic-be/internal/repository"

type PlantService interface {
	CreatePlant() string
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

func (s *plantService) GetPlants() {

}

func (s *plantService) DeletePlant() {

}
