// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation] class.
var (
	_SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperationClass     SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperationClass
	_SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperationClassOnce sync.Once
)

func getSLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperationClass() SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperationClass {
	_SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperationClassOnce.Do(func() {
		_SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperationClass = SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperationClass{class: objc.GetClass("SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation")}
	})
	return _SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperationClass
}

// GetSLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperationClass returns the class object for SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation.
func GetSLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperationClass() SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperationClass {
	return getSLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperationClass()
}

type SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperationClass) Alloc() SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation {
	rv := objc.Send[SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation.DrivingSpaceID]
//   - [SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation.FencePort]
//   - [SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation.HorizontalIndex]
//   - [SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation.Options]
//   - [SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation.Point]
//   - [SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation.VerticalIndex]
//   - [SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation.InitWithDrivingSpaceIDVerticalIndexHorizontalIndexPointOptionsFencePort]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation
type SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation struct {
	SLSAsynchronousBridgedWindowManagementOperation
}

// SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperationFromID constructs a [SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation] from an objc.ID.
func SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperationFromID(id objc.ID) SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation {
	return SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation{SLSAsynchronousBridgedWindowManagementOperation: SLSAsynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation implements ISLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation.
var _ ISLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation = SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation{}

// An interface definition for the [SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation] class.
//
// # Methods
//
//   - [ISLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation.DrivingSpaceID]
//   - [ISLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation.FencePort]
//   - [ISLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation.HorizontalIndex]
//   - [ISLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation.Options]
//   - [ISLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation.Point]
//   - [ISLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation.VerticalIndex]
//   - [ISLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation.InitWithDrivingSpaceIDVerticalIndexHorizontalIndexPointOptionsFencePort]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation
type ISLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation interface {
	ISLSAsynchronousBridgedWindowManagementOperation

	// Topic: Methods

	DrivingSpaceID() uint64
	FencePort() uint32
	HorizontalIndex() uint64
	Options() uint64
	Point() corefoundation.CGPoint
	VerticalIndex() uint64
	InitWithDrivingSpaceIDVerticalIndexHorizontalIndexPointOptionsFencePort(id uint64, index uint64, index2 uint64, point corefoundation.CGPoint, options uint64, port uint32) SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation
}

// Init initializes the instance.
func (s SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation) Init() SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation {
	rv := objc.Send[SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation) Autorelease() SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation {
	rv := objc.Send[SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation creates a new SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation instance.
func NewSLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation() SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation {
	class := getSLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperationClass()
	rv := objc.Send[SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation/initWithCoder:
func NewSLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperationWithCoder(coder objectivec.IObject) SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation {
	instance := getSLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation/initWithDrivingSpaceID:verticalIndex:horizontalIndex:point:options:fencePort:
func NewSLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperationWithDrivingSpaceIDVerticalIndexHorizontalIndexPointOptionsFencePort(id uint64, index uint64, index2 uint64, point corefoundation.CGPoint, options uint64, port uint32) SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation {
	instance := getSLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDrivingSpaceID:verticalIndex:horizontalIndex:point:options:fencePort:"), id, index, index2, point, options, port)
	return SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation/initWithDrivingSpaceID:verticalIndex:horizontalIndex:point:options:fencePort:
func (s SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation) InitWithDrivingSpaceIDVerticalIndexHorizontalIndexPointOptionsFencePort(id uint64, index uint64, index2 uint64, point corefoundation.CGPoint, options uint64, port uint32) SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation {
	rv := objc.Send[SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation](s.ID, objc.Sel("initWithDrivingSpaceID:verticalIndex:horizontalIndex:point:options:fencePort:"), id, index, index2, point, options, port)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation/drivingSpaceID
func (s SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation) DrivingSpaceID() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("drivingSpaceID"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation/fencePort
func (s SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation) FencePort() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("fencePort"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation/horizontalIndex
func (s SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation) HorizontalIndex() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("horizontalIndex"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation/options
func (s SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation) Options() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("options"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation/point
func (s SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation) Point() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](s.ID, objc.Sel("point"))
	return corefoundation.CGPoint(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation/verticalIndex
func (s SLSBridgedSpaceClientDrivenMoveSpacersToPointFencedOperation) VerticalIndex() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("verticalIndex"))
	return rv
}
