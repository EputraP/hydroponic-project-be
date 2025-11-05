package service

import (
	"hydroponic-be/internal/dto"
	"hydroponic-be/internal/model"
	repository "hydroponic-be/internal/repository/admin"
	"hydroponic-be/internal/util/logger"

	"github.com/google/uuid"
)

type PlantService interface {
	CreatePlant() string
	GetPlants() (*[]model.Plant, error)
	DeletePlant(plantId *uuid.UUID) (*dto.OnlyIdResponse, error)
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

func (s *plantService) GetPlants() (*[]model.Plant, error) {
	logger.Info("plantService", "Init GetPlants Service", nil)

	res, err := s.plantRepo.GetPlants()
	if err != nil {
		logger.Error("plantService", "Failed to fetch GetPlants Repo", map[string]string{
			"error": err.Error(),
		})
		return nil, err
	}

	logger.Info("plantService", "Finished GetPlants Service", map[string]string{})
	return res, nil
}

func (s *plantService) DeletePlant(plantId *uuid.UUID) (*dto.OnlyIdResponse, error) {
	logger.Info("plantService", "Init DeleteFarm Service", map[string]string{
		"plant_id": plantId.String(),
	})

	res, err := s.plantRepo.DeletePlant(&model.Plant{ID: *plantId})
	if err != nil {
		logger.Error("plantService", "Failed to execute DeleteFarm Repo", map[string]string{
			"plant_id": plantId.String(),
			"error":    err.Error(),
		})
		return nil, err
	}

	logger.Info("plantService", "Finished DelteFarm", map[string]string{
		"plant_id": res.ID.String(),
	})
	return &dto.OnlyIdResponse{
		ID: res.ID,
	}, nil
}
