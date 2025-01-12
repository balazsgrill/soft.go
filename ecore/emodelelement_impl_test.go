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

func discardEModelElement() {
	_ = assert.Equal
	_ = mock.Anything
	_ = testing.Coverage
}

func TestEModelElementAsEModelElement(t *testing.T) {
	o := newEModelElementImpl()
	assert.Equal(t, o, o.asEModelElement())
}

func TestEModelElementStaticClass(t *testing.T) {
	o := newEModelElementImpl()
	assert.Equal(t, GetPackage().GetEModelElement(), o.EStaticClass())
}

func TestEModelElementFeatureCount(t *testing.T) {
	o := newEModelElementImpl()
	assert.Equal(t, EMODEL_ELEMENT_FEATURE_COUNT, o.EStaticFeatureCount())
}

func TestEModelElementEAnnotationsGet(t *testing.T) {
	o := newEModelElementImpl()
	assert.NotNil(t, o.GetEAnnotations())
	assert.Panics(t, func() { _ = o.GetEAnnotations().Get(0).(EAnnotation) })
}

func TestEModelElementGetEAnnotationOperation(t *testing.T) {
	o := newEModelElementImpl()
	assert.Panics(t, func() { o.GetEAnnotation("") })
}

func TestEModelElementEGetFromID(t *testing.T) {
	o := newEModelElementImpl()
	assert.Panics(t, func() { o.EGetFromID(-1, true) })
	assert.Equal(t, o.GetEAnnotations(), o.EGetFromID(EMODEL_ELEMENT__EANNOTATIONS, true))
	assert.Equal(t, o.GetEAnnotations().(EObjectList).GetUnResolvedList(), o.EGetFromID(EMODEL_ELEMENT__EANNOTATIONS, false))
}

func TestEModelElementESetFromID(t *testing.T) {
	o := newEModelElementImpl()
	assert.Panics(t, func() { o.ESetFromID(-1, nil) })
	{
		// list with a value
		mockValue := NewMockEAnnotation(t)
		l := NewImmutableEList([]any{mockValue})
		mockValue.EXPECT().EInverseAdd(o, EANNOTATION__EMODEL_ELEMENT, mock.Anything).Return(nil).Once()

		// set list with new contents
		o.ESetFromID(EMODEL_ELEMENT__EANNOTATIONS, l)
		// checks
		assert.Equal(t, 1, o.GetEAnnotations().Size())
		assert.Equal(t, mockValue, o.GetEAnnotations().Get(0))
		mock.AssertExpectationsForObjects(t, mockValue)
	}

}

func TestEModelElementEIsSetFromID(t *testing.T) {
	o := newEModelElementImpl()
	assert.Panics(t, func() { o.EIsSetFromID(-1) })
	assert.False(t, o.EIsSetFromID(EMODEL_ELEMENT__EANNOTATIONS))
}

func TestEModelElementEUnsetFromID(t *testing.T) {
	o := newEModelElementImpl()
	assert.Panics(t, func() { o.EUnsetFromID(-1) })
	{
		o.EUnsetFromID(EMODEL_ELEMENT__EANNOTATIONS)
		v := o.EGetFromID(EMODEL_ELEMENT__EANNOTATIONS, false)
		assert.NotNil(t, v)
		l := v.(EList)
		assert.True(t, l.Empty())
	}
}

func TestEModelElementEInvokeFromID(t *testing.T) {
	o := newEModelElementImpl()
	assert.Panics(t, func() { o.EInvokeFromID(-1, nil) })
	assert.Panics(t, func() { o.EInvokeFromID(EMODEL_ELEMENT__GET_EANNOTATION_ESTRING, nil) })
}

func TestEModelElementEBasicInverseAdd(t *testing.T) {
	o := newEModelElementImpl()
	{
		mockObject := NewMockEObject(t)
		mockNotifications := NewMockENotificationChain(t)
		assert.Equal(t, mockNotifications, o.EBasicInverseAdd(mockObject, -1, mockNotifications))
	}
	{
		mockObject := NewMockEAnnotation(t)
		o.EBasicInverseAdd(mockObject, EMODEL_ELEMENT__EANNOTATIONS, nil)
		l := o.GetEAnnotations()
		assert.True(t, l.Contains(mockObject))
		mock.AssertExpectationsForObjects(t, mockObject)
	}

}

func TestEModelElementEBasicInverseRemove(t *testing.T) {
	o := newEModelElementImpl()
	{
		mockObject := NewMockEObject(t)
		mockNotifications := NewMockENotificationChain(t)
		assert.Equal(t, mockNotifications, o.EBasicInverseRemove(mockObject, -1, mockNotifications))
	}
	{
		// initialize list with a mock object
		mockObject := NewMockEAnnotation(t)
		mockObject.EXPECT().EInverseAdd(o, EANNOTATION__EMODEL_ELEMENT, mock.Anything).Return(nil).Once()

		l := o.GetEAnnotations()
		l.Add(mockObject)

		// basic inverse remove
		o.EBasicInverseRemove(mockObject, EMODEL_ELEMENT__EANNOTATIONS, nil)

		// check it was removed
		assert.False(t, l.Contains(mockObject))
		mock.AssertExpectationsForObjects(t, mockObject)
	}

}
