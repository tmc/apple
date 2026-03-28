// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZConsoleDeviceConfiguration] class.
var (
	_VZConsoleDeviceConfigurationClass     VZConsoleDeviceConfigurationClass
	_VZConsoleDeviceConfigurationClassOnce sync.Once
)

func getVZConsoleDeviceConfigurationClass() VZConsoleDeviceConfigurationClass {
	_VZConsoleDeviceConfigurationClassOnce.Do(func() {
		_VZConsoleDeviceConfigurationClass = VZConsoleDeviceConfigurationClass{class: objc.GetClass("VZConsoleDeviceConfiguration")}
	})
	return _VZConsoleDeviceConfigurationClass
}

// GetVZConsoleDeviceConfigurationClass returns the class object for VZConsoleDeviceConfiguration.
func GetVZConsoleDeviceConfigurationClass() VZConsoleDeviceConfigurationClass {
	return getVZConsoleDeviceConfigurationClass()
}

type VZConsoleDeviceConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZConsoleDeviceConfigurationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZConsoleDeviceConfigurationClass) Alloc() VZConsoleDeviceConfiguration {
	rv := objc.Send[VZConsoleDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZConsoleDeviceConfiguration._init]
//   - [VZConsoleDeviceConfiguration.MakeConsoleDeviceForVirtualMachineConsoleDeviceIndex]
//   - [VZConsoleDeviceConfiguration.ValidateWithError]
//   - [VZConsoleDeviceConfiguration._consoleDevice]
//   - [VZConsoleDeviceConfiguration.DebugDescription]
//   - [VZConsoleDeviceConfiguration.Description]
//   - [VZConsoleDeviceConfiguration.Hash]
//   - [VZConsoleDeviceConfiguration.Superclass]
// See: https://developer.apple.com/documentation/Virtualization/VZConsoleDeviceConfiguration
type VZConsoleDeviceConfiguration struct {
	objectivec.Object
}

// VZConsoleDeviceConfigurationFromID constructs a [VZConsoleDeviceConfiguration] from an objc.ID.
func VZConsoleDeviceConfigurationFromID(id objc.ID) VZConsoleDeviceConfiguration {
	return VZConsoleDeviceConfiguration{objectivec.Object{ID: id}}
}
// Ensure VZConsoleDeviceConfiguration implements IVZConsoleDeviceConfiguration.
var _ IVZConsoleDeviceConfiguration = VZConsoleDeviceConfiguration{}

// An interface definition for the [VZConsoleDeviceConfiguration] class.
//
// # Methods
//
//   - [IVZConsoleDeviceConfiguration._init]
//   - [IVZConsoleDeviceConfiguration.MakeConsoleDeviceForVirtualMachineConsoleDeviceIndex]
//   - [IVZConsoleDeviceConfiguration.ValidateWithError]
//   - [IVZConsoleDeviceConfiguration._consoleDevice]
//   - [IVZConsoleDeviceConfiguration.DebugDescription]
//   - [IVZConsoleDeviceConfiguration.Description]
//   - [IVZConsoleDeviceConfiguration.Hash]
//   - [IVZConsoleDeviceConfiguration.Superclass]
//
// See: https://developer.apple.com/documentation/Virtualization/VZConsoleDeviceConfiguration
type IVZConsoleDeviceConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	MakeConsoleDeviceForVirtualMachineConsoleDeviceIndex(machine objectivec.IObject, index uint64) objectivec.IObject
	ValidateWithError() (bool, error)
	_consoleDevice() objectivec.IObject
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (c VZConsoleDeviceConfiguration) Init() VZConsoleDeviceConfiguration {
	rv := objc.Send[VZConsoleDeviceConfiguration](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c VZConsoleDeviceConfiguration) Autorelease() VZConsoleDeviceConfiguration {
	rv := objc.Send[VZConsoleDeviceConfiguration](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZConsoleDeviceConfiguration creates a new VZConsoleDeviceConfiguration instance.
func NewVZConsoleDeviceConfiguration() VZConsoleDeviceConfiguration {
	class := getVZConsoleDeviceConfigurationClass()
	rv := objc.Send[VZConsoleDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZConsoleDeviceConfiguration/_init
func (c VZConsoleDeviceConfiguration) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZConsoleDeviceConfiguration/makeConsoleDeviceForVirtualMachine:consoleDeviceIndex:
func (c VZConsoleDeviceConfiguration) MakeConsoleDeviceForVirtualMachineConsoleDeviceIndex(machine objectivec.IObject, index uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("makeConsoleDeviceForVirtualMachine:consoleDeviceIndex:"), machine, index)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/Virtualization/VZConsoleDeviceConfiguration/validateWithError:
func (c VZConsoleDeviceConfiguration) ValidateWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](c.ID, objc.Sel("validateWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("validateWithError: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/Virtualization/VZConsoleDeviceConfiguration/_consoleDevice
func (c VZConsoleDeviceConfiguration) _consoleDevice() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("_consoleDevice"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/Virtualization/VZConsoleDeviceConfiguration/debugDescription
func (c VZConsoleDeviceConfiguration) DebugDescription() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/VZConsoleDeviceConfiguration/description
func (c VZConsoleDeviceConfiguration) Description() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/Virtualization/VZConsoleDeviceConfiguration/hash
func (c VZConsoleDeviceConfiguration) Hash() uint64 {
	rv := objc.Send[uint64](c.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/Virtualization/VZConsoleDeviceConfiguration/superclass
func (c VZConsoleDeviceConfiguration) Superclass() objc.Class {
	rv := objc.Send[objc.Class](c.ID, objc.Sel("superclass"))
	return rv
}

