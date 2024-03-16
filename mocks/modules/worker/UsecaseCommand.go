// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	context "context"
	request "worker-service/internal/modules/worker/models/request"

	mock "github.com/stretchr/testify/mock"
)

// UsecaseCommand is an autogenerated mock type for the UsecaseCommand type
type UsecaseCommand struct {
	mock.Mock
}

// CreateBankTicket provides a mock function with given fields: origCtx, payload
func (_m *UsecaseCommand) CreateBankTicket(origCtx context.Context, payload request.CreateTicketReq) (*string, error) {
	ret := _m.Called(origCtx, payload)

	if len(ret) == 0 {
		panic("no return value specified for CreateBankTicket")
	}

	var r0 *string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, request.CreateTicketReq) (*string, error)); ok {
		return rf(origCtx, payload)
	}
	if rf, ok := ret.Get(0).(func(context.Context, request.CreateTicketReq) *string); ok {
		r0 = rf(origCtx, payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, request.CreateTicketReq) error); ok {
		r1 = rf(origCtx, payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateOnlineBankTicket provides a mock function with given fields: origCtx, payload
func (_m *UsecaseCommand) CreateOnlineBankTicket(origCtx context.Context, payload request.CreateOnlineTicketReq) (*string, error) {
	ret := _m.Called(origCtx, payload)

	if len(ret) == 0 {
		panic("no return value specified for CreateOnlineBankTicket")
	}

	var r0 *string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, request.CreateOnlineTicketReq) (*string, error)); ok {
		return rf(origCtx, payload)
	}
	if rf, ok := ret.Get(0).(func(context.Context, request.CreateOnlineTicketReq) *string); ok {
		r0 = rf(origCtx, payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, request.CreateOnlineTicketReq) error); ok {
		r1 = rf(origCtx, payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAllExpiryBankTicket provides a mock function with given fields: origCtx
func (_m *UsecaseCommand) UpdateAllExpiryBankTicket(origCtx context.Context) (*string, error) {
	ret := _m.Called(origCtx)

	if len(ret) == 0 {
		panic("no return value specified for UpdateAllExpiryBankTicket")
	}

	var r0 *string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*string, error)); ok {
		return rf(origCtx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *string); ok {
		r0 = rf(origCtx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(origCtx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateAllExpiryPayment provides a mock function with given fields: origCtx
func (_m *UsecaseCommand) UpdateAllExpiryPayment(origCtx context.Context) (*string, error) {
	ret := _m.Called(origCtx)

	if len(ret) == 0 {
		panic("no return value specified for UpdateAllExpiryPayment")
	}

	var r0 *string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*string, error)); ok {
		return rf(origCtx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *string); ok {
		r0 = rf(origCtx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*string)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(origCtx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUsecaseCommand creates a new instance of UsecaseCommand. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUsecaseCommand(t interface {
	mock.TestingT
	Cleanup(func())
}) *UsecaseCommand {
	mock := &UsecaseCommand{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
