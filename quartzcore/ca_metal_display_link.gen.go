// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CAMetalDisplayLink] class.
var (
	_CAMetalDisplayLinkClass     CAMetalDisplayLinkClass
	_CAMetalDisplayLinkClassOnce sync.Once
)

func getCAMetalDisplayLinkClass() CAMetalDisplayLinkClass {
	_CAMetalDisplayLinkClassOnce.Do(func() {
		_CAMetalDisplayLinkClass = CAMetalDisplayLinkClass{class: objc.GetClass("CAMetalDisplayLink")}
	})
	return _CAMetalDisplayLinkClass
}

// GetCAMetalDisplayLinkClass returns the class object for CAMetalDisplayLink.
func GetCAMetalDisplayLinkClass() CAMetalDisplayLinkClass {
	return getCAMetalDisplayLinkClass()
}

type CAMetalDisplayLinkClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CAMetalDisplayLinkClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CAMetalDisplayLinkClass) Alloc() CAMetalDisplayLink {
	rv := objc.Send[CAMetalDisplayLink](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A class your Metal app uses to register for callbacks to synchronize its
// animations for a display.
//
// # Overview
//
// [CAMetalDisplayLink] instances are a specialized way to interact with
// variable-rate displays when you need more control over the timing window to
// render your app’s frames. Controlling the timing window and rendering
// delay for frames can help you achieve smoother frame rates and avoid visual
// artifacts.
//
// Your app initializes a new Metal display link by providing a target
// [CAMetalLayer]. Set this instance’s [CAMetalDisplayLink.Delegate] property to an
// implementation that encodes the rendering work for Metal to perform. With a
// set delegate, synchronize the display with a run loop to perform rendering
// on by calling the [CAMetalDisplayLink.AddToRunLoopForMode] method.
//
// Once you associate the display link with a run loop, the system calls the
// delegate’s [MetalDisplayLinkNeedsUpdate] method to request new frames.
// This method receives update requests based on the [CAMetalDisplayLink.PreferredFrameRateRange]
// and [CAMetalDisplayLink.PreferredFrameLatency] of the display link. The system makes a best
// effort to make callbacks at appropriate times. Your app should complete any
// commits to the Metal device’s [MTLCommandQueue] for rendering the display
// layer before calling [present()] on a drawable element.
//
// Your app can disable notifications by setting [CAMetalDisplayLink.Paused] to `true`. When your
// app finishes with a display link, call [CAMetalDisplayLink.Invalidate]to remove it from all
// run loops and the target.
//
// # Creating a Display Link
//
//   - [CAMetalDisplayLink.InitWithMetalLayer]: Creates a display link for Metal from a Core Animation layer.
//
// # Configuring a Display Link
//
//   - [CAMetalDisplayLink.PreferredFrameRateRange]: A range of frequencies your app allows for frame updates, affecting how often the system invokes your delegate’s callback.
//   - [CAMetalDisplayLink.SetPreferredFrameRateRange]
//   - [CAMetalDisplayLink.PreferredFrameLatency]: The amount of time, in frames, your app requests to render a frame.
//   - [CAMetalDisplayLink.SetPreferredFrameLatency]
//   - [CAMetalDisplayLink.Delegate]: An instance of a type your app implements that responds to the system’s callbacks.
//   - [CAMetalDisplayLink.SetDelegate]
//
// # Registering for Callbacks
//
//   - [CAMetalDisplayLink.AddToRunLoopForMode]: Registers the display link with a run loop.
//
// # Pausing Callbacks
//
//   - [CAMetalDisplayLink.Paused]: A Boolean value that indicates whether the system suspends the display link’s notifications to the target.
//   - [CAMetalDisplayLink.SetPaused]
//
// # Deregistering for callbacks
//
//   - [CAMetalDisplayLink.RemoveFromRunLoopForMode]: Removes a mode’s display link from a run loop.
//   - [CAMetalDisplayLink.Invalidate]: Removes the display link from all run loops for all modes.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMetalDisplayLink
//
// [MTLCommandQueue]: https://developer.apple.com/documentation/Metal/MTLCommandQueue
// [present()]: https://developer.apple.com/documentation/Metal/MTLDrawable/present()
type CAMetalDisplayLink struct {
	objectivec.Object
}

// CAMetalDisplayLinkFromID constructs a [CAMetalDisplayLink] from an objc.ID.
//
// A class your Metal app uses to register for callbacks to synchronize its
// animations for a display.
func CAMetalDisplayLinkFromID(id objc.ID) CAMetalDisplayLink {
	return CAMetalDisplayLink{objectivec.Object{ID: id}}
}

// NOTE: CAMetalDisplayLink adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CAMetalDisplayLink] class.
//
// # Creating a Display Link
//
//   - [ICAMetalDisplayLink.InitWithMetalLayer]: Creates a display link for Metal from a Core Animation layer.
//
// # Configuring a Display Link
//
//   - [ICAMetalDisplayLink.PreferredFrameRateRange]: A range of frequencies your app allows for frame updates, affecting how often the system invokes your delegate’s callback.
//   - [ICAMetalDisplayLink.SetPreferredFrameRateRange]
//   - [ICAMetalDisplayLink.PreferredFrameLatency]: The amount of time, in frames, your app requests to render a frame.
//   - [ICAMetalDisplayLink.SetPreferredFrameLatency]
//   - [ICAMetalDisplayLink.Delegate]: An instance of a type your app implements that responds to the system’s callbacks.
//   - [ICAMetalDisplayLink.SetDelegate]
//
// # Registering for Callbacks
//
//   - [ICAMetalDisplayLink.AddToRunLoopForMode]: Registers the display link with a run loop.
//
// # Pausing Callbacks
//
//   - [ICAMetalDisplayLink.Paused]: A Boolean value that indicates whether the system suspends the display link’s notifications to the target.
//   - [ICAMetalDisplayLink.SetPaused]
//
// # Deregistering for callbacks
//
//   - [ICAMetalDisplayLink.RemoveFromRunLoopForMode]: Removes a mode’s display link from a run loop.
//   - [ICAMetalDisplayLink.Invalidate]: Removes the display link from all run loops for all modes.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMetalDisplayLink
type ICAMetalDisplayLink interface {
	objectivec.IObject

	// Topic: Creating a Display Link

	// Creates a display link for Metal from a Core Animation layer.
	InitWithMetalLayer(layer ICAMetalLayer) CAMetalDisplayLink

	// Topic: Configuring a Display Link

	// A range of frequencies your app allows for frame updates, affecting how often the system invokes your delegate’s callback.
	PreferredFrameRateRange() CAFrameRateRange
	SetPreferredFrameRateRange(value CAFrameRateRange)
	// The amount of time, in frames, your app requests to render a frame.
	PreferredFrameLatency() float32
	SetPreferredFrameLatency(value float32)
	// An instance of a type your app implements that responds to the system’s callbacks.
	Delegate() CAMetalDisplayLinkDelegate
	SetDelegate(value CAMetalDisplayLinkDelegate)

	// Topic: Registering for Callbacks

	// Registers the display link with a run loop.
	AddToRunLoopForMode(runloop foundation.NSRunLoop, mode foundation.NSString)

	// Topic: Pausing Callbacks

	// A Boolean value that indicates whether the system suspends the display link’s notifications to the target.
	Paused() bool
	SetPaused(value bool)

	// Topic: Deregistering for callbacks

	// Removes a mode’s display link from a run loop.
	RemoveFromRunLoopForMode(runloop foundation.NSRunLoop, mode foundation.NSString)
	// Removes the display link from all run loops for all modes.
	Invalidate()
}

// Init initializes the instance.
func (m CAMetalDisplayLink) Init() CAMetalDisplayLink {
	rv := objc.Send[CAMetalDisplayLink](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m CAMetalDisplayLink) Autorelease() CAMetalDisplayLink {
	rv := objc.Send[CAMetalDisplayLink](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewCAMetalDisplayLink creates a new CAMetalDisplayLink instance.
func NewCAMetalDisplayLink() CAMetalDisplayLink {
	class := getCAMetalDisplayLinkClass()
	rv := objc.Send[CAMetalDisplayLink](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a display link for Metal from a Core Animation layer.
//
// layer: A Core Animation layer for Metal.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMetalDisplayLink/init(metalLayer:)
func NewMetalDisplayLinkWithMetalLayer(layer ICAMetalLayer) CAMetalDisplayLink {
	instance := getCAMetalDisplayLinkClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMetalLayer:"), layer)
	return CAMetalDisplayLinkFromID(rv)
}

// Creates a display link for Metal from a Core Animation layer.
//
// layer: A Core Animation layer for Metal.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMetalDisplayLink/init(metalLayer:)
func (m CAMetalDisplayLink) InitWithMetalLayer(layer ICAMetalLayer) CAMetalDisplayLink {
	rv := objc.Send[CAMetalDisplayLink](m.ID, objc.Sel("initWithMetalLayer:"), layer)
	return rv
}

// Registers the display link with a run loop.
//
// runloop: A run loop instance the method associates with the display link.
//
// mode: A run loop mode for the display link.
//
// # Discussion
//
// You can associate the display link with any of the [RunLoop] modes,
// multiple input modes, or a custom mode. When the run loop is in `mode`, the
// display link notifies its delegate when the system prepares the next frame.
//
// You can remove the display link from a run loop by calling
// [RemoveFromRunLoopForMode], or from all run loops with [Invalidate].
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMetalDisplayLink/add(to:forMode:)
//
// [RunLoop]: https://developer.apple.com/documentation/Foundation/RunLoop
func (m CAMetalDisplayLink) AddToRunLoopForMode(runloop foundation.NSRunLoop, mode foundation.NSString) {
	objc.Send[objc.ID](m.ID, objc.Sel("addToRunLoop:forMode:"), runloop, mode)
}

// Removes a mode’s display link from a run loop.
//
// runloop: A run loop the method disassociates the display link from for `mode`.
//
// mode: A run loop mode the method disassociates the display link for `runloop`.
//
// # Discussion
//
// The run loop releases the display link if it no longer associates with any
// run modes.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMetalDisplayLink/remove(from:forMode:)
func (m CAMetalDisplayLink) RemoveFromRunLoopForMode(runloop foundation.NSRunLoop, mode foundation.NSString) {
	objc.Send[objc.ID](m.ID, objc.Sel("removeFromRunLoop:forMode:"), runloop, mode)
}

// Removes the display link from all run loops for all modes.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMetalDisplayLink/invalidate()
func (m CAMetalDisplayLink) Invalidate() {
	objc.Send[objc.ID](m.ID, objc.Sel("invalidate"))
}

// A range of frequencies your app allows for frame updates, affecting how
// often the system invokes your delegate’s callback.
//
// # Discussion
//
// The display link makes a best attempt to invoke your app’s callback
// within the frequency range you set to this property. However, the system
// also takes into account the device’s hardware capabilities and the other
// tasks your game or app is running.
//
// The system can change the available range of frame rates because it factors
// in system policies and a person’s preferences. For example, Low Power
// Mode, critical thermal state, and accessibility settings can affect the
// system’s frame rate.
//
// The system typically provides a consistent frame rate by choosing one
// that’s a factor of the display’s maximum refresh rate. For example, a
// display link could invoke your callback 60 times per second for a display
// with a refresh rate of 60 hertz. However, the display link could invoke
// your callback less frequently, such as 30, 20, or 15 hertz, by setting a
// range with smaller values.
//
// See [Optimizing ProMotion refresh rates for iPhone 13 Pro and iPad Pro] for
// more information.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMetalDisplayLink/preferredFrameRateRange
//
// [Optimizing ProMotion refresh rates for iPhone 13 Pro and iPad Pro]: https://developer.apple.com/documentation/QuartzCore/optimizing-promotion-refresh-rates-for-iphone-13-pro-and-ipad-pro
func (m CAMetalDisplayLink) PreferredFrameRateRange() CAFrameRateRange {
	rv := objc.Send[CAFrameRateRange](m.ID, objc.Sel("preferredFrameRateRange"))
	return CAFrameRateRange(rv)
}
func (m CAMetalDisplayLink) SetPreferredFrameRateRange(value CAFrameRateRange) {
	objc.Send[struct{}](m.ID, objc.Sel("setPreferredFrameRateRange:"), value)
}

// The amount of time, in frames, your app requests to render a frame.
//
// # Discussion
//
// The final latency may be bigger if the system needs more time, such as for
// windowed modes on macOS.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMetalDisplayLink/preferredFrameLatency
func (m CAMetalDisplayLink) PreferredFrameLatency() float32 {
	rv := objc.Send[float32](m.ID, objc.Sel("preferredFrameLatency"))
	return rv
}
func (m CAMetalDisplayLink) SetPreferredFrameLatency(value float32) {
	objc.Send[struct{}](m.ID, objc.Sel("setPreferredFrameLatency:"), value)
}

// An instance of a type your app implements that responds to the system’s
// callbacks.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMetalDisplayLink/delegate
func (m CAMetalDisplayLink) Delegate() CAMetalDisplayLinkDelegate {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("delegate"))
	return CAMetalDisplayLinkDelegateObjectFromID(rv)
}
func (m CAMetalDisplayLink) SetDelegate(value CAMetalDisplayLinkDelegate) {
	objc.Send[struct{}](m.ID, objc.Sel("setDelegate:"), value)
}

// A Boolean value that indicates whether the system suspends the display
// link’s notifications to the target.
//
// # Discussion
//
// You can instruct the display link to stop sending notifications to the
// delegate by setting the property to true. The property defaults to false.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMetalDisplayLink/isPaused
func (m CAMetalDisplayLink) Paused() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isPaused"))
	return rv
}
func (m CAMetalDisplayLink) SetPaused(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setPaused:"), value)
}
