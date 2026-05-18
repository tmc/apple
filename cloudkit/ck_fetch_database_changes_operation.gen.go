// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
)

// The class instance for the [CKFetchDatabaseChangesOperation] class.
var (
	_CKFetchDatabaseChangesOperationClass     CKFetchDatabaseChangesOperationClass
	_CKFetchDatabaseChangesOperationClassOnce sync.Once
)

func getCKFetchDatabaseChangesOperationClass() CKFetchDatabaseChangesOperationClass {
	_CKFetchDatabaseChangesOperationClassOnce.Do(func() {
		_CKFetchDatabaseChangesOperationClass = CKFetchDatabaseChangesOperationClass{class: objc.GetClass("CKFetchDatabaseChangesOperation")}
	})
	return _CKFetchDatabaseChangesOperationClass
}

// GetCKFetchDatabaseChangesOperationClass returns the class object for CKFetchDatabaseChangesOperation.
func GetCKFetchDatabaseChangesOperationClass() CKFetchDatabaseChangesOperationClass {
	return getCKFetchDatabaseChangesOperationClass()
}

type CKFetchDatabaseChangesOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKFetchDatabaseChangesOperationClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKFetchDatabaseChangesOperationClass) Alloc() CKFetchDatabaseChangesOperation {
	rv := objc.Send[CKFetchDatabaseChangesOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An operation that fetches database changes.
//
// # Overview
//
// Use this operation to fetch all record zone changes in a database. This
// includes new record zones, changed zones — including deleted or purged
// zones — and zones that contain record changes. When you create the
// operation, you provide a server change token, which is an opaque token that
// represents a specific point in the database’s history. CloudKit returns
// only the changes that occur after that point. For your app’s first fetch,
// or to refetch every change in the database’s history, use `nil` instead.
//
// The operation yields new change tokens during its execution, and issues a
// final change token when it completes without error. The change tokens
// conform to [NSSecureCoding] and are safe to cache on-disk. This
// operation’s tokens aren’t compatible with
// [CKFetchRecordZoneChangesOperation] so you should segregate them in your
// cache. Don’t infer any behavior or order from the tokens’ contents.
//
// When your app launches for the first time, use this operation to fetch all
// the database’s changes. Cache the results on-device and use
// [CKDatabaseSubscription] to subscribe to future changes. Fetch those
// changes on receipt of the push notifications the subscription generates.
// It’s not necessary to perform a fetch each time your app launches, or to
// schedule fetches at regular intervals.
//
// The operation calls [CKFetchDatabaseChangesOperation.RecordZoneWithIDChangedBlock] for each zone that
// contains record changes. It also calls it for new and modified record
// zones. Store the IDs that CloudKit provides to this callback. Use those IDs
// with [CKFetchRecordZoneChangesOperation] to fetch the corresponding
// changes. There are similar callbacks for deleted and purged record zones.
//
// To run the operation, add it to the corresponding database’s operation
// queue. The operation executes its callbacks on a private serial queue.
//
// The following example shows how to create the operation, configure its
// callbacks, and execute it. For brevity, it omits the delete, and purge
// callbacks.
//
// # Creating an Operation
//
//   - [CKFetchDatabaseChangesOperation.InitWithPreviousServerChangeToken]: Creates an operation for fetching database changes.
//
// # Configuring the Operation
//
//   - [CKFetchDatabaseChangesOperation.FetchAllChanges]: A Boolean value that indicates whether to send repeated requests to the server.
//   - [CKFetchDatabaseChangesOperation.SetFetchAllChanges]
//   - [CKFetchDatabaseChangesOperation.PreviousServerChangeToken]: The server change token.
//   - [CKFetchDatabaseChangesOperation.SetPreviousServerChangeToken]
//   - [CKFetchDatabaseChangesOperation.ResultsLimit]: The maximum number of results that the operation fetches.
//   - [CKFetchDatabaseChangesOperation.SetResultsLimit]
//
// # Processing the Operation’s Results
//
//   - [CKFetchDatabaseChangesOperation.RecordZoneWithIDChangedBlock]: The closure to execute with a single record zone change.
//   - [CKFetchDatabaseChangesOperation.SetRecordZoneWithIDChangedBlock]
//   - [CKFetchDatabaseChangesOperation.RecordZoneWithIDWasDeletedBlock]: The closure to execute when a record zone no longer exists.
//   - [CKFetchDatabaseChangesOperation.SetRecordZoneWithIDWasDeletedBlock]
//   - [CKFetchDatabaseChangesOperation.RecordZoneWithIDWasDeletedDueToUserEncryptedDataResetBlock]: The closure to execute when a user-invoked account reset deletes a record zone.
//   - [CKFetchDatabaseChangesOperation.SetRecordZoneWithIDWasDeletedDueToUserEncryptedDataResetBlock]
//   - [CKFetchDatabaseChangesOperation.RecordZoneWithIDWasPurgedBlock]: The closure to execute when CloudKit purges a record zone.
//   - [CKFetchDatabaseChangesOperation.SetRecordZoneWithIDWasPurgedBlock]
//   - [CKFetchDatabaseChangesOperation.ChangeTokenUpdatedBlock]: The closure to execute when the change token updates.
//   - [CKFetchDatabaseChangesOperation.SetChangeTokenUpdatedBlock]
//
// # Instance Properties
//
//   - [CKFetchDatabaseChangesOperation.FetchDatabaseChangesResultBlock]: The closure to execute when the operation finishes.
//   - [CKFetchDatabaseChangesOperation.SetFetchDatabaseChangesResultBlock]
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchDatabaseChangesOperation
//
// [NSSecureCoding]: https://developer.apple.com/documentation/Foundation/NSSecureCoding
type CKFetchDatabaseChangesOperation struct {
	CKDatabaseOperation
}

// CKFetchDatabaseChangesOperationFromID constructs a [CKFetchDatabaseChangesOperation] from an objc.ID.
//
// An operation that fetches database changes.
func CKFetchDatabaseChangesOperationFromID(id objc.ID) CKFetchDatabaseChangesOperation {
	return CKFetchDatabaseChangesOperation{CKDatabaseOperation: CKDatabaseOperationFromID(id)}
}

// NOTE: CKFetchDatabaseChangesOperation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKFetchDatabaseChangesOperation] class.
//
// # Creating an Operation
//
//   - [ICKFetchDatabaseChangesOperation.InitWithPreviousServerChangeToken]: Creates an operation for fetching database changes.
//
// # Configuring the Operation
//
//   - [ICKFetchDatabaseChangesOperation.FetchAllChanges]: A Boolean value that indicates whether to send repeated requests to the server.
//   - [ICKFetchDatabaseChangesOperation.SetFetchAllChanges]
//   - [ICKFetchDatabaseChangesOperation.PreviousServerChangeToken]: The server change token.
//   - [ICKFetchDatabaseChangesOperation.SetPreviousServerChangeToken]
//   - [ICKFetchDatabaseChangesOperation.ResultsLimit]: The maximum number of results that the operation fetches.
//   - [ICKFetchDatabaseChangesOperation.SetResultsLimit]
//
// # Processing the Operation’s Results
//
//   - [ICKFetchDatabaseChangesOperation.RecordZoneWithIDChangedBlock]: The closure to execute with a single record zone change.
//   - [ICKFetchDatabaseChangesOperation.SetRecordZoneWithIDChangedBlock]
//   - [ICKFetchDatabaseChangesOperation.RecordZoneWithIDWasDeletedBlock]: The closure to execute when a record zone no longer exists.
//   - [ICKFetchDatabaseChangesOperation.SetRecordZoneWithIDWasDeletedBlock]
//   - [ICKFetchDatabaseChangesOperation.RecordZoneWithIDWasDeletedDueToUserEncryptedDataResetBlock]: The closure to execute when a user-invoked account reset deletes a record zone.
//   - [ICKFetchDatabaseChangesOperation.SetRecordZoneWithIDWasDeletedDueToUserEncryptedDataResetBlock]
//   - [ICKFetchDatabaseChangesOperation.RecordZoneWithIDWasPurgedBlock]: The closure to execute when CloudKit purges a record zone.
//   - [ICKFetchDatabaseChangesOperation.SetRecordZoneWithIDWasPurgedBlock]
//   - [ICKFetchDatabaseChangesOperation.ChangeTokenUpdatedBlock]: The closure to execute when the change token updates.
//   - [ICKFetchDatabaseChangesOperation.SetChangeTokenUpdatedBlock]
//
// # Instance Properties
//
//   - [ICKFetchDatabaseChangesOperation.FetchDatabaseChangesResultBlock]: The closure to execute when the operation finishes.
//   - [ICKFetchDatabaseChangesOperation.SetFetchDatabaseChangesResultBlock]
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchDatabaseChangesOperation
type ICKFetchDatabaseChangesOperation interface {
	ICKDatabaseOperation

	// Topic: Creating an Operation

	// Creates an operation for fetching database changes.
	InitWithPreviousServerChangeToken(previousServerChangeToken ICKServerChangeToken) CKFetchDatabaseChangesOperation

	// Topic: Configuring the Operation

	// A Boolean value that indicates whether to send repeated requests to the server.
	FetchAllChanges() bool
	SetFetchAllChanges(value bool)
	// The server change token.
	PreviousServerChangeToken() ICKServerChangeToken
	SetPreviousServerChangeToken(value ICKServerChangeToken)
	// The maximum number of results that the operation fetches.
	ResultsLimit() uint
	SetResultsLimit(value uint)

	// Topic: Processing the Operation’s Results

	// The closure to execute with a single record zone change.
	RecordZoneWithIDChangedBlock() CKRecordZoneIDHandler
	SetRecordZoneWithIDChangedBlock(value CKRecordZoneIDHandler)
	// The closure to execute when a record zone no longer exists.
	RecordZoneWithIDWasDeletedBlock() CKRecordZoneIDHandler
	SetRecordZoneWithIDWasDeletedBlock(value CKRecordZoneIDHandler)
	// The closure to execute when a user-invoked account reset deletes a record zone.
	RecordZoneWithIDWasDeletedDueToUserEncryptedDataResetBlock() CKRecordZoneIDHandler
	SetRecordZoneWithIDWasDeletedDueToUserEncryptedDataResetBlock(value CKRecordZoneIDHandler)
	// The closure to execute when CloudKit purges a record zone.
	RecordZoneWithIDWasPurgedBlock() CKRecordZoneIDHandler
	SetRecordZoneWithIDWasPurgedBlock(value CKRecordZoneIDHandler)
	// The closure to execute when the change token updates.
	ChangeTokenUpdatedBlock() CKServerChangeTokenHandler
	SetChangeTokenUpdatedBlock(value CKServerChangeTokenHandler)

	// Topic: Instance Properties

	// The closure to execute when the operation finishes.
	FetchDatabaseChangesResultBlock() unsafe.Pointer
	SetFetchDatabaseChangesResultBlock(value unsafe.Pointer)
}

// Init initializes the instance.
func (c CKFetchDatabaseChangesOperation) Init() CKFetchDatabaseChangesOperation {
	rv := objc.Send[CKFetchDatabaseChangesOperation](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKFetchDatabaseChangesOperation) Autorelease() CKFetchDatabaseChangesOperation {
	rv := objc.Send[CKFetchDatabaseChangesOperation](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKFetchDatabaseChangesOperation creates a new CKFetchDatabaseChangesOperation instance.
func NewCKFetchDatabaseChangesOperation() CKFetchDatabaseChangesOperation {
	class := getCKFetchDatabaseChangesOperationClass()
	rv := objc.Send[CKFetchDatabaseChangesOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an operation for fetching database changes.
//
// previousServerChangeToken: The change token that CloudKit uses to determine which database changes to
// return.
//
// # Discussion
//
// After creating the operation, assign a handler to the
// [fetchDatabaseChangesCompletionBlock] property so that you can process the
// operation’s results.
//
// If this is your first fetch, or if you want to refetch all zones, pass
// `nil` for the change token. If you provide a change token from a previous
// [CKFetchDatabaseChangesOperation], CloudKit returns only the zones with
// changes since that token. The per-database [CKServerChangeToken] isn’t
// the same as the per-record zone [CKServerChangeToken] from
// [CKFetchRecordZoneChangesOperation].
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchDatabaseChangesOperation/init(previousServerChangeToken:)
//
// [fetchDatabaseChangesCompletionBlock]: https://developer.apple.com/documentation/CloudKit/CKFetchDatabaseChangesOperation/fetchDatabaseChangesCompletionBlock
func NewCKFetchDatabaseChangesOperationWithPreviousServerChangeToken(previousServerChangeToken ICKServerChangeToken) CKFetchDatabaseChangesOperation {
	instance := getCKFetchDatabaseChangesOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPreviousServerChangeToken:"), previousServerChangeToken)
	return CKFetchDatabaseChangesOperationFromID(rv)
}

// Creates an operation for fetching database changes.
//
// previousServerChangeToken: The change token that CloudKit uses to determine which database changes to
// return.
//
// # Discussion
//
// After creating the operation, assign a handler to the
// [fetchDatabaseChangesCompletionBlock] property so that you can process the
// operation’s results.
//
// If this is your first fetch, or if you want to refetch all zones, pass
// `nil` for the change token. If you provide a change token from a previous
// [CKFetchDatabaseChangesOperation], CloudKit returns only the zones with
// changes since that token. The per-database [CKServerChangeToken] isn’t
// the same as the per-record zone [CKServerChangeToken] from
// [CKFetchRecordZoneChangesOperation].
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchDatabaseChangesOperation/init(previousServerChangeToken:)
//
// [fetchDatabaseChangesCompletionBlock]: https://developer.apple.com/documentation/CloudKit/CKFetchDatabaseChangesOperation/fetchDatabaseChangesCompletionBlock
func (c CKFetchDatabaseChangesOperation) InitWithPreviousServerChangeToken(previousServerChangeToken ICKServerChangeToken) CKFetchDatabaseChangesOperation {
	rv := objc.Send[CKFetchDatabaseChangesOperation](c.ID, objc.Sel("initWithPreviousServerChangeToken:"), previousServerChangeToken)
	return rv
}

// A Boolean value that indicates whether to send repeated requests to the
// server.
//
// # Discussion
//
// If true, the operation sends repeat requests to the server until it fetches
// all changes. CloudKit executes the handler you set on the
// [ChangeTokenUpdatedBlock] property with a change token after each request.
//
// The default value is true.
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchDatabaseChangesOperation/fetchAllChanges
func (c CKFetchDatabaseChangesOperation) FetchAllChanges() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("fetchAllChanges"))
	return rv
}
func (c CKFetchDatabaseChangesOperation) SetFetchAllChanges(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setFetchAllChanges:"), value)
}

// The server change token.
//
// # Discussion
//
// Assign the token you receive from the [fetchDatabaseChangesCompletionBlock]
// to this property. Doing so yields only the changes that occur after your
// most recent fetch operation. If you specify `nil` for this parameter, the
// operation fetches all changes.
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchDatabaseChangesOperation/previousServerChangeToken
//
// [fetchDatabaseChangesCompletionBlock]: https://developer.apple.com/documentation/CloudKit/CKFetchDatabaseChangesOperation/fetchDatabaseChangesCompletionBlock
func (c CKFetchDatabaseChangesOperation) PreviousServerChangeToken() ICKServerChangeToken {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("previousServerChangeToken"))
	return CKServerChangeTokenFromID(objc.ID(rv))
}
func (c CKFetchDatabaseChangesOperation) SetPreviousServerChangeToken(value ICKServerChangeToken) {
	objc.Send[struct{}](c.ID, objc.Sel("setPreviousServerChangeToken:"), value)
}

// The maximum number of results that the operation fetches.
//
// # Discussion
//
// Use this property to limit the number of changes this operation returns.
// When the operation reaches the limit, it updates the change token and
// returns it to indicate that more results are available.
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchDatabaseChangesOperation/resultsLimit
func (c CKFetchDatabaseChangesOperation) ResultsLimit() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("resultsLimit"))
	return rv
}
func (c CKFetchDatabaseChangesOperation) SetResultsLimit(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setResultsLimit:"), value)
}

// The closure to execute with a single record zone change.
//
// # Discussion
//
// The closure returns no value and takes the following parameter:
//
// `zoneID`: The ID of the record zone that contains changes.
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchDatabaseChangesOperation/recordZoneWithIDChangedBlock
func (c CKFetchDatabaseChangesOperation) RecordZoneWithIDChangedBlock() CKRecordZoneIDHandler {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("recordZoneWithIDChangedBlock"))
	_ = rv
	return nil
}
func (c CKFetchDatabaseChangesOperation) SetRecordZoneWithIDChangedBlock(value CKRecordZoneIDHandler) {
	block, cleanup := NewCKRecordZoneIDBlock(value)
	defer cleanup()
	objc.Send[struct{}](c.ID, objc.Sel("setRecordZoneWithIDChangedBlock:"), block)
}

// The closure to execute when a record zone no longer exists.
//
// # Discussion
//
// The closure returns no value and takes the following parameter:
//
// `zoneID`: The deleted record zone’s ID.
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchDatabaseChangesOperation/recordZoneWithIDWasDeletedBlock
func (c CKFetchDatabaseChangesOperation) RecordZoneWithIDWasDeletedBlock() CKRecordZoneIDHandler {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("recordZoneWithIDWasDeletedBlock"))
	_ = rv
	return nil
}
func (c CKFetchDatabaseChangesOperation) SetRecordZoneWithIDWasDeletedBlock(value CKRecordZoneIDHandler) {
	block, cleanup := NewCKRecordZoneIDBlock(value)
	defer cleanup()
	objc.Send[struct{}](c.ID, objc.Sel("setRecordZoneWithIDWasDeletedBlock:"), block)
}

// The closure to execute when a user-invoked account reset deletes a record
// zone.
//
// # Discussion
//
// The closure returns no value and takes a single parameter: the deleted
// record zone’s ID.
//
// The operation executes this closure, instead of
// [RecordZoneWithIDWasDeletedBlock], after a user action causes CloudKit to
// delete the record zone. Reupload any locally cached data to iCloud to
// minimize data loss.
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchDatabaseChangesOperation/recordZoneWithIDWasDeletedDueToUserEncryptedDataResetBlock
func (c CKFetchDatabaseChangesOperation) RecordZoneWithIDWasDeletedDueToUserEncryptedDataResetBlock() CKRecordZoneIDHandler {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("recordZoneWithIDWasDeletedDueToUserEncryptedDataResetBlock"))
	_ = rv
	return nil
}
func (c CKFetchDatabaseChangesOperation) SetRecordZoneWithIDWasDeletedDueToUserEncryptedDataResetBlock(value CKRecordZoneIDHandler) {
	block, cleanup := NewCKRecordZoneIDBlock(value)
	defer cleanup()
	objc.Send[struct{}](c.ID, objc.Sel("setRecordZoneWithIDWasDeletedDueToUserEncryptedDataResetBlock:"), block)
}

// The closure to execute when CloudKit purges a record zone.
//
// # Discussion
//
// The closure returns no value and takes the following parameter:
//
// `zoneID`: The purged record zone’s ID.
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchDatabaseChangesOperation/recordZoneWithIDWasPurgedBlock
func (c CKFetchDatabaseChangesOperation) RecordZoneWithIDWasPurgedBlock() CKRecordZoneIDHandler {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("recordZoneWithIDWasPurgedBlock"))
	_ = rv
	return nil
}
func (c CKFetchDatabaseChangesOperation) SetRecordZoneWithIDWasPurgedBlock(value CKRecordZoneIDHandler) {
	block, cleanup := NewCKRecordZoneIDBlock(value)
	defer cleanup()
	objc.Send[struct{}](c.ID, objc.Sel("setRecordZoneWithIDWasPurgedBlock:"), block)
}

// The closure to execute when the change token updates.
//
// # Discussion
//
// The closure executes periodically, and provides a new change token so that
// you don’t need to refetch previously fetched record zone changes in a
// subsequent operation.
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchDatabaseChangesOperation/changeTokenUpdatedBlock
func (c CKFetchDatabaseChangesOperation) ChangeTokenUpdatedBlock() CKServerChangeTokenHandler {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("changeTokenUpdatedBlock"))
	_ = rv
	return nil
}
func (c CKFetchDatabaseChangesOperation) SetChangeTokenUpdatedBlock(value CKServerChangeTokenHandler) {
	block, cleanup := NewCKServerChangeTokenBlock(value)
	defer cleanup()
	objc.Send[struct{}](c.ID, objc.Sel("setChangeTokenUpdatedBlock:"), block)
}

// The closure to execute when the operation finishes.
//
// See: https://developer.apple.com/documentation/cloudkit/ckfetchdatabasechangesoperation/fetchdatabasechangesresultblock
func (c CKFetchDatabaseChangesOperation) FetchDatabaseChangesResultBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("fetchDatabaseChangesResultBlock"))
	return rv
}
func (c CKFetchDatabaseChangesOperation) SetFetchDatabaseChangesResultBlock(value unsafe.Pointer) {
	objc.Send[struct{}](c.ID, objc.Sel("setFetchDatabaseChangesResultBlock:"), value)
}
