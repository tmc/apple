// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZPowerSourceDeviceConfiguration] class.
var (
	_VZPowerSourceDeviceConfigurationClass     VZPowerSourceDeviceConfigurationClass
	_VZPowerSourceDeviceConfigurationClassOnce sync.Once
)

func getVZPowerSourceDeviceConfigurationClass() VZPowerSourceDeviceConfigurationClass {
	_VZPowerSourceDeviceConfigurationClassOnce.Do(func() {
		_VZPowerSourceDeviceConfigurationClass = VZPowerSourceDeviceConfigurationClass{class: objc.GetClass("_VZPowerSourceDeviceConfiguration")}
	})
	return _VZPowerSourceDeviceConfigurationClass
}

// GetVZPowerSourceDeviceConfigurationClass returns the class object for _VZPowerSourceDeviceConfiguration.
func GetVZPowerSourceDeviceConfigurationClass() VZPowerSourceDeviceConfigurationClass {
	return getVZPowerSourceDeviceConfigurationClass()
}

type VZPowerSourceDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZPowerSourceDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZPowerSourceDeviceConfigurationClass) Alloc() VZPowerSourceDeviceConfiguration {
	rv := objc.Send[VZPowerSourceDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZPowerSourceDeviceConfiguration._init]
//   - [VZPowerSourceDeviceConfiguration._powerSourceDevice]
//   - [VZPowerSourceDeviceConfiguration.EncodeWithEncoder]
//   - [VZPowerSourceDeviceConfiguration.MakePowerSourceDeviceForVirtualMachinePowerSourceDeviceIndex]
//   - [VZPowerSourceDeviceConfiguration.DebugDescription]
//   - [VZPowerSourceDeviceConfiguration.Description]
//   - [VZPowerSourceDeviceConfiguration.Hash]
//   - [VZPowerSourceDeviceConfiguration.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZPowerSourceDeviceConfiguration
type VZPowerSourceDeviceConfiguration struct {
	objectivec.Object
}

// VZPowerSourceDeviceConfigurationFromID constructs a [VZPowerSourceDeviceConfiguration] from an objc.ID.
func VZPowerSourceDeviceConfigurationFromID(id objc.ID) VZPowerSourceDeviceConfiguration {
	return VZPowerSourceDeviceConfiguration{objectivec.Object{ID: id}}
}

// Ensure VZPowerSourceDeviceConfiguration implements IVZPowerSourceDeviceConfiguration.
var _ IVZPowerSourceDeviceConfiguration = VZPowerSourceDeviceConfiguration{}

// An interface definition for the [VZPowerSourceDeviceConfiguration] class.
//
// # Methods
//
//   - [IVZPowerSourceDeviceConfiguration._init]
//   - [IVZPowerSourceDeviceConfiguration._powerSourceDevice]
//   - [IVZPowerSourceDeviceConfiguration.EncodeWithEncoder]
//   - [IVZPowerSourceDeviceConfiguration.MakePowerSourceDeviceForVirtualMachinePowerSourceDeviceIndex]
//   - [IVZPowerSourceDeviceConfiguration.DebugDescription]
//   - [IVZPowerSourceDeviceConfiguration.Description]
//   - [IVZPowerSourceDeviceConfiguration.Hash]
//   - [IVZPowerSourceDeviceConfiguration.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZPowerSourceDeviceConfiguration
type IVZPowerSourceDeviceConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	_powerSourceDevice() objectivec.IObject
	EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject
	MakePowerSourceDeviceForVirtualMachinePowerSourceDeviceIndex(machine objectivec.IObject, index uint64) objectivec.IObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (v VZPowerSourceDeviceConfiguration) Init() VZPowerSourceDeviceConfiguration {
	rv := objc.Send[VZPowerSourceDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZPowerSourceDeviceConfiguration) Autorelease() VZPowerSourceDeviceConfiguration {
	rv := objc.Send[VZPowerSourceDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZPowerSourceDeviceConfiguration creates a new VZPowerSourceDeviceConfiguration instance.
func NewVZPowerSourceDeviceConfiguration() VZPowerSourceDeviceConfiguration {
	class := getVZPowerSourceDeviceConfigurationClass()
	rv := objc.Send[VZPowerSourceDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZPowerSourceDeviceConfiguration/_init
func (v VZPowerSourceDeviceConfiguration) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZPowerSourceDeviceConfiguration/encodeWithEncoder:
func (v VZPowerSourceDeviceConfiguration) EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("encodeWithEncoder:"), encoder)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZPowerSourceDeviceConfiguration/makePowerSourceDeviceForVirtualMachine:powerSourceDeviceIndex:
func (v VZPowerSourceDeviceConfiguration) MakePowerSourceDeviceForVirtualMachinePowerSourceDeviceIndex(machine objectivec.IObject, index uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("makePowerSourceDeviceForVirtualMachine:powerSourceDeviceIndex:"), machine, index)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZPowerSourceDeviceConfiguration/_powerSourceDevice
func (v VZPowerSourceDeviceConfiguration) _powerSourceDevice() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_powerSourceDevice"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZPowerSourceDeviceConfiguration/debugDescription
func (v VZPowerSourceDeviceConfiguration) DebugDescription() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/_VZPowerSourceDeviceConfiguration/description
func (v VZPowerSourceDeviceConfiguration) Description() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/_VZPowerSourceDeviceConfiguration/hash
func (v VZPowerSourceDeviceConfiguration) Hash() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZPowerSourceDeviceConfiguration/superclass
func (v VZPowerSourceDeviceConfiguration) Superclass() objc.Class {
	rv := objc.Send[objc.Class](v.ID, objc.Sel("superclass"))
	return rv
}
