// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSXPCListener] class.
var (
	_NSXPCListenerClass     NSXPCListenerClass
	_NSXPCListenerClassOnce sync.Once
)

func getNSXPCListenerClass() NSXPCListenerClass {
	_NSXPCListenerClassOnce.Do(func() {
		_NSXPCListenerClass = NSXPCListenerClass{class: objc.GetClass("NSXPCListener")}
	})
	return _NSXPCListenerClass
}

// GetNSXPCListenerClass returns the class object for NSXPCListener.
func GetNSXPCListenerClass() NSXPCListenerClass {
	return getNSXPCListenerClass()
}

type NSXPCListenerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSXPCListenerClass) Alloc() NSXPCListener {
	rv := objc.Send[NSXPCListener](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A listener that waits for new incoming connections, configures them, and
// accepts or rejects them.
//
// # Overview
// 
// Each XPC service, launchd agent, or launchd daemon typically has at least
// one [NSXPCListener] object that listens for connections to a specified
// service name. Each listener must have a delegate that conforms to the
// [NSXPCListenerDelegate] protocol. When the listener receives a new
// connection request, it creates a new [NSXPCConnection] object, then asks
// the delegate to inspect, configure, and resume the connection object by
// calling the delegate’s [ListenerShouldAcceptNewConnection] method.
//
// # Creating a listener
//
//   - [NSXPCListener.InitWithMachServiceName]: Initializes a listener in a LaunchAgent or LaunchDaemon which has a name advertised in a `launchd.Plist()` file.
//
// # Working with a delegate
//
//   - [NSXPCListener.Delegate]: The delegate for the listener.
//   - [NSXPCListener.SetDelegate]
//
// # Providing access to clients
//
//   - [NSXPCListener.Endpoint]: Returns an endpoint object that may be sent over an existing connection.
//
// # Managing connection state
//
//   - [NSXPCListener.Activate]: Activates the listener.
//   - [NSXPCListener.Resume]: Starts processing of incoming requests.
//   - [NSXPCListener.Invalidate]: Invalidates the listener.
//   - [NSXPCListener.Suspend]: Suspends the listener.
//
// # Working with code-signing
//
//   - [NSXPCListener.SetConnectionCodeSigningRequirement]: Sets the code signing requirement for connections to this listener.
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCListener
type NSXPCListener struct {
	objectivec.Object
}

// NSXPCListenerFromID constructs a [NSXPCListener] from an objc.ID.
//
// A listener that waits for new incoming connections, configures them, and
// accepts or rejects them.
func NSXPCListenerFromID(id objc.ID) NSXPCListener {
	return NSXPCListener{objectivec.Object{ID: id}}
}
// NOTE: NSXPCListener adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSXPCListener] class.
//
// # Creating a listener
//
//   - [INSXPCListener.InitWithMachServiceName]: Initializes a listener in a LaunchAgent or LaunchDaemon which has a name advertised in a `launchd.Plist()` file.
//
// # Working with a delegate
//
//   - [INSXPCListener.Delegate]: The delegate for the listener.
//   - [INSXPCListener.SetDelegate]
//
// # Providing access to clients
//
//   - [INSXPCListener.Endpoint]: Returns an endpoint object that may be sent over an existing connection.
//
// # Managing connection state
//
//   - [INSXPCListener.Activate]: Activates the listener.
//   - [INSXPCListener.Resume]: Starts processing of incoming requests.
//   - [INSXPCListener.Invalidate]: Invalidates the listener.
//   - [INSXPCListener.Suspend]: Suspends the listener.
//
// # Working with code-signing
//
//   - [INSXPCListener.SetConnectionCodeSigningRequirement]: Sets the code signing requirement for connections to this listener.
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCListener
type INSXPCListener interface {
	objectivec.IObject

	// Topic: Creating a listener

	// Initializes a listener in a LaunchAgent or LaunchDaemon which has a name advertised in a `launchd.Plist()` file.
	InitWithMachServiceName(name string) NSXPCListener

	// Topic: Working with a delegate

	// The delegate for the listener.
	Delegate() NSXPCListenerDelegate
	SetDelegate(value NSXPCListenerDelegate)

	// Topic: Providing access to clients

	// Returns an endpoint object that may be sent over an existing connection.
	Endpoint() INSXPCListenerEndpoint

	// Topic: Managing connection state

	// Activates the listener.
	Activate()
	// Starts processing of incoming requests.
	Resume()
	// Invalidates the listener.
	Invalidate()
	// Suspends the listener.
	Suspend()

	// Topic: Working with code-signing

	// Sets the code signing requirement for connections to this listener.
	SetConnectionCodeSigningRequirement(requirement string)
}

// Init initializes the instance.
func (x NSXPCListener) Init() NSXPCListener {
	rv := objc.Send[NSXPCListener](x.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (x NSXPCListener) Autorelease() NSXPCListener {
	rv := objc.Send[NSXPCListener](x.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSXPCListener creates a new NSXPCListener instance.
func NewNSXPCListener() NSXPCListener {
	class := getNSXPCListenerClass()
	rv := objc.Send[NSXPCListener](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a listener in a LaunchAgent or LaunchDaemon which has a name
// advertised in a `launchd.Plist()` file.
//
// # Discussion
// 
// For example, you might use this in an agent launched by launchd with a
// `launchd.Plist()` contained in `~/Library/LaunchAgents`, or a daemon
// launched by launchd with a `launchd.Plist()` contained in
// `/Library/LaunchDaemons`.
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCListener/init(machServiceName:)
func NewXPCListenerWithMachServiceName(name string) NSXPCListener {
	instance := getNSXPCListenerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMachServiceName:"), objc.String(name))
	return NSXPCListenerFromID(rv)
}

// Initializes a listener in a LaunchAgent or LaunchDaemon which has a name
// advertised in a `launchd.Plist()` file.
//
// # Discussion
// 
// For example, you might use this in an agent launched by launchd with a
// `launchd.Plist()` contained in `~/Library/LaunchAgents`, or a daemon
// launched by launchd with a `launchd.Plist()` contained in
// `/Library/LaunchDaemons`.
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCListener/init(machServiceName:)
func (x NSXPCListener) InitWithMachServiceName(name string) NSXPCListener {
	rv := objc.Send[NSXPCListener](x.ID, objc.Sel("initWithMachServiceName:"), objc.String(name))
	return rv
}

// Activates the listener.
//
// # Discussion
// 
// Connections start in an inactive state. You must call [Activate] on a
// connection before it can send or receive any messages.
// 
// Calling [Activate] on an active connection has no effect.
// 
// For backward compatibility reasons, calling [Resume] on an inactive and
// otherwise not suspended [NSXPCListener] has the same effect as calling
// [Activate]. For new code, prefer [Activate].
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCListener/activate()
func (x NSXPCListener) Activate() {
	objc.Send[objc.ID](x.ID, objc.Sel("activate"))
}

// Starts processing of incoming requests.
//
// # Discussion
// 
// All listeners start suspended and must be resumed before they begin
// processing incoming requests.
// 
// If called on the [ServiceListener] object, this method never returns.
// Therefore, you should call it as the last step inside the XPC service’s
// `main` function after setting up any desired initial state and configuring
// the listener itself.
// 
// If called on any other [NSXPCListener], the connection is resumed, and the
// method returns immediately.
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCListener/resume()
func (x NSXPCListener) Resume() {
	objc.Send[objc.ID](x.ID, objc.Sel("resume"))
}

// Invalidates the listener.
//
// # Discussion
// 
// After calling this method, no more connections are created. Once a listener
// is invalidated it may not be resumed or suspended.
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCListener/invalidate()
func (x NSXPCListener) Invalidate() {
	objc.Send[objc.ID](x.ID, objc.Sel("invalidate"))
}

// Suspends the listener.
//
// # Discussion
// 
// As you cannot invalidate a suspended listener, every call to [Suspend] that
// you make must be balanced by a call to [Resume].
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCListener/suspend()
func (x NSXPCListener) Suspend() {
	objc.Send[objc.ID](x.ID, objc.Sel("suspend"))
}

// Sets the code signing requirement for connections to this listener.
//
// requirement: A string that describes requirements expected of the connection peer. See
// [Code Signing Guide] for more information on the code signing format.
// //
// [Code Signing Guide]: https://developer.apple.com/library/archive/documentation/Security/Conceptual/CodeSigningGuide/
//
// # Discussion
// 
// Use this method to enforce a code-signing requirement on incoming XPC
// connections.
// 
// The following example shows how a listener can ensure that the XPC client
// service on the other end of a connection has a specific entitlement.
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCListener/setConnectionCodeSigningRequirement(_:)
func (x NSXPCListener) SetConnectionCodeSigningRequirement(requirement string) {
	objc.Send[objc.ID](x.ID, objc.Sel("setConnectionCodeSigningRequirement:"), objc.String(requirement))
}

// Returns the singleton listener used to listen for incoming connections in
// an XPC service.
//
// # Discussion
// 
// Calling the `resume` method on the returned object starts the listener and
// never returns. This method is typically called at the end of your `main`
// function.
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCListener/service()
func (_NSXPCListenerClass NSXPCListenerClass) ServiceListener() NSXPCListener {
	rv := objc.Send[objc.ID](objc.ID(_NSXPCListenerClass.class), objc.Sel("serviceListener"))
	return NSXPCListenerFromID(rv)
}

// Returns a new anonymous listener connection.
//
// # Discussion
// 
// Other processes can connect to this listener by passing this listener
// object’s [NSXPCListenerEndpoint] to the [InitWithListenerEndpoint] method
// of an [NSXPCConnection] object.
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCListener/anonymous()
func (_NSXPCListenerClass NSXPCListenerClass) AnonymousListener() NSXPCListener {
	rv := objc.Send[objc.ID](objc.ID(_NSXPCListenerClass.class), objc.Sel("anonymousListener"))
	return NSXPCListenerFromID(rv)
}

// The delegate for the listener.
//
// # Discussion
// 
// If no delegate is set, all new connections are rejected. See the
// documentation for [NSXPCListenerDelegate] for implementation details.
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCListener/delegate
func (x NSXPCListener) Delegate() NSXPCListenerDelegate {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("delegate"))
	return NSXPCListenerDelegateObjectFromID(rv)
}
func (x NSXPCListener) SetDelegate(value NSXPCListenerDelegate) {
	objc.Send[struct{}](x.ID, objc.Sel("setDelegate:"), value)
}

// Returns an endpoint object that may be sent over an existing connection.
//
// # Discussion
// 
// The receiver of the endpoint can use this object to create a new connection
// to this [NSXPCListener] object. The resulting [NSXPCListenerEndpoint]
// object uniquely names this listener object across connections.
//
// See: https://developer.apple.com/documentation/Foundation/NSXPCListener/endpoint
func (x NSXPCListener) Endpoint() INSXPCListenerEndpoint {
	rv := objc.Send[objc.ID](x.ID, objc.Sel("endpoint"))
	return NSXPCListenerEndpointFromID(objc.ID(rv))
}

