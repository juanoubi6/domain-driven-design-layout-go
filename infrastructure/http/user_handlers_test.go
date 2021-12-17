package http

import (
	"bytes"
	"domain-driven-design-layout/domain"
	"domain-driven-design-layout/domain/entities"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var mockCreateUserAction = new(domain.CreateUserMock)
var mockFindUserByIdAction = new(domain.FindUserByIdMock)
var mockFindUsersByIdListAction = new(domain.FindUsersByIdListMock)
var mockUpdateUserAction = new(domain.UpdateUserMock)
var mockDeleteUserAction = new(domain.DeleteUserMock)

type UserHandlersTestSuite struct {
	suite.Suite
	router *gin.Engine
}

func (suite *UserHandlersTestSuite) SetupTest() {
	mockCreateUserAction = new(domain.CreateUserMock)
	mockFindUserByIdAction = new(domain.FindUserByIdMock)
	mockFindUsersByIdListAction = new(domain.FindUsersByIdListMock)
	mockUpdateUserAction = new(domain.UpdateUserMock)
	mockDeleteUserAction = new(domain.DeleteUserMock)

	router := gin.New()

	userHandlers := &UserHandlers{
		createUserAction:        mockCreateUserAction,
		findUserByIdAction:      mockFindUserByIdAction,
		findUsersByIdListAction: mockFindUsersByIdListAction,
		updateUserAction:        mockUpdateUserAction,
		deleteUserAction:        mockDeleteUserAction,
	}

	router.POST("/users-api/users", userHandlers.CreateUser)
	router.POST("/users-api/users/list", userHandlers.FindUsersByIdList)
	router.GET("/users-api/users/:id", userHandlers.FindUserById)
	router.PUT("/users-api/users/:id", userHandlers.UpdateUser)
	router.DELETE("/users-api/users/:id", userHandlers.DeleteUser)

	suite.router = router
}

func TestUserHandlersTestSuite(t *testing.T) {
	suite.Run(t, new(UserHandlersTestSuite))
}

func (suite *UserHandlersTestSuite) TestUserHandlers_CreateUser_SuccessReturns201() {
	w := httptest.NewRecorder()

	body, _ := json.Marshal(createUserBodyRequest())
	req, _ := http.NewRequest("POST", "/users-api/users", bytes.NewBuffer(body))

	expected := createUser()

	mockCreateUserAction.On("Execute", mock.Anything).Return(expected, nil)

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 201, w.Code)
	mockCreateUserAction.AssertExpectations(suite.T())
}

func (suite *UserHandlersTestSuite) TestUserHandlers_CreateUser_InvalidBodyReturns400() {
	w := httptest.NewRecorder()

	body, _ := json.Marshal(map[string]interface{}{
		"first_name": 33,
	})
	req, _ := http.NewRequest("POST", "/users-api/users", bytes.NewBuffer(body))

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 400, w.Code)
}

func (suite *UserHandlersTestSuite) TestUserHandlers_CreateUser_Returns400OnActionFailure() {
	w := httptest.NewRecorder()

	body, _ := json.Marshal(createUserBodyRequest())
	req, _ := http.NewRequest("POST", "/users-api/users", bytes.NewBuffer(body))

	mockCreateUserAction.On("Execute", mock.Anything).Return(createUser(), errors.New("error"))

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 400, w.Code)
	mockCreateUserAction.AssertExpectations(suite.T())
}

func (suite *UserHandlersTestSuite) TestUserHandlers_FindUserById_SuccessReturns200() {
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/users-api/users/1", nil)

	expected := createUser()

	mockFindUserByIdAction.On("Execute", int64(1)).Return(&expected, nil)

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 200, w.Code)
	mockFindUserByIdAction.AssertExpectations(suite.T())
}

func (suite *UserHandlersTestSuite) TestUserHandlers_FindUserById_Returns404WhenUserCouldNotBeFound() {
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/users-api/users/1", nil)

	mockFindUserByIdAction.On("Execute", int64(1)).Return(nil, nil)

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 404, w.Code)
	mockFindUserByIdAction.AssertExpectations(suite.T())
}

func (suite *UserHandlersTestSuite) TestUserHandlers_FindUserById_Returns400OnActionFailure() {
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/users-api/users/1", nil)

	mockFindUserByIdAction.On("Execute", int64(1)).Return(nil, errors.New("error"))

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 400, w.Code)
	mockFindUserByIdAction.AssertExpectations(suite.T())
}

func createUserBodyRequest() map[string]interface{} {
	return map[string]interface{}{
		"first_name": "First",
		"last_name":  "Last",
		"birth_date": "1995-07-20T00:00:00.000Z",
		"addresses": []interface{}{
			map[string]interface{}{
				"street": "street 1",
				"number": 1,
				"city":   "Argentina",
			},
			map[string]interface{}{
				"street": "street 2",
				"number": 2,
				"city":   nil,
			},
		},
	}
}

func createUser() entities.User {
	country := "Argentina"

	return entities.User{
		ID:        1,
		FirstName: "test",
		LastName:  "name",
		BirthDate: time.Now(),
		Addresses: []entities.Address{
			{
				ID:     1,
				UserID: 1,
				Street: "street 1",
				Number: 1,
				City:   &country,
			},
			{
				ID:     2,
				UserID: 2,
				Street: "street 1",
				Number: 1,
				City:   nil,
			},
		},
	}
}
