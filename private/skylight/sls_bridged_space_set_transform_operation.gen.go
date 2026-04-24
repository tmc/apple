// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedSpaceSetTransformOperation] class.
var (
	_SLSBridgedSpaceSetTransformOperationClass     SLSBridgedSpaceSetTransformOperationClass
	_SLSBridgedSpaceSetTransformOperationClassOnce sync.Once
)

func getSLSBridgedSpaceSetTransformOperationClass() SLSBridgedSpaceSetTransformOperationClass {
	_SLSBridgedSpaceSetTransformOperationClassOnce.Do(func() {
		_SLSBridgedSpaceSetTransformOperationClass = SLSBridgedSpaceSetTransformOperationClass{class: objc.GetClass("SLSBridgedSpaceSetTransformOperation")}
	})
	return _SLSBridgedSpaceSetTransformOperationClass
}

// GetSLSBridgedSpaceSetTransformOperationClass returns the class object for SLSBridgedSpaceSetTransformOperation.
func GetSLSBridgedSpaceSetTransformOperationClass() SLSBridgedSpaceSetTransformOperationClass {
	return getSLSBridgedSpaceSetTransformOperationClass()
}

type SLSBridgedSpaceSetTransformOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedSpaceSetTransformOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedSpaceSetTransformOperationClass) Alloc() SLSBridgedSpaceSetTransformOperation {
	rv := objc.Send[SLSBridgedSpaceSetTransformOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedSpaceSetTransformOperation.Options]
//   - [SLSBridgedSpaceSetTransformOperation.SpaceID]
//   - [SLSBridgedSpaceSetTransformOperation.Transform]
//   - [SLSBridgedSpaceSetTransformOperation.InitWithSpaceIDTransformOptions]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetTransformOperation
type SLSBridgedSpaceSetTransformOperation struct {
	SLSAsynchronousBridgedWindowManagementOperation
}

// SLSBridgedSpaceSetTransformOperationFromID constructs a [SLSBridgedSpaceSetTransformOperation] from an objc.ID.
func SLSBridgedSpaceSetTransformOperationFromID(id objc.ID) SLSBridgedSpaceSetTransformOperation {
	return SLSBridgedSpaceSetTransformOperation{SLSAsynchronousBridgedWindowManagementOperation: SLSAsynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedSpaceSetTransformOperation implements ISLSBridgedSpaceSetTransformOperation.
var _ ISLSBridgedSpaceSetTransformOperation = SLSBridgedSpaceSetTransformOperation{}

// An interface definition for the [SLSBridgedSpaceSetTransformOperation] class.
//
// # Methods
//
//   - [ISLSBridgedSpaceSetTransformOperation.Options]
//   - [ISLSBridgedSpaceSetTransformOperation.SpaceID]
//   - [ISLSBridgedSpaceSetTransformOperation.Transform]
//   - [ISLSBridgedSpaceSetTransformOperation.InitWithSpaceIDTransformOptions]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetTransformOperation
type ISLSBridgedSpaceSetTransformOperation interface {
	ISLSAsynchronousBridgedWindowManagementOperation

	// Topic: Methods

	Options() uint32
	SpaceID() uint64
	Transform() corefoundation.CGAffineTransform
	InitWithSpaceIDTransformOptions(id uint64, transform corefoundation.CGAffineTransform, options uint32) SLSBridgedSpaceSetTransformOperation
}

// Init initializes the instance.
func (s SLSBridgedSpaceSetTransformOperation) Init() SLSBridgedSpaceSetTransformOperation {
	rv := objc.Send[SLSBridgedSpaceSetTransformOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedSpaceSetTransformOperation) Autorelease() SLSBridgedSpaceSetTransformOperation {
	rv := objc.Send[SLSBridgedSpaceSetTransformOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedSpaceSetTransformOperation creates a new SLSBridgedSpaceSetTransformOperation instance.
func NewSLSBridgedSpaceSetTransformOperation() SLSBridgedSpaceSetTransformOperation {
	class := getSLSBridgedSpaceSetTransformOperationClass()
	rv := objc.Send[SLSBridgedSpaceSetTransformOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetTransformOperation/initWithCoder:
func NewSLSBridgedSpaceSetTransformOperationWithCoder(coder objectivec.IObject) SLSBridgedSpaceSetTransformOperation {
	instance := getSLSBridgedSpaceSetTransformOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedSpaceSetTransformOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetTransformOperation/initWithSpaceID:transform:options:
func NewSLSBridgedSpaceSetTransformOperationWithSpaceIDTransformOptions(id uint64, transform corefoundation.CGAffineTransform, options uint32) SLSBridgedSpaceSetTransformOperation {
	instance := getSLSBridgedSpaceSetTransformOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpaceID:transform:options:"), id, transform, options)
	return SLSBridgedSpaceSetTransformOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetTransformOperation/initWithSpaceID:transform:options:
func (s SLSBridgedSpaceSetTransformOperation) InitWithSpaceIDTransformOptions(id uint64, transform corefoundation.CGAffineTransform, options uint32) SLSBridgedSpaceSetTransformOperation {
	rv := objc.Send[SLSBridgedSpaceSetTransformOperation](s.ID, objc.Sel("initWithSpaceID:transform:options:"), id, transform, options)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetTransformOperation/options
func (s SLSBridgedSpaceSetTransformOperation) Options() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("options"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetTransformOperation/spaceID
func (s SLSBridgedSpaceSetTransformOperation) SpaceID() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetTransformOperation/transform
func (s SLSBridgedSpaceSetTransformOperation) Transform() corefoundation.CGAffineTransform {
	rv := objc.Send[corefoundation.CGAffineTransform](s.ID, objc.Sel("transform"))
	return corefoundation.CGAffineTransform(rv)
}
