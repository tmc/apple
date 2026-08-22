// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZVirtioSoundDeviceInputStreamConfiguration] class.
var (
	_VZVirtioSoundDeviceInputStreamConfigurationClass     VZVirtioSoundDeviceInputStreamConfigurationClass
	_VZVirtioSoundDeviceInputStreamConfigurationClassOnce sync.Once
)

func getVZVirtioSoundDeviceInputStreamConfigurationClass() VZVirtioSoundDeviceInputStreamConfigurationClass {
	_VZVirtioSoundDeviceInputStreamConfigurationClassOnce.Do(func() {
		_VZVirtioSoundDeviceInputStreamConfigurationClass = VZVirtioSoundDeviceInputStreamConfigurationClass{class: objc.GetClass("VZVirtioSoundDeviceInputStreamConfiguration")}
	})
	return _VZVirtioSoundDeviceInputStreamConfigurationClass
}

// GetVZVirtioSoundDeviceInputStreamConfigurationClass returns the class object for VZVirtioSoundDeviceInputStreamConfiguration.
func GetVZVirtioSoundDeviceInputStreamConfigurationClass() VZVirtioSoundDeviceInputStreamConfigurationClass {
	return getVZVirtioSoundDeviceInputStreamConfigurationClass()
}

type VZVirtioSoundDeviceInputStreamConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtioSoundDeviceInputStreamConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioSoundDeviceInputStreamConfigurationClass) Alloc() VZVirtioSoundDeviceInputStreamConfiguration {
	rv := objc.SendIfResponds[VZVirtioSoundDeviceInputStreamConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

type VZVirtioSoundDeviceInputStreamConfiguration struct {
	VZVirtioSoundDeviceStreamConfiguration
}

// VZVirtioSoundDeviceInputStreamConfigurationFromID constructs a [VZVirtioSoundDeviceInputStreamConfiguration] from an objc.ID.
func VZVirtioSoundDeviceInputStreamConfigurationFromID(id objc.ID) VZVirtioSoundDeviceInputStreamConfiguration {
	return VZVirtioSoundDeviceInputStreamConfiguration{VZVirtioSoundDeviceStreamConfiguration: VZVirtioSoundDeviceStreamConfigurationFromID(id)}
}

// Ensure VZVirtioSoundDeviceInputStreamConfiguration implements IVZVirtioSoundDeviceInputStreamConfiguration.
var _ IVZVirtioSoundDeviceInputStreamConfiguration = VZVirtioSoundDeviceInputStreamConfiguration{}

// An interface definition for the [VZVirtioSoundDeviceInputStreamConfiguration] class.
type IVZVirtioSoundDeviceInputStreamConfiguration interface {
	IVZVirtioSoundDeviceStreamConfiguration
}

// Init initializes the instance.
func (v VZVirtioSoundDeviceInputStreamConfiguration) Init() VZVirtioSoundDeviceInputStreamConfiguration {
	rv := objc.SendIfResponds[VZVirtioSoundDeviceInputStreamConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioSoundDeviceInputStreamConfiguration) Autorelease() VZVirtioSoundDeviceInputStreamConfiguration {
	rv := objc.SendIfResponds[VZVirtioSoundDeviceInputStreamConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioSoundDeviceInputStreamConfiguration creates a new VZVirtioSoundDeviceInputStreamConfiguration instance.
func NewVZVirtioSoundDeviceInputStreamConfiguration() VZVirtioSoundDeviceInputStreamConfiguration {
	class := getVZVirtioSoundDeviceInputStreamConfigurationClass()
	rv := objc.SendIfResponds[VZVirtioSoundDeviceInputStreamConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}
