// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedWindowManagementOperationProcessIdentifierResult] class.
var (
	_SLSBridgedWindowManagementOperationProcessIdentifierResultClass     SLSBridgedWindowManagementOperationProcessIdentifierResultClass
	_SLSBridgedWindowManagementOperationProcessIdentifierResultClassOnce sync.Once
)

func getSLSBridgedWindowManagementOperationProcessIdentifierResultClass() SLSBridgedWindowManagementOperationProcessIdentifierResultClass {
	_SLSBridgedWindowManagementOperationProcessIdentifierResultClassOnce.Do(func() {
		_SLSBridgedWindowManagementOperationProcessIdentifierResultClass = SLSBridgedWindowManagementOperationProcessIdentifierResultClass{class: objc.GetClass("SLSBridgedWindowManagementOperationProcessIdentifierResult")}
	})
	return _SLSBridgedWindowManagementOperationProcessIdentifierResultClass
}

// GetSLSBridgedWindowManagementOperationProcessIdentifierResultClass returns the class object for SLSBridgedWindowManagementOperationProcessIdentifierResult.
func GetSLSBridgedWindowManagementOperationProcessIdentifierResultClass() SLSBridgedWindowManagementOperationProcessIdentifierResultClass {
	return getSLSBridgedWindowManagementOperationProcessIdentifierResultClass()
}

type SLSBridgedWindowManagementOperationProcessIdentifierResultClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedWindowManagementOperationProcessIdentifierResultClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedWindowManagementOperationProcessIdentifierResultClass) Alloc() SLSBridgedWindowManagementOperationProcessIdentifierResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationProcessIdentifierResult](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedWindowManagementOperationProcessIdentifierResult.ProcessIdentifier]
//   - [SLSBridgedWindowManagementOperationProcessIdentifierResult.InitWithProcessIdentifier]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationProcessIdentifierResult
type SLSBridgedWindowManagementOperationProcessIdentifierResult struct {
	SLSBridgedWindowManagementOperationResult
}

// SLSBridgedWindowManagementOperationProcessIdentifierResultFromID constructs a [SLSBridgedWindowManagementOperationProcessIdentifierResult] from an objc.ID.
func SLSBridgedWindowManagementOperationProcessIdentifierResultFromID(id objc.ID) SLSBridgedWindowManagementOperationProcessIdentifierResult {
	return SLSBridgedWindowManagementOperationProcessIdentifierResult{SLSBridgedWindowManagementOperationResult: SLSBridgedWindowManagementOperationResultFromID(id)}
}

// Ensure SLSBridgedWindowManagementOperationProcessIdentifierResult implements ISLSBridgedWindowManagementOperationProcessIdentifierResult.
var _ ISLSBridgedWindowManagementOperationProcessIdentifierResult = SLSBridgedWindowManagementOperationProcessIdentifierResult{}

// An interface definition for the [SLSBridgedWindowManagementOperationProcessIdentifierResult] class.
//
// # Methods
//
//   - [ISLSBridgedWindowManagementOperationProcessIdentifierResult.ProcessIdentifier]
//   - [ISLSBridgedWindowManagementOperationProcessIdentifierResult.InitWithProcessIdentifier]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationProcessIdentifierResult
type ISLSBridgedWindowManagementOperationProcessIdentifierResult interface {
	ISLSBridgedWindowManagementOperationResult

	// Topic: Methods

	ProcessIdentifier() int
	InitWithProcessIdentifier(identifier int) SLSBridgedWindowManagementOperationProcessIdentifierResult
}

// Init initializes the instance.
func (s SLSBridgedWindowManagementOperationProcessIdentifierResult) Init() SLSBridgedWindowManagementOperationProcessIdentifierResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationProcessIdentifierResult](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedWindowManagementOperationProcessIdentifierResult) Autorelease() SLSBridgedWindowManagementOperationProcessIdentifierResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationProcessIdentifierResult](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedWindowManagementOperationProcessIdentifierResult creates a new SLSBridgedWindowManagementOperationProcessIdentifierResult instance.
func NewSLSBridgedWindowManagementOperationProcessIdentifierResult() SLSBridgedWindowManagementOperationProcessIdentifierResult {
	class := getSLSBridgedWindowManagementOperationProcessIdentifierResultClass()
	rv := objc.Send[SLSBridgedWindowManagementOperationProcessIdentifierResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationProcessIdentifierResult/initWithCoder:
func NewSLSBridgedWindowManagementOperationProcessIdentifierResultWithCoder(coder objectivec.IObject) SLSBridgedWindowManagementOperationProcessIdentifierResult {
	instance := getSLSBridgedWindowManagementOperationProcessIdentifierResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedWindowManagementOperationProcessIdentifierResultFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationProcessIdentifierResult/initWithProcessIdentifier:
func NewSLSBridgedWindowManagementOperationProcessIdentifierResultWithProcessIdentifier(identifier int) SLSBridgedWindowManagementOperationProcessIdentifierResult {
	instance := getSLSBridgedWindowManagementOperationProcessIdentifierResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithProcessIdentifier:"), identifier)
	return SLSBridgedWindowManagementOperationProcessIdentifierResultFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationProcessIdentifierResult/initWithProcessIdentifier:
func (s SLSBridgedWindowManagementOperationProcessIdentifierResult) InitWithProcessIdentifier(identifier int) SLSBridgedWindowManagementOperationProcessIdentifierResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationProcessIdentifierResult](s.ID, objc.Sel("initWithProcessIdentifier:"), identifier)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationProcessIdentifierResult/processIdentifier
func (s SLSBridgedWindowManagementOperationProcessIdentifierResult) ProcessIdentifier() int {
	rv := objc.Send[int](s.ID, objc.Sel("processIdentifier"))
	return rv
}
