// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedSpaceCanCreateTileOperation] class.
var (
	_SLSBridgedSpaceCanCreateTileOperationClass     SLSBridgedSpaceCanCreateTileOperationClass
	_SLSBridgedSpaceCanCreateTileOperationClassOnce sync.Once
)

func getSLSBridgedSpaceCanCreateTileOperationClass() SLSBridgedSpaceCanCreateTileOperationClass {
	_SLSBridgedSpaceCanCreateTileOperationClassOnce.Do(func() {
		_SLSBridgedSpaceCanCreateTileOperationClass = SLSBridgedSpaceCanCreateTileOperationClass{class: objc.GetClass("SLSBridgedSpaceCanCreateTileOperation")}
	})
	return _SLSBridgedSpaceCanCreateTileOperationClass
}

// GetSLSBridgedSpaceCanCreateTileOperationClass returns the class object for SLSBridgedSpaceCanCreateTileOperation.
func GetSLSBridgedSpaceCanCreateTileOperationClass() SLSBridgedSpaceCanCreateTileOperationClass {
	return getSLSBridgedSpaceCanCreateTileOperationClass()
}

type SLSBridgedSpaceCanCreateTileOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedSpaceCanCreateTileOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedSpaceCanCreateTileOperationClass) Alloc() SLSBridgedSpaceCanCreateTileOperation {
	rv := objc.Send[SLSBridgedSpaceCanCreateTileOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedSpaceCanCreateTileOperation.MakeResultWithBoolValue]
//   - [SLSBridgedSpaceCanCreateTileOperation.MinSize]
//   - [SLSBridgedSpaceCanCreateTileOperation.SpaceID]
//   - [SLSBridgedSpaceCanCreateTileOperation.InitWithSpaceIDMinSize]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCanCreateTileOperation
type SLSBridgedSpaceCanCreateTileOperation struct {
	SLSSynchronousBridgedWindowManagementOperation
}

// SLSBridgedSpaceCanCreateTileOperationFromID constructs a [SLSBridgedSpaceCanCreateTileOperation] from an objc.ID.
func SLSBridgedSpaceCanCreateTileOperationFromID(id objc.ID) SLSBridgedSpaceCanCreateTileOperation {
	return SLSBridgedSpaceCanCreateTileOperation{SLSSynchronousBridgedWindowManagementOperation: SLSSynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedSpaceCanCreateTileOperation implements ISLSBridgedSpaceCanCreateTileOperation.
var _ ISLSBridgedSpaceCanCreateTileOperation = SLSBridgedSpaceCanCreateTileOperation{}

// An interface definition for the [SLSBridgedSpaceCanCreateTileOperation] class.
//
// # Methods
//
//   - [ISLSBridgedSpaceCanCreateTileOperation.MakeResultWithBoolValue]
//   - [ISLSBridgedSpaceCanCreateTileOperation.MinSize]
//   - [ISLSBridgedSpaceCanCreateTileOperation.SpaceID]
//   - [ISLSBridgedSpaceCanCreateTileOperation.InitWithSpaceIDMinSize]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCanCreateTileOperation
type ISLSBridgedSpaceCanCreateTileOperation interface {
	ISLSSynchronousBridgedWindowManagementOperation

	// Topic: Methods

	MakeResultWithBoolValue(value bool) objectivec.IObject
	MinSize() corefoundation.CGSize
	SpaceID() uint64
	InitWithSpaceIDMinSize(id uint64, size corefoundation.CGSize) SLSBridgedSpaceCanCreateTileOperation
}

// Init initializes the instance.
func (s SLSBridgedSpaceCanCreateTileOperation) Init() SLSBridgedSpaceCanCreateTileOperation {
	rv := objc.Send[SLSBridgedSpaceCanCreateTileOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedSpaceCanCreateTileOperation) Autorelease() SLSBridgedSpaceCanCreateTileOperation {
	rv := objc.Send[SLSBridgedSpaceCanCreateTileOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedSpaceCanCreateTileOperation creates a new SLSBridgedSpaceCanCreateTileOperation instance.
func NewSLSBridgedSpaceCanCreateTileOperation() SLSBridgedSpaceCanCreateTileOperation {
	class := getSLSBridgedSpaceCanCreateTileOperationClass()
	rv := objc.Send[SLSBridgedSpaceCanCreateTileOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCanCreateTileOperation/initWithCoder:
func NewSLSBridgedSpaceCanCreateTileOperationWithCoder(coder objectivec.IObject) SLSBridgedSpaceCanCreateTileOperation {
	instance := getSLSBridgedSpaceCanCreateTileOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedSpaceCanCreateTileOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCanCreateTileOperation/initWithSpaceID:minSize:
func NewSLSBridgedSpaceCanCreateTileOperationWithSpaceIDMinSize(id uint64, size corefoundation.CGSize) SLSBridgedSpaceCanCreateTileOperation {
	instance := getSLSBridgedSpaceCanCreateTileOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpaceID:minSize:"), id, size)
	return SLSBridgedSpaceCanCreateTileOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCanCreateTileOperation/makeResultWithBoolValue:
func (s SLSBridgedSpaceCanCreateTileOperation) MakeResultWithBoolValue(value bool) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("makeResultWithBoolValue:"), value)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCanCreateTileOperation/initWithSpaceID:minSize:
func (s SLSBridgedSpaceCanCreateTileOperation) InitWithSpaceIDMinSize(id uint64, size corefoundation.CGSize) SLSBridgedSpaceCanCreateTileOperation {
	rv := objc.Send[SLSBridgedSpaceCanCreateTileOperation](s.ID, objc.Sel("initWithSpaceID:minSize:"), id, size)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCanCreateTileOperation/minSize
func (s SLSBridgedSpaceCanCreateTileOperation) MinSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](s.ID, objc.Sel("minSize"))
	return corefoundation.CGSize(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCanCreateTileOperation/spaceID
func (s SLSBridgedSpaceCanCreateTileOperation) SpaceID() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}
