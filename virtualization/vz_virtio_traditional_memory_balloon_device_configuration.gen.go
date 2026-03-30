// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZVirtioTraditionalMemoryBalloonDeviceConfiguration] class.
var (
	_VZVirtioTraditionalMemoryBalloonDeviceConfigurationClass     VZVirtioTraditionalMemoryBalloonDeviceConfigurationClass
	_VZVirtioTraditionalMemoryBalloonDeviceConfigurationClassOnce sync.Once
)

func getVZVirtioTraditionalMemoryBalloonDeviceConfigurationClass() VZVirtioTraditionalMemoryBalloonDeviceConfigurationClass {
	_VZVirtioTraditionalMemoryBalloonDeviceConfigurationClassOnce.Do(func() {
		_VZVirtioTraditionalMemoryBalloonDeviceConfigurationClass = VZVirtioTraditionalMemoryBalloonDeviceConfigurationClass{class: objc.GetClass("VZVirtioTraditionalMemoryBalloonDeviceConfiguration")}
	})
	return _VZVirtioTraditionalMemoryBalloonDeviceConfigurationClass
}

// GetVZVirtioTraditionalMemoryBalloonDeviceConfigurationClass returns the class object for VZVirtioTraditionalMemoryBalloonDeviceConfiguration.
func GetVZVirtioTraditionalMemoryBalloonDeviceConfigurationClass() VZVirtioTraditionalMemoryBalloonDeviceConfigurationClass {
	return getVZVirtioTraditionalMemoryBalloonDeviceConfigurationClass()
}

type VZVirtioTraditionalMemoryBalloonDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtioTraditionalMemoryBalloonDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioTraditionalMemoryBalloonDeviceConfigurationClass) Alloc() VZVirtioTraditionalMemoryBalloonDeviceConfiguration {
	rv := objc.Send[VZVirtioTraditionalMemoryBalloonDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// A configuration object that provides a way to reclaim memory from the guest
// system.
//
// # Overview
//
// Create a [VZVirtioTraditionalMemoryBalloonDeviceConfiguration] object when
// you want the ability to reclaim memory from the guest operating system.
// After creating this object, add it to the [VZVirtioTraditionalMemoryBalloonDeviceConfiguration.MemoryBalloonDevices] property
// of your [VZVirtualMachineConfiguration] object. In response, the virtual
// machine provides a [VZVirtioTraditionalMemoryBalloonDevice] object, which
// you use to initiate memory-related requests with the guest system. Access
// that object from the [VZVirtioTraditionalMemoryBalloonDeviceConfiguration.MemoryBalloonDevices] property of [VZVirtualMachine].
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioTraditionalMemoryBalloonDeviceConfiguration
type VZVirtioTraditionalMemoryBalloonDeviceConfiguration struct {
	VZMemoryBalloonDeviceConfiguration
}

// VZVirtioTraditionalMemoryBalloonDeviceConfigurationFromID constructs a [VZVirtioTraditionalMemoryBalloonDeviceConfiguration] from an objc.ID.
//
// A configuration object that provides a way to reclaim memory from the guest
// system.
func VZVirtioTraditionalMemoryBalloonDeviceConfigurationFromID(id objc.ID) VZVirtioTraditionalMemoryBalloonDeviceConfiguration {
	return VZVirtioTraditionalMemoryBalloonDeviceConfiguration{VZMemoryBalloonDeviceConfiguration: VZMemoryBalloonDeviceConfigurationFromID(id)}
}

// NOTE: VZVirtioTraditionalMemoryBalloonDeviceConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZVirtioTraditionalMemoryBalloonDeviceConfiguration] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioTraditionalMemoryBalloonDeviceConfiguration
type IVZVirtioTraditionalMemoryBalloonDeviceConfiguration interface {
	IVZMemoryBalloonDeviceConfiguration

	// The array of devices that you use to adjust the amount of memory available to the guest system.
	MemoryBalloonDevices() IVZMemoryBalloonDevice
	SetMemoryBalloonDevices(value IVZMemoryBalloonDevice)
}

// Init initializes the instance.
func (v VZVirtioTraditionalMemoryBalloonDeviceConfiguration) Init() VZVirtioTraditionalMemoryBalloonDeviceConfiguration {
	rv := objc.Send[VZVirtioTraditionalMemoryBalloonDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioTraditionalMemoryBalloonDeviceConfiguration) Autorelease() VZVirtioTraditionalMemoryBalloonDeviceConfiguration {
	rv := objc.Send[VZVirtioTraditionalMemoryBalloonDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioTraditionalMemoryBalloonDeviceConfiguration creates a new VZVirtioTraditionalMemoryBalloonDeviceConfiguration instance.
func NewVZVirtioTraditionalMemoryBalloonDeviceConfiguration() VZVirtioTraditionalMemoryBalloonDeviceConfiguration {
	class := getVZVirtioTraditionalMemoryBalloonDeviceConfigurationClass()
	rv := objc.Send[VZVirtioTraditionalMemoryBalloonDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The array of devices that you use to adjust the amount of memory available
// to the guest system.
//
// See: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/memoryballoondevices
func (v VZVirtioTraditionalMemoryBalloonDeviceConfiguration) MemoryBalloonDevices() IVZMemoryBalloonDevice {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("memoryBalloonDevices"))
	return VZMemoryBalloonDeviceFromID(objc.ID(rv))
}
func (v VZVirtioTraditionalMemoryBalloonDeviceConfiguration) SetMemoryBalloonDevices(value IVZMemoryBalloonDevice) {
	objc.Send[struct{}](v.ID, objc.Sel("setMemoryBalloonDevices:"), value)
}
