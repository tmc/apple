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
	rv := objc.SendIfResponds[MLArrayDictionaryFeatureProvider](objc.ID(mc.class), objc.Sel("alloc"))
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
func (m MLArrayDictionaryFeatureProvider) Init() MLArrayDictionaryFeatureProvider {
	rv := objc.SendIfResponds[MLArrayDictionaryFeatureProvider](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLArrayDictionaryFeatureProvider) Autorelease() MLArrayDictionaryFeatureProvider {
	rv := objc.SendIfResponds[MLArrayDictionaryFeatureProvider](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLArrayDictionaryFeatureProvider creates a new MLArrayDictionaryFeatureProvider instance.
func NewMLArrayDictionaryFeatureProvider() MLArrayDictionaryFeatureProvider {
	class := getMLArrayDictionaryFeatureProviderClass()
	rv := objc.SendIfResponds[MLArrayDictionaryFeatureProvider](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewArrayDictionaryFeatureProviderWithCoder(coder objectivec.IObject) MLArrayDictionaryFeatureProvider {
	instance := getMLArrayDictionaryFeatureProviderClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLArrayDictionaryFeatureProviderFromID(rv)
}

func NewArrayDictionaryFeatureProviderWithDictionaryFeatureProviderArray(array objectivec.IObject) MLArrayDictionaryFeatureProvider {
	instance := getMLArrayDictionaryFeatureProviderClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDictionaryFeatureProviderArray:"), array)
	return MLArrayDictionaryFeatureProviderFromID(rv)
}

func (m MLArrayDictionaryFeatureProvider) EncodeWithCoder(coder foundation.INSCoder) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("encodeWithCoder:"), coder)
}
func (m MLArrayDictionaryFeatureProvider) FeaturesAtIndex(index int64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("featuresAtIndex:"), index)
	return objectivec.Object{ID: rv}
}
func (m MLArrayDictionaryFeatureProvider) InitWithCoder(coder foundation.INSCoder) MLArrayDictionaryFeatureProvider {
	rv := objc.SendIfResponds[MLArrayDictionaryFeatureProvider](m.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (m MLArrayDictionaryFeatureProvider) InitWithDictionaryFeatureProviderArray(array objectivec.IObject) MLArrayDictionaryFeatureProvider {
	rv := objc.SendIfResponds[MLArrayDictionaryFeatureProvider](m.ID, objc.Sel("initWithDictionaryFeatureProviderArray:"), array)
	return rv
}

func (_MLArrayDictionaryFeatureProviderClass MLArrayDictionaryFeatureProviderClass) SupportsSecureCoding() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_MLArrayDictionaryFeatureProviderClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

func (m MLArrayDictionaryFeatureProvider) Array() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("array"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLArrayDictionaryFeatureProvider) Count() int64 {
	rv := objc.SendIfResponds[int64](m.ID, objc.Sel("count"))
	return rv
}
