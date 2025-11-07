package repository

import (
	"fmt"
	model "hydroponic-be/internal/model/growing"
	"hydroponic-be/internal/util/logger"
	"strconv"

	"gorm.io/gorm"
)

type PlantGrowthRepository interface {
	CreatePlantGrowth(input *model.PlantGrowth) (*model.PlantGrowth, error)
	GetPlantGrowths() (*[]model.PlantGrowth, error)
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

func (r *plantGrowthRepository) GetPlantGrowths() (*[]model.PlantGrowth, error) {

	logger.Info("plantGrowthRepository", "Fetching GetPlantGrowths", map[string]string{})

	queryResult := &model.PlantGrowth{}
	result := []model.PlantGrowth{}

	sqlScript := `SELECT 
					id, 
					tower_id, 
					process_id, 
					"cycle", 
					hss, 
					hst, 
					height, 
					ph, 
					ppm, 
					water_level, 
					ovr_plant_condition, 
					remarks, 
					plant_amount
				FROM hydroponic_system.plant_growth
				where
					deleted_at is NULL
		`

	rows, err := r.db.Raw(sqlScript).Rows()
	if err != nil {
		logger.Error("plantGrowthRepository", "Failed to fetch GetPlantGrowths", map[string]string{
			"error": err.Error(),
		})
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&queryResult.ID,
			&queryResult.TowerId,
			&queryResult.ProcessId,
			&queryResult.Cycle,
			&queryResult.HSS,
			&queryResult.HST,
			&queryResult.Height,
			&queryResult.PH,
			&queryResult.PPM,
			&queryResult.WaterLevel,
			&queryResult.OVRPlantCondition,
			&queryResult.Remarks,
			&queryResult.PlantAmount,
		); err != nil {
			logger.Error("plantGrowthRepository", "Failed to fetch scan GetPlantGrowths result", map[string]string{
				"error": err.Error(),
			})
			return nil, fmt.Errorf("error scanning row: %w", err)
		}

		result = append(result, *queryResult)
	}

	logger.Info("plantGrowthRepository", "GetPlantGrowths fetched successfully", map[string]string{
		"fetched": strconv.Itoa(len(result)),
	})
	return &result, nil

}
