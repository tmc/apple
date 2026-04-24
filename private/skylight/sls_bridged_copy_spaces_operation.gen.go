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
	rv := objc.Send[SLSBridgedCopySpacesOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedCopySpacesOperation.MakeResultWithNumbers]
//   - [SLSBridgedCopySpacesOperation.Options]
//   - [SLSBridgedCopySpacesOperation.InitWithOptions]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopySpacesOperation
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
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopySpacesOperation
type ISLSBridgedCopySpacesOperation interface {
	ISLSSynchronousBridgedWindowManagementOperation

	// Topic: Methods

	MakeResultWithNumbers(numbers objectivec.IObject) objectivec.IObject
	Options() uint32
	InitWithOptions(options uint32) SLSBridgedCopySpacesOperation
}

// Init initializes the instance.
func (s SLSBridgedCopySpacesOperation) Init() SLSBridgedCopySpacesOperation {
	rv := objc.Send[SLSBridgedCopySpacesOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedCopySpacesOperation) Autorelease() SLSBridgedCopySpacesOperation {
	rv := objc.Send[SLSBridgedCopySpacesOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedCopySpacesOperation creates a new SLSBridgedCopySpacesOperation instance.
func NewSLSBridgedCopySpacesOperation() SLSBridgedCopySpacesOperation {
	class := getSLSBridgedCopySpacesOperationClass()
	rv := objc.Send[SLSBridgedCopySpacesOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopySpacesOperation/initWithCoder:
func NewSLSBridgedCopySpacesOperationWithCoder(coder objectivec.IObject) SLSBridgedCopySpacesOperation {
	instance := getSLSBridgedCopySpacesOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedCopySpacesOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopySpacesOperation/initWithOptions:
func NewSLSBridgedCopySpacesOperationWithOptions(options uint32) SLSBridgedCopySpacesOperation {
	instance := getSLSBridgedCopySpacesOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithOptions:"), options)
	return SLSBridgedCopySpacesOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopySpacesOperation/makeResultWithNumbers:
func (s SLSBridgedCopySpacesOperation) MakeResultWithNumbers(numbers objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("makeResultWithNumbers:"), numbers)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopySpacesOperation/initWithOptions:
func (s SLSBridgedCopySpacesOperation) InitWithOptions(options uint32) SLSBridgedCopySpacesOperation {
	rv := objc.Send[SLSBridgedCopySpacesOperation](s.ID, objc.Sel("initWithOptions:"), options)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedCopySpacesOperation/options
func (s SLSBridgedCopySpacesOperation) Options() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("options"))
	return rv
}
