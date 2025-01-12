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

// EStringToStringMapEntryImpl is the implementation of the model object 'EStringToStringMapEntry'
type EStringToStringMapEntryImpl struct {
	CompactEObjectContainer
	key   string
	value string
}

// newEStringToStringMapEntryImpl is the constructor of a EStringToStringMapEntryImpl
func newEStringToStringMapEntryImpl() *EStringToStringMapEntryImpl {
	eStringToStringMapEntry := new(EStringToStringMapEntryImpl)
	eStringToStringMapEntry.SetInterfaces(eStringToStringMapEntry)
	eStringToStringMapEntry.Initialize()
	return eStringToStringMapEntry
}

func (eStringToStringMapEntry *EStringToStringMapEntryImpl) Initialize() {
	eStringToStringMapEntry.CompactEObjectContainer.Initialize()
	eStringToStringMapEntry.key = ""
	eStringToStringMapEntry.value = ""

}

func (eStringToStringMapEntry *EStringToStringMapEntryImpl) asEStringToStringMapEntry() EStringToStringMapEntry {
	return eStringToStringMapEntry.GetInterfaces().(EStringToStringMapEntry)
}

func (eStringToStringMapEntry *EStringToStringMapEntryImpl) EStaticClass() EClass {
	return GetPackage().GetEStringToStringMapEntry()
}

func (eStringToStringMapEntry *EStringToStringMapEntryImpl) EStaticFeatureCount() int {
	return ESTRING_TO_STRING_MAP_ENTRY_FEATURE_COUNT
}
func (eStringToStringMapEntry *EStringToStringMapEntryImpl) GetKey() any {
	return eStringToStringMapEntry.GetTypedKey()
}

func (eStringToStringMapEntry *EStringToStringMapEntryImpl) SetKey(key any) {
	eStringToStringMapEntry.SetTypedKey(key.(string))
}

func (eStringToStringMapEntry *EStringToStringMapEntryImpl) GetValue() any {
	return eStringToStringMapEntry.GetTypedValue()
}

func (eStringToStringMapEntry *EStringToStringMapEntryImpl) SetValue(value any) {
	eStringToStringMapEntry.SetTypedValue(value.(string))
}

// GetTypedKey get the value of key
func (eStringToStringMapEntry *EStringToStringMapEntryImpl) GetTypedKey() string {
	return eStringToStringMapEntry.key
}

// SetTypedKey set the value of key
func (eStringToStringMapEntry *EStringToStringMapEntryImpl) SetTypedKey(newKey string) {
	oldKey := eStringToStringMapEntry.key
	eStringToStringMapEntry.key = newKey
	if eStringToStringMapEntry.ENotificationRequired() {
		eStringToStringMapEntry.ENotify(NewNotificationByFeatureID(eStringToStringMapEntry.AsEObject(), SET, ESTRING_TO_STRING_MAP_ENTRY__KEY, oldKey, newKey, NO_INDEX))
	}
}

// GetTypedValue get the value of value
func (eStringToStringMapEntry *EStringToStringMapEntryImpl) GetTypedValue() string {
	return eStringToStringMapEntry.value
}

// SetTypedValue set the value of value
func (eStringToStringMapEntry *EStringToStringMapEntryImpl) SetTypedValue(newValue string) {
	oldValue := eStringToStringMapEntry.value
	eStringToStringMapEntry.value = newValue
	if eStringToStringMapEntry.ENotificationRequired() {
		eStringToStringMapEntry.ENotify(NewNotificationByFeatureID(eStringToStringMapEntry.AsEObject(), SET, ESTRING_TO_STRING_MAP_ENTRY__VALUE, oldValue, newValue, NO_INDEX))
	}
}

func (eStringToStringMapEntry *EStringToStringMapEntryImpl) EGetFromID(featureID int, resolve bool) any {
	switch featureID {
	case ESTRING_TO_STRING_MAP_ENTRY__KEY:
		return eStringToStringMapEntry.asEStringToStringMapEntry().GetTypedKey()
	case ESTRING_TO_STRING_MAP_ENTRY__VALUE:
		return eStringToStringMapEntry.asEStringToStringMapEntry().GetTypedValue()
	default:
		return eStringToStringMapEntry.CompactEObjectContainer.EGetFromID(featureID, resolve)
	}
}

func (eStringToStringMapEntry *EStringToStringMapEntryImpl) ESetFromID(featureID int, newValue any) {
	switch featureID {
	case ESTRING_TO_STRING_MAP_ENTRY__KEY:
		eStringToStringMapEntry.asEStringToStringMapEntry().SetTypedKey(newValue.(string))
	case ESTRING_TO_STRING_MAP_ENTRY__VALUE:
		eStringToStringMapEntry.asEStringToStringMapEntry().SetTypedValue(newValue.(string))
	default:
		eStringToStringMapEntry.CompactEObjectContainer.ESetFromID(featureID, newValue)
	}
}

func (eStringToStringMapEntry *EStringToStringMapEntryImpl) EUnsetFromID(featureID int) {
	switch featureID {
	case ESTRING_TO_STRING_MAP_ENTRY__KEY:
		eStringToStringMapEntry.asEStringToStringMapEntry().SetTypedKey("")
	case ESTRING_TO_STRING_MAP_ENTRY__VALUE:
		eStringToStringMapEntry.asEStringToStringMapEntry().SetTypedValue("")
	default:
		eStringToStringMapEntry.CompactEObjectContainer.EUnsetFromID(featureID)
	}
}

func (eStringToStringMapEntry *EStringToStringMapEntryImpl) EIsSetFromID(featureID int) bool {
	switch featureID {
	case ESTRING_TO_STRING_MAP_ENTRY__KEY:
		return eStringToStringMapEntry.key != ""
	case ESTRING_TO_STRING_MAP_ENTRY__VALUE:
		return eStringToStringMapEntry.value != ""
	default:
		return eStringToStringMapEntry.CompactEObjectContainer.EIsSetFromID(featureID)
	}
}
