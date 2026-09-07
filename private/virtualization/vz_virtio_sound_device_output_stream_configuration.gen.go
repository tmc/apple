// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZVirtioSoundDeviceOutputStreamConfiguration] class.
var (
	_VZVirtioSoundDeviceOutputStreamConfigurationClass     VZVirtioSoundDeviceOutputStreamConfigurationClass
	_VZVirtioSoundDeviceOutputStreamConfigurationClassOnce sync.Once
)

func getVZVirtioSoundDeviceOutputStreamConfigurationClass() VZVirtioSoundDeviceOutputStreamConfigurationClass {
	_VZVirtioSoundDeviceOutputStreamConfigurationClassOnce.Do(func() {
		_VZVirtioSoundDeviceOutputStreamConfigurationClass = VZVirtioSoundDeviceOutputStreamConfigurationClass{class: objc.GetClass("VZVirtioSoundDeviceOutputStreamConfiguration")}
	})
	return _VZVirtioSoundDeviceOutputStreamConfigurationClass
}

// GetVZVirtioSoundDeviceOutputStreamConfigurationClass returns the class object for VZVirtioSoundDeviceOutputStreamConfiguration.
func GetVZVirtioSoundDeviceOutputStreamConfigurationClass() VZVirtioSoundDeviceOutputStreamConfigurationClass {
	return getVZVirtioSoundDeviceOutputStreamConfigurationClass()
}

type VZVirtioSoundDeviceOutputStreamConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtioSoundDeviceOutputStreamConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioSoundDeviceOutputStreamConfigurationClass) Alloc() VZVirtioSoundDeviceOutputStreamConfiguration {
	rv := objc.SendIfResponds[VZVirtioSoundDeviceOutputStreamConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

type VZVirtioSoundDeviceOutputStreamConfiguration struct {
	VZVirtioSoundDeviceStreamConfiguration
}

// VZVirtioSoundDeviceOutputStreamConfigurationFromID constructs a [VZVirtioSoundDeviceOutputStreamConfiguration] from an objc.ID.
func VZVirtioSoundDeviceOutputStreamConfigurationFromID(id objc.ID) VZVirtioSoundDeviceOutputStreamConfiguration {
	return VZVirtioSoundDeviceOutputStreamConfiguration{VZVirtioSoundDeviceStreamConfiguration: VZVirtioSoundDeviceStreamConfigurationFromID(id)}
}

// Ensure VZVirtioSoundDeviceOutputStreamConfiguration implements IVZVirtioSoundDeviceOutputStreamConfiguration.
var _ IVZVirtioSoundDeviceOutputStreamConfiguration = VZVirtioSoundDeviceOutputStreamConfiguration{}

// An interface definition for the [VZVirtioSoundDeviceOutputStreamConfiguration] class.
type IVZVirtioSoundDeviceOutputStreamConfiguration interface {
	IVZVirtioSoundDeviceStreamConfiguration
}

// Init initializes the instance.
func (v VZVirtioSoundDeviceOutputStreamConfiguration) Init() VZVirtioSoundDeviceOutputStreamConfiguration {
	rv := objc.SendIfResponds[VZVirtioSoundDeviceOutputStreamConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioSoundDeviceOutputStreamConfiguration) Autorelease() VZVirtioSoundDeviceOutputStreamConfiguration {
	rv := objc.SendIfResponds[VZVirtioSoundDeviceOutputStreamConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioSoundDeviceOutputStreamConfiguration creates a new VZVirtioSoundDeviceOutputStreamConfiguration instance.
func NewVZVirtioSoundDeviceOutputStreamConfiguration() VZVirtioSoundDeviceOutputStreamConfiguration {
	class := getVZVirtioSoundDeviceOutputStreamConfigurationClass()
	rv := objc.SendIfResponds[VZVirtioSoundDeviceOutputStreamConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}
