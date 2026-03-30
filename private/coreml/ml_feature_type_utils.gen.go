// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLFeatureTypeUtils] class.
var (
	_MLFeatureTypeUtilsClass     MLFeatureTypeUtilsClass
	_MLFeatureTypeUtilsClassOnce sync.Once
)

func getMLFeatureTypeUtilsClass() MLFeatureTypeUtilsClass {
	_MLFeatureTypeUtilsClassOnce.Do(func() {
		_MLFeatureTypeUtilsClass = MLFeatureTypeUtilsClass{class: objc.GetClass("MLFeatureTypeUtils")}
	})
	return _MLFeatureTypeUtilsClass
}

// GetMLFeatureTypeUtilsClass returns the class object for MLFeatureTypeUtils.
func GetMLFeatureTypeUtilsClass() MLFeatureTypeUtilsClass {
	return getMLFeatureTypeUtilsClass()
}

type MLFeatureTypeUtilsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLFeatureTypeUtilsClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLFeatureTypeUtilsClass) Alloc() MLFeatureTypeUtils {
	rv := objc.Send[MLFeatureTypeUtils](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLFeatureTypeUtils
type MLFeatureTypeUtils struct {
	objectivec.Object
}

// MLFeatureTypeUtilsFromID constructs a [MLFeatureTypeUtils] from an objc.ID.
func MLFeatureTypeUtilsFromID(id objc.ID) MLFeatureTypeUtils {
	return MLFeatureTypeUtils{objectivec.Object{ID: id}}
}

// Ensure MLFeatureTypeUtils implements IMLFeatureTypeUtils.
var _ IMLFeatureTypeUtils = MLFeatureTypeUtils{}

// An interface definition for the [MLFeatureTypeUtils] class.
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureTypeUtils
type IMLFeatureTypeUtils interface {
	objectivec.IObject
}

// Init initializes the instance.
func (f MLFeatureTypeUtils) Init() MLFeatureTypeUtils {
	rv := objc.Send[MLFeatureTypeUtils](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f MLFeatureTypeUtils) Autorelease() MLFeatureTypeUtils {
	rv := objc.Send[MLFeatureTypeUtils](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLFeatureTypeUtils creates a new MLFeatureTypeUtils instance.
func NewMLFeatureTypeUtils() MLFeatureTypeUtils {
	class := getMLFeatureTypeUtilsClass()
	rv := objc.Send[MLFeatureTypeUtils](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLFeatureTypeUtils/canShapeArrayBePromotedFrom:to:
func (_MLFeatureTypeUtilsClass MLFeatureTypeUtilsClass) CanShapeArrayBePromotedFromTo(from objectivec.IObject, to objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_MLFeatureTypeUtilsClass.class), objc.Sel("canShapeArrayBePromotedFrom:to:"), from, to)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLFeatureTypeUtils/descriptionForType:
func (_MLFeatureTypeUtilsClass MLFeatureTypeUtilsClass) DescriptionForType(type_ int64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLFeatureTypeUtilsClass.class), objc.Sel("descriptionForType:"), type_)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLFeatureTypeUtils/featureDescriptionWithName:consistentWithFeatureValues:error:
func (_MLFeatureTypeUtilsClass MLFeatureTypeUtilsClass) FeatureDescriptionWithNameConsistentWithFeatureValuesError(name objectivec.IObject, values objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLFeatureTypeUtilsClass.class), objc.Sel("featureDescriptionWithName:consistentWithFeatureValues:error:"), name, values, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLFeatureTypeUtils/featureTypeForObject:
func (_MLFeatureTypeUtilsClass MLFeatureTypeUtilsClass) FeatureTypeForObject(object objectivec.IObject) int64 {
	rv := objc.Send[int64](objc.ID(_MLFeatureTypeUtilsClass.class), objc.Sel("featureTypeForObject:"), object)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLFeatureTypeUtils/featureTypeForValuesInArray:error:
func (_MLFeatureTypeUtilsClass MLFeatureTypeUtilsClass) FeatureTypeForValuesInArrayError(array objectivec.IObject) (int64, error) {
	var errorPtr objc.ID
	rv := objc.Send[int64](objc.ID(_MLFeatureTypeUtilsClass.class), objc.Sel("featureTypeForValuesInArray:error:"), array, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return 0, foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLFeatureTypeUtils/featureValuesWithConsistentTypeFromArray:error:
func (_MLFeatureTypeUtilsClass MLFeatureTypeUtilsClass) FeatureValuesWithConsistentTypeFromArrayError(array objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLFeatureTypeUtilsClass.class), objc.Sel("featureValuesWithConsistentTypeFromArray:error:"), array, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
