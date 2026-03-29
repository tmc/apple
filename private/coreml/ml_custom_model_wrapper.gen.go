// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLCustomModelWrapper] class.
var (
	_MLCustomModelWrapperClass     MLCustomModelWrapperClass
	_MLCustomModelWrapperClassOnce sync.Once
)

func getMLCustomModelWrapperClass() MLCustomModelWrapperClass {
	_MLCustomModelWrapperClassOnce.Do(func() {
		_MLCustomModelWrapperClass = MLCustomModelWrapperClass{class: objc.GetClass("MLCustomModelWrapper")}
	})
	return _MLCustomModelWrapperClass
}

// GetMLCustomModelWrapperClass returns the class object for MLCustomModelWrapper.
func GetMLCustomModelWrapperClass() MLCustomModelWrapperClass {
	return getMLCustomModelWrapperClass()
}

type MLCustomModelWrapperClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLCustomModelWrapperClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLCustomModelWrapperClass) Alloc() MLCustomModelWrapper {
	rv := objc.Send[MLCustomModelWrapper](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLCustomModelWrapper.CustomModel]
//   - [MLCustomModelWrapper.SetCustomModel]
//   - [MLCustomModelWrapper.PredictionFromFeaturesOptionsError]
//   - [MLCustomModelWrapper.PredictionsFromBatchOptionsError]
//   - [MLCustomModelWrapper.InitWithModelDescriptionCustomModelConfiguration]
// See: https://developer.apple.com/documentation/CoreML/MLCustomModelWrapper
type MLCustomModelWrapper struct {
	MLModel
}

// MLCustomModelWrapperFromID constructs a [MLCustomModelWrapper] from an objc.ID.
func MLCustomModelWrapperFromID(id objc.ID) MLCustomModelWrapper {
	return MLCustomModelWrapper{MLModel: MLModelFromID(id)}
}
// Ensure MLCustomModelWrapper implements IMLCustomModelWrapper.
var _ IMLCustomModelWrapper = MLCustomModelWrapper{}

// An interface definition for the [MLCustomModelWrapper] class.
//
// # Methods
//
//   - [IMLCustomModelWrapper.CustomModel]
//   - [IMLCustomModelWrapper.SetCustomModel]
//   - [IMLCustomModelWrapper.PredictionFromFeaturesOptionsError]
//   - [IMLCustomModelWrapper.PredictionsFromBatchOptionsError]
//   - [IMLCustomModelWrapper.InitWithModelDescriptionCustomModelConfiguration]
//
// See: https://developer.apple.com/documentation/CoreML/MLCustomModelWrapper
type IMLCustomModelWrapper interface {
	IMLModel

	// Topic: Methods

	CustomModel() objectivec.Object
	SetCustomModel(value objectivec.Object)
	PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	PredictionsFromBatchOptionsError(batch objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	InitWithModelDescriptionCustomModelConfiguration(description objectivec.IObject, model objectivec.IObject, configuration objectivec.IObject) MLCustomModelWrapper
}

// Init initializes the instance.
func (c MLCustomModelWrapper) Init() MLCustomModelWrapper {
	rv := objc.Send[MLCustomModelWrapper](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MLCustomModelWrapper) Autorelease() MLCustomModelWrapper {
	rv := objc.Send[MLCustomModelWrapper](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLCustomModelWrapper creates a new MLCustomModelWrapper instance.
func NewMLCustomModelWrapper() MLCustomModelWrapper {
	class := getMLCustomModelWrapperClass()
	rv := objc.Send[MLCustomModelWrapper](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initDescriptionOnlyWithSpecification:configuration:error:
func NewCustomModelWrapperDescriptionOnlyWithSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (MLCustomModelWrapper, error) {
	var errorPtr objc.ID
	instance := getMLCustomModelWrapperClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initDescriptionOnlyWithSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLCustomModelWrapper{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLCustomModelWrapperFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initInterfaceAndMetadataWithCompiledArchive:error:
func NewCustomModelWrapperInterfaceAndMetadataWithCompiledArchiveError(archive unsafe.Pointer) (MLCustomModelWrapper, error) {
	var errorPtr objc.ID
	instance := getMLCustomModelWrapperClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initInterfaceAndMetadataWithCompiledArchive:error:"), archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLCustomModelWrapper{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLCustomModelWrapperFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithConfiguration:
func NewCustomModelWrapperWithConfiguration(configuration objectivec.IObject) MLCustomModelWrapper {
	instance := getMLCustomModelWrapperClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	return MLCustomModelWrapperFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithDescription:
func NewCustomModelWrapperWithDescription(description objectivec.IObject) MLCustomModelWrapper {
	instance := getMLCustomModelWrapperClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:"), description)
	return MLCustomModelWrapperFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithDescription:configuration:
func NewCustomModelWrapperWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLCustomModelWrapper {
	instance := getMLCustomModelWrapperClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return MLCustomModelWrapperFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLCustomModelWrapper/initWithModelDescription:customModel:configuration:
func NewCustomModelWrapperWithModelDescriptionCustomModelConfiguration(description objectivec.IObject, model objectivec.IObject, configuration objectivec.IObject) MLCustomModelWrapper {
	instance := getMLCustomModelWrapperClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithModelDescription:customModel:configuration:"), description, model, configuration)
	return MLCustomModelWrapperFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:
func NewCustomModelWrapperWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLCustomModelWrapper {
	instance := getMLCustomModelWrapperClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return MLCustomModelWrapperFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLCustomModelWrapper/predictionFromFeatures:options:error:
func (c MLCustomModelWrapper) PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](c.ID, objc.Sel("predictionFromFeatures:options:error:"), features, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLCustomModelWrapper/predictionsFromBatch:options:error:
func (c MLCustomModelWrapper) PredictionsFromBatchOptionsError(batch objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](c.ID, objc.Sel("predictionsFromBatch:options:error:"), batch, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLCustomModelWrapper/initWithModelDescription:customModel:configuration:
func (c MLCustomModelWrapper) InitWithModelDescriptionCustomModelConfiguration(description objectivec.IObject, model objectivec.IObject, configuration objectivec.IObject) MLCustomModelWrapper {
	rv := objc.Send[MLCustomModelWrapper](c.ID, objc.Sel("initWithModelDescription:customModel:configuration:"), description, model, configuration)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLCustomModelWrapper/customModel
func (c MLCustomModelWrapper) CustomModel() objectivec.Object {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("customModel"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (c MLCustomModelWrapper) SetCustomModel(value objectivec.Object) {
	objc.Send[struct{}](c.ID, objc.Sel("setCustomModel:"), value)
}

