// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SLSBridgedWindowManagementOperationSizeResult] class.
var (
	_SLSBridgedWindowManagementOperationSizeResultClass     SLSBridgedWindowManagementOperationSizeResultClass
	_SLSBridgedWindowManagementOperationSizeResultClassOnce sync.Once
)

func getSLSBridgedWindowManagementOperationSizeResultClass() SLSBridgedWindowManagementOperationSizeResultClass {
	_SLSBridgedWindowManagementOperationSizeResultClassOnce.Do(func() {
		_SLSBridgedWindowManagementOperationSizeResultClass = SLSBridgedWindowManagementOperationSizeResultClass{class: objc.GetClass("SLSBridgedWindowManagementOperationSizeResult")}
	})
	return _SLSBridgedWindowManagementOperationSizeResultClass
}

// GetSLSBridgedWindowManagementOperationSizeResultClass returns the class object for SLSBridgedWindowManagementOperationSizeResult.
func GetSLSBridgedWindowManagementOperationSizeResultClass() SLSBridgedWindowManagementOperationSizeResultClass {
	return getSLSBridgedWindowManagementOperationSizeResultClass()
}

type SLSBridgedWindowManagementOperationSizeResultClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SLSBridgedWindowManagementOperationSizeResultClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SLSBridgedWindowManagementOperationSizeResultClass) Alloc() SLSBridgedWindowManagementOperationSizeResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationSizeResult](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SLSBridgedWindowManagementOperationSizeResult.Size]
//   - [SLSBridgedWindowManagementOperationSizeResult.InitWithSize]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationSizeResult
type SLSBridgedWindowManagementOperationSizeResult struct {
	SLSBridgedWindowManagementOperationResult
}

// SLSBridgedWindowManagementOperationSizeResultFromID constructs a [SLSBridgedWindowManagementOperationSizeResult] from an objc.ID.
func SLSBridgedWindowManagementOperationSizeResultFromID(id objc.ID) SLSBridgedWindowManagementOperationSizeResult {
	return SLSBridgedWindowManagementOperationSizeResult{SLSBridgedWindowManagementOperationResult: SLSBridgedWindowManagementOperationResultFromID(id)}
}

// Ensure SLSBridgedWindowManagementOperationSizeResult implements ISLSBridgedWindowManagementOperationSizeResult.
var _ ISLSBridgedWindowManagementOperationSizeResult = SLSBridgedWindowManagementOperationSizeResult{}

// An interface definition for the [SLSBridgedWindowManagementOperationSizeResult] class.
//
// # Methods
//
//   - [ISLSBridgedWindowManagementOperationSizeResult.Size]
//   - [ISLSBridgedWindowManagementOperationSizeResult.InitWithSize]
//
// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationSizeResult
type ISLSBridgedWindowManagementOperationSizeResult interface {
	ISLSBridgedWindowManagementOperationResult

	// Topic: Methods

	Size() corefoundation.CGSize
	InitWithSize(size corefoundation.CGSize) SLSBridgedWindowManagementOperationSizeResult
}

// Init initializes the instance.
func (s SLSBridgedWindowManagementOperationSizeResult) Init() SLSBridgedWindowManagementOperationSizeResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationSizeResult](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SLSBridgedWindowManagementOperationSizeResult) Autorelease() SLSBridgedWindowManagementOperationSizeResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationSizeResult](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSLSBridgedWindowManagementOperationSizeResult creates a new SLSBridgedWindowManagementOperationSizeResult instance.
func NewSLSBridgedWindowManagementOperationSizeResult() SLSBridgedWindowManagementOperationSizeResult {
	class := getSLSBridgedWindowManagementOperationSizeResultClass()
	rv := objc.Send[SLSBridgedWindowManagementOperationSizeResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationSizeResult/initWithCoder:
func NewSLSBridgedWindowManagementOperationSizeResultWithCoder(coder objectivec.IObject) SLSBridgedWindowManagementOperationSizeResult {
	instance := getSLSBridgedWindowManagementOperationSizeResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SLSBridgedWindowManagementOperationSizeResultFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationSizeResult/initWithSize:
func NewSLSBridgedWindowManagementOperationSizeResultWithSize(size corefoundation.CGSize) SLSBridgedWindowManagementOperationSizeResult {
	instance := getSLSBridgedWindowManagementOperationSizeResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSize:"), size)
	return SLSBridgedWindowManagementOperationSizeResultFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationSizeResult/initWithSize:
func (s SLSBridgedWindowManagementOperationSizeResult) InitWithSize(size corefoundation.CGSize) SLSBridgedWindowManagementOperationSizeResult {
	rv := objc.Send[SLSBridgedWindowManagementOperationSizeResult](s.ID, objc.Sel("initWithSize:"), size)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SLSBridgedWindowManagementOperationSizeResult/size
func (s SLSBridgedWindowManagementOperationSizeResult) Size() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](s.ID, objc.Sel("size"))
	return corefoundation.CGSize(rv)
}
