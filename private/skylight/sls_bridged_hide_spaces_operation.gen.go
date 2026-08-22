// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedHideSpacesOperation] class.
var (
	_SLSBridgedHideSpacesOperationClass     SLSBridgedHideSpacesOperationClass
	_SLSBridgedHideSpacesOperationClassOnce sync.Once
)

func getSLSBridgedHideSpacesOperationClass() SLSBridgedHideSpacesOperationClass {
	_SLSBridgedHideSpacesOperationClassOnce.Do(func() {
		_SLSBridgedHideSpacesOperationClass = SLSBridgedHideSpacesOperationClass{class: objc.GetClass("SLSBridgedHideSpacesOperation")}
	})
	return _SLSBridgedHideSpacesOperationClass
}

// GetSLSBridgedHideSpacesOperationClass returns the class object for SLSBridgedHideSpacesOperation.
func GetSLSBridgedHideSpacesOperationClass() SLSBridgedHideSpacesOperationClass {
	return getSLSBridgedHideSpacesOperationClass()
}

type SLSBridgedHideSpacesOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedHideSpacesOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedHideSpacesOperationClass) Alloc() SLSBridgedHideSpacesOperation {
	rv := objc.SendIfResponds[SLSBridgedHideSpacesOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedHideSpacesOperation.Spaces]
//   - [SLSBridgedHideSpacesOperation.InitWithSpaces]
type SLSBridgedHideSpacesOperation struct {
	SLSAsynchronousBridgedWindowManagementOperation
}

// SLSBridgedHideSpacesOperationFromID constructs a [SLSBridgedHideSpacesOperation] from an objc.ID.
func SLSBridgedHideSpacesOperationFromID(id objc.ID) SLSBridgedHideSpacesOperation {
	return SLSBridgedHideSpacesOperation{SLSAsynchronousBridgedWindowManagementOperation: SLSAsynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedHideSpacesOperation implements ISLSBridgedHideSpacesOperation.
var _ ISLSBridgedHideSpacesOperation = SLSBridgedHideSpacesOperation{}

// An interface definition for the [SLSBridgedHideSpacesOperation] class.
//
// # Methods
//
//   - [ISLSBridgedHideSpacesOperation.Spaces]
//   - [ISLSBridgedHideSpacesOperation.InitWithSpaces]
type ISLSBridgedHideSpacesOperation interface {
	ISLSAsynchronousBridgedWindowManagementOperation

	// Topic: Methods

	Spaces() foundation.INSArray
	InitWithSpaces(spaces objectivec.IObject) SLSBridgedHideSpacesOperation
}

// Init initializes the instance.
func (s SLSBridgedHideSpacesOperation) Init() SLSBridgedHideSpacesOperation {
	rv := objc.SendIfResponds[SLSBridgedHideSpacesOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedHideSpacesOperation) Autorelease() SLSBridgedHideSpacesOperation {
	rv := objc.SendIfResponds[SLSBridgedHideSpacesOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedHideSpacesOperation creates a new SLSBridgedHideSpacesOperation instance.
func NewSLSBridgedHideSpacesOperation() SLSBridgedHideSpacesOperation {
	class := getSLSBridgedHideSpacesOperationClass()
	rv := objc.SendIfResponds[SLSBridgedHideSpacesOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSLSBridgedHideSpacesOperationWithCoder(coder objectivec.IObject) SLSBridgedHideSpacesOperation {
	instance := getSLSBridgedHideSpacesOperationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedHideSpacesOperationFromID(rv)
}

func NewSLSBridgedHideSpacesOperationWithSpaces(spaces objectivec.IObject) SLSBridgedHideSpacesOperation {
	instance := getSLSBridgedHideSpacesOperationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithSpaces:"), spaces)
	return SLSBridgedHideSpacesOperationFromID(rv)
}

func (s SLSBridgedHideSpacesOperation) InitWithSpaces(spaces objectivec.IObject) SLSBridgedHideSpacesOperation {
	rv := objc.SendIfResponds[SLSBridgedHideSpacesOperation](s.ID, objc.Sel("initWithSpaces:"), spaces)
	return rv
}

func (s SLSBridgedHideSpacesOperation) Spaces() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("spaces"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
