// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"fmt"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// A set of optional methods implemented by delegates of [NSAnimation](<doc://com.apple.appkit/documentation/AppKit/NSAnimation>) objects.
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimationDelegate
type NSAnimationDelegate interface {
	objectivec.IObject
}

// NSAnimationDelegateObject wraps an existing Objective-C object that conforms to the NSAnimationDelegate protocol.
type NSAnimationDelegateObject struct {
	objectivec.Object
}

func (o NSAnimationDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSAnimationDelegateObjectFromID constructs a [NSAnimationDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSAnimationDelegateObjectFromID(id objc.ID) NSAnimationDelegateObject {
	return NSAnimationDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
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
func (o NSAnimationDelegateObject) AnimationDidEnd(animation INSAnimation) {
	objc.Send[struct{}](o.ID, objc.Sel("animationDidEnd:"), animation)
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
func (o NSAnimationDelegateObject) AnimationDidStop(animation INSAnimation) {
	objc.Send[struct{}](o.ID, objc.Sel("animationDidStop:"), animation)
}

// Sent to the delegate just after an animation is started.
//
// animation: The [NSAnimation] object that was just started.
//
// # Return Value
//
// false to cancel the animation, true to have the animation proceed.
//
// # Discussion
//
// The delegate is sent this message just after `animation` receives a
// [StartAnimation] message. The delegate can use this method to prepare
// objects and resources for the effect.
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimationDelegate/animationShouldStart(_:)
func (o NSAnimationDelegateObject) AnimationShouldStart(animation INSAnimation) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("animationShouldStart:"), animation)
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
func (o NSAnimationDelegateObject) AnimationValueForProgress(animation INSAnimation, progress NSAnimationProgress) float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("animation:valueForProgress:"), animation, progress)
	return rv
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
// See: https://developer.apple.com/documentation/AppKit/NSAnimationDelegate/animation(_:didReachProgressMark:)
//
// [progressMarkNotification]: https://developer.apple.com/documentation/AppKit/NSAnimation/progressMarkNotification
func (o NSAnimationDelegateObject) AnimationDidReachProgressMark(animation INSAnimation, progress NSAnimationProgress) {
	objc.Send[struct{}](o.ID, objc.Sel("animation:didReachProgressMark:"), animation, progress)
}

// NSAnimationDelegateConfig holds optional typed callbacks for [NSAnimationDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nsanimationdelegate
type NSAnimationDelegateConfig struct {

	// Controlling and Monitoring an Animation
	// DidEnd — Sent to the delegate when the specified animation completes its run.
	DidEnd func(animation NSAnimation)
	// DidStop — Sent to the delegate when the specified animation is stopped before it completes its run.
	DidStop func(animation NSAnimation)
	// ShouldStart — Sent to the delegate just after an animation is started.
	ShouldStart func(animation NSAnimation) bool
	// ValueForProgress — Requests a custom curve value for the current progress value.
	ValueForProgress func(animation NSAnimation, progress NSAnimationProgress) float32

	// Managing Progress Marks
	// DidReachProgressMark — Sent to the delegate when an animation reaches a specific progress mark.
	DidReachProgressMark func(animation NSAnimation, progress NSAnimationProgress)
}

// NewNSAnimationDelegate creates an Objective-C object implementing the [NSAnimationDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSAnimationDelegateObject] satisfies the [NSAnimationDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nsanimationdelegate
func NewNSAnimationDelegate(config NSAnimationDelegateConfig) NSAnimationDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSAnimationDelegate_%d", n)

	var methods []objc.MethodDef

	if config.DidEnd != nil {
		fn := config.DidEnd
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("animationDidEnd:"),
			Fn: func(self objc.ID, _cmd objc.SEL, animationID objc.ID) {
				animation := NSAnimationFromID(animationID)
				fn(animation)
			},
		})
	}

	if config.DidStop != nil {
		fn := config.DidStop
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("animationDidStop:"),
			Fn: func(self objc.ID, _cmd objc.SEL, animationID objc.ID) {
				animation := NSAnimationFromID(animationID)
				fn(animation)
			},
		})
	}

	if config.ShouldStart != nil {
		fn := config.ShouldStart
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("animationShouldStart:"),
			Fn: func(self objc.ID, _cmd objc.SEL, animationID objc.ID) bool {
				animation := NSAnimationFromID(animationID)
				return fn(animation)
			},
		})
	}

	if config.ValueForProgress != nil {
		fn := config.ValueForProgress
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("animation:valueForProgress:"),
			Fn: func(self objc.ID, _cmd objc.SEL, animationID objc.ID, progress NSAnimationProgress) float32 {
				animation := NSAnimationFromID(animationID)
				return fn(animation, progress)
			},
		})
	}

	if config.DidReachProgressMark != nil {
		fn := config.DidReachProgressMark
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("animation:didReachProgressMark:"),
			Fn: func(self objc.ID, _cmd objc.SEL, animationID objc.ID, progress NSAnimationProgress) {
				animation := NSAnimationFromID(animationID)
				fn(animation, progress)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSAnimationDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSAnimationDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSAnimationDelegateObjectFromID(instance)
}
