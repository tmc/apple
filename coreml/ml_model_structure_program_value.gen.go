// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLModelStructureProgramValue] class.
var (
	_MLModelStructureProgramValueClass     MLModelStructureProgramValueClass
	_MLModelStructureProgramValueClassOnce sync.Once
)

func getMLModelStructureProgramValueClass() MLModelStructureProgramValueClass {
	_MLModelStructureProgramValueClassOnce.Do(func() {
		_MLModelStructureProgramValueClass = MLModelStructureProgramValueClass{class: objc.GetClass("MLModelStructureProgramValue")}
	})
	return _MLModelStructureProgramValueClass
}

// GetMLModelStructureProgramValueClass returns the class object for MLModelStructureProgramValue.
func GetMLModelStructureProgramValueClass() MLModelStructureProgramValueClass {
	return getMLModelStructureProgramValueClass()
}

type MLModelStructureProgramValueClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLModelStructureProgramValueClass) Alloc() MLModelStructureProgramValue {
	rv := objc.Send[MLModelStructureProgramValue](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// A class representing a constant value in the Program.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramValue
type MLModelStructureProgramValue struct {
	objectivec.Object
}

// MLModelStructureProgramValueFromID constructs a [MLModelStructureProgramValue] from an objc.ID.
//
// A class representing a constant value in the Program.
func MLModelStructureProgramValueFromID(id objc.ID) MLModelStructureProgramValue {
	return MLModelStructureProgramValue{objectivec.Object{id}}
}
// Ensure MLModelStructureProgramValue implements IMLModelStructureProgramValue.
var _ IMLModelStructureProgramValue = MLModelStructureProgramValue{}





// An interface definition for the [MLModelStructureProgramValue] class.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramValue
type IMLModelStructureProgramValue interface {
	objectivec.IObject
}





// Init initializes the instance.
func (m MLModelStructureProgramValue) Init() MLModelStructureProgramValue {
	rv := objc.Send[MLModelStructureProgramValue](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLModelStructureProgramValue) Autorelease() MLModelStructureProgramValue {
	rv := objc.Send[MLModelStructureProgramValue](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLModelStructureProgramValue creates a new MLModelStructureProgramValue instance.
func NewMLModelStructureProgramValue() MLModelStructureProgramValue {
	class := getMLModelStructureProgramValueClass()
	rv := objc.Send[MLModelStructureProgramValue](objc.ID(class.class), objc.Sel("new"))
	return rv
}



































