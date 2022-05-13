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
)

type AddressHandlersTestSuite struct {
	suite.Suite
	router                    *gin.Engine
	mockCreateAddressAction   *mocks.CreateAddressMock
	mockDeleteAddressAction   *mocks.DeleteAddressMock
	mockFindAddressByIdAction *mocks.FindAddressByIdMock
}

func (suite *AddressHandlersTestSuite) SetupTest() {
	mockCreateAddressAction := new(mocks.CreateAddressMock)
	mockDeleteAddressAction := new(mocks.DeleteAddressMock)
	mockFindAddressByIdAction := new(mocks.FindAddressByIdMock)

	router := gin.New()

	addressHandlers := &AddressHandlers{
		createAddressAction: mockCreateAddressAction,
		deleteAddressAction: mockDeleteAddressAction,
		findAddressById:     mockFindAddressByIdAction,
	}

	router.GET("/users-api/addresses/:addressID", addressHandlers.FindAddressById)
	router.POST("/users-api/user/:userID/addresses", addressHandlers.CreateAddress)
	router.DELETE("/users-api/user/:userID/addresses/:addressID", addressHandlers.DeleteAddress)

	suite.router = router
	suite.mockCreateAddressAction = mockCreateAddressAction
	suite.mockDeleteAddressAction = mockDeleteAddressAction
	suite.mockFindAddressByIdAction = mockFindAddressByIdAction
}

func TestAddressHandlersTestSuite(t *testing.T) {
	suite.Run(t, new(AddressHandlersTestSuite))
}

func (suite *AddressHandlersTestSuite) TestAddressHandlers_CreateAddress_SuccessReturns201() {
	w := httptest.NewRecorder()

	body, _ := json.Marshal(createAddressBodyRequest())
	req, _ := http.NewRequest("POST", "/users-api/user/1/addresses", bytes.NewBuffer(body))

	expected := createAddress()

	suite.mockCreateAddressAction.On("Execute", mock.AnythingOfType("*entities.AppContext"), int64(1), mock.Anything).Return(expected, nil)

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 201, w.Code)
	suite.mockCreateAddressAction.AssertExpectations(suite.T())
}

func (suite *AddressHandlersTestSuite) TestAddressHandlers_CreateAddress_InvalidBodyReturns400() {
	w := httptest.NewRecorder()

	body, _ := json.Marshal(map[string]interface{}{
		"first_name": 33,
	})
	req, _ := http.NewRequest("POST", "/users-api/user/1/addresses", bytes.NewBuffer(body))

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 400, w.Code)
}

func (suite *AddressHandlersTestSuite) TestAddressHandlers_CreateAddress_Returns400OnActionFailure() {
	w := httptest.NewRecorder()

	body, _ := json.Marshal(createAddressBodyRequest())
	req, _ := http.NewRequest("POST", "/users-api/user/1/addresses", bytes.NewBuffer(body))

	suite.mockCreateAddressAction.On("Execute", mock.AnythingOfType("*entities.AppContext"), int64(1), mock.Anything).Return(entities.Address{}, errors.New("error"))

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 400, w.Code)
	suite.mockCreateAddressAction.AssertExpectations(suite.T())
}

func (suite *AddressHandlersTestSuite) TestAddressHandlers_DeleteAddress_Returns400OnActionFailure() {
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("DELETE", "/users-api/user/1/addresses/10", nil)

	suite.mockDeleteAddressAction.On("Execute", mock.AnythingOfType("*entities.AppContext"), int64(10)).Return(errors.New("error"))

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 400, w.Code)
	suite.mockDeleteAddressAction.AssertExpectations(suite.T())
}

func (suite *AddressHandlersTestSuite) TestAddressHandlers_DeleteAddress_Returns200OnSuccess() {
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("DELETE", "/users-api/user/1/addresses/10", nil)

	suite.mockDeleteAddressAction.On("Execute", mock.AnythingOfType("*entities.AppContext"), int64(10)).Return(nil)

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 200, w.Code)
	suite.mockDeleteAddressAction.AssertExpectations(suite.T())
}

func (suite *AddressHandlersTestSuite) TestAddressHandlers_FindAddressById_Returns200OnSuccess() {
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/users-api/addresses/1", nil)

	expected := createAddress()

	suite.mockFindAddressByIdAction.On("Execute", mock.AnythingOfType("*entities.AppContext"), int64(1)).Return(&expected, nil)

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 200, w.Code)
	suite.mockFindAddressByIdAction.AssertExpectations(suite.T())

	responseBody, err := io.ReadAll(w.Body)
	if err != nil {
		suite.T().FailNow()
	}

	var addressResponse entities.Address
	if err = json.Unmarshal(responseBody, &addressResponse); err != nil {
		suite.T().FailNow()
	}

	assert.Equal(suite.T(), expected.ID, addressResponse.ID)
}

func (suite *AddressHandlersTestSuite) TestAddressHandlers_FindAddressById_Returns404WhenAddressCouldNotBeFound() {
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/users-api/addresses/1", nil)

	suite.mockFindAddressByIdAction.On("Execute", mock.AnythingOfType("*entities.AppContext"), int64(1)).Return(nil, nil)

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 404, w.Code)
	suite.mockFindAddressByIdAction.AssertExpectations(suite.T())
}

func (suite *AddressHandlersTestSuite) TestAddressHandlers_FindAddressById_Returns400OnActionFailure() {
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/users-api/addresses/1", nil)

	suite.mockFindAddressByIdAction.On("Execute", mock.AnythingOfType("*entities.AppContext"), int64(1)).Return(nil, errors.New("error"))

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 400, w.Code)
	suite.mockFindAddressByIdAction.AssertExpectations(suite.T())
}

func createAddressBodyRequest() map[string]interface{} {
	return map[string]interface{}{
		"street": "street 2",
		"number": 2,
		"city":   nil,
	}
}

func createAddress() entities.Address {
	return entities.Address{
		ID:     1,
		UserID: 1,
		Street: "street 1",
		Number: 1,
		City:   nil,
	}
}
