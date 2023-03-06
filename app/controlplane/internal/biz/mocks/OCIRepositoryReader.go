//
// Copyright 2023 The Chainloop Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	biz "github.com/chainloop-dev/chainloop/app/controlplane/internal/biz"
	backend "github.com/chainloop-dev/chainloop/internal/blobmanager"

	context "context"

	mock "github.com/stretchr/testify/mock"
)

// OCIRepositoryReader is an autogenerated mock type for the OCIRepositoryReader type
type OCIRepositoryReader struct {
	mock.Mock
}

// FindByID provides a mock function with given fields: ctx, ID
func (_m *OCIRepositoryReader) FindByID(ctx context.Context, ID string) (*biz.OCIRepository, error) {
	ret := _m.Called(ctx, ID)

	var r0 *biz.OCIRepository
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*biz.OCIRepository, error)); ok {
		return rf(ctx, ID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *biz.OCIRepository); ok {
		r0 = rf(ctx, ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*biz.OCIRepository)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindMainRepo provides a mock function with given fields: ctx, orgID
func (_m *OCIRepositoryReader) FindMainRepo(ctx context.Context, orgID string) (*biz.OCIRepository, error) {
	ret := _m.Called(ctx, orgID)

	var r0 *biz.OCIRepository
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*biz.OCIRepository, error)); ok {
		return rf(ctx, orgID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *biz.OCIRepository); ok {
		r0 = rf(ctx, orgID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*biz.OCIRepository)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, orgID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PerformValidation provides a mock function with given fields: ctx, ID
func (_m *OCIRepositoryReader) PerformValidation(ctx context.Context, ID string) error {
	ret := _m.Called(ctx, ID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, ID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// loadUploaderBackend provides a mock function with given fields: ctx, secretName
func (_m *OCIRepositoryReader) loadUploaderBackend(ctx context.Context, secretName string) (backend.Uploader, error) {
	ret := _m.Called(ctx, secretName)

	var r0 backend.Uploader
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (backend.Uploader, error)); ok {
		return rf(ctx, secretName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) backend.Uploader); ok {
		r0 = rf(ctx, secretName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(backend.Uploader)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, secretName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewOCIRepositoryReader interface {
	mock.TestingT
	Cleanup(func())
}

// NewOCIRepositoryReader creates a new instance of OCIRepositoryReader. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewOCIRepositoryReader(t mockConstructorTestingTNewOCIRepositoryReader) *OCIRepositoryReader {
	mock := &OCIRepositoryReader{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
