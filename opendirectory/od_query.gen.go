// Code generated from Apple documentation for OpenDirectory. DO NOT EDIT.

package opendirectory

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ODQuery] class.
var (
	_ODQueryClass     ODQueryClass
	_ODQueryClassOnce sync.Once
)

func getODQueryClass() ODQueryClass {
	_ODQueryClassOnce.Do(func() {
		_ODQueryClass = ODQueryClass{class: objc.GetClass("ODQuery")}
	})
	return _ODQueryClass
}

// GetODQueryClass returns the class object for ODQuery.
func GetODQueryClass() ODQueryClass {
	return getODQueryClass()
}

type ODQueryClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (oc ODQueryClass) Class() objc.Class {
	return oc.class
}

// Alloc allocates memory for a new instance of the class.
func (oc ODQueryClass) Alloc() ODQuery {
	rv := objc.Send[ODQuery](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// An [ODQuery] object serves as a Cocoa wrapper for an Open Directory query.
//
// # Creating and Initializing a Query
//
//   - [ODQuery.InitWithNodeForRecordTypesAttributeMatchTypeQueryValuesReturnAttributesMaximumResultsError]: Creates a query object with provided parameters.
//
// # Managing Asynchronous Queries
//
//   - [ODQuery.Delegate]: The query’s delegate.
//   - [ODQuery.SetDelegate]
//   - [ODQuery.OperationQueue]: The queue on which asynchronous results are delivered to the delegate.
//   - [ODQuery.SetOperationQueue]
//   - [ODQuery.ScheduleInRunLoopForMode]: Retrieves results from a query asynchronously by scheduling the query in a run loop.
//   - [ODQuery.RemoveFromRunLoopForMode]: Removes the query from a specified run loop.
//   - [ODQuery.Synchronize]: Restarts a query, disposing of any results it has obtained.
//
// # Managing Synchronous Queries
//
//   - [ODQuery.ResultsAllowingPartialError]: Returns results from a query synchronously.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODQuery
type ODQuery struct {
	objectivec.Object
}

// ODQueryFromID constructs a [ODQuery] from an objc.ID.
//
// An [ODQuery] object serves as a Cocoa wrapper for an Open Directory query.
func ODQueryFromID(id objc.ID) ODQuery {
	return ODQuery{objectivec.Object{ID: id}}
}

// NOTE: ODQuery adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [ODQuery] class.
//
// # Creating and Initializing a Query
//
//   - [IODQuery.InitWithNodeForRecordTypesAttributeMatchTypeQueryValuesReturnAttributesMaximumResultsError]: Creates a query object with provided parameters.
//
// # Managing Asynchronous Queries
//
//   - [IODQuery.Delegate]: The query’s delegate.
//   - [IODQuery.SetDelegate]
//   - [IODQuery.OperationQueue]: The queue on which asynchronous results are delivered to the delegate.
//   - [IODQuery.SetOperationQueue]
//   - [IODQuery.ScheduleInRunLoopForMode]: Retrieves results from a query asynchronously by scheduling the query in a run loop.
//   - [IODQuery.RemoveFromRunLoopForMode]: Removes the query from a specified run loop.
//   - [IODQuery.Synchronize]: Restarts a query, disposing of any results it has obtained.
//
// # Managing Synchronous Queries
//
//   - [IODQuery.ResultsAllowingPartialError]: Returns results from a query synchronously.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODQuery
type IODQuery interface {
	objectivec.IObject

	// Topic: Creating and Initializing a Query

	// Creates a query object with provided parameters.
	InitWithNodeForRecordTypesAttributeMatchTypeQueryValuesReturnAttributesMaximumResultsError(inNode IODNode, inRecordTypeOrList objectivec.IObject, inAttribute unsafe.Pointer, inMatchType ODMatchType, inQueryValueOrList objectivec.IObject, inReturnAttributeOrList objectivec.IObject, inMaximumResults int) (ODQuery, error)

	// Topic: Managing Asynchronous Queries

	// The query’s delegate.
	Delegate() ODQueryDelegate
	SetDelegate(value ODQueryDelegate)
	// The queue on which asynchronous results are delivered to the delegate.
	OperationQueue() foundation.OperationQueue
	SetOperationQueue(value foundation.OperationQueue)
	// Retrieves results from a query asynchronously by scheduling the query in a run loop.
	ScheduleInRunLoopForMode(inRunLoop foundation.NSRunLoop, inMode string)
	// Removes the query from a specified run loop.
	RemoveFromRunLoopForMode(inRunLoop foundation.NSRunLoop, inMode string)
	// Restarts a query, disposing of any results it has obtained.
	Synchronize()

	// Topic: Managing Synchronous Queries

	// Returns results from a query synchronously.
	ResultsAllowingPartialError(inAllowPartialResults bool) (foundation.INSArray, error)
}

// Init initializes the instance.
func (o ODQuery) Init() ODQuery {
	rv := objc.Send[ODQuery](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o ODQuery) Autorelease() ODQuery {
	rv := objc.Send[ODQuery](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewODQuery creates a new ODQuery instance.
func NewODQuery() ODQuery {
	class := getODQueryClass()
	rv := objc.Send[ODQuery](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a query object with provided parameters.
//
// inNode: The node to query.
//
// inRecordTypeOrList: The type or types of record to query. Can be an [NSString] object for a
// single type or an [NSArray] object containing [NSString] objects for
// multiple types.
//
// inAttribute: The name of the attribute to query.
//
// inMatchType: The type of query.
//
// inQueryValueOrList: The value or values to query in the attribute. Can be an [NSString] object
// or an [NSData] object for a single value, or an [NSArray] containing
// [NSString] and [NSData] objects for multiple values.
//
// inReturnAttributeOrList: The attribute or attributes to be returned from the query. Can be an
// [NSString] object for a single attribute or an [NSArray] object containing
// [NSString] objects for multiple attributes. Passing `nil` is equivalent to
// passing `kODAttributeTypeStandardOnly`.
//
// inMaximumResults: The maximum number of values to return.
//
// # Return Value
//
// The initialized query.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODQuery/init(node:forRecordTypes:attribute:matchType:queryValues:returnAttributes:maximumResults:)
func NewODQueryWithNodeForRecordTypesAttributeMatchTypeQueryValuesReturnAttributesMaximumResultsError(inNode IODNode, inRecordTypeOrList objectivec.IObject, inAttribute unsafe.Pointer, inMatchType ODMatchType, inQueryValueOrList objectivec.IObject, inReturnAttributeOrList objectivec.IObject, inMaximumResults int) (ODQuery, error) {
	var errorPtr objc.ID
	instance := getODQueryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithNode:forRecordTypes:attribute:matchType:queryValues:returnAttributes:maximumResults:error:"), inNode, inRecordTypeOrList, inAttribute, inMatchType, inQueryValueOrList, inReturnAttributeOrList, inMaximumResults, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return ODQuery{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return ODQuery{}, objc.ErrInitFailed
	}
	return ODQueryFromID(rv), nil
}

// Creates a query object with provided parameters.
//
// inNode: The node to query.
//
// inRecordTypeOrList: The type or types of record to query. Can be an [NSString] object for a
// single type or an [NSArray] object containing [NSString] objects for
// multiple types.
//
// inAttribute: The name of the attribute to query.
//
// inMatchType: The type of query.
//
// inQueryValueOrList: The value or values to query in the attribute. Can be an [NSString] object
// or an [NSData] object for a single value, or an [NSArray] containing
// [NSString] and [NSData] objects for multiple values.
//
// inReturnAttributeOrList: The attribute or attributes to be returned from the query. Can be an
// [NSString] object for a single attribute or an [NSArray] object containing
// [NSString] objects for multiple attributes. Passing `nil` is equivalent to
// passing `kODAttributeTypeStandardOnly`.
//
// inMaximumResults: The maximum number of values to return.
//
// # Return Value
//
// The initialized query.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODQuery/init(node:forRecordTypes:attribute:matchType:queryValues:returnAttributes:maximumResults:)
func (o ODQuery) InitWithNodeForRecordTypesAttributeMatchTypeQueryValuesReturnAttributesMaximumResultsError(inNode IODNode, inRecordTypeOrList objectivec.IObject, inAttribute unsafe.Pointer, inMatchType ODMatchType, inQueryValueOrList objectivec.IObject, inReturnAttributeOrList objectivec.IObject, inMaximumResults int) (ODQuery, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](o.ID, objc.Sel("initWithNode:forRecordTypes:attribute:matchType:queryValues:returnAttributes:maximumResults:error:"), inNode, inRecordTypeOrList, inAttribute, inMatchType, inQueryValueOrList, inReturnAttributeOrList, inMaximumResults, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return ODQuery{}, foundation.NSErrorFrom(errorPtr)
	}
	return ODQueryFromID(rv), nil

}

// Retrieves results from a query asynchronously by scheduling the query in a
// run loop.
//
// inRunLoop: The run loop.
//
// inMode: The mode of the run loop.
//
// # Discussion
//
// A delegate must be set prior to calling this method; otherwise, results may
// be lost due to the lack of a receiver.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODQuery/schedule(in:forMode:)
func (o ODQuery) ScheduleInRunLoopForMode(inRunLoop foundation.NSRunLoop, inMode string) {
	objc.Send[objc.ID](o.ID, objc.Sel("scheduleInRunLoop:forMode:"), inRunLoop, objc.String(inMode))
}

// Removes the query from a specified run loop.
//
// inRunLoop: The run loop.
//
// inMode: The mode to remove the query from.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODQuery/remove(from:forMode:)
func (o ODQuery) RemoveFromRunLoopForMode(inRunLoop foundation.NSRunLoop, inMode string) {
	objc.Send[objc.ID](o.ID, objc.Sel("removeFromRunLoop:forMode:"), inRunLoop, objc.String(inMode))
}

// Restarts a query, disposing of any results it has obtained.
//
// # Discussion
//
// If the query was originally scheduled in a run loop with
// [ODQuery.ScheduleInRunLoopForMode], the delegate is called with `inResults`
// set to `nil`, `[inError code]` set to `kODErrorQuerySynchronize`, and
// `[inError domain]` set to `kODErrorDomainFramework`.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODQuery/synchronize()
func (o ODQuery) Synchronize() {
	objc.Send[objc.ID](o.ID, objc.Sel("synchronize"))
}

// Returns results from a query synchronously.
//
// inAllowPartialResults: If true, only immediately available results are returned; otherwise, the
// function waits until all results are available.
//
// # Return Value
//
// The results of the query in an array of [ODRecord] objects.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODQuery/resultsAllowingPartial(_:)
func (o ODQuery) ResultsAllowingPartialError(inAllowPartialResults bool) (foundation.INSArray, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](o.ID, objc.Sel("resultsAllowingPartial:error:"), inAllowPartialResults, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return foundation.NSArrayFromID(rv), nil

}

// The query’s delegate.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODQuery/delegate
func (o ODQuery) Delegate() ODQueryDelegate {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("delegate"))
	return ODQueryDelegateObjectFromID(rv)
}
func (o ODQuery) SetDelegate(value ODQueryDelegate) {
	objc.Send[struct{}](o.ID, objc.Sel("setDelegate:"), value)
}

// The queue on which asynchronous results are delivered to the delegate.
//
// See: https://developer.apple.com/documentation/OpenDirectory/ODQuery/operationQueue
func (o ODQuery) OperationQueue() foundation.OperationQueue {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("operationQueue"))
	return foundation.OperationQueueFromID(objc.ID(rv))
}
func (o ODQuery) SetOperationQueue(value foundation.OperationQueue) {
	objc.Send[struct{}](o.ID, objc.Sel("setOperationQueue:"), value)
}
