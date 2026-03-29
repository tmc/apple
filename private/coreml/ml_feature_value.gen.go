// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLFeatureValue] class.
var (
	_MLFeatureValueClass     MLFeatureValueClass
	_MLFeatureValueClassOnce sync.Once
)

func getMLFeatureValueClass() MLFeatureValueClass {
	_MLFeatureValueClassOnce.Do(func() {
		_MLFeatureValueClass = MLFeatureValueClass{class: objc.GetClass("MLFeatureValue")}
	})
	return _MLFeatureValueClass
}

// GetMLFeatureValueClass returns the class object for MLFeatureValue.
func GetMLFeatureValueClass() MLFeatureValueClass {
	return getMLFeatureValueClass()
}

type MLFeatureValueClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLFeatureValueClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLFeatureValueClass) Alloc() MLFeatureValue {
	rv := objc.Send[MLFeatureValue](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLFeatureValue.DebugQuickLookObject]
//   - [MLFeatureValue.GetFeatureSize]
//   - [MLFeatureValue.GetFeatureSizeNdArrayMode]
//   - [MLFeatureValue.InternalStateValue]
//   - [MLFeatureValue.ObjectValue]
//   - [MLFeatureValue.SetObjectValue]
//   - [MLFeatureValue.StateValue]
//   - [MLFeatureValue.Value]
//   - [MLFeatureValue.SetValue]
//   - [MLFeatureValue.InitWithCoder]
//   - [MLFeatureValue.InitWithUndefinedValueAndType]
//   - [MLFeatureValue.InitWithValueType]
//   - [MLFeatureValue.Undefined]
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue
type MLFeatureValue struct {
	objectivec.Object
}

// MLFeatureValueFromID constructs a [MLFeatureValue] from an objc.ID.
func MLFeatureValueFromID(id objc.ID) MLFeatureValue {
	return MLFeatureValue{objectivec.Object{ID: id}}
}
// Ensure MLFeatureValue implements IMLFeatureValue.
var _ IMLFeatureValue = MLFeatureValue{}

// An interface definition for the [MLFeatureValue] class.
//
// # Methods
//
//   - [IMLFeatureValue.DebugQuickLookObject]
//   - [IMLFeatureValue.GetFeatureSize]
//   - [IMLFeatureValue.GetFeatureSizeNdArrayMode]
//   - [IMLFeatureValue.InternalStateValue]
//   - [IMLFeatureValue.ObjectValue]
//   - [IMLFeatureValue.SetObjectValue]
//   - [IMLFeatureValue.StateValue]
//   - [IMLFeatureValue.Value]
//   - [IMLFeatureValue.SetValue]
//   - [IMLFeatureValue.InitWithCoder]
//   - [IMLFeatureValue.InitWithUndefinedValueAndType]
//   - [IMLFeatureValue.InitWithValueType]
//   - [IMLFeatureValue.Undefined]
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue
type IMLFeatureValue interface {
	objectivec.IObject

	// Topic: Methods

	DebugQuickLookObject() objectivec.IObject
	GetFeatureSize(size []objectivec.IObject) objectivec.IObject
	GetFeatureSizeNdArrayMode(size []objectivec.IObject, mode bool) objectivec.IObject
	InternalStateValue() objectivec.IObject
	ObjectValue() objectivec.Object
	SetObjectValue(value objectivec.Object)
	StateValue() objectivec.IObject
	Value() objectivec.IObject
	SetValue(value objectivec.IObject)
	InitWithCoder(coder foundation.INSCoder) MLFeatureValue
	InitWithUndefinedValueAndType(type_ int64) MLFeatureValue
	InitWithValueType(value objectivec.IObject, type_ int64) MLFeatureValue
	Undefined() bool
}

// Init initializes the instance.
func (f MLFeatureValue) Init() MLFeatureValue {
	rv := objc.Send[MLFeatureValue](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f MLFeatureValue) Autorelease() MLFeatureValue {
	rv := objc.Send[MLFeatureValue](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLFeatureValue creates a new MLFeatureValue instance.
func NewMLFeatureValue() MLFeatureValue {
	class := getMLFeatureValueClass()
	rv := objc.Send[MLFeatureValue](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/initWithCoder:
func NewFeatureValueWithCoder(coder objectivec.IObject) MLFeatureValue {
	instance := getMLFeatureValueClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLFeatureValueFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/initWithUndefinedValueAndType:
func NewFeatureValueWithUndefinedValueAndType(type_ int64) MLFeatureValue {
	instance := getMLFeatureValueClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithUndefinedValueAndType:"), type_)
	return MLFeatureValueFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/initWithValue:type:
func NewFeatureValueWithValueType(value objectivec.IObject, type_ int64) MLFeatureValue {
	instance := getMLFeatureValueClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithValue:type:"), value, type_)
	return MLFeatureValueFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/debugQuickLookObject
func (f MLFeatureValue) DebugQuickLookObject() objectivec.IObject {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("debugQuickLookObject"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/getFeatureSize:
func (f MLFeatureValue) GetFeatureSize(size []objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("getFeatureSize:"), objectivec.IObjectSliceToNSArray(size))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/getFeatureSize:ndArrayMode:
func (f MLFeatureValue) GetFeatureSizeNdArrayMode(size []objectivec.IObject, mode bool) objectivec.IObject {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("getFeatureSize:ndArrayMode:"), objectivec.IObjectSliceToNSArray(size), mode)
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/internalStateValue
func (f MLFeatureValue) InternalStateValue() objectivec.IObject {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("internalStateValue"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/stateValue
func (f MLFeatureValue) StateValue() objectivec.IObject {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("stateValue"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/initWithCoder:
func (f MLFeatureValue) InitWithCoder(coder foundation.INSCoder) MLFeatureValue {
	rv := objc.Send[MLFeatureValue](f.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/initWithUndefinedValueAndType:
func (f MLFeatureValue) InitWithUndefinedValueAndType(type_ int64) MLFeatureValue {
	rv := objc.Send[MLFeatureValue](f.ID, objc.Sel("initWithUndefinedValueAndType:"), type_)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/initWithValue:type:
func (f MLFeatureValue) InitWithValueType(value objectivec.IObject, type_ int64) MLFeatureValue {
	rv := objc.Send[MLFeatureValue](f.ID, objc.Sel("initWithValue:type:"), value, type_)
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/featureValueOfType:fromObject:error:
func (_MLFeatureValueClass MLFeatureValueClass) FeatureValueOfTypeFromObjectError(type_ int64, object objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLFeatureValueClass.class), objc.Sel("featureValueOfType:fromObject:error:"), type_, object, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/featureValueWithDictionary:error:
func (_MLFeatureValueClass MLFeatureValueClass) FeatureValueWithDictionaryError(dictionary objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLFeatureValueClass.class), objc.Sel("featureValueWithDictionary:error:"), dictionary, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/featureValueWithDouble:
func (_MLFeatureValueClass MLFeatureValueClass) FeatureValueWithDouble(double float64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLFeatureValueClass.class), objc.Sel("featureValueWithDouble:"), double)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/featureValueWithImageAtURL:constraint:options:error:
func (_MLFeatureValueClass MLFeatureValueClass) FeatureValueWithImageAtURLConstraintOptionsError(url foundation.INSURL, constraint objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLFeatureValueClass.class), objc.Sel("featureValueWithImageAtURL:constraint:options:error:"), url, constraint, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/featureValueWithImageAtURL:orientation:constraint:options:error:
func (_MLFeatureValueClass MLFeatureValueClass) FeatureValueWithImageAtURLOrientationConstraintOptionsError(url foundation.INSURL, orientation uint32, constraint objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLFeatureValueClass.class), objc.Sel("featureValueWithImageAtURL:orientation:constraint:options:error:"), url, orientation, constraint, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/featureValueWithImageAtURL:orientation:pixelsWide:pixelsHigh:pixelFormatType:options:error:
func (_MLFeatureValueClass MLFeatureValueClass) FeatureValueWithImageAtURLOrientationPixelsWidePixelsHighPixelFormatTypeOptionsError(url foundation.INSURL, orientation uint32, wide int64, high int64, type_ uint32, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLFeatureValueClass.class), objc.Sel("featureValueWithImageAtURL:orientation:pixelsWide:pixelsHigh:pixelFormatType:options:error:"), url, orientation, wide, high, type_, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/featureValueWithImageAtURL:pixelsWide:pixelsHigh:pixelFormatType:options:error:
func (_MLFeatureValueClass MLFeatureValueClass) FeatureValueWithImageAtURLPixelsWidePixelsHighPixelFormatTypeOptionsError(url foundation.INSURL, wide int64, high int64, type_ uint32, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLFeatureValueClass.class), objc.Sel("featureValueWithImageAtURL:pixelsWide:pixelsHigh:pixelFormatType:options:error:"), url, wide, high, type_, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/featureValueWithInt64:
func (_MLFeatureValueClass MLFeatureValueClass) FeatureValueWithInt64(int64_ int64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLFeatureValueClass.class), objc.Sel("featureValueWithInt64:"), int64_)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/featureValueWithInt64KeyDictionary:
func (_MLFeatureValueClass MLFeatureValueClass) FeatureValueWithInt64KeyDictionary(dictionary objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLFeatureValueClass.class), objc.Sel("featureValueWithInt64KeyDictionary:"), dictionary)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/featureValueWithMultiArray:
func (_MLFeatureValueClass MLFeatureValueClass) FeatureValueWithMultiArray(array objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLFeatureValueClass.class), objc.Sel("featureValueWithMultiArray:"), array)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/featureValueWithSequence:
func (_MLFeatureValueClass MLFeatureValueClass) FeatureValueWithSequence(sequence objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLFeatureValueClass.class), objc.Sel("featureValueWithSequence:"), sequence)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/featureValueWithState:
func (_MLFeatureValueClass MLFeatureValueClass) FeatureValueWithState(state objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLFeatureValueClass.class), objc.Sel("featureValueWithState:"), state)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/featureValueWithString:
func (_MLFeatureValueClass MLFeatureValueClass) FeatureValueWithString(string_ objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLFeatureValueClass.class), objc.Sel("featureValueWithString:"), string_)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/featureValueWithStringKeyDictionary:
func (_MLFeatureValueClass MLFeatureValueClass) FeatureValueWithStringKeyDictionary(dictionary objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLFeatureValueClass.class), objc.Sel("featureValueWithStringKeyDictionary:"), dictionary)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/internalFeatureValueWithState:
func (_MLFeatureValueClass MLFeatureValueClass) InternalFeatureValueWithState(state objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLFeatureValueClass.class), objc.Sel("internalFeatureValueWithState:"), state)
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/supportsSecureCoding
func (_MLFeatureValueClass MLFeatureValueClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_MLFeatureValueClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/undefinedFeatureValueWithType:
func (_MLFeatureValueClass MLFeatureValueClass) UndefinedFeatureValueWithType(type_ int64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLFeatureValueClass.class), objc.Sel("undefinedFeatureValueWithType:"), type_)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/visionCropAndScaleOptionFromOptions:
func (_MLFeatureValueClass MLFeatureValueClass) VisionCropAndScaleOptionFromOptions(options objectivec.IObject) uint64 {
	rv := objc.Send[uint64](objc.ID(_MLFeatureValueClass.class), objc.Sel("visionCropAndScaleOptionFromOptions:"), options)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/objectValue
func (f MLFeatureValue) ObjectValue() objectivec.Object {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("objectValue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (f MLFeatureValue) SetObjectValue(value objectivec.Object) {
	objc.Send[struct{}](f.ID, objc.Sel("setObjectValue:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/undefined
func (f MLFeatureValue) Undefined() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("undefined"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLFeatureValue/value
func (f MLFeatureValue) Value() objectivec.IObject {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("value"))
	return objectivec.Object{ID: rv}
}
func (f MLFeatureValue) SetValue(value objectivec.IObject) {
	objc.Send[struct{}](f.ID, objc.Sel("setValue:"), value)
}

