// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOGPUMetalDeviceShmemPool] class.
var (
	_IOGPUMetalDeviceShmemPoolClass     IOGPUMetalDeviceShmemPoolClass
	_IOGPUMetalDeviceShmemPoolClassOnce sync.Once
)

func getIOGPUMetalDeviceShmemPoolClass() IOGPUMetalDeviceShmemPoolClass {
	_IOGPUMetalDeviceShmemPoolClassOnce.Do(func() {
		_IOGPUMetalDeviceShmemPoolClass = IOGPUMetalDeviceShmemPoolClass{class: objc.GetClass("IOGPUMetalDeviceShmemPool")}
	})
	return _IOGPUMetalDeviceShmemPoolClass
}

// GetIOGPUMetalDeviceShmemPoolClass returns the class object for IOGPUMetalDeviceShmemPool.
func GetIOGPUMetalDeviceShmemPoolClass() IOGPUMetalDeviceShmemPoolClass {
	return getIOGPUMetalDeviceShmemPoolClass()
}

type IOGPUMetalDeviceShmemPoolClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOGPUMetalDeviceShmemPoolClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOGPUMetalDeviceShmemPoolClass) Alloc() IOGPUMetalDeviceShmemPool {
	rv := objc.SendIfResponds[IOGPUMetalDeviceShmemPool](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [IOGPUMetalDeviceShmemPool.AllocatedSize]
//   - [IOGPUMetalDeviceShmemPool.AvailableCount]
//   - [IOGPUMetalDeviceShmemPool.Prune]
//   - [IOGPUMetalDeviceShmemPool.Purge]
//   - [IOGPUMetalDeviceShmemPool.SetShmemSize]
//   - [IOGPUMetalDeviceShmemPool.ShmemSize]
//   - [IOGPUMetalDeviceShmemPool.InitWithDeviceResourceClassShmemSizeShmemTypeOptions]
type IOGPUMetalDeviceShmemPool struct {
	objectivec.Object
}

// IOGPUMetalDeviceShmemPoolFromID constructs a [IOGPUMetalDeviceShmemPool] from an objc.ID.
func IOGPUMetalDeviceShmemPoolFromID(id objc.ID) IOGPUMetalDeviceShmemPool {
	return IOGPUMetalDeviceShmemPool{objectivec.Object{ID: id}}
}

// Ensure IOGPUMetalDeviceShmemPool implements IIOGPUMetalDeviceShmemPool.
var _ IIOGPUMetalDeviceShmemPool = IOGPUMetalDeviceShmemPool{}

// An interface definition for the [IOGPUMetalDeviceShmemPool] class.
//
// # Methods
//
//   - [IIOGPUMetalDeviceShmemPool.AllocatedSize]
//   - [IIOGPUMetalDeviceShmemPool.AvailableCount]
//   - [IIOGPUMetalDeviceShmemPool.Prune]
//   - [IIOGPUMetalDeviceShmemPool.Purge]
//   - [IIOGPUMetalDeviceShmemPool.SetShmemSize]
//   - [IIOGPUMetalDeviceShmemPool.ShmemSize]
//   - [IIOGPUMetalDeviceShmemPool.InitWithDeviceResourceClassShmemSizeShmemTypeOptions]
type IIOGPUMetalDeviceShmemPool interface {
	objectivec.IObject

	// Topic: Methods

	AllocatedSize() uint64
	AvailableCount() int
	Prune()
	Purge()
	SetShmemSize(size uint32)
	ShmemSize() uint32
	InitWithDeviceResourceClassShmemSizeShmemTypeOptions(device objectivec.IObject, class objectivec.Class, size uint32, type_ int, options objectivec.IObject) IOGPUMetalDeviceShmemPool
}

// Init initializes the instance.
func (i IOGPUMetalDeviceShmemPool) Init() IOGPUMetalDeviceShmemPool {
	rv := objc.SendIfResponds[IOGPUMetalDeviceShmemPool](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IOGPUMetalDeviceShmemPool) Autorelease() IOGPUMetalDeviceShmemPool {
	rv := objc.SendIfResponds[IOGPUMetalDeviceShmemPool](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOGPUMetalDeviceShmemPool creates a new IOGPUMetalDeviceShmemPool instance.
func NewIOGPUMetalDeviceShmemPool() IOGPUMetalDeviceShmemPool {
	class := getIOGPUMetalDeviceShmemPoolClass()
	rv := objc.SendIfResponds[IOGPUMetalDeviceShmemPool](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGPUMetalDeviceShmemPoolWithDeviceResourceClassShmemSizeShmemTypeOptions(device objectivec.IObject, class objectivec.Class, size uint32, type_ int, options objectivec.IObject) IOGPUMetalDeviceShmemPool {
	instance := getIOGPUMetalDeviceShmemPoolClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:resourceClass:shmemSize:shmemType:options:"), device, class, size, type_, options)
	return IOGPUMetalDeviceShmemPoolFromID(rv)
}

func (i IOGPUMetalDeviceShmemPool) AllocatedSize() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("allocatedSize"))
	return rv
}
func (i IOGPUMetalDeviceShmemPool) AvailableCount() int {
	rv := objc.SendIfResponds[int](i.ID, objc.Sel("availableCount"))
	return rv
}
func (i IOGPUMetalDeviceShmemPool) Prune() {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("prune"))
}
func (i IOGPUMetalDeviceShmemPool) Purge() {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("purge"))
}
func (i IOGPUMetalDeviceShmemPool) SetShmemSize(size uint32) {
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("setShmemSize:"), size)
}
func (i IOGPUMetalDeviceShmemPool) ShmemSize() uint32 {
	rv := objc.SendIfResponds[uint32](i.ID, objc.Sel("shmemSize"))
	return rv
}
func (i IOGPUMetalDeviceShmemPool) InitWithDeviceResourceClassShmemSizeShmemTypeOptions(device objectivec.IObject, class objectivec.Class, size uint32, type_ int, options objectivec.IObject) IOGPUMetalDeviceShmemPool {
	rv := objc.SendIfResponds[IOGPUMetalDeviceShmemPool](i.ID, objc.Sel("initWithDevice:resourceClass:shmemSize:shmemType:options:"), device, class, size, type_, options)
	return rv
}
