package handler

import (
	errs "hydroponic-be/internal/errors"
	"hydroponic-be/internal/service"
	"hydroponic-be/internal/util/logger"
	"hydroponic-be/internal/util/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func (h *ProcessHandler) GetModules(c *gin.Context) {

	logger.Info("processHandler", "Init GetModules handler", nil)

	resp, err := h.processService.GetModules()
	if err != nil {
		logger.Error("processHandler", "Failed to execute GetModules Service", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, err.Error())
		return
	}

	response.JSON(c, 200, "GetModules Success", resp)
}

func (h *ProcessHandler) GetSubModules(c *gin.Context) {

	logger.Info("processHandler", "Init GetSubModules handler", nil)

	processId := c.Param("processId")
	processIdParsed, paramErr := uuid.Parse(processId)
	if paramErr != nil {
		logger.Error("processHandler", "Invalid process ID parameter", map[string]string{
			"processId": processId,
		})
		response.Error(c, 400, errs.InvalidUUIDParamFormat.Error())
		return
	}

	resp, err := h.processService.GetSubModules(&processIdParsed)
	if err != nil {
		logger.Error("processHandler", "Failed to execute GetSubModules Service", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, err.Error())
		return
	}

	response.JSON(c, 200, "GetSubModules Success", resp)
}
