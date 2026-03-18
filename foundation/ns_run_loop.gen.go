// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [RunLoop] class.
var (
	_RunLoopClass     RunLoopClass
	_RunLoopClassOnce sync.Once
)

func getRunLoopClass() RunLoopClass {
	_RunLoopClassOnce.Do(func() {
		_RunLoopClass = RunLoopClass{class: objc.GetClass("NSRunLoop")}
	})
	return _RunLoopClass
}

// GetRunLoopClass returns the class object for NSRunLoop.
func GetRunLoopClass() RunLoopClass {
	return getRunLoopClass()
}

type RunLoopClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (rc RunLoopClass) Alloc() RunLoop {
	rv := objc.Send[RunLoop](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// The programmatic interface to objects that manage input sources.
//
// # Overview
// 
// A [NSRunLoop] object processes input for sources, such as mouse and
// keyboard events from the window system and [NSPort] objects. A [NSRunLoop]
// object also processes [NSTimer] events.
// 
// Your application neither creates nor explicitly manages [NSRunLoop]
// objects. The system creates a [NSRunLoop] object as needed for each
// [NSThread] object, including the application’s main thread. If you need
// to access the current thread’s run loop, use the class method
// [CurrentRunLoop].
// 
// Note that from the perspective of [NSRunLoop], [NSTimer] objects aren’t
// “input”—they’re a special type, and they don’t cause the run loop
// to return when they fire.
//
// # Accessing Run Loops and Modes
//
//   - [RunLoop.CurrentMode]: The receiver’s current input mode.
//   - [RunLoop.LimitDateForMode]: Performs one pass through the run loop in the specified mode and returns the date at which the next timer is scheduled to fire.
//   - [RunLoop.GetCFRunLoop]: Returns the receiver’s underlying run loop object.
//
// # Managing Timers
//
//   - [RunLoop.AddTimerForMode]: Registers a given timer with a given input mode.
//
// # Managing Ports
//
//   - [RunLoop.AddPortForMode]: Adds a port as an input source to the specified mode of the run loop.
//   - [RunLoop.RemovePortForMode]: Removes a port from the specified input mode of the run loop.
//
// # Running a Loop
//
//   - [RunLoop.Run]: Puts the receiver into a permanent loop, during which time it processes data from all attached input sources.
//   - [RunLoop.RunModeBeforeDate]: Runs the loop once, blocking for input in the specified mode until a given date.
//   - [RunLoop.RunUntilDate]: Runs the loop until the specified date, during which time it processes data from all attached input sources.
//   - [RunLoop.AcceptInputForModeBeforeDate]: Runs the loop once or until the specified date, accepting input only for the specified mode.
//
// # Scheduling and Canceling Tasks
//
//   - [RunLoop.PerformBlock]: Schedules a block that the run loop invokes.
//   - [RunLoop.PerformInModesBlock]: Schedules a block that the run loop invokes when it’s running in any of the specified modes.
//   - [RunLoop.PerformSelectorTargetArgumentOrderModes]: Schedules the sending of a message on the receiver.
//   - [RunLoop.CancelPerformSelectorTargetArgument]: Cancels the sending of a previously scheduled message.
//   - [RunLoop.CancelPerformSelectorsWithTarget]: Cancels all outstanding ordered performs scheduled with a given target.
//
// See: https://developer.apple.com/documentation/Foundation/RunLoop
type RunLoop struct {
	objectivec.Object
}

// RunLoopFromID constructs a [RunLoop] from an objc.ID.
//
// The programmatic interface to objects that manage input sources.
func RunLoopFromID(id objc.ID) RunLoop {
	return NSRunLoop{objectivec.Object{ID: id}}
}

// NSRunLoopFromID is an alias for [RunLoopFromID] for cross-framework compatibility.
func NSRunLoopFromID(id objc.ID) RunLoop { return RunLoopFromID(id) }
// NOTE: RunLoop adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [RunLoop] class.
//
// # Accessing Run Loops and Modes
//
//   - [IRunLoop.CurrentMode]: The receiver’s current input mode.
//   - [IRunLoop.LimitDateForMode]: Performs one pass through the run loop in the specified mode and returns the date at which the next timer is scheduled to fire.
//   - [IRunLoop.GetCFRunLoop]: Returns the receiver’s underlying run loop object.
//
// # Managing Timers
//
//   - [IRunLoop.AddTimerForMode]: Registers a given timer with a given input mode.
//
// # Managing Ports
//
//   - [IRunLoop.AddPortForMode]: Adds a port as an input source to the specified mode of the run loop.
//   - [IRunLoop.RemovePortForMode]: Removes a port from the specified input mode of the run loop.
//
// # Running a Loop
//
//   - [IRunLoop.Run]: Puts the receiver into a permanent loop, during which time it processes data from all attached input sources.
//   - [IRunLoop.RunModeBeforeDate]: Runs the loop once, blocking for input in the specified mode until a given date.
//   - [IRunLoop.RunUntilDate]: Runs the loop until the specified date, during which time it processes data from all attached input sources.
//   - [IRunLoop.AcceptInputForModeBeforeDate]: Runs the loop once or until the specified date, accepting input only for the specified mode.
//
// # Scheduling and Canceling Tasks
//
//   - [IRunLoop.PerformBlock]: Schedules a block that the run loop invokes.
//   - [IRunLoop.PerformInModesBlock]: Schedules a block that the run loop invokes when it’s running in any of the specified modes.
//   - [IRunLoop.PerformSelectorTargetArgumentOrderModes]: Schedules the sending of a message on the receiver.
//   - [IRunLoop.CancelPerformSelectorTargetArgument]: Cancels the sending of a previously scheduled message.
//   - [IRunLoop.CancelPerformSelectorsWithTarget]: Cancels all outstanding ordered performs scheduled with a given target.
//
// See: https://developer.apple.com/documentation/Foundation/RunLoop
type IRunLoop interface {
	objectivec.IObject

	// Topic: Accessing Run Loops and Modes

	// The receiver’s current input mode.
	CurrentMode() NSRunLoopMode
	// Performs one pass through the run loop in the specified mode and returns the date at which the next timer is scheduled to fire.
	LimitDateForMode(mode NSRunLoopMode) INSDate
	// Returns the receiver’s underlying run loop object.
	GetCFRunLoop() corefoundation.CFRunLoopRef

	// Topic: Managing Timers

	// Registers a given timer with a given input mode.
	AddTimerForMode(timer INSTimer, mode NSRunLoopMode)

	// Topic: Managing Ports

	// Adds a port as an input source to the specified mode of the run loop.
	AddPortForMode(aPort INSPort, mode NSRunLoopMode)
	// Removes a port from the specified input mode of the run loop.
	RemovePortForMode(aPort INSPort, mode NSRunLoopMode)

	// Topic: Running a Loop

	// Puts the receiver into a permanent loop, during which time it processes data from all attached input sources.
	Run()
	// Runs the loop once, blocking for input in the specified mode until a given date.
	RunModeBeforeDate(mode NSRunLoopMode, limitDate INSDate) bool
	// Runs the loop until the specified date, during which time it processes data from all attached input sources.
	RunUntilDate(limitDate INSDate)
	// Runs the loop once or until the specified date, accepting input only for the specified mode.
	AcceptInputForModeBeforeDate(mode NSRunLoopMode, limitDate INSDate)

	// Topic: Scheduling and Canceling Tasks

	// Schedules a block that the run loop invokes.
	PerformBlock(block VoidHandler)
	// Schedules a block that the run loop invokes when it’s running in any of the specified modes.
	PerformInModesBlock(modes []string, block VoidHandler)
	// Schedules the sending of a message on the receiver.
	PerformSelectorTargetArgumentOrderModes(aSelector objc.SEL, target objectivec.IObject, arg objectivec.IObject, order uint, modes []string)
	// Cancels the sending of a previously scheduled message.
	CancelPerformSelectorTargetArgument(aSelector objc.SEL, target objectivec.IObject, arg objectivec.IObject)
	// Cancels all outstanding ordered performs scheduled with a given target.
	CancelPerformSelectorsWithTarget(target objectivec.IObject)
}

// Init initializes the instance.
func (r RunLoop) Init() RunLoop {
	rv := objc.Send[RunLoop](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r RunLoop) Autorelease() RunLoop {
	rv := objc.Send[RunLoop](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewRunLoop creates a new RunLoop instance.
func NewRunLoop() RunLoop {
	class := getRunLoopClass()
	rv := objc.Send[RunLoop](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Performs one pass through the run loop in the specified mode and returns
// the date at which the next timer is scheduled to fire.
//
// mode: The run loop mode to search. You may specify custom modes or use one of the
// modes listed in `Run Loop Modes`.
//
// # Return Value
// 
// The date at which the next timer is scheduled to fire, or `nil` if there
// are no input sources for this mode.
//
// # Discussion
// 
// The run loop is entered with an immediate timeout, so the run loop does not
// block, waiting for input, if no input sources need processing.
//
// See: https://developer.apple.com/documentation/Foundation/RunLoop/limitDate(forMode:)
func (r RunLoop) LimitDateForMode(mode NSRunLoopMode) INSDate {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("limitDateForMode:"), objc.String(string(mode)))
	return NSDateFromID(rv)
}

// Returns the receiver’s underlying run loop object.
//
// # Return Value
// 
// The receiver’s underlying [CFRunLoop] object.
//
// [CFRunLoop]: https://developer.apple.com/documentation/CoreFoundation/CFRunLoop
//
// # Discussion
// 
// You can use the returned run loop to configure the current run loop using
// Core Foundation function calls. For example, you might use this function to
// set up a run loop observer.
//
// See: https://developer.apple.com/documentation/Foundation/RunLoop/getCFRunLoop()
func (r RunLoop) GetCFRunLoop() corefoundation.CFRunLoopRef {
	rv := objc.Send[corefoundation.CFRunLoopRef](r.ID, objc.Sel("getCFRunLoop"))
	return corefoundation.CFRunLoopRef(rv)
}

// Registers a given timer with a given input mode.
//
// timer: The timer to register with the receiver.
//
// mode: The mode in which to add `aTimer`. You may specify a custom mode or use one
// of the modes listed in `Run Loop Modes`.
//
// # Discussion
// 
// You can add a timer to multiple input modes. While running in the
// designated mode, the receiver causes the timer to fire on or after its
// scheduled fire date. Upon firing, the timer invokes its associated handler
// routine, which is a selector on a designated object.
// 
// The receiver retains `aTimer`. To remove a timer from all run loop modes on
// which it is installed, send an [Invalidate] message to the timer.
//
// See: https://developer.apple.com/documentation/Foundation/RunLoop/add(_:forMode:)-392ag
func (r RunLoop) AddTimerForMode(timer INSTimer, mode NSRunLoopMode) {
	objc.Send[objc.ID](r.ID, objc.Sel("addTimer:forMode:"), timer, objc.String(string(mode)))
}

// Adds a port as an input source to the specified mode of the run loop.
//
// aPort: The port to add to the receiver.
//
// mode: The mode in which to add `aPort`. You may specify a custom mode or use one
// of the modes listed in `Run Loop Modes`.
//
// # Discussion
// 
// This method schedules the port with the receiver. You can add a port to
// multiple input modes. When the receiver is running in the specified mode,
// it dispatches messages destined for that port to the port’s designated
// handler routine.
//
// See: https://developer.apple.com/documentation/Foundation/RunLoop/add(_:forMode:)-6z982
func (r RunLoop) AddPortForMode(aPort INSPort, mode NSRunLoopMode) {
	objc.Send[objc.ID](r.ID, objc.Sel("addPort:forMode:"), aPort, objc.String(string(mode)))
}

// Removes a port from the specified input mode of the run loop.
//
// aPort: The port to remove from the receiver.
//
// mode: The mode from which to remove `aPort`. You may specify a custom mode or use
// one of the modes listed in `Run Loop Modes`.
//
// # Discussion
// 
// If you added the port to multiple input modes, you must remove it from each
// mode separately.
//
// See: https://developer.apple.com/documentation/Foundation/RunLoop/remove(_:forMode:)
func (r RunLoop) RemovePortForMode(aPort INSPort, mode NSRunLoopMode) {
	objc.Send[objc.ID](r.ID, objc.Sel("removePort:forMode:"), aPort, objc.String(string(mode)))
}

// Puts the receiver into a permanent loop, during which time it processes
// data from all attached input sources.
//
// # Discussion
// 
// If no input sources or timers are attached to the run loop, this method
// exits immediately; otherwise, it runs the receiver in the
// [NSDefaultRunLoopMode] by repeatedly invoking [RunModeBeforeDate]. In other
// words, this method effectively begins an infinite loop that processes data
// from the run loop’s input sources and timers.
// 
// Manually removing all known input sources and timers from the run loop is
// not a guarantee that the run loop will exit. macOS can install and remove
// additional input sources as needed to process requests targeted at the
// receiver’s thread. Those sources could therefore prevent the run loop
// from exiting.
// 
// If you want the run loop to terminate, you shouldn’t use this method.
// Instead, use one of the other run methods and also check other arbitrary
// conditions of your own, in a loop. A simple example would be:
// 
// where `shouldKeepRunning` is set to [false] somewhere else in the program.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/Foundation/RunLoop/run()
func (r RunLoop) Run() {
	objc.Send[objc.ID](r.ID, objc.Sel("run"))
}

// Runs the loop once, blocking for input in the specified mode until a given
// date.
//
// mode: The mode in which to run. You may specify custom modes or use one of the
// modes listed in `Run Loop Modes`.
//
// limitDate: The date until which to block.
//
// # Return Value
// 
// [true] if the run loop ran and processed an input source or if the
// specified timeout value was reached; otherwise, [false] if the run loop
// could not be started.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// If no input sources or timers are attached to the run loop, this method
// exits immediately and returns [false]; otherwise, it returns after either
// the first input source is processed or `limitDate` is reached. Manually
// removing all known input sources and timers from the run loop does not
// guarantee that the run loop will exit immediately. macOS may install and
// remove additional input sources as needed to process requests targeted at
// the receiver’s thread. Those sources could therefore prevent the run loop
// from exiting.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/Foundation/RunLoop/run(mode:before:)
func (r RunLoop) RunModeBeforeDate(mode NSRunLoopMode, limitDate INSDate) bool {
	rv := objc.Send[bool](r.ID, objc.Sel("runMode:beforeDate:"), objc.String(string(mode)), limitDate)
	return rv
}

// Runs the loop until the specified date, during which time it processes data
// from all attached input sources.
//
// limitDate: The date up until which to run.
//
// # Discussion
// 
// If no input sources or timers are attached to the run loop, this method
// exits immediately; otherwise, it runs the receiver in the
// [NSDefaultRunLoopMode] by repeatedly invoking [RunModeBeforeDate] until the
// specified expiration date.
// 
// Manually removing all known input sources and timers from the run loop is
// not a guarantee that the run loop will exit. macOS can install and remove
// additional input sources as needed to process requests targeted at the
// receiver’s thread. Those sources could therefore prevent the run loop
// from exiting.
//
// See: https://developer.apple.com/documentation/Foundation/RunLoop/run(until:)
func (r RunLoop) RunUntilDate(limitDate INSDate) {
	objc.Send[objc.ID](r.ID, objc.Sel("runUntilDate:"), limitDate)
}

// Runs the loop once or until the specified date, accepting input only for
// the specified mode.
//
// mode: The mode in which to run. You may specify custom modes or use one of the
// modes listed in `Run Loop Modes`.
//
// limitDate: The date up until which to run.
//
// # Discussion
// 
// If no input sources or timers are attached to the run loop, this method
// exits immediately; otherwise, it runs the run loop once, returning as soon
// as one input source processes a message or the specifed time elapses.
// 
// Manually removing all known input sources and timers from the run loop is
// not a guarantee that the run loop will exit. macOS can install and remove
// additional input sources as needed to process requests targeted at the
// receiver’s thread. Those sources could therefore prevent the run loop
// from exiting.
//
// See: https://developer.apple.com/documentation/Foundation/RunLoop/acceptInput(forMode:before:)
func (r RunLoop) AcceptInputForModeBeforeDate(mode NSRunLoopMode, limitDate INSDate) {
	objc.Send[objc.ID](r.ID, objc.Sel("acceptInputForMode:beforeDate:"), objc.String(string(mode)), limitDate)
}

// Schedules a block that the run loop invokes.
//
// block: A block that the run loop invokes.
//
// See: https://developer.apple.com/documentation/Foundation/RunLoop/perform(_:)
func (r RunLoop) PerformBlock(block VoidHandler) {
_block0, _cleanup0 := NewVoidBlock(block)
	defer _cleanup0()
	objc.Send[objc.ID](r.ID, objc.Sel("performBlock:"), _block0)
}

// Schedules a block that the run loop invokes when it’s running in any of
// the specified modes.
//
// modes: An array of modes in which the run loop invokes the block.
//
// block: The block that the run loop invokes.
//
// See: https://developer.apple.com/documentation/Foundation/RunLoop/perform(inModes:block:)
func (r RunLoop) PerformInModesBlock(modes []string, block VoidHandler) {
_block1, _cleanup1 := NewVoidBlock(block)
	defer _cleanup1()
	objc.Send[objc.ID](r.ID, objc.Sel("performInModes:block:"), modes, _block1)
}

// Schedules the sending of a message on the receiver.
//
// aSelector: A selector that identifies the method to invoke. This method should not
// have a significant return value and should take a single argument of type
// id.
//
// target: The object that defines the selector in `aSelector`.
//
// arg: The argument to pass to the method when it is invoked. Pass `nil` if the
// method does not take an argument.
//
// order: The priority for the message. If multiple messages are scheduled, the
// messages with a lower order value are sent before messages with a higher
// order value.
//
// modes: An array of input modes for which the message may be sent. You may specify
// custom modes or use one of the modes listed in `Run Loop Modes`.
//
// # Discussion
// 
// This method sets up a timer to perform the `aSelector` message on the
// receiver at the start of the next run loop iteration. The timer is
// configured to run in the modes specified by the `modes` parameter. When the
// timer fires, the thread attempts to dequeue the message from the run loop
// and perform the selector. It succeeds if the run loop is running and in one
// of the specified modes; otherwise, the timer waits until the run loop is in
// one of those modes.
// 
// This method returns before the `aSelector` message is sent. The receiver
// retains the `target` and `anArgument` objects until the timer for the
// selector fires, and then releases them as part of its cleanup.
// 
// Use this method if you want multiple messages to be sent after the current
// event has been processed and you want to make sure these messages are sent
// in a certain order.
//
// See: https://developer.apple.com/documentation/Foundation/RunLoop/perform(_:target:argument:order:modes:)
func (r RunLoop) PerformSelectorTargetArgumentOrderModes(aSelector objc.SEL, target objectivec.IObject, arg objectivec.IObject, order uint, modes []string) {
	objc.Send[objc.ID](r.ID, objc.Sel("performSelector:target:argument:order:modes:"), aSelector, target, arg, order, objectivec.StringSliceToNSArray(modes))
}

// Cancels the sending of a previously scheduled message.
//
// aSelector: The previously-specified selector.
//
// target: The previously-specified target.
//
// arg: The previously-specified argument.
//
// # Discussion
// 
// You can use this method to cancel a message previously scheduled using the
// [PerformSelectorTargetArgumentOrderModes] method. The parameters identify
// the message you want to cancel and must match those originally specified
// when the selector was scheduled. This method removes the perform request
// from all modes of the run loop.
//
// See: https://developer.apple.com/documentation/Foundation/RunLoop/cancelPerform(_:target:argument:)
func (r RunLoop) CancelPerformSelectorTargetArgument(aSelector objc.SEL, target objectivec.IObject, arg objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("cancelPerformSelector:target:argument:"), aSelector, target, arg)
}

// Cancels all outstanding ordered performs scheduled with a given target.
//
// target: The previously-specified target.
//
// # Discussion
// 
// This method cancels the previously scheduled messages associated with the
// target, ignoring the selector and argument of the scheduled operation. This
// is in contrast to [CancelPerformSelectorTargetArgument], which requires you
// to match the selector and argument as well as the target. This method
// removes the perform requests for the object from all modes of the run loop.
//
// See: https://developer.apple.com/documentation/Foundation/RunLoop/cancelPerformSelectors(withTarget:)
func (r RunLoop) CancelPerformSelectorsWithTarget(target objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("cancelPerformSelectorsWithTarget:"), target)
}

// The receiver’s current input mode.
//
// # Discussion
// 
// The receiver’s current input mode. This method returns the current input
// mode while the receiver is running; otherwise, it returns `nil`.
// 
// The current mode is set by the methods that run the run loop, such as
// [AcceptInputForModeBeforeDate] and [RunModeBeforeDate].
//
// See: https://developer.apple.com/documentation/Foundation/RunLoop/currentMode
func (r RunLoop) CurrentMode() NSRunLoopMode {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("currentMode"))
	return NSRunLoopMode(NSStringFromID(rv).String())
}

// Returns the run loop for the current thread.
//
// # Return Value
// 
// The [NSRunLoop] object for the current thread.
// 
// # Discussion
// 
// If a run loop does not yet exist for the thread, one is created and
// returned.
//
// See: https://developer.apple.com/documentation/Foundation/RunLoop/current
func (_RunLoopClass RunLoopClass) CurrentRunLoop() RunLoop {
	rv := objc.Send[objc.ID](objc.ID(_RunLoopClass.class), objc.Sel("currentRunLoop"))
	return NSRunLoopFromID(objc.ID(rv))
}

// Returns the run loop of the main thread.
//
// # Return Value
// 
// An object representing the main thread’s run loop.
//
// See: https://developer.apple.com/documentation/Foundation/RunLoop/main
func (_RunLoopClass RunLoopClass) MainRunLoop() RunLoop {
	rv := objc.Send[objc.ID](objc.ID(_RunLoopClass.class), objc.Sel("mainRunLoop"))
	return NSRunLoopFromID(objc.ID(rv))
}

// PerformBlockSync is a synchronous wrapper around [RunLoop.PerformBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (r RunLoop) PerformBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	r.PerformBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

