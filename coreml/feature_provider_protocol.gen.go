// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// An interface that represents a collection of values for either a model’s input or its output.
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureProvider
type MLFeatureProvider interface {
	objectivec.IObject

	// Accesses the feature value given the feature’s name.
	//
	// See: https://developer.apple.com/documentation/CoreML/MLFeatureProvider/featureValue(for:)
	FeatureValueForName(featureName string) IMLFeatureValue

	// The set of valid feature names.
	//
	// See: https://developer.apple.com/documentation/CoreML/MLFeatureProvider/featureNames
	FeatureNames() foundation.INSSet
}

// MLFeatureProviderObject wraps an existing Objective-C object that conforms to the MLFeatureProvider protocol.
type MLFeatureProviderObject struct {
	objectivec.Object
}
func (o MLFeatureProviderObject) BaseObject() objectivec.Object {
	return o.Object
}

// MLFeatureProviderObjectFromID constructs a [MLFeatureProviderObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MLFeatureProviderObjectFromID(id objc.ID) MLFeatureProviderObject {
	return MLFeatureProviderObject{
		Object: objectivec.ObjectFromID(id),
	}
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
func (o MLFeatureProviderObject) FeatureValueForName(featureName string) IMLFeatureValue {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("featureValueForName:"), objc.String(featureName))
	return MLFeatureValueFromID(rv)
	}
// The set of valid feature names.
//
// See: https://developer.apple.com/documentation/CoreML/MLFeatureProvider/featureNames
func (o MLFeatureProviderObject) FeatureNames() foundation.INSSet {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("featureNames"))
	return foundation.NSSetFromID(rv)
	}

