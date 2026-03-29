// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLDataConversionUtils] class.
var (
	_MLDataConversionUtilsClass     MLDataConversionUtilsClass
	_MLDataConversionUtilsClassOnce sync.Once
)

func getMLDataConversionUtilsClass() MLDataConversionUtilsClass {
	_MLDataConversionUtilsClassOnce.Do(func() {
		_MLDataConversionUtilsClass = MLDataConversionUtilsClass{class: objc.GetClass("MLDataConversionUtils")}
	})
	return _MLDataConversionUtilsClass
}

// GetMLDataConversionUtilsClass returns the class object for MLDataConversionUtils.
func GetMLDataConversionUtilsClass() MLDataConversionUtilsClass {
	return getMLDataConversionUtilsClass()
}

type MLDataConversionUtilsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLDataConversionUtilsClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLDataConversionUtilsClass) Alloc() MLDataConversionUtils {
	rv := objc.Send[MLDataConversionUtils](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLDataConversionUtils
type MLDataConversionUtils struct {
	objectivec.Object
}

// MLDataConversionUtilsFromID constructs a [MLDataConversionUtils] from an objc.ID.
func MLDataConversionUtilsFromID(id objc.ID) MLDataConversionUtils {
	return MLDataConversionUtils{objectivec.Object{ID: id}}
}
// Ensure MLDataConversionUtils implements IMLDataConversionUtils.
var _ IMLDataConversionUtils = MLDataConversionUtils{}

// An interface definition for the [MLDataConversionUtils] class.
//
// See: https://developer.apple.com/documentation/CoreML/MLDataConversionUtils
type IMLDataConversionUtils interface {
	objectivec.IObject
}

// Init initializes the instance.
func (d MLDataConversionUtils) Init() MLDataConversionUtils {
	rv := objc.Send[MLDataConversionUtils](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d MLDataConversionUtils) Autorelease() MLDataConversionUtils {
	rv := objc.Send[MLDataConversionUtils](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLDataConversionUtils creates a new MLDataConversionUtils instance.
func NewMLDataConversionUtils() MLDataConversionUtils {
	class := getMLDataConversionUtilsClass()
	rv := objc.Send[MLDataConversionUtils](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLDataConversionUtils/batchProviderFromEspressoDataProvider:neuralNetworkEngine:options:error:
func (_MLDataConversionUtilsClass MLDataConversionUtilsClass) BatchProviderFromEspressoDataProviderNeuralNetworkEngineOptionsError(provider objectivec.IObject, engine objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLDataConversionUtilsClass.class), objc.Sel("batchProviderFromEspressoDataProvider:neuralNetworkEngine:options:error:"), provider, engine, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLDataConversionUtils/batchProviderFromMLComputeDataProvider:neuralNetworkEngine:options:error:
func (_MLDataConversionUtilsClass MLDataConversionUtilsClass) BatchProviderFromMLComputeDataProviderNeuralNetworkEngineOptionsError(provider objectivec.IObject, engine objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLDataConversionUtilsClass.class), objc.Sel("batchProviderFromMLComputeDataProvider:neuralNetworkEngine:options:error:"), provider, engine, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLDataConversionUtils/espressoDataProviderFromBatchProvider:forPrediction:neuralNetworkEngine:error:
func (_MLDataConversionUtilsClass MLDataConversionUtilsClass) EspressoDataProviderFromBatchProviderForPredictionNeuralNetworkEngineError(provider objectivec.IObject, prediction bool, engine objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLDataConversionUtilsClass.class), objc.Sel("espressoDataProviderFromBatchProvider:forPrediction:neuralNetworkEngine:error:"), provider, prediction, engine, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLDataConversionUtils/espressoDataProviderFromFeatureProvider:forPrediction:neuralNetworkEngine:error:
func (_MLDataConversionUtilsClass MLDataConversionUtilsClass) EspressoDataProviderFromFeatureProviderForPredictionNeuralNetworkEngineError(provider objectivec.IObject, prediction bool, engine objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLDataConversionUtilsClass.class), objc.Sel("espressoDataProviderFromFeatureProvider:forPrediction:neuralNetworkEngine:error:"), provider, prediction, engine, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLDataConversionUtils/featureProviderFomMLComputeDataProvider:neuralNetworkEngine:options:error:
func (_MLDataConversionUtilsClass MLDataConversionUtilsClass) FeatureProviderFomMLComputeDataProviderNeuralNetworkEngineOptionsError(provider objectivec.IObject, engine objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLDataConversionUtilsClass.class), objc.Sel("featureProviderFomMLComputeDataProvider:neuralNetworkEngine:options:error:"), provider, engine, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLDataConversionUtils/featureProviderFromEspressoDataProvider:neuralNetworkEngine:options:error:
func (_MLDataConversionUtilsClass MLDataConversionUtilsClass) FeatureProviderFromEspressoDataProviderNeuralNetworkEngineOptionsError(provider objectivec.IObject, engine objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLDataConversionUtilsClass.class), objc.Sel("featureProviderFromEspressoDataProvider:neuralNetworkEngine:options:error:"), provider, engine, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLDataConversionUtils/mlComputeDataProviderFromBatchProvider:batchSize:forPrediction:neuralNetworkEngine:error:
func (_MLDataConversionUtilsClass MLDataConversionUtilsClass) MlComputeDataProviderFromBatchProviderBatchSizeForPredictionNeuralNetworkEngineError(provider objectivec.IObject, size uint64, prediction bool, engine objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLDataConversionUtilsClass.class), objc.Sel("mlComputeDataProviderFromBatchProvider:batchSize:forPrediction:neuralNetworkEngine:error:"), provider, size, prediction, engine, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLDataConversionUtils/mlComputeDataTypeSize:
func (_MLDataConversionUtilsClass MLDataConversionUtilsClass) MlComputeDataTypeSize(size int64) uint64 {
	rv := objc.Send[uint64](objc.ID(_MLDataConversionUtilsClass.class), objc.Sel("mlComputeDataTypeSize:"), size)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLDataConversionUtils/populateEspressoShapeAndStridesFromInputShape:ndRepresentation:espressoShape:espressoStrides:error:
func (_MLDataConversionUtilsClass MLDataConversionUtilsClass) PopulateEspressoShapeAndStridesFromInputShapeNdRepresentationEspressoShapeEspressoStridesError(shape objectivec.IObject, representation bool, shape2 []objectivec.IObject, strides []objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_MLDataConversionUtilsClass.class), objc.Sel("populateEspressoShapeAndStridesFromInputShape:ndRepresentation:espressoShape:espressoStrides:error:"), shape, representation, objectivec.IObjectSliceToNSArray(shape2), objectivec.IObjectSliceToNSArray(strides), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("populateEspressoShapeAndStridesFromInputShape:ndRepresentation:espressoShape:espressoStrides:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLDataConversionUtils/populateShapeAndStrideFor:inputShape:outputShape:outputStrides:error:
func (_MLDataConversionUtilsClass MLDataConversionUtilsClass) PopulateShapeAndStrideForInputShapeOutputShapeOutputStridesError(for_ objectivec.IObject, shape objectivec.IObject, shape2 []objectivec.IObject, strides []objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_MLDataConversionUtilsClass.class), objc.Sel("populateShapeAndStrideFor:inputShape:outputShape:outputStrides:error:"), for_, shape, objectivec.IObjectSliceToNSArray(shape2), objectivec.IObjectSliceToNSArray(strides), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("populateShapeAndStrideFor:inputShape:outputShape:outputStrides:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLDataConversionUtils/sizeFromShape:
func (_MLDataConversionUtilsClass MLDataConversionUtilsClass) SizeFromShape(shape objectivec.IObject) uint64 {
	rv := objc.Send[uint64](objc.ID(_MLDataConversionUtilsClass.class), objc.Sel("sizeFromShape:"), shape)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLDataConversionUtils/stridesForShape:
func (_MLDataConversionUtilsClass MLDataConversionUtilsClass) StridesForShape(shape objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLDataConversionUtilsClass.class), objc.Sel("stridesForShape:"), shape)
	return objectivec.Object{ID: rv}
}

