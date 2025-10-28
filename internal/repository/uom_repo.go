package repository

import (
	"fmt"
	"hydroponic-be/internal/model"
	"hydroponic-be/internal/util/logger"
	"strconv"

	"gorm.io/gorm"
)

type UomRepository interface {
	GetUoMs() (*[]model.Uom, error)
}

type uomRepository struct {
	db *gorm.DB
}

func NewUomRepository(db *gorm.DB) UomRepository {
	return &uomRepository{
		db: db,
	}
}

func (r *uomRepository) GetUoMs() (*[]model.Uom, error) {

	logger.Info("uomRepository", "Fetching GetUoMs", map[string]string{})

	queryResult := &model.Uom{}
	result := []model.Uom{}

	sqlScript := `select
					id,
					name,
					symbol,
					description
				from hydroponic_system.uom
				where
					deleted_at is null
		`

	rows, err := r.db.Raw(sqlScript).Rows()
	if err != nil {
		logger.Error("uomRepository", "Failed to fetch GetUoMs", map[string]string{
			"error": err.Error(),
		})
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&queryResult.ID,
			&queryResult.Name,
			&queryResult.Symbol,
			&queryResult.Description,
		); err != nil {
			logger.Error("uomRepository", "Failed to fetch scan GetUoMs result", map[string]string{
				"error": err.Error(),
			})
			return nil, fmt.Errorf("error scanning row: %w", err)
		}

		result = append(result, *queryResult)
	}

	logger.Info("uomRepository", "GetUoMs fetched successfully", map[string]string{
		"fetched": strconv.Itoa(len(result)),
	})
	return &result, nil

}
