// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedWindowManagementOperationSpacersResult] class.
var (
	_SLSBridgedWindowManagementOperationSpacersResultClass     SLSBridgedWindowManagementOperationSpacersResultClass
	_SLSBridgedWindowManagementOperationSpacersResultClassOnce sync.Once
)

func getSLSBridgedWindowManagementOperationSpacersResultClass() SLSBridgedWindowManagementOperationSpacersResultClass {
	_SLSBridgedWindowManagementOperationSpacersResultClassOnce.Do(func() {
		_SLSBridgedWindowManagementOperationSpacersResultClass = SLSBridgedWindowManagementOperationSpacersResultClass{class: objc.GetClass("SLSBridgedWindowManagementOperationSpacersResult")}
	})
	return _SLSBridgedWindowManagementOperationSpacersResultClass
}

// GetSLSBridgedWindowManagementOperationSpacersResultClass returns the class object for SLSBridgedWindowManagementOperationSpacersResult.
func GetSLSBridgedWindowManagementOperationSpacersResultClass() SLSBridgedWindowManagementOperationSpacersResultClass {
	return getSLSBridgedWindowManagementOperationSpacersResultClass()
}

type SLSBridgedWindowManagementOperationSpacersResultClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedWindowManagementOperationSpacersResultClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedWindowManagementOperationSpacersResultClass) Alloc() SLSBridgedWindowManagementOperationSpacersResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationSpacersResult](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedWindowManagementOperationSpacersResult.HorizontalIndex]
//   - [SLSBridgedWindowManagementOperationSpacersResult.Rect]
//   - [SLSBridgedWindowManagementOperationSpacersResult.VerticalIndex]
//   - [SLSBridgedWindowManagementOperationSpacersResult.InitWithVerticalIndexHorizontalIndexRect]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationSpacersResult
type SLSBridgedWindowManagementOperationSpacersResult struct {
	SLSBridgedWindowManagementOperationResult
}

// SLSBridgedWindowManagementOperationSpacersResultFromID constructs a [SLSBridgedWindowManagementOperationSpacersResult] from an objc.ID.
func SLSBridgedWindowManagementOperationSpacersResultFromID(id objc.ID) SLSBridgedWindowManagementOperationSpacersResult {
	return SLSBridgedWindowManagementOperationSpacersResult{SLSBridgedWindowManagementOperationResult: SLSBridgedWindowManagementOperationResultFromID(id)}
}

// Ensure SLSBridgedWindowManagementOperationSpacersResult implements ISLSBridgedWindowManagementOperationSpacersResult.
var _ ISLSBridgedWindowManagementOperationSpacersResult = SLSBridgedWindowManagementOperationSpacersResult{}

// An interface definition for the [SLSBridgedWindowManagementOperationSpacersResult] class.
//
// # Methods
//
//   - [ISLSBridgedWindowManagementOperationSpacersResult.HorizontalIndex]
//   - [ISLSBridgedWindowManagementOperationSpacersResult.Rect]
//   - [ISLSBridgedWindowManagementOperationSpacersResult.VerticalIndex]
//   - [ISLSBridgedWindowManagementOperationSpacersResult.InitWithVerticalIndexHorizontalIndexRect]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationSpacersResult
type ISLSBridgedWindowManagementOperationSpacersResult interface {
	ISLSBridgedWindowManagementOperationResult

	// Topic: Methods

	HorizontalIndex() uint64
	Rect() corefoundation.CGRect
	VerticalIndex() uint64
	InitWithVerticalIndexHorizontalIndexRect(index uint64, index2 uint64, rect corefoundation.CGRect) SLSBridgedWindowManagementOperationSpacersResult
}

// Init initializes the instance.
func (s SLSBridgedWindowManagementOperationSpacersResult) Init() SLSBridgedWindowManagementOperationSpacersResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationSpacersResult](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedWindowManagementOperationSpacersResult) Autorelease() SLSBridgedWindowManagementOperationSpacersResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationSpacersResult](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedWindowManagementOperationSpacersResult creates a new SLSBridgedWindowManagementOperationSpacersResult instance.
func NewSLSBridgedWindowManagementOperationSpacersResult() SLSBridgedWindowManagementOperationSpacersResult {
	class := getSLSBridgedWindowManagementOperationSpacersResultClass()
	rv := objc.Send[SLSBridgedWindowManagementOperationSpacersResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationSpacersResult/initWithCoder:
func NewSLSBridgedWindowManagementOperationSpacersResultWithCoder(coder objectivec.IObject) SLSBridgedWindowManagementOperationSpacersResult {
	instance := getSLSBridgedWindowManagementOperationSpacersResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedWindowManagementOperationSpacersResultFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationSpacersResult/initWithVerticalIndex:horizontalIndex:rect:
func NewSLSBridgedWindowManagementOperationSpacersResultWithVerticalIndexHorizontalIndexRect(index uint64, index2 uint64, rect corefoundation.CGRect) SLSBridgedWindowManagementOperationSpacersResult {
	instance := getSLSBridgedWindowManagementOperationSpacersResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithVerticalIndex:horizontalIndex:rect:"), index, index2, rect)
	return SLSBridgedWindowManagementOperationSpacersResultFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationSpacersResult/initWithVerticalIndex:horizontalIndex:rect:
func (s SLSBridgedWindowManagementOperationSpacersResult) InitWithVerticalIndexHorizontalIndexRect(index uint64, index2 uint64, rect corefoundation.CGRect) SLSBridgedWindowManagementOperationSpacersResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationSpacersResult](s.ID, objc.Sel("initWithVerticalIndex:horizontalIndex:rect:"), index, index2, rect)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationSpacersResult/horizontalIndex
func (s SLSBridgedWindowManagementOperationSpacersResult) HorizontalIndex() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("horizontalIndex"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationSpacersResult/rect
func (s SLSBridgedWindowManagementOperationSpacersResult) Rect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s.ID, objc.Sel("rect"))
	return corefoundation.CGRect(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationSpacersResult/verticalIndex
func (s SLSBridgedWindowManagementOperationSpacersResult) VerticalIndex() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("verticalIndex"))
	return rv
}
