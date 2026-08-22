// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSBatchInsertResult] class.
var (
	_NSBatchInsertResultClass     NSBatchInsertResultClass
	_NSBatchInsertResultClassOnce sync.Once
)

func getNSBatchInsertResultClass() NSBatchInsertResultClass {
	_NSBatchInsertResultClassOnce.Do(func() {
		_NSBatchInsertResultClass = NSBatchInsertResultClass{class: objc.GetClass("NSBatchInsertResult")}
	})
	return _NSBatchInsertResultClass
}

// GetNSBatchInsertResultClass returns the class object for NSBatchInsertResult.
func GetNSBatchInsertResultClass() NSBatchInsertResultClass {
	return getNSBatchInsertResultClass()
}

type NSBatchInsertResultClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSBatchInsertResultClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSBatchInsertResultClass) Alloc() NSBatchInsertResult {
	rv := objc.Send[NSBatchInsertResult](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The result that Core Data returns when executing a batch-insertion request.
//
// # Accessing Results
//
//   - [NSBatchInsertResult.Result]: The result of a batch-insertion request.
//   - [NSBatchInsertResult.ResultType]: The type of result that Core Data returns from this request.
//
// See: https://developer.apple.com/documentation/CoreData/NSBatchInsertResult
type NSBatchInsertResult struct {
	NSPersistentStoreResult
}

// NSBatchInsertResultFromID constructs a [NSBatchInsertResult] from an objc.ID.
//
// The result that Core Data returns when executing a batch-insertion request.
func NSBatchInsertResultFromID(id objc.ID) NSBatchInsertResult {
	return NSBatchInsertResult{NSPersistentStoreResult: NSPersistentStoreResultFromID(id)}
}

// NOTE: NSBatchInsertResult adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSBatchInsertResult] class.
//
// # Accessing Results
//
//   - [INSBatchInsertResult.Result]: The result of a batch-insertion request.
//   - [INSBatchInsertResult.ResultType]: The type of result that Core Data returns from this request.
//
// See: https://developer.apple.com/documentation/CoreData/NSBatchInsertResult
type INSBatchInsertResult interface {
	INSPersistentStoreResult

	// Topic: Accessing Results

	// The result of a batch-insertion request.
	Result() objectivec.IObject
	// The type of result that Core Data returns from this request.
	ResultType() NSBatchInsertRequestResultType
}

// Init initializes the instance.
func (b NSBatchInsertResult) Init() NSBatchInsertResult {
	rv := objc.Send[NSBatchInsertResult](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b NSBatchInsertResult) Autorelease() NSBatchInsertResult {
	rv := objc.Send[NSBatchInsertResult](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSBatchInsertResult creates a new NSBatchInsertResult instance.
func NewNSBatchInsertResult() NSBatchInsertResult {
	class := getNSBatchInsertResultClass()
	rv := objc.Send[NSBatchInsertResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The result of a batch-insertion request.
//
// # Discussion
//
// Cast the result to the type corresponding to
// [NSBatchInsertResult.ResultType] to inspect it. The following example shows
// how to inspect a result type of
// [NSBatchInsertRequestResultType.statusOnly].
//
// See: https://developer.apple.com/documentation/CoreData/NSBatchInsertResult/result
//
// [NSBatchInsertRequestResultType.statusOnly]: https://developer.apple.com/documentation/CoreData/NSBatchInsertRequestResultType/statusOnly
func (b NSBatchInsertResult) Result() objectivec.IObject {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("result"))
	return objectivec.Object{ID: rv}
}

// The type of result that Core Data returns from this request.
//
// See: https://developer.apple.com/documentation/CoreData/NSBatchInsertResult/resultType
func (b NSBatchInsertResult) ResultType() NSBatchInsertRequestResultType {
	rv := objc.Send[NSBatchInsertRequestResultType](b.ID, objc.Sel("resultType"))
	return NSBatchInsertRequestResultType(rv)
}
