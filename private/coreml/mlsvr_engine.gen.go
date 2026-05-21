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
type IMLSVREngine interface {
	objectivec.IObject

	// Topic: Methods

	AllocSVMNodeVector(vector uint64) unsafe.Pointer
	DeallocSVMNodeVector(vector *SvmNode)
	FillSVMNodeVectorValuesCount(vector *SvmNode, values []float64, count uint64)
	FreeModelOnDealloc() bool
	SetFreeModelOnDealloc(value bool)
	InputSize() uint64
	IsInputSizeLowerBoundOnly() bool
	Model() unsafe.Pointer
	SetModel(value unsafe.Pointer)
	Predict(predict objectivec.IObject) objectivec.IObject
	InitWithLibSVMFile(sVMFile objectivec.IObject) MLSVREngine
	InitWithSVMModelFreeOnDeallocIsInputSizeLowerBoundOnlyInputSize(sVMModel *SvmModel, dealloc bool, only bool, size uint64) MLSVREngine
	InitWithSpecificationError(specification unsafe.Pointer) (MLSVREngine, error)
}

// Init initializes the instance.
func (m MLSVREngine) Init() MLSVREngine {
	rv := objc.Send[MLSVREngine](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLSVREngine) Autorelease() MLSVREngine {
	rv := objc.Send[MLSVREngine](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLSVREngine creates a new MLSVREngine instance.
func NewMLSVREngine() MLSVREngine {
	class := getMLSVREngineClass()
	rv := objc.Send[MLSVREngine](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSVREngineWithLibSVMFile(sVMFile objectivec.IObject) MLSVREngine {
	instance := getMLSVREngineClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLibSVMFile:"), sVMFile)
	return MLSVREngineFromID(rv)
}

func NewSVREngineWithSVMModelFreeOnDeallocIsInputSizeLowerBoundOnlyInputSize(sVMModel *SvmModel, dealloc bool, only bool, size uint64) MLSVREngine {
	instance := getMLSVREngineClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSVMModel:freeOnDealloc:isInputSizeLowerBoundOnly:inputSize:"), sVMModel, dealloc, only, size)
	return MLSVREngineFromID(rv)
}

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

func (m MLSVREngine) AllocSVMNodeVector(vector uint64) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m.ID, objc.Sel("allocSVMNodeVector:"), vector)
	return rv
}
func (m MLSVREngine) DeallocSVMNodeVector(vector *SvmNode) {
	objc.Send[objc.ID](m.ID, objc.Sel("deallocSVMNodeVector:"), vector)
}
func (m MLSVREngine) FillSVMNodeVectorValuesCount(vector *SvmNode, values []float64, count uint64) {
	objc.Send[objc.ID](m.ID, objc.Sel("fillSVMNodeVector:values:count:"), objc.CArray(vector), objc.CArray(values), count)
}
func (m MLSVREngine) Predict(predict objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predict:"), predict)
	return objectivec.Object{ID: rv}
}
func (m MLSVREngine) InitWithLibSVMFile(sVMFile objectivec.IObject) MLSVREngine {
	rv := objc.Send[MLSVREngine](m.ID, objc.Sel("initWithLibSVMFile:"), sVMFile)
	return rv
}
func (m MLSVREngine) InitWithSVMModelFreeOnDeallocIsInputSizeLowerBoundOnlyInputSize(sVMModel *SvmModel, dealloc bool, only bool, size uint64) MLSVREngine {
	rv := objc.Send[MLSVREngine](m.ID, objc.Sel("initWithSVMModel:freeOnDealloc:isInputSizeLowerBoundOnly:inputSize:"), sVMModel, dealloc, only, size)
	return rv
}
func (m MLSVREngine) InitWithSpecificationError(specification unsafe.Pointer) (MLSVREngine, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithSpecification:error:"), specification, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLSVREngine{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLSVREngineFromID(rv), nil

}

func (m MLSVREngine) FreeModelOnDealloc() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("freeModelOnDealloc"))
	return rv
}
func (m MLSVREngine) SetFreeModelOnDealloc(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setFreeModelOnDealloc:"), value)
}
func (m MLSVREngine) InputSize() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("inputSize"))
	return rv
}
func (m MLSVREngine) IsInputSizeLowerBoundOnly() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isInputSizeLowerBoundOnly"))
	return rv
}
func (m MLSVREngine) Model() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m.ID, objc.Sel("model"))
	return rv
}
func (m MLSVREngine) SetModel(value unsafe.Pointer) {
	objc.Send[struct{}](m.ID, objc.Sel("setModel:"), value)
}
