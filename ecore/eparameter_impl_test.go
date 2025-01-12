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

func discardEParameter() {
	_ = assert.Equal
	_ = mock.Anything
	_ = testing.Coverage
}

func TestEParameterAsEParameter(t *testing.T) {
	o := newEParameterImpl()
	assert.Equal(t, o, o.asEParameter())
}

func TestEParameterStaticClass(t *testing.T) {
	o := newEParameterImpl()
	assert.Equal(t, GetPackage().GetEParameter(), o.EStaticClass())
}

func TestEParameterFeatureCount(t *testing.T) {
	o := newEParameterImpl()
	assert.Equal(t, EPARAMETER_FEATURE_COUNT, o.EStaticFeatureCount())
}

func TestEParameterEOperationGet(t *testing.T) {
	// default
	o := newEParameterImpl()
	assert.Nil(t, o.GetEOperation())

	// set a mock container
	v := NewMockEOperation(t)
	o.ESetInternalContainer(v, EPARAMETER__EOPERATION)

	// no proxy
	v.EXPECT().EIsProxy().Return(false).Once()
	assert.Equal(t, v, o.GetEOperation())
}

func TestEParameterEGetFromID(t *testing.T) {
	o := newEParameterImpl()
	assert.Panics(t, func() { o.EGetFromID(-1, true) })
	assert.Equal(t, o.GetEOperation(), o.EGetFromID(EPARAMETER__EOPERATION, true))
}

func TestEParameterEIsSetFromID(t *testing.T) {
	o := newEParameterImpl()
	assert.Panics(t, func() { o.EIsSetFromID(-1) })
	assert.False(t, o.EIsSetFromID(EPARAMETER__EOPERATION))
}

func TestEParameterEBasicInverseAdd(t *testing.T) {
	o := newEParameterImpl()
	{
		mockObject := NewMockEObject(t)
		mockNotifications := NewMockENotificationChain(t)
		assert.Equal(t, mockNotifications, o.EBasicInverseAdd(mockObject, -1, mockNotifications))
	}
	{
		mockObject := NewMockEOperation(t)
		mockObject.EXPECT().EResource().Return(nil).Once()
		mockObject.EXPECT().EIsProxy().Return(false).Once()
		o.EBasicInverseAdd(mockObject, EPARAMETER__EOPERATION, nil)
		assert.Equal(t, mockObject, o.GetEOperation())
		mock.AssertExpectationsForObjects(t, mockObject)

		mockOther := NewMockEOperation(t)
		mockOther.EXPECT().EResource().Return(nil).Once()
		mockOther.EXPECT().EIsProxy().Return(false).Once()
		mockObject.EXPECT().EResource().Return(nil).Once()
		mockObject.EXPECT().EInverseRemove(o, EOPERATION__EPARAMETERS, nil).Return(nil).Once()
		o.EBasicInverseAdd(mockOther, EPARAMETER__EOPERATION, nil)
		assert.Equal(t, mockOther, o.GetEOperation())
		mock.AssertExpectationsForObjects(t, mockObject, mockOther)
	}

}

func TestEParameterEBasicInverseRemove(t *testing.T) {
	o := newEParameterImpl()
	{
		mockObject := NewMockEObject(t)
		mockNotifications := NewMockENotificationChain(t)
		assert.Equal(t, mockNotifications, o.EBasicInverseRemove(mockObject, -1, mockNotifications))
	}
	{
		mockObject := NewMockEOperation(t)
		o.EBasicInverseRemove(mockObject, EPARAMETER__EOPERATION, nil)
		mock.AssertExpectationsForObjects(t, mockObject)
	}

}
