// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

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

// Class returns the underlying Objective-C class pointer.
func (mc MLModelStructureProgramArgumentClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLModelStructureProgramArgumentClass) Alloc() MLModelStructureProgramArgument {
	rv := objc.Send[MLModelStructureProgramArgument](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLModelStructureProgramArgument.InitWithBindings]
//   - [MLModelStructureProgramArgument.InitWithMILArguments]
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramArgument
type MLModelStructureProgramArgument struct {
	objectivec.Object
}

// MLModelStructureProgramArgumentFromID constructs a [MLModelStructureProgramArgument] from an objc.ID.
func MLModelStructureProgramArgumentFromID(id objc.ID) MLModelStructureProgramArgument {
	return MLModelStructureProgramArgument{objectivec.Object{ID: id}}
}

// Ensure MLModelStructureProgramArgument implements IMLModelStructureProgramArgument.
var _ IMLModelStructureProgramArgument = MLModelStructureProgramArgument{}

// An interface definition for the [MLModelStructureProgramArgument] class.
//
// # Methods
//
//   - [IMLModelStructureProgramArgument.InitWithBindings]
//   - [IMLModelStructureProgramArgument.InitWithMILArguments]
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramArgument
type IMLModelStructureProgramArgument interface {
	objectivec.IObject

	// Topic: Methods

	InitWithBindings(bindings objectivec.IObject) MLModelStructureProgramArgument
	InitWithMILArguments(mILArguments unsafe.Pointer) MLModelStructureProgramArgument
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

// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramArgument/initWithBindings:
func NewModelStructureProgramArgumentWithBindings(bindings objectivec.IObject) MLModelStructureProgramArgument {
	instance := getMLModelStructureProgramArgumentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBindings:"), bindings)
	return MLModelStructureProgramArgumentFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramArgument/initWithMILArguments:
func NewModelStructureProgramArgumentWithMILArguments(mILArguments unsafe.Pointer) MLModelStructureProgramArgument {
	instance := getMLModelStructureProgramArgumentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMILArguments:"), mILArguments)
	return MLModelStructureProgramArgumentFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramArgument/initWithBindings:
func (m MLModelStructureProgramArgument) InitWithBindings(bindings objectivec.IObject) MLModelStructureProgramArgument {
	rv := objc.Send[MLModelStructureProgramArgument](m.ID, objc.Sel("initWithBindings:"), bindings)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelStructureProgramArgument/initWithMILArguments:
func (m MLModelStructureProgramArgument) InitWithMILArguments(mILArguments unsafe.Pointer) MLModelStructureProgramArgument {
	rv := objc.Send[MLModelStructureProgramArgument](m.ID, objc.Sel("initWithMILArguments:"), mILArguments)
	return rv
}
