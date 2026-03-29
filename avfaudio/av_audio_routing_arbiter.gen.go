// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAudioRoutingArbiter] class.
var (
	_AVAudioRoutingArbiterClass     AVAudioRoutingArbiterClass
	_AVAudioRoutingArbiterClassOnce sync.Once
)

func getAVAudioRoutingArbiterClass() AVAudioRoutingArbiterClass {
	_AVAudioRoutingArbiterClassOnce.Do(func() {
		_AVAudioRoutingArbiterClass = AVAudioRoutingArbiterClass{class: objc.GetClass("AVAudioRoutingArbiter")}
	})
	return _AVAudioRoutingArbiterClass
}

// GetAVAudioRoutingArbiterClass returns the class object for AVAudioRoutingArbiter.
func GetAVAudioRoutingArbiterClass() AVAudioRoutingArbiterClass {
	return getAVAudioRoutingArbiterClass()
}

type AVAudioRoutingArbiterClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAudioRoutingArbiterClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAudioRoutingArbiterClass) Alloc() AVAudioRoutingArbiter {
	rv := objc.Send[AVAudioRoutingArbiter](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object for configuring macOS apps to participate in AirPods Automatic
// Switching.
//
// # Overview
// 
// AirPods Automatic Switching is a feature of Apple operating systems that
// intelligently connects wireless headphones to the most appropriate audio
// device in a multidevice environment. For example, if a user plays a movie
// on iPad, and then locks the device and starts playing music on iPhone, the
// system automatically switches the source audio device from iPad to iPhone.
// 
// iOS apps automatically participate in AirPods Automatic Switching. To
// enable your macOS app to participate in this behavior, use
// [AVAudioRoutingArbiter] to indicate when your app starts and finishes
// playing or recording audio. For example, a Voice over IP (VoIP) app might
// request arbitration before starting a call, and when the arbitration
// completes, begin the VoIP session. Likewise, when the call ends, the app
// would end the VoIP session and leave arbitration.
//
// # Participating in AirPods Automatic Switching
//
//   - [AVAudioRoutingArbiter.BeginArbitrationWithCategoryCompletionHandler]: Begins routing arbitration to take ownership of a nearby Bluetooth audio route.
//   - [AVAudioRoutingArbiter.LeaveArbitration]: Stops an app’s participation in audio routing arbitration.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRoutingArbiter
type AVAudioRoutingArbiter struct {
	objectivec.Object
}

// AVAudioRoutingArbiterFromID constructs a [AVAudioRoutingArbiter] from an objc.ID.
//
// An object for configuring macOS apps to participate in AirPods Automatic
// Switching.
func AVAudioRoutingArbiterFromID(id objc.ID) AVAudioRoutingArbiter {
	return AVAudioRoutingArbiter{objectivec.Object{ID: id}}
}
// NOTE: AVAudioRoutingArbiter adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAudioRoutingArbiter] class.
//
// # Participating in AirPods Automatic Switching
//
//   - [IAVAudioRoutingArbiter.BeginArbitrationWithCategoryCompletionHandler]: Begins routing arbitration to take ownership of a nearby Bluetooth audio route.
//   - [IAVAudioRoutingArbiter.LeaveArbitration]: Stops an app’s participation in audio routing arbitration.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRoutingArbiter
type IAVAudioRoutingArbiter interface {
	objectivec.IObject

	// Topic: Participating in AirPods Automatic Switching

	// Begins routing arbitration to take ownership of a nearby Bluetooth audio route.
	BeginArbitrationWithCategoryCompletionHandler(category AVAudioRoutingArbitrationCategory, handler BoolErrorHandler)
	// Stops an app’s participation in audio routing arbitration.
	LeaveArbitration()
}

// Init initializes the instance.
func (a AVAudioRoutingArbiter) Init() AVAudioRoutingArbiter {
	rv := objc.Send[AVAudioRoutingArbiter](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAudioRoutingArbiter) Autorelease() AVAudioRoutingArbiter {
	rv := objc.Send[AVAudioRoutingArbiter](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAudioRoutingArbiter creates a new AVAudioRoutingArbiter instance.
func NewAVAudioRoutingArbiter() AVAudioRoutingArbiter {
	class := getAVAudioRoutingArbiterClass()
	rv := objc.Send[AVAudioRoutingArbiter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Begins routing arbitration to take ownership of a nearby Bluetooth audio
// route.
//
// category: A category that describes how the app uses audio.
//
// handler: A completion handler the system calls asynchronously when the system
// completes audio routing arbitration. This closure takes the following
// parameters:
// 
// defaultDeviceChanged: A Boolean value that indicates whether the system
// switched the AirPods to the macOS device. error: An error object that
// indicates why the request failed, or [nil] if the request succeeded.
// //
// [nil]: https://developer.apple.com/documentation/ObjectiveC/nil-227m0
//
// # Discussion
// 
// Call this method to tell the operating system to arbitrate with nearby
// Apple devices to take ownership of a supported Bluetooth audio device. When
// arbitration completes, the system calls the completion handler, passing a
// Boolean that indicates whether the audio device changed. In either case,
// begin using audio as normal.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRoutingArbiter/begin(category:completionHandler:)
func (a AVAudioRoutingArbiter) BeginArbitrationWithCategoryCompletionHandler(category AVAudioRoutingArbitrationCategory, handler BoolErrorHandler) {
_block1, _ := NewBoolErrorBlock(handler)
	objc.Send[objc.ID](a.ID, objc.Sel("beginArbitrationWithCategory:completionHandler:"), category, _block1)
}
// Stops an app’s participation in audio routing arbitration.
//
// # Discussion
// 
// Configure your app to notify the system when the app stops using audio for
// an undetermined duration. For example, for a Voice over IP (VoIP) app, call
// this method when the VoIP call ends. Calling this method allows the system
// to make an informed decision when multiple Apple devices are trying to take
// ownership of a Bluetooth audio route.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRoutingArbiter/leave()
func (a AVAudioRoutingArbiter) LeaveArbitration() {
	objc.Send[objc.ID](a.ID, objc.Sel("leaveArbitration"))
}

// The shared routing arbiter object.
//
// See: https://developer.apple.com/documentation/AVFAudio/AVAudioRoutingArbiter/shared
func (_AVAudioRoutingArbiterClass AVAudioRoutingArbiterClass) SharedRoutingArbiter() AVAudioRoutingArbiter {
	rv := objc.Send[objc.ID](objc.ID(_AVAudioRoutingArbiterClass.class), objc.Sel("sharedRoutingArbiter"))
	return AVAudioRoutingArbiterFromID(objc.ID(rv))
}

// BeginArbitrationWithCategory is a synchronous wrapper around [AVAudioRoutingArbiter.BeginArbitrationWithCategoryCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (a AVAudioRoutingArbiter) BeginArbitrationWithCategory(ctx context.Context, category AVAudioRoutingArbitrationCategory) (bool, error) {
	type result struct {
		val bool
		err error
	}
	done := make(chan result, 1)
	a.BeginArbitrationWithCategoryCompletionHandler(category, func(val bool, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return false, ctx.Err()
	}
}

