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

// The class instance for the [CKModifyRecordsOperation] class.
var (
	_CKModifyRecordsOperationClass     CKModifyRecordsOperationClass
	_CKModifyRecordsOperationClassOnce sync.Once
)

func getCKModifyRecordsOperationClass() CKModifyRecordsOperationClass {
	_CKModifyRecordsOperationClassOnce.Do(func() {
		_CKModifyRecordsOperationClass = CKModifyRecordsOperationClass{class: objc.GetClass("CKModifyRecordsOperation")}
	})
	return _CKModifyRecordsOperationClass
}

// GetCKModifyRecordsOperationClass returns the class object for CKModifyRecordsOperation.
func GetCKModifyRecordsOperationClass() CKModifyRecordsOperationClass {
	return getCKModifyRecordsOperationClass()
}

type CKModifyRecordsOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKModifyRecordsOperationClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKModifyRecordsOperationClass) Alloc() CKModifyRecordsOperation {
	rv := objc.Send[CKModifyRecordsOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An operation that modifies one or more records.
//
// # Overview
//
// After modifying the fields of a record, use this operation to save those
// changes to a database. You also use this operation to delete records
// permanently from a database.
//
// If you’re saving a record that contains a reference to another record,
// set the reference’s [CKReference.ReferenceAction] to indicate if the
// target record’s deletion should cascade to the saved record. This helps
// avoid orphaned records in explicit record hierarchies. When creating two
// new records that have a reference between them, use the same operation to
// save both records at the same time. During a save operation, CloudKit
// requires that the target record of the [CKRecord.Parent] reference, if set,
// exists in the database or is part of the same operation; all other
// reference fields are exempt from this requirement.
//
// When you save records, the value in the
// [CKModifyRecordsOperation.SavePolicy] property determines how to proceed
// when CloudKit detects conflicts. Because records can change between the
// time you fetch them and the time you save them, the save policy determines
// whether new changes overwrite existing changes. By default, the operation
// reports an error when there’s a newer version on the server. You can
// change the default setting to permit your changes to overwrite the server
// values wholly or partially.
//
// The handlers you assign to monitor progress of the operation execute
// serially on an internal queue that the operation manages. You must provide
// handlers capable of executing on a background thread, so any tasks that
// require access to the main thread must redirect accordingly.
//
// If you assign a completion handler to the [completionBlock] property of the
// operation, CloudKit calls it after the operation executes and returns the
// results. Use the completion handler to perform any housekeeping tasks for
// the operation, but don’t use it to process the results of the operation.
// The completion handler you provide should manage any failures of the
// operation, whether due to an error or an explicit cancellation.
//
// # Configuring the Modify Record Operation
//
//   - [CKModifyRecordsOperation.RecordsToSave]: The records to save to the database.
//   - [CKModifyRecordsOperation.SetRecordsToSave]
//   - [CKModifyRecordsOperation.RecordIDsToDelete]: The IDs of the records to delete permanently from the database.
//   - [CKModifyRecordsOperation.SetRecordIDsToDelete]
//   - [CKModifyRecordsOperation.ClientChangeTokenData]: A token that tracks local changes to records.
//   - [CKModifyRecordsOperation.SetClientChangeTokenData]
//   - [CKModifyRecordsOperation.Atomic]: A Boolean value that indicates whether the entire operation fails when CloudKit can’t update one or more records in a record zone.
//   - [CKModifyRecordsOperation.SetAtomic]
//   - [CKModifyRecordsOperation.SavePolicy]: The policy to use when saving changes to records.
//   - [CKModifyRecordsOperation.SetSavePolicy]
//
// # Processing the Modify Record Results
//
//   - [CKModifyRecordsOperation.PerRecordProgressBlock]: The closure to execute with progress information for individual records.
//   - [CKModifyRecordsOperation.SetPerRecordProgressBlock]
//
// # Instance Properties
//
//   - [CKModifyRecordsOperation.PerRecordDeleteBlock]: The closure to execute when CloudKit deletes a record.
//   - [CKModifyRecordsOperation.SetPerRecordDeleteBlock]
//   - [CKModifyRecordsOperation.PerRecordSaveBlock]: The closure to execute when CloudKit saves a record.
//   - [CKModifyRecordsOperation.SetPerRecordSaveBlock]
//
// See: https://developer.apple.com/documentation/CloudKit/CKModifyRecordsOperation
//
// [completionBlock]: https://developer.apple.com/documentation/Foundation/Operation/completionBlock
type CKModifyRecordsOperation struct {
	CKDatabaseOperation
}

// CKModifyRecordsOperationFromID constructs a [CKModifyRecordsOperation] from an objc.ID.
//
// An operation that modifies one or more records.
func CKModifyRecordsOperationFromID(id objc.ID) CKModifyRecordsOperation {
	return CKModifyRecordsOperation{CKDatabaseOperation: CKDatabaseOperationFromID(id)}
}

// NOTE: CKModifyRecordsOperation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKModifyRecordsOperation] class.
//
// # Configuring the Modify Record Operation
//
//   - [ICKModifyRecordsOperation.RecordsToSave]: The records to save to the database.
//   - [ICKModifyRecordsOperation.SetRecordsToSave]
//   - [ICKModifyRecordsOperation.RecordIDsToDelete]: The IDs of the records to delete permanently from the database.
//   - [ICKModifyRecordsOperation.SetRecordIDsToDelete]
//   - [ICKModifyRecordsOperation.ClientChangeTokenData]: A token that tracks local changes to records.
//   - [ICKModifyRecordsOperation.SetClientChangeTokenData]
//   - [ICKModifyRecordsOperation.Atomic]: A Boolean value that indicates whether the entire operation fails when CloudKit can’t update one or more records in a record zone.
//   - [ICKModifyRecordsOperation.SetAtomic]
//   - [ICKModifyRecordsOperation.SavePolicy]: The policy to use when saving changes to records.
//   - [ICKModifyRecordsOperation.SetSavePolicy]
//
// # Processing the Modify Record Results
//
//   - [ICKModifyRecordsOperation.PerRecordProgressBlock]: The closure to execute with progress information for individual records.
//   - [ICKModifyRecordsOperation.SetPerRecordProgressBlock]
//
// # Instance Properties
//
//   - [ICKModifyRecordsOperation.PerRecordDeleteBlock]: The closure to execute when CloudKit deletes a record.
//   - [ICKModifyRecordsOperation.SetPerRecordDeleteBlock]
//   - [ICKModifyRecordsOperation.PerRecordSaveBlock]: The closure to execute when CloudKit saves a record.
//   - [ICKModifyRecordsOperation.SetPerRecordSaveBlock]
//
// See: https://developer.apple.com/documentation/CloudKit/CKModifyRecordsOperation
type ICKModifyRecordsOperation interface {
	ICKDatabaseOperation

	// Topic: Configuring the Modify Record Operation

	// The records to save to the database.
	RecordsToSave() []CKRecord
	SetRecordsToSave(value []CKRecord)
	// The IDs of the records to delete permanently from the database.
	RecordIDsToDelete() []CKRecordID
	SetRecordIDsToDelete(value []CKRecordID)
	// A token that tracks local changes to records.
	ClientChangeTokenData() foundation.NSData
	SetClientChangeTokenData(value foundation.NSData)
	// A Boolean value that indicates whether the entire operation fails when CloudKit can’t update one or more records in a record zone.
	Atomic() bool
	SetAtomic(value bool)
	// The policy to use when saving changes to records.
	SavePolicy() CKRecordSavePolicy
	SetSavePolicy(value CKRecordSavePolicy)

	// Topic: Processing the Modify Record Results

	// The closure to execute with progress information for individual records.
	PerRecordProgressBlock() CKRecordFloat64Handler
	SetPerRecordProgressBlock(value CKRecordFloat64Handler)

	// Topic: Instance Properties

	// The closure to execute when CloudKit deletes a record.
	PerRecordDeleteBlock() unsafe.Pointer
	SetPerRecordDeleteBlock(value kernel.Pointer)
	// The closure to execute when CloudKit saves a record.
	PerRecordSaveBlock() unsafe.Pointer
	SetPerRecordSaveBlock(value kernel.Pointer)
}

// Init initializes the instance.
func (c CKModifyRecordsOperation) Init() CKModifyRecordsOperation {
	rv := objc.Send[CKModifyRecordsOperation](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKModifyRecordsOperation) Autorelease() CKModifyRecordsOperation {
	rv := objc.Send[CKModifyRecordsOperation](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKModifyRecordsOperation creates a new CKModifyRecordsOperation instance.
func NewCKModifyRecordsOperation() CKModifyRecordsOperation {
	class := getCKModifyRecordsOperationClass()
	rv := objc.Send[CKModifyRecordsOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The records to save to the database.
//
// # Discussion
//
// The initial value of the property is the array that you provide to the
// [initWithRecordsToSave:recordIDsToDelete:] method. You can modify this
// array as necessary before you execute the operation. The records must all
// target the same database, but can belong to different record zones.
//
// If you intend to change the value of this property, do so before you
// execute the operation or submit the operation to a queue.
//
// See: https://developer.apple.com/documentation/CloudKit/CKModifyRecordsOperation/recordsToSave
//
// [initWithRecordsToSave:recordIDsToDelete:]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordsOperation/initWithRecordsToSave:recordIDsToDelete:
func (c CKModifyRecordsOperation) RecordsToSave() []CKRecord {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("recordsToSave"))
	return objc.ConvertSlice(rv, func(id objc.ID) CKRecord {
		return CKRecordFromID(id)
	})
}
func (c CKModifyRecordsOperation) SetRecordsToSave(value []CKRecord) {
	objc.Send[struct{}](c.ID, objc.Sel("setRecordsToSave:"), objectivec.IObjectSliceToNSArray(value))
}

// The IDs of the records to delete permanently from the database.
//
// # Discussion
//
// An array of [CKRecordID] objects that identifies the records to delete. The
// initial value of the property is the array of record IDs that you provide
// to the [initWithRecordsToSave:recordIDsToDelete:] method.
//
// When deleting records, the operation reports progress only on the records
// with the IDs that you specify in this property. Deleting records can
// trigger the deletion of related records if there is an owner-owned
// relationship between the records involving a [CKReference] object. When
// additional deletions occur, CloudKit doesn’t pass them to the progress
// handler of the operation. For that reason, it’s important to understand
// the implications of the ownership model you use when you relate records to
// each other through a [CKReference] object. For more information about
// owner-owned relationships, see [CKReference].
//
// If you intend to change the value of this property, do so before you
// execute the operation or submit the operation to a queue.
//
// See: https://developer.apple.com/documentation/CloudKit/CKModifyRecordsOperation/recordIDsToDelete
//
// [initWithRecordsToSave:recordIDsToDelete:]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordsOperation/initWithRecordsToSave:recordIDsToDelete:
func (c CKModifyRecordsOperation) RecordIDsToDelete() []CKRecordID {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("recordIDsToDelete"))
	return objc.ConvertSlice(rv, func(id objc.ID) CKRecordID {
		return CKRecordIDFromID(id)
	})
}
func (c CKModifyRecordsOperation) SetRecordIDsToDelete(value []CKRecordID) {
	objc.Send[struct{}](c.ID, objc.Sel("setRecordIDsToDelete:"), objectivec.IObjectSliceToNSArray(value))
}

// A token that tracks local changes to records.
//
// # Discussion
//
// The default value is `nil`.
//
// When you modify records from a fetch operation, specify a token using this
// property to indicate which version of the record you most recently
// modified. Compare the token you supply to the token in the next record
// fetch to confirm the server successfully receives the device’s most
// recent modify request.
//
// If you intend to change the value of this property, do so before you
// execute the operation or submit the operation to a queue.
//
// See: https://developer.apple.com/documentation/CloudKit/CKModifyRecordsOperation/clientChangeTokenData
func (c CKModifyRecordsOperation) ClientChangeTokenData() foundation.NSData {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("clientChangeTokenData"))
	return foundation.NSDataFromID(objc.ID(rv))
}
func (c CKModifyRecordsOperation) SetClientChangeTokenData(value foundation.NSData) {
	objc.Send[struct{}](c.ID, objc.Sel("setClientChangeTokenData:"), value)
}

// A Boolean value that indicates whether the entire operation fails when
// CloudKit can’t update one or more records in a record zone.
//
// # Discussion
//
// Modifying records atomically prevents you from updating your data in a way
// that would leave it in an inconsistent state. You use atomic updates when
// you want to write multiple records to the same record zone. If there’s a
// failure to modify any of the records in a zone, CloudKit doesn’t change
// the other records in that same zone. The record zone must have the
// [CKRecordZoneCapabilityAtomic] capability for this behavior to apply. If a
// record zone doesn’t support the atomic capability, setting this property
// has no effect.
//
// The default value of this property is true, which causes all modifications
// within a single record zone to occur atomically. If your operation contains
// records in multiple record zones, a failure in one zone doesn’t prevent
// modifications to records in a different zone. Changing the value of this
// property to false causes CloudKit to modify records individually,
// regardless of whether the record zone supports atomic modifications.
//
// See: https://developer.apple.com/documentation/CloudKit/CKModifyRecordsOperation/isAtomic
func (c CKModifyRecordsOperation) Atomic() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("atomic"))
	return rv
}
func (c CKModifyRecordsOperation) SetAtomic(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setAtomic:"), value)
}

// The policy to use when saving changes to records.
//
// # Discussion
//
// The server uses this property to determine how to proceed when saving
// record changes. The exact behavior depends on the policy you choose:
//
// - Use [CKModifyRecordsOperation.RecordSavePolicy.ifServerRecordUnchanged]
// to only save a record when the change tag of the local copy matches that of
// the server’s copy. If the server record’s change tag is more recent,
// CloudKit discards the save and returns a [CKError.Code.serverRecordChanged]
// error. - Use [CKModifyRecordsOperation.RecordSavePolicy.changedKeys] to
// save only the fields of the record that contain changes. The server
// doesn’t compare record change tags when using this policy. - Use
// [CKModifyRecordsOperation.RecordSavePolicy.allKeys] to save every field of
// the record, even those without changes. The server doesn’t compare record
// change tags when using this policy.
//
// If you change the property’s value, do so before you execute the
// operation or submit the operation to a queue. The default value is
// [CKModifyRecordsOperation.RecordSavePolicy.ifServerRecordUnchanged].
//
// See: https://developer.apple.com/documentation/CloudKit/CKModifyRecordsOperation/savePolicy
//
// [CKError.Code.serverRecordChanged]: https://developer.apple.com/documentation/CloudKit/CKError/Code/serverRecordChanged
// [CKModifyRecordsOperation.RecordSavePolicy.allKeys]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordsOperation/RecordSavePolicy/allKeys
// [CKModifyRecordsOperation.RecordSavePolicy.changedKeys]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordsOperation/RecordSavePolicy/changedKeys
// [CKModifyRecordsOperation.RecordSavePolicy.ifServerRecordUnchanged]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordsOperation/RecordSavePolicy/ifServerRecordUnchanged
func (c CKModifyRecordsOperation) SavePolicy() CKRecordSavePolicy {
	rv := objc.Send[CKRecordSavePolicy](c.ID, objc.Sel("savePolicy"))
	return CKRecordSavePolicy(rv)
}
func (c CKModifyRecordsOperation) SetSavePolicy(value CKRecordSavePolicy) {
	objc.Send[struct{}](c.ID, objc.Sel("setSavePolicy:"), value)
}

// The closure to execute with progress information for individual records.
//
// # Discussion
//
// This property is a closure that returns no value and has the following
// parameters:
//
// - The record that CloudKit saves. - The amount of data, as a percentage,
// that CloudKit saves for the record. The range is `0.0` to `1.0`, where
// `0.0` indicates that CloudKit hasn’t saved any data, and `1.0` means that
// CloudKit has saved the entire record.
//
// The modify records operation executes this closure one or more times for
// each record in the [CKModifyRecordsOperation.RecordsToSave] property. Each
// time the closure executes, it executes serially with respect to the other
// progress closures of the operation. You can use this closure to track the
// ongoing progress of the operation.
//
// If you intend to use this closure to process results, set it before you
// execute the operation or add the operation to a queue.
//
// See: https://developer.apple.com/documentation/CloudKit/CKModifyRecordsOperation/perRecordProgressBlock
func (c CKModifyRecordsOperation) PerRecordProgressBlock() CKRecordFloat64Handler {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("perRecordProgressBlock"))
	_ = rv
	return nil
}
func (c CKModifyRecordsOperation) SetPerRecordProgressBlock(value CKRecordFloat64Handler) {
	block, cleanup := NewCKRecordFloat64Block(value)
	defer cleanup()
	objc.Send[struct{}](c.ID, objc.Sel("setPerRecordProgressBlock:"), block)
}

// The closure to execute when CloudKit deletes a record.
//
// See: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordsoperation/perrecorddeleteblock-9czoo
func (c CKModifyRecordsOperation) PerRecordDeleteBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("perRecordDeleteBlock"))
	return rv
}
func (c CKModifyRecordsOperation) SetPerRecordDeleteBlock(value kernel.Pointer) {
	objc.Send[struct{}](c.ID, objc.Sel("setPerRecordDeleteBlock:"), value)
}

// The closure to execute when CloudKit saves a record.
//
// See: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordsoperation/perrecordsaveblock-7yq9d
func (c CKModifyRecordsOperation) PerRecordSaveBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("perRecordSaveBlock"))
	return rv
}
func (c CKModifyRecordsOperation) SetPerRecordSaveBlock(value kernel.Pointer) {
	objc.Send[struct{}](c.ID, objc.Sel("setPerRecordSaveBlock:"), value)
}
