// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedSpaceSetShapeOperation] class.
var (
	_SLSBridgedSpaceSetShapeOperationClass     SLSBridgedSpaceSetShapeOperationClass
	_SLSBridgedSpaceSetShapeOperationClassOnce sync.Once
)

func getSLSBridgedSpaceSetShapeOperationClass() SLSBridgedSpaceSetShapeOperationClass {
	_SLSBridgedSpaceSetShapeOperationClassOnce.Do(func() {
		_SLSBridgedSpaceSetShapeOperationClass = SLSBridgedSpaceSetShapeOperationClass{class: objc.GetClass("SLSBridgedSpaceSetShapeOperation")}
	})
	return _SLSBridgedSpaceSetShapeOperationClass
}

// GetSLSBridgedSpaceSetShapeOperationClass returns the class object for SLSBridgedSpaceSetShapeOperation.
func GetSLSBridgedSpaceSetShapeOperationClass() SLSBridgedSpaceSetShapeOperationClass {
	return getSLSBridgedSpaceSetShapeOperationClass()
}

type SLSBridgedSpaceSetShapeOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedSpaceSetShapeOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedSpaceSetShapeOperationClass) Alloc() SLSBridgedSpaceSetShapeOperation {
	rv := objc.Send[SLSBridgedSpaceSetShapeOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedSpaceSetShapeOperation.CopyShape]
//   - [SLSBridgedSpaceSetShapeOperation.SpaceID]
//   - [SLSBridgedSpaceSetShapeOperation.InitWithSpaceIDShape]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetShapeOperation
type SLSBridgedSpaceSetShapeOperation struct {
	SLSAsynchronousBridgedWindowManagementOperation
}

// SLSBridgedSpaceSetShapeOperationFromID constructs a [SLSBridgedSpaceSetShapeOperation] from an objc.ID.
func SLSBridgedSpaceSetShapeOperationFromID(id objc.ID) SLSBridgedSpaceSetShapeOperation {
	return SLSBridgedSpaceSetShapeOperation{SLSAsynchronousBridgedWindowManagementOperation: SLSAsynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedSpaceSetShapeOperation implements ISLSBridgedSpaceSetShapeOperation.
var _ ISLSBridgedSpaceSetShapeOperation = SLSBridgedSpaceSetShapeOperation{}

// An interface definition for the [SLSBridgedSpaceSetShapeOperation] class.
//
// # Methods
//
//   - [ISLSBridgedSpaceSetShapeOperation.CopyShape]
//   - [ISLSBridgedSpaceSetShapeOperation.SpaceID]
//   - [ISLSBridgedSpaceSetShapeOperation.InitWithSpaceIDShape]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetShapeOperation
type ISLSBridgedSpaceSetShapeOperation interface {
	ISLSAsynchronousBridgedWindowManagementOperation

	// Topic: Methods

	CopyShape() uintptr
	SpaceID() uint64
	InitWithSpaceIDShape(id uint64, shape uintptr) SLSBridgedSpaceSetShapeOperation
}

// Init initializes the instance.
func (s SLSBridgedSpaceSetShapeOperation) Init() SLSBridgedSpaceSetShapeOperation {
	rv := objc.Send[SLSBridgedSpaceSetShapeOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedSpaceSetShapeOperation) Autorelease() SLSBridgedSpaceSetShapeOperation {
	rv := objc.Send[SLSBridgedSpaceSetShapeOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedSpaceSetShapeOperation creates a new SLSBridgedSpaceSetShapeOperation instance.
func NewSLSBridgedSpaceSetShapeOperation() SLSBridgedSpaceSetShapeOperation {
	class := getSLSBridgedSpaceSetShapeOperationClass()
	rv := objc.Send[SLSBridgedSpaceSetShapeOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetShapeOperation/initWithCoder:
func NewSLSBridgedSpaceSetShapeOperationWithCoder(coder objectivec.IObject) SLSBridgedSpaceSetShapeOperation {
	instance := getSLSBridgedSpaceSetShapeOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedSpaceSetShapeOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetShapeOperation/initWithSpaceID:shape:
func NewSLSBridgedSpaceSetShapeOperationWithSpaceIDShape(id uint64, shape uintptr) SLSBridgedSpaceSetShapeOperation {
	instance := getSLSBridgedSpaceSetShapeOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpaceID:shape:"), id, shape)
	return SLSBridgedSpaceSetShapeOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetShapeOperation/copyShape
func (s SLSBridgedSpaceSetShapeOperation) CopyShape() uintptr {
	rv := objc.Send[uintptr](s.ID, objc.Sel("copyShape"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetShapeOperation/initWithSpaceID:shape:
func (s SLSBridgedSpaceSetShapeOperation) InitWithSpaceIDShape(id uint64, shape uintptr) SLSBridgedSpaceSetShapeOperation {
	rv := objc.Send[SLSBridgedSpaceSetShapeOperation](s.ID, objc.Sel("initWithSpaceID:shape:"), id, shape)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetShapeOperation/spaceID
func (s SLSBridgedSpaceSetShapeOperation) SpaceID() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}
