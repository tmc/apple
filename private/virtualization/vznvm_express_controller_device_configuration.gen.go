// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZNVMExpressControllerDeviceConfiguration] class.
var (
	_VZNVMExpressControllerDeviceConfigurationClass     VZNVMExpressControllerDeviceConfigurationClass
	_VZNVMExpressControllerDeviceConfigurationClassOnce sync.Once
)

func getVZNVMExpressControllerDeviceConfigurationClass() VZNVMExpressControllerDeviceConfigurationClass {
	_VZNVMExpressControllerDeviceConfigurationClassOnce.Do(func() {
		_VZNVMExpressControllerDeviceConfigurationClass = VZNVMExpressControllerDeviceConfigurationClass{class: objc.GetClass("VZNVMExpressControllerDeviceConfiguration")}
	})
	return _VZNVMExpressControllerDeviceConfigurationClass
}

// GetVZNVMExpressControllerDeviceConfigurationClass returns the class object for VZNVMExpressControllerDeviceConfiguration.
func GetVZNVMExpressControllerDeviceConfigurationClass() VZNVMExpressControllerDeviceConfigurationClass {
	return getVZNVMExpressControllerDeviceConfigurationClass()
}

type VZNVMExpressControllerDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZNVMExpressControllerDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZNVMExpressControllerDeviceConfigurationClass) Alloc() VZNVMExpressControllerDeviceConfiguration {
	rv := objc.SendIfResponds[VZNVMExpressControllerDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

type VZNVMExpressControllerDeviceConfiguration struct {
	VZStorageDeviceConfiguration
}

// VZNVMExpressControllerDeviceConfigurationFromID constructs a [VZNVMExpressControllerDeviceConfiguration] from an objc.ID.
func VZNVMExpressControllerDeviceConfigurationFromID(id objc.ID) VZNVMExpressControllerDeviceConfiguration {
	return VZNVMExpressControllerDeviceConfiguration{VZStorageDeviceConfiguration: VZStorageDeviceConfigurationFromID(id)}
}

// Ensure VZNVMExpressControllerDeviceConfiguration implements IVZNVMExpressControllerDeviceConfiguration.
var _ IVZNVMExpressControllerDeviceConfiguration = VZNVMExpressControllerDeviceConfiguration{}

// An interface definition for the [VZNVMExpressControllerDeviceConfiguration] class.
type IVZNVMExpressControllerDeviceConfiguration interface {
	IVZStorageDeviceConfiguration
}

// Init initializes the instance.
func (v VZNVMExpressControllerDeviceConfiguration) Init() VZNVMExpressControllerDeviceConfiguration {
	rv := objc.SendIfResponds[VZNVMExpressControllerDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZNVMExpressControllerDeviceConfiguration) Autorelease() VZNVMExpressControllerDeviceConfiguration {
	rv := objc.SendIfResponds[VZNVMExpressControllerDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZNVMExpressControllerDeviceConfiguration creates a new VZNVMExpressControllerDeviceConfiguration instance.
func NewVZNVMExpressControllerDeviceConfiguration() VZNVMExpressControllerDeviceConfiguration {
	class := getVZNVMExpressControllerDeviceConfigurationClass()
	rv := objc.SendIfResponds[VZNVMExpressControllerDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}
