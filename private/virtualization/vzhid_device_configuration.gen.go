// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZHIDDeviceConfiguration] class.
var (
	_VZHIDDeviceConfigurationClass     VZHIDDeviceConfigurationClass
	_VZHIDDeviceConfigurationClassOnce sync.Once
)

func getVZHIDDeviceConfigurationClass() VZHIDDeviceConfigurationClass {
	_VZHIDDeviceConfigurationClassOnce.Do(func() {
		_VZHIDDeviceConfigurationClass = VZHIDDeviceConfigurationClass{class: objc.GetClass("_VZHIDDeviceConfiguration")}
	})
	return _VZHIDDeviceConfigurationClass
}

// GetVZHIDDeviceConfigurationClass returns the class object for _VZHIDDeviceConfiguration.
func GetVZHIDDeviceConfigurationClass() VZHIDDeviceConfigurationClass {
	return getVZHIDDeviceConfigurationClass()
}

type VZHIDDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZHIDDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZHIDDeviceConfigurationClass) Alloc() VZHIDDeviceConfiguration {
	rv := objc.Send[VZHIDDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZHIDDeviceConfiguration._hidDevice]
//   - [VZHIDDeviceConfiguration._init]
//   - [VZHIDDeviceConfiguration.EncodeWithEncoder]
//   - [VZHIDDeviceConfiguration.MakeHIDDeviceForVirtualMachineHidDeviceIndex]
//   - [VZHIDDeviceConfiguration.DebugDescription]
//   - [VZHIDDeviceConfiguration.Description]
//   - [VZHIDDeviceConfiguration.Hash]
//   - [VZHIDDeviceConfiguration.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZHIDDeviceConfiguration
type VZHIDDeviceConfiguration struct {
	objectivec.Object
}

// VZHIDDeviceConfigurationFromID constructs a [VZHIDDeviceConfiguration] from an objc.ID.
func VZHIDDeviceConfigurationFromID(id objc.ID) VZHIDDeviceConfiguration {
	return VZHIDDeviceConfiguration{objectivec.Object{ID: id}}
}

// Ensure VZHIDDeviceConfiguration implements IVZHIDDeviceConfiguration.
var _ IVZHIDDeviceConfiguration = VZHIDDeviceConfiguration{}

// An interface definition for the [VZHIDDeviceConfiguration] class.
//
// # Methods
//
//   - [IVZHIDDeviceConfiguration._hidDevice]
//   - [IVZHIDDeviceConfiguration._init]
//   - [IVZHIDDeviceConfiguration.EncodeWithEncoder]
//   - [IVZHIDDeviceConfiguration.MakeHIDDeviceForVirtualMachineHidDeviceIndex]
//   - [IVZHIDDeviceConfiguration.DebugDescription]
//   - [IVZHIDDeviceConfiguration.Description]
//   - [IVZHIDDeviceConfiguration.Hash]
//   - [IVZHIDDeviceConfiguration.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZHIDDeviceConfiguration
type IVZHIDDeviceConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	_hidDevice() objectivec.IObject
	_init() objectivec.IObject
	EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject
	MakeHIDDeviceForVirtualMachineHidDeviceIndex(machine objectivec.IObject, index uint64) objectivec.IObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (v VZHIDDeviceConfiguration) Init() VZHIDDeviceConfiguration {
	rv := objc.Send[VZHIDDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZHIDDeviceConfiguration) Autorelease() VZHIDDeviceConfiguration {
	rv := objc.Send[VZHIDDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZHIDDeviceConfiguration creates a new VZHIDDeviceConfiguration instance.
func NewVZHIDDeviceConfiguration() VZHIDDeviceConfiguration {
	class := getVZHIDDeviceConfigurationClass()
	rv := objc.Send[VZHIDDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZHIDDeviceConfiguration/_init
func (v VZHIDDeviceConfiguration) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZHIDDeviceConfiguration/encodeWithEncoder:
func (v VZHIDDeviceConfiguration) EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("encodeWithEncoder:"), encoder)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZHIDDeviceConfiguration/makeHIDDeviceForVirtualMachine:hidDeviceIndex:
func (v VZHIDDeviceConfiguration) MakeHIDDeviceForVirtualMachineHidDeviceIndex(machine objectivec.IObject, index uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("makeHIDDeviceForVirtualMachine:hidDeviceIndex:"), machine, index)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZHIDDeviceConfiguration/_hidDevice
func (v VZHIDDeviceConfiguration) _hidDevice() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_hidDevice"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZHIDDeviceConfiguration/debugDescription
func (v VZHIDDeviceConfiguration) DebugDescription() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/_VZHIDDeviceConfiguration/description
func (v VZHIDDeviceConfiguration) Description() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/_VZHIDDeviceConfiguration/hash
func (v VZHIDDeviceConfiguration) Hash() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZHIDDeviceConfiguration/superclass
func (v VZHIDDeviceConfiguration) Superclass() objc.Class {
	rv := objc.Send[objc.Class](v.ID, objc.Sel("superclass"))
	return rv
}
