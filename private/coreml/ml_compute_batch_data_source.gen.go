// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLComputeBatchDataSource] class.
var (
	_MLComputeBatchDataSourceClass     MLComputeBatchDataSourceClass
	_MLComputeBatchDataSourceClassOnce sync.Once
)

func getMLComputeBatchDataSourceClass() MLComputeBatchDataSourceClass {
	_MLComputeBatchDataSourceClassOnce.Do(func() {
		_MLComputeBatchDataSourceClass = MLComputeBatchDataSourceClass{class: objc.GetClass("MLComputeBatchDataSource")}
	})
	return _MLComputeBatchDataSourceClass
}

// GetMLComputeBatchDataSourceClass returns the class object for MLComputeBatchDataSource.
func GetMLComputeBatchDataSourceClass() MLComputeBatchDataSourceClass {
	return getMLComputeBatchDataSourceClass()
}

type MLComputeBatchDataSourceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLComputeBatchDataSourceClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLComputeBatchDataSourceClass) Alloc() MLComputeBatchDataSource {
	rv := objc.Send[MLComputeBatchDataSource](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLComputeBatchDataSource.BatchAtIndexError]
//   - [MLComputeBatchDataSource.BatchProvider]
//   - [MLComputeBatchDataSource.BatchSize]
//   - [MLComputeBatchDataSource.SetBatchSize]
//   - [MLComputeBatchDataSource.MlcDataSourceAtIndexError]
//   - [MLComputeBatchDataSource.NnEngine]
//   - [MLComputeBatchDataSource.NumberOfBatches]
//   - [MLComputeBatchDataSource.SizeOfBatchAtIndex]
//   - [MLComputeBatchDataSource.UseForPrediction]
//   - [MLComputeBatchDataSource.InitWithBatchProviderBatchSizeForPredictionNeuralNetworkEngineError]
//
// See: https://developer.apple.com/documentation/CoreML/MLComputeBatchDataSource
type MLComputeBatchDataSource struct {
	objectivec.Object
}

// MLComputeBatchDataSourceFromID constructs a [MLComputeBatchDataSource] from an objc.ID.
func MLComputeBatchDataSourceFromID(id objc.ID) MLComputeBatchDataSource {
	return MLComputeBatchDataSource{objectivec.Object{ID: id}}
}

// Ensure MLComputeBatchDataSource implements IMLComputeBatchDataSource.
var _ IMLComputeBatchDataSource = MLComputeBatchDataSource{}

// An interface definition for the [MLComputeBatchDataSource] class.
//
// # Methods
//
//   - [IMLComputeBatchDataSource.BatchAtIndexError]
//   - [IMLComputeBatchDataSource.BatchProvider]
//   - [IMLComputeBatchDataSource.BatchSize]
//   - [IMLComputeBatchDataSource.SetBatchSize]
//   - [IMLComputeBatchDataSource.MlcDataSourceAtIndexError]
//   - [IMLComputeBatchDataSource.NnEngine]
//   - [IMLComputeBatchDataSource.NumberOfBatches]
//   - [IMLComputeBatchDataSource.SizeOfBatchAtIndex]
//   - [IMLComputeBatchDataSource.UseForPrediction]
//   - [IMLComputeBatchDataSource.InitWithBatchProviderBatchSizeForPredictionNeuralNetworkEngineError]
//
// See: https://developer.apple.com/documentation/CoreML/MLComputeBatchDataSource
type IMLComputeBatchDataSource interface {
	objectivec.IObject

	// Topic: Methods

	BatchAtIndexError(index uint64) (objectivec.IObject, error)
	BatchProvider() objectivec.IObject
	BatchSize() uint64
	SetBatchSize(value uint64)
	MlcDataSourceAtIndexError(index int64) (objectivec.IObject, error)
	NnEngine() IMLNeuralNetworkEngine
	NumberOfBatches() uint64
	SizeOfBatchAtIndex(index uint64) uint64
	UseForPrediction() bool
	InitWithBatchProviderBatchSizeForPredictionNeuralNetworkEngineError(provider objectivec.IObject, size uint64, prediction bool, engine objectivec.IObject) (MLComputeBatchDataSource, error)
}

// Init initializes the instance.
func (c MLComputeBatchDataSource) Init() MLComputeBatchDataSource {
	rv := objc.Send[MLComputeBatchDataSource](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MLComputeBatchDataSource) Autorelease() MLComputeBatchDataSource {
	rv := objc.Send[MLComputeBatchDataSource](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLComputeBatchDataSource creates a new MLComputeBatchDataSource instance.
func NewMLComputeBatchDataSource() MLComputeBatchDataSource {
	class := getMLComputeBatchDataSourceClass()
	rv := objc.Send[MLComputeBatchDataSource](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLComputeBatchDataSource/initWithBatchProvider:batchSize:forPrediction:neuralNetworkEngine:error:
func NewComputeBatchDataSourceWithBatchProviderBatchSizeForPredictionNeuralNetworkEngineError(provider objectivec.IObject, size uint64, prediction bool, engine objectivec.IObject) (MLComputeBatchDataSource, error) {
	var errorPtr objc.ID
	instance := getMLComputeBatchDataSourceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBatchProvider:batchSize:forPrediction:neuralNetworkEngine:error:"), provider, size, prediction, engine, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLComputeBatchDataSource{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLComputeBatchDataSourceFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLComputeBatchDataSource/batchAtIndex:error:
func (c MLComputeBatchDataSource) BatchAtIndexError(index uint64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](c.ID, objc.Sel("batchAtIndex:error:"), index, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLComputeBatchDataSource/mlcDataSourceAtIndex:error:
func (c MLComputeBatchDataSource) MlcDataSourceAtIndexError(index int64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](c.ID, objc.Sel("mlcDataSourceAtIndex:error:"), index, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLComputeBatchDataSource/numberOfBatches
func (c MLComputeBatchDataSource) NumberOfBatches() uint64 {
	rv := objc.Send[uint64](c.ID, objc.Sel("numberOfBatches"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLComputeBatchDataSource/sizeOfBatchAtIndex:
func (c MLComputeBatchDataSource) SizeOfBatchAtIndex(index uint64) uint64 {
	rv := objc.Send[uint64](c.ID, objc.Sel("sizeOfBatchAtIndex:"), index)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLComputeBatchDataSource/initWithBatchProvider:batchSize:forPrediction:neuralNetworkEngine:error:
func (c MLComputeBatchDataSource) InitWithBatchProviderBatchSizeForPredictionNeuralNetworkEngineError(provider objectivec.IObject, size uint64, prediction bool, engine objectivec.IObject) (MLComputeBatchDataSource, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](c.ID, objc.Sel("initWithBatchProvider:batchSize:forPrediction:neuralNetworkEngine:error:"), provider, size, prediction, engine, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLComputeBatchDataSource{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLComputeBatchDataSourceFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/MLComputeBatchDataSource/batchProvider
func (c MLComputeBatchDataSource) BatchProvider() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("batchProvider"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLComputeBatchDataSource/batchSize
func (c MLComputeBatchDataSource) BatchSize() uint64 {
	rv := objc.Send[uint64](c.ID, objc.Sel("batchSize"))
	return rv
}
func (c MLComputeBatchDataSource) SetBatchSize(value uint64) {
	objc.Send[struct{}](c.ID, objc.Sel("setBatchSize:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLComputeBatchDataSource/nnEngine
func (c MLComputeBatchDataSource) NnEngine() IMLNeuralNetworkEngine {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("nnEngine"))
	return MLNeuralNetworkEngineFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLComputeBatchDataSource/useForPrediction
func (c MLComputeBatchDataSource) UseForPrediction() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("useForPrediction"))
	return rv
}
