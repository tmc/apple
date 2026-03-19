// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The elementary methods adopted by classes that define lock objects.
//
// See: https://developer.apple.com/documentation/Foundation/NSLocking
type NSLocking interface {
	objectivec.IObject

	// Attempts to acquire a lock, blocking a thread’s execution until the lock can be acquired.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSLocking/lock()
	Lock()

	// Relinquishes a previously acquired lock.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSLocking/unlock()
	Unlock()
}

// NSLockingObject wraps an existing Objective-C object that conforms to the NSLocking protocol.
type NSLockingObject struct {
	objectivec.Object
}
func (o NSLockingObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSLockingObjectFromID constructs a [NSLockingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSLockingObjectFromID(id objc.ID) NSLockingObject {
	return NSLockingObject{
		Object: objectivec.ObjectFromID(id),
	}
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
func (o NSLockingObject) Lock() {
	
	objc.Send[struct{}](o.ID, objc.Sel("lock"))
	}
// Relinquishes a previously acquired lock.
//
// See: https://developer.apple.com/documentation/Foundation/NSLocking/unlock()
func (o NSLockingObject) Unlock() {
	
	objc.Send[struct{}](o.ID, objc.Sel("unlock"))
	}

