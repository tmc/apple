// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLAccelerationStructurePassSampleBufferAttachmentDescriptor] class.
var (
	_MTLAccelerationStructurePassSampleBufferAttachmentDescriptorClass     MTLAccelerationStructurePassSampleBufferAttachmentDescriptorClass
	_MTLAccelerationStructurePassSampleBufferAttachmentDescriptorClassOnce sync.Once
)

func getMTLAccelerationStructurePassSampleBufferAttachmentDescriptorClass() MTLAccelerationStructurePassSampleBufferAttachmentDescriptorClass {
	_MTLAccelerationStructurePassSampleBufferAttachmentDescriptorClassOnce.Do(func() {
		_MTLAccelerationStructurePassSampleBufferAttachmentDescriptorClass = MTLAccelerationStructurePassSampleBufferAttachmentDescriptorClass{class: objc.GetClass("MTLAccelerationStructurePassSampleBufferAttachmentDescriptor")}
	})
	return _MTLAccelerationStructurePassSampleBufferAttachmentDescriptorClass
}

// GetMTLAccelerationStructurePassSampleBufferAttachmentDescriptorClass returns the class object for MTLAccelerationStructurePassSampleBufferAttachmentDescriptor.
func GetMTLAccelerationStructurePassSampleBufferAttachmentDescriptorClass() MTLAccelerationStructurePassSampleBufferAttachmentDescriptorClass {
	return getMTLAccelerationStructurePassSampleBufferAttachmentDescriptorClass()
}

type MTLAccelerationStructurePassSampleBufferAttachmentDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLAccelerationStructurePassSampleBufferAttachmentDescriptorClass) Alloc() MTLAccelerationStructurePassSampleBufferAttachmentDescriptor {
	rv := objc.Send[MTLAccelerationStructurePassSampleBufferAttachmentDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Instance Properties
//
//   - [MTLAccelerationStructurePassSampleBufferAttachmentDescriptor.EndOfEncoderSampleIndex]
//   - [MTLAccelerationStructurePassSampleBufferAttachmentDescriptor.SetEndOfEncoderSampleIndex]
//   - [MTLAccelerationStructurePassSampleBufferAttachmentDescriptor.SampleBuffer]: A specialized memory buffer that the GPU uses to store its counter data during the acceleration structure pass.
//   - [MTLAccelerationStructurePassSampleBufferAttachmentDescriptor.SetSampleBuffer]
//   - [MTLAccelerationStructurePassSampleBufferAttachmentDescriptor.StartOfEncoderSampleIndex]
//   - [MTLAccelerationStructurePassSampleBufferAttachmentDescriptor.SetStartOfEncoderSampleIndex]
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructurePassSampleBufferAttachmentDescriptor
type MTLAccelerationStructurePassSampleBufferAttachmentDescriptor struct {
	objectivec.Object
}

// MTLAccelerationStructurePassSampleBufferAttachmentDescriptorFromID constructs a [MTLAccelerationStructurePassSampleBufferAttachmentDescriptor] from an objc.ID.
func MTLAccelerationStructurePassSampleBufferAttachmentDescriptorFromID(id objc.ID) MTLAccelerationStructurePassSampleBufferAttachmentDescriptor {
	return MTLAccelerationStructurePassSampleBufferAttachmentDescriptor{objectivec.Object{ID: id}}
}
// NOTE: MTLAccelerationStructurePassSampleBufferAttachmentDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLAccelerationStructurePassSampleBufferAttachmentDescriptor] class.
//
// # Instance Properties
//
//   - [IMTLAccelerationStructurePassSampleBufferAttachmentDescriptor.EndOfEncoderSampleIndex]
//   - [IMTLAccelerationStructurePassSampleBufferAttachmentDescriptor.SetEndOfEncoderSampleIndex]
//   - [IMTLAccelerationStructurePassSampleBufferAttachmentDescriptor.SampleBuffer]: A specialized memory buffer that the GPU uses to store its counter data during the acceleration structure pass.
//   - [IMTLAccelerationStructurePassSampleBufferAttachmentDescriptor.SetSampleBuffer]
//   - [IMTLAccelerationStructurePassSampleBufferAttachmentDescriptor.StartOfEncoderSampleIndex]
//   - [IMTLAccelerationStructurePassSampleBufferAttachmentDescriptor.SetStartOfEncoderSampleIndex]
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructurePassSampleBufferAttachmentDescriptor
type IMTLAccelerationStructurePassSampleBufferAttachmentDescriptor interface {
	objectivec.IObject

	// Topic: Instance Properties

	EndOfEncoderSampleIndex() uint
	SetEndOfEncoderSampleIndex(value uint)
	// A specialized memory buffer that the GPU uses to store its counter data during the acceleration structure pass.
	SampleBuffer() MTLCounterSampleBuffer
	SetSampleBuffer(value MTLCounterSampleBuffer)
	StartOfEncoderSampleIndex() uint
	SetStartOfEncoderSampleIndex(value uint)
}

// Init initializes the instance.
func (a MTLAccelerationStructurePassSampleBufferAttachmentDescriptor) Init() MTLAccelerationStructurePassSampleBufferAttachmentDescriptor {
	rv := objc.Send[MTLAccelerationStructurePassSampleBufferAttachmentDescriptor](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a MTLAccelerationStructurePassSampleBufferAttachmentDescriptor) Autorelease() MTLAccelerationStructurePassSampleBufferAttachmentDescriptor {
	rv := objc.Send[MTLAccelerationStructurePassSampleBufferAttachmentDescriptor](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLAccelerationStructurePassSampleBufferAttachmentDescriptor creates a new MTLAccelerationStructurePassSampleBufferAttachmentDescriptor instance.
func NewMTLAccelerationStructurePassSampleBufferAttachmentDescriptor() MTLAccelerationStructurePassSampleBufferAttachmentDescriptor {
	class := getMTLAccelerationStructurePassSampleBufferAttachmentDescriptorClass()
	rv := objc.Send[MTLAccelerationStructurePassSampleBufferAttachmentDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructurePassSampleBufferAttachmentDescriptor/endOfEncoderSampleIndex
func (a MTLAccelerationStructurePassSampleBufferAttachmentDescriptor) EndOfEncoderSampleIndex() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("endOfEncoderSampleIndex"))
	return rv
}
func (a MTLAccelerationStructurePassSampleBufferAttachmentDescriptor) SetEndOfEncoderSampleIndex(value uint) {
	objc.Send[struct{}](a.ID, objc.Sel("setEndOfEncoderSampleIndex:"), value)
}
// A specialized memory buffer that the GPU uses to store its counter data
// during the acceleration structure pass.
//
// # Discussion
// 
// The property defaults to `nil`, which means the GPU doesn’t save any GPU
// counter information during the acceleration structure pass. See [Creating a
// counter sample buffer to store a GPU’s counter data during a pass] and
// [Sampling GPU data into counter sample buffers] for more information.
//
// [Creating a counter sample buffer to store a GPU’s counter data during a pass]: https://developer.apple.com/documentation/Metal/creating-a-counter-sample-buffer-to-store-a-gpus-counter-data-during-a-pass
// [Sampling GPU data into counter sample buffers]: https://developer.apple.com/documentation/Metal/sampling-gpu-data-into-counter-sample-buffers
//
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructurePassSampleBufferAttachmentDescriptor/sampleBuffer
func (a MTLAccelerationStructurePassSampleBufferAttachmentDescriptor) SampleBuffer() MTLCounterSampleBuffer {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("sampleBuffer"))
	return MTLCounterSampleBufferObjectFromID(rv)
}
func (a MTLAccelerationStructurePassSampleBufferAttachmentDescriptor) SetSampleBuffer(value MTLCounterSampleBuffer) {
	objc.Send[struct{}](a.ID, objc.Sel("setSampleBuffer:"), value)
}
// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructurePassSampleBufferAttachmentDescriptor/startOfEncoderSampleIndex
func (a MTLAccelerationStructurePassSampleBufferAttachmentDescriptor) StartOfEncoderSampleIndex() uint {
	rv := objc.Send[uint](a.ID, objc.Sel("startOfEncoderSampleIndex"))
	return rv
}
func (a MTLAccelerationStructurePassSampleBufferAttachmentDescriptor) SetStartOfEncoderSampleIndex(value uint) {
	objc.Send[struct{}](a.ID, objc.Sel("setStartOfEncoderSampleIndex:"), value)
}

