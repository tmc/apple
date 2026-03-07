// Code generated from Apple documentation. DO NOT EDIT.

package metal

import (
	"unsafe"
	"github.com/tmc/apple/foundation"
)

type MTL4CommitFeedbackHandler = func(MTL4CommitFeedback)







type MTL4NewBinaryFunctionCompletionHandler = func(MTL4BinaryFunction, foundation.NSError)







type MTL4NewMachineLearningPipelineStateCompletionHandler = func(MTL4MachineLearningPipelineState, foundation.NSError)







type MTLArgumentAccess = unsafe.Pointer







type MTLAutoreleasedComputePipelineReflection = *MTLComputePipelineReflection







type MTLAutoreleasedRenderPipelineReflection = *MTLRenderPipelineReflection








type MTLCommandBufferHandler = func(MTLCommandBuffer)







type MTLCommonCounter = string







type MTLCommonCounterSet = string







type MTLCoordinate2D = MTLSamplePosition







type MTLDeviceNotificationHandler = func(MTLDevice, string)







type MTLDeviceNotificationName = string







type MTLDrawablePresentedHandler = func(MTLDrawable)







type MTLGPUAddress = uint64







type MTLIOCommandBufferHandler = func(MTLIOCommandBuffer)







type MTLIOCompressionContext = unsafe.Pointer







type MTLNewComputePipelineStateCompletionHandler = func(MTLComputePipelineState, foundation.NSError)







type MTLNewComputePipelineStateWithReflectionCompletionHandler = func(MTLComputePipelineState, MTLComputePipelineReflection, foundation.NSError)







type MTLNewDynamicLibraryCompletionHandler = func(MTLDynamicLibrary, foundation.NSError)







type MTLNewLibraryCompletionHandler = func(MTLLibrary, foundation.NSError)







type MTLNewRenderPipelineStateCompletionHandler = func(MTLRenderPipelineState, foundation.NSError)







type MTLNewRenderPipelineStateWithReflectionCompletionHandler = func(MTLRenderPipelineState, MTLRenderPipelineReflection, foundation.NSError)









type MTLSharedEventNotificationBlock = func(MTLSharedEvent, uint64)







type MTLTimestamp = uint64







type NSDeviceCertification = int







type NSProcessPerformanceProfile = int





