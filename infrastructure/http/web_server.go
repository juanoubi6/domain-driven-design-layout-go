package http

import (
	"domain-driven-design-layout/infrastructure/config"
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
}

func (w *WebServer) setHandlers() {
	router := w.engine.Group(basePath)
	router.GET("health", w.health.Status)
	router.POST("/users", w.handlers.UserHandlers.CreateUser)
}
