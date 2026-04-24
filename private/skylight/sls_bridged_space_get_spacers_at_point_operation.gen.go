// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedSpaceGetSpacersAtPointOperation] class.
var (
	_SLSBridgedSpaceGetSpacersAtPointOperationClass     SLSBridgedSpaceGetSpacersAtPointOperationClass
	_SLSBridgedSpaceGetSpacersAtPointOperationClassOnce sync.Once
)

func getSLSBridgedSpaceGetSpacersAtPointOperationClass() SLSBridgedSpaceGetSpacersAtPointOperationClass {
	_SLSBridgedSpaceGetSpacersAtPointOperationClassOnce.Do(func() {
		_SLSBridgedSpaceGetSpacersAtPointOperationClass = SLSBridgedSpaceGetSpacersAtPointOperationClass{class: objc.GetClass("SLSBridgedSpaceGetSpacersAtPointOperation")}
	})
	return _SLSBridgedSpaceGetSpacersAtPointOperationClass
}

// GetSLSBridgedSpaceGetSpacersAtPointOperationClass returns the class object for SLSBridgedSpaceGetSpacersAtPointOperation.
func GetSLSBridgedSpaceGetSpacersAtPointOperationClass() SLSBridgedSpaceGetSpacersAtPointOperationClass {
	return getSLSBridgedSpaceGetSpacersAtPointOperationClass()
}

type SLSBridgedSpaceGetSpacersAtPointOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedSpaceGetSpacersAtPointOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedSpaceGetSpacersAtPointOperationClass) Alloc() SLSBridgedSpaceGetSpacersAtPointOperation {
	rv := objc.Send[SLSBridgedSpaceGetSpacersAtPointOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedSpaceGetSpacersAtPointOperation.MakeResultWithVerticalIndexHorizontalIndexRect]
//   - [SLSBridgedSpaceGetSpacersAtPointOperation.Point]
//   - [SLSBridgedSpaceGetSpacersAtPointOperation.SpaceID]
//   - [SLSBridgedSpaceGetSpacersAtPointOperation.InitWithSpaceIDPoint]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceGetSpacersAtPointOperation
type SLSBridgedSpaceGetSpacersAtPointOperation struct {
	SLSSynchronousBridgedWindowManagementOperation
}

// SLSBridgedSpaceGetSpacersAtPointOperationFromID constructs a [SLSBridgedSpaceGetSpacersAtPointOperation] from an objc.ID.
func SLSBridgedSpaceGetSpacersAtPointOperationFromID(id objc.ID) SLSBridgedSpaceGetSpacersAtPointOperation {
	return SLSBridgedSpaceGetSpacersAtPointOperation{SLSSynchronousBridgedWindowManagementOperation: SLSSynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedSpaceGetSpacersAtPointOperation implements ISLSBridgedSpaceGetSpacersAtPointOperation.
var _ ISLSBridgedSpaceGetSpacersAtPointOperation = SLSBridgedSpaceGetSpacersAtPointOperation{}

// An interface definition for the [SLSBridgedSpaceGetSpacersAtPointOperation] class.
//
// # Methods
//
//   - [ISLSBridgedSpaceGetSpacersAtPointOperation.MakeResultWithVerticalIndexHorizontalIndexRect]
//   - [ISLSBridgedSpaceGetSpacersAtPointOperation.Point]
//   - [ISLSBridgedSpaceGetSpacersAtPointOperation.SpaceID]
//   - [ISLSBridgedSpaceGetSpacersAtPointOperation.InitWithSpaceIDPoint]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceGetSpacersAtPointOperation
type ISLSBridgedSpaceGetSpacersAtPointOperation interface {
	ISLSSynchronousBridgedWindowManagementOperation

	// Topic: Methods

	MakeResultWithVerticalIndexHorizontalIndexRect(index uint64, index2 uint64, rect corefoundation.CGRect) objectivec.IObject
	Point() corefoundation.CGPoint
	SpaceID() uint64
	InitWithSpaceIDPoint(id uint64, point corefoundation.CGPoint) SLSBridgedSpaceGetSpacersAtPointOperation
}

// Init initializes the instance.
func (s SLSBridgedSpaceGetSpacersAtPointOperation) Init() SLSBridgedSpaceGetSpacersAtPointOperation {
	rv := objc.Send[SLSBridgedSpaceGetSpacersAtPointOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedSpaceGetSpacersAtPointOperation) Autorelease() SLSBridgedSpaceGetSpacersAtPointOperation {
	rv := objc.Send[SLSBridgedSpaceGetSpacersAtPointOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedSpaceGetSpacersAtPointOperation creates a new SLSBridgedSpaceGetSpacersAtPointOperation instance.
func NewSLSBridgedSpaceGetSpacersAtPointOperation() SLSBridgedSpaceGetSpacersAtPointOperation {
	class := getSLSBridgedSpaceGetSpacersAtPointOperationClass()
	rv := objc.Send[SLSBridgedSpaceGetSpacersAtPointOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceGetSpacersAtPointOperation/initWithCoder:
func NewSLSBridgedSpaceGetSpacersAtPointOperationWithCoder(coder objectivec.IObject) SLSBridgedSpaceGetSpacersAtPointOperation {
	instance := getSLSBridgedSpaceGetSpacersAtPointOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedSpaceGetSpacersAtPointOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceGetSpacersAtPointOperation/initWithSpaceID:point:
func NewSLSBridgedSpaceGetSpacersAtPointOperationWithSpaceIDPoint(id uint64, point corefoundation.CGPoint) SLSBridgedSpaceGetSpacersAtPointOperation {
	instance := getSLSBridgedSpaceGetSpacersAtPointOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpaceID:point:"), id, point)
	return SLSBridgedSpaceGetSpacersAtPointOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceGetSpacersAtPointOperation/makeResultWithVerticalIndex:horizontalIndex:rect:
func (s SLSBridgedSpaceGetSpacersAtPointOperation) MakeResultWithVerticalIndexHorizontalIndexRect(index uint64, index2 uint64, rect corefoundation.CGRect) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("makeResultWithVerticalIndex:horizontalIndex:rect:"), index, index2, rect)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceGetSpacersAtPointOperation/initWithSpaceID:point:
func (s SLSBridgedSpaceGetSpacersAtPointOperation) InitWithSpaceIDPoint(id uint64, point corefoundation.CGPoint) SLSBridgedSpaceGetSpacersAtPointOperation {
	rv := objc.Send[SLSBridgedSpaceGetSpacersAtPointOperation](s.ID, objc.Sel("initWithSpaceID:point:"), id, point)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceGetSpacersAtPointOperation/point
func (s SLSBridgedSpaceGetSpacersAtPointOperation) Point() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](s.ID, objc.Sel("point"))
	return corefoundation.CGPoint(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceGetSpacersAtPointOperation/spaceID
func (s SLSBridgedSpaceGetSpacersAtPointOperation) SpaceID() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}
