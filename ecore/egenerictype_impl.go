// *****************************************************************************
//
// This file is part of a MASA library or program.
// Refer to the included end-user license agreement for restrictions.
//
// Copyright (c) 2019 MASA Group
//
// *****************************************************************************

// *****************************************************************************
//
// Warning: This file was generated by soft.generator.go Generator
//
// *****************************************************************************

package ecore

// eGenericTypeImpl is the implementation of the model object 'EGenericType'
type eGenericTypeImpl struct {
	*EObjectImpl
	eClassifier    EClassifier
	eLowerBound    EGenericType
	eRawType       EClassifier
	eTypeArguments EList
	eTypeParameter ETypeParameter
	eUpperBound    EGenericType
}

// newEGenericTypeImpl is the constructor of a eGenericTypeImpl
func newEGenericTypeImpl() *eGenericTypeImpl {
	eGenericType := new(eGenericTypeImpl)
	eGenericType.EObjectImpl = NewEObjectImpl()
	eGenericType.SetInterfaces(eGenericType)

	return eGenericType
}

type eGenericTypeImplInitializers interface {
	initETypeArguments() EList
}

func (eGenericType *eGenericTypeImpl) getInitializers() eGenericTypeImplInitializers {
	return eGenericType.GetEObject().(eGenericTypeImplInitializers)
}

func (eGenericType *eGenericTypeImpl) EStaticClass() EClass {
	return GetPackage().GetEGenericType()
}

// IsInstance default implementation
func (eGenericType *eGenericTypeImpl) IsInstance(interface{}) bool {
	panic("IsInstance not implemented")
}

// GetEClassifier get the value of eClassifier
func (eGenericType *eGenericTypeImpl) GetEClassifier() EClassifier {
	return eGenericType.eClassifier

}

func (eGenericType *eGenericTypeImpl) basicGetEClassifier() EClassifier {
	return eGenericType.eClassifier

}

// SetEClassifier set the value of eClassifier
func (eGenericType *eGenericTypeImpl) SetEClassifier(newEClassifier EClassifier) {
	oldEClassifier := eGenericType.eClassifier
	eGenericType.eClassifier = newEClassifier
	if eGenericType.ENotificationRequired() {
		eGenericType.ENotify(NewNotificationByFeatureID(eGenericType.GetEObject(), SET, EGENERIC_TYPE__ECLASSIFIER, oldEClassifier, newEClassifier, NO_INDEX))
	}
}

// GetELowerBound get the value of eLowerBound
func (eGenericType *eGenericTypeImpl) GetELowerBound() EGenericType {
	return eGenericType.eLowerBound

}

// SetELowerBound set the value of eLowerBound
func (eGenericType *eGenericTypeImpl) SetELowerBound(newELowerBound EGenericType) {
	if newELowerBound != eGenericType.eLowerBound {
		var notifications ENotificationChain
		if eGenericType.eLowerBound != nil {
			notifications = eGenericType.eLowerBound.(EObjectInternal).EInverseRemove(eGenericType, EOPPOSITE_FEATURE_BASE-EGENERIC_TYPE__ELOWER_BOUND, notifications)
		}
		if newELowerBound != nil {
			notifications = newELowerBound.(EObjectInternal).EInverseAdd(eGenericType, EOPPOSITE_FEATURE_BASE-EGENERIC_TYPE__ELOWER_BOUND, notifications)
		}
		notifications = eGenericType.basicSetELowerBound(newELowerBound, notifications)
		if notifications != nil {
			notifications.Dispatch()
		}
	}
}

func (eGenericType *eGenericTypeImpl) basicSetELowerBound(newELowerBound EGenericType, msgs ENotificationChain) ENotificationChain {
	oldELowerBound := eGenericType.eLowerBound
	eGenericType.eLowerBound = newELowerBound
	notifications := msgs
	if eGenericType.ENotificationRequired() {
		notification := NewNotificationByFeatureID(eGenericType.GetEObject(), SET, EGENERIC_TYPE__ELOWER_BOUND, oldELowerBound, newELowerBound, NO_INDEX)
		if notifications != nil {
			notifications.Add(notification)
		} else {
			notifications = notification
		}
	}
	return notifications
}

// GetERawType get the value of eRawType
func (eGenericType *eGenericTypeImpl) GetERawType() EClassifier {
	return eGenericType.eRawType

}

func (eGenericType *eGenericTypeImpl) basicGetERawType() EClassifier {
	return eGenericType.eRawType

}

// GetETypeArguments get the value of eTypeArguments
func (eGenericType *eGenericTypeImpl) GetETypeArguments() EList {
	if eGenericType.eTypeArguments == nil {
		eGenericType.eTypeArguments = eGenericType.getInitializers().initETypeArguments()
	}
	return eGenericType.eTypeArguments

}

// GetETypeParameter get the value of eTypeParameter
func (eGenericType *eGenericTypeImpl) GetETypeParameter() ETypeParameter {
	return eGenericType.eTypeParameter

}

// SetETypeParameter set the value of eTypeParameter
func (eGenericType *eGenericTypeImpl) SetETypeParameter(newETypeParameter ETypeParameter) {
	oldETypeParameter := eGenericType.eTypeParameter
	eGenericType.eTypeParameter = newETypeParameter
	if eGenericType.ENotificationRequired() {
		eGenericType.ENotify(NewNotificationByFeatureID(eGenericType.GetEObject(), SET, EGENERIC_TYPE__ETYPE_PARAMETER, oldETypeParameter, newETypeParameter, NO_INDEX))
	}
}

// GetEUpperBound get the value of eUpperBound
func (eGenericType *eGenericTypeImpl) GetEUpperBound() EGenericType {
	return eGenericType.eUpperBound

}

// SetEUpperBound set the value of eUpperBound
func (eGenericType *eGenericTypeImpl) SetEUpperBound(newEUpperBound EGenericType) {
	if newEUpperBound != eGenericType.eUpperBound {
		var notifications ENotificationChain
		if eGenericType.eUpperBound != nil {
			notifications = eGenericType.eUpperBound.(EObjectInternal).EInverseRemove(eGenericType, EOPPOSITE_FEATURE_BASE-EGENERIC_TYPE__EUPPER_BOUND, notifications)
		}
		if newEUpperBound != nil {
			notifications = newEUpperBound.(EObjectInternal).EInverseAdd(eGenericType, EOPPOSITE_FEATURE_BASE-EGENERIC_TYPE__EUPPER_BOUND, notifications)
		}
		notifications = eGenericType.basicSetEUpperBound(newEUpperBound, notifications)
		if notifications != nil {
			notifications.Dispatch()
		}
	}
}

func (eGenericType *eGenericTypeImpl) basicSetEUpperBound(newEUpperBound EGenericType, msgs ENotificationChain) ENotificationChain {
	oldEUpperBound := eGenericType.eUpperBound
	eGenericType.eUpperBound = newEUpperBound
	notifications := msgs
	if eGenericType.ENotificationRequired() {
		notification := NewNotificationByFeatureID(eGenericType.GetEObject(), SET, EGENERIC_TYPE__EUPPER_BOUND, oldEUpperBound, newEUpperBound, NO_INDEX)
		if notifications != nil {
			notifications.Add(notification)
		} else {
			notifications = notification
		}
	}
	return notifications
}

func (eGenericType *eGenericTypeImpl) initETypeArguments() EList {
	return NewEObjectEList(eGenericType.GetEObjectInternal(), EGENERIC_TYPE__ETYPE_ARGUMENTS, -1, true, true, false, false, false)
}

func (eGenericType *eGenericTypeImpl) EGetFromID(featureID int, resolve, coreType bool) interface{} {
	switch featureID {
	case EGENERIC_TYPE__ECLASSIFIER:
		return eGenericType.GetEClassifier()
	case EGENERIC_TYPE__ELOWER_BOUND:
		return eGenericType.GetELowerBound()
	case EGENERIC_TYPE__ERAW_TYPE:
		return eGenericType.GetERawType()
	case EGENERIC_TYPE__ETYPE_ARGUMENTS:
		return eGenericType.GetETypeArguments()
	case EGENERIC_TYPE__ETYPE_PARAMETER:
		return eGenericType.GetETypeParameter()
	case EGENERIC_TYPE__EUPPER_BOUND:
		return eGenericType.GetEUpperBound()
	default:
		return eGenericType.EObjectImpl.EGetFromID(featureID, resolve, coreType)
	}
}

func (eGenericType *eGenericTypeImpl) ESetFromID(featureID int, newValue interface{}) {
	switch featureID {
	case EGENERIC_TYPE__ECLASSIFIER:
		e := newValue.(EClassifier)
		eGenericType.SetEClassifier(e)
	case EGENERIC_TYPE__ELOWER_BOUND:
		e := newValue.(EGenericType)
		eGenericType.SetELowerBound(e)
	case EGENERIC_TYPE__ETYPE_ARGUMENTS:
		e := newValue.(EList)
		eGenericType.GetETypeArguments().Clear()
		eGenericType.GetETypeArguments().Add(e)
	case EGENERIC_TYPE__ETYPE_PARAMETER:
		e := newValue.(ETypeParameter)
		eGenericType.SetETypeParameter(e)
	case EGENERIC_TYPE__EUPPER_BOUND:
		e := newValue.(EGenericType)
		eGenericType.SetEUpperBound(e)
	default:
		eGenericType.EObjectImpl.ESetFromID(featureID, newValue)
	}
}

func (eGenericType *eGenericTypeImpl) EUnsetFromID(featureID int) {
	switch featureID {
	case EGENERIC_TYPE__ECLASSIFIER:
		eGenericType.SetEClassifier(nil)
	case EGENERIC_TYPE__ELOWER_BOUND:
		eGenericType.SetELowerBound(nil)
	case EGENERIC_TYPE__ETYPE_ARGUMENTS:
		eGenericType.GetETypeArguments().Clear()
	case EGENERIC_TYPE__ETYPE_PARAMETER:
		eGenericType.SetETypeParameter(nil)
	case EGENERIC_TYPE__EUPPER_BOUND:
		eGenericType.SetEUpperBound(nil)
	default:
		eGenericType.EObjectImpl.EUnsetFromID(featureID)
	}
}

func (eGenericType *eGenericTypeImpl) EIsSetFromID(featureID int) bool {
	switch featureID {
	case EGENERIC_TYPE__ECLASSIFIER:
		return eGenericType.eClassifier != nil
	case EGENERIC_TYPE__ELOWER_BOUND:
		return eGenericType.eLowerBound != nil
	case EGENERIC_TYPE__ERAW_TYPE:
		return eGenericType.eRawType != nil
	case EGENERIC_TYPE__ETYPE_ARGUMENTS:
		return eGenericType.eTypeArguments != nil && eGenericType.eTypeArguments.Size() != 0
	case EGENERIC_TYPE__ETYPE_PARAMETER:
		return eGenericType.eTypeParameter != nil
	case EGENERIC_TYPE__EUPPER_BOUND:
		return eGenericType.eUpperBound != nil
	default:
		return eGenericType.EObjectImpl.EIsSetFromID(featureID)
	}
}

func (eGenericType *eGenericTypeImpl) EInvokeFromID(operationID int, arguments EList) interface{} {
	switch operationID {
	case EGENERIC_TYPE__IS_INSTANCE_EJAVAOBJECT:
		return eGenericType.IsInstance(arguments.Get(0))
	default:
		return eGenericType.EObjectImpl.EInvokeFromID(operationID, arguments)
	}
}

func (eGenericType *eGenericTypeImpl) EBasicInverseRemove(otherEnd EObject, featureID int, notifications ENotificationChain) ENotificationChain {
	switch featureID {
	case EGENERIC_TYPE__ELOWER_BOUND:
		return eGenericType.basicSetELowerBound(nil, notifications)
	case EGENERIC_TYPE__ETYPE_ARGUMENTS:
		list := eGenericType.GetETypeArguments().(ENotifyingList)
		return list.RemoveWithNotification(otherEnd, notifications)
	case EGENERIC_TYPE__EUPPER_BOUND:
		return eGenericType.basicSetEUpperBound(nil, notifications)
	default:
		return eGenericType.EObjectImpl.EBasicInverseRemove(otherEnd, featureID, notifications)
	}
}