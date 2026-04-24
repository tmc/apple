// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedWindowManagementOperationResult] class.
var (
	_SLSBridgedWindowManagementOperationResultClass     SLSBridgedWindowManagementOperationResultClass
	_SLSBridgedWindowManagementOperationResultClassOnce sync.Once
)

func getSLSBridgedWindowManagementOperationResultClass() SLSBridgedWindowManagementOperationResultClass {
	_SLSBridgedWindowManagementOperationResultClassOnce.Do(func() {
		_SLSBridgedWindowManagementOperationResultClass = SLSBridgedWindowManagementOperationResultClass{class: objc.GetClass("SLSBridgedWindowManagementOperationResult")}
	})
	return _SLSBridgedWindowManagementOperationResultClass
}

// GetSLSBridgedWindowManagementOperationResultClass returns the class object for SLSBridgedWindowManagementOperationResult.
func GetSLSBridgedWindowManagementOperationResultClass() SLSBridgedWindowManagementOperationResultClass {
	return getSLSBridgedWindowManagementOperationResultClass()
}

type SLSBridgedWindowManagementOperationResultClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedWindowManagementOperationResultClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedWindowManagementOperationResultClass) Alloc() SLSBridgedWindowManagementOperationResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationResult](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedWindowManagementOperationResult._init]
//   - [SLSBridgedWindowManagementOperationResult.EncodeWithCoder]
//   - [SLSBridgedWindowManagementOperationResult.InitWithCoder]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationResult
type SLSBridgedWindowManagementOperationResult struct {
	objectivec.Object
}

// SLSBridgedWindowManagementOperationResultFromID constructs a [SLSBridgedWindowManagementOperationResult] from an objc.ID.
func SLSBridgedWindowManagementOperationResultFromID(id objc.ID) SLSBridgedWindowManagementOperationResult {
	return SLSBridgedWindowManagementOperationResult{objectivec.Object{ID: id}}
}

// Ensure SLSBridgedWindowManagementOperationResult implements ISLSBridgedWindowManagementOperationResult.
var _ ISLSBridgedWindowManagementOperationResult = SLSBridgedWindowManagementOperationResult{}

// An interface definition for the [SLSBridgedWindowManagementOperationResult] class.
//
// # Methods
//
//   - [ISLSBridgedWindowManagementOperationResult._init]
//   - [ISLSBridgedWindowManagementOperationResult.EncodeWithCoder]
//   - [ISLSBridgedWindowManagementOperationResult.InitWithCoder]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationResult
type ISLSBridgedWindowManagementOperationResult interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
	EncodeWithCoder(coder foundation.INSCoder)
	InitWithCoder(coder foundation.INSCoder) SLSBridgedWindowManagementOperationResult
}

// Init initializes the instance.
func (s SLSBridgedWindowManagementOperationResult) Init() SLSBridgedWindowManagementOperationResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationResult](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedWindowManagementOperationResult) Autorelease() SLSBridgedWindowManagementOperationResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationResult](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedWindowManagementOperationResult creates a new SLSBridgedWindowManagementOperationResult instance.
func NewSLSBridgedWindowManagementOperationResult() SLSBridgedWindowManagementOperationResult {
	class := getSLSBridgedWindowManagementOperationResultClass()
	rv := objc.Send[SLSBridgedWindowManagementOperationResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationResult/initWithCoder:
func NewSLSBridgedWindowManagementOperationResultWithCoder(coder objectivec.IObject) SLSBridgedWindowManagementOperationResult {
	instance := getSLSBridgedWindowManagementOperationResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedWindowManagementOperationResultFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationResult/_init
func (s SLSBridgedWindowManagementOperationResult) _init() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationResult/encodeWithCoder:
func (s SLSBridgedWindowManagementOperationResult) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationResult/initWithCoder:
func (s SLSBridgedWindowManagementOperationResult) InitWithCoder(coder foundation.INSCoder) SLSBridgedWindowManagementOperationResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationResult](s.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationResult/supportsSecureCoding
func (_SLSBridgedWindowManagementOperationResultClass SLSBridgedWindowManagementOperationResultClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_SLSBridgedWindowManagementOperationResultClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}
