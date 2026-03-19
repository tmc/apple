// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [PortMessage] class.
var (
	_PortMessageClass     PortMessageClass
	_PortMessageClassOnce sync.Once
)

func getPortMessageClass() PortMessageClass {
	_PortMessageClassOnce.Do(func() {
		_PortMessageClass = PortMessageClass{class: objc.GetClass("NSPortMessage")}
	})
	return _PortMessageClass
}

// GetPortMessageClass returns the class object for NSPortMessage.
func GetPortMessageClass() PortMessageClass {
	return getPortMessageClass()
}

type PortMessageClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (pc PortMessageClass) Alloc() PortMessage {
	rv := objc.Send[PortMessage](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// A low-level, operating system-independent type for inter-application (and
// inter-thread) messages.
//
// # Overview
// 
// Port messages are used primarily by the distributed objects system. You
// should implement inter-application communication using distributed objects
// whenever possible and use [NSPortMessage] only when necessary.
// 
// An [NSPortMessage] object has three major parts: the send and receive
// ports, which are [NSPort] objects that link the sender of the message to
// the receiver, and the components, which form the body of the message. The
// components are held as an [NSArray] object containing [NSData] and [NSPort]
// objects. The [SendBeforeDate] message sends the components out through the
// send port; any replies to the message arrive on the receive port. See the
// [NSPort] class specification for information on handling incoming messages.
// 
// An [NSPortMessage] instance can be initialized with a pair of [NSPort]
// objects and an array of components. A port message’s body can contain
// only [NSPort] objects or [NSData] objects. In the distributed objects
// system the byte/character arrays are usually encoded [NSInvocation] objects
// that are being forwarded from a proxy to the corresponding real object.
// 
// An [NSPortMessage] object also maintains a message identifier, which can be
// used to indicate the class of a message, such as an Objective-C method
// invocation, a connection request, an error, and so on. Use the [Msgid] and
// [Msgid] methods to access the identifier.
//
// # Creating Instances
//
//   - [PortMessage.InitWithSendPortReceivePortComponents]: Initializes a newly allocated [NSPortMessage] object to send given data on a given port and to receiver replies on another given port.
//
// # Sending the Message
//
//   - [PortMessage.SendBeforeDate]: Attempts to send the message before the specified date.
//
// # Getting the Components
//
//   - [PortMessage.Components]: Returns the data components of the receiver.
//
// # Getting the Ports
//
//   - [PortMessage.ReceivePort]: For an outgoing message, returns the port on which replies to the receiver will arrive. For an incoming message, returns the port the receiver did arrive on.
//   - [PortMessage.SendPort]: For an outgoing message, returns the port the receiver will send itself through. For an incoming message, returns the port replies to the receiver should be sent through.
//
// # Accessing the Message ID
//
//   - [PortMessage.Msgid]: Returns the identifier for the receiver.
//   - [PortMessage.SetMsgid]
//
// See: https://developer.apple.com/documentation/Foundation/PortMessage
type PortMessage struct {
	objectivec.Object
}

// PortMessageFromID constructs a [PortMessage] from an objc.ID.
//
// A low-level, operating system-independent type for inter-application (and
// inter-thread) messages.
func PortMessageFromID(id objc.ID) PortMessage {
	return NSPortMessage{objectivec.Object{ID: id}}
}

// NSPortMessageFromID is an alias for [PortMessageFromID] for cross-framework compatibility.
func NSPortMessageFromID(id objc.ID) PortMessage { return PortMessageFromID(id) }
// NOTE: PortMessage adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [PortMessage] class.
//
// # Creating Instances
//
//   - [IPortMessage.InitWithSendPortReceivePortComponents]: Initializes a newly allocated [NSPortMessage] object to send given data on a given port and to receiver replies on another given port.
//
// # Sending the Message
//
//   - [IPortMessage.SendBeforeDate]: Attempts to send the message before the specified date.
//
// # Getting the Components
//
//   - [IPortMessage.Components]: Returns the data components of the receiver.
//
// # Getting the Ports
//
//   - [IPortMessage.ReceivePort]: For an outgoing message, returns the port on which replies to the receiver will arrive. For an incoming message, returns the port the receiver did arrive on.
//   - [IPortMessage.SendPort]: For an outgoing message, returns the port the receiver will send itself through. For an incoming message, returns the port replies to the receiver should be sent through.
//
// # Accessing the Message ID
//
//   - [IPortMessage.Msgid]: Returns the identifier for the receiver.
//   - [IPortMessage.SetMsgid]
//
// See: https://developer.apple.com/documentation/Foundation/PortMessage
type IPortMessage interface {
	objectivec.IObject

	// Topic: Creating Instances

	// Initializes a newly allocated [NSPortMessage] object to send given data on a given port and to receiver replies on another given port.
	InitWithSendPortReceivePortComponents(sendPort INSPort, replyPort INSPort, components INSArray) PortMessage

	// Topic: Sending the Message

	// Attempts to send the message before the specified date.
	SendBeforeDate(date INSDate) bool

	// Topic: Getting the Components

	// Returns the data components of the receiver.
	Components() INSArray

	// Topic: Getting the Ports

	// For an outgoing message, returns the port on which replies to the receiver will arrive. For an incoming message, returns the port the receiver did arrive on.
	ReceivePort() INSPort
	// For an outgoing message, returns the port the receiver will send itself through. For an incoming message, returns the port replies to the receiver should be sent through.
	SendPort() INSPort

	// Topic: Accessing the Message ID

	// Returns the identifier for the receiver.
	Msgid() uint32
	SetMsgid(value uint32)
}

// Init initializes the instance.
func (p PortMessage) Init() PortMessage {
	rv := objc.Send[PortMessage](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p PortMessage) Autorelease() PortMessage {
	rv := objc.Send[PortMessage](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewPortMessage creates a new PortMessage instance.
func NewPortMessage() PortMessage {
	class := getPortMessageClass()
	rv := objc.Send[PortMessage](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a newly allocated [NSPortMessage] object to send given data on
// a given port and to receiver replies on another given port.
//
// sendPort: The port on which the message is sent.
//
// replyPort: The port on which replies to the message arrive.
//
// components: The data to send in the message. `components` should contain only [NSData]
// and [NSPort] objects, and the contents of the [NSData] objects should be in
// network byte order.
//
// # Return Value
// 
// An [NSPortMessage] object initialized to send `components` on `sendPort`
// and to receiver replies on `receivePort`.
//
// # Discussion
// 
// An [NSPortMessage] object initialized with this method has a message
// identifier of 0.
// 
// This is the designated initializer for [NSPortMessage].
//
// See: https://developer.apple.com/documentation/Foundation/PortMessage/init(send:receive:components:)
func NewPortMessageWithSendPortReceivePortComponents(sendPort INSPort, replyPort INSPort, components INSArray) PortMessage {
	instance := getPortMessageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSendPort:receivePort:components:"), sendPort, replyPort, components)
	return PortMessageFromID(rv)
}

// Initializes a newly allocated [NSPortMessage] object to send given data on
// a given port and to receiver replies on another given port.
//
// sendPort: The port on which the message is sent.
//
// replyPort: The port on which replies to the message arrive.
//
// components: The data to send in the message. `components` should contain only [NSData]
// and [NSPort] objects, and the contents of the [NSData] objects should be in
// network byte order.
//
// # Return Value
// 
// An [NSPortMessage] object initialized to send `components` on `sendPort`
// and to receiver replies on `receivePort`.
//
// # Discussion
// 
// An [NSPortMessage] object initialized with this method has a message
// identifier of 0.
// 
// This is the designated initializer for [NSPortMessage].
//
// See: https://developer.apple.com/documentation/Foundation/PortMessage/init(send:receive:components:)
func (p PortMessage) InitWithSendPortReceivePortComponents(sendPort INSPort, replyPort INSPort, components INSArray) PortMessage {
	rv := objc.Send[PortMessage](p.ID, objc.Sel("initWithSendPort:receivePort:components:"), sendPort, replyPort, components)
	return rv
}
// Attempts to send the message before the specified date.
//
// date: The instant before which the message should be sent.
//
// # Return Value
// 
// [true] if the operation is successful, otherwise [false] (for example, if
// the operation times out).
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// If an error other than a time out occurs, this method could raise an
// [NSInvalidSendPortException], [NSInvalidReceivePortException], or an
// [NSPortSendException], depending on the type of send port and the type of
// error.
// 
// If the message cannot be sent immediately, the sending thread blocks until
// either the message is sent or `aDate` is reached. Sent messages are queued
// to minimize blocking, but failure can occur if multiple messages are sent
// to a port faster than the port’s owner can receive them, causing the
// queue to fill up. Therefore, select a value for `aDate` that provides
// enough time for the message to be processed before the next message is
// sent. See the [NSPort] class specification for information on receiving a
// port message.
//
// See: https://developer.apple.com/documentation/Foundation/PortMessage/send(before:)
func (p PortMessage) SendBeforeDate(date INSDate) bool {
	rv := objc.Send[bool](p.ID, objc.Sel("sendBeforeDate:"), date)
	return rv
}

// Returns the data components of the receiver.
//
// # Return Value
// 
// The data components of the receiver. See [NSPortMessage] for more
// information.
//
// See: https://developer.apple.com/documentation/Foundation/PortMessage/components
func (p PortMessage) Components() INSArray {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("components"))
	return NSArrayFromID(objc.ID(rv))
}
// For an outgoing message, returns the port on which replies to the receiver
// will arrive. For an incoming message, returns the port the receiver did
// arrive on.
//
// # Return Value
// 
// For an outgoing message, the port on which replies to the receiver will
// arrive. For an incoming message, the port the receiver did arrive on.
//
// See: https://developer.apple.com/documentation/Foundation/PortMessage/receivePort
func (p PortMessage) ReceivePort() INSPort {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("receivePort"))
	return NSPortFromID(objc.ID(rv))
}
// For an outgoing message, returns the port the receiver will send itself
// through. For an incoming message, returns the port replies to the receiver
// should be sent through.
//
// # Return Value
// 
// For an outgoing message, the port the receiver will send itself through
// when it receives a [SendBeforeDate] message. For an incoming message, the
// port replies to the receiver should be sent through.
//
// See: https://developer.apple.com/documentation/Foundation/PortMessage/sendPort
func (p PortMessage) SendPort() INSPort {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("sendPort"))
	return NSPortFromID(objc.ID(rv))
}
// Returns the identifier for the receiver.
//
// # Return Value
// 
// The identifier for the receiver.
// 
// # Discussion
// 
// Cooperating applications can use this to define different types of
// messages, such as connection requests, RPCs, errors, and so on.
//
// See: https://developer.apple.com/documentation/Foundation/PortMessage/msgid
func (p PortMessage) Msgid() uint32 {
	rv := objc.Send[uint32](p.ID, objc.Sel("msgid"))
	return rv
}
func (p PortMessage) SetMsgid(value uint32) {
	objc.Send[struct{}](p.ID, objc.Sel("setMsgid:"), value)
}

