// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/avfaudio"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [SKScene] class.
var (
	_SKSceneClass     SKSceneClass
	_SKSceneClassOnce sync.Once
)

func getSKSceneClass() SKSceneClass {
	_SKSceneClassOnce.Do(func() {
		_SKSceneClass = SKSceneClass{class: objc.GetClass("SKScene")}
	})
	return _SKSceneClass
}

// GetSKSceneClass returns the class object for SKScene.
func GetSKSceneClass() SKSceneClass {
	return getSKSceneClass()
}

type SKSceneClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SKSceneClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SKSceneClass) Alloc() SKScene {
	rv := objc.Send[SKScene](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// An object that organizes all of the active SpriteKit content.
//
// # Overview
//
// An [SKScene] object represents a scene of content in SpriteKit. A scene is
// the root node in a tree of SpriteKit nodes ([SKNode]). These nodes provide
// content that the scene animates and renders for display. To display a
// scene, you present it from an [SKView], [SKRenderer], or
// [WKInterfaceSKScene].
//
// [SKScene] is a subclass of [SKEffectNode] and enables certain effects to
// apply to the entire scene. Though applying effects to an entire scene can
// be an expensive operation, creativity, and ingenuity may help you find some
// interesting ways to use effects.
//
// # Creating a Scene Programmatically
//
//   - [SKScene.InitWithSize]: Initializes a new scene object.
//   - [SKScene.Size]: The dimensions of the scene, in points.
//   - [SKScene.SetSize]
//
// # Stretching Content to Fit the View
//
//   - [SKScene.ScaleMode]: A setting that defines how the scene is mapped to the view that presents it.
//   - [SKScene.SetScaleMode]
//
// # Configuring the Viewport
//
//   - [SKScene.Camera]: The camera node in the scene that determines what part of the scene’s coordinate space is visible in the view.
//   - [SKScene.SetCamera]
//   - [SKScene.AnchorPoint]: The point in the view’s frame that corresponds to the scene’s origin.
//   - [SKScene.SetAnchorPoint]
//
// # Responding to Loading and Resizing Events
//
//   - [SKScene.SceneDidLoad]: Tells you when the scene is presented.
//   - [SKScene.DidChangeSize]: Tells you when the scene’s size has changed.
//   - [SKScene.WillMoveFromView]: Tells you when the scene is about to be removed from a view.
//   - [SKScene.DidMoveToView]: Tells you when the scene is presented by a view.
//
// # Responding to Frame-Cycle Events
//
//   - [SKScene.Update]: Tells your app to perform any app-specific logic to update your scene.
//   - [SKScene.DidEvaluateActions]: Tells your app to peform any necessary logic after scene actions are evaluated.
//   - [SKScene.DidSimulatePhysics]: Tells your app to peform any necessary logic after physics simulations are performed.
//   - [SKScene.DidApplyConstraints]: Tells your app to peform any necessary logic after constraints are applied.
//   - [SKScene.DidFinishUpdate]: Tells your app to peform any necessary logic after the scene has finished all of the steps required to process animations.
//
// # Configuring a Delegate
//
//   - [SKScene.Delegate]: A delegate to be called during the animation loop.
//   - [SKScene.SetDelegate]
//
// # Setting the Background Appearance
//
//   - [SKScene.View]: The view that is currently presenting the scene.
//   - [SKScene.BackgroundColor]: The background color of the scene.
//   - [SKScene.SetBackgroundColor]
//
// # Configuring Physics Properties
//
//   - [SKScene.PhysicsWorld]: The physics simulation associated with the scene.
//
// # Adding Positional Audio
//
//   - [SKScene.Listener]: A node used to determine the position of the listener for positional audio in the scene.
//   - [SKScene.SetListener]
//   - [SKScene.AudioEngine]: The AVFoundation audio engine used to play audio from audio nodes contained in the scene.
//
// # Converting Between Coordinate Systems
//
//   - [SKScene.ConvertPointFromView]: Converts a point from view coordinates to scene coordinates.
//   - [SKScene.ConvertPointToView]: Converts a point from scene coordinates to view coordinates.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKScene
//
// [WKInterfaceSKScene]: https://developer.apple.com/documentation/WatchKit/WKInterfaceSKScene
type SKScene struct {
	SKEffectNode
}

// SKSceneFromID constructs a [SKScene] from an objc.ID.
//
// An object that organizes all of the active SpriteKit content.
func SKSceneFromID(id objc.ID) SKScene {
	return SKScene{SKEffectNode: SKEffectNodeFromID(id)}
}

// NOTE: SKScene adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SKScene] class.
//
// # Creating a Scene Programmatically
//
//   - [ISKScene.InitWithSize]: Initializes a new scene object.
//   - [ISKScene.Size]: The dimensions of the scene, in points.
//   - [ISKScene.SetSize]
//
// # Stretching Content to Fit the View
//
//   - [ISKScene.ScaleMode]: A setting that defines how the scene is mapped to the view that presents it.
//   - [ISKScene.SetScaleMode]
//
// # Configuring the Viewport
//
//   - [ISKScene.Camera]: The camera node in the scene that determines what part of the scene’s coordinate space is visible in the view.
//   - [ISKScene.SetCamera]
//   - [ISKScene.AnchorPoint]: The point in the view’s frame that corresponds to the scene’s origin.
//   - [ISKScene.SetAnchorPoint]
//
// # Responding to Loading and Resizing Events
//
//   - [ISKScene.SceneDidLoad]: Tells you when the scene is presented.
//   - [ISKScene.DidChangeSize]: Tells you when the scene’s size has changed.
//   - [ISKScene.WillMoveFromView]: Tells you when the scene is about to be removed from a view.
//   - [ISKScene.DidMoveToView]: Tells you when the scene is presented by a view.
//
// # Responding to Frame-Cycle Events
//
//   - [ISKScene.Update]: Tells your app to perform any app-specific logic to update your scene.
//   - [ISKScene.DidEvaluateActions]: Tells your app to peform any necessary logic after scene actions are evaluated.
//   - [ISKScene.DidSimulatePhysics]: Tells your app to peform any necessary logic after physics simulations are performed.
//   - [ISKScene.DidApplyConstraints]: Tells your app to peform any necessary logic after constraints are applied.
//   - [ISKScene.DidFinishUpdate]: Tells your app to peform any necessary logic after the scene has finished all of the steps required to process animations.
//
// # Configuring a Delegate
//
//   - [ISKScene.Delegate]: A delegate to be called during the animation loop.
//   - [ISKScene.SetDelegate]
//
// # Setting the Background Appearance
//
//   - [ISKScene.View]: The view that is currently presenting the scene.
//   - [ISKScene.BackgroundColor]: The background color of the scene.
//   - [ISKScene.SetBackgroundColor]
//
// # Configuring Physics Properties
//
//   - [ISKScene.PhysicsWorld]: The physics simulation associated with the scene.
//
// # Adding Positional Audio
//
//   - [ISKScene.Listener]: A node used to determine the position of the listener for positional audio in the scene.
//   - [ISKScene.SetListener]
//   - [ISKScene.AudioEngine]: The AVFoundation audio engine used to play audio from audio nodes contained in the scene.
//
// # Converting Between Coordinate Systems
//
//   - [ISKScene.ConvertPointFromView]: Converts a point from view coordinates to scene coordinates.
//   - [ISKScene.ConvertPointToView]: Converts a point from scene coordinates to view coordinates.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKScene
type ISKScene interface {
	ISKEffectNode

	// Topic: Creating a Scene Programmatically

	// Initializes a new scene object.
	InitWithSize(size corefoundation.CGSize) SKScene
	// The dimensions of the scene, in points.
	Size() corefoundation.CGSize
	SetSize(value corefoundation.CGSize)

	// Topic: Stretching Content to Fit the View

	// A setting that defines how the scene is mapped to the view that presents it.
	ScaleMode() SKSceneScaleMode
	SetScaleMode(value SKSceneScaleMode)

	// Topic: Configuring the Viewport

	// The camera node in the scene that determines what part of the scene’s coordinate space is visible in the view.
	Camera() ISKCameraNode
	SetCamera(value ISKCameraNode)
	// The point in the view’s frame that corresponds to the scene’s origin.
	AnchorPoint() corefoundation.CGPoint
	SetAnchorPoint(value corefoundation.CGPoint)

	// Topic: Responding to Loading and Resizing Events

	// Tells you when the scene is presented.
	SceneDidLoad()
	// Tells you when the scene’s size has changed.
	DidChangeSize(oldSize corefoundation.CGSize)
	// Tells you when the scene is about to be removed from a view.
	WillMoveFromView(view ISKView)
	// Tells you when the scene is presented by a view.
	DidMoveToView(view ISKView)

	// Topic: Responding to Frame-Cycle Events

	// Tells your app to perform any app-specific logic to update your scene.
	Update(currentTime foundation.NSTimeInterval)
	// Tells your app to peform any necessary logic after scene actions are evaluated.
	DidEvaluateActions()
	// Tells your app to peform any necessary logic after physics simulations are performed.
	DidSimulatePhysics()
	// Tells your app to peform any necessary logic after constraints are applied.
	DidApplyConstraints()
	// Tells your app to peform any necessary logic after the scene has finished all of the steps required to process animations.
	DidFinishUpdate()

	// Topic: Configuring a Delegate

	// A delegate to be called during the animation loop.
	Delegate() SKSceneDelegate
	SetDelegate(value SKSceneDelegate)

	// Topic: Setting the Background Appearance

	// The view that is currently presenting the scene.
	View() ISKView
	// The background color of the scene.
	BackgroundColor() appkit.NSColor
	SetBackgroundColor(value appkit.NSColor)

	// Topic: Configuring Physics Properties

	// The physics simulation associated with the scene.
	PhysicsWorld() ISKPhysicsWorld

	// Topic: Adding Positional Audio

	// A node used to determine the position of the listener for positional audio in the scene.
	Listener() ISKNode
	SetListener(value ISKNode)
	// The AVFoundation audio engine used to play audio from audio nodes contained in the scene.
	AudioEngine() avfaudio.AVAudioEngine

	// Topic: Converting Between Coordinate Systems

	// Converts a point from view coordinates to scene coordinates.
	ConvertPointFromView(point corefoundation.CGPoint) corefoundation.CGPoint
	// Converts a point from scene coordinates to view coordinates.
	ConvertPointToView(point corefoundation.CGPoint) corefoundation.CGPoint
}

// Init initializes the instance.
func (s SKScene) Init() SKScene {
	rv := objc.Send[SKScene](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SKScene) Autorelease() SKScene {
	rv := objc.Send[SKScene](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKScene creates a new SKScene instance.
func NewSKScene() SKScene {
	class := getSKSceneClass()
	rv := objc.Send[SKScene](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new node by loading an archive file from the game’s main
// bundle.
//
// filename: The name of the file, without a file extension. The file must be in the
// app’s main bundle and have a `XCUIElementTypeSks` filename extension.
//
// # Return Value
//
// The unarchived node object.
//
// # Discussion
//
// If you call this method on a subclass of the [SKScene] class and the object
// in the archive is an [SKScene] object, the returned object is initialized
// as if it is a member of the subclass. You use this behavior to create scene
// layouts in the Xcode Editor and provide custom behaviors in your subclass.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/init(fileNamed:)
func NewSceneNodeWithFileNamed(filename string) SKScene {
	rv := objc.Send[objc.ID](objc.ID(getSKSceneClass().class), objc.Sel("nodeWithFileNamed:"), objc.String(filename))
	return SKSceneFromID(rv)
}

// See: https://developer.apple.com/documentation/SpriteKit/SKNode/init(fileNamed:securelyWithClasses:)
func NewSceneNodeWithFileNamedSecurelyWithClassesAndError(filename string, classes foundation.INSSet) (SKScene, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getSKSceneClass().class), objc.Sel("nodeWithFileNamed:securelyWithClasses:andError:"), objc.String(filename), classes, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return SKScene{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return SKScene{}, objc.ErrInitFailed
	}
	return SKSceneFromID(rv), nil
}

// Called when a node is initialized from an .sks file.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/init(coder:)
func NewSceneWithCoder(aDecoder foundation.INSCoder) SKScene {
	instance := getSKSceneClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return SKSceneFromID(rv)
}

// Initializes a new scene object.
//
// size: The size of the scene in points.
//
// # Return Value
//
// A newly initialized scene object.
//
// # Discussion
//
// This is the class’s designated initializer method.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKScene/init(size:)
func NewSceneWithSize(size corefoundation.CGSize) SKScene {
	instance := getSKSceneClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSize:"), size)
	return SKSceneFromID(rv)
}

// Initializes a new scene object.
//
// size: The size of the scene in points.
//
// # Return Value
//
// A newly initialized scene object.
//
// # Discussion
//
// This is the class’s designated initializer method.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKScene/init(size:)
func (s SKScene) InitWithSize(size corefoundation.CGSize) SKScene {
	rv := objc.Send[SKScene](s.ID, objc.Sel("initWithSize:"), size)
	return rv
}

// Tells you when the scene is presented.
//
// # Discussion
//
// This method is intended to be overridden in a subclass. It is the preferred
// location to peform custom setup after the scene has been initialized or
// decoded.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKScene/sceneDidLoad()
func (s SKScene) SceneDidLoad() {
	objc.Send[objc.ID](s.ID, objc.Sel("sceneDidLoad"))
}

// Tells you when the scene’s size has changed.
//
// oldSize: The old size of the scene, in points.
//
// # Discussion
//
// This method is intended to be overridden in a subclass. Typically, you use
// this method to adjust the positions of nodes in the scene.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKScene/didChangeSize(_:)
func (s SKScene) DidChangeSize(oldSize corefoundation.CGSize) {
	objc.Send[objc.ID](s.ID, objc.Sel("didChangeSize:"), oldSize)
}

// Tells you when the scene is about to be removed from a view.
//
// view: The view that is presenting the scene.
//
// # Discussion
//
// This method is intended to be overridden in a subclass. You can use this
// method to implement any custom behavior for your scene when it is about to
// be removed from the view.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKScene/willMove(from:)
func (s SKScene) WillMoveFromView(view ISKView) {
	objc.Send[objc.ID](s.ID, objc.Sel("willMoveFromView:"), view)
}

// Tells you when the scene is presented by a view.
//
// view: The view that is presenting the scene.
//
// # Discussion
//
// This method is intended to be overridden in a subclass. You can use this
// method to implement any custom behavior for your scene when it is about to
// be presented by a view. For example, you might use this method to create
// the scene’s contents.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKScene/didMove(to:)
func (s SKScene) DidMoveToView(view ISKView) {
	objc.Send[objc.ID](s.ID, objc.Sel("didMoveToView:"), view)
}

// Tells your app to perform any app-specific logic to update your scene.
//
// currentTime: The current system time.
//
// # Discussion
//
// Do not call this method directly; it is called by the system exactly once
// per frame, so long as the scene is presented in a view and is not paused.
// This is the first method called when animating the scene, before any
// actions are evaluated and before any physics are simulated.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKScene/update(_:)
func (s SKScene) Update(currentTime foundation.NSTimeInterval) {
	objc.Send[objc.ID](s.ID, objc.Sel("update:"), currentTime)
}

// Tells your app to peform any necessary logic after scene actions are
// evaluated.
//
// # Discussion
//
// Do not call this method directly; it is called by the system exactly once
// per frame, so long as the scene is presented in a view and is not paused.
// It is called after any actions have been evaluated by nodes in the scene
// but before any physics are simulated.
//
// Any additional actions applied are not evaluated until the next update.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKScene/didEvaluateActions()
func (s SKScene) DidEvaluateActions() {
	objc.Send[objc.ID](s.ID, objc.Sel("didEvaluateActions"))
}

// Tells your app to peform any necessary logic after physics simulations are
// performed.
//
// # Discussion
//
// Do not call this method directly; it is called by the system exactly once
// per frame, so long as the scene is presented in a view and is not paused.
// It is called after physics has been simulated in the scene.
//
// Any additional actions applied are not evaluated until the next update.
//
// Any changes to physics bodies are not simulated until the next update.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKScene/didSimulatePhysics()
func (s SKScene) DidSimulatePhysics() {
	objc.Send[objc.ID](s.ID, objc.Sel("didSimulatePhysics"))
}

// Tells your app to peform any necessary logic after constraints are applied.
//
// # Discussion
//
// Do not call this method directly; it is called exactly once per frame, so
// long as the scene is presented in a view and is not paused. By default,
// this method does nothing. Your scene subclass should override this method
// and perform any necessary updates to the scene.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKScene/didApplyConstraints()
func (s SKScene) DidApplyConstraints() {
	objc.Send[objc.ID](s.ID, objc.Sel("didApplyConstraints"))
}

// Tells your app to peform any necessary logic after the scene has finished
// all of the steps required to process animations.
//
// # Discussion
//
// Do not call this method directly; it is called by the system exactly once
// per frame, so long as the scene is presented in a view and is not paused.
// It is called after all update logic has been completed and before the scene
// is rendered.
//
// Any additional actions applied are not evaluated until the next update.
//
// Any changes to physics bodies are not simulated until the next update.
//
// Any changes to constraints will not be applied until the next update.
//
// No further update logic will be applied to the scene after this call. Any
// values set on nodes here will be used when the scene is rendered for the
// current frame.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKScene/didFinishUpdate()
func (s SKScene) DidFinishUpdate() {
	objc.Send[objc.ID](s.ID, objc.Sel("didFinishUpdate"))
}

// Converts a point from view coordinates to scene coordinates.
//
// point: A point in view coordinates.
//
// # Return Value
//
// The same point in the scene’s coordinate system.
//
// # Discussion
//
// The scene must be presented in a view before calling this method.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKScene/convertPoint(fromView:)
func (s SKScene) ConvertPointFromView(point corefoundation.CGPoint) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](s.ID, objc.Sel("convertPointFromView:"), point)
	return corefoundation.CGPoint(rv)
}

// Converts a point from scene coordinates to view coordinates.
//
// point: A point in scene coordinates.
//
// # Return Value
//
// The same point in the view’s coordinate system.
//
// # Discussion
//
// The scene must be presented in a view before calling this method.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKScene/convertPoint(toView:)
func (s SKScene) ConvertPointToView(point corefoundation.CGPoint) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](s.ID, objc.Sel("convertPointToView:"), point)
	return corefoundation.CGPoint(rv)
}

// The dimensions of the scene, in points.
//
// # Discussion
//
// When a scene is first initialized, its size property is configured by the
// designated initializer. The size of the scene specifies the size of the
// visible portion of the scene in points. This is only used to specify the
// visible portion of the scene. Nodes in the tree can be positioned outside
// of this area; those nodes are still processed by the scene, but are ignored
// by the renderer.
//
// When a scene is presented, the [SKScene.Size] and [SKScene.AnchorPoint]
// properties determine the portion of the scene’s coordinate space that is
// visible in the view.
//
// If you set the [SKScene.Size] property to a new value, the scene’s
// [SKScene.DidChangeSize] method is called. This property can also change if
// the [SKScene.ScaleMode] property is set to [SKSceneScaleMode.resizeFill]
// and the presenting view is resized. After the scene’s size changes,
// future updates are rendered immediately at the new size.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKScene/size
//
// [SKSceneScaleMode.resizeFill]: https://developer.apple.com/documentation/SpriteKit/SKSceneScaleMode/resizeFill
func (s SKScene) Size() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](s.ID, objc.Sel("size"))
	return corefoundation.CGSize(rv)
}
func (s SKScene) SetSize(value corefoundation.CGSize) {
	objc.Send[struct{}](s.ID, objc.Sel("setSize:"), value)
}

// A setting that defines how the scene is mapped to the view that presents
// it.
//
// # Discussion
//
// It is possible for a scene’s size to differ from the size of the view it
// is presented in. The scale mode determines how the visible portion of the
// scene is mapped to the view. The possible values are listed in
// [SKSceneScaleMode]. The default value is [SKSceneScaleMode.fill].
//
// See: https://developer.apple.com/documentation/SpriteKit/SKScene/scaleMode
//
// [SKSceneScaleMode.fill]: https://developer.apple.com/documentation/SpriteKit/SKSceneScaleMode/fill
// [SKSceneScaleMode]: https://developer.apple.com/documentation/SpriteKit/SKSceneScaleMode
func (s SKScene) ScaleMode() SKSceneScaleMode {
	rv := objc.Send[SKSceneScaleMode](s.ID, objc.Sel("scaleMode"))
	return SKSceneScaleMode(rv)
}
func (s SKScene) SetScaleMode(value SKSceneScaleMode) {
	objc.Send[struct{}](s.ID, objc.Sel("setScaleMode:"), value)
}

// The camera node in the scene that determines what part of the scene’s
// coordinate space is visible in the view.
//
// # Discussion
//
// The default value of this property is `nil`, which means that the scene’s
// [SKScene.AnchorPoint] and [SKScene.Size] properties determine what portion
// of the scene is visible. If set to point to a camera node contained in the
// scene, the [SKScene.AnchorPoint] property is ignored and the scene is
// rendered using the camera node’s properties instead.
//
// A camera must be added as a child of the scene for it to render that scene.
//
// Listing 1 shows, in Swift, how to add a camera to an [SKScene] named
// `scene`. The camera is positioned in the center of the scene which gives
// the same result as rendering a camera-less scene with an
// [SKScene.AnchorPoint] of [zero].
//
// Listing 1. Adding a camera to a scene
//
// For more information, see [SKCameraNode].
//
// See: https://developer.apple.com/documentation/SpriteKit/SKScene/camera
//
// [zero]: https://developer.apple.com/documentation/CoreFoundation/CGPoint/zero
func (s SKScene) Camera() ISKCameraNode {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("camera"))
	return SKCameraNodeFromID(objc.ID(rv))
}
func (s SKScene) SetCamera(value ISKCameraNode) {
	objc.Send[struct{}](s.ID, objc.Sel("setCamera:"), value)
}

// The point in the view’s frame that corresponds to the scene’s origin.
//
// # Discussion
//
// When a scene is presented and a camera node has not been specified, the
// [SKScene.Size] and [SKScene.AnchorPoint] properties determine which part of
// the scene’s coordinate space is visible in the view.
//
// You specify the value using the unit coordinate space. The default value is
// `(0,0)`, which corresponds to the lower-left corner of the view’s frame
// rectangle.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKScene/anchorPoint
func (s SKScene) AnchorPoint() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](s.ID, objc.Sel("anchorPoint"))
	return corefoundation.CGPoint(rv)
}
func (s SKScene) SetAnchorPoint(value corefoundation.CGPoint) {
	objc.Send[struct{}](s.ID, objc.Sel("setAnchorPoint:"), value)
}

// A delegate to be called during the animation loop.
//
// # Discussion
//
// When a delegate is present, when any of the animation loop methods steps
// are executed, your delegate is called. Typically, you use a delegate when
// you do not want to implement a scene subclass or if you want to dynamically
// change the scene behavior at runtime.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKScene/delegate
func (s SKScene) Delegate() SKSceneDelegate {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("delegate"))
	return SKSceneDelegateObjectFromID(rv)
}
func (s SKScene) SetDelegate(value SKSceneDelegate) {
	objc.Send[struct{}](s.ID, objc.Sel("setDelegate:"), value)
}

// The view that is currently presenting the scene.
//
// # Discussion
//
// To present a scene, you call the [SKView.PresentScene] method or
// [SKView.PresentSceneTransition] method on the [SKView] class. If the scene
// is not currently presented, this property holds `nil`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKScene/view
func (s SKScene) View() ISKView {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("view"))
	return SKViewFromID(objc.ID(rv))
}

// The background color of the scene.
//
// # Discussion
//
// The default value is a gray color (RGBA `0.15, 0.15, 0.15, 1.0)`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKScene/backgroundColor
func (s SKScene) BackgroundColor() appkit.NSColor {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("backgroundColor"))
	return appkit.NSColorFromID(objc.ID(rv))
}
func (s SKScene) SetBackgroundColor(value appkit.NSColor) {
	objc.Send[struct{}](s.ID, objc.Sel("setBackgroundColor:"), value)
}

// The physics simulation associated with the scene.
//
// # Discussion
//
// Every scene automatically creates a physics world object to simulate
// physics on nodes in the scene. You use this property to access the
// scene’s global physics properties, such as gravity. To add physics to a
// particular node, see [SKNode.PhysicsBody].
//
// See: https://developer.apple.com/documentation/SpriteKit/SKScene/physicsWorld
func (s SKScene) PhysicsWorld() ISKPhysicsWorld {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("physicsWorld"))
	return SKPhysicsWorldFromID(objc.ID(rv))
}

// A node used to determine the position of the listener for positional audio
// in the scene.
//
// # Discussion
//
// The default value is `nil`, which means that the scene’s origin is used
// as the listener position for audio effects played by [SKAudioNode] objects
// in the scene. If a non-`nil` value is specified, it must be a node in the
// scene.
//
// Typically, you want the [SKScene.Camera] to be the listener so that audio
// nodes which are on screen are louder than off screen ones. In a game, the
// node that defines the player would likely be set as the `listener`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKScene/listener
func (s SKScene) Listener() ISKNode {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("listener"))
	return SKNodeFromID(objc.ID(rv))
}
func (s SKScene) SetListener(value ISKNode) {
	objc.Send[struct{}](s.ID, objc.Sel("setListener:"), value)
}

// The AVFoundation audio engine used to play audio from audio nodes contained
// in the scene.
//
// # Discussion
//
// An audio engine instance is automatically created for you when the scene is
// created. You can use methods and properties on a scene’s audio engine for
// overall control of all of its child audio nodes. The following code shows
// how a scene’s overall volume can be reduced from its default of 1.0 down
// to 0.2 and then paused:
//
// See: https://developer.apple.com/documentation/SpriteKit/SKScene/audioEngine
func (s SKScene) AudioEngine() avfaudio.AVAudioEngine {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("audioEngine"))
	return avfaudio.AVAudioEngineFromID(objc.ID(rv))
}
