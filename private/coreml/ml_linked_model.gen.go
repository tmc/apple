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

// The class instance for the [MLLinkedModel] class.
var (
	_MLLinkedModelClass     MLLinkedModelClass
	_MLLinkedModelClassOnce sync.Once
)

func getMLLinkedModelClass() MLLinkedModelClass {
	_MLLinkedModelClassOnce.Do(func() {
		_MLLinkedModelClass = MLLinkedModelClass{class: objc.GetClass("MLLinkedModel")}
	})
	return _MLLinkedModelClass
}

// GetMLLinkedModelClass returns the class object for MLLinkedModel.
func GetMLLinkedModelClass() MLLinkedModelClass {
	return getMLLinkedModelClass()
}

type MLLinkedModelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLLinkedModelClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLLinkedModelClass) Alloc() MLLinkedModel {
	rv := objc.Send[MLLinkedModel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLLinkedModel.LinkedModel]
//   - [MLLinkedModel.SetLinkedModel]
//   - [MLLinkedModel.ModelFileName]
//   - [MLLinkedModel.SetModelFileName]
//   - [MLLinkedModel.ModelSearchPath]
//   - [MLLinkedModel.SetModelSearchPath]
//   - [MLLinkedModel.ParameterValueForKeyError]
//   - [MLLinkedModel.PredictionFromFeaturesOptionsError]
//   - [MLLinkedModel.PredictionsFromBatchOptionsError]
//   - [MLLinkedModel.UpdateParameterDescriptionsByUnarchivingSpecification]
//   - [MLLinkedModel.InitWithLinkedModelModelFileNameModelSearchPathConfiguration]
//
// See: https://developer.apple.com/documentation/CoreML/MLLinkedModel
type MLLinkedModel struct {
	MLModel
}

// MLLinkedModelFromID constructs a [MLLinkedModel] from an objc.ID.
func MLLinkedModelFromID(id objc.ID) MLLinkedModel {
	return MLLinkedModel{MLModel: MLModelFromID(id)}
}

// Ensure MLLinkedModel implements IMLLinkedModel.
var _ IMLLinkedModel = MLLinkedModel{}

// An interface definition for the [MLLinkedModel] class.
//
// # Methods
//
//   - [IMLLinkedModel.LinkedModel]
//   - [IMLLinkedModel.SetLinkedModel]
//   - [IMLLinkedModel.ModelFileName]
//   - [IMLLinkedModel.SetModelFileName]
//   - [IMLLinkedModel.ModelSearchPath]
//   - [IMLLinkedModel.SetModelSearchPath]
//   - [IMLLinkedModel.ParameterValueForKeyError]
//   - [IMLLinkedModel.PredictionFromFeaturesOptionsError]
//   - [IMLLinkedModel.PredictionsFromBatchOptionsError]
//   - [IMLLinkedModel.UpdateParameterDescriptionsByUnarchivingSpecification]
//   - [IMLLinkedModel.InitWithLinkedModelModelFileNameModelSearchPathConfiguration]
//
// See: https://developer.apple.com/documentation/CoreML/MLLinkedModel
type IMLLinkedModel interface {
	IMLModel

	// Topic: Methods

	LinkedModel() IMLModel
	SetLinkedModel(value IMLModel)
	ModelFileName() string
	SetModelFileName(value string)
	ModelSearchPath() string
	SetModelSearchPath(value string)
	ParameterValueForKeyError(key objectivec.IObject) (objectivec.IObject, error)
	PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	PredictionsFromBatchOptionsError(batch objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	UpdateParameterDescriptionsByUnarchivingSpecification(specification unsafe.Pointer)
	InitWithLinkedModelModelFileNameModelSearchPathConfiguration(model objectivec.IObject, name objectivec.IObject, path objectivec.IObject, configuration objectivec.IObject) MLLinkedModel
}

// Init initializes the instance.
func (l MLLinkedModel) Init() MLLinkedModel {
	rv := objc.Send[MLLinkedModel](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l MLLinkedModel) Autorelease() MLLinkedModel {
	rv := objc.Send[MLLinkedModel](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLLinkedModel creates a new MLLinkedModel instance.
func NewMLLinkedModel() MLLinkedModel {
	class := getMLLinkedModelClass()
	rv := objc.Send[MLLinkedModel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModel/initDescriptionOnlyWithSpecification:configuration:error:
func NewLinkedModelDescriptionOnlyWithSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (MLLinkedModel, error) {
	var errorPtr objc.ID
	instance := getMLLinkedModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initDescriptionOnlyWithSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLLinkedModel{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLLinkedModelFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLModel/initInterfaceAndMetadataWithCompiledArchive:error:
func NewLinkedModelInterfaceAndMetadataWithCompiledArchiveError(archive unsafe.Pointer) (MLLinkedModel, error) {
	var errorPtr objc.ID
	instance := getMLLinkedModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initInterfaceAndMetadataWithCompiledArchive:error:"), archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLLinkedModel{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLLinkedModelFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithConfiguration:
func NewLinkedModelWithConfiguration(configuration objectivec.IObject) MLLinkedModel {
	instance := getMLLinkedModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	return MLLinkedModelFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithDescription:
func NewLinkedModelWithDescription(description objectivec.IObject) MLLinkedModel {
	instance := getMLLinkedModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:"), description)
	return MLLinkedModelFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithDescription:configuration:
func NewLinkedModelWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLLinkedModel {
	instance := getMLLinkedModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return MLLinkedModelFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLLinkedModel/initWithLinkedModel:modelFileName:modelSearchPath:configuration:
func NewLinkedModelWithLinkedModelModelFileNameModelSearchPathConfiguration(model objectivec.IObject, name objectivec.IObject, path objectivec.IObject, configuration objectivec.IObject) MLLinkedModel {
	instance := getMLLinkedModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLinkedModel:modelFileName:modelSearchPath:configuration:"), model, name, path, configuration)
	return MLLinkedModelFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:
func NewLinkedModelWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLLinkedModel {
	instance := getMLLinkedModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return MLLinkedModelFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLLinkedModel/parameterValueForKey:error:
func (l MLLinkedModel) ParameterValueForKeyError(key objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](l.ID, objc.Sel("parameterValueForKey:error:"), key, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLLinkedModel/predictionFromFeatures:options:error:
func (l MLLinkedModel) PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](l.ID, objc.Sel("predictionFromFeatures:options:error:"), features, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLLinkedModel/predictionsFromBatch:options:error:
func (l MLLinkedModel) PredictionsFromBatchOptionsError(batch objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](l.ID, objc.Sel("predictionsFromBatch:options:error:"), batch, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLLinkedModel/updateParameterDescriptionsByUnarchivingSpecification:
func (l MLLinkedModel) UpdateParameterDescriptionsByUnarchivingSpecification(specification unsafe.Pointer) {
	objc.Send[objc.ID](l.ID, objc.Sel("updateParameterDescriptionsByUnarchivingSpecification:"), specification)
}

// See: https://developer.apple.com/documentation/CoreML/MLLinkedModel/initWithLinkedModel:modelFileName:modelSearchPath:configuration:
func (l MLLinkedModel) InitWithLinkedModelModelFileNameModelSearchPathConfiguration(model objectivec.IObject, name objectivec.IObject, path objectivec.IObject, configuration objectivec.IObject) MLLinkedModel {
	rv := objc.Send[MLLinkedModel](l.ID, objc.Sel("initWithLinkedModel:modelFileName:modelSearchPath:configuration:"), model, name, path, configuration)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLLinkedModel/areFeaturesIn:modelNamed:aSubsetOf:error:
func (_MLLinkedModelClass MLLinkedModelClass) AreFeaturesInModelNamedASubsetOfError(in objectivec.IObject, named objectivec.IObject, of objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_MLLinkedModelClass.class), objc.Sel("areFeaturesIn:modelNamed:aSubsetOf:error:"), in, named, of, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("areFeaturesIn:modelNamed:aSubsetOf:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLLinkedModel/findFile:inSearchPath:basePath:
func (_MLLinkedModelClass MLLinkedModelClass) FindFileInSearchPathBasePath(file objectivec.IObject, path objectivec.IObject, path2 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLLinkedModelClass.class), objc.Sel("findFile:inSearchPath:basePath:"), file, path, path2)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLLinkedModel/loadModelFromSpecification:configuration:error:
func (_MLLinkedModelClass MLLinkedModelClass) LoadModelFromSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLLinkedModelClass.class), objc.Sel("loadModelFromSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLLinkedModel/linkedModel
func (l MLLinkedModel) LinkedModel() IMLModel {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("linkedModel"))
	return MLModelFromID(objc.ID(rv))
}
func (l MLLinkedModel) SetLinkedModel(value IMLModel) {
	objc.Send[struct{}](l.ID, objc.Sel("setLinkedModel:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLLinkedModel/modelFileName
func (l MLLinkedModel) ModelFileName() string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("modelFileName"))
	return foundation.NSStringFromID(rv).String()
}
func (l MLLinkedModel) SetModelFileName(value string) {
	objc.Send[struct{}](l.ID, objc.Sel("setModelFileName:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/CoreML/MLLinkedModel/modelSearchPath
func (l MLLinkedModel) ModelSearchPath() string {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("modelSearchPath"))
	return foundation.NSStringFromID(rv).String()
}
func (l MLLinkedModel) SetModelSearchPath(value string) {
	objc.Send[struct{}](l.ID, objc.Sel("setModelSearchPath:"), objc.String(value))
}
