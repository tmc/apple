// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/kernel"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOBluetoothHostController] class.
var (
	_IOBluetoothHostControllerClass     IOBluetoothHostControllerClass
	_IOBluetoothHostControllerClassOnce sync.Once
)

func getIOBluetoothHostControllerClass() IOBluetoothHostControllerClass {
	_IOBluetoothHostControllerClassOnce.Do(func() {
		_IOBluetoothHostControllerClass = IOBluetoothHostControllerClass{class: objc.GetClass("IOBluetoothHostController")}
	})
	return _IOBluetoothHostControllerClass
}

// GetIOBluetoothHostControllerClass returns the class object for IOBluetoothHostController.
func GetIOBluetoothHostControllerClass() IOBluetoothHostControllerClass {
	return getIOBluetoothHostControllerClass()
}

type IOBluetoothHostControllerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOBluetoothHostControllerClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOBluetoothHostControllerClass) Alloc() IOBluetoothHostController {
	rv := objc.Send[IOBluetoothHostController](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// This class is a representation of a Bluetooth Host Controller Interface
// that is present on the local computer (either plugged in externally or
// available internally).
//
// # Overview
//
// This object can be used to ask a Bluetooth HCI for certain pieces of
// information, and be used to make it perform certain functions.
//
// # Instance Properties
//
//   - [IOBluetoothHostController.Delegate]
//   - [IOBluetoothHostController.SetDelegate]
//   - [IOBluetoothHostController.PowerState]: Gets the controller power state
//
// # Instance Methods
//
//   - [IOBluetoothHostController.AddressAsString]: Convience routine to get the HCI controller’s Bluetooth address as an NSString object.
//   - [IOBluetoothHostController.ClassOfDevice]: Gets the current class of device value.
//   - [IOBluetoothHostController.NameAsString]: Gets the “friendly” name of HCI controller.
//   - [IOBluetoothHostController.SetClassOfDeviceForTimeInterval]: Sets the current class of device value, for the specified amount of time. Note that the time interval must be set and valid. The range of acceptable values is 30-120 seconds. Anything above or below will be rounded up, or down, as appropriate.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHostController
type IOBluetoothHostController struct {
	objectivec.Object
}

// IOBluetoothHostControllerFromID constructs a [IOBluetoothHostController] from an objc.ID.
//
// This class is a representation of a Bluetooth Host Controller Interface
// that is present on the local computer (either plugged in externally or
// available internally).
func IOBluetoothHostControllerFromID(id objc.ID) IOBluetoothHostController {
	return IOBluetoothHostController{objectivec.Object{ID: id}}
}

// NOTE: IOBluetoothHostController adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [IOBluetoothHostController] class.
//
// # Instance Properties
//
//   - [IIOBluetoothHostController.Delegate]
//   - [IIOBluetoothHostController.SetDelegate]
//   - [IIOBluetoothHostController.PowerState]: Gets the controller power state
//
// # Instance Methods
//
//   - [IIOBluetoothHostController.AddressAsString]: Convience routine to get the HCI controller’s Bluetooth address as an NSString object.
//   - [IIOBluetoothHostController.ClassOfDevice]: Gets the current class of device value.
//   - [IIOBluetoothHostController.NameAsString]: Gets the “friendly” name of HCI controller.
//   - [IIOBluetoothHostController.SetClassOfDeviceForTimeInterval]: Sets the current class of device value, for the specified amount of time. Note that the time interval must be set and valid. The range of acceptable values is 30-120 seconds. Anything above or below will be rounded up, or down, as appropriate.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHostController
type IIOBluetoothHostController interface {
	objectivec.IObject

	// Topic: Instance Properties

	Delegate() objectivec.IObject
	SetDelegate(value objectivec.IObject)
	// Gets the controller power state
	PowerState() BluetoothHCIPowerState

	// Topic: Instance Methods

	// Convience routine to get the HCI controller’s Bluetooth address as an NSString object.
	AddressAsString() string
	// Gets the current class of device value.
	ClassOfDevice() BluetoothClassOfDevice
	// Gets the “friendly” name of HCI controller.
	NameAsString() string
	// Sets the current class of device value, for the specified amount of time. Note that the time interval must be set and valid. The range of acceptable values is 30-120 seconds. Anything above or below will be rounded up, or down, as appropriate.
	SetClassOfDeviceForTimeInterval(classOfDevice BluetoothClassOfDevice, seconds foundation.NSTimeInterval) kernel.IOReturn
}

// Init initializes the instance.
func (b IOBluetoothHostController) Init() IOBluetoothHostController {
	rv := objc.Send[IOBluetoothHostController](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b IOBluetoothHostController) Autorelease() IOBluetoothHostController {
	rv := objc.Send[IOBluetoothHostController](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOBluetoothHostController creates a new IOBluetoothHostController instance.
func NewIOBluetoothHostController() IOBluetoothHostController {
	class := getIOBluetoothHostControllerClass()
	rv := objc.Send[IOBluetoothHostController](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Convience routine to get the HCI controller’s Bluetooth address as an
// NSString object.
//
// # Return Value
//
// Returns NSString *. nil if the address could not be retrieved.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHostController/addressAsString()
func (b IOBluetoothHostController) AddressAsString() string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("addressAsString"))
	return foundation.NSStringFromID(rv).String()
}

// Gets the current class of device value.
//
// # Return Value
//
// Returns the current class of device value.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHostController/classOfDevice()
func (b IOBluetoothHostController) ClassOfDevice() BluetoothClassOfDevice {
	rv := objc.Send[BluetoothClassOfDevice](b.ID, objc.Sel("classOfDevice"))
	return BluetoothClassOfDevice(rv)
}

// Gets the “friendly” name of HCI controller.
//
// # Return Value
//
// Returns NSString with the device name, nil if there is not one or it cannot
// be read.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHostController/nameAsString()
func (b IOBluetoothHostController) NameAsString() string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("nameAsString"))
	return foundation.NSStringFromID(rv).String()
}

// Sets the current class of device value, for the specified amount of time.
// Note that the time interval must be set and valid. The range of acceptable
// values is 30-120 seconds. Anything above or below will be rounded up, or
// down, as appropriate.
//
// # Return Value
//
// Returns the whether setting the class of device value was successful. 0 if
// success, error code otherwise.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHostController/setClassOfDevice(_:forTimeInterval:)
func (b IOBluetoothHostController) SetClassOfDeviceForTimeInterval(classOfDevice BluetoothClassOfDevice, seconds foundation.NSTimeInterval) kernel.IOReturn {
	rv := objc.Send[kernel.IOReturn](b.ID, objc.Sel("setClassOfDevice:forTimeInterval:"), classOfDevice, seconds)
	return kernel.IOReturn(rv)
}

// Gets the default HCI controller object.
//
// # Return Value
//
// A (autoreleased) pointer to the created IOBluetoothHostController object.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHostController/default()
func (_IOBluetoothHostControllerClass IOBluetoothHostControllerClass) DefaultController() IOBluetoothHostController {
	rv := objc.Send[objc.ID](objc.ID(_IOBluetoothHostControllerClass.class), objc.Sel("defaultController"))
	return IOBluetoothHostControllerFromID(rv)
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHostController/delegate
func (b IOBluetoothHostController) Delegate() objectivec.IObject {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("delegate"))
	return objectivec.Object{ID: rv}
}
func (b IOBluetoothHostController) SetDelegate(value objectivec.IObject) {
	objc.Send[struct{}](b.ID, objc.Sel("setDelegate:"), value)
}

// Gets the controller power state
//
// # Return Value
//
// The current controller’s power state. This will be 1 for on, or 0 for
// off. Only Apple Bluetooth adapters support power off
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHostController/powerState
func (b IOBluetoothHostController) PowerState() BluetoothHCIPowerState {
	rv := objc.Send[BluetoothHCIPowerState](b.ID, objc.Sel("powerState"))
	return BluetoothHCIPowerState(rv)
}
