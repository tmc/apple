// Code generated from Apple documentation for IOUSBHost. DO NOT EDIT.

package iousbhost

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [IOUSBHostPipe] class.
var (
	_IOUSBHostPipeClass     IOUSBHostPipeClass
	_IOUSBHostPipeClassOnce sync.Once
)

func getIOUSBHostPipeClass() IOUSBHostPipeClass {
	_IOUSBHostPipeClassOnce.Do(func() {
		_IOUSBHostPipeClass = IOUSBHostPipeClass{class: objc.GetClass("IOUSBHostPipe")}
	})
	return _IOUSBHostPipeClass
}

// GetIOUSBHostPipeClass returns the class object for IOUSBHostPipe.
func GetIOUSBHostPipeClass() IOUSBHostPipeClass {
	return getIOUSBHostPipeClass()
}

type IOUSBHostPipeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOUSBHostPipeClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOUSBHostPipeClass) Alloc() IOUSBHostPipe {
	rv := objc.Send[IOUSBHostPipe](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// The class that sends control, bulk, interrupt, and isochronous input/output
// requests for function drivers, and manages stream capabilities.
//
// # Overview
//
// The client creates pipe objects using
// [IOUSBHostInterface.CopyPipeWithAddressError].
//
// # Sending Bulk and Interrupt I/O
//
//   - [IOUSBHostPipe.EnqueueIORequestWithDataCompletionTimeoutErrorCompletionHandler]: Enqueues an input/output request on the pipe.
//   - [IOUSBHostPipe.ClearStallWithError]: Clears the halt condition of the pipe.
//
// # Managing Periodic Bandwidth
//
//   - [IOUSBHostPipe.AdjustPipeWithDescriptorsError]: Adjusts the behavior of periodic endpoints to consume a different amount of bus bandwidth.
//   - [IOUSBHostPipe.Descriptors]: A property that retrieves the current endpoint descriptors controlling the endpoint.
//   - [IOUSBHostPipe.OriginalDescriptors]: A property that retrieves the original endpoint descriptors from the pipe at the point of creation.
//
// # Enabling Power Savings
//
//   - [IOUSBHostPipe.SetIdleTimeoutError]: Sets the desired idle suspend timeout for the interface.
//   - [IOUSBHostPipe.IdleTimeout]: A property that retrieves the current idle suspend timeout.
//
// # Managing Streams
//
//   - [IOUSBHostPipe.EnableStreamsWithError]: Enables streams for the pipe.
//   - [IOUSBHostPipe.CopyStreamWithStreamIDError]: Returns the stream for a stream ID.
//   - [IOUSBHostPipe.DisableStreamsWithError]: Disables streams for the pipe.
//
// # Instance Methods
//
//   - [IOUSBHostPipe.EnqueueIORequestWithDataTransactionListTransactionListCountFirstFrameNumberOptionsErrorCompletionHandler]
//   - [IOUSBHostPipe.SendIORequestWithDataTransactionListTransactionListCountFirstFrameNumberOptionsError]
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostPipe
type IOUSBHostPipe struct {
	IOUSBHostIOSource
}

// IOUSBHostPipeFromID constructs a [IOUSBHostPipe] from an objc.ID.
//
// The class that sends control, bulk, interrupt, and isochronous input/output
// requests for function drivers, and manages stream capabilities.
func IOUSBHostPipeFromID(id objc.ID) IOUSBHostPipe {
	return IOUSBHostPipe{IOUSBHostIOSource: IOUSBHostIOSourceFromID(id)}
}

// NOTE: IOUSBHostPipe adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [IOUSBHostPipe] class.
//
// # Sending Bulk and Interrupt I/O
//
//   - [IIOUSBHostPipe.EnqueueIORequestWithDataCompletionTimeoutErrorCompletionHandler]: Enqueues an input/output request on the pipe.
//   - [IIOUSBHostPipe.ClearStallWithError]: Clears the halt condition of the pipe.
//
// # Managing Periodic Bandwidth
//
//   - [IIOUSBHostPipe.AdjustPipeWithDescriptorsError]: Adjusts the behavior of periodic endpoints to consume a different amount of bus bandwidth.
//   - [IIOUSBHostPipe.Descriptors]: A property that retrieves the current endpoint descriptors controlling the endpoint.
//   - [IIOUSBHostPipe.OriginalDescriptors]: A property that retrieves the original endpoint descriptors from the pipe at the point of creation.
//
// # Enabling Power Savings
//
//   - [IIOUSBHostPipe.SetIdleTimeoutError]: Sets the desired idle suspend timeout for the interface.
//   - [IIOUSBHostPipe.IdleTimeout]: A property that retrieves the current idle suspend timeout.
//
// # Managing Streams
//
//   - [IIOUSBHostPipe.EnableStreamsWithError]: Enables streams for the pipe.
//   - [IIOUSBHostPipe.CopyStreamWithStreamIDError]: Returns the stream for a stream ID.
//   - [IIOUSBHostPipe.DisableStreamsWithError]: Disables streams for the pipe.
//
// # Instance Methods
//
//   - [IIOUSBHostPipe.EnqueueIORequestWithDataTransactionListTransactionListCountFirstFrameNumberOptionsErrorCompletionHandler]
//   - [IIOUSBHostPipe.SendIORequestWithDataTransactionListTransactionListCountFirstFrameNumberOptionsError]
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostPipe
type IIOUSBHostPipe interface {
	IIOUSBHostIOSource

	// Topic: Sending Bulk and Interrupt I/O

	// Enqueues an input/output request on the pipe.
	EnqueueIORequestWithDataCompletionTimeoutErrorCompletionHandler(data foundation.NSMutableData, completionTimeout foundation.NSTimeInterval, error_ foundation.NSError, completionHandler Int32Uint32Handler) bool
	// Clears the halt condition of the pipe.
	ClearStallWithError() (bool, error)

	// Topic: Managing Periodic Bandwidth

	// Adjusts the behavior of periodic endpoints to consume a different amount of bus bandwidth.
	AdjustPipeWithDescriptorsError(descriptors *IOUSBHostIOSourceDescriptors) (bool, error)
	// A property that retrieves the current endpoint descriptors controlling the endpoint.
	Descriptors() *IOUSBHostIOSourceDescriptors
	// A property that retrieves the original endpoint descriptors from the pipe at the point of creation.
	OriginalDescriptors() *IOUSBHostIOSourceDescriptors

	// Topic: Enabling Power Savings

	// Sets the desired idle suspend timeout for the interface.
	SetIdleTimeoutError(idleTimeout foundation.NSTimeInterval) (bool, error)
	// A property that retrieves the current idle suspend timeout.
	IdleTimeout() foundation.NSTimeInterval

	// Topic: Managing Streams

	// Enables streams for the pipe.
	EnableStreamsWithError() (bool, error)
	// Returns the stream for a stream ID.
	CopyStreamWithStreamIDError(streamID uint) (IIOUSBHostStream, error)
	// Disables streams for the pipe.
	DisableStreamsWithError() (bool, error)

	// Topic: Instance Methods

	EnqueueIORequestWithDataTransactionListTransactionListCountFirstFrameNumberOptionsErrorCompletionHandler(data foundation.NSMutableData, transactionList *IOUSBHostIsochronousTransaction, transactionListCount uint, firstFrameNumber uint64, options IOUSBHostIsochronousTransferOptions, error_ foundation.NSError, completionHandler Int32IOUSBHostIsochronousTransactionHandler) bool
	SendIORequestWithDataTransactionListTransactionListCountFirstFrameNumberOptionsError(data foundation.NSMutableData, transactionList *IOUSBHostIsochronousTransaction, transactionListCount uint, firstFrameNumber uint64, options IOUSBHostIsochronousTransferOptions) (bool, error)
}

// Init initializes the instance.
func (u IOUSBHostPipe) Init() IOUSBHostPipe {
	rv := objc.Send[IOUSBHostPipe](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u IOUSBHostPipe) Autorelease() IOUSBHostPipe {
	rv := objc.Send[IOUSBHostPipe](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOUSBHostPipe creates a new IOUSBHostPipe instance.
func NewIOUSBHostPipe() IOUSBHostPipe {
	class := getIOUSBHostPipeClass()
	rv := objc.Send[IOUSBHostPipe](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Enqueues an input/output request on the pipe.
//
// data: An [NSMutableData] object defining the memory to use for the transfer. Use
// nil to send a zero-length packet.
//
// completionTimeout: A [TimeInterval] value representing the timeout of the request. If `0`, the
// request never times out. Use [IOUSBHostDefaultControlCompletionTimeout]
// unless there’s a need for a specific timeout.
//
// completionHandler: An [IOUSBHostCompletionHandler] that runs when the request completes, or
// times out after the call returns successfully. If the method returns with
// an error, the completion handler doesn’t run.
//
// # Discussion
//
// Use this method to issue an asynchronous input/output request on a bulk or
// interrupt pipe.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostPipe/enqueueIORequest(with:completionTimeout:completionHandler:)
//
// [NSMutableData]: https://developer.apple.com/documentation/Foundation/NSMutableData
// [IOUSBHostDefaultControlCompletionTimeout]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostDefaultControlCompletionTimeout
// [TimeInterval]: https://developer.apple.com/documentation/Foundation/TimeInterval
func (u IOUSBHostPipe) EnqueueIORequestWithDataCompletionTimeoutErrorCompletionHandler(data foundation.NSMutableData, completionTimeout foundation.NSTimeInterval, error_ foundation.NSError, completionHandler Int32Uint32Handler) bool {
	_block3, _ := NewInt32Uint32Block(completionHandler)
	rv := objc.Send[bool](u.ID, objc.Sel("enqueueIORequestWithData:completionTimeout:error:completionHandler:"), data, completionTimeout, error_, _block3)
	return rv
}

// Clears the halt condition of the pipe.
//
// # Discussion
//
// When a bulk or interrupt USB endpoint encounters any input/output error
// other than a timeout, it transitions to a halted state. It must also clear
// to perform additional input/output requests on the endpoint.
//
// This method clears the halted condition for the endpoint. It also sends a
// `CLEAR_TT_BUFFER` control request (See USB 2.0, 11.24.2.3.) to an
// intermediate hub, if necessary. All pending input/output requests on the
// endpoint abort, and the data toggle for the endpoint resets.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostPipe/clearStall()
func (u IOUSBHostPipe) ClearStallWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](u.ID, objc.Sel("clearStallWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("clearStallWithError: returned NO with nil NSError")
	}
	return rv, nil

}

// Adjusts the behavior of periodic endpoints to consume a different amount of
// bus bandwidth.
//
// descriptors: A reference to [IOUSBHostIOSourceDescriptors] describing the new endpoint
// policy.
//
// # Discussion
//
// During creation, periodic (interrupt and isochronous) endpoints reserve bus
// bandwidth to allow for maximum packet size, mult (the maximum number of
// packets that this endpoint supports), burst size, and the endpoint service
// interval.
//
// If an endpoint won’t use all of the allocated bandwidth, use
// [IOUSBHostPipe.AdjustPipeWithDescriptorsError] to reduce the bandwidth
// reserved for the endpoint. Copy the original endpoint descriptors, adjust
// maximum packet size, mult, burst size, and interval, then pass to
// [IOUSBHostPipe.AdjustPipeWithDescriptorsError]. The altered descriptors
// must pass validation from the kernel for policy changes to process.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostPipe/adjust(with:)
//
// [IOUSBHostIOSourceDescriptors]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostIOSourceDescriptors
func (u IOUSBHostPipe) AdjustPipeWithDescriptorsError(descriptors *IOUSBHostIOSourceDescriptors) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](u.ID, objc.Sel("adjustPipeWithDescriptors:error:"), unsafe.Pointer(descriptors), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("adjustPipeWithDescriptors:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Sets the desired idle suspend timeout for the interface.
//
// idleTimeout: The amount of time after all pipes are idle to wait before suspending the
// device.
//
// # Discussion
//
// After the interface idles, it defers electrical suspension of the device
// for the specified duration.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostPipe/setIdleTimeout(_:)
func (u IOUSBHostPipe) SetIdleTimeoutError(idleTimeout foundation.NSTimeInterval) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](u.ID, objc.Sel("setIdleTimeout:error:"), idleTimeout, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setIdleTimeout:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Enables streams for the pipe.
//
// # Discussion
//
// This method changes the operational mode of the pipe to allow streaming
// endpoint transfers. Call this method before
// [IOUSBHostPipe.CopyStreamWithStreamIDError].
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostPipe/enableStreams()
func (u IOUSBHostPipe) EnableStreamsWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](u.ID, objc.Sel("enableStreamsWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("enableStreamsWithError: returned NO with nil NSError")
	}
	return rv, nil

}

// Returns the stream for a stream ID.
//
// streamID: A stream ID in the range of 1 to n. Retrieve n can by calling
// [IOUSBGetEndpointMaxStreams] with the [IOUSBEndpointDescriptor].
//
// # Return Value
//
// A pointer to an [IOUSBHostStream]; otherwise, `nil` if the device or the
// underlying host controller doesn’t support the specified stream ID.
//
// # Discussion
//
// Call [IOUSBHostPipe.EnableStreamsWithError] before this method.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostPipe/copyStream(withStreamID:)
//
// [IOUSBEndpointDescriptor]: https://developer.apple.com/documentation/iokit/iousbendpointdescriptor
func (u IOUSBHostPipe) CopyStreamWithStreamIDError(streamID uint) (IIOUSBHostStream, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](u.ID, objc.Sel("copyStreamWithStreamID:error:"), streamID, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return IOUSBHostStream{}, foundation.NSErrorFrom(errorPtr)
	}
	return IOUSBHostStreamFromID(rv), nil

}

// Disables streams for the pipe.
//
// # Discussion
//
// This method changes the operational mode of the [IOUSBHostPipe] to disable
// streaming endpoint transfers. Before calling this method, set all stream
// contexts as nonactive on the device through an out-of-band (class-defined)
// mechanism, in accordance with USB 3.2, 8.12.1.4. This is necessary, as the
// method synchronously aborts any outstanding calls on existing
// [IOUSBHostStream] objects.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostPipe/disableStreams()
//
// [IOUSBHostPipe]: https://developer.apple.com/documentation/kernel/iousbhostpipe
func (u IOUSBHostPipe) DisableStreamsWithError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](u.ID, objc.Sel("disableStreamsWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("disableStreamsWithError: returned NO with nil NSError")
	}
	return rv, nil

}

// # Discussion
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostPipe/enqueueIORequest(with:transactionList:transactionListCount:firstFrameNumber:options:completionHandler:)
func (u IOUSBHostPipe) EnqueueIORequestWithDataTransactionListTransactionListCountFirstFrameNumberOptionsErrorCompletionHandler(data foundation.NSMutableData, transactionList *IOUSBHostIsochronousTransaction, transactionListCount uint, firstFrameNumber uint64, options IOUSBHostIsochronousTransferOptions, error_ foundation.NSError, completionHandler Int32IOUSBHostIsochronousTransactionHandler) bool {
	_block6, _ := NewInt32IOUSBHostIsochronousTransactionBlock(completionHandler)
	rv := objc.Send[bool](u.ID, objc.Sel("enqueueIORequestWithData:transactionList:transactionListCount:firstFrameNumber:options:error:completionHandler:"), data, transactionList, transactionListCount, firstFrameNumber, options, error_, _block6)
	return rv
}

// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostPipe/sendIORequest(with:transactionList:transactionListCount:firstFrameNumber:options:)
func (u IOUSBHostPipe) SendIORequestWithDataTransactionListTransactionListCountFirstFrameNumberOptionsError(data foundation.NSMutableData, transactionList *IOUSBHostIsochronousTransaction, transactionListCount uint, firstFrameNumber uint64, options IOUSBHostIsochronousTransferOptions) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](u.ID, objc.Sel("sendIORequestWithData:transactionList:transactionListCount:firstFrameNumber:options:error:"), data, unsafe.Pointer(transactionList), transactionListCount, firstFrameNumber, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("sendIORequestWithData:transactionList:transactionListCount:firstFrameNumber:options:error: returned NO with nil NSError")
	}
	return rv, nil

}

// A property that retrieves the current endpoint descriptors controlling the
// endpoint.
//
// # Return Value
//
// The current [IOUSBHostIOSourceDescriptors].
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostPipe/descriptors
//
// [IOUSBHostIOSourceDescriptors]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostIOSourceDescriptors
func (u IOUSBHostPipe) Descriptors() *IOUSBHostIOSourceDescriptors {
	rv := objc.Send[unsafe.Pointer](u.ID, objc.Sel("descriptors"))
	return (*IOUSBHostIOSourceDescriptors)(rv)
}

// A property that retrieves the original endpoint descriptors from the pipe
// at the point of creation.
//
// # Return Value
//
// The original [IOUSBHostIOSourceDescriptors].
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostPipe/originalDescriptors
//
// [IOUSBHostIOSourceDescriptors]: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostIOSourceDescriptors
func (u IOUSBHostPipe) OriginalDescriptors() *IOUSBHostIOSourceDescriptors {
	rv := objc.Send[unsafe.Pointer](u.ID, objc.Sel("originalDescriptors"))
	return (*IOUSBHostIOSourceDescriptors)(rv)
}

// A property that retrieves the current idle suspend timeout.
//
// # Return Value
//
// The amount of time after all pipes are idle to wait before suspending the
// device.
//
// See: https://developer.apple.com/documentation/IOUSBHost/IOUSBHostPipe/idleTimeout
func (u IOUSBHostPipe) IdleTimeout() foundation.NSTimeInterval {
	rv := objc.Send[foundation.NSTimeInterval](u.ID, objc.Sel("idleTimeout"))
	return foundation.NSTimeInterval(rv)
}
