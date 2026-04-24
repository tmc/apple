// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedWindowManagementOperationPropertyListArrayResult] class.
var (
	_SLSBridgedWindowManagementOperationPropertyListArrayResultClass     SLSBridgedWindowManagementOperationPropertyListArrayResultClass
	_SLSBridgedWindowManagementOperationPropertyListArrayResultClassOnce sync.Once
)

func getSLSBridgedWindowManagementOperationPropertyListArrayResultClass() SLSBridgedWindowManagementOperationPropertyListArrayResultClass {
	_SLSBridgedWindowManagementOperationPropertyListArrayResultClassOnce.Do(func() {
		_SLSBridgedWindowManagementOperationPropertyListArrayResultClass = SLSBridgedWindowManagementOperationPropertyListArrayResultClass{class: objc.GetClass("SLSBridgedWindowManagementOperationPropertyListArrayResult")}
	})
	return _SLSBridgedWindowManagementOperationPropertyListArrayResultClass
}

// GetSLSBridgedWindowManagementOperationPropertyListArrayResultClass returns the class object for SLSBridgedWindowManagementOperationPropertyListArrayResult.
func GetSLSBridgedWindowManagementOperationPropertyListArrayResultClass() SLSBridgedWindowManagementOperationPropertyListArrayResultClass {
	return getSLSBridgedWindowManagementOperationPropertyListArrayResultClass()
}

type SLSBridgedWindowManagementOperationPropertyListArrayResultClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedWindowManagementOperationPropertyListArrayResultClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedWindowManagementOperationPropertyListArrayResultClass) Alloc() SLSBridgedWindowManagementOperationPropertyListArrayResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationPropertyListArrayResult](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedWindowManagementOperationPropertyListArrayResult.PropertyListArray]
//   - [SLSBridgedWindowManagementOperationPropertyListArrayResult.InitWithPropertyListArray]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationPropertyListArrayResult
type SLSBridgedWindowManagementOperationPropertyListArrayResult struct {
	SLSBridgedWindowManagementOperationResult
}

// SLSBridgedWindowManagementOperationPropertyListArrayResultFromID constructs a [SLSBridgedWindowManagementOperationPropertyListArrayResult] from an objc.ID.
func SLSBridgedWindowManagementOperationPropertyListArrayResultFromID(id objc.ID) SLSBridgedWindowManagementOperationPropertyListArrayResult {
	return SLSBridgedWindowManagementOperationPropertyListArrayResult{SLSBridgedWindowManagementOperationResult: SLSBridgedWindowManagementOperationResultFromID(id)}
}

// Ensure SLSBridgedWindowManagementOperationPropertyListArrayResult implements ISLSBridgedWindowManagementOperationPropertyListArrayResult.
var _ ISLSBridgedWindowManagementOperationPropertyListArrayResult = SLSBridgedWindowManagementOperationPropertyListArrayResult{}

// An interface definition for the [SLSBridgedWindowManagementOperationPropertyListArrayResult] class.
//
// # Methods
//
//   - [ISLSBridgedWindowManagementOperationPropertyListArrayResult.PropertyListArray]
//   - [ISLSBridgedWindowManagementOperationPropertyListArrayResult.InitWithPropertyListArray]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationPropertyListArrayResult
type ISLSBridgedWindowManagementOperationPropertyListArrayResult interface {
	ISLSBridgedWindowManagementOperationResult

	// Topic: Methods

	PropertyListArray() foundation.INSArray
	InitWithPropertyListArray(array objectivec.IObject) SLSBridgedWindowManagementOperationPropertyListArrayResult
}

// Init initializes the instance.
func (s SLSBridgedWindowManagementOperationPropertyListArrayResult) Init() SLSBridgedWindowManagementOperationPropertyListArrayResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationPropertyListArrayResult](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedWindowManagementOperationPropertyListArrayResult) Autorelease() SLSBridgedWindowManagementOperationPropertyListArrayResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationPropertyListArrayResult](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedWindowManagementOperationPropertyListArrayResult creates a new SLSBridgedWindowManagementOperationPropertyListArrayResult instance.
func NewSLSBridgedWindowManagementOperationPropertyListArrayResult() SLSBridgedWindowManagementOperationPropertyListArrayResult {
	class := getSLSBridgedWindowManagementOperationPropertyListArrayResultClass()
	rv := objc.Send[SLSBridgedWindowManagementOperationPropertyListArrayResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationPropertyListArrayResult/initWithCoder:
func NewSLSBridgedWindowManagementOperationPropertyListArrayResultWithCoder(coder objectivec.IObject) SLSBridgedWindowManagementOperationPropertyListArrayResult {
	instance := getSLSBridgedWindowManagementOperationPropertyListArrayResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedWindowManagementOperationPropertyListArrayResultFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationPropertyListArrayResult/initWithPropertyListArray:
func NewSLSBridgedWindowManagementOperationPropertyListArrayResultWithPropertyListArray(array objectivec.IObject) SLSBridgedWindowManagementOperationPropertyListArrayResult {
	instance := getSLSBridgedWindowManagementOperationPropertyListArrayResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPropertyListArray:"), array)
	return SLSBridgedWindowManagementOperationPropertyListArrayResultFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationPropertyListArrayResult/initWithPropertyListArray:
func (s SLSBridgedWindowManagementOperationPropertyListArrayResult) InitWithPropertyListArray(array objectivec.IObject) SLSBridgedWindowManagementOperationPropertyListArrayResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationPropertyListArrayResult](s.ID, objc.Sel("initWithPropertyListArray:"), array)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationPropertyListArrayResult/propertyListArray
func (s SLSBridgedWindowManagementOperationPropertyListArrayResult) PropertyListArray() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("propertyListArray"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
