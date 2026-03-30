// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLArrayFeatureExtractor] class.
var (
	_MLArrayFeatureExtractorClass     MLArrayFeatureExtractorClass
	_MLArrayFeatureExtractorClassOnce sync.Once
)

func getMLArrayFeatureExtractorClass() MLArrayFeatureExtractorClass {
	_MLArrayFeatureExtractorClassOnce.Do(func() {
		_MLArrayFeatureExtractorClass = MLArrayFeatureExtractorClass{class: objc.GetClass("MLArrayFeatureExtractor")}
	})
	return _MLArrayFeatureExtractorClass
}

// GetMLArrayFeatureExtractorClass returns the class object for MLArrayFeatureExtractor.
func GetMLArrayFeatureExtractorClass() MLArrayFeatureExtractorClass {
	return getMLArrayFeatureExtractorClass()
}

type MLArrayFeatureExtractorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLArrayFeatureExtractorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLArrayFeatureExtractorClass) Alloc() MLArrayFeatureExtractor {
	rv := objc.Send[MLArrayFeatureExtractor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLArrayFeatureExtractor.ArrayColumnName]
//   - [MLArrayFeatureExtractor.ExtractIndices]
//   - [MLArrayFeatureExtractor.OutputType]
//   - [MLArrayFeatureExtractor.PredictionFromFeaturesOptionsError]
//   - [MLArrayFeatureExtractor.InitWithIndicesDataTransformerNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration]
//
// See: https://developer.apple.com/documentation/CoreML/MLArrayFeatureExtractor
type MLArrayFeatureExtractor struct {
	MLModel
}

// MLArrayFeatureExtractorFromID constructs a [MLArrayFeatureExtractor] from an objc.ID.
func MLArrayFeatureExtractorFromID(id objc.ID) MLArrayFeatureExtractor {
	return MLArrayFeatureExtractor{MLModel: MLModelFromID(id)}
}

// Ensure MLArrayFeatureExtractor implements IMLArrayFeatureExtractor.
var _ IMLArrayFeatureExtractor = MLArrayFeatureExtractor{}

// An interface definition for the [MLArrayFeatureExtractor] class.
//
// # Methods
//
//   - [IMLArrayFeatureExtractor.ArrayColumnName]
//   - [IMLArrayFeatureExtractor.ExtractIndices]
//   - [IMLArrayFeatureExtractor.OutputType]
//   - [IMLArrayFeatureExtractor.PredictionFromFeaturesOptionsError]
//   - [IMLArrayFeatureExtractor.InitWithIndicesDataTransformerNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration]
//
// See: https://developer.apple.com/documentation/CoreML/MLArrayFeatureExtractor
type IMLArrayFeatureExtractor interface {
	IMLModel

	// Topic: Methods

	ArrayColumnName() string
	ExtractIndices() foundation.INSArray
	OutputType() int64
	PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	InitWithIndicesDataTransformerNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(with objectivec.IObject, indices objectivec.IObject, name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLArrayFeatureExtractor
}

// Init initializes the instance.
func (a MLArrayFeatureExtractor) Init() MLArrayFeatureExtractor {
	rv := objc.Send[MLArrayFeatureExtractor](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a MLArrayFeatureExtractor) Autorelease() MLArrayFeatureExtractor {
	rv := objc.Send[MLArrayFeatureExtractor](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLArrayFeatureExtractor creates a new MLArrayFeatureExtractor instance.
func NewMLArrayFeatureExtractor() MLArrayFeatureExtractor {
	class := getMLArrayFeatureExtractorClass()
	rv := objc.Send[MLArrayFeatureExtractor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModel/initDescriptionOnlyWithSpecification:configuration:error:
func NewArrayFeatureExtractorDescriptionOnlyWithSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (MLArrayFeatureExtractor, error) {
	var errorPtr objc.ID
	instance := getMLArrayFeatureExtractorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initDescriptionOnlyWithSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLArrayFeatureExtractor{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLArrayFeatureExtractorFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLModel/initInterfaceAndMetadataWithCompiledArchive:error:
func NewArrayFeatureExtractorInterfaceAndMetadataWithCompiledArchiveError(archive unsafe.Pointer) (MLArrayFeatureExtractor, error) {
	var errorPtr objc.ID
	instance := getMLArrayFeatureExtractorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initInterfaceAndMetadataWithCompiledArchive:error:"), archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLArrayFeatureExtractor{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLArrayFeatureExtractorFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithConfiguration:
func NewArrayFeatureExtractorWithConfiguration(configuration objectivec.IObject) MLArrayFeatureExtractor {
	instance := getMLArrayFeatureExtractorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	return MLArrayFeatureExtractorFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithDescription:
func NewArrayFeatureExtractorWithDescription(description objectivec.IObject) MLArrayFeatureExtractor {
	instance := getMLArrayFeatureExtractorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:"), description)
	return MLArrayFeatureExtractorFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithDescription:configuration:
func NewArrayFeatureExtractorWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLArrayFeatureExtractor {
	instance := getMLArrayFeatureExtractorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return MLArrayFeatureExtractorFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLArrayFeatureExtractor/initWith:indices:dataTransformerName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:
func NewArrayFeatureExtractorWithIndicesDataTransformerNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(with objectivec.IObject, indices objectivec.IObject, name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLArrayFeatureExtractor {
	instance := getMLArrayFeatureExtractorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWith:indices:dataTransformerName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), with, indices, name, description, description2, names, names2, configuration)
	return MLArrayFeatureExtractorFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:
func NewArrayFeatureExtractorWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLArrayFeatureExtractor {
	instance := getMLArrayFeatureExtractorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return MLArrayFeatureExtractorFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLArrayFeatureExtractor/predictionFromFeatures:options:error:
func (a MLArrayFeatureExtractor) PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("predictionFromFeatures:options:error:"), features, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLArrayFeatureExtractor/initWith:indices:dataTransformerName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:
func (a MLArrayFeatureExtractor) InitWithIndicesDataTransformerNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(with objectivec.IObject, indices objectivec.IObject, name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLArrayFeatureExtractor {
	rv := objc.Send[MLArrayFeatureExtractor](a.ID, objc.Sel("initWith:indices:dataTransformerName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), with, indices, name, description, description2, names, names2, configuration)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLArrayFeatureExtractor/extractArrayElement:indices:dataTransformerName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:
func (_MLArrayFeatureExtractorClass MLArrayFeatureExtractorClass) ExtractArrayElementIndicesDataTransformerNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNames(element objectivec.IObject, indices objectivec.IObject, name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLArrayFeatureExtractorClass.class), objc.Sel("extractArrayElement:indices:dataTransformerName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:"), element, indices, name, description, description2, names, names2)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLArrayFeatureExtractor/extractArrayElement:indices:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:
func (_MLArrayFeatureExtractorClass MLArrayFeatureExtractorClass) ExtractArrayElementIndicesInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNames(element objectivec.IObject, indices objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLArrayFeatureExtractorClass.class), objc.Sel("extractArrayElement:indices:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:"), element, indices, description, description2, names, names2)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLArrayFeatureExtractor/loadModelFromSpecification:configuration:error:
func (_MLArrayFeatureExtractorClass MLArrayFeatureExtractorClass) LoadModelFromSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLArrayFeatureExtractorClass.class), objc.Sel("loadModelFromSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLArrayFeatureExtractor/arrayColumnName
func (a MLArrayFeatureExtractor) ArrayColumnName() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("arrayColumnName"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLArrayFeatureExtractor/extractIndices
func (a MLArrayFeatureExtractor) ExtractIndices() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("extractIndices"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLArrayFeatureExtractor/outputType
func (a MLArrayFeatureExtractor) OutputType() int64 {
	rv := objc.Send[int64](a.ID, objc.Sel("outputType"))
	return rv
}
