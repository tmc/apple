// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal
import (
	"github.com/tmc/apple/foundation"
)

// C struct types
// MTL4BufferRange
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4BufferRange
type MTL4BufferRange struct {
	BufferAddress MTLGPUAddress
	Length uint64

}

// MTL4CopySparseBufferMappingOperation - Groups together arguments for an operation to copy a sparse buffer mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CopySparseBufferMappingOperation
type MTL4CopySparseBufferMappingOperation struct {
	DestinationOffset uint // The origin in the destination buffer, in tiles.
	SourceRange foundation.NSRange // The range in the source buffer, in tiles.

}

// MTL4CopySparseTextureMappingOperation - Groups together arguments for an operation to copy a sparse texture mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4CopySparseTextureMappingOperation
type MTL4CopySparseTextureMappingOperation struct {
	DestinationLevel uint // The index of the mipmap level in the destination texture.
	DestinationOrigin MTLOrigin // The origin in the destination texture to copy into, in tiles.
	DestinationSlice uint // The index of the array slice in the destination texture to copy into.
	SourceLevel uint // The index of the mipmap level in the source texture.
	SourceRegion MTLRegion // The region in the source texture, in tiles.
	SourceSlice uint // The index of the array slice in the texture source of the copy operation.

}

// MTL4TimestampHeapEntry - Represents a timestamp data entry in a counter heap of type [MTL4CounterHeapTypeTimestamp].
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4TimestampHeapEntry
type MTL4TimestampHeapEntry struct {
	Timestamp uint64

}

// MTL4UpdateSparseBufferMappingOperation - Groups together arguments for an operation to update a sparse buffer mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4UpdateSparseBufferMappingOperation
type MTL4UpdateSparseBufferMappingOperation struct {
	BufferRange foundation.NSRange // The range in the buffer, in tiles.
	HeapOffset uint // The starting offset in the heap, in tiles.
	Mode MTLSparseTextureMappingMode // The mode of the mapping operation to perform.

}

// MTL4UpdateSparseTextureMappingOperation - Groups together arguments for an operation to update a sparse texture mapping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTL4UpdateSparseTextureMappingOperation
type MTL4UpdateSparseTextureMappingOperation struct {
	HeapOffset uint // The starting offset in the heap, in tiles.
	Mode MTLSparseTextureMappingMode // The mode of the mapping operation to perform.
	TextureLevel uint // The index of the mipmap level in the texture to update.
	TextureRegion MTLRegion // The region in the texture to update, in tiles.
	TextureSlice uint // The index of the array slice in the texture to update.

}

// MTLAccelerationStructureInstanceDescriptor - A description of an instance in an instanced geometry acceleration structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureInstanceDescriptor
type MTLAccelerationStructureInstanceDescriptor struct {
	AccelerationStructureIndex uint32 // The index of the acceleration structure to use for the instance.
	TransformationMatrix MTLPackedFloat4x3 // The transform for placing and orienting the instance in the scene.
	IntersectionFunctionTableOffset uint32 // An offset for determining which function in the intersection function table Metal needs to call when testing a ray against the instance.
	Options MTLAccelerationStructureInstanceOptions // The options for the instance.
	Mask uint32 // A mask to use for the instance when testing a ray against the geometry.

}

// MTLAccelerationStructureMotionInstanceDescriptor - A description of an instance in an instanced geometry acceleration structure, with the instance including a user identifier and motion data for the instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureMotionInstanceDescriptor
type MTLAccelerationStructureMotionInstanceDescriptor struct {
	AccelerationStructureIndex uint32 // The index of an acceleration structure which applies to the next acceleration-structure motion instance you create with the descriptor.
	MotionStartTime float32 // A starting time for the range of motion that the key-frame data represents.
	MotionEndTime float32 // An ending time for the range of motion that the key-frame data represents.
	MotionStartBorderMode MTLMotionBorderMode // A behavior that configures how a motion instance handles timestamps before a starting time.
	MotionEndBorderMode MTLMotionBorderMode // A behavior that configures how a motion instance handles timestamps after an ending time.
	MotionTransformsStartIndex uint32 // The index of motion data that represents the first key-frame motion data, which applies to the next acceleration-structure motion instance you create with the descriptor.
	MotionTransformsCount uint32 // The number of motion data key-frames, which applies to the next acceleration-structure motion instance you create with the descriptor.
	IntersectionFunctionTableOffset uint32 // An offset into the intersection-function table for ray tracing, which applies to the next acceleration-structure motion instance you create with the descriptor.
	Options MTLAccelerationStructureInstanceOptions // An option set which applies to the next acceleration structure motion-instance you create with the descriptor.
	Mask uint32 // A mask for testing ray-tracing rays with a scene’s geometry, which applies to the next acceleration-structure motion instance you create with the descriptor.
	UserID uint32 // An unique identifier, which applies to the next acceleration-structure motion instance you create with the descriptor.

}

// MTLAccelerationStructureSizes - The expected sizes for a ray-tracing acceleration structure.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureSizes
type MTLAccelerationStructureSizes struct {
	AccelerationStructureSize uint // The size of the acceleration structure, in bytes.
	BuildScratchBufferSize uint // The amount of scratch memory, in bytes, the GPU devices needs to build the acceleration structure.
	RefitScratchBufferSize uint // The amount of scratch memory, in bytes, the GPU device needs to refit the acceleration structure.

}

// MTLAccelerationStructureUserIDInstanceDescriptor - A description of an instance in an instanced geometry acceleration structure, with the instance including a user identifier for the instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAccelerationStructureUserIDInstanceDescriptor
type MTLAccelerationStructureUserIDInstanceDescriptor struct {
	AccelerationStructureIndex uint32 // The index of the acceleration structure to use for the instance.
	TransformationMatrix MTLPackedFloat4x3 // The transform for placing and orienting the instance in the scene.
	IntersectionFunctionTableOffset uint32 // An offset for determining which function in the intersection function table Metal calls when testing a ray against the instance.
	Options MTLAccelerationStructureInstanceOptions // The options for the instance.
	Mask uint32 // A mask to use for the instance when testing a ray against the geometry.
	UserID uint32 // The user identifier for the instance.

}

// MTLAxisAlignedBoundingBox - The bounds for an axis-aligned bounding box.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLAxisAlignedBoundingBox-c.struct
type MTLAxisAlignedBoundingBox struct {
	Min MTLPackedFloat3
	Max MTLPackedFloat3

}

// MTLClearColor - An RGBA value used for a color pixel.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLClearColor
type MTLClearColor struct {
	Red float64 // The red color channel.
	Green float64 // The green color channel.
	Blue float64 // The blue color channel.
	Alpha float64 // The alpha channel.

}

// MTLComponentTransform
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLComponentTransform
type MTLComponentTransform struct {
	Scale MTLPackedFloat3
	Shear MTLPackedFloat3
	Pivot MTLPackedFloat3
	Rotation MTLPackedFloatQuaternion
	Translation MTLPackedFloat3

}

// MTLCounterResultStageUtilization - The data structure for storing the data you resolve from a stage-utilization counter set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCounterResultStageUtilization
type MTLCounterResultStageUtilization struct {
	TotalCycles uint64 // The total number of cycles the GPU uses to run a pass.
	VertexCycles uint64 // The number of cycles the GPU uses to run vertex shaders during a pass.
	TessellationCycles uint64 // The number of cycles the GPU uses to run the tessellation stage during a pass.
	PostTessellationVertexCycles uint64 // The number of cycles the GPU uses to run post-tessellation vertex shaders during a pass.
	FragmentCycles uint64 // The number of cycles the GPU uses to run fragment shaders during a pass.
	RenderTargetCycles uint64 // The number of cycles the GPU uses to write data to render targets during a render pass.

}

// MTLCounterResultStatistic - The data structure for storing the data you resolve from a statistic counter set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCounterResultStatistic
type MTLCounterResultStatistic struct {
	TessellationInputPatches uint64 // The number of tessellation patches a render pass sends to the tessellation stage.
	VertexInvocations uint64 // The number of times a render pass calls any vertex shader.
	PostTessellationVertexInvocations uint64 // The number of vertices a render pass sends to a post-tessellation vertex shader.
	ClipperInvocations uint64 // The number of primitives a render pass sends to the clip stage.
	ClipperPrimitivesOut uint64 // The number of primitives the clip stage produces during a render pass.
	FragmentInvocations uint64 // The number of times a render pass calls fragment shaders.
	FragmentsPassed uint64 // The number of fragments a render pass sends to the visibility and blend stages because they pass the scissor, depth, and stencil tests.
	ComputeKernelInvocations uint64 // The number of times a pass calls any compute kernel.

}

// MTLCounterResultTimestamp - The data structure for storing the data you resolve from a timestamp counter set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLCounterResultTimestamp
type MTLCounterResultTimestamp struct {
	Timestamp uint64 // A timestamp value from a GPU at a particular point in time during an operation, typically at the beginning or ending of a render stage.

}

// MTLDispatchThreadgroupsIndirectArguments - The data layout required for arguments needed to specify the size of threadgroups.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDispatchThreadgroupsIndirectArguments
type MTLDispatchThreadgroupsIndirectArguments struct {
	ThreadgroupsPerGrid uint32 // The number of threadgroups for the grid, in each dimension.

}

// MTLDispatchThreadsIndirectArguments
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDispatchThreadsIndirectArguments
type MTLDispatchThreadsIndirectArguments struct {
	ThreadsPerGrid uint32
	ThreadsPerThreadgroup uint32

}

// MTLDrawIndexedPrimitivesIndirectArguments - The data layout required for drawing indexed primitives via indirect buffer calls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDrawIndexedPrimitivesIndirectArguments
type MTLDrawIndexedPrimitivesIndirectArguments struct {
	IndexCount uint32 // For each instance, the number of indices to read from the index buffer.
	InstanceCount uint32 // The number of instances to draw.
	IndexStart uint32 // The first index to draw.
	BaseVertex int32 // The first vertex to draw.
	BaseInstance uint32 // The first instance to draw.

}

// MTLDrawPatchIndirectArguments - The data layout required for drawing patches via indirect buffer calls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDrawPatchIndirectArguments
type MTLDrawPatchIndirectArguments struct {
	PatchCount uint32 // The number of patches in each instance.
	InstanceCount uint32 // The number of instances to draw.
	PatchStart uint32 // The patch start index.
	BaseInstance uint32 // The first instance to draw.

}

// MTLDrawPrimitivesIndirectArguments - The data layout required for drawing primitives via indirect buffer calls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLDrawPrimitivesIndirectArguments
type MTLDrawPrimitivesIndirectArguments struct {
	VertexCount uint32 // The number of vertices to draw.
	InstanceCount uint32 // The number of instances to draw.
	VertexStart uint32 // The first vertex to draw.
	BaseInstance uint32 // The first instance to draw.

}

// MTLIndirectAccelerationStructureInstanceDescriptor - A description of an instance in an instanced geometry acceleration structure that the GPU can populate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectAccelerationStructureInstanceDescriptor
type MTLIndirectAccelerationStructureInstanceDescriptor struct {
	AccelerationStructureID MTLResourceID
	IntersectionFunctionTableOffset uint32
	Mask uint32
	Options MTLAccelerationStructureInstanceOptions
	TransformationMatrix MTLPackedFloat4x3
	UserID uint32

}

// MTLIndirectAccelerationStructureMotionInstanceDescriptor - A description of an instance in an acceleration structure that the GPU can populate, with motion data for the instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectAccelerationStructureMotionInstanceDescriptor
type MTLIndirectAccelerationStructureMotionInstanceDescriptor struct {
	AccelerationStructureID MTLResourceID // The acceleration resource handle to use for this instance.
	MotionStartTime float32 // The start time of the motion instance.
	MotionStartBorderMode MTLMotionBorderMode // The motion border mode describing what happens if Metal samples the acceleration structure before the motion start time.
	MotionEndTime float32 // The end time of the motion instance.
	MotionEndBorderMode MTLMotionBorderMode // The motion border mode describing what happens if Metal samples the acceleration structure after the motion end time.
	MotionTransformsCount uint32 // The number of motion transforms belonging to the motion instance.
	MotionTransformsStartIndex uint32 // The index of the first set of transforms describing one keyframe of the animation.
	UserID uint32 // A user-assigned ID to help identify the instance.
	IntersectionFunctionTableOffset uint32 // An offset for determining which function in the intersection function table Metal calls when testing a ray against the instance.
	Mask uint32 // An instance mask to ignore geometry during ray tracing.
	Options MTLAccelerationStructureInstanceOptions // The options for this instance.

}

// MTLIndirectCommandBufferExecutionRange - A range of commands in an indirect command buffer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferExecutionRange
type MTLIndirectCommandBufferExecutionRange struct {
	Location uint32 // The first index in the command execution range.
	Length uint32 // The number of items in the command execution range.

}

// MTLIntersectionFunctionBufferArguments
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLIntersectionFunctionBufferArguments
type MTLIntersectionFunctionBufferArguments struct {
	IntersectionFunctionBuffer uint64
	IntersectionFunctionBufferSize uint64
	IntersectionFunctionStride uint64

}

// MTLMapIndirectArguments - The data layout for mapping sparse texture regions when using indirect commands.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLMapIndirectArguments
type MTLMapIndirectArguments struct {
	RegionOriginX uint32 // The x coordinate of the region to change, measured in tile coordinates.
	RegionOriginY uint32 // The y coordinate of the region to change, measured in tile coordinates.
	RegionOriginZ uint32 // The z coordinate of the region to change, measured in tile coordinates.
	RegionSizeWidth uint32 // The width of the region, measured in tile coordinates.
	RegionSizeHeight uint32 // The height of the region, measured in tile coordinates.
	RegionSizeDepth uint32 // The depth of the region, measured in tile coordinates.
	MipMapLevel uint32 // The mipmap to change.
	SliceId uint32 // The texture slice to change.

}

// MTLOrigin - The coordinates for the front upper-left corner of a region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLOrigin
type MTLOrigin struct {
	X int // The x coordinate of the origin.
	Y int // The y coordinate of the origin.
	Z uint // The z coordinate of the origin.

}

// MTLPackedFloat3 - A structure that contains three 32-bit floating-point values with no additional padding.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPackedFloat3-c.struct
type MTLPackedFloat3 struct {
	Elements [3]float32
	X float32
	Y float32
	Z float32

}

// MTLPackedFloat4x3 - A structure that contains the top three rows of a 4x4 matrix of 32-bit floating-point values, in column-major order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPackedFloat4x3-c.struct
type MTLPackedFloat4x3 struct {
	Columns [4]MTLPackedFloat3

}

// MTLPackedFloatQuaternion
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLPackedFloatQuaternion
type MTLPackedFloatQuaternion struct {
	X float32
	Y float32
	Z float32
	W float32

}

// MTLQuadTessellationFactorsHalf - The per-patch tessellation factors for a quad patch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLQuadTessellationFactorsHalf
type MTLQuadTessellationFactorsHalf struct {
	EdgeTessellationFactor uint16 // The edge tessellation factors, with each index value providing the tessellation factor for a particular edge.
	InsideTessellationFactor uint16 // The inside tessellation factors, with the value in index 0 providing the horizontal tessellation factor and the value in index 1 providing the vertical tessellation factor.

}

// MTLRegion - The bounds for a subset of an instance’s elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLRegion
type MTLRegion struct {
	Origin MTLOrigin // The coordinates of the front upper-left corner of the region.
	Size MTLSize // The dimensions of the region.

}

// MTLResourceID
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLResourceID
type MTLResourceID struct {

}

// MTLSamplePosition - A subpixel sample position for use in multisample antialiasing (MSAA).
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSamplePosition
type MTLSamplePosition struct {
	X float32 // The x position of the sample on the subpixel grid.
	Y float32 // The y position of the sample on the subpixel grid.

}

// MTLScissorRect - A rectangle for the scissor fragment test.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLScissorRect
type MTLScissorRect struct {
	Height uint // The height of the scissor rectangle, in pixels.
	Width int // The width of the scissor rectangle, in pixels.
	X int // The x window coordinate of the upper-left corner of the scissor rectangle.
	Y int // The y window coordinate of the upper-left corner of the scissor rectangle.

}

// MTLSize - A type that represents one, two, or three dimensions of a type instance, such as an array or texture.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSize
type MTLSize struct {
	Width int // A value for the x-axis dimension.
	Height uint // A value for the y-axis dimension.
	Depth uint // A value for the z-axis dimension.

}

// MTLSizeAndAlign - The size and alignment of a resource, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLSizeAndAlign
type MTLSizeAndAlign struct {
	Size uint // The size of a resource, in bytes.
	Align uint // The alignment of a resource, in bytes.

}

// MTLStageInRegionIndirectArguments - The data layout required for the arguments needed to specify the stage-in region.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLStageInRegionIndirectArguments
type MTLStageInRegionIndirectArguments struct {
	StageInOrigin uint32 // The location of the upper-left corner of the block.
	StageInSize uint32 // The size of the block.

}

// MTLTextureSwizzleChannels - A pattern that modifies the data read or sampled from a texture by rearranging or duplicating the elements of a vector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTextureSwizzleChannels
type MTLTextureSwizzleChannels struct {
	Red MTLTextureSwizzle // The data copied to the first output channel.
	Green MTLTextureSwizzle // The data copied to the second output channel.
	Blue MTLTextureSwizzle // The data copied to the third output channel.
	Alpha MTLTextureSwizzle // The data copied to the fourth output channel.

}

// MTLTriangleTessellationFactorsHalf - The per-patch tessellation factors for a triangle patch.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLTriangleTessellationFactorsHalf
type MTLTriangleTessellationFactorsHalf struct {
	InsideTessellationFactor uint16 // The inside tessellation factor.
	EdgeTessellationFactor uint16 // The edge tessellation factors, with each index value providing the tessellation factor for a particular edge.

}

// MTLVertexAmplificationViewMapping - An offset applied to a render target index and viewport index.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLVertexAmplificationViewMapping
type MTLVertexAmplificationViewMapping struct {
	RenderTargetArrayIndexOffset uint32 // An offset into the list of render targets.
	ViewportArrayIndexOffset uint32 // An offset into the list of viewports.

}

// MTLViewport - A 3D rectangular region for the viewport clipping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Metal/MTLViewport
type MTLViewport struct {
	OriginX float64 // The x coordinate of the upper-left corner of the viewport.
	OriginY float64 // The y coordinate of the upper-left corner of the viewport.
	Width float64 // The width of the viewport, in pixels.
	Height float64 // The height of the viewport, in pixels.
	Znear float64 // The z coordinate of the near clipping plane of the viewport.
	Zfar float64 // The z coordinate of the far clipping plane of the viewport.

}








