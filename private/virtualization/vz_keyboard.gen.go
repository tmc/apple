// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZKeyboard] class.
var (
	_VZKeyboardClass     VZKeyboardClass
	_VZKeyboardClassOnce sync.Once
)

func getVZKeyboardClass() VZKeyboardClass {
	_VZKeyboardClassOnce.Do(func() {
		_VZKeyboardClass = VZKeyboardClass{class: objc.GetClass("_VZKeyboard")}
	})
	return _VZKeyboardClass
}

// GetVZKeyboardClass returns the class object for _VZKeyboard.
func GetVZKeyboardClass() VZKeyboardClass {
	return getVZKeyboardClass()
}

type VZKeyboardClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZKeyboardClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZKeyboardClass) Alloc() VZKeyboard {
	rv := objc.Send[VZKeyboard](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZKeyboard.ActiveIndicators]
//   - [VZKeyboard.SendKeyEvents]
//   - [VZKeyboard.Type]
//   - [VZKeyboard.InitWithTypeVirtualMachineDeviceIdentifier]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZKeyboard
type VZKeyboard struct {
	objectivec.Object
}

// VZKeyboardFromID constructs a [VZKeyboard] from an objc.ID.
func VZKeyboardFromID(id objc.ID) VZKeyboard {
	return VZKeyboard{objectivec.Object{ID: id}}
}

// Ensure VZKeyboard implements IVZKeyboard.
var _ IVZKeyboard = VZKeyboard{}

// An interface definition for the [VZKeyboard] class.
//
// # Methods
//
//   - [IVZKeyboard.ActiveIndicators]
//   - [IVZKeyboard.SendKeyEvents]
//   - [IVZKeyboard.Type]
//   - [IVZKeyboard.InitWithTypeVirtualMachineDeviceIdentifier]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZKeyboard
type IVZKeyboard interface {
	objectivec.IObject

	// Topic: Methods

	ActiveIndicators() uint64
	SendKeyEvents(events objectivec.IObject)
	Type() int64
	InitWithTypeVirtualMachineDeviceIdentifier(type_ int64, machine objectivec.IObject, identifier uint32) VZKeyboard
}

// Init initializes the instance.
func (v VZKeyboard) Init() VZKeyboard {
	rv := objc.Send[VZKeyboard](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZKeyboard) Autorelease() VZKeyboard {
	rv := objc.Send[VZKeyboard](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZKeyboard creates a new VZKeyboard instance.
func NewVZKeyboard() VZKeyboard {
	class := getVZKeyboardClass()
	rv := objc.Send[VZKeyboard](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZKeyboard/initWithType:virtualMachine:deviceIdentifier:
func NewVZKeyboardWithTypeVirtualMachineDeviceIdentifier(type_ int64, machine objectivec.IObject, identifier uint32) VZKeyboard {
	instance := getVZKeyboardClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithType:virtualMachine:deviceIdentifier:"), type_, machine, identifier)
	return VZKeyboardFromID(rv)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZKeyboard/sendKeyEvents:
func (v VZKeyboard) SendKeyEvents(events objectivec.IObject) {
	objc.Send[objc.ID](v.ID, objc.Sel("sendKeyEvents:"), events)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZKeyboard/initWithType:virtualMachine:deviceIdentifier:
func (v VZKeyboard) InitWithTypeVirtualMachineDeviceIdentifier(type_ int64, machine objectivec.IObject, identifier uint32) VZKeyboard {
	rv := objc.Send[VZKeyboard](v.ID, objc.Sel("initWithType:virtualMachine:deviceIdentifier:"), type_, machine, identifier)
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZKeyboard/activeIndicators
func (v VZKeyboard) ActiveIndicators() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("activeIndicators"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZKeyboard/type
func (v VZKeyboard) Type() int64 {
	rv := objc.Send[int64](v.ID, objc.Sel("type"))
	return rv
}
