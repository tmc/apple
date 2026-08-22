// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOGPUMetalDeviceShmem] class.
var (
	_IOGPUMetalDeviceShmemClass     IOGPUMetalDeviceShmemClass
	_IOGPUMetalDeviceShmemClassOnce sync.Once
)

func getIOGPUMetalDeviceShmemClass() IOGPUMetalDeviceShmemClass {
	_IOGPUMetalDeviceShmemClassOnce.Do(func() {
		_IOGPUMetalDeviceShmemClass = IOGPUMetalDeviceShmemClass{class: objc.GetClass("IOGPUMetalDeviceShmem")}
	})
	return _IOGPUMetalDeviceShmemClass
}

// GetIOGPUMetalDeviceShmemClass returns the class object for IOGPUMetalDeviceShmem.
func GetIOGPUMetalDeviceShmemClass() IOGPUMetalDeviceShmemClass {
	return getIOGPUMetalDeviceShmemClass()
}

type IOGPUMetalDeviceShmemClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOGPUMetalDeviceShmemClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOGPUMetalDeviceShmemClass) Alloc() IOGPUMetalDeviceShmem {
	rv := objc.SendIfResponds[IOGPUMetalDeviceShmem](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [IOGPUMetalDeviceShmem.ShmemID]
//   - [IOGPUMetalDeviceShmem.ShmemSize]
//   - [IOGPUMetalDeviceShmem.VirtualAddress]
//   - [IOGPUMetalDeviceShmem.InitWithDeviceShmemSizeShmemType]
type IOGPUMetalDeviceShmem struct {
	objectivec.Object
}

// IOGPUMetalDeviceShmemFromID constructs a [IOGPUMetalDeviceShmem] from an objc.ID.
func IOGPUMetalDeviceShmemFromID(id objc.ID) IOGPUMetalDeviceShmem {
	return IOGPUMetalDeviceShmem{objectivec.Object{ID: id}}
}

// Ensure IOGPUMetalDeviceShmem implements IIOGPUMetalDeviceShmem.
var _ IIOGPUMetalDeviceShmem = IOGPUMetalDeviceShmem{}

// An interface definition for the [IOGPUMetalDeviceShmem] class.
//
// # Methods
//
//   - [IIOGPUMetalDeviceShmem.ShmemID]
//   - [IIOGPUMetalDeviceShmem.ShmemSize]
//   - [IIOGPUMetalDeviceShmem.VirtualAddress]
//   - [IIOGPUMetalDeviceShmem.InitWithDeviceShmemSizeShmemType]
type IIOGPUMetalDeviceShmem interface {
	objectivec.IObject

	// Topic: Methods

	ShmemID() uint32
	ShmemSize() uint32
	VirtualAddress() unsafe.Pointer
	InitWithDeviceShmemSizeShmemType(device objectivec.IObject, size uint32, type_ int) IOGPUMetalDeviceShmem
}

// Init initializes the instance.
func (i IOGPUMetalDeviceShmem) Init() IOGPUMetalDeviceShmem {
	rv := objc.SendIfResponds[IOGPUMetalDeviceShmem](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IOGPUMetalDeviceShmem) Autorelease() IOGPUMetalDeviceShmem {
	rv := objc.SendIfResponds[IOGPUMetalDeviceShmem](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOGPUMetalDeviceShmem creates a new IOGPUMetalDeviceShmem instance.
func NewIOGPUMetalDeviceShmem() IOGPUMetalDeviceShmem {
	class := getIOGPUMetalDeviceShmemClass()
	rv := objc.SendIfResponds[IOGPUMetalDeviceShmem](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGPUMetalDeviceShmemWithDeviceShmemSizeShmemType(device objectivec.IObject, size uint32, type_ int) IOGPUMetalDeviceShmem {
	instance := getIOGPUMetalDeviceShmemClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:shmemSize:shmemType:"), device, size, type_)
	return IOGPUMetalDeviceShmemFromID(rv)
}

func (i IOGPUMetalDeviceShmem) InitWithDeviceShmemSizeShmemType(device objectivec.IObject, size uint32, type_ int) IOGPUMetalDeviceShmem {
	rv := objc.SendIfResponds[IOGPUMetalDeviceShmem](i.ID, objc.Sel("initWithDevice:shmemSize:shmemType:"), device, size, type_)
	return rv
}

func (i IOGPUMetalDeviceShmem) ShmemID() uint32 {
	rv := objc.SendIfResponds[uint32](i.ID, objc.Sel("shmemID"))
	return rv
}
func (i IOGPUMetalDeviceShmem) ShmemSize() uint32 {
	rv := objc.SendIfResponds[uint32](i.ID, objc.Sel("shmemSize"))
	return rv
}
func (i IOGPUMetalDeviceShmem) VirtualAddress() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](i.ID, objc.Sel("virtualAddress"))
	return rv
}
