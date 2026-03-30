// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"fmt"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// A set of methods for fine-tuning a gesture recognizer’s behavior.
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizerDelegate
type NSGestureRecognizerDelegate interface {
	objectivec.IObject
}

// NSGestureRecognizerDelegateObject wraps an existing Objective-C object that conforms to the NSGestureRecognizerDelegate protocol.
type NSGestureRecognizerDelegateObject struct {
	objectivec.Object
}

func (o NSGestureRecognizerDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSGestureRecognizerDelegateObjectFromID constructs a [NSGestureRecognizerDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSGestureRecognizerDelegateObjectFromID(id objc.ID) NSGestureRecognizerDelegateObject {
	return NSGestureRecognizerDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Asks the delegate if a gesture recognizer should attempt to recognize
// gestures for a particular event.
//
// gestureRecognizer: The gesture recognizer object that is interpreting events. This is the
// object with which the delegate is associated.
//
// event: An event object associated with the request.
//
// # Return Value
//
// true to allow the gesture recognizer to begin recognizing gestures for the
// specified event, or false to prevent it from recognizing gestures for the
// specified event. If you do not implement this method, the default return
// value is true.
//
// # Discussion
//
// This method is called when a target view recognizes a new gesture event
// stream. The target view calls this method to determine whether the gesture
// recognizer should process events for the stream, or opt out of them.
// Returning false from this method causes the gesture recognizer to opt out,
// and prevents the other delegate methods from being called for the event
// stream.
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizerDelegate/gestureRecognizer(_:shouldAttemptToRecognizeWith:)
func (o NSGestureRecognizerDelegateObject) GestureRecognizerShouldAttemptToRecognizeWithEvent(gestureRecognizer INSGestureRecognizer, event INSEvent) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("gestureRecognizer:shouldAttemptToRecognizeWithEvent:"), gestureRecognizer, event)
	return rv
}

// Asks the delegate if a gesture recognizer should transition out of the
// Possible ([NSGestureRecognizerStatePossible]) state.
//
// gestureRecognizer: The gesture recognizer object that is interpreting events. This is the
// object with which the delegate is associated.
//
// # Return Value
//
// true to let the gesture recognizer transition out of the Possible
// ([NSGestureRecognizerStatePossible]) and continue trying to recognize the
// gesture or false to prevent it from trying to recognize its gesture. If you
// do not implement this method, the default return value is true.
//
// # Discussion
//
// When a gesture recognizer attempts to transition from the Possible
// ([NSGestureRecognizerStatePossible]) state to a different state, such as
// [NSGestureRecognizerStateBegan], the gesture recognizer calls this method
// to see if the transition should occur. Returning false from this delegate
// method causes the gesture recognizer to transition to the
// [NSGestureRecognizerStateFailed] state.
//
// For information about gesture states and transitions, see
// [NSGestureRecognizer] in [NSGestureRecognizer].
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizerDelegate/gestureRecognizerShouldBegin(_:)
func (o NSGestureRecognizerDelegateObject) GestureRecognizerShouldBegin(gestureRecognizer INSGestureRecognizer) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("gestureRecognizerShouldBegin:"), gestureRecognizer)
	return rv
}

// Asks the delegate if two gesture recognizers should be allowed to recognize
// their gestures simultaneously.
//
// gestureRecognizer: The first gesture recognizer to be considered. This is the object with
// which the delegate is associated.
//
// otherGestureRecognizer: The second gesture recognizer to be considered.
//
// # Return Value
//
// true to allow `gestureRecognizer` and `otherGestureRecognizer` to recognize
// their gestures simultaneously. If you do not implement this method, the
// default return value is false—no two gestures can be recognized
// simultaneously.
//
// # Discussion
//
// This method is called when recognition of a gesture by either
// `gestureRecognizer` and `otherGestureRecognizer` would block the other
// gesture recognizer from recognizing its gesture. Returning true is
// guaranteed to allow simultaneous recognition; returning false is not
// guaranteed to prevent simultaneous recognition because the other gesture
// recognizer’s delegate may return true.
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizerDelegate/gestureRecognizer(_:shouldRecognizeSimultaneouslyWith:)
func (o NSGestureRecognizerDelegateObject) GestureRecognizerShouldRecognizeSimultaneouslyWithGestureRecognizer(gestureRecognizer INSGestureRecognizer, otherGestureRecognizer INSGestureRecognizer) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("gestureRecognizer:shouldRecognizeSimultaneouslyWithGestureRecognizer:"), gestureRecognizer, otherGestureRecognizer)
	return rv
}

// Asks the delegate if the current gesture recognizer must wait to recognize
// its gesture until the specified gesture recognizer fails.
//
// gestureRecognizer: The gesture recognizer that might need to wait to recognize its gesture.
// This is the object with which the delegate is associated.
//
// otherGestureRecognizer: The gesture recognizer that must fail before the object in
// `gestureRecognizer` can recognize its gesture.
//
// # Return Value
//
// true if `otherGestureRecognizer` must fail before `gestureRecognizer` is
// allowed to recognize its gesture. If you do not implement this method, the
// default return value is false.
//
// # Discussion
//
// This method is called once per attempt to recognize, so you can change the
// failure requirements dynamically. The two gesture recognizers do not have
// to belong to the same view hierarchy.
//
// Returning true is guaranteed to set up the failure requirement; returning
// false does not prevent the failure requirement from being set up by the
// other gesture recognizer.
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizerDelegate/gestureRecognizer(_:shouldRequireFailureOf:)
func (o NSGestureRecognizerDelegateObject) GestureRecognizerShouldRequireFailureOfGestureRecognizer(gestureRecognizer INSGestureRecognizer, otherGestureRecognizer INSGestureRecognizer) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("gestureRecognizer:shouldRequireFailureOfGestureRecognizer:"), gestureRecognizer, otherGestureRecognizer)
	return rv
}

// Asks the delegate if the current gesture recognizer must fail before
// another gesture recognizer is allowed to recognize its gesture.
//
// gestureRecognizer: The gesture recognizer that must fail before `otherGestureRecognizer` can
// recognize its gesture. This is the object with which the delegate is
// associated.
//
// otherGestureRecognizer: The gesture recognizer that might need to wait to recognize its gesture.
//
// # Return Value
//
// true if `gestureRecognizer` must fail before `otherGestureRecognizer` is
// allowed to recognize its gesture. If you do not implement this method, the
// default return value is false.
//
// # Discussion
//
// This method is called once per attempt to recognize, so you can change the
// failure requirements dynamically. The two gesture recognizers do not have
// to belong to the same view hierarchy.
//
// Returning true is guaranteed to set up the failure requirement; returning
// false does not prevent the failure requirement from being set up by the
// other gesture recognizer.
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizerDelegate/gestureRecognizer(_:shouldBeRequiredToFailBy:)
func (o NSGestureRecognizerDelegateObject) GestureRecognizerShouldBeRequiredToFailByGestureRecognizer(gestureRecognizer INSGestureRecognizer, otherGestureRecognizer INSGestureRecognizer) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("gestureRecognizer:shouldBeRequiredToFailByGestureRecognizer:"), gestureRecognizer, otherGestureRecognizer)
	return rv
}

// Called, for a new touch, before the system calls the “ method on the
// gesture recognizer. Return [NO] to prevent the gesture recognizer from
// seeing this touch.
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizerDelegate/gestureRecognizer(_:shouldReceive:)
func (o NSGestureRecognizerDelegateObject) GestureRecognizerShouldReceiveTouch(gestureRecognizer INSGestureRecognizer, touch INSTouch) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("gestureRecognizer:shouldReceiveTouch:"), gestureRecognizer, touch)
	return rv
}

// NSGestureRecognizerDelegateConfig holds optional typed callbacks for [NSGestureRecognizerDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nsgesturerecognizerdelegate
type NSGestureRecognizerDelegateConfig struct {

	// Regulating Gesture Recognition
	// ShouldBegin — Asks the delegate if a gesture recognizer should transition out of the Possible ([NSGestureRecognizerStatePossible]) state.
	ShouldBegin func(gestureRecognizer NSGestureRecognizer) bool

	// Other Methods
	// ShouldAttemptToRecognizeWithEvent — Asks the delegate if a gesture recognizer should attempt to recognize gestures for a particular event.
	ShouldAttemptToRecognizeWithEvent func(gestureRecognizer NSGestureRecognizer, event NSEvent) bool
	// ShouldRecognizeSimultaneouslyWithGestureRecognizer — Asks the delegate if two gesture recognizers should be allowed to recognize their gestures simultaneously.
	ShouldRecognizeSimultaneouslyWithGestureRecognizer func(gestureRecognizer NSGestureRecognizer, otherGestureRecognizer NSGestureRecognizer) bool
	// ShouldRequireFailureOfGestureRecognizer — Asks the delegate if the current gesture recognizer must wait to recognize its gesture until the specified gesture recognizer fails.
	ShouldRequireFailureOfGestureRecognizer func(gestureRecognizer NSGestureRecognizer, otherGestureRecognizer NSGestureRecognizer) bool
	// ShouldBeRequiredToFailByGestureRecognizer — Asks the delegate if the current gesture recognizer must fail before another gesture recognizer is allowed to recognize its gesture.
	ShouldBeRequiredToFailByGestureRecognizer func(gestureRecognizer NSGestureRecognizer, otherGestureRecognizer NSGestureRecognizer) bool
	// ShouldReceiveTouch — Called, for a new touch, before the system calls the `` method on the gesture recognizer. Return [NO] to prevent the gesture recognizer from seeing this touch.
	ShouldReceiveTouch func(gestureRecognizer NSGestureRecognizer, touch NSTouch) bool
}

// NewNSGestureRecognizerDelegate creates an Objective-C object implementing the [NSGestureRecognizerDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSGestureRecognizerDelegateObject] satisfies the [NSGestureRecognizerDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nsgesturerecognizerdelegate
func NewNSGestureRecognizerDelegate(config NSGestureRecognizerDelegateConfig) NSGestureRecognizerDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSGestureRecognizerDelegate_%d", n)

	var methods []objc.MethodDef

	if config.ShouldAttemptToRecognizeWithEvent != nil {
		fn := config.ShouldAttemptToRecognizeWithEvent
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("gestureRecognizer:shouldAttemptToRecognizeWithEvent:"),
			Fn: func(self objc.ID, _cmd objc.SEL, gestureRecognizerID objc.ID, eventID objc.ID) bool {
				gestureRecognizer := NSGestureRecognizerFromID(gestureRecognizerID)
				event := NSEventFromID(eventID)
				return fn(gestureRecognizer, event)
			},
		})
	}

	if config.ShouldBegin != nil {
		fn := config.ShouldBegin
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("gestureRecognizerShouldBegin:"),
			Fn: func(self objc.ID, _cmd objc.SEL, gestureRecognizerID objc.ID) bool {
				gestureRecognizer := NSGestureRecognizerFromID(gestureRecognizerID)
				return fn(gestureRecognizer)
			},
		})
	}

	if config.ShouldRecognizeSimultaneouslyWithGestureRecognizer != nil {
		fn := config.ShouldRecognizeSimultaneouslyWithGestureRecognizer
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("gestureRecognizer:shouldRecognizeSimultaneouslyWithGestureRecognizer:"),
			Fn: func(self objc.ID, _cmd objc.SEL, gestureRecognizerID objc.ID, otherGestureRecognizerID objc.ID) bool {
				gestureRecognizer := NSGestureRecognizerFromID(gestureRecognizerID)
				otherGestureRecognizer := NSGestureRecognizerFromID(otherGestureRecognizerID)
				return fn(gestureRecognizer, otherGestureRecognizer)
			},
		})
	}

	if config.ShouldRequireFailureOfGestureRecognizer != nil {
		fn := config.ShouldRequireFailureOfGestureRecognizer
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("gestureRecognizer:shouldRequireFailureOfGestureRecognizer:"),
			Fn: func(self objc.ID, _cmd objc.SEL, gestureRecognizerID objc.ID, otherGestureRecognizerID objc.ID) bool {
				gestureRecognizer := NSGestureRecognizerFromID(gestureRecognizerID)
				otherGestureRecognizer := NSGestureRecognizerFromID(otherGestureRecognizerID)
				return fn(gestureRecognizer, otherGestureRecognizer)
			},
		})
	}

	if config.ShouldBeRequiredToFailByGestureRecognizer != nil {
		fn := config.ShouldBeRequiredToFailByGestureRecognizer
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("gestureRecognizer:shouldBeRequiredToFailByGestureRecognizer:"),
			Fn: func(self objc.ID, _cmd objc.SEL, gestureRecognizerID objc.ID, otherGestureRecognizerID objc.ID) bool {
				gestureRecognizer := NSGestureRecognizerFromID(gestureRecognizerID)
				otherGestureRecognizer := NSGestureRecognizerFromID(otherGestureRecognizerID)
				return fn(gestureRecognizer, otherGestureRecognizer)
			},
		})
	}

	if config.ShouldReceiveTouch != nil {
		fn := config.ShouldReceiveTouch
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("gestureRecognizer:shouldReceiveTouch:"),
			Fn: func(self objc.ID, _cmd objc.SEL, gestureRecognizerID objc.ID, touchID objc.ID) bool {
				gestureRecognizer := NSGestureRecognizerFromID(gestureRecognizerID)
				touch := NSTouchFromID(touchID)
				return fn(gestureRecognizer, touch)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSGestureRecognizerDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSGestureRecognizerDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSGestureRecognizerDelegateObjectFromID(instance)
}
