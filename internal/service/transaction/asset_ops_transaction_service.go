package service

import (
	dto "hydroponic-be/internal/dto/transaction"
	errs "hydroponic-be/internal/errors"
	model "hydroponic-be/internal/model/transaction"
	repositoryAdmin "hydroponic-be/internal/repository/admin"
	repository "hydroponic-be/internal/repository/transaction"
	"hydroponic-be/internal/util/logger"
)

type AssetOpsTransactionService interface {
	CreateAssetOpsTransaction(input *dto.AssetOpsTransaction) (*dto.AssetOpsTransaction, error)
}

type assetOpsTransactionService struct {
	assetOpsTransactionRepo repository.AssetOpsTransactionRepository
	assetRepo               repositoryAdmin.AssetRepository
}

type AssetOpsTransactionServiceConfig struct {
	AssetOpsTransactionRepo repository.AssetOpsTransactionRepository
	AssetRepo               repositoryAdmin.AssetRepository
}

func NewAssetOpsTransactionService(config AssetOpsTransactionServiceConfig) AssetOpsTransactionService {
	return &assetOpsTransactionService{
		assetOpsTransactionRepo: config.AssetOpsTransactionRepo,
		assetRepo:               config.AssetRepo,
	}
}

func (s *assetOpsTransactionService) CreateAssetOpsTransaction(input *dto.AssetOpsTransaction) (*dto.AssetOpsTransaction, error) {
	logger.Info("assetOpsTransactionService", "Creating new plant growth", map[string]string{})

	assetById, err := s.assetRepo.GetAssetById(&input.AssetId)
	if err != nil || len(*assetById) == 0 {
		logger.Error("assetOpsTransactionService", "Invalid AssetId provided", map[string]string{
			"error": errs.InvalidAssetId.Error(),
		})
		return nil, errs.InvalidAssetId
	}

	towerById, err := s.assetRepo.GetAssetById(&input.TowerId)
	if err != nil || len(*towerById) == 0 {
		logger.Error("assetOpsTransactionService", "Invalid TowerId provided", map[string]string{
			"error": errs.InvalidTowerId.Error(),
		})
		return nil, errs.InvalidTowerId
	}

	createdAssetOps, err := s.assetOpsTransactionRepo.CreateAssetOpsTransaction(&model.AssetOpsTransaction{
		TowerId: input.TowerId,
		AssetId: input.AssetId,
		Cycle:   input.Cycle,
		Value:   input.Value,
	})
	if err != nil {
		logger.Error("assetOpsTransactionService", "Error creating asset ops", map[string]string{
			"error": err.Error(),
		})
		return nil, errs.ErrorOnCreatingNewPlantGrowthRecord
	}

	logger.Info("assetOpsTransactionService", "Asset ops created successfully", map[string]string{})
	return &dto.AssetOpsTransaction{
		TowerId: createdAssetOps.TowerId,
		AssetId: createdAssetOps.AssetId,
		Value:   createdAssetOps.Value,
		Cycle:   createdAssetOps.Cycle,
	}, nil
}
