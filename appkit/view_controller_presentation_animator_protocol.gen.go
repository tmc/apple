// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A set of methods that let you define animations to play when transitioning between two view controllers.
//
// See: https://developer.apple.com/documentation/AppKit/NSViewControllerPresentationAnimator
type NSViewControllerPresentationAnimator interface {
	objectivec.IObject

	// Called when the specified view controller is about to be presented.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSViewControllerPresentationAnimator/animatePresentation(of:from:)
	AnimatePresentationOfViewControllerFromViewController(viewController INSViewController, fromViewController INSViewController)

	// Called when a previously-presented view controller is about to be dismissed.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSViewControllerPresentationAnimator/animateDismissal(of:from:)
	AnimateDismissalOfViewControllerFromViewController(viewController INSViewController, fromViewController INSViewController)
}

// NSViewControllerPresentationAnimatorObject wraps an existing Objective-C object that conforms to the NSViewControllerPresentationAnimator protocol.
type NSViewControllerPresentationAnimatorObject struct {
	objectivec.Object
}
func (o NSViewControllerPresentationAnimatorObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSViewControllerPresentationAnimatorObjectFromID constructs a [NSViewControllerPresentationAnimatorObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSViewControllerPresentationAnimatorObjectFromID(id objc.ID) NSViewControllerPresentationAnimatorObject {
	return NSViewControllerPresentationAnimatorObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Called when the specified view controller is about to be presented.
//
// viewController: The view controller that is being presented in place of the one in the
// `fromViewController` parameter.
//
// fromViewController: The view controller that is the parent of the one in the `viewController`
// parameter.
//
// # Discussion
// 
// To add custom presentation animation, Implement it in this method.
//
// See: https://developer.apple.com/documentation/AppKit/NSViewControllerPresentationAnimator/animatePresentation(of:from:)

func (o NSViewControllerPresentationAnimatorObject) AnimatePresentationOfViewControllerFromViewController(viewController INSViewController, fromViewController INSViewController) {
	
	objc.Send[struct{}](o.ID, objc.Sel("animatePresentationOfViewController:fromViewController:"), viewController, fromViewController)
	}

// Called when a previously-presented view controller is about to be
// dismissed.
//
// viewController: The view controller that is being dismissed.
//
// fromViewController: The view controller that is the parent of the one in the `viewController`
// parameter.
//
// # Discussion
// 
// To add custom view controller dismissal animation, Implement it in this
// method.
//
// See: https://developer.apple.com/documentation/AppKit/NSViewControllerPresentationAnimator/animateDismissal(of:from:)

func (o NSViewControllerPresentationAnimatorObject) AnimateDismissalOfViewControllerFromViewController(viewController INSViewController, fromViewController INSViewController) {
	
	objc.Send[struct{}](o.ID, objc.Sel("animateDismissalOfViewController:fromViewController:"), viewController, fromViewController)
	}

