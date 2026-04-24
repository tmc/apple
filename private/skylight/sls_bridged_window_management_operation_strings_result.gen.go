// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedWindowManagementOperationStringsResult] class.
var (
	_SLSBridgedWindowManagementOperationStringsResultClass     SLSBridgedWindowManagementOperationStringsResultClass
	_SLSBridgedWindowManagementOperationStringsResultClassOnce sync.Once
)

func getSLSBridgedWindowManagementOperationStringsResultClass() SLSBridgedWindowManagementOperationStringsResultClass {
	_SLSBridgedWindowManagementOperationStringsResultClassOnce.Do(func() {
		_SLSBridgedWindowManagementOperationStringsResultClass = SLSBridgedWindowManagementOperationStringsResultClass{class: objc.GetClass("SLSBridgedWindowManagementOperationStringsResult")}
	})
	return _SLSBridgedWindowManagementOperationStringsResultClass
}

// GetSLSBridgedWindowManagementOperationStringsResultClass returns the class object for SLSBridgedWindowManagementOperationStringsResult.
func GetSLSBridgedWindowManagementOperationStringsResultClass() SLSBridgedWindowManagementOperationStringsResultClass {
	return getSLSBridgedWindowManagementOperationStringsResultClass()
}

type SLSBridgedWindowManagementOperationStringsResultClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedWindowManagementOperationStringsResultClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedWindowManagementOperationStringsResultClass) Alloc() SLSBridgedWindowManagementOperationStringsResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationStringsResult](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedWindowManagementOperationStringsResult.Strings]
//   - [SLSBridgedWindowManagementOperationStringsResult.InitWithStrings]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationStringsResult
type SLSBridgedWindowManagementOperationStringsResult struct {
	SLSBridgedWindowManagementOperationResult
}

// SLSBridgedWindowManagementOperationStringsResultFromID constructs a [SLSBridgedWindowManagementOperationStringsResult] from an objc.ID.
func SLSBridgedWindowManagementOperationStringsResultFromID(id objc.ID) SLSBridgedWindowManagementOperationStringsResult {
	return SLSBridgedWindowManagementOperationStringsResult{SLSBridgedWindowManagementOperationResult: SLSBridgedWindowManagementOperationResultFromID(id)}
}

// Ensure SLSBridgedWindowManagementOperationStringsResult implements ISLSBridgedWindowManagementOperationStringsResult.
var _ ISLSBridgedWindowManagementOperationStringsResult = SLSBridgedWindowManagementOperationStringsResult{}

// An interface definition for the [SLSBridgedWindowManagementOperationStringsResult] class.
//
// # Methods
//
//   - [ISLSBridgedWindowManagementOperationStringsResult.Strings]
//   - [ISLSBridgedWindowManagementOperationStringsResult.InitWithStrings]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationStringsResult
type ISLSBridgedWindowManagementOperationStringsResult interface {
	ISLSBridgedWindowManagementOperationResult

	// Topic: Methods

	Strings() foundation.INSArray
	InitWithStrings(strings objectivec.IObject) SLSBridgedWindowManagementOperationStringsResult
}

// Init initializes the instance.
func (s SLSBridgedWindowManagementOperationStringsResult) Init() SLSBridgedWindowManagementOperationStringsResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationStringsResult](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedWindowManagementOperationStringsResult) Autorelease() SLSBridgedWindowManagementOperationStringsResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationStringsResult](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedWindowManagementOperationStringsResult creates a new SLSBridgedWindowManagementOperationStringsResult instance.
func NewSLSBridgedWindowManagementOperationStringsResult() SLSBridgedWindowManagementOperationStringsResult {
	class := getSLSBridgedWindowManagementOperationStringsResultClass()
	rv := objc.Send[SLSBridgedWindowManagementOperationStringsResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationStringsResult/initWithCoder:
func NewSLSBridgedWindowManagementOperationStringsResultWithCoder(coder objectivec.IObject) SLSBridgedWindowManagementOperationStringsResult {
	instance := getSLSBridgedWindowManagementOperationStringsResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedWindowManagementOperationStringsResultFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationStringsResult/initWithStrings:
func NewSLSBridgedWindowManagementOperationStringsResultWithStrings(strings objectivec.IObject) SLSBridgedWindowManagementOperationStringsResult {
	instance := getSLSBridgedWindowManagementOperationStringsResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithStrings:"), strings)
	return SLSBridgedWindowManagementOperationStringsResultFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationStringsResult/initWithStrings:
func (s SLSBridgedWindowManagementOperationStringsResult) InitWithStrings(strings objectivec.IObject) SLSBridgedWindowManagementOperationStringsResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationStringsResult](s.ID, objc.Sel("initWithStrings:"), strings)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationStringsResult/strings
func (s SLSBridgedWindowManagementOperationStringsResult) Strings() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("strings"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
