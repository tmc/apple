// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKSyncEngineSendChangesScope] class.
var (
	_CKSyncEngineSendChangesScopeClass     CKSyncEngineSendChangesScopeClass
	_CKSyncEngineSendChangesScopeClassOnce sync.Once
)

func getCKSyncEngineSendChangesScopeClass() CKSyncEngineSendChangesScopeClass {
	_CKSyncEngineSendChangesScopeClassOnce.Do(func() {
		_CKSyncEngineSendChangesScopeClass = CKSyncEngineSendChangesScopeClass{class: objc.GetClass("CKSyncEngineSendChangesScope")}
	})
	return _CKSyncEngineSendChangesScopeClass
}

// GetCKSyncEngineSendChangesScopeClass returns the class object for CKSyncEngineSendChangesScope.
func GetCKSyncEngineSendChangesScopeClass() CKSyncEngineSendChangesScopeClass {
	return getCKSyncEngineSendChangesScopeClass()
}

type CKSyncEngineSendChangesScopeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKSyncEngineSendChangesScopeClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKSyncEngineSendChangesScopeClass) Alloc() CKSyncEngineSendChangesScope {
	rv := objc.Send[CKSyncEngineSendChangesScope](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A scope in which the sync engine will send changes to the server.
//
// # Instance Properties
//
//   - [CKSyncEngineSendChangesScope.ExcludedZoneIDs]
//   - [CKSyncEngineSendChangesScope.RecordIDs]
//   - [CKSyncEngineSendChangesScope.ZoneIDs]
//
// # Instance Methods
//
//   - [CKSyncEngineSendChangesScope.ContainsPendingRecordZoneChange]
//   - [CKSyncEngineSendChangesScope.ContainsRecordID]
//   - [CKSyncEngineSendChangesScope.InitWithExcludedZoneIDs]
//   - [CKSyncEngineSendChangesScope.InitWithRecordIDs]
//   - [CKSyncEngineSendChangesScope.InitWithZoneIDs]
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesScope
type CKSyncEngineSendChangesScope struct {
	objectivec.Object
}

// CKSyncEngineSendChangesScopeFromID constructs a [CKSyncEngineSendChangesScope] from an objc.ID.
//
// A scope in which the sync engine will send changes to the server.
func CKSyncEngineSendChangesScopeFromID(id objc.ID) CKSyncEngineSendChangesScope {
	return CKSyncEngineSendChangesScope{objectivec.Object{ID: id}}
}

// NOTE: CKSyncEngineSendChangesScope adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKSyncEngineSendChangesScope] class.
//
// # Instance Properties
//
//   - [ICKSyncEngineSendChangesScope.ExcludedZoneIDs]
//   - [ICKSyncEngineSendChangesScope.RecordIDs]
//   - [ICKSyncEngineSendChangesScope.ZoneIDs]
//
// # Instance Methods
//
//   - [ICKSyncEngineSendChangesScope.ContainsPendingRecordZoneChange]
//   - [ICKSyncEngineSendChangesScope.ContainsRecordID]
//   - [ICKSyncEngineSendChangesScope.InitWithExcludedZoneIDs]
//   - [ICKSyncEngineSendChangesScope.InitWithRecordIDs]
//   - [ICKSyncEngineSendChangesScope.InitWithZoneIDs]
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesScope
type ICKSyncEngineSendChangesScope interface {
	objectivec.IObject

	// Topic: Instance Properties

	ExcludedZoneIDs() foundation.INSSet
	RecordIDs() foundation.INSSet
	ZoneIDs() foundation.INSSet

	// Topic: Instance Methods

	ContainsPendingRecordZoneChange(pendingRecordZoneChange ICKSyncEnginePendingRecordZoneChange) bool
	ContainsRecordID(recordID ICKRecordID) bool
	InitWithExcludedZoneIDs(excludedZoneIDs foundation.INSSet) CKSyncEngineSendChangesScope
	InitWithRecordIDs(recordIDs foundation.INSSet) CKSyncEngineSendChangesScope
	InitWithZoneIDs(zoneIDs foundation.INSSet) CKSyncEngineSendChangesScope
}

// Init initializes the instance.
func (c CKSyncEngineSendChangesScope) Init() CKSyncEngineSendChangesScope {
	rv := objc.Send[CKSyncEngineSendChangesScope](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKSyncEngineSendChangesScope) Autorelease() CKSyncEngineSendChangesScope {
	rv := objc.Send[CKSyncEngineSendChangesScope](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineSendChangesScope creates a new CKSyncEngineSendChangesScope instance.
func NewCKSyncEngineSendChangesScope() CKSyncEngineSendChangesScope {
	class := getCKSyncEngineSendChangesScopeClass()
	rv := objc.Send[CKSyncEngineSendChangesScope](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesScope/initWithExcludedZoneIDs:
func NewCKSyncEngineSendChangesScopeWithExcludedZoneIDs(excludedZoneIDs foundation.INSSet) CKSyncEngineSendChangesScope {
	instance := getCKSyncEngineSendChangesScopeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithExcludedZoneIDs:"), excludedZoneIDs)
	return CKSyncEngineSendChangesScopeFromID(rv)
}

// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesScope/initWithRecordIDs:
func NewCKSyncEngineSendChangesScopeWithRecordIDs(recordIDs foundation.INSSet) CKSyncEngineSendChangesScope {
	instance := getCKSyncEngineSendChangesScopeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRecordIDs:"), recordIDs)
	return CKSyncEngineSendChangesScopeFromID(rv)
}

// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesScope/initWithZoneIDs:
func NewCKSyncEngineSendChangesScopeWithZoneIDs(zoneIDs foundation.INSSet) CKSyncEngineSendChangesScope {
	instance := getCKSyncEngineSendChangesScopeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithZoneIDs:"), zoneIDs)
	return CKSyncEngineSendChangesScopeFromID(rv)
}

// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesScope/containsPendingRecordZoneChange:
func (c CKSyncEngineSendChangesScope) ContainsPendingRecordZoneChange(pendingRecordZoneChange ICKSyncEnginePendingRecordZoneChange) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("containsPendingRecordZoneChange:"), pendingRecordZoneChange)
	return rv
}

// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesScope/containsRecordID:
func (c CKSyncEngineSendChangesScope) ContainsRecordID(recordID ICKRecordID) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("containsRecordID:"), recordID)
	return rv
}

// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesScope/initWithExcludedZoneIDs:
func (c CKSyncEngineSendChangesScope) InitWithExcludedZoneIDs(excludedZoneIDs foundation.INSSet) CKSyncEngineSendChangesScope {
	rv := objc.Send[CKSyncEngineSendChangesScope](c.ID, objc.Sel("initWithExcludedZoneIDs:"), excludedZoneIDs)
	return rv
}

// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesScope/initWithRecordIDs:
func (c CKSyncEngineSendChangesScope) InitWithRecordIDs(recordIDs foundation.INSSet) CKSyncEngineSendChangesScope {
	rv := objc.Send[CKSyncEngineSendChangesScope](c.ID, objc.Sel("initWithRecordIDs:"), recordIDs)
	return rv
}

// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesScope/initWithZoneIDs:
func (c CKSyncEngineSendChangesScope) InitWithZoneIDs(zoneIDs foundation.INSSet) CKSyncEngineSendChangesScope {
	rv := objc.Send[CKSyncEngineSendChangesScope](c.ID, objc.Sel("initWithZoneIDs:"), zoneIDs)
	return rv
}

// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesScope/excludedZoneIDs
func (c CKSyncEngineSendChangesScope) ExcludedZoneIDs() foundation.INSSet {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("excludedZoneIDs"))
	return foundation.NSSetFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesScope/recordIDs
func (c CKSyncEngineSendChangesScope) RecordIDs() foundation.INSSet {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("recordIDs"))
	return foundation.NSSetFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineSendChangesScope/zoneIDs
func (c CKSyncEngineSendChangesScope) ZoneIDs() foundation.INSSet {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("zoneIDs"))
	return foundation.NSSetFromID(objc.ID(rv))
}
