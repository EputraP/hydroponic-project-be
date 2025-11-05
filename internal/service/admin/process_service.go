package service

import (
	"hydroponic-be/internal/dto"
	errs "hydroponic-be/internal/errors"
	"hydroponic-be/internal/model"
	repository "hydroponic-be/internal/repository/admin"
	"hydroponic-be/internal/util/logger"

	"github.com/google/uuid"
)

type ProcessService interface {
	GetProcesses() (*[]model.Process, error)
	GetModules() (*[]model.Process, error)
	GetSubModules(processId *uuid.UUID) ([]*dto.IdNameAndDescriptionResponse, error)
}

type processService struct {
	processRepo repository.ProcessRepository
}

type ProcessServiceConfig struct {
	ProcessRepo repository.ProcessRepository
}

func NewProcessService(config ProcessServiceConfig) ProcessService {
	return &processService{
		processRepo: config.ProcessRepo,
	}
}

func (s *processService) GetProcesses() (*[]model.Process, error) {
	logger.Info("processService", "Init GetProcesses Service", nil)

	res, err := s.processRepo.GetProcesses()
	if err != nil {
		logger.Error("processService", "Failed to fetch GetProcesses Repo", map[string]string{
			"error": err.Error(),
		})
		return nil, err
	}

	logger.Info("processService", "Finished GetProcesses Service", map[string]string{})
	return res, nil
}

func (s *processService) GetModules() (*[]model.Process, error) {
	logger.Info("processService", "Init GetModules Service", nil)

	res, err := s.processRepo.GetModules()
	if err != nil {
		logger.Error("processService", "Failed to fetch GetModules Repo", map[string]string{
			"error": err.Error(),
		})
		return nil, err
	}

	logger.Info("processService", "Finished GetModules Service", map[string]string{})
	return res, nil
}

func (s *processService) GetSubModules(processId *uuid.UUID) ([]*dto.IdNameAndDescriptionResponse, error) {
	logger.Info("processService", "Init GetSubModules Service", nil)

	resVal := []*dto.IdNameAndDescriptionResponse{}

	res, err := s.processRepo.GetSubModulesByProcessId(processId)
	if err != nil {
		logger.Error("processService", "Failed to fetch GetSubModules Repo", map[string]string{
			"error": err.Error(),
		})
		return nil, err
	}
	if len(*res) == 0 {
		return nil, errs.CantFindAnySubModules
	}

	for i := 0; i < len(*res); i++ {
		resVal = append(resVal, &dto.IdNameAndDescriptionResponse{
			ID:          (*res)[i].ID,
			Name:        (*res)[i].ProcessName,
			Description: (*res)[i].Description,
		})
	}

	logger.Info("processService", "Finished GetSubModules Service", map[string]string{})
	return resVal, nil
}
