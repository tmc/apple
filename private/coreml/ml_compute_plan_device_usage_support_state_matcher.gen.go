// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
	rv := objc.SendIfResponds[MLComputePlanDeviceUsageSupportStateMatcher](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLComputePlanDeviceUsageSupportStateMatcher.MatchingSupportStateForValidationMessage]
//   - [MLComputePlanDeviceUsageSupportStateMatcher.SupportStatePatterns]
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
type IMLComputePlanDeviceUsageSupportStateMatcher interface {
	objectivec.IObject

	// Topic: Methods

	MatchingSupportStateForValidationMessage(message objectivec.IObject) int64
	SupportStatePatterns() foundation.INSArray
}

// Init initializes the instance.
func (m MLComputePlanDeviceUsageSupportStateMatcher) Init() MLComputePlanDeviceUsageSupportStateMatcher {
	rv := objc.SendIfResponds[MLComputePlanDeviceUsageSupportStateMatcher](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLComputePlanDeviceUsageSupportStateMatcher) Autorelease() MLComputePlanDeviceUsageSupportStateMatcher {
	rv := objc.SendIfResponds[MLComputePlanDeviceUsageSupportStateMatcher](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLComputePlanDeviceUsageSupportStateMatcher creates a new MLComputePlanDeviceUsageSupportStateMatcher instance.
func NewMLComputePlanDeviceUsageSupportStateMatcher() MLComputePlanDeviceUsageSupportStateMatcher {
	class := getMLComputePlanDeviceUsageSupportStateMatcherClass()
	rv := objc.SendIfResponds[MLComputePlanDeviceUsageSupportStateMatcher](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (m MLComputePlanDeviceUsageSupportStateMatcher) MatchingSupportStateForValidationMessage(message objectivec.IObject) int64 {
	rv := objc.SendIfResponds[int64](m.ID, objc.Sel("matchingSupportStateForValidationMessage:"), message)
	return rv
}

func (_MLComputePlanDeviceUsageSupportStateMatcherClass MLComputePlanDeviceUsageSupportStateMatcherClass) SharedInstance() MLComputePlanDeviceUsageSupportStateMatcher {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLComputePlanDeviceUsageSupportStateMatcherClass.class), objc.Sel("sharedInstance"))
	return MLComputePlanDeviceUsageSupportStateMatcherFromID(rv)
}

func (m MLComputePlanDeviceUsageSupportStateMatcher) SupportStatePatterns() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("supportStatePatterns"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
