// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLModelStructureProgramNamedValueType] class.
var (
	_MLModelStructureProgramNamedValueTypeClass     MLModelStructureProgramNamedValueTypeClass
	_MLModelStructureProgramNamedValueTypeClassOnce sync.Once
)

func getMLModelStructureProgramNamedValueTypeClass() MLModelStructureProgramNamedValueTypeClass {
	_MLModelStructureProgramNamedValueTypeClassOnce.Do(func() {
		_MLModelStructureProgramNamedValueTypeClass = MLModelStructureProgramNamedValueTypeClass{class: objc.GetClass("MLModelStructureProgramNamedValueType")}
	})
	return _MLModelStructureProgramNamedValueTypeClass
}

// GetMLModelStructureProgramNamedValueTypeClass returns the class object for MLModelStructureProgramNamedValueType.
func GetMLModelStructureProgramNamedValueTypeClass() MLModelStructureProgramNamedValueTypeClass {
	return getMLModelStructureProgramNamedValueTypeClass()
}

type MLModelStructureProgramNamedValueTypeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLModelStructureProgramNamedValueTypeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLModelStructureProgramNamedValueTypeClass) Alloc() MLModelStructureProgramNamedValueType {
	rv := objc.Send[MLModelStructureProgramNamedValueType](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLModelStructureProgramNamedValueType.InitWithMILNamedValueType]
//   - [MLModelStructureProgramNamedValueType.InitWithNameType]
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramNamedValueType
type MLModelStructureProgramNamedValueType struct {
	objectivec.Object
}

// MLModelStructureProgramNamedValueTypeFromID constructs a [MLModelStructureProgramNamedValueType] from an objc.ID.
func MLModelStructureProgramNamedValueTypeFromID(id objc.ID) MLModelStructureProgramNamedValueType {
	return MLModelStructureProgramNamedValueType{objectivec.Object{ID: id}}
}

// Ensure MLModelStructureProgramNamedValueType implements IMLModelStructureProgramNamedValueType.
var _ IMLModelStructureProgramNamedValueType = MLModelStructureProgramNamedValueType{}

// An interface definition for the [MLModelStructureProgramNamedValueType] class.
//
// # Methods
//
//   - [IMLModelStructureProgramNamedValueType.InitWithMILNamedValueType]
//   - [IMLModelStructureProgramNamedValueType.InitWithNameType]
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramNamedValueType
type IMLModelStructureProgramNamedValueType interface {
	objectivec.IObject

	// Topic: Methods

	InitWithMILNamedValueType(type_ unsafe.Pointer) MLModelStructureProgramNamedValueType
	InitWithNameType(name objectivec.IObject, type_ objectivec.IObject) MLModelStructureProgramNamedValueType
}

// Init initializes the instance.
func (m MLModelStructureProgramNamedValueType) Init() MLModelStructureProgramNamedValueType {
	rv := objc.Send[MLModelStructureProgramNamedValueType](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLModelStructureProgramNamedValueType) Autorelease() MLModelStructureProgramNamedValueType {
	rv := objc.Send[MLModelStructureProgramNamedValueType](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLModelStructureProgramNamedValueType creates a new MLModelStructureProgramNamedValueType instance.
func NewMLModelStructureProgramNamedValueType() MLModelStructureProgramNamedValueType {
	class := getMLModelStructureProgramNamedValueTypeClass()
	rv := objc.Send[MLModelStructureProgramNamedValueType](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramNamedValueType/initWithMILNamedValueType:
func NewModelStructureProgramNamedValueTypeWithMILNamedValueType(type_ unsafe.Pointer) MLModelStructureProgramNamedValueType {
	instance := getMLModelStructureProgramNamedValueTypeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMILNamedValueType:"), type_)
	return MLModelStructureProgramNamedValueTypeFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramNamedValueType/initWithName:type:
func NewModelStructureProgramNamedValueTypeWithNameType(name objectivec.IObject, type_ objectivec.IObject) MLModelStructureProgramNamedValueType {
	instance := getMLModelStructureProgramNamedValueTypeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:type:"), name, type_)
	return MLModelStructureProgramNamedValueTypeFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramNamedValueType/initWithMILNamedValueType:
func (m MLModelStructureProgramNamedValueType) InitWithMILNamedValueType(type_ unsafe.Pointer) MLModelStructureProgramNamedValueType {
	rv := objc.Send[MLModelStructureProgramNamedValueType](m.ID, objc.Sel("initWithMILNamedValueType:"), type_)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramNamedValueType/initWithName:type:
func (m MLModelStructureProgramNamedValueType) InitWithNameType(name objectivec.IObject, type_ objectivec.IObject) MLModelStructureProgramNamedValueType {
	rv := objc.Send[MLModelStructureProgramNamedValueType](m.ID, objc.Sel("initWithName:type:"), name, type_)
	return rv
}
