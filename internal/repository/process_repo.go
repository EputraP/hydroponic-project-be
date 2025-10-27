package repository

import (
	"fmt"
	"hydroponic-be/internal/model"
	"hydroponic-be/internal/util/logger"
	"strconv"

	"gorm.io/gorm"
)

type ProcessRepository interface {
	GetProcesses() (*[]model.Process, error)
}

type processRepository struct {
	db *gorm.DB
}

func NewProcessRepository(db *gorm.DB) ProcessRepository {
	return &processRepository{
		db: db,
	}
}

func (r *processRepository) GetProcesses() (*[]model.Process, error) {

	logger.Info("processRepository", "Fetching GerProcesses", map[string]string{})

	queryResult := &model.Process{}
	result := []model.Process{}

	sqlScript := `select
					id,
					process_name,
					description
				from hydroponic_system.process
				where
					deleted_at is null
		`

	rows, err := r.db.Raw(sqlScript).Rows()
	if err != nil {
		logger.Error("processRepository", "Failed to fetch GetProcesses", map[string]string{
			"error": err.Error(),
		})
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(
			&queryResult.ID,
			&queryResult.ProcessName,
			&queryResult.Description,
		); err != nil {
			logger.Error("processRepository", "Failed to fetch scan GetProcesses result", map[string]string{
				"error": err.Error(),
			})
			return nil, fmt.Errorf("error scanning row: %w", err)
		}

		result = append(result, *queryResult)
	}

	logger.Info("processRepository", "GetProcesses fetched successfully", map[string]string{
		"fetched": strconv.Itoa(len(result)),
	})
	return &result, nil

}
