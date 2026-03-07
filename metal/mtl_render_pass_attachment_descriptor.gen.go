// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTLRenderPassAttachmentDescriptor] class.
var (
	_MTLRenderPassAttachmentDescriptorClass     MTLRenderPassAttachmentDescriptorClass
	_MTLRenderPassAttachmentDescriptorClassOnce sync.Once
)

func getMTLRenderPassAttachmentDescriptorClass() MTLRenderPassAttachmentDescriptorClass {
	_MTLRenderPassAttachmentDescriptorClassOnce.Do(func() {
		_MTLRenderPassAttachmentDescriptorClass = MTLRenderPassAttachmentDescriptorClass{class: objc.GetClass("MTLRenderPassAttachmentDescriptor")}
	})
	return _MTLRenderPassAttachmentDescriptorClass
}

// GetMTLRenderPassAttachmentDescriptorClass returns the class object for MTLRenderPassAttachmentDescriptor.
func GetMTLRenderPassAttachmentDescriptorClass() MTLRenderPassAttachmentDescriptorClass {
	return getMTLRenderPassAttachmentDescriptorClass()
}

type MTLRenderPassAttachmentDescriptorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTLRenderPassAttachmentDescriptorClass) Alloc() MTLRenderPassAttachmentDescriptor {
	rv := objc.Send[MTLRenderPassAttachmentDescriptor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}







// A render target that serves as the output destination for pixels generated
// by a render pass.
//
// # Overview
// 
// Use an [MTLRenderPassAttachmentDescriptor] instance to configure an
// individual render target of a framebuffer. Each
// [MTLRenderPassAttachmentDescriptor] instance specifies one texture that a
// graphics rendering pass can write into.
// 
// Typically, you don’t directly create [MTLRenderPassAttachmentDescriptor]
// instances. Instead, the [MTLRenderPassDescriptor] instance creates a
// default set of attachment instances. For each attachment that you intend to
// use as a render target, retrieve the [MTLRenderPassAttachmentDescriptor]
// instance from the render pass descriptor and configure its properties for
// use during this rendering pass.
// 
// You need to set the attachment’s [MTLRenderPassAttachmentDescriptor.Texture] property. The [MTLRenderPassAttachmentDescriptor.Level],
// [MTLRenderPassAttachmentDescriptor.Slice], and [MTLRenderPassAttachmentDescriptor.DepthPlane] properties specify the mipmap level, slice, and
// depth plane (for 3D textures) of the texture, respectively.
// 
// The [MTLRenderPassAttachmentDescriptor.LoadAction] and [MTLRenderPassAttachmentDescriptor.StoreAction] properties specify actions to perform at
// the start and end of a rendering pass for the attachment, respectively. For
// example, if you set the [MTLRenderPassAttachmentDescriptor.LoadAction] property of an attachment to
// [LoadActionClear], then the contents of the texture fill with a value for
// the type of attachment at the start of the rendering pass.
// 
// There are specific [MTLRenderPassAttachmentDescriptor] subclasses for
// color, depth, and stencil attachments. Each subclass provides additional
// properties to configure for that kind of attachment. The table below
// provides the list of subclasses.
// 
// [Table data omitted]
// 
// # Multisampling
// 
// To perform multisampled antialiased rendering, you use two textures. Attach
// to the [MTLRenderPassAttachmentDescriptor.Texture] property a [TextureType2DMultisample] texture, and a 2D or
// cube texture to the [MTLRenderPassAttachmentDescriptor.ResolveTexture] property. When a rendering command
// executes, it renders to the multisample texture. At the end of the render
// pass, the GPU resolves the contents of the multisample texture and writes
// the results into the resolve texture. The [MTLRenderPassAttachmentDescriptor.ResolveLevel], [MTLRenderPassAttachmentDescriptor.ResolveSlice],
// and [MTLRenderPassAttachmentDescriptor.ResolveDepthPlane] properties specify where the resolved image is
// written to. The attachment’s [MTLRenderPassAttachmentDescriptor.StoreAction] property determines what
// happens to the multisample texture after the GPU resolves its data.
//
// # Specifying the texture for the attachment
//
//   - [MTLRenderPassAttachmentDescriptor.Texture]: The texture object associated with this attachment.
//   - [MTLRenderPassAttachmentDescriptor.SetTexture]
//   - [MTLRenderPassAttachmentDescriptor.Level]: The mipmap level of the texture used for rendering to the attachment.
//   - [MTLRenderPassAttachmentDescriptor.SetLevel]
//   - [MTLRenderPassAttachmentDescriptor.Slice]: The slice of the texture used for rendering to the attachment.
//   - [MTLRenderPassAttachmentDescriptor.SetSlice]
//   - [MTLRenderPassAttachmentDescriptor.DepthPlane]: The depth plane of the texture used for rendering to the attachment.
//   - [MTLRenderPassAttachmentDescriptor.SetDepthPlane]
//
// # Specifying rendering pass actions
//
//   - [MTLRenderPassAttachmentDescriptor.LoadAction]: The action performed by this attachment at the start of a rendering pass for a render command encoder.
//   - [MTLRenderPassAttachmentDescriptor.SetLoadAction]
//   - [MTLRenderPassAttachmentDescriptor.StoreAction]: The action performed by this attachment at the end of a rendering pass for a render command encoder.
//   - [MTLRenderPassAttachmentDescriptor.SetStoreAction]
//   - [MTLRenderPassAttachmentDescriptor.StoreActionOptions]: The options that modify the store action performed by this attachment.
//   - [MTLRenderPassAttachmentDescriptor.SetStoreActionOptions]
//
// # Specifying the texture to resolve multisample data
//
//   - [MTLRenderPassAttachmentDescriptor.ResolveTexture]: The destination texture used when resolving multisampled texture data into single sample values.
//   - [MTLRenderPassAttachmentDescriptor.SetResolveTexture]
//   - [MTLRenderPassAttachmentDescriptor.ResolveLevel]: The mipmap level of the texture used for the multisample resolve action.
//   - [MTLRenderPassAttachmentDescriptor.SetResolveLevel]
//   - [MTLRenderPassAttachmentDescriptor.ResolveSlice]: The slice of the texture used for the multisample resolve action.
//   - [MTLRenderPassAttachmentDescriptor.SetResolveSlice]
//   - [MTLRenderPassAttachmentDescriptor.ResolveDepthPlane]: The depth plane of the texture used for the multisample resolve action.
//   - [MTLRenderPassAttachmentDescriptor.SetResolveDepthPlane]
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPassAttachmentDescriptor
type MTLRenderPassAttachmentDescriptor struct {
	objectivec.Object
}

// MTLRenderPassAttachmentDescriptorFromID constructs a [MTLRenderPassAttachmentDescriptor] from an objc.ID.
//
// A render target that serves as the output destination for pixels generated
// by a render pass.
func MTLRenderPassAttachmentDescriptorFromID(id objc.ID) MTLRenderPassAttachmentDescriptor {
	return MTLRenderPassAttachmentDescriptor{objectivec.Object{id}}
}
// NOTE: MTLRenderPassAttachmentDescriptor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [MTLRenderPassAttachmentDescriptor] class.
//
// # Specifying the texture for the attachment
//
//   - [IMTLRenderPassAttachmentDescriptor.Texture]: The texture object associated with this attachment.
//   - [IMTLRenderPassAttachmentDescriptor.SetTexture]
//   - [IMTLRenderPassAttachmentDescriptor.Level]: The mipmap level of the texture used for rendering to the attachment.
//   - [IMTLRenderPassAttachmentDescriptor.SetLevel]
//   - [IMTLRenderPassAttachmentDescriptor.Slice]: The slice of the texture used for rendering to the attachment.
//   - [IMTLRenderPassAttachmentDescriptor.SetSlice]
//   - [IMTLRenderPassAttachmentDescriptor.DepthPlane]: The depth plane of the texture used for rendering to the attachment.
//   - [IMTLRenderPassAttachmentDescriptor.SetDepthPlane]
//
// # Specifying rendering pass actions
//
//   - [IMTLRenderPassAttachmentDescriptor.LoadAction]: The action performed by this attachment at the start of a rendering pass for a render command encoder.
//   - [IMTLRenderPassAttachmentDescriptor.SetLoadAction]
//   - [IMTLRenderPassAttachmentDescriptor.StoreAction]: The action performed by this attachment at the end of a rendering pass for a render command encoder.
//   - [IMTLRenderPassAttachmentDescriptor.SetStoreAction]
//   - [IMTLRenderPassAttachmentDescriptor.StoreActionOptions]: The options that modify the store action performed by this attachment.
//   - [IMTLRenderPassAttachmentDescriptor.SetStoreActionOptions]
//
// # Specifying the texture to resolve multisample data
//
//   - [IMTLRenderPassAttachmentDescriptor.ResolveTexture]: The destination texture used when resolving multisampled texture data into single sample values.
//   - [IMTLRenderPassAttachmentDescriptor.SetResolveTexture]
//   - [IMTLRenderPassAttachmentDescriptor.ResolveLevel]: The mipmap level of the texture used for the multisample resolve action.
//   - [IMTLRenderPassAttachmentDescriptor.SetResolveLevel]
//   - [IMTLRenderPassAttachmentDescriptor.ResolveSlice]: The slice of the texture used for the multisample resolve action.
//   - [IMTLRenderPassAttachmentDescriptor.SetResolveSlice]
//   - [IMTLRenderPassAttachmentDescriptor.ResolveDepthPlane]: The depth plane of the texture used for the multisample resolve action.
//   - [IMTLRenderPassAttachmentDescriptor.SetResolveDepthPlane]
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPassAttachmentDescriptor
type IMTLRenderPassAttachmentDescriptor interface {
	objectivec.IObject

	// Topic: Specifying the texture for the attachment

	// The texture object associated with this attachment.
	Texture() MTLTexture
	SetTexture(value MTLTexture)
	// The mipmap level of the texture used for rendering to the attachment.
	Level() uint
	SetLevel(value uint)
	// The slice of the texture used for rendering to the attachment.
	Slice() uint
	SetSlice(value uint)
	// The depth plane of the texture used for rendering to the attachment.
	DepthPlane() uint
	SetDepthPlane(value uint)

	// Topic: Specifying rendering pass actions

	// The action performed by this attachment at the start of a rendering pass for a render command encoder.
	LoadAction() MTLLoadAction
	SetLoadAction(value MTLLoadAction)
	// The action performed by this attachment at the end of a rendering pass for a render command encoder.
	StoreAction() MTLStoreAction
	SetStoreAction(value MTLStoreAction)
	// The options that modify the store action performed by this attachment.
	StoreActionOptions() MTLStoreActionOptions
	SetStoreActionOptions(value MTLStoreActionOptions)

	// Topic: Specifying the texture to resolve multisample data

	// The destination texture used when resolving multisampled texture data into single sample values.
	ResolveTexture() MTLTexture
	SetResolveTexture(value MTLTexture)
	// The mipmap level of the texture used for the multisample resolve action.
	ResolveLevel() uint
	SetResolveLevel(value uint)
	// The slice of the texture used for the multisample resolve action.
	ResolveSlice() uint
	SetResolveSlice(value uint)
	// The depth plane of the texture used for the multisample resolve action.
	ResolveDepthPlane() uint
	SetResolveDepthPlane(value uint)
}





// Init initializes the instance.
func (r MTLRenderPassAttachmentDescriptor) Init() MTLRenderPassAttachmentDescriptor {
	rv := objc.Send[MTLRenderPassAttachmentDescriptor](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r MTLRenderPassAttachmentDescriptor) Autorelease() MTLRenderPassAttachmentDescriptor {
	rv := objc.Send[MTLRenderPassAttachmentDescriptor](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTLRenderPassAttachmentDescriptor creates a new MTLRenderPassAttachmentDescriptor instance.
func NewMTLRenderPassAttachmentDescriptor() MTLRenderPassAttachmentDescriptor {
	class := getMTLRenderPassAttachmentDescriptorClass()
	rv := objc.Send[MTLRenderPassAttachmentDescriptor](objc.ID(class.class), objc.Sel("new"))
	return rv
}





















// The texture object associated with this attachment.
//
// # Discussion
// 
// You need to set the attachment’s `texture` property, choosing an
// appropriate pixel format for the texture.
// 
// - To store color values in an attachment, use a texture with a
// color-renderable pixel format. - To store depth values, use a texture with
// a depth-renderable pixel format, such as [PixelFormatDepth32Float]. - To
// store stencil values, use a texture with a stencil-renderable pixel format,
// such as [PixelFormatStencil8].
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPassAttachmentDescriptor/texture
func (r MTLRenderPassAttachmentDescriptor) Texture() MTLTexture {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("texture"))
	return MTLTextureObjectFromID(rv)
}
func (r MTLRenderPassAttachmentDescriptor) SetTexture(value MTLTexture) {
	objc.Send[struct{}](r.ID, objc.Sel("setTexture:"), value)
}



// The mipmap level of the texture used for rendering to the attachment.
//
// # Discussion
// 
// The default value is `0`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPassAttachmentDescriptor/level
func (r MTLRenderPassAttachmentDescriptor) Level() uint {
	rv := objc.Send[uint](r.ID, objc.Sel("level"))
	return rv
}
func (r MTLRenderPassAttachmentDescriptor) SetLevel(value uint) {
	objc.Send[struct{}](r.ID, objc.Sel("setLevel:"), value)
}



// The slice of the texture used for rendering to the attachment.
//
// # Discussion
// 
// The default value is `0`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPassAttachmentDescriptor/slice
func (r MTLRenderPassAttachmentDescriptor) Slice() uint {
	rv := objc.Send[uint](r.ID, objc.Sel("slice"))
	return rv
}
func (r MTLRenderPassAttachmentDescriptor) SetSlice(value uint) {
	objc.Send[struct{}](r.ID, objc.Sel("setSlice:"), value)
}



// The depth plane of the texture used for rendering to the attachment.
//
// # Discussion
// 
// If the texture isn’t a 3D texture, then Metal ignores this property.
// 
// The default value is `0`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPassAttachmentDescriptor/depthPlane
func (r MTLRenderPassAttachmentDescriptor) DepthPlane() uint {
	rv := objc.Send[uint](r.ID, objc.Sel("depthPlane"))
	return rv
}
func (r MTLRenderPassAttachmentDescriptor) SetDepthPlane(value uint) {
	objc.Send[struct{}](r.ID, objc.Sel("setDepthPlane:"), value)
}



// The action performed by this attachment at the start of a rendering pass
// for a render command encoder.
//
// # Discussion
// 
// If your app renders all pixels of the render target for a given frame, use
// the [LoadActionDontCare] action, which allows the GPU to avoid loading the
// existing contents of the texture. Otherwise, use the [LoadActionClear]
// action to clear the previous contents of the render target or the
// [LoadActionLoad] action to preserve them. The [LoadActionClear] action also
// avoids the cost of loading the existing texture contents, but it still
// incurs the cost of filling the destination with a clear color.
// 
// For color render targets, the default value is [LoadActionDontCare]. For
// depth or stencil render targets, the default value is [LoadActionClear].
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPassAttachmentDescriptor/loadAction
func (r MTLRenderPassAttachmentDescriptor) LoadAction() MTLLoadAction {
	rv := objc.Send[MTLLoadAction](r.ID, objc.Sel("loadAction"))
	return MTLLoadAction(rv)
}
func (r MTLRenderPassAttachmentDescriptor) SetLoadAction(value MTLLoadAction) {
	objc.Send[struct{}](r.ID, objc.Sel("setLoadAction:"), value)
}



// The action performed by this attachment at the end of a rendering pass for
// a render command encoder.
//
// # Discussion
// 
// If your app doesn’t need the data in the texture after completing the
// rendering pass, use the [StoreActionDontCare] action. Otherwise, use the
// [StoreActionStore] action if the texture is directly stored or the
// [StoreActionMultisampleResolve] action if the texture is a multisampled
// texture. In some feature sets, you can use the
// [StoreActionStoreAndMultisampleResolve] action to store and resolve the
// texture in a single rendering pass. For more information, see:
// 
// - [Metal feature set tables (PDF)]
// - [Metal feature set tables (Numbers)]
// 
// When the store action is either [StoreActionMultisampleResolve] or
// [StoreActionStoreAndMultisampleResolve], the [ResolveTexture] property
// needs to be set to the texture to use as the target for the resolve action.
// Use the [ResolveLevel], [ResolveSlice], and [ResolveDepthPlane] properties
// to specify the mipmap level, cube slice, and depth plane of the resolve
// texture, respectively.
// 
// For color render targets, the default value is [StoreActionStore]. For
// depth or stencil render targets, the default value is
// [StoreActionDontCare].
//
// [Metal feature set tables (Numbers)]: https://developer.apple.com/metal/metal-feature-set-tables.zip
// [Metal feature set tables (PDF)]: https://developer.apple.com/metal/Metal-Feature-Set-Tables.pdf
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPassAttachmentDescriptor/storeAction
func (r MTLRenderPassAttachmentDescriptor) StoreAction() MTLStoreAction {
	rv := objc.Send[MTLStoreAction](r.ID, objc.Sel("storeAction"))
	return MTLStoreAction(rv)
}
func (r MTLRenderPassAttachmentDescriptor) SetStoreAction(value MTLStoreAction) {
	objc.Send[struct{}](r.ID, objc.Sel("setStoreAction:"), value)
}



// The options that modify the store action performed by this attachment.
//
// # Discussion
// 
// This property specifies additional behavior for the store action specified
// by the [StoreAction] property.
// 
// The default value is [StoreActionOptionNone].
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPassAttachmentDescriptor/storeActionOptions
func (r MTLRenderPassAttachmentDescriptor) StoreActionOptions() MTLStoreActionOptions {
	rv := objc.Send[MTLStoreActionOptions](r.ID, objc.Sel("storeActionOptions"))
	return MTLStoreActionOptions(rv)
}
func (r MTLRenderPassAttachmentDescriptor) SetStoreActionOptions(value MTLStoreActionOptions) {
	objc.Send[struct{}](r.ID, objc.Sel("setStoreActionOptions:"), value)
}



// The destination texture used when resolving multisampled texture data into
// single sample values.
//
// # Discussion
// 
// If the [StoreAction] value is set to [StoreActionMultisampleResolve] or
// [StoreActionStoreAndMultisampleResolve], then the [ResolveTexture] value
// needs to point to a valid texture. Otherwise, Metal ignores this property.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPassAttachmentDescriptor/resolveTexture
func (r MTLRenderPassAttachmentDescriptor) ResolveTexture() MTLTexture {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("resolveTexture"))
	return MTLTextureObjectFromID(rv)
}
func (r MTLRenderPassAttachmentDescriptor) SetResolveTexture(value MTLTexture) {
	objc.Send[struct{}](r.ID, objc.Sel("setResolveTexture:"), value)
}



// The mipmap level of the texture used for the multisample resolve action.
//
// # Discussion
// 
// If the value of [StoreAction] is set to [StoreActionMultisampleResolve] or
// [StoreActionStoreAndMultisampleResolve], set this property to point to a
// mipmap in the resolve texture.
// 
// The default value is `0`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPassAttachmentDescriptor/resolveLevel
func (r MTLRenderPassAttachmentDescriptor) ResolveLevel() uint {
	rv := objc.Send[uint](r.ID, objc.Sel("resolveLevel"))
	return rv
}
func (r MTLRenderPassAttachmentDescriptor) SetResolveLevel(value uint) {
	objc.Send[struct{}](r.ID, objc.Sel("setResolveLevel:"), value)
}



// The slice of the texture used for the multisample resolve action.
//
// # Discussion
// 
// If the value of [StoreAction] is set to [StoreActionMultisampleResolve] or
// [StoreActionStoreAndMultisampleResolve], set this property to point to a
// slice in the resolve texture.
// 
// The default value is `0`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPassAttachmentDescriptor/resolveSlice
func (r MTLRenderPassAttachmentDescriptor) ResolveSlice() uint {
	rv := objc.Send[uint](r.ID, objc.Sel("resolveSlice"))
	return rv
}
func (r MTLRenderPassAttachmentDescriptor) SetResolveSlice(value uint) {
	objc.Send[struct{}](r.ID, objc.Sel("setResolveSlice:"), value)
}



// The depth plane of the texture used for the multisample resolve action.
//
// # Discussion
// 
// If the value of [StoreAction] is set to [StoreActionMultisampleResolve] or
// [StoreActionStoreAndMultisampleResolve], set this property to point to a
// depth plane in the resolve texture.
// 
// The default value is `0`.
//
// See: https://developer.apple.com/documentation/Metal/MTLRenderPassAttachmentDescriptor/resolveDepthPlane
func (r MTLRenderPassAttachmentDescriptor) ResolveDepthPlane() uint {
	rv := objc.Send[uint](r.ID, objc.Sel("resolveDepthPlane"))
	return rv
}
func (r MTLRenderPassAttachmentDescriptor) SetResolveDepthPlane(value uint) {
	objc.Send[struct{}](r.ID, objc.Sel("setResolveDepthPlane:"), value)
}
























