// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLOneHotEncoder] class.
var (
	_MLOneHotEncoderClass     MLOneHotEncoderClass
	_MLOneHotEncoderClassOnce sync.Once
)

func getMLOneHotEncoderClass() MLOneHotEncoderClass {
	_MLOneHotEncoderClassOnce.Do(func() {
		_MLOneHotEncoderClass = MLOneHotEncoderClass{class: objc.GetClass("MLOneHotEncoder")}
	})
	return _MLOneHotEncoderClass
}

// GetMLOneHotEncoderClass returns the class object for MLOneHotEncoder.
func GetMLOneHotEncoderClass() MLOneHotEncoderClass {
	return getMLOneHotEncoderClass()
}

type MLOneHotEncoderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLOneHotEncoderClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLOneHotEncoderClass) Alloc() MLOneHotEncoder {
	rv := objc.Send[MLOneHotEncoder](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLOneHotEncoder.EncodeFeatureValue]
//   - [MLOneHotEncoder.EncodeFeatureValueIntString]
//   - [MLOneHotEncoder.FeatureEncoding]
//   - [MLOneHotEncoder.HandleUnknown]
//   - [MLOneHotEncoder.OuputSparse]
//   - [MLOneHotEncoder.UnknownDenseVector]
//   - [MLOneHotEncoder.InitWithDataTransformerNameOuputSparseHandleUnknownInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration]
//
// See: https://developer.apple.com/documentation/CoreML/MLOneHotEncoder
type MLOneHotEncoder struct {
	MLModelEngine
}

// MLOneHotEncoderFromID constructs a [MLOneHotEncoder] from an objc.ID.
func MLOneHotEncoderFromID(id objc.ID) MLOneHotEncoder {
	return MLOneHotEncoder{MLModelEngine: MLModelEngineFromID(id)}
}

// Ensure MLOneHotEncoder implements IMLOneHotEncoder.
var _ IMLOneHotEncoder = MLOneHotEncoder{}

// An interface definition for the [MLOneHotEncoder] class.
//
// # Methods
//
//   - [IMLOneHotEncoder.EncodeFeatureValue]
//   - [IMLOneHotEncoder.EncodeFeatureValueIntString]
//   - [IMLOneHotEncoder.FeatureEncoding]
//   - [IMLOneHotEncoder.HandleUnknown]
//   - [IMLOneHotEncoder.OuputSparse]
//   - [IMLOneHotEncoder.UnknownDenseVector]
//   - [IMLOneHotEncoder.InitWithDataTransformerNameOuputSparseHandleUnknownInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration]
//
// See: https://developer.apple.com/documentation/CoreML/MLOneHotEncoder
type IMLOneHotEncoder interface {
	IMLModelEngine

	// Topic: Methods

	EncodeFeatureValue(value objectivec.IObject) objectivec.IObject
	EncodeFeatureValueIntString(string_ uint64) objectivec.IObject
	FeatureEncoding() foundation.INSOrderedSet
	HandleUnknown() bool
	OuputSparse() bool
	UnknownDenseVector() objectivec.IObject
	InitWithDataTransformerNameOuputSparseHandleUnknownInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(with objectivec.IObject, name objectivec.IObject, sparse bool, unknown bool, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLOneHotEncoder
}

// Init initializes the instance.
func (o MLOneHotEncoder) Init() MLOneHotEncoder {
	rv := objc.Send[MLOneHotEncoder](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o MLOneHotEncoder) Autorelease() MLOneHotEncoder {
	rv := objc.Send[MLOneHotEncoder](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLOneHotEncoder creates a new MLOneHotEncoder instance.
func NewMLOneHotEncoder() MLOneHotEncoder {
	class := getMLOneHotEncoderClass()
	rv := objc.Send[MLOneHotEncoder](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLOneHotEncoder/initWith:dataTransformerName:ouputSparse:handleUnknown:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:
func NewOneHotEncoderWithDataTransformerNameOuputSparseHandleUnknownInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(with objectivec.IObject, name objectivec.IObject, sparse bool, unknown bool, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLOneHotEncoder {
	instance := getMLOneHotEncoderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWith:dataTransformerName:ouputSparse:handleUnknown:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), with, name, sparse, unknown, description, description2, names, names2, configuration)
	return MLOneHotEncoderFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelEngine/initWithDescription:configuration:
func NewOneHotEncoderWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLOneHotEncoder {
	instance := getMLOneHotEncoderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return MLOneHotEncoderFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelEngine/initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:
func NewOneHotEncoderWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLOneHotEncoder {
	instance := getMLOneHotEncoderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return MLOneHotEncoderFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLOneHotEncoder/encodeFeatureValue:
func (o MLOneHotEncoder) EncodeFeatureValue(value objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("encodeFeatureValue:"), value)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLOneHotEncoder/encodeFeatureValueIntString:
func (o MLOneHotEncoder) EncodeFeatureValueIntString(string_ uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("encodeFeatureValueIntString:"), string_)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLOneHotEncoder/unknownDenseVector
func (o MLOneHotEncoder) UnknownDenseVector() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("unknownDenseVector"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLOneHotEncoder/initWith:dataTransformerName:ouputSparse:handleUnknown:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:
func (o MLOneHotEncoder) InitWithDataTransformerNameOuputSparseHandleUnknownInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(with objectivec.IObject, name objectivec.IObject, sparse bool, unknown bool, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLOneHotEncoder {
	rv := objc.Send[MLOneHotEncoder](o.ID, objc.Sel("initWith:dataTransformerName:ouputSparse:handleUnknown:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), with, name, sparse, unknown, description, description2, names, names2, configuration)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLOneHotEncoder/featureEncoderFrom:dataTransformerName:ouputSparse:handleUnknown:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:
func (_MLOneHotEncoderClass MLOneHotEncoderClass) FeatureEncoderFromDataTransformerNameOuputSparseHandleUnknownInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNames(from objectivec.IObject, name objectivec.IObject, sparse bool, unknown bool, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLOneHotEncoderClass.class), objc.Sel("featureEncoderFrom:dataTransformerName:ouputSparse:handleUnknown:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:"), from, name, sparse, unknown, description, description2, names, names2)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLOneHotEncoder/featureEncoderFrom:inputDescription:orderedInputFeatureNames:
func (_MLOneHotEncoderClass MLOneHotEncoderClass) FeatureEncoderFromInputDescriptionOrderedInputFeatureNames(from objectivec.IObject, description objectivec.IObject, names objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLOneHotEncoderClass.class), objc.Sel("featureEncoderFrom:inputDescription:orderedInputFeatureNames:"), from, description, names)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLOneHotEncoder/featureEncoderFrom:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:
func (_MLOneHotEncoderClass MLOneHotEncoderClass) FeatureEncoderFromInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNames(from objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLOneHotEncoderClass.class), objc.Sel("featureEncoderFrom:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:"), from, description, description2, names, names2)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLOneHotEncoder/loadModelFromSpecification:configuration:error:
func (_MLOneHotEncoderClass MLOneHotEncoderClass) LoadModelFromSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLOneHotEncoderClass.class), objc.Sel("loadModelFromSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLOneHotEncoder/featureEncoding
func (o MLOneHotEncoder) FeatureEncoding() foundation.INSOrderedSet {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("featureEncoding"))
	return foundation.NSOrderedSetFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLOneHotEncoder/handleUnknown
func (o MLOneHotEncoder) HandleUnknown() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("handleUnknown"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLOneHotEncoder/ouputSparse
func (o MLOneHotEncoder) OuputSparse() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("ouputSparse"))
	return rv
}
