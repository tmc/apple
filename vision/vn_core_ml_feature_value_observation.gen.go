// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coreml"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [VNCoreMLFeatureValueObservation] class.
var (
	_VNCoreMLFeatureValueObservationClass     VNCoreMLFeatureValueObservationClass
	_VNCoreMLFeatureValueObservationClassOnce sync.Once
)

func getVNCoreMLFeatureValueObservationClass() VNCoreMLFeatureValueObservationClass {
	_VNCoreMLFeatureValueObservationClassOnce.Do(func() {
		_VNCoreMLFeatureValueObservationClass = VNCoreMLFeatureValueObservationClass{class: objc.GetClass("VNCoreMLFeatureValueObservation")}
	})
	return _VNCoreMLFeatureValueObservationClass
}

// GetVNCoreMLFeatureValueObservationClass returns the class object for VNCoreMLFeatureValueObservation.
func GetVNCoreMLFeatureValueObservationClass() VNCoreMLFeatureValueObservationClass {
	return getVNCoreMLFeatureValueObservationClass()
}

type VNCoreMLFeatureValueObservationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNCoreMLFeatureValueObservationClass) Alloc() VNCoreMLFeatureValueObservation {
	rv := objc.Send[VNCoreMLFeatureValueObservation](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}







// An object that represents a collection of key-value information that a Core
// ML image-analysis request produces.
//
// # Overview
// 
// This type of observation results from performing a [VNCoreMLRequest] image
// analysis with a Core ML model whose role is prediction rather than
// classification or image-to-image processing.
// 
// Vision infers that an [MLModel] object is a predictor model if that model
// predicts multiple features. You can tell that a model predicts multiple
// features when its [VNCoreMLFeatureValueObservation.ModelDescription] object has a `nil` value for its
// [VNCoreMLFeatureValueObservation.PredictedFeatureName] property, or when it inserts its output in an
// [VNCoreMLFeatureValueObservation.OutputDescriptionsByName] dictionary.
//
// [MLModel]: https://developer.apple.com/documentation/CoreML/MLModel
//
// # Obtaining Feature Values
//
//   - [VNCoreMLFeatureValueObservation.FeatureValue]: The feature result of a [VNCoreMLRequest](<doc://Vision/documentation/Vision/VNCoreMLRequest>) that outputs neither a classification nor an image.
//   - [VNCoreMLFeatureValueObservation.FeatureName]: The name used in the model description of the CoreML model that produced this observation.
//
// See: https://developer.apple.com/documentation/Vision/VNCoreMLFeatureValueObservation
type VNCoreMLFeatureValueObservation struct {
	VNObservation
}

// VNCoreMLFeatureValueObservationFromID constructs a [VNCoreMLFeatureValueObservation] from an objc.ID.
//
// An object that represents a collection of key-value information that a Core
// ML image-analysis request produces.
func VNCoreMLFeatureValueObservationFromID(id objc.ID) VNCoreMLFeatureValueObservation {
	return VNCoreMLFeatureValueObservation{VNObservation: VNObservationFromID(id)}
}
// NOTE: VNCoreMLFeatureValueObservation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [VNCoreMLFeatureValueObservation] class.
//
// # Obtaining Feature Values
//
//   - [IVNCoreMLFeatureValueObservation.FeatureValue]: The feature result of a [VNCoreMLRequest](<doc://Vision/documentation/Vision/VNCoreMLRequest>) that outputs neither a classification nor an image.
//   - [IVNCoreMLFeatureValueObservation.FeatureName]: The name used in the model description of the CoreML model that produced this observation.
//
// See: https://developer.apple.com/documentation/Vision/VNCoreMLFeatureValueObservation
type IVNCoreMLFeatureValueObservation interface {
	IVNObservation

	// Topic: Obtaining Feature Values

	// The feature result of a [VNCoreMLRequest](<doc://Vision/documentation/Vision/VNCoreMLRequest>) that outputs neither a classification nor an image.
	FeatureValue() coreml.MLFeatureValue
	// The name used in the model description of the CoreML model that produced this observation.
	FeatureName() string

	// Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view.
	ModelDescription() coreml.MLModelDescription
	SetModelDescription(value coreml.MLModelDescription)
	// A dictionary of output feature descriptions, which the model keys by the output’s name.
	OutputDescriptionsByName() coreml.MLFeatureDescription
	SetOutputDescriptionsByName(value coreml.MLFeatureDescription)
	// The name of the primary prediction feature output description.
	PredictedFeatureName() string
	SetPredictedFeatureName(value string)
	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (c VNCoreMLFeatureValueObservation) Init() VNCoreMLFeatureValueObservation {
	rv := objc.Send[VNCoreMLFeatureValueObservation](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c VNCoreMLFeatureValueObservation) Autorelease() VNCoreMLFeatureValueObservation {
	rv := objc.Send[VNCoreMLFeatureValueObservation](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNCoreMLFeatureValueObservation creates a new VNCoreMLFeatureValueObservation instance.
func NewVNCoreMLFeatureValueObservation() VNCoreMLFeatureValueObservation {
	class := getVNCoreMLFeatureValueObservationClass()
	rv := objc.Send[VNCoreMLFeatureValueObservation](objc.ID(class.class), objc.Sel("new"))
	return rv
}









func (c VNCoreMLFeatureValueObservation) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}












// The feature result of a [VNCoreMLRequest] that outputs neither a
// classification nor an image.
//
// # Discussion
// 
// Refer to [Core ML] documentation and the model itself to learn about proper
// handling of the content.
//
// [Core ML]: https://developer.apple.com/documentation/CoreML
//
// See: https://developer.apple.com/documentation/Vision/VNCoreMLFeatureValueObservation/featureValue
func (c VNCoreMLFeatureValueObservation) FeatureValue() coreml.MLFeatureValue {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("featureValue"))
	return coreml.MLFeatureValueFromID(objc.ID(rv))
}



// The name used in the model description of the CoreML model that produced
// this observation.
//
// See: https://developer.apple.com/documentation/Vision/VNCoreMLFeatureValueObservation/featureName
func (c VNCoreMLFeatureValueObservation) FeatureName() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("featureName"))
	return foundation.NSStringFromID(rv).String()
}



// Model information you use at runtime during development, which Xcode also
// displays in its Core ML model editor view.
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/modelDescription
func (c VNCoreMLFeatureValueObservation) ModelDescription() coreml.MLModelDescription {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("modelDescription"))
	return coreml.MLModelDescriptionFromID(objc.ID(rv))
}
func (c VNCoreMLFeatureValueObservation) SetModelDescription(value coreml.MLModelDescription) {
	objc.Send[struct{}](c.ID, objc.Sel("setModelDescription:"), value)
}



// A dictionary of output feature descriptions, which the model keys by the
// output’s name.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/outputDescriptionsByName
func (c VNCoreMLFeatureValueObservation) OutputDescriptionsByName() coreml.MLFeatureDescription {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("outputDescriptionsByName"))
	return coreml.MLFeatureDescriptionFromID(objc.ID(rv))
}
func (c VNCoreMLFeatureValueObservation) SetOutputDescriptionsByName(value coreml.MLFeatureDescription) {
	objc.Send[struct{}](c.ID, objc.Sel("setOutputDescriptionsByName:"), value)
}



// The name of the primary prediction feature output description.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/predictedFeatureName
func (c VNCoreMLFeatureValueObservation) PredictedFeatureName() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("predictedFeatureName"))
	return foundation.NSStringFromID(rv).String()
}
func (c VNCoreMLFeatureValueObservation) SetPredictedFeatureName(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setPredictedFeatureName:"), objc.String(value))
}



























