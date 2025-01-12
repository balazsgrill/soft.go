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

import "reflect"

// EStructuralFeature is the representation of the model object 'EStructuralFeature'
type EStructuralFeature interface {
	ETypedElement

	GetContainerClass() reflect.Type

	IsChangeable() bool
	SetChangeable(bool)

	IsVolatile() bool
	SetVolatile(bool)

	IsTransient() bool
	SetTransient(bool)

	GetDefaultValueLiteral() string
	SetDefaultValueLiteral(string)

	GetDefaultValue() any
	SetDefaultValue(any)

	IsUnsettable() bool
	SetUnsettable(bool)

	IsDerived() bool
	SetDerived(bool)

	GetFeatureID() int
	SetFeatureID(int)

	GetEContainingClass() EClass

	// Start of user code EStructuralFeature
	// End of user code
}
