// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSDistributedLock] class.
var (
	_NSDistributedLockClass     NSDistributedLockClass
	_NSDistributedLockClassOnce sync.Once
)

func getNSDistributedLockClass() NSDistributedLockClass {
	_NSDistributedLockClassOnce.Do(func() {
		_NSDistributedLockClass = NSDistributedLockClass{class: objc.GetClass("NSDistributedLock")}
	})
	return _NSDistributedLockClass
}

// GetNSDistributedLockClass returns the class object for NSDistributedLock.
func GetNSDistributedLockClass() NSDistributedLockClass {
	return getNSDistributedLockClass()
}

type NSDistributedLockClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSDistributedLockClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSDistributedLockClass) Alloc() NSDistributedLock {
	rv := objc.Send[NSDistributedLock](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A lock that multiple applications on multiple hosts can use to restrict
// access to some shared resource, such as a file.
//
// # Overview
//
// The lock is implemented by an entry (such as a file or directory) in the
// file system. For multiple applications to use an [NSDistributedLock] object
// to coordinate their activities, the lock must be writable on a file system
// accessible to all hosts on which the applications might be running.
//
// Use the [NSDistributedLock.TryLock] method to attempt to acquire a lock. You should generally
// use the [NSDistributedLock.Unlock] method to release the lock rather than [NSDistributedLock.BreakLock].
//
// [NSDistributedLock] doesn’t conform to the [NSLocking] protocol, nor does
// it have a `lock` method. The protocol’s [Lock] method is intended to
// block the execution of the thread until successful. For an
// [NSDistributedLock] object, this could mean polling the file system at some
// predetermined rate. A better solution is to provide the [NSDistributedLock.TryLock] method
// and let you determine the polling frequency that makes sense for your
// application.
//
// # Creating an NSDistributedLock
//
//   - [NSDistributedLock.InitWithPath]: Initializes an [NSDistributedLock] object to use as the lock the file-system entry specified by a given path.
//
// # Acquiring a Lock
//
//   - [NSDistributedLock.TryLock]: Attempts to acquire the receiver and immediately returns a Boolean value that indicates whether the attempt was successful.
//
// # Relinquishing a Lock
//
//   - [NSDistributedLock.BreakLock]: Forces the lock to be relinquished.
//   - [NSDistributedLock.Unlock]: Relinquishes the receiver.
//
// # Getting Lock Information
//
//   - [NSDistributedLock.LockDate]: Returns the time the receiver was acquired by any of the [NSDistributedLock] objects using the same path.
//
// See: https://developer.apple.com/documentation/Foundation/NSDistributedLock
type NSDistributedLock struct {
	objectivec.Object
}

// NSDistributedLockFromID constructs a [NSDistributedLock] from an objc.ID.
//
// A lock that multiple applications on multiple hosts can use to restrict
// access to some shared resource, such as a file.
func NSDistributedLockFromID(id objc.ID) NSDistributedLock {
	return NSDistributedLock{objectivec.Object{ID: id}}
}

// NOTE: NSDistributedLock adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSDistributedLock] class.
//
// # Creating an NSDistributedLock
//
//   - [INSDistributedLock.InitWithPath]: Initializes an [NSDistributedLock] object to use as the lock the file-system entry specified by a given path.
//
// # Acquiring a Lock
//
//   - [INSDistributedLock.TryLock]: Attempts to acquire the receiver and immediately returns a Boolean value that indicates whether the attempt was successful.
//
// # Relinquishing a Lock
//
//   - [INSDistributedLock.BreakLock]: Forces the lock to be relinquished.
//   - [INSDistributedLock.Unlock]: Relinquishes the receiver.
//
// # Getting Lock Information
//
//   - [INSDistributedLock.LockDate]: Returns the time the receiver was acquired by any of the [NSDistributedLock] objects using the same path.
//
// See: https://developer.apple.com/documentation/Foundation/NSDistributedLock
type INSDistributedLock interface {
	objectivec.IObject

	// Topic: Creating an NSDistributedLock

	// Initializes an [NSDistributedLock] object to use as the lock the file-system entry specified by a given path.
	InitWithPath(path string) NSDistributedLock

	// Topic: Acquiring a Lock

	// Attempts to acquire the receiver and immediately returns a Boolean value that indicates whether the attempt was successful.
	TryLock() bool

	// Topic: Relinquishing a Lock

	// Forces the lock to be relinquished.
	BreakLock()
	// Relinquishes the receiver.
	Unlock()

	// Topic: Getting Lock Information

	// Returns the time the receiver was acquired by any of the [NSDistributedLock] objects using the same path.
	LockDate() INSDate
}

// Init initializes the instance.
func (d NSDistributedLock) Init() NSDistributedLock {
	rv := objc.Send[NSDistributedLock](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d NSDistributedLock) Autorelease() NSDistributedLock {
	rv := objc.Send[NSDistributedLock](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSDistributedLock creates a new NSDistributedLock instance.
func NewNSDistributedLock() NSDistributedLock {
	class := getNSDistributedLockClass()
	rv := objc.Send[NSDistributedLock](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes an [NSDistributedLock] object to use as the lock the
// file-system entry specified by a given path.
//
// path: All of `path` up to the last component itself must exist. You can use
// [NSFileManager] to create (and set permissions) for any nonexistent
// intermediate directories.
//
// # Return Value
//
// An [NSDistributedLock] object initialized to use as the locking object the
// file-system entry specified by `path`.
//
// # Discussion
//
// For applications to use the lock, `path` must be accessible to—and
// writable by—all hosts on which the applications might be running.
//
// See: https://developer.apple.com/documentation/Foundation/NSDistributedLock/init(path:)
func NewDistributedLockWithPath(path string) NSDistributedLock {
	instance := getNSDistributedLockClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPath:"), objc.String(path))
	return NSDistributedLockFromID(rv)
}

// Initializes an [NSDistributedLock] object to use as the lock the
// file-system entry specified by a given path.
//
// path: All of `path` up to the last component itself must exist. You can use
// [NSFileManager] to create (and set permissions) for any nonexistent
// intermediate directories.
//
// # Return Value
//
// An [NSDistributedLock] object initialized to use as the locking object the
// file-system entry specified by `path`.
//
// # Discussion
//
// For applications to use the lock, `path` must be accessible to—and
// writable by—all hosts on which the applications might be running.
//
// See: https://developer.apple.com/documentation/Foundation/NSDistributedLock/init(path:)
func (d NSDistributedLock) InitWithPath(path string) NSDistributedLock {
	rv := objc.Send[NSDistributedLock](d.ID, objc.Sel("initWithPath:"), objc.String(path))
	return rv
}

// Attempts to acquire the receiver and immediately returns a Boolean value
// that indicates whether the attempt was successful.
//
// # Return Value
//
// true if the attempt to acquire the receiver was successful, otherwise
// false.
//
// # Discussion
//
// Raises [NSGenericException] if a file-system error occurs.
//
// See: https://developer.apple.com/documentation/Foundation/NSDistributedLock/try()
func (d NSDistributedLock) TryLock() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("tryLock"))
	return rv
}

// Forces the lock to be relinquished.
//
// # Discussion
//
// This method always succeeds unless the lock has been damaged. If another
// process has already unlocked or broken the lock, this method has no effect.
// You should generally use [Unlock] rather than [BreakLock] to relinquish a
// lock.
//
// Even if you break a lock, there’s no guarantee that you will then be able
// to acquire the lock—another process might get it before your [TryLock] is
// invoked.
//
// Raises an [NSGenericException] if the lock could not be removed.
//
// See: https://developer.apple.com/documentation/Foundation/NSDistributedLock/break()
func (d NSDistributedLock) BreakLock() {
	objc.Send[objc.ID](d.ID, objc.Sel("breakLock"))
}

// Relinquishes the receiver.
//
// # Discussion
//
// You should generally use the [Unlock] method rather than [BreakLock] to
// release a lock.
//
// An [NSGenericException] is raised if the receiver doesn’t already exist.
//
// See: https://developer.apple.com/documentation/Foundation/NSDistributedLock/unlock()
func (d NSDistributedLock) Unlock() {
	objc.Send[objc.ID](d.ID, objc.Sel("unlock"))
}

// Returns an [NSDistributedLock] object initialized to use as the locking
// object the file-system entry specified by a given path.
//
// path: All of `path` up to the last component itself must exist. You can use
// [NSFileManager] to create (and set permissions) for any nonexistent
// intermediate directories.
//
// # Return Value
//
// An [NSDistributedLock] object initialized to use as the locking object the
// file-system entry specified by `path`.
//
// # Discussion
//
// For applications to use the lock, `path` must be accessible to—and
// writable by—all hosts on which the applications might be running.
//
// See: https://developer.apple.com/documentation/Foundation/NSDistributedLock/lockWithPath:
func (_NSDistributedLockClass NSDistributedLockClass) LockWithPath(path string) NSDistributedLock {
	rv := objc.Send[objc.ID](objc.ID(_NSDistributedLockClass.class), objc.Sel("lockWithPath:"), objc.String(path))
	return NSDistributedLockFromID(rv)
}

// Returns the time the receiver was acquired by any of the
// [NSDistributedLock] objects using the same path.
//
// # Return Value
//
// The time the receiver was acquired by any of the [NSDistributedLock]
// objects using the same path. Returns `nil` if the lock doesn’t exist.
//
// # Discussion
//
// This method is potentially useful to applications that want to use an age
// heuristic to decide if a lock is too old and should be broken.
//
// If the creation date on the lock isn’t the date on which you locked it,
// you’ve lost the lock: it’s been broken since you last checked it.
//
// See: https://developer.apple.com/documentation/Foundation/NSDistributedLock/lockDate
func (d NSDistributedLock) LockDate() INSDate {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("lockDate"))
	return NSDateFromID(objc.ID(rv))
}
