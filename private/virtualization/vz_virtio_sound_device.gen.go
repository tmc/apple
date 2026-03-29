// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VZVirtioSoundDevice] class.
var (
	_VZVirtioSoundDeviceClass     VZVirtioSoundDeviceClass
	_VZVirtioSoundDeviceClassOnce sync.Once
)

func getVZVirtioSoundDeviceClass() VZVirtioSoundDeviceClass {
	_VZVirtioSoundDeviceClassOnce.Do(func() {
		_VZVirtioSoundDeviceClass = VZVirtioSoundDeviceClass{class: objc.GetClass("_VZVirtioSoundDevice")}
	})
	return _VZVirtioSoundDeviceClass
}

// GetVZVirtioSoundDeviceClass returns the class object for _VZVirtioSoundDevice.
func GetVZVirtioSoundDeviceClass() VZVirtioSoundDeviceClass {
	return getVZVirtioSoundDeviceClass()
}

type VZVirtioSoundDeviceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtioSoundDeviceClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioSoundDeviceClass) Alloc() VZVirtioSoundDevice {
	rv := objc.Send[VZVirtioSoundDevice](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZVirtioSoundDevice
type VZVirtioSoundDevice struct {
	VZAudioDevice
}

// VZVirtioSoundDeviceFromID constructs a [VZVirtioSoundDevice] from an objc.ID.
func VZVirtioSoundDeviceFromID(id objc.ID) VZVirtioSoundDevice {
	return VZVirtioSoundDevice{VZAudioDevice: VZAudioDeviceFromID(id)}
}
// Ensure VZVirtioSoundDevice implements IVZVirtioSoundDevice.
var _ IVZVirtioSoundDevice = VZVirtioSoundDevice{}

// An interface definition for the [VZVirtioSoundDevice] class.
//
// See: https://developer.apple.com/documentation/Virtualization/_VZVirtioSoundDevice
type IVZVirtioSoundDevice interface {
	IVZAudioDevice
}

// Init initializes the instance.
func (v VZVirtioSoundDevice) Init() VZVirtioSoundDevice {
	rv := objc.Send[VZVirtioSoundDevice](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioSoundDevice) Autorelease() VZVirtioSoundDevice {
	rv := objc.Send[VZVirtioSoundDevice](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioSoundDevice creates a new VZVirtioSoundDevice instance.
func NewVZVirtioSoundDevice() VZVirtioSoundDevice {
	class := getVZVirtioSoundDeviceClass()
	rv := objc.Send[VZVirtioSoundDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

