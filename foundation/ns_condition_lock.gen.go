// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSConditionLock] class.
var (
	_NSConditionLockClass     NSConditionLockClass
	_NSConditionLockClassOnce sync.Once
)

func getNSConditionLockClass() NSConditionLockClass {
	_NSConditionLockClassOnce.Do(func() {
		_NSConditionLockClass = NSConditionLockClass{class: objc.GetClass("NSConditionLock")}
	})
	return _NSConditionLockClass
}

// GetNSConditionLockClass returns the class object for NSConditionLock.
func GetNSConditionLockClass() NSConditionLockClass {
	return getNSConditionLockClass()
}

type NSConditionLockClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSConditionLockClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSConditionLockClass) Alloc() NSConditionLock {
	rv := objc.Send[NSConditionLock](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A lock that can be associated with specific, user-defined conditions.
//
// # Overview
// 
// Using an [NSConditionLock] object, you can ensure that a thread can acquire
// a lock only if a certain condition is met. Once it has acquired the lock
// and executed the critical section of code, the thread can relinquish the
// lock and set the associated condition to something new. The conditions
// themselves are arbitrary: you define them as needed for your application.
//
// # Initializing an NSConditionLock Object
//
//   - [NSConditionLock.InitWithCondition]: Initializes a newly allocated [NSConditionLock] object and sets its condition.
//
// # Accessing the Condition
//
//   - [NSConditionLock.Condition]: The condition associated with the receiver.
//
// # Acquiring and Releasing a Lock
//
//   - [NSConditionLock.LockBeforeDate]: Attempts to acquire a lock before a specified moment in time.
//   - [NSConditionLock.LockWhenCondition]: Attempts to acquire a lock.
//   - [NSConditionLock.LockWhenConditionBeforeDate]: Attempts to acquire a lock before a specified moment in time.
//   - [NSConditionLock.TryLock]: Attempts to acquire a lock without regard to the receiver’s condition.
//   - [NSConditionLock.TryLockWhenCondition]: Attempts to acquire a lock if the receiver’s condition is equal to the specified condition.
//   - [NSConditionLock.UnlockWithCondition]: Relinquishes the lock and sets the receiver’s condition.
//
// # Identifying the Condition Lock
//
//   - [NSConditionLock.Name]: The name associated with the receiver.
//   - [NSConditionLock.SetName]
//
// See: https://developer.apple.com/documentation/Foundation/NSConditionLock
type NSConditionLock struct {
	objectivec.Object
}

// NSConditionLockFromID constructs a [NSConditionLock] from an objc.ID.
//
// A lock that can be associated with specific, user-defined conditions.
func NSConditionLockFromID(id objc.ID) NSConditionLock {
	return NSConditionLock{objectivec.Object{ID: id}}
}
// NOTE: NSConditionLock adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSConditionLock] class.
//
// # Initializing an NSConditionLock Object
//
//   - [INSConditionLock.InitWithCondition]: Initializes a newly allocated [NSConditionLock] object and sets its condition.
//
// # Accessing the Condition
//
//   - [INSConditionLock.Condition]: The condition associated with the receiver.
//
// # Acquiring and Releasing a Lock
//
//   - [INSConditionLock.LockBeforeDate]: Attempts to acquire a lock before a specified moment in time.
//   - [INSConditionLock.LockWhenCondition]: Attempts to acquire a lock.
//   - [INSConditionLock.LockWhenConditionBeforeDate]: Attempts to acquire a lock before a specified moment in time.
//   - [INSConditionLock.TryLock]: Attempts to acquire a lock without regard to the receiver’s condition.
//   - [INSConditionLock.TryLockWhenCondition]: Attempts to acquire a lock if the receiver’s condition is equal to the specified condition.
//   - [INSConditionLock.UnlockWithCondition]: Relinquishes the lock and sets the receiver’s condition.
//
// # Identifying the Condition Lock
//
//   - [INSConditionLock.Name]: The name associated with the receiver.
//   - [INSConditionLock.SetName]
//
// See: https://developer.apple.com/documentation/Foundation/NSConditionLock
type INSConditionLock interface {
	objectivec.IObject

	// Topic: Initializing an NSConditionLock Object

	// Initializes a newly allocated [NSConditionLock] object and sets its condition.
	InitWithCondition(condition int) NSConditionLock

	// Topic: Accessing the Condition

	// The condition associated with the receiver.
	Condition() int

	// Topic: Acquiring and Releasing a Lock

	// Attempts to acquire a lock before a specified moment in time.
	LockBeforeDate(limit INSDate) bool
	// Attempts to acquire a lock.
	LockWhenCondition(condition int)
	// Attempts to acquire a lock before a specified moment in time.
	LockWhenConditionBeforeDate(condition int, limit INSDate) bool
	// Attempts to acquire a lock without regard to the receiver’s condition.
	TryLock() bool
	// Attempts to acquire a lock if the receiver’s condition is equal to the specified condition.
	TryLockWhenCondition(condition int) bool
	// Relinquishes the lock and sets the receiver’s condition.
	UnlockWithCondition(condition int)

	// Topic: Identifying the Condition Lock

	// The name associated with the receiver.
	Name() string
	SetName(value string)
}

// Init initializes the instance.
func (c NSConditionLock) Init() NSConditionLock {
	rv := objc.Send[NSConditionLock](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSConditionLock) Autorelease() NSConditionLock {
	rv := objc.Send[NSConditionLock](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSConditionLock creates a new NSConditionLock instance.
func NewNSConditionLock() NSConditionLock {
	class := getNSConditionLockClass()
	rv := objc.Send[NSConditionLock](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a newly allocated [NSConditionLock] object and sets its
// condition.
//
// condition: The user-defined condition for the lock. The value of `condition` is
// user-defined; see the class description for more information.
//
// # Return Value
// 
// An initialized condition lock object; may be different than the original
// receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSConditionLock/init(condition:)
func NewConditionLockWithCondition(condition int) NSConditionLock {
	instance := getNSConditionLockClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCondition:"), condition)
	return NSConditionLockFromID(rv)
}

// Initializes a newly allocated [NSConditionLock] object and sets its
// condition.
//
// condition: The user-defined condition for the lock. The value of `condition` is
// user-defined; see the class description for more information.
//
// # Return Value
// 
// An initialized condition lock object; may be different than the original
// receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSConditionLock/init(condition:)
func (c NSConditionLock) InitWithCondition(condition int) NSConditionLock {
	rv := objc.Send[NSConditionLock](c.ID, objc.Sel("initWithCondition:"), condition)
	return rv
}
// Attempts to acquire a lock before a specified moment in time.
//
// limit: The date by which the lock must be acquired or the attempt will time out.
//
// # Return Value
// 
// [true] if the lock is acquired within the time limit, [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The condition associated with the receiver isn’t taken into account in
// this operation. This method blocks the thread’s execution until the
// receiver acquires the lock or `limit` is reached.
//
// See: https://developer.apple.com/documentation/Foundation/NSConditionLock/lock(before:)
func (c NSConditionLock) LockBeforeDate(limit INSDate) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("lockBeforeDate:"), limit)
	return rv
}
// Attempts to acquire a lock.
//
// condition: The condition to match on.
//
// # Discussion
// 
// The receiver’s condition must be equal to `condition` before the locking
// operation will succeed. This method blocks the thread’s execution until
// the lock can be acquired.
//
// See: https://developer.apple.com/documentation/Foundation/NSConditionLock/lock(whenCondition:)
func (c NSConditionLock) LockWhenCondition(condition int) {
	objc.Send[objc.ID](c.ID, objc.Sel("lockWhenCondition:"), condition)
}
// Attempts to acquire a lock before a specified moment in time.
//
// condition: The condition to match on.
//
// limit: The date by which the lock must be acquired or the attempt will time out.
//
// # Return Value
// 
// [true] if the lock is acquired within the time limit, [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The receiver’s condition must be equal to `condition` before the locking
// operation will succeed. This method blocks the thread’s execution until
// the lock can be acquired or `limit` is reached.
//
// See: https://developer.apple.com/documentation/Foundation/NSConditionLock/lock(whenCondition:before:)
func (c NSConditionLock) LockWhenConditionBeforeDate(condition int, limit INSDate) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("lockWhenCondition:beforeDate:"), condition, limit)
	return rv
}
// Attempts to acquire a lock without regard to the receiver’s condition.
//
// # Return Value
// 
// [true] if the lock could be acquired, [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method returns immediately.
//
// See: https://developer.apple.com/documentation/Foundation/NSConditionLock/try()
func (c NSConditionLock) TryLock() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("tryLock"))
	return rv
}
// Attempts to acquire a lock if the receiver’s condition is equal to the
// specified condition.
//
// # Return Value
// 
// [true] if the lock could be acquired, [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// As part of its implementation, this method invokes
// [LockWhenConditionBeforeDate]. This method returns immediately.
//
// See: https://developer.apple.com/documentation/Foundation/NSConditionLock/tryLock(whenCondition:)
func (c NSConditionLock) TryLockWhenCondition(condition int) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("tryLockWhenCondition:"), condition)
	return rv
}
// Relinquishes the lock and sets the receiver’s condition.
//
// condition: The user-defined condition for the lock. The value of `condition` is
// user-defined; see the class description for more information.
//
// See: https://developer.apple.com/documentation/Foundation/NSConditionLock/unlock(withCondition:)
func (c NSConditionLock) UnlockWithCondition(condition int) {
	objc.Send[objc.ID](c.ID, objc.Sel("unlockWithCondition:"), condition)
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
func (c NSConditionLock) Lock() {
	objc.Send[objc.ID](c.ID, objc.Sel("lock"))
}
// Relinquishes a previously acquired lock.
//
// See: https://developer.apple.com/documentation/Foundation/NSLocking/unlock()
func (c NSConditionLock) Unlock() {
	objc.Send[objc.ID](c.ID, objc.Sel("unlock"))
}

// The condition associated with the receiver.
//
// # Discussion
// 
// If no condition has been set, returns 0.
//
// See: https://developer.apple.com/documentation/Foundation/NSConditionLock/condition
func (c NSConditionLock) Condition() int {
	rv := objc.Send[int](c.ID, objc.Sel("condition"))
	return rv
}
// The name associated with the receiver.
//
// # Discussion
// 
// You can use a name string to identify a condition lock within your code.
// Cocoa also uses this name as part of any error descriptions involving the
// receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSConditionLock/name
func (c NSConditionLock) Name() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("name"))
	return NSStringFromID(rv).String()
}
func (c NSConditionLock) SetName(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setName:"), objc.String(value))
}

