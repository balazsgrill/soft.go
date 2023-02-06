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

// mockXmlProcessorOption is an autogenerated mock type for the xmlProcessorOption type
type mockXmlProcessorOption struct {
	mock.Mock
}

type mockXmlProcessorOption_Expecter struct {
	mock *mock.Mock
}

func (_m *mockXmlProcessorOption) EXPECT() *mockXmlProcessorOption_Expecter {
	return &mockXmlProcessorOption_Expecter{mock: &_m.Mock}
}

// apply provides a mock function with given fields: _a0
func (_m *mockXmlProcessorOption) apply(_a0 *XMLProcessor) {
	_m.Called(_a0)
}

// mockXmlProcessorOption_apply_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'apply'
type mockXmlProcessorOption_apply_Call struct {
	*mock.Call
}

// apply is a helper method to define mock.On call
//   - _a0 *XMLProcessor
func (_e *mockXmlProcessorOption_Expecter) apply(_a0 interface{}) *mockXmlProcessorOption_apply_Call {
	return &mockXmlProcessorOption_apply_Call{Call: _e.mock.On("apply", _a0)}
}

func (_c *mockXmlProcessorOption_apply_Call) Run(run func(_a0 *XMLProcessor)) *mockXmlProcessorOption_apply_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*XMLProcessor))
	})
	return _c
}

func (_c *mockXmlProcessorOption_apply_Call) Return() *mockXmlProcessorOption_apply_Call {
	_c.Call.Return()
	return _c
}

type mockConstructorTestingTnewMockXmlProcessorOption interface {
	mock.TestingT
	Cleanup(func())
}

// newMockXmlProcessorOption creates a new instance of mockXmlProcessorOption. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newMockXmlProcessorOption(t mockConstructorTestingTnewMockXmlProcessorOption) *mockXmlProcessorOption {
	mock := &mockXmlProcessorOption{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
