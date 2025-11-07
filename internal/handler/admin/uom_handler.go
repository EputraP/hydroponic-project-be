package handler

import (
	service "hydroponic-be/internal/service/admin"
	"hydroponic-be/internal/util/logger"
	"hydroponic-be/internal/util/response"

	"github.com/gin-gonic/gin"
)

type UomHandler struct {
	uomService service.UomService
}

type UomHandlerConfig struct {
	UomService service.UomService
}

func NewUomHandler(config UomHandlerConfig) *UomHandler {
	return &UomHandler{
		uomService: config.UomService,
	}
}

func (h *UomHandler) GetUoms(c *gin.Context) {

	logger.Info("uomHandler", "Init GetUoms handler", nil)

	resp, err := h.uomService.GetUoms()
	if err != nil {
		logger.Error("uomHandler", "Failed to execute GetUoms Service", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, err.Error())
		return
	}

	response.JSON(c, 200, "GetUoms Success", resp)
}
