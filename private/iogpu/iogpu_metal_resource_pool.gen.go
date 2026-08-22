// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOGPUMetalResourcePool] class.
var (
	_IOGPUMetalResourcePoolClass     IOGPUMetalResourcePoolClass
	_IOGPUMetalResourcePoolClassOnce sync.Once
)

func getIOGPUMetalResourcePoolClass() IOGPUMetalResourcePoolClass {
	_IOGPUMetalResourcePoolClassOnce.Do(func() {
		_IOGPUMetalResourcePoolClass = IOGPUMetalResourcePoolClass{class: objc.GetClass("IOGPUMetalResourcePool")}
	})
	return _IOGPUMetalResourcePoolClass
}

// GetIOGPUMetalResourcePoolClass returns the class object for IOGPUMetalResourcePool.
func GetIOGPUMetalResourcePoolClass() IOGPUMetalResourcePoolClass {
	return getIOGPUMetalResourcePoolClass()
}

type IOGPUMetalResourcePoolClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOGPUMetalResourcePoolClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOGPUMetalResourcePoolClass) Alloc() IOGPUMetalResourcePool {
	rv := objc.SendIfResponds[IOGPUMetalResourcePool](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [IOGPUMetalResourcePool.AllocatedSize]
//   - [IOGPUMetalResourcePool.AvailableCount]
//   - [IOGPUMetalResourcePool.Purge]
//   - [IOGPUMetalResourcePool.PurgeWithLock]
//   - [IOGPUMetalResourcePool.ResourceArgs]
//   - [IOGPUMetalResourcePool.ResourceArgsSize]
//   - [IOGPUMetalResourcePool.SetResourceArgsResourceArgsSize]
//   - [IOGPUMetalResourcePool.UpdateResourcePurgeability]
//   - [IOGPUMetalResourcePool.InitWithDeviceResourceClassResourceArgsResourceArgsSizeOptions]
type IOGPUMetalResourcePool struct {
	objectivec.Object
}

// IOGPUMetalResourcePoolFromID constructs a [IOGPUMetalResourcePool] from an objc.ID.
func IOGPUMetalResourcePoolFromID(id objc.ID) IOGPUMetalResourcePool {
	return IOGPUMetalResourcePool{objectivec.Object{ID: id}}
}

// Ensure IOGPUMetalResourcePool implements IIOGPUMetalResourcePool.
var _ IIOGPUMetalResourcePool = IOGPUMetalResourcePool{}

// An interface definition for the [IOGPUMetalResourcePool] class.
//
// # Methods
//
//   - [IIOGPUMetalResourcePool.AllocatedSize]
//   - [IIOGPUMetalResourcePool.AvailableCount]
//   - [IIOGPUMetalResourcePool.Purge]
//   - [IIOGPUMetalResourcePool.PurgeWithLock]
//   - [IIOGPUMetalResourcePool.ResourceArgs]
//   - [IIOGPUMetalResourcePool.ResourceArgsSize]
//   - [IIOGPUMetalResourcePool.SetResourceArgsResourceArgsSize]
//   - [IIOGPUMetalResourcePool.UpdateResourcePurgeability]
//   - [IIOGPUMetalResourcePool.InitWithDeviceResourceClassResourceArgsResourceArgsSizeOptions]
type IIOGPUMetalResourcePool interface {
	objectivec.IObject

	// Topic: Methods

	AllocatedSize() uint64
	AvailableCount() int
	Purge()
	PurgeWithLock()
	ResourceArgs() *IOGPUNewResourceArgs
	ResourceArgsSize() uint32
	SetResourceArgsResourceArgsSize(args *IOGPUNewResourceArgs, size uint32)
	UpdateResourcePurgeability() bool
	InitWithDeviceResourceClassResourceArgsResourceArgsSizeOptions(device objectivec.IObject, class objectivec.Class, args *IOGPUNewResourceArgs, size uint32, options objectivec.IObject) IOGPUMetalResourcePool
}

// Init initializes the instance.
func (i IOGPUMetalResourcePool) Init() IOGPUMetalResourcePool {
	rv := objc.SendIfResponds[IOGPUMetalResourcePool](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IOGPUMetalResourcePool) Autorelease() IOGPUMetalResourcePool {
	rv := objc.SendIfResponds[IOGPUMetalResourcePool](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOGPUMetalResourcePool creates a new IOGPUMetalResourcePool instance.
func NewIOGPUMetalResourcePool() IOGPUMetalResourcePool {
	class := getIOGPUMetalResourcePoolClass()
	rv := objc.SendIfResponds[IOGPUMetalResourcePool](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGPUMetalResourcePoolWithDeviceResourceClassResourceArgsResourceArgsSizeOptions(device objectivec.IObject, class objectivec.Class, args *IOGPUNewResourceArgs, size uint32, options objectivec.IObject) IOGPUMetalResourcePool {
	instance := getIOGPUMetalResourcePoolClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:resourceClass:resourceArgs:resourceArgsSize:options:"), device, class, unsafe.Pointer(args), size, options)
	return IOGPUMetalResourcePoolFromID(rv)
}

func (i IOGPUMetalResourcePool) AllocatedSize() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("allocatedSize"))
	return rv
}
func (i IOGPUMetalResourcePool) AvailableCount() int {
	rv := objc.SendIfResponds[int](i.ID, objc.Sel("availableCount"))
	return rv
}
func (i IOGPUMetalResourcePool) Purge() {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("purge"))
}
func (i IOGPUMetalResourcePool) PurgeWithLock() {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("purgeWithLock"))
}
func (i IOGPUMetalResourcePool) SetResourceArgsResourceArgsSize(args *IOGPUNewResourceArgs, size uint32) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setResourceArgs:resourceArgsSize:"), unsafe.Pointer(args), size)
}
func (i IOGPUMetalResourcePool) UpdateResourcePurgeability() bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("updateResourcePurgeability"))
	return rv
}
func (i IOGPUMetalResourcePool) InitWithDeviceResourceClassResourceArgsResourceArgsSizeOptions(device objectivec.IObject, class objectivec.Class, args *IOGPUNewResourceArgs, size uint32, options objectivec.IObject) IOGPUMetalResourcePool {
	rv := objc.SendIfResponds[IOGPUMetalResourcePool](i.ID, objc.Sel("initWithDevice:resourceClass:resourceArgs:resourceArgsSize:options:"), device, class, unsafe.Pointer(args), size, options)
	return rv
}

func (i IOGPUMetalResourcePool) ResourceArgs() *IOGPUNewResourceArgs {
	rv := objc.SendIfResponds[unsafe.Pointer](i.ID, objc.Sel("resourceArgs"))
	return (*IOGPUNewResourceArgs)(rv)
}
func (i IOGPUMetalResourcePool) ResourceArgsSize() uint32 {
	rv := objc.SendIfResponds[uint32](i.ID, objc.Sel("resourceArgsSize"))
	return rv
}
