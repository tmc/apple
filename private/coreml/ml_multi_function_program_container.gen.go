// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLMultiFunctionProgramContainer] class.
var (
	_MLMultiFunctionProgramContainerClass     MLMultiFunctionProgramContainerClass
	_MLMultiFunctionProgramContainerClassOnce sync.Once
)

func getMLMultiFunctionProgramContainerClass() MLMultiFunctionProgramContainerClass {
	_MLMultiFunctionProgramContainerClassOnce.Do(func() {
		_MLMultiFunctionProgramContainerClass = MLMultiFunctionProgramContainerClass{class: objc.GetClass("MLMultiFunctionProgramContainer")}
	})
	return _MLMultiFunctionProgramContainerClass
}

// GetMLMultiFunctionProgramContainerClass returns the class object for MLMultiFunctionProgramContainer.
func GetMLMultiFunctionProgramContainerClass() MLMultiFunctionProgramContainerClass {
	return getMLMultiFunctionProgramContainerClass()
}

type MLMultiFunctionProgramContainerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLMultiFunctionProgramContainerClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLMultiFunctionProgramContainerClass) Alloc() MLMultiFunctionProgramContainer {
	rv := objc.SendIfResponds[MLMultiFunctionProgramContainer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLMultiFunctionProgramContainer.FunctionNameToInputLayersNames]
//   - [MLMultiFunctionProgramContainer.SetFunctionNameToInputLayersNames]
//   - [MLMultiFunctionProgramContainer.FunctionNameToOutputLayersNames]
//   - [MLMultiFunctionProgramContainer.SetFunctionNameToOutputLayersNames]
//   - [MLMultiFunctionProgramContainer.InitWithContainerProgramError]
type MLMultiFunctionProgramContainer struct {
	MLProgramContainer
}

// MLMultiFunctionProgramContainerFromID constructs a [MLMultiFunctionProgramContainer] from an objc.ID.
func MLMultiFunctionProgramContainerFromID(id objc.ID) MLMultiFunctionProgramContainer {
	return MLMultiFunctionProgramContainer{MLProgramContainer: MLProgramContainerFromID(id)}
}

// Ensure MLMultiFunctionProgramContainer implements IMLMultiFunctionProgramContainer.
var _ IMLMultiFunctionProgramContainer = MLMultiFunctionProgramContainer{}

// An interface definition for the [MLMultiFunctionProgramContainer] class.
//
// # Methods
//
//   - [IMLMultiFunctionProgramContainer.FunctionNameToInputLayersNames]
//   - [IMLMultiFunctionProgramContainer.SetFunctionNameToInputLayersNames]
//   - [IMLMultiFunctionProgramContainer.FunctionNameToOutputLayersNames]
//   - [IMLMultiFunctionProgramContainer.SetFunctionNameToOutputLayersNames]
//   - [IMLMultiFunctionProgramContainer.InitWithContainerProgramError]
type IMLMultiFunctionProgramContainer interface {
	IMLProgramContainer

	// Topic: Methods

	FunctionNameToInputLayersNames() foundation.INSDictionary
	SetFunctionNameToInputLayersNames(value foundation.INSDictionary)
	FunctionNameToOutputLayersNames() foundation.INSDictionary
	SetFunctionNameToOutputLayersNames(value foundation.INSDictionary)
	InitWithContainerProgramError(container objectivec.IObject, program unsafe.Pointer) (MLMultiFunctionProgramContainer, error)
}

// Init initializes the instance.
func (m MLMultiFunctionProgramContainer) Init() MLMultiFunctionProgramContainer {
	rv := objc.SendIfResponds[MLMultiFunctionProgramContainer](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLMultiFunctionProgramContainer) Autorelease() MLMultiFunctionProgramContainer {
	rv := objc.SendIfResponds[MLMultiFunctionProgramContainer](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLMultiFunctionProgramContainer creates a new MLMultiFunctionProgramContainer instance.
func NewMLMultiFunctionProgramContainer() MLMultiFunctionProgramContainer {
	class := getMLMultiFunctionProgramContainerClass()
	rv := objc.SendIfResponds[MLMultiFunctionProgramContainer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewMultiFunctionProgramContainerWithContainerProgramError(container objectivec.IObject, program unsafe.Pointer) (MLMultiFunctionProgramContainer, error) {
	var errorPtr objc.ID
	instance := getMLMultiFunctionProgramContainerClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithContainer:program:error:"), container, program, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLMultiFunctionProgramContainer{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return MLMultiFunctionProgramContainer{}, objc.ErrInitFailed
	}
	return MLMultiFunctionProgramContainerFromID(rv), nil
}

func NewMultiFunctionProgramContainerWithFeatureDescriptionsModelDescriptionOutputLayerNamesClassScoreVectorNameClassLabelsIsEncryptedModelVersionInfo(descriptions objectivec.IObject, description objectivec.IObject, names objectivec.IObject, name objectivec.IObject, labels objectivec.IObject, encrypted bool, info objectivec.IObject) MLMultiFunctionProgramContainer {
	instance := getMLMultiFunctionProgramContainerClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithFeatureDescriptions:modelDescription:outputLayerNames:classScoreVectorName:classLabels:isEncrypted:modelVersionInfo:"), descriptions, description, names, name, labels, encrypted, info)
	return MLMultiFunctionProgramContainerFromID(rv)
}

func NewMultiFunctionProgramContainerWithFeatureDescriptionsModelDescriptionOutputLayerNamesClassScoreVectorNameClassLabelsIsEncryptedModelVersionInfoCompilerVersionInfo(descriptions objectivec.IObject, description objectivec.IObject, names objectivec.IObject, name objectivec.IObject, labels objectivec.IObject, encrypted bool, info objectivec.IObject, info2 objectivec.IObject) MLMultiFunctionProgramContainer {
	instance := getMLMultiFunctionProgramContainerClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithFeatureDescriptions:modelDescription:outputLayerNames:classScoreVectorName:classLabels:isEncrypted:modelVersionInfo:compilerVersionInfo:"), descriptions, description, names, name, labels, encrypted, info, info2)
	return MLMultiFunctionProgramContainerFromID(rv)
}

func NewMultiFunctionProgramContainerWithFilePathInputLayerNamesOutputLayerNamesParameters(path objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, parameters objectivec.IObject) MLMultiFunctionProgramContainer {
	instance := getMLMultiFunctionProgramContainerClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithFilePath:inputLayerNames:outputLayerNames:parameters:"), path, names, names2, parameters)
	return MLMultiFunctionProgramContainerFromID(rv)
}

func (m MLMultiFunctionProgramContainer) InitWithContainerProgramError(container objectivec.IObject, program unsafe.Pointer) (MLMultiFunctionProgramContainer, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithContainer:program:error:"), container, program, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLMultiFunctionProgramContainer{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLMultiFunctionProgramContainerFromID(rv), nil

}

func (m MLMultiFunctionProgramContainer) FunctionNameToInputLayersNames() foundation.INSDictionary {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("functionNameToInputLayersNames"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLMultiFunctionProgramContainer) SetFunctionNameToInputLayersNames(value foundation.INSDictionary) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setFunctionNameToInputLayersNames:"), value)
}
func (m MLMultiFunctionProgramContainer) FunctionNameToOutputLayersNames() foundation.INSDictionary {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("functionNameToOutputLayersNames"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLMultiFunctionProgramContainer) SetFunctionNameToOutputLayersNames(value foundation.INSDictionary) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setFunctionNameToOutputLayersNames:"), value)
}
