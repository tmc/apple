// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTL4BinaryFunctionDescriptor] class.
var (
	_MTL4BinaryFunctionDescriptorClass     MTL4BinaryFunctionDescriptorClass
	_MTL4BinaryFunctionDescriptorClassOnce sync.Once
)

func getMTL4BinaryFunctionDescriptorClass() MTL4BinaryFunctionDescriptorClass {
	_MTL4BinaryFunctionDescriptorClassOnce.Do(func() {
		_MTL4BinaryFunctionDescriptorClass = MTL4BinaryFunctionDescriptorClass{class: objc.GetClass("MTL4BinaryFunctionDescriptor")}
	})
	return _MTL4BinaryFunctionDescriptorClass
}

// GetMTL4BinaryFunctionDescriptorClass returns the class object for MTL4BinaryFunctionDescriptor.
func GetMTL4BinaryFunctionDescriptorClass() MTL4BinaryFunctionDescriptorClass {
	return getMTL4BinaryFunctionDescriptorClass()
}

type MTL4BinaryFunctionDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTL4BinaryFunctionDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTL4BinaryFunctionDescriptorClass) Alloc() MTL4BinaryFunctionDescriptor {
	rv := objc.Send[MTL4BinaryFunctionDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// Base interface for other function-derived interfaces.
//
// # Instance Properties
//
//   - [MTL4BinaryFunctionDescriptor.FunctionDescriptor]: Provides the function descriptor corresponding to the function to compile into a binary function.
//   - [MTL4BinaryFunctionDescriptor.SetFunctionDescriptor]
//   - [MTL4BinaryFunctionDescriptor.Name]: Associates a string that uniquely identifies a binary function.
//   - [MTL4BinaryFunctionDescriptor.SetName]
//   - [MTL4BinaryFunctionDescriptor.Options]: Configure the options to use at binary function creation time.
//   - [MTL4BinaryFunctionDescriptor.SetOptions]
//
// See: https://developer.apple.com/documentation/Metal/MTL4BinaryFunctionDescriptor
type MTL4BinaryFunctionDescriptor struct {
	objectivec.Object
}

// MTL4BinaryFunctionDescriptorFromID constructs a [MTL4BinaryFunctionDescriptor] from an objc.ID.
//
// Base interface for other function-derived interfaces.
func MTL4BinaryFunctionDescriptorFromID(id objc.ID) MTL4BinaryFunctionDescriptor {
	return MTL4BinaryFunctionDescriptor{objectivec.Object{ID: id}}
}

// NOTE: MTL4BinaryFunctionDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTL4BinaryFunctionDescriptor] class.
//
// # Instance Properties
//
//   - [IMTL4BinaryFunctionDescriptor.FunctionDescriptor]: Provides the function descriptor corresponding to the function to compile into a binary function.
//   - [IMTL4BinaryFunctionDescriptor.SetFunctionDescriptor]
//   - [IMTL4BinaryFunctionDescriptor.Name]: Associates a string that uniquely identifies a binary function.
//   - [IMTL4BinaryFunctionDescriptor.SetName]
//   - [IMTL4BinaryFunctionDescriptor.Options]: Configure the options to use at binary function creation time.
//   - [IMTL4BinaryFunctionDescriptor.SetOptions]
//
// See: https://developer.apple.com/documentation/Metal/MTL4BinaryFunctionDescriptor
type IMTL4BinaryFunctionDescriptor interface {
	objectivec.IObject

	// Topic: Instance Properties

	// Provides the function descriptor corresponding to the function to compile into a binary function.
	FunctionDescriptor() IMTL4FunctionDescriptor
	SetFunctionDescriptor(value IMTL4FunctionDescriptor)
	// Associates a string that uniquely identifies a binary function.
	Name() string
	SetName(value string)
	// Configure the options to use at binary function creation time.
	Options() MTL4BinaryFunctionOptions
	SetOptions(value MTL4BinaryFunctionOptions)
}

// Init initializes the instance.
func (m MTL4BinaryFunctionDescriptor) Init() MTL4BinaryFunctionDescriptor {
	rv := objc.Send[MTL4BinaryFunctionDescriptor](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTL4BinaryFunctionDescriptor) Autorelease() MTL4BinaryFunctionDescriptor {
	rv := objc.Send[MTL4BinaryFunctionDescriptor](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4BinaryFunctionDescriptor creates a new MTL4BinaryFunctionDescriptor instance.
func NewMTL4BinaryFunctionDescriptor() MTL4BinaryFunctionDescriptor {
	class := getMTL4BinaryFunctionDescriptorClass()
	rv := objc.Send[MTL4BinaryFunctionDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Provides the function descriptor corresponding to the function to compile
// into a binary function.
//
// See: https://developer.apple.com/documentation/Metal/MTL4BinaryFunctionDescriptor/functionDescriptor
func (m MTL4BinaryFunctionDescriptor) FunctionDescriptor() IMTL4FunctionDescriptor {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("functionDescriptor"))
	return MTL4FunctionDescriptorFromID(objc.ID(rv))
}
func (m MTL4BinaryFunctionDescriptor) SetFunctionDescriptor(value IMTL4FunctionDescriptor) {
	objc.Send[struct{}](m.ID, objc.Sel("setFunctionDescriptor:"), value)
}

// Associates a string that uniquely identifies a binary function.
//
// # Discussion
//
// You can use this property to look up a corresponding binary function by
// name in a [MTL4Archive] instance.
//
// See: https://developer.apple.com/documentation/Metal/MTL4BinaryFunctionDescriptor/name
func (m MTL4BinaryFunctionDescriptor) Name() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (m MTL4BinaryFunctionDescriptor) SetName(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setName:"), objc.String(value))
}

// Configure the options to use at binary function creation time.
//
// See: https://developer.apple.com/documentation/Metal/MTL4BinaryFunctionDescriptor/options
func (m MTL4BinaryFunctionDescriptor) Options() MTL4BinaryFunctionOptions {
	rv := objc.Send[MTL4BinaryFunctionOptions](m.ID, objc.Sel("options"))
	return MTL4BinaryFunctionOptions(rv)
}
func (m MTL4BinaryFunctionDescriptor) SetOptions(value MTL4BinaryFunctionOptions) {
	objc.Send[struct{}](m.ID, objc.Sel("setOptions:"), value)
}
