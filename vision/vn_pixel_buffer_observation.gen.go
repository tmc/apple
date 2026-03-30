// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"

	"github.com/tmc/apple/coreml"
	"github.com/tmc/apple/corevideo"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [VNPixelBufferObservation] class.
var (
	_VNPixelBufferObservationClass     VNPixelBufferObservationClass
	_VNPixelBufferObservationClassOnce sync.Once
)

func getVNPixelBufferObservationClass() VNPixelBufferObservationClass {
	_VNPixelBufferObservationClassOnce.Do(func() {
		_VNPixelBufferObservationClass = VNPixelBufferObservationClass{class: objc.GetClass("VNPixelBufferObservation")}
	})
	return _VNPixelBufferObservationClass
}

// GetVNPixelBufferObservationClass returns the class object for VNPixelBufferObservation.
func GetVNPixelBufferObservationClass() VNPixelBufferObservationClass {
	return getVNPixelBufferObservationClass()
}

type VNPixelBufferObservationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VNPixelBufferObservationClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNPixelBufferObservationClass) Alloc() VNPixelBufferObservation {
	rv := objc.Send[VNPixelBufferObservation](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An object that represents an image that an image-analysis request produces.
//
// # Overview
//
// This type of observation results from performing a [VNCoreMLRequest] image
// analysis with a Core ML model that has an image-to-image processing role.
// For example, this observation might result from a model that analyzes the
// style of one image and then transfers that style to a different image.
//
// Vision infers that an [MLModel] object is an image-to-image model if that
// model includes an image. Its [VNPixelBufferObservation.ModelDescription] object includes an
// image-typed feature description in its [VNPixelBufferObservation.OutputDescriptionsByName]
// dictionary.
//
// # Parsing Observation Content
//
//   - [VNPixelBufferObservation.PixelBuffer]: The image that results from a request with image output.
//   - [VNPixelBufferObservation.FeatureName]: A feature name that the CoreML model defines.
//
// See: https://developer.apple.com/documentation/Vision/VNPixelBufferObservation
//
// [MLModel]: https://developer.apple.com/documentation/CoreML/MLModel
type VNPixelBufferObservation struct {
	VNObservation
}

// VNPixelBufferObservationFromID constructs a [VNPixelBufferObservation] from an objc.ID.
//
// An object that represents an image that an image-analysis request produces.
func VNPixelBufferObservationFromID(id objc.ID) VNPixelBufferObservation {
	return VNPixelBufferObservation{VNObservation: VNObservationFromID(id)}
}

// NOTE: VNPixelBufferObservation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNPixelBufferObservation] class.
//
// # Parsing Observation Content
//
//   - [IVNPixelBufferObservation.PixelBuffer]: The image that results from a request with image output.
//   - [IVNPixelBufferObservation.FeatureName]: A feature name that the CoreML model defines.
//
// See: https://developer.apple.com/documentation/Vision/VNPixelBufferObservation
type IVNPixelBufferObservation interface {
	IVNObservation

	// Topic: Parsing Observation Content

	// The image that results from a request with image output.
	PixelBuffer() corevideo.CVImageBufferRef
	// A feature name that the CoreML model defines.
	FeatureName() string

	// Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view.
	ModelDescription() coreml.MLModelDescription
	SetModelDescription(value coreml.MLModelDescription)
	// A dictionary of output feature descriptions, which the model keys by the output’s name.
	OutputDescriptionsByName() coreml.MLFeatureDescription
	SetOutputDescriptionsByName(value coreml.MLFeatureDescription)
}

// Init initializes the instance.
func (p VNPixelBufferObservation) Init() VNPixelBufferObservation {
	rv := objc.Send[VNPixelBufferObservation](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p VNPixelBufferObservation) Autorelease() VNPixelBufferObservation {
	rv := objc.Send[VNPixelBufferObservation](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNPixelBufferObservation creates a new VNPixelBufferObservation instance.
func NewVNPixelBufferObservation() VNPixelBufferObservation {
	class := getVNPixelBufferObservationClass()
	rv := objc.Send[VNPixelBufferObservation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The image that results from a request with image output.
//
// # Discussion
//
// [VNCoreMLRequest] produces observations that contain images in pixel buffer
// format. The confidence level is always `1.0`.
//
// See: https://developer.apple.com/documentation/Vision/VNPixelBufferObservation/pixelBuffer
func (p VNPixelBufferObservation) PixelBuffer() corevideo.CVImageBufferRef {
	rv := objc.Send[corevideo.CVImageBufferRef](p.ID, objc.Sel("pixelBuffer"))
	return corevideo.CVImageBufferRef(rv)
}

// A feature name that the CoreML model defines.
//
// # Discussion
//
// This value is nil if the observation isn’t the result of a
// [VNCoreMLRequest] operation.
//
// See: https://developer.apple.com/documentation/Vision/VNPixelBufferObservation/featureName
func (p VNPixelBufferObservation) FeatureName() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("featureName"))
	return foundation.NSStringFromID(rv).String()
}

// Model information you use at runtime during development, which Xcode also
// displays in its Core ML model editor view.
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/modelDescription
func (p VNPixelBufferObservation) ModelDescription() coreml.MLModelDescription {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("modelDescription"))
	return coreml.MLModelDescriptionFromID(objc.ID(rv))
}
func (p VNPixelBufferObservation) SetModelDescription(value coreml.MLModelDescription) {
	objc.Send[struct{}](p.ID, objc.Sel("setModelDescription:"), value)
}

// A dictionary of output feature descriptions, which the model keys by the
// output’s name.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/outputDescriptionsByName
func (p VNPixelBufferObservation) OutputDescriptionsByName() coreml.MLFeatureDescription {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("outputDescriptionsByName"))
	return coreml.MLFeatureDescriptionFromID(objc.ID(rv))
}
func (p VNPixelBufferObservation) SetOutputDescriptionsByName(value coreml.MLFeatureDescription) {
	objc.Send[struct{}](p.ID, objc.Sel("setOutputDescriptionsByName:"), value)
}
