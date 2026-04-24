// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedSpaceClientDrivenMoveSpacersToPointOperation] class.
var (
	_SLSBridgedSpaceClientDrivenMoveSpacersToPointOperationClass     SLSBridgedSpaceClientDrivenMoveSpacersToPointOperationClass
	_SLSBridgedSpaceClientDrivenMoveSpacersToPointOperationClassOnce sync.Once
)

func getSLSBridgedSpaceClientDrivenMoveSpacersToPointOperationClass() SLSBridgedSpaceClientDrivenMoveSpacersToPointOperationClass {
	_SLSBridgedSpaceClientDrivenMoveSpacersToPointOperationClassOnce.Do(func() {
		_SLSBridgedSpaceClientDrivenMoveSpacersToPointOperationClass = SLSBridgedSpaceClientDrivenMoveSpacersToPointOperationClass{class: objc.GetClass("SLSBridgedSpaceClientDrivenMoveSpacersToPointOperation")}
	})
	return _SLSBridgedSpaceClientDrivenMoveSpacersToPointOperationClass
}

// GetSLSBridgedSpaceClientDrivenMoveSpacersToPointOperationClass returns the class object for SLSBridgedSpaceClientDrivenMoveSpacersToPointOperation.
func GetSLSBridgedSpaceClientDrivenMoveSpacersToPointOperationClass() SLSBridgedSpaceClientDrivenMoveSpacersToPointOperationClass {
	return getSLSBridgedSpaceClientDrivenMoveSpacersToPointOperationClass()
}

type SLSBridgedSpaceClientDrivenMoveSpacersToPointOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedSpaceClientDrivenMoveSpacersToPointOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedSpaceClientDrivenMoveSpacersToPointOperationClass) Alloc() SLSBridgedSpaceClientDrivenMoveSpacersToPointOperation {
	rv := objc.Send[SLSBridgedSpaceClientDrivenMoveSpacersToPointOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedSpaceClientDrivenMoveSpacersToPointOperation.DrivingSpaceID]
//   - [SLSBridgedSpaceClientDrivenMoveSpacersToPointOperation.HorizontalIndex]
//   - [SLSBridgedSpaceClientDrivenMoveSpacersToPointOperation.Options]
//   - [SLSBridgedSpaceClientDrivenMoveSpacersToPointOperation.Point]
//   - [SLSBridgedSpaceClientDrivenMoveSpacersToPointOperation.VerticalIndex]
//   - [SLSBridgedSpaceClientDrivenMoveSpacersToPointOperation.InitWithDrivingSpaceIDVerticalIndexHorizontalIndexPointOptions]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceClientDrivenMoveSpacersToPointOperation
type SLSBridgedSpaceClientDrivenMoveSpacersToPointOperation struct {
	SLSAsynchronousBridgedWindowManagementOperation
}

// SLSBridgedSpaceClientDrivenMoveSpacersToPointOperationFromID constructs a [SLSBridgedSpaceClientDrivenMoveSpacersToPointOperation] from an objc.ID.
func SLSBridgedSpaceClientDrivenMoveSpacersToPointOperationFromID(id objc.ID) SLSBridgedSpaceClientDrivenMoveSpacersToPointOperation {
	return SLSBridgedSpaceClientDrivenMoveSpacersToPointOperation{SLSAsynchronousBridgedWindowManagementOperation: SLSAsynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedSpaceClientDrivenMoveSpacersToPointOperation implements ISLSBridgedSpaceClientDrivenMoveSpacersToPointOperation.
var _ ISLSBridgedSpaceClientDrivenMoveSpacersToPointOperation = SLSBridgedSpaceClientDrivenMoveSpacersToPointOperation{}

// An interface definition for the [SLSBridgedSpaceClientDrivenMoveSpacersToPointOperation] class.
//
// # Methods
//
//   - [ISLSBridgedSpaceClientDrivenMoveSpacersToPointOperation.DrivingSpaceID]
//   - [ISLSBridgedSpaceClientDrivenMoveSpacersToPointOperation.HorizontalIndex]
//   - [ISLSBridgedSpaceClientDrivenMoveSpacersToPointOperation.Options]
//   - [ISLSBridgedSpaceClientDrivenMoveSpacersToPointOperation.Point]
//   - [ISLSBridgedSpaceClientDrivenMoveSpacersToPointOperation.VerticalIndex]
//   - [ISLSBridgedSpaceClientDrivenMoveSpacersToPointOperation.InitWithDrivingSpaceIDVerticalIndexHorizontalIndexPointOptions]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceClientDrivenMoveSpacersToPointOperation
type ISLSBridgedSpaceClientDrivenMoveSpacersToPointOperation interface {
	ISLSAsynchronousBridgedWindowManagementOperation

	// Topic: Methods

	DrivingSpaceID() uint64
	HorizontalIndex() uint64
	Options() uint64
	Point() corefoundation.CGPoint
	VerticalIndex() uint64
	InitWithDrivingSpaceIDVerticalIndexHorizontalIndexPointOptions(id uint64, index uint64, index2 uint64, point corefoundation.CGPoint, options uint64) SLSBridgedSpaceClientDrivenMoveSpacersToPointOperation
}

// Init initializes the instance.
func (s SLSBridgedSpaceClientDrivenMoveSpacersToPointOperation) Init() SLSBridgedSpaceClientDrivenMoveSpacersToPointOperation {
	rv := objc.Send[SLSBridgedSpaceClientDrivenMoveSpacersToPointOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedSpaceClientDrivenMoveSpacersToPointOperation) Autorelease() SLSBridgedSpaceClientDrivenMoveSpacersToPointOperation {
	rv := objc.Send[SLSBridgedSpaceClientDrivenMoveSpacersToPointOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedSpaceClientDrivenMoveSpacersToPointOperation creates a new SLSBridgedSpaceClientDrivenMoveSpacersToPointOperation instance.
func NewSLSBridgedSpaceClientDrivenMoveSpacersToPointOperation() SLSBridgedSpaceClientDrivenMoveSpacersToPointOperation {
	class := getSLSBridgedSpaceClientDrivenMoveSpacersToPointOperationClass()
	rv := objc.Send[SLSBridgedSpaceClientDrivenMoveSpacersToPointOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceClientDrivenMoveSpacersToPointOperation/initWithCoder:
func NewSLSBridgedSpaceClientDrivenMoveSpacersToPointOperationWithCoder(coder objectivec.IObject) SLSBridgedSpaceClientDrivenMoveSpacersToPointOperation {
	instance := getSLSBridgedSpaceClientDrivenMoveSpacersToPointOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedSpaceClientDrivenMoveSpacersToPointOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceClientDrivenMoveSpacersToPointOperation/initWithDrivingSpaceID:verticalIndex:horizontalIndex:point:options:
func NewSLSBridgedSpaceClientDrivenMoveSpacersToPointOperationWithDrivingSpaceIDVerticalIndexHorizontalIndexPointOptions(id uint64, index uint64, index2 uint64, point corefoundation.CGPoint, options uint64) SLSBridgedSpaceClientDrivenMoveSpacersToPointOperation {
	instance := getSLSBridgedSpaceClientDrivenMoveSpacersToPointOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDrivingSpaceID:verticalIndex:horizontalIndex:point:options:"), id, index, index2, point, options)
	return SLSBridgedSpaceClientDrivenMoveSpacersToPointOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceClientDrivenMoveSpacersToPointOperation/initWithDrivingSpaceID:verticalIndex:horizontalIndex:point:options:
func (s SLSBridgedSpaceClientDrivenMoveSpacersToPointOperation) InitWithDrivingSpaceIDVerticalIndexHorizontalIndexPointOptions(id uint64, index uint64, index2 uint64, point corefoundation.CGPoint, options uint64) SLSBridgedSpaceClientDrivenMoveSpacersToPointOperation {
	rv := objc.Send[SLSBridgedSpaceClientDrivenMoveSpacersToPointOperation](s.ID, objc.Sel("initWithDrivingSpaceID:verticalIndex:horizontalIndex:point:options:"), id, index, index2, point, options)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceClientDrivenMoveSpacersToPointOperation/drivingSpaceID
func (s SLSBridgedSpaceClientDrivenMoveSpacersToPointOperation) DrivingSpaceID() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("drivingSpaceID"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceClientDrivenMoveSpacersToPointOperation/horizontalIndex
func (s SLSBridgedSpaceClientDrivenMoveSpacersToPointOperation) HorizontalIndex() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("horizontalIndex"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceClientDrivenMoveSpacersToPointOperation/options
func (s SLSBridgedSpaceClientDrivenMoveSpacersToPointOperation) Options() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("options"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceClientDrivenMoveSpacersToPointOperation/point
func (s SLSBridgedSpaceClientDrivenMoveSpacersToPointOperation) Point() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](s.ID, objc.Sel("point"))
	return corefoundation.CGPoint(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceClientDrivenMoveSpacersToPointOperation/verticalIndex
func (s SLSBridgedSpaceClientDrivenMoveSpacersToPointOperation) VerticalIndex() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("verticalIndex"))
	return rv
}
