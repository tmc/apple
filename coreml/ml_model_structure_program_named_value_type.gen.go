// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
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

// A class representing a named value type in a Program.
//
// # Accessing the value type properties
//
//   - [MLModelStructureProgramNamedValueType.Name]: The name of the parameter.
//   - [MLModelStructureProgramNamedValueType.Type]: The type of the parameter.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramNamedValueType
type MLModelStructureProgramNamedValueType struct {
	objectivec.Object
}

// MLModelStructureProgramNamedValueTypeFromID constructs a [MLModelStructureProgramNamedValueType] from an objc.ID.
//
// A class representing a named value type in a Program.
func MLModelStructureProgramNamedValueTypeFromID(id objc.ID) MLModelStructureProgramNamedValueType {
	return MLModelStructureProgramNamedValueType{objectivec.Object{ID: id}}
}

// Ensure MLModelStructureProgramNamedValueType implements IMLModelStructureProgramNamedValueType.
var _ IMLModelStructureProgramNamedValueType = MLModelStructureProgramNamedValueType{}

// An interface definition for the [MLModelStructureProgramNamedValueType] class.
//
// # Accessing the value type properties
//
//   - [IMLModelStructureProgramNamedValueType.Name]: The name of the parameter.
//   - [IMLModelStructureProgramNamedValueType.Type]: The type of the parameter.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramNamedValueType
type IMLModelStructureProgramNamedValueType interface {
	objectivec.IObject

	// Topic: Accessing the value type properties

	// The name of the parameter.
	Name() string
	// The type of the parameter.
	Type() IMLModelStructureProgramValueType
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

// The name of the parameter.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramNamedValueType/name
func (m MLModelStructureProgramNamedValueType) Name() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}

// The type of the parameter.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramNamedValueType/type
func (m MLModelStructureProgramNamedValueType) Type() IMLModelStructureProgramValueType {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("type"))
	return MLModelStructureProgramValueTypeFromID(objc.ID(rv))
}
