// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLComputePlanDeviceUsageSupportStatePattern] class.
var (
	_MLComputePlanDeviceUsageSupportStatePatternClass     MLComputePlanDeviceUsageSupportStatePatternClass
	_MLComputePlanDeviceUsageSupportStatePatternClassOnce sync.Once
)

func getMLComputePlanDeviceUsageSupportStatePatternClass() MLComputePlanDeviceUsageSupportStatePatternClass {
	_MLComputePlanDeviceUsageSupportStatePatternClassOnce.Do(func() {
		_MLComputePlanDeviceUsageSupportStatePatternClass = MLComputePlanDeviceUsageSupportStatePatternClass{class: objc.GetClass("MLComputePlanDeviceUsageSupportStatePattern")}
	})
	return _MLComputePlanDeviceUsageSupportStatePatternClass
}

// GetMLComputePlanDeviceUsageSupportStatePatternClass returns the class object for MLComputePlanDeviceUsageSupportStatePattern.
func GetMLComputePlanDeviceUsageSupportStatePatternClass() MLComputePlanDeviceUsageSupportStatePatternClass {
	return getMLComputePlanDeviceUsageSupportStatePatternClass()
}

type MLComputePlanDeviceUsageSupportStatePatternClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLComputePlanDeviceUsageSupportStatePatternClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLComputePlanDeviceUsageSupportStatePatternClass) Alloc() MLComputePlanDeviceUsageSupportStatePattern {
	rv := objc.Send[MLComputePlanDeviceUsageSupportStatePattern](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLComputePlanDeviceUsageSupportStatePattern.IsMatchForValidationMessage]
//   - [MLComputePlanDeviceUsageSupportStatePattern.Regex]
//   - [MLComputePlanDeviceUsageSupportStatePattern.SupportState]
//   - [MLComputePlanDeviceUsageSupportStatePattern.InitWithPatternSupportState]
//
// See: https://developer.apple.com/documentation/CoreML/MLComputePlanDeviceUsageSupportStatePattern
type MLComputePlanDeviceUsageSupportStatePattern struct {
	objectivec.Object
}

// MLComputePlanDeviceUsageSupportStatePatternFromID constructs a [MLComputePlanDeviceUsageSupportStatePattern] from an objc.ID.
func MLComputePlanDeviceUsageSupportStatePatternFromID(id objc.ID) MLComputePlanDeviceUsageSupportStatePattern {
	return MLComputePlanDeviceUsageSupportStatePattern{objectivec.Object{ID: id}}
}

// Ensure MLComputePlanDeviceUsageSupportStatePattern implements IMLComputePlanDeviceUsageSupportStatePattern.
var _ IMLComputePlanDeviceUsageSupportStatePattern = MLComputePlanDeviceUsageSupportStatePattern{}

// An interface definition for the [MLComputePlanDeviceUsageSupportStatePattern] class.
//
// # Methods
//
//   - [IMLComputePlanDeviceUsageSupportStatePattern.IsMatchForValidationMessage]
//   - [IMLComputePlanDeviceUsageSupportStatePattern.Regex]
//   - [IMLComputePlanDeviceUsageSupportStatePattern.SupportState]
//   - [IMLComputePlanDeviceUsageSupportStatePattern.InitWithPatternSupportState]
//
// See: https://developer.apple.com/documentation/CoreML/MLComputePlanDeviceUsageSupportStatePattern
type IMLComputePlanDeviceUsageSupportStatePattern interface {
	objectivec.IObject

	// Topic: Methods

	IsMatchForValidationMessage(message objectivec.IObject) bool
	Regex() foundation.NSRegularExpression
	SupportState() int64
	InitWithPatternSupportState(pattern objectivec.IObject, state int64) MLComputePlanDeviceUsageSupportStatePattern
}

// Init initializes the instance.
func (c MLComputePlanDeviceUsageSupportStatePattern) Init() MLComputePlanDeviceUsageSupportStatePattern {
	rv := objc.Send[MLComputePlanDeviceUsageSupportStatePattern](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MLComputePlanDeviceUsageSupportStatePattern) Autorelease() MLComputePlanDeviceUsageSupportStatePattern {
	rv := objc.Send[MLComputePlanDeviceUsageSupportStatePattern](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLComputePlanDeviceUsageSupportStatePattern creates a new MLComputePlanDeviceUsageSupportStatePattern instance.
func NewMLComputePlanDeviceUsageSupportStatePattern() MLComputePlanDeviceUsageSupportStatePattern {
	class := getMLComputePlanDeviceUsageSupportStatePatternClass()
	rv := objc.Send[MLComputePlanDeviceUsageSupportStatePattern](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLComputePlanDeviceUsageSupportStatePattern/initWithPattern:supportState:
func NewComputePlanDeviceUsageSupportStatePatternWithPatternSupportState(pattern objectivec.IObject, state int64) MLComputePlanDeviceUsageSupportStatePattern {
	instance := getMLComputePlanDeviceUsageSupportStatePatternClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPattern:supportState:"), pattern, state)
	return MLComputePlanDeviceUsageSupportStatePatternFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLComputePlanDeviceUsageSupportStatePattern/isMatchForValidationMessage:
func (c MLComputePlanDeviceUsageSupportStatePattern) IsMatchForValidationMessage(message objectivec.IObject) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isMatchForValidationMessage:"), message)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLComputePlanDeviceUsageSupportStatePattern/initWithPattern:supportState:
func (c MLComputePlanDeviceUsageSupportStatePattern) InitWithPatternSupportState(pattern objectivec.IObject, state int64) MLComputePlanDeviceUsageSupportStatePattern {
	rv := objc.Send[MLComputePlanDeviceUsageSupportStatePattern](c.ID, objc.Sel("initWithPattern:supportState:"), pattern, state)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLComputePlanDeviceUsageSupportStatePattern/deviceUsageSupportStatePatternWithPattern:supportState:
func (_MLComputePlanDeviceUsageSupportStatePatternClass MLComputePlanDeviceUsageSupportStatePatternClass) DeviceUsageSupportStatePatternWithPatternSupportState(pattern objectivec.IObject, state int64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLComputePlanDeviceUsageSupportStatePatternClass.class), objc.Sel("deviceUsageSupportStatePatternWithPattern:supportState:"), pattern, state)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLComputePlanDeviceUsageSupportStatePattern/regex
func (c MLComputePlanDeviceUsageSupportStatePattern) Regex() foundation.NSRegularExpression {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("regex"))
	return foundation.NSRegularExpressionFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLComputePlanDeviceUsageSupportStatePattern/supportState
func (c MLComputePlanDeviceUsageSupportStatePattern) SupportState() int64 {
	rv := objc.Send[int64](c.ID, objc.Sel("supportState"))
	return rv
}
