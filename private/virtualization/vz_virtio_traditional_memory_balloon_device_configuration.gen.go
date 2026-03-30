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

// See: https://developer.apple.com/documentation/Virtualization/VZVirtioTraditionalMemoryBalloonDeviceConfiguration
type VZVirtioTraditionalMemoryBalloonDeviceConfiguration struct {
	VZMemoryBalloonDeviceConfiguration
}

// VZVirtioTraditionalMemoryBalloonDeviceConfigurationFromID constructs a [VZVirtioTraditionalMemoryBalloonDeviceConfiguration] from an objc.ID.
func VZVirtioTraditionalMemoryBalloonDeviceConfigurationFromID(id objc.ID) VZVirtioTraditionalMemoryBalloonDeviceConfiguration {
	return VZVirtioTraditionalMemoryBalloonDeviceConfiguration{VZMemoryBalloonDeviceConfiguration: VZMemoryBalloonDeviceConfigurationFromID(id)}
}

// Ensure VZVirtioTraditionalMemoryBalloonDeviceConfiguration implements IVZVirtioTraditionalMemoryBalloonDeviceConfiguration.
var _ IVZVirtioTraditionalMemoryBalloonDeviceConfiguration = VZVirtioTraditionalMemoryBalloonDeviceConfiguration{}

// An interface definition for the [VZVirtioTraditionalMemoryBalloonDeviceConfiguration] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioTraditionalMemoryBalloonDeviceConfiguration
type IVZVirtioTraditionalMemoryBalloonDeviceConfiguration interface {
	IVZMemoryBalloonDeviceConfiguration
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
