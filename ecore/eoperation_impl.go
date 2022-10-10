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

// eOperationImpl is the implementation of the model object 'EOperation'
type eOperationImpl struct {
	eTypedElementExt
	eExceptions EList
	eParameters EList
	operationID int
}
type eOperationInitializers interface {
	initEExceptions() EList
	initEParameters() EList
}

// newEOperationImpl is the constructor of a eOperationImpl
func newEOperationImpl() *eOperationImpl {
	eOperation := new(eOperationImpl)
	eOperation.SetInterfaces(eOperation)
	eOperation.Initialize()
	return eOperation
}

func (eOperation *eOperationImpl) Initialize() {
	eOperation.eTypedElementExt.Initialize()
	eOperation.operationID = -1

}

func (eOperation *eOperationImpl) asEOperation() EOperation {
	return eOperation.GetInterfaces().(EOperation)
}

func (eOperation *eOperationImpl) asInitializers() eOperationInitializers {
	return eOperation.GetInterfaces().(eOperationInitializers)
}

func (eOperation *eOperationImpl) EStaticClass() EClass {
	return GetPackage().GetEOperation()
}

func (eOperation *eOperationImpl) EStaticFeatureCount() int {
	return EOPERATION_FEATURE_COUNT
}

// IsOverrideOf default implementation
func (eOperation *eOperationImpl) IsOverrideOf(EOperation) bool {
	panic("IsOverrideOf not implemented")
}

// GetEContainingClass get the value of eContainingClass
func (eOperation *eOperationImpl) GetEContainingClass() EClass {
	if eOperation.EContainerFeatureID() == EOPERATION__ECONTAINING_CLASS {
		return eOperation.EContainer().(EClass)
	}
	return nil
}

// GetEExceptions get the value of eExceptions
func (eOperation *eOperationImpl) GetEExceptions() EList {
	if eOperation.eExceptions == nil {
		eOperation.eExceptions = eOperation.asInitializers().initEExceptions()
	}
	return eOperation.eExceptions
}

// UnsetEExceptions unset the value of eExceptions
func (eOperation *eOperationImpl) UnsetEExceptions() {
	if eOperation.eExceptions != nil {
		eOperation.eExceptions.Clear()
	}
}

// GetEParameters get the value of eParameters
func (eOperation *eOperationImpl) GetEParameters() EList {
	if eOperation.eParameters == nil {
		eOperation.eParameters = eOperation.asInitializers().initEParameters()
	}
	return eOperation.eParameters
}

// GetOperationID get the value of operationID
func (eOperation *eOperationImpl) GetOperationID() int {
	return eOperation.operationID
}

// SetOperationID set the value of operationID
func (eOperation *eOperationImpl) SetOperationID(newOperationID int) {
	oldOperationID := eOperation.operationID
	eOperation.operationID = newOperationID
	if eOperation.ENotificationRequired() {
		eOperation.ENotify(NewNotificationByFeatureID(eOperation.AsEObject(), SET, EOPERATION__OPERATION_ID, oldOperationID, newOperationID, NO_INDEX))
	}
}

func (eOperation *eOperationImpl) initEExceptions() EList {
	return NewBasicEObjectList(eOperation.AsEObjectInternal(), EOPERATION__EEXCEPTIONS, -1, false, false, false, true, true)
}

func (eOperation *eOperationImpl) initEParameters() EList {
	return NewBasicEObjectList(eOperation.AsEObjectInternal(), EOPERATION__EPARAMETERS, EPARAMETER__EOPERATION, true, true, true, false, false)
}

func (eOperation *eOperationImpl) EGetFromID(featureID int, resolve bool) any {
	switch featureID {
	case EOPERATION__ECONTAINING_CLASS:
		return eOperation.asEOperation().GetEContainingClass()
	case EOPERATION__EEXCEPTIONS:
		list := eOperation.asEOperation().GetEExceptions()
		if !resolve {
			if objects, _ := list.(EObjectList); objects != nil {
				return objects.GetUnResolvedList()
			}
		}
		return list
	case EOPERATION__EPARAMETERS:
		return eOperation.asEOperation().GetEParameters()
	case EOPERATION__OPERATION_ID:
		return eOperation.asEOperation().GetOperationID()
	default:
		return eOperation.eTypedElementExt.EGetFromID(featureID, resolve)
	}
}

func (eOperation *eOperationImpl) ESetFromID(featureID int, newValue any) {
	switch featureID {
	case EOPERATION__EEXCEPTIONS:
		list := eOperation.asEOperation().GetEExceptions()
		list.Clear()
		list.AddAll(newValue.(EList))
	case EOPERATION__EPARAMETERS:
		list := eOperation.asEOperation().GetEParameters()
		list.Clear()
		list.AddAll(newValue.(EList))
	case EOPERATION__OPERATION_ID:
		eOperation.asEOperation().SetOperationID(newValue.(int))
	default:
		eOperation.eTypedElementExt.ESetFromID(featureID, newValue)
	}
}

func (eOperation *eOperationImpl) EUnsetFromID(featureID int) {
	switch featureID {
	case EOPERATION__EEXCEPTIONS:
		eOperation.asEOperation().UnsetEExceptions()
	case EOPERATION__EPARAMETERS:
		eOperation.asEOperation().GetEParameters().Clear()
	case EOPERATION__OPERATION_ID:
		eOperation.asEOperation().SetOperationID(-1)
	default:
		eOperation.eTypedElementExt.EUnsetFromID(featureID)
	}
}

func (eOperation *eOperationImpl) EIsSetFromID(featureID int) bool {
	switch featureID {
	case EOPERATION__ECONTAINING_CLASS:
		return eOperation.asEOperation().GetEContainingClass() != nil
	case EOPERATION__EEXCEPTIONS:
		return eOperation.eExceptions != nil && eOperation.eExceptions.Size() != 0
	case EOPERATION__EPARAMETERS:
		return eOperation.eParameters != nil && eOperation.eParameters.Size() != 0
	case EOPERATION__OPERATION_ID:
		return eOperation.operationID != -1
	default:
		return eOperation.eTypedElementExt.EIsSetFromID(featureID)
	}
}

func (eOperation *eOperationImpl) EInvokeFromID(operationID int, arguments EList) any {
	switch operationID {
	case EOPERATION__IS_OVERRIDE_OF_EOPERATION:
		return eOperation.asEOperation().IsOverrideOf(arguments.Get(0).(EOperation))
	default:
		return eOperation.eTypedElementExt.EInvokeFromID(operationID, arguments)
	}
}

func (eOperation *eOperationImpl) EBasicInverseAdd(otherEnd EObject, featureID int, notifications ENotificationChain) ENotificationChain {
	switch featureID {
	case EOPERATION__ECONTAINING_CLASS:
		msgs := notifications
		if eOperation.EInternalContainer() != nil {
			msgs = eOperation.EBasicRemoveFromContainer(msgs)
		}
		return eOperation.EBasicSetContainer(otherEnd, EOPERATION__ECONTAINING_CLASS, msgs)
	case EOPERATION__EPARAMETERS:
		list := eOperation.GetEParameters().(ENotifyingList)
		return list.AddWithNotification(otherEnd, notifications)
	default:
		return eOperation.eTypedElementExt.EBasicInverseAdd(otherEnd, featureID, notifications)
	}
}

func (eOperation *eOperationImpl) EBasicInverseRemove(otherEnd EObject, featureID int, notifications ENotificationChain) ENotificationChain {
	switch featureID {
	case EOPERATION__ECONTAINING_CLASS:
		return eOperation.EBasicSetContainer(nil, EOPERATION__ECONTAINING_CLASS, notifications)
	case EOPERATION__EPARAMETERS:
		list := eOperation.GetEParameters().(ENotifyingList)
		return list.RemoveWithNotification(otherEnd, notifications)
	default:
		return eOperation.eTypedElementExt.EBasicInverseRemove(otherEnd, featureID, notifications)
	}
}
