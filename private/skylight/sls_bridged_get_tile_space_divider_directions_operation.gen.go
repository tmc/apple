// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedGetTileSpaceDividerDirectionsOperation] class.
var (
	_SLSBridgedGetTileSpaceDividerDirectionsOperationClass     SLSBridgedGetTileSpaceDividerDirectionsOperationClass
	_SLSBridgedGetTileSpaceDividerDirectionsOperationClassOnce sync.Once
)

func getSLSBridgedGetTileSpaceDividerDirectionsOperationClass() SLSBridgedGetTileSpaceDividerDirectionsOperationClass {
	_SLSBridgedGetTileSpaceDividerDirectionsOperationClassOnce.Do(func() {
		_SLSBridgedGetTileSpaceDividerDirectionsOperationClass = SLSBridgedGetTileSpaceDividerDirectionsOperationClass{class: objc.GetClass("SLSBridgedGetTileSpaceDividerDirectionsOperation")}
	})
	return _SLSBridgedGetTileSpaceDividerDirectionsOperationClass
}

// GetSLSBridgedGetTileSpaceDividerDirectionsOperationClass returns the class object for SLSBridgedGetTileSpaceDividerDirectionsOperation.
func GetSLSBridgedGetTileSpaceDividerDirectionsOperationClass() SLSBridgedGetTileSpaceDividerDirectionsOperationClass {
	return getSLSBridgedGetTileSpaceDividerDirectionsOperationClass()
}

type SLSBridgedGetTileSpaceDividerDirectionsOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedGetTileSpaceDividerDirectionsOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedGetTileSpaceDividerDirectionsOperationClass) Alloc() SLSBridgedGetTileSpaceDividerDirectionsOperation {
	rv := objc.Send[SLSBridgedGetTileSpaceDividerDirectionsOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedGetTileSpaceDividerDirectionsOperation.MakeResultWithSpaceResizeDirections]
//   - [SLSBridgedGetTileSpaceDividerDirectionsOperation.SpaceID]
//   - [SLSBridgedGetTileSpaceDividerDirectionsOperation.InitWithSpaceID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedGetTileSpaceDividerDirectionsOperation
type SLSBridgedGetTileSpaceDividerDirectionsOperation struct {
	SLSSynchronousBridgedWindowManagementOperation
}

// SLSBridgedGetTileSpaceDividerDirectionsOperationFromID constructs a [SLSBridgedGetTileSpaceDividerDirectionsOperation] from an objc.ID.
func SLSBridgedGetTileSpaceDividerDirectionsOperationFromID(id objc.ID) SLSBridgedGetTileSpaceDividerDirectionsOperation {
	return SLSBridgedGetTileSpaceDividerDirectionsOperation{SLSSynchronousBridgedWindowManagementOperation: SLSSynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedGetTileSpaceDividerDirectionsOperation implements ISLSBridgedGetTileSpaceDividerDirectionsOperation.
var _ ISLSBridgedGetTileSpaceDividerDirectionsOperation = SLSBridgedGetTileSpaceDividerDirectionsOperation{}

// An interface definition for the [SLSBridgedGetTileSpaceDividerDirectionsOperation] class.
//
// # Methods
//
//   - [ISLSBridgedGetTileSpaceDividerDirectionsOperation.MakeResultWithSpaceResizeDirections]
//   - [ISLSBridgedGetTileSpaceDividerDirectionsOperation.SpaceID]
//   - [ISLSBridgedGetTileSpaceDividerDirectionsOperation.InitWithSpaceID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedGetTileSpaceDividerDirectionsOperation
type ISLSBridgedGetTileSpaceDividerDirectionsOperation interface {
	ISLSSynchronousBridgedWindowManagementOperation

	// Topic: Methods

	MakeResultWithSpaceResizeDirections(directions uint64) objectivec.IObject
	SpaceID() uint64
	InitWithSpaceID(id uint64) SLSBridgedGetTileSpaceDividerDirectionsOperation
}

// Init initializes the instance.
func (s SLSBridgedGetTileSpaceDividerDirectionsOperation) Init() SLSBridgedGetTileSpaceDividerDirectionsOperation {
	rv := objc.Send[SLSBridgedGetTileSpaceDividerDirectionsOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedGetTileSpaceDividerDirectionsOperation) Autorelease() SLSBridgedGetTileSpaceDividerDirectionsOperation {
	rv := objc.Send[SLSBridgedGetTileSpaceDividerDirectionsOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedGetTileSpaceDividerDirectionsOperation creates a new SLSBridgedGetTileSpaceDividerDirectionsOperation instance.
func NewSLSBridgedGetTileSpaceDividerDirectionsOperation() SLSBridgedGetTileSpaceDividerDirectionsOperation {
	class := getSLSBridgedGetTileSpaceDividerDirectionsOperationClass()
	rv := objc.Send[SLSBridgedGetTileSpaceDividerDirectionsOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedGetTileSpaceDividerDirectionsOperation/initWithCoder:
func NewSLSBridgedGetTileSpaceDividerDirectionsOperationWithCoder(coder objectivec.IObject) SLSBridgedGetTileSpaceDividerDirectionsOperation {
	instance := getSLSBridgedGetTileSpaceDividerDirectionsOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedGetTileSpaceDividerDirectionsOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedGetTileSpaceDividerDirectionsOperation/initWithSpaceID:
func NewSLSBridgedGetTileSpaceDividerDirectionsOperationWithSpaceID(id uint64) SLSBridgedGetTileSpaceDividerDirectionsOperation {
	instance := getSLSBridgedGetTileSpaceDividerDirectionsOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpaceID:"), id)
	return SLSBridgedGetTileSpaceDividerDirectionsOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedGetTileSpaceDividerDirectionsOperation/makeResultWithSpaceResizeDirections:
func (s SLSBridgedGetTileSpaceDividerDirectionsOperation) MakeResultWithSpaceResizeDirections(directions uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("makeResultWithSpaceResizeDirections:"), directions)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedGetTileSpaceDividerDirectionsOperation/initWithSpaceID:
func (s SLSBridgedGetTileSpaceDividerDirectionsOperation) InitWithSpaceID(id uint64) SLSBridgedGetTileSpaceDividerDirectionsOperation {
	rv := objc.Send[SLSBridgedGetTileSpaceDividerDirectionsOperation](s.ID, objc.Sel("initWithSpaceID:"), id)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedGetTileSpaceDividerDirectionsOperation/spaceID
func (s SLSBridgedGetTileSpaceDividerDirectionsOperation) SpaceID() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}
