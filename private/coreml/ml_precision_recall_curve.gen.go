// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLPrecisionRecallCurve] class.
var (
	_MLPrecisionRecallCurveClass     MLPrecisionRecallCurveClass
	_MLPrecisionRecallCurveClassOnce sync.Once
)

func getMLPrecisionRecallCurveClass() MLPrecisionRecallCurveClass {
	_MLPrecisionRecallCurveClassOnce.Do(func() {
		_MLPrecisionRecallCurveClass = MLPrecisionRecallCurveClass{class: objc.GetClass("MLPrecisionRecallCurve")}
	})
	return _MLPrecisionRecallCurveClass
}

// GetMLPrecisionRecallCurveClass returns the class object for MLPrecisionRecallCurve.
func GetMLPrecisionRecallCurveClass() MLPrecisionRecallCurveClass {
	return getMLPrecisionRecallCurveClass()
}

type MLPrecisionRecallCurveClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLPrecisionRecallCurveClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLPrecisionRecallCurveClass) Alloc() MLPrecisionRecallCurve {
	rv := objc.Send[MLPrecisionRecallCurve](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLPrecisionRecallCurve.EncodeWithCoder]
//   - [MLPrecisionRecallCurve.PrecisionConfidenceThresholds]
//   - [MLPrecisionRecallCurve.PrecisionValues]
//   - [MLPrecisionRecallCurve.RecallConfidenceThresholds]
//   - [MLPrecisionRecallCurve.RecallValues]
//   - [MLPrecisionRecallCurve.InitWithCoder]
//   - [MLPrecisionRecallCurve.InitWithPrecisionValuesPrecisionConfidenceThresholdsRecallValuesRecallConfidenceThresholds]
// See: https://developer.apple.com/documentation/CoreML/MLPrecisionRecallCurve
type MLPrecisionRecallCurve struct {
	objectivec.Object
}

// MLPrecisionRecallCurveFromID constructs a [MLPrecisionRecallCurve] from an objc.ID.
func MLPrecisionRecallCurveFromID(id objc.ID) MLPrecisionRecallCurve {
	return MLPrecisionRecallCurve{objectivec.Object{ID: id}}
}
// Ensure MLPrecisionRecallCurve implements IMLPrecisionRecallCurve.
var _ IMLPrecisionRecallCurve = MLPrecisionRecallCurve{}

// An interface definition for the [MLPrecisionRecallCurve] class.
//
// # Methods
//
//   - [IMLPrecisionRecallCurve.EncodeWithCoder]
//   - [IMLPrecisionRecallCurve.PrecisionConfidenceThresholds]
//   - [IMLPrecisionRecallCurve.PrecisionValues]
//   - [IMLPrecisionRecallCurve.RecallConfidenceThresholds]
//   - [IMLPrecisionRecallCurve.RecallValues]
//   - [IMLPrecisionRecallCurve.InitWithCoder]
//   - [IMLPrecisionRecallCurve.InitWithPrecisionValuesPrecisionConfidenceThresholdsRecallValuesRecallConfidenceThresholds]
//
// See: https://developer.apple.com/documentation/CoreML/MLPrecisionRecallCurve
type IMLPrecisionRecallCurve interface {
	objectivec.IObject

	// Topic: Methods

	EncodeWithCoder(coder foundation.INSCoder)
	PrecisionConfidenceThresholds() foundation.INSArray
	PrecisionValues() foundation.INSArray
	RecallConfidenceThresholds() foundation.INSArray
	RecallValues() foundation.INSArray
	InitWithCoder(coder foundation.INSCoder) MLPrecisionRecallCurve
	InitWithPrecisionValuesPrecisionConfidenceThresholdsRecallValuesRecallConfidenceThresholds(values objectivec.IObject, thresholds objectivec.IObject, values2 objectivec.IObject, thresholds2 objectivec.IObject) MLPrecisionRecallCurve
}

// Init initializes the instance.
func (p MLPrecisionRecallCurve) Init() MLPrecisionRecallCurve {
	rv := objc.Send[MLPrecisionRecallCurve](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p MLPrecisionRecallCurve) Autorelease() MLPrecisionRecallCurve {
	rv := objc.Send[MLPrecisionRecallCurve](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLPrecisionRecallCurve creates a new MLPrecisionRecallCurve instance.
func NewMLPrecisionRecallCurve() MLPrecisionRecallCurve {
	class := getMLPrecisionRecallCurveClass()
	rv := objc.Send[MLPrecisionRecallCurve](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLPrecisionRecallCurve/initWithCoder:
func NewPrecisionRecallCurveWithCoder(coder objectivec.IObject) MLPrecisionRecallCurve {
	instance := getMLPrecisionRecallCurveClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLPrecisionRecallCurveFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLPrecisionRecallCurve/initWithPrecisionValues:precisionConfidenceThresholds:recallValues:recallConfidenceThresholds:
func NewPrecisionRecallCurveWithPrecisionValuesPrecisionConfidenceThresholdsRecallValuesRecallConfidenceThresholds(values objectivec.IObject, thresholds objectivec.IObject, values2 objectivec.IObject, thresholds2 objectivec.IObject) MLPrecisionRecallCurve {
	instance := getMLPrecisionRecallCurveClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPrecisionValues:precisionConfidenceThresholds:recallValues:recallConfidenceThresholds:"), values, thresholds, values2, thresholds2)
	return MLPrecisionRecallCurveFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLPrecisionRecallCurve/encodeWithCoder:
func (p MLPrecisionRecallCurve) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](p.ID, objc.Sel("encodeWithCoder:"), coder)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLPrecisionRecallCurve/initWithCoder:
func (p MLPrecisionRecallCurve) InitWithCoder(coder foundation.INSCoder) MLPrecisionRecallCurve {
	rv := objc.Send[MLPrecisionRecallCurve](p.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLPrecisionRecallCurve/initWithPrecisionValues:precisionConfidenceThresholds:recallValues:recallConfidenceThresholds:
func (p MLPrecisionRecallCurve) InitWithPrecisionValuesPrecisionConfidenceThresholdsRecallValuesRecallConfidenceThresholds(values objectivec.IObject, thresholds objectivec.IObject, values2 objectivec.IObject, thresholds2 objectivec.IObject) MLPrecisionRecallCurve {
	rv := objc.Send[MLPrecisionRecallCurve](p.ID, objc.Sel("initWithPrecisionValues:precisionConfidenceThresholds:recallValues:recallConfidenceThresholds:"), values, thresholds, values2, thresholds2)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLPrecisionRecallCurve/supportsSecureCoding
func (_MLPrecisionRecallCurveClass MLPrecisionRecallCurveClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_MLPrecisionRecallCurveClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLPrecisionRecallCurve/precisionConfidenceThresholds
func (p MLPrecisionRecallCurve) PrecisionConfidenceThresholds() foundation.INSArray {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("precisionConfidenceThresholds"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLPrecisionRecallCurve/precisionValues
func (p MLPrecisionRecallCurve) PrecisionValues() foundation.INSArray {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("precisionValues"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLPrecisionRecallCurve/recallConfidenceThresholds
func (p MLPrecisionRecallCurve) RecallConfidenceThresholds() foundation.INSArray {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("recallConfidenceThresholds"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLPrecisionRecallCurve/recallValues
func (p MLPrecisionRecallCurve) RecallValues() foundation.INSArray {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("recallValues"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

