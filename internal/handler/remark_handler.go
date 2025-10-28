package handler

import (
	"fmt"
	"hydroponic-be/internal/service"
	"hydroponic-be/internal/util/logger"
	"hydroponic-be/internal/util/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func (h *RemarkHandler) GetRemarksByProcessId(c *gin.Context) {

	logger.Info("remarkHandler", "Init GetRemarksByProcessId handler", nil)

	processId := c.Params.ByName("processId")

	fmt.Println("processId " + processId)

	processIdUUID, err := uuid.Parse(processId)
	if err != nil {
		logger.Error("remarkHandler", "Invalid processId parameter", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, "Invalid processId parameter")
		return
	}

	resp, err := h.remarkService.GetRemarksByProcessId(&processIdUUID)
	if err != nil {
		logger.Error("remarkHandler", "Failed to execute GetRemarks Service", map[string]string{
			"error": err.Error(),
		})
		response.Error(c, 400, err.Error())
		return
	}

	response.JSON(c, 200, "GetRemarks Success", resp)
}
