// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLArrayDictionaryFeatureProvider] class.
var (
	_MLArrayDictionaryFeatureProviderClass     MLArrayDictionaryFeatureProviderClass
	_MLArrayDictionaryFeatureProviderClassOnce sync.Once
)

func getMLArrayDictionaryFeatureProviderClass() MLArrayDictionaryFeatureProviderClass {
	_MLArrayDictionaryFeatureProviderClassOnce.Do(func() {
		_MLArrayDictionaryFeatureProviderClass = MLArrayDictionaryFeatureProviderClass{class: objc.GetClass("MLArrayDictionaryFeatureProvider")}
	})
	return _MLArrayDictionaryFeatureProviderClass
}

// GetMLArrayDictionaryFeatureProviderClass returns the class object for MLArrayDictionaryFeatureProvider.
func GetMLArrayDictionaryFeatureProviderClass() MLArrayDictionaryFeatureProviderClass {
	return getMLArrayDictionaryFeatureProviderClass()
}

type MLArrayDictionaryFeatureProviderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLArrayDictionaryFeatureProviderClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLArrayDictionaryFeatureProviderClass) Alloc() MLArrayDictionaryFeatureProvider {
	rv := objc.Send[MLArrayDictionaryFeatureProvider](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLArrayDictionaryFeatureProvider.Array]
//   - [MLArrayDictionaryFeatureProvider.Count]
//   - [MLArrayDictionaryFeatureProvider.EncodeWithCoder]
//   - [MLArrayDictionaryFeatureProvider.FeaturesAtIndex]
//   - [MLArrayDictionaryFeatureProvider.InitWithCoder]
//   - [MLArrayDictionaryFeatureProvider.InitWithDictionaryFeatureProviderArray]
//
// See: https://developer.apple.com/documentation/CoreML/MLArrayDictionaryFeatureProvider
type MLArrayDictionaryFeatureProvider struct {
	objectivec.Object
}

// MLArrayDictionaryFeatureProviderFromID constructs a [MLArrayDictionaryFeatureProvider] from an objc.ID.
func MLArrayDictionaryFeatureProviderFromID(id objc.ID) MLArrayDictionaryFeatureProvider {
	return MLArrayDictionaryFeatureProvider{objectivec.Object{ID: id}}
}

// Ensure MLArrayDictionaryFeatureProvider implements IMLArrayDictionaryFeatureProvider.
var _ IMLArrayDictionaryFeatureProvider = MLArrayDictionaryFeatureProvider{}

// An interface definition for the [MLArrayDictionaryFeatureProvider] class.
//
// # Methods
//
//   - [IMLArrayDictionaryFeatureProvider.Array]
//   - [IMLArrayDictionaryFeatureProvider.Count]
//   - [IMLArrayDictionaryFeatureProvider.EncodeWithCoder]
//   - [IMLArrayDictionaryFeatureProvider.FeaturesAtIndex]
//   - [IMLArrayDictionaryFeatureProvider.InitWithCoder]
//   - [IMLArrayDictionaryFeatureProvider.InitWithDictionaryFeatureProviderArray]
//
// See: https://developer.apple.com/documentation/CoreML/MLArrayDictionaryFeatureProvider
type IMLArrayDictionaryFeatureProvider interface {
	objectivec.IObject

	// Topic: Methods

	Array() foundation.INSArray
	Count() int64
	EncodeWithCoder(coder foundation.INSCoder)
	FeaturesAtIndex(index int64) objectivec.IObject
	InitWithCoder(coder foundation.INSCoder) MLArrayDictionaryFeatureProvider
	InitWithDictionaryFeatureProviderArray(array objectivec.IObject) MLArrayDictionaryFeatureProvider
}

// Init initializes the instance.
func (a MLArrayDictionaryFeatureProvider) Init() MLArrayDictionaryFeatureProvider {
	rv := objc.Send[MLArrayDictionaryFeatureProvider](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a MLArrayDictionaryFeatureProvider) Autorelease() MLArrayDictionaryFeatureProvider {
	rv := objc.Send[MLArrayDictionaryFeatureProvider](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLArrayDictionaryFeatureProvider creates a new MLArrayDictionaryFeatureProvider instance.
func NewMLArrayDictionaryFeatureProvider() MLArrayDictionaryFeatureProvider {
	class := getMLArrayDictionaryFeatureProviderClass()
	rv := objc.Send[MLArrayDictionaryFeatureProvider](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLArrayDictionaryFeatureProvider/initWithCoder:
func NewArrayDictionaryFeatureProviderWithCoder(coder objectivec.IObject) MLArrayDictionaryFeatureProvider {
	instance := getMLArrayDictionaryFeatureProviderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLArrayDictionaryFeatureProviderFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLArrayDictionaryFeatureProvider/initWithDictionaryFeatureProviderArray:
func NewArrayDictionaryFeatureProviderWithDictionaryFeatureProviderArray(array objectivec.IObject) MLArrayDictionaryFeatureProvider {
	instance := getMLArrayDictionaryFeatureProviderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDictionaryFeatureProviderArray:"), array)
	return MLArrayDictionaryFeatureProviderFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLArrayDictionaryFeatureProvider/encodeWithCoder:
func (a MLArrayDictionaryFeatureProvider) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](a.ID, objc.Sel("encodeWithCoder:"), coder)
}

// See: https://developer.apple.com/documentation/CoreML/MLArrayDictionaryFeatureProvider/featuresAtIndex:
func (a MLArrayDictionaryFeatureProvider) FeaturesAtIndex(index int64) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("featuresAtIndex:"), index)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLArrayDictionaryFeatureProvider/initWithCoder:
func (a MLArrayDictionaryFeatureProvider) InitWithCoder(coder foundation.INSCoder) MLArrayDictionaryFeatureProvider {
	rv := objc.Send[MLArrayDictionaryFeatureProvider](a.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLArrayDictionaryFeatureProvider/initWithDictionaryFeatureProviderArray:
func (a MLArrayDictionaryFeatureProvider) InitWithDictionaryFeatureProviderArray(array objectivec.IObject) MLArrayDictionaryFeatureProvider {
	rv := objc.Send[MLArrayDictionaryFeatureProvider](a.ID, objc.Sel("initWithDictionaryFeatureProviderArray:"), array)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLArrayDictionaryFeatureProvider/supportsSecureCoding
func (_MLArrayDictionaryFeatureProviderClass MLArrayDictionaryFeatureProviderClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_MLArrayDictionaryFeatureProviderClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLArrayDictionaryFeatureProvider/array
func (a MLArrayDictionaryFeatureProvider) Array() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("array"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLArrayDictionaryFeatureProvider/count
func (a MLArrayDictionaryFeatureProvider) Count() int64 {
	rv := objc.Send[int64](a.ID, objc.Sel("count"))
	return rv
}
