// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"sync"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOGPUMetalCommandEncoder] class.
var (
	_IOGPUMetalCommandEncoderClass     IOGPUMetalCommandEncoderClass
	_IOGPUMetalCommandEncoderClassOnce sync.Once
)

func getIOGPUMetalCommandEncoderClass() IOGPUMetalCommandEncoderClass {
	_IOGPUMetalCommandEncoderClassOnce.Do(func() {
		_IOGPUMetalCommandEncoderClass = IOGPUMetalCommandEncoderClass{class: objc.GetClass("IOGPUMetalCommandEncoder")}
	})
	return _IOGPUMetalCommandEncoderClass
}

// GetIOGPUMetalCommandEncoderClass returns the class object for IOGPUMetalCommandEncoder.
func GetIOGPUMetalCommandEncoderClass() IOGPUMetalCommandEncoderClass {
	return getIOGPUMetalCommandEncoderClass()
}

type IOGPUMetalCommandEncoderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOGPUMetalCommandEncoderClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOGPUMetalCommandEncoderClass) Alloc() IOGPUMetalCommandEncoder {
	rv := objc.SendIfResponds[IOGPUMetalCommandEncoder](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [IOGPUMetalCommandEncoder.InitWithCommandBuffer]
type IOGPUMetalCommandEncoder struct {
	metal.MTLCommandEncoderObject
}

// IOGPUMetalCommandEncoderFromID constructs a [IOGPUMetalCommandEncoder] from an objc.ID.
func IOGPUMetalCommandEncoderFromID(id objc.ID) IOGPUMetalCommandEncoder {
	return IOGPUMetalCommandEncoder{MTLCommandEncoderObject: metal.MTLCommandEncoderObjectFromID(id)}
}

// Ensure IOGPUMetalCommandEncoder implements IIOGPUMetalCommandEncoder.
var _ IIOGPUMetalCommandEncoder = IOGPUMetalCommandEncoder{}

// An interface definition for the [IOGPUMetalCommandEncoder] class.
//
// # Methods
//
//   - [IIOGPUMetalCommandEncoder.InitWithCommandBuffer]
type IIOGPUMetalCommandEncoder interface {
	metal.MTLCommandEncoder

	// Topic: Methods

	InitWithCommandBuffer(buffer objectivec.IObject) IOGPUMetalCommandEncoder
}

// Init initializes the instance.
func (i IOGPUMetalCommandEncoder) Init() IOGPUMetalCommandEncoder {
	rv := objc.SendIfResponds[IOGPUMetalCommandEncoder](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IOGPUMetalCommandEncoder) Autorelease() IOGPUMetalCommandEncoder {
	rv := objc.SendIfResponds[IOGPUMetalCommandEncoder](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOGPUMetalCommandEncoder creates a new IOGPUMetalCommandEncoder instance.
func NewIOGPUMetalCommandEncoder() IOGPUMetalCommandEncoder {
	class := getIOGPUMetalCommandEncoderClass()
	rv := objc.SendIfResponds[IOGPUMetalCommandEncoder](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGPUMetalCommandEncoderWithCommandBuffer(buffer objectivec.IObject) IOGPUMetalCommandEncoder {
	instance := getIOGPUMetalCommandEncoderClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCommandBuffer:"), buffer)
	return IOGPUMetalCommandEncoderFromID(rv)
}

func (i IOGPUMetalCommandEncoder) InitWithCommandBuffer(buffer objectivec.IObject) IOGPUMetalCommandEncoder {
	rv := objc.SendIfResponds[IOGPUMetalCommandEncoder](i.ID, objc.Sel("initWithCommandBuffer:"), buffer)
	return rv
}
