// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedSpaceCopyNameOperation] class.
var (
	_SLSBridgedSpaceCopyNameOperationClass     SLSBridgedSpaceCopyNameOperationClass
	_SLSBridgedSpaceCopyNameOperationClassOnce sync.Once
)

func getSLSBridgedSpaceCopyNameOperationClass() SLSBridgedSpaceCopyNameOperationClass {
	_SLSBridgedSpaceCopyNameOperationClassOnce.Do(func() {
		_SLSBridgedSpaceCopyNameOperationClass = SLSBridgedSpaceCopyNameOperationClass{class: objc.GetClass("SLSBridgedSpaceCopyNameOperation")}
	})
	return _SLSBridgedSpaceCopyNameOperationClass
}

// GetSLSBridgedSpaceCopyNameOperationClass returns the class object for SLSBridgedSpaceCopyNameOperation.
func GetSLSBridgedSpaceCopyNameOperationClass() SLSBridgedSpaceCopyNameOperationClass {
	return getSLSBridgedSpaceCopyNameOperationClass()
}

type SLSBridgedSpaceCopyNameOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedSpaceCopyNameOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedSpaceCopyNameOperationClass) Alloc() SLSBridgedSpaceCopyNameOperation {
	rv := objc.SendIfResponds[SLSBridgedSpaceCopyNameOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedSpaceCopyNameOperation.MakeResultWithString]
//   - [SLSBridgedSpaceCopyNameOperation.SpaceID]
//   - [SLSBridgedSpaceCopyNameOperation.InitWithSpaceID]
type SLSBridgedSpaceCopyNameOperation struct {
	SLSSynchronousBridgedWindowManagementOperation
}

// SLSBridgedSpaceCopyNameOperationFromID constructs a [SLSBridgedSpaceCopyNameOperation] from an objc.ID.
func SLSBridgedSpaceCopyNameOperationFromID(id objc.ID) SLSBridgedSpaceCopyNameOperation {
	return SLSBridgedSpaceCopyNameOperation{SLSSynchronousBridgedWindowManagementOperation: SLSSynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedSpaceCopyNameOperation implements ISLSBridgedSpaceCopyNameOperation.
var _ ISLSBridgedSpaceCopyNameOperation = SLSBridgedSpaceCopyNameOperation{}

// An interface definition for the [SLSBridgedSpaceCopyNameOperation] class.
//
// # Methods
//
//   - [ISLSBridgedSpaceCopyNameOperation.MakeResultWithString]
//   - [ISLSBridgedSpaceCopyNameOperation.SpaceID]
//   - [ISLSBridgedSpaceCopyNameOperation.InitWithSpaceID]
type ISLSBridgedSpaceCopyNameOperation interface {
	ISLSSynchronousBridgedWindowManagementOperation

	// Topic: Methods

	MakeResultWithString(string_ objectivec.IObject) objectivec.IObject
	SpaceID() uint64
	InitWithSpaceID(id uint64) SLSBridgedSpaceCopyNameOperation
}

// Init initializes the instance.
func (s SLSBridgedSpaceCopyNameOperation) Init() SLSBridgedSpaceCopyNameOperation {
	rv := objc.SendIfResponds[SLSBridgedSpaceCopyNameOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedSpaceCopyNameOperation) Autorelease() SLSBridgedSpaceCopyNameOperation {
	rv := objc.SendIfResponds[SLSBridgedSpaceCopyNameOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedSpaceCopyNameOperation creates a new SLSBridgedSpaceCopyNameOperation instance.
func NewSLSBridgedSpaceCopyNameOperation() SLSBridgedSpaceCopyNameOperation {
	class := getSLSBridgedSpaceCopyNameOperationClass()
	rv := objc.SendIfResponds[SLSBridgedSpaceCopyNameOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSLSBridgedSpaceCopyNameOperationWithCoder(coder objectivec.IObject) SLSBridgedSpaceCopyNameOperation {
	instance := getSLSBridgedSpaceCopyNameOperationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedSpaceCopyNameOperationFromID(rv)
}

func NewSLSBridgedSpaceCopyNameOperationWithSpaceID(id uint64) SLSBridgedSpaceCopyNameOperation {
	instance := getSLSBridgedSpaceCopyNameOperationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithSpaceID:"), id)
	return SLSBridgedSpaceCopyNameOperationFromID(rv)
}

func (s SLSBridgedSpaceCopyNameOperation) MakeResultWithString(string_ objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("makeResultWithString:"), string_)
	return objectivec.Object{ID: rv}
}
func (s SLSBridgedSpaceCopyNameOperation) InitWithSpaceID(id uint64) SLSBridgedSpaceCopyNameOperation {
	rv := objc.SendIfResponds[SLSBridgedSpaceCopyNameOperation](s.ID, objc.Sel("initWithSpaceID:"), id)
	return rv
}

func (s SLSBridgedSpaceCopyNameOperation) SpaceID() uint64 {
	rv := objc.SendIfResponds[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}
