// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [InputStream] class.
var (
	_InputStreamClass     InputStreamClass
	_InputStreamClassOnce sync.Once
)

func getInputStreamClass() InputStreamClass {
	_InputStreamClassOnce.Do(func() {
		_InputStreamClass = InputStreamClass{class: objc.GetClass("NSInputStream")}
	})
	return _InputStreamClass
}

// GetInputStreamClass returns the class object for NSInputStream.
func GetInputStreamClass() InputStreamClass {
	return getInputStreamClass()
}

type InputStreamClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ic InputStreamClass) Alloc() InputStream {
	rv := objc.Send[InputStream](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}







// A stream that provides read-only stream functionality.
//
// # Overview
// 
// [NSInputStream] is “toll-free bridged” with its Core Foundation
// counterpart, [CFReadStream]. For more information on toll-free bridging,
// see [Toll-Free Bridging].
// 
// # Subclassing Notes
// 
// [NSInputStream] is an abstract superclass of a consisting of concrete
// subclasses of [NSStream] that provide standard read-only access to stream
// data. Although [NSInputStream] is probably sufficient for most situations
// requiring access to stream data, you can create a subclass of
// [NSInputStream] if you want more specialized behavior (for example, you
// want to record statistics on the data in a stream).
// 
// # Methods to Override
// 
// To create a subclass of [NSInputStream] you may have to implement
// initializers for the type of stream data supported and suitably
// re-implement existing initializers. You must also provide complete
// implementations of the following methods:
// 
// - [ReadMaxLength]
// 
// From the current read index, take up to the number of bytes specified in
// the second parameter from the stream and place them in the client-supplied
// buffer (first parameter). The buffer must be of the size specified by the
// second parameter. Return the actual number of bytes placed in the buffer;
// if there is nothing left in the stream, return `0`. Reset the index into
// the stream for the next read operation.
// 
// - [GetBufferLength]
// 
// Return in 0(1) a pointer to the subclass-allocated buffer (first
// parameter). Return by reference in the second parameter the number of bytes
// actually put into the buffer. The buffer’s contents are valid only until
// the next stream operation. Return [false] if you cannot access data in the
// buffer; otherwise, return [true]. If this method is not appropriate for
// your type of stream, you may return [false].
// 
// - [HasBytesAvailable]
// 
// Return [true] if there is more data to read in the stream, [false] if there
// is not. If you want to be semantically compatible with [NSInputStream],
// return [true] if a read must be attempted to determine if bytes are
// available.
//
// [CFReadStream]: https://developer.apple.com/documentation/CoreFoundation/CFReadStream
// [Toll-Free Bridging]: https://developer.apple.com/library/archive/documentation/General/Conceptual/CocoaEncyclopedia/Toll-FreeBridgin/Toll-FreeBridgin.html#//apple_ref/doc/uid/TP40010810-CH2
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Creating Streams
//
//   - [InputStream.InitWithData]: Initializes and returns an [NSInputStream] object for reading from a given [NSData] object.
//   - [InputStream.InitWithFileAtPath]: Initializes and returns an [NSInputStream] object that reads data from the file at a given path.
//   - [InputStream.InitWithURL]: Initializes and returns an [NSInputStream] object that reads data from the file at a given URL.
//
// # Using Streams
//
//   - [InputStream.ReadMaxLength]: Reads up to a given number of bytes into a given buffer.
//   - [InputStream.GetBufferLength]: Returns by reference a pointer to a read buffer and, by reference, the number of bytes available, and returns a Boolean value that indicates whether the buffer is available.
//   - [InputStream.HasBytesAvailable]: A Boolean value that indicates whether the receiver has bytes available to read.
//
// See: https://developer.apple.com/documentation/Foundation/InputStream
type InputStream struct {
	NSStream
}

// InputStreamFromID constructs a [InputStream] from an objc.ID.
//
// A stream that provides read-only stream functionality.
func InputStreamFromID(id objc.ID) InputStream {
	return NSInputStream{NSStream: NSStreamFromID(id)}
}

// NSInputStreamFromID is an alias for [InputStreamFromID] for cross-framework compatibility.
func NSInputStreamFromID(id objc.ID) InputStream { return InputStreamFromID(id) }
// NOTE: InputStream adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [InputStream] class.
//
// # Creating Streams
//
//   - [IInputStream.InitWithData]: Initializes and returns an [NSInputStream] object for reading from a given [NSData] object.
//   - [IInputStream.InitWithFileAtPath]: Initializes and returns an [NSInputStream] object that reads data from the file at a given path.
//   - [IInputStream.InitWithURL]: Initializes and returns an [NSInputStream] object that reads data from the file at a given URL.
//
// # Using Streams
//
//   - [IInputStream.ReadMaxLength]: Reads up to a given number of bytes into a given buffer.
//   - [IInputStream.GetBufferLength]: Returns by reference a pointer to a read buffer and, by reference, the number of bytes available, and returns a Boolean value that indicates whether the buffer is available.
//   - [IInputStream.HasBytesAvailable]: A Boolean value that indicates whether the receiver has bytes available to read.
//
// See: https://developer.apple.com/documentation/Foundation/InputStream
type IInputStream interface {
	INSStream

	// Topic: Creating Streams

	// Initializes and returns an [NSInputStream] object for reading from a given [NSData] object.
	InitWithData(data INSData) InputStream
	// Initializes and returns an [NSInputStream] object that reads data from the file at a given path.
	InitWithFileAtPath(path string) InputStream
	// Initializes and returns an [NSInputStream] object that reads data from the file at a given URL.
	InitWithURL(url INSURL) InputStream

	// Topic: Using Streams

	// Reads up to a given number of bytes into a given buffer.
	ReadMaxLength(buffer uint8, len_ uint) int
	// Returns by reference a pointer to a read buffer and, by reference, the number of bytes available, and returns a Boolean value that indicates whether the buffer is available.
	GetBufferLength(buffer uint8, len_ uint) bool
	// A Boolean value that indicates whether the receiver has bytes available to read.
	HasBytesAvailable() bool
}





// Init initializes the instance.
func (i InputStream) Init() InputStream {
	rv := objc.Send[InputStream](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i InputStream) Autorelease() InputStream {
	rv := objc.Send[InputStream](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewInputStream creates a new InputStream instance.
func NewInputStream() InputStream {
	class := getInputStreamClass()
	rv := objc.Send[InputStream](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Initializes and returns an [NSInputStream] object for reading from a given
// [NSData] object.
//
// data: The data object from which to read. The contents of `data` are copied.
//
// # Return Value
// 
// An initialized [NSInputStream] object for reading from `data`.
//
// # Discussion
// 
// The stream must be opened before it can be used.
//
// See: https://developer.apple.com/documentation/Foundation/InputStream/init(data:)
func NewInputStreamWithData(data INSData) InputStream {
	instance := getInputStreamClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:"), data)
	return InputStreamFromID(rv)
}


// Initializes and returns an [NSInputStream] object that reads data from the
// file at a given path.
//
// path: The path to the file.
//
// # Return Value
// 
// An initialized [NSInputStream] object that reads data from the file at
// `path`.
//
// # Discussion
// 
// The stream must be opened before it can be used.
//
// See: https://developer.apple.com/documentation/Foundation/InputStream/init(fileAtPath:)
func NewInputStreamWithFileAtPath(path string) InputStream {
	instance := getInputStreamClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFileAtPath:"), objc.String(path))
	return InputStreamFromID(rv)
}


// Initializes and returns an [NSInputStream] object that reads data from the
// file at a given URL.
//
// url: The URL to the file.
//
// # Return Value
// 
// An initialized [NSInputStream] object that reads data from the file at
// `url`.
//
// # Discussion
// 
// The stream must be opened before it can be used.
//
// See: https://developer.apple.com/documentation/Foundation/InputStream/init(url:)
func NewInputStreamWithURL(url INSURL) InputStream {
	instance := getInputStreamClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:"), url)
	return InputStreamFromID(rv)
}







// Initializes and returns an [NSInputStream] object for reading from a given
// [NSData] object.
//
// data: The data object from which to read. The contents of `data` are copied.
//
// # Return Value
// 
// An initialized [NSInputStream] object for reading from `data`.
//
// # Discussion
// 
// The stream must be opened before it can be used.
//
// See: https://developer.apple.com/documentation/Foundation/InputStream/init(data:)
func (i InputStream) InitWithData(data INSData) InputStream {
	rv := objc.Send[InputStream](i.ID, objc.Sel("initWithData:"), data)
	return rv
}

// Initializes and returns an [NSInputStream] object that reads data from the
// file at a given path.
//
// path: The path to the file.
//
// # Return Value
// 
// An initialized [NSInputStream] object that reads data from the file at
// `path`.
//
// # Discussion
// 
// The stream must be opened before it can be used.
//
// See: https://developer.apple.com/documentation/Foundation/InputStream/init(fileAtPath:)
func (i InputStream) InitWithFileAtPath(path string) InputStream {
	rv := objc.Send[InputStream](i.ID, objc.Sel("initWithFileAtPath:"), objc.String(path))
	return rv
}

// Initializes and returns an [NSInputStream] object that reads data from the
// file at a given URL.
//
// url: The URL to the file.
//
// # Return Value
// 
// An initialized [NSInputStream] object that reads data from the file at
// `url`.
//
// # Discussion
// 
// The stream must be opened before it can be used.
//
// See: https://developer.apple.com/documentation/Foundation/InputStream/init(url:)
func (i InputStream) InitWithURL(url INSURL) InputStream {
	rv := objc.Send[InputStream](i.ID, objc.Sel("initWithURL:"), url)
	return rv
}

// Reads up to a given number of bytes into a given buffer.
//
// buffer: A data buffer. The buffer must be large enough to contain the number of
// bytes specified by `len`.
//
// len: The maximum number of bytes to read.
//
// # Return Value
// 
// A number indicating the outcome of the operation:
//
// # Discussion
// 
// - A positive number indicates the number of bytes read. - `0` indicates
// that the end of the buffer was reached. - `-1` means that the operation
// failed; more information about the error can be obtained with
// [StreamError].
//
// See: https://developer.apple.com/documentation/Foundation/InputStream/read(_:maxLength:)
func (i InputStream) ReadMaxLength(buffer uint8, len_ uint) int {
	rv := objc.Send[int](i.ID, objc.Sel("read:maxLength:"), buffer, len_)
	return rv
}

// Returns by reference a pointer to a read buffer and, by reference, the
// number of bytes available, and returns a Boolean value that indicates
// whether the buffer is available.
//
// buffer: Upon return, contains a pointer to a read buffer. The buffer is only valid
// until the next stream operation is performed.
//
// len: Upon return, contains the number of bytes available.
//
// # Return Value
// 
// [true] if the buffer is available, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Subclasses of [NSInputStream] may return [false] if this operation is not
// appropriate for the stream type.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/Foundation/InputStream/getBuffer(_:length:)
func (i InputStream) GetBufferLength(buffer uint8, len_ uint) bool {
	rv := objc.Send[bool](i.ID, objc.Sel("getBuffer:length:"), buffer, len_)
	return rv
}





//
// See: https://developer.apple.com/documentation/Foundation/NSInputStream/inputStreamWithURL:
func (_InputStreamClass InputStreamClass) InputStreamWithURL(url INSURL) InputStream {
	rv := objc.Send[objc.ID](objc.ID(_InputStreamClass.class), objc.Sel("inputStreamWithURL:"), url)
	return NSInputStreamFromID(rv)
}

// Creates and returns an initialized [NSInputStream] object for reading from
// a given [NSData] object.
//
// data: The data object from which to read. The contents of `data` are copied.
//
// # Return Value
// 
// An initialized [NSInputStream] object for reading from `data`. If `data` is
// not an NSData object, this method returns `nil`.
//
// # Discussion
// 
// The stream must be opened before it can be used.
//
// See: https://developer.apple.com/documentation/Foundation/NSInputStream/inputStreamWithData:
func (_InputStreamClass InputStreamClass) InputStreamWithData(data INSData) InputStream {
	rv := objc.Send[objc.ID](objc.ID(_InputStreamClass.class), objc.Sel("inputStreamWithData:"), data)
	return NSInputStreamFromID(rv)
}

// Creates and returns an initialized [NSInputStream] object that reads data
// from the file at a given path.
//
// path: The path to the file.
//
// # Return Value
// 
// An initialized [NSInputStream] object that reads data from the file at
// `path`.
//
// # Discussion
// 
// The stream must be opened before it can be used.
//
// See: https://developer.apple.com/documentation/Foundation/NSInputStream/inputStreamWithFileAtPath:
func (_InputStreamClass InputStreamClass) InputStreamWithFileAtPath(path string) InputStream {
	rv := objc.Send[objc.ID](objc.ID(_InputStreamClass.class), objc.Sel("inputStreamWithFileAtPath:"), objc.String(path))
	return NSInputStreamFromID(rv)
}








// A Boolean value that indicates whether the receiver has bytes available to
// read.
//
// # Discussion
// 
// [true] if the receiver has bytes available to read, otherwise [false]. May
// also return [true] if a read must be attempted in order to determine the
// availability of bytes.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/InputStream/hasBytesAvailable
func (i InputStream) HasBytesAvailable() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("hasBytesAvailable"))
	return rv
}

























