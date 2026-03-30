// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLKNearestNeighborsClassifierParameters] class.
var (
	_MLKNearestNeighborsClassifierParametersClass     MLKNearestNeighborsClassifierParametersClass
	_MLKNearestNeighborsClassifierParametersClassOnce sync.Once
)

func getMLKNearestNeighborsClassifierParametersClass() MLKNearestNeighborsClassifierParametersClass {
	_MLKNearestNeighborsClassifierParametersClassOnce.Do(func() {
		_MLKNearestNeighborsClassifierParametersClass = MLKNearestNeighborsClassifierParametersClass{class: objc.GetClass("MLKNearestNeighborsClassifierParameters")}
	})
	return _MLKNearestNeighborsClassifierParametersClass
}

// GetMLKNearestNeighborsClassifierParametersClass returns the class object for MLKNearestNeighborsClassifierParameters.
func GetMLKNearestNeighborsClassifierParametersClass() MLKNearestNeighborsClassifierParametersClass {
	return getMLKNearestNeighborsClassifierParametersClass()
}

type MLKNearestNeighborsClassifierParametersClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLKNearestNeighborsClassifierParametersClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLKNearestNeighborsClassifierParametersClass) Alloc() MLKNearestNeighborsClassifierParameters {
	rv := objc.Send[MLKNearestNeighborsClassifierParameters](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLKNearestNeighborsClassifierParameters.DefaultLabel]
//   - [MLKNearestNeighborsClassifierParameters.SetDefaultLabel]
//   - [MLKNearestNeighborsClassifierParameters.IndexType]
//   - [MLKNearestNeighborsClassifierParameters.SetIndexType]
//   - [MLKNearestNeighborsClassifierParameters.LeafSize]
//   - [MLKNearestNeighborsClassifierParameters.SetLeafSize]
//   - [MLKNearestNeighborsClassifierParameters.NearestDistancesFeatureName]
//   - [MLKNearestNeighborsClassifierParameters.SetNearestDistancesFeatureName]
//   - [MLKNearestNeighborsClassifierParameters.NearestLabelsFeatureName]
//   - [MLKNearestNeighborsClassifierParameters.SetNearestLabelsFeatureName]
//   - [MLKNearestNeighborsClassifierParameters.NumberOfDimensions]
//   - [MLKNearestNeighborsClassifierParameters.SetNumberOfDimensions]
//   - [MLKNearestNeighborsClassifierParameters.WeightingScheme]
//   - [MLKNearestNeighborsClassifierParameters.SetWeightingScheme]
//
// See: https://developer.apple.com/documentation/CoreML/MLKNearestNeighborsClassifierParameters
type MLKNearestNeighborsClassifierParameters struct {
	objectivec.Object
}

// MLKNearestNeighborsClassifierParametersFromID constructs a [MLKNearestNeighborsClassifierParameters] from an objc.ID.
func MLKNearestNeighborsClassifierParametersFromID(id objc.ID) MLKNearestNeighborsClassifierParameters {
	return MLKNearestNeighborsClassifierParameters{objectivec.Object{ID: id}}
}

// Ensure MLKNearestNeighborsClassifierParameters implements IMLKNearestNeighborsClassifierParameters.
var _ IMLKNearestNeighborsClassifierParameters = MLKNearestNeighborsClassifierParameters{}

// An interface definition for the [MLKNearestNeighborsClassifierParameters] class.
//
// # Methods
//
//   - [IMLKNearestNeighborsClassifierParameters.DefaultLabel]
//   - [IMLKNearestNeighborsClassifierParameters.SetDefaultLabel]
//   - [IMLKNearestNeighborsClassifierParameters.IndexType]
//   - [IMLKNearestNeighborsClassifierParameters.SetIndexType]
//   - [IMLKNearestNeighborsClassifierParameters.LeafSize]
//   - [IMLKNearestNeighborsClassifierParameters.SetLeafSize]
//   - [IMLKNearestNeighborsClassifierParameters.NearestDistancesFeatureName]
//   - [IMLKNearestNeighborsClassifierParameters.SetNearestDistancesFeatureName]
//   - [IMLKNearestNeighborsClassifierParameters.NearestLabelsFeatureName]
//   - [IMLKNearestNeighborsClassifierParameters.SetNearestLabelsFeatureName]
//   - [IMLKNearestNeighborsClassifierParameters.NumberOfDimensions]
//   - [IMLKNearestNeighborsClassifierParameters.SetNumberOfDimensions]
//   - [IMLKNearestNeighborsClassifierParameters.WeightingScheme]
//   - [IMLKNearestNeighborsClassifierParameters.SetWeightingScheme]
//
// See: https://developer.apple.com/documentation/CoreML/MLKNearestNeighborsClassifierParameters
type IMLKNearestNeighborsClassifierParameters interface {
	objectivec.IObject

	// Topic: Methods

	DefaultLabel() objectivec.Object
	SetDefaultLabel(value objectivec.Object)
	IndexType() int64
	SetIndexType(value int64)
	LeafSize() uint64
	SetLeafSize(value uint64)
	NearestDistancesFeatureName() string
	SetNearestDistancesFeatureName(value string)
	NearestLabelsFeatureName() string
	SetNearestLabelsFeatureName(value string)
	NumberOfDimensions() uint64
	SetNumberOfDimensions(value uint64)
	WeightingScheme() int64
	SetWeightingScheme(value int64)
}

// Init initializes the instance.
func (k MLKNearestNeighborsClassifierParameters) Init() MLKNearestNeighborsClassifierParameters {
	rv := objc.Send[MLKNearestNeighborsClassifierParameters](k.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (k MLKNearestNeighborsClassifierParameters) Autorelease() MLKNearestNeighborsClassifierParameters {
	rv := objc.Send[MLKNearestNeighborsClassifierParameters](k.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLKNearestNeighborsClassifierParameters creates a new MLKNearestNeighborsClassifierParameters instance.
func NewMLKNearestNeighborsClassifierParameters() MLKNearestNeighborsClassifierParameters {
	class := getMLKNearestNeighborsClassifierParametersClass()
	rv := objc.Send[MLKNearestNeighborsClassifierParameters](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLKNearestNeighborsClassifierParameters/defaultLabel
func (k MLKNearestNeighborsClassifierParameters) DefaultLabel() objectivec.Object {
	rv := objc.Send[objc.ID](k.ID, objc.Sel("defaultLabel"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (k MLKNearestNeighborsClassifierParameters) SetDefaultLabel(value objectivec.Object) {
	objc.Send[struct{}](k.ID, objc.Sel("setDefaultLabel:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLKNearestNeighborsClassifierParameters/indexType
func (k MLKNearestNeighborsClassifierParameters) IndexType() int64 {
	rv := objc.Send[int64](k.ID, objc.Sel("indexType"))
	return rv
}
func (k MLKNearestNeighborsClassifierParameters) SetIndexType(value int64) {
	objc.Send[struct{}](k.ID, objc.Sel("setIndexType:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLKNearestNeighborsClassifierParameters/leafSize
func (k MLKNearestNeighborsClassifierParameters) LeafSize() uint64 {
	rv := objc.Send[uint64](k.ID, objc.Sel("leafSize"))
	return rv
}
func (k MLKNearestNeighborsClassifierParameters) SetLeafSize(value uint64) {
	objc.Send[struct{}](k.ID, objc.Sel("setLeafSize:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLKNearestNeighborsClassifierParameters/nearestDistancesFeatureName
func (k MLKNearestNeighborsClassifierParameters) NearestDistancesFeatureName() string {
	rv := objc.Send[objc.ID](k.ID, objc.Sel("nearestDistancesFeatureName"))
	return foundation.NSStringFromID(rv).String()
}
func (k MLKNearestNeighborsClassifierParameters) SetNearestDistancesFeatureName(value string) {
	objc.Send[struct{}](k.ID, objc.Sel("setNearestDistancesFeatureName:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/CoreML/MLKNearestNeighborsClassifierParameters/nearestLabelsFeatureName
func (k MLKNearestNeighborsClassifierParameters) NearestLabelsFeatureName() string {
	rv := objc.Send[objc.ID](k.ID, objc.Sel("nearestLabelsFeatureName"))
	return foundation.NSStringFromID(rv).String()
}
func (k MLKNearestNeighborsClassifierParameters) SetNearestLabelsFeatureName(value string) {
	objc.Send[struct{}](k.ID, objc.Sel("setNearestLabelsFeatureName:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/CoreML/MLKNearestNeighborsClassifierParameters/numberOfDimensions
func (k MLKNearestNeighborsClassifierParameters) NumberOfDimensions() uint64 {
	rv := objc.Send[uint64](k.ID, objc.Sel("numberOfDimensions"))
	return rv
}
func (k MLKNearestNeighborsClassifierParameters) SetNumberOfDimensions(value uint64) {
	objc.Send[struct{}](k.ID, objc.Sel("setNumberOfDimensions:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLKNearestNeighborsClassifierParameters/weightingScheme
func (k MLKNearestNeighborsClassifierParameters) WeightingScheme() int64 {
	rv := objc.Send[int64](k.ID, objc.Sel("weightingScheme"))
	return rv
}
func (k MLKNearestNeighborsClassifierParameters) SetWeightingScheme(value int64) {
	objc.Send[struct{}](k.ID, objc.Sel("setWeightingScheme:"), value)
}
