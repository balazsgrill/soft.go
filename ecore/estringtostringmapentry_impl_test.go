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

func discardEStringToStringMapEntry() {
	_ = assert.Equal
	_ = mock.Anything
	_ = testing.Coverage
}

func TestEStringToStringMapEntryAsEStringToStringMapEntry(t *testing.T) {
	o := newEStringToStringMapEntryImpl()
	assert.Equal(t, o, o.asEStringToStringMapEntry())
}

func TestEStringToStringMapEntryStaticClass(t *testing.T) {
	o := newEStringToStringMapEntryImpl()
	assert.Equal(t, GetPackage().GetEStringToStringMapEntry(), o.EStaticClass())
}

func TestEStringToStringMapEntryFeatureCount(t *testing.T) {
	o := newEStringToStringMapEntryImpl()
	assert.Equal(t, ESTRING_TO_STRING_MAP_ENTRY_FEATURE_COUNT, o.EStaticFeatureCount())
}

func TestEStringToStringMapEntryGetKey(t *testing.T) {
	o := newEStringToStringMapEntryImpl()
	assert.Equal(t, o.GetTypedKey(), o.GetKey())
}

func TestEStringToStringMapEntryGetValue(t *testing.T) {
	o := newEStringToStringMapEntryImpl()
	assert.Equal(t, o.GetTypedValue(), o.GetValue())
}

func TestEStringToStringMapEntryKeyGet(t *testing.T) {
	o := newEStringToStringMapEntryImpl()
	// get default value
	assert.Equal(t, string(""), o.GetTypedKey())
	// get initialized value
	v := string("Test String")
	o.key = v
	assert.Equal(t, v, o.GetTypedKey())
}

func TestEStringToStringMapEntryKeySet(t *testing.T) {
	o := newEStringToStringMapEntryImpl()
	v := string("Test String")
	mockAdapter := NewMockEAdapter(t)
	mockAdapter.EXPECT().SetTarget(o).Once()
	mockAdapter.EXPECT().NotifyChanged(mock.Anything).Once()
	o.EAdapters().Add(mockAdapter)
	o.SetTypedKey(v)
	mockAdapter.AssertExpectations(t)
}

func TestEStringToStringMapEntryValueGet(t *testing.T) {
	o := newEStringToStringMapEntryImpl()
	// get default value
	assert.Equal(t, string(""), o.GetTypedValue())
	// get initialized value
	v := string("Test String")
	o.value = v
	assert.Equal(t, v, o.GetTypedValue())
}

func TestEStringToStringMapEntryValueSet(t *testing.T) {
	o := newEStringToStringMapEntryImpl()
	v := string("Test String")
	mockAdapter := NewMockEAdapter(t)
	mockAdapter.EXPECT().SetTarget(o).Once()
	mockAdapter.EXPECT().NotifyChanged(mock.Anything).Once()
	o.EAdapters().Add(mockAdapter)
	o.SetTypedValue(v)
	mockAdapter.AssertExpectations(t)
}

func TestEStringToStringMapEntryEGetFromID(t *testing.T) {
	o := newEStringToStringMapEntryImpl()
	assert.Panics(t, func() { o.EGetFromID(-1, true) })
	assert.Equal(t, o.GetTypedKey(), o.EGetFromID(ESTRING_TO_STRING_MAP_ENTRY__KEY, true))
	assert.Equal(t, o.GetTypedValue(), o.EGetFromID(ESTRING_TO_STRING_MAP_ENTRY__VALUE, true))
}

func TestEStringToStringMapEntryESetFromID(t *testing.T) {
	o := newEStringToStringMapEntryImpl()
	assert.Panics(t, func() { o.ESetFromID(-1, nil) })
	{
		v := string("Test String")
		o.ESetFromID(ESTRING_TO_STRING_MAP_ENTRY__KEY, v)
		assert.Equal(t, v, o.EGetFromID(ESTRING_TO_STRING_MAP_ENTRY__KEY, false))
	}
	{
		v := string("Test String")
		o.ESetFromID(ESTRING_TO_STRING_MAP_ENTRY__VALUE, v)
		assert.Equal(t, v, o.EGetFromID(ESTRING_TO_STRING_MAP_ENTRY__VALUE, false))
	}

}

func TestEStringToStringMapEntryEIsSetFromID(t *testing.T) {
	o := newEStringToStringMapEntryImpl()
	assert.Panics(t, func() { o.EIsSetFromID(-1) })
	assert.False(t, o.EIsSetFromID(ESTRING_TO_STRING_MAP_ENTRY__KEY))
	assert.False(t, o.EIsSetFromID(ESTRING_TO_STRING_MAP_ENTRY__VALUE))
}

func TestEStringToStringMapEntryEUnsetFromID(t *testing.T) {
	o := newEStringToStringMapEntryImpl()
	assert.Panics(t, func() { o.EUnsetFromID(-1) })
	{
		o.EUnsetFromID(ESTRING_TO_STRING_MAP_ENTRY__KEY)
		v := o.EGetFromID(ESTRING_TO_STRING_MAP_ENTRY__KEY, false)
		assert.Equal(t, string(""), v)
	}
	{
		o.EUnsetFromID(ESTRING_TO_STRING_MAP_ENTRY__VALUE)
		v := o.EGetFromID(ESTRING_TO_STRING_MAP_ENTRY__VALUE, false)
		assert.Equal(t, string(""), v)
	}
}
