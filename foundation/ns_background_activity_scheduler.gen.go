// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSBackgroundActivityScheduler] class.
var (
	_NSBackgroundActivitySchedulerClass     NSBackgroundActivitySchedulerClass
	_NSBackgroundActivitySchedulerClassOnce sync.Once
)

func getNSBackgroundActivitySchedulerClass() NSBackgroundActivitySchedulerClass {
	_NSBackgroundActivitySchedulerClassOnce.Do(func() {
		_NSBackgroundActivitySchedulerClass = NSBackgroundActivitySchedulerClass{class: objc.GetClass("NSBackgroundActivityScheduler")}
	})
	return _NSBackgroundActivitySchedulerClass
}

// GetNSBackgroundActivitySchedulerClass returns the class object for NSBackgroundActivityScheduler.
func GetNSBackgroundActivitySchedulerClass() NSBackgroundActivitySchedulerClass {
	return getNSBackgroundActivitySchedulerClass()
}

type NSBackgroundActivitySchedulerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSBackgroundActivitySchedulerClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSBackgroundActivitySchedulerClass) Alloc() NSBackgroundActivityScheduler {
	rv := objc.Send[NSBackgroundActivityScheduler](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A task scheduler suitable for low priority operations that can run in the
// background.
//
// # Overview
//
// Use an [NSBackgroundActivityScheduler] object to schedule an arbitrary
// maintenance or background task. It’s similar to an [NSTimer] object, in
// that it lets you schedule a repeating or non-repeating task. However,
// [NSBackgroundActivityScheduler] gives the system flexibility to determine
// the most efficient time to execute based on energy usage, thermal
// conditions, and CPU use.
//
// For example, use an [NSBackgroundActivityScheduler] object to schedule:
//
// - Automatic saves - Backups - Data maintenance - Periodic content fetches -
// Installation of updates - Activities occurring in intervals of 10 minutes
// or more - Any other deferrable task
//
// For information about performing non-deferrable tasks efficiently, see
// [Specify Nondeferrable Background Activities] in [Energy Efficiency Guide
// for Mac Apps].
//
// # Create a Scheduler
//
// To initialize a scheduler, call [NSBackgroundActivityScheduler.InitWithIdentifier] for
// [NSBackgroundActivityScheduler], and pass it a unique identifier string in
// reverse DNS notation (`nil` and zero-length strings are not allowed) that
// remains constant across launches of your application.
//
// # Configure Scheduler Properties
//
// Configure the scheduler with any of the following scheduling properties:
//
// - [NSBackgroundActivityScheduler.Repeats]—If set to true, the activity is rescheduled at the specified
// interval after finishing. - [NSBackgroundActivityScheduler.Interval]—For repeating schedulers, the
// average interval between invocations of the activity. For nonrepeating
// schedulers, `interval` is the suggested interval of time between scheduling
// the activity and the invocation of the activity. - [NSBackgroundActivityScheduler.Tolerance]—The amount
// of time before or after the nominal fire date when the activity should be
// invoked. The nominal fire date is calculated by using the interval combined
// with the previous fire date or the time when the activity is started. These
// two properties create a window in time, during which the activity may be
// scheduled. The system will more aggressively schedule the activity as it
// nears the end of the grace period after the nominal fire date. The default
// value is half the interval. - [NSBackgroundActivityScheduler.QualityOfService]—The default value is
// [NSQualityOfServiceBackground]. If you upgrade the quality of service above
// this level, the system schedules the activity more aggressively. The
// default value is the recommended value for most activities. For information
// on quality of service, see [Prioritize Work at the Task Level] in [Energy
// Efficiency Guide for Mac Apps].
//
// The next three code examples demonstrate different scheduling scenarios.
//
// # Scheduling an activity to fire in the next 10 minutes
//
// # Scheduling an activity to fire between 15 and 45 minutes from now
//
// # Scheduling an activity to fire once each hour
//
// # Schedule Activity with scheduleWithBlock:
//
// When you’re ready to schedule the activity, call “ and provide a block
// of code to execute when the scheduler runs, as shown in the following
// example. The block will be called on a serial background queue appropriate
// for the level of quality of service specified. The system automatically
// uses the [BeginActivityWithOptionsReason] method (of [NSProcessInfo]) while
// invoking the block, choosing appropriate options based on the specified
// quality of service.
//
// When your block is called, it’s passed a completion handler as an
// argument. Configure the block to invoke this handler, passing it a result
// of type [NSBackgroundActivityScheduler.Result] to indicate whether the
// activity finished ([NSBackgroundActivityResultFinished]) or should be
// deferred ([NSBackgroundActivityResultDeferred]) and rescheduled for a later
// time. Failure to invoke the completion handler results in the activity not
// being rescheduled. For work that will be deferred and rescheduled, the
// block may optionally adjust scheduler properties, such as [NSBackgroundActivityScheduler.Interval] or
// [NSBackgroundActivityScheduler.Tolerance], before calling the completion handler.
//
// # Scheduling background activity
//
// # Detect Whether to Defer Activity
//
// It’s conceivable that while a lengthy activity is running, conditions may
// change, resulting in the activity now requiring deferral. For example,
// perhaps the user has unplugged the Mac and it’s now running on battery
// power. Your activity can call [NSBackgroundActivityScheduler.ShouldDefer] to determine whether this has
// occurred. A value of true indicates that the block should finish what
// it’s currently doing and invoke its completion handler with a value of
// [NSBackgroundActivityResultDeferred]. See the following example.
//
// # Detecting deferred background activity
//
// # Stop Activity
//
// Call [NSBackgroundActivityScheduler.Invalidate] to stop scheduling an activity, as shown in the following
// example.
//
// # Stopping background activity
//
// # Background Scheduler Attributes
//
//   - [NSBackgroundActivityScheduler.Identifier]: A unique reverse DNS notation string, such as `com.Example().MyApp.Updatecheck()`, that identifies the activity.
//   - [NSBackgroundActivityScheduler.Repeats]: A Boolean value indicating whether the activity should be rescheduled after it completes.
//   - [NSBackgroundActivityScheduler.SetRepeats]
//   - [NSBackgroundActivityScheduler.Interval]: An integer providing a suggested interval between scheduling and invoking the activity.
//   - [NSBackgroundActivityScheduler.SetInterval]
//   - [NSBackgroundActivityScheduler.QualityOfService]: A value of type [NSQualityOfService], which controls how aggressively the system schedules the activity.
//   - [NSBackgroundActivityScheduler.SetQualityOfService]
//   - [NSBackgroundActivityScheduler.ShouldDefer]: A Boolean value indicating whether your app should stop performing background activity and resume at a more optimal time.
//   - [NSBackgroundActivityScheduler.Tolerance]: A value of type [TimeInterval](<doc://com.apple.foundation/documentation/Foundation/TimeInterval>), which specifies a range of time during which the background activity may occur.
//   - [NSBackgroundActivityScheduler.SetTolerance]
//
// # Initializing Schedulers
//
//   - [NSBackgroundActivityScheduler.InitWithIdentifier]: Initializes a background activity scheduler object with a specified unique identifier.
//
// # Stopping Scheduled Activity
//
//   - [NSBackgroundActivityScheduler.Invalidate]: Prevents the background activity from being scheduled again.
//
// See: https://developer.apple.com/documentation/Foundation/NSBackgroundActivityScheduler
//
// [Energy Efficiency Guide for Mac Apps]: https://developer.apple.com/library/archive/documentation/Performance/Conceptual/power_efficiency_guidelines_osx/index.html#//apple_ref/doc/uid/TP40013929
// [NSBackgroundActivityScheduler.Result]: https://developer.apple.com/documentation/Foundation/NSBackgroundActivityScheduler/Result
// [Prioritize Work at the Task Level]: https://developer.apple.com/library/archive/documentation/Performance/Conceptual/power_efficiency_guidelines_osx/PrioritizeWorkAtTheTaskLevel.html#//apple_ref/doc/uid/TP40013929-CH35
// [Specify Nondeferrable Background Activities]: https://developer.apple.com/library/archive/documentation/Performance/Conceptual/power_efficiency_guidelines_osx/SchedulingBackgroundActivity.html#//apple_ref/doc/uid/TP40013929-CH32-SW10
type NSBackgroundActivityScheduler struct {
	objectivec.Object
}

// NSBackgroundActivitySchedulerFromID constructs a [NSBackgroundActivityScheduler] from an objc.ID.
//
// A task scheduler suitable for low priority operations that can run in the
// background.
func NSBackgroundActivitySchedulerFromID(id objc.ID) NSBackgroundActivityScheduler {
	return NSBackgroundActivityScheduler{objectivec.Object{ID: id}}
}

// NOTE: NSBackgroundActivityScheduler adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSBackgroundActivityScheduler] class.
//
// # Background Scheduler Attributes
//
//   - [INSBackgroundActivityScheduler.Identifier]: A unique reverse DNS notation string, such as `com.Example().MyApp.Updatecheck()`, that identifies the activity.
//   - [INSBackgroundActivityScheduler.Repeats]: A Boolean value indicating whether the activity should be rescheduled after it completes.
//   - [INSBackgroundActivityScheduler.SetRepeats]
//   - [INSBackgroundActivityScheduler.Interval]: An integer providing a suggested interval between scheduling and invoking the activity.
//   - [INSBackgroundActivityScheduler.SetInterval]
//   - [INSBackgroundActivityScheduler.QualityOfService]: A value of type [NSQualityOfService], which controls how aggressively the system schedules the activity.
//   - [INSBackgroundActivityScheduler.SetQualityOfService]
//   - [INSBackgroundActivityScheduler.ShouldDefer]: A Boolean value indicating whether your app should stop performing background activity and resume at a more optimal time.
//   - [INSBackgroundActivityScheduler.Tolerance]: A value of type [TimeInterval](<doc://com.apple.foundation/documentation/Foundation/TimeInterval>), which specifies a range of time during which the background activity may occur.
//   - [INSBackgroundActivityScheduler.SetTolerance]
//
// # Initializing Schedulers
//
//   - [INSBackgroundActivityScheduler.InitWithIdentifier]: Initializes a background activity scheduler object with a specified unique identifier.
//
// # Stopping Scheduled Activity
//
//   - [INSBackgroundActivityScheduler.Invalidate]: Prevents the background activity from being scheduled again.
//
// See: https://developer.apple.com/documentation/Foundation/NSBackgroundActivityScheduler
type INSBackgroundActivityScheduler interface {
	objectivec.IObject

	// Topic: Background Scheduler Attributes

	// A unique reverse DNS notation string, such as `com.Example().MyApp.Updatecheck()`, that identifies the activity.
	Identifier() string
	// A Boolean value indicating whether the activity should be rescheduled after it completes.
	Repeats() bool
	SetRepeats(value bool)
	// An integer providing a suggested interval between scheduling and invoking the activity.
	Interval() float64
	SetInterval(value float64)
	// A value of type [NSQualityOfService], which controls how aggressively the system schedules the activity.
	QualityOfService() QualityOfService
	SetQualityOfService(value QualityOfService)
	// A Boolean value indicating whether your app should stop performing background activity and resume at a more optimal time.
	ShouldDefer() bool
	// A value of type [TimeInterval](<doc://com.apple.foundation/documentation/Foundation/TimeInterval>), which specifies a range of time during which the background activity may occur.
	Tolerance() float64
	SetTolerance(value float64)

	// Topic: Initializing Schedulers

	// Initializes a background activity scheduler object with a specified unique identifier.
	InitWithIdentifier(identifier string) NSBackgroundActivityScheduler

	// Topic: Stopping Scheduled Activity

	// Prevents the background activity from being scheduled again.
	Invalidate()
}

// Init initializes the instance.
func (b NSBackgroundActivityScheduler) Init() NSBackgroundActivityScheduler {
	rv := objc.Send[NSBackgroundActivityScheduler](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b NSBackgroundActivityScheduler) Autorelease() NSBackgroundActivityScheduler {
	rv := objc.Send[NSBackgroundActivityScheduler](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSBackgroundActivityScheduler creates a new NSBackgroundActivityScheduler instance.
func NewNSBackgroundActivityScheduler() NSBackgroundActivityScheduler {
	class := getNSBackgroundActivitySchedulerClass()
	rv := objc.Send[NSBackgroundActivityScheduler](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a background activity scheduler object with a specified unique
// identifier.
//
// identifier: A unique string, in reverse DNS notation, that identifies the activity. For
// example, `com.Example().MyApp.Updatecheck()`. `nil` and zero-length strings
// are not allowed.
//
// # Return Value
//
// A new background activity scheduler object of type
// [NSBackgroundActivityScheduler].
//
// # Discussion
//
// The string passed to the `identifier` parameter should remain constant for
// an activity across launches of your app because the system uses this unique
// identifier to track the number of times the activity has run and to improve
// the heuristics for deciding when to run it again in the future. See
// [NSBackgroundActivityScheduler].
//
// See: https://developer.apple.com/documentation/Foundation/NSBackgroundActivityScheduler/init(identifier:)
func NewBackgroundActivitySchedulerWithIdentifier(identifier string) NSBackgroundActivityScheduler {
	instance := getNSBackgroundActivitySchedulerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIdentifier:"), objc.String(identifier))
	return NSBackgroundActivitySchedulerFromID(rv)
}

// Initializes a background activity scheduler object with a specified unique
// identifier.
//
// identifier: A unique string, in reverse DNS notation, that identifies the activity. For
// example, `com.Example().MyApp.Updatecheck()`. `nil` and zero-length strings
// are not allowed.
//
// # Return Value
//
// A new background activity scheduler object of type
// [NSBackgroundActivityScheduler].
//
// # Discussion
//
// The string passed to the `identifier` parameter should remain constant for
// an activity across launches of your app because the system uses this unique
// identifier to track the number of times the activity has run and to improve
// the heuristics for deciding when to run it again in the future. See
// [NSBackgroundActivityScheduler].
//
// See: https://developer.apple.com/documentation/Foundation/NSBackgroundActivityScheduler/init(identifier:)
func (b NSBackgroundActivityScheduler) InitWithIdentifier(identifier string) NSBackgroundActivityScheduler {
	rv := objc.Send[NSBackgroundActivityScheduler](b.ID, objc.Sel("initWithIdentifier:"), objc.String(identifier))
	return rv
}

// Prevents the background activity from being scheduled again.
//
// # Discussion
//
// When `invalidate` is used to stop an activity that is currently executing,
// the activity will still finish executing.
//
// See [NSBackgroundActivityScheduler].
//
// See: https://developer.apple.com/documentation/Foundation/NSBackgroundActivityScheduler/invalidate()
func (b NSBackgroundActivityScheduler) Invalidate() {
	objc.Send[objc.ID](b.ID, objc.Sel("invalidate"))
}

// A unique reverse DNS notation string, such as
// `com.Example().MyApp.Updatecheck()`, that identifies the activity.
//
// # Discussion
//
// This string should remain constant for an activity across launches of your
// app because the system uses this unique identifier to track the number of
// times the activity has run and to improve the heuristics for deciding when
// to run it again in the future. `nil` and zero-length strings are not
// allowed. See [NSBackgroundActivityScheduler].
//
// See: https://developer.apple.com/documentation/Foundation/NSBackgroundActivityScheduler/identifier
func (b NSBackgroundActivityScheduler) Identifier() string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("identifier"))
	return NSStringFromID(rv).String()
}

// A Boolean value indicating whether the activity should be rescheduled after
// it completes.
//
// # Discussion
//
// The default value for this property is false. See
// [NSBackgroundActivityScheduler].
//
// See: https://developer.apple.com/documentation/Foundation/NSBackgroundActivityScheduler/repeats
func (b NSBackgroundActivityScheduler) Repeats() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("repeats"))
	return rv
}
func (b NSBackgroundActivityScheduler) SetRepeats(value bool) {
	objc.Send[struct{}](b.ID, objc.Sel("setRepeats:"), value)
}

// An integer providing a suggested interval between scheduling and invoking
// the activity.
//
// # Discussion
//
// For repeating activities, the value of this property is also the suggested
// interval between invocations. See [NSBackgroundActivityScheduler].
//
// See: https://developer.apple.com/documentation/Foundation/NSBackgroundActivityScheduler/interval
func (b NSBackgroundActivityScheduler) Interval() float64 {
	rv := objc.Send[NSTimeInterval](b.ID, objc.Sel("interval"))
	return float64(rv)
}
func (b NSBackgroundActivityScheduler) SetInterval(value float64) {
	objc.Send[struct{}](b.ID, objc.Sel("setInterval:"), value)
}

// A value of type [NSQualityOfService], which controls how aggressively the
// system schedules the activity.
//
// # Discussion
//
// Options include:
//
// - NSQualityOfServiceUserInteractive - NSQualityOfServiceUserInitiated -
// NSQualityOfServiceUtility - NSQualityOfServiceBackground
//
// The default value is [NSQualityOfServiceBackground]. If you upgrade the
// quality of service above this level, the system schedules the activity more
// aggressively. The default value is the recommended value for most
// activities. See [NSBackgroundActivityScheduler]. For information about
// quality of service, see [Prioritize Work at the Task Level] in [Energy
// Efficiency Guide for Mac Apps].
//
// See: https://developer.apple.com/documentation/Foundation/NSBackgroundActivityScheduler/qualityOfService
//
// [Energy Efficiency Guide for Mac Apps]: https://developer.apple.com/library/archive/documentation/Performance/Conceptual/power_efficiency_guidelines_osx/index.html#//apple_ref/doc/uid/TP40013929
// [Prioritize Work at the Task Level]: https://developer.apple.com/library/archive/documentation/Performance/Conceptual/power_efficiency_guidelines_osx/PrioritizeWorkAtTheTaskLevel.html#//apple_ref/doc/uid/TP40013929-CH35
func (b NSBackgroundActivityScheduler) QualityOfService() QualityOfService {
	rv := objc.Send[QualityOfService](b.ID, objc.Sel("qualityOfService"))
	return QualityOfService(rv)
}
func (b NSBackgroundActivityScheduler) SetQualityOfService(value QualityOfService) {
	objc.Send[struct{}](b.ID, objc.Sel("setQualityOfService:"), value)
}

// A Boolean value indicating whether your app should stop performing
// background activity and resume at a more optimal time.
//
// # Discussion
//
// Your app can check the `shouldDefer` property while executing scheduled
// background activity. If this property contains a value of true, system
// conditions have changed since the time the activity started and deferral is
// recommended. For example, perhaps the user unplugged the Mac and it’s now
// running on battery power. In this case, your app should finish what it’s
// currently doing, save its state, and invoke its completion handler with a
// value of [NSBackgroundActivityResultDeferred]. The system will invoke your
// activity again at a more optimal time, and your app can restore its
// previous state and resume where it left off. See
// [NSBackgroundActivityScheduler] and [NSBackgroundActivityScheduler].
//
// See: https://developer.apple.com/documentation/Foundation/NSBackgroundActivityScheduler/shouldDefer
func (b NSBackgroundActivityScheduler) ShouldDefer() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("shouldDefer"))
	return rv
}

// A value of type [NSTimeInterval], which specifies a range of time during
// which the background activity may occur.
//
// # Discussion
//
// A nominal fire date for scheduled background activity is calculated based
// on a combination of the [Interval] property value and the time the activity
// began or the last execution date. The [Tolerance] property specifies a
// grace period—a range of time before and after the nominal fire date,
// during which the activity may be invoked. As the activity nears the end of
// its grace period, the system schedules the activity more aggressively. The
// default tolerance period is half the value of the [Interval] property. See
// [NSBackgroundActivityScheduler].
//
// See: https://developer.apple.com/documentation/Foundation/NSBackgroundActivityScheduler/tolerance
func (b NSBackgroundActivityScheduler) Tolerance() float64 {
	rv := objc.Send[NSTimeInterval](b.ID, objc.Sel("tolerance"))
	return float64(rv)
}
func (b NSBackgroundActivityScheduler) SetTolerance(value float64) {
	objc.Send[struct{}](b.ID, objc.Sel("setTolerance:"), value)
}
