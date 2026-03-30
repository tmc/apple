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

// The class instance for the [MLE5RangeShapeExecutionStreamOperationPool] class.
var (
	_MLE5RangeShapeExecutionStreamOperationPoolClass     MLE5RangeShapeExecutionStreamOperationPoolClass
	_MLE5RangeShapeExecutionStreamOperationPoolClassOnce sync.Once
)

func getMLE5RangeShapeExecutionStreamOperationPoolClass() MLE5RangeShapeExecutionStreamOperationPoolClass {
	_MLE5RangeShapeExecutionStreamOperationPoolClassOnce.Do(func() {
		_MLE5RangeShapeExecutionStreamOperationPoolClass = MLE5RangeShapeExecutionStreamOperationPoolClass{class: objc.GetClass("MLE5RangeShapeExecutionStreamOperationPool")}
	})
	return _MLE5RangeShapeExecutionStreamOperationPoolClass
}

// GetMLE5RangeShapeExecutionStreamOperationPoolClass returns the class object for MLE5RangeShapeExecutionStreamOperationPool.
func GetMLE5RangeShapeExecutionStreamOperationPoolClass() MLE5RangeShapeExecutionStreamOperationPoolClass {
	return getMLE5RangeShapeExecutionStreamOperationPoolClass()
}

type MLE5RangeShapeExecutionStreamOperationPoolClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLE5RangeShapeExecutionStreamOperationPoolClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLE5RangeShapeExecutionStreamOperationPoolClass) Alloc() MLE5RangeShapeExecutionStreamOperationPool {
	rv := objc.Send[MLE5RangeShapeExecutionStreamOperationPool](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLE5RangeShapeExecutionStreamOperationPool._makeAndPreloadOperationForFunctionError]
//   - [MLE5RangeShapeExecutionStreamOperationPool._putBack]
//   - [MLE5RangeShapeExecutionStreamOperationPool._takeOutAnyOperation]
//   - [MLE5RangeShapeExecutionStreamOperationPool.CompilerVersionInfo]
//   - [MLE5RangeShapeExecutionStreamOperationPool.Configuration]
//   - [MLE5RangeShapeExecutionStreamOperationPool.DefaultShapePool]
//   - [MLE5RangeShapeExecutionStreamOperationPool.MilDefaultShapeFunctionName]
//   - [MLE5RangeShapeExecutionStreamOperationPool.MilFunctionName]
//   - [MLE5RangeShapeExecutionStreamOperationPool.ModelDescription]
//   - [MLE5RangeShapeExecutionStreamOperationPool.ModelSignpostId]
//   - [MLE5RangeShapeExecutionStreamOperationPool.PrepareWithInitialPoolSizeError]
//   - [MLE5RangeShapeExecutionStreamOperationPool.ProgramLibrary]
//   - [MLE5RangeShapeExecutionStreamOperationPool.PutBack]
//   - [MLE5RangeShapeExecutionStreamOperationPool.SerialQueue]
//   - [MLE5RangeShapeExecutionStreamOperationPool.ShapeHashToPoolMap]
//   - [MLE5RangeShapeExecutionStreamOperationPool.TakeOutOperationForFeaturesError]
//   - [MLE5RangeShapeExecutionStreamOperationPool.InitWithProgramLibraryFunctionNameModelDescriptionConfigurationModelSignpostIdCompilerVersionInfo]
//   - [MLE5RangeShapeExecutionStreamOperationPool.DebugDescription]
//   - [MLE5RangeShapeExecutionStreamOperationPool.Description]
//   - [MLE5RangeShapeExecutionStreamOperationPool.Hash]
//   - [MLE5RangeShapeExecutionStreamOperationPool.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLE5RangeShapeExecutionStreamOperationPool
type MLE5RangeShapeExecutionStreamOperationPool struct {
	objectivec.Object
}

// MLE5RangeShapeExecutionStreamOperationPoolFromID constructs a [MLE5RangeShapeExecutionStreamOperationPool] from an objc.ID.
func MLE5RangeShapeExecutionStreamOperationPoolFromID(id objc.ID) MLE5RangeShapeExecutionStreamOperationPool {
	return MLE5RangeShapeExecutionStreamOperationPool{objectivec.Object{ID: id}}
}

// Ensure MLE5RangeShapeExecutionStreamOperationPool implements IMLE5RangeShapeExecutionStreamOperationPool.
var _ IMLE5RangeShapeExecutionStreamOperationPool = MLE5RangeShapeExecutionStreamOperationPool{}

// An interface definition for the [MLE5RangeShapeExecutionStreamOperationPool] class.
//
// # Methods
//
//   - [IMLE5RangeShapeExecutionStreamOperationPool._makeAndPreloadOperationForFunctionError]
//   - [IMLE5RangeShapeExecutionStreamOperationPool._putBack]
//   - [IMLE5RangeShapeExecutionStreamOperationPool._takeOutAnyOperation]
//   - [IMLE5RangeShapeExecutionStreamOperationPool.CompilerVersionInfo]
//   - [IMLE5RangeShapeExecutionStreamOperationPool.Configuration]
//   - [IMLE5RangeShapeExecutionStreamOperationPool.DefaultShapePool]
//   - [IMLE5RangeShapeExecutionStreamOperationPool.MilDefaultShapeFunctionName]
//   - [IMLE5RangeShapeExecutionStreamOperationPool.MilFunctionName]
//   - [IMLE5RangeShapeExecutionStreamOperationPool.ModelDescription]
//   - [IMLE5RangeShapeExecutionStreamOperationPool.ModelSignpostId]
//   - [IMLE5RangeShapeExecutionStreamOperationPool.PrepareWithInitialPoolSizeError]
//   - [IMLE5RangeShapeExecutionStreamOperationPool.ProgramLibrary]
//   - [IMLE5RangeShapeExecutionStreamOperationPool.PutBack]
//   - [IMLE5RangeShapeExecutionStreamOperationPool.SerialQueue]
//   - [IMLE5RangeShapeExecutionStreamOperationPool.ShapeHashToPoolMap]
//   - [IMLE5RangeShapeExecutionStreamOperationPool.TakeOutOperationForFeaturesError]
//   - [IMLE5RangeShapeExecutionStreamOperationPool.InitWithProgramLibraryFunctionNameModelDescriptionConfigurationModelSignpostIdCompilerVersionInfo]
//   - [IMLE5RangeShapeExecutionStreamOperationPool.DebugDescription]
//   - [IMLE5RangeShapeExecutionStreamOperationPool.Description]
//   - [IMLE5RangeShapeExecutionStreamOperationPool.Hash]
//   - [IMLE5RangeShapeExecutionStreamOperationPool.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLE5RangeShapeExecutionStreamOperationPool
type IMLE5RangeShapeExecutionStreamOperationPool interface {
	objectivec.IObject

	// Topic: Methods

	_makeAndPreloadOperationForFunctionError(function objectivec.IObject) (objectivec.IObject, error)
	_putBack(back objectivec.IObject)
	_takeOutAnyOperation() objectivec.IObject
	CompilerVersionInfo() IMLVersionInfo
	Configuration() IMLModelConfiguration
	DefaultShapePool() foundation.INSSet
	MilDefaultShapeFunctionName() string
	MilFunctionName() string
	ModelDescription() IMLModelDescription
	ModelSignpostId() uint64
	PrepareWithInitialPoolSizeError(size int64) (bool, error)
	ProgramLibrary() IMLE5ProgramLibrary
	PutBack(back objectivec.IObject)
	SerialQueue() objectivec.Object
	ShapeHashToPoolMap() foundation.INSDictionary
	TakeOutOperationForFeaturesError(features objectivec.IObject) (objectivec.IObject, error)
	InitWithProgramLibraryFunctionNameModelDescriptionConfigurationModelSignpostIdCompilerVersionInfo(library objectivec.IObject, name objectivec.IObject, description objectivec.IObject, configuration objectivec.IObject, id uint64, info objectivec.IObject) MLE5RangeShapeExecutionStreamOperationPool
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (e MLE5RangeShapeExecutionStreamOperationPool) Init() MLE5RangeShapeExecutionStreamOperationPool {
	rv := objc.Send[MLE5RangeShapeExecutionStreamOperationPool](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e MLE5RangeShapeExecutionStreamOperationPool) Autorelease() MLE5RangeShapeExecutionStreamOperationPool {
	rv := objc.Send[MLE5RangeShapeExecutionStreamOperationPool](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLE5RangeShapeExecutionStreamOperationPool creates a new MLE5RangeShapeExecutionStreamOperationPool instance.
func NewMLE5RangeShapeExecutionStreamOperationPool() MLE5RangeShapeExecutionStreamOperationPool {
	class := getMLE5RangeShapeExecutionStreamOperationPoolClass()
	rv := objc.Send[MLE5RangeShapeExecutionStreamOperationPool](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLE5RangeShapeExecutionStreamOperationPool/initWithProgramLibrary:functionName:modelDescription:configuration:modelSignpostId:compilerVersionInfo:
func NewE5RangeShapeExecutionStreamOperationPoolWithProgramLibraryFunctionNameModelDescriptionConfigurationModelSignpostIdCompilerVersionInfo(library objectivec.IObject, name objectivec.IObject, description objectivec.IObject, configuration objectivec.IObject, id uint64, info objectivec.IObject) MLE5RangeShapeExecutionStreamOperationPool {
	instance := getMLE5RangeShapeExecutionStreamOperationPoolClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithProgramLibrary:functionName:modelDescription:configuration:modelSignpostId:compilerVersionInfo:"), library, name, description, configuration, id, info)
	return MLE5RangeShapeExecutionStreamOperationPoolFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5RangeShapeExecutionStreamOperationPool/_makeAndPreloadOperationForFunction:error:
func (e MLE5RangeShapeExecutionStreamOperationPool) _makeAndPreloadOperationForFunctionError(function objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](e.ID, objc.Sel("_makeAndPreloadOperationForFunction:error:"), function, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// MakeAndPreloadOperationForFunctionError is an exported wrapper for the private method _makeAndPreloadOperationForFunctionError.
func (e MLE5RangeShapeExecutionStreamOperationPool) MakeAndPreloadOperationForFunctionError(function objectivec.IObject) (objectivec.IObject, error) {
	return e._makeAndPreloadOperationForFunctionError(function)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5RangeShapeExecutionStreamOperationPool/_putBack:
func (e MLE5RangeShapeExecutionStreamOperationPool) _putBack(back objectivec.IObject) {
	objc.Send[objc.ID](e.ID, objc.Sel("_putBack:"), back)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5RangeShapeExecutionStreamOperationPool/_takeOutAnyOperation
func (e MLE5RangeShapeExecutionStreamOperationPool) _takeOutAnyOperation() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("_takeOutAnyOperation"))
	return objectivec.Object{ID: rv}
}

// TakeOutAnyOperation is an exported wrapper for the private method _takeOutAnyOperation.
func (e MLE5RangeShapeExecutionStreamOperationPool) TakeOutAnyOperation() objectivec.IObject {
	return e._takeOutAnyOperation()
}

// See: https://developer.apple.com/documentation/CoreML/MLE5RangeShapeExecutionStreamOperationPool/prepareWithInitialPoolSize:error:
func (e MLE5RangeShapeExecutionStreamOperationPool) PrepareWithInitialPoolSizeError(size int64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](e.ID, objc.Sel("prepareWithInitialPoolSize:error:"), size, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("prepareWithInitialPoolSize:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLE5RangeShapeExecutionStreamOperationPool/putBack:
func (e MLE5RangeShapeExecutionStreamOperationPool) PutBack(back objectivec.IObject) {
	objc.Send[objc.ID](e.ID, objc.Sel("putBack:"), back)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5RangeShapeExecutionStreamOperationPool/takeOutOperationForFeatures:error:
func (e MLE5RangeShapeExecutionStreamOperationPool) TakeOutOperationForFeaturesError(features objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](e.ID, objc.Sel("takeOutOperationForFeatures:error:"), features, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLE5RangeShapeExecutionStreamOperationPool/initWithProgramLibrary:functionName:modelDescription:configuration:modelSignpostId:compilerVersionInfo:
func (e MLE5RangeShapeExecutionStreamOperationPool) InitWithProgramLibraryFunctionNameModelDescriptionConfigurationModelSignpostIdCompilerVersionInfo(library objectivec.IObject, name objectivec.IObject, description objectivec.IObject, configuration objectivec.IObject, id uint64, info objectivec.IObject) MLE5RangeShapeExecutionStreamOperationPool {
	rv := objc.Send[MLE5RangeShapeExecutionStreamOperationPool](e.ID, objc.Sel("initWithProgramLibrary:functionName:modelDescription:configuration:modelSignpostId:compilerVersionInfo:"), library, name, description, configuration, id, info)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLE5RangeShapeExecutionStreamOperationPool/compilerVersionInfo
func (e MLE5RangeShapeExecutionStreamOperationPool) CompilerVersionInfo() IMLVersionInfo {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("compilerVersionInfo"))
	return MLVersionInfoFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5RangeShapeExecutionStreamOperationPool/configuration
func (e MLE5RangeShapeExecutionStreamOperationPool) Configuration() IMLModelConfiguration {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("configuration"))
	return MLModelConfigurationFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5RangeShapeExecutionStreamOperationPool/debugDescription
func (e MLE5RangeShapeExecutionStreamOperationPool) DebugDescription() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLE5RangeShapeExecutionStreamOperationPool/defaultShapePool
func (e MLE5RangeShapeExecutionStreamOperationPool) DefaultShapePool() foundation.INSSet {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("defaultShapePool"))
	return foundation.NSSetFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5RangeShapeExecutionStreamOperationPool/description
func (e MLE5RangeShapeExecutionStreamOperationPool) Description() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLE5RangeShapeExecutionStreamOperationPool/hash
func (e MLE5RangeShapeExecutionStreamOperationPool) Hash() uint64 {
	rv := objc.Send[uint64](e.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLE5RangeShapeExecutionStreamOperationPool/milDefaultShapeFunctionName
func (e MLE5RangeShapeExecutionStreamOperationPool) MilDefaultShapeFunctionName() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("milDefaultShapeFunctionName"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLE5RangeShapeExecutionStreamOperationPool/milFunctionName
func (e MLE5RangeShapeExecutionStreamOperationPool) MilFunctionName() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("milFunctionName"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLE5RangeShapeExecutionStreamOperationPool/modelDescription
func (e MLE5RangeShapeExecutionStreamOperationPool) ModelDescription() IMLModelDescription {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("modelDescription"))
	return MLModelDescriptionFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5RangeShapeExecutionStreamOperationPool/modelSignpostId
func (e MLE5RangeShapeExecutionStreamOperationPool) ModelSignpostId() uint64 {
	rv := objc.Send[uint64](e.ID, objc.Sel("modelSignpostId"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLE5RangeShapeExecutionStreamOperationPool/programLibrary
func (e MLE5RangeShapeExecutionStreamOperationPool) ProgramLibrary() IMLE5ProgramLibrary {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("programLibrary"))
	return MLE5ProgramLibraryFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5RangeShapeExecutionStreamOperationPool/serialQueue
func (e MLE5RangeShapeExecutionStreamOperationPool) SerialQueue() objectivec.Object {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("serialQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5RangeShapeExecutionStreamOperationPool/shapeHashToPoolMap
func (e MLE5RangeShapeExecutionStreamOperationPool) ShapeHashToPoolMap() foundation.INSDictionary {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("shapeHashToPoolMap"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5RangeShapeExecutionStreamOperationPool/superclass
func (e MLE5RangeShapeExecutionStreamOperationPool) Superclass() objc.Class {
	rv := objc.Send[objc.Class](e.ID, objc.Sel("superclass"))
	return rv
}
