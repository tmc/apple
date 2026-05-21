// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKFetchRecordZonesOperation] class.
var (
	_CKFetchRecordZonesOperationClass     CKFetchRecordZonesOperationClass
	_CKFetchRecordZonesOperationClassOnce sync.Once
)

func getCKFetchRecordZonesOperationClass() CKFetchRecordZonesOperationClass {
	_CKFetchRecordZonesOperationClassOnce.Do(func() {
		_CKFetchRecordZonesOperationClass = CKFetchRecordZonesOperationClass{class: objc.GetClass("CKFetchRecordZonesOperation")}
	})
	return _CKFetchRecordZonesOperationClass
}

// GetCKFetchRecordZonesOperationClass returns the class object for CKFetchRecordZonesOperation.
func GetCKFetchRecordZonesOperationClass() CKFetchRecordZonesOperationClass {
	return getCKFetchRecordZonesOperationClass()
}

type CKFetchRecordZonesOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKFetchRecordZonesOperationClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKFetchRecordZonesOperationClass) Alloc() CKFetchRecordZonesOperation {
	rv := objc.Send[CKFetchRecordZonesOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An operation for retrieving record zones from a database.
//
// # Overview
//
// Use this operation object to fetch record zones so that you can ascertain
// their capabilities.
//
// If you assign a handler to the [completionBlock] property of the operation,
// CloudKit calls it after the operation executes and returns its results. You
// can use the handler to perform any housekeeping tasks that relate to the
// operation, but don’t use it to process the results of the operation. The
// handler you specify should manage any failures, whether due to an error or
// an explicit cancellation.
//
// # Initializing the Zone Fetch Operation
//
//   - [CKFetchRecordZonesOperation.InitWithRecordZoneIDs]: Creates an operation for fetching the specified record zones.
//
// # Configuring a Zone Fetch Operation
//
//   - [CKFetchRecordZonesOperation.RecordZoneIDs]: The IDs of the record zones to retrieve.
//   - [CKFetchRecordZonesOperation.SetRecordZoneIDs]
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZonesOperation
//
// [completionBlock]: https://developer.apple.com/documentation/Foundation/Operation/completionBlock
type CKFetchRecordZonesOperation struct {
	CKDatabaseOperation
}

// CKFetchRecordZonesOperationFromID constructs a [CKFetchRecordZonesOperation] from an objc.ID.
//
// An operation for retrieving record zones from a database.
func CKFetchRecordZonesOperationFromID(id objc.ID) CKFetchRecordZonesOperation {
	return CKFetchRecordZonesOperation{CKDatabaseOperation: CKDatabaseOperationFromID(id)}
}

// NOTE: CKFetchRecordZonesOperation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKFetchRecordZonesOperation] class.
//
// # Initializing the Zone Fetch Operation
//
//   - [ICKFetchRecordZonesOperation.InitWithRecordZoneIDs]: Creates an operation for fetching the specified record zones.
//
// # Configuring a Zone Fetch Operation
//
//   - [ICKFetchRecordZonesOperation.RecordZoneIDs]: The IDs of the record zones to retrieve.
//   - [ICKFetchRecordZonesOperation.SetRecordZoneIDs]
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZonesOperation
type ICKFetchRecordZonesOperation interface {
	ICKDatabaseOperation

	// Topic: Initializing the Zone Fetch Operation

	// Creates an operation for fetching the specified record zones.
	InitWithRecordZoneIDs(zoneIDs []CKRecordZoneID) CKFetchRecordZonesOperation

	// Topic: Configuring a Zone Fetch Operation

	// The IDs of the record zones to retrieve.
	RecordZoneIDs() []CKRecordZoneID
	SetRecordZoneIDs(value []CKRecordZoneID)
}

// Init initializes the instance.
func (c CKFetchRecordZonesOperation) Init() CKFetchRecordZonesOperation {
	rv := objc.Send[CKFetchRecordZonesOperation](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKFetchRecordZonesOperation) Autorelease() CKFetchRecordZonesOperation {
	rv := objc.Send[CKFetchRecordZonesOperation](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKFetchRecordZonesOperation creates a new CKFetchRecordZonesOperation instance.
func NewCKFetchRecordZonesOperation() CKFetchRecordZonesOperation {
	class := getCKFetchRecordZonesOperationClass()
	rv := objc.Send[CKFetchRecordZonesOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an operation for fetching the specified record zones.
//
// zoneIDs: An array of[CKRecordZoneID] objects that represents the zones you want to
// retrieve. If you provide an empty array, you must set the
// [CKFetchRecordZonesOperation.RecordZoneIDs] property before you execute the
// operation.
//
// # Discussion
//
// After creating the operation, assign a value to the
// [fetchRecordZonesCompletionBlock] property so you can process the results.
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZonesOperation/init(recordZoneIDs:)
//
// [fetchRecordZonesCompletionBlock]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZonesOperation/fetchRecordZonesCompletionBlock
func NewCKFetchRecordZonesOperationWithRecordZoneIDs(zoneIDs []CKRecordZoneID) CKFetchRecordZonesOperation {
	instance := getCKFetchRecordZonesOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRecordZoneIDs:"), objectivec.IObjectSliceToNSArray(zoneIDs))
	return CKFetchRecordZonesOperationFromID(rv)
}

// Creates an operation for fetching the specified record zones.
//
// zoneIDs: An array of[CKRecordZoneID] objects that represents the zones you want to
// retrieve. If you provide an empty array, you must set the
// [CKFetchRecordZonesOperation.RecordZoneIDs] property before you execute the
// operation.
//
// # Discussion
//
// After creating the operation, assign a value to the
// [fetchRecordZonesCompletionBlock] property so you can process the results.
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZonesOperation/init(recordZoneIDs:)
//
// [fetchRecordZonesCompletionBlock]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZonesOperation/fetchRecordZonesCompletionBlock
func (c CKFetchRecordZonesOperation) InitWithRecordZoneIDs(zoneIDs []CKRecordZoneID) CKFetchRecordZonesOperation {
	rv := objc.Send[CKFetchRecordZonesOperation](c.ID, objc.Sel("initWithRecordZoneIDs:"), objectivec.IObjectSliceToNSArray(zoneIDs))
	return rv
}

// Returns an operation for fetching all record zones in the current database.
//
// # Discussion
//
// Assign a value to the [fetchRecordZonesCompletionBlock] property of the
// operation that this method returns so that you can process the results.
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZonesOperation/fetchAllRecordZonesOperation()
//
// [fetchRecordZonesCompletionBlock]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZonesOperation/fetchRecordZonesCompletionBlock
func (_CKFetchRecordZonesOperationClass CKFetchRecordZonesOperationClass) FetchAllRecordZonesOperation() CKFetchRecordZonesOperation {
	rv := objc.Send[objc.ID](objc.ID(_CKFetchRecordZonesOperationClass.class), objc.Sel("fetchAllRecordZonesOperation"))
	return CKFetchRecordZonesOperationFromID(rv)
}

// The IDs of the record zones to retrieve.
//
// # Discussion
//
// Use this property to view or change the IDs of the record zones you want to
// retrieve. If you intend to change the value of this property, do so before
// you execute the operation or submit the operation to a queue.
//
// If you use the operation that
// [CKFetchRecordZonesOperationClass.FetchAllRecordZonesOperation] returns,
// CloudKit ignores the contents of this property and sets its value to `nil`.
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZonesOperation/recordZoneIDs
func (c CKFetchRecordZonesOperation) RecordZoneIDs() []CKRecordZoneID {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("recordZoneIDs"))
	return objc.ConvertSlice(rv, func(id objc.ID) CKRecordZoneID {
		return CKRecordZoneIDFromID(id)
	})
}
func (c CKFetchRecordZonesOperation) SetRecordZoneIDs(value []CKRecordZoneID) {
	objc.Send[struct{}](c.ID, objc.Sel("setRecordZoneIDs:"), objectivec.IObjectSliceToNSArray(value))
}
