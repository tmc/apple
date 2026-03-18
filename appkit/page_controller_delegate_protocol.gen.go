// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)
var _ = fmt.Sprintf

// The [NSPageControllerDelegate] protocol allows you to customize the behavior of instances of the NSPageController class.
//
// See: https://developer.apple.com/documentation/AppKit/NSPageControllerDelegate
type NSPageControllerDelegate interface {
	objectivec.IObject
}

// NSPageControllerDelegateObject wraps an existing Objective-C object that conforms to the NSPageControllerDelegate protocol.
type NSPageControllerDelegateObject struct {
	objectivec.Object
}
func (o NSPageControllerDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSPageControllerDelegateObjectFromID constructs a [NSPageControllerDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSPageControllerDelegateObjectFromID(id objc.ID) NSPageControllerDelegateObject {
	return NSPageControllerDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// This message is sent when the user begins a transition.
//
// pageController: The page controller.
//
// # Discussion
// 
// This message is sent when the user begins a transition whether via a swipe
// gesture of one of the page controller’s target-action navigation methods.
//
// See: https://developer.apple.com/documentation/AppKit/NSPageControllerDelegate/pageControllerWillStartLiveTransition(_:)

func (o NSPageControllerDelegateObject) PageControllerWillStartLiveTransition(pageController INSPageController) {
	
	objc.Send[struct{}](o.ID, objc.Sel("pageControllerWillStartLiveTransition:"), pageController)
	}

// This message is sent when a transition animation completes.
//
// pageController: The page controller.
//
// # Discussion
// 
// This message is sent when a transition animation completes either via swipe
// gesture or one of the page controller’s target-action navigation methods.
// 
// Your content view is still hidden and you must call the
// [CompleteTransition] method on `pageController` when your content is ready
// to show.
// 
// If completed successfully, a [PageControllerDidTransitionToObject] will
// already have been sent.
//
// See: https://developer.apple.com/documentation/AppKit/NSPageControllerDelegate/pageControllerDidEndLiveTransition(_:)

func (o NSPageControllerDelegateObject) PageControllerDidEndLiveTransition(pageController INSPageController) {
	
	objc.Send[struct{}](o.ID, objc.Sel("pageControllerDidEndLiveTransition:"), pageController)
	}

// This message is sent when any page transition is completed.
//
// pageController: The page controller.
//
// object: The object to display.
//
// See: https://developer.apple.com/documentation/AppKit/NSPageControllerDelegate/pageController(_:didTransitionTo:)

func (o NSPageControllerDelegateObject) PageControllerDidTransitionToObject(pageController INSPageController, object objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("pageController:didTransitionToObject:"), pageController, object)
	}

// Return the identifier of the view controller that owns a view to display
// the object.
//
// pageController: The page controller.
//
// object: The object to display.
//
// # Return Value
// 
// Returns a string identifier for the view controller for the specified
// object.
//
// # Discussion
// 
// If `pageController` does not have an unused view controller for this
// identifier, the you will be asked to create one via
// [PageControllerViewControllerForIdentifier].
//
// See: https://developer.apple.com/documentation/AppKit/NSPageControllerDelegate/pageController(_:identifierFor:)

func (o NSPageControllerDelegateObject) PageControllerIdentifierForObject(pageController INSPageController, object objectivec.IObject) NSPageControllerObjectIdentifier {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("pageController:identifierForObject:"), pageController, object)
	return NSPageControllerObjectIdentifier(foundation.NSStringFromID(rv).String())
	}

// Returns a view controller the page controller uses for managing the
// specified identifier.
//
// pageController: The page controller.
//
// identifier: The identifier for a view controller.
//
// # Return Value
// 
// Returns the view controller for the specified identifier.
//
// # Discussion
// 
// Your implementation of this method should return the requested view
// controller for the identifier or create and return a new view controller.
// 
// [NSPageController] will cache as many view controllers and views as
// necessary to maintain performance. This method is called whenever another
// instance is required.
// 
// The view controller may become the [SelectedViewController] after a
// transition if necessary.
//
// See: https://developer.apple.com/documentation/AppKit/NSPageControllerDelegate/pageController(_:viewControllerForIdentifier:)

func (o NSPageControllerDelegateObject) PageControllerViewControllerForIdentifier(pageController INSPageController, identifier NSPageControllerObjectIdentifier) INSViewController {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("pageController:viewControllerForIdentifier:"), pageController, objc.String(string(identifier)))
	return NSViewControllerFromID(rv)
	}

// Prepare the view controller and it’s view for drawing.
//
// pageController: The page controller.
//
// viewController: The view controller to prepare for drawing. You should setup the data
// sources and perform layout.
//
// object: The object to display.
//
// # Discussion
// 
// If this method is not implemented, then `viewController` object’s
// `representedObject` is set to the object.
// 
// This method is only useful if [PageControllerIdentifierForObject] and
// [PageControllerPrepareViewControllerWithObject] are implemented.
//
// See: https://developer.apple.com/documentation/AppKit/NSPageControllerDelegate/pageController(_:prepare:with:)

func (o NSPageControllerDelegateObject) PageControllerPrepareViewControllerWithObject(pageController INSPageController, viewController INSViewController, object objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("pageController:prepareViewController:withObject:"), pageController, viewController, object)
	}

// Returns the frame appropriate for displaying the specified object.
//
// pageController: The page controller.
//
// object: The object to display.
//
// # Return Value
// 
// The frame appropriate for displaying `object`.
//
// # Discussion
// 
// You only need to implement this if the view frame can differ between the
// page controller’s [ArrangedObjects].
// 
// This method must return immediately. Avoid file, network or any potentially
// blocking or lengthy work to provide an answer.
// 
// If this method is not implemented, all [ArrangedObjects] are assumed to
// have the same frame as the `pageController` object’s current
// [SelectedViewController] instance’s `view` or the bounds of `view` when
// [SelectedViewController] is `nil`.
// 
// This method is only useful if [PageControllerIdentifierForObject] and
// [PageControllerViewControllerForIdentifier] are implemented.
//
// See: https://developer.apple.com/documentation/AppKit/NSPageControllerDelegate/pageController(_:frameFor:)

func (o NSPageControllerDelegateObject) PageControllerFrameForObject(pageController INSPageController, object objectivec.IObject) corefoundation.CGRect {
	
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("pageController:frameForObject:"), pageController, object)
	return rv
	}

// NSPageControllerDelegateConfig holds optional typed callbacks for [NSPageControllerDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nspagecontrollerdelegate
type NSPageControllerDelegateConfig struct {

	// Transition Notification
	// PageControllerWillStartLiveTransition — This message is sent when the user begins a transition.
	PageControllerWillStartLiveTransition func(pageController NSPageController)
	// PageControllerDidEndLiveTransition — This message is sent when a transition animation completes.
	PageControllerDidEndLiveTransition func(pageController NSPageController)

	// Managing View Controllers
	// PageControllerViewControllerForIdentifier — Returns a view controller the page controller uses for managing the specified identifier.
	PageControllerViewControllerForIdentifier func(pageController NSPageController, identifier NSPageControllerObjectIdentifier) NSViewController
}

// NewNSPageControllerDelegate creates an Objective-C object implementing the [NSPageControllerDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSPageControllerDelegateObject] satisfies the [NSPageControllerDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nspagecontrollerdelegate
func NewNSPageControllerDelegate(config NSPageControllerDelegateConfig) NSPageControllerDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSPageControllerDelegate_%d", n)

	var methods []objc.MethodDef

	if config.PageControllerWillStartLiveTransition != nil {
		fn := config.PageControllerWillStartLiveTransition
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("pageControllerWillStartLiveTransition:"),
			Fn: func(self objc.ID, _cmd objc.SEL, pageControllerID objc.ID) {
				pageController := NSPageControllerFromID(pageControllerID)
				fn(pageController)
			},
		})
	}

	if config.PageControllerDidEndLiveTransition != nil {
		fn := config.PageControllerDidEndLiveTransition
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("pageControllerDidEndLiveTransition:"),
			Fn: func(self objc.ID, _cmd objc.SEL, pageControllerID objc.ID) {
				pageController := NSPageControllerFromID(pageControllerID)
				fn(pageController)
			},
		})
	}

	if config.PageControllerViewControllerForIdentifier != nil {
		fn := config.PageControllerViewControllerForIdentifier
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("pageController:viewControllerForIdentifier:"),
			Fn: func(self objc.ID, _cmd objc.SEL, pageControllerID objc.ID, identifierID objc.ID) objc.ID {
				pageController := NSPageControllerFromID(pageControllerID)
				identifier := NSPageControllerObjectIdentifier(objc.GoString(objc.Send[*byte](identifierID, objc.Sel("UTF8String"))))
				return fn(pageController, identifier).GetID()
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSPageControllerDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSPageControllerDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSPageControllerDelegateObjectFromID(instance)
}

