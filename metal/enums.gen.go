// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/Metal/MTL4AlphaToCoverageState
type MTL4AlphaToCoverageState int

const (
	// MTL4AlphaToCoverageStateDisabled: Disables alpha-to-coverage.
	MTL4AlphaToCoverageStateDisabled MTL4AlphaToCoverageState = 0
	// MTL4AlphaToCoverageStateEnabled: Enables alpha-to-coverage.
	MTL4AlphaToCoverageStateEnabled MTL4AlphaToCoverageState = 1
)

func (e MTL4AlphaToCoverageState) String() string {
	switch e {
	case MTL4AlphaToCoverageStateDisabled:
		return "MTL4AlphaToCoverageStateDisabled"
	case MTL4AlphaToCoverageStateEnabled:
		return "MTL4AlphaToCoverageStateEnabled"
	default:
		return fmt.Sprintf("MTL4AlphaToCoverageState(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTL4AlphaToOneState
type MTL4AlphaToOneState int

const (
	// MTL4AlphaToOneStateDisabled: Disables alpha-to-one.
	MTL4AlphaToOneStateDisabled MTL4AlphaToOneState = 0
	// MTL4AlphaToOneStateEnabled: Enables alpha-to-one.
	MTL4AlphaToOneStateEnabled MTL4AlphaToOneState = 1
)

func (e MTL4AlphaToOneState) String() string {
	switch e {
	case MTL4AlphaToOneStateDisabled:
		return "MTL4AlphaToOneStateDisabled"
	case MTL4AlphaToOneStateEnabled:
		return "MTL4AlphaToOneStateEnabled"
	default:
		return fmt.Sprintf("MTL4AlphaToOneState(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTL4BinaryFunctionOptions
type MTL4BinaryFunctionOptions uint

const (
	// MTL4BinaryFunctionOptionNone: Represents the default value: no options.
	MTL4BinaryFunctionOptionNone MTL4BinaryFunctionOptions = 0
	// MTL4BinaryFunctionOptionPipelineIndependent: Compiles the function to have its function handles return a constant MTLResourceID across all pipeline states.
	MTL4BinaryFunctionOptionPipelineIndependent MTL4BinaryFunctionOptions = 2
)

func (e MTL4BinaryFunctionOptions) String() string {
	switch e {
	case MTL4BinaryFunctionOptionNone:
		return "MTL4BinaryFunctionOptionNone"
	case MTL4BinaryFunctionOptionPipelineIndependent:
		return "MTL4BinaryFunctionOptionPipelineIndependent"
	default:
		return fmt.Sprintf("MTL4BinaryFunctionOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTL4BlendState
type MTL4BlendState int

const (
	// MTL4BlendStateDisabled: Disables blending.
	MTL4BlendStateDisabled MTL4BlendState = 0
	// MTL4BlendStateEnabled: Enables blending.
	MTL4BlendStateEnabled MTL4BlendState = 1
	// MTL4BlendStateUnspecialized: Defers determining the blending stage.
	MTL4BlendStateUnspecialized MTL4BlendState = 2
)

func (e MTL4BlendState) String() string {
	switch e {
	case MTL4BlendStateDisabled:
		return "MTL4BlendStateDisabled"
	case MTL4BlendStateEnabled:
		return "MTL4BlendStateEnabled"
	case MTL4BlendStateUnspecialized:
		return "MTL4BlendStateUnspecialized"
	default:
		return fmt.Sprintf("MTL4BlendState(%d)", e)
	}
}

type MTL4CommandQueueError int

const (
	MTL4CommandQueueErrorAccessRevoked MTL4CommandQueueError = 5
	MTL4CommandQueueErrorDeviceRemoved MTL4CommandQueueError = 4
	// MTL4CommandQueueErrorInternal: Indicates an internal problem in the Metal framework.
	MTL4CommandQueueErrorInternal MTL4CommandQueueError = 6
	// MTL4CommandQueueErrorNone: Indicates the absence of any problems.
	MTL4CommandQueueErrorNone MTL4CommandQueueError = 0
	// MTL4CommandQueueErrorNotPermitted: Indicates a process doesn’t have access to a GPU device.
	MTL4CommandQueueErrorNotPermitted MTL4CommandQueueError = 2
	// MTL4CommandQueueErrorOutOfMemory: Indicates the GPU doesn’t have sufficient memory to execute a command buffer.
	MTL4CommandQueueErrorOutOfMemory MTL4CommandQueueError = 3
	// MTL4CommandQueueErrorTimeout: Indicates the workload takes longer to execute than the system allows.
	MTL4CommandQueueErrorTimeout MTL4CommandQueueError = 1
)

func (e MTL4CommandQueueError) String() string {
	switch e {
	case MTL4CommandQueueErrorAccessRevoked:
		return "MTL4CommandQueueErrorAccessRevoked"
	case MTL4CommandQueueErrorDeviceRemoved:
		return "MTL4CommandQueueErrorDeviceRemoved"
	case MTL4CommandQueueErrorInternal:
		return "MTL4CommandQueueErrorInternal"
	case MTL4CommandQueueErrorNone:
		return "MTL4CommandQueueErrorNone"
	case MTL4CommandQueueErrorNotPermitted:
		return "MTL4CommandQueueErrorNotPermitted"
	case MTL4CommandQueueErrorOutOfMemory:
		return "MTL4CommandQueueErrorOutOfMemory"
	case MTL4CommandQueueErrorTimeout:
		return "MTL4CommandQueueErrorTimeout"
	default:
		return fmt.Sprintf("MTL4CommandQueueError(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTL4CompilerTaskStatus
type MTL4CompilerTaskStatus int

const (
	// MTL4CompilerTaskStatusCompiling: The compiler task is currently compiling.
	MTL4CompilerTaskStatusCompiling MTL4CompilerTaskStatus = 2
	// MTL4CompilerTaskStatusFinished: The compiler task is finished.
	MTL4CompilerTaskStatusFinished MTL4CompilerTaskStatus = 3
	// MTL4CompilerTaskStatusNone: No status.
	MTL4CompilerTaskStatusNone MTL4CompilerTaskStatus = 0
	// MTL4CompilerTaskStatusScheduled: The compiler task is currently scheduled.
	MTL4CompilerTaskStatusScheduled MTL4CompilerTaskStatus = 1
)

func (e MTL4CompilerTaskStatus) String() string {
	switch e {
	case MTL4CompilerTaskStatusCompiling:
		return "MTL4CompilerTaskStatusCompiling"
	case MTL4CompilerTaskStatusFinished:
		return "MTL4CompilerTaskStatusFinished"
	case MTL4CompilerTaskStatusNone:
		return "MTL4CompilerTaskStatusNone"
	case MTL4CompilerTaskStatusScheduled:
		return "MTL4CompilerTaskStatusScheduled"
	default:
		return fmt.Sprintf("MTL4CompilerTaskStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTL4CounterHeapType
type MTL4CounterHeapType int

const (
	// MTL4CounterHeapTypeInvalid: Specifies that MTL4CounterHeap entries contain invalid data.
	MTL4CounterHeapTypeInvalid MTL4CounterHeapType = 0
	// MTL4CounterHeapTypeTimestamp: Specifies that MTL4CounterHeap entries contain GPU timestamp data.
	MTL4CounterHeapTypeTimestamp MTL4CounterHeapType = 1
)

func (e MTL4CounterHeapType) String() string {
	switch e {
	case MTL4CounterHeapTypeInvalid:
		return "MTL4CounterHeapTypeInvalid"
	case MTL4CounterHeapTypeTimestamp:
		return "MTL4CounterHeapTypeTimestamp"
	default:
		return fmt.Sprintf("MTL4CounterHeapType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTL4IndirectCommandBufferSupportState
type MTL4IndirectCommandBufferSupportState int

const (
	// MTL4IndirectCommandBufferSupportStateDisabled: Disables support for indirect command buffers.
	MTL4IndirectCommandBufferSupportStateDisabled MTL4IndirectCommandBufferSupportState = 0
	// MTL4IndirectCommandBufferSupportStateEnabled: Enables support for indirect command buffers.
	MTL4IndirectCommandBufferSupportStateEnabled MTL4IndirectCommandBufferSupportState = 1
)

func (e MTL4IndirectCommandBufferSupportState) String() string {
	switch e {
	case MTL4IndirectCommandBufferSupportStateDisabled:
		return "MTL4IndirectCommandBufferSupportStateDisabled"
	case MTL4IndirectCommandBufferSupportStateEnabled:
		return "MTL4IndirectCommandBufferSupportStateEnabled"
	default:
		return fmt.Sprintf("MTL4IndirectCommandBufferSupportState(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTL4LogicalToPhysicalColorAttachmentMappingState
type MTL4LogicalToPhysicalColorAttachmentMappingState int

const (
	// MTL4LogicalToPhysicalColorAttachmentMappingStateIdentity: Treats the logical color attachment descriptor array for render and tile render pipelines to match the physical one.
	MTL4LogicalToPhysicalColorAttachmentMappingStateIdentity MTL4LogicalToPhysicalColorAttachmentMappingState = 0
	// MTL4LogicalToPhysicalColorAttachmentMappingStateInherited: Deduces the color attachment mapping by inheriting it from the color attachment map of the current encoder.
	MTL4LogicalToPhysicalColorAttachmentMappingStateInherited MTL4LogicalToPhysicalColorAttachmentMappingState = 1
)

func (e MTL4LogicalToPhysicalColorAttachmentMappingState) String() string {
	switch e {
	case MTL4LogicalToPhysicalColorAttachmentMappingStateIdentity:
		return "MTL4LogicalToPhysicalColorAttachmentMappingStateIdentity"
	case MTL4LogicalToPhysicalColorAttachmentMappingStateInherited:
		return "MTL4LogicalToPhysicalColorAttachmentMappingStateInherited"
	default:
		return fmt.Sprintf("MTL4LogicalToPhysicalColorAttachmentMappingState(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTL4PipelineDataSetSerializerConfiguration
type MTL4PipelineDataSetSerializerConfiguration uint

const (
	// MTL4PipelineDataSetSerializerConfigurationCaptureBinaries: Enables serializing pipeline binary functions.
	MTL4PipelineDataSetSerializerConfigurationCaptureBinaries MTL4PipelineDataSetSerializerConfiguration = 2
	// MTL4PipelineDataSetSerializerConfigurationCaptureDescriptors: Enables serializing pipeline scripts.
	MTL4PipelineDataSetSerializerConfigurationCaptureDescriptors MTL4PipelineDataSetSerializerConfiguration = 1
)

func (e MTL4PipelineDataSetSerializerConfiguration) String() string {
	switch e {
	case MTL4PipelineDataSetSerializerConfigurationCaptureBinaries:
		return "MTL4PipelineDataSetSerializerConfigurationCaptureBinaries"
	case MTL4PipelineDataSetSerializerConfigurationCaptureDescriptors:
		return "MTL4PipelineDataSetSerializerConfigurationCaptureDescriptors"
	default:
		return fmt.Sprintf("MTL4PipelineDataSetSerializerConfiguration(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTL4RenderEncoderOptions
type MTL4RenderEncoderOptions uint

const (
	// MTL4RenderEncoderOptionNone: Declares that this render pass doesn’t suspend nor resume.
	MTL4RenderEncoderOptionNone MTL4RenderEncoderOptions = 0
	// MTL4RenderEncoderOptionResuming: Configures the render pass to as .
	MTL4RenderEncoderOptionResuming MTL4RenderEncoderOptions = 2
	// MTL4RenderEncoderOptionSuspending: Configures the render pass as .
	MTL4RenderEncoderOptionSuspending MTL4RenderEncoderOptions = 1
)

func (e MTL4RenderEncoderOptions) String() string {
	switch e {
	case MTL4RenderEncoderOptionNone:
		return "MTL4RenderEncoderOptionNone"
	case MTL4RenderEncoderOptionResuming:
		return "MTL4RenderEncoderOptionResuming"
	case MTL4RenderEncoderOptionSuspending:
		return "MTL4RenderEncoderOptionSuspending"
	default:
		return fmt.Sprintf("MTL4RenderEncoderOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTL4ShaderReflection
type MTL4ShaderReflection uint

const (
	// MTL4ShaderReflectionBindingInfo: Requests reflection information for bindings.
	MTL4ShaderReflectionBindingInfo MTL4ShaderReflection = 1
	// MTL4ShaderReflectionBufferTypeInfo: Requests reflection information for buffer types.
	MTL4ShaderReflectionBufferTypeInfo MTL4ShaderReflection = 2
	// MTL4ShaderReflectionNone: Requests no information.
	MTL4ShaderReflectionNone MTL4ShaderReflection = 0
)

func (e MTL4ShaderReflection) String() string {
	switch e {
	case MTL4ShaderReflectionBindingInfo:
		return "MTL4ShaderReflectionBindingInfo"
	case MTL4ShaderReflectionBufferTypeInfo:
		return "MTL4ShaderReflectionBufferTypeInfo"
	case MTL4ShaderReflectionNone:
		return "MTL4ShaderReflectionNone"
	default:
		return fmt.Sprintf("MTL4ShaderReflection(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTL4TimestampGranularity
type MTL4TimestampGranularity int

const (
	// MTL4TimestampGranularityPrecise: A timestamp as precise as possible.
	MTL4TimestampGranularityPrecise MTL4TimestampGranularity = 1
	// MTL4TimestampGranularityRelaxed: A minimally-invasive timestamp which may be less precise.
	MTL4TimestampGranularityRelaxed MTL4TimestampGranularity = 0
)

func (e MTL4TimestampGranularity) String() string {
	switch e {
	case MTL4TimestampGranularityPrecise:
		return "MTL4TimestampGranularityPrecise"
	case MTL4TimestampGranularityRelaxed:
		return "MTL4TimestampGranularityRelaxed"
	default:
		return fmt.Sprintf("MTL4TimestampGranularity(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTL4VisibilityOptions
type MTL4VisibilityOptions uint

const (
	// MTL4VisibilityOptionDevice: Flushes caches to the GPU (device) memory coherence point.
	MTL4VisibilityOptionDevice MTL4VisibilityOptions = 1
	// MTL4VisibilityOptionNone: Don’t flush caches.
	MTL4VisibilityOptionNone MTL4VisibilityOptions = 0
	// MTL4VisibilityOptionResourceAlias: Flushes caches to ensure that aliased virtual addresses are memory consistent.
	MTL4VisibilityOptionResourceAlias MTL4VisibilityOptions = 2
)

func (e MTL4VisibilityOptions) String() string {
	switch e {
	case MTL4VisibilityOptionDevice:
		return "MTL4VisibilityOptionDevice"
	case MTL4VisibilityOptionNone:
		return "MTL4VisibilityOptionNone"
	case MTL4VisibilityOptionResourceAlias:
		return "MTL4VisibilityOptionResourceAlias"
	default:
		return fmt.Sprintf("MTL4VisibilityOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureInstanceDescriptorType
type MTLAccelerationStructureInstanceDescriptorType uint

const (
	// MTLAccelerationStructureInstanceDescriptorTypeDefault: An option specifying that the instance uses the default characteristics.
	MTLAccelerationStructureInstanceDescriptorTypeDefault MTLAccelerationStructureInstanceDescriptorType = 0
	// MTLAccelerationStructureInstanceDescriptorTypeIndirect: An option that enables an instance descriptor memory layout the GPU can populate.
	MTLAccelerationStructureInstanceDescriptorTypeIndirect MTLAccelerationStructureInstanceDescriptorType = 3
	// MTLAccelerationStructureInstanceDescriptorTypeIndirectMotion: An option specifying that the instance contains motion data, and enables using an instance descriptor memory layout that the GPU can populate.
	MTLAccelerationStructureInstanceDescriptorTypeIndirectMotion MTLAccelerationStructureInstanceDescriptorType = 4
	// MTLAccelerationStructureInstanceDescriptorTypeMotion: An option specifying that the instance contains motion data.
	MTLAccelerationStructureInstanceDescriptorTypeMotion MTLAccelerationStructureInstanceDescriptorType = 2
	// MTLAccelerationStructureInstanceDescriptorTypeUserID: An option specifying that the instance contains a user identifier.
	MTLAccelerationStructureInstanceDescriptorTypeUserID MTLAccelerationStructureInstanceDescriptorType = 1
)

func (e MTLAccelerationStructureInstanceDescriptorType) String() string {
	switch e {
	case MTLAccelerationStructureInstanceDescriptorTypeDefault:
		return "MTLAccelerationStructureInstanceDescriptorTypeDefault"
	case MTLAccelerationStructureInstanceDescriptorTypeIndirect:
		return "MTLAccelerationStructureInstanceDescriptorTypeIndirect"
	case MTLAccelerationStructureInstanceDescriptorTypeIndirectMotion:
		return "MTLAccelerationStructureInstanceDescriptorTypeIndirectMotion"
	case MTLAccelerationStructureInstanceDescriptorTypeMotion:
		return "MTLAccelerationStructureInstanceDescriptorTypeMotion"
	case MTLAccelerationStructureInstanceDescriptorTypeUserID:
		return "MTLAccelerationStructureInstanceDescriptorTypeUserID"
	default:
		return fmt.Sprintf("MTLAccelerationStructureInstanceDescriptorType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureInstanceOptions
type MTLAccelerationStructureInstanceOptions uint32

const (
	// MTLAccelerationStructureInstanceOptionDisableTriangleCulling: An option that turns off culling for this instance if ray intersector has culling enabled.
	MTLAccelerationStructureInstanceOptionDisableTriangleCulling MTLAccelerationStructureInstanceOptions = 1
	// MTLAccelerationStructureInstanceOptionNonOpaque: Specifies that intersectors should treat the instance as non-opaque.
	MTLAccelerationStructureInstanceOptionNonOpaque MTLAccelerationStructureInstanceOptions = 8
	// MTLAccelerationStructureInstanceOptionNone: Specifies the default behavior for resulting acceleration structure.
	MTLAccelerationStructureInstanceOptionNone MTLAccelerationStructureInstanceOptions = 0
	// MTLAccelerationStructureInstanceOptionOpaque: Specifies that intersectors should treat the instance as opaque.
	MTLAccelerationStructureInstanceOptionOpaque MTLAccelerationStructureInstanceOptions = 4
	// MTLAccelerationStructureInstanceOptionTriangleFrontFacingWindingCounterClockwise: Specifies that the instance specifies front facing triangles in counter-clockwise order.
	MTLAccelerationStructureInstanceOptionTriangleFrontFacingWindingCounterClockwise MTLAccelerationStructureInstanceOptions = 2
)

func (e MTLAccelerationStructureInstanceOptions) String() string {
	switch e {
	case MTLAccelerationStructureInstanceOptionDisableTriangleCulling:
		return "MTLAccelerationStructureInstanceOptionDisableTriangleCulling"
	case MTLAccelerationStructureInstanceOptionNonOpaque:
		return "MTLAccelerationStructureInstanceOptionNonOpaque"
	case MTLAccelerationStructureInstanceOptionNone:
		return "MTLAccelerationStructureInstanceOptionNone"
	case MTLAccelerationStructureInstanceOptionOpaque:
		return "MTLAccelerationStructureInstanceOptionOpaque"
	case MTLAccelerationStructureInstanceOptionTriangleFrontFacingWindingCounterClockwise:
		return "MTLAccelerationStructureInstanceOptionTriangleFrontFacingWindingCounterClockwise"
	default:
		return fmt.Sprintf("MTLAccelerationStructureInstanceOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureRefitOptions
type MTLAccelerationStructureRefitOptions uint

const (
	MTLAccelerationStructureRefitOptionPerPrimitiveData MTLAccelerationStructureRefitOptions = 2
	MTLAccelerationStructureRefitOptionVertexData       MTLAccelerationStructureRefitOptions = 1
)

func (e MTLAccelerationStructureRefitOptions) String() string {
	switch e {
	case MTLAccelerationStructureRefitOptionPerPrimitiveData:
		return "MTLAccelerationStructureRefitOptionPerPrimitiveData"
	case MTLAccelerationStructureRefitOptionVertexData:
		return "MTLAccelerationStructureRefitOptionVertexData"
	default:
		return fmt.Sprintf("MTLAccelerationStructureRefitOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureUsage
type MTLAccelerationStructureUsage uint

const (
	// MTLAccelerationStructureUsageExtendedLimits: An option that increases an acceleration structure’s storage capacity.
	MTLAccelerationStructureUsageExtendedLimits MTLAccelerationStructureUsage = 4
	// MTLAccelerationStructureUsageMinimizeMemory: An option that instructs Metal to prioritize building an acceleration structure that needs less memory.
	MTLAccelerationStructureUsageMinimizeMemory MTLAccelerationStructureUsage = 32
	// MTLAccelerationStructureUsageNone: A sentinel value the represents an empty set of options, which is the default behavior for building new acceleration structures.
	MTLAccelerationStructureUsageNone MTLAccelerationStructureUsage = 0
	// MTLAccelerationStructureUsagePreferFastBuild: An option that instructs Metal to build an acceleration structure quickly.
	MTLAccelerationStructureUsagePreferFastBuild MTLAccelerationStructureUsage = 2
	// MTLAccelerationStructureUsagePreferFastIntersection: An option that instructs Metal to prioritize building an acceleration structure with better intersection performance.
	MTLAccelerationStructureUsagePreferFastIntersection MTLAccelerationStructureUsage = 16
	// MTLAccelerationStructureUsageRefit: An option that lets you update an acceleration structure after creating it.
	MTLAccelerationStructureUsageRefit MTLAccelerationStructureUsage = 1
)

func (e MTLAccelerationStructureUsage) String() string {
	switch e {
	case MTLAccelerationStructureUsageExtendedLimits:
		return "MTLAccelerationStructureUsageExtendedLimits"
	case MTLAccelerationStructureUsageMinimizeMemory:
		return "MTLAccelerationStructureUsageMinimizeMemory"
	case MTLAccelerationStructureUsageNone:
		return "MTLAccelerationStructureUsageNone"
	case MTLAccelerationStructureUsagePreferFastBuild:
		return "MTLAccelerationStructureUsagePreferFastBuild"
	case MTLAccelerationStructureUsagePreferFastIntersection:
		return "MTLAccelerationStructureUsagePreferFastIntersection"
	case MTLAccelerationStructureUsageRefit:
		return "MTLAccelerationStructureUsageRefit"
	default:
		return fmt.Sprintf("MTLAccelerationStructureUsage(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLArgumentBuffersTier
type MTLArgumentBuffersTier uint

const (
	// MTLArgumentBuffersTier1: Support for tier 1 argument buffers.
	MTLArgumentBuffersTier1 MTLArgumentBuffersTier = 0
	// MTLArgumentBuffersTier2: Support for tier 2 argument buffers.
	MTLArgumentBuffersTier2 MTLArgumentBuffersTier = 1
)

func (e MTLArgumentBuffersTier) String() string {
	switch e {
	case MTLArgumentBuffersTier1:
		return "MTLArgumentBuffersTier1"
	case MTLArgumentBuffersTier2:
		return "MTLArgumentBuffersTier2"
	default:
		return fmt.Sprintf("MTLArgumentBuffersTier(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLArgumentType
type MTLArgumentType uint

const (
	// MTLArgumentTypeBuffer: The argument is a buffer.
	MTLArgumentTypeBuffer MTLArgumentType = 0
	// MTLArgumentTypeImageblock: The argument is an imageblock.
	MTLArgumentTypeImageblock MTLArgumentType = 17
	// MTLArgumentTypeImageblockData: The argument is imageblock data.
	MTLArgumentTypeImageblockData MTLArgumentType = 16
	// Deprecated.
	MTLArgumentTypeInstanceAccelerationStructure MTLArgumentType = 26
	// Deprecated.
	MTLArgumentTypeIntersectionFunctionTable MTLArgumentType = 27
	// Deprecated.
	MTLArgumentTypePrimitiveAccelerationStructure MTLArgumentType = 25
	// Deprecated.
	MTLArgumentTypeSampler MTLArgumentType = 3
	// Deprecated.
	MTLArgumentTypeTexture MTLArgumentType = 2
	// Deprecated.
	MTLArgumentTypeThreadgroupMemory MTLArgumentType = 1
	// Deprecated.
	MTLArgumentTypeVisibleFunctionTable MTLArgumentType = 24
)

func (e MTLArgumentType) String() string {
	switch e {
	case MTLArgumentTypeBuffer:
		return "MTLArgumentTypeBuffer"
	case MTLArgumentTypeImageblock:
		return "MTLArgumentTypeImageblock"
	case MTLArgumentTypeImageblockData:
		return "MTLArgumentTypeImageblockData"
	case MTLArgumentTypeInstanceAccelerationStructure:
		return "MTLArgumentTypeInstanceAccelerationStructure"
	case MTLArgumentTypeIntersectionFunctionTable:
		return "MTLArgumentTypeIntersectionFunctionTable"
	case MTLArgumentTypePrimitiveAccelerationStructure:
		return "MTLArgumentTypePrimitiveAccelerationStructure"
	case MTLArgumentTypeSampler:
		return "MTLArgumentTypeSampler"
	case MTLArgumentTypeTexture:
		return "MTLArgumentTypeTexture"
	case MTLArgumentTypeThreadgroupMemory:
		return "MTLArgumentTypeThreadgroupMemory"
	case MTLArgumentTypeVisibleFunctionTable:
		return "MTLArgumentTypeVisibleFunctionTable"
	default:
		return fmt.Sprintf("MTLArgumentType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLAttributeFormat
type MTLAttributeFormat uint

const (
	// MTLAttributeFormatChar: One signed 8-bit two’s complement value.
	MTLAttributeFormatChar MTLAttributeFormat = 46
	// MTLAttributeFormatChar2: Two signed 8-bit two’s complement values.
	MTLAttributeFormatChar2 MTLAttributeFormat = 4
	// MTLAttributeFormatChar2Normalized: Two signed normalized 8-bit two’s complement values.
	MTLAttributeFormatChar2Normalized MTLAttributeFormat = 10
	// MTLAttributeFormatChar3: Three signed 8-bit two’s complement values.
	MTLAttributeFormatChar3 MTLAttributeFormat = 5
	// MTLAttributeFormatChar3Normalized: Three signed normalized 8-bit two’s complement values.
	MTLAttributeFormatChar3Normalized MTLAttributeFormat = 11
	// MTLAttributeFormatChar4: Four signed 8-bit two’s complement values.
	MTLAttributeFormatChar4 MTLAttributeFormat = 6
	// MTLAttributeFormatChar4Normalized: Four signed normalized 8-bit two’s complement values.
	MTLAttributeFormatChar4Normalized MTLAttributeFormat = 12
	// MTLAttributeFormatCharNormalized: One signed normalized 8-bit two’s complement value.
	MTLAttributeFormatCharNormalized MTLAttributeFormat = 48
	// MTLAttributeFormatFloat: One single-precision floating-point value.
	MTLAttributeFormatFloat MTLAttributeFormat = 28
	// MTLAttributeFormatFloat2: Two single-precision floating-point values.
	MTLAttributeFormatFloat2 MTLAttributeFormat = 29
	// MTLAttributeFormatFloat3: Three single-precision floating-point values.
	MTLAttributeFormatFloat3 MTLAttributeFormat = 30
	// MTLAttributeFormatFloat4: Four single-precision floating-point values.
	MTLAttributeFormatFloat4 MTLAttributeFormat = 31
	// MTLAttributeFormatFloatRG11B10: One packed 32-bit value representing pixel data containing 11-bit float red and green channels, and a 10-bit float blue channel.
	MTLAttributeFormatFloatRG11B10 MTLAttributeFormat = 54
	// MTLAttributeFormatFloatRGB9E5: One packed 32-bit value representing pixel data containing 9-bit float red, green, and blue channels, and a 5-bit float shared exponent channel.
	MTLAttributeFormatFloatRGB9E5 MTLAttributeFormat = 55
	// MTLAttributeFormatHalf: One half-precision floating-point value.
	MTLAttributeFormatHalf MTLAttributeFormat = 53
	// MTLAttributeFormatHalf2: Two half-precision floating-point values.
	MTLAttributeFormatHalf2 MTLAttributeFormat = 25
	// MTLAttributeFormatHalf3: Three half-precision floating-point values.
	MTLAttributeFormatHalf3 MTLAttributeFormat = 26
	// MTLAttributeFormatHalf4: Four half-precision floating-point values.
	MTLAttributeFormatHalf4 MTLAttributeFormat = 27
	// MTLAttributeFormatInt: A 32-bit, signed integer value.
	MTLAttributeFormatInt MTLAttributeFormat = 32
	// MTLAttributeFormatInt1010102Normalized: One packed 32-bit value with four normalized signed two’s complement integer values, arranged as 10 bits, 10 bits, 10 bits, and 2 bits.
	MTLAttributeFormatInt1010102Normalized MTLAttributeFormat = 40
	// MTLAttributeFormatInt2: A two-component vector with 32-bit, signed integer values.
	MTLAttributeFormatInt2 MTLAttributeFormat = 33
	// MTLAttributeFormatInt3: A three-component vector with 32-bit, signed integer values.
	MTLAttributeFormatInt3 MTLAttributeFormat = 34
	// MTLAttributeFormatInt4: A four-component vector with 32-bit, signed integer values.
	MTLAttributeFormatInt4 MTLAttributeFormat = 35
	// MTLAttributeFormatInvalid: A sentinel value that represents an invalid attribute format.
	MTLAttributeFormatInvalid MTLAttributeFormat = 0
	// MTLAttributeFormatShort: A 16-bit, signed integer value.
	MTLAttributeFormatShort MTLAttributeFormat = 50
	// MTLAttributeFormatShort2: A two-component vector with 16-bit, signed integer values.
	MTLAttributeFormatShort2 MTLAttributeFormat = 16
	// MTLAttributeFormatShort2Normalized: A two-component vector with 16-bit, normalized, signed integer values.
	MTLAttributeFormatShort2Normalized MTLAttributeFormat = 22
	// MTLAttributeFormatShort3: A three-component vector with 16-bit, signed integer values.
	MTLAttributeFormatShort3 MTLAttributeFormat = 17
	// MTLAttributeFormatShort3Normalized: A three-component vector with 16-bit, normalized, signed integer values.
	MTLAttributeFormatShort3Normalized MTLAttributeFormat = 23
	// MTLAttributeFormatShort4: A four-component vector with 16-bit, signed integer values.
	MTLAttributeFormatShort4 MTLAttributeFormat = 18
	// MTLAttributeFormatShort4Normalized: A four-component vector with 16-bit, normalized, signed integer values.
	MTLAttributeFormatShort4Normalized MTLAttributeFormat = 24
	// MTLAttributeFormatShortNormalized: A 16-bit, normalized, signed integer value.
	MTLAttributeFormatShortNormalized MTLAttributeFormat = 52
	// MTLAttributeFormatUChar: An 8-bit, unsigned integer value.
	MTLAttributeFormatUChar MTLAttributeFormat = 45
	// MTLAttributeFormatUChar2: A two-component vector with 8-bit, unsigned integer values.
	MTLAttributeFormatUChar2 MTLAttributeFormat = 1
	// MTLAttributeFormatUChar2Normalized: A two-component vector with 8-bit, normalized, unsigned integer values.
	MTLAttributeFormatUChar2Normalized MTLAttributeFormat = 7
	// MTLAttributeFormatUChar3: A three-component vector with 8-bit, unsigned integer values.
	MTLAttributeFormatUChar3 MTLAttributeFormat = 2
	// MTLAttributeFormatUChar3Normalized: A three-component vector with 8-bit, normalized, unsigned integer values.
	MTLAttributeFormatUChar3Normalized MTLAttributeFormat = 8
	// MTLAttributeFormatUChar4: A four-component vector with 8-bit, unsigned integer values.
	MTLAttributeFormatUChar4 MTLAttributeFormat = 3
	// MTLAttributeFormatUChar4Normalized: A four-component vector with 8-bit, normalized, unsigned integer values.
	MTLAttributeFormatUChar4Normalized MTLAttributeFormat = 9
	// MTLAttributeFormatUChar4Normalized_BGRA: Four unsigned normalized 8-bit values, arranged as blue, green, red, and alpha components.
	MTLAttributeFormatUChar4Normalized_BGRA MTLAttributeFormat = 42
	// MTLAttributeFormatUCharNormalized: An 8-bit, normalized, unsigned integer value.
	MTLAttributeFormatUCharNormalized MTLAttributeFormat = 47
	// MTLAttributeFormatUInt: A 32-bit, unsigned integer value.
	MTLAttributeFormatUInt MTLAttributeFormat = 36
	// MTLAttributeFormatUInt1010102Normalized: One packed 32-bit value with four normalized unsigned integer values, arranged as 10 bits, 10 bits, 10 bits, and 2 bits.
	MTLAttributeFormatUInt1010102Normalized MTLAttributeFormat = 41
	// MTLAttributeFormatUInt2: A two-component vector with 32-bit, unsigned integer values.
	MTLAttributeFormatUInt2 MTLAttributeFormat = 37
	// MTLAttributeFormatUInt3: A three-component vector with 32-bit, unsigned integer values.
	MTLAttributeFormatUInt3 MTLAttributeFormat = 38
	// MTLAttributeFormatUInt4: A four-component vector with 32-bit, unsigned integer values.
	MTLAttributeFormatUInt4 MTLAttributeFormat = 39
	// MTLAttributeFormatUShort: A 16-bit, unsigned integer value.
	MTLAttributeFormatUShort MTLAttributeFormat = 49
	// MTLAttributeFormatUShort2: A two-component vector with 16-bit, unsigned integer values.
	MTLAttributeFormatUShort2 MTLAttributeFormat = 13
	// MTLAttributeFormatUShort2Normalized: Two unsigned normalized 16-bit values
	MTLAttributeFormatUShort2Normalized MTLAttributeFormat = 19
	// MTLAttributeFormatUShort3: A three-component vector with 16-bit, unsigned integer values.
	MTLAttributeFormatUShort3 MTLAttributeFormat = 14
	// MTLAttributeFormatUShort3Normalized: A three-component vector with 16-bit, normalized, unsigned integer values.
	MTLAttributeFormatUShort3Normalized MTLAttributeFormat = 20
	// MTLAttributeFormatUShort4: A four-component vector with 16-bit, unsigned integer values.
	MTLAttributeFormatUShort4 MTLAttributeFormat = 15
	// MTLAttributeFormatUShort4Normalized: A four-component vector with 16-bit, normalized, unsigned integer values.
	MTLAttributeFormatUShort4Normalized MTLAttributeFormat = 21
	// MTLAttributeFormatUShortNormalized: A 16-bit, normalized, unsigned integer value.
	MTLAttributeFormatUShortNormalized MTLAttributeFormat = 51
)

func (e MTLAttributeFormat) String() string {
	switch e {
	case MTLAttributeFormatChar:
		return "MTLAttributeFormatChar"
	case MTLAttributeFormatChar2:
		return "MTLAttributeFormatChar2"
	case MTLAttributeFormatChar2Normalized:
		return "MTLAttributeFormatChar2Normalized"
	case MTLAttributeFormatChar3:
		return "MTLAttributeFormatChar3"
	case MTLAttributeFormatChar3Normalized:
		return "MTLAttributeFormatChar3Normalized"
	case MTLAttributeFormatChar4:
		return "MTLAttributeFormatChar4"
	case MTLAttributeFormatChar4Normalized:
		return "MTLAttributeFormatChar4Normalized"
	case MTLAttributeFormatCharNormalized:
		return "MTLAttributeFormatCharNormalized"
	case MTLAttributeFormatFloat:
		return "MTLAttributeFormatFloat"
	case MTLAttributeFormatFloat2:
		return "MTLAttributeFormatFloat2"
	case MTLAttributeFormatFloat3:
		return "MTLAttributeFormatFloat3"
	case MTLAttributeFormatFloat4:
		return "MTLAttributeFormatFloat4"
	case MTLAttributeFormatFloatRG11B10:
		return "MTLAttributeFormatFloatRG11B10"
	case MTLAttributeFormatFloatRGB9E5:
		return "MTLAttributeFormatFloatRGB9E5"
	case MTLAttributeFormatHalf:
		return "MTLAttributeFormatHalf"
	case MTLAttributeFormatHalf2:
		return "MTLAttributeFormatHalf2"
	case MTLAttributeFormatHalf3:
		return "MTLAttributeFormatHalf3"
	case MTLAttributeFormatHalf4:
		return "MTLAttributeFormatHalf4"
	case MTLAttributeFormatInt:
		return "MTLAttributeFormatInt"
	case MTLAttributeFormatInt1010102Normalized:
		return "MTLAttributeFormatInt1010102Normalized"
	case MTLAttributeFormatInt2:
		return "MTLAttributeFormatInt2"
	case MTLAttributeFormatInt3:
		return "MTLAttributeFormatInt3"
	case MTLAttributeFormatInt4:
		return "MTLAttributeFormatInt4"
	case MTLAttributeFormatInvalid:
		return "MTLAttributeFormatInvalid"
	case MTLAttributeFormatShort:
		return "MTLAttributeFormatShort"
	case MTLAttributeFormatShort2:
		return "MTLAttributeFormatShort2"
	case MTLAttributeFormatShort2Normalized:
		return "MTLAttributeFormatShort2Normalized"
	case MTLAttributeFormatShort3:
		return "MTLAttributeFormatShort3"
	case MTLAttributeFormatShort3Normalized:
		return "MTLAttributeFormatShort3Normalized"
	case MTLAttributeFormatShort4:
		return "MTLAttributeFormatShort4"
	case MTLAttributeFormatShort4Normalized:
		return "MTLAttributeFormatShort4Normalized"
	case MTLAttributeFormatShortNormalized:
		return "MTLAttributeFormatShortNormalized"
	case MTLAttributeFormatUChar:
		return "MTLAttributeFormatUChar"
	case MTLAttributeFormatUChar2:
		return "MTLAttributeFormatUChar2"
	case MTLAttributeFormatUChar2Normalized:
		return "MTLAttributeFormatUChar2Normalized"
	case MTLAttributeFormatUChar3:
		return "MTLAttributeFormatUChar3"
	case MTLAttributeFormatUChar3Normalized:
		return "MTLAttributeFormatUChar3Normalized"
	case MTLAttributeFormatUChar4:
		return "MTLAttributeFormatUChar4"
	case MTLAttributeFormatUChar4Normalized:
		return "MTLAttributeFormatUChar4Normalized"
	case MTLAttributeFormatUChar4Normalized_BGRA:
		return "MTLAttributeFormatUChar4Normalized_BGRA"
	case MTLAttributeFormatUCharNormalized:
		return "MTLAttributeFormatUCharNormalized"
	case MTLAttributeFormatUInt:
		return "MTLAttributeFormatUInt"
	case MTLAttributeFormatUInt1010102Normalized:
		return "MTLAttributeFormatUInt1010102Normalized"
	case MTLAttributeFormatUInt2:
		return "MTLAttributeFormatUInt2"
	case MTLAttributeFormatUInt3:
		return "MTLAttributeFormatUInt3"
	case MTLAttributeFormatUInt4:
		return "MTLAttributeFormatUInt4"
	case MTLAttributeFormatUShort:
		return "MTLAttributeFormatUShort"
	case MTLAttributeFormatUShort2:
		return "MTLAttributeFormatUShort2"
	case MTLAttributeFormatUShort2Normalized:
		return "MTLAttributeFormatUShort2Normalized"
	case MTLAttributeFormatUShort3:
		return "MTLAttributeFormatUShort3"
	case MTLAttributeFormatUShort3Normalized:
		return "MTLAttributeFormatUShort3Normalized"
	case MTLAttributeFormatUShort4:
		return "MTLAttributeFormatUShort4"
	case MTLAttributeFormatUShort4Normalized:
		return "MTLAttributeFormatUShort4Normalized"
	case MTLAttributeFormatUShortNormalized:
		return "MTLAttributeFormatUShortNormalized"
	default:
		return fmt.Sprintf("MTLAttributeFormat(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLBarrierScope
type MTLBarrierScope uint

const (
	// MTLBarrierScopeBuffers: The barrier affects any buffer objects.
	MTLBarrierScopeBuffers MTLBarrierScope = 1
	// MTLBarrierScopeRenderTargets: The barrier affects any render targets.
	MTLBarrierScopeRenderTargets MTLBarrierScope = 4
	// MTLBarrierScopeTextures: The barrier affects textures.
	MTLBarrierScopeTextures MTLBarrierScope = 2
)

func (e MTLBarrierScope) String() string {
	switch e {
	case MTLBarrierScopeBuffers:
		return "MTLBarrierScopeBuffers"
	case MTLBarrierScopeRenderTargets:
		return "MTLBarrierScopeRenderTargets"
	case MTLBarrierScopeTextures:
		return "MTLBarrierScopeTextures"
	default:
		return fmt.Sprintf("MTLBarrierScope(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLBinaryArchiveError-swift.struct/Code
type MTLBinaryArchiveError int

const (
	// MTLBinaryArchiveErrorCompilationFailure: An error code that indicates the archive’s inability to compile its contents, typically when serializing it to a URL.
	MTLBinaryArchiveErrorCompilationFailure MTLBinaryArchiveError = 3
	// MTLBinaryArchiveErrorInternalError: An error code that indicates the Metal framework has an internal problem.
	MTLBinaryArchiveErrorInternalError MTLBinaryArchiveError = 4
	// MTLBinaryArchiveErrorInvalidFile: An error code that indicates an app is using an invalid reference to an archive file, typically related to a URL.
	MTLBinaryArchiveErrorInvalidFile MTLBinaryArchiveError = 1
	// MTLBinaryArchiveErrorNone: An error code that represents the absence of any problems.
	MTLBinaryArchiveErrorNone MTLBinaryArchiveError = 0
	// MTLBinaryArchiveErrorUnexpectedElement: An error code that indicates a problem with a configuration, typically in a descriptor or an archive’s inability to add linked functions.
	MTLBinaryArchiveErrorUnexpectedElement MTLBinaryArchiveError = 2
)

func (e MTLBinaryArchiveError) String() string {
	switch e {
	case MTLBinaryArchiveErrorCompilationFailure:
		return "MTLBinaryArchiveErrorCompilationFailure"
	case MTLBinaryArchiveErrorInternalError:
		return "MTLBinaryArchiveErrorInternalError"
	case MTLBinaryArchiveErrorInvalidFile:
		return "MTLBinaryArchiveErrorInvalidFile"
	case MTLBinaryArchiveErrorNone:
		return "MTLBinaryArchiveErrorNone"
	case MTLBinaryArchiveErrorUnexpectedElement:
		return "MTLBinaryArchiveErrorUnexpectedElement"
	default:
		return fmt.Sprintf("MTLBinaryArchiveError(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLBindingAccess
type MTLBindingAccess uint

const (
	// MTLArgumentAccessReadOnly: The function can only read its argument data.
	MTLArgumentAccessReadOnly MTLBindingAccess = 0
	// MTLArgumentAccessReadWrite: The function can either read or write its argument data.
	MTLArgumentAccessReadWrite MTLBindingAccess = 1
	// MTLArgumentAccessWriteOnly: The function can only write its argument data.
	MTLArgumentAccessWriteOnly MTLBindingAccess = 2
	MTLBindingAccessReadOnly   MTLBindingAccess = 0
	MTLBindingAccessReadWrite  MTLBindingAccess = 1
	MTLBindingAccessWriteOnly  MTLBindingAccess = 2
)

func (e MTLBindingAccess) String() string {
	switch e {
	case MTLArgumentAccessReadOnly:
		return "MTLArgumentAccessReadOnly"
	case MTLArgumentAccessReadWrite:
		return "MTLArgumentAccessReadWrite"
	case MTLArgumentAccessWriteOnly:
		return "MTLArgumentAccessWriteOnly"
	default:
		return fmt.Sprintf("MTLBindingAccess(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLBindingType
type MTLBindingType int

const (
	MTLBindingTypeBuffer                         MTLBindingType = 0
	MTLBindingTypeImageblock                     MTLBindingType = 17
	MTLBindingTypeImageblockData                 MTLBindingType = 16
	MTLBindingTypeInstanceAccelerationStructure  MTLBindingType = 26
	MTLBindingTypeIntersectionFunctionTable      MTLBindingType = 27
	MTLBindingTypeObjectPayload                  MTLBindingType = 34
	MTLBindingTypePrimitiveAccelerationStructure MTLBindingType = 25
	MTLBindingTypeSampler                        MTLBindingType = 3
	MTLBindingTypeTensor                         MTLBindingType = 37
	MTLBindingTypeTexture                        MTLBindingType = 2
	MTLBindingTypeThreadgroupMemory              MTLBindingType = 1
	MTLBindingTypeVisibleFunctionTable           MTLBindingType = 24
)

func (e MTLBindingType) String() string {
	switch e {
	case MTLBindingTypeBuffer:
		return "MTLBindingTypeBuffer"
	case MTLBindingTypeImageblock:
		return "MTLBindingTypeImageblock"
	case MTLBindingTypeImageblockData:
		return "MTLBindingTypeImageblockData"
	case MTLBindingTypeInstanceAccelerationStructure:
		return "MTLBindingTypeInstanceAccelerationStructure"
	case MTLBindingTypeIntersectionFunctionTable:
		return "MTLBindingTypeIntersectionFunctionTable"
	case MTLBindingTypeObjectPayload:
		return "MTLBindingTypeObjectPayload"
	case MTLBindingTypePrimitiveAccelerationStructure:
		return "MTLBindingTypePrimitiveAccelerationStructure"
	case MTLBindingTypeSampler:
		return "MTLBindingTypeSampler"
	case MTLBindingTypeTensor:
		return "MTLBindingTypeTensor"
	case MTLBindingTypeTexture:
		return "MTLBindingTypeTexture"
	case MTLBindingTypeThreadgroupMemory:
		return "MTLBindingTypeThreadgroupMemory"
	case MTLBindingTypeVisibleFunctionTable:
		return "MTLBindingTypeVisibleFunctionTable"
	default:
		return fmt.Sprintf("MTLBindingType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLBlendFactor
type MTLBlendFactor uint

const (
	// MTLBlendFactorBlendAlpha: Blend factor of alpha value.
	MTLBlendFactorBlendAlpha MTLBlendFactor = 13
	// MTLBlendFactorBlendColor: A blend factor that applies the blend color’s red, green, and blue components.
	MTLBlendFactorBlendColor MTLBlendFactor = 11
	// MTLBlendFactorDestinationAlpha: Blend factor of destination alpha.
	MTLBlendFactorDestinationAlpha MTLBlendFactor = 8
	// MTLBlendFactorDestinationColor: Blend factor of destination values.
	MTLBlendFactorDestinationColor MTLBlendFactor = 6
	// MTLBlendFactorOne: Blend factor of one.
	MTLBlendFactorOne MTLBlendFactor = 1
	// MTLBlendFactorOneMinusBlendAlpha: Blend factor of one minus alpha value.
	MTLBlendFactorOneMinusBlendAlpha MTLBlendFactor = 14
	// MTLBlendFactorOneMinusBlendColor: A blend factor that applies one minus the blend color’s red, green, and blue components.
	MTLBlendFactorOneMinusBlendColor MTLBlendFactor = 12
	// MTLBlendFactorOneMinusDestinationAlpha: Blend factor of one minus destination alpha.
	MTLBlendFactorOneMinusDestinationAlpha MTLBlendFactor = 9
	// MTLBlendFactorOneMinusDestinationColor: Blend factor of one minus destination values.
	MTLBlendFactorOneMinusDestinationColor MTLBlendFactor = 7
	// MTLBlendFactorOneMinusSource1Alpha: Blend factor of one minus source alpha.
	MTLBlendFactorOneMinusSource1Alpha MTLBlendFactor = 18
	// MTLBlendFactorOneMinusSource1Color: Blend factor of one minus source values.
	MTLBlendFactorOneMinusSource1Color MTLBlendFactor = 16
	// MTLBlendFactorOneMinusSourceAlpha: Blend factor of one minus source alpha.
	MTLBlendFactorOneMinusSourceAlpha MTLBlendFactor = 5
	// MTLBlendFactorOneMinusSourceColor: Blend factor of one minus source values.
	MTLBlendFactorOneMinusSourceColor MTLBlendFactor = 3
	// MTLBlendFactorSource1Alpha: Blend factor of source alpha.
	MTLBlendFactorSource1Alpha MTLBlendFactor = 17
	// MTLBlendFactorSource1Color: Blend factor of source values.
	MTLBlendFactorSource1Color MTLBlendFactor = 15
	// MTLBlendFactorSourceAlpha: Blend factor of source alpha.
	MTLBlendFactorSourceAlpha MTLBlendFactor = 4
	// MTLBlendFactorSourceAlphaSaturated: Blend factor of the minimum of either source alpha or one minus destination alpha.
	MTLBlendFactorSourceAlphaSaturated MTLBlendFactor = 10
	// MTLBlendFactorSourceColor: Blend factor of source values.
	MTLBlendFactorSourceColor MTLBlendFactor = 2
	// MTLBlendFactorUnspecialized: Defers assigning the blend factor.
	MTLBlendFactorUnspecialized MTLBlendFactor = 19
	// MTLBlendFactorZero: Blend factor of zero.
	MTLBlendFactorZero MTLBlendFactor = 0
)

func (e MTLBlendFactor) String() string {
	switch e {
	case MTLBlendFactorBlendAlpha:
		return "MTLBlendFactorBlendAlpha"
	case MTLBlendFactorBlendColor:
		return "MTLBlendFactorBlendColor"
	case MTLBlendFactorDestinationAlpha:
		return "MTLBlendFactorDestinationAlpha"
	case MTLBlendFactorDestinationColor:
		return "MTLBlendFactorDestinationColor"
	case MTLBlendFactorOne:
		return "MTLBlendFactorOne"
	case MTLBlendFactorOneMinusBlendAlpha:
		return "MTLBlendFactorOneMinusBlendAlpha"
	case MTLBlendFactorOneMinusBlendColor:
		return "MTLBlendFactorOneMinusBlendColor"
	case MTLBlendFactorOneMinusDestinationAlpha:
		return "MTLBlendFactorOneMinusDestinationAlpha"
	case MTLBlendFactorOneMinusDestinationColor:
		return "MTLBlendFactorOneMinusDestinationColor"
	case MTLBlendFactorOneMinusSource1Alpha:
		return "MTLBlendFactorOneMinusSource1Alpha"
	case MTLBlendFactorOneMinusSource1Color:
		return "MTLBlendFactorOneMinusSource1Color"
	case MTLBlendFactorOneMinusSourceAlpha:
		return "MTLBlendFactorOneMinusSourceAlpha"
	case MTLBlendFactorOneMinusSourceColor:
		return "MTLBlendFactorOneMinusSourceColor"
	case MTLBlendFactorSource1Alpha:
		return "MTLBlendFactorSource1Alpha"
	case MTLBlendFactorSource1Color:
		return "MTLBlendFactorSource1Color"
	case MTLBlendFactorSourceAlpha:
		return "MTLBlendFactorSourceAlpha"
	case MTLBlendFactorSourceAlphaSaturated:
		return "MTLBlendFactorSourceAlphaSaturated"
	case MTLBlendFactorSourceColor:
		return "MTLBlendFactorSourceColor"
	case MTLBlendFactorUnspecialized:
		return "MTLBlendFactorUnspecialized"
	case MTLBlendFactorZero:
		return "MTLBlendFactorZero"
	default:
		return fmt.Sprintf("MTLBlendFactor(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLBlendOperation
type MTLBlendOperation uint

const (
	// MTLBlendOperationAdd: Add portions of both source and destination pixel values.
	MTLBlendOperationAdd MTLBlendOperation = 0
	// MTLBlendOperationMax: Maximum of the source and destination pixel values.
	MTLBlendOperationMax MTLBlendOperation = 4
	// MTLBlendOperationMin: Minimum of the source and destination pixel values.
	MTLBlendOperationMin MTLBlendOperation = 3
	// MTLBlendOperationReverseSubtract: Subtract a portion of the source values from a portion of the destination pixel values.
	MTLBlendOperationReverseSubtract MTLBlendOperation = 2
	// MTLBlendOperationSubtract: Subtract a portion of the destination pixel values from a portion of the source.
	MTLBlendOperationSubtract MTLBlendOperation = 1
	// MTLBlendOperationUnspecialized: Defers assigning the blend operation.
	MTLBlendOperationUnspecialized MTLBlendOperation = 5
)

func (e MTLBlendOperation) String() string {
	switch e {
	case MTLBlendOperationAdd:
		return "MTLBlendOperationAdd"
	case MTLBlendOperationMax:
		return "MTLBlendOperationMax"
	case MTLBlendOperationMin:
		return "MTLBlendOperationMin"
	case MTLBlendOperationReverseSubtract:
		return "MTLBlendOperationReverseSubtract"
	case MTLBlendOperationSubtract:
		return "MTLBlendOperationSubtract"
	case MTLBlendOperationUnspecialized:
		return "MTLBlendOperationUnspecialized"
	default:
		return fmt.Sprintf("MTLBlendOperation(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLBlitOption
type MTLBlitOption uint

const (
	// MTLBlitOptionDepthFromDepthStencil: A blit option that copies the depth portion of a combined depth and stencil texture to or from a buffer.
	MTLBlitOptionDepthFromDepthStencil MTLBlitOption = 1
	// MTLBlitOptionNone: A blit option that clears other blit options, which removes any optional behavior for a blit operation.
	MTLBlitOptionNone MTLBlitOption = 0
	// MTLBlitOptionRowLinearPVRTC: A blit option that copies PVRTC data between a texture and a buffer.
	MTLBlitOptionRowLinearPVRTC MTLBlitOption = 4
	// MTLBlitOptionStencilFromDepthStencil: A blit option that copies the stencil portion of a combined depth and stencil texture to or from a buffer.
	MTLBlitOptionStencilFromDepthStencil MTLBlitOption = 2
)

func (e MTLBlitOption) String() string {
	switch e {
	case MTLBlitOptionDepthFromDepthStencil:
		return "MTLBlitOptionDepthFromDepthStencil"
	case MTLBlitOptionNone:
		return "MTLBlitOptionNone"
	case MTLBlitOptionRowLinearPVRTC:
		return "MTLBlitOptionRowLinearPVRTC"
	case MTLBlitOptionStencilFromDepthStencil:
		return "MTLBlitOptionStencilFromDepthStencil"
	default:
		return fmt.Sprintf("MTLBlitOption(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLBufferSparseTier
type MTLBufferSparseTier int

const (
	// MTLBufferSparseTier1: Indicates support for sparse buffers tier 1.
	MTLBufferSparseTier1 MTLBufferSparseTier = 1
	// MTLBufferSparseTierNone: Indicates that the buffer is not sparse.
	MTLBufferSparseTierNone MTLBufferSparseTier = 0
)

func (e MTLBufferSparseTier) String() string {
	switch e {
	case MTLBufferSparseTier1:
		return "MTLBufferSparseTier1"
	case MTLBufferSparseTierNone:
		return "MTLBufferSparseTierNone"
	default:
		return fmt.Sprintf("MTLBufferSparseTier(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLCPUCacheMode
type MTLCPUCacheMode uint

const (
	// MTLCPUCacheModeDefaultCache: The default CPU cache mode for the resource, which guarantees that read and write operations are executed in the expected order.
	MTLCPUCacheModeDefaultCache MTLCPUCacheMode = 0
	// MTLCPUCacheModeWriteCombined: A write-combined CPU cache mode that is optimized for resources that the CPU writes into, but never reads.
	MTLCPUCacheModeWriteCombined MTLCPUCacheMode = 1
)

func (e MTLCPUCacheMode) String() string {
	switch e {
	case MTLCPUCacheModeDefaultCache:
		return "MTLCPUCacheModeDefaultCache"
	case MTLCPUCacheModeWriteCombined:
		return "MTLCPUCacheModeWriteCombined"
	default:
		return fmt.Sprintf("MTLCPUCacheMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLCaptureDestination
type MTLCaptureDestination int

const (
	// MTLCaptureDestinationDeveloperTools: An option specifying that data should be captured to Xcode and that execution should stop in Xcode after the data is captured.
	MTLCaptureDestinationDeveloperTools MTLCaptureDestination = 1
	// MTLCaptureDestinationGPUTraceDocument: An option specifying that the captured command data should be saved to a GPU trace document.
	MTLCaptureDestinationGPUTraceDocument MTLCaptureDestination = 2
)

func (e MTLCaptureDestination) String() string {
	switch e {
	case MTLCaptureDestinationDeveloperTools:
		return "MTLCaptureDestinationDeveloperTools"
	case MTLCaptureDestinationGPUTraceDocument:
		return "MTLCaptureDestinationGPUTraceDocument"
	default:
		return fmt.Sprintf("MTLCaptureDestination(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLCaptureError
type MTLCaptureError int

const (
	// MTLCaptureErrorAlreadyCapturing: A capture error that indicates the session is already in progress.
	MTLCaptureErrorAlreadyCapturing MTLCaptureError = 2
	// MTLCaptureErrorInvalidDescriptor: A capture error that indicates your descriptor has invalid properties.
	MTLCaptureErrorInvalidDescriptor MTLCaptureError = 3
	// MTLCaptureErrorNotSupported: A capture error that indicates the capture options you’re requesting aren’t available.
	MTLCaptureErrorNotSupported MTLCaptureError = 1
)

func (e MTLCaptureError) String() string {
	switch e {
	case MTLCaptureErrorAlreadyCapturing:
		return "MTLCaptureErrorAlreadyCapturing"
	case MTLCaptureErrorInvalidDescriptor:
		return "MTLCaptureErrorInvalidDescriptor"
	case MTLCaptureErrorNotSupported:
		return "MTLCaptureErrorNotSupported"
	default:
		return fmt.Sprintf("MTLCaptureError(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLColorWriteMask
type MTLColorWriteMask uint

const (
	// MTLColorWriteMaskAll: All color channels are enabled.
	MTLColorWriteMaskAll MTLColorWriteMask = 15
	// MTLColorWriteMaskAlpha: The alpha color channel is enabled.
	MTLColorWriteMaskAlpha MTLColorWriteMask = 1
	// MTLColorWriteMaskBlue: The blue color channel is enabled.
	MTLColorWriteMaskBlue MTLColorWriteMask = 1
	// MTLColorWriteMaskGreen: The green color channel is enabled.
	MTLColorWriteMaskGreen MTLColorWriteMask = 1
	// MTLColorWriteMaskNone: All color channels are disabled.
	MTLColorWriteMaskNone MTLColorWriteMask = 0
	// MTLColorWriteMaskRed: The red color channel is enabled.
	MTLColorWriteMaskRed MTLColorWriteMask = 1
	// MTLColorWriteMaskUnspecialized: Defers assigning the color write mask.
	MTLColorWriteMaskUnspecialized MTLColorWriteMask = 16
)

func (e MTLColorWriteMask) String() string {
	switch e {
	case MTLColorWriteMaskAll:
		return "MTLColorWriteMaskAll"
	case MTLColorWriteMaskAlpha:
		return "MTLColorWriteMaskAlpha"
	case MTLColorWriteMaskNone:
		return "MTLColorWriteMaskNone"
	case MTLColorWriteMaskUnspecialized:
		return "MTLColorWriteMaskUnspecialized"
	default:
		return fmt.Sprintf("MTLColorWriteMask(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLCommandBufferError-swift.struct/Code
type MTLCommandBufferError int

const (
	// MTLCommandBufferErrorAccessRevoked: An error code that indicates the system has revoked the Metal device’s access because it’s responsible for too many timeouts or hangs.
	MTLCommandBufferErrorAccessRevoked MTLCommandBufferError = 4
	// MTLCommandBufferErrorBlacklisted: A former error code that indicates the system has revoked the Metal device’s access because it’s responsible for too many timeouts or hangs.
	MTLCommandBufferErrorBlacklisted MTLCommandBufferError = 4
	// MTLCommandBufferErrorDeviceRemoved: An error code that indicates a person physically removed the GPU device before the command buffer finished running.
	MTLCommandBufferErrorDeviceRemoved MTLCommandBufferError = 11
	// MTLCommandBufferErrorInternal: An error code that indicates the Metal framework has an internal problem.
	MTLCommandBufferErrorInternal MTLCommandBufferError = 1
	// MTLCommandBufferErrorInvalidResource: An error code that indicates the command buffer has an invalid reference to resource.
	MTLCommandBufferErrorInvalidResource MTLCommandBufferError = 9
	// MTLCommandBufferErrorMemoryless: An error code that indicates the GPU ran out of one or more of its internal resources that support memoryless render pass attachments.
	MTLCommandBufferErrorMemoryless MTLCommandBufferError = 10
	// MTLCommandBufferErrorNone: An error code that represents the absence of any problems.
	MTLCommandBufferErrorNone MTLCommandBufferError = 0
	// MTLCommandBufferErrorNotPermitted: An error code that indicates a process doesn’t have access to a GPU device.
	MTLCommandBufferErrorNotPermitted MTLCommandBufferError = 7
	// MTLCommandBufferErrorOutOfMemory: An error code that indicates the GPU device doesn’t have sufficient memory to execute a command buffer.
	MTLCommandBufferErrorOutOfMemory MTLCommandBufferError = 8
	// MTLCommandBufferErrorPageFault: An error code that indicates the command buffer generated a page fault the GPU can’t service.
	MTLCommandBufferErrorPageFault MTLCommandBufferError = 3
	// MTLCommandBufferErrorStackOverflow: An error code that indicates the GPU terminated the command buffer because a kernel function of tile shader used too many stack frames.
	MTLCommandBufferErrorStackOverflow MTLCommandBufferError = 12
	// MTLCommandBufferErrorTimeout: An error code that indicates the system interrupted and terminated the command buffer before it finished running.
	MTLCommandBufferErrorTimeout MTLCommandBufferError = 2
)

func (e MTLCommandBufferError) String() string {
	switch e {
	case MTLCommandBufferErrorAccessRevoked:
		return "MTLCommandBufferErrorAccessRevoked"
	case MTLCommandBufferErrorDeviceRemoved:
		return "MTLCommandBufferErrorDeviceRemoved"
	case MTLCommandBufferErrorInternal:
		return "MTLCommandBufferErrorInternal"
	case MTLCommandBufferErrorInvalidResource:
		return "MTLCommandBufferErrorInvalidResource"
	case MTLCommandBufferErrorMemoryless:
		return "MTLCommandBufferErrorMemoryless"
	case MTLCommandBufferErrorNone:
		return "MTLCommandBufferErrorNone"
	case MTLCommandBufferErrorNotPermitted:
		return "MTLCommandBufferErrorNotPermitted"
	case MTLCommandBufferErrorOutOfMemory:
		return "MTLCommandBufferErrorOutOfMemory"
	case MTLCommandBufferErrorPageFault:
		return "MTLCommandBufferErrorPageFault"
	case MTLCommandBufferErrorStackOverflow:
		return "MTLCommandBufferErrorStackOverflow"
	case MTLCommandBufferErrorTimeout:
		return "MTLCommandBufferErrorTimeout"
	default:
		return fmt.Sprintf("MTLCommandBufferError(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLCommandBufferErrorOption
type MTLCommandBufferErrorOption int

const (
	// MTLCommandBufferErrorOptionEncoderExecutionStatus: An option that instructs a command buffer to save additional details about a GPU runtime error.
	MTLCommandBufferErrorOptionEncoderExecutionStatus MTLCommandBufferErrorOption = 1
	// MTLCommandBufferErrorOptionNone: An option that clears a command buffer’s error options.
	MTLCommandBufferErrorOptionNone MTLCommandBufferErrorOption = 0
)

func (e MTLCommandBufferErrorOption) String() string {
	switch e {
	case MTLCommandBufferErrorOptionEncoderExecutionStatus:
		return "MTLCommandBufferErrorOptionEncoderExecutionStatus"
	case MTLCommandBufferErrorOptionNone:
		return "MTLCommandBufferErrorOptionNone"
	default:
		return fmt.Sprintf("MTLCommandBufferErrorOption(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLCommandBufferStatus
type MTLCommandBufferStatus int

const (
	// MTLCommandBufferStatusCommitted: A command buffer’s third state, which indicates the command queue is preparing to schedule the command buffer by resolving its dependencies.
	MTLCommandBufferStatusCommitted MTLCommandBufferStatus = 2
	// MTLCommandBufferStatusCompleted: A command buffer’s successful, final state, which indicates the GPU finished running the command buffer’s commands without any problems.
	MTLCommandBufferStatusCompleted MTLCommandBufferStatus = 4
	// MTLCommandBufferStatusEnqueued: A command buffer’s second state, which indicates its command queue is reserving a place for it.
	MTLCommandBufferStatusEnqueued MTLCommandBufferStatus = 1
	// MTLCommandBufferStatusError: A command buffer’s unsuccessful, final state, which indicates the GPU stopped running the buffer’s commands because of a runtime issue.
	MTLCommandBufferStatusError MTLCommandBufferStatus = 5
	// MTLCommandBufferStatusNotEnqueued: A command buffer’s initial state, which indicates its command queue isn’t reserving a place for it.
	MTLCommandBufferStatusNotEnqueued MTLCommandBufferStatus = 0
	// MTLCommandBufferStatusScheduled: A command buffer’s fourth state, which indicates the command buffer has its resources ready and is waiting for the GPU to run its commands.
	MTLCommandBufferStatusScheduled MTLCommandBufferStatus = 3
)

func (e MTLCommandBufferStatus) String() string {
	switch e {
	case MTLCommandBufferStatusCommitted:
		return "MTLCommandBufferStatusCommitted"
	case MTLCommandBufferStatusCompleted:
		return "MTLCommandBufferStatusCompleted"
	case MTLCommandBufferStatusEnqueued:
		return "MTLCommandBufferStatusEnqueued"
	case MTLCommandBufferStatusError:
		return "MTLCommandBufferStatusError"
	case MTLCommandBufferStatusNotEnqueued:
		return "MTLCommandBufferStatusNotEnqueued"
	case MTLCommandBufferStatusScheduled:
		return "MTLCommandBufferStatusScheduled"
	default:
		return fmt.Sprintf("MTLCommandBufferStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLCommandEncoderErrorState
type MTLCommandEncoderErrorState int

const (
	// MTLCommandEncoderErrorStateAffected: An error state that indicates the GPU failed to fully execute the commands because of an error.
	MTLCommandEncoderErrorStateAffected MTLCommandEncoderErrorState = 2
	// MTLCommandEncoderErrorStateCompleted: A state that indicates the GPU successfully executed the commands without any errors.
	MTLCommandEncoderErrorStateCompleted MTLCommandEncoderErrorState = 1
	// MTLCommandEncoderErrorStateFaulted: An error state that indicates the commands in the command buffer are the cause of an error.
	MTLCommandEncoderErrorStateFaulted MTLCommandEncoderErrorState = 4
	// MTLCommandEncoderErrorStatePending: An error state that indicates the GPU didn’t execute the commands.
	MTLCommandEncoderErrorStatePending MTLCommandEncoderErrorState = 3
	// MTLCommandEncoderErrorStateUnknown: An error state that indicates the command buffer doesn’t know the state of its commands on the GPU.
	MTLCommandEncoderErrorStateUnknown MTLCommandEncoderErrorState = 0
)

func (e MTLCommandEncoderErrorState) String() string {
	switch e {
	case MTLCommandEncoderErrorStateAffected:
		return "MTLCommandEncoderErrorStateAffected"
	case MTLCommandEncoderErrorStateCompleted:
		return "MTLCommandEncoderErrorStateCompleted"
	case MTLCommandEncoderErrorStateFaulted:
		return "MTLCommandEncoderErrorStateFaulted"
	case MTLCommandEncoderErrorStatePending:
		return "MTLCommandEncoderErrorStatePending"
	case MTLCommandEncoderErrorStateUnknown:
		return "MTLCommandEncoderErrorStateUnknown"
	default:
		return fmt.Sprintf("MTLCommandEncoderErrorState(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLCompareFunction
type MTLCompareFunction uint

const (
	// MTLCompareFunctionAlways: A new value always passes the comparison test.
	MTLCompareFunctionAlways MTLCompareFunction = 7
	// MTLCompareFunctionEqual: A new value passes the comparison test if it is equal to the existing value.
	MTLCompareFunctionEqual MTLCompareFunction = 2
	// MTLCompareFunctionGreater: A new value passes the comparison test if it is greater than the existing value.
	MTLCompareFunctionGreater MTLCompareFunction = 4
	// MTLCompareFunctionGreaterEqual: A new value passes the comparison test if it is greater than or equal to the existing value.
	MTLCompareFunctionGreaterEqual MTLCompareFunction = 6
	// MTLCompareFunctionLess: A new value passes the comparison test if it is less than the existing value.
	MTLCompareFunctionLess MTLCompareFunction = 1
	// MTLCompareFunctionLessEqual: A new value passes the comparison test if it is less than or equal to the existing value.
	MTLCompareFunctionLessEqual MTLCompareFunction = 3
	// MTLCompareFunctionNever: A new value never passes the comparison test.
	MTLCompareFunctionNever MTLCompareFunction = 0
	// MTLCompareFunctionNotEqual: A new value passes the comparison test if it is not equal to the existing value.
	MTLCompareFunctionNotEqual MTLCompareFunction = 5
)

func (e MTLCompareFunction) String() string {
	switch e {
	case MTLCompareFunctionAlways:
		return "MTLCompareFunctionAlways"
	case MTLCompareFunctionEqual:
		return "MTLCompareFunctionEqual"
	case MTLCompareFunctionGreater:
		return "MTLCompareFunctionGreater"
	case MTLCompareFunctionGreaterEqual:
		return "MTLCompareFunctionGreaterEqual"
	case MTLCompareFunctionLess:
		return "MTLCompareFunctionLess"
	case MTLCompareFunctionLessEqual:
		return "MTLCompareFunctionLessEqual"
	case MTLCompareFunctionNever:
		return "MTLCompareFunctionNever"
	case MTLCompareFunctionNotEqual:
		return "MTLCompareFunctionNotEqual"
	default:
		return fmt.Sprintf("MTLCompareFunction(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLCompileSymbolVisibility
type MTLCompileSymbolVisibility int

const (
	MTLCompileSymbolVisibilityDefault MTLCompileSymbolVisibility = 0
	MTLCompileSymbolVisibilityHidden  MTLCompileSymbolVisibility = 1
)

func (e MTLCompileSymbolVisibility) String() string {
	switch e {
	case MTLCompileSymbolVisibilityDefault:
		return "MTLCompileSymbolVisibilityDefault"
	case MTLCompileSymbolVisibilityHidden:
		return "MTLCompileSymbolVisibilityHidden"
	default:
		return fmt.Sprintf("MTLCompileSymbolVisibility(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLCounterSampleBufferError-swift.struct/Code
type MTLCounterSampleBufferError int

const (
	// MTLCounterSampleBufferErrorInternal: An error code that indicates the Metal framework has an internal problem.
	MTLCounterSampleBufferErrorInternal MTLCounterSampleBufferError = 2
	// MTLCounterSampleBufferErrorInvalid: An error code that indicates when a counter-sample buffer descriptor has at least one invalid property.
	MTLCounterSampleBufferErrorInvalid MTLCounterSampleBufferError = 1
	// MTLCounterSampleBufferErrorOutOfMemory: An error code that indicates the GPU device doesn’t have sufficient memory to create a counter sample buffer.
	MTLCounterSampleBufferErrorOutOfMemory MTLCounterSampleBufferError = 0
)

func (e MTLCounterSampleBufferError) String() string {
	switch e {
	case MTLCounterSampleBufferErrorInternal:
		return "MTLCounterSampleBufferErrorInternal"
	case MTLCounterSampleBufferErrorInvalid:
		return "MTLCounterSampleBufferErrorInvalid"
	case MTLCounterSampleBufferErrorOutOfMemory:
		return "MTLCounterSampleBufferErrorOutOfMemory"
	default:
		return fmt.Sprintf("MTLCounterSampleBufferError(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLCounterSamplingPoint
type MTLCounterSamplingPoint uint

const (
	// MTLCounterSamplingPointAtBlitBoundary: Counter sampling is allowed between blit commands in a blit pass.
	MTLCounterSamplingPointAtBlitBoundary MTLCounterSamplingPoint = 4
	// MTLCounterSamplingPointAtDispatchBoundary: Counter sampling is allowed between kernel dispatches in a compute pass.
	MTLCounterSamplingPointAtDispatchBoundary MTLCounterSamplingPoint = 2
	// MTLCounterSamplingPointAtDrawBoundary: Counter sampling is allowed between draw commands in a render pass.
	MTLCounterSamplingPointAtDrawBoundary MTLCounterSamplingPoint = 1
	// MTLCounterSamplingPointAtStageBoundary: Counter sampling is allowed at the start and end of a render pass’s vertex and fragment stages, and at the start and end of compute and blit passes.
	MTLCounterSamplingPointAtStageBoundary MTLCounterSamplingPoint = 0
	// MTLCounterSamplingPointAtTileDispatchBoundary: Counter sampling is allowed between tile dispatches in a render pass.
	MTLCounterSamplingPointAtTileDispatchBoundary MTLCounterSamplingPoint = 3
)

func (e MTLCounterSamplingPoint) String() string {
	switch e {
	case MTLCounterSamplingPointAtBlitBoundary:
		return "MTLCounterSamplingPointAtBlitBoundary"
	case MTLCounterSamplingPointAtDispatchBoundary:
		return "MTLCounterSamplingPointAtDispatchBoundary"
	case MTLCounterSamplingPointAtDrawBoundary:
		return "MTLCounterSamplingPointAtDrawBoundary"
	case MTLCounterSamplingPointAtStageBoundary:
		return "MTLCounterSamplingPointAtStageBoundary"
	case MTLCounterSamplingPointAtTileDispatchBoundary:
		return "MTLCounterSamplingPointAtTileDispatchBoundary"
	default:
		return fmt.Sprintf("MTLCounterSamplingPoint(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLCullMode
type MTLCullMode uint

const (
	// MTLCullModeBack: Culls back-facing primitives.
	MTLCullModeBack MTLCullMode = 2
	// MTLCullModeFront: Culls front-facing primitives.
	MTLCullModeFront MTLCullMode = 1
	// MTLCullModeNone: Does not cull any primitives.
	MTLCullModeNone MTLCullMode = 0
)

func (e MTLCullMode) String() string {
	switch e {
	case MTLCullModeBack:
		return "MTLCullModeBack"
	case MTLCullModeFront:
		return "MTLCullModeFront"
	case MTLCullModeNone:
		return "MTLCullModeNone"
	default:
		return fmt.Sprintf("MTLCullMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLCurveBasis
type MTLCurveBasis int

const (
	MTLCurveBasisBSpline    MTLCurveBasis = 0
	MTLCurveBasisBezier     MTLCurveBasis = 3
	MTLCurveBasisCatmullRom MTLCurveBasis = 1
	MTLCurveBasisLinear     MTLCurveBasis = 2
)

func (e MTLCurveBasis) String() string {
	switch e {
	case MTLCurveBasisBSpline:
		return "MTLCurveBasisBSpline"
	case MTLCurveBasisBezier:
		return "MTLCurveBasisBezier"
	case MTLCurveBasisCatmullRom:
		return "MTLCurveBasisCatmullRom"
	case MTLCurveBasisLinear:
		return "MTLCurveBasisLinear"
	default:
		return fmt.Sprintf("MTLCurveBasis(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLCurveEndCaps
type MTLCurveEndCaps int

const (
	MTLCurveEndCapsDisk   MTLCurveEndCaps = 1
	MTLCurveEndCapsNone   MTLCurveEndCaps = 0
	MTLCurveEndCapsSphere MTLCurveEndCaps = 2
)

func (e MTLCurveEndCaps) String() string {
	switch e {
	case MTLCurveEndCapsDisk:
		return "MTLCurveEndCapsDisk"
	case MTLCurveEndCapsNone:
		return "MTLCurveEndCapsNone"
	case MTLCurveEndCapsSphere:
		return "MTLCurveEndCapsSphere"
	default:
		return fmt.Sprintf("MTLCurveEndCaps(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLCurveType
type MTLCurveType int

const (
	MTLCurveTypeFlat  MTLCurveType = 1
	MTLCurveTypeRound MTLCurveType = 0
)

func (e MTLCurveType) String() string {
	switch e {
	case MTLCurveTypeFlat:
		return "MTLCurveTypeFlat"
	case MTLCurveTypeRound:
		return "MTLCurveTypeRound"
	default:
		return fmt.Sprintf("MTLCurveType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLDataType
type MTLDataType uint

const (
	// MTLDataTypeArray: An array instance.
	MTLDataTypeArray MTLDataType = 2
	// MTLDataTypeBFloat: A 16-bit, brain floating-point value.
	MTLDataTypeBFloat MTLDataType = 121
	// MTLDataTypeBFloat2: A two-component vector with 16-bit, brain floating-point values.
	MTLDataTypeBFloat2 MTLDataType = 122
	// MTLDataTypeBFloat3: A three-component vector with 16-bit, brain floating-point values.
	MTLDataTypeBFloat3 MTLDataType = 123
	// MTLDataTypeBFloat4: A four-component vector with 16-bit, brain floating-point values.
	MTLDataTypeBFloat4 MTLDataType = 124
	// MTLDataTypeBool: A Boolean value.
	MTLDataTypeBool MTLDataType = 53
	// MTLDataTypeBool2: A two-component Boolean vector.
	MTLDataTypeBool2 MTLDataType = 54
	// MTLDataTypeBool3: A three-component Boolean vector.
	MTLDataTypeBool3 MTLDataType = 55
	// MTLDataTypeBool4: A four-component Boolean vector.
	MTLDataTypeBool4 MTLDataType = 56
	// MTLDataTypeChar: An 8-bit, signed integer value.
	MTLDataTypeChar MTLDataType = 45
	// MTLDataTypeChar2: A two-component vector with 8-bit, signed integer values.
	MTLDataTypeChar2 MTLDataType = 46
	// MTLDataTypeChar3: A three-component vector with 8-bit, signed integer values.
	MTLDataTypeChar3 MTLDataType = 47
	// MTLDataTypeChar4: A four-component vector with 8-bit, signed integer values.
	MTLDataTypeChar4 MTLDataType = 48
	// MTLDataTypeComputePipeline: A Metal compute pipeline instance.
	MTLDataTypeComputePipeline MTLDataType = 79
	// MTLDataTypeDepthStencilState: Represents a data type corresponding to a depth-stencil state object.
	MTLDataTypeDepthStencilState MTLDataType = 139
	// MTLDataTypeFloat: A 32-bit floating-point value.
	MTLDataTypeFloat MTLDataType = 3
	// MTLDataTypeFloat2: A two-component vector with 32-bit floating-point values.
	MTLDataTypeFloat2 MTLDataType = 4
	// MTLDataTypeFloat2x2: A 2x2 component matrix with 32-bit floating-point values.
	MTLDataTypeFloat2x2 MTLDataType = 7
	// MTLDataTypeFloat2x3: A 2x3 component matrix with 32-bit floating-point values.
	MTLDataTypeFloat2x3 MTLDataType = 8
	// MTLDataTypeFloat2x4: A 2x4 component matrix with 32-bit floating-point values.
	MTLDataTypeFloat2x4 MTLDataType = 9
	// MTLDataTypeFloat3: A three-component vector with 32-bit floating-point values.
	MTLDataTypeFloat3 MTLDataType = 5
	// MTLDataTypeFloat3x2: A 3x2 component matrix with 32-bit floating-point values.
	MTLDataTypeFloat3x2 MTLDataType = 10
	// MTLDataTypeFloat3x3: A 3x3 component matrix with 32-bit floating-point values.
	MTLDataTypeFloat3x3 MTLDataType = 11
	// MTLDataTypeFloat3x4: A 3x4 component matrix with 32-bit floating-point values.
	MTLDataTypeFloat3x4 MTLDataType = 12
	// MTLDataTypeFloat4: A four-component vector with 32-bit floating-point values.
	MTLDataTypeFloat4 MTLDataType = 6
	// MTLDataTypeFloat4x2: A 4x2 component matrix with 32-bit floating-point values.
	MTLDataTypeFloat4x2 MTLDataType = 13
	// MTLDataTypeFloat4x3: A 4x3 component matrix with 32-bit floating-point values.
	MTLDataTypeFloat4x3 MTLDataType = 14
	// MTLDataTypeFloat4x4: A 4x4 component matrix with 32-bit floating-point values.
	MTLDataTypeFloat4x4 MTLDataType = 15
	// MTLDataTypeHalf: A 16-bit floating-point value.
	MTLDataTypeHalf MTLDataType = 16
	// MTLDataTypeHalf2: A two-component vector with 16-bit floating-point values.
	MTLDataTypeHalf2 MTLDataType = 17
	// MTLDataTypeHalf2x2: A 2x2 component matrix with 16-bit floating-point values.
	MTLDataTypeHalf2x2 MTLDataType = 20
	// MTLDataTypeHalf2x3: A 2x3 component matrix with 16-bit floating-point values.
	MTLDataTypeHalf2x3 MTLDataType = 21
	// MTLDataTypeHalf2x4: A 2x4 component matrix with 16-bit floating-point values.
	MTLDataTypeHalf2x4 MTLDataType = 22
	// MTLDataTypeHalf3: A three-component vector with 16-bit floating-point values.
	MTLDataTypeHalf3 MTLDataType = 18
	// MTLDataTypeHalf3x2: A 3x2 component matrix with 16-bit floating-point values.
	MTLDataTypeHalf3x2 MTLDataType = 23
	// MTLDataTypeHalf3x3: A 3x3 component matrix with 16-bit floating-point values.
	MTLDataTypeHalf3x3 MTLDataType = 24
	// MTLDataTypeHalf3x4: A 3x4 component matrix with 16-bit floating-point values.
	MTLDataTypeHalf3x4 MTLDataType = 25
	// MTLDataTypeHalf4: A four-component vector with 16-bit floating-point values.
	MTLDataTypeHalf4 MTLDataType = 19
	// MTLDataTypeHalf4x2: A 4x2 component matrix with 16-bit floating-point values.
	MTLDataTypeHalf4x2 MTLDataType = 26
	// MTLDataTypeHalf4x3: A 4x3 component matrix with 16-bit floating-point values.
	MTLDataTypeHalf4x3 MTLDataType = 27
	// MTLDataTypeHalf4x4: A 4x4 component matrix with 16-bit floating-point values.
	MTLDataTypeHalf4x4 MTLDataType = 28
	// MTLDataTypeIndirectCommandBuffer: An indirect command buffer resource instance.
	MTLDataTypeIndirectCommandBuffer MTLDataType = 80
	// MTLDataTypeInstanceAccelerationStructure: A high-level, ray-tracing acceleration structure for a set of low-level primitive instances.
	MTLDataTypeInstanceAccelerationStructure MTLDataType = 118
	// MTLDataTypeInt: A 32-bit, signed integer value.
	MTLDataTypeInt MTLDataType = 29
	// MTLDataTypeInt2: A two-component vector with 32-bit, signed integer values.
	MTLDataTypeInt2 MTLDataType = 30
	// MTLDataTypeInt3: A three-component vector with 32-bit, signed integer values.
	MTLDataTypeInt3 MTLDataType = 31
	// MTLDataTypeInt4: A four-component vector with 32-bit, signed integer values.
	MTLDataTypeInt4 MTLDataType = 32
	// MTLDataTypeIntersectionFunctionTable: A table of intersection functions that a render or compute pipeline can call.
	MTLDataTypeIntersectionFunctionTable MTLDataType = 116
	// MTLDataTypeLong: A 64-bit, signed integer value.
	MTLDataTypeLong MTLDataType = 81
	// MTLDataTypeLong2: A two-component vector with 64-bit, signed integer values.
	MTLDataTypeLong2 MTLDataType = 82
	// MTLDataTypeLong3: A three-component vector with 64-bit, signed integer values.
	MTLDataTypeLong3 MTLDataType = 83
	// MTLDataTypeLong4: A four-component vector with 64-bit, signed integer values.
	MTLDataTypeLong4 MTLDataType = 84
	// MTLDataTypeNone: A sentinel value that represents a GPU function parameter that doesn’t have a valid data type.
	MTLDataTypeNone MTLDataType = 0
	// MTLDataTypePointer: A pointer.
	MTLDataTypePointer MTLDataType = 60
	// MTLDataTypePrimitiveAccelerationStructure: A low-level ray-tracing acceleration structure for a set of primitives.
	MTLDataTypePrimitiveAccelerationStructure MTLDataType = 117
	// MTLDataTypeR16Snorm: An ordinary pixel with one component that’s a 16-bit, normalized, signed integer value.
	MTLDataTypeR16Snorm MTLDataType = 65
	// MTLDataTypeR16Unorm: An ordinary pixel with one component that’s a 16-bit, normalized, unsigned integer value.
	MTLDataTypeR16Unorm MTLDataType = 64
	// MTLDataTypeR8Snorm: An ordinary pixel with one component that’s an 8-bit, normalized, signed integer value.
	MTLDataTypeR8Snorm MTLDataType = 63
	// MTLDataTypeR8Unorm: An ordinary pixel with one component that’s an 8-bit, normalized, unsigned integer value.
	MTLDataTypeR8Unorm MTLDataType = 62
	// MTLDataTypeRG11B10Float: A packed 32-bit format with three floating-point color components, two of which are 11-bit values, and one is a 10-bit value.
	MTLDataTypeRG11B10Float MTLDataType = 76
	// MTLDataTypeRG16Snorm: An ordinary pixel with two components, each of which is a 16-bit, normalized, signed integer value.
	MTLDataTypeRG16Snorm MTLDataType = 69
	// MTLDataTypeRG16Unorm: An ordinary pixel with two components, each of which is a 16-bit, normalized, unsigned integer value.
	MTLDataTypeRG16Unorm MTLDataType = 68
	// MTLDataTypeRG8Snorm: An ordinary pixel with two components, each of which is an 8-bit, normalized, signed integer value.
	MTLDataTypeRG8Snorm MTLDataType = 67
	// MTLDataTypeRG8Unorm: An ordinary pixel with two components, each of which is an 8-bit, normalized, unsigned integer value.
	MTLDataTypeRG8Unorm MTLDataType = 66
	// MTLDataTypeRGB10A2Unorm: A packed 32-bit format with three color components, each of which is a 10-bit, normalized, unsigned integer value.
	MTLDataTypeRGB10A2Unorm MTLDataType = 75
	// MTLDataTypeRGB9E5Float: A packed 32-bit format with three color components, each of which is a 9-bit floating-point value.
	MTLDataTypeRGB9E5Float MTLDataType = 77
	// MTLDataTypeRGBA16Snorm: An ordinary pixel with four components, each of which is a 16-bit, normalized, signed integer value.
	MTLDataTypeRGBA16Snorm MTLDataType = 74
	// MTLDataTypeRGBA16Unorm: An ordinary pixel with four components, each of which is a 16-bit, normalized, unsigned integer value.
	MTLDataTypeRGBA16Unorm MTLDataType = 73
	// MTLDataTypeRGBA8Snorm: An ordinary pixel with four components, each of which is an 8-bit, normalized, signed integer value.
	MTLDataTypeRGBA8Snorm MTLDataType = 72
	// MTLDataTypeRGBA8Unorm: An ordinary pixel with four components, each of which is an 8-bit, normalized, unsigned integer value.
	MTLDataTypeRGBA8Unorm MTLDataType = 70
	// MTLDataTypeRGBA8Unorm_sRGB: An ordinary pixel with four components, each of which is an 8-bit, normalized, unsigned integer value in the sRGB color space.
	MTLDataTypeRGBA8Unorm_sRGB MTLDataType = 71
	// MTLDataTypeRenderPipeline: A Metal render pipeline instance.
	MTLDataTypeRenderPipeline MTLDataType = 78
	// MTLDataTypeSampler: A Metal texture sampler instance.
	MTLDataTypeSampler MTLDataType = 59
	// MTLDataTypeShort: A 16-bit, signed integer value.
	MTLDataTypeShort MTLDataType = 37
	// MTLDataTypeShort2: A two-component vector with 16-bit, signed integer values.
	MTLDataTypeShort2 MTLDataType = 38
	// MTLDataTypeShort3: A three-component vector with 16-bit, signed integer values.
	MTLDataTypeShort3 MTLDataType = 39
	// MTLDataTypeShort4: A four-component vector with 16-bit, signed integer values.
	MTLDataTypeShort4 MTLDataType = 40
	// MTLDataTypeStruct: A structure instance.
	MTLDataTypeStruct MTLDataType = 1
	// MTLDataTypeTensor: Represents a data type corresponding to a machine learning tensor.
	MTLDataTypeTensor MTLDataType = 140
	// MTLDataTypeTexture: A Metal texture resource instance.
	MTLDataTypeTexture MTLDataType = 58
	// MTLDataTypeUChar: An 8-bit, unsigned integer value.
	MTLDataTypeUChar MTLDataType = 49
	// MTLDataTypeUChar2: A two-component vector with 8-bit, unsigned integer values.
	MTLDataTypeUChar2 MTLDataType = 50
	// MTLDataTypeUChar3: A three-component vector with 8-bit, unsigned integer values.
	MTLDataTypeUChar3 MTLDataType = 51
	// MTLDataTypeUChar4: A four-component vector with 8-bit, unsigned integer values.
	MTLDataTypeUChar4 MTLDataType = 52
	// MTLDataTypeUInt: A 32-bit, unsigned integer value.
	MTLDataTypeUInt MTLDataType = 33
	// MTLDataTypeUInt2: A two-component vector with 32-bit, unsigned integer values.
	MTLDataTypeUInt2 MTLDataType = 34
	// MTLDataTypeUInt3: A three-component vector with 32-bit, unsigned integer values.
	MTLDataTypeUInt3 MTLDataType = 35
	// MTLDataTypeUInt4: A four-component vector with 32-bit, unsigned integer values.
	MTLDataTypeUInt4 MTLDataType = 36
	// MTLDataTypeULong: A 64-bit, unsigned integer value.
	MTLDataTypeULong MTLDataType = 85
	// MTLDataTypeULong2: A two-component vector with 64-bit, unsigned integer values.
	MTLDataTypeULong2 MTLDataType = 86
	// MTLDataTypeULong3: A three-component vector with 64-bit, unsigned integer values.
	MTLDataTypeULong3 MTLDataType = 87
	// MTLDataTypeULong4: A four-component vector with 64-bit, unsigned integer values.
	MTLDataTypeULong4 MTLDataType = 88
	// MTLDataTypeUShort: A 16-bit, unsigned integer value.
	MTLDataTypeUShort MTLDataType = 41
	// MTLDataTypeUShort2: A two-component vector with 16-bit, unsigned integer values.
	MTLDataTypeUShort2 MTLDataType = 42
	// MTLDataTypeUShort3: A three-component vector with 16-bit, unsigned integer values.
	MTLDataTypeUShort3 MTLDataType = 43
	// MTLDataTypeUShort4: A four-component vector with 16-bit, unsigned integer values.
	MTLDataTypeUShort4 MTLDataType = 44
	// MTLDataTypeVisibleFunctionTable: A table of visible functions that a render or compute pipeline can call.
	MTLDataTypeVisibleFunctionTable MTLDataType = 115
)

func (e MTLDataType) String() string {
	switch e {
	case MTLDataTypeArray:
		return "MTLDataTypeArray"
	case MTLDataTypeBFloat:
		return "MTLDataTypeBFloat"
	case MTLDataTypeBFloat2:
		return "MTLDataTypeBFloat2"
	case MTLDataTypeBFloat3:
		return "MTLDataTypeBFloat3"
	case MTLDataTypeBFloat4:
		return "MTLDataTypeBFloat4"
	case MTLDataTypeBool:
		return "MTLDataTypeBool"
	case MTLDataTypeBool2:
		return "MTLDataTypeBool2"
	case MTLDataTypeBool3:
		return "MTLDataTypeBool3"
	case MTLDataTypeBool4:
		return "MTLDataTypeBool4"
	case MTLDataTypeChar:
		return "MTLDataTypeChar"
	case MTLDataTypeChar2:
		return "MTLDataTypeChar2"
	case MTLDataTypeChar3:
		return "MTLDataTypeChar3"
	case MTLDataTypeChar4:
		return "MTLDataTypeChar4"
	case MTLDataTypeComputePipeline:
		return "MTLDataTypeComputePipeline"
	case MTLDataTypeDepthStencilState:
		return "MTLDataTypeDepthStencilState"
	case MTLDataTypeFloat:
		return "MTLDataTypeFloat"
	case MTLDataTypeFloat2:
		return "MTLDataTypeFloat2"
	case MTLDataTypeFloat2x2:
		return "MTLDataTypeFloat2x2"
	case MTLDataTypeFloat2x3:
		return "MTLDataTypeFloat2x3"
	case MTLDataTypeFloat2x4:
		return "MTLDataTypeFloat2x4"
	case MTLDataTypeFloat3:
		return "MTLDataTypeFloat3"
	case MTLDataTypeFloat3x2:
		return "MTLDataTypeFloat3x2"
	case MTLDataTypeFloat3x3:
		return "MTLDataTypeFloat3x3"
	case MTLDataTypeFloat3x4:
		return "MTLDataTypeFloat3x4"
	case MTLDataTypeFloat4:
		return "MTLDataTypeFloat4"
	case MTLDataTypeFloat4x2:
		return "MTLDataTypeFloat4x2"
	case MTLDataTypeFloat4x3:
		return "MTLDataTypeFloat4x3"
	case MTLDataTypeFloat4x4:
		return "MTLDataTypeFloat4x4"
	case MTLDataTypeHalf:
		return "MTLDataTypeHalf"
	case MTLDataTypeHalf2:
		return "MTLDataTypeHalf2"
	case MTLDataTypeHalf2x2:
		return "MTLDataTypeHalf2x2"
	case MTLDataTypeHalf2x3:
		return "MTLDataTypeHalf2x3"
	case MTLDataTypeHalf2x4:
		return "MTLDataTypeHalf2x4"
	case MTLDataTypeHalf3:
		return "MTLDataTypeHalf3"
	case MTLDataTypeHalf3x2:
		return "MTLDataTypeHalf3x2"
	case MTLDataTypeHalf3x3:
		return "MTLDataTypeHalf3x3"
	case MTLDataTypeHalf3x4:
		return "MTLDataTypeHalf3x4"
	case MTLDataTypeHalf4:
		return "MTLDataTypeHalf4"
	case MTLDataTypeHalf4x2:
		return "MTLDataTypeHalf4x2"
	case MTLDataTypeHalf4x3:
		return "MTLDataTypeHalf4x3"
	case MTLDataTypeHalf4x4:
		return "MTLDataTypeHalf4x4"
	case MTLDataTypeIndirectCommandBuffer:
		return "MTLDataTypeIndirectCommandBuffer"
	case MTLDataTypeInstanceAccelerationStructure:
		return "MTLDataTypeInstanceAccelerationStructure"
	case MTLDataTypeInt:
		return "MTLDataTypeInt"
	case MTLDataTypeInt2:
		return "MTLDataTypeInt2"
	case MTLDataTypeInt3:
		return "MTLDataTypeInt3"
	case MTLDataTypeInt4:
		return "MTLDataTypeInt4"
	case MTLDataTypeIntersectionFunctionTable:
		return "MTLDataTypeIntersectionFunctionTable"
	case MTLDataTypeLong:
		return "MTLDataTypeLong"
	case MTLDataTypeLong2:
		return "MTLDataTypeLong2"
	case MTLDataTypeLong3:
		return "MTLDataTypeLong3"
	case MTLDataTypeLong4:
		return "MTLDataTypeLong4"
	case MTLDataTypeNone:
		return "MTLDataTypeNone"
	case MTLDataTypePointer:
		return "MTLDataTypePointer"
	case MTLDataTypePrimitiveAccelerationStructure:
		return "MTLDataTypePrimitiveAccelerationStructure"
	case MTLDataTypeR16Snorm:
		return "MTLDataTypeR16Snorm"
	case MTLDataTypeR16Unorm:
		return "MTLDataTypeR16Unorm"
	case MTLDataTypeR8Snorm:
		return "MTLDataTypeR8Snorm"
	case MTLDataTypeR8Unorm:
		return "MTLDataTypeR8Unorm"
	case MTLDataTypeRG11B10Float:
		return "MTLDataTypeRG11B10Float"
	case MTLDataTypeRG16Snorm:
		return "MTLDataTypeRG16Snorm"
	case MTLDataTypeRG16Unorm:
		return "MTLDataTypeRG16Unorm"
	case MTLDataTypeRG8Snorm:
		return "MTLDataTypeRG8Snorm"
	case MTLDataTypeRG8Unorm:
		return "MTLDataTypeRG8Unorm"
	case MTLDataTypeRGB10A2Unorm:
		return "MTLDataTypeRGB10A2Unorm"
	case MTLDataTypeRGB9E5Float:
		return "MTLDataTypeRGB9E5Float"
	case MTLDataTypeRGBA16Snorm:
		return "MTLDataTypeRGBA16Snorm"
	case MTLDataTypeRGBA16Unorm:
		return "MTLDataTypeRGBA16Unorm"
	case MTLDataTypeRGBA8Snorm:
		return "MTLDataTypeRGBA8Snorm"
	case MTLDataTypeRGBA8Unorm:
		return "MTLDataTypeRGBA8Unorm"
	case MTLDataTypeRGBA8Unorm_sRGB:
		return "MTLDataTypeRGBA8Unorm_sRGB"
	case MTLDataTypeRenderPipeline:
		return "MTLDataTypeRenderPipeline"
	case MTLDataTypeSampler:
		return "MTLDataTypeSampler"
	case MTLDataTypeShort:
		return "MTLDataTypeShort"
	case MTLDataTypeShort2:
		return "MTLDataTypeShort2"
	case MTLDataTypeShort3:
		return "MTLDataTypeShort3"
	case MTLDataTypeShort4:
		return "MTLDataTypeShort4"
	case MTLDataTypeStruct:
		return "MTLDataTypeStruct"
	case MTLDataTypeTensor:
		return "MTLDataTypeTensor"
	case MTLDataTypeTexture:
		return "MTLDataTypeTexture"
	case MTLDataTypeUChar:
		return "MTLDataTypeUChar"
	case MTLDataTypeUChar2:
		return "MTLDataTypeUChar2"
	case MTLDataTypeUChar3:
		return "MTLDataTypeUChar3"
	case MTLDataTypeUChar4:
		return "MTLDataTypeUChar4"
	case MTLDataTypeUInt:
		return "MTLDataTypeUInt"
	case MTLDataTypeUInt2:
		return "MTLDataTypeUInt2"
	case MTLDataTypeUInt3:
		return "MTLDataTypeUInt3"
	case MTLDataTypeUInt4:
		return "MTLDataTypeUInt4"
	case MTLDataTypeULong:
		return "MTLDataTypeULong"
	case MTLDataTypeULong2:
		return "MTLDataTypeULong2"
	case MTLDataTypeULong3:
		return "MTLDataTypeULong3"
	case MTLDataTypeULong4:
		return "MTLDataTypeULong4"
	case MTLDataTypeUShort:
		return "MTLDataTypeUShort"
	case MTLDataTypeUShort2:
		return "MTLDataTypeUShort2"
	case MTLDataTypeUShort3:
		return "MTLDataTypeUShort3"
	case MTLDataTypeUShort4:
		return "MTLDataTypeUShort4"
	case MTLDataTypeVisibleFunctionTable:
		return "MTLDataTypeVisibleFunctionTable"
	default:
		return fmt.Sprintf("MTLDataType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLDepthClipMode
type MTLDepthClipMode uint

const (
	// MTLDepthClipModeClamp: Clamp fragments outside the near or far planes.
	MTLDepthClipModeClamp MTLDepthClipMode = 1
	// MTLDepthClipModeClip: Clip fragments outside the near or far planes.
	MTLDepthClipModeClip MTLDepthClipMode = 0
)

func (e MTLDepthClipMode) String() string {
	switch e {
	case MTLDepthClipModeClamp:
		return "MTLDepthClipModeClamp"
	case MTLDepthClipModeClip:
		return "MTLDepthClipModeClip"
	default:
		return fmt.Sprintf("MTLDepthClipMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLDeviceLocation
type MTLDeviceLocation uint

const (
	// MTLDeviceLocationBuiltIn: A location that indicates the GPU is permanently connected to the system internally.
	MTLDeviceLocationBuiltIn MTLDeviceLocation = 0
	// MTLDeviceLocationExternal: A GPU location that indicates a person connected the GPU to the system with an external interface, such as Thunderbolt.
	MTLDeviceLocationExternal MTLDeviceLocation = 2
	// MTLDeviceLocationSlot: A GPU location that indicates a person connected the GPU to a system’s internal slot.
	MTLDeviceLocationSlot MTLDeviceLocation = 1
	// MTLDeviceLocationUnspecified: A value that indicates the system can’t determine how the GPU connects to it.
	MTLDeviceLocationUnspecified MTLDeviceLocation = 0
)

func (e MTLDeviceLocation) String() string {
	switch e {
	case MTLDeviceLocationBuiltIn:
		return "MTLDeviceLocationBuiltIn"
	case MTLDeviceLocationExternal:
		return "MTLDeviceLocationExternal"
	case MTLDeviceLocationSlot:
		return "MTLDeviceLocationSlot"
	default:
		return fmt.Sprintf("MTLDeviceLocation(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLDispatchType
type MTLDispatchType uint

const (
	// MTLDispatchTypeConcurrent: Sets a command encoder to dispatch encoded commands concurrently during your pass.
	MTLDispatchTypeConcurrent MTLDispatchType = 1
	// MTLDispatchTypeSerial: Sets a command encoder to dispatch encoded commands serially during your pass.
	MTLDispatchTypeSerial MTLDispatchType = 0
)

func (e MTLDispatchType) String() string {
	switch e {
	case MTLDispatchTypeConcurrent:
		return "MTLDispatchTypeConcurrent"
	case MTLDispatchTypeSerial:
		return "MTLDispatchTypeSerial"
	default:
		return fmt.Sprintf("MTLDispatchType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLDynamicLibraryError-swift.struct/Code
type MTLDynamicLibraryError int

const (
	// MTLDynamicLibraryErrorCompilationFailure: An error code that indicates Metal couldn’t compile a dynamic library.
	MTLDynamicLibraryErrorCompilationFailure MTLDynamicLibraryError = 2
	// MTLDynamicLibraryErrorDependencyLoadFailure: An error code that indicates a dynamic library couldn’t link to other dynamic libraries.
	MTLDynamicLibraryErrorDependencyLoadFailure MTLDynamicLibraryError = 4
	// MTLDynamicLibraryErrorInvalidFile: An error code that indicates an app is using an invalid reference to a library file, typically related to a URL.
	MTLDynamicLibraryErrorInvalidFile MTLDynamicLibraryError = 1
	// MTLDynamicLibraryErrorNone: An error code that represents the absence of any problems.
	MTLDynamicLibraryErrorNone MTLDynamicLibraryError = 0
	// MTLDynamicLibraryErrorUnresolvedInstallName: An error code that indicates Metal couldn’t resolve the installation name for a new dynamic library.
	MTLDynamicLibraryErrorUnresolvedInstallName MTLDynamicLibraryError = 3
	// MTLDynamicLibraryErrorUnsupported: An error code that indicates the GPU device doesn’t support dynamic libraries.
	MTLDynamicLibraryErrorUnsupported MTLDynamicLibraryError = 5
)

func (e MTLDynamicLibraryError) String() string {
	switch e {
	case MTLDynamicLibraryErrorCompilationFailure:
		return "MTLDynamicLibraryErrorCompilationFailure"
	case MTLDynamicLibraryErrorDependencyLoadFailure:
		return "MTLDynamicLibraryErrorDependencyLoadFailure"
	case MTLDynamicLibraryErrorInvalidFile:
		return "MTLDynamicLibraryErrorInvalidFile"
	case MTLDynamicLibraryErrorNone:
		return "MTLDynamicLibraryErrorNone"
	case MTLDynamicLibraryErrorUnresolvedInstallName:
		return "MTLDynamicLibraryErrorUnresolvedInstallName"
	case MTLDynamicLibraryErrorUnsupported:
		return "MTLDynamicLibraryErrorUnsupported"
	default:
		return fmt.Sprintf("MTLDynamicLibraryError(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLFeatureSet
type MTLFeatureSet uint

const (
	// MTLFeatureSet_iOS_GPUFamily1_v1: The GPU family 1, version 1 feature set for iOS.
	MTLFeatureSet_iOS_GPUFamily1_v1 MTLFeatureSet = 0
	// MTLFeatureSet_iOS_GPUFamily1_v2: The GPU family 1, version 2 feature set for iOS.
	MTLFeatureSet_iOS_GPUFamily1_v2 MTLFeatureSet = 2
	// MTLFeatureSet_iOS_GPUFamily1_v3: The GPU family 1, version 3 feature set for iOS.
	MTLFeatureSet_iOS_GPUFamily1_v3 MTLFeatureSet = 5
	// MTLFeatureSet_iOS_GPUFamily1_v4: The GPU family 1, version 4 feature set for iOS.
	MTLFeatureSet_iOS_GPUFamily1_v4 MTLFeatureSet = 8
	// MTLFeatureSet_iOS_GPUFamily1_v5: The GPU family 1, version 5 feature set for iOS.
	MTLFeatureSet_iOS_GPUFamily1_v5 MTLFeatureSet = 12
	// MTLFeatureSet_iOS_GPUFamily2_v1: The GPU family 2, version 1 feature set for iOS.
	MTLFeatureSet_iOS_GPUFamily2_v1 MTLFeatureSet = 1
	// MTLFeatureSet_iOS_GPUFamily2_v2: The GPU family 2, version 2 feature set for iOS.
	MTLFeatureSet_iOS_GPUFamily2_v2 MTLFeatureSet = 3
	// MTLFeatureSet_iOS_GPUFamily2_v3: The GPU family 2, version 3 feature set for iOS.
	MTLFeatureSet_iOS_GPUFamily2_v3 MTLFeatureSet = 6
	// MTLFeatureSet_iOS_GPUFamily2_v4: The GPU family 2, version 4 feature set for iOS.
	MTLFeatureSet_iOS_GPUFamily2_v4 MTLFeatureSet = 9
	// MTLFeatureSet_iOS_GPUFamily2_v5: The GPU family 2, version 5 feature set for iOS.
	MTLFeatureSet_iOS_GPUFamily2_v5 MTLFeatureSet = 13
	// MTLFeatureSet_iOS_GPUFamily3_v1: The GPU family 3, version 1 feature set for iOS.
	MTLFeatureSet_iOS_GPUFamily3_v1 MTLFeatureSet = 4
	// MTLFeatureSet_iOS_GPUFamily3_v2: The GPU family 3, version 2 feature set for iOS.
	MTLFeatureSet_iOS_GPUFamily3_v2 MTLFeatureSet = 7
	// MTLFeatureSet_iOS_GPUFamily3_v3: The GPU family 3, version 3 feature set for iOS.
	MTLFeatureSet_iOS_GPUFamily3_v3 MTLFeatureSet = 10
	// MTLFeatureSet_iOS_GPUFamily3_v4: The GPU family 3, version 4 feature set for iOS.
	MTLFeatureSet_iOS_GPUFamily3_v4 MTLFeatureSet = 14
	// MTLFeatureSet_iOS_GPUFamily4_v1: The GPU family 4, version 1 feature set for iOS.
	MTLFeatureSet_iOS_GPUFamily4_v1 MTLFeatureSet = 11
	// MTLFeatureSet_iOS_GPUFamily4_v2: The GPU family 4, version 2 feature set for iOS.
	MTLFeatureSet_iOS_GPUFamily4_v2 MTLFeatureSet = 15
	// MTLFeatureSet_iOS_GPUFamily5_v1: The GPU family 5, version 1 feature set for iOS.
	MTLFeatureSet_iOS_GPUFamily5_v1 MTLFeatureSet = 16
	// Deprecated.
	MTLFeatureSet_OSX_GPUFamily1_v1 MTLFeatureSet = 10000
	// Deprecated.
	MTLFeatureSet_OSX_GPUFamily1_v2 MTLFeatureSet = 10001
	// Deprecated.
	MTLFeatureSet_OSX_ReadWriteTextureTier2 MTLFeatureSet = 10002
	// Deprecated.
	MTLFeatureSet_TVOS_GPUFamily1_v1 MTLFeatureSet = 30000
	// Deprecated.
	MTLFeatureSet_macOS_GPUFamily1_v1 MTLFeatureSet = 10000
	// Deprecated.
	MTLFeatureSet_macOS_GPUFamily1_v2 MTLFeatureSet = 10001
	// Deprecated.
	MTLFeatureSet_macOS_GPUFamily1_v3 MTLFeatureSet = 10003
	// Deprecated.
	MTLFeatureSet_macOS_GPUFamily1_v4 MTLFeatureSet = 10004
	// Deprecated.
	MTLFeatureSet_macOS_GPUFamily2_v1 MTLFeatureSet = 10005
	// Deprecated.
	MTLFeatureSet_macOS_ReadWriteTextureTier2 MTLFeatureSet = 10002
	// Deprecated.
	MTLFeatureSet_tvOS_GPUFamily1_v1 MTLFeatureSet = 30000
	// Deprecated.
	MTLFeatureSet_tvOS_GPUFamily1_v2 MTLFeatureSet = 30001
	// Deprecated.
	MTLFeatureSet_tvOS_GPUFamily1_v3 MTLFeatureSet = 30002
	// Deprecated.
	MTLFeatureSet_tvOS_GPUFamily1_v4 MTLFeatureSet = 30004
	// Deprecated.
	MTLFeatureSet_tvOS_GPUFamily2_v1 MTLFeatureSet = 30003
	// Deprecated.
	MTLFeatureSet_tvOS_GPUFamily2_v2 MTLFeatureSet = 30005
)

func (e MTLFeatureSet) String() string {
	switch e {
	case MTLFeatureSet_iOS_GPUFamily1_v1:
		return "MTLFeatureSet_iOS_GPUFamily1_v1"
	case MTLFeatureSet_iOS_GPUFamily1_v2:
		return "MTLFeatureSet_iOS_GPUFamily1_v2"
	case MTLFeatureSet_iOS_GPUFamily1_v3:
		return "MTLFeatureSet_iOS_GPUFamily1_v3"
	case MTLFeatureSet_iOS_GPUFamily1_v4:
		return "MTLFeatureSet_iOS_GPUFamily1_v4"
	case MTLFeatureSet_iOS_GPUFamily1_v5:
		return "MTLFeatureSet_iOS_GPUFamily1_v5"
	case MTLFeatureSet_iOS_GPUFamily2_v1:
		return "MTLFeatureSet_iOS_GPUFamily2_v1"
	case MTLFeatureSet_iOS_GPUFamily2_v2:
		return "MTLFeatureSet_iOS_GPUFamily2_v2"
	case MTLFeatureSet_iOS_GPUFamily2_v3:
		return "MTLFeatureSet_iOS_GPUFamily2_v3"
	case MTLFeatureSet_iOS_GPUFamily2_v4:
		return "MTLFeatureSet_iOS_GPUFamily2_v4"
	case MTLFeatureSet_iOS_GPUFamily2_v5:
		return "MTLFeatureSet_iOS_GPUFamily2_v5"
	case MTLFeatureSet_iOS_GPUFamily3_v1:
		return "MTLFeatureSet_iOS_GPUFamily3_v1"
	case MTLFeatureSet_iOS_GPUFamily3_v2:
		return "MTLFeatureSet_iOS_GPUFamily3_v2"
	case MTLFeatureSet_iOS_GPUFamily3_v3:
		return "MTLFeatureSet_iOS_GPUFamily3_v3"
	case MTLFeatureSet_iOS_GPUFamily3_v4:
		return "MTLFeatureSet_iOS_GPUFamily3_v4"
	case MTLFeatureSet_iOS_GPUFamily4_v1:
		return "MTLFeatureSet_iOS_GPUFamily4_v1"
	case MTLFeatureSet_iOS_GPUFamily4_v2:
		return "MTLFeatureSet_iOS_GPUFamily4_v2"
	case MTLFeatureSet_iOS_GPUFamily5_v1:
		return "MTLFeatureSet_iOS_GPUFamily5_v1"
	case MTLFeatureSet_OSX_GPUFamily1_v1:
		return "MTLFeatureSet_OSX_GPUFamily1_v1"
	case MTLFeatureSet_OSX_GPUFamily1_v2:
		return "MTLFeatureSet_OSX_GPUFamily1_v2"
	case MTLFeatureSet_OSX_ReadWriteTextureTier2:
		return "MTLFeatureSet_OSX_ReadWriteTextureTier2"
	case MTLFeatureSet_TVOS_GPUFamily1_v1:
		return "MTLFeatureSet_TVOS_GPUFamily1_v1"
	case MTLFeatureSet_macOS_GPUFamily1_v3:
		return "MTLFeatureSet_macOS_GPUFamily1_v3"
	case MTLFeatureSet_macOS_GPUFamily1_v4:
		return "MTLFeatureSet_macOS_GPUFamily1_v4"
	case MTLFeatureSet_macOS_GPUFamily2_v1:
		return "MTLFeatureSet_macOS_GPUFamily2_v1"
	case MTLFeatureSet_tvOS_GPUFamily1_v2:
		return "MTLFeatureSet_tvOS_GPUFamily1_v2"
	case MTLFeatureSet_tvOS_GPUFamily1_v3:
		return "MTLFeatureSet_tvOS_GPUFamily1_v3"
	case MTLFeatureSet_tvOS_GPUFamily1_v4:
		return "MTLFeatureSet_tvOS_GPUFamily1_v4"
	case MTLFeatureSet_tvOS_GPUFamily2_v1:
		return "MTLFeatureSet_tvOS_GPUFamily2_v1"
	case MTLFeatureSet_tvOS_GPUFamily2_v2:
		return "MTLFeatureSet_tvOS_GPUFamily2_v2"
	default:
		return fmt.Sprintf("MTLFeatureSet(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLFunctionLogType
type MTLFunctionLogType uint

const (
	// MTLFunctionLogTypeValidation: A message related to usage validation.
	MTLFunctionLogTypeValidation MTLFunctionLogType = 0
)

func (e MTLFunctionLogType) String() string {
	switch e {
	case MTLFunctionLogTypeValidation:
		return "MTLFunctionLogTypeValidation"
	default:
		return fmt.Sprintf("MTLFunctionLogType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLFunctionOptions
type MTLFunctionOptions uint

const (
	// MTLFunctionOptionCompileToBinary: An option that instructs the compiler to generate a binary format for dynamic linking.
	MTLFunctionOptionCompileToBinary MTLFunctionOptions = 1
	// MTLFunctionOptionFailOnBinaryArchiveMiss: An option that instructs the compiler to return an error when a GPU function isn’t in a binary archive.
	MTLFunctionOptionFailOnBinaryArchiveMiss MTLFunctionOptions = 4
	// MTLFunctionOptionNone: A sentinel value that represents an empty set of options, which is the default behavior for creating functions.
	MTLFunctionOptionNone MTLFunctionOptions = 0
	// MTLFunctionOptionPipelineIndependent: An option that generates the same function handle across all pipeline states that link a function, which lets you share function tables across pipeline states.
	MTLFunctionOptionPipelineIndependent MTLFunctionOptions = 8
	// MTLFunctionOptionStoreFunctionInMetalPipelinesScript: An option that instructs the compiler to store function information for inspecting binary archives.
	MTLFunctionOptionStoreFunctionInMetalPipelinesScript MTLFunctionOptions = 2
	// Deprecated.
	MTLFunctionOptionStoreFunctionInMetalScript MTLFunctionOptions = 2
)

func (e MTLFunctionOptions) String() string {
	switch e {
	case MTLFunctionOptionCompileToBinary:
		return "MTLFunctionOptionCompileToBinary"
	case MTLFunctionOptionFailOnBinaryArchiveMiss:
		return "MTLFunctionOptionFailOnBinaryArchiveMiss"
	case MTLFunctionOptionNone:
		return "MTLFunctionOptionNone"
	case MTLFunctionOptionPipelineIndependent:
		return "MTLFunctionOptionPipelineIndependent"
	case MTLFunctionOptionStoreFunctionInMetalPipelinesScript:
		return "MTLFunctionOptionStoreFunctionInMetalPipelinesScript"
	default:
		return fmt.Sprintf("MTLFunctionOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLFunctionType
type MTLFunctionType uint

const (
	// MTLFunctionTypeFragment: A fragment function you can use in a render pipeline state object.
	MTLFunctionTypeFragment MTLFunctionType = 2
	// MTLFunctionTypeIntersection: A function you can use in an intersection function table.
	MTLFunctionTypeIntersection MTLFunctionType = 6
	// MTLFunctionTypeKernel: A kernel you can use in a compute pipeline state object.
	MTLFunctionTypeKernel MTLFunctionType = 3
	MTLFunctionTypeMesh   MTLFunctionType = 7
	MTLFunctionTypeObject MTLFunctionType = 8
	// MTLFunctionTypeVertex: A vertex function you can use in a render pipeline state object.
	MTLFunctionTypeVertex MTLFunctionType = 1
	// MTLFunctionTypeVisible: A function you can use in a visible function table.
	MTLFunctionTypeVisible MTLFunctionType = 5
)

func (e MTLFunctionType) String() string {
	switch e {
	case MTLFunctionTypeFragment:
		return "MTLFunctionTypeFragment"
	case MTLFunctionTypeIntersection:
		return "MTLFunctionTypeIntersection"
	case MTLFunctionTypeKernel:
		return "MTLFunctionTypeKernel"
	case MTLFunctionTypeMesh:
		return "MTLFunctionTypeMesh"
	case MTLFunctionTypeObject:
		return "MTLFunctionTypeObject"
	case MTLFunctionTypeVertex:
		return "MTLFunctionTypeVertex"
	case MTLFunctionTypeVisible:
		return "MTLFunctionTypeVisible"
	default:
		return fmt.Sprintf("MTLFunctionType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLGPUFamily
type MTLGPUFamily int

const (
	// MTLGPUFamilyApple1: Represents the Apple family 1 GPU features that correspond to the Apple A7 GPUs.
	MTLGPUFamilyApple1  MTLGPUFamily = 1001
	MTLGPUFamilyApple10 MTLGPUFamily = 1010
	// MTLGPUFamilyApple2: Represents the Apple family 2 GPU features that correspond to the Apple A8 GPUs.
	MTLGPUFamilyApple2 MTLGPUFamily = 1002
	// MTLGPUFamilyApple3: Represents the Apple family 3 GPU features that correspond to the Apple A9 and A10 GPUs.
	MTLGPUFamilyApple3 MTLGPUFamily = 1003
	// MTLGPUFamilyApple4: Represents the Apple family 4 GPU features that correspond to the Apple A11 GPUs.
	MTLGPUFamilyApple4 MTLGPUFamily = 1004
	// MTLGPUFamilyApple5: Represents the Apple family 5 GPU features that correspond to the Apple A12 GPUs.
	MTLGPUFamilyApple5 MTLGPUFamily = 1005
	// MTLGPUFamilyApple6: Represents the Apple family 6 GPU features that correspond to the Apple A13 GPUs.
	MTLGPUFamilyApple6 MTLGPUFamily = 1006
	// MTLGPUFamilyApple7: Represents the Apple family 7 GPU features that correspond to the Apple A14 and M1 GPUs.
	MTLGPUFamilyApple7 MTLGPUFamily = 1007
	// MTLGPUFamilyApple8: Represents the Apple family 8 GPU features that correspond to the Apple A15, A16, and M2 GPUs.
	MTLGPUFamilyApple8 MTLGPUFamily = 1008
	// MTLGPUFamilyApple9: Represents the Apple family 9 GPU features that correspond to the Apple A17, M3, and M4 GPUs.
	MTLGPUFamilyApple9 MTLGPUFamily = 1009
	// MTLGPUFamilyCommon1: Represents the Common family 1 GPU features.
	MTLGPUFamilyCommon1 MTLGPUFamily = 3001
	// MTLGPUFamilyCommon2: Represents the Common family 2 GPU features.
	MTLGPUFamilyCommon2 MTLGPUFamily = 3002
	// MTLGPUFamilyCommon3: Represents the Common family 3 GPU features.
	MTLGPUFamilyCommon3 MTLGPUFamily = 3003
	// MTLGPUFamilyMac2: Represents the Mac family 2 GPU features.
	MTLGPUFamilyMac2 MTLGPUFamily = 2002
	// MTLGPUFamilyMetal3: Represents the Metal 3 features.
	MTLGPUFamilyMetal3 MTLGPUFamily = 5001
	MTLGPUFamilyMetal4 MTLGPUFamily = 5002
	// Deprecated.
	MTLGPUFamilyMac1 MTLGPUFamily = 2001
	// Deprecated.
	MTLGPUFamilyMacCatalyst1 MTLGPUFamily = 4001
	// Deprecated.
	MTLGPUFamilyMacCatalyst2 MTLGPUFamily = 4002
)

func (e MTLGPUFamily) String() string {
	switch e {
	case MTLGPUFamilyApple1:
		return "MTLGPUFamilyApple1"
	case MTLGPUFamilyApple10:
		return "MTLGPUFamilyApple10"
	case MTLGPUFamilyApple2:
		return "MTLGPUFamilyApple2"
	case MTLGPUFamilyApple3:
		return "MTLGPUFamilyApple3"
	case MTLGPUFamilyApple4:
		return "MTLGPUFamilyApple4"
	case MTLGPUFamilyApple5:
		return "MTLGPUFamilyApple5"
	case MTLGPUFamilyApple6:
		return "MTLGPUFamilyApple6"
	case MTLGPUFamilyApple7:
		return "MTLGPUFamilyApple7"
	case MTLGPUFamilyApple8:
		return "MTLGPUFamilyApple8"
	case MTLGPUFamilyApple9:
		return "MTLGPUFamilyApple9"
	case MTLGPUFamilyCommon1:
		return "MTLGPUFamilyCommon1"
	case MTLGPUFamilyCommon2:
		return "MTLGPUFamilyCommon2"
	case MTLGPUFamilyCommon3:
		return "MTLGPUFamilyCommon3"
	case MTLGPUFamilyMac2:
		return "MTLGPUFamilyMac2"
	case MTLGPUFamilyMetal3:
		return "MTLGPUFamilyMetal3"
	case MTLGPUFamilyMetal4:
		return "MTLGPUFamilyMetal4"
	case MTLGPUFamilyMac1:
		return "MTLGPUFamilyMac1"
	case MTLGPUFamilyMacCatalyst1:
		return "MTLGPUFamilyMacCatalyst1"
	case MTLGPUFamilyMacCatalyst2:
		return "MTLGPUFamilyMacCatalyst2"
	default:
		return fmt.Sprintf("MTLGPUFamily(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLHazardTrackingMode
type MTLHazardTrackingMode uint

const (
	// MTLHazardTrackingModeDefault: An option that applies the default tracking behavior in Metal based on the resource or heap type you’re creating.
	MTLHazardTrackingModeDefault MTLHazardTrackingMode = 0
	// MTLHazardTrackingModeTracked: An option that directs Metal to apply runtime safeguards that prevent memory hazards when commands access a resource.
	MTLHazardTrackingModeTracked MTLHazardTrackingMode = 2
	// MTLHazardTrackingModeUntracked: An option that disables automatic memory hazard tracking in Metal for a resource at runtime.
	MTLHazardTrackingModeUntracked MTLHazardTrackingMode = 1
)

func (e MTLHazardTrackingMode) String() string {
	switch e {
	case MTLHazardTrackingModeDefault:
		return "MTLHazardTrackingModeDefault"
	case MTLHazardTrackingModeTracked:
		return "MTLHazardTrackingModeTracked"
	case MTLHazardTrackingModeUntracked:
		return "MTLHazardTrackingModeUntracked"
	default:
		return fmt.Sprintf("MTLHazardTrackingMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLHeapType
type MTLHeapType int

const (
	// MTLHeapTypeAutomatic: A heap that automatically places new resource allocations.
	MTLHeapTypeAutomatic MTLHeapType = 0
	// MTLHeapTypePlacement: The app controls placement of resources on the heap.
	MTLHeapTypePlacement MTLHeapType = 1
	// MTLHeapTypeSparse: The heap contains sparse texture tiles.
	MTLHeapTypeSparse MTLHeapType = 2
)

func (e MTLHeapType) String() string {
	switch e {
	case MTLHeapTypeAutomatic:
		return "MTLHeapTypeAutomatic"
	case MTLHeapTypePlacement:
		return "MTLHeapTypePlacement"
	case MTLHeapTypeSparse:
		return "MTLHeapTypeSparse"
	default:
		return fmt.Sprintf("MTLHeapType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLIOCommandQueueType
type MTLIOCommandQueueType int

const (
	// MTLIOCommandQueueTypeConcurrent: Sets a new input/output command queue’s type to a queue that runs commands concurrently.
	MTLIOCommandQueueTypeConcurrent MTLIOCommandQueueType = 0
	// MTLIOCommandQueueTypeSerial: Sets a new input/output command queue’s type to a queue that runs commands serially.
	MTLIOCommandQueueTypeSerial MTLIOCommandQueueType = 1
)

func (e MTLIOCommandQueueType) String() string {
	switch e {
	case MTLIOCommandQueueTypeConcurrent:
		return "MTLIOCommandQueueTypeConcurrent"
	case MTLIOCommandQueueTypeSerial:
		return "MTLIOCommandQueueTypeSerial"
	default:
		return fmt.Sprintf("MTLIOCommandQueueType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLIOCompressionMethod
type MTLIOCompressionMethod int

const (
	// MTLIOCompressionMethodLZ4: Indicates that a file uses the LZ4 compression algorithm codec.
	MTLIOCompressionMethodLZ4 MTLIOCompressionMethod = 2
	// MTLIOCompressionMethodLZBitmap: Indicates that a file uses the LZBitmap compression algorithm codec.
	MTLIOCompressionMethodLZBitmap MTLIOCompressionMethod = 4
	// MTLIOCompressionMethodLZFSE: Indicates that a file uses the LZFSE compression algorithm codec.
	MTLIOCompressionMethodLZFSE MTLIOCompressionMethod = 1
	// MTLIOCompressionMethodLZMA: Indicates that a file uses the LZMA compression algorithm codec.
	MTLIOCompressionMethodLZMA MTLIOCompressionMethod = 3
	// MTLIOCompressionMethodZlib: Indicates that a file uses the zlib compression algorithm codec.
	MTLIOCompressionMethodZlib MTLIOCompressionMethod = 0
)

func (e MTLIOCompressionMethod) String() string {
	switch e {
	case MTLIOCompressionMethodLZ4:
		return "MTLIOCompressionMethodLZ4"
	case MTLIOCompressionMethodLZBitmap:
		return "MTLIOCompressionMethodLZBitmap"
	case MTLIOCompressionMethodLZFSE:
		return "MTLIOCompressionMethodLZFSE"
	case MTLIOCompressionMethodLZMA:
		return "MTLIOCompressionMethodLZMA"
	case MTLIOCompressionMethodZlib:
		return "MTLIOCompressionMethodZlib"
	default:
		return fmt.Sprintf("MTLIOCompressionMethod(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLIOCompressionStatus
type MTLIOCompressionStatus int

const (
	// MTLIOCompressionStatusComplete: Indicates the compression API successfully flushed and destroyed a compression context.
	MTLIOCompressionStatusComplete MTLIOCompressionStatus = 0
	// MTLIOCompressionStatusError: Indicates the compression API had an error while flushing and destroying a compression context.
	MTLIOCompressionStatusError MTLIOCompressionStatus = 1
)

func (e MTLIOCompressionStatus) String() string {
	switch e {
	case MTLIOCompressionStatusComplete:
		return "MTLIOCompressionStatusComplete"
	case MTLIOCompressionStatusError:
		return "MTLIOCompressionStatusError"
	default:
		return fmt.Sprintf("MTLIOCompressionStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLIOError-swift.struct/Code
type MTLIOError int

const (
	// MTLIOErrorInternal: An error code that represents a problem internal to the Metal framework.
	MTLIOErrorInternal MTLIOError = 2
	// MTLIOErrorURLInvalid: An error code that represents a problem with a file URL.
	MTLIOErrorURLInvalid MTLIOError = 1
)

func (e MTLIOError) String() string {
	switch e {
	case MTLIOErrorInternal:
		return "MTLIOErrorInternal"
	case MTLIOErrorURLInvalid:
		return "MTLIOErrorURLInvalid"
	default:
		return fmt.Sprintf("MTLIOError(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLIOPriority
type MTLIOPriority int

const (
	// MTLIOPriorityHigh: Sets a new input/output command queue’s priority to a high priority.
	MTLIOPriorityHigh MTLIOPriority = 0
	// MTLIOPriorityLow: Designates the low priority for a new input/output command queue.
	MTLIOPriorityLow MTLIOPriority = 2
	// MTLIOPriorityNormal: Designates the normal priority for a new input/output command queue.
	MTLIOPriorityNormal MTLIOPriority = 1
)

func (e MTLIOPriority) String() string {
	switch e {
	case MTLIOPriorityHigh:
		return "MTLIOPriorityHigh"
	case MTLIOPriorityLow:
		return "MTLIOPriorityLow"
	case MTLIOPriorityNormal:
		return "MTLIOPriorityNormal"
	default:
		return fmt.Sprintf("MTLIOPriority(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLIOStatus
type MTLIOStatus int

const (
	// MTLIOStatusCancelled: Indicates the GPU has successfully abandoned the input/output command buffer.
	MTLIOStatusCancelled MTLIOStatus = 1
	// MTLIOStatusComplete: Indicates the GPU has successfully finished executing the input/output command buffer.
	MTLIOStatusComplete MTLIOStatus = 3
	// MTLIOStatusError: Indicates the GPU experienced a problem with the input/output command buffer.
	MTLIOStatusError MTLIOStatus = 2
	// MTLIOStatusPending: Indicates the GPU hasn’t finished executing the input/output command buffer.
	MTLIOStatusPending MTLIOStatus = 0
)

func (e MTLIOStatus) String() string {
	switch e {
	case MTLIOStatusCancelled:
		return "MTLIOStatusCancelled"
	case MTLIOStatusComplete:
		return "MTLIOStatusComplete"
	case MTLIOStatusError:
		return "MTLIOStatusError"
	case MTLIOStatusPending:
		return "MTLIOStatusPending"
	default:
		return fmt.Sprintf("MTLIOStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLIndexType
type MTLIndexType uint

const (
	// MTLIndexTypeUInt16: A 16-bit unsigned integer used as a primitive index.
	MTLIndexTypeUInt16 MTLIndexType = 0
	// MTLIndexTypeUInt32: A 32-bit unsigned integer used as a primitive index.
	MTLIndexTypeUInt32 MTLIndexType = 1
)

func (e MTLIndexType) String() string {
	switch e {
	case MTLIndexTypeUInt16:
		return "MTLIndexTypeUInt16"
	case MTLIndexTypeUInt32:
		return "MTLIndexTypeUInt32"
	default:
		return fmt.Sprintf("MTLIndexType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLIndirectCommandType
type MTLIndirectCommandType uint

const (
	// MTLIndirectCommandTypeConcurrentDispatch: A compute command using a grid aligned to threadgroup boundaries.
	MTLIndirectCommandTypeConcurrentDispatch MTLIndirectCommandType = 32
	// MTLIndirectCommandTypeConcurrentDispatchThreads: A compute command using an arbitrarily sized grid.
	MTLIndirectCommandTypeConcurrentDispatchThreads MTLIndirectCommandType = 64
	// MTLIndirectCommandTypeDraw: A draw call command.
	MTLIndirectCommandTypeDraw MTLIndirectCommandType = 1
	// MTLIndirectCommandTypeDrawIndexed: An indexed draw call command.
	MTLIndirectCommandTypeDrawIndexed MTLIndirectCommandType = 2
	// MTLIndirectCommandTypeDrawIndexedPatches: An indexed draw call command for tessellated patches.
	MTLIndirectCommandTypeDrawIndexedPatches   MTLIndirectCommandType = 8
	MTLIndirectCommandTypeDrawMeshThreadgroups MTLIndirectCommandType = 128
	MTLIndirectCommandTypeDrawMeshThreads      MTLIndirectCommandType = 256
	// MTLIndirectCommandTypeDrawPatches: A draw call command for tessellated patches.
	MTLIndirectCommandTypeDrawPatches MTLIndirectCommandType = 4
)

func (e MTLIndirectCommandType) String() string {
	switch e {
	case MTLIndirectCommandTypeConcurrentDispatch:
		return "MTLIndirectCommandTypeConcurrentDispatch"
	case MTLIndirectCommandTypeConcurrentDispatchThreads:
		return "MTLIndirectCommandTypeConcurrentDispatchThreads"
	case MTLIndirectCommandTypeDraw:
		return "MTLIndirectCommandTypeDraw"
	case MTLIndirectCommandTypeDrawIndexed:
		return "MTLIndirectCommandTypeDrawIndexed"
	case MTLIndirectCommandTypeDrawIndexedPatches:
		return "MTLIndirectCommandTypeDrawIndexedPatches"
	case MTLIndirectCommandTypeDrawMeshThreadgroups:
		return "MTLIndirectCommandTypeDrawMeshThreadgroups"
	case MTLIndirectCommandTypeDrawMeshThreads:
		return "MTLIndirectCommandTypeDrawMeshThreads"
	case MTLIndirectCommandTypeDrawPatches:
		return "MTLIndirectCommandTypeDrawPatches"
	default:
		return fmt.Sprintf("MTLIndirectCommandType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLIntersectionFunctionSignature
type MTLIntersectionFunctionSignature uint

const (
	MTLIntersectionFunctionSignatureCurveData      MTLIntersectionFunctionSignature = 128
	MTLIntersectionFunctionSignatureExtendedLimits MTLIntersectionFunctionSignature = 32
	MTLIntersectionFunctionSignatureInstanceMotion MTLIntersectionFunctionSignature = 8
	// MTLIntersectionFunctionSignatureInstancing: A flag indicating that function signature uses instancing.
	MTLIntersectionFunctionSignatureInstancing MTLIntersectionFunctionSignature = 1
	// MTLIntersectionFunctionSignatureIntersectionFunctionBuffer: # Discussion
	MTLIntersectionFunctionSignatureIntersectionFunctionBuffer MTLIntersectionFunctionSignature = 256
	MTLIntersectionFunctionSignatureMaxLevels                  MTLIntersectionFunctionSignature = 64
	// MTLIntersectionFunctionSignatureNone: A constant indicating that the function uses the default signature.
	MTLIntersectionFunctionSignatureNone            MTLIntersectionFunctionSignature = 0
	MTLIntersectionFunctionSignaturePrimitiveMotion MTLIntersectionFunctionSignature = 16
	// MTLIntersectionFunctionSignatureTriangleData: A flag indicating that function signature uses triangle data.
	MTLIntersectionFunctionSignatureTriangleData MTLIntersectionFunctionSignature = 2
	// MTLIntersectionFunctionSignatureUserData: # Discussion
	MTLIntersectionFunctionSignatureUserData MTLIntersectionFunctionSignature = 512
	// MTLIntersectionFunctionSignatureWorldSpaceData: A flag indicating that function signature uses world space data.
	MTLIntersectionFunctionSignatureWorldSpaceData MTLIntersectionFunctionSignature = 4
)

func (e MTLIntersectionFunctionSignature) String() string {
	switch e {
	case MTLIntersectionFunctionSignatureCurveData:
		return "MTLIntersectionFunctionSignatureCurveData"
	case MTLIntersectionFunctionSignatureExtendedLimits:
		return "MTLIntersectionFunctionSignatureExtendedLimits"
	case MTLIntersectionFunctionSignatureInstanceMotion:
		return "MTLIntersectionFunctionSignatureInstanceMotion"
	case MTLIntersectionFunctionSignatureInstancing:
		return "MTLIntersectionFunctionSignatureInstancing"
	case MTLIntersectionFunctionSignatureIntersectionFunctionBuffer:
		return "MTLIntersectionFunctionSignatureIntersectionFunctionBuffer"
	case MTLIntersectionFunctionSignatureMaxLevels:
		return "MTLIntersectionFunctionSignatureMaxLevels"
	case MTLIntersectionFunctionSignatureNone:
		return "MTLIntersectionFunctionSignatureNone"
	case MTLIntersectionFunctionSignaturePrimitiveMotion:
		return "MTLIntersectionFunctionSignaturePrimitiveMotion"
	case MTLIntersectionFunctionSignatureTriangleData:
		return "MTLIntersectionFunctionSignatureTriangleData"
	case MTLIntersectionFunctionSignatureUserData:
		return "MTLIntersectionFunctionSignatureUserData"
	case MTLIntersectionFunctionSignatureWorldSpaceData:
		return "MTLIntersectionFunctionSignatureWorldSpaceData"
	default:
		return fmt.Sprintf("MTLIntersectionFunctionSignature(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLLanguageVersion
type MTLLanguageVersion uint

const (
	// MTLLanguageVersion1_1: Version 1.1 of the Metal shading language.
	MTLLanguageVersion1_1 MTLLanguageVersion = 65536
	// MTLLanguageVersion1_2: Version 1.2 of the Metal shading language.
	MTLLanguageVersion1_2 MTLLanguageVersion = 65536
	// MTLLanguageVersion2_0: Version 2.0 of the Metal shading language.
	MTLLanguageVersion2_0 MTLLanguageVersion = 131072
	// MTLLanguageVersion2_1: Version 2.1 of the Metal shading language.
	MTLLanguageVersion2_1 MTLLanguageVersion = 131072
	// MTLLanguageVersion2_2: Version 2.2 of the Metal shading language.
	MTLLanguageVersion2_2 MTLLanguageVersion = 131072
	// MTLLanguageVersion2_3: Version 2.3 of the Metal shading language.
	MTLLanguageVersion2_3 MTLLanguageVersion = 131072
	// MTLLanguageVersion2_4: Version 2.4 of the Metal shading language.
	MTLLanguageVersion2_4 MTLLanguageVersion = 131072
	// MTLLanguageVersion3_0: Version 3.0 of the Metal shading language.
	MTLLanguageVersion3_0 MTLLanguageVersion = 131073
	// MTLLanguageVersion3_1: Version 3.1 of the Metal shading language.
	MTLLanguageVersion3_1 MTLLanguageVersion = 131074
	// MTLLanguageVersion3_2: Version 3.2 of the Metal shading language.
	MTLLanguageVersion3_2 MTLLanguageVersion = 131075
	MTLLanguageVersion4_0 MTLLanguageVersion = 131076
	// Deprecated.
	MTLLanguageVersion1_0 MTLLanguageVersion = 65536
)

func (e MTLLanguageVersion) String() string {
	switch e {
	case MTLLanguageVersion1_1:
		return "MTLLanguageVersion1_1"
	case MTLLanguageVersion2_0:
		return "MTLLanguageVersion2_0"
	case MTLLanguageVersion3_0:
		return "MTLLanguageVersion3_0"
	case MTLLanguageVersion3_1:
		return "MTLLanguageVersion3_1"
	case MTLLanguageVersion3_2:
		return "MTLLanguageVersion3_2"
	case MTLLanguageVersion4_0:
		return "MTLLanguageVersion4_0"
	default:
		return fmt.Sprintf("MTLLanguageVersion(%d)", e)
	}
}

type MTLLibraryError int

const (
	MTLLibraryErrorCompileFailure   MTLLibraryError = 3
	MTLLibraryErrorCompileWarning   MTLLibraryError = 4
	MTLLibraryErrorFileNotFound     MTLLibraryError = 6
	MTLLibraryErrorFunctionNotFound MTLLibraryError = 5
	// MTLLibraryErrorInternal: The action caused an internal error.
	MTLLibraryErrorInternal MTLLibraryError = 2
	// MTLLibraryErrorUnsupported: Metal couldn’t support the requested action.
	MTLLibraryErrorUnsupported MTLLibraryError = 1
)

func (e MTLLibraryError) String() string {
	switch e {
	case MTLLibraryErrorCompileFailure:
		return "MTLLibraryErrorCompileFailure"
	case MTLLibraryErrorCompileWarning:
		return "MTLLibraryErrorCompileWarning"
	case MTLLibraryErrorFileNotFound:
		return "MTLLibraryErrorFileNotFound"
	case MTLLibraryErrorFunctionNotFound:
		return "MTLLibraryErrorFunctionNotFound"
	case MTLLibraryErrorInternal:
		return "MTLLibraryErrorInternal"
	case MTLLibraryErrorUnsupported:
		return "MTLLibraryErrorUnsupported"
	default:
		return fmt.Sprintf("MTLLibraryError(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLLibraryOptimizationLevel
type MTLLibraryOptimizationLevel int

const (
	// MTLLibraryOptimizationLevelDefault: An optimization option for the Metal compiler that prioritizes runtime performance.
	MTLLibraryOptimizationLevelDefault MTLLibraryOptimizationLevel = 0
	// MTLLibraryOptimizationLevelSize: An optimization option for the Metal compiler that prioritizes minimizing the size of its output binaries, which may also reduce compile time.
	MTLLibraryOptimizationLevelSize MTLLibraryOptimizationLevel = 1
)

func (e MTLLibraryOptimizationLevel) String() string {
	switch e {
	case MTLLibraryOptimizationLevelDefault:
		return "MTLLibraryOptimizationLevelDefault"
	case MTLLibraryOptimizationLevelSize:
		return "MTLLibraryOptimizationLevelSize"
	default:
		return fmt.Sprintf("MTLLibraryOptimizationLevel(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLLibraryType
type MTLLibraryType int

const (
	// MTLLibraryTypeDynamic: A library that you can dynamically link to from other libraries.
	MTLLibraryTypeDynamic MTLLibraryType = 1
	// MTLLibraryTypeExecutable: A library that can create pipeline state objects.
	MTLLibraryTypeExecutable MTLLibraryType = 0
)

func (e MTLLibraryType) String() string {
	switch e {
	case MTLLibraryTypeDynamic:
		return "MTLLibraryTypeDynamic"
	case MTLLibraryTypeExecutable:
		return "MTLLibraryTypeExecutable"
	default:
		return fmt.Sprintf("MTLLibraryType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLLoadAction
type MTLLoadAction uint

const (
	// MTLLoadActionClear: The GPU writes a value to every pixel in the attachment at the start of the render pass.
	MTLLoadActionClear MTLLoadAction = 2
	// MTLLoadActionDontCare: The GPU has permission to discard the existing contents of the attachment at the start of the render pass, replacing them with arbitrary data.
	MTLLoadActionDontCare MTLLoadAction = 0
	// MTLLoadActionLoad: The GPU preserves the existing contents of the attachment at the start of the render pass.
	MTLLoadActionLoad MTLLoadAction = 1
)

func (e MTLLoadAction) String() string {
	switch e {
	case MTLLoadActionClear:
		return "MTLLoadActionClear"
	case MTLLoadActionDontCare:
		return "MTLLoadActionDontCare"
	case MTLLoadActionLoad:
		return "MTLLoadActionLoad"
	default:
		return fmt.Sprintf("MTLLoadAction(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLLogLevel
type MTLLogLevel int

const (
	// MTLLogLevelDebug: The log level that captures diagnostic information.
	MTLLogLevelDebug MTLLogLevel = 1
	// MTLLogLevelError: The log level that captures error information.
	MTLLogLevelError MTLLogLevel = 4
	// MTLLogLevelFault: The log level that captures fault information.
	MTLLogLevelFault MTLLogLevel = 5
	// MTLLogLevelInfo: The log level that captures additional information.
	MTLLogLevelInfo MTLLogLevel = 2
	// MTLLogLevelNotice: The log level that captures notifications.
	MTLLogLevelNotice MTLLogLevel = 3
	// MTLLogLevelUndefined: The log level when the log level hasn’t been configured.
	MTLLogLevelUndefined MTLLogLevel = 0
)

func (e MTLLogLevel) String() string {
	switch e {
	case MTLLogLevelDebug:
		return "MTLLogLevelDebug"
	case MTLLogLevelError:
		return "MTLLogLevelError"
	case MTLLogLevelFault:
		return "MTLLogLevelFault"
	case MTLLogLevelInfo:
		return "MTLLogLevelInfo"
	case MTLLogLevelNotice:
		return "MTLLogLevelNotice"
	case MTLLogLevelUndefined:
		return "MTLLogLevelUndefined"
	default:
		return fmt.Sprintf("MTLLogLevel(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLLogStateError
type MTLLogStateError int

const (
	MTLLogStateErrorInvalid     MTLLogStateError = 2
	MTLLogStateErrorInvalidSize MTLLogStateError = 1
)

func (e MTLLogStateError) String() string {
	switch e {
	case MTLLogStateErrorInvalid:
		return "MTLLogStateErrorInvalid"
	case MTLLogStateErrorInvalidSize:
		return "MTLLogStateErrorInvalidSize"
	default:
		return fmt.Sprintf("MTLLogStateError(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLMathFloatingPointFunctions
type MTLMathFloatingPointFunctions int

const (
	// MTLMathFloatingPointFunctionsFast: An indication that Metal uses the fast version of the 32b floating-point math functions.
	MTLMathFloatingPointFunctionsFast MTLMathFloatingPointFunctions = 0
	// MTLMathFloatingPointFunctionsPrecise: An indication that Metal uses the precise version of the 32b floating-point math functions.
	MTLMathFloatingPointFunctionsPrecise MTLMathFloatingPointFunctions = 1
)

func (e MTLMathFloatingPointFunctions) String() string {
	switch e {
	case MTLMathFloatingPointFunctionsFast:
		return "MTLMathFloatingPointFunctionsFast"
	case MTLMathFloatingPointFunctionsPrecise:
		return "MTLMathFloatingPointFunctionsPrecise"
	default:
		return fmt.Sprintf("MTLMathFloatingPointFunctions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLMathMode
type MTLMathMode int

const (
	// MTLMathModeFast: An indicator of the mode the compiler uses to make aggressive, potentially lossy assumptions about floating-point math.
	MTLMathModeFast MTLMathMode = 2
	// MTLMathModeRelaxed: An indicator of the mode the compiler uses to make aggressive, potentially lossy assumptions about floating-point math, while honoring Inf/NaN.
	MTLMathModeRelaxed MTLMathMode = 1
	// MTLMathModeSafe: An indicator of the mode the compiler uses to disable unsafe floating-point optimizations by preventing the compiler from making any transformations that could affect the results.
	MTLMathModeSafe MTLMathMode = 0
)

func (e MTLMathMode) String() string {
	switch e {
	case MTLMathModeFast:
		return "MTLMathModeFast"
	case MTLMathModeRelaxed:
		return "MTLMathModeRelaxed"
	case MTLMathModeSafe:
		return "MTLMathModeSafe"
	default:
		return fmt.Sprintf("MTLMathMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLMatrixLayout
type MTLMatrixLayout int

const (
	MTLMatrixLayoutColumnMajor MTLMatrixLayout = 0
	MTLMatrixLayoutRowMajor    MTLMatrixLayout = 1
)

func (e MTLMatrixLayout) String() string {
	switch e {
	case MTLMatrixLayoutColumnMajor:
		return "MTLMatrixLayoutColumnMajor"
	case MTLMatrixLayoutRowMajor:
		return "MTLMatrixLayoutRowMajor"
	default:
		return fmt.Sprintf("MTLMatrixLayout(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLMotionBorderMode
type MTLMotionBorderMode uint32

const (
	// MTLMotionBorderModeClamp: A mode that specifies treating times outside the specified endpoint as if they were at the endpoint.
	MTLMotionBorderModeClamp MTLMotionBorderMode = 0
	// MTLMotionBorderModeVanish: A mode that specifies that times outside the specified endpoint need to prevent any ray-intersections with the primitive.
	MTLMotionBorderModeVanish MTLMotionBorderMode = 1
)

func (e MTLMotionBorderMode) String() string {
	switch e {
	case MTLMotionBorderModeClamp:
		return "MTLMotionBorderModeClamp"
	case MTLMotionBorderModeVanish:
		return "MTLMotionBorderModeVanish"
	default:
		return fmt.Sprintf("MTLMotionBorderMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLMultisampleDepthResolveFilter
type MTLMultisampleDepthResolveFilter uint

const (
	// MTLMultisampleDepthResolveFilterMax: The GPU compares all depth samples in the pixel and selects the sample with the largest value.
	MTLMultisampleDepthResolveFilterMax MTLMultisampleDepthResolveFilter = 2
	// MTLMultisampleDepthResolveFilterMin: The GPU compares all depth samples in the pixel and selects the sample with the smallest value.
	MTLMultisampleDepthResolveFilterMin MTLMultisampleDepthResolveFilter = 1
	// MTLMultisampleDepthResolveFilterSample0: No filter is applied.
	MTLMultisampleDepthResolveFilterSample0 MTLMultisampleDepthResolveFilter = 0
)

func (e MTLMultisampleDepthResolveFilter) String() string {
	switch e {
	case MTLMultisampleDepthResolveFilterMax:
		return "MTLMultisampleDepthResolveFilterMax"
	case MTLMultisampleDepthResolveFilterMin:
		return "MTLMultisampleDepthResolveFilterMin"
	case MTLMultisampleDepthResolveFilterSample0:
		return "MTLMultisampleDepthResolveFilterSample0"
	default:
		return fmt.Sprintf("MTLMultisampleDepthResolveFilter(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLMultisampleStencilResolveFilter
type MTLMultisampleStencilResolveFilter uint

const (
	// MTLMultisampleStencilResolveFilterDepthResolvedSample: Chooses the stencil sample corresponding to the depth sample selected by the depth resolve filter.
	MTLMultisampleStencilResolveFilterDepthResolvedSample MTLMultisampleStencilResolveFilter = 1
	// MTLMultisampleStencilResolveFilterSample0: Chooses the first stencil sample in the pixel.
	MTLMultisampleStencilResolveFilterSample0 MTLMultisampleStencilResolveFilter = 0
)

func (e MTLMultisampleStencilResolveFilter) String() string {
	switch e {
	case MTLMultisampleStencilResolveFilterDepthResolvedSample:
		return "MTLMultisampleStencilResolveFilterDepthResolvedSample"
	case MTLMultisampleStencilResolveFilterSample0:
		return "MTLMultisampleStencilResolveFilterSample0"
	default:
		return fmt.Sprintf("MTLMultisampleStencilResolveFilter(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLMutability
type MTLMutability uint

const (
	// MTLMutabilityDefault: The default behavior, based on the buffer’s type.
	MTLMutabilityDefault MTLMutability = 0
	// MTLMutabilityImmutable: An option that states that you can’t modify the buffer’s contents.
	MTLMutabilityImmutable MTLMutability = 2
	// MTLMutabilityMutable: An option that states that you can modify the buffer’s contents.
	MTLMutabilityMutable MTLMutability = 1
)

func (e MTLMutability) String() string {
	switch e {
	case MTLMutabilityDefault:
		return "MTLMutabilityDefault"
	case MTLMutabilityImmutable:
		return "MTLMutabilityImmutable"
	case MTLMutabilityMutable:
		return "MTLMutabilityMutable"
	default:
		return fmt.Sprintf("MTLMutability(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLPatchType
type MTLPatchType uint

const (
	// MTLPatchTypeNone: An option that indicates that this isn’t a post-tessellation vertex function.
	MTLPatchTypeNone MTLPatchType = 0
	// MTLPatchTypeQuad: A quad patch.
	MTLPatchTypeQuad MTLPatchType = 2
	// MTLPatchTypeTriangle: A triangle patch.
	MTLPatchTypeTriangle MTLPatchType = 1
)

func (e MTLPatchType) String() string {
	switch e {
	case MTLPatchTypeNone:
		return "MTLPatchTypeNone"
	case MTLPatchTypeQuad:
		return "MTLPatchTypeQuad"
	case MTLPatchTypeTriangle:
		return "MTLPatchTypeTriangle"
	default:
		return fmt.Sprintf("MTLPatchType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLPipelineOption
type MTLPipelineOption uint

const (
	// MTLPipelineOptionBindingInfo: An option that provides binding information for pipeline state resources.
	MTLPipelineOptionBindingInfo MTLPipelineOption = 1
	// MTLPipelineOptionBufferTypeInfo: An option instance that provides detailed buffer type information for buffer arguments.
	MTLPipelineOptionBufferTypeInfo MTLPipelineOption = 2
	// MTLPipelineOptionFailOnBinaryArchiveMiss: An option that instructs the compiler to return an error when a GPU function isn’t in a binary archive.
	MTLPipelineOptionFailOnBinaryArchiveMiss MTLPipelineOption = 4
	// MTLPipelineOptionNone: Don’t provide any reflection information.
	MTLPipelineOptionNone MTLPipelineOption = 0
	// Deprecated.
	MTLPipelineOptionArgumentInfo MTLPipelineOption = 1
)

func (e MTLPipelineOption) String() string {
	switch e {
	case MTLPipelineOptionBindingInfo:
		return "MTLPipelineOptionBindingInfo"
	case MTLPipelineOptionBufferTypeInfo:
		return "MTLPipelineOptionBufferTypeInfo"
	case MTLPipelineOptionFailOnBinaryArchiveMiss:
		return "MTLPipelineOptionFailOnBinaryArchiveMiss"
	case MTLPipelineOptionNone:
		return "MTLPipelineOptionNone"
	default:
		return fmt.Sprintf("MTLPipelineOption(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLPixelFormat
type MTLPixelFormat uint

const (
	// MTLPixelFormatA1BGR5Unorm: Packed 16-bit format with normalized unsigned integer color components: 5 bits each for BGR and 1 for alpha, packed into 16 bits.
	MTLPixelFormatA1BGR5Unorm MTLPixelFormat = 41
	// MTLPixelFormatA8Unorm: Ordinary format with one 8-bit normalized unsigned integer component.
	MTLPixelFormatA8Unorm MTLPixelFormat = 1
	// MTLPixelFormatABGR4Unorm: Packed 16-bit format with normalized unsigned integer color components: 4 bits each for ABGR, packed into 16 bits.
	MTLPixelFormatABGR4Unorm MTLPixelFormat = 42
	// MTLPixelFormatASTC_10x10_HDR: ASTC-compressed format with high-dynamic range content, a block width of 10, and a block height of 10.
	MTLPixelFormatASTC_10x10_HDR MTLPixelFormat = 234
	// MTLPixelFormatASTC_10x10_LDR: ASTC-compressed format with low-dynamic-range content, a block width of 10, and a block height of 10.
	MTLPixelFormatASTC_10x10_LDR MTLPixelFormat = 216
	// MTLPixelFormatASTC_10x10_sRGB: ASTC-compressed format with low-dynamic-range content, conversion between sRGB and linear space, a block width of 10, and a block height of 10.
	MTLPixelFormatASTC_10x10_sRGB MTLPixelFormat = 198
	// MTLPixelFormatASTC_10x5_HDR: ASTC-compressed format with high-dynamic range content, a block width of 10, and a block height of 5.
	MTLPixelFormatASTC_10x5_HDR MTLPixelFormat = 231
	// MTLPixelFormatASTC_10x5_LDR: ASTC-compressed format with low-dynamic-range content, a block width of 10, and a block height of 5.
	MTLPixelFormatASTC_10x5_LDR MTLPixelFormat = 213
	// MTLPixelFormatASTC_10x5_sRGB: ASTC-compressed format with low-dynamic-range content, conversion between sRGB and linear space, a block width of 10, and a block height of 5.
	MTLPixelFormatASTC_10x5_sRGB MTLPixelFormat = 195
	// MTLPixelFormatASTC_10x6_HDR: ASTC-compressed format with high-dynamic range content, a block width of 10, and a block height of 6.
	MTLPixelFormatASTC_10x6_HDR MTLPixelFormat = 232
	// MTLPixelFormatASTC_10x6_LDR: ASTC-compressed format with low-dynamic-range content, a block width of 10, and a block height of 6.
	MTLPixelFormatASTC_10x6_LDR MTLPixelFormat = 214
	// MTLPixelFormatASTC_10x6_sRGB: ASTC-compressed format with low-dynamic-range content, conversion between sRGB and linear space, a block width of 10, and a block height of 6.
	MTLPixelFormatASTC_10x6_sRGB MTLPixelFormat = 196
	// MTLPixelFormatASTC_10x8_HDR: ASTC-compressed format with high-dynamic range content, a block width of 10, and a block height of 8.
	MTLPixelFormatASTC_10x8_HDR MTLPixelFormat = 233
	// MTLPixelFormatASTC_10x8_LDR: ASTC-compressed format with low-dynamic-range content, a block width of 10, and a block height of 8.
	MTLPixelFormatASTC_10x8_LDR MTLPixelFormat = 215
	// MTLPixelFormatASTC_10x8_sRGB: ASTC-compressed format with low-dynamic-range content, conversion between sRGB and linear space, a block width of 10, and a block height of 8.
	MTLPixelFormatASTC_10x8_sRGB MTLPixelFormat = 197
	// MTLPixelFormatASTC_12x10_HDR: ASTC-compressed format with high-dynamic range content, a block width of 12, and a block height of 10.
	MTLPixelFormatASTC_12x10_HDR MTLPixelFormat = 235
	// MTLPixelFormatASTC_12x10_LDR: ASTC-compressed format with low-dynamic-range content, a block width of 12, and a block height of 10.
	MTLPixelFormatASTC_12x10_LDR MTLPixelFormat = 217
	// MTLPixelFormatASTC_12x10_sRGB: ASTC-compressed format with low-dynamic-range content, conversion between sRGB and linear space, a block width of 12, and a block height of 10.
	MTLPixelFormatASTC_12x10_sRGB MTLPixelFormat = 199
	// MTLPixelFormatASTC_12x12_HDR: ASTC-compressed format with high-dynamic range content, a block width of 12, and a block height of 12.
	MTLPixelFormatASTC_12x12_HDR MTLPixelFormat = 236
	// MTLPixelFormatASTC_12x12_LDR: ASTC-compressed format with low-dynamic-range content, a block width of 12, and a block height of 12.
	MTLPixelFormatASTC_12x12_LDR MTLPixelFormat = 218
	// MTLPixelFormatASTC_12x12_sRGB: ASTC-compressed format with low-dynamic-range content, conversion between sRGB and linear space, a block width of 12, and a block height of 12.
	MTLPixelFormatASTC_12x12_sRGB MTLPixelFormat = 200
	// MTLPixelFormatASTC_4x4_HDR: ASTC-compressed format with high-dynamic-range content, a block width of 4, and a block height of 4.
	MTLPixelFormatASTC_4x4_HDR MTLPixelFormat = 222
	// MTLPixelFormatASTC_4x4_LDR: ASTC-compressed format with low-dynamic-range content, a block width of 4, and a block height of 4.
	MTLPixelFormatASTC_4x4_LDR MTLPixelFormat = 204
	// MTLPixelFormatASTC_4x4_sRGB: ASTC-compressed format with low-dynamic-range content, conversion between sRGB and linear space, a block width of 4, and a block height of 4.
	MTLPixelFormatASTC_4x4_sRGB MTLPixelFormat = 186
	// MTLPixelFormatASTC_5x4_HDR: ASTC-compressed format with high-dynamic range content, a block width of 5, and a block height of 4.
	MTLPixelFormatASTC_5x4_HDR MTLPixelFormat = 223
	// MTLPixelFormatASTC_5x4_LDR: ASTC-compressed format with low-dynamic-range content, a block width of 5, and a block height of 4.
	MTLPixelFormatASTC_5x4_LDR MTLPixelFormat = 205
	// MTLPixelFormatASTC_5x4_sRGB: ASTC-compressed format with low-dynamic-range content, conversion between sRGB and linear space, a block width of 5, and a block height of 4.
	MTLPixelFormatASTC_5x4_sRGB MTLPixelFormat = 187
	// MTLPixelFormatASTC_5x5_HDR: ASTC-compressed format with high-dynamic range content, a block width of 5, and a block height of 5.
	MTLPixelFormatASTC_5x5_HDR MTLPixelFormat = 224
	// MTLPixelFormatASTC_5x5_LDR: ASTC-compressed format with low-dynamic-range content, a block width of 5, and a block height of 5.
	MTLPixelFormatASTC_5x5_LDR MTLPixelFormat = 206
	// MTLPixelFormatASTC_5x5_sRGB: ASTC-compressed format with low-dynamic-range content, conversion between sRGB and linear space, a block width of 5, and a block height of 5.
	MTLPixelFormatASTC_5x5_sRGB MTLPixelFormat = 188
	// MTLPixelFormatASTC_6x5_HDR: ASTC-compressed format with high-dynamic range content, a block width of 6, and a block height of 5.
	MTLPixelFormatASTC_6x5_HDR MTLPixelFormat = 225
	// MTLPixelFormatASTC_6x5_LDR: ASTC-compressed format with low-dynamic-range content, a block width of 6, and a block height of 5.
	MTLPixelFormatASTC_6x5_LDR MTLPixelFormat = 207
	// MTLPixelFormatASTC_6x5_sRGB: ASTC-compressed format with low-dynamic-range content, conversion between sRGB and linear space, a block width of 6, and a block height of 5.
	MTLPixelFormatASTC_6x5_sRGB MTLPixelFormat = 189
	// MTLPixelFormatASTC_6x6_HDR: ASTC-compressed format with high-dynamic range content, a block width of 6, and a block height of 6.
	MTLPixelFormatASTC_6x6_HDR MTLPixelFormat = 226
	// MTLPixelFormatASTC_6x6_LDR: ASTC-compressed format with low-dynamic-range content, a block width of 6, and a block height of 6.
	MTLPixelFormatASTC_6x6_LDR MTLPixelFormat = 208
	// MTLPixelFormatASTC_6x6_sRGB: ASTC-compressed format with low-dynamic-range content, conversion between sRGB and linear space, a block width of 6, and a block height of 6.
	MTLPixelFormatASTC_6x6_sRGB MTLPixelFormat = 190
	// MTLPixelFormatASTC_8x5_HDR: ASTC-compressed format with high-dynamic range content, a block width of 8, and a block height of 5.
	MTLPixelFormatASTC_8x5_HDR MTLPixelFormat = 228
	// MTLPixelFormatASTC_8x5_LDR: ASTC-compressed format with low-dynamic-range content, a block width of 8, and a block height of 5.
	MTLPixelFormatASTC_8x5_LDR MTLPixelFormat = 210
	// MTLPixelFormatASTC_8x5_sRGB: ASTC-compressed format with low-dynamic-range content, conversion between sRGB and linear space, a block width of 8, and a block height of 5.
	MTLPixelFormatASTC_8x5_sRGB MTLPixelFormat = 192
	// MTLPixelFormatASTC_8x6_HDR: ASTC-compressed format with high-dynamic range content, a block width of 8, and a block height of 6.
	MTLPixelFormatASTC_8x6_HDR MTLPixelFormat = 229
	// MTLPixelFormatASTC_8x6_LDR: ASTC-compressed format with low-dynamic-range content, a block width of 8, and a block height of 6.
	MTLPixelFormatASTC_8x6_LDR MTLPixelFormat = 211
	// MTLPixelFormatASTC_8x6_sRGB: ASTC-compressed format with low-dynamic-range content, conversion between sRGB and linear space, a block width of 8, and a block height of 6.
	MTLPixelFormatASTC_8x6_sRGB MTLPixelFormat = 193
	// MTLPixelFormatASTC_8x8_HDR: ASTC-compressed format with high-dynamic range content, a block width of 8, and a block height of 8.
	MTLPixelFormatASTC_8x8_HDR MTLPixelFormat = 230
	// MTLPixelFormatASTC_8x8_LDR: ASTC-compressed format with low-dynamic-range content, a block width of 8, and a block height of 8.
	MTLPixelFormatASTC_8x8_LDR MTLPixelFormat = 212
	// MTLPixelFormatASTC_8x8_sRGB: ASTC-compressed format with low-dynamic-range content, conversion between sRGB and linear space, a block width of 8, and a block height of 8.
	MTLPixelFormatASTC_8x8_sRGB MTLPixelFormat = 194
	// MTLPixelFormatB5G6R5Unorm: Packed 16-bit format with normalized unsigned integer color components: 5 bits for blue, 6 bits for green, 5 bits for red, packed into 16 bits.
	MTLPixelFormatB5G6R5Unorm MTLPixelFormat = 40
	// MTLPixelFormatBC1_RGBA: Compressed format with two 16-bit color components and one 32-bit descriptor component.
	MTLPixelFormatBC1_RGBA MTLPixelFormat = 130
	// MTLPixelFormatBC1_RGBA_sRGB: Compressed format with two 16-bit color components and one 32-bit descriptor component, with conversion between sRGB and linear space.
	MTLPixelFormatBC1_RGBA_sRGB MTLPixelFormat = 131
	// MTLPixelFormatBC2_RGBA: Compressed format with two 64-bit chunks.
	MTLPixelFormatBC2_RGBA MTLPixelFormat = 132
	// MTLPixelFormatBC2_RGBA_sRGB: Compressed format with two 64-bit chunks, with conversion between sRGB and linear space.
	MTLPixelFormatBC2_RGBA_sRGB MTLPixelFormat = 133
	// MTLPixelFormatBC3_RGBA: Compressed format with two 64-bit chunks.
	MTLPixelFormatBC3_RGBA MTLPixelFormat = 134
	// MTLPixelFormatBC3_RGBA_sRGB: Compressed format with two 64-bit chunks, with conversion between sRGB and linear space.
	MTLPixelFormatBC3_RGBA_sRGB MTLPixelFormat = 135
	// MTLPixelFormatBC4_RSnorm: Compressed format with one normalized signed integer component.
	MTLPixelFormatBC4_RSnorm MTLPixelFormat = 141
	// MTLPixelFormatBC4_RUnorm: Compressed format with one normalized unsigned integer component.
	MTLPixelFormatBC4_RUnorm MTLPixelFormat = 140
	// MTLPixelFormatBC5_RGSnorm: Compressed format with two normalized signed integer components.
	MTLPixelFormatBC5_RGSnorm MTLPixelFormat = 143
	// MTLPixelFormatBC5_RGUnorm: Compressed format with two normalized unsigned integer components.
	MTLPixelFormatBC5_RGUnorm MTLPixelFormat = 142
	// MTLPixelFormatBC6H_RGBFloat: Compressed format with four floating-point components.
	MTLPixelFormatBC6H_RGBFloat MTLPixelFormat = 150
	// MTLPixelFormatBC6H_RGBUfloat: Compressed format with four unsigned floating-point components.
	MTLPixelFormatBC6H_RGBUfloat MTLPixelFormat = 151
	// MTLPixelFormatBC7_RGBAUnorm: Compressed format with four normalized unsigned integer components.
	MTLPixelFormatBC7_RGBAUnorm MTLPixelFormat = 152
	// MTLPixelFormatBC7_RGBAUnorm_sRGB: Compressed format with four normalized unsigned integer components, with conversion between sRGB and linear space.
	MTLPixelFormatBC7_RGBAUnorm_sRGB MTLPixelFormat = 153
	// MTLPixelFormatBGR10A2Unorm: A 32-bit packed pixel format with four normalized unsigned integer components: 10-bit blue, 10-bit green, 10-bit red, and 2-bit alpha.
	MTLPixelFormatBGR10A2Unorm MTLPixelFormat = 94
	// MTLPixelFormatBGR10_XR: A 32-bit extended-range pixel format with three fixed-point components of 10-bit blue, 10-bit green, and 10-bit red.
	MTLPixelFormatBGR10_XR MTLPixelFormat = 554
	// MTLPixelFormatBGR10_XR_sRGB: A 32-bit extended-range pixel format with sRGB conversion and three fixed-point components of 10-bit blue, 10-bit green, and 10-bit red.
	MTLPixelFormatBGR10_XR_sRGB MTLPixelFormat = 555
	// MTLPixelFormatBGR5A1Unorm: Packed 16-bit format with normalized unsigned integer color components: 5 bits each for BGR and 1 for alpha, packed into 16 bits.
	MTLPixelFormatBGR5A1Unorm MTLPixelFormat = 43
	// MTLPixelFormatBGRA10_XR: A 64-bit extended-range pixel format with four fixed-point components of 10-bit blue, 10-bit green, 10-bit red, and 10-bit alpha.
	MTLPixelFormatBGRA10_XR MTLPixelFormat = 552
	// MTLPixelFormatBGRA10_XR_sRGB: A 64-bit extended-range pixel format with sRGB conversion and four fixed-point components of 10-bit blue, 10-bit green, 10-bit red, and 10-bit alpha.
	MTLPixelFormatBGRA10_XR_sRGB MTLPixelFormat = 553
	// MTLPixelFormatBGRA8Unorm: Ordinary format with four 8-bit normalized unsigned integer components in BGRA order.
	MTLPixelFormatBGRA8Unorm MTLPixelFormat = 80
	// MTLPixelFormatBGRA8Unorm_sRGB: Ordinary format with four 8-bit normalized unsigned integer components in BGRA order with conversion between sRGB and linear space.
	MTLPixelFormatBGRA8Unorm_sRGB MTLPixelFormat = 81
	// MTLPixelFormatBGRG422: A pixel format where the red and green components are subsampled horizontally.
	MTLPixelFormatBGRG422 MTLPixelFormat = 241
	// MTLPixelFormatDepth16Unorm: A pixel format for a depth-render target that has a 16-bit normalized, unsigned-integer component.
	MTLPixelFormatDepth16Unorm MTLPixelFormat = 250
	// MTLPixelFormatDepth24Unorm_Stencil8: A 32-bit combined depth and stencil pixel format with a 24-bit normalized unsigned integer for depth and an 8-bit unsigned integer for stencil.
	MTLPixelFormatDepth24Unorm_Stencil8 MTLPixelFormat = 255
	// MTLPixelFormatDepth32Float: A pixel format with one 32-bit floating-point component, used for a depth render target.
	MTLPixelFormatDepth32Float MTLPixelFormat = 252
	// MTLPixelFormatDepth32Float_Stencil8: A 40-bit combined depth and stencil pixel format with a 32-bit floating-point value for depth and an 8-bit unsigned integer for stencil.
	MTLPixelFormatDepth32Float_Stencil8 MTLPixelFormat = 260
	// MTLPixelFormatEAC_R11Snorm: Compressed format using EAC compression with one normalized signed integer component.
	MTLPixelFormatEAC_R11Snorm MTLPixelFormat = 172
	// MTLPixelFormatEAC_R11Unorm: Compressed format using EAC compression with one normalized unsigned integer component.
	MTLPixelFormatEAC_R11Unorm MTLPixelFormat = 170
	// MTLPixelFormatEAC_RG11Snorm: Compressed format using EAC compression with two normalized signed integer components.
	MTLPixelFormatEAC_RG11Snorm MTLPixelFormat = 176
	// MTLPixelFormatEAC_RG11Unorm: Compressed format using EAC compression with two normalized unsigned integer components.
	MTLPixelFormatEAC_RG11Unorm MTLPixelFormat = 174
	// MTLPixelFormatEAC_RGBA8: Compressed format using EAC compression with four 8-bit components.
	MTLPixelFormatEAC_RGBA8 MTLPixelFormat = 178
	// MTLPixelFormatEAC_RGBA8_sRGB: Compressed format using EAC compression with four 8-bit components with conversion between sRGB and linear space.
	MTLPixelFormatEAC_RGBA8_sRGB MTLPixelFormat = 179
	// MTLPixelFormatETC2_RGB8: Compressed format using ETC2 compression with three 8-bit components.
	MTLPixelFormatETC2_RGB8 MTLPixelFormat = 180
	// MTLPixelFormatETC2_RGB8A1: Compressed format using ETC2 compression with four 8-bit components.
	MTLPixelFormatETC2_RGB8A1 MTLPixelFormat = 182
	// MTLPixelFormatETC2_RGB8A1_sRGB: Compressed format using ETC2 compression with four 8-bit components with conversion between sRGB and linear space.
	MTLPixelFormatETC2_RGB8A1_sRGB MTLPixelFormat = 183
	// MTLPixelFormatETC2_RGB8_sRGB: Compressed format using ETC2 compression with three 8-bit components with conversion between sRGB and linear space.
	MTLPixelFormatETC2_RGB8_sRGB MTLPixelFormat = 181
	// MTLPixelFormatGBGR422: A pixel format where the red and green components are subsampled horizontally.
	MTLPixelFormatGBGR422 MTLPixelFormat = 240
	// MTLPixelFormatInvalid: The default value of the pixel format for the [MTLRenderPipelineState].
	MTLPixelFormatInvalid MTLPixelFormat = 0
	// MTLPixelFormatR16Float: Ordinary format with one 16-bit floating-point component.
	MTLPixelFormatR16Float MTLPixelFormat = 25
	// MTLPixelFormatR16Sint: Ordinary format with one 16-bit signed integer component.
	MTLPixelFormatR16Sint MTLPixelFormat = 24
	// MTLPixelFormatR16Snorm: Ordinary format with one 16-bit normalized signed integer component.
	MTLPixelFormatR16Snorm MTLPixelFormat = 22
	// MTLPixelFormatR16Uint: Ordinary format with one 16-bit unsigned integer component.
	MTLPixelFormatR16Uint MTLPixelFormat = 23
	// MTLPixelFormatR16Unorm: Ordinary format with one 16-bit normalized unsigned integer component.
	MTLPixelFormatR16Unorm MTLPixelFormat = 20
	// MTLPixelFormatR32Float: Ordinary format with one 32-bit floating-point component.
	MTLPixelFormatR32Float MTLPixelFormat = 55
	// MTLPixelFormatR32Sint: Ordinary format with one 32-bit signed integer component.
	MTLPixelFormatR32Sint MTLPixelFormat = 54
	// MTLPixelFormatR32Uint: Ordinary format with one 32-bit unsigned integer component.
	MTLPixelFormatR32Uint MTLPixelFormat = 53
	// MTLPixelFormatR8Sint: Ordinary format with one 8-bit signed integer component.
	MTLPixelFormatR8Sint MTLPixelFormat = 14
	// MTLPixelFormatR8Snorm: Ordinary format with one 8-bit normalized signed integer component.
	MTLPixelFormatR8Snorm MTLPixelFormat = 12
	// MTLPixelFormatR8Uint: Ordinary format with one 8-bit unsigned integer component.
	MTLPixelFormatR8Uint MTLPixelFormat = 13
	// MTLPixelFormatR8Unorm: Ordinary format with one 8-bit normalized unsigned integer component.
	MTLPixelFormatR8Unorm MTLPixelFormat = 10
	// MTLPixelFormatR8Unorm_sRGB: Ordinary format with one 8-bit normalized unsigned integer component with conversion between sRGB and linear space.
	MTLPixelFormatR8Unorm_sRGB MTLPixelFormat = 11
	// MTLPixelFormatRG11B10Float: 32-bit format with floating-point color components, 11 bits each for red and green and 10 bits for blue.
	MTLPixelFormatRG11B10Float MTLPixelFormat = 92
	// MTLPixelFormatRG16Float: Ordinary format with two 16-bit floating-point components.
	MTLPixelFormatRG16Float MTLPixelFormat = 65
	// MTLPixelFormatRG16Sint: Ordinary format with two 16-bit signed integer components.
	MTLPixelFormatRG16Sint MTLPixelFormat = 64
	// MTLPixelFormatRG16Snorm: Ordinary format with two 16-bit normalized signed integer components.
	MTLPixelFormatRG16Snorm MTLPixelFormat = 62
	// MTLPixelFormatRG16Uint: Ordinary format with two 16-bit unsigned integer components.
	MTLPixelFormatRG16Uint MTLPixelFormat = 63
	// MTLPixelFormatRG16Unorm: Ordinary format with two 16-bit normalized unsigned integer components.
	MTLPixelFormatRG16Unorm MTLPixelFormat = 60
	// MTLPixelFormatRG32Float: Ordinary format with two 32-bit floating-point components.
	MTLPixelFormatRG32Float MTLPixelFormat = 105
	// MTLPixelFormatRG32Sint: Ordinary format with two 32-bit signed integer components.
	MTLPixelFormatRG32Sint MTLPixelFormat = 104
	// MTLPixelFormatRG32Uint: Ordinary format with two 32-bit unsigned integer components.
	MTLPixelFormatRG32Uint MTLPixelFormat = 103
	// MTLPixelFormatRG8Sint: Ordinary format with two 8-bit signed integer components.
	MTLPixelFormatRG8Sint MTLPixelFormat = 34
	// MTLPixelFormatRG8Snorm: Ordinary format with two 8-bit normalized signed integer components.
	MTLPixelFormatRG8Snorm MTLPixelFormat = 32
	// MTLPixelFormatRG8Uint: Ordinary format with two 8-bit unsigned integer components.
	MTLPixelFormatRG8Uint MTLPixelFormat = 33
	// MTLPixelFormatRG8Unorm: Ordinary format with two 8-bit normalized unsigned integer components.
	MTLPixelFormatRG8Unorm MTLPixelFormat = 30
	// MTLPixelFormatRG8Unorm_sRGB: Ordinary format with two 8-bit normalized unsigned integer components with conversion between sRGB and linear space.
	MTLPixelFormatRG8Unorm_sRGB MTLPixelFormat = 31
	// MTLPixelFormatRGB10A2Uint: A 32-bit packed pixel format with four unsigned integer components: 10-bit red, 10-bit green, 10-bit blue, and 2-bit alpha.
	MTLPixelFormatRGB10A2Uint MTLPixelFormat = 91
	// MTLPixelFormatRGB10A2Unorm: A 32-bit packed pixel format with four normalized unsigned integer components: 10-bit red, 10-bit green, 10-bit blue, and 2-bit alpha.
	MTLPixelFormatRGB10A2Unorm MTLPixelFormat = 90
	// MTLPixelFormatRGB9E5Float: Packed 32-bit format with floating-point color components: 9 bits each for RGB and 5 bits for an exponent shared by RGB, packed into 32 bits.
	MTLPixelFormatRGB9E5Float MTLPixelFormat = 93
	// MTLPixelFormatRGBA16Float: Ordinary format with four 16-bit floating-point components in RGBA order.
	MTLPixelFormatRGBA16Float MTLPixelFormat = 115
	// MTLPixelFormatRGBA16Sint: Ordinary format with four 16-bit signed integer components in RGBA order.
	MTLPixelFormatRGBA16Sint MTLPixelFormat = 114
	// MTLPixelFormatRGBA16Snorm: Ordinary format with four 16-bit normalized signed integer components in RGBA order.
	MTLPixelFormatRGBA16Snorm MTLPixelFormat = 112
	// MTLPixelFormatRGBA16Uint: Ordinary format with four 16-bit unsigned integer components in RGBA order.
	MTLPixelFormatRGBA16Uint MTLPixelFormat = 113
	// MTLPixelFormatRGBA16Unorm: Ordinary format with four 16-bit normalized unsigned integer components in RGBA order.
	MTLPixelFormatRGBA16Unorm MTLPixelFormat = 110
	// MTLPixelFormatRGBA32Float: Ordinary format with four 32-bit floating-point components in RGBA order.
	MTLPixelFormatRGBA32Float MTLPixelFormat = 125
	// MTLPixelFormatRGBA32Sint: Ordinary format with four 32-bit signed integer components in RGBA order.
	MTLPixelFormatRGBA32Sint MTLPixelFormat = 124
	// MTLPixelFormatRGBA32Uint: Ordinary format with four 32-bit unsigned integer components in RGBA order.
	MTLPixelFormatRGBA32Uint MTLPixelFormat = 123
	// MTLPixelFormatRGBA8Sint: Ordinary format with four 8-bit signed integer components in RGBA order.
	MTLPixelFormatRGBA8Sint MTLPixelFormat = 74
	// MTLPixelFormatRGBA8Snorm: Ordinary format with four 8-bit normalized signed integer components in RGBA order.
	MTLPixelFormatRGBA8Snorm MTLPixelFormat = 72
	// MTLPixelFormatRGBA8Uint: Ordinary format with four 8-bit unsigned integer components in RGBA order.
	MTLPixelFormatRGBA8Uint MTLPixelFormat = 73
	// MTLPixelFormatRGBA8Unorm: Ordinary format with four 8-bit normalized unsigned integer components in RGBA order.
	MTLPixelFormatRGBA8Unorm MTLPixelFormat = 70
	// MTLPixelFormatRGBA8Unorm_sRGB: Ordinary format with four 8-bit normalized unsigned integer components in RGBA order with conversion between sRGB and linear space.
	MTLPixelFormatRGBA8Unorm_sRGB MTLPixelFormat = 71
	// MTLPixelFormatStencil8: A pixel format with an 8-bit unsigned integer component, used for a stencil render target.
	MTLPixelFormatStencil8 MTLPixelFormat = 253
	// MTLPixelFormatUnspecialized: # Discussion
	MTLPixelFormatUnspecialized MTLPixelFormat = 263
	// MTLPixelFormatX24_Stencil8: A stencil pixel format used to read the stencil value from a texture with a combined 24-bit depth and 8-bit stencil value.
	MTLPixelFormatX24_Stencil8 MTLPixelFormat = 262
	// MTLPixelFormatX32_Stencil8: A stencil pixel format used to read the stencil value from a texture with a combined 32-bit depth and 8-bit stencil value.
	MTLPixelFormatX32_Stencil8 MTLPixelFormat = 261
	// Deprecated.
	MTLPixelFormatPVRTC_RGBA_2BPP MTLPixelFormat = 164
	// Deprecated.
	MTLPixelFormatPVRTC_RGBA_2BPP_sRGB MTLPixelFormat = 165
	// Deprecated.
	MTLPixelFormatPVRTC_RGBA_4BPP MTLPixelFormat = 166
	// Deprecated.
	MTLPixelFormatPVRTC_RGBA_4BPP_sRGB MTLPixelFormat = 167
	// Deprecated.
	MTLPixelFormatPVRTC_RGB_2BPP MTLPixelFormat = 160
	// Deprecated.
	MTLPixelFormatPVRTC_RGB_2BPP_sRGB MTLPixelFormat = 161
	// Deprecated.
	MTLPixelFormatPVRTC_RGB_4BPP MTLPixelFormat = 162
	// Deprecated.
	MTLPixelFormatPVRTC_RGB_4BPP_sRGB MTLPixelFormat = 163
)

func (e MTLPixelFormat) String() string {
	switch e {
	case MTLPixelFormatA1BGR5Unorm:
		return "MTLPixelFormatA1BGR5Unorm"
	case MTLPixelFormatA8Unorm:
		return "MTLPixelFormatA8Unorm"
	case MTLPixelFormatABGR4Unorm:
		return "MTLPixelFormatABGR4Unorm"
	case MTLPixelFormatASTC_10x10_HDR:
		return "MTLPixelFormatASTC_10x10_HDR"
	case MTLPixelFormatASTC_10x10_LDR:
		return "MTLPixelFormatASTC_10x10_LDR"
	case MTLPixelFormatASTC_10x10_sRGB:
		return "MTLPixelFormatASTC_10x10_sRGB"
	case MTLPixelFormatASTC_10x5_HDR:
		return "MTLPixelFormatASTC_10x5_HDR"
	case MTLPixelFormatASTC_10x5_LDR:
		return "MTLPixelFormatASTC_10x5_LDR"
	case MTLPixelFormatASTC_10x5_sRGB:
		return "MTLPixelFormatASTC_10x5_sRGB"
	case MTLPixelFormatASTC_10x6_HDR:
		return "MTLPixelFormatASTC_10x6_HDR"
	case MTLPixelFormatASTC_10x6_LDR:
		return "MTLPixelFormatASTC_10x6_LDR"
	case MTLPixelFormatASTC_10x6_sRGB:
		return "MTLPixelFormatASTC_10x6_sRGB"
	case MTLPixelFormatASTC_10x8_HDR:
		return "MTLPixelFormatASTC_10x8_HDR"
	case MTLPixelFormatASTC_10x8_LDR:
		return "MTLPixelFormatASTC_10x8_LDR"
	case MTLPixelFormatASTC_10x8_sRGB:
		return "MTLPixelFormatASTC_10x8_sRGB"
	case MTLPixelFormatASTC_12x10_HDR:
		return "MTLPixelFormatASTC_12x10_HDR"
	case MTLPixelFormatASTC_12x10_LDR:
		return "MTLPixelFormatASTC_12x10_LDR"
	case MTLPixelFormatASTC_12x10_sRGB:
		return "MTLPixelFormatASTC_12x10_sRGB"
	case MTLPixelFormatASTC_12x12_HDR:
		return "MTLPixelFormatASTC_12x12_HDR"
	case MTLPixelFormatASTC_12x12_LDR:
		return "MTLPixelFormatASTC_12x12_LDR"
	case MTLPixelFormatASTC_12x12_sRGB:
		return "MTLPixelFormatASTC_12x12_sRGB"
	case MTLPixelFormatASTC_4x4_HDR:
		return "MTLPixelFormatASTC_4x4_HDR"
	case MTLPixelFormatASTC_4x4_LDR:
		return "MTLPixelFormatASTC_4x4_LDR"
	case MTLPixelFormatASTC_4x4_sRGB:
		return "MTLPixelFormatASTC_4x4_sRGB"
	case MTLPixelFormatASTC_5x4_HDR:
		return "MTLPixelFormatASTC_5x4_HDR"
	case MTLPixelFormatASTC_5x4_LDR:
		return "MTLPixelFormatASTC_5x4_LDR"
	case MTLPixelFormatASTC_5x4_sRGB:
		return "MTLPixelFormatASTC_5x4_sRGB"
	case MTLPixelFormatASTC_5x5_HDR:
		return "MTLPixelFormatASTC_5x5_HDR"
	case MTLPixelFormatASTC_5x5_LDR:
		return "MTLPixelFormatASTC_5x5_LDR"
	case MTLPixelFormatASTC_5x5_sRGB:
		return "MTLPixelFormatASTC_5x5_sRGB"
	case MTLPixelFormatASTC_6x5_HDR:
		return "MTLPixelFormatASTC_6x5_HDR"
	case MTLPixelFormatASTC_6x5_LDR:
		return "MTLPixelFormatASTC_6x5_LDR"
	case MTLPixelFormatASTC_6x5_sRGB:
		return "MTLPixelFormatASTC_6x5_sRGB"
	case MTLPixelFormatASTC_6x6_HDR:
		return "MTLPixelFormatASTC_6x6_HDR"
	case MTLPixelFormatASTC_6x6_LDR:
		return "MTLPixelFormatASTC_6x6_LDR"
	case MTLPixelFormatASTC_6x6_sRGB:
		return "MTLPixelFormatASTC_6x6_sRGB"
	case MTLPixelFormatASTC_8x5_HDR:
		return "MTLPixelFormatASTC_8x5_HDR"
	case MTLPixelFormatASTC_8x5_LDR:
		return "MTLPixelFormatASTC_8x5_LDR"
	case MTLPixelFormatASTC_8x5_sRGB:
		return "MTLPixelFormatASTC_8x5_sRGB"
	case MTLPixelFormatASTC_8x6_HDR:
		return "MTLPixelFormatASTC_8x6_HDR"
	case MTLPixelFormatASTC_8x6_LDR:
		return "MTLPixelFormatASTC_8x6_LDR"
	case MTLPixelFormatASTC_8x6_sRGB:
		return "MTLPixelFormatASTC_8x6_sRGB"
	case MTLPixelFormatASTC_8x8_HDR:
		return "MTLPixelFormatASTC_8x8_HDR"
	case MTLPixelFormatASTC_8x8_LDR:
		return "MTLPixelFormatASTC_8x8_LDR"
	case MTLPixelFormatASTC_8x8_sRGB:
		return "MTLPixelFormatASTC_8x8_sRGB"
	case MTLPixelFormatB5G6R5Unorm:
		return "MTLPixelFormatB5G6R5Unorm"
	case MTLPixelFormatBC1_RGBA:
		return "MTLPixelFormatBC1_RGBA"
	case MTLPixelFormatBC1_RGBA_sRGB:
		return "MTLPixelFormatBC1_RGBA_sRGB"
	case MTLPixelFormatBC2_RGBA:
		return "MTLPixelFormatBC2_RGBA"
	case MTLPixelFormatBC2_RGBA_sRGB:
		return "MTLPixelFormatBC2_RGBA_sRGB"
	case MTLPixelFormatBC3_RGBA:
		return "MTLPixelFormatBC3_RGBA"
	case MTLPixelFormatBC3_RGBA_sRGB:
		return "MTLPixelFormatBC3_RGBA_sRGB"
	case MTLPixelFormatBC4_RSnorm:
		return "MTLPixelFormatBC4_RSnorm"
	case MTLPixelFormatBC4_RUnorm:
		return "MTLPixelFormatBC4_RUnorm"
	case MTLPixelFormatBC5_RGSnorm:
		return "MTLPixelFormatBC5_RGSnorm"
	case MTLPixelFormatBC5_RGUnorm:
		return "MTLPixelFormatBC5_RGUnorm"
	case MTLPixelFormatBC6H_RGBFloat:
		return "MTLPixelFormatBC6H_RGBFloat"
	case MTLPixelFormatBC6H_RGBUfloat:
		return "MTLPixelFormatBC6H_RGBUfloat"
	case MTLPixelFormatBC7_RGBAUnorm:
		return "MTLPixelFormatBC7_RGBAUnorm"
	case MTLPixelFormatBC7_RGBAUnorm_sRGB:
		return "MTLPixelFormatBC7_RGBAUnorm_sRGB"
	case MTLPixelFormatBGR10A2Unorm:
		return "MTLPixelFormatBGR10A2Unorm"
	case MTLPixelFormatBGR10_XR:
		return "MTLPixelFormatBGR10_XR"
	case MTLPixelFormatBGR10_XR_sRGB:
		return "MTLPixelFormatBGR10_XR_sRGB"
	case MTLPixelFormatBGR5A1Unorm:
		return "MTLPixelFormatBGR5A1Unorm"
	case MTLPixelFormatBGRA10_XR:
		return "MTLPixelFormatBGRA10_XR"
	case MTLPixelFormatBGRA10_XR_sRGB:
		return "MTLPixelFormatBGRA10_XR_sRGB"
	case MTLPixelFormatBGRA8Unorm:
		return "MTLPixelFormatBGRA8Unorm"
	case MTLPixelFormatBGRA8Unorm_sRGB:
		return "MTLPixelFormatBGRA8Unorm_sRGB"
	case MTLPixelFormatBGRG422:
		return "MTLPixelFormatBGRG422"
	case MTLPixelFormatDepth16Unorm:
		return "MTLPixelFormatDepth16Unorm"
	case MTLPixelFormatDepth24Unorm_Stencil8:
		return "MTLPixelFormatDepth24Unorm_Stencil8"
	case MTLPixelFormatDepth32Float:
		return "MTLPixelFormatDepth32Float"
	case MTLPixelFormatDepth32Float_Stencil8:
		return "MTLPixelFormatDepth32Float_Stencil8"
	case MTLPixelFormatEAC_R11Snorm:
		return "MTLPixelFormatEAC_R11Snorm"
	case MTLPixelFormatEAC_R11Unorm:
		return "MTLPixelFormatEAC_R11Unorm"
	case MTLPixelFormatEAC_RG11Snorm:
		return "MTLPixelFormatEAC_RG11Snorm"
	case MTLPixelFormatEAC_RG11Unorm:
		return "MTLPixelFormatEAC_RG11Unorm"
	case MTLPixelFormatEAC_RGBA8:
		return "MTLPixelFormatEAC_RGBA8"
	case MTLPixelFormatEAC_RGBA8_sRGB:
		return "MTLPixelFormatEAC_RGBA8_sRGB"
	case MTLPixelFormatETC2_RGB8:
		return "MTLPixelFormatETC2_RGB8"
	case MTLPixelFormatETC2_RGB8A1:
		return "MTLPixelFormatETC2_RGB8A1"
	case MTLPixelFormatETC2_RGB8A1_sRGB:
		return "MTLPixelFormatETC2_RGB8A1_sRGB"
	case MTLPixelFormatETC2_RGB8_sRGB:
		return "MTLPixelFormatETC2_RGB8_sRGB"
	case MTLPixelFormatGBGR422:
		return "MTLPixelFormatGBGR422"
	case MTLPixelFormatInvalid:
		return "MTLPixelFormatInvalid"
	case MTLPixelFormatR16Float:
		return "MTLPixelFormatR16Float"
	case MTLPixelFormatR16Sint:
		return "MTLPixelFormatR16Sint"
	case MTLPixelFormatR16Snorm:
		return "MTLPixelFormatR16Snorm"
	case MTLPixelFormatR16Uint:
		return "MTLPixelFormatR16Uint"
	case MTLPixelFormatR16Unorm:
		return "MTLPixelFormatR16Unorm"
	case MTLPixelFormatR32Float:
		return "MTLPixelFormatR32Float"
	case MTLPixelFormatR32Sint:
		return "MTLPixelFormatR32Sint"
	case MTLPixelFormatR32Uint:
		return "MTLPixelFormatR32Uint"
	case MTLPixelFormatR8Sint:
		return "MTLPixelFormatR8Sint"
	case MTLPixelFormatR8Snorm:
		return "MTLPixelFormatR8Snorm"
	case MTLPixelFormatR8Uint:
		return "MTLPixelFormatR8Uint"
	case MTLPixelFormatR8Unorm:
		return "MTLPixelFormatR8Unorm"
	case MTLPixelFormatR8Unorm_sRGB:
		return "MTLPixelFormatR8Unorm_sRGB"
	case MTLPixelFormatRG11B10Float:
		return "MTLPixelFormatRG11B10Float"
	case MTLPixelFormatRG16Float:
		return "MTLPixelFormatRG16Float"
	case MTLPixelFormatRG16Sint:
		return "MTLPixelFormatRG16Sint"
	case MTLPixelFormatRG16Snorm:
		return "MTLPixelFormatRG16Snorm"
	case MTLPixelFormatRG16Uint:
		return "MTLPixelFormatRG16Uint"
	case MTLPixelFormatRG16Unorm:
		return "MTLPixelFormatRG16Unorm"
	case MTLPixelFormatRG32Float:
		return "MTLPixelFormatRG32Float"
	case MTLPixelFormatRG32Sint:
		return "MTLPixelFormatRG32Sint"
	case MTLPixelFormatRG32Uint:
		return "MTLPixelFormatRG32Uint"
	case MTLPixelFormatRG8Sint:
		return "MTLPixelFormatRG8Sint"
	case MTLPixelFormatRG8Snorm:
		return "MTLPixelFormatRG8Snorm"
	case MTLPixelFormatRG8Uint:
		return "MTLPixelFormatRG8Uint"
	case MTLPixelFormatRG8Unorm:
		return "MTLPixelFormatRG8Unorm"
	case MTLPixelFormatRG8Unorm_sRGB:
		return "MTLPixelFormatRG8Unorm_sRGB"
	case MTLPixelFormatRGB10A2Uint:
		return "MTLPixelFormatRGB10A2Uint"
	case MTLPixelFormatRGB10A2Unorm:
		return "MTLPixelFormatRGB10A2Unorm"
	case MTLPixelFormatRGB9E5Float:
		return "MTLPixelFormatRGB9E5Float"
	case MTLPixelFormatRGBA16Float:
		return "MTLPixelFormatRGBA16Float"
	case MTLPixelFormatRGBA16Sint:
		return "MTLPixelFormatRGBA16Sint"
	case MTLPixelFormatRGBA16Snorm:
		return "MTLPixelFormatRGBA16Snorm"
	case MTLPixelFormatRGBA16Uint:
		return "MTLPixelFormatRGBA16Uint"
	case MTLPixelFormatRGBA16Unorm:
		return "MTLPixelFormatRGBA16Unorm"
	case MTLPixelFormatRGBA32Float:
		return "MTLPixelFormatRGBA32Float"
	case MTLPixelFormatRGBA32Sint:
		return "MTLPixelFormatRGBA32Sint"
	case MTLPixelFormatRGBA32Uint:
		return "MTLPixelFormatRGBA32Uint"
	case MTLPixelFormatRGBA8Sint:
		return "MTLPixelFormatRGBA8Sint"
	case MTLPixelFormatRGBA8Snorm:
		return "MTLPixelFormatRGBA8Snorm"
	case MTLPixelFormatRGBA8Uint:
		return "MTLPixelFormatRGBA8Uint"
	case MTLPixelFormatRGBA8Unorm:
		return "MTLPixelFormatRGBA8Unorm"
	case MTLPixelFormatRGBA8Unorm_sRGB:
		return "MTLPixelFormatRGBA8Unorm_sRGB"
	case MTLPixelFormatStencil8:
		return "MTLPixelFormatStencil8"
	case MTLPixelFormatUnspecialized:
		return "MTLPixelFormatUnspecialized"
	case MTLPixelFormatX24_Stencil8:
		return "MTLPixelFormatX24_Stencil8"
	case MTLPixelFormatX32_Stencil8:
		return "MTLPixelFormatX32_Stencil8"
	case MTLPixelFormatPVRTC_RGBA_2BPP:
		return "MTLPixelFormatPVRTC_RGBA_2BPP"
	case MTLPixelFormatPVRTC_RGBA_2BPP_sRGB:
		return "MTLPixelFormatPVRTC_RGBA_2BPP_sRGB"
	case MTLPixelFormatPVRTC_RGBA_4BPP:
		return "MTLPixelFormatPVRTC_RGBA_4BPP"
	case MTLPixelFormatPVRTC_RGBA_4BPP_sRGB:
		return "MTLPixelFormatPVRTC_RGBA_4BPP_sRGB"
	case MTLPixelFormatPVRTC_RGB_2BPP:
		return "MTLPixelFormatPVRTC_RGB_2BPP"
	case MTLPixelFormatPVRTC_RGB_2BPP_sRGB:
		return "MTLPixelFormatPVRTC_RGB_2BPP_sRGB"
	case MTLPixelFormatPVRTC_RGB_4BPP:
		return "MTLPixelFormatPVRTC_RGB_4BPP"
	case MTLPixelFormatPVRTC_RGB_4BPP_sRGB:
		return "MTLPixelFormatPVRTC_RGB_4BPP_sRGB"
	default:
		return fmt.Sprintf("MTLPixelFormat(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLPrimitiveTopologyClass
type MTLPrimitiveTopologyClass uint

const (
	// MTLPrimitiveTopologyClassLine: A line primitive.
	MTLPrimitiveTopologyClassLine MTLPrimitiveTopologyClass = 2
	// MTLPrimitiveTopologyClassPoint: A point primitive.
	MTLPrimitiveTopologyClassPoint MTLPrimitiveTopologyClass = 1
	// MTLPrimitiveTopologyClassTriangle: A triangle primitive.
	MTLPrimitiveTopologyClassTriangle MTLPrimitiveTopologyClass = 3
	// MTLPrimitiveTopologyClassUnspecified: An unspecified primitive.
	MTLPrimitiveTopologyClassUnspecified MTLPrimitiveTopologyClass = 0
)

func (e MTLPrimitiveTopologyClass) String() string {
	switch e {
	case MTLPrimitiveTopologyClassLine:
		return "MTLPrimitiveTopologyClassLine"
	case MTLPrimitiveTopologyClassPoint:
		return "MTLPrimitiveTopologyClassPoint"
	case MTLPrimitiveTopologyClassTriangle:
		return "MTLPrimitiveTopologyClassTriangle"
	case MTLPrimitiveTopologyClassUnspecified:
		return "MTLPrimitiveTopologyClassUnspecified"
	default:
		return fmt.Sprintf("MTLPrimitiveTopologyClass(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLPrimitiveType
type MTLPrimitiveType uint

const (
	// MTLPrimitiveTypeLine: Rasterize a line between each separate pair of vertices, resulting in a series of unconnected lines.
	MTLPrimitiveTypeLine MTLPrimitiveType = 1
	// MTLPrimitiveTypeLineStrip: Rasterize a line between each pair of adjacent vertices, resulting in a series of connected lines (also called a polyline).
	MTLPrimitiveTypeLineStrip MTLPrimitiveType = 2
	// MTLPrimitiveTypePoint: Rasterize a point at each vertex.
	MTLPrimitiveTypePoint MTLPrimitiveType = 0
	// MTLPrimitiveTypeTriangle: For every separate set of three vertices, rasterize a triangle.
	MTLPrimitiveTypeTriangle MTLPrimitiveType = 3
	// MTLPrimitiveTypeTriangleStrip: For every three adjacent vertices, rasterize a triangle.
	MTLPrimitiveTypeTriangleStrip MTLPrimitiveType = 4
)

func (e MTLPrimitiveType) String() string {
	switch e {
	case MTLPrimitiveTypeLine:
		return "MTLPrimitiveTypeLine"
	case MTLPrimitiveTypeLineStrip:
		return "MTLPrimitiveTypeLineStrip"
	case MTLPrimitiveTypePoint:
		return "MTLPrimitiveTypePoint"
	case MTLPrimitiveTypeTriangle:
		return "MTLPrimitiveTypeTriangle"
	case MTLPrimitiveTypeTriangleStrip:
		return "MTLPrimitiveTypeTriangleStrip"
	default:
		return fmt.Sprintf("MTLPrimitiveType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLPurgeableState
type MTLPurgeableState uint

const (
	// MTLPurgeableStateEmpty: A state that indicates to the system that it needs to consider the contents of a resource as invalid, typically because you’re discarding it.
	MTLPurgeableStateEmpty MTLPurgeableState = 4
	// MTLPurgeableStateKeepCurrent: The current state is queried but doesn’t change.
	MTLPurgeableStateKeepCurrent MTLPurgeableState = 1
	// MTLPurgeableStateNonVolatile: The contents of the resource aren’t allowed to be discarded.
	MTLPurgeableStateNonVolatile MTLPurgeableState = 2
	// MTLPurgeableStateVolatile: The system is allowed to discard the resource to free up memory.
	MTLPurgeableStateVolatile MTLPurgeableState = 3
)

func (e MTLPurgeableState) String() string {
	switch e {
	case MTLPurgeableStateEmpty:
		return "MTLPurgeableStateEmpty"
	case MTLPurgeableStateKeepCurrent:
		return "MTLPurgeableStateKeepCurrent"
	case MTLPurgeableStateNonVolatile:
		return "MTLPurgeableStateNonVolatile"
	case MTLPurgeableStateVolatile:
		return "MTLPurgeableStateVolatile"
	default:
		return fmt.Sprintf("MTLPurgeableState(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLReadWriteTextureTier
type MTLReadWriteTextureTier uint

const (
	// MTLReadWriteTextureTier1: Indicates the system supports tier 1 read-write textures.
	MTLReadWriteTextureTier1 MTLReadWriteTextureTier = 1
	// MTLReadWriteTextureTier2: Indicates the system supports tier 2 read-write textures.
	MTLReadWriteTextureTier2 MTLReadWriteTextureTier = 2
	// MTLReadWriteTextureTierNone: Indicates the system doesn’t support read-write textures.
	MTLReadWriteTextureTierNone MTLReadWriteTextureTier = 0
)

func (e MTLReadWriteTextureTier) String() string {
	switch e {
	case MTLReadWriteTextureTier1:
		return "MTLReadWriteTextureTier1"
	case MTLReadWriteTextureTier2:
		return "MTLReadWriteTextureTier2"
	case MTLReadWriteTextureTierNone:
		return "MTLReadWriteTextureTierNone"
	default:
		return fmt.Sprintf("MTLReadWriteTextureTier(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLRenderStages
type MTLRenderStages uint

const (
	// MTLRenderStageFragment: The fragment rendering stage.
	MTLRenderStageFragment MTLRenderStages = 2
	// MTLRenderStageMesh: The mesh rendering stage.
	MTLRenderStageMesh MTLRenderStages = 16
	// MTLRenderStageObject: The object rendering stage.
	MTLRenderStageObject MTLRenderStages = 8
	// MTLRenderStageTile: The tile rendering stage.
	MTLRenderStageTile MTLRenderStages = 4
	// MTLRenderStageVertex: The vertex rendering stage.
	MTLRenderStageVertex MTLRenderStages = 1
)

func (e MTLRenderStages) String() string {
	switch e {
	case MTLRenderStageFragment:
		return "MTLRenderStageFragment"
	case MTLRenderStageMesh:
		return "MTLRenderStageMesh"
	case MTLRenderStageObject:
		return "MTLRenderStageObject"
	case MTLRenderStageTile:
		return "MTLRenderStageTile"
	case MTLRenderStageVertex:
		return "MTLRenderStageVertex"
	default:
		return fmt.Sprintf("MTLRenderStages(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLResourceOptions
type MTLResourceOptions uint

const (
	// MTLResourceCPUCacheModeDefaultCache: The default CPU cache mode for the resource, which guarantees that read and write operations are executed in the expected order.
	MTLResourceCPUCacheModeDefaultCache MTLResourceOptions = 0
	// MTLResourceCPUCacheModeWriteCombined: A write-combined CPU cache mode that is optimized for resources that the CPU writes into, but never reads.
	MTLResourceCPUCacheModeWriteCombined MTLResourceOptions = 1
	// MTLResourceHazardTrackingModeDefault: An option specifying that the default tracking mode should be used.
	MTLResourceHazardTrackingModeDefault MTLResourceOptions = 0
	// MTLResourceHazardTrackingModeTracked: An option that instructs Metal to apply safeguards for a resource at runtime to avoid memory hazards for the applicable commands.
	MTLResourceHazardTrackingModeTracked MTLResourceOptions = 2
	// MTLResourceHazardTrackingModeUntracked: A resource option that instructs Metal to ignore memory hazards for a resource at runtime.
	MTLResourceHazardTrackingModeUntracked MTLResourceOptions = 1
	// MTLResourceOptionCPUCacheModeDefault: This constant was deprecated in iOS 9.0 and macOS 10.11.
	MTLResourceOptionCPUCacheModeDefault MTLResourceOptions = 0
	// MTLResourceStorageModeManaged: The CPU and GPU may maintain separate copies of the resource, and any changes need to be explicitly synchronized.
	MTLResourceStorageModeManaged MTLResourceOptions = 1
	// MTLResourceStorageModeMemoryless: The resource’s contents are only available to the GPU, and only exist temporarily during a render pass.
	MTLResourceStorageModeMemoryless MTLResourceOptions = 3
	// MTLResourceStorageModePrivate: The resource is only available to the GPU.
	MTLResourceStorageModePrivate MTLResourceOptions = 2
	// MTLResourceStorageModeShared: The CPU and GPU share access to the resource, allocated in system memory.
	MTLResourceStorageModeShared MTLResourceOptions = 0
	// Deprecated.
	MTLResourceOptionCPUCacheModeWriteCombined MTLResourceOptions = 1
)

func (e MTLResourceOptions) String() string {
	switch e {
	case MTLResourceCPUCacheModeDefaultCache:
		return "MTLResourceCPUCacheModeDefaultCache"
	case MTLResourceCPUCacheModeWriteCombined:
		return "MTLResourceCPUCacheModeWriteCombined"
	case MTLResourceHazardTrackingModeTracked:
		return "MTLResourceHazardTrackingModeTracked"
	case MTLResourceStorageModeMemoryless:
		return "MTLResourceStorageModeMemoryless"
	default:
		return fmt.Sprintf("MTLResourceOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLResourceUsage
type MTLResourceUsage uint

const (
	// MTLResourceUsageRead: An option that enables reading from the resource.
	MTLResourceUsageRead MTLResourceUsage = 1
	// MTLResourceUsageWrite: An option that enables writing to the resource.
	MTLResourceUsageWrite MTLResourceUsage = 2
	// Deprecated.
	MTLResourceUsageSample MTLResourceUsage = 4
)

func (e MTLResourceUsage) String() string {
	switch e {
	case MTLResourceUsageRead:
		return "MTLResourceUsageRead"
	case MTLResourceUsageWrite:
		return "MTLResourceUsageWrite"
	case MTLResourceUsageSample:
		return "MTLResourceUsageSample"
	default:
		return fmt.Sprintf("MTLResourceUsage(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLSamplerAddressMode
type MTLSamplerAddressMode uint

const (
	// MTLSamplerAddressModeClampToBorderColor: Out-of-range texture coordinates return the value specified by the  property.
	MTLSamplerAddressModeClampToBorderColor MTLSamplerAddressMode = 5
	// MTLSamplerAddressModeClampToEdge: Texture coordinates are clamped between  and , inclusive.
	MTLSamplerAddressModeClampToEdge MTLSamplerAddressMode = 0
	// MTLSamplerAddressModeClampToZero: Out-of-range texture coordinates return transparent zero  for images with an alpha channel and return opaque zero  for images without an alpha channel.
	MTLSamplerAddressModeClampToZero MTLSamplerAddressMode = 4
	// MTLSamplerAddressModeMirrorClampToEdge: Between `-1.0` and `1.0`, the texture coordinates are mirrored across the axis; outside `-1.0` and `1.0`, texture coordinates are clamped.
	MTLSamplerAddressModeMirrorClampToEdge MTLSamplerAddressMode = 1
	// MTLSamplerAddressModeMirrorRepeat: Between `-1.0` and `1.0`, the texture coordinates are mirrored across the axis; outside `-1.0` and `1.0`, the image is repeated.
	MTLSamplerAddressModeMirrorRepeat MTLSamplerAddressMode = 3
	// MTLSamplerAddressModeRepeat: Texture coordinates wrap to the other side of the texture, effectively keeping only the fractional part of the texture coordinate.
	MTLSamplerAddressModeRepeat MTLSamplerAddressMode = 2
)

func (e MTLSamplerAddressMode) String() string {
	switch e {
	case MTLSamplerAddressModeClampToBorderColor:
		return "MTLSamplerAddressModeClampToBorderColor"
	case MTLSamplerAddressModeClampToEdge:
		return "MTLSamplerAddressModeClampToEdge"
	case MTLSamplerAddressModeClampToZero:
		return "MTLSamplerAddressModeClampToZero"
	case MTLSamplerAddressModeMirrorClampToEdge:
		return "MTLSamplerAddressModeMirrorClampToEdge"
	case MTLSamplerAddressModeMirrorRepeat:
		return "MTLSamplerAddressModeMirrorRepeat"
	case MTLSamplerAddressModeRepeat:
		return "MTLSamplerAddressModeRepeat"
	default:
		return fmt.Sprintf("MTLSamplerAddressMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLSamplerBorderColor
type MTLSamplerBorderColor uint

const (
	// MTLSamplerBorderColorOpaqueBlack: An opaque black color `(0,0,0,1)` for texture values outside the border
	MTLSamplerBorderColorOpaqueBlack MTLSamplerBorderColor = 1
	// MTLSamplerBorderColorOpaqueWhite: An opaque white color `(1,1,1,1)` for texture values outside the border.
	MTLSamplerBorderColorOpaqueWhite MTLSamplerBorderColor = 2
	// MTLSamplerBorderColorTransparentBlack: A transparent black color `(0,0,0,0)` for texture values outside the border.
	MTLSamplerBorderColorTransparentBlack MTLSamplerBorderColor = 0
)

func (e MTLSamplerBorderColor) String() string {
	switch e {
	case MTLSamplerBorderColorOpaqueBlack:
		return "MTLSamplerBorderColorOpaqueBlack"
	case MTLSamplerBorderColorOpaqueWhite:
		return "MTLSamplerBorderColorOpaqueWhite"
	case MTLSamplerBorderColorTransparentBlack:
		return "MTLSamplerBorderColorTransparentBlack"
	default:
		return fmt.Sprintf("MTLSamplerBorderColor(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLSamplerMinMagFilter
type MTLSamplerMinMagFilter uint

const (
	// MTLSamplerMinMagFilterLinear: Select two pixels in each dimension and interpolate linearly between them.
	MTLSamplerMinMagFilterLinear MTLSamplerMinMagFilter = 1
	// MTLSamplerMinMagFilterNearest: Select the single pixel nearest to the sample point.
	MTLSamplerMinMagFilterNearest MTLSamplerMinMagFilter = 0
)

func (e MTLSamplerMinMagFilter) String() string {
	switch e {
	case MTLSamplerMinMagFilterLinear:
		return "MTLSamplerMinMagFilterLinear"
	case MTLSamplerMinMagFilterNearest:
		return "MTLSamplerMinMagFilterNearest"
	default:
		return fmt.Sprintf("MTLSamplerMinMagFilter(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLSamplerMipFilter
type MTLSamplerMipFilter uint

const (
	// MTLSamplerMipFilterLinear: If the filter falls between mipmap levels, both levels are sampled and the results are determined by linear interpolation between levels.
	MTLSamplerMipFilterLinear MTLSamplerMipFilter = 2
	// MTLSamplerMipFilterNearest: The nearest mipmap level is selected.
	MTLSamplerMipFilterNearest MTLSamplerMipFilter = 1
	// MTLSamplerMipFilterNotMipmapped: The texture is sampled from mipmap level `0`, and other mipmap levels are ignored.
	MTLSamplerMipFilterNotMipmapped MTLSamplerMipFilter = 0
)

func (e MTLSamplerMipFilter) String() string {
	switch e {
	case MTLSamplerMipFilterLinear:
		return "MTLSamplerMipFilterLinear"
	case MTLSamplerMipFilterNearest:
		return "MTLSamplerMipFilterNearest"
	case MTLSamplerMipFilterNotMipmapped:
		return "MTLSamplerMipFilterNotMipmapped"
	default:
		return fmt.Sprintf("MTLSamplerMipFilter(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLSamplerReductionMode
type MTLSamplerReductionMode uint

const (
	// MTLSamplerReductionModeMaximum: A reduction mode that finds the maximum contributing sample value by separately evaluating each channel.
	MTLSamplerReductionModeMaximum MTLSamplerReductionMode = 2
	// MTLSamplerReductionModeMinimum: A reduction mode that finds the minimum contributing sample value by separately evaluating each channel.
	MTLSamplerReductionModeMinimum MTLSamplerReductionMode = 1
	// MTLSamplerReductionModeWeightedAverage: A reduction mode that adds together the product of each contributing sample value by its weight.
	MTLSamplerReductionModeWeightedAverage MTLSamplerReductionMode = 0
)

func (e MTLSamplerReductionMode) String() string {
	switch e {
	case MTLSamplerReductionModeMaximum:
		return "MTLSamplerReductionModeMaximum"
	case MTLSamplerReductionModeMinimum:
		return "MTLSamplerReductionModeMinimum"
	case MTLSamplerReductionModeWeightedAverage:
		return "MTLSamplerReductionModeWeightedAverage"
	default:
		return fmt.Sprintf("MTLSamplerReductionMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLShaderValidation
type MTLShaderValidation int

const (
	// MTLShaderValidationDefault: The default value when the property isn’t set.
	MTLShaderValidationDefault MTLShaderValidation = 0
	// MTLShaderValidationDisabled: Disables shader validation.
	MTLShaderValidationDisabled MTLShaderValidation = 2
	// MTLShaderValidationEnabled: Enables shader validation.
	MTLShaderValidationEnabled MTLShaderValidation = 1
)

func (e MTLShaderValidation) String() string {
	switch e {
	case MTLShaderValidationDefault:
		return "MTLShaderValidationDefault"
	case MTLShaderValidationDisabled:
		return "MTLShaderValidationDisabled"
	case MTLShaderValidationEnabled:
		return "MTLShaderValidationEnabled"
	default:
		return fmt.Sprintf("MTLShaderValidation(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLSparsePageSize
type MTLSparsePageSize int

const (
	// MTLSparsePageSize16: Represents a sparse texture’s page size of 16 kilobytes.
	MTLSparsePageSize16 MTLSparsePageSize = 101
	// MTLSparsePageSize256: Represents a sparse texture’s page size of 256 kilobytes.
	MTLSparsePageSize256 MTLSparsePageSize = 103
	// MTLSparsePageSize64: Represents a sparse texture’s page size of 64 kilobytes.
	MTLSparsePageSize64 MTLSparsePageSize = 102
)

func (e MTLSparsePageSize) String() string {
	switch e {
	case MTLSparsePageSize16:
		return "MTLSparsePageSize16"
	case MTLSparsePageSize256:
		return "MTLSparsePageSize256"
	case MTLSparsePageSize64:
		return "MTLSparsePageSize64"
	default:
		return fmt.Sprintf("MTLSparsePageSize(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLSparseTextureMappingMode
type MTLSparseTextureMappingMode uint

const (
	// MTLSparseTextureMappingModeMap: A request to map sparse tiles from the heap to a region in the texture.
	MTLSparseTextureMappingModeMap MTLSparseTextureMappingMode = 0
	// MTLSparseTextureMappingModeUnmap: A request to remove any mappings for a region in the texture.
	MTLSparseTextureMappingModeUnmap MTLSparseTextureMappingMode = 1
)

func (e MTLSparseTextureMappingMode) String() string {
	switch e {
	case MTLSparseTextureMappingModeMap:
		return "MTLSparseTextureMappingModeMap"
	case MTLSparseTextureMappingModeUnmap:
		return "MTLSparseTextureMappingModeUnmap"
	default:
		return fmt.Sprintf("MTLSparseTextureMappingMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLSparseTextureRegionAlignmentMode
type MTLSparseTextureRegionAlignmentMode uint

const (
	// MTLSparseTextureRegionAlignmentModeInward: The tile region ignores partially covered tiles.
	MTLSparseTextureRegionAlignmentModeInward MTLSparseTextureRegionAlignmentMode = 1
	// MTLSparseTextureRegionAlignmentModeOutward: The tile region includes any partially covered tiles.
	MTLSparseTextureRegionAlignmentModeOutward MTLSparseTextureRegionAlignmentMode = 0
)

func (e MTLSparseTextureRegionAlignmentMode) String() string {
	switch e {
	case MTLSparseTextureRegionAlignmentModeInward:
		return "MTLSparseTextureRegionAlignmentModeInward"
	case MTLSparseTextureRegionAlignmentModeOutward:
		return "MTLSparseTextureRegionAlignmentModeOutward"
	default:
		return fmt.Sprintf("MTLSparseTextureRegionAlignmentMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLStages
type MTLStages uint

const (
	// MTLStageAccelerationStructure: Represents all acceleration structure operations.
	MTLStageAccelerationStructure MTLStages = 536870912
	// MTLStageAll: Convenience mask representing all stages of GPU work.
	MTLStageAll MTLStages = 9223372036854775807
	// MTLStageBlit: Represents all blit operations in a pass.
	MTLStageBlit MTLStages = 268435456
	// MTLStageDispatch: Represents all compute dispatches in a compute pass.
	MTLStageDispatch MTLStages = 134217728
	// MTLStageFragment: Represents all fragment shader stage work in a render pass.
	MTLStageFragment MTLStages = 2
	// MTLStageMachineLearning: Represents all machine learning network dispatch operations.
	MTLStageMachineLearning MTLStages = 1073741824
	// MTLStageMesh: Represents all mesh shader stage work work in a render pass.
	MTLStageMesh MTLStages = 16
	// MTLStageObject: Represents all object shader stage work in a render pass.
	MTLStageObject MTLStages = 8
	// MTLStageResourceState: Represents all sparse and placement sparse resource mapping updates.
	MTLStageResourceState MTLStages = 67108864
	// MTLStageTile: Represents all tile shading stage work in a render pass.
	MTLStageTile MTLStages = 4
	// MTLStageVertex: Represents all vertex shader stage work in a render pass.
	MTLStageVertex MTLStages = 1
)

func (e MTLStages) String() string {
	switch e {
	case MTLStageAccelerationStructure:
		return "MTLStageAccelerationStructure"
	case MTLStageAll:
		return "MTLStageAll"
	case MTLStageBlit:
		return "MTLStageBlit"
	case MTLStageDispatch:
		return "MTLStageDispatch"
	case MTLStageFragment:
		return "MTLStageFragment"
	case MTLStageMachineLearning:
		return "MTLStageMachineLearning"
	case MTLStageMesh:
		return "MTLStageMesh"
	case MTLStageObject:
		return "MTLStageObject"
	case MTLStageResourceState:
		return "MTLStageResourceState"
	case MTLStageTile:
		return "MTLStageTile"
	case MTLStageVertex:
		return "MTLStageVertex"
	default:
		return fmt.Sprintf("MTLStages(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLStencilOperation
type MTLStencilOperation uint

const (
	// MTLStencilOperationDecrementClamp: A stencil operation that decreases a nonzero stencil value by one.
	MTLStencilOperationDecrementClamp MTLStencilOperation = 4
	// MTLStencilOperationDecrementWrap: A stencil operation that decreases a nonzero stencil value by one, or when it’s zero, resets it to the maximum representable value.
	MTLStencilOperationDecrementWrap MTLStencilOperation = 7
	// MTLStencilOperationIncrementClamp: A stencil operation that increases a stencil value by one, but only when the current value isn’t the maximum representable value.
	MTLStencilOperationIncrementClamp MTLStencilOperation = 3
	// MTLStencilOperationIncrementWrap: A stencil operation that decreases a nonzero stencil value by one, or when it’s the maximum representable value, resets it to zero.
	MTLStencilOperationIncrementWrap MTLStencilOperation = 6
	// MTLStencilOperationInvert: A stencil operation that applies a logical bitwise NOT to a stencil value.
	MTLStencilOperationInvert MTLStencilOperation = 5
	// MTLStencilOperationKeep: A stencil operation that doesn’t modify a stencil value.
	MTLStencilOperationKeep MTLStencilOperation = 0
	// MTLStencilOperationReplace: A stencil operation that replaces a stencil value with a reference value.
	MTLStencilOperationReplace MTLStencilOperation = 2
	// MTLStencilOperationZero: A stencil operation that sets a stencil value to zero.
	MTLStencilOperationZero MTLStencilOperation = 1
)

func (e MTLStencilOperation) String() string {
	switch e {
	case MTLStencilOperationDecrementClamp:
		return "MTLStencilOperationDecrementClamp"
	case MTLStencilOperationDecrementWrap:
		return "MTLStencilOperationDecrementWrap"
	case MTLStencilOperationIncrementClamp:
		return "MTLStencilOperationIncrementClamp"
	case MTLStencilOperationIncrementWrap:
		return "MTLStencilOperationIncrementWrap"
	case MTLStencilOperationInvert:
		return "MTLStencilOperationInvert"
	case MTLStencilOperationKeep:
		return "MTLStencilOperationKeep"
	case MTLStencilOperationReplace:
		return "MTLStencilOperationReplace"
	case MTLStencilOperationZero:
		return "MTLStencilOperationZero"
	default:
		return fmt.Sprintf("MTLStencilOperation(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLStepFunction
type MTLStepFunction uint

const (
	// MTLStepFunctionConstant: The function fetches attribute data once.
	MTLStepFunctionConstant MTLStepFunction = 0
	// MTLStepFunctionPerInstance: The function fetches data based on the instance index.
	MTLStepFunctionPerInstance MTLStepFunction = 2
	// MTLStepFunctionPerPatch: The post-tessellation function fetches data based on the patch index of the patch.
	MTLStepFunctionPerPatch MTLStepFunction = 3
	// MTLStepFunctionPerPatchControlPoint: The post-tessellation function fetches data based on the control-point indices associated with the patch.
	MTLStepFunctionPerPatchControlPoint MTLStepFunction = 4
	// MTLStepFunctionPerVertex: The vertex function fetches data for every vertex.
	MTLStepFunctionPerVertex MTLStepFunction = 1
	// MTLStepFunctionThreadPositionInGridX: The compute function fetches data based on the thread’s `x` coordinate.
	MTLStepFunctionThreadPositionInGridX MTLStepFunction = 5
	// MTLStepFunctionThreadPositionInGridXIndexed: The compute function fetches data by using the thread’s `x` coordinate to look up a value in the index buffer.
	MTLStepFunctionThreadPositionInGridXIndexed MTLStepFunction = 7
	// MTLStepFunctionThreadPositionInGridY: The compute function fetches data based on the thread’s `y` coordinate.
	MTLStepFunctionThreadPositionInGridY MTLStepFunction = 6
	// MTLStepFunctionThreadPositionInGridYIndexed: The compute function fetches data by using the thread’s `y` coordinate to look up a value in the index buffer.
	MTLStepFunctionThreadPositionInGridYIndexed MTLStepFunction = 8
)

func (e MTLStepFunction) String() string {
	switch e {
	case MTLStepFunctionConstant:
		return "MTLStepFunctionConstant"
	case MTLStepFunctionPerInstance:
		return "MTLStepFunctionPerInstance"
	case MTLStepFunctionPerPatch:
		return "MTLStepFunctionPerPatch"
	case MTLStepFunctionPerPatchControlPoint:
		return "MTLStepFunctionPerPatchControlPoint"
	case MTLStepFunctionPerVertex:
		return "MTLStepFunctionPerVertex"
	case MTLStepFunctionThreadPositionInGridX:
		return "MTLStepFunctionThreadPositionInGridX"
	case MTLStepFunctionThreadPositionInGridXIndexed:
		return "MTLStepFunctionThreadPositionInGridXIndexed"
	case MTLStepFunctionThreadPositionInGridY:
		return "MTLStepFunctionThreadPositionInGridY"
	case MTLStepFunctionThreadPositionInGridYIndexed:
		return "MTLStepFunctionThreadPositionInGridYIndexed"
	default:
		return fmt.Sprintf("MTLStepFunction(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLStitchedLibraryOptions
type MTLStitchedLibraryOptions uint

const (
	// MTLStitchedLibraryOptionFailOnBinaryArchiveMiss: An option that instructs the compiler to return an error when a GPU function for a stitched library isn’t in a binary archive.
	MTLStitchedLibraryOptionFailOnBinaryArchiveMiss            MTLStitchedLibraryOptions = 1
	MTLStitchedLibraryOptionNone                               MTLStitchedLibraryOptions = 0
	MTLStitchedLibraryOptionStoreLibraryInMetalPipelinesScript MTLStitchedLibraryOptions = 2
)

func (e MTLStitchedLibraryOptions) String() string {
	switch e {
	case MTLStitchedLibraryOptionFailOnBinaryArchiveMiss:
		return "MTLStitchedLibraryOptionFailOnBinaryArchiveMiss"
	case MTLStitchedLibraryOptionNone:
		return "MTLStitchedLibraryOptionNone"
	case MTLStitchedLibraryOptionStoreLibraryInMetalPipelinesScript:
		return "MTLStitchedLibraryOptionStoreLibraryInMetalPipelinesScript"
	default:
		return fmt.Sprintf("MTLStitchedLibraryOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLStorageMode
type MTLStorageMode uint

const (
	// MTLStorageModeManaged: The CPU and GPU may maintain separate copies of the resource, and any changes need to be explicitly synchronized.
	MTLStorageModeManaged MTLStorageMode = 1
	// MTLStorageModeMemoryless: The resource’s contents are only available to the GPU, and only exist temporarily during a render pass.
	MTLStorageModeMemoryless MTLStorageMode = 3
	// MTLStorageModePrivate: The resource is only available to the GPU.
	MTLStorageModePrivate MTLStorageMode = 2
	// MTLStorageModeShared: The CPU and GPU share access to the resource, allocated in system memory.
	MTLStorageModeShared MTLStorageMode = 0
)

func (e MTLStorageMode) String() string {
	switch e {
	case MTLStorageModeManaged:
		return "MTLStorageModeManaged"
	case MTLStorageModeMemoryless:
		return "MTLStorageModeMemoryless"
	case MTLStorageModePrivate:
		return "MTLStorageModePrivate"
	case MTLStorageModeShared:
		return "MTLStorageModeShared"
	default:
		return fmt.Sprintf("MTLStorageMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLStoreAction
type MTLStoreAction uint

const (
	// MTLStoreActionCustomSampleDepthStore: The GPU stores depth data in a sample-position–agnostic representation.
	MTLStoreActionCustomSampleDepthStore MTLStoreAction = 5
	// MTLStoreActionDontCare: The GPU has permission to discard the rendered contents of the attachment at the end of the render pass, replacing them with arbitrary data.
	MTLStoreActionDontCare MTLStoreAction = 0
	// MTLStoreActionMultisampleResolve: The GPU resolves the multisampled data to one sample per pixel and stores the data to the resolve texture, discarding the multisample data afterwards.
	MTLStoreActionMultisampleResolve MTLStoreAction = 2
	// MTLStoreActionStore: The GPU stores the rendered contents to the texture.
	MTLStoreActionStore MTLStoreAction = 1
	// MTLStoreActionStoreAndMultisampleResolve: The GPU stores the multisample data to the multisample texture, resolves the data to a sample per pixel, and stores the data to the resolve texture.
	MTLStoreActionStoreAndMultisampleResolve MTLStoreAction = 3
	// MTLStoreActionUnknown: The system selects a store action when it encodes the render pass.
	MTLStoreActionUnknown MTLStoreAction = 4
)

func (e MTLStoreAction) String() string {
	switch e {
	case MTLStoreActionCustomSampleDepthStore:
		return "MTLStoreActionCustomSampleDepthStore"
	case MTLStoreActionDontCare:
		return "MTLStoreActionDontCare"
	case MTLStoreActionMultisampleResolve:
		return "MTLStoreActionMultisampleResolve"
	case MTLStoreActionStore:
		return "MTLStoreActionStore"
	case MTLStoreActionStoreAndMultisampleResolve:
		return "MTLStoreActionStoreAndMultisampleResolve"
	case MTLStoreActionUnknown:
		return "MTLStoreActionUnknown"
	default:
		return fmt.Sprintf("MTLStoreAction(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLStoreActionOptions
type MTLStoreActionOptions uint

const (
	// MTLStoreActionOptionCustomSamplePositions: An option that stores data in a sample-position–agnostic representation.
	MTLStoreActionOptionCustomSamplePositions MTLStoreActionOptions = 1
	// MTLStoreActionOptionNone: An option that doesn’t modify the intended behavior of a store action.
	MTLStoreActionOptionNone MTLStoreActionOptions = 0
)

func (e MTLStoreActionOptions) String() string {
	switch e {
	case MTLStoreActionOptionCustomSamplePositions:
		return "MTLStoreActionOptionCustomSamplePositions"
	case MTLStoreActionOptionNone:
		return "MTLStoreActionOptionNone"
	default:
		return fmt.Sprintf("MTLStoreActionOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLTensorDataType
type MTLTensorDataType int

const (
	MTLTensorDataTypeBFloat16 MTLTensorDataType = 121
	MTLTensorDataTypeFloat16  MTLTensorDataType = 16
	MTLTensorDataTypeFloat32  MTLTensorDataType = 3
	MTLTensorDataTypeInt16    MTLTensorDataType = 37
	MTLTensorDataTypeInt32    MTLTensorDataType = 29
	MTLTensorDataTypeInt4     MTLTensorDataType = 143
	MTLTensorDataTypeInt8     MTLTensorDataType = 45
	MTLTensorDataTypeNone     MTLTensorDataType = 0
	MTLTensorDataTypeUInt16   MTLTensorDataType = 41
	MTLTensorDataTypeUInt32   MTLTensorDataType = 33
	MTLTensorDataTypeUInt4    MTLTensorDataType = 144
	MTLTensorDataTypeUInt8    MTLTensorDataType = 49
)

func (e MTLTensorDataType) String() string {
	switch e {
	case MTLTensorDataTypeBFloat16:
		return "MTLTensorDataTypeBFloat16"
	case MTLTensorDataTypeFloat16:
		return "MTLTensorDataTypeFloat16"
	case MTLTensorDataTypeFloat32:
		return "MTLTensorDataTypeFloat32"
	case MTLTensorDataTypeInt16:
		return "MTLTensorDataTypeInt16"
	case MTLTensorDataTypeInt32:
		return "MTLTensorDataTypeInt32"
	case MTLTensorDataTypeInt4:
		return "MTLTensorDataTypeInt4"
	case MTLTensorDataTypeInt8:
		return "MTLTensorDataTypeInt8"
	case MTLTensorDataTypeNone:
		return "MTLTensorDataTypeNone"
	case MTLTensorDataTypeUInt16:
		return "MTLTensorDataTypeUInt16"
	case MTLTensorDataTypeUInt32:
		return "MTLTensorDataTypeUInt32"
	case MTLTensorDataTypeUInt4:
		return "MTLTensorDataTypeUInt4"
	case MTLTensorDataTypeUInt8:
		return "MTLTensorDataTypeUInt8"
	default:
		return fmt.Sprintf("MTLTensorDataType(%d)", e)
	}
}

type MTLTensorError int

const (
	MTLTensorErrorInternalError     MTLTensorError = 1
	MTLTensorErrorInvalidDescriptor MTLTensorError = 2
	MTLTensorErrorNone              MTLTensorError = 0
)

func (e MTLTensorError) String() string {
	switch e {
	case MTLTensorErrorInternalError:
		return "MTLTensorErrorInternalError"
	case MTLTensorErrorInvalidDescriptor:
		return "MTLTensorErrorInvalidDescriptor"
	case MTLTensorErrorNone:
		return "MTLTensorErrorNone"
	default:
		return fmt.Sprintf("MTLTensorError(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLTensorUsage
type MTLTensorUsage uint

const (
	// MTLTensorUsageCompute: A tensor context that applies to compute encoders.
	MTLTensorUsageCompute MTLTensorUsage = 1
	// MTLTensorUsageMachineLearning: A tensor context that applies to machine learning encoders.
	MTLTensorUsageMachineLearning MTLTensorUsage = 4
	// MTLTensorUsageRender: A tensor context that applies to render encoders.
	MTLTensorUsageRender MTLTensorUsage = 2
)

func (e MTLTensorUsage) String() string {
	switch e {
	case MTLTensorUsageCompute:
		return "MTLTensorUsageCompute"
	case MTLTensorUsageMachineLearning:
		return "MTLTensorUsageMachineLearning"
	case MTLTensorUsageRender:
		return "MTLTensorUsageRender"
	default:
		return fmt.Sprintf("MTLTensorUsage(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLTessellationControlPointIndexType
type MTLTessellationControlPointIndexType uint

const (
	// MTLTessellationControlPointIndexTypeNone: No size.
	MTLTessellationControlPointIndexTypeNone MTLTessellationControlPointIndexType = 0
	// MTLTessellationControlPointIndexTypeUInt16: The size of a 16-bit unsigned integer.
	MTLTessellationControlPointIndexTypeUInt16 MTLTessellationControlPointIndexType = 1
	// MTLTessellationControlPointIndexTypeUInt32: The size of a 32-bit unsigned integer.
	MTLTessellationControlPointIndexTypeUInt32 MTLTessellationControlPointIndexType = 2
)

func (e MTLTessellationControlPointIndexType) String() string {
	switch e {
	case MTLTessellationControlPointIndexTypeNone:
		return "MTLTessellationControlPointIndexTypeNone"
	case MTLTessellationControlPointIndexTypeUInt16:
		return "MTLTessellationControlPointIndexTypeUInt16"
	case MTLTessellationControlPointIndexTypeUInt32:
		return "MTLTessellationControlPointIndexTypeUInt32"
	default:
		return fmt.Sprintf("MTLTessellationControlPointIndexType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLTessellationFactorFormat
type MTLTessellationFactorFormat uint

const (
	// MTLTessellationFactorFormatHalf: A 16-bit floating-point format.
	MTLTessellationFactorFormatHalf MTLTessellationFactorFormat = 0
)

func (e MTLTessellationFactorFormat) String() string {
	switch e {
	case MTLTessellationFactorFormatHalf:
		return "MTLTessellationFactorFormatHalf"
	default:
		return fmt.Sprintf("MTLTessellationFactorFormat(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLTessellationFactorStepFunction
type MTLTessellationFactorStepFunction uint

const (
	// MTLTessellationFactorStepFunctionConstant: A constant step function.
	MTLTessellationFactorStepFunctionConstant MTLTessellationFactorStepFunction = 0
	// MTLTessellationFactorStepFunctionPerInstance: A per-instance step function.
	MTLTessellationFactorStepFunctionPerInstance MTLTessellationFactorStepFunction = 2
	// MTLTessellationFactorStepFunctionPerPatch: A per-patch step function.
	MTLTessellationFactorStepFunctionPerPatch MTLTessellationFactorStepFunction = 1
	// MTLTessellationFactorStepFunctionPerPatchAndPerInstance: A per-patch and per-instance step function.
	MTLTessellationFactorStepFunctionPerPatchAndPerInstance MTLTessellationFactorStepFunction = 3
)

func (e MTLTessellationFactorStepFunction) String() string {
	switch e {
	case MTLTessellationFactorStepFunctionConstant:
		return "MTLTessellationFactorStepFunctionConstant"
	case MTLTessellationFactorStepFunctionPerInstance:
		return "MTLTessellationFactorStepFunctionPerInstance"
	case MTLTessellationFactorStepFunctionPerPatch:
		return "MTLTessellationFactorStepFunctionPerPatch"
	case MTLTessellationFactorStepFunctionPerPatchAndPerInstance:
		return "MTLTessellationFactorStepFunctionPerPatchAndPerInstance"
	default:
		return fmt.Sprintf("MTLTessellationFactorStepFunction(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLTessellationPartitionMode
type MTLTessellationPartitionMode uint

const (
	// MTLTessellationPartitionModeFractionalEven: A fractional even partitioning mode.
	MTLTessellationPartitionModeFractionalEven MTLTessellationPartitionMode = 3
	// MTLTessellationPartitionModeFractionalOdd: A fractional odd partitioning mode.
	MTLTessellationPartitionModeFractionalOdd MTLTessellationPartitionMode = 2
	// MTLTessellationPartitionModeInteger: An integer partitioning mode.
	MTLTessellationPartitionModeInteger MTLTessellationPartitionMode = 1
	// MTLTessellationPartitionModePow2: A power of two partitioning mode.
	MTLTessellationPartitionModePow2 MTLTessellationPartitionMode = 0
)

func (e MTLTessellationPartitionMode) String() string {
	switch e {
	case MTLTessellationPartitionModeFractionalEven:
		return "MTLTessellationPartitionModeFractionalEven"
	case MTLTessellationPartitionModeFractionalOdd:
		return "MTLTessellationPartitionModeFractionalOdd"
	case MTLTessellationPartitionModeInteger:
		return "MTLTessellationPartitionModeInteger"
	case MTLTessellationPartitionModePow2:
		return "MTLTessellationPartitionModePow2"
	default:
		return fmt.Sprintf("MTLTessellationPartitionMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLTextureCompressionType
type MTLTextureCompressionType int

const (
	MTLTextureCompressionTypeLossless MTLTextureCompressionType = 0
	MTLTextureCompressionTypeLossy    MTLTextureCompressionType = 1
)

func (e MTLTextureCompressionType) String() string {
	switch e {
	case MTLTextureCompressionTypeLossless:
		return "MTLTextureCompressionTypeLossless"
	case MTLTextureCompressionTypeLossy:
		return "MTLTextureCompressionTypeLossy"
	default:
		return fmt.Sprintf("MTLTextureCompressionType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLTextureSparseTier
type MTLTextureSparseTier int

const (
	// MTLTextureSparseTier1: Indicates support for sparse textures tier 1.
	MTLTextureSparseTier1 MTLTextureSparseTier = 1
	// MTLTextureSparseTier2: Indicates support for sparse textures tier 2.
	MTLTextureSparseTier2 MTLTextureSparseTier = 2
	// MTLTextureSparseTierNone: Indicates that the texture is not sparse.
	MTLTextureSparseTierNone MTLTextureSparseTier = 0
)

func (e MTLTextureSparseTier) String() string {
	switch e {
	case MTLTextureSparseTier1:
		return "MTLTextureSparseTier1"
	case MTLTextureSparseTier2:
		return "MTLTextureSparseTier2"
	case MTLTextureSparseTierNone:
		return "MTLTextureSparseTierNone"
	default:
		return fmt.Sprintf("MTLTextureSparseTier(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLTextureSwizzle
type MTLTextureSwizzle uint8

const (
	// MTLTextureSwizzleAlpha: The alpha channel of the source pixel is copied to the destination channel.
	MTLTextureSwizzleAlpha MTLTextureSwizzle = 5
	// MTLTextureSwizzleBlue: The blue channel of the source pixel is copied to the destination channel.
	MTLTextureSwizzleBlue MTLTextureSwizzle = 4
	// MTLTextureSwizzleGreen: The green channel of the source pixel is copied to the destination channel.
	MTLTextureSwizzleGreen MTLTextureSwizzle = 3
	// MTLTextureSwizzleOne: A value of `1.0` is copied to the destination channel.
	MTLTextureSwizzleOne MTLTextureSwizzle = 1
	// MTLTextureSwizzleRed: The red channel of the source pixel is copied to the destination channel.
	MTLTextureSwizzleRed MTLTextureSwizzle = 2
	// MTLTextureSwizzleZero: A value of `0.0` is copied to the destination channel.
	MTLTextureSwizzleZero MTLTextureSwizzle = 0
)

func (e MTLTextureSwizzle) String() string {
	switch e {
	case MTLTextureSwizzleAlpha:
		return "MTLTextureSwizzleAlpha"
	case MTLTextureSwizzleBlue:
		return "MTLTextureSwizzleBlue"
	case MTLTextureSwizzleGreen:
		return "MTLTextureSwizzleGreen"
	case MTLTextureSwizzleOne:
		return "MTLTextureSwizzleOne"
	case MTLTextureSwizzleRed:
		return "MTLTextureSwizzleRed"
	case MTLTextureSwizzleZero:
		return "MTLTextureSwizzleZero"
	default:
		return fmt.Sprintf("MTLTextureSwizzle(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLTextureType
type MTLTextureType uint

const (
	// MTLTextureType1D: A one-dimensional texture image.
	MTLTextureType1D MTLTextureType = 0
	// MTLTextureType1DArray: An array of one-dimensional texture images.
	MTLTextureType1DArray MTLTextureType = 1
	// MTLTextureType2D: A two-dimensional texture image.
	MTLTextureType2D MTLTextureType = 2
	// MTLTextureType2DArray: An array of two-dimensional texture images.
	MTLTextureType2DArray MTLTextureType = 3
	// MTLTextureType2DMultisample: A two-dimensional texture image that uses more than one sample for each pixel.
	MTLTextureType2DMultisample MTLTextureType = 4
	// MTLTextureType2DMultisampleArray: An array of two-dimensional texture images that use more than one sample for each pixel.
	MTLTextureType2DMultisampleArray MTLTextureType = 8
	// MTLTextureType3D: A three-dimensional texture image.
	MTLTextureType3D MTLTextureType = 7
	// MTLTextureTypeCube: A cube texture with six two-dimensional images.
	MTLTextureTypeCube MTLTextureType = 5
	// MTLTextureTypeCubeArray: An array of cube textures, each with six two-dimensional images.
	MTLTextureTypeCubeArray MTLTextureType = 6
	// MTLTextureTypeTextureBuffer: A texture buffer.
	MTLTextureTypeTextureBuffer MTLTextureType = 9
)

func (e MTLTextureType) String() string {
	switch e {
	case MTLTextureType1D:
		return "MTLTextureType1D"
	case MTLTextureType1DArray:
		return "MTLTextureType1DArray"
	case MTLTextureType2D:
		return "MTLTextureType2D"
	case MTLTextureType2DArray:
		return "MTLTextureType2DArray"
	case MTLTextureType2DMultisample:
		return "MTLTextureType2DMultisample"
	case MTLTextureType2DMultisampleArray:
		return "MTLTextureType2DMultisampleArray"
	case MTLTextureType3D:
		return "MTLTextureType3D"
	case MTLTextureTypeCube:
		return "MTLTextureTypeCube"
	case MTLTextureTypeCubeArray:
		return "MTLTextureTypeCubeArray"
	case MTLTextureTypeTextureBuffer:
		return "MTLTextureTypeTextureBuffer"
	default:
		return fmt.Sprintf("MTLTextureType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLTextureUsage
type MTLTextureUsage uint

const (
	// MTLTextureUsagePixelFormatView: An option to create texture views with a different component layout.
	MTLTextureUsagePixelFormatView MTLTextureUsage = 16
	// MTLTextureUsageRenderTarget: An option for rendering to the texture in a render pass.
	MTLTextureUsageRenderTarget MTLTextureUsage = 4
	// MTLTextureUsageShaderAtomic: An option that enables atomic memory operations on texture elements in shader code.
	MTLTextureUsageShaderAtomic MTLTextureUsage = 32
	// MTLTextureUsageShaderRead: An option for reading or sampling from the texture in a shader.
	MTLTextureUsageShaderRead MTLTextureUsage = 1
	// MTLTextureUsageShaderWrite: An option for writing to the texture in a shader.
	MTLTextureUsageShaderWrite MTLTextureUsage = 2
	// MTLTextureUsageUnknown: An option for a texture whose usage is unknown.
	MTLTextureUsageUnknown MTLTextureUsage = 0
)

func (e MTLTextureUsage) String() string {
	switch e {
	case MTLTextureUsagePixelFormatView:
		return "MTLTextureUsagePixelFormatView"
	case MTLTextureUsageRenderTarget:
		return "MTLTextureUsageRenderTarget"
	case MTLTextureUsageShaderAtomic:
		return "MTLTextureUsageShaderAtomic"
	case MTLTextureUsageShaderRead:
		return "MTLTextureUsageShaderRead"
	case MTLTextureUsageShaderWrite:
		return "MTLTextureUsageShaderWrite"
	case MTLTextureUsageUnknown:
		return "MTLTextureUsageUnknown"
	default:
		return fmt.Sprintf("MTLTextureUsage(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLTransformType
type MTLTransformType int

const (
	MTLTransformTypeComponent      MTLTransformType = 1
	MTLTransformTypePackedFloat4x3 MTLTransformType = 0
)

func (e MTLTransformType) String() string {
	switch e {
	case MTLTransformTypeComponent:
		return "MTLTransformTypeComponent"
	case MTLTransformTypePackedFloat4x3:
		return "MTLTransformTypePackedFloat4x3"
	default:
		return fmt.Sprintf("MTLTransformType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLTriangleFillMode
type MTLTriangleFillMode uint

const (
	// MTLTriangleFillModeFill: Rasterize triangle and triangle strip primitives as filled triangles.
	MTLTriangleFillModeFill MTLTriangleFillMode = 0
	// MTLTriangleFillModeLines: Rasterize triangle and triangle strip primitives as lines.
	MTLTriangleFillModeLines MTLTriangleFillMode = 1
)

func (e MTLTriangleFillMode) String() string {
	switch e {
	case MTLTriangleFillModeFill:
		return "MTLTriangleFillModeFill"
	case MTLTriangleFillModeLines:
		return "MTLTriangleFillModeLines"
	default:
		return fmt.Sprintf("MTLTriangleFillMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLVertexFormat
type MTLVertexFormat uint

const (
	// MTLVertexFormatChar: One signed 8-bit two’s complement value.
	MTLVertexFormatChar MTLVertexFormat = 46
	// MTLVertexFormatChar2: Two signed 8-bit two’s complement values.
	MTLVertexFormatChar2 MTLVertexFormat = 4
	// MTLVertexFormatChar2Normalized: Two signed normalized 8-bit two’s complement values.
	MTLVertexFormatChar2Normalized MTLVertexFormat = 10
	// MTLVertexFormatChar3: Three signed 8-bit two’s complement values.
	MTLVertexFormatChar3 MTLVertexFormat = 5
	// MTLVertexFormatChar3Normalized: Three signed normalized 8-bit two’s complement values.
	MTLVertexFormatChar3Normalized MTLVertexFormat = 11
	// MTLVertexFormatChar4: Four signed 8-bit two’s complement values.
	MTLVertexFormatChar4 MTLVertexFormat = 6
	// MTLVertexFormatChar4Normalized: Four signed normalized 8-bit two’s complement values.
	MTLVertexFormatChar4Normalized MTLVertexFormat = 12
	// MTLVertexFormatCharNormalized: One signed normalized 8-bit two’s complement value.
	MTLVertexFormatCharNormalized MTLVertexFormat = 48
	// MTLVertexFormatFloat: One single-precision floating-point value.
	MTLVertexFormatFloat MTLVertexFormat = 28
	// MTLVertexFormatFloat2: Two single-precision floating-point values.
	MTLVertexFormatFloat2 MTLVertexFormat = 29
	// MTLVertexFormatFloat3: Three single-precision floating-point values.
	MTLVertexFormatFloat3 MTLVertexFormat = 30
	// MTLVertexFormatFloat4: Four single-precision floating-point values.
	MTLVertexFormatFloat4       MTLVertexFormat = 31
	MTLVertexFormatFloatRG11B10 MTLVertexFormat = 54
	MTLVertexFormatFloatRGB9E5  MTLVertexFormat = 55
	// MTLVertexFormatHalf: One half-precision floating-point value.
	MTLVertexFormatHalf MTLVertexFormat = 53
	// MTLVertexFormatHalf2: Two half-precision floating-point values.
	MTLVertexFormatHalf2 MTLVertexFormat = 25
	// MTLVertexFormatHalf3: Three half-precision floating-point values.
	MTLVertexFormatHalf3 MTLVertexFormat = 26
	// MTLVertexFormatHalf4: Four half-precision floating-point values.
	MTLVertexFormatHalf4 MTLVertexFormat = 27
	// MTLVertexFormatInt: A 32-bit, signed integer value.
	MTLVertexFormatInt MTLVertexFormat = 32
	// MTLVertexFormatInt1010102Normalized: A four-component vector with 10-bit, normalized, signed integer values for red, green, and blue, and a 2-bit value for alpha.
	MTLVertexFormatInt1010102Normalized MTLVertexFormat = 40
	// MTLVertexFormatInt2: A two-component vector with 32-bit, signed integer values.
	MTLVertexFormatInt2 MTLVertexFormat = 33
	// MTLVertexFormatInt3: A three-component vector with 32-bit, signed integer values.
	MTLVertexFormatInt3 MTLVertexFormat = 34
	// MTLVertexFormatInt4: A four-component vector with 32-bit, signed integer values.
	MTLVertexFormatInt4 MTLVertexFormat = 35
	// MTLVertexFormatInvalid: A sentinel value that represents an empty set of vertex format options.
	MTLVertexFormatInvalid MTLVertexFormat = 0
	// MTLVertexFormatShort: A 16-bit, signed integer value.
	MTLVertexFormatShort MTLVertexFormat = 50
	// MTLVertexFormatShort2: A two-component vector with 16-bit, signed integer values.
	MTLVertexFormatShort2 MTLVertexFormat = 16
	// MTLVertexFormatShort2Normalized: A two-component vector with 16-bit, normalized, signed integer values.
	MTLVertexFormatShort2Normalized MTLVertexFormat = 22
	// MTLVertexFormatShort3: A three-component vector with 16-bit, signed integer values.
	MTLVertexFormatShort3 MTLVertexFormat = 17
	// MTLVertexFormatShort3Normalized: A three-component vector with 16-bit, normalized, signed integer values.
	MTLVertexFormatShort3Normalized MTLVertexFormat = 23
	// MTLVertexFormatShort4: A four-component vector with 16-bit, signed integer values.
	MTLVertexFormatShort4 MTLVertexFormat = 18
	// MTLVertexFormatShort4Normalized: A four-component vector with 16-bit, normalized, signed integer values.
	MTLVertexFormatShort4Normalized MTLVertexFormat = 24
	// MTLVertexFormatShortNormalized: A 16-bit, normalized, signed integer value.
	MTLVertexFormatShortNormalized MTLVertexFormat = 52
	// MTLVertexFormatUChar: An 8-bit, unsigned integer value.
	MTLVertexFormatUChar MTLVertexFormat = 45
	// MTLVertexFormatUChar2: A two-component vector with 8-bit, unsigned integer values.
	MTLVertexFormatUChar2 MTLVertexFormat = 1
	// MTLVertexFormatUChar2Normalized: A two-component vector with 8-bit, normalized, unsigned integer values.
	MTLVertexFormatUChar2Normalized MTLVertexFormat = 7
	// MTLVertexFormatUChar3: A three-component vector with 8-bit, unsigned integer values.
	MTLVertexFormatUChar3 MTLVertexFormat = 2
	// MTLVertexFormatUChar3Normalized: A three-component vector with 8-bit, normalized, unsigned integer values.
	MTLVertexFormatUChar3Normalized MTLVertexFormat = 8
	// MTLVertexFormatUChar4: A four-component vector with 8-bit, unsigned integer values.
	MTLVertexFormatUChar4 MTLVertexFormat = 3
	// MTLVertexFormatUChar4Normalized: A four-component vector with 8-bit, normalized, unsigned integer values.
	MTLVertexFormatUChar4Normalized MTLVertexFormat = 9
	// MTLVertexFormatUChar4Normalized_BGRA: A four-component vector with 8-bit, normalized, unsigned integer values for blue, green, red, and alpha.
	MTLVertexFormatUChar4Normalized_BGRA MTLVertexFormat = 42
	// MTLVertexFormatUCharNormalized: An 8-bit, normalized, unsigned integer value.
	MTLVertexFormatUCharNormalized MTLVertexFormat = 47
	// MTLVertexFormatUInt: A 32-bit, unsigned integer value.
	MTLVertexFormatUInt MTLVertexFormat = 36
	// MTLVertexFormatUInt1010102Normalized: A four-component vector with 10-bit, normalized, unsigned integer values for red, green, and blue, and a 2-bit value for alpha.
	MTLVertexFormatUInt1010102Normalized MTLVertexFormat = 41
	// MTLVertexFormatUInt2: A two-component vector with 32-bit, unsigned integer values.
	MTLVertexFormatUInt2 MTLVertexFormat = 37
	// MTLVertexFormatUInt3: A three-component vector with 32-bit, unsigned integer values.
	MTLVertexFormatUInt3 MTLVertexFormat = 38
	// MTLVertexFormatUInt4: A four-component vector with 32-bit, unsigned integer values.
	MTLVertexFormatUInt4 MTLVertexFormat = 39
	// MTLVertexFormatUShort: A 16-bit, unsigned integer value.
	MTLVertexFormatUShort MTLVertexFormat = 49
	// MTLVertexFormatUShort2: A two-component vector with 16-bit, unsigned integer values.
	MTLVertexFormatUShort2 MTLVertexFormat = 13
	// MTLVertexFormatUShort2Normalized: A two-component vector with 16-bit, normalized, unsigned integer values.
	MTLVertexFormatUShort2Normalized MTLVertexFormat = 19
	// MTLVertexFormatUShort3: A three-component vector with 16-bit, unsigned integer values.
	MTLVertexFormatUShort3 MTLVertexFormat = 14
	// MTLVertexFormatUShort3Normalized: A three-component vector with 16-bit, normalized, unsigned integer values.
	MTLVertexFormatUShort3Normalized MTLVertexFormat = 20
	// MTLVertexFormatUShort4: A four-component vector with 16-bit, unsigned integer values.
	MTLVertexFormatUShort4 MTLVertexFormat = 15
	// MTLVertexFormatUShort4Normalized: A four-component vector with 16-bit, normalized, unsigned integer values.
	MTLVertexFormatUShort4Normalized MTLVertexFormat = 21
	// MTLVertexFormatUShortNormalized: A 16-bit, normalized, unsigned integer value.
	MTLVertexFormatUShortNormalized MTLVertexFormat = 51
)

func (e MTLVertexFormat) String() string {
	switch e {
	case MTLVertexFormatChar:
		return "MTLVertexFormatChar"
	case MTLVertexFormatChar2:
		return "MTLVertexFormatChar2"
	case MTLVertexFormatChar2Normalized:
		return "MTLVertexFormatChar2Normalized"
	case MTLVertexFormatChar3:
		return "MTLVertexFormatChar3"
	case MTLVertexFormatChar3Normalized:
		return "MTLVertexFormatChar3Normalized"
	case MTLVertexFormatChar4:
		return "MTLVertexFormatChar4"
	case MTLVertexFormatChar4Normalized:
		return "MTLVertexFormatChar4Normalized"
	case MTLVertexFormatCharNormalized:
		return "MTLVertexFormatCharNormalized"
	case MTLVertexFormatFloat:
		return "MTLVertexFormatFloat"
	case MTLVertexFormatFloat2:
		return "MTLVertexFormatFloat2"
	case MTLVertexFormatFloat3:
		return "MTLVertexFormatFloat3"
	case MTLVertexFormatFloat4:
		return "MTLVertexFormatFloat4"
	case MTLVertexFormatFloatRG11B10:
		return "MTLVertexFormatFloatRG11B10"
	case MTLVertexFormatFloatRGB9E5:
		return "MTLVertexFormatFloatRGB9E5"
	case MTLVertexFormatHalf:
		return "MTLVertexFormatHalf"
	case MTLVertexFormatHalf2:
		return "MTLVertexFormatHalf2"
	case MTLVertexFormatHalf3:
		return "MTLVertexFormatHalf3"
	case MTLVertexFormatHalf4:
		return "MTLVertexFormatHalf4"
	case MTLVertexFormatInt:
		return "MTLVertexFormatInt"
	case MTLVertexFormatInt1010102Normalized:
		return "MTLVertexFormatInt1010102Normalized"
	case MTLVertexFormatInt2:
		return "MTLVertexFormatInt2"
	case MTLVertexFormatInt3:
		return "MTLVertexFormatInt3"
	case MTLVertexFormatInt4:
		return "MTLVertexFormatInt4"
	case MTLVertexFormatInvalid:
		return "MTLVertexFormatInvalid"
	case MTLVertexFormatShort:
		return "MTLVertexFormatShort"
	case MTLVertexFormatShort2:
		return "MTLVertexFormatShort2"
	case MTLVertexFormatShort2Normalized:
		return "MTLVertexFormatShort2Normalized"
	case MTLVertexFormatShort3:
		return "MTLVertexFormatShort3"
	case MTLVertexFormatShort3Normalized:
		return "MTLVertexFormatShort3Normalized"
	case MTLVertexFormatShort4:
		return "MTLVertexFormatShort4"
	case MTLVertexFormatShort4Normalized:
		return "MTLVertexFormatShort4Normalized"
	case MTLVertexFormatShortNormalized:
		return "MTLVertexFormatShortNormalized"
	case MTLVertexFormatUChar:
		return "MTLVertexFormatUChar"
	case MTLVertexFormatUChar2:
		return "MTLVertexFormatUChar2"
	case MTLVertexFormatUChar2Normalized:
		return "MTLVertexFormatUChar2Normalized"
	case MTLVertexFormatUChar3:
		return "MTLVertexFormatUChar3"
	case MTLVertexFormatUChar3Normalized:
		return "MTLVertexFormatUChar3Normalized"
	case MTLVertexFormatUChar4:
		return "MTLVertexFormatUChar4"
	case MTLVertexFormatUChar4Normalized:
		return "MTLVertexFormatUChar4Normalized"
	case MTLVertexFormatUChar4Normalized_BGRA:
		return "MTLVertexFormatUChar4Normalized_BGRA"
	case MTLVertexFormatUCharNormalized:
		return "MTLVertexFormatUCharNormalized"
	case MTLVertexFormatUInt:
		return "MTLVertexFormatUInt"
	case MTLVertexFormatUInt1010102Normalized:
		return "MTLVertexFormatUInt1010102Normalized"
	case MTLVertexFormatUInt2:
		return "MTLVertexFormatUInt2"
	case MTLVertexFormatUInt3:
		return "MTLVertexFormatUInt3"
	case MTLVertexFormatUInt4:
		return "MTLVertexFormatUInt4"
	case MTLVertexFormatUShort:
		return "MTLVertexFormatUShort"
	case MTLVertexFormatUShort2:
		return "MTLVertexFormatUShort2"
	case MTLVertexFormatUShort2Normalized:
		return "MTLVertexFormatUShort2Normalized"
	case MTLVertexFormatUShort3:
		return "MTLVertexFormatUShort3"
	case MTLVertexFormatUShort3Normalized:
		return "MTLVertexFormatUShort3Normalized"
	case MTLVertexFormatUShort4:
		return "MTLVertexFormatUShort4"
	case MTLVertexFormatUShort4Normalized:
		return "MTLVertexFormatUShort4Normalized"
	case MTLVertexFormatUShortNormalized:
		return "MTLVertexFormatUShortNormalized"
	default:
		return fmt.Sprintf("MTLVertexFormat(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLVertexStepFunction
type MTLVertexStepFunction uint

const (
	// MTLVertexStepFunctionConstant: The vertex function fetches attribute data once and uses that data for every vertex.
	MTLVertexStepFunctionConstant MTLVertexStepFunction = 0
	// MTLVertexStepFunctionPerInstance: The vertex function regularly fetches new attribute data for a number of instances that is determined by `stepRate`.
	MTLVertexStepFunctionPerInstance MTLVertexStepFunction = 2
	// MTLVertexStepFunctionPerPatch: The post-tessellation vertex function fetches data based on the patch index of the patch.
	MTLVertexStepFunctionPerPatch MTLVertexStepFunction = 3
	// MTLVertexStepFunctionPerPatchControlPoint: The post-tessellation vertex function fetches data based on the control-point indices associated with the patch.
	MTLVertexStepFunctionPerPatchControlPoint MTLVertexStepFunction = 4
	// MTLVertexStepFunctionPerVertex: The vertex function fetches and uses new attribute data for every vertex.
	MTLVertexStepFunctionPerVertex MTLVertexStepFunction = 1
)

func (e MTLVertexStepFunction) String() string {
	switch e {
	case MTLVertexStepFunctionConstant:
		return "MTLVertexStepFunctionConstant"
	case MTLVertexStepFunctionPerInstance:
		return "MTLVertexStepFunctionPerInstance"
	case MTLVertexStepFunctionPerPatch:
		return "MTLVertexStepFunctionPerPatch"
	case MTLVertexStepFunctionPerPatchControlPoint:
		return "MTLVertexStepFunctionPerPatchControlPoint"
	case MTLVertexStepFunctionPerVertex:
		return "MTLVertexStepFunctionPerVertex"
	default:
		return fmt.Sprintf("MTLVertexStepFunction(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLVisibilityResultMode
type MTLVisibilityResultMode int

const (
	// MTLVisibilityResultModeBoolean: The result records whether any samples passed depth and stencil tests.
	MTLVisibilityResultModeBoolean MTLVisibilityResultMode = 1
	// MTLVisibilityResultModeCounting: The result records how many samples passed depth and stencil tests.
	MTLVisibilityResultModeCounting MTLVisibilityResultMode = 2
	// MTLVisibilityResultModeDisabled: The result doesn’t contain any data because visibility testing was disabled.
	MTLVisibilityResultModeDisabled MTLVisibilityResultMode = 0
)

func (e MTLVisibilityResultMode) String() string {
	switch e {
	case MTLVisibilityResultModeBoolean:
		return "MTLVisibilityResultModeBoolean"
	case MTLVisibilityResultModeCounting:
		return "MTLVisibilityResultModeCounting"
	case MTLVisibilityResultModeDisabled:
		return "MTLVisibilityResultModeDisabled"
	default:
		return fmt.Sprintf("MTLVisibilityResultMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLVisibilityResultType
type MTLVisibilityResultType int

const (
	// MTLVisibilityResultTypeAccumulate: Accumulate visibility results data across multiple render passes.
	MTLVisibilityResultTypeAccumulate MTLVisibilityResultType = 1
	// MTLVisibilityResultTypeReset: Reset visibility result data when you create a render command encoder.
	MTLVisibilityResultTypeReset MTLVisibilityResultType = 0
)

func (e MTLVisibilityResultType) String() string {
	switch e {
	case MTLVisibilityResultTypeAccumulate:
		return "MTLVisibilityResultTypeAccumulate"
	case MTLVisibilityResultTypeReset:
		return "MTLVisibilityResultTypeReset"
	default:
		return fmt.Sprintf("MTLVisibilityResultType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Metal/MTLWinding
type MTLWinding uint

const (
	// MTLWindingClockwise: Primitives whose vertices are specified in clockwise order are front-facing.
	MTLWindingClockwise MTLWinding = 0
	// MTLWindingCounterClockwise: Primitives whose vertices are specified in counter-clockwise order are front-facing.
	MTLWindingCounterClockwise MTLWinding = 1
)

func (e MTLWinding) String() string {
	switch e {
	case MTLWindingClockwise:
		return "MTLWindingClockwise"
	case MTLWindingCounterClockwise:
		return "MTLWindingCounterClockwise"
	default:
		return fmt.Sprintf("MTLWinding(%d)", e)
	}
}
