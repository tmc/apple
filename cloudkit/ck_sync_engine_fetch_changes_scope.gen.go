// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKSyncEngineFetchChangesScope] class.
var (
	_CKSyncEngineFetchChangesScopeClass     CKSyncEngineFetchChangesScopeClass
	_CKSyncEngineFetchChangesScopeClassOnce sync.Once
)

func getCKSyncEngineFetchChangesScopeClass() CKSyncEngineFetchChangesScopeClass {
	_CKSyncEngineFetchChangesScopeClassOnce.Do(func() {
		_CKSyncEngineFetchChangesScopeClass = CKSyncEngineFetchChangesScopeClass{class: objc.GetClass("CKSyncEngineFetchChangesScope")}
	})
	return _CKSyncEngineFetchChangesScopeClass
}

// GetCKSyncEngineFetchChangesScopeClass returns the class object for CKSyncEngineFetchChangesScope.
func GetCKSyncEngineFetchChangesScopeClass() CKSyncEngineFetchChangesScopeClass {
	return getCKSyncEngineFetchChangesScopeClass()
}

type CKSyncEngineFetchChangesScopeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKSyncEngineFetchChangesScopeClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKSyncEngineFetchChangesScopeClass) Alloc() CKSyncEngineFetchChangesScope {
	rv := objc.Send[CKSyncEngineFetchChangesScope](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A scope in which the sync engine will fetch changes from the server.
//
// # Instance Properties
//
//   - [CKSyncEngineFetchChangesScope.ExcludedZoneIDs]
//   - [CKSyncEngineFetchChangesScope.ZoneIDs]
//
// # Instance Methods
//
//   - [CKSyncEngineFetchChangesScope.ContainsZoneID]
//   - [CKSyncEngineFetchChangesScope.InitWithExcludedZoneIDs]
//   - [CKSyncEngineFetchChangesScope.InitWithZoneIDs]
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchChangesScope
type CKSyncEngineFetchChangesScope struct {
	objectivec.Object
}

// CKSyncEngineFetchChangesScopeFromID constructs a [CKSyncEngineFetchChangesScope] from an objc.ID.
//
// A scope in which the sync engine will fetch changes from the server.
func CKSyncEngineFetchChangesScopeFromID(id objc.ID) CKSyncEngineFetchChangesScope {
	return CKSyncEngineFetchChangesScope{objectivec.Object{ID: id}}
}

// NOTE: CKSyncEngineFetchChangesScope adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKSyncEngineFetchChangesScope] class.
//
// # Instance Properties
//
//   - [ICKSyncEngineFetchChangesScope.ExcludedZoneIDs]
//   - [ICKSyncEngineFetchChangesScope.ZoneIDs]
//
// # Instance Methods
//
//   - [ICKSyncEngineFetchChangesScope.ContainsZoneID]
//   - [ICKSyncEngineFetchChangesScope.InitWithExcludedZoneIDs]
//   - [ICKSyncEngineFetchChangesScope.InitWithZoneIDs]
//
// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchChangesScope
type ICKSyncEngineFetchChangesScope interface {
	objectivec.IObject

	// Topic: Instance Properties

	ExcludedZoneIDs() foundation.INSSet
	ZoneIDs() foundation.INSSet

	// Topic: Instance Methods

	ContainsZoneID(zoneID ICKRecordZoneID) bool
	InitWithExcludedZoneIDs(zoneIDs foundation.INSSet) CKSyncEngineFetchChangesScope
	InitWithZoneIDs(zoneIDs foundation.INSSet) CKSyncEngineFetchChangesScope
}

// Init initializes the instance.
func (c CKSyncEngineFetchChangesScope) Init() CKSyncEngineFetchChangesScope {
	rv := objc.Send[CKSyncEngineFetchChangesScope](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKSyncEngineFetchChangesScope) Autorelease() CKSyncEngineFetchChangesScope {
	rv := objc.Send[CKSyncEngineFetchChangesScope](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKSyncEngineFetchChangesScope creates a new CKSyncEngineFetchChangesScope instance.
func NewCKSyncEngineFetchChangesScope() CKSyncEngineFetchChangesScope {
	class := getCKSyncEngineFetchChangesScopeClass()
	rv := objc.Send[CKSyncEngineFetchChangesScope](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchChangesScope/initWithExcludedZoneIDs:
func NewCKSyncEngineFetchChangesScopeWithExcludedZoneIDs(zoneIDs foundation.INSSet) CKSyncEngineFetchChangesScope {
	instance := getCKSyncEngineFetchChangesScopeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithExcludedZoneIDs:"), zoneIDs)
	return CKSyncEngineFetchChangesScopeFromID(rv)
}

// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchChangesScope/initWithZoneIDs:
func NewCKSyncEngineFetchChangesScopeWithZoneIDs(zoneIDs foundation.INSSet) CKSyncEngineFetchChangesScope {
	instance := getCKSyncEngineFetchChangesScopeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithZoneIDs:"), zoneIDs)
	return CKSyncEngineFetchChangesScopeFromID(rv)
}

// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchChangesScope/containsZoneID:
func (c CKSyncEngineFetchChangesScope) ContainsZoneID(zoneID ICKRecordZoneID) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("containsZoneID:"), zoneID)
	return rv
}

// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchChangesScope/initWithExcludedZoneIDs:
func (c CKSyncEngineFetchChangesScope) InitWithExcludedZoneIDs(zoneIDs foundation.INSSet) CKSyncEngineFetchChangesScope {
	rv := objc.Send[CKSyncEngineFetchChangesScope](c.ID, objc.Sel("initWithExcludedZoneIDs:"), zoneIDs)
	return rv
}

// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchChangesScope/initWithZoneIDs:
func (c CKSyncEngineFetchChangesScope) InitWithZoneIDs(zoneIDs foundation.INSSet) CKSyncEngineFetchChangesScope {
	rv := objc.Send[CKSyncEngineFetchChangesScope](c.ID, objc.Sel("initWithZoneIDs:"), zoneIDs)
	return rv
}

// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchChangesScope/excludedZoneIDs
func (c CKSyncEngineFetchChangesScope) ExcludedZoneIDs() foundation.INSSet {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("excludedZoneIDs"))
	return foundation.NSSetFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CloudKit/CKSyncEngineFetchChangesScope/zoneIDs
func (c CKSyncEngineFetchChangesScope) ZoneIDs() foundation.INSSet {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("zoneIDs"))
	return foundation.NSSetFromID(objc.ID(rv))
}
