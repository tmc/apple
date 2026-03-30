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

// The class instance for the [MLE5StaticShapeExecutionStreamOperationPool] class.
var (
	_MLE5StaticShapeExecutionStreamOperationPoolClass     MLE5StaticShapeExecutionStreamOperationPoolClass
	_MLE5StaticShapeExecutionStreamOperationPoolClassOnce sync.Once
)

func getMLE5StaticShapeExecutionStreamOperationPoolClass() MLE5StaticShapeExecutionStreamOperationPoolClass {
	_MLE5StaticShapeExecutionStreamOperationPoolClassOnce.Do(func() {
		_MLE5StaticShapeExecutionStreamOperationPoolClass = MLE5StaticShapeExecutionStreamOperationPoolClass{class: objc.GetClass("MLE5StaticShapeExecutionStreamOperationPool")}
	})
	return _MLE5StaticShapeExecutionStreamOperationPoolClass
}

// GetMLE5StaticShapeExecutionStreamOperationPoolClass returns the class object for MLE5StaticShapeExecutionStreamOperationPool.
func GetMLE5StaticShapeExecutionStreamOperationPoolClass() MLE5StaticShapeExecutionStreamOperationPoolClass {
	return getMLE5StaticShapeExecutionStreamOperationPoolClass()
}

type MLE5StaticShapeExecutionStreamOperationPoolClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLE5StaticShapeExecutionStreamOperationPoolClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLE5StaticShapeExecutionStreamOperationPoolClass) Alloc() MLE5StaticShapeExecutionStreamOperationPool {
	rv := objc.Send[MLE5StaticShapeExecutionStreamOperationPool](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLE5StaticShapeExecutionStreamOperationPool._putBack]
//   - [MLE5StaticShapeExecutionStreamOperationPool._takeOut]
//   - [MLE5StaticShapeExecutionStreamOperationPool.FunctionName]
//   - [MLE5StaticShapeExecutionStreamOperationPool.ModelConfiguration]
//   - [MLE5StaticShapeExecutionStreamOperationPool.ModelDescription]
//   - [MLE5StaticShapeExecutionStreamOperationPool.ModelSignpostId]
//   - [MLE5StaticShapeExecutionStreamOperationPool.NumberOfOperationsInUse]
//   - [MLE5StaticShapeExecutionStreamOperationPool.SetNumberOfOperationsInUse]
//   - [MLE5StaticShapeExecutionStreamOperationPool.PixelBufferPool]
//   - [MLE5StaticShapeExecutionStreamOperationPool.Pool]
//   - [MLE5StaticShapeExecutionStreamOperationPool.PrepareWithInitialPoolSizeError]
//   - [MLE5StaticShapeExecutionStreamOperationPool.ProgramLibrary]
//   - [MLE5StaticShapeExecutionStreamOperationPool.PutBack]
//   - [MLE5StaticShapeExecutionStreamOperationPool.SerialQueue]
//   - [MLE5StaticShapeExecutionStreamOperationPool.TakeOutOperationForFeaturesError]
//   - [MLE5StaticShapeExecutionStreamOperationPool.InitWithProgramLibraryFunctionNameModelDescriptionConfigurationModelSignpostIdCompilerVersionInfo]
//   - [MLE5StaticShapeExecutionStreamOperationPool.DebugDescription]
//   - [MLE5StaticShapeExecutionStreamOperationPool.Description]
//   - [MLE5StaticShapeExecutionStreamOperationPool.Hash]
//   - [MLE5StaticShapeExecutionStreamOperationPool.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLE5StaticShapeExecutionStreamOperationPool
type MLE5StaticShapeExecutionStreamOperationPool struct {
	objectivec.Object
}

// MLE5StaticShapeExecutionStreamOperationPoolFromID constructs a [MLE5StaticShapeExecutionStreamOperationPool] from an objc.ID.
func MLE5StaticShapeExecutionStreamOperationPoolFromID(id objc.ID) MLE5StaticShapeExecutionStreamOperationPool {
	return MLE5StaticShapeExecutionStreamOperationPool{objectivec.Object{ID: id}}
}

// Ensure MLE5StaticShapeExecutionStreamOperationPool implements IMLE5StaticShapeExecutionStreamOperationPool.
var _ IMLE5StaticShapeExecutionStreamOperationPool = MLE5StaticShapeExecutionStreamOperationPool{}

// An interface definition for the [MLE5StaticShapeExecutionStreamOperationPool] class.
//
// # Methods
//
//   - [IMLE5StaticShapeExecutionStreamOperationPool._putBack]
//   - [IMLE5StaticShapeExecutionStreamOperationPool._takeOut]
//   - [IMLE5StaticShapeExecutionStreamOperationPool.FunctionName]
//   - [IMLE5StaticShapeExecutionStreamOperationPool.ModelConfiguration]
//   - [IMLE5StaticShapeExecutionStreamOperationPool.ModelDescription]
//   - [IMLE5StaticShapeExecutionStreamOperationPool.ModelSignpostId]
//   - [IMLE5StaticShapeExecutionStreamOperationPool.NumberOfOperationsInUse]
//   - [IMLE5StaticShapeExecutionStreamOperationPool.SetNumberOfOperationsInUse]
//   - [IMLE5StaticShapeExecutionStreamOperationPool.PixelBufferPool]
//   - [IMLE5StaticShapeExecutionStreamOperationPool.Pool]
//   - [IMLE5StaticShapeExecutionStreamOperationPool.PrepareWithInitialPoolSizeError]
//   - [IMLE5StaticShapeExecutionStreamOperationPool.ProgramLibrary]
//   - [IMLE5StaticShapeExecutionStreamOperationPool.PutBack]
//   - [IMLE5StaticShapeExecutionStreamOperationPool.SerialQueue]
//   - [IMLE5StaticShapeExecutionStreamOperationPool.TakeOutOperationForFeaturesError]
//   - [IMLE5StaticShapeExecutionStreamOperationPool.InitWithProgramLibraryFunctionNameModelDescriptionConfigurationModelSignpostIdCompilerVersionInfo]
//   - [IMLE5StaticShapeExecutionStreamOperationPool.DebugDescription]
//   - [IMLE5StaticShapeExecutionStreamOperationPool.Description]
//   - [IMLE5StaticShapeExecutionStreamOperationPool.Hash]
//   - [IMLE5StaticShapeExecutionStreamOperationPool.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLE5StaticShapeExecutionStreamOperationPool
type IMLE5StaticShapeExecutionStreamOperationPool interface {
	objectivec.IObject

	// Topic: Methods

	_putBack(back objectivec.IObject)
	_takeOut() objectivec.IObject
	FunctionName() string
	ModelConfiguration() IMLModelConfiguration
	ModelDescription() IMLModelDescription
	ModelSignpostId() uint64
	NumberOfOperationsInUse() int64
	SetNumberOfOperationsInUse(value int64)
	PixelBufferPool() IMLPixelBufferPool
	Pool() foundation.INSSet
	PrepareWithInitialPoolSizeError(size int64) (bool, error)
	ProgramLibrary() IMLE5ProgramLibrary
	PutBack(back objectivec.IObject)
	SerialQueue() objectivec.Object
	TakeOutOperationForFeaturesError(features objectivec.IObject) (objectivec.IObject, error)
	InitWithProgramLibraryFunctionNameModelDescriptionConfigurationModelSignpostIdCompilerVersionInfo(library objectivec.IObject, name objectivec.IObject, description objectivec.IObject, configuration objectivec.IObject, id uint64, info objectivec.IObject) MLE5StaticShapeExecutionStreamOperationPool
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (e MLE5StaticShapeExecutionStreamOperationPool) Init() MLE5StaticShapeExecutionStreamOperationPool {
	rv := objc.Send[MLE5StaticShapeExecutionStreamOperationPool](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e MLE5StaticShapeExecutionStreamOperationPool) Autorelease() MLE5StaticShapeExecutionStreamOperationPool {
	rv := objc.Send[MLE5StaticShapeExecutionStreamOperationPool](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLE5StaticShapeExecutionStreamOperationPool creates a new MLE5StaticShapeExecutionStreamOperationPool instance.
func NewMLE5StaticShapeExecutionStreamOperationPool() MLE5StaticShapeExecutionStreamOperationPool {
	class := getMLE5StaticShapeExecutionStreamOperationPoolClass()
	rv := objc.Send[MLE5StaticShapeExecutionStreamOperationPool](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLE5StaticShapeExecutionStreamOperationPool/initWithProgramLibrary:functionName:modelDescription:configuration:modelSignpostId:compilerVersionInfo:
func NewE5StaticShapeExecutionStreamOperationPoolWithProgramLibraryFunctionNameModelDescriptionConfigurationModelSignpostIdCompilerVersionInfo(library objectivec.IObject, name objectivec.IObject, description objectivec.IObject, configuration objectivec.IObject, id uint64, info objectivec.IObject) MLE5StaticShapeExecutionStreamOperationPool {
	instance := getMLE5StaticShapeExecutionStreamOperationPoolClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithProgramLibrary:functionName:modelDescription:configuration:modelSignpostId:compilerVersionInfo:"), library, name, description, configuration, id, info)
	return MLE5StaticShapeExecutionStreamOperationPoolFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5StaticShapeExecutionStreamOperationPool/_putBack:
func (e MLE5StaticShapeExecutionStreamOperationPool) _putBack(back objectivec.IObject) {
	objc.Send[objc.ID](e.ID, objc.Sel("_putBack:"), back)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5StaticShapeExecutionStreamOperationPool/_takeOut
func (e MLE5StaticShapeExecutionStreamOperationPool) _takeOut() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("_takeOut"))
	return objectivec.Object{ID: rv}
}

// TakeOut is an exported wrapper for the private method _takeOut.
func (e MLE5StaticShapeExecutionStreamOperationPool) TakeOut() objectivec.IObject {
	return e._takeOut()
}

// See: https://developer.apple.com/documentation/CoreML/MLE5StaticShapeExecutionStreamOperationPool/prepareWithInitialPoolSize:error:
func (e MLE5StaticShapeExecutionStreamOperationPool) PrepareWithInitialPoolSizeError(size int64) (bool, error) {
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

// See: https://developer.apple.com/documentation/CoreML/MLE5StaticShapeExecutionStreamOperationPool/putBack:
func (e MLE5StaticShapeExecutionStreamOperationPool) PutBack(back objectivec.IObject) {
	objc.Send[objc.ID](e.ID, objc.Sel("putBack:"), back)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5StaticShapeExecutionStreamOperationPool/takeOutOperationForFeatures:error:
func (e MLE5StaticShapeExecutionStreamOperationPool) TakeOutOperationForFeaturesError(features objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](e.ID, objc.Sel("takeOutOperationForFeatures:error:"), features, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLE5StaticShapeExecutionStreamOperationPool/initWithProgramLibrary:functionName:modelDescription:configuration:modelSignpostId:compilerVersionInfo:
func (e MLE5StaticShapeExecutionStreamOperationPool) InitWithProgramLibraryFunctionNameModelDescriptionConfigurationModelSignpostIdCompilerVersionInfo(library objectivec.IObject, name objectivec.IObject, description objectivec.IObject, configuration objectivec.IObject, id uint64, info objectivec.IObject) MLE5StaticShapeExecutionStreamOperationPool {
	rv := objc.Send[MLE5StaticShapeExecutionStreamOperationPool](e.ID, objc.Sel("initWithProgramLibrary:functionName:modelDescription:configuration:modelSignpostId:compilerVersionInfo:"), library, name, description, configuration, id, info)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLE5StaticShapeExecutionStreamOperationPool/debugDescription
func (e MLE5StaticShapeExecutionStreamOperationPool) DebugDescription() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLE5StaticShapeExecutionStreamOperationPool/description
func (e MLE5StaticShapeExecutionStreamOperationPool) Description() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLE5StaticShapeExecutionStreamOperationPool/functionName
func (e MLE5StaticShapeExecutionStreamOperationPool) FunctionName() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("functionName"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLE5StaticShapeExecutionStreamOperationPool/hash
func (e MLE5StaticShapeExecutionStreamOperationPool) Hash() uint64 {
	rv := objc.Send[uint64](e.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLE5StaticShapeExecutionStreamOperationPool/modelConfiguration
func (e MLE5StaticShapeExecutionStreamOperationPool) ModelConfiguration() IMLModelConfiguration {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("modelConfiguration"))
	return MLModelConfigurationFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5StaticShapeExecutionStreamOperationPool/modelDescription
func (e MLE5StaticShapeExecutionStreamOperationPool) ModelDescription() IMLModelDescription {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("modelDescription"))
	return MLModelDescriptionFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5StaticShapeExecutionStreamOperationPool/modelSignpostId
func (e MLE5StaticShapeExecutionStreamOperationPool) ModelSignpostId() uint64 {
	rv := objc.Send[uint64](e.ID, objc.Sel("modelSignpostId"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLE5StaticShapeExecutionStreamOperationPool/numberOfOperationsInUse
func (e MLE5StaticShapeExecutionStreamOperationPool) NumberOfOperationsInUse() int64 {
	rv := objc.Send[int64](e.ID, objc.Sel("numberOfOperationsInUse"))
	return rv
}
func (e MLE5StaticShapeExecutionStreamOperationPool) SetNumberOfOperationsInUse(value int64) {
	objc.Send[struct{}](e.ID, objc.Sel("setNumberOfOperationsInUse:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLE5StaticShapeExecutionStreamOperationPool/pixelBufferPool
func (e MLE5StaticShapeExecutionStreamOperationPool) PixelBufferPool() IMLPixelBufferPool {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("pixelBufferPool"))
	return MLPixelBufferPoolFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5StaticShapeExecutionStreamOperationPool/pool
func (e MLE5StaticShapeExecutionStreamOperationPool) Pool() foundation.INSSet {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("pool"))
	return foundation.NSSetFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5StaticShapeExecutionStreamOperationPool/programLibrary
func (e MLE5StaticShapeExecutionStreamOperationPool) ProgramLibrary() IMLE5ProgramLibrary {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("programLibrary"))
	return MLE5ProgramLibraryFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5StaticShapeExecutionStreamOperationPool/serialQueue
func (e MLE5StaticShapeExecutionStreamOperationPool) SerialQueue() objectivec.Object {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("serialQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLE5StaticShapeExecutionStreamOperationPool/superclass
func (e MLE5StaticShapeExecutionStreamOperationPool) Superclass() objc.Class {
	rv := objc.Send[objc.Class](e.ID, objc.Sel("superclass"))
	return rv
}
