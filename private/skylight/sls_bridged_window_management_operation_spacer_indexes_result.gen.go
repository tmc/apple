// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedWindowManagementOperationSpacerIndexesResult] class.
var (
	_SLSBridgedWindowManagementOperationSpacerIndexesResultClass     SLSBridgedWindowManagementOperationSpacerIndexesResultClass
	_SLSBridgedWindowManagementOperationSpacerIndexesResultClassOnce sync.Once
)

func getSLSBridgedWindowManagementOperationSpacerIndexesResultClass() SLSBridgedWindowManagementOperationSpacerIndexesResultClass {
	_SLSBridgedWindowManagementOperationSpacerIndexesResultClassOnce.Do(func() {
		_SLSBridgedWindowManagementOperationSpacerIndexesResultClass = SLSBridgedWindowManagementOperationSpacerIndexesResultClass{class: objc.GetClass("SLSBridgedWindowManagementOperationSpacerIndexesResult")}
	})
	return _SLSBridgedWindowManagementOperationSpacerIndexesResultClass
}

// GetSLSBridgedWindowManagementOperationSpacerIndexesResultClass returns the class object for SLSBridgedWindowManagementOperationSpacerIndexesResult.
func GetSLSBridgedWindowManagementOperationSpacerIndexesResultClass() SLSBridgedWindowManagementOperationSpacerIndexesResultClass {
	return getSLSBridgedWindowManagementOperationSpacerIndexesResultClass()
}

type SLSBridgedWindowManagementOperationSpacerIndexesResultClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedWindowManagementOperationSpacerIndexesResultClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedWindowManagementOperationSpacerIndexesResultClass) Alloc() SLSBridgedWindowManagementOperationSpacerIndexesResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationSpacerIndexesResult](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedWindowManagementOperationSpacerIndexesResult.HorizontalIndex]
//   - [SLSBridgedWindowManagementOperationSpacerIndexesResult.VerticalIndex]
//   - [SLSBridgedWindowManagementOperationSpacerIndexesResult.InitWithVerticalIndexHorizontalIndex]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationSpacerIndexesResult
type SLSBridgedWindowManagementOperationSpacerIndexesResult struct {
	SLSBridgedWindowManagementOperationResult
}

// SLSBridgedWindowManagementOperationSpacerIndexesResultFromID constructs a [SLSBridgedWindowManagementOperationSpacerIndexesResult] from an objc.ID.
func SLSBridgedWindowManagementOperationSpacerIndexesResultFromID(id objc.ID) SLSBridgedWindowManagementOperationSpacerIndexesResult {
	return SLSBridgedWindowManagementOperationSpacerIndexesResult{SLSBridgedWindowManagementOperationResult: SLSBridgedWindowManagementOperationResultFromID(id)}
}

// Ensure SLSBridgedWindowManagementOperationSpacerIndexesResult implements ISLSBridgedWindowManagementOperationSpacerIndexesResult.
var _ ISLSBridgedWindowManagementOperationSpacerIndexesResult = SLSBridgedWindowManagementOperationSpacerIndexesResult{}

// An interface definition for the [SLSBridgedWindowManagementOperationSpacerIndexesResult] class.
//
// # Methods
//
//   - [ISLSBridgedWindowManagementOperationSpacerIndexesResult.HorizontalIndex]
//   - [ISLSBridgedWindowManagementOperationSpacerIndexesResult.VerticalIndex]
//   - [ISLSBridgedWindowManagementOperationSpacerIndexesResult.InitWithVerticalIndexHorizontalIndex]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationSpacerIndexesResult
type ISLSBridgedWindowManagementOperationSpacerIndexesResult interface {
	ISLSBridgedWindowManagementOperationResult

	// Topic: Methods

	HorizontalIndex() uint64
	VerticalIndex() uint64
	InitWithVerticalIndexHorizontalIndex(index uint64, index2 uint64) SLSBridgedWindowManagementOperationSpacerIndexesResult
}

// Init initializes the instance.
func (s SLSBridgedWindowManagementOperationSpacerIndexesResult) Init() SLSBridgedWindowManagementOperationSpacerIndexesResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationSpacerIndexesResult](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedWindowManagementOperationSpacerIndexesResult) Autorelease() SLSBridgedWindowManagementOperationSpacerIndexesResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationSpacerIndexesResult](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedWindowManagementOperationSpacerIndexesResult creates a new SLSBridgedWindowManagementOperationSpacerIndexesResult instance.
func NewSLSBridgedWindowManagementOperationSpacerIndexesResult() SLSBridgedWindowManagementOperationSpacerIndexesResult {
	class := getSLSBridgedWindowManagementOperationSpacerIndexesResultClass()
	rv := objc.Send[SLSBridgedWindowManagementOperationSpacerIndexesResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationSpacerIndexesResult/initWithCoder:
func NewSLSBridgedWindowManagementOperationSpacerIndexesResultWithCoder(coder objectivec.IObject) SLSBridgedWindowManagementOperationSpacerIndexesResult {
	instance := getSLSBridgedWindowManagementOperationSpacerIndexesResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedWindowManagementOperationSpacerIndexesResultFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationSpacerIndexesResult/initWithVerticalIndex:horizontalIndex:
func NewSLSBridgedWindowManagementOperationSpacerIndexesResultWithVerticalIndexHorizontalIndex(index uint64, index2 uint64) SLSBridgedWindowManagementOperationSpacerIndexesResult {
	instance := getSLSBridgedWindowManagementOperationSpacerIndexesResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithVerticalIndex:horizontalIndex:"), index, index2)
	return SLSBridgedWindowManagementOperationSpacerIndexesResultFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationSpacerIndexesResult/initWithVerticalIndex:horizontalIndex:
func (s SLSBridgedWindowManagementOperationSpacerIndexesResult) InitWithVerticalIndexHorizontalIndex(index uint64, index2 uint64) SLSBridgedWindowManagementOperationSpacerIndexesResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationSpacerIndexesResult](s.ID, objc.Sel("initWithVerticalIndex:horizontalIndex:"), index, index2)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationSpacerIndexesResult/horizontalIndex
func (s SLSBridgedWindowManagementOperationSpacerIndexesResult) HorizontalIndex() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("horizontalIndex"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationSpacerIndexesResult/verticalIndex
func (s SLSBridgedWindowManagementOperationSpacerIndexesResult) VerticalIndex() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("verticalIndex"))
	return rv
}
