// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"context"
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLModelAsset] class.
var (
	_MLModelAssetClass     MLModelAssetClass
	_MLModelAssetClassOnce sync.Once
)

func getMLModelAssetClass() MLModelAssetClass {
	_MLModelAssetClassOnce.Do(func() {
		_MLModelAssetClass = MLModelAssetClass{class: objc.GetClass("MLModelAsset")}
	})
	return _MLModelAssetClass
}

// GetMLModelAssetClass returns the class object for MLModelAsset.
func GetMLModelAssetClass() MLModelAssetClass {
	return getMLModelAssetClass()
}

type MLModelAssetClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLModelAssetClass) Alloc() MLModelAsset {
	rv := objc.Send[MLModelAsset](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// An abstraction of a compiled Core ML model asset.
//
// # Overview
// 
// [MLModelAsset] provides a unified interface by abstracting the compiled
// model representations for `XCUIElementTypeMlmodelc` files and in-memory
// representations.
// 
// To use an in-memory model, create an [MLModelAsset] with an in-memory model
// specification, then call [LoadModelAssetConfigurationCompletionHandler].
//
// # Getting function names
//
//   - [MLModelAsset.FunctionNamesWithCompletionHandler]: The list of function names in the model asset.
//
// # Getting the model description
//
//   - [MLModelAsset.ModelDescriptionWithCompletionHandler]: The default model descripton.
//   - [MLModelAsset.ModelDescriptionOfFunctionNamedCompletionHandler]: The model descripton for a specified function.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset
type MLModelAsset struct {
	objectivec.Object
}

// MLModelAssetFromID constructs a [MLModelAsset] from an objc.ID.
//
// An abstraction of a compiled Core ML model asset.
func MLModelAssetFromID(id objc.ID) MLModelAsset {
	return MLModelAsset{objectivec.Object{id}}
}
// NOTE: MLModelAsset adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [MLModelAsset] class.
//
// # Getting function names
//
//   - [IMLModelAsset.FunctionNamesWithCompletionHandler]: The list of function names in the model asset.
//
// # Getting the model description
//
//   - [IMLModelAsset.ModelDescriptionWithCompletionHandler]: The default model descripton.
//   - [IMLModelAsset.ModelDescriptionOfFunctionNamedCompletionHandler]: The model descripton for a specified function.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset
type IMLModelAsset interface {
	objectivec.IObject

	// Topic: Getting function names

	// The list of function names in the model asset.
	FunctionNamesWithCompletionHandler(handler ErrorHandler)

	// Topic: Getting the model description

	// The default model descripton.
	ModelDescriptionWithCompletionHandler(handler MLModelDescriptionErrorHandler)
	// The model descripton for a specified function.
	ModelDescriptionOfFunctionNamedCompletionHandler(functionName string, handler MLModelDescriptionErrorHandler)
}





// Init initializes the instance.
func (m MLModelAsset) Init() MLModelAsset {
	rv := objc.Send[MLModelAsset](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLModelAsset) Autorelease() MLModelAsset {
	rv := objc.Send[MLModelAsset](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLModelAsset creates a new MLModelAsset instance.
func NewMLModelAsset() MLModelAsset {
	class := getMLModelAssetClass()
	rv := objc.Send[MLModelAsset](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Construct a model asset from an ML Program specification by replacing blob
// file references with corresponding in-memory blobs.
//
// blobMapping: A dictionary with blob URL as the key and blob data as the value.
//
// # Discussion
// 
// An ML Program may use [BlobFileValue] syntax, which stores the blob data in
// external files and refers them by URL. This factory method enables
// in-memory workflow for such models by using the specified in-memory blob
// data in place of the external files.
// 
// The format of in-memory blobs must be the same as the external files. The
// dictionary must contain all the reference URLs used in the specification.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset/init(specification:blobMapping:)
func NewModelAssetWithSpecificationDataBlobMappingError(specificationData foundation.INSData, blobMapping foundation.INSDictionary) (MLModelAsset, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getMLModelAssetClass().class), objc.Sel("modelAssetWithSpecificationData:blobMapping:error:"), specificationData, blobMapping, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLModelAssetFromID(rv), foundation.NSErrorFrom(errorPtr)
	}
	return MLModelAssetFromID(rv), nil
}


// Creates a model asset from an in-memory model specification.
//
// specificationData: The contents of a `XCUIElementTypeMlmodel` as a data blob.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset/init(specification:)
func NewModelAssetWithSpecificationDataError(specificationData foundation.INSData) (MLModelAsset, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getMLModelAssetClass().class), objc.Sel("modelAssetWithSpecificationData:error:"), specificationData, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLModelAssetFromID(rv), foundation.NSErrorFrom(errorPtr)
	}
	return MLModelAssetFromID(rv), nil
}


// Constructs a ModelAsset from a compiled model URL.
//
// compiledModelURL: Location on the disk where the model asset is present.
//
// # Return Value
// 
// A model asset or nil if there is an error.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset/init(url:)
func NewModelAssetWithURLError(compiledModelURL foundation.INSURL) (MLModelAsset, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getMLModelAssetClass().class), objc.Sel("modelAssetWithURL:error:"), compiledModelURL, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLModelAssetFromID(rv), foundation.NSErrorFrom(errorPtr)
	}
	return MLModelAssetFromID(rv), nil
}







// The list of function names in the model asset.
//
// # Discussion
// 
// Some model types (e.g. ML Program) supports multiple functions. Use this
// method to query the function names.
// 
// The method vends the empty array when the model doesn’t use the
// multi-function configuration.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset/functionNames(completionHandler:)
func (m MLModelAsset) FunctionNamesWithCompletionHandler(handler ErrorHandler) {
		_block0, _cleanup0 := NewErrorBlock(handler)
	defer _cleanup0()
		objc.Send[objc.ID](m.ID, objc.Sel("functionNamesWithCompletionHandler:"), _block0)
}

// The default model descripton.
//
// # Discussion
// 
// Use this method to get the description of the model such as the feature
// descriptions, the model author, and other metadata.
// 
// For the multi-function model asset, this method vends the description for
// the default function. Use `modelDescription()` to get the model description
// of other functions.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset/modelDescription(completionHandler:)
func (m MLModelAsset) ModelDescriptionWithCompletionHandler(handler MLModelDescriptionErrorHandler) {
		_block0, _cleanup0 := NewMLModelDescriptionErrorBlock(handler)
	defer _cleanup0()
		objc.Send[objc.ID](m.ID, objc.Sel("modelDescriptionWithCompletionHandler:"), _block0)
}

// The model descripton for a specified function.
//
// # Discussion
// 
// Use this method to get the description of the model such as the feature
// descriptions, the model author, and other metadata.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelAsset/modelDescription(ofFunctionNamed:completionHandler:)
func (m MLModelAsset) ModelDescriptionOfFunctionNamedCompletionHandler(functionName string, handler MLModelDescriptionErrorHandler) {
		_block1, _cleanup1 := NewMLModelDescriptionErrorBlock(handler)
	defer _cleanup1()
		objc.Send[objc.ID](m.ID, objc.Sel("modelDescriptionOfFunctionNamed:completionHandler:"), objc.String(functionName), _block1)
}



























// FunctionNames is a synchronous wrapper around [MLModelAsset.FunctionNamesWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLModelAsset) FunctionNames(ctx context.Context) error {
	done := make(chan error, 1)
	m.FunctionNamesWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// ModelDescription is a synchronous wrapper around [MLModelAsset.ModelDescriptionWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLModelAsset) ModelDescription(ctx context.Context) (*MLModelDescription, error) {
	type result struct {
		val *MLModelDescription
		err error
	}
	done := make(chan result, 1)
	m.ModelDescriptionWithCompletionHandler(func(val *MLModelDescription, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// ModelDescriptionOfFunctionNamed is a synchronous wrapper around [MLModelAsset.ModelDescriptionOfFunctionNamedCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (m MLModelAsset) ModelDescriptionOfFunctionNamed(ctx context.Context, functionName string) (*MLModelDescription, error) {
	type result struct {
		val *MLModelDescription
		err error
	}
	done := make(chan result, 1)
	m.ModelDescriptionOfFunctionNamedCompletionHandler(functionName, func(val *MLModelDescription, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}






