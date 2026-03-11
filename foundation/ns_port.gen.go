// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [Port] class.
var (
	_PortClass     PortClass
	_PortClassOnce sync.Once
)

func getPortClass() PortClass {
	_PortClassOnce.Do(func() {
		_PortClass = PortClass{class: objc.GetClass("NSPort")}
	})
	return _PortClass
}

// GetPortClass returns the class object for NSPort.
func GetPortClass() PortClass {
	return getPortClass()
}

type PortClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (pc PortClass) Alloc() Port {
	rv := objc.Send[Port](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}







// An abstract class that represents a communication channel.
//
// # Overview
// 
// Communication occurs between [NSPort] objects, which typically reside in
// different threads or tasks. The distributed objects system uses [NSPort]
// objects to send [NSPortMessage] objects back and forth. Implement
// interapplication communication using distributed objects whenever possible
// and use [NSPort] objects only when necessary.
// 
// To receive incoming messages, add [NSPort] objects to an instance of
// [NSRunLoop] as input sources. [NSConnection] objects automatically add
// their receive port when initialized.
// 
// When the [NSPort] object receives a port message, it forwards the message
// to its delegate in a [HandleMachMessage] or [HandlePortMessage] message.
// The delegate should implement only one of these methods to process the
// incoming message in whatever form desired. [HandleMachMessage] provides a
// message as a raw Mach message beginning with a `msg_header_t` structure.
// [HandlePortMessage] provides a message as an instance of [NSPortMessage],
// which is an object-oriented wrapper for a Mach message. If a delegate has
// not been set, the [NSPort] object handles the message itself.
// 
// When you are finished using a port object, you must explicitly invalidate
// the port object prior to sending it a `release` message. Similarly, if your
// application uses garbage collection, you must invalidate the port object
// before removing any strong references to it. If you do not invalidate the
// port, the resulting port object may linger and create a memory leak. To
// invalidate the port object, invoke its [Invalidate] method.
// 
// Foundation defines three concrete subclasses of [NSPort]. [NSMachPort] and
// [NSMessagePort] allow local (on the same machine) communication only.
// [NSSocketPort] allows for both local and remote communication, but may be
// more expensive than the others for the local case. When creating an
// [NSPort] object, using [allocWithZone:] or [Port], an [NSMachPort] object
// is created instead.
// 
// For backward compatibility on Mach, `-[NSPort ]` returns an instance of the
// [NSMachPort] class when sent to this class. Otherwise, it returns an
// instance of a concrete subclass that can be used for messaging between
// threads or processes on the local machine, or, in the case of
// [NSSocketPort], between processes on separate machines.
//
// [NSConnection]: https://developer.apple.com/documentation/Foundation/NSConnection
// [allocWithZone:]: https://developer.apple.com/documentation/Foundation/nsport-allocwithzone
//
// # Validation
//
//   - [Port.Invalidate]: Marks the receiver as invalid and posts an [didBecomeInvalidNotification](<doc://com.apple.foundation/documentation/Foundation/Port/didBecomeInvalidNotification>) to the default notification center.
//   - [Port.Valid]: A Boolean value that indicates whether the receiver is valid.
//
// # Setting the Delegate
//
//   - [Port.SetDelegate]: Sets the receiver’s delegate to a given object.
//   - [Port.Delegate]: Returns the receiver’s delegate.
//
// # Setting Information
//
//   - [Port.SendBeforeDateComponentsFromReserved]: This method is provided for subclasses that have custom types of [NSPort].
//   - [Port.SendBeforeDateMsgidComponentsFromReserved]: This method is provided for subclasses that have custom types of [NSPort].
//   - [Port.ReservedSpaceLength]: The number of bytes of space reserved by the receiver for sending data.
//
// # Port Monitoring
//
//   - [Port.RemoveFromRunLoopForMode]: This method should be implemented by a subclass to stop monitoring of a port when removed from a give run loop in a given input mode.
//   - [Port.ScheduleInRunLoopForMode]: This method should be implemented by a subclass to set up monitoring of a port when added to a given run loop in a given input mode.
//
// See: https://developer.apple.com/documentation/Foundation/Port
type Port struct {
	objectivec.Object
}

// PortFromID constructs a [Port] from an objc.ID.
//
// An abstract class that represents a communication channel.
func PortFromID(id objc.ID) Port {
	return NSPort{objectivec.Object{id}}
}

// NSPortFromID is an alias for [PortFromID] for cross-framework compatibility.
func NSPortFromID(id objc.ID) Port { return PortFromID(id) }
// NOTE: Port adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [Port] class.
//
// # Validation
//
//   - [IPort.Invalidate]: Marks the receiver as invalid and posts an [didBecomeInvalidNotification](<doc://com.apple.foundation/documentation/Foundation/Port/didBecomeInvalidNotification>) to the default notification center.
//   - [IPort.Valid]: A Boolean value that indicates whether the receiver is valid.
//
// # Setting the Delegate
//
//   - [IPort.SetDelegate]: Sets the receiver’s delegate to a given object.
//   - [IPort.Delegate]: Returns the receiver’s delegate.
//
// # Setting Information
//
//   - [IPort.SendBeforeDateComponentsFromReserved]: This method is provided for subclasses that have custom types of [NSPort].
//   - [IPort.SendBeforeDateMsgidComponentsFromReserved]: This method is provided for subclasses that have custom types of [NSPort].
//   - [IPort.ReservedSpaceLength]: The number of bytes of space reserved by the receiver for sending data.
//
// # Port Monitoring
//
//   - [IPort.RemoveFromRunLoopForMode]: This method should be implemented by a subclass to stop monitoring of a port when removed from a give run loop in a given input mode.
//   - [IPort.ScheduleInRunLoopForMode]: This method should be implemented by a subclass to set up monitoring of a port when added to a given run loop in a given input mode.
//
// See: https://developer.apple.com/documentation/Foundation/Port
type IPort interface {
	objectivec.IObject
	NSCopying

	// Topic: Validation

	// Marks the receiver as invalid and posts an [didBecomeInvalidNotification](<doc://com.apple.foundation/documentation/Foundation/Port/didBecomeInvalidNotification>) to the default notification center.
	Invalidate()
	// A Boolean value that indicates whether the receiver is valid.
	Valid() bool

	// Topic: Setting the Delegate

	// Sets the receiver’s delegate to a given object.
	SetDelegate(anObject NSPortDelegate)
	// Returns the receiver’s delegate.
	Delegate() NSPortDelegate

	// Topic: Setting Information

	// This method is provided for subclasses that have custom types of [NSPort].
	SendBeforeDateComponentsFromReserved(limitDate INSDate, components INSArray, receivePort INSPort, headerSpaceReserved uint) bool
	// This method is provided for subclasses that have custom types of [NSPort].
	SendBeforeDateMsgidComponentsFromReserved(limitDate INSDate, msgID uint, components INSArray, receivePort INSPort, headerSpaceReserved uint) bool
	// The number of bytes of space reserved by the receiver for sending data.
	ReservedSpaceLength() uint

	// Topic: Port Monitoring

	// This method should be implemented by a subclass to stop monitoring of a port when removed from a give run loop in a given input mode.
	RemoveFromRunLoopForMode(runLoop INSRunLoop, mode NSRunLoopMode)
	// This method should be implemented by a subclass to set up monitoring of a port when added to a given run loop in a given input mode.
	ScheduleInRunLoopForMode(runLoop INSRunLoop, mode NSRunLoopMode)

	// Encodes the receiver using a given archiver.
	EncodeWithCoder(coder INSCoder)
	InitWithCoder(coder INSCoder) Port
}





// Init initializes the instance.
func (p Port) Init() Port {
	rv := objc.Send[Port](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p Port) Autorelease() Port {
	rv := objc.Send[Port](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewPort creates a new Port instance.
func NewPort() Port {
	class := getPortClass()
	rv := objc.Send[Port](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewPortWithCoder(coder INSCoder) Port {
	instance := getPortClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return PortFromID(rv)
}







// Marks the receiver as invalid and posts an [didBecomeInvalidNotification]
// to the default notification center.
//
// [didBecomeInvalidNotification]: https://developer.apple.com/documentation/Foundation/Port/didBecomeInvalidNotification
//
// # Discussion
// 
// You must call this method before releasing a port object (or removing
// strong references to it if your application is garbage collected).
//
// See: https://developer.apple.com/documentation/Foundation/Port/invalidate()
func (p Port) Invalidate() {
	objc.Send[objc.ID](p.ID, objc.Sel("invalidate"))
}

// Sets the receiver’s delegate to a given object.
//
// anObject: The delegate for the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/Port/setDelegate(_:)
func (p Port) SetDelegate(anObject NSPortDelegate) {
	objc.Send[objc.ID](p.ID, objc.Sel("setDelegate:"), anObject)
}

// Returns the receiver’s delegate.
//
// # Return Value
// 
// The receiver’s delegate.
//
// See: https://developer.apple.com/documentation/Foundation/Port/delegate()
func (p Port) Delegate() NSPortDelegate {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("delegate"))
	return NSPortDelegateObjectFromID(rv)
}

// This method is provided for subclasses that have custom types of [NSPort].
//
// limitDate: The last instant that a message may be sent.
//
// components: The message components.
//
// receivePort: The receive port.
//
// headerSpaceReserved: The number of bytes reserved for the header.
//
// # Discussion
// 
// [NSConnection] calls this method at the appropriate times. This method
// should not be called directly. This method could raise an
// [NSInvalidSendPortException], [NSInvalidReceivePortException], or an
// [NSPortSendException], depending on the type of send port and the type of
// error.
//
// See: https://developer.apple.com/documentation/Foundation/Port/send(before:components:from:reserved:)
func (p Port) SendBeforeDateComponentsFromReserved(limitDate INSDate, components INSArray, receivePort INSPort, headerSpaceReserved uint) bool {
	rv := objc.Send[bool](p.ID, objc.Sel("sendBeforeDate:components:from:reserved:"), limitDate, components, receivePort, headerSpaceReserved)
	return rv
}

// This method is provided for subclasses that have custom types of [NSPort].
//
// limitDate: The last instant that a message may be sent.
//
// msgID: The message ID.
//
// components: The message components.
//
// receivePort: The receive port.
//
// headerSpaceReserved: The number of bytes reserved for the header.
//
// # Discussion
// 
// [NSConnection] calls this method at the appropriate times. This method
// should not be called directly. This method could raise an
// [NSInvalidSendPortException], [NSInvalidReceivePortException], or an
// [NSPortSendException], depending on the type of send port and the type of
// error.
//
// See: https://developer.apple.com/documentation/Foundation/Port/send(before:msgid:components:from:reserved:)
func (p Port) SendBeforeDateMsgidComponentsFromReserved(limitDate INSDate, msgID uint, components INSArray, receivePort INSPort, headerSpaceReserved uint) bool {
	rv := objc.Send[bool](p.ID, objc.Sel("sendBeforeDate:msgid:components:from:reserved:"), limitDate, msgID, components, receivePort, headerSpaceReserved)
	return rv
}

// This method should be implemented by a subclass to stop monitoring of a
// port when removed from a give run loop in a given input mode.
//
// runLoop: The run loop from which to remove the receiver.
//
// mode: The run loop mode from which to remove the receiver
//
// # Discussion
// 
// This method should not be called directly.
//
// See: https://developer.apple.com/documentation/Foundation/Port/remove(from:forMode:)
func (p Port) RemoveFromRunLoopForMode(runLoop INSRunLoop, mode NSRunLoopMode) {
	objc.Send[objc.ID](p.ID, objc.Sel("removeFromRunLoop:forMode:"), runLoop, objc.String(string(mode)))
}

// This method should be implemented by a subclass to set up monitoring of a
// port when added to a given run loop in a given input mode.
//
// runLoop: The run loop to which to add the receiver.
//
// mode: The run loop mode to which to add the receiver
//
// # Discussion
// 
// This method should not be called directly.
//
// See: https://developer.apple.com/documentation/Foundation/Port/schedule(in:forMode:)
func (p Port) ScheduleInRunLoopForMode(runLoop INSRunLoop, mode NSRunLoopMode) {
	objc.Send[objc.ID](p.ID, objc.Sel("scheduleInRunLoop:forMode:"), runLoop, objc.String(string(mode)))
}

// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (p Port) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](p.ID, objc.Sel("encodeWithCoder:"), coder)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func (p Port) InitWithCoder(coder INSCoder) Port {
	rv := objc.Send[Port](p.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}





// Creates and returns a new [NSPort] object capable of both sending and
// receiving messages.
//
// # Return Value
// 
// A new [NSPort] object capable of both sending and receiving messages.
//
// See: https://developer.apple.com/documentation/Foundation/NSPort/port
func (_PortClass PortClass) Port() Port {
	rv := objc.Send[objc.ID](objc.ID(_PortClass.class), objc.Sel("port"))
	return NSPortFromID(rv)
}








// A Boolean value that indicates whether the receiver is valid.
//
// # Discussion
// 
// [false] if the receiver is known to be invalid, otherwise [true].
// 
// An [NSPort] object becomes invalid when its underlying communication
// resource, which is operating system dependent, is closed or damaged.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/Port/isValid
func (p Port) Valid() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isValid"))
	return rv
}



// The number of bytes of space reserved by the receiver for sending data.
//
// # Discussion
// 
// The number of bytes reserved by the receiver for sending data. The default
// length is `0`.
//
// See: https://developer.apple.com/documentation/Foundation/Port/reservedSpaceLength
func (p Port) ReservedSpaceLength() uint {
	rv := objc.Send[uint](p.ID, objc.Sel("reservedSpaceLength"))
	return rv
}














			// Protocol methods for NSCopying
			














