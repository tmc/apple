// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSPersistentStoreAsynchronousResult] class.
var (
	_NSPersistentStoreAsynchronousResultClass     NSPersistentStoreAsynchronousResultClass
	_NSPersistentStoreAsynchronousResultClassOnce sync.Once
)

func getNSPersistentStoreAsynchronousResultClass() NSPersistentStoreAsynchronousResultClass {
	_NSPersistentStoreAsynchronousResultClassOnce.Do(func() {
		_NSPersistentStoreAsynchronousResultClass = NSPersistentStoreAsynchronousResultClass{class: objc.GetClass("NSPersistentStoreAsynchronousResult")}
	})
	return _NSPersistentStoreAsynchronousResultClass
}

// GetNSPersistentStoreAsynchronousResultClass returns the class object for NSPersistentStoreAsynchronousResult.
func GetNSPersistentStoreAsynchronousResultClass() NSPersistentStoreAsynchronousResultClass {
	return getNSPersistentStoreAsynchronousResultClass()
}

type NSPersistentStoreAsynchronousResultClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSPersistentStoreAsynchronousResultClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPersistentStoreAsynchronousResultClass) Alloc() NSPersistentStoreAsynchronousResult {
	rv := objc.Send[NSPersistentStoreAsynchronousResult](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A concrete class used to represent the results of an asynchronous request.
//
// # Inspecting the Result
//
//   - [NSPersistentStoreAsynchronousResult.ManagedObjectContext]: The managed object context for the result.
//   - [NSPersistentStoreAsynchronousResult.OperationError]: An error that contains details if the asynchronous fetch request fails.
//   - [NSPersistentStoreAsynchronousResult.Progress]: An object that reports progress for the asynchronous fetch request.
//
// # Canceling the Result
//
//   - [NSPersistentStoreAsynchronousResult.Cancel]: Cancels the asynchronous fetch request.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreAsynchronousResult
type NSPersistentStoreAsynchronousResult struct {
	NSPersistentStoreResult
}

// NSPersistentStoreAsynchronousResultFromID constructs a [NSPersistentStoreAsynchronousResult] from an objc.ID.
//
// A concrete class used to represent the results of an asynchronous request.
func NSPersistentStoreAsynchronousResultFromID(id objc.ID) NSPersistentStoreAsynchronousResult {
	return NSPersistentStoreAsynchronousResult{NSPersistentStoreResult: NSPersistentStoreResultFromID(id)}
}

// NOTE: NSPersistentStoreAsynchronousResult adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPersistentStoreAsynchronousResult] class.
//
// # Inspecting the Result
//
//   - [INSPersistentStoreAsynchronousResult.ManagedObjectContext]: The managed object context for the result.
//   - [INSPersistentStoreAsynchronousResult.OperationError]: An error that contains details if the asynchronous fetch request fails.
//   - [INSPersistentStoreAsynchronousResult.Progress]: An object that reports progress for the asynchronous fetch request.
//
// # Canceling the Result
//
//   - [INSPersistentStoreAsynchronousResult.Cancel]: Cancels the asynchronous fetch request.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreAsynchronousResult
type INSPersistentStoreAsynchronousResult interface {
	INSPersistentStoreResult

	// Topic: Inspecting the Result

	// The managed object context for the result.
	ManagedObjectContext() INSManagedObjectContext
	// An error that contains details if the asynchronous fetch request fails.
	OperationError() foundation.NSError
	// An object that reports progress for the asynchronous fetch request.
	Progress() foundation.Progress

	// Topic: Canceling the Result

	// Cancels the asynchronous fetch request.
	Cancel()
}

// Init initializes the instance.
func (p NSPersistentStoreAsynchronousResult) Init() NSPersistentStoreAsynchronousResult {
	rv := objc.Send[NSPersistentStoreAsynchronousResult](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPersistentStoreAsynchronousResult) Autorelease() NSPersistentStoreAsynchronousResult {
	rv := objc.Send[NSPersistentStoreAsynchronousResult](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPersistentStoreAsynchronousResult creates a new NSPersistentStoreAsynchronousResult instance.
func NewNSPersistentStoreAsynchronousResult() NSPersistentStoreAsynchronousResult {
	class := getNSPersistentStoreAsynchronousResultClass()
	rv := objc.Send[NSPersistentStoreAsynchronousResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Cancels the asynchronous fetch request.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreAsynchronousResult/cancel()
func (p NSPersistentStoreAsynchronousResult) Cancel() {
	objc.Send[objc.ID](p.ID, objc.Sel("cancel"))
}

// The managed object context for the result.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreAsynchronousResult/managedObjectContext
func (p NSPersistentStoreAsynchronousResult) ManagedObjectContext() INSManagedObjectContext {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("managedObjectContext"))
	return NSManagedObjectContextFromID(objc.ID(rv))
}

// An error that contains details if the asynchronous fetch request fails.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreAsynchronousResult/operationError
func (p NSPersistentStoreAsynchronousResult) OperationError() foundation.NSError {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("operationError"))
	return foundation.NSErrorFromID(objc.ID(rv))
}

// An object that reports progress for the asynchronous fetch request.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreAsynchronousResult/progress
func (p NSPersistentStoreAsynchronousResult) Progress() foundation.Progress {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("progress"))
	return foundation.ProgressFromID(objc.ID(rv))
}
