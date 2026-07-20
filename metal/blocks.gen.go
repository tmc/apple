// Code generated from Apple documentation. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// ErrorHandler handles A completetion handler that you provide, which the task calls when it finishes compiling the binary function.
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
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}

// MTL4CommitFeedbackHandler handles Defines the block signature for a callback Metal invokes to provide your app feedback after completing a workload.

// NewMTL4CommitFeedbackHandlerBlock wraps a Go [MTL4CommitFeedbackHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewMTL4CommitFeedbackHandlerBlock(handler MTL4CommitFeedbackHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal MTL4CommitFeedback) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// MTL4NewBinaryFunctionCompletionHandler handles Provides a signature for a callback block that Metal calls when the compiler finishes a build task for a binary function.

// NewMTL4NewBinaryFunctionCompletionHandlerBlock wraps a Go [MTL4NewBinaryFunctionCompletionHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewMTL4NewBinaryFunctionCompletionHandlerBlock(handler MTL4NewBinaryFunctionCompletionHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive MTL4BinaryFunction, extra0 foundation.NSError) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// MTL4NewMachineLearningPipelineStateCompletionHandler handles Provides a signature for a callback block that Metal calls when the compiler finishes a build task for a machine learning pipeline state.

// NewMTL4NewMachineLearningPipelineStateCompletionHandlerBlock wraps a Go [MTL4NewMachineLearningPipelineStateCompletionHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewMTL4NewMachineLearningPipelineStateCompletionHandlerBlock(handler MTL4NewMachineLearningPipelineStateCompletionHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive MTL4MachineLearningPipelineState, extra0 foundation.NSError) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// MTLCommandBufferHandler handles A completion handler signature a GPU device calls when it finishes scheduling a command buffer, or when the GPU finishes running it.

// NewMTLCommandBufferHandlerBlock wraps a Go [MTLCommandBufferHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewMTLCommandBufferHandlerBlock(handler MTLCommandBufferHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal MTLCommandBuffer) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// MTLDeviceNotificationHandler handles A Swift closure or an Objective-C block that Metal calls when the system adds or removes a GPU device.

// NewMTLDeviceNotificationHandlerBlock wraps a Go [MTLDeviceNotificationHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewMTLDeviceNotificationHandlerBlock(handler MTLDeviceNotificationHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive MTLDevice, extra0 string) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// MTLDrawablePresentedHandler handles A block of code invoked after a drawable is presented.

// NewMTLDrawablePresentedHandlerBlock wraps a Go [MTLDrawablePresentedHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewMTLDrawablePresentedHandlerBlock(handler MTLDrawablePresentedHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal MTLDrawable) {
		handler(primitiveVal)
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
	if handler == nil {
		return 0, func() {}
	}
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

// MTLIOCommandBufferHandler handles A convenience type that defines the signature of an input/output command buffer’s completion handler.

// NewMTLIOCommandBufferHandlerBlock wraps a Go [MTLIOCommandBufferHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewMTLIOCommandBufferHandlerBlock(handler MTLIOCommandBufferHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitiveVal MTLIOCommandBuffer) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// MTLNewComputePipelineStateCompletionHandler handles A completion handler signature a method calls when it finishes creating a compute pipeline.

// NewMTLNewComputePipelineStateCompletionHandlerBlock wraps a Go [MTLNewComputePipelineStateCompletionHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewMTLNewComputePipelineStateCompletionHandlerBlock(handler MTLNewComputePipelineStateCompletionHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive MTLComputePipelineState, extra0 foundation.NSError) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// MTLNewComputePipelineStateWithReflectionCompletionHandler handles A completion handler signature a method calls when it finishes creating a compute pipeline and reflection information.

// NewMTLNewComputePipelineStateWithReflectionCompletionHandlerBlock wraps a Go [MTLNewComputePipelineStateWithReflectionCompletionHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewMTLNewComputePipelineStateWithReflectionCompletionHandlerBlock(handler MTLNewComputePipelineStateWithReflectionCompletionHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive MTLComputePipelineState, extra0 MTLComputePipelineReflection, extra1 foundation.NSError) {
		handler(primitive, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

// MTLNewDynamicLibraryCompletionHandler handles completion with primitive and object results.

// NewMTLNewDynamicLibraryCompletionHandlerBlock wraps a Go [MTLNewDynamicLibraryCompletionHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewMTLNewDynamicLibraryCompletionHandlerBlock(handler MTLNewDynamicLibraryCompletionHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive MTLDynamicLibrary, extra0 foundation.NSError) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// MTLNewLibraryCompletionHandler handles A completion handler signature a method calls when it finishes creating a Metal library.

// NewMTLNewLibraryCompletionHandlerBlock wraps a Go [MTLNewLibraryCompletionHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewMTLNewLibraryCompletionHandlerBlock(handler MTLNewLibraryCompletionHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive MTLLibrary, extra0 foundation.NSError) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// MTLNewRenderPipelineStateCompletionHandler handles A completion handler signature a method calls when it finishes creating a render pipeline.

// NewMTLNewRenderPipelineStateCompletionHandlerBlock wraps a Go [MTLNewRenderPipelineStateCompletionHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewMTLNewRenderPipelineStateCompletionHandlerBlock(handler MTLNewRenderPipelineStateCompletionHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive MTLRenderPipelineState, extra0 foundation.NSError) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// MTLNewRenderPipelineStateWithReflectionCompletionHandler handles A completion handler signature a method calls when it finishes creating a render pipeline and reflection information.

// NewMTLNewRenderPipelineStateWithReflectionCompletionHandlerBlock wraps a Go [MTLNewRenderPipelineStateWithReflectionCompletionHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewMTLNewRenderPipelineStateWithReflectionCompletionHandlerBlock(handler MTLNewRenderPipelineStateWithReflectionCompletionHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive MTLRenderPipelineState, extra0 MTLRenderPipelineReflection, extra1 foundation.NSError) {
		handler(primitive, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

// MTLSharedEventNotificationBlock handles A block of code invoked after a shareable event’s signal value equals or exceeds a given value.

// NewMTLSharedEventNotificationBlock wraps a Go [MTLSharedEventNotificationBlock] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewMTLSharedEventNotificationBlock(handler MTLSharedEventNotificationBlock) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive MTLSharedEvent, extra0 uint64) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// StringStringMTLLogLevelStringHandler is the signature for a completion handler block.
//
// Used by:
//   - [MTLLogState.AddLogHandler]
type StringStringMTLLogLevelStringHandler = func(*string, string, MTLLogLevel, string)

// NewStringStringMTLLogLevelStringBlock wraps a Go [StringStringMTLLogLevelStringHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MTLLogState.AddLogHandler]
func NewStringStringMTLLogLevelStringBlock(handler StringStringMTLLogLevelStringHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0ID objc.ID, extra1 MTLLogLevel, extra2ID objc.ID) {
		var result *string
		if resultID != 0 {
			v := objc.IDToString(resultID)
			result = &v
		}
		var extra0 string = string(extra0ID)
		var extra2 string = string(extra2ID)
		handler(result, extra0, extra1, extra2)
	})
	return objc.ID(block), func() { block.Release() }
}
