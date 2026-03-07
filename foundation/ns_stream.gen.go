// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [Stream] class.
var (
	_StreamClass     StreamClass
	_StreamClassOnce sync.Once
)

func getStreamClass() StreamClass {
	_StreamClassOnce.Do(func() {
		_StreamClass = StreamClass{class: objc.GetClass("NSStream")}
	})
	return _StreamClass
}

// GetStreamClass returns the class object for NSStream.
func GetStreamClass() StreamClass {
	return getStreamClass()
}

type StreamClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (sc StreamClass) Alloc() Stream {
	rv := objc.Send[Stream](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}







// An abstract class representing a stream.
//
// # Overview
// 
// This class’s interface is common to all Cocoa stream classes, including
// its concrete subclasses [NSInputStream] and [NSOutputStream].
// 
// [NSStream] objects provide an easy way to read and write data to and from a
// variety of media in a device-independent way. You can create stream objects
// for data located in memory, in a file, or on a network (using sockets), and
// you can use stream objects without loading all of the data into memory at
// once.
// 
// By default, [NSStream] instances that aren’t file-based are non-seekable,
// one-way streams (although custom seekable subclasses are possible). After
// you provide or consume data, you can’t retrieve the data from the stream.
// 
// # Subclassing Notes
// 
// [NSStream] is an abstract class, incapable of instantiation and intended
// for you to subclass it. It publishes a programmatic interface that all
// subclasses must adopt and provide implementations for. The two
// Apple-provided concrete subclasses of [NSStream], [NSInputStream] and
// [NSOutputStream], are suitable for most purposes. However, there might be
// situations when you want a peer subclass to [NSInputStream] and
// [NSOutputStream]. For example, you might want a class that implements a
// full-duplex (two-way) stream, or a class whose instances are capable of
// seeking through a stream.
// 
// # Methods to Override
// 
// All subclasses must fully implement the following methods:
// 
// - [Open] and [Close]
// 
// Implement [Open] to open the stream for reading or writing and make the
// stream available to the client directly or, if the stream object is
// scheduled on a run loop, to the delegate. Implement [Close] to close the
// stream and remove the stream object from the run loop, if necessary. A
// closed stream should still be able to accept new properties and report its
// current properties. Once you close a stream, you can’t reopen it.
// 
// - [Delegate]
// 
// Return and set the delegate. By a default, a stream object must be its own
// delegate; so a [Delegate] message with an argument of `nil` should restore
// this delegate. Don’t retain the delegate to prevent retain cycles.
// 
// To learn about delegates and delegation, read “Delegation” in Cocoa
// Fundamentals Guide.
// 
// - [ScheduleInRunLoopForMode] and [RemoveFromRunLoopForMode]
// 
// Implement [ScheduleInRunLoopForMode] to schedule the stream object on the
// specified run loop for the specified mode. Implement
// [RemoveFromRunLoopForMode] to remove the object from the run loop. See the
// documentation of the [NSRunLoop] class for details. Once the stream object
// for an open stream is scheduled on a run loop, it is the responsibility of
// the subclass as it processes stream data to send [StreamHandleEvent]
// messages to its delegate.
// 
// - [PropertyForKey] and [SetPropertyForKey]
// 
// Implement these methods to return and set, respectively, the property value
// for the specified key. You may add custom properties, but be sure to handle
// all properties defined by [NSStream] as well.
// 
// - [StreamStatus] and [StreamError]
// 
// Implement [StreamStatus] to return the current status of the stream as a
// [Stream.Status] constant; you may define new [Stream.Status] constants, but
// be sure to handle the system defined constants properly. Implement
// [StreamError] to return an [NSError] object representing the current error.
// You might decide to return a custom [NSError] object that can provide
// complete and localized information about the error.
//
// [Stream.Status]: https://developer.apple.com/documentation/Foundation/Stream/Status
//
// # Configuring Streams
//
//   - [Stream.PropertyForKey]: Returns the receiver’s property for a given key.
//   - [Stream.SetPropertyForKey]: Attempts to set the value of a given property of the receiver and returns a Boolean value that indicates whether the value is accepted by the receiver.
//   - [Stream.Delegate]: Sets the receiver’s delegate.
//   - [Stream.SetDelegate]
//
// # Using Streams
//
//   - [Stream.Open]: Opens the receiving stream.
//   - [Stream.Close]: Closes the receiver.
//
// # Managing Run Loops
//
//   - [Stream.ScheduleInRunLoopForMode]: Schedules the receiver on a given run loop in a given mode.
//   - [Stream.RemoveFromRunLoopForMode]: Removes the receiver from a given run loop running in a given mode.
//
// # Getting Stream Information
//
//   - [Stream.StreamStatus]: Returns the receiver’s status.
//   - [Stream.StreamError]: Returns an [NSError] object representing the stream error.
//
// # Constants
//
//   - [Stream.NSStreamSocketSSLErrorDomain]: The error domain used by 
//   - [Stream.NSStreamSOCKSErrorDomain]: The error domain used by 
//
// See: https://developer.apple.com/documentation/Foundation/Stream
type Stream struct {
	objectivec.Object
}

// StreamFromID constructs a [Stream] from an objc.ID.
//
// An abstract class representing a stream.
func StreamFromID(id objc.ID) Stream {
	return NSStream{objectivec.Object{id}}
}

// NSStreamFromID is an alias for [StreamFromID] for cross-framework compatibility.
func NSStreamFromID(id objc.ID) Stream { return StreamFromID(id) }
// NOTE: Stream adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [Stream] class.
//
// # Configuring Streams
//
//   - [IStream.PropertyForKey]: Returns the receiver’s property for a given key.
//   - [IStream.SetPropertyForKey]: Attempts to set the value of a given property of the receiver and returns a Boolean value that indicates whether the value is accepted by the receiver.
//   - [IStream.Delegate]: Sets the receiver’s delegate.
//   - [IStream.SetDelegate]
//
// # Using Streams
//
//   - [IStream.Open]: Opens the receiving stream.
//   - [IStream.Close]: Closes the receiver.
//
// # Managing Run Loops
//
//   - [IStream.ScheduleInRunLoopForMode]: Schedules the receiver on a given run loop in a given mode.
//   - [IStream.RemoveFromRunLoopForMode]: Removes the receiver from a given run loop running in a given mode.
//
// # Getting Stream Information
//
//   - [IStream.StreamStatus]: Returns the receiver’s status.
//   - [IStream.StreamError]: Returns an [NSError] object representing the stream error.
//
// # Constants
//
//   - [IStream.NSStreamSocketSSLErrorDomain]: The error domain used by 
//   - [IStream.NSStreamSOCKSErrorDomain]: The error domain used by 
//
// See: https://developer.apple.com/documentation/Foundation/Stream
type IStream interface {
	objectivec.IObject

	// Topic: Configuring Streams

	// Returns the receiver’s property for a given key.
	PropertyForKey(key NSStreamPropertyKey) objectivec.IObject
	// Attempts to set the value of a given property of the receiver and returns a Boolean value that indicates whether the value is accepted by the receiver.
	SetPropertyForKey(property objectivec.IObject, key NSStreamPropertyKey) bool
	// Sets the receiver’s delegate.
	Delegate() NSStreamDelegate
	SetDelegate(value NSStreamDelegate)

	// Topic: Using Streams

	// Opens the receiving stream.
	Open()
	// Closes the receiver.
	Close()

	// Topic: Managing Run Loops

	// Schedules the receiver on a given run loop in a given mode.
	ScheduleInRunLoopForMode(aRunLoop INSRunLoop, mode NSRunLoopMode)
	// Removes the receiver from a given run loop running in a given mode.
	RemoveFromRunLoopForMode(aRunLoop INSRunLoop, mode NSRunLoopMode)

	// Topic: Getting Stream Information

	// Returns the receiver’s status.
	StreamStatus() NSStreamStatus
	// Returns an [NSError] object representing the stream error.
	StreamError() INSError

	// Topic: Constants

	// The error domain used by 
	NSStreamSocketSSLErrorDomain() string
	// The error domain used by 
	NSStreamSOCKSErrorDomain() string
}





// Init initializes the instance.
func (s Stream) Init() Stream {
	rv := objc.Send[Stream](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s Stream) Autorelease() Stream {
	rv := objc.Send[Stream](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewStream creates a new Stream instance.
func NewStream() Stream {
	class := getStreamClass()
	rv := objc.Send[Stream](objc.ID(class.class), objc.Sel("new"))
	return rv
}










// Returns the receiver’s property for a given key.
//
// key: The key for one of the receiver’s properties. See Constants for a
// description of the available property-key constants and associated values.
//
// # Return Value
// 
// The receiver’s property for the key `key`.
//
// See: https://developer.apple.com/documentation/Foundation/Stream/property(forKey:)
func (s Stream) PropertyForKey(key NSStreamPropertyKey) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("propertyForKey:"), objc.String(string(key)))
	return objectivec.Object{ID: rv}
}

// Attempts to set the value of a given property of the receiver and returns a
// Boolean value that indicates whether the value is accepted by the receiver.
//
// property: The value for `key`.
//
// key: The key for one of the receiver’s properties. See Constants for a
// description of the available property-key constants and expected values.
//
// # Return Value
// 
// [true] if the value is accepted by the receiver, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/Stream/setProperty(_:forKey:)
func (s Stream) SetPropertyForKey(property objectivec.IObject, key NSStreamPropertyKey) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("setProperty:forKey:"), property, objc.String(string(key)))
	return rv
}

// Opens the receiving stream.
//
// # Discussion
// 
// A stream must be created before it can be opened. Once opened, a stream
// cannot be closed and reopened.
//
// See: https://developer.apple.com/documentation/Foundation/Stream/open()
func (s Stream) Open() {
	objc.Send[objc.ID](s.ID, objc.Sel("open"))
}

// Closes the receiver.
//
// # Discussion
// 
// Closing the stream terminates the flow of bytes and releases system
// resources that were reserved for the stream when it was opened. If the
// stream has been scheduled on a run loop, closing the stream implicitly
// removes the stream from the run loop. A stream that is closed can still be
// queried for its properties.
//
// See: https://developer.apple.com/documentation/Foundation/Stream/close()
func (s Stream) Close() {
	objc.Send[objc.ID](s.ID, objc.Sel("close"))
}

// Schedules the receiver on a given run loop in a given mode.
//
// aRunLoop: The run loop on which to schedule the receiver.
//
// mode: The mode for the run loop.
//
// # Discussion
// 
// Unless the client is polling the stream, it is responsible for ensuring
// that the stream is scheduled on at least one run loop and that at least one
// of the run loops on which the stream is scheduled is being run.
//
// See: https://developer.apple.com/documentation/Foundation/Stream/schedule(in:forMode:)
func (s Stream) ScheduleInRunLoopForMode(aRunLoop INSRunLoop, mode NSRunLoopMode) {
	objc.Send[objc.ID](s.ID, objc.Sel("scheduleInRunLoop:forMode:"), aRunLoop, objc.String(string(mode)))
}

// Removes the receiver from a given run loop running in a given mode.
//
// aRunLoop: The run loop on which the receiver was scheduled.
//
// mode: The mode for the run loop.
//
// See: https://developer.apple.com/documentation/Foundation/Stream/remove(from:forMode:)
func (s Stream) RemoveFromRunLoopForMode(aRunLoop INSRunLoop, mode NSRunLoopMode) {
	objc.Send[objc.ID](s.ID, objc.Sel("removeFromRunLoop:forMode:"), aRunLoop, objc.String(string(mode)))
}





// Creates and returns by reference a bound pair of input and output streams.
//
// bufferSize: The size of the buffer, in bytes, used to transfer data from `inputStream`
// to `outputStream`.
//
// inputStream: On return, contains an input stream.
//
// outputStream: On return, contains an output stream.
//
// # Discussion
// 
// The created streams are bound to one another, such that any data written to
// `outputStream` is received by `inputStream`.
// 
// This is a convenience method for calling
// [CFStreamCreateBoundPair(_:_:_:_:)] and bridging from the returned Core
// Foundation types.
//
// [CFStreamCreateBoundPair(_:_:_:_:)]: https://developer.apple.com/documentation/CoreFoundation/CFStreamCreateBoundPair(_:_:_:_:)
//
// See: https://developer.apple.com/documentation/Foundation/Stream/getBoundStreams(withBufferSize:inputStream:outputStream:)
func (_StreamClass StreamClass) GetBoundStreamsWithBufferSizeInputStreamOutputStream(bufferSize uint, inputStream INSInputStream, outputStream INSOutputStream) {
	objc.Send[objc.ID](objc.ID(_StreamClass.class), objc.Sel("getBoundStreamsWithBufferSize:inputStream:outputStream:"), bufferSize, inputStream, outputStream)
}








// Sets the receiver’s delegate.
//
// # Discussion
// 
// By default, a stream is its own delegate, and subclasses of [NSInputStream]
// and [NSOutputStream] must maintain this contract. If you override this
// method in a subclass, passing `nil` must restore the receiver as its own
// delegate. Delegates are not retained.
// 
// To learn about delegates and delegation, read “Delegation” in Cocoa
// Fundamentals Guide.
//
// See: https://developer.apple.com/documentation/Foundation/Stream/delegate
func (s Stream) Delegate() NSStreamDelegate {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("delegate"))
	return NSStreamDelegateObjectFromID(rv)
}
func (s Stream) SetDelegate(value NSStreamDelegate) {
	objc.Send[struct{}](s.ID, objc.Sel("setDelegate:"), value)
}



// Returns the receiver’s status.
//
// # Return Value
// 
// The receiver’s status.
// 
// # Discussion
// 
// See Constants for a description of the available NSStreamStatus constants.
//
// See: https://developer.apple.com/documentation/Foundation/Stream/streamStatus
func (s Stream) StreamStatus() NSStreamStatus {
	rv := objc.Send[NSStreamStatus](s.ID, objc.Sel("streamStatus"))
	return NSStreamStatus(rv)
}



// Returns an [NSError] object representing the stream error.
//
// # Return Value
// 
// An [NSError] object representing the stream error, or `nil` if no error has
// been encountered.
//
// See: https://developer.apple.com/documentation/Foundation/Stream/streamError
func (s Stream) StreamError() INSError {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("streamError"))
	return NSErrorFromID(objc.ID(rv))
}



// The error domain used by
//
// See: https://developer.apple.com/documentation/foundation/nsstreamsocketsslerrordomain
func (s Stream) NSStreamSocketSSLErrorDomain() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("NSStreamSocketSSLErrorDomain"))
	return NSStringFromID(rv).String()
}



// The error domain used by
//
// See: https://developer.apple.com/documentation/foundation/nsstreamsockserrordomain
func (s Stream) NSStreamSOCKSErrorDomain() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("NSStreamSOCKSErrorDomain"))
	return NSStringFromID(rv).String()
}























