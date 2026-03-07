// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSXPCConnection] class.
var (
	_NSXPCConnectionClass     NSXPCConnectionClass
	_NSXPCConnectionClassOnce sync.Once
)

func getNSXPCConnectionClass() NSXPCConnectionClass {
	_NSXPCConnectionClassOnce.Do(func() {
		_NSXPCConnectionClass = NSXPCConnectionClass{class: objc.GetClass("NSXPCConnection")}
	})
	return _NSXPCConnectionClass
}

// GetNSXPCConnectionClass returns the class object for NSXPCConnection.
func GetNSXPCConnectionClass() NSXPCConnectionClass {
	return getNSXPCConnectionClass()
}

type NSXPCConnectionClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSXPCConnectionClass) Alloc() NSXPCConnection {
	rv := objc.Send[NSXPCConnection](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// A bidirectional communication channel between two processes.
//
// # Overview
// 
// This class is the primary means of creating and configuring the
// communication mechanism between two processes. Each process has one
// instance of this class to represent the endpoint in the communication
// channel.
//
// # Creating a connection
//
//   - [NSXPCConnection.InitWithListenerEndpoint]: Initializes an [NSXPCConnection](<doc://com.apple.foundation/documentation/Foundation/NSXPCConnection>) object to connect to an [NSXPCListener](<doc://com.apple.foundation/documentation/Foundation/NSXPCListener>) object in another process, identified by an [NSXPCListenerEndpoint](<doc://com.apple.foundation/documentation/Foundation/NSXPCListenerEndpoint>) object.
//   - [NSXPCConnection.InitWithMachServiceNameOptions]: Initializes an [NSXPCConnection](<doc://com.apple.foundation/documentation/Foundation/NSXPCConnection>) object to connect to a LaunchAgent or LaunchDaemon with a name advertised in a `launchd.Plist()`.
//   - [NSXPCConnection.InitWithServiceName]: Initializes an [NSXPCConnection](<doc://com.apple.foundation/documentation/Foundation/NSXPCConnection>) object to connect to an [NSXPCListener](<doc://com.apple.foundation/documentation/Foundation/NSXPCListener>) object in an XPC service, identified by a service name.
//
// # Managing connection state
//
//   - [NSXPCConnection.Activate]: Activates the connection.
//   - [NSXPCConnection.Resume]: Starts or resumes handling of messages on a connection.
//   - [NSXPCConnection.Invalidate]: Invalidates the connection.
//   - [NSXPCConnection.Suspend]: Suspends the connection.
//   - [NSXPCConnection.InterruptionHandler]: An interruption handler that is called if the remote process exits or crashes.
//   - [NSXPCConnection.SetInterruptionHandler]
//   - [NSXPCConnection.InvalidationHandler]: An invalidation handler that is called if the connection can not be formed or the connection has terminated and may not be re-established.
//   - [NSXPCConnection.SetInvalidationHandler]
//   - [NSXPCConnection.ScheduleSendBarrierBlock]: Add a barrier block to execute on the connection.
//
// # Managing the connection interface
//
//   - [NSXPCConnection.ServiceName]: The name of the XPC service that this connection was configured to connect to.
//   - [NSXPCConnection.Endpoint]: If the connection was created with an [NSXPCListenerEndpoint](<doc://com.apple.foundation/documentation/Foundation/NSXPCListenerEndpoint>) object, returns the endpoint object used.
//   - [NSXPCConnection.ExportedInterface]: The [NSXPCInterface](<doc://com.apple.foundation/documentation/Foundation/NSXPCInterface>) object that describes the protocol for the exported object on this connection.
//   - [NSXPCConnection.SetExportedInterface]
//   - [NSXPCConnection.ExportedObject]: An exported object for the connection.
//   - [NSXPCConnection.SetExportedObject]
//   - [NSXPCConnection.RemoteObjectInterface]: Defines the [NSXPCInterface](<doc://com.apple.foundation/documentation/Foundation/NSXPCInterface>) object that describes the protocol for the object represented by the `remoteObjectProxy`.
//   - [NSXPCConnection.SetRemoteObjectInterface]
//
// # Working with security attributes
//
//   - [NSXPCConnection.AuditSessionIdentifier]: The BSM audit session identifier for the connecting process.
//   - [NSXPCConnection.ProcessIdentifier]: The process ID (PID) of the connecting process.
//   - [NSXPCConnection.EffectiveGroupIdentifier]: The effective group ID (EGID) of the connecting process.
//   - [NSXPCConnection.EffectiveUserIdentifier]: The effective user ID (EUID) of the connecting process.
//
// # Working with code signing
//
//   - [NSXPCConnection.SetCodeSigningRequirement]: Sets the code signing requirement for this connection.
//
// # Error codes
//
//   - [NSXPCConnection.NSXPCConnectionInterrupted]: The XPC connection was interrupted.
//   - [NSXPCConnection.SetNSXPCConnectionInterrupted]
//   - [NSXPCConnection.NSXPCConnectionInvalid]: The XPC connection was invalid.
//   - [NSXPCConnection.SetNSXPCConnectionInvalid]
//   - [NSXPCConnection.NSXPCConnectionReplyInvalid]: The XPC connection reply was invalid.
//   - [NSXPCConnection.SetNSXPCConnectionReplyInvalid]
//   - [NSXPCConnection.NSXPCConnectionErrorMinimum]: The lower bounds of XPC connection error code values.
//   - [NSXPCConnection.SetNSXPCConnectionErrorMinimum]
//   - [NSXPCConnection.NSXPCConnectionErrorMaximum]: The upper bounds of XPC connection error code values.
//   - [NSXPCConnection.SetNSXPCConnectionErrorMaximum]
//   - [NSXPCConnection.NSXPCConnectionCodeSigningRequirementFailure]: A code-signing requirement check failed.
//   - [NSXPCConnection.SetNSXPCConnectionCodeSigningRequirementFailure]
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCConnection
type NSXPCConnection struct {
	objectivec.Object
}

// NSXPCConnectionFromID constructs a [NSXPCConnection] from an objc.ID.
//
// A bidirectional communication channel between two processes.
func NSXPCConnectionFromID(id objc.ID) NSXPCConnection {
	return NSXPCConnection{objectivec.Object{id}}
}
// NOTE: NSXPCConnection adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSXPCConnection] class.
//
// # Creating a connection
//
//   - [INSXPCConnection.InitWithListenerEndpoint]: Initializes an [NSXPCConnection](<doc://com.apple.foundation/documentation/Foundation/NSXPCConnection>) object to connect to an [NSXPCListener](<doc://com.apple.foundation/documentation/Foundation/NSXPCListener>) object in another process, identified by an [NSXPCListenerEndpoint](<doc://com.apple.foundation/documentation/Foundation/NSXPCListenerEndpoint>) object.
//   - [INSXPCConnection.InitWithMachServiceNameOptions]: Initializes an [NSXPCConnection](<doc://com.apple.foundation/documentation/Foundation/NSXPCConnection>) object to connect to a LaunchAgent or LaunchDaemon with a name advertised in a `launchd.Plist()`.
//   - [INSXPCConnection.InitWithServiceName]: Initializes an [NSXPCConnection](<doc://com.apple.foundation/documentation/Foundation/NSXPCConnection>) object to connect to an [NSXPCListener](<doc://com.apple.foundation/documentation/Foundation/NSXPCListener>) object in an XPC service, identified by a service name.
//
// # Managing connection state
//
//   - [INSXPCConnection.Activate]: Activates the connection.
//   - [INSXPCConnection.Resume]: Starts or resumes handling of messages on a connection.
//   - [INSXPCConnection.Invalidate]: Invalidates the connection.
//   - [INSXPCConnection.Suspend]: Suspends the connection.
//   - [INSXPCConnection.InterruptionHandler]: An interruption handler that is called if the remote process exits or crashes.
//   - [INSXPCConnection.SetInterruptionHandler]
//   - [INSXPCConnection.InvalidationHandler]: An invalidation handler that is called if the connection can not be formed or the connection has terminated and may not be re-established.
//   - [INSXPCConnection.SetInvalidationHandler]
//   - [INSXPCConnection.ScheduleSendBarrierBlock]: Add a barrier block to execute on the connection.
//
// # Managing the connection interface
//
//   - [INSXPCConnection.ServiceName]: The name of the XPC service that this connection was configured to connect to.
//   - [INSXPCConnection.Endpoint]: If the connection was created with an [NSXPCListenerEndpoint](<doc://com.apple.foundation/documentation/Foundation/NSXPCListenerEndpoint>) object, returns the endpoint object used.
//   - [INSXPCConnection.ExportedInterface]: The [NSXPCInterface](<doc://com.apple.foundation/documentation/Foundation/NSXPCInterface>) object that describes the protocol for the exported object on this connection.
//   - [INSXPCConnection.SetExportedInterface]
//   - [INSXPCConnection.ExportedObject]: An exported object for the connection.
//   - [INSXPCConnection.SetExportedObject]
//   - [INSXPCConnection.RemoteObjectInterface]: Defines the [NSXPCInterface](<doc://com.apple.foundation/documentation/Foundation/NSXPCInterface>) object that describes the protocol for the object represented by the `remoteObjectProxy`.
//   - [INSXPCConnection.SetRemoteObjectInterface]
//
// # Working with security attributes
//
//   - [INSXPCConnection.AuditSessionIdentifier]: The BSM audit session identifier for the connecting process.
//   - [INSXPCConnection.ProcessIdentifier]: The process ID (PID) of the connecting process.
//   - [INSXPCConnection.EffectiveGroupIdentifier]: The effective group ID (EGID) of the connecting process.
//   - [INSXPCConnection.EffectiveUserIdentifier]: The effective user ID (EUID) of the connecting process.
//
// # Working with code signing
//
//   - [INSXPCConnection.SetCodeSigningRequirement]: Sets the code signing requirement for this connection.
//
// # Error codes
//
//   - [INSXPCConnection.NSXPCConnectionInterrupted]: The XPC connection was interrupted.
//   - [INSXPCConnection.SetNSXPCConnectionInterrupted]
//   - [INSXPCConnection.NSXPCConnectionInvalid]: The XPC connection was invalid.
//   - [INSXPCConnection.SetNSXPCConnectionInvalid]
//   - [INSXPCConnection.NSXPCConnectionReplyInvalid]: The XPC connection reply was invalid.
//   - [INSXPCConnection.SetNSXPCConnectionReplyInvalid]
//   - [INSXPCConnection.NSXPCConnectionErrorMinimum]: The lower bounds of XPC connection error code values.
//   - [INSXPCConnection.SetNSXPCConnectionErrorMinimum]
//   - [INSXPCConnection.NSXPCConnectionErrorMaximum]: The upper bounds of XPC connection error code values.
//   - [INSXPCConnection.SetNSXPCConnectionErrorMaximum]
//   - [INSXPCConnection.NSXPCConnectionCodeSigningRequirementFailure]: A code-signing requirement check failed.
//   - [INSXPCConnection.SetNSXPCConnectionCodeSigningRequirementFailure]
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCConnection
type INSXPCConnection interface {
	objectivec.IObject
	NSXPCProxyCreating

	// Topic: Creating a connection

	// Initializes an [NSXPCConnection](<doc://com.apple.foundation/documentation/Foundation/NSXPCConnection>) object to connect to an [NSXPCListener](<doc://com.apple.foundation/documentation/Foundation/NSXPCListener>) object in another process, identified by an [NSXPCListenerEndpoint](<doc://com.apple.foundation/documentation/Foundation/NSXPCListenerEndpoint>) object.
	InitWithListenerEndpoint(endpoint INSXPCListenerEndpoint) NSXPCConnection
	// Initializes an [NSXPCConnection](<doc://com.apple.foundation/documentation/Foundation/NSXPCConnection>) object to connect to a LaunchAgent or LaunchDaemon with a name advertised in a `launchd.Plist()`.
	InitWithMachServiceNameOptions(name string, options NSXPCConnectionOptions) NSXPCConnection
	// Initializes an [NSXPCConnection](<doc://com.apple.foundation/documentation/Foundation/NSXPCConnection>) object to connect to an [NSXPCListener](<doc://com.apple.foundation/documentation/Foundation/NSXPCListener>) object in an XPC service, identified by a service name.
	InitWithServiceName(serviceName string) NSXPCConnection

	// Topic: Managing connection state

	// Activates the connection.
	Activate()
	// Starts or resumes handling of messages on a connection.
	Resume()
	// Invalidates the connection.
	Invalidate()
	// Suspends the connection.
	Suspend()
	// An interruption handler that is called if the remote process exits or crashes.
	InterruptionHandler() VoidHandler
	SetInterruptionHandler(value VoidHandler)
	// An invalidation handler that is called if the connection can not be formed or the connection has terminated and may not be re-established.
	InvalidationHandler() VoidHandler
	SetInvalidationHandler(value VoidHandler)
	// Add a barrier block to execute on the connection.
	ScheduleSendBarrierBlock(block VoidHandler)

	// Topic: Managing the connection interface

	// The name of the XPC service that this connection was configured to connect to.
	ServiceName() string
	// If the connection was created with an [NSXPCListenerEndpoint](<doc://com.apple.foundation/documentation/Foundation/NSXPCListenerEndpoint>) object, returns the endpoint object used.
	Endpoint() INSXPCListenerEndpoint
	// The [NSXPCInterface](<doc://com.apple.foundation/documentation/Foundation/NSXPCInterface>) object that describes the protocol for the exported object on this connection.
	ExportedInterface() INSXPCInterface
	SetExportedInterface(value INSXPCInterface)
	// An exported object for the connection.
	ExportedObject() objectivec.IObject
	SetExportedObject(value objectivec.IObject)
	// Defines the [NSXPCInterface](<doc://com.apple.foundation/documentation/Foundation/NSXPCInterface>) object that describes the protocol for the object represented by the `remoteObjectProxy`.
	RemoteObjectInterface() INSXPCInterface
	SetRemoteObjectInterface(value INSXPCInterface)

	// Topic: Working with security attributes

	// The BSM audit session identifier for the connecting process.
	AuditSessionIdentifier() int32
	// The process ID (PID) of the connecting process.
	ProcessIdentifier() int32
	// The effective group ID (EGID) of the connecting process.
	EffectiveGroupIdentifier() uint32
	// The effective user ID (EUID) of the connecting process.
	EffectiveUserIdentifier() uint32

	// Topic: Working with code signing

	// Sets the code signing requirement for this connection.
	SetCodeSigningRequirement(requirement string)

	// Topic: Error codes

	// The XPC connection was interrupted.
	NSXPCConnectionInterrupted() int
	SetNSXPCConnectionInterrupted(value int)
	// The XPC connection was invalid.
	NSXPCConnectionInvalid() int
	SetNSXPCConnectionInvalid(value int)
	// The XPC connection reply was invalid.
	NSXPCConnectionReplyInvalid() int
	SetNSXPCConnectionReplyInvalid(value int)
	// The lower bounds of XPC connection error code values.
	NSXPCConnectionErrorMinimum() int
	SetNSXPCConnectionErrorMinimum(value int)
	// The upper bounds of XPC connection error code values.
	NSXPCConnectionErrorMaximum() int
	SetNSXPCConnectionErrorMaximum(value int)
	// A code-signing requirement check failed.
	NSXPCConnectionCodeSigningRequirementFailure() int
	SetNSXPCConnectionCodeSigningRequirementFailure(value int)
}





// Init initializes the instance.
func (x NSXPCConnection) Init() NSXPCConnection {
	rv := objc.Send[NSXPCConnection](x.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (x NSXPCConnection) Autorelease() NSXPCConnection {
	rv := objc.Send[NSXPCConnection](x.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSXPCConnection creates a new NSXPCConnection instance.
func NewNSXPCConnection() NSXPCConnection {
	class := getNSXPCConnectionClass()
	rv := objc.Send[NSXPCConnection](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Initializes an [NSXPCConnection] object to connect to an [NSXPCListener]
// object in another process, identified by an [NSXPCListenerEndpoint] object.
//
// endpoint: The desired listener endpoint for the service.
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCConnection/init(listenerEndpoint:)
func NewXPCConnectionWithListenerEndpoint(endpoint INSXPCListenerEndpoint) NSXPCConnection {
	instance := getNSXPCConnectionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithListenerEndpoint:"), endpoint)
	return NSXPCConnectionFromID(rv)
}


// Initializes an [NSXPCConnection] object to connect to a LaunchAgent or
// LaunchDaemon with a name advertised in a `launchd.Plist()`.
//
// # Discussion
// 
// For example, if an agent is managed with `launchd` and has a
// `launchd.Plist()` in `~/Library/LaunchAgents`, this method would create a
// connection to that agent. The agent should use [NSXPCListener] to wait for
// new connections.
// 
// If the connection is being made to a process that is running in a
// privileged Mach bootstrap context (for example, a daemon started by a
// `launchd` property list in `/Library/LaunchDaemons`), then pass the
// [NSXPCConnection] option.
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCConnection/init(machServiceName:options:)
func NewXPCConnectionWithMachServiceNameOptions(name string, options NSXPCConnectionOptions) NSXPCConnection {
	instance := getNSXPCConnectionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMachServiceName:options:"), objc.String(name), options)
	return NSXPCConnectionFromID(rv)
}


// Initializes an [NSXPCConnection] object to connect to an [NSXPCListener]
// object in an XPC service, identified by a service name.
//
// # Discussion
// 
// XPC services are helper processes that are usually part of your application
// bundle. The service should use [NSXPCListener] to wait for new connections.
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCConnection/init(serviceName:)
func NewXPCConnectionWithServiceName(serviceName string) NSXPCConnection {
	instance := getNSXPCConnectionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithServiceName:"), objc.String(serviceName))
	return NSXPCConnectionFromID(rv)
}







// Initializes an [NSXPCConnection] object to connect to an [NSXPCListener]
// object in another process, identified by an [NSXPCListenerEndpoint] object.
//
// endpoint: The desired listener endpoint for the service.
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCConnection/init(listenerEndpoint:)
func (x NSXPCConnection) InitWithListenerEndpoint(endpoint INSXPCListenerEndpoint) NSXPCConnection {
	rv := objc.Send[NSXPCConnection](x.ID, objc.Sel("initWithListenerEndpoint:"), endpoint)
	return rv
}

// Initializes an [NSXPCConnection] object to connect to a LaunchAgent or
// LaunchDaemon with a name advertised in a `launchd.Plist()`.
//
// # Discussion
// 
// For example, if an agent is managed with `launchd` and has a
// `launchd.Plist()` in `~/Library/LaunchAgents`, this method would create a
// connection to that agent. The agent should use [NSXPCListener] to wait for
// new connections.
// 
// If the connection is being made to a process that is running in a
// privileged Mach bootstrap context (for example, a daemon started by a
// `launchd` property list in `/Library/LaunchDaemons`), then pass the
// [NSXPCConnection] option.
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCConnection/init(machServiceName:options:)
func (x NSXPCConnection) InitWithMachServiceNameOptions(name string, options NSXPCConnectionOptions) NSXPCConnection {
	rv := objc.Send[NSXPCConnection](x.ID, objc.Sel("initWithMachServiceName:options:"), objc.String(name), options)
	return rv
}

// Initializes an [NSXPCConnection] object to connect to an [NSXPCListener]
// object in an XPC service, identified by a service name.
//
// # Discussion
// 
// XPC services are helper processes that are usually part of your application
// bundle. The service should use [NSXPCListener] to wait for new connections.
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCConnection/init(serviceName:)
func (x NSXPCConnection) InitWithServiceName(serviceName string) NSXPCConnection {
	rv := objc.Send[NSXPCConnection](x.ID, objc.Sel("initWithServiceName:"), objc.String(serviceName))
	return rv
}

// Activates the connection.
//
// # Discussion
// 
// Connections start in an inactive state. You must call [Activate] on a
// connection before it can send or receive any messages.
// 
// Calling [Activate] on an active connection has no effect.
// 
// For backward compatibility reasons, calling [Resume] on an inactive and
// otherwise not suspended [NSXPCConnection] has the same effect as calling
// [Activate]. For new code, prefer [Activate].
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCConnection/activate()
func (x NSXPCConnection) Activate() {
	objc.Send[objc.ID](x.ID, objc.Sel("activate"))
}

// Starts or resumes handling of messages on a connection.
//
// # Discussion
// 
// All connections start suspended. You must resume them before they start
// processing received messages or sending messages through the
// [RemoteObjectProxy] object.
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCConnection/resume()
func (x NSXPCConnection) Resume() {
	objc.Send[objc.ID](x.ID, objc.Sel("resume"))
}

// Invalidates the connection.
//
// # Discussion
// 
// When you call this method, all outstanding reply blocks, error handling
// blocks, and invalidation blocks are called on the message handling queue.
// The connection must be invalidated before it is deallocated. After a
// connection is invalidated, no more messages may be sent or received.
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCConnection/invalidate()
func (x NSXPCConnection) Invalidate() {
	objc.Send[objc.ID](x.ID, objc.Sel("invalidate"))
}

// Suspends the connection.
//
// # Discussion
// 
// As you cannot invalidate a suspended connection, every call to [Suspend]
// that you make must be balanced by a call to [Resume].
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCConnection/suspend()
func (x NSXPCConnection) Suspend() {
	objc.Send[objc.ID](x.ID, objc.Sel("suspend"))
}

// Add a barrier block to execute on the connection.
//
// block: A block or closure to execute. This block takes no parameters and returns
// no value.
//
// # Discussion
// 
// This barrier block runs after any outstanding send commands complete.
// However, the remote process isn’t guaranteed to receive the sent messages
// by the time the block executes. If you need to ensure the remote process
// received a message, wait for a reply from the process.
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCConnection/scheduleSendBarrierBlock(_:)
func (x NSXPCConnection) ScheduleSendBarrierBlock(block VoidHandler) {
		_block0, _cleanup0 := NewVoidBlock(block)
	defer _cleanup0()
		objc.Send[objc.ID](x.ID, objc.Sel("scheduleSendBarrierBlock:"), _block0)
}

// Returns a proxy for the remote object (that is, the object exported from
// the other side of this connection) with the specified error handler.
//
// # Discussion
// 
// See descriptions in [NSXPCProxyCreating] for more details.
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCConnection/remoteObjectProxyWithErrorHandler(_:)
func (x NSXPCConnection) RemoteObjectProxyWithErrorHandler(handler ErrorHandler) objectivec.IObject {
		_block0, _cleanup0 := NewErrorBlock(handler)
	defer _cleanup0()
		rv := objc.Send[objc.ID](x.ID, objc.Sel("remoteObjectProxyWithErrorHandler:"), _block0)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/Foundation/NSXPCConnection/synchronousRemoteObjectProxyWithErrorHandler(_:)
func (x NSXPCConnection) SynchronousRemoteObjectProxyWithErrorHandler(handler ErrorHandler) objectivec.IObject {
		_block0, _cleanup0 := NewErrorBlock(handler)
	defer _cleanup0()
		rv := objc.Send[objc.ID](x.ID, objc.Sel("synchronousRemoteObjectProxyWithErrorHandler:"), _block0)
	return objectivec.Object{ID: rv}
}

// Sets the code signing requirement for this connection.
//
// requirement: A string that describes requirements expected of the connection peer. See
// [Code Signing Guide] for more information on the code signing format.
// //
// [Code Signing Guide]: https://developer.apple.com/library/archive/documentation/Security/Conceptual/CodeSigningGuide/
//
// # Discussion
// 
// Use this method to enforce a code-signing requirement on the peer process
// of an XPC connection.
// 
// The following example shows how a client can ensure that the XPC service on
// the other end of a connection has a specific entitlement.
// 
// Calling this method with a malformed `requirement` results in a fatal error
// in Swift, or throws an exception in Objective-C. If new messages don’t
// match the requirement, the connection becomes invalidated.
// 
// Call this method before calling [Resume], since it’s an XPC error to call
// this method more than once.
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCConnection/setCodeSigningRequirement(_:)
func (x NSXPCConnection) SetCodeSigningRequirement(requirement string) {
	objc.Send[objc.ID](x.ID, objc.Sel("setCodeSigningRequirement:"), objc.String(requirement))
}





// Returns the current connection, in the context of a call to a method on
// your exported object.
//
// # Return Value
// 
// An [NSXPCConnection] object, representing a connection to another process.
//
// # Discussion
// 
// Use this method to determine what process invoked the current call.
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCConnection/current()
func (_NSXPCConnectionClass NSXPCConnectionClass) CurrentConnection() NSXPCConnection {
	rv := objc.Send[objc.ID](objc.ID(_NSXPCConnectionClass.class), objc.Sel("currentConnection"))
	return NSXPCConnectionFromID(rv)
}








// An interruption handler that is called if the remote process exits or
// crashes.
//
// # Discussion
// 
// It may be possible to re-establish the connection by simply sending another
// message. The handler is invoked on the same queue as reply messages and
// other handlers, and it is always executed after any other messages or reply
// block handlers (except for the invalidation handler).
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCConnection/interruptionHandler
func (x NSXPCConnection) InterruptionHandler() VoidHandler {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("interruptionHandler"))
	_ = rv
	return nil
}
func (x NSXPCConnection) SetInterruptionHandler(value VoidHandler) {
	block, cleanup := NewVoidBlock(value)
	defer cleanup()
	objc.Send[struct{}](x.ID, objc.Sel("setInterruptionHandler:"), block)
}



// An invalidation handler that is called if the connection can not be formed
// or the connection has terminated and may not be re-established.
//
// # Discussion
// 
// This handler is invoked on the same queue as reply messages and other
// handlers, and is always executed last (after the interruption handler, if
// required). You may not send messages over the connection from within an
// invalidation handler block.
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCConnection/invalidationHandler
func (x NSXPCConnection) InvalidationHandler() VoidHandler {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("invalidationHandler"))
	_ = rv
	return nil
}
func (x NSXPCConnection) SetInvalidationHandler(value VoidHandler) {
	block, cleanup := NewVoidBlock(value)
	defer cleanup()
	objc.Send[struct{}](x.ID, objc.Sel("setInvalidationHandler:"), block)
}



// The name of the XPC service that this connection was configured to connect
// to.
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCConnection/serviceName
func (x NSXPCConnection) ServiceName() string {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("serviceName"))
	return NSStringFromID(rv).String()
}



// If the connection was created with an [NSXPCListenerEndpoint] object,
// returns the endpoint object used.
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCConnection/endpoint
func (x NSXPCConnection) Endpoint() INSXPCListenerEndpoint {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("endpoint"))
	return NSXPCListenerEndpointFromID(objc.ID(rv))
}



// The [NSXPCInterface] object that describes the protocol for the exported
// object on this connection.
//
// # Discussion
// 
// This value is required if a exported object is set.
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCConnection/exportedInterface
func (x NSXPCConnection) ExportedInterface() INSXPCInterface {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("exportedInterface"))
	return NSXPCInterfaceFromID(objc.ID(rv))
}
func (x NSXPCConnection) SetExportedInterface(value INSXPCInterface) {
	objc.Send[struct{}](x.ID, objc.Sel("setExportedInterface:"), value)
}



// An exported object for the connection.
//
// # Discussion
// 
// Messages sent to the [RemoteObjectProxy] object from the other side of the
// connection are dispatched to this object. Messages delivered to exported
// objects are serialized and sent on a non-main queue. The receiver is
// responsible for handling the messages on a different queue or thread if it
// is required.
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCConnection/exportedObject
func (x NSXPCConnection) ExportedObject() objectivec.IObject {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("exportedObject"))
	return objectivec.Object{ID: rv}
}
func (x NSXPCConnection) SetExportedObject(value objectivec.IObject) {
	objc.Send[struct{}](x.ID, objc.Sel("setExportedObject:"), value)
}



// Defines the [NSXPCInterface] object that describes the protocol for the
// object represented by the `remoteObjectProxy`.
//
// # Discussion
// 
// The proxy represents the `exportedObject` on the [NSXPCConnection] in the
// other process.
// 
// This value is required if messages are sent over this connection.
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCConnection/remoteObjectInterface
func (x NSXPCConnection) RemoteObjectInterface() INSXPCInterface {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("remoteObjectInterface"))
	return NSXPCInterfaceFromID(objc.ID(rv))
}
func (x NSXPCConnection) SetRemoteObjectInterface(value INSXPCInterface) {
	objc.Send[struct{}](x.ID, objc.Sel("setRemoteObjectInterface:"), value)
}



// Returns a proxy for the remote object (that is, the `exportedObject` from
// the other side of this connection).
//
// # Discussion
// 
// See descriptions in [NSXPCProxyCreating] for more details.
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCConnection/remoteObjectProxy
func (x NSXPCConnection) RemoteObjectProxy() objectivec.IObject {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("remoteObjectProxy"))
	return objectivec.Object{ID: rv}
}



// The BSM audit session identifier for the connecting process.
//
// # Discussion
// 
// This attribute may be used by the listener delegate to accept or reject
// connections.
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCConnection/auditSessionIdentifier
func (x NSXPCConnection) AuditSessionIdentifier() int32 {
	rv := objc.Send[int32](x.ID, objc.Sel("auditSessionIdentifier"))
	return rv
}



// The process ID (PID) of the connecting process.
//
// # Discussion
// 
// This attribute may be used by the listener delegate to accept or reject
// connections.
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCConnection/processIdentifier
func (x NSXPCConnection) ProcessIdentifier() int32 {
	rv := objc.Send[int32](x.ID, objc.Sel("processIdentifier"))
	return rv
}



// The effective group ID (EGID) of the connecting process.
//
// # Discussion
// 
// This attribute may be used by the listener delegate to accept or reject
// connections.
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCConnection/effectiveGroupIdentifier
func (x NSXPCConnection) EffectiveGroupIdentifier() uint32 {
	rv := objc.Send[uint32](x.ID, objc.Sel("effectiveGroupIdentifier"))
	return rv
}



// The effective user ID (EUID) of the connecting process.
//
// # Discussion
// 
// This attribute may be used by the listener delegate to accept or reject
// connections.
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCConnection/effectiveUserIdentifier
func (x NSXPCConnection) EffectiveUserIdentifier() uint32 {
	rv := objc.Send[uint32](x.ID, objc.Sel("effectiveUserIdentifier"))
	return rv
}



// The XPC connection was interrupted.
//
// See: https://developer.apple.com/documentation/foundation/nsxpcconnectioninterrupted-swift.var
func (x NSXPCConnection) NSXPCConnectionInterrupted() int {
	rv := objc.Send[int](x.ID, objc.Sel("NSXPCConnectionInterrupted"))
	return rv
}
func (x NSXPCConnection) SetNSXPCConnectionInterrupted(value int) {
	objc.Send[struct{}](x.ID, objc.Sel("setNSXPCConnectionInterrupted:"), value)
}



// The XPC connection was invalid.
//
// See: https://developer.apple.com/documentation/foundation/nsxpcconnectioninvalid-swift.var
func (x NSXPCConnection) NSXPCConnectionInvalid() int {
	rv := objc.Send[int](x.ID, objc.Sel("NSXPCConnectionInvalid"))
	return rv
}
func (x NSXPCConnection) SetNSXPCConnectionInvalid(value int) {
	objc.Send[struct{}](x.ID, objc.Sel("setNSXPCConnectionInvalid:"), value)
}



// The XPC connection reply was invalid.
//
// See: https://developer.apple.com/documentation/foundation/nsxpcconnectionreplyinvalid-swift.var
func (x NSXPCConnection) NSXPCConnectionReplyInvalid() int {
	rv := objc.Send[int](x.ID, objc.Sel("NSXPCConnectionReplyInvalid"))
	return rv
}
func (x NSXPCConnection) SetNSXPCConnectionReplyInvalid(value int) {
	objc.Send[struct{}](x.ID, objc.Sel("setNSXPCConnectionReplyInvalid:"), value)
}



// The lower bounds of XPC connection error code values.
//
// See: https://developer.apple.com/documentation/foundation/nsxpcconnectionerrorminimum-swift.var
func (x NSXPCConnection) NSXPCConnectionErrorMinimum() int {
	rv := objc.Send[int](x.ID, objc.Sel("NSXPCConnectionErrorMinimum"))
	return rv
}
func (x NSXPCConnection) SetNSXPCConnectionErrorMinimum(value int) {
	objc.Send[struct{}](x.ID, objc.Sel("setNSXPCConnectionErrorMinimum:"), value)
}



// The upper bounds of XPC connection error code values.
//
// See: https://developer.apple.com/documentation/foundation/nsxpcconnectionerrormaximum-swift.var
func (x NSXPCConnection) NSXPCConnectionErrorMaximum() int {
	rv := objc.Send[int](x.ID, objc.Sel("NSXPCConnectionErrorMaximum"))
	return rv
}
func (x NSXPCConnection) SetNSXPCConnectionErrorMaximum(value int) {
	objc.Send[struct{}](x.ID, objc.Sel("setNSXPCConnectionErrorMaximum:"), value)
}



// A code-signing requirement check failed.
//
// See: https://developer.apple.com/documentation/foundation/nsxpcconnectioncodesigningrequirementfailure-swift.var
func (x NSXPCConnection) NSXPCConnectionCodeSigningRequirementFailure() int {
	rv := objc.Send[int](x.ID, objc.Sel("NSXPCConnectionCodeSigningRequirementFailure"))
	return rv
}
func (x NSXPCConnection) SetNSXPCConnectionCodeSigningRequirementFailure(value int) {
	objc.Send[struct{}](x.ID, objc.Sel("setNSXPCConnectionCodeSigningRequirementFailure:"), value)
}














			// Protocol methods for NSXPCProxyCreating
			










// ScheduleSendBarrierBlockSync is a synchronous wrapper around [NSXPCConnection.ScheduleSendBarrierBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (x NSXPCConnection) ScheduleSendBarrierBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	x.ScheduleSendBarrierBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// RemoteObjectProxyWithErrorHandlerSync is a synchronous wrapper around [NSXPCConnection.RemoteObjectProxyWithErrorHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (x NSXPCConnection) RemoteObjectProxyWithErrorHandlerSync(ctx context.Context) error {
	done := make(chan error, 1)
	x.RemoteObjectProxyWithErrorHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SynchronousRemoteObjectProxyWithErrorHandlerSync is a synchronous wrapper around [NSXPCConnection.SynchronousRemoteObjectProxyWithErrorHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (x NSXPCConnection) SynchronousRemoteObjectProxyWithErrorHandlerSync(ctx context.Context) error {
	done := make(chan error, 1)
	x.SynchronousRemoteObjectProxyWithErrorHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}






