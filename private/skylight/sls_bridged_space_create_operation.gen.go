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
	rv := objc.Send[SLSBridgedSpaceCreateOperation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedSpaceCreateOperation.MakeResultWithSpaceID]
//   - [SLSBridgedSpaceCreateOperation.Options]
//   - [SLSBridgedSpaceCreateOperation.Values]
//   - [SLSBridgedSpaceCreateOperation.InitWithOptionsValues]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCreateOperation
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
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCreateOperation
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
	rv := objc.Send[SLSBridgedSpaceCreateOperation](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedSpaceCreateOperation) Autorelease() SLSBridgedSpaceCreateOperation {
	rv := objc.Send[SLSBridgedSpaceCreateOperation](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedSpaceCreateOperation creates a new SLSBridgedSpaceCreateOperation instance.
func NewSLSBridgedSpaceCreateOperation() SLSBridgedSpaceCreateOperation {
	class := getSLSBridgedSpaceCreateOperationClass()
	rv := objc.Send[SLSBridgedSpaceCreateOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCreateOperation/initWithCoder:
func NewSLSBridgedSpaceCreateOperationWithCoder(coder objectivec.IObject) SLSBridgedSpaceCreateOperation {
	instance := getSLSBridgedSpaceCreateOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedSpaceCreateOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCreateOperation/initWithOptions:values:
func NewSLSBridgedSpaceCreateOperationWithOptionsValues(options uint32, values objectivec.IObject) SLSBridgedSpaceCreateOperation {
	instance := getSLSBridgedSpaceCreateOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithOptions:values:"), options, values)
	return SLSBridgedSpaceCreateOperationFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCreateOperation/makeResultWithSpaceID:
func (s SLSBridgedSpaceCreateOperation) MakeResultWithSpaceID(id uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("makeResultWithSpaceID:"), id)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCreateOperation/initWithOptions:values:
func (s SLSBridgedSpaceCreateOperation) InitWithOptionsValues(options uint32, values objectivec.IObject) SLSBridgedSpaceCreateOperation {
	rv := objc.Send[SLSBridgedSpaceCreateOperation](s.ID, objc.Sel("initWithOptions:values:"), options, values)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCreateOperation/options
func (s SLSBridgedSpaceCreateOperation) Options() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("options"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedSpaceCreateOperation/values
func (s SLSBridgedSpaceCreateOperation) Values() foundation.INSDictionary {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("values"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
