// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"sync"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOGPUMetal4ComputeCommandEncoder] class.
var (
	_IOGPUMetal4ComputeCommandEncoderClass     IOGPUMetal4ComputeCommandEncoderClass
	_IOGPUMetal4ComputeCommandEncoderClassOnce sync.Once
)

func getIOGPUMetal4ComputeCommandEncoderClass() IOGPUMetal4ComputeCommandEncoderClass {
	_IOGPUMetal4ComputeCommandEncoderClassOnce.Do(func() {
		_IOGPUMetal4ComputeCommandEncoderClass = IOGPUMetal4ComputeCommandEncoderClass{class: objc.GetClass("IOGPUMetal4ComputeCommandEncoder")}
	})
	return _IOGPUMetal4ComputeCommandEncoderClass
}

// GetIOGPUMetal4ComputeCommandEncoderClass returns the class object for IOGPUMetal4ComputeCommandEncoder.
func GetIOGPUMetal4ComputeCommandEncoderClass() IOGPUMetal4ComputeCommandEncoderClass {
	return getIOGPUMetal4ComputeCommandEncoderClass()
}

type IOGPUMetal4ComputeCommandEncoderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOGPUMetal4ComputeCommandEncoderClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOGPUMetal4ComputeCommandEncoderClass) Alloc() IOGPUMetal4ComputeCommandEncoder {
	rv := objc.SendIfResponds[IOGPUMetal4ComputeCommandEncoder](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [IOGPUMetal4ComputeCommandEncoder.GetType]
//   - [IOGPUMetal4ComputeCommandEncoder.InitWithCommandAllocator]
type IOGPUMetal4ComputeCommandEncoder struct {
	metal.MTL4ComputeCommandEncoderObject
}

// IOGPUMetal4ComputeCommandEncoderFromID constructs a [IOGPUMetal4ComputeCommandEncoder] from an objc.ID.
func IOGPUMetal4ComputeCommandEncoderFromID(id objc.ID) IOGPUMetal4ComputeCommandEncoder {
	return IOGPUMetal4ComputeCommandEncoder{MTL4ComputeCommandEncoderObject: metal.MTL4ComputeCommandEncoderObjectFromID(id)}
}

// Ensure IOGPUMetal4ComputeCommandEncoder implements IIOGPUMetal4ComputeCommandEncoder.
var _ IIOGPUMetal4ComputeCommandEncoder = IOGPUMetal4ComputeCommandEncoder{}

// An interface definition for the [IOGPUMetal4ComputeCommandEncoder] class.
//
// # Methods
//
//   - [IIOGPUMetal4ComputeCommandEncoder.GetType]
//   - [IIOGPUMetal4ComputeCommandEncoder.InitWithCommandAllocator]
type IIOGPUMetal4ComputeCommandEncoder interface {
	metal.MTL4ComputeCommandEncoder

	// Topic: Methods

	GetType() int64
	InitWithCommandAllocator(allocator objectivec.IObject) IOGPUMetal4ComputeCommandEncoder
}

// Init initializes the instance.
func (i IOGPUMetal4ComputeCommandEncoder) Init() IOGPUMetal4ComputeCommandEncoder {
	rv := objc.SendIfResponds[IOGPUMetal4ComputeCommandEncoder](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IOGPUMetal4ComputeCommandEncoder) Autorelease() IOGPUMetal4ComputeCommandEncoder {
	rv := objc.SendIfResponds[IOGPUMetal4ComputeCommandEncoder](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOGPUMetal4ComputeCommandEncoder creates a new IOGPUMetal4ComputeCommandEncoder instance.
func NewIOGPUMetal4ComputeCommandEncoder() IOGPUMetal4ComputeCommandEncoder {
	class := getIOGPUMetal4ComputeCommandEncoderClass()
	rv := objc.SendIfResponds[IOGPUMetal4ComputeCommandEncoder](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGPUMetal4ComputeCommandEncoderWithCommandAllocator(allocator objectivec.IObject) IOGPUMetal4ComputeCommandEncoder {
	instance := getIOGPUMetal4ComputeCommandEncoderClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCommandAllocator:"), allocator)
	return IOGPUMetal4ComputeCommandEncoderFromID(rv)
}

func (i IOGPUMetal4ComputeCommandEncoder) GetType() int64 {
	rv := objc.SendIfResponds[int64](i.ID, objc.Sel("getType"))
	return rv
}
func (i IOGPUMetal4ComputeCommandEncoder) InitWithCommandAllocator(allocator objectivec.IObject) IOGPUMetal4ComputeCommandEncoder {
	rv := objc.SendIfResponds[IOGPUMetal4ComputeCommandEncoder](i.ID, objc.Sel("initWithCommandAllocator:"), allocator)
	return rv
}
