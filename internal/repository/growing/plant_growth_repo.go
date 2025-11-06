package repository

import (
	"fmt"
	"hydroponic-be/internal/model"
	"hydroponic-be/internal/util/logger"

	"gorm.io/gorm"
)

type PlantGrowthRepository interface {
}

type plantGrowthRepository struct {
	db *gorm.DB
}

func NewPlantGrowthRepository(db *gorm.DB) PlantGrowthRepository {
	return &plantGrowthRepository{
		db: db,
	}
}

func (r *plantGrowthRepository) CreatePlantGrowth(input *model.Asset) (*model.Asset, error) {

	logger.Info("plantGrowthRepository", "Inserting CreatePlantGrowth", map[string]string{})

	sqlScript := `INSERT INTO hydroponic_system.plant_growth
						( tower_id, process_id, "cycle", hss, hst, height, ph, ppm, water_level, ovr_plant_condition, remarks, plant_amount, created_at)
				  VALUES(?, ?, 0, 0, 0, 0, 0, 0, 0, '', '', 0, '');
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
		logger.Error("plantGrowthRepository", "Failed to insert CreatePlantGrowth", map[string]string{
			"error": err.Error(),
		})
		return nil, err
	}
	defer rows.Close()

	logger.Info("plantGrowthRepository", "CreatePlantGrowth inserted successfully", map[string]string{
		"asset_name": input.AssetName,
	})
	return input, nil
}
