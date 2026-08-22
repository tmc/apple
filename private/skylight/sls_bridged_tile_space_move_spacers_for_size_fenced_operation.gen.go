// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedTileSpaceMoveSpacersForSizeFencedOperation] class.
var (
	_SLSBridgedTileSpaceMoveSpacersForSizeFencedOperationClass     SLSBridgedTileSpaceMoveSpacersForSizeFencedOperationClass
	_SLSBridgedTileSpaceMoveSpacersForSizeFencedOperationClassOnce sync.Once
)

func getSLSBridgedTileSpaceMoveSpacersForSizeFencedOperationClass() SLSBridgedTileSpaceMoveSpacersForSizeFencedOperationClass {
	_SLSBridgedTileSpaceMoveSpacersForSizeFencedOperationClassOnce.Do(func() {
		_SLSBridgedTileSpaceMoveSpacersForSizeFencedOperationClass = SLSBridgedTileSpaceMoveSpacersForSizeFencedOperationClass{class: objc.GetClass("SLSBridgedTileSpaceMoveSpacersForSizeFencedOperation")}
	})
	return _SLSBridgedTileSpaceMoveSpacersForSizeFencedOperationClass
}

// GetSLSBridgedTileSpaceMoveSpacersForSizeFencedOperationClass returns the class object for SLSBridgedTileSpaceMoveSpacersForSizeFencedOperation.
func GetSLSBridgedTileSpaceMoveSpacersForSizeFencedOperationClass() SLSBridgedTileSpaceMoveSpacersForSizeFencedOperationClass {
	return getSLSBridgedTileSpaceMoveSpacersForSizeFencedOperationClass()
}

type SLSBridgedTileSpaceMoveSpacersForSizeFencedOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedTileSpaceMoveSpacersForSizeFencedOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedTileSpaceMoveSpacersForSizeFencedOperationClass) Alloc() SLSBridgedTileSpaceMoveSpacersForSizeFencedOperation {
	rv := objc.SendIfResponds[SLSBridgedTileSpaceMoveSpacersForSizeFencedOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedTileSpaceMoveSpacersForSizeFencedOperation.FencePort]
//   - [SLSBridgedTileSpaceMoveSpacersForSizeFencedOperation.Size]
//   - [SLSBridgedTileSpaceMoveSpacersForSizeFencedOperation.TileSpaceID]
//   - [SLSBridgedTileSpaceMoveSpacersForSizeFencedOperation.InitWithTileSpaceIDSizeFencePort]
type SLSBridgedTileSpaceMoveSpacersForSizeFencedOperation struct {
	SLSAsynchronousBridgedWindowManagementOperation
}

// SLSBridgedTileSpaceMoveSpacersForSizeFencedOperationFromID constructs a [SLSBridgedTileSpaceMoveSpacersForSizeFencedOperation] from an objc.ID.
func SLSBridgedTileSpaceMoveSpacersForSizeFencedOperationFromID(id objc.ID) SLSBridgedTileSpaceMoveSpacersForSizeFencedOperation {
	return SLSBridgedTileSpaceMoveSpacersForSizeFencedOperation{SLSAsynchronousBridgedWindowManagementOperation: SLSAsynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedTileSpaceMoveSpacersForSizeFencedOperation implements ISLSBridgedTileSpaceMoveSpacersForSizeFencedOperation.
var _ ISLSBridgedTileSpaceMoveSpacersForSizeFencedOperation = SLSBridgedTileSpaceMoveSpacersForSizeFencedOperation{}

// An interface definition for the [SLSBridgedTileSpaceMoveSpacersForSizeFencedOperation] class.
//
// # Methods
//
//   - [ISLSBridgedTileSpaceMoveSpacersForSizeFencedOperation.FencePort]
//   - [ISLSBridgedTileSpaceMoveSpacersForSizeFencedOperation.Size]
//   - [ISLSBridgedTileSpaceMoveSpacersForSizeFencedOperation.TileSpaceID]
//   - [ISLSBridgedTileSpaceMoveSpacersForSizeFencedOperation.InitWithTileSpaceIDSizeFencePort]
type ISLSBridgedTileSpaceMoveSpacersForSizeFencedOperation interface {
	ISLSAsynchronousBridgedWindowManagementOperation

	// Topic: Methods

	FencePort() uint32
	Size() corefoundation.CGSize
	TileSpaceID() uint64
	InitWithTileSpaceIDSizeFencePort(id uint64, size corefoundation.CGSize, port uint32) SLSBridgedTileSpaceMoveSpacersForSizeFencedOperation
}

// Init initializes the instance.
func (s SLSBridgedTileSpaceMoveSpacersForSizeFencedOperation) Init() SLSBridgedTileSpaceMoveSpacersForSizeFencedOperation {
	rv := objc.SendIfResponds[SLSBridgedTileSpaceMoveSpacersForSizeFencedOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedTileSpaceMoveSpacersForSizeFencedOperation) Autorelease() SLSBridgedTileSpaceMoveSpacersForSizeFencedOperation {
	rv := objc.SendIfResponds[SLSBridgedTileSpaceMoveSpacersForSizeFencedOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedTileSpaceMoveSpacersForSizeFencedOperation creates a new SLSBridgedTileSpaceMoveSpacersForSizeFencedOperation instance.
func NewSLSBridgedTileSpaceMoveSpacersForSizeFencedOperation() SLSBridgedTileSpaceMoveSpacersForSizeFencedOperation {
	class := getSLSBridgedTileSpaceMoveSpacersForSizeFencedOperationClass()
	rv := objc.SendIfResponds[SLSBridgedTileSpaceMoveSpacersForSizeFencedOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSLSBridgedTileSpaceMoveSpacersForSizeFencedOperationWithCoder(coder objectivec.IObject) SLSBridgedTileSpaceMoveSpacersForSizeFencedOperation {
	instance := getSLSBridgedTileSpaceMoveSpacersForSizeFencedOperationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedTileSpaceMoveSpacersForSizeFencedOperationFromID(rv)
}

func NewSLSBridgedTileSpaceMoveSpacersForSizeFencedOperationWithTileSpaceIDSizeFencePort(id uint64, size corefoundation.CGSize, port uint32) SLSBridgedTileSpaceMoveSpacersForSizeFencedOperation {
	instance := getSLSBridgedTileSpaceMoveSpacersForSizeFencedOperationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithTileSpaceID:size:fencePort:"), id, size, port)
	return SLSBridgedTileSpaceMoveSpacersForSizeFencedOperationFromID(rv)
}

func (s SLSBridgedTileSpaceMoveSpacersForSizeFencedOperation) InitWithTileSpaceIDSizeFencePort(id uint64, size corefoundation.CGSize, port uint32) SLSBridgedTileSpaceMoveSpacersForSizeFencedOperation {
	rv := objc.SendIfResponds[SLSBridgedTileSpaceMoveSpacersForSizeFencedOperation](s.ID, objc.Sel("initWithTileSpaceID:size:fencePort:"), id, size, port)
	return rv
}

func (s SLSBridgedTileSpaceMoveSpacersForSizeFencedOperation) FencePort() uint32 {
	rv := objc.SendIfResponds[uint32](s.ID, objc.Sel("fencePort"))
	return rv
}
func (s SLSBridgedTileSpaceMoveSpacersForSizeFencedOperation) Size() corefoundation.CGSize {
	rv := objc.SendIfResponds[corefoundation.CGSize](s.ID, objc.Sel("size"))
	return corefoundation.CGSize(rv)
}
func (s SLSBridgedTileSpaceMoveSpacersForSizeFencedOperation) TileSpaceID() uint64 {
	rv := objc.SendIfResponds[uint64](s.ID, objc.Sel("tileSpaceID"))
	return rv
}
