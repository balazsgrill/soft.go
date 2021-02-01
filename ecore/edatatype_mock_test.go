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

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func discardMockEDataType() {
	_ = assert.Equal
	_ = testing.Coverage
}

// TestMockEDataTypeIsSerializable tests method IsSerializable
func TestMockEDataTypeIsSerializable(t *testing.T) {
	o := &MockEDataType{}
	r := bool(true)
	o.On("IsSerializable").Once().Return(r)
	o.On("IsSerializable").Once().Return(func() bool {
		return r
	})
	assert.Equal(t, r, o.IsSerializable())
	assert.Equal(t, r, o.IsSerializable())
	o.AssertExpectations(t)
}

// TestMockEDataTypeSetSerializable tests method SetSerializable
func TestMockEDataTypeSetSerializable(t *testing.T) {
	o := &MockEDataType{}
	v := bool(true)
	o.On("SetSerializable", v).Once()
	o.SetSerializable(v)
	o.AssertExpectations(t)
}
