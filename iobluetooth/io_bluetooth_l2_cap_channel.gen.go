// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/kernel"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOBluetoothL2CAPChannel] class.
var (
	_IOBluetoothL2CAPChannelClass     IOBluetoothL2CAPChannelClass
	_IOBluetoothL2CAPChannelClassOnce sync.Once
)

func getIOBluetoothL2CAPChannelClass() IOBluetoothL2CAPChannelClass {
	_IOBluetoothL2CAPChannelClassOnce.Do(func() {
		_IOBluetoothL2CAPChannelClass = IOBluetoothL2CAPChannelClass{class: objc.GetClass("IOBluetoothL2CAPChannel")}
	})
	return _IOBluetoothL2CAPChannelClass
}

// GetIOBluetoothL2CAPChannelClass returns the class object for IOBluetoothL2CAPChannel.
func GetIOBluetoothL2CAPChannelClass() IOBluetoothL2CAPChannelClass {
	return getIOBluetoothL2CAPChannelClass()
}

type IOBluetoothL2CAPChannelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOBluetoothL2CAPChannelClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOBluetoothL2CAPChannelClass) Alloc() IOBluetoothL2CAPChannel {
	rv := objc.Send[IOBluetoothL2CAPChannel](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// An instance of IOBluetoothL2CAPChannel represents a single open L2CAP
// channel.
//
// # Overview
//
// A client won’t create IOBluetoothL2CAPChannel objects directly. Instead,
// the IOBluetoothDevice’s L2CAP channel open API is responsible for opening
// a new L2CAP channel and returning an IOBluetoothL2CAPChannel instance
// representing that newly opened channel. Additionally, the IOBluetooth
// notification system will send notifications when new L2CAP channels are
// open (if requested).
//
// After a new L2CAP channel is opened, the L2CAP configuration process will
// not be completed until an incoming data listener is registered with the
// IOBluetoothL2CAPChannel object. The reason for this is to due to the
// limited buffering done of incoming L2CAP data. This way, we avoid the
// situation where incoming data is received before the client is ready for
// it. Once a client is done with an IOBluetoothL2CAPChannel that it opened,
// it should call -closeChannel. Additionally, if the client does not intend
// to use the connection to the remote device any further, it should call
// -closeConnection on the IOBluetoothDevice object.
//
// # Instance Properties
//
//   - [IOBluetoothL2CAPChannel.Device]: Returns the IOBluetoothDevice to which the target L2CAP channel is open.
//   - [IOBluetoothL2CAPChannel.IncomingMTU]: Returns the current incoming MTU for the L2CAP channel.
//   - [IOBluetoothL2CAPChannel.LocalChannelID]: Returns the local L2CAP channel ID for the target L2CAP channel.
//   - [IOBluetoothL2CAPChannel.ObjectID]: Returns the IOBluetoothObjectID of the given IOBluetoothL2CAPChannel.
//   - [IOBluetoothL2CAPChannel.OutgoingMTU]: Returns the current outgoing MTU for the L2CAP channel.
//   - [IOBluetoothL2CAPChannel.PSM]: Returns the PSM for the target L2CAP channel.
//   - [IOBluetoothL2CAPChannel.RemoteChannelID]: Returns the remote L2CAP channel ID for the target L2CAP channel.
//
// # Instance Methods
//
//   - [IOBluetoothL2CAPChannel.CloseChannel]: Initiates the close process on an open L2CAP channel.
//   - [IOBluetoothL2CAPChannel.Delegate]: Returns the currently assigned delegate
//   - [IOBluetoothL2CAPChannel.IsIncoming]: Returns TRUE if the channel is an incoming channel.
//   - [IOBluetoothL2CAPChannel.RegisterForChannelCloseNotificationSelector]: Allows a client to register for a channel close notification.
//   - [IOBluetoothL2CAPChannel.RequestRemoteMTU]: Initiates the process to reconfigure the L2CAP channel with a new outgoing MTU.
//   - [IOBluetoothL2CAPChannel.SetDelegate]: Allows an object to register itself as client of the L2CAP channel.
//   - [IOBluetoothL2CAPChannel.SetDelegateWithConfiguration]: Allows an object to register itself as client of the L2CAP channel.
//   - [IOBluetoothL2CAPChannel.WriteAsyncLengthRefcon]: Writes the given data over the target L2CAP channel asynchronously to the remote device.
//   - [IOBluetoothL2CAPChannel.WriteAsyncTrapLengthRefcon]
//   - [IOBluetoothL2CAPChannel.WriteSyncLength]: Writes the given data synchronously over the target L2CAP channel to the remote device.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel
type IOBluetoothL2CAPChannel struct {
	IOBluetoothObject
}

// IOBluetoothL2CAPChannelFromID constructs a [IOBluetoothL2CAPChannel] from an objc.ID.
//
// An instance of IOBluetoothL2CAPChannel represents a single open L2CAP
// channel.
func IOBluetoothL2CAPChannelFromID(id objc.ID) IOBluetoothL2CAPChannel {
	return IOBluetoothL2CAPChannel{IOBluetoothObject: IOBluetoothObjectFromID(id)}
}

// NOTE: IOBluetoothL2CAPChannel adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [IOBluetoothL2CAPChannel] class.
//
// # Instance Properties
//
//   - [IIOBluetoothL2CAPChannel.Device]: Returns the IOBluetoothDevice to which the target L2CAP channel is open.
//   - [IIOBluetoothL2CAPChannel.IncomingMTU]: Returns the current incoming MTU for the L2CAP channel.
//   - [IIOBluetoothL2CAPChannel.LocalChannelID]: Returns the local L2CAP channel ID for the target L2CAP channel.
//   - [IIOBluetoothL2CAPChannel.ObjectID]: Returns the IOBluetoothObjectID of the given IOBluetoothL2CAPChannel.
//   - [IIOBluetoothL2CAPChannel.OutgoingMTU]: Returns the current outgoing MTU for the L2CAP channel.
//   - [IIOBluetoothL2CAPChannel.PSM]: Returns the PSM for the target L2CAP channel.
//   - [IIOBluetoothL2CAPChannel.RemoteChannelID]: Returns the remote L2CAP channel ID for the target L2CAP channel.
//
// # Instance Methods
//
//   - [IIOBluetoothL2CAPChannel.CloseChannel]: Initiates the close process on an open L2CAP channel.
//   - [IIOBluetoothL2CAPChannel.Delegate]: Returns the currently assigned delegate
//   - [IIOBluetoothL2CAPChannel.IsIncoming]: Returns TRUE if the channel is an incoming channel.
//   - [IIOBluetoothL2CAPChannel.RegisterForChannelCloseNotificationSelector]: Allows a client to register for a channel close notification.
//   - [IIOBluetoothL2CAPChannel.RequestRemoteMTU]: Initiates the process to reconfigure the L2CAP channel with a new outgoing MTU.
//   - [IIOBluetoothL2CAPChannel.SetDelegate]: Allows an object to register itself as client of the L2CAP channel.
//   - [IIOBluetoothL2CAPChannel.SetDelegateWithConfiguration]: Allows an object to register itself as client of the L2CAP channel.
//   - [IIOBluetoothL2CAPChannel.WriteAsyncLengthRefcon]: Writes the given data over the target L2CAP channel asynchronously to the remote device.
//   - [IIOBluetoothL2CAPChannel.WriteAsyncTrapLengthRefcon]
//   - [IIOBluetoothL2CAPChannel.WriteSyncLength]: Writes the given data synchronously over the target L2CAP channel to the remote device.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel
type IIOBluetoothL2CAPChannel interface {
	IIOBluetoothObject

	// Topic: Instance Properties

	// Returns the IOBluetoothDevice to which the target L2CAP channel is open.
	Device() IIOBluetoothDevice
	// Returns the current incoming MTU for the L2CAP channel.
	IncomingMTU() BluetoothL2CAPMTU
	// Returns the local L2CAP channel ID for the target L2CAP channel.
	LocalChannelID() BluetoothL2CAPChannelID
	// Returns the IOBluetoothObjectID of the given IOBluetoothL2CAPChannel.
	ObjectID() IOBluetoothObjectID
	// Returns the current outgoing MTU for the L2CAP channel.
	OutgoingMTU() BluetoothL2CAPMTU
	// Returns the PSM for the target L2CAP channel.
	PSM() BluetoothL2CAPPSM
	// Returns the remote L2CAP channel ID for the target L2CAP channel.
	RemoteChannelID() BluetoothL2CAPChannelID

	// Topic: Instance Methods

	// Initiates the close process on an open L2CAP channel.
	CloseChannel() kernel.IOReturn
	// Returns the currently assigned delegate
	Delegate() objectivec.IObject
	// Returns TRUE if the channel is an incoming channel.
	IsIncoming() bool
	// Allows a client to register for a channel close notification.
	RegisterForChannelCloseNotificationSelector(observer objectivec.IObject, inSelector objc.SEL) IIOBluetoothUserNotification
	// Initiates the process to reconfigure the L2CAP channel with a new outgoing MTU.
	RequestRemoteMTU(remoteMTU BluetoothL2CAPMTU) kernel.IOReturn
	// Allows an object to register itself as client of the L2CAP channel.
	SetDelegate(channelDelegate objectivec.IObject) kernel.IOReturn
	// Allows an object to register itself as client of the L2CAP channel.
	SetDelegateWithConfiguration(channelDelegate objectivec.IObject, channelConfiguration foundation.INSDictionary) kernel.IOReturn
	// Writes the given data over the target L2CAP channel asynchronously to the remote device.
	WriteAsyncLengthRefcon(data unsafe.Pointer, length uint16, refcon uintptr) kernel.IOReturn
	WriteAsyncTrapLengthRefcon(data unsafe.Pointer, length uint16, refcon uintptr) kernel.IOReturn
	// Writes the given data synchronously over the target L2CAP channel to the remote device.
	WriteSyncLength(data unsafe.Pointer, length uint16) kernel.IOReturn
}

// Init initializes the instance.
func (b IOBluetoothL2CAPChannel) Init() IOBluetoothL2CAPChannel {
	rv := objc.Send[IOBluetoothL2CAPChannel](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b IOBluetoothL2CAPChannel) Autorelease() IOBluetoothL2CAPChannel {
	rv := objc.Send[IOBluetoothL2CAPChannel](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOBluetoothL2CAPChannel creates a new IOBluetoothL2CAPChannel instance.
func NewIOBluetoothL2CAPChannel() IOBluetoothL2CAPChannel {
	class := getIOBluetoothL2CAPChannelClass()
	rv := objc.Send[IOBluetoothL2CAPChannel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initiates the close process on an open L2CAP channel.
//
// # Return Value
//
// Returns kIOReturnSuccess on success.
//
// # Discussion
//
// This method may only be called by the client that opened the channel in the
// first place. In the future asynchronous and synchronous versions will be
// provided that let the client know when the close process has been finished.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/close()
func (b IOBluetoothL2CAPChannel) CloseChannel() kernel.IOReturn {
	rv := objc.Send[kernel.IOReturn](b.ID, objc.Sel("closeChannel"))
	return kernel.IOReturn(rv)
}

// Returns the currently assigned delegate
//
// # Return Value
//
// Returns the current delegate, or nil if one is not set.
//
// # Discussion
//
// An incoming channel is one that was initiated by a remote device.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/delegate()
func (b IOBluetoothL2CAPChannel) Delegate() objectivec.IObject {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("delegate"))
	return objectivec.Object{ID: rv}
}

// Returns TRUE if the channel is an incoming channel.
//
// # Return Value
//
// Returns TRUE if the channel is an incoming channel.
//
// # Discussion
//
// An incoming channel is one that was initiated by a remote device.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/isIncoming()
func (b IOBluetoothL2CAPChannel) IsIncoming() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("isIncoming"))
	return rv
}

// Allows a client to register for a channel close notification.
//
// observer: Target observer object
//
// inSelector: Selector to be sent to the observer when the L2CAP channel is closed.
//
// # Return Value
//
// Returns an IOBluetoothUserNotification representing the outstanding L2CAP
// channel close notification. To unregister the notification, call
// -unregister of the returned IOBluetoothUserNotification object. If an error
// is encountered creating the notification, nil is returned.
//
// # Discussion
//
// The given selector will be called on the target observer when the L2CAP
// channel is closed. The selector should contain two arguments. The first is
// the user notification object. The second is the IOBluetoothL2CAPChannel
// that was closed.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/register(forChannelCloseNotification:selector:)
func (b IOBluetoothL2CAPChannel) RegisterForChannelCloseNotificationSelector(observer objectivec.IObject, inSelector objc.SEL) IIOBluetoothUserNotification {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("registerForChannelCloseNotification:selector:"), observer, inSelector)
	return IOBluetoothUserNotificationFromID(rv)
}

// Initiates the process to reconfigure the L2CAP channel with a new outgoing
// MTU.
//
// remoteMTU: The desired outgoing MTU.
//
// # Return Value
//
// Returns kIOReturnSuccess if the channel re-configure process was
// successfully initiated.
//
// # Discussion
//
// Currently, this API does not give an indication that the re-config process
// has completed. In the future additional API will be available to provide
// that information both synchronously and asynchronously.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/requestRemoteMTU(_:)
func (b IOBluetoothL2CAPChannel) RequestRemoteMTU(remoteMTU BluetoothL2CAPMTU) kernel.IOReturn {
	rv := objc.Send[kernel.IOReturn](b.ID, objc.Sel("requestRemoteMTU:"), remoteMTU)
	return kernel.IOReturn(rv)
}

// Allows an object to register itself as client of the L2CAP channel.
//
// channelDelegate: The object that will play the role of channel delegate [NOTE the l2cap
// channel will retain the delegate].
//
// # Return Value
//
// Returns kIOReturnSuccess if the delegate is successfully registered.
//
// # Discussion
//
// A channel delegate is the object the L2CAP channel uses as target for data
// and events. The developer will implement only the the methods he/she is
// interested in. A list of the possible methods is at the end of this file in
// the definition of the informal protocol IOBluetoothL2CAPChannelDelegate. A
// newly opened L2CAP channel will not complete its configuration process
// until the client that opened it registers a connectionHandler. This
// prevents that case where incoming data is received before the client is
// ready.
//
// NOTE: This method is only available in macOS 10.2.5 (Bluetooth v1.2) or
// later.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/setDelegate(_:)
func (b IOBluetoothL2CAPChannel) SetDelegate(channelDelegate objectivec.IObject) kernel.IOReturn {
	rv := objc.Send[kernel.IOReturn](b.ID, objc.Sel("setDelegate:"), channelDelegate)
	return kernel.IOReturn(rv)
}

// Allows an object to register itself as client of the L2CAP channel.
//
// channelDelegate: The object that will play the role of channel delegate.
//
// channelConfiguration: The dictionary that describes the initial configuration for the channel.
//
// # Return Value
//
// Returns kIOReturnSuccess if the delegate is successfully registered.
//
// # Discussion
//
// A channel delegate is the object the L2CAP channel uses as target for data
// and events. The developer will implement only the the methods he/she is
// interested in. A list of the possible methods is at the end of this file in
// the definition of the informal protocol IOBluetoothL2CAPChannelDelegate. A
// newly opened L2CAP channel will not complete its configuration process
// until the client that opened it registers a connectionHandler. This
// prevents that case where incoming data is received before the client is
// ready.
//
// NOTE: This method is only available in macOS 10.5 (Bluetooth v2.0) or
// later.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/setDelegate(_:withConfiguration:)
func (b IOBluetoothL2CAPChannel) SetDelegateWithConfiguration(channelDelegate objectivec.IObject, channelConfiguration foundation.INSDictionary) kernel.IOReturn {
	rv := objc.Send[kernel.IOReturn](b.ID, objc.Sel("setDelegate:withConfiguration:"), channelDelegate, channelConfiguration)
	return kernel.IOReturn(rv)
}

// Writes the given data over the target L2CAP channel asynchronously to the
// remote device.
//
// data: Pointer to the buffer containing the data to send.
//
// length: The length of the given data buffer.
//
// refcon: User supplied value that gets passed to the write callback.
//
// # Return Value
//
// Returns kIOReturnSuccess if the data was buffered successfully.
//
// # Discussion
//
// The length of the data may not exceed the L2CAP channel’s ougoing MTU.
// When the data has been successfully passed to the hardware to be
// transmitted, the delegate method -l2capChannelWriteComplete:refcon:status:
// will be called with the refcon passed into this method.
//
// NOTE: This method is only available in macOS 10.2.5 (Bluetooth v1.2) or
// later.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/writeAsync(_:length:refcon:)
func (b IOBluetoothL2CAPChannel) WriteAsyncLengthRefcon(data unsafe.Pointer, length uint16, refcon uintptr) kernel.IOReturn {
	rv := objc.Send[kernel.IOReturn](b.ID, objc.Sel("writeAsync:length:refcon:"), data, length, refcon)
	return kernel.IOReturn(rv)
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/writeAsyncTrap(_:length:refcon:)
func (b IOBluetoothL2CAPChannel) WriteAsyncTrapLengthRefcon(data unsafe.Pointer, length uint16, refcon uintptr) kernel.IOReturn {
	rv := objc.Send[kernel.IOReturn](b.ID, objc.Sel("writeAsyncTrap:length:refcon:"), data, length, refcon)
	return kernel.IOReturn(rv)
}

// Writes the given data synchronously over the target L2CAP channel to the
// remote device.
//
// data: Pointer to the buffer containing the data to send.
//
// length: The length of the given data buffer.
//
// # Return Value
//
// Returns kIOReturnSuccess if the data was written successfully.
//
// # Discussion
//
// The length of the data may not exceed the L2CAP channel’s ougoing MTU.
// This method will block until the data has been successfully sent to the
// hardware for transmission (or an error occurs).
//
// NOTE: This method is only available in macOS 10.2.5 (Bluetooth v1.2) or
// later.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/writeSync(_:length:)
func (b IOBluetoothL2CAPChannel) WriteSyncLength(data unsafe.Pointer, length uint16) kernel.IOReturn {
	rv := objc.Send[kernel.IOReturn](b.ID, objc.Sel("writeSync:length:"), data, length)
	return kernel.IOReturn(rv)
}

// Allows a client to register for L2CAP channel open notifications for any
// L2CAP channel.
//
// object: Target object
//
// selector: Selector to be called on the target object when a new L2CAP channel is
// opened.
//
// # Return Value
//
// Returns an IOBluetoothUserNotification representing the outstanding L2CAP
// channel notification. To unregister the notification, call -unregister on
// the resulting IOBluetoothUserNotification object. If an error is
// encountered creating the notification, nil is returned. The returned
// IOBluetoothUserNotification will be valid for as long as the notification
// is registered. It is not necessary to retain the result. Once -unregister
// is called on it, it will no longer be valid.
//
// # Discussion
//
// The given selector will be called on the target object whenever any L2CAP
// channel is opened. The selector should accept two arguments. The first is
// the user notification object. The second is the IOBluetoothL2CAPChannel
// that was opened.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/register(forChannelOpenNotifications:selector:)
func (_IOBluetoothL2CAPChannelClass IOBluetoothL2CAPChannelClass) RegisterForChannelOpenNotificationsSelector(object objectivec.IObject, selector objc.SEL) IOBluetoothUserNotification {
	rv := objc.Send[objc.ID](objc.ID(_IOBluetoothL2CAPChannelClass.class), objc.Sel("registerForChannelOpenNotifications:selector:"), object, selector)
	return IOBluetoothUserNotificationFromID(rv)
}

// Allows a client to register for L2CAP channel open notifications for
// certain types of L2CAP channels.
//
// object: Target object
//
// selector: Selector to be called on the target object when a new L2CAP channel is
// opened.
//
// psm: PSM to match a new L2CAP channel. If the PSM doesn’t matter, 0 may be
// passed in.
//
// inDirection: The desired direction of the L2CAP channel -
// kIOBluetoothUserNotificationChannelDirectionAny if the direction doesn’t
// matter.
//
// # Return Value
//
// Returns an IOBluetoothUserNotification representing the outstanding L2CAP
// channel notification. To unregister the notification, call -unregister on
// the resulting IOBluetoothUserNotification object. If an error is
// encountered creating the notification, nil is returned. The returned
// IOBluetoothUserNotification will be valid for as long as the notification
// is registered. It is not necessary to retain the result. Once -unregister
// is called on it, it will no longer be valid.
//
// # Discussion
//
// The given selector will be called on the target object whenever an L2CAP
// channel with the given attributes is opened. The selector should accept two
// arguments. The first is the user notification object. The second is the
// IOBluetoothL2CAPChannel that was opened.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/register(forChannelOpenNotifications:selector:withPSM:direction:)
func (_IOBluetoothL2CAPChannelClass IOBluetoothL2CAPChannelClass) RegisterForChannelOpenNotificationsSelectorWithPSMDirection(object objectivec.IObject, selector objc.SEL, psm BluetoothL2CAPPSM, inDirection IOBluetoothUserNotificationChannelDirection) IOBluetoothUserNotification {
	rv := objc.Send[objc.ID](objc.ID(_IOBluetoothL2CAPChannelClass.class), objc.Sel("registerForChannelOpenNotifications:selector:withPSM:direction:"), object, selector, psm, inDirection)
	return IOBluetoothUserNotificationFromID(rv)
}

// Returns the IObluetoothL2CAPChannel with the given IOBluetoothObjectID.
//
// objectID: IOBluetoothObjectID of the desired IOBluetoothL2CAPChannel.
//
// # Return Value
//
// Returns the IOBluetoothL2CAPChannel that matches the given
// IOBluetoothObjectID if one exists. If no matching L2CAP channel exists, nil
// is returned.
//
// # Discussion
//
// The IOBluetoothObjectID can be used as a global reference for a given
// IOBluetoothL2CAPChannel. It allows two separate applications to refer to
// the same IOBluetoothL2CAPChannel object.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/withObjectID(_:)
func (_IOBluetoothL2CAPChannelClass IOBluetoothL2CAPChannelClass) WithObjectID(objectID IOBluetoothObjectID) IOBluetoothL2CAPChannel {
	rv := objc.Send[objc.ID](objc.ID(_IOBluetoothL2CAPChannelClass.class), objc.Sel("withObjectID:"), objectID)
	return IOBluetoothL2CAPChannelFromID(rv)
}

// Returns the IOBluetoothDevice to which the target L2CAP channel is open.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/device
func (b IOBluetoothL2CAPChannel) Device() IIOBluetoothDevice {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("device"))
	return IOBluetoothDeviceFromID(objc.ID(rv))
}

// Returns the current incoming MTU for the L2CAP channel.
//
// # Discussion
//
// The incoming MTU represents the maximum L2CAP packet size for packets being
// sent by the remote device.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/incomingMTU
func (b IOBluetoothL2CAPChannel) IncomingMTU() BluetoothL2CAPMTU {
	rv := objc.Send[BluetoothL2CAPMTU](b.ID, objc.Sel("incomingMTU"))
	return BluetoothL2CAPMTU(rv)
}

// Returns the local L2CAP channel ID for the target L2CAP channel.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/localChannelID
func (b IOBluetoothL2CAPChannel) LocalChannelID() BluetoothL2CAPChannelID {
	rv := objc.Send[BluetoothL2CAPChannelID](b.ID, objc.Sel("localChannelID"))
	return BluetoothL2CAPChannelID(rv)
}

// Returns the IOBluetoothObjectID of the given IOBluetoothL2CAPChannel.
//
// # Discussion
//
// The IOBluetoothObjectID can be used as a global reference for a given
// IOBluetoothL2CAPChannel. It allows two separate applications to refer to
// the same IOBluetoothL2CAPChannel.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/objectID
func (b IOBluetoothL2CAPChannel) ObjectID() IOBluetoothObjectID {
	rv := objc.Send[IOBluetoothObjectID](b.ID, objc.Sel("objectID"))
	return IOBluetoothObjectID(rv)
}

// Returns the current outgoing MTU for the L2CAP channel.
//
// # Discussion
//
// The outgoing MTU represents the maximum L2CAP packet size for packets being
// sent to the remote device.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/outgoingMTU
func (b IOBluetoothL2CAPChannel) OutgoingMTU() BluetoothL2CAPMTU {
	rv := objc.Send[BluetoothL2CAPMTU](b.ID, objc.Sel("outgoingMTU"))
	return BluetoothL2CAPMTU(rv)
}

// Returns the PSM for the target L2CAP channel.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/psm
func (b IOBluetoothL2CAPChannel) PSM() BluetoothL2CAPPSM {
	rv := objc.Send[BluetoothL2CAPPSM](b.ID, objc.Sel("PSM"))
	return BluetoothL2CAPPSM(rv)
}

// Returns the remote L2CAP channel ID for the target L2CAP channel.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothL2CAPChannel/remoteChannelID
func (b IOBluetoothL2CAPChannel) RemoteChannelID() BluetoothL2CAPChannelID {
	rv := objc.Send[BluetoothL2CAPChannelID](b.ID, objc.Sel("remoteChannelID"))
	return BluetoothL2CAPChannelID(rv)
}
