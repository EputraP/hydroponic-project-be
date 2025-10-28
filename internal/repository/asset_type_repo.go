package repository

import (
	"fmt"
	"hydroponic-be/internal/model"
	"hydroponic-be/internal/util/logger"
	"strconv"

	"gorm.io/gorm"
)

type AssetTypeRepository interface {
	GetAssetTypes() (*[]model.AssetType, error)
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
