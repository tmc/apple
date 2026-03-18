// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// A set of methods that defines a way to add animation to an existing class with a minimum of API impact.
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimatablePropertyContainer
type NSAnimatablePropertyContainer interface {
	objectivec.IObject

	// Returns a proxy object for the receiver that can be used to initiate implied animation for property changes.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAnimatablePropertyContainer/animator()
	Animator() objectivec.IObject

	// Sets the option dictionary that maps event trigger keys to animation objects.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAnimatablePropertyContainer/animations
	Animations() foundation.INSDictionary

	// Sets the option dictionary that maps event trigger keys to animation objects.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAnimatablePropertyContainer/animations
	SetAnimations(value foundation.INSDictionary)
}

// NSAnimatablePropertyContainerObject wraps an existing Objective-C object that conforms to the NSAnimatablePropertyContainer protocol.
type NSAnimatablePropertyContainerObject struct {
	objectivec.Object
}
func (o NSAnimatablePropertyContainerObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSAnimatablePropertyContainerObjectFromID constructs a [NSAnimatablePropertyContainerObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSAnimatablePropertyContainerObjectFromID(id objc.ID) NSAnimatablePropertyContainerObject {
	return NSAnimatablePropertyContainerObject{
		Object: objectivec.ObjectFromID(id),
	}
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

func (o NSAnimatablePropertyContainerObject) Animator() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("animator"))
	return objectivec.Object{ID: rv}
	}

// Sets the option dictionary that maps event trigger keys to animation
// objects.
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimatablePropertyContainer/animations

func (o NSAnimatablePropertyContainerObject) Animations() foundation.INSDictionary {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("animations"))
	return foundation.NSDictionaryFromID(rv)
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

func (o NSAnimatablePropertyContainerObject) AnimationForKey(key NSAnimatablePropertyKey) objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("animationForKey:"), objc.String(string(key)))
	return objectivec.Object{ID: rv}
	}

func (o NSAnimatablePropertyContainerObject) SetAnimations(value foundation.INSDictionary) {
	objc.Send[struct{}](o.ID, objc.Sel("setAnimations:"), value)
}

