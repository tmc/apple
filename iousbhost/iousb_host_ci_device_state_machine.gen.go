// Code generated from Apple documentation for IOUSBHost. DO NOT EDIT.

package iousbhost

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOUSBHostCIDeviceStateMachine] class.
var (
	_IOUSBHostCIDeviceStateMachineClass     IOUSBHostCIDeviceStateMachineClass
	_IOUSBHostCIDeviceStateMachineClassOnce sync.Once
)

func getIOUSBHostCIDeviceStateMachineClass() IOUSBHostCIDeviceStateMachineClass {
	_IOUSBHostCIDeviceStateMachineClassOnce.Do(func() {
		_IOUSBHostCIDeviceStateMachineClass = IOUSBHostCIDeviceStateMachineClass{class: objc.GetClass("IOUSBHostCIDeviceStateMachine")}
	})
	return _IOUSBHostCIDeviceStateMachineClass
}

// GetIOUSBHostCIDeviceStateMachineClass returns the class object for IOUSBHostCIDeviceStateMachine.
func GetIOUSBHostCIDeviceStateMachineClass() IOUSBHostCIDeviceStateMachineClass {
	return getIOUSBHostCIDeviceStateMachineClass()
}

type IOUSBHostCIDeviceStateMachineClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOUSBHostCIDeviceStateMachineClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOUSBHostCIDeviceStateMachineClass) Alloc() IOUSBHostCIDeviceStateMachine {
	rv := objc.Send[IOUSBHostCIDeviceStateMachine](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [IOUSBHostCIDeviceStateMachine.CompleteRoute]
//   - [IOUSBHostCIDeviceStateMachine.ControllerInterface]
//   - [IOUSBHostCIDeviceStateMachine.DeviceAddress]
//   - [IOUSBHostCIDeviceStateMachine.DeviceState]
//
// # Instance Methods
//
//   - [IOUSBHostCIDeviceStateMachine.InspectCommandError]
//   - [IOUSBHostCIDeviceStateMachine.RespondToCommandStatusError]
//   - [IOUSBHostCIDeviceStateMachine.RespondToCommandStatusDeviceAddressError]
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIDeviceStateMachine
type IOUSBHostCIDeviceStateMachine struct {
	objectivec.Object
}

// IOUSBHostCIDeviceStateMachineFromID constructs a [IOUSBHostCIDeviceStateMachine] from an objc.ID.
func IOUSBHostCIDeviceStateMachineFromID(id objc.ID) IOUSBHostCIDeviceStateMachine {
	return IOUSBHostCIDeviceStateMachine{objectivec.Object{ID: id}}
}

// NOTE: IOUSBHostCIDeviceStateMachine adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [IOUSBHostCIDeviceStateMachine] class.
//
// # Instance Properties
//
//   - [IIOUSBHostCIDeviceStateMachine.CompleteRoute]
//   - [IIOUSBHostCIDeviceStateMachine.ControllerInterface]
//   - [IIOUSBHostCIDeviceStateMachine.DeviceAddress]
//   - [IIOUSBHostCIDeviceStateMachine.DeviceState]
//
// # Instance Methods
//
//   - [IIOUSBHostCIDeviceStateMachine.InspectCommandError]
//   - [IIOUSBHostCIDeviceStateMachine.RespondToCommandStatusError]
//   - [IIOUSBHostCIDeviceStateMachine.RespondToCommandStatusDeviceAddressError]
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIDeviceStateMachine
type IIOUSBHostCIDeviceStateMachine interface {
	objectivec.IObject

	// Topic: Instance Properties

	CompleteRoute() uint
	ControllerInterface() IIOUSBHostControllerInterface
	DeviceAddress() uint
	DeviceState() IOUSBHostCIDeviceState

	// Topic: Instance Methods

	InspectCommandError(command *IOUSBHostCIMessage) (bool, error)
	RespondToCommandStatusError(command *IOUSBHostCIMessage, status IOUSBHostCIMessageStatus) (bool, error)
	RespondToCommandStatusDeviceAddressError(command *IOUSBHostCIMessage, status IOUSBHostCIMessageStatus, deviceAddress uint) (bool, error)
}

// Init initializes the instance.
func (u IOUSBHostCIDeviceStateMachine) Init() IOUSBHostCIDeviceStateMachine {
	rv := objc.Send[IOUSBHostCIDeviceStateMachine](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u IOUSBHostCIDeviceStateMachine) Autorelease() IOUSBHostCIDeviceStateMachine {
	rv := objc.Send[IOUSBHostCIDeviceStateMachine](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOUSBHostCIDeviceStateMachine creates a new IOUSBHostCIDeviceStateMachine instance.
func NewIOUSBHostCIDeviceStateMachine() IOUSBHostCIDeviceStateMachine {
	class := getIOUSBHostCIDeviceStateMachineClass()
	rv := objc.Send[IOUSBHostCIDeviceStateMachine](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIDeviceStateMachine/inspectCommand(_:)
func (u IOUSBHostCIDeviceStateMachine) InspectCommandError(command *IOUSBHostCIMessage) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](u.ID, objc.Sel("inspectCommand:error:"), unsafe.Pointer(command), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("inspectCommand:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIDeviceStateMachine/respond(toCommand:status:)
func (u IOUSBHostCIDeviceStateMachine) RespondToCommandStatusError(command *IOUSBHostCIMessage, status IOUSBHostCIMessageStatus) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](u.ID, objc.Sel("respondToCommand:status:error:"), unsafe.Pointer(command), status, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("respondToCommand:status:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIDeviceStateMachine/respond(toCommand:status:deviceAddress:)
func (u IOUSBHostCIDeviceStateMachine) RespondToCommandStatusDeviceAddressError(command *IOUSBHostCIMessage, status IOUSBHostCIMessageStatus, deviceAddress uint) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](u.ID, objc.Sel("respondToCommand:status:deviceAddress:error:"), unsafe.Pointer(command), status, deviceAddress, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("respondToCommand:status:deviceAddress:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIDeviceStateMachine/completeRoute
func (u IOUSBHostCIDeviceStateMachine) CompleteRoute() uint {
	rv := objc.Send[uint](u.ID, objc.Sel("completeRoute"))
	return rv
}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIDeviceStateMachine/controllerInterface
func (u IOUSBHostCIDeviceStateMachine) ControllerInterface() IIOUSBHostControllerInterface {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("controllerInterface"))
	return IOUSBHostControllerInterfaceFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIDeviceStateMachine/deviceAddress
func (u IOUSBHostCIDeviceStateMachine) DeviceAddress() uint {
	rv := objc.Send[uint](u.ID, objc.Sel("deviceAddress"))
	return rv
}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIDeviceStateMachine/deviceState
func (u IOUSBHostCIDeviceStateMachine) DeviceState() IOUSBHostCIDeviceState {
	rv := objc.Send[IOUSBHostCIDeviceState](u.ID, objc.Sel("deviceState"))
	return IOUSBHostCIDeviceState(rv)
}
