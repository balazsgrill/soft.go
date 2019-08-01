// Code generated by mockery v1.0.0. DO NOT EDIT.

package ecore

import mock "github.com/stretchr/testify/mock"

// MockEOperation is an autogenerated mock type for the EOperation type
type MockEOperation struct {
	mock.Mock
}

// EAdapters provides a mock function with given fields:
func (_m *MockEOperation) EAdapters() EList {
	ret := _m.Called()

	var r0 EList
	if rf, ok := ret.Get(0).(func() EList); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(EList)
		}
	}

	return r0
}

// EAllContents provides a mock function with given fields:
func (_m *MockEOperation) EAllContents() EIterator {
	ret := _m.Called()

	var r0 EIterator
	if rf, ok := ret.Get(0).(func() EIterator); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(EIterator)
		}
	}

	return r0
}

// EClass provides a mock function with given fields:
func (_m *MockEOperation) EClass() EClass {
	ret := _m.Called()

	var r0 EClass
	if rf, ok := ret.Get(0).(func() EClass); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(EClass)
		}
	}

	return r0
}

// EContainer provides a mock function with given fields:
func (_m *MockEOperation) EContainer() EObject {
	ret := _m.Called()

	var r0 EObject
	if rf, ok := ret.Get(0).(func() EObject); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(EObject)
		}
	}

	return r0
}

// EContainingFeature provides a mock function with given fields:
func (_m *MockEOperation) EContainingFeature() EStructuralFeature {
	ret := _m.Called()

	var r0 EStructuralFeature
	if rf, ok := ret.Get(0).(func() EStructuralFeature); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(EStructuralFeature)
		}
	}

	return r0
}

// EContainmentFeature provides a mock function with given fields:
func (_m *MockEOperation) EContainmentFeature() EReference {
	ret := _m.Called()

	var r0 EReference
	if rf, ok := ret.Get(0).(func() EReference); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(EReference)
		}
	}

	return r0
}

// EContents provides a mock function with given fields:
func (_m *MockEOperation) EContents() EList {
	ret := _m.Called()

	var r0 EList
	if rf, ok := ret.Get(0).(func() EList); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(EList)
		}
	}

	return r0
}

// ECrossReferences provides a mock function with given fields:
func (_m *MockEOperation) ECrossReferences() EList {
	ret := _m.Called()

	var r0 EList
	if rf, ok := ret.Get(0).(func() EList); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(EList)
		}
	}

	return r0
}

// EDeliver provides a mock function with given fields:
func (_m *MockEOperation) EDeliver() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// EGet provides a mock function with given fields: _a0
func (_m *MockEOperation) EGet(_a0 EStructuralFeature) interface{} {
	ret := _m.Called(_a0)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(EStructuralFeature) interface{}); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// EGetResolve provides a mock function with given fields: _a0, _a1
func (_m *MockEOperation) EGetResolve(_a0 EStructuralFeature, _a1 bool) interface{} {
	ret := _m.Called(_a0, _a1)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(EStructuralFeature, bool) interface{}); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// EInvoke provides a mock function with given fields: _a0, _a1
func (_m *MockEOperation) EInvoke(_a0 EOperation, _a1 EList) interface{} {
	ret := _m.Called(_a0, _a1)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(EOperation, EList) interface{}); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(interface{})
		}
	}

	return r0
}

// EIsProxy provides a mock function with given fields:
func (_m *MockEOperation) EIsProxy() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// EIsSet provides a mock function with given fields: _a0
func (_m *MockEOperation) EIsSet(_a0 EStructuralFeature) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(EStructuralFeature) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ENotificationRequired provides a mock function with given fields:
func (_m *MockEOperation) ENotificationRequired() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// ENotify provides a mock function with given fields: _a0
func (_m *MockEOperation) ENotify(_a0 ENotification) {
	_m.Called(_a0)
}

// EResource provides a mock function with given fields:
func (_m *MockEOperation) EResource() EResource {
	ret := _m.Called()

	var r0 EResource
	if rf, ok := ret.Get(0).(func() EResource); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(EResource)
		}
	}

	return r0
}

// ESet provides a mock function with given fields: _a0, _a1
func (_m *MockEOperation) ESet(_a0 EStructuralFeature, _a1 interface{}) {
	_m.Called(_a0, _a1)
}

// ESetDeliver provides a mock function with given fields: _a0
func (_m *MockEOperation) ESetDeliver(_a0 bool) {
	_m.Called(_a0)
}

// EUnset provides a mock function with given fields: _a0
func (_m *MockEOperation) EUnset(_a0 EStructuralFeature) {
	_m.Called(_a0)
}

// GetEAnnotation provides a mock function with given fields: _a0
func (_m *MockEOperation) GetEAnnotation(_a0 string) EAnnotation {
	ret := _m.Called(_a0)

	var r0 EAnnotation
	if rf, ok := ret.Get(0).(func(string) EAnnotation); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(EAnnotation)
		}
	}

	return r0
}

// GetEAnnotations provides a mock function with given fields:
func (_m *MockEOperation) GetEAnnotations() EList {
	ret := _m.Called()

	var r0 EList
	if rf, ok := ret.Get(0).(func() EList); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(EList)
		}
	}

	return r0
}

// GetEContainingClass provides a mock function with given fields:
func (_m *MockEOperation) GetEContainingClass() EClass {
	ret := _m.Called()

	var r0 EClass
	if rf, ok := ret.Get(0).(func() EClass); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(EClass)
		}
	}

	return r0
}

// GetEExceptions provides a mock function with given fields:
func (_m *MockEOperation) GetEExceptions() EList {
	ret := _m.Called()

	var r0 EList
	if rf, ok := ret.Get(0).(func() EList); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(EList)
		}
	}

	return r0
}

// GetEParameters provides a mock function with given fields:
func (_m *MockEOperation) GetEParameters() EList {
	ret := _m.Called()

	var r0 EList
	if rf, ok := ret.Get(0).(func() EList); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(EList)
		}
	}

	return r0
}

// GetEType provides a mock function with given fields:
func (_m *MockEOperation) GetEType() EClassifier {
	ret := _m.Called()

	var r0 EClassifier
	if rf, ok := ret.Get(0).(func() EClassifier); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(EClassifier)
		}
	}

	return r0
}

// GetLowerBound provides a mock function with given fields:
func (_m *MockEOperation) GetLowerBound() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GetName provides a mock function with given fields:
func (_m *MockEOperation) GetName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetOperationID provides a mock function with given fields:
func (_m *MockEOperation) GetOperationID() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// GetUpperBound provides a mock function with given fields:
func (_m *MockEOperation) GetUpperBound() int {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	return r0
}

// IsMany provides a mock function with given fields:
func (_m *MockEOperation) IsMany() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsOrdered provides a mock function with given fields:
func (_m *MockEOperation) IsOrdered() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsOverrideOf provides a mock function with given fields: _a0
func (_m *MockEOperation) IsOverrideOf(_a0 EOperation) bool {
	ret := _m.Called(_a0)

	var r0 bool
	if rf, ok := ret.Get(0).(func(EOperation) bool); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsRequired provides a mock function with given fields:
func (_m *MockEOperation) IsRequired() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// IsUnique provides a mock function with given fields:
func (_m *MockEOperation) IsUnique() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// SetEType provides a mock function with given fields: _a0
func (_m *MockEOperation) SetEType(_a0 EClassifier) {
	_m.Called(_a0)
}

// SetLowerBound provides a mock function with given fields: _a0
func (_m *MockEOperation) SetLowerBound(_a0 int) {
	_m.Called(_a0)
}

// SetName provides a mock function with given fields: _a0
func (_m *MockEOperation) SetName(_a0 string) {
	_m.Called(_a0)
}

// SetOperationID provides a mock function with given fields: _a0
func (_m *MockEOperation) SetOperationID(_a0 int) {
	_m.Called(_a0)
}

// SetOrdered provides a mock function with given fields: _a0
func (_m *MockEOperation) SetOrdered(_a0 bool) {
	_m.Called(_a0)
}

// SetUnique provides a mock function with given fields: _a0
func (_m *MockEOperation) SetUnique(_a0 bool) {
	_m.Called(_a0)
}

// SetUpperBound provides a mock function with given fields: _a0
func (_m *MockEOperation) SetUpperBound(_a0 int) {
	_m.Called(_a0)
}

// UnsetEExceptions provides a mock function with given fields:
func (_m *MockEOperation) UnsetEExceptions() {
	_m.Called()
}

// UnsetEType provides a mock function with given fields:
func (_m *MockEOperation) UnsetEType() {
	_m.Called()
}
