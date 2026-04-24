// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedSpaceGetTransformOperation] class.
var (
	_SLSBridgedSpaceGetTransformOperationClass     SLSBridgedSpaceGetTransformOperationClass
	_SLSBridgedSpaceGetTransformOperationClassOnce sync.Once
)

func getSLSBridgedSpaceGetTransformOperationClass() SLSBridgedSpaceGetTransformOperationClass {
	_SLSBridgedSpaceGetTransformOperationClassOnce.Do(func() {
		_SLSBridgedSpaceGetTransformOperationClass = SLSBridgedSpaceGetTransformOperationClass{class: objc.GetClass("SLSBridgedSpaceGetTransformOperation")}
	})
	return _SLSBridgedSpaceGetTransformOperationClass
}

// GetSLSBridgedSpaceGetTransformOperationClass returns the class object for SLSBridgedSpaceGetTransformOperation.
func GetSLSBridgedSpaceGetTransformOperationClass() SLSBridgedSpaceGetTransformOperationClass {
	return getSLSBridgedSpaceGetTransformOperationClass()
}

type SLSBridgedSpaceGetTransformOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedSpaceGetTransformOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedSpaceGetTransformOperationClass) Alloc() SLSBridgedSpaceGetTransformOperation {
	rv := objc.Send[SLSBridgedSpaceGetTransformOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedSpaceGetTransformOperation.MakeResultWithAffineTransformOptions]
//   - [SLSBridgedSpaceGetTransformOperation.SpaceID]
//   - [SLSBridgedSpaceGetTransformOperation.InitWithSpaceID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceGetTransformOperation
type SLSBridgedSpaceGetTransformOperation struct {
	SLSSynchronousBridgedWindowManagementOperation
}

// SLSBridgedSpaceGetTransformOperationFromID constructs a [SLSBridgedSpaceGetTransformOperation] from an objc.ID.
func SLSBridgedSpaceGetTransformOperationFromID(id objc.ID) SLSBridgedSpaceGetTransformOperation {
	return SLSBridgedSpaceGetTransformOperation{SLSSynchronousBridgedWindowManagementOperation: SLSSynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedSpaceGetTransformOperation implements ISLSBridgedSpaceGetTransformOperation.
var _ ISLSBridgedSpaceGetTransformOperation = SLSBridgedSpaceGetTransformOperation{}

// An interface definition for the [SLSBridgedSpaceGetTransformOperation] class.
//
// # Methods
//
//   - [ISLSBridgedSpaceGetTransformOperation.MakeResultWithAffineTransformOptions]
//   - [ISLSBridgedSpaceGetTransformOperation.SpaceID]
//   - [ISLSBridgedSpaceGetTransformOperation.InitWithSpaceID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceGetTransformOperation
type ISLSBridgedSpaceGetTransformOperation interface {
	ISLSSynchronousBridgedWindowManagementOperation

	// Topic: Methods

	MakeResultWithAffineTransformOptions(transform corefoundation.CGAffineTransform, options uint32) objectivec.IObject
	SpaceID() uint64
	InitWithSpaceID(id uint64) SLSBridgedSpaceGetTransformOperation
}

// Init initializes the instance.
func (s SLSBridgedSpaceGetTransformOperation) Init() SLSBridgedSpaceGetTransformOperation {
	rv := objc.Send[SLSBridgedSpaceGetTransformOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedSpaceGetTransformOperation) Autorelease() SLSBridgedSpaceGetTransformOperation {
	rv := objc.Send[SLSBridgedSpaceGetTransformOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedSpaceGetTransformOperation creates a new SLSBridgedSpaceGetTransformOperation instance.
func NewSLSBridgedSpaceGetTransformOperation() SLSBridgedSpaceGetTransformOperation {
	class := getSLSBridgedSpaceGetTransformOperationClass()
	rv := objc.Send[SLSBridgedSpaceGetTransformOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceGetTransformOperation/initWithCoder:
func NewSLSBridgedSpaceGetTransformOperationWithCoder(coder objectivec.IObject) SLSBridgedSpaceGetTransformOperation {
	instance := getSLSBridgedSpaceGetTransformOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedSpaceGetTransformOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceGetTransformOperation/initWithSpaceID:
func NewSLSBridgedSpaceGetTransformOperationWithSpaceID(id uint64) SLSBridgedSpaceGetTransformOperation {
	instance := getSLSBridgedSpaceGetTransformOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpaceID:"), id)
	return SLSBridgedSpaceGetTransformOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceGetTransformOperation/makeResultWithAffineTransform:options:
func (s SLSBridgedSpaceGetTransformOperation) MakeResultWithAffineTransformOptions(transform corefoundation.CGAffineTransform, options uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("makeResultWithAffineTransform:options:"), transform, options)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceGetTransformOperation/initWithSpaceID:
func (s SLSBridgedSpaceGetTransformOperation) InitWithSpaceID(id uint64) SLSBridgedSpaceGetTransformOperation {
	rv := objc.Send[SLSBridgedSpaceGetTransformOperation](s.ID, objc.Sel("initWithSpaceID:"), id)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceGetTransformOperation/spaceID
func (s SLSBridgedSpaceGetTransformOperation) SpaceID() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}
