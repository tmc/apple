// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/kernel"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOBluetoothDevicePair] class.
var (
	_IOBluetoothDevicePairClass     IOBluetoothDevicePairClass
	_IOBluetoothDevicePairClassOnce sync.Once
)

func getIOBluetoothDevicePairClass() IOBluetoothDevicePairClass {
	_IOBluetoothDevicePairClassOnce.Do(func() {
		_IOBluetoothDevicePairClass = IOBluetoothDevicePairClass{class: objc.GetClass("IOBluetoothDevicePair")}
	})
	return _IOBluetoothDevicePairClass
}

// GetIOBluetoothDevicePairClass returns the class object for IOBluetoothDevicePair.
func GetIOBluetoothDevicePairClass() IOBluetoothDevicePairClass {
	return getIOBluetoothDevicePairClass()
}

type IOBluetoothDevicePairClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOBluetoothDevicePairClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOBluetoothDevicePairClass) Alloc() IOBluetoothDevicePair {
	rv := objc.Send[IOBluetoothDevicePair](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// An instance of IOBluetoothDevicePair represents a pairing attempt to a
// remote Bluetooth device.
//
// # Overview
//
// Use the IOBluetoothDevicePair object to attempt to pair with any Bluetooth
// device. Once -start is invoked on it, progress is returned to the delegate
// via the messages defined below. This object enables you to pair with
// devices within your application without having to use the standard panels
// provided by the IOBluetoothUI framework, allowing you to write custom UI to
// select devices, and still handle the ability to perform device pairings.
//
// Of note is that this object MAY attempt to perform two low-level pairings,
// depending on the type of device you are attempting to pair. This is
// inconsequential to your code, however, as it occurs automatically and does
// not change the messaging.
//
// Once started, the pairing can be stopped. This will set the delegate to nil
// and then attempt to disconnect from the device if already connected.
//
// # Instance Properties
//
//   - [IOBluetoothDevicePair.Delegate]
//   - [IOBluetoothDevicePair.SetDelegate]
//
// # Instance Methods
//
//   - [IOBluetoothDevicePair.Device]: Get the IOBluetoothDevice being used by the object.
//   - [IOBluetoothDevicePair.ReplyPINCodePINCode]: This is the required reply to the devicePairingPINCodeRequest delegate message. Set the PIN code to use during pairing if required.
//   - [IOBluetoothDevicePair.ReplyUserConfirmation]: This is the required reply to the devicePairingUserConfirmationRequest delegate message.
//   - [IOBluetoothDevicePair.SetDevice]: Set the device object to pair with. It is retained by the object.
//   - [IOBluetoothDevicePair.Start]: Kicks off the pairing with the device.
//   - [IOBluetoothDevicePair.Stop]: Stops the current pairing. Removes the delegate and disconnects if device was connected.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevicePair
type IOBluetoothDevicePair struct {
	objectivec.Object
}

// IOBluetoothDevicePairFromID constructs a [IOBluetoothDevicePair] from an objc.ID.
//
// An instance of IOBluetoothDevicePair represents a pairing attempt to a
// remote Bluetooth device.
func IOBluetoothDevicePairFromID(id objc.ID) IOBluetoothDevicePair {
	return IOBluetoothDevicePair{objectivec.Object{ID: id}}
}

// NOTE: IOBluetoothDevicePair adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [IOBluetoothDevicePair] class.
//
// # Instance Properties
//
//   - [IIOBluetoothDevicePair.Delegate]
//   - [IIOBluetoothDevicePair.SetDelegate]
//
// # Instance Methods
//
//   - [IIOBluetoothDevicePair.Device]: Get the IOBluetoothDevice being used by the object.
//   - [IIOBluetoothDevicePair.ReplyPINCodePINCode]: This is the required reply to the devicePairingPINCodeRequest delegate message. Set the PIN code to use during pairing if required.
//   - [IIOBluetoothDevicePair.ReplyUserConfirmation]: This is the required reply to the devicePairingUserConfirmationRequest delegate message.
//   - [IIOBluetoothDevicePair.SetDevice]: Set the device object to pair with. It is retained by the object.
//   - [IIOBluetoothDevicePair.Start]: Kicks off the pairing with the device.
//   - [IIOBluetoothDevicePair.Stop]: Stops the current pairing. Removes the delegate and disconnects if device was connected.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevicePair
type IIOBluetoothDevicePair interface {
	objectivec.IObject

	// Topic: Instance Properties

	Delegate() objectivec.IObject
	SetDelegate(value objectivec.IObject)

	// Topic: Instance Methods

	// Get the IOBluetoothDevice being used by the object.
	Device() IIOBluetoothDevice
	// This is the required reply to the devicePairingPINCodeRequest delegate message. Set the PIN code to use during pairing if required.
	ReplyPINCodePINCode(PINCodeSize uint, PINCode *BluetoothPINCode)
	// This is the required reply to the devicePairingUserConfirmationRequest delegate message.
	ReplyUserConfirmation(reply bool)
	// Set the device object to pair with. It is retained by the object.
	SetDevice(inDevice IIOBluetoothDevice)
	// Kicks off the pairing with the device.
	Start() kernel.IOReturn
	// Stops the current pairing. Removes the delegate and disconnects if device was connected.
	Stop()
}

// Init initializes the instance.
func (b IOBluetoothDevicePair) Init() IOBluetoothDevicePair {
	rv := objc.Send[IOBluetoothDevicePair](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b IOBluetoothDevicePair) Autorelease() IOBluetoothDevicePair {
	rv := objc.Send[IOBluetoothDevicePair](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOBluetoothDevicePair creates a new IOBluetoothDevicePair instance.
func NewIOBluetoothDevicePair() IOBluetoothDevicePair {
	class := getIOBluetoothDevicePairClass()
	rv := objc.Send[IOBluetoothDevicePair](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an autorelease IOBluetoothDevicePair object with a device as the
// pairing target.
//
// device: An IOBluetoothDevice to attept to pair with. The device is retained.
//
// # Return Value
//
// Returns an IOReturn or Bluetooth error code, if the pairing could not be
// started.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevicePair/init(device:)
func NewBluetoothDevicePairWithDevice(device IIOBluetoothDevice) IOBluetoothDevicePair {
	rv := objc.Send[objc.ID](objc.ID(getIOBluetoothDevicePairClass().class), objc.Sel("pairWithDevice:"), device)
	return IOBluetoothDevicePairFromID(rv)
}

// Get the IOBluetoothDevice being used by the object.
//
// # Return Value
//
// Device The IOBluetoothDevice object that the IOBluetoothDevicePair object
// is pairing with, as specified in -setDevice: or pairWithDevice:
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevicePair/device()
func (b IOBluetoothDevicePair) Device() IIOBluetoothDevice {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("device"))
	return IOBluetoothDeviceFromID(rv)
}

// This is the required reply to the devicePairingPINCodeRequest delegate
// message. Set the PIN code to use during pairing if required.
//
// PINCodeSize: The PIN code length in octets (8 bits).
//
// PINCode: PIN code for the device. Can be up to a maximum of 128 bits.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevicePair/replyPINCode(_:pinCode:)
func (b IOBluetoothDevicePair) ReplyPINCodePINCode(PINCodeSize uint, PINCode *BluetoothPINCode) {
	objc.Send[objc.ID](b.ID, objc.Sel("replyPINCode:PINCode:"), PINCodeSize, unsafe.Pointer(PINCode))
}

// This is the required reply to the devicePairingUserConfirmationRequest
// delegate message.
//
// reply: A yes/no answer provide by the user to the numeric comparison presented.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevicePair/replyUserConfirmation(_:)
func (b IOBluetoothDevicePair) ReplyUserConfirmation(reply bool) {
	objc.Send[objc.ID](b.ID, objc.Sel("replyUserConfirmation:"), reply)
}

// Set the device object to pair with. It is retained by the object.
//
// inDevice: The IOBluetoothDevice object that the IOBluetoothDevicePair object with
// which to perform a pairing.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevicePair/setDevice(_:)
func (b IOBluetoothDevicePair) SetDevice(inDevice IIOBluetoothDevice) {
	objc.Send[objc.ID](b.ID, objc.Sel("setDevice:"), inDevice)
}

// Kicks off the pairing with the device.
//
// # Return Value
//
// Returns an IOReturn or Bluetooth error code, if the pairing could not be
// started.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevicePair/start()
func (b IOBluetoothDevicePair) Start() kernel.IOReturn {
	rv := objc.Send[kernel.IOReturn](b.ID, objc.Sel("start"))
	return kernel.IOReturn(rv)
}

// Stops the current pairing. Removes the delegate and disconnects if device
// was connected.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevicePair/stop()
func (b IOBluetoothDevicePair) Stop() {
	objc.Send[objc.ID](b.ID, objc.Sel("stop"))
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothDevicePair/delegate
func (b IOBluetoothDevicePair) Delegate() objectivec.IObject {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("delegate"))
	return objectivec.Object{ID: rv}
}
func (b IOBluetoothDevicePair) SetDelegate(value objectivec.IObject) {
	objc.Send[struct{}](b.ID, objc.Sel("setDelegate:"), value)
}
