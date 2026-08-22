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
	rv := objc.SendIfResponds[MLPredictionEventMetric](objc.ID(mc.class), objc.Sel("alloc"))
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
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (m MLPredictionEventMetric) Init() MLPredictionEventMetric {
	rv := objc.SendIfResponds[MLPredictionEventMetric](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLPredictionEventMetric) Autorelease() MLPredictionEventMetric {
	rv := objc.SendIfResponds[MLPredictionEventMetric](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLPredictionEventMetric creates a new MLPredictionEventMetric instance.
func NewMLPredictionEventMetric() MLPredictionEventMetric {
	class := getMLPredictionEventMetricClass()
	rv := objc.SendIfResponds[MLPredictionEventMetric](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewPredictionEventMetricWithBundleIdentifierModelNameFirstPartyExecutableModelTypeFeaturesPredictionDurationFeaturesPredictionCountSoFar(identifier objectivec.IObject, name objectivec.IObject, executable objectivec.IObject, type_ objectivec.IObject, duration float64, far int64) MLPredictionEventMetric {
	instance := getMLPredictionEventMetricClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithBundleIdentifier:modelName:firstPartyExecutable:modelType:featuresPredictionDuration:featuresPredictionCountSoFar:"), identifier, name, executable, type_, duration, far)
	return MLPredictionEventMetricFromID(rv)
}

func (m MLPredictionEventMetric) InitWithBundleIdentifierModelNameFirstPartyExecutableModelTypeFeaturesPredictionDurationFeaturesPredictionCountSoFar(identifier objectivec.IObject, name objectivec.IObject, executable objectivec.IObject, type_ objectivec.IObject, duration float64, far int64) MLPredictionEventMetric {
	rv := objc.SendIfResponds[MLPredictionEventMetric](m.ID, objc.Sel("initWithBundleIdentifier:modelName:firstPartyExecutable:modelType:featuresPredictionDuration:featuresPredictionCountSoFar:"), identifier, name, executable, type_, duration, far)
	return rv
}

func (m MLPredictionEventMetric) BundleIdentifier() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("bundleIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLPredictionEventMetric) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLPredictionEventMetric) Description() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLPredictionEventMetric) DictionaryRepresentation() foundation.INSDictionary {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("dictionaryRepresentation"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLPredictionEventMetric) FeaturesPredictionCountSoFar() int64 {
	rv := objc.SendIfResponds[int64](m.ID, objc.Sel("featuresPredictionCountSoFar"))
	return rv
}
func (m MLPredictionEventMetric) FeaturesPredictionDuration() float64 {
	rv := objc.SendIfResponds[float64](m.ID, objc.Sel("featuresPredictionDuration"))
	return rv
}
func (m MLPredictionEventMetric) FirstPartyExecutable() foundation.NSNumber {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("firstPartyExecutable"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (m MLPredictionEventMetric) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("hash"))
	return rv
}
func (m MLPredictionEventMetric) ModelName() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("modelName"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLPredictionEventMetric) ModelType() foundation.NSNumber {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("modelType"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (m MLPredictionEventMetric) Name() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLPredictionEventMetric) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](m.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
