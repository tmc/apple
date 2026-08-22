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

// The class instance for the [IOUSBHostCIEndpointStateMachine] class.
var (
	_IOUSBHostCIEndpointStateMachineClass     IOUSBHostCIEndpointStateMachineClass
	_IOUSBHostCIEndpointStateMachineClassOnce sync.Once
)

func getIOUSBHostCIEndpointStateMachineClass() IOUSBHostCIEndpointStateMachineClass {
	_IOUSBHostCIEndpointStateMachineClassOnce.Do(func() {
		_IOUSBHostCIEndpointStateMachineClass = IOUSBHostCIEndpointStateMachineClass{class: objc.GetClass("IOUSBHostCIEndpointStateMachine")}
	})
	return _IOUSBHostCIEndpointStateMachineClass
}

// GetIOUSBHostCIEndpointStateMachineClass returns the class object for IOUSBHostCIEndpointStateMachine.
func GetIOUSBHostCIEndpointStateMachineClass() IOUSBHostCIEndpointStateMachineClass {
	return getIOUSBHostCIEndpointStateMachineClass()
}

type IOUSBHostCIEndpointStateMachineClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOUSBHostCIEndpointStateMachineClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOUSBHostCIEndpointStateMachineClass) Alloc() IOUSBHostCIEndpointStateMachine {
	rv := objc.Send[IOUSBHostCIEndpointStateMachine](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [IOUSBHostCIEndpointStateMachine.ControllerInterface]
//   - [IOUSBHostCIEndpointStateMachine.CurrentTransferMessage]
//   - [IOUSBHostCIEndpointStateMachine.DeviceAddress]
//   - [IOUSBHostCIEndpointStateMachine.EndpointAddress]
//   - [IOUSBHostCIEndpointStateMachine.EndpointState]
//
// # Instance Methods
//
//   - [IOUSBHostCIEndpointStateMachine.EnqueueTransferCompletionForMessageStatusTransferLengthError]
//   - [IOUSBHostCIEndpointStateMachine.InspectCommandError]
//   - [IOUSBHostCIEndpointStateMachine.ProcessDoorbellError]
//   - [IOUSBHostCIEndpointStateMachine.RespondToCommandStatusError]
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIEndpointStateMachine
type IOUSBHostCIEndpointStateMachine struct {
	objectivec.Object
}

// IOUSBHostCIEndpointStateMachineFromID constructs a [IOUSBHostCIEndpointStateMachine] from an objc.ID.
func IOUSBHostCIEndpointStateMachineFromID(id objc.ID) IOUSBHostCIEndpointStateMachine {
	return IOUSBHostCIEndpointStateMachine{objectivec.Object{ID: id}}
}

// NOTE: IOUSBHostCIEndpointStateMachine adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [IOUSBHostCIEndpointStateMachine] class.
//
// # Instance Properties
//
//   - [IIOUSBHostCIEndpointStateMachine.ControllerInterface]
//   - [IIOUSBHostCIEndpointStateMachine.CurrentTransferMessage]
//   - [IIOUSBHostCIEndpointStateMachine.DeviceAddress]
//   - [IIOUSBHostCIEndpointStateMachine.EndpointAddress]
//   - [IIOUSBHostCIEndpointStateMachine.EndpointState]
//
// # Instance Methods
//
//   - [IIOUSBHostCIEndpointStateMachine.EnqueueTransferCompletionForMessageStatusTransferLengthError]
//   - [IIOUSBHostCIEndpointStateMachine.InspectCommandError]
//   - [IIOUSBHostCIEndpointStateMachine.ProcessDoorbellError]
//   - [IIOUSBHostCIEndpointStateMachine.RespondToCommandStatusError]
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIEndpointStateMachine
type IIOUSBHostCIEndpointStateMachine interface {
	objectivec.IObject

	// Topic: Instance Properties

	ControllerInterface() IIOUSBHostControllerInterface
	CurrentTransferMessage() *IOUSBHostCIMessage
	DeviceAddress() uint
	EndpointAddress() uint
	EndpointState() IOUSBHostCIEndpointState

	// Topic: Instance Methods

	EnqueueTransferCompletionForMessageStatusTransferLengthError(message *IOUSBHostCIMessage, status IOUSBHostCIMessageStatus, transferLength uint) (bool, error)
	InspectCommandError(command *IOUSBHostCIMessage) (bool, error)
	ProcessDoorbellError(doorbell uint32) (bool, error)
	RespondToCommandStatusError(command *IOUSBHostCIMessage, status IOUSBHostCIMessageStatus) (bool, error)
}

// Init initializes the instance.
func (u IOUSBHostCIEndpointStateMachine) Init() IOUSBHostCIEndpointStateMachine {
	rv := objc.Send[IOUSBHostCIEndpointStateMachine](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u IOUSBHostCIEndpointStateMachine) Autorelease() IOUSBHostCIEndpointStateMachine {
	rv := objc.Send[IOUSBHostCIEndpointStateMachine](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOUSBHostCIEndpointStateMachine creates a new IOUSBHostCIEndpointStateMachine instance.
func NewIOUSBHostCIEndpointStateMachine() IOUSBHostCIEndpointStateMachine {
	class := getIOUSBHostCIEndpointStateMachineClass()
	rv := objc.Send[IOUSBHostCIEndpointStateMachine](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIEndpointStateMachine/enqueueTransferCompletion(for:status:transferLength:)
func (u IOUSBHostCIEndpointStateMachine) EnqueueTransferCompletionForMessageStatusTransferLengthError(message *IOUSBHostCIMessage, status IOUSBHostCIMessageStatus, transferLength uint) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](u.ID, objc.Sel("enqueueTransferCompletionForMessage:status:transferLength:error:"), unsafe.Pointer(message), status, transferLength, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("enqueueTransferCompletionForMessage:status:transferLength:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIEndpointStateMachine/inspectCommand(_:)
func (u IOUSBHostCIEndpointStateMachine) InspectCommandError(command *IOUSBHostCIMessage) (bool, error) {
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

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIEndpointStateMachine/processDoorbell(_:)
func (u IOUSBHostCIEndpointStateMachine) ProcessDoorbellError(doorbell uint32) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](u.ID, objc.Sel("processDoorbell:error:"), doorbell, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("processDoorbell:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIEndpointStateMachine/respond(toCommand:status:)
func (u IOUSBHostCIEndpointStateMachine) RespondToCommandStatusError(command *IOUSBHostCIMessage, status IOUSBHostCIMessageStatus) (bool, error) {
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

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIEndpointStateMachine/controllerInterface
func (u IOUSBHostCIEndpointStateMachine) ControllerInterface() IIOUSBHostControllerInterface {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("controllerInterface"))
	return IOUSBHostControllerInterfaceFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIEndpointStateMachine/currentTransferMessage
func (u IOUSBHostCIEndpointStateMachine) CurrentTransferMessage() *IOUSBHostCIMessage {
	rv := objc.Send[unsafe.Pointer](u.ID, objc.Sel("currentTransferMessage"))
	return (*IOUSBHostCIMessage)(rv)
}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIEndpointStateMachine/deviceAddress
func (u IOUSBHostCIEndpointStateMachine) DeviceAddress() uint {
	rv := objc.Send[uint](u.ID, objc.Sel("deviceAddress"))
	return rv
}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIEndpointStateMachine/endpointAddress
func (u IOUSBHostCIEndpointStateMachine) EndpointAddress() uint {
	rv := objc.Send[uint](u.ID, objc.Sel("endpointAddress"))
	return rv
}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostCIEndpointStateMachine/endpointState
func (u IOUSBHostCIEndpointStateMachine) EndpointState() IOUSBHostCIEndpointState {
	rv := objc.Send[IOUSBHostCIEndpointState](u.ID, objc.Sel("endpointState"))
	return IOUSBHostCIEndpointState(rv)
}
