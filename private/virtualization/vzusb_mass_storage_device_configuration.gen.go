// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZUSBMassStorageDeviceConfiguration] class.
var (
	_VZUSBMassStorageDeviceConfigurationClass     VZUSBMassStorageDeviceConfigurationClass
	_VZUSBMassStorageDeviceConfigurationClassOnce sync.Once
)

func getVZUSBMassStorageDeviceConfigurationClass() VZUSBMassStorageDeviceConfigurationClass {
	_VZUSBMassStorageDeviceConfigurationClassOnce.Do(func() {
		_VZUSBMassStorageDeviceConfigurationClass = VZUSBMassStorageDeviceConfigurationClass{class: objc.GetClass("VZUSBMassStorageDeviceConfiguration")}
	})
	return _VZUSBMassStorageDeviceConfigurationClass
}

// GetVZUSBMassStorageDeviceConfigurationClass returns the class object for VZUSBMassStorageDeviceConfiguration.
func GetVZUSBMassStorageDeviceConfigurationClass() VZUSBMassStorageDeviceConfigurationClass {
	return getVZUSBMassStorageDeviceConfigurationClass()
}

type VZUSBMassStorageDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZUSBMassStorageDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZUSBMassStorageDeviceConfigurationClass) Alloc() VZUSBMassStorageDeviceConfiguration {
	rv := objc.Send[VZUSBMassStorageDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZUSBMassStorageDeviceConfiguration.IsDuplicateConfiguration]
//   - [VZUSBMassStorageDeviceConfiguration.MakeUSBDeviceWithVirtualMachine]
type VZUSBMassStorageDeviceConfiguration struct {
	VZStorageDeviceConfiguration
}

// VZUSBMassStorageDeviceConfigurationFromID constructs a [VZUSBMassStorageDeviceConfiguration] from an objc.ID.
func VZUSBMassStorageDeviceConfigurationFromID(id objc.ID) VZUSBMassStorageDeviceConfiguration {
	return VZUSBMassStorageDeviceConfiguration{VZStorageDeviceConfiguration: VZStorageDeviceConfigurationFromID(id)}
}

// Ensure VZUSBMassStorageDeviceConfiguration implements IVZUSBMassStorageDeviceConfiguration.
var _ IVZUSBMassStorageDeviceConfiguration = VZUSBMassStorageDeviceConfiguration{}

// An interface definition for the [VZUSBMassStorageDeviceConfiguration] class.
//
// # Methods
//
//   - [IVZUSBMassStorageDeviceConfiguration.IsDuplicateConfiguration]
//   - [IVZUSBMassStorageDeviceConfiguration.MakeUSBDeviceWithVirtualMachine]
type IVZUSBMassStorageDeviceConfiguration interface {
	IVZStorageDeviceConfiguration

	// Topic: Methods

	IsDuplicateConfiguration(configuration objectivec.IObject) bool
	MakeUSBDeviceWithVirtualMachine(machine objectivec.IObject) objectivec.IObject
}

// Init initializes the instance.
func (v VZUSBMassStorageDeviceConfiguration) Init() VZUSBMassStorageDeviceConfiguration {
	rv := objc.Send[VZUSBMassStorageDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZUSBMassStorageDeviceConfiguration) Autorelease() VZUSBMassStorageDeviceConfiguration {
	rv := objc.Send[VZUSBMassStorageDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZUSBMassStorageDeviceConfiguration creates a new VZUSBMassStorageDeviceConfiguration instance.
func NewVZUSBMassStorageDeviceConfiguration() VZUSBMassStorageDeviceConfiguration {
	class := getVZUSBMassStorageDeviceConfigurationClass()
	rv := objc.Send[VZUSBMassStorageDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZUSBMassStorageDeviceConfiguration) IsDuplicateConfiguration(configuration objectivec.IObject) bool {
	rv := objc.Send[bool](v.ID, objc.Sel("isDuplicateConfiguration:"), configuration)
	return rv
}
func (v VZUSBMassStorageDeviceConfiguration) MakeUSBDeviceWithVirtualMachine(machine objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("makeUSBDeviceWithVirtualMachine:"), machine)
	return objectivec.Object{ID: rv}
}
