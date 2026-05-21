// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
func (m MLNonMaximumSuppressionParameters) Init() MLNonMaximumSuppressionParameters {
	rv := objc.Send[MLNonMaximumSuppressionParameters](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLNonMaximumSuppressionParameters) Autorelease() MLNonMaximumSuppressionParameters {
	rv := objc.Send[MLNonMaximumSuppressionParameters](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLNonMaximumSuppressionParameters creates a new MLNonMaximumSuppressionParameters instance.
func NewMLNonMaximumSuppressionParameters() MLNonMaximumSuppressionParameters {
	class := getMLNonMaximumSuppressionParametersClass()
	rv := objc.Send[MLNonMaximumSuppressionParameters](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (m MLNonMaximumSuppressionParameters) ObjectBoundingBoxOutputDescription() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("objectBoundingBoxOutputDescription"))
	return objectivec.Object{ID: rv}
}

func (m MLNonMaximumSuppressionParameters) ConfidenceInputFeatureName() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("confidenceInputFeatureName"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLNonMaximumSuppressionParameters) SetConfidenceInputFeatureName(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setConfidenceInputFeatureName:"), objc.String(value))
}
func (m MLNonMaximumSuppressionParameters) ConfidenceOutputFeatureName() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("confidenceOutputFeatureName"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLNonMaximumSuppressionParameters) SetConfidenceOutputFeatureName(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setConfidenceOutputFeatureName:"), objc.String(value))
}
func (m MLNonMaximumSuppressionParameters) ConfidenceThreshold() float64 {
	rv := objc.Send[float64](m.ID, objc.Sel("confidenceThreshold"))
	return rv
}
func (m MLNonMaximumSuppressionParameters) SetConfidenceThreshold(value float64) {
	objc.Send[struct{}](m.ID, objc.Sel("setConfidenceThreshold:"), value)
}
func (m MLNonMaximumSuppressionParameters) ConfidenceThresholdInputFeatureName() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("confidenceThresholdInputFeatureName"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLNonMaximumSuppressionParameters) SetConfidenceThresholdInputFeatureName(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setConfidenceThresholdInputFeatureName:"), objc.String(value))
}
func (m MLNonMaximumSuppressionParameters) CoordinatesInputFeatureName() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("coordinatesInputFeatureName"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLNonMaximumSuppressionParameters) SetCoordinatesInputFeatureName(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setCoordinatesInputFeatureName:"), objc.String(value))
}
func (m MLNonMaximumSuppressionParameters) CoordinatesOutputFeatureName() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("coordinatesOutputFeatureName"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLNonMaximumSuppressionParameters) SetCoordinatesOutputFeatureName(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setCoordinatesOutputFeatureName:"), objc.String(value))
}
func (m MLNonMaximumSuppressionParameters) IouThreshold() float64 {
	rv := objc.Send[float64](m.ID, objc.Sel("iouThreshold"))
	return rv
}
func (m MLNonMaximumSuppressionParameters) SetIouThreshold(value float64) {
	objc.Send[struct{}](m.ID, objc.Sel("setIouThreshold:"), value)
}
func (m MLNonMaximumSuppressionParameters) IouThresholdInputFeatureName() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("iouThresholdInputFeatureName"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLNonMaximumSuppressionParameters) SetIouThresholdInputFeatureName(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setIouThresholdInputFeatureName:"), objc.String(value))
}
func (m MLNonMaximumSuppressionParameters) LabelNames() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("labelNames"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLNonMaximumSuppressionParameters) SetLabelNames(value foundation.INSArray) {
	objc.Send[struct{}](m.ID, objc.Sel("setLabelNames:"), value)
}
func (m MLNonMaximumSuppressionParameters) MaxBoxes() int64 {
	rv := objc.Send[int64](m.ID, objc.Sel("maxBoxes"))
	return rv
}
func (m MLNonMaximumSuppressionParameters) SetMaxBoxes(value int64) {
	objc.Send[struct{}](m.ID, objc.Sel("setMaxBoxes:"), value)
}
func (m MLNonMaximumSuppressionParameters) MinBoxes() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("minBoxes"))
	return rv
}
func (m MLNonMaximumSuppressionParameters) SetMinBoxes(value uint64) {
	objc.Send[struct{}](m.ID, objc.Sel("setMinBoxes:"), value)
}
func (m MLNonMaximumSuppressionParameters) NumClasses() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("numClasses"))
	return rv
}
func (m MLNonMaximumSuppressionParameters) SetNumClasses(value uint64) {
	objc.Send[struct{}](m.ID, objc.Sel("setNumClasses:"), value)
}
func (m MLNonMaximumSuppressionParameters) PerClass() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("perClass"))
	return rv
}
func (m MLNonMaximumSuppressionParameters) SetPerClass(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setPerClass:"), value)
}
func (m MLNonMaximumSuppressionParameters) SuppressionMethod() int {
	rv := objc.Send[int](m.ID, objc.Sel("suppressionMethod"))
	return rv
}
func (m MLNonMaximumSuppressionParameters) SetSuppressionMethod(value int) {
	objc.Send[struct{}](m.ID, objc.Sel("setSuppressionMethod:"), value)
}
