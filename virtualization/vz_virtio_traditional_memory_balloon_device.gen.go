// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VZVirtioTraditionalMemoryBalloonDevice] class.
var (
	_VZVirtioTraditionalMemoryBalloonDeviceClass     VZVirtioTraditionalMemoryBalloonDeviceClass
	_VZVirtioTraditionalMemoryBalloonDeviceClassOnce sync.Once
)

func getVZVirtioTraditionalMemoryBalloonDeviceClass() VZVirtioTraditionalMemoryBalloonDeviceClass {
	_VZVirtioTraditionalMemoryBalloonDeviceClassOnce.Do(func() {
		_VZVirtioTraditionalMemoryBalloonDeviceClass = VZVirtioTraditionalMemoryBalloonDeviceClass{class: objc.GetClass("VZVirtioTraditionalMemoryBalloonDevice")}
	})
	return _VZVirtioTraditionalMemoryBalloonDeviceClass
}

// GetVZVirtioTraditionalMemoryBalloonDeviceClass returns the class object for VZVirtioTraditionalMemoryBalloonDevice.
func GetVZVirtioTraditionalMemoryBalloonDeviceClass() VZVirtioTraditionalMemoryBalloonDeviceClass {
	return getVZVirtioTraditionalMemoryBalloonDeviceClass()
}

type VZVirtioTraditionalMemoryBalloonDeviceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtioTraditionalMemoryBalloonDeviceClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioTraditionalMemoryBalloonDeviceClass) Alloc() VZVirtioTraditionalMemoryBalloonDevice {
	rv := objc.Send[VZVirtioTraditionalMemoryBalloonDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// The object you use to change the amount of memory allocated to the guest
// system.
//
// # Overview
// 
// A [VZVirtioTraditionalMemoryBalloonDevice] object implements a
// Virtio-compliant balloon memory device, which lets you change the amount of
// physical memory assigned to the guest operating system. The virtual machine
// has no insight into the amount of memory its guest operating system uses. A
// memory balloon device lets you ask the guest operating system to relinquish
// memory voluntarily, which you might do if memory resources become scarce.
// 
// You don’t create a [VZVirtioTraditionalMemoryBalloonDevice] object
// directly. Instead, create a
// [VZVirtioTraditionalMemoryBalloonDeviceConfiguration] object and assign it
// to the [VZVirtioTraditionalMemoryBalloonDevice.MemoryBalloonDevices] property of your virtual machine
// configuration. In response, the virtual machine creates this object and
// assigns it to its [VZVirtioTraditionalMemoryBalloonDevice.MemoryBalloonDevices] property.
// 
// To use a memory balloon device, change the value in the
// [VZVirtioTraditionalMemoryBalloonDevice.TargetVirtualMachineMemorySize] property when your virtual machine is
// running. If the new value is smaller than the amount of currently assigned
// memory, the guest system may return a list of unused memory pages using the
// memory balloon device. If it does, the virtual machine releases those pages
// back to the host computer. If it doesn’t return any memory pages, the
// virtual machine leaves the guest’s memory size unchanged. If the new
// value is larger than the amount of currently assigned memory, the virtual
// machine reserves more pages for the guest operating system.
// 
// For optimal performance, the guest operating system should compact its
// memory before returning any pages back to the memory balloon device.
// Compacting the memory reduces fragmentation, and allows it to return
// contiguous blocks of free pages in the memory balloon device.
//
// # Changing the memory partition size
//
//   - [VZVirtioTraditionalMemoryBalloonDevice.TargetVirtualMachineMemorySize]: The target amount of memory, in bytes, to make available to the virtual machine.
//   - [VZVirtioTraditionalMemoryBalloonDevice.SetTargetVirtualMachineMemorySize]
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioTraditionalMemoryBalloonDevice
type VZVirtioTraditionalMemoryBalloonDevice struct {
	VZMemoryBalloonDevice
}

// VZVirtioTraditionalMemoryBalloonDeviceFromID constructs a [VZVirtioTraditionalMemoryBalloonDevice] from an objc.ID.
//
// The object you use to change the amount of memory allocated to the guest
// system.
func VZVirtioTraditionalMemoryBalloonDeviceFromID(id objc.ID) VZVirtioTraditionalMemoryBalloonDevice {
	return VZVirtioTraditionalMemoryBalloonDevice{VZMemoryBalloonDevice: VZMemoryBalloonDeviceFromID(id)}
}
// NOTE: VZVirtioTraditionalMemoryBalloonDevice adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZVirtioTraditionalMemoryBalloonDevice] class.
//
// # Changing the memory partition size
//
//   - [IVZVirtioTraditionalMemoryBalloonDevice.TargetVirtualMachineMemorySize]: The target amount of memory, in bytes, to make available to the virtual machine.
//   - [IVZVirtioTraditionalMemoryBalloonDevice.SetTargetVirtualMachineMemorySize]
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioTraditionalMemoryBalloonDevice
type IVZVirtioTraditionalMemoryBalloonDevice interface {
	IVZMemoryBalloonDevice

	// Topic: Changing the memory partition size

	// The target amount of memory, in bytes, to make available to the virtual machine.
	TargetVirtualMachineMemorySize() uint64
	SetTargetVirtualMachineMemorySize(value uint64)
}

// Init initializes the instance.
func (v VZVirtioTraditionalMemoryBalloonDevice) Init() VZVirtioTraditionalMemoryBalloonDevice {
	rv := objc.Send[VZVirtioTraditionalMemoryBalloonDevice](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioTraditionalMemoryBalloonDevice) Autorelease() VZVirtioTraditionalMemoryBalloonDevice {
	rv := objc.Send[VZVirtioTraditionalMemoryBalloonDevice](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioTraditionalMemoryBalloonDevice creates a new VZVirtioTraditionalMemoryBalloonDevice instance.
func NewVZVirtioTraditionalMemoryBalloonDevice() VZVirtioTraditionalMemoryBalloonDevice {
	class := getVZVirtioTraditionalMemoryBalloonDeviceClass()
	rv := objc.Send[VZVirtioTraditionalMemoryBalloonDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The target amount of memory, in bytes, to make available to the virtual
// machine.
//
// # Discussion
// 
// Use this property to change the amount of physical memory currently
// assigned to the guest operating system. If the new value is less than the
// current amount of physical memory, the virtual machine asks the guest
// operating system to return enough memory pages to fit within the new limit.
// If the guest complies, the virtual machine updates the guest’s physical
// memory size to the new value. If the guest doesn’t comply, the physical
// memory size remains unchanged.
// 
// The new value you specify must be a multiple of 1 megabyte — that is,
// 1024 * 1024 bytes. The value must also be less than or equal to the value
// in the [MemorySize] property of your [VZVirtualMachineConfiguration]
// object. If you specify a value that isn’t rounded to the nearest
// megabyte, the virtual machine rounds the memory size down to the nearest
// megabyte. If the resulting value is less than the value in
// [MinimumAllowedMemorySize], the virtual machine rounds up to the minimum
// value.
// 
// The virtual machine sets the initial value of this property to the value in
// the [MemorySize] property of your configuration object.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioTraditionalMemoryBalloonDevice/targetVirtualMachineMemorySize
func (v VZVirtioTraditionalMemoryBalloonDevice) TargetVirtualMachineMemorySize() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("targetVirtualMachineMemorySize"))
	return rv
}
func (v VZVirtioTraditionalMemoryBalloonDevice) SetTargetVirtualMachineMemorySize(value uint64) {
	objc.Send[struct{}](v.ID, objc.Sel("setTargetVirtualMachineMemorySize:"), value)
}

