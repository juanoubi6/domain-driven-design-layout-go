// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	entities "domain-driven-design-layout/domain/entities"

	mock "github.com/stretchr/testify/mock"
)

// FindUserByIdMock is an autogenerated mock type for the FindUserByIdMock type
type FindUserByIdMock struct {
	mock.Mock
}

// Execute provides a mock function with given fields: ctx, userID
func (_m *FindUserByIdMock) Execute(ctx entities.ApplicationContext, userID int64) (*entities.User, error) {
	ret := _m.Called(ctx, userID)

	var r0 *entities.User
	if rf, ok := ret.Get(0).(func(entities.ApplicationContext, int64) *entities.User); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entities.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entities.ApplicationContext, int64) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
