// Code generated from Apple documentation. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// ErrorHandler handles A Swift closure or an Objective-C block the method calls when it finishes creating the compute pipeline state.
//
// Used by:
//   - [MTL4Compiler.NewBinaryFunctionWithDescriptorCompilerTaskOptionsCompletionHandler]
//   - [MTL4Compiler.NewComputePipelineStateWithDescriptorCompilerTaskOptionsCompletionHandler]
//   - [MTL4Compiler.NewComputePipelineStateWithDescriptorDynamicLinkingDescriptorCompilerTaskOptionsCompletionHandler]
//   - [MTL4Compiler.NewDynamicLibraryCompletionHandler]
//   - [MTL4Compiler.NewDynamicLibraryWithURLCompletionHandler]
//   - [MTL4Compiler.NewLibraryWithDescriptorCompletionHandler]
//   - [MTL4Compiler.NewMachineLearningPipelineStateWithDescriptorCompletionHandler]
//   - [MTL4Compiler.NewRenderPipelineStateBySpecializationWithDescriptorPipelineCompletionHandler]
//   - [MTL4Compiler.NewRenderPipelineStateWithDescriptorCompilerTaskOptionsCompletionHandler]
//   - [MTL4Compiler.NewRenderPipelineStateWithDescriptorDynamicLinkingDescriptorCompilerTaskOptionsCompletionHandler]
//   - [MTLDevice.NewComputePipelineStateWithDescriptorOptionsCompletionHandler]
//   - [MTLDevice.NewComputePipelineStateWithFunctionCompletionHandler]
//   - [MTLDevice.NewComputePipelineStateWithFunctionOptionsCompletionHandler]
//   - [MTLDevice.NewLibraryWithSourceOptionsCompletionHandler]
//   - [MTLDevice.NewLibraryWithStitchedDescriptorCompletionHandler]
//   - [MTLDevice.NewRenderPipelineStateWithDescriptorCompletionHandler]
//   - [MTLDevice.NewRenderPipelineStateWithDescriptorOptionsCompletionHandler]
//   - [MTLDevice.NewRenderPipelineStateWithMeshDescriptorOptionsCompletionHandler]
//   - [MTLDevice.NewRenderPipelineStateWithTileDescriptorOptionsCompletionHandler]
type ErrorHandler = func()

// NewErrorBlock wraps a Go [ErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MTL4Compiler.NewBinaryFunctionWithDescriptorCompilerTaskOptionsCompletionHandler]
//   - [MTL4Compiler.NewComputePipelineStateWithDescriptorCompilerTaskOptionsCompletionHandler]
//   - [MTL4Compiler.NewComputePipelineStateWithDescriptorDynamicLinkingDescriptorCompilerTaskOptionsCompletionHandler]
//   - [MTL4Compiler.NewDynamicLibraryCompletionHandler]
//   - [MTL4Compiler.NewDynamicLibraryWithURLCompletionHandler]
//   - [MTL4Compiler.NewLibraryWithDescriptorCompletionHandler]
//   - [MTL4Compiler.NewMachineLearningPipelineStateWithDescriptorCompletionHandler]
//   - [MTL4Compiler.NewRenderPipelineStateBySpecializationWithDescriptorPipelineCompletionHandler]
//   - [MTL4Compiler.NewRenderPipelineStateWithDescriptorCompilerTaskOptionsCompletionHandler]
//   - [MTL4Compiler.NewRenderPipelineStateWithDescriptorDynamicLinkingDescriptorCompilerTaskOptionsCompletionHandler]
//   - [MTLDevice.NewComputePipelineStateWithDescriptorOptionsCompletionHandler]
//   - [MTLDevice.NewComputePipelineStateWithFunctionCompletionHandler]
//   - [MTLDevice.NewComputePipelineStateWithFunctionOptionsCompletionHandler]
//   - [MTLDevice.NewLibraryWithSourceOptionsCompletionHandler]
//   - [MTLDevice.NewLibraryWithStitchedDescriptorCompletionHandler]
//   - [MTLDevice.NewRenderPipelineStateWithDescriptorCompletionHandler]
//   - [MTLDevice.NewRenderPipelineStateWithDescriptorOptionsCompletionHandler]
//   - [MTLDevice.NewRenderPipelineStateWithMeshDescriptorOptionsCompletionHandler]
//   - [MTLDevice.NewRenderPipelineStateWithTileDescriptorOptionsCompletionHandler]
func NewErrorBlock(handler ErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}

// MTLFunctionErrorHandler handles A block of code that Metal calls after it creates the specialized function.
//   - function: A specialized function, or `nil` if an error occurred.
//   - error: An error object that describes compilation problems, if any. This object contains compiler errors if the specialized function is `nil`, and compiler warnings if Metal created the specialized function with warnings. If Metal created the function without errors or warnings, this error object is `nil`.
//
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [MTLLibrary.NewFunctionWithDescriptorCompletionHandler]
//   - [MTLLibrary.NewFunctionWithNameConstantValuesCompletionHandler]
//   - [MTLLibrary.NewIntersectionFunctionWithDescriptorCompletionHandler]
type MTLFunctionErrorHandler = func(MTLFunction, error)

// NewMTLFunctionErrorBlock wraps a Go [MTLFunctionErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MTLLibrary.NewFunctionWithDescriptorCompletionHandler]
//   - [MTLLibrary.NewFunctionWithNameConstantValuesCompletionHandler]
//   - [MTLLibrary.NewIntersectionFunctionWithDescriptorCompletionHandler]
func NewMTLFunctionErrorBlock(handler MTLFunctionErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result MTLFunction
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			result = MTLFunctionObjectFromID(resultID)
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// VoidHandler is the signature for a completion handler block.
//
// Used by:
//   - [MTLLogState.AddLogHandler]
type VoidHandler = func()

// NewVoidBlock wraps a Go [VoidHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MTLLogState.AddLogHandler]
func NewVoidBlock(handler VoidHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}
