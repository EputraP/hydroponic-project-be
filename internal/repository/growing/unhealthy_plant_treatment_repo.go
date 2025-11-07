package repository

import (
	"gorm.io/gorm"
)

type UnhealthyPlantTreatmentRepository interface {
}

type unhealthyPlantTreatmentRepository struct {
	db *gorm.DB
}

func NewUnhealthyPlantTreatmentRepository(db *gorm.DB) UnhealthyPlantTreatmentRepository {
	return &unhealthyPlantTreatmentRepository{
		db: db,
	}
}
