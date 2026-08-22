// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOGPUMetal4CommandAllocator] class.
var (
	_IOGPUMetal4CommandAllocatorClass     IOGPUMetal4CommandAllocatorClass
	_IOGPUMetal4CommandAllocatorClassOnce sync.Once
)

func getIOGPUMetal4CommandAllocatorClass() IOGPUMetal4CommandAllocatorClass {
	_IOGPUMetal4CommandAllocatorClassOnce.Do(func() {
		_IOGPUMetal4CommandAllocatorClass = IOGPUMetal4CommandAllocatorClass{class: objc.GetClass("IOGPUMetal4CommandAllocator")}
	})
	return _IOGPUMetal4CommandAllocatorClass
}

// GetIOGPUMetal4CommandAllocatorClass returns the class object for IOGPUMetal4CommandAllocator.
func GetIOGPUMetal4CommandAllocatorClass() IOGPUMetal4CommandAllocatorClass {
	return getIOGPUMetal4CommandAllocatorClass()
}

type IOGPUMetal4CommandAllocatorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOGPUMetal4CommandAllocatorClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOGPUMetal4CommandAllocatorClass) Alloc() IOGPUMetal4CommandAllocator {
	rv := objc.SendIfResponds[IOGPUMetal4CommandAllocator](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [IOGPUMetal4CommandAllocator.GetCommandBufferStorageRetainReferences]
//   - [IOGPUMetal4CommandAllocator.GetGeneration]
//   - [IOGPUMetal4CommandAllocator.ReturnCommandBufferStorageCommandAllocatorGeneration]
//   - [IOGPUMetal4CommandAllocator.SetCurrentCommandEncoder]
//   - [IOGPUMetal4CommandAllocator.SetHwResourcePoolCount]
//   - [IOGPUMetal4CommandAllocator.InitAllocatorWithDevice]
//   - [IOGPUMetal4CommandAllocator.InitWithDevice]
//   - [IOGPUMetal4CommandAllocator.InitWithDeviceDescriptor]
//   - [IOGPUMetal4CommandAllocator.InitWithDeviceAndAliasToDevicePools]
type IOGPUMetal4CommandAllocator struct {
	metal.MTL4CommandAllocatorObject
}

// IOGPUMetal4CommandAllocatorFromID constructs a [IOGPUMetal4CommandAllocator] from an objc.ID.
func IOGPUMetal4CommandAllocatorFromID(id objc.ID) IOGPUMetal4CommandAllocator {
	return IOGPUMetal4CommandAllocator{MTL4CommandAllocatorObject: metal.MTL4CommandAllocatorObjectFromID(id)}
}

// Ensure IOGPUMetal4CommandAllocator implements IIOGPUMetal4CommandAllocator.
var _ IIOGPUMetal4CommandAllocator = IOGPUMetal4CommandAllocator{}

// An interface definition for the [IOGPUMetal4CommandAllocator] class.
//
// # Methods
//
//   - [IIOGPUMetal4CommandAllocator.GetCommandBufferStorageRetainReferences]
//   - [IIOGPUMetal4CommandAllocator.GetGeneration]
//   - [IIOGPUMetal4CommandAllocator.ReturnCommandBufferStorageCommandAllocatorGeneration]
//   - [IIOGPUMetal4CommandAllocator.SetCurrentCommandEncoder]
//   - [IIOGPUMetal4CommandAllocator.SetHwResourcePoolCount]
//   - [IIOGPUMetal4CommandAllocator.InitAllocatorWithDevice]
//   - [IIOGPUMetal4CommandAllocator.InitWithDevice]
//   - [IIOGPUMetal4CommandAllocator.InitWithDeviceDescriptor]
//   - [IIOGPUMetal4CommandAllocator.InitWithDeviceAndAliasToDevicePools]
type IIOGPUMetal4CommandAllocator interface {
	metal.MTL4CommandAllocator

	// Topic: Methods

	GetCommandBufferStorageRetainReferences(storage uint64, references bool) *IOGPUMetalCommandBufferStorage
	GetGeneration() uint32
	ReturnCommandBufferStorageCommandAllocatorGeneration(storage *IOGPUMetalCommandBufferStorage, generation uint32)
	SetCurrentCommandEncoder(encoder objectivec.IObject)
	SetHwResourcePoolCount(pool []objectivec.IObject, count uint32)
	InitAllocatorWithDevice(device objectivec.IObject) IOGPUMetal4CommandAllocator
	InitWithDevice(device objectivec.IObject) IOGPUMetal4CommandAllocator
	InitWithDeviceDescriptor(device objectivec.IObject, descriptor objectivec.IObject) IOGPUMetal4CommandAllocator
	InitWithDeviceAndAliasToDevicePools(pools objectivec.IObject) IOGPUMetal4CommandAllocator
}

// Init initializes the instance.
func (i IOGPUMetal4CommandAllocator) Init() IOGPUMetal4CommandAllocator {
	rv := objc.SendIfResponds[IOGPUMetal4CommandAllocator](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IOGPUMetal4CommandAllocator) Autorelease() IOGPUMetal4CommandAllocator {
	rv := objc.SendIfResponds[IOGPUMetal4CommandAllocator](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOGPUMetal4CommandAllocator creates a new IOGPUMetal4CommandAllocator instance.
func NewIOGPUMetal4CommandAllocator() IOGPUMetal4CommandAllocator {
	class := getIOGPUMetal4CommandAllocatorClass()
	rv := objc.SendIfResponds[IOGPUMetal4CommandAllocator](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGPUMetal4CommandAllocatorWithDevice(device objectivec.IObject) IOGPUMetal4CommandAllocator {
	instance := getIOGPUMetal4CommandAllocatorClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return IOGPUMetal4CommandAllocatorFromID(rv)
}

func NewGPUMetal4CommandAllocatorWithDeviceAndAliasToDevicePools(pools objectivec.IObject) IOGPUMetal4CommandAllocator {
	instance := getIOGPUMetal4CommandAllocatorClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDeviceAndAliasToDevicePools:"), pools)
	return IOGPUMetal4CommandAllocatorFromID(rv)
}

func NewGPUMetal4CommandAllocatorWithDeviceDescriptor(device objectivec.IObject, descriptor objectivec.IObject) IOGPUMetal4CommandAllocator {
	instance := getIOGPUMetal4CommandAllocatorClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:descriptor:"), device, descriptor)
	return IOGPUMetal4CommandAllocatorFromID(rv)
}

func (i IOGPUMetal4CommandAllocator) GetCommandBufferStorageRetainReferences(storage uint64, references bool) *IOGPUMetalCommandBufferStorage {
	rv := objc.SendIfResponds[unsafe.Pointer](i.ID, objc.Sel("getCommandBufferStorage:retainReferences:"), storage, references)
	return (*IOGPUMetalCommandBufferStorage)(rv)
}
func (i IOGPUMetal4CommandAllocator) GetGeneration() uint32 {
	rv := objc.SendIfResponds[uint32](i.ID, objc.Sel("getGeneration"))
	return rv
}
func (i IOGPUMetal4CommandAllocator) ReturnCommandBufferStorageCommandAllocatorGeneration(storage *IOGPUMetalCommandBufferStorage, generation uint32) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("returnCommandBufferStorage:commandAllocatorGeneration:"), unsafe.Pointer(storage), generation)
}
func (i IOGPUMetal4CommandAllocator) SetCurrentCommandEncoder(encoder objectivec.IObject) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setCurrentCommandEncoder:"), encoder)
}
func (i IOGPUMetal4CommandAllocator) SetHwResourcePoolCount(pool []objectivec.IObject, count uint32) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setHwResourcePool:count:"), objc.CArray(pool), count)
}
func (i IOGPUMetal4CommandAllocator) InitAllocatorWithDevice(device objectivec.IObject) IOGPUMetal4CommandAllocator {
	rv := objc.SendIfResponds[IOGPUMetal4CommandAllocator](i.ID, objc.Sel("initAllocatorWithDevice:"), device)
	return rv
}
func (i IOGPUMetal4CommandAllocator) InitWithDevice(device objectivec.IObject) IOGPUMetal4CommandAllocator {
	rv := objc.SendIfResponds[IOGPUMetal4CommandAllocator](i.ID, objc.Sel("initWithDevice:"), device)
	return rv
}
func (i IOGPUMetal4CommandAllocator) InitWithDeviceDescriptor(device objectivec.IObject, descriptor objectivec.IObject) IOGPUMetal4CommandAllocator {
	rv := objc.SendIfResponds[IOGPUMetal4CommandAllocator](i.ID, objc.Sel("initWithDevice:descriptor:"), device, descriptor)
	return rv
}
func (i IOGPUMetal4CommandAllocator) InitWithDeviceAndAliasToDevicePools(pools objectivec.IObject) IOGPUMetal4CommandAllocator {
	rv := objc.SendIfResponds[IOGPUMetal4CommandAllocator](i.ID, objc.Sel("initWithDeviceAndAliasToDevicePools:"), pools)
	return rv
}
