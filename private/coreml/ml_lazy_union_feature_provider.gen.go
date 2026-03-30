// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLLazyUnionFeatureProvider] class.
var (
	_MLLazyUnionFeatureProviderClass     MLLazyUnionFeatureProviderClass
	_MLLazyUnionFeatureProviderClassOnce sync.Once
)

func getMLLazyUnionFeatureProviderClass() MLLazyUnionFeatureProviderClass {
	_MLLazyUnionFeatureProviderClassOnce.Do(func() {
		_MLLazyUnionFeatureProviderClass = MLLazyUnionFeatureProviderClass{class: objc.GetClass("MLLazyUnionFeatureProvider")}
	})
	return _MLLazyUnionFeatureProviderClass
}

// GetMLLazyUnionFeatureProviderClass returns the class object for MLLazyUnionFeatureProvider.
func GetMLLazyUnionFeatureProviderClass() MLLazyUnionFeatureProviderClass {
	return getMLLazyUnionFeatureProviderClass()
}

type MLLazyUnionFeatureProviderClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLLazyUnionFeatureProviderClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLLazyUnionFeatureProviderClass) Alloc() MLLazyUnionFeatureProvider {
	rv := objc.Send[MLLazyUnionFeatureProvider](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLLazyUnionFeatureProvider.FeatureNames]
//   - [MLLazyUnionFeatureProvider.FeatureValueForName]
//   - [MLLazyUnionFeatureProvider.First]
//   - [MLLazyUnionFeatureProvider.SetFirst]
//   - [MLLazyUnionFeatureProvider.Second]
//   - [MLLazyUnionFeatureProvider.SetSecond]
//   - [MLLazyUnionFeatureProvider.UnionFeatureProvider]
//   - [MLLazyUnionFeatureProvider.InitWithFeaturesFromAddedToFeaturesFrom]
//
// See: https://developer.apple.com/documentation/CoreML/MLLazyUnionFeatureProvider
type MLLazyUnionFeatureProvider struct {
	objectivec.Object
}

// MLLazyUnionFeatureProviderFromID constructs a [MLLazyUnionFeatureProvider] from an objc.ID.
func MLLazyUnionFeatureProviderFromID(id objc.ID) MLLazyUnionFeatureProvider {
	return MLLazyUnionFeatureProvider{objectivec.Object{ID: id}}
}

// Ensure MLLazyUnionFeatureProvider implements IMLLazyUnionFeatureProvider.
var _ IMLLazyUnionFeatureProvider = MLLazyUnionFeatureProvider{}

// An interface definition for the [MLLazyUnionFeatureProvider] class.
//
// # Methods
//
//   - [IMLLazyUnionFeatureProvider.FeatureNames]
//   - [IMLLazyUnionFeatureProvider.FeatureValueForName]
//   - [IMLLazyUnionFeatureProvider.First]
//   - [IMLLazyUnionFeatureProvider.SetFirst]
//   - [IMLLazyUnionFeatureProvider.Second]
//   - [IMLLazyUnionFeatureProvider.SetSecond]
//   - [IMLLazyUnionFeatureProvider.UnionFeatureProvider]
//   - [IMLLazyUnionFeatureProvider.InitWithFeaturesFromAddedToFeaturesFrom]
//
// See: https://developer.apple.com/documentation/CoreML/MLLazyUnionFeatureProvider
type IMLLazyUnionFeatureProvider interface {
	objectivec.IObject

	// Topic: Methods

	FeatureNames() foundation.INSSet
	FeatureValueForName(name objectivec.IObject) objectivec.IObject
	First() objectivec.IObject
	SetFirst(value objectivec.IObject)
	Second() objectivec.IObject
	SetSecond(value objectivec.IObject)
	UnionFeatureProvider() objectivec.IObject
	InitWithFeaturesFromAddedToFeaturesFrom(from objectivec.IObject, from2 objectivec.IObject) MLLazyUnionFeatureProvider
}

// Init initializes the instance.
func (l MLLazyUnionFeatureProvider) Init() MLLazyUnionFeatureProvider {
	rv := objc.Send[MLLazyUnionFeatureProvider](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l MLLazyUnionFeatureProvider) Autorelease() MLLazyUnionFeatureProvider {
	rv := objc.Send[MLLazyUnionFeatureProvider](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLLazyUnionFeatureProvider creates a new MLLazyUnionFeatureProvider instance.
func NewMLLazyUnionFeatureProvider() MLLazyUnionFeatureProvider {
	class := getMLLazyUnionFeatureProviderClass()
	rv := objc.Send[MLLazyUnionFeatureProvider](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLLazyUnionFeatureProvider/initWithFeaturesFrom:addedToFeaturesFrom:
func NewLazyUnionFeatureProviderWithFeaturesFromAddedToFeaturesFrom(from objectivec.IObject, from2 objectivec.IObject) MLLazyUnionFeatureProvider {
	instance := getMLLazyUnionFeatureProviderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFeaturesFrom:addedToFeaturesFrom:"), from, from2)
	return MLLazyUnionFeatureProviderFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLLazyUnionFeatureProvider/featureValueForName:
func (l MLLazyUnionFeatureProvider) FeatureValueForName(name objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("featureValueForName:"), name)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLLazyUnionFeatureProvider/unionFeatureProvider
func (l MLLazyUnionFeatureProvider) UnionFeatureProvider() objectivec.IObject {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("unionFeatureProvider"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLLazyUnionFeatureProvider/initWithFeaturesFrom:addedToFeaturesFrom:
func (l MLLazyUnionFeatureProvider) InitWithFeaturesFromAddedToFeaturesFrom(from objectivec.IObject, from2 objectivec.IObject) MLLazyUnionFeatureProvider {
	rv := objc.Send[MLLazyUnionFeatureProvider](l.ID, objc.Sel("initWithFeaturesFrom:addedToFeaturesFrom:"), from, from2)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLLazyUnionFeatureProvider/featureNames
func (l MLLazyUnionFeatureProvider) FeatureNames() foundation.INSSet {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("featureNames"))
	return foundation.NSSetFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLLazyUnionFeatureProvider/first
func (l MLLazyUnionFeatureProvider) First() objectivec.IObject {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("first"))
	return objectivec.Object{ID: rv}
}
func (l MLLazyUnionFeatureProvider) SetFirst(value objectivec.IObject) {
	objc.Send[struct{}](l.ID, objc.Sel("setFirst:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLLazyUnionFeatureProvider/second
func (l MLLazyUnionFeatureProvider) Second() objectivec.IObject {
	rv := objc.Send[objc.ID](l.ID, objc.Sel("second"))
	return objectivec.Object{ID: rv}
}
func (l MLLazyUnionFeatureProvider) SetSecond(value objectivec.IObject) {
	objc.Send[struct{}](l.ID, objc.Sel("setSecond:"), value)
}
