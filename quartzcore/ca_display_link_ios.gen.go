// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.
//go:build ios
// +build ios

package quartzcore

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

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
func (_CADisplayLinkClass CADisplayLinkClass) DisplayLinkWithTargetSelector(target objectivec.IObject, sel objc.SEL) CADisplayLink {
	rv := objc.Send[objc.ID](objc.ID(_CADisplayLinkClass.class), objc.Sel("displayLinkWithTarget:selector:"), target, sel)
	return CADisplayLinkFromID(rv)
}

// A frequency your app prefers for frame updates, affecting how often the
// system invokes your delegate’s callback.
//
// # Discussion
//
// The display link makes a best attempt to invoke your app’s callback at
// the frequency value you set to this property. However, the system also
// takes into account the device’s hardware capabilities and the other tasks
// your game or app is running.
//
// In iOS 15 and later, the system can change the available range of frame
// rates because it factors in system policies and a person’s preferences.
// For example, Low Power Mode, critical thermal state, and accessibility
// settings can affect the system’s frame rate.
//
// The system typically provides a consistent frame rate by choosing one
// that’s a factor of the display’s maximum refresh rate. For example, a
// display link could invoke your callback 60 times per second for a display
// with a refresh rate of 60 hertz. However, the display link could invoke
// your callback less frequently, such as 30, 20, or 15 times per second, by
// setting a smaller value.
//
// See [Optimizing ProMotion refresh rates for iPhone 13 Pro and iPad Pro] for
// more information.
//
// See: https://developer.apple.com/documentation/QuartzCore/CADisplayLink/preferredFramesPerSecond
//
// [Optimizing ProMotion refresh rates for iPhone 13 Pro and iPad Pro]: https://developer.apple.com/documentation/QuartzCore/optimizing-promotion-refresh-rates-for-iphone-13-pro-and-ipad-pro
func (d CADisplayLink) PreferredFramesPerSecond() int {
	rv := objc.Send[int](d.ID, objc.Sel("preferredFramesPerSecond"))
	return rv
}
func (d CADisplayLink) SetPreferredFramesPerSecond(value int) {
	objc.Send[struct{}](d.ID, objc.Sel("setPreferredFramesPerSecond:"), value)
}

// The number of frames that must pass before the display link notifies the
// target again.
//
// # Discussion
//
// The default value is `1`, which results in the system notifying your app at
// the refresh rate of the display. If you set the value to a value greater
// than `1`, the display link notifies your app at a fraction of the native
// refresh rate. For example, setting the interval to `2` causes the display
// link to fire every other frame, providing half the frame rate.
//
// Setting this value to less than `1` results in undefined behavior and is a
// programmer error.
//
// See: https://developer.apple.com/documentation/QuartzCore/CADisplayLink/frameInterval
func (d CADisplayLink) FrameInterval() int {
	rv := objc.Send[int](d.ID, objc.Sel("frameInterval"))
	return rv
}
func (d CADisplayLink) SetFrameInterval(value int) {
	objc.Send[struct{}](d.ID, objc.Sel("setFrameInterval:"), value)
}
