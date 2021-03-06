package http

import (
	"domain-driven-design-layout/infrastructure/config"
	"domain-driven-design-layout/infrastructure/http/middleware"
	"github.com/gin-gonic/gin"
)

const basePath = "users-api"

type WebServer struct {
	handlers *HttpHandlers
	health   *HealthHandler
	config   config.WebConfig
	engine   *gin.Engine
}

func NewWebServer(config config.WebConfig, handlers *HttpHandlers) (*WebServer, error) {
	return &WebServer{
		handlers: handlers,
		health:   handlers.HealthHandler,
		config:   config,
		engine:   gin.New(),
	}, nil
}

func (w *WebServer) BuildRouter() *gin.Engine {
	w.setMiddleware()
	w.setHandlers()
	return w.engine
}

func (w *WebServer) Start() error {
	w.setMiddleware()
	w.setHandlers()
	return w.engine.Run(w.config.Address)
}

func (w *WebServer) setMiddleware() {
	w.engine.Use(gin.LoggerWithWriter(gin.DefaultWriter, basePath+"/health"))
	w.engine.Use(middleware.CreateAppContextMiddleware())
	w.engine.Use(middleware.CreateResponseDurationMiddleware())
}

func (w *WebServer) setHandlers() {
	// Define metrics handler
	w.engine.GET("/metrics", w.handlers.MetricsHandler.Metrics)

	router := w.engine.Group(basePath)
	router.GET("/health", w.health.Status)

	// User routes
	router.POST("/users", w.handlers.UserHandlers.CreateUser)
	router.POST("/users/list", w.handlers.UserHandlers.FindUsersByIdList)
	router.GET("/users/:id", w.handlers.UserHandlers.FindUserById)
	router.PUT("/users", w.handlers.UserHandlers.UpdateUser)
	router.DELETE("/users/:id", w.handlers.UserHandlers.DeleteUser)

	// Address routes
	router.GET("/addresses/:addressID", w.handlers.AddressHandlers.FindAddressById)
	router.POST("/user/:userID/addresses", w.handlers.AddressHandlers.CreateAddress)
	router.DELETE("/user/:userID/addresses/:addressID", w.handlers.AddressHandlers.DeleteAddress)
}
