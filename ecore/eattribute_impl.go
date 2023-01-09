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

// EAttributeImpl is the implementation of the model object 'EAttribute'
type EAttributeImpl struct {
	EStructuralFeatureExt
	isID bool
}

type eAttributeBasics interface {
	basicGetEAttributeType() EDataType
}

// newEAttributeImpl is the constructor of a EAttributeImpl
func newEAttributeImpl() *EAttributeImpl {
	eAttribute := new(EAttributeImpl)
	eAttribute.SetInterfaces(eAttribute)
	eAttribute.Initialize()
	return eAttribute
}

func (eAttribute *EAttributeImpl) Initialize() {
	eAttribute.EStructuralFeatureExt.Initialize()
	eAttribute.isID = false

}

func (eAttribute *EAttributeImpl) asEAttribute() EAttribute {
	return eAttribute.GetInterfaces().(EAttribute)
}

func (eAttribute *EAttributeImpl) asBasics() eAttributeBasics {
	return eAttribute.GetInterfaces().(eAttributeBasics)
}

func (eAttribute *EAttributeImpl) EStaticClass() EClass {
	return GetPackage().GetEAttribute()
}

func (eAttribute *EAttributeImpl) EStaticFeatureCount() int {
	return EATTRIBUTE_FEATURE_COUNT
}

// GetEAttributeType get the value of eAttributeType
func (eAttribute *EAttributeImpl) GetEAttributeType() EDataType {
	panic("GetEAttributeType not implemented")
}

func (eAttribute *EAttributeImpl) basicGetEAttributeType() EDataType {
	panic("GetEAttributeType not implemented")
}

// IsID get the value of isID
func (eAttribute *EAttributeImpl) IsID() bool {
	return eAttribute.isID
}

// SetID set the value of isID
func (eAttribute *EAttributeImpl) SetID(newIsID bool) {
	oldIsID := eAttribute.isID
	eAttribute.isID = newIsID
	if eAttribute.ENotificationRequired() {
		eAttribute.ENotify(NewNotificationByFeatureID(eAttribute.AsEObject(), SET, EATTRIBUTE__ID, oldIsID, newIsID, NO_INDEX))
	}
}

func (eAttribute *EAttributeImpl) EGetFromID(featureID int, resolve bool) any {
	switch featureID {
	case EATTRIBUTE__EATTRIBUTE_TYPE:
		if resolve {
			return eAttribute.asEAttribute().GetEAttributeType()
		}
		return eAttribute.asBasics().basicGetEAttributeType()
	case EATTRIBUTE__ID:
		return eAttribute.asEAttribute().IsID()
	default:
		return eAttribute.EStructuralFeatureExt.EGetFromID(featureID, resolve)
	}
}

func (eAttribute *EAttributeImpl) ESetFromID(featureID int, newValue any) {
	switch featureID {
	case EATTRIBUTE__ID:
		eAttribute.asEAttribute().SetID(newValue.(bool))
	default:
		eAttribute.EStructuralFeatureExt.ESetFromID(featureID, newValue)
	}
}

func (eAttribute *EAttributeImpl) EUnsetFromID(featureID int) {
	switch featureID {
	case EATTRIBUTE__ID:
		eAttribute.asEAttribute().SetID(false)
	default:
		eAttribute.EStructuralFeatureExt.EUnsetFromID(featureID)
	}
}

func (eAttribute *EAttributeImpl) EIsSetFromID(featureID int) bool {
	switch featureID {
	case EATTRIBUTE__EATTRIBUTE_TYPE:
		return eAttribute.asEAttribute().GetEAttributeType() != nil
	case EATTRIBUTE__ID:
		return eAttribute.isID != false
	default:
		return eAttribute.EStructuralFeatureExt.EIsSetFromID(featureID)
	}
}
