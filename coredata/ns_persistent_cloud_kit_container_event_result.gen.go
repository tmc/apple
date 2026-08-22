// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSPersistentCloudKitContainerEventResult] class.
var (
	_NSPersistentCloudKitContainerEventResultClass     NSPersistentCloudKitContainerEventResultClass
	_NSPersistentCloudKitContainerEventResultClassOnce sync.Once
)

func getNSPersistentCloudKitContainerEventResultClass() NSPersistentCloudKitContainerEventResultClass {
	_NSPersistentCloudKitContainerEventResultClassOnce.Do(func() {
		_NSPersistentCloudKitContainerEventResultClass = NSPersistentCloudKitContainerEventResultClass{class: objc.GetClass("NSPersistentCloudKitContainerEventResult")}
	})
	return _NSPersistentCloudKitContainerEventResultClass
}

// GetNSPersistentCloudKitContainerEventResultClass returns the class object for NSPersistentCloudKitContainerEventResult.
func GetNSPersistentCloudKitContainerEventResultClass() NSPersistentCloudKitContainerEventResultClass {
	return getNSPersistentCloudKitContainerEventResultClass()
}

type NSPersistentCloudKitContainerEventResultClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSPersistentCloudKitContainerEventResultClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPersistentCloudKitContainerEventResultClass) Alloc() NSPersistentCloudKitContainerEventResult {
	rv := objc.Send[NSPersistentCloudKitContainerEventResult](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The result of a request to fetch persistent CloudKit container events.
//
// # Handling Event Results
//
//   - [NSPersistentCloudKitContainerEventResult.Result]: The result of the persistent CloudKit container event request, which the result type determines.
//   - [NSPersistentCloudKitContainerEventResult.ResultType]: The type of result that the CloudKit container event fetch request returns.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainerEventResult
type NSPersistentCloudKitContainerEventResult struct {
	NSPersistentStoreResult
}

// NSPersistentCloudKitContainerEventResultFromID constructs a [NSPersistentCloudKitContainerEventResult] from an objc.ID.
//
// The result of a request to fetch persistent CloudKit container events.
func NSPersistentCloudKitContainerEventResultFromID(id objc.ID) NSPersistentCloudKitContainerEventResult {
	return NSPersistentCloudKitContainerEventResult{NSPersistentStoreResult: NSPersistentStoreResultFromID(id)}
}

// NOTE: NSPersistentCloudKitContainerEventResult adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPersistentCloudKitContainerEventResult] class.
//
// # Handling Event Results
//
//   - [INSPersistentCloudKitContainerEventResult.Result]: The result of the persistent CloudKit container event request, which the result type determines.
//   - [INSPersistentCloudKitContainerEventResult.ResultType]: The type of result that the CloudKit container event fetch request returns.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainerEventResult
type INSPersistentCloudKitContainerEventResult interface {
	INSPersistentStoreResult

	// Topic: Handling Event Results

	// The result of the persistent CloudKit container event request, which the result type determines.
	Result() objectivec.IObject
	// The type of result that the CloudKit container event fetch request returns.
	ResultType() NSPersistentCloudKitContainerEventResultType
}

// Init initializes the instance.
func (p NSPersistentCloudKitContainerEventResult) Init() NSPersistentCloudKitContainerEventResult {
	rv := objc.Send[NSPersistentCloudKitContainerEventResult](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPersistentCloudKitContainerEventResult) Autorelease() NSPersistentCloudKitContainerEventResult {
	rv := objc.Send[NSPersistentCloudKitContainerEventResult](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPersistentCloudKitContainerEventResult creates a new NSPersistentCloudKitContainerEventResult instance.
func NewNSPersistentCloudKitContainerEventResult() NSPersistentCloudKitContainerEventResult {
	class := getNSPersistentCloudKitContainerEventResultClass()
	rv := objc.Send[NSPersistentCloudKitContainerEventResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The result of the persistent CloudKit container event request, which the
// result type determines.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainerEventResult/result
func (p NSPersistentCloudKitContainerEventResult) Result() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("result"))
	return objectivec.Object{ID: rv}
}

// The type of result that the CloudKit container event fetch request returns.
//
// See: https://developer.apple.com/documentation/CoreData/NSPersistentCloudKitContainerEventResult/resultType-swift.property
func (p NSPersistentCloudKitContainerEventResult) ResultType() NSPersistentCloudKitContainerEventResultType {
	rv := objc.Send[NSPersistentCloudKitContainerEventResultType](p.ID, objc.Sel("resultType"))
	return NSPersistentCloudKitContainerEventResultType(rv)
}
