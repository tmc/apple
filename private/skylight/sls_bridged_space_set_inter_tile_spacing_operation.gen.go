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
	rv := objc.Send[SLSBridgedSpaceSetInterTileSpacingOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedSpaceSetInterTileSpacingOperation.SpaceID]
//   - [SLSBridgedSpaceSetInterTileSpacingOperation.Spacing]
//   - [SLSBridgedSpaceSetInterTileSpacingOperation.InitWithSpaceIDSpacing]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetInterTileSpacingOperation
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
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetInterTileSpacingOperation
type ISLSBridgedSpaceSetInterTileSpacingOperation interface {
	ISLSAsynchronousBridgedWindowManagementOperation

	// Topic: Methods

	SpaceID() uint64
	Spacing() corefoundation.CGSize
	InitWithSpaceIDSpacing(id uint64, spacing corefoundation.CGSize) SLSBridgedSpaceSetInterTileSpacingOperation
}

// Init initializes the instance.
func (s SLSBridgedSpaceSetInterTileSpacingOperation) Init() SLSBridgedSpaceSetInterTileSpacingOperation {
	rv := objc.Send[SLSBridgedSpaceSetInterTileSpacingOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedSpaceSetInterTileSpacingOperation) Autorelease() SLSBridgedSpaceSetInterTileSpacingOperation {
	rv := objc.Send[SLSBridgedSpaceSetInterTileSpacingOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedSpaceSetInterTileSpacingOperation creates a new SLSBridgedSpaceSetInterTileSpacingOperation instance.
func NewSLSBridgedSpaceSetInterTileSpacingOperation() SLSBridgedSpaceSetInterTileSpacingOperation {
	class := getSLSBridgedSpaceSetInterTileSpacingOperationClass()
	rv := objc.Send[SLSBridgedSpaceSetInterTileSpacingOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetInterTileSpacingOperation/initWithCoder:
func NewSLSBridgedSpaceSetInterTileSpacingOperationWithCoder(coder objectivec.IObject) SLSBridgedSpaceSetInterTileSpacingOperation {
	instance := getSLSBridgedSpaceSetInterTileSpacingOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedSpaceSetInterTileSpacingOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetInterTileSpacingOperation/initWithSpaceID:spacing:
func NewSLSBridgedSpaceSetInterTileSpacingOperationWithSpaceIDSpacing(id uint64, spacing corefoundation.CGSize) SLSBridgedSpaceSetInterTileSpacingOperation {
	instance := getSLSBridgedSpaceSetInterTileSpacingOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpaceID:spacing:"), id, spacing)
	return SLSBridgedSpaceSetInterTileSpacingOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetInterTileSpacingOperation/initWithSpaceID:spacing:
func (s SLSBridgedSpaceSetInterTileSpacingOperation) InitWithSpaceIDSpacing(id uint64, spacing corefoundation.CGSize) SLSBridgedSpaceSetInterTileSpacingOperation {
	rv := objc.Send[SLSBridgedSpaceSetInterTileSpacingOperation](s.ID, objc.Sel("initWithSpaceID:spacing:"), id, spacing)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetInterTileSpacingOperation/spaceID
func (s SLSBridgedSpaceSetInterTileSpacingOperation) SpaceID() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetInterTileSpacingOperation/spacing
func (s SLSBridgedSpaceSetInterTileSpacingOperation) Spacing() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](s.ID, objc.Sel("spacing"))
	return corefoundation.CGSize(rv)
}
