// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZKeyboardConfiguration] class.
var (
	_VZKeyboardConfigurationClass     VZKeyboardConfigurationClass
	_VZKeyboardConfigurationClassOnce sync.Once
)

func getVZKeyboardConfigurationClass() VZKeyboardConfigurationClass {
	_VZKeyboardConfigurationClassOnce.Do(func() {
		_VZKeyboardConfigurationClass = VZKeyboardConfigurationClass{class: objc.GetClass("VZKeyboardConfiguration")}
	})
	return _VZKeyboardConfigurationClass
}

// GetVZKeyboardConfigurationClass returns the class object for VZKeyboardConfiguration.
func GetVZKeyboardConfigurationClass() VZKeyboardConfigurationClass {
	return getVZKeyboardConfigurationClass()
}

type VZKeyboardConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZKeyboardConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZKeyboardConfigurationClass) Alloc() VZKeyboardConfiguration {
	rv := objc.Send[VZKeyboardConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZKeyboardConfiguration._init]
//   - [VZKeyboardConfiguration.MakeKeyboardForVirtualMachineDeviceIdentifier]
//   - [VZKeyboardConfiguration.DebugDescription]
//   - [VZKeyboardConfiguration.Description]
//   - [VZKeyboardConfiguration.Hash]
//   - [VZKeyboardConfiguration.Superclass]
// See: https://developer.apple.com/documentation/Virtualization/VZKeyboardConfiguration
type VZKeyboardConfiguration struct {
	objectivec.Object
}

// VZKeyboardConfigurationFromID constructs a [VZKeyboardConfiguration] from an objc.ID.
func VZKeyboardConfigurationFromID(id objc.ID) VZKeyboardConfiguration {
	return VZKeyboardConfiguration{objectivec.Object{ID: id}}
}
// Ensure VZKeyboardConfiguration implements IVZKeyboardConfiguration.
var _ IVZKeyboardConfiguration = VZKeyboardConfiguration{}

// An interface definition for the [VZKeyboardConfiguration] class.
//
// # Methods
//
//   - [IVZKeyboardConfiguration._init]
//   - [IVZKeyboardConfiguration.MakeKeyboardForVirtualMachineDeviceIdentifier]
//   - [IVZKeyboardConfiguration.DebugDescription]
//   - [IVZKeyboardConfiguration.Description]
//   - [IVZKeyboardConfiguration.Hash]
//   - [IVZKeyboardConfiguration.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/VZKeyboardConfiguration
type IVZKeyboardConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	MakeKeyboardForVirtualMachineDeviceIdentifier(machine objectivec.IObject, identifier uint32) objectivec.IObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (k VZKeyboardConfiguration) Init() VZKeyboardConfiguration {
	rv := objc.Send[VZKeyboardConfiguration](k.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (k VZKeyboardConfiguration) Autorelease() VZKeyboardConfiguration {
	rv := objc.Send[VZKeyboardConfiguration](k.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZKeyboardConfiguration creates a new VZKeyboardConfiguration instance.
func NewVZKeyboardConfiguration() VZKeyboardConfiguration {
	class := getVZKeyboardConfigurationClass()
	rv := objc.Send[VZKeyboardConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZKeyboardConfiguration/_init
func (k VZKeyboardConfiguration) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](k.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZKeyboardConfiguration/makeKeyboardForVirtualMachine:deviceIdentifier:
func (k VZKeyboardConfiguration) MakeKeyboardForVirtualMachineDeviceIdentifier(machine objectivec.IObject, identifier uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](k.ID, objc.Sel("makeKeyboardForVirtualMachine:deviceIdentifier:"), machine, identifier)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/VZKeyboardConfiguration/debugDescription
func (k VZKeyboardConfiguration) DebugDescription() string {
	rv := objc.Send[objc.ID](k.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/VZKeyboardConfiguration/description
func (k VZKeyboardConfiguration) Description() string {
	rv := objc.Send[objc.ID](k.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/VZKeyboardConfiguration/hash
func (k VZKeyboardConfiguration) Hash() uint64 {
	rv := objc.Send[uint64](k.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/VZKeyboardConfiguration/superclass
func (k VZKeyboardConfiguration) Superclass() objc.Class {
	rv := objc.Send[objc.Class](k.ID, objc.Sel("superclass"))
	return rv
}

