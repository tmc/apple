// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLDataSource] class.
var (
	_MLDataSourceClass     MLDataSourceClass
	_MLDataSourceClassOnce sync.Once
)

func getMLDataSourceClass() MLDataSourceClass {
	_MLDataSourceClassOnce.Do(func() {
		_MLDataSourceClass = MLDataSourceClass{class: objc.GetClass("_MLDataSource")}
	})
	return _MLDataSourceClass
}

// GetMLDataSourceClass returns the class object for _MLDataSource.
func GetMLDataSourceClass() MLDataSourceClass {
	return getMLDataSourceClass()
}

type MLDataSourceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLDataSourceClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLDataSourceClass) Alloc() MLDataSource {
	rv := objc.Send[MLDataSource](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLDataSource.DataPointAtIndexError]
//   - [MLDataSource.DataTensorDictionary]
//   - [MLDataSource.SetDataTensorDictionary]
//   - [MLDataSource.NumberOfDataPoints]
//   - [MLDataSource.InitWithMLFeatureProviderForPredictionNeuralNetworkEngineError]
//   - [MLDataSource.DebugDescription]
//   - [MLDataSource.Description]
//   - [MLDataSource.Hash]
//   - [MLDataSource.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/_MLDataSource
type MLDataSource struct {
	objectivec.Object
}

// MLDataSourceFromID constructs a [MLDataSource] from an objc.ID.
func MLDataSourceFromID(id objc.ID) MLDataSource {
	return MLDataSource{objectivec.Object{ID: id}}
}

// Ensure MLDataSource implements IMLDataSource.
var _ IMLDataSource = MLDataSource{}

// An interface definition for the [MLDataSource] class.
//
// # Methods
//
//   - [IMLDataSource.DataPointAtIndexError]
//   - [IMLDataSource.DataTensorDictionary]
//   - [IMLDataSource.SetDataTensorDictionary]
//   - [IMLDataSource.NumberOfDataPoints]
//   - [IMLDataSource.InitWithMLFeatureProviderForPredictionNeuralNetworkEngineError]
//   - [IMLDataSource.DebugDescription]
//   - [IMLDataSource.Description]
//   - [IMLDataSource.Hash]
//   - [IMLDataSource.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/_MLDataSource
type IMLDataSource interface {
	objectivec.IObject

	// Topic: Methods

	DataPointAtIndexError(index uint64) (objectivec.IObject, error)
	DataTensorDictionary() foundation.INSDictionary
	SetDataTensorDictionary(value foundation.INSDictionary)
	NumberOfDataPoints() uint64
	InitWithMLFeatureProviderForPredictionNeuralNetworkEngineError(provider objectivec.IObject, prediction bool, engine objectivec.IObject) (MLDataSource, error)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (m MLDataSource) Init() MLDataSource {
	rv := objc.Send[MLDataSource](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLDataSource) Autorelease() MLDataSource {
	rv := objc.Send[MLDataSource](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLDataSource creates a new MLDataSource instance.
func NewMLDataSource() MLDataSource {
	class := getMLDataSourceClass()
	rv := objc.Send[MLDataSource](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/_MLDataSource/initWithMLFeatureProvider:forPrediction:neuralNetworkEngine:error:
func NewMLDataSourceWithMLFeatureProviderForPredictionNeuralNetworkEngineError(provider objectivec.IObject, prediction bool, engine objectivec.IObject) (MLDataSource, error) {
	var errorPtr objc.ID
	instance := getMLDataSourceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMLFeatureProvider:forPrediction:neuralNetworkEngine:error:"), provider, prediction, engine, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLDataSource{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLDataSourceFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/_MLDataSource/dataPointAtIndex:error:
func (m MLDataSource) DataPointAtIndexError(index uint64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("dataPointAtIndex:error:"), index, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/_MLDataSource/numberOfDataPoints
func (m MLDataSource) NumberOfDataPoints() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("numberOfDataPoints"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/_MLDataSource/initWithMLFeatureProvider:forPrediction:neuralNetworkEngine:error:
func (m MLDataSource) InitWithMLFeatureProviderForPredictionNeuralNetworkEngineError(provider objectivec.IObject, prediction bool, engine objectivec.IObject) (MLDataSource, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithMLFeatureProvider:forPrediction:neuralNetworkEngine:error:"), provider, prediction, engine, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLDataSource{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLDataSourceFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/_MLDataSource/dataTensorDictionary
func (m MLDataSource) DataTensorDictionary() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("dataTensorDictionary"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLDataSource) SetDataTensorDictionary(value foundation.INSDictionary) {
	objc.Send[struct{}](m.ID, objc.Sel("setDataTensorDictionary:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/_MLDataSource/debugDescription
func (m MLDataSource) DebugDescription() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/_MLDataSource/description
func (m MLDataSource) Description() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/_MLDataSource/hash
func (m MLDataSource) Hash() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/_MLDataSource/superclass
func (m MLDataSource) Superclass() objc.Class {
	rv := objc.Send[objc.Class](m.ID, objc.Sel("superclass"))
	return rv
}
