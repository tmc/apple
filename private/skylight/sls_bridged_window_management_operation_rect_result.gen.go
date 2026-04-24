// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedWindowManagementOperationRectResult] class.
var (
	_SLSBridgedWindowManagementOperationRectResultClass     SLSBridgedWindowManagementOperationRectResultClass
	_SLSBridgedWindowManagementOperationRectResultClassOnce sync.Once
)

func getSLSBridgedWindowManagementOperationRectResultClass() SLSBridgedWindowManagementOperationRectResultClass {
	_SLSBridgedWindowManagementOperationRectResultClassOnce.Do(func() {
		_SLSBridgedWindowManagementOperationRectResultClass = SLSBridgedWindowManagementOperationRectResultClass{class: objc.GetClass("SLSBridgedWindowManagementOperationRectResult")}
	})
	return _SLSBridgedWindowManagementOperationRectResultClass
}

// GetSLSBridgedWindowManagementOperationRectResultClass returns the class object for SLSBridgedWindowManagementOperationRectResult.
func GetSLSBridgedWindowManagementOperationRectResultClass() SLSBridgedWindowManagementOperationRectResultClass {
	return getSLSBridgedWindowManagementOperationRectResultClass()
}

type SLSBridgedWindowManagementOperationRectResultClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedWindowManagementOperationRectResultClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedWindowManagementOperationRectResultClass) Alloc() SLSBridgedWindowManagementOperationRectResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationRectResult](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedWindowManagementOperationRectResult.Rect]
//   - [SLSBridgedWindowManagementOperationRectResult.InitWithRect]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationRectResult
type SLSBridgedWindowManagementOperationRectResult struct {
	SLSBridgedWindowManagementOperationResult
}

// SLSBridgedWindowManagementOperationRectResultFromID constructs a [SLSBridgedWindowManagementOperationRectResult] from an objc.ID.
func SLSBridgedWindowManagementOperationRectResultFromID(id objc.ID) SLSBridgedWindowManagementOperationRectResult {
	return SLSBridgedWindowManagementOperationRectResult{SLSBridgedWindowManagementOperationResult: SLSBridgedWindowManagementOperationResultFromID(id)}
}

// Ensure SLSBridgedWindowManagementOperationRectResult implements ISLSBridgedWindowManagementOperationRectResult.
var _ ISLSBridgedWindowManagementOperationRectResult = SLSBridgedWindowManagementOperationRectResult{}

// An interface definition for the [SLSBridgedWindowManagementOperationRectResult] class.
//
// # Methods
//
//   - [ISLSBridgedWindowManagementOperationRectResult.Rect]
//   - [ISLSBridgedWindowManagementOperationRectResult.InitWithRect]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationRectResult
type ISLSBridgedWindowManagementOperationRectResult interface {
	ISLSBridgedWindowManagementOperationResult

	// Topic: Methods

	Rect() corefoundation.CGRect
	InitWithRect(rect corefoundation.CGRect) SLSBridgedWindowManagementOperationRectResult
}

// Init initializes the instance.
func (s SLSBridgedWindowManagementOperationRectResult) Init() SLSBridgedWindowManagementOperationRectResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationRectResult](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedWindowManagementOperationRectResult) Autorelease() SLSBridgedWindowManagementOperationRectResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationRectResult](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedWindowManagementOperationRectResult creates a new SLSBridgedWindowManagementOperationRectResult instance.
func NewSLSBridgedWindowManagementOperationRectResult() SLSBridgedWindowManagementOperationRectResult {
	class := getSLSBridgedWindowManagementOperationRectResultClass()
	rv := objc.Send[SLSBridgedWindowManagementOperationRectResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationRectResult/initWithCoder:
func NewSLSBridgedWindowManagementOperationRectResultWithCoder(coder objectivec.IObject) SLSBridgedWindowManagementOperationRectResult {
	instance := getSLSBridgedWindowManagementOperationRectResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedWindowManagementOperationRectResultFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationRectResult/initWithRect:
func NewSLSBridgedWindowManagementOperationRectResultWithRect(rect corefoundation.CGRect) SLSBridgedWindowManagementOperationRectResult {
	instance := getSLSBridgedWindowManagementOperationRectResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRect:"), rect)
	return SLSBridgedWindowManagementOperationRectResultFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationRectResult/initWithRect:
func (s SLSBridgedWindowManagementOperationRectResult) InitWithRect(rect corefoundation.CGRect) SLSBridgedWindowManagementOperationRectResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationRectResult](s.ID, objc.Sel("initWithRect:"), rect)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationRectResult/rect
func (s SLSBridgedWindowManagementOperationRectResult) Rect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s.ID, objc.Sel("rect"))
	return corefoundation.CGRect(rv)
}
