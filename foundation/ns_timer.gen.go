// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [Timer] class.
var (
	_TimerClass     TimerClass
	_TimerClassOnce sync.Once
)

func getTimerClass() TimerClass {
	_TimerClassOnce.Do(func() {
		_TimerClass = TimerClass{class: objc.GetClass("NSTimer")}
	})
	return _TimerClass
}

// GetTimerClass returns the class object for NSTimer.
func GetTimerClass() TimerClass {
	return getTimerClass()
}

type TimerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (tc TimerClass) Alloc() Timer {
	rv := objc.Send[Timer](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}







// A timer that fires after a certain time interval has elapsed, sending a
// specified message to a target object.
//
// # Overview
// 
// Timers work in conjunction with run loops. Run loops maintain strong
// references to their timers, so you don’t have to maintain your own strong
// reference to a timer after you have added it to a run loop.
// 
// To use a timer effectively, you should be aware of how run loops operate.
// See [Threading Programming Guide] for more information.
// 
// A timer is not a real-time mechanism. If a timer’s firing time occurs
// during a long run loop callout or while the run loop is in a mode that
// isn’t monitoring the timer, the timer doesn’t fire until the next time
// the run loop checks the timer. Therefore, the actual time at which a timer
// fires can be significantly later. See also [NSTimer].
// 
// [NSTimer] is toll-free bridged with its Core Foundation counterpart,
// [CFRunLoopTimer]. See [Toll-Free Bridging] for more information.
// 
// # Comparing Repeating and Nonrepeating Timers
// 
// You specify whether a timer is repeating or nonrepeating at creation time.
// A nonrepeating timer fires once and then invalidates itself automatically,
// thereby preventing the timer from firing again. By contrast, a repeating
// timer fires and then reschedules itself on the same run loop. A repeating
// timer always schedules itself based on the scheduled firing time, as
// opposed to the actual firing time. For example, if a timer is scheduled to
// fire at a particular time and every 5 seconds after that, the scheduled
// firing time will always fall on the original 5-second time intervals, even
// if the actual firing time gets delayed. If the firing time is delayed so
// far that it passes one or more of the scheduled firing times, the timer is
// fired only once for that time period; the timer is then rescheduled, after
// firing, for the next scheduled firing time in the future.
// 
// # Timer Tolerance
// 
// In iOS 7 and later and macOS 10.9 and later, you can specify a tolerance
// for a timer ([Tolerance]). This flexibility in when a timer fires improves
// the system’s ability to optimize for increased power savings and
// responsiveness. The timer may fire at any time between its scheduled fire
// date and the scheduled fire date plus the tolerance. The timer doesn’t
// fire before the scheduled fire date. For repeating timers, the next fire
// date is calculated from the original fire date regardless of tolerance
// applied at individual fire times, to avoid drift. The default value is
// zero, which means no additional tolerance is applied. The system reserves
// the right to apply a small amount of tolerance to certain timers regardless
// of the value of the [Tolerance] property.
// 
// As the user of the timer, you can determine the appropriate tolerance for a
// timer. A general rule, set the tolerance to at least 10% of the interval,
// for a repeating timer. Even a small amount of tolerance has significant
// positive impact on the power usage of your application. The system may
// enforce a maximum value for the tolerance.
// 
// # Scheduling Timers in Run Loops
// 
// You can register a timer in only one run loop at a time, although it can be
// added to multiple run loop modes within that run loop. There are three ways
// to create a timer:
// 
// - Use the [ScheduledTimerWithTimeIntervalInvocationRepeats] or
// [ScheduledTimerWithTimeIntervalTargetSelectorUserInfoRepeats] class method
// to create the timer and schedule it on the current run loop in the default
// mode. - Use the [TimerWithTimeIntervalInvocationRepeats] or
// [TimerWithTimeIntervalTargetSelectorUserInfoRepeats] class method to create
// the timer object without scheduling it on a run loop. (After creating it,
// you must add the timer to a run loop manually by calling the
// [AddTimerForMode] method of the corresponding [NSRunLoop] object.) -
// Allocate the timer and initialize it using the
// [InitWithFireDateIntervalTargetSelectorUserInfoRepeats] method. (After
// creating it, you must add the timer to a run loop manually by calling the
// [AddTimerForMode] method of the corresponding [NSRunLoop] object.)
// 
// Once scheduled on a run loop, the timer fires at the specified interval
// until it is invalidated. A nonrepeating timer invalidates itself
// immediately after it fires. However, for a repeating timer, you must
// invalidate the timer object yourself by calling its [Invalidate] method.
// Calling this method requests the removal of the timer from the current run
// loop; as a result, you should always call the [Invalidate] method from the
// same thread on which the timer was installed. Invalidating the timer
// immediately disables it so that it no longer affects the run loop. The run
// loop then removes the timer (and the strong reference it had to the timer),
// either just before the [Invalidate] method returns or at some later point.
// Once invalidated, timer objects cannot be reused.
// 
// After a repeating timer fires, it schedules the next firing for the nearest
// future date that is an integer multiple of the timer interval after the
// last scheduled fire date, within the specified [Tolerance]. If the time
// taken to call out to perform a selector or invocation is longer than the
// specified interval, the timer schedules only the next firing; that is, the
// timer doesn’t attempt to compensate for any missed firings that would
// have occurred while calling the specified selector or invocation.
// 
// # Subclassing Notes
// 
// Do not subclass [NSTimer].
//
// [CFRunLoopTimer]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoopTimer
// [Threading Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Multithreading/Introduction/Introduction.html#//apple_ref/doc/uid/10000057i
// [Toll-Free Bridging]: https://developer.apple.com/library/archive/documentation/General/Conceptual/CocoaEncyclopedia/Toll-FreeBridgin/Toll-FreeBridgin.html#//apple_ref/doc/uid/TP40010810-CH2
//
// # Creating a Timer
//
//   - [Timer.InitWithFireDateIntervalRepeatsBlock]: Initializes a timer for the specified date and time interval with the specified block.
//   - [Timer.InitWithFireDateIntervalTargetSelectorUserInfoRepeats]: Initializes a timer using the specified object and selector.
//
// # Firing a Timer
//
//   - [Timer.Fire]: Causes the timer’s message to be sent to its target.
//
// # Stopping a Timer
//
//   - [Timer.Invalidate]: Stops the timer from ever firing again and requests its removal from its run loop.
//
// # Retrieving Timer Information
//
//   - [Timer.Valid]: A Boolean value that indicates whether the timer is currently valid.
//   - [Timer.FireDate]: The date at which the timer will fire.
//   - [Timer.SetFireDate]
//   - [Timer.TimeInterval]: The timer’s time interval, in seconds.
//   - [Timer.UserInfo]: The receiver’s `userInfo` object.
//
// # Configuring Firing Tolerance
//
//   - [Timer.Tolerance]: The amount of time after the scheduled fire date that the timer may fire.
//   - [Timer.SetTolerance]
//
// See: https://developer.apple.com/documentation/Foundation/Timer
type Timer struct {
	objectivec.Object
}

// TimerFromID constructs a [Timer] from an objc.ID.
//
// A timer that fires after a certain time interval has elapsed, sending a
// specified message to a target object.
func TimerFromID(id objc.ID) Timer {
	return NSTimer{objectivec.Object{id}}
}

// NSTimerFromID is an alias for [TimerFromID] for cross-framework compatibility.
func NSTimerFromID(id objc.ID) Timer { return TimerFromID(id) }
// NOTE: Timer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [Timer] class.
//
// # Creating a Timer
//
//   - [ITimer.InitWithFireDateIntervalRepeatsBlock]: Initializes a timer for the specified date and time interval with the specified block.
//   - [ITimer.InitWithFireDateIntervalTargetSelectorUserInfoRepeats]: Initializes a timer using the specified object and selector.
//
// # Firing a Timer
//
//   - [ITimer.Fire]: Causes the timer’s message to be sent to its target.
//
// # Stopping a Timer
//
//   - [ITimer.Invalidate]: Stops the timer from ever firing again and requests its removal from its run loop.
//
// # Retrieving Timer Information
//
//   - [ITimer.Valid]: A Boolean value that indicates whether the timer is currently valid.
//   - [ITimer.FireDate]: The date at which the timer will fire.
//   - [ITimer.SetFireDate]
//   - [ITimer.TimeInterval]: The timer’s time interval, in seconds.
//   - [ITimer.UserInfo]: The receiver’s `userInfo` object.
//
// # Configuring Firing Tolerance
//
//   - [ITimer.Tolerance]: The amount of time after the scheduled fire date that the timer may fire.
//   - [ITimer.SetTolerance]
//
// See: https://developer.apple.com/documentation/Foundation/Timer
type ITimer interface {
	objectivec.IObject

	// Topic: Creating a Timer

	// Initializes a timer for the specified date and time interval with the specified block.
	InitWithFireDateIntervalRepeatsBlock(date INSDate, interval float64, repeats bool, block TimerHandler) Timer
	// Initializes a timer using the specified object and selector.
	InitWithFireDateIntervalTargetSelectorUserInfoRepeats(date INSDate, ti float64, t objectivec.IObject, s objc.SEL, ui objectivec.IObject, rep bool) Timer

	// Topic: Firing a Timer

	// Causes the timer’s message to be sent to its target.
	Fire()

	// Topic: Stopping a Timer

	// Stops the timer from ever firing again and requests its removal from its run loop.
	Invalidate()

	// Topic: Retrieving Timer Information

	// A Boolean value that indicates whether the timer is currently valid.
	Valid() bool
	// The date at which the timer will fire.
	FireDate() INSDate
	SetFireDate(value INSDate)
	// The timer’s time interval, in seconds.
	TimeInterval() float64
	// The receiver’s `userInfo` object.
	UserInfo() objectivec.IObject

	// Topic: Configuring Firing Tolerance

	// The amount of time after the scheduled fire date that the timer may fire.
	Tolerance() float64
	SetTolerance(value float64)
}





// Init initializes the instance.
func (t Timer) Init() Timer {
	rv := objc.Send[Timer](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t Timer) Autorelease() Timer {
	rv := objc.Send[Timer](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTimer creates a new Timer instance.
func NewTimer() Timer {
	class := getTimerClass()
	rv := objc.Send[Timer](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Initializes a timer using the specified object and selector.
//
// date: The time at which the timer should first fire.
//
// ti: For a repeating timer, this parameter contains the number of seconds
// between firings of the timer. If `ti` is less than or equal to `0.0`, this
// method chooses the nonnegative value of `0.0001` seconds instead.
//
// t: The object to which to send the message specified by `aSelector` when the
// timer fires. The timer maintains a strong reference to this object until it
// (the timer) is invalidated.
//
// s: The message to send to `target` when the timer fires.
// 
// The selector should have the following signature: `` (including a colon to
// indicate that the method takes an argument). The timer passes itself as the
// argument, thus the method would adopt the following pattern:
//
// ui: Custom user info for the timer. The timer maintains a strong reference to
// this object until it (the timer) is invalidated. This parameter may be
// `nil`.
//
// rep: If [true], the timer will repeatedly reschedule itself until invalidated.
// If [false], the timer will be invalidated after it fires.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// The receiver, initialized such that, when added to a run loop, it will fire
// at `date` and then, if `repeats` is [true], every `ti` after that.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// You must add the new timer to a run loop, using [AddTimerForMode]. Upon
// firing, the timer sends the message `aSelector` to `target`. (If the timer
// is configured to repeat, there is no need to subsequently re-add the timer
// to the run loop.)
//
// See: https://developer.apple.com/documentation/Foundation/Timer/init(fireAt:interval:target:selector:userInfo:repeats:)
func NewTimerWithFireDateIntervalTargetSelectorUserInfoRepeats(date INSDate, ti float64, t objectivec.IObject, s objc.SEL, ui objectivec.IObject, rep bool) Timer {
	instance := getTimerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFireDate:interval:target:selector:userInfo:repeats:"), date, ti, t, s, ui, rep)
	return TimerFromID(rv)
}


// Initializes a timer object with the specified invocation object.
//
// ti: The number of seconds between firings of the timer. If `ti` is less than or
// equal to `0.0`, this method chooses the nonnegative value of `0.0001`
// seconds instead.
//
// invocation: The invocation to use when the timer fires. The timer instructs the
// invocation object to maintain a strong reference to its arguments.
//
// yesOrNo: If [true], the timer will repeatedly reschedule itself until invalidated.
// If [false], the timer will be invalidated after it fires.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// A new [NSTimer] object, configured according to the specified parameters.
//
// # Discussion
// 
// You must add the new timer to a run loop, using [AddTimerForMode]. Then,
// after `ti` have elapsed, the timer fires, invoking `invocation`. (If the
// timer is configured to repeat, there is no need to subsequently re-add the
// timer to the run loop.)
//
// See: https://developer.apple.com/documentation/Foundation/Timer/init(timeInterval:invocation:repeats:)
func NewTimerWithTimeIntervalInvocationRepeats(ti float64, invocation INSInvocation, yesOrNo bool) Timer {
	rv := objc.Send[objc.ID](objc.ID(getTimerClass().class), objc.Sel("timerWithTimeInterval:invocation:repeats:"), ti, invocation, yesOrNo)
	return TimerFromID(rv)
}


// Initializes a timer object with the specified object and selector.
//
// ti: The number of seconds between firings of the timer. If `ti` is less than or
// equal to `0.0`, this method chooses the nonnegative value of `0.0001`
// seconds instead.
//
// aTarget: The object to which to send the message specified by `aSelector` when the
// timer fires. The timer maintains a strong reference to this object until it
// (the timer) is invalidated.
//
// aSelector: The message to send to `target` when the timer fires.
// 
// The selector should have the following signature: `` (including a colon to
// indicate that the method takes an argument). The timer passes itself as the
// argument, thus the method would adopt the following pattern:
//
// userInfo: Custom user info for the timer.
// 
// The timer maintains a strong reference to this object until it (the timer)
// is invalidated. This parameter may be `nil`.
//
// yesOrNo: If [true], the timer will repeatedly reschedule itself until invalidated.
// If [false], the timer will be invalidated after it fires.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// A new [NSTimer] object, configured according to the specified parameters.
//
// # Discussion
// 
// You must add the new timer to a run loop, using [AddTimerForMode]. Then,
// after `ti` seconds have elapsed, the timer fires, sending the message
// `aSelector` to `target`. (If the timer is configured to repeat, there is no
// need to subsequently re-add the timer to the run loop.)
//
// See: https://developer.apple.com/documentation/Foundation/Timer/init(timeInterval:target:selector:userInfo:repeats:)
func NewTimerWithTimeIntervalTargetSelectorUserInfoRepeats(ti float64, aTarget objectivec.IObject, aSelector objc.SEL, userInfo objectivec.IObject, yesOrNo bool) Timer {
	rv := objc.Send[objc.ID](objc.ID(getTimerClass().class), objc.Sel("timerWithTimeInterval:target:selector:userInfo:repeats:"), ti, aTarget, aSelector, userInfo, yesOrNo)
	return TimerFromID(rv)
}







// Initializes a timer for the specified date and time interval with the
// specified block.
//
// date: The time at which the timer should first fire.
//
// interval: For a repeating timer, this parameter contains the number of seconds
// between firings of the timer. If `interval` is less than or equal to `0.0`,
// this method chooses the nonnegative value of `0.0001` seconds instead.
//
// repeats: If [true], the timer will repeatedly reschedule itself until invalidated.
// If [false], the timer will be invalidated after it fires.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// block: A block to be executed when the timer fires. The block takes a single
// [NSTimer] parameter and has no return value.
//
// # Return Value
// 
// A new [NSTimer] object, configured according to the specified parameters.
//
// # Discussion
// 
// You must add the new timer to a run loop, using [AddTimerForMode]. Upon
// firing, after `interval` seconds have elapsed, the timer fires, executing
// `block`. (If the timer is configured to repeat, you don’t need to add the
// timer to the run loop again.)
//
// See: https://developer.apple.com/documentation/Foundation/Timer/init(fire:interval:repeats:block:)
func (t Timer) InitWithFireDateIntervalRepeatsBlock(date INSDate, interval float64, repeats bool, block TimerHandler) Timer {
		_block3, _cleanup3 := NewTimerBlock(block)
	defer _cleanup3()
		rv := objc.Send[objc.ID](t.ID, objc.Sel("initWithFireDate:interval:repeats:block:"), date, interval, repeats, _block3)
	return NSTimerFromID(rv)
}

// Initializes a timer using the specified object and selector.
//
// date: The time at which the timer should first fire.
//
// ti: For a repeating timer, this parameter contains the number of seconds
// between firings of the timer. If `ti` is less than or equal to `0.0`, this
// method chooses the nonnegative value of `0.0001` seconds instead.
//
// t: The object to which to send the message specified by `aSelector` when the
// timer fires. The timer maintains a strong reference to this object until it
// (the timer) is invalidated.
//
// s: The message to send to `target` when the timer fires.
// 
// The selector should have the following signature: `` (including a colon to
// indicate that the method takes an argument). The timer passes itself as the
// argument, thus the method would adopt the following pattern:
//
// ui: Custom user info for the timer. The timer maintains a strong reference to
// this object until it (the timer) is invalidated. This parameter may be
// `nil`.
//
// rep: If [true], the timer will repeatedly reschedule itself until invalidated.
// If [false], the timer will be invalidated after it fires.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// The receiver, initialized such that, when added to a run loop, it will fire
// at `date` and then, if `repeats` is [true], every `ti` after that.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// You must add the new timer to a run loop, using [AddTimerForMode]. Upon
// firing, the timer sends the message `aSelector` to `target`. (If the timer
// is configured to repeat, there is no need to subsequently re-add the timer
// to the run loop.)
//
// See: https://developer.apple.com/documentation/Foundation/Timer/init(fireAt:interval:target:selector:userInfo:repeats:)
func (t_ Timer) InitWithFireDateIntervalTargetSelectorUserInfoRepeats(date INSDate, ti float64, t objectivec.IObject, s objc.SEL, ui objectivec.IObject, rep bool) Timer {
	rv := objc.Send[Timer](t_.ID, objc.Sel("initWithFireDate:interval:target:selector:userInfo:repeats:"), date, ti, t, s, ui, rep)
	return rv
}

// Causes the timer’s message to be sent to its target.
//
// # Discussion
// 
// You can use this method to fire a repeating timer without interrupting its
// regular firing schedule. If the timer is non-repeating, it is automatically
// invalidated after firing, even if its scheduled fire date has not arrived.
//
// See: https://developer.apple.com/documentation/Foundation/Timer/fire()
func (t Timer) Fire() {
	objc.Send[objc.ID](t.ID, objc.Sel("fire"))
}

// Stops the timer from ever firing again and requests its removal from its
// run loop.
//
// # Discussion
// 
// This method is the only way to remove a timer from an [NSRunLoop] object.
// The [NSRunLoop] object removes its strong reference to the timer, either
// just before the [Invalidate] method returns or at some later point.
// 
// If it was configured with target and user info objects, the receiver
// removes its strong references to those objects as well.
// 
// # Special Considerations
// 
// You must send this message from the thread on which the timer was
// installed. If you send this message from another thread, the input source
// associated with the timer may not be removed from its run loop, which could
// prevent the thread from exiting properly.
//
// See: https://developer.apple.com/documentation/Foundation/Timer/invalidate()
func (t Timer) Invalidate() {
	objc.Send[objc.ID](t.ID, objc.Sel("invalidate"))
}





// Creates a timer and schedules it on the current run loop in the default
// mode.
//
// interval: The number of seconds between firings of the timer. If `interval` is less
// than or equal to `0.0`, this method chooses the nonnegative value of
// `0.0001` seconds instead.
//
// repeats: If `true`, the timer will repeatedly reschedule itself until invalidated.
// If `false`, the timer will be invalidated after it fires.
//
// block: A block to be executed when the timer fires.
// 
// The block takes a single [NSTimer] parameter and has no return value.
//
// # Return Value
// 
// A new [NSTimer] object, configured according to the specified parameters.
//
// # Discussion
// 
// After `interval` seconds have elapsed, the timer fires, executing `block`.
//
// See: https://developer.apple.com/documentation/Foundation/Timer/scheduledTimer(withTimeInterval:repeats:block:)
func (_TimerClass TimerClass) ScheduledTimerWithTimeIntervalRepeatsBlock(interval float64, repeats bool, block TimerHandler) Timer {
		_block2, _cleanup2 := NewTimerBlock(block)
	defer _cleanup2()
		rv := objc.Send[objc.ID](objc.ID(_TimerClass.class), objc.Sel("scheduledTimerWithTimeInterval:repeats:block:"), interval, repeats, _block2)
	return NSTimerFromID(rv)
}

// Creates a timer and schedules it on the current run loop in the default
// mode.
//
// ti: The number of seconds between firings of the timer. If `ti` is less than or
// equal to `0.0`, this method chooses the nonnegative value of `0.0001`
// seconds instead.
//
// aTarget: The object to which to send the message specified by `aSelector` when the
// timer fires. The timer maintains a strong reference to `target` until it
// (the timer) is invalidated.
//
// aSelector: The message to send to `target` when the timer fires.
// 
// The selector should have the following signature: `` (including a colon to
// indicate that the method takes an argument). The timer passes itself as the
// argument, thus the method would adopt the following pattern:
//
// userInfo: The user info for the timer. The timer maintains a strong reference to this
// object until it (the timer) is invalidated. This parameter may be `nil`.
//
// yesOrNo: If [true], the timer will repeatedly reschedule itself until invalidated.
// If [false], the timer will be invalidated after it fires.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// A new [NSTimer] object, configured according to the specified parameters.
//
// # Discussion
// 
// After `ti` seconds have elapsed, the timer fires, sending the message
// `aSelector` to `target`.
//
// See: https://developer.apple.com/documentation/Foundation/Timer/scheduledTimer(timeInterval:target:selector:userInfo:repeats:)
func (_TimerClass TimerClass) ScheduledTimerWithTimeIntervalTargetSelectorUserInfoRepeats(ti float64, aTarget objectivec.IObject, aSelector objc.SEL, userInfo objectivec.IObject, yesOrNo bool) Timer {
	rv := objc.Send[objc.ID](objc.ID(_TimerClass.class), objc.Sel("scheduledTimerWithTimeInterval:target:selector:userInfo:repeats:"), ti, aTarget, aSelector, userInfo, yesOrNo)
	return NSTimerFromID(rv)
}

// Creates a new timer and schedules it on the current run loop in the default
// mode.
//
// ti: The number of seconds between firings of the timer. If `ti` is less than or
// equal to `0.0`, this method chooses the nonnegative value of `0.0001`
// seconds instead.
//
// invocation: The invocation to use when the timer fires. The invocation object maintains
// a strong reference to its arguments until the timer is invalidated.
//
// yesOrNo: If [true], the timer will repeatedly reschedule itself until invalidated.
// If [false], the timer will be invalidated after it fires.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// A new [NSTimer] object, configured according to the specified parameters.
//
// # Discussion
// 
// After `ti` seconds have elapsed, the timer fires, invoking `invocation`.
//
// See: https://developer.apple.com/documentation/Foundation/Timer/scheduledTimer(timeInterval:invocation:repeats:)
func (_TimerClass TimerClass) ScheduledTimerWithTimeIntervalInvocationRepeats(ti float64, invocation INSInvocation, yesOrNo bool) Timer {
	rv := objc.Send[objc.ID](objc.ID(_TimerClass.class), objc.Sel("scheduledTimerWithTimeInterval:invocation:repeats:"), ti, invocation, yesOrNo)
	return NSTimerFromID(rv)
}

// Initializes a timer object with the specified time interval and block.
//
// interval: The number of seconds between firings of the timer. If `interval` is less
// than or equal to `0.0`, this method chooses the nonnegative value of
// `0.0001` seconds instead.
//
// repeats: If `true`, the timer will repeatedly reschedule itself until invalidated.
// If `false`, the timer will be invalidated after it fires.
//
// block: A block to be executed when the timer fires. The block takes a single
// [NSTimer] parameter and has no return value.
//
// # Return Value
// 
// A new [NSTimer] object, configured according to the specified parameters.
//
// # Discussion
// 
// You must add the new timer to a run loop, using [AddTimerForMode]. Then,
// after `interval` seconds have elapsed, the timer fires, executing `block`.
// (If the timer is configured to repeat, you don’t need to add the timer to
// the run loop again.)
//
// See: https://developer.apple.com/documentation/Foundation/Timer/init(timeInterval:repeats:block:)
func (_TimerClass TimerClass) TimerWithTimeIntervalRepeatsBlock(interval float64, repeats bool, block TimerHandler) Timer {
		_block2, _cleanup2 := NewTimerBlock(block)
	defer _cleanup2()
		rv := objc.Send[objc.ID](objc.ID(_TimerClass.class), objc.Sel("timerWithTimeInterval:repeats:block:"), interval, repeats, _block2)
	return NSTimerFromID(rv)
}








// A Boolean value that indicates whether the timer is currently valid.
//
// # Discussion
// 
// [true] if the receiver is still capable of firing or [false] if the timer
// has been invalidated and is no longer capable of firing.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/Timer/isValid
func (t Timer) Valid() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isValid"))
	return rv
}



// The date at which the timer will fire.
//
// # Discussion
// 
// If the timer is no longer valid, the last date at which the timer fired.
// 
// You can set this property to adjust the firing time of a repeating timer.
// Although resetting a timer’s next firing time is a relatively expensive
// operation, it may be more efficient in some situations. For example, you
// could use it in situations where you want to repeat an action multiple
// times in the future, but at irregular time intervals. Adjusting the firing
// time of a single timer would likely incur less expense than creating
// multiple timer objects, scheduling each one on a run loop, and then
// destroying them.
// 
// You should not change the fire date of a timer that has been invalidated,
// which includes non-repeating timers that have already fired. You could
// potentially change the fire date of a non-repeating timer that had not yet
// fired, although you should always do so from the thread to which the timer
// is attached to avoid potential race conditions.
// 
// Use the [Valid] method to verify that the timer is valid.
//
// See: https://developer.apple.com/documentation/Foundation/Timer/fireDate
func (t Timer) FireDate() INSDate {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("fireDate"))
	return NSDateFromID(objc.ID(rv))
}
func (t Timer) SetFireDate(value INSDate) {
	objc.Send[struct{}](t.ID, objc.Sel("setFireDate:"), value)
}



// The timer’s time interval, in seconds.
//
// # Discussion
// 
// If the timer is non-repeating, returns `0` even if a time interval was set.
//
// See: https://developer.apple.com/documentation/Foundation/Timer/timeInterval
func (t Timer) TimeInterval() float64 {
	rv := objc.Send[NSTimeInterval](t.ID, objc.Sel("timeInterval"))
	return float64(rv)
}



// The receiver’s `userInfo` object.
//
// # Discussion
// 
// Do not access this property after the timer is invalidated. Use [Valid] to
// test whether the timer is valid.
//
// See: https://developer.apple.com/documentation/Foundation/Timer/userInfo
func (t Timer) UserInfo() objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("userInfo"))
	return objectivec.Object{ID: rv}
}



// The amount of time after the scheduled fire date that the timer may fire.
//
// # Discussion
// 
// The default value is zero, which means no additional tolerance is applied.
// 
// Setting a tolerance for a timer allows it to fire later than the scheduled
// fire date. Allowing the system flexibility in when a timer fires increases
// the ability of the system to optimize for increased power savings and
// responsiveness.
// 
// The timer may fire at any time between its scheduled fire date and the
// scheduled fire date plus the tolerance. The timer will not fire before the
// scheduled fire date. For repeating timers, the next fire date is calculated
// from the original fire date regardless of tolerance applied at individual
// fire times, to avoid drift. The system reserves the right to apply a small
// amount of tolerance to certain timers regardless of the value of this
// property.
//
// See: https://developer.apple.com/documentation/Foundation/Timer/tolerance
func (t Timer) Tolerance() float64 {
	rv := objc.Send[NSTimeInterval](t.ID, objc.Sel("tolerance"))
	return float64(rv)
}
func (t Timer) SetTolerance(value float64) {
	objc.Send[struct{}](t.ID, objc.Sel("setTolerance:"), value)
}


















// ScheduledTimerWithTimeIntervalRepeatsBlockSync is a synchronous wrapper around [Timer.ScheduledTimerWithTimeIntervalRepeatsBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (tc TimerClass) ScheduledTimerWithTimeIntervalRepeatsBlockSync(ctx context.Context, interval float64, repeats bool) (*NSTimer, error) {
	done := make(chan *NSTimer, 1)
	tc.ScheduledTimerWithTimeIntervalRepeatsBlock(interval, repeats, func(val *NSTimer) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// TimerWithTimeIntervalRepeatsBlockSync is a synchronous wrapper around [Timer.TimerWithTimeIntervalRepeatsBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (tc TimerClass) TimerWithTimeIntervalRepeatsBlockSync(ctx context.Context, interval float64, repeats bool) (*NSTimer, error) {
	done := make(chan *NSTimer, 1)
	tc.TimerWithTimeIntervalRepeatsBlock(interval, repeats, func(val *NSTimer) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// InitWithFireDateIntervalRepeatsBlockSync is a synchronous wrapper around [Timer.InitWithFireDateIntervalRepeatsBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (t Timer) InitWithFireDateIntervalRepeatsBlockSync(ctx context.Context, date INSDate, interval float64, repeats bool) (*NSTimer, error) {
	done := make(chan *NSTimer, 1)
	t.InitWithFireDateIntervalRepeatsBlock(date, interval, repeats, func(val *NSTimer) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}






