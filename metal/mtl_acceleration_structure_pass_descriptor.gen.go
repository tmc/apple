// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLAccelerationStructurePassDescriptor] class.
var (
	_MTLAccelerationStructurePassDescriptorClass     MTLAccelerationStructurePassDescriptorClass
	_MTLAccelerationStructurePassDescriptorClassOnce sync.Once
)

func getMTLAccelerationStructurePassDescriptorClass() MTLAccelerationStructurePassDescriptorClass {
	_MTLAccelerationStructurePassDescriptorClassOnce.Do(func() {
		_MTLAccelerationStructurePassDescriptorClass = MTLAccelerationStructurePassDescriptorClass{class: objc.GetClass("MTLAccelerationStructurePassDescriptor")}
	})
	return _MTLAccelerationStructurePassDescriptorClass
}

// GetMTLAccelerationStructurePassDescriptorClass returns the class object for MTLAccelerationStructurePassDescriptor.
func GetMTLAccelerationStructurePassDescriptorClass() MTLAccelerationStructurePassDescriptorClass {
	return getMTLAccelerationStructurePassDescriptorClass()
}

type MTLAccelerationStructurePassDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLAccelerationStructurePassDescriptorClass) Alloc() MTLAccelerationStructurePassDescriptor {
	rv := objc.Send[MTLAccelerationStructurePassDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







//
// # Instance Properties
//
//   - [MTLAccelerationStructurePassDescriptor.SampleBufferAttachments]
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructurePassDescriptor
type MTLAccelerationStructurePassDescriptor struct {
	objectivec.Object
}

// MTLAccelerationStructurePassDescriptorFromID constructs a [MTLAccelerationStructurePassDescriptor] from an objc.ID.
func MTLAccelerationStructurePassDescriptorFromID(id objc.ID) MTLAccelerationStructurePassDescriptor {
	return MTLAccelerationStructurePassDescriptor{objectivec.Object{id}}
}
// NOTE: MTLAccelerationStructurePassDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [MTLAccelerationStructurePassDescriptor] class.
//
// # Instance Properties
//
//   - [IMTLAccelerationStructurePassDescriptor.SampleBufferAttachments]
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructurePassDescriptor
type IMTLAccelerationStructurePassDescriptor interface {
	objectivec.IObject

	// Topic: Instance Properties

	SampleBufferAttachments() IMTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray
}





// Init initializes the instance.
func (a MTLAccelerationStructurePassDescriptor) Init() MTLAccelerationStructurePassDescriptor {
	rv := objc.Send[MTLAccelerationStructurePassDescriptor](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a MTLAccelerationStructurePassDescriptor) Autorelease() MTLAccelerationStructurePassDescriptor {
	rv := objc.Send[MTLAccelerationStructurePassDescriptor](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLAccelerationStructurePassDescriptor creates a new MTLAccelerationStructurePassDescriptor instance.
func NewMTLAccelerationStructurePassDescriptor() MTLAccelerationStructurePassDescriptor {
	class := getMTLAccelerationStructurePassDescriptorClass()
	rv := objc.Send[MTLAccelerationStructurePassDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}














// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructurePassDescriptor/accelerationStructurePassDescriptor
func (_MTLAccelerationStructurePassDescriptorClass MTLAccelerationStructurePassDescriptorClass) AccelerationStructurePassDescriptor() MTLAccelerationStructurePassDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_MTLAccelerationStructurePassDescriptorClass.class), objc.Sel("accelerationStructurePassDescriptor"))
	return MTLAccelerationStructurePassDescriptorFromID(rv)
}








// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructurePassDescriptor/sampleBufferAttachments
func (a MTLAccelerationStructurePassDescriptor) SampleBufferAttachments() IMTLAccelerationStructurePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("sampleBufferAttachments"))
	return MTLAccelerationStructurePassSampleBufferAttachmentDescriptorArrayFromID(objc.ID(rv))
}
























