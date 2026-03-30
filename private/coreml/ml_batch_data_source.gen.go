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
	rv := objc.Send[MLBatchDataSource](objc.ID(mc.class), objc.Sel("alloc"))
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
//
// See: https://developer.apple.com/documentation/CoreML/_MLBatchDataSource
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
//
// See: https://developer.apple.com/documentation/CoreML/_MLBatchDataSource
type IMLBatchDataSource interface {
	objectivec.IObject

	// Topic: Methods

	BatchProvider() objectivec.IObject
	DataPointAtIndexError(index uint64) (objectivec.IObject, error)
	NnEngine() IMLNeuralNetworkEngine
	NumberOfDataPoints() uint64
	UseForPrediction() bool
	InitWithMLBatchProviderForPredictionNeuralNetworkEngineError(provider objectivec.IObject, prediction bool, engine objectivec.IObject) (MLBatchDataSource, error)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (m MLBatchDataSource) Init() MLBatchDataSource {
	rv := objc.Send[MLBatchDataSource](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLBatchDataSource) Autorelease() MLBatchDataSource {
	rv := objc.Send[MLBatchDataSource](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLBatchDataSource creates a new MLBatchDataSource instance.
func NewMLBatchDataSource() MLBatchDataSource {
	class := getMLBatchDataSourceClass()
	rv := objc.Send[MLBatchDataSource](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/_MLBatchDataSource/initWithMLBatchProvider:forPrediction:neuralNetworkEngine:error:
func NewMLBatchDataSourceWithMLBatchProviderForPredictionNeuralNetworkEngineError(provider objectivec.IObject, prediction bool, engine objectivec.IObject) (MLBatchDataSource, error) {
	var errorPtr objc.ID
	instance := getMLBatchDataSourceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMLBatchProvider:forPrediction:neuralNetworkEngine:error:"), provider, prediction, engine, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLBatchDataSource{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLBatchDataSourceFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/_MLBatchDataSource/dataPointAtIndex:error:
func (m MLBatchDataSource) DataPointAtIndexError(index uint64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("dataPointAtIndex:error:"), index, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/_MLBatchDataSource/numberOfDataPoints
func (m MLBatchDataSource) NumberOfDataPoints() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("numberOfDataPoints"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/_MLBatchDataSource/initWithMLBatchProvider:forPrediction:neuralNetworkEngine:error:
func (m MLBatchDataSource) InitWithMLBatchProviderForPredictionNeuralNetworkEngineError(provider objectivec.IObject, prediction bool, engine objectivec.IObject) (MLBatchDataSource, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithMLBatchProvider:forPrediction:neuralNetworkEngine:error:"), provider, prediction, engine, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLBatchDataSource{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLBatchDataSourceFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/_MLBatchDataSource/batchProvider
func (m MLBatchDataSource) BatchProvider() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("batchProvider"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/_MLBatchDataSource/debugDescription
func (m MLBatchDataSource) DebugDescription() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/_MLBatchDataSource/description
func (m MLBatchDataSource) Description() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/_MLBatchDataSource/hash
func (m MLBatchDataSource) Hash() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/_MLBatchDataSource/nnEngine
func (m MLBatchDataSource) NnEngine() IMLNeuralNetworkEngine {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("nnEngine"))
	return MLNeuralNetworkEngineFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/_MLBatchDataSource/superclass
func (m MLBatchDataSource) Superclass() objc.Class {
	rv := objc.Send[objc.Class](m.ID, objc.Sel("superclass"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/_MLBatchDataSource/useForPrediction
func (m MLBatchDataSource) UseForPrediction() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("useForPrediction"))
	return rv
}
