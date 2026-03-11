// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSPageController] class.
var (
	_NSPageControllerClass     NSPageControllerClass
	_NSPageControllerClassOnce sync.Once
)

func getNSPageControllerClass() NSPageControllerClass {
	_NSPageControllerClassOnce.Do(func() {
		_NSPageControllerClass = NSPageControllerClass{class: objc.GetClass("NSPageController")}
	})
	return _NSPageControllerClass
}

// GetNSPageControllerClass returns the class object for NSPageController.
func GetNSPageControllerClass() NSPageControllerClass {
	return getNSPageControllerClass()
}

type NSPageControllerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPageControllerClass) Alloc() NSPageController {
	rv := objc.Send[NSPageController](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// An object that controls swipe navigation and animations between views or
// view content.
//
// # Overview
// 
// [NSPageController] is useful for user interfaces which control navigating
// multiple pages as in a book or a web browser history. Page controller
// inherits from the [NSViewController] class . You must assign the `view`
// property to a view in your view hierarchy. The [NSPageController] class
// does not vend a view and does insert itself into the responder chain.
// 
// Conceptually, the page controller manages swiping between an array of
// pages, the [NSPageController.ArrangedObjects]. Using the [NSPageController.SelectedIndex] property, you can
// determine how many pages forward or backward the user may navigate.
// 
// # Page Controller Modes
// 
// There are two modes that an [NSPageController] instance may operate in,
// history mode and book mode. The main difference between the two modes is
// that expects `pageController.View()` to be the content and expects
// `pageController.View()` to be be a container for the content that you will
// supply by returning `viewControllers` in your delegate methods.
// 
// # History Mode
// 
// History mode is designed to be the easiest way to create a history user
// interface. The page controller will manage the history (the
// `arrangedObjects` property), snapshots, and user navigation between pages
// in the history.
// 
// As the user navigates to new content, add to the history by calling
// [NSPageController.NavigateForwardToObject]. The page controller will remove any
// `arrangedObjects` after the [NSPageController.SelectedIndex] and then add the object to the
// end of the `arrangedObjects` array and update the `selectedIndex` property.
// Just like navigating in a new direction in a web browser, all forward
// history is lost once the user starts navigating a new path. After returning
// from `` you are free to update the contents of `pageController.View()`.
// 
// # Delegate Method Invocation During History Mode Swiping
// 
// During swiping, the following optional delegate methods are called in the
// specified order:
// 
// The [PageControllerWillStartLiveTransition] delegate method is invoked when
// the user starts a swipe action. This is the appropriate point at which to
// save information that you may need to restore, such as a page’s scrolled
// location.
// 
// Upon returning from the this delegate method, `pageController.View()` is
// hidden. In it’s place the page controller shows a private view hierarchy
// to animate previously taken snapshots of the page history. This allows the
// page controller to remain responsive to the user without any required
// action by your application.Next, if implemented, the
// [PageControllerDidTransitionToObject] delegate method is invoked. This
// delegate method is called after a physically successful swipe, but before
// the animation has completed. The supplied object is the page the user
// navigated to – the new `selectedIndex` object in `arrangedObjects`. If
// background loading tasks need to be initiated this is the appropriate time
// to do so. However, do not block the main thread or the animation will
// stutter or pause.
// 
// Finally, the pageControllerDidEndLiveTransition: delegate method is invoked
// after the swipe and swipe animations are complete. You should any position
// settings or other display specific state stored in the
// [PageControllerWillStartLiveTransition] implementation. The
// `pageController.View()` is still hidden at this point and you must call
// [NSPageController.CompleteTransition] on the page controller instance to inform the instance
// to hide the private transition view and show `pageController.View()`. Often
// you do this immediately, however, if your content is not ready you can call
// this at a later.
// 
// # Book Mode (View Controller Mode)
// 
// Book mode is designed to give you more control over the swiping process and
// to facilitate more user interface designs than just history, although you
// can use book mode to create a history user interface.
// 
// In this mode, `pageController.View()` is a container view and the content
// views are vended by view controller instances supplied by the delegate
// object.
// 
// To enable book mode, you must implement the following two methods in your
// delegate: [PageControllerIdentifierForObject] and
// [PageControllerViewControllerForIdentifier].
// 
// The page controller instance caches the view controllers supplied for each
// identifier and only asks it’s delegate to create more if one does not
// already exists in its cache. If you have different type of views you want
// to swipe in, then supply a different identifier for each type.
// 
// When needed, you will be asked to prepare a view controller instance with a
// page via the optional delegate method
// [PageControllerPrepareViewControllerWithObject]. If you do not implement
// this method, then the `representedObject` of the view controller that would
// have been passed to this delegate method is set as the object.
// 
// The delegate will be asked to prepare a view controller with a `nil` object
// for each unique identifier it encounters. The NSPageController instance
// will use this to generate a default snapshot for that identifier.
// 
// When using the book mode, if `pageController.View()` is layer backed, live
// layers are used during transition instead of snapshots.
// 
// Generally, when using book mode, the set of pages are known and it is your
// responsibility to set the `arrangedObjects` array property and initially
// selected page using the `selectedIndex` property.
// 
// # Delegate Method Invocation During Book Mode Swiping
// 
// During swiping, the following optional delegate methods are called in the
// specified order:
// 
// The delegate method [PageControllerWillStartLiveTransition] is invoked when
// the user starts a swipe action. As in history mode, this is the appropriate
// point at which to save information that you may need to restore, such as a
// page’s scrolled location.
// 
// After returning from the [PageControllerWillStartLiveTransition] delegate
// method, the page controller takes a snapshot of the `view` in the specified
// [NSPageController.SelectedViewController] and then removes it from `pageController.View()`.
// The page controller replaces it with a private view hierarchy to animate
// previously taken snapshots. Unlike when building up a history, snapshots
// may not yet exist for the page being navigated to. In this case, a
// previously gathered default snapshot is used for that page’s identifier.
// Regardless, if using a default snapshot or a previously gathered snapshot
// of actual contents, a view controller is prepared for the page being
// navigated to, and passed to the delegate. This `viewController.View()` is
// then asked to draw on a background thread while swiping continues. Note
// that at this point the view does not reside within a window. Once the
// background threaded drawing completes, the initial snapshot is replaced
// with the newly generated snapshot.
// 
// Next the [PageControllerDidTransitionToObject] delegate method is invoked
// after a physically successful swipe, but before the animation has
// completed. The supplied object is the page the user navigated to - the new
// object in the [NSPageController.ArrangedObjects] array at the [NSPageController.SelectedIndex]. Note that the
// page controller’s [NSPageController.SelectedViewController] has not been updated yet. If
// you need to start some background loading tasks, now is the time to do it.
// Do not block the main thread or the animation will stutter or pause.
// 
// Finally the [PageControllerDidEndLiveTransition] method is invoked after
// the swipe and swipe animations are complete. The
// `selectedViewController.View()` is still detached at this point and you
// must call [NSPageController.CompleteTransition] on the page controller to hide the private
// transition view and update the [NSPageController.SelectedViewController]. Often you do this
// immediately, however, if your content is not ready you can call this at a
// later.
// 
// # Completing the Page Controller Transition
// 
// An [NSPageController] instance uses a private view hierarchy during
// swiping. To create a seamless transition to the new content, it is your
// responsibility to inform the page controller when you are ready to draw the
// new content. Ideally, the new content should match the snapshot so the user
// is none the wiser. You inform the page controller to complete the
// transition by calling [NSPageController.CompleteTransition]. If needed, a view controller is
// prepared and then the content view is shown (or added) to the view
// hierarchy and the private transition view is hidden.
// 
// During page controller initiated animations,
// [PageControllerWillStartLiveTransition] and
// [PageControllerDidEndLiveTransition] are invoked on the delegate. Generally
// during [PageControllerDidEndLiveTransition] you will call
// [NSPageController.CompleteTransition]. Programatic animations via the animator proxy do not
// call the delegate methods and you are responsible for calling
// [NSPageController.CompleteTransition] when the animation completes.This is easily done via a
// completion handler on an [NSAnimationContext] grouping. For example:
//
// # Customizing the Paged Interface Behavior
//
//   - [NSPageController.Delegate]: The page controller’s delegate object.
//   - [NSPageController.SetDelegate]
//
// # Page Controller Items
//
//   - [NSPageController.ArrangedObjects]: An array containing the objects displayed in the page controller’s view.
//   - [NSPageController.SetArrangedObjects]
//   - [NSPageController.NavigateForwardToObject]: Navigates to the specific object.
//   - [NSPageController.SelectedIndex]: The currently selected object in the arranged objects array.
//   - [NSPageController.SetSelectedIndex]
//
// # Page Controller Navigation
//
//   - [NSPageController.NavigateBack]: Navigates backwards in the page controller’s arranged objects array.
//   - [NSPageController.NavigateForward]: Navigates to the next object in the page controller’s arranged objects array, if appropriate.
//   - [NSPageController.TakeSelectedIndexFrom]: Navigates to the selected index, which is taken from the sender.
//
// # Transition Style
//
//   - [NSPageController.TransitionStyle]: The transition style the page controller uses when changing pages.
//   - [NSPageController.SetTransitionStyle]
//
// # Completing Page Transition
//
//   - [NSPageController.CompleteTransition]: Invoked when the page transition is completed.
//
// # Getting the View Controller
//
//   - [NSPageController.SelectedViewController]: The view controller associated with the selected object..
//
// See: https://developer.apple.com/documentation/AppKit/NSPageController
type NSPageController struct {
	NSViewController
}

// NSPageControllerFromID constructs a [NSPageController] from an objc.ID.
//
// An object that controls swipe navigation and animations between views or
// view content.
func NSPageControllerFromID(id objc.ID) NSPageController {
	return NSPageController{NSViewController: NSViewControllerFromID(id)}
}
// NOTE: NSPageController adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSPageController] class.
//
// # Customizing the Paged Interface Behavior
//
//   - [INSPageController.Delegate]: The page controller’s delegate object.
//   - [INSPageController.SetDelegate]
//
// # Page Controller Items
//
//   - [INSPageController.ArrangedObjects]: An array containing the objects displayed in the page controller’s view.
//   - [INSPageController.SetArrangedObjects]
//   - [INSPageController.NavigateForwardToObject]: Navigates to the specific object.
//   - [INSPageController.SelectedIndex]: The currently selected object in the arranged objects array.
//   - [INSPageController.SetSelectedIndex]
//
// # Page Controller Navigation
//
//   - [INSPageController.NavigateBack]: Navigates backwards in the page controller’s arranged objects array.
//   - [INSPageController.NavigateForward]: Navigates to the next object in the page controller’s arranged objects array, if appropriate.
//   - [INSPageController.TakeSelectedIndexFrom]: Navigates to the selected index, which is taken from the sender.
//
// # Transition Style
//
//   - [INSPageController.TransitionStyle]: The transition style the page controller uses when changing pages.
//   - [INSPageController.SetTransitionStyle]
//
// # Completing Page Transition
//
//   - [INSPageController.CompleteTransition]: Invoked when the page transition is completed.
//
// # Getting the View Controller
//
//   - [INSPageController.SelectedViewController]: The view controller associated with the selected object..
//
// See: https://developer.apple.com/documentation/AppKit/NSPageController
type INSPageController interface {
	INSViewController

	// Topic: Customizing the Paged Interface Behavior

	// The page controller’s delegate object.
	Delegate() NSPageControllerDelegate
	SetDelegate(value NSPageControllerDelegate)

	// Topic: Page Controller Items

	// An array containing the objects displayed in the page controller’s view.
	ArrangedObjects() foundation.INSArray
	SetArrangedObjects(value foundation.INSArray)
	// Navigates to the specific object.
	NavigateForwardToObject(object objectivec.IObject)
	// The currently selected object in the arranged objects array.
	SelectedIndex() int
	SetSelectedIndex(value int)

	// Topic: Page Controller Navigation

	// Navigates backwards in the page controller’s arranged objects array.
	NavigateBack(sender objectivec.IObject)
	// Navigates to the next object in the page controller’s arranged objects array, if appropriate.
	NavigateForward(sender objectivec.IObject)
	// Navigates to the selected index, which is taken from the sender.
	TakeSelectedIndexFrom(sender objectivec.IObject)

	// Topic: Transition Style

	// The transition style the page controller uses when changing pages.
	TransitionStyle() NSPageControllerTransitionStyle
	SetTransitionStyle(value NSPageControllerTransitionStyle)

	// Topic: Completing Page Transition

	// Invoked when the page transition is completed.
	CompleteTransition()

	// Topic: Getting the View Controller

	// The view controller associated with the selected object..
	SelectedViewController() INSViewController

	// Returns the animation that should be performed for the specified key.
	AnimationForKey(key NSAnimatablePropertyKey) objectivec.IObject
	// Returns a proxy object for the receiver that can be used to initiate implied animation for property changes.
	Animator() INSPageController
	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (p NSPageController) Init() NSPageController {
	rv := objc.Send[NSPageController](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPageController) Autorelease() NSPageController {
	rv := objc.Send[NSPageController](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPageController creates a new NSPageController instance.
func NewNSPageController() NSPageController {
	class := getNSPageControllerClass()
	rv := objc.Send[NSPageController](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/AppKit/NSViewController/init(coder:)
func NewPageControllerWithCoder(coder foundation.INSCoder) NSPageController {
	instance := getNSPageControllerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSPageControllerFromID(rv)
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
func NewPageControllerWithNibNameBundle(nibNameOrNil NSNibName, nibBundleOrNil foundation.NSBundle) NSPageController {
	instance := getNSPageControllerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithNibName:bundle:"), objc.String(string(nibNameOrNil)), nibBundleOrNil)
	return NSPageControllerFromID(rv)
}







// Navigates to the specific object.
//
// object: The object to display.
//
// # Discussion
// 
// Clears the [ArrangedObjects] array after the selected index, adds the
// argument to that array, and sets the [SelectedIndex] to the object’s
// index.
//
// See: https://developer.apple.com/documentation/AppKit/NSPageController/navigateForward(to:)
func (p NSPageController) NavigateForwardToObject(object objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("navigateForwardToObject:"), object)
}

// Navigates backwards in the page controller’s arranged objects array.
//
// sender: The sender.
//
// # Discussion
// 
// This method is typically invoked in response to a user interacting with a
// control, the `sender`.
// 
// This method is animated and invokes the delegate’s
// [PageControllerWillStartLiveTransition] and
// [PageControllerDidEndLiveTransition] methods.
//
// See: https://developer.apple.com/documentation/AppKit/NSPageController/navigateBack(_:)
func (p NSPageController) NavigateBack(sender objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("navigateBack:"), sender)
}

// Navigates to the next object in the page controller’s arranged objects
// array, if appropriate.
//
// sender: The sender.
//
// # Discussion
// 
// This method is typically invoked in response to a user interacting with a
// control, the `sender`.
// 
// This method is animated and invokes the delegate’s
// [PageControllerWillStartLiveTransition] and
// [PageControllerDidEndLiveTransition] methods.
//
// See: https://developer.apple.com/documentation/AppKit/NSPageController/navigateForward(_:)
func (p NSPageController) NavigateForward(sender objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("navigateForward:"), sender)
}

// Navigates to the selected index, which is taken from the sender.
//
// sender: The control that invoked the action.
//
// # Discussion
// 
// When invoked, this method causes the page controller’s view to display
// the object specified by the value taken from the `sender` control.
// 
// This method is animated and invokes the delegate’s
// [PageControllerWillStartLiveTransition] and
// [PageControllerDidEndLiveTransition] methods.
//
// See: https://developer.apple.com/documentation/AppKit/NSPageController/takeSelectedIndexFrom(_:)
func (p NSPageController) TakeSelectedIndexFrom(sender objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("takeSelectedIndexFrom:"), sender)
}

// Invoked when the page transition is completed.
//
// # Discussion
// 
// See [NSPageController] for a complete description.
//
// See: https://developer.apple.com/documentation/AppKit/NSPageController/completeTransition()
func (p NSPageController) CompleteTransition() {
	objc.Send[objc.ID](p.ID, objc.Sel("completeTransition"))
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
func (p NSPageController) AnimationForKey(key NSAnimatablePropertyKey) objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("animationForKey:"), objc.String(string(key)))
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
func (p NSPageController) Animator() INSPageController {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("animator"))
	return NSPageControllerFromID(rv)
}
func (p NSPageController) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](p.ID, objc.Sel("encodeWithCoder:"), coder)
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
func (_NSPageControllerClass NSPageControllerClass) DefaultAnimationForKey(key NSAnimatablePropertyKey) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_NSPageControllerClass.class), objc.Sel("defaultAnimationForKey:"), objc.String(string(key)))
	return objectivec.Object{ID: rv}
}








// The page controller’s delegate object.
//
// # Discussion
// 
// The delegate must conform to the [NSPageControllerDelegate] protocol.
//
// See: https://developer.apple.com/documentation/AppKit/NSPageController/delegate
func (p NSPageController) Delegate() NSPageControllerDelegate {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("delegate"))
	return NSPageControllerDelegateObjectFromID(rv)
}
func (p NSPageController) SetDelegate(value NSPageControllerDelegate) {
	objc.Send[struct{}](p.ID, objc.Sel("setDelegate:"), value)
}



// An array containing the objects displayed in the page controller’s view.
//
// # Discussion
// 
// The delegate will be asked for snapshots as they are needed. Alternatively,
// you may never directly set this array and use the
// -[NavigateForwardToObject] method to create a history as the user
// navigates.
// 
// This property is key-value observing compliant.
//
// See: https://developer.apple.com/documentation/AppKit/NSPageController/arrangedObjects
func (p NSPageController) ArrangedObjects() foundation.INSArray {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("arrangedObjects"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (p NSPageController) SetArrangedObjects(value foundation.INSArray) {
	objc.Send[struct{}](p.ID, objc.Sel("setArrangedObjects:"), value)
}



// The currently selected object in the arranged objects array.
//
// # Discussion
// 
// To animate a transition to a new index, use the NSPageController class
// animator object.
// 
// This property is key-value observing compliant.
//
// See: https://developer.apple.com/documentation/AppKit/NSPageController/selectedIndex
func (p NSPageController) SelectedIndex() int {
	rv := objc.Send[int](p.ID, objc.Sel("selectedIndex"))
	return rv
}
func (p NSPageController) SetSelectedIndex(value int) {
	objc.Send[struct{}](p.ID, objc.Sel("setSelectedIndex:"), value)
}



// The transition style the page controller uses when changing pages.
//
// # Discussion
// 
// The possible values for the transition style are discussed in
// [NSPageController.TransitionStyle].
// 
// The default value is [NSPageController.TransitionStyle.stackHistory].
//
// [NSPageController.TransitionStyle.stackHistory]: https://developer.apple.com/documentation/AppKit/NSPageController/TransitionStyle-swift.enum/stackHistory
// [NSPageController.TransitionStyle]: https://developer.apple.com/documentation/AppKit/NSPageController/TransitionStyle-swift.enum
//
// See: https://developer.apple.com/documentation/AppKit/NSPageController/transitionStyle-swift.property
func (p NSPageController) TransitionStyle() NSPageControllerTransitionStyle {
	rv := objc.Send[NSPageControllerTransitionStyle](p.ID, objc.Sel("transitionStyle"))
	return NSPageControllerTransitionStyle(rv)
}
func (p NSPageController) SetTransitionStyle(value NSPageControllerTransitionStyle) {
	objc.Send[struct{}](p.ID, objc.Sel("setTransitionStyle:"), value)
}



// The view controller associated with the selected object..
//
// # Discussion
// 
// May be `nil` if not using view controllers.
// 
// This property is only relevant in book mode. See [NSPageController] for
// details.
//
// See: https://developer.apple.com/documentation/AppKit/NSPageController/selectedViewController
func (p NSPageController) SelectedViewController() INSViewController {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("selectedViewController"))
	return NSViewControllerFromID(objc.ID(rv))
}



// Sets the option dictionary that maps event trigger keys to animation
// objects.
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimatablePropertyContainer/animations
func (p NSPageController) Animations() foundation.INSDictionary {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("animations"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (p NSPageController) SetAnimations(value foundation.INSDictionary) {
	objc.Send[struct{}](p.ID, objc.Sel("setAnimations:"), value)
}


































