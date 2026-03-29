// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLImputer] class.
var (
	_MLImputerClass     MLImputerClass
	_MLImputerClassOnce sync.Once
)

func getMLImputerClass() MLImputerClass {
	_MLImputerClassOnce.Do(func() {
		_MLImputerClass = MLImputerClass{class: objc.GetClass("MLImputer")}
	})
	return _MLImputerClass
}

// GetMLImputerClass returns the class object for MLImputer.
func GetMLImputerClass() MLImputerClass {
	return getMLImputerClass()
}

type MLImputerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLImputerClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLImputerClass) Alloc() MLImputer {
	rv := objc.Send[MLImputer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLImputer.ImputeValue]
//   - [MLImputer.ReplaceValue]
//   - [MLImputer.InitWithImputeValueReplaceValueInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfigurationError]
// See: https://developer.apple.com/documentation/CoreML/MLImputer
type MLImputer struct {
	MLModelEngine
}

// MLImputerFromID constructs a [MLImputer] from an objc.ID.
func MLImputerFromID(id objc.ID) MLImputer {
	return MLImputer{MLModelEngine: MLModelEngineFromID(id)}
}
// Ensure MLImputer implements IMLImputer.
var _ IMLImputer = MLImputer{}

// An interface definition for the [MLImputer] class.
//
// # Methods
//
//   - [IMLImputer.ImputeValue]
//   - [IMLImputer.ReplaceValue]
//   - [IMLImputer.InitWithImputeValueReplaceValueInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfigurationError]
//
// See: https://developer.apple.com/documentation/CoreML/MLImputer
type IMLImputer interface {
	IMLModelEngine

	// Topic: Methods

	ImputeValue() IMLFeatureValue
	ReplaceValue() IMLFeatureValue
	InitWithImputeValueReplaceValueInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfigurationError(with objectivec.IObject, value objectivec.IObject, value2 objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) (MLImputer, error)
}

// Init initializes the instance.
func (i MLImputer) Init() MLImputer {
	rv := objc.Send[MLImputer](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MLImputer) Autorelease() MLImputer {
	rv := objc.Send[MLImputer](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLImputer creates a new MLImputer instance.
func NewMLImputer() MLImputer {
	class := getMLImputerClass()
	rv := objc.Send[MLImputer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModelEngine/initWithDescription:configuration:
func NewImputerWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLImputer {
	instance := getMLImputerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return MLImputerFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLImputer/initWith:imputeValue:replaceValue:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:error:
func NewImputerWithImputeValueReplaceValueInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfigurationError(with objectivec.IObject, value objectivec.IObject, value2 objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) (MLImputer, error) {
	var errorPtr objc.ID
	instance := getMLImputerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWith:imputeValue:replaceValue:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:error:"), with, value, value2, description, description2, names, names2, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLImputer{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLImputerFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLModelEngine/initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:
func NewImputerWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLImputer {
	instance := getMLImputerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return MLImputerFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLImputer/initWith:imputeValue:replaceValue:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:error:
func (i MLImputer) InitWithImputeValueReplaceValueInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfigurationError(with objectivec.IObject, value objectivec.IObject, value2 objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) (MLImputer, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](i.ID, objc.Sel("initWith:imputeValue:replaceValue:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:error:"), with, value, value2, description, description2, names, names2, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLImputer{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLImputerFromID(rv), nil

}

//
// See: https://developer.apple.com/documentation/CoreML/MLImputer/imputeValueFrom:replaceValue:dataTransformerName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:error:
func (_MLImputerClass MLImputerClass) ImputeValueFromReplaceValueDataTransformerNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesError(from objectivec.IObject, value objectivec.IObject, name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLImputerClass.class), objc.Sel("imputeValueFrom:replaceValue:dataTransformerName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:error:"), from, value, name, description, description2, names, names2, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLImputer/imputeValueFrom:replaceValue:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:error:
func (_MLImputerClass MLImputerClass) ImputeValueFromReplaceValueInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesError(from objectivec.IObject, value objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLImputerClass.class), objc.Sel("imputeValueFrom:replaceValue:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:error:"), from, value, description, description2, names, names2, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLImputer/loadModelFromSpecification:configuration:error:
func (_MLImputerClass MLImputerClass) LoadModelFromSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLImputerClass.class), objc.Sel("loadModelFromSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLImputer/imputeValue
func (i MLImputer) ImputeValue() IMLFeatureValue {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("imputeValue"))
	return MLFeatureValueFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLImputer/replaceValue
func (i MLImputer) ReplaceValue() IMLFeatureValue {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("replaceValue"))
	return MLFeatureValueFromID(objc.ID(rv))
}

