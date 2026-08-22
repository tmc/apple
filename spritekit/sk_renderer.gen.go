// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SKRenderer] class.
var (
	_SKRendererClass     SKRendererClass
	_SKRendererClassOnce sync.Once
)

func getSKRendererClass() SKRendererClass {
	_SKRendererClassOnce.Do(func() {
		_SKRendererClass = SKRendererClass{class: objc.GetClass("SKRenderer")}
	})
	return _SKRendererClass
}

// GetSKRendererClass returns the class object for SKRenderer.
func GetSKRendererClass() SKRendererClass {
	return getSKRendererClass()
}

type SKRendererClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SKRendererClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SKRendererClass) Alloc() SKRenderer {
	rv := objc.Send[SKRenderer](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// An object that renders a scene into a custom Metal rendering pipeline and
// drives the scene update cycle.
//
// # Overview
//
// [SKRenderer] allows an app to mix SpriteKit and Metal content by rendering
// an [SKScene] into a Metal command buffer. The reasons an app may do this
// instead of displaying a scene in [SKView] are:
//
// - Apps that are built in Metal can mix in SpriteKit content. While it’s
// possible to add [SKView] as a subview to a Metal view, plugging
// [SKRenderer] into their Metal pipeline allows layering SpriteKit content at
// a specific z-position. - You might be writing a SpriteKit app and decide
// later to take full control over some portion of renderering by implementing
// it with Metal, yet continue to use SpriteKit for the rest of the app. For
// example, you might write the environmental effects layer of your app that
// does fog, clouds, and rain, with custom Metal shaders, and continue to
// layer content below and above that with SpriteKit.
//
// # First Steps
//
//   - [SKRenderer.Scene]: The scene this renderer will draw into the Metal command buffer.
//   - [SKRenderer.SetScene]
//
// # Rendering the Scene
//
//   - [SKRenderer.RenderWithViewportCommandBufferRenderPassDescriptor]
//   - [SKRenderer.RenderWithViewportRenderCommandEncoderRenderPassDescriptorCommandQueue]
//
// # Driving the Scene’s Update Cycle
//
//   - [SKRenderer.UpdateAtTime]
//
// # Configuring Performance Related Toggles
//
//   - [SKRenderer.IgnoresSiblingOrder]
//   - [SKRenderer.SetIgnoresSiblingOrder]
//   - [SKRenderer.ShouldCullNonVisibleNodes]
//   - [SKRenderer.SetShouldCullNonVisibleNodes]
//
// # Enabling Visual Statistics for Debugging
//
//   - [SKRenderer.ShowsNodeCount]: A Boolean value that indicates whether the view displays an overlay that shows physics bodies that are visible in the scene.
//   - [SKRenderer.SetShowsNodeCount]
//   - [SKRenderer.ShowsDrawCount]: A Boolean value that indicates whether the view displays the number of drawing passes it needed to render the view.
//   - [SKRenderer.SetShowsDrawCount]
//   - [SKRenderer.ShowsQuadCount]: A Boolean value that indicates whether the view displays the number of rectangles used to render the scene.
//   - [SKRenderer.SetShowsQuadCount]
//   - [SKRenderer.ShowsPhysics]: A Boolean value that indicates whether the view displays physics-related debugging information.
//   - [SKRenderer.SetShowsPhysics]
//   - [SKRenderer.ShowsFields]: A Boolean value that indicates whether the view displays information about physics fields in the scene.
//   - [SKRenderer.SetShowsFields]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKRenderer
type SKRenderer struct {
	objectivec.Object
}

// SKRendererFromID constructs a [SKRenderer] from an objc.ID.
//
// An object that renders a scene into a custom Metal rendering pipeline and
// drives the scene update cycle.
func SKRendererFromID(id objc.ID) SKRenderer {
	return SKRenderer{objectivec.Object{ID: id}}
}

// NOTE: SKRenderer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SKRenderer] class.
//
// # First Steps
//
//   - [ISKRenderer.Scene]: The scene this renderer will draw into the Metal command buffer.
//   - [ISKRenderer.SetScene]
//
// # Rendering the Scene
//
//   - [ISKRenderer.RenderWithViewportCommandBufferRenderPassDescriptor]
//   - [ISKRenderer.RenderWithViewportRenderCommandEncoderRenderPassDescriptorCommandQueue]
//
// # Driving the Scene’s Update Cycle
//
//   - [ISKRenderer.UpdateAtTime]
//
// # Configuring Performance Related Toggles
//
//   - [ISKRenderer.IgnoresSiblingOrder]
//   - [ISKRenderer.SetIgnoresSiblingOrder]
//   - [ISKRenderer.ShouldCullNonVisibleNodes]
//   - [ISKRenderer.SetShouldCullNonVisibleNodes]
//
// # Enabling Visual Statistics for Debugging
//
//   - [ISKRenderer.ShowsNodeCount]: A Boolean value that indicates whether the view displays an overlay that shows physics bodies that are visible in the scene.
//   - [ISKRenderer.SetShowsNodeCount]
//   - [ISKRenderer.ShowsDrawCount]: A Boolean value that indicates whether the view displays the number of drawing passes it needed to render the view.
//   - [ISKRenderer.SetShowsDrawCount]
//   - [ISKRenderer.ShowsQuadCount]: A Boolean value that indicates whether the view displays the number of rectangles used to render the scene.
//   - [ISKRenderer.SetShowsQuadCount]
//   - [ISKRenderer.ShowsPhysics]: A Boolean value that indicates whether the view displays physics-related debugging information.
//   - [ISKRenderer.SetShowsPhysics]
//   - [ISKRenderer.ShowsFields]: A Boolean value that indicates whether the view displays information about physics fields in the scene.
//   - [ISKRenderer.SetShowsFields]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKRenderer
type ISKRenderer interface {
	objectivec.IObject

	// Topic: First Steps

	// The scene this renderer will draw into the Metal command buffer.
	Scene() ISKScene
	SetScene(value ISKScene)

	// Topic: Rendering the Scene

	RenderWithViewportCommandBufferRenderPassDescriptor(viewport corefoundation.CGRect, commandBuffer metal.MTLCommandBuffer, renderPassDescriptor metal.MTLRenderPassDescriptor)
	RenderWithViewportRenderCommandEncoderRenderPassDescriptorCommandQueue(viewport corefoundation.CGRect, renderCommandEncoder metal.MTLRenderCommandEncoder, renderPassDescriptor metal.MTLRenderPassDescriptor, commandQueue metal.MTLCommandQueue)

	// Topic: Driving the Scene’s Update Cycle

	UpdateAtTime(currentTime foundation.NSTimeInterval)

	// Topic: Configuring Performance Related Toggles

	IgnoresSiblingOrder() bool
	SetIgnoresSiblingOrder(value bool)
	ShouldCullNonVisibleNodes() bool
	SetShouldCullNonVisibleNodes(value bool)

	// Topic: Enabling Visual Statistics for Debugging

	// A Boolean value that indicates whether the view displays an overlay that shows physics bodies that are visible in the scene.
	ShowsNodeCount() bool
	SetShowsNodeCount(value bool)
	// A Boolean value that indicates whether the view displays the number of drawing passes it needed to render the view.
	ShowsDrawCount() bool
	SetShowsDrawCount(value bool)
	// A Boolean value that indicates whether the view displays the number of rectangles used to render the scene.
	ShowsQuadCount() bool
	SetShowsQuadCount(value bool)
	// A Boolean value that indicates whether the view displays physics-related debugging information.
	ShowsPhysics() bool
	SetShowsPhysics(value bool)
	// A Boolean value that indicates whether the view displays information about physics fields in the scene.
	ShowsFields() bool
	SetShowsFields(value bool)
}

// Init initializes the instance.
func (r SKRenderer) Init() SKRenderer {
	rv := objc.Send[SKRenderer](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r SKRenderer) Autorelease() SKRenderer {
	rv := objc.Send[SKRenderer](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKRenderer creates a new SKRenderer instance.
func NewSKRenderer() SKRenderer {
	class := getSKRendererClass()
	rv := objc.Send[SKRenderer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes with a specific GPU to render into.
//
// device: A Metal device.
//
// # Return Value
//
// A new renderer object.
//
// # Discussion
//
// Pass in the same Metal device that is associated to the Metal command
// buffer passed into
// [SKRenderer.RenderWithViewportCommandBufferRenderPassDescriptor].
//
// See: https://developer.apple.com/documentation/SpriteKit/SKRenderer/init(device:)
func NewRendererWithDevice(device metal.MTLDevice) SKRenderer {
	rv := objc.Send[objc.ID](objc.ID(getSKRendererClass().class), objc.Sel("rendererWithDevice:"), device)
	return SKRendererFromID(rv)
}

// See: https://developer.apple.com/documentation/SpriteKit/SKRenderer/render(withViewport:commandBuffer:renderPassDescriptor:)
func (r SKRenderer) RenderWithViewportCommandBufferRenderPassDescriptor(viewport corefoundation.CGRect, commandBuffer metal.MTLCommandBuffer, renderPassDescriptor metal.MTLRenderPassDescriptor) {
	objc.Send[objc.ID](r.ID, objc.Sel("renderWithViewport:commandBuffer:renderPassDescriptor:"), viewport, commandBuffer, renderPassDescriptor)
}

// See: https://developer.apple.com/documentation/SpriteKit/SKRenderer/render(withViewport:renderCommandEncoder:renderPassDescriptor:commandQueue:)
func (r SKRenderer) RenderWithViewportRenderCommandEncoderRenderPassDescriptorCommandQueue(viewport corefoundation.CGRect, renderCommandEncoder metal.MTLRenderCommandEncoder, renderPassDescriptor metal.MTLRenderPassDescriptor, commandQueue metal.MTLCommandQueue) {
	objc.Send[objc.ID](r.ID, objc.Sel("renderWithViewport:renderCommandEncoder:renderPassDescriptor:commandQueue:"), viewport, renderCommandEncoder, renderPassDescriptor, commandQueue)
}

// See: https://developer.apple.com/documentation/SpriteKit/SKRenderer/update(atTime:)
func (r SKRenderer) UpdateAtTime(currentTime foundation.NSTimeInterval) {
	objc.Send[objc.ID](r.ID, objc.Sel("updateAtTime:"), currentTime)
}

// The scene this renderer will draw into the Metal command buffer.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKRenderer/scene
func (r SKRenderer) Scene() ISKScene {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("scene"))
	return SKSceneFromID(objc.ID(rv))
}
func (r SKRenderer) SetScene(value ISKScene) {
	objc.Send[struct{}](r.ID, objc.Sel("setScene:"), value)
}

// See: https://developer.apple.com/documentation/SpriteKit/SKRenderer/ignoresSiblingOrder
func (r SKRenderer) IgnoresSiblingOrder() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("ignoresSiblingOrder"))
	return rv
}
func (r SKRenderer) SetIgnoresSiblingOrder(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setIgnoresSiblingOrder:"), value)
}

// See: https://developer.apple.com/documentation/SpriteKit/SKRenderer/shouldCullNonVisibleNodes
func (r SKRenderer) ShouldCullNonVisibleNodes() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("shouldCullNonVisibleNodes"))
	return rv
}
func (r SKRenderer) SetShouldCullNonVisibleNodes(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setShouldCullNonVisibleNodes:"), value)
}

// A Boolean value that indicates whether the view displays an overlay that
// shows physics bodies that are visible in the scene.
//
// # Discussion
//
// When you enable this option, it shows the number of nodes currently in the
// scene’s node tree.
//
// You may achieve additional performance gain by actually removing nodes from
// the node tree manually which are off screen. For example, in the case of
// [SKRenderer.ShouldCullNonVisibleNodes], there would be less nodes for
// SpriteKit to test every frame whether they’re on screen.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKRenderer/showsNodeCount
func (r SKRenderer) ShowsNodeCount() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("showsNodeCount"))
	return rv
}
func (r SKRenderer) SetShowsNodeCount(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setShowsNodeCount:"), value)
}

// A Boolean value that indicates whether the view displays the number of
// drawing passes it needed to render the view.
//
// # Discussion
//
// Some operations in SpriteKit can require multiple rendering passes to draw
// a scene’s content. For example, an [SKEffectNode] object must render its
// children into a separate buffer, apply the effect, and then perform another
// pass to blend those results into its parent node. These additional
// rendering passes use more rendering resources, reducing your game’s frame
// rate or increasing its total power consumption. Use the draw count as
// another piece of data when you profile your game’s performance.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKRenderer/showsDrawCount
func (r SKRenderer) ShowsDrawCount() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("showsDrawCount"))
	return rv
}
func (r SKRenderer) SetShowsDrawCount(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setShowsDrawCount:"), value)
}

// A Boolean value that indicates whether the view displays the number of
// rectangles used to render the scene.
//
// # Discussion
//
// SpriteKit converts the node tree into one or more rendering passes. Each
// rendering pass is rendered using a series of textured rectangles (quads).
// The showsQuadCount property allows you to see the total number of quads
// that were used to render the scene’s contents. Use this as another piece
// of data when you profile your game’s performance. In most cases, fewer
// quads is better.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKRenderer/showsQuadCount
func (r SKRenderer) ShowsQuadCount() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("showsQuadCount"))
	return rv
}
func (r SKRenderer) SetShowsQuadCount(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setShowsQuadCount:"), value)
}

// A Boolean value that indicates whether the view displays physics-related
// debugging information.
//
// # Discussion
//
// When this debugging option is enabled, each time a frame is rendered, an
// overlay image is drawn on top of your scene that shows the positions and
// shapes of any physics bodies visible in the scene.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKRenderer/showsPhysics
func (r SKRenderer) ShowsPhysics() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("showsPhysics"))
	return rv
}
func (r SKRenderer) SetShowsPhysics(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setShowsPhysics:"), value)
}

// A Boolean value that indicates whether the view displays information about
// physics fields in the scene.
//
// # Discussion
//
// When this debugging option is enabled, each time a frame is rendered, an
// image is drawn behind your scene that shows the effects of any physics
// fields contained in the scene.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKRenderer/showsFields
func (r SKRenderer) ShowsFields() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("showsFields"))
	return rv
}
func (r SKRenderer) SetShowsFields(value bool) {
	objc.Send[struct{}](r.ID, objc.Sel("setShowsFields:"), value)
}
