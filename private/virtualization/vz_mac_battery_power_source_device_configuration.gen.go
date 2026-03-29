// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VZMacBatteryPowerSourceDeviceConfiguration] class.
var (
	_VZMacBatteryPowerSourceDeviceConfigurationClass     VZMacBatteryPowerSourceDeviceConfigurationClass
	_VZMacBatteryPowerSourceDeviceConfigurationClassOnce sync.Once
)

func getVZMacBatteryPowerSourceDeviceConfigurationClass() VZMacBatteryPowerSourceDeviceConfigurationClass {
	_VZMacBatteryPowerSourceDeviceConfigurationClassOnce.Do(func() {
		_VZMacBatteryPowerSourceDeviceConfigurationClass = VZMacBatteryPowerSourceDeviceConfigurationClass{class: objc.GetClass("_VZMacBatteryPowerSourceDeviceConfiguration")}
	})
	return _VZMacBatteryPowerSourceDeviceConfigurationClass
}

// GetVZMacBatteryPowerSourceDeviceConfigurationClass returns the class object for _VZMacBatteryPowerSourceDeviceConfiguration.
func GetVZMacBatteryPowerSourceDeviceConfigurationClass() VZMacBatteryPowerSourceDeviceConfigurationClass {
	return getVZMacBatteryPowerSourceDeviceConfigurationClass()
}

type VZMacBatteryPowerSourceDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZMacBatteryPowerSourceDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMacBatteryPowerSourceDeviceConfigurationClass) Alloc() VZMacBatteryPowerSourceDeviceConfiguration {
	rv := objc.Send[VZMacBatteryPowerSourceDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZMacBatteryPowerSourceDeviceConfiguration.Source]
//   - [VZMacBatteryPowerSourceDeviceConfiguration.SetSource]
// See: https://developer.apple.com/documentation/Virtualization/_VZMacBatteryPowerSourceDeviceConfiguration
type VZMacBatteryPowerSourceDeviceConfiguration struct {
	VZPowerSourceDeviceConfiguration
}

// VZMacBatteryPowerSourceDeviceConfigurationFromID constructs a [VZMacBatteryPowerSourceDeviceConfiguration] from an objc.ID.
func VZMacBatteryPowerSourceDeviceConfigurationFromID(id objc.ID) VZMacBatteryPowerSourceDeviceConfiguration {
	return VZMacBatteryPowerSourceDeviceConfiguration{VZPowerSourceDeviceConfiguration: VZPowerSourceDeviceConfigurationFromID(id)}
}
// Ensure VZMacBatteryPowerSourceDeviceConfiguration implements IVZMacBatteryPowerSourceDeviceConfiguration.
var _ IVZMacBatteryPowerSourceDeviceConfiguration = VZMacBatteryPowerSourceDeviceConfiguration{}

// An interface definition for the [VZMacBatteryPowerSourceDeviceConfiguration] class.
//
// # Methods
//
//   - [IVZMacBatteryPowerSourceDeviceConfiguration.Source]
//   - [IVZMacBatteryPowerSourceDeviceConfiguration.SetSource]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZMacBatteryPowerSourceDeviceConfiguration
type IVZMacBatteryPowerSourceDeviceConfiguration interface {
	IVZPowerSourceDeviceConfiguration

	// Topic: Methods

	Source() *VZMacBatterySource
	SetSource(value *VZMacBatterySource)
}

// Init initializes the instance.
func (v VZMacBatteryPowerSourceDeviceConfiguration) Init() VZMacBatteryPowerSourceDeviceConfiguration {
	rv := objc.Send[VZMacBatteryPowerSourceDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZMacBatteryPowerSourceDeviceConfiguration) Autorelease() VZMacBatteryPowerSourceDeviceConfiguration {
	rv := objc.Send[VZMacBatteryPowerSourceDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacBatteryPowerSourceDeviceConfiguration creates a new VZMacBatteryPowerSourceDeviceConfiguration instance.
func NewVZMacBatteryPowerSourceDeviceConfiguration() VZMacBatteryPowerSourceDeviceConfiguration {
	class := getVZMacBatteryPowerSourceDeviceConfigurationClass()
	rv := objc.Send[VZMacBatteryPowerSourceDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMacBatteryPowerSourceDeviceConfiguration/source
func (v VZMacBatteryPowerSourceDeviceConfiguration) Source() *VZMacBatterySource {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("source"))
	if rv == 0 {
		return nil
	}
	val := VZMacBatterySourceFromID(objc.ID(rv))
	return &val
}
func (v VZMacBatteryPowerSourceDeviceConfiguration) SetSource(value *VZMacBatterySource) {
	if value == nil {
		objc.Send[struct{}](v.ID, objc.Sel("setSource:"), objc.ID(0))
		return
	}
	objc.Send[struct{}](v.ID, objc.Sel("setSource:"), value)
}

