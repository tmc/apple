// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKSyncEngineFetchedRecordDeletion] class.
var (
	_CKSyncEngineFetchedRecordDeletionClass     CKSyncEngineFetchedRecordDeletionClass
	_CKSyncEngineFetchedRecordDeletionClassOnce sync.Once
)

func getCKSyncEngineFetchedRecordDeletionClass() CKSyncEngineFetchedRecordDeletionClass {
	_CKSyncEngineFetchedRecordDeletionClassOnce.Do(func() {
		_CKSyncEngineFetchedRecordDeletionClass = CKSyncEngineFetchedRecordDeletionClass{class: objc.GetClass("CKSyncEngineFetchedRecordDeletion")}
	})
	return _CKSyncEngineFetchedRecordDeletionClass
}

// GetCKSyncEngineFetchedRecordDeletionClass returns the class object for CKSyncEngineFetchedRecordDeletion.
func GetCKSyncEngineFetchedRecordDeletionClass() CKSyncEngineFetchedRecordDeletionClass {
	return getCKSyncEngineFetchedRecordDeletionClass()
}

type CKSyncEngineFetchedRecordDeletionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKSyncEngineFetchedRecordDeletionClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKSyncEngineFetchedRecordDeletionClass) Alloc() CKSyncEngineFetchedRecordDeletion {
	rv := objc.Send[CKSyncEngineFetchedRecordDeletion](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object that describes the deletion of an individual record.
//
// # Understanding the deletion
//
//   - [CKSyncEngineFetchedRecordDeletion.RecordID]: The deleted record’s unique identifier.
//   - [CKSyncEngineFetchedRecordDeletion.RecordType]: The record type of the deleted record.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchedRecordDeletion
type CKSyncEngineFetchedRecordDeletion struct {
	objectivec.Object
}

// CKSyncEngineFetchedRecordDeletionFromID constructs a [CKSyncEngineFetchedRecordDeletion] from an objc.ID.
//
// An object that describes the deletion of an individual record.
func CKSyncEngineFetchedRecordDeletionFromID(id objc.ID) CKSyncEngineFetchedRecordDeletion {
	return CKSyncEngineFetchedRecordDeletion{objectivec.Object{ID: id}}
}

// Ensure CKSyncEngineFetchedRecordDeletion implements ICKSyncEngineFetchedRecordDeletion.
var _ ICKSyncEngineFetchedRecordDeletion = CKSyncEngineFetchedRecordDeletion{}

// An interface definition for the [CKSyncEngineFetchedRecordDeletion] class.
//
// # Understanding the deletion
//
//   - [ICKSyncEngineFetchedRecordDeletion.RecordID]: The deleted record’s unique identifier.
//   - [ICKSyncEngineFetchedRecordDeletion.RecordType]: The record type of the deleted record.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchedRecordDeletion
type ICKSyncEngineFetchedRecordDeletion interface {
	objectivec.IObject

	// Topic: Understanding the deletion

	// The deleted record’s unique identifier.
	RecordID() ICKRecordID
	// The record type of the deleted record.
	RecordType() CKRecordType
}

// Init initializes the instance.
func (c CKSyncEngineFetchedRecordDeletion) Init() CKSyncEngineFetchedRecordDeletion {
	rv := objc.Send[CKSyncEngineFetchedRecordDeletion](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKSyncEngineFetchedRecordDeletion) Autorelease() CKSyncEngineFetchedRecordDeletion {
	rv := objc.Send[CKSyncEngineFetchedRecordDeletion](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineFetchedRecordDeletion creates a new CKSyncEngineFetchedRecordDeletion instance.
func NewCKSyncEngineFetchedRecordDeletion() CKSyncEngineFetchedRecordDeletion {
	class := getCKSyncEngineFetchedRecordDeletionClass()
	rv := objc.Send[CKSyncEngineFetchedRecordDeletion](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The deleted record’s unique identifier.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchedRecordDeletion/recordID
func (c CKSyncEngineFetchedRecordDeletion) RecordID() ICKRecordID {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("recordID"))
	return CKRecordIDFromID(objc.ID(rv))
}

// The record type of the deleted record.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchedRecordDeletion/recordType
func (c CKSyncEngineFetchedRecordDeletion) RecordType() CKRecordType {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("recordType"))
	return CKRecordType(foundation.NSStringFromID(rv).String())
}
