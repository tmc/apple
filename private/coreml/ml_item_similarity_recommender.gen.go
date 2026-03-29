// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLItemSimilarityRecommender] class.
var (
	_MLItemSimilarityRecommenderClass     MLItemSimilarityRecommenderClass
	_MLItemSimilarityRecommenderClassOnce sync.Once
)

func getMLItemSimilarityRecommenderClass() MLItemSimilarityRecommenderClass {
	_MLItemSimilarityRecommenderClassOnce.Do(func() {
		_MLItemSimilarityRecommenderClass = MLItemSimilarityRecommenderClass{class: objc.GetClass("MLItemSimilarityRecommender")}
	})
	return _MLItemSimilarityRecommenderClass
}

// GetMLItemSimilarityRecommenderClass returns the class object for MLItemSimilarityRecommender.
func GetMLItemSimilarityRecommenderClass() MLItemSimilarityRecommenderClass {
	return getMLItemSimilarityRecommenderClass()
}

type MLItemSimilarityRecommenderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLItemSimilarityRecommenderClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLItemSimilarityRecommenderClass) Alloc() MLItemSimilarityRecommender {
	rv := objc.Send[MLItemSimilarityRecommender](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLItemSimilarityRecommender._itemForIndexError]
//   - [MLItemSimilarityRecommender._mapItemSequenceDestError]
//   - [MLItemSimilarityRecommender.ModelData]
//   - [MLItemSimilarityRecommender.PredictionFromFeaturesOptionsError]
//   - [MLItemSimilarityRecommender.Metadata]
// See: https://developer.apple.com/documentation/CoreML/MLItemSimilarityRecommender
type MLItemSimilarityRecommender struct {
	MLModel
}

// MLItemSimilarityRecommenderFromID constructs a [MLItemSimilarityRecommender] from an objc.ID.
func MLItemSimilarityRecommenderFromID(id objc.ID) MLItemSimilarityRecommender {
	return MLItemSimilarityRecommender{MLModel: MLModelFromID(id)}
}
// Ensure MLItemSimilarityRecommender implements IMLItemSimilarityRecommender.
var _ IMLItemSimilarityRecommender = MLItemSimilarityRecommender{}

// An interface definition for the [MLItemSimilarityRecommender] class.
//
// # Methods
//
//   - [IMLItemSimilarityRecommender._itemForIndexError]
//   - [IMLItemSimilarityRecommender._mapItemSequenceDestError]
//   - [IMLItemSimilarityRecommender.ModelData]
//   - [IMLItemSimilarityRecommender.PredictionFromFeaturesOptionsError]
//   - [IMLItemSimilarityRecommender.Metadata]
//
// See: https://developer.apple.com/documentation/CoreML/MLItemSimilarityRecommender
type IMLItemSimilarityRecommender interface {
	IMLModel

	// Topic: Methods

	_itemForIndexError(index uint64) (objectivec.IObject, error)
	_mapItemSequenceDestError(sequence objectivec.IObject, dest unsafe.Pointer) (bool, error)
	ModelData() string
	PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	Metadata() IMLModelMetadata
}

// Init initializes the instance.
func (i MLItemSimilarityRecommender) Init() MLItemSimilarityRecommender {
	rv := objc.Send[MLItemSimilarityRecommender](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MLItemSimilarityRecommender) Autorelease() MLItemSimilarityRecommender {
	rv := objc.Send[MLItemSimilarityRecommender](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLItemSimilarityRecommender creates a new MLItemSimilarityRecommender instance.
func NewMLItemSimilarityRecommender() MLItemSimilarityRecommender {
	class := getMLItemSimilarityRecommenderClass()
	rv := objc.Send[MLItemSimilarityRecommender](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initDescriptionOnlyWithSpecification:configuration:error:
func NewItemSimilarityRecommenderDescriptionOnlyWithSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (MLItemSimilarityRecommender, error) {
	var errorPtr objc.ID
	instance := getMLItemSimilarityRecommenderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initDescriptionOnlyWithSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLItemSimilarityRecommender{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLItemSimilarityRecommenderFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initInterfaceAndMetadataWithCompiledArchive:error:
func NewItemSimilarityRecommenderInterfaceAndMetadataWithCompiledArchiveError(archive unsafe.Pointer) (MLItemSimilarityRecommender, error) {
	var errorPtr objc.ID
	instance := getMLItemSimilarityRecommenderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initInterfaceAndMetadataWithCompiledArchive:error:"), archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLItemSimilarityRecommender{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLItemSimilarityRecommenderFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithConfiguration:
func NewItemSimilarityRecommenderWithConfiguration(configuration objectivec.IObject) MLItemSimilarityRecommender {
	instance := getMLItemSimilarityRecommenderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	return MLItemSimilarityRecommenderFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithDescription:
func NewItemSimilarityRecommenderWithDescription(description objectivec.IObject) MLItemSimilarityRecommender {
	instance := getMLItemSimilarityRecommenderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:"), description)
	return MLItemSimilarityRecommenderFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithDescription:configuration:
func NewItemSimilarityRecommenderWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLItemSimilarityRecommender {
	instance := getMLItemSimilarityRecommenderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return MLItemSimilarityRecommenderFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:
func NewItemSimilarityRecommenderWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLItemSimilarityRecommender {
	instance := getMLItemSimilarityRecommenderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return MLItemSimilarityRecommenderFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLItemSimilarityRecommender/_itemForIndex:error:
func (i MLItemSimilarityRecommender) _itemForIndexError(index uint64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](i.ID, objc.Sel("_itemForIndex:error:"), index, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// ItemForIndexError is an exported wrapper for the private method _itemForIndexError.
func (i MLItemSimilarityRecommender) ItemForIndexError(index uint64) (objectivec.IObject, error) {
	return i._itemForIndexError(index)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLItemSimilarityRecommender/_mapItemSequence:dest:error:
func (i MLItemSimilarityRecommender) _mapItemSequenceDestError(sequence objectivec.IObject, dest unsafe.Pointer) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](i.ID, objc.Sel("_mapItemSequence:dest:error:"), sequence, dest, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("_mapItemSequence:dest:error: returned NO with nil NSError")
	}
	return rv, nil

}

// MapItemSequenceDestError is an exported wrapper for the private method _mapItemSequenceDestError.
func (i MLItemSimilarityRecommender) MapItemSequenceDestError(sequence objectivec.IObject, dest unsafe.Pointer) (bool, error) {
	return i._mapItemSequenceDestError(sequence, dest)
}
// See: https://developer.apple.com/documentation/CoreML/MLItemSimilarityRecommender/modelData
func (i MLItemSimilarityRecommender) ModelData() string {
	rv := objc.Send[*byte](i.ID, objc.Sel("modelData"))
	return objc.GoString(rv)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLItemSimilarityRecommender/predictionFromFeatures:options:error:
func (i MLItemSimilarityRecommender) PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](i.ID, objc.Sel("predictionFromFeatures:options:error:"), features, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

//
// See: https://developer.apple.com/documentation/CoreML/MLItemSimilarityRecommender/compileSpecification:toArchive:options:error:
func (_MLItemSimilarityRecommenderClass MLItemSimilarityRecommenderClass) CompileSpecificationToArchiveOptionsError(specification unsafe.Pointer, archive unsafe.Pointer, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLItemSimilarityRecommenderClass.class), objc.Sel("compileSpecification:toArchive:options:error:"), specification, archive, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLItemSimilarityRecommender/compiledVersionForSpecification:options:error:
func (_MLItemSimilarityRecommenderClass MLItemSimilarityRecommenderClass) CompiledVersionForSpecificationOptionsError(specification unsafe.Pointer, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLItemSimilarityRecommenderClass.class), objc.Sel("compiledVersionForSpecification:options:error:"), specification, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLItemSimilarityRecommender/loadModelFromCompiledArchive:modelVersionInfo:compilerVersionInfo:configuration:error:
func (_MLItemSimilarityRecommenderClass MLItemSimilarityRecommenderClass) LoadModelFromCompiledArchiveModelVersionInfoCompilerVersionInfoConfigurationError(archive unsafe.Pointer, info objectivec.IObject, info2 objectivec.IObject, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLItemSimilarityRecommenderClass.class), objc.Sel("loadModelFromCompiledArchive:modelVersionInfo:compilerVersionInfo:configuration:error:"), archive, info, info2, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLItemSimilarityRecommender/loadModelFromSpecification:configuration:error:
func (_MLItemSimilarityRecommenderClass MLItemSimilarityRecommenderClass) LoadModelFromSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLItemSimilarityRecommenderClass.class), objc.Sel("loadModelFromSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLItemSimilarityRecommender/loadModelFromSpecificationWithCompilationOptions:options:error:
func (_MLItemSimilarityRecommenderClass MLItemSimilarityRecommenderClass) LoadModelFromSpecificationWithCompilationOptionsOptionsError(options unsafe.Pointer, options2 objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLItemSimilarityRecommenderClass.class), objc.Sel("loadModelFromSpecificationWithCompilationOptions:options:error:"), options, options2, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLItemSimilarityRecommender/metadata
func (i MLItemSimilarityRecommender) Metadata() IMLModelMetadata {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("metadata"))
	return MLModelMetadataFromID(objc.ID(rv))
}

