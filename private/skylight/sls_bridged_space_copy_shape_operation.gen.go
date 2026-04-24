// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedSpaceCopyShapeOperation] class.
var (
	_SLSBridgedSpaceCopyShapeOperationClass     SLSBridgedSpaceCopyShapeOperationClass
	_SLSBridgedSpaceCopyShapeOperationClassOnce sync.Once
)

func getSLSBridgedSpaceCopyShapeOperationClass() SLSBridgedSpaceCopyShapeOperationClass {
	_SLSBridgedSpaceCopyShapeOperationClassOnce.Do(func() {
		_SLSBridgedSpaceCopyShapeOperationClass = SLSBridgedSpaceCopyShapeOperationClass{class: objc.GetClass("SLSBridgedSpaceCopyShapeOperation")}
	})
	return _SLSBridgedSpaceCopyShapeOperationClass
}

// GetSLSBridgedSpaceCopyShapeOperationClass returns the class object for SLSBridgedSpaceCopyShapeOperation.
func GetSLSBridgedSpaceCopyShapeOperationClass() SLSBridgedSpaceCopyShapeOperationClass {
	return getSLSBridgedSpaceCopyShapeOperationClass()
}

type SLSBridgedSpaceCopyShapeOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedSpaceCopyShapeOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedSpaceCopyShapeOperationClass) Alloc() SLSBridgedSpaceCopyShapeOperation {
	rv := objc.Send[SLSBridgedSpaceCopyShapeOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedSpaceCopyShapeOperation.MakeResultWithRegion]
//   - [SLSBridgedSpaceCopyShapeOperation.SpaceID]
//   - [SLSBridgedSpaceCopyShapeOperation.InitWithSpaceID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCopyShapeOperation
type SLSBridgedSpaceCopyShapeOperation struct {
	SLSSynchronousBridgedWindowManagementOperation
}

// SLSBridgedSpaceCopyShapeOperationFromID constructs a [SLSBridgedSpaceCopyShapeOperation] from an objc.ID.
func SLSBridgedSpaceCopyShapeOperationFromID(id objc.ID) SLSBridgedSpaceCopyShapeOperation {
	return SLSBridgedSpaceCopyShapeOperation{SLSSynchronousBridgedWindowManagementOperation: SLSSynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedSpaceCopyShapeOperation implements ISLSBridgedSpaceCopyShapeOperation.
var _ ISLSBridgedSpaceCopyShapeOperation = SLSBridgedSpaceCopyShapeOperation{}

// An interface definition for the [SLSBridgedSpaceCopyShapeOperation] class.
//
// # Methods
//
//   - [ISLSBridgedSpaceCopyShapeOperation.MakeResultWithRegion]
//   - [ISLSBridgedSpaceCopyShapeOperation.SpaceID]
//   - [ISLSBridgedSpaceCopyShapeOperation.InitWithSpaceID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCopyShapeOperation
type ISLSBridgedSpaceCopyShapeOperation interface {
	ISLSSynchronousBridgedWindowManagementOperation

	// Topic: Methods

	MakeResultWithRegion(region uintptr) objectivec.IObject
	SpaceID() uint64
	InitWithSpaceID(id uint64) SLSBridgedSpaceCopyShapeOperation
}

// Init initializes the instance.
func (s SLSBridgedSpaceCopyShapeOperation) Init() SLSBridgedSpaceCopyShapeOperation {
	rv := objc.Send[SLSBridgedSpaceCopyShapeOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedSpaceCopyShapeOperation) Autorelease() SLSBridgedSpaceCopyShapeOperation {
	rv := objc.Send[SLSBridgedSpaceCopyShapeOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedSpaceCopyShapeOperation creates a new SLSBridgedSpaceCopyShapeOperation instance.
func NewSLSBridgedSpaceCopyShapeOperation() SLSBridgedSpaceCopyShapeOperation {
	class := getSLSBridgedSpaceCopyShapeOperationClass()
	rv := objc.Send[SLSBridgedSpaceCopyShapeOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCopyShapeOperation/initWithCoder:
func NewSLSBridgedSpaceCopyShapeOperationWithCoder(coder objectivec.IObject) SLSBridgedSpaceCopyShapeOperation {
	instance := getSLSBridgedSpaceCopyShapeOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedSpaceCopyShapeOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCopyShapeOperation/initWithSpaceID:
func NewSLSBridgedSpaceCopyShapeOperationWithSpaceID(id uint64) SLSBridgedSpaceCopyShapeOperation {
	instance := getSLSBridgedSpaceCopyShapeOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpaceID:"), id)
	return SLSBridgedSpaceCopyShapeOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCopyShapeOperation/makeResultWithRegion:
func (s SLSBridgedSpaceCopyShapeOperation) MakeResultWithRegion(region uintptr) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("makeResultWithRegion:"), region)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCopyShapeOperation/initWithSpaceID:
func (s SLSBridgedSpaceCopyShapeOperation) InitWithSpaceID(id uint64) SLSBridgedSpaceCopyShapeOperation {
	rv := objc.Send[SLSBridgedSpaceCopyShapeOperation](s.ID, objc.Sel("initWithSpaceID:"), id)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCopyShapeOperation/spaceID
func (s SLSBridgedSpaceCopyShapeOperation) SpaceID() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}
