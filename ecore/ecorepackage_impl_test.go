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

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPackageInstance(t *testing.T) {
	p := GetPackage()
	assert.NotNil(t, p)
	assert.Equal(t, NAME, p.GetName())
	assert.Equal(t, NS_URI, p.GetNsURI())
	assert.Equal(t, NS_PREFIX, p.GetNsPrefix())

}

func TestPackageClasses(t *testing.T) {
	p := GetPackage()
	assert.NotNil(t, p.GetEAnnotationClass())
	assert.NotNil(t, p.GetEAnnotation_Source())
	assert.NotNil(t, p.GetEAnnotation_Details())
	assert.NotNil(t, p.GetEAnnotation_EModelElement())
	assert.NotNil(t, p.GetEAnnotation_Contents())
	assert.NotNil(t, p.GetEAnnotation_References())
	assert.NotNil(t, p.GetEAttribute())
	assert.NotNil(t, p.GetEAttribute_ID())
	assert.NotNil(t, p.GetEAttribute_EAttributeType())
	assert.NotNil(t, p.GetEClass())
	assert.NotNil(t, p.GetEClass_Abstract())
	assert.NotNil(t, p.GetEClass_Interface())
	assert.NotNil(t, p.GetEClass_EStructuralFeatures())
	assert.NotNil(t, p.GetEClass_EAttributes())
	assert.NotNil(t, p.GetEClass_EReferences())
	assert.NotNil(t, p.GetEClass_ESuperTypes())
	assert.NotNil(t, p.GetEClass_EOperations())
	assert.NotNil(t, p.GetEClass_EContainmentFeatures())
	assert.NotNil(t, p.GetEClass_ECrossReferenceFeatures())
	assert.NotNil(t, p.GetEClass_EAllAttributes())
	assert.NotNil(t, p.GetEClass_EAllReferences())
	assert.NotNil(t, p.GetEClass_EAllContainments())
	assert.NotNil(t, p.GetEClass_EAllOperations())
	assert.NotNil(t, p.GetEClass_EAllStructuralFeatures())
	assert.NotNil(t, p.GetEClass_EAllSuperTypes())
	assert.NotNil(t, p.GetEClass_EIDAttribute())
	assert.NotNil(t, p.GetEClass_IsSuperTypeOf_EClass())
	assert.NotNil(t, p.GetEClass_GetFeatureCount())
	assert.NotNil(t, p.GetEClass_GetEStructuralFeature_EInt())
	assert.NotNil(t, p.GetEClass_GetEStructuralFeature_EString())
	assert.NotNil(t, p.GetEClass_GetFeatureID_EStructuralFeature())
	assert.NotNil(t, p.GetEClass_GetOperationCount())
	assert.NotNil(t, p.GetEClass_GetEOperation_EInt())
	assert.NotNil(t, p.GetEClass_GetOperationID_EOperation())
	assert.NotNil(t, p.GetEClass_GetOverride_EOperation())
	assert.NotNil(t, p.GetEClass_GetFeatureType_EStructuralFeature())
	assert.NotNil(t, p.GetEClassifierClass())
	assert.NotNil(t, p.GetEClassifier_InstanceTypeName())
	assert.NotNil(t, p.GetEClassifier_InstanceClass())
	assert.NotNil(t, p.GetEClassifier_DefaultValue())
	assert.NotNil(t, p.GetEClassifier_EPackage())
	assert.NotNil(t, p.GetEClassifier_ClassifierID())
	assert.NotNil(t, p.GetEClassifier_IsInstance_EJavaObject())
	assert.NotNil(t, p.GetEDataType())
	assert.NotNil(t, p.GetEDataType_Serializable())
	assert.NotNil(t, p.GetEEnum())
	assert.NotNil(t, p.GetEEnum_ELiterals())
	assert.NotNil(t, p.GetEEnum_GetEEnumLiteral_EString())
	assert.NotNil(t, p.GetEEnum_GetEEnumLiteral_EInt())
	assert.NotNil(t, p.GetEEnum_GetEEnumLiteralByLiteral_EString())
	assert.NotNil(t, p.GetEEnumLiteral())
	assert.NotNil(t, p.GetEEnumLiteral_Value())
	assert.NotNil(t, p.GetEEnumLiteral_Instance())
	assert.NotNil(t, p.GetEEnumLiteral_Literal())
	assert.NotNil(t, p.GetEEnumLiteral_EEnum())
	assert.NotNil(t, p.GetEFactory())
	assert.NotNil(t, p.GetEFactory_EPackage())
	assert.NotNil(t, p.GetEFactory_Create_EClass())
	assert.NotNil(t, p.GetEFactory_CreateFromString_EDataType_EString())
	assert.NotNil(t, p.GetEFactory_ConvertToString_EDataType_EJavaObject())
	assert.NotNil(t, p.GetEGenericType())
	assert.NotNil(t, p.GetEGenericType_EUpperBound())
	assert.NotNil(t, p.GetEGenericType_ETypeArguments())
	assert.NotNil(t, p.GetEGenericType_ERawType())
	assert.NotNil(t, p.GetEGenericType_ELowerBound())
	assert.NotNil(t, p.GetEGenericType_ETypeParameter())
	assert.NotNil(t, p.GetEGenericType_EClassifier())
	assert.NotNil(t, p.GetEGenericType_IsInstance_EJavaObject())
	assert.NotNil(t, p.GetEModelElement())
	assert.NotNil(t, p.GetEModelElement_EAnnotations())
	assert.NotNil(t, p.GetEModelElement_GetEAnnotation_EString())
	assert.NotNil(t, p.GetENamedElement())
	assert.NotNil(t, p.GetENamedElement_Name())
	assert.NotNil(t, p.GetEObject())
	assert.NotNil(t, p.GetEObject_EClass())
	assert.NotNil(t, p.GetEObject_EIsProxy())
	assert.NotNil(t, p.GetEObject_EResource())
	assert.NotNil(t, p.GetEObject_EContainer())
	assert.NotNil(t, p.GetEObject_EContainingFeature())
	assert.NotNil(t, p.GetEObject_EContainmentFeature())
	assert.NotNil(t, p.GetEObject_EContents())
	assert.NotNil(t, p.GetEObject_EAllContents())
	assert.NotNil(t, p.GetEObject_ECrossReferences())
	assert.NotNil(t, p.GetEObject_EGet_EStructuralFeature())
	assert.NotNil(t, p.GetEObject_EGet_EStructuralFeature_EBoolean())
	assert.NotNil(t, p.GetEObject_ESet_EStructuralFeature_EJavaObject())
	assert.NotNil(t, p.GetEObject_EIsSet_EStructuralFeature())
	assert.NotNil(t, p.GetEObject_EUnset_EStructuralFeature())
	assert.NotNil(t, p.GetEObject_EInvoke_EOperation_EEList())
	assert.NotNil(t, p.GetEOperation())
	assert.NotNil(t, p.GetEOperation_EContainingClass())
	assert.NotNil(t, p.GetEOperation_EParameters())
	assert.NotNil(t, p.GetEOperation_EExceptions())
	assert.NotNil(t, p.GetEOperation_OperationID())
	assert.NotNil(t, p.GetEOperation_IsOverrideOf_EOperation())
	assert.NotNil(t, p.GetEPackage())
	assert.NotNil(t, p.GetEPackage_NsURI())
	assert.NotNil(t, p.GetEPackage_NsPrefix())
	assert.NotNil(t, p.GetEPackage_EFactoryInstance())
	assert.NotNil(t, p.GetEPackage_EClassifiers())
	assert.NotNil(t, p.GetEPackage_ESubPackages())
	assert.NotNil(t, p.GetEPackage_ESuperPackage())
	assert.NotNil(t, p.GetEPackage_GetEClassifier_EString())
	assert.NotNil(t, p.GetEParameter())
	assert.NotNil(t, p.GetEParameter_EOperation())
	assert.NotNil(t, p.GetEReference())
	assert.NotNil(t, p.GetEReference_Containment())
	assert.NotNil(t, p.GetEReference_Container())
	assert.NotNil(t, p.GetEReference_ResolveProxies())
	assert.NotNil(t, p.GetEReference_EOpposite())
	assert.NotNil(t, p.GetEReference_EReferenceType())
	assert.NotNil(t, p.GetEReference_EKeys())
	assert.NotNil(t, p.GetEStringToStringMapEntry())
	assert.NotNil(t, p.GetEStringToStringMapEntry_Key())
	assert.NotNil(t, p.GetEStringToStringMapEntry_Value())
	assert.NotNil(t, p.GetEStructuralFeature())
	assert.NotNil(t, p.GetEStructuralFeature_Changeable())
	assert.NotNil(t, p.GetEStructuralFeature_Volatile())
	assert.NotNil(t, p.GetEStructuralFeature_Transient())
	assert.NotNil(t, p.GetEStructuralFeature_DefaultValueLiteral())
	assert.NotNil(t, p.GetEStructuralFeature_DefaultValue())
	assert.NotNil(t, p.GetEStructuralFeature_Unsettable())
	assert.NotNil(t, p.GetEStructuralFeature_Derived())
	assert.NotNil(t, p.GetEStructuralFeature_EContainingClass())
	assert.NotNil(t, p.GetEStructuralFeature_FeatureID())
	assert.NotNil(t, p.GetEStructuralFeature_GetContainerClass())
	assert.NotNil(t, p.GetETypeParameter())
	assert.NotNil(t, p.GetETypeParameter_EBounds())
	assert.NotNil(t, p.GetETypedElement())
	assert.NotNil(t, p.GetETypedElement_Ordered())
	assert.NotNil(t, p.GetETypedElement_Unique())
	assert.NotNil(t, p.GetETypedElement_LowerBound())
	assert.NotNil(t, p.GetETypedElement_UpperBound())
	assert.NotNil(t, p.GetETypedElement_Many())
	assert.NotNil(t, p.GetETypedElement_Required())
	assert.NotNil(t, p.GetETypedElement_EType())
}

func TestPackageDataTypes(t *testing.T) {
	p := GetPackage()
	assert.NotNil(t, p.GetEBigDecimal())
	assert.NotNil(t, p.GetEBigInteger())
	assert.NotNil(t, p.GetEBoolean())
	assert.NotNil(t, p.GetEBooleanObject())
	assert.NotNil(t, p.GetEByte())
	assert.NotNil(t, p.GetEByteArray())
	assert.NotNil(t, p.GetEByteObject())
	assert.NotNil(t, p.GetEChar())
	assert.NotNil(t, p.GetECharacterObject())
	assert.NotNil(t, p.GetEDate())
	assert.NotNil(t, p.GetEDiagnosticChain())
	assert.NotNil(t, p.GetEDouble())
	assert.NotNil(t, p.GetEDoubleObject())
	assert.NotNil(t, p.GetEEList())
	assert.NotNil(t, p.GetEEnumerator())
	assert.NotNil(t, p.GetEFeatureMap())
	assert.NotNil(t, p.GetEFeatureMapEntry())
	assert.NotNil(t, p.GetEFloat())
	assert.NotNil(t, p.GetEFloatObject())
	assert.NotNil(t, p.GetEInt())
	assert.NotNil(t, p.GetEIntegerObject())
	assert.NotNil(t, p.GetEInvocationTargetException())
	assert.NotNil(t, p.GetEJavaClass())
	assert.NotNil(t, p.GetEJavaObject())
	assert.NotNil(t, p.GetELong())
	assert.NotNil(t, p.GetELongObject())
	assert.NotNil(t, p.GetEMap())
	assert.NotNil(t, p.GetEResource())
	assert.NotNil(t, p.GetEResourceSet())
	assert.NotNil(t, p.GetEShort())
	assert.NotNil(t, p.GetEShortObject())
	assert.NotNil(t, p.GetEString())
	assert.NotNil(t, p.GetETreeIterator())
}