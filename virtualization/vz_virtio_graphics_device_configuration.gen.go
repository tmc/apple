// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZVirtioGraphicsDeviceConfiguration] class.
var (
	_VZVirtioGraphicsDeviceConfigurationClass     VZVirtioGraphicsDeviceConfigurationClass
	_VZVirtioGraphicsDeviceConfigurationClassOnce sync.Once
)

func getVZVirtioGraphicsDeviceConfigurationClass() VZVirtioGraphicsDeviceConfigurationClass {
	_VZVirtioGraphicsDeviceConfigurationClassOnce.Do(func() {
		_VZVirtioGraphicsDeviceConfigurationClass = VZVirtioGraphicsDeviceConfigurationClass{class: objc.GetClass("VZVirtioGraphicsDeviceConfiguration")}
	})
	return _VZVirtioGraphicsDeviceConfigurationClass
}

// GetVZVirtioGraphicsDeviceConfigurationClass returns the class object for VZVirtioGraphicsDeviceConfiguration.
func GetVZVirtioGraphicsDeviceConfigurationClass() VZVirtioGraphicsDeviceConfigurationClass {
	return getVZVirtioGraphicsDeviceConfigurationClass()
}

type VZVirtioGraphicsDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtioGraphicsDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtioGraphicsDeviceConfigurationClass) Alloc() VZVirtioGraphicsDeviceConfiguration {
	rv := objc.Send[VZVirtioGraphicsDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// Configuration that represents the configuration of a Virtio graphics device
// for a Linux VM.
//
// # Instance properties
//
//   - [VZVirtioGraphicsDeviceConfiguration.Scanouts]: The array of output devices.
//   - [VZVirtioGraphicsDeviceConfiguration.SetScanouts]
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioGraphicsDeviceConfiguration
type VZVirtioGraphicsDeviceConfiguration struct {
	VZGraphicsDeviceConfiguration
}

// VZVirtioGraphicsDeviceConfigurationFromID constructs a [VZVirtioGraphicsDeviceConfiguration] from an objc.ID.
//
// Configuration that represents the configuration of a Virtio graphics device
// for a Linux VM.
func VZVirtioGraphicsDeviceConfigurationFromID(id objc.ID) VZVirtioGraphicsDeviceConfiguration {
	return VZVirtioGraphicsDeviceConfiguration{VZGraphicsDeviceConfiguration: VZGraphicsDeviceConfigurationFromID(id)}
}
// NOTE: VZVirtioGraphicsDeviceConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZVirtioGraphicsDeviceConfiguration] class.
//
// # Instance properties
//
//   - [IVZVirtioGraphicsDeviceConfiguration.Scanouts]: The array of output devices.
//   - [IVZVirtioGraphicsDeviceConfiguration.SetScanouts]
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioGraphicsDeviceConfiguration
type IVZVirtioGraphicsDeviceConfiguration interface {
	IVZGraphicsDeviceConfiguration

	// Topic: Instance properties

	// The array of output devices.
	Scanouts() []VZVirtioGraphicsScanoutConfiguration
	SetScanouts(value []VZVirtioGraphicsScanoutConfiguration)
}

// Init initializes the instance.
func (v VZVirtioGraphicsDeviceConfiguration) Init() VZVirtioGraphicsDeviceConfiguration {
	rv := objc.Send[VZVirtioGraphicsDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtioGraphicsDeviceConfiguration) Autorelease() VZVirtioGraphicsDeviceConfiguration {
	rv := objc.Send[VZVirtioGraphicsDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtioGraphicsDeviceConfiguration creates a new VZVirtioGraphicsDeviceConfiguration instance.
func NewVZVirtioGraphicsDeviceConfiguration() VZVirtioGraphicsDeviceConfiguration {
	class := getVZVirtioGraphicsDeviceConfigurationClass()
	rv := objc.Send[VZVirtioGraphicsDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The array of output devices.
//
// See: https://developer.apple.com/documentation/Virtualization/VZVirtioGraphicsDeviceConfiguration/scanouts
func (v VZVirtioGraphicsDeviceConfiguration) Scanouts() []VZVirtioGraphicsScanoutConfiguration {
	rv := objc.Send[[]objc.ID](v.ID, objc.Sel("scanouts"))
	return objc.ConvertSlice(rv, func(id objc.ID) VZVirtioGraphicsScanoutConfiguration {
		return VZVirtioGraphicsScanoutConfigurationFromID(id)
	})
}
func (v VZVirtioGraphicsDeviceConfiguration) SetScanouts(value []VZVirtioGraphicsScanoutConfiguration) {
	objc.Send[struct{}](v.ID, objc.Sel("setScanouts:"), objectivec.IObjectSliceToNSArray(value))
}

