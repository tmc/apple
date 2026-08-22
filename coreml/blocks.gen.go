// Code generated from Apple documentation. DO NOT EDIT.

package coreml

import (
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// MLComputePlanErrorHandler handles When the compute plan is constructed successfully or unsuccessfully, the completion handler is invoked with a valid MLComputePlan instance or NSError object.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [MLComputePlan.LoadContentsOfURLConfigurationCompletionHandler]
//   - [MLComputePlan.LoadModelAssetConfigurationCompletionHandler]
type MLComputePlanErrorHandler = func(*MLComputePlan, error)

// NewMLComputePlanErrorBlock wraps a Go [MLComputePlanErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MLComputePlan.LoadContentsOfURLConfigurationCompletionHandler]
//   - [MLComputePlan.LoadModelAssetConfigurationCompletionHandler]
func NewMLComputePlanErrorBlock(handler MLComputePlanErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *MLComputePlan
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := MLComputePlanFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// MLFeatureProviderErrorHandler handles The callback the system invokes when it completes the prediction.
//   - output: A feature provider that contains the outputs of the prediction.
//   - error: If an error occurs, an error object that describes the error; otherwise, `nil`.
//
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [MLModel.PredictionFromFeaturesCompletionHandler]
//   - [MLModel.PredictionFromFeaturesOptionsCompletionHandler]
//   - [MLModel.PredictionFromFeaturesUsingStateOptionsCompletionHandler]
type MLFeatureProviderErrorHandler = func(MLFeatureProvider, error)

// NewMLFeatureProviderErrorBlock wraps a Go [MLFeatureProviderErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MLModel.PredictionFromFeaturesCompletionHandler]
//   - [MLModel.PredictionFromFeaturesOptionsCompletionHandler]
//   - [MLModel.PredictionFromFeaturesUsingStateOptionsCompletionHandler]
func NewMLFeatureProviderErrorBlock(handler MLFeatureProviderErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result MLFeatureProvider
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			result = MLFeatureProviderObjectFromID(resultID)
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// MLModelDescriptionErrorHandler is the signature for a completion handler block.
//
// Used by:
//   - [MLModelAsset.ModelDescriptionOfFunctionNamedCompletionHandler]
//   - [MLModelAsset.ModelDescriptionWithCompletionHandler]
type MLModelDescriptionErrorHandler = func(*MLModelDescription, error)

// NewMLModelDescriptionErrorBlock wraps a Go [MLModelDescriptionErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MLModelAsset.ModelDescriptionOfFunctionNamedCompletionHandler]
//   - [MLModelAsset.ModelDescriptionWithCompletionHandler]
func NewMLModelDescriptionErrorBlock(handler MLModelDescriptionErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *MLModelDescription
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := MLModelDescriptionFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// MLModelErrorHandler handles The completion handler invoked when the load completes.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [MLModel.LoadContentsOfURLConfigurationCompletionHandler]
//   - [MLModel.LoadModelAssetConfigurationCompletionHandler]
type MLModelErrorHandler = func(*MLModel, error)

// NewMLModelErrorBlock wraps a Go [MLModelErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MLModel.LoadContentsOfURLConfigurationCompletionHandler]
//   - [MLModel.LoadModelAssetConfigurationCompletionHandler]
func NewMLModelErrorBlock(handler MLModelErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *MLModel
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := MLModelFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// MLModelStructureErrorHandler handles When the model structure is constructed successfully or unsuccessfully, the completion handler is invoked with a valid MLModelStructure instance or NSError object.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [MLModelStructure.LoadContentsOfURLCompletionHandler]
//   - [MLModelStructure.LoadModelAssetCompletionHandler]
type MLModelStructureErrorHandler = func(*MLModelStructure, error)

// NewMLModelStructureErrorBlock wraps a Go [MLModelStructureErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MLModelStructure.LoadContentsOfURLCompletionHandler]
//   - [MLModelStructure.LoadModelAssetCompletionHandler]
func NewMLModelStructureErrorBlock(handler MLModelStructureErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *MLModelStructure
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := MLModelStructureFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// MLMultiArrayHandler handles Block to access the state buffer through [MLMultiArray].
//
// Used by:
//   - [MLState.GetMultiArrayForStateNamedHandler]
type MLMultiArrayHandler = func(*MLMultiArray)

// NewMLMultiArrayBlock wraps a Go [MLMultiArrayHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MLState.GetMultiArrayForStateNamedHandler]
func NewMLMultiArrayBlock(handler MLMultiArrayHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *MLMultiArray
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := MLMultiArrayFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// MLUpdateContextHandler handles The closure an update task uses to notify your app.
//
// Used by:
//   - [MLUpdateProgressHandlers.InitForEventsProgressHandlerCompletionHandler]
//   - [MLUpdateTask.UpdateTaskForModelAtURLTrainingDataCompletionHandlerError]
//   - [MLUpdateTask.UpdateTaskForModelAtURLTrainingDataConfigurationCompletionHandlerError]
type MLUpdateContextHandler = func(*MLUpdateContext)

// NewMLUpdateContextBlock wraps a Go [MLUpdateContextHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MLUpdateProgressHandlers.InitForEventsProgressHandlerCompletionHandler]
//   - [MLUpdateTask.UpdateTaskForModelAtURLTrainingDataCompletionHandlerError]
//   - [MLUpdateTask.UpdateTaskForModelAtURLTrainingDataConfigurationCompletionHandlerError]
func NewMLUpdateContextBlock(handler MLUpdateContextHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *MLUpdateContext
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := MLUpdateContextFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// URLErrorHandler is the signature for a completion handler block.
//
// Used by:
//   - [MLModel.CompileModelAtURLCompletionHandler]
type URLErrorHandler = func(*foundation.NSURL, error)

// NewURLErrorBlock wraps a Go [URLErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MLModel.CompileModelAtURLCompletionHandler]
func NewURLErrorBlock(handler URLErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *foundation.NSURL
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := foundation.NSURLFromID(resultID)
			result = &v
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// UnsafePointerHandler handles In Swift, a closure the multiarray calls in its deinitializer.
//
// Used by:
//   - [MLMultiArray.InitWithDataPointerShapeDataTypeStridesDeallocatorError]
type UnsafePointerHandler = func(unsafe.Pointer)

// NewUnsafePointerBlock wraps a Go [UnsafePointerHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MLMultiArray.InitWithDataPointerShapeDataTypeStridesDeallocatorError]
func NewUnsafePointerBlock(handler UnsafePointerHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal unsafe.Pointer) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// UnsafePointerIntHandler handles The block to receive the buffer pointer and its size in bytes.
//
// Used by:
//   - [MLMultiArray.GetBytesWithHandler]
type UnsafePointerIntHandler = func(unsafe.Pointer, int)

// NewUnsafePointerIntBlock wraps a Go [UnsafePointerIntHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MLMultiArray.GetBytesWithHandler]
func NewUnsafePointerIntBlock(handler UnsafePointerIntHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive unsafe.Pointer, extra0 int) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// UnsafePointerIntNumberArrayHandler handles The block to receive the buffer pointer, its size in bytes, and strides.
//   - bytes: The pointer to the buffer.
//   - size: The size of the buffer.
//   - strides: The strides of the buffer in scalars. Note that this may be different from `strides`’s value prior to this method invocation.
//
// Used by:
//   - [MLMultiArray.GetMutableBytesWithHandler]
type UnsafePointerIntNumberArrayHandler = func(unsafe.Pointer, int, *foundation.NSArray)

// NewUnsafePointerIntNumberArrayBlock wraps a Go [UnsafePointerIntNumberArrayHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MLMultiArray.GetMutableBytesWithHandler]
func NewUnsafePointerIntNumberArrayBlock(handler UnsafePointerIntNumberArrayHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive unsafe.Pointer, extra0 int, extra1ID objc.ID) {
		var extra1 *foundation.NSArray
		if extra1ID != 0 {
			objc.Send[objc.ID](extra1ID, objc.Sel("retain"))
			v := foundation.NSArrayFromID(extra1ID)
			extra1 = &v
		}
		handler(primitive, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

// stringArrayErrorHandler is the signature for a completion handler block.
//
// Used by:
//   - [MLModelAsset.FunctionNamesWithCompletionHandler]
type stringArrayErrorHandler = func(*[]string, error)

// NewstringArrayErrorBlock wraps a Go [stringArrayErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MLModelAsset.FunctionNamesWithCompletionHandler]
func NewstringArrayErrorBlock(handler stringArrayErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result *[]string
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			obj := foundation.NSArrayFromID(resultID)
			count := obj.Count()
			res := make([]string, count)
			for i := uint(0); i < count; i++ {
				item := obj.ObjectAtIndex(i)
				res[i] = objc.IDToString(item.GetID())
			}
			result = &res
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}
