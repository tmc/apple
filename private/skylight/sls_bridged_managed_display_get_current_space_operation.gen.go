// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedManagedDisplayGetCurrentSpaceOperation] class.
var (
	_SLSBridgedManagedDisplayGetCurrentSpaceOperationClass     SLSBridgedManagedDisplayGetCurrentSpaceOperationClass
	_SLSBridgedManagedDisplayGetCurrentSpaceOperationClassOnce sync.Once
)

func getSLSBridgedManagedDisplayGetCurrentSpaceOperationClass() SLSBridgedManagedDisplayGetCurrentSpaceOperationClass {
	_SLSBridgedManagedDisplayGetCurrentSpaceOperationClassOnce.Do(func() {
		_SLSBridgedManagedDisplayGetCurrentSpaceOperationClass = SLSBridgedManagedDisplayGetCurrentSpaceOperationClass{class: objc.GetClass("SLSBridgedManagedDisplayGetCurrentSpaceOperation")}
	})
	return _SLSBridgedManagedDisplayGetCurrentSpaceOperationClass
}

// GetSLSBridgedManagedDisplayGetCurrentSpaceOperationClass returns the class object for SLSBridgedManagedDisplayGetCurrentSpaceOperation.
func GetSLSBridgedManagedDisplayGetCurrentSpaceOperationClass() SLSBridgedManagedDisplayGetCurrentSpaceOperationClass {
	return getSLSBridgedManagedDisplayGetCurrentSpaceOperationClass()
}

type SLSBridgedManagedDisplayGetCurrentSpaceOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedManagedDisplayGetCurrentSpaceOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedManagedDisplayGetCurrentSpaceOperationClass) Alloc() SLSBridgedManagedDisplayGetCurrentSpaceOperation {
	rv := objc.Send[SLSBridgedManagedDisplayGetCurrentSpaceOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedManagedDisplayGetCurrentSpaceOperation.DisplayIdentifier]
//   - [SLSBridgedManagedDisplayGetCurrentSpaceOperation.MakeResultWithSpaceID]
//   - [SLSBridgedManagedDisplayGetCurrentSpaceOperation.InitWithDisplayIdentifier]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedManagedDisplayGetCurrentSpaceOperation
type SLSBridgedManagedDisplayGetCurrentSpaceOperation struct {
	SLSSynchronousBridgedWindowManagementOperation
}

// SLSBridgedManagedDisplayGetCurrentSpaceOperationFromID constructs a [SLSBridgedManagedDisplayGetCurrentSpaceOperation] from an objc.ID.
func SLSBridgedManagedDisplayGetCurrentSpaceOperationFromID(id objc.ID) SLSBridgedManagedDisplayGetCurrentSpaceOperation {
	return SLSBridgedManagedDisplayGetCurrentSpaceOperation{SLSSynchronousBridgedWindowManagementOperation: SLSSynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedManagedDisplayGetCurrentSpaceOperation implements ISLSBridgedManagedDisplayGetCurrentSpaceOperation.
var _ ISLSBridgedManagedDisplayGetCurrentSpaceOperation = SLSBridgedManagedDisplayGetCurrentSpaceOperation{}

// An interface definition for the [SLSBridgedManagedDisplayGetCurrentSpaceOperation] class.
//
// # Methods
//
//   - [ISLSBridgedManagedDisplayGetCurrentSpaceOperation.DisplayIdentifier]
//   - [ISLSBridgedManagedDisplayGetCurrentSpaceOperation.MakeResultWithSpaceID]
//   - [ISLSBridgedManagedDisplayGetCurrentSpaceOperation.InitWithDisplayIdentifier]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedManagedDisplayGetCurrentSpaceOperation
type ISLSBridgedManagedDisplayGetCurrentSpaceOperation interface {
	ISLSSynchronousBridgedWindowManagementOperation

	// Topic: Methods

	DisplayIdentifier() string
	MakeResultWithSpaceID(id uint64) objectivec.IObject
	InitWithDisplayIdentifier(identifier objectivec.IObject) SLSBridgedManagedDisplayGetCurrentSpaceOperation
}

// Init initializes the instance.
func (s SLSBridgedManagedDisplayGetCurrentSpaceOperation) Init() SLSBridgedManagedDisplayGetCurrentSpaceOperation {
	rv := objc.Send[SLSBridgedManagedDisplayGetCurrentSpaceOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedManagedDisplayGetCurrentSpaceOperation) Autorelease() SLSBridgedManagedDisplayGetCurrentSpaceOperation {
	rv := objc.Send[SLSBridgedManagedDisplayGetCurrentSpaceOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedManagedDisplayGetCurrentSpaceOperation creates a new SLSBridgedManagedDisplayGetCurrentSpaceOperation instance.
func NewSLSBridgedManagedDisplayGetCurrentSpaceOperation() SLSBridgedManagedDisplayGetCurrentSpaceOperation {
	class := getSLSBridgedManagedDisplayGetCurrentSpaceOperationClass()
	rv := objc.Send[SLSBridgedManagedDisplayGetCurrentSpaceOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedManagedDisplayGetCurrentSpaceOperation/initWithCoder:
func NewSLSBridgedManagedDisplayGetCurrentSpaceOperationWithCoder(coder objectivec.IObject) SLSBridgedManagedDisplayGetCurrentSpaceOperation {
	instance := getSLSBridgedManagedDisplayGetCurrentSpaceOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedManagedDisplayGetCurrentSpaceOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedManagedDisplayGetCurrentSpaceOperation/initWithDisplayIdentifier:
func NewSLSBridgedManagedDisplayGetCurrentSpaceOperationWithDisplayIdentifier(identifier objectivec.IObject) SLSBridgedManagedDisplayGetCurrentSpaceOperation {
	instance := getSLSBridgedManagedDisplayGetCurrentSpaceOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDisplayIdentifier:"), identifier)
	return SLSBridgedManagedDisplayGetCurrentSpaceOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedManagedDisplayGetCurrentSpaceOperation/makeResultWithSpaceID:
func (s SLSBridgedManagedDisplayGetCurrentSpaceOperation) MakeResultWithSpaceID(id uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("makeResultWithSpaceID:"), id)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedManagedDisplayGetCurrentSpaceOperation/initWithDisplayIdentifier:
func (s SLSBridgedManagedDisplayGetCurrentSpaceOperation) InitWithDisplayIdentifier(identifier objectivec.IObject) SLSBridgedManagedDisplayGetCurrentSpaceOperation {
	rv := objc.Send[SLSBridgedManagedDisplayGetCurrentSpaceOperation](s.ID, objc.Sel("initWithDisplayIdentifier:"), identifier)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedManagedDisplayGetCurrentSpaceOperation/displayIdentifier
func (s SLSBridgedManagedDisplayGetCurrentSpaceOperation) DisplayIdentifier() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("displayIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
