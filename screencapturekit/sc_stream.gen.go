// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"context"
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SCStream] class.
var (
	_SCStreamClass     SCStreamClass
	_SCStreamClassOnce sync.Once
)

func getSCStreamClass() SCStreamClass {
	_SCStreamClassOnce.Do(func() {
		_SCStreamClass = SCStreamClass{class: objc.GetClass("SCStream")}
	})
	return _SCStreamClass
}

// GetSCStreamClass returns the class object for SCStream.
func GetSCStreamClass() SCStreamClass {
	return getSCStreamClass()
}

type SCStreamClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (sc SCStreamClass) Alloc() SCStream {
	rv := objc.Send[SCStream](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}







// An instance that represents a stream of shareable content.
//
// # Overview
// 
// Use a stream to capture video of screen content like apps and windows.
// Create a content stream by passing it an instance of [SCContentFilter] and
// an [SCStreamConfiguration] object. The stream uses the filter to determine
// which screen content to capture, and uses the configuration data to
// configure the output.
//
// # Creating a stream
//
//   - [SCStream.InitWithFilterConfigurationDelegate]: Creates a stream with a content filter and configuration.
//
// # Updating stream configuration
//
//   - [SCStream.UpdateConfigurationCompletionHandler]: Updates the stream with a new configuration.
//   - [SCStream.UpdateContentFilterCompletionHandler]: Updates the stream by applying a new content filter.
//
// # Adding and removing stream output
//
//   - [SCStream.AddStreamOutputTypeSampleHandlerQueueError]: Adds a destination that receives the stream output.
//   - [SCStream.RemoveStreamOutputTypeError]: Removes a destination from receiving stream output.
//
// # Adding and removing recording output
//
//   - [SCStream.AddRecordingOutputError]
//   - [SCStream.RemoveRecordingOutputError]
//
// # Starting and stopping a stream
//
//   - [SCStream.StartCaptureWithCompletionHandler]: Starts the stream with a callback to indicate whether it successfully starts.
//   - [SCStream.StopCaptureWithCompletionHandler]: Stops the stream.
//
// # Stream synchronization
//
//   - [SCStream.SynchronizationClock]: A clock to use for output synchronization.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStream
type SCStream struct {
	objectivec.Object
}

// SCStreamFromID constructs a [SCStream] from an objc.ID.
//
// An instance that represents a stream of shareable content.
func SCStreamFromID(id objc.ID) SCStream {
	return SCStream{objectivec.Object{id}}
}
// NOTE: SCStream adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [SCStream] class.
//
// # Creating a stream
//
//   - [ISCStream.InitWithFilterConfigurationDelegate]: Creates a stream with a content filter and configuration.
//
// # Updating stream configuration
//
//   - [ISCStream.UpdateConfigurationCompletionHandler]: Updates the stream with a new configuration.
//   - [ISCStream.UpdateContentFilterCompletionHandler]: Updates the stream by applying a new content filter.
//
// # Adding and removing stream output
//
//   - [ISCStream.AddStreamOutputTypeSampleHandlerQueueError]: Adds a destination that receives the stream output.
//   - [ISCStream.RemoveStreamOutputTypeError]: Removes a destination from receiving stream output.
//
// # Adding and removing recording output
//
//   - [ISCStream.AddRecordingOutputError]
//   - [ISCStream.RemoveRecordingOutputError]
//
// # Starting and stopping a stream
//
//   - [ISCStream.StartCaptureWithCompletionHandler]: Starts the stream with a callback to indicate whether it successfully starts.
//   - [ISCStream.StopCaptureWithCompletionHandler]: Stops the stream.
//
// # Stream synchronization
//
//   - [ISCStream.SynchronizationClock]: A clock to use for output synchronization.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStream
type ISCStream interface {
	objectivec.IObject

	// Topic: Creating a stream

	// Creates a stream with a content filter and configuration.
	InitWithFilterConfigurationDelegate(contentFilter ISCContentFilter, streamConfig ISCStreamConfiguration, delegate SCStreamDelegate) SCStream

	// Topic: Updating stream configuration

	// Updates the stream with a new configuration.
	UpdateConfigurationCompletionHandler(streamConfig ISCStreamConfiguration, completionHandler ErrorHandler)
	// Updates the stream by applying a new content filter.
	UpdateContentFilterCompletionHandler(contentFilter ISCContentFilter, completionHandler ErrorHandler)

	// Topic: Adding and removing stream output

	// Adds a destination that receives the stream output.
	AddStreamOutputTypeSampleHandlerQueueError(output SCStreamOutput, type_ SCStreamOutputType, sampleHandlerQueue dispatch.Queue) (bool, error)
	// Removes a destination from receiving stream output.
	RemoveStreamOutputTypeError(output SCStreamOutput, type_ SCStreamOutputType) (bool, error)

	// Topic: Adding and removing recording output

	AddRecordingOutputError(recordingOutput ISCRecordingOutput) (bool, error)
	RemoveRecordingOutputError(recordingOutput ISCRecordingOutput) (bool, error)

	// Topic: Starting and stopping a stream

	// Starts the stream with a callback to indicate whether it successfully starts.
	StartCaptureWithCompletionHandler(completionHandler ErrorHandler)
	// Stops the stream.
	StopCaptureWithCompletionHandler(completionHandler ErrorHandler)

	// Topic: Stream synchronization

	// A clock to use for output synchronization.
	SynchronizationClock() objectivec.IObject
}





// Init initializes the instance.
func (s SCStream) Init() SCStream {
	rv := objc.Send[SCStream](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SCStream) Autorelease() SCStream {
	rv := objc.Send[SCStream](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSCStream creates a new SCStream instance.
func NewSCStream() SCStream {
	class := getSCStreamClass()
	rv := objc.Send[SCStream](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Creates a stream with a content filter and configuration.
//
// contentFilter: The content to capture.
//
// streamConfig: The configuration to apply to the stream.
//
// delegate: An optional object that responds to stream events.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStream/init(filter:configuration:delegate:)
func NewStreamWithFilterConfigurationDelegate(contentFilter ISCContentFilter, streamConfig ISCStreamConfiguration, delegate SCStreamDelegate) SCStream {
	instance := getSCStreamClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFilter:configuration:delegate:"), contentFilter, streamConfig, delegate)
	return SCStreamFromID(rv)
}







// Creates a stream with a content filter and configuration.
//
// contentFilter: The content to capture.
//
// streamConfig: The configuration to apply to the stream.
//
// delegate: An optional object that responds to stream events.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStream/init(filter:configuration:delegate:)
func (s SCStream) InitWithFilterConfigurationDelegate(contentFilter ISCContentFilter, streamConfig ISCStreamConfiguration, delegate SCStreamDelegate) SCStream {
	rv := objc.Send[SCStream](s.ID, objc.Sel("initWithFilter:configuration:delegate:"), contentFilter, streamConfig, delegate)
	return rv
}

// Updates the stream with a new configuration.
//
// streamConfig: An object that provides the updated stream configuration.
//
// completionHandler: A completion handler the system calls when this method completes. It
// includes an error if the update fails.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStream/updateConfiguration(_:completionHandler:)
func (s SCStream) UpdateConfigurationCompletionHandler(streamConfig ISCStreamConfiguration, completionHandler ErrorHandler) {
		_block1, _cleanup1 := NewErrorBlock(completionHandler)
	defer _cleanup1()
		objc.Send[objc.ID](s.ID, objc.Sel("updateConfiguration:completionHandler:"), streamConfig, _block1)
}

// Updates the stream by applying a new content filter.
//
// contentFilter: The content filter to apply.
//
// completionHandler: A completion handler the system calls when this method completes. It
// includes an error if the update fails.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStream/updateContentFilter(_:completionHandler:)
func (s SCStream) UpdateContentFilterCompletionHandler(contentFilter ISCContentFilter, completionHandler ErrorHandler) {
		_block1, _cleanup1 := NewErrorBlock(completionHandler)
	defer _cleanup1()
		objc.Send[objc.ID](s.ID, objc.Sel("updateContentFilter:completionHandler:"), contentFilter, _block1)
}

// Adds a destination that receives the stream output.
//
// output: The object that conforms to the stream output protocol.
//
// type: The stream output type.
//
// sampleHandlerQueue: The queue that receives the stream output.
//
// # Discussion
// 
// Use this method to attach an object that conforms to [SCStreamOutput] to
// receive stream content. Optionally, provide a [DispatchQueue] to send
// output to a queue that’s responsible for processing the output.
//
// [DispatchQueue]: https://developer.apple.com/documentation/Dispatch/DispatchQueue
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStream/addStreamOutput(_:type:sampleHandlerQueue:)
func (s SCStream) AddStreamOutputTypeSampleHandlerQueueError(output SCStreamOutput, type_ SCStreamOutputType, sampleHandlerQueue dispatch.Queue) (bool, error) {
			var errorPtr objc.ID
	rv := objc.Send[bool](s.ID, objc.Sel("addStreamOutput:type:sampleHandlerQueue:error:"), output, type_, uintptr(sampleHandlerQueue.Handle()), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("addStreamOutput:type:sampleHandlerQueue:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Removes a destination from receiving stream output.
//
// output: The object to remove that conforms to the stream output protocol.
//
// type: The stream output type.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStream/removeStreamOutput(_:type:)
func (s SCStream) RemoveStreamOutputTypeError(output SCStreamOutput, type_ SCStreamOutputType) (bool, error) {
			var errorPtr objc.ID
	rv := objc.Send[bool](s.ID, objc.Sel("removeStreamOutput:type:error:"), output, type_, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("removeStreamOutput:type:error: returned NO with nil NSError")
	}
	return rv, nil

}

//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStream/addRecordingOutput(_:)
func (s SCStream) AddRecordingOutputError(recordingOutput ISCRecordingOutput) (bool, error) {
			var errorPtr objc.ID
	rv := objc.Send[bool](s.ID, objc.Sel("addRecordingOutput:error:"), recordingOutput, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("addRecordingOutput:error: returned NO with nil NSError")
	}
	return rv, nil

}

//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStream/removeRecordingOutput(_:)
func (s SCStream) RemoveRecordingOutputError(recordingOutput ISCRecordingOutput) (bool, error) {
			var errorPtr objc.ID
	rv := objc.Send[bool](s.ID, objc.Sel("removeRecordingOutput:error:"), recordingOutput, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("removeRecordingOutput:error: returned NO with nil NSError")
	}
	return rv, nil

}

// Starts the stream with a callback to indicate whether it successfully
// starts.
//
// completionHandler: A completion handler that provides an error if the stream fails to start.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStream/startCapture(completionHandler:)
func (s SCStream) StartCaptureWithCompletionHandler(completionHandler ErrorHandler) {
		_block0, _cleanup0 := NewErrorBlock(completionHandler)
	defer _cleanup0()
		objc.Send[objc.ID](s.ID, objc.Sel("startCaptureWithCompletionHandler:"), _block0)
}

// Stops the stream.
//
// completionHandler: A completion handler that provides an error if the stream fails to stop.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStream/stopCapture(completionHandler:)
func (s SCStream) StopCaptureWithCompletionHandler(completionHandler ErrorHandler) {
		_block0, _cleanup0 := NewErrorBlock(completionHandler)
	defer _cleanup0()
		objc.Send[objc.ID](s.ID, objc.Sel("stopCaptureWithCompletionHandler:"), _block0)
}












// A clock to use for output synchronization.
//
// # Discussion
// 
// The synchronization clock provides the timebase for sample buffers that the
// stream outputs. Use it to synchronize with the clocks of other media
// sources, such as the [synchronizationClock] of [AVCaptureSession].
//
// [AVCaptureSession]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession
// [synchronizationClock]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/synchronizationClock
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStream/synchronizationClock
func (s SCStream) SynchronizationClock() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("synchronizationClock"))
	return objectivec.Object{ID: rv}
}


















// UpdateConfiguration is a synchronous wrapper around [SCStream.UpdateConfigurationCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (s SCStream) UpdateConfiguration(ctx context.Context, streamConfig ISCStreamConfiguration) error {
	done := make(chan error, 1)
	s.UpdateConfigurationCompletionHandler(streamConfig, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// UpdateContentFilter is a synchronous wrapper around [SCStream.UpdateContentFilterCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (s SCStream) UpdateContentFilter(ctx context.Context, contentFilter ISCContentFilter) error {
	done := make(chan error, 1)
	s.UpdateContentFilterCompletionHandler(contentFilter, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// StartCapture is a synchronous wrapper around [SCStream.StartCaptureWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (s SCStream) StartCapture(ctx context.Context) error {
	done := make(chan error, 1)
	s.StartCaptureWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// StopCapture is a synchronous wrapper around [SCStream.StopCaptureWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (s SCStream) StopCapture(ctx context.Context) error {
	done := make(chan error, 1)
	s.StopCaptureWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}






