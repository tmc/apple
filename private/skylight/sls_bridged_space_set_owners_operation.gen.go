// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedSpaceSetOwnersOperation] class.
var (
	_SLSBridgedSpaceSetOwnersOperationClass     SLSBridgedSpaceSetOwnersOperationClass
	_SLSBridgedSpaceSetOwnersOperationClassOnce sync.Once
)

func getSLSBridgedSpaceSetOwnersOperationClass() SLSBridgedSpaceSetOwnersOperationClass {
	_SLSBridgedSpaceSetOwnersOperationClassOnce.Do(func() {
		_SLSBridgedSpaceSetOwnersOperationClass = SLSBridgedSpaceSetOwnersOperationClass{class: objc.GetClass("SLSBridgedSpaceSetOwnersOperation")}
	})
	return _SLSBridgedSpaceSetOwnersOperationClass
}

// GetSLSBridgedSpaceSetOwnersOperationClass returns the class object for SLSBridgedSpaceSetOwnersOperation.
func GetSLSBridgedSpaceSetOwnersOperationClass() SLSBridgedSpaceSetOwnersOperationClass {
	return getSLSBridgedSpaceSetOwnersOperationClass()
}

type SLSBridgedSpaceSetOwnersOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedSpaceSetOwnersOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedSpaceSetOwnersOperationClass) Alloc() SLSBridgedSpaceSetOwnersOperation {
	rv := objc.Send[SLSBridgedSpaceSetOwnersOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedSpaceSetOwnersOperation.Owners]
//   - [SLSBridgedSpaceSetOwnersOperation.SpaceID]
//   - [SLSBridgedSpaceSetOwnersOperation.InitWithSpaceIDOwners]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetOwnersOperation
type SLSBridgedSpaceSetOwnersOperation struct {
	SLSAsynchronousBridgedWindowManagementOperation
}

// SLSBridgedSpaceSetOwnersOperationFromID constructs a [SLSBridgedSpaceSetOwnersOperation] from an objc.ID.
func SLSBridgedSpaceSetOwnersOperationFromID(id objc.ID) SLSBridgedSpaceSetOwnersOperation {
	return SLSBridgedSpaceSetOwnersOperation{SLSAsynchronousBridgedWindowManagementOperation: SLSAsynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedSpaceSetOwnersOperation implements ISLSBridgedSpaceSetOwnersOperation.
var _ ISLSBridgedSpaceSetOwnersOperation = SLSBridgedSpaceSetOwnersOperation{}

// An interface definition for the [SLSBridgedSpaceSetOwnersOperation] class.
//
// # Methods
//
//   - [ISLSBridgedSpaceSetOwnersOperation.Owners]
//   - [ISLSBridgedSpaceSetOwnersOperation.SpaceID]
//   - [ISLSBridgedSpaceSetOwnersOperation.InitWithSpaceIDOwners]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetOwnersOperation
type ISLSBridgedSpaceSetOwnersOperation interface {
	ISLSAsynchronousBridgedWindowManagementOperation

	// Topic: Methods

	Owners() foundation.INSArray
	SpaceID() uint64
	InitWithSpaceIDOwners(id uint64, owners objectivec.IObject) SLSBridgedSpaceSetOwnersOperation
}

// Init initializes the instance.
func (s SLSBridgedSpaceSetOwnersOperation) Init() SLSBridgedSpaceSetOwnersOperation {
	rv := objc.Send[SLSBridgedSpaceSetOwnersOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedSpaceSetOwnersOperation) Autorelease() SLSBridgedSpaceSetOwnersOperation {
	rv := objc.Send[SLSBridgedSpaceSetOwnersOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedSpaceSetOwnersOperation creates a new SLSBridgedSpaceSetOwnersOperation instance.
func NewSLSBridgedSpaceSetOwnersOperation() SLSBridgedSpaceSetOwnersOperation {
	class := getSLSBridgedSpaceSetOwnersOperationClass()
	rv := objc.Send[SLSBridgedSpaceSetOwnersOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetOwnersOperation/initWithCoder:
func NewSLSBridgedSpaceSetOwnersOperationWithCoder(coder objectivec.IObject) SLSBridgedSpaceSetOwnersOperation {
	instance := getSLSBridgedSpaceSetOwnersOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedSpaceSetOwnersOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetOwnersOperation/initWithSpaceID:owners:
func NewSLSBridgedSpaceSetOwnersOperationWithSpaceIDOwners(id uint64, owners objectivec.IObject) SLSBridgedSpaceSetOwnersOperation {
	instance := getSLSBridgedSpaceSetOwnersOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpaceID:owners:"), id, owners)
	return SLSBridgedSpaceSetOwnersOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetOwnersOperation/initWithSpaceID:owners:
func (s SLSBridgedSpaceSetOwnersOperation) InitWithSpaceIDOwners(id uint64, owners objectivec.IObject) SLSBridgedSpaceSetOwnersOperation {
	rv := objc.Send[SLSBridgedSpaceSetOwnersOperation](s.ID, objc.Sel("initWithSpaceID:owners:"), id, owners)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetOwnersOperation/owners
func (s SLSBridgedSpaceSetOwnersOperation) Owners() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("owners"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetOwnersOperation/spaceID
func (s SLSBridgedSpaceSetOwnersOperation) SpaceID() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}
