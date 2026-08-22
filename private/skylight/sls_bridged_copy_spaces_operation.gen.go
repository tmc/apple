// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedCopySpacesOperation] class.
var (
	_SLSBridgedCopySpacesOperationClass     SLSBridgedCopySpacesOperationClass
	_SLSBridgedCopySpacesOperationClassOnce sync.Once
)

func getSLSBridgedCopySpacesOperationClass() SLSBridgedCopySpacesOperationClass {
	_SLSBridgedCopySpacesOperationClassOnce.Do(func() {
		_SLSBridgedCopySpacesOperationClass = SLSBridgedCopySpacesOperationClass{class: objc.GetClass("SLSBridgedCopySpacesOperation")}
	})
	return _SLSBridgedCopySpacesOperationClass
}

// GetSLSBridgedCopySpacesOperationClass returns the class object for SLSBridgedCopySpacesOperation.
func GetSLSBridgedCopySpacesOperationClass() SLSBridgedCopySpacesOperationClass {
	return getSLSBridgedCopySpacesOperationClass()
}

type SLSBridgedCopySpacesOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedCopySpacesOperationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedCopySpacesOperationClass) Alloc() SLSBridgedCopySpacesOperation {
	rv := objc.SendIfResponds[SLSBridgedCopySpacesOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedCopySpacesOperation.MakeResultWithNumbers]
//   - [SLSBridgedCopySpacesOperation.Options]
//   - [SLSBridgedCopySpacesOperation.InitWithOptions]
type SLSBridgedCopySpacesOperation struct {
	SLSSynchronousBridgedWindowManagementOperation
}

// SLSBridgedCopySpacesOperationFromID constructs a [SLSBridgedCopySpacesOperation] from an objc.ID.
func SLSBridgedCopySpacesOperationFromID(id objc.ID) SLSBridgedCopySpacesOperation {
	return SLSBridgedCopySpacesOperation{SLSSynchronousBridgedWindowManagementOperation: SLSSynchronousBridgedWindowManagementOperationFromID(id)}
}

// Ensure SLSBridgedCopySpacesOperation implements ISLSBridgedCopySpacesOperation.
var _ ISLSBridgedCopySpacesOperation = SLSBridgedCopySpacesOperation{}

// An interface definition for the [SLSBridgedCopySpacesOperation] class.
//
// # Methods
//
//   - [ISLSBridgedCopySpacesOperation.MakeResultWithNumbers]
//   - [ISLSBridgedCopySpacesOperation.Options]
//   - [ISLSBridgedCopySpacesOperation.InitWithOptions]
type ISLSBridgedCopySpacesOperation interface {
	ISLSSynchronousBridgedWindowManagementOperation

	// Topic: Methods

	MakeResultWithNumbers(numbers objectivec.IObject) objectivec.IObject
	Options() uint32
	InitWithOptions(options uint32) SLSBridgedCopySpacesOperation
}

// Init initializes the instance.
func (s SLSBridgedCopySpacesOperation) Init() SLSBridgedCopySpacesOperation {
	rv := objc.SendIfResponds[SLSBridgedCopySpacesOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedCopySpacesOperation) Autorelease() SLSBridgedCopySpacesOperation {
	rv := objc.SendIfResponds[SLSBridgedCopySpacesOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedCopySpacesOperation creates a new SLSBridgedCopySpacesOperation instance.
func NewSLSBridgedCopySpacesOperation() SLSBridgedCopySpacesOperation {
	class := getSLSBridgedCopySpacesOperationClass()
	rv := objc.SendIfResponds[SLSBridgedCopySpacesOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSLSBridgedCopySpacesOperationWithCoder(coder objectivec.IObject) SLSBridgedCopySpacesOperation {
	instance := getSLSBridgedCopySpacesOperationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedCopySpacesOperationFromID(rv)
}

func NewSLSBridgedCopySpacesOperationWithOptions(options uint32) SLSBridgedCopySpacesOperation {
	instance := getSLSBridgedCopySpacesOperationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithOptions:"), options)
	return SLSBridgedCopySpacesOperationFromID(rv)
}

func (s SLSBridgedCopySpacesOperation) MakeResultWithNumbers(numbers objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("makeResultWithNumbers:"), numbers)
	return objectivec.Object{ID: rv}
}
func (s SLSBridgedCopySpacesOperation) InitWithOptions(options uint32) SLSBridgedCopySpacesOperation {
	rv := objc.SendIfResponds[SLSBridgedCopySpacesOperation](s.ID, objc.Sel("initWithOptions:"), options)
	return rv
}

func (s SLSBridgedCopySpacesOperation) Options() uint32 {
	rv := objc.SendIfResponds[uint32](s.ID, objc.Sel("options"))
	return rv
}
