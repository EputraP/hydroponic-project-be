package repository

import (
	"fmt"
	"hydroponic-be/internal/dto"
	"hydroponic-be/internal/util/logger"
	"strconv"

	"gorm.io/gorm"
)

type PlantRepository interface {
	CreateFarm() string
	GetPlants() (*[]dto.GetPlants, error)
}

type plantRepository struct {
	db *gorm.DB
}

func NewPlantRepository(db *gorm.DB) PlantRepository {
	return &plantRepository{
		db: db,
	}
}

func (r *plantRepository) CreateFarm() string {
	return "Adili Jokowi"
}

func (r *plantRepository) GetPlants() (*[]dto.GetPlants, error) {

	logger.Info("plantRepository", "Fetching GetPlants", map[string]string{})

	queryResult := &dto.GetPlants{}
	result := []dto.GetPlants{}

	sqlScript := `select
					id,
					plant_name,
					scientific_name,
					variety,
					plant_type,
					description,
					ph_min,
					ph_max,
					ppm_min,
					ppm_max,
					light_hours,
					optimal_temperature_min,
					optimal_temperature_max,
					harvest_days,
					germination_days,
					hss_days,
					hst_days
				from hydroponic_system.plants
				where
					deleted_at is null
		`

	rows, err := r.db.Raw(sqlScript).Rows()
	if err != nil {
		logger.Error("plantRepository", "Failed to fetch GetPlants", map[string]string{
			"error": err.Error(),
		})
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&queryResult.ID,
			&queryResult.PlantName,
			&queryResult.ScientificName,
			&queryResult.Variety,
			&queryResult.PlantType,
			&queryResult.Description,
			&queryResult.PHMin,
			&queryResult.PHMax,
			&queryResult.PPMMin,
			&queryResult.PPMMax,
			&queryResult.LightHours,
			&queryResult.TemperatureMin,
			&queryResult.TemperatureMax,
			&queryResult.HarvestDays,
			&queryResult.GerminationDays,
			&queryResult.HSSDays,
			&queryResult.HSTDays,
		); err != nil {
			logger.Error("plantRepository", "Failed to fetch scan GetPlants result", map[string]string{
				"error": err.Error(),
			})
			return nil, fmt.Errorf("error scanning row: %w", err)
		}

		result = append(result, *queryResult)
	}

	logger.Info("plantRepository", "GetPlants fetched successfully", map[string]string{
		"fetched": strconv.Itoa(len(result)),
	})
	return &result, nil

}
