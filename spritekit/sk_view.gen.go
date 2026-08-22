// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SKView] class.
var (
	_SKViewClass     SKViewClass
	_SKViewClassOnce sync.Once
)

func getSKViewClass() SKViewClass {
	_SKViewClassOnce.Do(func() {
		_SKViewClass = SKViewClass{class: objc.GetClass("SKView")}
	})
	return _SKViewClass
}

// GetSKViewClass returns the class object for SKView.
func GetSKViewClass() SKViewClass {
	return getSKViewClass()
}

type SKViewClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SKViewClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SKViewClass) Alloc() SKView {
	rv := objc.Send[SKView](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// A view subclass that renders a SpriteKit scene.
//
// # Overview
//
// You present a scene by calling the view’s [SKView.PresentScene] method.
// When a scene is presented by the view, it alternates between running its
// simulation (which animates the content) and rendering the content for
// display. You can pause the scene by setting the view’s [SKView.Paused]
// property to true.
//
// # Displaying a Scene
//
//   - [SKView.Scene]: The scene currently presented by this view.
//   - [SKView.PresentScene]: Presents a scene.
//   - [SKView.PresentSceneTransition]: Transitions from the current scene to a new scene.
//
// # Controlling the Timing of a Scene’s Rendering
//
//   - [SKView.IsPaused]: A Boolean value that indicates whether the view’s scene animations are paused.
//   - [SKView.SetPaused]
//   - [SKView.PreferredFramesPerSecond]: The animation frame rate that the view uses to render its scene.
//   - [SKView.SetPreferredFramesPerSecond]
//   - [SKView.Delegate]: A delegate that allows dynamic control of the view’s render rate.
//   - [SKView.SetDelegate]
//
// # Configuring Performance Related Toggles
//
//   - [SKView.IgnoresSiblingOrder]: A Boolean value that indicates whether parent-child and sibling relationships affect the rendering order of nodes in the scene.
//   - [SKView.SetIgnoresSiblingOrder]
//   - [SKView.ShouldCullNonVisibleNodes]: A Boolean value that indicates whether the view automatically culls non-visible nodes from the rendering tree.
//   - [SKView.SetShouldCullNonVisibleNodes]
//   - [SKView.AllowsTransparency]: A Boolean property that indicates whether the view is rendered using transparency.
//   - [SKView.SetAllowsTransparency]
//   - [SKView.IsAsynchronous]: A Boolean value that indicates whether the content is rendered asynchronously.
//   - [SKView.SetAsynchronous]
//
// # Enabling Visual Statistics for Debugging
//
//   - [SKView.ShowsFPS]: A Boolean value that indicates whether the view displays a frame rate indicator.
//   - [SKView.SetShowsFPS]
//   - [SKView.ShowsNodeCount]: A Boolean value that indicates whether the view displays an overlay that shows physics bodies that are visible in the scene.
//   - [SKView.SetShowsNodeCount]
//   - [SKView.ShowsDrawCount]: A Boolean value that indicates whether the view displays the number of drawing passes it needed to render the view.
//   - [SKView.SetShowsDrawCount]
//   - [SKView.ShowsQuadCount]: A Boolean value that indicates whether the view displays the number of rectangles used to render the scene.
//   - [SKView.SetShowsQuadCount]
//   - [SKView.ShowsPhysics]: A Boolean value that indicates whether the view displays physics-related debugging information.
//   - [SKView.SetShowsPhysics]
//   - [SKView.ShowsFields]: A Boolean value that indicates whether the view displays information about physics fields in the scene.
//   - [SKView.SetShowsFields]
//
// # Converting Between View and Scene Coordinates
//
//   - [SKView.ConvertPointFromScene]: Converts a point from scene coordinates to view coordinates.
//   - [SKView.ConvertPointToScene]: Converts a point from view coordinates to scene coordinates.
//
// # Snapshotting Nodes to a Texture
//
//   - [SKView.TextureFromNodeCrop]: Renders a portion of a node’s contents and returns the rendered image as a texture.
//   - [SKView.TextureFromNode]: Renders the contents of a node tree and returns the rendered image as a texture.
//
// # Instance Properties
//
//   - [SKView.DisableDepthStencilBuffer]
//   - [SKView.SetDisableDepthStencilBuffer]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKView
type SKView struct {
	appkit.NSView
}

// SKViewFromID constructs a [SKView] from an objc.ID.
//
// A view subclass that renders a SpriteKit scene.
func SKViewFromID(id objc.ID) SKView {
	return SKView{NSView: appkit.NSViewFromID(id)}
}

// NOTE: SKView adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SKView] class.
//
// # Displaying a Scene
//
//   - [ISKView.Scene]: The scene currently presented by this view.
//   - [ISKView.PresentScene]: Presents a scene.
//   - [ISKView.PresentSceneTransition]: Transitions from the current scene to a new scene.
//
// # Controlling the Timing of a Scene’s Rendering
//
//   - [ISKView.IsPaused]: A Boolean value that indicates whether the view’s scene animations are paused.
//   - [ISKView.SetPaused]
//   - [ISKView.PreferredFramesPerSecond]: The animation frame rate that the view uses to render its scene.
//   - [ISKView.SetPreferredFramesPerSecond]
//   - [ISKView.Delegate]: A delegate that allows dynamic control of the view’s render rate.
//   - [ISKView.SetDelegate]
//
// # Configuring Performance Related Toggles
//
//   - [ISKView.IgnoresSiblingOrder]: A Boolean value that indicates whether parent-child and sibling relationships affect the rendering order of nodes in the scene.
//   - [ISKView.SetIgnoresSiblingOrder]
//   - [ISKView.ShouldCullNonVisibleNodes]: A Boolean value that indicates whether the view automatically culls non-visible nodes from the rendering tree.
//   - [ISKView.SetShouldCullNonVisibleNodes]
//   - [ISKView.AllowsTransparency]: A Boolean property that indicates whether the view is rendered using transparency.
//   - [ISKView.SetAllowsTransparency]
//   - [ISKView.IsAsynchronous]: A Boolean value that indicates whether the content is rendered asynchronously.
//   - [ISKView.SetAsynchronous]
//
// # Enabling Visual Statistics for Debugging
//
//   - [ISKView.ShowsFPS]: A Boolean value that indicates whether the view displays a frame rate indicator.
//   - [ISKView.SetShowsFPS]
//   - [ISKView.ShowsNodeCount]: A Boolean value that indicates whether the view displays an overlay that shows physics bodies that are visible in the scene.
//   - [ISKView.SetShowsNodeCount]
//   - [ISKView.ShowsDrawCount]: A Boolean value that indicates whether the view displays the number of drawing passes it needed to render the view.
//   - [ISKView.SetShowsDrawCount]
//   - [ISKView.ShowsQuadCount]: A Boolean value that indicates whether the view displays the number of rectangles used to render the scene.
//   - [ISKView.SetShowsQuadCount]
//   - [ISKView.ShowsPhysics]: A Boolean value that indicates whether the view displays physics-related debugging information.
//   - [ISKView.SetShowsPhysics]
//   - [ISKView.ShowsFields]: A Boolean value that indicates whether the view displays information about physics fields in the scene.
//   - [ISKView.SetShowsFields]
//
// # Converting Between View and Scene Coordinates
//
//   - [ISKView.ConvertPointFromScene]: Converts a point from scene coordinates to view coordinates.
//   - [ISKView.ConvertPointToScene]: Converts a point from view coordinates to scene coordinates.
//
// # Snapshotting Nodes to a Texture
//
//   - [ISKView.TextureFromNodeCrop]: Renders a portion of a node’s contents and returns the rendered image as a texture.
//   - [ISKView.TextureFromNode]: Renders the contents of a node tree and returns the rendered image as a texture.
//
// # Instance Properties
//
//   - [ISKView.DisableDepthStencilBuffer]
//   - [ISKView.SetDisableDepthStencilBuffer]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKView
type ISKView interface {
	appkit.INSView

	// Topic: Displaying a Scene

	// The scene currently presented by this view.
	Scene() ISKScene
	// Presents a scene.
	PresentScene(scene ISKScene)
	// Transitions from the current scene to a new scene.
	PresentSceneTransition(scene ISKScene, transition ISKTransition)

	// Topic: Controlling the Timing of a Scene’s Rendering

	// A Boolean value that indicates whether the view’s scene animations are paused.
	IsPaused() bool
	SetPaused(value bool)
	// The animation frame rate that the view uses to render its scene.
	PreferredFramesPerSecond() int
	SetPreferredFramesPerSecond(value int)
	// A delegate that allows dynamic control of the view’s render rate.
	Delegate() objectivec.Object
	SetDelegate(value objectivec.Object)

	// Topic: Configuring Performance Related Toggles

	// A Boolean value that indicates whether parent-child and sibling relationships affect the rendering order of nodes in the scene.
	IgnoresSiblingOrder() bool
	SetIgnoresSiblingOrder(value bool)
	// A Boolean value that indicates whether the view automatically culls non-visible nodes from the rendering tree.
	ShouldCullNonVisibleNodes() bool
	SetShouldCullNonVisibleNodes(value bool)
	// A Boolean property that indicates whether the view is rendered using transparency.
	AllowsTransparency() bool
	SetAllowsTransparency(value bool)
	// A Boolean value that indicates whether the content is rendered asynchronously.
	IsAsynchronous() bool
	SetAsynchronous(value bool)

	// Topic: Enabling Visual Statistics for Debugging

	// A Boolean value that indicates whether the view displays a frame rate indicator.
	ShowsFPS() bool
	SetShowsFPS(value bool)
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

	// Topic: Converting Between View and Scene Coordinates

	// Converts a point from scene coordinates to view coordinates.
	ConvertPointFromScene(point corefoundation.CGPoint, scene ISKScene) corefoundation.CGPoint
	// Converts a point from view coordinates to scene coordinates.
	ConvertPointToScene(point corefoundation.CGPoint, scene ISKScene) corefoundation.CGPoint

	// Topic: Snapshotting Nodes to a Texture

	// Renders a portion of a node’s contents and returns the rendered image as a texture.
	TextureFromNodeCrop(node ISKNode, crop corefoundation.CGRect) ISKTexture
	// Renders the contents of a node tree and returns the rendered image as a texture.
	TextureFromNode(node ISKNode) ISKTexture

	// Topic: Instance Properties

	DisableDepthStencilBuffer() bool
	SetDisableDepthStencilBuffer(value bool)
}

// Init initializes the instance.
func (v SKView) Init() SKView {
	rv := objc.Send[SKView](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v SKView) Autorelease() SKView {
	rv := objc.Send[SKView](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKView creates a new SKView instance.
func NewSKView() SKView {
	class := getSKViewClass()
	rv := objc.Send[SKView](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Presents a scene.
//
// scene: The scene to present.
//
// # Discussion
//
// The new scene immediately replaces the current scene, if one exists.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKView/presentScene(_:)
func (v SKView) PresentScene(scene ISKScene) {
	objc.Send[objc.ID](v.ID, objc.Sel("presentScene:"), scene)
}

// Transitions from the current scene to a new scene.
//
// scene: The scene to present.
//
// transition: A transition used to animate between the two scenes.
//
// # Discussion
//
// If there is currently a scene presented by the view, the view’s
// [SKView.Scene] property is updated immediately, the transition is executed
// to swap between the scenes. Otherwise, the new scene is presented
// immediately and the transition property is ignored.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKView/presentScene(_:transition:)
func (v SKView) PresentSceneTransition(scene ISKScene, transition ISKTransition) {
	objc.Send[objc.ID](v.ID, objc.Sel("presentScene:transition:"), scene, transition)
}

// Converts a point from scene coordinates to view coordinates.
//
// point: A point in scene coordinates.
//
// scene: A scene.
//
// # Return Value
//
// The same point in the view’s coordinate system.
//
// # Discussion
//
// This method performs the coordinate conversion as if the scene is presented
// inside the view.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKView/convert(_:from:)
func (v SKView) ConvertPointFromScene(point corefoundation.CGPoint, scene ISKScene) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](v.ID, objc.Sel("convertPoint:fromScene:"), point, scene)
	return corefoundation.CGPoint(rv)
}

// Converts a point from view coordinates to scene coordinates.
//
// point: A point in view coordinates.
//
// scene: A scene.
//
// # Return Value
//
// The same point in the scene’s coordinate system.
//
// # Discussion
//
// This method performs the coordinate conversion as if the scene is presented
// inside the view.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKView/convert(_:to:)
func (v SKView) ConvertPointToScene(point corefoundation.CGPoint, scene ISKScene) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](v.ID, objc.Sel("convertPoint:toScene:"), point, scene)
	return corefoundation.CGPoint(rv)
}

// Renders a portion of a node’s contents and returns the rendered image as
// a texture.
//
// node: The node object that is the root node of the tree you want to render to the
// texture.
//
// crop: A rectangle in the node’s coordinate system that describes the area to be
// rendered.
//
// # Return Value
//
// A SpriteKit texture that holds the rendered image.
//
// # Discussion
//
// The node being rendered does not need to appear in the view’s presented
// scene. The new texture is created with a size equal to the size of the
// `crop` rectangle. If the node is not a scene node, it is rendered with a
// clear background color (`[“SKColor` `clear]`).
//
// See: https://developer.apple.com/documentation/SpriteKit/SKView/texture(from:crop:)
func (v SKView) TextureFromNodeCrop(node ISKNode, crop corefoundation.CGRect) ISKTexture {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("textureFromNode:crop:"), node, crop)
	return SKTextureFromID(rv)
}

// Renders the contents of a node tree and returns the rendered image as a
// texture.
//
// node: The node object that is the root node of the tree you want to render to the
// texture.
//
// # Return Value
//
// A SpriteKit texture that holds the rendered image.
//
// # Discussion
//
// The node being rendered does not need to appear in the view’s presented
// scene. The new texture is created with a size equal to the rectangle
// returned by the node’s [SKNode.CalculateAccumulatedFrame] method. If the
// node is not a scene node, it is rendered with a clear background color
// (`[“SKColor` `clear]`).
//
// See: https://developer.apple.com/documentation/SpriteKit/SKView/texture(from:)
func (v SKView) TextureFromNode(node ISKNode) ISKTexture {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("textureFromNode:"), node)
	return SKTextureFromID(rv)
}

// The scene currently presented by this view.
//
// # Discussion
//
// The default value is `nil`.
//
// You call [SKView.PresentScene] to assign a value to this property.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKView/scene
func (v SKView) Scene() ISKScene {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("scene"))
	return SKSceneFromID(objc.ID(rv))
}

// A Boolean value that indicates whether the view’s scene animations are
// paused.
//
// # Discussion
//
// If the value is true, the scene’s content is fixed onscreen. No actions
// are executed and no physics simulation is performed.
//
// When an application moves from an active to an inactive state,
// [SKView.Paused] is automatically set to true. When an application returns
// to an active state, [SKView.Paused] is automatically set to its previous
// value.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKView/isPaused
func (v SKView) IsPaused() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("isPaused"))
	return rv
}
func (v SKView) SetPaused(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setPaused:"), value)
}

// The animation frame rate that the view uses to render its scene.
//
// # Discussion
//
// When your application sets its preferred frame rate, the view chooses a
// frame rate as close to that as possible based on the capabilities of the
// screen the view is displayed on. The actual frame rate chosen is usually a
// factor of the maximum refresh rate of the screen to provide a consistent
// frame rate. For example, if the maximum refresh rate of the screen is 60
// frames per second, that is also the highest frame rate the view sets as the
// actual frame rate. However, if you ask for a lower frame rate, the view
// might choose 30, 20, or 15 frames per second, or another rate, as the
// actual frame rate.
//
// Your application should choose a frame rate that it can consistently
// maintain.
//
// The default value is 60 frames per second.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKView/preferredFramesPerSecond
func (v SKView) PreferredFramesPerSecond() int {
	rv := objc.Send[int](v.ID, objc.Sel("preferredFramesPerSecond"))
	return rv
}
func (v SKView) SetPreferredFramesPerSecond(value int) {
	objc.Send[struct{}](v.ID, objc.Sel("setPreferredFramesPerSecond:"), value)
}

// A delegate that allows dynamic control of the view’s render rate.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKView/delegate
func (v SKView) Delegate() objectivec.Object {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("delegate"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (v SKView) SetDelegate(value objectivec.Object) {
	objc.Send[struct{}](v.ID, objc.Sel("setDelegate:"), value)
}

// A Boolean value that indicates whether parent-child and sibling
// relationships affect the rendering order of nodes in the scene.
//
// # Discussion
//
// The default value is false, which means that when multiple nodes share the
// same z position, those nodes are sorted and rendered in a deterministic
// order. Parents are rendered before their children, and siblings are
// rendered in array order. When this property is set to true, the position of
// the nodes in the tree is ignored when determining the rendering order. The
// rendering order of nodes at the same z position is arbitrary and may change
// every time a new frame is rendered. When sibling and parent order is
// ignored, SpriteKit applies additional optimizations to improve rendering
// performance. If you need nodes to be rendered in a specific and
// deterministic order, you must set the z position of those nodes.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKView/ignoresSiblingOrder
func (v SKView) IgnoresSiblingOrder() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("ignoresSiblingOrder"))
	return rv
}
func (v SKView) SetIgnoresSiblingOrder(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setIgnoresSiblingOrder:"), value)
}

// A Boolean value that indicates whether the view automatically culls
// non-visible nodes from the rendering tree.
//
// # Discussion
//
// The default value is true, meaning that when the scene is rendered, the
// scene first searches the tree for invisible or offscreen nodes and culls
// them from the list of nodes to be rendered. Then the remaining (visible)
// nodes are processed and rendered. This is normally the desired behavior,
// because Scene Kit avoids expensive processing on nodes that cannot affect
// the final output. However, if your game is already managing the contents of
// the scene’s node tree (for example, by removing nodes from the tree when
// they are offscreen), you can set this to false to disable automatic scene
// culling. Disabling scene culling removes the performance overhead of this
// check, but each invisible or offscreen node present in the node tree
// reduces the performance of the renderer.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKView/shouldCullNonVisibleNodes
func (v SKView) ShouldCullNonVisibleNodes() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("shouldCullNonVisibleNodes"))
	return rv
}
func (v SKView) SetShouldCullNonVisibleNodes(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setShouldCullNonVisibleNodes:"), value)
}

// A Boolean property that indicates whether the view is rendered using
// transparency.
//
// # Discussion
//
// This property tells the drawing system as to how it should treat the view.
// If set to false, the drawing system treats the view as fully opaque, which
// allows the drawing system to optimize some drawing operations and improve
// performance. If set to true, the drawing system composites the view
// normally with other content. The default value of this property is false.
//
// An opaque view is expected to fill its bounds with entirely opaque
// content—that is, the content should have an alpha value of 1.0. If the
// view is opaque and either does not fill its bounds or contains wholly or
// partially transparent content, the results are unpredictable. Always set
// the value of this property to false if the view is fully or partially
// transparent.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKView/allowsTransparency
func (v SKView) AllowsTransparency() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("allowsTransparency"))
	return rv
}
func (v SKView) SetAllowsTransparency(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setAllowsTransparency:"), value)
}

// A Boolean value that indicates whether the content is rendered
// asynchronously.
//
// # Discussion
//
// The default value is true. If the value is false, the contents of this view
// are synchronized with Core Animation updates.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKView/isAsynchronous
func (v SKView) IsAsynchronous() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("isAsynchronous"))
	return rv
}
func (v SKView) SetAsynchronous(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setAsynchronous:"), value)
}

// A Boolean value that indicates whether the view displays a frame rate
// indicator.
//
// # Discussion
//
// The frame rate is a good indicator of the performance of your scene. Avoid
// creating scenes that have widely varying frame rates.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKView/showsFPS
func (v SKView) ShowsFPS() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("showsFPS"))
	return rv
}
func (v SKView) SetShowsFPS(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setShowsFPS:"), value)
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
// [SKView.ShouldCullNonVisibleNodes], there would be less nodes for SpriteKit
// to test every frame whether they’re on screen.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKView/showsNodeCount
func (v SKView) ShowsNodeCount() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("showsNodeCount"))
	return rv
}
func (v SKView) SetShowsNodeCount(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setShowsNodeCount:"), value)
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
// See: https://developer.apple.com/documentation/SpriteKit/SKView/showsDrawCount
func (v SKView) ShowsDrawCount() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("showsDrawCount"))
	return rv
}
func (v SKView) SetShowsDrawCount(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setShowsDrawCount:"), value)
}

// A Boolean value that indicates whether the view displays the number of
// rectangles used to render the scene.
//
// # Discussion
//
// SpriteKit converts the node tree into one or more rendering passes. Each
// rendering pass is rendered using a series of textured rectangles (quads).
// The [SKView.ShowsQuadCount] property allows you to see the total number of
// quads that were used to render the scene’s contents. Use this as another
// piece of data when you profile your game’s performance. In most cases,
// fewer quads is better.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKView/showsQuadCount
func (v SKView) ShowsQuadCount() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("showsQuadCount"))
	return rv
}
func (v SKView) SetShowsQuadCount(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setShowsQuadCount:"), value)
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
// See: https://developer.apple.com/documentation/SpriteKit/SKView/showsPhysics
func (v SKView) ShowsPhysics() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("showsPhysics"))
	return rv
}
func (v SKView) SetShowsPhysics(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setShowsPhysics:"), value)
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
// See: https://developer.apple.com/documentation/SpriteKit/SKView/showsFields
func (v SKView) ShowsFields() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("showsFields"))
	return rv
}
func (v SKView) SetShowsFields(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setShowsFields:"), value)
}

// See: https://developer.apple.com/documentation/SpriteKit/SKView/disableDepthStencilBuffer
func (v SKView) DisableDepthStencilBuffer() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("disableDepthStencilBuffer"))
	return rv
}
func (v SKView) SetDisableDepthStencilBuffer(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setDisableDepthStencilBuffer:"), value)
}
