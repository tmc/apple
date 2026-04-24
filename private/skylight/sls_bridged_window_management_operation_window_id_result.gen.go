// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedWindowManagementOperationWindowIDResult] class.
var (
	_SLSBridgedWindowManagementOperationWindowIDResultClass     SLSBridgedWindowManagementOperationWindowIDResultClass
	_SLSBridgedWindowManagementOperationWindowIDResultClassOnce sync.Once
)

func getSLSBridgedWindowManagementOperationWindowIDResultClass() SLSBridgedWindowManagementOperationWindowIDResultClass {
	_SLSBridgedWindowManagementOperationWindowIDResultClassOnce.Do(func() {
		_SLSBridgedWindowManagementOperationWindowIDResultClass = SLSBridgedWindowManagementOperationWindowIDResultClass{class: objc.GetClass("SLSBridgedWindowManagementOperationWindowIDResult")}
	})
	return _SLSBridgedWindowManagementOperationWindowIDResultClass
}

// GetSLSBridgedWindowManagementOperationWindowIDResultClass returns the class object for SLSBridgedWindowManagementOperationWindowIDResult.
func GetSLSBridgedWindowManagementOperationWindowIDResultClass() SLSBridgedWindowManagementOperationWindowIDResultClass {
	return getSLSBridgedWindowManagementOperationWindowIDResultClass()
}

type SLSBridgedWindowManagementOperationWindowIDResultClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedWindowManagementOperationWindowIDResultClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedWindowManagementOperationWindowIDResultClass) Alloc() SLSBridgedWindowManagementOperationWindowIDResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationWindowIDResult](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedWindowManagementOperationWindowIDResult.WindowID]
//   - [SLSBridgedWindowManagementOperationWindowIDResult.InitWithWindowID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationWindowIDResult
type SLSBridgedWindowManagementOperationWindowIDResult struct {
	SLSBridgedWindowManagementOperationResult
}

// SLSBridgedWindowManagementOperationWindowIDResultFromID constructs a [SLSBridgedWindowManagementOperationWindowIDResult] from an objc.ID.
func SLSBridgedWindowManagementOperationWindowIDResultFromID(id objc.ID) SLSBridgedWindowManagementOperationWindowIDResult {
	return SLSBridgedWindowManagementOperationWindowIDResult{SLSBridgedWindowManagementOperationResult: SLSBridgedWindowManagementOperationResultFromID(id)}
}

// Ensure SLSBridgedWindowManagementOperationWindowIDResult implements ISLSBridgedWindowManagementOperationWindowIDResult.
var _ ISLSBridgedWindowManagementOperationWindowIDResult = SLSBridgedWindowManagementOperationWindowIDResult{}

// An interface definition for the [SLSBridgedWindowManagementOperationWindowIDResult] class.
//
// # Methods
//
//   - [ISLSBridgedWindowManagementOperationWindowIDResult.WindowID]
//   - [ISLSBridgedWindowManagementOperationWindowIDResult.InitWithWindowID]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationWindowIDResult
type ISLSBridgedWindowManagementOperationWindowIDResult interface {
	ISLSBridgedWindowManagementOperationResult

	// Topic: Methods

	WindowID() uint32
	InitWithWindowID(id uint32) SLSBridgedWindowManagementOperationWindowIDResult
}

// Init initializes the instance.
func (s SLSBridgedWindowManagementOperationWindowIDResult) Init() SLSBridgedWindowManagementOperationWindowIDResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationWindowIDResult](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedWindowManagementOperationWindowIDResult) Autorelease() SLSBridgedWindowManagementOperationWindowIDResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationWindowIDResult](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedWindowManagementOperationWindowIDResult creates a new SLSBridgedWindowManagementOperationWindowIDResult instance.
func NewSLSBridgedWindowManagementOperationWindowIDResult() SLSBridgedWindowManagementOperationWindowIDResult {
	class := getSLSBridgedWindowManagementOperationWindowIDResultClass()
	rv := objc.Send[SLSBridgedWindowManagementOperationWindowIDResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationWindowIDResult/initWithCoder:
func NewSLSBridgedWindowManagementOperationWindowIDResultWithCoder(coder objectivec.IObject) SLSBridgedWindowManagementOperationWindowIDResult {
	instance := getSLSBridgedWindowManagementOperationWindowIDResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedWindowManagementOperationWindowIDResultFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationWindowIDResult/initWithWindowID:
func NewSLSBridgedWindowManagementOperationWindowIDResultWithWindowID(id uint32) SLSBridgedWindowManagementOperationWindowIDResult {
	instance := getSLSBridgedWindowManagementOperationWindowIDResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithWindowID:"), id)
	return SLSBridgedWindowManagementOperationWindowIDResultFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationWindowIDResult/initWithWindowID:
func (s SLSBridgedWindowManagementOperationWindowIDResult) InitWithWindowID(id uint32) SLSBridgedWindowManagementOperationWindowIDResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationWindowIDResult](s.ID, objc.Sel("initWithWindowID:"), id)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationWindowIDResult/windowID
func (s SLSBridgedWindowManagementOperationWindowIDResult) WindowID() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("windowID"))
	return rv
}
