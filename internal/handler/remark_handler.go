package handler

import (
	"hydroponic-be/internal/service"
	"hydroponic-be/internal/util/logger"
	"hydroponic-be/internal/util/response"

	"github.com/gin-gonic/gin"
)

type RemarkHandler struct {
	remarkService service.RemarkService
}

type RemarkHandlerConfig struct {
	RemarkService service.RemarkService
}

func NewRemarkHandler(config RemarkHandlerConfig) *RemarkHandler {
	return &RemarkHandler{
		remarkService: config.RemarkService,
	}
}

func (h *RemarkHandler) GetRemarks(c *gin.Context) {

	logger.Info("remarkHandler", "Init GerRemarks handler", nil)

	resp, err := h.remarkService.GetRemarks()
	if err != nil {
		logger.Error("remarkHandler", "Failed to execute GetRemarks Service", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, err.Error())
		return
	}

	response.JSON(c, 200, "GetRemarks Success", resp)
}
