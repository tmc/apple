// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedSpaceCopyValuesOperation] class.
var (
	_SLSBridgedSpaceCopyValuesOperationClass     SLSBridgedSpaceCopyValuesOperationClass
	_SLSBridgedSpaceCopyValuesOperationClassOnce sync.Once
)

func getSLSBridgedSpaceCopyValuesOperationClass() SLSBridgedSpaceCopyValuesOperationClass {
	_SLSBridgedSpaceCopyValuesOperationClassOnce.Do(func() {
		_SLSBridgedSpaceCopyValuesOperationClass = SLSBridgedSpaceCopyValuesOperationClass{class: objc.GetClass("SLSBridgedSpaceCopyValuesOperation")}
	})
	return _SLSBridgedSpaceCopyValuesOperationClass
}

// GetSLSBridgedSpaceCopyValuesOperationClass returns the class object for SLSBridgedSpaceCopyValuesOperation.
func GetSLSBridgedSpaceCopyValuesOperationClass() SLSBridgedSpaceCopyValuesOperationClass {
	return getSLSBridgedSpaceCopyValuesOperationClass()
}

type SLSBridgedSpaceCopyValuesOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedSpaceCopyValuesOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedSpaceCopyValuesOperationClass) Alloc() SLSBridgedSpaceCopyValuesOperation {
	rv := objc.SendIfResponds[SLSBridgedSpaceCopyValuesOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedSpaceCopyValuesOperation.MakeResultWithPropertyListDictionary]
//   - [SLSBridgedSpaceCopyValuesOperation.SpaceID]
//   - [SLSBridgedSpaceCopyValuesOperation.InitWithSpaceID]
type SLSBridgedSpaceCopyValuesOperation struct {
	SLSSynchronousBridgedWindowManagementOperation
}

// SLSBridgedSpaceCopyValuesOperationFromID constructs a [SLSBridgedSpaceCopyValuesOperation] from an objc.ID.
func SLSBridgedSpaceCopyValuesOperationFromID(id objc.ID) SLSBridgedSpaceCopyValuesOperation {
	return SLSBridgedSpaceCopyValuesOperation{SLSSynchronousBridgedWindowManagementOperation: SLSSynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedSpaceCopyValuesOperation implements ISLSBridgedSpaceCopyValuesOperation.
var _ ISLSBridgedSpaceCopyValuesOperation = SLSBridgedSpaceCopyValuesOperation{}

// An interface definition for the [SLSBridgedSpaceCopyValuesOperation] class.
//
// # Methods
//
//   - [ISLSBridgedSpaceCopyValuesOperation.MakeResultWithPropertyListDictionary]
//   - [ISLSBridgedSpaceCopyValuesOperation.SpaceID]
//   - [ISLSBridgedSpaceCopyValuesOperation.InitWithSpaceID]
type ISLSBridgedSpaceCopyValuesOperation interface {
	ISLSSynchronousBridgedWindowManagementOperation

	// Topic: Methods

	MakeResultWithPropertyListDictionary(dictionary objectivec.IObject) objectivec.IObject
	SpaceID() uint64
	InitWithSpaceID(id uint64) SLSBridgedSpaceCopyValuesOperation
}

// Init initializes the instance.
func (s SLSBridgedSpaceCopyValuesOperation) Init() SLSBridgedSpaceCopyValuesOperation {
	rv := objc.SendIfResponds[SLSBridgedSpaceCopyValuesOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedSpaceCopyValuesOperation) Autorelease() SLSBridgedSpaceCopyValuesOperation {
	rv := objc.SendIfResponds[SLSBridgedSpaceCopyValuesOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedSpaceCopyValuesOperation creates a new SLSBridgedSpaceCopyValuesOperation instance.
func NewSLSBridgedSpaceCopyValuesOperation() SLSBridgedSpaceCopyValuesOperation {
	class := getSLSBridgedSpaceCopyValuesOperationClass()
	rv := objc.SendIfResponds[SLSBridgedSpaceCopyValuesOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSLSBridgedSpaceCopyValuesOperationWithCoder(coder objectivec.IObject) SLSBridgedSpaceCopyValuesOperation {
	instance := getSLSBridgedSpaceCopyValuesOperationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedSpaceCopyValuesOperationFromID(rv)
}

func NewSLSBridgedSpaceCopyValuesOperationWithSpaceID(id uint64) SLSBridgedSpaceCopyValuesOperation {
	instance := getSLSBridgedSpaceCopyValuesOperationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithSpaceID:"), id)
	return SLSBridgedSpaceCopyValuesOperationFromID(rv)
}

func (s SLSBridgedSpaceCopyValuesOperation) MakeResultWithPropertyListDictionary(dictionary objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("makeResultWithPropertyListDictionary:"), dictionary)
	return objectivec.Object{ID: rv}
}
func (s SLSBridgedSpaceCopyValuesOperation) InitWithSpaceID(id uint64) SLSBridgedSpaceCopyValuesOperation {
	rv := objc.SendIfResponds[SLSBridgedSpaceCopyValuesOperation](s.ID, objc.Sel("initWithSpaceID:"), id)
	return rv
}

func (s SLSBridgedSpaceCopyValuesOperation) SpaceID() uint64 {
	rv := objc.SendIfResponds[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}
