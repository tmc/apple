// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZAudioDeviceConfiguration] class.
var (
	_VZAudioDeviceConfigurationClass     VZAudioDeviceConfigurationClass
	_VZAudioDeviceConfigurationClassOnce sync.Once
)

func getVZAudioDeviceConfigurationClass() VZAudioDeviceConfigurationClass {
	_VZAudioDeviceConfigurationClassOnce.Do(func() {
		_VZAudioDeviceConfigurationClass = VZAudioDeviceConfigurationClass{class: objc.GetClass("VZAudioDeviceConfiguration")}
	})
	return _VZAudioDeviceConfigurationClass
}

// GetVZAudioDeviceConfigurationClass returns the class object for VZAudioDeviceConfiguration.
func GetVZAudioDeviceConfigurationClass() VZAudioDeviceConfigurationClass {
	return getVZAudioDeviceConfigurationClass()
}

type VZAudioDeviceConfigurationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZAudioDeviceConfigurationClass) Alloc() VZAudioDeviceConfiguration {
	rv := objc.Send[VZAudioDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// The base class for an audio device configuration.
//
// # Overview
// 
// Don’t instantiate this abstract class directly. Instead, instantiate one
// of its subclasses such as [VZVirtioSoundDeviceConfiguration].
//
// See: https://developer.apple.com/documentation/Virtualization/VZAudioDeviceConfiguration
type VZAudioDeviceConfiguration struct {
	objectivec.Object
}

// VZAudioDeviceConfigurationFromID constructs a [VZAudioDeviceConfiguration] from an objc.ID.
//
// The base class for an audio device configuration.
func VZAudioDeviceConfigurationFromID(id objc.ID) VZAudioDeviceConfiguration {
	return VZAudioDeviceConfiguration{objectivec.Object{ID: id}}
}
// NOTE: VZAudioDeviceConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZAudioDeviceConfiguration] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZAudioDeviceConfiguration
type IVZAudioDeviceConfiguration interface {
	objectivec.IObject
}

// Init initializes the instance.
func (a VZAudioDeviceConfiguration) Init() VZAudioDeviceConfiguration {
	rv := objc.Send[VZAudioDeviceConfiguration](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a VZAudioDeviceConfiguration) Autorelease() VZAudioDeviceConfiguration {
	rv := objc.Send[VZAudioDeviceConfiguration](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZAudioDeviceConfiguration creates a new VZAudioDeviceConfiguration instance.
func NewVZAudioDeviceConfiguration() VZAudioDeviceConfiguration {
	class := getVZAudioDeviceConfigurationClass()
	rv := objc.Send[VZAudioDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

