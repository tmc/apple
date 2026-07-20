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
//   - [VZHIDDeviceConfiguration.MakeHIDDeviceForVirtualMachineHidDeviceIndex]
//   - [VZHIDDeviceConfiguration.DebugDescription]
//   - [VZHIDDeviceConfiguration.Description]
//   - [VZHIDDeviceConfiguration.Hash]
//   - [VZHIDDeviceConfiguration.Superclass]
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
//   - [IVZHIDDeviceConfiguration.MakeHIDDeviceForVirtualMachineHidDeviceIndex]
//   - [IVZHIDDeviceConfiguration.DebugDescription]
//   - [IVZHIDDeviceConfiguration.Description]
//   - [IVZHIDDeviceConfiguration.Hash]
//   - [IVZHIDDeviceConfiguration.Superclass]
type IVZHIDDeviceConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	_hidDevice() AvpHidGenericDevice
	_init() objectivec.IObject
	MakeHIDDeviceForVirtualMachineHidDeviceIndex(machine objectivec.IObject, index uint64) objectivec.IObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
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

func (v VZHIDDeviceConfiguration) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}
func (v VZHIDDeviceConfiguration) MakeHIDDeviceForVirtualMachineHidDeviceIndex(machine objectivec.IObject, index uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("makeHIDDeviceForVirtualMachine:hidDeviceIndex:"), machine, index)
	return objectivec.Object{ID: rv}
}

func (v VZHIDDeviceConfiguration) _hidDevice() AvpHidGenericDevice {
	rv := objc.Send[AvpHidGenericDevice](v.ID, objc.Sel("_hidDevice"))
	return AvpHidGenericDevice(rv)
}

// CanHidDevice reports whether the receiver responds to the private selector _hidDevice.
func (v VZHIDDeviceConfiguration) CanHidDevice() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_hidDevice"))
}

// HidDevice is an exported wrapper for the private property _hidDevice.
func (v VZHIDDeviceConfiguration) HidDevice() (AvpHidGenericDevice, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_hidDevice")) {
		return AvpHidGenericDevice{}, &objc.UnrecognizedSelectorError{Selector: "_hidDevice"}
	}
	return v._hidDevice(), nil
}
func (v VZHIDDeviceConfiguration) DebugDescription() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZHIDDeviceConfiguration) Description() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZHIDDeviceConfiguration) Hash() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("hash"))
	return rv
}
func (v VZHIDDeviceConfiguration) Superclass() objectivec.Class {
	rv := objc.Send[objectivec.Class](v.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
