// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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

// # Methods
//
//   - [MLItemSimilarityRecommender._itemForIndexError]
//   - [MLItemSimilarityRecommender._mapItemSequenceDestError]
//   - [MLItemSimilarityRecommender.ModelData]
//   - [MLItemSimilarityRecommender.PredictionFromFeaturesOptionsError]
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
type IMLItemSimilarityRecommender interface {
	IMLModel

	// Topic: Methods

	_itemForIndexError(index uint64) (objectivec.IObject, error)
	_mapItemSequenceDestError(sequence objectivec.IObject, dest unsafe.Pointer) (bool, error)
	ModelData() string
	PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
}

// Init initializes the instance.
func (m MLItemSimilarityRecommender) Init() MLItemSimilarityRecommender {
	rv := objc.Send[MLItemSimilarityRecommender](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLItemSimilarityRecommender) Autorelease() MLItemSimilarityRecommender {
	rv := objc.Send[MLItemSimilarityRecommender](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLItemSimilarityRecommender creates a new MLItemSimilarityRecommender instance.
func NewMLItemSimilarityRecommender() MLItemSimilarityRecommender {
	class := getMLItemSimilarityRecommenderClass()
	rv := objc.Send[MLItemSimilarityRecommender](objc.ID(class.class), objc.Sel("new"))
	return rv
}

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

func NewItemSimilarityRecommenderWithConfiguration(configuration objectivec.IObject) MLItemSimilarityRecommender {
	instance := getMLItemSimilarityRecommenderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	return MLItemSimilarityRecommenderFromID(rv)
}

func NewItemSimilarityRecommenderWithDescription(description objectivec.IObject) MLItemSimilarityRecommender {
	instance := getMLItemSimilarityRecommenderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:"), description)
	return MLItemSimilarityRecommenderFromID(rv)
}

func NewItemSimilarityRecommenderWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLItemSimilarityRecommender {
	instance := getMLItemSimilarityRecommenderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return MLItemSimilarityRecommenderFromID(rv)
}

func NewItemSimilarityRecommenderWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLItemSimilarityRecommender {
	instance := getMLItemSimilarityRecommenderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return MLItemSimilarityRecommenderFromID(rv)
}

func (m MLItemSimilarityRecommender) _itemForIndexError(index uint64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("_itemForIndex:error:"), index, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// ItemForIndexError is an exported wrapper for the private method _itemForIndexError.
func (m MLItemSimilarityRecommender) ItemForIndexError(index uint64) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_itemForIndex:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_itemForIndex:error:"}
		return nil, err
	}
	return m._itemForIndexError(index)
}

// CanItemForIndexError reports whether the receiver responds to the private selector _itemForIndex:error:.
func (m MLItemSimilarityRecommender) CanItemForIndexError() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_itemForIndex:error:"))
}
func (m MLItemSimilarityRecommender) _mapItemSequenceDestError(sequence objectivec.IObject, dest unsafe.Pointer) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("_mapItemSequence:dest:error:"), sequence, dest, unsafe.Pointer(&errorPtr))
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
func (m MLItemSimilarityRecommender) MapItemSequenceDestError(sequence objectivec.IObject, dest unsafe.Pointer) (bool, error) {
	if !objc.RespondsToSelector(m.ID, objc.Sel("_mapItemSequence:dest:error:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_mapItemSequence:dest:error:"}
		return false, err
	}
	return m._mapItemSequenceDestError(sequence, dest)
}

// CanMapItemSequenceDestError reports whether the receiver responds to the private selector _mapItemSequence:dest:error:.
func (m MLItemSimilarityRecommender) CanMapItemSequenceDestError() bool {
	return objc.RespondsToSelector(m.ID, objc.Sel("_mapItemSequence:dest:error:"))
}
func (m MLItemSimilarityRecommender) ModelData() string {
	rv := objc.Send[*byte](m.ID, objc.Sel("modelData"))
	return objc.GoString(rv)
}
func (m MLItemSimilarityRecommender) PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictionFromFeatures:options:error:"), features, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

func (_MLItemSimilarityRecommenderClass MLItemSimilarityRecommenderClass) CompileSpecificationToArchiveOptionsError(specification unsafe.Pointer, archive unsafe.Pointer, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLItemSimilarityRecommenderClass.class), objc.Sel("compileSpecification:toArchive:options:error:"), specification, archive, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_MLItemSimilarityRecommenderClass MLItemSimilarityRecommenderClass) CompiledVersionForSpecificationOptionsError(specification unsafe.Pointer, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLItemSimilarityRecommenderClass.class), objc.Sel("compiledVersionForSpecification:options:error:"), specification, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_MLItemSimilarityRecommenderClass MLItemSimilarityRecommenderClass) LoadModelFromCompiledArchiveModelVersionInfoCompilerVersionInfoConfigurationError(archive unsafe.Pointer, info objectivec.IObject, info2 objectivec.IObject, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLItemSimilarityRecommenderClass.class), objc.Sel("loadModelFromCompiledArchive:modelVersionInfo:compilerVersionInfo:configuration:error:"), archive, info, info2, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_MLItemSimilarityRecommenderClass MLItemSimilarityRecommenderClass) LoadModelFromSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLItemSimilarityRecommenderClass.class), objc.Sel("loadModelFromSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_MLItemSimilarityRecommenderClass MLItemSimilarityRecommenderClass) LoadModelFromSpecificationWithCompilationOptionsOptionsError(options unsafe.Pointer, options2 objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLItemSimilarityRecommenderClass.class), objc.Sel("loadModelFromSpecificationWithCompilationOptions:options:error:"), options, options2, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
