package http

import (
	"bytes"
	"domain-driven-design-layout/domain/entities"
	"domain-driven-design-layout/domain/mocks"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var mockCreateUserAction = new(mocks.CreateUserMock)
var mockFindUserByIdAction = new(mocks.FindUserByIdMock)
var mockFindUsersByIdListAction = new(mocks.FindUsersByIdListMock)
var mockUpdateUserAction = new(mocks.UpdateUserMock)
var mockDeleteUserAction = new(mocks.DeleteUserMock)

type UserHandlersTestSuite struct {
	suite.Suite
	router *gin.Engine
}

func (suite *UserHandlersTestSuite) SetupTest() {
	mockCreateUserAction = new(mocks.CreateUserMock)
	mockFindUserByIdAction = new(mocks.FindUserByIdMock)
	mockFindUsersByIdListAction = new(mocks.FindUsersByIdListMock)
	mockUpdateUserAction = new(mocks.UpdateUserMock)
	mockDeleteUserAction = new(mocks.DeleteUserMock)

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
	router.PUT("/users-api/users", userHandlers.UpdateUser)
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

	mockCreateUserAction.On("Execute", mock.AnythingOfType("*entities.AppContext"), mock.Anything).Return(expected, nil)

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

	mockCreateUserAction.On("Execute", mock.AnythingOfType("*entities.AppContext"), mock.Anything).Return(createUser(), errors.New("error"))

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 400, w.Code)
	mockCreateUserAction.AssertExpectations(suite.T())
}

func (suite *UserHandlersTestSuite) TestUserHandlers_FindUserById_Returns200OnSuccess() {
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/users-api/users/1", nil)

	expected := createUser()

	mockFindUserByIdAction.On("Execute", mock.AnythingOfType("*entities.AppContext"), int64(1)).Return(&expected, nil)

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 200, w.Code)
	mockFindUserByIdAction.AssertExpectations(suite.T())

	responseBody, err := io.ReadAll(w.Body)
	if err != nil {
		suite.T().FailNow()
	}

	var userResponse entities.User
	if err = json.Unmarshal(responseBody, &userResponse); err != nil {
		suite.T().FailNow()
	}

	assert.Equal(suite.T(), expected.ID, userResponse.ID)

}

func (suite *UserHandlersTestSuite) TestUserHandlers_FindUserById_Returns404WhenUserCouldNotBeFound() {
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/users-api/users/1", nil)

	mockFindUserByIdAction.On("Execute", mock.AnythingOfType("*entities.AppContext"), int64(1)).Return(nil, nil)

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 404, w.Code)
	mockFindUserByIdAction.AssertExpectations(suite.T())
}

func (suite *UserHandlersTestSuite) TestUserHandlers_FindUserById_Returns400OnActionFailure() {
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/users-api/users/1", nil)

	mockFindUserByIdAction.On("Execute", mock.AnythingOfType("*entities.AppContext"), int64(1)).Return(nil, errors.New("error"))

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 400, w.Code)
	mockFindUserByIdAction.AssertExpectations(suite.T())
}

func (suite *UserHandlersTestSuite) TestUserHandlers_FindUsersByIdList_Returns200OnSuccess() {
	w := httptest.NewRecorder()

	body, _ := json.Marshal(findUsersByIdListBodyRequest())
	req, _ := http.NewRequest("POST", "/users-api/users/list", bytes.NewBuffer(body))

	expected := []entities.User{createUser()}

	mockFindUsersByIdListAction.On("Execute", mock.AnythingOfType("*entities.AppContext"), []int64{1, 2, 3}).Return(expected, nil)

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 200, w.Code)
	mockFindUsersByIdListAction.AssertExpectations(suite.T())

	responseBody, err := io.ReadAll(w.Body)
	if err != nil {
		suite.T().FailNow()
	}

	var usersResponse []entities.User
	if err = json.Unmarshal(responseBody, &usersResponse); err != nil {
		suite.T().FailNow()
	}

	assert.Equal(suite.T(), len(expected), len(usersResponse))
}

func (suite *UserHandlersTestSuite) TestUserHandlers_FindUsersByIdList_Returns400OnActionFailure() {
	w := httptest.NewRecorder()

	body, _ := json.Marshal(findUsersByIdListBodyRequest())
	req, _ := http.NewRequest("POST", "/users-api/users/list", bytes.NewBuffer(body))

	mockFindUsersByIdListAction.On("Execute", mock.AnythingOfType("*entities.AppContext"), mock.Anything).Return(nil, errors.New("error"))

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 400, w.Code)
	mockFindUsersByIdListAction.AssertExpectations(suite.T())
}

func (suite *UserHandlersTestSuite) TestUserHandlers_FindUsersByIdList_InvalidBodyReturns400() {
	w := httptest.NewRecorder()

	body, _ := json.Marshal(map[string]interface{}{
		"invalid_value": 33,
	})
	req, _ := http.NewRequest("POST", "/users-api/users/list", bytes.NewBuffer(body))

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 400, w.Code)
}

func (suite *UserHandlersTestSuite) TestUserHandlers_UpdateUser_Returns200OnSuccess() {
	w := httptest.NewRecorder()

	body, _ := json.Marshal(updateUserBodyRequest())
	req, _ := http.NewRequest("PUT", "/users-api/users", bytes.NewBuffer(body))

	expected := createUser()

	mockUpdateUserAction.On("Execute", mock.AnythingOfType("*entities.AppContext"), mock.Anything).Return(expected, nil)

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 200, w.Code)
	mockUpdateUserAction.AssertExpectations(suite.T())

	responseBody, err := io.ReadAll(w.Body)
	if err != nil {
		suite.T().FailNow()
	}

	var userResponse entities.User
	if err = json.Unmarshal(responseBody, &userResponse); err != nil {
		suite.T().FailNow()
	}

	assert.Equal(suite.T(), expected.ID, userResponse.ID)
}

func (suite *UserHandlersTestSuite) TestUserHandlers_UpdateUser_InvalidBodyReturns400() {
	w := httptest.NewRecorder()

	body, _ := json.Marshal(map[string]interface{}{
		"invalid": 33,
	})
	req, _ := http.NewRequest("PUT", "/users-api/users", bytes.NewBuffer(body))

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 400, w.Code)
}

func (suite *UserHandlersTestSuite) TestUserHandlers_UpdateUser_Returns400OnActionFailure() {
	w := httptest.NewRecorder()

	body, _ := json.Marshal(updateUserBodyRequest())
	req, _ := http.NewRequest("PUT", "/users-api/users", bytes.NewBuffer(body))

	mockUpdateUserAction.On("Execute", mock.AnythingOfType("*entities.AppContext"), mock.Anything).Return(createUser(), errors.New("error"))

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 400, w.Code)
	mockUpdateUserAction.AssertExpectations(suite.T())
}

func (suite *UserHandlersTestSuite) TestUserHandlers_DeleteUser_Returns400OnActionFailure() {
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("DELETE", "/users-api/users/1", nil)

	mockDeleteUserAction.On("Execute", mock.AnythingOfType("*entities.AppContext"), int64(1)).Return(errors.New("error"))

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 400, w.Code)
	mockDeleteUserAction.AssertExpectations(suite.T())
}

func (suite *UserHandlersTestSuite) TestUserHandlers_DeleteUser_Returns200OnSuccess() {
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("DELETE", "/users-api/users/1", nil)

	mockDeleteUserAction.On("Execute", mock.AnythingOfType("*entities.AppContext"), int64(1)).Return(nil)

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 200, w.Code)
	mockDeleteUserAction.AssertExpectations(suite.T())
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

func findUsersByIdListBodyRequest() map[string]interface{} {
	return map[string]interface{}{
		"user_ids": []int64{1, 2, 3},
	}
}

func updateUserBodyRequest() map[string]interface{} {
	return map[string]interface{}{
		"id":         1,
		"first_name": "First",
		"last_name":  "Last",
		"birth_date": "1995-07-20T00:00:00.000Z",
	}
}
