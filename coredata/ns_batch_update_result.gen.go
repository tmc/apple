// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSBatchUpdateResult] class.
var (
	_NSBatchUpdateResultClass     NSBatchUpdateResultClass
	_NSBatchUpdateResultClassOnce sync.Once
)

func getNSBatchUpdateResultClass() NSBatchUpdateResultClass {
	_NSBatchUpdateResultClassOnce.Do(func() {
		_NSBatchUpdateResultClass = NSBatchUpdateResultClass{class: objc.GetClass("NSBatchUpdateResult")}
	})
	return _NSBatchUpdateResultClass
}

// GetNSBatchUpdateResultClass returns the class object for NSBatchUpdateResult.
func GetNSBatchUpdateResultClass() NSBatchUpdateResultClass {
	return getNSBatchUpdateResultClass()
}

type NSBatchUpdateResultClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSBatchUpdateResultClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSBatchUpdateResultClass) Alloc() NSBatchUpdateResult {
	rv := objc.Send[NSBatchUpdateResult](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The result returned when executing a batch update request.
//
// # Accessing Results
//
//   - [NSBatchUpdateResult.Result]: The result of a batch-update request, either the number of updated objects, the identifiers of the updated objects, or a status value.
//   - [NSBatchUpdateResult.ResultType]: The type of result that Core Data returns from the request.
//
// See: https://developer.apple.com/documentation/CoreData/NSBatchUpdateResult
type NSBatchUpdateResult struct {
	NSPersistentStoreResult
}

// NSBatchUpdateResultFromID constructs a [NSBatchUpdateResult] from an objc.ID.
//
// The result returned when executing a batch update request.
func NSBatchUpdateResultFromID(id objc.ID) NSBatchUpdateResult {
	return NSBatchUpdateResult{NSPersistentStoreResult: NSPersistentStoreResultFromID(id)}
}

// NOTE: NSBatchUpdateResult adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSBatchUpdateResult] class.
//
// # Accessing Results
//
//   - [INSBatchUpdateResult.Result]: The result of a batch-update request, either the number of updated objects, the identifiers of the updated objects, or a status value.
//   - [INSBatchUpdateResult.ResultType]: The type of result that Core Data returns from the request.
//
// See: https://developer.apple.com/documentation/CoreData/NSBatchUpdateResult
type INSBatchUpdateResult interface {
	INSPersistentStoreResult

	// Topic: Accessing Results

	// The result of a batch-update request, either the number of updated objects, the identifiers of the updated objects, or a status value.
	Result() objectivec.IObject
	// The type of result that Core Data returns from the request.
	ResultType() NSBatchUpdateRequestResultType
}

// Init initializes the instance.
func (b NSBatchUpdateResult) Init() NSBatchUpdateResult {
	rv := objc.Send[NSBatchUpdateResult](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b NSBatchUpdateResult) Autorelease() NSBatchUpdateResult {
	rv := objc.Send[NSBatchUpdateResult](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSBatchUpdateResult creates a new NSBatchUpdateResult instance.
func NewNSBatchUpdateResult() NSBatchUpdateResult {
	class := getNSBatchUpdateResultClass()
	rv := objc.Send[NSBatchUpdateResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The result of a batch-update request, either the number of updated objects,
// the identifiers of the updated objects, or a status value.
//
// See: https://developer.apple.com/documentation/CoreData/NSBatchUpdateResult/result
func (b NSBatchUpdateResult) Result() objectivec.IObject {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("result"))
	return objectivec.Object{ID: rv}
}

// The type of result that Core Data returns from the request.
//
// See: https://developer.apple.com/documentation/CoreData/NSBatchUpdateResult/resultType
func (b NSBatchUpdateResult) ResultType() NSBatchUpdateRequestResultType {
	rv := objc.Send[NSBatchUpdateRequestResultType](b.ID, objc.Sel("resultType"))
	return NSBatchUpdateRequestResultType(rv)
}
