// Code generated from Apple documentation. DO NOT EDIT.

package metal

import (
	"unsafe"
	"github.com/tmc/apple/foundation"
)

// MTL4CommitFeedbackHandler is defines the block signature for a callback Metal invokes to provide your app feedback after completing a workload.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommitFeedbackHandler
type MTL4CommitFeedbackHandler = func(MTL4CommitFeedback)

// MTL4NewBinaryFunctionCompletionHandler is provides a signature for a callback block that Metal calls when the compiler finishes a build task for a binary function.
//
// See: https://developer.apple.com/documentation/Metal/MTL4NewBinaryFunctionCompletionHandler
type MTL4NewBinaryFunctionCompletionHandler = func(MTL4BinaryFunction, foundation.NSError)

// MTL4NewMachineLearningPipelineStateCompletionHandler is provides a signature for a callback block that Metal calls when the compiler finishes a build task for a machine learning pipeline state.
//
// See: https://developer.apple.com/documentation/Metal/MTL4NewMachineLearningPipelineStateCompletionHandler
type MTL4NewMachineLearningPipelineStateCompletionHandler = func(MTL4MachineLearningPipelineState, foundation.NSError)

// MTLArgumentAccess is function access restrictions to argument data in the shading language code.
//
// Deprecated: Deprecated since macOS 14.0.
//
// See: https://developer.apple.com/documentation/Metal/MTLArgumentAccess
type MTLArgumentAccess = unsafe.Pointer

// MTLAutoreleasedComputePipelineReflection is a convenience type alias for an autoreleased compute pipeline reflection object.
//
// See: https://developer.apple.com/documentation/Metal/MTLAutoreleasedComputePipelineReflection
type MTLAutoreleasedComputePipelineReflection = *MTLComputePipelineReflection

// MTLAutoreleasedRenderPipelineReflection is a convenience type alias for an autoreleased pipeline reflection instance.
//
// See: https://developer.apple.com/documentation/Metal/MTLAutoreleasedRenderPipelineReflection
type MTLAutoreleasedRenderPipelineReflection = *MTLRenderPipelineReflection

// MTLCommandBufferHandler is a completion handler signature a GPU device calls when it finishes scheduling a command buffer, or when the GPU finishes running it.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandBufferHandler
type MTLCommandBufferHandler = func(MTLCommandBuffer)

// MTLCommonCounter is the name of a specific counter that can appear in a GPU device’s counter sets.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommonCounter
type MTLCommonCounter = string

// MTLCommonCounterSet is the name of a specific counter set that a GPU device can support.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommonCounterSet
type MTLCommonCounterSet = string

// MTLCoordinate2D is a coordinate in the viewport.
//
// See: https://developer.apple.com/documentation/Metal/MTLCoordinate2D
type MTLCoordinate2D = MTLSamplePosition

// MTLDeviceNotificationHandler is a Swift closure or an Objective-C block that Metal calls when the system adds or removes a GPU device.
//
// See: https://developer.apple.com/documentation/Metal/MTLDeviceNotificationHandler
type MTLDeviceNotificationHandler = func(MTLDevice, string)

// MTLDeviceNotificationName is a notification that represents a change to a GPU device in the system.
//
// See: https://developer.apple.com/documentation/Metal/MTLDeviceNotificationName
type MTLDeviceNotificationName = string

// MTLDrawablePresentedHandler is a block of code invoked after a drawable is presented.
//
// See: https://developer.apple.com/documentation/Metal/MTLDrawablePresentedHandler
type MTLDrawablePresentedHandler = func(MTLDrawable)

// MTLGPUAddress is a 64-bit unsigned integer type appropriate for storing GPU addresses.
//
// See: https://developer.apple.com/documentation/Metal/MTLGPUAddress
type MTLGPUAddress = uint64

// MTLIOCommandBufferHandler is a convenience type that defines the signature of an input/output command buffer’s completion handler.
//
// See: https://developer.apple.com/documentation/Metal/MTLIOCommandBufferHandler
type MTLIOCommandBufferHandler = func(MTLIOCommandBuffer)

// MTLIOCompressionContext is a pointer that represents the state of a file compression session in progress.
//
// See: https://developer.apple.com/documentation/Metal/MTLIOCompressionContext
type MTLIOCompressionContext = unsafe.Pointer

// MTLNewComputePipelineStateCompletionHandler is a completion handler signature a method calls when it finishes creating a compute pipeline.
//
// See: https://developer.apple.com/documentation/Metal/MTLNewComputePipelineStateCompletionHandler
type MTLNewComputePipelineStateCompletionHandler = func(MTLComputePipelineState, foundation.NSError)

// MTLNewComputePipelineStateWithReflectionCompletionHandler is a completion handler signature a method calls when it finishes creating a compute pipeline and reflection information.
//
// See: https://developer.apple.com/documentation/Metal/MTLNewComputePipelineStateWithReflectionCompletionHandler
type MTLNewComputePipelineStateWithReflectionCompletionHandler = func(MTLComputePipelineState, MTLComputePipelineReflection, foundation.NSError)

// See: https://developer.apple.com/documentation/Metal/MTLNewDynamicLibraryCompletionHandler
type MTLNewDynamicLibraryCompletionHandler = func(MTLDynamicLibrary, foundation.NSError)

// MTLNewLibraryCompletionHandler is a completion handler signature a method calls when it finishes creating a Metal library.
//
// See: https://developer.apple.com/documentation/Metal/MTLNewLibraryCompletionHandler
type MTLNewLibraryCompletionHandler = func(MTLLibrary, foundation.NSError)

// MTLNewRenderPipelineStateCompletionHandler is a completion handler signature a method calls when it finishes creating a render pipeline.
//
// See: https://developer.apple.com/documentation/Metal/MTLNewRenderPipelineStateCompletionHandler
type MTLNewRenderPipelineStateCompletionHandler = func(MTLRenderPipelineState, foundation.NSError)

// MTLNewRenderPipelineStateWithReflectionCompletionHandler is a completion handler signature a method calls when it finishes creating a render pipeline and reflection information.
//
// See: https://developer.apple.com/documentation/Metal/MTLNewRenderPipelineStateWithReflectionCompletionHandler
type MTLNewRenderPipelineStateWithReflectionCompletionHandler = func(MTLRenderPipelineState, MTLRenderPipelineReflection, foundation.NSError)

// MTLSharedEventNotificationBlock is a block of code invoked after a shareable event’s signal value equals or exceeds a given value.
//
// See: https://developer.apple.com/documentation/Metal/MTLSharedEventNotificationBlock
type MTLSharedEventNotificationBlock = func(MTLSharedEvent, uint64)

// MTLTimestamp is the number of nanoseconds for a point in absolute time or Mach absolute time.
//
// See: https://developer.apple.com/documentation/Metal/MTLTimestamp
type MTLTimestamp = uint64

// See: https://developer.apple.com/documentation/Metal/NSDeviceCertification
type NSDeviceCertification = int

// NSProcessPerformanceProfile is a value describing the device’s performance profile.
//
// See: https://developer.apple.com/documentation/Metal/NSProcessPerformanceProfile
type NSProcessPerformanceProfile = int

