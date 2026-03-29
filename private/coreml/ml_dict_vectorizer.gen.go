// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLDictVectorizer] class.
var (
	_MLDictVectorizerClass     MLDictVectorizerClass
	_MLDictVectorizerClassOnce sync.Once
)

func getMLDictVectorizerClass() MLDictVectorizerClass {
	_MLDictVectorizerClassOnce.Do(func() {
		_MLDictVectorizerClass = MLDictVectorizerClass{class: objc.GetClass("MLDictVectorizer")}
	})
	return _MLDictVectorizerClass
}

// GetMLDictVectorizerClass returns the class object for MLDictVectorizer.
func GetMLDictVectorizerClass() MLDictVectorizerClass {
	return getMLDictVectorizerClass()
}

type MLDictVectorizerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLDictVectorizerClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLDictVectorizerClass) Alloc() MLDictVectorizer {
	rv := objc.Send[MLDictVectorizer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLDictVectorizer.CategoryName]
//   - [MLDictVectorizer.ConstructDictionaryError]
//   - [MLDictVectorizer.InitWithDataTransformerNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfigurationError]
// See: https://developer.apple.com/documentation/CoreML/MLDictVectorizer
type MLDictVectorizer struct {
	MLModelEngine
}

// MLDictVectorizerFromID constructs a [MLDictVectorizer] from an objc.ID.
func MLDictVectorizerFromID(id objc.ID) MLDictVectorizer {
	return MLDictVectorizer{MLModelEngine: MLModelEngineFromID(id)}
}
// Ensure MLDictVectorizer implements IMLDictVectorizer.
var _ IMLDictVectorizer = MLDictVectorizer{}

// An interface definition for the [MLDictVectorizer] class.
//
// # Methods
//
//   - [IMLDictVectorizer.CategoryName]
//   - [IMLDictVectorizer.ConstructDictionaryError]
//   - [IMLDictVectorizer.InitWithDataTransformerNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfigurationError]
//
// See: https://developer.apple.com/documentation/CoreML/MLDictVectorizer
type IMLDictVectorizer interface {
	IMLModelEngine

	// Topic: Methods

	CategoryName() foundation.INSOrderedSet
	ConstructDictionaryError(dictionary objectivec.IObject) (objectivec.IObject, error)
	InitWithDataTransformerNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfigurationError(with objectivec.IObject, name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) (MLDictVectorizer, error)
}

// Init initializes the instance.
func (d MLDictVectorizer) Init() MLDictVectorizer {
	rv := objc.Send[MLDictVectorizer](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d MLDictVectorizer) Autorelease() MLDictVectorizer {
	rv := objc.Send[MLDictVectorizer](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLDictVectorizer creates a new MLDictVectorizer instance.
func NewMLDictVectorizer() MLDictVectorizer {
	class := getMLDictVectorizerClass()
	rv := objc.Send[MLDictVectorizer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLDictVectorizer/initWith:dataTransformerName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:error:
func NewDictVectorizerWithDataTransformerNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfigurationError(with objectivec.IObject, name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) (MLDictVectorizer, error) {
	var errorPtr objc.ID
	instance := getMLDictVectorizerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWith:dataTransformerName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:error:"), with, name, description, description2, names, names2, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLDictVectorizer{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLDictVectorizerFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModelEngine/initWithDescription:configuration:
func NewDictVectorizerWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLDictVectorizer {
	instance := getMLDictVectorizerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return MLDictVectorizerFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModelEngine/initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:
func NewDictVectorizerWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLDictVectorizer {
	instance := getMLDictVectorizerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return MLDictVectorizerFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLDictVectorizer/constructDictionary:error:
func (d MLDictVectorizer) ConstructDictionaryError(dictionary objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("constructDictionary:error:"), dictionary, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLDictVectorizer/initWith:dataTransformerName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:error:
func (d MLDictVectorizer) InitWithDataTransformerNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfigurationError(with objectivec.IObject, name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) (MLDictVectorizer, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initWith:dataTransformerName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:error:"), with, name, description, description2, names, names2, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLDictVectorizer{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLDictVectorizerFromID(rv), nil

}

//
// See: https://developer.apple.com/documentation/CoreML/MLDictVectorizer/categoryName:dataTransformerName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:error:
func (_MLDictVectorizerClass MLDictVectorizerClass) CategoryNameDataTransformerNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesError(name objectivec.IObject, name2 objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLDictVectorizerClass.class), objc.Sel("categoryName:dataTransformerName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:error:"), name, name2, description, description2, names, names2, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLDictVectorizer/categoryName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:error:
func (_MLDictVectorizerClass MLDictVectorizerClass) CategoryNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesError(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLDictVectorizerClass.class), objc.Sel("categoryName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:error:"), name, description, description2, names, names2, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLDictVectorizer/loadModelFromSpecification:configuration:error:
func (_MLDictVectorizerClass MLDictVectorizerClass) LoadModelFromSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLDictVectorizerClass.class), objc.Sel("loadModelFromSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLDictVectorizer/categoryName
func (d MLDictVectorizer) CategoryName() foundation.INSOrderedSet {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("categoryName"))
	return foundation.NSOrderedSetFromID(objc.ID(rv))
}

