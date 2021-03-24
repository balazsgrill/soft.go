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
	"net/url"
	"testing"
)

func discardEDataType() {
	_ = assert.Equal
	_ = mock.Anything
	_ = testing.Coverage
	_ = url.Parse
}

func TestEDataTypeAsEDataType(t *testing.T) {
	o := newEDataTypeImpl()
	assert.Equal(t, o, o.asEDataType())
}

func TestEDataTypeStaticClass(t *testing.T) {
	o := newEDataTypeImpl()
	assert.Equal(t, GetPackage().GetEDataType(), o.EStaticClass())
}

func TestEDataTypeFeatureCount(t *testing.T) {
	o := newEDataTypeImpl()
	assert.Equal(t, EDATA_TYPE_FEATURE_COUNT, o.EStaticFeatureCount())
}

func TestEDataTypeSerializableGet(t *testing.T) {
	o := newEDataTypeImpl()
	// get default value
	assert.Equal(t, bool(true), o.IsSerializable())
	// get initialized value
	v := bool(true)
	o.isSerializable = v
	assert.Equal(t, v, o.IsSerializable())
}

func TestEDataTypeSerializableSet(t *testing.T) {
	o := newEDataTypeImpl()
	v := bool(true)
	mockAdapter := new(MockEAdapter)
	mockAdapter.On("SetTarget", o).Once()
	mockAdapter.On("NotifyChanged", mock.Anything).Once()
	o.EAdapters().Add(mockAdapter)
	o.SetSerializable(v)
	mockAdapter.AssertExpectations(t)
}

func TestEDataTypeEGetFromID(t *testing.T) {
	o := newEDataTypeImpl()
	assert.Panics(t, func() { o.EGetFromID(-1, true) })
	assert.Equal(t, o.IsSerializable(), o.EGetFromID(EDATA_TYPE__SERIALIZABLE, true))
}

func TestEDataTypeESetFromID(t *testing.T) {
	o := newEDataTypeImpl()
	assert.Panics(t, func() { o.ESetFromID(-1, nil) })
	{
		v := bool(true)
		o.ESetFromID(EDATA_TYPE__SERIALIZABLE, v)
		assert.Equal(t, v, o.EGetFromID(EDATA_TYPE__SERIALIZABLE, false))
	}

}

func TestEDataTypeEIsSetFromID(t *testing.T) {
	o := newEDataTypeImpl()
	assert.Panics(t, func() { o.EIsSetFromID(-1) })
	assert.False(t, o.EIsSetFromID(EDATA_TYPE__SERIALIZABLE))
}

func TestEDataTypeEUnsetFromID(t *testing.T) {
	o := newEDataTypeImpl()
	assert.Panics(t, func() { o.EUnsetFromID(-1) })
	{
		o.EUnsetFromID(EDATA_TYPE__SERIALIZABLE)
		v := o.EGetFromID(EDATA_TYPE__SERIALIZABLE, false)
		assert.Equal(t, bool(true), v)
	}
}
