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
	rv := objc.SendIfResponds[MLComputeBatchDataSource](objc.ID(mc.class), objc.Sel("alloc"))
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
type IMLComputeBatchDataSource interface {
	objectivec.IObject

	// Topic: Methods

	BatchAtIndexError(index uint64) (objectivec.IObject, error)
	BatchProvider() unsafe.Pointer
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
func (m MLComputeBatchDataSource) Init() MLComputeBatchDataSource {
	rv := objc.SendIfResponds[MLComputeBatchDataSource](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLComputeBatchDataSource) Autorelease() MLComputeBatchDataSource {
	rv := objc.SendIfResponds[MLComputeBatchDataSource](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLComputeBatchDataSource creates a new MLComputeBatchDataSource instance.
func NewMLComputeBatchDataSource() MLComputeBatchDataSource {
	class := getMLComputeBatchDataSourceClass()
	rv := objc.SendIfResponds[MLComputeBatchDataSource](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewComputeBatchDataSourceWithBatchProviderBatchSizeForPredictionNeuralNetworkEngineError(provider objectivec.IObject, size uint64, prediction bool, engine objectivec.IObject) (MLComputeBatchDataSource, error) {
	var errorPtr objc.ID
	instance := getMLComputeBatchDataSourceClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithBatchProvider:batchSize:forPrediction:neuralNetworkEngine:error:"), provider, size, prediction, engine, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLComputeBatchDataSource{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return MLComputeBatchDataSource{}, objc.ErrInitFailed
	}
	return MLComputeBatchDataSourceFromID(rv), nil
}

func (m MLComputeBatchDataSource) BatchAtIndexError(index uint64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("batchAtIndex:error:"), index, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLComputeBatchDataSource) MlcDataSourceAtIndexError(index int64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("mlcDataSourceAtIndex:error:"), index, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLComputeBatchDataSource) NumberOfBatches() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("numberOfBatches"))
	return rv
}
func (m MLComputeBatchDataSource) SizeOfBatchAtIndex(index uint64) uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("sizeOfBatchAtIndex:"), index)
	return rv
}
func (m MLComputeBatchDataSource) InitWithBatchProviderBatchSizeForPredictionNeuralNetworkEngineError(provider objectivec.IObject, size uint64, prediction bool, engine objectivec.IObject) (MLComputeBatchDataSource, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithBatchProvider:batchSize:forPrediction:neuralNetworkEngine:error:"), provider, size, prediction, engine, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLComputeBatchDataSource{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLComputeBatchDataSourceFromID(rv), nil

}

func (m MLComputeBatchDataSource) BatchProvider() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](m.ID, objc.Sel("batchProvider"))
	return rv
}
func (m MLComputeBatchDataSource) BatchSize() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("batchSize"))
	return rv
}
func (m MLComputeBatchDataSource) SetBatchSize(value uint64) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setBatchSize:"), value)
}
func (m MLComputeBatchDataSource) NnEngine() IMLNeuralNetworkEngine {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("nnEngine"))
	return MLNeuralNetworkEngineFromID(objc.ID(rv))
}
func (m MLComputeBatchDataSource) UseForPrediction() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("useForPrediction"))
	return rv
}
