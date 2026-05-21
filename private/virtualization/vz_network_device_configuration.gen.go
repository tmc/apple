// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZNetworkDeviceConfiguration] class.
var (
	_VZNetworkDeviceConfigurationClass     VZNetworkDeviceConfigurationClass
	_VZNetworkDeviceConfigurationClassOnce sync.Once
)

func getVZNetworkDeviceConfigurationClass() VZNetworkDeviceConfigurationClass {
	_VZNetworkDeviceConfigurationClassOnce.Do(func() {
		_VZNetworkDeviceConfigurationClass = VZNetworkDeviceConfigurationClass{class: objc.GetClass("VZNetworkDeviceConfiguration")}
	})
	return _VZNetworkDeviceConfigurationClass
}

// GetVZNetworkDeviceConfigurationClass returns the class object for VZNetworkDeviceConfiguration.
func GetVZNetworkDeviceConfigurationClass() VZNetworkDeviceConfigurationClass {
	return getVZNetworkDeviceConfigurationClass()
}

type VZNetworkDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZNetworkDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZNetworkDeviceConfigurationClass) Alloc() VZNetworkDeviceConfiguration {
	rv := objc.Send[VZNetworkDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZNetworkDeviceConfiguration._init]
//   - [VZNetworkDeviceConfiguration.MakeNetworkDeviceForVirtualMachineNetworkDeviceIndex]
//   - [VZNetworkDeviceConfiguration._networkDevice]
//   - [VZNetworkDeviceConfiguration.DebugDescription]
//   - [VZNetworkDeviceConfiguration.Description]
//   - [VZNetworkDeviceConfiguration.Hash]
//   - [VZNetworkDeviceConfiguration.Superclass]
type VZNetworkDeviceConfiguration struct {
	objectivec.Object
}

// VZNetworkDeviceConfigurationFromID constructs a [VZNetworkDeviceConfiguration] from an objc.ID.
func VZNetworkDeviceConfigurationFromID(id objc.ID) VZNetworkDeviceConfiguration {
	return VZNetworkDeviceConfiguration{objectivec.Object{ID: id}}
}

// Ensure VZNetworkDeviceConfiguration implements IVZNetworkDeviceConfiguration.
var _ IVZNetworkDeviceConfiguration = VZNetworkDeviceConfiguration{}

// An interface definition for the [VZNetworkDeviceConfiguration] class.
//
// # Methods
//
//   - [IVZNetworkDeviceConfiguration._init]
//   - [IVZNetworkDeviceConfiguration.MakeNetworkDeviceForVirtualMachineNetworkDeviceIndex]
//   - [IVZNetworkDeviceConfiguration._networkDevice]
//   - [IVZNetworkDeviceConfiguration.DebugDescription]
//   - [IVZNetworkDeviceConfiguration.Description]
//   - [IVZNetworkDeviceConfiguration.Hash]
//   - [IVZNetworkDeviceConfiguration.Superclass]
type IVZNetworkDeviceConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	MakeNetworkDeviceForVirtualMachineNetworkDeviceIndex(machine objectivec.IObject, index uint64) objectivec.IObject
	_networkDevice() unsafe.Pointer
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (v VZNetworkDeviceConfiguration) Init() VZNetworkDeviceConfiguration {
	rv := objc.Send[VZNetworkDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZNetworkDeviceConfiguration) Autorelease() VZNetworkDeviceConfiguration {
	rv := objc.Send[VZNetworkDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZNetworkDeviceConfiguration creates a new VZNetworkDeviceConfiguration instance.
func NewVZNetworkDeviceConfiguration() VZNetworkDeviceConfiguration {
	class := getVZNetworkDeviceConfigurationClass()
	rv := objc.Send[VZNetworkDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZNetworkDeviceConfiguration) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}
func (v VZNetworkDeviceConfiguration) MakeNetworkDeviceForVirtualMachineNetworkDeviceIndex(machine objectivec.IObject, index uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("makeNetworkDeviceForVirtualMachine:networkDeviceIndex:"), machine, index)
	return objectivec.Object{ID: rv}
}

func (v VZNetworkDeviceConfiguration) _networkDevice() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v.ID, objc.Sel("_networkDevice"))
	return rv
}

// CanNetworkDevice reports whether the receiver responds to the private selector _networkDevice.
func (v VZNetworkDeviceConfiguration) CanNetworkDevice() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_networkDevice"))
}

// NetworkDevice is an exported wrapper for the private property _networkDevice.
func (v VZNetworkDeviceConfiguration) NetworkDevice() (unsafe.Pointer, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_networkDevice")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_networkDevice"}
	}
	return v._networkDevice(), nil
}
func (v VZNetworkDeviceConfiguration) DebugDescription() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZNetworkDeviceConfiguration) Description() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZNetworkDeviceConfiguration) Hash() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("hash"))
	return rv
}
func (v VZNetworkDeviceConfiguration) Superclass() objectivec.Class {
	rv := objc.Send[objectivec.Class](v.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
