// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLAppleSoundAnalysisPreprocessing] class.
var (
	_MLAppleSoundAnalysisPreprocessingClass     MLAppleSoundAnalysisPreprocessingClass
	_MLAppleSoundAnalysisPreprocessingClassOnce sync.Once
)

func getMLAppleSoundAnalysisPreprocessingClass() MLAppleSoundAnalysisPreprocessingClass {
	_MLAppleSoundAnalysisPreprocessingClassOnce.Do(func() {
		_MLAppleSoundAnalysisPreprocessingClass = MLAppleSoundAnalysisPreprocessingClass{class: objc.GetClass("MLAppleSoundAnalysisPreprocessing")}
	})
	return _MLAppleSoundAnalysisPreprocessingClass
}

// GetMLAppleSoundAnalysisPreprocessingClass returns the class object for MLAppleSoundAnalysisPreprocessing.
func GetMLAppleSoundAnalysisPreprocessingClass() MLAppleSoundAnalysisPreprocessingClass {
	return getMLAppleSoundAnalysisPreprocessingClass()
}

type MLAppleSoundAnalysisPreprocessingClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLAppleSoundAnalysisPreprocessingClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLAppleSoundAnalysisPreprocessingClass) Alloc() MLAppleSoundAnalysisPreprocessing {
	rv := objc.Send[MLAppleSoundAnalysisPreprocessing](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLAppleSoundAnalysisPreprocessing.PredictionFromFeaturesOptionsError]
//   - [MLAppleSoundAnalysisPreprocessing.InitWithDescriptionConfigurationError]
// See: https://developer.apple.com/documentation/CoreML/MLAppleSoundAnalysisPreprocessing
type MLAppleSoundAnalysisPreprocessing struct {
	MLModel
}

// MLAppleSoundAnalysisPreprocessingFromID constructs a [MLAppleSoundAnalysisPreprocessing] from an objc.ID.
func MLAppleSoundAnalysisPreprocessingFromID(id objc.ID) MLAppleSoundAnalysisPreprocessing {
	return MLAppleSoundAnalysisPreprocessing{MLModel: MLModelFromID(id)}
}
// Ensure MLAppleSoundAnalysisPreprocessing implements IMLAppleSoundAnalysisPreprocessing.
var _ IMLAppleSoundAnalysisPreprocessing = MLAppleSoundAnalysisPreprocessing{}

// An interface definition for the [MLAppleSoundAnalysisPreprocessing] class.
//
// # Methods
//
//   - [IMLAppleSoundAnalysisPreprocessing.PredictionFromFeaturesOptionsError]
//   - [IMLAppleSoundAnalysisPreprocessing.InitWithDescriptionConfigurationError]
//
// See: https://developer.apple.com/documentation/CoreML/MLAppleSoundAnalysisPreprocessing
type IMLAppleSoundAnalysisPreprocessing interface {
	IMLModel

	// Topic: Methods

	PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	InitWithDescriptionConfigurationError(description objectivec.IObject, configuration objectivec.IObject) (MLAppleSoundAnalysisPreprocessing, error)
}

// Init initializes the instance.
func (a MLAppleSoundAnalysisPreprocessing) Init() MLAppleSoundAnalysisPreprocessing {
	rv := objc.Send[MLAppleSoundAnalysisPreprocessing](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a MLAppleSoundAnalysisPreprocessing) Autorelease() MLAppleSoundAnalysisPreprocessing {
	rv := objc.Send[MLAppleSoundAnalysisPreprocessing](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLAppleSoundAnalysisPreprocessing creates a new MLAppleSoundAnalysisPreprocessing instance.
func NewMLAppleSoundAnalysisPreprocessing() MLAppleSoundAnalysisPreprocessing {
	class := getMLAppleSoundAnalysisPreprocessingClass()
	rv := objc.Send[MLAppleSoundAnalysisPreprocessing](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initDescriptionOnlyWithSpecification:configuration:error:
func NewAppleSoundAnalysisPreprocessingDescriptionOnlyWithSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (MLAppleSoundAnalysisPreprocessing, error) {
	var errorPtr objc.ID
	instance := getMLAppleSoundAnalysisPreprocessingClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initDescriptionOnlyWithSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleSoundAnalysisPreprocessing{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleSoundAnalysisPreprocessingFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initInterfaceAndMetadataWithCompiledArchive:error:
func NewAppleSoundAnalysisPreprocessingInterfaceAndMetadataWithCompiledArchiveError(archive unsafe.Pointer) (MLAppleSoundAnalysisPreprocessing, error) {
	var errorPtr objc.ID
	instance := getMLAppleSoundAnalysisPreprocessingClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initInterfaceAndMetadataWithCompiledArchive:error:"), archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleSoundAnalysisPreprocessing{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleSoundAnalysisPreprocessingFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithConfiguration:
func NewAppleSoundAnalysisPreprocessingWithConfiguration(configuration objectivec.IObject) MLAppleSoundAnalysisPreprocessing {
	instance := getMLAppleSoundAnalysisPreprocessingClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	return MLAppleSoundAnalysisPreprocessingFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithDescription:
func NewAppleSoundAnalysisPreprocessingWithDescription(description objectivec.IObject) MLAppleSoundAnalysisPreprocessing {
	instance := getMLAppleSoundAnalysisPreprocessingClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:"), description)
	return MLAppleSoundAnalysisPreprocessingFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithDescription:configuration:
func NewAppleSoundAnalysisPreprocessingWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLAppleSoundAnalysisPreprocessing {
	instance := getMLAppleSoundAnalysisPreprocessingClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return MLAppleSoundAnalysisPreprocessingFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLAppleSoundAnalysisPreprocessing/initWithDescription:configuration:error:
func NewAppleSoundAnalysisPreprocessingWithDescriptionConfigurationError(description objectivec.IObject, configuration objectivec.IObject) (MLAppleSoundAnalysisPreprocessing, error) {
	var errorPtr objc.ID
	instance := getMLAppleSoundAnalysisPreprocessingClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:error:"), description, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleSoundAnalysisPreprocessing{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleSoundAnalysisPreprocessingFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:
func NewAppleSoundAnalysisPreprocessingWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLAppleSoundAnalysisPreprocessing {
	instance := getMLAppleSoundAnalysisPreprocessingClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return MLAppleSoundAnalysisPreprocessingFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLAppleSoundAnalysisPreprocessing/predictionFromFeatures:options:error:
func (a MLAppleSoundAnalysisPreprocessing) PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("predictionFromFeatures:options:error:"), features, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLAppleSoundAnalysisPreprocessing/initWithDescription:configuration:error:
func (a MLAppleSoundAnalysisPreprocessing) InitWithDescriptionConfigurationError(description objectivec.IObject, configuration objectivec.IObject) (MLAppleSoundAnalysisPreprocessing, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("initWithDescription:configuration:error:"), description, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLAppleSoundAnalysisPreprocessing{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLAppleSoundAnalysisPreprocessingFromID(rv), nil

}

//
// See: https://developer.apple.com/documentation/CoreML/MLAppleSoundAnalysisPreprocessing/loadModelFromSpecification:configuration:error:
func (_MLAppleSoundAnalysisPreprocessingClass MLAppleSoundAnalysisPreprocessingClass) LoadModelFromSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLAppleSoundAnalysisPreprocessingClass.class), objc.Sel("loadModelFromSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

