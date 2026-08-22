// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLBatchProviderUtils] class.
var (
	_MLBatchProviderUtilsClass     MLBatchProviderUtilsClass
	_MLBatchProviderUtilsClassOnce sync.Once
)

func getMLBatchProviderUtilsClass() MLBatchProviderUtilsClass {
	_MLBatchProviderUtilsClassOnce.Do(func() {
		_MLBatchProviderUtilsClass = MLBatchProviderUtilsClass{class: objc.GetClass("MLBatchProviderUtils")}
	})
	return _MLBatchProviderUtilsClass
}

// GetMLBatchProviderUtilsClass returns the class object for MLBatchProviderUtils.
func GetMLBatchProviderUtilsClass() MLBatchProviderUtilsClass {
	return getMLBatchProviderUtilsClass()
}

type MLBatchProviderUtilsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLBatchProviderUtilsClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLBatchProviderUtilsClass) Alloc() MLBatchProviderUtils {
	rv := objc.SendIfResponds[MLBatchProviderUtils](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

type MLBatchProviderUtils struct {
	objectivec.Object
}

// MLBatchProviderUtilsFromID constructs a [MLBatchProviderUtils] from an objc.ID.
func MLBatchProviderUtilsFromID(id objc.ID) MLBatchProviderUtils {
	return MLBatchProviderUtils{objectivec.Object{ID: id}}
}

// Ensure MLBatchProviderUtils implements IMLBatchProviderUtils.
var _ IMLBatchProviderUtils = MLBatchProviderUtils{}

// An interface definition for the [MLBatchProviderUtils] class.
type IMLBatchProviderUtils interface {
	objectivec.IObject
}

// Init initializes the instance.
func (m MLBatchProviderUtils) Init() MLBatchProviderUtils {
	rv := objc.SendIfResponds[MLBatchProviderUtils](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLBatchProviderUtils) Autorelease() MLBatchProviderUtils {
	rv := objc.SendIfResponds[MLBatchProviderUtils](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLBatchProviderUtils creates a new MLBatchProviderUtils instance.
func NewMLBatchProviderUtils() MLBatchProviderUtils {
	class := getMLBatchProviderUtilsClass()
	rv := objc.SendIfResponds[MLBatchProviderUtils](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (_MLBatchProviderUtilsClass MLBatchProviderUtilsClass) BatchFromConcatinatingBatches(batches objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLBatchProviderUtilsClass.class), objc.Sel("batchFromConcatinatingBatches:"), batches)
	return objectivec.Object{ID: rv}
}
func (_MLBatchProviderUtilsClass MLBatchProviderUtilsClass) BatchWithSubsetOfFeaturesNamedFromBatch(named objectivec.IObject, batch objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLBatchProviderUtilsClass.class), objc.Sel("batchWithSubsetOfFeaturesNamed:fromBatch:"), named, batch)
	return objectivec.Object{ID: rv}
}
func (_MLBatchProviderUtilsClass MLBatchProviderUtilsClass) DictionaryFromBatch(batch objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLBatchProviderUtilsClass.class), objc.Sel("dictionaryFromBatch:"), batch)
	return objectivec.Object{ID: rv}
}
func (_MLBatchProviderUtilsClass MLBatchProviderUtilsClass) DictionaryFromBatchFeatureNames(batch objectivec.IObject, names objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLBatchProviderUtilsClass.class), objc.Sel("dictionaryFromBatch:featureNames:"), batch, names)
	return objectivec.Object{ID: rv}
}
func (_MLBatchProviderUtilsClass MLBatchProviderUtilsClass) FeatureDescriptionsByNameForBatchError(batch objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLBatchProviderUtilsClass.class), objc.Sel("featureDescriptionsByNameForBatch:error:"), batch, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_MLBatchProviderUtilsClass MLBatchProviderUtilsClass) FeatureNamesInBatch(batch objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLBatchProviderUtilsClass.class), objc.Sel("featureNamesInBatch:"), batch)
	return objectivec.Object{ID: rv}
}
func (_MLBatchProviderUtilsClass MLBatchProviderUtilsClass) FeatureProviderArrayFromBatch(batch objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLBatchProviderUtilsClass.class), objc.Sel("featureProviderArrayFromBatch:"), batch)
	return objectivec.Object{ID: rv}
}
func (_MLBatchProviderUtilsClass MLBatchProviderUtilsClass) FeatureValueArrayForNameBatchError(name objectivec.IObject, batch objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLBatchProviderUtilsClass.class), objc.Sel("featureValueArrayForName:batch:error:"), name, batch, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_MLBatchProviderUtilsClass MLBatchProviderUtilsClass) LazyBatchIndexIntoBatchIndicesError(batch objectivec.IObject, indices objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLBatchProviderUtilsClass.class), objc.Sel("lazyBatchIndexIntoBatch:indices:error:"), batch, indices, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_MLBatchProviderUtilsClass MLBatchProviderUtilsClass) LazyBatchWindowIntoBatchStartIndexWindowLengthError(batch objectivec.IObject, index uint64, length uint64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLBatchProviderUtilsClass.class), objc.Sel("lazyBatchWindowIntoBatch:startIndex:windowLength:error:"), batch, index, length, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_MLBatchProviderUtilsClass MLBatchProviderUtilsClass) LazyBatchWithFeaturesInBatchAddedToBatchError(batch objectivec.IObject, batch2 objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLBatchProviderUtilsClass.class), objc.Sel("lazyBatchWithFeaturesInBatch:addedToBatch:error:"), batch, batch2, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_MLBatchProviderUtilsClass MLBatchProviderUtilsClass) VectorizeFeaturesNamedFromBatchIntoRowsOfDoubleMatrixError(named objectivec.IObject, batch objectivec.IObject, matrix objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_MLBatchProviderUtilsClass.class), objc.Sel("vectorizeFeaturesNamed:fromBatch:intoRowsOfDoubleMatrix:error:"), named, batch, matrix, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("vectorizeFeaturesNamed:fromBatch:intoRowsOfDoubleMatrix:error: returned NO with nil NSError")
	}
	return rv, nil

}
