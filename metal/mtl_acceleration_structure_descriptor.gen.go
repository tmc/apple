// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLAccelerationStructureDescriptor] class.
var (
	_MTLAccelerationStructureDescriptorClass     MTLAccelerationStructureDescriptorClass
	_MTLAccelerationStructureDescriptorClassOnce sync.Once
)

func getMTLAccelerationStructureDescriptorClass() MTLAccelerationStructureDescriptorClass {
	_MTLAccelerationStructureDescriptorClassOnce.Do(func() {
		_MTLAccelerationStructureDescriptorClass = MTLAccelerationStructureDescriptorClass{class: objc.GetClass("MTLAccelerationStructureDescriptor")}
	})
	return _MTLAccelerationStructureDescriptorClass
}

// GetMTLAccelerationStructureDescriptorClass returns the class object for MTLAccelerationStructureDescriptor.
func GetMTLAccelerationStructureDescriptorClass() MTLAccelerationStructureDescriptorClass {
	return getMTLAccelerationStructureDescriptorClass()
}

type MTLAccelerationStructureDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLAccelerationStructureDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLAccelerationStructureDescriptorClass) Alloc() MTLAccelerationStructureDescriptor {
	rv := objc.Send[MTLAccelerationStructureDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A base class for classes that define the configuration for a new
// acceleration structure.
//
// # Overview
//
// This is the base class for other acceleration structure descriptors.
// Don’t use this class directly. Use one of the derived classes instead, as
// [MTLAccelerationStructure] describes.
//
// # Specifying usage options
//
//   - [MTLAccelerationStructureDescriptor.Usage]: The options that describe how you intend to use the acceleration structure.
//   - [MTLAccelerationStructureDescriptor.SetUsage]
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureDescriptor
type MTLAccelerationStructureDescriptor struct {
	objectivec.Object
}

// MTLAccelerationStructureDescriptorFromID constructs a [MTLAccelerationStructureDescriptor] from an objc.ID.
//
// A base class for classes that define the configuration for a new
// acceleration structure.
func MTLAccelerationStructureDescriptorFromID(id objc.ID) MTLAccelerationStructureDescriptor {
	return MTLAccelerationStructureDescriptor{objectivec.Object{ID: id}}
}

// NOTE: MTLAccelerationStructureDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLAccelerationStructureDescriptor] class.
//
// # Specifying usage options
//
//   - [IMTLAccelerationStructureDescriptor.Usage]: The options that describe how you intend to use the acceleration structure.
//   - [IMTLAccelerationStructureDescriptor.SetUsage]
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureDescriptor
type IMTLAccelerationStructureDescriptor interface {
	objectivec.IObject

	// Topic: Specifying usage options

	// The options that describe how you intend to use the acceleration structure.
	Usage() MTLAccelerationStructureUsage
	SetUsage(value MTLAccelerationStructureUsage)
}

// Init initializes the instance.
func (a MTLAccelerationStructureDescriptor) Init() MTLAccelerationStructureDescriptor {
	rv := objc.Send[MTLAccelerationStructureDescriptor](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a MTLAccelerationStructureDescriptor) Autorelease() MTLAccelerationStructureDescriptor {
	rv := objc.Send[MTLAccelerationStructureDescriptor](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLAccelerationStructureDescriptor creates a new MTLAccelerationStructureDescriptor instance.
func NewMTLAccelerationStructureDescriptor() MTLAccelerationStructureDescriptor {
	class := getMTLAccelerationStructureDescriptorClass()
	rv := objc.Send[MTLAccelerationStructureDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The options that describe how you intend to use the acceleration structure.
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureDescriptor/usage
func (a MTLAccelerationStructureDescriptor) Usage() MTLAccelerationStructureUsage {
	rv := objc.Send[MTLAccelerationStructureUsage](a.ID, objc.Sel("usage"))
	return MTLAccelerationStructureUsage(rv)
}
func (a MTLAccelerationStructureDescriptor) SetUsage(value MTLAccelerationStructureUsage) {
	objc.Send[struct{}](a.ID, objc.Sel("setUsage:"), value)
}
