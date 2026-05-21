// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
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

// Class returns the underlying Objective-C class pointer.
func (vc VNCoreMLFeatureValueObservationClass) Class() objc.Class {
	return vc.class
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
// features when its [modelDescription] object has a `nil` value for its
// [predictedFeatureName] property, or when it inserts its output in an
// [outputDescriptionsByName] dictionary.
//
// # Obtaining Feature Values
//
//   - [VNCoreMLFeatureValueObservation.FeatureValue]: The feature result of a [VNCoreMLRequest](<doc://Vision/documentation/Vision/VNCoreMLRequest>) that outputs neither a classification nor an image.
//   - [VNCoreMLFeatureValueObservation.FeatureName]: The name used in the model description of the CoreML model that produced this observation.
//
// See: https://developer.apple.com/documentation/Vision/VNCoreMLFeatureValueObservation
//
// [MLModel]: https://developer.apple.com/documentation/CoreML/MLModel
// [modelDescription]: https://developer.apple.com/documentation/CoreML/MLModel/modelDescription
// [outputDescriptionsByName]: https://developer.apple.com/documentation/CoreML/MLModelDescription/outputDescriptionsByName
// [predictedFeatureName]: https://developer.apple.com/documentation/CoreML/MLModelDescription/predictedFeatureName
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
	FeatureValue() objectivec.IObject
	// The name used in the model description of the CoreML model that produced this observation.
	FeatureName() string
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

// See: https://developer.apple.com/documentation/Vision/VNObservation/init(coder:)
func NewCoreMLFeatureValueObservationWithCoder(coder foundation.INSCoder) VNCoreMLFeatureValueObservation {
	instance := getVNCoreMLFeatureValueObservationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return VNCoreMLFeatureValueObservationFromID(rv)
}

// The feature result of a [VNCoreMLRequest] that outputs neither a
// classification nor an image.
//
// # Discussion
//
// Refer to [Core ML] documentation and the model itself to learn about proper
// handling of the content.
//
// See: https://developer.apple.com/documentation/Vision/VNCoreMLFeatureValueObservation/featureValue
//
// [Core ML]: https://developer.apple.com/documentation/CoreML
func (c VNCoreMLFeatureValueObservation) FeatureValue() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("featureValue"))
	return objectivec.Object{ID: rv}
}

// The name used in the model description of the CoreML model that produced
// this observation.
//
// See: https://developer.apple.com/documentation/Vision/VNCoreMLFeatureValueObservation/featureName
func (c VNCoreMLFeatureValueObservation) FeatureName() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("featureName"))
	return foundation.NSStringFromID(rv).String()
}
