// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLVisibleFunctionTableDescriptor] class.
var (
	_MTLVisibleFunctionTableDescriptorClass     MTLVisibleFunctionTableDescriptorClass
	_MTLVisibleFunctionTableDescriptorClassOnce sync.Once
)

func getMTLVisibleFunctionTableDescriptorClass() MTLVisibleFunctionTableDescriptorClass {
	_MTLVisibleFunctionTableDescriptorClassOnce.Do(func() {
		_MTLVisibleFunctionTableDescriptorClass = MTLVisibleFunctionTableDescriptorClass{class: objc.GetClass("MTLVisibleFunctionTableDescriptor")}
	})
	return _MTLVisibleFunctionTableDescriptorClass
}

// GetMTLVisibleFunctionTableDescriptorClass returns the class object for MTLVisibleFunctionTableDescriptor.
func GetMTLVisibleFunctionTableDescriptorClass() MTLVisibleFunctionTableDescriptorClass {
	return getMTLVisibleFunctionTableDescriptorClass()
}

type MTLVisibleFunctionTableDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLVisibleFunctionTableDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLVisibleFunctionTableDescriptorClass) Alloc() MTLVisibleFunctionTableDescriptor {
	rv := objc.Send[MTLVisibleFunctionTableDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A specification of how to create a visible function table.
//
// # Configuring the function table
//
//   - [MTLVisibleFunctionTableDescriptor.FunctionCount]: The number of entries in the function table.
//   - [MTLVisibleFunctionTableDescriptor.SetFunctionCount]
//
// See: https://developer.apple.com/documentation/Metal/MTLVisibleFunctionTableDescriptor
type MTLVisibleFunctionTableDescriptor struct {
	objectivec.Object
}

// MTLVisibleFunctionTableDescriptorFromID constructs a [MTLVisibleFunctionTableDescriptor] from an objc.ID.
//
// A specification of how to create a visible function table.
func MTLVisibleFunctionTableDescriptorFromID(id objc.ID) MTLVisibleFunctionTableDescriptor {
	return MTLVisibleFunctionTableDescriptor{objectivec.Object{ID: id}}
}
// NOTE: MTLVisibleFunctionTableDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLVisibleFunctionTableDescriptor] class.
//
// # Configuring the function table
//
//   - [IMTLVisibleFunctionTableDescriptor.FunctionCount]: The number of entries in the function table.
//   - [IMTLVisibleFunctionTableDescriptor.SetFunctionCount]
//
// See: https://developer.apple.com/documentation/Metal/MTLVisibleFunctionTableDescriptor
type IMTLVisibleFunctionTableDescriptor interface {
	objectivec.IObject

	// Topic: Configuring the function table

	// The number of entries in the function table.
	FunctionCount() uint
	SetFunctionCount(value uint)
}

// Init initializes the instance.
func (v MTLVisibleFunctionTableDescriptor) Init() MTLVisibleFunctionTableDescriptor {
	rv := objc.Send[MTLVisibleFunctionTableDescriptor](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v MTLVisibleFunctionTableDescriptor) Autorelease() MTLVisibleFunctionTableDescriptor {
	rv := objc.Send[MTLVisibleFunctionTableDescriptor](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLVisibleFunctionTableDescriptor creates a new MTLVisibleFunctionTableDescriptor instance.
func NewMTLVisibleFunctionTableDescriptor() MTLVisibleFunctionTableDescriptor {
	class := getMTLVisibleFunctionTableDescriptorClass()
	rv := objc.Send[MTLVisibleFunctionTableDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a default visible function table descriptor.
//
// See: https://developer.apple.com/documentation/Metal/MTLVisibleFunctionTableDescriptor/visibleFunctionTableDescriptor
func (_MTLVisibleFunctionTableDescriptorClass MTLVisibleFunctionTableDescriptorClass) VisibleFunctionTableDescriptor() MTLVisibleFunctionTableDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_MTLVisibleFunctionTableDescriptorClass.class), objc.Sel("visibleFunctionTableDescriptor"))
	return MTLVisibleFunctionTableDescriptorFromID(rv)
}

// The number of entries in the function table.
//
// See: https://developer.apple.com/documentation/Metal/MTLVisibleFunctionTableDescriptor/functionCount
func (v MTLVisibleFunctionTableDescriptor) FunctionCount() uint {
	rv := objc.Send[uint](v.ID, objc.Sel("functionCount"))
	return rv
}
func (v MTLVisibleFunctionTableDescriptor) SetFunctionCount(value uint) {
	objc.Send[struct{}](v.ID, objc.Sel("setFunctionCount:"), value)
}

