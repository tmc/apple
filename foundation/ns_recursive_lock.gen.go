// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSRecursiveLock] class.
var (
	_NSRecursiveLockClass     NSRecursiveLockClass
	_NSRecursiveLockClassOnce sync.Once
)

func getNSRecursiveLockClass() NSRecursiveLockClass {
	_NSRecursiveLockClassOnce.Do(func() {
		_NSRecursiveLockClass = NSRecursiveLockClass{class: objc.GetClass("NSRecursiveLock")}
	})
	return _NSRecursiveLockClass
}

// GetNSRecursiveLockClass returns the class object for NSRecursiveLock.
func GetNSRecursiveLockClass() NSRecursiveLockClass {
	return getNSRecursiveLockClass()
}

type NSRecursiveLockClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSRecursiveLockClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSRecursiveLockClass) Alloc() NSRecursiveLock {
	rv := objc.Send[NSRecursiveLock](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A lock that may be acquired multiple times by the same thread without
// causing a deadlock.
//
// # Overview
// 
// [NSRecursiveLock] defines a lock that may be acquired multiple times by the
// same thread without causing a deadlock, a situation where a thread is
// permanently blocked waiting for itself to relinquish a lock. While the
// locking thread has one or more locks, all other threads are prevented from
// accessing the code protected by the lock.
//
// # Acquiring a Lock
//
//   - [NSRecursiveLock.LockBeforeDate]: Attempts to acquire a lock before a given date.
//   - [NSRecursiveLock.TryLock]: Attempts to acquire a lock, and immediately returns a Boolean value that indicates whether the attempt was successful.
//
// # Naming the Lock
//
//   - [NSRecursiveLock.Name]: The name associated with the receiver.
//   - [NSRecursiveLock.SetName]
//
// See: https://developer.apple.com/documentation/Foundation/NSRecursiveLock
type NSRecursiveLock struct {
	objectivec.Object
}

// NSRecursiveLockFromID constructs a [NSRecursiveLock] from an objc.ID.
//
// A lock that may be acquired multiple times by the same thread without
// causing a deadlock.
func NSRecursiveLockFromID(id objc.ID) NSRecursiveLock {
	return NSRecursiveLock{objectivec.Object{ID: id}}
}
// NOTE: NSRecursiveLock adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSRecursiveLock] class.
//
// # Acquiring a Lock
//
//   - [INSRecursiveLock.LockBeforeDate]: Attempts to acquire a lock before a given date.
//   - [INSRecursiveLock.TryLock]: Attempts to acquire a lock, and immediately returns a Boolean value that indicates whether the attempt was successful.
//
// # Naming the Lock
//
//   - [INSRecursiveLock.Name]: The name associated with the receiver.
//   - [INSRecursiveLock.SetName]
//
// See: https://developer.apple.com/documentation/Foundation/NSRecursiveLock
type INSRecursiveLock interface {
	objectivec.IObject

	// Topic: Acquiring a Lock

	// Attempts to acquire a lock before a given date.
	LockBeforeDate(limit INSDate) bool
	// Attempts to acquire a lock, and immediately returns a Boolean value that indicates whether the attempt was successful.
	TryLock() bool

	// Topic: Naming the Lock

	// The name associated with the receiver.
	Name() string
	SetName(value string)
}

// Init initializes the instance.
func (r NSRecursiveLock) Init() NSRecursiveLock {
	rv := objc.Send[NSRecursiveLock](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r NSRecursiveLock) Autorelease() NSRecursiveLock {
	rv := objc.Send[NSRecursiveLock](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSRecursiveLock creates a new NSRecursiveLock instance.
func NewNSRecursiveLock() NSRecursiveLock {
	class := getNSRecursiveLockClass()
	rv := objc.Send[NSRecursiveLock](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Attempts to acquire a lock before a given date.
//
// limit: The time before which the lock should be acquired.
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
// See: https://developer.apple.com/documentation/Foundation/NSRecursiveLock/lock(before:)
func (r NSRecursiveLock) LockBeforeDate(limit INSDate) bool {
	rv := objc.Send[bool](r.ID, objc.Sel("lockBeforeDate:"), limit)
	return rv
}
// Attempts to acquire a lock, and immediately returns a Boolean value that
// indicates whether the attempt was successful.
//
// # Return Value
// 
// [true] if successful, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Foundation/NSRecursiveLock/try()
func (r NSRecursiveLock) TryLock() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("tryLock"))
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
func (r NSRecursiveLock) Lock() {
	objc.Send[objc.ID](r.ID, objc.Sel("lock"))
}
// Relinquishes a previously acquired lock.
//
// See: https://developer.apple.com/documentation/Foundation/NSLocking/unlock()
func (r NSRecursiveLock) Unlock() {
	objc.Send[objc.ID](r.ID, objc.Sel("unlock"))
}

// The name associated with the receiver.
//
// # Discussion
// 
// You can use a name string to identify a lock within your code. Cocoa also
// uses this name as part of any error descriptions involving the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSRecursiveLock/name
func (r NSRecursiveLock) Name() string {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("name"))
	return NSStringFromID(rv).String()
}
func (r NSRecursiveLock) SetName(value string) {
	objc.Send[struct{}](r.ID, objc.Sel("setName:"), objc.String(value))
}

