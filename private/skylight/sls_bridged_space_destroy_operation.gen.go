// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedSpaceDestroyOperation] class.
var (
	_SLSBridgedSpaceDestroyOperationClass     SLSBridgedSpaceDestroyOperationClass
	_SLSBridgedSpaceDestroyOperationClassOnce sync.Once
)

func getSLSBridgedSpaceDestroyOperationClass() SLSBridgedSpaceDestroyOperationClass {
	_SLSBridgedSpaceDestroyOperationClassOnce.Do(func() {
		_SLSBridgedSpaceDestroyOperationClass = SLSBridgedSpaceDestroyOperationClass{class: objc.GetClass("SLSBridgedSpaceDestroyOperation")}
	})
	return _SLSBridgedSpaceDestroyOperationClass
}

// GetSLSBridgedSpaceDestroyOperationClass returns the class object for SLSBridgedSpaceDestroyOperation.
func GetSLSBridgedSpaceDestroyOperationClass() SLSBridgedSpaceDestroyOperationClass {
	return getSLSBridgedSpaceDestroyOperationClass()
}

type SLSBridgedSpaceDestroyOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedSpaceDestroyOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedSpaceDestroyOperationClass) Alloc() SLSBridgedSpaceDestroyOperation {
	rv := objc.Send[SLSBridgedSpaceDestroyOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedSpaceDestroyOperation.SpaceID]
//   - [SLSBridgedSpaceDestroyOperation.InitWithSpaceID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceDestroyOperation
type SLSBridgedSpaceDestroyOperation struct {
	SLSAsynchronousBridgedWindowManagementOperation
}

// SLSBridgedSpaceDestroyOperationFromID constructs a [SLSBridgedSpaceDestroyOperation] from an objc.ID.
func SLSBridgedSpaceDestroyOperationFromID(id objc.ID) SLSBridgedSpaceDestroyOperation {
	return SLSBridgedSpaceDestroyOperation{SLSAsynchronousBridgedWindowManagementOperation: SLSAsynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedSpaceDestroyOperation implements ISLSBridgedSpaceDestroyOperation.
var _ ISLSBridgedSpaceDestroyOperation = SLSBridgedSpaceDestroyOperation{}

// An interface definition for the [SLSBridgedSpaceDestroyOperation] class.
//
// # Methods
//
//   - [ISLSBridgedSpaceDestroyOperation.SpaceID]
//   - [ISLSBridgedSpaceDestroyOperation.InitWithSpaceID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceDestroyOperation
type ISLSBridgedSpaceDestroyOperation interface {
	ISLSAsynchronousBridgedWindowManagementOperation

	// Topic: Methods

	SpaceID() uint64
	InitWithSpaceID(id uint64) SLSBridgedSpaceDestroyOperation
}

// Init initializes the instance.
func (s SLSBridgedSpaceDestroyOperation) Init() SLSBridgedSpaceDestroyOperation {
	rv := objc.Send[SLSBridgedSpaceDestroyOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedSpaceDestroyOperation) Autorelease() SLSBridgedSpaceDestroyOperation {
	rv := objc.Send[SLSBridgedSpaceDestroyOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedSpaceDestroyOperation creates a new SLSBridgedSpaceDestroyOperation instance.
func NewSLSBridgedSpaceDestroyOperation() SLSBridgedSpaceDestroyOperation {
	class := getSLSBridgedSpaceDestroyOperationClass()
	rv := objc.Send[SLSBridgedSpaceDestroyOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceDestroyOperation/initWithCoder:
func NewSLSBridgedSpaceDestroyOperationWithCoder(coder objectivec.IObject) SLSBridgedSpaceDestroyOperation {
	instance := getSLSBridgedSpaceDestroyOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedSpaceDestroyOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceDestroyOperation/initWithSpaceID:
func NewSLSBridgedSpaceDestroyOperationWithSpaceID(id uint64) SLSBridgedSpaceDestroyOperation {
	instance := getSLSBridgedSpaceDestroyOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpaceID:"), id)
	return SLSBridgedSpaceDestroyOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceDestroyOperation/initWithSpaceID:
func (s SLSBridgedSpaceDestroyOperation) InitWithSpaceID(id uint64) SLSBridgedSpaceDestroyOperation {
	rv := objc.Send[SLSBridgedSpaceDestroyOperation](s.ID, objc.Sel("initWithSpaceID:"), id)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceDestroyOperation/spaceID
func (s SLSBridgedSpaceDestroyOperation) SpaceID() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}
