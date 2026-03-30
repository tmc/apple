// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
)

// The class instance for the [OutputStream] class.
var (
	_OutputStreamClass     OutputStreamClass
	_OutputStreamClassOnce sync.Once
)

func getOutputStreamClass() OutputStreamClass {
	_OutputStreamClassOnce.Do(func() {
		_OutputStreamClass = OutputStreamClass{class: objc.GetClass("NSOutputStream")}
	})
	return _OutputStreamClass
}

// GetOutputStreamClass returns the class object for NSOutputStream.
func GetOutputStreamClass() OutputStreamClass {
	return getOutputStreamClass()
}

type OutputStreamClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (oc OutputStreamClass) Class() objc.Class {
	return oc.class
}

// Alloc allocates memory for a new instance of the class.
func (oc OutputStreamClass) Alloc() OutputStream {
	rv := objc.Send[OutputStream](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// A stream that provides write-only stream functionality.
//
// # Overview
//
// [NSOutputStream] is “toll-free bridged” with its Core Foundation
// counterpart, [CFWriteStream]. For more information on toll-free bridging,
// see [Toll-Free Bridging].
//
// # Subclassing Notes
//
// [NSOutputStream] is a concrete subclass of [NSStream] that lets you write
// data to a stream. Although [NSOutputStream] is probably sufficient for most
// situations requiring this capability, you can create a subclass of
// [NSOutputStream] if you want more specialized behavior (for example, you
// want to record statistics on the data in a stream).
//
// # Methods to Override
//
// To create a subclass of [NSOutputStream] you may have to implement
// initializers for the type of stream data supported and suitably reimplement
// existing initializers. You must also provide complete implementations of
// the following methods:
//
// - [WriteMaxLength]
//
// From the current write pointer, take up to the number of bytes specified in
// the “ parameter from the client-supplied buffer (first parameter) and put
// them onto the stream. The buffer must be of the size specified by the
// second parameter. To prepare for the next operation, offset the write
// pointer by the number of bytes written. Return a signed integer based on
// the outcome of the current operation:
//
// - If the write operation is successful, return the actual number of bytes
// put onto the stream. - If the stream is of a fixed length and has reached
// its capacity, return `0`. - If there was an error writing to the stream,
// return `-1`. - [HasSpaceAvailable]
//
// Return true if the stream can currently accept more data, false if it
// cannot. If you want to be semantically compatible with [NSOutputStream],
// return true if a write must be attempted to determine if space is
// available.
//
// # Creating Streams
//
//   - [OutputStream.InitToMemory]: Returns an initialized output stream that will write to memory.
//   - [OutputStream.InitToBufferCapacity]: Returns an initialized output stream that can write to a provided buffer.
//   - [OutputStream.InitToFileAtPathAppend]: Returns an initialized output stream for writing to a specified file.
//   - [OutputStream.InitWithURLAppend]: Returns an initialized output stream for writing to a specified URL.
//
// # Using Streams
//
//   - [OutputStream.HasSpaceAvailable]: A boolean value that indicates whether the receiver can be written to.
//   - [OutputStream.WriteMaxLength]: Writes the contents of a provided data buffer to the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/OutputStream
//
// [CFWriteStream]: https://developer.apple.com/documentation/CoreFoundation/CFWriteStream
// [Toll-Free Bridging]: https://developer.apple.com/library/archive/documentation/General/Conceptual/CocoaEncyclopedia/Toll-FreeBridgin/Toll-FreeBridgin.html#//apple_ref/doc/uid/TP40010810-CH2
type OutputStream struct {
	NSStream
}

// OutputStreamFromID constructs a [OutputStream] from an objc.ID.
//
// A stream that provides write-only stream functionality.
func OutputStreamFromID(id objc.ID) OutputStream {
	return NSOutputStream{NSStream: NSStreamFromID(id)}
}

// NSOutputStreamFromID is an alias for [OutputStreamFromID] for cross-framework compatibility.
func NSOutputStreamFromID(id objc.ID) OutputStream { return OutputStreamFromID(id) }

// NOTE: OutputStream adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [OutputStream] class.
//
// # Creating Streams
//
//   - [IOutputStream.InitToMemory]: Returns an initialized output stream that will write to memory.
//   - [IOutputStream.InitToBufferCapacity]: Returns an initialized output stream that can write to a provided buffer.
//   - [IOutputStream.InitToFileAtPathAppend]: Returns an initialized output stream for writing to a specified file.
//   - [IOutputStream.InitWithURLAppend]: Returns an initialized output stream for writing to a specified URL.
//
// # Using Streams
//
//   - [IOutputStream.HasSpaceAvailable]: A boolean value that indicates whether the receiver can be written to.
//   - [IOutputStream.WriteMaxLength]: Writes the contents of a provided data buffer to the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/OutputStream
type IOutputStream interface {
	INSStream

	// Topic: Creating Streams

	// Returns an initialized output stream that will write to memory.
	InitToMemory() OutputStream
	// Returns an initialized output stream that can write to a provided buffer.
	InitToBufferCapacity(buffer unsafe.Pointer, capacity uint) OutputStream
	// Returns an initialized output stream for writing to a specified file.
	InitToFileAtPathAppend(path string, shouldAppend bool) OutputStream
	// Returns an initialized output stream for writing to a specified URL.
	InitWithURLAppend(url INSURL, shouldAppend bool) OutputStream

	// Topic: Using Streams

	// A boolean value that indicates whether the receiver can be written to.
	HasSpaceAvailable() bool
	// Writes the contents of a provided data buffer to the receiver.
	WriteMaxLength(buffer unsafe.Pointer, len_ uint) int
}

// Init initializes the instance.
func (o OutputStream) Init() OutputStream {
	rv := objc.Send[OutputStream](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o OutputStream) Autorelease() OutputStream {
	rv := objc.Send[OutputStream](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewOutputStream creates a new OutputStream instance.
func NewOutputStream() OutputStream {
	class := getOutputStreamClass()
	rv := objc.Send[OutputStream](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns an initialized output stream that can write to a provided buffer.
//
// buffer: The buffer the output stream will write to.
//
// capacity: The size of the buffer in bytes.
//
// # Return Value
//
// An initialized output stream that can write to `buffer`.
//
// # Discussion
//
// The stream must be opened before it can be used.
//
// When the number of bytes written to `buffer` has reached `capacity`, the
// stream’s [StreamStatus] will return [NSStreamStatusAtEnd].
//
// See: https://developer.apple.com/documentation/Foundation/OutputStream/init(toBuffer:capacity:)
func NewOutputStreamToBufferCapacity(buffer unsafe.Pointer, capacity uint) OutputStream {
	instance := getOutputStreamClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initToBuffer:capacity:"), buffer, capacity)
	return OutputStreamFromID(rv)
}

// Returns an initialized output stream for writing to a specified file.
//
// path: The path to the file the output stream will write to.
//
// shouldAppend: true if newly written data should be appended to any existing file
// contents, otherwise false.
//
// # Return Value
//
// An initialized output stream that can write to `path`.
//
// # Discussion
//
// The stream must be opened before it can be used.
//
// See: https://developer.apple.com/documentation/Foundation/OutputStream/init(toFileAtPath:append:)
func NewOutputStreamToFileAtPathAppend(path string, shouldAppend bool) OutputStream {
	instance := getOutputStreamClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initToFileAtPath:append:"), objc.String(path), shouldAppend)
	return OutputStreamFromID(rv)
}

// Returns an initialized output stream that will write to memory.
//
// # Return Value
//
// An initialized output stream that will write stream data to memory.
//
// # Discussion
//
// The stream must be opened before it can be used.
//
// The contents of the memory stream are retrieved by passing the constant
// [NSStreamDataWrittenToMemoryStreamKey] to [PropertyForKey].
//
// See: https://developer.apple.com/documentation/Foundation/OutputStream/init(toMemory:)
func NewOutputStreamToMemory() OutputStream {
	instance := getOutputStreamClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initToMemory"))
	return OutputStreamFromID(rv)
}

// Returns an initialized output stream for writing to a specified URL.
//
// url: The URL to the file the output stream will write to.
//
// shouldAppend: true if newly written data should be appended to any existing file
// contents, otherwise false.
//
// # Return Value
//
// An initialized output stream that can write to `url`.
//
// # Discussion
//
// The stream must be opened before it can be used.
//
// See: https://developer.apple.com/documentation/Foundation/OutputStream/init(url:append:)
func NewOutputStreamWithURLAppend(url INSURL, shouldAppend bool) OutputStream {
	instance := getOutputStreamClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:append:"), url, shouldAppend)
	return OutputStreamFromID(rv)
}

// Returns an initialized output stream that will write to memory.
//
// # Return Value
//
// An initialized output stream that will write stream data to memory.
//
// # Discussion
//
// The stream must be opened before it can be used.
//
// The contents of the memory stream are retrieved by passing the constant
// [NSStreamDataWrittenToMemoryStreamKey] to [PropertyForKey].
//
// See: https://developer.apple.com/documentation/Foundation/OutputStream/init(toMemory:)
func (o OutputStream) InitToMemory() OutputStream {
	rv := objc.Send[OutputStream](o.ID, objc.Sel("initToMemory"))
	return rv
}

// Returns an initialized output stream that can write to a provided buffer.
//
// buffer: The buffer the output stream will write to.
//
// capacity: The size of the buffer in bytes.
//
// # Return Value
//
// An initialized output stream that can write to `buffer`.
//
// # Discussion
//
// The stream must be opened before it can be used.
//
// When the number of bytes written to `buffer` has reached `capacity`, the
// stream’s [StreamStatus] will return [NSStreamStatusAtEnd].
//
// See: https://developer.apple.com/documentation/Foundation/OutputStream/init(toBuffer:capacity:)
func (o OutputStream) InitToBufferCapacity(buffer unsafe.Pointer, capacity uint) OutputStream {
	rv := objc.Send[OutputStream](o.ID, objc.Sel("initToBuffer:capacity:"), buffer, capacity)
	return rv
}

// Returns an initialized output stream for writing to a specified file.
//
// path: The path to the file the output stream will write to.
//
// shouldAppend: true if newly written data should be appended to any existing file
// contents, otherwise false.
//
// # Return Value
//
// An initialized output stream that can write to `path`.
//
// # Discussion
//
// The stream must be opened before it can be used.
//
// See: https://developer.apple.com/documentation/Foundation/OutputStream/init(toFileAtPath:append:)
func (o OutputStream) InitToFileAtPathAppend(path string, shouldAppend bool) OutputStream {
	rv := objc.Send[OutputStream](o.ID, objc.Sel("initToFileAtPath:append:"), objc.String(path), shouldAppend)
	return rv
}

// Returns an initialized output stream for writing to a specified URL.
//
// url: The URL to the file the output stream will write to.
//
// shouldAppend: true if newly written data should be appended to any existing file
// contents, otherwise false.
//
// # Return Value
//
// An initialized output stream that can write to `url`.
//
// # Discussion
//
// The stream must be opened before it can be used.
//
// See: https://developer.apple.com/documentation/Foundation/OutputStream/init(url:append:)
func (o OutputStream) InitWithURLAppend(url INSURL, shouldAppend bool) OutputStream {
	rv := objc.Send[OutputStream](o.ID, objc.Sel("initWithURL:append:"), url, shouldAppend)
	return rv
}

// Writes the contents of a provided data buffer to the receiver.
//
// buffer: The data to write.
//
// len: The length of the data buffer, in bytes.
//
// # Return Value
//
// A number indicating the outcome of the operation:
//
// - A positive number indicates the number of bytes written. - `0` indicates
// that a fixed-length stream and has reached its capacity. - `-1` means that
// the operation failed; more information about the error can be obtained with
// [StreamError].
//
// See: https://developer.apple.com/documentation/Foundation/OutputStream/write(_:maxLength:)
func (o OutputStream) WriteMaxLength(buffer unsafe.Pointer, len_ uint) int {
	rv := objc.Send[int](o.ID, objc.Sel("write:maxLength:"), buffer, len_)
	return rv
}

// Creates and returns an initialized output stream that will write stream
// data to memory.
//
// # Return Value
//
// An initialized output stream that will write stream data to memory.
//
// # Discussion
//
// The stream must be opened before it can be used.
//
// You retrieve the contents of the memory stream by sending the message
// [PropertyForKey] to the receiver with an argument of
// [NSStreamDataWrittenToMemoryStreamKey].
//
// See: https://developer.apple.com/documentation/Foundation/OutputStream/toMemory()
func (_OutputStreamClass OutputStreamClass) OutputStreamToMemory() OutputStream {
	rv := objc.Send[objc.ID](objc.ID(_OutputStreamClass.class), objc.Sel("outputStreamToMemory"))
	return NSOutputStreamFromID(rv)
}

// Creates and returns an initialized output stream that can write to a
// provided buffer.
//
// buffer: The buffer the output stream will write to.
//
// capacity: The size of the buffer in bytes.
//
// # Return Value
//
// An initialized output stream that can write to `buffer`.
//
// # Discussion
//
// The stream must be opened before it can be used.
//
// When the number of bytes written to `buffer` has reached `capacity`, the
// stream’s [StreamStatus] will return [NSStreamStatusAtEnd].
//
// See: https://developer.apple.com/documentation/Foundation/NSOutputStream/outputStreamToBuffer:capacity:
func (_OutputStreamClass OutputStreamClass) OutputStreamToBufferCapacity(buffer unsafe.Pointer, capacity uint) OutputStream {
	rv := objc.Send[objc.ID](objc.ID(_OutputStreamClass.class), objc.Sel("outputStreamToBuffer:capacity:"), buffer, capacity)
	return NSOutputStreamFromID(rv)
}

// Creates and returns an initialized output stream for writing to a specified
// file.
//
// path: The path to the file the output stream will write to.
//
// shouldAppend: true if newly written data should be appended to any existing file
// contents, otherwise false.
//
// # Return Value
//
// An initialized output stream that can write to `path`.
//
// # Discussion
//
// The stream must be opened before it can be used.
//
// See: https://developer.apple.com/documentation/Foundation/NSOutputStream/outputStreamToFileAtPath:append:
func (_OutputStreamClass OutputStreamClass) OutputStreamToFileAtPathAppend(path string, shouldAppend bool) OutputStream {
	rv := objc.Send[objc.ID](objc.ID(_OutputStreamClass.class), objc.Sel("outputStreamToFileAtPath:append:"), objc.String(path), shouldAppend)
	return NSOutputStreamFromID(rv)
}

// See: https://developer.apple.com/documentation/Foundation/NSOutputStream/outputStreamWithURL:append:
func (_OutputStreamClass OutputStreamClass) OutputStreamWithURLAppend(url INSURL, shouldAppend bool) OutputStream {
	rv := objc.Send[objc.ID](objc.ID(_OutputStreamClass.class), objc.Sel("outputStreamWithURL:append:"), url, shouldAppend)
	return NSOutputStreamFromID(rv)
}

// A boolean value that indicates whether the receiver can be written to.
//
// # Discussion
//
// true if the receiver can be written to or if a write must be attempted in
// order to determine if space is available, false otherwise.
//
// See: https://developer.apple.com/documentation/Foundation/OutputStream/hasSpaceAvailable
func (o OutputStream) HasSpaceAvailable() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("hasSpaceAvailable"))
	return rv
}
