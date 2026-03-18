// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLIndirectCommandBufferDescriptor] class.
var (
	_MTLIndirectCommandBufferDescriptorClass     MTLIndirectCommandBufferDescriptorClass
	_MTLIndirectCommandBufferDescriptorClassOnce sync.Once
)

func getMTLIndirectCommandBufferDescriptorClass() MTLIndirectCommandBufferDescriptorClass {
	_MTLIndirectCommandBufferDescriptorClassOnce.Do(func() {
		_MTLIndirectCommandBufferDescriptorClass = MTLIndirectCommandBufferDescriptorClass{class: objc.GetClass("MTLIndirectCommandBufferDescriptor")}
	})
	return _MTLIndirectCommandBufferDescriptorClass
}

// GetMTLIndirectCommandBufferDescriptorClass returns the class object for MTLIndirectCommandBufferDescriptor.
func GetMTLIndirectCommandBufferDescriptorClass() MTLIndirectCommandBufferDescriptorClass {
	return getMTLIndirectCommandBufferDescriptorClass()
}

type MTLIndirectCommandBufferDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLIndirectCommandBufferDescriptorClass) Alloc() MTLIndirectCommandBufferDescriptor {
	rv := objc.Send[MTLIndirectCommandBufferDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A configuration you create to customize an indirect command buffer.
//
// # Declaring command types to encode
//
//   - [MTLIndirectCommandBufferDescriptor.CommandTypes]: The set of command types that you can encode into the indirect command buffer.
//   - [MTLIndirectCommandBufferDescriptor.SetCommandTypes]
//
// # Declaring command inheritance
//
//   - [MTLIndirectCommandBufferDescriptor.InheritBuffers]: A Boolean value that determines where commands in the indirect command buffer get their buffer arguments from when you execute them.
//   - [MTLIndirectCommandBufferDescriptor.SetInheritBuffers]
//   - [MTLIndirectCommandBufferDescriptor.InheritPipelineState]: A Boolean value that determines where commands in the indirect command buffer get their pipeline state from when you execute them.
//   - [MTLIndirectCommandBufferDescriptor.SetInheritPipelineState]
//
// # Declaring the maximum number of argument buffers per command
//
//   - [MTLIndirectCommandBufferDescriptor.MaxVertexBufferBindCount]: The maximum number of buffers that you can set per command for the vertex stage.
//   - [MTLIndirectCommandBufferDescriptor.SetMaxVertexBufferBindCount]
//   - [MTLIndirectCommandBufferDescriptor.MaxFragmentBufferBindCount]: The maximum number of buffers that you can set per command for the fragment stage.
//   - [MTLIndirectCommandBufferDescriptor.SetMaxFragmentBufferBindCount]
//   - [MTLIndirectCommandBufferDescriptor.MaxKernelBufferBindCount]: The maximum number of buffers that you can set per command for the compute kernel.
//   - [MTLIndirectCommandBufferDescriptor.SetMaxKernelBufferBindCount]
//
// # Instance Properties
//
//   - [MTLIndirectCommandBufferDescriptor.InheritCullMode]: Configures whether the indirect command buffer inherits the cull mode from the encoder.
//   - [MTLIndirectCommandBufferDescriptor.SetInheritCullMode]
//   - [MTLIndirectCommandBufferDescriptor.InheritDepthBias]: Configures whether the indirect command buffer inherits the depth bias from the encoder.
//   - [MTLIndirectCommandBufferDescriptor.SetInheritDepthBias]
//   - [MTLIndirectCommandBufferDescriptor.InheritDepthClipMode]: Configures whether the indirect command buffer inherits the depth clip mode from the encoder.
//   - [MTLIndirectCommandBufferDescriptor.SetInheritDepthClipMode]
//   - [MTLIndirectCommandBufferDescriptor.InheritDepthStencilState]: Configures whether the indirect command buffer inherits the depth stencil state from the encoder.
//   - [MTLIndirectCommandBufferDescriptor.SetInheritDepthStencilState]
//   - [MTLIndirectCommandBufferDescriptor.InheritFrontFacingWinding]: Configures whether the indirect command buffer inherits the front facing winding from the encoder.
//   - [MTLIndirectCommandBufferDescriptor.SetInheritFrontFacingWinding]
//   - [MTLIndirectCommandBufferDescriptor.InheritTriangleFillMode]: Configures whether the indirect command buffer inherits the triangle fill mode from the encoder.
//   - [MTLIndirectCommandBufferDescriptor.SetInheritTriangleFillMode]
//   - [MTLIndirectCommandBufferDescriptor.MaxKernelThreadgroupMemoryBindCount]
//   - [MTLIndirectCommandBufferDescriptor.SetMaxKernelThreadgroupMemoryBindCount]
//   - [MTLIndirectCommandBufferDescriptor.MaxMeshBufferBindCount]
//   - [MTLIndirectCommandBufferDescriptor.SetMaxMeshBufferBindCount]
//   - [MTLIndirectCommandBufferDescriptor.MaxObjectBufferBindCount]
//   - [MTLIndirectCommandBufferDescriptor.SetMaxObjectBufferBindCount]
//   - [MTLIndirectCommandBufferDescriptor.MaxObjectThreadgroupMemoryBindCount]
//   - [MTLIndirectCommandBufferDescriptor.SetMaxObjectThreadgroupMemoryBindCount]
//   - [MTLIndirectCommandBufferDescriptor.SupportColorAttachmentMapping]: Specifies if the indirect command buffer should support color attachment mapping.
//   - [MTLIndirectCommandBufferDescriptor.SetSupportColorAttachmentMapping]
//   - [MTLIndirectCommandBufferDescriptor.SupportDynamicAttributeStride]
//   - [MTLIndirectCommandBufferDescriptor.SetSupportDynamicAttributeStride]
//   - [MTLIndirectCommandBufferDescriptor.SupportRayTracing]
//   - [MTLIndirectCommandBufferDescriptor.SetSupportRayTracing]
//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor
type MTLIndirectCommandBufferDescriptor struct {
	objectivec.Object
}

// MTLIndirectCommandBufferDescriptorFromID constructs a [MTLIndirectCommandBufferDescriptor] from an objc.ID.
//
// A configuration you create to customize an indirect command buffer.
func MTLIndirectCommandBufferDescriptorFromID(id objc.ID) MTLIndirectCommandBufferDescriptor {
	return MTLIndirectCommandBufferDescriptor{objectivec.Object{ID: id}}
}
// NOTE: MTLIndirectCommandBufferDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTLIndirectCommandBufferDescriptor] class.
//
// # Declaring command types to encode
//
//   - [IMTLIndirectCommandBufferDescriptor.CommandTypes]: The set of command types that you can encode into the indirect command buffer.
//   - [IMTLIndirectCommandBufferDescriptor.SetCommandTypes]
//
// # Declaring command inheritance
//
//   - [IMTLIndirectCommandBufferDescriptor.InheritBuffers]: A Boolean value that determines where commands in the indirect command buffer get their buffer arguments from when you execute them.
//   - [IMTLIndirectCommandBufferDescriptor.SetInheritBuffers]
//   - [IMTLIndirectCommandBufferDescriptor.InheritPipelineState]: A Boolean value that determines where commands in the indirect command buffer get their pipeline state from when you execute them.
//   - [IMTLIndirectCommandBufferDescriptor.SetInheritPipelineState]
//
// # Declaring the maximum number of argument buffers per command
//
//   - [IMTLIndirectCommandBufferDescriptor.MaxVertexBufferBindCount]: The maximum number of buffers that you can set per command for the vertex stage.
//   - [IMTLIndirectCommandBufferDescriptor.SetMaxVertexBufferBindCount]
//   - [IMTLIndirectCommandBufferDescriptor.MaxFragmentBufferBindCount]: The maximum number of buffers that you can set per command for the fragment stage.
//   - [IMTLIndirectCommandBufferDescriptor.SetMaxFragmentBufferBindCount]
//   - [IMTLIndirectCommandBufferDescriptor.MaxKernelBufferBindCount]: The maximum number of buffers that you can set per command for the compute kernel.
//   - [IMTLIndirectCommandBufferDescriptor.SetMaxKernelBufferBindCount]
//
// # Instance Properties
//
//   - [IMTLIndirectCommandBufferDescriptor.InheritCullMode]: Configures whether the indirect command buffer inherits the cull mode from the encoder.
//   - [IMTLIndirectCommandBufferDescriptor.SetInheritCullMode]
//   - [IMTLIndirectCommandBufferDescriptor.InheritDepthBias]: Configures whether the indirect command buffer inherits the depth bias from the encoder.
//   - [IMTLIndirectCommandBufferDescriptor.SetInheritDepthBias]
//   - [IMTLIndirectCommandBufferDescriptor.InheritDepthClipMode]: Configures whether the indirect command buffer inherits the depth clip mode from the encoder.
//   - [IMTLIndirectCommandBufferDescriptor.SetInheritDepthClipMode]
//   - [IMTLIndirectCommandBufferDescriptor.InheritDepthStencilState]: Configures whether the indirect command buffer inherits the depth stencil state from the encoder.
//   - [IMTLIndirectCommandBufferDescriptor.SetInheritDepthStencilState]
//   - [IMTLIndirectCommandBufferDescriptor.InheritFrontFacingWinding]: Configures whether the indirect command buffer inherits the front facing winding from the encoder.
//   - [IMTLIndirectCommandBufferDescriptor.SetInheritFrontFacingWinding]
//   - [IMTLIndirectCommandBufferDescriptor.InheritTriangleFillMode]: Configures whether the indirect command buffer inherits the triangle fill mode from the encoder.
//   - [IMTLIndirectCommandBufferDescriptor.SetInheritTriangleFillMode]
//   - [IMTLIndirectCommandBufferDescriptor.MaxKernelThreadgroupMemoryBindCount]
//   - [IMTLIndirectCommandBufferDescriptor.SetMaxKernelThreadgroupMemoryBindCount]
//   - [IMTLIndirectCommandBufferDescriptor.MaxMeshBufferBindCount]
//   - [IMTLIndirectCommandBufferDescriptor.SetMaxMeshBufferBindCount]
//   - [IMTLIndirectCommandBufferDescriptor.MaxObjectBufferBindCount]
//   - [IMTLIndirectCommandBufferDescriptor.SetMaxObjectBufferBindCount]
//   - [IMTLIndirectCommandBufferDescriptor.MaxObjectThreadgroupMemoryBindCount]
//   - [IMTLIndirectCommandBufferDescriptor.SetMaxObjectThreadgroupMemoryBindCount]
//   - [IMTLIndirectCommandBufferDescriptor.SupportColorAttachmentMapping]: Specifies if the indirect command buffer should support color attachment mapping.
//   - [IMTLIndirectCommandBufferDescriptor.SetSupportColorAttachmentMapping]
//   - [IMTLIndirectCommandBufferDescriptor.SupportDynamicAttributeStride]
//   - [IMTLIndirectCommandBufferDescriptor.SetSupportDynamicAttributeStride]
//   - [IMTLIndirectCommandBufferDescriptor.SupportRayTracing]
//   - [IMTLIndirectCommandBufferDescriptor.SetSupportRayTracing]
//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor
type IMTLIndirectCommandBufferDescriptor interface {
	objectivec.IObject

	// Topic: Declaring command types to encode

	// The set of command types that you can encode into the indirect command buffer.
	CommandTypes() MTLIndirectCommandType
	SetCommandTypes(value MTLIndirectCommandType)

	// Topic: Declaring command inheritance

	// A Boolean value that determines where commands in the indirect command buffer get their buffer arguments from when you execute them.
	InheritBuffers() bool
	SetInheritBuffers(value bool)
	// A Boolean value that determines where commands in the indirect command buffer get their pipeline state from when you execute them.
	InheritPipelineState() bool
	SetInheritPipelineState(value bool)

	// Topic: Declaring the maximum number of argument buffers per command

	// The maximum number of buffers that you can set per command for the vertex stage.
	MaxVertexBufferBindCount() uint
	SetMaxVertexBufferBindCount(value uint)
	// The maximum number of buffers that you can set per command for the fragment stage.
	MaxFragmentBufferBindCount() uint
	SetMaxFragmentBufferBindCount(value uint)
	// The maximum number of buffers that you can set per command for the compute kernel.
	MaxKernelBufferBindCount() uint
	SetMaxKernelBufferBindCount(value uint)

	// Topic: Instance Properties

	// Configures whether the indirect command buffer inherits the cull mode from the encoder.
	InheritCullMode() bool
	SetInheritCullMode(value bool)
	// Configures whether the indirect command buffer inherits the depth bias from the encoder.
	InheritDepthBias() bool
	SetInheritDepthBias(value bool)
	// Configures whether the indirect command buffer inherits the depth clip mode from the encoder.
	InheritDepthClipMode() bool
	SetInheritDepthClipMode(value bool)
	// Configures whether the indirect command buffer inherits the depth stencil state from the encoder.
	InheritDepthStencilState() bool
	SetInheritDepthStencilState(value bool)
	// Configures whether the indirect command buffer inherits the front facing winding from the encoder.
	InheritFrontFacingWinding() bool
	SetInheritFrontFacingWinding(value bool)
	// Configures whether the indirect command buffer inherits the triangle fill mode from the encoder.
	InheritTriangleFillMode() bool
	SetInheritTriangleFillMode(value bool)
	MaxKernelThreadgroupMemoryBindCount() uint
	SetMaxKernelThreadgroupMemoryBindCount(value uint)
	MaxMeshBufferBindCount() uint
	SetMaxMeshBufferBindCount(value uint)
	MaxObjectBufferBindCount() uint
	SetMaxObjectBufferBindCount(value uint)
	MaxObjectThreadgroupMemoryBindCount() uint
	SetMaxObjectThreadgroupMemoryBindCount(value uint)
	// Specifies if the indirect command buffer should support color attachment mapping.
	SupportColorAttachmentMapping() bool
	SetSupportColorAttachmentMapping(value bool)
	SupportDynamicAttributeStride() bool
	SetSupportDynamicAttributeStride(value bool)
	SupportRayTracing() bool
	SetSupportRayTracing(value bool)
}

// Init initializes the instance.
func (i MTLIndirectCommandBufferDescriptor) Init() MTLIndirectCommandBufferDescriptor {
	rv := objc.Send[MTLIndirectCommandBufferDescriptor](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i MTLIndirectCommandBufferDescriptor) Autorelease() MTLIndirectCommandBufferDescriptor {
	rv := objc.Send[MTLIndirectCommandBufferDescriptor](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLIndirectCommandBufferDescriptor creates a new MTLIndirectCommandBufferDescriptor instance.
func NewMTLIndirectCommandBufferDescriptor() MTLIndirectCommandBufferDescriptor {
	class := getMTLIndirectCommandBufferDescriptorClass()
	rv := objc.Send[MTLIndirectCommandBufferDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The set of command types that you can encode into the indirect command
// buffer.
//
// # Discussion
// 
// When you create the indirect command buffer, Metal allocates memory for
// each command it can hold. It needs to allocate enough memory to hold any
// command that you might later encode. To save space, specify only the
// command types you are going to encode in the indirect command buffer.
// 
// You can’t combine rendering and compute commands in the same indirect
// command buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/commandTypes
func (i MTLIndirectCommandBufferDescriptor) CommandTypes() MTLIndirectCommandType {
	rv := objc.Send[MTLIndirectCommandType](i.ID, objc.Sel("commandTypes"))
	return MTLIndirectCommandType(rv)
}
func (i MTLIndirectCommandBufferDescriptor) SetCommandTypes(value MTLIndirectCommandType) {
	objc.Send[struct{}](i.ID, objc.Sel("setCommandTypes:"), value)
}

// A Boolean value that determines where commands in the indirect command
// buffer get their buffer arguments from when you execute them.
//
// # Discussion
// 
// Always set this property explicitly.
// 
// If you set the value to [true], don’t set buffer arguments when you
// encode commands into the indirect command buffer. The commands use
// (inherit) the buffer arguments that you set on the parent encoder.
// 
// If you set the value to [false], set the buffer arguments when you encode
// the commands into the indirect command buffer. The commands ignore any
// buffer arguments set on the parent encoder.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/inheritBuffers
func (i MTLIndirectCommandBufferDescriptor) InheritBuffers() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("inheritBuffers"))
	return rv
}
func (i MTLIndirectCommandBufferDescriptor) SetInheritBuffers(value bool) {
	objc.Send[struct{}](i.ID, objc.Sel("setInheritBuffers:"), value)
}

// A Boolean value that determines where commands in the indirect command
// buffer get their pipeline state from when you execute them.
//
// # Discussion
// 
// The default value is [false]. If the value is [false], set the pipeline
// state object when you encode the commands into the indirect command buffer.
// The commands ignore any pipeline state object set on the parent encoder.
// 
// If you set the value to [true], don’t set a pipeline state object when
// you encode commands into the indirect command buffer. The commands use
// (inherit) the pipeline stage object that you set on the parent encoder.
// 
// This property doesn’t exist in iOS 12 and earlier, and tvOS 12 and
// earlier. If you create an indirect command buffer on those systems, it
// inherits the pipeline state, exactly as if the property existed, with a
// value of [true]. If you need your app to run on earlier versions of iOS,
// use an availability attribute to set the property conditionally:
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/inheritPipelineState
func (i MTLIndirectCommandBufferDescriptor) InheritPipelineState() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("inheritPipelineState"))
	return rv
}
func (i MTLIndirectCommandBufferDescriptor) SetInheritPipelineState(value bool) {
	objc.Send[struct{}](i.ID, objc.Sel("setInheritPipelineState:"), value)
}

// The maximum number of buffers that you can set per command for the vertex
// stage.
//
// # Discussion
// 
// Metal ignores this property if [InheritBuffers] is [true] or if you
// configured [CommandTypes] for compute commands. Metal needs to reserve
// enough memory in each command to store this many arguments. Use the
// smallest value that works for all commands you plan to encode into the
// indirect command buffer.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/maxVertexBufferBindCount
func (i MTLIndirectCommandBufferDescriptor) MaxVertexBufferBindCount() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("maxVertexBufferBindCount"))
	return rv
}
func (i MTLIndirectCommandBufferDescriptor) SetMaxVertexBufferBindCount(value uint) {
	objc.Send[struct{}](i.ID, objc.Sel("setMaxVertexBufferBindCount:"), value)
}

// The maximum number of buffers that you can set per command for the fragment
// stage.
//
// # Discussion
// 
// Metal ignores this property if [InheritBuffers] is [true] or if you
// configured [CommandTypes] for compute commands. Metal needs to reserve
// enough memory in each command to store this many arguments. Use the
// smallest value that works for all commands you plan to encode into the
// indirect command buffer.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/maxFragmentBufferBindCount
func (i MTLIndirectCommandBufferDescriptor) MaxFragmentBufferBindCount() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("maxFragmentBufferBindCount"))
	return rv
}
func (i MTLIndirectCommandBufferDescriptor) SetMaxFragmentBufferBindCount(value uint) {
	objc.Send[struct{}](i.ID, objc.Sel("setMaxFragmentBufferBindCount:"), value)
}

// The maximum number of buffers that you can set per command for the compute
// kernel.
//
// # Discussion
// 
// Metal ignores this property if [InheritBuffers] is [true] or if you
// configured [CommandTypes] for rendering commands. Metal needs to reserve
// enough memory in each command to store this many arguments. Use the
// smallest value that works for all commands you plan to encode into the
// indirect command buffer.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/maxKernelBufferBindCount
func (i MTLIndirectCommandBufferDescriptor) MaxKernelBufferBindCount() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("maxKernelBufferBindCount"))
	return rv
}
func (i MTLIndirectCommandBufferDescriptor) SetMaxKernelBufferBindCount(value uint) {
	objc.Send[struct{}](i.ID, objc.Sel("setMaxKernelBufferBindCount:"), value)
}

// Configures whether the indirect command buffer inherits the cull mode from
// the encoder.
//
// # Discussion
// 
// The property’s default value is [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/inheritCullMode
func (i MTLIndirectCommandBufferDescriptor) InheritCullMode() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("inheritCullMode"))
	return rv
}
func (i MTLIndirectCommandBufferDescriptor) SetInheritCullMode(value bool) {
	objc.Send[struct{}](i.ID, objc.Sel("setInheritCullMode:"), value)
}

// Configures whether the indirect command buffer inherits the depth bias from
// the encoder.
//
// # Discussion
// 
// The property’s default value is [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/inheritDepthBias
func (i MTLIndirectCommandBufferDescriptor) InheritDepthBias() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("inheritDepthBias"))
	return rv
}
func (i MTLIndirectCommandBufferDescriptor) SetInheritDepthBias(value bool) {
	objc.Send[struct{}](i.ID, objc.Sel("setInheritDepthBias:"), value)
}

// Configures whether the indirect command buffer inherits the depth clip mode
// from the encoder.
//
// # Discussion
// 
// The property’s default value is [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/inheritDepthClipMode
func (i MTLIndirectCommandBufferDescriptor) InheritDepthClipMode() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("inheritDepthClipMode"))
	return rv
}
func (i MTLIndirectCommandBufferDescriptor) SetInheritDepthClipMode(value bool) {
	objc.Send[struct{}](i.ID, objc.Sel("setInheritDepthClipMode:"), value)
}

// Configures whether the indirect command buffer inherits the depth stencil
// state from the encoder.
//
// # Discussion
// 
// The property’s default value is [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/inheritDepthStencilState
func (i MTLIndirectCommandBufferDescriptor) InheritDepthStencilState() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("inheritDepthStencilState"))
	return rv
}
func (i MTLIndirectCommandBufferDescriptor) SetInheritDepthStencilState(value bool) {
	objc.Send[struct{}](i.ID, objc.Sel("setInheritDepthStencilState:"), value)
}

// Configures whether the indirect command buffer inherits the front facing
// winding from the encoder.
//
// # Discussion
// 
// The property’s default value is [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/inheritFrontFacingWinding
func (i MTLIndirectCommandBufferDescriptor) InheritFrontFacingWinding() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("inheritFrontFacingWinding"))
	return rv
}
func (i MTLIndirectCommandBufferDescriptor) SetInheritFrontFacingWinding(value bool) {
	objc.Send[struct{}](i.ID, objc.Sel("setInheritFrontFacingWinding:"), value)
}

// Configures whether the indirect command buffer inherits the triangle fill
// mode from the encoder.
//
// # Discussion
// 
// The property’s default value is [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/inheritTriangleFillMode
func (i MTLIndirectCommandBufferDescriptor) InheritTriangleFillMode() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("inheritTriangleFillMode"))
	return rv
}
func (i MTLIndirectCommandBufferDescriptor) SetInheritTriangleFillMode(value bool) {
	objc.Send[struct{}](i.ID, objc.Sel("setInheritTriangleFillMode:"), value)
}

// See: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/maxKernelThreadgroupMemoryBindCount
func (i MTLIndirectCommandBufferDescriptor) MaxKernelThreadgroupMemoryBindCount() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("maxKernelThreadgroupMemoryBindCount"))
	return rv
}
func (i MTLIndirectCommandBufferDescriptor) SetMaxKernelThreadgroupMemoryBindCount(value uint) {
	objc.Send[struct{}](i.ID, objc.Sel("setMaxKernelThreadgroupMemoryBindCount:"), value)
}

// See: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/maxMeshBufferBindCount
func (i MTLIndirectCommandBufferDescriptor) MaxMeshBufferBindCount() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("maxMeshBufferBindCount"))
	return rv
}
func (i MTLIndirectCommandBufferDescriptor) SetMaxMeshBufferBindCount(value uint) {
	objc.Send[struct{}](i.ID, objc.Sel("setMaxMeshBufferBindCount:"), value)
}

// See: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/maxObjectBufferBindCount
func (i MTLIndirectCommandBufferDescriptor) MaxObjectBufferBindCount() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("maxObjectBufferBindCount"))
	return rv
}
func (i MTLIndirectCommandBufferDescriptor) SetMaxObjectBufferBindCount(value uint) {
	objc.Send[struct{}](i.ID, objc.Sel("setMaxObjectBufferBindCount:"), value)
}

// See: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/maxObjectThreadgroupMemoryBindCount
func (i MTLIndirectCommandBufferDescriptor) MaxObjectThreadgroupMemoryBindCount() uint {
	rv := objc.Send[uint](i.ID, objc.Sel("maxObjectThreadgroupMemoryBindCount"))
	return rv
}
func (i MTLIndirectCommandBufferDescriptor) SetMaxObjectThreadgroupMemoryBindCount(value uint) {
	objc.Send[struct{}](i.ID, objc.Sel("setMaxObjectThreadgroupMemoryBindCount:"), value)
}

// Specifies if the indirect command buffer should support color attachment
// mapping.
//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/supportColorAttachmentMapping
func (i MTLIndirectCommandBufferDescriptor) SupportColorAttachmentMapping() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("supportColorAttachmentMapping"))
	return rv
}
func (i MTLIndirectCommandBufferDescriptor) SetSupportColorAttachmentMapping(value bool) {
	objc.Send[struct{}](i.ID, objc.Sel("setSupportColorAttachmentMapping:"), value)
}

// See: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/supportDynamicAttributeStride
func (i MTLIndirectCommandBufferDescriptor) SupportDynamicAttributeStride() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("supportDynamicAttributeStride"))
	return rv
}
func (i MTLIndirectCommandBufferDescriptor) SetSupportDynamicAttributeStride(value bool) {
	objc.Send[struct{}](i.ID, objc.Sel("setSupportDynamicAttributeStride:"), value)
}

// See: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBufferDescriptor/supportRayTracing
func (i MTLIndirectCommandBufferDescriptor) SupportRayTracing() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("supportRayTracing"))
	return rv
}
func (i MTLIndirectCommandBufferDescriptor) SetSupportRayTracing(value bool) {
	objc.Send[struct{}](i.ID, objc.Sel("setSupportRayTracing:"), value)
}

