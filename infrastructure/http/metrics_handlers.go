package http

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

type MetricsHandler struct {
	prometheusHandler http.Handler
}

func NewMetricsHandler() (*MetricsHandler, error) {
	return &MetricsHandler{
		prometheusHandler: promhttp.Handler(),
	}, nil
}

func (h *MetricsHandler) Metrics(c *gin.Context) {
	h.prometheusHandler.ServeHTTP(c.Writer, c.Request)
}
