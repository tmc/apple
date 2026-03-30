// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AVPlayerItemRenderedLegibleOutput] class.
var (
	_AVPlayerItemRenderedLegibleOutputClass     AVPlayerItemRenderedLegibleOutputClass
	_AVPlayerItemRenderedLegibleOutputClassOnce sync.Once
)

func getAVPlayerItemRenderedLegibleOutputClass() AVPlayerItemRenderedLegibleOutputClass {
	_AVPlayerItemRenderedLegibleOutputClassOnce.Do(func() {
		_AVPlayerItemRenderedLegibleOutputClass = AVPlayerItemRenderedLegibleOutputClass{class: objc.GetClass("AVPlayerItemRenderedLegibleOutput")}
	})
	return _AVPlayerItemRenderedLegibleOutputClass
}

// GetAVPlayerItemRenderedLegibleOutputClass returns the class object for AVPlayerItemRenderedLegibleOutput.
func GetAVPlayerItemRenderedLegibleOutputClass() AVPlayerItemRenderedLegibleOutputClass {
	return getAVPlayerItemRenderedLegibleOutputClass()
}

type AVPlayerItemRenderedLegibleOutputClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVPlayerItemRenderedLegibleOutputClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVPlayerItemRenderedLegibleOutputClass) Alloc() AVPlayerItemRenderedLegibleOutput {
	rv := objc.Send[AVPlayerItemRenderedLegibleOutput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A player item output that vends media with a legible characteristic as
// rendered pixel buffers.
//
// # Creating an output
//
//   - [AVPlayerItemRenderedLegibleOutput.InitWithVideoDisplaySize]: Creates a rendered legible output object.
//
// # Configuring an output
//
//   - [AVPlayerItemRenderedLegibleOutput.AdvanceIntervalForDelegateInvocation]: Permits advance invocation of the associated delegate, if any.
//   - [AVPlayerItemRenderedLegibleOutput.SetAdvanceIntervalForDelegateInvocation]
//   - [AVPlayerItemRenderedLegibleOutput.VideoDisplaySize]: Set the video display size to use for rendering of pixel buffers.
//   - [AVPlayerItemRenderedLegibleOutput.SetVideoDisplaySize]
//
// # Setting a delegate
//
//   - [AVPlayerItemRenderedLegibleOutput.Delegate]: A delegate object for this output.
//   - [AVPlayerItemRenderedLegibleOutput.SetDelegateQueue]: Sets the delegate object and the queue on which it’s invoked.
//   - [AVPlayerItemRenderedLegibleOutput.DelegateQueue]: The dispatch queue on which the output calls the delegate object.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemRenderedLegibleOutput
type AVPlayerItemRenderedLegibleOutput struct {
	AVPlayerItemOutput
}

// AVPlayerItemRenderedLegibleOutputFromID constructs a [AVPlayerItemRenderedLegibleOutput] from an objc.ID.
//
// A player item output that vends media with a legible characteristic as
// rendered pixel buffers.
func AVPlayerItemRenderedLegibleOutputFromID(id objc.ID) AVPlayerItemRenderedLegibleOutput {
	return AVPlayerItemRenderedLegibleOutput{AVPlayerItemOutput: AVPlayerItemOutputFromID(id)}
}

// NOTE: AVPlayerItemRenderedLegibleOutput adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVPlayerItemRenderedLegibleOutput] class.
//
// # Creating an output
//
//   - [IAVPlayerItemRenderedLegibleOutput.InitWithVideoDisplaySize]: Creates a rendered legible output object.
//
// # Configuring an output
//
//   - [IAVPlayerItemRenderedLegibleOutput.AdvanceIntervalForDelegateInvocation]: Permits advance invocation of the associated delegate, if any.
//   - [IAVPlayerItemRenderedLegibleOutput.SetAdvanceIntervalForDelegateInvocation]
//   - [IAVPlayerItemRenderedLegibleOutput.VideoDisplaySize]: Set the video display size to use for rendering of pixel buffers.
//   - [IAVPlayerItemRenderedLegibleOutput.SetVideoDisplaySize]
//
// # Setting a delegate
//
//   - [IAVPlayerItemRenderedLegibleOutput.Delegate]: A delegate object for this output.
//   - [IAVPlayerItemRenderedLegibleOutput.SetDelegateQueue]: Sets the delegate object and the queue on which it’s invoked.
//   - [IAVPlayerItemRenderedLegibleOutput.DelegateQueue]: The dispatch queue on which the output calls the delegate object.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemRenderedLegibleOutput
type IAVPlayerItemRenderedLegibleOutput interface {
	IAVPlayerItemOutput

	// Topic: Creating an output

	// Creates a rendered legible output object.
	InitWithVideoDisplaySize(videoDisplaySize corefoundation.CGSize) AVPlayerItemRenderedLegibleOutput

	// Topic: Configuring an output

	// Permits advance invocation of the associated delegate, if any.
	AdvanceIntervalForDelegateInvocation() float64
	SetAdvanceIntervalForDelegateInvocation(value float64)
	// Set the video display size to use for rendering of pixel buffers.
	VideoDisplaySize() corefoundation.CGSize
	SetVideoDisplaySize(value corefoundation.CGSize)

	// Topic: Setting a delegate

	// A delegate object for this output.
	Delegate() AVPlayerItemRenderedLegibleOutputPushDelegate
	// Sets the delegate object and the queue on which it’s invoked.
	SetDelegateQueue(delegate AVPlayerItemRenderedLegibleOutputPushDelegate, delegateQueue dispatch.Queue)
	// The dispatch queue on which the output calls the delegate object.
	DelegateQueue() dispatch.Queue
}

// Init initializes the instance.
func (p AVPlayerItemRenderedLegibleOutput) Init() AVPlayerItemRenderedLegibleOutput {
	rv := objc.Send[AVPlayerItemRenderedLegibleOutput](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p AVPlayerItemRenderedLegibleOutput) Autorelease() AVPlayerItemRenderedLegibleOutput {
	rv := objc.Send[AVPlayerItemRenderedLegibleOutput](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPlayerItemRenderedLegibleOutput creates a new AVPlayerItemRenderedLegibleOutput instance.
func NewAVPlayerItemRenderedLegibleOutput() AVPlayerItemRenderedLegibleOutput {
	class := getAVPlayerItemRenderedLegibleOutputClass()
	rv := objc.Send[AVPlayerItemRenderedLegibleOutput](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a rendered legible output object.
//
// videoDisplaySize: The size of the video display.
//
// # Discussion
//
// You can also choose to reset the [VideoDisplaySize] value after
// initialization or during playback.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemRenderedLegibleOutput/init(videoDisplay:)
func NewPlayerItemRenderedLegibleOutputWithVideoDisplaySize(videoDisplaySize corefoundation.CGSize) AVPlayerItemRenderedLegibleOutput {
	instance := getAVPlayerItemRenderedLegibleOutputClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithVideoDisplaySize:"), videoDisplaySize)
	return AVPlayerItemRenderedLegibleOutputFromID(rv)
}

// Creates a rendered legible output object.
//
// videoDisplaySize: The size of the video display.
//
// # Discussion
//
// You can also choose to reset the [VideoDisplaySize] value after
// initialization or during playback.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemRenderedLegibleOutput/init(videoDisplay:)
func (p AVPlayerItemRenderedLegibleOutput) InitWithVideoDisplaySize(videoDisplaySize corefoundation.CGSize) AVPlayerItemRenderedLegibleOutput {
	rv := objc.Send[AVPlayerItemRenderedLegibleOutput](p.ID, objc.Sel("initWithVideoDisplaySize:"), videoDisplaySize)
	return rv
}

// Sets the delegate object and the queue on which it’s invoked.
//
// delegate: A delegate object for this output.
//
// delegateQueue: A dispatch queue on which the system calls all delegate methods.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemRenderedLegibleOutput/setDelegate(_:queue:)
func (p AVPlayerItemRenderedLegibleOutput) SetDelegateQueue(delegate AVPlayerItemRenderedLegibleOutputPushDelegate, delegateQueue dispatch.Queue) {
	objc.Send[objc.ID](p.ID, objc.Sel("setDelegate:queue:"), delegate, uintptr(delegateQueue.Handle()))
}

// Permits advance invocation of the associated delegate, if any.
//
// # Discussion
//
// Use this property specify the number of seconds early to invoke the
// delegate object. When possible, an AVPlayerItemLegibleOutput uses this
// value to call its delegate earlier than it would otherwise.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemRenderedLegibleOutput/advanceIntervalForDelegateInvocation
func (p AVPlayerItemRenderedLegibleOutput) AdvanceIntervalForDelegateInvocation() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("advanceIntervalForDelegateInvocation"))
	return rv
}
func (p AVPlayerItemRenderedLegibleOutput) SetAdvanceIntervalForDelegateInvocation(value float64) {
	objc.Send[struct{}](p.ID, objc.Sel("setAdvanceIntervalForDelegateInvocation:"), value)
}

// Set the video display size to use for rendering of pixel buffers.
//
// # Discussion
//
// The output renders the pixel buffers according to the width and height of
// display area. If you set this property during the presentation time of a
// vended caption image, the output vends a new image rendered at the new
// size.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemRenderedLegibleOutput/videoDisplaySize
func (p AVPlayerItemRenderedLegibleOutput) VideoDisplaySize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](p.ID, objc.Sel("videoDisplaySize"))
	return corefoundation.CGSize(rv)
}
func (p AVPlayerItemRenderedLegibleOutput) SetVideoDisplaySize(value corefoundation.CGSize) {
	objc.Send[struct{}](p.ID, objc.Sel("setVideoDisplaySize:"), value)
}

// A delegate object for this output.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemRenderedLegibleOutput/delegate
func (p AVPlayerItemRenderedLegibleOutput) Delegate() AVPlayerItemRenderedLegibleOutputPushDelegate {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("delegate"))
	return AVPlayerItemRenderedLegibleOutputPushDelegateObjectFromID(rv)
}

// The dispatch queue on which the output calls the delegate object.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemRenderedLegibleOutput/delegateQueue
func (p AVPlayerItemRenderedLegibleOutput) DelegateQueue() dispatch.Queue {
	rv := objc.Send[uintptr](p.ID, objc.Sel("delegateQueue"))
	return dispatch.QueueFromHandle(rv)
}
