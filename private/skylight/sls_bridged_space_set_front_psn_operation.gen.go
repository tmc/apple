// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedSpaceSetFrontPSNOperation] class.
var (
	_SLSBridgedSpaceSetFrontPSNOperationClass     SLSBridgedSpaceSetFrontPSNOperationClass
	_SLSBridgedSpaceSetFrontPSNOperationClassOnce sync.Once
)

func getSLSBridgedSpaceSetFrontPSNOperationClass() SLSBridgedSpaceSetFrontPSNOperationClass {
	_SLSBridgedSpaceSetFrontPSNOperationClassOnce.Do(func() {
		_SLSBridgedSpaceSetFrontPSNOperationClass = SLSBridgedSpaceSetFrontPSNOperationClass{class: objc.GetClass("SLSBridgedSpaceSetFrontPSNOperation")}
	})
	return _SLSBridgedSpaceSetFrontPSNOperationClass
}

// GetSLSBridgedSpaceSetFrontPSNOperationClass returns the class object for SLSBridgedSpaceSetFrontPSNOperation.
func GetSLSBridgedSpaceSetFrontPSNOperationClass() SLSBridgedSpaceSetFrontPSNOperationClass {
	return getSLSBridgedSpaceSetFrontPSNOperationClass()
}

type SLSBridgedSpaceSetFrontPSNOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedSpaceSetFrontPSNOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedSpaceSetFrontPSNOperationClass) Alloc() SLSBridgedSpaceSetFrontPSNOperation {
	rv := objc.Send[SLSBridgedSpaceSetFrontPSNOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedSpaceSetFrontPSNOperation.Psn]
//   - [SLSBridgedSpaceSetFrontPSNOperation.SpaceID]
//   - [SLSBridgedSpaceSetFrontPSNOperation.InitWithSpaceIDPsn]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetFrontPSNOperation
type SLSBridgedSpaceSetFrontPSNOperation struct {
	SLSAsynchronousBridgedWindowManagementOperation
}

// SLSBridgedSpaceSetFrontPSNOperationFromID constructs a [SLSBridgedSpaceSetFrontPSNOperation] from an objc.ID.
func SLSBridgedSpaceSetFrontPSNOperationFromID(id objc.ID) SLSBridgedSpaceSetFrontPSNOperation {
	return SLSBridgedSpaceSetFrontPSNOperation{SLSAsynchronousBridgedWindowManagementOperation: SLSAsynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedSpaceSetFrontPSNOperation implements ISLSBridgedSpaceSetFrontPSNOperation.
var _ ISLSBridgedSpaceSetFrontPSNOperation = SLSBridgedSpaceSetFrontPSNOperation{}

// An interface definition for the [SLSBridgedSpaceSetFrontPSNOperation] class.
//
// # Methods
//
//   - [ISLSBridgedSpaceSetFrontPSNOperation.Psn]
//   - [ISLSBridgedSpaceSetFrontPSNOperation.SpaceID]
//   - [ISLSBridgedSpaceSetFrontPSNOperation.InitWithSpaceIDPsn]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetFrontPSNOperation
type ISLSBridgedSpaceSetFrontPSNOperation interface {
	ISLSAsynchronousBridgedWindowManagementOperation

	// Topic: Methods

	Psn() objectivec.IObject
	SpaceID() uint64
	InitWithSpaceIDPsn(id uint64, psn objectivec.IObject) SLSBridgedSpaceSetFrontPSNOperation
}

// Init initializes the instance.
func (s SLSBridgedSpaceSetFrontPSNOperation) Init() SLSBridgedSpaceSetFrontPSNOperation {
	rv := objc.Send[SLSBridgedSpaceSetFrontPSNOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedSpaceSetFrontPSNOperation) Autorelease() SLSBridgedSpaceSetFrontPSNOperation {
	rv := objc.Send[SLSBridgedSpaceSetFrontPSNOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedSpaceSetFrontPSNOperation creates a new SLSBridgedSpaceSetFrontPSNOperation instance.
func NewSLSBridgedSpaceSetFrontPSNOperation() SLSBridgedSpaceSetFrontPSNOperation {
	class := getSLSBridgedSpaceSetFrontPSNOperationClass()
	rv := objc.Send[SLSBridgedSpaceSetFrontPSNOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetFrontPSNOperation/initWithCoder:
func NewSLSBridgedSpaceSetFrontPSNOperationWithCoder(coder objectivec.IObject) SLSBridgedSpaceSetFrontPSNOperation {
	instance := getSLSBridgedSpaceSetFrontPSNOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedSpaceSetFrontPSNOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetFrontPSNOperation/initWithSpaceID:psn:
func NewSLSBridgedSpaceSetFrontPSNOperationWithSpaceIDPsn(id uint64, psn objectivec.IObject) SLSBridgedSpaceSetFrontPSNOperation {
	instance := getSLSBridgedSpaceSetFrontPSNOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpaceID:psn:"), id, psn)
	return SLSBridgedSpaceSetFrontPSNOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetFrontPSNOperation/initWithSpaceID:psn:
func (s SLSBridgedSpaceSetFrontPSNOperation) InitWithSpaceIDPsn(id uint64, psn objectivec.IObject) SLSBridgedSpaceSetFrontPSNOperation {
	rv := objc.Send[SLSBridgedSpaceSetFrontPSNOperation](s.ID, objc.Sel("initWithSpaceID:psn:"), id, psn)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetFrontPSNOperation/psn
func (s SLSBridgedSpaceSetFrontPSNOperation) Psn() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("psn"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetFrontPSNOperation/spaceID
func (s SLSBridgedSpaceSetFrontPSNOperation) SpaceID() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}
