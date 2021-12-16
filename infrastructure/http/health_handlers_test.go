package http

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func createHealthHandlerRouter() *gin.Engine {
	router := gin.New()

	healthHandlers := &HealthHandler{}

	router.GET("/users-api/health", healthHandlers.Status)

	return router
}

func TestHealthHandler_Status_Success(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users-api/health", nil)

	createHealthHandlerRouter().ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
