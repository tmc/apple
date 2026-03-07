// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSProxy] class.
var (
	_NSProxyClass     NSProxyClass
	_NSProxyClassOnce sync.Once
)

func getNSProxyClass() NSProxyClass {
	_NSProxyClassOnce.Do(func() {
		_NSProxyClass = NSProxyClass{class: objc.GetClass("NSProxy")}
	})
	return _NSProxyClass
}

// GetNSProxyClass returns the class object for NSProxy.
func GetNSProxyClass() NSProxyClass {
	return getNSProxyClass()
}

type NSProxyClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSProxyClass) Alloc() NSProxy {
	rv := objc.Send[NSProxy](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// An abstract superclass defining an API for objects that act as stand-ins
// for other objects or for objects that don’t exist yet.
//
// # Overview
// 
// Typically, a message to a proxy is forwarded to the real object or causes
// the proxy to load (or transform itself into) the real object. Subclasses of
// [NSProxy] can be used to implement transparent distributed messaging (for
// example, [NSDistantObject]) or for lazy instantiation of objects that are
// expensive to create.
// 
// [NSProxy] implements the basic methods required of a root class, including
// those defined in the [NSObjectProtocol] protocol. However, as an abstract
// class it doesn’t provide an initialization method, and it raises an
// exception upon receiving any message it doesn’t respond to. A concrete
// subclass must therefore provide an initialization or creation method and
// override the [NSProxy.ForwardInvocation] and [NSProxy.MethodSignatureForSelector] methods
// to handle messages that it doesn’t implement itself. A subclass’s
// implementation of [NSProxy.ForwardInvocation] should do whatever is needed to
// process the invocation, such as forwarding the invocation over the network
// or loading the real object and passing it the invocation.
// [NSProxy.MethodSignatureForSelector] is required to provide argument type
// information for a given message; a subclass’s implementation should be
// able to determine the argument types for the messages it needs to forward
// and should construct an [NSMethodSignature] object accordingly. See the
// [NSDistantObject], [NSInvocation], and [NSMethodSignature] class
// specifications for more information.
//
// [NSDistantObject]: https://developer.apple.com/documentation/Foundation/NSDistantObject
// [NSObjectProtocol]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol
//
// # Deallocating Instances
//
//   - [NSProxy.Dealloc]: Deallocates the memory occupied by the receiver.
//
// # Finalizing an Object
//
//   - [NSProxy.Finalize]: The garbage collector invokes this method on the receiver before disposing of the memory it uses.
//
// # Handling Unimplemented Methods
//
//   - [NSProxy.ForwardInvocation]: Passes a given invocation to the real object the proxy represents.
//
// # Describing a Proxy Class or Object
//
//   - [NSProxy.Description]: A string containing the real class name and the id of the receiver as a hexadecimal number.
//   - [NSProxy.DebugDescription]
//
// See: https://developer.apple.com/documentation/Foundation/NSProxy
type NSProxy struct {
	objectivec.Object
}

// NSProxyFromID constructs a [NSProxy] from an objc.ID.
//
// An abstract superclass defining an API for objects that act as stand-ins
// for other objects or for objects that don’t exist yet.
func NSProxyFromID(id objc.ID) NSProxy {
	return NSProxy{objectivec.Object{id}}
}
// NOTE: NSProxy adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSProxy] class.
//
// # Deallocating Instances
//
//   - [INSProxy.Dealloc]: Deallocates the memory occupied by the receiver.
//
// # Finalizing an Object
//
//   - [INSProxy.Finalize]: The garbage collector invokes this method on the receiver before disposing of the memory it uses.
//
// # Handling Unimplemented Methods
//
//   - [INSProxy.ForwardInvocation]: Passes a given invocation to the real object the proxy represents.
//
// # Describing a Proxy Class or Object
//
//   - [INSProxy.Description]: A string containing the real class name and the id of the receiver as a hexadecimal number.
//   - [INSProxy.DebugDescription]
//
// See: https://developer.apple.com/documentation/Foundation/NSProxy
type INSProxy interface {
	objectivec.IObject

	// Topic: Deallocating Instances

	// Deallocates the memory occupied by the receiver.
	Dealloc()

	// Topic: Finalizing an Object

	// The garbage collector invokes this method on the receiver before disposing of the memory it uses.
	Finalize()

	// Topic: Handling Unimplemented Methods

	// Passes a given invocation to the real object the proxy represents.
	ForwardInvocation(invocation INSInvocation)

	// Topic: Describing a Proxy Class or Object

	// A string containing the real class name and the id of the receiver as a hexadecimal number.
	Description() string
	DebugDescription() string

	AllowsWeakReference() bool
	RetainWeakReference() bool
	// Raises [NSInvalidArgumentException]. Override this method in your concrete subclass to return a proper [NSMethodSignature] object for the given selector and the class your proxy objects stand in for.
	MethodSignatureForSelector(sel objc.SEL) INSMethodSignature
}





// Init initializes the instance.
func (p NSProxy) Init() NSProxy {
	rv := objc.Send[NSProxy](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSProxy) Autorelease() NSProxy {
	rv := objc.Send[NSProxy](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSProxy creates a new NSProxy instance.
func NewNSProxy() NSProxy {
	class := getNSProxyClass()
	rv := objc.Send[NSProxy](objc.ID(class.class), objc.Sel("new"))
	return rv
}










// Deallocates the memory occupied by the receiver.
//
// # Discussion
// 
// This method behaves as described in the [NSObject] class specification
// under the [dealloc] instance method.
//
// [dealloc]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/dealloc
//
// See: https://developer.apple.com/documentation/Foundation/NSProxy/dealloc()
func (p NSProxy) Dealloc() {
	objc.Send[objc.ID](p.ID, objc.Sel("dealloc"))
}

// The garbage collector invokes this method on the receiver before disposing
// of the memory it uses.
//
// # Discussion
// 
// This method behaves as described in the [NSObject] class specification
// under the [finalize()] instance method. Note that a `finalize` method must
// be thread-safe.
//
// [finalize()]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/finalize()
//
// See: https://developer.apple.com/documentation/Foundation/NSProxy/finalize()
func (p NSProxy) Finalize() {
	objc.Send[objc.ID](p.ID, objc.Sel("finalize"))
}

// Passes a given invocation to the real object the proxy represents.
//
// invocation: The invocation to forward.
//
// # Discussion
// 
// [NSProxy]’s implementation merely raises [NSInvalidArgumentException].
// Override this method in your subclass to handle `invocation` appropriately,
// at the very least by setting its return value.
// 
// For example, if your proxy merely forwards messages to an instance variable
// named `realObject`, it can implement [ForwardInvocation] like this:
//
// See: https://developer.apple.com/documentation/Foundation/NSProxy/forwardInvocation(_:)
func (p NSProxy) ForwardInvocation(invocation INSInvocation) {
	objc.Send[objc.ID](p.ID, objc.Sel("forwardInvocation:"), invocation)
}

// See: https://developer.apple.com/documentation/Foundation/NSProxy/allowsWeakReference
func (p NSProxy) AllowsWeakReference() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("allowsWeakReference"))
	return rv
}

// See: https://developer.apple.com/documentation/Foundation/NSProxy/retainWeakReference
func (p NSProxy) RetainWeakReference() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("retainWeakReference"))
	return rv
}

// Raises [NSInvalidArgumentException]. Override this method in your concrete
// subclass to return a proper [NSMethodSignature] object for the given
// selector and the class your proxy objects stand in for.
//
// sel: The selector for which to return a method signature.
//
// # Return Value
// 
// Not applicable. The implementation provided by [NSProxy] raises an
// exception.
//
// # Discussion
// 
// Be sure to avoid an infinite loop when necessary by checking that `sel`
// isn’t the selector for this method itself and by not sending any message
// that might invoke this method.
// 
// For example, if your proxy merely forwards messages to an instance variable
// named `realObject`, it can implement [MethodSignatureForSelector] like
// this:
//
// See: https://developer.apple.com/documentation/Foundation/NSProxy/methodSignatureForSelector:
func (p NSProxy) MethodSignatureForSelector(sel objc.SEL) INSMethodSignature {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("methodSignatureForSelector:"), sel)
	return NSMethodSignatureFromID(rv)
}












// A string containing the real class name and the id of the receiver as a
// hexadecimal number.
//
// See: https://developer.apple.com/documentation/Foundation/NSProxy/description
func (p NSProxy) Description() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("description"))
	return NSStringFromID(rv).String()
}



// See: https://developer.apple.com/documentation/Foundation/NSProxy/debugDescription
func (p NSProxy) DebugDescription() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("debugDescription"))
	return NSStringFromID(rv).String()
}


















