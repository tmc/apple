// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedWindowManagementOperationStringResult] class.
var (
	_SLSBridgedWindowManagementOperationStringResultClass     SLSBridgedWindowManagementOperationStringResultClass
	_SLSBridgedWindowManagementOperationStringResultClassOnce sync.Once
)

func getSLSBridgedWindowManagementOperationStringResultClass() SLSBridgedWindowManagementOperationStringResultClass {
	_SLSBridgedWindowManagementOperationStringResultClassOnce.Do(func() {
		_SLSBridgedWindowManagementOperationStringResultClass = SLSBridgedWindowManagementOperationStringResultClass{class: objc.GetClass("SLSBridgedWindowManagementOperationStringResult")}
	})
	return _SLSBridgedWindowManagementOperationStringResultClass
}

// GetSLSBridgedWindowManagementOperationStringResultClass returns the class object for SLSBridgedWindowManagementOperationStringResult.
func GetSLSBridgedWindowManagementOperationStringResultClass() SLSBridgedWindowManagementOperationStringResultClass {
	return getSLSBridgedWindowManagementOperationStringResultClass()
}

type SLSBridgedWindowManagementOperationStringResultClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedWindowManagementOperationStringResultClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedWindowManagementOperationStringResultClass) Alloc() SLSBridgedWindowManagementOperationStringResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationStringResult](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedWindowManagementOperationStringResult.InitWithString]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationStringResult
type SLSBridgedWindowManagementOperationStringResult struct {
	SLSBridgedWindowManagementOperationResult
}

// SLSBridgedWindowManagementOperationStringResultFromID constructs a [SLSBridgedWindowManagementOperationStringResult] from an objc.ID.
func SLSBridgedWindowManagementOperationStringResultFromID(id objc.ID) SLSBridgedWindowManagementOperationStringResult {
	return SLSBridgedWindowManagementOperationStringResult{SLSBridgedWindowManagementOperationResult: SLSBridgedWindowManagementOperationResultFromID(id)}
}

// Ensure SLSBridgedWindowManagementOperationStringResult implements ISLSBridgedWindowManagementOperationStringResult.
var _ ISLSBridgedWindowManagementOperationStringResult = SLSBridgedWindowManagementOperationStringResult{}

// An interface definition for the [SLSBridgedWindowManagementOperationStringResult] class.
//
// # Methods
//
//   - [ISLSBridgedWindowManagementOperationStringResult.InitWithString]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationStringResult
type ISLSBridgedWindowManagementOperationStringResult interface {
	ISLSBridgedWindowManagementOperationResult

	// Topic: Methods

	InitWithString(string_ objectivec.IObject) SLSBridgedWindowManagementOperationStringResult
}

// Init initializes the instance.
func (s SLSBridgedWindowManagementOperationStringResult) Init() SLSBridgedWindowManagementOperationStringResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationStringResult](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedWindowManagementOperationStringResult) Autorelease() SLSBridgedWindowManagementOperationStringResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationStringResult](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedWindowManagementOperationStringResult creates a new SLSBridgedWindowManagementOperationStringResult instance.
func NewSLSBridgedWindowManagementOperationStringResult() SLSBridgedWindowManagementOperationStringResult {
	class := getSLSBridgedWindowManagementOperationStringResultClass()
	rv := objc.Send[SLSBridgedWindowManagementOperationStringResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationStringResult/initWithCoder:
func NewSLSBridgedWindowManagementOperationStringResultWithCoder(coder objectivec.IObject) SLSBridgedWindowManagementOperationStringResult {
	instance := getSLSBridgedWindowManagementOperationStringResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedWindowManagementOperationStringResultFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationStringResult/initWithString:
func NewSLSBridgedWindowManagementOperationStringResultWithString(string_ objectivec.IObject) SLSBridgedWindowManagementOperationStringResult {
	instance := getSLSBridgedWindowManagementOperationStringResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithString:"), string_)
	return SLSBridgedWindowManagementOperationStringResultFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationStringResult/initWithString:
func (s SLSBridgedWindowManagementOperationStringResult) InitWithString(string_ objectivec.IObject) SLSBridgedWindowManagementOperationStringResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationStringResult](s.ID, objc.Sel("initWithString:"), string_)
	return rv
}
