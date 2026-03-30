// Code generated from Apple documentation for NetworkExtension. DO NOT EDIT.

package networkextension

import (
	"context"
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NEVPNConnection] class.
var (
	_NEVPNConnectionClass     NEVPNConnectionClass
	_NEVPNConnectionClassOnce sync.Once
)

func getNEVPNConnectionClass() NEVPNConnectionClass {
	_NEVPNConnectionClassOnce.Do(func() {
		_NEVPNConnectionClass = NEVPNConnectionClass{class: objc.GetClass("NEVPNConnection")}
	})
	return _NEVPNConnectionClass
}

// GetNEVPNConnectionClass returns the class object for NEVPNConnection.
func GetNEVPNConnectionClass() NEVPNConnectionClass {
	return getNEVPNConnectionClass()
}

type NEVPNConnectionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NEVPNConnectionClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NEVPNConnectionClass) Alloc() NEVPNConnection {
	rv := objc.Send[NEVPNConnection](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object to start and stop a Personal VPN connection and get its status.
//
// # Overview
//
// [NEVPNConnection] objects are not instantiated directly. Instead, each
// [NEVPNManager] object has an associated [NEVPNConnection] object as a
// read-only property.
//
// The [NEVPNConnection] class provides methods for starting and stopping the
// VPN programmatically. The other way that the VPN can be started and stopped
// is through VPN On Demand. See the `onDemandRules` property in
// [NEVPNManager] and [NEOnDemandRule].
//
// Instances of this class are thread safe.
//
// # Controlling the VPN connection
//
//   - [NEVPNConnection.StartVPNTunnelAndReturnError]: Start the process of connecting the VPN.
//   - [NEVPNConnection.StartVPNTunnelWithOptionsAndReturnError]: Start the process of connecting the VPN.
//   - [NEVPNConnection.NEVPNConnectionStartOptionUsername]
//   - [NEVPNConnection.NEVPNConnectionStartOptionPassword]
//   - [NEVPNConnection.StopVPNTunnel]: Start the process of disconnecting the VPN.
//
// # Getting VPN connection status
//
//   - [NEVPNConnection.Manager]
//   - [NEVPNConnection.Status]: The current status of the VPN connection.
//   - [NEVPNConnection.ConnectedDate]: The date and time when the connection status changed to [NEVPNStatusConnected].
//
// # Notifications
//
//   - [NEVPNConnection.NEVPNStatusDidChange]: Posted when the status of the VPN connection changes.
//
// # Handling errors
//
//   - [NEVPNConnection.FetchLastDisconnectErrorWithCompletionHandler]: Retrives the most recent error that caused the VPN to disconnect.
//   - [NEVPNConnection.NEVPNConnectionErrorDomain]: The domain for errors resulting from VPN connection calls.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnection
type NEVPNConnection struct {
	objectivec.Object
}

// NEVPNConnectionFromID constructs a [NEVPNConnection] from an objc.ID.
//
// An object to start and stop a Personal VPN connection and get its status.
func NEVPNConnectionFromID(id objc.ID) NEVPNConnection {
	return NEVPNConnection{objectivec.Object{ID: id}}
}

// NOTE: NEVPNConnection adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NEVPNConnection] class.
//
// # Controlling the VPN connection
//
//   - [INEVPNConnection.StartVPNTunnelAndReturnError]: Start the process of connecting the VPN.
//   - [INEVPNConnection.StartVPNTunnelWithOptionsAndReturnError]: Start the process of connecting the VPN.
//   - [INEVPNConnection.NEVPNConnectionStartOptionUsername]
//   - [INEVPNConnection.NEVPNConnectionStartOptionPassword]
//   - [INEVPNConnection.StopVPNTunnel]: Start the process of disconnecting the VPN.
//
// # Getting VPN connection status
//
//   - [INEVPNConnection.Manager]
//   - [INEVPNConnection.Status]: The current status of the VPN connection.
//   - [INEVPNConnection.ConnectedDate]: The date and time when the connection status changed to [NEVPNStatusConnected].
//
// # Notifications
//
//   - [INEVPNConnection.NEVPNStatusDidChange]: Posted when the status of the VPN connection changes.
//
// # Handling errors
//
//   - [INEVPNConnection.FetchLastDisconnectErrorWithCompletionHandler]: Retrives the most recent error that caused the VPN to disconnect.
//   - [INEVPNConnection.NEVPNConnectionErrorDomain]: The domain for errors resulting from VPN connection calls.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnection
type INEVPNConnection interface {
	objectivec.IObject

	// Topic: Controlling the VPN connection

	// Start the process of connecting the VPN.
	StartVPNTunnelAndReturnError() (bool, error)
	// Start the process of connecting the VPN.
	StartVPNTunnelWithOptionsAndReturnError(options foundation.INSDictionary) (bool, error)
	NEVPNConnectionStartOptionUsername() string
	NEVPNConnectionStartOptionPassword() string
	// Start the process of disconnecting the VPN.
	StopVPNTunnel()

	// Topic: Getting VPN connection status

	Manager() INEVPNManager
	// The current status of the VPN connection.
	Status() NEVPNStatus
	// The date and time when the connection status changed to [NEVPNStatusConnected].
	ConnectedDate() foundation.INSDate

	// Topic: Notifications

	// Posted when the status of the VPN connection changes.
	NEVPNStatusDidChange() foundation.NSString

	// Topic: Handling errors

	// Retrives the most recent error that caused the VPN to disconnect.
	FetchLastDisconnectErrorWithCompletionHandler(handler ErrorHandler)
	// The domain for errors resulting from VPN connection calls.
	NEVPNConnectionErrorDomain() string
}

// Init initializes the instance.
func (v NEVPNConnection) Init() NEVPNConnection {
	rv := objc.Send[NEVPNConnection](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v NEVPNConnection) Autorelease() NEVPNConnection {
	rv := objc.Send[NEVPNConnection](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewNEVPNConnection creates a new NEVPNConnection instance.
func NewNEVPNConnection() NEVPNConnection {
	class := getNEVPNConnectionClass()
	rv := objc.Send[NEVPNConnection](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Start the process of connecting the VPN.
//
// # Discussion
//
// This method returns immediately after starting the process of connecting
// the VPN. In order to be notified when the VPN is fully connected, register
// to observe the [NEVPNStatusDidChangeNotification] notification on the
// [NEVPNConnection] object, and examine the status property when the
// notification is received.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnection/startVPNTunnel()
//
// [NEVPNStatusDidChangeNotification]: https://developer.apple.com/documentation/NetworkExtension/NEVPNStatusDidChangeNotification
func (v NEVPNConnection) StartVPNTunnelAndReturnError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("startVPNTunnelAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("startVPNTunnelAndReturnError: returned NO with nil NSError")
	}
	return rv, nil

}

// Start the process of connecting the VPN.
//
// options: An [NSDictionary] that will be passed to the tunnel provider during the
// process of starting the tunnel. See Constants, below.
//
// # Discussion
//
// This method returns immediately after starting the process of connecting
// the VPN. In order to be notified when the VPN is fully connected, register
// to observe the [NEVPNStatusDidChangeNotification] notification on the
// [NEVPNConnection] object, and examine the status property when the
// notification is received.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnection/startVPNTunnel(options:)
//
// [NEVPNStatusDidChangeNotification]: https://developer.apple.com/documentation/NetworkExtension/NEVPNStatusDidChangeNotification
func (v NEVPNConnection) StartVPNTunnelWithOptionsAndReturnError(options foundation.INSDictionary) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](v.ID, objc.Sel("startVPNTunnelWithOptions:andReturnError:"), options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("startVPNTunnelWithOptions:andReturnError: returned NO with nil NSError")
	}
	return rv, nil

}

// Start the process of disconnecting the VPN.
//
// # Discussion
//
// This method returns immediately after starting the process of disconnecting
// the VPN. In order to be notified when the VPN is fully disconnected,
// register to observe the [NEVPNStatusDidChangeNotification] notification on
// the [NEVPNConnection] object and examine the status property when the
// notification is received.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnection/stopVPNTunnel()
//
// [NEVPNStatusDidChangeNotification]: https://developer.apple.com/documentation/NetworkExtension/NEVPNStatusDidChangeNotification
func (v NEVPNConnection) StopVPNTunnel() {
	objc.Send[objc.ID](v.ID, objc.Sel("stopVPNTunnel"))
}

// Retrives the most recent error that caused the VPN to disconnect.
//
// handler: An error handler that receives the last disconnect error as a parameter.
//
// # Discussion
//
// If VPN system (including the IPsec client) generated the error, then the
// error uses the [NEVPNConnectionErrorDomain] error domain. If the error came
// from a tunnel provider app extension instead, then the error is the
// [NSError] that the provider passed when disconnecting the tunnel.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnection/fetchLastDisconnectError(completionHandler:)
//
// [NSError]: https://developer.apple.com/documentation/Foundation/NSError
func (v NEVPNConnection) FetchLastDisconnectErrorWithCompletionHandler(handler ErrorHandler) {
	_block0, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](v.ID, objc.Sel("fetchLastDisconnectErrorWithCompletionHandler:"), _block0)
}

// See: https://developer.apple.com/documentation/networkextension/nevpnconnectionstartoptionusername
func (v NEVPNConnection) NEVPNConnectionStartOptionUsername() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("NEVPNConnectionStartOptionUsername"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/networkextension/nevpnconnectionstartoptionpassword
func (v NEVPNConnection) NEVPNConnectionStartOptionPassword() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("NEVPNConnectionStartOptionPassword"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnection/manager
func (v NEVPNConnection) Manager() INEVPNManager {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("manager"))
	return NEVPNManagerFromID(objc.ID(rv))
}

// The current status of the VPN connection.
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnection/status
func (v NEVPNConnection) Status() NEVPNStatus {
	rv := objc.Send[NEVPNStatus](v.ID, objc.Sel("status"))
	return NEVPNStatus(rv)
}

// The date and time when the connection status changed to
// [NEVPNStatusConnected].
//
// # Discussion
//
// This property contains the date and time when the connection status changed
// to [NEVPNStatusConnected] after previously being set to
// [NEVPNStatusDisconnected]. This property is set to nil whenever the status
// changes to [NEVPNStatusDisconnected].
//
// See: https://developer.apple.com/documentation/NetworkExtension/NEVPNConnection/connectedDate
func (v NEVPNConnection) ConnectedDate() foundation.INSDate {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("connectedDate"))
	return foundation.NSDateFromID(objc.ID(rv))
}

// Posted when the status of the VPN connection changes.
//
// See: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NEVPNStatusDidChange
func (v NEVPNConnection) NEVPNStatusDidChange() foundation.NSString {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("NEVPNStatusDidChange"))
	return foundation.NSStringFromID(objc.ID(rv))
}

// The domain for errors resulting from VPN connection calls.
//
// See: https://developer.apple.com/documentation/networkextension/nevpnconnectionerrordomain
func (v NEVPNConnection) NEVPNConnectionErrorDomain() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("NEVPNConnectionErrorDomain"))
	return foundation.NSStringFromID(rv).String()
}

// FetchLastDisconnectError is a synchronous wrapper around [NEVPNConnection.FetchLastDisconnectErrorWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (v NEVPNConnection) FetchLastDisconnectError(ctx context.Context) error {
	done := make(chan error, 1)
	v.FetchLastDisconnectErrorWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
