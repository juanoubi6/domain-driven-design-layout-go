package http

import (
	"domain-driven-design-layout/infrastructure/builder"
	"domain-driven-design-layout/infrastructure/config"
	"domain-driven-design-layout/infrastructure/http/handlers"
)

type HttpHandlers struct {
	HealthHandler *handlers.HealthHandler
	UserHandlers  *handlers.UserHandlers
}

func CreateHttpHandlers(actions *builder.Actions, config config.WebConfig) (*HttpHandlers, error) {
	userHandlers, err := handlers.NewUserHandlers(actions, config)
	if err != nil {
		return nil, err
	}

	healthHandler := handlers.HealthHandler{}

	return &HttpHandlers{
		HealthHandler: &healthHandler,
		UserHandlers:  userHandlers,
	}, nil
}
