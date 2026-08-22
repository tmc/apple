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
	rv := objc.SendIfResponds[SLSBridgedSpaceSetOwnersOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedSpaceSetOwnersOperation.Owners]
//   - [SLSBridgedSpaceSetOwnersOperation.SpaceID]
//   - [SLSBridgedSpaceSetOwnersOperation.InitWithSpaceIDOwners]
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
type ISLSBridgedSpaceSetOwnersOperation interface {
	ISLSAsynchronousBridgedWindowManagementOperation

	// Topic: Methods

	Owners() foundation.INSArray
	SpaceID() uint64
	InitWithSpaceIDOwners(id uint64, owners objectivec.IObject) SLSBridgedSpaceSetOwnersOperation
}

// Init initializes the instance.
func (s SLSBridgedSpaceSetOwnersOperation) Init() SLSBridgedSpaceSetOwnersOperation {
	rv := objc.SendIfResponds[SLSBridgedSpaceSetOwnersOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedSpaceSetOwnersOperation) Autorelease() SLSBridgedSpaceSetOwnersOperation {
	rv := objc.SendIfResponds[SLSBridgedSpaceSetOwnersOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedSpaceSetOwnersOperation creates a new SLSBridgedSpaceSetOwnersOperation instance.
func NewSLSBridgedSpaceSetOwnersOperation() SLSBridgedSpaceSetOwnersOperation {
	class := getSLSBridgedSpaceSetOwnersOperationClass()
	rv := objc.SendIfResponds[SLSBridgedSpaceSetOwnersOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSLSBridgedSpaceSetOwnersOperationWithCoder(coder objectivec.IObject) SLSBridgedSpaceSetOwnersOperation {
	instance := getSLSBridgedSpaceSetOwnersOperationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedSpaceSetOwnersOperationFromID(rv)
}

func NewSLSBridgedSpaceSetOwnersOperationWithSpaceIDOwners(id uint64, owners objectivec.IObject) SLSBridgedSpaceSetOwnersOperation {
	instance := getSLSBridgedSpaceSetOwnersOperationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithSpaceID:owners:"), id, owners)
	return SLSBridgedSpaceSetOwnersOperationFromID(rv)
}

func (s SLSBridgedSpaceSetOwnersOperation) InitWithSpaceIDOwners(id uint64, owners objectivec.IObject) SLSBridgedSpaceSetOwnersOperation {
	rv := objc.SendIfResponds[SLSBridgedSpaceSetOwnersOperation](s.ID, objc.Sel("initWithSpaceID:owners:"), id, owners)
	return rv
}

func (s SLSBridgedSpaceSetOwnersOperation) Owners() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("owners"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (s SLSBridgedSpaceSetOwnersOperation) SpaceID() uint64 {
	rv := objc.SendIfResponds[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}
