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

// EClass is the representation of the model object 'EClass'
type EClass interface {
	EClassifier

	IsSuperTypeOf(EClass) bool
	GetFeatureCount() int
	GetEStructuralFeature(int) EStructuralFeature
	GetEStructuralFeatureFromString(string) EStructuralFeature
	GetFeatureID(EStructuralFeature) int
	GetOperationCount() int
	GetEOperation(int) EOperation
	GetOperationID(EOperation) int
	GetOverride(EOperation) EOperation
	GetFeatureType(EStructuralFeature) EClassifier

	IsAbstract() bool
	SetAbstract(bool)

	IsInterface() bool
	SetInterface(bool)

	GetEStructuralFeatures() EList

	GetEAttributes() EList

	GetEReferences() EList

	GetESuperTypes() EList

	GetEOperations() EList

	GetEContainmentFeatures() EList

	GetECrossReferenceFeatures() EList

	GetEAllAttributes() EList

	GetEAllReferences() EList

	GetEAllContainments() EList

	GetEAllOperations() EList

	GetEAllStructuralFeatures() EList

	GetEAllSuperTypes() EList

	GetEIDAttribute() EAttribute

	// Start of user code EClass
	// End of user code
}
