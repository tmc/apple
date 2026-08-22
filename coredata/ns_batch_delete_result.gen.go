// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSBatchDeleteResult] class.
var (
	_NSBatchDeleteResultClass     NSBatchDeleteResultClass
	_NSBatchDeleteResultClassOnce sync.Once
)

func getNSBatchDeleteResultClass() NSBatchDeleteResultClass {
	_NSBatchDeleteResultClassOnce.Do(func() {
		_NSBatchDeleteResultClass = NSBatchDeleteResultClass{class: objc.GetClass("NSBatchDeleteResult")}
	})
	return _NSBatchDeleteResultClass
}

// GetNSBatchDeleteResultClass returns the class object for NSBatchDeleteResult.
func GetNSBatchDeleteResultClass() NSBatchDeleteResultClass {
	return getNSBatchDeleteResultClass()
}

type NSBatchDeleteResultClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSBatchDeleteResultClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSBatchDeleteResultClass) Alloc() NSBatchDeleteResult {
	rv := objc.Send[NSBatchDeleteResult](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that describes the result of a batch delete request.
//
// # Accessing the Result
//
//   - [NSBatchDeleteResult.Result]: The value the request returns after it executes.
//   - [NSBatchDeleteResult.ResultType]: The data type of the request’s result value.
//
// See: https://developer.apple.com/documentation/CoreData/NSBatchDeleteResult
type NSBatchDeleteResult struct {
	NSPersistentStoreResult
}

// NSBatchDeleteResultFromID constructs a [NSBatchDeleteResult] from an objc.ID.
//
// An object that describes the result of a batch delete request.
func NSBatchDeleteResultFromID(id objc.ID) NSBatchDeleteResult {
	return NSBatchDeleteResult{NSPersistentStoreResult: NSPersistentStoreResultFromID(id)}
}

// NOTE: NSBatchDeleteResult adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSBatchDeleteResult] class.
//
// # Accessing the Result
//
//   - [INSBatchDeleteResult.Result]: The value the request returns after it executes.
//   - [INSBatchDeleteResult.ResultType]: The data type of the request’s result value.
//
// See: https://developer.apple.com/documentation/CoreData/NSBatchDeleteResult
type INSBatchDeleteResult interface {
	INSPersistentStoreResult

	// Topic: Accessing the Result

	// The value the request returns after it executes.
	Result() objectivec.IObject
	// The data type of the request’s result value.
	ResultType() NSBatchDeleteRequestResultType
}

// Init initializes the instance.
func (b NSBatchDeleteResult) Init() NSBatchDeleteResult {
	rv := objc.Send[NSBatchDeleteResult](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b NSBatchDeleteResult) Autorelease() NSBatchDeleteResult {
	rv := objc.Send[NSBatchDeleteResult](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSBatchDeleteResult creates a new NSBatchDeleteResult instance.
func NewNSBatchDeleteResult() NSBatchDeleteResult {
	class := getNSBatchDeleteResultClass()
	rv := objc.Send[NSBatchDeleteResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The value the request returns after it executes.
//
// # Discussion
//
// Use [NSBatchDeleteResult.ResultType] to determine the kind of value this
// property contains, and then cast to the appropriate type as the following
// example shows:
//
// See: https://developer.apple.com/documentation/CoreData/NSBatchDeleteResult/result
func (b NSBatchDeleteResult) Result() objectivec.IObject {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("result"))
	return objectivec.Object{ID: rv}
}

// The data type of the request’s result value.
//
// # Discussion
//
// This property’s value is set to the request’s
// [NSBatchDeleteRequest.ResultType] property.
//
// See: https://developer.apple.com/documentation/CoreData/NSBatchDeleteResult/resultType
func (b NSBatchDeleteResult) ResultType() NSBatchDeleteRequestResultType {
	rv := objc.Send[NSBatchDeleteRequestResultType](b.ID, objc.Sel("resultType"))
	return NSBatchDeleteRequestResultType(rv)
}
