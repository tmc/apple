// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOGPUMetalIndirectCommandBuffer] class.
var (
	_IOGPUMetalIndirectCommandBufferClass     IOGPUMetalIndirectCommandBufferClass
	_IOGPUMetalIndirectCommandBufferClassOnce sync.Once
)

func getIOGPUMetalIndirectCommandBufferClass() IOGPUMetalIndirectCommandBufferClass {
	_IOGPUMetalIndirectCommandBufferClassOnce.Do(func() {
		_IOGPUMetalIndirectCommandBufferClass = IOGPUMetalIndirectCommandBufferClass{class: objc.GetClass("IOGPUMetalIndirectCommandBuffer")}
	})
	return _IOGPUMetalIndirectCommandBufferClass
}

// GetIOGPUMetalIndirectCommandBufferClass returns the class object for IOGPUMetalIndirectCommandBuffer.
func GetIOGPUMetalIndirectCommandBufferClass() IOGPUMetalIndirectCommandBufferClass {
	return getIOGPUMetalIndirectCommandBufferClass()
}

type IOGPUMetalIndirectCommandBufferClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOGPUMetalIndirectCommandBufferClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOGPUMetalIndirectCommandBufferClass) Alloc() IOGPUMetalIndirectCommandBuffer {
	rv := objc.SendIfResponds[IOGPUMetalIndirectCommandBuffer](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [IOGPUMetalIndirectCommandBuffer.CommandBufferType]
//   - [IOGPUMetalIndirectCommandBuffer.Descriptor]
//   - [IOGPUMetalIndirectCommandBuffer.GetHeaderHeaderSize]
//   - [IOGPUMetalIndirectCommandBuffer.GpuHandle]
//   - [IOGPUMetalIndirectCommandBuffer.GpuResourceID]
//   - [IOGPUMetalIndirectCommandBuffer.IndirectComputeCommandAtIndex]
//   - [IOGPUMetalIndirectCommandBuffer.IndirectRenderCommandAtIndex]
//   - [IOGPUMetalIndirectCommandBuffer.PrivateICBuffer]
//   - [IOGPUMetalIndirectCommandBuffer.ResetWithRange]
//   - [IOGPUMetalIndirectCommandBuffer.Size]
//   - [IOGPUMetalIndirectCommandBuffer.UniqueIdentifier]
//   - [IOGPUMetalIndirectCommandBuffer.InitWithBufferDescriptorMaxCommandCount]
type IOGPUMetalIndirectCommandBuffer struct {
	IOGPUMetalResource
}

// IOGPUMetalIndirectCommandBufferFromID constructs a [IOGPUMetalIndirectCommandBuffer] from an objc.ID.
func IOGPUMetalIndirectCommandBufferFromID(id objc.ID) IOGPUMetalIndirectCommandBuffer {
	return IOGPUMetalIndirectCommandBuffer{IOGPUMetalResource: IOGPUMetalResourceFromID(id)}
}

// Ensure IOGPUMetalIndirectCommandBuffer implements IIOGPUMetalIndirectCommandBuffer.
var _ IIOGPUMetalIndirectCommandBuffer = IOGPUMetalIndirectCommandBuffer{}

// An interface definition for the [IOGPUMetalIndirectCommandBuffer] class.
//
// # Methods
//
//   - [IIOGPUMetalIndirectCommandBuffer.CommandBufferType]
//   - [IIOGPUMetalIndirectCommandBuffer.Descriptor]
//   - [IIOGPUMetalIndirectCommandBuffer.GetHeaderHeaderSize]
//   - [IIOGPUMetalIndirectCommandBuffer.GpuHandle]
//   - [IIOGPUMetalIndirectCommandBuffer.GpuResourceID]
//   - [IIOGPUMetalIndirectCommandBuffer.IndirectComputeCommandAtIndex]
//   - [IIOGPUMetalIndirectCommandBuffer.IndirectRenderCommandAtIndex]
//   - [IIOGPUMetalIndirectCommandBuffer.PrivateICBuffer]
//   - [IIOGPUMetalIndirectCommandBuffer.ResetWithRange]
//   - [IIOGPUMetalIndirectCommandBuffer.Size]
//   - [IIOGPUMetalIndirectCommandBuffer.UniqueIdentifier]
//   - [IIOGPUMetalIndirectCommandBuffer.InitWithBufferDescriptorMaxCommandCount]
type IIOGPUMetalIndirectCommandBuffer interface {
	IIOGPUMetalResource

	// Topic: Methods

	CommandBufferType() uint64
	Descriptor() metal.MTLIndirectCommandBufferDescriptor
	GetHeaderHeaderSize(header unsafe.Pointer, size *uint64)
	GpuHandle() uint64
	GpuResourceID() metal.MTLResourceID
	IndirectComputeCommandAtIndex(index uint64) objectivec.IObject
	IndirectRenderCommandAtIndex(index uint64) objectivec.IObject
	PrivateICBuffer() IIOGPUMetalBuffer
	ResetWithRange(range_ foundation.NSRange)
	Size() uint64
	UniqueIdentifier() uint64
	InitWithBufferDescriptorMaxCommandCount(buffer objectivec.IObject, descriptor objectivec.IObject, count uint64) IOGPUMetalIndirectCommandBuffer
}

// Init initializes the instance.
func (i IOGPUMetalIndirectCommandBuffer) Init() IOGPUMetalIndirectCommandBuffer {
	rv := objc.SendIfResponds[IOGPUMetalIndirectCommandBuffer](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IOGPUMetalIndirectCommandBuffer) Autorelease() IOGPUMetalIndirectCommandBuffer {
	rv := objc.SendIfResponds[IOGPUMetalIndirectCommandBuffer](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOGPUMetalIndirectCommandBuffer creates a new IOGPUMetalIndirectCommandBuffer instance.
func NewIOGPUMetalIndirectCommandBuffer() IOGPUMetalIndirectCommandBuffer {
	class := getIOGPUMetalIndirectCommandBufferClass()
	rv := objc.SendIfResponds[IOGPUMetalIndirectCommandBuffer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGPUMetalIndirectCommandBufferMemorylessDescriptor(memoryless objectivec.IObject, descriptor objectivec.IObject) IOGPUMetalIndirectCommandBuffer {
	instance := getIOGPUMetalIndirectCommandBufferClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initMemoryless:descriptor:"), memoryless, descriptor)
	return IOGPUMetalIndirectCommandBufferFromID(rv)
}

func NewGPUMetalIndirectCommandBufferStandinWithDevice(device objectivec.IObject) IOGPUMetalIndirectCommandBuffer {
	instance := getIOGPUMetalIndirectCommandBufferClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initStandinWithDevice:"), device)
	return IOGPUMetalIndirectCommandBufferFromID(rv)
}

func NewGPUMetalIndirectCommandBufferWithBufferDescriptorMaxCommandCount(buffer objectivec.IObject, descriptor objectivec.IObject, count uint64) IOGPUMetalIndirectCommandBuffer {
	instance := getIOGPUMetalIndirectCommandBufferClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithBuffer:descriptor:maxCommandCount:"), buffer, descriptor, count)
	return IOGPUMetalIndirectCommandBufferFromID(rv)
}

func NewGPUMetalIndirectCommandBufferWithDeviceOptionsArgsArgsSize(device objectivec.IObject, options uint64, args *IOGPUNewResourceArgs, size uint32) IOGPUMetalIndirectCommandBuffer {
	instance := getIOGPUMetalIndirectCommandBufferClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:options:args:argsSize:"), device, options, unsafe.Pointer(args), size)
	return IOGPUMetalIndirectCommandBufferFromID(rv)
}

func NewGPUMetalIndirectCommandBufferWithDeviceRemoteStorageResourceOptionsArgsArgsSize(device objectivec.IObject, resource objectivec.IObject, options uint64, args *IOGPUNewResourceArgs, size uint32) IOGPUMetalIndirectCommandBuffer {
	instance := getIOGPUMetalIndirectCommandBufferClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:remoteStorageResource:options:args:argsSize:"), device, resource, options, unsafe.Pointer(args), size)
	return IOGPUMetalIndirectCommandBufferFromID(rv)
}

func NewGPUMetalIndirectCommandBufferWithResource(resource objectivec.IObject) IOGPUMetalIndirectCommandBuffer {
	instance := getIOGPUMetalIndirectCommandBufferClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithResource:"), resource)
	return IOGPUMetalIndirectCommandBufferFromID(rv)
}

func (i IOGPUMetalIndirectCommandBuffer) GetHeaderHeaderSize(header unsafe.Pointer, size *uint64) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("getHeader:headerSize:"), header, unsafe.Pointer(size))
}
func (i IOGPUMetalIndirectCommandBuffer) IndirectComputeCommandAtIndex(index uint64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("indirectComputeCommandAtIndex:"), index)
	return objectivec.Object{ID: rv}
}
func (i IOGPUMetalIndirectCommandBuffer) IndirectRenderCommandAtIndex(index uint64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("indirectRenderCommandAtIndex:"), index)
	return objectivec.Object{ID: rv}
}
func (i IOGPUMetalIndirectCommandBuffer) ResetWithRange(range_ foundation.NSRange) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("resetWithRange:"), range_)
}
func (i IOGPUMetalIndirectCommandBuffer) UniqueIdentifier() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("uniqueIdentifier"))
	return rv
}
func (i IOGPUMetalIndirectCommandBuffer) InitWithBufferDescriptorMaxCommandCount(buffer objectivec.IObject, descriptor objectivec.IObject, count uint64) IOGPUMetalIndirectCommandBuffer {
	rv := objc.SendIfResponds[IOGPUMetalIndirectCommandBuffer](i.ID, objc.Sel("initWithBuffer:descriptor:maxCommandCount:"), buffer, descriptor, count)
	return rv
}

func (i IOGPUMetalIndirectCommandBuffer) CommandBufferType() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("commandBufferType"))
	return rv
}
func (i IOGPUMetalIndirectCommandBuffer) Descriptor() metal.MTLIndirectCommandBufferDescriptor {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("descriptor"))
	return metal.MTLIndirectCommandBufferDescriptorFromID(objc.ID(rv))
}
func (i IOGPUMetalIndirectCommandBuffer) GpuHandle() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("gpuHandle"))
	return rv
}
func (i IOGPUMetalIndirectCommandBuffer) GpuResourceID() metal.MTLResourceID {
	rv := objc.SendIfResponds[metal.MTLResourceID](i.ID, objc.Sel("gpuResourceID"))
	return metal.MTLResourceID(rv)
}
func (i IOGPUMetalIndirectCommandBuffer) PrivateICBuffer() IIOGPUMetalBuffer {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("privateICBuffer"))
	return IOGPUMetalBufferFromID(objc.ID(rv))
}
func (i IOGPUMetalIndirectCommandBuffer) Size() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("size"))
	return rv
}
