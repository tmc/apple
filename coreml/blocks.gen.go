// Code generated from Apple documentation. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// ErrorHandler handles The closures the task calls during the update process.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [MLModelAsset.FunctionNamesWithCompletionHandler]
//   - [MLUpdateTask.UpdateTaskForModelAtURLTrainingDataConfigurationProgressHandlersError]
//   - [MLUpdateTask.UpdateTaskForModelAtURLTrainingDataProgressHandlersError]
type ErrorHandler = func(error)

// NewErrorBlock wraps a Go [ErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MLModelAsset.FunctionNamesWithCompletionHandler]
//   - [MLUpdateTask.UpdateTaskForModelAtURLTrainingDataConfigurationProgressHandlersError]
//   - [MLUpdateTask.UpdateTaskForModelAtURLTrainingDataProgressHandlersError]
func NewErrorBlock(handler ErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, errID objc.ID) {
		handler(foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

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
//   - [MLModel.LoadModelAssetConfigurationCompletionHandler]
type MLModelErrorHandler = func(*MLModel, error)

// NewMLModelErrorBlock wraps a Go [MLModelErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MLModel.LoadModelAssetConfigurationCompletionHandler]
func NewMLModelErrorBlock(handler MLModelErrorHandler) (objc.ID, func()) {
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

// VoidHandler handles In Swift, a closure the multiarray calls in its deinitializer.
//
// Used by:
//   - [MLMultiArray.InitWithDataPointerShapeDataTypeStridesDeallocatorError]
type VoidHandler = func()

// NewVoidBlock wraps a Go [VoidHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MLMultiArray.InitWithDataPointerShapeDataTypeStridesDeallocatorError]
func NewVoidBlock(handler VoidHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}

