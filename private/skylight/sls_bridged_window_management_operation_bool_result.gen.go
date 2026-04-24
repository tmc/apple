// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedWindowManagementOperationBoolResult] class.
var (
	_SLSBridgedWindowManagementOperationBoolResultClass     SLSBridgedWindowManagementOperationBoolResultClass
	_SLSBridgedWindowManagementOperationBoolResultClassOnce sync.Once
)

func getSLSBridgedWindowManagementOperationBoolResultClass() SLSBridgedWindowManagementOperationBoolResultClass {
	_SLSBridgedWindowManagementOperationBoolResultClassOnce.Do(func() {
		_SLSBridgedWindowManagementOperationBoolResultClass = SLSBridgedWindowManagementOperationBoolResultClass{class: objc.GetClass("SLSBridgedWindowManagementOperationBoolResult")}
	})
	return _SLSBridgedWindowManagementOperationBoolResultClass
}

// GetSLSBridgedWindowManagementOperationBoolResultClass returns the class object for SLSBridgedWindowManagementOperationBoolResult.
func GetSLSBridgedWindowManagementOperationBoolResultClass() SLSBridgedWindowManagementOperationBoolResultClass {
	return getSLSBridgedWindowManagementOperationBoolResultClass()
}

type SLSBridgedWindowManagementOperationBoolResultClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedWindowManagementOperationBoolResultClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedWindowManagementOperationBoolResultClass) Alloc() SLSBridgedWindowManagementOperationBoolResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationBoolResult](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedWindowManagementOperationBoolResult.BoolValue]
//   - [SLSBridgedWindowManagementOperationBoolResult.InitWithBoolValue]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationBoolResult
type SLSBridgedWindowManagementOperationBoolResult struct {
	SLSBridgedWindowManagementOperationResult
}

// SLSBridgedWindowManagementOperationBoolResultFromID constructs a [SLSBridgedWindowManagementOperationBoolResult] from an objc.ID.
func SLSBridgedWindowManagementOperationBoolResultFromID(id objc.ID) SLSBridgedWindowManagementOperationBoolResult {
	return SLSBridgedWindowManagementOperationBoolResult{SLSBridgedWindowManagementOperationResult: SLSBridgedWindowManagementOperationResultFromID(id)}
}

// Ensure SLSBridgedWindowManagementOperationBoolResult implements ISLSBridgedWindowManagementOperationBoolResult.
var _ ISLSBridgedWindowManagementOperationBoolResult = SLSBridgedWindowManagementOperationBoolResult{}

// An interface definition for the [SLSBridgedWindowManagementOperationBoolResult] class.
//
// # Methods
//
//   - [ISLSBridgedWindowManagementOperationBoolResult.BoolValue]
//   - [ISLSBridgedWindowManagementOperationBoolResult.InitWithBoolValue]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationBoolResult
type ISLSBridgedWindowManagementOperationBoolResult interface {
	ISLSBridgedWindowManagementOperationResult

	// Topic: Methods

	BoolValue() bool
	InitWithBoolValue(value bool) SLSBridgedWindowManagementOperationBoolResult
}

// Init initializes the instance.
func (s SLSBridgedWindowManagementOperationBoolResult) Init() SLSBridgedWindowManagementOperationBoolResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationBoolResult](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedWindowManagementOperationBoolResult) Autorelease() SLSBridgedWindowManagementOperationBoolResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationBoolResult](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedWindowManagementOperationBoolResult creates a new SLSBridgedWindowManagementOperationBoolResult instance.
func NewSLSBridgedWindowManagementOperationBoolResult() SLSBridgedWindowManagementOperationBoolResult {
	class := getSLSBridgedWindowManagementOperationBoolResultClass()
	rv := objc.Send[SLSBridgedWindowManagementOperationBoolResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationBoolResult/initWithBoolValue:
func NewSLSBridgedWindowManagementOperationBoolResultWithBoolValue(value bool) SLSBridgedWindowManagementOperationBoolResult {
	instance := getSLSBridgedWindowManagementOperationBoolResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBoolValue:"), value)
	return SLSBridgedWindowManagementOperationBoolResultFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationBoolResult/initWithCoder:
func NewSLSBridgedWindowManagementOperationBoolResultWithCoder(coder objectivec.IObject) SLSBridgedWindowManagementOperationBoolResult {
	instance := getSLSBridgedWindowManagementOperationBoolResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedWindowManagementOperationBoolResultFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationBoolResult/initWithBoolValue:
func (s SLSBridgedWindowManagementOperationBoolResult) InitWithBoolValue(value bool) SLSBridgedWindowManagementOperationBoolResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationBoolResult](s.ID, objc.Sel("initWithBoolValue:"), value)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationBoolResult/boolValue
func (s SLSBridgedWindowManagementOperationBoolResult) BoolValue() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("boolValue"))
	return rv
}
