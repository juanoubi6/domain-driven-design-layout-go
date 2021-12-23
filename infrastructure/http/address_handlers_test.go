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
)

var mockCreateAddressAction = new(domain.CreateAddressMock)
var mockDeleteAddressAction = new(domain.DeleteAddressMock)

type AddressHandlersTestSuite struct {
	suite.Suite
	router *gin.Engine
}

func (suite *AddressHandlersTestSuite) SetupTest() {
	mockCreateAddressAction = new(domain.CreateAddressMock)
	mockDeleteAddressAction = new(domain.DeleteAddressMock)

	router := gin.New()

	addressHandlers := &AddressHandlers{
		createAddressAction: mockCreateAddressAction,
		deleteAddressAction: mockDeleteAddressAction,
	}

	router.POST("/users-api/users/:userID/addresses", addressHandlers.CreateAddress)
	router.DELETE("/users-api/users/:userID/addresses/:addressID", addressHandlers.DeleteAddress)

	suite.router = router
}

func TestAddressHandlersTestSuite(t *testing.T) {
	suite.Run(t, new(AddressHandlersTestSuite))
}

func (suite *AddressHandlersTestSuite) TestAddressHandlers_CreateAddress_SuccessReturns201() {
	w := httptest.NewRecorder()

	body, _ := json.Marshal(createAddressBodyRequest())
	req, _ := http.NewRequest("POST", "/users-api/users/1/addresses", bytes.NewBuffer(body))

	expected := createAddress()

	mockCreateAddressAction.On("Execute", int64(1), mock.Anything).Return(expected, nil)

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 201, w.Code)
	mockCreateAddressAction.AssertExpectations(suite.T())
}

func (suite *AddressHandlersTestSuite) TestAddressHandlers_CreateAddress_InvalidBodyReturns400() {
	w := httptest.NewRecorder()

	body, _ := json.Marshal(map[string]interface{}{
		"first_name": 33,
	})
	req, _ := http.NewRequest("POST", "/users-api/users/1/addresses", bytes.NewBuffer(body))

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 400, w.Code)
}

func (suite *AddressHandlersTestSuite) TestAddressHandlers_CreateAddress_Returns400OnActionFailure() {
	w := httptest.NewRecorder()

	body, _ := json.Marshal(createAddressBodyRequest())
	req, _ := http.NewRequest("POST", "/users-api/users/1/addresses", bytes.NewBuffer(body))

	mockCreateAddressAction.On("Execute", int64(1), mock.Anything).Return(entities.Address{}, errors.New("error"))

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 400, w.Code)
	mockCreateAddressAction.AssertExpectations(suite.T())
}

func (suite *AddressHandlersTestSuite) TestAddressHandlers_DeleteAddress_Returns400OnActionFailure() {
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("DELETE", "/users-api/users/1/addresses/10", nil)

	mockDeleteAddressAction.On("Execute", int64(10)).Return(errors.New("error"))

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 400, w.Code)
	mockDeleteAddressAction.AssertExpectations(suite.T())
}

func (suite *AddressHandlersTestSuite) TestAddressHandlers_DeleteAddress_Returns200OnSuccess() {
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("DELETE", "/users-api/users/1/addresses/10", nil)

	mockDeleteAddressAction.On("Execute", int64(10)).Return(nil)

	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), 200, w.Code)
	mockDeleteAddressAction.AssertExpectations(suite.T())
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
