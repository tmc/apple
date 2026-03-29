// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLCategoricalMapping] class.
var (
	_MLCategoricalMappingClass     MLCategoricalMappingClass
	_MLCategoricalMappingClassOnce sync.Once
)

func getMLCategoricalMappingClass() MLCategoricalMappingClass {
	_MLCategoricalMappingClassOnce.Do(func() {
		_MLCategoricalMappingClass = MLCategoricalMappingClass{class: objc.GetClass("MLCategoricalMapping")}
	})
	return _MLCategoricalMappingClass
}

// GetMLCategoricalMappingClass returns the class object for MLCategoricalMapping.
func GetMLCategoricalMappingClass() MLCategoricalMappingClass {
	return getMLCategoricalMappingClass()
}

type MLCategoricalMappingClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLCategoricalMappingClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLCategoricalMappingClass) Alloc() MLCategoricalMapping {
	rv := objc.Send[MLCategoricalMapping](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLCategoricalMapping.MapFeatureError]
//   - [MLCategoricalMapping.Mapping]
//   - [MLCategoricalMapping.ValueOnUnknown]
//   - [MLCategoricalMapping.InitWithMappingValueOnUnknownDataTransformerNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration]
// See: https://developer.apple.com/documentation/CoreML/MLCategoricalMapping
type MLCategoricalMapping struct {
	MLModelEngine
}

// MLCategoricalMappingFromID constructs a [MLCategoricalMapping] from an objc.ID.
func MLCategoricalMappingFromID(id objc.ID) MLCategoricalMapping {
	return MLCategoricalMapping{MLModelEngine: MLModelEngineFromID(id)}
}
// Ensure MLCategoricalMapping implements IMLCategoricalMapping.
var _ IMLCategoricalMapping = MLCategoricalMapping{}

// An interface definition for the [MLCategoricalMapping] class.
//
// # Methods
//
//   - [IMLCategoricalMapping.MapFeatureError]
//   - [IMLCategoricalMapping.Mapping]
//   - [IMLCategoricalMapping.ValueOnUnknown]
//   - [IMLCategoricalMapping.InitWithMappingValueOnUnknownDataTransformerNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration]
//
// See: https://developer.apple.com/documentation/CoreML/MLCategoricalMapping
type IMLCategoricalMapping interface {
	IMLModelEngine

	// Topic: Methods

	MapFeatureError(feature objectivec.IObject) (objectivec.IObject, error)
	Mapping() foundation.INSDictionary
	ValueOnUnknown() IMLFeatureValue
	InitWithMappingValueOnUnknownDataTransformerNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(mapping objectivec.IObject, unknown objectivec.IObject, name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLCategoricalMapping
}

// Init initializes the instance.
func (c MLCategoricalMapping) Init() MLCategoricalMapping {
	rv := objc.Send[MLCategoricalMapping](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MLCategoricalMapping) Autorelease() MLCategoricalMapping {
	rv := objc.Send[MLCategoricalMapping](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLCategoricalMapping creates a new MLCategoricalMapping instance.
func NewMLCategoricalMapping() MLCategoricalMapping {
	class := getMLCategoricalMappingClass()
	rv := objc.Send[MLCategoricalMapping](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModelEngine/initWithDescription:configuration:
func NewCategoricalMappingWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLCategoricalMapping {
	instance := getMLCategoricalMappingClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return MLCategoricalMappingFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLCategoricalMapping/initWithMapping:valueOnUnknown:dataTransformerName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:
func NewCategoricalMappingWithMappingValueOnUnknownDataTransformerNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(mapping objectivec.IObject, unknown objectivec.IObject, name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLCategoricalMapping {
	instance := getMLCategoricalMappingClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMapping:valueOnUnknown:dataTransformerName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), mapping, unknown, name, description, description2, names, names2, configuration)
	return MLCategoricalMappingFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModelEngine/initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:
func NewCategoricalMappingWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLCategoricalMapping {
	instance := getMLCategoricalMappingClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return MLCategoricalMappingFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLCategoricalMapping/mapFeature:error:
func (c MLCategoricalMapping) MapFeatureError(feature objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](c.ID, objc.Sel("mapFeature:error:"), feature, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLCategoricalMapping/initWithMapping:valueOnUnknown:dataTransformerName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:
func (c MLCategoricalMapping) InitWithMappingValueOnUnknownDataTransformerNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(mapping objectivec.IObject, unknown objectivec.IObject, name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLCategoricalMapping {
	rv := objc.Send[MLCategoricalMapping](c.ID, objc.Sel("initWithMapping:valueOnUnknown:dataTransformerName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), mapping, unknown, name, description, description2, names, names2, configuration)
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLCategoricalMapping/loadModelFromSpecification:configuration:error:
func (_MLCategoricalMappingClass MLCategoricalMappingClass) LoadModelFromSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLCategoricalMappingClass.class), objc.Sel("loadModelFromSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLCategoricalMapping/mapping
func (c MLCategoricalMapping) Mapping() foundation.INSDictionary {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("mapping"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLCategoricalMapping/valueOnUnknown
func (c MLCategoricalMapping) ValueOnUnknown() IMLFeatureValue {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("valueOnUnknown"))
	return MLFeatureValueFromID(objc.ID(rv))
}

