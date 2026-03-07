// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSAnimation] class.
var (
	_NSAnimationClass     NSAnimationClass
	_NSAnimationClassOnce sync.Once
)

func getNSAnimationClass() NSAnimationClass {
	_NSAnimationClassOnce.Do(func() {
		_NSAnimationClass = NSAnimationClass{class: objc.GetClass("NSAnimation")}
	})
	return _NSAnimationClass
}

// GetNSAnimationClass returns the class object for NSAnimation.
func GetNSAnimationClass() NSAnimationClass {
	return getNSAnimationClass()
}

type NSAnimationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSAnimationClass) Alloc() NSAnimation {
	rv := objc.Send[NSAnimation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// An object that manages the timing and progress of animations in the user
// interface.
//
// # Overview
// 
// [NSAnimation] also lets you link together multiple animations so that when
// one animation ends another one starts. It does not provide any drawing
// support for animation and does not directly deal with views, targets, or
// actions.
// 
// [NSAnimation] objects have several characteristics, including duration,
// frame rate, and animation curve, which describes the relative speed of the
// animation over its course. You can set progress marks in an animation, each
// of which specifies a percentage of the animation completed; when an
// animation reaches a progress mark, it notifies its delegate and posts a
// notification to any observers. Animations execute in one of three blocking
// modes: blocking, non-blocking on the main thread, and non-blocking on a
// separate thread. The non-blocking modes permit the handling of user events
// while the animation is running.
// 
// # Subclassing Notes
// 
// The usual usage pattern for [NSAnimation] is to make a subclass that
// overrides (at least) the [NSAnimation.CurrentProgress] property to invoke the
// superclass implementation and then perform whatever animation action is
// needed. The method implementation might use the [NSAnimation.CurrentValue] property and
// then use that value to update some drawing; as a consequence of getting the
// current value, the method [AnimationValueForProgress] is sent to the
// delegate (if there is a delegate that implements the method). For more
// information on subclassing [NSAnimation], see [Animation Programming Guide
// for Cocoa].
//
// [Animation Programming Guide for Cocoa]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/AnimationGuide/Introduction/Introduction.html#//apple_ref/doc/uid/TP40003592
//
// # Initializing an NSAnimation Object
//
//   - [NSAnimation.InitWithDurationAnimationCurve]: Returns an [NSAnimation] object initialized with the specified duration and animation-curve values.
//
// # Configuring an Animation
//
//   - [NSAnimation.AnimationBlockingMode]: The blocking mode of the animation.
//   - [NSAnimation.SetAnimationBlockingMode]
//   - [NSAnimation.RunLoopModesForAnimating]: An array of strings representing the run loop modes in which the animation can run.
//   - [NSAnimation.AnimationCurve]: The timing curve for the animation.
//   - [NSAnimation.SetAnimationCurve]
//   - [NSAnimation.Duration]: The duration of the animation, in seconds.
//   - [NSAnimation.SetDuration]
//   - [NSAnimation.FrameRate]: The number of frame updates per second to generate for the animation.
//   - [NSAnimation.SetFrameRate]
//
// # Managing the Delegate
//
//   - [NSAnimation.Delegate]: The animation delegate.
//   - [NSAnimation.SetDelegate]
//
// # Controlling and Monitoring an Animation
//
//   - [NSAnimation.StartAnimation]: Starts the animation represented by the receiver.
//   - [NSAnimation.StopAnimation]: Stops the animation represented by the receiver.
//   - [NSAnimation.Animating]: A Boolean value indicating whether the animation is in progress.
//   - [NSAnimation.CurrentProgress]: The current progress of the animation.
//   - [NSAnimation.SetCurrentProgress]
//   - [NSAnimation.CurrentValue]: The current value of the animation effect, based on the current progress
//
// # Managing Progress Marks
//
//   - [NSAnimation.AddProgressMark]: Adds the progress mark to the receiver.
//   - [NSAnimation.RemoveProgressMark]: Removes progress mark from the receiver.
//   - [NSAnimation.ProgressMarks]: An array of floating-point numbers representing current progress marks.
//   - [NSAnimation.SetProgressMarks]
//
// # Linking Animations Together
//
//   - [NSAnimation.StartWhenAnimationReachesProgress]: Starts running the animation represented by the receiver when another animation reaches a specific progress mark.
//   - [NSAnimation.StopWhenAnimationReachesProgress]: Stops running the animation represented by the receiver when another animation reaches a specific progress mark.
//   - [NSAnimation.ClearStartAnimation]: Clears linkage to another animation that causes the receiver to start.
//   - [NSAnimation.ClearStopAnimation]: Clears linkage to another animation that causes the receiver to stop.
//
// # Initializers
//
//   - [NSAnimation.InitWithCoder]
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimation
type NSAnimation struct {
	objectivec.Object
}

// NSAnimationFromID constructs a [NSAnimation] from an objc.ID.
//
// An object that manages the timing and progress of animations in the user
// interface.
func NSAnimationFromID(id objc.ID) NSAnimation {
	return NSAnimation{objectivec.Object{id}}
}
// NOTE: NSAnimation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSAnimation] class.
//
// # Initializing an NSAnimation Object
//
//   - [INSAnimation.InitWithDurationAnimationCurve]: Returns an [NSAnimation] object initialized with the specified duration and animation-curve values.
//
// # Configuring an Animation
//
//   - [INSAnimation.AnimationBlockingMode]: The blocking mode of the animation.
//   - [INSAnimation.SetAnimationBlockingMode]
//   - [INSAnimation.RunLoopModesForAnimating]: An array of strings representing the run loop modes in which the animation can run.
//   - [INSAnimation.AnimationCurve]: The timing curve for the animation.
//   - [INSAnimation.SetAnimationCurve]
//   - [INSAnimation.Duration]: The duration of the animation, in seconds.
//   - [INSAnimation.SetDuration]
//   - [INSAnimation.FrameRate]: The number of frame updates per second to generate for the animation.
//   - [INSAnimation.SetFrameRate]
//
// # Managing the Delegate
//
//   - [INSAnimation.Delegate]: The animation delegate.
//   - [INSAnimation.SetDelegate]
//
// # Controlling and Monitoring an Animation
//
//   - [INSAnimation.StartAnimation]: Starts the animation represented by the receiver.
//   - [INSAnimation.StopAnimation]: Stops the animation represented by the receiver.
//   - [INSAnimation.Animating]: A Boolean value indicating whether the animation is in progress.
//   - [INSAnimation.CurrentProgress]: The current progress of the animation.
//   - [INSAnimation.SetCurrentProgress]
//   - [INSAnimation.CurrentValue]: The current value of the animation effect, based on the current progress
//
// # Managing Progress Marks
//
//   - [INSAnimation.AddProgressMark]: Adds the progress mark to the receiver.
//   - [INSAnimation.RemoveProgressMark]: Removes progress mark from the receiver.
//   - [INSAnimation.ProgressMarks]: An array of floating-point numbers representing current progress marks.
//   - [INSAnimation.SetProgressMarks]
//
// # Linking Animations Together
//
//   - [INSAnimation.StartWhenAnimationReachesProgress]: Starts running the animation represented by the receiver when another animation reaches a specific progress mark.
//   - [INSAnimation.StopWhenAnimationReachesProgress]: Stops running the animation represented by the receiver when another animation reaches a specific progress mark.
//   - [INSAnimation.ClearStartAnimation]: Clears linkage to another animation that causes the receiver to start.
//   - [INSAnimation.ClearStopAnimation]: Clears linkage to another animation that causes the receiver to stop.
//
// # Initializers
//
//   - [INSAnimation.InitWithCoder]
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimation
type INSAnimation interface {
	objectivec.IObject

	// Topic: Initializing an NSAnimation Object

	// Returns an [NSAnimation] object initialized with the specified duration and animation-curve values.
	InitWithDurationAnimationCurve(duration float64, animationCurve NSAnimationCurve) NSAnimation

	// Topic: Configuring an Animation

	// The blocking mode of the animation.
	AnimationBlockingMode() NSAnimationBlockingMode
	SetAnimationBlockingMode(value NSAnimationBlockingMode)
	// An array of strings representing the run loop modes in which the animation can run.
	RunLoopModesForAnimating() []string
	// The timing curve for the animation.
	AnimationCurve() NSAnimationCurve
	SetAnimationCurve(value NSAnimationCurve)
	// The duration of the animation, in seconds.
	Duration() float64
	SetDuration(value float64)
	// The number of frame updates per second to generate for the animation.
	FrameRate() float32
	SetFrameRate(value float32)

	// Topic: Managing the Delegate

	// The animation delegate.
	Delegate() NSAnimationDelegate
	SetDelegate(value NSAnimationDelegate)

	// Topic: Controlling and Monitoring an Animation

	// Starts the animation represented by the receiver.
	StartAnimation()
	// Stops the animation represented by the receiver.
	StopAnimation()
	// A Boolean value indicating whether the animation is in progress.
	Animating() bool
	// The current progress of the animation.
	CurrentProgress() NSAnimationProgress
	SetCurrentProgress(value NSAnimationProgress)
	// The current value of the animation effect, based on the current progress
	CurrentValue() float32

	// Topic: Managing Progress Marks

	// Adds the progress mark to the receiver.
	AddProgressMark(progressMark NSAnimationProgress)
	// Removes progress mark from the receiver.
	RemoveProgressMark(progressMark NSAnimationProgress)
	// An array of floating-point numbers representing current progress marks.
	ProgressMarks() []foundation.NSNumber
	SetProgressMarks(value []foundation.NSNumber)

	// Topic: Linking Animations Together

	// Starts running the animation represented by the receiver when another animation reaches a specific progress mark.
	StartWhenAnimationReachesProgress(animation INSAnimation, startProgress NSAnimationProgress)
	// Stops running the animation represented by the receiver when another animation reaches a specific progress mark.
	StopWhenAnimationReachesProgress(animation INSAnimation, stopProgress NSAnimationProgress)
	// Clears linkage to another animation that causes the receiver to start.
	ClearStartAnimation()
	// Clears linkage to another animation that causes the receiver to stop.
	ClearStopAnimation()

	// Topic: Initializers

	InitWithCoder(coder foundation.INSCoder) NSAnimation

	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (a NSAnimation) Init() NSAnimation {
	rv := objc.Send[NSAnimation](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a NSAnimation) Autorelease() NSAnimation {
	rv := objc.Send[NSAnimation](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSAnimation creates a new NSAnimation instance.
func NewNSAnimation() NSAnimation {
	class := getNSAnimationClass()
	rv := objc.Send[NSAnimation](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/AppKit/NSAnimation/init(coder:)
func NewAnimationWithCoder(coder foundation.INSCoder) NSAnimation {
	instance := getNSAnimationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSAnimationFromID(rv)
}


// Returns an [NSAnimation] object initialized with the specified duration and
// animation-curve values.
//
// duration: The number of seconds over which the animation occurs. Specifying a
// negative number raises an exception.
//
// animationCurve: An [NSAnimationCurve] constant that describes the relative speed of the
// animation over its course; if it is zero, the default curve
// ([NSAnimationEaseInOut]) is used.
//
// # Return Value
// 
// An initialized [NSAnimation] instance. Returns `nil` if the object could
// not be initialized.
//
// # Discussion
// 
// You can always later change the duration of an [NSAnimation] object by
// changing the [Duration] property, even while the animation is running. See
// “Constants” for descriptions of the NSAnimationCurve constants.
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimation/init(duration:animationCurve:)
func NewAnimationWithDurationAnimationCurve(duration float64, animationCurve NSAnimationCurve) NSAnimation {
	instance := getNSAnimationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDuration:animationCurve:"), duration, animationCurve)
	return NSAnimationFromID(rv)
}







// Returns an [NSAnimation] object initialized with the specified duration and
// animation-curve values.
//
// duration: The number of seconds over which the animation occurs. Specifying a
// negative number raises an exception.
//
// animationCurve: An [NSAnimationCurve] constant that describes the relative speed of the
// animation over its course; if it is zero, the default curve
// ([NSAnimationEaseInOut]) is used.
//
// # Return Value
// 
// An initialized [NSAnimation] instance. Returns `nil` if the object could
// not be initialized.
//
// # Discussion
// 
// You can always later change the duration of an [NSAnimation] object by
// changing the [Duration] property, even while the animation is running. See
// “Constants” for descriptions of the NSAnimationCurve constants.
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimation/init(duration:animationCurve:)
func (a NSAnimation) InitWithDurationAnimationCurve(duration float64, animationCurve NSAnimationCurve) NSAnimation {
	rv := objc.Send[NSAnimation](a.ID, objc.Sel("initWithDuration:animationCurve:"), duration, animationCurve)
	return rv
}

// Starts the animation represented by the receiver.
//
// # Discussion
// 
// A strong reference to the animation is maintained until the end of the
// animation or until its [StopAnimation] method is called. If the blocking
// mode is [NSAnimation.BlockingMode.blocking], this method returns after the
// animation has completed or the delegate sends it [StopAnimation]. If the
// receiver has a progress of `1.0`, it starts again at `0.0`.
//
// [NSAnimation.BlockingMode.blocking]: https://developer.apple.com/documentation/AppKit/NSAnimation/BlockingMode/blocking
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimation/start()
func (a NSAnimation) StartAnimation() {
	objc.Send[objc.ID](a.ID, objc.Sel("startAnimation"))
}

// Stops the animation represented by the receiver.
//
// # Discussion
// 
// The current progress of the receiver is not reset. When this method is sent
// to instances of [NSViewAnimation] (a subclass of [NSAnimation]) the
// receiver moves to the end frame location.
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimation/stop()
func (a NSAnimation) StopAnimation() {
	objc.Send[objc.ID](a.ID, objc.Sel("stopAnimation"))
}

// Adds the progress mark to the receiver.
//
// progressMark: A `float` value (typed as NSAnimationProgress) between 0.0 and 1.0. Values
// outside that range are pinned to 0.0 or 1.0, whichever is nearest.
//
// # Discussion
// 
// A progress mark represents a percentage of the animation completed. When
// the animation reaches a progress mark, an [AnimationDidReachProgressMark]
// message is sent to the delegate and an [progressMarkNotification] is
// broadcast to all observers. You might receive multiple notifications of
// progress advances over multiple marks.
//
// [progressMarkNotification]: https://developer.apple.com/documentation/AppKit/NSAnimation/progressMarkNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimation/addProgressMark(_:)
func (a NSAnimation) AddProgressMark(progressMark NSAnimationProgress) {
	objc.Send[objc.ID](a.ID, objc.Sel("addProgressMark:"), progressMark)
}

// Removes progress mark from the receiver.
//
// progressMark: A `float` value (typed as NSAnimationProgress) that indicates the portion
// of the animation completed. The value should correspond to a progress mark
// set with [AddProgressMark] or [NSAnimation].
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimation/removeProgressMark(_:)
func (a NSAnimation) RemoveProgressMark(progressMark NSAnimationProgress) {
	objc.Send[objc.ID](a.ID, objc.Sel("removeProgressMark:"), progressMark)
}

// Starts running the animation represented by the receiver when another
// animation reaches a specific progress mark.
//
// animation: The other [NSAnimation] object with which the receiver is linked.
//
// startProgress: A `float` value (typed as NSAnimationProgress) that specifies a progress
// mark of the other animation.
//
// # Discussion
// 
// This method links the running of two animations together. You can set only
// one [NSAnimation] object as a start animation and one as a stop animation
// at any one time. Setting a new start animation removes any animation
// previously set.
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimation/start(when:reachesProgress:)
func (a NSAnimation) StartWhenAnimationReachesProgress(animation INSAnimation, startProgress NSAnimationProgress) {
	objc.Send[objc.ID](a.ID, objc.Sel("startWhenAnimation:reachesProgress:"), animation, startProgress)
}

// Stops running the animation represented by the receiver when another
// animation reaches a specific progress mark.
//
// animation: The other [NSAnimation] object with which the receiver is linked.
//
// stopProgress: A `float` value (typed as NSAnimationProgress) that specifies a progress
// mark of the other animation.
//
// # Discussion
// 
// This method links the running of two animations together. You can set only
// one [NSAnimation] object as a start animation and one as a stop animation
// at any one time. Setting a new stop animation removes any animation
// previously set.
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimation/stop(when:reachesProgress:)
func (a NSAnimation) StopWhenAnimationReachesProgress(animation INSAnimation, stopProgress NSAnimationProgress) {
	objc.Send[objc.ID](a.ID, objc.Sel("stopWhenAnimation:reachesProgress:"), animation, stopProgress)
}

// Clears linkage to another animation that causes the receiver to start.
//
// # Discussion
// 
// The linkage to the other animation is made with
// [StartWhenAnimationReachesProgress].
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimation/clearStart()
func (a NSAnimation) ClearStartAnimation() {
	objc.Send[objc.ID](a.ID, objc.Sel("clearStartAnimation"))
}

// Clears linkage to another animation that causes the receiver to stop.
//
// # Discussion
// 
// The linkage to the other animation is made with
// [StopWhenAnimationReachesProgress].
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimation/clearStop()
func (a NSAnimation) ClearStopAnimation() {
	objc.Send[objc.ID](a.ID, objc.Sel("clearStopAnimation"))
}

//
// See: https://developer.apple.com/documentation/AppKit/NSAnimation/init(coder:)
func (a NSAnimation) InitWithCoder(coder foundation.INSCoder) NSAnimation {
	rv := objc.Send[NSAnimation](a.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (a NSAnimation) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](a.ID, objc.Sel("encodeWithCoder:"), coder)
}












// The blocking mode of the animation.
//
// # Discussion
// 
// The value in this property determines whether the animation blocks a given
// thread. The default value of this property is
// [NSAnimation.BlockingMode.blocking], which means that the animation runs on
// the main thread in a custom run-loop mode that blocks user events. When
// changing the value of this property, the new blocking mode takes effect the
// next time the animation is started and has no effect on an in-progress
// animation.
// 
// If you set the block mode to [NSAnimation.BlockingMode.nonblocking], the
// animation runs in the main thread in one of the standard run-loop modes or
// in a mode returned from [NSAnimation]. If you set the mode to
// [NSAnimation.BlockingMode.nonblockingThreaded], a new thread is spawned to
// run the animation.
//
// [NSAnimation.BlockingMode.blocking]: https://developer.apple.com/documentation/AppKit/NSAnimation/BlockingMode/blocking
// [NSAnimation.BlockingMode.nonblockingThreaded]: https://developer.apple.com/documentation/AppKit/NSAnimation/BlockingMode/nonblockingThreaded
// [NSAnimation.BlockingMode.nonblocking]: https://developer.apple.com/documentation/AppKit/NSAnimation/BlockingMode/nonblocking
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimation/animationBlockingMode
func (a NSAnimation) AnimationBlockingMode() NSAnimationBlockingMode {
	rv := objc.Send[NSAnimationBlockingMode](a.ID, objc.Sel("animationBlockingMode"))
	return NSAnimationBlockingMode(rv)
}
func (a NSAnimation) SetAnimationBlockingMode(value NSAnimationBlockingMode) {
	objc.Send[struct{}](a.ID, objc.Sel("setAnimationBlockingMode:"), value)
}



// An array of strings representing the run loop modes in which the animation
// can run.
//
// # Discussion
// 
// The default value of this property is `nil`, which indicates that the
// animation can be run in the default, modal, or event-tracking modes. The
// value of this property is ignored if the animation blocking mode is
// something other than [NSAnimation.BlockingMode.nonblocking].
// 
// For information about run loop modes and for constants, see [RunLoop].
//
// [NSAnimation.BlockingMode.nonblocking]: https://developer.apple.com/documentation/AppKit/NSAnimation/BlockingMode/nonblocking
// [RunLoop]: https://developer.apple.com/documentation/Foundation/RunLoop
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimation/runLoopModesForAnimating
func (a NSAnimation) RunLoopModesForAnimating() []string {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("runLoopModesForAnimating"))
	return objc.ConvertSliceToStrings(rv)
}



// The timing curve for the animation.
//
// # Discussion
// 
// The animation curve describes the relative frame rate over the course of
// the animation; predefined curves are linear, ease in (slow down near end),
// ease out (slowly speed up at start), and ease in-ease out (S-curve).
// Changing the value of this property changes the timing of an in-progress
// animation. The value of this property is ignored if the delegate implements
// the [AnimationValueForProgress] method.
// 
// Setting this property to an invalid value raises an exception. For a list
// of valid animation values, see [NSAnimation.Curve].
//
// [NSAnimation.Curve]: https://developer.apple.com/documentation/AppKit/NSAnimation/Curve
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimation/animationCurve
func (a NSAnimation) AnimationCurve() NSAnimationCurve {
	rv := objc.Send[NSAnimationCurve](a.ID, objc.Sel("animationCurve"))
	return NSAnimationCurve(rv)
}
func (a NSAnimation) SetAnimationCurve(value NSAnimationCurve) {
	objc.Send[struct{}](a.ID, objc.Sel("setAnimationCurve:"), value)
}



// The duration of the animation, in seconds.
//
// # Discussion
// 
// The value of this property must be greater than or equal to `0`. Setting
// the duration to a negative value raises an exception.
// 
// You can change the duration of an animation while it is running. Setting
// the duration to a value that is less than the current progress value ends
// an in-progress animation.
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimation/duration
func (a NSAnimation) Duration() float64 {
	rv := objc.Send[float64](a.ID, objc.Sel("duration"))
	return rv
}
func (a NSAnimation) SetDuration(value float64) {
	objc.Send[struct{}](a.ID, objc.Sel("setDuration:"), value)
}



// The number of frame updates per second to generate for the animation.
//
// # Discussion
// 
// The value of this property must be greater than or equal to `0`. Specifying
// a value of `0.0` causes the animation to run as fast as possible. Setting
// the property to a negative value raises an exception.
// 
// The frame rate is not guaranteed due to differences among systems for the
// time needed to process a frame. You can change the frame rate while an
// animation is running and the new value is used at the next frame. The
// default frame rate is set to a reasonable value (which is subject to future
// change).
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimation/frameRate
func (a NSAnimation) FrameRate() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("frameRate"))
	return rv
}
func (a NSAnimation) SetFrameRate(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setFrameRate:"), value)
}



// The animation delegate.
//
// # Discussion
// 
// The delegate must conform to the [NSAnimationDelegate].
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimation/delegate
func (a NSAnimation) Delegate() NSAnimationDelegate {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("delegate"))
	return NSAnimationDelegateObjectFromID(rv)
}
func (a NSAnimation) SetDelegate(value NSAnimationDelegate) {
	objc.Send[struct{}](a.ID, objc.Sel("setDelegate:"), value)
}



// A Boolean value indicating whether the animation is in progress.
//
// # Discussion
// 
// The value of this property is [true] when the animation is in progress or
// [false] when it is stopped.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimation/isAnimating
func (a NSAnimation) Animating() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isAnimating"))
	return rv
}



// The current progress of the animation.
//
// # Discussion
// 
// This property contains the completion percentage of the animation. Valid
// values are in the range `0.0` to `1.0`, where `0.0` represents the
// beginning of the animation and `1.0` represents the end of the animation.
// 
// Changing the value of this property adjusts the progress of a running
// animation. Setting this property to a value less than `0.0` sets the value
// of the property to `0.0`. Similarly, specifying a value greater than `1.0`
// changes the value of the property to `1.0`. The [NSAnimation] class updates
// the value of this property during the animation. To perform additional
// tasks at specific progress points, use the delegate’s
// [AnimationValueForProgress] method.
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimation/currentProgress
func (a NSAnimation) CurrentProgress() NSAnimationProgress {
	rv := objc.Send[NSAnimationProgress](a.ID, objc.Sel("currentProgress"))
	return NSAnimationProgress(rv)
}
func (a NSAnimation) SetCurrentProgress(value NSAnimationProgress) {
	objc.Send[struct{}](a.ID, objc.Sel("setCurrentProgress:"), value)
}



// The current value of the animation effect, based on the current progress
//
// # Discussion
// 
// An [NSAnimation] object gets the current value from delegate’s
// [AnimationValueForProgress] method. If that method is not implemented, the
// animation computes the current value from the current progress by factoring
// in the animation curve. An animation object does not access this property
// directly. Instances of [NSAnimation] subclasses or other objects can invoke
// this method on a periodic basis to get the current value.
// 
// Subclasses may override this property and return a custom curve value
// instead of implementing [AnimationValueForProgress], thereby saving on the
// overhead of using a delegate. The current value can be less than `0.0` or
// greater than `1.0`. For example, if you make the value greater than `1.0`
// you can achieve a “rubber effect” where the size of a view is
// temporarily larger before its final size.
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimation/currentValue
func (a NSAnimation) CurrentValue() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("currentValue"))
	return rv
}



// An array of floating-point numbers representing current progress marks.
//
// # Discussion
// 
// The value of this property is an array of [NSNumber] objects, each of which
// contains a float value, which are typed to the [NSAnimationProgress] type.
// If there are no progress marks, the array is empty. Setting the value of
// this property is `nil` clears all progress marks.
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimation/progressMarks
func (a NSAnimation) ProgressMarks() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("progressMarks"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}
func (a NSAnimation) SetProgressMarks(value []foundation.NSNumber) {
	objc.Send[struct{}](a.ID, objc.Sel("setProgressMarks:"), objectivec.IObjectSliceToNSArray(value))
}

























