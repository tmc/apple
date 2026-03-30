// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSCondition] class.
var (
	_NSConditionClass     NSConditionClass
	_NSConditionClassOnce sync.Once
)

func getNSConditionClass() NSConditionClass {
	_NSConditionClassOnce.Do(func() {
		_NSConditionClass = NSConditionClass{class: objc.GetClass("NSCondition")}
	})
	return _NSConditionClass
}

// GetNSConditionClass returns the class object for NSCondition.
func GetNSConditionClass() NSConditionClass {
	return getNSConditionClass()
}

type NSConditionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSConditionClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSConditionClass) Alloc() NSCondition {
	rv := objc.Send[NSCondition](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A condition variable whose semantics follow those used for POSIX-style
// conditions.
//
// # Overview
//
// A condition object acts as both a lock and a checkpoint in a given thread.
// The lock protects your code while it tests the condition and performs the
// task triggered by the condition. The checkpoint behavior requires that the
// condition be true before the thread proceeds with its task. While the
// condition is not true, the thread blocks. It remains blocked until another
// thread signals the condition object.
//
// The semantics for using an [NSCondition] object are as follows:
//
// - Lock the condition object. - Test a boolean predicate. (This predicate is
// a boolean flag or other variable in your code that indicates whether it is
// safe to perform the task protected by the condition.) - If the boolean
// predicate is false, call the condition object’s [NSCondition.Wait] or [NSCondition.WaitUntilDate]
// method to block the thread. Upon returning from these methods, go to step 2
// to retest your boolean predicate. (Continue waiting and retesting the
// predicate until it is true.) - If the boolean predicate is true, perform
// the task. - Optionally update any predicates (or signal any conditions)
// affected by your task. - When your task is done, unlock the condition
// object.
//
// The pseudocode for performing the preceding steps would therefore look
// something like the following:
//
// Whenever you use a condition object, the first step is to lock the
// condition. Locking the condition ensures that your predicate and task code
// are protected from interference by other threads using the same condition.
// Once you have completed your task, you can set other predicates or signal
// other conditions based on the needs of your code. You should always set
// predicates and signal conditions while holding the condition object’s
// lock.
//
// When a thread waits on a condition, the condition object unlocks its lock
// and blocks the thread. When the condition is signaled, the system wakes up
// the thread. The condition object then reacquires its lock before returning
// from the [NSCondition.Wait] or [NSCondition.WaitUntilDate] method. Thus, from the point of view of
// the thread, it is as if it always held the lock.
//
// A boolean predicate is an important part of the semantics of using
// conditions because of the way signaling works. Signaling a condition does
// not guarantee that the condition itself is true. There are timing issues
// involved in signaling that may cause false signals to appear. Using a
// predicate ensures that these spurious signals do not cause you to perform
// work before it is safe to do so. The predicate itself is simply a flag or
// other variable in your code that you test in order to acquire a Boolean
// result.
//
// For more information on how to use conditions, see Using POSIX Thread Locks
// in [Threading Programming Guide].
//
// # Waiting for the Lock
//
//   - [NSCondition.Wait]: Blocks the current thread until the condition is signaled.
//   - [NSCondition.WaitUntilDate]: Blocks the current thread until the condition is signaled or the specified time limit is reached.
//
// # Signaling Waiting Threads
//
//   - [NSCondition.Signal]: Signals the condition, waking up one thread waiting on it.
//   - [NSCondition.Broadcast]: Signals the condition, waking up all threads waiting on it.
//
// # Identifying the Condition
//
//   - [NSCondition.Name]: The name of the condition.
//   - [NSCondition.SetName]
//
// See: https://developer.apple.com/documentation/Foundation/NSCondition
//
// [Threading Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Multithreading/Introduction/Introduction.html#//apple_ref/doc/uid/10000057i
type NSCondition struct {
	objectivec.Object
}

// NSConditionFromID constructs a [NSCondition] from an objc.ID.
//
// A condition variable whose semantics follow those used for POSIX-style
// conditions.
func NSConditionFromID(id objc.ID) NSCondition {
	return NSCondition{objectivec.Object{ID: id}}
}

// NOTE: NSCondition adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSCondition] class.
//
// # Waiting for the Lock
//
//   - [INSCondition.Wait]: Blocks the current thread until the condition is signaled.
//   - [INSCondition.WaitUntilDate]: Blocks the current thread until the condition is signaled or the specified time limit is reached.
//
// # Signaling Waiting Threads
//
//   - [INSCondition.Signal]: Signals the condition, waking up one thread waiting on it.
//   - [INSCondition.Broadcast]: Signals the condition, waking up all threads waiting on it.
//
// # Identifying the Condition
//
//   - [INSCondition.Name]: The name of the condition.
//   - [INSCondition.SetName]
//
// See: https://developer.apple.com/documentation/Foundation/NSCondition
type INSCondition interface {
	objectivec.IObject

	// Topic: Waiting for the Lock

	// Blocks the current thread until the condition is signaled.
	Wait()
	// Blocks the current thread until the condition is signaled or the specified time limit is reached.
	WaitUntilDate(limit INSDate) bool

	// Topic: Signaling Waiting Threads

	// Signals the condition, waking up one thread waiting on it.
	Signal()
	// Signals the condition, waking up all threads waiting on it.
	Broadcast()

	// Topic: Identifying the Condition

	// The name of the condition.
	Name() string
	SetName(value string)
}

// Init initializes the instance.
func (c NSCondition) Init() NSCondition {
	rv := objc.Send[NSCondition](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSCondition) Autorelease() NSCondition {
	rv := objc.Send[NSCondition](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSCondition creates a new NSCondition instance.
func NewNSCondition() NSCondition {
	class := getNSConditionClass()
	rv := objc.Send[NSCondition](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Blocks the current thread until the condition is signaled.
//
// # Discussion
//
// You must lock the receiver prior to calling this method.
//
// See: https://developer.apple.com/documentation/Foundation/NSCondition/wait()
func (c NSCondition) Wait() {
	objc.Send[objc.ID](c.ID, objc.Sel("wait"))
}

// Blocks the current thread until the condition is signaled or the specified
// time limit is reached.
//
// limit: The time at which to wake up the thread if the condition has not been
// signaled.
//
// # Return Value
//
// true if the condition was signaled; otherwise, false if the time limit was
// reached.
//
// # Discussion
//
// You must lock the receiver prior to calling this method.
//
// See: https://developer.apple.com/documentation/Foundation/NSCondition/wait(until:)
func (c NSCondition) WaitUntilDate(limit INSDate) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("waitUntilDate:"), limit)
	return rv
}

// Signals the condition, waking up one thread waiting on it.
//
// # Discussion
//
// You use this method to wake up one thread that is waiting on the condition.
// You may call this method multiple times to wake up multiple threads. If no
// threads are waiting on the condition, this method does nothing.
//
// To avoid race conditions, you should invoke this method only while the
// receiver is locked.
//
// See: https://developer.apple.com/documentation/Foundation/NSCondition/signal()
func (c NSCondition) Signal() {
	objc.Send[objc.ID](c.ID, objc.Sel("signal"))
}

// Signals the condition, waking up all threads waiting on it.
//
// # Discussion
//
// If no threads are waiting on the condition, this method does nothing.
//
// To avoid race conditions, you should invoke this method only while the
// receiver is locked.
//
// See: https://developer.apple.com/documentation/Foundation/NSCondition/broadcast()
func (c NSCondition) Broadcast() {
	objc.Send[objc.ID](c.ID, objc.Sel("broadcast"))
}

// Attempts to acquire a lock, blocking a thread’s execution until the lock
// can be acquired.
//
// # Discussion
//
// An application protects a critical section of code by requiring a thread to
// acquire a lock before executing the code. Once the critical section is
// completed, the thread relinquishes the lock by invoking [Unlock].
//
// See: https://developer.apple.com/documentation/Foundation/NSLocking/lock()
func (c NSCondition) Lock() {
	objc.Send[objc.ID](c.ID, objc.Sel("lock"))
}

// Relinquishes a previously acquired lock.
//
// See: https://developer.apple.com/documentation/Foundation/NSLocking/unlock()
func (c NSCondition) Unlock() {
	objc.Send[objc.ID](c.ID, objc.Sel("unlock"))
}

// The name of the condition.
//
// # Discussion
//
// You can use a name string to identify a condition object within your code.
// Cocoa also uses this name as part of any error descriptions involving the
// condition.
//
// See: https://developer.apple.com/documentation/Foundation/NSCondition/name
func (c NSCondition) Name() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("name"))
	return NSStringFromID(rv).String()
}
func (c NSCondition) SetName(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setName:"), objc.String(value))
}
