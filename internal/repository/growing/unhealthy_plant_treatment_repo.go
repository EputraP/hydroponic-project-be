package repository

import (
	"fmt"
	model "hydroponic-be/internal/model/growing"
	"hydroponic-be/internal/util/logger"

	"gorm.io/gorm"
)

type UnhealthyPlantTreatmentRepository interface {
	CreateUnhealthyPlantTreatment(input *model.UnhealthyPlantTreatment) (*model.UnhealthyPlantTreatment, error)
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
