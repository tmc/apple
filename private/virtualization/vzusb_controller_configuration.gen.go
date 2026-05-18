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
func (v VZUSBControllerConfiguration) Init() VZUSBControllerConfiguration {
	rv := objc.Send[VZUSBControllerConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZUSBControllerConfiguration) Autorelease() VZUSBControllerConfiguration {
	rv := objc.Send[VZUSBControllerConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZUSBControllerConfiguration creates a new VZUSBControllerConfiguration instance.
func NewVZUSBControllerConfiguration() VZUSBControllerConfiguration {
	class := getVZUSBControllerConfigurationClass()
	rv := objc.Send[VZUSBControllerConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZUSBControllerConfiguration/_init
func (v VZUSBControllerConfiguration) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/VZUSBControllerConfiguration/makeUSBControllerForVirtualMachine:usbControllerIndex:usbDevices:
func (v VZUSBControllerConfiguration) MakeUSBControllerForVirtualMachineUsbControllerIndexUsbDevices(machine objectivec.IObject, index uint64, devices objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("makeUSBControllerForVirtualMachine:usbControllerIndex:usbDevices:"), machine, index, devices)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/VZUSBControllerConfiguration/_usbDevices
func (v VZUSBControllerConfiguration) _usbDevices() foundation.INSArray {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_usbDevices"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// CanUsbDevices reports whether the receiver responds to the private selector _usbDevices.
func (v VZUSBControllerConfiguration) CanUsbDevices() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_usbDevices"))
}

// UsbDevices is an exported wrapper for the private property _usbDevices.
func (v VZUSBControllerConfiguration) UsbDevices() (foundation.INSArray, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_usbDevices")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_usbDevices"}
	}
	return v._usbDevices(), nil
}

// See: https://developer.apple.com/documentation/Virtualization/VZUSBControllerConfiguration/debugDescription
func (v VZUSBControllerConfiguration) DebugDescription() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/VZUSBControllerConfiguration/description
func (v VZUSBControllerConfiguration) Description() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/VZUSBControllerConfiguration/hash
func (v VZUSBControllerConfiguration) Hash() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZUSBControllerConfiguration/superclass
func (v VZUSBControllerConfiguration) Superclass() objc.Class {
	rv := objc.Send[objc.Class](v.ID, objc.Sel("superclass"))
	return rv
}
