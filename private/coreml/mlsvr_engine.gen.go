// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLSVREngine] class.
var (
	_MLSVREngineClass     MLSVREngineClass
	_MLSVREngineClassOnce sync.Once
)

func getMLSVREngineClass() MLSVREngineClass {
	_MLSVREngineClassOnce.Do(func() {
		_MLSVREngineClass = MLSVREngineClass{class: objc.GetClass("MLSVREngine")}
	})
	return _MLSVREngineClass
}

// GetMLSVREngineClass returns the class object for MLSVREngine.
func GetMLSVREngineClass() MLSVREngineClass {
	return getMLSVREngineClass()
}

type MLSVREngineClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLSVREngineClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLSVREngineClass) Alloc() MLSVREngine {
	rv := objc.Send[MLSVREngine](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLSVREngine.AllocSVMNodeVector]
//   - [MLSVREngine.DeallocSVMNodeVector]
//   - [MLSVREngine.FillSVMNodeVectorValuesCount]
//   - [MLSVREngine.FreeModelOnDealloc]
//   - [MLSVREngine.SetFreeModelOnDealloc]
//   - [MLSVREngine.InputSize]
//   - [MLSVREngine.IsInputSizeLowerBoundOnly]
//   - [MLSVREngine.Model]
//   - [MLSVREngine.SetModel]
//   - [MLSVREngine.Predict]
//   - [MLSVREngine.InitWithLibSVMFile]
//   - [MLSVREngine.InitWithSVMModelFreeOnDeallocIsInputSizeLowerBoundOnlyInputSize]
//   - [MLSVREngine.InitWithSpecificationError]
//
// See: https://developer.apple.com/documentation/CoreML/MLSVREngine
type MLSVREngine struct {
	objectivec.Object
}

// MLSVREngineFromID constructs a [MLSVREngine] from an objc.ID.
func MLSVREngineFromID(id objc.ID) MLSVREngine {
	return MLSVREngine{objectivec.Object{ID: id}}
}

// Ensure MLSVREngine implements IMLSVREngine.
var _ IMLSVREngine = MLSVREngine{}

// An interface definition for the [MLSVREngine] class.
//
// # Methods
//
//   - [IMLSVREngine.AllocSVMNodeVector]
//   - [IMLSVREngine.DeallocSVMNodeVector]
//   - [IMLSVREngine.FillSVMNodeVectorValuesCount]
//   - [IMLSVREngine.FreeModelOnDealloc]
//   - [IMLSVREngine.SetFreeModelOnDealloc]
//   - [IMLSVREngine.InputSize]
//   - [IMLSVREngine.IsInputSizeLowerBoundOnly]
//   - [IMLSVREngine.Model]
//   - [IMLSVREngine.SetModel]
//   - [IMLSVREngine.Predict]
//   - [IMLSVREngine.InitWithLibSVMFile]
//   - [IMLSVREngine.InitWithSVMModelFreeOnDeallocIsInputSizeLowerBoundOnlyInputSize]
//   - [IMLSVREngine.InitWithSpecificationError]
//
// See: https://developer.apple.com/documentation/CoreML/MLSVREngine
type IMLSVREngine interface {
	objectivec.IObject

	// Topic: Methods

	AllocSVMNodeVector(vector uint64) objectivec.IObject
	DeallocSVMNodeVector(vector objectivec.IObject)
	FillSVMNodeVectorValuesCount(vector objectivec.IObject, values []float64, count uint64)
	FreeModelOnDealloc() bool
	SetFreeModelOnDealloc(value bool)
	InputSize() uint64
	IsInputSizeLowerBoundOnly() bool
	Model() objectivec.IObject
	SetModel(value objectivec.IObject)
	Predict(predict objectivec.IObject) objectivec.IObject
	InitWithLibSVMFile(sVMFile objectivec.IObject) MLSVREngine
	InitWithSVMModelFreeOnDeallocIsInputSizeLowerBoundOnlyInputSize(sVMModel objectivec.IObject, dealloc bool, only bool, size uint64) MLSVREngine
	InitWithSpecificationError(specification unsafe.Pointer) (MLSVREngine, error)
}

// Init initializes the instance.
func (s MLSVREngine) Init() MLSVREngine {
	rv := objc.Send[MLSVREngine](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s MLSVREngine) Autorelease() MLSVREngine {
	rv := objc.Send[MLSVREngine](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLSVREngine creates a new MLSVREngine instance.
func NewMLSVREngine() MLSVREngine {
	class := getMLSVREngineClass()
	rv := objc.Send[MLSVREngine](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLSVREngine/initWithLibSVMFile:
func NewSVREngineWithLibSVMFile(sVMFile objectivec.IObject) MLSVREngine {
	instance := getMLSVREngineClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLibSVMFile:"), sVMFile)
	return MLSVREngineFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLSVREngine/initWithSVMModel:freeOnDealloc:isInputSizeLowerBoundOnly:inputSize:
func NewSVREngineWithSVMModelFreeOnDeallocIsInputSizeLowerBoundOnlyInputSize(sVMModel objectivec.IObject, dealloc bool, only bool, size uint64) MLSVREngine {
	instance := getMLSVREngineClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSVMModel:freeOnDealloc:isInputSizeLowerBoundOnly:inputSize:"), sVMModel, dealloc, only, size)
	return MLSVREngineFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLSVREngine/initWithSpecification:error:
func NewSVREngineWithSpecificationError(specification unsafe.Pointer) (MLSVREngine, error) {
	var errorPtr objc.ID
	instance := getMLSVREngineClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpecification:error:"), specification, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLSVREngine{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLSVREngineFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLSVREngine/allocSVMNodeVector:
func (s MLSVREngine) AllocSVMNodeVector(vector uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("allocSVMNodeVector:"), vector)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLSVREngine/deallocSVMNodeVector:
func (s MLSVREngine) DeallocSVMNodeVector(vector objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("deallocSVMNodeVector:"), vector)
}

// See: https://developer.apple.com/documentation/CoreML/MLSVREngine/fillSVMNodeVector:values:count:
func (s MLSVREngine) FillSVMNodeVectorValuesCount(vector objectivec.IObject, values []float64, count uint64) {
	objc.Send[objc.ID](s.ID, objc.Sel("fillSVMNodeVector:values:count:"), vector, objc.CArray(values), count)
}

// See: https://developer.apple.com/documentation/CoreML/MLSVREngine/predict:
func (s MLSVREngine) Predict(predict objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("predict:"), predict)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLSVREngine/initWithLibSVMFile:
func (s MLSVREngine) InitWithLibSVMFile(sVMFile objectivec.IObject) MLSVREngine {
	rv := objc.Send[MLSVREngine](s.ID, objc.Sel("initWithLibSVMFile:"), sVMFile)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLSVREngine/initWithSVMModel:freeOnDealloc:isInputSizeLowerBoundOnly:inputSize:
func (s MLSVREngine) InitWithSVMModelFreeOnDeallocIsInputSizeLowerBoundOnlyInputSize(sVMModel objectivec.IObject, dealloc bool, only bool, size uint64) MLSVREngine {
	rv := objc.Send[MLSVREngine](s.ID, objc.Sel("initWithSVMModel:freeOnDealloc:isInputSizeLowerBoundOnly:inputSize:"), sVMModel, dealloc, only, size)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLSVREngine/initWithSpecification:error:
func (s MLSVREngine) InitWithSpecificationError(specification unsafe.Pointer) (MLSVREngine, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](s.ID, objc.Sel("initWithSpecification:error:"), specification, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLSVREngine{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLSVREngineFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/MLSVREngine/freeModelOnDealloc
func (s MLSVREngine) FreeModelOnDealloc() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("freeModelOnDealloc"))
	return rv
}
func (s MLSVREngine) SetFreeModelOnDealloc(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setFreeModelOnDealloc:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLSVREngine/inputSize
func (s MLSVREngine) InputSize() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("inputSize"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLSVREngine/isInputSizeLowerBoundOnly
func (s MLSVREngine) IsInputSizeLowerBoundOnly() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isInputSizeLowerBoundOnly"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLSVREngine/model
func (s MLSVREngine) Model() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("model"))
	return objectivec.Object{ID: rv}
}
func (s MLSVREngine) SetModel(value objectivec.IObject) {
	objc.Send[struct{}](s.ID, objc.Sel("setModel:"), value)
}
