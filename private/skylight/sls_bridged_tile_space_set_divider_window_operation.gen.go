// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedTileSpaceSetDividerWindowOperation] class.
var (
	_SLSBridgedTileSpaceSetDividerWindowOperationClass     SLSBridgedTileSpaceSetDividerWindowOperationClass
	_SLSBridgedTileSpaceSetDividerWindowOperationClassOnce sync.Once
)

func getSLSBridgedTileSpaceSetDividerWindowOperationClass() SLSBridgedTileSpaceSetDividerWindowOperationClass {
	_SLSBridgedTileSpaceSetDividerWindowOperationClassOnce.Do(func() {
		_SLSBridgedTileSpaceSetDividerWindowOperationClass = SLSBridgedTileSpaceSetDividerWindowOperationClass{class: objc.GetClass("SLSBridgedTileSpaceSetDividerWindowOperation")}
	})
	return _SLSBridgedTileSpaceSetDividerWindowOperationClass
}

// GetSLSBridgedTileSpaceSetDividerWindowOperationClass returns the class object for SLSBridgedTileSpaceSetDividerWindowOperation.
func GetSLSBridgedTileSpaceSetDividerWindowOperationClass() SLSBridgedTileSpaceSetDividerWindowOperationClass {
	return getSLSBridgedTileSpaceSetDividerWindowOperationClass()
}

type SLSBridgedTileSpaceSetDividerWindowOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedTileSpaceSetDividerWindowOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedTileSpaceSetDividerWindowOperationClass) Alloc() SLSBridgedTileSpaceSetDividerWindowOperation {
	rv := objc.SendIfResponds[SLSBridgedTileSpaceSetDividerWindowOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedTileSpaceSetDividerWindowOperation.Direction]
//   - [SLSBridgedTileSpaceSetDividerWindowOperation.MakeResultWithVerticalIndexHorizontalIndex]
//   - [SLSBridgedTileSpaceSetDividerWindowOperation.SpaceID]
//   - [SLSBridgedTileSpaceSetDividerWindowOperation.WindowID]
//   - [SLSBridgedTileSpaceSetDividerWindowOperation.InitWithSpaceIDDirectionWindowID]
type SLSBridgedTileSpaceSetDividerWindowOperation struct {
	SLSSynchronousBridgedWindowManagementOperation
}

// SLSBridgedTileSpaceSetDividerWindowOperationFromID constructs a [SLSBridgedTileSpaceSetDividerWindowOperation] from an objc.ID.
func SLSBridgedTileSpaceSetDividerWindowOperationFromID(id objc.ID) SLSBridgedTileSpaceSetDividerWindowOperation {
	return SLSBridgedTileSpaceSetDividerWindowOperation{SLSSynchronousBridgedWindowManagementOperation: SLSSynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedTileSpaceSetDividerWindowOperation implements ISLSBridgedTileSpaceSetDividerWindowOperation.
var _ ISLSBridgedTileSpaceSetDividerWindowOperation = SLSBridgedTileSpaceSetDividerWindowOperation{}

// An interface definition for the [SLSBridgedTileSpaceSetDividerWindowOperation] class.
//
// # Methods
//
//   - [ISLSBridgedTileSpaceSetDividerWindowOperation.Direction]
//   - [ISLSBridgedTileSpaceSetDividerWindowOperation.MakeResultWithVerticalIndexHorizontalIndex]
//   - [ISLSBridgedTileSpaceSetDividerWindowOperation.SpaceID]
//   - [ISLSBridgedTileSpaceSetDividerWindowOperation.WindowID]
//   - [ISLSBridgedTileSpaceSetDividerWindowOperation.InitWithSpaceIDDirectionWindowID]
type ISLSBridgedTileSpaceSetDividerWindowOperation interface {
	ISLSSynchronousBridgedWindowManagementOperation

	// Topic: Methods

	Direction() uint64
	MakeResultWithVerticalIndexHorizontalIndex(index uint64, index2 uint64) objectivec.IObject
	SpaceID() uint64
	WindowID() uint32
	InitWithSpaceIDDirectionWindowID(id uint64, direction uint64, id2 uint32) SLSBridgedTileSpaceSetDividerWindowOperation
}

// Init initializes the instance.
func (s SLSBridgedTileSpaceSetDividerWindowOperation) Init() SLSBridgedTileSpaceSetDividerWindowOperation {
	rv := objc.SendIfResponds[SLSBridgedTileSpaceSetDividerWindowOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedTileSpaceSetDividerWindowOperation) Autorelease() SLSBridgedTileSpaceSetDividerWindowOperation {
	rv := objc.SendIfResponds[SLSBridgedTileSpaceSetDividerWindowOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedTileSpaceSetDividerWindowOperation creates a new SLSBridgedTileSpaceSetDividerWindowOperation instance.
func NewSLSBridgedTileSpaceSetDividerWindowOperation() SLSBridgedTileSpaceSetDividerWindowOperation {
	class := getSLSBridgedTileSpaceSetDividerWindowOperationClass()
	rv := objc.SendIfResponds[SLSBridgedTileSpaceSetDividerWindowOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSLSBridgedTileSpaceSetDividerWindowOperationWithCoder(coder objectivec.IObject) SLSBridgedTileSpaceSetDividerWindowOperation {
	instance := getSLSBridgedTileSpaceSetDividerWindowOperationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedTileSpaceSetDividerWindowOperationFromID(rv)
}

func NewSLSBridgedTileSpaceSetDividerWindowOperationWithSpaceIDDirectionWindowID(id uint64, direction uint64, id2 uint32) SLSBridgedTileSpaceSetDividerWindowOperation {
	instance := getSLSBridgedTileSpaceSetDividerWindowOperationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithSpaceID:direction:windowID:"), id, direction, id2)
	return SLSBridgedTileSpaceSetDividerWindowOperationFromID(rv)
}

func (s SLSBridgedTileSpaceSetDividerWindowOperation) MakeResultWithVerticalIndexHorizontalIndex(index uint64, index2 uint64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("makeResultWithVerticalIndex:horizontalIndex:"), index, index2)
	return objectivec.Object{ID: rv}
}
func (s SLSBridgedTileSpaceSetDividerWindowOperation) InitWithSpaceIDDirectionWindowID(id uint64, direction uint64, id2 uint32) SLSBridgedTileSpaceSetDividerWindowOperation {
	rv := objc.SendIfResponds[SLSBridgedTileSpaceSetDividerWindowOperation](s.ID, objc.Sel("initWithSpaceID:direction:windowID:"), id, direction, id2)
	return rv
}

func (s SLSBridgedTileSpaceSetDividerWindowOperation) Direction() uint64 {
	rv := objc.SendIfResponds[uint64](s.ID, objc.Sel("direction"))
	return rv
}
func (s SLSBridgedTileSpaceSetDividerWindowOperation) SpaceID() uint64 {
	rv := objc.SendIfResponds[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}
func (s SLSBridgedTileSpaceSetDividerWindowOperation) WindowID() uint32 {
	rv := objc.SendIfResponds[uint32](s.ID, objc.Sel("windowID"))
	return rv
}
