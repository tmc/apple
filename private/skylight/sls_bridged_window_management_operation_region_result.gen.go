// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedWindowManagementOperationRegionResult] class.
var (
	_SLSBridgedWindowManagementOperationRegionResultClass     SLSBridgedWindowManagementOperationRegionResultClass
	_SLSBridgedWindowManagementOperationRegionResultClassOnce sync.Once
)

func getSLSBridgedWindowManagementOperationRegionResultClass() SLSBridgedWindowManagementOperationRegionResultClass {
	_SLSBridgedWindowManagementOperationRegionResultClassOnce.Do(func() {
		_SLSBridgedWindowManagementOperationRegionResultClass = SLSBridgedWindowManagementOperationRegionResultClass{class: objc.GetClass("SLSBridgedWindowManagementOperationRegionResult")}
	})
	return _SLSBridgedWindowManagementOperationRegionResultClass
}

// GetSLSBridgedWindowManagementOperationRegionResultClass returns the class object for SLSBridgedWindowManagementOperationRegionResult.
func GetSLSBridgedWindowManagementOperationRegionResultClass() SLSBridgedWindowManagementOperationRegionResultClass {
	return getSLSBridgedWindowManagementOperationRegionResultClass()
}

type SLSBridgedWindowManagementOperationRegionResultClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedWindowManagementOperationRegionResultClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedWindowManagementOperationRegionResultClass) Alloc() SLSBridgedWindowManagementOperationRegionResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationRegionResult](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedWindowManagementOperationRegionResult.CopyRegion]
//   - [SLSBridgedWindowManagementOperationRegionResult.InitWithRegion]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationRegionResult
type SLSBridgedWindowManagementOperationRegionResult struct {
	SLSBridgedWindowManagementOperationResult
}

// SLSBridgedWindowManagementOperationRegionResultFromID constructs a [SLSBridgedWindowManagementOperationRegionResult] from an objc.ID.
func SLSBridgedWindowManagementOperationRegionResultFromID(id objc.ID) SLSBridgedWindowManagementOperationRegionResult {
	return SLSBridgedWindowManagementOperationRegionResult{SLSBridgedWindowManagementOperationResult: SLSBridgedWindowManagementOperationResultFromID(id)}
}

// Ensure SLSBridgedWindowManagementOperationRegionResult implements ISLSBridgedWindowManagementOperationRegionResult.
var _ ISLSBridgedWindowManagementOperationRegionResult = SLSBridgedWindowManagementOperationRegionResult{}

// An interface definition for the [SLSBridgedWindowManagementOperationRegionResult] class.
//
// # Methods
//
//   - [ISLSBridgedWindowManagementOperationRegionResult.CopyRegion]
//   - [ISLSBridgedWindowManagementOperationRegionResult.InitWithRegion]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationRegionResult
type ISLSBridgedWindowManagementOperationRegionResult interface {
	ISLSBridgedWindowManagementOperationResult

	// Topic: Methods

	CopyRegion() uintptr
	InitWithRegion(region uintptr) SLSBridgedWindowManagementOperationRegionResult
}

// Init initializes the instance.
func (s SLSBridgedWindowManagementOperationRegionResult) Init() SLSBridgedWindowManagementOperationRegionResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationRegionResult](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedWindowManagementOperationRegionResult) Autorelease() SLSBridgedWindowManagementOperationRegionResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationRegionResult](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedWindowManagementOperationRegionResult creates a new SLSBridgedWindowManagementOperationRegionResult instance.
func NewSLSBridgedWindowManagementOperationRegionResult() SLSBridgedWindowManagementOperationRegionResult {
	class := getSLSBridgedWindowManagementOperationRegionResultClass()
	rv := objc.Send[SLSBridgedWindowManagementOperationRegionResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationRegionResult/initWithCoder:
func NewSLSBridgedWindowManagementOperationRegionResultWithCoder(coder objectivec.IObject) SLSBridgedWindowManagementOperationRegionResult {
	instance := getSLSBridgedWindowManagementOperationRegionResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedWindowManagementOperationRegionResultFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationRegionResult/initWithRegion:
func NewSLSBridgedWindowManagementOperationRegionResultWithRegion(region uintptr) SLSBridgedWindowManagementOperationRegionResult {
	instance := getSLSBridgedWindowManagementOperationRegionResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRegion:"), region)
	return SLSBridgedWindowManagementOperationRegionResultFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationRegionResult/copyRegion
func (s SLSBridgedWindowManagementOperationRegionResult) CopyRegion() uintptr {
	rv := objc.Send[uintptr](s.ID, objc.Sel("copyRegion"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationRegionResult/initWithRegion:
func (s SLSBridgedWindowManagementOperationRegionResult) InitWithRegion(region uintptr) SLSBridgedWindowManagementOperationRegionResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationRegionResult](s.ID, objc.Sel("initWithRegion:"), region)
	return rv
}
