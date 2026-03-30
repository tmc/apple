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

// See: https://developer.apple.com/documentation/CoreML/MLProgramContainer
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
//
// See: https://developer.apple.com/documentation/CoreML/MLProgramContainer
type IMLProgramContainer interface {
	IMLNeuralNetworkContainer
}

// Init initializes the instance.
func (p MLProgramContainer) Init() MLProgramContainer {
	rv := objc.Send[MLProgramContainer](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p MLProgramContainer) Autorelease() MLProgramContainer {
	rv := objc.Send[MLProgramContainer](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLProgramContainer creates a new MLProgramContainer instance.
func NewMLProgramContainer() MLProgramContainer {
	class := getMLProgramContainerClass()
	rv := objc.Send[MLProgramContainer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkContainer/initWithFeatureDescriptions:modelDescription:outputLayerNames:classScoreVectorName:classLabels:isEncrypted:modelVersionInfo:
func NewProgramContainerWithFeatureDescriptionsModelDescriptionOutputLayerNamesClassScoreVectorNameClassLabelsIsEncryptedModelVersionInfo(descriptions objectivec.IObject, description objectivec.IObject, names objectivec.IObject, name objectivec.IObject, labels objectivec.IObject, encrypted bool, info objectivec.IObject) MLProgramContainer {
	instance := getMLProgramContainerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFeatureDescriptions:modelDescription:outputLayerNames:classScoreVectorName:classLabels:isEncrypted:modelVersionInfo:"), descriptions, description, names, name, labels, encrypted, info)
	return MLProgramContainerFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkContainer/initWithFeatureDescriptions:modelDescription:outputLayerNames:classScoreVectorName:classLabels:isEncrypted:modelVersionInfo:compilerVersionInfo:
func NewProgramContainerWithFeatureDescriptionsModelDescriptionOutputLayerNamesClassScoreVectorNameClassLabelsIsEncryptedModelVersionInfoCompilerVersionInfo(descriptions objectivec.IObject, description objectivec.IObject, names objectivec.IObject, name objectivec.IObject, labels objectivec.IObject, encrypted bool, info objectivec.IObject, info2 objectivec.IObject) MLProgramContainer {
	instance := getMLProgramContainerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFeatureDescriptions:modelDescription:outputLayerNames:classScoreVectorName:classLabels:isEncrypted:modelVersionInfo:compilerVersionInfo:"), descriptions, description, names, name, labels, encrypted, info, info2)
	return MLProgramContainerFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkContainer/initWithFilePath:inputLayerNames:outputLayerNames:parameters:
func NewProgramContainerWithFilePathInputLayerNamesOutputLayerNamesParameters(path objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, parameters objectivec.IObject) MLProgramContainer {
	instance := getMLProgramContainerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFilePath:inputLayerNames:outputLayerNames:parameters:"), path, names, names2, parameters)
	return MLProgramContainerFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLProgramContainer/loadProgramAtURL:error:
func (_MLProgramContainerClass MLProgramContainerClass) LoadProgramAtURLError(url foundation.INSURL) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLProgramContainerClass.class), objc.Sel("loadProgramAtURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLProgramContainer/loadProgramFromCompiledArchive:error:
func (_MLProgramContainerClass MLProgramContainerClass) LoadProgramFromCompiledArchiveError(archive unsafe.Pointer) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLProgramContainerClass.class), objc.Sel("loadProgramFromCompiledArchive:error:"), archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLProgramContainer/populateInputNameToShapeMap:fromContainer:forFunction:program:withValidation:error:
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

// See: https://developer.apple.com/documentation/CoreML/MLProgramContainer/updateOptionalDefaultValueParametersInContainer:usingProgram:functionName:modelDescription:
func (_MLProgramContainerClass MLProgramContainerClass) UpdateOptionalDefaultValueParametersInContainerUsingProgramFunctionNameModelDescription(container objectivec.IObject, program unsafe.Pointer, name objectivec.IObject, description objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_MLProgramContainerClass.class), objc.Sel("updateOptionalDefaultValueParametersInContainer:usingProgram:functionName:modelDescription:"), container, program, name, description)
}
