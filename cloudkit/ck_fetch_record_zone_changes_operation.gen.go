// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/kernel"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKFetchRecordZoneChangesOperation] class.
var (
	_CKFetchRecordZoneChangesOperationClass     CKFetchRecordZoneChangesOperationClass
	_CKFetchRecordZoneChangesOperationClassOnce sync.Once
)

func getCKFetchRecordZoneChangesOperationClass() CKFetchRecordZoneChangesOperationClass {
	_CKFetchRecordZoneChangesOperationClassOnce.Do(func() {
		_CKFetchRecordZoneChangesOperationClass = CKFetchRecordZoneChangesOperationClass{class: objc.GetClass("CKFetchRecordZoneChangesOperation")}
	})
	return _CKFetchRecordZoneChangesOperationClass
}

// GetCKFetchRecordZoneChangesOperationClass returns the class object for CKFetchRecordZoneChangesOperation.
func GetCKFetchRecordZoneChangesOperationClass() CKFetchRecordZoneChangesOperationClass {
	return getCKFetchRecordZoneChangesOperationClass()
}

type CKFetchRecordZoneChangesOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKFetchRecordZoneChangesOperationClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKFetchRecordZoneChangesOperationClass) Alloc() CKFetchRecordZoneChangesOperation {
	rv := objc.Send[CKFetchRecordZoneChangesOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An operation that fetches record zone changes.
//
// # Overview
//
// Use this operation to fetch record changes in one or more record zones,
// such as those that occur during record creation, modification, and
// deletion. You provide a configuration object for each record zone to query
// for changes. The configuration contains a server change token, which is an
// opaque pointer to a specific change in the zone’s history. CloudKit
// returns only the changes that occur after that point. For the first time
// you fetch a record zone’s changes, or to refetch all changes in a
// zone’s history, use `nil` instead.
//
// CloudKit processes the record zones in succession, and returns the changes
// for each zone in batches. Each batch yields a new change token. If all
// batches return without error, the operation issues a final change token for
// that zone. The change tokens conform to [NSSecureCoding] and are safe to
// cache on-disk. This operation’s tokens aren’t compatible with
// [CKFetchDatabaseChangesOperation] so you should segregate them in your
// app’s cache. Don’t infer behavior or order from the tokens’ contents.
//
// If you create record zones in the private database, fetch all changes the
// first time the app launches. Cache the results on-device and use
// [CKRecordZoneSubscription] to subscribe to future changes. Fetch those
// changes on receipt of the push notifications the subscription generates. If
// you use the shared database, subscribe to changes with
// [CKDatabaseSubscription] instead. When a user participates in sharing,
// CloudKit adds and removes record zones. This means you don’t know in
// advance which zones exist in the shared database. Use
// [CKFetchDatabaseChangesOperation] to fetch shared record zones on receipt
// of the subscription’s push notifications. Then fetch the changes in those
// zones using this operation. Regardless of which database you use, it’s
// not necessary to perform fetches each time your app launches, or to
// schedule fetches at regular intervals.
//
// To run the operation, add it to the corresponding database’s operation
// queue. The operation executes its callbacks on a private serial queue.
//
// The following example demonstrates how to create the operation, configure
// its callbacks, and execute it. For brevity, it omits the delete and
// operation completion callbacks.
//
// # Configuring the Zone Change Operation
//
//   - [CKFetchRecordZoneChangesOperation.ConfigurationsByRecordZoneID]: A dictionary of configurations for fetching change operations by zone identifier.
//   - [CKFetchRecordZoneChangesOperation.SetConfigurationsByRecordZoneID]
//   - [CKFetchRecordZoneChangesOperation.FetchAllChanges]: A Boolean value that indicates whether to send repeated requests to the server.
//   - [CKFetchRecordZoneChangesOperation.SetFetchAllChanges]
//   - [CKFetchRecordZoneChangesOperation.RecordZoneIDs]: The IDs of the record zones that contain the records to fetch.
//   - [CKFetchRecordZoneChangesOperation.SetRecordZoneIDs]
//
// # Processing the Zone Change Operation Results
//
//   - [CKFetchRecordZoneChangesOperation.RecordWithIDWasDeletedBlock]: The closure to execute when a record no longer exists.
//   - [CKFetchRecordZoneChangesOperation.SetRecordWithIDWasDeletedBlock]
//   - [CKFetchRecordZoneChangesOperation.RecordZoneChangeTokensUpdatedBlock]: The closure to execute when the change token updates.
//   - [CKFetchRecordZoneChangesOperation.SetRecordZoneChangeTokensUpdatedBlock]
//
// # Instance Properties
//
//   - [CKFetchRecordZoneChangesOperation.RecordWasChangedBlock]: The closure to execute with the results of retrieving a record change.
//   - [CKFetchRecordZoneChangesOperation.SetRecordWasChangedBlock]
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation
//
// [NSSecureCoding]: https://developer.apple.com/documentation/Foundation/NSSecureCoding
type CKFetchRecordZoneChangesOperation struct {
	CKDatabaseOperation
}

// CKFetchRecordZoneChangesOperationFromID constructs a [CKFetchRecordZoneChangesOperation] from an objc.ID.
//
// An operation that fetches record zone changes.
func CKFetchRecordZoneChangesOperationFromID(id objc.ID) CKFetchRecordZoneChangesOperation {
	return CKFetchRecordZoneChangesOperation{CKDatabaseOperation: CKDatabaseOperationFromID(id)}
}

// NOTE: CKFetchRecordZoneChangesOperation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKFetchRecordZoneChangesOperation] class.
//
// # Configuring the Zone Change Operation
//
//   - [ICKFetchRecordZoneChangesOperation.ConfigurationsByRecordZoneID]: A dictionary of configurations for fetching change operations by zone identifier.
//   - [ICKFetchRecordZoneChangesOperation.SetConfigurationsByRecordZoneID]
//   - [ICKFetchRecordZoneChangesOperation.FetchAllChanges]: A Boolean value that indicates whether to send repeated requests to the server.
//   - [ICKFetchRecordZoneChangesOperation.SetFetchAllChanges]
//   - [ICKFetchRecordZoneChangesOperation.RecordZoneIDs]: The IDs of the record zones that contain the records to fetch.
//   - [ICKFetchRecordZoneChangesOperation.SetRecordZoneIDs]
//
// # Processing the Zone Change Operation Results
//
//   - [ICKFetchRecordZoneChangesOperation.RecordWithIDWasDeletedBlock]: The closure to execute when a record no longer exists.
//   - [ICKFetchRecordZoneChangesOperation.SetRecordWithIDWasDeletedBlock]
//   - [ICKFetchRecordZoneChangesOperation.RecordZoneChangeTokensUpdatedBlock]: The closure to execute when the change token updates.
//   - [ICKFetchRecordZoneChangesOperation.SetRecordZoneChangeTokensUpdatedBlock]
//
// # Instance Properties
//
//   - [ICKFetchRecordZoneChangesOperation.RecordWasChangedBlock]: The closure to execute with the results of retrieving a record change.
//   - [ICKFetchRecordZoneChangesOperation.SetRecordWasChangedBlock]
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation
type ICKFetchRecordZoneChangesOperation interface {
	ICKDatabaseOperation

	// Topic: Configuring the Zone Change Operation

	// A dictionary of configurations for fetching change operations by zone identifier.
	ConfigurationsByRecordZoneID() foundation.INSDictionary
	SetConfigurationsByRecordZoneID(value foundation.INSDictionary)
	// A Boolean value that indicates whether to send repeated requests to the server.
	FetchAllChanges() bool
	SetFetchAllChanges(value bool)
	// The IDs of the record zones that contain the records to fetch.
	RecordZoneIDs() []CKRecordZoneID
	SetRecordZoneIDs(value []CKRecordZoneID)

	// Topic: Processing the Zone Change Operation Results

	// The closure to execute when a record no longer exists.
	RecordWithIDWasDeletedBlock() unsafe.Pointer
	SetRecordWithIDWasDeletedBlock(value kernel.Pointer)
	// The closure to execute when the change token updates.
	RecordZoneChangeTokensUpdatedBlock() CKRecordZoneIDCKServerChangeTokenDataHandler
	SetRecordZoneChangeTokensUpdatedBlock(value CKRecordZoneIDCKServerChangeTokenDataHandler)

	// Topic: Instance Properties

	// The closure to execute with the results of retrieving a record change.
	RecordWasChangedBlock() unsafe.Pointer
	SetRecordWasChangedBlock(value kernel.Pointer)

	// Configuration options for each record zone that the operation retrieves.
	OptionsByRecordZoneID() unsafe.Pointer
	SetOptionsByRecordZoneID(value kernel.Pointer)
}

// Init initializes the instance.
func (c CKFetchRecordZoneChangesOperation) Init() CKFetchRecordZoneChangesOperation {
	rv := objc.Send[CKFetchRecordZoneChangesOperation](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKFetchRecordZoneChangesOperation) Autorelease() CKFetchRecordZoneChangesOperation {
	rv := objc.Send[CKFetchRecordZoneChangesOperation](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKFetchRecordZoneChangesOperation creates a new CKFetchRecordZoneChangesOperation instance.
func NewCKFetchRecordZoneChangesOperation() CKFetchRecordZoneChangesOperation {
	class := getCKFetchRecordZoneChangesOperationClass()
	rv := objc.Send[CKFetchRecordZoneChangesOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A dictionary of configurations for fetching change operations by zone
// identifier.
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/configurationsByRecordZoneID
func (c CKFetchRecordZoneChangesOperation) ConfigurationsByRecordZoneID() foundation.INSDictionary {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("configurationsByRecordZoneID"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (c CKFetchRecordZoneChangesOperation) SetConfigurationsByRecordZoneID(value foundation.INSDictionary) {
	objc.Send[struct{}](c.ID, objc.Sel("setConfigurationsByRecordZoneID:"), value)
}

// A Boolean value that indicates whether to send repeated requests to the
// server.
//
// # Discussion
//
// If true, the operation sends repeat requests to the server until it fetches
// all changes. CloudKit executes the handler you set on the
// [recordZoneFetchCompletionBlock] property with a change token after each
// request.
//
// The default value is true.
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/fetchAllChanges
//
// [recordZoneFetchCompletionBlock]: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/recordZoneFetchCompletionBlock
func (c CKFetchRecordZoneChangesOperation) FetchAllChanges() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("fetchAllChanges"))
	return rv
}
func (c CKFetchRecordZoneChangesOperation) SetFetchAllChanges(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setFetchAllChanges:"), value)
}

// The IDs of the record zones that contain the records to fetch.
//
// # Discussion
//
// Typically, you set the value of this property when you create the
// operation. If you intend to change the record zone IDs, update the value
// before you execute the operation or submit it to a queue.
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/recordZoneIDs
func (c CKFetchRecordZoneChangesOperation) RecordZoneIDs() []CKRecordZoneID {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("recordZoneIDs"))
	return objc.ConvertSlice(rv, func(id objc.ID) CKRecordZoneID {
		return CKRecordZoneIDFromID(id)
	})
}
func (c CKFetchRecordZoneChangesOperation) SetRecordZoneIDs(value []CKRecordZoneID) {
	objc.Send[struct{}](c.ID, objc.Sel("setRecordZoneIDs:"), objectivec.IObjectSliceToNSArray(value))
}

// The closure to execute when a record no longer exists.
//
// See: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonechangesoperation/recordwithidwasdeletedblock-3z14c
func (c CKFetchRecordZoneChangesOperation) RecordWithIDWasDeletedBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("recordWithIDWasDeletedBlock"))
	return rv
}
func (c CKFetchRecordZoneChangesOperation) SetRecordWithIDWasDeletedBlock(value kernel.Pointer) {
	objc.Send[struct{}](c.ID, objc.Sel("setRecordWithIDWasDeletedBlock:"), value)
}

// The closure to execute when the change token updates.
//
// # Discussion
//
// The closure returns no value and takes the following parameters:
//
// - The record zone’s ID. - The new change token from the server. You can
// store this token locally and use it during subsequent fetch operations to
// limit the results to records that change after this operation executes. -
// The most recent client change token from the device. If the change token
// isn’t the most recent change token you provided, the server might not
// have received the associated changes.
//
// The operation executes this closure once for each record zone. Each time
// the closure executes, it executes serially with respect to the other blocks
// of the operation.
//
// Set this property before you execute the operation or submit it to a queue.
//
// See: https://developer.apple.com/documentation/CloudKit/CKFetchRecordZoneChangesOperation/recordZoneChangeTokensUpdatedBlock
func (c CKFetchRecordZoneChangesOperation) RecordZoneChangeTokensUpdatedBlock() CKRecordZoneIDCKServerChangeTokenDataHandler {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("recordZoneChangeTokensUpdatedBlock"))
	_ = rv
	return nil
}
func (c CKFetchRecordZoneChangesOperation) SetRecordZoneChangeTokensUpdatedBlock(value CKRecordZoneIDCKServerChangeTokenDataHandler) {
	block, cleanup := NewCKRecordZoneIDCKServerChangeTokenDataBlock(value)
	defer cleanup()
	objc.Send[struct{}](c.ID, objc.Sel("setRecordZoneChangeTokensUpdatedBlock:"), block)
}

// The closure to execute with the results of retrieving a record change.
//
// See: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonechangesoperation/recordwaschangedblock-x5bw
func (c CKFetchRecordZoneChangesOperation) RecordWasChangedBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("recordWasChangedBlock"))
	return rv
}
func (c CKFetchRecordZoneChangesOperation) SetRecordWasChangedBlock(value kernel.Pointer) {
	objc.Send[struct{}](c.ID, objc.Sel("setRecordWasChangedBlock:"), value)
}

// Configuration options for each record zone that the operation retrieves.
//
// See: https://developer.apple.com/documentation/cloudkit/ckfetchrecordzonechangesoperation/optionsbyrecordzoneid
func (c CKFetchRecordZoneChangesOperation) OptionsByRecordZoneID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("optionsByRecordZoneID"))
	return rv
}
func (c CKFetchRecordZoneChangesOperation) SetOptionsByRecordZoneID(value kernel.Pointer) {
	objc.Send[struct{}](c.ID, objc.Sel("setOptionsByRecordZoneID:"), value)
}
