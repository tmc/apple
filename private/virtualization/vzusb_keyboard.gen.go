// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZUSBKeyboard] class.
var (
	_VZUSBKeyboardClass     VZUSBKeyboardClass
	_VZUSBKeyboardClassOnce sync.Once
)

func getVZUSBKeyboardClass() VZUSBKeyboardClass {
	_VZUSBKeyboardClassOnce.Do(func() {
		_VZUSBKeyboardClass = VZUSBKeyboardClass{class: objc.GetClass("_VZUSBKeyboard")}
	})
	return _VZUSBKeyboardClass
}

// GetVZUSBKeyboardClass returns the class object for _VZUSBKeyboard.
func GetVZUSBKeyboardClass() VZUSBKeyboardClass {
	return getVZUSBKeyboardClass()
}

type VZUSBKeyboardClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZUSBKeyboardClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZUSBKeyboardClass) Alloc() VZUSBKeyboard {
	rv := objc.Send[VZUSBKeyboard](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZUSBKeyboard.InitWithConfiguration]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZUSBKeyboard
type VZUSBKeyboard struct {
	VZKeyboard
}

// VZUSBKeyboardFromID constructs a [VZUSBKeyboard] from an objc.ID.
func VZUSBKeyboardFromID(id objc.ID) VZUSBKeyboard {
	return VZUSBKeyboard{VZKeyboard: VZKeyboardFromID(id)}
}

// Ensure VZUSBKeyboard implements IVZUSBKeyboard.
var _ IVZUSBKeyboard = VZUSBKeyboard{}

// An interface definition for the [VZUSBKeyboard] class.
//
// # Methods
//
//   - [IVZUSBKeyboard.InitWithConfiguration]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZUSBKeyboard
type IVZUSBKeyboard interface {
	IVZKeyboard

	// Topic: Methods

	InitWithConfiguration(configuration objectivec.IObject) VZUSBKeyboard
}

// Init initializes the instance.
func (v VZUSBKeyboard) Init() VZUSBKeyboard {
	rv := objc.Send[VZUSBKeyboard](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZUSBKeyboard) Autorelease() VZUSBKeyboard {
	rv := objc.Send[VZUSBKeyboard](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZUSBKeyboard creates a new VZUSBKeyboard instance.
func NewVZUSBKeyboard() VZUSBKeyboard {
	class := getVZUSBKeyboardClass()
	rv := objc.Send[VZUSBKeyboard](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZUSBKeyboard/initWithConfiguration:
func NewVZUSBKeyboardWithConfiguration(configuration objectivec.IObject) VZUSBKeyboard {
	instance := getVZUSBKeyboardClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	return VZUSBKeyboardFromID(rv)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZKeyboard/initWithType:virtualMachine:deviceIdentifier:
func NewVZUSBKeyboardWithTypeVirtualMachineDeviceIdentifier(type_ int64, machine objectivec.IObject, identifier uint32) VZUSBKeyboard {
	instance := getVZUSBKeyboardClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithType:virtualMachine:deviceIdentifier:"), type_, machine, identifier)
	return VZUSBKeyboardFromID(rv)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZUSBKeyboard/initWithConfiguration:
func (v VZUSBKeyboard) InitWithConfiguration(configuration objectivec.IObject) VZUSBKeyboard {
	rv := objc.Send[VZUSBKeyboard](v.ID, objc.Sel("initWithConfiguration:"), configuration)
	return rv
}
