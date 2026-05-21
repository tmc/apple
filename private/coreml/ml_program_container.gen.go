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

// The class instance for the [MLProgramContainer] class.
var (
	_MLProgramContainerClass     MLProgramContainerClass
	_MLProgramContainerClassOnce sync.Once
)

func getMLProgramContainerClass() MLProgramContainerClass {
	_MLProgramContainerClassOnce.Do(func() {
		_MLProgramContainerClass = MLProgramContainerClass{class: objc.GetClass("MLProgramContainer")}
	})
	return _MLProgramContainerClass
}

// GetMLProgramContainerClass returns the class object for MLProgramContainer.
func GetMLProgramContainerClass() MLProgramContainerClass {
	return getMLProgramContainerClass()
}

type MLProgramContainerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLProgramContainerClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLProgramContainerClass) Alloc() MLProgramContainer {
	rv := objc.Send[MLProgramContainer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

type MLProgramContainer struct {
	MLNeuralNetworkContainer
}

// MLProgramContainerFromID constructs a [MLProgramContainer] from an objc.ID.
func MLProgramContainerFromID(id objc.ID) MLProgramContainer {
	return MLProgramContainer{MLNeuralNetworkContainer: MLNeuralNetworkContainerFromID(id)}
}

// Ensure MLProgramContainer implements IMLProgramContainer.
var _ IMLProgramContainer = MLProgramContainer{}

// An interface definition for the [MLProgramContainer] class.
type IMLProgramContainer interface {
	IMLNeuralNetworkContainer
}

// Init initializes the instance.
func (m MLProgramContainer) Init() MLProgramContainer {
	rv := objc.Send[MLProgramContainer](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLProgramContainer) Autorelease() MLProgramContainer {
	rv := objc.Send[MLProgramContainer](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLProgramContainer creates a new MLProgramContainer instance.
func NewMLProgramContainer() MLProgramContainer {
	class := getMLProgramContainerClass()
	rv := objc.Send[MLProgramContainer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewProgramContainerWithFeatureDescriptionsModelDescriptionOutputLayerNamesClassScoreVectorNameClassLabelsIsEncryptedModelVersionInfo(descriptions objectivec.IObject, description objectivec.IObject, names objectivec.IObject, name objectivec.IObject, labels objectivec.IObject, encrypted bool, info objectivec.IObject) MLProgramContainer {
	instance := getMLProgramContainerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFeatureDescriptions:modelDescription:outputLayerNames:classScoreVectorName:classLabels:isEncrypted:modelVersionInfo:"), descriptions, description, names, name, labels, encrypted, info)
	return MLProgramContainerFromID(rv)
}

func NewProgramContainerWithFeatureDescriptionsModelDescriptionOutputLayerNamesClassScoreVectorNameClassLabelsIsEncryptedModelVersionInfoCompilerVersionInfo(descriptions objectivec.IObject, description objectivec.IObject, names objectivec.IObject, name objectivec.IObject, labels objectivec.IObject, encrypted bool, info objectivec.IObject, info2 objectivec.IObject) MLProgramContainer {
	instance := getMLProgramContainerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFeatureDescriptions:modelDescription:outputLayerNames:classScoreVectorName:classLabels:isEncrypted:modelVersionInfo:compilerVersionInfo:"), descriptions, description, names, name, labels, encrypted, info, info2)
	return MLProgramContainerFromID(rv)
}

func NewProgramContainerWithFilePathInputLayerNamesOutputLayerNamesParameters(path objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, parameters objectivec.IObject) MLProgramContainer {
	instance := getMLProgramContainerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFilePath:inputLayerNames:outputLayerNames:parameters:"), path, names, names2, parameters)
	return MLProgramContainerFromID(rv)
}

func (_MLProgramContainerClass MLProgramContainerClass) PopulateInputNameToShapeMapFromContainerForFunctionProgramWithValidationError(map_ unsafe.Pointer, container objectivec.IObject, function unsafe.Pointer, program unsafe.Pointer, validation bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_MLProgramContainerClass.class), objc.Sel("populateInputNameToShapeMap:fromContainer:forFunction:program:withValidation:error:"), map_, container, function, program, validation, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("populateInputNameToShapeMap:fromContainer:forFunction:program:withValidation:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (_MLProgramContainerClass MLProgramContainerClass) UpdateOptionalDefaultValueParametersInContainerUsingProgramFunctionNameModelDescription(container objectivec.IObject, program unsafe.Pointer, name objectivec.IObject, description objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_MLProgramContainerClass.class), objc.Sel("updateOptionalDefaultValueParametersInContainer:usingProgram:functionName:modelDescription:"), container, program, name, description)
}
