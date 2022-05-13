// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	entities "domain-driven-design-layout/domain/entities"

	mock "github.com/stretchr/testify/mock"
)

// CreateUserMock is an autogenerated mock type for the CreateUserMock type
type CreateUserMock struct {
	mock.Mock
}

// Execute provides a mock function with given fields: ctx, proto
func (_m *CreateUserMock) Execute(ctx entities.ApplicationContext, proto entities.UserPrototype) (entities.User, error) {
	ret := _m.Called(ctx, proto)

	var r0 entities.User
	if rf, ok := ret.Get(0).(func(entities.ApplicationContext, entities.UserPrototype) entities.User); ok {
		r0 = rf(ctx, proto)
	} else {
		r0 = ret.Get(0).(entities.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(entities.ApplicationContext, entities.UserPrototype) error); ok {
		r1 = rf(ctx, proto)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
