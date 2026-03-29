// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLClassConfidenceThresholding] class.
var (
	_MLClassConfidenceThresholdingClass     MLClassConfidenceThresholdingClass
	_MLClassConfidenceThresholdingClassOnce sync.Once
)

func getMLClassConfidenceThresholdingClass() MLClassConfidenceThresholdingClass {
	_MLClassConfidenceThresholdingClassOnce.Do(func() {
		_MLClassConfidenceThresholdingClass = MLClassConfidenceThresholdingClass{class: objc.GetClass("MLClassConfidenceThresholding")}
	})
	return _MLClassConfidenceThresholdingClass
}

// GetMLClassConfidenceThresholdingClass returns the class object for MLClassConfidenceThresholding.
func GetMLClassConfidenceThresholdingClass() MLClassConfidenceThresholdingClass {
	return getMLClassConfidenceThresholdingClass()
}

type MLClassConfidenceThresholdingClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLClassConfidenceThresholdingClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLClassConfidenceThresholdingClass) Alloc() MLClassConfidenceThresholding {
	rv := objc.Send[MLClassConfidenceThresholding](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLClassConfidenceThresholding.InputFeatureConformer]
//   - [MLClassConfidenceThresholding.ParameterContainer]
//   - [MLClassConfidenceThresholding.SetParameterContainer]
//   - [MLClassConfidenceThresholding.ParameterValueForKeyError]
//   - [MLClassConfidenceThresholding.PredictionFromFeaturesOptionsError]
//   - [MLClassConfidenceThresholding.InitWithDescriptionConfigurationPrecisionRecallCurvesError]
// See: https://developer.apple.com/documentation/CoreML/MLClassConfidenceThresholding
type MLClassConfidenceThresholding struct {
	MLModel
}

// MLClassConfidenceThresholdingFromID constructs a [MLClassConfidenceThresholding] from an objc.ID.
func MLClassConfidenceThresholdingFromID(id objc.ID) MLClassConfidenceThresholding {
	return MLClassConfidenceThresholding{MLModel: MLModelFromID(id)}
}
// Ensure MLClassConfidenceThresholding implements IMLClassConfidenceThresholding.
var _ IMLClassConfidenceThresholding = MLClassConfidenceThresholding{}

// An interface definition for the [MLClassConfidenceThresholding] class.
//
// # Methods
//
//   - [IMLClassConfidenceThresholding.InputFeatureConformer]
//   - [IMLClassConfidenceThresholding.ParameterContainer]
//   - [IMLClassConfidenceThresholding.SetParameterContainer]
//   - [IMLClassConfidenceThresholding.ParameterValueForKeyError]
//   - [IMLClassConfidenceThresholding.PredictionFromFeaturesOptionsError]
//   - [IMLClassConfidenceThresholding.InitWithDescriptionConfigurationPrecisionRecallCurvesError]
//
// See: https://developer.apple.com/documentation/CoreML/MLClassConfidenceThresholding
type IMLClassConfidenceThresholding interface {
	IMLModel

	// Topic: Methods

	InputFeatureConformer() IMLFeatureProviderConformer
	ParameterContainer() IMLParameterContainer
	SetParameterContainer(value IMLParameterContainer)
	ParameterValueForKeyError(key objectivec.IObject) (objectivec.IObject, error)
	PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	InitWithDescriptionConfigurationPrecisionRecallCurvesError(description objectivec.IObject, configuration objectivec.IObject, curves objectivec.IObject) (MLClassConfidenceThresholding, error)
}

// Init initializes the instance.
func (c MLClassConfidenceThresholding) Init() MLClassConfidenceThresholding {
	rv := objc.Send[MLClassConfidenceThresholding](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MLClassConfidenceThresholding) Autorelease() MLClassConfidenceThresholding {
	rv := objc.Send[MLClassConfidenceThresholding](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLClassConfidenceThresholding creates a new MLClassConfidenceThresholding instance.
func NewMLClassConfidenceThresholding() MLClassConfidenceThresholding {
	class := getMLClassConfidenceThresholdingClass()
	rv := objc.Send[MLClassConfidenceThresholding](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initDescriptionOnlyWithSpecification:configuration:error:
func NewClassConfidenceThresholdingDescriptionOnlyWithSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (MLClassConfidenceThresholding, error) {
	var errorPtr objc.ID
	instance := getMLClassConfidenceThresholdingClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initDescriptionOnlyWithSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLClassConfidenceThresholding{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLClassConfidenceThresholdingFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initInterfaceAndMetadataWithCompiledArchive:error:
func NewClassConfidenceThresholdingInterfaceAndMetadataWithCompiledArchiveError(archive unsafe.Pointer) (MLClassConfidenceThresholding, error) {
	var errorPtr objc.ID
	instance := getMLClassConfidenceThresholdingClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initInterfaceAndMetadataWithCompiledArchive:error:"), archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLClassConfidenceThresholding{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLClassConfidenceThresholdingFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithConfiguration:
func NewClassConfidenceThresholdingWithConfiguration(configuration objectivec.IObject) MLClassConfidenceThresholding {
	instance := getMLClassConfidenceThresholdingClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	return MLClassConfidenceThresholdingFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithDescription:
func NewClassConfidenceThresholdingWithDescription(description objectivec.IObject) MLClassConfidenceThresholding {
	instance := getMLClassConfidenceThresholdingClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:"), description)
	return MLClassConfidenceThresholdingFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithDescription:configuration:
func NewClassConfidenceThresholdingWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLClassConfidenceThresholding {
	instance := getMLClassConfidenceThresholdingClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return MLClassConfidenceThresholdingFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLClassConfidenceThresholding/initWithDescription:configuration:precisionRecallCurves:error:
func NewClassConfidenceThresholdingWithDescriptionConfigurationPrecisionRecallCurvesError(description objectivec.IObject, configuration objectivec.IObject, curves objectivec.IObject) (MLClassConfidenceThresholding, error) {
	var errorPtr objc.ID
	instance := getMLClassConfidenceThresholdingClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:precisionRecallCurves:error:"), description, configuration, curves, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLClassConfidenceThresholding{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLClassConfidenceThresholdingFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:
func NewClassConfidenceThresholdingWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLClassConfidenceThresholding {
	instance := getMLClassConfidenceThresholdingClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return MLClassConfidenceThresholdingFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLClassConfidenceThresholding/parameterValueForKey:error:
func (c MLClassConfidenceThresholding) ParameterValueForKeyError(key objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](c.ID, objc.Sel("parameterValueForKey:error:"), key, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLClassConfidenceThresholding/predictionFromFeatures:options:error:
func (c MLClassConfidenceThresholding) PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](c.ID, objc.Sel("predictionFromFeatures:options:error:"), features, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLClassConfidenceThresholding/initWithDescription:configuration:precisionRecallCurves:error:
func (c MLClassConfidenceThresholding) InitWithDescriptionConfigurationPrecisionRecallCurvesError(description objectivec.IObject, configuration objectivec.IObject, curves objectivec.IObject) (MLClassConfidenceThresholding, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](c.ID, objc.Sel("initWithDescription:configuration:precisionRecallCurves:error:"), description, configuration, curves, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLClassConfidenceThresholding{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLClassConfidenceThresholdingFromID(rv), nil

}

//
// See: https://developer.apple.com/documentation/CoreML/MLClassConfidenceThresholding/loadModelFromSpecification:configuration:error:
func (_MLClassConfidenceThresholdingClass MLClassConfidenceThresholdingClass) LoadModelFromSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLClassConfidenceThresholdingClass.class), objc.Sel("loadModelFromSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLClassConfidenceThresholding/inputFeatureConformer
func (c MLClassConfidenceThresholding) InputFeatureConformer() IMLFeatureProviderConformer {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("inputFeatureConformer"))
	return MLFeatureProviderConformerFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLClassConfidenceThresholding/parameterContainer
func (c MLClassConfidenceThresholding) ParameterContainer() IMLParameterContainer {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("parameterContainer"))
	return MLParameterContainerFromID(objc.ID(rv))
}
func (c MLClassConfidenceThresholding) SetParameterContainer(value IMLParameterContainer) {
	objc.Send[struct{}](c.ID, objc.Sel("setParameterContainer:"), value)
}

