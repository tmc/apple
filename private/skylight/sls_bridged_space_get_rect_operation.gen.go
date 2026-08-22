// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedSpaceGetRectOperation] class.
var (
	_SLSBridgedSpaceGetRectOperationClass     SLSBridgedSpaceGetRectOperationClass
	_SLSBridgedSpaceGetRectOperationClassOnce sync.Once
)

func getSLSBridgedSpaceGetRectOperationClass() SLSBridgedSpaceGetRectOperationClass {
	_SLSBridgedSpaceGetRectOperationClassOnce.Do(func() {
		_SLSBridgedSpaceGetRectOperationClass = SLSBridgedSpaceGetRectOperationClass{class: objc.GetClass("SLSBridgedSpaceGetRectOperation")}
	})
	return _SLSBridgedSpaceGetRectOperationClass
}

// GetSLSBridgedSpaceGetRectOperationClass returns the class object for SLSBridgedSpaceGetRectOperation.
func GetSLSBridgedSpaceGetRectOperationClass() SLSBridgedSpaceGetRectOperationClass {
	return getSLSBridgedSpaceGetRectOperationClass()
}

type SLSBridgedSpaceGetRectOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedSpaceGetRectOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedSpaceGetRectOperationClass) Alloc() SLSBridgedSpaceGetRectOperation {
	rv := objc.SendIfResponds[SLSBridgedSpaceGetRectOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedSpaceGetRectOperation.MakeResultWithRect]
//   - [SLSBridgedSpaceGetRectOperation.SpaceID]
//   - [SLSBridgedSpaceGetRectOperation.InitWithSpaceID]
type SLSBridgedSpaceGetRectOperation struct {
	SLSSynchronousBridgedWindowManagementOperation
}

// SLSBridgedSpaceGetRectOperationFromID constructs a [SLSBridgedSpaceGetRectOperation] from an objc.ID.
func SLSBridgedSpaceGetRectOperationFromID(id objc.ID) SLSBridgedSpaceGetRectOperation {
	return SLSBridgedSpaceGetRectOperation{SLSSynchronousBridgedWindowManagementOperation: SLSSynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedSpaceGetRectOperation implements ISLSBridgedSpaceGetRectOperation.
var _ ISLSBridgedSpaceGetRectOperation = SLSBridgedSpaceGetRectOperation{}

// An interface definition for the [SLSBridgedSpaceGetRectOperation] class.
//
// # Methods
//
//   - [ISLSBridgedSpaceGetRectOperation.MakeResultWithRect]
//   - [ISLSBridgedSpaceGetRectOperation.SpaceID]
//   - [ISLSBridgedSpaceGetRectOperation.InitWithSpaceID]
type ISLSBridgedSpaceGetRectOperation interface {
	ISLSSynchronousBridgedWindowManagementOperation

	// Topic: Methods

	MakeResultWithRect(rect corefoundation.CGRect) objectivec.IObject
	SpaceID() uint64
	InitWithSpaceID(id uint64) SLSBridgedSpaceGetRectOperation
}

// Init initializes the instance.
func (s SLSBridgedSpaceGetRectOperation) Init() SLSBridgedSpaceGetRectOperation {
	rv := objc.SendIfResponds[SLSBridgedSpaceGetRectOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedSpaceGetRectOperation) Autorelease() SLSBridgedSpaceGetRectOperation {
	rv := objc.SendIfResponds[SLSBridgedSpaceGetRectOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedSpaceGetRectOperation creates a new SLSBridgedSpaceGetRectOperation instance.
func NewSLSBridgedSpaceGetRectOperation() SLSBridgedSpaceGetRectOperation {
	class := getSLSBridgedSpaceGetRectOperationClass()
	rv := objc.SendIfResponds[SLSBridgedSpaceGetRectOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSLSBridgedSpaceGetRectOperationWithCoder(coder objectivec.IObject) SLSBridgedSpaceGetRectOperation {
	instance := getSLSBridgedSpaceGetRectOperationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedSpaceGetRectOperationFromID(rv)
}

func NewSLSBridgedSpaceGetRectOperationWithSpaceID(id uint64) SLSBridgedSpaceGetRectOperation {
	instance := getSLSBridgedSpaceGetRectOperationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithSpaceID:"), id)
	return SLSBridgedSpaceGetRectOperationFromID(rv)
}

func (s SLSBridgedSpaceGetRectOperation) MakeResultWithRect(rect corefoundation.CGRect) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("makeResultWithRect:"), rect)
	return objectivec.Object{ID: rv}
}
func (s SLSBridgedSpaceGetRectOperation) InitWithSpaceID(id uint64) SLSBridgedSpaceGetRectOperation {
	rv := objc.SendIfResponds[SLSBridgedSpaceGetRectOperation](s.ID, objc.Sel("initWithSpaceID:"), id)
	return rv
}

func (s SLSBridgedSpaceGetRectOperation) SpaceID() uint64 {
	rv := objc.SendIfResponds[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}
