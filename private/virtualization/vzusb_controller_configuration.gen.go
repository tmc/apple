// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZUSBControllerConfiguration] class.
var (
	_VZUSBControllerConfigurationClass     VZUSBControllerConfigurationClass
	_VZUSBControllerConfigurationClassOnce sync.Once
)

func getVZUSBControllerConfigurationClass() VZUSBControllerConfigurationClass {
	_VZUSBControllerConfigurationClassOnce.Do(func() {
		_VZUSBControllerConfigurationClass = VZUSBControllerConfigurationClass{class: objc.GetClass("VZUSBControllerConfiguration")}
	})
	return _VZUSBControllerConfigurationClass
}

// GetVZUSBControllerConfigurationClass returns the class object for VZUSBControllerConfiguration.
func GetVZUSBControllerConfigurationClass() VZUSBControllerConfigurationClass {
	return getVZUSBControllerConfigurationClass()
}

type VZUSBControllerConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZUSBControllerConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZUSBControllerConfigurationClass) Alloc() VZUSBControllerConfiguration {
	rv := objc.Send[VZUSBControllerConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZUSBControllerConfiguration._init]
//   - [VZUSBControllerConfiguration._usbDevices]
//   - [VZUSBControllerConfiguration.MakeUSBControllerForVirtualMachineUsbControllerIndexUsbDevices]
//   - [VZUSBControllerConfiguration.DebugDescription]
//   - [VZUSBControllerConfiguration.Description]
//   - [VZUSBControllerConfiguration.Hash]
//   - [VZUSBControllerConfiguration.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/VZUSBControllerConfiguration
type VZUSBControllerConfiguration struct {
	objectivec.Object
}

// VZUSBControllerConfigurationFromID constructs a [VZUSBControllerConfiguration] from an objc.ID.
func VZUSBControllerConfigurationFromID(id objc.ID) VZUSBControllerConfiguration {
	return VZUSBControllerConfiguration{objectivec.Object{ID: id}}
}

// Ensure VZUSBControllerConfiguration implements IVZUSBControllerConfiguration.
var _ IVZUSBControllerConfiguration = VZUSBControllerConfiguration{}

// An interface definition for the [VZUSBControllerConfiguration] class.
//
// # Methods
//
//   - [IVZUSBControllerConfiguration._init]
//   - [IVZUSBControllerConfiguration._usbDevices]
//   - [IVZUSBControllerConfiguration.MakeUSBControllerForVirtualMachineUsbControllerIndexUsbDevices]
//   - [IVZUSBControllerConfiguration.DebugDescription]
//   - [IVZUSBControllerConfiguration.Description]
//   - [IVZUSBControllerConfiguration.Hash]
//   - [IVZUSBControllerConfiguration.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/VZUSBControllerConfiguration
type IVZUSBControllerConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	_usbDevices() foundation.INSArray
	MakeUSBControllerForVirtualMachineUsbControllerIndexUsbDevices(machine objectivec.IObject, index uint64, devices objectivec.IObject) objectivec.IObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (u VZUSBControllerConfiguration) Init() VZUSBControllerConfiguration {
	rv := objc.Send[VZUSBControllerConfiguration](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u VZUSBControllerConfiguration) Autorelease() VZUSBControllerConfiguration {
	rv := objc.Send[VZUSBControllerConfiguration](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZUSBControllerConfiguration creates a new VZUSBControllerConfiguration instance.
func NewVZUSBControllerConfiguration() VZUSBControllerConfiguration {
	class := getVZUSBControllerConfigurationClass()
	rv := objc.Send[VZUSBControllerConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZUSBControllerConfiguration/_init
func (u VZUSBControllerConfiguration) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/VZUSBControllerConfiguration/makeUSBControllerForVirtualMachine:usbControllerIndex:usbDevices:
func (u VZUSBControllerConfiguration) MakeUSBControllerForVirtualMachineUsbControllerIndexUsbDevices(machine objectivec.IObject, index uint64, devices objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("makeUSBControllerForVirtualMachine:usbControllerIndex:usbDevices:"), machine, index, devices)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/VZUSBControllerConfiguration/_usbDevices
func (u VZUSBControllerConfiguration) _usbDevices() foundation.INSArray {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("_usbDevices"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/Virtualization/VZUSBControllerConfiguration/debugDescription
func (u VZUSBControllerConfiguration) DebugDescription() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/VZUSBControllerConfiguration/description
func (u VZUSBControllerConfiguration) Description() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/VZUSBControllerConfiguration/hash
func (u VZUSBControllerConfiguration) Hash() uint64 {
	rv := objc.Send[uint64](u.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZUSBControllerConfiguration/superclass
func (u VZUSBControllerConfiguration) Superclass() objc.Class {
	rv := objc.Send[objc.Class](u.ID, objc.Sel("superclass"))
	return rv
}
