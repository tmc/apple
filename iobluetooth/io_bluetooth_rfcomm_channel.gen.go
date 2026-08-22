// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/kernel"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOBluetoothRFCOMMChannel] class.
var (
	_IOBluetoothRFCOMMChannelClass     IOBluetoothRFCOMMChannelClass
	_IOBluetoothRFCOMMChannelClassOnce sync.Once
)

func getIOBluetoothRFCOMMChannelClass() IOBluetoothRFCOMMChannelClass {
	_IOBluetoothRFCOMMChannelClassOnce.Do(func() {
		_IOBluetoothRFCOMMChannelClass = IOBluetoothRFCOMMChannelClass{class: objc.GetClass("IOBluetoothRFCOMMChannel")}
	})
	return _IOBluetoothRFCOMMChannelClass
}

// GetIOBluetoothRFCOMMChannelClass returns the class object for IOBluetoothRFCOMMChannel.
func GetIOBluetoothRFCOMMChannelClass() IOBluetoothRFCOMMChannelClass {
	return getIOBluetoothRFCOMMChannelClass()
}

type IOBluetoothRFCOMMChannelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOBluetoothRFCOMMChannelClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOBluetoothRFCOMMChannelClass) Alloc() IOBluetoothRFCOMMChannel {
	rv := objc.Send[IOBluetoothRFCOMMChannel](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// An instance of this class represents an RFCOMM channel as defined by the
// Bluetooth SDP spec..
//
// # Overview
//
// An RFCOMM channel object can be obtained by opening an RFCOMM channel in a
// device, or by requesting a notification when a channel is created (this is
// commonly used to provide services).
//
// # Instance Methods
//
//   - [IOBluetoothRFCOMMChannel.CloseChannel]: Close the channel.
//   - [IOBluetoothRFCOMMChannel.Delegate]: Returns the object delegate
//   - [IOBluetoothRFCOMMChannel.GetDevice]: Returns the Bluetooth Device that carries the rfcomm data.
//   - [IOBluetoothRFCOMMChannel.GetChannelID]: Returns the object rfcomm channel ID.
//   - [IOBluetoothRFCOMMChannel.GetMTU]: Returns the channel maximum transfer unit.
//   - [IOBluetoothRFCOMMChannel.GetObjectID]: Returns the IOBluetoothObjectID of the given IOBluetoothRFCOMMChannel.
//   - [IOBluetoothRFCOMMChannel.GetRFCOMMChannelRef]: Returns an IOBluetoothRFCOMMChannelRef representation of the target IOBluetoothRFCOMMChannel object.
//   - [IOBluetoothRFCOMMChannel.IsIncoming]: Returns the direction of the channel. An incoming channel is one that was opened by the remote device.
//   - [IOBluetoothRFCOMMChannel.IsOpen]: Returns the state of the channel.
//   - [IOBluetoothRFCOMMChannel.IsTransmissionPaused]: Returns TRUE if flow control is off.
//   - [IOBluetoothRFCOMMChannel.RegisterForChannelCloseNotificationSelector]: Allows a client to register for a channel close notification.
//   - [IOBluetoothRFCOMMChannel.SendRemoteLineStatus]: Sends an error to the remote side.
//   - [IOBluetoothRFCOMMChannel.SetDelegate]: Allows an object to register itself as a client of the RFCOMM channel.
//   - [IOBluetoothRFCOMMChannel.SetSerialParametersDataBitsParityStopBits]: Changes the parameters of the serial connection.
//   - [IOBluetoothRFCOMMChannel.WriteAsyncLengthRefcon]: Sends a block of data in the channel asynchronously.
//   - [IOBluetoothRFCOMMChannel.WriteSyncLength]: Sends a block of data in the channel synchronously.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel
type IOBluetoothRFCOMMChannel struct {
	IOBluetoothObject
}

// IOBluetoothRFCOMMChannelFromID constructs a [IOBluetoothRFCOMMChannel] from an objc.ID.
//
// An instance of this class represents an RFCOMM channel as defined by the
// Bluetooth SDP spec..
func IOBluetoothRFCOMMChannelFromID(id objc.ID) IOBluetoothRFCOMMChannel {
	return IOBluetoothRFCOMMChannel{IOBluetoothObject: IOBluetoothObjectFromID(id)}
}

// NOTE: IOBluetoothRFCOMMChannel adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [IOBluetoothRFCOMMChannel] class.
//
// # Instance Methods
//
//   - [IIOBluetoothRFCOMMChannel.CloseChannel]: Close the channel.
//   - [IIOBluetoothRFCOMMChannel.Delegate]: Returns the object delegate
//   - [IIOBluetoothRFCOMMChannel.GetDevice]: Returns the Bluetooth Device that carries the rfcomm data.
//   - [IIOBluetoothRFCOMMChannel.GetChannelID]: Returns the object rfcomm channel ID.
//   - [IIOBluetoothRFCOMMChannel.GetMTU]: Returns the channel maximum transfer unit.
//   - [IIOBluetoothRFCOMMChannel.GetObjectID]: Returns the IOBluetoothObjectID of the given IOBluetoothRFCOMMChannel.
//   - [IIOBluetoothRFCOMMChannel.GetRFCOMMChannelRef]: Returns an IOBluetoothRFCOMMChannelRef representation of the target IOBluetoothRFCOMMChannel object.
//   - [IIOBluetoothRFCOMMChannel.IsIncoming]: Returns the direction of the channel. An incoming channel is one that was opened by the remote device.
//   - [IIOBluetoothRFCOMMChannel.IsOpen]: Returns the state of the channel.
//   - [IIOBluetoothRFCOMMChannel.IsTransmissionPaused]: Returns TRUE if flow control is off.
//   - [IIOBluetoothRFCOMMChannel.RegisterForChannelCloseNotificationSelector]: Allows a client to register for a channel close notification.
//   - [IIOBluetoothRFCOMMChannel.SendRemoteLineStatus]: Sends an error to the remote side.
//   - [IIOBluetoothRFCOMMChannel.SetDelegate]: Allows an object to register itself as a client of the RFCOMM channel.
//   - [IIOBluetoothRFCOMMChannel.SetSerialParametersDataBitsParityStopBits]: Changes the parameters of the serial connection.
//   - [IIOBluetoothRFCOMMChannel.WriteAsyncLengthRefcon]: Sends a block of data in the channel asynchronously.
//   - [IIOBluetoothRFCOMMChannel.WriteSyncLength]: Sends a block of data in the channel synchronously.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel
type IIOBluetoothRFCOMMChannel interface {
	IIOBluetoothObject

	// Topic: Instance Methods

	// Close the channel.
	CloseChannel() kernel.IOReturn
	// Returns the object delegate
	Delegate() objectivec.IObject
	// Returns the Bluetooth Device that carries the rfcomm data.
	GetDevice() IIOBluetoothDevice
	// Returns the object rfcomm channel ID.
	GetChannelID() BluetoothRFCOMMChannelID
	// Returns the channel maximum transfer unit.
	GetMTU() BluetoothRFCOMMMTU
	// Returns the IOBluetoothObjectID of the given IOBluetoothRFCOMMChannel.
	GetObjectID() IOBluetoothObjectID
	// Returns an IOBluetoothRFCOMMChannelRef representation of the target IOBluetoothRFCOMMChannel object.
	GetRFCOMMChannelRef() IOBluetoothRFCOMMChannelRef
	// Returns the direction of the channel. An incoming channel is one that was opened by the remote device.
	IsIncoming() bool
	// Returns the state of the channel.
	IsOpen() bool
	// Returns TRUE if flow control is off.
	IsTransmissionPaused() bool
	// Allows a client to register for a channel close notification.
	RegisterForChannelCloseNotificationSelector(observer objectivec.IObject, inSelector objc.SEL) IIOBluetoothUserNotification
	// Sends an error to the remote side.
	SendRemoteLineStatus(lineStatus BluetoothRFCOMMLineStatus) kernel.IOReturn
	// Allows an object to register itself as a client of the RFCOMM channel.
	SetDelegate(delegate objectivec.IObject) kernel.IOReturn
	// Changes the parameters of the serial connection.
	SetSerialParametersDataBitsParityStopBits(speed uint32, nBits uint8, parity BluetoothRFCOMMParityType, bitStop uint8) kernel.IOReturn
	// Sends a block of data in the channel asynchronously.
	WriteAsyncLengthRefcon(data unsafe.Pointer, length uint16, refcon uintptr) kernel.IOReturn
	// Sends a block of data in the channel synchronously.
	WriteSyncLength(data unsafe.Pointer, length uint16) kernel.IOReturn
}

// Init initializes the instance.
func (b IOBluetoothRFCOMMChannel) Init() IOBluetoothRFCOMMChannel {
	rv := objc.Send[IOBluetoothRFCOMMChannel](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b IOBluetoothRFCOMMChannel) Autorelease() IOBluetoothRFCOMMChannel {
	rv := objc.Send[IOBluetoothRFCOMMChannel](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOBluetoothRFCOMMChannel creates a new IOBluetoothRFCOMMChannel instance.
func NewIOBluetoothRFCOMMChannel() IOBluetoothRFCOMMChannel {
	class := getIOBluetoothRFCOMMChannelClass()
	rv := objc.Send[IOBluetoothRFCOMMChannel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Close the channel.
//
// # Return Value
//
// An error code value. 0 if successful.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/close()
func (b IOBluetoothRFCOMMChannel) CloseChannel() kernel.IOReturn {
	rv := objc.Send[kernel.IOReturn](b.ID, objc.Sel("closeChannel"))
	return kernel.IOReturn(rv)
}

// Returns the object delegate
//
// # Return Value
//
// # The current delegate, or nil
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/delegate()
func (b IOBluetoothRFCOMMChannel) Delegate() objectivec.IObject {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("delegate"))
	return objectivec.Object{ID: rv}
}

// Returns the Bluetooth Device that carries the rfcomm data.
//
// # Return Value
//
// The IOBluetoothDevice object .
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/getDevice()
func (b IOBluetoothRFCOMMChannel) GetDevice() IIOBluetoothDevice {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("getDevice"))
	return IOBluetoothDeviceFromID(rv)
}

// Returns the object rfcomm channel ID.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/getID()
func (b IOBluetoothRFCOMMChannel) GetChannelID() BluetoothRFCOMMChannelID {
	rv := objc.Send[BluetoothRFCOMMChannelID](b.ID, objc.Sel("getChannelID"))
	return BluetoothRFCOMMChannelID(rv)
}

// Returns the channel maximum transfer unit.
//
// # Return Value
//
// Channel MTU size .
//
// # Discussion
//
// Returns the length of the largest chunk of data that this channel can
// carry. If the caller wishes to use the write:length:sleep: api the length
// of the data can not be bigger than the channel MTU (maximum transfer unit).
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/getMTU()
func (b IOBluetoothRFCOMMChannel) GetMTU() BluetoothRFCOMMMTU {
	rv := objc.Send[BluetoothRFCOMMMTU](b.ID, objc.Sel("getMTU"))
	return BluetoothRFCOMMMTU(rv)
}

// Returns the IOBluetoothObjectID of the given IOBluetoothRFCOMMChannel.
//
// # Return Value
//
// Returns the IOBluetoothObjectID of the given IOBluetoothRFCOMMChannel.
//
// # Discussion
//
// The IOBluetoothObjectID can be used as a global reference for a given
// IOBluetoothRFCOMMChannel. It allows two separate applications to refer to
// the same IOBluetoothRFCOMMChannel.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/getObjectID()
func (b IOBluetoothRFCOMMChannel) GetObjectID() IOBluetoothObjectID {
	rv := objc.Send[IOBluetoothObjectID](b.ID, objc.Sel("getObjectID"))
	return IOBluetoothObjectID(rv)
}

// Returns an IOBluetoothRFCOMMChannelRef representation of the target
// IOBluetoothRFCOMMChannel object.
//
// # Return Value
//
// Returns an IOBluetoothRFCOMMChannelRef representation of the target
// IOBluetoothRFCOMMChannel object.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/getRef()
func (b IOBluetoothRFCOMMChannel) GetRFCOMMChannelRef() IOBluetoothRFCOMMChannelRef {
	rv := objc.Send[IOBluetoothRFCOMMChannelRef](b.ID, objc.Sel("getRFCOMMChannelRef"))
	return IOBluetoothRFCOMMChannelRef(rv)
}

// Returns the direction of the channel. An incoming channel is one that was
// opened by the remote device.
//
// # Return Value
//
// Returns TRUE if the channel was opened by the remote device, FALSE if the
// channel was opened by this object.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/isIncoming()
func (b IOBluetoothRFCOMMChannel) IsIncoming() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("isIncoming"))
	return rv
}

// Returns the state of the channel.
//
// # Return Value
//
// TRUE if the channel state is open, FALSE otherwise.
//
// # Discussion
//
// note that “not open” means closed, opening and closing.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/isOpen()
func (b IOBluetoothRFCOMMChannel) IsOpen() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("isOpen"))
	return rv
}

// Returns TRUE if flow control is off.
//
// # Return Value
//
// TRUE if the action of sending data will block the current thread, FALSE
// otherwise.
//
// # Discussion
//
// Returns true if the remote device flow control is stopping out
// transmission. This is useful because we do not buffer data, we stop the
// transmitting actor. With this method the transmitter can check if sending
// data is going to be successful or is going to block.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/isTransmissionPaused()
func (b IOBluetoothRFCOMMChannel) IsTransmissionPaused() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("isTransmissionPaused"))
	return rv
}

// Allows a client to register for a channel close notification.
//
// observer: Target observer object
//
// inSelector: Selector to be sent to the observer when the RFCOMM channel is closed.
//
// # Return Value
//
// Returns an IOBluetoothUserNotification representing the outstanding RFCOMM
// channel close notification. To unregister the notification, call
// -unregister of the returned IOBluetoothUserNotification object. If an error
// is encountered creating the notification, nil is returned.
//
// # Discussion
//
// The given selector will be called on the target observer when the RFCOMM
// channel is closed. The selector should contain two arguments. The first is
// the user notification object. The second is the IOBluetoothRFCOMMChannel
// that was closed.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/register(forChannelCloseNotification:selector:)
func (b IOBluetoothRFCOMMChannel) RegisterForChannelCloseNotificationSelector(observer objectivec.IObject, inSelector objc.SEL) IIOBluetoothUserNotification {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("registerForChannelCloseNotification:selector:"), observer, inSelector)
	return IOBluetoothUserNotificationFromID(rv)
}

// Sends an error to the remote side.
//
// lineStatus: The error type. The error code can be NoError, OverrunError, ParityError or
// FramingError.
//
// # Return Value
//
// An error code value. 0 if successful.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/sendRemoteLineStatus(_:)
func (b IOBluetoothRFCOMMChannel) SendRemoteLineStatus(lineStatus BluetoothRFCOMMLineStatus) kernel.IOReturn {
	rv := objc.Send[kernel.IOReturn](b.ID, objc.Sel("sendRemoteLineStatus:"), lineStatus)
	return kernel.IOReturn(rv)
}

// Allows an object to register itself as a client of the RFCOMM channel.
//
// delegate: The object that will play the role of channel delegate [NOTE the rfcomm
// channel will reatin the delegate].
//
// # Return Value
//
// Returns kIOReturnSuccess if the delegate is successfully registered.
//
// # Discussion
//
// A channel delegate is the object the RFCOMM channel uses as target for data
// and events. The developer will implement only the the methods he/she is
// interested in. A list of the possible methods is at the end of this file in
// the definition of the informal protocol IOBluetoothRFCOMMChannelDelegate.
//
// NOTE: This method is only available in macOS 10.2.5 (Bluetooth v1.2) or
// later. NOTE: Before OS X 10.6, the delegate was retained. On 10.6 and
// later, it is not.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/setDelegate(_:)
func (b IOBluetoothRFCOMMChannel) SetDelegate(delegate objectivec.IObject) kernel.IOReturn {
	rv := objc.Send[kernel.IOReturn](b.ID, objc.Sel("setDelegate:"), delegate)
	return kernel.IOReturn(rv)
}

// Changes the parameters of the serial connection.
//
// speed: The baudrate.
//
// nBits: Number of data bits.
//
// parity: The type of parity can be NoParity, OddParity, EvenParity or MaxParity.
//
// bitStop: Number of stop bits.
//
// # Return Value
//
// An error code value. 0 if successful.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/setSerialParameters(_:dataBits:parity:stopBits:)
func (b IOBluetoothRFCOMMChannel) SetSerialParametersDataBitsParityStopBits(speed uint32, nBits uint8, parity BluetoothRFCOMMParityType, bitStop uint8) kernel.IOReturn {
	rv := objc.Send[kernel.IOReturn](b.ID, objc.Sel("setSerialParameters:dataBits:parity:stopBits:"), speed, nBits, parity, bitStop)
	return kernel.IOReturn(rv)
}

// Sends a block of data in the channel asynchronously.
//
// data: A pointer to the data buffer to be sent.
//
// length: The length of the buffer to be sent (in bytes).
//
// refcon: User supplied value that gets passed to the write callback.
//
// # Return Value
//
// Returns kIOReturnSuccess if the data was buffered successfully.
//
// # Discussion
//
// The number of bytes to be sent must not exceed the channel MTU. If the
// return value is an error condition none of the data was sent. Once the data
// has been successfully passed to the hardware to be transmitted, the
// delegate method -rfcommChannelWriteComplete:refcon:status: will be called
// with the refcon that was passed to this method.
//
// NOTE: This method is only available in macOS 10.2.5 (Bluetooth v1.2) or
// later.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/writeAsync(_:length:refcon:)
func (b IOBluetoothRFCOMMChannel) WriteAsyncLengthRefcon(data unsafe.Pointer, length uint16, refcon uintptr) kernel.IOReturn {
	rv := objc.Send[kernel.IOReturn](b.ID, objc.Sel("writeAsync:length:refcon:"), data, length, refcon)
	return kernel.IOReturn(rv)
}

// Sends a block of data in the channel synchronously.
//
// data: A pointer to the data buffer to be sent.
//
// length: The length of the buffer to be sent (in bytes).
//
// # Return Value
//
// Returns kIOReturnSuccess if the data was written successfully.
//
// # Discussion
//
// Sends data through the channel. The number of bytes to be sent must not
// exceed the channel MTU. If the return value is an error condition none of
// the data was sent. This method will block until the data has been
// successfully sent to the hardware for transmission (or until an error
// occurs).
//
// NOTE: This method is only available in macOS 10.2.5 (Bluetooth v1.2) or
// later.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/writeSync(_:length:)
func (b IOBluetoothRFCOMMChannel) WriteSyncLength(data unsafe.Pointer, length uint16) kernel.IOReturn {
	rv := objc.Send[kernel.IOReturn](b.ID, objc.Sel("writeSync:length:"), data, length)
	return kernel.IOReturn(rv)
}

// Allows a client to register for RFCOMM channel open notifications for any
// RFCOMM channel.
//
// object: Target object
//
// selector: Selector to be called on the target object when a new RFCOMM channel is
// opened. the format for the selector is: -(void)
// selectorName:(IOBluetoothUserNotification *)inNotification
// channel:(IOBluetoothRFCOMMChannel *)newChannel
//
// # Return Value
//
// Returns an IOBluetoothUserNotification representing the outstanding RFCOMM
// channel notification. To unregister the notification, call -unregister on
// the resulting IOBluetoothUserNotification object. If an error is
// encountered creating the notification, nil is returned. The returned
// IOBluetoothUserNotification will be valid for as long as the notification
// is registered. It is not necessary to retain the result. Once -unregister
// is called on it, it will no longer be valid.
//
// # Discussion
//
// The given selector will be called on the target object whenever any RFCOMM
// channel is opened. The selector should accept two arguments. The first is
// the user notification object. The second is the IOBluetoothRFCOMMChannel
// that was opened.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/register(forChannelOpenNotifications:selector:)
func (_IOBluetoothRFCOMMChannelClass IOBluetoothRFCOMMChannelClass) RegisterForChannelOpenNotificationsSelector(object objectivec.IObject, selector objc.SEL) IOBluetoothUserNotification {
	rv := objc.Send[objc.ID](objc.ID(_IOBluetoothRFCOMMChannelClass.class), objc.Sel("registerForChannelOpenNotifications:selector:"), object, selector)
	return IOBluetoothUserNotificationFromID(rv)
}

// Allows a client to register for RFCOMM channel open notifications for
// certain types of RFCOMM channels.
//
// object: Target object
//
// selector: Selector to be called on the target object when a new RFCOMM channel is
// opened. the format for the selector is: -(void)
// selectorName:(IOBluetoothUserNotification *)inNotification
// channel:(IOBluetoothRFCOMMChannel *)newChannel
//
// channelID: RFCOMM channel ID to match a new RFCOMM channel. If the channel ID
// doesn’t matter, 0 may be passed in.
//
// inDirection: The desired direction of the RFCOMM channel -
// kIOBluetoothUserNotificationChannelDirectionAny if the direction doesn’t
// matter.
//
// # Discussion
//
// The given selector will be called on the target object whenever an RFCOMM
// channel with the given attributes is opened. The selector should accept two
// arguments. The first is the user notification object. The second is the
// IOBluetoothRFCOMMChannel that was opened.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/register(forChannelOpenNotifications:selector:withChannelID:direction:)
func (_IOBluetoothRFCOMMChannelClass IOBluetoothRFCOMMChannelClass) RegisterForChannelOpenNotificationsSelectorWithChannelIDDirection(object objectivec.IObject, selector objc.SEL, channelID BluetoothRFCOMMChannelID, inDirection IOBluetoothUserNotificationChannelDirection) IOBluetoothUserNotification {
	rv := objc.Send[objc.ID](objc.ID(_IOBluetoothRFCOMMChannelClass.class), objc.Sel("registerForChannelOpenNotifications:selector:withChannelID:direction:"), object, selector, channelID, inDirection)
	return IOBluetoothUserNotificationFromID(rv)
}

// Returns the IObluetoothRFCOMMChannel with the given IOBluetoothObjectID.
//
// objectID: IOBluetoothObjectID of the desired IObluetoothRFCOMMChannel.
//
// # Return Value
//
// Returns the IObluetoothRFCOMMChannel that matches the given
// IOBluetoothObjectID if one exists. If no matching RFCOMM channel exists,
// nil is returned.
//
// # Discussion
//
// The IOBluetoothObjectID can be used as a global reference for a given
// IObluetoothRFCOMMChannel. It allows two separate applications to refer to
// the same IObluetoothRFCOMMChannel object.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/withObjectID(_:)
func (_IOBluetoothRFCOMMChannelClass IOBluetoothRFCOMMChannelClass) WithObjectID(objectID IOBluetoothObjectID) IOBluetoothRFCOMMChannel {
	rv := objc.Send[objc.ID](objc.ID(_IOBluetoothRFCOMMChannelClass.class), objc.Sel("withObjectID:"), objectID)
	return IOBluetoothRFCOMMChannelFromID(rv)
}

// Method call to convert an IOBluetoothRFCOMMChannelRef into an
// IOBluetoothRFCOMMChannel *.
//
// rfcommChannelRef: IOBluetoothRFCOMMChannelRef for which an IOBluetoothRFCOMMChannel * is
// desired.
//
// # Return Value
//
// Returns the IOBluetoothRFCOMMChannel * for the given
// IOBluetoothRFCOMMChannelRef.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannel/withRFCOMMChannelRef(_:)
func (_IOBluetoothRFCOMMChannelClass IOBluetoothRFCOMMChannelClass) WithRFCOMMChannelRef(rfcommChannelRef IOBluetoothRFCOMMChannelRef) IOBluetoothRFCOMMChannel {
	rv := objc.Send[objc.ID](objc.ID(_IOBluetoothRFCOMMChannelClass.class), objc.Sel("withRFCOMMChannelRef:"), rfcommChannelRef)
	return IOBluetoothRFCOMMChannelFromID(rv)
}
