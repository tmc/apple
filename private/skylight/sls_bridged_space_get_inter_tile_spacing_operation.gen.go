// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedSpaceGetInterTileSpacingOperation] class.
var (
	_SLSBridgedSpaceGetInterTileSpacingOperationClass     SLSBridgedSpaceGetInterTileSpacingOperationClass
	_SLSBridgedSpaceGetInterTileSpacingOperationClassOnce sync.Once
)

func getSLSBridgedSpaceGetInterTileSpacingOperationClass() SLSBridgedSpaceGetInterTileSpacingOperationClass {
	_SLSBridgedSpaceGetInterTileSpacingOperationClassOnce.Do(func() {
		_SLSBridgedSpaceGetInterTileSpacingOperationClass = SLSBridgedSpaceGetInterTileSpacingOperationClass{class: objc.GetClass("SLSBridgedSpaceGetInterTileSpacingOperation")}
	})
	return _SLSBridgedSpaceGetInterTileSpacingOperationClass
}

// GetSLSBridgedSpaceGetInterTileSpacingOperationClass returns the class object for SLSBridgedSpaceGetInterTileSpacingOperation.
func GetSLSBridgedSpaceGetInterTileSpacingOperationClass() SLSBridgedSpaceGetInterTileSpacingOperationClass {
	return getSLSBridgedSpaceGetInterTileSpacingOperationClass()
}

type SLSBridgedSpaceGetInterTileSpacingOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedSpaceGetInterTileSpacingOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedSpaceGetInterTileSpacingOperationClass) Alloc() SLSBridgedSpaceGetInterTileSpacingOperation {
	rv := objc.Send[SLSBridgedSpaceGetInterTileSpacingOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedSpaceGetInterTileSpacingOperation.MakeResultWithSize]
//   - [SLSBridgedSpaceGetInterTileSpacingOperation.SpaceID]
//   - [SLSBridgedSpaceGetInterTileSpacingOperation.InitWithSpaceID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceGetInterTileSpacingOperation
type SLSBridgedSpaceGetInterTileSpacingOperation struct {
	SLSSynchronousBridgedWindowManagementOperation
}

// SLSBridgedSpaceGetInterTileSpacingOperationFromID constructs a [SLSBridgedSpaceGetInterTileSpacingOperation] from an objc.ID.
func SLSBridgedSpaceGetInterTileSpacingOperationFromID(id objc.ID) SLSBridgedSpaceGetInterTileSpacingOperation {
	return SLSBridgedSpaceGetInterTileSpacingOperation{SLSSynchronousBridgedWindowManagementOperation: SLSSynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedSpaceGetInterTileSpacingOperation implements ISLSBridgedSpaceGetInterTileSpacingOperation.
var _ ISLSBridgedSpaceGetInterTileSpacingOperation = SLSBridgedSpaceGetInterTileSpacingOperation{}

// An interface definition for the [SLSBridgedSpaceGetInterTileSpacingOperation] class.
//
// # Methods
//
//   - [ISLSBridgedSpaceGetInterTileSpacingOperation.MakeResultWithSize]
//   - [ISLSBridgedSpaceGetInterTileSpacingOperation.SpaceID]
//   - [ISLSBridgedSpaceGetInterTileSpacingOperation.InitWithSpaceID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceGetInterTileSpacingOperation
type ISLSBridgedSpaceGetInterTileSpacingOperation interface {
	ISLSSynchronousBridgedWindowManagementOperation

	// Topic: Methods

	MakeResultWithSize(size corefoundation.CGSize) objectivec.IObject
	SpaceID() uint64
	InitWithSpaceID(id uint64) SLSBridgedSpaceGetInterTileSpacingOperation
}

// Init initializes the instance.
func (s SLSBridgedSpaceGetInterTileSpacingOperation) Init() SLSBridgedSpaceGetInterTileSpacingOperation {
	rv := objc.Send[SLSBridgedSpaceGetInterTileSpacingOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedSpaceGetInterTileSpacingOperation) Autorelease() SLSBridgedSpaceGetInterTileSpacingOperation {
	rv := objc.Send[SLSBridgedSpaceGetInterTileSpacingOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedSpaceGetInterTileSpacingOperation creates a new SLSBridgedSpaceGetInterTileSpacingOperation instance.
func NewSLSBridgedSpaceGetInterTileSpacingOperation() SLSBridgedSpaceGetInterTileSpacingOperation {
	class := getSLSBridgedSpaceGetInterTileSpacingOperationClass()
	rv := objc.Send[SLSBridgedSpaceGetInterTileSpacingOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceGetInterTileSpacingOperation/initWithCoder:
func NewSLSBridgedSpaceGetInterTileSpacingOperationWithCoder(coder objectivec.IObject) SLSBridgedSpaceGetInterTileSpacingOperation {
	instance := getSLSBridgedSpaceGetInterTileSpacingOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedSpaceGetInterTileSpacingOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceGetInterTileSpacingOperation/initWithSpaceID:
func NewSLSBridgedSpaceGetInterTileSpacingOperationWithSpaceID(id uint64) SLSBridgedSpaceGetInterTileSpacingOperation {
	instance := getSLSBridgedSpaceGetInterTileSpacingOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpaceID:"), id)
	return SLSBridgedSpaceGetInterTileSpacingOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceGetInterTileSpacingOperation/makeResultWithSize:
func (s SLSBridgedSpaceGetInterTileSpacingOperation) MakeResultWithSize(size corefoundation.CGSize) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("makeResultWithSize:"), size)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceGetInterTileSpacingOperation/initWithSpaceID:
func (s SLSBridgedSpaceGetInterTileSpacingOperation) InitWithSpaceID(id uint64) SLSBridgedSpaceGetInterTileSpacingOperation {
	rv := objc.Send[SLSBridgedSpaceGetInterTileSpacingOperation](s.ID, objc.Sel("initWithSpaceID:"), id)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceGetInterTileSpacingOperation/spaceID
func (s SLSBridgedSpaceGetInterTileSpacingOperation) SpaceID() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}
