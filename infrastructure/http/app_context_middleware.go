package http

import (
	"context"
	"domain-driven-design-layout/domain/entities"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
)

const AppContextKey = "appContextKey"
const CorrelationIDHeader = "Correlation-Id"

func CreateAppContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		appCtx := CreateAppContext(c.Request.Context(), c.Request.Header.Get(CorrelationIDHeader))
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), AppContextKey, appCtx))
		c.Next()
	}
}

func CreateAppContext(parent context.Context, correlationID string) *entities.AppContext {
	if correlationID == "" {
		correlationID = uuid.New().String()
	}

	appCtx := &entities.AppContext{
		CorrelationID: correlationID,
		Logger: logrus.StandardLogger().WithFields(logrus.Fields{
			"correlationID": correlationID,
		}),
		Context: parent,
	}

	return appCtx
}

func GetAppContext(r *http.Request) *entities.AppContext {
	appCtx, ok := r.Context().Value(AppContextKey).(*entities.AppContext)
	if ok {
		return appCtx
	}

	return CreateAppContext(r.Context(), "")
}
