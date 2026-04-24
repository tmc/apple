// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedSpaceRemoveValuesForKeysOperation] class.
var (
	_SLSBridgedSpaceRemoveValuesForKeysOperationClass     SLSBridgedSpaceRemoveValuesForKeysOperationClass
	_SLSBridgedSpaceRemoveValuesForKeysOperationClassOnce sync.Once
)

func getSLSBridgedSpaceRemoveValuesForKeysOperationClass() SLSBridgedSpaceRemoveValuesForKeysOperationClass {
	_SLSBridgedSpaceRemoveValuesForKeysOperationClassOnce.Do(func() {
		_SLSBridgedSpaceRemoveValuesForKeysOperationClass = SLSBridgedSpaceRemoveValuesForKeysOperationClass{class: objc.GetClass("SLSBridgedSpaceRemoveValuesForKeysOperation")}
	})
	return _SLSBridgedSpaceRemoveValuesForKeysOperationClass
}

// GetSLSBridgedSpaceRemoveValuesForKeysOperationClass returns the class object for SLSBridgedSpaceRemoveValuesForKeysOperation.
func GetSLSBridgedSpaceRemoveValuesForKeysOperationClass() SLSBridgedSpaceRemoveValuesForKeysOperationClass {
	return getSLSBridgedSpaceRemoveValuesForKeysOperationClass()
}

type SLSBridgedSpaceRemoveValuesForKeysOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedSpaceRemoveValuesForKeysOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedSpaceRemoveValuesForKeysOperationClass) Alloc() SLSBridgedSpaceRemoveValuesForKeysOperation {
	rv := objc.Send[SLSBridgedSpaceRemoveValuesForKeysOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedSpaceRemoveValuesForKeysOperation.Keys]
//   - [SLSBridgedSpaceRemoveValuesForKeysOperation.SpaceID]
//   - [SLSBridgedSpaceRemoveValuesForKeysOperation.InitWithSpaceIDKeys]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceRemoveValuesForKeysOperation
type SLSBridgedSpaceRemoveValuesForKeysOperation struct {
	SLSAsynchronousBridgedWindowManagementOperation
}

// SLSBridgedSpaceRemoveValuesForKeysOperationFromID constructs a [SLSBridgedSpaceRemoveValuesForKeysOperation] from an objc.ID.
func SLSBridgedSpaceRemoveValuesForKeysOperationFromID(id objc.ID) SLSBridgedSpaceRemoveValuesForKeysOperation {
	return SLSBridgedSpaceRemoveValuesForKeysOperation{SLSAsynchronousBridgedWindowManagementOperation: SLSAsynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedSpaceRemoveValuesForKeysOperation implements ISLSBridgedSpaceRemoveValuesForKeysOperation.
var _ ISLSBridgedSpaceRemoveValuesForKeysOperation = SLSBridgedSpaceRemoveValuesForKeysOperation{}

// An interface definition for the [SLSBridgedSpaceRemoveValuesForKeysOperation] class.
//
// # Methods
//
//   - [ISLSBridgedSpaceRemoveValuesForKeysOperation.Keys]
//   - [ISLSBridgedSpaceRemoveValuesForKeysOperation.SpaceID]
//   - [ISLSBridgedSpaceRemoveValuesForKeysOperation.InitWithSpaceIDKeys]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceRemoveValuesForKeysOperation
type ISLSBridgedSpaceRemoveValuesForKeysOperation interface {
	ISLSAsynchronousBridgedWindowManagementOperation

	// Topic: Methods

	Keys() foundation.INSArray
	SpaceID() uint64
	InitWithSpaceIDKeys(id uint64, keys objectivec.IObject) SLSBridgedSpaceRemoveValuesForKeysOperation
}

// Init initializes the instance.
func (s SLSBridgedSpaceRemoveValuesForKeysOperation) Init() SLSBridgedSpaceRemoveValuesForKeysOperation {
	rv := objc.Send[SLSBridgedSpaceRemoveValuesForKeysOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedSpaceRemoveValuesForKeysOperation) Autorelease() SLSBridgedSpaceRemoveValuesForKeysOperation {
	rv := objc.Send[SLSBridgedSpaceRemoveValuesForKeysOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedSpaceRemoveValuesForKeysOperation creates a new SLSBridgedSpaceRemoveValuesForKeysOperation instance.
func NewSLSBridgedSpaceRemoveValuesForKeysOperation() SLSBridgedSpaceRemoveValuesForKeysOperation {
	class := getSLSBridgedSpaceRemoveValuesForKeysOperationClass()
	rv := objc.Send[SLSBridgedSpaceRemoveValuesForKeysOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceRemoveValuesForKeysOperation/initWithCoder:
func NewSLSBridgedSpaceRemoveValuesForKeysOperationWithCoder(coder objectivec.IObject) SLSBridgedSpaceRemoveValuesForKeysOperation {
	instance := getSLSBridgedSpaceRemoveValuesForKeysOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedSpaceRemoveValuesForKeysOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceRemoveValuesForKeysOperation/initWithSpaceID:keys:
func NewSLSBridgedSpaceRemoveValuesForKeysOperationWithSpaceIDKeys(id uint64, keys objectivec.IObject) SLSBridgedSpaceRemoveValuesForKeysOperation {
	instance := getSLSBridgedSpaceRemoveValuesForKeysOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpaceID:keys:"), id, keys)
	return SLSBridgedSpaceRemoveValuesForKeysOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceRemoveValuesForKeysOperation/initWithSpaceID:keys:
func (s SLSBridgedSpaceRemoveValuesForKeysOperation) InitWithSpaceIDKeys(id uint64, keys objectivec.IObject) SLSBridgedSpaceRemoveValuesForKeysOperation {
	rv := objc.Send[SLSBridgedSpaceRemoveValuesForKeysOperation](s.ID, objc.Sel("initWithSpaceID:keys:"), id, keys)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceRemoveValuesForKeysOperation/keys
func (s SLSBridgedSpaceRemoveValuesForKeysOperation) Keys() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("keys"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceRemoveValuesForKeysOperation/spaceID
func (s SLSBridgedSpaceRemoveValuesForKeysOperation) SpaceID() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}
