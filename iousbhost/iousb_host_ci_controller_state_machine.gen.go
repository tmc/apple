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

// The class instance for the [IOUSBHostCIControllerStateMachine] class.
var (
	_IOUSBHostCIControllerStateMachineClass     IOUSBHostCIControllerStateMachineClass
	_IOUSBHostCIControllerStateMachineClassOnce sync.Once
)

func getIOUSBHostCIControllerStateMachineClass() IOUSBHostCIControllerStateMachineClass {
	_IOUSBHostCIControllerStateMachineClassOnce.Do(func() {
		_IOUSBHostCIControllerStateMachineClass = IOUSBHostCIControllerStateMachineClass{class: objc.GetClass("IOUSBHostCIControllerStateMachine")}
	})
	return _IOUSBHostCIControllerStateMachineClass
}

// GetIOUSBHostCIControllerStateMachineClass returns the class object for IOUSBHostCIControllerStateMachine.
func GetIOUSBHostCIControllerStateMachineClass() IOUSBHostCIControllerStateMachineClass {
	return getIOUSBHostCIControllerStateMachineClass()
}

type IOUSBHostCIControllerStateMachineClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOUSBHostCIControllerStateMachineClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOUSBHostCIControllerStateMachineClass) Alloc() IOUSBHostCIControllerStateMachine {
	rv := objc.Send[IOUSBHostCIControllerStateMachine](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [IOUSBHostCIControllerStateMachine.ControllerInterface]
//   - [IOUSBHostCIControllerStateMachine.ControllerState]
//
// # Instance Methods
//
//   - [IOUSBHostCIControllerStateMachine.EnqueueUpdatedFrameTimestampError]
//   - [IOUSBHostCIControllerStateMachine.InspectCommandError]
//   - [IOUSBHostCIControllerStateMachine.RespondToCommandStatusError]
//   - [IOUSBHostCIControllerStateMachine.RespondToCommandStatusFrameTimestampError]
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIControllerStateMachine
type IOUSBHostCIControllerStateMachine struct {
	objectivec.Object
}

// IOUSBHostCIControllerStateMachineFromID constructs a [IOUSBHostCIControllerStateMachine] from an objc.ID.
func IOUSBHostCIControllerStateMachineFromID(id objc.ID) IOUSBHostCIControllerStateMachine {
	return IOUSBHostCIControllerStateMachine{objectivec.Object{ID: id}}
}

// NOTE: IOUSBHostCIControllerStateMachine adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [IOUSBHostCIControllerStateMachine] class.
//
// # Instance Properties
//
//   - [IIOUSBHostCIControllerStateMachine.ControllerInterface]
//   - [IIOUSBHostCIControllerStateMachine.ControllerState]
//
// # Instance Methods
//
//   - [IIOUSBHostCIControllerStateMachine.EnqueueUpdatedFrameTimestampError]
//   - [IIOUSBHostCIControllerStateMachine.InspectCommandError]
//   - [IIOUSBHostCIControllerStateMachine.RespondToCommandStatusError]
//   - [IIOUSBHostCIControllerStateMachine.RespondToCommandStatusFrameTimestampError]
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIControllerStateMachine
type IIOUSBHostCIControllerStateMachine interface {
	objectivec.IObject

	// Topic: Instance Properties

	ControllerInterface() IIOUSBHostControllerInterface
	ControllerState() IOUSBHostCIControllerState

	// Topic: Instance Methods

	EnqueueUpdatedFrameTimestampError(frame uint64, timestamp uint64) (bool, error)
	InspectCommandError(command *IOUSBHostCIMessage) (bool, error)
	RespondToCommandStatusError(command *IOUSBHostCIMessage, status IOUSBHostCIMessageStatus) (bool, error)
	RespondToCommandStatusFrameTimestampError(command *IOUSBHostCIMessage, status IOUSBHostCIMessageStatus, frame uint64, timestamp uint64) (bool, error)
}

// Init initializes the instance.
func (u IOUSBHostCIControllerStateMachine) Init() IOUSBHostCIControllerStateMachine {
	rv := objc.Send[IOUSBHostCIControllerStateMachine](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u IOUSBHostCIControllerStateMachine) Autorelease() IOUSBHostCIControllerStateMachine {
	rv := objc.Send[IOUSBHostCIControllerStateMachine](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOUSBHostCIControllerStateMachine creates a new IOUSBHostCIControllerStateMachine instance.
func NewIOUSBHostCIControllerStateMachine() IOUSBHostCIControllerStateMachine {
	class := getIOUSBHostCIControllerStateMachineClass()
	rv := objc.Send[IOUSBHostCIControllerStateMachine](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIControllerStateMachine/enqueueUpdatedFrame(_:timestamp:)
func (u IOUSBHostCIControllerStateMachine) EnqueueUpdatedFrameTimestampError(frame uint64, timestamp uint64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](u.ID, objc.Sel("enqueueUpdatedFrame:timestamp:error:"), frame, timestamp, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("enqueueUpdatedFrame:timestamp:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIControllerStateMachine/inspectCommand(_:)
func (u IOUSBHostCIControllerStateMachine) InspectCommandError(command *IOUSBHostCIMessage) (bool, error) {
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

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIControllerStateMachine/respond(toCommand:status:)
func (u IOUSBHostCIControllerStateMachine) RespondToCommandStatusError(command *IOUSBHostCIMessage, status IOUSBHostCIMessageStatus) (bool, error) {
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

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIControllerStateMachine/respond(toCommand:status:frame:timestamp:)
func (u IOUSBHostCIControllerStateMachine) RespondToCommandStatusFrameTimestampError(command *IOUSBHostCIMessage, status IOUSBHostCIMessageStatus, frame uint64, timestamp uint64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](u.ID, objc.Sel("respondToCommand:status:frame:timestamp:error:"), unsafe.Pointer(command), status, frame, timestamp, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("respondToCommand:status:frame:timestamp:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIControllerStateMachine/controllerInterface
func (u IOUSBHostCIControllerStateMachine) ControllerInterface() IIOUSBHostControllerInterface {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("controllerInterface"))
	return IOUSBHostControllerInterfaceFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIControllerStateMachine/controllerState
func (u IOUSBHostCIControllerStateMachine) ControllerState() IOUSBHostCIControllerState {
	rv := objc.Send[IOUSBHostCIControllerState](u.ID, objc.Sel("controllerState"))
	return IOUSBHostCIControllerState(rv)
}
