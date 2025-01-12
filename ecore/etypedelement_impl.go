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

// ETypedElementImpl is the implementation of the model object 'ETypedElement'
type ETypedElementImpl struct {
	ENamedElementImpl
	eType      EClassifier
	isOrdered  bool
	isUnique   bool
	lowerBound int
	upperBound int
}

type eTypedElementBasics interface {
	basicGetEType() EClassifier
}

// newETypedElementImpl is the constructor of a ETypedElementImpl
func newETypedElementImpl() *ETypedElementImpl {
	eTypedElement := new(ETypedElementImpl)
	eTypedElement.SetInterfaces(eTypedElement)
	eTypedElement.Initialize()
	return eTypedElement
}

func (eTypedElement *ETypedElementImpl) Initialize() {
	eTypedElement.ENamedElementImpl.Initialize()
	eTypedElement.isOrdered = true
	eTypedElement.isUnique = true
	eTypedElement.lowerBound = 0
	eTypedElement.upperBound = 1

}

func (eTypedElement *ETypedElementImpl) asETypedElement() ETypedElement {
	return eTypedElement.GetInterfaces().(ETypedElement)
}

func (eTypedElement *ETypedElementImpl) asBasics() eTypedElementBasics {
	return eTypedElement.GetInterfaces().(eTypedElementBasics)
}

func (eTypedElement *ETypedElementImpl) EStaticClass() EClass {
	return GetPackage().GetETypedElement()
}

func (eTypedElement *ETypedElementImpl) EStaticFeatureCount() int {
	return ETYPED_ELEMENT_FEATURE_COUNT
}

// GetEType get the value of eType
func (eTypedElement *ETypedElementImpl) GetEType() EClassifier {
	if eTypedElement.eType != nil && eTypedElement.eType.EIsProxy() {
		oldEType := eTypedElement.eType
		newEType := eTypedElement.EResolveProxy(oldEType).(EClassifier)
		eTypedElement.eType = newEType
		if newEType != oldEType {
			if eTypedElement.ENotificationRequired() {
				eTypedElement.ENotify(NewNotificationByFeatureID(eTypedElement, RESOLVE, ETYPED_ELEMENT__ETYPE, oldEType, newEType, NO_INDEX))
			}
		}
	}
	return eTypedElement.eType
}

func (eTypedElement *ETypedElementImpl) basicGetEType() EClassifier {
	return eTypedElement.eType
}

// SetEType set the value of eType
func (eTypedElement *ETypedElementImpl) SetEType(newEType EClassifier) {
	oldEType := eTypedElement.eType
	eTypedElement.eType = newEType
	if eTypedElement.ENotificationRequired() {
		eTypedElement.ENotify(NewNotificationByFeatureID(eTypedElement.AsEObject(), SET, ETYPED_ELEMENT__ETYPE, oldEType, newEType, NO_INDEX))
	}
}

// UnsetEType unset the value of eType
func (eTypedElement *ETypedElementImpl) UnsetEType() {
	oldEType := eTypedElement.eType
	eTypedElement.eType = nil
	if eTypedElement.ENotificationRequired() {
		eTypedElement.ENotify(NewNotificationByFeatureID(eTypedElement.AsEObject(), UNSET, ETYPED_ELEMENT__ETYPE, oldEType, nil, NO_INDEX))
	}
}

// IsMany get the value of isMany
func (eTypedElement *ETypedElementImpl) IsMany() bool {
	panic("IsMany not implemented")
}

// IsOrdered get the value of isOrdered
func (eTypedElement *ETypedElementImpl) IsOrdered() bool {
	return eTypedElement.isOrdered
}

// SetOrdered set the value of isOrdered
func (eTypedElement *ETypedElementImpl) SetOrdered(newIsOrdered bool) {
	oldIsOrdered := eTypedElement.isOrdered
	eTypedElement.isOrdered = newIsOrdered
	if eTypedElement.ENotificationRequired() {
		eTypedElement.ENotify(NewNotificationByFeatureID(eTypedElement.AsEObject(), SET, ETYPED_ELEMENT__ORDERED, oldIsOrdered, newIsOrdered, NO_INDEX))
	}
}

// IsRequired get the value of isRequired
func (eTypedElement *ETypedElementImpl) IsRequired() bool {
	panic("IsRequired not implemented")
}

// IsUnique get the value of isUnique
func (eTypedElement *ETypedElementImpl) IsUnique() bool {
	return eTypedElement.isUnique
}

// SetUnique set the value of isUnique
func (eTypedElement *ETypedElementImpl) SetUnique(newIsUnique bool) {
	oldIsUnique := eTypedElement.isUnique
	eTypedElement.isUnique = newIsUnique
	if eTypedElement.ENotificationRequired() {
		eTypedElement.ENotify(NewNotificationByFeatureID(eTypedElement.AsEObject(), SET, ETYPED_ELEMENT__UNIQUE, oldIsUnique, newIsUnique, NO_INDEX))
	}
}

// GetLowerBound get the value of lowerBound
func (eTypedElement *ETypedElementImpl) GetLowerBound() int {
	return eTypedElement.lowerBound
}

// SetLowerBound set the value of lowerBound
func (eTypedElement *ETypedElementImpl) SetLowerBound(newLowerBound int) {
	oldLowerBound := eTypedElement.lowerBound
	eTypedElement.lowerBound = newLowerBound
	if eTypedElement.ENotificationRequired() {
		eTypedElement.ENotify(NewNotificationByFeatureID(eTypedElement.AsEObject(), SET, ETYPED_ELEMENT__LOWER_BOUND, oldLowerBound, newLowerBound, NO_INDEX))
	}
}

// GetUpperBound get the value of upperBound
func (eTypedElement *ETypedElementImpl) GetUpperBound() int {
	return eTypedElement.upperBound
}

// SetUpperBound set the value of upperBound
func (eTypedElement *ETypedElementImpl) SetUpperBound(newUpperBound int) {
	oldUpperBound := eTypedElement.upperBound
	eTypedElement.upperBound = newUpperBound
	if eTypedElement.ENotificationRequired() {
		eTypedElement.ENotify(NewNotificationByFeatureID(eTypedElement.AsEObject(), SET, ETYPED_ELEMENT__UPPER_BOUND, oldUpperBound, newUpperBound, NO_INDEX))
	}
}

func (eTypedElement *ETypedElementImpl) EGetFromID(featureID int, resolve bool) any {
	switch featureID {
	case ETYPED_ELEMENT__ETYPE:
		if resolve {
			return eTypedElement.asETypedElement().GetEType()
		}
		return eTypedElement.asBasics().basicGetEType()
	case ETYPED_ELEMENT__LOWER_BOUND:
		return eTypedElement.asETypedElement().GetLowerBound()
	case ETYPED_ELEMENT__MANY:
		return eTypedElement.asETypedElement().IsMany()
	case ETYPED_ELEMENT__ORDERED:
		return eTypedElement.asETypedElement().IsOrdered()
	case ETYPED_ELEMENT__REQUIRED:
		return eTypedElement.asETypedElement().IsRequired()
	case ETYPED_ELEMENT__UNIQUE:
		return eTypedElement.asETypedElement().IsUnique()
	case ETYPED_ELEMENT__UPPER_BOUND:
		return eTypedElement.asETypedElement().GetUpperBound()
	default:
		return eTypedElement.ENamedElementImpl.EGetFromID(featureID, resolve)
	}
}

func (eTypedElement *ETypedElementImpl) ESetFromID(featureID int, newValue any) {
	switch featureID {
	case ETYPED_ELEMENT__ETYPE:
		newValueOrNil, _ := newValue.(EClassifier)
		eTypedElement.asETypedElement().SetEType(newValueOrNil)
	case ETYPED_ELEMENT__LOWER_BOUND:
		eTypedElement.asETypedElement().SetLowerBound(newValue.(int))
	case ETYPED_ELEMENT__ORDERED:
		eTypedElement.asETypedElement().SetOrdered(newValue.(bool))
	case ETYPED_ELEMENT__UNIQUE:
		eTypedElement.asETypedElement().SetUnique(newValue.(bool))
	case ETYPED_ELEMENT__UPPER_BOUND:
		eTypedElement.asETypedElement().SetUpperBound(newValue.(int))
	default:
		eTypedElement.ENamedElementImpl.ESetFromID(featureID, newValue)
	}
}

func (eTypedElement *ETypedElementImpl) EUnsetFromID(featureID int) {
	switch featureID {
	case ETYPED_ELEMENT__ETYPE:
		eTypedElement.asETypedElement().UnsetEType()
	case ETYPED_ELEMENT__LOWER_BOUND:
		eTypedElement.asETypedElement().SetLowerBound(0)
	case ETYPED_ELEMENT__ORDERED:
		eTypedElement.asETypedElement().SetOrdered(true)
	case ETYPED_ELEMENT__UNIQUE:
		eTypedElement.asETypedElement().SetUnique(true)
	case ETYPED_ELEMENT__UPPER_BOUND:
		eTypedElement.asETypedElement().SetUpperBound(1)
	default:
		eTypedElement.ENamedElementImpl.EUnsetFromID(featureID)
	}
}

func (eTypedElement *ETypedElementImpl) EIsSetFromID(featureID int) bool {
	switch featureID {
	case ETYPED_ELEMENT__ETYPE:
		return eTypedElement.eType != nil
	case ETYPED_ELEMENT__LOWER_BOUND:
		return eTypedElement.lowerBound != 0
	case ETYPED_ELEMENT__MANY:
		return eTypedElement.asETypedElement().IsMany() != false
	case ETYPED_ELEMENT__ORDERED:
		return eTypedElement.isOrdered != true
	case ETYPED_ELEMENT__REQUIRED:
		return eTypedElement.asETypedElement().IsRequired() != false
	case ETYPED_ELEMENT__UNIQUE:
		return eTypedElement.isUnique != true
	case ETYPED_ELEMENT__UPPER_BOUND:
		return eTypedElement.upperBound != 1
	default:
		return eTypedElement.ENamedElementImpl.EIsSetFromID(featureID)
	}
}
