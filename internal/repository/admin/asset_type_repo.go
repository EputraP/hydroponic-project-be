package repository

import (
	"fmt"
	model "hydroponic-be/internal/model/admin"
	"hydroponic-be/internal/util/logger"
	"strconv"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AssetTypeRepository interface {
	GetAssetTypes() (*[]model.AssetType, error)
	GetAssetTypeById(input *uuid.UUID) (*[]model.AssetType, error)
}

type assetTypeRepository struct {
	db *gorm.DB
}

func NewAssetTypeRepository(db *gorm.DB) AssetTypeRepository {
	return &assetTypeRepository{
		db: db,
	}
}

func (r *assetTypeRepository) GetAssetTypes() (*[]model.AssetType, error) {

	logger.Info("assetTypeRepository", "Fetching GetAssetTypes", map[string]string{})

	queryResult := &model.AssetType{}
	result := []model.AssetType{}

	sqlScript := `select
					id,
					type_name,
					description
				from hydroponic_system.asset_types
				where
					deleted_at is null
		`

	rows, err := r.db.Raw(sqlScript).Rows()
	if err != nil {
		logger.Error("assetTypeRepository", "Failed to fetch GetAssetTypes", map[string]string{
			"error": err.Error(),
		})
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&queryResult.ID,
			&queryResult.TypeName,
			&queryResult.Description,
		); err != nil {
			logger.Error("assetTypeRepository", "Failed to fetch scan GetAssetTypes result", map[string]string{
				"error": err.Error(),
			})
			return nil, fmt.Errorf("error scanning row: %w", err)
		}

		result = append(result, *queryResult)
	}

	logger.Info("assetTypeRepository", "GetAssetTypes fetched successfully", map[string]string{
		"fetched": strconv.Itoa(len(result)),
	})
	return &result, nil

}

func (r *assetTypeRepository) GetAssetTypeById(input *uuid.UUID) (*[]model.AssetType, error) {

	logger.Info("assetTypeRepository", "Fetching GetAssetTypeById", map[string]string{})

	queryResult := &model.AssetType{}
	result := []model.AssetType{}

	sqlScript := `select
					id,
					type_name,
					description
				from hydroponic_system.asset_types
				where
					deleted_at is null and id = '%s'
		`

	sqlScript = fmt.Sprintf(sqlScript, input.String())
	rows, err := r.db.Raw(sqlScript).Rows()
	if err != nil {
		logger.Error("assetTypeRepository", "Failed to fetch GetAssetTypeById", map[string]string{
			"error": err.Error(),
		})
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&queryResult.ID,
			&queryResult.TypeName,
			&queryResult.Description,
		); err != nil {
			logger.Error("assetTypeRepository", "Failed to fetch scan GetAssetTypeById result", map[string]string{
				"error": err.Error(),
			})
			return nil, fmt.Errorf("error scanning row: %w", err)
		}

		result = append(result, *queryResult)
	}

	logger.Info("assetTypeRepository", "GetAssetTypeById fetched successfully", map[string]string{
		"fetched": strconv.Itoa(len(result)),
	})
	return &result, nil

}
