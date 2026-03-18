// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLTileRenderPipelineDescriptor] class.
var (
	_MTLTileRenderPipelineDescriptorClass     MTLTileRenderPipelineDescriptorClass
	_MTLTileRenderPipelineDescriptorClassOnce sync.Once
)

func getMTLTileRenderPipelineDescriptorClass() MTLTileRenderPipelineDescriptorClass {
	_MTLTileRenderPipelineDescriptorClassOnce.Do(func() {
		_MTLTileRenderPipelineDescriptorClass = MTLTileRenderPipelineDescriptorClass{class: objc.GetClass("MTLTileRenderPipelineDescriptor")}
	})
	return _MTLTileRenderPipelineDescriptorClass
}

// GetMTLTileRenderPipelineDescriptorClass returns the class object for MTLTileRenderPipelineDescriptor.
func GetMTLTileRenderPipelineDescriptorClass() MTLTileRenderPipelineDescriptorClass {
	return getMTLTileRenderPipelineDescriptorClass()
}

type MTLTileRenderPipelineDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLTileRenderPipelineDescriptorClass) Alloc() MTLTileRenderPipelineDescriptor {
	rv := objc.Send[MTLTileRenderPipelineDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An object that configures new render pipeline state objects for tile
// shading.
//
// # Identifying the render pipeline
//
//   - [MTLTileRenderPipelineDescriptor.Label]: A string that identifies the tile pipeline descriptor.
//   - [MTLTileRenderPipelineDescriptor.SetLabel]
//
// # Specifying graphics functions and associated data
//
//   - [MTLTileRenderPipelineDescriptor.TileFunction]: The compute kernel or fragment function the pipeline calls.
//   - [MTLTileRenderPipelineDescriptor.SetTileFunction]
//   - [MTLTileRenderPipelineDescriptor.TileBuffers]: An array that contains the buffer mutability options for a render pipeline’s tile function.
//   - [MTLTileRenderPipelineDescriptor.MaxCallStackDepth]: The maximum call stack depth for indirect function calls in tile shaders.
//   - [MTLTileRenderPipelineDescriptor.SetMaxCallStackDepth]
//
// # Specifying rasterization and visibility state
//
//   - [MTLTileRenderPipelineDescriptor.ThreadgroupSizeMatchesTileSize]: A Boolean value that indicates whether all threadgroups for this pipeline completely cover tiles.
//   - [MTLTileRenderPipelineDescriptor.SetThreadgroupSizeMatchesTileSize]
//   - [MTLTileRenderPipelineDescriptor.RasterSampleCount]: The number of samples in each fragment.
//   - [MTLTileRenderPipelineDescriptor.SetRasterSampleCount]
//
// # Specifying rendering pipeline state
//
//   - [MTLTileRenderPipelineDescriptor.Reset]: Specifies the default rendering pipeline state values for the descriptor.
//   - [MTLTileRenderPipelineDescriptor.ColorAttachments]: An array of attachments that store color data.
//
// # Specifying threads per threadgroup
//
//   - [MTLTileRenderPipelineDescriptor.MaxTotalThreadsPerThreadgroup]: The maximum number of threads in a threadgroup when dispatching a command using the pipeline.
//   - [MTLTileRenderPipelineDescriptor.SetMaxTotalThreadsPerThreadgroup]
//
// # Specifying precompiled shader binaries
//
//   - [MTLTileRenderPipelineDescriptor.SupportAddingBinaryFunctions]: A Boolean value that indicates whether you can use the pipeline to create new pipelines by adding binary functions to its callable functions list.
//   - [MTLTileRenderPipelineDescriptor.SetSupportAddingBinaryFunctions]
//   - [MTLTileRenderPipelineDescriptor.BinaryArchives]: An array of binary archives to search for precompiled versions of the shader.
//   - [MTLTileRenderPipelineDescriptor.SetBinaryArchives]
//
// # Specifying callable functions for the pipeline
//
//   - [MTLTileRenderPipelineDescriptor.LinkedFunctions]: Functions that you can specify as function arguments for the tile shader when encoding commands that use the pipeline.
//   - [MTLTileRenderPipelineDescriptor.SetLinkedFunctions]
//
// # Specifying shader validation
//
//   - [MTLTileRenderPipelineDescriptor.ShaderValidation]: A value that enables or disables shader validation for the pipeline.
//   - [MTLTileRenderPipelineDescriptor.SetShaderValidation]
//
// # Instance Properties
//
//   - [MTLTileRenderPipelineDescriptor.PreloadedLibraries]
//   - [MTLTileRenderPipelineDescriptor.SetPreloadedLibraries]
//   - [MTLTileRenderPipelineDescriptor.RequiredThreadsPerThreadgroup]
//   - [MTLTileRenderPipelineDescriptor.SetRequiredThreadsPerThreadgroup]
//
// See: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor
type MTLTileRenderPipelineDescriptor struct {
	objectivec.Object
}

// MTLTileRenderPipelineDescriptorFromID constructs a [MTLTileRenderPipelineDescriptor] from an objc.ID.
//
// An object that configures new render pipeline state objects for tile
// shading.
func MTLTileRenderPipelineDescriptorFromID(id objc.ID) MTLTileRenderPipelineDescriptor {
	return MTLTileRenderPipelineDescriptor{objectivec.Object{ID: id}}
}
// NOTE: MTLTileRenderPipelineDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLTileRenderPipelineDescriptor] class.
//
// # Identifying the render pipeline
//
//   - [IMTLTileRenderPipelineDescriptor.Label]: A string that identifies the tile pipeline descriptor.
//   - [IMTLTileRenderPipelineDescriptor.SetLabel]
//
// # Specifying graphics functions and associated data
//
//   - [IMTLTileRenderPipelineDescriptor.TileFunction]: The compute kernel or fragment function the pipeline calls.
//   - [IMTLTileRenderPipelineDescriptor.SetTileFunction]
//   - [IMTLTileRenderPipelineDescriptor.TileBuffers]: An array that contains the buffer mutability options for a render pipeline’s tile function.
//   - [IMTLTileRenderPipelineDescriptor.MaxCallStackDepth]: The maximum call stack depth for indirect function calls in tile shaders.
//   - [IMTLTileRenderPipelineDescriptor.SetMaxCallStackDepth]
//
// # Specifying rasterization and visibility state
//
//   - [IMTLTileRenderPipelineDescriptor.ThreadgroupSizeMatchesTileSize]: A Boolean value that indicates whether all threadgroups for this pipeline completely cover tiles.
//   - [IMTLTileRenderPipelineDescriptor.SetThreadgroupSizeMatchesTileSize]
//   - [IMTLTileRenderPipelineDescriptor.RasterSampleCount]: The number of samples in each fragment.
//   - [IMTLTileRenderPipelineDescriptor.SetRasterSampleCount]
//
// # Specifying rendering pipeline state
//
//   - [IMTLTileRenderPipelineDescriptor.Reset]: Specifies the default rendering pipeline state values for the descriptor.
//   - [IMTLTileRenderPipelineDescriptor.ColorAttachments]: An array of attachments that store color data.
//
// # Specifying threads per threadgroup
//
//   - [IMTLTileRenderPipelineDescriptor.MaxTotalThreadsPerThreadgroup]: The maximum number of threads in a threadgroup when dispatching a command using the pipeline.
//   - [IMTLTileRenderPipelineDescriptor.SetMaxTotalThreadsPerThreadgroup]
//
// # Specifying precompiled shader binaries
//
//   - [IMTLTileRenderPipelineDescriptor.SupportAddingBinaryFunctions]: A Boolean value that indicates whether you can use the pipeline to create new pipelines by adding binary functions to its callable functions list.
//   - [IMTLTileRenderPipelineDescriptor.SetSupportAddingBinaryFunctions]
//   - [IMTLTileRenderPipelineDescriptor.BinaryArchives]: An array of binary archives to search for precompiled versions of the shader.
//   - [IMTLTileRenderPipelineDescriptor.SetBinaryArchives]
//
// # Specifying callable functions for the pipeline
//
//   - [IMTLTileRenderPipelineDescriptor.LinkedFunctions]: Functions that you can specify as function arguments for the tile shader when encoding commands that use the pipeline.
//   - [IMTLTileRenderPipelineDescriptor.SetLinkedFunctions]
//
// # Specifying shader validation
//
//   - [IMTLTileRenderPipelineDescriptor.ShaderValidation]: A value that enables or disables shader validation for the pipeline.
//   - [IMTLTileRenderPipelineDescriptor.SetShaderValidation]
//
// # Instance Properties
//
//   - [IMTLTileRenderPipelineDescriptor.PreloadedLibraries]
//   - [IMTLTileRenderPipelineDescriptor.SetPreloadedLibraries]
//   - [IMTLTileRenderPipelineDescriptor.RequiredThreadsPerThreadgroup]
//   - [IMTLTileRenderPipelineDescriptor.SetRequiredThreadsPerThreadgroup]
//
// See: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor
type IMTLTileRenderPipelineDescriptor interface {
	objectivec.IObject

	// Topic: Identifying the render pipeline

	// A string that identifies the tile pipeline descriptor.
	Label() string
	SetLabel(value string)

	// Topic: Specifying graphics functions and associated data

	// The compute kernel or fragment function the pipeline calls.
	TileFunction() MTLFunction
	SetTileFunction(value MTLFunction)
	// An array that contains the buffer mutability options for a render pipeline’s tile function.
	TileBuffers() IMTLPipelineBufferDescriptorArray
	// The maximum call stack depth for indirect function calls in tile shaders.
	MaxCallStackDepth() uint
	SetMaxCallStackDepth(value uint)

	// Topic: Specifying rasterization and visibility state

	// A Boolean value that indicates whether all threadgroups for this pipeline completely cover tiles.
	ThreadgroupSizeMatchesTileSize() bool
	SetThreadgroupSizeMatchesTileSize(value bool)
	// The number of samples in each fragment.
	RasterSampleCount() uint
	SetRasterSampleCount(value uint)

	// Topic: Specifying rendering pipeline state

	// Specifies the default rendering pipeline state values for the descriptor.
	Reset()
	// An array of attachments that store color data.
	ColorAttachments() IMTLTileRenderPipelineColorAttachmentDescriptorArray

	// Topic: Specifying threads per threadgroup

	// The maximum number of threads in a threadgroup when dispatching a command using the pipeline.
	MaxTotalThreadsPerThreadgroup() uint
	SetMaxTotalThreadsPerThreadgroup(value uint)

	// Topic: Specifying precompiled shader binaries

	// A Boolean value that indicates whether you can use the pipeline to create new pipelines by adding binary functions to its callable functions list.
	SupportAddingBinaryFunctions() bool
	SetSupportAddingBinaryFunctions(value bool)
	// An array of binary archives to search for precompiled versions of the shader.
	BinaryArchives() []objectivec.IObject
	SetBinaryArchives(value []objectivec.IObject)

	// Topic: Specifying callable functions for the pipeline

	// Functions that you can specify as function arguments for the tile shader when encoding commands that use the pipeline.
	LinkedFunctions() IMTLLinkedFunctions
	SetLinkedFunctions(value IMTLLinkedFunctions)

	// Topic: Specifying shader validation

	// A value that enables or disables shader validation for the pipeline.
	ShaderValidation() MTLShaderValidation
	SetShaderValidation(value MTLShaderValidation)

	// Topic: Instance Properties

	PreloadedLibraries() []objectivec.IObject
	SetPreloadedLibraries(value []objectivec.IObject)
	RequiredThreadsPerThreadgroup() MTLSize
	SetRequiredThreadsPerThreadgroup(value MTLSize)
}

// Init initializes the instance.
func (t MTLTileRenderPipelineDescriptor) Init() MTLTileRenderPipelineDescriptor {
	rv := objc.Send[MTLTileRenderPipelineDescriptor](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t MTLTileRenderPipelineDescriptor) Autorelease() MTLTileRenderPipelineDescriptor {
	rv := objc.Send[MTLTileRenderPipelineDescriptor](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLTileRenderPipelineDescriptor creates a new MTLTileRenderPipelineDescriptor instance.
func NewMTLTileRenderPipelineDescriptor() MTLTileRenderPipelineDescriptor {
	class := getMTLTileRenderPipelineDescriptorClass()
	rv := objc.Send[MTLTileRenderPipelineDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Specifies the default rendering pipeline state values for the descriptor.
//
// See: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/reset()
func (t MTLTileRenderPipelineDescriptor) Reset() {
	objc.Send[objc.ID](t.ID, objc.Sel("reset"))
}

// A string that identifies the tile pipeline descriptor.
//
// See: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/label
func (t MTLTileRenderPipelineDescriptor) Label() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}
func (t MTLTileRenderPipelineDescriptor) SetLabel(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setLabel:"), objc.String(value))
}

// The compute kernel or fragment function the pipeline calls.
//
// # Discussion
// 
// Kernel-based and fragment-based tile pipeline dispatches act as a barrier
// against previous draw commands and other dispatches. Kernel-based pipelines
// wait until all prior access to the tile completes. Fragment-based pipelines
// wait only until all prior access to the fragment’s location completes.
//
// See: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/tileFunction
func (t MTLTileRenderPipelineDescriptor) TileFunction() MTLFunction {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("tileFunction"))
	return MTLFunctionObjectFromID(rv)
}
func (t MTLTileRenderPipelineDescriptor) SetTileFunction(value MTLFunction) {
	objc.Send[struct{}](t.ID, objc.Sel("setTileFunction:"), value)
}

// An array that contains the buffer mutability options for a render
// pipeline’s tile function.
//
// # Discussion
// 
// This property returns an array of [MTLPipelineBufferDescriptor] objects,
// with each array index corresponding to the same index in the buffer
// argument table for the render pipeline’s tile shader.
//
// See: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/tileBuffers
func (t MTLTileRenderPipelineDescriptor) TileBuffers() IMTLPipelineBufferDescriptorArray {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("tileBuffers"))
	return MTLPipelineBufferDescriptorArrayFromID(objc.ID(rv))
}

// The maximum call stack depth for indirect function calls in tile shaders.
//
// # Discussion
// 
// The property’s default value is `1`. Change its value if you use
// recursive functions in your tile dispatch.
// 
// The maximum call stack depth applies only to indirect function calls in
// your shader, and affects the upper bound of stack memory for each thread.
// Indirect function calls include those to visible functions, intersection
// functions, and to dynamic libraries.
//
// See: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/maxCallStackDepth
func (t MTLTileRenderPipelineDescriptor) MaxCallStackDepth() uint {
	rv := objc.Send[uint](t.ID, objc.Sel("maxCallStackDepth"))
	return rv
}
func (t MTLTileRenderPipelineDescriptor) SetMaxCallStackDepth(value uint) {
	objc.Send[struct{}](t.ID, objc.Sel("setMaxCallStackDepth:"), value)
}

// A Boolean value that indicates whether all threadgroups for this pipeline
// completely cover tiles.
//
// # Discussion
// 
// Metal can optimize code generation when the threadgroup and tile sizes
// match.
//
// See: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/threadgroupSizeMatchesTileSize
func (t MTLTileRenderPipelineDescriptor) ThreadgroupSizeMatchesTileSize() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("threadgroupSizeMatchesTileSize"))
	return rv
}
func (t MTLTileRenderPipelineDescriptor) SetThreadgroupSizeMatchesTileSize(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setThreadgroupSizeMatchesTileSize:"), value)
}

// The number of samples in each fragment.
//
// # Discussion
// 
// The default value is `1`. This value is used only if the pipeline render
// targets support multisampling. If the render targets don’t support
// multisampling, then this value needs to be `1`.
// 
// When you create a [MTLRenderCommandEncoder], the [SampleCount] value of all
// attachments need to match this `sampleCount` value. Furthermore, the
// texture type of all attachments need to be [TextureType2DMultisample].
// 
// Support for different sample count values varies by device instance. Call
// the [SupportsTextureSampleCount] method on an [MTLDevice] instance to
// determine whether it supports a specific sample count.
//
// See: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/rasterSampleCount
func (t MTLTileRenderPipelineDescriptor) RasterSampleCount() uint {
	rv := objc.Send[uint](t.ID, objc.Sel("rasterSampleCount"))
	return rv
}
func (t MTLTileRenderPipelineDescriptor) SetRasterSampleCount(value uint) {
	objc.Send[struct{}](t.ID, objc.Sel("setRasterSampleCount:"), value)
}

// An array of attachments that store color data.
//
// See: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/colorAttachments
func (t MTLTileRenderPipelineDescriptor) ColorAttachments() IMTLTileRenderPipelineColorAttachmentDescriptorArray {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("colorAttachments"))
	return MTLTileRenderPipelineColorAttachmentDescriptorArrayFromID(objc.ID(rv))
}

// The maximum number of threads in a threadgroup when dispatching a command
// using the pipeline.
//
// See: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/maxTotalThreadsPerThreadgroup
func (t MTLTileRenderPipelineDescriptor) MaxTotalThreadsPerThreadgroup() uint {
	rv := objc.Send[uint](t.ID, objc.Sel("maxTotalThreadsPerThreadgroup"))
	return rv
}
func (t MTLTileRenderPipelineDescriptor) SetMaxTotalThreadsPerThreadgroup(value uint) {
	objc.Send[struct{}](t.ID, objc.Sel("setMaxTotalThreadsPerThreadgroup:"), value)
}

// A Boolean value that indicates whether you can use the pipeline to create
// new pipelines by adding binary functions to its callable functions list.
//
// See: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/supportAddingBinaryFunctions
func (t MTLTileRenderPipelineDescriptor) SupportAddingBinaryFunctions() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("supportAddingBinaryFunctions"))
	return rv
}
func (t MTLTileRenderPipelineDescriptor) SetSupportAddingBinaryFunctions(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setSupportAddingBinaryFunctions:"), value)
}

// An array of binary archives to search for precompiled versions of the
// shader.
//
// See: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/binaryArchives
func (t MTLTileRenderPipelineDescriptor) BinaryArchives() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("binaryArchives"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
func (t MTLTileRenderPipelineDescriptor) SetBinaryArchives(value []objectivec.IObject) {
	objc.Send[struct{}](t.ID, objc.Sel("setBinaryArchives:"), objectivec.IObjectSliceToNSArray(value))
}

// Functions that you can specify as function arguments for the tile shader
// when encoding commands that use the pipeline.
//
// See: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/linkedFunctions
func (t MTLTileRenderPipelineDescriptor) LinkedFunctions() IMTLLinkedFunctions {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("linkedFunctions"))
	return MTLLinkedFunctionsFromID(objc.ID(rv))
}
func (t MTLTileRenderPipelineDescriptor) SetLinkedFunctions(value IMTLLinkedFunctions) {
	objc.Send[struct{}](t.ID, objc.Sel("setLinkedFunctions:"), value)
}

// A value that enables or disables shader validation for the pipeline.
//
// # Discussion
// 
// You can override the value using either of these environment variables:
// `MTL_SHADER_VALIDATION_ENABLE_PIPELINES` or
// `MTL_SHADER_VALIDATION_DISABLE_PIPELINES.`
//
// See: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/shaderValidation
func (t MTLTileRenderPipelineDescriptor) ShaderValidation() MTLShaderValidation {
	rv := objc.Send[MTLShaderValidation](t.ID, objc.Sel("shaderValidation"))
	return MTLShaderValidation(rv)
}
func (t MTLTileRenderPipelineDescriptor) SetShaderValidation(value MTLShaderValidation) {
	objc.Send[struct{}](t.ID, objc.Sel("setShaderValidation:"), value)
}

// See: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/preloadedLibraries
func (t MTLTileRenderPipelineDescriptor) PreloadedLibraries() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("preloadedLibraries"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
func (t MTLTileRenderPipelineDescriptor) SetPreloadedLibraries(value []objectivec.IObject) {
	objc.Send[struct{}](t.ID, objc.Sel("setPreloadedLibraries:"), objectivec.IObjectSliceToNSArray(value))
}

//
// # Discussion
// 
// Sets the required threads-per-threadgroup during tile dispatches. The
// `threadsPerTile` argument of any tile dispatch must match to this value if
// it is set. Optional, unless the pipeline is going to use CooperativeTensors
// in which case this must be set. Setting this to a size of 0 in every
// dimension disables this property
//
// See: https://developer.apple.com/documentation/Metal/MTLTileRenderPipelineDescriptor/requiredThreadsPerThreadgroup
func (t MTLTileRenderPipelineDescriptor) RequiredThreadsPerThreadgroup() MTLSize {
	rv := objc.Send[MTLSize](t.ID, objc.Sel("requiredThreadsPerThreadgroup"))
	return MTLSize(rv)
}
func (t MTLTileRenderPipelineDescriptor) SetRequiredThreadsPerThreadgroup(value MTLSize) {
	objc.Send[struct{}](t.ID, objc.Sel("setRequiredThreadsPerThreadgroup:"), value)
}

