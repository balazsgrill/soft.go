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

// MockEAnnotation is an mock type for the EAnnotation type
type MockEAnnotation struct {
	mock.Mock
	MockEAnnotation_Prototype
}

// MockEAnnotation_Prototype is the mock implementation of all EAnnotation methods ( inherited and declared )
type MockEAnnotation_Prototype struct {
	mock *mock.Mock
	MockEModelElement_Prototype
	MockEAnnotation_Prototype_Methods
}

func (_mp *MockEAnnotation_Prototype) SetMock(mock *mock.Mock) {
	_mp.mock = mock
	_mp.MockEModelElement_Prototype.SetMock(mock)
	_mp.MockEAnnotation_Prototype_Methods.SetMock(mock)
}

// MockEAnnotation_Expecter is the expecter implementation for all EAnnotation methods ( inherited and declared )
type MockEAnnotation_Expecter struct {
	MockEModelElement_Expecter
	MockEAnnotation_Expecter_Methods
}

func (_me *MockEAnnotation_Expecter) SetMock(mock *mock.Mock) {
	_me.MockEModelElement_Expecter.SetMock(mock)
	_me.MockEAnnotation_Expecter_Methods.SetMock(mock)
}

func (eAnnotation *MockEAnnotation_Prototype) EXPECT() *MockEAnnotation_Expecter {
	expecter := &MockEAnnotation_Expecter{}
	expecter.SetMock(eAnnotation.mock)
	return expecter
}

// MockEAnnotation_Prototype_Methods is the mock implementation of EAnnotation declared methods
type MockEAnnotation_Prototype_Methods struct {
	mock *mock.Mock
}

func (_mdp *MockEAnnotation_Prototype_Methods) SetMock(mock *mock.Mock) {
	_mdp.mock = mock
}

// MockEAnnotation_Expecter_Methods is the expecter implementation of EAnnotation declared methods
type MockEAnnotation_Expecter_Methods struct {
	mock *mock.Mock
}

func (_mde *MockEAnnotation_Expecter_Methods) SetMock(mock *mock.Mock) {
	_mde.mock = mock
}

// GetContents get the value of contents
func (eAnnotation *MockEAnnotation_Prototype_Methods) GetContents() EList {
	ret := eAnnotation.mock.Called()

	var r EList
	if rf, ok := ret.Get(0).(func() EList); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EList)
		}
	}

	return r
}

type MockEAnnotation_GetContents_Call struct {
	*mock.Call
}

func (e *MockEAnnotation_Expecter_Methods) GetContents() *MockEAnnotation_GetContents_Call {
	return &MockEAnnotation_GetContents_Call{Call: e.mock.On("GetContents")}
}

func (c *MockEAnnotation_GetContents_Call) Run(run func()) *MockEAnnotation_GetContents_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEAnnotation_GetContents_Call) Return(contents EList) *MockEAnnotation_GetContents_Call {
	c.Call.Return(contents)
	return c
}

// GetDetails get the value of details
func (eAnnotation *MockEAnnotation_Prototype_Methods) GetDetails() EMap {
	ret := eAnnotation.mock.Called()

	var r EMap
	if rf, ok := ret.Get(0).(func() EMap); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EMap)
		}
	}

	return r
}

type MockEAnnotation_GetDetails_Call struct {
	*mock.Call
}

func (e *MockEAnnotation_Expecter_Methods) GetDetails() *MockEAnnotation_GetDetails_Call {
	return &MockEAnnotation_GetDetails_Call{Call: e.mock.On("GetDetails")}
}

func (c *MockEAnnotation_GetDetails_Call) Run(run func()) *MockEAnnotation_GetDetails_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEAnnotation_GetDetails_Call) Return(details EMap) *MockEAnnotation_GetDetails_Call {
	c.Call.Return(details)
	return c
}

// GetEModelElement get the value of eModelElement
func (eAnnotation *MockEAnnotation_Prototype_Methods) GetEModelElement() EModelElement {
	ret := eAnnotation.mock.Called()

	var r EModelElement
	if rf, ok := ret.Get(0).(func() EModelElement); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EModelElement)
		}
	}

	return r
}

type MockEAnnotation_GetEModelElement_Call struct {
	*mock.Call
}

func (e *MockEAnnotation_Expecter_Methods) GetEModelElement() *MockEAnnotation_GetEModelElement_Call {
	return &MockEAnnotation_GetEModelElement_Call{Call: e.mock.On("GetEModelElement")}
}

func (c *MockEAnnotation_GetEModelElement_Call) Run(run func()) *MockEAnnotation_GetEModelElement_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEAnnotation_GetEModelElement_Call) Return(eModelElement EModelElement) *MockEAnnotation_GetEModelElement_Call {
	c.Call.Return(eModelElement)
	return c
}

// SetEModelElement provides mock implementation for setting the value of eModelElement
func (eAnnotation *MockEAnnotation_Prototype_Methods) SetEModelElement(eModelElement EModelElement) {
	eAnnotation.mock.Called(eModelElement)
}

type MockEAnnotation_SetEModelElement_Call struct {
	*mock.Call
}

// SetEModelElement is a helper method to define mock.On call
// - eModelElement EModelElement
func (e *MockEAnnotation_Expecter_Methods) SetEModelElement(eModelElement any) *MockEAnnotation_SetEModelElement_Call {
	return &MockEAnnotation_SetEModelElement_Call{Call: e.mock.On("SetEModelElement", eModelElement)}
}

func (c *MockEAnnotation_SetEModelElement_Call) Run(run func(eModelElement EModelElement)) *MockEAnnotation_SetEModelElement_Call {
	c.Call.Run(func(args mock.Arguments) {
		run(args[0].(EModelElement))
	})
	return c
}

func (c *MockEAnnotation_SetEModelElement_Call) Return() *MockEAnnotation_SetEModelElement_Call {
	c.Call.Return()
	return c
}

// GetReferences get the value of references
func (eAnnotation *MockEAnnotation_Prototype_Methods) GetReferences() EList {
	ret := eAnnotation.mock.Called()

	var r EList
	if rf, ok := ret.Get(0).(func() EList); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EList)
		}
	}

	return r
}

type MockEAnnotation_GetReferences_Call struct {
	*mock.Call
}

func (e *MockEAnnotation_Expecter_Methods) GetReferences() *MockEAnnotation_GetReferences_Call {
	return &MockEAnnotation_GetReferences_Call{Call: e.mock.On("GetReferences")}
}

func (c *MockEAnnotation_GetReferences_Call) Run(run func()) *MockEAnnotation_GetReferences_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEAnnotation_GetReferences_Call) Return(references EList) *MockEAnnotation_GetReferences_Call {
	c.Call.Return(references)
	return c
}

// GetSource get the value of source
func (eAnnotation *MockEAnnotation_Prototype_Methods) GetSource() string {
	ret := eAnnotation.mock.Called()

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

type MockEAnnotation_GetSource_Call struct {
	*mock.Call
}

func (e *MockEAnnotation_Expecter_Methods) GetSource() *MockEAnnotation_GetSource_Call {
	return &MockEAnnotation_GetSource_Call{Call: e.mock.On("GetSource")}
}

func (c *MockEAnnotation_GetSource_Call) Run(run func()) *MockEAnnotation_GetSource_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEAnnotation_GetSource_Call) Return(source string) *MockEAnnotation_GetSource_Call {
	c.Call.Return(source)
	return c
}

// SetSource provides mock implementation for setting the value of source
func (eAnnotation *MockEAnnotation_Prototype_Methods) SetSource(source string) {
	eAnnotation.mock.Called(source)
}

type MockEAnnotation_SetSource_Call struct {
	*mock.Call
}

// SetSource is a helper method to define mock.On call
// - source string
func (e *MockEAnnotation_Expecter_Methods) SetSource(source any) *MockEAnnotation_SetSource_Call {
	return &MockEAnnotation_SetSource_Call{Call: e.mock.On("SetSource", source)}
}

func (c *MockEAnnotation_SetSource_Call) Run(run func(source string)) *MockEAnnotation_SetSource_Call {
	c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return c
}

func (c *MockEAnnotation_SetSource_Call) Return() *MockEAnnotation_SetSource_Call {
	c.Call.Return()
	return c
}

type mockConstructorTestingTNewMockEAnnotation interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockEAnnotation creates a new instance of MockEAnnotation. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockEAnnotation(t mockConstructorTestingTNewMockEAnnotation) *MockEAnnotation {
	mock := &MockEAnnotation{}
	mock.SetMock(&mock.Mock)
	mock.Mock.Test(t)
	t.Cleanup(func() { mock.AssertExpectations(t) })
	return mock
}
