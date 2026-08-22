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
	rv := objc.SendIfResponds[SLSBridgedWindowManagementOperationSpaceManagementModeResult](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedWindowManagementOperationSpaceManagementModeResult.SpaceManagementMode]
//   - [SLSBridgedWindowManagementOperationSpaceManagementModeResult.InitWithSpaceManagementMode]
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
type ISLSBridgedWindowManagementOperationSpaceManagementModeResult interface {
	ISLSBridgedWindowManagementOperationResult

	// Topic: Methods

	SpaceManagementMode() uint64
	InitWithSpaceManagementMode(mode uint64) SLSBridgedWindowManagementOperationSpaceManagementModeResult
}

// Init initializes the instance.
func (s SLSBridgedWindowManagementOperationSpaceManagementModeResult) Init() SLSBridgedWindowManagementOperationSpaceManagementModeResult {
	rv := objc.SendIfResponds[SLSBridgedWindowManagementOperationSpaceManagementModeResult](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedWindowManagementOperationSpaceManagementModeResult) Autorelease() SLSBridgedWindowManagementOperationSpaceManagementModeResult {
	rv := objc.SendIfResponds[SLSBridgedWindowManagementOperationSpaceManagementModeResult](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedWindowManagementOperationSpaceManagementModeResult creates a new SLSBridgedWindowManagementOperationSpaceManagementModeResult instance.
func NewSLSBridgedWindowManagementOperationSpaceManagementModeResult() SLSBridgedWindowManagementOperationSpaceManagementModeResult {
	class := getSLSBridgedWindowManagementOperationSpaceManagementModeResultClass()
	rv := objc.SendIfResponds[SLSBridgedWindowManagementOperationSpaceManagementModeResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSLSBridgedWindowManagementOperationSpaceManagementModeResultWithCoder(coder objectivec.IObject) SLSBridgedWindowManagementOperationSpaceManagementModeResult {
	instance := getSLSBridgedWindowManagementOperationSpaceManagementModeResultClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedWindowManagementOperationSpaceManagementModeResultFromID(rv)
}

func NewSLSBridgedWindowManagementOperationSpaceManagementModeResultWithSpaceManagementMode(mode uint64) SLSBridgedWindowManagementOperationSpaceManagementModeResult {
	instance := getSLSBridgedWindowManagementOperationSpaceManagementModeResultClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithSpaceManagementMode:"), mode)
	return SLSBridgedWindowManagementOperationSpaceManagementModeResultFromID(rv)
}

func (s SLSBridgedWindowManagementOperationSpaceManagementModeResult) InitWithSpaceManagementMode(mode uint64) SLSBridgedWindowManagementOperationSpaceManagementModeResult {
	rv := objc.SendIfResponds[SLSBridgedWindowManagementOperationSpaceManagementModeResult](s.ID, objc.Sel("initWithSpaceManagementMode:"), mode)
	return rv
}

func (s SLSBridgedWindowManagementOperationSpaceManagementModeResult) SpaceManagementMode() uint64 {
	rv := objc.SendIfResponds[uint64](s.ID, objc.Sel("spaceManagementMode"))
	return rv
}
