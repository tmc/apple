// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"sync"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOGPUMetal4MachineLearningCommandEncoder] class.
var (
	_IOGPUMetal4MachineLearningCommandEncoderClass     IOGPUMetal4MachineLearningCommandEncoderClass
	_IOGPUMetal4MachineLearningCommandEncoderClassOnce sync.Once
)

func getIOGPUMetal4MachineLearningCommandEncoderClass() IOGPUMetal4MachineLearningCommandEncoderClass {
	_IOGPUMetal4MachineLearningCommandEncoderClassOnce.Do(func() {
		_IOGPUMetal4MachineLearningCommandEncoderClass = IOGPUMetal4MachineLearningCommandEncoderClass{class: objc.GetClass("IOGPUMetal4MachineLearningCommandEncoder")}
	})
	return _IOGPUMetal4MachineLearningCommandEncoderClass
}

// GetIOGPUMetal4MachineLearningCommandEncoderClass returns the class object for IOGPUMetal4MachineLearningCommandEncoder.
func GetIOGPUMetal4MachineLearningCommandEncoderClass() IOGPUMetal4MachineLearningCommandEncoderClass {
	return getIOGPUMetal4MachineLearningCommandEncoderClass()
}

type IOGPUMetal4MachineLearningCommandEncoderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOGPUMetal4MachineLearningCommandEncoderClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOGPUMetal4MachineLearningCommandEncoderClass) Alloc() IOGPUMetal4MachineLearningCommandEncoder {
	rv := objc.SendIfResponds[IOGPUMetal4MachineLearningCommandEncoder](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [IOGPUMetal4MachineLearningCommandEncoder.GetType]
//   - [IOGPUMetal4MachineLearningCommandEncoder.InitWithCommandBufferAllocator]
type IOGPUMetal4MachineLearningCommandEncoder struct {
	metal.MTL4MachineLearningCommandEncoderObject
}

// IOGPUMetal4MachineLearningCommandEncoderFromID constructs a [IOGPUMetal4MachineLearningCommandEncoder] from an objc.ID.
func IOGPUMetal4MachineLearningCommandEncoderFromID(id objc.ID) IOGPUMetal4MachineLearningCommandEncoder {
	return IOGPUMetal4MachineLearningCommandEncoder{MTL4MachineLearningCommandEncoderObject: metal.MTL4MachineLearningCommandEncoderObjectFromID(id)}
}

// Ensure IOGPUMetal4MachineLearningCommandEncoder implements IIOGPUMetal4MachineLearningCommandEncoder.
var _ IIOGPUMetal4MachineLearningCommandEncoder = IOGPUMetal4MachineLearningCommandEncoder{}

// An interface definition for the [IOGPUMetal4MachineLearningCommandEncoder] class.
//
// # Methods
//
//   - [IIOGPUMetal4MachineLearningCommandEncoder.GetType]
//   - [IIOGPUMetal4MachineLearningCommandEncoder.InitWithCommandBufferAllocator]
type IIOGPUMetal4MachineLearningCommandEncoder interface {
	metal.MTL4MachineLearningCommandEncoder

	// Topic: Methods

	GetType() int64
	InitWithCommandBufferAllocator(buffer objectivec.IObject, allocator objectivec.IObject) IOGPUMetal4MachineLearningCommandEncoder
}

// Init initializes the instance.
func (i IOGPUMetal4MachineLearningCommandEncoder) Init() IOGPUMetal4MachineLearningCommandEncoder {
	rv := objc.SendIfResponds[IOGPUMetal4MachineLearningCommandEncoder](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IOGPUMetal4MachineLearningCommandEncoder) Autorelease() IOGPUMetal4MachineLearningCommandEncoder {
	rv := objc.SendIfResponds[IOGPUMetal4MachineLearningCommandEncoder](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOGPUMetal4MachineLearningCommandEncoder creates a new IOGPUMetal4MachineLearningCommandEncoder instance.
func NewIOGPUMetal4MachineLearningCommandEncoder() IOGPUMetal4MachineLearningCommandEncoder {
	class := getIOGPUMetal4MachineLearningCommandEncoderClass()
	rv := objc.SendIfResponds[IOGPUMetal4MachineLearningCommandEncoder](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGPUMetal4MachineLearningCommandEncoderWithCommandBufferAllocator(buffer objectivec.IObject, allocator objectivec.IObject) IOGPUMetal4MachineLearningCommandEncoder {
	instance := getIOGPUMetal4MachineLearningCommandEncoderClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCommandBuffer:allocator:"), buffer, allocator)
	return IOGPUMetal4MachineLearningCommandEncoderFromID(rv)
}

func (i IOGPUMetal4MachineLearningCommandEncoder) GetType() int64 {
	rv := objc.SendIfResponds[int64](i.ID, objc.Sel("getType"))
	return rv
}
func (i IOGPUMetal4MachineLearningCommandEncoder) InitWithCommandBufferAllocator(buffer objectivec.IObject, allocator objectivec.IObject) IOGPUMetal4MachineLearningCommandEncoder {
	rv := objc.SendIfResponds[IOGPUMetal4MachineLearningCommandEncoder](i.ID, objc.Sel("initWithCommandBuffer:allocator:"), buffer, allocator)
	return rv
}
