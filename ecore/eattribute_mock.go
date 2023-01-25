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

type MockEAttribute struct {
	MockEStructuralFeature
}

type MockEAttribute_Expecter struct {
	MockEStructuralFeature_Expecter
}

func (eAttribute *MockEAttribute) EXPECT() *MockEAttribute_Expecter {
	e := &MockEAttribute_Expecter{}
	e.Mock = &eAttribute.Mock
	return e
}

// GetEAttributeType get the value of eAttributeType
func (eAttribute *MockEAttribute) GetEAttributeType() EDataType {
	ret := eAttribute.Called()

	var r EDataType
	if rf, ok := ret.Get(0).(func() EDataType); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EDataType)
		}
	}

	return r
}

type MockEAttribute_GetEAttributeType_Call struct {
	*mock.Call
}

func (e *MockEAttribute_Expecter) GetEAttributeType() *MockEAttribute_GetEAttributeType_Call {
	return &MockEAttribute_GetEAttributeType_Call{Call: e.Mock.On("GetEAttributeType")}
}

func (c *MockEAttribute_GetEAttributeType_Call) Run(run func()) *MockEAttribute_GetEAttributeType_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEAttribute_GetEAttributeType_Call) Return(eAttributeType EDataType) *MockEAttribute_GetEAttributeType_Call {
	c.Call.Return(eAttributeType)
	return c
}

// IsID get the value of isID
func (eAttribute *MockEAttribute) IsID() bool {
	ret := eAttribute.Called()

	var r bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(bool)
		}
	}

	return r
}

type MockEAttribute_IsID_Call struct {
	*mock.Call
}

func (e *MockEAttribute_Expecter) IsID() *MockEAttribute_IsID_Call {
	return &MockEAttribute_IsID_Call{Call: e.Mock.On("IsID")}
}

func (c *MockEAttribute_IsID_Call) Run(run func()) *MockEAttribute_IsID_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEAttribute_IsID_Call) Return(isID bool) *MockEAttribute_IsID_Call {
	c.Call.Return(isID)
	return c
}

// SetID provides mock implementation for setting the value of isID
func (eAttribute *MockEAttribute) SetID(isID bool) {
	eAttribute.Called(isID)
}

type MockEAttribute_SetID_Call struct {
	*mock.Call
}

// SetIDis a helper method to define mock.On call
// - isID bool
func (e *MockEAttribute_Expecter) SetID(isID any) *MockEAttribute_SetID_Call {
	return &MockEAttribute_SetID_Call{Call: e.Mock.On("SetID", isID)}
}

func (c *MockEAttribute_SetID_Call) Run(run func(isID bool)) *MockEAttribute_SetID_Call {
	c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool))
	})
	return c
}

func (c *MockEAttribute_SetID_Call) Return() *MockEAttribute_SetID_Call {
	c.Call.Return()
	return c
}

type mockConstructorTestingTNewMockEAttribute interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockEAttribute creates a new instance of MockEAttribute. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockEAttribute(t mockConstructorTestingTNewMockEAttribute) *MockEAttribute {
	mock := &MockEAttribute{}
	mock.Mock.Test(t)
	t.Cleanup(func() { mock.AssertExpectations(t) })
	return mock
}
