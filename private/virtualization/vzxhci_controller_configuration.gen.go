// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZXHCIControllerConfiguration] class.
var (
	_VZXHCIControllerConfigurationClass     VZXHCIControllerConfigurationClass
	_VZXHCIControllerConfigurationClassOnce sync.Once
)

func getVZXHCIControllerConfigurationClass() VZXHCIControllerConfigurationClass {
	_VZXHCIControllerConfigurationClassOnce.Do(func() {
		_VZXHCIControllerConfigurationClass = VZXHCIControllerConfigurationClass{class: objc.GetClass("VZXHCIControllerConfiguration")}
	})
	return _VZXHCIControllerConfigurationClass
}

// GetVZXHCIControllerConfigurationClass returns the class object for VZXHCIControllerConfiguration.
func GetVZXHCIControllerConfigurationClass() VZXHCIControllerConfigurationClass {
	return getVZXHCIControllerConfigurationClass()
}

type VZXHCIControllerConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZXHCIControllerConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZXHCIControllerConfigurationClass) Alloc() VZXHCIControllerConfiguration {
	rv := objc.Send[VZXHCIControllerConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZXHCIControllerConfiguration
type VZXHCIControllerConfiguration struct {
	VZUSBControllerConfiguration
}

// VZXHCIControllerConfigurationFromID constructs a [VZXHCIControllerConfiguration] from an objc.ID.
func VZXHCIControllerConfigurationFromID(id objc.ID) VZXHCIControllerConfiguration {
	return VZXHCIControllerConfiguration{VZUSBControllerConfiguration: VZUSBControllerConfigurationFromID(id)}
}

// Ensure VZXHCIControllerConfiguration implements IVZXHCIControllerConfiguration.
var _ IVZXHCIControllerConfiguration = VZXHCIControllerConfiguration{}

// An interface definition for the [VZXHCIControllerConfiguration] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZXHCIControllerConfiguration
type IVZXHCIControllerConfiguration interface {
	IVZUSBControllerConfiguration
}

// Init initializes the instance.
func (x VZXHCIControllerConfiguration) Init() VZXHCIControllerConfiguration {
	rv := objc.Send[VZXHCIControllerConfiguration](x.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (x VZXHCIControllerConfiguration) Autorelease() VZXHCIControllerConfiguration {
	rv := objc.Send[VZXHCIControllerConfiguration](x.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZXHCIControllerConfiguration creates a new VZXHCIControllerConfiguration instance.
func NewVZXHCIControllerConfiguration() VZXHCIControllerConfiguration {
	class := getVZXHCIControllerConfigurationClass()
	rv := objc.Send[VZXHCIControllerConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}
