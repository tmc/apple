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

// The class instance for the [MLProgramE5Container] class.
var (
	_MLProgramE5ContainerClass     MLProgramE5ContainerClass
	_MLProgramE5ContainerClassOnce sync.Once
)

func getMLProgramE5ContainerClass() MLProgramE5ContainerClass {
	_MLProgramE5ContainerClassOnce.Do(func() {
		_MLProgramE5ContainerClass = MLProgramE5ContainerClass{class: objc.GetClass("MLProgramE5Container")}
	})
	return _MLProgramE5ContainerClass
}

// GetMLProgramE5ContainerClass returns the class object for MLProgramE5Container.
func GetMLProgramE5ContainerClass() MLProgramE5ContainerClass {
	return getMLProgramE5ContainerClass()
}

type MLProgramE5ContainerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLProgramE5ContainerClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLProgramE5ContainerClass) Alloc() MLProgramE5Container {
	rv := objc.Send[MLProgramE5Container](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLProgramE5Container.URLOfMILText]
//   - [MLProgramE5Container.ClassScoreVectorNameOfFunctionNamed]
//   - [MLProgramE5Container.CompilerOutput]
//   - [MLProgramE5Container.CompilerVersionInfo]
//   - [MLProgramE5Container.FindPrecompiledE5BundleAndReturnError]
//   - [MLProgramE5Container.FunctionInfoArray]
//   - [MLProgramE5Container.ModelAssetDescription]
//   - [MLProgramE5Container.ModelVersionInfo]
//   - [MLProgramE5Container.OptionalInputDefaultValuesForFunctionNamed]
//   - [MLProgramE5Container.InitWithCompiledArchiveModelVersionInfoCompilerVersionInfoError]
//
// See: https://developer.apple.com/documentation/CoreML/MLProgramE5Container
type MLProgramE5Container struct {
	objectivec.Object
}

// MLProgramE5ContainerFromID constructs a [MLProgramE5Container] from an objc.ID.
func MLProgramE5ContainerFromID(id objc.ID) MLProgramE5Container {
	return MLProgramE5Container{objectivec.Object{ID: id}}
}

// Ensure MLProgramE5Container implements IMLProgramE5Container.
var _ IMLProgramE5Container = MLProgramE5Container{}

// An interface definition for the [MLProgramE5Container] class.
//
// # Methods
//
//   - [IMLProgramE5Container.URLOfMILText]
//   - [IMLProgramE5Container.ClassScoreVectorNameOfFunctionNamed]
//   - [IMLProgramE5Container.CompilerOutput]
//   - [IMLProgramE5Container.CompilerVersionInfo]
//   - [IMLProgramE5Container.FindPrecompiledE5BundleAndReturnError]
//   - [IMLProgramE5Container.FunctionInfoArray]
//   - [IMLProgramE5Container.ModelAssetDescription]
//   - [IMLProgramE5Container.ModelVersionInfo]
//   - [IMLProgramE5Container.OptionalInputDefaultValuesForFunctionNamed]
//   - [IMLProgramE5Container.InitWithCompiledArchiveModelVersionInfoCompilerVersionInfoError]
//
// See: https://developer.apple.com/documentation/CoreML/MLProgramE5Container
type IMLProgramE5Container interface {
	objectivec.IObject

	// Topic: Methods

	URLOfMILText() foundation.INSURL
	ClassScoreVectorNameOfFunctionNamed(named objectivec.IObject) objectivec.IObject
	CompilerOutput() IMLCompilerNeuralNetworkOutput
	CompilerVersionInfo() IMLVersionInfo
	FindPrecompiledE5BundleAndReturnError() (objectivec.IObject, error)
	FunctionInfoArray() foundation.INSArray
	ModelAssetDescription() IMLModelAssetDescription
	ModelVersionInfo() IMLVersionInfo
	OptionalInputDefaultValuesForFunctionNamed(named objectivec.IObject) objectivec.IObject
	InitWithCompiledArchiveModelVersionInfoCompilerVersionInfoError(archive unsafe.Pointer, info objectivec.IObject, info2 objectivec.IObject) (MLProgramE5Container, error)
}

// Init initializes the instance.
func (p MLProgramE5Container) Init() MLProgramE5Container {
	rv := objc.Send[MLProgramE5Container](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p MLProgramE5Container) Autorelease() MLProgramE5Container {
	rv := objc.Send[MLProgramE5Container](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLProgramE5Container creates a new MLProgramE5Container instance.
func NewMLProgramE5Container() MLProgramE5Container {
	class := getMLProgramE5ContainerClass()
	rv := objc.Send[MLProgramE5Container](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLProgramE5Container/initWithCompiledArchive:modelVersionInfo:compilerVersionInfo:error:
func NewProgramE5ContainerWithCompiledArchiveModelVersionInfoCompilerVersionInfoError(archive unsafe.Pointer, info objectivec.IObject, info2 objectivec.IObject) (MLProgramE5Container, error) {
	var errorPtr objc.ID
	instance := getMLProgramE5ContainerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCompiledArchive:modelVersionInfo:compilerVersionInfo:error:"), archive, info, info2, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLProgramE5Container{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLProgramE5ContainerFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLProgramE5Container/classScoreVectorNameOfFunctionNamed:
func (p MLProgramE5Container) ClassScoreVectorNameOfFunctionNamed(named objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("classScoreVectorNameOfFunctionNamed:"), named)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLProgramE5Container/findPrecompiledE5BundleAndReturnError:
func (p MLProgramE5Container) FindPrecompiledE5BundleAndReturnError() (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](p.ID, objc.Sel("findPrecompiledE5BundleAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLProgramE5Container/optionalInputDefaultValuesForFunctionNamed:
func (p MLProgramE5Container) OptionalInputDefaultValuesForFunctionNamed(named objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("optionalInputDefaultValuesForFunctionNamed:"), named)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLProgramE5Container/initWithCompiledArchive:modelVersionInfo:compilerVersionInfo:error:
func (p MLProgramE5Container) InitWithCompiledArchiveModelVersionInfoCompilerVersionInfoError(archive unsafe.Pointer, info objectivec.IObject, info2 objectivec.IObject) (MLProgramE5Container, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](p.ID, objc.Sel("initWithCompiledArchive:modelVersionInfo:compilerVersionInfo:error:"), archive, info, info2, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLProgramE5Container{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLProgramE5ContainerFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/MLProgramE5Container/_getDefaultFunctionName:modelDescription:fromModelAssetDescription:
func (_MLProgramE5ContainerClass MLProgramE5ContainerClass) _getDefaultFunctionNameModelDescriptionFromModelAssetDescription(name []objectivec.IObject, description []objectivec.IObject, description2 objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_MLProgramE5ContainerClass.class), objc.Sel("_getDefaultFunctionName:modelDescription:fromModelAssetDescription:"), objectivec.IObjectSliceToNSArray(name), objectivec.IObjectSliceToNSArray(description), description2)
}

// GetDefaultFunctionNameModelDescriptionFromModelAssetDescription is an exported wrapper for the private method _getDefaultFunctionNameModelDescriptionFromModelAssetDescription.
func (_MLProgramE5ContainerClass MLProgramE5ContainerClass) GetDefaultFunctionNameModelDescriptionFromModelAssetDescription(name []objectivec.IObject, description []objectivec.IObject, description2 objectivec.IObject) {
	_MLProgramE5ContainerClass._getDefaultFunctionNameModelDescriptionFromModelAssetDescription(name, description, description2)
}

// See: https://developer.apple.com/documentation/CoreML/MLProgramE5Container/deduceFunctionNameToCompute:modelDescription:fromConfiguration:modelAssetDescription:error:
func (_MLProgramE5ContainerClass MLProgramE5ContainerClass) DeduceFunctionNameToComputeModelDescriptionFromConfigurationModelAssetDescriptionError(compute []objectivec.IObject, description []objectivec.IObject, configuration objectivec.IObject, description2 objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_MLProgramE5ContainerClass.class), objc.Sel("deduceFunctionNameToCompute:modelDescription:fromConfiguration:modelAssetDescription:error:"), objectivec.IObjectSliceToNSArray(compute), objectivec.IObjectSliceToNSArray(description), configuration, description2, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("deduceFunctionNameToCompute:modelDescription:fromConfiguration:modelAssetDescription:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLProgramE5Container/URLOfMILText
func (p MLProgramE5Container) URLOfMILText() foundation.INSURL {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("URLOfMILText"))
	return foundation.NSURLFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLProgramE5Container/compilerOutput
func (p MLProgramE5Container) CompilerOutput() IMLCompilerNeuralNetworkOutput {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("compilerOutput"))
	return MLCompilerNeuralNetworkOutputFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLProgramE5Container/compilerVersionInfo
func (p MLProgramE5Container) CompilerVersionInfo() IMLVersionInfo {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("compilerVersionInfo"))
	return MLVersionInfoFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLProgramE5Container/functionInfoArray
func (p MLProgramE5Container) FunctionInfoArray() foundation.INSArray {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("functionInfoArray"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLProgramE5Container/modelAssetDescription
func (p MLProgramE5Container) ModelAssetDescription() IMLModelAssetDescription {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("modelAssetDescription"))
	return MLModelAssetDescriptionFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLProgramE5Container/modelVersionInfo
func (p MLProgramE5Container) ModelVersionInfo() IMLVersionInfo {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("modelVersionInfo"))
	return MLVersionInfoFromID(objc.ID(rv))
}
