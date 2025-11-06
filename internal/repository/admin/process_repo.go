package repository

import (
	"fmt"
	model "hydroponic-be/internal/model/admin"
	"hydroponic-be/internal/util/logger"
	"strconv"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProcessRepository interface {
	GetProcesses() (*[]model.Process, error)
	GetModules() (*[]model.Process, error)
	GetSubModulesByProcessId(processId *uuid.UUID) (*[]model.Process, error)
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
					deleted_at is null AND
					process_id IS NULL AND
					description != 'Module'
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

func (r *processRepository) GetModules() (*[]model.Process, error) {

	logger.Info("processRepository", "Fetching GetModules", map[string]string{})

	queryResult := &model.Process{}
	result := []model.Process{}

	sqlScript := `select
					id,
					process_name,
					description,
					process_id
				from hydroponic_system.process
				where
					deleted_at is null AND
					process_id IS NULL AND
					description = 'Module'

		`

	rows, err := r.db.Raw(sqlScript).Rows()
	if err != nil {
		logger.Error("processRepository", "Failed to fetch GetModules", map[string]string{
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
			&queryResult.ProcessId,
		); err != nil {
			logger.Error("processRepository", "Failed to fetch scan GetModules result", map[string]string{
				"error": err.Error(),
			})
			return nil, fmt.Errorf("error scanning row: %w", err)
		}

		result = append(result, *queryResult)
	}

	logger.Info("processRepository", "GetModules fetched successfully", map[string]string{
		"fetched": strconv.Itoa(len(result)),
	})
	return &result, nil

}

func (r *processRepository) GetSubModulesByProcessId(processId *uuid.UUID) (*[]model.Process, error) {

	logger.Info("processRepository", "Fetching GetSubModulesByProcessId", map[string]string{})

	queryResult := &model.Process{}
	result := []model.Process{}

	sqlScript := `select
					id,
					process_name,
					description,
					process_id
				from hydroponic_system.process
				where
					deleted_at is null AND
					description = 'Sub Module'
					AND process_id = '` + processId.String() + `'
		`

	rows, err := r.db.Raw(sqlScript).Rows()
	if err != nil {
		logger.Error("processRepository", "Failed to fetch GetSubModulesByProcessId", map[string]string{
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
			&queryResult.ProcessId,
		); err != nil {
			logger.Error("processRepository", "Failed to fetch scan GetSubModulesByProcessId result", map[string]string{
				"error": err.Error(),
			})
			return nil, fmt.Errorf("error scanning row: %w", err)
		}

		result = append(result, *queryResult)
	}

	logger.Info("processRepository", "GetSubModulesByProcessId fetched successfully", map[string]string{
		"fetched": strconv.Itoa(len(result)),
	})
	return &result, nil

}
