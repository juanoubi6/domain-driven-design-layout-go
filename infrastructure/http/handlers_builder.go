package http

import (
	"domain-driven-design-layout/infrastructure/builder"
)

type HttpHandlers struct {
	HealthHandler *HealthHandler
	UserHandlers  *UserHandlers
}

func CreateHttpHandlers(actions *builder.Actions) (*HttpHandlers, error) {
	userHandlers, err := NewUserHandlers(actions)
	if err != nil {
		return nil, err
	}

	healthHandler, err := NewHealthHandler()
	if err != nil {
		return nil, err
	}

	return &HttpHandlers{
		HealthHandler: healthHandler,
		UserHandlers:  userHandlers,
	}, nil
}
