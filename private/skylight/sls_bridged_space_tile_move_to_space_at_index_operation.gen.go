// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedSpaceTileMoveToSpaceAtIndexOperation] class.
var (
	_SLSBridgedSpaceTileMoveToSpaceAtIndexOperationClass     SLSBridgedSpaceTileMoveToSpaceAtIndexOperationClass
	_SLSBridgedSpaceTileMoveToSpaceAtIndexOperationClassOnce sync.Once
)

func getSLSBridgedSpaceTileMoveToSpaceAtIndexOperationClass() SLSBridgedSpaceTileMoveToSpaceAtIndexOperationClass {
	_SLSBridgedSpaceTileMoveToSpaceAtIndexOperationClassOnce.Do(func() {
		_SLSBridgedSpaceTileMoveToSpaceAtIndexOperationClass = SLSBridgedSpaceTileMoveToSpaceAtIndexOperationClass{class: objc.GetClass("SLSBridgedSpaceTileMoveToSpaceAtIndexOperation")}
	})
	return _SLSBridgedSpaceTileMoveToSpaceAtIndexOperationClass
}

// GetSLSBridgedSpaceTileMoveToSpaceAtIndexOperationClass returns the class object for SLSBridgedSpaceTileMoveToSpaceAtIndexOperation.
func GetSLSBridgedSpaceTileMoveToSpaceAtIndexOperationClass() SLSBridgedSpaceTileMoveToSpaceAtIndexOperationClass {
	return getSLSBridgedSpaceTileMoveToSpaceAtIndexOperationClass()
}

type SLSBridgedSpaceTileMoveToSpaceAtIndexOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedSpaceTileMoveToSpaceAtIndexOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedSpaceTileMoveToSpaceAtIndexOperationClass) Alloc() SLSBridgedSpaceTileMoveToSpaceAtIndexOperation {
	rv := objc.Send[SLSBridgedSpaceTileMoveToSpaceAtIndexOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedSpaceTileMoveToSpaceAtIndexOperation.Index]
//   - [SLSBridgedSpaceTileMoveToSpaceAtIndexOperation.ParentID]
//   - [SLSBridgedSpaceTileMoveToSpaceAtIndexOperation.TileID]
//   - [SLSBridgedSpaceTileMoveToSpaceAtIndexOperation.InitWithTileIDParentIDIndex]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceTileMoveToSpaceAtIndexOperation
type SLSBridgedSpaceTileMoveToSpaceAtIndexOperation struct {
	SLSAsynchronousBridgedWindowManagementOperation
}

// SLSBridgedSpaceTileMoveToSpaceAtIndexOperationFromID constructs a [SLSBridgedSpaceTileMoveToSpaceAtIndexOperation] from an objc.ID.
func SLSBridgedSpaceTileMoveToSpaceAtIndexOperationFromID(id objc.ID) SLSBridgedSpaceTileMoveToSpaceAtIndexOperation {
	return SLSBridgedSpaceTileMoveToSpaceAtIndexOperation{SLSAsynchronousBridgedWindowManagementOperation: SLSAsynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedSpaceTileMoveToSpaceAtIndexOperation implements ISLSBridgedSpaceTileMoveToSpaceAtIndexOperation.
var _ ISLSBridgedSpaceTileMoveToSpaceAtIndexOperation = SLSBridgedSpaceTileMoveToSpaceAtIndexOperation{}

// An interface definition for the [SLSBridgedSpaceTileMoveToSpaceAtIndexOperation] class.
//
// # Methods
//
//   - [ISLSBridgedSpaceTileMoveToSpaceAtIndexOperation.Index]
//   - [ISLSBridgedSpaceTileMoveToSpaceAtIndexOperation.ParentID]
//   - [ISLSBridgedSpaceTileMoveToSpaceAtIndexOperation.TileID]
//   - [ISLSBridgedSpaceTileMoveToSpaceAtIndexOperation.InitWithTileIDParentIDIndex]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceTileMoveToSpaceAtIndexOperation
type ISLSBridgedSpaceTileMoveToSpaceAtIndexOperation interface {
	ISLSAsynchronousBridgedWindowManagementOperation

	// Topic: Methods

	Index() uint64
	ParentID() uint64
	TileID() uint64
	InitWithTileIDParentIDIndex(id uint64, id2 uint64, index uint64) SLSBridgedSpaceTileMoveToSpaceAtIndexOperation
}

// Init initializes the instance.
func (s SLSBridgedSpaceTileMoveToSpaceAtIndexOperation) Init() SLSBridgedSpaceTileMoveToSpaceAtIndexOperation {
	rv := objc.Send[SLSBridgedSpaceTileMoveToSpaceAtIndexOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedSpaceTileMoveToSpaceAtIndexOperation) Autorelease() SLSBridgedSpaceTileMoveToSpaceAtIndexOperation {
	rv := objc.Send[SLSBridgedSpaceTileMoveToSpaceAtIndexOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedSpaceTileMoveToSpaceAtIndexOperation creates a new SLSBridgedSpaceTileMoveToSpaceAtIndexOperation instance.
func NewSLSBridgedSpaceTileMoveToSpaceAtIndexOperation() SLSBridgedSpaceTileMoveToSpaceAtIndexOperation {
	class := getSLSBridgedSpaceTileMoveToSpaceAtIndexOperationClass()
	rv := objc.Send[SLSBridgedSpaceTileMoveToSpaceAtIndexOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceTileMoveToSpaceAtIndexOperation/initWithCoder:
func NewSLSBridgedSpaceTileMoveToSpaceAtIndexOperationWithCoder(coder objectivec.IObject) SLSBridgedSpaceTileMoveToSpaceAtIndexOperation {
	instance := getSLSBridgedSpaceTileMoveToSpaceAtIndexOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedSpaceTileMoveToSpaceAtIndexOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceTileMoveToSpaceAtIndexOperation/initWithTileID:parentID:index:
func NewSLSBridgedSpaceTileMoveToSpaceAtIndexOperationWithTileIDParentIDIndex(id uint64, id2 uint64, index uint64) SLSBridgedSpaceTileMoveToSpaceAtIndexOperation {
	instance := getSLSBridgedSpaceTileMoveToSpaceAtIndexOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTileID:parentID:index:"), id, id2, index)
	return SLSBridgedSpaceTileMoveToSpaceAtIndexOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceTileMoveToSpaceAtIndexOperation/initWithTileID:parentID:index:
func (s SLSBridgedSpaceTileMoveToSpaceAtIndexOperation) InitWithTileIDParentIDIndex(id uint64, id2 uint64, index uint64) SLSBridgedSpaceTileMoveToSpaceAtIndexOperation {
	rv := objc.Send[SLSBridgedSpaceTileMoveToSpaceAtIndexOperation](s.ID, objc.Sel("initWithTileID:parentID:index:"), id, id2, index)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceTileMoveToSpaceAtIndexOperation/index
func (s SLSBridgedSpaceTileMoveToSpaceAtIndexOperation) Index() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("index"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceTileMoveToSpaceAtIndexOperation/parentID
func (s SLSBridgedSpaceTileMoveToSpaceAtIndexOperation) ParentID() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("parentID"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceTileMoveToSpaceAtIndexOperation/tileID
func (s SLSBridgedSpaceTileMoveToSpaceAtIndexOperation) TileID() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("tileID"))
	return rv
}
