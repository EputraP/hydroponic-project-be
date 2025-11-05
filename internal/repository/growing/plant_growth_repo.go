package repository

import (
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
