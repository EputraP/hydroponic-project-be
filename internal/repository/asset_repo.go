package repository

import (
	"gorm.io/gorm"
)

type AssetRepository interface {
}

type assetRepository struct {
	db *gorm.DB
}

func NewAssetRepository(db *gorm.DB) RemarkRepository {
	return &remarkRepository{
		db: db,
	}
}
