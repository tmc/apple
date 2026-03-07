// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [Thread] class.
var (
	_ThreadClass     ThreadClass
	_ThreadClassOnce sync.Once
)

func getThreadClass() ThreadClass {
	_ThreadClassOnce.Do(func() {
		_ThreadClass = ThreadClass{class: objc.GetClass("NSThread")}
	})
	return _ThreadClass
}

// GetThreadClass returns the class object for NSThread.
func GetThreadClass() ThreadClass {
	return getThreadClass()
}

type ThreadClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (tc ThreadClass) Alloc() Thread {
	rv := objc.Send[Thread](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}







// A thread of execution.
//
// # Overview
// 
// Use this class when you want to have an Objective-C method run in its own
// thread of execution. Threads are especially useful when you need to perform
// a lengthy task, but don’t want it to block the execution of the rest of
// the application. In particular, you can use threads to avoid blocking the
// main thread of the application, which handles user interface and
// event-related actions. Threads can also be used to divide a large job into
// several smaller jobs, which can lead to performance increases on multi-core
// computers.
// 
// The [NSThread] class supports semantics similar to those of [NSOperation]
// for monitoring the runtime condition of a thread. You can use these
// semantics to cancel the execution of a thread or determine if the thread is
// still executing or has finished its task. Canceling a thread requires
// support from your thread code; see the description for [Cancel] for more
// information.
// 
// # Subclassing Notes
// 
// You can subclass [NSThread] and override the [Main] method to implement
// your thread’s main entry point. If you override [Main], you do not need
// to invoke the inherited behavior by calling `super`.
//
// # Initializing an NSThread Object
//
//   - [Thread.InitWithTargetSelectorObject]: Returns an [NSThread] object initialized with the given arguments.
//
// # Starting a Thread
//
//   - [Thread.Start]: Starts the receiver.
//   - [Thread.Main]: The main entry point routine for the thread.
//
// # Stopping a Thread
//
//   - [Thread.Cancel]: Changes the cancelled state of the receiver to indicate that it should exit.
//
// # Determining the Thread’s Execution State
//
//   - [Thread.Executing]: A Boolean value that indicates whether the receiver is executing.
//   - [Thread.Finished]: A Boolean value that indicates whether the receiver has finished execution.
//   - [Thread.Cancelled]: A Boolean value that indicates whether the receiver is cancelled.
//
// # Working with the Main Thread
//
//   - [Thread.IsMainThread]: A Boolean value that indicates whether the receiver is the main thread.
//
// # Working with Thread Properties
//
//   - [Thread.ThreadDictionary]: The thread object’s dictionary.
//   - [Thread.NSAssertionHandlerKey]: A key with a corresponding value in the thread dictionary.
//   - [Thread.Name]: The name of the receiver.
//   - [Thread.SetName]
//   - [Thread.StackSize]: The stack size of the receiver, in bytes.
//   - [Thread.SetStackSize]
//
// # Prioritizing Thread Work
//
//   - [Thread.QualityOfService]
//   - [Thread.SetQualityOfService]
//   - [Thread.ThreadPriority]: The receiver’s priority
//   - [Thread.SetThreadPriority]
//
// # Notifications
//
//   - [Thread.NSDidBecomeSingleThreaded]: Not implemented.
//   - [Thread.NSThreadWillExit]: An 
//   - [Thread.NSWillBecomeMultiThreaded]: Posted when the first thread is detached from the current thread. The 
//
// # Initializers
//
//   - [Thread.InitWithBlock]
//
// See: https://developer.apple.com/documentation/Foundation/Thread
type Thread struct {
	objectivec.Object
}

// ThreadFromID constructs a [Thread] from an objc.ID.
//
// A thread of execution.
func ThreadFromID(id objc.ID) Thread {
	return NSThread{objectivec.Object{id}}
}

// NSThreadFromID is an alias for [ThreadFromID] for cross-framework compatibility.
func NSThreadFromID(id objc.ID) Thread { return ThreadFromID(id) }
// NOTE: Thread adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [Thread] class.
//
// # Initializing an NSThread Object
//
//   - [IThread.InitWithTargetSelectorObject]: Returns an [NSThread] object initialized with the given arguments.
//
// # Starting a Thread
//
//   - [IThread.Start]: Starts the receiver.
//   - [IThread.Main]: The main entry point routine for the thread.
//
// # Stopping a Thread
//
//   - [IThread.Cancel]: Changes the cancelled state of the receiver to indicate that it should exit.
//
// # Determining the Thread’s Execution State
//
//   - [IThread.Executing]: A Boolean value that indicates whether the receiver is executing.
//   - [IThread.Finished]: A Boolean value that indicates whether the receiver has finished execution.
//   - [IThread.Cancelled]: A Boolean value that indicates whether the receiver is cancelled.
//
// # Working with the Main Thread
//
//   - [IThread.IsMainThread]: A Boolean value that indicates whether the receiver is the main thread.
//
// # Working with Thread Properties
//
//   - [IThread.ThreadDictionary]: The thread object’s dictionary.
//   - [IThread.NSAssertionHandlerKey]: A key with a corresponding value in the thread dictionary.
//   - [IThread.Name]: The name of the receiver.
//   - [IThread.SetName]
//   - [IThread.StackSize]: The stack size of the receiver, in bytes.
//   - [IThread.SetStackSize]
//
// # Prioritizing Thread Work
//
//   - [IThread.QualityOfService]
//   - [IThread.SetQualityOfService]
//   - [IThread.ThreadPriority]: The receiver’s priority
//   - [IThread.SetThreadPriority]
//
// # Notifications
//
//   - [IThread.NSDidBecomeSingleThreaded]: Not implemented.
//   - [IThread.NSThreadWillExit]: An 
//   - [IThread.NSWillBecomeMultiThreaded]: Posted when the first thread is detached from the current thread. The 
//
// # Initializers
//
//   - [IThread.InitWithBlock]
//
// See: https://developer.apple.com/documentation/Foundation/Thread
type IThread interface {
	objectivec.IObject

	// Topic: Initializing an NSThread Object

	// Returns an [NSThread] object initialized with the given arguments.
	InitWithTargetSelectorObject(target objectivec.IObject, selector objc.SEL, argument objectivec.IObject) Thread

	// Topic: Starting a Thread

	// Starts the receiver.
	Start()
	// The main entry point routine for the thread.
	Main()

	// Topic: Stopping a Thread

	// Changes the cancelled state of the receiver to indicate that it should exit.
	Cancel()

	// Topic: Determining the Thread’s Execution State

	// A Boolean value that indicates whether the receiver is executing.
	Executing() bool
	// A Boolean value that indicates whether the receiver has finished execution.
	Finished() bool
	// A Boolean value that indicates whether the receiver is cancelled.
	Cancelled() bool

	// Topic: Working with the Main Thread

	// A Boolean value that indicates whether the receiver is the main thread.
	IsMainThread() bool

	// Topic: Working with Thread Properties

	// The thread object’s dictionary.
	ThreadDictionary() INSDictionary
	// A key with a corresponding value in the thread dictionary.
	NSAssertionHandlerKey() string
	// The name of the receiver.
	Name() string
	SetName(value string)
	// The stack size of the receiver, in bytes.
	StackSize() uint
	SetStackSize(value uint)

	// Topic: Prioritizing Thread Work

	QualityOfService() QualityOfService
	SetQualityOfService(value QualityOfService)
	// The receiver’s priority
	ThreadPriority() float64
	SetThreadPriority(value float64)

	// Topic: Notifications

	// Not implemented.
	NSDidBecomeSingleThreaded() NSNotificationName
	// An 
	NSThreadWillExit() NSNotificationName
	// Posted when the first thread is detached from the current thread. The 
	NSWillBecomeMultiThreaded() NSNotificationName

	// Topic: Initializers

	InitWithBlock(block VoidHandler) Thread
}





// Init initializes the instance.
func (t Thread) Init() Thread {
	rv := objc.Send[Thread](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t Thread) Autorelease() Thread {
	rv := objc.Send[Thread](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewThread creates a new Thread instance.
func NewThread() Thread {
	class := getThreadClass()
	rv := objc.Send[Thread](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Returns an [NSThread] object initialized with the given arguments.
//
// target: The object to which the message specified by `selector` is sent.
//
// selector: The selector for the message to send to `target`. This selector must take
// only one argument and must not have a return value.
//
// argument: The single argument passed to the target. May be `nil`.
//
// # Return Value
// 
// An [NSThread] object initialized with the given arguments.
//
// # Discussion
// 
// The objects `target` and `argument` are retained during the execution of
// the detached thread. They are released when the thread finally exits.
//
// See: https://developer.apple.com/documentation/Foundation/Thread/init(target:selector:object:)
func NewThreadWithTargetSelectorObject(target objectivec.IObject, selector objc.SEL, argument objectivec.IObject) Thread {
	instance := getThreadClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTarget:selector:object:"), target, selector, argument)
	return ThreadFromID(rv)
}







// Returns an [NSThread] object initialized with the given arguments.
//
// target: The object to which the message specified by `selector` is sent.
//
// selector: The selector for the message to send to `target`. This selector must take
// only one argument and must not have a return value.
//
// argument: The single argument passed to the target. May be `nil`.
//
// # Return Value
// 
// An [NSThread] object initialized with the given arguments.
//
// # Discussion
// 
// The objects `target` and `argument` are retained during the execution of
// the detached thread. They are released when the thread finally exits.
//
// See: https://developer.apple.com/documentation/Foundation/Thread/init(target:selector:object:)
func (t Thread) InitWithTargetSelectorObject(target objectivec.IObject, selector objc.SEL, argument objectivec.IObject) Thread {
	rv := objc.Send[Thread](t.ID, objc.Sel("initWithTarget:selector:object:"), target, selector, argument)
	return rv
}

// Starts the receiver.
//
// # Discussion
// 
// This method asynchronously spawns the new thread and invokes the
// receiver’s [Main] method on the new thread. The [Executing] property
// returns [true] once the thread starts executing, which may occur after the
// [Start] method returns.
// 
// If you initialized the receiver with a target and selector, the default
// [Main] method invokes that selector automatically.
// 
// If this thread is the first thread detached in the application, this method
// posts the [NSWillBecomeMultiThreaded] with object `nil` to the default
// notification center.
//
// [NSWillBecomeMultiThreaded]: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSWillBecomeMultiThreaded
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/Thread/start()
func (t Thread) Start() {
	objc.Send[objc.ID](t.ID, objc.Sel("start"))
}

// The main entry point routine for the thread.
//
// # Discussion
// 
// The default implementation of this method takes the target and selector
// used to initialize the receiver and invokes the selector on the specified
// target. If you subclass [NSThread], you can override this method and use it
// to implement the main body of your thread instead. If you do so, you do not
// need to invoke `super`.
// 
// You should never invoke this method directly. You should always start your
// thread by invoking the [Start] method.
//
// See: https://developer.apple.com/documentation/Foundation/Thread/main()
func (t Thread) Main() {
	objc.Send[objc.ID](t.ID, objc.Sel("main"))
}

// Changes the cancelled state of the receiver to indicate that it should
// exit.
//
// # Discussion
// 
// The semantics of this method are the same as those used for [NSOperation].
// This method sets state information in the receiver that is then reflected
// by the [Cancelled] property. Threads that support cancellation should
// periodically call the [Cancelled] method to determine if the thread has in
// fact been cancelled, and exit if it has been.
// 
// For more information about cancellation and operation objects, see
// [NSOperation].
//
// See: https://developer.apple.com/documentation/Foundation/Thread/cancel()
func (t Thread) Cancel() {
	objc.Send[objc.ID](t.ID, objc.Sel("cancel"))
}

//
// See: https://developer.apple.com/documentation/Foundation/Thread/init(block:)
func (t Thread) InitWithBlock(block VoidHandler) Thread {
		_block0, _cleanup0 := NewVoidBlock(block)
	defer _cleanup0()
		rv := objc.Send[objc.ID](t.ID, objc.Sel("initWithBlock:"), _block0)
	return NSThreadFromID(rv)
}





// Detaches a new thread and uses the specified selector as the thread entry
// point.
//
// selector: The selector for the message to send to the target. This selector must take
// only one argument and must not have a return value.
//
// target: The object that will receive the message `aSelector` on the new thread.
//
// argument: The single argument passed to the target. May be `nil`.
//
// # Discussion
// 
// The objects `aTarget` and `anArgument` are retained during the execution of
// the detached thread, then released. The detached thread is exited (using
// the [Exit] class method) as soon as `aTarget` has completed executing the
// `aSelector` method.
// 
// If this thread is the first thread detached in the application, this method
// posts the [NSWillBecomeMultiThreaded] with object `nil` to the default
// notification center.
//
// [NSWillBecomeMultiThreaded]: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSWillBecomeMultiThreaded
//
// See: https://developer.apple.com/documentation/Foundation/Thread/detachNewThreadSelector(_:toTarget:with:)
func (_ThreadClass ThreadClass) DetachNewThreadSelectorToTargetWithObject(selector objc.SEL, target objectivec.IObject, argument objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_ThreadClass.class), objc.Sel("detachNewThreadSelector:toTarget:withObject:"), selector, target, argument)
}

// Blocks the current thread until the time specified.
//
// date: The time at which to resume processing.
//
// # Discussion
// 
// No run loop processing occurs while the thread is blocked.
//
// See: https://developer.apple.com/documentation/Foundation/Thread/sleep(until:)
func (_ThreadClass ThreadClass) SleepUntilDate(date INSDate) {
	objc.Send[objc.ID](objc.ID(_ThreadClass.class), objc.Sel("sleepUntilDate:"), date)
}

// Sleeps the thread for a given time interval.
//
// ti: The duration of the sleep.
//
// # Discussion
// 
// No run loop processing occurs while the thread is blocked.
//
// See: https://developer.apple.com/documentation/Foundation/Thread/sleep(forTimeInterval:)
func (_ThreadClass ThreadClass) SleepForTimeInterval(ti float64) {
	objc.Send[objc.ID](objc.ID(_ThreadClass.class), objc.Sel("sleepForTimeInterval:"), ti)
}

// Terminates the current thread.
//
// # Discussion
// 
// This method uses the [CurrentThread] class method to access the current
// thread. Before exiting the thread, this method posts the [NSThreadWillExit]
// with the thread being exited to the default notification center. Because
// notifications are delivered synchronously, all observers of
// [NSThreadWillExit] are guaranteed to receive the notification before the
// thread exits.
// 
// Invoking this method should be avoided as it does not give your thread a
// chance to clean up any resources it allocated during its execution.
//
// [NSThreadWillExit]: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/NSThreadWillExit
//
// See: https://developer.apple.com/documentation/Foundation/Thread/exit()
func (_ThreadClass ThreadClass) Exit() {
	objc.Send[objc.ID](objc.ID(_ThreadClass.class), objc.Sel("exit"))
}

// Returns whether the application is multithreaded.
//
// # Return Value
// 
// [true] if the application is multithreaded, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// An application is considered multithreaded if a thread was ever detached
// from the main thread using either
// [DetachNewThreadSelectorToTargetWithObject] or [Start]. If you detached a
// thread in your application using a non-Cocoa API, such as the POSIX or
// Multiprocessing Services APIs, this method could still return [false]. The
// detached thread does not have to be currently running for the application
// to be considered multithreaded—this method only indicates whether a
// single thread has been spawned.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/Foundation/Thread/isMultiThreaded()
func (_ThreadClass ThreadClass) IsMultiThreaded() bool {
	rv := objc.Send[bool](objc.ID(_ThreadClass.class), objc.Sel("isMultiThreaded"))
	return rv
}

// Returns the current thread’s priority.
//
// # Return Value
// 
// The current thread’s priority, which is specified by a floating point
// number from 0.0 to 1.0, where 1.0 is highest priority.
//
// # Discussion
// 
// The priorities in this range are mapped to the operating system’s
// priority values. A “typical” thread priority might be 0.5, but because
// the priority is determined by the kernel, there is no guarantee what this
// value actually will be.
//
// See: https://developer.apple.com/documentation/Foundation/Thread/threadPriority()
func (_ThreadClass ThreadClass) ThreadPriority() float64 {
	rv := objc.Send[float64](objc.ID(_ThreadClass.class), objc.Sel("threadPriority"))
	return rv
}

// Sets the current thread’s priority.
//
// p: The new priority, specified with a floating point number from 0.0 to 1.0,
// where 1.0 is highest priority.
//
// # Return Value
// 
// [true] if the priority assignment succeeded, [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The priorities in this range are mapped to the operating system’s
// priority values.
//
// See: https://developer.apple.com/documentation/Foundation/Thread/setThreadPriority(_:)
func (_ThreadClass ThreadClass) SetThreadPriority(p float64) bool {
	rv := objc.Send[bool](objc.ID(_ThreadClass.class), objc.Sel("setThreadPriority:"), p)
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/Thread/detachNewThread(_:)
func (_ThreadClass ThreadClass) DetachNewThreadWithBlock(block VoidHandler) {
		_block0, _cleanup0 := NewVoidBlock(block)
	defer _cleanup0()
		objc.Send[objc.ID](objc.ID(_ThreadClass.class), objc.Sel("detachNewThreadWithBlock:"), _block0)
}








// A Boolean value that indicates whether the receiver is executing.
//
// # Discussion
// 
// [true] if the receiver is executing, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/Thread/isExecuting
func (t Thread) Executing() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isExecuting"))
	return rv
}



// A Boolean value that indicates whether the receiver has finished execution.
//
// # Discussion
// 
// [true] if the receiver has finished execution, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/Thread/isFinished
func (t Thread) Finished() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isFinished"))
	return rv
}



// A Boolean value that indicates whether the receiver is cancelled.
//
// # Discussion
// 
// [true] if the receiver has been cancelled, otherwise [false].
// 
// If your thread supports cancellation, it should check this property
// periodically and exit if it ever returns [true].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/Thread/isCancelled
func (t Thread) Cancelled() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isCancelled"))
	return rv
}



// A Boolean value that indicates whether the receiver is the main thread.
//
// # Discussion
// 
// [true] if the receiver is the main thread, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/Thread/isMainThread-swift.property
func (t Thread) IsMainThread() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isMainThread"))
	return rv
}



// The thread object’s dictionary.
//
// # Discussion
// 
// You can use the returned dictionary to store thread-specific data. The
// thread dictionary is not used during any manipulations of the [NSThread]
// object—it is simply a place where you can store any interesting data. For
// example, Foundation uses it to store the thread’s default [NSConnection]
// and [NSAssertionHandler] instances. You may define your own keys for the
// dictionary.
//
// See: https://developer.apple.com/documentation/Foundation/Thread/threadDictionary
func (t Thread) ThreadDictionary() INSDictionary {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("threadDictionary"))
	return NSDictionaryFromID(objc.ID(rv))
}



// A key with a corresponding value in the thread dictionary.
//
// See: https://developer.apple.com/documentation/foundation/nsassertionhandlerkey
func (t Thread) NSAssertionHandlerKey() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("NSAssertionHandlerKey"))
	return NSStringFromID(rv).String()
}



// The name of the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/Thread/name
func (t Thread) Name() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("name"))
	return NSStringFromID(rv).String()
}
func (t Thread) SetName(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setName:"), objc.String(value))
}



// The stack size of the receiver, in bytes.
//
// # Discussion
// 
// This value must be in bytes and a multiple of 4KB.
// 
// To change the stack size, you must set this property before starting your
// thread. Setting the stack size after the thread has started changes the
// attribute size (which is reflected by the [StackSize] method), but it does
// not affect the actual number of pages set aside for the thread.
//
// See: https://developer.apple.com/documentation/Foundation/Thread/stackSize
func (t Thread) StackSize() uint {
	rv := objc.Send[uint](t.ID, objc.Sel("stackSize"))
	return rv
}
func (t Thread) SetStackSize(value uint) {
	objc.Send[struct{}](t.ID, objc.Sel("setStackSize:"), value)
}



// See: https://developer.apple.com/documentation/Foundation/Thread/qualityOfService
func (t Thread) QualityOfService() QualityOfService {
	rv := objc.Send[QualityOfService](t.ID, objc.Sel("qualityOfService"))
	return QualityOfService(rv)
}
func (t Thread) SetQualityOfService(value QualityOfService) {
	objc.Send[struct{}](t.ID, objc.Sel("setQualityOfService:"), value)
}



// The receiver’s priority
//
// # Discussion
// 
// The thread’s priority, which is specified by a floating point number from
// 0.0 to 1.0, where 1.0 is highest priority.
// 
// The priorities in this range are mapped to the operating system’s
// priority values. A “typical” thread priority might be 0.5, but because
// the priority is determined by the kernel, there is no guarantee what this
// value actually will be.
//
// See: https://developer.apple.com/documentation/Foundation/Thread/threadPriority
func (t Thread) ThreadPriority() float64 {
	rv := objc.Send[float64](t.ID, objc.Sel("threadPriority"))
	return rv
}
func (t Thread) SetThreadPriority(value float64) {
	objc.Send[struct{}](t.ID, objc.Sel("setThreadPriority:"), value)
}



// Not implemented.
//
// See: https://developer.apple.com/documentation/foundation/nsnotification/name-swift.struct/nsdidbecomesinglethreaded
func (t Thread) NSDidBecomeSingleThreaded() NSNotificationName {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("NSDidBecomeSingleThreadedNotification"))
	return NSNotificationName(NSStringFromID(rv).String())
}



// An
//
// See: https://developer.apple.com/documentation/foundation/nsnotification/name-swift.struct/nsthreadwillexit
func (t Thread) NSThreadWillExit() NSNotificationName {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("NSThreadWillExitNotification"))
	return NSNotificationName(NSStringFromID(rv).String())
}



// Posted when the first thread is detached from the current thread. The
//
// See: https://developer.apple.com/documentation/foundation/nsnotification/name-swift.struct/nswillbecomemultithreaded
func (t Thread) NSWillBecomeMultiThreaded() NSNotificationName {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("NSWillBecomeMultiThreadedNotification"))
	return NSNotificationName(NSStringFromID(rv).String())
}







// Returns the [NSThread] object representing the main thread.
//
// # Return Value
// 
// The [NSThread] object representing the main thread.
//
// See: https://developer.apple.com/documentation/Foundation/Thread/main
func (_ThreadClass ThreadClass) MainThread() Thread {
	rv := objc.Send[objc.ID](objc.ID(_ThreadClass.class), objc.Sel("mainThread"))
	return NSThreadFromID(objc.ID(rv))
}



// Returns the thread object representing the current thread of execution.
//
// # Return Value
// 
// A thread object representing the current thread of execution.
//
// See: https://developer.apple.com/documentation/Foundation/Thread/current
func (_ThreadClass ThreadClass) CurrentThread() Thread {
	rv := objc.Send[objc.ID](objc.ID(_ThreadClass.class), objc.Sel("currentThread"))
	return NSThreadFromID(objc.ID(rv))
}



// Returns an array containing the call stack return addresses.
//
// # Return Value
// 
// An array containing the call stack return addresses. Each element is an
// [NSNumber] object containing an [NSUInteger] value.
//
// See: https://developer.apple.com/documentation/Foundation/Thread/callStackReturnAddresses
func (_ThreadClass ThreadClass) CallStackReturnAddresses() []NSNumber {
	rv := objc.Send[[]objc.ID](objc.ID(_ThreadClass.class), objc.Sel("callStackReturnAddresses"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSNumber {
		return NSNumberFromID(id)
	})
}



// Returns an array containing the call stack symbols.
//
// # Return Value
// 
// An array containing the call stack symbols. Each element is an [NSString]
// object with a value in a format determined by the `backtrace_symbols()`
// function. For more information, see backtrace_symbols(3) macOS Developer
// Tools Manual Page.
// 
// # Discussion
// 
// The return value describes the call stack backtrace of the current thread
// at the moment this method was called.
//
// See: https://developer.apple.com/documentation/Foundation/Thread/callStackSymbols
func (_ThreadClass ThreadClass) CallStackSymbols() []string {
	rv := objc.Send[[]objc.ID](objc.ID(_ThreadClass.class), objc.Sel("callStackSymbols"))
	return objc.ConvertSliceToStrings(rv)
}














// InitWithBlockSync is a synchronous wrapper around [Thread.InitWithBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (t Thread) InitWithBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	t.InitWithBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// DetachNewThreadWithBlockSync is a synchronous wrapper around [Thread.DetachNewThreadWithBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (tc ThreadClass) DetachNewThreadWithBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	tc.DetachNewThreadWithBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}






