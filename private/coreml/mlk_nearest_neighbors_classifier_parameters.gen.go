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
func (m MLKNearestNeighborsClassifierParameters) Init() MLKNearestNeighborsClassifierParameters {
	rv := objc.Send[MLKNearestNeighborsClassifierParameters](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLKNearestNeighborsClassifierParameters) Autorelease() MLKNearestNeighborsClassifierParameters {
	rv := objc.Send[MLKNearestNeighborsClassifierParameters](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLKNearestNeighborsClassifierParameters creates a new MLKNearestNeighborsClassifierParameters instance.
func NewMLKNearestNeighborsClassifierParameters() MLKNearestNeighborsClassifierParameters {
	class := getMLKNearestNeighborsClassifierParametersClass()
	rv := objc.Send[MLKNearestNeighborsClassifierParameters](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (m MLKNearestNeighborsClassifierParameters) DefaultLabel() objectivec.Object {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("defaultLabel"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (m MLKNearestNeighborsClassifierParameters) SetDefaultLabel(value objectivec.Object) {
	objc.Send[struct{}](m.ID, objc.Sel("setDefaultLabel:"), value)
}
func (m MLKNearestNeighborsClassifierParameters) IndexType() int64 {
	rv := objc.Send[int64](m.ID, objc.Sel("indexType"))
	return rv
}
func (m MLKNearestNeighborsClassifierParameters) SetIndexType(value int64) {
	objc.Send[struct{}](m.ID, objc.Sel("setIndexType:"), value)
}
func (m MLKNearestNeighborsClassifierParameters) LeafSize() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("leafSize"))
	return rv
}
func (m MLKNearestNeighborsClassifierParameters) SetLeafSize(value uint64) {
	objc.Send[struct{}](m.ID, objc.Sel("setLeafSize:"), value)
}
func (m MLKNearestNeighborsClassifierParameters) NearestDistancesFeatureName() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("nearestDistancesFeatureName"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLKNearestNeighborsClassifierParameters) SetNearestDistancesFeatureName(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setNearestDistancesFeatureName:"), objc.String(value))
}
func (m MLKNearestNeighborsClassifierParameters) NearestLabelsFeatureName() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("nearestLabelsFeatureName"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLKNearestNeighborsClassifierParameters) SetNearestLabelsFeatureName(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setNearestLabelsFeatureName:"), objc.String(value))
}
func (m MLKNearestNeighborsClassifierParameters) NumberOfDimensions() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("numberOfDimensions"))
	return rv
}
func (m MLKNearestNeighborsClassifierParameters) SetNumberOfDimensions(value uint64) {
	objc.Send[struct{}](m.ID, objc.Sel("setNumberOfDimensions:"), value)
}
func (m MLKNearestNeighborsClassifierParameters) WeightingScheme() int64 {
	rv := objc.Send[int64](m.ID, objc.Sel("weightingScheme"))
	return rv
}
func (m MLKNearestNeighborsClassifierParameters) SetWeightingScheme(value int64) {
	objc.Send[struct{}](m.ID, objc.Sel("setWeightingScheme:"), value)
}
