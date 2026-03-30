// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLNormalizer] class.
var (
	_MLNormalizerClass     MLNormalizerClass
	_MLNormalizerClassOnce sync.Once
)

func getMLNormalizerClass() MLNormalizerClass {
	_MLNormalizerClassOnce.Do(func() {
		_MLNormalizerClass = MLNormalizerClass{class: objc.GetClass("MLNormalizer")}
	})
	return _MLNormalizerClass
}

// GetMLNormalizerClass returns the class object for MLNormalizer.
func GetMLNormalizerClass() MLNormalizerClass {
	return getMLNormalizerClass()
}

type MLNormalizerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLNormalizerClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLNormalizerClass) Alloc() MLNormalizer {
	rv := objc.Send[MLNormalizer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLNormalizer.Norm]
//   - [MLNormalizer.InitWithDataTransformerNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration]
//
// See: https://developer.apple.com/documentation/CoreML/MLNormalizer
type MLNormalizer struct {
	MLModelEngine
}

// MLNormalizerFromID constructs a [MLNormalizer] from an objc.ID.
func MLNormalizerFromID(id objc.ID) MLNormalizer {
	return MLNormalizer{MLModelEngine: MLModelEngineFromID(id)}
}

// Ensure MLNormalizer implements IMLNormalizer.
var _ IMLNormalizer = MLNormalizer{}

// An interface definition for the [MLNormalizer] class.
//
// # Methods
//
//   - [IMLNormalizer.Norm]
//   - [IMLNormalizer.InitWithDataTransformerNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration]
//
// See: https://developer.apple.com/documentation/CoreML/MLNormalizer
type IMLNormalizer interface {
	IMLModelEngine

	// Topic: Methods

	Norm() int
	InitWithDataTransformerNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(with int, name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLNormalizer
}

// Init initializes the instance.
func (n MLNormalizer) Init() MLNormalizer {
	rv := objc.Send[MLNormalizer](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n MLNormalizer) Autorelease() MLNormalizer {
	rv := objc.Send[MLNormalizer](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLNormalizer creates a new MLNormalizer instance.
func NewMLNormalizer() MLNormalizer {
	class := getMLNormalizerClass()
	rv := objc.Send[MLNormalizer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLNormalizer/initWith:dataTransformerName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:
func NewNormalizerWithDataTransformerNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(with int, name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLNormalizer {
	instance := getMLNormalizerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWith:dataTransformerName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), with, name, description, description2, names, names2, configuration)
	return MLNormalizerFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelEngine/initWithDescription:configuration:
func NewNormalizerWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLNormalizer {
	instance := getMLNormalizerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return MLNormalizerFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelEngine/initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:
func NewNormalizerWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLNormalizer {
	instance := getMLNormalizerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return MLNormalizerFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLNormalizer/initWith:dataTransformerName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:
func (n MLNormalizer) InitWithDataTransformerNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(with int, name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLNormalizer {
	rv := objc.Send[MLNormalizer](n.ID, objc.Sel("initWith:dataTransformerName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), with, name, description, description2, names, names2, configuration)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLNormalizer/inputDescriptionFrom:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:
func (_MLNormalizerClass MLNormalizerClass) InputDescriptionFromOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNames(from objectivec.IObject, description objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLNormalizerClass.class), objc.Sel("inputDescriptionFrom:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:"), from, description, names, names2)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLNormalizer/loadModelFromSpecification:configuration:error:
func (_MLNormalizerClass MLNormalizerClass) LoadModelFromSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLNormalizerClass.class), objc.Sel("loadModelFromSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLNormalizer/normFrom:dataTransformerName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:
func (_MLNormalizerClass MLNormalizerClass) NormFromDataTransformerNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNames(from int, name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLNormalizerClass.class), objc.Sel("normFrom:dataTransformerName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:"), from, name, description, description2, names, names2)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLNormalizer/normFrom:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:
func (_MLNormalizerClass MLNormalizerClass) NormFromInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNames(from int, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLNormalizerClass.class), objc.Sel("normFrom:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:"), from, description, description2, names, names2)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLNormalizer/norm
func (n MLNormalizer) Norm() int {
	rv := objc.Send[int](n.ID, objc.Sel("norm"))
	return rv
}
