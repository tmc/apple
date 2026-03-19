// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CAMetalLayer] class.
var (
	_CAMetalLayerClass     CAMetalLayerClass
	_CAMetalLayerClassOnce sync.Once
)

func getCAMetalLayerClass() CAMetalLayerClass {
	_CAMetalLayerClassOnce.Do(func() {
		_CAMetalLayerClass = CAMetalLayerClass{class: objc.GetClass("CAMetalLayer")}
	})
	return _CAMetalLayerClass
}

// GetCAMetalLayerClass returns the class object for CAMetalLayer.
func GetCAMetalLayerClass() CAMetalLayerClass {
	return getCAMetalLayerClass()
}

type CAMetalLayerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (cc CAMetalLayerClass) Alloc() CAMetalLayer {
	rv := objc.Send[CAMetalLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A Core Animation layer that Metal can render into, typically displayed
// onscreen.
//
// # Overview
// 
// Use a [CAMetalLayer] when you want to use Metal to render a layer’s
// contents; for example, to render into a view. Consider using [MTKView]
// instead, because this class automatically wraps a [CAMetalLayer] object and
// provides a higher-level abstraction.
// 
// If you’re using UIKit, to create a view that uses a [CAMetalLayer],
// create a subclass of [UIView] and override its [CAMetalLayer.LayerClass] class method to
// return a [CAMetalLayer]:
// 
// If you’re using AppKit, configure an [NSView] object to use a backing
// layer and assign a [CAMetalLayer] object to the view:
// 
// Adjust the layer’s properties to configure its underlying pixel format
// and other display behaviors.
// 
// # Rendering the Layer’s Contents
// 
// A [CAMetalLayer] creates a pool of Metal drawable objects
// ([CAMetalDrawable]). At any given time, one of these drawable objects
// contains the contents of the layer. To change the layer’s contents, ask
// the layer for a drawable object, render into it, and then update the
// layer’s contents to point to the new drawable.
// 
// Call the layer’s [CAMetalLayer.NextDrawable] method to obtain a drawable object. Get
// the drawable object’s texture and create a render pass that renders to
// that texture, as shown in the code below:
// 
// To change the layer’s contents to the new drawable, call the
// [present(_:)] method (or one of its variants) on the command buffer
// containing the encoded render pass, passing in the drawable object to
// present.
// 
// # Keeping References to Drawables
// 
// The layer reuses a drawable only if it isn’t onscreen and there are no
// strong references to it. Further, if a drawable isn’t available when you
// call [CAMetalLayer.NextDrawable], the system waits for one to become available. To avoid
// stalls in your app, request a new drawable only when you need it, and
// release any references to it as quickly as possible after you’re done
// with it.
// 
// For example, before retrieving a new drawable, you might perform other work
// on the CPU or submit commands to the GPU that don’t require the drawable.
// Then, obtain the drawable and encode a command buffer to render into it, as
// described above. After you commit this command buffer, release all strong
// references to the drawable. If you don’t release drawables correctly, the
// layer runs out of drawables, and future calls to [CAMetalLayer.NextDrawable] return
// `nil`.
// 
// # Releasing the Drawable
// 
// Don’t release the drawable explicitly; instead, embed your render loop
// within an autorelease pool block:
// 
// This block releases drawables promptly and avoids possible deadlock
// situations with multiple drawables. Release drawables as soon as possible
// after committing your onscreen render pass.
//
// [MTKView]: https://developer.apple.com/documentation/MetalKit/MTKView
// [NSView]: https://developer.apple.com/documentation/AppKit/NSView
// [UIView]: https://developer.apple.com/library/archive/releasenotes/iPhone/RN-iPhoneSDK/index.html#//apple_ref/doc/uid/TP40007428-CH1-SW18
// [present(_:)]: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/present(_:)
//
// # Configuring the Metal Device
//
//   - [CAMetalLayer.Device]: The Metal device responsible for the layer’s drawable resources.
//   - [CAMetalLayer.SetDevice]
//   - [CAMetalLayer.PreferredDevice]: The device object that the system recommends using for this layer.
//
// # Configuring the Layer’s Drawable Objects
//
//   - [CAMetalLayer.PixelFormat]: The pixel format of the layer’s textures.
//   - [CAMetalLayer.SetPixelFormat]
//   - [CAMetalLayer.Colorspace]: The color space of the rendered content.
//   - [CAMetalLayer.SetColorspace]
//   - [CAMetalLayer.FramebufferOnly]: A Boolean value that determines whether the layer’s textures are used only for rendering.
//   - [CAMetalLayer.SetFramebufferOnly]
//   - [CAMetalLayer.DrawableSize]: The size, in pixels, of textures for rendering layer content.
//   - [CAMetalLayer.SetDrawableSize]
//
// # Configuring Presentation Behavior
//
//   - [CAMetalLayer.PresentsWithTransaction]: A Boolean value that determines whether the layer presents its content using a Core Animation transaction.
//   - [CAMetalLayer.SetPresentsWithTransaction]
//   - [CAMetalLayer.DisplaySyncEnabled]: A Boolean value that determines whether the layer synchronizes its updates to the display’s refresh rate.
//   - [CAMetalLayer.SetDisplaySyncEnabled]
//
// # Configuring Extended Dynamic Range Behavior
//
//   - [CAMetalLayer.EDRMetadata]: Metadata describing the tone mapping to apply to the extended dynamic range (EDR) values in the layer.
//   - [CAMetalLayer.SetEDRMetadata]
//
// # Obtaining a Metal Drawable
//
//   - [CAMetalLayer.NextDrawable]: Waits until a Metal drawable is available, and then returns it.
//   - [CAMetalLayer.MaximumDrawableCount]: The number of Metal drawables in the resource pool managed by Core Animation.
//   - [CAMetalLayer.SetMaximumDrawableCount]
//   - [CAMetalLayer.AllowsNextDrawableTimeout]: A Boolean value that determines whether requests for a new buffer expire if the system can’t satisfy them.
//   - [CAMetalLayer.SetAllowsNextDrawableTimeout]
//
// # Configuring the Metal Performance HUD
//
//   - [CAMetalLayer.DeveloperHUDProperties]: The properties of the Metal performance heads-up display.
//   - [CAMetalLayer.SetDeveloperHUDProperties]
//
// # Instance Properties
//
//   - [CAMetalLayer.ResidencySet]
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer
type CAMetalLayer struct {
	CALayer
}

// CAMetalLayerFromID constructs a [CAMetalLayer] from an objc.ID.
//
// A Core Animation layer that Metal can render into, typically displayed
// onscreen.
func CAMetalLayerFromID(id objc.ID) CAMetalLayer {
	return CAMetalLayer{CALayer: CALayerFromID(id)}
}
// NOTE: CAMetalLayer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CAMetalLayer] class.
//
// # Configuring the Metal Device
//
//   - [ICAMetalLayer.Device]: The Metal device responsible for the layer’s drawable resources.
//   - [ICAMetalLayer.SetDevice]
//   - [ICAMetalLayer.PreferredDevice]: The device object that the system recommends using for this layer.
//
// # Configuring the Layer’s Drawable Objects
//
//   - [ICAMetalLayer.PixelFormat]: The pixel format of the layer’s textures.
//   - [ICAMetalLayer.SetPixelFormat]
//   - [ICAMetalLayer.Colorspace]: The color space of the rendered content.
//   - [ICAMetalLayer.SetColorspace]
//   - [ICAMetalLayer.FramebufferOnly]: A Boolean value that determines whether the layer’s textures are used only for rendering.
//   - [ICAMetalLayer.SetFramebufferOnly]
//   - [ICAMetalLayer.DrawableSize]: The size, in pixels, of textures for rendering layer content.
//   - [ICAMetalLayer.SetDrawableSize]
//
// # Configuring Presentation Behavior
//
//   - [ICAMetalLayer.PresentsWithTransaction]: A Boolean value that determines whether the layer presents its content using a Core Animation transaction.
//   - [ICAMetalLayer.SetPresentsWithTransaction]
//   - [ICAMetalLayer.DisplaySyncEnabled]: A Boolean value that determines whether the layer synchronizes its updates to the display’s refresh rate.
//   - [ICAMetalLayer.SetDisplaySyncEnabled]
//
// # Configuring Extended Dynamic Range Behavior
//
//   - [ICAMetalLayer.EDRMetadata]: Metadata describing the tone mapping to apply to the extended dynamic range (EDR) values in the layer.
//   - [ICAMetalLayer.SetEDRMetadata]
//
// # Obtaining a Metal Drawable
//
//   - [ICAMetalLayer.NextDrawable]: Waits until a Metal drawable is available, and then returns it.
//   - [ICAMetalLayer.MaximumDrawableCount]: The number of Metal drawables in the resource pool managed by Core Animation.
//   - [ICAMetalLayer.SetMaximumDrawableCount]
//   - [ICAMetalLayer.AllowsNextDrawableTimeout]: A Boolean value that determines whether requests for a new buffer expire if the system can’t satisfy them.
//   - [ICAMetalLayer.SetAllowsNextDrawableTimeout]
//
// # Configuring the Metal Performance HUD
//
//   - [ICAMetalLayer.DeveloperHUDProperties]: The properties of the Metal performance heads-up display.
//   - [ICAMetalLayer.SetDeveloperHUDProperties]
//
// # Instance Properties
//
//   - [ICAMetalLayer.ResidencySet]
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer
type ICAMetalLayer interface {
	ICALayer

	// Topic: Configuring the Metal Device

	// The Metal device responsible for the layer’s drawable resources.
	Device() objectivec.IObject
	SetDevice(value objectivec.IObject)
	// The device object that the system recommends using for this layer.
	PreferredDevice() objectivec.IObject

	// Topic: Configuring the Layer’s Drawable Objects

	// The pixel format of the layer’s textures.
	PixelFormat() objectivec.IObject
	SetPixelFormat(value objectivec.IObject)
	// The color space of the rendered content.
	Colorspace() coregraphics.CGColorSpaceRef
	SetColorspace(value coregraphics.CGColorSpaceRef)
	// A Boolean value that determines whether the layer’s textures are used only for rendering.
	FramebufferOnly() bool
	SetFramebufferOnly(value bool)
	// The size, in pixels, of textures for rendering layer content.
	DrawableSize() corefoundation.CGSize
	SetDrawableSize(value corefoundation.CGSize)

	// Topic: Configuring Presentation Behavior

	// A Boolean value that determines whether the layer presents its content using a Core Animation transaction.
	PresentsWithTransaction() bool
	SetPresentsWithTransaction(value bool)
	// A Boolean value that determines whether the layer synchronizes its updates to the display’s refresh rate.
	DisplaySyncEnabled() bool
	SetDisplaySyncEnabled(value bool)

	// Topic: Configuring Extended Dynamic Range Behavior

	// Metadata describing the tone mapping to apply to the extended dynamic range (EDR) values in the layer.
	EDRMetadata() ICAEDRMetadata
	SetEDRMetadata(value ICAEDRMetadata)

	// Topic: Obtaining a Metal Drawable

	// Waits until a Metal drawable is available, and then returns it.
	NextDrawable() CAMetalDrawable
	// The number of Metal drawables in the resource pool managed by Core Animation.
	MaximumDrawableCount() uint
	SetMaximumDrawableCount(value uint)
	// A Boolean value that determines whether requests for a new buffer expire if the system can’t satisfy them.
	AllowsNextDrawableTimeout() bool
	SetAllowsNextDrawableTimeout(value bool)

	// Topic: Configuring the Metal Performance HUD

	// The properties of the Metal performance heads-up display.
	DeveloperHUDProperties() foundation.INSDictionary
	SetDeveloperHUDProperties(value foundation.INSDictionary)

	// Topic: Instance Properties

	ResidencySet() objectivec.IObject

	// A positive integer that identifies the drawable.
	DrawableID() int
	SetDrawableID(value int)
	// The host time, in seconds, when the drawable was displayed onscreen.
	PresentedTime() float64
	SetPresentedTime(value float64)
}

// Init initializes the instance.
func (m CAMetalLayer) Init() CAMetalLayer {
	rv := objc.Send[CAMetalLayer](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m CAMetalLayer) Autorelease() CAMetalLayer {
	rv := objc.Send[CAMetalLayer](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewCAMetalLayer creates a new CAMetalLayer instance.
func NewCAMetalLayer() CAMetalLayer {
	class := getCAMetalLayerClass()
	rv := objc.Send[CAMetalLayer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Override to copy or initialize custom fields of the specified layer.
//
// layer: The layer from which custom fields should be copied.
//
// # Return Value
// 
// A layer instance with any custom instance variables copied from `layer`.
//
// # Discussion
// 
// This initializer is used to create shadow copies of layers, for example,
// for the [PresentationLayer] method. Using this method in any other
// situation will produce undefined behavior. For example, do not use this
// method to initialize a new layer with an existing layer’s content.
// 
// If you are implementing a custom layer subclass, you can override this
// method and use it to copy the values of instance variables into the new
// object. Subclasses should always invoke the superclass implementation.
// 
// This method is the designated initializer for layer objects in the
// presentation layer.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/init(layer:)
func NewMetalLayerWithLayer(layer objectivec.IObject) CAMetalLayer {
	instance := getCAMetalLayerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLayer:"), layer)
	return CAMetalLayerFromID(rv)
}

// Waits until a Metal drawable is available, and then returns it.
//
// # Return Value
// 
// A Metal drawable. Use the drawable’s [Texture] property to configure a
// [MTLRenderPipelineColorAttachmentDescriptor] object for rendering to the
// layer.
//
// [MTLRenderPipelineColorAttachmentDescriptor]: https://developer.apple.com/documentation/Metal/MTLRenderPipelineColorAttachmentDescriptor
//
// # Discussion
// 
// A [CAMetalLayer] object maintains an internal pool of textures for
// displaying layer content, each wrapped in a [CAMetalDrawable] object. Use
// this method to retrieve the next available drawable from the pool. If all
// drawables are in use, the layer waits up to one second for one to become
// available, after which it returns `nil`. The [AllowsNextDrawableTimeout]
// property affects this behavior.
// 
// This method returns `nil` if the layer’s [PixelFormat] or other
// properties are invalid.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/nextDrawable()
func (m CAMetalLayer) NextDrawable() CAMetalDrawable {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("nextDrawable"))
	return CAMetalDrawableObjectFromID(rv)
}

// The Metal device responsible for the layer’s drawable resources.
//
// # Discussion
// 
// This property determines which device object Metal uses to create its
// [MTLTexture] objects. When you retrieve a drawable object and its
// associated texture, you must render to the texture using the same device
// object.
// 
// The default value is `nil`—you must set the device for a layer before
// rendering.
//
// [MTLTexture]: https://developer.apple.com/documentation/Metal/MTLTexture
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/device
func (m CAMetalLayer) Device() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("device"))
	return objectivec.Object{ID: rv}
}
func (m CAMetalLayer) SetDevice(value objectivec.IObject) {
	objc.Send[struct{}](m.ID, objc.Sel("setDevice:"), value)
}
// The device object that the system recommends using for this layer.
//
// # Discussion
// 
// On systems with a single GPU, this method returns the default device
// object; see [MTLCreateSystemDefaultDevice()]. On systems with more than one
// GPU, this method returns the [MTLDevice] that was last used to composite
// and present the [CAMetalLayer]. This device object usually corresponds to
// the GPU associated with the screen that’s displaying the layer. If you
// set the layer’s [Device] property to this device object, you reduce the
// number of cross-GPU texture copies that Core Animation must perform to
// present the layer’s contents onscreen.
//
// [MTLCreateSystemDefaultDevice()]: https://developer.apple.com/documentation/Metal/MTLCreateSystemDefaultDevice()
// [MTLDevice]: https://developer.apple.com/documentation/Metal/MTLDevice
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/preferredDevice
func (m CAMetalLayer) PreferredDevice() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("preferredDevice"))
	return objectivec.Object{ID: rv}
}
// The pixel format of the layer’s textures.
//
// # Discussion
// 
// The default value is [MTLPixelFormat.bgra8Unorm].
// 
// You must use one of the following formats:
// 
// - [MTLPixelFormat.bgra8Unorm] - [MTLPixelFormat.bgra8Unorm_srgb] -
// [MTLPixelFormat.rgba16Float] - [MTLPixelFormat.rgb10a2Unorm] -
// [MTLPixelFormat.bgr10a2Unorm] - [MTLPixelFormat.bgra10_xr] -
// [MTLPixelFormat.bgra10_xr_srgb] - [MTLPixelFormat.bgr10_xr] -
// [MTLPixelFormat.bgr10_xr_srgb]
//
// [MTLPixelFormat.bgr10_xr]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/bgr10_xr
// [MTLPixelFormat.bgr10_xr_srgb]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/bgr10_xr_srgb
// [MTLPixelFormat.bgr10a2Unorm]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/bgr10a2Unorm
// [MTLPixelFormat.bgra10_xr]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/bgra10_xr
// [MTLPixelFormat.bgra10_xr_srgb]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/bgra10_xr_srgb
// [MTLPixelFormat.bgra8Unorm]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/bgra8Unorm
// [MTLPixelFormat.bgra8Unorm_srgb]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/bgra8Unorm_srgb
// [MTLPixelFormat.rgb10a2Unorm]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/rgb10a2Unorm
// [MTLPixelFormat.rgba16Float]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/rgba16Float
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/pixelFormat
func (m CAMetalLayer) PixelFormat() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("pixelFormat"))
	return objectivec.Object{ID: rv}
}
func (m CAMetalLayer) SetPixelFormat(value objectivec.IObject) {
	objc.Send[struct{}](m.ID, objc.Sel("setPixelFormat:"), value)
}
// The color space of the rendered content.
//
// # Discussion
// 
// Set this value to specify a color space for the contents of the layer. When
// a color space is present, Core Animation performs any necessary color space
// transformations when compositing this content.
// 
// The default value is `nil`, indicating that the rendered content isn’t
// color-matched.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/colorspace
func (m CAMetalLayer) Colorspace() coregraphics.CGColorSpaceRef {
	rv := objc.Send[coregraphics.CGColorSpaceRef](m.ID, objc.Sel("colorspace"))
	return coregraphics.CGColorSpaceRef(rv)
}
func (m CAMetalLayer) SetColorspace(value coregraphics.CGColorSpaceRef) {
	objc.Send[struct{}](m.ID, objc.Sel("setColorspace:"), value)
}
// A Boolean value that determines whether the layer’s textures are used
// only for rendering.
//
// # Discussion
// 
// If the value is [true] (the default), the [CAMetalLayer] class allocates
// its [MTLTexture] objects with only the [renderTarget] usage flag. Core
// Animation can then optimize the texture for display purposes. However, you
// may not sample, read from, or write to those textures. To support sampling
// and pixel read/write operations (at a cost to performance), set this value
// to [false].
//
// [MTLTexture]: https://developer.apple.com/documentation/Metal/MTLTexture
// [false]: https://developer.apple.com/documentation/Swift/false
// [renderTarget]: https://developer.apple.com/documentation/Metal/MTLTextureUsage/renderTarget
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/framebufferOnly
func (m CAMetalLayer) FramebufferOnly() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("framebufferOnly"))
	return rv
}
func (m CAMetalLayer) SetFramebufferOnly(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setFramebufferOnly:"), value)
}
// The size, in pixels, of textures for rendering layer content.
//
// # Discussion
// 
// By default, a layer creates textures sized to match its content—that is,
// this property’s value is the layer’s [Bounds] size multiplied by its
// [ContentsScale] factor.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/drawableSize
func (m CAMetalLayer) DrawableSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](m.ID, objc.Sel("drawableSize"))
	return corefoundation.CGSize(rv)
}
func (m CAMetalLayer) SetDrawableSize(value corefoundation.CGSize) {
	objc.Send[struct{}](m.ID, objc.Sel("setDrawableSize:"), value)
}
// A Boolean value that determines whether the layer presents its content
// using a Core Animation transaction.
//
// # Discussion
// 
// By default, this value is [false]; [CAMetalLayer] displays the output of a
// rendering pass to the display as quickly as possible and asynchronously to
// any Core Animation transactions. Core Animation doesn’t guarantee that
// the Metal content arrives in the same frame as other Core Animation
// content. This behavior could be an issue if, for example, your app draws
// [UIKit] content over the top of your [CAMetalLayer].
// 
// Setting this value to [true] makes the layer draw its contents
// synchronously, using whichever Core Animation transaction is current at the
// time you call the drawable’s [present()] method. To ensure that a
// transaction is available when you schedule the drawable to be presented,
// first commit the command buffer containing your Metal rendering commands.
// Then, call its [waitUntilScheduled()] method to synchronously wait until
// the command queue schedules the command buffer to execute on the GPU.
// Finally, call the drawable’s [present()] method.
//
// [UIKit]: https://developer.apple.com/library/archive/releasenotes/General/RN-iOSSDK-6_1/index.html#//apple_ref/doc/uid/TP40012869-CH1-SW17
// [false]: https://developer.apple.com/documentation/Swift/false
// [present()]: https://developer.apple.com/documentation/Metal/MTLDrawable/present()
// [true]: https://developer.apple.com/documentation/Swift/true
// [waitUntilScheduled()]: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/waitUntilScheduled()
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/presentsWithTransaction
func (m CAMetalLayer) PresentsWithTransaction() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("presentsWithTransaction"))
	return rv
}
func (m CAMetalLayer) SetPresentsWithTransaction(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setPresentsWithTransaction:"), value)
}
// A Boolean value that determines whether the layer synchronizes its updates
// to the display’s refresh rate.
//
// # Discussion
// 
// Set this value to [true] to synchronize the presentation of the layer’s
// contents with the display’s refresh, also known as or . If [false], the
// layer presents new content more quickly, but possibly with brief visual
// artifacts ().
// 
// The default value is [true].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/displaySyncEnabled
func (m CAMetalLayer) DisplaySyncEnabled() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("displaySyncEnabled"))
	return rv
}
func (m CAMetalLayer) SetDisplaySyncEnabled(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setDisplaySyncEnabled:"), value)
}
// Metadata describing the tone mapping to apply to the extended dynamic range
// (EDR) values in the layer.
//
// # Discussion
// 
// You must set this property before calling [NextDrawable].
// 
// The default value is `nil`, which means that the system doesn’t perform
// any tone mapping of data prior to passing it on to the display. Values
// above the maximum ([maximumExtendedDynamicRangeColorComponentValue]) may be
// clipped.
// 
// If non-`nil`, the system uses the metadata provided to tone map values to
// the display, based on the display’s current characteristics. You must
// also set [PixelFormat] to a pixel format that supports pixel values greater
// than `1.0` (such as [MTLPixelFormat.rgba16Float]) and [colorspace] to a
// color space that supports a linear transfer function.
// 
// The tone mapping process requires significant amounts of memory and GPU
// processing.
//
// [MTLPixelFormat.rgba16Float]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/rgba16Float
// [colorspace]: https://developer.apple.com/documentation/QuartzCore/CAOpenGLLayer/colorspace
// [maximumExtendedDynamicRangeColorComponentValue]: https://developer.apple.com/documentation/AppKit/NSScreen/maximumExtendedDynamicRangeColorComponentValue
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/edrMetadata
func (m CAMetalLayer) EDRMetadata() ICAEDRMetadata {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("EDRMetadata"))
	return CAEDRMetadataFromID(objc.ID(rv))
}
func (m CAMetalLayer) SetEDRMetadata(value ICAEDRMetadata) {
	objc.Send[struct{}](m.ID, objc.Sel("setEDRMetadata:"), value)
}
// The number of Metal drawables in the resource pool managed by Core
// Animation.
//
// # Discussion
// 
// You can set this value to `2` or `3` only; if you pass a different value,
// Core Animation ignores the value and throws an exception.
// 
// The default value is `3`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/maximumDrawableCount
func (m CAMetalLayer) MaximumDrawableCount() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("maximumDrawableCount"))
	return rv
}
func (m CAMetalLayer) SetMaximumDrawableCount(value uint) {
	objc.Send[struct{}](m.ID, objc.Sel("setMaximumDrawableCount:"), value)
}
// A Boolean value that determines whether requests for a new buffer expire if
// the system can’t satisfy them.
//
// # Discussion
// 
// If [true], the [NextDrawable] method returns [nil] if it can’t provide a
// drawable object within one second. If [false], the [NextDrawable] method
// waits indefinitely for a drawable to become available.
// 
// The default value is [true].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [nil]: https://developer.apple.com/documentation/ObjectiveC/nil-227m0
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/allowsNextDrawableTimeout
func (m CAMetalLayer) AllowsNextDrawableTimeout() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("allowsNextDrawableTimeout"))
	return rv
}
func (m CAMetalLayer) SetAllowsNextDrawableTimeout(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setAllowsNextDrawableTimeout:"), value)
}
// The properties of the Metal performance heads-up display.
//
// # Discussion
// 
// The Metal performance HUD provides real-time statistics and logging,
// including CPU and GPU render time and frame-presentation deadlines.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/developerHUDProperties
func (m CAMetalLayer) DeveloperHUDProperties() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("developerHUDProperties"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m CAMetalLayer) SetDeveloperHUDProperties(value foundation.INSDictionary) {
	objc.Send[struct{}](m.ID, objc.Sel("setDeveloperHUDProperties:"), value)
}
// See: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/residencySet
func (m CAMetalLayer) ResidencySet() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("residencySet"))
	return objectivec.Object{ID: rv}
}
// A positive integer that identifies the drawable.
//
// See: https://developer.apple.com/documentation/Metal/MTLDrawable/drawableID
func (m CAMetalLayer) DrawableID() int {
	rv := objc.Send[int](m.ID, objc.Sel("drawableID"))
	return rv
}
func (m CAMetalLayer) SetDrawableID(value int) {
	objc.Send[struct{}](m.ID, objc.Sel("setDrawableID:"), value)
}
// The host time, in seconds, when the drawable was displayed onscreen.
//
// See: https://developer.apple.com/documentation/Metal/MTLDrawable/presentedTime
func (m CAMetalLayer) PresentedTime() float64 {
	rv := objc.Send[float64](m.ID, objc.Sel("presentedTime"))
	return rv
}
func (m CAMetalLayer) SetPresentedTime(value float64) {
	objc.Send[struct{}](m.ID, objc.Sel("setPresentedTime:"), value)
}

// Returns the class used to create the layer for instances of this class.
//
// See: https://developer.apple.com/documentation/UIKit/UIView/layerClass
func (_CAMetalLayerClass CAMetalLayerClass) LayerClass() objc.Class {
	rv := objc.Send[objc.Class](objc.ID(_CAMetalLayerClass.class), objc.Sel("layerClass"))
	return rv
}
func (_CAMetalLayerClass CAMetalLayerClass) SetLayerClass(value objc.Class) {
	objc.Send[struct{}](objc.ID(_CAMetalLayerClass.class), objc.Sel("setLayerClass:"), value)
}

