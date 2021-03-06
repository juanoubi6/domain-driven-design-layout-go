// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	entities "domain-driven-design-layout/domain/entities"

	mock "github.com/stretchr/testify/mock"
)

// CreateAddressMock is an autogenerated mock type for the CreateAddressMock type
type CreateAddressMock struct {
	mock.Mock
}

// Execute provides a mock function with given fields: ctx, userID, prototype
func (_m *CreateAddressMock) Execute(ctx entities.ApplicationContext, userID int64, prototype entities.AddressPrototype) (entities.Address, error) {
	ret := _m.Called(ctx, userID, prototype)

	var r0 entities.Address
	if rf, ok := ret.Get(0).(func(entities.ApplicationContext, int64, entities.AddressPrototype) entities.Address); ok {
		r0 = rf(ctx, userID, prototype)
	} else {
		r0 = ret.Get(0).(entities.Address)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entities.ApplicationContext, int64, entities.AddressPrototype) error); ok {
		r1 = rf(ctx, userID, prototype)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
