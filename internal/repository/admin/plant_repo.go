package repository

import (
	"fmt"
	model "hydroponic-be/internal/model/admin"
	"hydroponic-be/internal/util/logger"
	"strconv"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PlantRepository interface {
	CreateFarm() string
	GetPlants() (*[]model.Plant, error)
	GetPlantById(input *uuid.UUID) (*[]model.Plant, error)
	DeletePlant(inputModel *model.Plant) (*model.OnlyIdResponse, error)
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

func (r *plantRepository) GetPlants() (*[]model.Plant, error) {

	logger.Info("plantRepository", "Fetching GetPlants", map[string]string{})

	queryResult := &model.Plant{}
	result := []model.Plant{}

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

func (r *plantRepository) GetPlantById(input *uuid.UUID) (*[]model.Plant, error) {

	logger.Info("plantRepository", "Fetching GetPlantById", map[string]string{
		"plantID": input.String(),
	})

	queryResult := &model.Plant{}
	result := []model.Plant{}

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
					deleted_at is null and id = '%s'
		`
	sqlScript = fmt.Sprintf(sqlScript, input.String())

	rows, err := r.db.Raw(sqlScript).Rows()
	if err != nil {
		logger.Error("plantRepository", "Failed to fetch GetPlantById", map[string]string{
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
			logger.Error("plantRepository", "Failed to fetch scan GetPlantById result", map[string]string{
				"error": err.Error(),
			})
			return nil, fmt.Errorf("error scanning row: %w", err)
		}

		result = append(result, *queryResult)
	}

	logger.Info("plantRepository", "GetPlantById fetched successfully", map[string]string{
		"fetched": strconv.Itoa(len(result)),
	})
	return &result, nil

}

func (r *plantRepository) DeletePlant(inputModel *model.Plant) (*model.OnlyIdResponse, error) {
	logger.Info("plantRepository", "Deleting plant", map[string]string{
		"plantID": inputModel.ID.String(),
	})

	result := &model.OnlyIdResponse{}

	sqlScript := `UPDATE hydroponic_system.plants 
				SET deleted_at = %s 
				WHERE id = '%s' 
				RETURNING id`
	sqlScript = fmt.Sprintf(sqlScript, "NOW()", inputModel.ID.String())

	rows, err := r.db.Raw(sqlScript).Rows()
	if err != nil {
		logger.Error("plantRepository", "Failed to fetch DeletePlant", map[string]string{
			"error": err.Error(),
		})
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&result.ID,
		); err != nil {
			logger.Error("plantRepository", "Failed to fetch scan DeletePlant result", map[string]string{
				"error": err.Error(),
			})
			return nil, fmt.Errorf("error scanning row: %w", err)
		}
	}

	logger.Info("plantRepository", "Plant deleted successfully", map[string]string{
		"farmID": inputModel.ID.String(),
	})
	return result, nil
}
