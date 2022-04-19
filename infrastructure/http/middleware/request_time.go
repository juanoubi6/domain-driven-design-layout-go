package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

func CreateResponseDurationMiddleware() gin.HandlerFunc {
	buckets := []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10}

	responseDurationHistogram := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_server_request_duration_seconds",
		Help:    "Response time for handler in seconds",
		Buckets: buckets,
	}, []string{"route", "method"})

	prometheus.MustRegister(responseDurationHistogram)

	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)

		route := c.FullPath()
		method := c.Request.Method

		responseDurationHistogram.WithLabelValues(route, method).Observe(duration.Seconds())
	}
}
