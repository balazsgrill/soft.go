// *****************************************************************************
// Copyright(c) 2021 MASA Group
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// *****************************************************************************

package ecore

type BasicEObjectMap struct {
	BasicEMap
	entryClass EClass
}

type basicEObjectMapList struct {
	basicEObjectList
	m *BasicEObjectMap
}

func newBasicEObjectMapList(m *BasicEObjectMap, owner EObjectInternal, featureID, inverseFeatureID int, unset bool) *basicEObjectMapList {
	l := new(basicEObjectMapList)
	l.interfaces = l
	l.data = []interface{}{}
	l.isUnique = true
	l.owner = owner
	l.featureID = featureID
	l.inverseFeatureID = inverseFeatureID
	l.containment = true
	l.inverse = inverseFeatureID != -1
	l.opposite = l.inverse
	l.proxies = false
	l.unset = unset
	l.m = m
	return l
}

func (ml *basicEObjectMapList) didAdd(index int, elem interface{}) {
	entry := elem.(EMapEntry)
	ml.m.doAdd(entry)
}

func (ml *basicEObjectMapList) didSet(index int, newElem interface{}, oldElem interface{}) {
	newEntry := newElem.(EMapEntry)
	oldEntry := oldElem.(EMapEntry)
	ml.m.doRemove(oldEntry)
	ml.m.doAdd(newEntry)
}

func (ml *basicEObjectMapList) didRemove(index int, oldElem interface{}) {
	oldEntry := oldElem.(EMapEntry)
	ml.m.doRemove(oldEntry)
}

func (ml *basicEObjectMapList) didClear(oldObjects []interface{}) {
	ml.m.doClear()
}

func NewBasicEObjectMap(entryClass EClass, owner EObjectInternal, featureID int, inverseFeatureID int, unset bool) *BasicEObjectMap {
	basicEObjectMap := &BasicEObjectMap{}
	basicEObjectMap.EList = newBasicEObjectMapList(basicEObjectMap, owner, featureID, inverseFeatureID, unset)
	basicEObjectMap.interfaces = basicEObjectMap
	basicEObjectMap.mapData = make(map[interface{}]interface{})
	basicEObjectMap.entryClass = entryClass
	return basicEObjectMap
}

func (m *BasicEObjectMap) newEntry(key interface{}, value interface{}) EMapEntry {
	eFactory := m.entryClass.GetEPackage().GetEFactoryInstance()
	eEntry := eFactory.Create(m.entryClass).(EMapEntry)
	eEntry.SetKey(key)
	eEntry.SetValue(value)
	return eEntry
}
