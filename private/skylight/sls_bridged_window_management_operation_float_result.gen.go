// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedWindowManagementOperationFloatResult] class.
var (
	_SLSBridgedWindowManagementOperationFloatResultClass     SLSBridgedWindowManagementOperationFloatResultClass
	_SLSBridgedWindowManagementOperationFloatResultClassOnce sync.Once
)

func getSLSBridgedWindowManagementOperationFloatResultClass() SLSBridgedWindowManagementOperationFloatResultClass {
	_SLSBridgedWindowManagementOperationFloatResultClassOnce.Do(func() {
		_SLSBridgedWindowManagementOperationFloatResultClass = SLSBridgedWindowManagementOperationFloatResultClass{class: objc.GetClass("SLSBridgedWindowManagementOperationFloatResult")}
	})
	return _SLSBridgedWindowManagementOperationFloatResultClass
}

// GetSLSBridgedWindowManagementOperationFloatResultClass returns the class object for SLSBridgedWindowManagementOperationFloatResult.
func GetSLSBridgedWindowManagementOperationFloatResultClass() SLSBridgedWindowManagementOperationFloatResultClass {
	return getSLSBridgedWindowManagementOperationFloatResultClass()
}

type SLSBridgedWindowManagementOperationFloatResultClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedWindowManagementOperationFloatResultClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedWindowManagementOperationFloatResultClass) Alloc() SLSBridgedWindowManagementOperationFloatResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationFloatResult](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedWindowManagementOperationFloatResult.FloatValue]
//   - [SLSBridgedWindowManagementOperationFloatResult.InitWithFloatValue]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationFloatResult
type SLSBridgedWindowManagementOperationFloatResult struct {
	SLSBridgedWindowManagementOperationResult
}

// SLSBridgedWindowManagementOperationFloatResultFromID constructs a [SLSBridgedWindowManagementOperationFloatResult] from an objc.ID.
func SLSBridgedWindowManagementOperationFloatResultFromID(id objc.ID) SLSBridgedWindowManagementOperationFloatResult {
	return SLSBridgedWindowManagementOperationFloatResult{SLSBridgedWindowManagementOperationResult: SLSBridgedWindowManagementOperationResultFromID(id)}
}

// Ensure SLSBridgedWindowManagementOperationFloatResult implements ISLSBridgedWindowManagementOperationFloatResult.
var _ ISLSBridgedWindowManagementOperationFloatResult = SLSBridgedWindowManagementOperationFloatResult{}

// An interface definition for the [SLSBridgedWindowManagementOperationFloatResult] class.
//
// # Methods
//
//   - [ISLSBridgedWindowManagementOperationFloatResult.FloatValue]
//   - [ISLSBridgedWindowManagementOperationFloatResult.InitWithFloatValue]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationFloatResult
type ISLSBridgedWindowManagementOperationFloatResult interface {
	ISLSBridgedWindowManagementOperationResult

	// Topic: Methods

	FloatValue() float32
	InitWithFloatValue(value float32) SLSBridgedWindowManagementOperationFloatResult
}

// Init initializes the instance.
func (s SLSBridgedWindowManagementOperationFloatResult) Init() SLSBridgedWindowManagementOperationFloatResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationFloatResult](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedWindowManagementOperationFloatResult) Autorelease() SLSBridgedWindowManagementOperationFloatResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationFloatResult](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedWindowManagementOperationFloatResult creates a new SLSBridgedWindowManagementOperationFloatResult instance.
func NewSLSBridgedWindowManagementOperationFloatResult() SLSBridgedWindowManagementOperationFloatResult {
	class := getSLSBridgedWindowManagementOperationFloatResultClass()
	rv := objc.Send[SLSBridgedWindowManagementOperationFloatResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationFloatResult/initWithCoder:
func NewSLSBridgedWindowManagementOperationFloatResultWithCoder(coder objectivec.IObject) SLSBridgedWindowManagementOperationFloatResult {
	instance := getSLSBridgedWindowManagementOperationFloatResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedWindowManagementOperationFloatResultFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationFloatResult/initWithFloatValue:
func NewSLSBridgedWindowManagementOperationFloatResultWithFloatValue(value float32) SLSBridgedWindowManagementOperationFloatResult {
	instance := getSLSBridgedWindowManagementOperationFloatResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFloatValue:"), value)
	return SLSBridgedWindowManagementOperationFloatResultFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationFloatResult/initWithFloatValue:
func (s SLSBridgedWindowManagementOperationFloatResult) InitWithFloatValue(value float32) SLSBridgedWindowManagementOperationFloatResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationFloatResult](s.ID, objc.Sel("initWithFloatValue:"), value)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationFloatResult/floatValue
func (s SLSBridgedWindowManagementOperationFloatResult) FloatValue() float32 {
	rv := objc.Send[float32](s.ID, objc.Sel("floatValue"))
	return rv
}
