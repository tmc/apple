// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedSpaceSetValuesOperation] class.
var (
	_SLSBridgedSpaceSetValuesOperationClass     SLSBridgedSpaceSetValuesOperationClass
	_SLSBridgedSpaceSetValuesOperationClassOnce sync.Once
)

func getSLSBridgedSpaceSetValuesOperationClass() SLSBridgedSpaceSetValuesOperationClass {
	_SLSBridgedSpaceSetValuesOperationClassOnce.Do(func() {
		_SLSBridgedSpaceSetValuesOperationClass = SLSBridgedSpaceSetValuesOperationClass{class: objc.GetClass("SLSBridgedSpaceSetValuesOperation")}
	})
	return _SLSBridgedSpaceSetValuesOperationClass
}

// GetSLSBridgedSpaceSetValuesOperationClass returns the class object for SLSBridgedSpaceSetValuesOperation.
func GetSLSBridgedSpaceSetValuesOperationClass() SLSBridgedSpaceSetValuesOperationClass {
	return getSLSBridgedSpaceSetValuesOperationClass()
}

type SLSBridgedSpaceSetValuesOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedSpaceSetValuesOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedSpaceSetValuesOperationClass) Alloc() SLSBridgedSpaceSetValuesOperation {
	rv := objc.Send[SLSBridgedSpaceSetValuesOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedSpaceSetValuesOperation.SpaceID]
//   - [SLSBridgedSpaceSetValuesOperation.Values]
//   - [SLSBridgedSpaceSetValuesOperation.InitWithSpaceIDValues]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetValuesOperation
type SLSBridgedSpaceSetValuesOperation struct {
	SLSAsynchronousBridgedWindowManagementOperation
}

// SLSBridgedSpaceSetValuesOperationFromID constructs a [SLSBridgedSpaceSetValuesOperation] from an objc.ID.
func SLSBridgedSpaceSetValuesOperationFromID(id objc.ID) SLSBridgedSpaceSetValuesOperation {
	return SLSBridgedSpaceSetValuesOperation{SLSAsynchronousBridgedWindowManagementOperation: SLSAsynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedSpaceSetValuesOperation implements ISLSBridgedSpaceSetValuesOperation.
var _ ISLSBridgedSpaceSetValuesOperation = SLSBridgedSpaceSetValuesOperation{}

// An interface definition for the [SLSBridgedSpaceSetValuesOperation] class.
//
// # Methods
//
//   - [ISLSBridgedSpaceSetValuesOperation.SpaceID]
//   - [ISLSBridgedSpaceSetValuesOperation.Values]
//   - [ISLSBridgedSpaceSetValuesOperation.InitWithSpaceIDValues]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetValuesOperation
type ISLSBridgedSpaceSetValuesOperation interface {
	ISLSAsynchronousBridgedWindowManagementOperation

	// Topic: Methods

	SpaceID() uint64
	Values() foundation.INSDictionary
	InitWithSpaceIDValues(id uint64, values objectivec.IObject) SLSBridgedSpaceSetValuesOperation
}

// Init initializes the instance.
func (s SLSBridgedSpaceSetValuesOperation) Init() SLSBridgedSpaceSetValuesOperation {
	rv := objc.Send[SLSBridgedSpaceSetValuesOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedSpaceSetValuesOperation) Autorelease() SLSBridgedSpaceSetValuesOperation {
	rv := objc.Send[SLSBridgedSpaceSetValuesOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedSpaceSetValuesOperation creates a new SLSBridgedSpaceSetValuesOperation instance.
func NewSLSBridgedSpaceSetValuesOperation() SLSBridgedSpaceSetValuesOperation {
	class := getSLSBridgedSpaceSetValuesOperationClass()
	rv := objc.Send[SLSBridgedSpaceSetValuesOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetValuesOperation/initWithCoder:
func NewSLSBridgedSpaceSetValuesOperationWithCoder(coder objectivec.IObject) SLSBridgedSpaceSetValuesOperation {
	instance := getSLSBridgedSpaceSetValuesOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedSpaceSetValuesOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetValuesOperation/initWithSpaceID:values:
func NewSLSBridgedSpaceSetValuesOperationWithSpaceIDValues(id uint64, values objectivec.IObject) SLSBridgedSpaceSetValuesOperation {
	instance := getSLSBridgedSpaceSetValuesOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpaceID:values:"), id, values)
	return SLSBridgedSpaceSetValuesOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetValuesOperation/initWithSpaceID:values:
func (s SLSBridgedSpaceSetValuesOperation) InitWithSpaceIDValues(id uint64, values objectivec.IObject) SLSBridgedSpaceSetValuesOperation {
	rv := objc.Send[SLSBridgedSpaceSetValuesOperation](s.ID, objc.Sel("initWithSpaceID:values:"), id, values)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetValuesOperation/spaceID
func (s SLSBridgedSpaceSetValuesOperation) SpaceID() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceSetValuesOperation/values
func (s SLSBridgedSpaceSetValuesOperation) Values() foundation.INSDictionary {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("values"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
