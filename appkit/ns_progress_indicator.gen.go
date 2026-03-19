// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSProgressIndicator] class.
var (
	_NSProgressIndicatorClass     NSProgressIndicatorClass
	_NSProgressIndicatorClassOnce sync.Once
)

func getNSProgressIndicatorClass() NSProgressIndicatorClass {
	_NSProgressIndicatorClassOnce.Do(func() {
		_NSProgressIndicatorClass = NSProgressIndicatorClass{class: objc.GetClass("NSProgressIndicator")}
	})
	return _NSProgressIndicatorClass
}

// GetNSProgressIndicatorClass returns the class object for NSProgressIndicator.
func GetNSProgressIndicatorClass() NSProgressIndicatorClass {
	return getNSProgressIndicatorClass()
}

type NSProgressIndicatorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSProgressIndicatorClass) Alloc() NSProgressIndicator {
	rv := objc.Send[NSProgressIndicator](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An interface that provides visual feedback to the user about the status of
// an ongoing task.
//
// # Overview
// 
// Progress indicators can be determinate or indeterminate. A determinate
// indicator displays the completion percentage of a task. An indeterminate
// indicator shows that the app is busy without providing a visual indication
// of how long the task will take.
//
// # Animating the progress indicator
//
//   - [NSProgressIndicator.StartAnimation]: Starts the animation of an indeterminate progress indicator.
//   - [NSProgressIndicator.StopAnimation]: Stops the animation of an indeterminate progress indicator.
//   - [NSProgressIndicator.UsesThreadedAnimation]: A Boolean that indicates whether the progress indicator implements animation in a separate thread.
//   - [NSProgressIndicator.SetUsesThreadedAnimation]
//
// # Advancing the progress bar
//
//   - [NSProgressIndicator.IncrementBy]: Advances the progress bar of a determinate progress indicator by the specified amount.
//   - [NSProgressIndicator.DoubleValue]: The value that indicates the current extent of the progress indicator.
//   - [NSProgressIndicator.SetDoubleValue]
//   - [NSProgressIndicator.MinValue]: The minimum value for the progress indicator.
//   - [NSProgressIndicator.SetMinValue]
//   - [NSProgressIndicator.MaxValue]: The maximum value for the progress indicator.
//   - [NSProgressIndicator.SetMaxValue]
//
// # Observing the progress bar
//
//   - [NSProgressIndicator.ObservedProgress]: The progress object to use for updating the progress view.
//   - [NSProgressIndicator.SetObservedProgress]
//
// # Setting the appearance
//
//   - [NSProgressIndicator.ControlSize]: The size of the progress indicator.
//   - [NSProgressIndicator.SetControlSize]
//   - [NSProgressIndicator.ControlTint]: The progress indicator’s control tint.
//   - [NSProgressIndicator.SetControlTint]
//   - [NSProgressIndicator.Bezeled]: A Boolean that indicates whether the progress indicator’s frame has a three-dimensional bezel.
//   - [NSProgressIndicator.SetBezeled]
//   - [NSProgressIndicator.Indeterminate]: A Boolean that indicates whether the progress indicator is indeterminate.
//   - [NSProgressIndicator.SetIndeterminate]
//   - [NSProgressIndicator.Style]: The style of the progress indicator (bar or spinning).
//   - [NSProgressIndicator.SetStyle]
//   - [NSProgressIndicator.SizeToFit]: This action method resizes the progress indicator to an appropriate size depending on the value of [style](<doc://com.apple.appkit/documentation/AppKit/NSProgressIndicator/style-swift.property>).
//   - [NSProgressIndicator.DisplayedWhenStopped]: A Boolean that indicates whether the progress indicator hides itself when it isn’t animating.
//   - [NSProgressIndicator.SetDisplayedWhenStopped]
//
// See: https://developer.apple.com/documentation/AppKit/NSProgressIndicator
type NSProgressIndicator struct {
	NSView
}

// NSProgressIndicatorFromID constructs a [NSProgressIndicator] from an objc.ID.
//
// An interface that provides visual feedback to the user about the status of
// an ongoing task.
func NSProgressIndicatorFromID(id objc.ID) NSProgressIndicator {
	return NSProgressIndicator{NSView: NSViewFromID(id)}
}
// NOTE: NSProgressIndicator adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSProgressIndicator] class.
//
// # Animating the progress indicator
//
//   - [INSProgressIndicator.StartAnimation]: Starts the animation of an indeterminate progress indicator.
//   - [INSProgressIndicator.StopAnimation]: Stops the animation of an indeterminate progress indicator.
//   - [INSProgressIndicator.UsesThreadedAnimation]: A Boolean that indicates whether the progress indicator implements animation in a separate thread.
//   - [INSProgressIndicator.SetUsesThreadedAnimation]
//
// # Advancing the progress bar
//
//   - [INSProgressIndicator.IncrementBy]: Advances the progress bar of a determinate progress indicator by the specified amount.
//   - [INSProgressIndicator.DoubleValue]: The value that indicates the current extent of the progress indicator.
//   - [INSProgressIndicator.SetDoubleValue]
//   - [INSProgressIndicator.MinValue]: The minimum value for the progress indicator.
//   - [INSProgressIndicator.SetMinValue]
//   - [INSProgressIndicator.MaxValue]: The maximum value for the progress indicator.
//   - [INSProgressIndicator.SetMaxValue]
//
// # Observing the progress bar
//
//   - [INSProgressIndicator.ObservedProgress]: The progress object to use for updating the progress view.
//   - [INSProgressIndicator.SetObservedProgress]
//
// # Setting the appearance
//
//   - [INSProgressIndicator.ControlSize]: The size of the progress indicator.
//   - [INSProgressIndicator.SetControlSize]
//   - [INSProgressIndicator.ControlTint]: The progress indicator’s control tint.
//   - [INSProgressIndicator.SetControlTint]
//   - [INSProgressIndicator.Bezeled]: A Boolean that indicates whether the progress indicator’s frame has a three-dimensional bezel.
//   - [INSProgressIndicator.SetBezeled]
//   - [INSProgressIndicator.Indeterminate]: A Boolean that indicates whether the progress indicator is indeterminate.
//   - [INSProgressIndicator.SetIndeterminate]
//   - [INSProgressIndicator.Style]: The style of the progress indicator (bar or spinning).
//   - [INSProgressIndicator.SetStyle]
//   - [INSProgressIndicator.SizeToFit]: This action method resizes the progress indicator to an appropriate size depending on the value of [style](<doc://com.apple.appkit/documentation/AppKit/NSProgressIndicator/style-swift.property>).
//   - [INSProgressIndicator.DisplayedWhenStopped]: A Boolean that indicates whether the progress indicator hides itself when it isn’t animating.
//   - [INSProgressIndicator.SetDisplayedWhenStopped]
//
// See: https://developer.apple.com/documentation/AppKit/NSProgressIndicator
type INSProgressIndicator interface {
	INSView
	NSAccessibilityGroup
	NSAccessibilityProgressIndicator

	// Topic: Animating the progress indicator

	// Starts the animation of an indeterminate progress indicator.
	StartAnimation(sender objectivec.IObject)
	// Stops the animation of an indeterminate progress indicator.
	StopAnimation(sender objectivec.IObject)
	// A Boolean that indicates whether the progress indicator implements animation in a separate thread.
	UsesThreadedAnimation() bool
	SetUsesThreadedAnimation(value bool)

	// Topic: Advancing the progress bar

	// Advances the progress bar of a determinate progress indicator by the specified amount.
	IncrementBy(delta float64)
	// The value that indicates the current extent of the progress indicator.
	DoubleValue() float64
	SetDoubleValue(value float64)
	// The minimum value for the progress indicator.
	MinValue() float64
	SetMinValue(value float64)
	// The maximum value for the progress indicator.
	MaxValue() float64
	SetMaxValue(value float64)

	// Topic: Observing the progress bar

	// The progress object to use for updating the progress view.
	ObservedProgress() foundation.NSProgress
	SetObservedProgress(value foundation.NSProgress)

	// Topic: Setting the appearance

	// The size of the progress indicator.
	ControlSize() NSControlSize
	SetControlSize(value NSControlSize)
	// The progress indicator’s control tint.
	ControlTint() NSControlTint
	SetControlTint(value NSControlTint)
	// A Boolean that indicates whether the progress indicator’s frame has a three-dimensional bezel.
	Bezeled() bool
	SetBezeled(value bool)
	// A Boolean that indicates whether the progress indicator is indeterminate.
	Indeterminate() bool
	SetIndeterminate(value bool)
	// The style of the progress indicator (bar or spinning).
	Style() NSProgressIndicatorStyle
	SetStyle(value NSProgressIndicatorStyle)
	// This action method resizes the progress indicator to an appropriate size depending on the value of [style](<doc://com.apple.appkit/documentation/AppKit/NSProgressIndicator/style-swift.property>).
	SizeToFit()
	// A Boolean that indicates whether the progress indicator hides itself when it isn’t animating.
	DisplayedWhenStopped() bool
	SetDisplayedWhenStopped(value bool)
}

// Init initializes the instance.
func (p NSProgressIndicator) Init() NSProgressIndicator {
	rv := objc.Send[NSProgressIndicator](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSProgressIndicator) Autorelease() NSProgressIndicator {
	rv := objc.Send[NSProgressIndicator](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSProgressIndicator creates a new NSProgressIndicator instance.
func NewNSProgressIndicator() NSProgressIndicator {
	class := getNSProgressIndicatorClass()
	rv := objc.Send[NSProgressIndicator](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a view using from data in the specified coder object.
//
// coder: The coder object that contains the view’s configuration details.
//
// # Return Value
// 
// An initialized view or `nil` if AppKit couldn’t create the object.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/init(coder:)
func NewProgressIndicatorWithCoder(coder foundation.INSCoder) NSProgressIndicator {
	instance := getNSProgressIndicatorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSProgressIndicatorFromID(rv)
}

// Initializes and returns a newly allocated [NSView] object with a specified
// frame rectangle.
//
// frameRect: The frame rectangle for the created view object.
//
// # Return Value
// 
// An initialized view or `nil` if AppKit couldn’t create the object.
//
// # Discussion
// 
// Insert the view into your window’s view hieararchy before you can do
// anything with it. This method is the designated initializer for the
// [NSView] class.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/init(frame:)
func NewProgressIndicatorWithFrame(frameRect corefoundation.CGRect) NSProgressIndicator {
	instance := getNSProgressIndicatorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSProgressIndicatorFromID(rv)
}

// Starts the animation of an indeterminate progress indicator.
//
// sender: The object sending the message.
//
// # Discussion
// 
// Does nothing for a determinate progress indicator.
//
// See: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/startAnimation(_:)
func (p NSProgressIndicator) StartAnimation(sender objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("startAnimation:"), sender)
}
// Stops the animation of an indeterminate progress indicator.
//
// sender: The object sending the message.
//
// # Discussion
// 
// Does nothing for a determinate progress indicator.
//
// See: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/stopAnimation(_:)
func (p NSProgressIndicator) StopAnimation(sender objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("stopAnimation:"), sender)
}
// Advances the progress bar of a determinate progress indicator by the
// specified amount.
//
// delta: The amount by which to increment the progress bar. For example, if you want
// to advance a progress bar from 0.0 to 100.0 in 20 steps, you would invoke
// [IncrementBy] 20 times with a delta value of 5.0.
//
// See: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/increment(by:)
func (p NSProgressIndicator) IncrementBy(delta float64) {
	objc.Send[objc.ID](p.ID, objc.Sel("incrementBy:"), delta)
}
// This action method resizes the progress indicator to an appropriate size
// depending on the value of [Style].
//
// # Discussion
// 
// Use this after you use [Style] to re-size the progress indicator.
//
// See: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/sizeToFit()
func (p NSProgressIndicator) SizeToFit() {
	objc.Send[objc.ID](p.ID, objc.Sel("sizeToFit"))
}
// Returns the progress indicator’s value.
//
// # Return Value
// 
// The value of the progress indicator.
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityValue] property.
//
// [accessibilityValue]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityValue
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProgressIndicator/accessibilityValue()
func (p NSProgressIndicator) AccessibilityValue() foundation.NSNumber {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityValue"))
	return foundation.NSNumberFromID(rv)
}

// A Boolean that indicates whether the progress indicator implements
// animation in a separate thread.
//
// # Discussion
// 
// When the value of this property is [true], animation of the progress
// indicator occurs in a separate thread.
// 
// If the app becomes multithreaded as a result of an invocation of this
// method, the app’s performance could become noticeably slower.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/usesThreadedAnimation
func (p NSProgressIndicator) UsesThreadedAnimation() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("usesThreadedAnimation"))
	return rv
}
func (p NSProgressIndicator) SetUsesThreadedAnimation(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setUsesThreadedAnimation:"), value)
}
// The value that indicates the current extent of the progress indicator.
//
// # Discussion
// 
// By default, a determinate progress indicator goes from `0.0` to `100.0`. If
// the progress bar has advanced halfway across the view, this value would be
// `50.0`.
// 
// An indeterminate progress indicator does not use this value.
//
// See: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/doubleValue
func (p NSProgressIndicator) DoubleValue() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("doubleValue"))
	return rv
}
func (p NSProgressIndicator) SetDoubleValue(value float64) {
	objc.Send[struct{}](p.ID, objc.Sel("setDoubleValue:"), value)
}
// The minimum value for the progress indicator.
//
// # Discussion
// 
// By default, a determinate progress indicator goes from `0.0` to `100.0`, so
// the default value of this property is `0.0`.
// 
// An indeterminate progress indicator does not use this value.
//
// See: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/minValue
func (p NSProgressIndicator) MinValue() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("minValue"))
	return rv
}
func (p NSProgressIndicator) SetMinValue(value float64) {
	objc.Send[struct{}](p.ID, objc.Sel("setMinValue:"), value)
}
// The maximum value for the progress indicator.
//
// # Discussion
// 
// By default, a determinate progress indicator goes from `0.0` to `100.0`, so
// the default value of this property is `100.0`.
// 
// An indeterminate progress indicator does not use this value.
//
// See: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/maxValue
func (p NSProgressIndicator) MaxValue() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("maxValue"))
	return rv
}
func (p NSProgressIndicator) SetMaxValue(value float64) {
	objc.Send[struct{}](p.ID, objc.Sel("setMaxValue:"), value)
}
// The progress object to use for updating the progress view.
//
// # Discussion
// 
// Set this property when you want the progress view to automatically update
// its progress value using the information it receives from the [Progress]
// object. Setting this property also modifies the [Indeterminate],
// [MinValue], [MaxValue], and [DoubleValue] properties of the indicator. Set
// the property to `nil` when you want to update the progress manually. The
// default value of this property is `nil`.
// 
// For more information on configuring a progress object, see [Progress].
//
// [Progress]: https://developer.apple.com/documentation/Foundation/Progress
//
// See: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/observedProgress
func (p NSProgressIndicator) ObservedProgress() foundation.NSProgress {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("observedProgress"))
	return foundation.NSProgressFromID(objc.ID(rv))
}
func (p NSProgressIndicator) SetObservedProgress(value foundation.NSProgress) {
	objc.Send[struct{}](p.ID, objc.Sel("setObservedProgress:"), value)
}
// The size of the progress indicator.
//
// # Discussion
// 
// See [NSCell] for possible values.
//
// See: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/controlSize
func (p NSProgressIndicator) ControlSize() NSControlSize {
	rv := objc.Send[NSControlSize](p.ID, objc.Sel("controlSize"))
	return NSControlSize(rv)
}
func (p NSProgressIndicator) SetControlSize(value NSControlSize) {
	objc.Send[struct{}](p.ID, objc.Sel("setControlSize:"), value)
}
// The progress indicator’s control tint.
//
// # Discussion
// 
// See [NSCell] for possible values.
//
// See: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/controlTint
func (p NSProgressIndicator) ControlTint() NSControlTint {
	rv := objc.Send[NSControlTint](p.ID, objc.Sel("controlTint"))
	return NSControlTint(rv)
}
func (p NSProgressIndicator) SetControlTint(value NSControlTint) {
	objc.Send[struct{}](p.ID, objc.Sel("setControlTint:"), value)
}
// A Boolean that indicates whether the progress indicator’s frame has a
// three-dimensional bezel.
//
// # Discussion
// 
// When the value of this property is [true], the progress indicator is
// bezeled.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/isBezeled
func (p NSProgressIndicator) Bezeled() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isBezeled"))
	return rv
}
func (p NSProgressIndicator) SetBezeled(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setBezeled:"), value)
}
// A Boolean that indicates whether the progress indicator is indeterminate.
//
// # Discussion
// 
// A determinate indicator displays how much of the task has been completed.
// An indeterminate indicator shows simply that the app is busy.
// 
// When the value of this property is [true], the progress indicator is
// indeterminate; otherwise, it is determinate.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/isIndeterminate
func (p NSProgressIndicator) Indeterminate() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isIndeterminate"))
	return rv
}
func (p NSProgressIndicator) SetIndeterminate(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setIndeterminate:"), value)
}
// The style of the progress indicator (bar or spinning).
//
// # Discussion
// 
// See [NSProgressIndicator.Style] for possible values.
//
// [NSProgressIndicator.Style]: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/Style-swift.enum
//
// See: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/style-swift.property
func (p NSProgressIndicator) Style() NSProgressIndicatorStyle {
	rv := objc.Send[NSProgressIndicatorStyle](p.ID, objc.Sel("style"))
	return NSProgressIndicatorStyle(rv)
}
func (p NSProgressIndicator) SetStyle(value NSProgressIndicatorStyle) {
	objc.Send[struct{}](p.ID, objc.Sel("setStyle:"), value)
}
// A Boolean that indicates whether the progress indicator hides itself when
// it isn’t animating.
//
// # Discussion
// 
// When the value of this property is [false], the progress indicator is
// hidden when it isn’t animating. The default value of this property is
// [true].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSProgressIndicator/isDisplayedWhenStopped
func (p NSProgressIndicator) DisplayedWhenStopped() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isDisplayedWhenStopped"))
	return rv
}
func (p NSProgressIndicator) SetDisplayedWhenStopped(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setDisplayedWhenStopped:"), value)
}

			// Protocol methods for NSAccessibilityProgressIndicator
			
// Returns the accessibility element’s frame in screen coordinates.
//
// # Return Value
// 
// The element’s frame in screen coordinates.
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityFrame] property. This method is called whenever accessibility
// clients request the [size] or [position] attributes.
//
// [accessibilityFrame]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFrame
// [position]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/position
// [size]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/size
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityFrame()
func (o NSProgressIndicator) AccessibilityFrame() corefoundation.CGRect {
	
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("accessibilityFrame"))
	return rv
	}
// Returns the accessibility element’s parent in the accessibility
// hierarchy.
//
// # Return Value
// 
// The element’s parent in the accessibility hierarchy.
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityParent] property.
//
// [accessibilityParent]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityParent
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityParent()
func (o NSProgressIndicator) AccessibilityParent() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityParent"))
	return objectivec.Object{ID: rv}
	}
// Returns the accessibility element’s identity.
//
// # Return Value
// 
// Returns the unique ID for the accessibility element. It is often used in
// automated testing.
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityIdentifier] property.
//
// [accessibilityIdentifier]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityIdentifier
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityIdentifier()
func (o NSProgressIndicator) AccessibilityIdentifier() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityIdentifier"))
	return foundation.NSStringFromID(rv).String()
	}
// Returns a Boolean value that indicates whether the accessibility element
// has the keyboard focus.
//
// # Return Value
// 
// [true] if this element has the keyboard focus; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityFocused] property.
//
// [accessibilityFocused]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFocused
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/isAccessibilityFocused()
func (o NSProgressIndicator) IsAccessibilityFocused() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFocused"))
	return rv
	}

