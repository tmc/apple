// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corevideo"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLAppleImageFeatureExtractor] class.
var (
	_MLAppleImageFeatureExtractorClass     MLAppleImageFeatureExtractorClass
	_MLAppleImageFeatureExtractorClassOnce sync.Once
)

func getMLAppleImageFeatureExtractorClass() MLAppleImageFeatureExtractorClass {
	_MLAppleImageFeatureExtractorClassOnce.Do(func() {
		_MLAppleImageFeatureExtractorClass = MLAppleImageFeatureExtractorClass{class: objc.GetClass("MLAppleImageFeatureExtractor")}
	})
	return _MLAppleImageFeatureExtractorClass
}

// GetMLAppleImageFeatureExtractorClass returns the class object for MLAppleImageFeatureExtractor.
func GetMLAppleImageFeatureExtractorClass() MLAppleImageFeatureExtractorClass {
	return getMLAppleImageFeatureExtractorClass()
}

type MLAppleImageFeatureExtractorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLAppleImageFeatureExtractorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLAppleImageFeatureExtractorClass) Alloc() MLAppleImageFeatureExtractor {
	rv := objc.Send[MLAppleImageFeatureExtractor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLAppleImageFeatureExtractor.ComputeScenePrintFeaturesHandleUseCPUOnlyError]
//   - [MLAppleImageFeatureExtractor.FeatureValueFromObjectPrintKeyShape]
//   - [MLAppleImageFeatureExtractor.FeatureValueFromScenePrintElementSize]
//   - [MLAppleImageFeatureExtractor.Parameters]
//   - [MLAppleImageFeatureExtractor.PredictionFromFeaturesOptionsError]
//   - [MLAppleImageFeatureExtractor.InitWithParametersModelDescriptionFeatureExtractorTypeConfigurationError]
// See: https://developer.apple.com/documentation/CoreML/MLAppleImageFeatureExtractor
type MLAppleImageFeatureExtractor struct {
	MLModel
}

// MLAppleImageFeatureExtractorFromID constructs a [MLAppleImageFeatureExtractor] from an objc.ID.
func MLAppleImageFeatureExtractorFromID(id objc.ID) MLAppleImageFeatureExtractor {
	return MLAppleImageFeatureExtractor{MLModel: MLModelFromID(id)}
}
// Ensure MLAppleImageFeatureExtractor implements IMLAppleImageFeatureExtractor.
var _ IMLAppleImageFeatureExtractor = MLAppleImageFeatureExtractor{}

// An interface definition for the [MLAppleImageFeatureExtractor] class.
//
// # Methods
//
//   - [IMLAppleImageFeatureExtractor.ComputeScenePrintFeaturesHandleUseCPUOnlyError]
//   - [IMLAppleImageFeatureExtractor.FeatureValueFromObjectPrintKeyShape]
//   - [IMLAppleImageFeatureExtractor.FeatureValueFromScenePrintElementSize]
//   - [IMLAppleImageFeatureExtractor.Parameters]
//   - [IMLAppleImageFeatureExtractor.PredictionFromFeaturesOptionsError]
//   - [IMLAppleImageFeatureExtractor.InitWithParametersModelDescriptionFeatureExtractorTypeConfigurationError]
//
// See: https://developer.apple.com/documentation/CoreML/MLAppleImageFeatureExtractor
type IMLAppleImageFeatureExtractor interface {
	IMLModel

	// Topic: Methods

	ComputeScenePrintFeaturesHandleUseCPUOnlyError(features corevideo.CVImageBufferRef, handle objectivec.IObject, cPUOnly bool) (objectivec.IObject, error)
	FeatureValueFromObjectPrintKeyShape(print_ objectivec.IObject, key objectivec.IObject, shape objectivec.IObject) objectivec.IObject
	FeatureValueFromScenePrintElementSize(print_ objectivec.IObject, size uint64) objectivec.IObject
	Parameters() IMLAppleImageFeatureExtractorParameters
	PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	InitWithParametersModelDescriptionFeatureExtractorTypeConfigurationError(parameters objectivec.IObject, description objectivec.IObject, type_ int, configuration objectivec.IObject) (MLAppleImageFeatureExtractor, error)
}

// Init initializes the instance.
func (a MLAppleImageFeatureExtractor) Init() MLAppleImageFeatureExtractor {
	rv := objc.Send[MLAppleImageFeatureExtractor](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a MLAppleImageFeatureExtractor) Autorelease() MLAppleImageFeatureExtractor {
	rv := objc.Send[MLAppleImageFeatureExtractor](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLAppleImageFeatureExtractor creates a new MLAppleImageFeatureExtractor instance.
func NewMLAppleImageFeatureExtractor() MLAppleImageFeatureExtractor {
	class := getMLAppleImageFeatureExtractorClass()
	rv := objc.Send[MLAppleImageFeatureExtractor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initDescriptionOnlyWithSpecification:configuration:error:
func NewAppleImageFeatureExtractorDescriptionOnlyWithSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (MLAppleImageFeatureExtractor, error) {
	var errorPtr objc.ID
	instance := getMLAppleImageFeatureExtractorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initDescriptionOnlyWithSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleImageFeatureExtractor{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleImageFeatureExtractorFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initInterfaceAndMetadataWithCompiledArchive:error:
func NewAppleImageFeatureExtractorInterfaceAndMetadataWithCompiledArchiveError(archive unsafe.Pointer) (MLAppleImageFeatureExtractor, error) {
	var errorPtr objc.ID
	instance := getMLAppleImageFeatureExtractorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initInterfaceAndMetadataWithCompiledArchive:error:"), archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleImageFeatureExtractor{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleImageFeatureExtractorFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithConfiguration:
func NewAppleImageFeatureExtractorWithConfiguration(configuration objectivec.IObject) MLAppleImageFeatureExtractor {
	instance := getMLAppleImageFeatureExtractorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	return MLAppleImageFeatureExtractorFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithDescription:
func NewAppleImageFeatureExtractorWithDescription(description objectivec.IObject) MLAppleImageFeatureExtractor {
	instance := getMLAppleImageFeatureExtractorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:"), description)
	return MLAppleImageFeatureExtractorFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithDescription:configuration:
func NewAppleImageFeatureExtractorWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLAppleImageFeatureExtractor {
	instance := getMLAppleImageFeatureExtractorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return MLAppleImageFeatureExtractorFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:
func NewAppleImageFeatureExtractorWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLAppleImageFeatureExtractor {
	instance := getMLAppleImageFeatureExtractorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return MLAppleImageFeatureExtractorFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLAppleImageFeatureExtractor/initWithParameters:modelDescription:featureExtractorType:configuration:error:
func NewAppleImageFeatureExtractorWithParametersModelDescriptionFeatureExtractorTypeConfigurationError(parameters objectivec.IObject, description objectivec.IObject, type_ int, configuration objectivec.IObject) (MLAppleImageFeatureExtractor, error) {
	var errorPtr objc.ID
	instance := getMLAppleImageFeatureExtractorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithParameters:modelDescription:featureExtractorType:configuration:error:"), parameters, description, type_, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleImageFeatureExtractor{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleImageFeatureExtractorFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLAppleImageFeatureExtractor/computeScenePrintFeatures:handle:useCPUOnly:error:
func (a MLAppleImageFeatureExtractor) ComputeScenePrintFeaturesHandleUseCPUOnlyError(features corevideo.CVImageBufferRef, handle objectivec.IObject, cPUOnly bool) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("computeScenePrintFeatures:handle:useCPUOnly:error:"), features, handle, cPUOnly, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLAppleImageFeatureExtractor/featureValueFromObjectPrint:key:shape:
func (a MLAppleImageFeatureExtractor) FeatureValueFromObjectPrintKeyShape(print_ objectivec.IObject, key objectivec.IObject, shape objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("featureValueFromObjectPrint:key:shape:"), print_, key, shape)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLAppleImageFeatureExtractor/featureValueFromScenePrint:elementSize:
func (a MLAppleImageFeatureExtractor) FeatureValueFromScenePrintElementSize(print_ objectivec.IObject, size uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("featureValueFromScenePrint:elementSize:"), print_, size)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLAppleImageFeatureExtractor/predictionFromFeatures:options:error:
func (a MLAppleImageFeatureExtractor) PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("predictionFromFeatures:options:error:"), features, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLAppleImageFeatureExtractor/initWithParameters:modelDescription:featureExtractorType:configuration:error:
func (a MLAppleImageFeatureExtractor) InitWithParametersModelDescriptionFeatureExtractorTypeConfigurationError(parameters objectivec.IObject, description objectivec.IObject, type_ int, configuration objectivec.IObject) (MLAppleImageFeatureExtractor, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("initWithParameters:modelDescription:featureExtractorType:configuration:error:"), parameters, description, type_, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleImageFeatureExtractor{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleImageFeatureExtractorFromID(rv), nil

}

//
// See: https://developer.apple.com/documentation/CoreML/MLAppleImageFeatureExtractor/loadModelFromSpecification:configuration:error:
func (_MLAppleImageFeatureExtractorClass MLAppleImageFeatureExtractorClass) LoadModelFromSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLAppleImageFeatureExtractorClass.class), objc.Sel("loadModelFromSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLAppleImageFeatureExtractor/parameters
func (a MLAppleImageFeatureExtractor) Parameters() IMLAppleImageFeatureExtractorParameters {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("parameters"))
	return MLAppleImageFeatureExtractorParametersFromID(objc.ID(rv))
}

