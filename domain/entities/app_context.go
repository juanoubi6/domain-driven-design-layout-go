package entities

import (
	"context"
	"github.com/sirupsen/logrus"
)

type ApplicationContext interface {
	context.Context
}

type AppContext struct {
	CorrelationID string
	Logger        logrus.FieldLogger
	context.Context
}
