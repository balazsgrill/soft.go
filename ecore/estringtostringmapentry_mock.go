// Code generated by soft.generator.go. DO NOT EDIT.

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
	"github.com/stretchr/testify/mock"
)

// MockEStringToStringMapEntry is an mock type for the EStringToStringMapEntry type
type MockEStringToStringMapEntry struct {
	mock.Mock
	MockEStringToStringMapEntry_Prototype
}

// MockEStringToStringMapEntry_Prototype is the mock implementation of all EStringToStringMapEntry methods ( inherited and declared )
type MockEStringToStringMapEntry_Prototype struct {
	mock *mock.Mock
	MockEObjectInternal_Prototype
	MockEMapEntry_Prototype
	MockEStringToStringMapEntry_Prototype_Methods
}

func (_mp *MockEStringToStringMapEntry_Prototype) SetMock(mock *mock.Mock) {
	_mp.mock = mock
	_mp.MockEObjectInternal_Prototype.SetMock(mock)
	_mp.MockEMapEntry_Prototype.SetMock(mock)
	_mp.MockEStringToStringMapEntry_Prototype_Methods.SetMock(mock)
}

// MockEStringToStringMapEntry_Expecter is the expecter implementation for all EStringToStringMapEntry methods ( inherited and declared )
type MockEStringToStringMapEntry_Expecter struct {
	MockEObjectInternal_Expecter
	MockEMapEntry_Expecter
	MockEStringToStringMapEntry_Expecter_Methods
}

func (_me *MockEStringToStringMapEntry_Expecter) SetMock(mock *mock.Mock) {
	_me.MockEObjectInternal_Expecter.SetMock(mock)
	_me.MockEMapEntry_Expecter.SetMock(mock)
	_me.MockEStringToStringMapEntry_Expecter_Methods.SetMock(mock)
}

func (eStringToStringMapEntry *MockEStringToStringMapEntry_Prototype) EXPECT() *MockEStringToStringMapEntry_Expecter {
	expecter := &MockEStringToStringMapEntry_Expecter{}
	expecter.SetMock(eStringToStringMapEntry.mock)
	return expecter
}

// MockEStringToStringMapEntry_Prototype_Methods is the mock implementation of EStringToStringMapEntry declared methods
type MockEStringToStringMapEntry_Prototype_Methods struct {
	mock *mock.Mock
}

func (_mdp *MockEStringToStringMapEntry_Prototype_Methods) SetMock(mock *mock.Mock) {
	_mdp.mock = mock
}

// MockEStringToStringMapEntry_Expecter_Methods is the expecter implementation of EStringToStringMapEntry declared methods
type MockEStringToStringMapEntry_Expecter_Methods struct {
	mock *mock.Mock
}

func (_mde *MockEStringToStringMapEntry_Expecter_Methods) SetMock(mock *mock.Mock) {
	_mde.mock = mock
}

// GetTypedKey get the value of key
func (eStringToStringMapEntry *MockEStringToStringMapEntry_Prototype_Methods) GetTypedKey() string {
	ret := eStringToStringMapEntry.mock.Called()

	var r string
	if rf, ok := ret.Get(0).(func() string); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(string)
		}
	}

	return r
}

type MockEStringToStringMapEntry_GetTypedKey_Call struct {
	*mock.Call
}

func (e *MockEStringToStringMapEntry_Expecter_Methods) GetTypedKey() *MockEStringToStringMapEntry_GetTypedKey_Call {
	return &MockEStringToStringMapEntry_GetTypedKey_Call{Call: e.mock.On("GetTypedKey")}
}

func (c *MockEStringToStringMapEntry_GetTypedKey_Call) Run(run func()) *MockEStringToStringMapEntry_GetTypedKey_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEStringToStringMapEntry_GetTypedKey_Call) Return(key string) *MockEStringToStringMapEntry_GetTypedKey_Call {
	c.Call.Return(key)
	return c
}

// SetTypedKey provides mock implementation for setting the value of key
func (eStringToStringMapEntry *MockEStringToStringMapEntry_Prototype_Methods) SetTypedKey(key string) {
	eStringToStringMapEntry.mock.Called(key)
}

type MockEStringToStringMapEntry_SetTypedKey_Call struct {
	*mock.Call
}

// SetTypedKey is a helper method to define mock.On call
// - key string
func (e *MockEStringToStringMapEntry_Expecter_Methods) SetTypedKey(key any) *MockEStringToStringMapEntry_SetTypedKey_Call {
	return &MockEStringToStringMapEntry_SetTypedKey_Call{Call: e.mock.On("SetTypedKey", key)}
}

func (c *MockEStringToStringMapEntry_SetTypedKey_Call) Run(run func(key string)) *MockEStringToStringMapEntry_SetTypedKey_Call {
	c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return c
}

func (c *MockEStringToStringMapEntry_SetTypedKey_Call) Return() *MockEStringToStringMapEntry_SetTypedKey_Call {
	c.Call.Return()
	return c
}

// GetTypedValue get the value of value
func (eStringToStringMapEntry *MockEStringToStringMapEntry_Prototype_Methods) GetTypedValue() string {
	ret := eStringToStringMapEntry.mock.Called()

	var r string
	if rf, ok := ret.Get(0).(func() string); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(string)
		}
	}

	return r
}

type MockEStringToStringMapEntry_GetTypedValue_Call struct {
	*mock.Call
}

func (e *MockEStringToStringMapEntry_Expecter_Methods) GetTypedValue() *MockEStringToStringMapEntry_GetTypedValue_Call {
	return &MockEStringToStringMapEntry_GetTypedValue_Call{Call: e.mock.On("GetTypedValue")}
}

func (c *MockEStringToStringMapEntry_GetTypedValue_Call) Run(run func()) *MockEStringToStringMapEntry_GetTypedValue_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEStringToStringMapEntry_GetTypedValue_Call) Return(value string) *MockEStringToStringMapEntry_GetTypedValue_Call {
	c.Call.Return(value)
	return c
}

// SetTypedValue provides mock implementation for setting the value of value
func (eStringToStringMapEntry *MockEStringToStringMapEntry_Prototype_Methods) SetTypedValue(value string) {
	eStringToStringMapEntry.mock.Called(value)
}

type MockEStringToStringMapEntry_SetTypedValue_Call struct {
	*mock.Call
}

// SetTypedValue is a helper method to define mock.On call
// - value string
func (e *MockEStringToStringMapEntry_Expecter_Methods) SetTypedValue(value any) *MockEStringToStringMapEntry_SetTypedValue_Call {
	return &MockEStringToStringMapEntry_SetTypedValue_Call{Call: e.mock.On("SetTypedValue", value)}
}

func (c *MockEStringToStringMapEntry_SetTypedValue_Call) Run(run func(value string)) *MockEStringToStringMapEntry_SetTypedValue_Call {
	c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return c
}

func (c *MockEStringToStringMapEntry_SetTypedValue_Call) Return() *MockEStringToStringMapEntry_SetTypedValue_Call {
	c.Call.Return()
	return c
}

type mockConstructorTestingTNewMockEStringToStringMapEntry interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockEStringToStringMapEntry creates a new instance of MockEStringToStringMapEntry. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockEStringToStringMapEntry(t mockConstructorTestingTNewMockEStringToStringMapEntry) *MockEStringToStringMapEntry {
	mock := &MockEStringToStringMapEntry{}
	mock.SetMock(&mock.Mock)
	mock.Mock.Test(t)
	t.Cleanup(func() { mock.AssertExpectations(t) })
	return mock
}
