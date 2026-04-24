// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedSpaceGetSizeForProposedTileOperation] class.
var (
	_SLSBridgedSpaceGetSizeForProposedTileOperationClass     SLSBridgedSpaceGetSizeForProposedTileOperationClass
	_SLSBridgedSpaceGetSizeForProposedTileOperationClassOnce sync.Once
)

func getSLSBridgedSpaceGetSizeForProposedTileOperationClass() SLSBridgedSpaceGetSizeForProposedTileOperationClass {
	_SLSBridgedSpaceGetSizeForProposedTileOperationClassOnce.Do(func() {
		_SLSBridgedSpaceGetSizeForProposedTileOperationClass = SLSBridgedSpaceGetSizeForProposedTileOperationClass{class: objc.GetClass("SLSBridgedSpaceGetSizeForProposedTileOperation")}
	})
	return _SLSBridgedSpaceGetSizeForProposedTileOperationClass
}

// GetSLSBridgedSpaceGetSizeForProposedTileOperationClass returns the class object for SLSBridgedSpaceGetSizeForProposedTileOperation.
func GetSLSBridgedSpaceGetSizeForProposedTileOperationClass() SLSBridgedSpaceGetSizeForProposedTileOperationClass {
	return getSLSBridgedSpaceGetSizeForProposedTileOperationClass()
}

type SLSBridgedSpaceGetSizeForProposedTileOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedSpaceGetSizeForProposedTileOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedSpaceGetSizeForProposedTileOperationClass) Alloc() SLSBridgedSpaceGetSizeForProposedTileOperation {
	rv := objc.Send[SLSBridgedSpaceGetSizeForProposedTileOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedSpaceGetSizeForProposedTileOperation.MakeResultWithSize]
//   - [SLSBridgedSpaceGetSizeForProposedTileOperation.SpaceID]
//   - [SLSBridgedSpaceGetSizeForProposedTileOperation.Values]
//   - [SLSBridgedSpaceGetSizeForProposedTileOperation.InitWithSpaceIDValues]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceGetSizeForProposedTileOperation
type SLSBridgedSpaceGetSizeForProposedTileOperation struct {
	SLSSynchronousBridgedWindowManagementOperation
}

// SLSBridgedSpaceGetSizeForProposedTileOperationFromID constructs a [SLSBridgedSpaceGetSizeForProposedTileOperation] from an objc.ID.
func SLSBridgedSpaceGetSizeForProposedTileOperationFromID(id objc.ID) SLSBridgedSpaceGetSizeForProposedTileOperation {
	return SLSBridgedSpaceGetSizeForProposedTileOperation{SLSSynchronousBridgedWindowManagementOperation: SLSSynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedSpaceGetSizeForProposedTileOperation implements ISLSBridgedSpaceGetSizeForProposedTileOperation.
var _ ISLSBridgedSpaceGetSizeForProposedTileOperation = SLSBridgedSpaceGetSizeForProposedTileOperation{}

// An interface definition for the [SLSBridgedSpaceGetSizeForProposedTileOperation] class.
//
// # Methods
//
//   - [ISLSBridgedSpaceGetSizeForProposedTileOperation.MakeResultWithSize]
//   - [ISLSBridgedSpaceGetSizeForProposedTileOperation.SpaceID]
//   - [ISLSBridgedSpaceGetSizeForProposedTileOperation.Values]
//   - [ISLSBridgedSpaceGetSizeForProposedTileOperation.InitWithSpaceIDValues]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceGetSizeForProposedTileOperation
type ISLSBridgedSpaceGetSizeForProposedTileOperation interface {
	ISLSSynchronousBridgedWindowManagementOperation

	// Topic: Methods

	MakeResultWithSize(size corefoundation.CGSize) objectivec.IObject
	SpaceID() uint64
	Values() foundation.INSDictionary
	InitWithSpaceIDValues(id uint64, values objectivec.IObject) SLSBridgedSpaceGetSizeForProposedTileOperation
}

// Init initializes the instance.
func (s SLSBridgedSpaceGetSizeForProposedTileOperation) Init() SLSBridgedSpaceGetSizeForProposedTileOperation {
	rv := objc.Send[SLSBridgedSpaceGetSizeForProposedTileOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedSpaceGetSizeForProposedTileOperation) Autorelease() SLSBridgedSpaceGetSizeForProposedTileOperation {
	rv := objc.Send[SLSBridgedSpaceGetSizeForProposedTileOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedSpaceGetSizeForProposedTileOperation creates a new SLSBridgedSpaceGetSizeForProposedTileOperation instance.
func NewSLSBridgedSpaceGetSizeForProposedTileOperation() SLSBridgedSpaceGetSizeForProposedTileOperation {
	class := getSLSBridgedSpaceGetSizeForProposedTileOperationClass()
	rv := objc.Send[SLSBridgedSpaceGetSizeForProposedTileOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceGetSizeForProposedTileOperation/initWithCoder:
func NewSLSBridgedSpaceGetSizeForProposedTileOperationWithCoder(coder objectivec.IObject) SLSBridgedSpaceGetSizeForProposedTileOperation {
	instance := getSLSBridgedSpaceGetSizeForProposedTileOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedSpaceGetSizeForProposedTileOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceGetSizeForProposedTileOperation/initWithSpaceID:values:
func NewSLSBridgedSpaceGetSizeForProposedTileOperationWithSpaceIDValues(id uint64, values objectivec.IObject) SLSBridgedSpaceGetSizeForProposedTileOperation {
	instance := getSLSBridgedSpaceGetSizeForProposedTileOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpaceID:values:"), id, values)
	return SLSBridgedSpaceGetSizeForProposedTileOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceGetSizeForProposedTileOperation/makeResultWithSize:
func (s SLSBridgedSpaceGetSizeForProposedTileOperation) MakeResultWithSize(size corefoundation.CGSize) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("makeResultWithSize:"), size)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceGetSizeForProposedTileOperation/initWithSpaceID:values:
func (s SLSBridgedSpaceGetSizeForProposedTileOperation) InitWithSpaceIDValues(id uint64, values objectivec.IObject) SLSBridgedSpaceGetSizeForProposedTileOperation {
	rv := objc.Send[SLSBridgedSpaceGetSizeForProposedTileOperation](s.ID, objc.Sel("initWithSpaceID:values:"), id, values)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceGetSizeForProposedTileOperation/spaceID
func (s SLSBridgedSpaceGetSizeForProposedTileOperation) SpaceID() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceGetSizeForProposedTileOperation/values
func (s SLSBridgedSpaceGetSizeForProposedTileOperation) Values() foundation.INSDictionary {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("values"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
