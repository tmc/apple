// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSProtocolChecker] class.
var (
	_NSProtocolCheckerClass     NSProtocolCheckerClass
	_NSProtocolCheckerClassOnce sync.Once
)

func getNSProtocolCheckerClass() NSProtocolCheckerClass {
	_NSProtocolCheckerClassOnce.Do(func() {
		_NSProtocolCheckerClass = NSProtocolCheckerClass{class: objc.GetClass("NSProtocolChecker")}
	})
	return _NSProtocolCheckerClass
}

// GetNSProtocolCheckerClass returns the class object for NSProtocolChecker.
func GetNSProtocolCheckerClass() NSProtocolCheckerClass {
	return getNSProtocolCheckerClass()
}

type NSProtocolCheckerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSProtocolCheckerClass) Alloc() NSProtocolChecker {
	rv := objc.Send[NSProtocolChecker](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that restricts the messages that can be sent to another object
// (referred to as the checker’s delegate).
//
// # Overview
// 
// A [NSProtocolChecker] object can be particularly useful when an object with
// many methods, only a few of which ought to be remotely accessible, is made
// available using the distributed objects system.
// 
// A protocol checker acts as a kind of proxy; when it receives a message that
// is in its designated protocol, it forwards the message to its target and
// consequently appears to be the target object itself. However, when it
// receives a message not in its protocol, it raises an
// [invalidArgumentException] to indicate that the message isn’t allowed,
// whether or not the target object implements the method.
// 
// Typically, an object that is to be distributed (yet must restrict messages)
// creates an [NSProtocolChecker] for itself and returns the checker rather
// than returning itself in response to any messages. The object might also
// register the checker as the root object of an NSConnection.
// 
// The object should be careful about vending references to `self`—the
// protocol checker will convert a return value of `self` to indicate the
// checker rather than the object for any messages forwarded by the checker,
// but direct references to the object (bypassing the checker) could be passed
// around by other objects.
//
// [invalidArgumentException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/invalidArgumentException
//
// # Creating a checker
//
//   - [NSProtocolChecker.InitWithTargetProtocol]: Initializes a newly allocated [NSProtocolChecker] instance that will forward any messages in `aProtocol` to `anObject`, the protocol checker’s target.
//
// # Getting information
//
//   - [NSProtocolChecker.Protocol]: Returns the protocol object the receiver uses.
//   - [NSProtocolChecker.Target]: Returns the target of the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSProtocolChecker
type NSProtocolChecker struct {
	NSProxy
}

// NSProtocolCheckerFromID constructs a [NSProtocolChecker] from an objc.ID.
//
// An object that restricts the messages that can be sent to another object
// (referred to as the checker’s delegate).
func NSProtocolCheckerFromID(id objc.ID) NSProtocolChecker {
	return NSProtocolChecker{NSProxy: NSProxyFromID(id)}
}
// NOTE: NSProtocolChecker adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSProtocolChecker] class.
//
// # Creating a checker
//
//   - [INSProtocolChecker.InitWithTargetProtocol]: Initializes a newly allocated [NSProtocolChecker] instance that will forward any messages in `aProtocol` to `anObject`, the protocol checker’s target.
//
// # Getting information
//
//   - [INSProtocolChecker.Protocol]: Returns the protocol object the receiver uses.
//   - [INSProtocolChecker.Target]: Returns the target of the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSProtocolChecker
type INSProtocolChecker interface {
	INSProxy

	// Topic: Creating a checker

	// Initializes a newly allocated [NSProtocolChecker] instance that will forward any messages in `aProtocol` to `anObject`, the protocol checker’s target.
	InitWithTargetProtocol(anObject objectivec.Object, aProtocol *objectivec.Protocol) NSProtocolChecker

	// Topic: Getting information

	// Returns the protocol object the receiver uses.
	Protocol() *objectivec.Protocol
	// Returns the target of the receiver.
	Target() objectivec.Object

	// Name of an exception that occurs when you pass an invalid argument to a method, such as a 
	InvalidArgumentException() NSExceptionName
}

// Init initializes the instance.
func (p NSProtocolChecker) Init() NSProtocolChecker {
	rv := objc.Send[NSProtocolChecker](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSProtocolChecker) Autorelease() NSProtocolChecker {
	rv := objc.Send[NSProtocolChecker](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSProtocolChecker creates a new NSProtocolChecker instance.
func NewNSProtocolChecker() NSProtocolChecker {
	class := getNSProtocolCheckerClass()
	rv := objc.Send[NSProtocolChecker](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a newly allocated [NSProtocolChecker] instance that will
// forward any messages in `aProtocol` to `anObject`, the protocol checker’s
// target.
//
// # Discussion
// 
// Thus, the checker can be vended in lieu of `anObject` to restrict the
// messages that can be sent to `anObject`. If `anObject` is allowed to be
// freed or dereferenced by clients, the `free` method should be included in
// `aProtocol`.
//
// See: https://developer.apple.com/documentation/Foundation/NSProtocolChecker/init(target:protocol:)
func NewProtocolCheckerWithTargetProtocol(anObject objectivec.Object, aProtocol *objectivec.Protocol) NSProtocolChecker {
	instance := getNSProtocolCheckerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTarget:protocol:"), anObject, aProtocol)
	return NSProtocolCheckerFromID(rv)
}

// Initializes a newly allocated [NSProtocolChecker] instance that will
// forward any messages in `aProtocol` to `anObject`, the protocol checker’s
// target.
//
// # Discussion
// 
// Thus, the checker can be vended in lieu of `anObject` to restrict the
// messages that can be sent to `anObject`. If `anObject` is allowed to be
// freed or dereferenced by clients, the `free` method should be included in
// `aProtocol`.
//
// See: https://developer.apple.com/documentation/Foundation/NSProtocolChecker/init(target:protocol:)
func (p NSProtocolChecker) InitWithTargetProtocol(anObject objectivec.Object, aProtocol *objectivec.Protocol) NSProtocolChecker {
	rv := objc.Send[NSProtocolChecker](p.ID, objc.Sel("initWithTarget:protocol:"), anObject, aProtocol)
	return rv
}

// Allocates and initializes an [NSProtocolChecker] instance that will forward
// any messages in `aProtocol` to `anObject`, the protocol checker’s target.
//
// # Discussion
// 
// Thus, the checker can be vended in lieu of `anObject` to restrict the
// messages that can be sent to `anObject`. Returns the new instance.
//
// See: https://developer.apple.com/documentation/Foundation/NSProtocolChecker/protocolCheckerWithTarget:protocol:
func (_NSProtocolCheckerClass NSProtocolCheckerClass) ProtocolCheckerWithTargetProtocol(anObject objectivec.Object, aProtocol *objectivec.Protocol) NSProtocolChecker {
	rv := objc.Send[objc.ID](objc.ID(_NSProtocolCheckerClass.class), objc.Sel("protocolCheckerWithTarget:protocol:"), anObject, aProtocol)
	return NSProtocolCheckerFromID(rv)
}

// Returns the protocol object the receiver uses.
//
// See: https://developer.apple.com/documentation/Foundation/NSProtocolChecker/protocol
func (p NSProtocolChecker) Protocol() *objectivec.Protocol {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("protocol"))
	if rv == 0 {
		return nil
	}
	val := objectivec.ProtocolFromID(objc.ID(rv))
	return &val
}
// Returns the target of the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSProtocolChecker/target
func (p NSProtocolChecker) Target() objectivec.Object {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("target"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
// Name of an exception that occurs when you pass an invalid argument to a
// method, such as a
//
// See: https://developer.apple.com/documentation/foundation/nsexceptionname/invalidargumentexception
func (p NSProtocolChecker) InvalidArgumentException() NSExceptionName {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("NSInvalidArgumentException"))
	return NSExceptionName(NSStringFromID(rv).String())
}

