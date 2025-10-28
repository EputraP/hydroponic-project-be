package repository

import (
	"fmt"
	"hydroponic-be/internal/model"
	"hydroponic-be/internal/util/logger"
	"strconv"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RemarkRepository interface {
	GetRemarks() (*[]model.Remark, error)
	GetRemarksByProcessId(processId *uuid.UUID) (*[]model.Remark, error)
}

type remarkRepository struct {
	db *gorm.DB
}

func NewRemarkRepository(db *gorm.DB) RemarkRepository {
	return &remarkRepository{
		db: db,
	}
}

func (r *remarkRepository) GetRemarks() (*[]model.Remark, error) {

	logger.Info("remarkRepository", "Fetching GetRemarks", map[string]string{})

	queryResult := &model.Remark{}
	result := []model.Remark{}

	sqlScript := `select
					id,
					remark,
					description
				from hydroponic_system.remarks
				where
					deleted_at is null
		`

	rows, err := r.db.Raw(sqlScript).Rows()
	if err != nil {
		logger.Error("remarkRepository", "Failed to fetch GetRemarks", map[string]string{
			"error": err.Error(),
		})
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&queryResult.ID,
			&queryResult.Remark,
			&queryResult.Description,
		); err != nil {
			logger.Error("remarkRepository", "Failed to fetch scan GetRemarks result", map[string]string{
				"error": err.Error(),
			})
			return nil, fmt.Errorf("error scanning row: %w", err)
		}

		result = append(result, *queryResult)
	}

	logger.Info("remarkRepository", "GetRemarks fetched successfully", map[string]string{
		"fetched": strconv.Itoa(len(result)),
	})
	return &result, nil
}
func (r *remarkRepository) GetRemarksByProcessId(processId *uuid.UUID) (*[]model.Remark, error) {

	logger.Info("remarkRepository", "Fetching GetRemarksByProcessId", map[string]string{})

	queryResult := &model.Remark{}
	result := []model.Remark{}

	sqlScript := `select
					id,
					remark,
					description
				from hydroponic_system.remarks
				where
					deleted_at is null AND
					process_id = '%s'
		`

	sqlScript = fmt.Sprintf(sqlScript,
		processId.String(),
	)
	rows, err := r.db.Raw(sqlScript).Rows()
	if err != nil {
		logger.Error("remarkRepository", "Failed to fetch GetRemarksByProcessId", map[string]string{
			"error": err.Error(),
		})
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&queryResult.ID,
			&queryResult.Remark,
			&queryResult.Description,
		); err != nil {
			logger.Error("remarkRepository", "Failed to fetch scan GetRemarksByProcessId result", map[string]string{
				"error": err.Error(),
			})
			return nil, fmt.Errorf("error scanning row: %w", err)
		}

		result = append(result, *queryResult)
	}

	logger.Info("remarkRepository", "GetRemarksByProcessId fetched successfully", map[string]string{
		"fetched": strconv.Itoa(len(result)),
	})
	return &result, nil
}
