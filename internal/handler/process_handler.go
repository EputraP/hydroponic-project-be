package handler

import (
	"hydroponic-be/internal/service"
	"hydroponic-be/internal/util/logger"
	"hydroponic-be/internal/util/response"

	"github.com/gin-gonic/gin"
)

type ProcessHandler struct {
	processService service.ProcessService
}

type ProcessHandlerConfig struct {
	ProcessService service.ProcessService
}

func NewProcessHandler(config ProcessHandlerConfig) *ProcessHandler {
	return &ProcessHandler{
		processService: config.ProcessService,
	}
}

func (h *ProcessHandler) GetProcesses(c *gin.Context) {

	logger.Info("processHandler", "Init GetProcesses handler", nil)

	resp, err := h.processService.GetProcesses()
	if err != nil {
		logger.Error("processHandler", "Failed to execute GetProcesses Service", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, err.Error())
		return
	}

	response.JSON(c, 200, "GetProcesses Success", resp)
}
