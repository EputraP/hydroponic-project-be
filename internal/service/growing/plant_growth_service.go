package service

import (
	dto "hydroponic-be/internal/dto/growing"
	errs "hydroponic-be/internal/errors"
	model "hydroponic-be/internal/model/growing"
	adminRepository "hydroponic-be/internal/repository/admin"
	repository "hydroponic-be/internal/repository/growing"
	"hydroponic-be/internal/util/logger"
)

type PlantGrowthService interface {
	CreatePlantGrowth(input *dto.PlantGrowth) (*dto.PlantGrowth, error)
}

type plantGrowthService struct {
	plantGrowthRepo repository.PlantGrowthRepository
	assetRepo       adminRepository.AssetRepository
	processRepo     adminRepository.ProcessRepository
}

type PlantGrowthServiceConfig struct {
	PlantGrowthRepo repository.PlantGrowthRepository
	AssetRepo       adminRepository.AssetRepository
	ProcessRepo     adminRepository.ProcessRepository
}

func NewPlantGrowthService(config PlantGrowthServiceConfig) PlantGrowthService {
	return &plantGrowthService{
		plantGrowthRepo: config.PlantGrowthRepo,
		assetRepo:       config.AssetRepo,
		processRepo:     config.ProcessRepo,
	}
}

func (s *plantGrowthService) CreatePlantGrowth(input *dto.PlantGrowth) (*dto.PlantGrowth, error) {
	logger.Info("plantGrowthService", "Creating new plant growth", map[string]string{})

	assetById, err := s.assetRepo.GetAssetById(&input.TowerId)
	if err != nil || len(*assetById) == 0 {
		logger.Error("plantGrowthService", "Invalid AssetId provided", map[string]string{
			"error": errs.InvalidAssetId.Error(),
		})
		return nil, errs.InvalidAssetId
	}

	processById, err := s.processRepo.GetProcessById(&input.ProcessId)
	if err != nil || len(*processById) == 0 {
		logger.Error("plantGrowthService", "Invalid ProcessIs provided", map[string]string{
			"error": errs.InvalidProcessId.Error(),
		})
		return nil, errs.InvalidProcessId
	}

	createdAsset, err := s.plantGrowthRepo.CreatePlantGrowth(&model.PlantGrowth{
		TowerId:           input.TowerId,
		ProcessId:         input.ProcessId,
		Cycle:             input.Cycle,
		HSS:               input.HSS,
		HST:               input.HST,
		Height:            input.Height,
		PH:                input.PH,
		PPM:               input.PPM,
		WaterLevel:        input.WaterLevel,
		OVRPlantCondition: input.OVRPlantCondition,
		Remarks:           input.Remarks,
		PlantAmount:       input.PlantAmount,
	})
	if err != nil {
		logger.Error("plantGrowthService", "Error creating asset", map[string]string{
			"error": err.Error(),
		})
		return nil, errs.ErrorOnCreatingNewPlantGrowthRecord
	}

	logger.Info("assetService", "Asset created successfully", map[string]string{})
	return &dto.PlantGrowth{
		TowerId:           createdAsset.TowerId,
		ProcessId:         createdAsset.ProcessId,
		Cycle:             createdAsset.Cycle,
		HSS:               createdAsset.HSS,
		HST:               createdAsset.HST,
		Height:            createdAsset.Height,
		PH:                createdAsset.PH,
		PPM:               createdAsset.PPM,
		WaterLevel:        createdAsset.WaterLevel,
		OVRPlantCondition: createdAsset.OVRPlantCondition,
		Remarks:           createdAsset.Remarks,
		PlantAmount:       createdAsset.PlantAmount,
	}, nil
}
