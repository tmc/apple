// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedSpaceResetMenuBarOperation] class.
var (
	_SLSBridgedSpaceResetMenuBarOperationClass     SLSBridgedSpaceResetMenuBarOperationClass
	_SLSBridgedSpaceResetMenuBarOperationClassOnce sync.Once
)

func getSLSBridgedSpaceResetMenuBarOperationClass() SLSBridgedSpaceResetMenuBarOperationClass {
	_SLSBridgedSpaceResetMenuBarOperationClassOnce.Do(func() {
		_SLSBridgedSpaceResetMenuBarOperationClass = SLSBridgedSpaceResetMenuBarOperationClass{class: objc.GetClass("SLSBridgedSpaceResetMenuBarOperation")}
	})
	return _SLSBridgedSpaceResetMenuBarOperationClass
}

// GetSLSBridgedSpaceResetMenuBarOperationClass returns the class object for SLSBridgedSpaceResetMenuBarOperation.
func GetSLSBridgedSpaceResetMenuBarOperationClass() SLSBridgedSpaceResetMenuBarOperationClass {
	return getSLSBridgedSpaceResetMenuBarOperationClass()
}

type SLSBridgedSpaceResetMenuBarOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedSpaceResetMenuBarOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedSpaceResetMenuBarOperationClass) Alloc() SLSBridgedSpaceResetMenuBarOperation {
	rv := objc.Send[SLSBridgedSpaceResetMenuBarOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedSpaceResetMenuBarOperation.SpaceID]
//   - [SLSBridgedSpaceResetMenuBarOperation.InitWithSpaceID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceResetMenuBarOperation
type SLSBridgedSpaceResetMenuBarOperation struct {
	SLSAsynchronousBridgedWindowManagementOperation
}

// SLSBridgedSpaceResetMenuBarOperationFromID constructs a [SLSBridgedSpaceResetMenuBarOperation] from an objc.ID.
func SLSBridgedSpaceResetMenuBarOperationFromID(id objc.ID) SLSBridgedSpaceResetMenuBarOperation {
	return SLSBridgedSpaceResetMenuBarOperation{SLSAsynchronousBridgedWindowManagementOperation: SLSAsynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedSpaceResetMenuBarOperation implements ISLSBridgedSpaceResetMenuBarOperation.
var _ ISLSBridgedSpaceResetMenuBarOperation = SLSBridgedSpaceResetMenuBarOperation{}

// An interface definition for the [SLSBridgedSpaceResetMenuBarOperation] class.
//
// # Methods
//
//   - [ISLSBridgedSpaceResetMenuBarOperation.SpaceID]
//   - [ISLSBridgedSpaceResetMenuBarOperation.InitWithSpaceID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceResetMenuBarOperation
type ISLSBridgedSpaceResetMenuBarOperation interface {
	ISLSAsynchronousBridgedWindowManagementOperation

	// Topic: Methods

	SpaceID() uint64
	InitWithSpaceID(id uint64) SLSBridgedSpaceResetMenuBarOperation
}

// Init initializes the instance.
func (s SLSBridgedSpaceResetMenuBarOperation) Init() SLSBridgedSpaceResetMenuBarOperation {
	rv := objc.Send[SLSBridgedSpaceResetMenuBarOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedSpaceResetMenuBarOperation) Autorelease() SLSBridgedSpaceResetMenuBarOperation {
	rv := objc.Send[SLSBridgedSpaceResetMenuBarOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedSpaceResetMenuBarOperation creates a new SLSBridgedSpaceResetMenuBarOperation instance.
func NewSLSBridgedSpaceResetMenuBarOperation() SLSBridgedSpaceResetMenuBarOperation {
	class := getSLSBridgedSpaceResetMenuBarOperationClass()
	rv := objc.Send[SLSBridgedSpaceResetMenuBarOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceResetMenuBarOperation/initWithCoder:
func NewSLSBridgedSpaceResetMenuBarOperationWithCoder(coder objectivec.IObject) SLSBridgedSpaceResetMenuBarOperation {
	instance := getSLSBridgedSpaceResetMenuBarOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedSpaceResetMenuBarOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceResetMenuBarOperation/initWithSpaceID:
func NewSLSBridgedSpaceResetMenuBarOperationWithSpaceID(id uint64) SLSBridgedSpaceResetMenuBarOperation {
	instance := getSLSBridgedSpaceResetMenuBarOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpaceID:"), id)
	return SLSBridgedSpaceResetMenuBarOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceResetMenuBarOperation/initWithSpaceID:
func (s SLSBridgedSpaceResetMenuBarOperation) InitWithSpaceID(id uint64) SLSBridgedSpaceResetMenuBarOperation {
	rv := objc.Send[SLSBridgedSpaceResetMenuBarOperation](s.ID, objc.Sel("initWithSpaceID:"), id)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceResetMenuBarOperation/spaceID
func (s SLSBridgedSpaceResetMenuBarOperation) SpaceID() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}
