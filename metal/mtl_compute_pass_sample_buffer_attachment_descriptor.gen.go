// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLComputePassSampleBufferAttachmentDescriptor] class.
var (
	_MTLComputePassSampleBufferAttachmentDescriptorClass     MTLComputePassSampleBufferAttachmentDescriptorClass
	_MTLComputePassSampleBufferAttachmentDescriptorClassOnce sync.Once
)

func getMTLComputePassSampleBufferAttachmentDescriptorClass() MTLComputePassSampleBufferAttachmentDescriptorClass {
	_MTLComputePassSampleBufferAttachmentDescriptorClassOnce.Do(func() {
		_MTLComputePassSampleBufferAttachmentDescriptorClass = MTLComputePassSampleBufferAttachmentDescriptorClass{class: objc.GetClass("MTLComputePassSampleBufferAttachmentDescriptor")}
	})
	return _MTLComputePassSampleBufferAttachmentDescriptorClass
}

// GetMTLComputePassSampleBufferAttachmentDescriptorClass returns the class object for MTLComputePassSampleBufferAttachmentDescriptor.
func GetMTLComputePassSampleBufferAttachmentDescriptorClass() MTLComputePassSampleBufferAttachmentDescriptorClass {
	return getMTLComputePassSampleBufferAttachmentDescriptorClass()
}

type MTLComputePassSampleBufferAttachmentDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLComputePassSampleBufferAttachmentDescriptorClass) Alloc() MTLComputePassSampleBufferAttachmentDescriptor {
	rv := objc.Send[MTLComputePassSampleBufferAttachmentDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A configuration that instructs the GPU where to store counter data from the
// beginning and end of a compute pass.
//
// # Overview
// 
// For more context about configuring sample buffer attachments for compute
// passes, see [Sampling GPU data into counter sample buffers]. That article
// is one of a series in [GPU counters and counter sample buffers] about
// sampling Metal hardware counters for performance measurement.
//
// [GPU counters and counter sample buffers]: https://developer.apple.com/documentation/Metal/gpu-counters-and-counter-sample-buffers
// [Sampling GPU data into counter sample buffers]: https://developer.apple.com/documentation/Metal/sampling-gpu-data-into-counter-sample-buffers
//
// # Configuring the sample buffer attachment
//
//   - [MTLComputePassSampleBufferAttachmentDescriptor.SampleBuffer]: A specialized memory buffer that the GPU uses to store its counter data during a compute pass.
//   - [MTLComputePassSampleBufferAttachmentDescriptor.SetSampleBuffer]
//   - [MTLComputePassSampleBufferAttachmentDescriptor.StartOfEncoderSampleIndex]: An index within a counter sample buffer that tells the GPU where to store counter data from the start of a compute pass.
//   - [MTLComputePassSampleBufferAttachmentDescriptor.SetStartOfEncoderSampleIndex]
//   - [MTLComputePassSampleBufferAttachmentDescriptor.EndOfEncoderSampleIndex]: An index within a counter sample buffer that tells the GPU where to store counter data from the end of a compute pass.
//   - [MTLComputePassSampleBufferAttachmentDescriptor.SetEndOfEncoderSampleIndex]
//
// See: https://developer.apple.com/documentation/Metal/MTLComputePassSampleBufferAttachmentDescriptor
type MTLComputePassSampleBufferAttachmentDescriptor struct {
	objectivec.Object
}

// MTLComputePassSampleBufferAttachmentDescriptorFromID constructs a [MTLComputePassSampleBufferAttachmentDescriptor] from an objc.ID.
//
// A configuration that instructs the GPU where to store counter data from the
// beginning and end of a compute pass.
func MTLComputePassSampleBufferAttachmentDescriptorFromID(id objc.ID) MTLComputePassSampleBufferAttachmentDescriptor {
	return MTLComputePassSampleBufferAttachmentDescriptor{objectivec.Object{ID: id}}
}
// NOTE: MTLComputePassSampleBufferAttachmentDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLComputePassSampleBufferAttachmentDescriptor] class.
//
// # Configuring the sample buffer attachment
//
//   - [IMTLComputePassSampleBufferAttachmentDescriptor.SampleBuffer]: A specialized memory buffer that the GPU uses to store its counter data during a compute pass.
//   - [IMTLComputePassSampleBufferAttachmentDescriptor.SetSampleBuffer]
//   - [IMTLComputePassSampleBufferAttachmentDescriptor.StartOfEncoderSampleIndex]: An index within a counter sample buffer that tells the GPU where to store counter data from the start of a compute pass.
//   - [IMTLComputePassSampleBufferAttachmentDescriptor.SetStartOfEncoderSampleIndex]
//   - [IMTLComputePassSampleBufferAttachmentDescriptor.EndOfEncoderSampleIndex]: An index within a counter sample buffer that tells the GPU where to store counter data from the end of a compute pass.
//   - [IMTLComputePassSampleBufferAttachmentDescriptor.SetEndOfEncoderSampleIndex]
//
// See: https://developer.apple.com/documentation/Metal/MTLComputePassSampleBufferAttachmentDescriptor
type IMTLComputePassSampleBufferAttachmentDescriptor interface {
	objectivec.IObject

	// Topic: Configuring the sample buffer attachment

	// A specialized memory buffer that the GPU uses to store its counter data during a compute pass.
	SampleBuffer() MTLCounterSampleBuffer
	SetSampleBuffer(value MTLCounterSampleBuffer)
	// An index within a counter sample buffer that tells the GPU where to store counter data from the start of a compute pass.
	StartOfEncoderSampleIndex() uint
	SetStartOfEncoderSampleIndex(value uint)
	// An index within a counter sample buffer that tells the GPU where to store counter data from the end of a compute pass.
	EndOfEncoderSampleIndex() uint
	SetEndOfEncoderSampleIndex(value uint)
}

// Init initializes the instance.
func (c MTLComputePassSampleBufferAttachmentDescriptor) Init() MTLComputePassSampleBufferAttachmentDescriptor {
	rv := objc.Send[MTLComputePassSampleBufferAttachmentDescriptor](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MTLComputePassSampleBufferAttachmentDescriptor) Autorelease() MTLComputePassSampleBufferAttachmentDescriptor {
	rv := objc.Send[MTLComputePassSampleBufferAttachmentDescriptor](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLComputePassSampleBufferAttachmentDescriptor creates a new MTLComputePassSampleBufferAttachmentDescriptor instance.
func NewMTLComputePassSampleBufferAttachmentDescriptor() MTLComputePassSampleBufferAttachmentDescriptor {
	class := getMTLComputePassSampleBufferAttachmentDescriptorClass()
	rv := objc.Send[MTLComputePassSampleBufferAttachmentDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A specialized memory buffer that the GPU uses to store its counter data
// during a compute pass.
//
// # Discussion
// 
// The property defaults to `nil`, which means the GPU doesn’t save any GPU
// counter information during the compute pass. For more information, see
// [Creating a counter sample buffer to store a GPU’s counter data during a
// pass] and [Sampling GPU data into counter sample buffers].
//
// [Creating a counter sample buffer to store a GPU’s counter data during a pass]: https://developer.apple.com/documentation/Metal/creating-a-counter-sample-buffer-to-store-a-gpus-counter-data-during-a-pass
// [Sampling GPU data into counter sample buffers]: https://developer.apple.com/documentation/Metal/sampling-gpu-data-into-counter-sample-buffers
//
// See: https://developer.apple.com/documentation/Metal/MTLComputePassSampleBufferAttachmentDescriptor/sampleBuffer
func (c MTLComputePassSampleBufferAttachmentDescriptor) SampleBuffer() MTLCounterSampleBuffer {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("sampleBuffer"))
	return MTLCounterSampleBufferObjectFromID(rv)
}
func (c MTLComputePassSampleBufferAttachmentDescriptor) SetSampleBuffer(value MTLCounterSampleBuffer) {
	objc.Send[struct{}](c.ID, objc.Sel("setSampleBuffer:"), value)
}
// An index within a counter sample buffer that tells the GPU where to store
// counter data from the start of a compute pass.
//
// # Discussion
// 
// This property indicates where the GPU stores the counter data within an
// [MTLCounterSampleBuffer] instance that it samples at the beginning of a
// compute pass.
// 
// You can tell the GPU to skip sampling at the start of the compute pass by
// assigning [MTLCounterDontSample] to this property.
//
// [MTLCounterDontSample]: https://developer.apple.com/documentation/Metal/MTLCounterDontSample
//
// See: https://developer.apple.com/documentation/Metal/MTLComputePassSampleBufferAttachmentDescriptor/startOfEncoderSampleIndex
func (c MTLComputePassSampleBufferAttachmentDescriptor) StartOfEncoderSampleIndex() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("startOfEncoderSampleIndex"))
	return rv
}
func (c MTLComputePassSampleBufferAttachmentDescriptor) SetStartOfEncoderSampleIndex(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setStartOfEncoderSampleIndex:"), value)
}
// An index within a counter sample buffer that tells the GPU where to store
// counter data from the end of a compute pass.
//
// # Discussion
// 
// This property indicates where the GPU stores the counter data within an
// [MTLCounterSampleBuffer] instance that it samples at the end of a compute
// pass.
// 
// You can tell the GPU to skip sampling at the end of the compute pass by
// assigning [MTLCounterDontSample] to this property.
//
// [MTLCounterDontSample]: https://developer.apple.com/documentation/Metal/MTLCounterDontSample
//
// See: https://developer.apple.com/documentation/Metal/MTLComputePassSampleBufferAttachmentDescriptor/endOfEncoderSampleIndex
func (c MTLComputePassSampleBufferAttachmentDescriptor) EndOfEncoderSampleIndex() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("endOfEncoderSampleIndex"))
	return rv
}
func (c MTLComputePassSampleBufferAttachmentDescriptor) SetEndOfEncoderSampleIndex(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setEndOfEncoderSampleIndex:"), value)
}

