// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZPointingDeviceConfiguration] class.
var (
	_VZPointingDeviceConfigurationClass     VZPointingDeviceConfigurationClass
	_VZPointingDeviceConfigurationClassOnce sync.Once
)

func getVZPointingDeviceConfigurationClass() VZPointingDeviceConfigurationClass {
	_VZPointingDeviceConfigurationClassOnce.Do(func() {
		_VZPointingDeviceConfigurationClass = VZPointingDeviceConfigurationClass{class: objc.GetClass("VZPointingDeviceConfiguration")}
	})
	return _VZPointingDeviceConfigurationClass
}

// GetVZPointingDeviceConfigurationClass returns the class object for VZPointingDeviceConfiguration.
func GetVZPointingDeviceConfigurationClass() VZPointingDeviceConfigurationClass {
	return getVZPointingDeviceConfigurationClass()
}

type VZPointingDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZPointingDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZPointingDeviceConfigurationClass) Alloc() VZPointingDeviceConfiguration {
	rv := objc.Send[VZPointingDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZPointingDeviceConfiguration._init]
//   - [VZPointingDeviceConfiguration._pointingDevice]
//   - [VZPointingDeviceConfiguration.MakePointingDeviceForVirtualMachinePointingDeviceIndex]
//   - [VZPointingDeviceConfiguration.DebugDescription]
//   - [VZPointingDeviceConfiguration.Description]
//   - [VZPointingDeviceConfiguration.Hash]
//   - [VZPointingDeviceConfiguration.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/VZPointingDeviceConfiguration
type VZPointingDeviceConfiguration struct {
	objectivec.Object
}

// VZPointingDeviceConfigurationFromID constructs a [VZPointingDeviceConfiguration] from an objc.ID.
func VZPointingDeviceConfigurationFromID(id objc.ID) VZPointingDeviceConfiguration {
	return VZPointingDeviceConfiguration{objectivec.Object{ID: id}}
}

// Ensure VZPointingDeviceConfiguration implements IVZPointingDeviceConfiguration.
var _ IVZPointingDeviceConfiguration = VZPointingDeviceConfiguration{}

// An interface definition for the [VZPointingDeviceConfiguration] class.
//
// # Methods
//
//   - [IVZPointingDeviceConfiguration._init]
//   - [IVZPointingDeviceConfiguration._pointingDevice]
//   - [IVZPointingDeviceConfiguration.MakePointingDeviceForVirtualMachinePointingDeviceIndex]
//   - [IVZPointingDeviceConfiguration.DebugDescription]
//   - [IVZPointingDeviceConfiguration.Description]
//   - [IVZPointingDeviceConfiguration.Hash]
//   - [IVZPointingDeviceConfiguration.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/VZPointingDeviceConfiguration
type IVZPointingDeviceConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	_pointingDevice() int
	MakePointingDeviceForVirtualMachinePointingDeviceIndex(machine objectivec.IObject, index uint64) objectivec.IObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (p VZPointingDeviceConfiguration) Init() VZPointingDeviceConfiguration {
	rv := objc.Send[VZPointingDeviceConfiguration](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p VZPointingDeviceConfiguration) Autorelease() VZPointingDeviceConfiguration {
	rv := objc.Send[VZPointingDeviceConfiguration](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZPointingDeviceConfiguration creates a new VZPointingDeviceConfiguration instance.
func NewVZPointingDeviceConfiguration() VZPointingDeviceConfiguration {
	class := getVZPointingDeviceConfigurationClass()
	rv := objc.Send[VZPointingDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZPointingDeviceConfiguration/_init
func (p VZPointingDeviceConfiguration) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/VZPointingDeviceConfiguration/makePointingDeviceForVirtualMachine:pointingDeviceIndex:
func (p VZPointingDeviceConfiguration) MakePointingDeviceForVirtualMachinePointingDeviceIndex(machine objectivec.IObject, index uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("makePointingDeviceForVirtualMachine:pointingDeviceIndex:"), machine, index)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/VZPointingDeviceConfiguration/_pointingDevice
func (p VZPointingDeviceConfiguration) _pointingDevice() int {
	rv := objc.Send[int](p.ID, objc.Sel("_pointingDevice"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZPointingDeviceConfiguration/debugDescription
func (p VZPointingDeviceConfiguration) DebugDescription() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/VZPointingDeviceConfiguration/description
func (p VZPointingDeviceConfiguration) Description() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/Virtualization/VZPointingDeviceConfiguration/hash
func (p VZPointingDeviceConfiguration) Hash() uint64 {
	rv := objc.Send[uint64](p.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZPointingDeviceConfiguration/superclass
func (p VZPointingDeviceConfiguration) Superclass() objc.Class {
	rv := objc.Send[objc.Class](p.ID, objc.Sel("superclass"))
	return rv
}
