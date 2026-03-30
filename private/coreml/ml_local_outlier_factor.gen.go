// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLLocalOutlierFactor] class.
var (
	_MLLocalOutlierFactorClass     MLLocalOutlierFactorClass
	_MLLocalOutlierFactorClassOnce sync.Once
)

func getMLLocalOutlierFactorClass() MLLocalOutlierFactorClass {
	_MLLocalOutlierFactorClassOnce.Do(func() {
		_MLLocalOutlierFactorClass = MLLocalOutlierFactorClass{class: objc.GetClass("MLLocalOutlierFactor")}
	})
	return _MLLocalOutlierFactorClass
}

// GetMLLocalOutlierFactorClass returns the class object for MLLocalOutlierFactor.
func GetMLLocalOutlierFactorClass() MLLocalOutlierFactorClass {
	return getMLLocalOutlierFactorClass()
}

type MLLocalOutlierFactorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLLocalOutlierFactorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLLocalOutlierFactorClass) Alloc() MLLocalOutlierFactor {
	rv := objc.Send[MLLocalOutlierFactor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLLocalOutlierFactor.ComputeLOFForQueryPoint]
//   - [MLLocalOutlierFactor.FindNearestNeighborsToIndex]
//   - [MLLocalOutlierFactor.FindNearestNeighborsToQueryPoint]
//   - [MLLocalOutlierFactor.InputMultiArrayError]
//   - [MLLocalOutlierFactor.KDistanceToIndex]
//   - [MLLocalOutlierFactor.LocalReachabilityDensityForIndex]
//   - [MLLocalOutlierFactor.LocalReachabilityDensityForQeuryPoint]
//   - [MLLocalOutlierFactor.LocalReachabilityDensityOfNeighbors]
//   - [MLLocalOutlierFactor.ParameterValueForKeyError]
//   - [MLLocalOutlierFactor.PredictionFromFeaturesOptionsError]
//   - [MLLocalOutlierFactor.UpdateToValidDistance]
//   - [MLLocalOutlierFactor.InitWithKNearestNeighborsModelAtURLConfigurationError]
//
// See: https://developer.apple.com/documentation/CoreML/MLLocalOutlierFactor
type MLLocalOutlierFactor struct {
	MLModel
}

// MLLocalOutlierFactorFromID constructs a [MLLocalOutlierFactor] from an objc.ID.
func MLLocalOutlierFactorFromID(id objc.ID) MLLocalOutlierFactor {
	return MLLocalOutlierFactor{MLModel: MLModelFromID(id)}
}

// Ensure MLLocalOutlierFactor implements IMLLocalOutlierFactor.
var _ IMLLocalOutlierFactor = MLLocalOutlierFactor{}

// An interface definition for the [MLLocalOutlierFactor] class.
//
// # Methods
//
//   - [IMLLocalOutlierFactor.ComputeLOFForQueryPoint]
//   - [IMLLocalOutlierFactor.FindNearestNeighborsToIndex]
//   - [IMLLocalOutlierFactor.FindNearestNeighborsToQueryPoint]
//   - [IMLLocalOutlierFactor.InputMultiArrayError]
//   - [IMLLocalOutlierFactor.KDistanceToIndex]
//   - [IMLLocalOutlierFactor.LocalReachabilityDensityForIndex]
//   - [IMLLocalOutlierFactor.LocalReachabilityDensityForQeuryPoint]
//   - [IMLLocalOutlierFactor.LocalReachabilityDensityOfNeighbors]
//   - [IMLLocalOutlierFactor.ParameterValueForKeyError]
//   - [IMLLocalOutlierFactor.PredictionFromFeaturesOptionsError]
//   - [IMLLocalOutlierFactor.UpdateToValidDistance]
//   - [IMLLocalOutlierFactor.InitWithKNearestNeighborsModelAtURLConfigurationError]
//
// See: https://developer.apple.com/documentation/CoreML/MLLocalOutlierFactor
type IMLLocalOutlierFactor interface {
	IMLModel

	// Topic: Methods

	ComputeLOFForQueryPoint(point objectivec.IObject) float64
	FindNearestNeighborsToIndex(index uint64) objectivec.IObject
	FindNearestNeighborsToQueryPoint(point unsafe.Pointer) objectivec.IObject
	InputMultiArrayError(array objectivec.IObject) (objectivec.IObject, error)
	KDistanceToIndex(index uint64) float32
	LocalReachabilityDensityForIndex(index uint64) float64
	LocalReachabilityDensityForQeuryPoint(point unsafe.Pointer) float64
	LocalReachabilityDensityOfNeighbors(neighbors unsafe.Pointer) float64
	ParameterValueForKeyError(key objectivec.IObject) (objectivec.IObject, error)
	PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	UpdateToValidDistance(distance unsafe.Pointer)
	InitWithKNearestNeighborsModelAtURLConfigurationError(url foundation.INSURL, configuration objectivec.IObject) (MLLocalOutlierFactor, error)
}

// Init initializes the instance.
func (l MLLocalOutlierFactor) Init() MLLocalOutlierFactor {
	rv := objc.Send[MLLocalOutlierFactor](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l MLLocalOutlierFactor) Autorelease() MLLocalOutlierFactor {
	rv := objc.Send[MLLocalOutlierFactor](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLLocalOutlierFactor creates a new MLLocalOutlierFactor instance.
func NewMLLocalOutlierFactor() MLLocalOutlierFactor {
	class := getMLLocalOutlierFactorClass()
	rv := objc.Send[MLLocalOutlierFactor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModel/initDescriptionOnlyWithSpecification:configuration:error:
func NewLocalOutlierFactorDescriptionOnlyWithSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (MLLocalOutlierFactor, error) {
	var errorPtr objc.ID
	instance := getMLLocalOutlierFactorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initDescriptionOnlyWithSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLLocalOutlierFactor{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLLocalOutlierFactorFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLModel/initInterfaceAndMetadataWithCompiledArchive:error:
func NewLocalOutlierFactorInterfaceAndMetadataWithCompiledArchiveError(archive unsafe.Pointer) (MLLocalOutlierFactor, error) {
	var errorPtr objc.ID
	instance := getMLLocalOutlierFactorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initInterfaceAndMetadataWithCompiledArchive:error:"), archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLLocalOutlierFactor{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLLocalOutlierFactorFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithConfiguration:
func NewLocalOutlierFactorWithConfiguration(configuration objectivec.IObject) MLLocalOutlierFactor {
	instance := getMLLocalOutlierFactorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	return MLLocalOutlierFactorFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithDescription:
func NewLocalOutlierFactorWithDescription(description objectivec.IObject) MLLocalOutlierFactor {
	instance := getMLLocalOutlierFactorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:"), description)
	return MLLocalOutlierFactorFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithDescription:configuration:
func NewLocalOutlierFactorWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLLocalOutlierFactor {
	instance := getMLLocalOutlierFactorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return MLLocalOutlierFactorFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLLocalOutlierFactor/initWithKNearestNeighborsModelAtURL:configuration:error:
func NewLocalOutlierFactorWithKNearestNeighborsModelAtURLConfigurationError(url foundation.INSURL, configuration objectivec.IObject) (MLLocalOutlierFactor, error) {
	var errorPtr objc.ID
	instance := getMLLocalOutlierFactorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithKNearestNeighborsModelAtURL:configuration:error:"), url, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLLocalOutlierFactor{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLLocalOutlierFactorFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:
func NewLocalOutlierFactorWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLLocalOutlierFactor {
	instance := getMLLocalOutlierFactorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return MLLocalOutlierFactorFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLLocalOutlierFactor/computeLOFForQueryPoint:
func (l MLLocalOutlierFactor) ComputeLOFForQueryPoint(point objectivec.IObject) float64 {
	rv := objc.Send[float64](l.ID, objc.Sel("computeLOFForQueryPoint:"), point)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLLocalOutlierFactor/findNearestNeighborsToIndex:
func (l MLLocalOutlierFactor) FindNearestNeighborsToIndex(index uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("findNearestNeighborsToIndex:"), index)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLLocalOutlierFactor/findNearestNeighborsToQueryPoint:
func (l MLLocalOutlierFactor) FindNearestNeighborsToQueryPoint(point unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("findNearestNeighborsToQueryPoint:"), point)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLLocalOutlierFactor/inputMultiArray:error:
func (l MLLocalOutlierFactor) InputMultiArrayError(array objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](l.ID, objc.Sel("inputMultiArray:error:"), array, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLLocalOutlierFactor/kDistanceToIndex:
func (l MLLocalOutlierFactor) KDistanceToIndex(index uint64) float32 {
	rv := objc.Send[float32](l.ID, objc.Sel("kDistanceToIndex:"), index)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLLocalOutlierFactor/localReachabilityDensityForIndex:
func (l MLLocalOutlierFactor) LocalReachabilityDensityForIndex(index uint64) float64 {
	rv := objc.Send[float64](l.ID, objc.Sel("localReachabilityDensityForIndex:"), index)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLLocalOutlierFactor/localReachabilityDensityForQeuryPoint:
func (l MLLocalOutlierFactor) LocalReachabilityDensityForQeuryPoint(point unsafe.Pointer) float64 {
	rv := objc.Send[float64](l.ID, objc.Sel("localReachabilityDensityForQeuryPoint:"), point)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLLocalOutlierFactor/localReachabilityDensityOfNeighbors:
func (l MLLocalOutlierFactor) LocalReachabilityDensityOfNeighbors(neighbors unsafe.Pointer) float64 {
	rv := objc.Send[float64](l.ID, objc.Sel("localReachabilityDensityOfNeighbors:"), neighbors)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLLocalOutlierFactor/parameterValueForKey:error:
func (l MLLocalOutlierFactor) ParameterValueForKeyError(key objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](l.ID, objc.Sel("parameterValueForKey:error:"), key, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLLocalOutlierFactor/predictionFromFeatures:options:error:
func (l MLLocalOutlierFactor) PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](l.ID, objc.Sel("predictionFromFeatures:options:error:"), features, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLLocalOutlierFactor/updateToValidDistance:
func (l MLLocalOutlierFactor) UpdateToValidDistance(distance unsafe.Pointer) {
	objc.Send[objc.ID](l.ID, objc.Sel("updateToValidDistance:"), distance)
}

// See: https://developer.apple.com/documentation/CoreML/MLLocalOutlierFactor/initWithKNearestNeighborsModelAtURL:configuration:error:
func (l MLLocalOutlierFactor) InitWithKNearestNeighborsModelAtURLConfigurationError(url foundation.INSURL, configuration objectivec.IObject) (MLLocalOutlierFactor, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](l.ID, objc.Sel("initWithKNearestNeighborsModelAtURL:configuration:error:"), url, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLLocalOutlierFactor{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLLocalOutlierFactorFromID(rv), nil

}
