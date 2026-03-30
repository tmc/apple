// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLModelStructureProgramBinding] class.
var (
	_MLModelStructureProgramBindingClass     MLModelStructureProgramBindingClass
	_MLModelStructureProgramBindingClassOnce sync.Once
)

func getMLModelStructureProgramBindingClass() MLModelStructureProgramBindingClass {
	_MLModelStructureProgramBindingClassOnce.Do(func() {
		_MLModelStructureProgramBindingClass = MLModelStructureProgramBindingClass{class: objc.GetClass("MLModelStructureProgramBinding")}
	})
	return _MLModelStructureProgramBindingClass
}

// GetMLModelStructureProgramBindingClass returns the class object for MLModelStructureProgramBinding.
func GetMLModelStructureProgramBindingClass() MLModelStructureProgramBindingClass {
	return getMLModelStructureProgramBindingClass()
}

type MLModelStructureProgramBindingClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLModelStructureProgramBindingClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLModelStructureProgramBindingClass) Alloc() MLModelStructureProgramBinding {
	rv := objc.Send[MLModelStructureProgramBinding](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLModelStructureProgramBinding.InitWithName]
//   - [MLModelStructureProgramBinding.InitWithValue]
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramBinding
type MLModelStructureProgramBinding struct {
	objectivec.Object
}

// MLModelStructureProgramBindingFromID constructs a [MLModelStructureProgramBinding] from an objc.ID.
func MLModelStructureProgramBindingFromID(id objc.ID) MLModelStructureProgramBinding {
	return MLModelStructureProgramBinding{objectivec.Object{ID: id}}
}

// Ensure MLModelStructureProgramBinding implements IMLModelStructureProgramBinding.
var _ IMLModelStructureProgramBinding = MLModelStructureProgramBinding{}

// An interface definition for the [MLModelStructureProgramBinding] class.
//
// # Methods
//
//   - [IMLModelStructureProgramBinding.InitWithName]
//   - [IMLModelStructureProgramBinding.InitWithValue]
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramBinding
type IMLModelStructureProgramBinding interface {
	objectivec.IObject

	// Topic: Methods

	InitWithName(name objectivec.IObject) MLModelStructureProgramBinding
	InitWithValue(value objectivec.IObject) MLModelStructureProgramBinding
}

// Init initializes the instance.
func (m MLModelStructureProgramBinding) Init() MLModelStructureProgramBinding {
	rv := objc.Send[MLModelStructureProgramBinding](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLModelStructureProgramBinding) Autorelease() MLModelStructureProgramBinding {
	rv := objc.Send[MLModelStructureProgramBinding](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLModelStructureProgramBinding creates a new MLModelStructureProgramBinding instance.
func NewMLModelStructureProgramBinding() MLModelStructureProgramBinding {
	class := getMLModelStructureProgramBindingClass()
	rv := objc.Send[MLModelStructureProgramBinding](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramBinding/initWithName:
func NewModelStructureProgramBindingWithName(name objectivec.IObject) MLModelStructureProgramBinding {
	instance := getMLModelStructureProgramBindingClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:"), name)
	return MLModelStructureProgramBindingFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramBinding/initWithValue:
func NewModelStructureProgramBindingWithValue(value objectivec.IObject) MLModelStructureProgramBinding {
	instance := getMLModelStructureProgramBindingClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithValue:"), value)
	return MLModelStructureProgramBindingFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramBinding/initWithName:
func (m MLModelStructureProgramBinding) InitWithName(name objectivec.IObject) MLModelStructureProgramBinding {
	rv := objc.Send[MLModelStructureProgramBinding](m.ID, objc.Sel("initWithName:"), name)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramBinding/initWithValue:
func (m MLModelStructureProgramBinding) InitWithValue(value objectivec.IObject) MLModelStructureProgramBinding {
	rv := objc.Send[MLModelStructureProgramBinding](m.ID, objc.Sel("initWithValue:"), value)
	return rv
}
