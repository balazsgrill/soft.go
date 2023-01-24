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

func discardMockEOperation() {
	_ = assert.Equal
	_ = testing.Coverage
}

type mockEOperationRun struct {
	mock.Mock
}

func (m *mockEOperationRun) Run(args ...any) {
	m.Called(args...)
}

type mockConstructorTestingTmockEOperationRun interface {
	mock.TestingT
	Cleanup(func())
}

// newMockEOperationRun creates a new instance of mockEOperationRun. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newMockEOperationRun(t mockConstructorTestingTmockEOperationRun, args ...any) *mockEOperationRun {
	mock := &mockEOperationRun{}
	mock.Test(t)
	mock.On("Run", args...).Once()
	t.Cleanup(func() { mock.AssertExpectations(t) })
	return mock
}

// TestMockEOperationGetEContainingClass tests method GetEContainingClass
func TestMockEOperationGetEContainingClass(t *testing.T) {
	o := NewMockEOperation(t)
	r := new(MockEClass)
	m := newMockEOperationRun(t)
	o.EXPECT().GetEContainingClass().Run(func() { m.Run() }).Return(r).Once()
	o.EXPECT().GetEContainingClass().Once().Return(func() EClass { return r })
	assert.Equal(t, r, o.GetEContainingClass())
	assert.Equal(t, r, o.GetEContainingClass())
}

// TestMockEOperationGetEExceptions tests method GetEExceptions
func TestMockEOperationGetEExceptions(t *testing.T) {
	o := &MockEOperation{}
	l := &MockEList{}
	m := newMockEOperationRun(t)
	o.EXPECT().GetEExceptions().Run(func() { m.Run() }).Return(l).Once()
	o.EXPECT().GetEExceptions().Once().Return(func() EList { return l })
	assert.Equal(t, l, o.GetEExceptions())
	assert.Equal(t, l, o.GetEExceptions())
}

// TestMockEOperationUnsetEExceptions tests method UnsetEExceptions
func TestMockEOperationUnsetEExceptions(t *testing.T) {
	o := NewMockEOperation(t)
	m := newMockEOperationRun(t)
	o.EXPECT().UnsetEExceptions().Run(func() { m.Run() }).Once()
	o.UnsetEExceptions()
}

// TestMockEOperationGetEParameters tests method GetEParameters
func TestMockEOperationGetEParameters(t *testing.T) {
	o := &MockEOperation{}
	l := &MockEList{}
	m := newMockEOperationRun(t)
	o.EXPECT().GetEParameters().Run(func() { m.Run() }).Return(l).Once()
	o.EXPECT().GetEParameters().Once().Return(func() EList { return l })
	assert.Equal(t, l, o.GetEParameters())
	assert.Equal(t, l, o.GetEParameters())
}

// TestMockEOperationGetOperationID tests method GetOperationID
func TestMockEOperationGetOperationID(t *testing.T) {
	o := NewMockEOperation(t)
	r := int(45)
	m := newMockEOperationRun(t)
	o.EXPECT().GetOperationID().Run(func() { m.Run() }).Return(r).Once()
	o.EXPECT().GetOperationID().Once().Return(func() int { return r })
	assert.Equal(t, r, o.GetOperationID())
	assert.Equal(t, r, o.GetOperationID())
}

// TestMockEOperationSetOperationID tests method SetOperationID
func TestMockEOperationSetOperationID(t *testing.T) {
	o := NewMockEOperation(t)
	v := int(45)
	m := newMockEOperationRun(t, v)
	o.EXPECT().SetOperationID(v).Run(func(_p0 int) { m.Run(_p0) }).Once()
	o.SetOperationID(v)
}

// TestMockEOperationIsOverrideOf tests method IsOverrideOf
func TestMockEOperationIsOverrideOf(t *testing.T) {
	o := &MockEOperation{}
	someOperation := new(MockEOperation)
	r := bool(true)
	o.On("IsOverrideOf", someOperation).Return(r).Once()
	o.On("IsOverrideOf", someOperation).Return(func() bool {
		return r
	}).Once()
	assert.Equal(t, r, o.IsOverrideOf(someOperation))
	assert.Equal(t, r, o.IsOverrideOf(someOperation))
	o.AssertExpectations(t)
}
