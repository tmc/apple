// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedSpaceGetAbsoluteLevelOperation] class.
var (
	_SLSBridgedSpaceGetAbsoluteLevelOperationClass     SLSBridgedSpaceGetAbsoluteLevelOperationClass
	_SLSBridgedSpaceGetAbsoluteLevelOperationClassOnce sync.Once
)

func getSLSBridgedSpaceGetAbsoluteLevelOperationClass() SLSBridgedSpaceGetAbsoluteLevelOperationClass {
	_SLSBridgedSpaceGetAbsoluteLevelOperationClassOnce.Do(func() {
		_SLSBridgedSpaceGetAbsoluteLevelOperationClass = SLSBridgedSpaceGetAbsoluteLevelOperationClass{class: objc.GetClass("SLSBridgedSpaceGetAbsoluteLevelOperation")}
	})
	return _SLSBridgedSpaceGetAbsoluteLevelOperationClass
}

// GetSLSBridgedSpaceGetAbsoluteLevelOperationClass returns the class object for SLSBridgedSpaceGetAbsoluteLevelOperation.
func GetSLSBridgedSpaceGetAbsoluteLevelOperationClass() SLSBridgedSpaceGetAbsoluteLevelOperationClass {
	return getSLSBridgedSpaceGetAbsoluteLevelOperationClass()
}

type SLSBridgedSpaceGetAbsoluteLevelOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedSpaceGetAbsoluteLevelOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedSpaceGetAbsoluteLevelOperationClass) Alloc() SLSBridgedSpaceGetAbsoluteLevelOperation {
	rv := objc.SendIfResponds[SLSBridgedSpaceGetAbsoluteLevelOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedSpaceGetAbsoluteLevelOperation.MakeResultWithInt32Value]
//   - [SLSBridgedSpaceGetAbsoluteLevelOperation.SpaceID]
//   - [SLSBridgedSpaceGetAbsoluteLevelOperation.InitWithSpaceID]
type SLSBridgedSpaceGetAbsoluteLevelOperation struct {
	SLSSynchronousBridgedWindowManagementOperation
}

// SLSBridgedSpaceGetAbsoluteLevelOperationFromID constructs a [SLSBridgedSpaceGetAbsoluteLevelOperation] from an objc.ID.
func SLSBridgedSpaceGetAbsoluteLevelOperationFromID(id objc.ID) SLSBridgedSpaceGetAbsoluteLevelOperation {
	return SLSBridgedSpaceGetAbsoluteLevelOperation{SLSSynchronousBridgedWindowManagementOperation: SLSSynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedSpaceGetAbsoluteLevelOperation implements ISLSBridgedSpaceGetAbsoluteLevelOperation.
var _ ISLSBridgedSpaceGetAbsoluteLevelOperation = SLSBridgedSpaceGetAbsoluteLevelOperation{}

// An interface definition for the [SLSBridgedSpaceGetAbsoluteLevelOperation] class.
//
// # Methods
//
//   - [ISLSBridgedSpaceGetAbsoluteLevelOperation.MakeResultWithInt32Value]
//   - [ISLSBridgedSpaceGetAbsoluteLevelOperation.SpaceID]
//   - [ISLSBridgedSpaceGetAbsoluteLevelOperation.InitWithSpaceID]
type ISLSBridgedSpaceGetAbsoluteLevelOperation interface {
	ISLSSynchronousBridgedWindowManagementOperation

	// Topic: Methods

	MakeResultWithInt32Value(int32Value int) objectivec.IObject
	SpaceID() uint64
	InitWithSpaceID(id uint64) SLSBridgedSpaceGetAbsoluteLevelOperation
}

// Init initializes the instance.
func (s SLSBridgedSpaceGetAbsoluteLevelOperation) Init() SLSBridgedSpaceGetAbsoluteLevelOperation {
	rv := objc.SendIfResponds[SLSBridgedSpaceGetAbsoluteLevelOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedSpaceGetAbsoluteLevelOperation) Autorelease() SLSBridgedSpaceGetAbsoluteLevelOperation {
	rv := objc.SendIfResponds[SLSBridgedSpaceGetAbsoluteLevelOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedSpaceGetAbsoluteLevelOperation creates a new SLSBridgedSpaceGetAbsoluteLevelOperation instance.
func NewSLSBridgedSpaceGetAbsoluteLevelOperation() SLSBridgedSpaceGetAbsoluteLevelOperation {
	class := getSLSBridgedSpaceGetAbsoluteLevelOperationClass()
	rv := objc.SendIfResponds[SLSBridgedSpaceGetAbsoluteLevelOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSLSBridgedSpaceGetAbsoluteLevelOperationWithCoder(coder objectivec.IObject) SLSBridgedSpaceGetAbsoluteLevelOperation {
	instance := getSLSBridgedSpaceGetAbsoluteLevelOperationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedSpaceGetAbsoluteLevelOperationFromID(rv)
}

func NewSLSBridgedSpaceGetAbsoluteLevelOperationWithSpaceID(id uint64) SLSBridgedSpaceGetAbsoluteLevelOperation {
	instance := getSLSBridgedSpaceGetAbsoluteLevelOperationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithSpaceID:"), id)
	return SLSBridgedSpaceGetAbsoluteLevelOperationFromID(rv)
}

func (s SLSBridgedSpaceGetAbsoluteLevelOperation) MakeResultWithInt32Value(int32Value int) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("makeResultWithInt32Value:"), int32Value)
	return objectivec.Object{ID: rv}
}
func (s SLSBridgedSpaceGetAbsoluteLevelOperation) InitWithSpaceID(id uint64) SLSBridgedSpaceGetAbsoluteLevelOperation {
	rv := objc.SendIfResponds[SLSBridgedSpaceGetAbsoluteLevelOperation](s.ID, objc.Sel("initWithSpaceID:"), id)
	return rv
}

func (s SLSBridgedSpaceGetAbsoluteLevelOperation) SpaceID() uint64 {
	rv := objc.SendIfResponds[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}
