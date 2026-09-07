// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLBatchDataSource] class.
var (
	_MLBatchDataSourceClass     MLBatchDataSourceClass
	_MLBatchDataSourceClassOnce sync.Once
)

func getMLBatchDataSourceClass() MLBatchDataSourceClass {
	_MLBatchDataSourceClassOnce.Do(func() {
		_MLBatchDataSourceClass = MLBatchDataSourceClass{class: objc.GetClass("_MLBatchDataSource")}
	})
	return _MLBatchDataSourceClass
}

// GetMLBatchDataSourceClass returns the class object for _MLBatchDataSource.
func GetMLBatchDataSourceClass() MLBatchDataSourceClass {
	return getMLBatchDataSourceClass()
}

type MLBatchDataSourceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLBatchDataSourceClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLBatchDataSourceClass) Alloc() MLBatchDataSource {
	rv := objc.SendIfResponds[MLBatchDataSource](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLBatchDataSource.BatchProvider]
//   - [MLBatchDataSource.DataPointAtIndexError]
//   - [MLBatchDataSource.NnEngine]
//   - [MLBatchDataSource.NumberOfDataPoints]
//   - [MLBatchDataSource.UseForPrediction]
//   - [MLBatchDataSource.InitWithMLBatchProviderForPredictionNeuralNetworkEngineError]
//   - [MLBatchDataSource.DebugDescription]
//   - [MLBatchDataSource.Description]
//   - [MLBatchDataSource.Hash]
//   - [MLBatchDataSource.Superclass]
type MLBatchDataSource struct {
	objectivec.Object
}

// MLBatchDataSourceFromID constructs a [MLBatchDataSource] from an objc.ID.
func MLBatchDataSourceFromID(id objc.ID) MLBatchDataSource {
	return MLBatchDataSource{objectivec.Object{ID: id}}
}

// Ensure MLBatchDataSource implements IMLBatchDataSource.
var _ IMLBatchDataSource = MLBatchDataSource{}

// An interface definition for the [MLBatchDataSource] class.
//
// # Methods
//
//   - [IMLBatchDataSource.BatchProvider]
//   - [IMLBatchDataSource.DataPointAtIndexError]
//   - [IMLBatchDataSource.NnEngine]
//   - [IMLBatchDataSource.NumberOfDataPoints]
//   - [IMLBatchDataSource.UseForPrediction]
//   - [IMLBatchDataSource.InitWithMLBatchProviderForPredictionNeuralNetworkEngineError]
//   - [IMLBatchDataSource.DebugDescription]
//   - [IMLBatchDataSource.Description]
//   - [IMLBatchDataSource.Hash]
//   - [IMLBatchDataSource.Superclass]
type IMLBatchDataSource interface {
	objectivec.IObject

	// Topic: Methods

	BatchProvider() unsafe.Pointer
	DataPointAtIndexError(index uint64) (objectivec.IObject, error)
	NnEngine() IMLNeuralNetworkEngine
	NumberOfDataPoints() uint64
	UseForPrediction() bool
	InitWithMLBatchProviderForPredictionNeuralNetworkEngineError(provider objectivec.IObject, prediction bool, engine objectivec.IObject) (MLBatchDataSource, error)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (m MLBatchDataSource) Init() MLBatchDataSource {
	rv := objc.SendIfResponds[MLBatchDataSource](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLBatchDataSource) Autorelease() MLBatchDataSource {
	rv := objc.SendIfResponds[MLBatchDataSource](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLBatchDataSource creates a new MLBatchDataSource instance.
func NewMLBatchDataSource() MLBatchDataSource {
	class := getMLBatchDataSourceClass()
	rv := objc.SendIfResponds[MLBatchDataSource](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewMLBatchDataSourceWithMLBatchProviderForPredictionNeuralNetworkEngineError(provider objectivec.IObject, prediction bool, engine objectivec.IObject) (MLBatchDataSource, error) {
	var errorPtr objc.ID
	instance := getMLBatchDataSourceClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithMLBatchProvider:forPrediction:neuralNetworkEngine:error:"), provider, prediction, engine, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLBatchDataSource{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return MLBatchDataSource{}, objc.ErrInitFailed
	}
	return MLBatchDataSourceFromID(rv), nil
}

func (m MLBatchDataSource) DataPointAtIndexError(index uint64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("dataPointAtIndex:error:"), index, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLBatchDataSource) NumberOfDataPoints() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("numberOfDataPoints"))
	return rv
}
func (m MLBatchDataSource) InitWithMLBatchProviderForPredictionNeuralNetworkEngineError(provider objectivec.IObject, prediction bool, engine objectivec.IObject) (MLBatchDataSource, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithMLBatchProvider:forPrediction:neuralNetworkEngine:error:"), provider, prediction, engine, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return *new(MLBatchDataSource), foundation.NSErrorFrom(errorPtr)
	}
	return MLBatchDataSourceFromID(rv), nil

}

func (m MLBatchDataSource) BatchProvider() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](m.ID, objc.Sel("batchProvider"))
	return rv
}
func (m MLBatchDataSource) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLBatchDataSource) Description() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLBatchDataSource) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("hash"))
	return rv
}
func (m MLBatchDataSource) NnEngine() IMLNeuralNetworkEngine {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("nnEngine"))
	return MLNeuralNetworkEngineFromID(objc.ID(rv))
}
func (m MLBatchDataSource) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](m.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
func (m MLBatchDataSource) UseForPrediction() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("useForPrediction"))
	return rv
}
