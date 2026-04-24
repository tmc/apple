// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedSpaceSetAlphaOperation] class.
var (
	_SLSBridgedSpaceSetAlphaOperationClass     SLSBridgedSpaceSetAlphaOperationClass
	_SLSBridgedSpaceSetAlphaOperationClassOnce sync.Once
)

func getSLSBridgedSpaceSetAlphaOperationClass() SLSBridgedSpaceSetAlphaOperationClass {
	_SLSBridgedSpaceSetAlphaOperationClassOnce.Do(func() {
		_SLSBridgedSpaceSetAlphaOperationClass = SLSBridgedSpaceSetAlphaOperationClass{class: objc.GetClass("SLSBridgedSpaceSetAlphaOperation")}
	})
	return _SLSBridgedSpaceSetAlphaOperationClass
}

// GetSLSBridgedSpaceSetAlphaOperationClass returns the class object for SLSBridgedSpaceSetAlphaOperation.
func GetSLSBridgedSpaceSetAlphaOperationClass() SLSBridgedSpaceSetAlphaOperationClass {
	return getSLSBridgedSpaceSetAlphaOperationClass()
}

type SLSBridgedSpaceSetAlphaOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedSpaceSetAlphaOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedSpaceSetAlphaOperationClass) Alloc() SLSBridgedSpaceSetAlphaOperation {
	rv := objc.Send[SLSBridgedSpaceSetAlphaOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedSpaceSetAlphaOperation.Alpha]
//   - [SLSBridgedSpaceSetAlphaOperation.SpaceID]
//   - [SLSBridgedSpaceSetAlphaOperation.InitWithSpaceIDAlpha]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetAlphaOperation
type SLSBridgedSpaceSetAlphaOperation struct {
	SLSAsynchronousBridgedWindowManagementOperation
}

// SLSBridgedSpaceSetAlphaOperationFromID constructs a [SLSBridgedSpaceSetAlphaOperation] from an objc.ID.
func SLSBridgedSpaceSetAlphaOperationFromID(id objc.ID) SLSBridgedSpaceSetAlphaOperation {
	return SLSBridgedSpaceSetAlphaOperation{SLSAsynchronousBridgedWindowManagementOperation: SLSAsynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedSpaceSetAlphaOperation implements ISLSBridgedSpaceSetAlphaOperation.
var _ ISLSBridgedSpaceSetAlphaOperation = SLSBridgedSpaceSetAlphaOperation{}

// An interface definition for the [SLSBridgedSpaceSetAlphaOperation] class.
//
// # Methods
//
//   - [ISLSBridgedSpaceSetAlphaOperation.Alpha]
//   - [ISLSBridgedSpaceSetAlphaOperation.SpaceID]
//   - [ISLSBridgedSpaceSetAlphaOperation.InitWithSpaceIDAlpha]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetAlphaOperation
type ISLSBridgedSpaceSetAlphaOperation interface {
	ISLSAsynchronousBridgedWindowManagementOperation

	// Topic: Methods

	Alpha() float32
	SpaceID() uint64
	InitWithSpaceIDAlpha(id uint64, alpha float32) SLSBridgedSpaceSetAlphaOperation
}

// Init initializes the instance.
func (s SLSBridgedSpaceSetAlphaOperation) Init() SLSBridgedSpaceSetAlphaOperation {
	rv := objc.Send[SLSBridgedSpaceSetAlphaOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedSpaceSetAlphaOperation) Autorelease() SLSBridgedSpaceSetAlphaOperation {
	rv := objc.Send[SLSBridgedSpaceSetAlphaOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedSpaceSetAlphaOperation creates a new SLSBridgedSpaceSetAlphaOperation instance.
func NewSLSBridgedSpaceSetAlphaOperation() SLSBridgedSpaceSetAlphaOperation {
	class := getSLSBridgedSpaceSetAlphaOperationClass()
	rv := objc.Send[SLSBridgedSpaceSetAlphaOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetAlphaOperation/initWithCoder:
func NewSLSBridgedSpaceSetAlphaOperationWithCoder(coder objectivec.IObject) SLSBridgedSpaceSetAlphaOperation {
	instance := getSLSBridgedSpaceSetAlphaOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedSpaceSetAlphaOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetAlphaOperation/initWithSpaceID:alpha:
func NewSLSBridgedSpaceSetAlphaOperationWithSpaceIDAlpha(id uint64, alpha float32) SLSBridgedSpaceSetAlphaOperation {
	instance := getSLSBridgedSpaceSetAlphaOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpaceID:alpha:"), id, alpha)
	return SLSBridgedSpaceSetAlphaOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetAlphaOperation/initWithSpaceID:alpha:
func (s SLSBridgedSpaceSetAlphaOperation) InitWithSpaceIDAlpha(id uint64, alpha float32) SLSBridgedSpaceSetAlphaOperation {
	rv := objc.Send[SLSBridgedSpaceSetAlphaOperation](s.ID, objc.Sel("initWithSpaceID:alpha:"), id, alpha)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetAlphaOperation/alpha
func (s SLSBridgedSpaceSetAlphaOperation) Alpha() float32 {
	rv := objc.Send[float32](s.ID, objc.Sel("alpha"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetAlphaOperation/spaceID
func (s SLSBridgedSpaceSetAlphaOperation) SpaceID() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}
