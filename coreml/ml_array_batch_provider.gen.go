// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLArrayBatchProvider] class.
var (
	_MLArrayBatchProviderClass     MLArrayBatchProviderClass
	_MLArrayBatchProviderClassOnce sync.Once
)

func getMLArrayBatchProviderClass() MLArrayBatchProviderClass {
	_MLArrayBatchProviderClassOnce.Do(func() {
		_MLArrayBatchProviderClass = MLArrayBatchProviderClass{class: objc.GetClass("MLArrayBatchProvider")}
	})
	return _MLArrayBatchProviderClass
}

// GetMLArrayBatchProviderClass returns the class object for MLArrayBatchProvider.
func GetMLArrayBatchProviderClass() MLArrayBatchProviderClass {
	return getMLArrayBatchProviderClass()
}

type MLArrayBatchProviderClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLArrayBatchProviderClass) Alloc() MLArrayBatchProvider {
	rv := objc.Send[MLArrayBatchProvider](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// A convenience wrapper for batches of feature providers.
//
// # Overview
// 
// This batch provider supports an array of feature providers or a dictionary
// of arrays of feature values.
//
// # Creating a batch provider
//
//   - [MLArrayBatchProvider.InitWithFeatureProviderArray]: Creates the batch provider based on the array of feature providers.
//   - [MLArrayBatchProvider.InitWithDictionaryError]: Creates a batch provider based on feature names and their associated arrays of data.
//
// # Accessing the feature providers
//
//   - [MLArrayBatchProvider.Array]: The array of feature providers.
//
// See: https://developer.apple.com/documentation/CoreML/MLArrayBatchProvider
type MLArrayBatchProvider struct {
	objectivec.Object
}

// MLArrayBatchProviderFromID constructs a [MLArrayBatchProvider] from an objc.ID.
//
// A convenience wrapper for batches of feature providers.
func MLArrayBatchProviderFromID(id objc.ID) MLArrayBatchProvider {
	return MLArrayBatchProvider{objectivec.Object{id}}
}
// NOTE: MLArrayBatchProvider adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [MLArrayBatchProvider] class.
//
// # Creating a batch provider
//
//   - [IMLArrayBatchProvider.InitWithFeatureProviderArray]: Creates the batch provider based on the array of feature providers.
//   - [IMLArrayBatchProvider.InitWithDictionaryError]: Creates a batch provider based on feature names and their associated arrays of data.
//
// # Accessing the feature providers
//
//   - [IMLArrayBatchProvider.Array]: The array of feature providers.
//
// See: https://developer.apple.com/documentation/CoreML/MLArrayBatchProvider
type IMLArrayBatchProvider interface {
	objectivec.IObject
	MLBatchProvider

	// Topic: Creating a batch provider

	// Creates the batch provider based on the array of feature providers.
	InitWithFeatureProviderArray(array []objectivec.IObject) MLArrayBatchProvider
	// Creates a batch provider based on feature names and their associated arrays of data.
	InitWithDictionaryError(dictionary foundation.INSDictionary) (MLArrayBatchProvider, error)

	// Topic: Accessing the feature providers

	// The array of feature providers.
	Array() []objectivec.IObject
}





// Init initializes the instance.
func (a MLArrayBatchProvider) Init() MLArrayBatchProvider {
	rv := objc.Send[MLArrayBatchProvider](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a MLArrayBatchProvider) Autorelease() MLArrayBatchProvider {
	rv := objc.Send[MLArrayBatchProvider](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLArrayBatchProvider creates a new MLArrayBatchProvider instance.
func NewMLArrayBatchProvider() MLArrayBatchProvider {
	class := getMLArrayBatchProviderClass()
	rv := objc.Send[MLArrayBatchProvider](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Creates a batch provider based on feature names and their associated arrays
// of data.
//
// dictionary: A dictionary which maps feature names to an array of values. The error case
// occurs when all the arrays do not have the same length or the values in an
// aray are not expressible as an [MLFeatureValue].
//
// # Discussion
// 
// This initializer is convenient when the data are available as individual
// arrays.
//
// See: https://developer.apple.com/documentation/CoreML/MLArrayBatchProvider/init(dictionary:)
func NewArrayBatchProviderWithDictionaryError(dictionary foundation.INSDictionary) (MLArrayBatchProvider, error) {
	var errorPtr objc.ID
	instance := getMLArrayBatchProviderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDictionary:error:"), dictionary, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLArrayBatchProviderFromID(rv), foundation.NSErrorFrom(errorPtr)
	}
	return MLArrayBatchProviderFromID(rv), nil
}


// Creates the batch provider based on the array of feature providers.
//
// array: The array of feature providers for the batch.
//
// See: https://developer.apple.com/documentation/CoreML/MLArrayBatchProvider/init(array:)
func NewArrayBatchProviderWithFeatureProviderArray(array []objectivec.IObject) MLArrayBatchProvider {
	instance := getMLArrayBatchProviderClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFeatureProviderArray:"), objectivec.IObjectSliceToNSArray(array))
	return MLArrayBatchProviderFromID(rv)
}







// Creates the batch provider based on the array of feature providers.
//
// array: The array of feature providers for the batch.
//
// See: https://developer.apple.com/documentation/CoreML/MLArrayBatchProvider/init(array:)
func (a MLArrayBatchProvider) InitWithFeatureProviderArray(array []objectivec.IObject) MLArrayBatchProvider {
	rv := objc.Send[MLArrayBatchProvider](a.ID, objc.Sel("initWithFeatureProviderArray:"), objectivec.IObjectSliceToNSArray(array))
	return rv
}

// Creates a batch provider based on feature names and their associated arrays
// of data.
//
// dictionary: A dictionary which maps feature names to an array of values. The error case
// occurs when all the arrays do not have the same length or the values in an
// aray are not expressible as an [MLFeatureValue].
//
// # Discussion
// 
// This initializer is convenient when the data are available as individual
// arrays.
//
// See: https://developer.apple.com/documentation/CoreML/MLArrayBatchProvider/init(dictionary:)
func (a MLArrayBatchProvider) InitWithDictionaryError(dictionary foundation.INSDictionary) (MLArrayBatchProvider, error) {
			var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("initWithDictionary:error:"), dictionary, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLArrayBatchProvider{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLArrayBatchProviderFromID(rv), nil

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
func (a MLArrayBatchProvider) FeaturesAtIndex(index int) MLFeatureProvider {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("featuresAtIndex:"), index)
	return MLFeatureProviderObjectFromID(rv)
}












// The array of feature providers.
//
// See: https://developer.apple.com/documentation/CoreML/MLArrayBatchProvider/array
func (a MLArrayBatchProvider) Array() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("array"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}



// The number of feature providers in this batch.
//
// See: https://developer.apple.com/documentation/CoreML/MLBatchProvider/count
func (a MLArrayBatchProvider) Count() int {
	rv := objc.Send[int](a.ID, objc.Sel("count"))
	return rv
}













			// Protocol methods for MLBatchProvider
			














