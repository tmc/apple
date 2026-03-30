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

// The class instance for the [MLDelegateUpdatableModel] class.
var (
	_MLDelegateUpdatableModelClass     MLDelegateUpdatableModelClass
	_MLDelegateUpdatableModelClassOnce sync.Once
)

func getMLDelegateUpdatableModelClass() MLDelegateUpdatableModelClass {
	_MLDelegateUpdatableModelClassOnce.Do(func() {
		_MLDelegateUpdatableModelClass = MLDelegateUpdatableModelClass{class: objc.GetClass("MLDelegateUpdatableModel")}
	})
	return _MLDelegateUpdatableModelClass
}

// GetMLDelegateUpdatableModelClass returns the class object for MLDelegateUpdatableModel.
func GetMLDelegateUpdatableModelClass() MLDelegateUpdatableModelClass {
	return getMLDelegateUpdatableModelClass()
}

type MLDelegateUpdatableModelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLDelegateUpdatableModelClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLDelegateUpdatableModelClass) Alloc() MLDelegateUpdatableModel {
	rv := objc.Send[MLDelegateUpdatableModel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLDelegateUpdatableModel.CancelUpdate]
//   - [MLDelegateUpdatableModel.ResumeUpdate]
//   - [MLDelegateUpdatableModel.ResumeUpdateWithParameters]
//   - [MLDelegateUpdatableModel.SetUpdateProgressHandlersDispatchQueue]
//   - [MLDelegateUpdatableModel.UpdatableEngine]
//   - [MLDelegateUpdatableModel.UpdateModelWithData]
//   - [MLDelegateUpdatableModel.WriteToURLError]
//   - [MLDelegateUpdatableModel.Metadata]
//
// See: https://developer.apple.com/documentation/CoreML/MLDelegateUpdatableModel
type MLDelegateUpdatableModel struct {
	MLDelegateModel
}

// MLDelegateUpdatableModelFromID constructs a [MLDelegateUpdatableModel] from an objc.ID.
func MLDelegateUpdatableModelFromID(id objc.ID) MLDelegateUpdatableModel {
	return MLDelegateUpdatableModel{MLDelegateModel: MLDelegateModelFromID(id)}
}

// Ensure MLDelegateUpdatableModel implements IMLDelegateUpdatableModel.
var _ IMLDelegateUpdatableModel = MLDelegateUpdatableModel{}

// An interface definition for the [MLDelegateUpdatableModel] class.
//
// # Methods
//
//   - [IMLDelegateUpdatableModel.CancelUpdate]
//   - [IMLDelegateUpdatableModel.ResumeUpdate]
//   - [IMLDelegateUpdatableModel.ResumeUpdateWithParameters]
//   - [IMLDelegateUpdatableModel.SetUpdateProgressHandlersDispatchQueue]
//   - [IMLDelegateUpdatableModel.UpdatableEngine]
//   - [IMLDelegateUpdatableModel.UpdateModelWithData]
//   - [IMLDelegateUpdatableModel.WriteToURLError]
//   - [IMLDelegateUpdatableModel.Metadata]
//
// See: https://developer.apple.com/documentation/CoreML/MLDelegateUpdatableModel
type IMLDelegateUpdatableModel interface {
	IMLDelegateModel

	// Topic: Methods

	CancelUpdate()
	ResumeUpdate()
	ResumeUpdateWithParameters(parameters objectivec.IObject)
	SetUpdateProgressHandlersDispatchQueue(handlers ErrorHandler, queue objectivec.IObject)
	UpdatableEngine() objectivec.IObject
	UpdateModelWithData(data objectivec.IObject)
	WriteToURLError(url foundation.INSURL) (bool, error)
	Metadata() IMLModelMetadata
}

// Init initializes the instance.
func (d MLDelegateUpdatableModel) Init() MLDelegateUpdatableModel {
	rv := objc.Send[MLDelegateUpdatableModel](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d MLDelegateUpdatableModel) Autorelease() MLDelegateUpdatableModel {
	rv := objc.Send[MLDelegateUpdatableModel](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLDelegateUpdatableModel creates a new MLDelegateUpdatableModel instance.
func NewMLDelegateUpdatableModel() MLDelegateUpdatableModel {
	class := getMLDelegateUpdatableModelClass()
	rv := objc.Send[MLDelegateUpdatableModel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModel/initDescriptionOnlyWithSpecification:configuration:error:
func NewDelegateUpdatableModelDescriptionOnlyWithSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (MLDelegateUpdatableModel, error) {
	var errorPtr objc.ID
	instance := getMLDelegateUpdatableModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initDescriptionOnlyWithSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLDelegateUpdatableModel{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLDelegateUpdatableModelFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLModel/initInterfaceAndMetadataWithCompiledArchive:error:
func NewDelegateUpdatableModelInterfaceAndMetadataWithCompiledArchiveError(archive unsafe.Pointer) (MLDelegateUpdatableModel, error) {
	var errorPtr objc.ID
	instance := getMLDelegateUpdatableModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initInterfaceAndMetadataWithCompiledArchive:error:"), archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLDelegateUpdatableModel{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLDelegateUpdatableModelFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithConfiguration:
func NewDelegateUpdatableModelWithConfiguration(configuration objectivec.IObject) MLDelegateUpdatableModel {
	instance := getMLDelegateUpdatableModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	return MLDelegateUpdatableModelFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithDescription:
func NewDelegateUpdatableModelWithDescription(description objectivec.IObject) MLDelegateUpdatableModel {
	instance := getMLDelegateUpdatableModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:"), description)
	return MLDelegateUpdatableModelFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithDescription:configuration:
func NewDelegateUpdatableModelWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLDelegateUpdatableModel {
	instance := getMLDelegateUpdatableModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return MLDelegateUpdatableModelFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLDelegateUpdatableModel/initWithEngine:error:
func NewDelegateUpdatableModelWithEngineError(engine objectivec.IObject) (MLDelegateUpdatableModel, error) {
	var errorPtr objc.ID
	instance := getMLDelegateUpdatableModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithEngine:error:"), engine, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLDelegateUpdatableModel{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLDelegateUpdatableModelFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:
func NewDelegateUpdatableModelWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLDelegateUpdatableModel {
	instance := getMLDelegateUpdatableModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return MLDelegateUpdatableModelFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLDelegateUpdatableModel/cancelUpdate
func (d MLDelegateUpdatableModel) CancelUpdate() {
	objc.Send[objc.ID](d.ID, objc.Sel("cancelUpdate"))
}

// See: https://developer.apple.com/documentation/CoreML/MLDelegateUpdatableModel/resumeUpdate
func (d MLDelegateUpdatableModel) ResumeUpdate() {
	objc.Send[objc.ID](d.ID, objc.Sel("resumeUpdate"))
}

// See: https://developer.apple.com/documentation/CoreML/MLDelegateUpdatableModel/resumeUpdateWithParameters:
func (d MLDelegateUpdatableModel) ResumeUpdateWithParameters(parameters objectivec.IObject) {
	objc.Send[objc.ID](d.ID, objc.Sel("resumeUpdateWithParameters:"), parameters)
}

// See: https://developer.apple.com/documentation/CoreML/MLDelegateUpdatableModel/setUpdateProgressHandlers:dispatchQueue:
func (d MLDelegateUpdatableModel) SetUpdateProgressHandlersDispatchQueue(handlers ErrorHandler, queue objectivec.IObject) {
	_block0, _ := NewErrorBlock(handlers)
	objc.Send[objc.ID](d.ID, objc.Sel("setUpdateProgressHandlers:dispatchQueue:"), _block0, queue)
}

// See: https://developer.apple.com/documentation/CoreML/MLDelegateUpdatableModel/updateModelWithData:
func (d MLDelegateUpdatableModel) UpdateModelWithData(data objectivec.IObject) {
	objc.Send[objc.ID](d.ID, objc.Sel("updateModelWithData:"), data)
}

// See: https://developer.apple.com/documentation/CoreML/MLDelegateUpdatableModel/writeToURL:error:
func (d MLDelegateUpdatableModel) WriteToURLError(url foundation.INSURL) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("writeToURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("writeToURL:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLDelegateUpdatableModel/loadModelFromCompiledArchive:modelVersionInfo:compilerVersionInfo:configuration:error:
func (_MLDelegateUpdatableModelClass MLDelegateUpdatableModelClass) LoadModelFromCompiledArchiveModelVersionInfoCompilerVersionInfoConfigurationError(archive unsafe.Pointer, info objectivec.IObject, info2 objectivec.IObject, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLDelegateUpdatableModelClass.class), objc.Sel("loadModelFromCompiledArchive:modelVersionInfo:compilerVersionInfo:configuration:error:"), archive, info, info2, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLDelegateUpdatableModel/metadata
func (d MLDelegateUpdatableModel) Metadata() IMLModelMetadata {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("metadata"))
	return MLModelMetadataFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLDelegateUpdatableModel/updatableEngine
func (d MLDelegateUpdatableModel) UpdatableEngine() objectivec.IObject {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("updatableEngine"))
	return objectivec.Object{ID: rv}
}
