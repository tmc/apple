// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOBluetoothHandsFreeAudioGateway] class.
var (
	_IOBluetoothHandsFreeAudioGatewayClass     IOBluetoothHandsFreeAudioGatewayClass
	_IOBluetoothHandsFreeAudioGatewayClassOnce sync.Once
)

func getIOBluetoothHandsFreeAudioGatewayClass() IOBluetoothHandsFreeAudioGatewayClass {
	_IOBluetoothHandsFreeAudioGatewayClassOnce.Do(func() {
		_IOBluetoothHandsFreeAudioGatewayClass = IOBluetoothHandsFreeAudioGatewayClass{class: objc.GetClass("IOBluetoothHandsFreeAudioGateway")}
	})
	return _IOBluetoothHandsFreeAudioGatewayClass
}

// GetIOBluetoothHandsFreeAudioGatewayClass returns the class object for IOBluetoothHandsFreeAudioGateway.
func GetIOBluetoothHandsFreeAudioGatewayClass() IOBluetoothHandsFreeAudioGatewayClass {
	return getIOBluetoothHandsFreeAudioGatewayClass()
}

type IOBluetoothHandsFreeAudioGatewayClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOBluetoothHandsFreeAudioGatewayClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOBluetoothHandsFreeAudioGatewayClass) Alloc() IOBluetoothHandsFreeAudioGateway {
	rv := objc.Send[IOBluetoothHandsFreeAudioGateway](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// An object that sends data to a connected Bluetooth hands-free phone or
// headset and processes commands from it.
//
// # Overview
//
// This class represents the audio gateway portion of a Bluetooth audio
// profile.
//
// # Showing Status Indicators
//
//   - [IOBluetoothHandsFreeAudioGateway.CreateIndicatorMinMaxCurrentValue]: Sends a request to the Bluetooth device to show or update a status indicator.
//
// # Sending and Receiving Commands
//
//   - [IOBluetoothHandsFreeAudioGateway.SendResponse]: Sends data followed by a success message to a connected Bluetooth hands-free phone or headset.
//   - [IOBluetoothHandsFreeAudioGateway.SendResponseWithOK]: Sends data followed by an optional success message to a connected Bluetooth hands-free phone or headset.
//   - [IOBluetoothHandsFreeAudioGateway.SendOKResponse]: Sends a success message to a connected Bluetooth hands-free phone or headset.
//   - [IOBluetoothHandsFreeAudioGateway.ProcessATCommand]: Processes a command from a connected Bluetooth hands-free phone or headset.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeAudioGateway
type IOBluetoothHandsFreeAudioGateway struct {
	IOBluetoothHandsFree
}

// IOBluetoothHandsFreeAudioGatewayFromID constructs a [IOBluetoothHandsFreeAudioGateway] from an objc.ID.
//
// An object that sends data to a connected Bluetooth hands-free phone or
// headset and processes commands from it.
func IOBluetoothHandsFreeAudioGatewayFromID(id objc.ID) IOBluetoothHandsFreeAudioGateway {
	return IOBluetoothHandsFreeAudioGateway{IOBluetoothHandsFree: IOBluetoothHandsFreeFromID(id)}
}

// NOTE: IOBluetoothHandsFreeAudioGateway adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [IOBluetoothHandsFreeAudioGateway] class.
//
// # Showing Status Indicators
//
//   - [IIOBluetoothHandsFreeAudioGateway.CreateIndicatorMinMaxCurrentValue]: Sends a request to the Bluetooth device to show or update a status indicator.
//
// # Sending and Receiving Commands
//
//   - [IIOBluetoothHandsFreeAudioGateway.SendResponse]: Sends data followed by a success message to a connected Bluetooth hands-free phone or headset.
//   - [IIOBluetoothHandsFreeAudioGateway.SendResponseWithOK]: Sends data followed by an optional success message to a connected Bluetooth hands-free phone or headset.
//   - [IIOBluetoothHandsFreeAudioGateway.SendOKResponse]: Sends a success message to a connected Bluetooth hands-free phone or headset.
//   - [IIOBluetoothHandsFreeAudioGateway.ProcessATCommand]: Processes a command from a connected Bluetooth hands-free phone or headset.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeAudioGateway
type IIOBluetoothHandsFreeAudioGateway interface {
	IIOBluetoothHandsFree

	// Topic: Showing Status Indicators

	// Sends a request to the Bluetooth device to show or update a status indicator.
	CreateIndicatorMinMaxCurrentValue(indicatorName string, minValue int32, maxValue int32, currentValue int32)

	// Topic: Sending and Receiving Commands

	// Sends data followed by a success message to a connected Bluetooth hands-free phone or headset.
	SendResponse(response string)
	// Sends data followed by an optional success message to a connected Bluetooth hands-free phone or headset.
	SendResponseWithOK(response string, withOK bool)
	// Sends a success message to a connected Bluetooth hands-free phone or headset.
	SendOKResponse()
	// Processes a command from a connected Bluetooth hands-free phone or headset.
	ProcessATCommand(atCommand string)
}

// Init initializes the instance.
func (b IOBluetoothHandsFreeAudioGateway) Init() IOBluetoothHandsFreeAudioGateway {
	rv := objc.Send[IOBluetoothHandsFreeAudioGateway](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b IOBluetoothHandsFreeAudioGateway) Autorelease() IOBluetoothHandsFreeAudioGateway {
	rv := objc.Send[IOBluetoothHandsFreeAudioGateway](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOBluetoothHandsFreeAudioGateway creates a new IOBluetoothHandsFreeAudioGateway instance.
func NewIOBluetoothHandsFreeAudioGateway() IOBluetoothHandsFreeAudioGateway {
	class := getIOBluetoothHandsFreeAudioGatewayClass()
	rv := objc.Send[IOBluetoothHandsFreeAudioGateway](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an object that controls a connected Bluetooth hands-free phone or
// headset.
//
// device: A remote Bluetooth phone or headset.
//
// inDelegate: A delegate that conforms to the [IOBluetoothHandsFreeAudioGatewayDelegate]
// protocol.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeAudioGateway/init(device:delegate:)
func NewBluetoothHandsFreeAudioGatewayWithDeviceDelegate(device IIOBluetoothDevice, inDelegate objectivec.IObject) IOBluetoothHandsFreeAudioGateway {
	instance := getIOBluetoothHandsFreeAudioGatewayClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:delegate:"), device, inDelegate)
	return IOBluetoothHandsFreeAudioGatewayFromID(rv)
}

// Sends a request to the Bluetooth device to show or update a status
// indicator.
//
// indicatorName: The name of the indicator. Use one of the following constants:
//
// - [IOBluetoothHandsFreeIndicatorService] -
// [IOBluetoothHandsFreeIndicatorCall] -
// [IOBluetoothHandsFreeIndicatorCallSetup] -
// [IOBluetoothHandsFreeIndicatorCallHeld] -
// [IOBluetoothHandsFreeIndicatorSignal] - [IOBluetoothHandsFreeIndicatorRoam]
// - [IOBluetoothHandsFreeIndicatorBattChg]
//
// minValue: The minimum value for the indicator.
//
// maxValue: The maximum value for the indicator.
//
// currentValue: The current value of the indicator.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeAudioGateway/createIndicator(_:min:max:currentValue:)
//
// [IOBluetoothHandsFreeIndicatorBattChg]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeIndicatorBattChg
// [IOBluetoothHandsFreeIndicatorCallHeld]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeIndicatorCallHeld
// [IOBluetoothHandsFreeIndicatorCallSetup]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeIndicatorCallSetup
// [IOBluetoothHandsFreeIndicatorCall]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeIndicatorCall
// [IOBluetoothHandsFreeIndicatorRoam]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeIndicatorRoam
// [IOBluetoothHandsFreeIndicatorService]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeIndicatorService
// [IOBluetoothHandsFreeIndicatorSignal]: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeIndicatorSignal
func (b IOBluetoothHandsFreeAudioGateway) CreateIndicatorMinMaxCurrentValue(indicatorName string, minValue int32, maxValue int32, currentValue int32) {
	objc.Send[objc.ID](b.ID, objc.Sel("createIndicator:min:max:currentValue:"), objc.String(indicatorName), minValue, maxValue, currentValue)
}

// Sends data followed by a success message to a connected Bluetooth
// hands-free phone or headset.
//
// response: A string containing the data.
//
// # Discussion
//
// Calling this method has the same result as calling `sendResponse(response,
// true)`.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeAudioGateway/sendResponse(_:)
func (b IOBluetoothHandsFreeAudioGateway) SendResponse(response string) {
	objc.Send[objc.ID](b.ID, objc.Sel("sendResponse:"), objc.String(response))
}

// Sends data followed by an optional success message to a connected Bluetooth
// hands-free phone or headset.
//
// response: A string containing the data.
//
// withOK: If `true`, send an [OK] message after sending the response.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeAudioGateway/sendResponse(_:withOK:)
func (b IOBluetoothHandsFreeAudioGateway) SendResponseWithOK(response string, withOK bool) {
	objc.Send[objc.ID](b.ID, objc.Sel("sendResponse:withOK:"), objc.String(response), withOK)
}

// Sends a success message to a connected Bluetooth hands-free phone or
// headset.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeAudioGateway/sendOKResponse()
func (b IOBluetoothHandsFreeAudioGateway) SendOKResponse() {
	objc.Send[objc.ID](b.ID, objc.Sel("sendOKResponse"))
}

// Processes a command from a connected Bluetooth hands-free phone or headset.
//
// atCommand: A string containing the AT command sent from the hands-free Bluetooth
// device.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeAudioGateway/process(atCommand:)
func (b IOBluetoothHandsFreeAudioGateway) ProcessATCommand(atCommand string) {
	objc.Send[objc.ID](b.ID, objc.Sel("processATCommand:"), objc.String(atCommand))
}
