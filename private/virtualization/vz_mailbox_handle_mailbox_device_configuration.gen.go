// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VZMailboxHandleMailboxDeviceConfiguration] class.
var (
	_VZMailboxHandleMailboxDeviceConfigurationClass     VZMailboxHandleMailboxDeviceConfigurationClass
	_VZMailboxHandleMailboxDeviceConfigurationClassOnce sync.Once
)

func getVZMailboxHandleMailboxDeviceConfigurationClass() VZMailboxHandleMailboxDeviceConfigurationClass {
	_VZMailboxHandleMailboxDeviceConfigurationClassOnce.Do(func() {
		_VZMailboxHandleMailboxDeviceConfigurationClass = VZMailboxHandleMailboxDeviceConfigurationClass{class: objc.GetClass("_VZMailboxHandleMailboxDeviceConfiguration")}
	})
	return _VZMailboxHandleMailboxDeviceConfigurationClass
}

// GetVZMailboxHandleMailboxDeviceConfigurationClass returns the class object for _VZMailboxHandleMailboxDeviceConfiguration.
func GetVZMailboxHandleMailboxDeviceConfigurationClass() VZMailboxHandleMailboxDeviceConfigurationClass {
	return getVZMailboxHandleMailboxDeviceConfigurationClass()
}

type VZMailboxHandleMailboxDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZMailboxHandleMailboxDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMailboxHandleMailboxDeviceConfigurationClass) Alloc() VZMailboxHandleMailboxDeviceConfiguration {
	rv := objc.Send[VZMailboxHandleMailboxDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMailboxHandleMailboxDeviceConfiguration
type VZMailboxHandleMailboxDeviceConfiguration struct {
	VZMailboxDeviceConfiguration
}

// VZMailboxHandleMailboxDeviceConfigurationFromID constructs a [VZMailboxHandleMailboxDeviceConfiguration] from an objc.ID.
func VZMailboxHandleMailboxDeviceConfigurationFromID(id objc.ID) VZMailboxHandleMailboxDeviceConfiguration {
	return VZMailboxHandleMailboxDeviceConfiguration{VZMailboxDeviceConfiguration: VZMailboxDeviceConfigurationFromID(id)}
}
// Ensure VZMailboxHandleMailboxDeviceConfiguration implements IVZMailboxHandleMailboxDeviceConfiguration.
var _ IVZMailboxHandleMailboxDeviceConfiguration = VZMailboxHandleMailboxDeviceConfiguration{}

// An interface definition for the [VZMailboxHandleMailboxDeviceConfiguration] class.
//
// See: https://developer.apple.com/documentation/Virtualization/_VZMailboxHandleMailboxDeviceConfiguration
type IVZMailboxHandleMailboxDeviceConfiguration interface {
	IVZMailboxDeviceConfiguration
}

// Init initializes the instance.
func (v VZMailboxHandleMailboxDeviceConfiguration) Init() VZMailboxHandleMailboxDeviceConfiguration {
	rv := objc.Send[VZMailboxHandleMailboxDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZMailboxHandleMailboxDeviceConfiguration) Autorelease() VZMailboxHandleMailboxDeviceConfiguration {
	rv := objc.Send[VZMailboxHandleMailboxDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMailboxHandleMailboxDeviceConfiguration creates a new VZMailboxHandleMailboxDeviceConfiguration instance.
func NewVZMailboxHandleMailboxDeviceConfiguration() VZMailboxHandleMailboxDeviceConfiguration {
	class := getVZMailboxHandleMailboxDeviceConfigurationClass()
	rv := objc.Send[VZMailboxHandleMailboxDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

