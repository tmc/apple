// Code generated from Apple documentation. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// MTL4BinaryFunctionErrorHandler handles A completetion handler that you provide, which the task calls when it finishes compiling the binary function.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [MTL4Compiler.NewBinaryFunctionWithDescriptorCompilerTaskOptionsCompletionHandler]
type MTL4BinaryFunctionErrorHandler = func(MTL4BinaryFunction, error)

// NewMTL4BinaryFunctionErrorBlock wraps a Go [MTL4BinaryFunctionErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MTL4Compiler.NewBinaryFunctionWithDescriptorCompilerTaskOptionsCompletionHandler]
func NewMTL4BinaryFunctionErrorBlock(handler MTL4BinaryFunctionErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result MTL4BinaryFunction
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			result = MTL4BinaryFunctionObjectFromID(resultID)
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// MTL4CommitFeedbackHandler handles MTL4CommitFeedbackHandler that Metal invokes.
//
// Used by:
//   - [MTL4CommitOptions.AddFeedbackHandler]

// NewMTL4CommitFeedbackBlock wraps a Go [MTL4CommitFeedbackHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MTL4CommitOptions.AddFeedbackHandler]
func NewMTL4CommitFeedbackBlock(handler MTL4CommitFeedbackHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result MTL4CommitFeedback
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			result = MTL4CommitFeedbackObjectFromID(resultID)
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// MTL4MachineLearningPipelineStateErrorHandler handles A block Metal calls when it finishes the build task.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [MTL4Compiler.NewMachineLearningPipelineStateWithDescriptorCompletionHandler]
type MTL4MachineLearningPipelineStateErrorHandler = func(MTL4MachineLearningPipelineState, error)

// NewMTL4MachineLearningPipelineStateErrorBlock wraps a Go [MTL4MachineLearningPipelineStateErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MTL4Compiler.NewMachineLearningPipelineStateWithDescriptorCompletionHandler]
func NewMTL4MachineLearningPipelineStateErrorBlock(handler MTL4MachineLearningPipelineStateErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result MTL4MachineLearningPipelineState
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			result = MTL4MachineLearningPipelineStateObjectFromID(resultID)
		}
		handler(result, foundation.SafeErrorFrom(errID))
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
	block := objc.NewBlock(func(b objc.Block, primitiveID objc.ID, extra0 foundation.NSError) {
		var primitive MTL4BinaryFunction
		if primitiveID != 0 {
			objc.Send[objc.ID](primitiveID, objc.Sel("retain"))
			primitive = MTL4BinaryFunctionObjectFromID(primitiveID)
		}
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
	block := objc.NewBlock(func(b objc.Block, primitiveID objc.ID, extra0 foundation.NSError) {
		var primitive MTL4MachineLearningPipelineState
		if primitiveID != 0 {
			objc.Send[objc.ID](primitiveID, objc.Sel("retain"))
			primitive = MTL4MachineLearningPipelineStateObjectFromID(primitiveID)
		}
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// MTLCommandBufferHandler handles A Swift closure or an Objective-C block that Metal calls after it schedules the command buffer to run on the GPU.
//
// Used by:
//   - [MTLCommandBuffer.AddCompletedHandler]
//   - [MTLCommandBuffer.AddScheduledHandler]

// NewMTLCommandBufferBlock wraps a Go [MTLCommandBufferHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MTLCommandBuffer.AddCompletedHandler]
//   - [MTLCommandBuffer.AddScheduledHandler]
func NewMTLCommandBufferBlock(handler MTLCommandBufferHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result MTLCommandBuffer
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			result = MTLCommandBufferObjectFromID(resultID)
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// MTLComputePipelineStateErrorHandler handles A block Metal calls when it finishes the build task.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [MTL4Compiler.NewComputePipelineStateWithDescriptorCompilerTaskOptionsCompletionHandler]
//   - [MTL4Compiler.NewComputePipelineStateWithDescriptorDynamicLinkingDescriptorCompilerTaskOptionsCompletionHandler]
//   - [MTLDevice.NewComputePipelineStateWithFunctionCompletionHandler]
type MTLComputePipelineStateErrorHandler = func(MTLComputePipelineState, error)

// NewMTLComputePipelineStateErrorBlock wraps a Go [MTLComputePipelineStateErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MTL4Compiler.NewComputePipelineStateWithDescriptorCompilerTaskOptionsCompletionHandler]
//   - [MTL4Compiler.NewComputePipelineStateWithDescriptorDynamicLinkingDescriptorCompilerTaskOptionsCompletionHandler]
//   - [MTLDevice.NewComputePipelineStateWithFunctionCompletionHandler]
func NewMTLComputePipelineStateErrorBlock(handler MTLComputePipelineStateErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result MTLComputePipelineState
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			result = MTLComputePipelineStateObjectFromID(resultID)
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// MTLComputePipelineStateMTLComputePipelineReflectionErrorHandler handles A Swift closure or an Objective-C block the method calls when it finishes creating the compute pipeline state.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [MTLDevice.NewComputePipelineStateWithDescriptorOptionsCompletionHandler]
//   - [MTLDevice.NewComputePipelineStateWithFunctionOptionsCompletionHandler]
type MTLComputePipelineStateMTLComputePipelineReflectionErrorHandler = func(MTLComputePipelineState, *MTLComputePipelineReflection, error)

// NewMTLComputePipelineStateMTLComputePipelineReflectionErrorBlock wraps a Go [MTLComputePipelineStateMTLComputePipelineReflectionErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MTLDevice.NewComputePipelineStateWithDescriptorOptionsCompletionHandler]
//   - [MTLDevice.NewComputePipelineStateWithFunctionOptionsCompletionHandler]
func NewMTLComputePipelineStateMTLComputePipelineReflectionErrorBlock(handler MTLComputePipelineStateMTLComputePipelineReflectionErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0ID objc.ID, errID objc.ID) {
		var result MTLComputePipelineState
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			result = MTLComputePipelineStateObjectFromID(resultID)
		}
		var extra0 *MTLComputePipelineReflection
		if extra0ID != 0 {
			objc.Send[objc.ID](extra0ID, objc.Sel("retain"))
			v := MTLComputePipelineReflectionFromID(extra0ID)
			extra0 = &v
		}
		handler(result, extra0, foundation.SafeErrorFrom(errID))
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
	block := objc.NewBlock(func(b objc.Block, primitiveID objc.ID, extra0 string) {
		var primitive MTLDevice
		if primitiveID != 0 {
			objc.Send[objc.ID](primitiveID, objc.Sel("retain"))
			primitive = MTLDeviceObjectFromID(primitiveID)
		}
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// MTLDrawableHandler handles A block of code to be invoked.
//
// Used by:
//   - [MTLDrawable.AddPresentedHandler]
type MTLDrawableHandler = func(MTLDrawable)

// NewMTLDrawableBlock wraps a Go [MTLDrawableHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MTLDrawable.AddPresentedHandler]
func NewMTLDrawableBlock(handler MTLDrawableHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result MTLDrawable
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			result = MTLDrawableObjectFromID(resultID)
		}
		handler(result)
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
	block := objc.NewBlock(func(b objc.Block, primitiveID objc.ID) {
		var primitiveVal MTLDrawable
		if primitiveID != 0 {
			objc.Send[objc.ID](primitiveID, objc.Sel("retain"))
			primitiveVal = MTLDrawableObjectFromID(primitiveID)
		}
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// MTLDynamicLibraryErrorHandler handles A block Metal calls when it finishes the build task.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [MTL4Compiler.NewDynamicLibraryCompletionHandler]
//   - [MTL4Compiler.NewDynamicLibraryWithURLCompletionHandler]
type MTLDynamicLibraryErrorHandler = func(MTLDynamicLibrary, error)

// NewMTLDynamicLibraryErrorBlock wraps a Go [MTLDynamicLibraryErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MTL4Compiler.NewDynamicLibraryCompletionHandler]
//   - [MTL4Compiler.NewDynamicLibraryWithURLCompletionHandler]
func NewMTLDynamicLibraryErrorBlock(handler MTLDynamicLibraryErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result MTLDynamicLibrary
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			result = MTLDynamicLibraryObjectFromID(resultID)
		}
		handler(result, foundation.SafeErrorFrom(errID))
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

// MTLIOCommandBufferHandler handles A Swift closure or an Objective-C block with your code.
//
// Used by:
//   - [MTLIOCommandBuffer.AddCompletedHandler]

// NewMTLIOCommandBufferBlock wraps a Go [MTLIOCommandBufferHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MTLIOCommandBuffer.AddCompletedHandler]
func NewMTLIOCommandBufferBlock(handler MTLIOCommandBufferHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result MTLIOCommandBuffer
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			result = MTLIOCommandBufferObjectFromID(resultID)
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// MTLLibraryErrorHandler handles A block Metal calls when it finishes the build task.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [MTL4Compiler.NewLibraryWithDescriptorCompletionHandler]
//   - [MTLDevice.NewLibraryWithSourceOptionsCompletionHandler]
//   - [MTLDevice.NewLibraryWithStitchedDescriptorCompletionHandler]
type MTLLibraryErrorHandler = func(MTLLibrary, error)

// NewMTLLibraryErrorBlock wraps a Go [MTLLibraryErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MTL4Compiler.NewLibraryWithDescriptorCompletionHandler]
//   - [MTLDevice.NewLibraryWithSourceOptionsCompletionHandler]
//   - [MTLDevice.NewLibraryWithStitchedDescriptorCompletionHandler]
func NewMTLLibraryErrorBlock(handler MTLLibraryErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result MTLLibrary
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			result = MTLLibraryObjectFromID(resultID)
		}
		handler(result, foundation.SafeErrorFrom(errID))
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
	block := objc.NewBlock(func(b objc.Block, primitiveID objc.ID, extra0 foundation.NSError) {
		var primitive MTLComputePipelineState
		if primitiveID != 0 {
			objc.Send[objc.ID](primitiveID, objc.Sel("retain"))
			primitive = MTLComputePipelineStateObjectFromID(primitiveID)
		}
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
	block := objc.NewBlock(func(b objc.Block, primitiveID objc.ID, extra0 MTLComputePipelineReflection, extra1 foundation.NSError) {
		var primitive MTLComputePipelineState
		if primitiveID != 0 {
			objc.Send[objc.ID](primitiveID, objc.Sel("retain"))
			primitive = MTLComputePipelineStateObjectFromID(primitiveID)
		}
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
	block := objc.NewBlock(func(b objc.Block, primitiveID objc.ID, extra0 foundation.NSError) {
		var primitive MTLDynamicLibrary
		if primitiveID != 0 {
			objc.Send[objc.ID](primitiveID, objc.Sel("retain"))
			primitive = MTLDynamicLibraryObjectFromID(primitiveID)
		}
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
	block := objc.NewBlock(func(b objc.Block, primitiveID objc.ID, extra0 foundation.NSError) {
		var primitive MTLLibrary
		if primitiveID != 0 {
			objc.Send[objc.ID](primitiveID, objc.Sel("retain"))
			primitive = MTLLibraryObjectFromID(primitiveID)
		}
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
	block := objc.NewBlock(func(b objc.Block, primitiveID objc.ID, extra0 foundation.NSError) {
		var primitive MTLRenderPipelineState
		if primitiveID != 0 {
			objc.Send[objc.ID](primitiveID, objc.Sel("retain"))
			primitive = MTLRenderPipelineStateObjectFromID(primitiveID)
		}
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
	block := objc.NewBlock(func(b objc.Block, primitiveID objc.ID, extra0 MTLRenderPipelineReflection, extra1 foundation.NSError) {
		var primitive MTLRenderPipelineState
		if primitiveID != 0 {
			objc.Send[objc.ID](primitiveID, objc.Sel("retain"))
			primitive = MTLRenderPipelineStateObjectFromID(primitiveID)
		}
		handler(primitive, extra0, extra1)
	})
	return objc.ID(block), func() { block.Release() }
}

// MTLRenderPipelineStateErrorHandler handles A block Metal calls when it finishes the build task.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [MTL4Compiler.NewRenderPipelineStateBySpecializationWithDescriptorPipelineCompletionHandler]
//   - [MTL4Compiler.NewRenderPipelineStateWithDescriptorCompilerTaskOptionsCompletionHandler]
//   - [MTL4Compiler.NewRenderPipelineStateWithDescriptorDynamicLinkingDescriptorCompilerTaskOptionsCompletionHandler]
//   - [MTLDevice.NewRenderPipelineStateWithDescriptorCompletionHandler]
type MTLRenderPipelineStateErrorHandler = func(MTLRenderPipelineState, error)

// NewMTLRenderPipelineStateErrorBlock wraps a Go [MTLRenderPipelineStateErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MTL4Compiler.NewRenderPipelineStateBySpecializationWithDescriptorPipelineCompletionHandler]
//   - [MTL4Compiler.NewRenderPipelineStateWithDescriptorCompilerTaskOptionsCompletionHandler]
//   - [MTL4Compiler.NewRenderPipelineStateWithDescriptorDynamicLinkingDescriptorCompilerTaskOptionsCompletionHandler]
//   - [MTLDevice.NewRenderPipelineStateWithDescriptorCompletionHandler]
func NewMTLRenderPipelineStateErrorBlock(handler MTLRenderPipelineStateErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, errID objc.ID) {
		var result MTLRenderPipelineState
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			result = MTLRenderPipelineStateObjectFromID(resultID)
		}
		handler(result, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// MTLRenderPipelineStateMTLRenderPipelineReflectionErrorHandler handles A Swift closure or an Objective-C block the method calls when it finishes creating the render pipeline state.
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [MTLDevice.NewRenderPipelineStateWithDescriptorOptionsCompletionHandler]
//   - [MTLDevice.NewRenderPipelineStateWithMeshDescriptorOptionsCompletionHandler]
//   - [MTLDevice.NewRenderPipelineStateWithTileDescriptorOptionsCompletionHandler]
type MTLRenderPipelineStateMTLRenderPipelineReflectionErrorHandler = func(MTLRenderPipelineState, *MTLRenderPipelineReflection, error)

// NewMTLRenderPipelineStateMTLRenderPipelineReflectionErrorBlock wraps a Go [MTLRenderPipelineStateMTLRenderPipelineReflectionErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MTLDevice.NewRenderPipelineStateWithDescriptorOptionsCompletionHandler]
//   - [MTLDevice.NewRenderPipelineStateWithMeshDescriptorOptionsCompletionHandler]
//   - [MTLDevice.NewRenderPipelineStateWithTileDescriptorOptionsCompletionHandler]
func NewMTLRenderPipelineStateMTLRenderPipelineReflectionErrorBlock(handler MTLRenderPipelineStateMTLRenderPipelineReflectionErrorHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0ID objc.ID, errID objc.ID) {
		var result MTLRenderPipelineState
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			result = MTLRenderPipelineStateObjectFromID(resultID)
		}
		var extra0 *MTLRenderPipelineReflection
		if extra0ID != 0 {
			objc.Send[objc.ID](extra0ID, objc.Sel("retain"))
			v := MTLRenderPipelineReflectionFromID(extra0ID)
			extra0 = &v
		}
		handler(result, extra0, foundation.SafeErrorFrom(errID))
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
	block := objc.NewBlock(func(b objc.Block, primitiveID objc.ID, extra0 uint64) {
		var primitive MTLSharedEvent
		if primitiveID != 0 {
			objc.Send[objc.ID](primitiveID, objc.Sel("retain"))
			primitive = MTLSharedEventObjectFromID(primitiveID)
		}
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// MTLSharedEventUint64Handler handles The notification handler to call.
//
// Used by:
//   - [MTLSharedEvent.NotifyListenerAtValueBlock]
type MTLSharedEventUint64Handler = func(MTLSharedEvent, uint64)

// NewMTLSharedEventUint64Block wraps a Go [MTLSharedEventUint64Handler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MTLSharedEvent.NotifyListenerAtValueBlock]
func NewMTLSharedEventUint64Block(handler MTLSharedEventUint64Handler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0 uint64) {
		var result MTLSharedEvent
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			result = MTLSharedEventObjectFromID(resultID)
		}
		handler(result, extra0)
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
		var extra0 string = objc.IDToString(extra0ID)
		var extra2 string = objc.IDToString(extra2ID)
		handler(result, extra0, extra1, extra2)
	})
	return objc.ID(block), func() { block.Release() }
}

// UnsafePointerUintHandler handles A block the framework invokes when it deallocates the buffer so that your app can release the underlying memory; otherwise `nil` to opt out.
//
// Used by:
//   - [MTLDevice.NewBufferWithBytesNoCopyLengthOptionsDeallocator]
type UnsafePointerUintHandler = func(unsafe.Pointer, uint)

// NewUnsafePointerUintBlock wraps a Go [UnsafePointerUintHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [MTLDevice.NewBufferWithBytesNoCopyLengthOptionsDeallocator]
func NewUnsafePointerUintBlock(handler UnsafePointerUintHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, primitive unsafe.Pointer, extra0 uint) {
		handler(primitive, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}
