// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MTL4RenderPipelineDescriptor] class.
var (
	_MTL4RenderPipelineDescriptorClass     MTL4RenderPipelineDescriptorClass
	_MTL4RenderPipelineDescriptorClassOnce sync.Once
)

func getMTL4RenderPipelineDescriptorClass() MTL4RenderPipelineDescriptorClass {
	_MTL4RenderPipelineDescriptorClassOnce.Do(func() {
		_MTL4RenderPipelineDescriptorClass = MTL4RenderPipelineDescriptorClass{class: objc.GetClass("MTL4RenderPipelineDescriptor")}
	})
	return _MTL4RenderPipelineDescriptorClass
}

// GetMTL4RenderPipelineDescriptorClass returns the class object for MTL4RenderPipelineDescriptor.
func GetMTL4RenderPipelineDescriptorClass() MTL4RenderPipelineDescriptorClass {
	return getMTL4RenderPipelineDescriptorClass()
}

type MTL4RenderPipelineDescriptorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MTL4RenderPipelineDescriptorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTL4RenderPipelineDescriptorClass) Alloc() MTL4RenderPipelineDescriptor {
	rv := objc.Send[MTL4RenderPipelineDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// Groups together properties to create a render pipeline state object.
//
// # Overview
// 
// Compared to [MTLRenderPipelineDescriptor], this interface doesn’t offer a
// mechanism to hint to Metal mutability of vertex and fragment buffers.
// Additionally, using this descriptor, you don’t specify binary archives.
//
// # Instance Properties
//
//   - [MTL4RenderPipelineDescriptor.AlphaToCoverageState]: Indicates whether to read and use the alpha channel fragment output of color attachments to compute a sample coverage mask.
//   - [MTL4RenderPipelineDescriptor.SetAlphaToCoverageState]
//   - [MTL4RenderPipelineDescriptor.AlphaToOneState]: Indicates whether the pipeline forces alpha channel values of color attachments to the largest representable value.
//   - [MTL4RenderPipelineDescriptor.SetAlphaToOneState]
//   - [MTL4RenderPipelineDescriptor.ColorAttachmentMappingState]: Configures a logical-to-physical rendering remap state.
//   - [MTL4RenderPipelineDescriptor.SetColorAttachmentMappingState]
//   - [MTL4RenderPipelineDescriptor.ColorAttachments]: Accesses an array containing descriptions of the color attachments this pipeline writes to.
//   - [MTL4RenderPipelineDescriptor.FragmentFunctionDescriptor]: Assigns the shader function that this pipeline executes for each fragment.
//   - [MTL4RenderPipelineDescriptor.SetFragmentFunctionDescriptor]
//   - [MTL4RenderPipelineDescriptor.FragmentStaticLinkingDescriptor]: Provides static linking information for the fragment stage of the render pipeline.
//   - [MTL4RenderPipelineDescriptor.SetFragmentStaticLinkingDescriptor]
//   - [MTL4RenderPipelineDescriptor.InputPrimitiveTopology]: Assigns type of primitive topology this pipeline renders.
//   - [MTL4RenderPipelineDescriptor.SetInputPrimitiveTopology]
//   - [MTL4RenderPipelineDescriptor.RasterizationEnabled]: Determines whether the pipeline rasterizes primitives.
//   - [MTL4RenderPipelineDescriptor.SetRasterizationEnabled]
//   - [MTL4RenderPipelineDescriptor.MaxVertexAmplificationCount]: Determines the maximum value that can you can pass as the pipeline’s amplification count.
//   - [MTL4RenderPipelineDescriptor.SetMaxVertexAmplificationCount]
//   - [MTL4RenderPipelineDescriptor.RasterSampleCount]: Controls the number of samples this pipeline applies for each fragment.
//   - [MTL4RenderPipelineDescriptor.SetRasterSampleCount]
//   - [MTL4RenderPipelineDescriptor.SupportFragmentBinaryLinking]: Indicates whether you can use the pipeline to create new pipelines by adding binary functions to the fragment shader function’s callable functions list.
//   - [MTL4RenderPipelineDescriptor.SetSupportFragmentBinaryLinking]
//   - [MTL4RenderPipelineDescriptor.SupportIndirectCommandBuffers]: Indicates whether the pipeline supports indirect command buffers.
//   - [MTL4RenderPipelineDescriptor.SetSupportIndirectCommandBuffers]
//   - [MTL4RenderPipelineDescriptor.SupportVertexBinaryLinking]: Indicates whether you can use the render pipeline to create new pipelines by adding binary functions to the vertex shader function’s callable functions list.
//   - [MTL4RenderPipelineDescriptor.SetSupportVertexBinaryLinking]
//   - [MTL4RenderPipelineDescriptor.VertexDescriptor]: Configures an optional vertex descriptor for the vertex input.
//   - [MTL4RenderPipelineDescriptor.SetVertexDescriptor]
//   - [MTL4RenderPipelineDescriptor.VertexFunctionDescriptor]: Assigns the shader function that this pipeline executes for each vertex.
//   - [MTL4RenderPipelineDescriptor.SetVertexFunctionDescriptor]
//   - [MTL4RenderPipelineDescriptor.VertexStaticLinkingDescriptor]: Provides static linking information for the vertex stage of the render pipeline.
//   - [MTL4RenderPipelineDescriptor.SetVertexStaticLinkingDescriptor]
//
// # Instance Methods
//
//   - [MTL4RenderPipelineDescriptor.Reset]: Resets this descriptor to its default state.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor
type MTL4RenderPipelineDescriptor struct {
	MTL4PipelineDescriptor
}

// MTL4RenderPipelineDescriptorFromID constructs a [MTL4RenderPipelineDescriptor] from an objc.ID.
//
// Groups together properties to create a render pipeline state object.
func MTL4RenderPipelineDescriptorFromID(id objc.ID) MTL4RenderPipelineDescriptor {
	return MTL4RenderPipelineDescriptor{MTL4PipelineDescriptor: MTL4PipelineDescriptorFromID(id)}
}
// NOTE: MTL4RenderPipelineDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTL4RenderPipelineDescriptor] class.
//
// # Instance Properties
//
//   - [IMTL4RenderPipelineDescriptor.AlphaToCoverageState]: Indicates whether to read and use the alpha channel fragment output of color attachments to compute a sample coverage mask.
//   - [IMTL4RenderPipelineDescriptor.SetAlphaToCoverageState]
//   - [IMTL4RenderPipelineDescriptor.AlphaToOneState]: Indicates whether the pipeline forces alpha channel values of color attachments to the largest representable value.
//   - [IMTL4RenderPipelineDescriptor.SetAlphaToOneState]
//   - [IMTL4RenderPipelineDescriptor.ColorAttachmentMappingState]: Configures a logical-to-physical rendering remap state.
//   - [IMTL4RenderPipelineDescriptor.SetColorAttachmentMappingState]
//   - [IMTL4RenderPipelineDescriptor.ColorAttachments]: Accesses an array containing descriptions of the color attachments this pipeline writes to.
//   - [IMTL4RenderPipelineDescriptor.FragmentFunctionDescriptor]: Assigns the shader function that this pipeline executes for each fragment.
//   - [IMTL4RenderPipelineDescriptor.SetFragmentFunctionDescriptor]
//   - [IMTL4RenderPipelineDescriptor.FragmentStaticLinkingDescriptor]: Provides static linking information for the fragment stage of the render pipeline.
//   - [IMTL4RenderPipelineDescriptor.SetFragmentStaticLinkingDescriptor]
//   - [IMTL4RenderPipelineDescriptor.InputPrimitiveTopology]: Assigns type of primitive topology this pipeline renders.
//   - [IMTL4RenderPipelineDescriptor.SetInputPrimitiveTopology]
//   - [IMTL4RenderPipelineDescriptor.RasterizationEnabled]: Determines whether the pipeline rasterizes primitives.
//   - [IMTL4RenderPipelineDescriptor.SetRasterizationEnabled]
//   - [IMTL4RenderPipelineDescriptor.MaxVertexAmplificationCount]: Determines the maximum value that can you can pass as the pipeline’s amplification count.
//   - [IMTL4RenderPipelineDescriptor.SetMaxVertexAmplificationCount]
//   - [IMTL4RenderPipelineDescriptor.RasterSampleCount]: Controls the number of samples this pipeline applies for each fragment.
//   - [IMTL4RenderPipelineDescriptor.SetRasterSampleCount]
//   - [IMTL4RenderPipelineDescriptor.SupportFragmentBinaryLinking]: Indicates whether you can use the pipeline to create new pipelines by adding binary functions to the fragment shader function’s callable functions list.
//   - [IMTL4RenderPipelineDescriptor.SetSupportFragmentBinaryLinking]
//   - [IMTL4RenderPipelineDescriptor.SupportIndirectCommandBuffers]: Indicates whether the pipeline supports indirect command buffers.
//   - [IMTL4RenderPipelineDescriptor.SetSupportIndirectCommandBuffers]
//   - [IMTL4RenderPipelineDescriptor.SupportVertexBinaryLinking]: Indicates whether you can use the render pipeline to create new pipelines by adding binary functions to the vertex shader function’s callable functions list.
//   - [IMTL4RenderPipelineDescriptor.SetSupportVertexBinaryLinking]
//   - [IMTL4RenderPipelineDescriptor.VertexDescriptor]: Configures an optional vertex descriptor for the vertex input.
//   - [IMTL4RenderPipelineDescriptor.SetVertexDescriptor]
//   - [IMTL4RenderPipelineDescriptor.VertexFunctionDescriptor]: Assigns the shader function that this pipeline executes for each vertex.
//   - [IMTL4RenderPipelineDescriptor.SetVertexFunctionDescriptor]
//   - [IMTL4RenderPipelineDescriptor.VertexStaticLinkingDescriptor]: Provides static linking information for the vertex stage of the render pipeline.
//   - [IMTL4RenderPipelineDescriptor.SetVertexStaticLinkingDescriptor]
//
// # Instance Methods
//
//   - [IMTL4RenderPipelineDescriptor.Reset]: Resets this descriptor to its default state.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor
type IMTL4RenderPipelineDescriptor interface {
	IMTL4PipelineDescriptor

	// Topic: Instance Properties

	// Indicates whether to read and use the alpha channel fragment output of color attachments to compute a sample coverage mask.
	AlphaToCoverageState() MTL4AlphaToCoverageState
	SetAlphaToCoverageState(value MTL4AlphaToCoverageState)
	// Indicates whether the pipeline forces alpha channel values of color attachments to the largest representable value.
	AlphaToOneState() MTL4AlphaToOneState
	SetAlphaToOneState(value MTL4AlphaToOneState)
	// Configures a logical-to-physical rendering remap state.
	ColorAttachmentMappingState() MTL4LogicalToPhysicalColorAttachmentMappingState
	SetColorAttachmentMappingState(value MTL4LogicalToPhysicalColorAttachmentMappingState)
	// Accesses an array containing descriptions of the color attachments this pipeline writes to.
	ColorAttachments() IMTL4RenderPipelineColorAttachmentDescriptorArray
	// Assigns the shader function that this pipeline executes for each fragment.
	FragmentFunctionDescriptor() IMTL4FunctionDescriptor
	SetFragmentFunctionDescriptor(value IMTL4FunctionDescriptor)
	// Provides static linking information for the fragment stage of the render pipeline.
	FragmentStaticLinkingDescriptor() IMTL4StaticLinkingDescriptor
	SetFragmentStaticLinkingDescriptor(value IMTL4StaticLinkingDescriptor)
	// Assigns type of primitive topology this pipeline renders.
	InputPrimitiveTopology() MTLPrimitiveTopologyClass
	SetInputPrimitiveTopology(value MTLPrimitiveTopologyClass)
	// Determines whether the pipeline rasterizes primitives.
	RasterizationEnabled() bool
	SetRasterizationEnabled(value bool)
	// Determines the maximum value that can you can pass as the pipeline’s amplification count.
	MaxVertexAmplificationCount() uint
	SetMaxVertexAmplificationCount(value uint)
	// Controls the number of samples this pipeline applies for each fragment.
	RasterSampleCount() uint
	SetRasterSampleCount(value uint)
	// Indicates whether you can use the pipeline to create new pipelines by adding binary functions to the fragment shader function’s callable functions list.
	SupportFragmentBinaryLinking() bool
	SetSupportFragmentBinaryLinking(value bool)
	// Indicates whether the pipeline supports indirect command buffers.
	SupportIndirectCommandBuffers() MTL4IndirectCommandBufferSupportState
	SetSupportIndirectCommandBuffers(value MTL4IndirectCommandBufferSupportState)
	// Indicates whether you can use the render pipeline to create new pipelines by adding binary functions to the vertex shader function’s callable functions list.
	SupportVertexBinaryLinking() bool
	SetSupportVertexBinaryLinking(value bool)
	// Configures an optional vertex descriptor for the vertex input.
	VertexDescriptor() IMTLVertexDescriptor
	SetVertexDescriptor(value IMTLVertexDescriptor)
	// Assigns the shader function that this pipeline executes for each vertex.
	VertexFunctionDescriptor() IMTL4FunctionDescriptor
	SetVertexFunctionDescriptor(value IMTL4FunctionDescriptor)
	// Provides static linking information for the vertex stage of the render pipeline.
	VertexStaticLinkingDescriptor() IMTL4StaticLinkingDescriptor
	SetVertexStaticLinkingDescriptor(value IMTL4StaticLinkingDescriptor)

	// Topic: Instance Methods

	// Resets this descriptor to its default state.
	Reset()
}

// Init initializes the instance.
func (m MTL4RenderPipelineDescriptor) Init() MTL4RenderPipelineDescriptor {
	rv := objc.Send[MTL4RenderPipelineDescriptor](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTL4RenderPipelineDescriptor) Autorelease() MTL4RenderPipelineDescriptor {
	rv := objc.Send[MTL4RenderPipelineDescriptor](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTL4RenderPipelineDescriptor creates a new MTL4RenderPipelineDescriptor instance.
func NewMTL4RenderPipelineDescriptor() MTL4RenderPipelineDescriptor {
	class := getMTL4RenderPipelineDescriptorClass()
	rv := objc.Send[MTL4RenderPipelineDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Resets this descriptor to its default state.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor/reset()
func (m MTL4RenderPipelineDescriptor) Reset() {
	objc.Send[objc.ID](m.ID, objc.Sel("reset"))
}

// Indicates whether to read and use the alpha channel fragment output of
// color attachments to compute a sample coverage mask.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor/alphaToCoverageState
func (m MTL4RenderPipelineDescriptor) AlphaToCoverageState() MTL4AlphaToCoverageState {
	rv := objc.Send[MTL4AlphaToCoverageState](m.ID, objc.Sel("alphaToCoverageState"))
	return MTL4AlphaToCoverageState(rv)
}
func (m MTL4RenderPipelineDescriptor) SetAlphaToCoverageState(value MTL4AlphaToCoverageState) {
	objc.Send[struct{}](m.ID, objc.Sel("setAlphaToCoverageState:"), value)
}
// Indicates whether the pipeline forces alpha channel values of color
// attachments to the largest representable value.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor/alphaToOneState
func (m MTL4RenderPipelineDescriptor) AlphaToOneState() MTL4AlphaToOneState {
	rv := objc.Send[MTL4AlphaToOneState](m.ID, objc.Sel("alphaToOneState"))
	return MTL4AlphaToOneState(rv)
}
func (m MTL4RenderPipelineDescriptor) SetAlphaToOneState(value MTL4AlphaToOneState) {
	objc.Send[struct{}](m.ID, objc.Sel("setAlphaToOneState:"), value)
}
// Configures a logical-to-physical rendering remap state.
//
// # Discussion
// 
// Use this property to assign how a [MTL4RenderCommandEncoder] instance maps
// the output of your fragment shader to physical color attachments.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor/colorAttachmentMappingState
func (m MTL4RenderPipelineDescriptor) ColorAttachmentMappingState() MTL4LogicalToPhysicalColorAttachmentMappingState {
	rv := objc.Send[MTL4LogicalToPhysicalColorAttachmentMappingState](m.ID, objc.Sel("colorAttachmentMappingState"))
	return MTL4LogicalToPhysicalColorAttachmentMappingState(rv)
}
func (m MTL4RenderPipelineDescriptor) SetColorAttachmentMappingState(value MTL4LogicalToPhysicalColorAttachmentMappingState) {
	objc.Send[struct{}](m.ID, objc.Sel("setColorAttachmentMappingState:"), value)
}
// Accesses an array containing descriptions of the color attachments this
// pipeline writes to.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor/colorAttachments
func (m MTL4RenderPipelineDescriptor) ColorAttachments() IMTL4RenderPipelineColorAttachmentDescriptorArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("colorAttachments"))
	return MTL4RenderPipelineColorAttachmentDescriptorArrayFromID(objc.ID(rv))
}
// Assigns the shader function that this pipeline executes for each fragment.
//
// # Discussion
// 
// When you don’t specify a fragment function, you need to disable
// rasterization by setting property [RasterizationEnabled] to false.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor/fragmentFunctionDescriptor
func (m MTL4RenderPipelineDescriptor) FragmentFunctionDescriptor() IMTL4FunctionDescriptor {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("fragmentFunctionDescriptor"))
	return MTL4FunctionDescriptorFromID(objc.ID(rv))
}
func (m MTL4RenderPipelineDescriptor) SetFragmentFunctionDescriptor(value IMTL4FunctionDescriptor) {
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
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor/fragmentStaticLinkingDescriptor
func (m MTL4RenderPipelineDescriptor) FragmentStaticLinkingDescriptor() IMTL4StaticLinkingDescriptor {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("fragmentStaticLinkingDescriptor"))
	return MTL4StaticLinkingDescriptorFromID(objc.ID(rv))
}
func (m MTL4RenderPipelineDescriptor) SetFragmentStaticLinkingDescriptor(value IMTL4StaticLinkingDescriptor) {
	objc.Send[struct{}](m.ID, objc.Sel("setFragmentStaticLinkingDescriptor:"), value)
}
// Assigns type of primitive topology this pipeline renders.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor/inputPrimitiveTopology
func (m MTL4RenderPipelineDescriptor) InputPrimitiveTopology() MTLPrimitiveTopologyClass {
	rv := objc.Send[MTLPrimitiveTopologyClass](m.ID, objc.Sel("inputPrimitiveTopology"))
	return MTLPrimitiveTopologyClass(rv)
}
func (m MTL4RenderPipelineDescriptor) SetInputPrimitiveTopology(value MTLPrimitiveTopologyClass) {
	objc.Send[struct{}](m.ID, objc.Sel("setInputPrimitiveTopology:"), value)
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
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor/isRasterizationEnabled
func (m MTL4RenderPipelineDescriptor) RasterizationEnabled() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isRasterizationEnabled"))
	return rv
}
func (m MTL4RenderPipelineDescriptor) SetRasterizationEnabled(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setRasterizationEnabled:"), value)
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
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor/maxVertexAmplificationCount
func (m MTL4RenderPipelineDescriptor) MaxVertexAmplificationCount() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("maxVertexAmplificationCount"))
	return rv
}
func (m MTL4RenderPipelineDescriptor) SetMaxVertexAmplificationCount(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setMaxVertexAmplificationCount:"), value)
}
// Controls the number of samples this pipeline applies for each fragment.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor/rasterSampleCount
func (m MTL4RenderPipelineDescriptor) RasterSampleCount() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("rasterSampleCount"))
	return rv
}
func (m MTL4RenderPipelineDescriptor) SetRasterSampleCount(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setRasterSampleCount:"), value)
}
// Indicates whether you can use the pipeline to create new pipelines by
// adding binary functions to the fragment shader function’s callable
// functions list.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor/supportFragmentBinaryLinking
func (m MTL4RenderPipelineDescriptor) SupportFragmentBinaryLinking() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("supportFragmentBinaryLinking"))
	return rv
}
func (m MTL4RenderPipelineDescriptor) SetSupportFragmentBinaryLinking(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setSupportFragmentBinaryLinking:"), value)
}
// Indicates whether the pipeline supports indirect command buffers.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor/supportIndirectCommandBuffers
func (m MTL4RenderPipelineDescriptor) SupportIndirectCommandBuffers() MTL4IndirectCommandBufferSupportState {
	rv := objc.Send[MTL4IndirectCommandBufferSupportState](m.ID, objc.Sel("supportIndirectCommandBuffers"))
	return MTL4IndirectCommandBufferSupportState(rv)
}
func (m MTL4RenderPipelineDescriptor) SetSupportIndirectCommandBuffers(value MTL4IndirectCommandBufferSupportState) {
	objc.Send[struct{}](m.ID, objc.Sel("setSupportIndirectCommandBuffers:"), value)
}
// Indicates whether you can use the render pipeline to create new pipelines
// by adding binary functions to the vertex shader function’s callable
// functions list.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor/supportVertexBinaryLinking
func (m MTL4RenderPipelineDescriptor) SupportVertexBinaryLinking() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("supportVertexBinaryLinking"))
	return rv
}
func (m MTL4RenderPipelineDescriptor) SetSupportVertexBinaryLinking(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setSupportVertexBinaryLinking:"), value)
}
// Configures an optional vertex descriptor for the vertex input.
//
// # Discussion
// 
// A vertex descriptor specifies the layout of your vertex data, allowing your
// vertex shaders to access the content in your vertex arrays via the
// `[[stage_in]]` attribute in Metal Shading Language.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor/vertexDescriptor
func (m MTL4RenderPipelineDescriptor) VertexDescriptor() IMTLVertexDescriptor {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("vertexDescriptor"))
	return MTLVertexDescriptorFromID(objc.ID(rv))
}
func (m MTL4RenderPipelineDescriptor) SetVertexDescriptor(value IMTLVertexDescriptor) {
	objc.Send[struct{}](m.ID, objc.Sel("setVertexDescriptor:"), value)
}
// Assigns the shader function that this pipeline executes for each vertex.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor/vertexFunctionDescriptor
func (m MTL4RenderPipelineDescriptor) VertexFunctionDescriptor() IMTL4FunctionDescriptor {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("vertexFunctionDescriptor"))
	return MTL4FunctionDescriptorFromID(objc.ID(rv))
}
func (m MTL4RenderPipelineDescriptor) SetVertexFunctionDescriptor(value IMTL4FunctionDescriptor) {
	objc.Send[struct{}](m.ID, objc.Sel("setVertexFunctionDescriptor:"), value)
}
// Provides static linking information for the vertex stage of the render
// pipeline.
//
// # Discussion
// 
// Use this property to link extra shader functions to the vertex stage of the
// render pipeline.
//
// See: https://developer.apple.com/documentation/Metal/MTL4RenderPipelineDescriptor/vertexStaticLinkingDescriptor
func (m MTL4RenderPipelineDescriptor) VertexStaticLinkingDescriptor() IMTL4StaticLinkingDescriptor {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("vertexStaticLinkingDescriptor"))
	return MTL4StaticLinkingDescriptorFromID(objc.ID(rv))
}
func (m MTL4RenderPipelineDescriptor) SetVertexStaticLinkingDescriptor(value IMTL4StaticLinkingDescriptor) {
	objc.Send[struct{}](m.ID, objc.Sel("setVertexStaticLinkingDescriptor:"), value)
}

