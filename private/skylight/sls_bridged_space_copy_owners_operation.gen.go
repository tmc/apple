// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedSpaceCopyOwnersOperation] class.
var (
	_SLSBridgedSpaceCopyOwnersOperationClass     SLSBridgedSpaceCopyOwnersOperationClass
	_SLSBridgedSpaceCopyOwnersOperationClassOnce sync.Once
)

func getSLSBridgedSpaceCopyOwnersOperationClass() SLSBridgedSpaceCopyOwnersOperationClass {
	_SLSBridgedSpaceCopyOwnersOperationClassOnce.Do(func() {
		_SLSBridgedSpaceCopyOwnersOperationClass = SLSBridgedSpaceCopyOwnersOperationClass{class: objc.GetClass("SLSBridgedSpaceCopyOwnersOperation")}
	})
	return _SLSBridgedSpaceCopyOwnersOperationClass
}

// GetSLSBridgedSpaceCopyOwnersOperationClass returns the class object for SLSBridgedSpaceCopyOwnersOperation.
func GetSLSBridgedSpaceCopyOwnersOperationClass() SLSBridgedSpaceCopyOwnersOperationClass {
	return getSLSBridgedSpaceCopyOwnersOperationClass()
}

type SLSBridgedSpaceCopyOwnersOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedSpaceCopyOwnersOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedSpaceCopyOwnersOperationClass) Alloc() SLSBridgedSpaceCopyOwnersOperation {
	rv := objc.SendIfResponds[SLSBridgedSpaceCopyOwnersOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedSpaceCopyOwnersOperation.MakeResultWithNumbers]
//   - [SLSBridgedSpaceCopyOwnersOperation.SpaceID]
//   - [SLSBridgedSpaceCopyOwnersOperation.InitWithSpaceID]
type SLSBridgedSpaceCopyOwnersOperation struct {
	SLSSynchronousBridgedWindowManagementOperation
}

// SLSBridgedSpaceCopyOwnersOperationFromID constructs a [SLSBridgedSpaceCopyOwnersOperation] from an objc.ID.
func SLSBridgedSpaceCopyOwnersOperationFromID(id objc.ID) SLSBridgedSpaceCopyOwnersOperation {
	return SLSBridgedSpaceCopyOwnersOperation{SLSSynchronousBridgedWindowManagementOperation: SLSSynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedSpaceCopyOwnersOperation implements ISLSBridgedSpaceCopyOwnersOperation.
var _ ISLSBridgedSpaceCopyOwnersOperation = SLSBridgedSpaceCopyOwnersOperation{}

// An interface definition for the [SLSBridgedSpaceCopyOwnersOperation] class.
//
// # Methods
//
//   - [ISLSBridgedSpaceCopyOwnersOperation.MakeResultWithNumbers]
//   - [ISLSBridgedSpaceCopyOwnersOperation.SpaceID]
//   - [ISLSBridgedSpaceCopyOwnersOperation.InitWithSpaceID]
type ISLSBridgedSpaceCopyOwnersOperation interface {
	ISLSSynchronousBridgedWindowManagementOperation

	// Topic: Methods

	MakeResultWithNumbers(numbers objectivec.IObject) objectivec.IObject
	SpaceID() uint64
	InitWithSpaceID(id uint64) SLSBridgedSpaceCopyOwnersOperation
}

// Init initializes the instance.
func (s SLSBridgedSpaceCopyOwnersOperation) Init() SLSBridgedSpaceCopyOwnersOperation {
	rv := objc.SendIfResponds[SLSBridgedSpaceCopyOwnersOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedSpaceCopyOwnersOperation) Autorelease() SLSBridgedSpaceCopyOwnersOperation {
	rv := objc.SendIfResponds[SLSBridgedSpaceCopyOwnersOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedSpaceCopyOwnersOperation creates a new SLSBridgedSpaceCopyOwnersOperation instance.
func NewSLSBridgedSpaceCopyOwnersOperation() SLSBridgedSpaceCopyOwnersOperation {
	class := getSLSBridgedSpaceCopyOwnersOperationClass()
	rv := objc.SendIfResponds[SLSBridgedSpaceCopyOwnersOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSLSBridgedSpaceCopyOwnersOperationWithCoder(coder objectivec.IObject) SLSBridgedSpaceCopyOwnersOperation {
	instance := getSLSBridgedSpaceCopyOwnersOperationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedSpaceCopyOwnersOperationFromID(rv)
}

func NewSLSBridgedSpaceCopyOwnersOperationWithSpaceID(id uint64) SLSBridgedSpaceCopyOwnersOperation {
	instance := getSLSBridgedSpaceCopyOwnersOperationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithSpaceID:"), id)
	return SLSBridgedSpaceCopyOwnersOperationFromID(rv)
}

func (s SLSBridgedSpaceCopyOwnersOperation) MakeResultWithNumbers(numbers objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("makeResultWithNumbers:"), numbers)
	return objectivec.Object{ID: rv}
}
func (s SLSBridgedSpaceCopyOwnersOperation) InitWithSpaceID(id uint64) SLSBridgedSpaceCopyOwnersOperation {
	rv := objc.SendIfResponds[SLSBridgedSpaceCopyOwnersOperation](s.ID, objc.Sel("initWithSpaceID:"), id)
	return rv
}

func (s SLSBridgedSpaceCopyOwnersOperation) SpaceID() uint64 {
	rv := objc.SendIfResponds[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}
