package http

import (
	"domain-driven-design-layout/domain/actions/users"
	"domain-driven-design-layout/infrastructure/builder"
	"domain-driven-design-layout/infrastructure/http/requests"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandlers struct {
	createUserAction users.CreateUser
}

func NewUserHandlers(actions *builder.Actions) (*UserHandlers, error) {
	return &UserHandlers{createUserAction: actions.CreateUser}, nil
}

func (r *UserHandlers) CreateUser(c *gin.Context) {
	var request requests.CreateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := r.createUserAction.Execute(request.ToUserPrototype())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}
