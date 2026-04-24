// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedWindowManagementOperationSpaceManagementModeResult] class.
var (
	_SLSBridgedWindowManagementOperationSpaceManagementModeResultClass     SLSBridgedWindowManagementOperationSpaceManagementModeResultClass
	_SLSBridgedWindowManagementOperationSpaceManagementModeResultClassOnce sync.Once
)

func getSLSBridgedWindowManagementOperationSpaceManagementModeResultClass() SLSBridgedWindowManagementOperationSpaceManagementModeResultClass {
	_SLSBridgedWindowManagementOperationSpaceManagementModeResultClassOnce.Do(func() {
		_SLSBridgedWindowManagementOperationSpaceManagementModeResultClass = SLSBridgedWindowManagementOperationSpaceManagementModeResultClass{class: objc.GetClass("SLSBridgedWindowManagementOperationSpaceManagementModeResult")}
	})
	return _SLSBridgedWindowManagementOperationSpaceManagementModeResultClass
}

// GetSLSBridgedWindowManagementOperationSpaceManagementModeResultClass returns the class object for SLSBridgedWindowManagementOperationSpaceManagementModeResult.
func GetSLSBridgedWindowManagementOperationSpaceManagementModeResultClass() SLSBridgedWindowManagementOperationSpaceManagementModeResultClass {
	return getSLSBridgedWindowManagementOperationSpaceManagementModeResultClass()
}

type SLSBridgedWindowManagementOperationSpaceManagementModeResultClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedWindowManagementOperationSpaceManagementModeResultClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedWindowManagementOperationSpaceManagementModeResultClass) Alloc() SLSBridgedWindowManagementOperationSpaceManagementModeResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationSpaceManagementModeResult](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedWindowManagementOperationSpaceManagementModeResult.SpaceManagementMode]
//   - [SLSBridgedWindowManagementOperationSpaceManagementModeResult.InitWithSpaceManagementMode]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationSpaceManagementModeResult
type SLSBridgedWindowManagementOperationSpaceManagementModeResult struct {
	SLSBridgedWindowManagementOperationResult
}

// SLSBridgedWindowManagementOperationSpaceManagementModeResultFromID constructs a [SLSBridgedWindowManagementOperationSpaceManagementModeResult] from an objc.ID.
func SLSBridgedWindowManagementOperationSpaceManagementModeResultFromID(id objc.ID) SLSBridgedWindowManagementOperationSpaceManagementModeResult {
	return SLSBridgedWindowManagementOperationSpaceManagementModeResult{SLSBridgedWindowManagementOperationResult: SLSBridgedWindowManagementOperationResultFromID(id)}
}

// Ensure SLSBridgedWindowManagementOperationSpaceManagementModeResult implements ISLSBridgedWindowManagementOperationSpaceManagementModeResult.
var _ ISLSBridgedWindowManagementOperationSpaceManagementModeResult = SLSBridgedWindowManagementOperationSpaceManagementModeResult{}

// An interface definition for the [SLSBridgedWindowManagementOperationSpaceManagementModeResult] class.
//
// # Methods
//
//   - [ISLSBridgedWindowManagementOperationSpaceManagementModeResult.SpaceManagementMode]
//   - [ISLSBridgedWindowManagementOperationSpaceManagementModeResult.InitWithSpaceManagementMode]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationSpaceManagementModeResult
type ISLSBridgedWindowManagementOperationSpaceManagementModeResult interface {
	ISLSBridgedWindowManagementOperationResult

	// Topic: Methods

	SpaceManagementMode() uint64
	InitWithSpaceManagementMode(mode uint64) SLSBridgedWindowManagementOperationSpaceManagementModeResult
}

// Init initializes the instance.
func (s SLSBridgedWindowManagementOperationSpaceManagementModeResult) Init() SLSBridgedWindowManagementOperationSpaceManagementModeResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationSpaceManagementModeResult](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedWindowManagementOperationSpaceManagementModeResult) Autorelease() SLSBridgedWindowManagementOperationSpaceManagementModeResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationSpaceManagementModeResult](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedWindowManagementOperationSpaceManagementModeResult creates a new SLSBridgedWindowManagementOperationSpaceManagementModeResult instance.
func NewSLSBridgedWindowManagementOperationSpaceManagementModeResult() SLSBridgedWindowManagementOperationSpaceManagementModeResult {
	class := getSLSBridgedWindowManagementOperationSpaceManagementModeResultClass()
	rv := objc.Send[SLSBridgedWindowManagementOperationSpaceManagementModeResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationSpaceManagementModeResult/initWithCoder:
func NewSLSBridgedWindowManagementOperationSpaceManagementModeResultWithCoder(coder objectivec.IObject) SLSBridgedWindowManagementOperationSpaceManagementModeResult {
	instance := getSLSBridgedWindowManagementOperationSpaceManagementModeResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedWindowManagementOperationSpaceManagementModeResultFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationSpaceManagementModeResult/initWithSpaceManagementMode:
func NewSLSBridgedWindowManagementOperationSpaceManagementModeResultWithSpaceManagementMode(mode uint64) SLSBridgedWindowManagementOperationSpaceManagementModeResult {
	instance := getSLSBridgedWindowManagementOperationSpaceManagementModeResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpaceManagementMode:"), mode)
	return SLSBridgedWindowManagementOperationSpaceManagementModeResultFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationSpaceManagementModeResult/initWithSpaceManagementMode:
func (s SLSBridgedWindowManagementOperationSpaceManagementModeResult) InitWithSpaceManagementMode(mode uint64) SLSBridgedWindowManagementOperationSpaceManagementModeResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationSpaceManagementModeResult](s.ID, objc.Sel("initWithSpaceManagementMode:"), mode)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationSpaceManagementModeResult/spaceManagementMode
func (s SLSBridgedWindowManagementOperationSpaceManagementModeResult) SpaceManagementMode() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("spaceManagementMode"))
	return rv
}
