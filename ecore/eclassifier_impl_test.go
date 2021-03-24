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
	"reflect"
	"testing"
)

func discardEClassifier() {
	_ = assert.Equal
	_ = mock.Anything
	_ = testing.Coverage
	_ = url.Parse
}

func TestEClassifierAsEClassifier(t *testing.T) {
	o := newEClassifierImpl()
	assert.Equal(t, o, o.asEClassifier())
}

func TestEClassifierStaticClass(t *testing.T) {
	o := newEClassifierImpl()
	assert.Equal(t, GetPackage().GetEClassifierClass(), o.EStaticClass())
}

func TestEClassifierFeatureCount(t *testing.T) {
	o := newEClassifierImpl()
	assert.Equal(t, ECLASSIFIER_FEATURE_COUNT, o.EStaticFeatureCount())
}

func TestEClassifierClassifierIDGet(t *testing.T) {
	o := newEClassifierImpl()
	// get default value
	assert.Equal(t, int(-1), o.GetClassifierID())
	// get initialized value
	v := int(45)
	o.classifierID = v
	assert.Equal(t, v, o.GetClassifierID())
}

func TestEClassifierClassifierIDSet(t *testing.T) {
	o := newEClassifierImpl()
	v := int(45)
	mockAdapter := new(MockEAdapter)
	mockAdapter.On("SetTarget", o).Once()
	mockAdapter.On("NotifyChanged", mock.Anything).Once()
	o.EAdapters().Add(mockAdapter)
	o.SetClassifierID(v)
	mockAdapter.AssertExpectations(t)
}

func TestEClassifierDefaultValueGet(t *testing.T) {
	o := newEClassifierImpl()
	assert.Panics(t, func() { o.GetDefaultValue() })
}

func TestEClassifierEPackageGet(t *testing.T) {
	// default
	o := newEClassifierImpl()
	assert.Nil(t, o.GetEPackage())

	// set a mock container
	v := new(MockEPackage)
	o.ESetInternalContainer(v, ECLASSIFIER__EPACKAGE)

	// no proxy
	v.On("EIsProxy").Return(false)
	assert.Equal(t, v, o.GetEPackage())
}

func TestEClassifierInstanceClassGet(t *testing.T) {
	o := newEClassifierImpl()
	// get default value
	assert.Nil(t, o.GetInstanceClass())
	// get initialized value
	v := reflect.Type(reflect.TypeOf(""))
	o.instanceClass = v
	assert.Equal(t, v, o.GetInstanceClass())
}

func TestEClassifierInstanceClassSet(t *testing.T) {
	o := newEClassifierImpl()
	v := reflect.Type(reflect.TypeOf(""))
	mockAdapter := new(MockEAdapter)
	mockAdapter.On("SetTarget", o).Once()
	mockAdapter.On("NotifyChanged", mock.Anything).Once()
	o.EAdapters().Add(mockAdapter)
	o.SetInstanceClass(v)
	mockAdapter.AssertExpectations(t)
}

func TestEClassifierInstanceTypeNameGet(t *testing.T) {
	o := newEClassifierImpl()
	// get default value
	assert.Equal(t, string(""), o.GetInstanceTypeName())
	// get initialized value
	v := string("Test String")
	o.instanceTypeName = v
	assert.Equal(t, v, o.GetInstanceTypeName())
}

func TestEClassifierInstanceTypeNameSet(t *testing.T) {
	o := newEClassifierImpl()
	v := string("Test String")
	mockAdapter := new(MockEAdapter)
	mockAdapter.On("SetTarget", o).Once()
	mockAdapter.On("NotifyChanged", mock.Anything).Once()
	o.EAdapters().Add(mockAdapter)
	o.SetInstanceTypeName(v)
	mockAdapter.AssertExpectations(t)
}

func TestEClassifierIsInstanceOperation(t *testing.T) {
	o := newEClassifierImpl()
	assert.Panics(t, func() { o.IsInstance(nil) })
}

func TestEClassifierEGetFromID(t *testing.T) {
	o := newEClassifierImpl()
	assert.Panics(t, func() { o.EGetFromID(-1, true) })
	assert.Equal(t, o.GetClassifierID(), o.EGetFromID(ECLASSIFIER__CLASSIFIER_ID, true))
	assert.Panics(t, func() { o.EGetFromID(ECLASSIFIER__DEFAULT_VALUE, true) })
	assert.Panics(t, func() { o.EGetFromID(ECLASSIFIER__DEFAULT_VALUE, false) })
	assert.Equal(t, o.GetEPackage(), o.EGetFromID(ECLASSIFIER__EPACKAGE, true))
	assert.Equal(t, o.GetInstanceClass(), o.EGetFromID(ECLASSIFIER__INSTANCE_CLASS, true))
	assert.Equal(t, o.GetInstanceTypeName(), o.EGetFromID(ECLASSIFIER__INSTANCE_TYPE_NAME, true))
}

func TestEClassifierESetFromID(t *testing.T) {
	o := newEClassifierImpl()
	assert.Panics(t, func() { o.ESetFromID(-1, nil) })
	{
		v := int(45)
		o.ESetFromID(ECLASSIFIER__CLASSIFIER_ID, v)
		assert.Equal(t, v, o.EGetFromID(ECLASSIFIER__CLASSIFIER_ID, false))
	}
	{
		v := reflect.Type(reflect.TypeOf(""))
		o.ESetFromID(ECLASSIFIER__INSTANCE_CLASS, v)
		assert.Equal(t, v, o.EGetFromID(ECLASSIFIER__INSTANCE_CLASS, false))
	}
	{
		v := string("Test String")
		o.ESetFromID(ECLASSIFIER__INSTANCE_TYPE_NAME, v)
		assert.Equal(t, v, o.EGetFromID(ECLASSIFIER__INSTANCE_TYPE_NAME, false))
	}

}

func TestEClassifierEIsSetFromID(t *testing.T) {
	o := newEClassifierImpl()
	assert.Panics(t, func() { o.EIsSetFromID(-1) })
	assert.False(t, o.EIsSetFromID(ECLASSIFIER__CLASSIFIER_ID))
	assert.Panics(t, func() { o.EIsSetFromID(ECLASSIFIER__DEFAULT_VALUE) })
	assert.False(t, o.EIsSetFromID(ECLASSIFIER__EPACKAGE))
	assert.False(t, o.EIsSetFromID(ECLASSIFIER__INSTANCE_CLASS))
	assert.False(t, o.EIsSetFromID(ECLASSIFIER__INSTANCE_TYPE_NAME))
}

func TestEClassifierEUnsetFromID(t *testing.T) {
	o := newEClassifierImpl()
	assert.Panics(t, func() { o.EUnsetFromID(-1) })
	{
		o.EUnsetFromID(ECLASSIFIER__CLASSIFIER_ID)
		v := o.EGetFromID(ECLASSIFIER__CLASSIFIER_ID, false)
		assert.Equal(t, int(-1), v)
	}
	{
		o.EUnsetFromID(ECLASSIFIER__INSTANCE_CLASS)
		v := o.EGetFromID(ECLASSIFIER__INSTANCE_CLASS, false)
		assert.Nil(t, v)
	}
	{
		o.EUnsetFromID(ECLASSIFIER__INSTANCE_TYPE_NAME)
		v := o.EGetFromID(ECLASSIFIER__INSTANCE_TYPE_NAME, false)
		assert.Equal(t, string(""), v)
	}
}

func TestEClassifierEInvokeFromID(t *testing.T) {
	o := newEClassifierImpl()
	assert.Panics(t, func() { o.EInvokeFromID(-1, nil) })
	assert.Panics(t, func() { o.EInvokeFromID(ECLASSIFIER__IS_INSTANCE_EJAVAOBJECT, nil) })
}

func TestEClassifierEBasicInverseAdd(t *testing.T) {
	o := newEClassifierImpl()
	{
		mockObject := new(MockEObject)
		mockNotifications := new(MockENotificationChain)
		assert.Equal(t, mockNotifications, o.EBasicInverseAdd(mockObject, -1, mockNotifications))
	}
	{
		mockObject := new(MockEPackage)
		mockObject.On("EResource").Return(nil).Once()
		mockObject.On("EIsProxy").Return(false).Once()
		o.EBasicInverseAdd(mockObject, ECLASSIFIER__EPACKAGE, nil)
		assert.Equal(t, mockObject, o.GetEPackage())
		mock.AssertExpectationsForObjects(t, mockObject)

		mockOther := new(MockEPackage)
		mockOther.On("EResource").Return(nil).Once()
		mockOther.On("EIsProxy").Return(false).Once()
		mockObject.On("EResource").Return(nil).Once()
		mockObject.On("EInverseRemove", o, EPACKAGE__ECLASSIFIERS, nil).Return(nil).Once()
		o.EBasicInverseAdd(mockOther, ECLASSIFIER__EPACKAGE, nil)
		assert.Equal(t, mockOther, o.GetEPackage())
		mock.AssertExpectationsForObjects(t, mockObject, mockOther)
	}

}

func TestEClassifierEBasicInverseRemove(t *testing.T) {
	o := newEClassifierImpl()
	{
		mockObject := new(MockEObject)
		mockNotifications := new(MockENotificationChain)
		assert.Equal(t, mockNotifications, o.EBasicInverseRemove(mockObject, -1, mockNotifications))
	}
	{
		mockObject := new(MockEPackage)
		o.EBasicInverseRemove(mockObject, ECLASSIFIER__EPACKAGE, nil)
		mock.AssertExpectationsForObjects(t, mockObject)
	}

}
