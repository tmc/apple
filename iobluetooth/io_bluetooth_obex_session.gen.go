// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/kernel"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOBluetoothOBEXSession] class.
var (
	_IOBluetoothOBEXSessionClass     IOBluetoothOBEXSessionClass
	_IOBluetoothOBEXSessionClassOnce sync.Once
)

func getIOBluetoothOBEXSessionClass() IOBluetoothOBEXSessionClass {
	_IOBluetoothOBEXSessionClassOnce.Do(func() {
		_IOBluetoothOBEXSessionClass = IOBluetoothOBEXSessionClass{class: objc.GetClass("IOBluetoothOBEXSession")}
	})
	return _IOBluetoothOBEXSessionClass
}

// GetIOBluetoothOBEXSessionClass returns the class object for IOBluetoothOBEXSession.
func GetIOBluetoothOBEXSessionClass() IOBluetoothOBEXSessionClass {
	return getIOBluetoothOBEXSessionClass()
}

type IOBluetoothOBEXSessionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOBluetoothOBEXSessionClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOBluetoothOBEXSessionClass) Alloc() IOBluetoothOBEXSession {
	rv := objc.Send[IOBluetoothOBEXSession](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// An OBEX Session with a Bluetooth RFCOMM channel as the transport.
//
// # Initializers
//
//   - [IOBluetoothOBEXSession.InitWithDeviceChannelID]: Initializes a Bluetooth-based OBEX Session using a Bluetooth device.
//   - [IOBluetoothOBEXSession.InitWithIncomingRFCOMMChannelEventSelectorSelectorTargetRefCon]: Initializes a Bluetooth-based OBEX Session using an incoming RFCOMM channel.
//   - [IOBluetoothOBEXSession.InitWithSDPServiceRecord]: Initializes a Bluetooth-based OBEX Session using an SDP service record.
//
// # Instance Methods
//
//   - [IOBluetoothOBEXSession.GetDevice]: Get the Bluetooth Device being used by the session object.
//   - [IOBluetoothOBEXSession.GetRFCOMMChannel]: Get the Bluetooth RFCOMM channel being used by the session object.
//   - [IOBluetoothOBEXSession.IsSessionTargetAMac]: Tells whether the target device is a Mac by checking its service record.
//   - [IOBluetoothOBEXSession.RestartTransmission]: If the transmission was stopped due to the lack of buffers this call restarts it.
//   - [IOBluetoothOBEXSession.SendBufferTroughChannel]: Sends the next block of data through the rfcomm channel.
//   - [IOBluetoothOBEXSession.SetOBEXSessionOpenConnectionCallbackRefCon]: For C API support. Allows you to set the callback to be invoked when the OBEX connection is actually opened.
//   - [IOBluetoothOBEXSession.SetOpenTransportConnectionAsyncSelectorTargetRefCon]: Allows you to set the selector to be used when a transport connection is opened, or fails to open.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession
type IOBluetoothOBEXSession struct {
	OBEXSession
}

// IOBluetoothOBEXSessionFromID constructs a [IOBluetoothOBEXSession] from an objc.ID.
//
// An OBEX Session with a Bluetooth RFCOMM channel as the transport.
func IOBluetoothOBEXSessionFromID(id objc.ID) IOBluetoothOBEXSession {
	return IOBluetoothOBEXSession{OBEXSession: OBEXSessionFromID(id)}
}

// NOTE: IOBluetoothOBEXSession adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [IOBluetoothOBEXSession] class.
//
// # Initializers
//
//   - [IIOBluetoothOBEXSession.InitWithDeviceChannelID]: Initializes a Bluetooth-based OBEX Session using a Bluetooth device.
//   - [IIOBluetoothOBEXSession.InitWithIncomingRFCOMMChannelEventSelectorSelectorTargetRefCon]: Initializes a Bluetooth-based OBEX Session using an incoming RFCOMM channel.
//   - [IIOBluetoothOBEXSession.InitWithSDPServiceRecord]: Initializes a Bluetooth-based OBEX Session using an SDP service record.
//
// # Instance Methods
//
//   - [IIOBluetoothOBEXSession.GetDevice]: Get the Bluetooth Device being used by the session object.
//   - [IIOBluetoothOBEXSession.GetRFCOMMChannel]: Get the Bluetooth RFCOMM channel being used by the session object.
//   - [IIOBluetoothOBEXSession.IsSessionTargetAMac]: Tells whether the target device is a Mac by checking its service record.
//   - [IIOBluetoothOBEXSession.RestartTransmission]: If the transmission was stopped due to the lack of buffers this call restarts it.
//   - [IIOBluetoothOBEXSession.SendBufferTroughChannel]: Sends the next block of data through the rfcomm channel.
//   - [IIOBluetoothOBEXSession.SetOBEXSessionOpenConnectionCallbackRefCon]: For C API support. Allows you to set the callback to be invoked when the OBEX connection is actually opened.
//   - [IIOBluetoothOBEXSession.SetOpenTransportConnectionAsyncSelectorTargetRefCon]: Allows you to set the selector to be used when a transport connection is opened, or fails to open.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession
type IIOBluetoothOBEXSession interface {
	IOBEXSession
	IOBluetoothRFCOMMChannelDelegate

	// Topic: Initializers

	// Initializes a Bluetooth-based OBEX Session using a Bluetooth device.
	InitWithDeviceChannelID(inDevice IIOBluetoothDevice, inChannelID BluetoothRFCOMMChannelID) IOBluetoothOBEXSession
	// Initializes a Bluetooth-based OBEX Session using an incoming RFCOMM channel.
	InitWithIncomingRFCOMMChannelEventSelectorSelectorTargetRefCon(inChannel IIOBluetoothRFCOMMChannel, inEventSelector objc.SEL, inEventSelectorTarget objectivec.IObject, inUserRefCon uintptr) IOBluetoothOBEXSession
	// Initializes a Bluetooth-based OBEX Session using an SDP service record.
	InitWithSDPServiceRecord(inSDPServiceRecord IIOBluetoothSDPServiceRecord) IOBluetoothOBEXSession

	// Topic: Instance Methods

	// Get the Bluetooth Device being used by the session object.
	GetDevice() IIOBluetoothDevice
	// Get the Bluetooth RFCOMM channel being used by the session object.
	GetRFCOMMChannel() IIOBluetoothRFCOMMChannel
	// Tells whether the target device is a Mac by checking its service record.
	IsSessionTargetAMac() bool
	// If the transmission was stopped due to the lack of buffers this call restarts it.
	RestartTransmission()
	// Sends the next block of data through the rfcomm channel.
	SendBufferTroughChannel() kernel.IOReturn
	// For C API support. Allows you to set the callback to be invoked when the OBEX connection is actually opened.
	SetOBEXSessionOpenConnectionCallbackRefCon(inCallback IOBluetoothOBEXSessionOpenConnectionCallback, inUserRefCon uintptr)
	// Allows you to set the selector to be used when a transport connection is opened, or fails to open.
	SetOpenTransportConnectionAsyncSelectorTargetRefCon(inSelector objc.SEL, inSelectorTarget objectivec.IObject, inUserRefCon uintptr)
}

// Init initializes the instance.
func (b IOBluetoothOBEXSession) Init() IOBluetoothOBEXSession {
	rv := objc.Send[IOBluetoothOBEXSession](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b IOBluetoothOBEXSession) Autorelease() IOBluetoothOBEXSession {
	rv := objc.Send[IOBluetoothOBEXSession](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOBluetoothOBEXSession creates a new IOBluetoothOBEXSession instance.
func NewIOBluetoothOBEXSession() IOBluetoothOBEXSession {
	class := getIOBluetoothOBEXSessionClass()
	rv := objc.Send[IOBluetoothOBEXSession](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a Bluetooth-based OBEX Session using a Bluetooth device.
//
// inDevice: The bluetooth device on which to open the OBEXSession.
//
// inChannelID: The RFCOMM channel ID to use when opening the connection.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/init(device:channelID:)
func NewBluetoothOBEXSessionWithDeviceChannelID(inDevice IIOBluetoothDevice, inChannelID BluetoothRFCOMMChannelID) IOBluetoothOBEXSession {
	instance := getIOBluetoothOBEXSessionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:channelID:"), inDevice, inChannelID)
	return IOBluetoothOBEXSessionFromID(rv)
}

// Initializes a Bluetooth-based OBEX Session using an incoming RFCOMM
// channel.
//
// inChannel: RFCOMM channel ID of the desired channel to be used.
//
// inEventSelector: The selector to be called when an event is received.
//
// inEventSelectorTarget: The target object that get the selector message.
//
// inUserRefCon: Caller reference constant, pass whatever you want, it will be returned to
// you in the selector.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/init(incomingRFCOMMChannel:eventSelector:selectorTarget:refCon:)
func NewBluetoothOBEXSessionWithIncomingRFCOMMChannelEventSelectorSelectorTargetRefCon(inChannel IIOBluetoothRFCOMMChannel, inEventSelector objc.SEL, inEventSelectorTarget objectivec.IObject, inUserRefCon uintptr) IOBluetoothOBEXSession {
	instance := getIOBluetoothOBEXSessionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIncomingRFCOMMChannel:eventSelector:selectorTarget:refCon:"), inChannel, inEventSelector, inEventSelectorTarget, inUserRefCon)
	return IOBluetoothOBEXSessionFromID(rv)
}

// Initializes a Bluetooth-based OBEX Session using an SDP service record.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/init(sdpServiceRecord:)
func NewBluetoothOBEXSessionWithSDPServiceRecord(inSDPServiceRecord IIOBluetoothSDPServiceRecord) IOBluetoothOBEXSession {
	instance := getIOBluetoothOBEXSessionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSDPServiceRecord:"), inSDPServiceRecord)
	return IOBluetoothOBEXSessionFromID(rv)
}

// Initializes a Bluetooth-based OBEX Session using a Bluetooth device.
//
// inDevice: The bluetooth device on which to open the OBEXSession.
//
// inChannelID: The RFCOMM channel ID to use when opening the connection.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/init(device:channelID:)
func (b IOBluetoothOBEXSession) InitWithDeviceChannelID(inDevice IIOBluetoothDevice, inChannelID BluetoothRFCOMMChannelID) IOBluetoothOBEXSession {
	rv := objc.Send[IOBluetoothOBEXSession](b.ID, objc.Sel("initWithDevice:channelID:"), inDevice, inChannelID)
	return rv
}

// Initializes a Bluetooth-based OBEX Session using an incoming RFCOMM
// channel.
//
// inChannel: RFCOMM channel ID of the desired channel to be used.
//
// inEventSelector: The selector to be called when an event is received.
//
// inEventSelectorTarget: The target object that get the selector message.
//
// inUserRefCon: Caller reference constant, pass whatever you want, it will be returned to
// you in the selector.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/init(incomingRFCOMMChannel:eventSelector:selectorTarget:refCon:)
func (b IOBluetoothOBEXSession) InitWithIncomingRFCOMMChannelEventSelectorSelectorTargetRefCon(inChannel IIOBluetoothRFCOMMChannel, inEventSelector objc.SEL, inEventSelectorTarget objectivec.IObject, inUserRefCon uintptr) IOBluetoothOBEXSession {
	rv := objc.Send[IOBluetoothOBEXSession](b.ID, objc.Sel("initWithIncomingRFCOMMChannel:eventSelector:selectorTarget:refCon:"), inChannel, inEventSelector, inEventSelectorTarget, inUserRefCon)
	return rv
}

// Initializes a Bluetooth-based OBEX Session using an SDP service record.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/init(sdpServiceRecord:)
func (b IOBluetoothOBEXSession) InitWithSDPServiceRecord(inSDPServiceRecord IIOBluetoothSDPServiceRecord) IOBluetoothOBEXSession {
	rv := objc.Send[IOBluetoothOBEXSession](b.ID, objc.Sel("initWithSDPServiceRecord:"), inSDPServiceRecord)
	return rv
}

// Get the Bluetooth Device being used by the session object.
//
// # Return Value
//
// An IOBluetoothDevice object.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/getDevice()
func (b IOBluetoothOBEXSession) GetDevice() IIOBluetoothDevice {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("getDevice"))
	return IOBluetoothDeviceFromID(rv)
}

// Get the Bluetooth RFCOMM channel being used by the session object.
//
// # Return Value
//
// A IOBluetoothRFCOMMChannel object.
//
// # Discussion
//
// This could potentially be nil even though you have a valid OBEX session,
// because the RFCOMM channel is only valid when the session is connected.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/getRFCOMMChannel()
func (b IOBluetoothOBEXSession) GetRFCOMMChannel() IIOBluetoothRFCOMMChannel {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("getRFCOMMChannel"))
	return IOBluetoothRFCOMMChannelFromID(rv)
}

// Tells whether the target device is a Mac by checking its service record.
//
// # Return Value
//
// TRUE only if device service record has Mac entry, FALSE for all else.
//
// # Discussion
//
// Tells whether the target device is a Mac by checking its service record.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/isSessionTargetAMac()
func (b IOBluetoothOBEXSession) IsSessionTargetAMac() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("isSessionTargetAMac"))
	return rv
}

// If the transmission was stopped due to the lack of buffers this call
// restarts it.
//
// # Discussion
//
// If the transmission was stopeed due to the lack of buffers this call
// restarts it.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/restartTransmission()
func (b IOBluetoothOBEXSession) RestartTransmission() {
	objc.Send[objc.ID](b.ID, objc.Sel("restartTransmission"))
}

// Sends the next block of data through the rfcomm channel.
//
// # Discussion
//
// Since a send in the rfcomm channel is broken in multiple write calls (this
// actually is true only if the size is grater than the rfcomm MTU). Each
// write call is performed by sendBufferTroughChannel. This should never need
// to be overwritten.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/sendBufferTroughChannel()
func (b IOBluetoothOBEXSession) SendBufferTroughChannel() kernel.IOReturn {
	rv := objc.Send[kernel.IOReturn](b.ID, objc.Sel("sendBufferTroughChannel"))
	return kernel.IOReturn(rv)
}

// For C API support. Allows you to set the callback to be invoked when the
// OBEX connection is actually opened.
//
// inCallback: Function to call on the target.
//
// inUserRefCon: User’s reference constant, will be returned on the callback.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/setOBEXSessionOpenConnectionCallback(_:refCon:)
func (b IOBluetoothOBEXSession) SetOBEXSessionOpenConnectionCallbackRefCon(inCallback IOBluetoothOBEXSessionOpenConnectionCallback, inUserRefCon uintptr) {
	objc.Send[objc.ID](b.ID, objc.Sel("setOBEXSessionOpenConnectionCallback:refCon:"), inCallback, inUserRefCon)
}

// Allows you to set the selector to be used when a transport connection is
// opened, or fails to open.
//
// inSelector: Selector to call on the target.
//
// inSelectorTarget: Target to be called with the selector.
//
// inUserRefCon: User’s refCon that will get passed to them when their selector is
// invoked.
//
// # Discussion
//
// You do not need to call this on the session typically, unless you have
// subclassed the OBEXSession to implement a new transport and that transport
// supports async opening of connections. If it does not support async open,
// then using this is pointless.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/setOpenTransportConnectionAsyncSelector(_:target:refCon:)
func (b IOBluetoothOBEXSession) SetOpenTransportConnectionAsyncSelectorTargetRefCon(inSelector objc.SEL, inSelectorTarget objectivec.IObject, inUserRefCon uintptr) {
	objc.Send[objc.ID](b.ID, objc.Sel("setOpenTransportConnectionAsyncSelector:target:refCon:"), inSelector, inSelectorTarget, inUserRefCon)
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannelDelegate/rfcommChannelClosed(_:)
func (b IOBluetoothOBEXSession) RfcommChannelClosed(rfcommChannel IIOBluetoothRFCOMMChannel) {
	objc.Send[objc.ID](b.ID, objc.Sel("rfcommChannelClosed:"), rfcommChannel)
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannelDelegate/rfcommChannelControlSignalsChanged(_:)
func (b IOBluetoothOBEXSession) RfcommChannelControlSignalsChanged(rfcommChannel IIOBluetoothRFCOMMChannel) {
	objc.Send[objc.ID](b.ID, objc.Sel("rfcommChannelControlSignalsChanged:"), rfcommChannel)
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannelDelegate/rfcommChannelData(_:data:length:)
func (b IOBluetoothOBEXSession) RfcommChannelDataDataLength(rfcommChannel IIOBluetoothRFCOMMChannel, dataPointer unsafe.Pointer, dataLength uintptr) {
	objc.Send[objc.ID](b.ID, objc.Sel("rfcommChannelData:data:length:"), rfcommChannel, dataPointer, dataLength)
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannelDelegate/rfcommChannelFlowControlChanged(_:)
func (b IOBluetoothOBEXSession) RfcommChannelFlowControlChanged(rfcommChannel IIOBluetoothRFCOMMChannel) {
	objc.Send[objc.ID](b.ID, objc.Sel("rfcommChannelFlowControlChanged:"), rfcommChannel)
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannelDelegate/rfcommChannelOpenComplete(_:status:)
func (b IOBluetoothOBEXSession) RfcommChannelOpenCompleteStatus(rfcommChannel IIOBluetoothRFCOMMChannel, error_ kernel.IOReturn) {
	objc.Send[objc.ID](b.ID, objc.Sel("rfcommChannelOpenComplete:status:"), rfcommChannel, error_)
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannelDelegate/rfcommChannelQueueSpaceAvailable(_:)
func (b IOBluetoothOBEXSession) RfcommChannelQueueSpaceAvailable(rfcommChannel IIOBluetoothRFCOMMChannel) {
	objc.Send[objc.ID](b.ID, objc.Sel("rfcommChannelQueueSpaceAvailable:"), rfcommChannel)
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannelDelegate/rfcommChannelWriteComplete(_:refcon:status:)
func (b IOBluetoothOBEXSession) RfcommChannelWriteCompleteRefconStatus(rfcommChannel IIOBluetoothRFCOMMChannel, refcon uintptr, error_ kernel.IOReturn) {
	objc.Send[objc.ID](b.ID, objc.Sel("rfcommChannelWriteComplete:refcon:status:"), rfcommChannel, refcon, error_)
}

// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothRFCOMMChannelDelegate/rfcommChannelWriteComplete(_:refcon:status:bytesWritten:)
func (b IOBluetoothOBEXSession) RfcommChannelWriteCompleteRefconStatusBytesWritten(rfcommChannel IIOBluetoothRFCOMMChannel, refcon uintptr, error_ kernel.IOReturn, length uintptr) {
	objc.Send[objc.ID](b.ID, objc.Sel("rfcommChannelWriteComplete:refcon:status:bytesWritten:"), rfcommChannel, refcon, error_, length)
}

// Creates a Bluetooth-based OBEX Session using a Bluetooth device and a
// Bluetooth RFCOMM channel ID.
//
// inDevice: A valid Bluetooth device describing which device you want to connect to
// with Bluetooth/OBEX.
//
// inRFCOMMChannelID: An RFCOMM Channel ID numbe that is available on the remote device. This
// channel will be used when the transport connection is attempted.
//
// # Discussion
//
// Note that this does NOT mean the transport connection was open. It will be
// opened when OBEXConnect is invoked on the session object.
//
// IMPORTANT NOTE In Bluetooth framework version 1.0.0, the session returned
// will NOT be autoreleased as it should be according to objc convention. This
// has been changed starting in Bluetooth version 1.0.1 and later, so it WILL
// be autoreleased upon return, so you will need to retain it if you want to
// reference it later.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/withDevice(_:channelID:)
func (_IOBluetoothOBEXSessionClass IOBluetoothOBEXSessionClass) WithDeviceChannelID(inDevice IIOBluetoothDevice, inRFCOMMChannelID BluetoothRFCOMMChannelID) IOBluetoothOBEXSession {
	rv := objc.Send[objc.ID](objc.ID(_IOBluetoothOBEXSessionClass.class), objc.Sel("withDevice:channelID:"), inDevice, inRFCOMMChannelID)
	return IOBluetoothOBEXSessionFromID(rv)
}

// Creates a Bluetooth-based OBEX Session using an incoming RFCOMM channel.
//
// inChannel: The channel to use to create a connection to a device.
//
// inEventSelector: The selector that gets called when an event occurs on the OBEX Session.
//
// inEventSelectorTarget: The object that is used to call the above selector.
//
// inUserRefCon: The reference constant. Pass whatever you wish - it will be returned to you
// in the selector.
//
// # Discussion
//
// IMPORTANT NOTE In Bluetooth framework version 1.0.0, the session returned
// will NOT be autoreleased as it should be according to objc convention. This
// has been changed starting in Bluetooth version 1.0.1 and later, so it WILL
// be autoreleased upon return, so you will need to retain it if you want to
// reference it later.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/withIncomingRFCOMMChannel(_:eventSelector:selectorTarget:refCon:)
func (_IOBluetoothOBEXSessionClass IOBluetoothOBEXSessionClass) WithIncomingRFCOMMChannelEventSelectorSelectorTargetRefCon(inChannel IIOBluetoothRFCOMMChannel, inEventSelector objc.SEL, inEventSelectorTarget objectivec.IObject, inUserRefCon uintptr) IOBluetoothOBEXSession {
	rv := objc.Send[objc.ID](objc.ID(_IOBluetoothOBEXSessionClass.class), objc.Sel("withIncomingRFCOMMChannel:eventSelector:selectorTarget:refCon:"), inChannel, inEventSelector, inEventSelectorTarget, inUserRefCon)
	return IOBluetoothOBEXSessionFromID(rv)
}

// Creates a Bluetooth-based OBEX Session using an SDP service record,
// typically obtained from a device/service browser window controller.
//
// inSDPServiceRecord: A valid SDP service record describing the service (and RFCOMM channel) you
// want to connect to with Bluetooth/OBEX.
//
// # Return Value
//
// An OBEX session representing the device/rfcomm channel found in the service
// record. nil if we failed.
//
// # Discussion
//
// Note that this does NOT mean the transport connection was open. It will be
// opened when OBEXConnect is invoked on the session object.
//
// IMPORTANT NOTE In Bluetooth framework version 1.0.0, the session returned
// will NOT be autoreleased as it should be according to objc convention. This
// has been changed starting in Bluetooth version 1.0.1 and later, so it WILL
// be autoreleased upon return, so you will need to retain it if you want to
// reference it later.
//
// See: https://developer.apple.com/documentation/IOBluetooth/IOBluetoothOBEXSession/withSDPServiceRecord(_:)
func (_IOBluetoothOBEXSessionClass IOBluetoothOBEXSessionClass) WithSDPServiceRecord(inSDPServiceRecord IIOBluetoothSDPServiceRecord) IOBluetoothOBEXSession {
	rv := objc.Send[objc.ID](objc.ID(_IOBluetoothOBEXSessionClass.class), objc.Sel("withSDPServiceRecord:"), inSDPServiceRecord)
	return IOBluetoothOBEXSessionFromID(rv)
}

// Protocol methods for IOBluetoothRFCOMMChannelDelegate
