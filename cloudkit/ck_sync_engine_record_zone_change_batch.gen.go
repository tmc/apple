// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKSyncEngineRecordZoneChangeBatch] class.
var (
	_CKSyncEngineRecordZoneChangeBatchClass     CKSyncEngineRecordZoneChangeBatchClass
	_CKSyncEngineRecordZoneChangeBatchClassOnce sync.Once
)

func getCKSyncEngineRecordZoneChangeBatchClass() CKSyncEngineRecordZoneChangeBatchClass {
	_CKSyncEngineRecordZoneChangeBatchClassOnce.Do(func() {
		_CKSyncEngineRecordZoneChangeBatchClass = CKSyncEngineRecordZoneChangeBatchClass{class: objc.GetClass("CKSyncEngineRecordZoneChangeBatch")}
	})
	return _CKSyncEngineRecordZoneChangeBatchClass
}

// GetCKSyncEngineRecordZoneChangeBatchClass returns the class object for CKSyncEngineRecordZoneChangeBatch.
func GetCKSyncEngineRecordZoneChangeBatchClass() CKSyncEngineRecordZoneChangeBatchClass {
	return getCKSyncEngineRecordZoneChangeBatchClass()
}

type CKSyncEngineRecordZoneChangeBatchClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKSyncEngineRecordZoneChangeBatchClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKSyncEngineRecordZoneChangeBatchClass) Alloc() CKSyncEngineRecordZoneChangeBatch {
	rv := objc.Send[CKSyncEngineRecordZoneChangeBatch](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object that contains the record changes for a single send operation.
//
// # Creating a batch
//
//   - [CKSyncEngineRecordZoneChangeBatch.InitWithPendingChangesRecordProvider]: Creates a batch of records to modify using the provided record zone changes.
//   - [CKSyncEngineRecordZoneChangeBatch.InitWithRecordsToSaveRecordIDsToDeleteAtomicByZone]: Creates a batch of records to modify.
//
// # Managing atomicity
//
//   - [CKSyncEngineRecordZoneChangeBatch.AtomicByZone]: A Boolean value that determines whether CloudKit modifies records atomically by record zone.
//   - [CKSyncEngineRecordZoneChangeBatch.SetAtomicByZone]
//
// # Managing the records
//
//   - [CKSyncEngineRecordZoneChangeBatch.RecordIDsToDelete]: The unique identifiers of the records to delete.
//   - [CKSyncEngineRecordZoneChangeBatch.RecordsToSave]: The records to save.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineRecordZoneChangeBatch
type CKSyncEngineRecordZoneChangeBatch struct {
	objectivec.Object
}

// CKSyncEngineRecordZoneChangeBatchFromID constructs a [CKSyncEngineRecordZoneChangeBatch] from an objc.ID.
//
// An object that contains the record changes for a single send operation.
func CKSyncEngineRecordZoneChangeBatchFromID(id objc.ID) CKSyncEngineRecordZoneChangeBatch {
	return CKSyncEngineRecordZoneChangeBatch{objectivec.Object{ID: id}}
}

// Ensure CKSyncEngineRecordZoneChangeBatch implements ICKSyncEngineRecordZoneChangeBatch.
var _ ICKSyncEngineRecordZoneChangeBatch = CKSyncEngineRecordZoneChangeBatch{}

// An interface definition for the [CKSyncEngineRecordZoneChangeBatch] class.
//
// # Creating a batch
//
//   - [ICKSyncEngineRecordZoneChangeBatch.InitWithPendingChangesRecordProvider]: Creates a batch of records to modify using the provided record zone changes.
//   - [ICKSyncEngineRecordZoneChangeBatch.InitWithRecordsToSaveRecordIDsToDeleteAtomicByZone]: Creates a batch of records to modify.
//
// # Managing atomicity
//
//   - [ICKSyncEngineRecordZoneChangeBatch.AtomicByZone]: A Boolean value that determines whether CloudKit modifies records atomically by record zone.
//   - [ICKSyncEngineRecordZoneChangeBatch.SetAtomicByZone]
//
// # Managing the records
//
//   - [ICKSyncEngineRecordZoneChangeBatch.RecordIDsToDelete]: The unique identifiers of the records to delete.
//   - [ICKSyncEngineRecordZoneChangeBatch.RecordsToSave]: The records to save.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineRecordZoneChangeBatch
type ICKSyncEngineRecordZoneChangeBatch interface {
	objectivec.IObject

	// Topic: Creating a batch

	// Creates a batch of records to modify using the provided record zone changes.
	InitWithPendingChangesRecordProvider(pendingChanges []CKSyncEnginePendingRecordZoneChange, recordProvider CKRecordIDHandler) CKSyncEngineRecordZoneChangeBatch
	// Creates a batch of records to modify.
	InitWithRecordsToSaveRecordIDsToDeleteAtomicByZone(recordsToSave []CKRecord, recordIDsToDelete []CKRecordID, atomicByZone bool) CKSyncEngineRecordZoneChangeBatch

	// Topic: Managing atomicity

	// A Boolean value that determines whether CloudKit modifies records atomically by record zone.
	AtomicByZone() bool
	SetAtomicByZone(value bool)

	// Topic: Managing the records

	// The unique identifiers of the records to delete.
	RecordIDsToDelete() []CKRecordID
	// The records to save.
	RecordsToSave() []CKRecord
}

// Init initializes the instance.
func (c CKSyncEngineRecordZoneChangeBatch) Init() CKSyncEngineRecordZoneChangeBatch {
	rv := objc.Send[CKSyncEngineRecordZoneChangeBatch](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKSyncEngineRecordZoneChangeBatch) Autorelease() CKSyncEngineRecordZoneChangeBatch {
	rv := objc.Send[CKSyncEngineRecordZoneChangeBatch](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineRecordZoneChangeBatch creates a new CKSyncEngineRecordZoneChangeBatch instance.
func NewCKSyncEngineRecordZoneChangeBatch() CKSyncEngineRecordZoneChangeBatch {
	class := getCKSyncEngineRecordZoneChangeBatchClass()
	rv := objc.Send[CKSyncEngineRecordZoneChangeBatch](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a batch of records to modify.
//
// recordsToSave: The records to save.
//
// recordIDsToDelete: The identifiers of the records to delete.
//
// atomicByZone: A Boolean value that determines whether CloudKit modifies the specified
// records atomically by record zone.
//
// # Return Value
//
// An initialized change batch, or `nil` if CloudKit can’t create one.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineRecordZoneChangeBatch/initWithRecordsToSave:recordIDsToDelete:atomicByZone:
func NewCKSyncEngineRecordZoneChangeBatchWithRecordsToSaveRecordIDsToDeleteAtomicByZone(recordsToSave []CKRecord, recordIDsToDelete []CKRecordID, atomicByZone bool) CKSyncEngineRecordZoneChangeBatch {
	instance := getCKSyncEngineRecordZoneChangeBatchClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRecordsToSave:recordIDsToDelete:atomicByZone:"), objectivec.IObjectSliceToNSArray(recordsToSave), objectivec.IObjectSliceToNSArray(recordIDsToDelete), atomicByZone)
	return CKSyncEngineRecordZoneChangeBatchFromID(rv)
}

// Creates a batch of records to modify using the provided record zone
// changes.
//
// pendingChanges: The record zone changes to process.
//
// recordProvider: A block that returns the record for the specified record identifier.
//
// # Return Value
//
// The batch of records to modify, or `nil` if there are no pending changes.
//
// # Discussion
//
// This method iterates over `pendingChanges` and adds the necessary
// information to the new batch, until there are no more changes or the size
// of the batch reaches the maximum limit. If the type of change is a record
// save, the method asks the specified `recordProvider` block for that record.
// If the closure returns `nil`, the method skips that change.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineRecordZoneChangeBatch/initWithPendingChanges:recordProvider:
func (c CKSyncEngineRecordZoneChangeBatch) InitWithPendingChangesRecordProvider(pendingChanges []CKSyncEnginePendingRecordZoneChange, recordProvider CKRecordIDHandler) CKSyncEngineRecordZoneChangeBatch {
	_block1, _ := NewCKRecordIDBlock(recordProvider)
	rv := objc.Send[CKSyncEngineRecordZoneChangeBatch](c.ID, objc.Sel("initWithPendingChanges:recordProvider:"), pendingChanges, _block1)
	return rv
}

// Creates a batch of records to modify.
//
// recordsToSave: The records to save.
//
// recordIDsToDelete: The identifiers of the records to delete.
//
// atomicByZone: A Boolean value that determines whether CloudKit modifies the specified
// records atomically by record zone.
//
// # Return Value
//
// An initialized change batch, or `nil` if CloudKit can’t create one.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineRecordZoneChangeBatch/initWithRecordsToSave:recordIDsToDelete:atomicByZone:
func (c CKSyncEngineRecordZoneChangeBatch) InitWithRecordsToSaveRecordIDsToDeleteAtomicByZone(recordsToSave []CKRecord, recordIDsToDelete []CKRecordID, atomicByZone bool) CKSyncEngineRecordZoneChangeBatch {
	rv := objc.Send[CKSyncEngineRecordZoneChangeBatch](c.ID, objc.Sel("initWithRecordsToSave:recordIDsToDelete:atomicByZone:"), objectivec.IObjectSliceToNSArray(recordsToSave), objectivec.IObjectSliceToNSArray(recordIDsToDelete), atomicByZone)
	return rv
}

// A Boolean value that determines whether CloudKit modifies records
// atomically by record zone.
//
// # Discussion
//
// When true, CloudKit processes record changes atomically by record zone, and
// if any individual change fails, all other changes in that record’s record
// zone fail and return an error of type [CKError.Code.batchRequestFailed].
//
// The default value is false.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineRecordZoneChangeBatch/atomicByZone
//
// [CKError.Code.batchRequestFailed]: https://developer.apple.com/documentation/CloudKit/CKError/Code/batchRequestFailed
func (c CKSyncEngineRecordZoneChangeBatch) AtomicByZone() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("atomicByZone"))
	return rv
}
func (c CKSyncEngineRecordZoneChangeBatch) SetAtomicByZone(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setAtomicByZone:"), value)
}

// The unique identifiers of the records to delete.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineRecordZoneChangeBatch/recordIDsToDelete
func (c CKSyncEngineRecordZoneChangeBatch) RecordIDsToDelete() []CKRecordID {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("recordIDsToDelete"))
	return objc.ConvertSlice(rv, func(id objc.ID) CKRecordID {
		return CKRecordIDFromID(id)
	})
}

// The records to save.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineRecordZoneChangeBatch/recordsToSave
func (c CKSyncEngineRecordZoneChangeBatch) RecordsToSave() []CKRecord {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("recordsToSave"))
	return objc.ConvertSlice(rv, func(id objc.ID) CKRecord {
		return CKRecordFromID(id)
	})
}
