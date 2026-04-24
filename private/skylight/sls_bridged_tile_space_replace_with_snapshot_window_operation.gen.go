// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedTileSpaceReplaceWithSnapshotWindowOperation] class.
var (
	_SLSBridgedTileSpaceReplaceWithSnapshotWindowOperationClass     SLSBridgedTileSpaceReplaceWithSnapshotWindowOperationClass
	_SLSBridgedTileSpaceReplaceWithSnapshotWindowOperationClassOnce sync.Once
)

func getSLSBridgedTileSpaceReplaceWithSnapshotWindowOperationClass() SLSBridgedTileSpaceReplaceWithSnapshotWindowOperationClass {
	_SLSBridgedTileSpaceReplaceWithSnapshotWindowOperationClassOnce.Do(func() {
		_SLSBridgedTileSpaceReplaceWithSnapshotWindowOperationClass = SLSBridgedTileSpaceReplaceWithSnapshotWindowOperationClass{class: objc.GetClass("SLSBridgedTileSpaceReplaceWithSnapshotWindowOperation")}
	})
	return _SLSBridgedTileSpaceReplaceWithSnapshotWindowOperationClass
}

// GetSLSBridgedTileSpaceReplaceWithSnapshotWindowOperationClass returns the class object for SLSBridgedTileSpaceReplaceWithSnapshotWindowOperation.
func GetSLSBridgedTileSpaceReplaceWithSnapshotWindowOperationClass() SLSBridgedTileSpaceReplaceWithSnapshotWindowOperationClass {
	return getSLSBridgedTileSpaceReplaceWithSnapshotWindowOperationClass()
}

type SLSBridgedTileSpaceReplaceWithSnapshotWindowOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedTileSpaceReplaceWithSnapshotWindowOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedTileSpaceReplaceWithSnapshotWindowOperationClass) Alloc() SLSBridgedTileSpaceReplaceWithSnapshotWindowOperation {
	rv := objc.Send[SLSBridgedTileSpaceReplaceWithSnapshotWindowOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedTileSpaceReplaceWithSnapshotWindowOperation.MakeResultWithWindowID]
//   - [SLSBridgedTileSpaceReplaceWithSnapshotWindowOperation.SpaceID]
//   - [SLSBridgedTileSpaceReplaceWithSnapshotWindowOperation.InitWithSpaceID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedTileSpaceReplaceWithSnapshotWindowOperation
type SLSBridgedTileSpaceReplaceWithSnapshotWindowOperation struct {
	SLSSynchronousBridgedWindowManagementOperation
}

// SLSBridgedTileSpaceReplaceWithSnapshotWindowOperationFromID constructs a [SLSBridgedTileSpaceReplaceWithSnapshotWindowOperation] from an objc.ID.
func SLSBridgedTileSpaceReplaceWithSnapshotWindowOperationFromID(id objc.ID) SLSBridgedTileSpaceReplaceWithSnapshotWindowOperation {
	return SLSBridgedTileSpaceReplaceWithSnapshotWindowOperation{SLSSynchronousBridgedWindowManagementOperation: SLSSynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedTileSpaceReplaceWithSnapshotWindowOperation implements ISLSBridgedTileSpaceReplaceWithSnapshotWindowOperation.
var _ ISLSBridgedTileSpaceReplaceWithSnapshotWindowOperation = SLSBridgedTileSpaceReplaceWithSnapshotWindowOperation{}

// An interface definition for the [SLSBridgedTileSpaceReplaceWithSnapshotWindowOperation] class.
//
// # Methods
//
//   - [ISLSBridgedTileSpaceReplaceWithSnapshotWindowOperation.MakeResultWithWindowID]
//   - [ISLSBridgedTileSpaceReplaceWithSnapshotWindowOperation.SpaceID]
//   - [ISLSBridgedTileSpaceReplaceWithSnapshotWindowOperation.InitWithSpaceID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedTileSpaceReplaceWithSnapshotWindowOperation
type ISLSBridgedTileSpaceReplaceWithSnapshotWindowOperation interface {
	ISLSSynchronousBridgedWindowManagementOperation

	// Topic: Methods

	MakeResultWithWindowID(id uint32) objectivec.IObject
	SpaceID() uint64
	InitWithSpaceID(id uint64) SLSBridgedTileSpaceReplaceWithSnapshotWindowOperation
}

// Init initializes the instance.
func (s SLSBridgedTileSpaceReplaceWithSnapshotWindowOperation) Init() SLSBridgedTileSpaceReplaceWithSnapshotWindowOperation {
	rv := objc.Send[SLSBridgedTileSpaceReplaceWithSnapshotWindowOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedTileSpaceReplaceWithSnapshotWindowOperation) Autorelease() SLSBridgedTileSpaceReplaceWithSnapshotWindowOperation {
	rv := objc.Send[SLSBridgedTileSpaceReplaceWithSnapshotWindowOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedTileSpaceReplaceWithSnapshotWindowOperation creates a new SLSBridgedTileSpaceReplaceWithSnapshotWindowOperation instance.
func NewSLSBridgedTileSpaceReplaceWithSnapshotWindowOperation() SLSBridgedTileSpaceReplaceWithSnapshotWindowOperation {
	class := getSLSBridgedTileSpaceReplaceWithSnapshotWindowOperationClass()
	rv := objc.Send[SLSBridgedTileSpaceReplaceWithSnapshotWindowOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedTileSpaceReplaceWithSnapshotWindowOperation/initWithCoder:
func NewSLSBridgedTileSpaceReplaceWithSnapshotWindowOperationWithCoder(coder objectivec.IObject) SLSBridgedTileSpaceReplaceWithSnapshotWindowOperation {
	instance := getSLSBridgedTileSpaceReplaceWithSnapshotWindowOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedTileSpaceReplaceWithSnapshotWindowOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedTileSpaceReplaceWithSnapshotWindowOperation/initWithSpaceID:
func NewSLSBridgedTileSpaceReplaceWithSnapshotWindowOperationWithSpaceID(id uint64) SLSBridgedTileSpaceReplaceWithSnapshotWindowOperation {
	instance := getSLSBridgedTileSpaceReplaceWithSnapshotWindowOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpaceID:"), id)
	return SLSBridgedTileSpaceReplaceWithSnapshotWindowOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedTileSpaceReplaceWithSnapshotWindowOperation/makeResultWithWindowID:
func (s SLSBridgedTileSpaceReplaceWithSnapshotWindowOperation) MakeResultWithWindowID(id uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("makeResultWithWindowID:"), id)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedTileSpaceReplaceWithSnapshotWindowOperation/initWithSpaceID:
func (s SLSBridgedTileSpaceReplaceWithSnapshotWindowOperation) InitWithSpaceID(id uint64) SLSBridgedTileSpaceReplaceWithSnapshotWindowOperation {
	rv := objc.Send[SLSBridgedTileSpaceReplaceWithSnapshotWindowOperation](s.ID, objc.Sel("initWithSpaceID:"), id)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedTileSpaceReplaceWithSnapshotWindowOperation/spaceID
func (s SLSBridgedTileSpaceReplaceWithSnapshotWindowOperation) SpaceID() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}
