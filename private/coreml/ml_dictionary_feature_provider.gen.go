// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLDictionaryFeatureProvider] class.
var (
	_MLDictionaryFeatureProviderClass     MLDictionaryFeatureProviderClass
	_MLDictionaryFeatureProviderClassOnce sync.Once
)

func getMLDictionaryFeatureProviderClass() MLDictionaryFeatureProviderClass {
	_MLDictionaryFeatureProviderClassOnce.Do(func() {
		_MLDictionaryFeatureProviderClass = MLDictionaryFeatureProviderClass{class: objc.GetClass("MLDictionaryFeatureProvider")}
	})
	return _MLDictionaryFeatureProviderClass
}

// GetMLDictionaryFeatureProviderClass returns the class object for MLDictionaryFeatureProvider.
func GetMLDictionaryFeatureProviderClass() MLDictionaryFeatureProviderClass {
	return getMLDictionaryFeatureProviderClass()
}

type MLDictionaryFeatureProviderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLDictionaryFeatureProviderClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLDictionaryFeatureProviderClass) Alloc() MLDictionaryFeatureProvider {
	rv := objc.Send[MLDictionaryFeatureProvider](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLDictionaryFeatureProvider.CachedFeatureNames]
//   - [MLDictionaryFeatureProvider.InitWithCoder]
//   - [MLDictionaryFeatureProvider.InitWithFeatureProvider]
//   - [MLDictionaryFeatureProvider.InitWithFeatureProviderFeatureNames]
//   - [MLDictionaryFeatureProvider.InitWithFeatureValueDictionary]
//   - [MLDictionaryFeatureProvider.Dictionary]
//   - [MLDictionaryFeatureProvider.SetDictionary]
// See: https://developer.apple.com/documentation/CoreML/MLDictionaryFeatureProvider
type MLDictionaryFeatureProvider struct {
	objectivec.Object
}

// MLDictionaryFeatureProviderFromID constructs a [MLDictionaryFeatureProvider] from an objc.ID.
func MLDictionaryFeatureProviderFromID(id objc.ID) MLDictionaryFeatureProvider {
	return MLDictionaryFeatureProvider{objectivec.Object{ID: id}}
}
// Ensure MLDictionaryFeatureProvider implements IMLDictionaryFeatureProvider.
var _ IMLDictionaryFeatureProvider = MLDictionaryFeatureProvider{}

// An interface definition for the [MLDictionaryFeatureProvider] class.
//
// # Methods
//
//   - [IMLDictionaryFeatureProvider.CachedFeatureNames]
//   - [IMLDictionaryFeatureProvider.InitWithCoder]
//   - [IMLDictionaryFeatureProvider.InitWithFeatureProvider]
//   - [IMLDictionaryFeatureProvider.InitWithFeatureProviderFeatureNames]
//   - [IMLDictionaryFeatureProvider.InitWithFeatureValueDictionary]
//   - [IMLDictionaryFeatureProvider.Dictionary]
//   - [IMLDictionaryFeatureProvider.SetDictionary]
//
// See: https://developer.apple.com/documentation/CoreML/MLDictionaryFeatureProvider
type IMLDictionaryFeatureProvider interface {
	objectivec.IObject

	// Topic: Methods

	CachedFeatureNames() foundation.INSSet
	InitWithCoder(coder foundation.INSCoder) MLDictionaryFeatureProvider
	InitWithFeatureProvider(provider objectivec.IObject) MLDictionaryFeatureProvider
	InitWithFeatureProviderFeatureNames(provider objectivec.IObject, names objectivec.IObject) MLDictionaryFeatureProvider
	InitWithFeatureValueDictionary(dictionary objectivec.IObject) MLDictionaryFeatureProvider
	Dictionary() foundation.INSDictionary
	SetDictionary(value foundation.INSDictionary)
}

// Init initializes the instance.
func (d MLDictionaryFeatureProvider) Init() MLDictionaryFeatureProvider {
	rv := objc.Send[MLDictionaryFeatureProvider](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d MLDictionaryFeatureProvider) Autorelease() MLDictionaryFeatureProvider {
	rv := objc.Send[MLDictionaryFeatureProvider](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLDictionaryFeatureProvider creates a new MLDictionaryFeatureProvider instance.
func NewMLDictionaryFeatureProvider() MLDictionaryFeatureProvider {
	class := getMLDictionaryFeatureProviderClass()
	rv := objc.Send[MLDictionaryFeatureProvider](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLDictionaryFeatureProvider/initWithCoder:
func NewDictionaryFeatureProviderWithCoder(coder objectivec.IObject) MLDictionaryFeatureProvider {
	instance := getMLDictionaryFeatureProviderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLDictionaryFeatureProviderFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLDictionaryFeatureProvider/initWithFeatureProvider:
func NewDictionaryFeatureProviderWithFeatureProvider(provider objectivec.IObject) MLDictionaryFeatureProvider {
	instance := getMLDictionaryFeatureProviderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFeatureProvider:"), provider)
	return MLDictionaryFeatureProviderFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLDictionaryFeatureProvider/initWithFeatureProvider:featureNames:
func NewDictionaryFeatureProviderWithFeatureProviderFeatureNames(provider objectivec.IObject, names objectivec.IObject) MLDictionaryFeatureProvider {
	instance := getMLDictionaryFeatureProviderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFeatureProvider:featureNames:"), provider, names)
	return MLDictionaryFeatureProviderFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLDictionaryFeatureProvider/initWithFeatureValueDictionary:
func NewDictionaryFeatureProviderWithFeatureValueDictionary(dictionary objectivec.IObject) MLDictionaryFeatureProvider {
	instance := getMLDictionaryFeatureProviderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFeatureValueDictionary:"), dictionary)
	return MLDictionaryFeatureProviderFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLDictionaryFeatureProvider/initWithCoder:
func (d MLDictionaryFeatureProvider) InitWithCoder(coder foundation.INSCoder) MLDictionaryFeatureProvider {
	rv := objc.Send[MLDictionaryFeatureProvider](d.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLDictionaryFeatureProvider/initWithFeatureProvider:
func (d MLDictionaryFeatureProvider) InitWithFeatureProvider(provider objectivec.IObject) MLDictionaryFeatureProvider {
	rv := objc.Send[MLDictionaryFeatureProvider](d.ID, objc.Sel("initWithFeatureProvider:"), provider)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLDictionaryFeatureProvider/initWithFeatureProvider:featureNames:
func (d MLDictionaryFeatureProvider) InitWithFeatureProviderFeatureNames(provider objectivec.IObject, names objectivec.IObject) MLDictionaryFeatureProvider {
	rv := objc.Send[MLDictionaryFeatureProvider](d.ID, objc.Sel("initWithFeatureProvider:featureNames:"), provider, names)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLDictionaryFeatureProvider/initWithFeatureValueDictionary:
func (d MLDictionaryFeatureProvider) InitWithFeatureValueDictionary(dictionary objectivec.IObject) MLDictionaryFeatureProvider {
	rv := objc.Send[MLDictionaryFeatureProvider](d.ID, objc.Sel("initWithFeatureValueDictionary:"), dictionary)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLDictionaryFeatureProvider/supportsSecureCoding
func (_MLDictionaryFeatureProviderClass MLDictionaryFeatureProviderClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_MLDictionaryFeatureProviderClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLDictionaryFeatureProvider/cachedFeatureNames
func (d MLDictionaryFeatureProvider) CachedFeatureNames() foundation.INSSet {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("cachedFeatureNames"))
	return foundation.NSSetFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLDictionaryFeatureProvider/dictionary
func (d MLDictionaryFeatureProvider) Dictionary() foundation.INSDictionary {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("dictionary"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (d MLDictionaryFeatureProvider) SetDictionary(value foundation.INSDictionary) {
	objc.Send[struct{}](d.ID, objc.Sel("setDictionary:"), value)
}

