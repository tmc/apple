// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSPersistentStoreResult] class.
var (
	_NSPersistentStoreResultClass     NSPersistentStoreResultClass
	_NSPersistentStoreResultClassOnce sync.Once
)

func getNSPersistentStoreResultClass() NSPersistentStoreResultClass {
	_NSPersistentStoreResultClassOnce.Do(func() {
		_NSPersistentStoreResultClass = NSPersistentStoreResultClass{class: objc.GetClass("NSPersistentStoreResult")}
	})
	return _NSPersistentStoreResultClass
}

// GetNSPersistentStoreResultClass returns the class object for NSPersistentStoreResult.
func GetNSPersistentStoreResultClass() NSPersistentStoreResultClass {
	return getNSPersistentStoreResultClass()
}

type NSPersistentStoreResultClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSPersistentStoreResultClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPersistentStoreResultClass) Alloc() NSPersistentStoreResult {
	rv := objc.Send[NSPersistentStoreResult](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The abstract base class for results returned from a persistent store
// coordinator.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreResult
type NSPersistentStoreResult struct {
	objectivec.Object
}

// NSPersistentStoreResultFromID constructs a [NSPersistentStoreResult] from an objc.ID.
//
// The abstract base class for results returned from a persistent store
// coordinator.
func NSPersistentStoreResultFromID(id objc.ID) NSPersistentStoreResult {
	return NSPersistentStoreResult{objectivec.Object{ID: id}}
}

// NOTE: NSPersistentStoreResult adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPersistentStoreResult] class.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentStoreResult
type INSPersistentStoreResult interface {
	objectivec.IObject
}

// Init initializes the instance.
func (p NSPersistentStoreResult) Init() NSPersistentStoreResult {
	rv := objc.Send[NSPersistentStoreResult](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPersistentStoreResult) Autorelease() NSPersistentStoreResult {
	rv := objc.Send[NSPersistentStoreResult](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPersistentStoreResult creates a new NSPersistentStoreResult instance.
func NewNSPersistentStoreResult() NSPersistentStoreResult {
	class := getNSPersistentStoreResultClass()
	rv := objc.Send[NSPersistentStoreResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}
