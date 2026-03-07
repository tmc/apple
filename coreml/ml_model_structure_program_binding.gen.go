// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
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

// Alloc allocates memory for a new instance of the class.
func (mc MLModelStructureProgramBindingClass) Alloc() MLModelStructureProgramBinding {
	rv := objc.Send[MLModelStructureProgramBinding](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// A class representing a binding in the Program
//
// # Overview
// 
// A Binding is either a previously defined name of a variable or a constant
// value in the Program.
//
// # Accessing the program binding properties
//
//   - [MLModelStructureProgramBinding.Name]: The name of the variable in the Program.
//   - [MLModelStructureProgramBinding.Value]: The compile time constant value in the Program.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramBinding
type MLModelStructureProgramBinding struct {
	objectivec.Object
}

// MLModelStructureProgramBindingFromID constructs a [MLModelStructureProgramBinding] from an objc.ID.
//
// A class representing a binding in the Program
func MLModelStructureProgramBindingFromID(id objc.ID) MLModelStructureProgramBinding {
	return MLModelStructureProgramBinding{objectivec.Object{id}}
}
// Ensure MLModelStructureProgramBinding implements IMLModelStructureProgramBinding.
var _ IMLModelStructureProgramBinding = MLModelStructureProgramBinding{}





// An interface definition for the [MLModelStructureProgramBinding] class.
//
// # Accessing the program binding properties
//
//   - [IMLModelStructureProgramBinding.Name]: The name of the variable in the Program.
//   - [IMLModelStructureProgramBinding.Value]: The compile time constant value in the Program.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramBinding
type IMLModelStructureProgramBinding interface {
	objectivec.IObject

	// Topic: Accessing the program binding properties

	// The name of the variable in the Program.
	Name() string
	// The compile time constant value in the Program.
	Value() IMLModelStructureProgramValue
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





















// The name of the variable in the Program.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramBinding/name
func (m MLModelStructureProgramBinding) Name() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}



// The compile time constant value in the Program.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramBinding/value
func (m MLModelStructureProgramBinding) Value() IMLModelStructureProgramValue {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("value"))
	return MLModelStructureProgramValueFromID(objc.ID(rv))
}

















