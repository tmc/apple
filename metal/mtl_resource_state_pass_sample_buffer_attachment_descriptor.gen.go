// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLResourceStatePassSampleBufferAttachmentDescriptor] class.
var (
	_MTLResourceStatePassSampleBufferAttachmentDescriptorClass     MTLResourceStatePassSampleBufferAttachmentDescriptorClass
	_MTLResourceStatePassSampleBufferAttachmentDescriptorClassOnce sync.Once
)

func getMTLResourceStatePassSampleBufferAttachmentDescriptorClass() MTLResourceStatePassSampleBufferAttachmentDescriptorClass {
	_MTLResourceStatePassSampleBufferAttachmentDescriptorClassOnce.Do(func() {
		_MTLResourceStatePassSampleBufferAttachmentDescriptorClass = MTLResourceStatePassSampleBufferAttachmentDescriptorClass{class: objc.GetClass("MTLResourceStatePassSampleBufferAttachmentDescriptor")}
	})
	return _MTLResourceStatePassSampleBufferAttachmentDescriptorClass
}

// GetMTLResourceStatePassSampleBufferAttachmentDescriptorClass returns the class object for MTLResourceStatePassSampleBufferAttachmentDescriptor.
func GetMTLResourceStatePassSampleBufferAttachmentDescriptorClass() MTLResourceStatePassSampleBufferAttachmentDescriptorClass {
	return getMTLResourceStatePassSampleBufferAttachmentDescriptorClass()
}

type MTLResourceStatePassSampleBufferAttachmentDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLResourceStatePassSampleBufferAttachmentDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLResourceStatePassSampleBufferAttachmentDescriptorClass) Alloc() MTLResourceStatePassSampleBufferAttachmentDescriptor {
	rv := objc.Send[MTLResourceStatePassSampleBufferAttachmentDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A description of where to store GPU counter information at the start and
// end of a resource state pass.
//
// # Configuring the sample buffer attachment
//
//   - [MTLResourceStatePassSampleBufferAttachmentDescriptor.SampleBuffer]: A specialized memory buffer that the GPU uses to store its counter data during the resource state pass.
//   - [MTLResourceStatePassSampleBufferAttachmentDescriptor.SetSampleBuffer]
//   - [MTLResourceStatePassSampleBufferAttachmentDescriptor.StartOfEncoderSampleIndex]: The index the Metal device object should use to store GPU counters when starting the resource state pass.
//   - [MTLResourceStatePassSampleBufferAttachmentDescriptor.SetStartOfEncoderSampleIndex]
//   - [MTLResourceStatePassSampleBufferAttachmentDescriptor.EndOfEncoderSampleIndex]: The index the Metal device object should use to store GPU counters when ending the resource state pass.
//   - [MTLResourceStatePassSampleBufferAttachmentDescriptor.SetEndOfEncoderSampleIndex]
//
// See: https://developer.apple.com/documentation/Metal/MTLResourceStatePassSampleBufferAttachmentDescriptor
type MTLResourceStatePassSampleBufferAttachmentDescriptor struct {
	objectivec.Object
}

// MTLResourceStatePassSampleBufferAttachmentDescriptorFromID constructs a [MTLResourceStatePassSampleBufferAttachmentDescriptor] from an objc.ID.
//
// A description of where to store GPU counter information at the start and
// end of a resource state pass.
func MTLResourceStatePassSampleBufferAttachmentDescriptorFromID(id objc.ID) MTLResourceStatePassSampleBufferAttachmentDescriptor {
	return MTLResourceStatePassSampleBufferAttachmentDescriptor{objectivec.Object{ID: id}}
}
// NOTE: MTLResourceStatePassSampleBufferAttachmentDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLResourceStatePassSampleBufferAttachmentDescriptor] class.
//
// # Configuring the sample buffer attachment
//
//   - [IMTLResourceStatePassSampleBufferAttachmentDescriptor.SampleBuffer]: A specialized memory buffer that the GPU uses to store its counter data during the resource state pass.
//   - [IMTLResourceStatePassSampleBufferAttachmentDescriptor.SetSampleBuffer]
//   - [IMTLResourceStatePassSampleBufferAttachmentDescriptor.StartOfEncoderSampleIndex]: The index the Metal device object should use to store GPU counters when starting the resource state pass.
//   - [IMTLResourceStatePassSampleBufferAttachmentDescriptor.SetStartOfEncoderSampleIndex]
//   - [IMTLResourceStatePassSampleBufferAttachmentDescriptor.EndOfEncoderSampleIndex]: The index the Metal device object should use to store GPU counters when ending the resource state pass.
//   - [IMTLResourceStatePassSampleBufferAttachmentDescriptor.SetEndOfEncoderSampleIndex]
//
// See: https://developer.apple.com/documentation/Metal/MTLResourceStatePassSampleBufferAttachmentDescriptor
type IMTLResourceStatePassSampleBufferAttachmentDescriptor interface {
	objectivec.IObject

	// Topic: Configuring the sample buffer attachment

	// A specialized memory buffer that the GPU uses to store its counter data during the resource state pass.
	SampleBuffer() MTLCounterSampleBuffer
	SetSampleBuffer(value MTLCounterSampleBuffer)
	// The index the Metal device object should use to store GPU counters when starting the resource state pass.
	StartOfEncoderSampleIndex() uint
	SetStartOfEncoderSampleIndex(value uint)
	// The index the Metal device object should use to store GPU counters when ending the resource state pass.
	EndOfEncoderSampleIndex() uint
	SetEndOfEncoderSampleIndex(value uint)
}

// Init initializes the instance.
func (r MTLResourceStatePassSampleBufferAttachmentDescriptor) Init() MTLResourceStatePassSampleBufferAttachmentDescriptor {
	rv := objc.Send[MTLResourceStatePassSampleBufferAttachmentDescriptor](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MTLResourceStatePassSampleBufferAttachmentDescriptor) Autorelease() MTLResourceStatePassSampleBufferAttachmentDescriptor {
	rv := objc.Send[MTLResourceStatePassSampleBufferAttachmentDescriptor](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLResourceStatePassSampleBufferAttachmentDescriptor creates a new MTLResourceStatePassSampleBufferAttachmentDescriptor instance.
func NewMTLResourceStatePassSampleBufferAttachmentDescriptor() MTLResourceStatePassSampleBufferAttachmentDescriptor {
	class := getMTLResourceStatePassSampleBufferAttachmentDescriptorClass()
	rv := objc.Send[MTLResourceStatePassSampleBufferAttachmentDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A specialized memory buffer that the GPU uses to store its counter data
// during the resource state pass.
//
// # Discussion
// 
// The property defaults to `nil`, which means the GPU doesn’t save any GPU
// counter information during the resource state pass. For more information,
// see [Creating a counter sample buffer to store a GPU’s counter data
// during a pass] and [Sampling GPU data into counter sample buffers].
//
// [Creating a counter sample buffer to store a GPU’s counter data during a pass]: https://developer.apple.com/documentation/Metal/creating-a-counter-sample-buffer-to-store-a-gpus-counter-data-during-a-pass
// [Sampling GPU data into counter sample buffers]: https://developer.apple.com/documentation/Metal/sampling-gpu-data-into-counter-sample-buffers
//
// See: https://developer.apple.com/documentation/Metal/MTLResourceStatePassSampleBufferAttachmentDescriptor/sampleBuffer
func (r MTLResourceStatePassSampleBufferAttachmentDescriptor) SampleBuffer() MTLCounterSampleBuffer {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("sampleBuffer"))
	return MTLCounterSampleBufferObjectFromID(rv)
}
func (r MTLResourceStatePassSampleBufferAttachmentDescriptor) SetSampleBuffer(value MTLCounterSampleBuffer) {
	objc.Send[struct{}](r.ID, objc.Sel("setSampleBuffer:"), value)
}
// The index the Metal device object should use to store GPU counters when
// starting the resource state pass.
//
// # Discussion
// 
// Specify [MTLCounterDontSample] if you don’t want to sample GPU counters
// at the start of the resource state pass. Otherwise, specify an index within
// the sample buffer where you want the GPU to write the sample data.
// 
// On devices that don’t support [CounterSamplingPointAtStageBoundary] you
// need to set the value to [MTLCounterDontSample].
//
// [MTLCounterDontSample]: https://developer.apple.com/documentation/Metal/MTLCounterDontSample
//
// See: https://developer.apple.com/documentation/Metal/MTLResourceStatePassSampleBufferAttachmentDescriptor/startOfEncoderSampleIndex
func (r MTLResourceStatePassSampleBufferAttachmentDescriptor) StartOfEncoderSampleIndex() uint {
	rv := objc.Send[uint](r.ID, objc.Sel("startOfEncoderSampleIndex"))
	return rv
}
func (r MTLResourceStatePassSampleBufferAttachmentDescriptor) SetStartOfEncoderSampleIndex(value uint) {
	objc.Send[struct{}](r.ID, objc.Sel("setStartOfEncoderSampleIndex:"), value)
}
// The index the Metal device object should use to store GPU counters when
// ending the resource state pass.
//
// # Discussion
// 
// Specify [MTLCounterDontSample] if you don’t want to sample GPU counters
// at the end of the resource state pass. Otherwise, specify an index within
// the sample buffer where you want the GPU to write the sample data.
// 
// On devices that don’t support [CounterSamplingPointAtStageBoundary] you
// need to set the value to [MTLCounterDontSample].
//
// [MTLCounterDontSample]: https://developer.apple.com/documentation/Metal/MTLCounterDontSample
//
// See: https://developer.apple.com/documentation/Metal/MTLResourceStatePassSampleBufferAttachmentDescriptor/endOfEncoderSampleIndex
func (r MTLResourceStatePassSampleBufferAttachmentDescriptor) EndOfEncoderSampleIndex() uint {
	rv := objc.Send[uint](r.ID, objc.Sel("endOfEncoderSampleIndex"))
	return rv
}
func (r MTLResourceStatePassSampleBufferAttachmentDescriptor) SetEndOfEncoderSampleIndex(value uint) {
	objc.Send[struct{}](r.ID, objc.Sel("setEndOfEncoderSampleIndex:"), value)
}

