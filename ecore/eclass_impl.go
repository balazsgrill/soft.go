// *****************************************************************************
//
// This file is part of a MASA library or program.
// Refer to the included end-user license agreement for restrictions.
//
// Copyright (c) 2020 MASA Group
//
// *****************************************************************************

// *****************************************************************************
//
// Warning: This file was generated by soft.generator.go Generator
//
// *****************************************************************************

package ecore

// eClassImpl is the implementation of the model object 'EClass'
type eClassImpl struct {
	*eClassifierExt
	eAllAttributes          EList
	eAllContainments        EList
	eAllOperations          EList
	eAllReferences          EList
	eAllStructuralFeatures  EList
	eAllSuperTypes          EList
	eAttributes             EList
	eContainmentFeatures    EList
	eCrossReferenceFeatures EList
	eIDAttribute            EAttribute
	eOperations             EList
	eReferences             EList
	eStructuralFeatures     EList
	eSuperTypes             EList
	isAbstract              bool
	isInterface             bool
}

// newEClassImpl is the constructor of a eClassImpl
func newEClassImpl() *eClassImpl {
	eClass := new(eClassImpl)
	eClass.eClassifierExt = newEClassifierExt()
	eClass.SetInterfaces(eClass)
	eClass.isAbstract = false
	eClass.isInterface = false

	return eClass
}

type eClassImplInitializers interface {
	initEAllAttributes()
	initEAllContainments()
	initEAllOperations()
	initEAllReferences()
	initEAllStructuralFeatures()
	initEAllSuperTypes()
	initEAttributes()
	initEContainmentFeatures()
	initECrossReferenceFeatures()
	initEIDAttribute()
	initEOperations() EList
	initEReferences()
	initEStructuralFeatures() EList
	initESuperTypes() EList
}

func (eClass *eClassImpl) getInitializers() eClassImplInitializers {
	return eClass.AsEObject().(eClassImplInitializers)
}

func (eClass *eClassImpl) asEClass() EClass {
	return eClass.GetInterfaces().(EClass)
}

func (eClass *eClassImpl) EStaticClass() EClass {
	return GetPackage().GetEClass()
}

// GetEOperation default implementation
func (eClass *eClassImpl) GetEOperation(int) EOperation {
	panic("GetEOperation not implemented")
}

// GetEStructuralFeature default implementation
func (eClass *eClassImpl) GetEStructuralFeature(int) EStructuralFeature {
	panic("GetEStructuralFeature not implemented")
}

// GetEStructuralFeatureFromString default implementation
func (eClass *eClassImpl) GetEStructuralFeatureFromString(string) EStructuralFeature {
	panic("GetEStructuralFeatureFromString not implemented")
}

// GetFeatureCount default implementation
func (eClass *eClassImpl) GetFeatureCount() int {
	panic("GetFeatureCount not implemented")
}

// GetFeatureID default implementation
func (eClass *eClassImpl) GetFeatureID(EStructuralFeature) int {
	panic("GetFeatureID not implemented")
}

// GetFeatureType default implementation
func (eClass *eClassImpl) GetFeatureType(EStructuralFeature) EClassifier {
	panic("GetFeatureType not implemented")
}

// GetOperationCount default implementation
func (eClass *eClassImpl) GetOperationCount() int {
	panic("GetOperationCount not implemented")
}

// GetOperationID default implementation
func (eClass *eClassImpl) GetOperationID(EOperation) int {
	panic("GetOperationID not implemented")
}

// GetOverride default implementation
func (eClass *eClassImpl) GetOverride(EOperation) EOperation {
	panic("GetOverride not implemented")
}

// IsSuperTypeOf default implementation
func (eClass *eClassImpl) IsSuperTypeOf(EClass) bool {
	panic("IsSuperTypeOf not implemented")
}

// GetEAllAttributes get the value of eAllAttributes
func (eClass *eClassImpl) GetEAllAttributes() EList {
	eClass.getInitializers().initEAllAttributes()
	return eClass.eAllAttributes
}

// GetEAllContainments get the value of eAllContainments
func (eClass *eClassImpl) GetEAllContainments() EList {
	eClass.getInitializers().initEAllContainments()
	return eClass.eAllContainments
}

// GetEAllOperations get the value of eAllOperations
func (eClass *eClassImpl) GetEAllOperations() EList {
	eClass.getInitializers().initEAllOperations()
	return eClass.eAllOperations
}

// GetEAllReferences get the value of eAllReferences
func (eClass *eClassImpl) GetEAllReferences() EList {
	eClass.getInitializers().initEAllReferences()
	return eClass.eAllReferences
}

// GetEAllStructuralFeatures get the value of eAllStructuralFeatures
func (eClass *eClassImpl) GetEAllStructuralFeatures() EList {
	eClass.getInitializers().initEAllStructuralFeatures()
	return eClass.eAllStructuralFeatures
}

// GetEAllSuperTypes get the value of eAllSuperTypes
func (eClass *eClassImpl) GetEAllSuperTypes() EList {
	eClass.getInitializers().initEAllSuperTypes()
	return eClass.eAllSuperTypes
}

// GetEAttributes get the value of eAttributes
func (eClass *eClassImpl) GetEAttributes() EList {
	eClass.getInitializers().initEAttributes()
	return eClass.eAttributes
}

// GetEContainmentFeatures get the value of eContainmentFeatures
func (eClass *eClassImpl) GetEContainmentFeatures() EList {
	eClass.getInitializers().initEContainmentFeatures()
	return eClass.eContainmentFeatures
}

// GetECrossReferenceFeatures get the value of eCrossReferenceFeatures
func (eClass *eClassImpl) GetECrossReferenceFeatures() EList {
	eClass.getInitializers().initECrossReferenceFeatures()
	return eClass.eCrossReferenceFeatures
}

// GetEIDAttribute get the value of eIDAttribute
func (eClass *eClassImpl) GetEIDAttribute() EAttribute {
	eClass.getInitializers().initEIDAttribute()
	return eClass.eIDAttribute
}

// GetEOperations get the value of eOperations
func (eClass *eClassImpl) GetEOperations() EList {
	if eClass.eOperations == nil {
		eClass.eOperations = eClass.getInitializers().initEOperations()
	}
	return eClass.eOperations
}

// GetEReferences get the value of eReferences
func (eClass *eClassImpl) GetEReferences() EList {
	eClass.getInitializers().initEReferences()
	return eClass.eReferences
}

// GetEStructuralFeatures get the value of eStructuralFeatures
func (eClass *eClassImpl) GetEStructuralFeatures() EList {
	if eClass.eStructuralFeatures == nil {
		eClass.eStructuralFeatures = eClass.getInitializers().initEStructuralFeatures()
	}
	return eClass.eStructuralFeatures
}

// GetESuperTypes get the value of eSuperTypes
func (eClass *eClassImpl) GetESuperTypes() EList {
	if eClass.eSuperTypes == nil {
		eClass.eSuperTypes = eClass.getInitializers().initESuperTypes()
	}
	return eClass.eSuperTypes
}

// IsAbstract get the value of isAbstract
func (eClass *eClassImpl) IsAbstract() bool {
	return eClass.isAbstract
}

// SetAbstract set the value of isAbstract
func (eClass *eClassImpl) SetAbstract(newIsAbstract bool) {
	oldIsAbstract := eClass.isAbstract
	eClass.isAbstract = newIsAbstract
	if eClass.ENotificationRequired() {
		eClass.ENotify(NewNotificationByFeatureID(eClass.AsEObject(), SET, ECLASS__ABSTRACT, oldIsAbstract, newIsAbstract, NO_INDEX))
	}
}

// IsInterface get the value of isInterface
func (eClass *eClassImpl) IsInterface() bool {
	return eClass.isInterface
}

// SetInterface set the value of isInterface
func (eClass *eClassImpl) SetInterface(newIsInterface bool) {
	oldIsInterface := eClass.isInterface
	eClass.isInterface = newIsInterface
	if eClass.ENotificationRequired() {
		eClass.ENotify(NewNotificationByFeatureID(eClass.AsEObject(), SET, ECLASS__INTERFACE, oldIsInterface, newIsInterface, NO_INDEX))
	}
}

func (eClass *eClassImpl) initEAllAttributes() {
	eClass.eAllAttributes = NewBasicEObjectList(eClass.AsEObjectInternal(), ECLASS__EALL_ATTRIBUTES, -1, false, false, false, true, false)
}

func (eClass *eClassImpl) initEAllContainments() {
	eClass.eAllContainments = NewBasicEObjectList(eClass.AsEObjectInternal(), ECLASS__EALL_CONTAINMENTS, -1, false, false, false, true, false)
}

func (eClass *eClassImpl) initEAllOperations() {
	eClass.eAllOperations = NewBasicEObjectList(eClass.AsEObjectInternal(), ECLASS__EALL_OPERATIONS, -1, false, false, false, true, false)
}

func (eClass *eClassImpl) initEAllReferences() {
	eClass.eAllReferences = NewBasicEObjectList(eClass.AsEObjectInternal(), ECLASS__EALL_REFERENCES, -1, false, false, false, true, false)
}

func (eClass *eClassImpl) initEAllStructuralFeatures() {
	eClass.eAllStructuralFeatures = NewBasicEObjectList(eClass.AsEObjectInternal(), ECLASS__EALL_STRUCTURAL_FEATURES, -1, false, false, false, true, false)
}

func (eClass *eClassImpl) initEAllSuperTypes() {
	eClass.eAllSuperTypes = NewBasicEObjectList(eClass.AsEObjectInternal(), ECLASS__EALL_SUPER_TYPES, -1, false, false, false, true, false)
}

func (eClass *eClassImpl) initEAttributes() {
	eClass.eAttributes = NewBasicEObjectList(eClass.AsEObjectInternal(), ECLASS__EATTRIBUTES, -1, false, false, false, true, false)
}

func (eClass *eClassImpl) initEContainmentFeatures() {
	eClass.eContainmentFeatures = NewBasicEObjectList(eClass.AsEObjectInternal(), ECLASS__ECONTAINMENT_FEATURES, -1, false, false, false, true, false)
}

func (eClass *eClassImpl) initECrossReferenceFeatures() {
	eClass.eCrossReferenceFeatures = NewBasicEObjectList(eClass.AsEObjectInternal(), ECLASS__ECROSS_REFERENCE_FEATURES, -1, false, false, false, true, false)
}

func (eClass *eClassImpl) initEIDAttribute() {
	panic("initEIDAttribute not implemented")
}

func (eClass *eClassImpl) initEOperations() EList {
	return NewBasicEObjectList(eClass.AsEObjectInternal(), ECLASS__EOPERATIONS, EOPERATION__ECONTAINING_CLASS, true, true, true, false, false)
}

func (eClass *eClassImpl) initEReferences() {
	eClass.eReferences = NewBasicEObjectList(eClass.AsEObjectInternal(), ECLASS__EREFERENCES, -1, false, false, false, true, false)
}

func (eClass *eClassImpl) initEStructuralFeatures() EList {
	return NewBasicEObjectList(eClass.AsEObjectInternal(), ECLASS__ESTRUCTURAL_FEATURES, ESTRUCTURAL_FEATURE__ECONTAINING_CLASS, true, true, true, false, false)
}

func (eClass *eClassImpl) initESuperTypes() EList {
	return NewBasicEObjectList(eClass.AsEObjectInternal(), ECLASS__ESUPER_TYPES, -1, false, false, false, true, false)
}

func (eClass *eClassImpl) EGetFromID(featureID int, resolve bool) interface{} {
	switch featureID {
	case ECLASS__ABSTRACT:
		return eClass.asEClass().IsAbstract()
	case ECLASS__EALL_ATTRIBUTES:
		list := eClass.asEClass().GetEAllAttributes()
		if !resolve {
			if objects, _ := list.(EObjectList); objects != nil {
				return objects.GetUnResolvedList()
			}
		}
		return list
	case ECLASS__EALL_CONTAINMENTS:
		list := eClass.asEClass().GetEAllContainments()
		if !resolve {
			if objects, _ := list.(EObjectList); objects != nil {
				return objects.GetUnResolvedList()
			}
		}
		return list
	case ECLASS__EALL_OPERATIONS:
		list := eClass.asEClass().GetEAllOperations()
		if !resolve {
			if objects, _ := list.(EObjectList); objects != nil {
				return objects.GetUnResolvedList()
			}
		}
		return list
	case ECLASS__EALL_REFERENCES:
		list := eClass.asEClass().GetEAllReferences()
		if !resolve {
			if objects, _ := list.(EObjectList); objects != nil {
				return objects.GetUnResolvedList()
			}
		}
		return list
	case ECLASS__EALL_STRUCTURAL_FEATURES:
		list := eClass.asEClass().GetEAllStructuralFeatures()
		if !resolve {
			if objects, _ := list.(EObjectList); objects != nil {
				return objects.GetUnResolvedList()
			}
		}
		return list
	case ECLASS__EALL_SUPER_TYPES:
		list := eClass.asEClass().GetEAllSuperTypes()
		if !resolve {
			if objects, _ := list.(EObjectList); objects != nil {
				return objects.GetUnResolvedList()
			}
		}
		return list
	case ECLASS__EATTRIBUTES:
		list := eClass.asEClass().GetEAttributes()
		if !resolve {
			if objects, _ := list.(EObjectList); objects != nil {
				return objects.GetUnResolvedList()
			}
		}
		return list
	case ECLASS__ECONTAINMENT_FEATURES:
		list := eClass.asEClass().GetEContainmentFeatures()
		if !resolve {
			if objects, _ := list.(EObjectList); objects != nil {
				return objects.GetUnResolvedList()
			}
		}
		return list
	case ECLASS__ECROSS_REFERENCE_FEATURES:
		list := eClass.asEClass().GetECrossReferenceFeatures()
		if !resolve {
			if objects, _ := list.(EObjectList); objects != nil {
				return objects.GetUnResolvedList()
			}
		}
		return list
	case ECLASS__EID_ATTRIBUTE:
		return eClass.asEClass().GetEIDAttribute()
	case ECLASS__EOPERATIONS:
		return eClass.asEClass().GetEOperations()
	case ECLASS__EREFERENCES:
		list := eClass.asEClass().GetEReferences()
		if !resolve {
			if objects, _ := list.(EObjectList); objects != nil {
				return objects.GetUnResolvedList()
			}
		}
		return list
	case ECLASS__ESTRUCTURAL_FEATURES:
		return eClass.asEClass().GetEStructuralFeatures()
	case ECLASS__ESUPER_TYPES:
		list := eClass.asEClass().GetESuperTypes()
		if !resolve {
			if objects, _ := list.(EObjectList); objects != nil {
				return objects.GetUnResolvedList()
			}
		}
		return list
	case ECLASS__INTERFACE:
		return eClass.asEClass().IsInterface()
	default:
		return eClass.eClassifierExt.EGetFromID(featureID, resolve)
	}
}

func (eClass *eClassImpl) ESetFromID(featureID int, newValue interface{}) {
	switch featureID {
	case ECLASS__ABSTRACT:
		eClass.asEClass().SetAbstract(newValue.(bool))
	case ECLASS__EOPERATIONS:
		list := eClass.asEClass().GetEOperations()
		list.Clear()
		list.AddAll(newValue.(EList))
	case ECLASS__ESTRUCTURAL_FEATURES:
		list := eClass.asEClass().GetEStructuralFeatures()
		list.Clear()
		list.AddAll(newValue.(EList))
	case ECLASS__ESUPER_TYPES:
		list := eClass.asEClass().GetESuperTypes()
		list.Clear()
		list.AddAll(newValue.(EList))
	case ECLASS__INTERFACE:
		eClass.asEClass().SetInterface(newValue.(bool))
	default:
		eClass.eClassifierExt.ESetFromID(featureID, newValue)
	}
}

func (eClass *eClassImpl) EUnsetFromID(featureID int) {
	switch featureID {
	case ECLASS__ABSTRACT:
		eClass.asEClass().SetAbstract(false)
	case ECLASS__EOPERATIONS:
		eClass.asEClass().GetEOperations().Clear()
	case ECLASS__ESTRUCTURAL_FEATURES:
		eClass.asEClass().GetEStructuralFeatures().Clear()
	case ECLASS__ESUPER_TYPES:
		eClass.asEClass().GetESuperTypes().Clear()
	case ECLASS__INTERFACE:
		eClass.asEClass().SetInterface(false)
	default:
		eClass.eClassifierExt.EUnsetFromID(featureID)
	}
}

func (eClass *eClassImpl) EIsSetFromID(featureID int) bool {
	switch featureID {
	case ECLASS__ABSTRACT:
		return eClass.isAbstract != false
	case ECLASS__EALL_ATTRIBUTES:
		return eClass.eAllAttributes != nil && eClass.eAllAttributes.Size() != 0
	case ECLASS__EALL_CONTAINMENTS:
		return eClass.eAllContainments != nil && eClass.eAllContainments.Size() != 0
	case ECLASS__EALL_OPERATIONS:
		return eClass.eAllOperations != nil && eClass.eAllOperations.Size() != 0
	case ECLASS__EALL_REFERENCES:
		return eClass.eAllReferences != nil && eClass.eAllReferences.Size() != 0
	case ECLASS__EALL_STRUCTURAL_FEATURES:
		return eClass.eAllStructuralFeatures != nil && eClass.eAllStructuralFeatures.Size() != 0
	case ECLASS__EALL_SUPER_TYPES:
		return eClass.eAllSuperTypes != nil && eClass.eAllSuperTypes.Size() != 0
	case ECLASS__EATTRIBUTES:
		return eClass.eAttributes != nil && eClass.eAttributes.Size() != 0
	case ECLASS__ECONTAINMENT_FEATURES:
		return eClass.eContainmentFeatures != nil && eClass.eContainmentFeatures.Size() != 0
	case ECLASS__ECROSS_REFERENCE_FEATURES:
		return eClass.eCrossReferenceFeatures != nil && eClass.eCrossReferenceFeatures.Size() != 0
	case ECLASS__EID_ATTRIBUTE:
		return eClass.eIDAttribute != nil
	case ECLASS__EOPERATIONS:
		return eClass.eOperations != nil && eClass.eOperations.Size() != 0
	case ECLASS__EREFERENCES:
		return eClass.eReferences != nil && eClass.eReferences.Size() != 0
	case ECLASS__ESTRUCTURAL_FEATURES:
		return eClass.eStructuralFeatures != nil && eClass.eStructuralFeatures.Size() != 0
	case ECLASS__ESUPER_TYPES:
		return eClass.eSuperTypes != nil && eClass.eSuperTypes.Size() != 0
	case ECLASS__INTERFACE:
		return eClass.isInterface != false
	default:
		return eClass.eClassifierExt.EIsSetFromID(featureID)
	}
}

func (eClass *eClassImpl) EInvokeFromID(operationID int, arguments EList) interface{} {
	switch operationID {
	case ECLASS__GET_EOPERATION_EINT:
		return eClass.asEClass().GetEOperation(arguments.Get(0).(int))
	case ECLASS__GET_ESTRUCTURAL_FEATURE_EINT:
		return eClass.asEClass().GetEStructuralFeature(arguments.Get(0).(int))
	case ECLASS__GET_ESTRUCTURAL_FEATURE_ESTRING:
		return eClass.asEClass().GetEStructuralFeatureFromString(arguments.Get(0).(string))
	case ECLASS__GET_FEATURE_COUNT:
		return eClass.asEClass().GetFeatureCount()
	case ECLASS__GET_FEATURE_ID_ESTRUCTURALFEATURE:
		return eClass.asEClass().GetFeatureID(arguments.Get(0).(EStructuralFeature))
	case ECLASS__GET_FEATURE_TYPE_ESTRUCTURALFEATURE:
		return eClass.asEClass().GetFeatureType(arguments.Get(0).(EStructuralFeature))
	case ECLASS__GET_OPERATION_COUNT:
		return eClass.asEClass().GetOperationCount()
	case ECLASS__GET_OPERATION_ID_EOPERATION:
		return eClass.asEClass().GetOperationID(arguments.Get(0).(EOperation))
	case ECLASS__GET_OVERRIDE_EOPERATION:
		return eClass.asEClass().GetOverride(arguments.Get(0).(EOperation))
	case ECLASS__IS_SUPER_TYPE_OF_ECLASS:
		return eClass.asEClass().IsSuperTypeOf(arguments.Get(0).(EClass))
	default:
		return eClass.eClassifierExt.EInvokeFromID(operationID, arguments)
	}
}

func (eClass *eClassImpl) EBasicInverseAdd(otherEnd EObject, featureID int, notifications ENotificationChain) ENotificationChain {
	switch featureID {
	case ECLASS__EOPERATIONS:
		list := eClass.GetEOperations().(ENotifyingList)
		return list.AddWithNotification(otherEnd, notifications)
	case ECLASS__ESTRUCTURAL_FEATURES:
		list := eClass.GetEStructuralFeatures().(ENotifyingList)
		return list.AddWithNotification(otherEnd, notifications)
	default:
		return eClass.eClassifierExt.EBasicInverseAdd(otherEnd, featureID, notifications)
	}
}

func (eClass *eClassImpl) EBasicInverseRemove(otherEnd EObject, featureID int, notifications ENotificationChain) ENotificationChain {
	switch featureID {
	case ECLASS__EOPERATIONS:
		list := eClass.GetEOperations().(ENotifyingList)
		return list.RemoveWithNotification(otherEnd, notifications)
	case ECLASS__ESTRUCTURAL_FEATURES:
		list := eClass.GetEStructuralFeatures().(ENotifyingList)
		return list.RemoveWithNotification(otherEnd, notifications)
	default:
		return eClass.eClassifierExt.EBasicInverseRemove(otherEnd, featureID, notifications)
	}
}
