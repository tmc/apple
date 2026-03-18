// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLModelStructureProgramValueType] class.
var (
	_MLModelStructureProgramValueTypeClass     MLModelStructureProgramValueTypeClass
	_MLModelStructureProgramValueTypeClassOnce sync.Once
)

func getMLModelStructureProgramValueTypeClass() MLModelStructureProgramValueTypeClass {
	_MLModelStructureProgramValueTypeClassOnce.Do(func() {
		_MLModelStructureProgramValueTypeClass = MLModelStructureProgramValueTypeClass{class: objc.GetClass("MLModelStructureProgramValueType")}
	})
	return _MLModelStructureProgramValueTypeClass
}

// GetMLModelStructureProgramValueTypeClass returns the class object for MLModelStructureProgramValueType.
func GetMLModelStructureProgramValueTypeClass() MLModelStructureProgramValueTypeClass {
	return getMLModelStructureProgramValueTypeClass()
}

type MLModelStructureProgramValueTypeClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLModelStructureProgramValueTypeClass) Alloc() MLModelStructureProgramValueType {
	rv := objc.Send[MLModelStructureProgramValueType](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A class representing the type of a value or a variable in the Program.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramValueType
type MLModelStructureProgramValueType struct {
	objectivec.Object
}

// MLModelStructureProgramValueTypeFromID constructs a [MLModelStructureProgramValueType] from an objc.ID.
//
// A class representing the type of a value or a variable in the Program.
func MLModelStructureProgramValueTypeFromID(id objc.ID) MLModelStructureProgramValueType {
	return MLModelStructureProgramValueType{objectivec.Object{ID: id}}
}
// Ensure MLModelStructureProgramValueType implements IMLModelStructureProgramValueType.
var _ IMLModelStructureProgramValueType = MLModelStructureProgramValueType{}

// An interface definition for the [MLModelStructureProgramValueType] class.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramValueType
type IMLModelStructureProgramValueType interface {
	objectivec.IObject
}

// Init initializes the instance.
func (m MLModelStructureProgramValueType) Init() MLModelStructureProgramValueType {
	rv := objc.Send[MLModelStructureProgramValueType](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLModelStructureProgramValueType) Autorelease() MLModelStructureProgramValueType {
	rv := objc.Send[MLModelStructureProgramValueType](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLModelStructureProgramValueType creates a new MLModelStructureProgramValueType instance.
func NewMLModelStructureProgramValueType() MLModelStructureProgramValueType {
	class := getMLModelStructureProgramValueTypeClass()
	rv := objc.Send[MLModelStructureProgramValueType](objc.ID(class.class), objc.Sel("new"))
	return rv
}

