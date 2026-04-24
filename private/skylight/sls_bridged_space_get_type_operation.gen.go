// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedSpaceGetTypeOperation] class.
var (
	_SLSBridgedSpaceGetTypeOperationClass     SLSBridgedSpaceGetTypeOperationClass
	_SLSBridgedSpaceGetTypeOperationClassOnce sync.Once
)

func getSLSBridgedSpaceGetTypeOperationClass() SLSBridgedSpaceGetTypeOperationClass {
	_SLSBridgedSpaceGetTypeOperationClassOnce.Do(func() {
		_SLSBridgedSpaceGetTypeOperationClass = SLSBridgedSpaceGetTypeOperationClass{class: objc.GetClass("SLSBridgedSpaceGetTypeOperation")}
	})
	return _SLSBridgedSpaceGetTypeOperationClass
}

// GetSLSBridgedSpaceGetTypeOperationClass returns the class object for SLSBridgedSpaceGetTypeOperation.
func GetSLSBridgedSpaceGetTypeOperationClass() SLSBridgedSpaceGetTypeOperationClass {
	return getSLSBridgedSpaceGetTypeOperationClass()
}

type SLSBridgedSpaceGetTypeOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedSpaceGetTypeOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedSpaceGetTypeOperationClass) Alloc() SLSBridgedSpaceGetTypeOperation {
	rv := objc.Send[SLSBridgedSpaceGetTypeOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedSpaceGetTypeOperation.MakeResultWithWorkspaceType]
//   - [SLSBridgedSpaceGetTypeOperation.SpaceID]
//   - [SLSBridgedSpaceGetTypeOperation.InitWithSpaceID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceGetTypeOperation
type SLSBridgedSpaceGetTypeOperation struct {
	SLSSynchronousBridgedWindowManagementOperation
}

// SLSBridgedSpaceGetTypeOperationFromID constructs a [SLSBridgedSpaceGetTypeOperation] from an objc.ID.
func SLSBridgedSpaceGetTypeOperationFromID(id objc.ID) SLSBridgedSpaceGetTypeOperation {
	return SLSBridgedSpaceGetTypeOperation{SLSSynchronousBridgedWindowManagementOperation: SLSSynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedSpaceGetTypeOperation implements ISLSBridgedSpaceGetTypeOperation.
var _ ISLSBridgedSpaceGetTypeOperation = SLSBridgedSpaceGetTypeOperation{}

// An interface definition for the [SLSBridgedSpaceGetTypeOperation] class.
//
// # Methods
//
//   - [ISLSBridgedSpaceGetTypeOperation.MakeResultWithWorkspaceType]
//   - [ISLSBridgedSpaceGetTypeOperation.SpaceID]
//   - [ISLSBridgedSpaceGetTypeOperation.InitWithSpaceID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceGetTypeOperation
type ISLSBridgedSpaceGetTypeOperation interface {
	ISLSSynchronousBridgedWindowManagementOperation

	// Topic: Methods

	MakeResultWithWorkspaceType(type_ int) objectivec.IObject
	SpaceID() uint64
	InitWithSpaceID(id uint64) SLSBridgedSpaceGetTypeOperation
}

// Init initializes the instance.
func (s SLSBridgedSpaceGetTypeOperation) Init() SLSBridgedSpaceGetTypeOperation {
	rv := objc.Send[SLSBridgedSpaceGetTypeOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedSpaceGetTypeOperation) Autorelease() SLSBridgedSpaceGetTypeOperation {
	rv := objc.Send[SLSBridgedSpaceGetTypeOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedSpaceGetTypeOperation creates a new SLSBridgedSpaceGetTypeOperation instance.
func NewSLSBridgedSpaceGetTypeOperation() SLSBridgedSpaceGetTypeOperation {
	class := getSLSBridgedSpaceGetTypeOperationClass()
	rv := objc.Send[SLSBridgedSpaceGetTypeOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceGetTypeOperation/initWithCoder:
func NewSLSBridgedSpaceGetTypeOperationWithCoder(coder objectivec.IObject) SLSBridgedSpaceGetTypeOperation {
	instance := getSLSBridgedSpaceGetTypeOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedSpaceGetTypeOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceGetTypeOperation/initWithSpaceID:
func NewSLSBridgedSpaceGetTypeOperationWithSpaceID(id uint64) SLSBridgedSpaceGetTypeOperation {
	instance := getSLSBridgedSpaceGetTypeOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpaceID:"), id)
	return SLSBridgedSpaceGetTypeOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceGetTypeOperation/makeResultWithWorkspaceType:
func (s SLSBridgedSpaceGetTypeOperation) MakeResultWithWorkspaceType(type_ int) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("makeResultWithWorkspaceType:"), type_)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceGetTypeOperation/initWithSpaceID:
func (s SLSBridgedSpaceGetTypeOperation) InitWithSpaceID(id uint64) SLSBridgedSpaceGetTypeOperation {
	rv := objc.Send[SLSBridgedSpaceGetTypeOperation](s.ID, objc.Sel("initWithSpaceID:"), id)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceGetTypeOperation/spaceID
func (s SLSBridgedSpaceGetTypeOperation) SpaceID() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}
