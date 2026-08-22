// Code generated from Apple documentation for IOBluetooth. DO NOT EDIT.

package iobluetooth

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [OBEXSession] class.
var (
	_OBEXSessionClass     OBEXSessionClass
	_OBEXSessionClassOnce sync.Once
)

func getOBEXSessionClass() OBEXSessionClass {
	_OBEXSessionClassOnce.Do(func() {
		_OBEXSessionClass = OBEXSessionClass{class: objc.GetClass("OBEXSession")}
	})
	return _OBEXSessionClass
}

// GetOBEXSessionClass returns the class object for OBEXSession.
func GetOBEXSessionClass() OBEXSessionClass {
	return getOBEXSessionClass()
}

type OBEXSessionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (oc OBEXSessionClass) Class() objc.Class {
	return oc.class
}

// Alloc allocates memory for a new instance of the class.
func (oc OBEXSessionClass) Alloc() OBEXSession {
	rv := objc.Send[OBEXSession](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// Object representing an OBEX connection to a remote target.
//
// # Overview
//
// You will have no need for a obtaining/using a raw OBEXSession, since it
// requires an underlying transport to do anything useful. However, once you
// have an object that is a subclass of this class, you can use the functions
// herein to manipulate that OBEXSession. First off, you will want to use
// OBEXConnect (if you are a client session) to actually cause the transport
// to open a connection to a remote target and establish an OBEX connection
// over it. From there you can issue more commands based on the responses from
// a server.
//
// If you are a server session, the first thing you should receive is an
// OBEXConnect command packet, and you will want to issue an
// OBEXConnectResponse packet, with your reesponse to that command (success,
// denied, bad request, etc.).
//
// You can use the session accessors to access certain information, such as
// the negotiated max packet length.
//
// If you wish to implement your own OBEXSession over a transport such as
// ethernet, you will need to see the end of the file to determine which
// functions to override, and what to pass to those functions.
//
// No timeout mechanism has been implemented so far for an OBEXSessions. If
// you need timeouts, you will need to implement them yourself. This is being
// explored for a future revision. However, be aware that the OBEX
// Specification does not explicitly require timeouts, so be sure you allow
// ample time for commands to complete, as some devices may be slow when
// sending large amounts of data.
//
// # Instance Methods
//
//   - [OBEXSession.ClientHandleIncomingData]: Tranport subclasses need to invoke this from their own data-receive handlers. For example, when data is received over a Bluetooth RFCOMM channel in the IOBluetoothOBEXSession, it in turn calls this to dispatch the data. If you do not handle this case, your server session will not work, guaranteed.
//   - [OBEXSession.CloseTransportConnection]: You must override this - it will be called when the transport connection should be shutdown.
//   - [OBEXSession.GetAvailableCommandPayloadLength]: Determine the maximum amount of data you can send in a particular command as an OBEX client session.
//   - [OBEXSession.GetAvailableCommandResponsePayloadLength]: Determine the maximum amount of data you can send in a particular command response as an OBEX server session.
//   - [OBEXSession.GetMaxPacketLength]: Gets current max packet length.
//   - [OBEXSession.HasOpenOBEXConnection]: Has a successful connect packet been sent and received? This API tells you so.
//   - [OBEXSession.HasOpenTransportConnection]: You must override this - it will be called periodically to determine if a transport connection is open or not.
//   - [OBEXSession.OBEXAbortOptionalHeadersLengthEventSelectorSelectorTargetRefCon]: Send an OBEX Abort command to the session’s target.
//   - [OBEXSession.OBEXAbortResponseOptionalHeadersOptionalHeadersLengthEventSelectorSelectorTargetRefCon]: Send an abort response to a session’s target.
//   - [OBEXSession.OBEXConnectMaxPacketLengthOptionalHeadersOptionalHeadersLengthEventSelectorSelectorTargetRefCon]: Initiate an OBEX connection to a device. Causes underlying transport (Bluetooth, et al) to attempt to connect to a remote device. After success, an OBEX connect packet is sent to establish the OBEX Connection.
//   - [OBEXSession.OBEXConnectResponseFlagsMaxPacketLengthOptionalHeadersOptionalHeadersLengthEventSelectorSelectorTargetRefCon]: Send a connect response to a session’s target.
//   - [OBEXSession.OBEXDisconnectOptionalHeadersLengthEventSelectorSelectorTargetRefCon]: Send an OBEX Disconnect command to the session’s target. THIS DOES NOT necessarily close the underlying transport connection. Deleting the session will ensure that closure.
//   - [OBEXSession.OBEXDisconnectResponseOptionalHeadersOptionalHeadersLengthEventSelectorSelectorTargetRefCon]: Send a disconnect response to a session’s target.
//   - [OBEXSession.OBEXGetHeadersHeadersLengthEventSelectorSelectorTargetRefCon]: Send an OBEX Get command to the session’s target.
//   - [OBEXSession.OBEXGetResponseOptionalHeadersOptionalHeadersLengthEventSelectorSelectorTargetRefCon]: Send a get response to a session’s target.
//   - [OBEXSession.OBEXPutHeadersDataHeadersDataLengthBodyDataBodyDataLengthEventSelectorSelectorTargetRefCon]: Send an OBEX Put command to the session’s target.
//   - [OBEXSession.OBEXPutResponseOptionalHeadersOptionalHeadersLengthEventSelectorSelectorTargetRefCon]: Send a put response to a session’s target.
//   - [OBEXSession.OBEXSetPathConstantsOptionalHeadersOptionalHeadersLengthEventSelectorSelectorTargetRefCon]: Send an OBEX SetPath command to the session’s target.
//   - [OBEXSession.OBEXSetPathResponseOptionalHeadersOptionalHeadersLengthEventSelectorSelectorTargetRefCon]: Send a set path response to a session’s target.
//   - [OBEXSession.OpenTransportConnectionSelectorTargetRefCon]: Opens a transport connection to a device. A Bluetooth connection is one example of a transport.
//   - [OBEXSession.SendDataToTransportDataLength]: You must override this to send data over your transport. This does nothing by default, it will return a kOBEXUnsupportedError.
//   - [OBEXSession.ServerHandleIncomingData]: Tranport subclasses need to invoke this from their own data-receive handlers. For example, when data is received over a Bluetooth RFCOMM channel in the IOBluetoothOBEXSession, it in turn calls this to dispatch the data. If you do not handle this case, your server session will not work, guaranteed.
//   - [OBEXSession.SetEventCallback]: Sets the C-API callback used when the session recieves data.
//   - [OBEXSession.SetEventRefCon]: Sets the C-API callback refCon used when the session recieves data.
//   - [OBEXSession.SetEventSelectorTargetRefCon]: Allow you to set a selector to be called when events occur on the OBEX session.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXSession
type OBEXSession struct {
	objectivec.Object
}

// OBEXSessionFromID constructs a [OBEXSession] from an objc.ID.
//
// Object representing an OBEX connection to a remote target.
func OBEXSessionFromID(id objc.ID) OBEXSession {
	return OBEXSession{objectivec.Object{ID: id}}
}

// NOTE: OBEXSession adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [OBEXSession] class.
//
// # Instance Methods
//
//   - [IOBEXSession.ClientHandleIncomingData]: Tranport subclasses need to invoke this from their own data-receive handlers. For example, when data is received over a Bluetooth RFCOMM channel in the IOBluetoothOBEXSession, it in turn calls this to dispatch the data. If you do not handle this case, your server session will not work, guaranteed.
//   - [IOBEXSession.CloseTransportConnection]: You must override this - it will be called when the transport connection should be shutdown.
//   - [IOBEXSession.GetAvailableCommandPayloadLength]: Determine the maximum amount of data you can send in a particular command as an OBEX client session.
//   - [IOBEXSession.GetAvailableCommandResponsePayloadLength]: Determine the maximum amount of data you can send in a particular command response as an OBEX server session.
//   - [IOBEXSession.GetMaxPacketLength]: Gets current max packet length.
//   - [IOBEXSession.HasOpenOBEXConnection]: Has a successful connect packet been sent and received? This API tells you so.
//   - [IOBEXSession.HasOpenTransportConnection]: You must override this - it will be called periodically to determine if a transport connection is open or not.
//   - [IOBEXSession.OBEXAbortOptionalHeadersLengthEventSelectorSelectorTargetRefCon]: Send an OBEX Abort command to the session’s target.
//   - [IOBEXSession.OBEXAbortResponseOptionalHeadersOptionalHeadersLengthEventSelectorSelectorTargetRefCon]: Send an abort response to a session’s target.
//   - [IOBEXSession.OBEXConnectMaxPacketLengthOptionalHeadersOptionalHeadersLengthEventSelectorSelectorTargetRefCon]: Initiate an OBEX connection to a device. Causes underlying transport (Bluetooth, et al) to attempt to connect to a remote device. After success, an OBEX connect packet is sent to establish the OBEX Connection.
//   - [IOBEXSession.OBEXConnectResponseFlagsMaxPacketLengthOptionalHeadersOptionalHeadersLengthEventSelectorSelectorTargetRefCon]: Send a connect response to a session’s target.
//   - [IOBEXSession.OBEXDisconnectOptionalHeadersLengthEventSelectorSelectorTargetRefCon]: Send an OBEX Disconnect command to the session’s target. THIS DOES NOT necessarily close the underlying transport connection. Deleting the session will ensure that closure.
//   - [IOBEXSession.OBEXDisconnectResponseOptionalHeadersOptionalHeadersLengthEventSelectorSelectorTargetRefCon]: Send a disconnect response to a session’s target.
//   - [IOBEXSession.OBEXGetHeadersHeadersLengthEventSelectorSelectorTargetRefCon]: Send an OBEX Get command to the session’s target.
//   - [IOBEXSession.OBEXGetResponseOptionalHeadersOptionalHeadersLengthEventSelectorSelectorTargetRefCon]: Send a get response to a session’s target.
//   - [IOBEXSession.OBEXPutHeadersDataHeadersDataLengthBodyDataBodyDataLengthEventSelectorSelectorTargetRefCon]: Send an OBEX Put command to the session’s target.
//   - [IOBEXSession.OBEXPutResponseOptionalHeadersOptionalHeadersLengthEventSelectorSelectorTargetRefCon]: Send a put response to a session’s target.
//   - [IOBEXSession.OBEXSetPathConstantsOptionalHeadersOptionalHeadersLengthEventSelectorSelectorTargetRefCon]: Send an OBEX SetPath command to the session’s target.
//   - [IOBEXSession.OBEXSetPathResponseOptionalHeadersOptionalHeadersLengthEventSelectorSelectorTargetRefCon]: Send a set path response to a session’s target.
//   - [IOBEXSession.OpenTransportConnectionSelectorTargetRefCon]: Opens a transport connection to a device. A Bluetooth connection is one example of a transport.
//   - [IOBEXSession.SendDataToTransportDataLength]: You must override this to send data over your transport. This does nothing by default, it will return a kOBEXUnsupportedError.
//   - [IOBEXSession.ServerHandleIncomingData]: Tranport subclasses need to invoke this from their own data-receive handlers. For example, when data is received over a Bluetooth RFCOMM channel in the IOBluetoothOBEXSession, it in turn calls this to dispatch the data. If you do not handle this case, your server session will not work, guaranteed.
//   - [IOBEXSession.SetEventCallback]: Sets the C-API callback used when the session recieves data.
//   - [IOBEXSession.SetEventRefCon]: Sets the C-API callback refCon used when the session recieves data.
//   - [IOBEXSession.SetEventSelectorTargetRefCon]: Allow you to set a selector to be called when events occur on the OBEX session.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXSession
type IOBEXSession interface {
	objectivec.IObject

	// Topic: Instance Methods

	// Tranport subclasses need to invoke this from their own data-receive handlers. For example, when data is received over a Bluetooth RFCOMM channel in the IOBluetoothOBEXSession, it in turn calls this to dispatch the data. If you do not handle this case, your server session will not work, guaranteed.
	ClientHandleIncomingData(event *OBEXTransportEvent)
	// You must override this - it will be called when the transport connection should be shutdown.
	CloseTransportConnection() OBEXError
	// Determine the maximum amount of data you can send in a particular command as an OBEX client session.
	GetAvailableCommandPayloadLength(inOpCode OBEXOpCode) OBEXMaxPacketLength
	// Determine the maximum amount of data you can send in a particular command response as an OBEX server session.
	GetAvailableCommandResponsePayloadLength(inOpCode OBEXOpCode) OBEXMaxPacketLength
	// Gets current max packet length.
	GetMaxPacketLength() OBEXMaxPacketLength
	// Has a successful connect packet been sent and received? This API tells you so.
	HasOpenOBEXConnection() bool
	// You must override this - it will be called periodically to determine if a transport connection is open or not.
	HasOpenTransportConnection() bool
	// Send an OBEX Abort command to the session’s target.
	OBEXAbortOptionalHeadersLengthEventSelectorSelectorTargetRefCon(inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon uintptr) OBEXError
	// Send an abort response to a session’s target.
	OBEXAbortResponseOptionalHeadersOptionalHeadersLengthEventSelectorSelectorTargetRefCon(inResponseOpCode OBEXOpCode, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon uintptr) OBEXError
	// Initiate an OBEX connection to a device. Causes underlying transport (Bluetooth, et al) to attempt to connect to a remote device. After success, an OBEX connect packet is sent to establish the OBEX Connection.
	OBEXConnectMaxPacketLengthOptionalHeadersOptionalHeadersLengthEventSelectorSelectorTargetRefCon(inFlags OBEXFlags, inMaxPacketLength OBEXMaxPacketLength, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon uintptr) OBEXError
	// Send a connect response to a session’s target.
	OBEXConnectResponseFlagsMaxPacketLengthOptionalHeadersOptionalHeadersLengthEventSelectorSelectorTargetRefCon(inResponseOpCode OBEXOpCode, inFlags OBEXFlags, inMaxPacketLength OBEXMaxPacketLength, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon uintptr) OBEXError
	// Send an OBEX Disconnect command to the session’s target. THIS DOES NOT necessarily close the underlying transport connection. Deleting the session will ensure that closure.
	OBEXDisconnectOptionalHeadersLengthEventSelectorSelectorTargetRefCon(inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon uintptr) OBEXError
	// Send a disconnect response to a session’s target.
	OBEXDisconnectResponseOptionalHeadersOptionalHeadersLengthEventSelectorSelectorTargetRefCon(inResponseOpCode OBEXOpCode, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon uintptr) OBEXError
	// Send an OBEX Get command to the session’s target.
	OBEXGetHeadersHeadersLengthEventSelectorSelectorTargetRefCon(isFinalChunk bool, inHeaders unsafe.Pointer, inHeadersLength uintptr, inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon uintptr) OBEXError
	// Send a get response to a session’s target.
	OBEXGetResponseOptionalHeadersOptionalHeadersLengthEventSelectorSelectorTargetRefCon(inResponseOpCode OBEXOpCode, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon uintptr) OBEXError
	// Send an OBEX Put command to the session’s target.
	OBEXPutHeadersDataHeadersDataLengthBodyDataBodyDataLengthEventSelectorSelectorTargetRefCon(isFinalChunk bool, inHeadersData unsafe.Pointer, inHeadersDataLength uintptr, inBodyData unsafe.Pointer, inBodyDataLength uintptr, inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon uintptr) OBEXError
	// Send a put response to a session’s target.
	OBEXPutResponseOptionalHeadersOptionalHeadersLengthEventSelectorSelectorTargetRefCon(inResponseOpCode OBEXOpCode, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon uintptr) OBEXError
	// Send an OBEX SetPath command to the session’s target.
	OBEXSetPathConstantsOptionalHeadersOptionalHeadersLengthEventSelectorSelectorTargetRefCon(inFlags OBEXFlags, inConstants OBEXConstants, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon uintptr) OBEXError
	// Send a set path response to a session’s target.
	OBEXSetPathResponseOptionalHeadersOptionalHeadersLengthEventSelectorSelectorTargetRefCon(inResponseOpCode OBEXOpCode, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon uintptr) OBEXError
	// Opens a transport connection to a device. A Bluetooth connection is one example of a transport.
	OpenTransportConnectionSelectorTargetRefCon(inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon uintptr) OBEXError
	// You must override this to send data over your transport. This does nothing by default, it will return a kOBEXUnsupportedError.
	SendDataToTransportDataLength(inDataToSend unsafe.Pointer, inDataLength uintptr) OBEXError
	// Tranport subclasses need to invoke this from their own data-receive handlers. For example, when data is received over a Bluetooth RFCOMM channel in the IOBluetoothOBEXSession, it in turn calls this to dispatch the data. If you do not handle this case, your server session will not work, guaranteed.
	ServerHandleIncomingData(event *OBEXTransportEvent)
	// Sets the C-API callback used when the session recieves data.
	SetEventCallback(inEventCallback OBEXSessionEventCallback)
	// Sets the C-API callback refCon used when the session recieves data.
	SetEventRefCon(inRefCon unsafe.Pointer)
	// Allow you to set a selector to be called when events occur on the OBEX session.
	SetEventSelectorTargetRefCon(inEventSelector objc.SEL, inEventSelectorTarget objectivec.IObject, inUserRefCon uintptr)
}

// Init initializes the instance.
func (o OBEXSession) Init() OBEXSession {
	rv := objc.Send[OBEXSession](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o OBEXSession) Autorelease() OBEXSession {
	rv := objc.Send[OBEXSession](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewOBEXSession creates a new OBEXSession instance.
func NewOBEXSession() OBEXSession {
	class := getOBEXSessionClass()
	rv := objc.Send[OBEXSession](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Tranport subclasses need to invoke this from their own data-receive
// handlers. For example, when data is received over a Bluetooth RFCOMM
// channel in the IOBluetoothOBEXSession, it in turn calls this to dispatch
// the data. If you do not handle this case, your server session will not
// work, guaranteed.
//
// event: New event received from the transport.
//
// # Discussion
//
// Tranport subclasses must call this for OBEX client sessions to work!
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/clientHandleIncomingData(_:)
func (o OBEXSession) ClientHandleIncomingData(event *OBEXTransportEvent) {
	objc.Send[objc.ID](o.ID, objc.Sel("clientHandleIncomingData:"), unsafe.Pointer(event))
}

// You must override this - it will be called when the transport connection
// should be shutdown.
//
// # Return Value
//
// Return whether or not the transport connection was closed successfully or
// not. Return OBEXSuccess ( 0 ) on success, otherwise an error code.
//
// # Discussion
//
// Tranport subclasses must override this! When called you should take
// whatever steps are necessary to actually close down the transport
// connection.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/closeTransportConnection()
func (o OBEXSession) CloseTransportConnection() OBEXError {
	rv := objc.Send[OBEXError](o.ID, objc.Sel("closeTransportConnection"))
	return OBEXError(rv)
}

// Determine the maximum amount of data you can send in a particular command
// as an OBEX client session.
//
// inOpCode: The opcode you are interested in sending (as a client).
//
// # Return Value
//
// The maximum amount of data a particular packet can handle, after accounting
// for any command overhead.
//
// # Discussion
//
// Each OBEX Command has a certain amount of overhead. Since the negotiated
// max packet length does not indicate what the maximum data amount you can
// send in a particular command’s packet, you can use this function to
// determine how much data to provide in optional headers or body data
// headers.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/getAvailableCommandPayloadLength(_:)
func (o OBEXSession) GetAvailableCommandPayloadLength(inOpCode OBEXOpCode) OBEXMaxPacketLength {
	rv := objc.Send[OBEXMaxPacketLength](o.ID, objc.Sel("getAvailableCommandPayloadLength:"), inOpCode)
	return OBEXMaxPacketLength(rv)
}

// Determine the maximum amount of data you can send in a particular command
// response as an OBEX server session.
//
// inOpCode: The opcode you are interested in responding to (as a server).
//
// # Return Value
//
// The maximum amount of data a particular packet can handle, after accounting
// for any command response overhead.
//
// # Discussion
//
// Each OBEX Command response has a certain amount of overhead. Since the
// negotiated max packet length does not indicate what the maximum data amount
// you can send in a particular response’s packet, you can use this function
// to determine how much data to provide in optional headers or body data
// headers.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/getAvailableCommandResponsePayloadLength(_:)
func (o OBEXSession) GetAvailableCommandResponsePayloadLength(inOpCode OBEXOpCode) OBEXMaxPacketLength {
	rv := objc.Send[OBEXMaxPacketLength](o.ID, objc.Sel("getAvailableCommandResponsePayloadLength:"), inOpCode)
	return OBEXMaxPacketLength(rv)
}

// Gets current max packet length.
//
// # Return Value
//
// Max packet length.
//
// # Discussion
//
// This value could change before and after a connect command has been sent or
// a connect command response has been received, since the recipient could
// negotiate a lower max packet size.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/getMaxPacketLength()
func (o OBEXSession) GetMaxPacketLength() OBEXMaxPacketLength {
	rv := objc.Send[OBEXMaxPacketLength](o.ID, objc.Sel("getMaxPacketLength"))
	return OBEXMaxPacketLength(rv)
}

// Has a successful connect packet been sent and received? This API tells you
// so.
//
// # Return Value
//
// True or false, we are OBEX-connected to another OBEX entity.
//
// # Discussion
//
// A “transport” connection may exist (such as a Bluetooth baseband
// connection), but the OBEX connection may not be established over that
// transport. If it has been, this function returns true.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/hasOpenOBEXConnection()
func (o OBEXSession) HasOpenOBEXConnection() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("hasOpenOBEXConnection"))
	return rv
}

// You must override this - it will be called periodically to determine if a
// transport connection is open or not.
//
// # Return Value
//
// Return whether or not the transport connection is still open.
//
// # Discussion
//
// Tranport subclasses must override this! When called you simply return if
// the transport connection is still open or not.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/hasOpenTransportConnection()
func (o OBEXSession) HasOpenTransportConnection() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("hasOpenTransportConnection"))
	return rv
}

// Send an OBEX Abort command to the session’s target.
//
// inOptionalHeaders: Can be NULL. Ptr to some data you want to send as your optional headers.
// Use the provided header contruction kit in OBEX.h and OBEXHeadersToBytes()
// for convenience.
//
// inOptionalHeadersLength: Length of data in ptr passed in above.
//
// inSelector: A VALID selector to be called when something interesting happens due to
// this call. Selector in your target object MUST have the following
// signature, or it will not be called properly (look for error messages in
// Console.app):
//
// - (void)OBEXAbortHandler:(const OBEXSessionEvent*)inSessionEvent;
//
// inTarget: A VALID target object for the selector.
//
// inUserRefCon: Whatever you want to pass here. It will be passed back to you in the refCon
// portion of the OBEXSessionEvent struct. nil is, of course, OK here.
//
// # Discussion
//
// A NULL selector or target will result in an error. After return, the data
// passed in will have been sent over the transport. You will receive a
// response to your command on your selector.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/obexAbort(_:optionalHeadersLength:eventSelector:selectorTarget:refCon:)
func (o OBEXSession) OBEXAbortOptionalHeadersLengthEventSelectorSelectorTargetRefCon(inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon uintptr) OBEXError {
	rv := objc.Send[OBEXError](o.ID, objc.Sel("OBEXAbort:optionalHeadersLength:eventSelector:selectorTarget:refCon:"), inOptionalHeaders, inOptionalHeadersLength, inSelector, inTarget, inUserRefCon)
	return OBEXError(rv)
}

// Send an abort response to a session’s target.
//
// inOptionalHeaders: Can be NULL. Ptr to some data you want to send as your optional headers.
// Use the provided header contruction kit in OBEX.h and OBEXHeadersToBytes()
// for convenience.
//
// inOptionalHeadersLength: Length of data in ptr passed in above.
//
// inSelector: A VALID selector to be called when something interesting happens due to
// this call. Selector in your target object MUST have the following
// signature, or it will not be called properly (look for error messages in
// Console.app):
//
// `- (void)OBEXAbortResponseHandler:(const OBEXSessionEvent*)inSessionEvent;`
//
// inTarget: A VALID target object for the selector.
//
// inUserRefCon: Whatever you want to pass here. It will be passed back to you in the refCon
// portion of the OBEXSessionEvent struct. nil is, of course, OK here.
//
// # Discussion
//
// A NULL selector or target will result in an error. After return, the data
// passed in will have been sent over the underlying OBEX transport. You will
// receive any responses to your command response on your selector.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/obexAbortResponse(_:optionalHeaders:optionalHeadersLength:eventSelector:selectorTarget:refCon:)
func (o OBEXSession) OBEXAbortResponseOptionalHeadersOptionalHeadersLengthEventSelectorSelectorTargetRefCon(inResponseOpCode OBEXOpCode, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon uintptr) OBEXError {
	rv := objc.Send[OBEXError](o.ID, objc.Sel("OBEXAbortResponse:optionalHeaders:optionalHeadersLength:eventSelector:selectorTarget:refCon:"), inResponseOpCode, inOptionalHeaders, inOptionalHeadersLength, inSelector, inTarget, inUserRefCon)
	return OBEXError(rv)
}

// Initiate an OBEX connection to a device. Causes underlying transport
// (Bluetooth, et al) to attempt to connect to a remote device. After success,
// an OBEX connect packet is sent to establish the OBEX Connection.
//
// inFlags: OBEX connect flags. See OBEX.h for possibilities.
//
// inMaxPacketLength: Maximum packet size you want to support. May be negotiated down, depending
// on target device.
//
// inOptionalHeaders: Can be NULL. Ptr to some data you want to send as your optional headers.
// Use the provided header contruction kit in OBEX.h and OBEXHeadersToBytes()
// for convenience.
//
// inOptionalHeadersLength: Length of data in ptr passed in above.
//
// inSelector: A VALID selector to be called when something interesting happens due to
// this call. Selector in your target object MUST have the following
// signature, or it will not be called properly (look for error messages in
// Console.app):
//
// - (void)OBEXConnectHandler:(const OBEXSessionEvent*)inSessionEvent;
//
// inTarget: A VALID target object for the selector.
//
// inUserRefCon: Whatever you want to pass here. It will be passed back to you in the refCon
// portion of the OBEXSessionEvent struct. nil is, of course, OK here.
//
// # Discussion
//
// A NULL selector or target will result in an error. After return, the data
// passed in will have been sent over the transport. You will receive a
// response to your command on your selector. If you have already established
// an OBEX connection and you call this again you will get an
// ‘kOBEXSessionAlreadyConnectedError’ as a result.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/obexConnect(_:maxPacketLength:optionalHeaders:optionalHeadersLength:eventSelector:selectorTarget:refCon:)
func (o OBEXSession) OBEXConnectMaxPacketLengthOptionalHeadersOptionalHeadersLengthEventSelectorSelectorTargetRefCon(inFlags OBEXFlags, inMaxPacketLength OBEXMaxPacketLength, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon uintptr) OBEXError {
	rv := objc.Send[OBEXError](o.ID, objc.Sel("OBEXConnect:maxPacketLength:optionalHeaders:optionalHeadersLength:eventSelector:selectorTarget:refCon:"), inFlags, inMaxPacketLength, inOptionalHeaders, inOptionalHeadersLength, inSelector, inTarget, inUserRefCon)
	return OBEXError(rv)
}

// Send a connect response to a session’s target.
//
// inResponseOpCode: OBEX response constant. See OBEX.h for possibilities.
//
// inFlags: OBEX connect flags. See OBEX.h for possibilities.
//
// inMaxPacketLength: Maximum packet size you want your OBEX session to communicate with. This
// MUST be lower than the max packet size the client has reported to you in
// the connect command you received from it.
//
// inOptionalHeaders: Can be NULL. Ptr to some data you want to send as your optional headers.
// Use the provided header contruction kit in OBEX.h and OBEXHeadersToBytes()
// for convenience.
//
// inOptionalHeadersLength: Length of data in ptr passed in above.
//
// inSelector: A VALID selector to be called when something interesting happens due to
// this call. Selector in your target object MUST have the following
// signature, or it will not be called properly (look for error messages in
// Console.app):
//
// - (void)OBEXConnectResponseHandler:(const OBEXSessionEvent*)inSessionEvent;
//
// inTarget: A VALID target object for the selector.
//
// inUserRefCon: Whatever you want to pass here. It will be passed back to you in the refCon
// portion of the OBEXSessionEvent struct. nil is, of course, OK here.
//
// # Discussion
//
// A NULL selector or target will result in an error. After return, the data
// passed in will have been sent over the underlying OBEX transport. You will
// receive any responses to your command response on your selector.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/obexConnectResponse(_:flags:maxPacketLength:optionalHeaders:optionalHeadersLength:eventSelector:selectorTarget:refCon:)
func (o OBEXSession) OBEXConnectResponseFlagsMaxPacketLengthOptionalHeadersOptionalHeadersLengthEventSelectorSelectorTargetRefCon(inResponseOpCode OBEXOpCode, inFlags OBEXFlags, inMaxPacketLength OBEXMaxPacketLength, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon uintptr) OBEXError {
	rv := objc.Send[OBEXError](o.ID, objc.Sel("OBEXConnectResponse:flags:maxPacketLength:optionalHeaders:optionalHeadersLength:eventSelector:selectorTarget:refCon:"), inResponseOpCode, inFlags, inMaxPacketLength, inOptionalHeaders, inOptionalHeadersLength, inSelector, inTarget, inUserRefCon)
	return OBEXError(rv)
}

// Send an OBEX Disconnect command to the session’s target. THIS DOES NOT
// necessarily close the underlying transport connection. Deleting the session
// will ensure that closure.
//
// inOptionalHeaders: Can be NULL. Ptr to some data you want to send as your optional headers.
// Use the provided header contruction kit in OBEX.h and OBEXHeadersToBytes()
// for convenience.
//
// inOptionalHeadersLength: Length of data in ptr passed in above.
//
// inSelector: A VALID selector to be called when something interesting happens due to
// this call. Selector in your target object MUST have the following
// signature, or it will not be called properly (look for error messages in
// Console.app):
//
// - (void)OBEXDisconnectHandler:(const OBEXSessionEvent*)inSessionEvent;
//
// inTarget: A VALID target object for the selector.
//
// inUserRefCon: Whatever you want to pass here. It will be passed back to you in the refCon
// portion of the OBEXSessionEvent struct. nil is, of course, OK here.
//
// # Discussion
//
// A NULL selector or target will result in an error. After return, the data
// passed in will have been sent over the transport. You will receive a
// response to your command on your selector. Be careful not to exceed the max
// packet length in your optional headers, or your command will be rejected.
// It is recommended that you call getMaxPacketLength on your session before
// issuing this command so you know how much data the session’s target will
// accept in a single transaction.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/obexDisconnect(_:optionalHeadersLength:eventSelector:selectorTarget:refCon:)
func (o OBEXSession) OBEXDisconnectOptionalHeadersLengthEventSelectorSelectorTargetRefCon(inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon uintptr) OBEXError {
	rv := objc.Send[OBEXError](o.ID, objc.Sel("OBEXDisconnect:optionalHeadersLength:eventSelector:selectorTarget:refCon:"), inOptionalHeaders, inOptionalHeadersLength, inSelector, inTarget, inUserRefCon)
	return OBEXError(rv)
}

// Send a disconnect response to a session’s target.
//
// inOptionalHeaders: Can be NULL. Ptr to some data you want to send as your optional headers.
// Use the provided header contruction kit in OBEX.h and OBEXHeadersToBytes()
// for convenience.
//
// inOptionalHeadersLength: Length of data in ptr passed in above.
//
// inSelector: A VALID selector to be called when something interesting happens due to
// this call. Selector in your target object MUST have the following
// signature, or it will not be called properly (look for error messages in
// Console.app):
//
// - (void)OBEXDisconnectResponseHandler:(const
// OBEXSessionEvent*)inSessionEvent;
//
// inTarget: A VALID target object for the selector.
//
// inUserRefCon: Whatever you want to pass here. It will be passed back to you in the refCon
// portion of the OBEXSessionEvent struct. nil is, of course, OK here.
//
// # Discussion
//
// A NULL selector or target will result in an error. After return, the data
// passed in will have been sent over the underlying OBEX transport. You will
// receive any responses to your command response on your selector.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/obexDisconnectResponse(_:optionalHeaders:optionalHeadersLength:eventSelector:selectorTarget:refCon:)
func (o OBEXSession) OBEXDisconnectResponseOptionalHeadersOptionalHeadersLengthEventSelectorSelectorTargetRefCon(inResponseOpCode OBEXOpCode, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon uintptr) OBEXError {
	rv := objc.Send[OBEXError](o.ID, objc.Sel("OBEXDisconnectResponse:optionalHeaders:optionalHeadersLength:eventSelector:selectorTarget:refCon:"), inResponseOpCode, inOptionalHeaders, inOptionalHeadersLength, inSelector, inTarget, inUserRefCon)
	return OBEXError(rv)
}

// Send an OBEX Get command to the session’s target.
//
// isFinalChunk: Specify if this request is complete in one shot - that all the headers you
// are supplying will fit in the negotiated max packet length.
//
// inHeaders: Can be NULL. Ptr to some data you want to send as your headers, such as
// Length, Name, etc. Use the provided header contruction kit in OBEX.h and
// OBEXHeadersToBytes() for your convenience.
//
// inHeadersLength: Length of data in ptr passed in above.
//
// inSelector: A VALID selector to be called when something interesting happens due to
// this call. Selector in your target object MUST have the following
// signature, or it will not be called properly (look for error messages in
// Console.app):
//
// - (void)OBEXGetHandler:(const OBEXSessionEvent*)inSessionEvent;
//
// inTarget: A VALID target object for the selector.
//
// inUserRefCon: Whatever you want to pass here. It will be passed back to you in the refCon
// portion of the OBEXSessionEvent struct. nil is, of course, OK here.
//
// # Discussion
//
// A NULL selector or target will result in an error. After return, the data
// passed in will have been sent over the transport. You will receive a
// response to your command on your selector.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/obexGet(_:headers:headersLength:eventSelector:selectorTarget:refCon:)
func (o OBEXSession) OBEXGetHeadersHeadersLengthEventSelectorSelectorTargetRefCon(isFinalChunk bool, inHeaders unsafe.Pointer, inHeadersLength uintptr, inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon uintptr) OBEXError {
	rv := objc.Send[OBEXError](o.ID, objc.Sel("OBEXGet:headers:headersLength:eventSelector:selectorTarget:refCon:"), isFinalChunk, inHeaders, inHeadersLength, inSelector, inTarget, inUserRefCon)
	return OBEXError(rv)
}

// Send a get response to a session’s target.
//
// inOptionalHeaders: Can be NULL. Ptr to some data you want to send as your optional headers.
// Use the provided header contruction kit in OBEX.h and OBEXHeadersToBytes()
// for convenience.
//
// inOptionalHeadersLength: Length of data in ptr passed in above.
//
// inSelector: A VALID selector to be called when something interesting happens due to
// this call. Selector in your target object MUST have the following
// signature, or it will not be called properly (look for error messages in
// Console.app):
//
// - (void)OBEXGetResponseHandler:(const OBEXSessionEvent*)inSessionEvent;
//
// inTarget: A VALID target object for the selector.
//
// inUserRefCon: Whatever you want to pass here. It will be passed back to you in the refCon
// portion of the OBEXSessionEvent struct. nil is, of course, OK here.
//
// # Discussion
//
// A NULL selector or target will result in an error. After return, the data
// passed in will have been sent over the underlying OBEX transport. You will
// receive any responses to your command response on your selector.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/obexGetResponse(_:optionalHeaders:optionalHeadersLength:eventSelector:selectorTarget:refCon:)
func (o OBEXSession) OBEXGetResponseOptionalHeadersOptionalHeadersLengthEventSelectorSelectorTargetRefCon(inResponseOpCode OBEXOpCode, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon uintptr) OBEXError {
	rv := objc.Send[OBEXError](o.ID, objc.Sel("OBEXGetResponse:optionalHeaders:optionalHeadersLength:eventSelector:selectorTarget:refCon:"), inResponseOpCode, inOptionalHeaders, inOptionalHeadersLength, inSelector, inTarget, inUserRefCon)
	return OBEXError(rv)
}

// Send an OBEX Put command to the session’s target.
//
// isFinalChunk: Specify if this request is complete in one shot - that all the headers you
// are supplying will fit in the negotiated max packet length.
//
// inHeadersData: Can be NULL. Ptr to some data you want to send as your headers, such as
// Length, Name, etc. Use the provided header contruction kit in OBEX.h and
// OBEXHeadersToBytes() for convenience.
//
// inHeadersDataLength: Length of data in ptr passed in above.
//
// inBodyData: Can be NULL. Ptr to some data you want to send as your BODY header. Do not
// construct a real OBEX header here, it will be done for you - just pass a
// pointer to your data, we’ll do the rest. HOWEVER, be aware that some
// overhead (3 bytes) will be added to the data in constructing the BODY
// header for you.
//
// inBodyDataLength: Length of data in ptr passed in above.
//
// inSelector: A VALID selector to be called when something interesting happens due to
// this call. Selector in your target object MUST have the following
// signature, or it will not be called properly (look for error messages in
// Console.app):
//
// - (void)OBEXPutHandler:(const OBEXSessionEvent*)inSessionEvent;
//
// inTarget: A VALID target object for the selector.
//
// inUserRefCon: Whatever you want to pass here. It will be passed back to you in the refCon
// portion of the OBEXSessionEvent struct. nil is, of course, OK here.
//
// # Discussion
//
// A NULL selector or target will result in an error. After return, the data
// passed in will have been sent over the transport. You will receive a
// response to your command on your selector.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/obexPut(_:headersData:headersDataLength:bodyData:bodyDataLength:eventSelector:selectorTarget:refCon:)
func (o OBEXSession) OBEXPutHeadersDataHeadersDataLengthBodyDataBodyDataLengthEventSelectorSelectorTargetRefCon(isFinalChunk bool, inHeadersData unsafe.Pointer, inHeadersDataLength uintptr, inBodyData unsafe.Pointer, inBodyDataLength uintptr, inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon uintptr) OBEXError {
	rv := objc.Send[OBEXError](o.ID, objc.Sel("OBEXPut:headersData:headersDataLength:bodyData:bodyDataLength:eventSelector:selectorTarget:refCon:"), isFinalChunk, inHeadersData, inHeadersDataLength, inBodyData, inBodyDataLength, inSelector, inTarget, inUserRefCon)
	return OBEXError(rv)
}

// Send a put response to a session’s target.
//
// inOptionalHeaders: Can be NULL. Ptr to some data you want to send as your optional headers.
// Use the provided header contruction kit in OBEX.h and OBEXHeadersToBytes()
// for convenience.
//
// inOptionalHeadersLength: Length of data in ptr passed in above.
//
// inSelector: A VALID selector to be called when something interesting happens due to
// this call. Selector in your target object MUST have the following
// signature, or it will not be called properly (look for error messages in
// Console.app):
//
// - (void)OBEXPutResponseHandler:(const OBEXSessionEvent*)inSessionEvent;
//
// inTarget: A VALID target object for the selector.
//
// inUserRefCon: Whatever you want to pass here. It will be passed back to you in the refCon
// portion of the OBEXSessionEvent struct. nil is, of course, OK here.
//
// # Discussion
//
// A NULL selector or target will result in an error. After return, the data
// passed in will have been sent over the underlying OBEX transport. You will
// receive any responses to your command response on your selector.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/obexPutResponse(_:optionalHeaders:optionalHeadersLength:eventSelector:selectorTarget:refCon:)
func (o OBEXSession) OBEXPutResponseOptionalHeadersOptionalHeadersLengthEventSelectorSelectorTargetRefCon(inResponseOpCode OBEXOpCode, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon uintptr) OBEXError {
	rv := objc.Send[OBEXError](o.ID, objc.Sel("OBEXPutResponse:optionalHeaders:optionalHeadersLength:eventSelector:selectorTarget:refCon:"), inResponseOpCode, inOptionalHeaders, inOptionalHeadersLength, inSelector, inTarget, inUserRefCon)
	return OBEXError(rv)
}

// Send an OBEX SetPath command to the session’s target.
//
// inFlags: OBEX setpath flags. See OBEX.h for possibilities.
//
// inConstants: OBEX setpath constants. See OBEX.h for possibilities.
//
// inOptionalHeaders: Can be NULL. Ptr to some data you want to send as your optional headers.
// Use the provided header contruction kit in OBEX.h and OBEXHeadersToBytes()
// for convenience.
//
// inOptionalHeadersLength: Length of data in ptr passed in above.
//
// inSelector: A VALID selector to be called when something interesting happens due to
// this call. Selector in your target object MUST have the following
// signature, or it will not be called properly (look for error messages in
// Console.app):
//
// - (void)OBEXSetPathHandler:(const OBEXSessionEvent*)inSessionEvent;
//
// inTarget: A VALID target object for the selector.
//
// inUserRefCon: Whatever you want to pass here. It will be passed back to you in the refCon
// portion of the OBEXSessionEvent struct. nil is, of course, OK here.
//
// # Discussion
//
// A NULL selector or target will result in an error. After return, the data
// passed in will have been sent over the transport. You will receive a
// response to your command on your selector.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/obexSetPath(_:constants:optionalHeaders:optionalHeadersLength:eventSelector:selectorTarget:refCon:)
func (o OBEXSession) OBEXSetPathConstantsOptionalHeadersOptionalHeadersLengthEventSelectorSelectorTargetRefCon(inFlags OBEXFlags, inConstants OBEXConstants, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon uintptr) OBEXError {
	rv := objc.Send[OBEXError](o.ID, objc.Sel("OBEXSetPath:constants:optionalHeaders:optionalHeadersLength:eventSelector:selectorTarget:refCon:"), inFlags, inConstants, inOptionalHeaders, inOptionalHeadersLength, inSelector, inTarget, inUserRefCon)
	return OBEXError(rv)
}

// Send a set path response to a session’s target.
//
// inOptionalHeaders: Can be NULL. Ptr to some data you want to send as your optional headers.
// Use the provided header contruction kit in OBEX.h and OBEXHeadersToBytes()
// for convenience.
//
// inOptionalHeadersLength: Length of data in ptr passed in above.
//
// inSelector: A VALID selector to be called when something interesting happens due to
// this call. Selector in your target object MUST have the following
// signature, or it will not be called properly (look for error messages in
// Console.app):
//
// - (void)OBEXSetPathResponseHandler:(const OBEXSessionEvent*)inSessionEvent;
//
// inTarget: A VALID target object for the selector.
//
// inUserRefCon: Whatever you want to pass here. It will be passed back to you in the refCon
// portion of the OBEXSessionEvent struct. nil is, of course, OK here.
//
// # Discussion
//
// A NULL selector or target will result in an error. After return, the data
// passed in will have been sent over the underlying OBEX transport. You will
// receive any responses to your command response on your selector.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/obexSetPathResponse(_:optionalHeaders:optionalHeadersLength:eventSelector:selectorTarget:refCon:)
func (o OBEXSession) OBEXSetPathResponseOptionalHeadersOptionalHeadersLengthEventSelectorSelectorTargetRefCon(inResponseOpCode OBEXOpCode, inOptionalHeaders unsafe.Pointer, inOptionalHeadersLength uintptr, inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon uintptr) OBEXError {
	rv := objc.Send[OBEXError](o.ID, objc.Sel("OBEXSetPathResponse:optionalHeaders:optionalHeadersLength:eventSelector:selectorTarget:refCon:"), inResponseOpCode, inOptionalHeaders, inOptionalHeadersLength, inSelector, inTarget, inUserRefCon)
	return OBEXError(rv)
}

// Opens a transport connection to a device. A Bluetooth connection is one
// example of a transport.
//
// inSelector: Selector to call for success, failure or timeout.
//
// inTarget: Target on which to call the selector.
//
// inUserRefCon: Caller’s reference constant.
//
// # Discussion
//
// Tranport subclasses must override this! when called you should attempt to
// open your transport connection, and if you are successful, return
// kOBEXSuccess, otherwise an interesting error code.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/openTransportConnection(_:selectorTarget:refCon:)
func (o OBEXSession) OpenTransportConnectionSelectorTargetRefCon(inSelector objc.SEL, inTarget objectivec.IObject, inUserRefCon uintptr) OBEXError {
	rv := objc.Send[OBEXError](o.ID, objc.Sel("openTransportConnection:selectorTarget:refCon:"), inSelector, inTarget, inUserRefCon)
	return OBEXError(rv)
}

// You must override this to send data over your transport. This does nothing
// by default, it will return a kOBEXUnsupportedError.
//
// inDataToSend: Data to shove over the transport to a remote OBEX session.
//
// inDataLength: Length of data passed in.
//
// # Discussion
//
// Tranport subclasses must override this! When called you should send the
// data over the transport to the remote session.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/sendData(toTransport:dataLength:)
func (o OBEXSession) SendDataToTransportDataLength(inDataToSend unsafe.Pointer, inDataLength uintptr) OBEXError {
	rv := objc.Send[OBEXError](o.ID, objc.Sel("sendDataToTransport:dataLength:"), inDataToSend, inDataLength)
	return OBEXError(rv)
}

// Tranport subclasses need to invoke this from their own data-receive
// handlers. For example, when data is received over a Bluetooth RFCOMM
// channel in the IOBluetoothOBEXSession, it in turn calls this to dispatch
// the data. If you do not handle this case, your server session will not
// work, guaranteed.
//
// event: New event received from the transport.
//
// # Discussion
//
// Tranport subclasses must call this for OBEX server sessions to work!
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/serverHandleIncomingData(_:)
func (o OBEXSession) ServerHandleIncomingData(event *OBEXTransportEvent) {
	objc.Send[objc.ID](o.ID, objc.Sel("serverHandleIncomingData:"), unsafe.Pointer(event))
}

// Sets the C-API callback used when the session recieves data.
//
// inEventCallback: Function to callback. Should be non-NULL, unless you are attempting to
// clear the callback, but doing that doesn’t make much sense.
//
// # Discussion
//
// This is really not intended for client sessions. Only subclasses would
// really be interested in using this. They should set these when their
// subclass object is created, because otherwise they will have no way of
// receiving the initial command data packet. This is a partner to
// setEventRefCon, described below.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/setEventCallback(_:)
func (o OBEXSession) SetEventCallback(inEventCallback OBEXSessionEventCallback) {
	objc.Send[objc.ID](o.ID, objc.Sel("setEventCallback:"), inEventCallback)
}

// Sets the C-API callback refCon used when the session recieves data.
//
// inRefCon: User’s refCon that will get passed when their event callback is invoked.
//
// # Discussion
//
// This is really not intended for client sessions. Only subclasses would
// really be interested in using this. They should set these when their
// subclass object is created, because otherwise they will have no context in
// their callback.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/setEventRefCon(_:)
func (o OBEXSession) SetEventRefCon(inRefCon unsafe.Pointer) {
	objc.Send[objc.ID](o.ID, objc.Sel("setEventRefCon:"), inRefCon)
}

// Allow you to set a selector to be called when events occur on the OBEX
// session.
//
// inEventSelector: Selector to call on the target.
//
// inEventSelectorTarget: Target to be called with the selector.
//
// inUserRefCon: User’s refCon that will get passed when their event callback is invoked.
//
// # Discussion
//
// Really not needed to be used, since the event selector will get set when an
// OBEX command is sent out.
//
// See: https://developer.apple.com/documentation/IOBluetooth/OBEXSession/setEventSelector(_:target:refCon:)
func (o OBEXSession) SetEventSelectorTargetRefCon(inEventSelector objc.SEL, inEventSelectorTarget objectivec.IObject, inUserRefCon uintptr) {
	objc.Send[objc.ID](o.ID, objc.Sel("setEventSelector:target:refCon:"), inEventSelector, inEventSelectorTarget, inUserRefCon)
}
