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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestBasicEObjectListAccessors(t *testing.T) {
	{
		list := NewBasicEObjectList(nil, 1, -1, false, true, false, false, false)
		assert.Equal(t, nil, list.GetNotifier())
		assert.Equal(t, nil, list.GetFeature())
		assert.Equal(t, 1, list.GetFeatureID())
	}
	{
		mockOwner := NewMockEObjectInternal(t)
		list := NewBasicEObjectList(mockOwner, 1, -1, false, true, false, false, false)
		assert.Equal(t, mockOwner, list.GetNotifier())
		assert.Equal(t, 1, list.GetFeatureID())
		mockClass := NewMockEClass(t)
		mockFeature := NewMockEStructuralFeature(t)
		mockClass.EXPECT().GetEStructuralFeature(1).Return(mockFeature)
		mockOwner.EXPECT().EClass().Return(mockClass)
		assert.Equal(t, mockFeature, list.GetFeature())
		mock.AssertExpectationsForObjects(t, mockOwner, mockClass, mockFeature, mockClass)
	}
}

func TestBasicEObjectListInverseNoOpposite(t *testing.T) {
	mockOwner := NewMockEObjectInternal(t)
	mockOwner.EXPECT().EDeliver().Return(false)

	mockObject := NewMockEObjectInternal(t)
	list := NewBasicEObjectList(mockOwner, 1, -1, false, true, false, false, false)
	mockObject.EXPECT().EInverseAdd(mockOwner, -2, nil).Return(nil)

	assert.True(t, list.Add(mockObject))

	mockObject.EXPECT().EInverseRemove(mockOwner, -2, nil).Return(nil)
	assert.True(t, list.Remove(mockObject))

	mock.AssertExpectationsForObjects(t, mockObject, mockOwner)
}

func TestBasicEObjectListInverseOpposite(t *testing.T) {
	mockOwner := NewMockEObjectInternal(t)
	mockOwner.EXPECT().EDeliver().Return(false)

	mockObject := NewMockEObjectInternal(t)
	list := NewBasicEObjectList(mockOwner, 1, 2, false, true, true, false, false)

	mockObject.EXPECT().EInverseAdd(mockOwner, 2, nil).Return(nil)
	assert.True(t, list.Add(mockObject))

	mockObject.EXPECT().EInverseRemove(mockOwner, 2, nil).Return(nil)
	assert.True(t, list.Remove(mockObject))

	mock.AssertExpectationsForObjects(t, mockObject, mockOwner)
}

func TestBasicEObjectListContains(t *testing.T) {
	mockOwner := NewMockEObjectInternal(t)
	mockOwner.EXPECT().EDeliver().Return(false)

	// no proxy
	{
		list := NewBasicEObjectList(mockOwner, 1, 2, false, false, false, false, false)
		mockObject := NewMockEObjectInternal(t)
		list.Add(mockObject)
		assert.True(t, list.Contains(mockObject))
		assert.True(t, list.Contains(mockObject))

		mock.AssertExpectationsForObjects(t, mockObject, mockOwner)
	}

	// with proxy
	{
		list := NewBasicEObjectList(mockOwner, 1, 2, false, false, false, true, false)
		mockObject := NewMockEObjectInternal(t)
		list.Add(mockObject)
		assert.True(t, list.Contains(mockObject))

		mockResolved := NewMockEObjectInternal(t)
		mockOwner.EXPECT().EResolveProxy(mockObject).Return(mockResolved)
		mockObject.EXPECT().EIsProxy().Return(true)
		assert.True(t, list.Contains(mockResolved))

		mock.AssertExpectationsForObjects(t, mockObject, mockOwner)
	}
}

func TestBasicEObjectListGet(t *testing.T) {
	mockOwner := NewMockEObjectInternal(t)
	mockOwner.EXPECT().EDeliver().Return(false)

	list := NewBasicEObjectList(mockOwner, 1, 2, false, false, false, false, false)
	mockObject := NewMockEObjectInternal(t)
	list.Add(mockObject)
	assert.Equal(t, mockObject, list.Get(0))
	mock.AssertExpectationsForObjects(t, mockObject, mockOwner)

	// with proxy
	{
		list := NewBasicEObjectList(mockOwner, 1, 2, false, false, false, true, false)
		mockObject := NewMockEObjectInternal(t)
		list.Add(mockObject)

		mockResolved := NewMockEObjectInternal(t)
		mockOwner.EXPECT().EResolveProxy(mockObject).Return(mockResolved)
		mockObject.EXPECT().EIsProxy().Return(true)
		assert.Equal(t, mockResolved, list.Get(0))

		mock.AssertExpectationsForObjects(t, mockObject, mockOwner)
	}
}

func TestBasicEObjectListGetProxy(t *testing.T) {

	mockOwner := NewMockEObjectInternal(t)
	mockOwner.EXPECT().EDeliver().Return(false)

	list := NewBasicEObjectList(mockOwner, 1, 2, false, false, false, true, false)
	mockObject := NewMockEObjectInternal(t)
	list.Add(mockObject)

	mockResolved := NewMockEObjectInternal(t)
	mockOwner.EXPECT().EResolveProxy(mockObject).Return(mockResolved)
	mockObject.EXPECT().EIsProxy().Return(true)
	assert.Equal(t, mockResolved, list.Get(0))
	mock.AssertExpectationsForObjects(t, mockObject, mockOwner)
}

func TestBasicEObjectListGetProxyContainment(t *testing.T) {

	mockOwner := NewMockEObjectInternal(t)
	mockOwner.EXPECT().EDeliver().Return(false).Once()

	list := NewBasicEObjectList(mockOwner, 1, 2, true, false, false, true, false)
	mockObject := NewMockEObjectInternal(t)
	list.Add(mockObject)
	mock.AssertExpectationsForObjects(t, mockObject, mockOwner)

	mockAdapter := NewMockEAdapter(t)
	mockResolved := NewMockEObjectInternal(t)
	mockResolved.EXPECT().EInternalContainer().Return(nil).Once()
	mockObject.EXPECT().EIsProxy().Return(true)
	mockOwner.EXPECT().EDeliver().Return(true).Once()
	mockOwner.EXPECT().EAdapters().Return(NewImmutableEList([]any{mockAdapter}))
	mockOwner.EXPECT().EResolveProxy(mockObject).Return(mockResolved)
	mockOwner.EXPECT().ENotify(mock.MatchedBy(func(n ENotification) bool {
		return n.GetNotifier() == mockOwner &&
			n.GetFeatureID() == 1 &&
			n.GetNewValue() == mockResolved &&
			n.GetOldValue() == mockObject &&
			n.GetEventType() == RESOLVE &&
			n.GetPosition() == 0
	}))
	assert.Equal(t, mockResolved, list.Get(0))
	mock.AssertExpectationsForObjects(t, mockAdapter, mockObject, mockOwner, mockResolved)
}

func TestBasicEObjectListRemoveAll(t *testing.T) {
	mockOwner := NewMockEObjectInternal(t)
	mockOwner.EXPECT().EDeliver().Return(false)
	mockObject := NewMockEObjectInternal(t)
	mockOther := NewMockEObjectInternal(t)
	mockResolved := NewMockEObjectInternal(t)
	{
		list := NewBasicEObjectList(mockOwner, 1, 2, false, false, false, false, false)
		list.Add(mockObject)
		list.Add(mockOther)
		list.RemoveAll(NewImmutableEList([]any{mockObject, mockResolved}))
		assert.Equal(t, []any{mockOther}, list.ToArray())
	}
	{
		list := NewBasicEObjectList(mockOwner, 1, 2, false, false, false, true, false)
		mockObject.EXPECT().EIsProxy().Return(false).Once()
		list.Add(mockObject)
		list.Add(mockOther)

		mockOther.EXPECT().EIsProxy().Return(true).Once()
		mockOwner.EXPECT().EResolveProxy(mockOther).Return(mockResolved).Once()
		list.RemoveAll(NewImmutableEList([]any{mockObject, mockResolved}))
		assert.Equal(t, []any{}, list.ToArray())
	}

}

func TestBasicEObjectListUnResolved(t *testing.T) {
	mockOwner := NewMockEObjectInternal(t)
	// no proxy
	{
		list := NewBasicEObjectList(mockOwner, 1, 2, false, false, false, false, false)
		assert.Equal(t, list, list.GetUnResolvedList())
	}
	// with proxy
	{
		list := NewBasicEObjectList(mockOwner, 1, 2, false, false, false, true, false)
		unresolved := list.GetUnResolvedList()
		assert.NotEqual(t, list, unresolved)
		unresolvedAsEObjectList, _ := unresolved.(EObjectList)
		assert.NotNil(t, unresolvedAsEObjectList)
		assert.Equal(t, unresolved, unresolvedAsEObjectList.GetUnResolvedList())
		unresolvedAsENotifyingList, _ := unresolved.(ENotifyingList)
		assert.NotNil(t, unresolvedAsENotifyingList)
	}
}

func TestBasicEObjectListUnResolvedGet(t *testing.T) {
	mockOwner := NewMockEObjectInternal(t)
	mockOwner.EXPECT().EDeliver().Return(false)

	// add an object to unresolved
	list := NewBasicEObjectList(mockOwner, 1, 2, false, false, false, true, false)
	unresolved := list.GetUnResolvedList()
	mockObject := NewMockEObjectInternal(t)
	unresolved.Add(mockObject)
	// check that in unresolved it is the same
	assert.Equal(t, mockObject, unresolved.Get(0))

	// check that in original list , there is a resolution
	mockResolved := NewMockEObjectInternal(t)
	mockOwner.EXPECT().EResolveProxy(mockObject).Once().Return(mockResolved)
	mockObject.EXPECT().EIsProxy().Return(true)
	assert.Equal(t, mockResolved, list.Get(0))

	// check that now it is the resolved one in the unresolved list
	assert.Equal(t, mockResolved, unresolved.Get(0))
	assert.Panics(t, func() { unresolved.Get(1) })

	mock.AssertExpectationsForObjects(t, mockOwner, mockObject, mockResolved)
}

func TestBasicEObjectListUnResolvedContains(t *testing.T) {
	mockOwner := NewMockEObjectInternal(t)
	mockOwner.EXPECT().EDeliver().Return(false)

	// add an object to unresolved
	list := NewBasicEObjectList(mockOwner, 1, 2, false, false, false, true, false)
	unresolved := list.GetUnResolvedList()
	mockObject := NewMockEObjectInternal(t)
	unresolved.Add(mockObject)

	assert.True(t, unresolved.Contains(mockObject))

	// check that in original list , there is a resolution
	mockResolved := NewMockEObjectInternal(t)
	mockOwner.EXPECT().EResolveProxy(mockObject).Once().Return(mockResolved)
	mockObject.EXPECT().EIsProxy().Return(true)
	assert.True(t, !unresolved.Contains(mockResolved))
	assert.True(t, list.Contains(mockResolved))

	mock.AssertExpectationsForObjects(t, mockOwner, mockObject, mockResolved)
}

func TestBasicEObjectListUnResolvedSet(t *testing.T) {
	mockOwner := NewMockEObjectInternal(t)
	mockOwner.EXPECT().EDeliver().Return(false)

	// add an object to unresolved
	list := NewBasicEObjectList(mockOwner, 1, 2, false, false, false, true, false)
	unresolved := list.GetUnResolvedList()
	mockObject := NewMockEObjectInternal(t)
	unresolved.Add(mockObject)

	// set first index as another object & check that it has been replaced
	mockObject1 := NewMockEObjectInternal(t)
	unresolved.Set(0, mockObject1)
	assert.Equal(t, mockObject1, unresolved.Get(0))

	// check that invalid range is supported
	assert.Panics(t, func() { unresolved.Set(1, mockObject) })

	mock.AssertExpectationsForObjects(t, mockOwner, mockObject)
}

func TestBasicEObjectListUnResolvedSetAlreadyPresent(t *testing.T) {
	mockOwner := NewMockEObjectInternal(t)
	mockObject1 := NewMockEObjectInternal(t)
	mockObject2 := NewMockEObjectInternal(t)

	list := NewBasicEObjectList(mockOwner, 1, 2, false, false, false, true, false)
	unresolved := list.GetUnResolvedList()

	// add an object to unresolved
	mockOwner.EXPECT().EDeliver().Return(false).Once()
	assert.True(t, unresolved.AddAll(NewImmutableEList([]any{mockObject1, mockObject2})))
	mock.AssertExpectationsForObjects(t, mockOwner, mockObject1)

	assert.Panics(t, func() { unresolved.Set(1, mockObject1) })
	mock.AssertExpectationsForObjects(t, mockOwner, mockObject1)
}

func TestBasicEObjectListUnResolvedAdd(t *testing.T) {
	mockOwner := NewMockEObjectInternal(t)
	mockObject := NewMockEObjectInternal(t)

	list := NewBasicEObjectList(mockOwner, 1, 2, false, false, false, true, false)
	unresolved := list.GetUnResolvedList()

	// add an object to unresolved
	mockOwner.EXPECT().EDeliver().Return(false).Once()
	assert.True(t, unresolved.Add(mockObject))
	mock.AssertExpectationsForObjects(t, mockOwner, mockObject)

	// add an same object to unresolved
	assert.False(t, unresolved.Add(mockObject))
	mock.AssertExpectationsForObjects(t, mockOwner, mockObject)
}

func TestBasicEObjectListUnResolvedAddAll(t *testing.T) {
	mockOwner := NewMockEObjectInternal(t)
	mockObject1 := NewMockEObjectInternal(t)
	mockObject2 := NewMockEObjectInternal(t)

	list := NewBasicEObjectList(mockOwner, 1, 2, false, false, false, true, false)
	unresolved := list.GetUnResolvedList()

	// add an object to unresolved
	mockOwner.EXPECT().EDeliver().Return(false).Once()
	assert.True(t, unresolved.AddAll(NewImmutableEList([]any{mockObject1})))
	mock.AssertExpectationsForObjects(t, mockOwner, mockObject1)

	// add two object with one already in the list
	mockOwner.EXPECT().EDeliver().Return(false).Once()
	assert.True(t, unresolved.AddAll(NewImmutableEList([]any{mockObject1, mockObject2})))
	mock.AssertExpectationsForObjects(t, mockOwner, mockObject1, mockObject2)

	// add two object with already in the list
	assert.False(t, unresolved.AddAll(NewImmutableEList([]any{mockObject1, mockObject2})))
	mock.AssertExpectationsForObjects(t, mockOwner, mockObject1, mockObject2)
}

func TestBasicEObjectListUnResolvedInsert(t *testing.T) {
	mockOwner := NewMockEObjectInternal(t)
	mockObject := NewMockEObjectInternal(t)

	list := NewBasicEObjectList(mockOwner, 1, 2, false, false, false, true, false)
	unresolved := list.GetUnResolvedList()

	assert.Panics(t, func() {
		unresolved.Insert(1, mockObject)
	})

	mockOwner.EXPECT().EDeliver().Return(false).Once()
	assert.True(t, unresolved.Insert(0, mockObject))
	mock.AssertExpectationsForObjects(t, mockOwner, mockObject)

	assert.False(t, unresolved.Insert(0, mockObject))
	mock.AssertExpectationsForObjects(t, mockOwner, mockObject)
}

func TestBasicEObjectListUnResolvedInsertAll(t *testing.T) {
	mockOwner := NewMockEObjectInternal(t)
	mockObject1 := NewMockEObjectInternal(t)
	mockObject2 := NewMockEObjectInternal(t)

	list := NewBasicEObjectList(mockOwner, 1, 2, false, false, false, true, false)
	unresolved := list.GetUnResolvedList()

	assert.Panics(t, func() {
		unresolved.InsertAll(1, NewImmutableEList([]any{mockObject1}))
	})

	mockOwner.EXPECT().EDeliver().Return(false).Once()
	assert.True(t, unresolved.InsertAll(0, NewImmutableEList([]any{mockObject1})))
	mock.AssertExpectationsForObjects(t, mockOwner, mockObject1)

	mockOwner.EXPECT().EDeliver().Return(false).Once()
	assert.True(t, unresolved.InsertAll(0, NewImmutableEList([]any{mockObject1, mockObject2})))
	mock.AssertExpectationsForObjects(t, mockOwner, mockObject1, mockObject2)

	assert.False(t, unresolved.InsertAll(0, NewImmutableEList([]any{mockObject1, mockObject2})))
	mock.AssertExpectationsForObjects(t, mockOwner, mockObject1, mockObject2)
}

func TestBasicEObjectListUnResolvedMoveObject(t *testing.T) {
	mockOwner := NewMockEObjectInternal(t)
	mockObject1 := NewMockEObjectInternal(t)
	mockObject2 := NewMockEObjectInternal(t)

	list := NewBasicEObjectList(mockOwner, 1, 2, false, false, false, true, false)
	unresolved := list.GetUnResolvedList()

	assert.Panics(t, func() {
		unresolved.MoveObject(0, mockObject2)
	})

	mockOwner.EXPECT().EDeliver().Return(false).Once()
	unresolved.AddAll(NewImmutableEList([]any{mockObject1, mockObject2}))
	mock.AssertExpectationsForObjects(t, mockOwner, mockObject1, mockObject2)

	mockOwner.EXPECT().EDeliver().Return(false).Once()
	unresolved.MoveObject(0, mockObject2)
	assert.Equal(t, []any{mockObject2, mockObject1}, unresolved.ToArray())
	mock.AssertExpectationsForObjects(t, mockOwner, mockObject1, mockObject2)
}

func TestBasicEObjectListUnResolvedMoveIndex(t *testing.T) {

	mockOwner := NewMockEObjectInternal(t)
	mockObject1 := NewMockEObjectInternal(t)
	mockObject2 := NewMockEObjectInternal(t)

	list := NewBasicEObjectList(mockOwner, 1, 2, false, false, false, true, false)
	unresolved := list.GetUnResolvedList()

	mockOwner.EXPECT().EDeliver().Return(false).Once()
	unresolved.AddAll(NewImmutableEList([]any{mockObject1, mockObject2}))
	mock.AssertExpectationsForObjects(t, mockOwner, mockObject1, mockObject2)

	mockOwner.EXPECT().EDeliver().Return(false).Once()
	unresolved.Move(0, 1)
	assert.Equal(t, []any{mockObject2, mockObject1}, unresolved.ToArray())
	mock.AssertExpectationsForObjects(t, mockOwner, mockObject1, mockObject2)

}

func TestBasicEObjectListUnResolvedRemoveAt(t *testing.T) {
	mockOwner := NewMockEObjectInternal(t)
	mockObject1 := NewMockEObjectInternal(t)
	mockObject2 := NewMockEObjectInternal(t)

	list := NewBasicEObjectList(mockOwner, 1, 2, false, false, false, true, false)
	unresolved := list.GetUnResolvedList()

	mockOwner.EXPECT().EDeliver().Return(false).Once()
	unresolved.AddAll(NewImmutableEList([]any{mockObject1, mockObject2}))
	mock.AssertExpectationsForObjects(t, mockOwner, mockObject1, mockObject2)

	mockOwner.EXPECT().EDeliver().Return(false).Once()
	unresolved.RemoveAt(0)
	assert.Equal(t, []any{mockObject2}, unresolved.ToArray())
	mock.AssertExpectationsForObjects(t, mockOwner, mockObject1, mockObject2)
}

func TestBasicEObjectListUnResolvedRemoveRange(t *testing.T) {
	mockOwner := NewMockEObjectInternal(t)
	mockObject1 := NewMockEObjectInternal(t)
	mockObject2 := NewMockEObjectInternal(t)

	list := NewBasicEObjectList(mockOwner, 1, 2, false, false, false, true, false)
	unresolved := list.GetUnResolvedList()

	mockOwner.EXPECT().EDeliver().Return(false).Once()
	unresolved.AddAll(NewImmutableEList([]any{mockObject1, mockObject2}))
	mock.AssertExpectationsForObjects(t, mockOwner, mockObject1, mockObject2)

	mockOwner.EXPECT().EDeliver().Return(false).Once()
	unresolved.RemoveRange(0, 2)
	assert.Equal(t, []any{}, unresolved.ToArray())
	mock.AssertExpectationsForObjects(t, mockOwner, mockObject1, mockObject2)
}

func TestBasicEObjectListUnResolvedRemove(t *testing.T) {
	mockOwner := NewMockEObjectInternal(t)
	mockObject1 := NewMockEObjectInternal(t)
	mockObject2 := NewMockEObjectInternal(t)

	list := NewBasicEObjectList(mockOwner, 1, 2, false, false, false, true, false)
	unresolved := list.GetUnResolvedList()

	mockOwner.EXPECT().EDeliver().Return(false).Once()
	unresolved.AddAll(NewImmutableEList([]any{mockObject1}))
	mock.AssertExpectationsForObjects(t, mockOwner, mockObject1)

	mockOwner.EXPECT().EDeliver().Return(false).Once()
	assert.False(t, unresolved.Remove(mockObject2))
	assert.True(t, unresolved.Remove(mockObject1))
	assert.Equal(t, []any{}, unresolved.ToArray())
	mock.AssertExpectationsForObjects(t, mockOwner, mockObject1, mockObject2)
}

func TestBasicEObjectListUnResolvedRemoveAll(t *testing.T) {
	mockOwner := NewMockEObjectInternal(t)
	mockObject1 := NewMockEObjectInternal(t)
	mockObject2 := NewMockEObjectInternal(t)
	mockObject3 := NewMockEObjectInternal(t)

	list := NewBasicEObjectList(mockOwner, 1, 2, false, false, false, true, false)
	unresolved := list.GetUnResolvedList()

	mockOwner.EXPECT().EDeliver().Return(false).Once()
	unresolved.AddAll(NewImmutableEList([]any{mockObject1, mockObject2, mockObject3}))
	mock.AssertExpectationsForObjects(t, mockOwner, mockObject1)

	mockOwner.EXPECT().EDeliver().Return(false).Once()
	assert.True(t, unresolved.RemoveAll(NewImmutableEList([]any{mockObject1, mockObject2})))
	assert.Equal(t, []any{mockObject3}, unresolved.ToArray())
	mock.AssertExpectationsForObjects(t, mockOwner, mockObject1, mockObject2, mockObject3)
}

func TestBasicEObjectListUnResolvedClear(t *testing.T) {

	mockOwner := NewMockEObjectInternal(t)
	mockObject1 := NewMockEObjectInternal(t)
	mockObject2 := NewMockEObjectInternal(t)

	list := NewBasicEObjectList(mockOwner, 1, 2, false, false, false, true, false)
	unresolved := list.GetUnResolvedList()

	mockOwner.EXPECT().EDeliver().Return(false).Once()
	unresolved.AddAll(NewImmutableEList([]any{mockObject1, mockObject2}))
	mock.AssertExpectationsForObjects(t, mockOwner, mockObject1, mockObject2)

	mockOwner.EXPECT().EDeliver().Return(false).Once()
	unresolved.Clear()
	mock.AssertExpectationsForObjects(t, mockOwner, mockObject1, mockObject2)
}

func TestBasicEObjectListUnResolvedAccessors(t *testing.T) {
	mockOwner := NewMockEObjectInternal(t)
	mockClass := NewMockEClass(t)
	mockFeature := NewMockEStructuralFeature(t)
	list := NewBasicEObjectList(mockOwner, 1, 2, false, false, false, true, false)
	unresolved, _ := list.GetUnResolvedList().(ENotifyingList)
	assert.NotNil(t, unresolved)
	assert.Equal(t, mockOwner, unresolved.GetNotifier())
	assert.Equal(t, 1, unresolved.GetFeatureID())

	mockOwner.EXPECT().EClass().Return(mockClass).Once()
	mockClass.EXPECT().GetEStructuralFeature(1).Return(mockFeature).Once()
	assert.Equal(t, mockFeature, unresolved.GetFeature())
	mock.AssertExpectationsForObjects(t, mockOwner, mockClass, mockFeature)

}

func TestBasicEObjectListUnResolvedAddWithNotification(t *testing.T) {
	mockOwner := NewMockEObjectInternal(t)
	mockObject1 := NewMockEObjectInternal(t)

	list := NewBasicEObjectList(mockOwner, 1, 2, false, false, false, true, false)
	unresolved, _ := list.GetUnResolvedList().(ENotifyingList)

	mockOwner.EXPECT().EDeliver().Return(false).Once()
	unresolved.AddWithNotification(mockObject1, nil)
	assert.Equal(t, []any{mockObject1}, unresolved.ToArray())
	mock.AssertExpectationsForObjects(t, mockOwner, mockObject1)
}

func TestBasicEObjectListUnResolvedRemoveWithNotification(t *testing.T) {
	mockOwner := NewMockEObjectInternal(t)
	mockObject1 := NewMockEObjectInternal(t)
	mockObject2 := NewMockEObjectInternal(t)

	list := NewBasicEObjectList(mockOwner, 1, 2, false, false, false, true, false)
	unresolved, _ := list.GetUnResolvedList().(ENotifyingList)

	mockOwner.EXPECT().EDeliver().Return(false).Once()
	unresolved.AddAll(NewImmutableEList([]any{mockObject1, mockObject2}))
	assert.Equal(t, []any{mockObject1, mockObject2}, unresolved.ToArray())
	mock.AssertExpectationsForObjects(t, mockOwner, mockObject1, mockObject2)

	mockOwner.EXPECT().EDeliver().Return(false).Once()
	unresolved.RemoveWithNotification(mockObject2, nil)
	assert.Equal(t, []any{mockObject1}, unresolved.ToArray())
	mock.AssertExpectationsForObjects(t, mockOwner, mockObject1, mockObject2)

	mockNotificationChain := NewMockENotificationChain(t)
	assert.Equal(t, mockNotificationChain, unresolved.RemoveWithNotification(mockObject2, mockNotificationChain))
}

func TestBasicEObjectListUnResolvedSetWithNotification(t *testing.T) {
	mockOwner := NewMockEObjectInternal(t)
	mockObject1 := NewMockEObjectInternal(t)
	mockObject2 := NewMockEObjectInternal(t)

	list := NewBasicEObjectList(mockOwner, 1, 2, false, false, false, true, false)
	unresolved, _ := list.GetUnResolvedList().(ENotifyingList)

	mockOwner.EXPECT().EDeliver().Return(false).Once()
	unresolved.Add(mockObject1)
	assert.Equal(t, []any{mockObject1}, unresolved.ToArray())
	mock.AssertExpectationsForObjects(t, mockOwner, mockObject1)

	mockOwner.EXPECT().EDeliver().Return(false).Once()
	unresolved.SetWithNotification(0, mockObject2, nil)
	assert.Equal(t, []any{mockObject2}, unresolved.ToArray())
	mock.AssertExpectationsForObjects(t, mockOwner, mockObject1, mockObject2)
}
