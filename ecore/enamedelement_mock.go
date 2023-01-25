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

type MockENamedElement struct {
	MockEModelElement
}

type MockENamedElement_Expecter struct {
	MockEModelElement_Expecter
}

func (eNamedElement *MockENamedElement) EXPECT() *MockENamedElement_Expecter {
	e := &MockENamedElement_Expecter{}
	e.Mock = &eNamedElement.Mock
	return e
}

// GetName get the value of name
func (eNamedElement *MockENamedElement) GetName() string {
	ret := eNamedElement.Called()

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

type MockENamedElement_GetName_Call struct {
	*mock.Call
}

func (e *MockENamedElement_Expecter) GetName() *MockENamedElement_GetName_Call {
	return &MockENamedElement_GetName_Call{Call: e.Mock.On("GetName")}
}

func (c *MockENamedElement_GetName_Call) Run(run func()) *MockENamedElement_GetName_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockENamedElement_GetName_Call) Return(name string) *MockENamedElement_GetName_Call {
	c.Call.Return(name)
	return c
}

// SetName provides mock implementation for setting the value of name
func (eNamedElement *MockENamedElement) SetName(name string) {
	eNamedElement.Called(name)
}

type MockENamedElement_SetName_Call struct {
	*mock.Call
}

// SetNameis a helper method to define mock.On call
// - name string
func (e *MockENamedElement_Expecter) SetName(name any) *MockENamedElement_SetName_Call {
	return &MockENamedElement_SetName_Call{Call: e.Mock.On("SetName", name)}
}

func (c *MockENamedElement_SetName_Call) Run(run func(name string)) *MockENamedElement_SetName_Call {
	c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return c
}

func (c *MockENamedElement_SetName_Call) Return() *MockENamedElement_SetName_Call {
	c.Call.Return()
	return c
}

type mockConstructorTestingTNewMockENamedElement interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockENamedElement creates a new instance of MockENamedElement. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockENamedElement(t mockConstructorTestingTNewMockENamedElement) *MockENamedElement {
	mock := &MockENamedElement{}
	mock.Mock.Test(t)
	t.Cleanup(func() { mock.AssertExpectations(t) })
	return mock
}
