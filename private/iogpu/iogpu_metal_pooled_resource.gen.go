// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOGPUMetalPooledResource] class.
var (
	_IOGPUMetalPooledResourceClass     IOGPUMetalPooledResourceClass
	_IOGPUMetalPooledResourceClassOnce sync.Once
)

func getIOGPUMetalPooledResourceClass() IOGPUMetalPooledResourceClass {
	_IOGPUMetalPooledResourceClassOnce.Do(func() {
		_IOGPUMetalPooledResourceClass = IOGPUMetalPooledResourceClass{class: objc.GetClass("IOGPUMetalPooledResource")}
	})
	return _IOGPUMetalPooledResourceClass
}

// GetIOGPUMetalPooledResourceClass returns the class object for IOGPUMetalPooledResource.
func GetIOGPUMetalPooledResourceClass() IOGPUMetalPooledResourceClass {
	return getIOGPUMetalPooledResourceClass()
}

type IOGPUMetalPooledResourceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOGPUMetalPooledResourceClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOGPUMetalPooledResourceClass) Alloc() IOGPUMetalPooledResource {
	rv := objc.SendIfResponds[IOGPUMetalPooledResource](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

type IOGPUMetalPooledResource struct {
	IOGPUMetalResource
}

// IOGPUMetalPooledResourceFromID constructs a [IOGPUMetalPooledResource] from an objc.ID.
func IOGPUMetalPooledResourceFromID(id objc.ID) IOGPUMetalPooledResource {
	return IOGPUMetalPooledResource{IOGPUMetalResource: IOGPUMetalResourceFromID(id)}
}

// Ensure IOGPUMetalPooledResource implements IIOGPUMetalPooledResource.
var _ IIOGPUMetalPooledResource = IOGPUMetalPooledResource{}

// An interface definition for the [IOGPUMetalPooledResource] class.
type IIOGPUMetalPooledResource interface {
	IIOGPUMetalResource
}

// Init initializes the instance.
func (i IOGPUMetalPooledResource) Init() IOGPUMetalPooledResource {
	rv := objc.SendIfResponds[IOGPUMetalPooledResource](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IOGPUMetalPooledResource) Autorelease() IOGPUMetalPooledResource {
	rv := objc.SendIfResponds[IOGPUMetalPooledResource](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOGPUMetalPooledResource creates a new IOGPUMetalPooledResource instance.
func NewIOGPUMetalPooledResource() IOGPUMetalPooledResource {
	class := getIOGPUMetalPooledResourceClass()
	rv := objc.SendIfResponds[IOGPUMetalPooledResource](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGPUMetalPooledResourceMemorylessDescriptor(memoryless objectivec.IObject, descriptor objectivec.IObject) IOGPUMetalPooledResource {
	instance := getIOGPUMetalPooledResourceClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initMemoryless:descriptor:"), memoryless, descriptor)
	return IOGPUMetalPooledResourceFromID(rv)
}

func NewGPUMetalPooledResourceStandinWithDevice(device objectivec.IObject) IOGPUMetalPooledResource {
	instance := getIOGPUMetalPooledResourceClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initStandinWithDevice:"), device)
	return IOGPUMetalPooledResourceFromID(rv)
}

func NewGPUMetalPooledResourceWithDeviceOptionsArgsArgsSize(device objectivec.IObject, options uint64, args *IOGPUNewResourceArgs, size uint32) IOGPUMetalPooledResource {
	instance := getIOGPUMetalPooledResourceClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:options:args:argsSize:"), device, options, unsafe.Pointer(args), size)
	return IOGPUMetalPooledResourceFromID(rv)
}

func NewGPUMetalPooledResourceWithDeviceRemoteStorageResourceOptionsArgsArgsSize(device objectivec.IObject, resource objectivec.IObject, options uint64, args *IOGPUNewResourceArgs, size uint32) IOGPUMetalPooledResource {
	instance := getIOGPUMetalPooledResourceClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:remoteStorageResource:options:args:argsSize:"), device, resource, options, unsafe.Pointer(args), size)
	return IOGPUMetalPooledResourceFromID(rv)
}

func NewGPUMetalPooledResourceWithResource(resource objectivec.IObject) IOGPUMetalPooledResource {
	instance := getIOGPUMetalPooledResourceClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithResource:"), resource)
	return IOGPUMetalPooledResourceFromID(rv)
}
