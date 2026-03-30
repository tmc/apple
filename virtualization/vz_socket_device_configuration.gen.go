// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZSocketDeviceConfiguration] class.
var (
	_VZSocketDeviceConfigurationClass     VZSocketDeviceConfigurationClass
	_VZSocketDeviceConfigurationClassOnce sync.Once
)

func getVZSocketDeviceConfigurationClass() VZSocketDeviceConfigurationClass {
	_VZSocketDeviceConfigurationClassOnce.Do(func() {
		_VZSocketDeviceConfigurationClass = VZSocketDeviceConfigurationClass{class: objc.GetClass("VZSocketDeviceConfiguration")}
	})
	return _VZSocketDeviceConfigurationClass
}

// GetVZSocketDeviceConfigurationClass returns the class object for VZSocketDeviceConfiguration.
func GetVZSocketDeviceConfigurationClass() VZSocketDeviceConfigurationClass {
	return getVZSocketDeviceConfigurationClass()
}

type VZSocketDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZSocketDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZSocketDeviceConfigurationClass) Alloc() VZSocketDeviceConfiguration {
	rv := objc.Send[VZSocketDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// The common configuration traits for socket device requests.
//
// # Overview
//
// Don’t create a [VZSocketDeviceConfiguration] object directly. Instead,
// create a [VZVirtioSocketDeviceConfiguration] object and add it to your
// virtual machine’s configuration.
//
// See: https://developer.apple.com/documentation/Virtualization/VZSocketDeviceConfiguration
type VZSocketDeviceConfiguration struct {
	objectivec.Object
}

// VZSocketDeviceConfigurationFromID constructs a [VZSocketDeviceConfiguration] from an objc.ID.
//
// The common configuration traits for socket device requests.
func VZSocketDeviceConfigurationFromID(id objc.ID) VZSocketDeviceConfiguration {
	return VZSocketDeviceConfiguration{objectivec.Object{ID: id}}
}

// NOTE: VZSocketDeviceConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZSocketDeviceConfiguration] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZSocketDeviceConfiguration
type IVZSocketDeviceConfiguration interface {
	objectivec.IObject
}

// Init initializes the instance.
func (s VZSocketDeviceConfiguration) Init() VZSocketDeviceConfiguration {
	rv := objc.Send[VZSocketDeviceConfiguration](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s VZSocketDeviceConfiguration) Autorelease() VZSocketDeviceConfiguration {
	rv := objc.Send[VZSocketDeviceConfiguration](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZSocketDeviceConfiguration creates a new VZSocketDeviceConfiguration instance.
func NewVZSocketDeviceConfiguration() VZSocketDeviceConfiguration {
	class := getVZSocketDeviceConfigurationClass()
	rv := objc.Send[VZSocketDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}
