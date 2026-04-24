// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedTileSpaceMoveSpacersForSizeOperation] class.
var (
	_SLSBridgedTileSpaceMoveSpacersForSizeOperationClass     SLSBridgedTileSpaceMoveSpacersForSizeOperationClass
	_SLSBridgedTileSpaceMoveSpacersForSizeOperationClassOnce sync.Once
)

func getSLSBridgedTileSpaceMoveSpacersForSizeOperationClass() SLSBridgedTileSpaceMoveSpacersForSizeOperationClass {
	_SLSBridgedTileSpaceMoveSpacersForSizeOperationClassOnce.Do(func() {
		_SLSBridgedTileSpaceMoveSpacersForSizeOperationClass = SLSBridgedTileSpaceMoveSpacersForSizeOperationClass{class: objc.GetClass("SLSBridgedTileSpaceMoveSpacersForSizeOperation")}
	})
	return _SLSBridgedTileSpaceMoveSpacersForSizeOperationClass
}

// GetSLSBridgedTileSpaceMoveSpacersForSizeOperationClass returns the class object for SLSBridgedTileSpaceMoveSpacersForSizeOperation.
func GetSLSBridgedTileSpaceMoveSpacersForSizeOperationClass() SLSBridgedTileSpaceMoveSpacersForSizeOperationClass {
	return getSLSBridgedTileSpaceMoveSpacersForSizeOperationClass()
}

type SLSBridgedTileSpaceMoveSpacersForSizeOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedTileSpaceMoveSpacersForSizeOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedTileSpaceMoveSpacersForSizeOperationClass) Alloc() SLSBridgedTileSpaceMoveSpacersForSizeOperation {
	rv := objc.Send[SLSBridgedTileSpaceMoveSpacersForSizeOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedTileSpaceMoveSpacersForSizeOperation.Size]
//   - [SLSBridgedTileSpaceMoveSpacersForSizeOperation.TileSpaceID]
//   - [SLSBridgedTileSpaceMoveSpacersForSizeOperation.InitWithTileSpaceIDSize]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedTileSpaceMoveSpacersForSizeOperation
type SLSBridgedTileSpaceMoveSpacersForSizeOperation struct {
	SLSAsynchronousBridgedWindowManagementOperation
}

// SLSBridgedTileSpaceMoveSpacersForSizeOperationFromID constructs a [SLSBridgedTileSpaceMoveSpacersForSizeOperation] from an objc.ID.
func SLSBridgedTileSpaceMoveSpacersForSizeOperationFromID(id objc.ID) SLSBridgedTileSpaceMoveSpacersForSizeOperation {
	return SLSBridgedTileSpaceMoveSpacersForSizeOperation{SLSAsynchronousBridgedWindowManagementOperation: SLSAsynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedTileSpaceMoveSpacersForSizeOperation implements ISLSBridgedTileSpaceMoveSpacersForSizeOperation.
var _ ISLSBridgedTileSpaceMoveSpacersForSizeOperation = SLSBridgedTileSpaceMoveSpacersForSizeOperation{}

// An interface definition for the [SLSBridgedTileSpaceMoveSpacersForSizeOperation] class.
//
// # Methods
//
//   - [ISLSBridgedTileSpaceMoveSpacersForSizeOperation.Size]
//   - [ISLSBridgedTileSpaceMoveSpacersForSizeOperation.TileSpaceID]
//   - [ISLSBridgedTileSpaceMoveSpacersForSizeOperation.InitWithTileSpaceIDSize]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedTileSpaceMoveSpacersForSizeOperation
type ISLSBridgedTileSpaceMoveSpacersForSizeOperation interface {
	ISLSAsynchronousBridgedWindowManagementOperation

	// Topic: Methods

	Size() corefoundation.CGSize
	TileSpaceID() uint64
	InitWithTileSpaceIDSize(id uint64, size corefoundation.CGSize) SLSBridgedTileSpaceMoveSpacersForSizeOperation
}

// Init initializes the instance.
func (s SLSBridgedTileSpaceMoveSpacersForSizeOperation) Init() SLSBridgedTileSpaceMoveSpacersForSizeOperation {
	rv := objc.Send[SLSBridgedTileSpaceMoveSpacersForSizeOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedTileSpaceMoveSpacersForSizeOperation) Autorelease() SLSBridgedTileSpaceMoveSpacersForSizeOperation {
	rv := objc.Send[SLSBridgedTileSpaceMoveSpacersForSizeOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedTileSpaceMoveSpacersForSizeOperation creates a new SLSBridgedTileSpaceMoveSpacersForSizeOperation instance.
func NewSLSBridgedTileSpaceMoveSpacersForSizeOperation() SLSBridgedTileSpaceMoveSpacersForSizeOperation {
	class := getSLSBridgedTileSpaceMoveSpacersForSizeOperationClass()
	rv := objc.Send[SLSBridgedTileSpaceMoveSpacersForSizeOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedTileSpaceMoveSpacersForSizeOperation/initWithCoder:
func NewSLSBridgedTileSpaceMoveSpacersForSizeOperationWithCoder(coder objectivec.IObject) SLSBridgedTileSpaceMoveSpacersForSizeOperation {
	instance := getSLSBridgedTileSpaceMoveSpacersForSizeOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedTileSpaceMoveSpacersForSizeOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedTileSpaceMoveSpacersForSizeOperation/initWithTileSpaceID:size:
func NewSLSBridgedTileSpaceMoveSpacersForSizeOperationWithTileSpaceIDSize(id uint64, size corefoundation.CGSize) SLSBridgedTileSpaceMoveSpacersForSizeOperation {
	instance := getSLSBridgedTileSpaceMoveSpacersForSizeOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTileSpaceID:size:"), id, size)
	return SLSBridgedTileSpaceMoveSpacersForSizeOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedTileSpaceMoveSpacersForSizeOperation/initWithTileSpaceID:size:
func (s SLSBridgedTileSpaceMoveSpacersForSizeOperation) InitWithTileSpaceIDSize(id uint64, size corefoundation.CGSize) SLSBridgedTileSpaceMoveSpacersForSizeOperation {
	rv := objc.Send[SLSBridgedTileSpaceMoveSpacersForSizeOperation](s.ID, objc.Sel("initWithTileSpaceID:size:"), id, size)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedTileSpaceMoveSpacersForSizeOperation/size
func (s SLSBridgedTileSpaceMoveSpacersForSizeOperation) Size() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](s.ID, objc.Sel("size"))
	return corefoundation.CGSize(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedTileSpaceMoveSpacersForSizeOperation/tileSpaceID
func (s SLSBridgedTileSpaceMoveSpacersForSizeOperation) TileSpaceID() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("tileSpaceID"))
	return rv
}
