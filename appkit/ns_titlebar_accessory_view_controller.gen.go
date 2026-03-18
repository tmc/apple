// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSTitlebarAccessoryViewController] class.
var (
	_NSTitlebarAccessoryViewControllerClass     NSTitlebarAccessoryViewControllerClass
	_NSTitlebarAccessoryViewControllerClassOnce sync.Once
)

func getNSTitlebarAccessoryViewControllerClass() NSTitlebarAccessoryViewControllerClass {
	_NSTitlebarAccessoryViewControllerClassOnce.Do(func() {
		_NSTitlebarAccessoryViewControllerClass = NSTitlebarAccessoryViewControllerClass{class: objc.GetClass("NSTitlebarAccessoryViewController")}
	})
	return _NSTitlebarAccessoryViewControllerClass
}

// GetNSTitlebarAccessoryViewControllerClass returns the class object for NSTitlebarAccessoryViewController.
func GetNSTitlebarAccessoryViewControllerClass() NSTitlebarAccessoryViewControllerClass {
	return getNSTitlebarAccessoryViewControllerClass()
}

type NSTitlebarAccessoryViewControllerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTitlebarAccessoryViewControllerClass) Alloc() NSTitlebarAccessoryViewController {
	rv := objc.Send[NSTitlebarAccessoryViewController](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that manages a custom view—known as an accessory view—in the
// title bar–toolbar area of a window.
//
// # Overview
// 
// Because a title bar accessory view controller is contained in a visual
// effect view (that is, [NSVisualEffectView]), it automatically handles the
// blur behind the accessory view and the size and location changes for the
// content of the view when a window goes in and out of full screen mode. If
// you’re currently using [NSToolbar] fullscreen accessory APIs, such as
// [fullScreenAccessoryView], you should use
// [NSTitlebarAccessoryViewController] APIs instead.
// 
// Typically, you create an [NSTitlebarAccessoryViewController] object, give
// it your custom view, set the [NSTitlebarAccessoryViewController.LayoutAttribute] property to ensure that it
// displays correctly in relation to the title bar, and add the view
// controller to your window. For more information about [NSWindow] methods
// you can use to add and remove a title bar accessory view controller, see
// Managing Title Bars.
// 
// Don’t override the `view` property in your
// [NSTitlebarAccessoryViewController] subclass. Instead, you can override
// [LoadView], and set the `view` property in that method.
//
// [fullScreenAccessoryView]: https://developer.apple.com/documentation/AppKit/NSToolbar/fullScreenAccessoryView
//
// # Configuring a title bar accessory view controller
//
//   - [NSTitlebarAccessoryViewController.FullScreenMinHeight]: The visual minimum height of an accessory view that displays below the title bar when the window is in full screen mode.
//   - [NSTitlebarAccessoryViewController.SetFullScreenMinHeight]
//   - [NSTitlebarAccessoryViewController.LayoutAttribute]: The location of the accessory view, in relation to the window’s title bar.
//   - [NSTitlebarAccessoryViewController.SetLayoutAttribute]
//
// # Configuring the scroll edge effect
//
//   - [NSTitlebarAccessoryViewController.PreferredScrollEdgeEffectStyle]: The titlebar accessory’s preferred effect for content scrolling behind it.
//   - [NSTitlebarAccessoryViewController.SetPreferredScrollEdgeEffectStyle]
//
// # Instance Properties
//
//   - [NSTitlebarAccessoryViewController.AutomaticallyAdjustsSize]
//   - [NSTitlebarAccessoryViewController.SetAutomaticallyAdjustsSize]
//   - [NSTitlebarAccessoryViewController.Hidden]
//   - [NSTitlebarAccessoryViewController.SetHidden]
//
// See: https://developer.apple.com/documentation/AppKit/NSTitlebarAccessoryViewController
type NSTitlebarAccessoryViewController struct {
	NSViewController
}

// NSTitlebarAccessoryViewControllerFromID constructs a [NSTitlebarAccessoryViewController] from an objc.ID.
//
// An object that manages a custom view—known as an accessory view—in the
// title bar–toolbar area of a window.
func NSTitlebarAccessoryViewControllerFromID(id objc.ID) NSTitlebarAccessoryViewController {
	return NSTitlebarAccessoryViewController{NSViewController: NSViewControllerFromID(id)}
}
// NOTE: NSTitlebarAccessoryViewController adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSTitlebarAccessoryViewController] class.
//
// # Configuring a title bar accessory view controller
//
//   - [INSTitlebarAccessoryViewController.FullScreenMinHeight]: The visual minimum height of an accessory view that displays below the title bar when the window is in full screen mode.
//   - [INSTitlebarAccessoryViewController.SetFullScreenMinHeight]
//   - [INSTitlebarAccessoryViewController.LayoutAttribute]: The location of the accessory view, in relation to the window’s title bar.
//   - [INSTitlebarAccessoryViewController.SetLayoutAttribute]
//
// # Configuring the scroll edge effect
//
//   - [INSTitlebarAccessoryViewController.PreferredScrollEdgeEffectStyle]: The titlebar accessory’s preferred effect for content scrolling behind it.
//   - [INSTitlebarAccessoryViewController.SetPreferredScrollEdgeEffectStyle]
//
// # Instance Properties
//
//   - [INSTitlebarAccessoryViewController.AutomaticallyAdjustsSize]
//   - [INSTitlebarAccessoryViewController.SetAutomaticallyAdjustsSize]
//   - [INSTitlebarAccessoryViewController.Hidden]
//   - [INSTitlebarAccessoryViewController.SetHidden]
//
// See: https://developer.apple.com/documentation/AppKit/NSTitlebarAccessoryViewController
type INSTitlebarAccessoryViewController interface {
	INSViewController
	NSAnimationDelegate

	// Topic: Configuring a title bar accessory view controller

	// The visual minimum height of an accessory view that displays below the title bar when the window is in full screen mode.
	FullScreenMinHeight() float64
	SetFullScreenMinHeight(value float64)
	// The location of the accessory view, in relation to the window’s title bar.
	LayoutAttribute() NSLayoutAttribute
	SetLayoutAttribute(value NSLayoutAttribute)

	// Topic: Configuring the scroll edge effect

	// The titlebar accessory’s preferred effect for content scrolling behind it.
	PreferredScrollEdgeEffectStyle() INSScrollEdgeEffectStyle
	SetPreferredScrollEdgeEffectStyle(value INSScrollEdgeEffectStyle)

	// Topic: Instance Properties

	AutomaticallyAdjustsSize() bool
	SetAutomaticallyAdjustsSize(value bool)
	Hidden() bool
	SetHidden(value bool)

	// The toolbar’s full screen accessory view.
	FullScreenAccessoryView() INSView
	SetFullScreenAccessoryView(value INSView)
	// Returns the animation that should be performed for the specified key.
	AnimationForKey(key NSAnimatablePropertyKey) objectivec.IObject
	// Returns a proxy object for the receiver that can be used to initiate implied animation for property changes.
	Animator() INSTitlebarAccessoryViewController
}

// Init initializes the instance.
func (t NSTitlebarAccessoryViewController) Init() NSTitlebarAccessoryViewController {
	rv := objc.Send[NSTitlebarAccessoryViewController](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTitlebarAccessoryViewController) Autorelease() NSTitlebarAccessoryViewController {
	rv := objc.Send[NSTitlebarAccessoryViewController](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTitlebarAccessoryViewController creates a new NSTitlebarAccessoryViewController instance.
func NewNSTitlebarAccessoryViewController() NSTitlebarAccessoryViewController {
	class := getNSTitlebarAccessoryViewControllerClass()
	rv := objc.Send[NSTitlebarAccessoryViewController](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/init(coder:)
func NewTitlebarAccessoryViewControllerWithCoder(coder foundation.INSCoder) NSTitlebarAccessoryViewController {
	instance := getNSTitlebarAccessoryViewControllerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSTitlebarAccessoryViewControllerFromID(rv)
}

// Returns a view controller object initialized to the nib file in the
// specified bundle.
//
// nibNameOrNil: The name of the nib file, without any leading path information.
//
// nibBundleOrNil: The bundle in which to search for the nib file. If you specify `nil`, this
// method looks for the nib file in the main bundle.
//
// # Return Value
// 
// The initialized [NSViewController] object.
//
// # Discussion
// 
// The [NSViewController] object looks for the nib file in the bundle’s
// language-specific project directories first, followed by the Resources
// directory.
// 
// The specified nib file should typically have the class of the file’s
// owner set to [NSViewController], or a custom subclass, with the `view`
// outlet connected to a view.
// 
// If you pass in `nil` for `nibNameOrNil`, [NibName] returns `nil` and
// [LoadView] throws an exception; in this case you must set [View] before
// [View] is invoked, or override [LoadView].
//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/init(nibName:bundle:)
func NewTitlebarAccessoryViewControllerWithNibNameBundle(nibNameOrNil NSNibName, nibBundleOrNil foundation.NSBundle) NSTitlebarAccessoryViewController {
	instance := getNSTitlebarAccessoryViewControllerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithNibName:bundle:"), objc.String(string(nibNameOrNil)), nibBundleOrNil)
	return NSTitlebarAccessoryViewControllerFromID(rv)
}

// Sent to the delegate when the specified animation completes its run.
//
// animation: The [NSAnimation] instance that completed its run.
//
// # Discussion
// 
// When an [NSAnimation] object reaches the end of its planned duration, it
// has a progress value of 1.0.
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimationDelegate/animationDidEnd(_:)
func (t NSTitlebarAccessoryViewController) AnimationDidEnd(animation INSAnimation) {
	objc.Send[objc.ID](t.ID, objc.Sel("animationDidEnd:"), animation)
}

// Sent to the delegate when an animation reaches a specific progress mark.
//
// animation: A running [NSAnimation] object that has reached a progress mark.
//
// progress: A `float` value (typed as [NSAnimationProgress]) that indicates a progress
// mark of `animation`.
//
// # Discussion
// 
// The delegate typically implements this method to perform some animation
// effect for the time slice indicated by `progress`, such as redrawing
// objects in a view with new coordinates or changing the frame location or
// size of a window or view. As an alternative to this delegation message, you
// may choose to observe the [progressMarkNotification] notification.
//
// [progressMarkNotification]: https://developer.apple.com/documentation/AppKit/NSAnimation/progressMarkNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimationDelegate/animation(_:didReachProgressMark:)
func (t NSTitlebarAccessoryViewController) AnimationDidReachProgressMark(animation INSAnimation, progress NSAnimationProgress) {
	objc.Send[objc.ID](t.ID, objc.Sel("animation:didReachProgressMark:"), animation, progress)
}

// Sent to the delegate when the specified animation is stopped before it
// completes its run.
//
// animation: The [NSAnimation] instance that was stopped.
//
// # Discussion
// 
// An [NSAnimation] object stops running when it receives a [StopAnimation]
// message.
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimationDelegate/animationDidStop(_:)
func (t NSTitlebarAccessoryViewController) AnimationDidStop(animation INSAnimation) {
	objc.Send[objc.ID](t.ID, objc.Sel("animationDidStop:"), animation)
}

// Returns the animation that should be performed for the specified key.
//
// key: The action name or property specified as a string.
//
// # Return Value
// 
// The animation to perform. A subclass of [CAAnimation].
//
// [CAAnimation]: https://developer.apple.com/documentation/QuartzCore/CAAnimation
//
// # Discussion
// 
// When the action specified by `key` is triggered for an object, this method
// is consulted to find the animation, if any, that should be performed in
// response.
// 
// Like its Core Animation [CALayer] counterpart, [action(forKey:)], this
// method is a funnel point that defines the order in which the search for an
// animation proceeds.It first checks the receiver’s Getting the Animator
// Proxy dictionary for a value matching `key`, then falls back to [Animator]
// for the receiver’s class.
// 
// Subclasses should not typically need to override this method.
//
// [CALayer]: https://developer.apple.com/documentation/QuartzCore/CALayer
// [action(forKey:)]: https://developer.apple.com/documentation/QuartzCore/CALayer/action(forKey:)
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimatablePropertyContainer/animation(forKey:)
func (t NSTitlebarAccessoryViewController) AnimationForKey(key NSAnimatablePropertyKey) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("animationForKey:"), objc.String(string(key)))
	return objectivec.Object{ID: rv}
}

// Sent to the delegate just after an animation is started.
//
// animation: The [NSAnimation] object that was just started.
//
// # Return Value
// 
// [false] to cancel the animation, [true] to have the animation proceed.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The delegate is sent this message just after `animation` receives a
// [StartAnimation] message. The delegate can use this method to prepare
// objects and resources for the effect.
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimationDelegate/animationShouldStart(_:)
func (t NSTitlebarAccessoryViewController) AnimationShouldStart(animation INSAnimation) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("animationShouldStart:"), animation)
	return rv
}

// Requests a custom curve value for the current progress value.
//
// animation: An [NSAnimation] object that is running.
//
// progress: A `float` value (typed as [NSAnimationProgress]) that indicates a progress
// mark of `animation`. This value is always between 0.0 and 1.0.
//
// # Return Value
// 
// A `float` value representing a custom curve.
//
// # Discussion
// 
// The delegate can compute and return a custom curve value for the given
// progress value. If the delegate does not implement this method,
// [NSAnimation] computes the current curve value.
// 
// The animation:valueForProgress: message is sent to the delegate when an
// [NSAnimation] object receives a [CurrentValue] message. The value the
// delegate returns is used as the value of [CurrentValue]; if there is no
// delegate, or it doesn’t implement animation:valueForProgress:,
// [NSAnimation] computes and returns the current value. [NSAnimation] does
// not invoke [CurrentValue]itself, but subclasses might.
// 
// See the description of [CurrentValue] for more information.
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimationDelegate/animation(_:valueForProgress:)
func (t NSTitlebarAccessoryViewController) AnimationValueForProgress(animation INSAnimation, progress NSAnimationProgress) float32 {
	rv := objc.Send[float32](t.ID, objc.Sel("animation:valueForProgress:"), animation, progress)
	return rv
}

// Returns a proxy object for the receiver that can be used to initiate
// implied animation for property changes.
//
// # Return Value
// 
// Returns a proxy object for the receiver that can initiate implied
// animations in response to property changes.
//
// # Discussion
// 
// The animator proxy object should be treated as if it was the receiver
// itself, and may be passed to any code that accepts the receiver as a
// parameter.
// 
// Sending key-value coding compliant “set” messages to the proxy will
// trigger animation for automatically animated properties of its target
// object, if the active [NSAnimationContext] in the current thread has a
// duration value greater than zero, and an animation for the property key is
// found by the [NSAnimatablePropertyContainer] search mechanism.
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimatablePropertyContainer/animator()
func (t NSTitlebarAccessoryViewController) Animator() INSTitlebarAccessoryViewController {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("animator"))
	return NSTitlebarAccessoryViewControllerFromID(rv)
}

// Returns the default animation that should be performed for the specified
// key.
//
// key: The action name or property specified as a string.
//
// # Return Value
// 
// The animation to perform. A subclass of [CAAnimation].
//
// [CAAnimation]: https://developer.apple.com/documentation/QuartzCore/CAAnimation
//
// # Discussion
// 
// The [NSAnimatablePropertyContainer] method consults this class method when
// its search of the receivers Getting the Animator Proxy dictionary fails to
// return an animation for `key`.
// 
// An animatable property container should implement this method to return a
// default animation to be performed for each key that it wants to make
// auto-animatable, where `key` usually references a property of the receiver,
// but can also specify a special animation trigger
// ([NSAnimationTriggerOrderIn] or [NSAnimationTriggerOrderOut]).
// 
// A developer implementing a custom view subclass, can enable automatic
// animation for properties by overriding this method, and having it return
// the desired default [CAAnimation] subclass to use for each of the property
// keys of interest. The override should defer to super for any keys it
// doesn’t specifically handle, facilitating inheritance of default
// animation specifications. The following is an example of such an
// implementation.
//
// [CAAnimation]: https://developer.apple.com/documentation/QuartzCore/CAAnimation
// [NSAnimationTriggerOrderIn]: https://developer.apple.com/documentation/AppKit/NSAnimationTriggerOrderIn
// [NSAnimationTriggerOrderOut]: https://developer.apple.com/documentation/AppKit/NSAnimationTriggerOrderOut
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimatablePropertyContainer/defaultAnimation(forKey:)
func (_NSTitlebarAccessoryViewControllerClass NSTitlebarAccessoryViewControllerClass) DefaultAnimationForKey(key NSAnimatablePropertyKey) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_NSTitlebarAccessoryViewControllerClass.class), objc.Sel("defaultAnimationForKey:"), objc.String(string(key)))
	return objectivec.Object{ID: rv}
}

// The visual minimum height of an accessory view that displays below the
// title bar when the window is in full screen mode.
//
// # Discussion
// 
// The [FullScreenMinHeight] property applies only to an
// [NSTitlebarAccessoryViewController] object in which [LayoutAttribute] is
// set to [NSLayoutConstraint.Attribute.bottom]. The minimum height you set
// determines how much of your accessory view is visible when the menu bar is
// hidden during full screen mode. By default, the minimum height is `0`,
// which means that the accessory view is hidden when the menu bar is hidden.
// When a user reveals the menu bar in this scenario, the accessory view is
// also revealed.
// 
// To persistently show a portion of the accessory view, set this property to
// a value greater than `0`. For example, if you have a fixed height accessory
// view, you can set [FullScreenMinHeight] to
// `view.FrameXCUIElementTypeSizeXCUIElementTypeHeight()` to show the view
// regardless of whether the menu bar is hidden. Note that the view’s height
// is never actually changed when it is hidden or revealed; instead, it is
// automatically clipped by an internal clip view.
//
// [NSLayoutConstraint.Attribute.bottom]: https://developer.apple.com/documentation/UIKit/NSLayoutConstraint/Attribute/bottom
//
// See: https://developer.apple.com/documentation/AppKit/NSTitlebarAccessoryViewController/fullScreenMinHeight
func (t NSTitlebarAccessoryViewController) FullScreenMinHeight() float64 {
	rv := objc.Send[float64](t.ID, objc.Sel("fullScreenMinHeight"))
	return rv
}
func (t NSTitlebarAccessoryViewController) SetFullScreenMinHeight(value float64) {
	objc.Send[struct{}](t.ID, objc.Sel("setFullScreenMinHeight:"), value)
}

// The location of the accessory view, in relation to the window’s title
// bar.
//
// # Discussion
// 
// The default value of this property is
// [NSLayoutConstraint.Attribute.bottom], which means that the accessory view
// should display below the title bar. You can also set this property to
// [NSLayoutConstraint.Attribute.right] or (in apps linked on macOS 10.11 or
// later) [NSLayoutConstraint.Attribute.left]. All other values are invalid
// and will cause an assertion to be raised.
//
// [NSLayoutConstraint.Attribute.bottom]: https://developer.apple.com/documentation/UIKit/NSLayoutConstraint/Attribute/bottom
// [NSLayoutConstraint.Attribute.left]: https://developer.apple.com/documentation/UIKit/NSLayoutConstraint/Attribute/left
// [NSLayoutConstraint.Attribute.right]: https://developer.apple.com/documentation/UIKit/NSLayoutConstraint/Attribute/right
//
// See: https://developer.apple.com/documentation/AppKit/NSTitlebarAccessoryViewController/layoutAttribute
func (t NSTitlebarAccessoryViewController) LayoutAttribute() NSLayoutAttribute {
	rv := objc.Send[NSLayoutAttribute](t.ID, objc.Sel("layoutAttribute"))
	return NSLayoutAttribute(rv)
}
func (t NSTitlebarAccessoryViewController) SetLayoutAttribute(value NSLayoutAttribute) {
	objc.Send[struct{}](t.ID, objc.Sel("setLayoutAttribute:"), value)
}

// The titlebar accessory’s preferred effect for content scrolling behind
// it.
//
// # Discussion
// 
// To allow for a soft edge on the bottom edge of a titlebar accessory:
//
// See: https://developer.apple.com/documentation/AppKit/NSTitlebarAccessoryViewController/preferredScrollEdgeEffectStyle
func (t NSTitlebarAccessoryViewController) PreferredScrollEdgeEffectStyle() INSScrollEdgeEffectStyle {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("preferredScrollEdgeEffectStyle"))
	return NSScrollEdgeEffectStyleFromID(objc.ID(rv))
}
func (t NSTitlebarAccessoryViewController) SetPreferredScrollEdgeEffectStyle(value INSScrollEdgeEffectStyle) {
	objc.Send[struct{}](t.ID, objc.Sel("setPreferredScrollEdgeEffectStyle:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSTitlebarAccessoryViewController/automaticallyAdjustsSize
func (t NSTitlebarAccessoryViewController) AutomaticallyAdjustsSize() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("automaticallyAdjustsSize"))
	return rv
}
func (t NSTitlebarAccessoryViewController) SetAutomaticallyAdjustsSize(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setAutomaticallyAdjustsSize:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSTitlebarAccessoryViewController/isHidden
func (t NSTitlebarAccessoryViewController) Hidden() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isHidden"))
	return rv
}
func (t NSTitlebarAccessoryViewController) SetHidden(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setHidden:"), value)
}

// Sets the option dictionary that maps event trigger keys to animation
// objects.
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimatablePropertyContainer/animations
func (t NSTitlebarAccessoryViewController) Animations() foundation.INSDictionary {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("animations"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (t NSTitlebarAccessoryViewController) SetAnimations(value foundation.INSDictionary) {
	objc.Send[struct{}](t.ID, objc.Sel("setAnimations:"), value)
}

// The toolbar’s full screen accessory view.
//
// See: https://developer.apple.com/documentation/appkit/nstoolbar/fullscreenaccessoryview
func (t NSTitlebarAccessoryViewController) FullScreenAccessoryView() INSView {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("fullScreenAccessoryView"))
	return NSViewFromID(objc.ID(rv))
}
func (t NSTitlebarAccessoryViewController) SetFullScreenAccessoryView(value INSView) {
	objc.Send[struct{}](t.ID, objc.Sel("setFullScreenAccessoryView:"), value)
}

			// Protocol methods for NSAnimationDelegate
			

