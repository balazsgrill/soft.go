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
	"github.com/stretchr/testify/mock"
)

type MockEClass struct {
	MockEClassifier
}

type MockEClass_Expecter struct {
	MockEClassifier_Expecter
}

func (eClass *MockEClass) EXPECT() *MockEClass_Expecter {
	e := &MockEClass_Expecter{}
	e.Mock = &eClass.Mock
	return e
}

// IsAbstract get the value of isAbstract
func (eClass *MockEClass) IsAbstract() bool {
	ret := eClass.Called()

	var r bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(bool)
		}
	}

	return r
}

type MockEClass_IsAbstract_Call struct {
	*mock.Call
}

func (e *MockEClass_Expecter) IsAbstract() *MockEClass_IsAbstract_Call {
	return &MockEClass_IsAbstract_Call{Call: e.Mock.On("IsAbstract")}
}

func (c *MockEClass_IsAbstract_Call) Run(run func()) *MockEClass_IsAbstract_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEClass_IsAbstract_Call) Return(isAbstract bool) *MockEClass_IsAbstract_Call {
	c.Call.Return(isAbstract)
	return c
}

// SetAbstract provides mock implementation for setting the value of isAbstract
func (eClass *MockEClass) SetAbstract(isAbstract bool) {
	eClass.Called(isAbstract)
}

type MockEClass_SetAbstract_Call struct {
	*mock.Call
}

// SetAbstractis a helper method to define mock.On call
// - isAbstract bool
func (e *MockEClass_Expecter) SetAbstract(isAbstract any) *MockEClass_SetAbstract_Call {
	return &MockEClass_SetAbstract_Call{Call: e.Mock.On("SetAbstract", isAbstract)}
}

func (c *MockEClass_SetAbstract_Call) Run(run func(isAbstract bool)) *MockEClass_SetAbstract_Call {
	c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool))
	})
	return c
}

func (c *MockEClass_SetAbstract_Call) Return() *MockEClass_SetAbstract_Call {
	c.Call.Return()
	return c
}

// GetEAllAttributes get the value of eAllAttributes
func (eClass *MockEClass) GetEAllAttributes() EList {
	ret := eClass.Called()

	var r EList
	if rf, ok := ret.Get(0).(func() EList); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EList)
		}
	}

	return r
}

type MockEClass_GetEAllAttributes_Call struct {
	*mock.Call
}

func (e *MockEClass_Expecter) GetEAllAttributes() *MockEClass_GetEAllAttributes_Call {
	return &MockEClass_GetEAllAttributes_Call{Call: e.Mock.On("GetEAllAttributes")}
}

func (c *MockEClass_GetEAllAttributes_Call) Run(run func()) *MockEClass_GetEAllAttributes_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEClass_GetEAllAttributes_Call) Return(eAllAttributes EList) *MockEClass_GetEAllAttributes_Call {
	c.Call.Return(eAllAttributes)
	return c
}

// GetEAllContainments get the value of eAllContainments
func (eClass *MockEClass) GetEAllContainments() EList {
	ret := eClass.Called()

	var r EList
	if rf, ok := ret.Get(0).(func() EList); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EList)
		}
	}

	return r
}

type MockEClass_GetEAllContainments_Call struct {
	*mock.Call
}

func (e *MockEClass_Expecter) GetEAllContainments() *MockEClass_GetEAllContainments_Call {
	return &MockEClass_GetEAllContainments_Call{Call: e.Mock.On("GetEAllContainments")}
}

func (c *MockEClass_GetEAllContainments_Call) Run(run func()) *MockEClass_GetEAllContainments_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEClass_GetEAllContainments_Call) Return(eAllContainments EList) *MockEClass_GetEAllContainments_Call {
	c.Call.Return(eAllContainments)
	return c
}

// GetEAllOperations get the value of eAllOperations
func (eClass *MockEClass) GetEAllOperations() EList {
	ret := eClass.Called()

	var r EList
	if rf, ok := ret.Get(0).(func() EList); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EList)
		}
	}

	return r
}

type MockEClass_GetEAllOperations_Call struct {
	*mock.Call
}

func (e *MockEClass_Expecter) GetEAllOperations() *MockEClass_GetEAllOperations_Call {
	return &MockEClass_GetEAllOperations_Call{Call: e.Mock.On("GetEAllOperations")}
}

func (c *MockEClass_GetEAllOperations_Call) Run(run func()) *MockEClass_GetEAllOperations_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEClass_GetEAllOperations_Call) Return(eAllOperations EList) *MockEClass_GetEAllOperations_Call {
	c.Call.Return(eAllOperations)
	return c
}

// GetEAllReferences get the value of eAllReferences
func (eClass *MockEClass) GetEAllReferences() EList {
	ret := eClass.Called()

	var r EList
	if rf, ok := ret.Get(0).(func() EList); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EList)
		}
	}

	return r
}

type MockEClass_GetEAllReferences_Call struct {
	*mock.Call
}

func (e *MockEClass_Expecter) GetEAllReferences() *MockEClass_GetEAllReferences_Call {
	return &MockEClass_GetEAllReferences_Call{Call: e.Mock.On("GetEAllReferences")}
}

func (c *MockEClass_GetEAllReferences_Call) Run(run func()) *MockEClass_GetEAllReferences_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEClass_GetEAllReferences_Call) Return(eAllReferences EList) *MockEClass_GetEAllReferences_Call {
	c.Call.Return(eAllReferences)
	return c
}

// GetEAllStructuralFeatures get the value of eAllStructuralFeatures
func (eClass *MockEClass) GetEAllStructuralFeatures() EList {
	ret := eClass.Called()

	var r EList
	if rf, ok := ret.Get(0).(func() EList); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EList)
		}
	}

	return r
}

type MockEClass_GetEAllStructuralFeatures_Call struct {
	*mock.Call
}

func (e *MockEClass_Expecter) GetEAllStructuralFeatures() *MockEClass_GetEAllStructuralFeatures_Call {
	return &MockEClass_GetEAllStructuralFeatures_Call{Call: e.Mock.On("GetEAllStructuralFeatures")}
}

func (c *MockEClass_GetEAllStructuralFeatures_Call) Run(run func()) *MockEClass_GetEAllStructuralFeatures_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEClass_GetEAllStructuralFeatures_Call) Return(eAllStructuralFeatures EList) *MockEClass_GetEAllStructuralFeatures_Call {
	c.Call.Return(eAllStructuralFeatures)
	return c
}

// GetEAllSuperTypes get the value of eAllSuperTypes
func (eClass *MockEClass) GetEAllSuperTypes() EList {
	ret := eClass.Called()

	var r EList
	if rf, ok := ret.Get(0).(func() EList); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EList)
		}
	}

	return r
}

type MockEClass_GetEAllSuperTypes_Call struct {
	*mock.Call
}

func (e *MockEClass_Expecter) GetEAllSuperTypes() *MockEClass_GetEAllSuperTypes_Call {
	return &MockEClass_GetEAllSuperTypes_Call{Call: e.Mock.On("GetEAllSuperTypes")}
}

func (c *MockEClass_GetEAllSuperTypes_Call) Run(run func()) *MockEClass_GetEAllSuperTypes_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEClass_GetEAllSuperTypes_Call) Return(eAllSuperTypes EList) *MockEClass_GetEAllSuperTypes_Call {
	c.Call.Return(eAllSuperTypes)
	return c
}

// GetEAttributes get the value of eAttributes
func (eClass *MockEClass) GetEAttributes() EList {
	ret := eClass.Called()

	var r EList
	if rf, ok := ret.Get(0).(func() EList); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EList)
		}
	}

	return r
}

type MockEClass_GetEAttributes_Call struct {
	*mock.Call
}

func (e *MockEClass_Expecter) GetEAttributes() *MockEClass_GetEAttributes_Call {
	return &MockEClass_GetEAttributes_Call{Call: e.Mock.On("GetEAttributes")}
}

func (c *MockEClass_GetEAttributes_Call) Run(run func()) *MockEClass_GetEAttributes_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEClass_GetEAttributes_Call) Return(eAttributes EList) *MockEClass_GetEAttributes_Call {
	c.Call.Return(eAttributes)
	return c
}

// GetEContainmentFeatures get the value of eContainmentFeatures
func (eClass *MockEClass) GetEContainmentFeatures() EList {
	ret := eClass.Called()

	var r EList
	if rf, ok := ret.Get(0).(func() EList); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EList)
		}
	}

	return r
}

type MockEClass_GetEContainmentFeatures_Call struct {
	*mock.Call
}

func (e *MockEClass_Expecter) GetEContainmentFeatures() *MockEClass_GetEContainmentFeatures_Call {
	return &MockEClass_GetEContainmentFeatures_Call{Call: e.Mock.On("GetEContainmentFeatures")}
}

func (c *MockEClass_GetEContainmentFeatures_Call) Run(run func()) *MockEClass_GetEContainmentFeatures_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEClass_GetEContainmentFeatures_Call) Return(eContainmentFeatures EList) *MockEClass_GetEContainmentFeatures_Call {
	c.Call.Return(eContainmentFeatures)
	return c
}

// GetECrossReferenceFeatures get the value of eCrossReferenceFeatures
func (eClass *MockEClass) GetECrossReferenceFeatures() EList {
	ret := eClass.Called()

	var r EList
	if rf, ok := ret.Get(0).(func() EList); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EList)
		}
	}

	return r
}

type MockEClass_GetECrossReferenceFeatures_Call struct {
	*mock.Call
}

func (e *MockEClass_Expecter) GetECrossReferenceFeatures() *MockEClass_GetECrossReferenceFeatures_Call {
	return &MockEClass_GetECrossReferenceFeatures_Call{Call: e.Mock.On("GetECrossReferenceFeatures")}
}

func (c *MockEClass_GetECrossReferenceFeatures_Call) Run(run func()) *MockEClass_GetECrossReferenceFeatures_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEClass_GetECrossReferenceFeatures_Call) Return(eCrossReferenceFeatures EList) *MockEClass_GetECrossReferenceFeatures_Call {
	c.Call.Return(eCrossReferenceFeatures)
	return c
}

// GetEIDAttribute get the value of eIDAttribute
func (eClass *MockEClass) GetEIDAttribute() EAttribute {
	ret := eClass.Called()

	var r EAttribute
	if rf, ok := ret.Get(0).(func() EAttribute); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EAttribute)
		}
	}

	return r
}

type MockEClass_GetEIDAttribute_Call struct {
	*mock.Call
}

func (e *MockEClass_Expecter) GetEIDAttribute() *MockEClass_GetEIDAttribute_Call {
	return &MockEClass_GetEIDAttribute_Call{Call: e.Mock.On("GetEIDAttribute")}
}

func (c *MockEClass_GetEIDAttribute_Call) Run(run func()) *MockEClass_GetEIDAttribute_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEClass_GetEIDAttribute_Call) Return(eIDAttribute EAttribute) *MockEClass_GetEIDAttribute_Call {
	c.Call.Return(eIDAttribute)
	return c
}

// GetEOperations get the value of eOperations
func (eClass *MockEClass) GetEOperations() EList {
	ret := eClass.Called()

	var r EList
	if rf, ok := ret.Get(0).(func() EList); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EList)
		}
	}

	return r
}

type MockEClass_GetEOperations_Call struct {
	*mock.Call
}

func (e *MockEClass_Expecter) GetEOperations() *MockEClass_GetEOperations_Call {
	return &MockEClass_GetEOperations_Call{Call: e.Mock.On("GetEOperations")}
}

func (c *MockEClass_GetEOperations_Call) Run(run func()) *MockEClass_GetEOperations_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEClass_GetEOperations_Call) Return(eOperations EList) *MockEClass_GetEOperations_Call {
	c.Call.Return(eOperations)
	return c
}

// GetEReferences get the value of eReferences
func (eClass *MockEClass) GetEReferences() EList {
	ret := eClass.Called()

	var r EList
	if rf, ok := ret.Get(0).(func() EList); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EList)
		}
	}

	return r
}

type MockEClass_GetEReferences_Call struct {
	*mock.Call
}

func (e *MockEClass_Expecter) GetEReferences() *MockEClass_GetEReferences_Call {
	return &MockEClass_GetEReferences_Call{Call: e.Mock.On("GetEReferences")}
}

func (c *MockEClass_GetEReferences_Call) Run(run func()) *MockEClass_GetEReferences_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEClass_GetEReferences_Call) Return(eReferences EList) *MockEClass_GetEReferences_Call {
	c.Call.Return(eReferences)
	return c
}

// GetEStructuralFeatures get the value of eStructuralFeatures
func (eClass *MockEClass) GetEStructuralFeatures() EList {
	ret := eClass.Called()

	var r EList
	if rf, ok := ret.Get(0).(func() EList); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EList)
		}
	}

	return r
}

type MockEClass_GetEStructuralFeatures_Call struct {
	*mock.Call
}

func (e *MockEClass_Expecter) GetEStructuralFeatures() *MockEClass_GetEStructuralFeatures_Call {
	return &MockEClass_GetEStructuralFeatures_Call{Call: e.Mock.On("GetEStructuralFeatures")}
}

func (c *MockEClass_GetEStructuralFeatures_Call) Run(run func()) *MockEClass_GetEStructuralFeatures_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEClass_GetEStructuralFeatures_Call) Return(eStructuralFeatures EList) *MockEClass_GetEStructuralFeatures_Call {
	c.Call.Return(eStructuralFeatures)
	return c
}

// GetESuperTypes get the value of eSuperTypes
func (eClass *MockEClass) GetESuperTypes() EList {
	ret := eClass.Called()

	var r EList
	if rf, ok := ret.Get(0).(func() EList); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EList)
		}
	}

	return r
}

type MockEClass_GetESuperTypes_Call struct {
	*mock.Call
}

func (e *MockEClass_Expecter) GetESuperTypes() *MockEClass_GetESuperTypes_Call {
	return &MockEClass_GetESuperTypes_Call{Call: e.Mock.On("GetESuperTypes")}
}

func (c *MockEClass_GetESuperTypes_Call) Run(run func()) *MockEClass_GetESuperTypes_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEClass_GetESuperTypes_Call) Return(eSuperTypes EList) *MockEClass_GetESuperTypes_Call {
	c.Call.Return(eSuperTypes)
	return c
}

// IsInterface get the value of isInterface
func (eClass *MockEClass) IsInterface() bool {
	ret := eClass.Called()

	var r bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(bool)
		}
	}

	return r
}

type MockEClass_IsInterface_Call struct {
	*mock.Call
}

func (e *MockEClass_Expecter) IsInterface() *MockEClass_IsInterface_Call {
	return &MockEClass_IsInterface_Call{Call: e.Mock.On("IsInterface")}
}

func (c *MockEClass_IsInterface_Call) Run(run func()) *MockEClass_IsInterface_Call {
	c.Call.Run(func(mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEClass_IsInterface_Call) Return(isInterface bool) *MockEClass_IsInterface_Call {
	c.Call.Return(isInterface)
	return c
}

// SetInterface provides mock implementation for setting the value of isInterface
func (eClass *MockEClass) SetInterface(isInterface bool) {
	eClass.Called(isInterface)
}

type MockEClass_SetInterface_Call struct {
	*mock.Call
}

// SetInterfaceis a helper method to define mock.On call
// - isInterface bool
func (e *MockEClass_Expecter) SetInterface(isInterface any) *MockEClass_SetInterface_Call {
	return &MockEClass_SetInterface_Call{Call: e.Mock.On("SetInterface", isInterface)}
}

func (c *MockEClass_SetInterface_Call) Run(run func(isInterface bool)) *MockEClass_SetInterface_Call {
	c.Call.Run(func(args mock.Arguments) {
		run(args[0].(bool))
	})
	return c
}

func (c *MockEClass_SetInterface_Call) Return() *MockEClass_SetInterface_Call {
	c.Call.Return()
	return c
}

// GetEOperation provides mock implementation
func (eClass *MockEClass) GetEOperation(operationID int) EOperation {
	ret := eClass.Called(operationID)

	var r EOperation
	if rf, ok := ret.Get(0).(func() EOperation); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EOperation)
		}
	}

	return r
}

type MockEClass_GetEOperation_Call struct {
	*mock.Call
}

// GetEOperationis a helper method to define mock.On call
// - operationID int
func (e *MockEClass_Expecter) GetEOperation(operationID any) *MockEClass_GetEOperation_Call {
	return &MockEClass_GetEOperation_Call{Call: e.Mock.On("GetEOperation", operationID)}
}

func (c *MockEClass_GetEOperation_Call) Run(run func(int)) *MockEClass_GetEOperation_Call {
	c.Call.Run(func(_args mock.Arguments) {
		run(_args[0].(int))
	})
	return c
}

func (c *MockEClass_GetEOperation_Call) Return(_a0 EOperation) *MockEClass_GetEOperation_Call {
	c.Call.Return(_a0)
	return c
}

// GetEStructuralFeature provides mock implementation
func (eClass *MockEClass) GetEStructuralFeature(featureID int) EStructuralFeature {
	ret := eClass.Called(featureID)

	var r EStructuralFeature
	if rf, ok := ret.Get(0).(func() EStructuralFeature); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EStructuralFeature)
		}
	}

	return r
}

type MockEClass_GetEStructuralFeature_Call struct {
	*mock.Call
}

// GetEStructuralFeatureis a helper method to define mock.On call
// - featureID int
func (e *MockEClass_Expecter) GetEStructuralFeature(featureID any) *MockEClass_GetEStructuralFeature_Call {
	return &MockEClass_GetEStructuralFeature_Call{Call: e.Mock.On("GetEStructuralFeature", featureID)}
}

func (c *MockEClass_GetEStructuralFeature_Call) Run(run func(int)) *MockEClass_GetEStructuralFeature_Call {
	c.Call.Run(func(_args mock.Arguments) {
		run(_args[0].(int))
	})
	return c
}

func (c *MockEClass_GetEStructuralFeature_Call) Return(_a0 EStructuralFeature) *MockEClass_GetEStructuralFeature_Call {
	c.Call.Return(_a0)
	return c
}

// GetEStructuralFeatureFromName provides mock implementation
func (eClass *MockEClass) GetEStructuralFeatureFromName(featureName string) EStructuralFeature {
	ret := eClass.Called(featureName)

	var r EStructuralFeature
	if rf, ok := ret.Get(0).(func() EStructuralFeature); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EStructuralFeature)
		}
	}

	return r
}

type MockEClass_GetEStructuralFeatureFromName_Call struct {
	*mock.Call
}

// GetEStructuralFeatureFromNameis a helper method to define mock.On call
// - featureName string
func (e *MockEClass_Expecter) GetEStructuralFeatureFromName(featureName any) *MockEClass_GetEStructuralFeatureFromName_Call {
	return &MockEClass_GetEStructuralFeatureFromName_Call{Call: e.Mock.On("GetEStructuralFeatureFromName", featureName)}
}

func (c *MockEClass_GetEStructuralFeatureFromName_Call) Run(run func(string)) *MockEClass_GetEStructuralFeatureFromName_Call {
	c.Call.Run(func(_args mock.Arguments) {
		run(_args[0].(string))
	})
	return c
}

func (c *MockEClass_GetEStructuralFeatureFromName_Call) Return(_a0 EStructuralFeature) *MockEClass_GetEStructuralFeatureFromName_Call {
	c.Call.Return(_a0)
	return c
}

// GetFeatureCount provides mock implementation
func (eClass *MockEClass) GetFeatureCount() int {
	ret := eClass.Called()

	var r int
	if rf, ok := ret.Get(0).(func() int); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(int)
		}
	}

	return r
}

type MockEClass_GetFeatureCount_Call struct {
	*mock.Call
}

// GetFeatureCountis a helper method to define mock.On call
func (e *MockEClass_Expecter) GetFeatureCount() *MockEClass_GetFeatureCount_Call {
	return &MockEClass_GetFeatureCount_Call{Call: e.Mock.On("GetFeatureCount")}
}

func (c *MockEClass_GetFeatureCount_Call) Run(run func()) *MockEClass_GetFeatureCount_Call {
	c.Call.Run(func(_args mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEClass_GetFeatureCount_Call) Return(_a0 int) *MockEClass_GetFeatureCount_Call {
	c.Call.Return(_a0)
	return c
}

// GetFeatureID provides mock implementation
func (eClass *MockEClass) GetFeatureID(feature EStructuralFeature) int {
	ret := eClass.Called(feature)

	var r int
	if rf, ok := ret.Get(0).(func() int); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(int)
		}
	}

	return r
}

type MockEClass_GetFeatureID_Call struct {
	*mock.Call
}

// GetFeatureIDis a helper method to define mock.On call
// - feature EStructuralFeature
func (e *MockEClass_Expecter) GetFeatureID(feature any) *MockEClass_GetFeatureID_Call {
	return &MockEClass_GetFeatureID_Call{Call: e.Mock.On("GetFeatureID", feature)}
}

func (c *MockEClass_GetFeatureID_Call) Run(run func(EStructuralFeature)) *MockEClass_GetFeatureID_Call {
	c.Call.Run(func(_args mock.Arguments) {
		run(_args[0].(EStructuralFeature))
	})
	return c
}

func (c *MockEClass_GetFeatureID_Call) Return(_a0 int) *MockEClass_GetFeatureID_Call {
	c.Call.Return(_a0)
	return c
}

// GetFeatureType provides mock implementation
func (eClass *MockEClass) GetFeatureType(feature EStructuralFeature) EClassifier {
	ret := eClass.Called(feature)

	var r EClassifier
	if rf, ok := ret.Get(0).(func() EClassifier); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EClassifier)
		}
	}

	return r
}

type MockEClass_GetFeatureType_Call struct {
	*mock.Call
}

// GetFeatureTypeis a helper method to define mock.On call
// - feature EStructuralFeature
func (e *MockEClass_Expecter) GetFeatureType(feature any) *MockEClass_GetFeatureType_Call {
	return &MockEClass_GetFeatureType_Call{Call: e.Mock.On("GetFeatureType", feature)}
}

func (c *MockEClass_GetFeatureType_Call) Run(run func(EStructuralFeature)) *MockEClass_GetFeatureType_Call {
	c.Call.Run(func(_args mock.Arguments) {
		run(_args[0].(EStructuralFeature))
	})
	return c
}

func (c *MockEClass_GetFeatureType_Call) Return(_a0 EClassifier) *MockEClass_GetFeatureType_Call {
	c.Call.Return(_a0)
	return c
}

// GetOperationCount provides mock implementation
func (eClass *MockEClass) GetOperationCount() int {
	ret := eClass.Called()

	var r int
	if rf, ok := ret.Get(0).(func() int); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(int)
		}
	}

	return r
}

type MockEClass_GetOperationCount_Call struct {
	*mock.Call
}

// GetOperationCountis a helper method to define mock.On call
func (e *MockEClass_Expecter) GetOperationCount() *MockEClass_GetOperationCount_Call {
	return &MockEClass_GetOperationCount_Call{Call: e.Mock.On("GetOperationCount")}
}

func (c *MockEClass_GetOperationCount_Call) Run(run func()) *MockEClass_GetOperationCount_Call {
	c.Call.Run(func(_args mock.Arguments) {
		run()
	})
	return c
}

func (c *MockEClass_GetOperationCount_Call) Return(_a0 int) *MockEClass_GetOperationCount_Call {
	c.Call.Return(_a0)
	return c
}

// GetOperationID provides mock implementation
func (eClass *MockEClass) GetOperationID(operation EOperation) int {
	ret := eClass.Called(operation)

	var r int
	if rf, ok := ret.Get(0).(func() int); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(int)
		}
	}

	return r
}

type MockEClass_GetOperationID_Call struct {
	*mock.Call
}

// GetOperationIDis a helper method to define mock.On call
// - operation EOperation
func (e *MockEClass_Expecter) GetOperationID(operation any) *MockEClass_GetOperationID_Call {
	return &MockEClass_GetOperationID_Call{Call: e.Mock.On("GetOperationID", operation)}
}

func (c *MockEClass_GetOperationID_Call) Run(run func(EOperation)) *MockEClass_GetOperationID_Call {
	c.Call.Run(func(_args mock.Arguments) {
		run(_args[0].(EOperation))
	})
	return c
}

func (c *MockEClass_GetOperationID_Call) Return(_a0 int) *MockEClass_GetOperationID_Call {
	c.Call.Return(_a0)
	return c
}

// GetOverride provides mock implementation
func (eClass *MockEClass) GetOverride(operation EOperation) EOperation {
	ret := eClass.Called(operation)

	var r EOperation
	if rf, ok := ret.Get(0).(func() EOperation); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(EOperation)
		}
	}

	return r
}

type MockEClass_GetOverride_Call struct {
	*mock.Call
}

// GetOverrideis a helper method to define mock.On call
// - operation EOperation
func (e *MockEClass_Expecter) GetOverride(operation any) *MockEClass_GetOverride_Call {
	return &MockEClass_GetOverride_Call{Call: e.Mock.On("GetOverride", operation)}
}

func (c *MockEClass_GetOverride_Call) Run(run func(EOperation)) *MockEClass_GetOverride_Call {
	c.Call.Run(func(_args mock.Arguments) {
		run(_args[0].(EOperation))
	})
	return c
}

func (c *MockEClass_GetOverride_Call) Return(_a0 EOperation) *MockEClass_GetOverride_Call {
	c.Call.Return(_a0)
	return c
}

// IsSuperTypeOf provides mock implementation
func (eClass *MockEClass) IsSuperTypeOf(someClass EClass) bool {
	ret := eClass.Called(someClass)

	var r bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r = rf()
	} else {
		if ret.Get(0) != nil {
			r = ret.Get(0).(bool)
		}
	}

	return r
}

type MockEClass_IsSuperTypeOf_Call struct {
	*mock.Call
}

// IsSuperTypeOfis a helper method to define mock.On call
// - someClass EClass
func (e *MockEClass_Expecter) IsSuperTypeOf(someClass any) *MockEClass_IsSuperTypeOf_Call {
	return &MockEClass_IsSuperTypeOf_Call{Call: e.Mock.On("IsSuperTypeOf", someClass)}
}

func (c *MockEClass_IsSuperTypeOf_Call) Run(run func(EClass)) *MockEClass_IsSuperTypeOf_Call {
	c.Call.Run(func(_args mock.Arguments) {
		run(_args[0].(EClass))
	})
	return c
}

func (c *MockEClass_IsSuperTypeOf_Call) Return(_a0 bool) *MockEClass_IsSuperTypeOf_Call {
	c.Call.Return(_a0)
	return c
}

type mockConstructorTestingTNewMockEClass interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockEClass creates a new instance of MockEClass. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockEClass(t mockConstructorTestingTNewMockEClass) *MockEClass {
	mock := &MockEClass{}
	mock.Mock.Test(t)
	t.Cleanup(func() { mock.AssertExpectations(t) })
	return mock
}
