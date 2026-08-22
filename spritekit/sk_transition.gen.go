// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/coreimage"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SKTransition] class.
var (
	_SKTransitionClass     SKTransitionClass
	_SKTransitionClassOnce sync.Once
)

func getSKTransitionClass() SKTransitionClass {
	_SKTransitionClassOnce.Do(func() {
		_SKTransitionClass = SKTransitionClass{class: objc.GetClass("SKTransition")}
	})
	return _SKTransitionClass
}

// GetSKTransitionClass returns the class object for SKTransition.
func GetSKTransitionClass() SKTransitionClass {
	return getSKTransitionClass()
}

type SKTransitionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SKTransitionClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SKTransitionClass) Alloc() SKTransition {
	rv := objc.Send[SKTransition](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// An object used to perform an animated transition to a new scene.
//
// # Overview
//
// Scenes are the basic building blocks of games. Typically, you design
// self-contained scenes for the parts of your game, and then transition
// between these scenes as necessary. For example, you might create different
// scene classes to represent any or all of the following concepts:
//
// - A loading scene to display while other content is loaded - A main menu
// scene to choose what kind of game the user wants to play - A scene to
// configure the details of the specific kind of game the user chose - A scene
// that provides the gameplay - A scene displayed when gameplay ends
//
// When you present a new scene in a view that is already presenting a scene,
// you have the option of using a transition to animate the change from the
// old scene to the new scene. Using a transition provides continuity so that
// the scene change is not quite so abrupt.
//
// # Pausing
//
//   - [SKTransition.PausesIncomingScene]: A Boolean value that determines whether the incoming scene is paused during the transition.
//   - [SKTransition.SetPausesIncomingScene]
//   - [SKTransition.PausesOutgoingScene]: A Boolean value that determines whether the outgoing scene is paused during the transition.
//   - [SKTransition.SetPausesOutgoingScene]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTransition
type SKTransition struct {
	objectivec.Object
}

// SKTransitionFromID constructs a [SKTransition] from an objc.ID.
//
// An object used to perform an animated transition to a new scene.
func SKTransitionFromID(id objc.ID) SKTransition {
	return SKTransition{objectivec.Object{ID: id}}
}

// NOTE: SKTransition adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SKTransition] class.
//
// # Pausing
//
//   - [ISKTransition.PausesIncomingScene]: A Boolean value that determines whether the incoming scene is paused during the transition.
//   - [ISKTransition.SetPausesIncomingScene]
//   - [ISKTransition.PausesOutgoingScene]: A Boolean value that determines whether the outgoing scene is paused during the transition.
//   - [ISKTransition.SetPausesOutgoingScene]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTransition
type ISKTransition interface {
	objectivec.IObject

	// Topic: Pausing

	// A Boolean value that determines whether the incoming scene is paused during the transition.
	PausesIncomingScene() bool
	SetPausesIncomingScene(value bool)
	// A Boolean value that determines whether the outgoing scene is paused during the transition.
	PausesOutgoingScene() bool
	SetPausesOutgoingScene(value bool)
}

// Init initializes the instance.
func (t SKTransition) Init() SKTransition {
	rv := objc.Send[SKTransition](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t SKTransition) Autorelease() SKTransition {
	rv := objc.Send[SKTransition](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKTransition creates a new SKTransition instance.
func NewSKTransition() SKTransition {
	class := getSKTransitionClass()
	rv := objc.Send[SKTransition](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a transition that uses a Core Image filter to perform the
// transition.
//
// filter: A Core Image filter.
//
// sec: The duration of the transition.
//
// # Return Value
//
// A new transition.
//
// # Discussion
//
// The filter used to perform the transition must a be filter that requires
// only two image parameters (`inputImage`, `inputTargetImage`) and generates
// a single image (`outputImage`). The transition automatically sets the
// filter’s `inputImage`, `inputTargetImage`, and `inputTime` properties.
// You must set up any other filter properties before creating the transition.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTransition/init(ciFilter:duration:)
func NewTransitionWithCIFilterDuration(filter *coreimage.CIFilter, sec foundation.NSTimeInterval) SKTransition {
	rv := objc.Send[objc.ID](objc.ID(getSKTransitionClass().class), objc.Sel("transitionWithCIFilter:duration:"), filter.ID, sec)
	return SKTransitionFromID(rv)
}

// Creates a cross fade transition.
//
// sec: The duration of the transition.
//
// # Return Value
//
// A new transition.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTransition/crossFade(withDuration:)
func (_SKTransitionClass SKTransitionClass) CrossFadeWithDuration(sec foundation.NSTimeInterval) SKTransition {
	rv := objc.Send[objc.ID](objc.ID(_SKTransitionClass.class), objc.Sel("crossFadeWithDuration:"), sec)
	return SKTransitionFromID(rv)
}

// Creates a transition where the new scene appears as a pair of closing
// horizontal doors.
//
// sec: The duration of the transition.
//
// # Return Value
//
// A new transition.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTransition/doorsCloseHorizontal(withDuration:)
func (_SKTransitionClass SKTransitionClass) DoorsCloseHorizontalWithDuration(sec foundation.NSTimeInterval) SKTransition {
	rv := objc.Send[objc.ID](objc.ID(_SKTransitionClass.class), objc.Sel("doorsCloseHorizontalWithDuration:"), sec)
	return SKTransitionFromID(rv)
}

// Creates a transition where the new scene appears as a pair of closing
// vertical doors.
//
// sec: The duration of the transition.
//
// # Return Value
//
// A new transition.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTransition/doorsCloseVertical(withDuration:)
func (_SKTransitionClass SKTransitionClass) DoorsCloseVerticalWithDuration(sec foundation.NSTimeInterval) SKTransition {
	rv := objc.Send[objc.ID](objc.ID(_SKTransitionClass.class), objc.Sel("doorsCloseVerticalWithDuration:"), sec)
	return SKTransitionFromID(rv)
}

// Creates a transition where the new scene appears as a pair of opening
// horizontal doors.
//
// sec: The duration of the transition.
//
// # Return Value
//
// A new transition.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTransition/doorsOpenHorizontal(withDuration:)
func (_SKTransitionClass SKTransitionClass) DoorsOpenHorizontalWithDuration(sec foundation.NSTimeInterval) SKTransition {
	rv := objc.Send[objc.ID](objc.ID(_SKTransitionClass.class), objc.Sel("doorsOpenHorizontalWithDuration:"), sec)
	return SKTransitionFromID(rv)
}

// Creates a transition where the new scene appears as a pair of opening
// vertical doors.
//
// sec: The duration of the transition.
//
// # Return Value
//
// A new transition.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTransition/doorsOpenVertical(withDuration:)
func (_SKTransitionClass SKTransitionClass) DoorsOpenVerticalWithDuration(sec foundation.NSTimeInterval) SKTransition {
	rv := objc.Send[objc.ID](objc.ID(_SKTransitionClass.class), objc.Sel("doorsOpenVerticalWithDuration:"), sec)
	return SKTransitionFromID(rv)
}

// Creates a transition where the previous scene disappears as a pair of
// opening doors.
//
// sec: The duration of the transition.
//
// # Return Value
//
// A new transition.
//
// # Discussion
//
// The new scene starts in the background and moves closer as the doors open.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTransition/doorway(withDuration:)
func (_SKTransitionClass SKTransitionClass) DoorwayWithDuration(sec foundation.NSTimeInterval) SKTransition {
	rv := objc.Send[objc.ID](objc.ID(_SKTransitionClass.class), objc.Sel("doorwayWithDuration:"), sec)
	return SKTransitionFromID(rv)
}

// Creates a transition that first fades to a constant color and then fades to
// the new scene.
//
// color: The color to use as the fade color.
//
// sec: The duration of the transition.
//
// # Return Value
//
// A new transition.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTransition/fade(with:duration:)
func (_SKTransitionClass SKTransitionClass) FadeWithColorDuration(color appkit.NSColor, sec foundation.NSTimeInterval) SKTransition {
	rv := objc.Send[objc.ID](objc.ID(_SKTransitionClass.class), objc.Sel("fadeWithColor:duration:"), color, sec)
	return SKTransitionFromID(rv)
}

// Creates a transition that first fades to black and then fades to the new
// scene.
//
// sec: The duration of the transition.
//
// # Return Value
//
// A new transition.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTransition/fade(withDuration:)
func (_SKTransitionClass SKTransitionClass) FadeWithDuration(sec foundation.NSTimeInterval) SKTransition {
	rv := objc.Send[objc.ID](objc.ID(_SKTransitionClass.class), objc.Sel("fadeWithDuration:"), sec)
	return SKTransitionFromID(rv)
}

// Creates a transition where the two scenes are flipped across a horizontal
// line running through the center of the view.
//
// sec: The duration of the transition.
//
// # Return Value
//
// A new transition.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTransition/flipHorizontal(withDuration:)
func (_SKTransitionClass SKTransitionClass) FlipHorizontalWithDuration(sec foundation.NSTimeInterval) SKTransition {
	rv := objc.Send[objc.ID](objc.ID(_SKTransitionClass.class), objc.Sel("flipHorizontalWithDuration:"), sec)
	return SKTransitionFromID(rv)
}

// Creates a transition where the two scenes are flipped across a vertical
// line running through the center of the view.
//
// sec: The duration of the transition.
//
// # Return Value
//
// A new transition.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTransition/flipVertical(withDuration:)
func (_SKTransitionClass SKTransitionClass) FlipVerticalWithDuration(sec foundation.NSTimeInterval) SKTransition {
	rv := objc.Send[objc.ID](objc.ID(_SKTransitionClass.class), objc.Sel("flipVerticalWithDuration:"), sec)
	return SKTransitionFromID(rv)
}

// Creates a transition where the new scene moves in on top of the old scene.
//
// direction: The direction of the move. Possible values are described in
// [SKTransitionDirection].
//
// sec: The duration of the transition.
//
// # Return Value
//
// A new transition.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTransition/moveIn(with:duration:)
//
// [SKTransitionDirection]: https://developer.apple.com/documentation/SpriteKit/SKTransitionDirection
func (_SKTransitionClass SKTransitionClass) MoveInWithDirectionDuration(direction SKTransitionDirection, sec foundation.NSTimeInterval) SKTransition {
	rv := objc.Send[objc.ID](objc.ID(_SKTransitionClass.class), objc.Sel("moveInWithDirection:duration:"), direction, sec)
	return SKTransitionFromID(rv)
}

// Creates a transition where the new scene moves in, pushing the old scene
// out of the view.
//
// direction: The direction of the push. Possible values are described in
// [SKTransitionDirection].
//
// sec: The duration of the transition.
//
// # Return Value
//
// A new transition.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTransition/push(with:duration:)
//
// [SKTransitionDirection]: https://developer.apple.com/documentation/SpriteKit/SKTransitionDirection
func (_SKTransitionClass SKTransitionClass) PushWithDirectionDuration(direction SKTransitionDirection, sec foundation.NSTimeInterval) SKTransition {
	rv := objc.Send[objc.ID](objc.ID(_SKTransitionClass.class), objc.Sel("pushWithDirection:duration:"), direction, sec)
	return SKTransitionFromID(rv)
}

// Creates a transition where the old scene moves out of the view, revealing
// the new scene underneath it.
//
// direction: The direction of the reveal. Possible values are described in
// [SKTransitionDirection].
//
// sec: The duration of the transition.
//
// # Return Value
//
// A new transition.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTransition/reveal(with:duration:)
//
// [SKTransitionDirection]: https://developer.apple.com/documentation/SpriteKit/SKTransitionDirection
func (_SKTransitionClass SKTransitionClass) RevealWithDirectionDuration(direction SKTransitionDirection, sec foundation.NSTimeInterval) SKTransition {
	rv := objc.Send[objc.ID](objc.ID(_SKTransitionClass.class), objc.Sel("revealWithDirection:duration:"), direction, sec)
	return SKTransitionFromID(rv)
}

// A Boolean value that determines whether the incoming scene is paused during
// the transition.
//
// # Discussion
//
// The default value is true.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTransition/pausesIncomingScene
func (t SKTransition) PausesIncomingScene() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("pausesIncomingScene"))
	return rv
}
func (t SKTransition) SetPausesIncomingScene(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setPausesIncomingScene:"), value)
}

// A Boolean value that determines whether the outgoing scene is paused during
// the transition.
//
// # Discussion
//
// The default value is true.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKTransition/pausesOutgoingScene
func (t SKTransition) PausesOutgoingScene() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("pausesOutgoingScene"))
	return rv
}
func (t SKTransition) SetPausesOutgoingScene(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setPausesOutgoingScene:"), value)
}
