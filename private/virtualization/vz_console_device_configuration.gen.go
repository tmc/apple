// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
	rv := objc.SendIfResponds[VZConsoleDeviceConfiguration](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZConsoleDeviceConfiguration._consoleDevice]
//   - [VZConsoleDeviceConfiguration._init]
//   - [VZConsoleDeviceConfiguration.MakeConsoleDeviceForVirtualMachineConsoleDeviceIndex]
//   - [VZConsoleDeviceConfiguration.ValidateWithError]
//   - [VZConsoleDeviceConfiguration.DebugDescription]
//   - [VZConsoleDeviceConfiguration.Description]
//   - [VZConsoleDeviceConfiguration.Hash]
//   - [VZConsoleDeviceConfiguration.Superclass]
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
//   - [IVZConsoleDeviceConfiguration._consoleDevice]
//   - [IVZConsoleDeviceConfiguration._init]
//   - [IVZConsoleDeviceConfiguration.MakeConsoleDeviceForVirtualMachineConsoleDeviceIndex]
//   - [IVZConsoleDeviceConfiguration.ValidateWithError]
//   - [IVZConsoleDeviceConfiguration.DebugDescription]
//   - [IVZConsoleDeviceConfiguration.Description]
//   - [IVZConsoleDeviceConfiguration.Hash]
//   - [IVZConsoleDeviceConfiguration.Superclass]
type IVZConsoleDeviceConfiguration interface {
	objectivec.IObject

	// Topic: Methods

	_consoleDevice() unsafe.Pointer
	_init() objectivec.IObject
	MakeConsoleDeviceForVirtualMachineConsoleDeviceIndex(machine objectivec.IObject, index uint64) objectivec.IObject
	ValidateWithError() (bool, error)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (v VZConsoleDeviceConfiguration) Init() VZConsoleDeviceConfiguration {
	rv := objc.SendIfResponds[VZConsoleDeviceConfiguration](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZConsoleDeviceConfiguration) Autorelease() VZConsoleDeviceConfiguration {
	rv := objc.SendIfResponds[VZConsoleDeviceConfiguration](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZConsoleDeviceConfiguration creates a new VZConsoleDeviceConfiguration instance.
func NewVZConsoleDeviceConfiguration() VZConsoleDeviceConfiguration {
	class := getVZConsoleDeviceConfigurationClass()
	rv := objc.SendIfResponds[VZConsoleDeviceConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZConsoleDeviceConfiguration) _init() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}
func (v VZConsoleDeviceConfiguration) MakeConsoleDeviceForVirtualMachineConsoleDeviceIndex(machine objectivec.IObject, index uint64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("makeConsoleDeviceForVirtualMachine:consoleDeviceIndex:"), machine, index)
	return objectivec.Object{ID: rv}
}
func (v VZConsoleDeviceConfiguration) ValidateWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("validateWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("validateWithError: returned NO with nil NSError")
	}
	return rv, nil

}

func (v VZConsoleDeviceConfiguration) _consoleDevice() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](v.ID, objc.Sel("_consoleDevice"))
	return rv
}

// CanConsoleDevice reports whether the receiver responds to the private selector _consoleDevice.
func (v VZConsoleDeviceConfiguration) CanConsoleDevice() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_consoleDevice"))
}

// ConsoleDevice is an exported wrapper for the private property _consoleDevice.
func (v VZConsoleDeviceConfiguration) ConsoleDevice() (unsafe.Pointer, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_consoleDevice")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_consoleDevice"}
	}
	return v._consoleDevice(), nil
}
func (v VZConsoleDeviceConfiguration) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZConsoleDeviceConfiguration) Description() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZConsoleDeviceConfiguration) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](v.ID, objc.Sel("hash"))
	return rv
}
func (v VZConsoleDeviceConfiguration) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](v.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
