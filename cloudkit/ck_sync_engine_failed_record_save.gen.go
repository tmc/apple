// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKSyncEngineFailedRecordSave] class.
var (
	_CKSyncEngineFailedRecordSaveClass     CKSyncEngineFailedRecordSaveClass
	_CKSyncEngineFailedRecordSaveClassOnce sync.Once
)

func getCKSyncEngineFailedRecordSaveClass() CKSyncEngineFailedRecordSaveClass {
	_CKSyncEngineFailedRecordSaveClassOnce.Do(func() {
		_CKSyncEngineFailedRecordSaveClass = CKSyncEngineFailedRecordSaveClass{class: objc.GetClass("CKSyncEngineFailedRecordSave")}
	})
	return _CKSyncEngineFailedRecordSaveClass
}

// GetCKSyncEngineFailedRecordSaveClass returns the class object for CKSyncEngineFailedRecordSave.
func GetCKSyncEngineFailedRecordSaveClass() CKSyncEngineFailedRecordSaveClass {
	return getCKSyncEngineFailedRecordSaveClass()
}

type CKSyncEngineFailedRecordSaveClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKSyncEngineFailedRecordSaveClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKSyncEngineFailedRecordSaveClass) Alloc() CKSyncEngineFailedRecordSave {
	rv := objc.Send[CKSyncEngineFailedRecordSave](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A type that describes an unsuccessful attempt to modify an individual
// record.
//
// # Accessing the record
//
//   - [CKSyncEngineFailedRecordSave.Record]: The record that CloudKit is unable to modify.
//
// # Accessing the error
//
//   - [CKSyncEngineFailedRecordSave.Error]: A error that describes the reason for the unsuccessful attempt to modify the associated record.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFailedRecordSave
type CKSyncEngineFailedRecordSave struct {
	objectivec.Object
}

// CKSyncEngineFailedRecordSaveFromID constructs a [CKSyncEngineFailedRecordSave] from an objc.ID.
//
// A type that describes an unsuccessful attempt to modify an individual
// record.
func CKSyncEngineFailedRecordSaveFromID(id objc.ID) CKSyncEngineFailedRecordSave {
	return CKSyncEngineFailedRecordSave{objectivec.Object{ID: id}}
}

// Ensure CKSyncEngineFailedRecordSave implements ICKSyncEngineFailedRecordSave.
var _ ICKSyncEngineFailedRecordSave = CKSyncEngineFailedRecordSave{}

// An interface definition for the [CKSyncEngineFailedRecordSave] class.
//
// # Accessing the record
//
//   - [ICKSyncEngineFailedRecordSave.Record]: The record that CloudKit is unable to modify.
//
// # Accessing the error
//
//   - [ICKSyncEngineFailedRecordSave.Error]: A error that describes the reason for the unsuccessful attempt to modify the associated record.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFailedRecordSave
type ICKSyncEngineFailedRecordSave interface {
	objectivec.IObject

	// Topic: Accessing the record

	// The record that CloudKit is unable to modify.
	Record() ICKRecord

	// Topic: Accessing the error

	// A error that describes the reason for the unsuccessful attempt to modify the associated record.
	Error() foundation.INSError
}

// Init initializes the instance.
func (c CKSyncEngineFailedRecordSave) Init() CKSyncEngineFailedRecordSave {
	rv := objc.Send[CKSyncEngineFailedRecordSave](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKSyncEngineFailedRecordSave) Autorelease() CKSyncEngineFailedRecordSave {
	rv := objc.Send[CKSyncEngineFailedRecordSave](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineFailedRecordSave creates a new CKSyncEngineFailedRecordSave instance.
func NewCKSyncEngineFailedRecordSave() CKSyncEngineFailedRecordSave {
	class := getCKSyncEngineFailedRecordSaveClass()
	rv := objc.Send[CKSyncEngineFailedRecordSave](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The record that CloudKit is unable to modify.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFailedRecordSave/record
func (c CKSyncEngineFailedRecordSave) Record() ICKRecord {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("record"))
	return CKRecordFromID(objc.ID(rv))
}

// A error that describes the reason for the unsuccessful attempt to modify
// the associated record.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFailedRecordSave/error
func (c CKSyncEngineFailedRecordSave) Error() foundation.INSError {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("error"))
	return foundation.NSErrorFromID(objc.ID(rv))
}
