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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func discardMockENamedElement() {
	_ = assert.Equal
	_ = testing.Coverage
}

type mockENamedElementRun struct {
	mock.Mock
}

func (m *mockENamedElementRun) Run(args ...any) {
	m.Called(args...)
}

type mockConstructorTestingTmockENamedElementRun interface {
	mock.TestingT
	Cleanup(func())
}

// newMockENamedElementRun creates a new instance of mockENamedElementRun. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newMockENamedElementRun(t mockConstructorTestingTmockENamedElementRun, args ...any) *mockENamedElementRun {
	mock := &mockENamedElementRun{}
	mock.Test(t)
	mock.On("Run", args...).Once()
	t.Cleanup(func() { mock.AssertExpectations(t) })
	return mock
}

// TestMockENamedElementGetName tests method GetName
func TestMockENamedElementGetName(t *testing.T) {
	o := NewMockENamedElement(t)
	r := string("Test String")
	m := newMockENamedElementRun(t)
	o.EXPECT().GetName().Run(func() { m.Run() }).Return(r).Once()
	o.EXPECT().GetName().Once().Return(func() string { return r })
	assert.Equal(t, r, o.GetName())
	assert.Equal(t, r, o.GetName())
}

// TestMockENamedElementSetName tests method SetName
func TestMockENamedElementSetName(t *testing.T) {
	o := NewMockENamedElement(t)
	v := string("Test String")
	m := newMockENamedElementRun(t, v)
	o.EXPECT().SetName(v).Run(func(_p0 string) { m.Run(_p0) }).Once()
	o.SetName(v)
}
