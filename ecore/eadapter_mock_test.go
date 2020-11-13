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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMockEAdapter_GetTarget(t *testing.T) {
	mockNotifier := new(MockENotifier)
	mockAdapter := new(MockEAdapter)
	mockAdapter.On("GetTarget").Return(mockNotifier).Once()
	assert.Equal(t, mockNotifier, mockAdapter.GetTarget())
	mock.AssertExpectationsForObjects(t, mockNotifier, mockAdapter)

	mockAdapter.On("GetTarget").Return(func() ENotifier { return mockNotifier }).Once()
	assert.Equal(t, mockNotifier, mockAdapter.GetTarget())
	mock.AssertExpectationsForObjects(t, mockNotifier, mockAdapter)
}
