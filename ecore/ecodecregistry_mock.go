// Code generated by mockery v2.16.0. DO NOT EDIT.

// *****************************************************************************
// Copyright(c) 2021 MASA Group
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// *****************************************************************************

package ecore

import mock "github.com/stretchr/testify/mock"

// MockECodecRegistry is an autogenerated mock type for the ECodecRegistry type
type MockECodecRegistry struct {
	mock.Mock
}

type MockECodecRegistry_Expecter struct {
	mock *mock.Mock
}

func (_m *MockECodecRegistry) EXPECT() *MockECodecRegistry_Expecter {
	return &MockECodecRegistry_Expecter{mock: &_m.Mock}
}

// GetCodec provides a mock function with given fields: uri
func (_m *MockECodecRegistry) GetCodec(uri *URI) ECodec {
	ret := _m.Called(uri)

	var r0 ECodec
	if rf, ok := ret.Get(0).(func(*URI) ECodec); ok {
		r0 = rf(uri)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ECodec)
		}
	}

	return r0
}

// MockECodecRegistry_GetCodec_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCodec'
type MockECodecRegistry_GetCodec_Call struct {
	*mock.Call
}

// GetCodec is a helper method to define mock.On call
//   - uri *URI
func (_e *MockECodecRegistry_Expecter) GetCodec(uri interface{}) *MockECodecRegistry_GetCodec_Call {
	return &MockECodecRegistry_GetCodec_Call{Call: _e.mock.On("GetCodec", uri)}
}

func (_c *MockECodecRegistry_GetCodec_Call) Run(run func(uri *URI)) *MockECodecRegistry_GetCodec_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*URI))
	})
	return _c
}

func (_c *MockECodecRegistry_GetCodec_Call) Return(_a0 ECodec) *MockECodecRegistry_GetCodec_Call {
	_c.Call.Return(_a0)
	return _c
}

// GetExtensionToCodecMap provides a mock function with given fields:
func (_m *MockECodecRegistry) GetExtensionToCodecMap() map[string]ECodec {
	ret := _m.Called()

	var r0 map[string]ECodec
	if rf, ok := ret.Get(0).(func() map[string]ECodec); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]ECodec)
		}
	}

	return r0
}

// MockECodecRegistry_GetExtensionToCodecMap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetExtensionToCodecMap'
type MockECodecRegistry_GetExtensionToCodecMap_Call struct {
	*mock.Call
}

// GetExtensionToCodecMap is a helper method to define mock.On call
func (_e *MockECodecRegistry_Expecter) GetExtensionToCodecMap() *MockECodecRegistry_GetExtensionToCodecMap_Call {
	return &MockECodecRegistry_GetExtensionToCodecMap_Call{Call: _e.mock.On("GetExtensionToCodecMap")}
}

func (_c *MockECodecRegistry_GetExtensionToCodecMap_Call) Run(run func()) *MockECodecRegistry_GetExtensionToCodecMap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockECodecRegistry_GetExtensionToCodecMap_Call) Return(_a0 map[string]ECodec) *MockECodecRegistry_GetExtensionToCodecMap_Call {
	_c.Call.Return(_a0)
	return _c
}

// GetProtocolToCodecMap provides a mock function with given fields:
func (_m *MockECodecRegistry) GetProtocolToCodecMap() map[string]ECodec {
	ret := _m.Called()

	var r0 map[string]ECodec
	if rf, ok := ret.Get(0).(func() map[string]ECodec); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]ECodec)
		}
	}

	return r0
}

// MockECodecRegistry_GetProtocolToCodecMap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetProtocolToCodecMap'
type MockECodecRegistry_GetProtocolToCodecMap_Call struct {
	*mock.Call
}

// GetProtocolToCodecMap is a helper method to define mock.On call
func (_e *MockECodecRegistry_Expecter) GetProtocolToCodecMap() *MockECodecRegistry_GetProtocolToCodecMap_Call {
	return &MockECodecRegistry_GetProtocolToCodecMap_Call{Call: _e.mock.On("GetProtocolToCodecMap")}
}

func (_c *MockECodecRegistry_GetProtocolToCodecMap_Call) Run(run func()) *MockECodecRegistry_GetProtocolToCodecMap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockECodecRegistry_GetProtocolToCodecMap_Call) Return(_a0 map[string]ECodec) *MockECodecRegistry_GetProtocolToCodecMap_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewMockECodecRegistry interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockECodecRegistry creates a new instance of MockECodecRegistry. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockECodecRegistry(t mockConstructorTestingTNewMockECodecRegistry) *MockECodecRegistry {
	mock := &MockECodecRegistry{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}