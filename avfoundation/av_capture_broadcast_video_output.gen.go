// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AVCaptureBroadcastVideoOutput] class.
var (
	_AVCaptureBroadcastVideoOutputClass     AVCaptureBroadcastVideoOutputClass
	_AVCaptureBroadcastVideoOutputClassOnce sync.Once
)

func getAVCaptureBroadcastVideoOutputClass() AVCaptureBroadcastVideoOutputClass {
	_AVCaptureBroadcastVideoOutputClassOnce.Do(func() {
		_AVCaptureBroadcastVideoOutputClass = AVCaptureBroadcastVideoOutputClass{class: objc.GetClass("AVCaptureBroadcastVideoOutput")}
	})
	return _AVCaptureBroadcastVideoOutputClass
}

// GetAVCaptureBroadcastVideoOutputClass returns the class object for AVCaptureBroadcastVideoOutput.
func GetAVCaptureBroadcastVideoOutputClass() AVCaptureBroadcastVideoOutputClass {
	return getAVCaptureBroadcastVideoOutputClass()
}

type AVCaptureBroadcastVideoOutputClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVCaptureBroadcastVideoOutputClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCaptureBroadcastVideoOutputClass) Alloc() AVCaptureBroadcastVideoOutput {
	rv := objc.Send[AVCaptureBroadcastVideoOutput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// [AVCaptureBroadcastVideoOutput] is a subclass of [AVCaptureOutput] that
// delivers broadcast-quality video and ancillary data through the device’s
// DisplayPort hardware interface (USB-C DP Alt Mode)
//
// # Overview
//
// Not all [AVCaptureDeviceFormat] instances support
// [AVCaptureBroadcastVideoOutput]. Before adding this output to a session,
// check the device format’s
// `AVCaptureDeviceFormat.UnsupportedCaptureOutputClasses()` property to
// verify that [AVCaptureBroadcastVideoOutput] is not listed. If the current
// format does not support broadcast video output, the connection will be
// marked inactive and no samples will be delivered.
//
// # Managing the Output
//
//   - [AVCaptureBroadcastVideoOutput.Delegate]: The receiver’s delegate.
//   - [AVCaptureBroadcastVideoOutput.DelegateCallbackQueue]: The dispatch queue on which all [AVCaptureBroadcastVideoOutputDelegate](<https://developer.apple.com/documentation/AVFoundation/AVCaptureBroadcastVideoOutputDelegate>) methods will be called.
//   - [AVCaptureBroadcastVideoOutput.SetDelegateQueue]: Sets the receiver’s delegate and the dispatch queue on which the delegate will be called.
//
// # Managing Video Output
//
//   - [AVCaptureBroadcastVideoOutput.VideoSettings]: The current video output settings for the broadcast video output.
//   - [AVCaptureBroadcastVideoOutput.MaxBufferedFrameCount]: This represents the maximum count of buffered frames. By default the value is 0, which means late frames are immediately dropped to maintain minimal latency.
//   - [AVCaptureBroadcastVideoOutput.SetMaxBufferedFrameCount]
//   - [AVCaptureBroadcastVideoOutput.ResetFrameBuffer]: Tells the broadcast video output to reset the frame buffer and drop all currently buffered frames.
//   - [AVCaptureBroadcastVideoOutput.DroppedFrameReplacementPolicy]: The strategy used to replace dropped video frames.
//   - [AVCaptureBroadcastVideoOutput.SetDroppedFrameReplacementPolicy]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureBroadcastVideoOutput
type AVCaptureBroadcastVideoOutput struct {
	AVCaptureOutput
}

// AVCaptureBroadcastVideoOutputFromID constructs a [AVCaptureBroadcastVideoOutput] from an objc.ID.
//
// [AVCaptureBroadcastVideoOutput] is a subclass of [AVCaptureOutput] that
// delivers broadcast-quality video and ancillary data through the device’s
// DisplayPort hardware interface (USB-C DP Alt Mode)
func AVCaptureBroadcastVideoOutputFromID(id objc.ID) AVCaptureBroadcastVideoOutput {
	return AVCaptureBroadcastVideoOutput{AVCaptureOutput: AVCaptureOutputFromID(id)}
}

// NOTE: AVCaptureBroadcastVideoOutput adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCaptureBroadcastVideoOutput] class.
//
// # Managing the Output
//
//   - [IAVCaptureBroadcastVideoOutput.Delegate]: The receiver’s delegate.
//   - [IAVCaptureBroadcastVideoOutput.DelegateCallbackQueue]: The dispatch queue on which all [AVCaptureBroadcastVideoOutputDelegate](<https://developer.apple.com/documentation/AVFoundation/AVCaptureBroadcastVideoOutputDelegate>) methods will be called.
//   - [IAVCaptureBroadcastVideoOutput.SetDelegateQueue]: Sets the receiver’s delegate and the dispatch queue on which the delegate will be called.
//
// # Managing Video Output
//
//   - [IAVCaptureBroadcastVideoOutput.VideoSettings]: The current video output settings for the broadcast video output.
//   - [IAVCaptureBroadcastVideoOutput.MaxBufferedFrameCount]: This represents the maximum count of buffered frames. By default the value is 0, which means late frames are immediately dropped to maintain minimal latency.
//   - [IAVCaptureBroadcastVideoOutput.SetMaxBufferedFrameCount]
//   - [IAVCaptureBroadcastVideoOutput.ResetFrameBuffer]: Tells the broadcast video output to reset the frame buffer and drop all currently buffered frames.
//   - [IAVCaptureBroadcastVideoOutput.DroppedFrameReplacementPolicy]: The strategy used to replace dropped video frames.
//   - [IAVCaptureBroadcastVideoOutput.SetDroppedFrameReplacementPolicy]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureBroadcastVideoOutput
type IAVCaptureBroadcastVideoOutput interface {
	IAVCaptureOutput

	// Topic: Managing the Output

	// The receiver’s delegate.
	Delegate() AVCaptureBroadcastVideoOutputDelegate
	// The dispatch queue on which all [AVCaptureBroadcastVideoOutputDelegate](<https://developer.apple.com/documentation/AVFoundation/AVCaptureBroadcastVideoOutputDelegate>) methods will be called.
	DelegateCallbackQueue() dispatch.Queue
	// Sets the receiver’s delegate and the dispatch queue on which the delegate will be called.
	SetDelegateQueue(delegate AVCaptureBroadcastVideoOutputDelegate, delegateCallbackQueue dispatch.Queue)

	// Topic: Managing Video Output

	// The current video output settings for the broadcast video output.
	VideoSettings() foundation.INSDictionary
	// This represents the maximum count of buffered frames. By default the value is 0, which means late frames are immediately dropped to maintain minimal latency.
	MaxBufferedFrameCount() int
	SetMaxBufferedFrameCount(value int)
	// Tells the broadcast video output to reset the frame buffer and drop all currently buffered frames.
	ResetFrameBuffer()
	// The strategy used to replace dropped video frames.
	DroppedFrameReplacementPolicy() AVCaptureBroadcastVideoOutputDroppedFrameReplacementPolicy
	SetDroppedFrameReplacementPolicy(value AVCaptureBroadcastVideoOutputDroppedFrameReplacementPolicy)
}

// Init initializes the instance.
func (c AVCaptureBroadcastVideoOutput) Init() AVCaptureBroadcastVideoOutput {
	rv := objc.Send[AVCaptureBroadcastVideoOutput](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCaptureBroadcastVideoOutput) Autorelease() AVCaptureBroadcastVideoOutput {
	rv := objc.Send[AVCaptureBroadcastVideoOutput](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureBroadcastVideoOutput creates a new AVCaptureBroadcastVideoOutput instance.
func NewAVCaptureBroadcastVideoOutput() AVCaptureBroadcastVideoOutput {
	class := getAVCaptureBroadcastVideoOutputClass()
	rv := objc.Send[AVCaptureBroadcastVideoOutput](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Sets the receiver’s delegate and the dispatch queue on which the delegate
// will be called.
//
// delegate: An object conforming to the [AVCaptureBroadcastVideoOutputDelegate]
// protocol that will receive broadcast video output notifications.
//
// delegateCallbackQueue: A dispatch queue on which all [AVCaptureBroadcastVideoOutputDelegate]
// methods will be called.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureBroadcastVideoOutput/setDelegate(_:queue:)
func (c AVCaptureBroadcastVideoOutput) SetDelegateQueue(delegate AVCaptureBroadcastVideoOutputDelegate, delegateCallbackQueue dispatch.Queue) {
	objc.Send[objc.ID](c.ID, objc.Sel("setDelegate:queue:"), delegate, uintptr(delegateCallbackQueue.Handle()))
}

// Tells the broadcast video output to reset the frame buffer and drop all
// currently buffered frames.
//
// # Discussion
//
// This method can be called when buffered video frames should be dropped.
// This will force all those frames to be dropped and reset the buffered frame
// count to 0.
//
// Use this method in scenarios where you need to clear pending frames, such
// as:
//
// - Pausing or stopping broadcast: Drop pending frames that should not be
// transmitted - Reducing accumulated latency: If buffering has built up
// significant delay, reset to return to real-time output
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureBroadcastVideoOutput/resetFrameBuffer()
func (c AVCaptureBroadcastVideoOutput) ResetFrameBuffer() {
	objc.Send[objc.ID](c.ID, objc.Sel("resetFrameBuffer"))
}

// The receiver’s delegate.
//
// # Discussion
//
// The value of this property is an object conforming to the
// [AVCaptureBroadcastVideoOutputDelegate] protocol that will be able to
// monitor the broadcast output operations.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureBroadcastVideoOutput/delegate
func (c AVCaptureBroadcastVideoOutput) Delegate() AVCaptureBroadcastVideoOutputDelegate {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("delegate"))
	return AVCaptureBroadcastVideoOutputDelegateObjectFromID(rv)
}

// The dispatch queue on which all [AVCaptureBroadcastVideoOutputDelegate]
// methods will be called.
//
// # Discussion
//
// The value of this property is a dispatch queue on which all delegate method
// calls will be serialized. If you have not called the
// [AVCaptureBroadcastVideoOutput.SetDelegateQueue] method, the value of this
// property will be `nil`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureBroadcastVideoOutput/delegateCallbackQueue
func (c AVCaptureBroadcastVideoOutput) DelegateCallbackQueue() dispatch.Queue {
	rv := objc.Send[uintptr](c.ID, objc.Sel("delegateCallbackQueue"))
	return dispatch.QueueFromHandle(rv)
}

// The current video output settings for the broadcast video output.
//
// # Discussion
//
// This read-only property reports the actual video format and output settings
// currently being used for broadcast video output. The value is a dictionary
// containing metadata descriptors conforming to SMPTE ST 377 (Material
// Exchange Format) using Universal Labels (ULs) for professional broadcast
// interoperability.
//
// The settings reflect the format negotiated between the camera capture
// pipeline and the connected broadcast video destination, taking into
// account:
//
// - Camera native capture format capabilities - Connected broadcast video
// destination capabilities - System performance constraints - Display
// transport bandwidth limitations
//
// This property will return `nil` when no broadcast video destination is
// connected or when the output pipeline is not active.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureBroadcastVideoOutput/videoSettings
func (c AVCaptureBroadcastVideoOutput) VideoSettings() foundation.INSDictionary {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("videoSettings"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// This represents the maximum count of buffered frames. By default the value
// is 0, which means late frames are immediately dropped to maintain minimal
// latency.
//
// # Discussion
//
// When set to a value greater than 0, the buffer absorbs minor timing jitter
// in the capture pipeline, reducing the possibility of dropping frames during
// temporary processing variations. Frames accumulate in the buffer up to the
// specified limit. Once the buffer reaches
// [AVCaptureBroadcastVideoOutput.MaxBufferedFrameCount], the oldest frame is
// removed to make room for each new incoming frame, maintaining a rolling
// window of buffered content.
//
// Calling [AVCaptureBroadcastVideoOutput.ResetFrameBuffer] clears all
// buffered frames and resets the buffer count back to 0, allowing the buffer
// to fill again from empty.
//
// The maximum supported value can be retrieved using
// [AVCaptureBroadcastVideoOutputClass.MaxSupportedBufferedFrameCount].
// Setting a value higher than the maximum supported value will raise an
// [NSInvalidArgumentException].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureBroadcastVideoOutput/maxBufferedFrameCount
func (c AVCaptureBroadcastVideoOutput) MaxBufferedFrameCount() int {
	rv := objc.Send[int](c.ID, objc.Sel("maxBufferedFrameCount"))
	return rv
}
func (c AVCaptureBroadcastVideoOutput) SetMaxBufferedFrameCount(value int) {
	objc.Send[struct{}](c.ID, objc.Sel("setMaxBufferedFrameCount:"), value)
}

// The strategy used to replace dropped video frames.
//
// # Discussion
//
// This property determines how the broadcast video output handles dropped
// frames. The default value is
// [AVCaptureBroadcastVideoOutputDroppedFrameReplacementPolicyRepeatPreviousFrame].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureBroadcastVideoOutput/droppedFrameReplacementPolicy-swift.property
func (c AVCaptureBroadcastVideoOutput) DroppedFrameReplacementPolicy() AVCaptureBroadcastVideoOutputDroppedFrameReplacementPolicy {
	rv := objc.Send[AVCaptureBroadcastVideoOutputDroppedFrameReplacementPolicy](c.ID, objc.Sel("droppedFrameReplacementPolicy"))
	return AVCaptureBroadcastVideoOutputDroppedFrameReplacementPolicy(rv)
}
func (c AVCaptureBroadcastVideoOutput) SetDroppedFrameReplacementPolicy(value AVCaptureBroadcastVideoOutputDroppedFrameReplacementPolicy) {
	objc.Send[struct{}](c.ID, objc.Sel("setDroppedFrameReplacementPolicy:"), value)
}

// The maximum value supported for maxBufferedFrameCount.
//
// # Discussion
//
// This class property returns the system-imposed limit for buffered frame
// count to ensure optimal performance and memory usage in broadcast
// workflows. The limit is determined based on system capabilities.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureBroadcastVideoOutput/maxSupportedBufferedFrameCount
func (_AVCaptureBroadcastVideoOutputClass AVCaptureBroadcastVideoOutputClass) MaxSupportedBufferedFrameCount() int {
	rv := objc.Send[int](objc.ID(_AVCaptureBroadcastVideoOutputClass.class), objc.Sel("maxSupportedBufferedFrameCount"))
	return rv
}
