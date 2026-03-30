// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

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
//
// See: https://developer.apple.com/documentation/Virtualization/VZNetworkDeviceConfiguration
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
//
// See: https://developer.apple.com/documentation/Virtualization/VZNetworkDeviceConfiguration
type IVZNetworkDeviceConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	MakeNetworkDeviceForVirtualMachineNetworkDeviceIndex(machine objectivec.IObject, index uint64) objectivec.IObject
	_networkDevice() objectivec.IObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (n VZNetworkDeviceConfiguration) Init() VZNetworkDeviceConfiguration {
	rv := objc.Send[VZNetworkDeviceConfiguration](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n VZNetworkDeviceConfiguration) Autorelease() VZNetworkDeviceConfiguration {
	rv := objc.Send[VZNetworkDeviceConfiguration](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZNetworkDeviceConfiguration creates a new VZNetworkDeviceConfiguration instance.
func NewVZNetworkDeviceConfiguration() VZNetworkDeviceConfiguration {
	class := getVZNetworkDeviceConfigurationClass()
	rv := objc.Send[VZNetworkDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZNetworkDeviceConfiguration/_init
func (n VZNetworkDeviceConfiguration) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/VZNetworkDeviceConfiguration/makeNetworkDeviceForVirtualMachine:networkDeviceIndex:
func (n VZNetworkDeviceConfiguration) MakeNetworkDeviceForVirtualMachineNetworkDeviceIndex(machine objectivec.IObject, index uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("makeNetworkDeviceForVirtualMachine:networkDeviceIndex:"), machine, index)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/VZNetworkDeviceConfiguration/_networkDevice
func (n VZNetworkDeviceConfiguration) _networkDevice() objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("_networkDevice"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/VZNetworkDeviceConfiguration/debugDescription
func (n VZNetworkDeviceConfiguration) DebugDescription() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/VZNetworkDeviceConfiguration/description
func (n VZNetworkDeviceConfiguration) Description() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/VZNetworkDeviceConfiguration/hash
func (n VZNetworkDeviceConfiguration) Hash() uint64 {
	rv := objc.Send[uint64](n.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZNetworkDeviceConfiguration/superclass
func (n VZNetworkDeviceConfiguration) Superclass() objc.Class {
	rv := objc.Send[objc.Class](n.ID, objc.Sel("superclass"))
	return rv
}
