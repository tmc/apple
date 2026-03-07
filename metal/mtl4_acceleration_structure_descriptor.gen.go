// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MTL4AccelerationStructureDescriptor] class.
var (
	_MTL4AccelerationStructureDescriptorClass     MTL4AccelerationStructureDescriptorClass
	_MTL4AccelerationStructureDescriptorClassOnce sync.Once
)

func getMTL4AccelerationStructureDescriptorClass() MTL4AccelerationStructureDescriptorClass {
	_MTL4AccelerationStructureDescriptorClassOnce.Do(func() {
		_MTL4AccelerationStructureDescriptorClass = MTL4AccelerationStructureDescriptorClass{class: objc.GetClass("MTL4AccelerationStructureDescriptor")}
	})
	return _MTL4AccelerationStructureDescriptorClass
}

// GetMTL4AccelerationStructureDescriptorClass returns the class object for MTL4AccelerationStructureDescriptor.
func GetMTL4AccelerationStructureDescriptorClass() MTL4AccelerationStructureDescriptorClass {
	return getMTL4AccelerationStructureDescriptorClass()
}

type MTL4AccelerationStructureDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTL4AccelerationStructureDescriptorClass) Alloc() MTL4AccelerationStructureDescriptor {
	rv := objc.Send[MTL4AccelerationStructureDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// Base class for Metal 4 acceleration structure descriptors.
//
// # Overview
// 
// Don’t use this class directly. Use one of its subclasses instead.
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureDescriptor
type MTL4AccelerationStructureDescriptor struct {
	MTLAccelerationStructureDescriptor
}

// MTL4AccelerationStructureDescriptorFromID constructs a [MTL4AccelerationStructureDescriptor] from an objc.ID.
//
// Base class for Metal 4 acceleration structure descriptors.
func MTL4AccelerationStructureDescriptorFromID(id objc.ID) MTL4AccelerationStructureDescriptor {
	return MTL4AccelerationStructureDescriptor{MTLAccelerationStructureDescriptor: MTLAccelerationStructureDescriptorFromID(id)}
}
// NOTE: MTL4AccelerationStructureDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [MTL4AccelerationStructureDescriptor] class.
//
// See: https://developer.apple.com/documentation/Metal/MTL4AccelerationStructureDescriptor
type IMTL4AccelerationStructureDescriptor interface {
	IMTLAccelerationStructureDescriptor
}





// Init initializes the instance.
func (m MTL4AccelerationStructureDescriptor) Init() MTL4AccelerationStructureDescriptor {
	rv := objc.Send[MTL4AccelerationStructureDescriptor](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTL4AccelerationStructureDescriptor) Autorelease() MTL4AccelerationStructureDescriptor {
	rv := objc.Send[MTL4AccelerationStructureDescriptor](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4AccelerationStructureDescriptor creates a new MTL4AccelerationStructureDescriptor instance.
func NewMTL4AccelerationStructureDescriptor() MTL4AccelerationStructureDescriptor {
	class := getMTL4AccelerationStructureDescriptorClass()
	rv := objc.Send[MTL4AccelerationStructureDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}










































