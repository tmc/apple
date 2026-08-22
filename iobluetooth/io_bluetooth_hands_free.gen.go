// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOBluetoothHandsFree] class.
var (
	_IOBluetoothHandsFreeClass     IOBluetoothHandsFreeClass
	_IOBluetoothHandsFreeClassOnce sync.Once
)

func getIOBluetoothHandsFreeClass() IOBluetoothHandsFreeClass {
	_IOBluetoothHandsFreeClassOnce.Do(func() {
		_IOBluetoothHandsFreeClass = IOBluetoothHandsFreeClass{class: objc.GetClass("IOBluetoothHandsFree")}
	})
	return _IOBluetoothHandsFreeClass
}

// GetIOBluetoothHandsFreeClass returns the class object for IOBluetoothHandsFree.
func GetIOBluetoothHandsFreeClass() IOBluetoothHandsFreeClass {
	return getIOBluetoothHandsFreeClass()
}

type IOBluetoothHandsFreeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOBluetoothHandsFreeClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOBluetoothHandsFreeClass) Alloc() IOBluetoothHandsFree {
	rv := objc.Send[IOBluetoothHandsFree](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// Hands free profile class.
//
// # Overview
//
// Superclass of IOBluetoothHandsFreeDevice and
// IOBluetoothHandsFreeAudioGateway classes. Contains the common code used to
// support the bluetoooth hands free profile.
//
// # Initializers
//
//   - [IOBluetoothHandsFree.InitWithDeviceDelegate]: Create a new IOBluetoothHandsFree object
//
// # Instance Properties
//
//   - [IOBluetoothHandsFree.Delegate]: Return the delegate
//   - [IOBluetoothHandsFree.SetDelegate]
//   - [IOBluetoothHandsFree.Device]: Return the IOBluetoothDevice.
//   - [IOBluetoothHandsFree.DeviceCallHoldModes]: Return the device’s supported call hold modes.
//   - [IOBluetoothHandsFree.DeviceSupportedFeatures]: Return the device’s supported features.
//   - [IOBluetoothHandsFree.DeviceSupportedSMSServices]: Return the device’s supported SMS services.
//   - [IOBluetoothHandsFree.InputVolume]: Return the input volume
//   - [IOBluetoothHandsFree.SetInputVolume]
//   - [IOBluetoothHandsFree.IsConnected]
//   - [IOBluetoothHandsFree.IsInputMuted]: Return the input mute state.
//   - [IOBluetoothHandsFree.SetInputMuted]
//   - [IOBluetoothHandsFree.IsOutputMuted]: Return the output mute state.
//   - [IOBluetoothHandsFree.SetOutputMuted]
//   - [IOBluetoothHandsFree.IsSMSEnabled]: Return YES if the device has SMS enabled.
//   - [IOBluetoothHandsFree.OutputVolume]: Return the output volume
//   - [IOBluetoothHandsFree.SetOutputVolume]
//   - [IOBluetoothHandsFree.SMSMode]: Return the device’s SMS mode.
//   - [IOBluetoothHandsFree.SupportedFeatures]: Set the supported features
//   - [IOBluetoothHandsFree.SetSupportedFeatures]
//
// # Instance Methods
//
//   - [IOBluetoothHandsFree.Connect]: Connect to the device
//   - [IOBluetoothHandsFree.ConnectSCO]: Open a SCO connection with the device
//   - [IOBluetoothHandsFree.Disconnect]: Disconnect from the device
//   - [IOBluetoothHandsFree.DisconnectSCO]: Disconnect the SCO connection with the device
//   - [IOBluetoothHandsFree.Indicator]: Return an indicator’s value
//   - [IOBluetoothHandsFree.IsSCOConnected]: Determine if there is a SCO connection to the device
//   - [IOBluetoothHandsFree.SetIndicatorValue]: Set an indicator’s value
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree
type IOBluetoothHandsFree struct {
	objectivec.Object
}

// IOBluetoothHandsFreeFromID constructs a [IOBluetoothHandsFree] from an objc.ID.
//
// Hands free profile class.
func IOBluetoothHandsFreeFromID(id objc.ID) IOBluetoothHandsFree {
	return IOBluetoothHandsFree{objectivec.Object{ID: id}}
}

// NOTE: IOBluetoothHandsFree adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [IOBluetoothHandsFree] class.
//
// # Initializers
//
//   - [IIOBluetoothHandsFree.InitWithDeviceDelegate]: Create a new IOBluetoothHandsFree object
//
// # Instance Properties
//
//   - [IIOBluetoothHandsFree.Delegate]: Return the delegate
//   - [IIOBluetoothHandsFree.SetDelegate]
//   - [IIOBluetoothHandsFree.Device]: Return the IOBluetoothDevice.
//   - [IIOBluetoothHandsFree.DeviceCallHoldModes]: Return the device’s supported call hold modes.
//   - [IIOBluetoothHandsFree.DeviceSupportedFeatures]: Return the device’s supported features.
//   - [IIOBluetoothHandsFree.DeviceSupportedSMSServices]: Return the device’s supported SMS services.
//   - [IIOBluetoothHandsFree.InputVolume]: Return the input volume
//   - [IIOBluetoothHandsFree.SetInputVolume]
//   - [IIOBluetoothHandsFree.IsConnected]
//   - [IIOBluetoothHandsFree.IsInputMuted]: Return the input mute state.
//   - [IIOBluetoothHandsFree.SetInputMuted]
//   - [IIOBluetoothHandsFree.IsOutputMuted]: Return the output mute state.
//   - [IIOBluetoothHandsFree.SetOutputMuted]
//   - [IIOBluetoothHandsFree.IsSMSEnabled]: Return YES if the device has SMS enabled.
//   - [IIOBluetoothHandsFree.OutputVolume]: Return the output volume
//   - [IIOBluetoothHandsFree.SetOutputVolume]
//   - [IIOBluetoothHandsFree.SMSMode]: Return the device’s SMS mode.
//   - [IIOBluetoothHandsFree.SupportedFeatures]: Set the supported features
//   - [IIOBluetoothHandsFree.SetSupportedFeatures]
//
// # Instance Methods
//
//   - [IIOBluetoothHandsFree.Connect]: Connect to the device
//   - [IIOBluetoothHandsFree.ConnectSCO]: Open a SCO connection with the device
//   - [IIOBluetoothHandsFree.Disconnect]: Disconnect from the device
//   - [IIOBluetoothHandsFree.DisconnectSCO]: Disconnect the SCO connection with the device
//   - [IIOBluetoothHandsFree.Indicator]: Return an indicator’s value
//   - [IIOBluetoothHandsFree.IsSCOConnected]: Determine if there is a SCO connection to the device
//   - [IIOBluetoothHandsFree.SetIndicatorValue]: Set an indicator’s value
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree
type IIOBluetoothHandsFree interface {
	objectivec.IObject

	// Topic: Initializers

	// Create a new IOBluetoothHandsFree object
	InitWithDeviceDelegate(device IIOBluetoothDevice, inDelegate IOBluetoothHandsFreeDelegate) IOBluetoothHandsFree

	// Topic: Instance Properties

	// Return the delegate
	Delegate() IOBluetoothHandsFreeDelegate
	SetDelegate(value IOBluetoothHandsFreeDelegate)
	// Return the IOBluetoothDevice.
	Device() IIOBluetoothDevice
	// Return the device’s supported call hold modes.
	DeviceCallHoldModes() uint32
	// Return the device’s supported features.
	DeviceSupportedFeatures() uint32
	// Return the device’s supported SMS services.
	DeviceSupportedSMSServices() uint32
	// Return the input volume
	InputVolume() float32
	SetInputVolume(value float32)
	IsConnected() bool
	// Return the input mute state.
	IsInputMuted() bool
	SetInputMuted(value bool)
	// Return the output mute state.
	IsOutputMuted() bool
	SetOutputMuted(value bool)
	// Return YES if the device has SMS enabled.
	IsSMSEnabled() bool
	// Return the output volume
	OutputVolume() float32
	SetOutputVolume(value float32)
	// Return the device’s SMS mode.
	SMSMode() IOBluetoothSMSMode
	// Set the supported features
	SupportedFeatures() uint32
	SetSupportedFeatures(value uint32)

	// Topic: Instance Methods

	// Connect to the device
	Connect()
	// Open a SCO connection with the device
	ConnectSCO()
	// Disconnect from the device
	Disconnect()
	// Disconnect the SCO connection with the device
	DisconnectSCO()
	// Return an indicator’s value
	Indicator(indicatorName string) int32
	// Determine if there is a SCO connection to the device
	IsSCOConnected() bool
	// Set an indicator’s value
	SetIndicatorValue(indicatorName string, indicatorValue int32)
}

// Init initializes the instance.
func (b IOBluetoothHandsFree) Init() IOBluetoothHandsFree {
	rv := objc.Send[IOBluetoothHandsFree](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b IOBluetoothHandsFree) Autorelease() IOBluetoothHandsFree {
	rv := objc.Send[IOBluetoothHandsFree](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOBluetoothHandsFree creates a new IOBluetoothHandsFree instance.
func NewIOBluetoothHandsFree() IOBluetoothHandsFree {
	class := getIOBluetoothHandsFreeClass()
	rv := objc.Send[IOBluetoothHandsFree](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Create a new IOBluetoothHandsFree object
//
// device: An IOBluetoothDevice
//
// inDelegate: An object to act as delegate that implements the
// IOBluetoothHandsFreeDelegate protocol.
//
// # Return Value
//
// A newly created IOBluetoothHandsFreeAudioGateway object on success, nil on
// failure
//
// # Discussion
//
// This method should be called on a subclass (IOBluetoothHandsFreeDevice or
// IOBluetoothHandsFreeAudioGateway) to get full functionality.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/init(device:delegate:)
func NewBluetoothHandsFreeWithDeviceDelegate(device IIOBluetoothDevice, inDelegate IOBluetoothHandsFreeDelegate) IOBluetoothHandsFree {
	instance := getIOBluetoothHandsFreeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:delegate:"), device, inDelegate)
	return IOBluetoothHandsFreeFromID(rv)
}

// Create a new IOBluetoothHandsFree object
//
// device: An IOBluetoothDevice
//
// inDelegate: An object to act as delegate that implements the
// IOBluetoothHandsFreeDelegate protocol.
//
// # Return Value
//
// A newly created IOBluetoothHandsFreeAudioGateway object on success, nil on
// failure
//
// # Discussion
//
// This method should be called on a subclass (IOBluetoothHandsFreeDevice or
// IOBluetoothHandsFreeAudioGateway) to get full functionality.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/init(device:delegate:)
func (b IOBluetoothHandsFree) InitWithDeviceDelegate(device IIOBluetoothDevice, inDelegate IOBluetoothHandsFreeDelegate) IOBluetoothHandsFree {
	rv := objc.Send[IOBluetoothHandsFree](b.ID, objc.Sel("initWithDevice:delegate:"), device, inDelegate)
	return rv
}

// Connect to the device
//
// # Discussion
//
// Connects to the device and sets up a service level connection (RFCOMM
// channel). Delegate methods will be called once the connection is complete
// or a failure occurs.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/connect()
func (b IOBluetoothHandsFree) Connect() {
	objc.Send[objc.ID](b.ID, objc.Sel("connect"))
}

// Open a SCO connection with the device
//
// # Discussion
//
// Opens a SCO connection with the device. The device must already have a
// service level connection or this will return immediately. Delegate methods
// will be called once the connection is complete of a failure occurs.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/connectSCO()
func (b IOBluetoothHandsFree) ConnectSCO() {
	objc.Send[objc.ID](b.ID, objc.Sel("connectSCO"))
}

// Disconnect from the device
//
// # Discussion
//
// Disconnects from the device, closes the SCO and service level connection if
// they are connected. Delegate methods will be called once the disconnection
// is complete.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/disconnect()
func (b IOBluetoothHandsFree) Disconnect() {
	objc.Send[objc.ID](b.ID, objc.Sel("disconnect"))
}

// Disconnect the SCO connection with the device
//
// # Discussion
//
// Disconnects the SCO connection with the device (if one exists). Delegate
// methods will be called once the disconnection is complete.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/disconnectSCO()
func (b IOBluetoothHandsFree) DisconnectSCO() {
	objc.Send[objc.ID](b.ID, objc.Sel("disconnectSCO"))
}

// Return an indicator’s value
//
// indicatorName: See “Hands free indicator constants,” for standard indicator names.
//
// # Discussion
//
// Returns an indicator’s value.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/indicator(_:)
func (b IOBluetoothHandsFree) Indicator(indicatorName string) int32 {
	rv := objc.Send[int32](b.ID, objc.Sel("indicator:"), objc.String(indicatorName))
	return rv
}

// Determine if there is a SCO connection to the device
//
// # Return Value
//
// YES if there is a SCO connection to the device; otherwise, NO.
//
// # Discussion
//
// Determines if there is a SCO connection to the device.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/isSCOConnected()
func (b IOBluetoothHandsFree) IsSCOConnected() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("isSCOConnected"))
	return rv
}

// Set an indicator’s value
//
// indicatorName: See “Hands free indicator constants,” for standard indicator names.
//
// indicatorValue: Will set the indicator value as long as it is within the min and max values
// allowed.
//
// # Discussion
//
// Sets an indicator’s value.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/setIndicator(_:value:)
func (b IOBluetoothHandsFree) SetIndicatorValue(indicatorName string, indicatorValue int32) {
	objc.Send[objc.ID](b.ID, objc.Sel("setIndicator:value:"), objc.String(indicatorName), indicatorValue)
}

// Return the delegate
//
// # Return Value
//
// The delegate for the hands free object or nil if it doesn’t have a
// delegate.
//
// # Discussion
//
// Returns the hands free object’s delegate.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/delegate
func (b IOBluetoothHandsFree) Delegate() IOBluetoothHandsFreeDelegate {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("delegate"))
	return IOBluetoothHandsFreeDelegateObjectFromID(rv)
}
func (b IOBluetoothHandsFree) SetDelegate(value IOBluetoothHandsFreeDelegate) {
	objc.Send[struct{}](b.ID, objc.Sel("setDelegate:"), value)
}

// Return the IOBluetoothDevice.
//
// # Return Value
//
// # The IOBluetoothDevice object
//
// # Discussion
//
// Returns the IOBluetoothDevice to connect with.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/device
func (b IOBluetoothHandsFree) Device() IIOBluetoothDevice {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("device"))
	return IOBluetoothDeviceFromID(objc.ID(rv))
}

// Return the device’s supported call hold modes.
//
// # Return Value
//
// # The SMS services supported
//
// # Discussion
//
// Returns the device’s supported call hold modes bitmap. The values are
// described in “IOBluetoothHandsFreeCallHoldModes.”
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/deviceCallHoldModes
func (b IOBluetoothHandsFree) DeviceCallHoldModes() uint32 {
	rv := objc.Send[uint32](b.ID, objc.Sel("deviceCallHoldModes"))
	return rv
}

// Return the device’s supported features.
//
// # Return Value
//
// # The device features bitmap
//
// # Discussion
//
// Returns the device’s supported features bitmap. The values are described
// in “IOBluetoothHandsFreeDeviceFeatures and
// IOBluetoothHandsFreeAudioGatewayFeatures.”
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/deviceSupportedFeatures
func (b IOBluetoothHandsFree) DeviceSupportedFeatures() uint32 {
	rv := objc.Send[uint32](b.ID, objc.Sel("deviceSupportedFeatures"))
	return rv
}

// Return the device’s supported SMS services.
//
// # Return Value
//
// # The SMS services supported
//
// # Discussion
//
// Returns the device’s supported SMS services bitmap. The values are
// described in “IOBluetoothHandsFreeSMSSupport.”
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/deviceSupportedSMSServices
func (b IOBluetoothHandsFree) DeviceSupportedSMSServices() uint32 {
	rv := objc.Send[uint32](b.ID, objc.Sel("deviceSupportedSMSServices"))
	return rv
}

// Return the input volume
//
// # Return Value
//
// # The input volume
//
// # Discussion
//
// Returns the input volume between 0 and 1. 0 is the same as mute.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/inputVolume
func (b IOBluetoothHandsFree) InputVolume() float32 {
	rv := objc.Send[float32](b.ID, objc.Sel("inputVolume"))
	return rv
}
func (b IOBluetoothHandsFree) SetInputVolume(value float32) {
	objc.Send[struct{}](b.ID, objc.Sel("setInputVolume:"), value)
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/isConnected
func (b IOBluetoothHandsFree) IsConnected() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("isConnected"))
	return rv
}

// Return the input mute state.
//
// # Discussion
//
// Returns the inputs mute state.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/isInputMuted
func (b IOBluetoothHandsFree) IsInputMuted() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("isInputMuted"))
	return rv
}
func (b IOBluetoothHandsFree) SetInputMuted(value bool) {
	objc.Send[struct{}](b.ID, objc.Sel("setInputMuted:"), value)
}

// Return the output mute state.
//
// # Discussion
//
// Returns the outputs mute state.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/isOutputMuted
func (b IOBluetoothHandsFree) IsOutputMuted() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("isOutputMuted"))
	return rv
}
func (b IOBluetoothHandsFree) SetOutputMuted(value bool) {
	objc.Send[struct{}](b.ID, objc.Sel("setOutputMuted:"), value)
}

// Return YES if the device has SMS enabled.
//
// # Discussion
//
// Returns YES if the device has SMS enabled (by responding to a CMGF
// command). NO if the device has not set an SMS mode or doesn’t support
// SMS.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/isSMSEnabled
func (b IOBluetoothHandsFree) IsSMSEnabled() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("isSMSEnabled"))
	return rv
}

// Return the output volume
//
// # Return Value
//
// # The output volume
//
// # Discussion
//
// Returns the output volume between 0 and 1. 0 is the same as mute.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/outputVolume
func (b IOBluetoothHandsFree) OutputVolume() float32 {
	rv := objc.Send[float32](b.ID, objc.Sel("outputVolume"))
	return rv
}
func (b IOBluetoothHandsFree) SetOutputVolume(value float32) {
	objc.Send[struct{}](b.ID, objc.Sel("setOutputVolume:"), value)
}

// Return the device’s SMS mode.
//
// # Return Value
//
// # The SMS mode
//
// # Discussion
//
// Returns the device’s SMS mode. The values are described in
// “IOBluetoothSMSMode.”
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/smsMode
func (b IOBluetoothHandsFree) SMSMode() IOBluetoothSMSMode {
	rv := objc.Send[IOBluetoothSMSMode](b.ID, objc.Sel("SMSMode"))
	return IOBluetoothSMSMode(rv)
}

// Set the supported features
//
// # Discussion
//
// Sets the supported features bitmap. The values are described in
// “IOBluetoothHandsFreeDeviceFeatures and
// IOBluetoothHandsFreeAudioGatewayFeatures.”
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFree/supportedFeatures
func (b IOBluetoothHandsFree) SupportedFeatures() uint32 {
	rv := objc.Send[uint32](b.ID, objc.Sel("supportedFeatures"))
	return rv
}
func (b IOBluetoothHandsFree) SetSupportedFeatures(value uint32) {
	objc.Send[struct{}](b.ID, objc.Sel("setSupportedFeatures:"), value)
}
