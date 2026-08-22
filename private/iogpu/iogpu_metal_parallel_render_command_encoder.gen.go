// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"sync"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOGPUMetalParallelRenderCommandEncoder] class.
var (
	_IOGPUMetalParallelRenderCommandEncoderClass     IOGPUMetalParallelRenderCommandEncoderClass
	_IOGPUMetalParallelRenderCommandEncoderClassOnce sync.Once
)

func getIOGPUMetalParallelRenderCommandEncoderClass() IOGPUMetalParallelRenderCommandEncoderClass {
	_IOGPUMetalParallelRenderCommandEncoderClassOnce.Do(func() {
		_IOGPUMetalParallelRenderCommandEncoderClass = IOGPUMetalParallelRenderCommandEncoderClass{class: objc.GetClass("IOGPUMetalParallelRenderCommandEncoder")}
	})
	return _IOGPUMetalParallelRenderCommandEncoderClass
}

// GetIOGPUMetalParallelRenderCommandEncoderClass returns the class object for IOGPUMetalParallelRenderCommandEncoder.
func GetIOGPUMetalParallelRenderCommandEncoderClass() IOGPUMetalParallelRenderCommandEncoderClass {
	return getIOGPUMetalParallelRenderCommandEncoderClass()
}

type IOGPUMetalParallelRenderCommandEncoderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOGPUMetalParallelRenderCommandEncoderClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOGPUMetalParallelRenderCommandEncoderClass) Alloc() IOGPUMetalParallelRenderCommandEncoder {
	rv := objc.SendIfResponds[IOGPUMetalParallelRenderCommandEncoder](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [IOGPUMetalParallelRenderCommandEncoder._renderCommandEncoderCommon]
//   - [IOGPUMetalParallelRenderCommandEncoder.GetType]
//   - [IOGPUMetalParallelRenderCommandEncoder.IsMemorylessRender]
//   - [IOGPUMetalParallelRenderCommandEncoder.InitWithCommandBufferRenderPassDescriptor]
type IOGPUMetalParallelRenderCommandEncoder struct {
	metal.MTLParallelRenderCommandEncoderObject
}

// IOGPUMetalParallelRenderCommandEncoderFromID constructs a [IOGPUMetalParallelRenderCommandEncoder] from an objc.ID.
func IOGPUMetalParallelRenderCommandEncoderFromID(id objc.ID) IOGPUMetalParallelRenderCommandEncoder {
	return IOGPUMetalParallelRenderCommandEncoder{MTLParallelRenderCommandEncoderObject: metal.MTLParallelRenderCommandEncoderObjectFromID(id)}
}

// Ensure IOGPUMetalParallelRenderCommandEncoder implements IIOGPUMetalParallelRenderCommandEncoder.
var _ IIOGPUMetalParallelRenderCommandEncoder = IOGPUMetalParallelRenderCommandEncoder{}

// An interface definition for the [IOGPUMetalParallelRenderCommandEncoder] class.
//
// # Methods
//
//   - [IIOGPUMetalParallelRenderCommandEncoder._renderCommandEncoderCommon]
//   - [IIOGPUMetalParallelRenderCommandEncoder.GetType]
//   - [IIOGPUMetalParallelRenderCommandEncoder.IsMemorylessRender]
//   - [IIOGPUMetalParallelRenderCommandEncoder.InitWithCommandBufferRenderPassDescriptor]
type IIOGPUMetalParallelRenderCommandEncoder interface {
	metal.MTLParallelRenderCommandEncoder

	// Topic: Methods

	_renderCommandEncoderCommon() objectivec.IObject
	GetType() uint64
	IsMemorylessRender() bool
	InitWithCommandBufferRenderPassDescriptor(buffer objectivec.IObject, descriptor objectivec.IObject) IOGPUMetalParallelRenderCommandEncoder
}

// Init initializes the instance.
func (i IOGPUMetalParallelRenderCommandEncoder) Init() IOGPUMetalParallelRenderCommandEncoder {
	rv := objc.SendIfResponds[IOGPUMetalParallelRenderCommandEncoder](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IOGPUMetalParallelRenderCommandEncoder) Autorelease() IOGPUMetalParallelRenderCommandEncoder {
	rv := objc.SendIfResponds[IOGPUMetalParallelRenderCommandEncoder](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOGPUMetalParallelRenderCommandEncoder creates a new IOGPUMetalParallelRenderCommandEncoder instance.
func NewIOGPUMetalParallelRenderCommandEncoder() IOGPUMetalParallelRenderCommandEncoder {
	class := getIOGPUMetalParallelRenderCommandEncoderClass()
	rv := objc.SendIfResponds[IOGPUMetalParallelRenderCommandEncoder](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGPUMetalParallelRenderCommandEncoderWithCommandBufferRenderPassDescriptor(buffer objectivec.IObject, descriptor objectivec.IObject) IOGPUMetalParallelRenderCommandEncoder {
	instance := getIOGPUMetalParallelRenderCommandEncoderClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCommandBuffer:renderPassDescriptor:"), buffer, descriptor)
	return IOGPUMetalParallelRenderCommandEncoderFromID(rv)
}

func (i IOGPUMetalParallelRenderCommandEncoder) _renderCommandEncoderCommon() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("_renderCommandEncoderCommon"))
	return objectivec.Object{ID: rv}
}

// RenderCommandEncoderCommon is an exported wrapper for the private method _renderCommandEncoderCommon.
func (i IOGPUMetalParallelRenderCommandEncoder) RenderCommandEncoderCommon() (objectivec.IObject, error) {
	if !objc.RespondsToSelector(i.ID, objc.Sel("_renderCommandEncoderCommon")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_renderCommandEncoderCommon"}
		return nil, err
	}
	return i._renderCommandEncoderCommon(), nil
}

// CanRenderCommandEncoderCommon reports whether the receiver responds to the private selector _renderCommandEncoderCommon.
func (i IOGPUMetalParallelRenderCommandEncoder) CanRenderCommandEncoderCommon() bool {
	return objc.RespondsToSelector(i.ID, objc.Sel("_renderCommandEncoderCommon"))
}
func (i IOGPUMetalParallelRenderCommandEncoder) GetType() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("getType"))
	return rv
}
func (i IOGPUMetalParallelRenderCommandEncoder) IsMemorylessRender() bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("isMemorylessRender"))
	return rv
}
func (i IOGPUMetalParallelRenderCommandEncoder) InitWithCommandBufferRenderPassDescriptor(buffer objectivec.IObject, descriptor objectivec.IObject) IOGPUMetalParallelRenderCommandEncoder {
	rv := objc.SendIfResponds[IOGPUMetalParallelRenderCommandEncoder](i.ID, objc.Sel("initWithCommandBuffer:renderPassDescriptor:"), buffer, descriptor)
	return rv
}
