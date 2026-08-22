// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLSVMEngine] class.
var (
	_MLSVMEngineClass     MLSVMEngineClass
	_MLSVMEngineClassOnce sync.Once
)

func getMLSVMEngineClass() MLSVMEngineClass {
	_MLSVMEngineClassOnce.Do(func() {
		_MLSVMEngineClass = MLSVMEngineClass{class: objc.GetClass("MLSVMEngine")}
	})
	return _MLSVMEngineClass
}

// GetMLSVMEngineClass returns the class object for MLSVMEngine.
func GetMLSVMEngineClass() MLSVMEngineClass {
	return getMLSVMEngineClass()
}

type MLSVMEngineClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLSVMEngineClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLSVMEngineClass) Alloc() MLSVMEngine {
	rv := objc.SendIfResponds[MLSVMEngine](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLSVMEngine.AllocSVMNodeVector]
//   - [MLSVMEngine.ClassLabels]
//   - [MLSVMEngine.SetClassLabels]
//   - [MLSVMEngine.DeallocSVMNodeVector]
//   - [MLSVMEngine.FillSVMNodeVectorValuesCount]
//   - [MLSVMEngine.FreeModelOnDealloc]
//   - [MLSVMEngine.SetFreeModelOnDealloc]
//   - [MLSVMEngine.HasProbabilityPredictionEnabled]
//   - [MLSVMEngine.InputSize]
//   - [MLSVMEngine.SetInputSize]
//   - [MLSVMEngine.IsInputSizeLowerBoundOnly]
//   - [MLSVMEngine.SetIsInputSizeLowerBoundOnly]
//   - [MLSVMEngine.Model]
//   - [MLSVMEngine.SetModel]
//   - [MLSVMEngine.NumberOfClasses]
//   - [MLSVMEngine.Predict]
//   - [MLSVMEngine.PredictProbabilitiesProbabilities]
//   - [MLSVMEngine.InitWithLibSVMFileClassLabels]
//   - [MLSVMEngine.InitWithSVMModelFreeOnDeallocIsInputSizeLowerBoundOnlyInputSizeClassLabels]
//   - [MLSVMEngine.InitWithSpecificationError]
type MLSVMEngine struct {
	objectivec.Object
}

// MLSVMEngineFromID constructs a [MLSVMEngine] from an objc.ID.
func MLSVMEngineFromID(id objc.ID) MLSVMEngine {
	return MLSVMEngine{objectivec.Object{ID: id}}
}

// Ensure MLSVMEngine implements IMLSVMEngine.
var _ IMLSVMEngine = MLSVMEngine{}

// An interface definition for the [MLSVMEngine] class.
//
// # Methods
//
//   - [IMLSVMEngine.AllocSVMNodeVector]
//   - [IMLSVMEngine.ClassLabels]
//   - [IMLSVMEngine.SetClassLabels]
//   - [IMLSVMEngine.DeallocSVMNodeVector]
//   - [IMLSVMEngine.FillSVMNodeVectorValuesCount]
//   - [IMLSVMEngine.FreeModelOnDealloc]
//   - [IMLSVMEngine.SetFreeModelOnDealloc]
//   - [IMLSVMEngine.HasProbabilityPredictionEnabled]
//   - [IMLSVMEngine.InputSize]
//   - [IMLSVMEngine.SetInputSize]
//   - [IMLSVMEngine.IsInputSizeLowerBoundOnly]
//   - [IMLSVMEngine.SetIsInputSizeLowerBoundOnly]
//   - [IMLSVMEngine.Model]
//   - [IMLSVMEngine.SetModel]
//   - [IMLSVMEngine.NumberOfClasses]
//   - [IMLSVMEngine.Predict]
//   - [IMLSVMEngine.PredictProbabilitiesProbabilities]
//   - [IMLSVMEngine.InitWithLibSVMFileClassLabels]
//   - [IMLSVMEngine.InitWithSVMModelFreeOnDeallocIsInputSizeLowerBoundOnlyInputSizeClassLabels]
//   - [IMLSVMEngine.InitWithSpecificationError]
type IMLSVMEngine interface {
	objectivec.IObject

	// Topic: Methods

	AllocSVMNodeVector(vector uint64) unsafe.Pointer
	ClassLabels() foundation.INSArray
	SetClassLabels(value foundation.INSArray)
	DeallocSVMNodeVector(vector *SvmNode)
	FillSVMNodeVectorValuesCount(vector *SvmNode, values []float64, count uint64)
	FreeModelOnDealloc() bool
	SetFreeModelOnDealloc(value bool)
	HasProbabilityPredictionEnabled() bool
	InputSize() uint64
	SetInputSize(value uint64)
	IsInputSizeLowerBoundOnly() bool
	SetIsInputSizeLowerBoundOnly(value bool)
	Model() unsafe.Pointer
	SetModel(value unsafe.Pointer)
	NumberOfClasses() uint64
	Predict(predict objectivec.IObject) objectivec.IObject
	PredictProbabilitiesProbabilities(probabilities objectivec.IObject, probabilities2 []float64)
	InitWithLibSVMFileClassLabels(sVMFile objectivec.IObject, labels objectivec.IObject) MLSVMEngine
	InitWithSVMModelFreeOnDeallocIsInputSizeLowerBoundOnlyInputSizeClassLabels(sVMModel *SvmModel, dealloc bool, only bool, size uint64, labels objectivec.IObject) MLSVMEngine
	InitWithSpecificationError(specification unsafe.Pointer) (MLSVMEngine, error)
}

// Init initializes the instance.
func (m MLSVMEngine) Init() MLSVMEngine {
	rv := objc.SendIfResponds[MLSVMEngine](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLSVMEngine) Autorelease() MLSVMEngine {
	rv := objc.SendIfResponds[MLSVMEngine](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLSVMEngine creates a new MLSVMEngine instance.
func NewMLSVMEngine() MLSVMEngine {
	class := getMLSVMEngineClass()
	rv := objc.SendIfResponds[MLSVMEngine](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSVMEngineWithLibSVMFileClassLabels(sVMFile objectivec.IObject, labels objectivec.IObject) MLSVMEngine {
	instance := getMLSVMEngineClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithLibSVMFile:classLabels:"), sVMFile, labels)
	return MLSVMEngineFromID(rv)
}

func NewSVMEngineWithSVMModelFreeOnDeallocIsInputSizeLowerBoundOnlyInputSizeClassLabels(sVMModel *SvmModel, dealloc bool, only bool, size uint64, labels objectivec.IObject) MLSVMEngine {
	instance := getMLSVMEngineClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithSVMModel:freeOnDealloc:isInputSizeLowerBoundOnly:inputSize:classLabels:"), sVMModel, dealloc, only, size, labels)
	return MLSVMEngineFromID(rv)
}

func NewSVMEngineWithSpecificationError(specification unsafe.Pointer) (MLSVMEngine, error) {
	var errorPtr objc.ID
	instance := getMLSVMEngineClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithSpecification:error:"), specification, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLSVMEngine{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return MLSVMEngine{}, objc.ErrInitFailed
	}
	return MLSVMEngineFromID(rv), nil
}

func (m MLSVMEngine) AllocSVMNodeVector(vector uint64) unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](m.ID, objc.Sel("allocSVMNodeVector:"), vector)
	return rv
}
func (m MLSVMEngine) DeallocSVMNodeVector(vector *SvmNode) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("deallocSVMNodeVector:"), vector)
}
func (m MLSVMEngine) FillSVMNodeVectorValuesCount(vector *SvmNode, values []float64, count uint64) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("fillSVMNodeVector:values:count:"), objc.CArray(vector), objc.CArray(values), count)
}
func (m MLSVMEngine) HasProbabilityPredictionEnabled() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("hasProbabilityPredictionEnabled"))
	return rv
}
func (m MLSVMEngine) Predict(predict objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("predict:"), predict)
	return objectivec.Object{ID: rv}
}
func (m MLSVMEngine) PredictProbabilitiesProbabilities(probabilities objectivec.IObject, probabilities2 []float64) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("predictProbabilities:probabilities:"), probabilities, probabilities2)
}
func (m MLSVMEngine) InitWithLibSVMFileClassLabels(sVMFile objectivec.IObject, labels objectivec.IObject) MLSVMEngine {
	rv := objc.SendIfResponds[MLSVMEngine](m.ID, objc.Sel("initWithLibSVMFile:classLabels:"), sVMFile, labels)
	return rv
}
func (m MLSVMEngine) InitWithSVMModelFreeOnDeallocIsInputSizeLowerBoundOnlyInputSizeClassLabels(sVMModel *SvmModel, dealloc bool, only bool, size uint64, labels objectivec.IObject) MLSVMEngine {
	rv := objc.SendIfResponds[MLSVMEngine](m.ID, objc.Sel("initWithSVMModel:freeOnDealloc:isInputSizeLowerBoundOnly:inputSize:classLabels:"), sVMModel, dealloc, only, size, labels)
	return rv
}
func (m MLSVMEngine) InitWithSpecificationError(specification unsafe.Pointer) (MLSVMEngine, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithSpecification:error:"), specification, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLSVMEngine{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLSVMEngineFromID(rv), nil

}

func (m MLSVMEngine) ClassLabels() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("classLabels"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLSVMEngine) SetClassLabels(value foundation.INSArray) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setClassLabels:"), value)
}
func (m MLSVMEngine) FreeModelOnDealloc() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("freeModelOnDealloc"))
	return rv
}
func (m MLSVMEngine) SetFreeModelOnDealloc(value bool) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setFreeModelOnDealloc:"), value)
}
func (m MLSVMEngine) InputSize() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("inputSize"))
	return rv
}
func (m MLSVMEngine) SetInputSize(value uint64) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setInputSize:"), value)
}
func (m MLSVMEngine) IsInputSizeLowerBoundOnly() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("isInputSizeLowerBoundOnly"))
	return rv
}
func (m MLSVMEngine) SetIsInputSizeLowerBoundOnly(value bool) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setIsInputSizeLowerBoundOnly:"), value)
}
func (m MLSVMEngine) Model() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](m.ID, objc.Sel("model"))
	return rv
}
func (m MLSVMEngine) SetModel(value unsafe.Pointer) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setModel:"), value)
}
func (m MLSVMEngine) NumberOfClasses() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("numberOfClasses"))
	return rv
}
