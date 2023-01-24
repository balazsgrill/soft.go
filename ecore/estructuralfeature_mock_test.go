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
	"reflect"
	"testing"
)

func discardMockEStructuralFeature() {
	_ = assert.Equal
	_ = testing.Coverage
}

type mockEStructuralFeatureRun struct {
	mock.Mock
}

func (m *mockEStructuralFeatureRun) Run(args ...any) {
	m.Called(args...)
}

type mockConstructorTestingTmockEStructuralFeatureRun interface {
	mock.TestingT
	Cleanup(func())
}

// newMockEStructuralFeatureRun creates a new instance of mockEStructuralFeatureRun. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newMockEStructuralFeatureRun(t mockConstructorTestingTmockEStructuralFeatureRun, args ...any) *mockEStructuralFeatureRun {
	mock := &mockEStructuralFeatureRun{}
	mock.Test(t)
	mock.On("Run", args...).Once()
	t.Cleanup(func() { mock.AssertExpectations(t) })
	return mock
}

// TestMockEStructuralFeatureIsChangeable tests method IsChangeable
func TestMockEStructuralFeatureIsChangeable(t *testing.T) {
	o := NewMockEStructuralFeature(t)
	r := bool(true)
	m := newMockEStructuralFeatureRun(t)
	o.EXPECT().IsChangeable().Run(func() { m.Run() }).Return(r).Once()
	o.EXPECT().IsChangeable().Once().Return(func() bool { return r })
	assert.Equal(t, r, o.IsChangeable())
	assert.Equal(t, r, o.IsChangeable())
}

// TestMockEStructuralFeatureSetChangeable tests method SetChangeable
func TestMockEStructuralFeatureSetChangeable(t *testing.T) {
	o := NewMockEStructuralFeature(t)
	v := bool(true)
	m := newMockEStructuralFeatureRun(t, v)
	o.EXPECT().SetChangeable(v).Run(func(_p0 bool) { m.Run(_p0) }).Once()
	o.SetChangeable(v)
}

// TestMockEStructuralFeatureGetDefaultValue tests method GetDefaultValue
func TestMockEStructuralFeatureGetDefaultValue(t *testing.T) {
	o := NewMockEStructuralFeature(t)
	r := any(nil)
	m := newMockEStructuralFeatureRun(t)
	o.EXPECT().GetDefaultValue().Run(func() { m.Run() }).Return(r).Once()
	o.EXPECT().GetDefaultValue().Once().Return(func() any { return r })
	assert.Equal(t, r, o.GetDefaultValue())
	assert.Equal(t, r, o.GetDefaultValue())
}

// TestMockEStructuralFeatureSetDefaultValue tests method SetDefaultValue
func TestMockEStructuralFeatureSetDefaultValue(t *testing.T) {
	o := NewMockEStructuralFeature(t)
	v := any(nil)
	m := newMockEStructuralFeatureRun(t, v)
	o.EXPECT().SetDefaultValue(v).Run(func(_p0 any) { m.Run(_p0) }).Once()
	o.SetDefaultValue(v)
}

// TestMockEStructuralFeatureGetDefaultValueLiteral tests method GetDefaultValueLiteral
func TestMockEStructuralFeatureGetDefaultValueLiteral(t *testing.T) {
	o := NewMockEStructuralFeature(t)
	r := string("Test String")
	m := newMockEStructuralFeatureRun(t)
	o.EXPECT().GetDefaultValueLiteral().Run(func() { m.Run() }).Return(r).Once()
	o.EXPECT().GetDefaultValueLiteral().Once().Return(func() string { return r })
	assert.Equal(t, r, o.GetDefaultValueLiteral())
	assert.Equal(t, r, o.GetDefaultValueLiteral())
}

// TestMockEStructuralFeatureSetDefaultValueLiteral tests method SetDefaultValueLiteral
func TestMockEStructuralFeatureSetDefaultValueLiteral(t *testing.T) {
	o := NewMockEStructuralFeature(t)
	v := string("Test String")
	m := newMockEStructuralFeatureRun(t, v)
	o.EXPECT().SetDefaultValueLiteral(v).Run(func(_p0 string) { m.Run(_p0) }).Once()
	o.SetDefaultValueLiteral(v)
}

// TestMockEStructuralFeatureIsDerived tests method IsDerived
func TestMockEStructuralFeatureIsDerived(t *testing.T) {
	o := NewMockEStructuralFeature(t)
	r := bool(true)
	m := newMockEStructuralFeatureRun(t)
	o.EXPECT().IsDerived().Run(func() { m.Run() }).Return(r).Once()
	o.EXPECT().IsDerived().Once().Return(func() bool { return r })
	assert.Equal(t, r, o.IsDerived())
	assert.Equal(t, r, o.IsDerived())
}

// TestMockEStructuralFeatureSetDerived tests method SetDerived
func TestMockEStructuralFeatureSetDerived(t *testing.T) {
	o := NewMockEStructuralFeature(t)
	v := bool(true)
	m := newMockEStructuralFeatureRun(t, v)
	o.EXPECT().SetDerived(v).Run(func(_p0 bool) { m.Run(_p0) }).Once()
	o.SetDerived(v)
}

// TestMockEStructuralFeatureGetEContainingClass tests method GetEContainingClass
func TestMockEStructuralFeatureGetEContainingClass(t *testing.T) {
	o := NewMockEStructuralFeature(t)
	r := new(MockEClass)
	m := newMockEStructuralFeatureRun(t)
	o.EXPECT().GetEContainingClass().Run(func() { m.Run() }).Return(r).Once()
	o.EXPECT().GetEContainingClass().Once().Return(func() EClass { return r })
	assert.Equal(t, r, o.GetEContainingClass())
	assert.Equal(t, r, o.GetEContainingClass())
}

// TestMockEStructuralFeatureGetFeatureID tests method GetFeatureID
func TestMockEStructuralFeatureGetFeatureID(t *testing.T) {
	o := NewMockEStructuralFeature(t)
	r := int(45)
	m := newMockEStructuralFeatureRun(t)
	o.EXPECT().GetFeatureID().Run(func() { m.Run() }).Return(r).Once()
	o.EXPECT().GetFeatureID().Once().Return(func() int { return r })
	assert.Equal(t, r, o.GetFeatureID())
	assert.Equal(t, r, o.GetFeatureID())
}

// TestMockEStructuralFeatureSetFeatureID tests method SetFeatureID
func TestMockEStructuralFeatureSetFeatureID(t *testing.T) {
	o := NewMockEStructuralFeature(t)
	v := int(45)
	m := newMockEStructuralFeatureRun(t, v)
	o.EXPECT().SetFeatureID(v).Run(func(_p0 int) { m.Run(_p0) }).Once()
	o.SetFeatureID(v)
}

// TestMockEStructuralFeatureIsTransient tests method IsTransient
func TestMockEStructuralFeatureIsTransient(t *testing.T) {
	o := NewMockEStructuralFeature(t)
	r := bool(true)
	m := newMockEStructuralFeatureRun(t)
	o.EXPECT().IsTransient().Run(func() { m.Run() }).Return(r).Once()
	o.EXPECT().IsTransient().Once().Return(func() bool { return r })
	assert.Equal(t, r, o.IsTransient())
	assert.Equal(t, r, o.IsTransient())
}

// TestMockEStructuralFeatureSetTransient tests method SetTransient
func TestMockEStructuralFeatureSetTransient(t *testing.T) {
	o := NewMockEStructuralFeature(t)
	v := bool(true)
	m := newMockEStructuralFeatureRun(t, v)
	o.EXPECT().SetTransient(v).Run(func(_p0 bool) { m.Run(_p0) }).Once()
	o.SetTransient(v)
}

// TestMockEStructuralFeatureIsUnsettable tests method IsUnsettable
func TestMockEStructuralFeatureIsUnsettable(t *testing.T) {
	o := NewMockEStructuralFeature(t)
	r := bool(true)
	m := newMockEStructuralFeatureRun(t)
	o.EXPECT().IsUnsettable().Run(func() { m.Run() }).Return(r).Once()
	o.EXPECT().IsUnsettable().Once().Return(func() bool { return r })
	assert.Equal(t, r, o.IsUnsettable())
	assert.Equal(t, r, o.IsUnsettable())
}

// TestMockEStructuralFeatureSetUnsettable tests method SetUnsettable
func TestMockEStructuralFeatureSetUnsettable(t *testing.T) {
	o := NewMockEStructuralFeature(t)
	v := bool(true)
	m := newMockEStructuralFeatureRun(t, v)
	o.EXPECT().SetUnsettable(v).Run(func(_p0 bool) { m.Run(_p0) }).Once()
	o.SetUnsettable(v)
}

// TestMockEStructuralFeatureIsVolatile tests method IsVolatile
func TestMockEStructuralFeatureIsVolatile(t *testing.T) {
	o := NewMockEStructuralFeature(t)
	r := bool(true)
	m := newMockEStructuralFeatureRun(t)
	o.EXPECT().IsVolatile().Run(func() { m.Run() }).Return(r).Once()
	o.EXPECT().IsVolatile().Once().Return(func() bool { return r })
	assert.Equal(t, r, o.IsVolatile())
	assert.Equal(t, r, o.IsVolatile())
}

// TestMockEStructuralFeatureSetVolatile tests method SetVolatile
func TestMockEStructuralFeatureSetVolatile(t *testing.T) {
	o := NewMockEStructuralFeature(t)
	v := bool(true)
	m := newMockEStructuralFeatureRun(t, v)
	o.EXPECT().SetVolatile(v).Run(func(_p0 bool) { m.Run(_p0) }).Once()
	o.SetVolatile(v)
}

// TestMockEStructuralFeatureGetContainerClass tests method GetContainerClass
func TestMockEStructuralFeatureGetContainerClass(t *testing.T) {
	o := &MockEStructuralFeature{}
	r := reflect.Type(reflect.TypeOf(""))
	o.On("GetContainerClass").Return(r).Once()
	o.On("GetContainerClass").Return(func() reflect.Type {
		return r
	}).Once()
	assert.Equal(t, r, o.GetContainerClass())
	assert.Equal(t, r, o.GetContainerClass())
	o.AssertExpectations(t)
}
