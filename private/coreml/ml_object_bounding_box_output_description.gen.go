// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLObjectBoundingBoxOutputDescription] class.
var (
	_MLObjectBoundingBoxOutputDescriptionClass     MLObjectBoundingBoxOutputDescriptionClass
	_MLObjectBoundingBoxOutputDescriptionClassOnce sync.Once
)

func getMLObjectBoundingBoxOutputDescriptionClass() MLObjectBoundingBoxOutputDescriptionClass {
	_MLObjectBoundingBoxOutputDescriptionClassOnce.Do(func() {
		_MLObjectBoundingBoxOutputDescriptionClass = MLObjectBoundingBoxOutputDescriptionClass{class: objc.GetClass("MLObjectBoundingBoxOutputDescription")}
	})
	return _MLObjectBoundingBoxOutputDescriptionClass
}

// GetMLObjectBoundingBoxOutputDescriptionClass returns the class object for MLObjectBoundingBoxOutputDescription.
func GetMLObjectBoundingBoxOutputDescriptionClass() MLObjectBoundingBoxOutputDescriptionClass {
	return getMLObjectBoundingBoxOutputDescriptionClass()
}

type MLObjectBoundingBoxOutputDescriptionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLObjectBoundingBoxOutputDescriptionClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLObjectBoundingBoxOutputDescriptionClass) Alloc() MLObjectBoundingBoxOutputDescription {
	rv := objc.Send[MLObjectBoundingBoxOutputDescription](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLObjectBoundingBoxOutputDescription.ConfidenceFeatureName]
//   - [MLObjectBoundingBoxOutputDescription.SetConfidenceFeatureName]
//   - [MLObjectBoundingBoxOutputDescription.CoordinatesFeatureName]
//   - [MLObjectBoundingBoxOutputDescription.SetCoordinatesFeatureName]
//   - [MLObjectBoundingBoxOutputDescription.Format]
//   - [MLObjectBoundingBoxOutputDescription.SetFormat]
//   - [MLObjectBoundingBoxOutputDescription.LabelNames]
//   - [MLObjectBoundingBoxOutputDescription.SetLabelNames]
//
// See: https://developer.apple.com/documentation/CoreML/MLObjectBoundingBoxOutputDescription
type MLObjectBoundingBoxOutputDescription struct {
	objectivec.Object
}

// MLObjectBoundingBoxOutputDescriptionFromID constructs a [MLObjectBoundingBoxOutputDescription] from an objc.ID.
func MLObjectBoundingBoxOutputDescriptionFromID(id objc.ID) MLObjectBoundingBoxOutputDescription {
	return MLObjectBoundingBoxOutputDescription{objectivec.Object{ID: id}}
}

// Ensure MLObjectBoundingBoxOutputDescription implements IMLObjectBoundingBoxOutputDescription.
var _ IMLObjectBoundingBoxOutputDescription = MLObjectBoundingBoxOutputDescription{}

// An interface definition for the [MLObjectBoundingBoxOutputDescription] class.
//
// # Methods
//
//   - [IMLObjectBoundingBoxOutputDescription.ConfidenceFeatureName]
//   - [IMLObjectBoundingBoxOutputDescription.SetConfidenceFeatureName]
//   - [IMLObjectBoundingBoxOutputDescription.CoordinatesFeatureName]
//   - [IMLObjectBoundingBoxOutputDescription.SetCoordinatesFeatureName]
//   - [IMLObjectBoundingBoxOutputDescription.Format]
//   - [IMLObjectBoundingBoxOutputDescription.SetFormat]
//   - [IMLObjectBoundingBoxOutputDescription.LabelNames]
//   - [IMLObjectBoundingBoxOutputDescription.SetLabelNames]
//
// See: https://developer.apple.com/documentation/CoreML/MLObjectBoundingBoxOutputDescription
type IMLObjectBoundingBoxOutputDescription interface {
	objectivec.IObject

	// Topic: Methods

	ConfidenceFeatureName() string
	SetConfidenceFeatureName(value string)
	CoordinatesFeatureName() string
	SetCoordinatesFeatureName(value string)
	Format() int
	SetFormat(value int)
	LabelNames() foundation.INSArray
	SetLabelNames(value foundation.INSArray)
}

// Init initializes the instance.
func (o MLObjectBoundingBoxOutputDescription) Init() MLObjectBoundingBoxOutputDescription {
	rv := objc.Send[MLObjectBoundingBoxOutputDescription](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o MLObjectBoundingBoxOutputDescription) Autorelease() MLObjectBoundingBoxOutputDescription {
	rv := objc.Send[MLObjectBoundingBoxOutputDescription](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLObjectBoundingBoxOutputDescription creates a new MLObjectBoundingBoxOutputDescription instance.
func NewMLObjectBoundingBoxOutputDescription() MLObjectBoundingBoxOutputDescription {
	class := getMLObjectBoundingBoxOutputDescriptionClass()
	rv := objc.Send[MLObjectBoundingBoxOutputDescription](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLObjectBoundingBoxOutputDescription/confidenceFeatureName
func (o MLObjectBoundingBoxOutputDescription) ConfidenceFeatureName() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("confidenceFeatureName"))
	return foundation.NSStringFromID(rv).String()
}
func (o MLObjectBoundingBoxOutputDescription) SetConfidenceFeatureName(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setConfidenceFeatureName:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/CoreML/MLObjectBoundingBoxOutputDescription/coordinatesFeatureName
func (o MLObjectBoundingBoxOutputDescription) CoordinatesFeatureName() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("coordinatesFeatureName"))
	return foundation.NSStringFromID(rv).String()
}
func (o MLObjectBoundingBoxOutputDescription) SetCoordinatesFeatureName(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setCoordinatesFeatureName:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/CoreML/MLObjectBoundingBoxOutputDescription/format
func (o MLObjectBoundingBoxOutputDescription) Format() int {
	rv := objc.Send[int](o.ID, objc.Sel("format"))
	return rv
}
func (o MLObjectBoundingBoxOutputDescription) SetFormat(value int) {
	objc.Send[struct{}](o.ID, objc.Sel("setFormat:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLObjectBoundingBoxOutputDescription/labelNames
func (o MLObjectBoundingBoxOutputDescription) LabelNames() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("labelNames"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (o MLObjectBoundingBoxOutputDescription) SetLabelNames(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setLabelNames:"), value)
}
