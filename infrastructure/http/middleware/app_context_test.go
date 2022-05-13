package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

var expectedCorrelationID = ""

type AppContextMiddlewareTestSuite struct {
	suite.Suite
	router *gin.Engine
}

func (s *AppContextMiddlewareTestSuite) SetupSuite() {
	ginEngine := gin.New()
	ginEngine.Use(CreateAppContextMiddleware())

	ginEngine.GET("/testPath", func(c *gin.Context) {
		appCtx := GetAppContext(c.Request)
		expectedCorrelationID = appCtx.GetCorrelationID()

		c.JSON(http.StatusOK, "OK")
	})

	s.router = ginEngine
}

func (s *AppContextMiddlewareTestSuite) SetupTest() {
	expectedCorrelationID = ""
}

func TestAddressHandlersTestSuite(t *testing.T) {
	suite.Run(t, new(AppContextMiddlewareTestSuite))
}

func (s *AppContextMiddlewareTestSuite) Test_CreateAppContextMiddleware_DecoratesHandler() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/testPath", http.NoBody)

	s.router.ServeHTTP(w, req)

	assert.Equal(s.T(), http.StatusOK, w.Code)
	assert.NotEqualf(s.T(), "", expectedCorrelationID, "Should not be equal")
}

func (s *AppContextMiddlewareTestSuite) Test_CreateAppContextMiddleware_UsesCorrelationIDFromRequestHeader() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/testPath", http.NoBody)
	req.Header.Set(CorrelationIDHeader, "someCorrelationID")

	s.router.ServeHTTP(w, req)

	assert.Equal(s.T(), http.StatusOK, w.Code)
	assert.Equal(s.T(), "someCorrelationID", expectedCorrelationID, "Should not be equal")
}
