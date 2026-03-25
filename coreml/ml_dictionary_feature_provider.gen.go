// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
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

// Alloc allocates memory for a new instance of the class.
func (mc MLDictionaryFeatureProviderClass) Alloc() MLDictionaryFeatureProvider {
	rv := objc.Send[MLDictionaryFeatureProvider](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A convenience wrapper for the given dictionary of data.
//
// # Overview
// 
// If your input data is stored in a dictionary, consider this type of
// [MLFeatureProvider] that is backed by a dictionary. It is a convenience
// interface, saving you the trouble of iterating through the dictionary to
// assign all of its values.
//
// # Creating the provider
//
//   - [MLDictionaryFeatureProvider.InitWithDictionaryError]: Creates the feature provider based on a dictionary.
//
// # Accessing the features
//
//   - [MLDictionaryFeatureProvider.ObjectForKeyedSubscript]: Subscript interface for the feature provider to pass through to the dictionary.
//   - [MLDictionaryFeatureProvider.Dictionary]: The backing dictionary.
//
// See: https://developer.apple.com/documentation/CoreML/MLDictionaryFeatureProvider
type MLDictionaryFeatureProvider struct {
	objectivec.Object
}

// MLDictionaryFeatureProviderFromID constructs a [MLDictionaryFeatureProvider] from an objc.ID.
//
// A convenience wrapper for the given dictionary of data.
func MLDictionaryFeatureProviderFromID(id objc.ID) MLDictionaryFeatureProvider {
	return MLDictionaryFeatureProvider{objectivec.Object{ID: id}}
}
// NOTE: MLDictionaryFeatureProvider adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MLDictionaryFeatureProvider] class.
//
// # Creating the provider
//
//   - [IMLDictionaryFeatureProvider.InitWithDictionaryError]: Creates the feature provider based on a dictionary.
//
// # Accessing the features
//
//   - [IMLDictionaryFeatureProvider.ObjectForKeyedSubscript]: Subscript interface for the feature provider to pass through to the dictionary.
//   - [IMLDictionaryFeatureProvider.Dictionary]: The backing dictionary.
//
// See: https://developer.apple.com/documentation/CoreML/MLDictionaryFeatureProvider
type IMLDictionaryFeatureProvider interface {
	objectivec.IObject
	MLFeatureProvider

	// Topic: Creating the provider

	// Creates the feature provider based on a dictionary.
	InitWithDictionaryError(dictionary foundation.INSDictionary) (MLDictionaryFeatureProvider, error)

	// Topic: Accessing the features

	// Subscript interface for the feature provider to pass through to the dictionary.
	ObjectForKeyedSubscript(featureName string) IMLFeatureValue
	// The backing dictionary.
	Dictionary() foundation.INSDictionary

	EncodeWithCoder(coder foundation.INSCoder)
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

// Creates the feature provider based on a dictionary.
//
// dictionary: The dictionary of feature names and feature values.
//
// See: https://developer.apple.com/documentation/CoreML/MLDictionaryFeatureProvider/init(dictionary:)
func NewDictionaryFeatureProviderWithDictionaryError(dictionary foundation.INSDictionary) (MLDictionaryFeatureProvider, error) {
	var errorPtr objc.ID
	instance := getMLDictionaryFeatureProviderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDictionary:error:"), dictionary, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLDictionaryFeatureProvider{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLDictionaryFeatureProviderFromID(rv), nil
}

// Creates the feature provider based on a dictionary.
//
// dictionary: The dictionary of feature names and feature values.
//
// See: https://developer.apple.com/documentation/CoreML/MLDictionaryFeatureProvider/init(dictionary:)
func (d MLDictionaryFeatureProvider) InitWithDictionaryError(dictionary foundation.INSDictionary) (MLDictionaryFeatureProvider, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initWithDictionary:error:"), dictionary, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLDictionaryFeatureProvider{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLDictionaryFeatureProviderFromID(rv), nil

}
// Subscript interface for the feature provider to pass through to the
// dictionary.
//
// See: https://developer.apple.com/documentation/CoreML/MLDictionaryFeatureProvider/subscript(_:)
func (d MLDictionaryFeatureProvider) ObjectForKeyedSubscript(featureName string) IMLFeatureValue {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("objectForKeyedSubscript:"), objc.String(featureName))
	return MLFeatureValueFromID(rv)
}
// Accesses the feature value given the feature’s name.
//
// featureName: The name of the feature of the desired value.
//
// # Return Value
// 
// The value of the feature, or nil if no value exists for that name.
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureProvider/featureValue(for:)
func (d MLDictionaryFeatureProvider) FeatureValueForName(featureName string) IMLFeatureValue {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("featureValueForName:"), objc.String(featureName))
	return MLFeatureValueFromID(rv)
}
func (d MLDictionaryFeatureProvider) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](d.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The backing dictionary.
//
// See: https://developer.apple.com/documentation/CoreML/MLDictionaryFeatureProvider/dictionary
func (d MLDictionaryFeatureProvider) Dictionary() foundation.INSDictionary {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("dictionary"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
// The set of valid feature names.
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureProvider/featureNames
func (d MLDictionaryFeatureProvider) FeatureNames() foundation.INSSet {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("featureNames"))
	return foundation.NSSetFromID(objc.ID(rv))
}

			// Protocol methods for MLFeatureProvider
			

