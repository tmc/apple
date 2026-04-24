// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedMoveManagedSpaceToDisplayIndexOperation] class.
var (
	_SLSBridgedMoveManagedSpaceToDisplayIndexOperationClass     SLSBridgedMoveManagedSpaceToDisplayIndexOperationClass
	_SLSBridgedMoveManagedSpaceToDisplayIndexOperationClassOnce sync.Once
)

func getSLSBridgedMoveManagedSpaceToDisplayIndexOperationClass() SLSBridgedMoveManagedSpaceToDisplayIndexOperationClass {
	_SLSBridgedMoveManagedSpaceToDisplayIndexOperationClassOnce.Do(func() {
		_SLSBridgedMoveManagedSpaceToDisplayIndexOperationClass = SLSBridgedMoveManagedSpaceToDisplayIndexOperationClass{class: objc.GetClass("SLSBridgedMoveManagedSpaceToDisplayIndexOperation")}
	})
	return _SLSBridgedMoveManagedSpaceToDisplayIndexOperationClass
}

// GetSLSBridgedMoveManagedSpaceToDisplayIndexOperationClass returns the class object for SLSBridgedMoveManagedSpaceToDisplayIndexOperation.
func GetSLSBridgedMoveManagedSpaceToDisplayIndexOperationClass() SLSBridgedMoveManagedSpaceToDisplayIndexOperationClass {
	return getSLSBridgedMoveManagedSpaceToDisplayIndexOperationClass()
}

type SLSBridgedMoveManagedSpaceToDisplayIndexOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedMoveManagedSpaceToDisplayIndexOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedMoveManagedSpaceToDisplayIndexOperationClass) Alloc() SLSBridgedMoveManagedSpaceToDisplayIndexOperation {
	rv := objc.Send[SLSBridgedMoveManagedSpaceToDisplayIndexOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedMoveManagedSpaceToDisplayIndexOperation.DisplayIdentifier]
//   - [SLSBridgedMoveManagedSpaceToDisplayIndexOperation.Index]
//   - [SLSBridgedMoveManagedSpaceToDisplayIndexOperation.SpaceID]
//   - [SLSBridgedMoveManagedSpaceToDisplayIndexOperation.InitWithSpaceIDDisplayIdentifierIndex]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedMoveManagedSpaceToDisplayIndexOperation
type SLSBridgedMoveManagedSpaceToDisplayIndexOperation struct {
	SLSAsynchronousBridgedWindowManagementOperation
}

// SLSBridgedMoveManagedSpaceToDisplayIndexOperationFromID constructs a [SLSBridgedMoveManagedSpaceToDisplayIndexOperation] from an objc.ID.
func SLSBridgedMoveManagedSpaceToDisplayIndexOperationFromID(id objc.ID) SLSBridgedMoveManagedSpaceToDisplayIndexOperation {
	return SLSBridgedMoveManagedSpaceToDisplayIndexOperation{SLSAsynchronousBridgedWindowManagementOperation: SLSAsynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedMoveManagedSpaceToDisplayIndexOperation implements ISLSBridgedMoveManagedSpaceToDisplayIndexOperation.
var _ ISLSBridgedMoveManagedSpaceToDisplayIndexOperation = SLSBridgedMoveManagedSpaceToDisplayIndexOperation{}

// An interface definition for the [SLSBridgedMoveManagedSpaceToDisplayIndexOperation] class.
//
// # Methods
//
//   - [ISLSBridgedMoveManagedSpaceToDisplayIndexOperation.DisplayIdentifier]
//   - [ISLSBridgedMoveManagedSpaceToDisplayIndexOperation.Index]
//   - [ISLSBridgedMoveManagedSpaceToDisplayIndexOperation.SpaceID]
//   - [ISLSBridgedMoveManagedSpaceToDisplayIndexOperation.InitWithSpaceIDDisplayIdentifierIndex]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedMoveManagedSpaceToDisplayIndexOperation
type ISLSBridgedMoveManagedSpaceToDisplayIndexOperation interface {
	ISLSAsynchronousBridgedWindowManagementOperation

	// Topic: Methods

	DisplayIdentifier() string
	Index() uint32
	SpaceID() uint64
	InitWithSpaceIDDisplayIdentifierIndex(id uint64, identifier objectivec.IObject, index uint32) SLSBridgedMoveManagedSpaceToDisplayIndexOperation
}

// Init initializes the instance.
func (s SLSBridgedMoveManagedSpaceToDisplayIndexOperation) Init() SLSBridgedMoveManagedSpaceToDisplayIndexOperation {
	rv := objc.Send[SLSBridgedMoveManagedSpaceToDisplayIndexOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedMoveManagedSpaceToDisplayIndexOperation) Autorelease() SLSBridgedMoveManagedSpaceToDisplayIndexOperation {
	rv := objc.Send[SLSBridgedMoveManagedSpaceToDisplayIndexOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedMoveManagedSpaceToDisplayIndexOperation creates a new SLSBridgedMoveManagedSpaceToDisplayIndexOperation instance.
func NewSLSBridgedMoveManagedSpaceToDisplayIndexOperation() SLSBridgedMoveManagedSpaceToDisplayIndexOperation {
	class := getSLSBridgedMoveManagedSpaceToDisplayIndexOperationClass()
	rv := objc.Send[SLSBridgedMoveManagedSpaceToDisplayIndexOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedMoveManagedSpaceToDisplayIndexOperation/initWithCoder:
func NewSLSBridgedMoveManagedSpaceToDisplayIndexOperationWithCoder(coder objectivec.IObject) SLSBridgedMoveManagedSpaceToDisplayIndexOperation {
	instance := getSLSBridgedMoveManagedSpaceToDisplayIndexOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedMoveManagedSpaceToDisplayIndexOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedMoveManagedSpaceToDisplayIndexOperation/initWithSpaceID:displayIdentifier:index:
func NewSLSBridgedMoveManagedSpaceToDisplayIndexOperationWithSpaceIDDisplayIdentifierIndex(id uint64, identifier objectivec.IObject, index uint32) SLSBridgedMoveManagedSpaceToDisplayIndexOperation {
	instance := getSLSBridgedMoveManagedSpaceToDisplayIndexOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpaceID:displayIdentifier:index:"), id, identifier, index)
	return SLSBridgedMoveManagedSpaceToDisplayIndexOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedMoveManagedSpaceToDisplayIndexOperation/initWithSpaceID:displayIdentifier:index:
func (s SLSBridgedMoveManagedSpaceToDisplayIndexOperation) InitWithSpaceIDDisplayIdentifierIndex(id uint64, identifier objectivec.IObject, index uint32) SLSBridgedMoveManagedSpaceToDisplayIndexOperation {
	rv := objc.Send[SLSBridgedMoveManagedSpaceToDisplayIndexOperation](s.ID, objc.Sel("initWithSpaceID:displayIdentifier:index:"), id, identifier, index)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedMoveManagedSpaceToDisplayIndexOperation/displayIdentifier
func (s SLSBridgedMoveManagedSpaceToDisplayIndexOperation) DisplayIdentifier() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("displayIdentifier"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedMoveManagedSpaceToDisplayIndexOperation/index
func (s SLSBridgedMoveManagedSpaceToDisplayIndexOperation) Index() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("index"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedMoveManagedSpaceToDisplayIndexOperation/spaceID
func (s SLSBridgedMoveManagedSpaceToDisplayIndexOperation) SpaceID() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}
