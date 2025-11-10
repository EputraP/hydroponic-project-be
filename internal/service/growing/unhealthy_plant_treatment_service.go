package service

import (
	dto "hydroponic-be/internal/dto/growing"
	errs "hydroponic-be/internal/errors"
	model "hydroponic-be/internal/model/growing"
	repositoryAdmin "hydroponic-be/internal/repository/admin"
	repository "hydroponic-be/internal/repository/growing"
	"hydroponic-be/internal/util/logger"
)

type UnhealthyPlantTreatmentService interface {
	CreateUnhealthyPlantTreatment(input *dto.UnhealthyPlantTreatment) (*dto.UnhealthyPlantTreatment, error)
	GetUnhealthyPlantTreatment() (*[]model.UnhealthyPlantTreatment, error)
}

type unhealthyPlantTreatmentService struct {
	unhealthyPlantTreatmentRepo repository.UnhealthyPlantTreatmentRepository
	assetRepo                   repositoryAdmin.AssetRepository
}

type UnhealthyPlantTreatmentServiceConfig struct {
	UnhealthyPlantTreatmentRepo repository.UnhealthyPlantTreatmentRepository
	AssetRepo                   repositoryAdmin.AssetRepository
}

func NewUnhealthyPlantTreatmentService(config UnhealthyPlantTreatmentServiceConfig) UnhealthyPlantTreatmentService {
	return &unhealthyPlantTreatmentService{
		unhealthyPlantTreatmentRepo: config.UnhealthyPlantTreatmentRepo,
		assetRepo:                   config.AssetRepo,
	}
}

func (s *unhealthyPlantTreatmentService) CreateUnhealthyPlantTreatment(input *dto.UnhealthyPlantTreatment) (*dto.UnhealthyPlantTreatment, error) {
	logger.Info("unhealthyPlantTreatmentService", "Creating new Unhealthy Plant Treatment", map[string]string{})

	assetById, err := s.assetRepo.GetAssetById(&input.TowerId)
	if err != nil || len(*assetById) == 0 {
		logger.Error("unhealthyPlantTreatmentService", "Invalid ProcessIs provided", map[string]string{
			"error": errs.InvalidAssetId.Error(),
		})
		return nil, errs.InvalidAssetId
	}

	createdPlantTreatment, err := s.unhealthyPlantTreatmentRepo.CreateUnhealthyPlantTreatment(
		&model.UnhealthyPlantTreatment{
			TowerId:        input.TowerId,
			Cycle:          input.Cycle,
			TowerHoleNo:    input.TowerHoleNo,
			Treatment:      input.Treatment,
			AfterTreatment: input.AfterTreatment,
		},
	)
	if err != nil {
		logger.Error("unhealthyPlantTreatmentService", "Error creating asset", map[string]string{
			"error": err.Error(),
		})
		return nil, errs.ErrorOnCreatingNewPlantGrowthRecord
	}

	logger.Info("unhealthyPlantTreatmentService", "Asset created successfully", map[string]string{})
	return &dto.UnhealthyPlantTreatment{
		TowerId:        createdPlantTreatment.TowerId,
		Cycle:          createdPlantTreatment.Cycle,
		TowerHoleNo:    createdPlantTreatment.TowerHoleNo,
		Treatment:      createdPlantTreatment.Treatment,
		AfterTreatment: createdPlantTreatment.AfterTreatment,
	}, nil
}

func (s *unhealthyPlantTreatmentService) GetUnhealthyPlantTreatment() (*[]model.UnhealthyPlantTreatment, error) {
	logger.Info("unhealthyPlantTreatmentService", "Init GetUnhealthyPlantTreatment Service", nil)

	res, err := s.unhealthyPlantTreatmentRepo.GetUnhealthyPlantTreatment()
	if err != nil {
		logger.Error("unhealthyPlantTreatmentService", "Failed to fetch GetUnhealthyPlantTreatment Repo", map[string]string{
			"error": err.Error(),
		})
		return nil, err
	}

	logger.Info("unhealthyPlantTreatmentService", "Finished GetUnhealthyPlantTreatment Service", map[string]string{})
	return res, nil
}
