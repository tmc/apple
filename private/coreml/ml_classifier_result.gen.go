// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLClassifierResult] class.
var (
	_MLClassifierResultClass     MLClassifierResultClass
	_MLClassifierResultClassOnce sync.Once
)

func getMLClassifierResultClass() MLClassifierResultClass {
	_MLClassifierResultClassOnce.Do(func() {
		_MLClassifierResultClass = MLClassifierResultClass{class: objc.GetClass("MLClassifierResult")}
	})
	return _MLClassifierResultClass
}

// GetMLClassifierResultClass returns the class object for MLClassifierResult.
func GetMLClassifierResultClass() MLClassifierResultClass {
	return getMLClassifierResultClass()
}

type MLClassifierResultClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLClassifierResultClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLClassifierResultClass) Alloc() MLClassifierResult {
	rv := objc.Send[MLClassifierResult](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLClassifierResult.AdditionalFeatures]
//   - [MLClassifierResult.AsFeatureDictionaryWithPredictedClassDescriptionClassProbabilityDescription]
//   - [MLClassifierResult.ClassProbability]
//   - [MLClassifierResult.PredictedClass]
//   - [MLClassifierResult.PredictedClassFeatureType]
//   - [MLClassifierResult.InitWithClassProbabilityAdditionalFeaturesClassLabelOfMaxProbability]
//   - [MLClassifierResult.InitWithIntClassProbabilityClassFeatureTypeAdditionalFeatures]
//   - [MLClassifierResult.InitWithStringClassProbabilityClassFeatureTypeAdditionalFeatures]
//
// See: https://developer.apple.com/documentation/CoreML/MLClassifierResult
type MLClassifierResult struct {
	objectivec.Object
}

// MLClassifierResultFromID constructs a [MLClassifierResult] from an objc.ID.
func MLClassifierResultFromID(id objc.ID) MLClassifierResult {
	return MLClassifierResult{objectivec.Object{ID: id}}
}

// Ensure MLClassifierResult implements IMLClassifierResult.
var _ IMLClassifierResult = MLClassifierResult{}

// An interface definition for the [MLClassifierResult] class.
//
// # Methods
//
//   - [IMLClassifierResult.AdditionalFeatures]
//   - [IMLClassifierResult.AsFeatureDictionaryWithPredictedClassDescriptionClassProbabilityDescription]
//   - [IMLClassifierResult.ClassProbability]
//   - [IMLClassifierResult.PredictedClass]
//   - [IMLClassifierResult.PredictedClassFeatureType]
//   - [IMLClassifierResult.InitWithClassProbabilityAdditionalFeaturesClassLabelOfMaxProbability]
//   - [IMLClassifierResult.InitWithIntClassProbabilityClassFeatureTypeAdditionalFeatures]
//   - [IMLClassifierResult.InitWithStringClassProbabilityClassFeatureTypeAdditionalFeatures]
//
// See: https://developer.apple.com/documentation/CoreML/MLClassifierResult
type IMLClassifierResult interface {
	objectivec.IObject

	// Topic: Methods

	AdditionalFeatures() objectivec.IObject
	AsFeatureDictionaryWithPredictedClassDescriptionClassProbabilityDescription(description objectivec.IObject, description2 objectivec.IObject) objectivec.IObject
	ClassProbability() foundation.INSDictionary
	PredictedClass() IMLFeatureValue
	PredictedClassFeatureType() int64
	InitWithClassProbabilityAdditionalFeaturesClassLabelOfMaxProbability(probability objectivec.IObject, features objectivec.IObject, probability2 objectivec.IObject) MLClassifierResult
	InitWithIntClassProbabilityClassFeatureTypeAdditionalFeatures(probability objectivec.IObject, type_ int64, features objectivec.IObject) MLClassifierResult
	InitWithStringClassProbabilityClassFeatureTypeAdditionalFeatures(probability objectivec.IObject, type_ int64, features objectivec.IObject) MLClassifierResult
}

// Init initializes the instance.
func (c MLClassifierResult) Init() MLClassifierResult {
	rv := objc.Send[MLClassifierResult](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MLClassifierResult) Autorelease() MLClassifierResult {
	rv := objc.Send[MLClassifierResult](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLClassifierResult creates a new MLClassifierResult instance.
func NewMLClassifierResult() MLClassifierResult {
	class := getMLClassifierResultClass()
	rv := objc.Send[MLClassifierResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLClassifierResult/initWithClassProbability:additionalFeatures:classLabelOfMaxProbability:
func NewClassifierResultWithClassProbabilityAdditionalFeaturesClassLabelOfMaxProbability(probability objectivec.IObject, features objectivec.IObject, probability2 objectivec.IObject) MLClassifierResult {
	instance := getMLClassifierResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithClassProbability:additionalFeatures:classLabelOfMaxProbability:"), probability, features, probability2)
	return MLClassifierResultFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLClassifierResult/initWithIntClassProbability:classFeatureType:additionalFeatures:
func NewClassifierResultWithIntClassProbabilityClassFeatureTypeAdditionalFeatures(probability objectivec.IObject, type_ int64, features objectivec.IObject) MLClassifierResult {
	instance := getMLClassifierResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIntClassProbability:classFeatureType:additionalFeatures:"), probability, type_, features)
	return MLClassifierResultFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLClassifierResult/initWithStringClassProbability:classFeatureType:additionalFeatures:
func NewClassifierResultWithStringClassProbabilityClassFeatureTypeAdditionalFeatures(probability objectivec.IObject, type_ int64, features objectivec.IObject) MLClassifierResult {
	instance := getMLClassifierResultClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithStringClassProbability:classFeatureType:additionalFeatures:"), probability, type_, features)
	return MLClassifierResultFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLClassifierResult/asFeatureDictionaryWithPredictedClassDescription:classProbabilityDescription:
func (c MLClassifierResult) AsFeatureDictionaryWithPredictedClassDescriptionClassProbabilityDescription(description objectivec.IObject, description2 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("asFeatureDictionaryWithPredictedClassDescription:classProbabilityDescription:"), description, description2)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLClassifierResult/initWithClassProbability:additionalFeatures:classLabelOfMaxProbability:
func (c MLClassifierResult) InitWithClassProbabilityAdditionalFeaturesClassLabelOfMaxProbability(probability objectivec.IObject, features objectivec.IObject, probability2 objectivec.IObject) MLClassifierResult {
	rv := objc.Send[MLClassifierResult](c.ID, objc.Sel("initWithClassProbability:additionalFeatures:classLabelOfMaxProbability:"), probability, features, probability2)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLClassifierResult/initWithIntClassProbability:classFeatureType:additionalFeatures:
func (c MLClassifierResult) InitWithIntClassProbabilityClassFeatureTypeAdditionalFeatures(probability objectivec.IObject, type_ int64, features objectivec.IObject) MLClassifierResult {
	rv := objc.Send[MLClassifierResult](c.ID, objc.Sel("initWithIntClassProbability:classFeatureType:additionalFeatures:"), probability, type_, features)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLClassifierResult/initWithStringClassProbability:classFeatureType:additionalFeatures:
func (c MLClassifierResult) InitWithStringClassProbabilityClassFeatureTypeAdditionalFeatures(probability objectivec.IObject, type_ int64, features objectivec.IObject) MLClassifierResult {
	rv := objc.Send[MLClassifierResult](c.ID, objc.Sel("initWithStringClassProbability:classFeatureType:additionalFeatures:"), probability, type_, features)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLClassifierResult/resultWithClassProbability:additionalFeatures:classLabelOfMaxProbability:
func (_MLClassifierResultClass MLClassifierResultClass) ResultWithClassProbabilityAdditionalFeaturesClassLabelOfMaxProbability(probability objectivec.IObject, features objectivec.IObject, probability2 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLClassifierResultClass.class), objc.Sel("resultWithClassProbability:additionalFeatures:classLabelOfMaxProbability:"), probability, features, probability2)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLClassifierResult/resultWithIntClassProbability:
func (_MLClassifierResultClass MLClassifierResultClass) ResultWithIntClassProbability(probability objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLClassifierResultClass.class), objc.Sel("resultWithIntClassProbability:"), probability)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLClassifierResult/resultWithIntClassProbability:additionalFeatures:
func (_MLClassifierResultClass MLClassifierResultClass) ResultWithIntClassProbabilityAdditionalFeatures(probability objectivec.IObject, features objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLClassifierResultClass.class), objc.Sel("resultWithIntClassProbability:additionalFeatures:"), probability, features)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLClassifierResult/resultWithStringClassProbability:
func (_MLClassifierResultClass MLClassifierResultClass) ResultWithStringClassProbability(probability objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLClassifierResultClass.class), objc.Sel("resultWithStringClassProbability:"), probability)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLClassifierResult/resultWithStringClassProbability:additionalFeatures:
func (_MLClassifierResultClass MLClassifierResultClass) ResultWithStringClassProbabilityAdditionalFeatures(probability objectivec.IObject, features objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLClassifierResultClass.class), objc.Sel("resultWithStringClassProbability:additionalFeatures:"), probability, features)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLClassifierResult/additionalFeatures
func (c MLClassifierResult) AdditionalFeatures() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("additionalFeatures"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLClassifierResult/classProbability
func (c MLClassifierResult) ClassProbability() foundation.INSDictionary {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("classProbability"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLClassifierResult/predictedClass
func (c MLClassifierResult) PredictedClass() IMLFeatureValue {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("predictedClass"))
	return MLFeatureValueFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLClassifierResult/predictedClassFeatureType
func (c MLClassifierResult) PredictedClassFeatureType() int64 {
	rv := objc.Send[int64](c.ID, objc.Sel("predictedClassFeatureType"))
	return rv
}
