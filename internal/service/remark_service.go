package service

import (
	"hydroponic-be/internal/model"
	"hydroponic-be/internal/repository"
	"hydroponic-be/internal/util/logger"
)

type RemarkService interface {
	GetRemarks() (*[]model.Remark, error)
}

type remarkService struct {
	remarkRepo repository.RemarkRepository
}

type RemarkServiceConfig struct {
	RemarkRepo repository.RemarkRepository
}

func NewRemarkService(config RemarkServiceConfig) RemarkService {
	return &remarkService{
		remarkRepo: config.RemarkRepo,
	}
}

func (s *remarkService) GetRemarks() (*[]model.Remark, error) {
	logger.Info("remarkService", "Init GetRemarks Service", nil)

	res, err := s.remarkRepo.GetRemarks()
	if err != nil {
		logger.Error("remarkService", "Failed to fetch GetRemarks Repo", map[string]string{
			"error": err.Error(),
		})
		return nil, err
	}

	logger.Info("remarkService", "Finished GetRemarks Service", map[string]string{})
	return res, nil
}
