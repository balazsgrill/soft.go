package ecore

import (
	io "io"
	url "net/url"
)

// MockEResource is an autogenerated mock type for the EResource type
type MockEResource struct {
	MockENotifier
}

// Attached provides a mock function with given fields: object
func (_m *MockEResource) Attached(object EObject) {
	_m.Called(object)
}

// Detached provides a mock function with given fields: object
func (_m *MockEResource) Detached(object EObject) {
	_m.Called(object)
}

// GetAllContents provides a mock function with given fields:
func (_m *MockEResource) GetAllContents() EIterator {
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

// GetContents provides a mock function with given fields:
func (_m *MockEResource) GetContents() EList {
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

// GetEObject provides a mock function with given fields: _a0
func (_m *MockEResource) GetEObject(_a0 string) EObject {
	ret := _m.Called(_a0)

	var r0 EObject
	if rf, ok := ret.Get(0).(func(string) EObject); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(EObject)
		}
	}

	return r0
}

// GetErrors provides a mock function with given fields:
func (_m *MockEResource) GetErrors() EList {
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

// GetIDManager provides a mock function with given fields:
func (_m *MockEResource) GetObjectIDManager() EObjectIDManager {
	ret := _m.Called()

	var r0 EObjectIDManager
	if rf, ok := ret.Get(0).(func() EObjectIDManager); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(EObjectIDManager)
		}
	}

	return r0
}

// GetResourceSet provides a mock function with given fields:
func (_m *MockEResource) GetResourceSet() EResourceSet {
	ret := _m.Called()

	var r0 EResourceSet
	if rf, ok := ret.Get(0).(func() EResourceSet); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(EResourceSet)
		}
	}

	return r0
}

// GetURI provides a mock function with given fields:
func (_m *MockEResource) GetURI() *url.URL {
	ret := _m.Called()

	var r0 *url.URL
	if rf, ok := ret.Get(0).(func() *url.URL); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*url.URL)
		}
	}

	return r0
}

// GetURIFragment provides a mock function with given fields: _a0
func (_m *MockEResource) GetURIFragment(_a0 EObject) string {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(EObject) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetWarnings provides a mock function with given fields:
func (_m *MockEResource) GetWarnings() EList {
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

// IsLoaded provides a mock function with given fields:
func (_m *MockEResource) IsLoaded() bool {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// Load provides a mock function with given fields:
func (_m *MockEResource) Load() {
	_m.Called()
}

func (_m *MockEResource) LoadWithOptions(o map[string]interface{}) {
	_m.Called(o)
}

// LoadWithReader provides a mock function with given fields: r
func (_m *MockEResource) LoadWithReader(r io.Reader, o map[string]interface{}) {
	_m.Called(r, o)
}

// Save provides a mock function with given fields:
func (_m *MockEResource) Save() {
	_m.Called()
}

func (_m *MockEResource) SaveWithOptions(o map[string]interface{}) {
	_m.Called(o)
}

// SaveWithWriter provides a mock function with given fields: w
func (_m *MockEResource) SaveWithWriter(w io.Writer, o map[string]interface{}) {
	_m.Called(w, o)
}

// SetIDManager provides a mock function with given fields: _a0
func (_m *MockEResource) SetObjectIDManager(_a0 EObjectIDManager) {
	_m.Called(_a0)
}

// SetURI provides a mock function with given fields: _a0
func (_m *MockEResource) SetURI(_a0 *url.URL) {
	_m.Called(_a0)
}

// Unload provides a mock function with given fields:
func (_m *MockEResource) Unload() {
	_m.Called()
}
