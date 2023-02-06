// *****************************************************************************
// Copyright(c) 2021 MASA Group
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// *****************************************************************************

package ecore

import (
	io "io"

	mock "github.com/stretchr/testify/mock"
)

// MockEResourceCodec is mock type for the EResourceCodec type
type MockEResourceCodec struct {
	mock.Mock
}

type MockEResourceCodec_Expecter struct {
	mock *mock.Mock
}

func (_m *MockEResourceCodec) EXPECT() *MockEResourceCodec_Expecter {
	return &MockEResourceCodec_Expecter{mock: &_m.Mock}
}

// NewDecoder provides a mock function with given fields: resource, r, options
func (_m *MockEResourceCodec) NewDecoder(resource EResource, r io.Reader, options map[string]interface{}) EResourceDecoder {
	ret := _m.Called(resource, r, options)

	var r0 EResourceDecoder
	if rf, ok := ret.Get(0).(func(EResource, io.Reader, map[string]interface{}) EResourceDecoder); ok {
		r0 = rf(resource, r, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(EResourceDecoder)
		}
	}

	return r0
}

// MockEResourceCodec_NewDecoder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NewDecoder'
type MockEResourceCodec_NewDecoder_Call struct {
	*mock.Call
}

// NewDecoder is a helper method to define mock.On call
//   - resource EResource
//   - r io.Reader
//   - options map[string]interface{}
func (_e *MockEResourceCodec_Expecter) NewDecoder(resource interface{}, r interface{}, options interface{}) *MockEResourceCodec_NewDecoder_Call {
	return &MockEResourceCodec_NewDecoder_Call{Call: _e.mock.On("NewDecoder", resource, r, options)}
}

func (_c *MockEResourceCodec_NewDecoder_Call) Run(run func(resource EResource, r io.Reader, options map[string]interface{})) *MockEResourceCodec_NewDecoder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(EResource), args[1].(io.Reader), args[2].(map[string]interface{}))
	})
	return _c
}

func (_c *MockEResourceCodec_NewDecoder_Call) Return(_a0 EResourceDecoder) *MockEResourceCodec_NewDecoder_Call {
	_c.Call.Return(_a0)
	return _c
}

// NewEncoder provides a mock function with given fields: resource, w, options
func (_m *MockEResourceCodec) NewEncoder(resource EResource, w io.Writer, options map[string]interface{}) EResourceEncoder {
	ret := _m.Called(resource, w, options)

	var r0 EResourceEncoder
	if rf, ok := ret.Get(0).(func(EResource, io.Writer, map[string]interface{}) EResourceEncoder); ok {
		r0 = rf(resource, w, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(EResourceEncoder)
		}
	}

	return r0
}

// MockEResourceCodec_NewEncoder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'NewEncoder'
type MockEResourceCodec_NewEncoder_Call struct {
	*mock.Call
}

// NewEncoder is a helper method to define mock.On call
//   - resource EResource
//   - w io.Writer
//   - options map[string]interface{}
func (_e *MockEResourceCodec_Expecter) NewEncoder(resource interface{}, w interface{}, options interface{}) *MockEResourceCodec_NewEncoder_Call {
	return &MockEResourceCodec_NewEncoder_Call{Call: _e.mock.On("NewEncoder", resource, w, options)}
}

func (_c *MockEResourceCodec_NewEncoder_Call) Run(run func(resource EResource, w io.Writer, options map[string]interface{})) *MockEResourceCodec_NewEncoder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(EResource), args[1].(io.Writer), args[2].(map[string]interface{}))
	})
	return _c
}

func (_c *MockEResourceCodec_NewEncoder_Call) Return(_a0 EResourceEncoder) *MockEResourceCodec_NewEncoder_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewMockEResourceCodec interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockEResourceCodec creates a new instance of MockEResourceCodec. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockEResourceCodec(t mockConstructorTestingTNewMockEResourceCodec) *MockEResourceCodec {
	mock := &MockEResourceCodec{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
