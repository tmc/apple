// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedGetSpaceNeedsSafeApertureOperation] class.
var (
	_SLSBridgedGetSpaceNeedsSafeApertureOperationClass     SLSBridgedGetSpaceNeedsSafeApertureOperationClass
	_SLSBridgedGetSpaceNeedsSafeApertureOperationClassOnce sync.Once
)

func getSLSBridgedGetSpaceNeedsSafeApertureOperationClass() SLSBridgedGetSpaceNeedsSafeApertureOperationClass {
	_SLSBridgedGetSpaceNeedsSafeApertureOperationClassOnce.Do(func() {
		_SLSBridgedGetSpaceNeedsSafeApertureOperationClass = SLSBridgedGetSpaceNeedsSafeApertureOperationClass{class: objc.GetClass("SLSBridgedGetSpaceNeedsSafeApertureOperation")}
	})
	return _SLSBridgedGetSpaceNeedsSafeApertureOperationClass
}

// GetSLSBridgedGetSpaceNeedsSafeApertureOperationClass returns the class object for SLSBridgedGetSpaceNeedsSafeApertureOperation.
func GetSLSBridgedGetSpaceNeedsSafeApertureOperationClass() SLSBridgedGetSpaceNeedsSafeApertureOperationClass {
	return getSLSBridgedGetSpaceNeedsSafeApertureOperationClass()
}

type SLSBridgedGetSpaceNeedsSafeApertureOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedGetSpaceNeedsSafeApertureOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedGetSpaceNeedsSafeApertureOperationClass) Alloc() SLSBridgedGetSpaceNeedsSafeApertureOperation {
	rv := objc.Send[SLSBridgedGetSpaceNeedsSafeApertureOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedGetSpaceNeedsSafeApertureOperation.MakeResultWithBoolValue]
//   - [SLSBridgedGetSpaceNeedsSafeApertureOperation.SpaceID]
//   - [SLSBridgedGetSpaceNeedsSafeApertureOperation.InitWithSpaceID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedGetSpaceNeedsSafeApertureOperation
type SLSBridgedGetSpaceNeedsSafeApertureOperation struct {
	SLSSynchronousBridgedWindowManagementOperation
}

// SLSBridgedGetSpaceNeedsSafeApertureOperationFromID constructs a [SLSBridgedGetSpaceNeedsSafeApertureOperation] from an objc.ID.
func SLSBridgedGetSpaceNeedsSafeApertureOperationFromID(id objc.ID) SLSBridgedGetSpaceNeedsSafeApertureOperation {
	return SLSBridgedGetSpaceNeedsSafeApertureOperation{SLSSynchronousBridgedWindowManagementOperation: SLSSynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedGetSpaceNeedsSafeApertureOperation implements ISLSBridgedGetSpaceNeedsSafeApertureOperation.
var _ ISLSBridgedGetSpaceNeedsSafeApertureOperation = SLSBridgedGetSpaceNeedsSafeApertureOperation{}

// An interface definition for the [SLSBridgedGetSpaceNeedsSafeApertureOperation] class.
//
// # Methods
//
//   - [ISLSBridgedGetSpaceNeedsSafeApertureOperation.MakeResultWithBoolValue]
//   - [ISLSBridgedGetSpaceNeedsSafeApertureOperation.SpaceID]
//   - [ISLSBridgedGetSpaceNeedsSafeApertureOperation.InitWithSpaceID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedGetSpaceNeedsSafeApertureOperation
type ISLSBridgedGetSpaceNeedsSafeApertureOperation interface {
	ISLSSynchronousBridgedWindowManagementOperation

	// Topic: Methods

	MakeResultWithBoolValue(value bool) objectivec.IObject
	SpaceID() uint64
	InitWithSpaceID(id uint64) SLSBridgedGetSpaceNeedsSafeApertureOperation
}

// Init initializes the instance.
func (s SLSBridgedGetSpaceNeedsSafeApertureOperation) Init() SLSBridgedGetSpaceNeedsSafeApertureOperation {
	rv := objc.Send[SLSBridgedGetSpaceNeedsSafeApertureOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedGetSpaceNeedsSafeApertureOperation) Autorelease() SLSBridgedGetSpaceNeedsSafeApertureOperation {
	rv := objc.Send[SLSBridgedGetSpaceNeedsSafeApertureOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedGetSpaceNeedsSafeApertureOperation creates a new SLSBridgedGetSpaceNeedsSafeApertureOperation instance.
func NewSLSBridgedGetSpaceNeedsSafeApertureOperation() SLSBridgedGetSpaceNeedsSafeApertureOperation {
	class := getSLSBridgedGetSpaceNeedsSafeApertureOperationClass()
	rv := objc.Send[SLSBridgedGetSpaceNeedsSafeApertureOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedGetSpaceNeedsSafeApertureOperation/initWithCoder:
func NewSLSBridgedGetSpaceNeedsSafeApertureOperationWithCoder(coder objectivec.IObject) SLSBridgedGetSpaceNeedsSafeApertureOperation {
	instance := getSLSBridgedGetSpaceNeedsSafeApertureOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedGetSpaceNeedsSafeApertureOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedGetSpaceNeedsSafeApertureOperation/initWithSpaceID:
func NewSLSBridgedGetSpaceNeedsSafeApertureOperationWithSpaceID(id uint64) SLSBridgedGetSpaceNeedsSafeApertureOperation {
	instance := getSLSBridgedGetSpaceNeedsSafeApertureOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpaceID:"), id)
	return SLSBridgedGetSpaceNeedsSafeApertureOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedGetSpaceNeedsSafeApertureOperation/makeResultWithBoolValue:
func (s SLSBridgedGetSpaceNeedsSafeApertureOperation) MakeResultWithBoolValue(value bool) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("makeResultWithBoolValue:"), value)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedGetSpaceNeedsSafeApertureOperation/initWithSpaceID:
func (s SLSBridgedGetSpaceNeedsSafeApertureOperation) InitWithSpaceID(id uint64) SLSBridgedGetSpaceNeedsSafeApertureOperation {
	rv := objc.Send[SLSBridgedGetSpaceNeedsSafeApertureOperation](s.ID, objc.Sel("initWithSpaceID:"), id)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedGetSpaceNeedsSafeApertureOperation/spaceID
func (s SLSBridgedGetSpaceNeedsSafeApertureOperation) SpaceID() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}
