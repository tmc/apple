// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MTLIntersectionFunctionDescriptor] class.
var (
	_MTLIntersectionFunctionDescriptorClass     MTLIntersectionFunctionDescriptorClass
	_MTLIntersectionFunctionDescriptorClassOnce sync.Once
)

func getMTLIntersectionFunctionDescriptorClass() MTLIntersectionFunctionDescriptorClass {
	_MTLIntersectionFunctionDescriptorClassOnce.Do(func() {
		_MTLIntersectionFunctionDescriptorClass = MTLIntersectionFunctionDescriptorClass{class: objc.GetClass("MTLIntersectionFunctionDescriptor")}
	})
	return _MTLIntersectionFunctionDescriptorClass
}

// GetMTLIntersectionFunctionDescriptorClass returns the class object for MTLIntersectionFunctionDescriptor.
func GetMTLIntersectionFunctionDescriptorClass() MTLIntersectionFunctionDescriptorClass {
	return getMTLIntersectionFunctionDescriptorClass()
}

type MTLIntersectionFunctionDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLIntersectionFunctionDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLIntersectionFunctionDescriptorClass) Alloc() MTLIntersectionFunctionDescriptor {
	rv := objc.Send[MTLIntersectionFunctionDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A description of an intersection function that performs an intersection
// test.
//
// # Overview
//
// This class doesn’t add any additional API over its parent class.
//
// See: https://developer.apple.com/documentation/Metal/MTLIntersectionFunctionDescriptor
type MTLIntersectionFunctionDescriptor struct {
	MTLFunctionDescriptor
}

// MTLIntersectionFunctionDescriptorFromID constructs a [MTLIntersectionFunctionDescriptor] from an objc.ID.
//
// A description of an intersection function that performs an intersection
// test.
func MTLIntersectionFunctionDescriptorFromID(id objc.ID) MTLIntersectionFunctionDescriptor {
	return MTLIntersectionFunctionDescriptor{MTLFunctionDescriptor: MTLFunctionDescriptorFromID(id)}
}

// NOTE: MTLIntersectionFunctionDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLIntersectionFunctionDescriptor] class.
//
// See: https://developer.apple.com/documentation/Metal/MTLIntersectionFunctionDescriptor
type IMTLIntersectionFunctionDescriptor interface {
	IMTLFunctionDescriptor
}

// Init initializes the instance.
func (i MTLIntersectionFunctionDescriptor) Init() MTLIntersectionFunctionDescriptor {
	rv := objc.Send[MTLIntersectionFunctionDescriptor](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MTLIntersectionFunctionDescriptor) Autorelease() MTLIntersectionFunctionDescriptor {
	rv := objc.Send[MTLIntersectionFunctionDescriptor](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLIntersectionFunctionDescriptor creates a new MTLIntersectionFunctionDescriptor instance.
func NewMTLIntersectionFunctionDescriptor() MTLIntersectionFunctionDescriptor {
	class := getMTLIntersectionFunctionDescriptorClass()
	rv := objc.Send[MTLIntersectionFunctionDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}
