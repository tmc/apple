// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLPredictionEventMetric] class.
var (
	_MLPredictionEventMetricClass     MLPredictionEventMetricClass
	_MLPredictionEventMetricClassOnce sync.Once
)

func getMLPredictionEventMetricClass() MLPredictionEventMetricClass {
	_MLPredictionEventMetricClassOnce.Do(func() {
		_MLPredictionEventMetricClass = MLPredictionEventMetricClass{class: objc.GetClass("MLPredictionEventMetric")}
	})
	return _MLPredictionEventMetricClass
}

// GetMLPredictionEventMetricClass returns the class object for MLPredictionEventMetric.
func GetMLPredictionEventMetricClass() MLPredictionEventMetricClass {
	return getMLPredictionEventMetricClass()
}

type MLPredictionEventMetricClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLPredictionEventMetricClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLPredictionEventMetricClass) Alloc() MLPredictionEventMetric {
	rv := objc.Send[MLPredictionEventMetric](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLPredictionEventMetric.BundleIdentifier]
//   - [MLPredictionEventMetric.DictionaryRepresentation]
//   - [MLPredictionEventMetric.FeaturesPredictionCountSoFar]
//   - [MLPredictionEventMetric.FeaturesPredictionDuration]
//   - [MLPredictionEventMetric.FirstPartyExecutable]
//   - [MLPredictionEventMetric.ModelName]
//   - [MLPredictionEventMetric.ModelType]
//   - [MLPredictionEventMetric.Name]
//   - [MLPredictionEventMetric.InitWithBundleIdentifierModelNameFirstPartyExecutableModelTypeFeaturesPredictionDurationFeaturesPredictionCountSoFar]
//   - [MLPredictionEventMetric.DebugDescription]
//   - [MLPredictionEventMetric.Description]
//   - [MLPredictionEventMetric.Hash]
//   - [MLPredictionEventMetric.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLPredictionEventMetric
type MLPredictionEventMetric struct {
	objectivec.Object
}

// MLPredictionEventMetricFromID constructs a [MLPredictionEventMetric] from an objc.ID.
func MLPredictionEventMetricFromID(id objc.ID) MLPredictionEventMetric {
	return MLPredictionEventMetric{objectivec.Object{ID: id}}
}

// Ensure MLPredictionEventMetric implements IMLPredictionEventMetric.
var _ IMLPredictionEventMetric = MLPredictionEventMetric{}

// An interface definition for the [MLPredictionEventMetric] class.
//
// # Methods
//
//   - [IMLPredictionEventMetric.BundleIdentifier]
//   - [IMLPredictionEventMetric.DictionaryRepresentation]
//   - [IMLPredictionEventMetric.FeaturesPredictionCountSoFar]
//   - [IMLPredictionEventMetric.FeaturesPredictionDuration]
//   - [IMLPredictionEventMetric.FirstPartyExecutable]
//   - [IMLPredictionEventMetric.ModelName]
//   - [IMLPredictionEventMetric.ModelType]
//   - [IMLPredictionEventMetric.Name]
//   - [IMLPredictionEventMetric.InitWithBundleIdentifierModelNameFirstPartyExecutableModelTypeFeaturesPredictionDurationFeaturesPredictionCountSoFar]
//   - [IMLPredictionEventMetric.DebugDescription]
//   - [IMLPredictionEventMetric.Description]
//   - [IMLPredictionEventMetric.Hash]
//   - [IMLPredictionEventMetric.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLPredictionEventMetric
type IMLPredictionEventMetric interface {
	objectivec.IObject

	// Topic: Methods

	BundleIdentifier() string
	DictionaryRepresentation() foundation.INSDictionary
	FeaturesPredictionCountSoFar() int64
	FeaturesPredictionDuration() float64
	FirstPartyExecutable() foundation.NSNumber
	ModelName() string
	ModelType() foundation.NSNumber
	Name() string
	InitWithBundleIdentifierModelNameFirstPartyExecutableModelTypeFeaturesPredictionDurationFeaturesPredictionCountSoFar(identifier objectivec.IObject, name objectivec.IObject, executable objectivec.IObject, type_ objectivec.IObject, duration float64, far int64) MLPredictionEventMetric
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (p MLPredictionEventMetric) Init() MLPredictionEventMetric {
	rv := objc.Send[MLPredictionEventMetric](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p MLPredictionEventMetric) Autorelease() MLPredictionEventMetric {
	rv := objc.Send[MLPredictionEventMetric](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLPredictionEventMetric creates a new MLPredictionEventMetric instance.
func NewMLPredictionEventMetric() MLPredictionEventMetric {
	class := getMLPredictionEventMetricClass()
	rv := objc.Send[MLPredictionEventMetric](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLPredictionEventMetric/initWithBundleIdentifier:modelName:firstPartyExecutable:modelType:featuresPredictionDuration:featuresPredictionCountSoFar:
func NewPredictionEventMetricWithBundleIdentifierModelNameFirstPartyExecutableModelTypeFeaturesPredictionDurationFeaturesPredictionCountSoFar(identifier objectivec.IObject, name objectivec.IObject, executable objectivec.IObject, type_ objectivec.IObject, duration float64, far int64) MLPredictionEventMetric {
	instance := getMLPredictionEventMetricClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBundleIdentifier:modelName:firstPartyExecutable:modelType:featuresPredictionDuration:featuresPredictionCountSoFar:"), identifier, name, executable, type_, duration, far)
	return MLPredictionEventMetricFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLPredictionEventMetric/initWithBundleIdentifier:modelName:firstPartyExecutable:modelType:featuresPredictionDuration:featuresPredictionCountSoFar:
func (p MLPredictionEventMetric) InitWithBundleIdentifierModelNameFirstPartyExecutableModelTypeFeaturesPredictionDurationFeaturesPredictionCountSoFar(identifier objectivec.IObject, name objectivec.IObject, executable objectivec.IObject, type_ objectivec.IObject, duration float64, far int64) MLPredictionEventMetric {
	rv := objc.Send[MLPredictionEventMetric](p.ID, objc.Sel("initWithBundleIdentifier:modelName:firstPartyExecutable:modelType:featuresPredictionDuration:featuresPredictionCountSoFar:"), identifier, name, executable, type_, duration, far)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLPredictionEventMetric/bundleIdentifier
func (p MLPredictionEventMetric) BundleIdentifier() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("bundleIdentifier"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLPredictionEventMetric/debugDescription
func (p MLPredictionEventMetric) DebugDescription() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLPredictionEventMetric/description
func (p MLPredictionEventMetric) Description() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLPredictionEventMetric/dictionaryRepresentation
func (p MLPredictionEventMetric) DictionaryRepresentation() foundation.INSDictionary {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("dictionaryRepresentation"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLPredictionEventMetric/featuresPredictionCountSoFar
func (p MLPredictionEventMetric) FeaturesPredictionCountSoFar() int64 {
	rv := objc.Send[int64](p.ID, objc.Sel("featuresPredictionCountSoFar"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLPredictionEventMetric/featuresPredictionDuration
func (p MLPredictionEventMetric) FeaturesPredictionDuration() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("featuresPredictionDuration"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLPredictionEventMetric/firstPartyExecutable
func (p MLPredictionEventMetric) FirstPartyExecutable() foundation.NSNumber {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("firstPartyExecutable"))
	return foundation.NSNumberFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLPredictionEventMetric/hash
func (p MLPredictionEventMetric) Hash() uint64 {
	rv := objc.Send[uint64](p.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLPredictionEventMetric/modelName
func (p MLPredictionEventMetric) ModelName() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("modelName"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLPredictionEventMetric/modelType
func (p MLPredictionEventMetric) ModelType() foundation.NSNumber {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("modelType"))
	return foundation.NSNumberFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLPredictionEventMetric/name
func (p MLPredictionEventMetric) Name() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLPredictionEventMetric/superclass
func (p MLPredictionEventMetric) Superclass() objc.Class {
	rv := objc.Send[objc.Class](p.ID, objc.Sel("superclass"))
	return rv
}
