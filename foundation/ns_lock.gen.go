// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSLock] class.
var (
	_NSLockClass     NSLockClass
	_NSLockClassOnce sync.Once
)

func getNSLockClass() NSLockClass {
	_NSLockClassOnce.Do(func() {
		_NSLockClass = NSLockClass{class: objc.GetClass("NSLock")}
	})
	return _NSLockClass
}

// GetNSLockClass returns the class object for NSLock.
func GetNSLockClass() NSLockClass {
	return getNSLockClass()
}

type NSLockClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSLockClass) Alloc() NSLock {
	rv := objc.Send[NSLock](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// An object that coordinates the operation of multiple threads of execution
// within the same application.
//
// # Overview
// 
// An [NSLock] object can be used to mediate access to an application’s
// global data or to protect a critical section of code, allowing it to run
// atomically.
// 
// You should not use this class to implement a recursive lock. Calling the
// `lock` method twice on the same thread will lock up your thread
// permanently. Use the [NSRecursiveLock] class to implement recursive locks
// instead.
// 
// Unlocking a lock that is not locked is considered a programmer error and
// should be fixed in your code. The [NSLock] class reports such errors by
// printing an error message to the console when they occur.
//
// # Acquiring a Lock
//
//   - [NSLock.LockBeforeDate]: Attempts to acquire a lock before a given time and returns a Boolean value indicating whether the attempt was successful.
//   - [NSLock.TryLock]: Attempts to acquire a lock and immediately returns a Boolean value that indicates whether the attempt was successful.
//
// # Naming the Lock
//
//   - [NSLock.Name]: The name associated with the receiver.
//   - [NSLock.SetName]
//
// See: https://developer.apple.com/documentation/Foundation/NSLock
type NSLock struct {
	objectivec.Object
}

// NSLockFromID constructs a [NSLock] from an objc.ID.
//
// An object that coordinates the operation of multiple threads of execution
// within the same application.
func NSLockFromID(id objc.ID) NSLock {
	return NSLock{objectivec.Object{id}}
}
// NOTE: NSLock adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSLock] class.
//
// # Acquiring a Lock
//
//   - [INSLock.LockBeforeDate]: Attempts to acquire a lock before a given time and returns a Boolean value indicating whether the attempt was successful.
//   - [INSLock.TryLock]: Attempts to acquire a lock and immediately returns a Boolean value that indicates whether the attempt was successful.
//
// # Naming the Lock
//
//   - [INSLock.Name]: The name associated with the receiver.
//   - [INSLock.SetName]
//
// See: https://developer.apple.com/documentation/Foundation/NSLock
type INSLock interface {
	objectivec.IObject

	// Topic: Acquiring a Lock

	// Attempts to acquire a lock before a given time and returns a Boolean value indicating whether the attempt was successful.
	LockBeforeDate(limit INSDate) bool
	// Attempts to acquire a lock and immediately returns a Boolean value that indicates whether the attempt was successful.
	TryLock() bool

	// Topic: Naming the Lock

	// The name associated with the receiver.
	Name() string
	SetName(value string)
}





// Init initializes the instance.
func (l NSLock) Init() NSLock {
	rv := objc.Send[NSLock](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l NSLock) Autorelease() NSLock {
	rv := objc.Send[NSLock](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSLock creates a new NSLock instance.
func NewNSLock() NSLock {
	class := getNSLockClass()
	rv := objc.Send[NSLock](objc.ID(class.class), objc.Sel("new"))
	return rv
}










// Attempts to acquire a lock before a given time and returns a Boolean value
// indicating whether the attempt was successful.
//
// limit: The time limit for attempting to acquire a lock.
//
// # Return Value
// 
// [true] if the lock is acquired before `limit`, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The thread is blocked until the receiver acquires the lock or `limit` is
// reached.
//
// See: https://developer.apple.com/documentation/Foundation/NSLock/lock(before:)
func (l NSLock) LockBeforeDate(limit INSDate) bool {
	rv := objc.Send[bool](l.ID, objc.Sel("lockBeforeDate:"), limit)
	return rv
}

// Attempts to acquire a lock and immediately returns a Boolean value that
// indicates whether the attempt was successful.
//
// # Return Value
// 
// [true] if the lock was acquired, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSLock/try()
func (l NSLock) TryLock() bool {
	rv := objc.Send[bool](l.ID, objc.Sel("tryLock"))
	return rv
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
func (l NSLock) Lock() {
	objc.Send[objc.ID](l.ID, objc.Sel("lock"))
}

// Relinquishes a previously acquired lock.
//
// See: https://developer.apple.com/documentation/Foundation/NSLocking/unlock()
func (l NSLock) Unlock() {
	objc.Send[objc.ID](l.ID, objc.Sel("unlock"))
}












// The name associated with the receiver.
//
// # Discussion
// 
// You can use a name string to identify a lock within your code. Cocoa also
// uses this name as part of any error descriptions involving the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSLock/name
func (l NSLock) Name() string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("name"))
	return NSStringFromID(rv).String()
}
func (l NSLock) SetName(value string) {
	objc.Send[struct{}](l.ID, objc.Sel("setName:"), objc.String(value))
}


























