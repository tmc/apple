// Code generated from Apple documentation for Vision. DO NOT EDIT.

package vision

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coreml"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VNCoreMLModel] class.
var (
	_VNCoreMLModelClass     VNCoreMLModelClass
	_VNCoreMLModelClassOnce sync.Once
)

func getVNCoreMLModelClass() VNCoreMLModelClass {
	_VNCoreMLModelClassOnce.Do(func() {
		_VNCoreMLModelClass = VNCoreMLModelClass{class: objc.GetClass("VNCoreMLModel")}
	})
	return _VNCoreMLModelClass
}

// GetVNCoreMLModelClass returns the class object for VNCoreMLModel.
func GetVNCoreMLModelClass() VNCoreMLModelClass {
	return getVNCoreMLModelClass()
}

type VNCoreMLModelClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (vc VNCoreMLModelClass) Alloc() VNCoreMLModel {
	rv := objc.Send[VNCoreMLModel](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// A container for the model to use with Vision requests.
//
// # Overview
// 
// A [Core ML] model encapsulates the information trained from a data set used
// to drive Vision recognition requests. See [Getting a Core ML Model] for
// instructions on training your own model. Once you train the model, use this
// class to initialize a [VNCoreMLRequest] for identification.
//
// [Core ML]: https://developer.apple.com/documentation/CoreML
// [Getting a Core ML Model]: https://developer.apple.com/documentation/CoreML/getting-a-core-ml-model
//
// # Providing Features
//
//   - [VNCoreMLModel.FeatureProvider]: An optional object to support inputs outside Vision.
//   - [VNCoreMLModel.SetFeatureProvider]
//   - [VNCoreMLModel.InputImageFeatureName]: The name of the feature value that Vision sets from the request handler.
//   - [VNCoreMLModel.SetInputImageFeatureName]
//
// See: https://developer.apple.com/documentation/Vision/VNCoreMLModel
type VNCoreMLModel struct {
	objectivec.Object
}

// VNCoreMLModelFromID constructs a [VNCoreMLModel] from an objc.ID.
//
// A container for the model to use with Vision requests.
func VNCoreMLModelFromID(id objc.ID) VNCoreMLModel {
	return VNCoreMLModel{objectivec.Object{ID: id}}
}
// NOTE: VNCoreMLModel adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VNCoreMLModel] class.
//
// # Providing Features
//
//   - [IVNCoreMLModel.FeatureProvider]: An optional object to support inputs outside Vision.
//   - [IVNCoreMLModel.SetFeatureProvider]
//   - [IVNCoreMLModel.InputImageFeatureName]: The name of the feature value that Vision sets from the request handler.
//   - [IVNCoreMLModel.SetInputImageFeatureName]
//
// See: https://developer.apple.com/documentation/Vision/VNCoreMLModel
type IVNCoreMLModel interface {
	objectivec.IObject

	// Topic: Providing Features

	// An optional object to support inputs outside Vision.
	FeatureProvider() coreml.MLFeatureProvider
	SetFeatureProvider(value coreml.MLFeatureProvider)
	// The name of the feature value that Vision sets from the request handler.
	InputImageFeatureName() string
	SetInputImageFeatureName(value string)

	// The model to base the image analysis request on.
	Model() IVNCoreMLModel
	SetModel(value IVNCoreMLModel)
}

// Init initializes the instance.
func (c VNCoreMLModel) Init() VNCoreMLModel {
	rv := objc.Send[VNCoreMLModel](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c VNCoreMLModel) Autorelease() VNCoreMLModel {
	rv := objc.Send[VNCoreMLModel](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewVNCoreMLModel creates a new VNCoreMLModel instance.
func NewVNCoreMLModel() VNCoreMLModel {
	class := getVNCoreMLModelClass()
	rv := objc.Send[VNCoreMLModel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a model container to use with a Core ML request.
//
// model: The model to create the model container from.
//
// # Discussion
// 
// This method may fail if the framework doesn’t support the Core ML model.
// For example, a model that doesn’t accept an image as any of its inputs
// will yield an [ErrorInvalidModel] error.
//
// See: https://developer.apple.com/documentation/Vision/VNCoreMLModel/init(for:)
func NewCoreMLModelForMLModelError(model coreml.MLModel) (VNCoreMLModel, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getVNCoreMLModelClass().class), objc.Sel("modelForMLModel:error:"), model, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return VNCoreMLModelFromID(rv), foundation.NSErrorFrom(errorPtr)
	}
	return VNCoreMLModelFromID(rv), nil
}

// An optional object to support inputs outside Vision.
//
// # Discussion
// 
// This optional object conforms to the [MLFeatureProvider] protocol that the
// model uses to predict inputs that are not supplied by Vision. Vision
// provides the MLModel with the image for the [InputImageFeatureName] via the
// [VNRequestHandler].
// 
// A feature provider is necessary for models that have more than one required
// input. Models with only one image input won’t use the feature provider.
//
// [MLFeatureProvider]: https://developer.apple.com/documentation/CoreML/MLFeatureProvider
//
// See: https://developer.apple.com/documentation/Vision/VNCoreMLModel/featureProvider
func (c VNCoreMLModel) FeatureProvider() coreml.MLFeatureProvider {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("featureProvider"))
	return coreml.MLFeatureProviderObjectFromID(rv)
}
func (c VNCoreMLModel) SetFeatureProvider(value coreml.MLFeatureProvider) {
	objc.Send[struct{}](c.ID, objc.Sel("setFeatureProvider:"), value)
}
// The name of the feature value that Vision sets from the request handler.
//
// # Discussion
// 
// By default, Vision uses the first input found, but you can manually set
// that input to another [FeatureName] instead.
//
// See: https://developer.apple.com/documentation/Vision/VNCoreMLModel/inputImageFeatureName
func (c VNCoreMLModel) InputImageFeatureName() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("inputImageFeatureName"))
	return foundation.NSStringFromID(rv).String()
}
func (c VNCoreMLModel) SetInputImageFeatureName(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setInputImageFeatureName:"), objc.String(value))
}
// The model to base the image analysis request on.
//
// See: https://developer.apple.com/documentation/vision/vncoremlrequest/model
func (c VNCoreMLModel) Model() IVNCoreMLModel {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("model"))
	return VNCoreMLModelFromID(objc.ID(rv))
}
func (c VNCoreMLModel) SetModel(value IVNCoreMLModel) {
	objc.Send[struct{}](c.ID, objc.Sel("setModel:"), value)
}

