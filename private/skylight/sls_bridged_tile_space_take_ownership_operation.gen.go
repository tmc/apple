// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedTileSpaceTakeOwnershipOperation] class.
var (
	_SLSBridgedTileSpaceTakeOwnershipOperationClass     SLSBridgedTileSpaceTakeOwnershipOperationClass
	_SLSBridgedTileSpaceTakeOwnershipOperationClassOnce sync.Once
)

func getSLSBridgedTileSpaceTakeOwnershipOperationClass() SLSBridgedTileSpaceTakeOwnershipOperationClass {
	_SLSBridgedTileSpaceTakeOwnershipOperationClassOnce.Do(func() {
		_SLSBridgedTileSpaceTakeOwnershipOperationClass = SLSBridgedTileSpaceTakeOwnershipOperationClass{class: objc.GetClass("SLSBridgedTileSpaceTakeOwnershipOperation")}
	})
	return _SLSBridgedTileSpaceTakeOwnershipOperationClass
}

// GetSLSBridgedTileSpaceTakeOwnershipOperationClass returns the class object for SLSBridgedTileSpaceTakeOwnershipOperation.
func GetSLSBridgedTileSpaceTakeOwnershipOperationClass() SLSBridgedTileSpaceTakeOwnershipOperationClass {
	return getSLSBridgedTileSpaceTakeOwnershipOperationClass()
}

type SLSBridgedTileSpaceTakeOwnershipOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedTileSpaceTakeOwnershipOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedTileSpaceTakeOwnershipOperationClass) Alloc() SLSBridgedTileSpaceTakeOwnershipOperation {
	rv := objc.Send[SLSBridgedTileSpaceTakeOwnershipOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedTileSpaceTakeOwnershipOperation.SpaceID]
//   - [SLSBridgedTileSpaceTakeOwnershipOperation.InitWithSpaceID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedTileSpaceTakeOwnershipOperation
type SLSBridgedTileSpaceTakeOwnershipOperation struct {
	SLSAsynchronousBridgedWindowManagementOperation
}

// SLSBridgedTileSpaceTakeOwnershipOperationFromID constructs a [SLSBridgedTileSpaceTakeOwnershipOperation] from an objc.ID.
func SLSBridgedTileSpaceTakeOwnershipOperationFromID(id objc.ID) SLSBridgedTileSpaceTakeOwnershipOperation {
	return SLSBridgedTileSpaceTakeOwnershipOperation{SLSAsynchronousBridgedWindowManagementOperation: SLSAsynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedTileSpaceTakeOwnershipOperation implements ISLSBridgedTileSpaceTakeOwnershipOperation.
var _ ISLSBridgedTileSpaceTakeOwnershipOperation = SLSBridgedTileSpaceTakeOwnershipOperation{}

// An interface definition for the [SLSBridgedTileSpaceTakeOwnershipOperation] class.
//
// # Methods
//
//   - [ISLSBridgedTileSpaceTakeOwnershipOperation.SpaceID]
//   - [ISLSBridgedTileSpaceTakeOwnershipOperation.InitWithSpaceID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedTileSpaceTakeOwnershipOperation
type ISLSBridgedTileSpaceTakeOwnershipOperation interface {
	ISLSAsynchronousBridgedWindowManagementOperation

	// Topic: Methods

	SpaceID() uint64
	InitWithSpaceID(id uint64) SLSBridgedTileSpaceTakeOwnershipOperation
}

// Init initializes the instance.
func (s SLSBridgedTileSpaceTakeOwnershipOperation) Init() SLSBridgedTileSpaceTakeOwnershipOperation {
	rv := objc.Send[SLSBridgedTileSpaceTakeOwnershipOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedTileSpaceTakeOwnershipOperation) Autorelease() SLSBridgedTileSpaceTakeOwnershipOperation {
	rv := objc.Send[SLSBridgedTileSpaceTakeOwnershipOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedTileSpaceTakeOwnershipOperation creates a new SLSBridgedTileSpaceTakeOwnershipOperation instance.
func NewSLSBridgedTileSpaceTakeOwnershipOperation() SLSBridgedTileSpaceTakeOwnershipOperation {
	class := getSLSBridgedTileSpaceTakeOwnershipOperationClass()
	rv := objc.Send[SLSBridgedTileSpaceTakeOwnershipOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedTileSpaceTakeOwnershipOperation/initWithCoder:
func NewSLSBridgedTileSpaceTakeOwnershipOperationWithCoder(coder objectivec.IObject) SLSBridgedTileSpaceTakeOwnershipOperation {
	instance := getSLSBridgedTileSpaceTakeOwnershipOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedTileSpaceTakeOwnershipOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedTileSpaceTakeOwnershipOperation/initWithSpaceID:
func NewSLSBridgedTileSpaceTakeOwnershipOperationWithSpaceID(id uint64) SLSBridgedTileSpaceTakeOwnershipOperation {
	instance := getSLSBridgedTileSpaceTakeOwnershipOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpaceID:"), id)
	return SLSBridgedTileSpaceTakeOwnershipOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedTileSpaceTakeOwnershipOperation/initWithSpaceID:
func (s SLSBridgedTileSpaceTakeOwnershipOperation) InitWithSpaceID(id uint64) SLSBridgedTileSpaceTakeOwnershipOperation {
	rv := objc.Send[SLSBridgedTileSpaceTakeOwnershipOperation](s.ID, objc.Sel("initWithSpaceID:"), id)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedTileSpaceTakeOwnershipOperation/spaceID
func (s SLSBridgedTileSpaceTakeOwnershipOperation) SpaceID() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}
