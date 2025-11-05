package service

import (
	"hydroponic-be/internal/model"
	repository "hydroponic-be/internal/repository/admin"
	"hydroponic-be/internal/util/logger"
)

type UomService interface {
	GetUoms() (*[]model.Uom, error)
}

type uomService struct {
	uomRepo repository.UomRepository
}

type UomServiceConfig struct {
	UomRepo repository.UomRepository
}

func NewUomService(config UomServiceConfig) UomService {
	return &uomService{
		uomRepo: config.UomRepo,
	}
}

func (s *uomService) GetUoms() (*[]model.Uom, error) {
	logger.Info("uomService", "Init GetUoms Service", nil)

	res, err := s.uomRepo.GetUoMs()
	if err != nil {
		logger.Error("uomService", "Failed to fetch GetUoms Repo", map[string]string{
			"error": err.Error(),
		})
		return nil, err
	}

	logger.Info("uomService", "Finished GetUoms Service", map[string]string{})
	return res, nil
}
