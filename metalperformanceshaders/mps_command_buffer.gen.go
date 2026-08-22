// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"sync"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSCommandBuffer] class.
var (
	_MPSCommandBufferClass     MPSCommandBufferClass
	_MPSCommandBufferClassOnce sync.Once
)

func getMPSCommandBufferClass() MPSCommandBufferClass {
	_MPSCommandBufferClassOnce.Do(func() {
		_MPSCommandBufferClass = MPSCommandBufferClass{class: objc.GetClass("MPSCommandBuffer")}
	})
	return _MPSCommandBufferClass
}

// GetMPSCommandBufferClass returns the class object for MPSCommandBuffer.
func GetMPSCommandBufferClass() MPSCommandBufferClass {
	return getMPSCommandBufferClass()
}

type MPSCommandBufferClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSCommandBufferClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSCommandBufferClass) Alloc() MPSCommandBuffer {
	rv := objc.Send[MPSCommandBuffer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Initializers
//
//   - [MPSCommandBuffer.InitWithCommandBuffer]
//
// # Instance Properties
//
//   - [MPSCommandBuffer.CommandBuffer]
//   - [MPSCommandBuffer.HeapProvider]
//   - [MPSCommandBuffer.SetHeapProvider]
//   - [MPSCommandBuffer.Predicate]
//   - [MPSCommandBuffer.SetPredicate]
//   - [MPSCommandBuffer.RootCommandBuffer]
//
// # Instance Methods
//
//   - [MPSCommandBuffer.CommitAndContinue]
//   - [MPSCommandBuffer.PrefetchHeapForWorkloadSize]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCommandBuffer
type MPSCommandBuffer struct {
	objectivec.Object
}

// MPSCommandBufferFromID constructs a [MPSCommandBuffer] from an objc.ID.
func MPSCommandBufferFromID(id objc.ID) MPSCommandBuffer {
	return MPSCommandBuffer{objectivec.Object{ID: id}}
}

// NOTE: MPSCommandBuffer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSCommandBuffer] class.
//
// # Initializers
//
//   - [IMPSCommandBuffer.InitWithCommandBuffer]
//
// # Instance Properties
//
//   - [IMPSCommandBuffer.CommandBuffer]
//   - [IMPSCommandBuffer.HeapProvider]
//   - [IMPSCommandBuffer.SetHeapProvider]
//   - [IMPSCommandBuffer.Predicate]
//   - [IMPSCommandBuffer.SetPredicate]
//   - [IMPSCommandBuffer.RootCommandBuffer]
//
// # Instance Methods
//
//   - [IMPSCommandBuffer.CommitAndContinue]
//   - [IMPSCommandBuffer.PrefetchHeapForWorkloadSize]
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCommandBuffer
type IMPSCommandBuffer interface {
	objectivec.IObject

	// Topic: Initializers

	InitWithCommandBuffer(commandBuffer metal.MTLCommandBuffer) MPSCommandBuffer

	// Topic: Instance Properties

	CommandBuffer() metal.MTLCommandBuffer
	HeapProvider() MPSHeapProvider
	SetHeapProvider(value MPSHeapProvider)
	Predicate() IMPSPredicate
	SetPredicate(value IMPSPredicate)
	RootCommandBuffer() metal.MTLCommandBuffer

	// Topic: Instance Methods

	CommitAndContinue()
	PrefetchHeapForWorkloadSize(size uintptr)
}

// Init initializes the instance.
func (c MPSCommandBuffer) Init() MPSCommandBuffer {
	rv := objc.Send[MPSCommandBuffer](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MPSCommandBuffer) Autorelease() MPSCommandBuffer {
	rv := objc.Send[MPSCommandBuffer](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSCommandBuffer creates a new MPSCommandBuffer instance.
func NewMPSCommandBuffer() MPSCommandBuffer {
	class := getMPSCommandBufferClass()
	rv := objc.Send[MPSCommandBuffer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCommandBuffer/init(from:)
func NewCommandBufferFromCommandQueue(commandQueue metal.MTLCommandQueue) MPSCommandBuffer {
	rv := objc.Send[objc.ID](objc.ID(getMPSCommandBufferClass().class), objc.Sel("commandBufferFromCommandQueue:"), commandQueue)
	return MPSCommandBufferFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCommandBuffer/init(commandBuffer:)
func NewCommandBufferWithCommandBuffer(commandBuffer metal.MTLCommandBuffer) MPSCommandBuffer {
	instance := getMPSCommandBufferClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCommandBuffer:"), commandBuffer)
	return MPSCommandBufferFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCommandBuffer/init(commandBuffer:)
func (c MPSCommandBuffer) InitWithCommandBuffer(commandBuffer metal.MTLCommandBuffer) MPSCommandBuffer {
	rv := objc.Send[MPSCommandBuffer](c.ID, objc.Sel("initWithCommandBuffer:"), commandBuffer)
	return rv
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCommandBuffer/commitAndContinue()
func (c MPSCommandBuffer) CommitAndContinue() {
	objc.Send[objc.ID](c.ID, objc.Sel("commitAndContinue"))
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCommandBuffer/prefetchHeap(forWorkloadSize:)
func (c MPSCommandBuffer) PrefetchHeapForWorkloadSize(size uintptr) {
	objc.Send[objc.ID](c.ID, objc.Sel("prefetchHeapForWorkloadSize:"), size)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCommandBuffer/commandBufferWithCommandBuffer:
func (_MPSCommandBufferClass MPSCommandBufferClass) CommandBufferWithCommandBuffer(commandBuffer metal.MTLCommandBuffer) MPSCommandBuffer {
	rv := objc.Send[objc.ID](objc.ID(_MPSCommandBufferClass.class), objc.Sel("commandBufferWithCommandBuffer:"), commandBuffer)
	return MPSCommandBufferFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCommandBuffer/commandBuffer
func (c MPSCommandBuffer) CommandBuffer() metal.MTLCommandBuffer {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("commandBuffer"))
	return metal.MTLCommandBufferObjectFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCommandBuffer/heapProvider
func (c MPSCommandBuffer) HeapProvider() MPSHeapProvider {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("heapProvider"))
	return MPSHeapProviderObjectFromID(rv)
}
func (c MPSCommandBuffer) SetHeapProvider(value MPSHeapProvider) {
	objc.Send[struct{}](c.ID, objc.Sel("setHeapProvider:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCommandBuffer/predicate
func (c MPSCommandBuffer) Predicate() IMPSPredicate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("predicate"))
	return MPSPredicateFromID(objc.ID(rv))
}
func (c MPSCommandBuffer) SetPredicate(value IMPSPredicate) {
	objc.Send[struct{}](c.ID, objc.Sel("setPredicate:"), value)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSCommandBuffer/rootCommandBuffer
func (c MPSCommandBuffer) RootCommandBuffer() metal.MTLCommandBuffer {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("rootCommandBuffer"))
	return metal.MTLCommandBufferObjectFromID(rv)
}
