package service

import (
	"hydroponic-be/internal/model"
	"hydroponic-be/internal/repository"
	"hydroponic-be/internal/util/logger"
)

type ProcessService interface {
	GetProcesses() (*[]model.Process, error)
	GetModules() (*[]model.Process, error)
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
