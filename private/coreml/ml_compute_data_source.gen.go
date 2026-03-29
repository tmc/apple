// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corevideo"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLComputeDataSource] class.
var (
	_MLComputeDataSourceClass     MLComputeDataSourceClass
	_MLComputeDataSourceClassOnce sync.Once
)

func getMLComputeDataSourceClass() MLComputeDataSourceClass {
	_MLComputeDataSourceClassOnce.Do(func() {
		_MLComputeDataSourceClass = MLComputeDataSourceClass{class: objc.GetClass("MLComputeDataSource")}
	})
	return _MLComputeDataSourceClass
}

// GetMLComputeDataSourceClass returns the class object for MLComputeDataSource.
func GetMLComputeDataSourceClass() MLComputeDataSourceClass {
	return getMLComputeDataSourceClass()
}

type MLComputeDataSourceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLComputeDataSourceClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLComputeDataSourceClass) Alloc() MLComputeDataSource {
	rv := objc.Send[MLComputeDataSource](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLComputeDataSource.DataCHWFromChanneledPixelTypeChannelOrderIsBGRError]
//   - [MLComputeDataSource.DataCHWFromPixelBufferChannelOrderIsBGRError]
//   - [MLComputeDataSource.DataCHWFromPixelTypeGray8Error]
//   - [MLComputeDataSource.DataDictionary]
//   - [MLComputeDataSource.OneHotEncodedDataFromFeatureValueWithNNEngineError]
//   - [MLComputeDataSource.InitWithFeatureProviderForPredictionNeuralNetworkEngineError]
// See: https://developer.apple.com/documentation/CoreML/MLComputeDataSource
type MLComputeDataSource struct {
	objectivec.Object
}

// MLComputeDataSourceFromID constructs a [MLComputeDataSource] from an objc.ID.
func MLComputeDataSourceFromID(id objc.ID) MLComputeDataSource {
	return MLComputeDataSource{objectivec.Object{ID: id}}
}
// Ensure MLComputeDataSource implements IMLComputeDataSource.
var _ IMLComputeDataSource = MLComputeDataSource{}

// An interface definition for the [MLComputeDataSource] class.
//
// # Methods
//
//   - [IMLComputeDataSource.DataCHWFromChanneledPixelTypeChannelOrderIsBGRError]
//   - [IMLComputeDataSource.DataCHWFromPixelBufferChannelOrderIsBGRError]
//   - [IMLComputeDataSource.DataCHWFromPixelTypeGray8Error]
//   - [IMLComputeDataSource.DataDictionary]
//   - [IMLComputeDataSource.OneHotEncodedDataFromFeatureValueWithNNEngineError]
//   - [IMLComputeDataSource.InitWithFeatureProviderForPredictionNeuralNetworkEngineError]
//
// See: https://developer.apple.com/documentation/CoreML/MLComputeDataSource
type IMLComputeDataSource interface {
	objectivec.IObject

	// Topic: Methods

	DataCHWFromChanneledPixelTypeChannelOrderIsBGRError(type_ corevideo.CVImageBufferRef, bgr bool) (objectivec.IObject, error)
	DataCHWFromPixelBufferChannelOrderIsBGRError(buffer corevideo.CVImageBufferRef, bgr bool) (objectivec.IObject, error)
	DataCHWFromPixelTypeGray8Error(gray8 corevideo.CVImageBufferRef) (objectivec.IObject, error)
	DataDictionary() foundation.INSDictionary
	OneHotEncodedDataFromFeatureValueWithNNEngineError(value objectivec.IObject, nNEngine objectivec.IObject) (objectivec.IObject, error)
	InitWithFeatureProviderForPredictionNeuralNetworkEngineError(provider objectivec.IObject, prediction bool, engine objectivec.IObject) (MLComputeDataSource, error)
}

// Init initializes the instance.
func (c MLComputeDataSource) Init() MLComputeDataSource {
	rv := objc.Send[MLComputeDataSource](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MLComputeDataSource) Autorelease() MLComputeDataSource {
	rv := objc.Send[MLComputeDataSource](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLComputeDataSource creates a new MLComputeDataSource instance.
func NewMLComputeDataSource() MLComputeDataSource {
	class := getMLComputeDataSourceClass()
	rv := objc.Send[MLComputeDataSource](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLComputeDataSource/initWithFeatureProvider:forPrediction:neuralNetworkEngine:error:
func NewComputeDataSourceWithFeatureProviderForPredictionNeuralNetworkEngineError(provider objectivec.IObject, prediction bool, engine objectivec.IObject) (MLComputeDataSource, error) {
	var errorPtr objc.ID
	instance := getMLComputeDataSourceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFeatureProvider:forPrediction:neuralNetworkEngine:error:"), provider, prediction, engine, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLComputeDataSource{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLComputeDataSourceFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLComputeDataSource/dataCHWFromChanneledPixelType:channelOrderIsBGR:error:
func (c MLComputeDataSource) DataCHWFromChanneledPixelTypeChannelOrderIsBGRError(type_ corevideo.CVImageBufferRef, bgr bool) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](c.ID, objc.Sel("dataCHWFromChanneledPixelType:channelOrderIsBGR:error:"), type_, bgr, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLComputeDataSource/dataCHWFromPixelBuffer:channelOrderIsBGR:error:
func (c MLComputeDataSource) DataCHWFromPixelBufferChannelOrderIsBGRError(buffer corevideo.CVImageBufferRef, bgr bool) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](c.ID, objc.Sel("dataCHWFromPixelBuffer:channelOrderIsBGR:error:"), buffer, bgr, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLComputeDataSource/dataCHWFromPixelTypeGray8:error:
func (c MLComputeDataSource) DataCHWFromPixelTypeGray8Error(gray8 corevideo.CVImageBufferRef) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](c.ID, objc.Sel("dataCHWFromPixelTypeGray8:error:"), gray8, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLComputeDataSource/oneHotEncodedDataFromFeatureValue:withNNEngine:error:
func (c MLComputeDataSource) OneHotEncodedDataFromFeatureValueWithNNEngineError(value objectivec.IObject, nNEngine objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](c.ID, objc.Sel("oneHotEncodedDataFromFeatureValue:withNNEngine:error:"), value, nNEngine, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLComputeDataSource/initWithFeatureProvider:forPrediction:neuralNetworkEngine:error:
func (c MLComputeDataSource) InitWithFeatureProviderForPredictionNeuralNetworkEngineError(provider objectivec.IObject, prediction bool, engine objectivec.IObject) (MLComputeDataSource, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](c.ID, objc.Sel("initWithFeatureProvider:forPrediction:neuralNetworkEngine:error:"), provider, prediction, engine, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLComputeDataSource{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLComputeDataSourceFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/MLComputeDataSource/dataDictionary
func (c MLComputeDataSource) DataDictionary() foundation.INSDictionary {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("dataDictionary"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

