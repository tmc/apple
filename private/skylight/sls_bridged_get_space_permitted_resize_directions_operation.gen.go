// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedGetSpacePermittedResizeDirectionsOperation] class.
var (
	_SLSBridgedGetSpacePermittedResizeDirectionsOperationClass     SLSBridgedGetSpacePermittedResizeDirectionsOperationClass
	_SLSBridgedGetSpacePermittedResizeDirectionsOperationClassOnce sync.Once
)

func getSLSBridgedGetSpacePermittedResizeDirectionsOperationClass() SLSBridgedGetSpacePermittedResizeDirectionsOperationClass {
	_SLSBridgedGetSpacePermittedResizeDirectionsOperationClassOnce.Do(func() {
		_SLSBridgedGetSpacePermittedResizeDirectionsOperationClass = SLSBridgedGetSpacePermittedResizeDirectionsOperationClass{class: objc.GetClass("SLSBridgedGetSpacePermittedResizeDirectionsOperation")}
	})
	return _SLSBridgedGetSpacePermittedResizeDirectionsOperationClass
}

// GetSLSBridgedGetSpacePermittedResizeDirectionsOperationClass returns the class object for SLSBridgedGetSpacePermittedResizeDirectionsOperation.
func GetSLSBridgedGetSpacePermittedResizeDirectionsOperationClass() SLSBridgedGetSpacePermittedResizeDirectionsOperationClass {
	return getSLSBridgedGetSpacePermittedResizeDirectionsOperationClass()
}

type SLSBridgedGetSpacePermittedResizeDirectionsOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedGetSpacePermittedResizeDirectionsOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedGetSpacePermittedResizeDirectionsOperationClass) Alloc() SLSBridgedGetSpacePermittedResizeDirectionsOperation {
	rv := objc.Send[SLSBridgedGetSpacePermittedResizeDirectionsOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedGetSpacePermittedResizeDirectionsOperation.HorizontalIndex]
//   - [SLSBridgedGetSpacePermittedResizeDirectionsOperation.MakeResultWithSpaceResizeDirections]
//   - [SLSBridgedGetSpacePermittedResizeDirectionsOperation.SpaceID]
//   - [SLSBridgedGetSpacePermittedResizeDirectionsOperation.VerticalIndex]
//   - [SLSBridgedGetSpacePermittedResizeDirectionsOperation.InitWithSpaceIDVerticalIndexHorizontalIndex]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedGetSpacePermittedResizeDirectionsOperation
type SLSBridgedGetSpacePermittedResizeDirectionsOperation struct {
	SLSSynchronousBridgedWindowManagementOperation
}

// SLSBridgedGetSpacePermittedResizeDirectionsOperationFromID constructs a [SLSBridgedGetSpacePermittedResizeDirectionsOperation] from an objc.ID.
func SLSBridgedGetSpacePermittedResizeDirectionsOperationFromID(id objc.ID) SLSBridgedGetSpacePermittedResizeDirectionsOperation {
	return SLSBridgedGetSpacePermittedResizeDirectionsOperation{SLSSynchronousBridgedWindowManagementOperation: SLSSynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedGetSpacePermittedResizeDirectionsOperation implements ISLSBridgedGetSpacePermittedResizeDirectionsOperation.
var _ ISLSBridgedGetSpacePermittedResizeDirectionsOperation = SLSBridgedGetSpacePermittedResizeDirectionsOperation{}

// An interface definition for the [SLSBridgedGetSpacePermittedResizeDirectionsOperation] class.
//
// # Methods
//
//   - [ISLSBridgedGetSpacePermittedResizeDirectionsOperation.HorizontalIndex]
//   - [ISLSBridgedGetSpacePermittedResizeDirectionsOperation.MakeResultWithSpaceResizeDirections]
//   - [ISLSBridgedGetSpacePermittedResizeDirectionsOperation.SpaceID]
//   - [ISLSBridgedGetSpacePermittedResizeDirectionsOperation.VerticalIndex]
//   - [ISLSBridgedGetSpacePermittedResizeDirectionsOperation.InitWithSpaceIDVerticalIndexHorizontalIndex]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedGetSpacePermittedResizeDirectionsOperation
type ISLSBridgedGetSpacePermittedResizeDirectionsOperation interface {
	ISLSSynchronousBridgedWindowManagementOperation

	// Topic: Methods

	HorizontalIndex() uint64
	MakeResultWithSpaceResizeDirections(directions uint64) objectivec.IObject
	SpaceID() uint64
	VerticalIndex() uint64
	InitWithSpaceIDVerticalIndexHorizontalIndex(id uint64, index uint64, index2 uint64) SLSBridgedGetSpacePermittedResizeDirectionsOperation
}

// Init initializes the instance.
func (s SLSBridgedGetSpacePermittedResizeDirectionsOperation) Init() SLSBridgedGetSpacePermittedResizeDirectionsOperation {
	rv := objc.Send[SLSBridgedGetSpacePermittedResizeDirectionsOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedGetSpacePermittedResizeDirectionsOperation) Autorelease() SLSBridgedGetSpacePermittedResizeDirectionsOperation {
	rv := objc.Send[SLSBridgedGetSpacePermittedResizeDirectionsOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedGetSpacePermittedResizeDirectionsOperation creates a new SLSBridgedGetSpacePermittedResizeDirectionsOperation instance.
func NewSLSBridgedGetSpacePermittedResizeDirectionsOperation() SLSBridgedGetSpacePermittedResizeDirectionsOperation {
	class := getSLSBridgedGetSpacePermittedResizeDirectionsOperationClass()
	rv := objc.Send[SLSBridgedGetSpacePermittedResizeDirectionsOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedGetSpacePermittedResizeDirectionsOperation/initWithCoder:
func NewSLSBridgedGetSpacePermittedResizeDirectionsOperationWithCoder(coder objectivec.IObject) SLSBridgedGetSpacePermittedResizeDirectionsOperation {
	instance := getSLSBridgedGetSpacePermittedResizeDirectionsOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedGetSpacePermittedResizeDirectionsOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedGetSpacePermittedResizeDirectionsOperation/initWithSpaceID:verticalIndex:horizontalIndex:
func NewSLSBridgedGetSpacePermittedResizeDirectionsOperationWithSpaceIDVerticalIndexHorizontalIndex(id uint64, index uint64, index2 uint64) SLSBridgedGetSpacePermittedResizeDirectionsOperation {
	instance := getSLSBridgedGetSpacePermittedResizeDirectionsOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpaceID:verticalIndex:horizontalIndex:"), id, index, index2)
	return SLSBridgedGetSpacePermittedResizeDirectionsOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedGetSpacePermittedResizeDirectionsOperation/makeResultWithSpaceResizeDirections:
func (s SLSBridgedGetSpacePermittedResizeDirectionsOperation) MakeResultWithSpaceResizeDirections(directions uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("makeResultWithSpaceResizeDirections:"), directions)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedGetSpacePermittedResizeDirectionsOperation/initWithSpaceID:verticalIndex:horizontalIndex:
func (s SLSBridgedGetSpacePermittedResizeDirectionsOperation) InitWithSpaceIDVerticalIndexHorizontalIndex(id uint64, index uint64, index2 uint64) SLSBridgedGetSpacePermittedResizeDirectionsOperation {
	rv := objc.Send[SLSBridgedGetSpacePermittedResizeDirectionsOperation](s.ID, objc.Sel("initWithSpaceID:verticalIndex:horizontalIndex:"), id, index, index2)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedGetSpacePermittedResizeDirectionsOperation/horizontalIndex
func (s SLSBridgedGetSpacePermittedResizeDirectionsOperation) HorizontalIndex() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("horizontalIndex"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedGetSpacePermittedResizeDirectionsOperation/spaceID
func (s SLSBridgedGetSpacePermittedResizeDirectionsOperation) SpaceID() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedGetSpacePermittedResizeDirectionsOperation/verticalIndex
func (s SLSBridgedGetSpacePermittedResizeDirectionsOperation) VerticalIndex() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("verticalIndex"))
	return rv
}
