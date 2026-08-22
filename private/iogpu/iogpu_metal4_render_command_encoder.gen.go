// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"sync"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOGPUMetal4RenderCommandEncoder] class.
var (
	_IOGPUMetal4RenderCommandEncoderClass     IOGPUMetal4RenderCommandEncoderClass
	_IOGPUMetal4RenderCommandEncoderClassOnce sync.Once
)

func getIOGPUMetal4RenderCommandEncoderClass() IOGPUMetal4RenderCommandEncoderClass {
	_IOGPUMetal4RenderCommandEncoderClassOnce.Do(func() {
		_IOGPUMetal4RenderCommandEncoderClass = IOGPUMetal4RenderCommandEncoderClass{class: objc.GetClass("IOGPUMetal4RenderCommandEncoder")}
	})
	return _IOGPUMetal4RenderCommandEncoderClass
}

// GetIOGPUMetal4RenderCommandEncoderClass returns the class object for IOGPUMetal4RenderCommandEncoder.
func GetIOGPUMetal4RenderCommandEncoderClass() IOGPUMetal4RenderCommandEncoderClass {
	return getIOGPUMetal4RenderCommandEncoderClass()
}

type IOGPUMetal4RenderCommandEncoderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOGPUMetal4RenderCommandEncoderClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOGPUMetal4RenderCommandEncoderClass) Alloc() IOGPUMetal4RenderCommandEncoder {
	rv := objc.SendIfResponds[IOGPUMetal4RenderCommandEncoder](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [IOGPUMetal4RenderCommandEncoder.GetType]
//   - [IOGPUMetal4RenderCommandEncoder.InitWithCommandAllocator]
type IOGPUMetal4RenderCommandEncoder struct {
	metal.MTL4RenderCommandEncoderObject
}

// IOGPUMetal4RenderCommandEncoderFromID constructs a [IOGPUMetal4RenderCommandEncoder] from an objc.ID.
func IOGPUMetal4RenderCommandEncoderFromID(id objc.ID) IOGPUMetal4RenderCommandEncoder {
	return IOGPUMetal4RenderCommandEncoder{MTL4RenderCommandEncoderObject: metal.MTL4RenderCommandEncoderObjectFromID(id)}
}

// Ensure IOGPUMetal4RenderCommandEncoder implements IIOGPUMetal4RenderCommandEncoder.
var _ IIOGPUMetal4RenderCommandEncoder = IOGPUMetal4RenderCommandEncoder{}

// An interface definition for the [IOGPUMetal4RenderCommandEncoder] class.
//
// # Methods
//
//   - [IIOGPUMetal4RenderCommandEncoder.GetType]
//   - [IIOGPUMetal4RenderCommandEncoder.InitWithCommandAllocator]
type IIOGPUMetal4RenderCommandEncoder interface {
	metal.MTL4RenderCommandEncoder

	// Topic: Methods

	GetType() int64
	InitWithCommandAllocator(allocator objectivec.IObject) IOGPUMetal4RenderCommandEncoder
}

// Init initializes the instance.
func (i IOGPUMetal4RenderCommandEncoder) Init() IOGPUMetal4RenderCommandEncoder {
	rv := objc.SendIfResponds[IOGPUMetal4RenderCommandEncoder](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IOGPUMetal4RenderCommandEncoder) Autorelease() IOGPUMetal4RenderCommandEncoder {
	rv := objc.SendIfResponds[IOGPUMetal4RenderCommandEncoder](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOGPUMetal4RenderCommandEncoder creates a new IOGPUMetal4RenderCommandEncoder instance.
func NewIOGPUMetal4RenderCommandEncoder() IOGPUMetal4RenderCommandEncoder {
	class := getIOGPUMetal4RenderCommandEncoderClass()
	rv := objc.SendIfResponds[IOGPUMetal4RenderCommandEncoder](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGPUMetal4RenderCommandEncoderWithCommandAllocator(allocator objectivec.IObject) IOGPUMetal4RenderCommandEncoder {
	instance := getIOGPUMetal4RenderCommandEncoderClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCommandAllocator:"), allocator)
	return IOGPUMetal4RenderCommandEncoderFromID(rv)
}

func (i IOGPUMetal4RenderCommandEncoder) GetType() int64 {
	rv := objc.SendIfResponds[int64](i.ID, objc.Sel("getType"))
	return rv
}
func (i IOGPUMetal4RenderCommandEncoder) InitWithCommandAllocator(allocator objectivec.IObject) IOGPUMetal4RenderCommandEncoder {
	rv := objc.SendIfResponds[IOGPUMetal4RenderCommandEncoder](i.ID, objc.Sel("initWithCommandAllocator:"), allocator)
	return rv
}
