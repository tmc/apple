// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKSyncEngineFailedZoneSave] class.
var (
	_CKSyncEngineFailedZoneSaveClass     CKSyncEngineFailedZoneSaveClass
	_CKSyncEngineFailedZoneSaveClassOnce sync.Once
)

func getCKSyncEngineFailedZoneSaveClass() CKSyncEngineFailedZoneSaveClass {
	_CKSyncEngineFailedZoneSaveClassOnce.Do(func() {
		_CKSyncEngineFailedZoneSaveClass = CKSyncEngineFailedZoneSaveClass{class: objc.GetClass("CKSyncEngineFailedZoneSave")}
	})
	return _CKSyncEngineFailedZoneSaveClass
}

// GetCKSyncEngineFailedZoneSaveClass returns the class object for CKSyncEngineFailedZoneSave.
func GetCKSyncEngineFailedZoneSaveClass() CKSyncEngineFailedZoneSaveClass {
	return getCKSyncEngineFailedZoneSaveClass()
}

type CKSyncEngineFailedZoneSaveClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKSyncEngineFailedZoneSaveClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKSyncEngineFailedZoneSaveClass) Alloc() CKSyncEngineFailedZoneSave {
	rv := objc.Send[CKSyncEngineFailedZoneSave](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object that describes an unsuccessful attempt to modify a single record
// zone.
//
// # Accessing the record zone
//
//   - [CKSyncEngineFailedZoneSave.RecordZone]: The record zone that CloudKit is unable to modify.
//
// # Accessing the error
//
//   - [CKSyncEngineFailedZoneSave.Error]: A error that describes the reason for the unsuccessful attempt to modify the associated record zone.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFailedZoneSave
type CKSyncEngineFailedZoneSave struct {
	objectivec.Object
}

// CKSyncEngineFailedZoneSaveFromID constructs a [CKSyncEngineFailedZoneSave] from an objc.ID.
//
// An object that describes an unsuccessful attempt to modify a single record
// zone.
func CKSyncEngineFailedZoneSaveFromID(id objc.ID) CKSyncEngineFailedZoneSave {
	return CKSyncEngineFailedZoneSave{objectivec.Object{ID: id}}
}

// Ensure CKSyncEngineFailedZoneSave implements ICKSyncEngineFailedZoneSave.
var _ ICKSyncEngineFailedZoneSave = CKSyncEngineFailedZoneSave{}

// An interface definition for the [CKSyncEngineFailedZoneSave] class.
//
// # Accessing the record zone
//
//   - [ICKSyncEngineFailedZoneSave.RecordZone]: The record zone that CloudKit is unable to modify.
//
// # Accessing the error
//
//   - [ICKSyncEngineFailedZoneSave.Error]: A error that describes the reason for the unsuccessful attempt to modify the associated record zone.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFailedZoneSave
type ICKSyncEngineFailedZoneSave interface {
	objectivec.IObject

	// Topic: Accessing the record zone

	// The record zone that CloudKit is unable to modify.
	RecordZone() ICKRecordZone

	// Topic: Accessing the error

	// A error that describes the reason for the unsuccessful attempt to modify the associated record zone.
	Error() foundation.NSError
}

// Init initializes the instance.
func (c CKSyncEngineFailedZoneSave) Init() CKSyncEngineFailedZoneSave {
	rv := objc.Send[CKSyncEngineFailedZoneSave](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKSyncEngineFailedZoneSave) Autorelease() CKSyncEngineFailedZoneSave {
	rv := objc.Send[CKSyncEngineFailedZoneSave](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineFailedZoneSave creates a new CKSyncEngineFailedZoneSave instance.
func NewCKSyncEngineFailedZoneSave() CKSyncEngineFailedZoneSave {
	class := getCKSyncEngineFailedZoneSaveClass()
	rv := objc.Send[CKSyncEngineFailedZoneSave](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The record zone that CloudKit is unable to modify.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFailedZoneSave/recordZone
func (c CKSyncEngineFailedZoneSave) RecordZone() ICKRecordZone {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("recordZone"))
	return CKRecordZoneFromID(objc.ID(rv))
}

// A error that describes the reason for the unsuccessful attempt to modify
// the associated record zone.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFailedZoneSave/error
func (c CKSyncEngineFailedZoneSave) Error() foundation.NSError {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("error"))
	return foundation.NSErrorFromID(objc.ID(rv))
}
