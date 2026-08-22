// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedWillSwitchSpacesOperation] class.
var (
	_SLSBridgedWillSwitchSpacesOperationClass     SLSBridgedWillSwitchSpacesOperationClass
	_SLSBridgedWillSwitchSpacesOperationClassOnce sync.Once
)

func getSLSBridgedWillSwitchSpacesOperationClass() SLSBridgedWillSwitchSpacesOperationClass {
	_SLSBridgedWillSwitchSpacesOperationClassOnce.Do(func() {
		_SLSBridgedWillSwitchSpacesOperationClass = SLSBridgedWillSwitchSpacesOperationClass{class: objc.GetClass("SLSBridgedWillSwitchSpacesOperation")}
	})
	return _SLSBridgedWillSwitchSpacesOperationClass
}

// GetSLSBridgedWillSwitchSpacesOperationClass returns the class object for SLSBridgedWillSwitchSpacesOperation.
func GetSLSBridgedWillSwitchSpacesOperationClass() SLSBridgedWillSwitchSpacesOperationClass {
	return getSLSBridgedWillSwitchSpacesOperationClass()
}

type SLSBridgedWillSwitchSpacesOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedWillSwitchSpacesOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedWillSwitchSpacesOperationClass) Alloc() SLSBridgedWillSwitchSpacesOperation {
	rv := objc.SendIfResponds[SLSBridgedWillSwitchSpacesOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedWillSwitchSpacesOperation.Spaces]
//   - [SLSBridgedWillSwitchSpacesOperation.InitWithSpaces]
type SLSBridgedWillSwitchSpacesOperation struct {
	SLSAsynchronousBridgedWindowManagementOperation
}

// SLSBridgedWillSwitchSpacesOperationFromID constructs a [SLSBridgedWillSwitchSpacesOperation] from an objc.ID.
func SLSBridgedWillSwitchSpacesOperationFromID(id objc.ID) SLSBridgedWillSwitchSpacesOperation {
	return SLSBridgedWillSwitchSpacesOperation{SLSAsynchronousBridgedWindowManagementOperation: SLSAsynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedWillSwitchSpacesOperation implements ISLSBridgedWillSwitchSpacesOperation.
var _ ISLSBridgedWillSwitchSpacesOperation = SLSBridgedWillSwitchSpacesOperation{}

// An interface definition for the [SLSBridgedWillSwitchSpacesOperation] class.
//
// # Methods
//
//   - [ISLSBridgedWillSwitchSpacesOperation.Spaces]
//   - [ISLSBridgedWillSwitchSpacesOperation.InitWithSpaces]
type ISLSBridgedWillSwitchSpacesOperation interface {
	ISLSAsynchronousBridgedWindowManagementOperation

	// Topic: Methods

	Spaces() foundation.INSArray
	InitWithSpaces(spaces objectivec.IObject) SLSBridgedWillSwitchSpacesOperation
}

// Init initializes the instance.
func (s SLSBridgedWillSwitchSpacesOperation) Init() SLSBridgedWillSwitchSpacesOperation {
	rv := objc.SendIfResponds[SLSBridgedWillSwitchSpacesOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedWillSwitchSpacesOperation) Autorelease() SLSBridgedWillSwitchSpacesOperation {
	rv := objc.SendIfResponds[SLSBridgedWillSwitchSpacesOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedWillSwitchSpacesOperation creates a new SLSBridgedWillSwitchSpacesOperation instance.
func NewSLSBridgedWillSwitchSpacesOperation() SLSBridgedWillSwitchSpacesOperation {
	class := getSLSBridgedWillSwitchSpacesOperationClass()
	rv := objc.SendIfResponds[SLSBridgedWillSwitchSpacesOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSLSBridgedWillSwitchSpacesOperationWithCoder(coder objectivec.IObject) SLSBridgedWillSwitchSpacesOperation {
	instance := getSLSBridgedWillSwitchSpacesOperationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedWillSwitchSpacesOperationFromID(rv)
}

func NewSLSBridgedWillSwitchSpacesOperationWithSpaces(spaces objectivec.IObject) SLSBridgedWillSwitchSpacesOperation {
	instance := getSLSBridgedWillSwitchSpacesOperationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithSpaces:"), spaces)
	return SLSBridgedWillSwitchSpacesOperationFromID(rv)
}

func (s SLSBridgedWillSwitchSpacesOperation) InitWithSpaces(spaces objectivec.IObject) SLSBridgedWillSwitchSpacesOperation {
	rv := objc.SendIfResponds[SLSBridgedWillSwitchSpacesOperation](s.ID, objc.Sel("initWithSpaces:"), spaces)
	return rv
}

func (s SLSBridgedWillSwitchSpacesOperation) Spaces() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("spaces"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
