// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedSpaceGetAlphaOperation] class.
var (
	_SLSBridgedSpaceGetAlphaOperationClass     SLSBridgedSpaceGetAlphaOperationClass
	_SLSBridgedSpaceGetAlphaOperationClassOnce sync.Once
)

func getSLSBridgedSpaceGetAlphaOperationClass() SLSBridgedSpaceGetAlphaOperationClass {
	_SLSBridgedSpaceGetAlphaOperationClassOnce.Do(func() {
		_SLSBridgedSpaceGetAlphaOperationClass = SLSBridgedSpaceGetAlphaOperationClass{class: objc.GetClass("SLSBridgedSpaceGetAlphaOperation")}
	})
	return _SLSBridgedSpaceGetAlphaOperationClass
}

// GetSLSBridgedSpaceGetAlphaOperationClass returns the class object for SLSBridgedSpaceGetAlphaOperation.
func GetSLSBridgedSpaceGetAlphaOperationClass() SLSBridgedSpaceGetAlphaOperationClass {
	return getSLSBridgedSpaceGetAlphaOperationClass()
}

type SLSBridgedSpaceGetAlphaOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedSpaceGetAlphaOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedSpaceGetAlphaOperationClass) Alloc() SLSBridgedSpaceGetAlphaOperation {
	rv := objc.Send[SLSBridgedSpaceGetAlphaOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedSpaceGetAlphaOperation.MakeResultWithFloatValue]
//   - [SLSBridgedSpaceGetAlphaOperation.SpaceID]
//   - [SLSBridgedSpaceGetAlphaOperation.InitWithSpaceID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceGetAlphaOperation
type SLSBridgedSpaceGetAlphaOperation struct {
	SLSSynchronousBridgedWindowManagementOperation
}

// SLSBridgedSpaceGetAlphaOperationFromID constructs a [SLSBridgedSpaceGetAlphaOperation] from an objc.ID.
func SLSBridgedSpaceGetAlphaOperationFromID(id objc.ID) SLSBridgedSpaceGetAlphaOperation {
	return SLSBridgedSpaceGetAlphaOperation{SLSSynchronousBridgedWindowManagementOperation: SLSSynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedSpaceGetAlphaOperation implements ISLSBridgedSpaceGetAlphaOperation.
var _ ISLSBridgedSpaceGetAlphaOperation = SLSBridgedSpaceGetAlphaOperation{}

// An interface definition for the [SLSBridgedSpaceGetAlphaOperation] class.
//
// # Methods
//
//   - [ISLSBridgedSpaceGetAlphaOperation.MakeResultWithFloatValue]
//   - [ISLSBridgedSpaceGetAlphaOperation.SpaceID]
//   - [ISLSBridgedSpaceGetAlphaOperation.InitWithSpaceID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceGetAlphaOperation
type ISLSBridgedSpaceGetAlphaOperation interface {
	ISLSSynchronousBridgedWindowManagementOperation

	// Topic: Methods

	MakeResultWithFloatValue(value float32) objectivec.IObject
	SpaceID() uint64
	InitWithSpaceID(id uint64) SLSBridgedSpaceGetAlphaOperation
}

// Init initializes the instance.
func (s SLSBridgedSpaceGetAlphaOperation) Init() SLSBridgedSpaceGetAlphaOperation {
	rv := objc.Send[SLSBridgedSpaceGetAlphaOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedSpaceGetAlphaOperation) Autorelease() SLSBridgedSpaceGetAlphaOperation {
	rv := objc.Send[SLSBridgedSpaceGetAlphaOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedSpaceGetAlphaOperation creates a new SLSBridgedSpaceGetAlphaOperation instance.
func NewSLSBridgedSpaceGetAlphaOperation() SLSBridgedSpaceGetAlphaOperation {
	class := getSLSBridgedSpaceGetAlphaOperationClass()
	rv := objc.Send[SLSBridgedSpaceGetAlphaOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceGetAlphaOperation/initWithCoder:
func NewSLSBridgedSpaceGetAlphaOperationWithCoder(coder objectivec.IObject) SLSBridgedSpaceGetAlphaOperation {
	instance := getSLSBridgedSpaceGetAlphaOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedSpaceGetAlphaOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceGetAlphaOperation/initWithSpaceID:
func NewSLSBridgedSpaceGetAlphaOperationWithSpaceID(id uint64) SLSBridgedSpaceGetAlphaOperation {
	instance := getSLSBridgedSpaceGetAlphaOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpaceID:"), id)
	return SLSBridgedSpaceGetAlphaOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceGetAlphaOperation/makeResultWithFloatValue:
func (s SLSBridgedSpaceGetAlphaOperation) MakeResultWithFloatValue(value float32) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("makeResultWithFloatValue:"), value)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceGetAlphaOperation/initWithSpaceID:
func (s SLSBridgedSpaceGetAlphaOperation) InitWithSpaceID(id uint64) SLSBridgedSpaceGetAlphaOperation {
	rv := objc.Send[SLSBridgedSpaceGetAlphaOperation](s.ID, objc.Sel("initWithSpaceID:"), id)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceGetAlphaOperation/spaceID
func (s SLSBridgedSpaceGetAlphaOperation) SpaceID() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}
