// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLMeshRenderPipelineDescriptor] class.
var (
	_MTLMeshRenderPipelineDescriptorClass     MTLMeshRenderPipelineDescriptorClass
	_MTLMeshRenderPipelineDescriptorClassOnce sync.Once
)

func getMTLMeshRenderPipelineDescriptorClass() MTLMeshRenderPipelineDescriptorClass {
	_MTLMeshRenderPipelineDescriptorClassOnce.Do(func() {
		_MTLMeshRenderPipelineDescriptorClass = MTLMeshRenderPipelineDescriptorClass{class: objc.GetClass("MTLMeshRenderPipelineDescriptor")}
	})
	return _MTLMeshRenderPipelineDescriptorClass
}

// GetMTLMeshRenderPipelineDescriptorClass returns the class object for MTLMeshRenderPipelineDescriptor.
func GetMTLMeshRenderPipelineDescriptorClass() MTLMeshRenderPipelineDescriptorClass {
	return getMTLMeshRenderPipelineDescriptorClass()
}

type MTLMeshRenderPipelineDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLMeshRenderPipelineDescriptorClass) Alloc() MTLMeshRenderPipelineDescriptor {
	rv := objc.Send[MTLMeshRenderPipelineDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// An object that configures new render pipeline state objects for mesh
// shading.
//
// # Instance Properties
//
//   - [MTLMeshRenderPipelineDescriptor.BinaryArchives]
//   - [MTLMeshRenderPipelineDescriptor.SetBinaryArchives]
//   - [MTLMeshRenderPipelineDescriptor.ColorAttachments]
//   - [MTLMeshRenderPipelineDescriptor.DepthAttachmentPixelFormat]
//   - [MTLMeshRenderPipelineDescriptor.SetDepthAttachmentPixelFormat]
//   - [MTLMeshRenderPipelineDescriptor.FragmentBuffers]
//   - [MTLMeshRenderPipelineDescriptor.FragmentFunction]
//   - [MTLMeshRenderPipelineDescriptor.SetFragmentFunction]
//   - [MTLMeshRenderPipelineDescriptor.FragmentLinkedFunctions]
//   - [MTLMeshRenderPipelineDescriptor.SetFragmentLinkedFunctions]
//   - [MTLMeshRenderPipelineDescriptor.AlphaToCoverageEnabled]
//   - [MTLMeshRenderPipelineDescriptor.SetAlphaToCoverageEnabled]
//   - [MTLMeshRenderPipelineDescriptor.AlphaToOneEnabled]
//   - [MTLMeshRenderPipelineDescriptor.SetAlphaToOneEnabled]
//   - [MTLMeshRenderPipelineDescriptor.RasterizationEnabled]
//   - [MTLMeshRenderPipelineDescriptor.SetRasterizationEnabled]
//   - [MTLMeshRenderPipelineDescriptor.Label]
//   - [MTLMeshRenderPipelineDescriptor.SetLabel]
//   - [MTLMeshRenderPipelineDescriptor.MaxTotalThreadgroupsPerMeshGrid]
//   - [MTLMeshRenderPipelineDescriptor.SetMaxTotalThreadgroupsPerMeshGrid]
//   - [MTLMeshRenderPipelineDescriptor.MaxTotalThreadsPerMeshThreadgroup]
//   - [MTLMeshRenderPipelineDescriptor.SetMaxTotalThreadsPerMeshThreadgroup]
//   - [MTLMeshRenderPipelineDescriptor.MaxTotalThreadsPerObjectThreadgroup]
//   - [MTLMeshRenderPipelineDescriptor.SetMaxTotalThreadsPerObjectThreadgroup]
//   - [MTLMeshRenderPipelineDescriptor.MaxVertexAmplificationCount]
//   - [MTLMeshRenderPipelineDescriptor.SetMaxVertexAmplificationCount]
//   - [MTLMeshRenderPipelineDescriptor.MeshBuffers]
//   - [MTLMeshRenderPipelineDescriptor.MeshFunction]
//   - [MTLMeshRenderPipelineDescriptor.SetMeshFunction]
//   - [MTLMeshRenderPipelineDescriptor.MeshLinkedFunctions]
//   - [MTLMeshRenderPipelineDescriptor.SetMeshLinkedFunctions]
//   - [MTLMeshRenderPipelineDescriptor.MeshThreadgroupSizeIsMultipleOfThreadExecutionWidth]
//   - [MTLMeshRenderPipelineDescriptor.SetMeshThreadgroupSizeIsMultipleOfThreadExecutionWidth]
//   - [MTLMeshRenderPipelineDescriptor.ObjectBuffers]
//   - [MTLMeshRenderPipelineDescriptor.ObjectFunction]
//   - [MTLMeshRenderPipelineDescriptor.SetObjectFunction]
//   - [MTLMeshRenderPipelineDescriptor.ObjectLinkedFunctions]
//   - [MTLMeshRenderPipelineDescriptor.SetObjectLinkedFunctions]
//   - [MTLMeshRenderPipelineDescriptor.ObjectThreadgroupSizeIsMultipleOfThreadExecutionWidth]
//   - [MTLMeshRenderPipelineDescriptor.SetObjectThreadgroupSizeIsMultipleOfThreadExecutionWidth]
//   - [MTLMeshRenderPipelineDescriptor.PayloadMemoryLength]
//   - [MTLMeshRenderPipelineDescriptor.SetPayloadMemoryLength]
//   - [MTLMeshRenderPipelineDescriptor.RasterSampleCount]
//   - [MTLMeshRenderPipelineDescriptor.SetRasterSampleCount]
//   - [MTLMeshRenderPipelineDescriptor.RequiredThreadsPerMeshThreadgroup]
//   - [MTLMeshRenderPipelineDescriptor.SetRequiredThreadsPerMeshThreadgroup]
//   - [MTLMeshRenderPipelineDescriptor.RequiredThreadsPerObjectThreadgroup]
//   - [MTLMeshRenderPipelineDescriptor.SetRequiredThreadsPerObjectThreadgroup]
//   - [MTLMeshRenderPipelineDescriptor.ShaderValidation]: A value that enables or disables shader validation for the pipeline.
//   - [MTLMeshRenderPipelineDescriptor.SetShaderValidation]
//   - [MTLMeshRenderPipelineDescriptor.StencilAttachmentPixelFormat]
//   - [MTLMeshRenderPipelineDescriptor.SetStencilAttachmentPixelFormat]
//   - [MTLMeshRenderPipelineDescriptor.SupportIndirectCommandBuffers]
//   - [MTLMeshRenderPipelineDescriptor.SetSupportIndirectCommandBuffers]
//
// # Instance Methods
//
//   - [MTLMeshRenderPipelineDescriptor.Reset]
//
// See: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor
type MTLMeshRenderPipelineDescriptor struct {
	objectivec.Object
}

// MTLMeshRenderPipelineDescriptorFromID constructs a [MTLMeshRenderPipelineDescriptor] from an objc.ID.
//
// An object that configures new render pipeline state objects for mesh
// shading.
func MTLMeshRenderPipelineDescriptorFromID(id objc.ID) MTLMeshRenderPipelineDescriptor {
	return MTLMeshRenderPipelineDescriptor{objectivec.Object{ID: id}}
}
// NOTE: MTLMeshRenderPipelineDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLMeshRenderPipelineDescriptor] class.
//
// # Instance Properties
//
//   - [IMTLMeshRenderPipelineDescriptor.BinaryArchives]
//   - [IMTLMeshRenderPipelineDescriptor.SetBinaryArchives]
//   - [IMTLMeshRenderPipelineDescriptor.ColorAttachments]
//   - [IMTLMeshRenderPipelineDescriptor.DepthAttachmentPixelFormat]
//   - [IMTLMeshRenderPipelineDescriptor.SetDepthAttachmentPixelFormat]
//   - [IMTLMeshRenderPipelineDescriptor.FragmentBuffers]
//   - [IMTLMeshRenderPipelineDescriptor.FragmentFunction]
//   - [IMTLMeshRenderPipelineDescriptor.SetFragmentFunction]
//   - [IMTLMeshRenderPipelineDescriptor.FragmentLinkedFunctions]
//   - [IMTLMeshRenderPipelineDescriptor.SetFragmentLinkedFunctions]
//   - [IMTLMeshRenderPipelineDescriptor.AlphaToCoverageEnabled]
//   - [IMTLMeshRenderPipelineDescriptor.SetAlphaToCoverageEnabled]
//   - [IMTLMeshRenderPipelineDescriptor.AlphaToOneEnabled]
//   - [IMTLMeshRenderPipelineDescriptor.SetAlphaToOneEnabled]
//   - [IMTLMeshRenderPipelineDescriptor.RasterizationEnabled]
//   - [IMTLMeshRenderPipelineDescriptor.SetRasterizationEnabled]
//   - [IMTLMeshRenderPipelineDescriptor.Label]
//   - [IMTLMeshRenderPipelineDescriptor.SetLabel]
//   - [IMTLMeshRenderPipelineDescriptor.MaxTotalThreadgroupsPerMeshGrid]
//   - [IMTLMeshRenderPipelineDescriptor.SetMaxTotalThreadgroupsPerMeshGrid]
//   - [IMTLMeshRenderPipelineDescriptor.MaxTotalThreadsPerMeshThreadgroup]
//   - [IMTLMeshRenderPipelineDescriptor.SetMaxTotalThreadsPerMeshThreadgroup]
//   - [IMTLMeshRenderPipelineDescriptor.MaxTotalThreadsPerObjectThreadgroup]
//   - [IMTLMeshRenderPipelineDescriptor.SetMaxTotalThreadsPerObjectThreadgroup]
//   - [IMTLMeshRenderPipelineDescriptor.MaxVertexAmplificationCount]
//   - [IMTLMeshRenderPipelineDescriptor.SetMaxVertexAmplificationCount]
//   - [IMTLMeshRenderPipelineDescriptor.MeshBuffers]
//   - [IMTLMeshRenderPipelineDescriptor.MeshFunction]
//   - [IMTLMeshRenderPipelineDescriptor.SetMeshFunction]
//   - [IMTLMeshRenderPipelineDescriptor.MeshLinkedFunctions]
//   - [IMTLMeshRenderPipelineDescriptor.SetMeshLinkedFunctions]
//   - [IMTLMeshRenderPipelineDescriptor.MeshThreadgroupSizeIsMultipleOfThreadExecutionWidth]
//   - [IMTLMeshRenderPipelineDescriptor.SetMeshThreadgroupSizeIsMultipleOfThreadExecutionWidth]
//   - [IMTLMeshRenderPipelineDescriptor.ObjectBuffers]
//   - [IMTLMeshRenderPipelineDescriptor.ObjectFunction]
//   - [IMTLMeshRenderPipelineDescriptor.SetObjectFunction]
//   - [IMTLMeshRenderPipelineDescriptor.ObjectLinkedFunctions]
//   - [IMTLMeshRenderPipelineDescriptor.SetObjectLinkedFunctions]
//   - [IMTLMeshRenderPipelineDescriptor.ObjectThreadgroupSizeIsMultipleOfThreadExecutionWidth]
//   - [IMTLMeshRenderPipelineDescriptor.SetObjectThreadgroupSizeIsMultipleOfThreadExecutionWidth]
//   - [IMTLMeshRenderPipelineDescriptor.PayloadMemoryLength]
//   - [IMTLMeshRenderPipelineDescriptor.SetPayloadMemoryLength]
//   - [IMTLMeshRenderPipelineDescriptor.RasterSampleCount]
//   - [IMTLMeshRenderPipelineDescriptor.SetRasterSampleCount]
//   - [IMTLMeshRenderPipelineDescriptor.RequiredThreadsPerMeshThreadgroup]
//   - [IMTLMeshRenderPipelineDescriptor.SetRequiredThreadsPerMeshThreadgroup]
//   - [IMTLMeshRenderPipelineDescriptor.RequiredThreadsPerObjectThreadgroup]
//   - [IMTLMeshRenderPipelineDescriptor.SetRequiredThreadsPerObjectThreadgroup]
//   - [IMTLMeshRenderPipelineDescriptor.ShaderValidation]: A value that enables or disables shader validation for the pipeline.
//   - [IMTLMeshRenderPipelineDescriptor.SetShaderValidation]
//   - [IMTLMeshRenderPipelineDescriptor.StencilAttachmentPixelFormat]
//   - [IMTLMeshRenderPipelineDescriptor.SetStencilAttachmentPixelFormat]
//   - [IMTLMeshRenderPipelineDescriptor.SupportIndirectCommandBuffers]
//   - [IMTLMeshRenderPipelineDescriptor.SetSupportIndirectCommandBuffers]
//
// # Instance Methods
//
//   - [IMTLMeshRenderPipelineDescriptor.Reset]
//
// See: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor
type IMTLMeshRenderPipelineDescriptor interface {
	objectivec.IObject

	// Topic: Instance Properties

	BinaryArchives() []objectivec.IObject
	SetBinaryArchives(value []objectivec.IObject)
	ColorAttachments() IMTLRenderPipelineColorAttachmentDescriptorArray
	DepthAttachmentPixelFormat() MTLPixelFormat
	SetDepthAttachmentPixelFormat(value MTLPixelFormat)
	FragmentBuffers() IMTLPipelineBufferDescriptorArray
	FragmentFunction() MTLFunction
	SetFragmentFunction(value MTLFunction)
	FragmentLinkedFunctions() IMTLLinkedFunctions
	SetFragmentLinkedFunctions(value IMTLLinkedFunctions)
	AlphaToCoverageEnabled() bool
	SetAlphaToCoverageEnabled(value bool)
	AlphaToOneEnabled() bool
	SetAlphaToOneEnabled(value bool)
	RasterizationEnabled() bool
	SetRasterizationEnabled(value bool)
	Label() string
	SetLabel(value string)
	MaxTotalThreadgroupsPerMeshGrid() uint
	SetMaxTotalThreadgroupsPerMeshGrid(value uint)
	MaxTotalThreadsPerMeshThreadgroup() uint
	SetMaxTotalThreadsPerMeshThreadgroup(value uint)
	MaxTotalThreadsPerObjectThreadgroup() uint
	SetMaxTotalThreadsPerObjectThreadgroup(value uint)
	MaxVertexAmplificationCount() uint
	SetMaxVertexAmplificationCount(value uint)
	MeshBuffers() IMTLPipelineBufferDescriptorArray
	MeshFunction() MTLFunction
	SetMeshFunction(value MTLFunction)
	MeshLinkedFunctions() IMTLLinkedFunctions
	SetMeshLinkedFunctions(value IMTLLinkedFunctions)
	MeshThreadgroupSizeIsMultipleOfThreadExecutionWidth() bool
	SetMeshThreadgroupSizeIsMultipleOfThreadExecutionWidth(value bool)
	ObjectBuffers() IMTLPipelineBufferDescriptorArray
	ObjectFunction() MTLFunction
	SetObjectFunction(value MTLFunction)
	ObjectLinkedFunctions() IMTLLinkedFunctions
	SetObjectLinkedFunctions(value IMTLLinkedFunctions)
	ObjectThreadgroupSizeIsMultipleOfThreadExecutionWidth() bool
	SetObjectThreadgroupSizeIsMultipleOfThreadExecutionWidth(value bool)
	PayloadMemoryLength() uint
	SetPayloadMemoryLength(value uint)
	RasterSampleCount() uint
	SetRasterSampleCount(value uint)
	RequiredThreadsPerMeshThreadgroup() MTLSize
	SetRequiredThreadsPerMeshThreadgroup(value MTLSize)
	RequiredThreadsPerObjectThreadgroup() MTLSize
	SetRequiredThreadsPerObjectThreadgroup(value MTLSize)
	// A value that enables or disables shader validation for the pipeline.
	ShaderValidation() MTLShaderValidation
	SetShaderValidation(value MTLShaderValidation)
	StencilAttachmentPixelFormat() MTLPixelFormat
	SetStencilAttachmentPixelFormat(value MTLPixelFormat)
	SupportIndirectCommandBuffers() bool
	SetSupportIndirectCommandBuffers(value bool)

	// Topic: Instance Methods

	Reset()
}

// Init initializes the instance.
func (m MTLMeshRenderPipelineDescriptor) Init() MTLMeshRenderPipelineDescriptor {
	rv := objc.Send[MTLMeshRenderPipelineDescriptor](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MTLMeshRenderPipelineDescriptor) Autorelease() MTLMeshRenderPipelineDescriptor {
	rv := objc.Send[MTLMeshRenderPipelineDescriptor](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLMeshRenderPipelineDescriptor creates a new MTLMeshRenderPipelineDescriptor instance.
func NewMTLMeshRenderPipelineDescriptor() MTLMeshRenderPipelineDescriptor {
	class := getMTLMeshRenderPipelineDescriptorClass()
	rv := objc.Send[MTLMeshRenderPipelineDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/reset()
func (m MTLMeshRenderPipelineDescriptor) Reset() {
	objc.Send[objc.ID](m.ID, objc.Sel("reset"))
}

// See: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/binaryArchives
func (m MTLMeshRenderPipelineDescriptor) BinaryArchives() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("binaryArchives"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
func (m MTLMeshRenderPipelineDescriptor) SetBinaryArchives(value []objectivec.IObject) {
	objc.Send[struct{}](m.ID, objc.Sel("setBinaryArchives:"), objectivec.IObjectSliceToNSArray(value))
}
// See: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/colorAttachments
func (m MTLMeshRenderPipelineDescriptor) ColorAttachments() IMTLRenderPipelineColorAttachmentDescriptorArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("colorAttachments"))
	return MTLRenderPipelineColorAttachmentDescriptorArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/depthAttachmentPixelFormat
func (m MTLMeshRenderPipelineDescriptor) DepthAttachmentPixelFormat() MTLPixelFormat {
	rv := objc.Send[MTLPixelFormat](m.ID, objc.Sel("depthAttachmentPixelFormat"))
	return MTLPixelFormat(rv)
}
func (m MTLMeshRenderPipelineDescriptor) SetDepthAttachmentPixelFormat(value MTLPixelFormat) {
	objc.Send[struct{}](m.ID, objc.Sel("setDepthAttachmentPixelFormat:"), value)
}
// See: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/fragmentBuffers
func (m MTLMeshRenderPipelineDescriptor) FragmentBuffers() IMTLPipelineBufferDescriptorArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("fragmentBuffers"))
	return MTLPipelineBufferDescriptorArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/fragmentFunction
func (m MTLMeshRenderPipelineDescriptor) FragmentFunction() MTLFunction {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("fragmentFunction"))
	return MTLFunctionObjectFromID(rv)
}
func (m MTLMeshRenderPipelineDescriptor) SetFragmentFunction(value MTLFunction) {
	objc.Send[struct{}](m.ID, objc.Sel("setFragmentFunction:"), value)
}
// See: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/fragmentLinkedFunctions
func (m MTLMeshRenderPipelineDescriptor) FragmentLinkedFunctions() IMTLLinkedFunctions {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("fragmentLinkedFunctions"))
	return MTLLinkedFunctionsFromID(objc.ID(rv))
}
func (m MTLMeshRenderPipelineDescriptor) SetFragmentLinkedFunctions(value IMTLLinkedFunctions) {
	objc.Send[struct{}](m.ID, objc.Sel("setFragmentLinkedFunctions:"), value)
}
// See: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/isAlphaToCoverageEnabled
func (m MTLMeshRenderPipelineDescriptor) AlphaToCoverageEnabled() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isAlphaToCoverageEnabled"))
	return rv
}
func (m MTLMeshRenderPipelineDescriptor) SetAlphaToCoverageEnabled(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setAlphaToCoverageEnabled:"), value)
}
// See: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/isAlphaToOneEnabled
func (m MTLMeshRenderPipelineDescriptor) AlphaToOneEnabled() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isAlphaToOneEnabled"))
	return rv
}
func (m MTLMeshRenderPipelineDescriptor) SetAlphaToOneEnabled(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setAlphaToOneEnabled:"), value)
}
// See: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/isRasterizationEnabled
func (m MTLMeshRenderPipelineDescriptor) RasterizationEnabled() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isRasterizationEnabled"))
	return rv
}
func (m MTLMeshRenderPipelineDescriptor) SetRasterizationEnabled(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setRasterizationEnabled:"), value)
}
// See: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/label
func (m MTLMeshRenderPipelineDescriptor) Label() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}
func (m MTLMeshRenderPipelineDescriptor) SetLabel(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setLabel:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/maxTotalThreadgroupsPerMeshGrid
func (m MTLMeshRenderPipelineDescriptor) MaxTotalThreadgroupsPerMeshGrid() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("maxTotalThreadgroupsPerMeshGrid"))
	return rv
}
func (m MTLMeshRenderPipelineDescriptor) SetMaxTotalThreadgroupsPerMeshGrid(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setMaxTotalThreadgroupsPerMeshGrid:"), value)
}
// See: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/maxTotalThreadsPerMeshThreadgroup
func (m MTLMeshRenderPipelineDescriptor) MaxTotalThreadsPerMeshThreadgroup() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("maxTotalThreadsPerMeshThreadgroup"))
	return rv
}
func (m MTLMeshRenderPipelineDescriptor) SetMaxTotalThreadsPerMeshThreadgroup(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setMaxTotalThreadsPerMeshThreadgroup:"), value)
}
// See: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/maxTotalThreadsPerObjectThreadgroup
func (m MTLMeshRenderPipelineDescriptor) MaxTotalThreadsPerObjectThreadgroup() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("maxTotalThreadsPerObjectThreadgroup"))
	return rv
}
func (m MTLMeshRenderPipelineDescriptor) SetMaxTotalThreadsPerObjectThreadgroup(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setMaxTotalThreadsPerObjectThreadgroup:"), value)
}
// See: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/maxVertexAmplificationCount
func (m MTLMeshRenderPipelineDescriptor) MaxVertexAmplificationCount() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("maxVertexAmplificationCount"))
	return rv
}
func (m MTLMeshRenderPipelineDescriptor) SetMaxVertexAmplificationCount(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setMaxVertexAmplificationCount:"), value)
}
// See: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/meshBuffers
func (m MTLMeshRenderPipelineDescriptor) MeshBuffers() IMTLPipelineBufferDescriptorArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("meshBuffers"))
	return MTLPipelineBufferDescriptorArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/meshFunction
func (m MTLMeshRenderPipelineDescriptor) MeshFunction() MTLFunction {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("meshFunction"))
	return MTLFunctionObjectFromID(rv)
}
func (m MTLMeshRenderPipelineDescriptor) SetMeshFunction(value MTLFunction) {
	objc.Send[struct{}](m.ID, objc.Sel("setMeshFunction:"), value)
}
// See: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/meshLinkedFunctions
func (m MTLMeshRenderPipelineDescriptor) MeshLinkedFunctions() IMTLLinkedFunctions {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("meshLinkedFunctions"))
	return MTLLinkedFunctionsFromID(objc.ID(rv))
}
func (m MTLMeshRenderPipelineDescriptor) SetMeshLinkedFunctions(value IMTLLinkedFunctions) {
	objc.Send[struct{}](m.ID, objc.Sel("setMeshLinkedFunctions:"), value)
}
// See: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/meshThreadgroupSizeIsMultipleOfThreadExecutionWidth
func (m MTLMeshRenderPipelineDescriptor) MeshThreadgroupSizeIsMultipleOfThreadExecutionWidth() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("meshThreadgroupSizeIsMultipleOfThreadExecutionWidth"))
	return rv
}
func (m MTLMeshRenderPipelineDescriptor) SetMeshThreadgroupSizeIsMultipleOfThreadExecutionWidth(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setMeshThreadgroupSizeIsMultipleOfThreadExecutionWidth:"), value)
}
// See: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/objectBuffers
func (m MTLMeshRenderPipelineDescriptor) ObjectBuffers() IMTLPipelineBufferDescriptorArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("objectBuffers"))
	return MTLPipelineBufferDescriptorArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/objectFunction
func (m MTLMeshRenderPipelineDescriptor) ObjectFunction() MTLFunction {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("objectFunction"))
	return MTLFunctionObjectFromID(rv)
}
func (m MTLMeshRenderPipelineDescriptor) SetObjectFunction(value MTLFunction) {
	objc.Send[struct{}](m.ID, objc.Sel("setObjectFunction:"), value)
}
// See: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/objectLinkedFunctions
func (m MTLMeshRenderPipelineDescriptor) ObjectLinkedFunctions() IMTLLinkedFunctions {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("objectLinkedFunctions"))
	return MTLLinkedFunctionsFromID(objc.ID(rv))
}
func (m MTLMeshRenderPipelineDescriptor) SetObjectLinkedFunctions(value IMTLLinkedFunctions) {
	objc.Send[struct{}](m.ID, objc.Sel("setObjectLinkedFunctions:"), value)
}
// See: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/objectThreadgroupSizeIsMultipleOfThreadExecutionWidth
func (m MTLMeshRenderPipelineDescriptor) ObjectThreadgroupSizeIsMultipleOfThreadExecutionWidth() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("objectThreadgroupSizeIsMultipleOfThreadExecutionWidth"))
	return rv
}
func (m MTLMeshRenderPipelineDescriptor) SetObjectThreadgroupSizeIsMultipleOfThreadExecutionWidth(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setObjectThreadgroupSizeIsMultipleOfThreadExecutionWidth:"), value)
}
// See: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/payloadMemoryLength
func (m MTLMeshRenderPipelineDescriptor) PayloadMemoryLength() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("payloadMemoryLength"))
	return rv
}
func (m MTLMeshRenderPipelineDescriptor) SetPayloadMemoryLength(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setPayloadMemoryLength:"), value)
}
// See: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/rasterSampleCount
func (m MTLMeshRenderPipelineDescriptor) RasterSampleCount() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("rasterSampleCount"))
	return rv
}
func (m MTLMeshRenderPipelineDescriptor) SetRasterSampleCount(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setRasterSampleCount:"), value)
}
//
// # Discussion
// 
// Sets the required mesh threads-per-threadgroup during mesh draws. The
// `threadsPerMeshThreadgroup` argument of any draw must match to this value
// if it is set. Optional, unless the pipeline is going to use
// CooperativeTensors in which case this must be set. Setting this to a size
// of 0 in every dimension disables this property
//
// See: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/requiredThreadsPerMeshThreadgroup
func (m MTLMeshRenderPipelineDescriptor) RequiredThreadsPerMeshThreadgroup() MTLSize {
	rv := objc.Send[MTLSize](m.ID, objc.Sel("requiredThreadsPerMeshThreadgroup"))
	return MTLSize(rv)
}
func (m MTLMeshRenderPipelineDescriptor) SetRequiredThreadsPerMeshThreadgroup(value MTLSize) {
	objc.Send[struct{}](m.ID, objc.Sel("setRequiredThreadsPerMeshThreadgroup:"), value)
}
//
// # Discussion
// 
// Sets the required object threads-per-threadgroup during mesh draws. The
// `threadsPerObjectThreadgroup` argument of any draw must match to this value
// if it is set. Optional, unless the pipeline is going to use
// CooperativeTensors in which case this must be set. Setting this to a size
// of 0 in every dimension disables this property
//
// See: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/requiredThreadsPerObjectThreadgroup
func (m MTLMeshRenderPipelineDescriptor) RequiredThreadsPerObjectThreadgroup() MTLSize {
	rv := objc.Send[MTLSize](m.ID, objc.Sel("requiredThreadsPerObjectThreadgroup"))
	return MTLSize(rv)
}
func (m MTLMeshRenderPipelineDescriptor) SetRequiredThreadsPerObjectThreadgroup(value MTLSize) {
	objc.Send[struct{}](m.ID, objc.Sel("setRequiredThreadsPerObjectThreadgroup:"), value)
}
// A value that enables or disables shader validation for the pipeline.
//
// # Discussion
// 
// You can override the value using either of these environment variables:
// `MTL_SHADER_VALIDATION_ENABLE_PIPELINES` or
// `MTL_SHADER_VALIDATION_DISABLE_PIPELINES.`
//
// See: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/shaderValidation
func (m MTLMeshRenderPipelineDescriptor) ShaderValidation() MTLShaderValidation {
	rv := objc.Send[MTLShaderValidation](m.ID, objc.Sel("shaderValidation"))
	return MTLShaderValidation(rv)
}
func (m MTLMeshRenderPipelineDescriptor) SetShaderValidation(value MTLShaderValidation) {
	objc.Send[struct{}](m.ID, objc.Sel("setShaderValidation:"), value)
}
// See: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/stencilAttachmentPixelFormat
func (m MTLMeshRenderPipelineDescriptor) StencilAttachmentPixelFormat() MTLPixelFormat {
	rv := objc.Send[MTLPixelFormat](m.ID, objc.Sel("stencilAttachmentPixelFormat"))
	return MTLPixelFormat(rv)
}
func (m MTLMeshRenderPipelineDescriptor) SetStencilAttachmentPixelFormat(value MTLPixelFormat) {
	objc.Send[struct{}](m.ID, objc.Sel("setStencilAttachmentPixelFormat:"), value)
}
// See: https://developer.apple.com/documentation/Metal/MTLMeshRenderPipelineDescriptor/supportIndirectCommandBuffers
func (m MTLMeshRenderPipelineDescriptor) SupportIndirectCommandBuffers() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("supportIndirectCommandBuffers"))
	return rv
}
func (m MTLMeshRenderPipelineDescriptor) SetSupportIndirectCommandBuffers(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setSupportIndirectCommandBuffers:"), value)
}

