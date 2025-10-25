package repository

import (
	"gorm.io/gorm"
)

type PlantRepository interface {
	CreateFarm() string
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
