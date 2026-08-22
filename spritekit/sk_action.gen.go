// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SKAction] class.
var (
	_SKActionClass     SKActionClass
	_SKActionClassOnce sync.Once
)

func getSKActionClass() SKActionClass {
	_SKActionClassOnce.Do(func() {
		_SKActionClass = SKActionClass{class: objc.GetClass("SKAction")}
	})
	return _SKActionClass
}

// GetSKActionClass returns the class object for SKAction.
func GetSKActionClass() SKActionClass {
	return getSKActionClass()
}

type SKActionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SKActionClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SKActionClass) Alloc() SKAction {
	rv := objc.Send[SKAction](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// An object that is run by a node to change its structure or content.
//
// # Overview
//
// [SKAction] is an animation that is executed by a node in the scene. Actions
// are used to change a node in some way (like move its position over time),
// but you can also use actions to change the scene, like doing a fadeout.
// When the scene processes its nodes, the actions associated with those nodes
// are processed.
//
// # Controlling Action Timing
//
//   - [SKAction.Duration]: The duration required to complete an action.
//   - [SKAction.SetDuration]
//   - [SKAction.TimingMode]: A setting that controls the speed curve of an animation.
//   - [SKAction.SetTimingMode]
//   - [SKAction.TimingFunction]: A block used to customize the timing function.
//   - [SKAction.SetTimingFunction]
//   - [SKAction.Speed]: A speed factor that modifies how fast an action runs.
//   - [SKAction.SetSpeed]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKAction
type SKAction struct {
	objectivec.Object
}

// SKActionFromID constructs a [SKAction] from an objc.ID.
//
// An object that is run by a node to change its structure or content.
func SKActionFromID(id objc.ID) SKAction {
	return SKAction{objectivec.Object{ID: id}}
}

// NOTE: SKAction adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SKAction] class.
//
// # Controlling Action Timing
//
//   - [ISKAction.Duration]: The duration required to complete an action.
//   - [ISKAction.SetDuration]
//   - [ISKAction.TimingMode]: A setting that controls the speed curve of an animation.
//   - [ISKAction.SetTimingMode]
//   - [ISKAction.TimingFunction]: A block used to customize the timing function.
//   - [ISKAction.SetTimingFunction]
//   - [ISKAction.Speed]: A speed factor that modifies how fast an action runs.
//   - [ISKAction.SetSpeed]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKAction
type ISKAction interface {
	objectivec.IObject

	// Topic: Controlling Action Timing

	// The duration required to complete an action.
	Duration() foundation.NSTimeInterval
	SetDuration(value foundation.NSTimeInterval)
	// A setting that controls the speed curve of an animation.
	TimingMode() SKActionTimingMode
	SetTimingMode(value SKActionTimingMode)
	// A block used to customize the timing function.
	TimingFunction() Float32Float32Handler
	SetTimingFunction(value Float32Float32Handler)
	// A speed factor that modifies how fast an action runs.
	Speed() float64
	SetSpeed(value float64)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (a SKAction) Init() SKAction {
	rv := objc.Send[SKAction](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a SKAction) Autorelease() SKAction {
	rv := objc.Send[SKAction](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKAction creates a new SKAction instance.
func NewSKAction() SKAction {
	class := getSKActionClass()
	rv := objc.Send[SKAction](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (a SKAction) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](a.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The duration required to complete an action.
//
// # Discussion
//
// This is the expected duration of an action’s animation. The actual time
// an action takes to complete is modified by the [SKAction.Speed] property of
// the action and the [SKNode.Speed] property of the node on which it
// executes.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKAction/duration
func (a SKAction) Duration() foundation.NSTimeInterval {
	rv := objc.Send[foundation.NSTimeInterval](a.ID, objc.Sel("duration"))
	return foundation.NSTimeInterval(rv)
}
func (a SKAction) SetDuration(value foundation.NSTimeInterval) {
	objc.Send[struct{}](a.ID, objc.Sel("setDuration:"), value)
}

// A setting that controls the speed curve of an animation.
//
// # Discussion
//
// The possible values for this property are listed in [SKActionTimingMode].
// The default value is [SKActionTimingMode.linear].
//
// See: https://developer.apple.com/documentation/SpriteKit/SKAction/timingMode
//
// [SKActionTimingMode.linear]: https://developer.apple.com/documentation/SpriteKit/SKActionTimingMode/linear
// [SKActionTimingMode]: https://developer.apple.com/documentation/SpriteKit/SKActionTimingMode
func (a SKAction) TimingMode() SKActionTimingMode {
	rv := objc.Send[SKActionTimingMode](a.ID, objc.Sel("timingMode"))
	return SKActionTimingMode(rv)
}
func (a SKAction) SetTimingMode(value SKActionTimingMode) {
	objc.Send[struct{}](a.ID, objc.Sel("setTimingMode:"), value)
}

// A block used to customize the timing function.
//
// # Discussion
//
// If a timing function is provided, after the normal timing mode is applied,
// the result is sent to the timing function. The return
// [SKActionTimingFunction] value of the timing function determines the actual
// time used to perform the animation.
//
// The following code shows how you can create a custom timing function using
// [simd_smoothstep(_:_:_:)] interpolation:
//
// If the above code is combined with a vertical linear move action, the path
// taken by a node running this action describes the curve illustrated below:
//
// [media-2759776]
//
// See: https://developer.apple.com/documentation/SpriteKit/SKAction/timingFunction
//
// [simd_smoothstep(_:_:_:)]: https://developer.apple.com/documentation/simd/simd_smoothstep(_:_:_:)-5839l
func (a SKAction) TimingFunction() Float32Float32Handler {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("timingFunction"))
	_ = rv
	return nil
}
func (a SKAction) SetTimingFunction(value Float32Float32Handler) {
	block, cleanup := NewFloat32Float32Block(value)
	defer cleanup()
	objc.Send[struct{}](a.ID, objc.Sel("setTimingFunction:"), block)
}

// A speed factor that modifies how fast an action runs.
//
// # Discussion
//
// The speed factor adjusts how fast an action’s animation runs. For
// example, a speed factor of `2.0` means the animation runs twice as fast.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKAction/speed
func (a SKAction) Speed() float64 {
	rv := objc.Send[float64](a.ID, objc.Sel("speed"))
	return rv
}
func (a SKAction) SetSpeed(value float64) {
	objc.Send[struct{}](a.ID, objc.Sel("setSpeed:"), value)
}
