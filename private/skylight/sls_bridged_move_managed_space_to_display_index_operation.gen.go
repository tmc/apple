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
	rv := objc.SendIfResponds[SLSBridgedMoveManagedSpaceToDisplayIndexOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedMoveManagedSpaceToDisplayIndexOperation.DisplayIdentifier]
//   - [SLSBridgedMoveManagedSpaceToDisplayIndexOperation.Index]
//   - [SLSBridgedMoveManagedSpaceToDisplayIndexOperation.SpaceID]
//   - [SLSBridgedMoveManagedSpaceToDisplayIndexOperation.InitWithSpaceIDDisplayIdentifierIndex]
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
	rv := objc.SendIfResponds[SLSBridgedMoveManagedSpaceToDisplayIndexOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedMoveManagedSpaceToDisplayIndexOperation) Autorelease() SLSBridgedMoveManagedSpaceToDisplayIndexOperation {
	rv := objc.SendIfResponds[SLSBridgedMoveManagedSpaceToDisplayIndexOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedMoveManagedSpaceToDisplayIndexOperation creates a new SLSBridgedMoveManagedSpaceToDisplayIndexOperation instance.
func NewSLSBridgedMoveManagedSpaceToDisplayIndexOperation() SLSBridgedMoveManagedSpaceToDisplayIndexOperation {
	class := getSLSBridgedMoveManagedSpaceToDisplayIndexOperationClass()
	rv := objc.SendIfResponds[SLSBridgedMoveManagedSpaceToDisplayIndexOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSLSBridgedMoveManagedSpaceToDisplayIndexOperationWithCoder(coder objectivec.IObject) SLSBridgedMoveManagedSpaceToDisplayIndexOperation {
	instance := getSLSBridgedMoveManagedSpaceToDisplayIndexOperationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedMoveManagedSpaceToDisplayIndexOperationFromID(rv)
}

func NewSLSBridgedMoveManagedSpaceToDisplayIndexOperationWithSpaceIDDisplayIdentifierIndex(id uint64, identifier objectivec.IObject, index uint32) SLSBridgedMoveManagedSpaceToDisplayIndexOperation {
	instance := getSLSBridgedMoveManagedSpaceToDisplayIndexOperationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithSpaceID:displayIdentifier:index:"), id, identifier, index)
	return SLSBridgedMoveManagedSpaceToDisplayIndexOperationFromID(rv)
}

func (s SLSBridgedMoveManagedSpaceToDisplayIndexOperation) InitWithSpaceIDDisplayIdentifierIndex(id uint64, identifier objectivec.IObject, index uint32) SLSBridgedMoveManagedSpaceToDisplayIndexOperation {
	rv := objc.SendIfResponds[SLSBridgedMoveManagedSpaceToDisplayIndexOperation](s.ID, objc.Sel("initWithSpaceID:displayIdentifier:index:"), id, identifier, index)
	return rv
}

func (s SLSBridgedMoveManagedSpaceToDisplayIndexOperation) DisplayIdentifier() string {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("displayIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (s SLSBridgedMoveManagedSpaceToDisplayIndexOperation) Index() uint32 {
	rv := objc.SendIfResponds[uint32](s.ID, objc.Sel("index"))
	return rv
}
func (s SLSBridgedMoveManagedSpaceToDisplayIndexOperation) SpaceID() uint64 {
	rv := objc.SendIfResponds[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}
