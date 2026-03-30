// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLFeatureVectorizer] class.
var (
	_MLFeatureVectorizerClass     MLFeatureVectorizerClass
	_MLFeatureVectorizerClassOnce sync.Once
)

func getMLFeatureVectorizerClass() MLFeatureVectorizerClass {
	_MLFeatureVectorizerClassOnce.Do(func() {
		_MLFeatureVectorizerClass = MLFeatureVectorizerClass{class: objc.GetClass("MLFeatureVectorizer")}
	})
	return _MLFeatureVectorizerClass
}

// GetMLFeatureVectorizerClass returns the class object for MLFeatureVectorizer.
func GetMLFeatureVectorizerClass() MLFeatureVectorizerClass {
	return getMLFeatureVectorizerClass()
}

type MLFeatureVectorizerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLFeatureVectorizerClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLFeatureVectorizerClass) Alloc() MLFeatureVectorizer {
	rv := objc.Send[MLFeatureVectorizer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLFeatureVectorizer.ColumnNameEncoding]
//   - [MLFeatureVectorizer.DimensionEncoding]
//   - [MLFeatureVectorizer.VectorizeOneHotEncoderDictIndexError]
//   - [MLFeatureVectorizer.InitWithDimensionEncodingDataTransformerNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration]
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureVectorizer
type MLFeatureVectorizer struct {
	MLModelEngine
}

// MLFeatureVectorizerFromID constructs a [MLFeatureVectorizer] from an objc.ID.
func MLFeatureVectorizerFromID(id objc.ID) MLFeatureVectorizer {
	return MLFeatureVectorizer{MLModelEngine: MLModelEngineFromID(id)}
}

// Ensure MLFeatureVectorizer implements IMLFeatureVectorizer.
var _ IMLFeatureVectorizer = MLFeatureVectorizer{}

// An interface definition for the [MLFeatureVectorizer] class.
//
// # Methods
//
//   - [IMLFeatureVectorizer.ColumnNameEncoding]
//   - [IMLFeatureVectorizer.DimensionEncoding]
//   - [IMLFeatureVectorizer.VectorizeOneHotEncoderDictIndexError]
//   - [IMLFeatureVectorizer.InitWithDimensionEncodingDataTransformerNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration]
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureVectorizer
type IMLFeatureVectorizer interface {
	IMLModelEngine

	// Topic: Methods

	ColumnNameEncoding() foundation.INSArray
	DimensionEncoding() foundation.INSArray
	VectorizeOneHotEncoderDictIndexError(dict objectivec.IObject, index uint64) (objectivec.IObject, error)
	InitWithDimensionEncodingDataTransformerNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(with objectivec.IObject, encoding objectivec.IObject, name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLFeatureVectorizer
}

// Init initializes the instance.
func (f MLFeatureVectorizer) Init() MLFeatureVectorizer {
	rv := objc.Send[MLFeatureVectorizer](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f MLFeatureVectorizer) Autorelease() MLFeatureVectorizer {
	rv := objc.Send[MLFeatureVectorizer](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLFeatureVectorizer creates a new MLFeatureVectorizer instance.
func NewMLFeatureVectorizer() MLFeatureVectorizer {
	class := getMLFeatureVectorizerClass()
	rv := objc.Send[MLFeatureVectorizer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelEngine/initWithDescription:configuration:
func NewFeatureVectorizerWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLFeatureVectorizer {
	instance := getMLFeatureVectorizerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return MLFeatureVectorizerFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLFeatureVectorizer/initWith:dimensionEncoding:dataTransformerName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:
func NewFeatureVectorizerWithDimensionEncodingDataTransformerNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(with objectivec.IObject, encoding objectivec.IObject, name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLFeatureVectorizer {
	instance := getMLFeatureVectorizerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWith:dimensionEncoding:dataTransformerName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), with, encoding, name, description, description2, names, names2, configuration)
	return MLFeatureVectorizerFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelEngine/initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:
func NewFeatureVectorizerWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLFeatureVectorizer {
	instance := getMLFeatureVectorizerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return MLFeatureVectorizerFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLFeatureVectorizer/vectorizeOneHotEncoderDict:index:error:
func (f MLFeatureVectorizer) VectorizeOneHotEncoderDictIndexError(dict objectivec.IObject, index uint64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](f.ID, objc.Sel("vectorizeOneHotEncoderDict:index:error:"), dict, index, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLFeatureVectorizer/initWith:dimensionEncoding:dataTransformerName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:
func (f MLFeatureVectorizer) InitWithDimensionEncodingDataTransformerNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(with objectivec.IObject, encoding objectivec.IObject, name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLFeatureVectorizer {
	rv := objc.Send[MLFeatureVectorizer](f.ID, objc.Sel("initWith:dimensionEncoding:dataTransformerName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), with, encoding, name, description, description2, names, names2, configuration)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLFeatureVectorizer/loadModelFromSpecification:configuration:error:
func (_MLFeatureVectorizerClass MLFeatureVectorizerClass) LoadModelFromSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLFeatureVectorizerClass.class), objc.Sel("loadModelFromSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLFeatureVectorizer/columnNameEncoding
func (f MLFeatureVectorizer) ColumnNameEncoding() foundation.INSArray {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("columnNameEncoding"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLFeatureVectorizer/dimensionEncoding
func (f MLFeatureVectorizer) DimensionEncoding() foundation.INSArray {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("dimensionEncoding"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
