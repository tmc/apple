// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedCopyManagedDisplayForSpaceOperation] class.
var (
	_SLSBridgedCopyManagedDisplayForSpaceOperationClass     SLSBridgedCopyManagedDisplayForSpaceOperationClass
	_SLSBridgedCopyManagedDisplayForSpaceOperationClassOnce sync.Once
)

func getSLSBridgedCopyManagedDisplayForSpaceOperationClass() SLSBridgedCopyManagedDisplayForSpaceOperationClass {
	_SLSBridgedCopyManagedDisplayForSpaceOperationClassOnce.Do(func() {
		_SLSBridgedCopyManagedDisplayForSpaceOperationClass = SLSBridgedCopyManagedDisplayForSpaceOperationClass{class: objc.GetClass("SLSBridgedCopyManagedDisplayForSpaceOperation")}
	})
	return _SLSBridgedCopyManagedDisplayForSpaceOperationClass
}

// GetSLSBridgedCopyManagedDisplayForSpaceOperationClass returns the class object for SLSBridgedCopyManagedDisplayForSpaceOperation.
func GetSLSBridgedCopyManagedDisplayForSpaceOperationClass() SLSBridgedCopyManagedDisplayForSpaceOperationClass {
	return getSLSBridgedCopyManagedDisplayForSpaceOperationClass()
}

type SLSBridgedCopyManagedDisplayForSpaceOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedCopyManagedDisplayForSpaceOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedCopyManagedDisplayForSpaceOperationClass) Alloc() SLSBridgedCopyManagedDisplayForSpaceOperation {
	rv := objc.Send[SLSBridgedCopyManagedDisplayForSpaceOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedCopyManagedDisplayForSpaceOperation.MakeResultWithString]
//   - [SLSBridgedCopyManagedDisplayForSpaceOperation.SpaceID]
//   - [SLSBridgedCopyManagedDisplayForSpaceOperation.InitWithSpaceID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyManagedDisplayForSpaceOperation
type SLSBridgedCopyManagedDisplayForSpaceOperation struct {
	SLSSynchronousBridgedWindowManagementOperation
}

// SLSBridgedCopyManagedDisplayForSpaceOperationFromID constructs a [SLSBridgedCopyManagedDisplayForSpaceOperation] from an objc.ID.
func SLSBridgedCopyManagedDisplayForSpaceOperationFromID(id objc.ID) SLSBridgedCopyManagedDisplayForSpaceOperation {
	return SLSBridgedCopyManagedDisplayForSpaceOperation{SLSSynchronousBridgedWindowManagementOperation: SLSSynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedCopyManagedDisplayForSpaceOperation implements ISLSBridgedCopyManagedDisplayForSpaceOperation.
var _ ISLSBridgedCopyManagedDisplayForSpaceOperation = SLSBridgedCopyManagedDisplayForSpaceOperation{}

// An interface definition for the [SLSBridgedCopyManagedDisplayForSpaceOperation] class.
//
// # Methods
//
//   - [ISLSBridgedCopyManagedDisplayForSpaceOperation.MakeResultWithString]
//   - [ISLSBridgedCopyManagedDisplayForSpaceOperation.SpaceID]
//   - [ISLSBridgedCopyManagedDisplayForSpaceOperation.InitWithSpaceID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyManagedDisplayForSpaceOperation
type ISLSBridgedCopyManagedDisplayForSpaceOperation interface {
	ISLSSynchronousBridgedWindowManagementOperation

	// Topic: Methods

	MakeResultWithString(string_ objectivec.IObject) objectivec.IObject
	SpaceID() uint64
	InitWithSpaceID(id uint64) SLSBridgedCopyManagedDisplayForSpaceOperation
}

// Init initializes the instance.
func (s SLSBridgedCopyManagedDisplayForSpaceOperation) Init() SLSBridgedCopyManagedDisplayForSpaceOperation {
	rv := objc.Send[SLSBridgedCopyManagedDisplayForSpaceOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedCopyManagedDisplayForSpaceOperation) Autorelease() SLSBridgedCopyManagedDisplayForSpaceOperation {
	rv := objc.Send[SLSBridgedCopyManagedDisplayForSpaceOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedCopyManagedDisplayForSpaceOperation creates a new SLSBridgedCopyManagedDisplayForSpaceOperation instance.
func NewSLSBridgedCopyManagedDisplayForSpaceOperation() SLSBridgedCopyManagedDisplayForSpaceOperation {
	class := getSLSBridgedCopyManagedDisplayForSpaceOperationClass()
	rv := objc.Send[SLSBridgedCopyManagedDisplayForSpaceOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyManagedDisplayForSpaceOperation/initWithCoder:
func NewSLSBridgedCopyManagedDisplayForSpaceOperationWithCoder(coder objectivec.IObject) SLSBridgedCopyManagedDisplayForSpaceOperation {
	instance := getSLSBridgedCopyManagedDisplayForSpaceOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedCopyManagedDisplayForSpaceOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyManagedDisplayForSpaceOperation/initWithSpaceID:
func NewSLSBridgedCopyManagedDisplayForSpaceOperationWithSpaceID(id uint64) SLSBridgedCopyManagedDisplayForSpaceOperation {
	instance := getSLSBridgedCopyManagedDisplayForSpaceOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpaceID:"), id)
	return SLSBridgedCopyManagedDisplayForSpaceOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyManagedDisplayForSpaceOperation/makeResultWithString:
func (s SLSBridgedCopyManagedDisplayForSpaceOperation) MakeResultWithString(string_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("makeResultWithString:"), string_)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyManagedDisplayForSpaceOperation/initWithSpaceID:
func (s SLSBridgedCopyManagedDisplayForSpaceOperation) InitWithSpaceID(id uint64) SLSBridgedCopyManagedDisplayForSpaceOperation {
	rv := objc.Send[SLSBridgedCopyManagedDisplayForSpaceOperation](s.ID, objc.Sel("initWithSpaceID:"), id)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopyManagedDisplayForSpaceOperation/spaceID
func (s SLSBridgedCopyManagedDisplayForSpaceOperation) SpaceID() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}
