// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLRenderPassSampleBufferAttachmentDescriptor] class.
var (
	_MTLRenderPassSampleBufferAttachmentDescriptorClass     MTLRenderPassSampleBufferAttachmentDescriptorClass
	_MTLRenderPassSampleBufferAttachmentDescriptorClassOnce sync.Once
)

func getMTLRenderPassSampleBufferAttachmentDescriptorClass() MTLRenderPassSampleBufferAttachmentDescriptorClass {
	_MTLRenderPassSampleBufferAttachmentDescriptorClassOnce.Do(func() {
		_MTLRenderPassSampleBufferAttachmentDescriptorClass = MTLRenderPassSampleBufferAttachmentDescriptorClass{class: objc.GetClass("MTLRenderPassSampleBufferAttachmentDescriptor")}
	})
	return _MTLRenderPassSampleBufferAttachmentDescriptorClass
}

// GetMTLRenderPassSampleBufferAttachmentDescriptorClass returns the class object for MTLRenderPassSampleBufferAttachmentDescriptor.
func GetMTLRenderPassSampleBufferAttachmentDescriptorClass() MTLRenderPassSampleBufferAttachmentDescriptorClass {
	return getMTLRenderPassSampleBufferAttachmentDescriptorClass()
}

type MTLRenderPassSampleBufferAttachmentDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLRenderPassSampleBufferAttachmentDescriptorClass) Alloc() MTLRenderPassSampleBufferAttachmentDescriptor {
	rv := objc.Send[MTLRenderPassSampleBufferAttachmentDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A description of where to store GPU counter information at the start and
// end of a render pass.
//
// # Configuring the sample buffer attachment
//
//   - [MTLRenderPassSampleBufferAttachmentDescriptor.SampleBuffer]: A specialized memory buffer that the GPU uses to store its counter data during the render pass.
//   - [MTLRenderPassSampleBufferAttachmentDescriptor.SetSampleBuffer]
//   - [MTLRenderPassSampleBufferAttachmentDescriptor.StartOfVertexSampleIndex]: The index the Metal device object should use to store GPU counters when starting the render pass’s vertex stage.
//   - [MTLRenderPassSampleBufferAttachmentDescriptor.SetStartOfVertexSampleIndex]
//   - [MTLRenderPassSampleBufferAttachmentDescriptor.EndOfVertexSampleIndex]: The index the Metal device object should use to store GPU counters when ending the render pass’s vertex stage.
//   - [MTLRenderPassSampleBufferAttachmentDescriptor.SetEndOfVertexSampleIndex]
//   - [MTLRenderPassSampleBufferAttachmentDescriptor.StartOfFragmentSampleIndex]: The index the Metal device object should use to store GPU counters when starting the render pass’s fragment stage.
//   - [MTLRenderPassSampleBufferAttachmentDescriptor.SetStartOfFragmentSampleIndex]
//   - [MTLRenderPassSampleBufferAttachmentDescriptor.EndOfFragmentSampleIndex]: The index the Metal device object should use to store GPU counters when ending the render pass’s fragment stage.
//   - [MTLRenderPassSampleBufferAttachmentDescriptor.SetEndOfFragmentSampleIndex]
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPassSampleBufferAttachmentDescriptor
type MTLRenderPassSampleBufferAttachmentDescriptor struct {
	objectivec.Object
}

// MTLRenderPassSampleBufferAttachmentDescriptorFromID constructs a [MTLRenderPassSampleBufferAttachmentDescriptor] from an objc.ID.
//
// A description of where to store GPU counter information at the start and
// end of a render pass.
func MTLRenderPassSampleBufferAttachmentDescriptorFromID(id objc.ID) MTLRenderPassSampleBufferAttachmentDescriptor {
	return MTLRenderPassSampleBufferAttachmentDescriptor{objectivec.Object{ID: id}}
}
// NOTE: MTLRenderPassSampleBufferAttachmentDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLRenderPassSampleBufferAttachmentDescriptor] class.
//
// # Configuring the sample buffer attachment
//
//   - [IMTLRenderPassSampleBufferAttachmentDescriptor.SampleBuffer]: A specialized memory buffer that the GPU uses to store its counter data during the render pass.
//   - [IMTLRenderPassSampleBufferAttachmentDescriptor.SetSampleBuffer]
//   - [IMTLRenderPassSampleBufferAttachmentDescriptor.StartOfVertexSampleIndex]: The index the Metal device object should use to store GPU counters when starting the render pass’s vertex stage.
//   - [IMTLRenderPassSampleBufferAttachmentDescriptor.SetStartOfVertexSampleIndex]
//   - [IMTLRenderPassSampleBufferAttachmentDescriptor.EndOfVertexSampleIndex]: The index the Metal device object should use to store GPU counters when ending the render pass’s vertex stage.
//   - [IMTLRenderPassSampleBufferAttachmentDescriptor.SetEndOfVertexSampleIndex]
//   - [IMTLRenderPassSampleBufferAttachmentDescriptor.StartOfFragmentSampleIndex]: The index the Metal device object should use to store GPU counters when starting the render pass’s fragment stage.
//   - [IMTLRenderPassSampleBufferAttachmentDescriptor.SetStartOfFragmentSampleIndex]
//   - [IMTLRenderPassSampleBufferAttachmentDescriptor.EndOfFragmentSampleIndex]: The index the Metal device object should use to store GPU counters when ending the render pass’s fragment stage.
//   - [IMTLRenderPassSampleBufferAttachmentDescriptor.SetEndOfFragmentSampleIndex]
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPassSampleBufferAttachmentDescriptor
type IMTLRenderPassSampleBufferAttachmentDescriptor interface {
	objectivec.IObject

	// Topic: Configuring the sample buffer attachment

	// A specialized memory buffer that the GPU uses to store its counter data during the render pass.
	SampleBuffer() MTLCounterSampleBuffer
	SetSampleBuffer(value MTLCounterSampleBuffer)
	// The index the Metal device object should use to store GPU counters when starting the render pass’s vertex stage.
	StartOfVertexSampleIndex() uint
	SetStartOfVertexSampleIndex(value uint)
	// The index the Metal device object should use to store GPU counters when ending the render pass’s vertex stage.
	EndOfVertexSampleIndex() uint
	SetEndOfVertexSampleIndex(value uint)
	// The index the Metal device object should use to store GPU counters when starting the render pass’s fragment stage.
	StartOfFragmentSampleIndex() uint
	SetStartOfFragmentSampleIndex(value uint)
	// The index the Metal device object should use to store GPU counters when ending the render pass’s fragment stage.
	EndOfFragmentSampleIndex() uint
	SetEndOfFragmentSampleIndex(value uint)
}

// Init initializes the instance.
func (r MTLRenderPassSampleBufferAttachmentDescriptor) Init() MTLRenderPassSampleBufferAttachmentDescriptor {
	rv := objc.Send[MTLRenderPassSampleBufferAttachmentDescriptor](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MTLRenderPassSampleBufferAttachmentDescriptor) Autorelease() MTLRenderPassSampleBufferAttachmentDescriptor {
	rv := objc.Send[MTLRenderPassSampleBufferAttachmentDescriptor](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLRenderPassSampleBufferAttachmentDescriptor creates a new MTLRenderPassSampleBufferAttachmentDescriptor instance.
func NewMTLRenderPassSampleBufferAttachmentDescriptor() MTLRenderPassSampleBufferAttachmentDescriptor {
	class := getMTLRenderPassSampleBufferAttachmentDescriptorClass()
	rv := objc.Send[MTLRenderPassSampleBufferAttachmentDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A specialized memory buffer that the GPU uses to store its counter data
// during the render pass.
//
// # Discussion
// 
// The property defaults to `nil`, which means the GPU doesn’t save any GPU
// counter information during the render pass. For more information, see
// [Creating a counter sample buffer to store a GPU’s counter data during a
// pass] and [Sampling GPU data into counter sample buffers].
//
// [Creating a counter sample buffer to store a GPU’s counter data during a pass]: https://developer.apple.com/documentation/Metal/creating-a-counter-sample-buffer-to-store-a-gpus-counter-data-during-a-pass
// [Sampling GPU data into counter sample buffers]: https://developer.apple.com/documentation/Metal/sampling-gpu-data-into-counter-sample-buffers
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPassSampleBufferAttachmentDescriptor/sampleBuffer
func (r MTLRenderPassSampleBufferAttachmentDescriptor) SampleBuffer() MTLCounterSampleBuffer {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("sampleBuffer"))
	return MTLCounterSampleBufferObjectFromID(rv)
}
func (r MTLRenderPassSampleBufferAttachmentDescriptor) SetSampleBuffer(value MTLCounterSampleBuffer) {
	objc.Send[struct{}](r.ID, objc.Sel("setSampleBuffer:"), value)
}
// The index the Metal device object should use to store GPU counters when
// starting the render pass’s vertex stage.
//
// # Discussion
// 
// Specify [MTLCounterDontSample] if you don’t want to sample GPU counters
// at the start of the vertex stage. Otherwise, specify an index within the
// sample buffer where you want the GPU to write the sample data.
// 
// On devices that don’t support [CounterSamplingPointAtStageBoundary] you
// need to set the value to [MTLCounterDontSample].
//
// [MTLCounterDontSample]: https://developer.apple.com/documentation/Metal/MTLCounterDontSample
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPassSampleBufferAttachmentDescriptor/startOfVertexSampleIndex
func (r MTLRenderPassSampleBufferAttachmentDescriptor) StartOfVertexSampleIndex() uint {
	rv := objc.Send[uint](r.ID, objc.Sel("startOfVertexSampleIndex"))
	return rv
}
func (r MTLRenderPassSampleBufferAttachmentDescriptor) SetStartOfVertexSampleIndex(value uint) {
	objc.Send[struct{}](r.ID, objc.Sel("setStartOfVertexSampleIndex:"), value)
}
// The index the Metal device object should use to store GPU counters when
// ending the render pass’s vertex stage.
//
// # Discussion
// 
// Specify [MTLCounterDontSample] if you don’t want to sample GPU counters
// at the end of the vertex stage. Otherwise, specify an index within the
// sample buffer where you want the GPU to write the sample data.
// 
// On devices that don’t support [CounterSamplingPointAtStageBoundary] you
// need to set the value to [MTLCounterDontSample].
//
// [MTLCounterDontSample]: https://developer.apple.com/documentation/Metal/MTLCounterDontSample
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPassSampleBufferAttachmentDescriptor/endOfVertexSampleIndex
func (r MTLRenderPassSampleBufferAttachmentDescriptor) EndOfVertexSampleIndex() uint {
	rv := objc.Send[uint](r.ID, objc.Sel("endOfVertexSampleIndex"))
	return rv
}
func (r MTLRenderPassSampleBufferAttachmentDescriptor) SetEndOfVertexSampleIndex(value uint) {
	objc.Send[struct{}](r.ID, objc.Sel("setEndOfVertexSampleIndex:"), value)
}
// The index the Metal device object should use to store GPU counters when
// starting the render pass’s fragment stage.
//
// # Discussion
// 
// Specify [MTLCounterDontSample] if you don’t want to sample GPU counters
// at the start of the fragment stage. Otherwise, specify an index within the
// sample buffer where you want the GPU to write the sample data.
// 
// On devices that don’t support [CounterSamplingPointAtStageBoundary] you
// need to set the value to [MTLCounterDontSample].
//
// [MTLCounterDontSample]: https://developer.apple.com/documentation/Metal/MTLCounterDontSample
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPassSampleBufferAttachmentDescriptor/startOfFragmentSampleIndex
func (r MTLRenderPassSampleBufferAttachmentDescriptor) StartOfFragmentSampleIndex() uint {
	rv := objc.Send[uint](r.ID, objc.Sel("startOfFragmentSampleIndex"))
	return rv
}
func (r MTLRenderPassSampleBufferAttachmentDescriptor) SetStartOfFragmentSampleIndex(value uint) {
	objc.Send[struct{}](r.ID, objc.Sel("setStartOfFragmentSampleIndex:"), value)
}
// The index the Metal device object should use to store GPU counters when
// ending the render pass’s fragment stage.
//
// # Discussion
// 
// Specify [MTLCounterDontSample] if you don’t want to sample GPU counters
// at the end of the fragment stage. Otherwise, specify an index within the
// sample buffer where you want the GPU to write the sample data.
// 
// On devices that don’t support [CounterSamplingPointAtStageBoundary] you
// need to set the value to [MTLCounterDontSample].
//
// [MTLCounterDontSample]: https://developer.apple.com/documentation/Metal/MTLCounterDontSample
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPassSampleBufferAttachmentDescriptor/endOfFragmentSampleIndex
func (r MTLRenderPassSampleBufferAttachmentDescriptor) EndOfFragmentSampleIndex() uint {
	rv := objc.Send[uint](r.ID, objc.Sel("endOfFragmentSampleIndex"))
	return rv
}
func (r MTLRenderPassSampleBufferAttachmentDescriptor) SetEndOfFragmentSampleIndex(value uint) {
	objc.Send[struct{}](r.ID, objc.Sel("setEndOfFragmentSampleIndex:"), value)
}

