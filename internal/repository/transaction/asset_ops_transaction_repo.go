package repository

import (
	"fmt"
	model "hydroponic-be/internal/model/transaction"
	"hydroponic-be/internal/util/logger"

	"gorm.io/gorm"
)

type AssetOpsTransactionRepository interface {
	CreateAssetOpsTransaction(input *model.AssetOpsTransaction) (*model.AssetOpsTransaction, error)
}

type assetOpsTransactionRepository struct {
	db *gorm.DB
}

func NewAssetOpsTransactionRepository(db *gorm.DB) AssetOpsTransactionRepository {
	return &assetOpsTransactionRepository{
		db: db,
	}
}

func (r *assetOpsTransactionRepository) CreateAssetOpsTransaction(input *model.AssetOpsTransaction) (*model.AssetOpsTransaction, error) {

	logger.Info("assetOpsTransactionRepository", "Inserting CreateAssetOpsTransaction", map[string]string{})

	sqlScript := `INSERT INTO hydroponic_system.assets_ops_transaction
					(asset_id, tower_id, "cycle", value, created_at)
				VALUES('%s', '%s', %d, %d, NOW());
		`
	sqlScript = fmt.Sprintf(sqlScript,
		input.AssetId,
		input.TowerId,
		input.Cycle,
		input.Value,
	)

	rows, err := r.db.Raw(sqlScript).Rows()
	if err != nil {
		logger.Error("assetOpsTransactionRepository", "Failed to insert CreateAssetOpsTransaction", map[string]string{
			"error": err.Error(),
		})
		return nil, err
	}
	defer rows.Close()

	logger.Info("assetOpsTransactionRepository", "CreateAssetOpsTransaction inserted successfully", map[string]string{})
	return input, nil
}
