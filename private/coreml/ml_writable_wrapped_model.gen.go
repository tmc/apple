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

// The class instance for the [MLWritableWrappedModel] class.
var (
	_MLWritableWrappedModelClass     MLWritableWrappedModelClass
	_MLWritableWrappedModelClassOnce sync.Once
)

func getMLWritableWrappedModelClass() MLWritableWrappedModelClass {
	_MLWritableWrappedModelClassOnce.Do(func() {
		_MLWritableWrappedModelClass = MLWritableWrappedModelClass{class: objc.GetClass("MLWritableWrappedModel")}
	})
	return _MLWritableWrappedModelClass
}

// GetMLWritableWrappedModelClass returns the class object for MLWritableWrappedModel.
func GetMLWritableWrappedModelClass() MLWritableWrappedModelClass {
	return getMLWritableWrappedModelClass()
}

type MLWritableWrappedModelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLWritableWrappedModelClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLWritableWrappedModelClass) Alloc() MLWritableWrappedModel {
	rv := objc.Send[MLWritableWrappedModel](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLWritableWrappedModel.WriteToURLError]
type MLWritableWrappedModel struct {
	MLWrappedModel
}

// MLWritableWrappedModelFromID constructs a [MLWritableWrappedModel] from an objc.ID.
func MLWritableWrappedModelFromID(id objc.ID) MLWritableWrappedModel {
	return MLWritableWrappedModel{MLWrappedModel: MLWrappedModelFromID(id)}
}

// Ensure MLWritableWrappedModel implements IMLWritableWrappedModel.
var _ IMLWritableWrappedModel = MLWritableWrappedModel{}

// An interface definition for the [MLWritableWrappedModel] class.
//
// # Methods
//
//   - [IMLWritableWrappedModel.WriteToURLError]
type IMLWritableWrappedModel interface {
	IMLWrappedModel

	// Topic: Methods

	WriteToURLError(url foundation.NSURL) (bool, error)
}

// Init initializes the instance.
func (m MLWritableWrappedModel) Init() MLWritableWrappedModel {
	rv := objc.Send[MLWritableWrappedModel](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLWritableWrappedModel) Autorelease() MLWritableWrappedModel {
	rv := objc.Send[MLWritableWrappedModel](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLWritableWrappedModel creates a new MLWritableWrappedModel instance.
func NewMLWritableWrappedModel() MLWritableWrappedModel {
	class := getMLWritableWrappedModelClass()
	rv := objc.Send[MLWritableWrappedModel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewWritableWrappedModelDescriptionOnlyWithSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (MLWritableWrappedModel, error) {
	var errorPtr objc.ID
	instance := getMLWritableWrappedModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initDescriptionOnlyWithSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLWritableWrappedModel{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLWritableWrappedModelFromID(rv), nil
}

func NewWritableWrappedModelInterfaceAndMetadataWithCompiledArchiveError(archive unsafe.Pointer) (MLWritableWrappedModel, error) {
	var errorPtr objc.ID
	instance := getMLWritableWrappedModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initInterfaceAndMetadataWithCompiledArchive:error:"), archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLWritableWrappedModel{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLWritableWrappedModelFromID(rv), nil
}

func NewWritableWrappedModelWithConfiguration(configuration objectivec.IObject) MLWritableWrappedModel {
	instance := getMLWritableWrappedModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	return MLWritableWrappedModelFromID(rv)
}

func NewWritableWrappedModelWithDescription(description objectivec.IObject) MLWritableWrappedModel {
	instance := getMLWritableWrappedModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:"), description)
	return MLWritableWrappedModelFromID(rv)
}

func NewWritableWrappedModelWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLWritableWrappedModel {
	instance := getMLWritableWrappedModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return MLWritableWrappedModelFromID(rv)
}

func NewWritableWrappedModelWithInnerModel(model objectivec.IObject) MLWritableWrappedModel {
	instance := getMLWritableWrappedModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithInnerModel:"), model)
	return MLWritableWrappedModelFromID(rv)
}

func NewWritableWrappedModelWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLWritableWrappedModel {
	instance := getMLWritableWrappedModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return MLWritableWrappedModelFromID(rv)
}

func (m MLWritableWrappedModel) WriteToURLError(url foundation.NSURL) (bool, error) {
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

func (_MLWritableWrappedModelClass MLWritableWrappedModelClass) WrapperAroundWritableModel(model objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLWritableWrappedModelClass.class), objc.Sel("wrapperAroundWritableModel:"), model)
	return objectivec.Object{ID: rv}
}
