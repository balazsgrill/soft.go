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

// MockETypedElement is an mock type for the ETypedElement type
type MockETypedElement struct {
	mock.Mock
	MockETypedElement_Prototype
}

// MockETypedElement_Prototype is the mock implementation of all ETypedElement methods ( inherited and declared )
type MockETypedElement_Prototype struct {
	mock *mock.Mock
	MockENamedElement_Prototype
	MockETypedElement_Prototype_Methods
}

func (_mp *MockETypedElement_Prototype) SetMock(mock *mock.Mock) {
	_mp.mock = mock
	_mp.MockENamedElement_Prototype.SetMock(mock)
	_mp.MockETypedElement_Prototype_Methods.SetMock(mock)
}

// MockETypedElement_Expecter is the expecter implementation for all ETypedElement methods ( inherited and declared )
type MockETypedElement_Expecter struct {
	MockENamedElement_Expecter
	MockETypedElement_Expecter_Methods
}

func (_me *MockETypedElement_Expecter) SetMock(mock *mock.Mock) {
	_me.MockENamedElement_Expecter.SetMock(mock)
	_me.MockETypedElement_Expecter_Methods.SetMock(mock)
}

func (eTypedElement *MockETypedElement_Prototype) EXPECT() *MockETypedElement_Expecter {
	expecter := &MockETypedElement_Expecter{}
	expecter.SetMock(eTypedElement.mock)
	return expecter
}

// MockETypedElement_Prototype_Methods is the mock implementation of ETypedElement declared methods
type MockETypedElement_Prototype_Methods struct {
	mock *mock.Mock
}

func (_mdp *MockETypedElement_Prototype_Methods) SetMock(mock *mock.Mock) {
	_mdp.mock = mock
}

// MockETypedElement_Expecter_Methods is the expecter implementation of ETypedElement declared methods
type MockETypedElement_Expecter_Methods struct {
	mock *mock.Mock
}

func (_mde *MockETypedElement_Expecter_Methods) SetMock(mock *mock.Mock) {
	_mde.mock = mock
}

// GetEType get the value of eType
func (eTypedElement *MockETypedElement_Prototype_Methods) GetEType() EClassifier {
	ret := eTypedElement.mock.Called()

	var r EClassifier
	if rf, ok := ret.Get(0).(func() EClassifier); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EClassifier)
		}
	}

	return r
}

type MockETypedElement_GetEType_Call struct {
	*mock.Call
}

func (e *MockETypedElement_Expecter_Methods) GetEType() *MockETypedElement_GetEType_Call {
	return &MockETypedElement_GetEType_Call{Call: e.mock.On("GetEType")}
}

func (c *MockETypedElement_GetEType_Call) Run(run func()) *MockETypedElement_GetEType_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockETypedElement_GetEType_Call) Return(eType EClassifier) *MockETypedElement_GetEType_Call {
	c.Call.Return(eType)
	return c
}

// SetEType provides mock implementation for setting the value of eType
func (eTypedElement *MockETypedElement_Prototype_Methods) SetEType(eType EClassifier) {
	eTypedElement.mock.Called(eType)
}

type MockETypedElement_SetEType_Call struct {
	*mock.Call
}

// SetEType is a helper method to define mock.On call
// - eType EClassifier
func (e *MockETypedElement_Expecter_Methods) SetEType(eType any) *MockETypedElement_SetEType_Call {
	return &MockETypedElement_SetEType_Call{Call: e.mock.On("SetEType", eType)}
}

func (c *MockETypedElement_SetEType_Call) Run(run func(eType EClassifier)) *MockETypedElement_SetEType_Call {
	c.Call.Run(func(args mock.Arguments) {
		run(args[0].(EClassifier))
	})
	return c
}

func (c *MockETypedElement_SetEType_Call) Return() *MockETypedElement_SetEType_Call {
	c.Call.Return()
	return c
}

// UnsetEType provides mock implementation for unset the value of eType
func (eTypedElement *MockETypedElement_Prototype_Methods) UnsetEType() {
	eTypedElement.mock.Called()
}

type MockETypedElement_UnsetEType_Call struct {
	*mock.Call
}

func (e *MockETypedElement_Expecter_Methods) UnsetEType() *MockETypedElement_UnsetEType_Call {
	return &MockETypedElement_UnsetEType_Call{Call: e.mock.On("UnsetEType")}
}

func (c *MockETypedElement_UnsetEType_Call) Run(run func()) *MockETypedElement_UnsetEType_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockETypedElement_UnsetEType_Call) Return() *MockETypedElement_UnsetEType_Call {
	c.Call.Return()
	return c
}

// GetLowerBound get the value of lowerBound
func (eTypedElement *MockETypedElement_Prototype_Methods) GetLowerBound() int {
	ret := eTypedElement.mock.Called()

	var r int
	if rf, ok := ret.Get(0).(func() int); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(int)
		}
	}

	return r
}

type MockETypedElement_GetLowerBound_Call struct {
	*mock.Call
}

func (e *MockETypedElement_Expecter_Methods) GetLowerBound() *MockETypedElement_GetLowerBound_Call {
	return &MockETypedElement_GetLowerBound_Call{Call: e.mock.On("GetLowerBound")}
}

func (c *MockETypedElement_GetLowerBound_Call) Run(run func()) *MockETypedElement_GetLowerBound_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockETypedElement_GetLowerBound_Call) Return(lowerBound int) *MockETypedElement_GetLowerBound_Call {
	c.Call.Return(lowerBound)
	return c
}

// SetLowerBound provides mock implementation for setting the value of lowerBound
func (eTypedElement *MockETypedElement_Prototype_Methods) SetLowerBound(lowerBound int) {
	eTypedElement.mock.Called(lowerBound)
}

type MockETypedElement_SetLowerBound_Call struct {
	*mock.Call
}

// SetLowerBound is a helper method to define mock.On call
// - lowerBound int
func (e *MockETypedElement_Expecter_Methods) SetLowerBound(lowerBound any) *MockETypedElement_SetLowerBound_Call {
	return &MockETypedElement_SetLowerBound_Call{Call: e.mock.On("SetLowerBound", lowerBound)}
}

func (c *MockETypedElement_SetLowerBound_Call) Run(run func(lowerBound int)) *MockETypedElement_SetLowerBound_Call {
	c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return c
}

func (c *MockETypedElement_SetLowerBound_Call) Return() *MockETypedElement_SetLowerBound_Call {
	c.Call.Return()
	return c
}

// IsMany get the value of isMany
func (eTypedElement *MockETypedElement_Prototype_Methods) IsMany() bool {
	ret := eTypedElement.mock.Called()

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

type MockETypedElement_IsMany_Call struct {
	*mock.Call
}

func (e *MockETypedElement_Expecter_Methods) IsMany() *MockETypedElement_IsMany_Call {
	return &MockETypedElement_IsMany_Call{Call: e.mock.On("IsMany")}
}

func (c *MockETypedElement_IsMany_Call) Run(run func()) *MockETypedElement_IsMany_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockETypedElement_IsMany_Call) Return(isMany bool) *MockETypedElement_IsMany_Call {
	c.Call.Return(isMany)
	return c
}

// IsOrdered get the value of isOrdered
func (eTypedElement *MockETypedElement_Prototype_Methods) IsOrdered() bool {
	ret := eTypedElement.mock.Called()

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

type MockETypedElement_IsOrdered_Call struct {
	*mock.Call
}

func (e *MockETypedElement_Expecter_Methods) IsOrdered() *MockETypedElement_IsOrdered_Call {
	return &MockETypedElement_IsOrdered_Call{Call: e.mock.On("IsOrdered")}
}

func (c *MockETypedElement_IsOrdered_Call) Run(run func()) *MockETypedElement_IsOrdered_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockETypedElement_IsOrdered_Call) Return(isOrdered bool) *MockETypedElement_IsOrdered_Call {
	c.Call.Return(isOrdered)
	return c
}

// SetOrdered provides mock implementation for setting the value of isOrdered
func (eTypedElement *MockETypedElement_Prototype_Methods) SetOrdered(isOrdered bool) {
	eTypedElement.mock.Called(isOrdered)
}

type MockETypedElement_SetOrdered_Call struct {
	*mock.Call
}

// SetOrdered is a helper method to define mock.On call
// - isOrdered bool
func (e *MockETypedElement_Expecter_Methods) SetOrdered(isOrdered any) *MockETypedElement_SetOrdered_Call {
	return &MockETypedElement_SetOrdered_Call{Call: e.mock.On("SetOrdered", isOrdered)}
}

func (c *MockETypedElement_SetOrdered_Call) Run(run func(isOrdered bool)) *MockETypedElement_SetOrdered_Call {
	c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool))
	})
	return c
}

func (c *MockETypedElement_SetOrdered_Call) Return() *MockETypedElement_SetOrdered_Call {
	c.Call.Return()
	return c
}

// IsRequired get the value of isRequired
func (eTypedElement *MockETypedElement_Prototype_Methods) IsRequired() bool {
	ret := eTypedElement.mock.Called()

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

type MockETypedElement_IsRequired_Call struct {
	*mock.Call
}

func (e *MockETypedElement_Expecter_Methods) IsRequired() *MockETypedElement_IsRequired_Call {
	return &MockETypedElement_IsRequired_Call{Call: e.mock.On("IsRequired")}
}

func (c *MockETypedElement_IsRequired_Call) Run(run func()) *MockETypedElement_IsRequired_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockETypedElement_IsRequired_Call) Return(isRequired bool) *MockETypedElement_IsRequired_Call {
	c.Call.Return(isRequired)
	return c
}

// IsUnique get the value of isUnique
func (eTypedElement *MockETypedElement_Prototype_Methods) IsUnique() bool {
	ret := eTypedElement.mock.Called()

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

type MockETypedElement_IsUnique_Call struct {
	*mock.Call
}

func (e *MockETypedElement_Expecter_Methods) IsUnique() *MockETypedElement_IsUnique_Call {
	return &MockETypedElement_IsUnique_Call{Call: e.mock.On("IsUnique")}
}

func (c *MockETypedElement_IsUnique_Call) Run(run func()) *MockETypedElement_IsUnique_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockETypedElement_IsUnique_Call) Return(isUnique bool) *MockETypedElement_IsUnique_Call {
	c.Call.Return(isUnique)
	return c
}

// SetUnique provides mock implementation for setting the value of isUnique
func (eTypedElement *MockETypedElement_Prototype_Methods) SetUnique(isUnique bool) {
	eTypedElement.mock.Called(isUnique)
}

type MockETypedElement_SetUnique_Call struct {
	*mock.Call
}

// SetUnique is a helper method to define mock.On call
// - isUnique bool
func (e *MockETypedElement_Expecter_Methods) SetUnique(isUnique any) *MockETypedElement_SetUnique_Call {
	return &MockETypedElement_SetUnique_Call{Call: e.mock.On("SetUnique", isUnique)}
}

func (c *MockETypedElement_SetUnique_Call) Run(run func(isUnique bool)) *MockETypedElement_SetUnique_Call {
	c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool))
	})
	return c
}

func (c *MockETypedElement_SetUnique_Call) Return() *MockETypedElement_SetUnique_Call {
	c.Call.Return()
	return c
}

// GetUpperBound get the value of upperBound
func (eTypedElement *MockETypedElement_Prototype_Methods) GetUpperBound() int {
	ret := eTypedElement.mock.Called()

	var r int
	if rf, ok := ret.Get(0).(func() int); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(int)
		}
	}

	return r
}

type MockETypedElement_GetUpperBound_Call struct {
	*mock.Call
}

func (e *MockETypedElement_Expecter_Methods) GetUpperBound() *MockETypedElement_GetUpperBound_Call {
	return &MockETypedElement_GetUpperBound_Call{Call: e.mock.On("GetUpperBound")}
}

func (c *MockETypedElement_GetUpperBound_Call) Run(run func()) *MockETypedElement_GetUpperBound_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockETypedElement_GetUpperBound_Call) Return(upperBound int) *MockETypedElement_GetUpperBound_Call {
	c.Call.Return(upperBound)
	return c
}

// SetUpperBound provides mock implementation for setting the value of upperBound
func (eTypedElement *MockETypedElement_Prototype_Methods) SetUpperBound(upperBound int) {
	eTypedElement.mock.Called(upperBound)
}

type MockETypedElement_SetUpperBound_Call struct {
	*mock.Call
}

// SetUpperBound is a helper method to define mock.On call
// - upperBound int
func (e *MockETypedElement_Expecter_Methods) SetUpperBound(upperBound any) *MockETypedElement_SetUpperBound_Call {
	return &MockETypedElement_SetUpperBound_Call{Call: e.mock.On("SetUpperBound", upperBound)}
}

func (c *MockETypedElement_SetUpperBound_Call) Run(run func(upperBound int)) *MockETypedElement_SetUpperBound_Call {
	c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int))
	})
	return c
}

func (c *MockETypedElement_SetUpperBound_Call) Return() *MockETypedElement_SetUpperBound_Call {
	c.Call.Return()
	return c
}

type mockConstructorTestingTNewMockETypedElement interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockETypedElement creates a new instance of MockETypedElement. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockETypedElement(t mockConstructorTestingTNewMockETypedElement) *MockETypedElement {
	mock := &MockETypedElement{}
	mock.SetMock(&mock.Mock)
	mock.Mock.Test(t)
	t.Cleanup(func() { mock.AssertExpectations(t) })
	return mock
}
