// Code generated from Apple documentation for MetalKit. DO NOT EDIT.

package metalkit

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MTKView] class.
var (
	_MTKViewClass     MTKViewClass
	_MTKViewClassOnce sync.Once
)

func getMTKViewClass() MTKViewClass {
	_MTKViewClassOnce.Do(func() {
		_MTKViewClass = MTKViewClass{class: objc.GetClass("MTKView")}
	})
	return _MTKViewClass
}

// GetMTKViewClass returns the class object for MTKView.
func GetMTKViewClass() MTKViewClass {
	return getMTKViewClass()
}

type MTKViewClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MTKViewClass) Alloc() MTKView {
	rv := objc.Send[MTKView](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A specialized view that creates, configures, and displays Metal objects.
//
// # Overview
// 
// The [MTKView] class provides a default implementation of a Metal-aware view
// that you can use to render graphics using Metal and display them onscreen.
// When asked, the view provides a [MTLRenderPassDescriptor] object that
// points at a texture for you to render new contents into. Optionally, an
// [MTKView] can create depth and stencil textures for you and any
// intermediate textures needed for antialiasing. The view uses a
// [CAMetalLayer] to manage the Metal drawable objects.
// 
// The view requires a [MTLDevice] object to manage the Metal objects it
// creates for you. You must set the [MTKView.Device] property and, optionally, modify
// the view’s drawable properties before drawing.
// 
// # Configuring the Drawing Behavior
// 
// The MTKView class supports three drawing modes:
// 
// - Timed updates: The view redraws its contents based on an internal timer.
// In this case, which is the default behavior, both [MTKView.Paused] and
// [MTKView.EnableSetNeedsDisplay] are set to [false]. Use this mode for games and
// other animated content that’s regularly updated. - Draw notifications:
// The view redraws itself when something invalidates its contents, usually
// because of a call to [setNeedsDisplay()] or some other view-related
// behavior. In this case, set [MTKView.Paused] and [MTKView.EnableSetNeedsDisplay] to [true].
// Use this mode for apps with a more traditional workflow, where updates
// happen when data changes, but not on a regular timed interval. - Explicit
// drawing: The view redraws its contents only when you explicitly call the
// [MTKView.Draw] method. In this case, set [MTKView.Paused] to [true] and
// [MTKView.EnableSetNeedsDisplay] to [false]. Use this mode to create your own custom
// workflow.
// 
// # Drawing the View’s Contents
// 
// Regardless of drawing mode, when the view needs to update its contents, it
// calls the [draw(_:)] method when that method has been overridden by a
// subclass, or [DrawInMTKView] on the view’s delegate if the subclass
// doesn’t override it. You should either subclass [MTKView] or provide a
// delegate, but not both.
// 
// In your drawing method, you obtain a render pass descriptor from the view,
// render into it, and then present the associated drawable.
// 
// # Obtaining a Drawable from a MetalKit View
// 
// Each [MTKView] is backed by a [CAMetalLayer]. In your renderer, implement
// the [MTKViewDelegate] protocol to interact with a MetalKit view. Call the
// MetalKit view’s [MTKView.CurrentRenderPassDescriptor] property to obtain a render
// pass descriptor configured for the current frame:
// 
// When you read this property, Core Animation implicitly obtains a drawable
// for the current frame and stores it in the [MTKView.CurrentDrawable] property. It
// then configures a render pass descriptor to draw into that drawable,
// including any depth, stencil, and antialiasing textures as necessary. The
// view configures this render pass using the default store and load actions.
// You can adjust the descriptor further before using it to create a
// [MTLRenderCommandEncoder].
// 
// Obtain drawables as late as possible; preferably, immediately before
// encoding your onscreen render pass.
// 
// # Registering the Drawable’s Presentation
// 
// After rendering the contents, you must present the drawable to update the
// view’s contents. The most convenient way to present the content is to
// call the [present(_:)] method on the command buffer. Then, call the
// [commit()] method to submit the command buffer to a GPU:
// 
// When a command queue schedules a command buffer for execution, the drawable
// tracks all render or write requests on itself in that command buffer. The
// operating system doesn’t present the drawable onscreen until the commands
// have finished executing. By asking the command buffer to present the
// drawable, you guarantee that presentation happens after the command queue
// has scheduled this command buffer. Don’t wait for the command buffer to
// finish executing before registering the drawable’s presentation.
//
// [CAMetalLayer]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer
// [MTLDevice]: https://developer.apple.com/documentation/Metal/MTLDevice
// [MTLRenderCommandEncoder]: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder
// [MTLRenderPassDescriptor]: https://developer.apple.com/documentation/Metal/MTLRenderPassDescriptor
// [commit()]: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/commit()
// [draw(_:)]: https://developer.apple.com/documentation/AppKit/NSView/draw(_:)
// [false]: https://developer.apple.com/documentation/Swift/false
// [present(_:)]: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/present(_:)
// [setNeedsDisplay()]: https://developer.apple.com/documentation/UIKit/UIView/setNeedsDisplay()
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Creating a View
//
//   - [MTKView.InitWithFrameDevice]: Initializes a view with the specified frame rectangle and Metal device.
//
// # Configuring the Delegate
//
//   - [MTKView.Delegate]: The view’s delegate.
//   - [MTKView.SetDelegate]
//
// # Configuring the Metal Device
//
//   - [MTKView.Device]: The device object the view uses to create its Metal objects.
//   - [MTKView.SetDevice]
//   - [MTKView.PreferredDevice]: The device object that the system recommends using for this view.
//
// # Configuring the Color Render Target
//
//   - [MTKView.ColorPixelFormat]: The color pixel format for the current drawable’s texture.
//   - [MTKView.SetColorPixelFormat]
//   - [MTKView.Colorspace]: The color space of the rendered content.
//   - [MTKView.SetColorspace]
//   - [MTKView.FramebufferOnly]: A Boolean value that determines whether the drawable’s textures are used only for rendering.
//   - [MTKView.SetFramebufferOnly]
//   - [MTKView.DrawableSize]: The current size of drawable textures.
//   - [MTKView.SetDrawableSize]
//   - [MTKView.PreferredDrawableSize]: The recommended dimensions of the drawable.
//   - [MTKView.AutoResizeDrawable]: A Boolean value that controls whether to resize the drawable as the view changes size.
//   - [MTKView.SetAutoResizeDrawable]
//   - [MTKView.ClearColor]: The color to use to clear the color target when creating a render pass descriptor.
//   - [MTKView.SetClearColor]
//
// # Configuring the Render Target Properties
//
//   - [MTKView.DepthStencilPixelFormat]: The format used to generate the [depthStencilTexture](<doc://com.apple.metalkit/documentation/MetalKit/MTKView/depthStencilTexture>) object.
//   - [MTKView.SetDepthStencilPixelFormat]
//   - [MTKView.DepthStencilAttachmentTextureUsage]: The texture usage characteristics that the view uses when creating the depth and stencil textures.
//   - [MTKView.SetDepthStencilAttachmentTextureUsage]
//   - [MTKView.ClearDepth]: The depth value to use to clear the depth target when creating a render pass descriptor.
//   - [MTKView.SetClearDepth]
//   - [MTKView.ClearStencil]: The stencil value to use to clear the stencil target when creating a render pass descriptor.
//   - [MTKView.SetClearStencil]
//
// # Configuring Multisampling
//
//   - [MTKView.SampleCount]: The sample count used to generate the [multisampleColorTexture](<doc://com.apple.metalkit/documentation/MetalKit/MTKView/multisampleColorTexture>) object.
//   - [MTKView.SetSampleCount]
//   - [MTKView.MultisampleColorAttachmentTextureUsage]: The texture usage characteristics that the view uses when creating multisample textures.
//   - [MTKView.SetMultisampleColorAttachmentTextureUsage]
//
// # Retrieving Render Target Information
//
//   - [MTKView.CurrentRenderPassDescriptor]: A render pass descriptor to draw into the current drawable.
//   - [MTKView.CurrentDrawable]: The drawable to use for the current frame.
//   - [MTKView.DepthStencilTexture]: A packed depth and stencil texture associated with the current drawable object’s texture.
//   - [MTKView.DepthStencilStorageMode]: The storage mode that the packed depth and stencil texture use.
//   - [MTKView.SetDepthStencilStorageMode]
//   - [MTKView.MultisampleColorTexture]: The multisample color sample texture to render into.
//
// # Configuring Drawing Behavior
//
//   - [MTKView.PreferredFramesPerSecond]: The rate at which the view redraws its contents.
//   - [MTKView.SetPreferredFramesPerSecond]
//   - [MTKView.Paused]: A Boolean value that indicates whether the draw loop is paused.
//   - [MTKView.SetPaused]
//   - [MTKView.EnableSetNeedsDisplay]: A Boolean value that indicates whether the view responds to [setNeedsDisplay()](<doc://com.apple.documentation/documentation/UIKit/UIView/setNeedsDisplay()>).
//   - [MTKView.SetEnableSetNeedsDisplay]
//   - [MTKView.Draw]: Redraws the view’s contents immediately.
//   - [MTKView.PresentsWithTransaction]: A Boolean value that determines whether the view presents its content using a Core Animation transaction.
//   - [MTKView.SetPresentsWithTransaction]
//
// # Releasing Memory
//
//   - [MTKView.ReleaseDrawables]: Releases the [depthStencilTexture](<doc://com.apple.metalkit/documentation/MetalKit/MTKView/depthStencilTexture>) and [multisampleColorTexture](<doc://com.apple.metalkit/documentation/MetalKit/MTKView/multisampleColorTexture>) objects.
//
// # Instance Properties
//
//   - [MTKView.CurrentMTL4RenderPassDescriptor]
//
// See: https://developer.apple.com/documentation/MetalKit/MTKView
type MTKView struct {
	appkit.NSView
}

// MTKViewFromID constructs a [MTKView] from an objc.ID.
//
// A specialized view that creates, configures, and displays Metal objects.
func MTKViewFromID(id objc.ID) MTKView {
	return MTKView{NSView: appkit.NSViewFromID(id)}
}
// NOTE: MTKView adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MTKView] class.
//
// # Creating a View
//
//   - [IMTKView.InitWithFrameDevice]: Initializes a view with the specified frame rectangle and Metal device.
//
// # Configuring the Delegate
//
//   - [IMTKView.Delegate]: The view’s delegate.
//   - [IMTKView.SetDelegate]
//
// # Configuring the Metal Device
//
//   - [IMTKView.Device]: The device object the view uses to create its Metal objects.
//   - [IMTKView.SetDevice]
//   - [IMTKView.PreferredDevice]: The device object that the system recommends using for this view.
//
// # Configuring the Color Render Target
//
//   - [IMTKView.ColorPixelFormat]: The color pixel format for the current drawable’s texture.
//   - [IMTKView.SetColorPixelFormat]
//   - [IMTKView.Colorspace]: The color space of the rendered content.
//   - [IMTKView.SetColorspace]
//   - [IMTKView.FramebufferOnly]: A Boolean value that determines whether the drawable’s textures are used only for rendering.
//   - [IMTKView.SetFramebufferOnly]
//   - [IMTKView.DrawableSize]: The current size of drawable textures.
//   - [IMTKView.SetDrawableSize]
//   - [IMTKView.PreferredDrawableSize]: The recommended dimensions of the drawable.
//   - [IMTKView.AutoResizeDrawable]: A Boolean value that controls whether to resize the drawable as the view changes size.
//   - [IMTKView.SetAutoResizeDrawable]
//   - [IMTKView.ClearColor]: The color to use to clear the color target when creating a render pass descriptor.
//   - [IMTKView.SetClearColor]
//
// # Configuring the Render Target Properties
//
//   - [IMTKView.DepthStencilPixelFormat]: The format used to generate the [depthStencilTexture](<doc://com.apple.metalkit/documentation/MetalKit/MTKView/depthStencilTexture>) object.
//   - [IMTKView.SetDepthStencilPixelFormat]
//   - [IMTKView.DepthStencilAttachmentTextureUsage]: The texture usage characteristics that the view uses when creating the depth and stencil textures.
//   - [IMTKView.SetDepthStencilAttachmentTextureUsage]
//   - [IMTKView.ClearDepth]: The depth value to use to clear the depth target when creating a render pass descriptor.
//   - [IMTKView.SetClearDepth]
//   - [IMTKView.ClearStencil]: The stencil value to use to clear the stencil target when creating a render pass descriptor.
//   - [IMTKView.SetClearStencil]
//
// # Configuring Multisampling
//
//   - [IMTKView.SampleCount]: The sample count used to generate the [multisampleColorTexture](<doc://com.apple.metalkit/documentation/MetalKit/MTKView/multisampleColorTexture>) object.
//   - [IMTKView.SetSampleCount]
//   - [IMTKView.MultisampleColorAttachmentTextureUsage]: The texture usage characteristics that the view uses when creating multisample textures.
//   - [IMTKView.SetMultisampleColorAttachmentTextureUsage]
//
// # Retrieving Render Target Information
//
//   - [IMTKView.CurrentRenderPassDescriptor]: A render pass descriptor to draw into the current drawable.
//   - [IMTKView.CurrentDrawable]: The drawable to use for the current frame.
//   - [IMTKView.DepthStencilTexture]: A packed depth and stencil texture associated with the current drawable object’s texture.
//   - [IMTKView.DepthStencilStorageMode]: The storage mode that the packed depth and stencil texture use.
//   - [IMTKView.SetDepthStencilStorageMode]
//   - [IMTKView.MultisampleColorTexture]: The multisample color sample texture to render into.
//
// # Configuring Drawing Behavior
//
//   - [IMTKView.PreferredFramesPerSecond]: The rate at which the view redraws its contents.
//   - [IMTKView.SetPreferredFramesPerSecond]
//   - [IMTKView.Paused]: A Boolean value that indicates whether the draw loop is paused.
//   - [IMTKView.SetPaused]
//   - [IMTKView.EnableSetNeedsDisplay]: A Boolean value that indicates whether the view responds to [setNeedsDisplay()](<doc://com.apple.documentation/documentation/UIKit/UIView/setNeedsDisplay()>).
//   - [IMTKView.SetEnableSetNeedsDisplay]
//   - [IMTKView.Draw]: Redraws the view’s contents immediately.
//   - [IMTKView.PresentsWithTransaction]: A Boolean value that determines whether the view presents its content using a Core Animation transaction.
//   - [IMTKView.SetPresentsWithTransaction]
//
// # Releasing Memory
//
//   - [IMTKView.ReleaseDrawables]: Releases the [depthStencilTexture](<doc://com.apple.metalkit/documentation/MetalKit/MTKView/depthStencilTexture>) and [multisampleColorTexture](<doc://com.apple.metalkit/documentation/MetalKit/MTKView/multisampleColorTexture>) objects.
//
// # Instance Properties
//
//   - [IMTKView.CurrentMTL4RenderPassDescriptor]
//
// See: https://developer.apple.com/documentation/MetalKit/MTKView
type IMTKView interface {
	appkit.INSView

	// Topic: Creating a View

	// Initializes a view with the specified frame rectangle and Metal device.
	InitWithFrameDevice(frameRect corefoundation.CGRect, device metal.MTLDevice) MTKView

	// Topic: Configuring the Delegate

	// The view’s delegate.
	Delegate() MTKViewDelegate
	SetDelegate(value MTKViewDelegate)

	// Topic: Configuring the Metal Device

	// The device object the view uses to create its Metal objects.
	Device() metal.MTLDevice
	SetDevice(value metal.MTLDevice)
	// The device object that the system recommends using for this view.
	PreferredDevice() metal.MTLDevice

	// Topic: Configuring the Color Render Target

	// The color pixel format for the current drawable’s texture.
	ColorPixelFormat() unsafe.Pointer
	SetColorPixelFormat(value metal.MTLPixelFormat)
	// The color space of the rendered content.
	Colorspace() coregraphics.CGColorSpaceRef
	SetColorspace(value coregraphics.CGColorSpaceRef)
	// A Boolean value that determines whether the drawable’s textures are used only for rendering.
	FramebufferOnly() bool
	SetFramebufferOnly(value bool)
	// The current size of drawable textures.
	DrawableSize() corefoundation.CGSize
	SetDrawableSize(value corefoundation.CGSize)
	// The recommended dimensions of the drawable.
	PreferredDrawableSize() corefoundation.CGSize
	// A Boolean value that controls whether to resize the drawable as the view changes size.
	AutoResizeDrawable() bool
	SetAutoResizeDrawable(value bool)
	// The color to use to clear the color target when creating a render pass descriptor.
	ClearColor() metal.MTLClearColor
	SetClearColor(value metal.MTLClearColor)

	// Topic: Configuring the Render Target Properties

	// The format used to generate the [depthStencilTexture](<doc://com.apple.metalkit/documentation/MetalKit/MTKView/depthStencilTexture>) object.
	DepthStencilPixelFormat() unsafe.Pointer
	SetDepthStencilPixelFormat(value metal.MTLPixelFormat)
	// The texture usage characteristics that the view uses when creating the depth and stencil textures.
	DepthStencilAttachmentTextureUsage() unsafe.Pointer
	SetDepthStencilAttachmentTextureUsage(value metal.MTLTextureUsage)
	// The depth value to use to clear the depth target when creating a render pass descriptor.
	ClearDepth() float64
	SetClearDepth(value float64)
	// The stencil value to use to clear the stencil target when creating a render pass descriptor.
	ClearStencil() uint32
	SetClearStencil(value uint32)

	// Topic: Configuring Multisampling

	// The sample count used to generate the [multisampleColorTexture](<doc://com.apple.metalkit/documentation/MetalKit/MTKView/multisampleColorTexture>) object.
	SampleCount() uint
	SetSampleCount(value uint)
	// The texture usage characteristics that the view uses when creating multisample textures.
	MultisampleColorAttachmentTextureUsage() unsafe.Pointer
	SetMultisampleColorAttachmentTextureUsage(value metal.MTLTextureUsage)

	// Topic: Retrieving Render Target Information

	// A render pass descriptor to draw into the current drawable.
	CurrentRenderPassDescriptor() metal.MTLRenderPassDescriptor
	// The drawable to use for the current frame.
	CurrentDrawable() objectivec.IObject
	// A packed depth and stencil texture associated with the current drawable object’s texture.
	DepthStencilTexture() metal.MTLTexture
	// The storage mode that the packed depth and stencil texture use.
	DepthStencilStorageMode() unsafe.Pointer
	SetDepthStencilStorageMode(value metal.MTLStorageMode)
	// The multisample color sample texture to render into.
	MultisampleColorTexture() metal.MTLTexture

	// Topic: Configuring Drawing Behavior

	// The rate at which the view redraws its contents.
	PreferredFramesPerSecond() int
	SetPreferredFramesPerSecond(value int)
	// A Boolean value that indicates whether the draw loop is paused.
	Paused() bool
	SetPaused(value bool)
	// A Boolean value that indicates whether the view responds to [setNeedsDisplay()](<doc://com.apple.documentation/documentation/UIKit/UIView/setNeedsDisplay()>).
	EnableSetNeedsDisplay() bool
	SetEnableSetNeedsDisplay(value bool)
	// Redraws the view’s contents immediately.
	Draw()
	// A Boolean value that determines whether the view presents its content using a Core Animation transaction.
	PresentsWithTransaction() bool
	SetPresentsWithTransaction(value bool)

	// Topic: Releasing Memory

	// Releases the [depthStencilTexture](<doc://com.apple.metalkit/documentation/MetalKit/MTKView/depthStencilTexture>) and [multisampleColorTexture](<doc://com.apple.metalkit/documentation/MetalKit/MTKView/multisampleColorTexture>) objects.
	ReleaseDrawables()

	// Topic: Instance Properties

	CurrentMTL4RenderPassDescriptor() metal.MTL4RenderPassDescriptor
}

// Init initializes the instance.
func (v MTKView) Init() MTKView {
	rv := objc.Send[MTKView](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v MTKView) Autorelease() MTKView {
	rv := objc.Send[MTKView](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewMTKView creates a new MTKView instance.
func NewMTKView() MTKView {
	class := getMTKViewClass()
	rv := objc.Send[MTKView](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a view from data in a given unarchiver.
//
// coder: An unarchiver object.
//
// # Return Value
// 
// An initialized view object.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKView/init(coder:)
func NewViewWithCoder(coder foundation.INSCoder) MTKView {
	instance := getMTKViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MTKViewFromID(rv)
}

// Initializes a view with the specified frame rectangle and Metal device.
//
// frameRect: The frame rectangle for the view.
//
// device: The Metal device object to use.
//
// # Return Value
// 
// An initialized view object.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKView/init(frame:device:)
func NewViewWithFrameDevice(frameRect corefoundation.CGRect, device metal.MTLDevice) MTKView {
	instance := getMTKViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:device:"), frameRect, device)
	return MTKViewFromID(rv)
}

// Initializes a view with the specified frame rectangle and Metal device.
//
// frameRect: The frame rectangle for the view.
//
// device: The Metal device object to use.
//
// # Return Value
// 
// An initialized view object.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKView/init(frame:device:)
func (v MTKView) InitWithFrameDevice(frameRect corefoundation.CGRect, device metal.MTLDevice) MTKView {
	rv := objc.Send[MTKView](v.ID, objc.Sel("initWithFrame:device:"), frameRect, device)
	return rv
}
// Redraws the view’s contents immediately.
//
// # Discussion
// 
// This method manually tells the view to redraw its contents. Calling this
// method causes the view to call either the [DrawInMTKView] method of the
// view’s [Delegate], or the [draw(_:)] method of the [MTKView] subclass.
// Never call this method inside either drawing function.
//
// [draw(_:)]: https://developer.apple.com/documentation/UIKit/UIView/draw(_:)
//
// See: https://developer.apple.com/documentation/MetalKit/MTKView/draw()
func (v MTKView) Draw() {
	objc.Send[objc.ID](v.ID, objc.Sel("draw"))
}
// Releases the [DepthStencilTexture] and [MultisampleColorTexture] objects.
//
// # Discussion
// 
// Call this method when your app is moving to the background or when the view
// won’t display content for a significant period of time. The texture
// objects that this class creates consume a large amount of memory, so
// freeing them makes that memory available to other parts of your app.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKView/releaseDrawables()
func (v MTKView) ReleaseDrawables() {
	objc.Send[objc.ID](v.ID, objc.Sel("releaseDrawables"))
}

// The view’s delegate.
//
// # Discussion
// 
// A delegate is optional. If you provide one, the view calls the delegate
// when it needs to update its contents. You should either provide a delegate
// or subclass the view to override the [draw(_:)] method, but not both.
//
// [draw(_:)]: https://developer.apple.com/documentation/UIKit/UIView/draw(_:)
//
// See: https://developer.apple.com/documentation/MetalKit/MTKView/delegate
func (v MTKView) Delegate() MTKViewDelegate {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("delegate"))
	return MTKViewDelegateObjectFromID(rv)
}
func (v MTKView) SetDelegate(value MTKViewDelegate) {
	objc.Send[struct{}](v.ID, objc.Sel("setDelegate:"), value)
}
// The device object the view uses to create its Metal objects.
//
// # Discussion
// 
// The default value is `nil`. You must explicitly set the device object.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKView/device
func (v MTKView) Device() metal.MTLDevice {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("device"))
	return metal.MTLDeviceObjectFromID(rv)
}
func (v MTKView) SetDevice(value metal.MTLDevice) {
	objc.Send[struct{}](v.ID, objc.Sel("setDevice:"), value)
}
// The device object that the system recommends using for this view.
//
// # Discussion
// 
// On systems with a single GPU, this method returns the default Metal system
// device; see [MTLCreateSystemDefaultDevice()]. On systems with more than one
// GPU, this method returns the [MTLDevice] that was last used to composite
// and present the view’s contents. This device object usually corresponds
// to the GPU associated with the screen that’s displaying the view. If you
// set the view’s [Device] property to this device object, you reduce the
// number of cross-GPU texture copies that Core Animation must make to present
// the view’s contents onscreen.
//
// [MTLCreateSystemDefaultDevice()]: https://developer.apple.com/documentation/Metal/MTLCreateSystemDefaultDevice()
// [MTLDevice]: https://developer.apple.com/documentation/Metal/MTLDevice
//
// See: https://developer.apple.com/documentation/MetalKit/MTKView/preferredDevice
func (v MTKView) PreferredDevice() metal.MTLDevice {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("preferredDevice"))
	return metal.MTLDeviceObjectFromID(rv)
}
// The color pixel format for the current drawable’s texture.
//
// # Discussion
// 
// The pixel format must be one that the underlying [CAMetalLayer] can use.
// See [pixelFormat].
// 
// The default value is [MTLPixelFormat.bgra8Unorm].
//
// [CAMetalLayer]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer
// [MTLPixelFormat.bgra8Unorm]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/bgra8Unorm
// [pixelFormat]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/pixelFormat
//
// See: https://developer.apple.com/documentation/MetalKit/MTKView/colorPixelFormat
func (v MTKView) ColorPixelFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v.ID, objc.Sel("colorPixelFormat"))
	return rv
}
func (v MTKView) SetColorPixelFormat(value metal.MTLPixelFormat) {
	objc.Send[struct{}](v.ID, objc.Sel("setColorPixelFormat:"), value)
}
// The color space of the rendered content.
//
// # Discussion
// 
// The default value is `nil`, indicating that the rendered content isn’t
// color-matched. If you set this to a different color space, Core Animation
// performs any necessary color transformations when compositing the view’s
// contents.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKView/colorspace
func (v MTKView) Colorspace() coregraphics.CGColorSpaceRef {
	rv := objc.Send[coregraphics.CGColorSpaceRef](v.ID, objc.Sel("colorspace"))
	return coregraphics.CGColorSpaceRef(rv)
}
func (v MTKView) SetColorspace(value coregraphics.CGColorSpaceRef) {
	objc.Send[struct{}](v.ID, objc.Sel("setColorspace:"), value)
}
// A Boolean value that determines whether the drawable’s textures are used
// only for rendering.
//
// # Discussion
// 
// If the value is [true] (the default), the underlying [CAMetalLayer] object
// allocates its textures with only the [renderTarget] usage flag. Core
// Animation can then optimize the textures for display purposes. However, you
// may not sample, read from, or write to those textures. If the value is
// [false], you can sample or perform read/write operations on the textures,
// but at a cost to performance.
//
// [CAMetalLayer]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer
// [false]: https://developer.apple.com/documentation/Swift/false
// [renderTarget]: https://developer.apple.com/documentation/Metal/MTLTextureUsage/renderTarget
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/MetalKit/MTKView/framebufferOnly
func (v MTKView) FramebufferOnly() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("framebufferOnly"))
	return rv
}
func (v MTKView) SetFramebufferOnly(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setFramebufferOnly:"), value)
}
// The current size of drawable textures.
//
// # Discussion
// 
// Changing this value adjusts the size of any color, depth, stencil, and
// multisampling textures created by the view. If [AutoResizeDrawable] is
// [true], this property is updated automatically whenever the view’s size
// changes. If [AutoResizeDrawable] is [false], set this value to change the
// size of the texture objects.
// 
// The default value is derived from the current view’s size, in native
// pixels.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/MetalKit/MTKView/drawableSize
func (v MTKView) DrawableSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](v.ID, objc.Sel("drawableSize"))
	return corefoundation.CGSize(rv)
}
func (v MTKView) SetDrawableSize(value corefoundation.CGSize) {
	objc.Send[struct{}](v.ID, objc.Sel("setDrawableSize:"), value)
}
// The recommended dimensions of the drawable.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKView/preferredDrawableSize
func (v MTKView) PreferredDrawableSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](v.ID, objc.Sel("preferredDrawableSize"))
	return corefoundation.CGSize(rv)
}
// A Boolean value that controls whether to resize the drawable as the view
// changes size.
//
// # Discussion
// 
// If the value is [true], the view automatically resizes its underlying
// color, depth, stencil, and multisample textures when the view is resized.
// If the value is [false], you must explicitly set [DrawableSize] to change
// the size of these objects.
// 
// The default value is [true].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/MetalKit/MTKView/autoResizeDrawable
func (v MTKView) AutoResizeDrawable() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("autoResizeDrawable"))
	return rv
}
func (v MTKView) SetAutoResizeDrawable(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setAutoResizeDrawable:"), value)
}
// The color to use to clear the color target when creating a render pass
// descriptor.
//
// # Discussion
// 
// When the view creates a render pass, it sets the load action for the color
// render target to [MTLLoadAction.clear] and uses this color as the clear
// color. The default value is `(0.0, 0.0, 0.0, 1.0)`.
//
// [MTLLoadAction.clear]: https://developer.apple.com/documentation/Metal/MTLLoadAction/clear
//
// See: https://developer.apple.com/documentation/MetalKit/MTKView/clearColor
func (v MTKView) ClearColor() metal.MTLClearColor {
	rv := objc.Send[metal.MTLClearColor](v.ID, objc.Sel("clearColor"))
	return metal.MTLClearColor(rv)
}
func (v MTKView) SetClearColor(value metal.MTLClearColor) {
	objc.Send[struct{}](v.ID, objc.Sel("setClearColor:"), value)
}
// The format used to generate the [DepthStencilTexture] object.
//
// # Discussion
// 
// The default value is [MTLPixelFormat.invalid], which means that the view
// doesn’t create a depth and stencil texture. If you set it to a different
// format, the view automatically creates those textures for you and
// configures them as part of any render passes that the view creates.
//
// [MTLPixelFormat.invalid]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/invalid
//
// See: https://developer.apple.com/documentation/MetalKit/MTKView/depthStencilPixelFormat
func (v MTKView) DepthStencilPixelFormat() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v.ID, objc.Sel("depthStencilPixelFormat"))
	return rv
}
func (v MTKView) SetDepthStencilPixelFormat(value metal.MTLPixelFormat) {
	objc.Send[struct{}](v.ID, objc.Sel("setDepthStencilPixelFormat:"), value)
}
// The texture usage characteristics that the view uses when creating the
// depth and stencil textures.
//
// # Discussion
// 
// The default value is [renderTarget].
//
// [renderTarget]: https://developer.apple.com/documentation/Metal/MTLTextureUsage/renderTarget
//
// See: https://developer.apple.com/documentation/MetalKit/MTKView/depthStencilAttachmentTextureUsage
func (v MTKView) DepthStencilAttachmentTextureUsage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v.ID, objc.Sel("depthStencilAttachmentTextureUsage"))
	return rv
}
func (v MTKView) SetDepthStencilAttachmentTextureUsage(value metal.MTLTextureUsage) {
	objc.Send[struct{}](v.ID, objc.Sel("setDepthStencilAttachmentTextureUsage:"), value)
}
// The depth value to use to clear the depth target when creating a render
// pass descriptor.
//
// # Discussion
// 
// If you specified that you want a depth texture, the view configures any
// render passes to use the depth texture, with a load action of
// [MTLLoadAction.clear] and the value of this property as the value to clear
// it to. The default value is `1.0`.
//
// [MTLLoadAction.clear]: https://developer.apple.com/documentation/Metal/MTLLoadAction/clear
//
// See: https://developer.apple.com/documentation/MetalKit/MTKView/clearDepth
func (v MTKView) ClearDepth() float64 {
	rv := objc.Send[float64](v.ID, objc.Sel("clearDepth"))
	return rv
}
func (v MTKView) SetClearDepth(value float64) {
	objc.Send[struct{}](v.ID, objc.Sel("setClearDepth:"), value)
}
// The stencil value to use to clear the stencil target when creating a render
// pass descriptor.
//
// # Discussion
// 
// If you specified that you want a stencil texture, the view configures any
// render passes to use the stencil texture, with a load action of
// [MTLLoadAction.clear] and the value of this property as the value to clear
// it to. The default value is `0`.
//
// [MTLLoadAction.clear]: https://developer.apple.com/documentation/Metal/MTLLoadAction/clear
//
// See: https://developer.apple.com/documentation/MetalKit/MTKView/clearStencil
func (v MTKView) ClearStencil() uint32 {
	rv := objc.Send[uint32](v.ID, objc.Sel("clearStencil"))
	return rv
}
func (v MTKView) SetClearStencil(value uint32) {
	objc.Send[struct{}](v.ID, objc.Sel("setClearStencil:"), value)
}
// The sample count used to generate the [MultisampleColorTexture] object.
//
// # Discussion
// 
// Support for different sample count values varies by device object. Call the
// [supportsTextureSampleCount(_:)] method to determine if the device object
// supports the sample count you want.
// 
// The default value is `1`. When you set a value greater than `1`, the view
// creates and configures an intermediate set of multisample textures. The
// pixel format is the same as the one specified for the drawable; see
// [ColorPixelFormat]. When the view creates a render pass descriptor, the
// render pass uses those intermediate textures as the color render targets,
// with a store action to resolve these multisample textures into the
// drawable’s texture ([MTLStoreAction.multisampleResolve]).
//
// [MTLStoreAction.multisampleResolve]: https://developer.apple.com/documentation/Metal/MTLStoreAction/multisampleResolve
// [supportsTextureSampleCount(_:)]: https://developer.apple.com/documentation/Metal/MTLDevice/supportsTextureSampleCount(_:)
//
// See: https://developer.apple.com/documentation/MetalKit/MTKView/sampleCount
func (v MTKView) SampleCount() uint {
	rv := objc.Send[uint](v.ID, objc.Sel("sampleCount"))
	return rv
}
func (v MTKView) SetSampleCount(value uint) {
	objc.Send[struct{}](v.ID, objc.Sel("setSampleCount:"), value)
}
// The texture usage characteristics that the view uses when creating
// multisample textures.
//
// # Discussion
// 
// The default value is [renderTarget].
//
// [renderTarget]: https://developer.apple.com/documentation/Metal/MTLTextureUsage/renderTarget
//
// See: https://developer.apple.com/documentation/MetalKit/MTKView/multisampleColorAttachmentTextureUsage
func (v MTKView) MultisampleColorAttachmentTextureUsage() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v.ID, objc.Sel("multisampleColorAttachmentTextureUsage"))
	return rv
}
func (v MTKView) SetMultisampleColorAttachmentTextureUsage(value metal.MTLTextureUsage) {
	objc.Send[struct{}](v.ID, objc.Sel("setMultisampleColorAttachmentTextureUsage:"), value)
}
// A render pass descriptor to draw into the current drawable.
//
// # Discussion
// 
// Reading this property creates and returns a new render pass descriptor to
// render into the current drawable’s texture. [MTKView] doesn’t use this
// descriptor, and there’s no requirement for your application to use it.
// 
// This property is `nil` if the view’s [Device] property isn’t set or if
// [CurrentDrawable] is `nil`. Your app should check that
// [CurrentRenderPassDescriptor] isn’t `nil` before attempting to use it.
// 
// The view configures the render pass as follows:
// 
// - If multisampling isn’t enabled—The color attachment at index 0 of the
// render pass descriptor points to the texture assigned to the current
// drawable, with a load action of [MTLLoadAction.clear] and a store action of
// [MTLStoreAction.store]. - If you’ve enabled multisampling—The color
// attachment at index 0 of the render pass descriptor points to the
// multisample texture, the resolve texture points to the texture assigned to
// the current drawable, and the attachment has a load action of
// [MTLLoadAction.clear] and a store action of
// [MTLStoreAction.multisampleResolve]. - If you’ve specified a depth or
// stencil target—The render pass configures the appropriate targets, with a
// load action of [MTLLoadAction.clear] and a store action of
// [MTLStoreAction.dontCare].
//
// [MTLLoadAction.clear]: https://developer.apple.com/documentation/Metal/MTLLoadAction/clear
// [MTLStoreAction.dontCare]: https://developer.apple.com/documentation/Metal/MTLStoreAction/dontCare
// [MTLStoreAction.multisampleResolve]: https://developer.apple.com/documentation/Metal/MTLStoreAction/multisampleResolve
// [MTLStoreAction.store]: https://developer.apple.com/documentation/Metal/MTLStoreAction/store
//
// See: https://developer.apple.com/documentation/MetalKit/MTKView/currentRenderPassDescriptor
func (v MTKView) CurrentRenderPassDescriptor() metal.MTLRenderPassDescriptor {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("currentRenderPassDescriptor"))
	return metal.MTLRenderPassDescriptorFromID(objc.ID(rv))
}
// The drawable to use for the current frame.
//
// # Discussion
// 
// If all drawable objects are in use, the value of this property is `nil`.
// Your app should check that [CurrentDrawable] isn’t `nil` before
// attempting to draw. The view changes the value of this property only after
// returning from a drawing function, either [draw(_:)] from a subclassed
// instance of the view, or [DrawInMTKView] from the view’s delegate.
// 
// Use a [MTLRenderCommandEncoder] object to render into the drawable’s
// texture and present it for display (typically registered using the
// [present(_:)] method of a command buffer). Try to minimize the time between
// when you fetch the drawable and when you submit the command buffer that
// uses it. For more information, see [CAMetalLayer].
//
// [CAMetalLayer]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer#3385893
// [MTLRenderCommandEncoder]: https://developer.apple.com/documentation/Metal/MTLRenderCommandEncoder
// [draw(_:)]: https://developer.apple.com/documentation/UIKit/UIView/draw(_:)
// [present(_:)]: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/present(_:)
//
// See: https://developer.apple.com/documentation/MetalKit/MTKView/currentDrawable
func (v MTKView) CurrentDrawable() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("currentDrawable"))
	return objectivec.Object{ID: rv}
}
// A packed depth and stencil texture associated with the current drawable
// object’s texture.
//
// # Discussion
// 
// The value of [DepthStencilPixelFormat] determines the format of this
// texture.
// 
// The default value is `nil`. This value is also `nil` if the specified pixel
// format is [MTLPixelFormat.invalid].
//
// [MTLPixelFormat.invalid]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/invalid
//
// See: https://developer.apple.com/documentation/MetalKit/MTKView/depthStencilTexture
func (v MTKView) DepthStencilTexture() metal.MTLTexture {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("depthStencilTexture"))
	return metal.MTLTextureObjectFromID(rv)
}
// The storage mode that the packed depth and stencil texture use.
//
// # Discussion
// 
// The default value is [MTLStorageMode.private].
//
// [MTLStorageMode.private]: https://developer.apple.com/documentation/Metal/MTLStorageMode/private
//
// See: https://developer.apple.com/documentation/MetalKit/MTKView/depthStencilStorageMode
func (v MTKView) DepthStencilStorageMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](v.ID, objc.Sel("depthStencilStorageMode"))
	return rv
}
func (v MTKView) SetDepthStencilStorageMode(value metal.MTLStorageMode) {
	objc.Send[struct{}](v.ID, objc.Sel("setDepthStencilStorageMode:"), value)
}
// The multisample color sample texture to render into.
//
// # Discussion
// 
// The format of this texture is determined by the value of the
// [ColorPixelFormat] and [SampleCount] properties.
// 
// The default value is `nil`. This value is also `nil` if the specified pixel
// format is [MTLPixelFormat.invalid], or if [SampleCount] is less than or
// equal to 1.
//
// [MTLPixelFormat.invalid]: https://developer.apple.com/documentation/Metal/MTLPixelFormat/invalid
//
// See: https://developer.apple.com/documentation/MetalKit/MTKView/multisampleColorTexture
func (v MTKView) MultisampleColorTexture() metal.MTLTexture {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("multisampleColorTexture"))
	return metal.MTLTextureObjectFromID(rv)
}
// The rate at which the view redraws its contents.
//
// # Discussion
// 
// When your application sets its preferred frame rate, the view chooses a
// frame rate as close to that as possible based on the capabilities of the
// screen the view is displayed on. To provide a consistent frame rate, the
// actual frame rate chosen is usually a factor of the maximum refresh rate of
// the screen. For example, if the maximum refresh rate of the screen is `60`
// frames per second, that’s also the highest frame rate the view sets as
// the actual frame rate. However, if you ask for a lower frame rate, the view
// might choose `30`, `20`, or `15` frames per second, or another factor, as
// the actual frame rate.
// 
// Your application should choose a frame rate that it can consistently
// maintain. The default value is `60` frames per second.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKView/preferredFramesPerSecond
func (v MTKView) PreferredFramesPerSecond() int {
	rv := objc.Send[int](v.ID, objc.Sel("preferredFramesPerSecond"))
	return rv
}
func (v MTKView) SetPreferredFramesPerSecond(value int) {
	objc.Send[struct{}](v.ID, objc.Sel("setPreferredFramesPerSecond:"), value)
}
// A Boolean value that indicates whether the draw loop is paused.
//
// # Discussion
// 
// If the value is [false], the view periodically redraws the contents, at a
// frame rate set by the value of [PreferredFramesPerSecond].
// 
// The default value is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/MetalKit/MTKView/isPaused
func (v MTKView) Paused() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("isPaused"))
	return rv
}
func (v MTKView) SetPaused(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setPaused:"), value)
}
// A Boolean value that indicates whether the view responds to
// [setNeedsDisplay()].
//
// [setNeedsDisplay()]: https://developer.apple.com/documentation/UIKit/UIView/setNeedsDisplay()
//
// # Discussion
// 
// If this value and the value of [Paused] are [true], the view behaves
// similarly to a [UIView] object, responding to calls to [setNeedsDisplay()].
// In this case, the view’s internal draw loop is paused and updates are
// event-driven instead.
// 
// The default value is [false].
//
// [UIView]: https://developer.apple.com/documentation/UIKit/UIView
// [false]: https://developer.apple.com/documentation/Swift/false
// [setNeedsDisplay()]: https://developer.apple.com/documentation/UIKit/UIView/setNeedsDisplay()
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/MetalKit/MTKView/enableSetNeedsDisplay
func (v MTKView) EnableSetNeedsDisplay() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("enableSetNeedsDisplay"))
	return rv
}
func (v MTKView) SetEnableSetNeedsDisplay(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setEnableSetNeedsDisplay:"), value)
}
// A Boolean value that determines whether the view presents its content using
// a Core Animation transaction.
//
// # Discussion
// 
// This property mirrors the value on the underlying [CAMetalLayer] object,
// and determines whether the view synchronizes updates to its own contents
// with other content changes in Core Animation. For more information about
// how this property affects your rendering code, see
// [presentsWithTransaction].
//
// [CAMetalLayer]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer
// [presentsWithTransaction]: https://developer.apple.com/documentation/QuartzCore/CAMetalLayer/presentsWithTransaction
//
// See: https://developer.apple.com/documentation/MetalKit/MTKView/presentsWithTransaction
func (v MTKView) PresentsWithTransaction() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("presentsWithTransaction"))
	return rv
}
func (v MTKView) SetPresentsWithTransaction(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setPresentsWithTransaction:"), value)
}
//
// # Discussion
// 
// A render pass descriptor generated from the currentDrawable’s texture and
// the view’s depth, stencil, and sample buffers and clear values.
// 
// This is a convience property. The view does not use this descriptor and
// there is no requirement for an app to use this descriptor.
//
// See: https://developer.apple.com/documentation/MetalKit/MTKView/currentMTL4RenderPassDescriptor
func (v MTKView) CurrentMTL4RenderPassDescriptor() metal.MTL4RenderPassDescriptor {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("currentMTL4RenderPassDescriptor"))
	return metal.MTL4RenderPassDescriptorFromID(objc.ID(rv))
}

