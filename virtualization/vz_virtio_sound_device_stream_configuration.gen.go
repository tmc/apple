// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZVirtioSoundDeviceStreamConfiguration] class.
var (
	_VZVirtioSoundDeviceStreamConfigurationClass     VZVirtioSoundDeviceStreamConfigurationClass
	_VZVirtioSoundDeviceStreamConfigurationClassOnce sync.Once
)

func getVZVirtioSoundDeviceStreamConfigurationClass() VZVirtioSoundDeviceStreamConfigurationClass {
	_VZVirtioSoundDeviceStreamConfigurationClassOnce.Do(func() {
		_VZVirtioSoundDeviceStreamConfigurationClass = VZVirtioSoundDeviceStreamConfigurationClass{class: objc.GetClass("VZVirtioSoundDeviceStreamConfiguration")}
	})
	return _VZVirtioSoundDeviceStreamConfigurationClass
}

// GetVZVirtioSoundDeviceStreamConfigurationClass returns the class object for VZVirtioSoundDeviceStreamConfiguration.
func GetVZVirtioSoundDeviceStreamConfigurationClass() VZVirtioSoundDeviceStreamConfigurationClass {
	return getVZVirtioSoundDeviceStreamConfigurationClass()
}

type VZVirtioSoundDeviceStreamConfigurationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioSoundDeviceStreamConfigurationClass) Alloc() VZVirtioSoundDeviceStreamConfiguration {
	rv := objc.Send[VZVirtioSoundDeviceStreamConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// An object that defines a Virtio sound device stream configuration.
//
// # Overview
// 
// A [VZVirtioSoundDeviceStreamConfiguration] object represents a PCM stream
// of audio data. Don’t instantiate this class directly. Instead,
// instantiate one of its subclasses such as
// [VZVirtioSoundDeviceInputStreamConfiguration] or
// [VZVirtioSoundDeviceOutputStreamConfiguration].
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioSoundDeviceStreamConfiguration
type VZVirtioSoundDeviceStreamConfiguration struct {
	objectivec.Object
}

// VZVirtioSoundDeviceStreamConfigurationFromID constructs a [VZVirtioSoundDeviceStreamConfiguration] from an objc.ID.
//
// An object that defines a Virtio sound device stream configuration.
func VZVirtioSoundDeviceStreamConfigurationFromID(id objc.ID) VZVirtioSoundDeviceStreamConfiguration {
	return VZVirtioSoundDeviceStreamConfiguration{objectivec.Object{id}}
}
// NOTE: VZVirtioSoundDeviceStreamConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VZVirtioSoundDeviceStreamConfiguration] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioSoundDeviceStreamConfiguration
type IVZVirtioSoundDeviceStreamConfiguration interface {
	objectivec.IObject
}





// Init initializes the instance.
func (v VZVirtioSoundDeviceStreamConfiguration) Init() VZVirtioSoundDeviceStreamConfiguration {
	rv := objc.Send[VZVirtioSoundDeviceStreamConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioSoundDeviceStreamConfiguration) Autorelease() VZVirtioSoundDeviceStreamConfiguration {
	rv := objc.Send[VZVirtioSoundDeviceStreamConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioSoundDeviceStreamConfiguration creates a new VZVirtioSoundDeviceStreamConfiguration instance.
func NewVZVirtioSoundDeviceStreamConfiguration() VZVirtioSoundDeviceStreamConfiguration {
	class := getVZVirtioSoundDeviceStreamConfigurationClass()
	rv := objc.Send[VZVirtioSoundDeviceStreamConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}










































