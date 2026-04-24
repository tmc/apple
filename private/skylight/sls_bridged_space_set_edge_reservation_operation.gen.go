// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedSpaceSetEdgeReservationOperation] class.
var (
	_SLSBridgedSpaceSetEdgeReservationOperationClass     SLSBridgedSpaceSetEdgeReservationOperationClass
	_SLSBridgedSpaceSetEdgeReservationOperationClassOnce sync.Once
)

func getSLSBridgedSpaceSetEdgeReservationOperationClass() SLSBridgedSpaceSetEdgeReservationOperationClass {
	_SLSBridgedSpaceSetEdgeReservationOperationClassOnce.Do(func() {
		_SLSBridgedSpaceSetEdgeReservationOperationClass = SLSBridgedSpaceSetEdgeReservationOperationClass{class: objc.GetClass("SLSBridgedSpaceSetEdgeReservationOperation")}
	})
	return _SLSBridgedSpaceSetEdgeReservationOperationClass
}

// GetSLSBridgedSpaceSetEdgeReservationOperationClass returns the class object for SLSBridgedSpaceSetEdgeReservationOperation.
func GetSLSBridgedSpaceSetEdgeReservationOperationClass() SLSBridgedSpaceSetEdgeReservationOperationClass {
	return getSLSBridgedSpaceSetEdgeReservationOperationClass()
}

type SLSBridgedSpaceSetEdgeReservationOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedSpaceSetEdgeReservationOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedSpaceSetEdgeReservationOperationClass) Alloc() SLSBridgedSpaceSetEdgeReservationOperation {
	rv := objc.Send[SLSBridgedSpaceSetEdgeReservationOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedSpaceSetEdgeReservationOperation.Bottom]
//   - [SLSBridgedSpaceSetEdgeReservationOperation.EdgeMask]
//   - [SLSBridgedSpaceSetEdgeReservationOperation.Left]
//   - [SLSBridgedSpaceSetEdgeReservationOperation.Right]
//   - [SLSBridgedSpaceSetEdgeReservationOperation.SpaceID]
//   - [SLSBridgedSpaceSetEdgeReservationOperation.Top]
//   - [SLSBridgedSpaceSetEdgeReservationOperation.InitWithSpaceIDEdgeMaskLeftRightTopBottom]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetEdgeReservationOperation
type SLSBridgedSpaceSetEdgeReservationOperation struct {
	SLSAsynchronousBridgedWindowManagementOperation
}

// SLSBridgedSpaceSetEdgeReservationOperationFromID constructs a [SLSBridgedSpaceSetEdgeReservationOperation] from an objc.ID.
func SLSBridgedSpaceSetEdgeReservationOperationFromID(id objc.ID) SLSBridgedSpaceSetEdgeReservationOperation {
	return SLSBridgedSpaceSetEdgeReservationOperation{SLSAsynchronousBridgedWindowManagementOperation: SLSAsynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedSpaceSetEdgeReservationOperation implements ISLSBridgedSpaceSetEdgeReservationOperation.
var _ ISLSBridgedSpaceSetEdgeReservationOperation = SLSBridgedSpaceSetEdgeReservationOperation{}

// An interface definition for the [SLSBridgedSpaceSetEdgeReservationOperation] class.
//
// # Methods
//
//   - [ISLSBridgedSpaceSetEdgeReservationOperation.Bottom]
//   - [ISLSBridgedSpaceSetEdgeReservationOperation.EdgeMask]
//   - [ISLSBridgedSpaceSetEdgeReservationOperation.Left]
//   - [ISLSBridgedSpaceSetEdgeReservationOperation.Right]
//   - [ISLSBridgedSpaceSetEdgeReservationOperation.SpaceID]
//   - [ISLSBridgedSpaceSetEdgeReservationOperation.Top]
//   - [ISLSBridgedSpaceSetEdgeReservationOperation.InitWithSpaceIDEdgeMaskLeftRightTopBottom]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetEdgeReservationOperation
type ISLSBridgedSpaceSetEdgeReservationOperation interface {
	ISLSAsynchronousBridgedWindowManagementOperation

	// Topic: Methods

	Bottom() float64
	EdgeMask() uint64
	Left() float64
	Right() float64
	SpaceID() uint64
	Top() float64
	InitWithSpaceIDEdgeMaskLeftRightTopBottom(id uint64, mask uint64, left float64, right float64, top float64, bottom float64) SLSBridgedSpaceSetEdgeReservationOperation
}

// Init initializes the instance.
func (s SLSBridgedSpaceSetEdgeReservationOperation) Init() SLSBridgedSpaceSetEdgeReservationOperation {
	rv := objc.Send[SLSBridgedSpaceSetEdgeReservationOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedSpaceSetEdgeReservationOperation) Autorelease() SLSBridgedSpaceSetEdgeReservationOperation {
	rv := objc.Send[SLSBridgedSpaceSetEdgeReservationOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedSpaceSetEdgeReservationOperation creates a new SLSBridgedSpaceSetEdgeReservationOperation instance.
func NewSLSBridgedSpaceSetEdgeReservationOperation() SLSBridgedSpaceSetEdgeReservationOperation {
	class := getSLSBridgedSpaceSetEdgeReservationOperationClass()
	rv := objc.Send[SLSBridgedSpaceSetEdgeReservationOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetEdgeReservationOperation/initWithCoder:
func NewSLSBridgedSpaceSetEdgeReservationOperationWithCoder(coder objectivec.IObject) SLSBridgedSpaceSetEdgeReservationOperation {
	instance := getSLSBridgedSpaceSetEdgeReservationOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedSpaceSetEdgeReservationOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetEdgeReservationOperation/initWithSpaceID:edgeMask:left:right:top:bottom:
func NewSLSBridgedSpaceSetEdgeReservationOperationWithSpaceIDEdgeMaskLeftRightTopBottom(id uint64, mask uint64, left float64, right float64, top float64, bottom float64) SLSBridgedSpaceSetEdgeReservationOperation {
	instance := getSLSBridgedSpaceSetEdgeReservationOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpaceID:edgeMask:left:right:top:bottom:"), id, mask, left, right, top, bottom)
	return SLSBridgedSpaceSetEdgeReservationOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetEdgeReservationOperation/initWithSpaceID:edgeMask:left:right:top:bottom:
func (s SLSBridgedSpaceSetEdgeReservationOperation) InitWithSpaceIDEdgeMaskLeftRightTopBottom(id uint64, mask uint64, left float64, right float64, top float64, bottom float64) SLSBridgedSpaceSetEdgeReservationOperation {
	rv := objc.Send[SLSBridgedSpaceSetEdgeReservationOperation](s.ID, objc.Sel("initWithSpaceID:edgeMask:left:right:top:bottom:"), id, mask, left, right, top, bottom)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetEdgeReservationOperation/bottom
func (s SLSBridgedSpaceSetEdgeReservationOperation) Bottom() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("bottom"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetEdgeReservationOperation/edgeMask
func (s SLSBridgedSpaceSetEdgeReservationOperation) EdgeMask() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("edgeMask"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetEdgeReservationOperation/left
func (s SLSBridgedSpaceSetEdgeReservationOperation) Left() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("left"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetEdgeReservationOperation/right
func (s SLSBridgedSpaceSetEdgeReservationOperation) Right() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("right"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetEdgeReservationOperation/spaceID
func (s SLSBridgedSpaceSetEdgeReservationOperation) SpaceID() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetEdgeReservationOperation/top
func (s SLSBridgedSpaceSetEdgeReservationOperation) Top() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("top"))
	return rv
}
