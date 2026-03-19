// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/quartzcore"
)

// The class instance for the [NSAnimationContext] class.
var (
	_NSAnimationContextClass     NSAnimationContextClass
	_NSAnimationContextClassOnce sync.Once
)

func getNSAnimationContextClass() NSAnimationContextClass {
	_NSAnimationContextClassOnce.Do(func() {
		_NSAnimationContextClass = NSAnimationContextClass{class: objc.GetClass("NSAnimationContext")}
	})
	return _NSAnimationContextClass
}

// GetNSAnimationContextClass returns the class object for NSAnimationContext.
func GetNSAnimationContextClass() NSAnimationContextClass {
	return getNSAnimationContextClass()
}

type NSAnimationContextClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSAnimationContextClass) Alloc() NSAnimationContext {
	rv := objc.Send[NSAnimationContext](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An animation context, which contains information about environment and
// state.
//
// # Overview
// 
// [NSAnimationContext] is analogous to [CATransaction] and is similar in
// overall concept to [NSGraphicsContext]. Each thread maintains its own stack
// of nestable [NSAnimationContext] instances, with each new instance
// initialized as a copy of the instance below (so, inheriting its current
// properties).
// 
// Multiple [NSAnimationContext] instances can be nested, allowing a given
// block of code to initiate animations using its own specified duration
// without affecting animations initiated by surrounding code.
//
// [CATransaction]: https://developer.apple.com/documentation/QuartzCore/CATransaction
//
// # Animation Completion Handlers
//
//   - [NSAnimationContext.CompletionHandler]: A completion Block that is called when the animations in the grouping are completed.
//   - [NSAnimationContext.SetCompletionHandler]
//
// # Modifying the Animation Duration
//
//   - [NSAnimationContext.Duration]: The duration used by animations created as a result of setting new values for an animatable property.
//   - [NSAnimationContext.SetDuration]
//   - [NSAnimationContext.TimingFunction]: The timing function used for all animations within this animation proxy group.
//   - [NSAnimationContext.SetTimingFunction]
//
// # Implicit Animation
//
//   - [NSAnimationContext.AllowsImplicitAnimation]: Determine if animations are enabled or not for animations that occur as a result of another property change.
//   - [NSAnimationContext.SetAllowsImplicitAnimation]
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimationContext
type NSAnimationContext struct {
	objectivec.Object
}

// NSAnimationContextFromID constructs a [NSAnimationContext] from an objc.ID.
//
// An animation context, which contains information about environment and
// state.
func NSAnimationContextFromID(id objc.ID) NSAnimationContext {
	return NSAnimationContext{objectivec.Object{ID: id}}
}
// NOTE: NSAnimationContext adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSAnimationContext] class.
//
// # Animation Completion Handlers
//
//   - [INSAnimationContext.CompletionHandler]: A completion Block that is called when the animations in the grouping are completed.
//   - [INSAnimationContext.SetCompletionHandler]
//
// # Modifying the Animation Duration
//
//   - [INSAnimationContext.Duration]: The duration used by animations created as a result of setting new values for an animatable property.
//   - [INSAnimationContext.SetDuration]
//   - [INSAnimationContext.TimingFunction]: The timing function used for all animations within this animation proxy group.
//   - [INSAnimationContext.SetTimingFunction]
//
// # Implicit Animation
//
//   - [INSAnimationContext.AllowsImplicitAnimation]: Determine if animations are enabled or not for animations that occur as a result of another property change.
//   - [INSAnimationContext.SetAllowsImplicitAnimation]
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimationContext
type INSAnimationContext interface {
	objectivec.IObject

	// Topic: Animation Completion Handlers

	// A completion Block that is called when the animations in the grouping are completed.
	CompletionHandler() VoidHandler
	SetCompletionHandler(value VoidHandler)

	// Topic: Modifying the Animation Duration

	// The duration used by animations created as a result of setting new values for an animatable property.
	Duration() float64
	SetDuration(value float64)
	// The timing function used for all animations within this animation proxy group.
	TimingFunction() quartzcore.CAMediaTimingFunction
	SetTimingFunction(value quartzcore.CAMediaTimingFunction)

	// Topic: Implicit Animation

	// Determine if animations are enabled or not for animations that occur as a result of another property change.
	AllowsImplicitAnimation() bool
	SetAllowsImplicitAnimation(value bool)
}

// Init initializes the instance.
func (a NSAnimationContext) Init() NSAnimationContext {
	rv := objc.Send[NSAnimationContext](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a NSAnimationContext) Autorelease() NSAnimationContext {
	rv := objc.Send[NSAnimationContext](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSAnimationContext creates a new NSAnimationContext instance.
func NewNSAnimationContext() NSAnimationContext {
	class := getNSAnimationContextClass()
	rv := objc.Send[NSAnimationContext](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new animation grouping.
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimationContext/beginGrouping()
func (_NSAnimationContextClass NSAnimationContextClass) BeginGrouping() {
	objc.Send[objc.ID](objc.ID(_NSAnimationContextClass.class), objc.Sel("beginGrouping"))
}
// Ends the current animation grouping.
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimationContext/endGrouping()
func (_NSAnimationContextClass NSAnimationContextClass) EndGrouping() {
	objc.Send[objc.ID](objc.ID(_NSAnimationContextClass.class), objc.Sel("endGrouping"))
}
// Allows you to specify a completion block body after the set of animation
// actions whose completion will trigger the completion block.
//
// changes: A block object containing animations for this transaction group.
// 
// The `context` parameter passes the thread’s current [NSAnimationContext]
// to the Block as a convenience, so code within the Block that wants to
// change or query properties of the current `context` does not have to call
// [CurrentContext].
// 
// The block object returns no value.
//
// completionHandler: A Block object called when animations for this transaction group are
// completed.
// 
// The Block object takes no parameters and returns no value.
//
// # Discussion
// 
// Using this method allows you to more naturally group animations and an
// completion Block.
// 
// An example use is as follows. Using this method you would write the
// following code fragment:
// 
// The above code is semantically equivalent to the following:
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimationContext/runAnimationGroup(_:completionHandler:)
func (_NSAnimationContextClass NSAnimationContextClass) RunAnimationGroupCompletionHandler(changes AnimationContextHandler, completionHandler VoidHandler) {
_block0, _cleanup0 := NewAnimationContextBlock(changes)
	defer _cleanup0()
	_block1, _cleanup1 := NewVoidBlock(completionHandler)
	defer _cleanup1()
	objc.Send[objc.ID](objc.ID(_NSAnimationContextClass.class), objc.Sel("runAnimationGroup:completionHandler:"), _block0, _block1)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimationContext/runAnimationGroup(_:)
func (_NSAnimationContextClass NSAnimationContextClass) RunAnimationGroup(changes AnimationContextHandler) {
_block0, _cleanup0 := NewAnimationContextBlock(changes)
	defer _cleanup0()
	objc.Send[objc.ID](objc.ID(_NSAnimationContextClass.class), objc.Sel("runAnimationGroup:"), _block0)
}

// A completion Block that is called when the animations in the grouping are
// completed.
//
// # Discussion
// 
// If set to a non-`nil` value, a context’s `completionHandler` is
// guaranteed to be called on the main thread as soon as all animations
// subsequently added to the current [NSAnimationContext] grouping have
// completed or been cancelled.
// 
// This method drives the underlying [CATransaction][completionBlock()]
// property, although the Application Kit may assign a different, intermediary
// `completionBlock` to the current [CATransaction].
// 
// The completion handler waits for all animations to which the handler
// applies, independent of whether they are evaluated by the Application Kit
// or delegated to Core Animation for evaluation in the render tree before
// firing.
// 
// If no animations are added before the current grouping is ended—or the
// completionHandler is set to a different value—the handler will be invoked
// immediately.
//
// [completionBlock()]: https://developer.apple.com/documentation/QuartzCore/CATransaction/completionBlock()
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimationContext/completionHandler
func (a NSAnimationContext) CompletionHandler() VoidHandler {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("completionHandler"))
	_ = rv
	return nil
}
func (a NSAnimationContext) SetCompletionHandler(value VoidHandler) {
	block, cleanup := NewVoidBlock(value)
	defer cleanup()
	objc.Send[struct{}](a.ID, objc.Sel("setCompletionHandler:"), block)
}
// The duration used by animations created as a result of setting new values
// for an animatable property.
//
// # Discussion
// 
// Any animations that occur as a result of setting the values of animatable
// properties in the current context will run for this duration.
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimationContext/duration
func (a NSAnimationContext) Duration() float64 {
	rv := objc.Send[float64](a.ID, objc.Sel("duration"))
	return rv
}
func (a NSAnimationContext) SetDuration(value float64) {
	objc.Send[struct{}](a.ID, objc.Sel("setDuration:"), value)
}
// The timing function used for all animations within this animation proxy
// group.
//
// # Discussion
// 
// The NSAnimationContext timing function is analogous to the CATransaction
// [setAnimationTimingFunction(_:)] method.
// 
// Animations initiated through the “animator” proxy syntax, that do not
// have an explicitly specified timing functions, will inherit the enclosing
// [NSAnimationContext] instance’s [TimingFunction] if it is not `nil`
// (which is the default).
// 
// As with the existing [Duration] property, changing a timing function causes
// the same change in the underlying CATransaction instance’s
// [animationTimingFunction()].
// 
// Also as with the [Duration] property, you may change the timingFunction any
// number of times within a given NSAnimationContext [BeginGrouping] and
// [EndGrouping] block. Changes to the `timingFunction` will apply to any
// animations that are subsequently initiated in that NSAnimationContext
// grouping, until the `timingFunction` is possibly changed again.
//
// [animationTimingFunction()]: https://developer.apple.com/documentation/QuartzCore/CATransaction/animationTimingFunction()
// [setAnimationTimingFunction(_:)]: https://developer.apple.com/documentation/QuartzCore/CATransaction/setAnimationTimingFunction(_:)
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimationContext/timingFunction
func (a NSAnimationContext) TimingFunction() quartzcore.CAMediaTimingFunction {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("timingFunction"))
	return quartzcore.CAMediaTimingFunctionFromID(objc.ID(rv))
}
func (a NSAnimationContext) SetTimingFunction(value quartzcore.CAMediaTimingFunction) {
	objc.Send[struct{}](a.ID, objc.Sel("setTimingFunction:"), value)
}
// Determine if animations are enabled or not for animations that occur as a
// result of another property change.
//
// # Discussion
// 
// Using the [Animator] proxy will automatically set `allowsImplicitAnimation`
// to [true]. When [true], other properties can implicitly animate along with
// the initially changed property.
// 
// For instance, calling `[[view animator] frame]` will allow subviews to also
// animate their frame positions. When the value is [false] the behavior is
// diabled.
// 
// The default value is [false].
// 
// This is only applicable when layer backed on OS v10.8 and later.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimationContext/allowsImplicitAnimation
func (a NSAnimationContext) AllowsImplicitAnimation() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("allowsImplicitAnimation"))
	return rv
}
func (a NSAnimationContext) SetAllowsImplicitAnimation(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setAllowsImplicitAnimation:"), value)
}

// Returns the current animation context.
//
// # Return Value
// 
// The current animation context.
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimationContext/current
func (_NSAnimationContextClass NSAnimationContextClass) CurrentContext() NSAnimationContext {
	rv := objc.Send[objc.ID](objc.ID(_NSAnimationContextClass.class), objc.Sel("currentContext"))
	return NSAnimationContextFromID(objc.ID(rv))
}

// RunAnimationGroupSync is a synchronous wrapper around [NSAnimationContext.RunAnimationGroupCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (ac NSAnimationContextClass) RunAnimationGroupSync(ctx context.Context, changes AnimationContextHandler) error {
	done := make(chan struct{}, 1)
	ac.RunAnimationGroupCompletionHandler(changes, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// RunAnimationGroupSyncSync is a synchronous wrapper around [NSAnimationContext.RunAnimationGroup].
// It blocks until the completion handler fires or the context is cancelled.
func (ac NSAnimationContextClass) RunAnimationGroupSyncSync(ctx context.Context) (*NSAnimationContext, error) {
	done := make(chan *NSAnimationContext, 1)
	ac.RunAnimationGroup(func(val *NSAnimationContext) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

