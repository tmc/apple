// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedSpaceSetInterTileSpacingOperation] class.
var (
	_SLSBridgedSpaceSetInterTileSpacingOperationClass     SLSBridgedSpaceSetInterTileSpacingOperationClass
	_SLSBridgedSpaceSetInterTileSpacingOperationClassOnce sync.Once
)

func getSLSBridgedSpaceSetInterTileSpacingOperationClass() SLSBridgedSpaceSetInterTileSpacingOperationClass {
	_SLSBridgedSpaceSetInterTileSpacingOperationClassOnce.Do(func() {
		_SLSBridgedSpaceSetInterTileSpacingOperationClass = SLSBridgedSpaceSetInterTileSpacingOperationClass{class: objc.GetClass("SLSBridgedSpaceSetInterTileSpacingOperation")}
	})
	return _SLSBridgedSpaceSetInterTileSpacingOperationClass
}

// GetSLSBridgedSpaceSetInterTileSpacingOperationClass returns the class object for SLSBridgedSpaceSetInterTileSpacingOperation.
func GetSLSBridgedSpaceSetInterTileSpacingOperationClass() SLSBridgedSpaceSetInterTileSpacingOperationClass {
	return getSLSBridgedSpaceSetInterTileSpacingOperationClass()
}

type SLSBridgedSpaceSetInterTileSpacingOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedSpaceSetInterTileSpacingOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedSpaceSetInterTileSpacingOperationClass) Alloc() SLSBridgedSpaceSetInterTileSpacingOperation {
	rv := objc.SendIfResponds[SLSBridgedSpaceSetInterTileSpacingOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedSpaceSetInterTileSpacingOperation.SpaceID]
//   - [SLSBridgedSpaceSetInterTileSpacingOperation.Spacing]
//   - [SLSBridgedSpaceSetInterTileSpacingOperation.InitWithSpaceIDSpacing]
type SLSBridgedSpaceSetInterTileSpacingOperation struct {
	SLSAsynchronousBridgedWindowManagementOperation
}

// SLSBridgedSpaceSetInterTileSpacingOperationFromID constructs a [SLSBridgedSpaceSetInterTileSpacingOperation] from an objc.ID.
func SLSBridgedSpaceSetInterTileSpacingOperationFromID(id objc.ID) SLSBridgedSpaceSetInterTileSpacingOperation {
	return SLSBridgedSpaceSetInterTileSpacingOperation{SLSAsynchronousBridgedWindowManagementOperation: SLSAsynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedSpaceSetInterTileSpacingOperation implements ISLSBridgedSpaceSetInterTileSpacingOperation.
var _ ISLSBridgedSpaceSetInterTileSpacingOperation = SLSBridgedSpaceSetInterTileSpacingOperation{}

// An interface definition for the [SLSBridgedSpaceSetInterTileSpacingOperation] class.
//
// # Methods
//
//   - [ISLSBridgedSpaceSetInterTileSpacingOperation.SpaceID]
//   - [ISLSBridgedSpaceSetInterTileSpacingOperation.Spacing]
//   - [ISLSBridgedSpaceSetInterTileSpacingOperation.InitWithSpaceIDSpacing]
type ISLSBridgedSpaceSetInterTileSpacingOperation interface {
	ISLSAsynchronousBridgedWindowManagementOperation

	// Topic: Methods

	SpaceID() uint64
	Spacing() corefoundation.CGSize
	InitWithSpaceIDSpacing(id uint64, spacing corefoundation.CGSize) SLSBridgedSpaceSetInterTileSpacingOperation
}

// Init initializes the instance.
func (s SLSBridgedSpaceSetInterTileSpacingOperation) Init() SLSBridgedSpaceSetInterTileSpacingOperation {
	rv := objc.SendIfResponds[SLSBridgedSpaceSetInterTileSpacingOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedSpaceSetInterTileSpacingOperation) Autorelease() SLSBridgedSpaceSetInterTileSpacingOperation {
	rv := objc.SendIfResponds[SLSBridgedSpaceSetInterTileSpacingOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedSpaceSetInterTileSpacingOperation creates a new SLSBridgedSpaceSetInterTileSpacingOperation instance.
func NewSLSBridgedSpaceSetInterTileSpacingOperation() SLSBridgedSpaceSetInterTileSpacingOperation {
	class := getSLSBridgedSpaceSetInterTileSpacingOperationClass()
	rv := objc.SendIfResponds[SLSBridgedSpaceSetInterTileSpacingOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSLSBridgedSpaceSetInterTileSpacingOperationWithCoder(coder objectivec.IObject) SLSBridgedSpaceSetInterTileSpacingOperation {
	instance := getSLSBridgedSpaceSetInterTileSpacingOperationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedSpaceSetInterTileSpacingOperationFromID(rv)
}

func NewSLSBridgedSpaceSetInterTileSpacingOperationWithSpaceIDSpacing(id uint64, spacing corefoundation.CGSize) SLSBridgedSpaceSetInterTileSpacingOperation {
	instance := getSLSBridgedSpaceSetInterTileSpacingOperationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithSpaceID:spacing:"), id, spacing)
	return SLSBridgedSpaceSetInterTileSpacingOperationFromID(rv)
}

func (s SLSBridgedSpaceSetInterTileSpacingOperation) InitWithSpaceIDSpacing(id uint64, spacing corefoundation.CGSize) SLSBridgedSpaceSetInterTileSpacingOperation {
	rv := objc.SendIfResponds[SLSBridgedSpaceSetInterTileSpacingOperation](s.ID, objc.Sel("initWithSpaceID:spacing:"), id, spacing)
	return rv
}

func (s SLSBridgedSpaceSetInterTileSpacingOperation) SpaceID() uint64 {
	rv := objc.SendIfResponds[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}
func (s SLSBridgedSpaceSetInterTileSpacingOperation) Spacing() corefoundation.CGSize {
	rv := objc.SendIfResponds[corefoundation.CGSize](s.ID, objc.Sel("spacing"))
	return corefoundation.CGSize(rv)
}
