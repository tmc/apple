// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedSpaceCopyTileSpacesOperation] class.
var (
	_SLSBridgedSpaceCopyTileSpacesOperationClass     SLSBridgedSpaceCopyTileSpacesOperationClass
	_SLSBridgedSpaceCopyTileSpacesOperationClassOnce sync.Once
)

func getSLSBridgedSpaceCopyTileSpacesOperationClass() SLSBridgedSpaceCopyTileSpacesOperationClass {
	_SLSBridgedSpaceCopyTileSpacesOperationClassOnce.Do(func() {
		_SLSBridgedSpaceCopyTileSpacesOperationClass = SLSBridgedSpaceCopyTileSpacesOperationClass{class: objc.GetClass("SLSBridgedSpaceCopyTileSpacesOperation")}
	})
	return _SLSBridgedSpaceCopyTileSpacesOperationClass
}

// GetSLSBridgedSpaceCopyTileSpacesOperationClass returns the class object for SLSBridgedSpaceCopyTileSpacesOperation.
func GetSLSBridgedSpaceCopyTileSpacesOperationClass() SLSBridgedSpaceCopyTileSpacesOperationClass {
	return getSLSBridgedSpaceCopyTileSpacesOperationClass()
}

type SLSBridgedSpaceCopyTileSpacesOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedSpaceCopyTileSpacesOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedSpaceCopyTileSpacesOperationClass) Alloc() SLSBridgedSpaceCopyTileSpacesOperation {
	rv := objc.Send[SLSBridgedSpaceCopyTileSpacesOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedSpaceCopyTileSpacesOperation.MakeResultWithNumbers]
//   - [SLSBridgedSpaceCopyTileSpacesOperation.SpaceID]
//   - [SLSBridgedSpaceCopyTileSpacesOperation.InitWithSpaceID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCopyTileSpacesOperation
type SLSBridgedSpaceCopyTileSpacesOperation struct {
	SLSSynchronousBridgedWindowManagementOperation
}

// SLSBridgedSpaceCopyTileSpacesOperationFromID constructs a [SLSBridgedSpaceCopyTileSpacesOperation] from an objc.ID.
func SLSBridgedSpaceCopyTileSpacesOperationFromID(id objc.ID) SLSBridgedSpaceCopyTileSpacesOperation {
	return SLSBridgedSpaceCopyTileSpacesOperation{SLSSynchronousBridgedWindowManagementOperation: SLSSynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedSpaceCopyTileSpacesOperation implements ISLSBridgedSpaceCopyTileSpacesOperation.
var _ ISLSBridgedSpaceCopyTileSpacesOperation = SLSBridgedSpaceCopyTileSpacesOperation{}

// An interface definition for the [SLSBridgedSpaceCopyTileSpacesOperation] class.
//
// # Methods
//
//   - [ISLSBridgedSpaceCopyTileSpacesOperation.MakeResultWithNumbers]
//   - [ISLSBridgedSpaceCopyTileSpacesOperation.SpaceID]
//   - [ISLSBridgedSpaceCopyTileSpacesOperation.InitWithSpaceID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCopyTileSpacesOperation
type ISLSBridgedSpaceCopyTileSpacesOperation interface {
	ISLSSynchronousBridgedWindowManagementOperation

	// Topic: Methods

	MakeResultWithNumbers(numbers objectivec.IObject) objectivec.IObject
	SpaceID() uint64
	InitWithSpaceID(id uint64) SLSBridgedSpaceCopyTileSpacesOperation
}

// Init initializes the instance.
func (s SLSBridgedSpaceCopyTileSpacesOperation) Init() SLSBridgedSpaceCopyTileSpacesOperation {
	rv := objc.Send[SLSBridgedSpaceCopyTileSpacesOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedSpaceCopyTileSpacesOperation) Autorelease() SLSBridgedSpaceCopyTileSpacesOperation {
	rv := objc.Send[SLSBridgedSpaceCopyTileSpacesOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedSpaceCopyTileSpacesOperation creates a new SLSBridgedSpaceCopyTileSpacesOperation instance.
func NewSLSBridgedSpaceCopyTileSpacesOperation() SLSBridgedSpaceCopyTileSpacesOperation {
	class := getSLSBridgedSpaceCopyTileSpacesOperationClass()
	rv := objc.Send[SLSBridgedSpaceCopyTileSpacesOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCopyTileSpacesOperation/initWithCoder:
func NewSLSBridgedSpaceCopyTileSpacesOperationWithCoder(coder objectivec.IObject) SLSBridgedSpaceCopyTileSpacesOperation {
	instance := getSLSBridgedSpaceCopyTileSpacesOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedSpaceCopyTileSpacesOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCopyTileSpacesOperation/initWithSpaceID:
func NewSLSBridgedSpaceCopyTileSpacesOperationWithSpaceID(id uint64) SLSBridgedSpaceCopyTileSpacesOperation {
	instance := getSLSBridgedSpaceCopyTileSpacesOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpaceID:"), id)
	return SLSBridgedSpaceCopyTileSpacesOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCopyTileSpacesOperation/makeResultWithNumbers:
func (s SLSBridgedSpaceCopyTileSpacesOperation) MakeResultWithNumbers(numbers objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("makeResultWithNumbers:"), numbers)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCopyTileSpacesOperation/initWithSpaceID:
func (s SLSBridgedSpaceCopyTileSpacesOperation) InitWithSpaceID(id uint64) SLSBridgedSpaceCopyTileSpacesOperation {
	rv := objc.Send[SLSBridgedSpaceCopyTileSpacesOperation](s.ID, objc.Sel("initWithSpaceID:"), id)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCopyTileSpacesOperation/spaceID
func (s SLSBridgedSpaceCopyTileSpacesOperation) SpaceID() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}
