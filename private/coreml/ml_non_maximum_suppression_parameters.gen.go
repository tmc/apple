// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLNonMaximumSuppressionParameters] class.
var (
	_MLNonMaximumSuppressionParametersClass     MLNonMaximumSuppressionParametersClass
	_MLNonMaximumSuppressionParametersClassOnce sync.Once
)

func getMLNonMaximumSuppressionParametersClass() MLNonMaximumSuppressionParametersClass {
	_MLNonMaximumSuppressionParametersClassOnce.Do(func() {
		_MLNonMaximumSuppressionParametersClass = MLNonMaximumSuppressionParametersClass{class: objc.GetClass("MLNonMaximumSuppressionParameters")}
	})
	return _MLNonMaximumSuppressionParametersClass
}

// GetMLNonMaximumSuppressionParametersClass returns the class object for MLNonMaximumSuppressionParameters.
func GetMLNonMaximumSuppressionParametersClass() MLNonMaximumSuppressionParametersClass {
	return getMLNonMaximumSuppressionParametersClass()
}

type MLNonMaximumSuppressionParametersClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLNonMaximumSuppressionParametersClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLNonMaximumSuppressionParametersClass) Alloc() MLNonMaximumSuppressionParameters {
	rv := objc.Send[MLNonMaximumSuppressionParameters](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLNonMaximumSuppressionParameters.ConfidenceInputFeatureName]
//   - [MLNonMaximumSuppressionParameters.SetConfidenceInputFeatureName]
//   - [MLNonMaximumSuppressionParameters.ConfidenceOutputFeatureName]
//   - [MLNonMaximumSuppressionParameters.SetConfidenceOutputFeatureName]
//   - [MLNonMaximumSuppressionParameters.ConfidenceThreshold]
//   - [MLNonMaximumSuppressionParameters.SetConfidenceThreshold]
//   - [MLNonMaximumSuppressionParameters.ConfidenceThresholdInputFeatureName]
//   - [MLNonMaximumSuppressionParameters.SetConfidenceThresholdInputFeatureName]
//   - [MLNonMaximumSuppressionParameters.CoordinatesInputFeatureName]
//   - [MLNonMaximumSuppressionParameters.SetCoordinatesInputFeatureName]
//   - [MLNonMaximumSuppressionParameters.CoordinatesOutputFeatureName]
//   - [MLNonMaximumSuppressionParameters.SetCoordinatesOutputFeatureName]
//   - [MLNonMaximumSuppressionParameters.IouThreshold]
//   - [MLNonMaximumSuppressionParameters.SetIouThreshold]
//   - [MLNonMaximumSuppressionParameters.IouThresholdInputFeatureName]
//   - [MLNonMaximumSuppressionParameters.SetIouThresholdInputFeatureName]
//   - [MLNonMaximumSuppressionParameters.LabelNames]
//   - [MLNonMaximumSuppressionParameters.SetLabelNames]
//   - [MLNonMaximumSuppressionParameters.MaxBoxes]
//   - [MLNonMaximumSuppressionParameters.SetMaxBoxes]
//   - [MLNonMaximumSuppressionParameters.MinBoxes]
//   - [MLNonMaximumSuppressionParameters.SetMinBoxes]
//   - [MLNonMaximumSuppressionParameters.NumClasses]
//   - [MLNonMaximumSuppressionParameters.SetNumClasses]
//   - [MLNonMaximumSuppressionParameters.ObjectBoundingBoxOutputDescription]
//   - [MLNonMaximumSuppressionParameters.PerClass]
//   - [MLNonMaximumSuppressionParameters.SetPerClass]
//   - [MLNonMaximumSuppressionParameters.SuppressionMethod]
//   - [MLNonMaximumSuppressionParameters.SetSuppressionMethod]
// See: https://developer.apple.com/documentation/CoreML/MLNonMaximumSuppressionParameters
type MLNonMaximumSuppressionParameters struct {
	objectivec.Object
}

// MLNonMaximumSuppressionParametersFromID constructs a [MLNonMaximumSuppressionParameters] from an objc.ID.
func MLNonMaximumSuppressionParametersFromID(id objc.ID) MLNonMaximumSuppressionParameters {
	return MLNonMaximumSuppressionParameters{objectivec.Object{ID: id}}
}
// Ensure MLNonMaximumSuppressionParameters implements IMLNonMaximumSuppressionParameters.
var _ IMLNonMaximumSuppressionParameters = MLNonMaximumSuppressionParameters{}

// An interface definition for the [MLNonMaximumSuppressionParameters] class.
//
// # Methods
//
//   - [IMLNonMaximumSuppressionParameters.ConfidenceInputFeatureName]
//   - [IMLNonMaximumSuppressionParameters.SetConfidenceInputFeatureName]
//   - [IMLNonMaximumSuppressionParameters.ConfidenceOutputFeatureName]
//   - [IMLNonMaximumSuppressionParameters.SetConfidenceOutputFeatureName]
//   - [IMLNonMaximumSuppressionParameters.ConfidenceThreshold]
//   - [IMLNonMaximumSuppressionParameters.SetConfidenceThreshold]
//   - [IMLNonMaximumSuppressionParameters.ConfidenceThresholdInputFeatureName]
//   - [IMLNonMaximumSuppressionParameters.SetConfidenceThresholdInputFeatureName]
//   - [IMLNonMaximumSuppressionParameters.CoordinatesInputFeatureName]
//   - [IMLNonMaximumSuppressionParameters.SetCoordinatesInputFeatureName]
//   - [IMLNonMaximumSuppressionParameters.CoordinatesOutputFeatureName]
//   - [IMLNonMaximumSuppressionParameters.SetCoordinatesOutputFeatureName]
//   - [IMLNonMaximumSuppressionParameters.IouThreshold]
//   - [IMLNonMaximumSuppressionParameters.SetIouThreshold]
//   - [IMLNonMaximumSuppressionParameters.IouThresholdInputFeatureName]
//   - [IMLNonMaximumSuppressionParameters.SetIouThresholdInputFeatureName]
//   - [IMLNonMaximumSuppressionParameters.LabelNames]
//   - [IMLNonMaximumSuppressionParameters.SetLabelNames]
//   - [IMLNonMaximumSuppressionParameters.MaxBoxes]
//   - [IMLNonMaximumSuppressionParameters.SetMaxBoxes]
//   - [IMLNonMaximumSuppressionParameters.MinBoxes]
//   - [IMLNonMaximumSuppressionParameters.SetMinBoxes]
//   - [IMLNonMaximumSuppressionParameters.NumClasses]
//   - [IMLNonMaximumSuppressionParameters.SetNumClasses]
//   - [IMLNonMaximumSuppressionParameters.ObjectBoundingBoxOutputDescription]
//   - [IMLNonMaximumSuppressionParameters.PerClass]
//   - [IMLNonMaximumSuppressionParameters.SetPerClass]
//   - [IMLNonMaximumSuppressionParameters.SuppressionMethod]
//   - [IMLNonMaximumSuppressionParameters.SetSuppressionMethod]
//
// See: https://developer.apple.com/documentation/CoreML/MLNonMaximumSuppressionParameters
type IMLNonMaximumSuppressionParameters interface {
	objectivec.IObject

	// Topic: Methods

	ConfidenceInputFeatureName() string
	SetConfidenceInputFeatureName(value string)
	ConfidenceOutputFeatureName() string
	SetConfidenceOutputFeatureName(value string)
	ConfidenceThreshold() float64
	SetConfidenceThreshold(value float64)
	ConfidenceThresholdInputFeatureName() string
	SetConfidenceThresholdInputFeatureName(value string)
	CoordinatesInputFeatureName() string
	SetCoordinatesInputFeatureName(value string)
	CoordinatesOutputFeatureName() string
	SetCoordinatesOutputFeatureName(value string)
	IouThreshold() float64
	SetIouThreshold(value float64)
	IouThresholdInputFeatureName() string
	SetIouThresholdInputFeatureName(value string)
	LabelNames() foundation.INSArray
	SetLabelNames(value foundation.INSArray)
	MaxBoxes() int64
	SetMaxBoxes(value int64)
	MinBoxes() uint64
	SetMinBoxes(value uint64)
	NumClasses() uint64
	SetNumClasses(value uint64)
	ObjectBoundingBoxOutputDescription() objectivec.IObject
	PerClass() bool
	SetPerClass(value bool)
	SuppressionMethod() int
	SetSuppressionMethod(value int)
}

// Init initializes the instance.
func (n MLNonMaximumSuppressionParameters) Init() MLNonMaximumSuppressionParameters {
	rv := objc.Send[MLNonMaximumSuppressionParameters](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n MLNonMaximumSuppressionParameters) Autorelease() MLNonMaximumSuppressionParameters {
	rv := objc.Send[MLNonMaximumSuppressionParameters](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLNonMaximumSuppressionParameters creates a new MLNonMaximumSuppressionParameters instance.
func NewMLNonMaximumSuppressionParameters() MLNonMaximumSuppressionParameters {
	class := getMLNonMaximumSuppressionParametersClass()
	rv := objc.Send[MLNonMaximumSuppressionParameters](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLNonMaximumSuppressionParameters/objectBoundingBoxOutputDescription
func (n MLNonMaximumSuppressionParameters) ObjectBoundingBoxOutputDescription() objectivec.IObject {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("objectBoundingBoxOutputDescription"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLNonMaximumSuppressionParameters/confidenceInputFeatureName
func (n MLNonMaximumSuppressionParameters) ConfidenceInputFeatureName() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("confidenceInputFeatureName"))
	return foundation.NSStringFromID(rv).String()
}
func (n MLNonMaximumSuppressionParameters) SetConfidenceInputFeatureName(value string) {
	objc.Send[struct{}](n.ID, objc.Sel("setConfidenceInputFeatureName:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/CoreML/MLNonMaximumSuppressionParameters/confidenceOutputFeatureName
func (n MLNonMaximumSuppressionParameters) ConfidenceOutputFeatureName() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("confidenceOutputFeatureName"))
	return foundation.NSStringFromID(rv).String()
}
func (n MLNonMaximumSuppressionParameters) SetConfidenceOutputFeatureName(value string) {
	objc.Send[struct{}](n.ID, objc.Sel("setConfidenceOutputFeatureName:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/CoreML/MLNonMaximumSuppressionParameters/confidenceThreshold
func (n MLNonMaximumSuppressionParameters) ConfidenceThreshold() float64 {
	rv := objc.Send[float64](n.ID, objc.Sel("confidenceThreshold"))
	return rv
}
func (n MLNonMaximumSuppressionParameters) SetConfidenceThreshold(value float64) {
	objc.Send[struct{}](n.ID, objc.Sel("setConfidenceThreshold:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLNonMaximumSuppressionParameters/confidenceThresholdInputFeatureName
func (n MLNonMaximumSuppressionParameters) ConfidenceThresholdInputFeatureName() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("confidenceThresholdInputFeatureName"))
	return foundation.NSStringFromID(rv).String()
}
func (n MLNonMaximumSuppressionParameters) SetConfidenceThresholdInputFeatureName(value string) {
	objc.Send[struct{}](n.ID, objc.Sel("setConfidenceThresholdInputFeatureName:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/CoreML/MLNonMaximumSuppressionParameters/coordinatesInputFeatureName
func (n MLNonMaximumSuppressionParameters) CoordinatesInputFeatureName() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("coordinatesInputFeatureName"))
	return foundation.NSStringFromID(rv).String()
}
func (n MLNonMaximumSuppressionParameters) SetCoordinatesInputFeatureName(value string) {
	objc.Send[struct{}](n.ID, objc.Sel("setCoordinatesInputFeatureName:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/CoreML/MLNonMaximumSuppressionParameters/coordinatesOutputFeatureName
func (n MLNonMaximumSuppressionParameters) CoordinatesOutputFeatureName() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("coordinatesOutputFeatureName"))
	return foundation.NSStringFromID(rv).String()
}
func (n MLNonMaximumSuppressionParameters) SetCoordinatesOutputFeatureName(value string) {
	objc.Send[struct{}](n.ID, objc.Sel("setCoordinatesOutputFeatureName:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/CoreML/MLNonMaximumSuppressionParameters/iouThreshold
func (n MLNonMaximumSuppressionParameters) IouThreshold() float64 {
	rv := objc.Send[float64](n.ID, objc.Sel("iouThreshold"))
	return rv
}
func (n MLNonMaximumSuppressionParameters) SetIouThreshold(value float64) {
	objc.Send[struct{}](n.ID, objc.Sel("setIouThreshold:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLNonMaximumSuppressionParameters/iouThresholdInputFeatureName
func (n MLNonMaximumSuppressionParameters) IouThresholdInputFeatureName() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("iouThresholdInputFeatureName"))
	return foundation.NSStringFromID(rv).String()
}
func (n MLNonMaximumSuppressionParameters) SetIouThresholdInputFeatureName(value string) {
	objc.Send[struct{}](n.ID, objc.Sel("setIouThresholdInputFeatureName:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/CoreML/MLNonMaximumSuppressionParameters/labelNames
func (n MLNonMaximumSuppressionParameters) LabelNames() foundation.INSArray {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("labelNames"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (n MLNonMaximumSuppressionParameters) SetLabelNames(value foundation.INSArray) {
	objc.Send[struct{}](n.ID, objc.Sel("setLabelNames:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLNonMaximumSuppressionParameters/maxBoxes
func (n MLNonMaximumSuppressionParameters) MaxBoxes() int64 {
	rv := objc.Send[int64](n.ID, objc.Sel("maxBoxes"))
	return rv
}
func (n MLNonMaximumSuppressionParameters) SetMaxBoxes(value int64) {
	objc.Send[struct{}](n.ID, objc.Sel("setMaxBoxes:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLNonMaximumSuppressionParameters/minBoxes
func (n MLNonMaximumSuppressionParameters) MinBoxes() uint64 {
	rv := objc.Send[uint64](n.ID, objc.Sel("minBoxes"))
	return rv
}
func (n MLNonMaximumSuppressionParameters) SetMinBoxes(value uint64) {
	objc.Send[struct{}](n.ID, objc.Sel("setMinBoxes:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLNonMaximumSuppressionParameters/numClasses
func (n MLNonMaximumSuppressionParameters) NumClasses() uint64 {
	rv := objc.Send[uint64](n.ID, objc.Sel("numClasses"))
	return rv
}
func (n MLNonMaximumSuppressionParameters) SetNumClasses(value uint64) {
	objc.Send[struct{}](n.ID, objc.Sel("setNumClasses:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLNonMaximumSuppressionParameters/perClass
func (n MLNonMaximumSuppressionParameters) PerClass() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("perClass"))
	return rv
}
func (n MLNonMaximumSuppressionParameters) SetPerClass(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setPerClass:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLNonMaximumSuppressionParameters/suppressionMethod
func (n MLNonMaximumSuppressionParameters) SuppressionMethod() int {
	rv := objc.Send[int](n.ID, objc.Sel("suppressionMethod"))
	return rv
}
func (n MLNonMaximumSuppressionParameters) SetSuppressionMethod(value int) {
	objc.Send[struct{}](n.ID, objc.Sel("setSuppressionMethod:"), value)
}

