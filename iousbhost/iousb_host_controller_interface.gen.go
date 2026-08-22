// Code generated from Apple documentation for IOUSBHost. DO NOT EDIT.

package iousbhost

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOUSBHostControllerInterface] class.
var (
	_IOUSBHostControllerInterfaceClass     IOUSBHostControllerInterfaceClass
	_IOUSBHostControllerInterfaceClassOnce sync.Once
)

func getIOUSBHostControllerInterfaceClass() IOUSBHostControllerInterfaceClass {
	_IOUSBHostControllerInterfaceClassOnce.Do(func() {
		_IOUSBHostControllerInterfaceClass = IOUSBHostControllerInterfaceClass{class: objc.GetClass("IOUSBHostControllerInterface")}
	})
	return _IOUSBHostControllerInterfaceClass
}

// GetIOUSBHostControllerInterfaceClass returns the class object for IOUSBHostControllerInterface.
func GetIOUSBHostControllerInterfaceClass() IOUSBHostControllerInterfaceClass {
	return getIOUSBHostControllerInterfaceClass()
}

type IOUSBHostControllerInterfaceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOUSBHostControllerInterfaceClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOUSBHostControllerInterfaceClass) Alloc() IOUSBHostControllerInterface {
	rv := objc.Send[IOUSBHostControllerInterface](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [IOUSBHostControllerInterface.Capabilities]
//   - [IOUSBHostControllerInterface.ControllerStateMachine]
//   - [IOUSBHostControllerInterface.InterruptRateHz]
//   - [IOUSBHostControllerInterface.SetInterruptRateHz]
//   - [IOUSBHostControllerInterface.Queue]
//   - [IOUSBHostControllerInterface.Uuid]
//
// # Instance Methods
//
//   - [IOUSBHostControllerInterface.CapabilitiesForPort]
//   - [IOUSBHostControllerInterface.DescriptionForMessage]
//   - [IOUSBHostControllerInterface.Destroy]
//   - [IOUSBHostControllerInterface.EnqueueInterruptError]
//   - [IOUSBHostControllerInterface.EnqueueInterruptExpediteError]
//   - [IOUSBHostControllerInterface.EnqueueInterruptsCountError]
//   - [IOUSBHostControllerInterface.EnqueueInterruptsCountExpediteError]
//   - [IOUSBHostControllerInterface.GetPortStateMachineForCommandError]
//   - [IOUSBHostControllerInterface.GetPortStateMachineForPortError]
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostControllerInterface
type IOUSBHostControllerInterface struct {
	objectivec.Object
}

// IOUSBHostControllerInterfaceFromID constructs a [IOUSBHostControllerInterface] from an objc.ID.
func IOUSBHostControllerInterfaceFromID(id objc.ID) IOUSBHostControllerInterface {
	return IOUSBHostControllerInterface{objectivec.Object{ID: id}}
}

// NOTE: IOUSBHostControllerInterface adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [IOUSBHostControllerInterface] class.
//
// # Instance Properties
//
//   - [IIOUSBHostControllerInterface.Capabilities]
//   - [IIOUSBHostControllerInterface.ControllerStateMachine]
//   - [IIOUSBHostControllerInterface.InterruptRateHz]
//   - [IIOUSBHostControllerInterface.SetInterruptRateHz]
//   - [IIOUSBHostControllerInterface.Queue]
//   - [IIOUSBHostControllerInterface.Uuid]
//
// # Instance Methods
//
//   - [IIOUSBHostControllerInterface.CapabilitiesForPort]
//   - [IIOUSBHostControllerInterface.DescriptionForMessage]
//   - [IIOUSBHostControllerInterface.Destroy]
//   - [IIOUSBHostControllerInterface.EnqueueInterruptError]
//   - [IIOUSBHostControllerInterface.EnqueueInterruptExpediteError]
//   - [IIOUSBHostControllerInterface.EnqueueInterruptsCountError]
//   - [IIOUSBHostControllerInterface.EnqueueInterruptsCountExpediteError]
//   - [IIOUSBHostControllerInterface.GetPortStateMachineForCommandError]
//   - [IIOUSBHostControllerInterface.GetPortStateMachineForPortError]
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostControllerInterface
type IIOUSBHostControllerInterface interface {
	objectivec.IObject

	// Topic: Instance Properties

	Capabilities() *IOUSBHostCIMessage
	ControllerStateMachine() IIOUSBHostCIControllerStateMachine
	InterruptRateHz() uint
	SetInterruptRateHz(value uint)
	Queue() dispatch.Queue
	Uuid() foundation.NSUUID

	// Topic: Instance Methods

	CapabilitiesForPort(port uint) *IOUSBHostCIMessage
	DescriptionForMessage(message *IOUSBHostCIMessage) string
	Destroy()
	EnqueueInterruptError(interrupt *IOUSBHostCIMessage) (bool, error)
	EnqueueInterruptExpediteError(interrupt *IOUSBHostCIMessage, expedite bool) (bool, error)
	EnqueueInterruptsCountError(interrupts *IOUSBHostCIMessage, count uint) (bool, error)
	EnqueueInterruptsCountExpediteError(interrupts *IOUSBHostCIMessage, count uint, expedite bool) (bool, error)
	GetPortStateMachineForCommandError(command *IOUSBHostCIMessage) (IIOUSBHostCIPortStateMachine, error)
	GetPortStateMachineForPortError(port uint) (IIOUSBHostCIPortStateMachine, error)
}

// Init initializes the instance.
func (u IOUSBHostControllerInterface) Init() IOUSBHostControllerInterface {
	rv := objc.Send[IOUSBHostControllerInterface](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u IOUSBHostControllerInterface) Autorelease() IOUSBHostControllerInterface {
	rv := objc.Send[IOUSBHostControllerInterface](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOUSBHostControllerInterface creates a new IOUSBHostControllerInterface instance.
func NewIOUSBHostControllerInterface() IOUSBHostControllerInterface {
	class := getIOUSBHostControllerInterfaceClass()
	rv := objc.Send[IOUSBHostControllerInterface](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostControllerInterface/capabilities(forPort:)
func (u IOUSBHostControllerInterface) CapabilitiesForPort(port uint) *IOUSBHostCIMessage {
	rv := objc.Send[unsafe.Pointer](u.ID, objc.Sel("capabilitiesForPort:"), port)
	return (*IOUSBHostCIMessage)(rv)
}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostControllerInterface/description(for:)
func (u IOUSBHostControllerInterface) DescriptionForMessage(message *IOUSBHostCIMessage) string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("descriptionForMessage:"), unsafe.Pointer(message))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostControllerInterface/destroy()
func (u IOUSBHostControllerInterface) Destroy() {
	objc.Send[objc.ID](u.ID, objc.Sel("destroy"))
}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostControllerInterface/enqueueInterrupt(_:)
func (u IOUSBHostControllerInterface) EnqueueInterruptError(interrupt *IOUSBHostCIMessage) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](u.ID, objc.Sel("enqueueInterrupt:error:"), unsafe.Pointer(interrupt), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("enqueueInterrupt:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostControllerInterface/enqueueInterrupt(_:expedite:)
func (u IOUSBHostControllerInterface) EnqueueInterruptExpediteError(interrupt *IOUSBHostCIMessage, expedite bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](u.ID, objc.Sel("enqueueInterrupt:expedite:error:"), unsafe.Pointer(interrupt), expedite, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("enqueueInterrupt:expedite:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostControllerInterface/enqueueInterrupts(_:count:)
func (u IOUSBHostControllerInterface) EnqueueInterruptsCountError(interrupts *IOUSBHostCIMessage, count uint) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](u.ID, objc.Sel("enqueueInterrupts:count:error:"), objc.CArray(interrupts), count, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("enqueueInterrupts:count:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostControllerInterface/enqueueInterrupts(_:count:expedite:)
func (u IOUSBHostControllerInterface) EnqueueInterruptsCountExpediteError(interrupts *IOUSBHostCIMessage, count uint, expedite bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](u.ID, objc.Sel("enqueueInterrupts:count:expedite:error:"), objc.CArray(interrupts), count, expedite, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("enqueueInterrupts:count:expedite:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostControllerInterface/getPortStateMachine(forCommand:error:)
func (u IOUSBHostControllerInterface) GetPortStateMachineForCommandError(command *IOUSBHostCIMessage) (IIOUSBHostCIPortStateMachine, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](u.ID, objc.Sel("getPortStateMachineForCommand:error:"), unsafe.Pointer(command), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return IOUSBHostCIPortStateMachine{}, foundation.NSErrorFrom(errorPtr)
	}
	return IOUSBHostCIPortStateMachineFromID(rv), nil

}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostControllerInterface/getPortStateMachine(forPort:error:)
func (u IOUSBHostControllerInterface) GetPortStateMachineForPortError(port uint) (IIOUSBHostCIPortStateMachine, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](u.ID, objc.Sel("getPortStateMachineForPort:error:"), port, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return IOUSBHostCIPortStateMachine{}, foundation.NSErrorFrom(errorPtr)
	}
	return IOUSBHostCIPortStateMachineFromID(rv), nil

}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostControllerInterface/capabilities
func (u IOUSBHostControllerInterface) Capabilities() *IOUSBHostCIMessage {
	rv := objc.Send[unsafe.Pointer](u.ID, objc.Sel("capabilities"))
	return (*IOUSBHostCIMessage)(rv)
}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostControllerInterface/controllerStateMachine
func (u IOUSBHostControllerInterface) ControllerStateMachine() IIOUSBHostCIControllerStateMachine {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("controllerStateMachine"))
	return IOUSBHostCIControllerStateMachineFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostControllerInterface/interruptRateHz
func (u IOUSBHostControllerInterface) InterruptRateHz() uint {
	rv := objc.Send[uint](u.ID, objc.Sel("interruptRateHz"))
	return rv
}
func (u IOUSBHostControllerInterface) SetInterruptRateHz(value uint) {
	objc.Send[struct{}](u.ID, objc.Sel("setInterruptRateHz:"), value)
}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostControllerInterface/queue
func (u IOUSBHostControllerInterface) Queue() dispatch.Queue {
	rv := objc.Send[uintptr](u.ID, objc.Sel("queue"))
	return dispatch.QueueFromHandle(rv)
}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostControllerInterface/uuid
func (u IOUSBHostControllerInterface) Uuid() foundation.NSUUID {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("uuid"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}
