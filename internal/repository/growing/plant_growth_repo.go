package repository

import (
	"fmt"
	model "hydroponic-be/internal/model/growing"
	"hydroponic-be/internal/util/logger"

	"gorm.io/gorm"
)

type PlantGrowthRepository interface {
	CreatePlantGrowth(input *model.PlantGrowth) (*model.PlantGrowth, error)
}

type plantGrowthRepository struct {
	db *gorm.DB
}

func NewPlantGrowthRepository(db *gorm.DB) PlantGrowthRepository {
	return &plantGrowthRepository{
		db: db,
	}
}

func (r *plantGrowthRepository) CreatePlantGrowth(input *model.PlantGrowth) (*model.PlantGrowth, error) {

	logger.Info("plantGrowthRepository", "Inserting CreatePlantGrowth", map[string]string{})

	sqlScript := `INSERT INTO hydroponic_system.plant_growth
						( tower_id, process_id, "cycle", hss, hst, height, ph, ppm, water_level, ovr_plant_condition, remarks, plant_amount, created_at)
				  VALUES('%s', '%s', %d, %d, %d, %d, %f, %f, %d, '%s', '%s', %d, NOW());
		`
	sqlScript = fmt.Sprintf(sqlScript,
		input.TowerId,
		input.ProcessId,
		input.Cycle,
		input.HSS,
		input.HST,
		input.Height,
		input.PH,
		input.PPM,
		input.WaterLevel,
		input.OVRPlantCondition,
		input.Remarks,
		input.PlantAmount,
	)

	rows, err := r.db.Raw(sqlScript).Rows()
	if err != nil {
		logger.Error("plantGrowthRepository", "Failed to insert CreatePlantGrowth", map[string]string{
			"error": err.Error(),
		})
		return nil, err
	}
	defer rows.Close()

	logger.Info("plantGrowthRepository", "CreatePlantGrowth inserted successfully", map[string]string{})
	return input, nil
}
