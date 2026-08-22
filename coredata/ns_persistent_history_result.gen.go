// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSPersistentHistoryResult] class.
var (
	_NSPersistentHistoryResultClass     NSPersistentHistoryResultClass
	_NSPersistentHistoryResultClassOnce sync.Once
)

func getNSPersistentHistoryResultClass() NSPersistentHistoryResultClass {
	_NSPersistentHistoryResultClassOnce.Do(func() {
		_NSPersistentHistoryResultClass = NSPersistentHistoryResultClass{class: objc.GetClass("NSPersistentHistoryResult")}
	})
	return _NSPersistentHistoryResultClass
}

// GetNSPersistentHistoryResultClass returns the class object for NSPersistentHistoryResult.
func GetNSPersistentHistoryResultClass() NSPersistentHistoryResultClass {
	return getNSPersistentHistoryResultClass()
}

type NSPersistentHistoryResultClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSPersistentHistoryResultClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPersistentHistoryResultClass) Alloc() NSPersistentHistoryResult {
	rv := objc.Send[NSPersistentHistoryResult](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The result of a request to fetch persistent history.
//
// # Inspecting History Results
//
//   - [NSPersistentHistoryResult.Result]: The result of the history request determined by the persistent history result type.
//   - [NSPersistentHistoryResult.ResultType]: The type of result that the persistent history change request returns.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryResult
type NSPersistentHistoryResult struct {
	NSPersistentStoreResult
}

// NSPersistentHistoryResultFromID constructs a [NSPersistentHistoryResult] from an objc.ID.
//
// The result of a request to fetch persistent history.
func NSPersistentHistoryResultFromID(id objc.ID) NSPersistentHistoryResult {
	return NSPersistentHistoryResult{NSPersistentStoreResult: NSPersistentStoreResultFromID(id)}
}

// NOTE: NSPersistentHistoryResult adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPersistentHistoryResult] class.
//
// # Inspecting History Results
//
//   - [INSPersistentHistoryResult.Result]: The result of the history request determined by the persistent history result type.
//   - [INSPersistentHistoryResult.ResultType]: The type of result that the persistent history change request returns.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryResult
type INSPersistentHistoryResult interface {
	INSPersistentStoreResult

	// Topic: Inspecting History Results

	// The result of the history request determined by the persistent history result type.
	Result() objectivec.IObject
	// The type of result that the persistent history change request returns.
	ResultType() NSPersistentHistoryResultType
}

// Init initializes the instance.
func (p NSPersistentHistoryResult) Init() NSPersistentHistoryResult {
	rv := objc.Send[NSPersistentHistoryResult](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPersistentHistoryResult) Autorelease() NSPersistentHistoryResult {
	rv := objc.Send[NSPersistentHistoryResult](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPersistentHistoryResult creates a new NSPersistentHistoryResult instance.
func NewNSPersistentHistoryResult() NSPersistentHistoryResult {
	class := getNSPersistentHistoryResultClass()
	rv := objc.Send[NSPersistentHistoryResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The result of the history request determined by the persistent history
// result type.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryResult/result
func (p NSPersistentHistoryResult) Result() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("result"))
	return objectivec.Object{ID: rv}
}

// The type of result that the persistent history change request returns.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentHistoryResult/resultType
func (p NSPersistentHistoryResult) ResultType() NSPersistentHistoryResultType {
	rv := objc.Send[NSPersistentHistoryResultType](p.ID, objc.Sel("resultType"))
	return NSPersistentHistoryResultType(rv)
}
