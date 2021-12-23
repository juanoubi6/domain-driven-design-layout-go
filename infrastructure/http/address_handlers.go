package http

import (
	"domain-driven-design-layout/domain/actions/addresses"
	"domain-driven-design-layout/infrastructure/builder"
	"domain-driven-design-layout/infrastructure/http/requests"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AddressHandlers struct {
	createAddressAction addresses.CreateAddress
	deleteAddressAction addresses.DeleteAddress
}

func NewAddressHandlers(actions *builder.Actions) (*AddressHandlers, error) {
	return &AddressHandlers{
		createAddressAction: actions.CreateAddress,
		deleteAddressAction: actions.DeleteAddress,
	}, nil
}

func (r *AddressHandlers) CreateAddress(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid address id value in URL"})
		return
	}

	var request requests.CreateAddressRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdAddress, err := r.createAddressAction.Execute(int64(userID), request.ToAddressPrototype())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdAddress)
}

func (r *AddressHandlers) DeleteAddress(c *gin.Context) {
	addressId, err := strconv.Atoi(c.Param("addressID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid address id value in URL"})
		return
	}

	if err := r.deleteAddressAction.Execute(int64(addressId)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}