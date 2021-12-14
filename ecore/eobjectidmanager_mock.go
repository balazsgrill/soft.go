// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package ecore

import mock "github.com/stretchr/testify/mock"

// MockEObjectIDManager is an autogenerated mock type for the EObjectIDManager type
type MockEObjectIDManager struct {
	mock.Mock
}

// Clear provides a mock function with given fields:
func (_m *MockEObjectIDManager) Clear() {
	_m.Called()
}

// GetEObject provides a mock function with given fields: _a0
func (_m *MockEObjectIDManager) GetEObject(_a0 interface{}) EObject {
	ret := _m.Called(_a0)

	var r0 EObject
	if rf, ok := ret.Get(0).(func(interface{}) EObject); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(EObject)
		}
	}

	return r0
}

// GetID provides a mock function with given fields: _a0
func (_m *MockEObjectIDManager) GetID(_a0 EObject) interface{} {
	ret := _m.Called(_a0)

	var r0 interface{}
	if rf, ok := ret.Get(0).(func(EObject) interface{}); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0)
	}

	return r0
}

// Register provides a mock function with given fields: _a0
func (_m *MockEObjectIDManager) Register(_a0 EObject) {
	_m.Called(_a0)
}

// UnRegister provides a mock function with given fields: _a0
func (_m *MockEObjectIDManager) UnRegister(_a0 EObject) {
	_m.Called(_a0)
}

func (_m *MockEObjectIDManager) SetID(_a0 EObject, _a1 interface{}) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(EObject, interface{}) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(error)
		}
	}

	return r0
}
