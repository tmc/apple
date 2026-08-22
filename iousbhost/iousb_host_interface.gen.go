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

// The class instance for the [IOUSBHostInterface] class.
var (
	_IOUSBHostInterfaceClass     IOUSBHostInterfaceClass
	_IOUSBHostInterfaceClassOnce sync.Once
)

func getIOUSBHostInterfaceClass() IOUSBHostInterfaceClass {
	_IOUSBHostInterfaceClassOnce.Do(func() {
		_IOUSBHostInterfaceClass = IOUSBHostInterfaceClass{class: objc.GetClass("IOUSBHostInterface")}
	})
	return _IOUSBHostInterfaceClass
}

// GetIOUSBHostInterfaceClass returns the class object for IOUSBHostInterface.
func GetIOUSBHostInterfaceClass() IOUSBHostInterfaceClass {
	return getIOUSBHostInterfaceClass()
}

type IOUSBHostInterfaceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOUSBHostInterfaceClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOUSBHostInterfaceClass) Alloc() IOUSBHostInterface {
	rv := objc.Send[IOUSBHostInterface](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// The class for accessing USB-related services.
//
// # Overview
//
// Use this class to create pipes, retrieve descriptors, send device requests,
// and enable power savings. Create an instance of the class with
// [initWithIOService:options:queue:error:interestHandler:].
//
// # Retrieving Function Descriptors
//
//   - [IOUSBHostInterface.ConfigurationDescriptor]: The configuration descriptor for the interface.
//   - [IOUSBHostInterface.InterfaceDescriptor]: The descriptor for the interface.
//
// # Managing Pipes
//
//   - [IOUSBHostInterface.SelectAlternateSettingError]: Selects an alternative setting for the interface.
//   - [IOUSBHostInterface.CopyPipeWithAddressError]: Copies a pipe for a specific endpoint address.
//
// # Enabling Power Savings
//
//   - [IOUSBHostInterface.IdleTimeout]: The current idle suspend timeout.
//   - [IOUSBHostInterface.SetIdleTimeoutError]: Sets the desired idle suspend timeout for the interface.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostInterface
//
// [initWithIOService:options:queue:error:interestHandler:]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostObject/initWithIOService:options:queue:error:interestHandler:
type IOUSBHostInterface struct {
	IOUSBHostObject
}

// IOUSBHostInterfaceFromID constructs a [IOUSBHostInterface] from an objc.ID.
//
// The class for accessing USB-related services.
func IOUSBHostInterfaceFromID(id objc.ID) IOUSBHostInterface {
	return IOUSBHostInterface{IOUSBHostObject: IOUSBHostObjectFromID(id)}
}

// NOTE: IOUSBHostInterface adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [IOUSBHostInterface] class.
//
// # Retrieving Function Descriptors
//
//   - [IIOUSBHostInterface.ConfigurationDescriptor]: The configuration descriptor for the interface.
//   - [IIOUSBHostInterface.InterfaceDescriptor]: The descriptor for the interface.
//
// # Managing Pipes
//
//   - [IIOUSBHostInterface.SelectAlternateSettingError]: Selects an alternative setting for the interface.
//   - [IIOUSBHostInterface.CopyPipeWithAddressError]: Copies a pipe for a specific endpoint address.
//
// # Enabling Power Savings
//
//   - [IIOUSBHostInterface.IdleTimeout]: The current idle suspend timeout.
//   - [IIOUSBHostInterface.SetIdleTimeoutError]: Sets the desired idle suspend timeout for the interface.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostInterface
type IIOUSBHostInterface interface {
	IIOUSBHostObject

	// Topic: Retrieving Function Descriptors

	// The configuration descriptor for the interface.
	ConfigurationDescriptor() objectivec.IObject
	// The descriptor for the interface.
	InterfaceDescriptor() objectivec.IObject

	// Topic: Managing Pipes

	// Selects an alternative setting for the interface.
	SelectAlternateSettingError(alternateSetting uint) (bool, error)
	// Copies a pipe for a specific endpoint address.
	CopyPipeWithAddressError(address uint) (IIOUSBHostPipe, error)

	// Topic: Enabling Power Savings

	// The current idle suspend timeout.
	IdleTimeout() foundation.NSTimeInterval
	// Sets the desired idle suspend timeout for the interface.
	SetIdleTimeoutError(idleTimeout foundation.NSTimeInterval) (bool, error)
}

// Init initializes the instance.
func (u IOUSBHostInterface) Init() IOUSBHostInterface {
	rv := objc.Send[IOUSBHostInterface](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u IOUSBHostInterface) Autorelease() IOUSBHostInterface {
	rv := objc.Send[IOUSBHostInterface](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOUSBHostInterface creates a new IOUSBHostInterface instance.
func NewIOUSBHostInterface() IOUSBHostInterface {
	class := getIOUSBHostInterfaceClass()
	rv := objc.Send[IOUSBHostInterface](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Selects an alternative setting for the interface.
//
// alternateSetting: The alternative interface number to activate.
//
// # Discussion
//
// Use this method to select an alternative setting for the interface. The
// operation aborts all pending input/output requests on the interface’s
// pipes, and closes all open pipes. It also selects the new alternative
// setting through the `SET_INTERFACE` control request (See USB 3.2, 9.4.10.).
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostInterface/selectAlternateSetting(_:)
func (u IOUSBHostInterface) SelectAlternateSettingError(alternateSetting uint) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](u.ID, objc.Sel("selectAlternateSetting:error:"), alternateSetting, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("selectAlternateSetting:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Copies a pipe for a specific endpoint address.
//
// address: The endpoint address of the pipe.
//
// # Return Value
//
// An [IOUSBHostPipe] or nil if the system can’t create the pipe.
//
// # Discussion
//
// If the pipe returns successfully, the method maintains a reference to the
// [IOUSBHostInterface].
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostInterface/copyPipe(withAddress:)
func (u IOUSBHostInterface) CopyPipeWithAddressError(address uint) (IIOUSBHostPipe, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](u.ID, objc.Sel("copyPipeWithAddress:error:"), address, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return IOUSBHostPipe{}, foundation.NSErrorFrom(errorPtr)
	}
	return IOUSBHostPipeFromID(rv), nil

}

// Sets the desired idle suspend timeout for the interface.
//
// idleTimeout: The amount of time after all pipes are idle to wait before suspending the
// device.
//
// # Discussion
//
// After the interface idles, it defers electrical suspension of the device
// for the specified duration.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostInterface/setIdleTimeout(_:)
func (u IOUSBHostInterface) SetIdleTimeoutError(idleTimeout foundation.NSTimeInterval) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](u.ID, objc.Sel("setIdleTimeout:error:"), idleTimeout, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setIdleTimeout:error: returned NO with nil NSError")
	}
	return rv, nil

}

// The configuration descriptor for the interface.
//
// # Return Value
//
// A pointer to the device’s configuration descriptor.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostInterface/configurationDescriptor
func (u IOUSBHostInterface) ConfigurationDescriptor() objectivec.IObject {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("configurationDescriptor"))
	return objectivec.Object{ID: rv}
}

// The descriptor for the interface.
//
// # Return Value
//
// A pointer to the interface’s descriptor.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostInterface/interfaceDescriptor
func (u IOUSBHostInterface) InterfaceDescriptor() objectivec.IObject {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("interfaceDescriptor"))
	return objectivec.Object{ID: rv}
}

// The current idle suspend timeout.
//
// # Return Value
//
// The amount of time after all pipes are idle to wait before suspending the
// device.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostInterface/idleTimeout
func (u IOUSBHostInterface) IdleTimeout() foundation.NSTimeInterval {
	rv := objc.Send[foundation.NSTimeInterval](u.ID, objc.Sel("idleTimeout"))
	return foundation.NSTimeInterval(rv)
}
