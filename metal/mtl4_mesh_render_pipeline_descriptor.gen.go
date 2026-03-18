// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MTL4MeshRenderPipelineDescriptor] class.
var (
	_MTL4MeshRenderPipelineDescriptorClass     MTL4MeshRenderPipelineDescriptorClass
	_MTL4MeshRenderPipelineDescriptorClassOnce sync.Once
)

func getMTL4MeshRenderPipelineDescriptorClass() MTL4MeshRenderPipelineDescriptorClass {
	_MTL4MeshRenderPipelineDescriptorClassOnce.Do(func() {
		_MTL4MeshRenderPipelineDescriptorClass = MTL4MeshRenderPipelineDescriptorClass{class: objc.GetClass("MTL4MeshRenderPipelineDescriptor")}
	})
	return _MTL4MeshRenderPipelineDescriptorClass
}

// GetMTL4MeshRenderPipelineDescriptorClass returns the class object for MTL4MeshRenderPipelineDescriptor.
func GetMTL4MeshRenderPipelineDescriptorClass() MTL4MeshRenderPipelineDescriptorClass {
	return getMTL4MeshRenderPipelineDescriptorClass()
}

type MTL4MeshRenderPipelineDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTL4MeshRenderPipelineDescriptorClass) Alloc() MTL4MeshRenderPipelineDescriptor {
	rv := objc.Send[MTL4MeshRenderPipelineDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// Groups together properties you use to create a mesh render pipeline state
// object.
//
// # Overview
// 
// Compared to [MTLMeshRenderPipelineDescriptor], this interface doesn’t
// offer a mechanism to hint to Metal mutability of object, mesh, or fragment
// buffers. Additionally, when you use this descriptor, you don’t specify
// binary archives.
//
// # Instance Properties
//
//   - [MTL4MeshRenderPipelineDescriptor.AlphaToCoverageState]: Indicates whether to read and use the alpha channel fragment output of color attachments to compute a sample coverage mask.
//   - [MTL4MeshRenderPipelineDescriptor.SetAlphaToCoverageState]
//   - [MTL4MeshRenderPipelineDescriptor.AlphaToOneState]: Indicates whether the pipeline forces alpha channel values of color attachments to the largest representable value.
//   - [MTL4MeshRenderPipelineDescriptor.SetAlphaToOneState]
//   - [MTL4MeshRenderPipelineDescriptor.ColorAttachmentMappingState]: Sets the logical-to-physical rendering remap state.
//   - [MTL4MeshRenderPipelineDescriptor.SetColorAttachmentMappingState]
//   - [MTL4MeshRenderPipelineDescriptor.ColorAttachments]: Accesses an array containing descriptions of the color attachments this pipeline writes to.
//   - [MTL4MeshRenderPipelineDescriptor.FragmentFunctionDescriptor]: Assigns a function descriptor representing the function this pipeline executes for each fragment.
//   - [MTL4MeshRenderPipelineDescriptor.SetFragmentFunctionDescriptor]
//   - [MTL4MeshRenderPipelineDescriptor.FragmentStaticLinkingDescriptor]: Provides static linking information for the fragment stage of the render pipeline.
//   - [MTL4MeshRenderPipelineDescriptor.SetFragmentStaticLinkingDescriptor]
//   - [MTL4MeshRenderPipelineDescriptor.RasterizationEnabled]: Determines whether the pipeline rasterizes primitives.
//   - [MTL4MeshRenderPipelineDescriptor.SetRasterizationEnabled]
//   - [MTL4MeshRenderPipelineDescriptor.MaxTotalThreadgroupsPerMeshGrid]: Controls the largest number of threads the pipeline state can execute when the object stage of a mesh render pipeline you create from this descriptor dispatches its mesh stage.
//   - [MTL4MeshRenderPipelineDescriptor.SetMaxTotalThreadgroupsPerMeshGrid]
//   - [MTL4MeshRenderPipelineDescriptor.MaxTotalThreadsPerMeshThreadgroup]: Controls the largest number of threads the pipeline state can execute in a single mesh shader threadgroup dispatch.
//   - [MTL4MeshRenderPipelineDescriptor.SetMaxTotalThreadsPerMeshThreadgroup]
//   - [MTL4MeshRenderPipelineDescriptor.MaxTotalThreadsPerObjectThreadgroup]: Controls the largest number of threads the pipeline state can execute in a single object shader threadgroup dispatch.
//   - [MTL4MeshRenderPipelineDescriptor.SetMaxTotalThreadsPerObjectThreadgroup]
//   - [MTL4MeshRenderPipelineDescriptor.MaxVertexAmplificationCount]: Determines the maximum value that can you can pass as the pipeline’s amplification count.
//   - [MTL4MeshRenderPipelineDescriptor.SetMaxVertexAmplificationCount]
//   - [MTL4MeshRenderPipelineDescriptor.MeshFunctionDescriptor]: Assigns a function descriptor representing the function this pipeline executes for each primitive in the mesh shader stage.
//   - [MTL4MeshRenderPipelineDescriptor.SetMeshFunctionDescriptor]
//   - [MTL4MeshRenderPipelineDescriptor.MeshStaticLinkingDescriptor]: Provides static linking information for the mesh stage of the render pipeline.
//   - [MTL4MeshRenderPipelineDescriptor.SetMeshStaticLinkingDescriptor]
//   - [MTL4MeshRenderPipelineDescriptor.MeshThreadgroupSizeIsMultipleOfThreadExecutionWidth]: Provides a guarantee to Metal regarding the number of threadgroup threads for the mesh stage of a pipeline you create from this descriptor.
//   - [MTL4MeshRenderPipelineDescriptor.SetMeshThreadgroupSizeIsMultipleOfThreadExecutionWidth]
//   - [MTL4MeshRenderPipelineDescriptor.ObjectFunctionDescriptor]: Assigns a function descriptor representing the function this pipeline executes for each  in the object shader stage.
//   - [MTL4MeshRenderPipelineDescriptor.SetObjectFunctionDescriptor]
//   - [MTL4MeshRenderPipelineDescriptor.ObjectStaticLinkingDescriptor]: Provides static linking information for the object stage of the render pipeline.
//   - [MTL4MeshRenderPipelineDescriptor.SetObjectStaticLinkingDescriptor]
//   - [MTL4MeshRenderPipelineDescriptor.ObjectThreadgroupSizeIsMultipleOfThreadExecutionWidth]: Provides a guarantee to Metal regarding the number of threadgroup threads for the object stage of a pipeline you create from this descriptor.
//   - [MTL4MeshRenderPipelineDescriptor.SetObjectThreadgroupSizeIsMultipleOfThreadExecutionWidth]
//   - [MTL4MeshRenderPipelineDescriptor.PayloadMemoryLength]: Reserves storage for the object-to-mesh stage payload.
//   - [MTL4MeshRenderPipelineDescriptor.SetPayloadMemoryLength]
//   - [MTL4MeshRenderPipelineDescriptor.RasterSampleCount]: Sets number of samples this pipeline applies for each fragment.
//   - [MTL4MeshRenderPipelineDescriptor.SetRasterSampleCount]
//   - [MTL4MeshRenderPipelineDescriptor.RequiredThreadsPerMeshThreadgroup]: Controls the required number of mesh threads-per-threadgroup when drawing with a mesh shader pipeline you create from this descriptor.
//   - [MTL4MeshRenderPipelineDescriptor.SetRequiredThreadsPerMeshThreadgroup]
//   - [MTL4MeshRenderPipelineDescriptor.RequiredThreadsPerObjectThreadgroup]: Controls the required number of object threads-per-threadgroup when drawing with a mesh shader pipeline you create from this descriptor.
//   - [MTL4MeshRenderPipelineDescriptor.SetRequiredThreadsPerObjectThreadgroup]
//   - [MTL4MeshRenderPipelineDescriptor.SupportFragmentBinaryLinking]: Indicates whether you can use the render pipeline to create new pipelines by adding binary functions to the fragment shader function’s callable functions list.
//   - [MTL4MeshRenderPipelineDescriptor.SetSupportFragmentBinaryLinking]
//   - [MTL4MeshRenderPipelineDescriptor.SupportIndirectCommandBuffers]: Indicates whether the pipeline supports indirect command buffers.
//   - [MTL4MeshRenderPipelineDescriptor.SetSupportIndirectCommandBuffers]
//   - [MTL4MeshRenderPipelineDescriptor.SupportMeshBinaryLinking]: Indicates whether you can use the render pipeline to create new pipelines by adding binary functions to the mesh shader function’s callable functions list.
//   - [MTL4MeshRenderPipelineDescriptor.SetSupportMeshBinaryLinking]
//   - [MTL4MeshRenderPipelineDescriptor.SupportObjectBinaryLinking]: Indicates whether you can use the render pipeline to create new pipelines by adding binary functions to the object shader function’s callable functions list.
//   - [MTL4MeshRenderPipelineDescriptor.SetSupportObjectBinaryLinking]
//
// # Instance Methods
//
//   - [MTL4MeshRenderPipelineDescriptor.Reset]: Resets this descriptor to its default state.
//
// See: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor
type MTL4MeshRenderPipelineDescriptor struct {
	MTL4PipelineDescriptor
}

// MTL4MeshRenderPipelineDescriptorFromID constructs a [MTL4MeshRenderPipelineDescriptor] from an objc.ID.
//
// Groups together properties you use to create a mesh render pipeline state
// object.
func MTL4MeshRenderPipelineDescriptorFromID(id objc.ID) MTL4MeshRenderPipelineDescriptor {
	return MTL4MeshRenderPipelineDescriptor{MTL4PipelineDescriptor: MTL4PipelineDescriptorFromID(id)}
}
// NOTE: MTL4MeshRenderPipelineDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTL4MeshRenderPipelineDescriptor] class.
//
// # Instance Properties
//
//   - [IMTL4MeshRenderPipelineDescriptor.AlphaToCoverageState]: Indicates whether to read and use the alpha channel fragment output of color attachments to compute a sample coverage mask.
//   - [IMTL4MeshRenderPipelineDescriptor.SetAlphaToCoverageState]
//   - [IMTL4MeshRenderPipelineDescriptor.AlphaToOneState]: Indicates whether the pipeline forces alpha channel values of color attachments to the largest representable value.
//   - [IMTL4MeshRenderPipelineDescriptor.SetAlphaToOneState]
//   - [IMTL4MeshRenderPipelineDescriptor.ColorAttachmentMappingState]: Sets the logical-to-physical rendering remap state.
//   - [IMTL4MeshRenderPipelineDescriptor.SetColorAttachmentMappingState]
//   - [IMTL4MeshRenderPipelineDescriptor.ColorAttachments]: Accesses an array containing descriptions of the color attachments this pipeline writes to.
//   - [IMTL4MeshRenderPipelineDescriptor.FragmentFunctionDescriptor]: Assigns a function descriptor representing the function this pipeline executes for each fragment.
//   - [IMTL4MeshRenderPipelineDescriptor.SetFragmentFunctionDescriptor]
//   - [IMTL4MeshRenderPipelineDescriptor.FragmentStaticLinkingDescriptor]: Provides static linking information for the fragment stage of the render pipeline.
//   - [IMTL4MeshRenderPipelineDescriptor.SetFragmentStaticLinkingDescriptor]
//   - [IMTL4MeshRenderPipelineDescriptor.RasterizationEnabled]: Determines whether the pipeline rasterizes primitives.
//   - [IMTL4MeshRenderPipelineDescriptor.SetRasterizationEnabled]
//   - [IMTL4MeshRenderPipelineDescriptor.MaxTotalThreadgroupsPerMeshGrid]: Controls the largest number of threads the pipeline state can execute when the object stage of a mesh render pipeline you create from this descriptor dispatches its mesh stage.
//   - [IMTL4MeshRenderPipelineDescriptor.SetMaxTotalThreadgroupsPerMeshGrid]
//   - [IMTL4MeshRenderPipelineDescriptor.MaxTotalThreadsPerMeshThreadgroup]: Controls the largest number of threads the pipeline state can execute in a single mesh shader threadgroup dispatch.
//   - [IMTL4MeshRenderPipelineDescriptor.SetMaxTotalThreadsPerMeshThreadgroup]
//   - [IMTL4MeshRenderPipelineDescriptor.MaxTotalThreadsPerObjectThreadgroup]: Controls the largest number of threads the pipeline state can execute in a single object shader threadgroup dispatch.
//   - [IMTL4MeshRenderPipelineDescriptor.SetMaxTotalThreadsPerObjectThreadgroup]
//   - [IMTL4MeshRenderPipelineDescriptor.MaxVertexAmplificationCount]: Determines the maximum value that can you can pass as the pipeline’s amplification count.
//   - [IMTL4MeshRenderPipelineDescriptor.SetMaxVertexAmplificationCount]
//   - [IMTL4MeshRenderPipelineDescriptor.MeshFunctionDescriptor]: Assigns a function descriptor representing the function this pipeline executes for each primitive in the mesh shader stage.
//   - [IMTL4MeshRenderPipelineDescriptor.SetMeshFunctionDescriptor]
//   - [IMTL4MeshRenderPipelineDescriptor.MeshStaticLinkingDescriptor]: Provides static linking information for the mesh stage of the render pipeline.
//   - [IMTL4MeshRenderPipelineDescriptor.SetMeshStaticLinkingDescriptor]
//   - [IMTL4MeshRenderPipelineDescriptor.MeshThreadgroupSizeIsMultipleOfThreadExecutionWidth]: Provides a guarantee to Metal regarding the number of threadgroup threads for the mesh stage of a pipeline you create from this descriptor.
//   - [IMTL4MeshRenderPipelineDescriptor.SetMeshThreadgroupSizeIsMultipleOfThreadExecutionWidth]
//   - [IMTL4MeshRenderPipelineDescriptor.ObjectFunctionDescriptor]: Assigns a function descriptor representing the function this pipeline executes for each  in the object shader stage.
//   - [IMTL4MeshRenderPipelineDescriptor.SetObjectFunctionDescriptor]
//   - [IMTL4MeshRenderPipelineDescriptor.ObjectStaticLinkingDescriptor]: Provides static linking information for the object stage of the render pipeline.
//   - [IMTL4MeshRenderPipelineDescriptor.SetObjectStaticLinkingDescriptor]
//   - [IMTL4MeshRenderPipelineDescriptor.ObjectThreadgroupSizeIsMultipleOfThreadExecutionWidth]: Provides a guarantee to Metal regarding the number of threadgroup threads for the object stage of a pipeline you create from this descriptor.
//   - [IMTL4MeshRenderPipelineDescriptor.SetObjectThreadgroupSizeIsMultipleOfThreadExecutionWidth]
//   - [IMTL4MeshRenderPipelineDescriptor.PayloadMemoryLength]: Reserves storage for the object-to-mesh stage payload.
//   - [IMTL4MeshRenderPipelineDescriptor.SetPayloadMemoryLength]
//   - [IMTL4MeshRenderPipelineDescriptor.RasterSampleCount]: Sets number of samples this pipeline applies for each fragment.
//   - [IMTL4MeshRenderPipelineDescriptor.SetRasterSampleCount]
//   - [IMTL4MeshRenderPipelineDescriptor.RequiredThreadsPerMeshThreadgroup]: Controls the required number of mesh threads-per-threadgroup when drawing with a mesh shader pipeline you create from this descriptor.
//   - [IMTL4MeshRenderPipelineDescriptor.SetRequiredThreadsPerMeshThreadgroup]
//   - [IMTL4MeshRenderPipelineDescriptor.RequiredThreadsPerObjectThreadgroup]: Controls the required number of object threads-per-threadgroup when drawing with a mesh shader pipeline you create from this descriptor.
//   - [IMTL4MeshRenderPipelineDescriptor.SetRequiredThreadsPerObjectThreadgroup]
//   - [IMTL4MeshRenderPipelineDescriptor.SupportFragmentBinaryLinking]: Indicates whether you can use the render pipeline to create new pipelines by adding binary functions to the fragment shader function’s callable functions list.
//   - [IMTL4MeshRenderPipelineDescriptor.SetSupportFragmentBinaryLinking]
//   - [IMTL4MeshRenderPipelineDescriptor.SupportIndirectCommandBuffers]: Indicates whether the pipeline supports indirect command buffers.
//   - [IMTL4MeshRenderPipelineDescriptor.SetSupportIndirectCommandBuffers]
//   - [IMTL4MeshRenderPipelineDescriptor.SupportMeshBinaryLinking]: Indicates whether you can use the render pipeline to create new pipelines by adding binary functions to the mesh shader function’s callable functions list.
//   - [IMTL4MeshRenderPipelineDescriptor.SetSupportMeshBinaryLinking]
//   - [IMTL4MeshRenderPipelineDescriptor.SupportObjectBinaryLinking]: Indicates whether you can use the render pipeline to create new pipelines by adding binary functions to the object shader function’s callable functions list.
//   - [IMTL4MeshRenderPipelineDescriptor.SetSupportObjectBinaryLinking]
//
// # Instance Methods
//
//   - [IMTL4MeshRenderPipelineDescriptor.Reset]: Resets this descriptor to its default state.
//
// See: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor
type IMTL4MeshRenderPipelineDescriptor interface {
	IMTL4PipelineDescriptor

	// Topic: Instance Properties

	// Indicates whether to read and use the alpha channel fragment output of color attachments to compute a sample coverage mask.
	AlphaToCoverageState() MTL4AlphaToCoverageState
	SetAlphaToCoverageState(value MTL4AlphaToCoverageState)
	// Indicates whether the pipeline forces alpha channel values of color attachments to the largest representable value.
	AlphaToOneState() MTL4AlphaToOneState
	SetAlphaToOneState(value MTL4AlphaToOneState)
	// Sets the logical-to-physical rendering remap state.
	ColorAttachmentMappingState() MTL4LogicalToPhysicalColorAttachmentMappingState
	SetColorAttachmentMappingState(value MTL4LogicalToPhysicalColorAttachmentMappingState)
	// Accesses an array containing descriptions of the color attachments this pipeline writes to.
	ColorAttachments() IMTL4RenderPipelineColorAttachmentDescriptorArray
	// Assigns a function descriptor representing the function this pipeline executes for each fragment.
	FragmentFunctionDescriptor() IMTL4FunctionDescriptor
	SetFragmentFunctionDescriptor(value IMTL4FunctionDescriptor)
	// Provides static linking information for the fragment stage of the render pipeline.
	FragmentStaticLinkingDescriptor() IMTL4StaticLinkingDescriptor
	SetFragmentStaticLinkingDescriptor(value IMTL4StaticLinkingDescriptor)
	// Determines whether the pipeline rasterizes primitives.
	RasterizationEnabled() bool
	SetRasterizationEnabled(value bool)
	// Controls the largest number of threads the pipeline state can execute when the object stage of a mesh render pipeline you create from this descriptor dispatches its mesh stage.
	MaxTotalThreadgroupsPerMeshGrid() uint
	SetMaxTotalThreadgroupsPerMeshGrid(value uint)
	// Controls the largest number of threads the pipeline state can execute in a single mesh shader threadgroup dispatch.
	MaxTotalThreadsPerMeshThreadgroup() uint
	SetMaxTotalThreadsPerMeshThreadgroup(value uint)
	// Controls the largest number of threads the pipeline state can execute in a single object shader threadgroup dispatch.
	MaxTotalThreadsPerObjectThreadgroup() uint
	SetMaxTotalThreadsPerObjectThreadgroup(value uint)
	// Determines the maximum value that can you can pass as the pipeline’s amplification count.
	MaxVertexAmplificationCount() uint
	SetMaxVertexAmplificationCount(value uint)
	// Assigns a function descriptor representing the function this pipeline executes for each primitive in the mesh shader stage.
	MeshFunctionDescriptor() IMTL4FunctionDescriptor
	SetMeshFunctionDescriptor(value IMTL4FunctionDescriptor)
	// Provides static linking information for the mesh stage of the render pipeline.
	MeshStaticLinkingDescriptor() IMTL4StaticLinkingDescriptor
	SetMeshStaticLinkingDescriptor(value IMTL4StaticLinkingDescriptor)
	// Provides a guarantee to Metal regarding the number of threadgroup threads for the mesh stage of a pipeline you create from this descriptor.
	MeshThreadgroupSizeIsMultipleOfThreadExecutionWidth() bool
	SetMeshThreadgroupSizeIsMultipleOfThreadExecutionWidth(value bool)
	// Assigns a function descriptor representing the function this pipeline executes for each  in the object shader stage.
	ObjectFunctionDescriptor() IMTL4FunctionDescriptor
	SetObjectFunctionDescriptor(value IMTL4FunctionDescriptor)
	// Provides static linking information for the object stage of the render pipeline.
	ObjectStaticLinkingDescriptor() IMTL4StaticLinkingDescriptor
	SetObjectStaticLinkingDescriptor(value IMTL4StaticLinkingDescriptor)
	// Provides a guarantee to Metal regarding the number of threadgroup threads for the object stage of a pipeline you create from this descriptor.
	ObjectThreadgroupSizeIsMultipleOfThreadExecutionWidth() bool
	SetObjectThreadgroupSizeIsMultipleOfThreadExecutionWidth(value bool)
	// Reserves storage for the object-to-mesh stage payload.
	PayloadMemoryLength() uint
	SetPayloadMemoryLength(value uint)
	// Sets number of samples this pipeline applies for each fragment.
	RasterSampleCount() uint
	SetRasterSampleCount(value uint)
	// Controls the required number of mesh threads-per-threadgroup when drawing with a mesh shader pipeline you create from this descriptor.
	RequiredThreadsPerMeshThreadgroup() MTLSize
	SetRequiredThreadsPerMeshThreadgroup(value MTLSize)
	// Controls the required number of object threads-per-threadgroup when drawing with a mesh shader pipeline you create from this descriptor.
	RequiredThreadsPerObjectThreadgroup() MTLSize
	SetRequiredThreadsPerObjectThreadgroup(value MTLSize)
	// Indicates whether you can use the render pipeline to create new pipelines by adding binary functions to the fragment shader function’s callable functions list.
	SupportFragmentBinaryLinking() bool
	SetSupportFragmentBinaryLinking(value bool)
	// Indicates whether the pipeline supports indirect command buffers.
	SupportIndirectCommandBuffers() MTL4IndirectCommandBufferSupportState
	SetSupportIndirectCommandBuffers(value MTL4IndirectCommandBufferSupportState)
	// Indicates whether you can use the render pipeline to create new pipelines by adding binary functions to the mesh shader function’s callable functions list.
	SupportMeshBinaryLinking() bool
	SetSupportMeshBinaryLinking(value bool)
	// Indicates whether you can use the render pipeline to create new pipelines by adding binary functions to the object shader function’s callable functions list.
	SupportObjectBinaryLinking() bool
	SetSupportObjectBinaryLinking(value bool)

	// Topic: Instance Methods

	// Resets this descriptor to its default state.
	Reset()
}

// Init initializes the instance.
func (m MTL4MeshRenderPipelineDescriptor) Init() MTL4MeshRenderPipelineDescriptor {
	rv := objc.Send[MTL4MeshRenderPipelineDescriptor](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTL4MeshRenderPipelineDescriptor) Autorelease() MTL4MeshRenderPipelineDescriptor {
	rv := objc.Send[MTL4MeshRenderPipelineDescriptor](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4MeshRenderPipelineDescriptor creates a new MTL4MeshRenderPipelineDescriptor instance.
func NewMTL4MeshRenderPipelineDescriptor() MTL4MeshRenderPipelineDescriptor {
	class := getMTL4MeshRenderPipelineDescriptorClass()
	rv := objc.Send[MTL4MeshRenderPipelineDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Resets this descriptor to its default state.
//
// See: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/reset()
func (m MTL4MeshRenderPipelineDescriptor) Reset() {
	objc.Send[objc.ID](m.ID, objc.Sel("reset"))
}

// Indicates whether to read and use the alpha channel fragment output of
// color attachments to compute a sample coverage mask.
//
// See: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/alphaToCoverageState
func (m MTL4MeshRenderPipelineDescriptor) AlphaToCoverageState() MTL4AlphaToCoverageState {
	rv := objc.Send[MTL4AlphaToCoverageState](m.ID, objc.Sel("alphaToCoverageState"))
	return MTL4AlphaToCoverageState(rv)
}
func (m MTL4MeshRenderPipelineDescriptor) SetAlphaToCoverageState(value MTL4AlphaToCoverageState) {
	objc.Send[struct{}](m.ID, objc.Sel("setAlphaToCoverageState:"), value)
}

// Indicates whether the pipeline forces alpha channel values of color
// attachments to the largest representable value.
//
// See: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/alphaToOneState
func (m MTL4MeshRenderPipelineDescriptor) AlphaToOneState() MTL4AlphaToOneState {
	rv := objc.Send[MTL4AlphaToOneState](m.ID, objc.Sel("alphaToOneState"))
	return MTL4AlphaToOneState(rv)
}
func (m MTL4MeshRenderPipelineDescriptor) SetAlphaToOneState(value MTL4AlphaToOneState) {
	objc.Send[struct{}](m.ID, objc.Sel("setAlphaToOneState:"), value)
}

// Sets the logical-to-physical rendering remap state.
//
// # Discussion
// 
// Use this property to assign how a [MTL4RenderCommandEncoder] instance maps
// the output of your fragment shader to physical color attachments.
//
// See: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/colorAttachmentMappingState
func (m MTL4MeshRenderPipelineDescriptor) ColorAttachmentMappingState() MTL4LogicalToPhysicalColorAttachmentMappingState {
	rv := objc.Send[MTL4LogicalToPhysicalColorAttachmentMappingState](m.ID, objc.Sel("colorAttachmentMappingState"))
	return MTL4LogicalToPhysicalColorAttachmentMappingState(rv)
}
func (m MTL4MeshRenderPipelineDescriptor) SetColorAttachmentMappingState(value MTL4LogicalToPhysicalColorAttachmentMappingState) {
	objc.Send[struct{}](m.ID, objc.Sel("setColorAttachmentMappingState:"), value)
}

// Accesses an array containing descriptions of the color attachments this
// pipeline writes to.
//
// See: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/colorAttachments
func (m MTL4MeshRenderPipelineDescriptor) ColorAttachments() IMTL4RenderPipelineColorAttachmentDescriptorArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("colorAttachments"))
	return MTL4RenderPipelineColorAttachmentDescriptorArrayFromID(objc.ID(rv))
}

// Assigns a function descriptor representing the function this pipeline
// executes for each fragment.
//
// See: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/fragmentFunctionDescriptor
func (m MTL4MeshRenderPipelineDescriptor) FragmentFunctionDescriptor() IMTL4FunctionDescriptor {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("fragmentFunctionDescriptor"))
	return MTL4FunctionDescriptorFromID(objc.ID(rv))
}
func (m MTL4MeshRenderPipelineDescriptor) SetFragmentFunctionDescriptor(value IMTL4FunctionDescriptor) {
	objc.Send[struct{}](m.ID, objc.Sel("setFragmentFunctionDescriptor:"), value)
}

// Provides static linking information for the fragment stage of the render
// pipeline.
//
// # Discussion
// 
// Use this property to link extra shader functions to the fragment stage of
// the render pipeline.
//
// See: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/fragmentStaticLinkingDescriptor
func (m MTL4MeshRenderPipelineDescriptor) FragmentStaticLinkingDescriptor() IMTL4StaticLinkingDescriptor {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("fragmentStaticLinkingDescriptor"))
	return MTL4StaticLinkingDescriptorFromID(objc.ID(rv))
}
func (m MTL4MeshRenderPipelineDescriptor) SetFragmentStaticLinkingDescriptor(value IMTL4StaticLinkingDescriptor) {
	objc.Send[struct{}](m.ID, objc.Sel("setFragmentStaticLinkingDescriptor:"), value)
}

// Determines whether the pipeline rasterizes primitives.
//
// # Discussion
// 
// By default, this value is [true], specifying that this pipeline rasterizes
// primitives. Set this property to [false] when you don’t provide a
// fragment shader function via function [FragmentFunctionDescriptor].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/isRasterizationEnabled
func (m MTL4MeshRenderPipelineDescriptor) RasterizationEnabled() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isRasterizationEnabled"))
	return rv
}
func (m MTL4MeshRenderPipelineDescriptor) SetRasterizationEnabled(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setRasterizationEnabled:"), value)
}

// Controls the largest number of threads the pipeline state can execute when
// the object stage of a mesh render pipeline you create from this descriptor
// dispatches its mesh stage.
//
// # Discussion
// 
// This number represents the maximum size of the product of the components of
// the parameter you pass to Metal Shading Language’s built-in function
// `:set_threadgroups_per_grid`.
// 
// The default value of this property is `0`, which indicates that the Metal
// Shading Language attribute `[[max_total_threadgroups_per_mesh_grid(N)]]`
// you attach to the pipeline’s mesh shader function determines the value of
// this property.
// 
// When you specify both the `[[max_total_threadgroups_per_mesh_grid(N)]]`
// attribute and this property, you are responsible for making sure these
// values match.
// 
// Additionally, you are responsible for ensuring this value doesn’t exceed
// the “maximum threads per mesh grid” device limit documented in the
// “Metal Feature Set Tables” PDF:
// [https://developer.apple.com/metal/Metal-Feature-Set-Tables.pdf].
//
// [https://developer.apple.com/metal/Metal-Feature-Set-Tables.pdf]: https://developer.apple.com/metal/Metal-Feature-Set-Tables.pdf
//
// See: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/maxTotalThreadgroupsPerMeshGrid
func (m MTL4MeshRenderPipelineDescriptor) MaxTotalThreadgroupsPerMeshGrid() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("maxTotalThreadgroupsPerMeshGrid"))
	return rv
}
func (m MTL4MeshRenderPipelineDescriptor) SetMaxTotalThreadgroupsPerMeshGrid(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setMaxTotalThreadgroupsPerMeshGrid:"), value)
}

// Controls the largest number of threads the pipeline state can execute in a
// single mesh shader threadgroup dispatch.
//
// # Discussion
// 
// This number represents the maximum size of the product of the components of
// parameter `threadsPerMeshThreadgroup` that Metal can use when drawing with
// this pipeline in mesh shader dispatch methods, such as
// [DrawMeshThreadgroupsThreadsPerObjectThreadgroupThreadsPerMeshThreadgroup].
// 
// The compiler’s optimizer can use the value of this property to generate
// more efficient code, specifically when the value doesn’t exceed the
// thread execution width of the underlying GPU.
// 
// The default value of this property is `0`, thish indicates that the Metal
// Shader Language attribute `[[max_total_threads_per_threadgroup]]` you
// attache to the pipeline’s mesh shader function determines the value of
// this property.
// 
// When you specify both the `[[max_total_threads_per_threadgroup(N)]]`
// attribute and this property, you are responsible for making sure these
// values match.
// 
// Additionally, you are responsible for ensuring this value doesn’t exceed
// the “maximum threads per threadgroup” device limit documented in the
// “Metal Feature Set Tables” PDF:
// [https://developer.apple.com/metal/Metal-Feature-Set-Tables.pdf].
//
// [https://developer.apple.com/metal/Metal-Feature-Set-Tables.pdf]: https://developer.apple.com/metal/Metal-Feature-Set-Tables.pdf
//
// See: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/maxTotalThreadsPerMeshThreadgroup
func (m MTL4MeshRenderPipelineDescriptor) MaxTotalThreadsPerMeshThreadgroup() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("maxTotalThreadsPerMeshThreadgroup"))
	return rv
}
func (m MTL4MeshRenderPipelineDescriptor) SetMaxTotalThreadsPerMeshThreadgroup(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setMaxTotalThreadsPerMeshThreadgroup:"), value)
}

// Controls the largest number of threads the pipeline state can execute in a
// single object shader threadgroup dispatch.
//
// # Discussion
// 
// This number represents the maximum size of the product of the components of
// parameter `threadsPerObjectThreadgroup` that Metal can use when drawing
// with this pipeline in mesh shader dispatch methods, such as
// [DrawMeshThreadgroupsThreadsPerObjectThreadgroupThreadsPerMeshThreadgroup].
// 
// The compiler’s optimizer can use the value of this property to generate
// more efficient code, specifically when the value doesn’t exceed the
// thread execution width of the underlying GPU.
// 
// The default value of this property is `0`, which indicates that the number
// you pass to attribute `[[max_total_threads_per_threadgroup(N)]]` of the
// pipeline’s object function determines the maximum total threads per
// threadgroup.
// 
// When you specify both the `[[max_total_threads_per_threadgroup(N)]]`
// attribute and this property, you are responsible for making sure these
// values match.
// 
// Additionally, you are responsible for ensuring this value doesn’t exceed
// the “maximum threads per threadgroup” device limit documented in the
// “Metal Feature Set Tables” PDF:
// [https://developer.apple.com/metal/Metal-Feature-Set-Tables.pdf].
//
// [https://developer.apple.com/metal/Metal-Feature-Set-Tables.pdf]: https://developer.apple.com/metal/Metal-Feature-Set-Tables.pdf
//
// See: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/maxTotalThreadsPerObjectThreadgroup
func (m MTL4MeshRenderPipelineDescriptor) MaxTotalThreadsPerObjectThreadgroup() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("maxTotalThreadsPerObjectThreadgroup"))
	return rv
}
func (m MTL4MeshRenderPipelineDescriptor) SetMaxTotalThreadsPerObjectThreadgroup(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setMaxTotalThreadsPerObjectThreadgroup:"), value)
}

// Determines the maximum value that can you can pass as the pipeline’s
// amplification count.
//
// # Discussion
// 
// This property controls the maximum count you pass to
// [SetVertexAmplificationCountViewMappings] when using vertex amplification
// with this pipeline.
//
// See: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/maxVertexAmplificationCount
func (m MTL4MeshRenderPipelineDescriptor) MaxVertexAmplificationCount() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("maxVertexAmplificationCount"))
	return rv
}
func (m MTL4MeshRenderPipelineDescriptor) SetMaxVertexAmplificationCount(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setMaxVertexAmplificationCount:"), value)
}

// Assigns a function descriptor representing the function this pipeline
// executes for each primitive in the mesh shader stage.
//
// See: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/meshFunctionDescriptor
func (m MTL4MeshRenderPipelineDescriptor) MeshFunctionDescriptor() IMTL4FunctionDescriptor {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("meshFunctionDescriptor"))
	return MTL4FunctionDescriptorFromID(objc.ID(rv))
}
func (m MTL4MeshRenderPipelineDescriptor) SetMeshFunctionDescriptor(value IMTL4FunctionDescriptor) {
	objc.Send[struct{}](m.ID, objc.Sel("setMeshFunctionDescriptor:"), value)
}

// Provides static linking information for the mesh stage of the render
// pipeline.
//
// # Discussion
// 
// Use this property to link extra shader functions to the mesh stage of the
// render pipeline.
//
// See: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/meshStaticLinkingDescriptor
func (m MTL4MeshRenderPipelineDescriptor) MeshStaticLinkingDescriptor() IMTL4StaticLinkingDescriptor {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("meshStaticLinkingDescriptor"))
	return MTL4StaticLinkingDescriptorFromID(objc.ID(rv))
}
func (m MTL4MeshRenderPipelineDescriptor) SetMeshStaticLinkingDescriptor(value IMTL4StaticLinkingDescriptor) {
	objc.Send[struct{}](m.ID, objc.Sel("setMeshStaticLinkingDescriptor:"), value)
}

// Provides a guarantee to Metal regarding the number of threadgroup threads
// for the mesh stage of a pipeline you create from this descriptor.
//
// # Discussion
// 
// If you set this property to [true], you state to Metal that when you use a
// mesh render pipeline you create from this descriptor, the number of
// threadgroup threads you dispatch for the mesh stage is a multiple of its
// [MeshThreadExecutionWidth]. The compiler’s optimizer can use this
// guarantee to generate more efficient code.
// 
// This property’s default value is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/meshThreadgroupSizeIsMultipleOfThreadExecutionWidth
func (m MTL4MeshRenderPipelineDescriptor) MeshThreadgroupSizeIsMultipleOfThreadExecutionWidth() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("meshThreadgroupSizeIsMultipleOfThreadExecutionWidth"))
	return rv
}
func (m MTL4MeshRenderPipelineDescriptor) SetMeshThreadgroupSizeIsMultipleOfThreadExecutionWidth(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setMeshThreadgroupSizeIsMultipleOfThreadExecutionWidth:"), value)
}

// Assigns a function descriptor representing the function this pipeline
// executes for each in the object shader stage.
//
// See: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/objectFunctionDescriptor
func (m MTL4MeshRenderPipelineDescriptor) ObjectFunctionDescriptor() IMTL4FunctionDescriptor {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("objectFunctionDescriptor"))
	return MTL4FunctionDescriptorFromID(objc.ID(rv))
}
func (m MTL4MeshRenderPipelineDescriptor) SetObjectFunctionDescriptor(value IMTL4FunctionDescriptor) {
	objc.Send[struct{}](m.ID, objc.Sel("setObjectFunctionDescriptor:"), value)
}

// Provides static linking information for the object stage of the render
// pipeline.
//
// # Discussion
// 
// Use this property to link extra shader functions to the object stage of the
// render pipeline.
//
// See: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/objectStaticLinkingDescriptor
func (m MTL4MeshRenderPipelineDescriptor) ObjectStaticLinkingDescriptor() IMTL4StaticLinkingDescriptor {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("objectStaticLinkingDescriptor"))
	return MTL4StaticLinkingDescriptorFromID(objc.ID(rv))
}
func (m MTL4MeshRenderPipelineDescriptor) SetObjectStaticLinkingDescriptor(value IMTL4StaticLinkingDescriptor) {
	objc.Send[struct{}](m.ID, objc.Sel("setObjectStaticLinkingDescriptor:"), value)
}

// Provides a guarantee to Metal regarding the number of threadgroup threads
// for the object stage of a pipeline you create from this descriptor.
//
// # Discussion
// 
// If you set this property to [true], you state to Metal that when you use a
// mesh render pipeline you create from this descriptor, the number of
// threadgroup threads you dispatch for the object stage is a multiple of its
// [ObjectThreadExecutionWidth]. The compiler’s optimizer can use this
// guarantee to generate more efficient code.
// 
// This property’s default value is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/objectThreadgroupSizeIsMultipleOfThreadExecutionWidth
func (m MTL4MeshRenderPipelineDescriptor) ObjectThreadgroupSizeIsMultipleOfThreadExecutionWidth() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("objectThreadgroupSizeIsMultipleOfThreadExecutionWidth"))
	return rv
}
func (m MTL4MeshRenderPipelineDescriptor) SetObjectThreadgroupSizeIsMultipleOfThreadExecutionWidth(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setObjectThreadgroupSizeIsMultipleOfThreadExecutionWidth:"), value)
}

// Reserves storage for the object-to-mesh stage payload.
//
// # Discussion
// 
// This property determines the size, in bytes, of the buffer you indicate via
// the Metal Shading Language `[[payload]]` attribute in the object and mesh
// shader functions of the mesh render pipeline.
// 
// If this value is `0`, Metal derives the size from the (dereferenced) type
// you declare for the payload in the object shader function. If the type is a
// pointer, Metal reserves space for a single element.
// 
// The default value is `0`.
//
// See: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/payloadMemoryLength
func (m MTL4MeshRenderPipelineDescriptor) PayloadMemoryLength() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("payloadMemoryLength"))
	return rv
}
func (m MTL4MeshRenderPipelineDescriptor) SetPayloadMemoryLength(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setPayloadMemoryLength:"), value)
}

// Sets number of samples this pipeline applies for each fragment.
//
// See: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/rasterSampleCount
func (m MTL4MeshRenderPipelineDescriptor) RasterSampleCount() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("rasterSampleCount"))
	return rv
}
func (m MTL4MeshRenderPipelineDescriptor) SetRasterSampleCount(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setRasterSampleCount:"), value)
}

// Controls the required number of mesh threads-per-threadgroup when drawing
// with a mesh shader pipeline you create from this descriptor.
//
// # Discussion
// 
// This argument is optional, unless this pipeline uses [CooperativeTensors],
// in which case you are responsible for providing it.
// 
// When this value is set to non-zero, you are responsible for ensuring the
// parameter `threadsPerMeshThreadgroup` in any mesh dispatch draw calls that
// use this mesh render pipeline, such as
// [DrawMeshThreadgroupsThreadsPerObjectThreadgroupThreadsPerMeshThreadgroup],
// match it.
// 
// Setting this value to a size of 0 in every dimension disables this
// property.
//
// See: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/requiredThreadsPerMeshThreadgroup
func (m MTL4MeshRenderPipelineDescriptor) RequiredThreadsPerMeshThreadgroup() MTLSize {
	rv := objc.Send[MTLSize](m.ID, objc.Sel("requiredThreadsPerMeshThreadgroup"))
	return MTLSize(rv)
}
func (m MTL4MeshRenderPipelineDescriptor) SetRequiredThreadsPerMeshThreadgroup(value MTLSize) {
	objc.Send[struct{}](m.ID, objc.Sel("setRequiredThreadsPerMeshThreadgroup:"), value)
}

// Controls the required number of object threads-per-threadgroup when drawing
// with a mesh shader pipeline you create from this descriptor.
//
// # Discussion
// 
// This argument is optional, unless this pipeline uses [CooperativeTensors],
// in which case you are responsible for providing it.
// 
// When this value is set to non-zero, you are responsible for ensuring the
// parameter `threadsPerObjectThreadgroup` in any mesh dispatch draw calls
// that use this mesh render pipeline, such as
// [DrawMeshThreadgroupsThreadsPerObjectThreadgroupThreadsPerMeshThreadgroup],
// match it.
// 
// Setting this value to a size of 0 in every dimension disables this
// property.
//
// See: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/requiredThreadsPerObjectThreadgroup
func (m MTL4MeshRenderPipelineDescriptor) RequiredThreadsPerObjectThreadgroup() MTLSize {
	rv := objc.Send[MTLSize](m.ID, objc.Sel("requiredThreadsPerObjectThreadgroup"))
	return MTLSize(rv)
}
func (m MTL4MeshRenderPipelineDescriptor) SetRequiredThreadsPerObjectThreadgroup(value MTLSize) {
	objc.Send[struct{}](m.ID, objc.Sel("setRequiredThreadsPerObjectThreadgroup:"), value)
}

// Indicates whether you can use the render pipeline to create new pipelines
// by adding binary functions to the fragment shader function’s callable
// functions list.
//
// See: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/supportFragmentBinaryLinking
func (m MTL4MeshRenderPipelineDescriptor) SupportFragmentBinaryLinking() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("supportFragmentBinaryLinking"))
	return rv
}
func (m MTL4MeshRenderPipelineDescriptor) SetSupportFragmentBinaryLinking(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setSupportFragmentBinaryLinking:"), value)
}

// Indicates whether the pipeline supports indirect command buffers.
//
// See: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/supportIndirectCommandBuffers
func (m MTL4MeshRenderPipelineDescriptor) SupportIndirectCommandBuffers() MTL4IndirectCommandBufferSupportState {
	rv := objc.Send[MTL4IndirectCommandBufferSupportState](m.ID, objc.Sel("supportIndirectCommandBuffers"))
	return MTL4IndirectCommandBufferSupportState(rv)
}
func (m MTL4MeshRenderPipelineDescriptor) SetSupportIndirectCommandBuffers(value MTL4IndirectCommandBufferSupportState) {
	objc.Send[struct{}](m.ID, objc.Sel("setSupportIndirectCommandBuffers:"), value)
}

// Indicates whether you can use the render pipeline to create new pipelines
// by adding binary functions to the mesh shader function’s callable
// functions list.
//
// See: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/supportMeshBinaryLinking
func (m MTL4MeshRenderPipelineDescriptor) SupportMeshBinaryLinking() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("supportMeshBinaryLinking"))
	return rv
}
func (m MTL4MeshRenderPipelineDescriptor) SetSupportMeshBinaryLinking(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setSupportMeshBinaryLinking:"), value)
}

// Indicates whether you can use the render pipeline to create new pipelines
// by adding binary functions to the object shader function’s callable
// functions list.
//
// See: https://developer.apple.com/documentation/Metal/MTL4MeshRenderPipelineDescriptor/supportObjectBinaryLinking
func (m MTL4MeshRenderPipelineDescriptor) SupportObjectBinaryLinking() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("supportObjectBinaryLinking"))
	return rv
}
func (m MTL4MeshRenderPipelineDescriptor) SetSupportObjectBinaryLinking(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setSupportObjectBinaryLinking:"), value)
}

