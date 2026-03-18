// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// An interface that represents a collection of feature providers.
//
// See: https://developer.apple.com/documentation/CoreML/MLBatchProvider
type MLBatchProvider interface {
	objectivec.IObject

	// Returns the feature provider at the given index.
	//
	// See: https://developer.apple.com/documentation/CoreML/MLBatchProvider/features(at:)
	FeaturesAtIndex(index int) MLFeatureProvider

	// The number of feature providers in this batch.
	//
	// See: https://developer.apple.com/documentation/CoreML/MLBatchProvider/count
	Count() int
}

// MLBatchProviderObject wraps an existing Objective-C object that conforms to the MLBatchProvider protocol.
type MLBatchProviderObject struct {
	objectivec.Object
}
func (o MLBatchProviderObject) BaseObject() objectivec.Object {
	return o.Object
}

// MLBatchProviderObjectFromID constructs a [MLBatchProviderObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MLBatchProviderObjectFromID(id objc.ID) MLBatchProviderObject {
	return MLBatchProviderObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Returns the feature provider at the given index.
//
// index: The index of the desired feature provider.
//
// # Return Value
// 
// The feature provider at the given index.
//
// See: https://developer.apple.com/documentation/CoreML/MLBatchProvider/features(at:)

func (o MLBatchProviderObject) FeaturesAtIndex(index int) MLFeatureProvider {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("featuresAtIndex:"), index)
	return MLFeatureProviderObjectFromID(rv)
	}

// The number of feature providers in this batch.
//
// See: https://developer.apple.com/documentation/CoreML/MLBatchProvider/count

func (o MLBatchProviderObject) Count() int {
	
	rv := objc.Send[int](o.ID, objc.Sel("count"))
	return rv
	}

