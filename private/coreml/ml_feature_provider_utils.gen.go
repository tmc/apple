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

// The class instance for the [MLFeatureProviderUtils] class.
var (
	_MLFeatureProviderUtilsClass     MLFeatureProviderUtilsClass
	_MLFeatureProviderUtilsClassOnce sync.Once
)

func getMLFeatureProviderUtilsClass() MLFeatureProviderUtilsClass {
	_MLFeatureProviderUtilsClassOnce.Do(func() {
		_MLFeatureProviderUtilsClass = MLFeatureProviderUtilsClass{class: objc.GetClass("MLFeatureProviderUtils")}
	})
	return _MLFeatureProviderUtilsClass
}

// GetMLFeatureProviderUtilsClass returns the class object for MLFeatureProviderUtils.
func GetMLFeatureProviderUtilsClass() MLFeatureProviderUtilsClass {
	return getMLFeatureProviderUtilsClass()
}

type MLFeatureProviderUtilsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLFeatureProviderUtilsClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLFeatureProviderUtilsClass) Alloc() MLFeatureProviderUtils {
	rv := objc.Send[MLFeatureProviderUtils](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLFeatureProviderUtils
type MLFeatureProviderUtils struct {
	objectivec.Object
}

// MLFeatureProviderUtilsFromID constructs a [MLFeatureProviderUtils] from an objc.ID.
func MLFeatureProviderUtilsFromID(id objc.ID) MLFeatureProviderUtils {
	return MLFeatureProviderUtils{objectivec.Object{ID: id}}
}
// Ensure MLFeatureProviderUtils implements IMLFeatureProviderUtils.
var _ IMLFeatureProviderUtils = MLFeatureProviderUtils{}

// An interface definition for the [MLFeatureProviderUtils] class.
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureProviderUtils
type IMLFeatureProviderUtils interface {
	objectivec.IObject
}

// Init initializes the instance.
func (f MLFeatureProviderUtils) Init() MLFeatureProviderUtils {
	rv := objc.Send[MLFeatureProviderUtils](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f MLFeatureProviderUtils) Autorelease() MLFeatureProviderUtils {
	rv := objc.Send[MLFeatureProviderUtils](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLFeatureProviderUtils creates a new MLFeatureProviderUtils instance.
func NewMLFeatureProviderUtils() MLFeatureProviderUtils {
	class := getMLFeatureProviderUtilsClass()
	rv := objc.Send[MLFeatureProviderUtils](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureProviderUtils/_featureValuesForNames:providedBy:error:
func (_MLFeatureProviderUtilsClass MLFeatureProviderUtilsClass) _featureValuesForNamesProvidedByError(names objectivec.IObject, by objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLFeatureProviderUtilsClass.class), objc.Sel("_featureValuesForNames:providedBy:error:"), names, by, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// FeatureValuesForNamesProvidedByError is an exported wrapper for the private method _featureValuesForNamesProvidedByError.
func (_MLFeatureProviderUtilsClass MLFeatureProviderUtilsClass) FeatureValuesForNamesProvidedByError(names objectivec.IObject, by objectivec.IObject) (objectivec.IObject, error) {
	return _MLFeatureProviderUtilsClass._featureValuesForNamesProvidedByError(names, by)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureProviderUtils/_vectorizeWithoutSizeCheckFeatureValues:intoDoubleVector:stride:error:
func (_MLFeatureProviderUtilsClass MLFeatureProviderUtilsClass) _vectorizeWithoutSizeCheckFeatureValuesIntoDoubleVectorStrideError(values objectivec.IObject, stride uint64) (float64, error) {
	var vector float64
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_MLFeatureProviderUtilsClass.class), objc.Sel("_vectorizeWithoutSizeCheckFeatureValues:intoDoubleVector:stride:error:"), values, unsafe.Pointer(&vector), stride, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0.0, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return 0.0, errors.New("_vectorizeWithoutSizeCheckFeatureValues:intoDoubleVector:stride:error: returned NO with nil NSError")
	}
	return vector, nil
}

// VectorizeWithoutSizeCheckFeatureValuesIntoDoubleVectorStrideError is an exported wrapper for the private method _vectorizeWithoutSizeCheckFeatureValuesIntoDoubleVectorStrideError.
func (_MLFeatureProviderUtilsClass MLFeatureProviderUtilsClass) VectorizeWithoutSizeCheckFeatureValuesIntoDoubleVectorStrideError(values objectivec.IObject, stride uint64) (float64, error) {
	return _MLFeatureProviderUtilsClass._vectorizeWithoutSizeCheckFeatureValuesIntoDoubleVectorStrideError(values, stride)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureProviderUtils/_vectorizedSizeOfFeatureValues:error:
func (_MLFeatureProviderUtilsClass MLFeatureProviderUtilsClass) _vectorizedSizeOfFeatureValuesError(values objectivec.IObject) (int64, error) {
	var errorPtr objc.ID
	rv := objc.Send[int64](objc.ID(_MLFeatureProviderUtilsClass.class), objc.Sel("_vectorizedSizeOfFeatureValues:error:"), values, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}

// VectorizedSizeOfFeatureValuesError is an exported wrapper for the private method _vectorizedSizeOfFeatureValuesError.
func (_MLFeatureProviderUtilsClass MLFeatureProviderUtilsClass) VectorizedSizeOfFeatureValuesError(values objectivec.IObject) (int64, error) {
	return _MLFeatureProviderUtilsClass._vectorizedSizeOfFeatureValuesError(values)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureProviderUtils/canVectorizeAllFeaturesWithDescriptions:
func (_MLFeatureProviderUtilsClass MLFeatureProviderUtilsClass) CanVectorizeAllFeaturesWithDescriptions(descriptions objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_MLFeatureProviderUtilsClass.class), objc.Sel("canVectorizeAllFeaturesWithDescriptions:"), descriptions)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureProviderUtils/canVectorizeFeatureWithDescription:
func (_MLFeatureProviderUtilsClass MLFeatureProviderUtilsClass) CanVectorizeFeatureWithDescription(description objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_MLFeatureProviderUtilsClass.class), objc.Sel("canVectorizeFeatureWithDescription:"), description)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureProviderUtils/lazyProviderWithFeaturesProvidedBy:addedToFeaturesProvidedBy:
func (_MLFeatureProviderUtilsClass MLFeatureProviderUtilsClass) LazyProviderWithFeaturesProvidedByAddedToFeaturesProvidedBy(by objectivec.IObject, by2 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLFeatureProviderUtilsClass.class), objc.Sel("lazyProviderWithFeaturesProvidedBy:addedToFeaturesProvidedBy:"), by, by2)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureProviderUtils/providerWithSubsetOfFeaturesNamed:providedBy:
func (_MLFeatureProviderUtilsClass MLFeatureProviderUtilsClass) ProviderWithSubsetOfFeaturesNamedProvidedBy(named objectivec.IObject, by objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLFeatureProviderUtilsClass.class), objc.Sel("providerWithSubsetOfFeaturesNamed:providedBy:"), named, by)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureProviderUtils/vectorizeFeaturesProvidedBy:featureNames:error:
func (_MLFeatureProviderUtilsClass MLFeatureProviderUtilsClass) VectorizeFeaturesProvidedByFeatureNamesError(by objectivec.IObject, names objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLFeatureProviderUtilsClass.class), objc.Sel("vectorizeFeaturesProvidedBy:featureNames:error:"), by, names, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureProviderUtils/vectorizeFeaturesProvidedBy:featureNames:intoDoubleVector:length:stride:error:
func (_MLFeatureProviderUtilsClass MLFeatureProviderUtilsClass) VectorizeFeaturesProvidedByFeatureNamesIntoDoubleVectorLengthStrideError(by objectivec.IObject, names objectivec.IObject, length uint64, stride uint64) (float64, error) {
	var vector float64
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_MLFeatureProviderUtilsClass.class), objc.Sel("vectorizeFeaturesProvidedBy:featureNames:intoDoubleVector:length:stride:error:"), by, names, unsafe.Pointer(&vector), length, stride, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0.0, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return 0.0, errors.New("vectorizeFeaturesProvidedBy:featureNames:intoDoubleVector:length:stride:error: returned NO with nil NSError")
	}
	return vector, nil
}

