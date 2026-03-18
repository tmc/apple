// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLModelStructureProgramArgument] class.
var (
	_MLModelStructureProgramArgumentClass     MLModelStructureProgramArgumentClass
	_MLModelStructureProgramArgumentClassOnce sync.Once
)

func getMLModelStructureProgramArgumentClass() MLModelStructureProgramArgumentClass {
	_MLModelStructureProgramArgumentClassOnce.Do(func() {
		_MLModelStructureProgramArgumentClass = MLModelStructureProgramArgumentClass{class: objc.GetClass("MLModelStructureProgramArgument")}
	})
	return _MLModelStructureProgramArgumentClass
}

// GetMLModelStructureProgramArgumentClass returns the class object for MLModelStructureProgramArgument.
func GetMLModelStructureProgramArgumentClass() MLModelStructureProgramArgumentClass {
	return getMLModelStructureProgramArgumentClass()
}

type MLModelStructureProgramArgumentClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLModelStructureProgramArgumentClass) Alloc() MLModelStructureProgramArgument {
	rv := objc.Send[MLModelStructureProgramArgument](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A class representing an argument in the Program.
//
// # Accessing the bindings
//
//   - [MLModelStructureProgramArgument.Bindings]: The array of bindings.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramArgument
type MLModelStructureProgramArgument struct {
	objectivec.Object
}

// MLModelStructureProgramArgumentFromID constructs a [MLModelStructureProgramArgument] from an objc.ID.
//
// A class representing an argument in the Program.
func MLModelStructureProgramArgumentFromID(id objc.ID) MLModelStructureProgramArgument {
	return MLModelStructureProgramArgument{objectivec.Object{ID: id}}
}
// Ensure MLModelStructureProgramArgument implements IMLModelStructureProgramArgument.
var _ IMLModelStructureProgramArgument = MLModelStructureProgramArgument{}

// An interface definition for the [MLModelStructureProgramArgument] class.
//
// # Accessing the bindings
//
//   - [IMLModelStructureProgramArgument.Bindings]: The array of bindings.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramArgument
type IMLModelStructureProgramArgument interface {
	objectivec.IObject

	// Topic: Accessing the bindings

	// The array of bindings.
	Bindings() []MLModelStructureProgramBinding
}

// Init initializes the instance.
func (m MLModelStructureProgramArgument) Init() MLModelStructureProgramArgument {
	rv := objc.Send[MLModelStructureProgramArgument](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLModelStructureProgramArgument) Autorelease() MLModelStructureProgramArgument {
	rv := objc.Send[MLModelStructureProgramArgument](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLModelStructureProgramArgument creates a new MLModelStructureProgramArgument instance.
func NewMLModelStructureProgramArgument() MLModelStructureProgramArgument {
	class := getMLModelStructureProgramArgumentClass()
	rv := objc.Send[MLModelStructureProgramArgument](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The array of bindings.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramArgument/bindings
func (m MLModelStructureProgramArgument) Bindings() []MLModelStructureProgramBinding {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("bindings"))
	return objc.ConvertSlice(rv, func(id objc.ID) MLModelStructureProgramBinding {
		return MLModelStructureProgramBindingFromID(id)
	})
}

