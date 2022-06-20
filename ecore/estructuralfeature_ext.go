// *****************************************************************************
//
// This file is part of a MASA library or program.
// Refer to the included end-user license agreement for restrictions.
//
// Copyright (c) 2019 MASA Group
//
// *****************************************************************************

package ecore

// eStructuralFeatureExt is the extension of the model object 'EStructuralFeature'
type eStructuralFeatureExt struct {
	eStructuralFeatureImpl
	defaultValue        interface{}
	defaultValueFactory EFactory
}

func newEStructuralFeatureExt() *eStructuralFeatureExt {
	eStructuralFeature := new(eStructuralFeatureExt)
	eStructuralFeature.SetInterfaces(eStructuralFeature)
	eStructuralFeature.Initialize()
	return eStructuralFeature
}

// GetDefaultValue get the value of defaultValue
func (eStructuralFeature *eStructuralFeatureExt) GetDefaultValue() interface{} {
	eType := eStructuralFeature.GetEType()
	defaultValueLiteral := eStructuralFeature.GetDefaultValueLiteral()
	if eType != nil && len(defaultValueLiteral) == 0 {
		if eStructuralFeature.IsMany() {
			return nil
		} else {
			return eType.GetDefaultValue()
		}
	} else if eDataType, _ := eType.(EDataType); eDataType != nil {
		if ePackage := eType.GetEPackage(); ePackage != nil {
			if factory := ePackage.GetEFactoryInstance(); factory != eStructuralFeature.defaultValueFactory {
				if eDataType.IsSerializable() {
					eStructuralFeature.defaultValue = factory.CreateFromString(eDataType, defaultValueLiteral)
				}
				eStructuralFeature.defaultValueFactory = factory
			}
		}
		return eStructuralFeature.defaultValue
	}
	return nil
}

// SetDefaultValue set the value of defaultValue
func (eStructuralFeature *eStructuralFeatureExt) SetDefaultValue(newDefaultValue interface{}) {
	eType := eStructuralFeature.GetEType()
	if eDataType, _ := eType.(EDataType); eDataType != nil {
		factory := eDataType.GetEPackage().GetEFactoryInstance()
		literal := factory.ConvertToString(eDataType, newDefaultValue)
		eStructuralFeature.eStructuralFeatureImpl.SetDefaultValueLiteral(literal)
		eStructuralFeature.defaultValueFactory = nil // reset default value
	} else {
		panic("Cannot serialize value to object without an EDataType eType")
	}
}

// SetDefaultValueLiteral set the value of defaultValueLiteral
func (eStructuralFeature *eStructuralFeatureExt) SetDefaultValueLiteral(newDefaultValueLiteral string) {
	eStructuralFeature.defaultValueFactory = nil // reset default value
	eStructuralFeature.eStructuralFeatureImpl.SetDefaultValueLiteral(newDefaultValueLiteral)
}

// GetFeatureID get the value of featureID
func (eStructuralFeature *eStructuralFeatureExt) GetFeatureID() int {
	return eStructuralFeature.featureID
}

// SetFeatureID set the value of featureID
func (eStructuralFeature *eStructuralFeatureExt) SetFeatureID(newFeatureID int) {
	eStructuralFeature.featureID = newFeatureID
}

func IsBidirectional(feature EStructuralFeature) bool {
	ref, isRef := feature.(EReference)
	if isRef {
		return ref.GetEOpposite() != nil
	}
	return false
}

func IsContainer(feature EStructuralFeature) bool {
	ref, isRef := feature.(EReference)
	if isRef {
		opposite := ref.GetEOpposite()
		if opposite != nil {
			return opposite.IsContainment()
		}
	}
	return false
}

func IsContains(feature EStructuralFeature) bool {
	ref, isRef := feature.(EReference)
	if isRef {
		return ref.IsContainment()
	}
	return false
}

func IsProxy(feature EStructuralFeature) bool {
	ref, isRef := feature.(EReference)
	if isRef {
		return ref.IsResolveProxies()
	}
	return false
}

func IsMapType(feature EStructuralFeature) bool {
	if eType := feature.GetEType(); eType != nil {
		instanceTypeName := eType.GetInstanceTypeName()
		return instanceTypeName == "java.util.Map.Entry" || instanceTypeName == "java.util.Map$Entry"
	}
	return false
}
