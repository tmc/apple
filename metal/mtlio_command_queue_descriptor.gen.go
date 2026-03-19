// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLIOCommandQueueDescriptor] class.
var (
	_MTLIOCommandQueueDescriptorClass     MTLIOCommandQueueDescriptorClass
	_MTLIOCommandQueueDescriptorClassOnce sync.Once
)

func getMTLIOCommandQueueDescriptorClass() MTLIOCommandQueueDescriptorClass {
	_MTLIOCommandQueueDescriptorClassOnce.Do(func() {
		_MTLIOCommandQueueDescriptorClass = MTLIOCommandQueueDescriptorClass{class: objc.GetClass("MTLIOCommandQueueDescriptor")}
	})
	return _MTLIOCommandQueueDescriptorClass
}

// GetMTLIOCommandQueueDescriptorClass returns the class object for MTLIOCommandQueueDescriptor.
func GetMTLIOCommandQueueDescriptorClass() MTLIOCommandQueueDescriptorClass {
	return getMTLIOCommandQueueDescriptorClass()
}

type MTLIOCommandQueueDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLIOCommandQueueDescriptorClass) Alloc() MTLIOCommandQueueDescriptor {
	rv := objc.Send[MTLIOCommandQueueDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A configuration template you use to create a new input/output command
// queue.
//
// # Overview
// 
// Use this descriptor type to configure the settings of each input/output
// command queue that you create using [NewIOCommandQueueWithDescriptorError].
// To create additional input/output command queues, you can reuse a
// descriptor instance and optionally reconfigure its properties.
// 
// Create each input/output queue to meet your apps needs by setting the
// descriptor’s properties.
// 
// - Select a queue’s relative level of importance with the [MTLIOCommandQueueDescriptor.Priority]
// property. - Create a queue that runs multiple input/output command buffers
// in parallel by setting the [Type] property to
// [IOCommandQueueTypeConcurrent]. - Decide how many individual commands a
// queue can run simultaneously with the [MTLIOCommandQueueDescriptor.MaxCommandsInFlight] property. -
// Choose how many command buffers a queue can have waiting to run with
// [MTLIOCommandQueueDescriptor.MaxCommandBufferCount] property. - Take control of the queue’s scratch
// memory allocation by implementing [MTLIOScratchBufferAllocator] and assign
// an instance of it to the [MTLIOCommandQueueDescriptor.ScratchBufferAllocator] property.
//
// # Configuring the input/output command queue
//
//   - [MTLIOCommandQueueDescriptor.Priority]: Configures the priority for a new input/output command queue.
//   - [MTLIOCommandQueueDescriptor.SetPriority]
//   - [MTLIOCommandQueueDescriptor.Type]: Configures the queue type for a new input/output command queue.
//   - [MTLIOCommandQueueDescriptor.SetType]
//   - [MTLIOCommandQueueDescriptor.MaxCommandsInFlight]: Sets the largest number of individual commands that an input/output command queue can run at a time.
//   - [MTLIOCommandQueueDescriptor.SetMaxCommandsInFlight]
//   - [MTLIOCommandQueueDescriptor.MaxCommandBufferCount]: Sets the largest number of outstanding input/output command buffers a queue can have at any point in time.
//   - [MTLIOCommandQueueDescriptor.SetMaxCommandBufferCount]
//
// # Providing your own a scratch buffer
//
//   - [MTLIOCommandQueueDescriptor.ScratchBufferAllocator]: An optional memory allocator that you implement to manage the scratch memory that an input/output command queue requests.
//   - [MTLIOCommandQueueDescriptor.SetScratchBufferAllocator]
//
// See: https://developer.apple.com/documentation/Metal/MTLIOCommandQueueDescriptor
type MTLIOCommandQueueDescriptor struct {
	objectivec.Object
}

// MTLIOCommandQueueDescriptorFromID constructs a [MTLIOCommandQueueDescriptor] from an objc.ID.
//
// A configuration template you use to create a new input/output command
// queue.
func MTLIOCommandQueueDescriptorFromID(id objc.ID) MTLIOCommandQueueDescriptor {
	return MTLIOCommandQueueDescriptor{objectivec.Object{ID: id}}
}
// NOTE: MTLIOCommandQueueDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLIOCommandQueueDescriptor] class.
//
// # Configuring the input/output command queue
//
//   - [IMTLIOCommandQueueDescriptor.Priority]: Configures the priority for a new input/output command queue.
//   - [IMTLIOCommandQueueDescriptor.SetPriority]
//   - [IMTLIOCommandQueueDescriptor.Type]: Configures the queue type for a new input/output command queue.
//   - [IMTLIOCommandQueueDescriptor.SetType]
//   - [IMTLIOCommandQueueDescriptor.MaxCommandsInFlight]: Sets the largest number of individual commands that an input/output command queue can run at a time.
//   - [IMTLIOCommandQueueDescriptor.SetMaxCommandsInFlight]
//   - [IMTLIOCommandQueueDescriptor.MaxCommandBufferCount]: Sets the largest number of outstanding input/output command buffers a queue can have at any point in time.
//   - [IMTLIOCommandQueueDescriptor.SetMaxCommandBufferCount]
//
// # Providing your own a scratch buffer
//
//   - [IMTLIOCommandQueueDescriptor.ScratchBufferAllocator]: An optional memory allocator that you implement to manage the scratch memory that an input/output command queue requests.
//   - [IMTLIOCommandQueueDescriptor.SetScratchBufferAllocator]
//
// See: https://developer.apple.com/documentation/Metal/MTLIOCommandQueueDescriptor
type IMTLIOCommandQueueDescriptor interface {
	objectivec.IObject

	// Topic: Configuring the input/output command queue

	// Configures the priority for a new input/output command queue.
	Priority() MTLIOPriority
	SetPriority(value MTLIOPriority)
	// Configures the queue type for a new input/output command queue.
	Type() MTLIOCommandQueueType
	SetType(value MTLIOCommandQueueType)
	// Sets the largest number of individual commands that an input/output command queue can run at a time.
	MaxCommandsInFlight() uint
	SetMaxCommandsInFlight(value uint)
	// Sets the largest number of outstanding input/output command buffers a queue can have at any point in time.
	MaxCommandBufferCount() uint
	SetMaxCommandBufferCount(value uint)

	// Topic: Providing your own a scratch buffer

	// An optional memory allocator that you implement to manage the scratch memory that an input/output command queue requests.
	ScratchBufferAllocator() MTLIOScratchBufferAllocator
	SetScratchBufferAllocator(value MTLIOScratchBufferAllocator)
}

// Init initializes the instance.
func (i MTLIOCommandQueueDescriptor) Init() MTLIOCommandQueueDescriptor {
	rv := objc.Send[MTLIOCommandQueueDescriptor](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MTLIOCommandQueueDescriptor) Autorelease() MTLIOCommandQueueDescriptor {
	rv := objc.Send[MTLIOCommandQueueDescriptor](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLIOCommandQueueDescriptor creates a new MTLIOCommandQueueDescriptor instance.
func NewMTLIOCommandQueueDescriptor() MTLIOCommandQueueDescriptor {
	class := getMTLIOCommandQueueDescriptorClass()
	rv := objc.Send[MTLIOCommandQueueDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Configures the priority for a new input/output command queue.
//
// See: https://developer.apple.com/documentation/Metal/MTLIOCommandQueueDescriptor/priority
func (i MTLIOCommandQueueDescriptor) Priority() MTLIOPriority {
	rv := objc.Send[MTLIOPriority](i.ID, objc.Sel("priority"))
	return MTLIOPriority(rv)
}
func (i MTLIOCommandQueueDescriptor) SetPriority(value MTLIOPriority) {
	objc.Send[struct{}](i.ID, objc.Sel("setPriority:"), value)
}
// Configures the queue type for a new input/output command queue.
//
// See: https://developer.apple.com/documentation/Metal/MTLIOCommandQueueDescriptor/type
func (i MTLIOCommandQueueDescriptor) Type() MTLIOCommandQueueType {
	rv := objc.Send[MTLIOCommandQueueType](i.ID, objc.Sel("type"))
	return MTLIOCommandQueueType(rv)
}
func (i MTLIOCommandQueueDescriptor) SetType(value MTLIOCommandQueueType) {
	objc.Send[struct{}](i.ID, objc.Sel("setType:"), value)
}
// Sets the largest number of individual commands that an input/output command
// queue can run at a time.
//
// # Discussion
// 
// Set to `0` to instruct Metal to select an appropriate value for you —
// based on the system’s available memory.
//
// See: https://developer.apple.com/documentation/Metal/MTLIOCommandQueueDescriptor/maxCommandsInFlight
func (i MTLIOCommandQueueDescriptor) MaxCommandsInFlight() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("maxCommandsInFlight"))
	return rv
}
func (i MTLIOCommandQueueDescriptor) SetMaxCommandsInFlight(value uint) {
	objc.Send[struct{}](i.ID, objc.Sel("setMaxCommandsInFlight:"), value)
}
// Sets the largest number of outstanding input/output command buffers a queue
// can have at any point in time.
//
// # Discussion
// 
// The input/output command buffers that count against this limit are those
// that are currently executing in a queue or waiting to execute. The command
// buffers that have finished executing no longer count against this limit.
//
// See: https://developer.apple.com/documentation/Metal/MTLIOCommandQueueDescriptor/maxCommandBufferCount
func (i MTLIOCommandQueueDescriptor) MaxCommandBufferCount() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("maxCommandBufferCount"))
	return rv
}
func (i MTLIOCommandQueueDescriptor) SetMaxCommandBufferCount(value uint) {
	objc.Send[struct{}](i.ID, objc.Sel("setMaxCommandBufferCount:"), value)
}
// An optional memory allocator that you implement to manage the scratch
// memory that an input/output command queue requests.
//
// # Discussion
// 
// Your app can manage an input/output command queue’s scratch memory by an
// implementing [MTLIOScratchBufferAllocator] in one of your types, and
// assigning an instance of it to [ScratchBufferAllocator]. Otherwise, set to
// `nil` to instruct the input/output command queue to allocate and manage its
// own scratch buffers.
//
// See: https://developer.apple.com/documentation/Metal/MTLIOCommandQueueDescriptor/scratchBufferAllocator
func (i MTLIOCommandQueueDescriptor) ScratchBufferAllocator() MTLIOScratchBufferAllocator {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("scratchBufferAllocator"))
	return MTLIOScratchBufferAllocatorObjectFromID(rv)
}
func (i MTLIOCommandQueueDescriptor) SetScratchBufferAllocator(value MTLIOScratchBufferAllocator) {
	objc.Send[struct{}](i.ID, objc.Sel("setScratchBufferAllocator:"), value)
}

