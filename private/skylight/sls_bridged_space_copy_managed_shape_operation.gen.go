// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedSpaceCopyManagedShapeOperation] class.
var (
	_SLSBridgedSpaceCopyManagedShapeOperationClass     SLSBridgedSpaceCopyManagedShapeOperationClass
	_SLSBridgedSpaceCopyManagedShapeOperationClassOnce sync.Once
)

func getSLSBridgedSpaceCopyManagedShapeOperationClass() SLSBridgedSpaceCopyManagedShapeOperationClass {
	_SLSBridgedSpaceCopyManagedShapeOperationClassOnce.Do(func() {
		_SLSBridgedSpaceCopyManagedShapeOperationClass = SLSBridgedSpaceCopyManagedShapeOperationClass{class: objc.GetClass("SLSBridgedSpaceCopyManagedShapeOperation")}
	})
	return _SLSBridgedSpaceCopyManagedShapeOperationClass
}

// GetSLSBridgedSpaceCopyManagedShapeOperationClass returns the class object for SLSBridgedSpaceCopyManagedShapeOperation.
func GetSLSBridgedSpaceCopyManagedShapeOperationClass() SLSBridgedSpaceCopyManagedShapeOperationClass {
	return getSLSBridgedSpaceCopyManagedShapeOperationClass()
}

type SLSBridgedSpaceCopyManagedShapeOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedSpaceCopyManagedShapeOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedSpaceCopyManagedShapeOperationClass) Alloc() SLSBridgedSpaceCopyManagedShapeOperation {
	rv := objc.Send[SLSBridgedSpaceCopyManagedShapeOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedSpaceCopyManagedShapeOperation.MakeResultWithRegion]
//   - [SLSBridgedSpaceCopyManagedShapeOperation.SpaceID]
//   - [SLSBridgedSpaceCopyManagedShapeOperation.InitWithSpaceID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCopyManagedShapeOperation
type SLSBridgedSpaceCopyManagedShapeOperation struct {
	SLSSynchronousBridgedWindowManagementOperation
}

// SLSBridgedSpaceCopyManagedShapeOperationFromID constructs a [SLSBridgedSpaceCopyManagedShapeOperation] from an objc.ID.
func SLSBridgedSpaceCopyManagedShapeOperationFromID(id objc.ID) SLSBridgedSpaceCopyManagedShapeOperation {
	return SLSBridgedSpaceCopyManagedShapeOperation{SLSSynchronousBridgedWindowManagementOperation: SLSSynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedSpaceCopyManagedShapeOperation implements ISLSBridgedSpaceCopyManagedShapeOperation.
var _ ISLSBridgedSpaceCopyManagedShapeOperation = SLSBridgedSpaceCopyManagedShapeOperation{}

// An interface definition for the [SLSBridgedSpaceCopyManagedShapeOperation] class.
//
// # Methods
//
//   - [ISLSBridgedSpaceCopyManagedShapeOperation.MakeResultWithRegion]
//   - [ISLSBridgedSpaceCopyManagedShapeOperation.SpaceID]
//   - [ISLSBridgedSpaceCopyManagedShapeOperation.InitWithSpaceID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCopyManagedShapeOperation
type ISLSBridgedSpaceCopyManagedShapeOperation interface {
	ISLSSynchronousBridgedWindowManagementOperation

	// Topic: Methods

	MakeResultWithRegion(region uintptr) objectivec.IObject
	SpaceID() uint64
	InitWithSpaceID(id uint64) SLSBridgedSpaceCopyManagedShapeOperation
}

// Init initializes the instance.
func (s SLSBridgedSpaceCopyManagedShapeOperation) Init() SLSBridgedSpaceCopyManagedShapeOperation {
	rv := objc.Send[SLSBridgedSpaceCopyManagedShapeOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedSpaceCopyManagedShapeOperation) Autorelease() SLSBridgedSpaceCopyManagedShapeOperation {
	rv := objc.Send[SLSBridgedSpaceCopyManagedShapeOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedSpaceCopyManagedShapeOperation creates a new SLSBridgedSpaceCopyManagedShapeOperation instance.
func NewSLSBridgedSpaceCopyManagedShapeOperation() SLSBridgedSpaceCopyManagedShapeOperation {
	class := getSLSBridgedSpaceCopyManagedShapeOperationClass()
	rv := objc.Send[SLSBridgedSpaceCopyManagedShapeOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCopyManagedShapeOperation/initWithCoder:
func NewSLSBridgedSpaceCopyManagedShapeOperationWithCoder(coder objectivec.IObject) SLSBridgedSpaceCopyManagedShapeOperation {
	instance := getSLSBridgedSpaceCopyManagedShapeOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedSpaceCopyManagedShapeOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCopyManagedShapeOperation/initWithSpaceID:
func NewSLSBridgedSpaceCopyManagedShapeOperationWithSpaceID(id uint64) SLSBridgedSpaceCopyManagedShapeOperation {
	instance := getSLSBridgedSpaceCopyManagedShapeOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpaceID:"), id)
	return SLSBridgedSpaceCopyManagedShapeOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCopyManagedShapeOperation/makeResultWithRegion:
func (s SLSBridgedSpaceCopyManagedShapeOperation) MakeResultWithRegion(region uintptr) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("makeResultWithRegion:"), region)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCopyManagedShapeOperation/initWithSpaceID:
func (s SLSBridgedSpaceCopyManagedShapeOperation) InitWithSpaceID(id uint64) SLSBridgedSpaceCopyManagedShapeOperation {
	rv := objc.Send[SLSBridgedSpaceCopyManagedShapeOperation](s.ID, objc.Sel("initWithSpaceID:"), id)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCopyManagedShapeOperation/spaceID
func (s SLSBridgedSpaceCopyManagedShapeOperation) SpaceID() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}
