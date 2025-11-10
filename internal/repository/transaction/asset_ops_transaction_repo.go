package repository

import (
	"gorm.io/gorm"
)

type AssetOpsTransactionRepository interface {
}

type assetOpsTransactionRepository struct {
	db *gorm.DB
}

func NewAssetOpsTransactionRepository(db *gorm.DB) AssetOpsTransactionRepository {
	return &assetOpsTransactionRepository{
		db: db,
	}
}
