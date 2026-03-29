// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [AVPlayerItemVideoOutput] class.
var (
	_AVPlayerItemVideoOutputClass     AVPlayerItemVideoOutputClass
	_AVPlayerItemVideoOutputClassOnce sync.Once
)

func getAVPlayerItemVideoOutputClass() AVPlayerItemVideoOutputClass {
	_AVPlayerItemVideoOutputClassOnce.Do(func() {
		_AVPlayerItemVideoOutputClass = AVPlayerItemVideoOutputClass{class: objc.GetClass("AVPlayerItemVideoOutput")}
	})
	return _AVPlayerItemVideoOutputClass
}

// GetAVPlayerItemVideoOutputClass returns the class object for AVPlayerItemVideoOutput.
func GetAVPlayerItemVideoOutputClass() AVPlayerItemVideoOutputClass {
	return getAVPlayerItemVideoOutputClass()
}

type AVPlayerItemVideoOutputClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVPlayerItemVideoOutputClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVPlayerItemVideoOutputClass) Alloc() AVPlayerItemVideoOutput {
	rv := objc.Send[AVPlayerItemVideoOutput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that outputs video frames from a player item.
//
// # Creating a video output
//
//   - [AVPlayerItemVideoOutput.InitWithOutputSettings]: Creates a video output object initialized with the specified output settings.
//
// # Configuring the delegate
//
//   - [AVPlayerItemVideoOutput.SetDelegateQueue]: Sets the delegate and dispatch queue for the receiver.
//   - [AVPlayerItemVideoOutput.Delegate]: The delegate for the video output object.
//   - [AVPlayerItemVideoOutput.DelegateQueue]: The dispatch queue on which to call delegate methods.
//
// # Notifying the delegate of changes
//
//   - [AVPlayerItemVideoOutput.RequestNotificationOfMediaDataChangeWithAdvanceInterval]: Tells the receiver that the video out put client is entering a quiescent state.
//
// # Getting pixel buffer data
//
//   - [AVPlayerItemVideoOutput.HasNewPixelBufferForItemTime]: Returns a Boolean value that indicates whether video output is available for the specified item time.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemVideoOutput
type AVPlayerItemVideoOutput struct {
	AVPlayerItemOutput
}

// AVPlayerItemVideoOutputFromID constructs a [AVPlayerItemVideoOutput] from an objc.ID.
//
// An object that outputs video frames from a player item.
func AVPlayerItemVideoOutputFromID(id objc.ID) AVPlayerItemVideoOutput {
	return AVPlayerItemVideoOutput{AVPlayerItemOutput: AVPlayerItemOutputFromID(id)}
}
// NOTE: AVPlayerItemVideoOutput adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVPlayerItemVideoOutput] class.
//
// # Creating a video output
//
//   - [IAVPlayerItemVideoOutput.InitWithOutputSettings]: Creates a video output object initialized with the specified output settings.
//
// # Configuring the delegate
//
//   - [IAVPlayerItemVideoOutput.SetDelegateQueue]: Sets the delegate and dispatch queue for the receiver.
//   - [IAVPlayerItemVideoOutput.Delegate]: The delegate for the video output object.
//   - [IAVPlayerItemVideoOutput.DelegateQueue]: The dispatch queue on which to call delegate methods.
//
// # Notifying the delegate of changes
//
//   - [IAVPlayerItemVideoOutput.RequestNotificationOfMediaDataChangeWithAdvanceInterval]: Tells the receiver that the video out put client is entering a quiescent state.
//
// # Getting pixel buffer data
//
//   - [IAVPlayerItemVideoOutput.HasNewPixelBufferForItemTime]: Returns a Boolean value that indicates whether video output is available for the specified item time.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemVideoOutput
type IAVPlayerItemVideoOutput interface {
	IAVPlayerItemOutput

	// Topic: Creating a video output

	// Creates a video output object initialized with the specified output settings.
	InitWithOutputSettings(outputSettings foundation.INSDictionary) AVPlayerItemVideoOutput

	// Topic: Configuring the delegate

	// Sets the delegate and dispatch queue for the receiver.
	SetDelegateQueue(delegate AVPlayerItemOutputPullDelegate, delegateQueue dispatch.Queue)
	// The delegate for the video output object.
	Delegate() AVPlayerItemOutputPullDelegate
	// The dispatch queue on which to call delegate methods.
	DelegateQueue() dispatch.Queue

	// Topic: Notifying the delegate of changes

	// Tells the receiver that the video out put client is entering a quiescent state.
	RequestNotificationOfMediaDataChangeWithAdvanceInterval(interval float64)

	// Topic: Getting pixel buffer data

	// Returns a Boolean value that indicates whether video output is available for the specified item time.
	HasNewPixelBufferForItemTime(itemTime coremedia.CMTime) bool
}

// Init initializes the instance.
func (p AVPlayerItemVideoOutput) Init() AVPlayerItemVideoOutput {
	rv := objc.Send[AVPlayerItemVideoOutput](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p AVPlayerItemVideoOutput) Autorelease() AVPlayerItemVideoOutput {
	rv := objc.Send[AVPlayerItemVideoOutput](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPlayerItemVideoOutput creates a new AVPlayerItemVideoOutput instance.
func NewAVPlayerItemVideoOutput() AVPlayerItemVideoOutput {
	class := getAVPlayerItemVideoOutputClass()
	rv := objc.Send[AVPlayerItemVideoOutput](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a video output object initialized with the specified output
// settings.
//
// outputSettings: The client requirements for output [CVPixelBuffer] objects, expressed using
// the constants in `AVVideoSettings.H()`.
// //
// [CVPixelBuffer]: https://developer.apple.com/documentation/CoreVideo/cvpixelbuffer-q2e
//
// # Discussion
// 
// For uncompressed video output, start with `kCVPixelBuffer*` keys in ``. In
// addition to the keys in `CVPixelBuffer.H()`, uncompressed video settings
// dictionaries may also provide a value for [AVVideoAllowWideColorKey].
//
// [AVVideoAllowWideColorKey]: https://developer.apple.com/documentation/AVFoundation/AVVideoAllowWideColorKey
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemVideoOutput/init(outputSettings:)
func NewPlayerItemVideoOutputWithOutputSettings(outputSettings foundation.INSDictionary) AVPlayerItemVideoOutput {
	instance := getAVPlayerItemVideoOutputClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithOutputSettings:"), outputSettings)
	return AVPlayerItemVideoOutputFromID(rv)
}

// Creates a video output object initialized with the specified output
// settings.
//
// outputSettings: The client requirements for output [CVPixelBuffer] objects, expressed using
// the constants in `AVVideoSettings.H()`.
// //
// [CVPixelBuffer]: https://developer.apple.com/documentation/CoreVideo/cvpixelbuffer-q2e
//
// # Discussion
// 
// For uncompressed video output, start with `kCVPixelBuffer*` keys in ``. In
// addition to the keys in `CVPixelBuffer.H()`, uncompressed video settings
// dictionaries may also provide a value for [AVVideoAllowWideColorKey].
//
// [AVVideoAllowWideColorKey]: https://developer.apple.com/documentation/AVFoundation/AVVideoAllowWideColorKey
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemVideoOutput/init(outputSettings:)
func (p AVPlayerItemVideoOutput) InitWithOutputSettings(outputSettings foundation.INSDictionary) AVPlayerItemVideoOutput {
	rv := objc.Send[AVPlayerItemVideoOutput](p.ID, objc.Sel("initWithOutputSettings:"), outputSettings)
	return rv
}
// Sets the delegate and dispatch queue for the receiver.
//
// delegate: The delegate object for the receiver. You may specify `nil` for this
// parameter.
//
// delegateQueue: The dispatch queue on which to call delegate methods. If you specify `nil`
// for this parameter, the video output object calls the delegate on the
// dispatch queue for your app’s main thread.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemVideoOutput/setDelegate(_:queue:)
func (p AVPlayerItemVideoOutput) SetDelegateQueue(delegate AVPlayerItemOutputPullDelegate, delegateQueue dispatch.Queue) {
	objc.Send[objc.ID](p.ID, objc.Sel("setDelegate:queue:"), delegate, uintptr(delegateQueue.Handle()))
}
// Tells the receiver that the video out put client is entering a quiescent
// state.
//
// interval: The amount of time to wait before notifying the delegate of the media
// change.
//
// # Discussion
// 
// Call this method before you suspend your use of a [CVDisplayLink] type or a
// [CADisplayLink] object. After the interval expires, the video output object
// notifies its delegate that it should resume the display link. If the
// interval value you specify is large, the delegate is notified as soon as
// possible rather than waiting.
// 
// Do not call this method repeatedly to force the delegate to be notified for
// each sample.
//
// [CADisplayLink]: https://developer.apple.com/documentation/QuartzCore/CADisplayLink
// [CVDisplayLink]: https://developer.apple.com/documentation/CoreVideo/CVDisplayLink
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemVideoOutput/requestNotificationOfMediaDataChange(withAdvanceInterval:)
func (p AVPlayerItemVideoOutput) RequestNotificationOfMediaDataChangeWithAdvanceInterval(interval float64) {
	objc.Send[objc.ID](p.ID, objc.Sel("requestNotificationOfMediaDataChangeWithAdvanceInterval:"), interval)
}
// Returns a Boolean value that indicates whether video output is available
// for the specified item time.
//
// itemTime: The item time to query. The time value is relative to the [AVPlayerItem]
// object with which the receiver is associated.
//
// # Return Value
// 
// [true] if there is available video output that has not been previously
// acquired or [false] if there is not.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method returns [true] if the video data at the specified time has not
// yet been acquired or is different from the video that was acquired
// previously. If you require multiple objects to acquire video output from
// the same [AVPlayerItem] object, you should create separate
// [AVPlayerItemVideoOutput] objects for each.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemVideoOutput/hasNewPixelBuffer(forItemTime:)
func (p AVPlayerItemVideoOutput) HasNewPixelBufferForItemTime(itemTime coremedia.CMTime) bool {
	rv := objc.Send[bool](p.ID, objc.Sel("hasNewPixelBufferForItemTime:"), itemTime)
	return rv
}

// The delegate for the video output object.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemVideoOutput/delegate
func (p AVPlayerItemVideoOutput) Delegate() AVPlayerItemOutputPullDelegate {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("delegate"))
	return AVPlayerItemOutputPullDelegateObjectFromID(rv)
}
// The dispatch queue on which to call delegate methods.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemVideoOutput/delegateQueue
func (p AVPlayerItemVideoOutput) DelegateQueue() dispatch.Queue {
	rv := objc.Send[uintptr](p.ID, objc.Sel("delegateQueue"))
	return dispatch.QueueFromHandle(rv)
}

