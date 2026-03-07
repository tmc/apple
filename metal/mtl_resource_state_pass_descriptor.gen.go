// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLResourceStatePassDescriptor] class.
var (
	_MTLResourceStatePassDescriptorClass     MTLResourceStatePassDescriptorClass
	_MTLResourceStatePassDescriptorClassOnce sync.Once
)

func getMTLResourceStatePassDescriptorClass() MTLResourceStatePassDescriptorClass {
	_MTLResourceStatePassDescriptorClassOnce.Do(func() {
		_MTLResourceStatePassDescriptorClass = MTLResourceStatePassDescriptorClass{class: objc.GetClass("MTLResourceStatePassDescriptor")}
	})
	return _MTLResourceStatePassDescriptorClass
}

// GetMTLResourceStatePassDescriptorClass returns the class object for MTLResourceStatePassDescriptor.
func GetMTLResourceStatePassDescriptorClass() MTLResourceStatePassDescriptorClass {
	return getMTLResourceStatePassDescriptorClass()
}

type MTLResourceStatePassDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLResourceStatePassDescriptorClass) Alloc() MTLResourceStatePassDescriptor {
	rv := objc.Send[MTLResourceStatePassDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// A configuration for a resource state pass, used to create a resource state
// command encoder.
//
// # Specifying sample buffers for GPU counters
//
//   - [MTLResourceStatePassDescriptor.SampleBufferAttachments]: The array of sample buffers that the resource state pass can access.
//
// See: https://developer.apple.com/documentation/Metal/MTLResourceStatePassDescriptor
type MTLResourceStatePassDescriptor struct {
	objectivec.Object
}

// MTLResourceStatePassDescriptorFromID constructs a [MTLResourceStatePassDescriptor] from an objc.ID.
//
// A configuration for a resource state pass, used to create a resource state
// command encoder.
func MTLResourceStatePassDescriptorFromID(id objc.ID) MTLResourceStatePassDescriptor {
	return MTLResourceStatePassDescriptor{objectivec.Object{id}}
}
// NOTE: MTLResourceStatePassDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [MTLResourceStatePassDescriptor] class.
//
// # Specifying sample buffers for GPU counters
//
//   - [IMTLResourceStatePassDescriptor.SampleBufferAttachments]: The array of sample buffers that the resource state pass can access.
//
// See: https://developer.apple.com/documentation/Metal/MTLResourceStatePassDescriptor
type IMTLResourceStatePassDescriptor interface {
	objectivec.IObject

	// Topic: Specifying sample buffers for GPU counters

	// The array of sample buffers that the resource state pass can access.
	SampleBufferAttachments() IMTLResourceStatePassSampleBufferAttachmentDescriptorArray
}





// Init initializes the instance.
func (r MTLResourceStatePassDescriptor) Init() MTLResourceStatePassDescriptor {
	rv := objc.Send[MTLResourceStatePassDescriptor](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MTLResourceStatePassDescriptor) Autorelease() MTLResourceStatePassDescriptor {
	rv := objc.Send[MTLResourceStatePassDescriptor](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLResourceStatePassDescriptor creates a new MTLResourceStatePassDescriptor instance.
func NewMTLResourceStatePassDescriptor() MTLResourceStatePassDescriptor {
	class := getMTLResourceStatePassDescriptorClass()
	rv := objc.Send[MTLResourceStatePassDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}














// Creates a new resource state pass descriptor.
//
// # Return Value
// 
// A new resource state pass descriptor.
//
// See: https://developer.apple.com/documentation/Metal/MTLResourceStatePassDescriptor/resourceStatePassDescriptor
func (_MTLResourceStatePassDescriptorClass MTLResourceStatePassDescriptorClass) ResourceStatePassDescriptor() MTLResourceStatePassDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_MTLResourceStatePassDescriptorClass.class), objc.Sel("resourceStatePassDescriptor"))
	return MTLResourceStatePassDescriptorFromID(rv)
}








// The array of sample buffers that the resource state pass can access.
//
// See: https://developer.apple.com/documentation/Metal/MTLResourceStatePassDescriptor/sampleBufferAttachments
func (r MTLResourceStatePassDescriptor) SampleBufferAttachments() IMTLResourceStatePassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("sampleBufferAttachments"))
	return MTLResourceStatePassSampleBufferAttachmentDescriptorArrayFromID(objc.ID(rv))
}
























