// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedMoveWindowsToManagedSpaceOperation] class.
var (
	_SLSBridgedMoveWindowsToManagedSpaceOperationClass     SLSBridgedMoveWindowsToManagedSpaceOperationClass
	_SLSBridgedMoveWindowsToManagedSpaceOperationClassOnce sync.Once
)

func getSLSBridgedMoveWindowsToManagedSpaceOperationClass() SLSBridgedMoveWindowsToManagedSpaceOperationClass {
	_SLSBridgedMoveWindowsToManagedSpaceOperationClassOnce.Do(func() {
		_SLSBridgedMoveWindowsToManagedSpaceOperationClass = SLSBridgedMoveWindowsToManagedSpaceOperationClass{class: objc.GetClass("SLSBridgedMoveWindowsToManagedSpaceOperation")}
	})
	return _SLSBridgedMoveWindowsToManagedSpaceOperationClass
}

// GetSLSBridgedMoveWindowsToManagedSpaceOperationClass returns the class object for SLSBridgedMoveWindowsToManagedSpaceOperation.
func GetSLSBridgedMoveWindowsToManagedSpaceOperationClass() SLSBridgedMoveWindowsToManagedSpaceOperationClass {
	return getSLSBridgedMoveWindowsToManagedSpaceOperationClass()
}

type SLSBridgedMoveWindowsToManagedSpaceOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedMoveWindowsToManagedSpaceOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedMoveWindowsToManagedSpaceOperationClass) Alloc() SLSBridgedMoveWindowsToManagedSpaceOperation {
	rv := objc.SendIfResponds[SLSBridgedMoveWindowsToManagedSpaceOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedMoveWindowsToManagedSpaceOperation.SpaceID]
//   - [SLSBridgedMoveWindowsToManagedSpaceOperation.Windows]
//   - [SLSBridgedMoveWindowsToManagedSpaceOperation.InitWithWindowsSpaceID]
type SLSBridgedMoveWindowsToManagedSpaceOperation struct {
	SLSAsynchronousBridgedWindowManagementOperation
}

// SLSBridgedMoveWindowsToManagedSpaceOperationFromID constructs a [SLSBridgedMoveWindowsToManagedSpaceOperation] from an objc.ID.
func SLSBridgedMoveWindowsToManagedSpaceOperationFromID(id objc.ID) SLSBridgedMoveWindowsToManagedSpaceOperation {
	return SLSBridgedMoveWindowsToManagedSpaceOperation{SLSAsynchronousBridgedWindowManagementOperation: SLSAsynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedMoveWindowsToManagedSpaceOperation implements ISLSBridgedMoveWindowsToManagedSpaceOperation.
var _ ISLSBridgedMoveWindowsToManagedSpaceOperation = SLSBridgedMoveWindowsToManagedSpaceOperation{}

// An interface definition for the [SLSBridgedMoveWindowsToManagedSpaceOperation] class.
//
// # Methods
//
//   - [ISLSBridgedMoveWindowsToManagedSpaceOperation.SpaceID]
//   - [ISLSBridgedMoveWindowsToManagedSpaceOperation.Windows]
//   - [ISLSBridgedMoveWindowsToManagedSpaceOperation.InitWithWindowsSpaceID]
type ISLSBridgedMoveWindowsToManagedSpaceOperation interface {
	ISLSAsynchronousBridgedWindowManagementOperation

	// Topic: Methods

	SpaceID() uint64
	Windows() foundation.INSArray
	InitWithWindowsSpaceID(windows objectivec.IObject, id uint64) SLSBridgedMoveWindowsToManagedSpaceOperation
}

// Init initializes the instance.
func (s SLSBridgedMoveWindowsToManagedSpaceOperation) Init() SLSBridgedMoveWindowsToManagedSpaceOperation {
	rv := objc.SendIfResponds[SLSBridgedMoveWindowsToManagedSpaceOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedMoveWindowsToManagedSpaceOperation) Autorelease() SLSBridgedMoveWindowsToManagedSpaceOperation {
	rv := objc.SendIfResponds[SLSBridgedMoveWindowsToManagedSpaceOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedMoveWindowsToManagedSpaceOperation creates a new SLSBridgedMoveWindowsToManagedSpaceOperation instance.
func NewSLSBridgedMoveWindowsToManagedSpaceOperation() SLSBridgedMoveWindowsToManagedSpaceOperation {
	class := getSLSBridgedMoveWindowsToManagedSpaceOperationClass()
	rv := objc.SendIfResponds[SLSBridgedMoveWindowsToManagedSpaceOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSLSBridgedMoveWindowsToManagedSpaceOperationWithCoder(coder objectivec.IObject) SLSBridgedMoveWindowsToManagedSpaceOperation {
	instance := getSLSBridgedMoveWindowsToManagedSpaceOperationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedMoveWindowsToManagedSpaceOperationFromID(rv)
}

func NewSLSBridgedMoveWindowsToManagedSpaceOperationWithWindowsSpaceID(windows objectivec.IObject, id uint64) SLSBridgedMoveWindowsToManagedSpaceOperation {
	instance := getSLSBridgedMoveWindowsToManagedSpaceOperationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithWindows:spaceID:"), windows, id)
	return SLSBridgedMoveWindowsToManagedSpaceOperationFromID(rv)
}

func (s SLSBridgedMoveWindowsToManagedSpaceOperation) InitWithWindowsSpaceID(windows objectivec.IObject, id uint64) SLSBridgedMoveWindowsToManagedSpaceOperation {
	rv := objc.SendIfResponds[SLSBridgedMoveWindowsToManagedSpaceOperation](s.ID, objc.Sel("initWithWindows:spaceID:"), windows, id)
	return rv
}

func (s SLSBridgedMoveWindowsToManagedSpaceOperation) SpaceID() uint64 {
	rv := objc.SendIfResponds[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}
func (s SLSBridgedMoveWindowsToManagedSpaceOperation) Windows() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("windows"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
