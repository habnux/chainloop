// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	credentials "github.com/chainloop-dev/chainloop/internal/credentials"
	mock "github.com/stretchr/testify/mock"
)

// Writer is an autogenerated mock type for the Writer type
type Writer struct {
	mock.Mock
}

// DeleteCreds provides a mock function with given fields: ctx, credID
func (_m *Writer) DeleteCreds(ctx context.Context, credID string) error {
	ret := _m.Called(ctx, credID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, credID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveAPICreds provides a mock function with given fields: ctx, org, creds
func (_m *Writer) SaveAPICreds(ctx context.Context, org string, creds *credentials.APICreds) (string, error) {
	ret := _m.Called(ctx, org, creds)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *credentials.APICreds) (string, error)); ok {
		return rf(ctx, org, creds)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *credentials.APICreds) string); ok {
		r0 = rf(ctx, org, creds)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *credentials.APICreds) error); ok {
		r1 = rf(ctx, org, creds)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveOCICreds provides a mock function with given fields: ctx, org, creds
func (_m *Writer) SaveOCICreds(ctx context.Context, org string, creds *credentials.OCIKeypair) (string, error) {
	ret := _m.Called(ctx, org, creds)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, *credentials.OCIKeypair) (string, error)); ok {
		return rf(ctx, org, creds)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, *credentials.OCIKeypair) string); ok {
		r0 = rf(ctx, org, creds)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, *credentials.OCIKeypair) error); ok {
		r1 = rf(ctx, org, creds)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewWriter interface {
	mock.TestingT
	Cleanup(func())
}

// NewWriter creates a new instance of Writer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewWriter(t mockConstructorTestingTNewWriter) *Writer {
	mock := &Writer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
