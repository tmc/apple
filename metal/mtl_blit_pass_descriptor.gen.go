// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLBlitPassDescriptor] class.
var (
	_MTLBlitPassDescriptorClass     MTLBlitPassDescriptorClass
	_MTLBlitPassDescriptorClassOnce sync.Once
)

func getMTLBlitPassDescriptorClass() MTLBlitPassDescriptorClass {
	_MTLBlitPassDescriptorClassOnce.Do(func() {
		_MTLBlitPassDescriptorClass = MTLBlitPassDescriptorClass{class: objc.GetClass("MTLBlitPassDescriptor")}
	})
	return _MTLBlitPassDescriptorClass
}

// GetMTLBlitPassDescriptorClass returns the class object for MTLBlitPassDescriptor.
func GetMTLBlitPassDescriptorClass() MTLBlitPassDescriptorClass {
	return getMTLBlitPassDescriptorClass()
}

type MTLBlitPassDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLBlitPassDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLBlitPassDescriptorClass) Alloc() MTLBlitPassDescriptor {
	rv := objc.Send[MTLBlitPassDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A configuration you create to customize a blit command encoder, which
// affects the runtime behavior of the blit pass you encode with it.
//
// # Overview
// 
// You can customize an encoder for a blit pass by creating and configuring an
// [MTLBlitPassDescriptor] instance and passing it to
// [BlitCommandEncoderWithDescriptor].
//
// # Configuring sample buffer attachment descriptors for a blit pass
//
//   - [MTLBlitPassDescriptor.SampleBufferAttachments]: An array of counter sample buffer attachments that you configure for a blit pass.
//
// See: https://developer.apple.com/documentation/Metal/MTLBlitPassDescriptor
type MTLBlitPassDescriptor struct {
	objectivec.Object
}

// MTLBlitPassDescriptorFromID constructs a [MTLBlitPassDescriptor] from an objc.ID.
//
// A configuration you create to customize a blit command encoder, which
// affects the runtime behavior of the blit pass you encode with it.
func MTLBlitPassDescriptorFromID(id objc.ID) MTLBlitPassDescriptor {
	return MTLBlitPassDescriptor{objectivec.Object{ID: id}}
}
// NOTE: MTLBlitPassDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLBlitPassDescriptor] class.
//
// # Configuring sample buffer attachment descriptors for a blit pass
//
//   - [IMTLBlitPassDescriptor.SampleBufferAttachments]: An array of counter sample buffer attachments that you configure for a blit pass.
//
// See: https://developer.apple.com/documentation/Metal/MTLBlitPassDescriptor
type IMTLBlitPassDescriptor interface {
	objectivec.IObject

	// Topic: Configuring sample buffer attachment descriptors for a blit pass

	// An array of counter sample buffer attachments that you configure for a blit pass.
	SampleBufferAttachments() IMTLBlitPassSampleBufferAttachmentDescriptorArray
}

// Init initializes the instance.
func (b MTLBlitPassDescriptor) Init() MTLBlitPassDescriptor {
	rv := objc.Send[MTLBlitPassDescriptor](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b MTLBlitPassDescriptor) Autorelease() MTLBlitPassDescriptor {
	rv := objc.Send[MTLBlitPassDescriptor](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLBlitPassDescriptor creates a new MTLBlitPassDescriptor instance.
func NewMTLBlitPassDescriptor() MTLBlitPassDescriptor {
	class := getMTLBlitPassDescriptorClass()
	rv := objc.Send[MTLBlitPassDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new blit pass descriptor with a default configuration.
//
// See: https://developer.apple.com/documentation/Metal/MTLBlitPassDescriptor/blitPassDescriptor
func (_MTLBlitPassDescriptorClass MTLBlitPassDescriptorClass) BlitPassDescriptor() MTLBlitPassDescriptor {
	rv := objc.Send[objc.ID](objc.ID(_MTLBlitPassDescriptorClass.class), objc.Sel("blitPassDescriptor"))
	return MTLBlitPassDescriptorFromID(rv)
}

// An array of counter sample buffer attachments that you configure for a blit
// pass.
//
// # Discussion
// 
// See [Sampling GPU data into counter sample buffers] for more context about
// configuring this property. That article is one of a series of articles in
// [GPU counters and counter sample buffers].
//
// [GPU counters and counter sample buffers]: https://developer.apple.com/documentation/Metal/gpu-counters-and-counter-sample-buffers
// [Sampling GPU data into counter sample buffers]: https://developer.apple.com/documentation/Metal/sampling-gpu-data-into-counter-sample-buffers
//
// See: https://developer.apple.com/documentation/Metal/MTLBlitPassDescriptor/sampleBufferAttachments
func (b MTLBlitPassDescriptor) SampleBufferAttachments() IMTLBlitPassSampleBufferAttachmentDescriptorArray {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("sampleBufferAttachments"))
	return MTLBlitPassSampleBufferAttachmentDescriptorArrayFromID(objc.ID(rv))
}

