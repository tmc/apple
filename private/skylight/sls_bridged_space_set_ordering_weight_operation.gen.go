// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedSpaceSetOrderingWeightOperation] class.
var (
	_SLSBridgedSpaceSetOrderingWeightOperationClass     SLSBridgedSpaceSetOrderingWeightOperationClass
	_SLSBridgedSpaceSetOrderingWeightOperationClassOnce sync.Once
)

func getSLSBridgedSpaceSetOrderingWeightOperationClass() SLSBridgedSpaceSetOrderingWeightOperationClass {
	_SLSBridgedSpaceSetOrderingWeightOperationClassOnce.Do(func() {
		_SLSBridgedSpaceSetOrderingWeightOperationClass = SLSBridgedSpaceSetOrderingWeightOperationClass{class: objc.GetClass("SLSBridgedSpaceSetOrderingWeightOperation")}
	})
	return _SLSBridgedSpaceSetOrderingWeightOperationClass
}

// GetSLSBridgedSpaceSetOrderingWeightOperationClass returns the class object for SLSBridgedSpaceSetOrderingWeightOperation.
func GetSLSBridgedSpaceSetOrderingWeightOperationClass() SLSBridgedSpaceSetOrderingWeightOperationClass {
	return getSLSBridgedSpaceSetOrderingWeightOperationClass()
}

type SLSBridgedSpaceSetOrderingWeightOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedSpaceSetOrderingWeightOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedSpaceSetOrderingWeightOperationClass) Alloc() SLSBridgedSpaceSetOrderingWeightOperation {
	rv := objc.Send[SLSBridgedSpaceSetOrderingWeightOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedSpaceSetOrderingWeightOperation.SpaceID]
//   - [SLSBridgedSpaceSetOrderingWeightOperation.Weight]
//   - [SLSBridgedSpaceSetOrderingWeightOperation.InitWithSpaceIDWeight]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetOrderingWeightOperation
type SLSBridgedSpaceSetOrderingWeightOperation struct {
	SLSAsynchronousBridgedWindowManagementOperation
}

// SLSBridgedSpaceSetOrderingWeightOperationFromID constructs a [SLSBridgedSpaceSetOrderingWeightOperation] from an objc.ID.
func SLSBridgedSpaceSetOrderingWeightOperationFromID(id objc.ID) SLSBridgedSpaceSetOrderingWeightOperation {
	return SLSBridgedSpaceSetOrderingWeightOperation{SLSAsynchronousBridgedWindowManagementOperation: SLSAsynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedSpaceSetOrderingWeightOperation implements ISLSBridgedSpaceSetOrderingWeightOperation.
var _ ISLSBridgedSpaceSetOrderingWeightOperation = SLSBridgedSpaceSetOrderingWeightOperation{}

// An interface definition for the [SLSBridgedSpaceSetOrderingWeightOperation] class.
//
// # Methods
//
//   - [ISLSBridgedSpaceSetOrderingWeightOperation.SpaceID]
//   - [ISLSBridgedSpaceSetOrderingWeightOperation.Weight]
//   - [ISLSBridgedSpaceSetOrderingWeightOperation.InitWithSpaceIDWeight]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetOrderingWeightOperation
type ISLSBridgedSpaceSetOrderingWeightOperation interface {
	ISLSAsynchronousBridgedWindowManagementOperation

	// Topic: Methods

	SpaceID() uint64
	Weight() int
	InitWithSpaceIDWeight(id uint64, weight int) SLSBridgedSpaceSetOrderingWeightOperation
}

// Init initializes the instance.
func (s SLSBridgedSpaceSetOrderingWeightOperation) Init() SLSBridgedSpaceSetOrderingWeightOperation {
	rv := objc.Send[SLSBridgedSpaceSetOrderingWeightOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedSpaceSetOrderingWeightOperation) Autorelease() SLSBridgedSpaceSetOrderingWeightOperation {
	rv := objc.Send[SLSBridgedSpaceSetOrderingWeightOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedSpaceSetOrderingWeightOperation creates a new SLSBridgedSpaceSetOrderingWeightOperation instance.
func NewSLSBridgedSpaceSetOrderingWeightOperation() SLSBridgedSpaceSetOrderingWeightOperation {
	class := getSLSBridgedSpaceSetOrderingWeightOperationClass()
	rv := objc.Send[SLSBridgedSpaceSetOrderingWeightOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetOrderingWeightOperation/initWithCoder:
func NewSLSBridgedSpaceSetOrderingWeightOperationWithCoder(coder objectivec.IObject) SLSBridgedSpaceSetOrderingWeightOperation {
	instance := getSLSBridgedSpaceSetOrderingWeightOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedSpaceSetOrderingWeightOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetOrderingWeightOperation/initWithSpaceID:weight:
func NewSLSBridgedSpaceSetOrderingWeightOperationWithSpaceIDWeight(id uint64, weight int) SLSBridgedSpaceSetOrderingWeightOperation {
	instance := getSLSBridgedSpaceSetOrderingWeightOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpaceID:weight:"), id, weight)
	return SLSBridgedSpaceSetOrderingWeightOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetOrderingWeightOperation/initWithSpaceID:weight:
func (s SLSBridgedSpaceSetOrderingWeightOperation) InitWithSpaceIDWeight(id uint64, weight int) SLSBridgedSpaceSetOrderingWeightOperation {
	rv := objc.Send[SLSBridgedSpaceSetOrderingWeightOperation](s.ID, objc.Sel("initWithSpaceID:weight:"), id, weight)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetOrderingWeightOperation/spaceID
func (s SLSBridgedSpaceSetOrderingWeightOperation) SpaceID() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetOrderingWeightOperation/weight
func (s SLSBridgedSpaceSetOrderingWeightOperation) Weight() int {
	rv := objc.Send[int](s.ID, objc.Sel("weight"))
	return rv
}
