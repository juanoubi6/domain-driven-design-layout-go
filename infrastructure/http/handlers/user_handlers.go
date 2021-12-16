package handlers

import (
	"domain-driven-design-layout/infrastructure/builder"
	"domain-driven-design-layout/infrastructure/config"
	"domain-driven-design-layout/infrastructure/http/requests"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandlers struct {
	actions *builder.Actions
	config  config.WebConfig
}

func NewUserHandlers(actions *builder.Actions, config config.WebConfig) (*UserHandlers, error) {
	return &UserHandlers{
		actions: actions,
		config:  config,
	}, nil
}

func (r *UserHandlers) CreateUser(c *gin.Context) {
	var request requests.CreateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := r.actions.CreateUser.Execute(request.ToUserPrototype())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}
