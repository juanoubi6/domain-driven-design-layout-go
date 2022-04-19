package http

import (
	"domain-driven-design-layout/domain/actions/users"
	"domain-driven-design-layout/infrastructure/builder"
	"domain-driven-design-layout/infrastructure/http/requests"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserHandlers struct {
	createUserAction        users.CreateUser
	findUserByIdAction      users.FindUserById
	findUsersByIdListAction users.FindUsersByIdList
	updateUserAction        users.UpdateUser
	deleteUserAction        users.DeleteUser
}

func NewUserHandlers(actions *builder.Actions) (*UserHandlers, error) {
	return &UserHandlers{
		createUserAction:        actions.CreateUser,
		findUserByIdAction:      actions.FindUserById,
		findUsersByIdListAction: actions.FindUsersByIdList,
		updateUserAction:        actions.UpdateUser,
		deleteUserAction:        actions.DeleteUser,
	}, nil
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

func (r *UserHandlers) FindUserById(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id value in URL"})
		return
	}

	appCtx := GetAppContext(c.Request)
	appCtx.Logger.WithField("userID", userId).Info("Received request to find user by ID")

	user, err := r.findUserByIdAction.Execute(int64(userId))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User could not be found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (r *UserHandlers) FindUsersByIdList(c *gin.Context) {
	var request requests.FindUsersByIdListRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userList, err := r.findUsersByIdListAction.Execute(request.UserIDs)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userList)
}

func (r *UserHandlers) UpdateUser(c *gin.Context) {
	var request requests.UpdateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedUser, err := r.updateUserAction.Execute(request.ToUser())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

func (r *UserHandlers) DeleteUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id value in URL"})
		return
	}

	if err := r.deleteUserAction.Execute(int64(userId)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}
