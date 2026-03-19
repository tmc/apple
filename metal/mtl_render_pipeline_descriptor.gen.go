// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLRenderPipelineDescriptor] class.
var (
	_MTLRenderPipelineDescriptorClass     MTLRenderPipelineDescriptorClass
	_MTLRenderPipelineDescriptorClassOnce sync.Once
)

func getMTLRenderPipelineDescriptorClass() MTLRenderPipelineDescriptorClass {
	_MTLRenderPipelineDescriptorClassOnce.Do(func() {
		_MTLRenderPipelineDescriptorClass = MTLRenderPipelineDescriptorClass{class: objc.GetClass("MTLRenderPipelineDescriptor")}
	})
	return _MTLRenderPipelineDescriptorClass
}

// GetMTLRenderPipelineDescriptorClass returns the class object for MTLRenderPipelineDescriptor.
func GetMTLRenderPipelineDescriptorClass() MTLRenderPipelineDescriptorClass {
	return getMTLRenderPipelineDescriptorClass()
}

type MTLRenderPipelineDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLRenderPipelineDescriptorClass) Alloc() MTLRenderPipelineDescriptor {
	rv := objc.Send[MTLRenderPipelineDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An argument of options you pass to a GPU device to get a render pipeline
// state.
//
// # Overview
// 
// An [MTLRenderPipelineDescriptor] instance configures the state of the
// pipeline to use during a rendering pass, including rasterization (such as
// multisampling), visibility, blending, tessellation, and graphics function
// state. Use standard allocation and initialization techniques to create an
// [MTLRenderPipelineDescriptor] object. Then configure and use the descriptor
// to create an [MTLRenderPipelineState] object.
// 
// To specify the vertex or fragment function in the rendering pipeline
// descriptor, set the [MTLRenderPipelineDescriptor.VertexFunction] or [MTLRenderPipelineDescriptor.FragmentFunction] property,
// respectively, to the desired [MTLFunction] object. The system ignores the
// tessellation stage properties if you don’t set the [MTLRenderPipelineDescriptor.VertexFunction]
// property to a post-tessellation vertex function. A vertex function is a
// post-tessellation vertex function if the `[[ patch(patch-type, N) ]]`
// attribute precedes the function’s signature in your Metal Shading
// Language source. See the “Post-Tessellation Vertex Functions” section
// of [Metal Shading Language Specification] for more information.
// 
// Setting the [MTLRenderPipelineDescriptor.FragmentFunction] property to `nil` disables the rasterization
// of pixels into the color attachment. This action is typically for
// outputting vertex function data into a buffer object, or for depth-only
// rendering.
// 
// If the vertex shader has an argument with per-vertex input attributes, set
// the [VertexDescriptor] property to an [MTLVertexDescriptor] object that
// describes the organization of that vertex data.
// 
// # Multisampling and the render pipeline
// 
// If a color attachment supports multisampling (essentially, the attachment
// is an [TextureType2DMultisample] type color texture), you can create
// multiple samples per fragment, and the following rendering pipeline
// descriptor properties determine coverage:
// 
// - [MTLRenderPipelineDescriptor.RasterSampleCount] is the number of samples for each pixel. - If
// [MTLRenderPipelineDescriptor.AlphaToCoverageEnabled] is [true], the GPU uses the alpha channel fragment
// output for [MTLRenderPipelineDescriptor.ColorAttachments] to compute a coverage mask that affects the
// values the GPU writes to all attachments (color, depth, and stencil). - If
// [MTLRenderPipelineDescriptor.AlphaToOneEnabled] is [true], the GPU changes alpha channel fragment
// values for [MTLRenderPipelineDescriptor.ColorAttachments] to `1.0`, which is the largest representable
// value.
// 
// If [MTLRenderPipelineDescriptor.AlphaToCoverageEnabled] is [true], an implementation-defined
// `coverageToMask` function uses the alpha channel fragment output from
// [MTLRenderPipelineDescriptor.ColorAttachments] to create an intermediate coverage mask, which sets a
// number of bits in its output proportionally to the value of the
// floating-point input. For example, if the input is `0.0f`, the function
// sets the output to `0x0`. If the input is `1.0f`, the function sets all
// output bits (in effect, `~0x0`). If the input is `0.5f`, the function sets
// half of the bits, according to the implementation, which often uses dither
// patterns.
// 
// To determine a final coverage mask, the function performs a logical [AND]
// on the resulting coverage mask `alphaCoverageMask` with the masks from the
// rasterizer and fragment shader, as the following code shows:
//
// [Metal Shading Language Specification]: https://developer.apple.com/metal/Metal-Shading-Language-Specification.pdf
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Identifying the render pipeline state object
//
//   - [MTLRenderPipelineDescriptor.Label]: A string that identifies the render pipeline descriptor.
//   - [MTLRenderPipelineDescriptor.SetLabel]
//
// # Specifying graphics functions and associated data
//
//   - [MTLRenderPipelineDescriptor.VertexFunction]: The vertex function the pipeline calls to process vertices.
//   - [MTLRenderPipelineDescriptor.SetVertexFunction]
//   - [MTLRenderPipelineDescriptor.FragmentFunction]: The fragment function the pipeline calls to process fragments.
//   - [MTLRenderPipelineDescriptor.SetFragmentFunction]
//   - [MTLRenderPipelineDescriptor.MaxVertexCallStackDepth]: The maximum function call depth from the top-most vertex shader function.
//   - [MTLRenderPipelineDescriptor.SetMaxVertexCallStackDepth]
//   - [MTLRenderPipelineDescriptor.MaxFragmentCallStackDepth]: The maximum function call depth from the top-most fragment shader function.
//   - [MTLRenderPipelineDescriptor.SetMaxFragmentCallStackDepth]
//
// # Specifying buffer layouts and fetch behavior
//
//   - [MTLRenderPipelineDescriptor.VertexDescriptor]: The organization of vertex data in an attribute’s argument table.
//   - [MTLRenderPipelineDescriptor.SetVertexDescriptor]
//
// # Specifying buffer mutability
//
//   - [MTLRenderPipelineDescriptor.VertexBuffers]: An array that contains the buffer mutability options for a render pipeline’s vertex function.
//   - [MTLRenderPipelineDescriptor.FragmentBuffers]: An array that contains the buffer mutability options for a render pipeline’s fragment function.
//
// # Specifying rendering pipeline state
//
//   - [MTLRenderPipelineDescriptor.Reset]: Specifies the default rendering pipeline state values for the descriptor.
//   - [MTLRenderPipelineDescriptor.ColorAttachments]: An array of attachments that store color data.
//   - [MTLRenderPipelineDescriptor.DepthAttachmentPixelFormat]: The pixel format of the attachment that stores depth data.
//   - [MTLRenderPipelineDescriptor.SetDepthAttachmentPixelFormat]
//   - [MTLRenderPipelineDescriptor.StencilAttachmentPixelFormat]: The pixel format of the attachment that stores stencil data.
//   - [MTLRenderPipelineDescriptor.SetStencilAttachmentPixelFormat]
//
// # Specifying rasterization and visibility state
//
//   - [MTLRenderPipelineDescriptor.AlphaToCoverageEnabled]: A Boolean value that indicates whether to read and use the alpha channel fragment output for color attachments to compute a sample coverage mask.
//   - [MTLRenderPipelineDescriptor.SetAlphaToCoverageEnabled]
//   - [MTLRenderPipelineDescriptor.AlphaToOneEnabled]: A Boolean value that indicates whether to force alpha channel values for color attachments to the largest representable value.
//   - [MTLRenderPipelineDescriptor.SetAlphaToOneEnabled]
//   - [MTLRenderPipelineDescriptor.RasterizationEnabled]: A Boolean value that determines whether the pipeline rasterizes primitives.
//   - [MTLRenderPipelineDescriptor.SetRasterizationEnabled]
//   - [MTLRenderPipelineDescriptor.InputPrimitiveTopology]: The type of primitive topology the pipeline renders.
//   - [MTLRenderPipelineDescriptor.SetInputPrimitiveTopology]
//   - [MTLRenderPipelineDescriptor.RasterSampleCount]: The number of samples the pipeline applies for each fragment.
//   - [MTLRenderPipelineDescriptor.SetRasterSampleCount]
//
// # Specifying tessellation state
//
//   - [MTLRenderPipelineDescriptor.MaxTessellationFactor]: The maximum tessellation factor that the tessellator uses when tessellating patches.
//   - [MTLRenderPipelineDescriptor.SetMaxTessellationFactor]
//   - [MTLRenderPipelineDescriptor.TessellationFactorScaleEnabled]: A Boolean value that determines whether the pipeline scales the tessellation factor.
//   - [MTLRenderPipelineDescriptor.SetTessellationFactorScaleEnabled]
//   - [MTLRenderPipelineDescriptor.TessellationFactorFormat]: The format of the tessellation factors in the tessellation factor buffer.
//   - [MTLRenderPipelineDescriptor.SetTessellationFactorFormat]
//   - [MTLRenderPipelineDescriptor.TessellationControlPointIndexType]: The size of the control point indices in a control point index buffer.
//   - [MTLRenderPipelineDescriptor.SetTessellationControlPointIndexType]
//   - [MTLRenderPipelineDescriptor.TessellationFactorStepFunction]: The step function for determining the tessellation factors for a patch from the tessellation factor buffer.
//   - [MTLRenderPipelineDescriptor.SetTessellationFactorStepFunction]
//   - [MTLRenderPipelineDescriptor.TessellationOutputWindingOrder]: The winding order of triangles from the tessellator.
//   - [MTLRenderPipelineDescriptor.SetTessellationOutputWindingOrder]
//   - [MTLRenderPipelineDescriptor.TessellationPartitionMode]: The partitioning mode that the tessellator uses to derive the number and spacing of segments for subdividing a corresponding edge.
//   - [MTLRenderPipelineDescriptor.SetTessellationPartitionMode]
//
// # Specifying indirect command buffers usage
//
//   - [MTLRenderPipelineDescriptor.SupportIndirectCommandBuffers]: A Boolean value that determines whether you can encode commands into an indirect command buffer using the render pipeline.
//   - [MTLRenderPipelineDescriptor.SetSupportIndirectCommandBuffers]
//
// # Specifying the maximum vertex amplification count
//
//   - [MTLRenderPipelineDescriptor.MaxVertexAmplificationCount]: The maximum vertex amplification count you can set when encoding render commands.
//   - [MTLRenderPipelineDescriptor.SetMaxVertexAmplificationCount]
//
// # Specifying precompiled shader binaries
//
//   - [MTLRenderPipelineDescriptor.SupportAddingVertexBinaryFunctions]: A Boolean value that indicates whether you can use the pipeline to create new pipelines by adding binary functions to the vertex shader’s callable functions list.
//   - [MTLRenderPipelineDescriptor.SetSupportAddingVertexBinaryFunctions]
//   - [MTLRenderPipelineDescriptor.SupportAddingFragmentBinaryFunctions]: A Boolean value that indicates whether you can use the pipeline to create new pipelines by adding binary functions to the fragment shader’s callable functions list.
//   - [MTLRenderPipelineDescriptor.SetSupportAddingFragmentBinaryFunctions]
//   - [MTLRenderPipelineDescriptor.BinaryArchives]: An array of binary archives to search for precompiled versions of the shader.
//   - [MTLRenderPipelineDescriptor.SetBinaryArchives]
//
// # Specifying callable functions for the pipeline
//
//   - [MTLRenderPipelineDescriptor.VertexLinkedFunctions]: Functions that you can specify as function arguments for the vertex shader when encoding commands that use the pipeline.
//   - [MTLRenderPipelineDescriptor.SetVertexLinkedFunctions]
//   - [MTLRenderPipelineDescriptor.FragmentLinkedFunctions]: Functions that you can specify as function arguments for the fragment shader when encoding commands that use the pipeline.
//   - [MTLRenderPipelineDescriptor.SetFragmentLinkedFunctions]
//
// # Specifying shader validation
//
//   - [MTLRenderPipelineDescriptor.ShaderValidation]: A value that enables or disables shader validation for the pipeline.
//   - [MTLRenderPipelineDescriptor.SetShaderValidation]
//
// # Instance Properties
//
//   - [MTLRenderPipelineDescriptor.FragmentPreloadedLibraries]
//   - [MTLRenderPipelineDescriptor.SetFragmentPreloadedLibraries]
//   - [MTLRenderPipelineDescriptor.VertexPreloadedLibraries]
//   - [MTLRenderPipelineDescriptor.SetVertexPreloadedLibraries]
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor
type MTLRenderPipelineDescriptor struct {
	objectivec.Object
}

// MTLRenderPipelineDescriptorFromID constructs a [MTLRenderPipelineDescriptor] from an objc.ID.
//
// An argument of options you pass to a GPU device to get a render pipeline
// state.
func MTLRenderPipelineDescriptorFromID(id objc.ID) MTLRenderPipelineDescriptor {
	return MTLRenderPipelineDescriptor{objectivec.Object{ID: id}}
}
// NOTE: MTLRenderPipelineDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLRenderPipelineDescriptor] class.
//
// # Identifying the render pipeline state object
//
//   - [IMTLRenderPipelineDescriptor.Label]: A string that identifies the render pipeline descriptor.
//   - [IMTLRenderPipelineDescriptor.SetLabel]
//
// # Specifying graphics functions and associated data
//
//   - [IMTLRenderPipelineDescriptor.VertexFunction]: The vertex function the pipeline calls to process vertices.
//   - [IMTLRenderPipelineDescriptor.SetVertexFunction]
//   - [IMTLRenderPipelineDescriptor.FragmentFunction]: The fragment function the pipeline calls to process fragments.
//   - [IMTLRenderPipelineDescriptor.SetFragmentFunction]
//   - [IMTLRenderPipelineDescriptor.MaxVertexCallStackDepth]: The maximum function call depth from the top-most vertex shader function.
//   - [IMTLRenderPipelineDescriptor.SetMaxVertexCallStackDepth]
//   - [IMTLRenderPipelineDescriptor.MaxFragmentCallStackDepth]: The maximum function call depth from the top-most fragment shader function.
//   - [IMTLRenderPipelineDescriptor.SetMaxFragmentCallStackDepth]
//
// # Specifying buffer layouts and fetch behavior
//
//   - [IMTLRenderPipelineDescriptor.VertexDescriptor]: The organization of vertex data in an attribute’s argument table.
//   - [IMTLRenderPipelineDescriptor.SetVertexDescriptor]
//
// # Specifying buffer mutability
//
//   - [IMTLRenderPipelineDescriptor.VertexBuffers]: An array that contains the buffer mutability options for a render pipeline’s vertex function.
//   - [IMTLRenderPipelineDescriptor.FragmentBuffers]: An array that contains the buffer mutability options for a render pipeline’s fragment function.
//
// # Specifying rendering pipeline state
//
//   - [IMTLRenderPipelineDescriptor.Reset]: Specifies the default rendering pipeline state values for the descriptor.
//   - [IMTLRenderPipelineDescriptor.ColorAttachments]: An array of attachments that store color data.
//   - [IMTLRenderPipelineDescriptor.DepthAttachmentPixelFormat]: The pixel format of the attachment that stores depth data.
//   - [IMTLRenderPipelineDescriptor.SetDepthAttachmentPixelFormat]
//   - [IMTLRenderPipelineDescriptor.StencilAttachmentPixelFormat]: The pixel format of the attachment that stores stencil data.
//   - [IMTLRenderPipelineDescriptor.SetStencilAttachmentPixelFormat]
//
// # Specifying rasterization and visibility state
//
//   - [IMTLRenderPipelineDescriptor.AlphaToCoverageEnabled]: A Boolean value that indicates whether to read and use the alpha channel fragment output for color attachments to compute a sample coverage mask.
//   - [IMTLRenderPipelineDescriptor.SetAlphaToCoverageEnabled]
//   - [IMTLRenderPipelineDescriptor.AlphaToOneEnabled]: A Boolean value that indicates whether to force alpha channel values for color attachments to the largest representable value.
//   - [IMTLRenderPipelineDescriptor.SetAlphaToOneEnabled]
//   - [IMTLRenderPipelineDescriptor.RasterizationEnabled]: A Boolean value that determines whether the pipeline rasterizes primitives.
//   - [IMTLRenderPipelineDescriptor.SetRasterizationEnabled]
//   - [IMTLRenderPipelineDescriptor.InputPrimitiveTopology]: The type of primitive topology the pipeline renders.
//   - [IMTLRenderPipelineDescriptor.SetInputPrimitiveTopology]
//   - [IMTLRenderPipelineDescriptor.RasterSampleCount]: The number of samples the pipeline applies for each fragment.
//   - [IMTLRenderPipelineDescriptor.SetRasterSampleCount]
//
// # Specifying tessellation state
//
//   - [IMTLRenderPipelineDescriptor.MaxTessellationFactor]: The maximum tessellation factor that the tessellator uses when tessellating patches.
//   - [IMTLRenderPipelineDescriptor.SetMaxTessellationFactor]
//   - [IMTLRenderPipelineDescriptor.TessellationFactorScaleEnabled]: A Boolean value that determines whether the pipeline scales the tessellation factor.
//   - [IMTLRenderPipelineDescriptor.SetTessellationFactorScaleEnabled]
//   - [IMTLRenderPipelineDescriptor.TessellationFactorFormat]: The format of the tessellation factors in the tessellation factor buffer.
//   - [IMTLRenderPipelineDescriptor.SetTessellationFactorFormat]
//   - [IMTLRenderPipelineDescriptor.TessellationControlPointIndexType]: The size of the control point indices in a control point index buffer.
//   - [IMTLRenderPipelineDescriptor.SetTessellationControlPointIndexType]
//   - [IMTLRenderPipelineDescriptor.TessellationFactorStepFunction]: The step function for determining the tessellation factors for a patch from the tessellation factor buffer.
//   - [IMTLRenderPipelineDescriptor.SetTessellationFactorStepFunction]
//   - [IMTLRenderPipelineDescriptor.TessellationOutputWindingOrder]: The winding order of triangles from the tessellator.
//   - [IMTLRenderPipelineDescriptor.SetTessellationOutputWindingOrder]
//   - [IMTLRenderPipelineDescriptor.TessellationPartitionMode]: The partitioning mode that the tessellator uses to derive the number and spacing of segments for subdividing a corresponding edge.
//   - [IMTLRenderPipelineDescriptor.SetTessellationPartitionMode]
//
// # Specifying indirect command buffers usage
//
//   - [IMTLRenderPipelineDescriptor.SupportIndirectCommandBuffers]: A Boolean value that determines whether you can encode commands into an indirect command buffer using the render pipeline.
//   - [IMTLRenderPipelineDescriptor.SetSupportIndirectCommandBuffers]
//
// # Specifying the maximum vertex amplification count
//
//   - [IMTLRenderPipelineDescriptor.MaxVertexAmplificationCount]: The maximum vertex amplification count you can set when encoding render commands.
//   - [IMTLRenderPipelineDescriptor.SetMaxVertexAmplificationCount]
//
// # Specifying precompiled shader binaries
//
//   - [IMTLRenderPipelineDescriptor.SupportAddingVertexBinaryFunctions]: A Boolean value that indicates whether you can use the pipeline to create new pipelines by adding binary functions to the vertex shader’s callable functions list.
//   - [IMTLRenderPipelineDescriptor.SetSupportAddingVertexBinaryFunctions]
//   - [IMTLRenderPipelineDescriptor.SupportAddingFragmentBinaryFunctions]: A Boolean value that indicates whether you can use the pipeline to create new pipelines by adding binary functions to the fragment shader’s callable functions list.
//   - [IMTLRenderPipelineDescriptor.SetSupportAddingFragmentBinaryFunctions]
//   - [IMTLRenderPipelineDescriptor.BinaryArchives]: An array of binary archives to search for precompiled versions of the shader.
//   - [IMTLRenderPipelineDescriptor.SetBinaryArchives]
//
// # Specifying callable functions for the pipeline
//
//   - [IMTLRenderPipelineDescriptor.VertexLinkedFunctions]: Functions that you can specify as function arguments for the vertex shader when encoding commands that use the pipeline.
//   - [IMTLRenderPipelineDescriptor.SetVertexLinkedFunctions]
//   - [IMTLRenderPipelineDescriptor.FragmentLinkedFunctions]: Functions that you can specify as function arguments for the fragment shader when encoding commands that use the pipeline.
//   - [IMTLRenderPipelineDescriptor.SetFragmentLinkedFunctions]
//
// # Specifying shader validation
//
//   - [IMTLRenderPipelineDescriptor.ShaderValidation]: A value that enables or disables shader validation for the pipeline.
//   - [IMTLRenderPipelineDescriptor.SetShaderValidation]
//
// # Instance Properties
//
//   - [IMTLRenderPipelineDescriptor.FragmentPreloadedLibraries]
//   - [IMTLRenderPipelineDescriptor.SetFragmentPreloadedLibraries]
//   - [IMTLRenderPipelineDescriptor.VertexPreloadedLibraries]
//   - [IMTLRenderPipelineDescriptor.SetVertexPreloadedLibraries]
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor
type IMTLRenderPipelineDescriptor interface {
	objectivec.IObject

	// Topic: Identifying the render pipeline state object

	// A string that identifies the render pipeline descriptor.
	Label() string
	SetLabel(value string)

	// Topic: Specifying graphics functions and associated data

	// The vertex function the pipeline calls to process vertices.
	VertexFunction() MTLFunction
	SetVertexFunction(value MTLFunction)
	// The fragment function the pipeline calls to process fragments.
	FragmentFunction() MTLFunction
	SetFragmentFunction(value MTLFunction)
	// The maximum function call depth from the top-most vertex shader function.
	MaxVertexCallStackDepth() uint
	SetMaxVertexCallStackDepth(value uint)
	// The maximum function call depth from the top-most fragment shader function.
	MaxFragmentCallStackDepth() uint
	SetMaxFragmentCallStackDepth(value uint)

	// Topic: Specifying buffer layouts and fetch behavior

	// The organization of vertex data in an attribute’s argument table.
	VertexDescriptor() IMTLVertexDescriptor
	SetVertexDescriptor(value IMTLVertexDescriptor)

	// Topic: Specifying buffer mutability

	// An array that contains the buffer mutability options for a render pipeline’s vertex function.
	VertexBuffers() IMTLPipelineBufferDescriptorArray
	// An array that contains the buffer mutability options for a render pipeline’s fragment function.
	FragmentBuffers() IMTLPipelineBufferDescriptorArray

	// Topic: Specifying rendering pipeline state

	// Specifies the default rendering pipeline state values for the descriptor.
	Reset()
	// An array of attachments that store color data.
	ColorAttachments() IMTLRenderPipelineColorAttachmentDescriptorArray
	// The pixel format of the attachment that stores depth data.
	DepthAttachmentPixelFormat() MTLPixelFormat
	SetDepthAttachmentPixelFormat(value MTLPixelFormat)
	// The pixel format of the attachment that stores stencil data.
	StencilAttachmentPixelFormat() MTLPixelFormat
	SetStencilAttachmentPixelFormat(value MTLPixelFormat)

	// Topic: Specifying rasterization and visibility state

	// A Boolean value that indicates whether to read and use the alpha channel fragment output for color attachments to compute a sample coverage mask.
	AlphaToCoverageEnabled() bool
	SetAlphaToCoverageEnabled(value bool)
	// A Boolean value that indicates whether to force alpha channel values for color attachments to the largest representable value.
	AlphaToOneEnabled() bool
	SetAlphaToOneEnabled(value bool)
	// A Boolean value that determines whether the pipeline rasterizes primitives.
	RasterizationEnabled() bool
	SetRasterizationEnabled(value bool)
	// The type of primitive topology the pipeline renders.
	InputPrimitiveTopology() MTLPrimitiveTopologyClass
	SetInputPrimitiveTopology(value MTLPrimitiveTopologyClass)
	// The number of samples the pipeline applies for each fragment.
	RasterSampleCount() uint
	SetRasterSampleCount(value uint)

	// Topic: Specifying tessellation state

	// The maximum tessellation factor that the tessellator uses when tessellating patches.
	MaxTessellationFactor() uint
	SetMaxTessellationFactor(value uint)
	// A Boolean value that determines whether the pipeline scales the tessellation factor.
	TessellationFactorScaleEnabled() bool
	SetTessellationFactorScaleEnabled(value bool)
	// The format of the tessellation factors in the tessellation factor buffer.
	TessellationFactorFormat() MTLTessellationFactorFormat
	SetTessellationFactorFormat(value MTLTessellationFactorFormat)
	// The size of the control point indices in a control point index buffer.
	TessellationControlPointIndexType() MTLTessellationControlPointIndexType
	SetTessellationControlPointIndexType(value MTLTessellationControlPointIndexType)
	// The step function for determining the tessellation factors for a patch from the tessellation factor buffer.
	TessellationFactorStepFunction() MTLTessellationFactorStepFunction
	SetTessellationFactorStepFunction(value MTLTessellationFactorStepFunction)
	// The winding order of triangles from the tessellator.
	TessellationOutputWindingOrder() MTLWinding
	SetTessellationOutputWindingOrder(value MTLWinding)
	// The partitioning mode that the tessellator uses to derive the number and spacing of segments for subdividing a corresponding edge.
	TessellationPartitionMode() MTLTessellationPartitionMode
	SetTessellationPartitionMode(value MTLTessellationPartitionMode)

	// Topic: Specifying indirect command buffers usage

	// A Boolean value that determines whether you can encode commands into an indirect command buffer using the render pipeline.
	SupportIndirectCommandBuffers() bool
	SetSupportIndirectCommandBuffers(value bool)

	// Topic: Specifying the maximum vertex amplification count

	// The maximum vertex amplification count you can set when encoding render commands.
	MaxVertexAmplificationCount() uint
	SetMaxVertexAmplificationCount(value uint)

	// Topic: Specifying precompiled shader binaries

	// A Boolean value that indicates whether you can use the pipeline to create new pipelines by adding binary functions to the vertex shader’s callable functions list.
	SupportAddingVertexBinaryFunctions() bool
	SetSupportAddingVertexBinaryFunctions(value bool)
	// A Boolean value that indicates whether you can use the pipeline to create new pipelines by adding binary functions to the fragment shader’s callable functions list.
	SupportAddingFragmentBinaryFunctions() bool
	SetSupportAddingFragmentBinaryFunctions(value bool)
	// An array of binary archives to search for precompiled versions of the shader.
	BinaryArchives() []objectivec.IObject
	SetBinaryArchives(value []objectivec.IObject)

	// Topic: Specifying callable functions for the pipeline

	// Functions that you can specify as function arguments for the vertex shader when encoding commands that use the pipeline.
	VertexLinkedFunctions() IMTLLinkedFunctions
	SetVertexLinkedFunctions(value IMTLLinkedFunctions)
	// Functions that you can specify as function arguments for the fragment shader when encoding commands that use the pipeline.
	FragmentLinkedFunctions() IMTLLinkedFunctions
	SetFragmentLinkedFunctions(value IMTLLinkedFunctions)

	// Topic: Specifying shader validation

	// A value that enables or disables shader validation for the pipeline.
	ShaderValidation() MTLShaderValidation
	SetShaderValidation(value MTLShaderValidation)

	// Topic: Instance Properties

	FragmentPreloadedLibraries() []objectivec.IObject
	SetFragmentPreloadedLibraries(value []objectivec.IObject)
	VertexPreloadedLibraries() []objectivec.IObject
	SetVertexPreloadedLibraries(value []objectivec.IObject)
}

// Init initializes the instance.
func (r MTLRenderPipelineDescriptor) Init() MTLRenderPipelineDescriptor {
	rv := objc.Send[MTLRenderPipelineDescriptor](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MTLRenderPipelineDescriptor) Autorelease() MTLRenderPipelineDescriptor {
	rv := objc.Send[MTLRenderPipelineDescriptor](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLRenderPipelineDescriptor creates a new MTLRenderPipelineDescriptor instance.
func NewMTLRenderPipelineDescriptor() MTLRenderPipelineDescriptor {
	class := getMTLRenderPipelineDescriptorClass()
	rv := objc.Send[MTLRenderPipelineDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Specifies the default rendering pipeline state values for the descriptor.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/reset()
func (r MTLRenderPipelineDescriptor) Reset() {
	objc.Send[objc.ID](r.ID, objc.Sel("reset"))
}

// A string that identifies the render pipeline descriptor.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/label
func (r MTLRenderPipelineDescriptor) Label() string {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}
func (r MTLRenderPipelineDescriptor) SetLabel(value string) {
	objc.Send[struct{}](r.ID, objc.Sel("setLabel:"), objc.String(value))
}
// The vertex function the pipeline calls to process vertices.
//
// # Discussion
// 
// The default value is `nil`. The vertex function needs to be specified. The
// vertex function can be either a regular vertex function or a
// post-tessellation vertex function.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/vertexFunction
func (r MTLRenderPipelineDescriptor) VertexFunction() MTLFunction {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("vertexFunction"))
	return MTLFunctionObjectFromID(rv)
}
func (r MTLRenderPipelineDescriptor) SetVertexFunction(value MTLFunction) {
	objc.Send[struct{}](r.ID, objc.Sel("setVertexFunction:"), value)
}
// The fragment function the pipeline calls to process fragments.
//
// # Discussion
// 
// The default value is `nil`. If this value is `nil`, then there is no
// fragment function and therefore no writes to the color render target occur.
// Depth and stencil writes and visibility result counting can still proceed.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/fragmentFunction
func (r MTLRenderPipelineDescriptor) FragmentFunction() MTLFunction {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("fragmentFunction"))
	return MTLFunctionObjectFromID(rv)
}
func (r MTLRenderPipelineDescriptor) SetFragmentFunction(value MTLFunction) {
	objc.Send[struct{}](r.ID, objc.Sel("setFragmentFunction:"), value)
}
// The maximum function call depth from the top-most vertex shader function.
//
// # Discussion
// 
// The default value is 1.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/maxVertexCallStackDepth
func (r MTLRenderPipelineDescriptor) MaxVertexCallStackDepth() uint {
	rv := objc.Send[uint](r.ID, objc.Sel("maxVertexCallStackDepth"))
	return rv
}
func (r MTLRenderPipelineDescriptor) SetMaxVertexCallStackDepth(value uint) {
	objc.Send[struct{}](r.ID, objc.Sel("setMaxVertexCallStackDepth:"), value)
}
// The maximum function call depth from the top-most fragment shader function.
//
// # Discussion
// 
// The default value is 1.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/maxFragmentCallStackDepth
func (r MTLRenderPipelineDescriptor) MaxFragmentCallStackDepth() uint {
	rv := objc.Send[uint](r.ID, objc.Sel("maxFragmentCallStackDepth"))
	return rv
}
func (r MTLRenderPipelineDescriptor) SetMaxFragmentCallStackDepth(value uint) {
	objc.Send[struct{}](r.ID, objc.Sel("setMaxFragmentCallStackDepth:"), value)
}
// The organization of vertex data in an attribute’s argument table.
//
// # Discussion
// 
// An [MTLVertexDescriptor] instance is used to describe the organization of
// per-vertex input structs passed in an argument of a vertex shader function.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/vertexDescriptor
func (r MTLRenderPipelineDescriptor) VertexDescriptor() IMTLVertexDescriptor {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("vertexDescriptor"))
	return MTLVertexDescriptorFromID(objc.ID(rv))
}
func (r MTLRenderPipelineDescriptor) SetVertexDescriptor(value IMTLVertexDescriptor) {
	objc.Send[struct{}](r.ID, objc.Sel("setVertexDescriptor:"), value)
}
// An array that contains the buffer mutability options for a render
// pipeline’s vertex function.
//
// # Discussion
// 
// This property returns an array of [MTLPipelineBufferDescriptor] instances,
// where each element corresponds to the same index in the buffer argument
// table for the render pipeline’s vertex function.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/vertexBuffers
func (r MTLRenderPipelineDescriptor) VertexBuffers() IMTLPipelineBufferDescriptorArray {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("vertexBuffers"))
	return MTLPipelineBufferDescriptorArrayFromID(objc.ID(rv))
}
// An array that contains the buffer mutability options for a render
// pipeline’s fragment function.
//
// # Discussion
// 
// This property returns an array of [MTLPipelineBufferDescriptor] instances,
// where each element corresponds to the index in the buffer argument table
// for the render pipeline’s fragment function.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/fragmentBuffers
func (r MTLRenderPipelineDescriptor) FragmentBuffers() IMTLPipelineBufferDescriptorArray {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("fragmentBuffers"))
	return MTLPipelineBufferDescriptorArrayFromID(objc.ID(rv))
}
// An array of attachments that store color data.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/colorAttachments
func (r MTLRenderPipelineDescriptor) ColorAttachments() IMTLRenderPipelineColorAttachmentDescriptorArray {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("colorAttachments"))
	return MTLRenderPipelineColorAttachmentDescriptorArrayFromID(objc.ID(rv))
}
// The pixel format of the attachment that stores depth data.
//
// # Discussion
// 
// By default, the pixel format of the rendering pipeline state for each
// attachment is [MTLPixelFormatInvalid].
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/depthAttachmentPixelFormat
func (r MTLRenderPipelineDescriptor) DepthAttachmentPixelFormat() MTLPixelFormat {
	rv := objc.Send[MTLPixelFormat](r.ID, objc.Sel("depthAttachmentPixelFormat"))
	return MTLPixelFormat(rv)
}
func (r MTLRenderPipelineDescriptor) SetDepthAttachmentPixelFormat(value MTLPixelFormat) {
	objc.Send[struct{}](r.ID, objc.Sel("setDepthAttachmentPixelFormat:"), value)
}
// The pixel format of the attachment that stores stencil data.
//
// # Discussion
// 
// By default, the pixel format of the rendering pipeline state for each
// attachment is [MTLPixelFormatInvalid].
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/stencilAttachmentPixelFormat
func (r MTLRenderPipelineDescriptor) StencilAttachmentPixelFormat() MTLPixelFormat {
	rv := objc.Send[MTLPixelFormat](r.ID, objc.Sel("stencilAttachmentPixelFormat"))
	return MTLPixelFormat(rv)
}
func (r MTLRenderPipelineDescriptor) SetStencilAttachmentPixelFormat(value MTLPixelFormat) {
	objc.Send[struct{}](r.ID, objc.Sel("setStencilAttachmentPixelFormat:"), value)
}
// A Boolean value that indicates whether to read and use the alpha channel
// fragment output for color attachments to compute a sample coverage mask.
//
// # Discussion
// 
// The default value is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/isAlphaToCoverageEnabled
func (r MTLRenderPipelineDescriptor) AlphaToCoverageEnabled() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("isAlphaToCoverageEnabled"))
	return rv
}
func (r MTLRenderPipelineDescriptor) SetAlphaToCoverageEnabled(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setAlphaToCoverageEnabled:"), value)
}
// A Boolean value that indicates whether to force alpha channel values for
// color attachments to the largest representable value.
//
// # Discussion
// 
// The default value is [false].
// 
// If enabled, alpha channel fragment values are only forced for
// `colorAttachments[0]`. Other attachments are unaffected.
// 
// You may use `alphaToOneEnabled` when you want to write an alpha value that
// represents partial coverage of the pixel, but also want to disable blending
// (by forcing alpha to one).
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/isAlphaToOneEnabled
func (r MTLRenderPipelineDescriptor) AlphaToOneEnabled() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("isAlphaToOneEnabled"))
	return rv
}
func (r MTLRenderPipelineDescriptor) SetAlphaToOneEnabled(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setAlphaToOneEnabled:"), value)
}
// A Boolean value that determines whether the pipeline rasterizes primitives.
//
// # Discussion
// 
// The default value is [true], indicating that primitives are rasterized. If
// the value is [false], then primitives are dropped prior to rasterization
// (i.e. rasterization is disabled). Disabling rasterization may be useful to
// gather data from vertex-only transformations.
// 
// When this value is [false], no fragments are processed and the vertex
// shader function needs to return `void`.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/isRasterizationEnabled
func (r MTLRenderPipelineDescriptor) RasterizationEnabled() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("isRasterizationEnabled"))
	return rv
}
func (r MTLRenderPipelineDescriptor) SetRasterizationEnabled(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setRasterizationEnabled:"), value)
}
// The type of primitive topology the pipeline renders.
//
// # Discussion
// 
// Your app needs to specify this value when layered rendering is enabled.
// 
// The default value is [MTLPrimitiveTopologyClassUnspecified].
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/inputPrimitiveTopology
func (r MTLRenderPipelineDescriptor) InputPrimitiveTopology() MTLPrimitiveTopologyClass {
	rv := objc.Send[MTLPrimitiveTopologyClass](r.ID, objc.Sel("inputPrimitiveTopology"))
	return MTLPrimitiveTopologyClass(rv)
}
func (r MTLRenderPipelineDescriptor) SetInputPrimitiveTopology(value MTLPrimitiveTopologyClass) {
	objc.Send[struct{}](r.ID, objc.Sel("setInputPrimitiveTopology:"), value)
}
// The number of samples the pipeline applies for each fragment.
//
// # Discussion
// 
// The render pipeline state honors this property only if the pipeline render
// targets support multisampling.
// 
// When your create an [MTLRenderCommandEncoder] instance, this property’s
// value needs to be equal to the number of render target textures.
// Furthermore, the texture type of all render target textures need to be
// [TextureType2DMultisample].
// 
// The number of samples a GPU supports varies by device. You can check
// whether an [MTLDevice] instance supports a specific sample count by calling
// its [SupportsTextureSampleCount] method.
// 
// The default value for this property is `1`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/rasterSampleCount
func (r MTLRenderPipelineDescriptor) RasterSampleCount() uint {
	rv := objc.Send[uint](r.ID, objc.Sel("rasterSampleCount"))
	return rv
}
func (r MTLRenderPipelineDescriptor) SetRasterSampleCount(value uint) {
	objc.Send[struct{}](r.ID, objc.Sel("setRasterSampleCount:"), value)
}
// The maximum tessellation factor that the tessellator uses when tessellating
// patches.
//
// # Discussion
// 
// The default value is `16` and the maximum value is `64`. Any value in
// between needs to be set according to the partitioning mode specified by the
// [TessellationPartitionMode] property:
// 
// - For the [TessellationPartitionModePow2] partitioning mode, the value
// needs to be a power of two. - For the [TessellationPartitionModeInteger]
// partitioning mode, the value can be an even or odd number. - For the
// [TessellationPartitionModeFractionalOdd] or
// [TessellationPartitionModeFractionalEven] partitioning mode, the value
// needs to be an even number.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/maxTessellationFactor
func (r MTLRenderPipelineDescriptor) MaxTessellationFactor() uint {
	rv := objc.Send[uint](r.ID, objc.Sel("maxTessellationFactor"))
	return rv
}
func (r MTLRenderPipelineDescriptor) SetMaxTessellationFactor(value uint) {
	objc.Send[struct{}](r.ID, objc.Sel("setMaxTessellationFactor:"), value)
}
// A Boolean value that determines whether the pipeline scales the
// tessellation factor.
//
// # Discussion
// 
// The default value is [false].
// 
// If this value is [true], a scale factor is applied to the tessellation
// factors after the patch cull check is performed but before the tessellation
// factors are clamped to the value of [MaxTessellationFactor]. The scale
// factor is applied only if the patch is not culled.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/isTessellationFactorScaleEnabled
func (r MTLRenderPipelineDescriptor) TessellationFactorScaleEnabled() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("isTessellationFactorScaleEnabled"))
	return rv
}
func (r MTLRenderPipelineDescriptor) SetTessellationFactorScaleEnabled(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setTessellationFactorScaleEnabled:"), value)
}
// The format of the tessellation factors in the tessellation factor buffer.
//
// # Discussion
// 
// The default value is [TessellationFactorFormatHalf].
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/tessellationFactorFormat
func (r MTLRenderPipelineDescriptor) TessellationFactorFormat() MTLTessellationFactorFormat {
	rv := objc.Send[MTLTessellationFactorFormat](r.ID, objc.Sel("tessellationFactorFormat"))
	return MTLTessellationFactorFormat(rv)
}
func (r MTLRenderPipelineDescriptor) SetTessellationFactorFormat(value MTLTessellationFactorFormat) {
	objc.Send[struct{}](r.ID, objc.Sel("setTessellationFactorFormat:"), value)
}
// The size of the control point indices in a control point index buffer.
//
// # Discussion
// 
// The default value is [TessellationControlPointIndexTypeNone]; use this
// value when drawing patches without a control point index buffer. This value
// needs to be either [TessellationControlPointIndexTypeUInt16] or
// [TessellationControlPointIndexTypeUInt32] when drawing patches with indexed
// control points.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/tessellationControlPointIndexType
func (r MTLRenderPipelineDescriptor) TessellationControlPointIndexType() MTLTessellationControlPointIndexType {
	rv := objc.Send[MTLTessellationControlPointIndexType](r.ID, objc.Sel("tessellationControlPointIndexType"))
	return MTLTessellationControlPointIndexType(rv)
}
func (r MTLRenderPipelineDescriptor) SetTessellationControlPointIndexType(value MTLTessellationControlPointIndexType) {
	objc.Send[struct{}](r.ID, objc.Sel("setTessellationControlPointIndexType:"), value)
}
// The step function for determining the tessellation factors for a patch from
// the tessellation factor buffer.
//
// # Discussion
// 
// The default value is [TessellationFactorStepFunctionConstant].
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/tessellationFactorStepFunction
func (r MTLRenderPipelineDescriptor) TessellationFactorStepFunction() MTLTessellationFactorStepFunction {
	rv := objc.Send[MTLTessellationFactorStepFunction](r.ID, objc.Sel("tessellationFactorStepFunction"))
	return MTLTessellationFactorStepFunction(rv)
}
func (r MTLRenderPipelineDescriptor) SetTessellationFactorStepFunction(value MTLTessellationFactorStepFunction) {
	objc.Send[struct{}](r.ID, objc.Sel("setTessellationFactorStepFunction:"), value)
}
// The winding order of triangles from the tessellator.
//
// # Discussion
// 
// The default value is [WindingClockwise].
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/tessellationOutputWindingOrder
func (r MTLRenderPipelineDescriptor) TessellationOutputWindingOrder() MTLWinding {
	rv := objc.Send[MTLWinding](r.ID, objc.Sel("tessellationOutputWindingOrder"))
	return MTLWinding(rv)
}
func (r MTLRenderPipelineDescriptor) SetTessellationOutputWindingOrder(value MTLWinding) {
	objc.Send[struct{}](r.ID, objc.Sel("setTessellationOutputWindingOrder:"), value)
}
// The partitioning mode that the tessellator uses to derive the number and
// spacing of segments for subdividing a corresponding edge.
//
// # Discussion
// 
// The default value is [TessellationPartitionModePow2].
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/tessellationPartitionMode
func (r MTLRenderPipelineDescriptor) TessellationPartitionMode() MTLTessellationPartitionMode {
	rv := objc.Send[MTLTessellationPartitionMode](r.ID, objc.Sel("tessellationPartitionMode"))
	return MTLTessellationPartitionMode(rv)
}
func (r MTLRenderPipelineDescriptor) SetTessellationPartitionMode(value MTLTessellationPartitionMode) {
	objc.Send[struct{}](r.ID, objc.Sel("setTessellationPartitionMode:"), value)
}
// A Boolean value that determines whether you can encode commands into an
// indirect command buffer using the render pipeline.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/supportIndirectCommandBuffers
func (r MTLRenderPipelineDescriptor) SupportIndirectCommandBuffers() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("supportIndirectCommandBuffers"))
	return rv
}
func (r MTLRenderPipelineDescriptor) SetSupportIndirectCommandBuffers(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setSupportIndirectCommandBuffers:"), value)
}
// The maximum vertex amplification count you can set when encoding render
// commands.
//
// # Discussion
// 
// Before setting this property, call the [SupportsVertexAmplificationCount]
// method on the device object to determine whether that amplification count
// is supported.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/maxVertexAmplificationCount
func (r MTLRenderPipelineDescriptor) MaxVertexAmplificationCount() uint {
	rv := objc.Send[uint](r.ID, objc.Sel("maxVertexAmplificationCount"))
	return rv
}
func (r MTLRenderPipelineDescriptor) SetMaxVertexAmplificationCount(value uint) {
	objc.Send[struct{}](r.ID, objc.Sel("setMaxVertexAmplificationCount:"), value)
}
// A Boolean value that indicates whether you can use the pipeline to create
// new pipelines by adding binary functions to the vertex shader’s callable
// functions list.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/supportAddingVertexBinaryFunctions
func (r MTLRenderPipelineDescriptor) SupportAddingVertexBinaryFunctions() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("supportAddingVertexBinaryFunctions"))
	return rv
}
func (r MTLRenderPipelineDescriptor) SetSupportAddingVertexBinaryFunctions(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setSupportAddingVertexBinaryFunctions:"), value)
}
// A Boolean value that indicates whether you can use the pipeline to create
// new pipelines by adding binary functions to the fragment shader’s
// callable functions list.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/supportAddingFragmentBinaryFunctions
func (r MTLRenderPipelineDescriptor) SupportAddingFragmentBinaryFunctions() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("supportAddingFragmentBinaryFunctions"))
	return rv
}
func (r MTLRenderPipelineDescriptor) SetSupportAddingFragmentBinaryFunctions(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setSupportAddingFragmentBinaryFunctions:"), value)
}
// An array of binary archives to search for precompiled versions of the
// shader.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/binaryArchives
func (r MTLRenderPipelineDescriptor) BinaryArchives() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](r.ID, objc.Sel("binaryArchives"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
func (r MTLRenderPipelineDescriptor) SetBinaryArchives(value []objectivec.IObject) {
	objc.Send[struct{}](r.ID, objc.Sel("setBinaryArchives:"), objectivec.IObjectSliceToNSArray(value))
}
// Functions that you can specify as function arguments for the vertex shader
// when encoding commands that use the pipeline.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/vertexLinkedFunctions
func (r MTLRenderPipelineDescriptor) VertexLinkedFunctions() IMTLLinkedFunctions {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("vertexLinkedFunctions"))
	return MTLLinkedFunctionsFromID(objc.ID(rv))
}
func (r MTLRenderPipelineDescriptor) SetVertexLinkedFunctions(value IMTLLinkedFunctions) {
	objc.Send[struct{}](r.ID, objc.Sel("setVertexLinkedFunctions:"), value)
}
// Functions that you can specify as function arguments for the fragment
// shader when encoding commands that use the pipeline.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/fragmentLinkedFunctions
func (r MTLRenderPipelineDescriptor) FragmentLinkedFunctions() IMTLLinkedFunctions {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("fragmentLinkedFunctions"))
	return MTLLinkedFunctionsFromID(objc.ID(rv))
}
func (r MTLRenderPipelineDescriptor) SetFragmentLinkedFunctions(value IMTLLinkedFunctions) {
	objc.Send[struct{}](r.ID, objc.Sel("setFragmentLinkedFunctions:"), value)
}
// A value that enables or disables shader validation for the pipeline.
//
// # Discussion
// 
// You can override the value using either of these environment variables:
// `MTL_SHADER_VALIDATION_ENABLE_PIPELINES` or
// `MTL_SHADER_VALIDATION_DISABLE_PIPELINES.`
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/shaderValidation
func (r MTLRenderPipelineDescriptor) ShaderValidation() MTLShaderValidation {
	rv := objc.Send[MTLShaderValidation](r.ID, objc.Sel("shaderValidation"))
	return MTLShaderValidation(rv)
}
func (r MTLRenderPipelineDescriptor) SetShaderValidation(value MTLShaderValidation) {
	objc.Send[struct{}](r.ID, objc.Sel("setShaderValidation:"), value)
}
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/fragmentPreloadedLibraries
func (r MTLRenderPipelineDescriptor) FragmentPreloadedLibraries() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](r.ID, objc.Sel("fragmentPreloadedLibraries"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
func (r MTLRenderPipelineDescriptor) SetFragmentPreloadedLibraries(value []objectivec.IObject) {
	objc.Send[struct{}](r.ID, objc.Sel("setFragmentPreloadedLibraries:"), objectivec.IObjectSliceToNSArray(value))
}
// See: https://developer.apple.com/documentation/Metal/MTLRenderPipelineDescriptor/vertexPreloadedLibraries
func (r MTLRenderPipelineDescriptor) VertexPreloadedLibraries() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](r.ID, objc.Sel("vertexPreloadedLibraries"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
func (r MTLRenderPipelineDescriptor) SetVertexPreloadedLibraries(value []objectivec.IObject) {
	objc.Send[struct{}](r.ID, objc.Sel("setVertexPreloadedLibraries:"), objectivec.IObjectSliceToNSArray(value))
}

