// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"context"
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NETunnelProviderSession] class.
var (
	_NETunnelProviderSessionClass     NETunnelProviderSessionClass
	_NETunnelProviderSessionClassOnce sync.Once
)

func getNETunnelProviderSessionClass() NETunnelProviderSessionClass {
	_NETunnelProviderSessionClassOnce.Do(func() {
		_NETunnelProviderSessionClass = NETunnelProviderSessionClass{class: objc.GetClass("NETunnelProviderSession")}
	})
	return _NETunnelProviderSessionClass
}

// GetNETunnelProviderSessionClass returns the class object for NETunnelProviderSession.
func GetNETunnelProviderSessionClass() NETunnelProviderSessionClass {
	return getNETunnelProviderSessionClass()
}

type NETunnelProviderSessionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NETunnelProviderSessionClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NETunnelProviderSessionClass) Alloc() NETunnelProviderSession {
	rv := objc.Send[NETunnelProviderSession](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object to start and stop a tunnel connection and get its status.
//
// # Overview
//
// [NETunnelProviderSession] objects control network tunnel connections
// provided by Tunnel Provider extensions.
//
// [NETunnelProviderSession] objects are not instantiated directly. Instead,
// each [NETunnelProviderManager] object has an associated
// [NETunnelProviderSession] as a read-only property.
//
// # Controlling the tunnel connection
//
//   - [NETunnelProviderSession.StartTunnelWithOptionsAndReturnError]: Start the process of connecting the tunnel.
//   - [NETunnelProviderSession.StopTunnel]: Start the process of disconnecting the tunnel.
//
// # Communicating with the tunnel provider
//
//   - [NETunnelProviderSession.SendProviderMessageReturnErrorResponseHandler]: Send a message to the Tunnel Provider extension. If the extension is not running, it should be launched to handle the message. If this method can’t start sending the message it reports an error in the `returnError` parameter. If an error occurs while sending the message or returning the result, `nil` should be sent to the response handler as notification.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderSession
type NETunnelProviderSession struct {
	NEVPNConnection
}

// NETunnelProviderSessionFromID constructs a [NETunnelProviderSession] from an objc.ID.
//
// An object to start and stop a tunnel connection and get its status.
func NETunnelProviderSessionFromID(id objc.ID) NETunnelProviderSession {
	return NETunnelProviderSession{NEVPNConnection: NEVPNConnectionFromID(id)}
}

// NOTE: NETunnelProviderSession adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NETunnelProviderSession] class.
//
// # Controlling the tunnel connection
//
//   - [INETunnelProviderSession.StartTunnelWithOptionsAndReturnError]: Start the process of connecting the tunnel.
//   - [INETunnelProviderSession.StopTunnel]: Start the process of disconnecting the tunnel.
//
// # Communicating with the tunnel provider
//
//   - [INETunnelProviderSession.SendProviderMessageReturnErrorResponseHandler]: Send a message to the Tunnel Provider extension. If the extension is not running, it should be launched to handle the message. If this method can’t start sending the message it reports an error in the `returnError` parameter. If an error occurs while sending the message or returning the result, `nil` should be sent to the response handler as notification.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderSession
type INETunnelProviderSession interface {
	INEVPNConnection

	// Topic: Controlling the tunnel connection

	// Start the process of connecting the tunnel.
	StartTunnelWithOptionsAndReturnError(options foundation.INSDictionary) (bool, error)
	// Start the process of disconnecting the tunnel.
	StopTunnel()

	// Topic: Communicating with the tunnel provider

	// Send a message to the Tunnel Provider extension. If the extension is not running, it should be launched to handle the message. If this method can’t start sending the message it reports an error in the `returnError` parameter. If an error occurs while sending the message or returning the result, `nil` should be sent to the response handler as notification.
	SendProviderMessageReturnErrorResponseHandler(messageData foundation.INSData, error_ foundation.INSError, responseHandler DataHandler) bool
}

// Init initializes the instance.
func (t NETunnelProviderSession) Init() NETunnelProviderSession {
	rv := objc.Send[NETunnelProviderSession](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NETunnelProviderSession) Autorelease() NETunnelProviderSession {
	rv := objc.Send[NETunnelProviderSession](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNETunnelProviderSession creates a new NETunnelProviderSession instance.
func NewNETunnelProviderSession() NETunnelProviderSession {
	class := getNETunnelProviderSessionClass()
	rv := objc.Send[NETunnelProviderSession](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Start the process of connecting the tunnel.
//
// options: A dictionary containing options to be passed to the Tunnel Provider
// extension.
//
// # Discussion
//
// This method returns immediately after starting the process of connecting
// the tunnel. In order to be notified when the tunnel is fully connected,
// register to observe the [NEVPNStatusDidChangeNotification] notification on
// the [NETunnelProviderSession] object and examine its status property when
// the notification is received.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderSession/startTunnel(options:)
//
// [NEVPNStatusDidChangeNotification]: https://developer.apple.com/documentation/NetworkExtension/NEVPNStatusDidChangeNotification
func (t NETunnelProviderSession) StartTunnelWithOptionsAndReturnError(options foundation.INSDictionary) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](t.ID, objc.Sel("startTunnelWithOptions:andReturnError:"), options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("startTunnelWithOptions:andReturnError: returned NO with nil NSError")
	}
	return rv, nil

}

// Start the process of disconnecting the tunnel.
//
// # Discussion
//
// This method returns immediately after starting the process of disconnecting
// the tunnel. In order to be notified when the tunnel is fully disconnected,
// register to observe the [NEVPNStatusDidChangeNotification] notification on
// the [NETunnelProviderSession] object and examine its status property when
// the notification is received.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderSession/stopTunnel()
//
// [NEVPNStatusDidChangeNotification]: https://developer.apple.com/documentation/NetworkExtension/NEVPNStatusDidChangeNotification
func (t NETunnelProviderSession) StopTunnel() {
	objc.Send[objc.ID](t.ID, objc.Sel("stopTunnel"))
}

// Send a message to the Tunnel Provider extension. If the extension is not
// running, it should be launched to handle the message. If this method
// can’t start sending the message it reports an error in the `returnError`
// parameter. If an error occurs while sending the message or returning the
// result, `nil` should be sent to the response handler as notification.
//
// messageData: An [NSData] object containing the message to be sent.
//
// responseHandler: An optional block that handles the response from the Tunnel Provider
// extension. Pass nil if no response is expected.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/NetworkExtension/NETunnelProviderSession/sendProviderMessage(_:responseHandler:)
//
// [NSData]: https://developer.apple.com/documentation/Foundation/NSData
func (t NETunnelProviderSession) SendProviderMessageReturnErrorResponseHandler(messageData foundation.INSData, error_ foundation.INSError, responseHandler DataHandler) bool {
	_block2, _ := NewDataBlock(responseHandler)
	rv := objc.Send[bool](t.ID, objc.Sel("sendProviderMessage:returnError:responseHandler:"), messageData, error_, _block2)
	return rv
}

// SendProviderMessageReturnErrorResponseHandlerSync is a synchronous wrapper around [NETunnelProviderSession.SendProviderMessageReturnErrorResponseHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (t NETunnelProviderSession) SendProviderMessageReturnErrorResponseHandlerSync(ctx context.Context, messageData foundation.INSData, error_ foundation.INSError) (*foundation.NSData, error) {
	done := make(chan *foundation.NSData, 1)
	t.SendProviderMessageReturnErrorResponseHandler(messageData, error_, func(val *foundation.NSData) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
