// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/corevideo"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
	rv := objc.SendIfResponds[MLAppleImageFeatureExtractor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLAppleImageFeatureExtractor.ComputeScenePrintFeaturesHandleUseCPUOnlyError]
//   - [MLAppleImageFeatureExtractor.FeatureValueFromObjectPrintKeyShape]
//   - [MLAppleImageFeatureExtractor.FeatureValueFromScenePrintElementSize]
//   - [MLAppleImageFeatureExtractor.Parameters]
//   - [MLAppleImageFeatureExtractor.PredictionFromFeaturesOptionsError]
//   - [MLAppleImageFeatureExtractor.InitWithParametersModelDescriptionFeatureExtractorTypeConfigurationError]
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
func (m MLAppleImageFeatureExtractor) Init() MLAppleImageFeatureExtractor {
	rv := objc.SendIfResponds[MLAppleImageFeatureExtractor](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLAppleImageFeatureExtractor) Autorelease() MLAppleImageFeatureExtractor {
	rv := objc.SendIfResponds[MLAppleImageFeatureExtractor](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLAppleImageFeatureExtractor creates a new MLAppleImageFeatureExtractor instance.
func NewMLAppleImageFeatureExtractor() MLAppleImageFeatureExtractor {
	class := getMLAppleImageFeatureExtractorClass()
	rv := objc.SendIfResponds[MLAppleImageFeatureExtractor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewAppleImageFeatureExtractorDescriptionOnlyWithSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (MLAppleImageFeatureExtractor, error) {
	var errorPtr objc.ID
	instance := getMLAppleImageFeatureExtractorClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initDescriptionOnlyWithSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleImageFeatureExtractor{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return MLAppleImageFeatureExtractor{}, objc.ErrInitFailed
	}
	return MLAppleImageFeatureExtractorFromID(rv), nil
}

func NewAppleImageFeatureExtractorInterfaceAndMetadataWithCompiledArchiveError(archive unsafe.Pointer) (MLAppleImageFeatureExtractor, error) {
	var errorPtr objc.ID
	instance := getMLAppleImageFeatureExtractorClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initInterfaceAndMetadataWithCompiledArchive:error:"), archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleImageFeatureExtractor{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return MLAppleImageFeatureExtractor{}, objc.ErrInitFailed
	}
	return MLAppleImageFeatureExtractorFromID(rv), nil
}

func NewAppleImageFeatureExtractorWithConfiguration(configuration objectivec.IObject) MLAppleImageFeatureExtractor {
	instance := getMLAppleImageFeatureExtractorClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	return MLAppleImageFeatureExtractorFromID(rv)
}

func NewAppleImageFeatureExtractorWithDescription(description objectivec.IObject) MLAppleImageFeatureExtractor {
	instance := getMLAppleImageFeatureExtractorClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDescription:"), description)
	return MLAppleImageFeatureExtractorFromID(rv)
}

func NewAppleImageFeatureExtractorWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLAppleImageFeatureExtractor {
	instance := getMLAppleImageFeatureExtractorClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return MLAppleImageFeatureExtractorFromID(rv)
}

func NewAppleImageFeatureExtractorWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLAppleImageFeatureExtractor {
	instance := getMLAppleImageFeatureExtractorClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return MLAppleImageFeatureExtractorFromID(rv)
}

func NewAppleImageFeatureExtractorWithParametersModelDescriptionFeatureExtractorTypeConfigurationError(parameters objectivec.IObject, description objectivec.IObject, type_ int, configuration objectivec.IObject) (MLAppleImageFeatureExtractor, error) {
	var errorPtr objc.ID
	instance := getMLAppleImageFeatureExtractorClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithParameters:modelDescription:featureExtractorType:configuration:error:"), parameters, description, type_, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleImageFeatureExtractor{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return MLAppleImageFeatureExtractor{}, objc.ErrInitFailed
	}
	return MLAppleImageFeatureExtractorFromID(rv), nil
}

func (m MLAppleImageFeatureExtractor) ComputeScenePrintFeaturesHandleUseCPUOnlyError(features corevideo.CVImageBufferRef, handle objectivec.IObject, cPUOnly bool) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("computeScenePrintFeatures:handle:useCPUOnly:error:"), features, handle, cPUOnly, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLAppleImageFeatureExtractor) FeatureValueFromObjectPrintKeyShape(print_ objectivec.IObject, key objectivec.IObject, shape objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("featureValueFromObjectPrint:key:shape:"), print_, key, shape)
	return objectivec.Object{ID: rv}
}
func (m MLAppleImageFeatureExtractor) FeatureValueFromScenePrintElementSize(print_ objectivec.IObject, size uint64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("featureValueFromScenePrint:elementSize:"), print_, size)
	return objectivec.Object{ID: rv}
}
func (m MLAppleImageFeatureExtractor) PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictionFromFeatures:options:error:"), features, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLAppleImageFeatureExtractor) InitWithParametersModelDescriptionFeatureExtractorTypeConfigurationError(parameters objectivec.IObject, description objectivec.IObject, type_ int, configuration objectivec.IObject) (MLAppleImageFeatureExtractor, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithParameters:modelDescription:featureExtractorType:configuration:error:"), parameters, description, type_, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleImageFeatureExtractor{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleImageFeatureExtractorFromID(rv), nil

}

func (_MLAppleImageFeatureExtractorClass MLAppleImageFeatureExtractorClass) LoadModelFromSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLAppleImageFeatureExtractorClass.class), objc.Sel("loadModelFromSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

func (m MLAppleImageFeatureExtractor) Parameters() IMLAppleImageFeatureExtractorParameters {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("parameters"))
	return MLAppleImageFeatureExtractorParametersFromID(objc.ID(rv))
}
