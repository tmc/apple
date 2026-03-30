// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLRegressorResult] class.
var (
	_MLRegressorResultClass     MLRegressorResultClass
	_MLRegressorResultClassOnce sync.Once
)

func getMLRegressorResultClass() MLRegressorResultClass {
	_MLRegressorResultClassOnce.Do(func() {
		_MLRegressorResultClass = MLRegressorResultClass{class: objc.GetClass("MLRegressorResult")}
	})
	return _MLRegressorResultClass
}

// GetMLRegressorResultClass returns the class object for MLRegressorResult.
func GetMLRegressorResultClass() MLRegressorResultClass {
	return getMLRegressorResultClass()
}

type MLRegressorResultClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLRegressorResultClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLRegressorResultClass) Alloc() MLRegressorResult {
	rv := objc.Send[MLRegressorResult](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLRegressorResult.AdditionalFeatures]
//   - [MLRegressorResult.AsFeatureDictionaryWithPredictedValueDescription]
//   - [MLRegressorResult.PredictedValue]
//   - [MLRegressorResult.InitWithValueAdditionalFeatures]
//
// See: https://developer.apple.com/documentation/CoreML/MLRegressorResult
type MLRegressorResult struct {
	objectivec.Object
}

// MLRegressorResultFromID constructs a [MLRegressorResult] from an objc.ID.
func MLRegressorResultFromID(id objc.ID) MLRegressorResult {
	return MLRegressorResult{objectivec.Object{ID: id}}
}

// Ensure MLRegressorResult implements IMLRegressorResult.
var _ IMLRegressorResult = MLRegressorResult{}

// An interface definition for the [MLRegressorResult] class.
//
// # Methods
//
//   - [IMLRegressorResult.AdditionalFeatures]
//   - [IMLRegressorResult.AsFeatureDictionaryWithPredictedValueDescription]
//   - [IMLRegressorResult.PredictedValue]
//   - [IMLRegressorResult.InitWithValueAdditionalFeatures]
//
// See: https://developer.apple.com/documentation/CoreML/MLRegressorResult
type IMLRegressorResult interface {
	objectivec.IObject

	// Topic: Methods

	AdditionalFeatures() objectivec.IObject
	AsFeatureDictionaryWithPredictedValueDescription(description objectivec.IObject) objectivec.IObject
	PredictedValue() IMLMultiArray
	InitWithValueAdditionalFeatures(value objectivec.IObject, features objectivec.IObject) MLRegressorResult
}

// Init initializes the instance.
func (r MLRegressorResult) Init() MLRegressorResult {
	rv := objc.Send[MLRegressorResult](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MLRegressorResult) Autorelease() MLRegressorResult {
	rv := objc.Send[MLRegressorResult](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLRegressorResult creates a new MLRegressorResult instance.
func NewMLRegressorResult() MLRegressorResult {
	class := getMLRegressorResultClass()
	rv := objc.Send[MLRegressorResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLRegressorResult/initWithValue:additionalFeatures:
func NewRegressorResultWithValueAdditionalFeatures(value objectivec.IObject, features objectivec.IObject) MLRegressorResult {
	instance := getMLRegressorResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithValue:additionalFeatures:"), value, features)
	return MLRegressorResultFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLRegressorResult/asFeatureDictionaryWithPredictedValueDescription:
func (r MLRegressorResult) AsFeatureDictionaryWithPredictedValueDescription(description objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("asFeatureDictionaryWithPredictedValueDescription:"), description)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLRegressorResult/initWithValue:additionalFeatures:
func (r MLRegressorResult) InitWithValueAdditionalFeatures(value objectivec.IObject, features objectivec.IObject) MLRegressorResult {
	rv := objc.Send[MLRegressorResult](r.ID, objc.Sel("initWithValue:additionalFeatures:"), value, features)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLRegressorResult/resultWithValue:
func (_MLRegressorResultClass MLRegressorResultClass) ResultWithValue(value objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLRegressorResultClass.class), objc.Sel("resultWithValue:"), value)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLRegressorResult/resultWithValue:additionalFeatures:
func (_MLRegressorResultClass MLRegressorResultClass) ResultWithValueAdditionalFeatures(value objectivec.IObject, features objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLRegressorResultClass.class), objc.Sel("resultWithValue:additionalFeatures:"), value, features)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLRegressorResult/additionalFeatures
func (r MLRegressorResult) AdditionalFeatures() objectivec.IObject {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("additionalFeatures"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLRegressorResult/predictedValue
func (r MLRegressorResult) PredictedValue() IMLMultiArray {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("predictedValue"))
	return MLMultiArrayFromID(objc.ID(rv))
}
