package repository

import (
	"fmt"
	model "hydroponic-be/internal/model/admin"
	"hydroponic-be/internal/util/logger"
	"strconv"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UomRepository interface {
	GetUoMs() (*[]model.Uom, error)
	GetUoMById(input *uuid.UUID) (*[]model.Uom, error)
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

func (r *uomRepository) GetUoMById(input *uuid.UUID) (*[]model.Uom, error) {

	logger.Info("uomRepository", "Fetching GetUoMById", map[string]string{})

	queryResult := &model.Uom{}
	result := []model.Uom{}

	sqlScript := `select
					id,
					name,
					symbol,
					description
				from hydroponic_system.uom
				where
					deleted_at is null and id = '%s'
		`

	sqlScript = fmt.Sprintf(sqlScript, input.String())

	rows, err := r.db.Raw(sqlScript).Rows()
	if err != nil {
		logger.Error("uomRepository", "Failed to fetch GetUoMById", map[string]string{
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
			logger.Error("uomRepository", "Failed to fetch scan GetUoMById result", map[string]string{
				"error": err.Error(),
			})
			return nil, fmt.Errorf("error scanning row: %w", err)
		}

		result = append(result, *queryResult)
	}

	logger.Info("uomRepository", "GetUoMById fetched successfully", map[string]string{
		"fetched": strconv.Itoa(len(result)),
	})
	return &result, nil

}
