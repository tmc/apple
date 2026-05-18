// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKModifyRecordZonesOperation] class.
var (
	_CKModifyRecordZonesOperationClass     CKModifyRecordZonesOperationClass
	_CKModifyRecordZonesOperationClassOnce sync.Once
)

func getCKModifyRecordZonesOperationClass() CKModifyRecordZonesOperationClass {
	_CKModifyRecordZonesOperationClassOnce.Do(func() {
		_CKModifyRecordZonesOperationClass = CKModifyRecordZonesOperationClass{class: objc.GetClass("CKModifyRecordZonesOperation")}
	})
	return _CKModifyRecordZonesOperationClass
}

// GetCKModifyRecordZonesOperationClass returns the class object for CKModifyRecordZonesOperation.
func GetCKModifyRecordZonesOperationClass() CKModifyRecordZonesOperationClass {
	return getCKModifyRecordZonesOperationClass()
}

type CKModifyRecordZonesOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKModifyRecordZonesOperationClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKModifyRecordZonesOperationClass) Alloc() CKModifyRecordZonesOperation {
	rv := objc.Send[CKModifyRecordZonesOperation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An operation that modifies one or more record zones.
//
// # Overview
//
// After you create one or more record zones, use this operation to save those
// zones to the database. You can also use the operation to delete record
// zones and their records.
//
// If you assign a handler to the [CKModifyRecordZonesOperation.CompletionBlock] property of the operation,
// CloudKit calls the handler after the operation executes and returns its
// results. Use the handler to perform housekeeping tasks for the operation,
// but don’t use it to process the results of the operation. The handler you
// provide should manage any failures of the operation, whether due to an
// error or an explicit cancellation.
//
// # Configuring the Modify Zones Operation
//
//   - [CKModifyRecordZonesOperation.RecordZonesToSave]: The record zones to save to the database.
//   - [CKModifyRecordZonesOperation.SetRecordZonesToSave]
//   - [CKModifyRecordZonesOperation.RecordZoneIDsToDelete]: The IDs of the record zones to delete permanently from the database.
//   - [CKModifyRecordZonesOperation.SetRecordZoneIDsToDelete]
//
// # Instance Properties
//
//   - [CKModifyRecordZonesOperation.ModifyRecordZonesResultBlock]: The closure to execute after CloudKit modifies all of the record zones.
//   - [CKModifyRecordZonesOperation.SetModifyRecordZonesResultBlock]
//   - [CKModifyRecordZonesOperation.PerRecordZoneDeleteBlock]: The closure to execute when CloudKit deletes a record zone.
//   - [CKModifyRecordZonesOperation.SetPerRecordZoneDeleteBlock]
//   - [CKModifyRecordZonesOperation.PerRecordZoneSaveBlock]: The closure to execute when CloudKit saves a record zone.
//   - [CKModifyRecordZonesOperation.SetPerRecordZoneSaveBlock]
//
// See: https://developer.apple.com/documentation/CloudKit/CKModifyRecordZonesOperation
type CKModifyRecordZonesOperation struct {
	CKDatabaseOperation
}

// CKModifyRecordZonesOperationFromID constructs a [CKModifyRecordZonesOperation] from an objc.ID.
//
// An operation that modifies one or more record zones.
func CKModifyRecordZonesOperationFromID(id objc.ID) CKModifyRecordZonesOperation {
	return CKModifyRecordZonesOperation{CKDatabaseOperation: CKDatabaseOperationFromID(id)}
}

// NOTE: CKModifyRecordZonesOperation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKModifyRecordZonesOperation] class.
//
// # Configuring the Modify Zones Operation
//
//   - [ICKModifyRecordZonesOperation.RecordZonesToSave]: The record zones to save to the database.
//   - [ICKModifyRecordZonesOperation.SetRecordZonesToSave]
//   - [ICKModifyRecordZonesOperation.RecordZoneIDsToDelete]: The IDs of the record zones to delete permanently from the database.
//   - [ICKModifyRecordZonesOperation.SetRecordZoneIDsToDelete]
//
// # Instance Properties
//
//   - [ICKModifyRecordZonesOperation.ModifyRecordZonesResultBlock]: The closure to execute after CloudKit modifies all of the record zones.
//   - [ICKModifyRecordZonesOperation.SetModifyRecordZonesResultBlock]
//   - [ICKModifyRecordZonesOperation.PerRecordZoneDeleteBlock]: The closure to execute when CloudKit deletes a record zone.
//   - [ICKModifyRecordZonesOperation.SetPerRecordZoneDeleteBlock]
//   - [ICKModifyRecordZonesOperation.PerRecordZoneSaveBlock]: The closure to execute when CloudKit saves a record zone.
//   - [ICKModifyRecordZonesOperation.SetPerRecordZoneSaveBlock]
//
// See: https://developer.apple.com/documentation/CloudKit/CKModifyRecordZonesOperation
type ICKModifyRecordZonesOperation interface {
	ICKDatabaseOperation

	// Topic: Configuring the Modify Zones Operation

	// The record zones to save to the database.
	RecordZonesToSave() []CKRecordZone
	SetRecordZonesToSave(value []CKRecordZone)
	// The IDs of the record zones to delete permanently from the database.
	RecordZoneIDsToDelete() []CKRecordZoneID
	SetRecordZoneIDsToDelete(value []CKRecordZoneID)

	// Topic: Instance Properties

	// The closure to execute after CloudKit modifies all of the record zones.
	ModifyRecordZonesResultBlock() unsafe.Pointer
	SetModifyRecordZonesResultBlock(value unsafe.Pointer)
	// The closure to execute when CloudKit deletes a record zone.
	PerRecordZoneDeleteBlock() unsafe.Pointer
	SetPerRecordZoneDeleteBlock(value unsafe.Pointer)
	// The closure to execute when CloudKit saves a record zone.
	PerRecordZoneSaveBlock() unsafe.Pointer
	SetPerRecordZoneSaveBlock(value unsafe.Pointer)
}

// Init initializes the instance.
func (c CKModifyRecordZonesOperation) Init() CKModifyRecordZonesOperation {
	rv := objc.Send[CKModifyRecordZonesOperation](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKModifyRecordZonesOperation) Autorelease() CKModifyRecordZonesOperation {
	rv := objc.Send[CKModifyRecordZonesOperation](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKModifyRecordZonesOperation creates a new CKModifyRecordZonesOperation instance.
func NewCKModifyRecordZonesOperation() CKModifyRecordZonesOperation {
	class := getCKModifyRecordZonesOperationClass()
	rv := objc.Send[CKModifyRecordZonesOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The record zones to save to the database.
//
// # Discussion
//
// The initial value of the property is the array that you provide to the
// [initWithRecordZonesToSave:recordZoneIDsToDelete:] method. You can modify
// this array as necessary before you execute the operation. The record zones
// must all target the same database. You can specify `nil`, or an empty
// array, for this property.
//
// If you intend to change the value of this property, do so before you
// execute the operation or submit the operation to a queue.
//
// See: https://developer.apple.com/documentation/CloudKit/CKModifyRecordZonesOperation/recordZonesToSave
//
// [initWithRecordZonesToSave:recordZoneIDsToDelete:]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordZonesOperation/initWithRecordZonesToSave:recordZoneIDsToDelete:
func (c CKModifyRecordZonesOperation) RecordZonesToSave() []CKRecordZone {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("recordZonesToSave"))
	return objc.ConvertSlice(rv, func(id objc.ID) CKRecordZone {
		return CKRecordZoneFromID(id)
	})
}
func (c CKModifyRecordZonesOperation) SetRecordZonesToSave(value []CKRecordZone) {
	objc.Send[struct{}](c.ID, objc.Sel("setRecordZonesToSave:"), objectivec.IObjectSliceToNSArray(value))
}

// The IDs of the record zones to delete permanently from the database.
//
// # Discussion
//
// The initial value of the property is the array of zone IDs that you provide
// to the [initWithRecordZonesToSave:recordZoneIDsToDelete:] method. You can
// modify this array as necessary before you execute the operation. The record
// zones must all target the same database. You can specify `nil`, or an empty
// array, for this property.
//
// If you intend to change the value of this property, do so before you
// execute the operation or submit the operation to a queue.
//
// See: https://developer.apple.com/documentation/CloudKit/CKModifyRecordZonesOperation/recordZoneIDsToDelete
//
// [initWithRecordZonesToSave:recordZoneIDsToDelete:]: https://developer.apple.com/documentation/CloudKit/CKModifyRecordZonesOperation/initWithRecordZonesToSave:recordZoneIDsToDelete:
func (c CKModifyRecordZonesOperation) RecordZoneIDsToDelete() []CKRecordZoneID {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("recordZoneIDsToDelete"))
	return objc.ConvertSlice(rv, func(id objc.ID) CKRecordZoneID {
		return CKRecordZoneIDFromID(id)
	})
}
func (c CKModifyRecordZonesOperation) SetRecordZoneIDsToDelete(value []CKRecordZoneID) {
	objc.Send[struct{}](c.ID, objc.Sel("setRecordZoneIDsToDelete:"), objectivec.IObjectSliceToNSArray(value))
}

// The closure to execute after CloudKit modifies all of the record zones.
//
// See: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordzonesoperation/modifyrecordzonesresultblock
func (c CKModifyRecordZonesOperation) ModifyRecordZonesResultBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("modifyRecordZonesResultBlock"))
	return rv
}
func (c CKModifyRecordZonesOperation) SetModifyRecordZonesResultBlock(value unsafe.Pointer) {
	objc.Send[struct{}](c.ID, objc.Sel("setModifyRecordZonesResultBlock:"), value)
}

// The closure to execute when CloudKit deletes a record zone.
//
// See: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordzonesoperation/perrecordzonedeleteblock-6c82y
func (c CKModifyRecordZonesOperation) PerRecordZoneDeleteBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("perRecordZoneDeleteBlock"))
	return rv
}
func (c CKModifyRecordZonesOperation) SetPerRecordZoneDeleteBlock(value unsafe.Pointer) {
	objc.Send[struct{}](c.ID, objc.Sel("setPerRecordZoneDeleteBlock:"), value)
}

// The closure to execute when CloudKit saves a record zone.
//
// See: https://developer.apple.com/documentation/cloudkit/ckmodifyrecordzonesoperation/perrecordzonesaveblock-1m45y
func (c CKModifyRecordZonesOperation) PerRecordZoneSaveBlock() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("perRecordZoneSaveBlock"))
	return rv
}
func (c CKModifyRecordZonesOperation) SetPerRecordZoneSaveBlock(value unsafe.Pointer) {
	objc.Send[struct{}](c.ID, objc.Sel("setPerRecordZoneSaveBlock:"), value)
}
