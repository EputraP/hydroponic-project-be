package repository

import (
	"fmt"
	model "hydroponic-be/internal/model/growing"
	"hydroponic-be/internal/util/logger"
	"strconv"

	"gorm.io/gorm"
)

type UnhealthyPlantTreatmentRepository interface {
	CreateUnhealthyPlantTreatment(input *model.UnhealthyPlantTreatment) (*model.UnhealthyPlantTreatment, error)
	GetUnhealthyPlantTreatment() (*[]model.UnhealthyPlantTreatment, error)
}

type unhealthyPlantTreatmentRepository struct {
	db *gorm.DB
}

func NewUnhealthyPlantTreatmentRepository(db *gorm.DB) UnhealthyPlantTreatmentRepository {
	return &unhealthyPlantTreatmentRepository{
		db: db,
	}
}

func (r *unhealthyPlantTreatmentRepository) CreateUnhealthyPlantTreatment(input *model.UnhealthyPlantTreatment) (*model.UnhealthyPlantTreatment, error) {

	logger.Info("unhealthyPlantTreatmentRepository", "Inserting CreateUnhelathyPlantTreatment", map[string]string{})

	sqlScript := `INSERT INTO hydroponic_system.unhealthy_plant_treatment
				(tower_id, "cycle", tower_hole_no, treatment, after_treatment, created_at)
				VALUES('%s', %d, %d, '%s', '%s', NOW());
		`
	sqlScript = fmt.Sprintf(sqlScript,
		input.TowerId,
		input.Cycle,
		input.TowerHoleNo,
		input.Treatment,
		input.AfterTreatment,
	)

	rows, err := r.db.Raw(sqlScript).Rows()
	if err != nil {
		logger.Error("unhealthyPlantTreatmentRepository", "Failed to insert CreateUnhelathyPlantTreatment", map[string]string{
			"error": err.Error(),
		})
		return nil, err
	}
	defer rows.Close()

	logger.Info("unhealthyPlantTreatmentRepository", "CreateUnhelathyPlantTreatment inserted successfully", map[string]string{})
	return input, nil
}

func (r *unhealthyPlantTreatmentRepository) GetUnhealthyPlantTreatment() (*[]model.UnhealthyPlantTreatment, error) {

	logger.Info("unhealthyPlantTreatmentRepository", "Fetching GetUnhealthyPlantTreatment", map[string]string{})

	queryResult := &model.UnhealthyPlantTreatment{}
	result := []model.UnhealthyPlantTreatment{}

	sqlScript := `SELECT 
					id, 
					tower_id, 
					"cycle", 
					tower_hole_no, 
					treatment, 
					after_treatment
				FROM 
					hydroponic_system.unhealthy_plant_treatment
				where
					deleted_at is NULL;
		`

	rows, err := r.db.Raw(sqlScript).Rows()
	if err != nil {
		logger.Error("unhealthyPlantTreatmentRepository", "Failed to fetch GetUnhealthyPlantTreatment", map[string]string{
			"error": err.Error(),
		})
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&queryResult.ID,
			&queryResult.TowerId,
			&queryResult.Cycle,
			&queryResult.TowerHoleNo,
			&queryResult.Treatment,
			&queryResult.AfterTreatment,
		); err != nil {
			logger.Error("unhealthyPlantTreatmentRepository", "Failed to fetch scan GetUnhealthyPlantTreatment result", map[string]string{
				"error": err.Error(),
			})
			return nil, fmt.Errorf("error scanning row: %w", err)
		}

		result = append(result, *queryResult)
	}

	logger.Info("unhealthyPlantTreatmentRepository", "GetUnhealthyPlantTreatment fetched successfully", map[string]string{
		"fetched": strconv.Itoa(len(result)),
	})
	return &result, nil
}
