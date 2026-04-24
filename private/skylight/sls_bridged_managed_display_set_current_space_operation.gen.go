// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedManagedDisplaySetCurrentSpaceOperation] class.
var (
	_SLSBridgedManagedDisplaySetCurrentSpaceOperationClass     SLSBridgedManagedDisplaySetCurrentSpaceOperationClass
	_SLSBridgedManagedDisplaySetCurrentSpaceOperationClassOnce sync.Once
)

func getSLSBridgedManagedDisplaySetCurrentSpaceOperationClass() SLSBridgedManagedDisplaySetCurrentSpaceOperationClass {
	_SLSBridgedManagedDisplaySetCurrentSpaceOperationClassOnce.Do(func() {
		_SLSBridgedManagedDisplaySetCurrentSpaceOperationClass = SLSBridgedManagedDisplaySetCurrentSpaceOperationClass{class: objc.GetClass("SLSBridgedManagedDisplaySetCurrentSpaceOperation")}
	})
	return _SLSBridgedManagedDisplaySetCurrentSpaceOperationClass
}

// GetSLSBridgedManagedDisplaySetCurrentSpaceOperationClass returns the class object for SLSBridgedManagedDisplaySetCurrentSpaceOperation.
func GetSLSBridgedManagedDisplaySetCurrentSpaceOperationClass() SLSBridgedManagedDisplaySetCurrentSpaceOperationClass {
	return getSLSBridgedManagedDisplaySetCurrentSpaceOperationClass()
}

type SLSBridgedManagedDisplaySetCurrentSpaceOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedManagedDisplaySetCurrentSpaceOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedManagedDisplaySetCurrentSpaceOperationClass) Alloc() SLSBridgedManagedDisplaySetCurrentSpaceOperation {
	rv := objc.Send[SLSBridgedManagedDisplaySetCurrentSpaceOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedManagedDisplaySetCurrentSpaceOperation.DisplayIdentifier]
//   - [SLSBridgedManagedDisplaySetCurrentSpaceOperation.SpaceID]
//   - [SLSBridgedManagedDisplaySetCurrentSpaceOperation.InitWithDisplayIdentifierSpaceID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedManagedDisplaySetCurrentSpaceOperation
type SLSBridgedManagedDisplaySetCurrentSpaceOperation struct {
	SLSAsynchronousBridgedWindowManagementOperation
}

// SLSBridgedManagedDisplaySetCurrentSpaceOperationFromID constructs a [SLSBridgedManagedDisplaySetCurrentSpaceOperation] from an objc.ID.
func SLSBridgedManagedDisplaySetCurrentSpaceOperationFromID(id objc.ID) SLSBridgedManagedDisplaySetCurrentSpaceOperation {
	return SLSBridgedManagedDisplaySetCurrentSpaceOperation{SLSAsynchronousBridgedWindowManagementOperation: SLSAsynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedManagedDisplaySetCurrentSpaceOperation implements ISLSBridgedManagedDisplaySetCurrentSpaceOperation.
var _ ISLSBridgedManagedDisplaySetCurrentSpaceOperation = SLSBridgedManagedDisplaySetCurrentSpaceOperation{}

// An interface definition for the [SLSBridgedManagedDisplaySetCurrentSpaceOperation] class.
//
// # Methods
//
//   - [ISLSBridgedManagedDisplaySetCurrentSpaceOperation.DisplayIdentifier]
//   - [ISLSBridgedManagedDisplaySetCurrentSpaceOperation.SpaceID]
//   - [ISLSBridgedManagedDisplaySetCurrentSpaceOperation.InitWithDisplayIdentifierSpaceID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedManagedDisplaySetCurrentSpaceOperation
type ISLSBridgedManagedDisplaySetCurrentSpaceOperation interface {
	ISLSAsynchronousBridgedWindowManagementOperation

	// Topic: Methods

	DisplayIdentifier() string
	SpaceID() uint64
	InitWithDisplayIdentifierSpaceID(identifier objectivec.IObject, id uint64) SLSBridgedManagedDisplaySetCurrentSpaceOperation
}

// Init initializes the instance.
func (s SLSBridgedManagedDisplaySetCurrentSpaceOperation) Init() SLSBridgedManagedDisplaySetCurrentSpaceOperation {
	rv := objc.Send[SLSBridgedManagedDisplaySetCurrentSpaceOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedManagedDisplaySetCurrentSpaceOperation) Autorelease() SLSBridgedManagedDisplaySetCurrentSpaceOperation {
	rv := objc.Send[SLSBridgedManagedDisplaySetCurrentSpaceOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedManagedDisplaySetCurrentSpaceOperation creates a new SLSBridgedManagedDisplaySetCurrentSpaceOperation instance.
func NewSLSBridgedManagedDisplaySetCurrentSpaceOperation() SLSBridgedManagedDisplaySetCurrentSpaceOperation {
	class := getSLSBridgedManagedDisplaySetCurrentSpaceOperationClass()
	rv := objc.Send[SLSBridgedManagedDisplaySetCurrentSpaceOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedManagedDisplaySetCurrentSpaceOperation/initWithCoder:
func NewSLSBridgedManagedDisplaySetCurrentSpaceOperationWithCoder(coder objectivec.IObject) SLSBridgedManagedDisplaySetCurrentSpaceOperation {
	instance := getSLSBridgedManagedDisplaySetCurrentSpaceOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedManagedDisplaySetCurrentSpaceOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedManagedDisplaySetCurrentSpaceOperation/initWithDisplayIdentifier:spaceID:
func NewSLSBridgedManagedDisplaySetCurrentSpaceOperationWithDisplayIdentifierSpaceID(identifier objectivec.IObject, id uint64) SLSBridgedManagedDisplaySetCurrentSpaceOperation {
	instance := getSLSBridgedManagedDisplaySetCurrentSpaceOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDisplayIdentifier:spaceID:"), identifier, id)
	return SLSBridgedManagedDisplaySetCurrentSpaceOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedManagedDisplaySetCurrentSpaceOperation/initWithDisplayIdentifier:spaceID:
func (s SLSBridgedManagedDisplaySetCurrentSpaceOperation) InitWithDisplayIdentifierSpaceID(identifier objectivec.IObject, id uint64) SLSBridgedManagedDisplaySetCurrentSpaceOperation {
	rv := objc.Send[SLSBridgedManagedDisplaySetCurrentSpaceOperation](s.ID, objc.Sel("initWithDisplayIdentifier:spaceID:"), identifier, id)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedManagedDisplaySetCurrentSpaceOperation/displayIdentifier
func (s SLSBridgedManagedDisplaySetCurrentSpaceOperation) DisplayIdentifier() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("displayIdentifier"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedManagedDisplaySetCurrentSpaceOperation/spaceID
func (s SLSBridgedManagedDisplaySetCurrentSpaceOperation) SpaceID() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}
