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

// The class instance for the [IOUSBHostCIPortStateMachine] class.
var (
	_IOUSBHostCIPortStateMachineClass     IOUSBHostCIPortStateMachineClass
	_IOUSBHostCIPortStateMachineClassOnce sync.Once
)

func getIOUSBHostCIPortStateMachineClass() IOUSBHostCIPortStateMachineClass {
	_IOUSBHostCIPortStateMachineClassOnce.Do(func() {
		_IOUSBHostCIPortStateMachineClass = IOUSBHostCIPortStateMachineClass{class: objc.GetClass("IOUSBHostCIPortStateMachine")}
	})
	return _IOUSBHostCIPortStateMachineClass
}

// GetIOUSBHostCIPortStateMachineClass returns the class object for IOUSBHostCIPortStateMachine.
func GetIOUSBHostCIPortStateMachineClass() IOUSBHostCIPortStateMachineClass {
	return getIOUSBHostCIPortStateMachineClass()
}

type IOUSBHostCIPortStateMachineClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOUSBHostCIPortStateMachineClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOUSBHostCIPortStateMachineClass) Alloc() IOUSBHostCIPortStateMachine {
	rv := objc.Send[IOUSBHostCIPortStateMachine](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [IOUSBHostCIPortStateMachine.Connected]
//   - [IOUSBHostCIPortStateMachine.SetConnected]
//   - [IOUSBHostCIPortStateMachine.ControllerInterface]
//   - [IOUSBHostCIPortStateMachine.LinkState]
//   - [IOUSBHostCIPortStateMachine.Overcurrent]
//   - [IOUSBHostCIPortStateMachine.SetOvercurrent]
//   - [IOUSBHostCIPortStateMachine.PortNumber]
//   - [IOUSBHostCIPortStateMachine.PortState]
//   - [IOUSBHostCIPortStateMachine.PortStatus]
//   - [IOUSBHostCIPortStateMachine.Powered]
//   - [IOUSBHostCIPortStateMachine.SetPowered]
//   - [IOUSBHostCIPortStateMachine.Speed]
//
// # Instance Methods
//
//   - [IOUSBHostCIPortStateMachine.InspectCommandError]
//   - [IOUSBHostCIPortStateMachine.RespondToCommandStatusError]
//   - [IOUSBHostCIPortStateMachine.UpdateLinkStateSpeedInhibitLinkStateChangeError]
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIPortStateMachine
type IOUSBHostCIPortStateMachine struct {
	objectivec.Object
}

// IOUSBHostCIPortStateMachineFromID constructs a [IOUSBHostCIPortStateMachine] from an objc.ID.
func IOUSBHostCIPortStateMachineFromID(id objc.ID) IOUSBHostCIPortStateMachine {
	return IOUSBHostCIPortStateMachine{objectivec.Object{ID: id}}
}

// NOTE: IOUSBHostCIPortStateMachine adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [IOUSBHostCIPortStateMachine] class.
//
// # Instance Properties
//
//   - [IIOUSBHostCIPortStateMachine.Connected]
//   - [IIOUSBHostCIPortStateMachine.SetConnected]
//   - [IIOUSBHostCIPortStateMachine.ControllerInterface]
//   - [IIOUSBHostCIPortStateMachine.LinkState]
//   - [IIOUSBHostCIPortStateMachine.Overcurrent]
//   - [IIOUSBHostCIPortStateMachine.SetOvercurrent]
//   - [IIOUSBHostCIPortStateMachine.PortNumber]
//   - [IIOUSBHostCIPortStateMachine.PortState]
//   - [IIOUSBHostCIPortStateMachine.PortStatus]
//   - [IIOUSBHostCIPortStateMachine.Powered]
//   - [IIOUSBHostCIPortStateMachine.SetPowered]
//   - [IIOUSBHostCIPortStateMachine.Speed]
//
// # Instance Methods
//
//   - [IIOUSBHostCIPortStateMachine.InspectCommandError]
//   - [IIOUSBHostCIPortStateMachine.RespondToCommandStatusError]
//   - [IIOUSBHostCIPortStateMachine.UpdateLinkStateSpeedInhibitLinkStateChangeError]
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIPortStateMachine
type IIOUSBHostCIPortStateMachine interface {
	objectivec.IObject

	// Topic: Instance Properties

	Connected() bool
	SetConnected(value bool)
	ControllerInterface() IIOUSBHostControllerInterface
	LinkState() IOUSBHostCILinkState
	Overcurrent() bool
	SetOvercurrent(value bool)
	PortNumber() uint
	PortState() IOUSBHostCIPortState
	PortStatus() uint32
	Powered() bool
	SetPowered(value bool)
	Speed() IOUSBHostCIDeviceSpeed

	// Topic: Instance Methods

	InspectCommandError(command *IOUSBHostCIMessage) (bool, error)
	RespondToCommandStatusError(command *IOUSBHostCIMessage, status IOUSBHostCIMessageStatus) (bool, error)
	UpdateLinkStateSpeedInhibitLinkStateChangeError(linkState IOUSBHostCILinkState, speed IOUSBHostCIDeviceSpeed, inhibitLinkStateChange bool) (bool, error)
}

// Init initializes the instance.
func (u IOUSBHostCIPortStateMachine) Init() IOUSBHostCIPortStateMachine {
	rv := objc.Send[IOUSBHostCIPortStateMachine](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u IOUSBHostCIPortStateMachine) Autorelease() IOUSBHostCIPortStateMachine {
	rv := objc.Send[IOUSBHostCIPortStateMachine](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOUSBHostCIPortStateMachine creates a new IOUSBHostCIPortStateMachine instance.
func NewIOUSBHostCIPortStateMachine() IOUSBHostCIPortStateMachine {
	class := getIOUSBHostCIPortStateMachineClass()
	rv := objc.Send[IOUSBHostCIPortStateMachine](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIPortStateMachine/inspectCommand(_:)
func (u IOUSBHostCIPortStateMachine) InspectCommandError(command *IOUSBHostCIMessage) (bool, error) {
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

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIPortStateMachine/respond(toCommand:status:)
func (u IOUSBHostCIPortStateMachine) RespondToCommandStatusError(command *IOUSBHostCIMessage, status IOUSBHostCIMessageStatus) (bool, error) {
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

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIPortStateMachine/updateLinkState(_:speed:inhibitLinkStateChange:)
func (u IOUSBHostCIPortStateMachine) UpdateLinkStateSpeedInhibitLinkStateChangeError(linkState IOUSBHostCILinkState, speed IOUSBHostCIDeviceSpeed, inhibitLinkStateChange bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](u.ID, objc.Sel("updateLinkState:speed:inhibitLinkStateChange:error:"), linkState, speed, inhibitLinkStateChange, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("updateLinkState:speed:inhibitLinkStateChange:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIPortStateMachine/connected
func (u IOUSBHostCIPortStateMachine) Connected() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("connected"))
	return rv
}
func (u IOUSBHostCIPortStateMachine) SetConnected(value bool) {
	objc.Send[struct{}](u.ID, objc.Sel("setConnected:"), value)
}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIPortStateMachine/controllerInterface
func (u IOUSBHostCIPortStateMachine) ControllerInterface() IIOUSBHostControllerInterface {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("controllerInterface"))
	return IOUSBHostControllerInterfaceFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIPortStateMachine/linkState
func (u IOUSBHostCIPortStateMachine) LinkState() IOUSBHostCILinkState {
	rv := objc.Send[IOUSBHostCILinkState](u.ID, objc.Sel("linkState"))
	return IOUSBHostCILinkState(rv)
}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIPortStateMachine/overcurrent
func (u IOUSBHostCIPortStateMachine) Overcurrent() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("overcurrent"))
	return rv
}
func (u IOUSBHostCIPortStateMachine) SetOvercurrent(value bool) {
	objc.Send[struct{}](u.ID, objc.Sel("setOvercurrent:"), value)
}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIPortStateMachine/portNumber
func (u IOUSBHostCIPortStateMachine) PortNumber() uint {
	rv := objc.Send[uint](u.ID, objc.Sel("portNumber"))
	return rv
}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIPortStateMachine/portState
func (u IOUSBHostCIPortStateMachine) PortState() IOUSBHostCIPortState {
	rv := objc.Send[IOUSBHostCIPortState](u.ID, objc.Sel("portState"))
	return IOUSBHostCIPortState(rv)
}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIPortStateMachine/portStatus
func (u IOUSBHostCIPortStateMachine) PortStatus() uint32 {
	rv := objc.Send[uint32](u.ID, objc.Sel("portStatus"))
	return rv
}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIPortStateMachine/powered
func (u IOUSBHostCIPortStateMachine) Powered() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("powered"))
	return rv
}
func (u IOUSBHostCIPortStateMachine) SetPowered(value bool) {
	objc.Send[struct{}](u.ID, objc.Sel("setPowered:"), value)
}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIPortStateMachine/speed
func (u IOUSBHostCIPortStateMachine) Speed() IOUSBHostCIDeviceSpeed {
	rv := objc.Send[IOUSBHostCIDeviceSpeed](u.ID, objc.Sel("speed"))
	return IOUSBHostCIDeviceSpeed(rv)
}
