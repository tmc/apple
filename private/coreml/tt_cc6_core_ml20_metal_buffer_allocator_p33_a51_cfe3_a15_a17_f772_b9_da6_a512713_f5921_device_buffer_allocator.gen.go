// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [DeviceBufferAllocator] class.
var (
	_DeviceBufferAllocatorClass     DeviceBufferAllocatorClass
	_DeviceBufferAllocatorClassOnce sync.Once
)

func getDeviceBufferAllocatorClass() DeviceBufferAllocatorClass {
	_DeviceBufferAllocatorClassOnce.Do(func() {
		_DeviceBufferAllocatorClass = DeviceBufferAllocatorClass{class: objc.GetClass("_TtCC6CoreML20MetalBufferAllocatorP33_A51CFE3A15A17F772B9DA6A512713F5921DeviceBufferAllocator")}
	})
	return _DeviceBufferAllocatorClass
}

// GetDeviceBufferAllocatorClass returns the class object for _TtCC6CoreML20MetalBufferAllocatorP33_A51CFE3A15A17F772B9DA6A512713F5921DeviceBufferAllocator.
func GetDeviceBufferAllocatorClass() DeviceBufferAllocatorClass {
	return getDeviceBufferAllocatorClass()
}

type DeviceBufferAllocatorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (dc DeviceBufferAllocatorClass) Class() objc.Class {
	return dc.class
}

// Alloc allocates memory for a new instance of the class.
func (dc DeviceBufferAllocatorClass) Alloc() DeviceBufferAllocator {
	rv := objc.SendIfResponds[DeviceBufferAllocator](objc.ID(dc.class), objc.Sel("alloc"))
	return rv
}

type DeviceBufferAllocator struct {
	objectivec.Object
}

// DeviceBufferAllocatorFromID constructs a [DeviceBufferAllocator] from an objc.ID.
func DeviceBufferAllocatorFromID(id objc.ID) DeviceBufferAllocator {
	return DeviceBufferAllocator{objectivec.Object{ID: id}}
}

// Ensure DeviceBufferAllocator implements IDeviceBufferAllocator.
var _ IDeviceBufferAllocator = DeviceBufferAllocator{}

// An interface definition for the [DeviceBufferAllocator] class.
type IDeviceBufferAllocator interface {
	objectivec.IObject
}

// Init initializes the instance.
func (d DeviceBufferAllocator) Init() DeviceBufferAllocator {
	rv := objc.SendIfResponds[DeviceBufferAllocator](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d DeviceBufferAllocator) Autorelease() DeviceBufferAllocator {
	rv := objc.SendIfResponds[DeviceBufferAllocator](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewDeviceBufferAllocator creates a new DeviceBufferAllocator instance.
func NewDeviceBufferAllocator() DeviceBufferAllocator {
	class := getDeviceBufferAllocatorClass()
	rv := objc.SendIfResponds[DeviceBufferAllocator](objc.ID(class.class), objc.Sel("new"))
	return rv
}
