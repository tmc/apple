// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOGPUMetalAccelerationStructure] class.
var (
	_IOGPUMetalAccelerationStructureClass     IOGPUMetalAccelerationStructureClass
	_IOGPUMetalAccelerationStructureClassOnce sync.Once
)

func getIOGPUMetalAccelerationStructureClass() IOGPUMetalAccelerationStructureClass {
	_IOGPUMetalAccelerationStructureClassOnce.Do(func() {
		_IOGPUMetalAccelerationStructureClass = IOGPUMetalAccelerationStructureClass{class: objc.GetClass("IOGPUMetalAccelerationStructure")}
	})
	return _IOGPUMetalAccelerationStructureClass
}

// GetIOGPUMetalAccelerationStructureClass returns the class object for IOGPUMetalAccelerationStructure.
func GetIOGPUMetalAccelerationStructureClass() IOGPUMetalAccelerationStructureClass {
	return getIOGPUMetalAccelerationStructureClass()
}

type IOGPUMetalAccelerationStructureClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOGPUMetalAccelerationStructureClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOGPUMetalAccelerationStructureClass) Alloc() IOGPUMetalAccelerationStructure {
	rv := objc.SendIfResponds[IOGPUMetalAccelerationStructure](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [IOGPUMetalAccelerationStructure.AccelerationStructureUniqueIdentifier]
//   - [IOGPUMetalAccelerationStructure.Buffer]
//   - [IOGPUMetalAccelerationStructure.BufferOffset]
//   - [IOGPUMetalAccelerationStructure.CopyPropertiesFromBuffer]
//   - [IOGPUMetalAccelerationStructure.Descriptor]
//   - [IOGPUMetalAccelerationStructure.SetDescriptor]
//   - [IOGPUMetalAccelerationStructure.GpuHandle]
//   - [IOGPUMetalAccelerationStructure.GpuResourceID]
//   - [IOGPUMetalAccelerationStructure.ResourceIndex]
//   - [IOGPUMetalAccelerationStructure.Size]
//   - [IOGPUMetalAccelerationStructure.UniqueIdentifier]
//   - [IOGPUMetalAccelerationStructure.InitWithBufferOffset]
//   - [IOGPUMetalAccelerationStructure.InitWithBufferOffsetResourceIndex]
type IOGPUMetalAccelerationStructure struct {
	IOGPUMetalResource
}

// IOGPUMetalAccelerationStructureFromID constructs a [IOGPUMetalAccelerationStructure] from an objc.ID.
func IOGPUMetalAccelerationStructureFromID(id objc.ID) IOGPUMetalAccelerationStructure {
	return IOGPUMetalAccelerationStructure{IOGPUMetalResource: IOGPUMetalResourceFromID(id)}
}

// Ensure IOGPUMetalAccelerationStructure implements IIOGPUMetalAccelerationStructure.
var _ IIOGPUMetalAccelerationStructure = IOGPUMetalAccelerationStructure{}

// An interface definition for the [IOGPUMetalAccelerationStructure] class.
//
// # Methods
//
//   - [IIOGPUMetalAccelerationStructure.AccelerationStructureUniqueIdentifier]
//   - [IIOGPUMetalAccelerationStructure.Buffer]
//   - [IIOGPUMetalAccelerationStructure.BufferOffset]
//   - [IIOGPUMetalAccelerationStructure.CopyPropertiesFromBuffer]
//   - [IIOGPUMetalAccelerationStructure.Descriptor]
//   - [IIOGPUMetalAccelerationStructure.SetDescriptor]
//   - [IIOGPUMetalAccelerationStructure.GpuHandle]
//   - [IIOGPUMetalAccelerationStructure.GpuResourceID]
//   - [IIOGPUMetalAccelerationStructure.ResourceIndex]
//   - [IIOGPUMetalAccelerationStructure.Size]
//   - [IIOGPUMetalAccelerationStructure.UniqueIdentifier]
//   - [IIOGPUMetalAccelerationStructure.InitWithBufferOffset]
//   - [IIOGPUMetalAccelerationStructure.InitWithBufferOffsetResourceIndex]
type IIOGPUMetalAccelerationStructure interface {
	IIOGPUMetalResource

	// Topic: Methods

	AccelerationStructureUniqueIdentifier() uint64
	Buffer() IIOGPUMetalBuffer
	BufferOffset() uint64
	CopyPropertiesFromBuffer(buffer objectivec.IObject)
	Descriptor() metal.MTLAccelerationStructureDescriptor
	SetDescriptor(value metal.MTLAccelerationStructureDescriptor)
	GpuHandle() uint64
	GpuResourceID() metal.MTLResourceID
	ResourceIndex() uint64
	Size() uint64
	UniqueIdentifier() uint64
	InitWithBufferOffset(buffer objectivec.IObject, offset uint64) IOGPUMetalAccelerationStructure
	InitWithBufferOffsetResourceIndex(buffer objectivec.IObject, offset uint64, index uint64) IOGPUMetalAccelerationStructure
}

// Init initializes the instance.
func (i IOGPUMetalAccelerationStructure) Init() IOGPUMetalAccelerationStructure {
	rv := objc.SendIfResponds[IOGPUMetalAccelerationStructure](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IOGPUMetalAccelerationStructure) Autorelease() IOGPUMetalAccelerationStructure {
	rv := objc.SendIfResponds[IOGPUMetalAccelerationStructure](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOGPUMetalAccelerationStructure creates a new IOGPUMetalAccelerationStructure instance.
func NewIOGPUMetalAccelerationStructure() IOGPUMetalAccelerationStructure {
	class := getIOGPUMetalAccelerationStructureClass()
	rv := objc.SendIfResponds[IOGPUMetalAccelerationStructure](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGPUMetalAccelerationStructureMemorylessDescriptor(memoryless objectivec.IObject, descriptor objectivec.IObject) IOGPUMetalAccelerationStructure {
	instance := getIOGPUMetalAccelerationStructureClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initMemoryless:descriptor:"), memoryless, descriptor)
	return IOGPUMetalAccelerationStructureFromID(rv)
}

func NewGPUMetalAccelerationStructureStandinWithDevice(device objectivec.IObject) IOGPUMetalAccelerationStructure {
	instance := getIOGPUMetalAccelerationStructureClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initStandinWithDevice:"), device)
	return IOGPUMetalAccelerationStructureFromID(rv)
}

func NewGPUMetalAccelerationStructureWithBufferOffset(buffer objectivec.IObject, offset uint64) IOGPUMetalAccelerationStructure {
	instance := getIOGPUMetalAccelerationStructureClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithBuffer:offset:"), buffer, offset)
	return IOGPUMetalAccelerationStructureFromID(rv)
}

func NewGPUMetalAccelerationStructureWithBufferOffsetResourceIndex(buffer objectivec.IObject, offset uint64, index uint64) IOGPUMetalAccelerationStructure {
	instance := getIOGPUMetalAccelerationStructureClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithBuffer:offset:resourceIndex:"), buffer, offset, index)
	return IOGPUMetalAccelerationStructureFromID(rv)
}

func NewGPUMetalAccelerationStructureWithDeviceOptionsArgsArgsSize(device objectivec.IObject, options uint64, args *IOGPUNewResourceArgs, size uint32) IOGPUMetalAccelerationStructure {
	instance := getIOGPUMetalAccelerationStructureClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:options:args:argsSize:"), device, options, unsafe.Pointer(args), size)
	return IOGPUMetalAccelerationStructureFromID(rv)
}

func NewGPUMetalAccelerationStructureWithDeviceRemoteStorageResourceOptionsArgsArgsSize(device objectivec.IObject, resource objectivec.IObject, options uint64, args *IOGPUNewResourceArgs, size uint32) IOGPUMetalAccelerationStructure {
	instance := getIOGPUMetalAccelerationStructureClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:remoteStorageResource:options:args:argsSize:"), device, resource, options, unsafe.Pointer(args), size)
	return IOGPUMetalAccelerationStructureFromID(rv)
}

func NewGPUMetalAccelerationStructureWithResource(resource objectivec.IObject) IOGPUMetalAccelerationStructure {
	instance := getIOGPUMetalAccelerationStructureClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithResource:"), resource)
	return IOGPUMetalAccelerationStructureFromID(rv)
}

func (i IOGPUMetalAccelerationStructure) CopyPropertiesFromBuffer(buffer objectivec.IObject) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("copyPropertiesFromBuffer:"), buffer)
}
func (i IOGPUMetalAccelerationStructure) InitWithBufferOffset(buffer objectivec.IObject, offset uint64) IOGPUMetalAccelerationStructure {
	rv := objc.SendIfResponds[IOGPUMetalAccelerationStructure](i.ID, objc.Sel("initWithBuffer:offset:"), buffer, offset)
	return rv
}
func (i IOGPUMetalAccelerationStructure) InitWithBufferOffsetResourceIndex(buffer objectivec.IObject, offset uint64, index uint64) IOGPUMetalAccelerationStructure {
	rv := objc.SendIfResponds[IOGPUMetalAccelerationStructure](i.ID, objc.Sel("initWithBuffer:offset:resourceIndex:"), buffer, offset, index)
	return rv
}

func (i IOGPUMetalAccelerationStructure) AccelerationStructureUniqueIdentifier() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("accelerationStructureUniqueIdentifier"))
	return rv
}
func (i IOGPUMetalAccelerationStructure) Buffer() IIOGPUMetalBuffer {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("buffer"))
	return IOGPUMetalBufferFromID(objc.ID(rv))
}
func (i IOGPUMetalAccelerationStructure) BufferOffset() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("bufferOffset"))
	return rv
}
func (i IOGPUMetalAccelerationStructure) Descriptor() metal.MTLAccelerationStructureDescriptor {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("descriptor"))
	return metal.MTLAccelerationStructureDescriptorFromID(objc.ID(rv))
}
func (i IOGPUMetalAccelerationStructure) SetDescriptor(value metal.MTLAccelerationStructureDescriptor) {
	objc.SendIfResponds[struct{}](i.ID, objc.Sel("setDescriptor:"), value)
}
func (i IOGPUMetalAccelerationStructure) GpuHandle() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("gpuHandle"))
	return rv
}
func (i IOGPUMetalAccelerationStructure) GpuResourceID() metal.MTLResourceID {
	rv := objc.SendIfResponds[metal.MTLResourceID](i.ID, objc.Sel("gpuResourceID"))
	return metal.MTLResourceID(rv)
}
func (i IOGPUMetalAccelerationStructure) ResourceIndex() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("resourceIndex"))
	return rv
}
func (i IOGPUMetalAccelerationStructure) Size() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("size"))
	return rv
}
func (i IOGPUMetalAccelerationStructure) UniqueIdentifier() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("uniqueIdentifier"))
	return rv
}
