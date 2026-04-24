// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedSpaceCreateTileOperation] class.
var (
	_SLSBridgedSpaceCreateTileOperationClass     SLSBridgedSpaceCreateTileOperationClass
	_SLSBridgedSpaceCreateTileOperationClassOnce sync.Once
)

func getSLSBridgedSpaceCreateTileOperationClass() SLSBridgedSpaceCreateTileOperationClass {
	_SLSBridgedSpaceCreateTileOperationClassOnce.Do(func() {
		_SLSBridgedSpaceCreateTileOperationClass = SLSBridgedSpaceCreateTileOperationClass{class: objc.GetClass("SLSBridgedSpaceCreateTileOperation")}
	})
	return _SLSBridgedSpaceCreateTileOperationClass
}

// GetSLSBridgedSpaceCreateTileOperationClass returns the class object for SLSBridgedSpaceCreateTileOperation.
func GetSLSBridgedSpaceCreateTileOperationClass() SLSBridgedSpaceCreateTileOperationClass {
	return getSLSBridgedSpaceCreateTileOperationClass()
}

type SLSBridgedSpaceCreateTileOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedSpaceCreateTileOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedSpaceCreateTileOperationClass) Alloc() SLSBridgedSpaceCreateTileOperation {
	rv := objc.Send[SLSBridgedSpaceCreateTileOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedSpaceCreateTileOperation.MakeResultWithSpaceID]
//   - [SLSBridgedSpaceCreateTileOperation.SpaceID]
//   - [SLSBridgedSpaceCreateTileOperation.Values]
//   - [SLSBridgedSpaceCreateTileOperation.InitWithSpaceIDValues]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCreateTileOperation
type SLSBridgedSpaceCreateTileOperation struct {
	SLSSynchronousBridgedWindowManagementOperation
}

// SLSBridgedSpaceCreateTileOperationFromID constructs a [SLSBridgedSpaceCreateTileOperation] from an objc.ID.
func SLSBridgedSpaceCreateTileOperationFromID(id objc.ID) SLSBridgedSpaceCreateTileOperation {
	return SLSBridgedSpaceCreateTileOperation{SLSSynchronousBridgedWindowManagementOperation: SLSSynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedSpaceCreateTileOperation implements ISLSBridgedSpaceCreateTileOperation.
var _ ISLSBridgedSpaceCreateTileOperation = SLSBridgedSpaceCreateTileOperation{}

// An interface definition for the [SLSBridgedSpaceCreateTileOperation] class.
//
// # Methods
//
//   - [ISLSBridgedSpaceCreateTileOperation.MakeResultWithSpaceID]
//   - [ISLSBridgedSpaceCreateTileOperation.SpaceID]
//   - [ISLSBridgedSpaceCreateTileOperation.Values]
//   - [ISLSBridgedSpaceCreateTileOperation.InitWithSpaceIDValues]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCreateTileOperation
type ISLSBridgedSpaceCreateTileOperation interface {
	ISLSSynchronousBridgedWindowManagementOperation

	// Topic: Methods

	MakeResultWithSpaceID(id uint64) objectivec.IObject
	SpaceID() uint64
	Values() foundation.INSDictionary
	InitWithSpaceIDValues(id uint64, values objectivec.IObject) SLSBridgedSpaceCreateTileOperation
}

// Init initializes the instance.
func (s SLSBridgedSpaceCreateTileOperation) Init() SLSBridgedSpaceCreateTileOperation {
	rv := objc.Send[SLSBridgedSpaceCreateTileOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedSpaceCreateTileOperation) Autorelease() SLSBridgedSpaceCreateTileOperation {
	rv := objc.Send[SLSBridgedSpaceCreateTileOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedSpaceCreateTileOperation creates a new SLSBridgedSpaceCreateTileOperation instance.
func NewSLSBridgedSpaceCreateTileOperation() SLSBridgedSpaceCreateTileOperation {
	class := getSLSBridgedSpaceCreateTileOperationClass()
	rv := objc.Send[SLSBridgedSpaceCreateTileOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCreateTileOperation/initWithCoder:
func NewSLSBridgedSpaceCreateTileOperationWithCoder(coder objectivec.IObject) SLSBridgedSpaceCreateTileOperation {
	instance := getSLSBridgedSpaceCreateTileOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedSpaceCreateTileOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCreateTileOperation/initWithSpaceID:values:
func NewSLSBridgedSpaceCreateTileOperationWithSpaceIDValues(id uint64, values objectivec.IObject) SLSBridgedSpaceCreateTileOperation {
	instance := getSLSBridgedSpaceCreateTileOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpaceID:values:"), id, values)
	return SLSBridgedSpaceCreateTileOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCreateTileOperation/makeResultWithSpaceID:
func (s SLSBridgedSpaceCreateTileOperation) MakeResultWithSpaceID(id uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("makeResultWithSpaceID:"), id)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCreateTileOperation/initWithSpaceID:values:
func (s SLSBridgedSpaceCreateTileOperation) InitWithSpaceIDValues(id uint64, values objectivec.IObject) SLSBridgedSpaceCreateTileOperation {
	rv := objc.Send[SLSBridgedSpaceCreateTileOperation](s.ID, objc.Sel("initWithSpaceID:values:"), id, values)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCreateTileOperation/spaceID
func (s SLSBridgedSpaceCreateTileOperation) SpaceID() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCreateTileOperation/values
func (s SLSBridgedSpaceCreateTileOperation) Values() foundation.INSDictionary {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("values"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
