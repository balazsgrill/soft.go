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

func discardMockEGenericType() {
	_ = assert.Equal
	_ = testing.Coverage
}

type mockEGenericTypeRun struct {
	mock.Mock
}

func (m *mockEGenericTypeRun) Run(args ...any) {
	m.Called(args...)
}

type mockConstructorTestingTmockEGenericTypeRun interface {
	mock.TestingT
	Cleanup(func())
}

// newMockEGenericTypeRun creates a new instance of mockEGenericTypeRun. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newMockEGenericTypeRun(t mockConstructorTestingTmockEGenericTypeRun, args ...any) *mockEGenericTypeRun {
	mock := &mockEGenericTypeRun{}
	mock.Test(t)
	mock.On("Run", args...).Once()
	t.Cleanup(func() { mock.AssertExpectations(t) })
	return mock
}

// TestMockEGenericTypeGetEClassifier tests method GetEClassifier
func TestMockEGenericTypeGetEClassifier(t *testing.T) {
	o := NewMockEGenericType(t)
	r := new(MockEClassifier)
	m := newMockEGenericTypeRun(t)
	o.EXPECT().GetEClassifier().Return(r).Run(func() { m.Run() }).Once()
	o.EXPECT().GetEClassifier().Once().Return(func() EClassifier { return r })
	assert.Equal(t, r, o.GetEClassifier())
	assert.Equal(t, r, o.GetEClassifier())
}

// TestMockEGenericTypeSetEClassifier tests method SetEClassifier
func TestMockEGenericTypeSetEClassifier(t *testing.T) {
	o := NewMockEGenericType(t)
	v := new(MockEClassifier)
	m := newMockEGenericTypeRun(t, v)
	o.EXPECT().SetEClassifier(v).Return().Run(func(_p0 EClassifier) { m.Run(_p0) }).Once()
	o.SetEClassifier(v)
}

// TestMockEGenericTypeGetELowerBound tests method GetELowerBound
func TestMockEGenericTypeGetELowerBound(t *testing.T) {
	o := NewMockEGenericType(t)
	r := new(MockEGenericType)
	m := newMockEGenericTypeRun(t)
	o.EXPECT().GetELowerBound().Return(r).Run(func() { m.Run() }).Once()
	o.EXPECT().GetELowerBound().Once().Return(func() EGenericType { return r })
	assert.Equal(t, r, o.GetELowerBound())
	assert.Equal(t, r, o.GetELowerBound())
}

// TestMockEGenericTypeSetELowerBound tests method SetELowerBound
func TestMockEGenericTypeSetELowerBound(t *testing.T) {
	o := NewMockEGenericType(t)
	v := new(MockEGenericType)
	m := newMockEGenericTypeRun(t, v)
	o.EXPECT().SetELowerBound(v).Return().Run(func(_p0 EGenericType) { m.Run(_p0) }).Once()
	o.SetELowerBound(v)
}

// TestMockEGenericTypeGetERawType tests method GetERawType
func TestMockEGenericTypeGetERawType(t *testing.T) {
	o := NewMockEGenericType(t)
	r := new(MockEClassifier)
	m := newMockEGenericTypeRun(t)
	o.EXPECT().GetERawType().Return(r).Run(func() { m.Run() }).Once()
	o.EXPECT().GetERawType().Once().Return(func() EClassifier { return r })
	assert.Equal(t, r, o.GetERawType())
	assert.Equal(t, r, o.GetERawType())
}

// TestMockEGenericTypeGetETypeArguments tests method GetETypeArguments
func TestMockEGenericTypeGetETypeArguments(t *testing.T) {
	o := NewMockEGenericType(t)
	l := NewMockEList(t)
	m := newMockEGenericTypeRun(t)
	o.EXPECT().GetETypeArguments().Return(l).Run(func() { m.Run() }).Once()
	o.EXPECT().GetETypeArguments().Once().Return(func() EList { return l })
	assert.Equal(t, l, o.GetETypeArguments())
	assert.Equal(t, l, o.GetETypeArguments())
}

// TestMockEGenericTypeGetETypeParameter tests method GetETypeParameter
func TestMockEGenericTypeGetETypeParameter(t *testing.T) {
	o := NewMockEGenericType(t)
	r := new(MockETypeParameter)
	m := newMockEGenericTypeRun(t)
	o.EXPECT().GetETypeParameter().Return(r).Run(func() { m.Run() }).Once()
	o.EXPECT().GetETypeParameter().Once().Return(func() ETypeParameter { return r })
	assert.Equal(t, r, o.GetETypeParameter())
	assert.Equal(t, r, o.GetETypeParameter())
}

// TestMockEGenericTypeSetETypeParameter tests method SetETypeParameter
func TestMockEGenericTypeSetETypeParameter(t *testing.T) {
	o := NewMockEGenericType(t)
	v := new(MockETypeParameter)
	m := newMockEGenericTypeRun(t, v)
	o.EXPECT().SetETypeParameter(v).Return().Run(func(_p0 ETypeParameter) { m.Run(_p0) }).Once()
	o.SetETypeParameter(v)
}

// TestMockEGenericTypeGetEUpperBound tests method GetEUpperBound
func TestMockEGenericTypeGetEUpperBound(t *testing.T) {
	o := NewMockEGenericType(t)
	r := new(MockEGenericType)
	m := newMockEGenericTypeRun(t)
	o.EXPECT().GetEUpperBound().Return(r).Run(func() { m.Run() }).Once()
	o.EXPECT().GetEUpperBound().Once().Return(func() EGenericType { return r })
	assert.Equal(t, r, o.GetEUpperBound())
	assert.Equal(t, r, o.GetEUpperBound())
}

// TestMockEGenericTypeSetEUpperBound tests method SetEUpperBound
func TestMockEGenericTypeSetEUpperBound(t *testing.T) {
	o := NewMockEGenericType(t)
	v := new(MockEGenericType)
	m := newMockEGenericTypeRun(t, v)
	o.EXPECT().SetEUpperBound(v).Return().Run(func(_p0 EGenericType) { m.Run(_p0) }).Once()
	o.SetEUpperBound(v)
}

// TestMockEGenericTypeIsInstance tests method IsInstance
func TestMockEGenericTypeIsInstance(t *testing.T) {
	o := &MockEGenericType{}
	object := any(nil)
	m := newMockEGenericTypeRun(t, object)
	r := bool(true)
	o.EXPECT().IsInstance(object).Return(r).Run(func(object any) { m.Run(object) }).Once()
	o.EXPECT().IsInstance(object).Once().Return(func() bool {
		return r
	})
	assert.Equal(t, r, o.IsInstance(object))
	assert.Equal(t, r, o.IsInstance(object))
	o.AssertExpectations(t)
}
