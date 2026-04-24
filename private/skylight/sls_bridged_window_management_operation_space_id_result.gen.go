// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedWindowManagementOperationSpaceIDResult] class.
var (
	_SLSBridgedWindowManagementOperationSpaceIDResultClass     SLSBridgedWindowManagementOperationSpaceIDResultClass
	_SLSBridgedWindowManagementOperationSpaceIDResultClassOnce sync.Once
)

func getSLSBridgedWindowManagementOperationSpaceIDResultClass() SLSBridgedWindowManagementOperationSpaceIDResultClass {
	_SLSBridgedWindowManagementOperationSpaceIDResultClassOnce.Do(func() {
		_SLSBridgedWindowManagementOperationSpaceIDResultClass = SLSBridgedWindowManagementOperationSpaceIDResultClass{class: objc.GetClass("SLSBridgedWindowManagementOperationSpaceIDResult")}
	})
	return _SLSBridgedWindowManagementOperationSpaceIDResultClass
}

// GetSLSBridgedWindowManagementOperationSpaceIDResultClass returns the class object for SLSBridgedWindowManagementOperationSpaceIDResult.
func GetSLSBridgedWindowManagementOperationSpaceIDResultClass() SLSBridgedWindowManagementOperationSpaceIDResultClass {
	return getSLSBridgedWindowManagementOperationSpaceIDResultClass()
}

type SLSBridgedWindowManagementOperationSpaceIDResultClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedWindowManagementOperationSpaceIDResultClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedWindowManagementOperationSpaceIDResultClass) Alloc() SLSBridgedWindowManagementOperationSpaceIDResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationSpaceIDResult](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedWindowManagementOperationSpaceIDResult.SpaceID]
//   - [SLSBridgedWindowManagementOperationSpaceIDResult.InitWithSpaceID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationSpaceIDResult
type SLSBridgedWindowManagementOperationSpaceIDResult struct {
	SLSBridgedWindowManagementOperationResult
}

// SLSBridgedWindowManagementOperationSpaceIDResultFromID constructs a [SLSBridgedWindowManagementOperationSpaceIDResult] from an objc.ID.
func SLSBridgedWindowManagementOperationSpaceIDResultFromID(id objc.ID) SLSBridgedWindowManagementOperationSpaceIDResult {
	return SLSBridgedWindowManagementOperationSpaceIDResult{SLSBridgedWindowManagementOperationResult: SLSBridgedWindowManagementOperationResultFromID(id)}
}

// Ensure SLSBridgedWindowManagementOperationSpaceIDResult implements ISLSBridgedWindowManagementOperationSpaceIDResult.
var _ ISLSBridgedWindowManagementOperationSpaceIDResult = SLSBridgedWindowManagementOperationSpaceIDResult{}

// An interface definition for the [SLSBridgedWindowManagementOperationSpaceIDResult] class.
//
// # Methods
//
//   - [ISLSBridgedWindowManagementOperationSpaceIDResult.SpaceID]
//   - [ISLSBridgedWindowManagementOperationSpaceIDResult.InitWithSpaceID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationSpaceIDResult
type ISLSBridgedWindowManagementOperationSpaceIDResult interface {
	ISLSBridgedWindowManagementOperationResult

	// Topic: Methods

	SpaceID() uint64
	InitWithSpaceID(id uint64) SLSBridgedWindowManagementOperationSpaceIDResult
}

// Init initializes the instance.
func (s SLSBridgedWindowManagementOperationSpaceIDResult) Init() SLSBridgedWindowManagementOperationSpaceIDResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationSpaceIDResult](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedWindowManagementOperationSpaceIDResult) Autorelease() SLSBridgedWindowManagementOperationSpaceIDResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationSpaceIDResult](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedWindowManagementOperationSpaceIDResult creates a new SLSBridgedWindowManagementOperationSpaceIDResult instance.
func NewSLSBridgedWindowManagementOperationSpaceIDResult() SLSBridgedWindowManagementOperationSpaceIDResult {
	class := getSLSBridgedWindowManagementOperationSpaceIDResultClass()
	rv := objc.Send[SLSBridgedWindowManagementOperationSpaceIDResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationSpaceIDResult/initWithCoder:
func NewSLSBridgedWindowManagementOperationSpaceIDResultWithCoder(coder objectivec.IObject) SLSBridgedWindowManagementOperationSpaceIDResult {
	instance := getSLSBridgedWindowManagementOperationSpaceIDResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedWindowManagementOperationSpaceIDResultFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationSpaceIDResult/initWithSpaceID:
func NewSLSBridgedWindowManagementOperationSpaceIDResultWithSpaceID(id uint64) SLSBridgedWindowManagementOperationSpaceIDResult {
	instance := getSLSBridgedWindowManagementOperationSpaceIDResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpaceID:"), id)
	return SLSBridgedWindowManagementOperationSpaceIDResultFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationSpaceIDResult/initWithSpaceID:
func (s SLSBridgedWindowManagementOperationSpaceIDResult) InitWithSpaceID(id uint64) SLSBridgedWindowManagementOperationSpaceIDResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationSpaceIDResult](s.ID, objc.Sel("initWithSpaceID:"), id)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationSpaceIDResult/spaceID
func (s SLSBridgedWindowManagementOperationSpaceIDResult) SpaceID() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("spaceID"))
	return rv
}
