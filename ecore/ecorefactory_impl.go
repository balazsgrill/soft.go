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

import "strconv"
import "time"

type ecoreFactoryInternal interface {
	createEBigDecimalFromString(eDataType EDataType, literalValue string) any
	createEBigIntegerFromString(eDataType EDataType, literalValue string) any
	createEBooleanFromString(eDataType EDataType, literalValue string) any
	createEBooleanObjectFromString(eDataType EDataType, literalValue string) any
	createEByteFromString(eDataType EDataType, literalValue string) any
	createEByteArrayFromString(eDataType EDataType, literalValue string) any
	createEByteObjectFromString(eDataType EDataType, literalValue string) any
	createECharFromString(eDataType EDataType, literalValue string) any
	createECharacterObjectFromString(eDataType EDataType, literalValue string) any
	createEDateFromString(eDataType EDataType, literalValue string) any
	createEDoubleFromString(eDataType EDataType, literalValue string) any
	createEDoubleObjectFromString(eDataType EDataType, literalValue string) any
	createEFloatFromString(eDataType EDataType, literalValue string) any
	createEFloatObjectFromString(eDataType EDataType, literalValue string) any
	createEIntFromString(eDataType EDataType, literalValue string) any
	createEIntegerObjectFromString(eDataType EDataType, literalValue string) any
	createEJavaClassFromString(eDataType EDataType, literalValue string) any
	createEJavaObjectFromString(eDataType EDataType, literalValue string) any
	createELongFromString(eDataType EDataType, literalValue string) any
	createELongObjectFromString(eDataType EDataType, literalValue string) any
	createEShortFromString(eDataType EDataType, literalValue string) any
	createEShortObjectFromString(eDataType EDataType, literalValue string) any
	createEStringFromString(eDataType EDataType, literalValue string) any
	convertEBigDecimalToString(eDataType EDataType, literalValue any) string
	convertEBigIntegerToString(eDataType EDataType, literalValue any) string
	convertEBooleanToString(eDataType EDataType, literalValue any) string
	convertEBooleanObjectToString(eDataType EDataType, literalValue any) string
	convertEByteToString(eDataType EDataType, literalValue any) string
	convertEByteArrayToString(eDataType EDataType, literalValue any) string
	convertEByteObjectToString(eDataType EDataType, literalValue any) string
	convertECharToString(eDataType EDataType, literalValue any) string
	convertECharacterObjectToString(eDataType EDataType, literalValue any) string
	convertEDateToString(eDataType EDataType, literalValue any) string
	convertEDoubleToString(eDataType EDataType, literalValue any) string
	convertEDoubleObjectToString(eDataType EDataType, literalValue any) string
	convertEFloatToString(eDataType EDataType, literalValue any) string
	convertEFloatObjectToString(eDataType EDataType, literalValue any) string
	convertEIntToString(eDataType EDataType, literalValue any) string
	convertEIntegerObjectToString(eDataType EDataType, literalValue any) string
	convertEJavaClassToString(eDataType EDataType, literalValue any) string
	convertEJavaObjectToString(eDataType EDataType, literalValue any) string
	convertELongToString(eDataType EDataType, literalValue any) string
	convertELongObjectToString(eDataType EDataType, literalValue any) string
	convertEShortToString(eDataType EDataType, literalValue any) string
	convertEShortObjectToString(eDataType EDataType, literalValue any) string
	convertEStringToString(eDataType EDataType, literalValue any) string
}

type EcoreFactoryImpl struct {
	EFactoryExt
}

func newEcoreFactoryImpl() *EcoreFactoryImpl {
	factory := new(EcoreFactoryImpl)
	factory.SetInterfaces(factory)
	factory.Initialize()
	return factory
}

func (ecoreFactoryImpl *EcoreFactoryImpl) AsInternal() ecoreFactoryInternal {
	return ecoreFactoryImpl.GetInterfaces().(ecoreFactoryInternal)
}

func (ecoreFactoryImpl *EcoreFactoryImpl) AsEFactory() EcoreFactory {
	return ecoreFactoryImpl.GetInterfaces().(EcoreFactory)
}

func (ecoreFactoryImpl *EcoreFactoryImpl) Create(eClass EClass) EObject {
	classID := eClass.GetClassifierID()
	switch classID {
	case EANNOTATION:
		return ecoreFactoryImpl.AsEFactory().CreateEAnnotation()
	case EATTRIBUTE:
		return ecoreFactoryImpl.AsEFactory().CreateEAttribute()
	case ECLASS:
		return ecoreFactoryImpl.AsEFactory().CreateEClass()
	case EDATA_TYPE:
		return ecoreFactoryImpl.AsEFactory().CreateEDataType()
	case EENUM:
		return ecoreFactoryImpl.AsEFactory().CreateEEnum()
	case EENUM_LITERAL:
		return ecoreFactoryImpl.AsEFactory().CreateEEnumLiteral()
	case EFACTORY:
		return ecoreFactoryImpl.AsEFactory().CreateEFactory()
	case EGENERIC_TYPE:
		return ecoreFactoryImpl.AsEFactory().CreateEGenericType()
	case EOBJECT:
		return ecoreFactoryImpl.AsEFactory().CreateEObject()
	case EOPERATION:
		return ecoreFactoryImpl.AsEFactory().CreateEOperation()
	case EPACKAGE:
		return ecoreFactoryImpl.AsEFactory().CreateEPackage()
	case EPARAMETER:
		return ecoreFactoryImpl.AsEFactory().CreateEParameter()
	case EREFERENCE:
		return ecoreFactoryImpl.AsEFactory().CreateEReference()
	case ESTRING_TO_STRING_MAP_ENTRY:
		return ecoreFactoryImpl.AsEFactory().CreateEStringToStringMapEntry()
	case ETYPE_PARAMETER:
		return ecoreFactoryImpl.AsEFactory().CreateETypeParameter()
	default:
		panic("Create: " + strconv.Itoa(classID) + " not found")
	}
}

func (ecoreFactoryImpl *EcoreFactoryImpl) CreateEAnnotation() EAnnotation {
	return newEAnnotationImpl()
}

func (ecoreFactoryImpl *EcoreFactoryImpl) CreateEAnnotationFromContainer(eContainer EModelElement) EAnnotation {
	element := ecoreFactoryImpl.AsEFactory().CreateEAnnotation()
	if eContainer != nil {
		eContainer.GetEAnnotations().Add(element)
	}
	return element
}

func (ecoreFactoryImpl *EcoreFactoryImpl) CreateEAttribute() EAttribute {
	return newEAttributeExt()
}

func (ecoreFactoryImpl *EcoreFactoryImpl) CreateEAttributeFromContainer(eContainer EClass) EAttribute {
	element := ecoreFactoryImpl.AsEFactory().CreateEAttribute()
	if eContainer != nil {
		eContainer.GetEStructuralFeatures().Add(element)
	}
	return element
}

func (ecoreFactoryImpl *EcoreFactoryImpl) CreateEAttributeFromContainerAndClassID(eContainer EClass, classID int) EAttribute {
	element := ecoreFactoryImpl.AsEFactory().CreateEAttribute()
	element.SetFeatureID(classID)
	if eContainer != nil {
		eContainer.GetEStructuralFeatures().Add(element)
	}
	return element
}

func (ecoreFactoryImpl *EcoreFactoryImpl) CreateEClass() EClass {
	return newEClassExt()
}

func (ecoreFactoryImpl *EcoreFactoryImpl) CreateEClassFromContainer(eContainer EPackage) EClass {
	element := ecoreFactoryImpl.AsEFactory().CreateEClass()
	if eContainer != nil {
		eContainer.GetEClassifiers().Add(element)
	}
	return element
}

func (ecoreFactoryImpl *EcoreFactoryImpl) CreateEClassFromContainerAndClassID(eContainer EPackage, classID int) EClass {
	element := ecoreFactoryImpl.AsEFactory().CreateEClass()
	element.SetClassifierID(classID)
	if eContainer != nil {
		eContainer.GetEClassifiers().Add(element)
	}
	return element
}

func (ecoreFactoryImpl *EcoreFactoryImpl) CreateEDataType() EDataType {
	return newEDataTypeExt()
}

func (ecoreFactoryImpl *EcoreFactoryImpl) CreateEDataTypeFromContainer(eContainer EPackage) EDataType {
	element := ecoreFactoryImpl.AsEFactory().CreateEDataType()
	if eContainer != nil {
		eContainer.GetEClassifiers().Add(element)
	}
	return element
}

func (ecoreFactoryImpl *EcoreFactoryImpl) CreateEDataTypeFromContainerAndClassID(eContainer EPackage, classID int) EDataType {
	element := ecoreFactoryImpl.AsEFactory().CreateEDataType()
	element.SetClassifierID(classID)
	if eContainer != nil {
		eContainer.GetEClassifiers().Add(element)
	}
	return element
}

func (ecoreFactoryImpl *EcoreFactoryImpl) CreateEEnum() EEnum {
	return newEEnumExt()
}

func (ecoreFactoryImpl *EcoreFactoryImpl) CreateEEnumFromContainer(eContainer EPackage) EEnum {
	element := ecoreFactoryImpl.AsEFactory().CreateEEnum()
	if eContainer != nil {
		eContainer.GetEClassifiers().Add(element)
	}
	return element
}

func (ecoreFactoryImpl *EcoreFactoryImpl) CreateEEnumFromContainerAndClassID(eContainer EPackage, classID int) EEnum {
	element := ecoreFactoryImpl.AsEFactory().CreateEEnum()
	element.SetClassifierID(classID)
	if eContainer != nil {
		eContainer.GetEClassifiers().Add(element)
	}
	return element
}

func (ecoreFactoryImpl *EcoreFactoryImpl) CreateEEnumLiteral() EEnumLiteral {
	return newEEnumLiteralExt()
}

func (ecoreFactoryImpl *EcoreFactoryImpl) CreateEEnumLiteralFromContainer(eContainer EEnum) EEnumLiteral {
	element := ecoreFactoryImpl.AsEFactory().CreateEEnumLiteral()
	if eContainer != nil {
		eContainer.GetELiterals().Add(element)
	}
	return element
}

func (ecoreFactoryImpl *EcoreFactoryImpl) CreateEFactory() EFactory {
	return newEFactoryExt()
}

func (ecoreFactoryImpl *EcoreFactoryImpl) CreateEFactoryFromContainer(eContainer EPackage) EFactory {
	element := ecoreFactoryImpl.AsEFactory().CreateEFactory()
	if eContainer != nil {
		eContainer.SetEFactoryInstance(element)
	}
	return element
}

func (ecoreFactoryImpl *EcoreFactoryImpl) CreateEGenericType() EGenericType {
	return newEGenericTypeImpl()
}

func (ecoreFactoryImpl *EcoreFactoryImpl) CreateEObject() EObject {
	return newEObjectImpl()
}

func (ecoreFactoryImpl *EcoreFactoryImpl) CreateEOperation() EOperation {
	return newEOperationExt()
}

func (ecoreFactoryImpl *EcoreFactoryImpl) CreateEOperationFromContainer(eContainer EClass) EOperation {
	element := ecoreFactoryImpl.AsEFactory().CreateEOperation()
	if eContainer != nil {
		eContainer.GetEOperations().Add(element)
	}
	return element
}

func (ecoreFactoryImpl *EcoreFactoryImpl) CreateEOperationFromContainerAndClassID(eContainer EClass, classID int) EOperation {
	element := ecoreFactoryImpl.AsEFactory().CreateEOperation()
	element.SetOperationID(classID)
	if eContainer != nil {
		eContainer.GetEOperations().Add(element)
	}
	return element
}

func (ecoreFactoryImpl *EcoreFactoryImpl) CreateEPackage() EPackage {
	return newEPackageExt()
}

func (ecoreFactoryImpl *EcoreFactoryImpl) CreateEPackageFromContainer(eContainer EPackage) EPackage {
	element := ecoreFactoryImpl.AsEFactory().CreateEPackage()
	if eContainer != nil {
		eContainer.GetESubPackages().Add(element)
	}
	return element
}

func (ecoreFactoryImpl *EcoreFactoryImpl) CreateEParameter() EParameter {
	return newEParameterImpl()
}

func (ecoreFactoryImpl *EcoreFactoryImpl) CreateEParameterFromContainer(eContainer EOperation) EParameter {
	element := ecoreFactoryImpl.AsEFactory().CreateEParameter()
	if eContainer != nil {
		eContainer.GetEParameters().Add(element)
	}
	return element
}

func (ecoreFactoryImpl *EcoreFactoryImpl) CreateEReference() EReference {
	return newEReferenceExt()
}

func (ecoreFactoryImpl *EcoreFactoryImpl) CreateEReferenceFromContainer(eContainer EClass) EReference {
	element := ecoreFactoryImpl.AsEFactory().CreateEReference()
	if eContainer != nil {
		eContainer.GetEStructuralFeatures().Add(element)
	}
	return element
}

func (ecoreFactoryImpl *EcoreFactoryImpl) CreateEReferenceFromContainerAndClassID(eContainer EClass, classID int) EReference {
	element := ecoreFactoryImpl.AsEFactory().CreateEReference()
	element.SetFeatureID(classID)
	if eContainer != nil {
		eContainer.GetEStructuralFeatures().Add(element)
	}
	return element
}

func (ecoreFactoryImpl *EcoreFactoryImpl) CreateEStringToStringMapEntry() EStringToStringMapEntry {
	return newEStringToStringMapEntryImpl()
}

func (ecoreFactoryImpl *EcoreFactoryImpl) CreateETypeParameter() ETypeParameter {
	return newETypeParameterImpl()
}

func (ecoreFactoryImpl *EcoreFactoryImpl) CreateFromString(eDataType EDataType, literalValue string) any {
	classID := eDataType.GetClassifierID()
	switch classID {
	case EBIG_DECIMAL:
		return ecoreFactoryImpl.AsInternal().createEBigDecimalFromString(eDataType, literalValue)
	case EBIG_INTEGER:
		return ecoreFactoryImpl.AsInternal().createEBigIntegerFromString(eDataType, literalValue)
	case EBOOLEAN:
		return ecoreFactoryImpl.AsInternal().createEBooleanFromString(eDataType, literalValue)
	case EBOOLEAN_OBJECT:
		return ecoreFactoryImpl.AsInternal().createEBooleanObjectFromString(eDataType, literalValue)
	case EBYTE:
		return ecoreFactoryImpl.AsInternal().createEByteFromString(eDataType, literalValue)
	case EBYTE_ARRAY:
		return ecoreFactoryImpl.AsInternal().createEByteArrayFromString(eDataType, literalValue)
	case EBYTE_OBJECT:
		return ecoreFactoryImpl.AsInternal().createEByteObjectFromString(eDataType, literalValue)
	case ECHAR:
		return ecoreFactoryImpl.AsInternal().createECharFromString(eDataType, literalValue)
	case ECHARACTER_OBJECT:
		return ecoreFactoryImpl.AsInternal().createECharacterObjectFromString(eDataType, literalValue)
	case EDATE:
		return ecoreFactoryImpl.AsInternal().createEDateFromString(eDataType, literalValue)
	case EDOUBLE:
		return ecoreFactoryImpl.AsInternal().createEDoubleFromString(eDataType, literalValue)
	case EDOUBLE_OBJECT:
		return ecoreFactoryImpl.AsInternal().createEDoubleObjectFromString(eDataType, literalValue)
	case EFLOAT:
		return ecoreFactoryImpl.AsInternal().createEFloatFromString(eDataType, literalValue)
	case EFLOAT_OBJECT:
		return ecoreFactoryImpl.AsInternal().createEFloatObjectFromString(eDataType, literalValue)
	case EINT:
		return ecoreFactoryImpl.AsInternal().createEIntFromString(eDataType, literalValue)
	case EINTEGER_OBJECT:
		return ecoreFactoryImpl.AsInternal().createEIntegerObjectFromString(eDataType, literalValue)
	case EJAVA_CLASS:
		return ecoreFactoryImpl.AsInternal().createEJavaClassFromString(eDataType, literalValue)
	case EJAVA_OBJECT:
		return ecoreFactoryImpl.AsInternal().createEJavaObjectFromString(eDataType, literalValue)
	case ELONG:
		return ecoreFactoryImpl.AsInternal().createELongFromString(eDataType, literalValue)
	case ELONG_OBJECT:
		return ecoreFactoryImpl.AsInternal().createELongObjectFromString(eDataType, literalValue)
	case ESHORT:
		return ecoreFactoryImpl.AsInternal().createEShortFromString(eDataType, literalValue)
	case ESHORT_OBJECT:
		return ecoreFactoryImpl.AsInternal().createEShortObjectFromString(eDataType, literalValue)
	case ESTRING:
		return ecoreFactoryImpl.AsInternal().createEStringFromString(eDataType, literalValue)
	default:
		panic("The datatype '" + eDataType.GetName() + "' is not a valid classifier")
	}
}

func (ecoreFactoryImpl *EcoreFactoryImpl) ConvertToString(eDataType EDataType, instanceValue any) string {
	classID := eDataType.GetClassifierID()
	switch classID {
	case EBIG_DECIMAL:
		return ecoreFactoryImpl.AsInternal().convertEBigDecimalToString(eDataType, instanceValue)
	case EBIG_INTEGER:
		return ecoreFactoryImpl.AsInternal().convertEBigIntegerToString(eDataType, instanceValue)
	case EBOOLEAN:
		return ecoreFactoryImpl.AsInternal().convertEBooleanToString(eDataType, instanceValue)
	case EBOOLEAN_OBJECT:
		return ecoreFactoryImpl.AsInternal().convertEBooleanObjectToString(eDataType, instanceValue)
	case EBYTE:
		return ecoreFactoryImpl.AsInternal().convertEByteToString(eDataType, instanceValue)
	case EBYTE_ARRAY:
		return ecoreFactoryImpl.AsInternal().convertEByteArrayToString(eDataType, instanceValue)
	case EBYTE_OBJECT:
		return ecoreFactoryImpl.AsInternal().convertEByteObjectToString(eDataType, instanceValue)
	case ECHAR:
		return ecoreFactoryImpl.AsInternal().convertECharToString(eDataType, instanceValue)
	case ECHARACTER_OBJECT:
		return ecoreFactoryImpl.AsInternal().convertECharacterObjectToString(eDataType, instanceValue)
	case EDATE:
		return ecoreFactoryImpl.AsInternal().convertEDateToString(eDataType, instanceValue)
	case EDOUBLE:
		return ecoreFactoryImpl.AsInternal().convertEDoubleToString(eDataType, instanceValue)
	case EDOUBLE_OBJECT:
		return ecoreFactoryImpl.AsInternal().convertEDoubleObjectToString(eDataType, instanceValue)
	case EFLOAT:
		return ecoreFactoryImpl.AsInternal().convertEFloatToString(eDataType, instanceValue)
	case EFLOAT_OBJECT:
		return ecoreFactoryImpl.AsInternal().convertEFloatObjectToString(eDataType, instanceValue)
	case EINT:
		return ecoreFactoryImpl.AsInternal().convertEIntToString(eDataType, instanceValue)
	case EINTEGER_OBJECT:
		return ecoreFactoryImpl.AsInternal().convertEIntegerObjectToString(eDataType, instanceValue)
	case EJAVA_CLASS:
		return ecoreFactoryImpl.AsInternal().convertEJavaClassToString(eDataType, instanceValue)
	case EJAVA_OBJECT:
		return ecoreFactoryImpl.AsInternal().convertEJavaObjectToString(eDataType, instanceValue)
	case ELONG:
		return ecoreFactoryImpl.AsInternal().convertELongToString(eDataType, instanceValue)
	case ELONG_OBJECT:
		return ecoreFactoryImpl.AsInternal().convertELongObjectToString(eDataType, instanceValue)
	case ESHORT:
		return ecoreFactoryImpl.AsInternal().convertEShortToString(eDataType, instanceValue)
	case ESHORT_OBJECT:
		return ecoreFactoryImpl.AsInternal().convertEShortObjectToString(eDataType, instanceValue)
	case ESTRING:
		return ecoreFactoryImpl.AsInternal().convertEStringToString(eDataType, instanceValue)
	default:
		panic("The datatype '" + eDataType.GetName() + "' is not a valid classifier")
	}
}
func (ecoreFactoryImpl *EcoreFactoryImpl) createEBigDecimalFromString(eDataType EDataType, literalValue string) any {
	value, _ := strconv.ParseFloat(literalValue, 64)
	return value
}

func (ecoreFactoryImpl *EcoreFactoryImpl) convertEBigDecimalToString(eDataType EDataType, instanceValue any) string {
	v, _ := instanceValue.(float64)
	return strconv.FormatFloat(v, 'f', -1, 64)
}

func (ecoreFactoryImpl *EcoreFactoryImpl) createEBigIntegerFromString(eDataType EDataType, literalValue string) any {
	value, _ := strconv.ParseInt(literalValue, 10, 64)
	return value
}

func (ecoreFactoryImpl *EcoreFactoryImpl) convertEBigIntegerToString(eDataType EDataType, instanceValue any) string {
	v, _ := instanceValue.(int64)
	return strconv.FormatInt(v, 10)
}

func (ecoreFactoryImpl *EcoreFactoryImpl) createEBooleanFromString(eDataType EDataType, literalValue string) any {
	value, _ := strconv.ParseBool(literalValue)
	return value
}

func (ecoreFactoryImpl *EcoreFactoryImpl) convertEBooleanToString(eDataType EDataType, instanceValue any) string {
	v, _ := instanceValue.(bool)
	return strconv.FormatBool(v)
}

func (ecoreFactoryImpl *EcoreFactoryImpl) createEBooleanObjectFromString(eDataType EDataType, literalValue string) any {
	value, _ := strconv.ParseBool(literalValue)
	return value
}

func (ecoreFactoryImpl *EcoreFactoryImpl) convertEBooleanObjectToString(eDataType EDataType, instanceValue any) string {
	v, _ := instanceValue.(bool)
	return strconv.FormatBool(v)
}

func (ecoreFactoryImpl *EcoreFactoryImpl) createEByteFromString(eDataType EDataType, literalValue string) any {
	if len(literalValue) == 0 {
		return "golang\u0000"
	} else {
		return []byte(literalValue)[0]
	}
}

func (ecoreFactoryImpl *EcoreFactoryImpl) convertEByteToString(eDataType EDataType, instanceValue any) string {
	b := instanceValue.(byte)
	return string([]byte{b})
}

func (ecoreFactoryImpl *EcoreFactoryImpl) createEByteArrayFromString(eDataType EDataType, literalValue string) any {
	return []byte(literalValue)
}

func (ecoreFactoryImpl *EcoreFactoryImpl) convertEByteArrayToString(eDataType EDataType, instanceValue any) string {
	b := instanceValue.([]byte)
	return string(b)
}

func (ecoreFactoryImpl *EcoreFactoryImpl) createEByteObjectFromString(eDataType EDataType, literalValue string) any {
	if len(literalValue) == 0 {
		return "golang\u0000"
	} else {
		return []byte(literalValue)[0]
	}
}

func (ecoreFactoryImpl *EcoreFactoryImpl) convertEByteObjectToString(eDataType EDataType, instanceValue any) string {
	b := instanceValue.(byte)
	return string([]byte{b})
}

func (ecoreFactoryImpl *EcoreFactoryImpl) createECharFromString(eDataType EDataType, literalValue string) any {
	if len(literalValue) == 0 {
		return "golang\u0000"
	} else {
		return []byte(literalValue)[0]
	}
}

func (ecoreFactoryImpl *EcoreFactoryImpl) convertECharToString(eDataType EDataType, instanceValue any) string {
	b := instanceValue.(byte)
	return string([]byte{b})
}

func (ecoreFactoryImpl *EcoreFactoryImpl) createECharacterObjectFromString(eDataType EDataType, literalValue string) any {
	if len(literalValue) == 0 {
		return "golang\u0000"
	} else {
		return []byte(literalValue)[0]
	}
}

func (ecoreFactoryImpl *EcoreFactoryImpl) convertECharacterObjectToString(eDataType EDataType, instanceValue any) string {
	b := instanceValue.(byte)
	return string([]byte{b})
}

const (
	dateFormat string = "2006-01-02T15:04:05.999Z"
)

func (ecoreFactoryImpl *EcoreFactoryImpl) createEDateFromString(eDataType EDataType, literalValue string) any {
	t, _ := time.Parse(dateFormat, literalValue)
	return &t
}

func (ecoreFactoryImpl *EcoreFactoryImpl) convertEDateToString(eDataType EDataType, instanceValue any) string {
	t, _ := instanceValue.(*time.Time)
	return t.Format(dateFormat)
}

func (ecoreFactoryImpl *EcoreFactoryImpl) createEDoubleFromString(eDataType EDataType, literalValue string) any {
	value, _ := strconv.ParseFloat(literalValue, 64)
	return value
}

func (ecoreFactoryImpl *EcoreFactoryImpl) convertEDoubleToString(eDataType EDataType, instanceValue any) string {
	v, _ := instanceValue.(float64)
	return strconv.FormatFloat(v, 'f', -1, 64)
}

func (ecoreFactoryImpl *EcoreFactoryImpl) createEDoubleObjectFromString(eDataType EDataType, literalValue string) any {
	value, _ := strconv.ParseFloat(literalValue, 64)
	return value
}

func (ecoreFactoryImpl *EcoreFactoryImpl) convertEDoubleObjectToString(eDataType EDataType, instanceValue any) string {
	v, _ := instanceValue.(float64)
	return strconv.FormatFloat(v, 'f', -1, 64)
}

func (ecoreFactoryImpl *EcoreFactoryImpl) createEFloatFromString(eDataType EDataType, literalValue string) any {
	value, _ := strconv.ParseFloat(literalValue, 32)
	return float32(value)
}

func (ecoreFactoryImpl *EcoreFactoryImpl) convertEFloatToString(eDataType EDataType, instanceValue any) string {
	v, _ := instanceValue.(float32)
	return strconv.FormatFloat(float64(v), 'f', -1, 32)
}

func (ecoreFactoryImpl *EcoreFactoryImpl) createEFloatObjectFromString(eDataType EDataType, literalValue string) any {
	value, _ := strconv.ParseFloat(literalValue, 32)
	return float32(value)
}

func (ecoreFactoryImpl *EcoreFactoryImpl) convertEFloatObjectToString(eDataType EDataType, instanceValue any) string {
	v, _ := instanceValue.(float32)
	return strconv.FormatFloat(float64(v), 'f', -1, 32)
}

func (ecoreFactoryImpl *EcoreFactoryImpl) createEIntFromString(eDataType EDataType, literalValue string) any {
	value, _ := strconv.Atoi(literalValue)
	return value
}

func (ecoreFactoryImpl *EcoreFactoryImpl) convertEIntToString(eDataType EDataType, instanceValue any) string {
	v, _ := instanceValue.(int)
	return strconv.Itoa(v)
}

func (ecoreFactoryImpl *EcoreFactoryImpl) createEIntegerObjectFromString(eDataType EDataType, literalValue string) any {
	value, _ := strconv.Atoi(literalValue)
	return value
}

func (ecoreFactoryImpl *EcoreFactoryImpl) convertEIntegerObjectToString(eDataType EDataType, instanceValue any) string {
	v, _ := instanceValue.(int)
	return strconv.Itoa(v)
}

func (ecoreFactoryImpl *EcoreFactoryImpl) createEJavaClassFromString(eDataType EDataType, literalValue string) any {
	panic("NotImplementedException")
}

func (ecoreFactoryImpl *EcoreFactoryImpl) convertEJavaClassToString(eDataType EDataType, instanceValue any) string {
	panic("NotImplementedException")
}

func (ecoreFactoryImpl *EcoreFactoryImpl) createEJavaObjectFromString(eDataType EDataType, literalValue string) any {
	panic("NotImplementedException")
}

func (ecoreFactoryImpl *EcoreFactoryImpl) convertEJavaObjectToString(eDataType EDataType, instanceValue any) string {
	panic("NotImplementedException")
}

func (ecoreFactoryImpl *EcoreFactoryImpl) createELongFromString(eDataType EDataType, literalValue string) any {
	value, _ := strconv.ParseInt(literalValue, 10, 64)
	return value
}

func (ecoreFactoryImpl *EcoreFactoryImpl) convertELongToString(eDataType EDataType, instanceValue any) string {
	v, _ := instanceValue.(int64)
	return strconv.FormatInt(v, 10)
}

func (ecoreFactoryImpl *EcoreFactoryImpl) createELongObjectFromString(eDataType EDataType, literalValue string) any {
	value, _ := strconv.Atoi(literalValue)
	return value
}

func (ecoreFactoryImpl *EcoreFactoryImpl) convertELongObjectToString(eDataType EDataType, instanceValue any) string {
	v, _ := instanceValue.(int)
	return strconv.Itoa(v)
}

func (ecoreFactoryImpl *EcoreFactoryImpl) createEShortFromString(eDataType EDataType, literalValue string) any {
	value, _ := strconv.ParseInt(literalValue, 10, 16)
	return int16(value)
}

func (ecoreFactoryImpl *EcoreFactoryImpl) convertEShortToString(eDataType EDataType, instanceValue any) string {
	v, _ := instanceValue.(int16)
	return strconv.FormatInt(int64(v), 10)
}

func (ecoreFactoryImpl *EcoreFactoryImpl) createEShortObjectFromString(eDataType EDataType, literalValue string) any {
	value, _ := strconv.ParseInt(literalValue, 10, 16)
	return int16(value)
}

func (ecoreFactoryImpl *EcoreFactoryImpl) convertEShortObjectToString(eDataType EDataType, instanceValue any) string {
	v, _ := instanceValue.(int16)
	return strconv.FormatInt(int64(v), 10)
}

func (ecoreFactoryImpl *EcoreFactoryImpl) createEStringFromString(eDataType EDataType, literalValue string) any {
	return literalValue
}

func (ecoreFactoryImpl *EcoreFactoryImpl) convertEStringToString(eDataType EDataType, instanceValue any) string {
	v, _ := instanceValue.(string)
	return v
}
