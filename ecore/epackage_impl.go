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

// EPackageImpl is the implementation of the model object 'EPackage'
type EPackageImpl struct {
	ENamedElementImpl
	eClassifiers     EList
	eFactoryInstance EFactory
	eSubPackages     EList
	nsPrefix         string
	nsURI            string
}
type ePackageInitializers interface {
	initEClassifiers() EList
	initESubPackages() EList
}

type ePackageBasics interface {
	basicSetEFactoryInstance(EFactory, ENotificationChain) ENotificationChain
}

// newEPackageImpl is the constructor of a EPackageImpl
func newEPackageImpl() *EPackageImpl {
	ePackage := new(EPackageImpl)
	ePackage.SetInterfaces(ePackage)
	ePackage.Initialize()
	return ePackage
}

func (ePackage *EPackageImpl) Initialize() {
	ePackage.ENamedElementImpl.Initialize()
	ePackage.nsPrefix = ""
	ePackage.nsURI = ""

}

func (ePackage *EPackageImpl) asEPackage() EPackage {
	return ePackage.GetInterfaces().(EPackage)
}

func (ePackage *EPackageImpl) asInitializers() ePackageInitializers {
	return ePackage.GetInterfaces().(ePackageInitializers)
}

func (ePackage *EPackageImpl) asBasics() ePackageBasics {
	return ePackage.GetInterfaces().(ePackageBasics)
}

func (ePackage *EPackageImpl) EStaticClass() EClass {
	return GetPackage().GetEPackage()
}

func (ePackage *EPackageImpl) EStaticFeatureCount() int {
	return EPACKAGE_FEATURE_COUNT
}

// GetEClassifier default implementation
func (ePackage *EPackageImpl) GetEClassifier(string) EClassifier {
	panic("GetEClassifier not implemented")
}

// GetEClassifiers get the value of eClassifiers
func (ePackage *EPackageImpl) GetEClassifiers() EList {
	if ePackage.eClassifiers == nil {
		ePackage.eClassifiers = ePackage.asInitializers().initEClassifiers()
	}
	return ePackage.eClassifiers
}

// GetEFactoryInstance get the value of eFactoryInstance
func (ePackage *EPackageImpl) GetEFactoryInstance() EFactory {
	return ePackage.eFactoryInstance
}

// SetEFactoryInstance set the value of eFactoryInstance
func (ePackage *EPackageImpl) SetEFactoryInstance(newEFactoryInstance EFactory) {
	if newEFactoryInstance != ePackage.eFactoryInstance {
		var notifications ENotificationChain
		if oldEFactoryInstanceInternal, _ := ePackage.eFactoryInstance.(EObjectInternal); oldEFactoryInstanceInternal != nil {
			notifications = oldEFactoryInstanceInternal.EInverseRemove(ePackage, EFACTORY__EPACKAGE, notifications)
		}
		if newEFactoryInstanceInternal, _ := newEFactoryInstance.(EObjectInternal); newEFactoryInstanceInternal != nil {
			notifications = newEFactoryInstanceInternal.EInverseAdd(ePackage.AsEObject(), EFACTORY__EPACKAGE, notifications)
		}
		notifications = ePackage.asBasics().basicSetEFactoryInstance(newEFactoryInstance, notifications)
		if notifications != nil {
			notifications.Dispatch()
		}
	}
}

func (ePackage *EPackageImpl) basicSetEFactoryInstance(newEFactoryInstance EFactory, msgs ENotificationChain) ENotificationChain {
	oldEFactoryInstance := ePackage.eFactoryInstance
	ePackage.eFactoryInstance = newEFactoryInstance
	notifications := msgs
	if ePackage.ENotificationRequired() {
		notification := NewNotificationByFeatureID(ePackage.AsEObject(), SET, EPACKAGE__EFACTORY_INSTANCE, oldEFactoryInstance, newEFactoryInstance, NO_INDEX)
		if notifications != nil {
			notifications.Add(notification)
		} else {
			notifications = notification
		}
	}
	return notifications
}

// GetESubPackages get the value of eSubPackages
func (ePackage *EPackageImpl) GetESubPackages() EList {
	if ePackage.eSubPackages == nil {
		ePackage.eSubPackages = ePackage.asInitializers().initESubPackages()
	}
	return ePackage.eSubPackages
}

// GetESuperPackage get the value of eSuperPackage
func (ePackage *EPackageImpl) GetESuperPackage() EPackage {
	if ePackage.EContainerFeatureID() == EPACKAGE__ESUPER_PACKAGE {
		return ePackage.EContainer().(EPackage)
	}
	return nil
}

// GetNsPrefix get the value of nsPrefix
func (ePackage *EPackageImpl) GetNsPrefix() string {
	return ePackage.nsPrefix
}

// SetNsPrefix set the value of nsPrefix
func (ePackage *EPackageImpl) SetNsPrefix(newNsPrefix string) {
	oldNsPrefix := ePackage.nsPrefix
	ePackage.nsPrefix = newNsPrefix
	if ePackage.ENotificationRequired() {
		ePackage.ENotify(NewNotificationByFeatureID(ePackage.AsEObject(), SET, EPACKAGE__NS_PREFIX, oldNsPrefix, newNsPrefix, NO_INDEX))
	}
}

// GetNsURI get the value of nsURI
func (ePackage *EPackageImpl) GetNsURI() string {
	return ePackage.nsURI
}

// SetNsURI set the value of nsURI
func (ePackage *EPackageImpl) SetNsURI(newNsURI string) {
	oldNsURI := ePackage.nsURI
	ePackage.nsURI = newNsURI
	if ePackage.ENotificationRequired() {
		ePackage.ENotify(NewNotificationByFeatureID(ePackage.AsEObject(), SET, EPACKAGE__NS_URI, oldNsURI, newNsURI, NO_INDEX))
	}
}

func (ePackage *EPackageImpl) initEClassifiers() EList {
	return NewBasicEObjectList(ePackage.AsEObjectInternal(), EPACKAGE__ECLASSIFIERS, ECLASSIFIER__EPACKAGE, true, true, true, false, false)
}

func (ePackage *EPackageImpl) initESubPackages() EList {
	return NewBasicEObjectList(ePackage.AsEObjectInternal(), EPACKAGE__ESUB_PACKAGES, EPACKAGE__ESUPER_PACKAGE, true, true, true, false, false)
}

func (ePackage *EPackageImpl) EGetFromID(featureID int, resolve bool) any {
	switch featureID {
	case EPACKAGE__ECLASSIFIERS:
		return ePackage.asEPackage().GetEClassifiers()
	case EPACKAGE__EFACTORY_INSTANCE:
		return ePackage.asEPackage().GetEFactoryInstance()
	case EPACKAGE__ESUB_PACKAGES:
		return ePackage.asEPackage().GetESubPackages()
	case EPACKAGE__ESUPER_PACKAGE:
		return ePackage.asEPackage().GetESuperPackage()
	case EPACKAGE__NS_PREFIX:
		return ePackage.asEPackage().GetNsPrefix()
	case EPACKAGE__NS_URI:
		return ePackage.asEPackage().GetNsURI()
	default:
		return ePackage.ENamedElementImpl.EGetFromID(featureID, resolve)
	}
}

func (ePackage *EPackageImpl) ESetFromID(featureID int, newValue any) {
	switch featureID {
	case EPACKAGE__ECLASSIFIERS:
		list := ePackage.asEPackage().GetEClassifiers()
		list.Clear()
		list.AddAll(newValue.(EList))
	case EPACKAGE__EFACTORY_INSTANCE:
		ePackage.asEPackage().SetEFactoryInstance(newValue.(EFactory))
	case EPACKAGE__ESUB_PACKAGES:
		list := ePackage.asEPackage().GetESubPackages()
		list.Clear()
		list.AddAll(newValue.(EList))
	case EPACKAGE__NS_PREFIX:
		ePackage.asEPackage().SetNsPrefix(newValue.(string))
	case EPACKAGE__NS_URI:
		ePackage.asEPackage().SetNsURI(newValue.(string))
	default:
		ePackage.ENamedElementImpl.ESetFromID(featureID, newValue)
	}
}

func (ePackage *EPackageImpl) EUnsetFromID(featureID int) {
	switch featureID {
	case EPACKAGE__ECLASSIFIERS:
		ePackage.asEPackage().GetEClassifiers().Clear()
	case EPACKAGE__EFACTORY_INSTANCE:
		ePackage.asEPackage().SetEFactoryInstance(nil)
	case EPACKAGE__ESUB_PACKAGES:
		ePackage.asEPackage().GetESubPackages().Clear()
	case EPACKAGE__NS_PREFIX:
		ePackage.asEPackage().SetNsPrefix("")
	case EPACKAGE__NS_URI:
		ePackage.asEPackage().SetNsURI("")
	default:
		ePackage.ENamedElementImpl.EUnsetFromID(featureID)
	}
}

func (ePackage *EPackageImpl) EIsSetFromID(featureID int) bool {
	switch featureID {
	case EPACKAGE__ECLASSIFIERS:
		return ePackage.eClassifiers != nil && ePackage.eClassifiers.Size() != 0
	case EPACKAGE__EFACTORY_INSTANCE:
		return ePackage.eFactoryInstance != nil
	case EPACKAGE__ESUB_PACKAGES:
		return ePackage.eSubPackages != nil && ePackage.eSubPackages.Size() != 0
	case EPACKAGE__ESUPER_PACKAGE:
		return ePackage.asEPackage().GetESuperPackage() != nil
	case EPACKAGE__NS_PREFIX:
		return ePackage.nsPrefix != ""
	case EPACKAGE__NS_URI:
		return ePackage.nsURI != ""
	default:
		return ePackage.ENamedElementImpl.EIsSetFromID(featureID)
	}
}

func (ePackage *EPackageImpl) EInvokeFromID(operationID int, arguments EList) any {
	switch operationID {
	case EPACKAGE__GET_ECLASSIFIER_ESTRING:
		return ePackage.asEPackage().GetEClassifier(arguments.Get(0).(string))
	default:
		return ePackage.ENamedElementImpl.EInvokeFromID(operationID, arguments)
	}
}

func (ePackage *EPackageImpl) EBasicInverseAdd(otherEnd EObject, featureID int, notifications ENotificationChain) ENotificationChain {
	switch featureID {
	case EPACKAGE__ECLASSIFIERS:
		list := ePackage.GetEClassifiers().(ENotifyingList)
		return list.AddWithNotification(otherEnd, notifications)
	case EPACKAGE__EFACTORY_INSTANCE:
		msgs := notifications
		eFactoryInstance := ePackage.eFactoryInstance
		if eFactoryInstance != nil {
			msgs = eFactoryInstance.(EObjectInternal).EInverseRemove(ePackage.AsEObject(), EOPPOSITE_FEATURE_BASE-EPACKAGE__EFACTORY_INSTANCE, msgs)
		}
		return ePackage.asBasics().basicSetEFactoryInstance(otherEnd.(EFactory), msgs)
	case EPACKAGE__ESUB_PACKAGES:
		list := ePackage.GetESubPackages().(ENotifyingList)
		return list.AddWithNotification(otherEnd, notifications)
	case EPACKAGE__ESUPER_PACKAGE:
		msgs := notifications
		if ePackage.EInternalContainer() != nil {
			msgs = ePackage.EBasicRemoveFromContainer(msgs)
		}
		return ePackage.EBasicSetContainer(otherEnd, EPACKAGE__ESUPER_PACKAGE, msgs)
	default:
		return ePackage.ENamedElementImpl.EBasicInverseAdd(otherEnd, featureID, notifications)
	}
}

func (ePackage *EPackageImpl) EBasicInverseRemove(otherEnd EObject, featureID int, notifications ENotificationChain) ENotificationChain {
	switch featureID {
	case EPACKAGE__ECLASSIFIERS:
		list := ePackage.GetEClassifiers().(ENotifyingList)
		return list.RemoveWithNotification(otherEnd, notifications)
	case EPACKAGE__EFACTORY_INSTANCE:
		return ePackage.asBasics().basicSetEFactoryInstance(nil, notifications)
	case EPACKAGE__ESUB_PACKAGES:
		list := ePackage.GetESubPackages().(ENotifyingList)
		return list.RemoveWithNotification(otherEnd, notifications)
	case EPACKAGE__ESUPER_PACKAGE:
		return ePackage.EBasicSetContainer(nil, EPACKAGE__ESUPER_PACKAGE, notifications)
	default:
		return ePackage.ENamedElementImpl.EBasicInverseRemove(otherEnd, featureID, notifications)
	}
}
