// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLCommandQueueDescriptor] class.
var (
	_MTLCommandQueueDescriptorClass     MTLCommandQueueDescriptorClass
	_MTLCommandQueueDescriptorClassOnce sync.Once
)

func getMTLCommandQueueDescriptorClass() MTLCommandQueueDescriptorClass {
	_MTLCommandQueueDescriptorClassOnce.Do(func() {
		_MTLCommandQueueDescriptorClass = MTLCommandQueueDescriptorClass{class: objc.GetClass("MTLCommandQueueDescriptor")}
	})
	return _MTLCommandQueueDescriptorClass
}

// GetMTLCommandQueueDescriptorClass returns the class object for MTLCommandQueueDescriptor.
func GetMTLCommandQueueDescriptorClass() MTLCommandQueueDescriptorClass {
	return getMTLCommandQueueDescriptorClass()
}

type MTLCommandQueueDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTLCommandQueueDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLCommandQueueDescriptorClass) Alloc() MTLCommandQueueDescriptor {
	rv := objc.Send[MTLCommandQueueDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A configuration that customizes the behavior for a new command queue.
//
// # Instance Properties
//
//   - [MTLCommandQueueDescriptor.LogState]: The shader logging configuration that the command queue uses.
//   - [MTLCommandQueueDescriptor.SetLogState]
//   - [MTLCommandQueueDescriptor.MaxCommandBufferCount]: An integer that sets the maximum number of uncompleted command buffers the queue can allow.
//   - [MTLCommandQueueDescriptor.SetMaxCommandBufferCount]
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandQueueDescriptor
type MTLCommandQueueDescriptor struct {
	objectivec.Object
}

// MTLCommandQueueDescriptorFromID constructs a [MTLCommandQueueDescriptor] from an objc.ID.
//
// A configuration that customizes the behavior for a new command queue.
func MTLCommandQueueDescriptorFromID(id objc.ID) MTLCommandQueueDescriptor {
	return MTLCommandQueueDescriptor{objectivec.Object{ID: id}}
}
// NOTE: MTLCommandQueueDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLCommandQueueDescriptor] class.
//
// # Instance Properties
//
//   - [IMTLCommandQueueDescriptor.LogState]: The shader logging configuration that the command queue uses.
//   - [IMTLCommandQueueDescriptor.SetLogState]
//   - [IMTLCommandQueueDescriptor.MaxCommandBufferCount]: An integer that sets the maximum number of uncompleted command buffers the queue can allow.
//   - [IMTLCommandQueueDescriptor.SetMaxCommandBufferCount]
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandQueueDescriptor
type IMTLCommandQueueDescriptor interface {
	objectivec.IObject

	// Topic: Instance Properties

	// The shader logging configuration that the command queue uses.
	LogState() MTLLogState
	SetLogState(value MTLLogState)
	// An integer that sets the maximum number of uncompleted command buffers the queue can allow.
	MaxCommandBufferCount() uint
	SetMaxCommandBufferCount(value uint)

	// The domain for Metal command buffer errors.
	MTLCommandBufferErrorDomain() string
}

// Init initializes the instance.
func (c MTLCommandQueueDescriptor) Init() MTLCommandQueueDescriptor {
	rv := objc.Send[MTLCommandQueueDescriptor](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MTLCommandQueueDescriptor) Autorelease() MTLCommandQueueDescriptor {
	rv := objc.Send[MTLCommandQueueDescriptor](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLCommandQueueDescriptor creates a new MTLCommandQueueDescriptor instance.
func NewMTLCommandQueueDescriptor() MTLCommandQueueDescriptor {
	class := getMTLCommandQueueDescriptorClass()
	rv := objc.Send[MTLCommandQueueDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The shader logging configuration that the command queue uses.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandQueueDescriptor/logState
func (c MTLCommandQueueDescriptor) LogState() MTLLogState {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("logState"))
	return MTLLogStateObjectFromID(rv)
}
func (c MTLCommandQueueDescriptor) SetLogState(value MTLLogState) {
	objc.Send[struct{}](c.ID, objc.Sel("setLogState:"), value)
}
// An integer that sets the maximum number of uncompleted command buffers the
// queue can allow.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandQueueDescriptor/maxCommandBufferCount
func (c MTLCommandQueueDescriptor) MaxCommandBufferCount() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("maxCommandBufferCount"))
	return rv
}
func (c MTLCommandQueueDescriptor) SetMaxCommandBufferCount(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setMaxCommandBufferCount:"), value)
}
// The domain for Metal command buffer errors.
//
// See: https://developer.apple.com/documentation/metal/mtlcommandbuffererrordomain
func (c MTLCommandQueueDescriptor) MTLCommandBufferErrorDomain() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("MTLCommandBufferErrorDomain"))
	return foundation.NSStringFromID(rv).String()
}

