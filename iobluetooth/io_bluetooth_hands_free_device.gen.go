// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOBluetoothHandsFreeDevice] class.
var (
	_IOBluetoothHandsFreeDeviceClass     IOBluetoothHandsFreeDeviceClass
	_IOBluetoothHandsFreeDeviceClassOnce sync.Once
)

func getIOBluetoothHandsFreeDeviceClass() IOBluetoothHandsFreeDeviceClass {
	_IOBluetoothHandsFreeDeviceClassOnce.Do(func() {
		_IOBluetoothHandsFreeDeviceClass = IOBluetoothHandsFreeDeviceClass{class: objc.GetClass("IOBluetoothHandsFreeDevice")}
	})
	return _IOBluetoothHandsFreeDeviceClass
}

// GetIOBluetoothHandsFreeDeviceClass returns the class object for IOBluetoothHandsFreeDevice.
func GetIOBluetoothHandsFreeDeviceClass() IOBluetoothHandsFreeDeviceClass {
	return getIOBluetoothHandsFreeDeviceClass()
}

type IOBluetoothHandsFreeDeviceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOBluetoothHandsFreeDeviceClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOBluetoothHandsFreeDeviceClass) Alloc() IOBluetoothHandsFreeDevice {
	rv := objc.Send[IOBluetoothHandsFreeDevice](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// An object you use to manage phone calls on a connected Bluetooth hands-free
// phone or headset.
//
// # Accepting Calls
//
//   - [IOBluetoothHandsFreeDevice.AcceptCall]: Accepts an incoming call.
//   - [IOBluetoothHandsFreeDevice.AcceptCallOnPhone]: Accepts an incoming call and transfers the audio to the managed hands-free phone or headset.
//   - [IOBluetoothHandsFreeDevice.CallTransfer]: Ends all calls that are active or on hold, and accepts any waiting calls.
//
// # Dialing Calls
//
//   - [IOBluetoothHandsFreeDevice.DialNumber]: Calls the phone number on a hands-free phone or headset.
//   - [IOBluetoothHandsFreeDevice.MemoryDial]: Calls the phone number stored in a speed dial or memory slot of the hands-free phone or headset.
//   - [IOBluetoothHandsFreeDevice.Redial]: Calls the number stored on the hands-free phone or headset again.
//
// # Holding Calls
//
//   - [IOBluetoothHandsFreeDevice.HoldCall]: Places all active calls on hold and accepts a held or waiting call.
//   - [IOBluetoothHandsFreeDevice.AddHeldCall]: Adds held calls to the current conversation.
//   - [IOBluetoothHandsFreeDevice.PlaceAllOthersOnHold]: Places all calls except the call with the specified index on hold.
//
// # Ending Calls
//
//   - [IOBluetoothHandsFreeDevice.EndCall]: Ends the current call or refuses an incoming call.
//   - [IOBluetoothHandsFreeDevice.ReleaseCall]: Ends the call with the specified index.
//   - [IOBluetoothHandsFreeDevice.ReleaseActiveCalls]: Ends all active calls and accepts a held or waiting call.
//   - [IOBluetoothHandsFreeDevice.ReleaseHeldCalls]: Ends all calls that are on hold or returns a busy signal for a waiting call.
//
// # Sending Messages and Commands
//
//   - [IOBluetoothHandsFreeDevice.SendSMSMessage]: Sends a text message to a phone number.
//   - [IOBluetoothHandsFreeDevice.SendDTMF]: Sends the tone associated with a phone key to the hands-free Bluetooth device.
//   - [IOBluetoothHandsFreeDevice.SendATCommand]: Sends an AT command to the Bluetooth audio gateway.
//   - [IOBluetoothHandsFreeDevice.SendATCommandTimeoutSelectorTarget]: Send an AT command to the Bluetooth audio gateway and performs a selector on completion or timeout.
//
// # Requesting Status Information
//
//   - [IOBluetoothHandsFreeDevice.SubscriberNumber]: Requests that the Bluetooth audio gateway send the subscriber number to the delegate.
//   - [IOBluetoothHandsFreeDevice.CurrentCallList]: Requests that the Bluetooth audio gateway send the delegate a list of calls that are active, on hold, or being set up.
//
// # Transferring Audio
//
//   - [IOBluetoothHandsFreeDevice.TransferAudioToComputer]: Moves the audio for current and future calls to a Mac.
//   - [IOBluetoothHandsFreeDevice.TransferAudioToPhone]: Moves the audio for current or future calls to a phone.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice
type IOBluetoothHandsFreeDevice struct {
	IOBluetoothHandsFree
}

// IOBluetoothHandsFreeDeviceFromID constructs a [IOBluetoothHandsFreeDevice] from an objc.ID.
//
// An object you use to manage phone calls on a connected Bluetooth hands-free
// phone or headset.
func IOBluetoothHandsFreeDeviceFromID(id objc.ID) IOBluetoothHandsFreeDevice {
	return IOBluetoothHandsFreeDevice{IOBluetoothHandsFree: IOBluetoothHandsFreeFromID(id)}
}

// NOTE: IOBluetoothHandsFreeDevice adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [IOBluetoothHandsFreeDevice] class.
//
// # Accepting Calls
//
//   - [IIOBluetoothHandsFreeDevice.AcceptCall]: Accepts an incoming call.
//   - [IIOBluetoothHandsFreeDevice.AcceptCallOnPhone]: Accepts an incoming call and transfers the audio to the managed hands-free phone or headset.
//   - [IIOBluetoothHandsFreeDevice.CallTransfer]: Ends all calls that are active or on hold, and accepts any waiting calls.
//
// # Dialing Calls
//
//   - [IIOBluetoothHandsFreeDevice.DialNumber]: Calls the phone number on a hands-free phone or headset.
//   - [IIOBluetoothHandsFreeDevice.MemoryDial]: Calls the phone number stored in a speed dial or memory slot of the hands-free phone or headset.
//   - [IIOBluetoothHandsFreeDevice.Redial]: Calls the number stored on the hands-free phone or headset again.
//
// # Holding Calls
//
//   - [IIOBluetoothHandsFreeDevice.HoldCall]: Places all active calls on hold and accepts a held or waiting call.
//   - [IIOBluetoothHandsFreeDevice.AddHeldCall]: Adds held calls to the current conversation.
//   - [IIOBluetoothHandsFreeDevice.PlaceAllOthersOnHold]: Places all calls except the call with the specified index on hold.
//
// # Ending Calls
//
//   - [IIOBluetoothHandsFreeDevice.EndCall]: Ends the current call or refuses an incoming call.
//   - [IIOBluetoothHandsFreeDevice.ReleaseCall]: Ends the call with the specified index.
//   - [IIOBluetoothHandsFreeDevice.ReleaseActiveCalls]: Ends all active calls and accepts a held or waiting call.
//   - [IIOBluetoothHandsFreeDevice.ReleaseHeldCalls]: Ends all calls that are on hold or returns a busy signal for a waiting call.
//
// # Sending Messages and Commands
//
//   - [IIOBluetoothHandsFreeDevice.SendSMSMessage]: Sends a text message to a phone number.
//   - [IIOBluetoothHandsFreeDevice.SendDTMF]: Sends the tone associated with a phone key to the hands-free Bluetooth device.
//   - [IIOBluetoothHandsFreeDevice.SendATCommand]: Sends an AT command to the Bluetooth audio gateway.
//   - [IIOBluetoothHandsFreeDevice.SendATCommandTimeoutSelectorTarget]: Send an AT command to the Bluetooth audio gateway and performs a selector on completion or timeout.
//
// # Requesting Status Information
//
//   - [IIOBluetoothHandsFreeDevice.SubscriberNumber]: Requests that the Bluetooth audio gateway send the subscriber number to the delegate.
//   - [IIOBluetoothHandsFreeDevice.CurrentCallList]: Requests that the Bluetooth audio gateway send the delegate a list of calls that are active, on hold, or being set up.
//
// # Transferring Audio
//
//   - [IIOBluetoothHandsFreeDevice.TransferAudioToComputer]: Moves the audio for current and future calls to a Mac.
//   - [IIOBluetoothHandsFreeDevice.TransferAudioToPhone]: Moves the audio for current or future calls to a phone.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice
type IIOBluetoothHandsFreeDevice interface {
	IIOBluetoothHandsFree

	// Topic: Accepting Calls

	// Accepts an incoming call.
	AcceptCall()
	// Accepts an incoming call and transfers the audio to the managed hands-free phone or headset.
	AcceptCallOnPhone()
	// Ends all calls that are active or on hold, and accepts any waiting calls.
	CallTransfer()

	// Topic: Dialing Calls

	// Calls the phone number on a hands-free phone or headset.
	DialNumber(aNumber string)
	// Calls the phone number stored in a speed dial or memory slot of the hands-free phone or headset.
	MemoryDial(memoryLocation int32)
	// Calls the number stored on the hands-free phone or headset again.
	Redial()

	// Topic: Holding Calls

	// Places all active calls on hold and accepts a held or waiting call.
	HoldCall()
	// Adds held calls to the current conversation.
	AddHeldCall()
	// Places all calls except the call with the specified index on hold.
	PlaceAllOthersOnHold(index int32)

	// Topic: Ending Calls

	// Ends the current call or refuses an incoming call.
	EndCall()
	// Ends the call with the specified index.
	ReleaseCall(index int32)
	// Ends all active calls and accepts a held or waiting call.
	ReleaseActiveCalls()
	// Ends all calls that are on hold or returns a busy signal for a waiting call.
	ReleaseHeldCalls()

	// Topic: Sending Messages and Commands

	// Sends a text message to a phone number.
	SendSMSMessage(aNumber string, aMessage string)
	// Sends the tone associated with a phone key to the hands-free Bluetooth device.
	SendDTMF(character string)
	// Sends an AT command to the Bluetooth audio gateway.
	SendATCommand(atCommand string)
	// Send an AT command to the Bluetooth audio gateway and performs a selector on completion or timeout.
	SendATCommandTimeoutSelectorTarget(atCommand string, timeout float32, selector objc.SEL, target objectivec.IObject)

	// Topic: Requesting Status Information

	// Requests that the Bluetooth audio gateway send the subscriber number to the delegate.
	SubscriberNumber()
	// Requests that the Bluetooth audio gateway send the delegate a list of calls that are active, on hold, or being set up.
	CurrentCallList()

	// Topic: Transferring Audio

	// Moves the audio for current and future calls to a Mac.
	TransferAudioToComputer()
	// Moves the audio for current or future calls to a phone.
	TransferAudioToPhone()
}

// Init initializes the instance.
func (b IOBluetoothHandsFreeDevice) Init() IOBluetoothHandsFreeDevice {
	rv := objc.Send[IOBluetoothHandsFreeDevice](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b IOBluetoothHandsFreeDevice) Autorelease() IOBluetoothHandsFreeDevice {
	rv := objc.Send[IOBluetoothHandsFreeDevice](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOBluetoothHandsFreeDevice creates a new IOBluetoothHandsFreeDevice instance.
func NewIOBluetoothHandsFreeDevice() IOBluetoothHandsFreeDevice {
	class := getIOBluetoothHandsFreeDeviceClass()
	rv := objc.Send[IOBluetoothHandsFreeDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an object to manage phone calls on a hands-free Bluetooth device.
//
// device: A Bluetooth device.
//
// delegate: A delegate that conforms to the [IOBluetoothHandsFreeDeviceDelegate]
// protocol.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/init(device:delegate:)
func NewBluetoothHandsFreeDeviceWithDeviceDelegate(device IIOBluetoothDevice, delegate objectivec.IObject) IOBluetoothHandsFreeDevice {
	instance := getIOBluetoothHandsFreeDeviceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:delegate:"), device, delegate)
	return IOBluetoothHandsFreeDeviceFromID(rv)
}

// Accepts an incoming call.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/acceptCall()
func (b IOBluetoothHandsFreeDevice) AcceptCall() {
	objc.Send[objc.ID](b.ID, objc.Sel("acceptCall"))
}

// Accepts an incoming call and transfers the audio to the managed hands-free
// phone or headset.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/acceptCallOnPhone()
func (b IOBluetoothHandsFreeDevice) AcceptCallOnPhone() {
	objc.Send[objc.ID](b.ID, objc.Sel("acceptCallOnPhone"))
}

// Ends all calls that are active or on hold, and accepts any waiting calls.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/callTransfer()
func (b IOBluetoothHandsFreeDevice) CallTransfer() {
	objc.Send[objc.ID](b.ID, objc.Sel("callTransfer"))
}

// Calls the phone number on a hands-free phone or headset.
//
// aNumber: A string containing a phone number.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/dialNumber(_:)
func (b IOBluetoothHandsFreeDevice) DialNumber(aNumber string) {
	objc.Send[objc.ID](b.ID, objc.Sel("dialNumber:"), objc.String(aNumber))
}

// Calls the phone number stored in a speed dial or memory slot of the
// hands-free phone or headset.
//
// memoryLocation: The speed dial or other memory index of a phone number.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/memoryDial(_:)
func (b IOBluetoothHandsFreeDevice) MemoryDial(memoryLocation int32) {
	objc.Send[objc.ID](b.ID, objc.Sel("memoryDial:"), memoryLocation)
}

// Calls the number stored on the hands-free phone or headset again.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/redial()
func (b IOBluetoothHandsFreeDevice) Redial() {
	objc.Send[objc.ID](b.ID, objc.Sel("redial"))
}

// Places all active calls on hold and accepts a held or waiting call.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/holdCall()
func (b IOBluetoothHandsFreeDevice) HoldCall() {
	objc.Send[objc.ID](b.ID, objc.Sel("holdCall"))
}

// Adds held calls to the current conversation.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/addHeldCall()
func (b IOBluetoothHandsFreeDevice) AddHeldCall() {
	objc.Send[objc.ID](b.ID, objc.Sel("addHeldCall"))
}

// Places all calls except the call with the specified index on hold.
//
// index: The index of the call that remains active.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/placeAllOthers(onHold:)
func (b IOBluetoothHandsFreeDevice) PlaceAllOthersOnHold(index int32) {
	objc.Send[objc.ID](b.ID, objc.Sel("placeAllOthersOnHold:"), index)
}

// Ends the current call or refuses an incoming call.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/endCall()
func (b IOBluetoothHandsFreeDevice) EndCall() {
	objc.Send[objc.ID](b.ID, objc.Sel("endCall"))
}

// Ends the call with the specified index.
//
// index: The index of the call to end.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/releaseCall(_:)
func (b IOBluetoothHandsFreeDevice) ReleaseCall(index int32) {
	objc.Send[objc.ID](b.ID, objc.Sel("releaseCall:"), index)
}

// Ends all active calls and accepts a held or waiting call.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/releaseActiveCalls()
func (b IOBluetoothHandsFreeDevice) ReleaseActiveCalls() {
	objc.Send[objc.ID](b.ID, objc.Sel("releaseActiveCalls"))
}

// Ends all calls that are on hold or returns a busy signal for a waiting
// call.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/releaseHeldCalls()
func (b IOBluetoothHandsFreeDevice) ReleaseHeldCalls() {
	objc.Send[objc.ID](b.ID, objc.Sel("releaseHeldCalls"))
}

// Sends a text message to a phone number.
//
// aNumber: The phone number to send the message to.
//
// aMessage: A string containing a message. The message must be no longer than 160
// characters.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/sendSMS(_:message:)
func (b IOBluetoothHandsFreeDevice) SendSMSMessage(aNumber string, aMessage string) {
	objc.Send[objc.ID](b.ID, objc.Sel("sendSMS:message:"), objc.String(aNumber), objc.String(aMessage))
}

// Sends the tone associated with a phone key to the hands-free Bluetooth
// device.
//
// character: The phone keypad character for the tone. The character must be one of the
// following:
//
// - `0-9`
// - `#`
// - `*`
// - `A-D`
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/sendDTMF(_:)
func (b IOBluetoothHandsFreeDevice) SendDTMF(character string) {
	objc.Send[objc.ID](b.ID, objc.Sel("sendDTMF:"), objc.String(character))
}

// Sends an AT command to the Bluetooth audio gateway.
//
// atCommand: A string containing the AT command.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/send(atCommand:)
func (b IOBluetoothHandsFreeDevice) SendATCommand(atCommand string) {
	objc.Send[objc.ID](b.ID, objc.Sel("sendATCommand:"), objc.String(atCommand))
}

// Send an AT command to the Bluetooth audio gateway and performs a selector
// on completion or timeout.
//
// atCommand: A string containing the AT command.
//
// timeout: The number of seconds until the message times out.
//
// selector: The function to call on completion or timeout.
//
// target: The target object for the completion call.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/send(atCommand:timeout:selector:target:)
func (b IOBluetoothHandsFreeDevice) SendATCommandTimeoutSelectorTarget(atCommand string, timeout float32, selector objc.SEL, target objectivec.IObject) {
	objc.Send[objc.ID](b.ID, objc.Sel("sendATCommand:timeout:selector:target:"), objc.String(atCommand), timeout, selector, target)
}

// Requests that the Bluetooth audio gateway send the subscriber number to the
// delegate.
//
// # Discussion
//
// The subscriber number is sent to the [HandsFreeSubscriberNumber] function
// of the delegate.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/subscriberNumber()
func (b IOBluetoothHandsFreeDevice) SubscriberNumber() {
	objc.Send[objc.ID](b.ID, objc.Sel("subscriberNumber"))
}

// Requests that the Bluetooth audio gateway send the delegate a list of calls
// that are active, on hold, or being set up.
//
// # Discussion
//
// The [HandsFreeCurrentCall] function of the delegate is called once for each
// current call.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/currentCallList()
func (b IOBluetoothHandsFreeDevice) CurrentCallList() {
	objc.Send[objc.ID](b.ID, objc.Sel("currentCallList"))
}

// Moves the audio for current and future calls to a Mac.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/transferAudioToComputer()
func (b IOBluetoothHandsFreeDevice) TransferAudioToComputer() {
	objc.Send[objc.ID](b.ID, objc.Sel("transferAudioToComputer"))
}

// Moves the audio for current or future calls to a phone.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothHandsFreeDevice/transferAudioToPhone()
func (b IOBluetoothHandsFreeDevice) TransferAudioToPhone() {
	objc.Send[objc.ID](b.ID, objc.Sel("transferAudioToPhone"))
}
