package repository

import (
	"fmt"
	"hydroponic-be/internal/model"
	"hydroponic-be/internal/util/logger"

	"gorm.io/gorm"
)

type AssetRepository interface {
	CreateAsset(input *model.Asset) (*model.Asset, error)
}

type assetRepository struct {
	db *gorm.DB
}

func NewAssetRepository(db *gorm.DB) RemarkRepository {
	return &remarkRepository{
		db: db,
	}
}

func (r *assetRepository) CreateAsset(input *model.Asset) (*model.Asset, error) {

	logger.Info("assetRepository", "Inserting CreateAsset", map[string]string{})

	sqlScript := `INSERT INTO hydroponic_system.assets
					(uom_id, 
					plant_id, 
					asset_name, 
					price, 
					value, 
					"cycle", 
					created_at,  
					asset_type_id)
				VALUES
					('%s' 
					'%s', 
					%s, 
					%d, 
					%d, 
					%d, 
					NOW(), 
					'%s');'
		`
	sqlScript = fmt.Sprintf(sqlScript,
		input.UOMID,
		input.PlantID,
		input.AssetName,
		input.Price,
		input.Value,
		input.Cycle,
		input.AssetTypeID,
	)

	rows, err := r.db.Raw(sqlScript).Rows()
	if err != nil {
		logger.Error("assetRepository", "Failed to insert CreateAsset", map[string]string{
			"error": err.Error(),
		})
		return nil, err
	}
	defer rows.Close()

	logger.Info("assetRepository", "CreateAsset inserted successfully", map[string]string{
		"asset_name": input.AssetName,
	})
	return input, nil
}
