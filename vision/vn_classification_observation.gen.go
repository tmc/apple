// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coreml"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [VNClassificationObservation] class.
var (
	_VNClassificationObservationClass     VNClassificationObservationClass
	_VNClassificationObservationClassOnce sync.Once
)

func getVNClassificationObservationClass() VNClassificationObservationClass {
	_VNClassificationObservationClassOnce.Do(func() {
		_VNClassificationObservationClass = VNClassificationObservationClass{class: objc.GetClass("VNClassificationObservation")}
	})
	return _VNClassificationObservationClass
}

// GetVNClassificationObservationClass returns the class object for VNClassificationObservation.
func GetVNClassificationObservationClass() VNClassificationObservationClass {
	return getVNClassificationObservationClass()
}

type VNClassificationObservationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNClassificationObservationClass) Alloc() VNClassificationObservation {
	rv := objc.Send[VNClassificationObservation](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// An object that represents classification information that an image-analysis
// request produces.
//
// # Overview
// 
// This type of observation results from performing a [VNCoreMLRequest] image
// analysis with a Core ML model whose role is classification (rather than
// prediction or image-to-image processing). Vision infers that an [MLModel]
// object is a classifier model if that model predicts a single feature. That
// is, the model’s [VNClassificationObservation.ModelDescription] object has a non-`nil` value for its
// [VNClassificationObservation.PredictedFeatureName] property.
//
// [MLModel]: https://developer.apple.com/documentation/CoreML/MLModel
//
// # Determining Classification
//
//   - [VNClassificationObservation.Identifier]: Classification label identifying the type of observation.
//
// # Measuring Confidence and Precision
//
//   - [VNClassificationObservation.HasPrecisionRecallCurve]: A Boolean variable indicating whether the observation contains precision and recall curves.
//   - [VNClassificationObservation.HasMinimumPrecisionForRecall]: Determines whether the observation for a specific recall has a minimum precision value.
//   - [VNClassificationObservation.HasMinimumRecallForPrecision]: Determines whether the observation for a specific precision has a minimum recall value.
//
// See: https://developer.apple.com/documentation/Vision/VNClassificationObservation
type VNClassificationObservation struct {
	VNObservation
}

// VNClassificationObservationFromID constructs a [VNClassificationObservation] from an objc.ID.
//
// An object that represents classification information that an image-analysis
// request produces.
func VNClassificationObservationFromID(id objc.ID) VNClassificationObservation {
	return VNClassificationObservation{VNObservation: VNObservationFromID(id)}
}
// NOTE: VNClassificationObservation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNClassificationObservation] class.
//
// # Determining Classification
//
//   - [IVNClassificationObservation.Identifier]: Classification label identifying the type of observation.
//
// # Measuring Confidence and Precision
//
//   - [IVNClassificationObservation.HasPrecisionRecallCurve]: A Boolean variable indicating whether the observation contains precision and recall curves.
//   - [IVNClassificationObservation.HasMinimumPrecisionForRecall]: Determines whether the observation for a specific recall has a minimum precision value.
//   - [IVNClassificationObservation.HasMinimumRecallForPrecision]: Determines whether the observation for a specific precision has a minimum recall value.
//
// See: https://developer.apple.com/documentation/Vision/VNClassificationObservation
type IVNClassificationObservation interface {
	IVNObservation

	// Topic: Determining Classification

	// Classification label identifying the type of observation.
	Identifier() string

	// Topic: Measuring Confidence and Precision

	// A Boolean variable indicating whether the observation contains precision and recall curves.
	HasPrecisionRecallCurve() bool
	// Determines whether the observation for a specific recall has a minimum precision value.
	HasMinimumPrecisionForRecall(minimumPrecision float32, recall float32) bool
	// Determines whether the observation for a specific precision has a minimum recall value.
	HasMinimumRecallForPrecision(minimumRecall float32, precision float32) bool

	// Model information you use at runtime during development, which Xcode also displays in its Core ML model editor view.
	ModelDescription() coreml.MLModelDescription
	SetModelDescription(value coreml.MLModelDescription)
	// The name of the primary prediction feature output description.
	PredictedFeatureName() string
	SetPredictedFeatureName(value string)
}

// Init initializes the instance.
func (c VNClassificationObservation) Init() VNClassificationObservation {
	rv := objc.Send[VNClassificationObservation](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c VNClassificationObservation) Autorelease() VNClassificationObservation {
	rv := objc.Send[VNClassificationObservation](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNClassificationObservation creates a new VNClassificationObservation instance.
func NewVNClassificationObservation() VNClassificationObservation {
	class := getVNClassificationObservationClass()
	rv := objc.Send[VNClassificationObservation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Determines whether the observation for a specific recall has a minimum
// precision value.
//
// minimumPrecision: The minimum percentage of classification results that are relevant.
//
// recall: The percentage of relevant results that the algorithm correctly classified.
//
// # Return Value
// 
// A Boolean indicating whether or not this classification observation
// provides a minimum percentage of relevant results that meet the desired
// recall criterion.
//
// See: https://developer.apple.com/documentation/Vision/VNClassificationObservation/hasMinimumPrecision(_:forRecall:)
func (c VNClassificationObservation) HasMinimumPrecisionForRecall(minimumPrecision float32, recall float32) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("hasMinimumPrecision:forRecall:"), minimumPrecision, recall)
	return rv
}
// Determines whether the observation for a specific precision has a minimum
// recall value.
//
// minimumRecall: The minimum percentage of relevant results that the algorithm correctly
// classified.
//
// precision: The percentage of classification results that are relevant.
//
// # Return Value
// 
// A Boolean indicating whether or not this classification observation
// provides a minimum percentage of relevant results that meet the desired
// precision criterion.
//
// See: https://developer.apple.com/documentation/Vision/VNClassificationObservation/hasMinimumRecall(_:forPrecision:)
func (c VNClassificationObservation) HasMinimumRecallForPrecision(minimumRecall float32, precision float32) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("hasMinimumRecall:forPrecision:"), minimumRecall, precision)
	return rv
}

// Classification label identifying the type of observation.
//
// # Discussion
// 
// An example classification could be a string like `cat` or `hotdog`. The
// model used for the classification defines the domain of strings that may
// result. Usually, these strings are unlocalized technical labels not meant
// for direct presentation to the end user.
//
// See: https://developer.apple.com/documentation/Vision/VNClassificationObservation/identifier
func (c VNClassificationObservation) Identifier() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("identifier"))
	return foundation.NSStringFromID(rv).String()
}
// A Boolean variable indicating whether the observation contains precision
// and recall curves.
//
// # Discussion
// 
// Precision refers to the percentage of your classification results that are
// relevant, while recall refers to the percentage of total relevant results
// correctly classified.
// 
// If this property is [true], then you can call precision and recall-related
// methods in this observation. If this property is [false], then the
// precision and recall-related methods won’t return meaningful data.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Vision/VNClassificationObservation/hasPrecisionRecallCurve
func (c VNClassificationObservation) HasPrecisionRecallCurve() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("hasPrecisionRecallCurve"))
	return rv
}
// Model information you use at runtime during development, which Xcode also
// displays in its Core ML model editor view.
//
// See: https://developer.apple.com/documentation/CoreML/MLModel/modelDescription
func (c VNClassificationObservation) ModelDescription() coreml.MLModelDescription {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("modelDescription"))
	return coreml.MLModelDescriptionFromID(objc.ID(rv))
}
func (c VNClassificationObservation) SetModelDescription(value coreml.MLModelDescription) {
	objc.Send[struct{}](c.ID, objc.Sel("setModelDescription:"), value)
}
// The name of the primary prediction feature output description.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/predictedFeatureName
func (c VNClassificationObservation) PredictedFeatureName() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("predictedFeatureName"))
	return foundation.NSStringFromID(rv).String()
}
func (c VNClassificationObservation) SetPredictedFeatureName(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setPredictedFeatureName:"), objc.String(value))
}

