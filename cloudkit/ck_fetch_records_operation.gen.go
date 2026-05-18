// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKFetchRecordsOperation] class.
var (
	_CKFetchRecordsOperationClass     CKFetchRecordsOperationClass
	_CKFetchRecordsOperationClassOnce sync.Once
)

func getCKFetchRecordsOperationClass() CKFetchRecordsOperationClass {
	_CKFetchRecordsOperationClassOnce.Do(func() {
		_CKFetchRecordsOperationClass = CKFetchRecordsOperationClass{class: objc.GetClass("CKFetchRecordsOperation")}
	})
	return _CKFetchRecordsOperationClass
}

// GetCKFetchRecordsOperationClass returns the class object for CKFetchRecordsOperation.
func GetCKFetchRecordsOperationClass() CKFetchRecordsOperationClass {
	return getCKFetchRecordsOperationClass()
}

type CKFetchRecordsOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKFetchRecordsOperationClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKFetchRecordsOperationClass) Alloc() CKFetchRecordsOperation {
	rv := objc.Send[CKFetchRecordsOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An operation for retrieving records from a database.
//
// # Overview
//
// Use this operation to retrieve the entire contents of each record or only a
// subset of its contained values. As records become available, the operation
// object reports progress about the state of the operation to several
// different blocks, which you can use to process the results.
//
// Fetching records is a common use of CloudKit, even if your app doesn’t
// cache record IDs locally. For example, when you fetch a record related to
// the current record through a [CKReference] object, you use the ID in the
// reference to perform the fetch.
//
// The handlers you assign to process the fetched records execute serially on
// an internal queue that the fetch operation manages. You must provide
// handlers capable of executing on a background thread, so any tasks that
// require access to the main thread must redirect accordingly.
//
// In addition to data records, a fetch records operation can fetch the
// current user record. The [CKFetchRecordsOperation.FetchCurrentUserRecordOperation] method returns a
// specially configured operation object that retrieves the current user
// record. That record is a standard [CKRecord] object that has no content
// initially. You can add data to the user record and save it as necessary.
// Don’t store sensitive personal information, such as passwords, in the
// user record because other users of your app can access the discoverable
// user record in a public database. If you must store sensitive information
// about a user, do so in a separate record that is accessible only to that
// user.
//
// If you assign a closure to the [CKFetchRecordsOperation.CompletionBlock] property of the operation
// object, CloudKit calls it after the operation executes and returns its
// results. Use a closure to perform any housekeeping tasks for the operation,
// but don’t use it to process the results of the operation. The closure you
// specify should handle the failure of the operation to complete its task,
// whether due to an error or an explicit cancellation.
//
// # Creating a Record Fetch Operation
//
//   - [CKFetchRecordsOperation.InitWithRecordIDs]: Creates a fetch operation for retrieving the records with the specified IDs.
//
// # Configuring a Record Fetch Operation
//
//   - [CKFetchRecordsOperation.RecordIDs]: The record IDs of the records to fetch.
//   - [CKFetchRecordsOperation.SetRecordIDs]
//   - [CKFetchRecordsOperation.DesiredKeys]: The fields of the records to fetch.
//   - [CKFetchRecordsOperation.SetDesiredKeys]
//
// # Processing Record Fetch Results
//
//   - [CKFetchRecordsOperation.PerRecordProgressBlock]: The closure to execute with progress information for individual records.
//   - [CKFetchRecordsOperation.SetPerRecordProgressBlock]
//
// # Instance Properties
//
//   - [CKFetchRecordsOperation.FetchRecordsResultBlock]: The closure to execute after CloudKit retrieves all of the records.
//   - [CKFetchRecordsOperation.SetFetchRecordsResultBlock]
//   - [CKFetchRecordsOperation.PerRecordResultBlock]: The closure to execute when a record becomes available.
//   - [CKFetchRecordsOperation.SetPerRecordResultBlock]
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchRecordsOperation
type CKFetchRecordsOperation struct {
	CKDatabaseOperation
}

// CKFetchRecordsOperationFromID constructs a [CKFetchRecordsOperation] from an objc.ID.
//
// An operation for retrieving records from a database.
func CKFetchRecordsOperationFromID(id objc.ID) CKFetchRecordsOperation {
	return CKFetchRecordsOperation{CKDatabaseOperation: CKDatabaseOperationFromID(id)}
}

// NOTE: CKFetchRecordsOperation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKFetchRecordsOperation] class.
//
// # Creating a Record Fetch Operation
//
//   - [ICKFetchRecordsOperation.InitWithRecordIDs]: Creates a fetch operation for retrieving the records with the specified IDs.
//
// # Configuring a Record Fetch Operation
//
//   - [ICKFetchRecordsOperation.RecordIDs]: The record IDs of the records to fetch.
//   - [ICKFetchRecordsOperation.SetRecordIDs]
//   - [ICKFetchRecordsOperation.DesiredKeys]: The fields of the records to fetch.
//   - [ICKFetchRecordsOperation.SetDesiredKeys]
//
// # Processing Record Fetch Results
//
//   - [ICKFetchRecordsOperation.PerRecordProgressBlock]: The closure to execute with progress information for individual records.
//   - [ICKFetchRecordsOperation.SetPerRecordProgressBlock]
//
// # Instance Properties
//
//   - [ICKFetchRecordsOperation.FetchRecordsResultBlock]: The closure to execute after CloudKit retrieves all of the records.
//   - [ICKFetchRecordsOperation.SetFetchRecordsResultBlock]
//   - [ICKFetchRecordsOperation.PerRecordResultBlock]: The closure to execute when a record becomes available.
//   - [ICKFetchRecordsOperation.SetPerRecordResultBlock]
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchRecordsOperation
type ICKFetchRecordsOperation interface {
	ICKDatabaseOperation

	// Topic: Creating a Record Fetch Operation

	// Creates a fetch operation for retrieving the records with the specified IDs.
	InitWithRecordIDs(recordIDs []CKRecordID) CKFetchRecordsOperation

	// Topic: Configuring a Record Fetch Operation

	// The record IDs of the records to fetch.
	RecordIDs() []CKRecordID
	SetRecordIDs(value []CKRecordID)
	// The fields of the records to fetch.
	DesiredKeys() string
	SetDesiredKeys(value string)

	// Topic: Processing Record Fetch Results

	// The closure to execute with progress information for individual records.
	PerRecordProgressBlock() func(*CKRecordID, float64)
	SetPerRecordProgressBlock(value objc.ID)

	// Topic: Instance Properties

	// The closure to execute after CloudKit retrieves all of the records.
	FetchRecordsResultBlock() unsafe.Pointer
	SetFetchRecordsResultBlock(value unsafe.Pointer)
	// The closure to execute when a record becomes available.
	PerRecordResultBlock() unsafe.Pointer
	SetPerRecordResultBlock(value unsafe.Pointer)
}

// Init initializes the instance.
func (c CKFetchRecordsOperation) Init() CKFetchRecordsOperation {
	rv := objc.Send[CKFetchRecordsOperation](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKFetchRecordsOperation) Autorelease() CKFetchRecordsOperation {
	rv := objc.Send[CKFetchRecordsOperation](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKFetchRecordsOperation creates a new CKFetchRecordsOperation instance.
func NewCKFetchRecordsOperation() CKFetchRecordsOperation {
	class := getCKFetchRecordsOperationClass()
	rv := objc.Send[CKFetchRecordsOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a fetch operation for retrieving the records with the specified
// IDs.
//
// recordIDs: An array of [CKRecordID] objects that represents the records you want to
// retrieve. If you provide an empty array, you must set the [RecordIDs]
// property before you execute the operation.
//
// # Discussion
//
// A fetch operation retrieves all of a record’s fields, including any
// assets that those fields reference. If you want to minimize the amount of
// data that the operation returns, configure the [desiredKeys] property with
// only the keys that contain the values that you have an interest in.
//
// After initializing the operation, you must associate at least one progress
// handler with the operation (excluding the completion handler) to process
// the results.
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchRecordsOperation/init(recordIDs:)
//
// [desiredKeys]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordsOperation/desiredKeys-34l1l
func NewCKFetchRecordsOperationWithRecordIDs(recordIDs []CKRecordID) CKFetchRecordsOperation {
	instance := getCKFetchRecordsOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRecordIDs:"), objectivec.IObjectSliceToNSArray(recordIDs))
	return CKFetchRecordsOperationFromID(rv)
}

// Creates a fetch operation for retrieving the records with the specified
// IDs.
//
// recordIDs: An array of [CKRecordID] objects that represents the records you want to
// retrieve. If you provide an empty array, you must set the [RecordIDs]
// property before you execute the operation.
//
// # Discussion
//
// A fetch operation retrieves all of a record’s fields, including any
// assets that those fields reference. If you want to minimize the amount of
// data that the operation returns, configure the [desiredKeys] property with
// only the keys that contain the values that you have an interest in.
//
// After initializing the operation, you must associate at least one progress
// handler with the operation (excluding the completion handler) to process
// the results.
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchRecordsOperation/init(recordIDs:)
//
// [desiredKeys]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordsOperation/desiredKeys-34l1l
func (c CKFetchRecordsOperation) InitWithRecordIDs(recordIDs []CKRecordID) CKFetchRecordsOperation {
	rv := objc.Send[CKFetchRecordsOperation](c.ID, objc.Sel("initWithRecordIDs:"), objectivec.IObjectSliceToNSArray(recordIDs))
	return rv
}

// Returns a fetch operation for retrieving the current user record.
//
// # Discussion
//
// The returned operation object searches for the single record that
// corresponds to the current user record. You must associate at least one
// progress handler with the operation object (excluding the completion
// handler) to process the results.
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchRecordsOperation/fetchCurrentUserRecordOperation()
func (_CKFetchRecordsOperationClass CKFetchRecordsOperationClass) FetchCurrentUserRecordOperation() CKFetchRecordsOperation {
	rv := objc.Send[objc.ID](objc.ID(_CKFetchRecordsOperationClass.class), objc.Sel("fetchCurrentUserRecordOperation"))
	return CKFetchRecordsOperationFromID(rv)
}

// The record IDs of the records to fetch.
//
// # Discussion
//
// Use this property to view or change the IDs of the records you want to
// retrieve. If you use the operation that [FetchCurrentUserRecordOperation]
// returns, CloudKit ignores the contents of this property and sets its value
// to `nil`.
//
// If you intend to specify a value other than `nil`, do so before you execute
// the operation or add the operation to a queue. The records you fetch
// don’t need to be in the same record zone. The record ID for each record
// provides the zone information that CloudKit needs to fetch the
// corresponding record.
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchRecordsOperation/recordIDs
func (c CKFetchRecordsOperation) RecordIDs() []CKRecordID {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("recordIDs"))
	return objc.ConvertSlice(rv, func(id objc.ID) CKRecordID {
		return CKRecordIDFromID(id)
	})
}
func (c CKFetchRecordsOperation) SetRecordIDs(value []CKRecordID) {
	objc.Send[struct{}](c.ID, objc.Sel("setRecordIDs:"), objectivec.IObjectSliceToNSArray(value))
}

// The fields of the records to fetch.
//
// See: https://developer.apple.com/documentation/cloudkit/ckfetchrecordsoperation/desiredkeys-31bbq
func (c CKFetchRecordsOperation) DesiredKeys() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("desiredKeys"))
	return foundation.NSStringFromID(rv).String()
}
func (c CKFetchRecordsOperation) SetDesiredKeys(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setDesiredKeys:"), objc.String(value))
}

// The closure to execute with progress information for individual records.
//
// # Discussion
//
// This property is a closure that returns no value and has the following
// parameters:
//
// - The ID of the record to retrieve. - The amount of data, as a percentage,
// that CloudKit downloads for the record. The range is `0.0` to `1.0`, where
// `0.0` indicates that CloudKit hasn’t downloaded anything, and `1.0` means
// the download is complete.
//
// The fetch operation executes this closure one or more times for each record
// ID in the [RecordIDs] property. Each time the closure executes, it executes
// serially with respect to the other progress closures of the operation. You
// can use this closure to track the ongoing progress of the download
// operation.
//
// If you intend to use this closure to process results, set it before you
// execute the operation or add the operation to a queue.
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchRecordsOperation/perRecordProgressBlock
func (c CKFetchRecordsOperation) PerRecordProgressBlock() func(*CKRecordID, float64) {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("perRecordProgressBlock"))
	_ = rv
	return nil
}
func (c CKFetchRecordsOperation) SetPerRecordProgressBlock(value objc.ID) {
	objc.Send[struct{}](c.ID, objc.Sel("setPerRecordProgressBlock:"), value)
}

// The closure to execute after CloudKit retrieves all of the records.
//
// See: https://developer.apple.com/documentation/cloudkit/ckfetchrecordsoperation/fetchrecordsresultblock
func (c CKFetchRecordsOperation) FetchRecordsResultBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("fetchRecordsResultBlock"))
	return rv
}
func (c CKFetchRecordsOperation) SetFetchRecordsResultBlock(value unsafe.Pointer) {
	objc.Send[struct{}](c.ID, objc.Sel("setFetchRecordsResultBlock:"), value)
}

// The closure to execute when a record becomes available.
//
// See: https://developer.apple.com/documentation/cloudkit/ckfetchrecordsoperation/perrecordresultblock
func (c CKFetchRecordsOperation) PerRecordResultBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("perRecordResultBlock"))
	return rv
}
func (c CKFetchRecordsOperation) SetPerRecordResultBlock(value unsafe.Pointer) {
	objc.Send[struct{}](c.ID, objc.Sel("setPerRecordResultBlock:"), value)
}
