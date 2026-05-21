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
type IMLDelegateUpdatableModel interface {
	IMLDelegateModel

	// Topic: Methods

	CancelUpdate()
	ResumeUpdate()
	ResumeUpdateWithParameters(parameters objectivec.IObject)
	SetUpdateProgressHandlersDispatchQueue(handlers objectivec.IObject, queue objectivec.IObject)
	UpdatableEngine() unsafe.Pointer
	UpdateModelWithData(data objectivec.IObject)
	WriteToURLError(url foundation.NSURL) (bool, error)
}

// Init initializes the instance.
func (m MLDelegateUpdatableModel) Init() MLDelegateUpdatableModel {
	rv := objc.Send[MLDelegateUpdatableModel](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLDelegateUpdatableModel) Autorelease() MLDelegateUpdatableModel {
	rv := objc.Send[MLDelegateUpdatableModel](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLDelegateUpdatableModel creates a new MLDelegateUpdatableModel instance.
func NewMLDelegateUpdatableModel() MLDelegateUpdatableModel {
	class := getMLDelegateUpdatableModelClass()
	rv := objc.Send[MLDelegateUpdatableModel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

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

func NewDelegateUpdatableModelWithConfiguration(configuration objectivec.IObject) MLDelegateUpdatableModel {
	instance := getMLDelegateUpdatableModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	return MLDelegateUpdatableModelFromID(rv)
}

func NewDelegateUpdatableModelWithDescription(description objectivec.IObject) MLDelegateUpdatableModel {
	instance := getMLDelegateUpdatableModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:"), description)
	return MLDelegateUpdatableModelFromID(rv)
}

func NewDelegateUpdatableModelWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLDelegateUpdatableModel {
	instance := getMLDelegateUpdatableModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return MLDelegateUpdatableModelFromID(rv)
}

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

func NewDelegateUpdatableModelWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLDelegateUpdatableModel {
	instance := getMLDelegateUpdatableModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return MLDelegateUpdatableModelFromID(rv)
}

func (m MLDelegateUpdatableModel) CancelUpdate() {
	objc.Send[objc.ID](m.ID, objc.Sel("cancelUpdate"))
}
func (m MLDelegateUpdatableModel) ResumeUpdate() {
	objc.Send[objc.ID](m.ID, objc.Sel("resumeUpdate"))
}
func (m MLDelegateUpdatableModel) ResumeUpdateWithParameters(parameters objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("resumeUpdateWithParameters:"), parameters)
}
func (m MLDelegateUpdatableModel) SetUpdateProgressHandlersDispatchQueue(handlers objectivec.IObject, queue objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("setUpdateProgressHandlers:dispatchQueue:"), handlers, queue)
}
func (m MLDelegateUpdatableModel) UpdateModelWithData(data objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("updateModelWithData:"), data)
}
func (m MLDelegateUpdatableModel) WriteToURLError(url foundation.NSURL) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("writeToURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("writeToURL:error: returned NO with nil NSError")
	}
	return rv, nil

}

func (_MLDelegateUpdatableModelClass MLDelegateUpdatableModelClass) LoadModelFromCompiledArchiveModelVersionInfoCompilerVersionInfoConfigurationError(archive MLModelInputArchiverRef, info objectivec.IObject, info2 objectivec.IObject, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLDelegateUpdatableModelClass.class), objc.Sel("loadModelFromCompiledArchive:modelVersionInfo:compilerVersionInfo:configuration:error:"), archive, info, info2, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

func (m MLDelegateUpdatableModel) UpdatableEngine() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m.ID, objc.Sel("updatableEngine"))
	return rv
}
