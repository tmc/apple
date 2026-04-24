// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedWindowManagementOperationNumbersResult] class.
var (
	_SLSBridgedWindowManagementOperationNumbersResultClass     SLSBridgedWindowManagementOperationNumbersResultClass
	_SLSBridgedWindowManagementOperationNumbersResultClassOnce sync.Once
)

func getSLSBridgedWindowManagementOperationNumbersResultClass() SLSBridgedWindowManagementOperationNumbersResultClass {
	_SLSBridgedWindowManagementOperationNumbersResultClassOnce.Do(func() {
		_SLSBridgedWindowManagementOperationNumbersResultClass = SLSBridgedWindowManagementOperationNumbersResultClass{class: objc.GetClass("SLSBridgedWindowManagementOperationNumbersResult")}
	})
	return _SLSBridgedWindowManagementOperationNumbersResultClass
}

// GetSLSBridgedWindowManagementOperationNumbersResultClass returns the class object for SLSBridgedWindowManagementOperationNumbersResult.
func GetSLSBridgedWindowManagementOperationNumbersResultClass() SLSBridgedWindowManagementOperationNumbersResultClass {
	return getSLSBridgedWindowManagementOperationNumbersResultClass()
}

type SLSBridgedWindowManagementOperationNumbersResultClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedWindowManagementOperationNumbersResultClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedWindowManagementOperationNumbersResultClass) Alloc() SLSBridgedWindowManagementOperationNumbersResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationNumbersResult](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedWindowManagementOperationNumbersResult.Numbers]
//   - [SLSBridgedWindowManagementOperationNumbersResult.InitWithNumbers]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationNumbersResult
type SLSBridgedWindowManagementOperationNumbersResult struct {
	SLSBridgedWindowManagementOperationResult
}

// SLSBridgedWindowManagementOperationNumbersResultFromID constructs a [SLSBridgedWindowManagementOperationNumbersResult] from an objc.ID.
func SLSBridgedWindowManagementOperationNumbersResultFromID(id objc.ID) SLSBridgedWindowManagementOperationNumbersResult {
	return SLSBridgedWindowManagementOperationNumbersResult{SLSBridgedWindowManagementOperationResult: SLSBridgedWindowManagementOperationResultFromID(id)}
}

// Ensure SLSBridgedWindowManagementOperationNumbersResult implements ISLSBridgedWindowManagementOperationNumbersResult.
var _ ISLSBridgedWindowManagementOperationNumbersResult = SLSBridgedWindowManagementOperationNumbersResult{}

// An interface definition for the [SLSBridgedWindowManagementOperationNumbersResult] class.
//
// # Methods
//
//   - [ISLSBridgedWindowManagementOperationNumbersResult.Numbers]
//   - [ISLSBridgedWindowManagementOperationNumbersResult.InitWithNumbers]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationNumbersResult
type ISLSBridgedWindowManagementOperationNumbersResult interface {
	ISLSBridgedWindowManagementOperationResult

	// Topic: Methods

	Numbers() foundation.INSArray
	InitWithNumbers(numbers objectivec.IObject) SLSBridgedWindowManagementOperationNumbersResult
}

// Init initializes the instance.
func (s SLSBridgedWindowManagementOperationNumbersResult) Init() SLSBridgedWindowManagementOperationNumbersResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationNumbersResult](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedWindowManagementOperationNumbersResult) Autorelease() SLSBridgedWindowManagementOperationNumbersResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationNumbersResult](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedWindowManagementOperationNumbersResult creates a new SLSBridgedWindowManagementOperationNumbersResult instance.
func NewSLSBridgedWindowManagementOperationNumbersResult() SLSBridgedWindowManagementOperationNumbersResult {
	class := getSLSBridgedWindowManagementOperationNumbersResultClass()
	rv := objc.Send[SLSBridgedWindowManagementOperationNumbersResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationNumbersResult/initWithCoder:
func NewSLSBridgedWindowManagementOperationNumbersResultWithCoder(coder objectivec.IObject) SLSBridgedWindowManagementOperationNumbersResult {
	instance := getSLSBridgedWindowManagementOperationNumbersResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedWindowManagementOperationNumbersResultFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationNumbersResult/initWithNumbers:
func NewSLSBridgedWindowManagementOperationNumbersResultWithNumbers(numbers objectivec.IObject) SLSBridgedWindowManagementOperationNumbersResult {
	instance := getSLSBridgedWindowManagementOperationNumbersResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithNumbers:"), numbers)
	return SLSBridgedWindowManagementOperationNumbersResultFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationNumbersResult/initWithNumbers:
func (s SLSBridgedWindowManagementOperationNumbersResult) InitWithNumbers(numbers objectivec.IObject) SLSBridgedWindowManagementOperationNumbersResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationNumbersResult](s.ID, objc.Sel("initWithNumbers:"), numbers)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationNumbersResult/numbers
func (s SLSBridgedWindowManagementOperationNumbersResult) Numbers() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("numbers"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
