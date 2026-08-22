// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
	rv := objc.SendIfResponds[MLDictionaryFeatureProvider](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLDictionaryFeatureProvider.CachedFeatureNames]
//   - [MLDictionaryFeatureProvider.CountByEnumeratingWithStateObjectsCount]
//   - [MLDictionaryFeatureProvider.InitWithFeatureProvider]
//   - [MLDictionaryFeatureProvider.InitWithFeatureProviderFeatureNames]
//   - [MLDictionaryFeatureProvider.InitWithFeatureValueDictionary]
//   - [MLDictionaryFeatureProvider.Dictionary]
//   - [MLDictionaryFeatureProvider.SetDictionary]
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
//   - [IMLDictionaryFeatureProvider.CountByEnumeratingWithStateObjectsCount]
//   - [IMLDictionaryFeatureProvider.InitWithFeatureProvider]
//   - [IMLDictionaryFeatureProvider.InitWithFeatureProviderFeatureNames]
//   - [IMLDictionaryFeatureProvider.InitWithFeatureValueDictionary]
//   - [IMLDictionaryFeatureProvider.Dictionary]
//   - [IMLDictionaryFeatureProvider.SetDictionary]
type IMLDictionaryFeatureProvider interface {
	objectivec.IObject

	// Topic: Methods

	CachedFeatureNames() foundation.INSSet
	CountByEnumeratingWithStateObjectsCount(state unsafe.Pointer, objects []objectivec.IObject, count uint64) uint64
	InitWithFeatureProvider(provider objectivec.IObject) MLDictionaryFeatureProvider
	InitWithFeatureProviderFeatureNames(provider objectivec.IObject, names objectivec.IObject) MLDictionaryFeatureProvider
	InitWithFeatureValueDictionary(dictionary objectivec.IObject) MLDictionaryFeatureProvider
	Dictionary() foundation.INSDictionary
	SetDictionary(value foundation.INSDictionary)
}

// Init initializes the instance.
func (m MLDictionaryFeatureProvider) Init() MLDictionaryFeatureProvider {
	rv := objc.SendIfResponds[MLDictionaryFeatureProvider](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLDictionaryFeatureProvider) Autorelease() MLDictionaryFeatureProvider {
	rv := objc.SendIfResponds[MLDictionaryFeatureProvider](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLDictionaryFeatureProvider creates a new MLDictionaryFeatureProvider instance.
func NewMLDictionaryFeatureProvider() MLDictionaryFeatureProvider {
	class := getMLDictionaryFeatureProviderClass()
	rv := objc.SendIfResponds[MLDictionaryFeatureProvider](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewDictionaryFeatureProviderWithFeatureProvider(provider objectivec.IObject) MLDictionaryFeatureProvider {
	instance := getMLDictionaryFeatureProviderClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithFeatureProvider:"), provider)
	return MLDictionaryFeatureProviderFromID(rv)
}

func NewDictionaryFeatureProviderWithFeatureProviderFeatureNames(provider objectivec.IObject, names objectivec.IObject) MLDictionaryFeatureProvider {
	instance := getMLDictionaryFeatureProviderClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithFeatureProvider:featureNames:"), provider, names)
	return MLDictionaryFeatureProviderFromID(rv)
}

func NewDictionaryFeatureProviderWithFeatureValueDictionary(dictionary objectivec.IObject) MLDictionaryFeatureProvider {
	instance := getMLDictionaryFeatureProviderClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithFeatureValueDictionary:"), dictionary)
	return MLDictionaryFeatureProviderFromID(rv)
}

func (m MLDictionaryFeatureProvider) CountByEnumeratingWithStateObjectsCount(state unsafe.Pointer, objects []objectivec.IObject, count uint64) uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("countByEnumeratingWithState:objects:count:"), objc.CArray(state), objc.CArray(objects), count)
	return rv
}
func (m MLDictionaryFeatureProvider) InitWithFeatureProvider(provider objectivec.IObject) MLDictionaryFeatureProvider {
	rv := objc.SendIfResponds[MLDictionaryFeatureProvider](m.ID, objc.Sel("initWithFeatureProvider:"), provider)
	return rv
}
func (m MLDictionaryFeatureProvider) InitWithFeatureProviderFeatureNames(provider objectivec.IObject, names objectivec.IObject) MLDictionaryFeatureProvider {
	rv := objc.SendIfResponds[MLDictionaryFeatureProvider](m.ID, objc.Sel("initWithFeatureProvider:featureNames:"), provider, names)
	return rv
}
func (m MLDictionaryFeatureProvider) InitWithFeatureValueDictionary(dictionary objectivec.IObject) MLDictionaryFeatureProvider {
	rv := objc.SendIfResponds[MLDictionaryFeatureProvider](m.ID, objc.Sel("initWithFeatureValueDictionary:"), dictionary)
	return rv
}

func (_MLDictionaryFeatureProviderClass MLDictionaryFeatureProviderClass) SupportsSecureCoding() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_MLDictionaryFeatureProviderClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

func (m MLDictionaryFeatureProvider) CachedFeatureNames() foundation.INSSet {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("cachedFeatureNames"))
	return foundation.NSSetFromID(objc.ID(rv))
}
func (m MLDictionaryFeatureProvider) Dictionary() foundation.INSDictionary {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("dictionary"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLDictionaryFeatureProvider) SetDictionary(value foundation.INSDictionary) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setDictionary:"), value)
}
