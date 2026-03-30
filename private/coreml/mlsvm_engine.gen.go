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
	rv := objc.Send[MLSVMEngine](objc.ID(mc.class), objc.Sel("alloc"))
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
//
// See: https://developer.apple.com/documentation/CoreML/MLSVMEngine
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
//
// See: https://developer.apple.com/documentation/CoreML/MLSVMEngine
type IMLSVMEngine interface {
	objectivec.IObject

	// Topic: Methods

	AllocSVMNodeVector(vector uint64) objectivec.IObject
	ClassLabels() foundation.INSArray
	SetClassLabels(value foundation.INSArray)
	DeallocSVMNodeVector(vector objectivec.IObject)
	FillSVMNodeVectorValuesCount(vector objectivec.IObject, values []float64, count uint64)
	FreeModelOnDealloc() bool
	SetFreeModelOnDealloc(value bool)
	HasProbabilityPredictionEnabled() bool
	InputSize() uint64
	SetInputSize(value uint64)
	IsInputSizeLowerBoundOnly() bool
	SetIsInputSizeLowerBoundOnly(value bool)
	Model() objectivec.IObject
	SetModel(value objectivec.IObject)
	NumberOfClasses() uint64
	Predict(predict objectivec.IObject) objectivec.IObject
	PredictProbabilitiesProbabilities(probabilities objectivec.IObject, probabilities2 []float64)
	InitWithLibSVMFileClassLabels(sVMFile objectivec.IObject, labels objectivec.IObject) MLSVMEngine
	InitWithSVMModelFreeOnDeallocIsInputSizeLowerBoundOnlyInputSizeClassLabels(sVMModel objectivec.IObject, dealloc bool, only bool, size uint64, labels objectivec.IObject) MLSVMEngine
	InitWithSpecificationError(specification unsafe.Pointer) (MLSVMEngine, error)
}

// Init initializes the instance.
func (s MLSVMEngine) Init() MLSVMEngine {
	rv := objc.Send[MLSVMEngine](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s MLSVMEngine) Autorelease() MLSVMEngine {
	rv := objc.Send[MLSVMEngine](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLSVMEngine creates a new MLSVMEngine instance.
func NewMLSVMEngine() MLSVMEngine {
	class := getMLSVMEngineClass()
	rv := objc.Send[MLSVMEngine](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLSVMEngine/initWithLibSVMFile:classLabels:
func NewSVMEngineWithLibSVMFileClassLabels(sVMFile objectivec.IObject, labels objectivec.IObject) MLSVMEngine {
	instance := getMLSVMEngineClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLibSVMFile:classLabels:"), sVMFile, labels)
	return MLSVMEngineFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLSVMEngine/initWithSVMModel:freeOnDealloc:isInputSizeLowerBoundOnly:inputSize:classLabels:
func NewSVMEngineWithSVMModelFreeOnDeallocIsInputSizeLowerBoundOnlyInputSizeClassLabels(sVMModel objectivec.IObject, dealloc bool, only bool, size uint64, labels objectivec.IObject) MLSVMEngine {
	instance := getMLSVMEngineClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSVMModel:freeOnDealloc:isInputSizeLowerBoundOnly:inputSize:classLabels:"), sVMModel, dealloc, only, size, labels)
	return MLSVMEngineFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLSVMEngine/initWithSpecification:error:
func NewSVMEngineWithSpecificationError(specification unsafe.Pointer) (MLSVMEngine, error) {
	var errorPtr objc.ID
	instance := getMLSVMEngineClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpecification:error:"), specification, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLSVMEngine{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLSVMEngineFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLSVMEngine/allocSVMNodeVector:
func (s MLSVMEngine) AllocSVMNodeVector(vector uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("allocSVMNodeVector:"), vector)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLSVMEngine/deallocSVMNodeVector:
func (s MLSVMEngine) DeallocSVMNodeVector(vector objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("deallocSVMNodeVector:"), vector)
}

// See: https://developer.apple.com/documentation/CoreML/MLSVMEngine/fillSVMNodeVector:values:count:
func (s MLSVMEngine) FillSVMNodeVectorValuesCount(vector objectivec.IObject, values []float64, count uint64) {
	objc.Send[objc.ID](s.ID, objc.Sel("fillSVMNodeVector:values:count:"), vector, objc.CArray(values), count)
}

// See: https://developer.apple.com/documentation/CoreML/MLSVMEngine/hasProbabilityPredictionEnabled
func (s MLSVMEngine) HasProbabilityPredictionEnabled() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("hasProbabilityPredictionEnabled"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLSVMEngine/predict:
func (s MLSVMEngine) Predict(predict objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("predict:"), predict)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLSVMEngine/predictProbabilities:probabilities:
func (s MLSVMEngine) PredictProbabilitiesProbabilities(probabilities objectivec.IObject, probabilities2 []float64) {
	objc.Send[objc.ID](s.ID, objc.Sel("predictProbabilities:probabilities:"), probabilities, probabilities2)
}

// See: https://developer.apple.com/documentation/CoreML/MLSVMEngine/initWithLibSVMFile:classLabels:
func (s MLSVMEngine) InitWithLibSVMFileClassLabels(sVMFile objectivec.IObject, labels objectivec.IObject) MLSVMEngine {
	rv := objc.Send[MLSVMEngine](s.ID, objc.Sel("initWithLibSVMFile:classLabels:"), sVMFile, labels)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLSVMEngine/initWithSVMModel:freeOnDealloc:isInputSizeLowerBoundOnly:inputSize:classLabels:
func (s MLSVMEngine) InitWithSVMModelFreeOnDeallocIsInputSizeLowerBoundOnlyInputSizeClassLabels(sVMModel objectivec.IObject, dealloc bool, only bool, size uint64, labels objectivec.IObject) MLSVMEngine {
	rv := objc.Send[MLSVMEngine](s.ID, objc.Sel("initWithSVMModel:freeOnDealloc:isInputSizeLowerBoundOnly:inputSize:classLabels:"), sVMModel, dealloc, only, size, labels)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLSVMEngine/initWithSpecification:error:
func (s MLSVMEngine) InitWithSpecificationError(specification unsafe.Pointer) (MLSVMEngine, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](s.ID, objc.Sel("initWithSpecification:error:"), specification, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLSVMEngine{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLSVMEngineFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/MLSVMEngine/classLabels
func (s MLSVMEngine) ClassLabels() foundation.INSArray {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("classLabels"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (s MLSVMEngine) SetClassLabels(value foundation.INSArray) {
	objc.Send[struct{}](s.ID, objc.Sel("setClassLabels:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLSVMEngine/freeModelOnDealloc
func (s MLSVMEngine) FreeModelOnDealloc() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("freeModelOnDealloc"))
	return rv
}
func (s MLSVMEngine) SetFreeModelOnDealloc(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setFreeModelOnDealloc:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLSVMEngine/inputSize
func (s MLSVMEngine) InputSize() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("inputSize"))
	return rv
}
func (s MLSVMEngine) SetInputSize(value uint64) {
	objc.Send[struct{}](s.ID, objc.Sel("setInputSize:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLSVMEngine/isInputSizeLowerBoundOnly
func (s MLSVMEngine) IsInputSizeLowerBoundOnly() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isInputSizeLowerBoundOnly"))
	return rv
}
func (s MLSVMEngine) SetIsInputSizeLowerBoundOnly(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setIsInputSizeLowerBoundOnly:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLSVMEngine/model
func (s MLSVMEngine) Model() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("model"))
	return objectivec.Object{ID: rv}
}
func (s MLSVMEngine) SetModel(value objectivec.IObject) {
	objc.Send[struct{}](s.ID, objc.Sel("setModel:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLSVMEngine/numberOfClasses
func (s MLSVMEngine) NumberOfClasses() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("numberOfClasses"))
	return rv
}
