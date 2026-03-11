// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSSplitViewItemAccessoryViewController] class.
var (
	_NSSplitViewItemAccessoryViewControllerClass     NSSplitViewItemAccessoryViewControllerClass
	_NSSplitViewItemAccessoryViewControllerClassOnce sync.Once
)

func getNSSplitViewItemAccessoryViewControllerClass() NSSplitViewItemAccessoryViewControllerClass {
	_NSSplitViewItemAccessoryViewControllerClassOnce.Do(func() {
		_NSSplitViewItemAccessoryViewControllerClass = NSSplitViewItemAccessoryViewControllerClass{class: objc.GetClass("NSSplitViewItemAccessoryViewController")}
	})
	return _NSSplitViewItemAccessoryViewControllerClass
}

// GetNSSplitViewItemAccessoryViewControllerClass returns the class object for NSSplitViewItemAccessoryViewController.
func GetNSSplitViewItemAccessoryViewControllerClass() NSSplitViewItemAccessoryViewControllerClass {
	return getNSSplitViewItemAccessoryViewControllerClass()
}

type NSSplitViewItemAccessoryViewControllerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSplitViewItemAccessoryViewControllerClass) Alloc() NSSplitViewItemAccessoryViewController {
	rv := objc.Send[NSSplitViewItemAccessoryViewController](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







//
// # Configuring the scroll edge effect
//
//   - [NSSplitViewItemAccessoryViewController.PreferredScrollEdgeEffectStyle]: The split view item accessory’s preferred effect for content scrolling behind it.
//   - [NSSplitViewItemAccessoryViewController.SetPreferredScrollEdgeEffectStyle]
//
// # Instance Properties
//
//   - [NSSplitViewItemAccessoryViewController.AutomaticallyAppliesContentInsets]: Whether or not standard content insets should be applied to the view. Defaults to YES.
//   - [NSSplitViewItemAccessoryViewController.SetAutomaticallyAppliesContentInsets]
//   - [NSSplitViewItemAccessoryViewController.Hidden]: When set, this property will collapse the accessory view to 0 height (animatable) but not remove it from the window. Set through the animator object to animate it.
//   - [NSSplitViewItemAccessoryViewController.SetHidden]
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewItemAccessoryViewController
type NSSplitViewItemAccessoryViewController struct {
	NSViewController
}

// NSSplitViewItemAccessoryViewControllerFromID constructs a [NSSplitViewItemAccessoryViewController] from an objc.ID.
func NSSplitViewItemAccessoryViewControllerFromID(id objc.ID) NSSplitViewItemAccessoryViewController {
	return NSSplitViewItemAccessoryViewController{NSViewController: NSViewControllerFromID(id)}
}
// NOTE: NSSplitViewItemAccessoryViewController adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSSplitViewItemAccessoryViewController] class.
//
// # Configuring the scroll edge effect
//
//   - [INSSplitViewItemAccessoryViewController.PreferredScrollEdgeEffectStyle]: The split view item accessory’s preferred effect for content scrolling behind it.
//   - [INSSplitViewItemAccessoryViewController.SetPreferredScrollEdgeEffectStyle]
//
// # Instance Properties
//
//   - [INSSplitViewItemAccessoryViewController.AutomaticallyAppliesContentInsets]: Whether or not standard content insets should be applied to the view. Defaults to YES.
//   - [INSSplitViewItemAccessoryViewController.SetAutomaticallyAppliesContentInsets]
//   - [INSSplitViewItemAccessoryViewController.Hidden]: When set, this property will collapse the accessory view to 0 height (animatable) but not remove it from the window. Set through the animator object to animate it.
//   - [INSSplitViewItemAccessoryViewController.SetHidden]
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewItemAccessoryViewController
type INSSplitViewItemAccessoryViewController interface {
	INSViewController

	// Topic: Configuring the scroll edge effect

	// The split view item accessory’s preferred effect for content scrolling behind it.
	PreferredScrollEdgeEffectStyle() INSScrollEdgeEffectStyle
	SetPreferredScrollEdgeEffectStyle(value INSScrollEdgeEffectStyle)

	// Topic: Instance Properties

	// Whether or not standard content insets should be applied to the view. Defaults to YES.
	AutomaticallyAppliesContentInsets() bool
	SetAutomaticallyAppliesContentInsets(value bool)
	// When set, this property will collapse the accessory view to 0 height (animatable) but not remove it from the window. Set through the animator object to animate it.
	Hidden() bool
	SetHidden(value bool)

	BottomAlignedAccessoryViewControllers() INSSplitViewItemAccessoryViewController
	SetBottomAlignedAccessoryViewControllers(value INSSplitViewItemAccessoryViewController)
	// The following methods allow you to add accessory views to the top/bottom of this splitViewItem. See 
	TopAlignedAccessoryViewControllers() INSSplitViewItemAccessoryViewController
	SetTopAlignedAccessoryViewControllers(value INSSplitViewItemAccessoryViewController)
	// Returns the animation that should be performed for the specified key.
	AnimationForKey(key NSAnimatablePropertyKey) objectivec.IObject
	// Returns a proxy object for the receiver that can be used to initiate implied animation for property changes.
	Animator() INSSplitViewItemAccessoryViewController
	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (s NSSplitViewItemAccessoryViewController) Init() NSSplitViewItemAccessoryViewController {
	rv := objc.Send[NSSplitViewItemAccessoryViewController](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSplitViewItemAccessoryViewController) Autorelease() NSSplitViewItemAccessoryViewController {
	rv := objc.Send[NSSplitViewItemAccessoryViewController](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSplitViewItemAccessoryViewController creates a new NSSplitViewItemAccessoryViewController instance.
func NewNSSplitViewItemAccessoryViewController() NSSplitViewItemAccessoryViewController {
	class := getNSSplitViewItemAccessoryViewControllerClass()
	rv := objc.Send[NSSplitViewItemAccessoryViewController](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/init(coder:)
func NewSplitViewItemAccessoryViewControllerWithCoder(coder foundation.INSCoder) NSSplitViewItemAccessoryViewController {
	instance := getNSSplitViewItemAccessoryViewControllerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSSplitViewItemAccessoryViewControllerFromID(rv)
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
func NewSplitViewItemAccessoryViewControllerWithNibNameBundle(nibNameOrNil NSNibName, nibBundleOrNil foundation.NSBundle) NSSplitViewItemAccessoryViewController {
	instance := getNSSplitViewItemAccessoryViewControllerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithNibName:bundle:"), objc.String(string(nibNameOrNil)), nibBundleOrNil)
	return NSSplitViewItemAccessoryViewControllerFromID(rv)
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
func (s NSSplitViewItemAccessoryViewController) AnimationForKey(key NSAnimatablePropertyKey) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("animationForKey:"), objc.String(string(key)))
	return objectivec.Object{ID: rv}
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
func (s NSSplitViewItemAccessoryViewController) Animator() INSSplitViewItemAccessoryViewController {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("animator"))
	return NSSplitViewItemAccessoryViewControllerFromID(rv)
}
func (s NSSplitViewItemAccessoryViewController) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
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
func (_NSSplitViewItemAccessoryViewControllerClass NSSplitViewItemAccessoryViewControllerClass) DefaultAnimationForKey(key NSAnimatablePropertyKey) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_NSSplitViewItemAccessoryViewControllerClass.class), objc.Sel("defaultAnimationForKey:"), objc.String(string(key)))
	return objectivec.Object{ID: rv}
}








// The split view item accessory’s preferred effect for content scrolling
// behind it.
//
// # Discussion
// 
// To allow for a soft edge on the interior edge of a titlebar accessory:
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewItemAccessoryViewController/preferredScrollEdgeEffectStyle
func (s NSSplitViewItemAccessoryViewController) PreferredScrollEdgeEffectStyle() INSScrollEdgeEffectStyle {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("preferredScrollEdgeEffectStyle"))
	return NSScrollEdgeEffectStyleFromID(objc.ID(rv))
}
func (s NSSplitViewItemAccessoryViewController) SetPreferredScrollEdgeEffectStyle(value INSScrollEdgeEffectStyle) {
	objc.Send[struct{}](s.ID, objc.Sel("setPreferredScrollEdgeEffectStyle:"), value)
}



// Whether or not standard content insets should be applied to the view.
// Defaults to YES.
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewItemAccessoryViewController/automaticallyAppliesContentInsets
func (s NSSplitViewItemAccessoryViewController) AutomaticallyAppliesContentInsets() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("automaticallyAppliesContentInsets"))
	return rv
}
func (s NSSplitViewItemAccessoryViewController) SetAutomaticallyAppliesContentInsets(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setAutomaticallyAppliesContentInsets:"), value)
}



// When set, this property will collapse the accessory view to 0 height
// (animatable) but not remove it from the window. Set through the animator
// object to animate it.
//
// See: https://developer.apple.com/documentation/AppKit/NSSplitViewItemAccessoryViewController/isHidden
func (s NSSplitViewItemAccessoryViewController) Hidden() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isHidden"))
	return rv
}
func (s NSSplitViewItemAccessoryViewController) SetHidden(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setHidden:"), value)
}



// Sets the option dictionary that maps event trigger keys to animation
// objects.
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimatablePropertyContainer/animations
func (s NSSplitViewItemAccessoryViewController) Animations() foundation.INSDictionary {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("animations"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (s NSSplitViewItemAccessoryViewController) SetAnimations(value foundation.INSDictionary) {
	objc.Send[struct{}](s.ID, objc.Sel("setAnimations:"), value)
}



// See: https://developer.apple.com/documentation/appkit/nssplitviewitem/bottomalignedaccessoryviewcontrollers
func (s NSSplitViewItemAccessoryViewController) BottomAlignedAccessoryViewControllers() INSSplitViewItemAccessoryViewController {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("bottomAlignedAccessoryViewControllers"))
	return NSSplitViewItemAccessoryViewControllerFromID(objc.ID(rv))
}
func (s NSSplitViewItemAccessoryViewController) SetBottomAlignedAccessoryViewControllers(value INSSplitViewItemAccessoryViewController) {
	objc.Send[struct{}](s.ID, objc.Sel("setBottomAlignedAccessoryViewControllers:"), value)
}



// The following methods allow you to add accessory views to the top/bottom of
// this splitViewItem. See
//
// See: https://developer.apple.com/documentation/appkit/nssplitviewitem/topalignedaccessoryviewcontrollers
func (s NSSplitViewItemAccessoryViewController) TopAlignedAccessoryViewControllers() INSSplitViewItemAccessoryViewController {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("topAlignedAccessoryViewControllers"))
	return NSSplitViewItemAccessoryViewControllerFromID(objc.ID(rv))
}
func (s NSSplitViewItemAccessoryViewController) SetTopAlignedAccessoryViewControllers(value INSSplitViewItemAccessoryViewController) {
	objc.Send[struct{}](s.ID, objc.Sel("setTopAlignedAccessoryViewControllers:"), value)
}


































