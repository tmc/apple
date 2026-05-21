// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedWindowManagementOperationSpaceResizeDirectionsResult] class.
var (
	_SLSBridgedWindowManagementOperationSpaceResizeDirectionsResultClass     SLSBridgedWindowManagementOperationSpaceResizeDirectionsResultClass
	_SLSBridgedWindowManagementOperationSpaceResizeDirectionsResultClassOnce sync.Once
)

func getSLSBridgedWindowManagementOperationSpaceResizeDirectionsResultClass() SLSBridgedWindowManagementOperationSpaceResizeDirectionsResultClass {
	_SLSBridgedWindowManagementOperationSpaceResizeDirectionsResultClassOnce.Do(func() {
		_SLSBridgedWindowManagementOperationSpaceResizeDirectionsResultClass = SLSBridgedWindowManagementOperationSpaceResizeDirectionsResultClass{class: objc.GetClass("SLSBridgedWindowManagementOperationSpaceResizeDirectionsResult")}
	})
	return _SLSBridgedWindowManagementOperationSpaceResizeDirectionsResultClass
}

// GetSLSBridgedWindowManagementOperationSpaceResizeDirectionsResultClass returns the class object for SLSBridgedWindowManagementOperationSpaceResizeDirectionsResult.
func GetSLSBridgedWindowManagementOperationSpaceResizeDirectionsResultClass() SLSBridgedWindowManagementOperationSpaceResizeDirectionsResultClass {
	return getSLSBridgedWindowManagementOperationSpaceResizeDirectionsResultClass()
}

type SLSBridgedWindowManagementOperationSpaceResizeDirectionsResultClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedWindowManagementOperationSpaceResizeDirectionsResultClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedWindowManagementOperationSpaceResizeDirectionsResultClass) Alloc() SLSBridgedWindowManagementOperationSpaceResizeDirectionsResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationSpaceResizeDirectionsResult](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedWindowManagementOperationSpaceResizeDirectionsResult.SpaceResizeDirections]
//   - [SLSBridgedWindowManagementOperationSpaceResizeDirectionsResult.InitWithSpaceResizeDirections]
type SLSBridgedWindowManagementOperationSpaceResizeDirectionsResult struct {
	SLSBridgedWindowManagementOperationResult
}

// SLSBridgedWindowManagementOperationSpaceResizeDirectionsResultFromID constructs a [SLSBridgedWindowManagementOperationSpaceResizeDirectionsResult] from an objc.ID.
func SLSBridgedWindowManagementOperationSpaceResizeDirectionsResultFromID(id objc.ID) SLSBridgedWindowManagementOperationSpaceResizeDirectionsResult {
	return SLSBridgedWindowManagementOperationSpaceResizeDirectionsResult{SLSBridgedWindowManagementOperationResult: SLSBridgedWindowManagementOperationResultFromID(id)}
}

// Ensure SLSBridgedWindowManagementOperationSpaceResizeDirectionsResult implements ISLSBridgedWindowManagementOperationSpaceResizeDirectionsResult.
var _ ISLSBridgedWindowManagementOperationSpaceResizeDirectionsResult = SLSBridgedWindowManagementOperationSpaceResizeDirectionsResult{}

// An interface definition for the [SLSBridgedWindowManagementOperationSpaceResizeDirectionsResult] class.
//
// # Methods
//
//   - [ISLSBridgedWindowManagementOperationSpaceResizeDirectionsResult.SpaceResizeDirections]
//   - [ISLSBridgedWindowManagementOperationSpaceResizeDirectionsResult.InitWithSpaceResizeDirections]
type ISLSBridgedWindowManagementOperationSpaceResizeDirectionsResult interface {
	ISLSBridgedWindowManagementOperationResult

	// Topic: Methods

	SpaceResizeDirections() uint64
	InitWithSpaceResizeDirections(directions uint64) SLSBridgedWindowManagementOperationSpaceResizeDirectionsResult
}

// Init initializes the instance.
func (s SLSBridgedWindowManagementOperationSpaceResizeDirectionsResult) Init() SLSBridgedWindowManagementOperationSpaceResizeDirectionsResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationSpaceResizeDirectionsResult](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedWindowManagementOperationSpaceResizeDirectionsResult) Autorelease() SLSBridgedWindowManagementOperationSpaceResizeDirectionsResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationSpaceResizeDirectionsResult](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedWindowManagementOperationSpaceResizeDirectionsResult creates a new SLSBridgedWindowManagementOperationSpaceResizeDirectionsResult instance.
func NewSLSBridgedWindowManagementOperationSpaceResizeDirectionsResult() SLSBridgedWindowManagementOperationSpaceResizeDirectionsResult {
	class := getSLSBridgedWindowManagementOperationSpaceResizeDirectionsResultClass()
	rv := objc.Send[SLSBridgedWindowManagementOperationSpaceResizeDirectionsResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSLSBridgedWindowManagementOperationSpaceResizeDirectionsResultWithCoder(coder objectivec.IObject) SLSBridgedWindowManagementOperationSpaceResizeDirectionsResult {
	instance := getSLSBridgedWindowManagementOperationSpaceResizeDirectionsResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedWindowManagementOperationSpaceResizeDirectionsResultFromID(rv)
}

func NewSLSBridgedWindowManagementOperationSpaceResizeDirectionsResultWithSpaceResizeDirections(directions uint64) SLSBridgedWindowManagementOperationSpaceResizeDirectionsResult {
	instance := getSLSBridgedWindowManagementOperationSpaceResizeDirectionsResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpaceResizeDirections:"), directions)
	return SLSBridgedWindowManagementOperationSpaceResizeDirectionsResultFromID(rv)
}

func (s SLSBridgedWindowManagementOperationSpaceResizeDirectionsResult) InitWithSpaceResizeDirections(directions uint64) SLSBridgedWindowManagementOperationSpaceResizeDirectionsResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationSpaceResizeDirectionsResult](s.ID, objc.Sel("initWithSpaceResizeDirections:"), directions)
	return rv
}

func (s SLSBridgedWindowManagementOperationSpaceResizeDirectionsResult) SpaceResizeDirections() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("spaceResizeDirections"))
	return rv
}
