// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedShowSpacesOperation] class.
var (
	_SLSBridgedShowSpacesOperationClass     SLSBridgedShowSpacesOperationClass
	_SLSBridgedShowSpacesOperationClassOnce sync.Once
)

func getSLSBridgedShowSpacesOperationClass() SLSBridgedShowSpacesOperationClass {
	_SLSBridgedShowSpacesOperationClassOnce.Do(func() {
		_SLSBridgedShowSpacesOperationClass = SLSBridgedShowSpacesOperationClass{class: objc.GetClass("SLSBridgedShowSpacesOperation")}
	})
	return _SLSBridgedShowSpacesOperationClass
}

// GetSLSBridgedShowSpacesOperationClass returns the class object for SLSBridgedShowSpacesOperation.
func GetSLSBridgedShowSpacesOperationClass() SLSBridgedShowSpacesOperationClass {
	return getSLSBridgedShowSpacesOperationClass()
}

type SLSBridgedShowSpacesOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedShowSpacesOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedShowSpacesOperationClass) Alloc() SLSBridgedShowSpacesOperation {
	rv := objc.Send[SLSBridgedShowSpacesOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedShowSpacesOperation.Spaces]
//   - [SLSBridgedShowSpacesOperation.InitWithSpaces]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedShowSpacesOperation
type SLSBridgedShowSpacesOperation struct {
	SLSAsynchronousBridgedWindowManagementOperation
}

// SLSBridgedShowSpacesOperationFromID constructs a [SLSBridgedShowSpacesOperation] from an objc.ID.
func SLSBridgedShowSpacesOperationFromID(id objc.ID) SLSBridgedShowSpacesOperation {
	return SLSBridgedShowSpacesOperation{SLSAsynchronousBridgedWindowManagementOperation: SLSAsynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedShowSpacesOperation implements ISLSBridgedShowSpacesOperation.
var _ ISLSBridgedShowSpacesOperation = SLSBridgedShowSpacesOperation{}

// An interface definition for the [SLSBridgedShowSpacesOperation] class.
//
// # Methods
//
//   - [ISLSBridgedShowSpacesOperation.Spaces]
//   - [ISLSBridgedShowSpacesOperation.InitWithSpaces]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedShowSpacesOperation
type ISLSBridgedShowSpacesOperation interface {
	ISLSAsynchronousBridgedWindowManagementOperation

	// Topic: Methods

	Spaces() foundation.INSArray
	InitWithSpaces(spaces objectivec.IObject) SLSBridgedShowSpacesOperation
}

// Init initializes the instance.
func (s SLSBridgedShowSpacesOperation) Init() SLSBridgedShowSpacesOperation {
	rv := objc.Send[SLSBridgedShowSpacesOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedShowSpacesOperation) Autorelease() SLSBridgedShowSpacesOperation {
	rv := objc.Send[SLSBridgedShowSpacesOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedShowSpacesOperation creates a new SLSBridgedShowSpacesOperation instance.
func NewSLSBridgedShowSpacesOperation() SLSBridgedShowSpacesOperation {
	class := getSLSBridgedShowSpacesOperationClass()
	rv := objc.Send[SLSBridgedShowSpacesOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedShowSpacesOperation/initWithCoder:
func NewSLSBridgedShowSpacesOperationWithCoder(coder objectivec.IObject) SLSBridgedShowSpacesOperation {
	instance := getSLSBridgedShowSpacesOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedShowSpacesOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedShowSpacesOperation/initWithSpaces:
func NewSLSBridgedShowSpacesOperationWithSpaces(spaces objectivec.IObject) SLSBridgedShowSpacesOperation {
	instance := getSLSBridgedShowSpacesOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpaces:"), spaces)
	return SLSBridgedShowSpacesOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedShowSpacesOperation/initWithSpaces:
func (s SLSBridgedShowSpacesOperation) InitWithSpaces(spaces objectivec.IObject) SLSBridgedShowSpacesOperation {
	rv := objc.Send[SLSBridgedShowSpacesOperation](s.ID, objc.Sel("initWithSpaces:"), spaces)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedShowSpacesOperation/spaces
func (s SLSBridgedShowSpacesOperation) Spaces() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("spaces"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
