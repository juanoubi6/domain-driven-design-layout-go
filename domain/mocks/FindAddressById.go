// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	entities "domain-driven-design-layout/domain/entities"

	mock "github.com/stretchr/testify/mock"
)

// FindAddressByIdMock is an autogenerated mock type for the FindAddressByIdMock type
type FindAddressByIdMock struct {
	mock.Mock
}

// Execute provides a mock function with given fields: ctx, addressID
func (_m *FindAddressByIdMock) Execute(ctx entities.ApplicationContext, addressID int64) (*entities.Address, error) {
	ret := _m.Called(ctx, addressID)

	var r0 *entities.Address
	if rf, ok := ret.Get(0).(func(entities.ApplicationContext, int64) *entities.Address); ok {
		r0 = rf(ctx, addressID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.Address)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entities.ApplicationContext, int64) error); ok {
		r1 = rf(ctx, addressID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
