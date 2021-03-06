// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	context "context"
	entities "domain-driven-design-layout/domain/entities"

	mock "github.com/stretchr/testify/mock"
)

// TxRepositoryCreatorMock is an autogenerated mock type for the TxRepositoryCreatorMock type
type TxRepositoryCreatorMock struct {
	mock.Mock
}

// CreateTxMainDatabase provides a mock function with given fields: ctx
func (_m *TxRepositoryCreatorMock) CreateTxMainDatabase(ctx context.Context) (entities.MainDatabase, error) {
	ret := _m.Called(ctx)

	var r0 entities.MainDatabase
	if rf, ok := ret.Get(0).(func(context.Context) entities.MainDatabase); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(entities.MainDatabase)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
