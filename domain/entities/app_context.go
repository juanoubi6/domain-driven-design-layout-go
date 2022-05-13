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

func (appCtx *AppContext) GetCorrelationID() string {
	return appCtx.CorrelationID
}

func (appCtx *AppContext) GetLogger() logrus.FieldLogger {
	return appCtx.Logger
}

func CreateEmptyAppContext() AppContext {
	return AppContext{}
}
