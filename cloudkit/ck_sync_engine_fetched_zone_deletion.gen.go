// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKSyncEngineFetchedZoneDeletion] class.
var (
	_CKSyncEngineFetchedZoneDeletionClass     CKSyncEngineFetchedZoneDeletionClass
	_CKSyncEngineFetchedZoneDeletionClassOnce sync.Once
)

func getCKSyncEngineFetchedZoneDeletionClass() CKSyncEngineFetchedZoneDeletionClass {
	_CKSyncEngineFetchedZoneDeletionClassOnce.Do(func() {
		_CKSyncEngineFetchedZoneDeletionClass = CKSyncEngineFetchedZoneDeletionClass{class: objc.GetClass("CKSyncEngineFetchedZoneDeletion")}
	})
	return _CKSyncEngineFetchedZoneDeletionClass
}

// GetCKSyncEngineFetchedZoneDeletionClass returns the class object for CKSyncEngineFetchedZoneDeletion.
func GetCKSyncEngineFetchedZoneDeletionClass() CKSyncEngineFetchedZoneDeletionClass {
	return getCKSyncEngineFetchedZoneDeletionClass()
}

type CKSyncEngineFetchedZoneDeletionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKSyncEngineFetchedZoneDeletionClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKSyncEngineFetchedZoneDeletionClass) Alloc() CKSyncEngineFetchedZoneDeletion {
	rv := objc.Send[CKSyncEngineFetchedZoneDeletion](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object that describes the deletion of a record zone.
//
// # Understanding the deletion
//
//   - [CKSyncEngineFetchedZoneDeletion.ZoneID]: The identifier of the deleted record zone.
//   - [CKSyncEngineFetchedZoneDeletion.Reason]: The reason for the deletion.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchedZoneDeletion
type CKSyncEngineFetchedZoneDeletion struct {
	objectivec.Object
}

// CKSyncEngineFetchedZoneDeletionFromID constructs a [CKSyncEngineFetchedZoneDeletion] from an objc.ID.
//
// An object that describes the deletion of a record zone.
func CKSyncEngineFetchedZoneDeletionFromID(id objc.ID) CKSyncEngineFetchedZoneDeletion {
	return CKSyncEngineFetchedZoneDeletion{objectivec.Object{ID: id}}
}

// Ensure CKSyncEngineFetchedZoneDeletion implements ICKSyncEngineFetchedZoneDeletion.
var _ ICKSyncEngineFetchedZoneDeletion = CKSyncEngineFetchedZoneDeletion{}

// An interface definition for the [CKSyncEngineFetchedZoneDeletion] class.
//
// # Understanding the deletion
//
//   - [ICKSyncEngineFetchedZoneDeletion.ZoneID]: The identifier of the deleted record zone.
//   - [ICKSyncEngineFetchedZoneDeletion.Reason]: The reason for the deletion.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchedZoneDeletion
type ICKSyncEngineFetchedZoneDeletion interface {
	objectivec.IObject

	// Topic: Understanding the deletion

	// The identifier of the deleted record zone.
	ZoneID() ICKRecordZoneID
	// The reason for the deletion.
	Reason() CKSyncEngineZoneDeletionReason
}

// Init initializes the instance.
func (c CKSyncEngineFetchedZoneDeletion) Init() CKSyncEngineFetchedZoneDeletion {
	rv := objc.Send[CKSyncEngineFetchedZoneDeletion](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKSyncEngineFetchedZoneDeletion) Autorelease() CKSyncEngineFetchedZoneDeletion {
	rv := objc.Send[CKSyncEngineFetchedZoneDeletion](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineFetchedZoneDeletion creates a new CKSyncEngineFetchedZoneDeletion instance.
func NewCKSyncEngineFetchedZoneDeletion() CKSyncEngineFetchedZoneDeletion {
	class := getCKSyncEngineFetchedZoneDeletionClass()
	rv := objc.Send[CKSyncEngineFetchedZoneDeletion](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The identifier of the deleted record zone.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchedZoneDeletion/zoneID
func (c CKSyncEngineFetchedZoneDeletion) ZoneID() ICKRecordZoneID {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("zoneID"))
	return CKRecordZoneIDFromID(objc.ID(rv))
}

// The reason for the deletion.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchedZoneDeletion/reason
func (c CKSyncEngineFetchedZoneDeletion) Reason() CKSyncEngineZoneDeletionReason {
	rv := objc.Send[CKSyncEngineZoneDeletionReason](c.ID, objc.Sel("reason"))
	return CKSyncEngineZoneDeletionReason(rv)
}
