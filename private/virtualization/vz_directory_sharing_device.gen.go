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

// # Methods
//
//   - [VZDirectorySharingDevice._initWithVirtualMachineDirectorySharingDeviceIndex]
//
// See: https://developer.apple.com/documentation/Virtualization/VZDirectorySharingDevice
type VZDirectorySharingDevice struct {
	objectivec.Object
}

// VZDirectorySharingDeviceFromID constructs a [VZDirectorySharingDevice] from an objc.ID.
func VZDirectorySharingDeviceFromID(id objc.ID) VZDirectorySharingDevice {
	return VZDirectorySharingDevice{objectivec.Object{ID: id}}
}

// Ensure VZDirectorySharingDevice implements IVZDirectorySharingDevice.
var _ IVZDirectorySharingDevice = VZDirectorySharingDevice{}

// An interface definition for the [VZDirectorySharingDevice] class.
//
// # Methods
//
//   - [IVZDirectorySharingDevice._initWithVirtualMachineDirectorySharingDeviceIndex]
//
// See: https://developer.apple.com/documentation/Virtualization/VZDirectorySharingDevice
type IVZDirectorySharingDevice interface {
	objectivec.IObject

	// Topic: Methods

	_initWithVirtualMachineDirectorySharingDeviceIndex(machine objectivec.IObject, index uint64) objectivec.IObject
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

// See: https://developer.apple.com/documentation/Virtualization/VZDirectorySharingDevice/_initWithVirtualMachine:directorySharingDeviceIndex:
func (d VZDirectorySharingDevice) _initWithVirtualMachineDirectorySharingDeviceIndex(machine objectivec.IObject, index uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("_initWithVirtualMachine:directorySharingDeviceIndex:"), machine, index)
	return objectivec.Object{ID: rv}
}

// InitWithVirtualMachineDirectorySharingDeviceIndex is an exported wrapper for the private method _initWithVirtualMachineDirectorySharingDeviceIndex.
func (d VZDirectorySharingDevice) InitWithVirtualMachineDirectorySharingDeviceIndex(machine objectivec.IObject, index uint64) objectivec.IObject {
	return d._initWithVirtualMachineDirectorySharingDeviceIndex(machine, index)
}
