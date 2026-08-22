// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedSpaceCreateOperation] class.
var (
	_SLSBridgedSpaceCreateOperationClass     SLSBridgedSpaceCreateOperationClass
	_SLSBridgedSpaceCreateOperationClassOnce sync.Once
)

func getSLSBridgedSpaceCreateOperationClass() SLSBridgedSpaceCreateOperationClass {
	_SLSBridgedSpaceCreateOperationClassOnce.Do(func() {
		_SLSBridgedSpaceCreateOperationClass = SLSBridgedSpaceCreateOperationClass{class: objc.GetClass("SLSBridgedSpaceCreateOperation")}
	})
	return _SLSBridgedSpaceCreateOperationClass
}

// GetSLSBridgedSpaceCreateOperationClass returns the class object for SLSBridgedSpaceCreateOperation.
func GetSLSBridgedSpaceCreateOperationClass() SLSBridgedSpaceCreateOperationClass {
	return getSLSBridgedSpaceCreateOperationClass()
}

type SLSBridgedSpaceCreateOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedSpaceCreateOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedSpaceCreateOperationClass) Alloc() SLSBridgedSpaceCreateOperation {
	rv := objc.SendIfResponds[SLSBridgedSpaceCreateOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedSpaceCreateOperation.MakeResultWithSpaceID]
//   - [SLSBridgedSpaceCreateOperation.Options]
//   - [SLSBridgedSpaceCreateOperation.Values]
//   - [SLSBridgedSpaceCreateOperation.InitWithOptionsValues]
type SLSBridgedSpaceCreateOperation struct {
	SLSSynchronousBridgedWindowManagementOperation
}

// SLSBridgedSpaceCreateOperationFromID constructs a [SLSBridgedSpaceCreateOperation] from an objc.ID.
func SLSBridgedSpaceCreateOperationFromID(id objc.ID) SLSBridgedSpaceCreateOperation {
	return SLSBridgedSpaceCreateOperation{SLSSynchronousBridgedWindowManagementOperation: SLSSynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedSpaceCreateOperation implements ISLSBridgedSpaceCreateOperation.
var _ ISLSBridgedSpaceCreateOperation = SLSBridgedSpaceCreateOperation{}

// An interface definition for the [SLSBridgedSpaceCreateOperation] class.
//
// # Methods
//
//   - [ISLSBridgedSpaceCreateOperation.MakeResultWithSpaceID]
//   - [ISLSBridgedSpaceCreateOperation.Options]
//   - [ISLSBridgedSpaceCreateOperation.Values]
//   - [ISLSBridgedSpaceCreateOperation.InitWithOptionsValues]
type ISLSBridgedSpaceCreateOperation interface {
	ISLSSynchronousBridgedWindowManagementOperation

	// Topic: Methods

	MakeResultWithSpaceID(id uint64) objectivec.IObject
	Options() uint32
	Values() foundation.INSDictionary
	InitWithOptionsValues(options uint32, values objectivec.IObject) SLSBridgedSpaceCreateOperation
}

// Init initializes the instance.
func (s SLSBridgedSpaceCreateOperation) Init() SLSBridgedSpaceCreateOperation {
	rv := objc.SendIfResponds[SLSBridgedSpaceCreateOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedSpaceCreateOperation) Autorelease() SLSBridgedSpaceCreateOperation {
	rv := objc.SendIfResponds[SLSBridgedSpaceCreateOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedSpaceCreateOperation creates a new SLSBridgedSpaceCreateOperation instance.
func NewSLSBridgedSpaceCreateOperation() SLSBridgedSpaceCreateOperation {
	class := getSLSBridgedSpaceCreateOperationClass()
	rv := objc.SendIfResponds[SLSBridgedSpaceCreateOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSLSBridgedSpaceCreateOperationWithCoder(coder objectivec.IObject) SLSBridgedSpaceCreateOperation {
	instance := getSLSBridgedSpaceCreateOperationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedSpaceCreateOperationFromID(rv)
}

func NewSLSBridgedSpaceCreateOperationWithOptionsValues(options uint32, values objectivec.IObject) SLSBridgedSpaceCreateOperation {
	instance := getSLSBridgedSpaceCreateOperationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithOptions:values:"), options, values)
	return SLSBridgedSpaceCreateOperationFromID(rv)
}

func (s SLSBridgedSpaceCreateOperation) MakeResultWithSpaceID(id uint64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("makeResultWithSpaceID:"), id)
	return objectivec.Object{ID: rv}
}
func (s SLSBridgedSpaceCreateOperation) InitWithOptionsValues(options uint32, values objectivec.IObject) SLSBridgedSpaceCreateOperation {
	rv := objc.SendIfResponds[SLSBridgedSpaceCreateOperation](s.ID, objc.Sel("initWithOptions:values:"), options, values)
	return rv
}

func (s SLSBridgedSpaceCreateOperation) Options() uint32 {
	rv := objc.SendIfResponds[uint32](s.ID, objc.Sel("options"))
	return rv
}
func (s SLSBridgedSpaceCreateOperation) Values() foundation.INSDictionary {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("values"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
