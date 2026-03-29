// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZDirectorySharingDevice] class.
var (
	_VZDirectorySharingDeviceClass     VZDirectorySharingDeviceClass
	_VZDirectorySharingDeviceClassOnce sync.Once
)

func getVZDirectorySharingDeviceClass() VZDirectorySharingDeviceClass {
	_VZDirectorySharingDeviceClassOnce.Do(func() {
		_VZDirectorySharingDeviceClass = VZDirectorySharingDeviceClass{class: objc.GetClass("VZDirectorySharingDevice")}
	})
	return _VZDirectorySharingDeviceClass
}

// GetVZDirectorySharingDeviceClass returns the class object for VZDirectorySharingDevice.
func GetVZDirectorySharingDeviceClass() VZDirectorySharingDeviceClass {
	return getVZDirectorySharingDeviceClass()
}

type VZDirectorySharingDeviceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZDirectorySharingDeviceClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZDirectorySharingDeviceClass) Alloc() VZDirectorySharingDevice {
	rv := objc.Send[VZDirectorySharingDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// The base class that represents a directory sharing device in a VM.
//
// # Overview
// 
// Don’t instantiate [VZDirectorySharingDevice] directly; configure a
// directory sharing device first by using [VZVirtualMachineConfiguration]
// through a subclass of [VZDirectorySharingDeviceConfiguration].
// 
// When you create a [VZVirtualMachine] from the configuration, the directory
// sharing devices are available through the
// `VZVirtualMachine.DirectorySharingDevices()` property.
// 
// The real type of [VZDirectorySharingDevice] corresponds to the type used by
// the configuration. For example, a [VZVirtioFileSystemDeviceConfiguration]
// leads to a device of type [VZVirtioFileSystemDevice].
//
// See: https://developer.apple.com/documentation/Virtualization/VZDirectorySharingDevice
type VZDirectorySharingDevice struct {
	objectivec.Object
}

// VZDirectorySharingDeviceFromID constructs a [VZDirectorySharingDevice] from an objc.ID.
//
// The base class that represents a directory sharing device in a VM.
func VZDirectorySharingDeviceFromID(id objc.ID) VZDirectorySharingDevice {
	return VZDirectorySharingDevice{objectivec.Object{ID: id}}
}
// NOTE: VZDirectorySharingDevice adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZDirectorySharingDevice] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZDirectorySharingDevice
type IVZDirectorySharingDevice interface {
	objectivec.IObject
}

// Init initializes the instance.
func (d VZDirectorySharingDevice) Init() VZDirectorySharingDevice {
	rv := objc.Send[VZDirectorySharingDevice](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d VZDirectorySharingDevice) Autorelease() VZDirectorySharingDevice {
	rv := objc.Send[VZDirectorySharingDevice](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZDirectorySharingDevice creates a new VZDirectorySharingDevice instance.
func NewVZDirectorySharingDevice() VZDirectorySharingDevice {
	class := getVZDirectorySharingDeviceClass()
	rv := objc.Send[VZDirectorySharingDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

