// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedWindowManagementOperationAffineTransformWithOptionsResult] class.
var (
	_SLSBridgedWindowManagementOperationAffineTransformWithOptionsResultClass     SLSBridgedWindowManagementOperationAffineTransformWithOptionsResultClass
	_SLSBridgedWindowManagementOperationAffineTransformWithOptionsResultClassOnce sync.Once
)

func getSLSBridgedWindowManagementOperationAffineTransformWithOptionsResultClass() SLSBridgedWindowManagementOperationAffineTransformWithOptionsResultClass {
	_SLSBridgedWindowManagementOperationAffineTransformWithOptionsResultClassOnce.Do(func() {
		_SLSBridgedWindowManagementOperationAffineTransformWithOptionsResultClass = SLSBridgedWindowManagementOperationAffineTransformWithOptionsResultClass{class: objc.GetClass("SLSBridgedWindowManagementOperationAffineTransformWithOptionsResult")}
	})
	return _SLSBridgedWindowManagementOperationAffineTransformWithOptionsResultClass
}

// GetSLSBridgedWindowManagementOperationAffineTransformWithOptionsResultClass returns the class object for SLSBridgedWindowManagementOperationAffineTransformWithOptionsResult.
func GetSLSBridgedWindowManagementOperationAffineTransformWithOptionsResultClass() SLSBridgedWindowManagementOperationAffineTransformWithOptionsResultClass {
	return getSLSBridgedWindowManagementOperationAffineTransformWithOptionsResultClass()
}

type SLSBridgedWindowManagementOperationAffineTransformWithOptionsResultClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedWindowManagementOperationAffineTransformWithOptionsResultClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedWindowManagementOperationAffineTransformWithOptionsResultClass) Alloc() SLSBridgedWindowManagementOperationAffineTransformWithOptionsResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationAffineTransformWithOptionsResult](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedWindowManagementOperationAffineTransformWithOptionsResult.AffineTransform]
//   - [SLSBridgedWindowManagementOperationAffineTransformWithOptionsResult.Options]
//   - [SLSBridgedWindowManagementOperationAffineTransformWithOptionsResult.InitWithAffineTransformOptions]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationAffineTransformWithOptionsResult
type SLSBridgedWindowManagementOperationAffineTransformWithOptionsResult struct {
	SLSBridgedWindowManagementOperationResult
}

// SLSBridgedWindowManagementOperationAffineTransformWithOptionsResultFromID constructs a [SLSBridgedWindowManagementOperationAffineTransformWithOptionsResult] from an objc.ID.
func SLSBridgedWindowManagementOperationAffineTransformWithOptionsResultFromID(id objc.ID) SLSBridgedWindowManagementOperationAffineTransformWithOptionsResult {
	return SLSBridgedWindowManagementOperationAffineTransformWithOptionsResult{SLSBridgedWindowManagementOperationResult: SLSBridgedWindowManagementOperationResultFromID(id)}
}

// Ensure SLSBridgedWindowManagementOperationAffineTransformWithOptionsResult implements ISLSBridgedWindowManagementOperationAffineTransformWithOptionsResult.
var _ ISLSBridgedWindowManagementOperationAffineTransformWithOptionsResult = SLSBridgedWindowManagementOperationAffineTransformWithOptionsResult{}

// An interface definition for the [SLSBridgedWindowManagementOperationAffineTransformWithOptionsResult] class.
//
// # Methods
//
//   - [ISLSBridgedWindowManagementOperationAffineTransformWithOptionsResult.AffineTransform]
//   - [ISLSBridgedWindowManagementOperationAffineTransformWithOptionsResult.Options]
//   - [ISLSBridgedWindowManagementOperationAffineTransformWithOptionsResult.InitWithAffineTransformOptions]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationAffineTransformWithOptionsResult
type ISLSBridgedWindowManagementOperationAffineTransformWithOptionsResult interface {
	ISLSBridgedWindowManagementOperationResult

	// Topic: Methods

	AffineTransform() corefoundation.CGAffineTransform
	Options() uint32
	InitWithAffineTransformOptions(transform corefoundation.CGAffineTransform, options uint32) SLSBridgedWindowManagementOperationAffineTransformWithOptionsResult
}

// Init initializes the instance.
func (s SLSBridgedWindowManagementOperationAffineTransformWithOptionsResult) Init() SLSBridgedWindowManagementOperationAffineTransformWithOptionsResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationAffineTransformWithOptionsResult](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedWindowManagementOperationAffineTransformWithOptionsResult) Autorelease() SLSBridgedWindowManagementOperationAffineTransformWithOptionsResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationAffineTransformWithOptionsResult](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedWindowManagementOperationAffineTransformWithOptionsResult creates a new SLSBridgedWindowManagementOperationAffineTransformWithOptionsResult instance.
func NewSLSBridgedWindowManagementOperationAffineTransformWithOptionsResult() SLSBridgedWindowManagementOperationAffineTransformWithOptionsResult {
	class := getSLSBridgedWindowManagementOperationAffineTransformWithOptionsResultClass()
	rv := objc.Send[SLSBridgedWindowManagementOperationAffineTransformWithOptionsResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationAffineTransformWithOptionsResult/initWithAffineTransform:options:
func NewSLSBridgedWindowManagementOperationAffineTransformWithOptionsResultWithAffineTransformOptions(transform corefoundation.CGAffineTransform, options uint32) SLSBridgedWindowManagementOperationAffineTransformWithOptionsResult {
	instance := getSLSBridgedWindowManagementOperationAffineTransformWithOptionsResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAffineTransform:options:"), transform, options)
	return SLSBridgedWindowManagementOperationAffineTransformWithOptionsResultFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationAffineTransformWithOptionsResult/initWithCoder:
func NewSLSBridgedWindowManagementOperationAffineTransformWithOptionsResultWithCoder(coder objectivec.IObject) SLSBridgedWindowManagementOperationAffineTransformWithOptionsResult {
	instance := getSLSBridgedWindowManagementOperationAffineTransformWithOptionsResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedWindowManagementOperationAffineTransformWithOptionsResultFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationAffineTransformWithOptionsResult/initWithAffineTransform:options:
func (s SLSBridgedWindowManagementOperationAffineTransformWithOptionsResult) InitWithAffineTransformOptions(transform corefoundation.CGAffineTransform, options uint32) SLSBridgedWindowManagementOperationAffineTransformWithOptionsResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationAffineTransformWithOptionsResult](s.ID, objc.Sel("initWithAffineTransform:options:"), transform, options)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationAffineTransformWithOptionsResult/affineTransform
func (s SLSBridgedWindowManagementOperationAffineTransformWithOptionsResult) AffineTransform() corefoundation.CGAffineTransform {
	rv := objc.Send[corefoundation.CGAffineTransform](s.ID, objc.Sel("affineTransform"))
	return corefoundation.CGAffineTransform(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationAffineTransformWithOptionsResult/options
func (s SLSBridgedWindowManagementOperationAffineTransformWithOptionsResult) Options() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("options"))
	return rv
}
