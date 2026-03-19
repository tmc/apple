// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CADisplayLink] class.
var (
	_CADisplayLinkClass     CADisplayLinkClass
	_CADisplayLinkClassOnce sync.Once
)

func getCADisplayLinkClass() CADisplayLinkClass {
	_CADisplayLinkClassOnce.Do(func() {
		_CADisplayLinkClass = CADisplayLinkClass{class: objc.GetClass("CADisplayLink")}
	})
	return _CADisplayLinkClass
}

// GetCADisplayLinkClass returns the class object for CADisplayLink.
func GetCADisplayLinkClass() CADisplayLinkClass {
	return getCADisplayLinkClass()
}

type CADisplayLinkClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (cc CADisplayLinkClass) Alloc() CADisplayLink {
	rv := objc.Send[CADisplayLink](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A timer object that allows your app to synchronize its drawing to the
// refresh rate of the display.
//
// # Overview
// 
// Your app initializes a new display link by providing a target object and a
// selector to call when the system updates the screen. To synchronize your
// display loop with the display, your application adds it to a run loop using
// the [CADisplayLink.AddToRunLoopForMode] method.
// 
// Once you associate the display link with a run loop, the system calls the
// selector on the target when the screen’s contents need to update. The
// target can read the display link’s [CADisplayLink.Timestamp] property to retrieve the
// time the system displayed the previous frame. For example, an app that
// displays movies might use `timestamp` to calculate which video frame to
// display next. An app that performs its own animations might use `timestamp`
// to determine where and how visible objects appear in the upcoming frame.
// 
// The [CADisplayLink.Duration] property provides the amount of time between frames at the
// [CADisplayLink.MaximumFramesPerSecond]. To calculate the actual frame duration, use
// [CADisplayLink.TargetTimestamp] - [CADisplayLink.Timestamp]. You can use this value in your app to
// calculate the frame rate of the display, the approximate time the system
// displays the next frame, and to adjust the drawing behavior so that the
// next frame is ready in time to display.
// 
// Your app can disable notifications by setting [CADisplayLink.Paused] to `true`. Also, if
// your app can’t provide frames in the time the system provides, you may
// want to choose a slower frame rate. An app with a slower but consistent
// frame rate appears smoother to the user than an app that skips frames. You
// can define the number of frames per second by setting
// [CADisplayLink.PreferredFramesPerSecond].
// 
// When your app finishes with a display link, call [CADisplayLink.Invalidate] to remove it
// from all run loops and to disassociate it from the target.
// 
// The code listing below shows how to create a display link and add it to the
// current run loop. The display link invokes the step function, which prints
// the target timestamp with each screen update.
// 
// You shouldn’t subclass [CADisplayLink].
// 
// # Preferred and Actual Frame Rates
// 
// You control a display link’s frame rate (the number of times the system
// calls the selector of its target, per second) by setting
// [CADisplayLink.PreferredFramesPerSecond]. However, the actual frames per second may
// differ from the preferred value you set; actual frame rates are always a
// factor of the maximum refresh rate of the device. For example, if your
// device’s maximum refresh rate is 60 frames per second (defined by
// [CADisplayLink.MaximumFramesPerSecond]), actual frame rates include 15, 20, 30, and 60
// frames per second. If you set a display link’s preferred frame rate to a
// value higher than the maximum, the actual frame rate is the maximum.
// 
// In iOS 15, frame rate availability can change due to the system factoring
// in the system policy and user preference — including Low Power Mode,
// critical thermal state, and accessibility settings.
// 
// The system rounds, to the nearest factor, preferred frame rates that
// aren’t a divisor of the maximum frame rate. For example, setting a
// preferred frame rate to either 26 or 35 frames per second on a device with
// a maximum refresh rate of 60 frames per second yields an actual frame rate
// of 30 times per second.
// 
// The code listing below shows how to calculate the actual frame rate by
// dividing 1 by your display link’s [CADisplayLink.Timestamp] subtracted from its
// [CADisplayLink.TargetTimestamp].
//
// # Configuring a Display Link
//
//   - [CADisplayLink.Duration]: The time interval between screen refresh updates.
//   - [CADisplayLink.PreferredFrameRateRange]: A range of frequencies your app allows for frame updates, affecting how often the system invokes your delegate’s callback.
//   - [CADisplayLink.SetPreferredFrameRateRange]
//   - [CADisplayLink.Paused]: A Boolean value that indicates whether the system suspends the display link’s notifications to the target.
//   - [CADisplayLink.SetPaused]
//   - [CADisplayLink.Timestamp]: The time interval that represents when the last frame displayed.
//   - [CADisplayLink.TargetTimestamp]: The time interval that represents when the next frame displays.
//
// # Scheduling a Display Link to Send Notifications
//
//   - [CADisplayLink.AddToRunLoopForMode]: Registers the display link with a run loop.
//   - [CADisplayLink.RemoveFromRunLoopForMode]: Removes the display link from the run loop for the given mode.
//   - [CADisplayLink.Invalidate]: Removes the display link from all run loop modes.
//
// See: https://developer.apple.com/documentation/QuartzCore/CADisplayLink
type CADisplayLink struct {
	objectivec.Object
}

// CADisplayLinkFromID constructs a [CADisplayLink] from an objc.ID.
//
// A timer object that allows your app to synchronize its drawing to the
// refresh rate of the display.
func CADisplayLinkFromID(id objc.ID) CADisplayLink {
	return CADisplayLink{objectivec.Object{ID: id}}
}
// NOTE: CADisplayLink adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CADisplayLink] class.
//
// # Configuring a Display Link
//
//   - [ICADisplayLink.Duration]: The time interval between screen refresh updates.
//   - [ICADisplayLink.PreferredFrameRateRange]: A range of frequencies your app allows for frame updates, affecting how often the system invokes your delegate’s callback.
//   - [ICADisplayLink.SetPreferredFrameRateRange]
//   - [ICADisplayLink.Paused]: A Boolean value that indicates whether the system suspends the display link’s notifications to the target.
//   - [ICADisplayLink.SetPaused]
//   - [ICADisplayLink.Timestamp]: The time interval that represents when the last frame displayed.
//   - [ICADisplayLink.TargetTimestamp]: The time interval that represents when the next frame displays.
//
// # Scheduling a Display Link to Send Notifications
//
//   - [ICADisplayLink.AddToRunLoopForMode]: Registers the display link with a run loop.
//   - [ICADisplayLink.RemoveFromRunLoopForMode]: Removes the display link from the run loop for the given mode.
//   - [ICADisplayLink.Invalidate]: Removes the display link from all run loop modes.
//
// See: https://developer.apple.com/documentation/QuartzCore/CADisplayLink
type ICADisplayLink interface {
	objectivec.IObject

	// Topic: Configuring a Display Link

	// The time interval between screen refresh updates.
	Duration() float64
	// A range of frequencies your app allows for frame updates, affecting how often the system invokes your delegate’s callback.
	PreferredFrameRateRange() CAFrameRateRange
	SetPreferredFrameRateRange(value CAFrameRateRange)
	// A Boolean value that indicates whether the system suspends the display link’s notifications to the target.
	Paused() bool
	SetPaused(value bool)
	// The time interval that represents when the last frame displayed.
	Timestamp() float64
	// The time interval that represents when the next frame displays.
	TargetTimestamp() float64

	// Topic: Scheduling a Display Link to Send Notifications

	// Registers the display link with a run loop.
	AddToRunLoopForMode(runloop foundation.NSRunLoop, mode foundation.NSString)
	// Removes the display link from the run loop for the given mode.
	RemoveFromRunLoopForMode(runloop foundation.NSRunLoop, mode foundation.NSString)
	// Removes the display link from all run loop modes.
	Invalidate()

	// The maximum number of frames per second a screen can render.
	MaximumFramesPerSecond() int
	SetMaximumFramesPerSecond(value int)
}

// Init initializes the instance.
func (d CADisplayLink) Init() CADisplayLink {
	rv := objc.Send[CADisplayLink](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d CADisplayLink) Autorelease() CADisplayLink {
	rv := objc.Send[CADisplayLink](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewCADisplayLink creates a new CADisplayLink instance.
func NewCADisplayLink() CADisplayLink {
	class := getCADisplayLinkClass()
	rv := objc.Send[CADisplayLink](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a display link for a target that calls its selector.
//
// target: An object in your app that you want the system to notify each time it
// updates a display.
//
// sel: A selector instance that represents a method for `target`.
//
// # Return Value
// 
// A new [CADisplayLink] object.
//
// # Discussion
// 
// The selector on the target must be a method with the following signature,
// where sender is the display link returned by this method.
// 
// The newly constructed display link retains the target.
//
// See: https://developer.apple.com/documentation/QuartzCore/CADisplayLink/init(target:selector:)
func NewDisplayLinkWithTargetSelector(target objectivec.IObject, sel objc.SEL) CADisplayLink {
	rv := objc.Send[objc.ID](objc.ID(getCADisplayLinkClass().class), objc.Sel("displayLinkWithTarget:selector:"), target, sel)
	return CADisplayLinkFromID(rv)
}

// Registers the display link with a run loop.
//
// runloop: The run loop to associate with the display link.
//
// mode: The mode in which to add the display link to the run loop.
//
// # Discussion
// 
// You can associate a display link with multiple input modes. While the run
// loop is executing in a mode you specify, the display link notifies the
// target when the system requires new frames.
// 
// You can specify a custom mode or use one of the modes listed in [RunLoop].
// 
// The run loop retains the display link. To remove the display link from all
// run loops, call [Invalidate].
//
// [RunLoop]: https://developer.apple.com/documentation/Foundation/RunLoop
//
// See: https://developer.apple.com/documentation/QuartzCore/CADisplayLink/add(to:forMode:)
func (d CADisplayLink) AddToRunLoopForMode(runloop foundation.NSRunLoop, mode foundation.NSString) {
	objc.Send[objc.ID](d.ID, objc.Sel("addToRunLoop:forMode:"), runloop, mode)
}
// Removes the display link from the run loop for the given mode.
//
// runloop: The run loop you associate with the display link.
//
// mode: The run loop mode in which the display link is running.
//
// # Discussion
// 
// The run loop releases the display link if it’s no longer associated with
// any run modes.
//
// See: https://developer.apple.com/documentation/QuartzCore/CADisplayLink/remove(from:forMode:)
func (d CADisplayLink) RemoveFromRunLoopForMode(runloop foundation.NSRunLoop, mode foundation.NSString) {
	objc.Send[objc.ID](d.ID, objc.Sel("removeFromRunLoop:forMode:"), runloop, mode)
}
// Removes the display link from all run loop modes.
//
// # Discussion
// 
// When you remove the display link from all run loop mode, the system
// releases it. The display link also releases the target.
// 
// This method is thread safe, so you can call it from a thread separate to
// the one in which the display link runs.
//
// See: https://developer.apple.com/documentation/QuartzCore/CADisplayLink/invalidate()
func (d CADisplayLink) Invalidate() {
	objc.Send[objc.ID](d.ID, objc.Sel("invalidate"))
}

// The time interval between screen refresh updates.
//
// # Discussion
// 
// This value is in an undefined state until the system calls the target’s
// selector at least once.
// 
// You calculate the expected amount of time your app has to render each frame
// by using [TargetTimestamp]-[Timestamp]. Use
// [TargetTimestamp]-[CACurrentMediaTime] to calculate the actual amount of
// time.
//
// See: https://developer.apple.com/documentation/QuartzCore/CADisplayLink/duration
func (d CADisplayLink) Duration() float64 {
	rv := objc.Send[float64](d.ID, objc.Sel("duration"))
	return rv
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
// [Optimizing ProMotion refresh rates for iPhone 13 Pro and iPad Pro]: https://developer.apple.com/documentation/QuartzCore/optimizing-promotion-refresh-rates-for-iphone-13-pro-and-ipad-pro
//
// See: https://developer.apple.com/documentation/QuartzCore/CADisplayLink/preferredFrameRateRange
func (d CADisplayLink) PreferredFrameRateRange() CAFrameRateRange {
	rv := objc.Send[CAFrameRateRange](d.ID, objc.Sel("preferredFrameRateRange"))
	return CAFrameRateRange(rv)
}
func (d CADisplayLink) SetPreferredFrameRateRange(value CAFrameRateRange) {
	objc.Send[struct{}](d.ID, objc.Sel("setPreferredFrameRateRange:"), value)
}
// A Boolean value that indicates whether the system suspends the display
// link’s notifications to the target.
//
// # Discussion
// 
// The default value is [false]. If [true], the display link doesn’t send
// notifications to the target.
// 
// This property is thread safe, so you can set it from a thread separate to
// the one in which the display link runs.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/QuartzCore/CADisplayLink/isPaused
func (d CADisplayLink) Paused() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("isPaused"))
	return rv
}
func (d CADisplayLink) SetPaused(value bool) {
	objc.Send[struct{}](d.ID, objc.Sel("setPaused:"), value)
}
// The time interval that represents when the last frame displayed.
//
// # Discussion
// 
// If you need to calculate what to display next, use [TargetTimestamp]
// instead.
//
// See: https://developer.apple.com/documentation/QuartzCore/CADisplayLink/timestamp
func (d CADisplayLink) Timestamp() float64 {
	rv := objc.Send[float64](d.ID, objc.Sel("timestamp"))
	return rv
}
// The time interval that represents when the next frame displays.
//
// # Discussion
// 
// You can use the target timestamp to cancel or pause long running processes
// that may overrun the available time between frames in order to maintain a
// consistent frame rate.
// 
// The following code shows how you can create a display link and register it
// with a run loop. The `step(``)` function attempts to sum the square roots
// of all numbers up to [max], but with each iteration checks the current time
// ([CACurrentMediaTime]) against the [TargetTimestamp]. If the time taken to
// complete the calculation is later than the target timestamp, the function
// breaks the loop:
//
// [max]: https://developer.apple.com/documentation/Swift/Int/max
//
// See: https://developer.apple.com/documentation/QuartzCore/CADisplayLink/targetTimestamp
func (d CADisplayLink) TargetTimestamp() float64 {
	rv := objc.Send[float64](d.ID, objc.Sel("targetTimestamp"))
	return rv
}
// The maximum number of frames per second a screen can render.
//
// See: https://developer.apple.com/documentation/UIKit/UIScreen/maximumFramesPerSecond
func (d CADisplayLink) MaximumFramesPerSecond() int {
	rv := objc.Send[int](d.ID, objc.Sel("maximumFramesPerSecond"))
	return rv
}
func (d CADisplayLink) SetMaximumFramesPerSecond(value int) {
	objc.Send[struct{}](d.ID, objc.Sel("setMaximumFramesPerSecond:"), value)
}

