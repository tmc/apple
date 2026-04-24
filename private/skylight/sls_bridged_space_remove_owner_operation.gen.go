// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedSpaceRemoveOwnerOperation] class.
var (
	_SLSBridgedSpaceRemoveOwnerOperationClass     SLSBridgedSpaceRemoveOwnerOperationClass
	_SLSBridgedSpaceRemoveOwnerOperationClassOnce sync.Once
)

func getSLSBridgedSpaceRemoveOwnerOperationClass() SLSBridgedSpaceRemoveOwnerOperationClass {
	_SLSBridgedSpaceRemoveOwnerOperationClassOnce.Do(func() {
		_SLSBridgedSpaceRemoveOwnerOperationClass = SLSBridgedSpaceRemoveOwnerOperationClass{class: objc.GetClass("SLSBridgedSpaceRemoveOwnerOperation")}
	})
	return _SLSBridgedSpaceRemoveOwnerOperationClass
}

// GetSLSBridgedSpaceRemoveOwnerOperationClass returns the class object for SLSBridgedSpaceRemoveOwnerOperation.
func GetSLSBridgedSpaceRemoveOwnerOperationClass() SLSBridgedSpaceRemoveOwnerOperationClass {
	return getSLSBridgedSpaceRemoveOwnerOperationClass()
}

type SLSBridgedSpaceRemoveOwnerOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedSpaceRemoveOwnerOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedSpaceRemoveOwnerOperationClass) Alloc() SLSBridgedSpaceRemoveOwnerOperation {
	rv := objc.Send[SLSBridgedSpaceRemoveOwnerOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedSpaceRemoveOwnerOperation.Owner]
//   - [SLSBridgedSpaceRemoveOwnerOperation.SpaceID]
//   - [SLSBridgedSpaceRemoveOwnerOperation.InitWithSpaceIDOwner]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceRemoveOwnerOperation
type SLSBridgedSpaceRemoveOwnerOperation struct {
	SLSAsynchronousBridgedWindowManagementOperation
}

// SLSBridgedSpaceRemoveOwnerOperationFromID constructs a [SLSBridgedSpaceRemoveOwnerOperation] from an objc.ID.
func SLSBridgedSpaceRemoveOwnerOperationFromID(id objc.ID) SLSBridgedSpaceRemoveOwnerOperation {
	return SLSBridgedSpaceRemoveOwnerOperation{SLSAsynchronousBridgedWindowManagementOperation: SLSAsynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedSpaceRemoveOwnerOperation implements ISLSBridgedSpaceRemoveOwnerOperation.
var _ ISLSBridgedSpaceRemoveOwnerOperation = SLSBridgedSpaceRemoveOwnerOperation{}

// An interface definition for the [SLSBridgedSpaceRemoveOwnerOperation] class.
//
// # Methods
//
//   - [ISLSBridgedSpaceRemoveOwnerOperation.Owner]
//   - [ISLSBridgedSpaceRemoveOwnerOperation.SpaceID]
//   - [ISLSBridgedSpaceRemoveOwnerOperation.InitWithSpaceIDOwner]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceRemoveOwnerOperation
type ISLSBridgedSpaceRemoveOwnerOperation interface {
	ISLSAsynchronousBridgedWindowManagementOperation

	// Topic: Methods

	Owner() int
	SpaceID() uint64
	InitWithSpaceIDOwner(id uint64, owner int) SLSBridgedSpaceRemoveOwnerOperation
}

// Init initializes the instance.
func (s SLSBridgedSpaceRemoveOwnerOperation) Init() SLSBridgedSpaceRemoveOwnerOperation {
	rv := objc.Send[SLSBridgedSpaceRemoveOwnerOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedSpaceRemoveOwnerOperation) Autorelease() SLSBridgedSpaceRemoveOwnerOperation {
	rv := objc.Send[SLSBridgedSpaceRemoveOwnerOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedSpaceRemoveOwnerOperation creates a new SLSBridgedSpaceRemoveOwnerOperation instance.
func NewSLSBridgedSpaceRemoveOwnerOperation() SLSBridgedSpaceRemoveOwnerOperation {
	class := getSLSBridgedSpaceRemoveOwnerOperationClass()
	rv := objc.Send[SLSBridgedSpaceRemoveOwnerOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceRemoveOwnerOperation/initWithCoder:
func NewSLSBridgedSpaceRemoveOwnerOperationWithCoder(coder objectivec.IObject) SLSBridgedSpaceRemoveOwnerOperation {
	instance := getSLSBridgedSpaceRemoveOwnerOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedSpaceRemoveOwnerOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceRemoveOwnerOperation/initWithSpaceID:owner:
func NewSLSBridgedSpaceRemoveOwnerOperationWithSpaceIDOwner(id uint64, owner int) SLSBridgedSpaceRemoveOwnerOperation {
	instance := getSLSBridgedSpaceRemoveOwnerOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpaceID:owner:"), id, owner)
	return SLSBridgedSpaceRemoveOwnerOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceRemoveOwnerOperation/initWithSpaceID:owner:
func (s SLSBridgedSpaceRemoveOwnerOperation) InitWithSpaceIDOwner(id uint64, owner int) SLSBridgedSpaceRemoveOwnerOperation {
	rv := objc.Send[SLSBridgedSpaceRemoveOwnerOperation](s.ID, objc.Sel("initWithSpaceID:owner:"), id, owner)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceRemoveOwnerOperation/owner
func (s SLSBridgedSpaceRemoveOwnerOperation) Owner() int {
	rv := objc.Send[int](s.ID, objc.Sel("owner"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceRemoveOwnerOperation/spaceID
func (s SLSBridgedSpaceRemoveOwnerOperation) SpaceID() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}
