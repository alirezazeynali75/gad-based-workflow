package api

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)


type Handlers struct {
	logger *slog.Logger
}


func NewHandlers(logger *slog.Logger) *Handlers {
	return &Handlers{
		logger: logger.With("handlers"),
	}
}


func (h *Handlers) RegisterRoutes(router *gin.Engine) {
    // Register routes for all handlers
    router.POST("/trim", h.Trim)
    router.POST("/uppercase", h.Uppercase)
    router.POST("/uppercase-with-increase", h.UppercaseWithIncrease)
    router.POST("/all", h.All)
}