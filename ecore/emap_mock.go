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

// MockEMap is an autogenerated mock type for the EMap type
type MockEMap struct {
	MockEList
}

type MockEMap_Expecter struct {
	MockEList_Expecter
}

func (_m *MockEMap) EXPECT() *MockEMap_Expecter {
	e := &MockEMap_Expecter{}
	e.Mock = &_m.Mock
	return e
}

// ContainsKey provides a mock function with given fields: key
func (_m *MockEMap) ContainsKey(key any) bool {
	ret := _m.Called(key)

	var r0 bool
	if rf, ok := ret.Get(0).(func(any) bool); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockEMap_ContainsKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ContainsKey'
type MockEMap_ContainsKey_Call struct {
	*mock.Call
}

// ContainsKey is a helper method to define mock.On call
//   - key any
func (_e *MockEMap_Expecter) ContainsKey(key any) *MockEMap_ContainsKey_Call {
	return &MockEMap_ContainsKey_Call{Call: _e.Mock.On("ContainsKey", key)}
}

func (_c *MockEMap_ContainsKey_Call) Run(run func(key any)) *MockEMap_ContainsKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0])
	})
	return _c
}

func (_c *MockEMap_ContainsKey_Call) Return(_a0 bool) *MockEMap_ContainsKey_Call {
	_c.Call.Return(_a0)
	return _c
}

// ContainsValue provides a mock function with given fields: value
func (_m *MockEMap) ContainsValue(value any) bool {
	ret := _m.Called(value)

	var r0 bool
	if rf, ok := ret.Get(0).(func(any) bool); ok {
		r0 = rf(value)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockEMap_ContainsValue_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ContainsValue'
type MockEMap_ContainsValue_Call struct {
	*mock.Call
}

// ContainsValue is a helper method to define mock.On call
//   - value any
func (_e *MockEMap_Expecter) ContainsValue(value any) *MockEMap_ContainsValue_Call {
	return &MockEMap_ContainsValue_Call{Call: _e.Mock.On("ContainsValue", value)}
}

func (_c *MockEMap_ContainsValue_Call) Run(run func(value any)) *MockEMap_ContainsValue_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0])
	})
	return _c
}

func (_c *MockEMap_ContainsValue_Call) Return(_a0 bool) *MockEMap_ContainsValue_Call {
	_c.Call.Return(_a0)
	return _c
}

// GetValue provides a mock function with given fields: key
func (_m *MockEMap) GetValue(key any) any {
	ret := _m.Called(key)

	var r0 any
	if rf, ok := ret.Get(0).(func(any) any); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0)
		}
	}

	return r0
}

// MockEMap_GetValue_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetValue'
type MockEMap_GetValue_Call struct {
	*mock.Call
}

// GetValue is a helper method to define mock.On call
//   - key any
func (_e *MockEMap_Expecter) GetValue(key any) *MockEMap_GetValue_Call {
	return &MockEMap_GetValue_Call{Call: _e.Mock.On("GetValue", key)}
}

func (_c *MockEMap_GetValue_Call) Run(run func(key any)) *MockEMap_GetValue_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0])
	})
	return _c
}

func (_c *MockEMap_GetValue_Call) Return(_a0 any) *MockEMap_GetValue_Call {
	_c.Call.Return(_a0)
	return _c
}

// Put provides a mock function with given fields: key, value
func (_m *MockEMap) Put(key any, value any) {
	_m.Called(key, value)
}

// MockEMap_Put_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Put'
type MockEMap_Put_Call struct {
	*mock.Call
}

// Put is a helper method to define mock.On call
//   - key any
//   - value any
func (_e *MockEMap_Expecter) Put(key any, value any) *MockEMap_Put_Call {
	return &MockEMap_Put_Call{Call: _e.Mock.On("Put", key, value)}
}

func (_c *MockEMap_Put_Call) Run(run func(key any, value any)) *MockEMap_Put_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0], args[1])
	})
	return _c
}

func (_c *MockEMap_Put_Call) Return() *MockEMap_Put_Call {
	_c.Call.Return()
	return _c
}

// RemoveKey provides a mock function with given fields: key
func (_m *MockEMap) RemoveKey(key any) any {
	ret := _m.Called(key)

	var r0 any
	if rf, ok := ret.Get(0).(func(any) any); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0)
		}
	}

	return r0
}

// MockEMap_RemoveKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveKey'
type MockEMap_RemoveKey_Call struct {
	*mock.Call
}

// RemoveKey is a helper method to define mock.On call
//   - key any
func (_e *MockEMap_Expecter) RemoveKey(key any) *MockEMap_RemoveKey_Call {
	return &MockEMap_RemoveKey_Call{Call: _e.Mock.On("RemoveKey", key)}
}

func (_c *MockEMap_RemoveKey_Call) Run(run func(key any)) *MockEMap_RemoveKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0])
	})
	return _c
}

func (_c *MockEMap_RemoveKey_Call) Return(_a0 any) *MockEMap_RemoveKey_Call {
	_c.Call.Return(_a0)
	return _c
}

// ToMap provides a mock function with given fields:
func (_m *MockEMap) ToMap() map[any]any {
	ret := _m.Called()

	var r0 map[any]any
	if rf, ok := ret.Get(0).(func() map[any]any); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[any]any)
		}
	}

	return r0
}

// MockEMap_ToMap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ToMap'
type MockEMap_ToMap_Call struct {
	*mock.Call
}

// ToMap is a helper method to define mock.On call
func (_e *MockEMap_Expecter) ToMap() *MockEMap_ToMap_Call {
	return &MockEMap_ToMap_Call{Call: _e.Mock.On("ToMap")}
}

func (_c *MockEMap_ToMap_Call) Run(run func()) *MockEMap_ToMap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockEMap_ToMap_Call) Return(_a0 map[any]any) *MockEMap_ToMap_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewMockEMap interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockEMap creates a new instance of MockEMap. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockEMap(t mockConstructorTestingTNewMockEMap) *MockEMap {
	mock := &MockEMap{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
