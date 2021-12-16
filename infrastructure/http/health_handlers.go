package http

import (
	"github.com/gin-gonic/gin"
)

type HealthHandler struct{}

func NewHealthHandler() (*HealthHandler, error) {
	return &HealthHandler{}, nil
}

func (h *HealthHandler) Status(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "OK",
	})
}
