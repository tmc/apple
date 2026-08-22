// Code generated from Apple documentation for IOUSBHost. DO NOT EDIT.

package iousbhost

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [IOUSBHostStream] class.
var (
	_IOUSBHostStreamClass     IOUSBHostStreamClass
	_IOUSBHostStreamClassOnce sync.Once
)

func getIOUSBHostStreamClass() IOUSBHostStreamClass {
	_IOUSBHostStreamClassOnce.Do(func() {
		_IOUSBHostStreamClass = IOUSBHostStreamClass{class: objc.GetClass("IOUSBHostStream")}
	})
	return _IOUSBHostStreamClass
}

// GetIOUSBHostStreamClass returns the class object for IOUSBHostStream.
func GetIOUSBHostStreamClass() IOUSBHostStreamClass {
	return getIOUSBHostStreamClass()
}

type IOUSBHostStreamClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOUSBHostStreamClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOUSBHostStreamClass) Alloc() IOUSBHostStream {
	rv := objc.Send[IOUSBHostStream](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// The class responsible for sending stream data for function drivers.
//
// # Overview
//
// The [IOUSBHostPipe.CopyStreamWithStreamIDError] method creates stream
// objects.
//
// # Sending I/O
//
//   - [IOUSBHostStream.EnqueueIORequestWithDataErrorCompletionHandler]: Enqueues an input/output request on the stream.
//   - [IOUSBHostStream.AbortWithOptionError]: Aborts pending input/output requests.
//   - [IOUSBHostStream.AbortWithError]: Aborts pending input/output requests synchronously.
//
// # Getting the Pipe Object
//
//   - [IOUSBHostStream.HostPipe]: The pipe that creates the stream.
//
// # Getting the Stream ID
//
//   - [IOUSBHostStream.StreamID]: The ID for the stream.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostStream
type IOUSBHostStream struct {
	IOUSBHostIOSource
}

// IOUSBHostStreamFromID constructs a [IOUSBHostStream] from an objc.ID.
//
// The class responsible for sending stream data for function drivers.
func IOUSBHostStreamFromID(id objc.ID) IOUSBHostStream {
	return IOUSBHostStream{IOUSBHostIOSource: IOUSBHostIOSourceFromID(id)}
}

// NOTE: IOUSBHostStream adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [IOUSBHostStream] class.
//
// # Sending I/O
//
//   - [IIOUSBHostStream.EnqueueIORequestWithDataErrorCompletionHandler]: Enqueues an input/output request on the stream.
//   - [IIOUSBHostStream.AbortWithOptionError]: Aborts pending input/output requests.
//   - [IIOUSBHostStream.AbortWithError]: Aborts pending input/output requests synchronously.
//
// # Getting the Pipe Object
//
//   - [IIOUSBHostStream.HostPipe]: The pipe that creates the stream.
//
// # Getting the Stream ID
//
//   - [IIOUSBHostStream.StreamID]: The ID for the stream.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostStream
type IIOUSBHostStream interface {
	IIOUSBHostIOSource

	// Topic: Sending I/O

	// Enqueues an input/output request on the stream.
	EnqueueIORequestWithDataErrorCompletionHandler(data foundation.NSMutableData, error_ foundation.NSError, completionHandler Int32Uint32Handler) bool
	// Aborts pending input/output requests.
	AbortWithOptionError(option IOUSBHostAbortOption) (bool, error)
	// Aborts pending input/output requests synchronously.
	AbortWithError() (bool, error)

	// Topic: Getting the Pipe Object

	// The pipe that creates the stream.
	HostPipe() IIOUSBHostPipe

	// Topic: Getting the Stream ID

	// The ID for the stream.
	StreamID() uint
}

// Init initializes the instance.
func (u IOUSBHostStream) Init() IOUSBHostStream {
	rv := objc.Send[IOUSBHostStream](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u IOUSBHostStream) Autorelease() IOUSBHostStream {
	rv := objc.Send[IOUSBHostStream](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOUSBHostStream creates a new IOUSBHostStream instance.
func NewIOUSBHostStream() IOUSBHostStream {
	class := getIOUSBHostStreamClass()
	rv := objc.Send[IOUSBHostStream](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Enqueues an input/output request on the stream.
//
// data: An [NSMutableData] object defining the memory to use for the transfer.
//
// completionHandler: An [IOUSBHostCompletionHandler] that runs when the request completes. The
// `completionHandler` doesn’t run if the method returns an error.
//
// # Discussion
//
// This method sends an asynchronous request on the stream.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostStream/enqueueIORequest(with:completionHandler:)
//
// [NSMutableData]: https://developer.apple.com/documentation/Foundation/NSMutableData
func (u IOUSBHostStream) EnqueueIORequestWithDataErrorCompletionHandler(data foundation.NSMutableData, error_ foundation.NSError, completionHandler Int32Uint32Handler) bool {
	_block2, _ := NewInt32Uint32Block(completionHandler)
	rv := objc.Send[bool](u.ID, objc.Sel("enqueueIORequestWithData:error:completionHandler:"), data, error_, _block2)
	return rv
}

// Aborts pending input/output requests.
//
// option: A set of options. [IOUSBHostAbortOption.synchronous] is the default.
//
// # Discussion
//
// Set the stream context as nonactive on the device with an out-of-band
// (class-defined) mechanism before calling this method, in accordance with
// USB 3.2, 8.12.1.4. The device won’t select a nonactive stream to become
// the current stream on the endpoint.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostStream/abort(with:)
//
// [IOUSBHostAbortOption.synchronous]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostAbortOption/synchronous
func (u IOUSBHostStream) AbortWithOptionError(option IOUSBHostAbortOption) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](u.ID, objc.Sel("abortWithOption:error:"), option, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("abortWithOption:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Aborts pending input/output requests synchronously.
//
// # Discussion
//
// Set the stream context as nonactive on the device with an out-of-band
// (class-defined) mechanism before calling this method, in accordance with
// USB 3.2, 8.12.1.4. The device won’t select a nonactive stream to become
// the current stream on the endpoint.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostStream/abort()
func (u IOUSBHostStream) AbortWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](u.ID, objc.Sel("abortWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("abortWithError: returned NO with nil NSError")
	}
	return rv, nil

}

// The pipe that creates the stream.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostStream/hostPipe
func (u IOUSBHostStream) HostPipe() IIOUSBHostPipe {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("hostPipe"))
	return IOUSBHostPipeFromID(objc.ID(rv))
}

// The ID for the stream.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostStream/streamID
func (u IOUSBHostStream) StreamID() uint {
	rv := objc.Send[uint](u.ID, objc.Sel("streamID"))
	return rv
}
