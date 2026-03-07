// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLBlitPassSampleBufferAttachmentDescriptor] class.
var (
	_MTLBlitPassSampleBufferAttachmentDescriptorClass     MTLBlitPassSampleBufferAttachmentDescriptorClass
	_MTLBlitPassSampleBufferAttachmentDescriptorClassOnce sync.Once
)

func getMTLBlitPassSampleBufferAttachmentDescriptorClass() MTLBlitPassSampleBufferAttachmentDescriptorClass {
	_MTLBlitPassSampleBufferAttachmentDescriptorClassOnce.Do(func() {
		_MTLBlitPassSampleBufferAttachmentDescriptorClass = MTLBlitPassSampleBufferAttachmentDescriptorClass{class: objc.GetClass("MTLBlitPassSampleBufferAttachmentDescriptor")}
	})
	return _MTLBlitPassSampleBufferAttachmentDescriptorClass
}

// GetMTLBlitPassSampleBufferAttachmentDescriptorClass returns the class object for MTLBlitPassSampleBufferAttachmentDescriptor.
func GetMTLBlitPassSampleBufferAttachmentDescriptorClass() MTLBlitPassSampleBufferAttachmentDescriptorClass {
	return getMTLBlitPassSampleBufferAttachmentDescriptorClass()
}

type MTLBlitPassSampleBufferAttachmentDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLBlitPassSampleBufferAttachmentDescriptorClass) Alloc() MTLBlitPassSampleBufferAttachmentDescriptor {
	rv := objc.Send[MTLBlitPassSampleBufferAttachmentDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// A configuration that instructs the GPU where to store counter data from the
// beginning and end of a blit pass.
//
// # Overview
// 
// See [Sampling GPU data into counter sample buffers] for more context about
// configuring instances of this type. That article is one of a series of
// articles in [GPU counters and counter sample buffers].
//
// [GPU counters and counter sample buffers]: https://developer.apple.com/documentation/Metal/gpu-counters-and-counter-sample-buffers
// [Sampling GPU data into counter sample buffers]: https://developer.apple.com/documentation/Metal/sampling-gpu-data-into-counter-sample-buffers
//
// # Configuring the sample buffer attachment
//
//   - [MTLBlitPassSampleBufferAttachmentDescriptor.SampleBuffer]: A specialized memory buffer that the GPU uses to store its counter data during the blit pass.
//   - [MTLBlitPassSampleBufferAttachmentDescriptor.SetSampleBuffer]
//   - [MTLBlitPassSampleBufferAttachmentDescriptor.StartOfEncoderSampleIndex]: An index within a counter sample buffer that tells the GPU where to store counter data from the start of a blit pass.
//   - [MTLBlitPassSampleBufferAttachmentDescriptor.SetStartOfEncoderSampleIndex]
//   - [MTLBlitPassSampleBufferAttachmentDescriptor.EndOfEncoderSampleIndex]: An index within a counter sample buffer that tells the GPU where to store counter data from the end of a blit pass.
//   - [MTLBlitPassSampleBufferAttachmentDescriptor.SetEndOfEncoderSampleIndex]
//
// See: https://developer.apple.com/documentation/Metal/MTLBlitPassSampleBufferAttachmentDescriptor
type MTLBlitPassSampleBufferAttachmentDescriptor struct {
	objectivec.Object
}

// MTLBlitPassSampleBufferAttachmentDescriptorFromID constructs a [MTLBlitPassSampleBufferAttachmentDescriptor] from an objc.ID.
//
// A configuration that instructs the GPU where to store counter data from the
// beginning and end of a blit pass.
func MTLBlitPassSampleBufferAttachmentDescriptorFromID(id objc.ID) MTLBlitPassSampleBufferAttachmentDescriptor {
	return MTLBlitPassSampleBufferAttachmentDescriptor{objectivec.Object{id}}
}
// NOTE: MTLBlitPassSampleBufferAttachmentDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [MTLBlitPassSampleBufferAttachmentDescriptor] class.
//
// # Configuring the sample buffer attachment
//
//   - [IMTLBlitPassSampleBufferAttachmentDescriptor.SampleBuffer]: A specialized memory buffer that the GPU uses to store its counter data during the blit pass.
//   - [IMTLBlitPassSampleBufferAttachmentDescriptor.SetSampleBuffer]
//   - [IMTLBlitPassSampleBufferAttachmentDescriptor.StartOfEncoderSampleIndex]: An index within a counter sample buffer that tells the GPU where to store counter data from the start of a blit pass.
//   - [IMTLBlitPassSampleBufferAttachmentDescriptor.SetStartOfEncoderSampleIndex]
//   - [IMTLBlitPassSampleBufferAttachmentDescriptor.EndOfEncoderSampleIndex]: An index within a counter sample buffer that tells the GPU where to store counter data from the end of a blit pass.
//   - [IMTLBlitPassSampleBufferAttachmentDescriptor.SetEndOfEncoderSampleIndex]
//
// See: https://developer.apple.com/documentation/Metal/MTLBlitPassSampleBufferAttachmentDescriptor
type IMTLBlitPassSampleBufferAttachmentDescriptor interface {
	objectivec.IObject

	// Topic: Configuring the sample buffer attachment

	// A specialized memory buffer that the GPU uses to store its counter data during the blit pass.
	SampleBuffer() MTLCounterSampleBuffer
	SetSampleBuffer(value MTLCounterSampleBuffer)
	// An index within a counter sample buffer that tells the GPU where to store counter data from the start of a blit pass.
	StartOfEncoderSampleIndex() uint
	SetStartOfEncoderSampleIndex(value uint)
	// An index within a counter sample buffer that tells the GPU where to store counter data from the end of a blit pass.
	EndOfEncoderSampleIndex() uint
	SetEndOfEncoderSampleIndex(value uint)
}





// Init initializes the instance.
func (b MTLBlitPassSampleBufferAttachmentDescriptor) Init() MTLBlitPassSampleBufferAttachmentDescriptor {
	rv := objc.Send[MTLBlitPassSampleBufferAttachmentDescriptor](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b MTLBlitPassSampleBufferAttachmentDescriptor) Autorelease() MTLBlitPassSampleBufferAttachmentDescriptor {
	rv := objc.Send[MTLBlitPassSampleBufferAttachmentDescriptor](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLBlitPassSampleBufferAttachmentDescriptor creates a new MTLBlitPassSampleBufferAttachmentDescriptor instance.
func NewMTLBlitPassSampleBufferAttachmentDescriptor() MTLBlitPassSampleBufferAttachmentDescriptor {
	class := getMTLBlitPassSampleBufferAttachmentDescriptorClass()
	rv := objc.Send[MTLBlitPassSampleBufferAttachmentDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}





















// A specialized memory buffer that the GPU uses to store its counter data
// during the blit pass.
//
// # Discussion
// 
// The property defaults to `nil`, which means the GPU doesn’t save any GPU
// counter information during the blit pass. For more information, see
// [Creating a counter sample buffer to store a GPU’s counter data during a
// pass] and [Sampling GPU data into counter sample buffers].
//
// [Creating a counter sample buffer to store a GPU’s counter data during a pass]: https://developer.apple.com/documentation/Metal/creating-a-counter-sample-buffer-to-store-a-gpus-counter-data-during-a-pass
// [Sampling GPU data into counter sample buffers]: https://developer.apple.com/documentation/Metal/sampling-gpu-data-into-counter-sample-buffers
//
// See: https://developer.apple.com/documentation/Metal/MTLBlitPassSampleBufferAttachmentDescriptor/sampleBuffer
func (b MTLBlitPassSampleBufferAttachmentDescriptor) SampleBuffer() MTLCounterSampleBuffer {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("sampleBuffer"))
	return MTLCounterSampleBufferObjectFromID(rv)
}
func (b MTLBlitPassSampleBufferAttachmentDescriptor) SetSampleBuffer(value MTLCounterSampleBuffer) {
	objc.Send[struct{}](b.ID, objc.Sel("setSampleBuffer:"), value)
}



// An index within a counter sample buffer that tells the GPU where to store
// counter data from the start of a blit pass.
//
// # Discussion
// 
// This property indicates where the GPU stores the counter data within an
// [MTLCounterSampleBuffer] instance that it samples at the beginning of a
// blit pass.
// 
// You can tell the GPU to skip sampling at the start of the blit pass by
// assigning [MTLCounterDontSample] to this property.
//
// [MTLCounterDontSample]: https://developer.apple.com/documentation/Metal/MTLCounterDontSample
//
// See: https://developer.apple.com/documentation/Metal/MTLBlitPassSampleBufferAttachmentDescriptor/startOfEncoderSampleIndex
func (b MTLBlitPassSampleBufferAttachmentDescriptor) StartOfEncoderSampleIndex() uint {
	rv := objc.Send[uint](b.ID, objc.Sel("startOfEncoderSampleIndex"))
	return rv
}
func (b MTLBlitPassSampleBufferAttachmentDescriptor) SetStartOfEncoderSampleIndex(value uint) {
	objc.Send[struct{}](b.ID, objc.Sel("setStartOfEncoderSampleIndex:"), value)
}



// An index within a counter sample buffer that tells the GPU where to store
// counter data from the end of a blit pass.
//
// # Discussion
// 
// This property indicates where the GPU stores the counter data within an
// [MTLCounterSampleBuffer] instance that it samples at the end of a blit
// pass.
// 
// You can tell the GPU to skip sampling at the end of the blit pass by
// assigning [MTLCounterDontSample] to this property.
//
// [MTLCounterDontSample]: https://developer.apple.com/documentation/Metal/MTLCounterDontSample
//
// See: https://developer.apple.com/documentation/Metal/MTLBlitPassSampleBufferAttachmentDescriptor/endOfEncoderSampleIndex
func (b MTLBlitPassSampleBufferAttachmentDescriptor) EndOfEncoderSampleIndex() uint {
	rv := objc.Send[uint](b.ID, objc.Sel("endOfEncoderSampleIndex"))
	return rv
}
func (b MTLBlitPassSampleBufferAttachmentDescriptor) SetEndOfEncoderSampleIndex(value uint) {
	objc.Send[struct{}](b.ID, objc.Sel("setEndOfEncoderSampleIndex:"), value)
}
























