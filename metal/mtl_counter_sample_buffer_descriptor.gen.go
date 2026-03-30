// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLCounterSampleBufferDescriptor] class.
var (
	_MTLCounterSampleBufferDescriptorClass     MTLCounterSampleBufferDescriptorClass
	_MTLCounterSampleBufferDescriptorClassOnce sync.Once
)

func getMTLCounterSampleBufferDescriptorClass() MTLCounterSampleBufferDescriptorClass {
	_MTLCounterSampleBufferDescriptorClassOnce.Do(func() {
		_MTLCounterSampleBufferDescriptorClass = MTLCounterSampleBufferDescriptorClass{class: objc.GetClass("MTLCounterSampleBufferDescriptor")}
	})
	return _MTLCounterSampleBufferDescriptorClass
}

// GetMTLCounterSampleBufferDescriptorClass returns the class object for MTLCounterSampleBufferDescriptor.
func GetMTLCounterSampleBufferDescriptorClass() MTLCounterSampleBufferDescriptorClass {
	return getMTLCounterSampleBufferDescriptorClass()
}

type MTLCounterSampleBufferDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLCounterSampleBufferDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLCounterSampleBufferDescriptorClass) Alloc() MTLCounterSampleBufferDescriptor {
	rv := objc.Send[MTLCounterSampleBufferDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A group of properties that configures the counter sample buffers you create
// with it.
//
// # Overview
//
// To create a new counter sample buffer, create and configure an
// [MTLCounterSampleBufferDescriptor] instance, and then call an [MTLDevice]
// instance’s [NewCounterSampleBufferWithDescriptorError] method. See
// [Creating a counter sample buffer to store a GPU’s counter data during a
// pass].
//
// Each new sample counter buffer inherits the values of the descriptor’s
// properties when you create it. You can modify a descriptor and reuse it to
// create other counter sample buffers, which has no effect on existing
// counter sample buffers.
//
// # Configuring a descriptor for a counter sample buffer
//
//   - [MTLCounterSampleBufferDescriptor.CounterSet]: A GPU device’s counter set instance that you want to sample.
//   - [MTLCounterSampleBufferDescriptor.SetCounterSet]
//   - [MTLCounterSampleBufferDescriptor.Label]: The name for the counter sample buffer you create with the descriptor.
//   - [MTLCounterSampleBufferDescriptor.SetLabel]
//   - [MTLCounterSampleBufferDescriptor.SampleCount]: The number of instances of a counter set’s data that a counter sample buffer can store.
//   - [MTLCounterSampleBufferDescriptor.SetSampleCount]
//   - [MTLCounterSampleBufferDescriptor.StorageMode]: The memory storage mode for the counter sample buffers you create with the descriptor.
//   - [MTLCounterSampleBufferDescriptor.SetStorageMode]
//
// See: https://developer.apple.com/documentation/Metal/MTLCounterSampleBufferDescriptor
//
// [Creating a counter sample buffer to store a GPU’s counter data during a pass]: https://developer.apple.com/documentation/Metal/creating-a-counter-sample-buffer-to-store-a-gpus-counter-data-during-a-pass
type MTLCounterSampleBufferDescriptor struct {
	objectivec.Object
}

// MTLCounterSampleBufferDescriptorFromID constructs a [MTLCounterSampleBufferDescriptor] from an objc.ID.
//
// A group of properties that configures the counter sample buffers you create
// with it.
func MTLCounterSampleBufferDescriptorFromID(id objc.ID) MTLCounterSampleBufferDescriptor {
	return MTLCounterSampleBufferDescriptor{objectivec.Object{ID: id}}
}

// NOTE: MTLCounterSampleBufferDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLCounterSampleBufferDescriptor] class.
//
// # Configuring a descriptor for a counter sample buffer
//
//   - [IMTLCounterSampleBufferDescriptor.CounterSet]: A GPU device’s counter set instance that you want to sample.
//   - [IMTLCounterSampleBufferDescriptor.SetCounterSet]
//   - [IMTLCounterSampleBufferDescriptor.Label]: The name for the counter sample buffer you create with the descriptor.
//   - [IMTLCounterSampleBufferDescriptor.SetLabel]
//   - [IMTLCounterSampleBufferDescriptor.SampleCount]: The number of instances of a counter set’s data that a counter sample buffer can store.
//   - [IMTLCounterSampleBufferDescriptor.SetSampleCount]
//   - [IMTLCounterSampleBufferDescriptor.StorageMode]: The memory storage mode for the counter sample buffers you create with the descriptor.
//   - [IMTLCounterSampleBufferDescriptor.SetStorageMode]
//
// See: https://developer.apple.com/documentation/Metal/MTLCounterSampleBufferDescriptor
type IMTLCounterSampleBufferDescriptor interface {
	objectivec.IObject

	// Topic: Configuring a descriptor for a counter sample buffer

	// A GPU device’s counter set instance that you want to sample.
	CounterSet() MTLCounterSet
	SetCounterSet(value MTLCounterSet)
	// The name for the counter sample buffer you create with the descriptor.
	Label() string
	SetLabel(value string)
	// The number of instances of a counter set’s data that a counter sample buffer can store.
	SampleCount() uint
	SetSampleCount(value uint)
	// The memory storage mode for the counter sample buffers you create with the descriptor.
	StorageMode() MTLStorageMode
	SetStorageMode(value MTLStorageMode)

	// A sentinel value that instructs an encoder to skip sampling a counter as the GPU runs the encoder’s pass.
	MTLCounterDontSample() int
	SetMTLCounterDontSample(value int)
}

// Init initializes the instance.
func (c MTLCounterSampleBufferDescriptor) Init() MTLCounterSampleBufferDescriptor {
	rv := objc.Send[MTLCounterSampleBufferDescriptor](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MTLCounterSampleBufferDescriptor) Autorelease() MTLCounterSampleBufferDescriptor {
	rv := objc.Send[MTLCounterSampleBufferDescriptor](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLCounterSampleBufferDescriptor creates a new MTLCounterSampleBufferDescriptor instance.
func NewMTLCounterSampleBufferDescriptor() MTLCounterSampleBufferDescriptor {
	class := getMTLCounterSampleBufferDescriptorClass()
	rv := objc.Send[MTLCounterSampleBufferDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A GPU device’s counter set instance that you want to sample.
//
// # Discussion
//
// Assign this property to one of the counter sets in an [MTLDevice]
// instance’s [CounterSets] property.
//
// See: https://developer.apple.com/documentation/Metal/MTLCounterSampleBufferDescriptor/counterSet
func (c MTLCounterSampleBufferDescriptor) CounterSet() MTLCounterSet {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("counterSet"))
	return MTLCounterSetObjectFromID(rv)
}
func (c MTLCounterSampleBufferDescriptor) SetCounterSet(value MTLCounterSet) {
	objc.Send[struct{}](c.ID, objc.Sel("setCounterSet:"), value)
}

// The name for the counter sample buffer you create with the descriptor.
//
// See: https://developer.apple.com/documentation/Metal/MTLCounterSampleBufferDescriptor/label
func (c MTLCounterSampleBufferDescriptor) Label() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}
func (c MTLCounterSampleBufferDescriptor) SetLabel(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setLabel:"), objc.String(value))
}

// The number of instances of a counter set’s data that a counter sample
// buffer can store.
//
// # Discussion
//
// The counter sample buffer instances you create with the
// [MTLCounterSampleBufferDescriptor] can store [SampleCount] instances of a
// counter set.
//
// See: https://developer.apple.com/documentation/Metal/MTLCounterSampleBufferDescriptor/sampleCount
func (c MTLCounterSampleBufferDescriptor) SampleCount() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("sampleCount"))
	return rv
}
func (c MTLCounterSampleBufferDescriptor) SetSampleCount(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setSampleCount:"), value)
}

// The memory storage mode for the counter sample buffers you create with the
// descriptor.
//
// # Discussion
//
// To access a counter sample buffer with the CPU, set this property to
// [MTLStorageModeShared], otherwise [MTLStorageModePrivate].
//
// See: https://developer.apple.com/documentation/Metal/MTLCounterSampleBufferDescriptor/storageMode
func (c MTLCounterSampleBufferDescriptor) StorageMode() MTLStorageMode {
	rv := objc.Send[MTLStorageMode](c.ID, objc.Sel("storageMode"))
	return MTLStorageMode(rv)
}
func (c MTLCounterSampleBufferDescriptor) SetStorageMode(value MTLStorageMode) {
	objc.Send[struct{}](c.ID, objc.Sel("setStorageMode:"), value)
}

// A sentinel value that instructs an encoder to skip sampling a counter as
// the GPU runs the encoder’s pass.
//
// See: https://developer.apple.com/documentation/metal/mtlcounterdontsample
func (c MTLCounterSampleBufferDescriptor) MTLCounterDontSample() int {
	rv := objc.Send[int](c.ID, objc.Sel("MTLCounterDontSample"))
	return rv
}
func (c MTLCounterSampleBufferDescriptor) SetMTLCounterDontSample(value int) {
	objc.Send[struct{}](c.ID, objc.Sel("setMTLCounterDontSample:"), value)
}
