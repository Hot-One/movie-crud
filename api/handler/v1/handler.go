package v1

import (
	"movie-crud/api/status_http"
	"movie-crud/config"
	"movie-crud/pkg/logger"
	"movie-crud/storage"

	"github.com/gin-gonic/gin"
)

type handlerV1 struct {
	log     logger.Logger
	cfg     config.Config
	storage storage.IStorage
}

type HandlerV1Config struct {
	Logger  logger.Logger
	Cfg     config.Config
	Storage storage.IStorage
}

func New(c *HandlerV1Config) *handlerV1 {
	return &handlerV1{
		log:     c.Logger,
		cfg:     c.Cfg,
		storage: c.Storage,
	}
}

func (h *handlerV1) handleResponse(c *gin.Context, status status_http.Status, data any) {
	switch code := status.Code; {
	case code < 300:
		h.log.Info(
			"response",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
		)
	case code < 400:
		h.log.Warn(
			"response",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
		)
	default:
		h.log.Error(
			"response",
			logger.Int("code", status.Code),
			logger.String("status", status.Status),
			logger.Any("description", status.Description),
		)
	}

	c.JSON(status.Code, status_http.Response{
		Status:      status.Status,
		Description: status.Description,
		Data:        data,
	})
}

func (h *handlerV1) getDefaultLimit(limit int32) int32 {
	if limit == 0 {
		return 10
	}

	return limit
}

func (h *handlerV1) getDefaultOffset(offset int32) int32 {
	if offset == 0 {
		return 0
	}

	return offset
}
