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

func discardMockEStringToStringMapEntry() {
	_ = assert.Equal
	_ = testing.Coverage
}

type mockEStringToStringMapEntryRun struct {
	mock.Mock
}

func (m *mockEStringToStringMapEntryRun) Run(args ...any) {
	m.Called(args...)
}

type mockConstructorTestingTmockEStringToStringMapEntryRun interface {
	mock.TestingT
	Cleanup(func())
}

// newMockEStringToStringMapEntryRun creates a new instance of mockEStringToStringMapEntryRun. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newMockEStringToStringMapEntryRun(t mockConstructorTestingTmockEStringToStringMapEntryRun, args ...any) *mockEStringToStringMapEntryRun {
	mock := &mockEStringToStringMapEntryRun{}
	mock.Test(t)
	mock.On("Run", args...).Once()
	t.Cleanup(func() { mock.AssertExpectations(t) })
	return mock
}

// TestMockEStringToStringMapEntryGetTypedKey tests method GetTypedKey
func TestMockEStringToStringMapEntryGetTypedKey(t *testing.T) {
	o := NewMockEStringToStringMapEntry(t)
	r := string("Test String")
	m := newMockEStringToStringMapEntryRun(t)
	o.EXPECT().GetTypedKey().Run(func() { m.Run() }).Return(r).Once()
	o.EXPECT().GetTypedKey().Once().Return(func() string { return r })
	assert.Equal(t, r, o.GetTypedKey())
	assert.Equal(t, r, o.GetTypedKey())
}

// TestMockEStringToStringMapEntrySetTypedKey tests method SetTypedKey
func TestMockEStringToStringMapEntrySetTypedKey(t *testing.T) {
	o := NewMockEStringToStringMapEntry(t)
	v := string("Test String")
	m := newMockEStringToStringMapEntryRun(t, v)
	o.EXPECT().SetTypedKey(v).Run(func(_p0 string) { m.Run(_p0) }).Once()
	o.SetTypedKey(v)
}

// TestMockEStringToStringMapEntryGetTypedValue tests method GetTypedValue
func TestMockEStringToStringMapEntryGetTypedValue(t *testing.T) {
	o := NewMockEStringToStringMapEntry(t)
	r := string("Test String")
	m := newMockEStringToStringMapEntryRun(t)
	o.EXPECT().GetTypedValue().Run(func() { m.Run() }).Return(r).Once()
	o.EXPECT().GetTypedValue().Once().Return(func() string { return r })
	assert.Equal(t, r, o.GetTypedValue())
	assert.Equal(t, r, o.GetTypedValue())
}

// TestMockEStringToStringMapEntrySetTypedValue tests method SetTypedValue
func TestMockEStringToStringMapEntrySetTypedValue(t *testing.T) {
	o := NewMockEStringToStringMapEntry(t)
	v := string("Test String")
	m := newMockEStringToStringMapEntryRun(t, v)
	o.EXPECT().SetTypedValue(v).Run(func(_p0 string) { m.Run(_p0) }).Once()
	o.SetTypedValue(v)
}
