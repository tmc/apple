// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSecureCursorAssertionManager] class.
var (
	_SLSecureCursorAssertionManagerClass     SLSecureCursorAssertionManagerClass
	_SLSecureCursorAssertionManagerClassOnce sync.Once
)

func getSLSecureCursorAssertionManagerClass() SLSecureCursorAssertionManagerClass {
	_SLSecureCursorAssertionManagerClassOnce.Do(func() {
		_SLSecureCursorAssertionManagerClass = SLSecureCursorAssertionManagerClass{class: objc.GetClass("SLSecureCursorAssertionManager")}
	})
	return _SLSecureCursorAssertionManagerClass
}

// GetSLSecureCursorAssertionManagerClass returns the class object for SLSecureCursorAssertionManager.
func GetSLSecureCursorAssertionManagerClass() SLSecureCursorAssertionManagerClass {
	return getSLSecureCursorAssertionManagerClass()
}

type SLSecureCursorAssertionManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSecureCursorAssertionManagerClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSecureCursorAssertionManagerClass) Alloc() SLSecureCursorAssertionManager {
	rv := objc.Send[SLSecureCursorAssertionManager](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSecureCursorAssertionManager.InvalidateAll]
//   - [SLSecureCursorAssertionManager.InvalidateUUID]
//   - [SLSecureCursorAssertionManager.StrongAssertionUUIDs]
//   - [SLSecureCursorAssertionManager.TakeAssertion]
//   - [SLSecureCursorAssertionManager.UnmapUUID]
//   - [SLSecureCursorAssertionManager.WeakAssertionMap]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSecureCursorAssertionManager
type SLSecureCursorAssertionManager struct {
	objectivec.Object
}

// SLSecureCursorAssertionManagerFromID constructs a [SLSecureCursorAssertionManager] from an objc.ID.
func SLSecureCursorAssertionManagerFromID(id objc.ID) SLSecureCursorAssertionManager {
	return SLSecureCursorAssertionManager{objectivec.Object{ID: id}}
}

// Ensure SLSecureCursorAssertionManager implements ISLSecureCursorAssertionManager.
var _ ISLSecureCursorAssertionManager = SLSecureCursorAssertionManager{}

// An interface definition for the [SLSecureCursorAssertionManager] class.
//
// # Methods
//
//   - [ISLSecureCursorAssertionManager.InvalidateAll]
//   - [ISLSecureCursorAssertionManager.InvalidateUUID]
//   - [ISLSecureCursorAssertionManager.StrongAssertionUUIDs]
//   - [ISLSecureCursorAssertionManager.TakeAssertion]
//   - [ISLSecureCursorAssertionManager.UnmapUUID]
//   - [ISLSecureCursorAssertionManager.WeakAssertionMap]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSecureCursorAssertionManager
type ISLSecureCursorAssertionManager interface {
	objectivec.IObject

	// Topic: Methods

	InvalidateAll()
	InvalidateUUID(uuid objectivec.IObject)
	StrongAssertionUUIDs() foundation.NSHashTable
	TakeAssertion() objectivec.IObject
	UnmapUUID(uuid objectivec.IObject)
	WeakAssertionMap() foundation.NSMapTable
}

// Init initializes the instance.
func (s SLSecureCursorAssertionManager) Init() SLSecureCursorAssertionManager {
	rv := objc.Send[SLSecureCursorAssertionManager](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSecureCursorAssertionManager) Autorelease() SLSecureCursorAssertionManager {
	rv := objc.Send[SLSecureCursorAssertionManager](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSecureCursorAssertionManager creates a new SLSecureCursorAssertionManager instance.
func NewSLSecureCursorAssertionManager() SLSecureCursorAssertionManager {
	class := getSLSecureCursorAssertionManagerClass()
	rv := objc.Send[SLSecureCursorAssertionManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSecureCursorAssertionManager/invalidateAll
func (s SLSecureCursorAssertionManager) InvalidateAll() {
	objc.Send[objc.ID](s.ID, objc.Sel("invalidateAll"))
}

// See: https://developer.apple.com/documentation/SkyLight/SLSecureCursorAssertionManager/invalidateUUID:
func (s SLSecureCursorAssertionManager) InvalidateUUID(uuid objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("invalidateUUID:"), uuid)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSecureCursorAssertionManager/takeAssertion
func (s SLSecureCursorAssertionManager) TakeAssertion() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("takeAssertion"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSecureCursorAssertionManager/unmapUUID:
func (s SLSecureCursorAssertionManager) UnmapUUID(uuid objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("unmapUUID:"), uuid)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSecureCursorAssertionManager/sharedManager
func (_SLSecureCursorAssertionManagerClass SLSecureCursorAssertionManagerClass) SharedManager() SLSecureCursorAssertionManager {
	rv := objc.Send[objc.ID](objc.ID(_SLSecureCursorAssertionManagerClass.class), objc.Sel("sharedManager"))
	return SLSecureCursorAssertionManagerFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSecureCursorAssertionManager/strongAssertionUUIDs
func (s SLSecureCursorAssertionManager) StrongAssertionUUIDs() foundation.NSHashTable {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("strongAssertionUUIDs"))
	return foundation.NSHashTableFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/SkyLight/SLSecureCursorAssertionManager/weakAssertionMap
func (s SLSecureCursorAssertionManager) WeakAssertionMap() foundation.NSMapTable {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("weakAssertionMap"))
	return foundation.NSMapTableFromID(objc.ID(rv))
}
