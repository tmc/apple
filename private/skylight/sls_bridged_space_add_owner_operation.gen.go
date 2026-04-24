// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedSpaceAddOwnerOperation] class.
var (
	_SLSBridgedSpaceAddOwnerOperationClass     SLSBridgedSpaceAddOwnerOperationClass
	_SLSBridgedSpaceAddOwnerOperationClassOnce sync.Once
)

func getSLSBridgedSpaceAddOwnerOperationClass() SLSBridgedSpaceAddOwnerOperationClass {
	_SLSBridgedSpaceAddOwnerOperationClassOnce.Do(func() {
		_SLSBridgedSpaceAddOwnerOperationClass = SLSBridgedSpaceAddOwnerOperationClass{class: objc.GetClass("SLSBridgedSpaceAddOwnerOperation")}
	})
	return _SLSBridgedSpaceAddOwnerOperationClass
}

// GetSLSBridgedSpaceAddOwnerOperationClass returns the class object for SLSBridgedSpaceAddOwnerOperation.
func GetSLSBridgedSpaceAddOwnerOperationClass() SLSBridgedSpaceAddOwnerOperationClass {
	return getSLSBridgedSpaceAddOwnerOperationClass()
}

type SLSBridgedSpaceAddOwnerOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedSpaceAddOwnerOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedSpaceAddOwnerOperationClass) Alloc() SLSBridgedSpaceAddOwnerOperation {
	rv := objc.Send[SLSBridgedSpaceAddOwnerOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedSpaceAddOwnerOperation.Owner]
//   - [SLSBridgedSpaceAddOwnerOperation.SpaceID]
//   - [SLSBridgedSpaceAddOwnerOperation.InitWithSpaceIDOwner]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceAddOwnerOperation
type SLSBridgedSpaceAddOwnerOperation struct {
	SLSAsynchronousBridgedWindowManagementOperation
}

// SLSBridgedSpaceAddOwnerOperationFromID constructs a [SLSBridgedSpaceAddOwnerOperation] from an objc.ID.
func SLSBridgedSpaceAddOwnerOperationFromID(id objc.ID) SLSBridgedSpaceAddOwnerOperation {
	return SLSBridgedSpaceAddOwnerOperation{SLSAsynchronousBridgedWindowManagementOperation: SLSAsynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedSpaceAddOwnerOperation implements ISLSBridgedSpaceAddOwnerOperation.
var _ ISLSBridgedSpaceAddOwnerOperation = SLSBridgedSpaceAddOwnerOperation{}

// An interface definition for the [SLSBridgedSpaceAddOwnerOperation] class.
//
// # Methods
//
//   - [ISLSBridgedSpaceAddOwnerOperation.Owner]
//   - [ISLSBridgedSpaceAddOwnerOperation.SpaceID]
//   - [ISLSBridgedSpaceAddOwnerOperation.InitWithSpaceIDOwner]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceAddOwnerOperation
type ISLSBridgedSpaceAddOwnerOperation interface {
	ISLSAsynchronousBridgedWindowManagementOperation

	// Topic: Methods

	Owner() int
	SpaceID() uint64
	InitWithSpaceIDOwner(id uint64, owner int) SLSBridgedSpaceAddOwnerOperation
}

// Init initializes the instance.
func (s SLSBridgedSpaceAddOwnerOperation) Init() SLSBridgedSpaceAddOwnerOperation {
	rv := objc.Send[SLSBridgedSpaceAddOwnerOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedSpaceAddOwnerOperation) Autorelease() SLSBridgedSpaceAddOwnerOperation {
	rv := objc.Send[SLSBridgedSpaceAddOwnerOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedSpaceAddOwnerOperation creates a new SLSBridgedSpaceAddOwnerOperation instance.
func NewSLSBridgedSpaceAddOwnerOperation() SLSBridgedSpaceAddOwnerOperation {
	class := getSLSBridgedSpaceAddOwnerOperationClass()
	rv := objc.Send[SLSBridgedSpaceAddOwnerOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceAddOwnerOperation/initWithCoder:
func NewSLSBridgedSpaceAddOwnerOperationWithCoder(coder objectivec.IObject) SLSBridgedSpaceAddOwnerOperation {
	instance := getSLSBridgedSpaceAddOwnerOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedSpaceAddOwnerOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceAddOwnerOperation/initWithSpaceID:owner:
func NewSLSBridgedSpaceAddOwnerOperationWithSpaceIDOwner(id uint64, owner int) SLSBridgedSpaceAddOwnerOperation {
	instance := getSLSBridgedSpaceAddOwnerOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpaceID:owner:"), id, owner)
	return SLSBridgedSpaceAddOwnerOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceAddOwnerOperation/initWithSpaceID:owner:
func (s SLSBridgedSpaceAddOwnerOperation) InitWithSpaceIDOwner(id uint64, owner int) SLSBridgedSpaceAddOwnerOperation {
	rv := objc.Send[SLSBridgedSpaceAddOwnerOperation](s.ID, objc.Sel("initWithSpaceID:owner:"), id, owner)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceAddOwnerOperation/owner
func (s SLSBridgedSpaceAddOwnerOperation) Owner() int {
	rv := objc.Send[int](s.ID, objc.Sel("owner"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceAddOwnerOperation/spaceID
func (s SLSBridgedSpaceAddOwnerOperation) SpaceID() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}
