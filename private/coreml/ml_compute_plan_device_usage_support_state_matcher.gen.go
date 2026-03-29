// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLComputePlanDeviceUsageSupportStateMatcher] class.
var (
	_MLComputePlanDeviceUsageSupportStateMatcherClass     MLComputePlanDeviceUsageSupportStateMatcherClass
	_MLComputePlanDeviceUsageSupportStateMatcherClassOnce sync.Once
)

func getMLComputePlanDeviceUsageSupportStateMatcherClass() MLComputePlanDeviceUsageSupportStateMatcherClass {
	_MLComputePlanDeviceUsageSupportStateMatcherClassOnce.Do(func() {
		_MLComputePlanDeviceUsageSupportStateMatcherClass = MLComputePlanDeviceUsageSupportStateMatcherClass{class: objc.GetClass("MLComputePlanDeviceUsageSupportStateMatcher")}
	})
	return _MLComputePlanDeviceUsageSupportStateMatcherClass
}

// GetMLComputePlanDeviceUsageSupportStateMatcherClass returns the class object for MLComputePlanDeviceUsageSupportStateMatcher.
func GetMLComputePlanDeviceUsageSupportStateMatcherClass() MLComputePlanDeviceUsageSupportStateMatcherClass {
	return getMLComputePlanDeviceUsageSupportStateMatcherClass()
}

type MLComputePlanDeviceUsageSupportStateMatcherClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLComputePlanDeviceUsageSupportStateMatcherClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLComputePlanDeviceUsageSupportStateMatcherClass) Alloc() MLComputePlanDeviceUsageSupportStateMatcher {
	rv := objc.Send[MLComputePlanDeviceUsageSupportStateMatcher](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLComputePlanDeviceUsageSupportStateMatcher.MatchingSupportStateForValidationMessage]
//   - [MLComputePlanDeviceUsageSupportStateMatcher.SupportStatePatterns]
// See: https://developer.apple.com/documentation/CoreML/MLComputePlanDeviceUsageSupportStateMatcher
type MLComputePlanDeviceUsageSupportStateMatcher struct {
	objectivec.Object
}

// MLComputePlanDeviceUsageSupportStateMatcherFromID constructs a [MLComputePlanDeviceUsageSupportStateMatcher] from an objc.ID.
func MLComputePlanDeviceUsageSupportStateMatcherFromID(id objc.ID) MLComputePlanDeviceUsageSupportStateMatcher {
	return MLComputePlanDeviceUsageSupportStateMatcher{objectivec.Object{ID: id}}
}
// Ensure MLComputePlanDeviceUsageSupportStateMatcher implements IMLComputePlanDeviceUsageSupportStateMatcher.
var _ IMLComputePlanDeviceUsageSupportStateMatcher = MLComputePlanDeviceUsageSupportStateMatcher{}

// An interface definition for the [MLComputePlanDeviceUsageSupportStateMatcher] class.
//
// # Methods
//
//   - [IMLComputePlanDeviceUsageSupportStateMatcher.MatchingSupportStateForValidationMessage]
//   - [IMLComputePlanDeviceUsageSupportStateMatcher.SupportStatePatterns]
//
// See: https://developer.apple.com/documentation/CoreML/MLComputePlanDeviceUsageSupportStateMatcher
type IMLComputePlanDeviceUsageSupportStateMatcher interface {
	objectivec.IObject

	// Topic: Methods

	MatchingSupportStateForValidationMessage(message objectivec.IObject) int64
	SupportStatePatterns() foundation.INSArray
}

// Init initializes the instance.
func (c MLComputePlanDeviceUsageSupportStateMatcher) Init() MLComputePlanDeviceUsageSupportStateMatcher {
	rv := objc.Send[MLComputePlanDeviceUsageSupportStateMatcher](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MLComputePlanDeviceUsageSupportStateMatcher) Autorelease() MLComputePlanDeviceUsageSupportStateMatcher {
	rv := objc.Send[MLComputePlanDeviceUsageSupportStateMatcher](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLComputePlanDeviceUsageSupportStateMatcher creates a new MLComputePlanDeviceUsageSupportStateMatcher instance.
func NewMLComputePlanDeviceUsageSupportStateMatcher() MLComputePlanDeviceUsageSupportStateMatcher {
	class := getMLComputePlanDeviceUsageSupportStateMatcherClass()
	rv := objc.Send[MLComputePlanDeviceUsageSupportStateMatcher](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLComputePlanDeviceUsageSupportStateMatcher/matchingSupportStateForValidationMessage:
func (c MLComputePlanDeviceUsageSupportStateMatcher) MatchingSupportStateForValidationMessage(message objectivec.IObject) int64 {
	rv := objc.Send[int64](c.ID, objc.Sel("matchingSupportStateForValidationMessage:"), message)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLComputePlanDeviceUsageSupportStateMatcher/sharedInstance
func (_MLComputePlanDeviceUsageSupportStateMatcherClass MLComputePlanDeviceUsageSupportStateMatcherClass) SharedInstance() MLComputePlanDeviceUsageSupportStateMatcher {
	rv := objc.Send[objc.ID](objc.ID(_MLComputePlanDeviceUsageSupportStateMatcherClass.class), objc.Sel("sharedInstance"))
	return MLComputePlanDeviceUsageSupportStateMatcherFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLComputePlanDeviceUsageSupportStateMatcher/supportStatePatterns
func (c MLComputePlanDeviceUsageSupportStateMatcher) SupportStatePatterns() foundation.INSArray {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("supportStatePatterns"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

