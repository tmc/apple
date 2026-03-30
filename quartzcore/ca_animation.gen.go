// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CAAnimation] class.
var (
	_CAAnimationClass     CAAnimationClass
	_CAAnimationClassOnce sync.Once
)

func getCAAnimationClass() CAAnimationClass {
	_CAAnimationClassOnce.Do(func() {
		_CAAnimationClass = CAAnimationClass{class: objc.GetClass("CAAnimation")}
	})
	return _CAAnimationClass
}

// GetCAAnimationClass returns the class object for CAAnimation.
func GetCAAnimationClass() CAAnimationClass {
	return getCAAnimationClass()
}

type CAAnimationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CAAnimationClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CAAnimationClass) Alloc() CAAnimation {
	rv := objc.Send[CAAnimation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// The abstract superclass for animations in Core Animation.
//
// # Overview
//
// [CAAnimation] provides the basic support for the [CAMediaTiming] and
// [CAAction] protocols. You do not create instance of [CAAnimation]: to
// animate Core Animation layers or SceneKit objects, create instances of the
// concrete subclasses [CABasicAnimation], [CAKeyframeAnimation],
// [CAAnimationGroup], or [CATransition].
//
// # Animating Core Animation Layers
//
// You can animate the contents of your iOS or macOS app’s user interface by
// attaching animations to [CALayer] objects. For more information, see [Core
// Animation Programming Guide].
//
// # Animating Scene Kit Content
//
// In Scene Kit, animation objects represent not only property-based
// animations, but also animations of geometry data created with external 3D
// authoring tools and loaded from a scene file. You use the properties of the
// [CAAnimation] object representing a geometry animation to control its
// timing, monitor its progress, and attach actions for Scene Kit to trigger
// during the animation. You can attach animations to Scene Kit objects that
// adopt the [SCNAnimatable] protocol, including nodes, geometries, and
// materials.
//
// In a Scene Kit app, [CAAnimation] objects support additional methods and
// properties, listed under Controlling SceneKit Animation Timing, Fading
// between SceneKit Animations, and Attaching SceneKit Animation Events.
//
// # Animation Attributes
//
//   - [CAAnimation.RemovedOnCompletion]: Determines if the animation is removed from the target layer’s animations upon completion.
//   - [CAAnimation.SetRemovedOnCompletion]
//   - [CAAnimation.TimingFunction]: An optional timing function defining the pacing of the animation.
//   - [CAAnimation.SetTimingFunction]
//
// # Designating a Delegate
//
//   - [CAAnimation.Delegate]: Specifies the receiver’s delegate object.
//   - [CAAnimation.SetDelegate]
//
// # Archiving Properties
//
//   - [CAAnimation.ShouldArchiveValueForKey]: Specifies whether the value of the property for a given key is archived.
//
// # Controlling SceneKit Animation Timing
//
//   - [CAAnimation.UsesSceneTimeBase]: For animations attached to SceneKit objects, a Boolean value that determines whether the animation is evaluated using the scene time or the system time.
//   - [CAAnimation.SetUsesSceneTimeBase]
//
// # Fading between SceneKit Animations
//
//   - [CAAnimation.FadeInDuration]: For animations attached to SceneKit objects, the duration for transitioning into the animation’s effect as it begins.
//   - [CAAnimation.SetFadeInDuration]
//   - [CAAnimation.FadeOutDuration]: For animations attached to SceneKit objects, the duration for transitioning out of the animation’s effect as it ends.
//   - [CAAnimation.SetFadeOutDuration]
//
// # Attaching SceneKit Animation Events
//
//   - [CAAnimation.AnimationEvents]: For animations attached to SceneKit objects, a list of events attached to an animation.
//   - [CAAnimation.SetAnimationEvents]
//
// # Instance Properties
//
//   - [CAAnimation.PreferredFrameRateRange]
//   - [CAAnimation.SetPreferredFrameRateRange]
//
// See: https://developer.apple.com/documentation/QuartzCore/CAAnimation
//
// [Core Animation Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/CoreAnimation_guide/Introduction/Introduction.html#//apple_ref/doc/uid/TP40004514
// [SCNAnimatable]: https://developer.apple.com/documentation/SceneKit/SCNAnimatable
type CAAnimation struct {
	objectivec.Object
}

// CAAnimationFromID constructs a [CAAnimation] from an objc.ID.
//
// The abstract superclass for animations in Core Animation.
func CAAnimationFromID(id objc.ID) CAAnimation {
	return CAAnimation{objectivec.Object{ID: id}}
}

// NOTE: CAAnimation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CAAnimation] class.
//
// # Animation Attributes
//
//   - [ICAAnimation.RemovedOnCompletion]: Determines if the animation is removed from the target layer’s animations upon completion.
//   - [ICAAnimation.SetRemovedOnCompletion]
//   - [ICAAnimation.TimingFunction]: An optional timing function defining the pacing of the animation.
//   - [ICAAnimation.SetTimingFunction]
//
// # Designating a Delegate
//
//   - [ICAAnimation.Delegate]: Specifies the receiver’s delegate object.
//   - [ICAAnimation.SetDelegate]
//
// # Archiving Properties
//
//   - [ICAAnimation.ShouldArchiveValueForKey]: Specifies whether the value of the property for a given key is archived.
//
// # Controlling SceneKit Animation Timing
//
//   - [ICAAnimation.UsesSceneTimeBase]: For animations attached to SceneKit objects, a Boolean value that determines whether the animation is evaluated using the scene time or the system time.
//   - [ICAAnimation.SetUsesSceneTimeBase]
//
// # Fading between SceneKit Animations
//
//   - [ICAAnimation.FadeInDuration]: For animations attached to SceneKit objects, the duration for transitioning into the animation’s effect as it begins.
//   - [ICAAnimation.SetFadeInDuration]
//   - [ICAAnimation.FadeOutDuration]: For animations attached to SceneKit objects, the duration for transitioning out of the animation’s effect as it ends.
//   - [ICAAnimation.SetFadeOutDuration]
//
// # Attaching SceneKit Animation Events
//
//   - [ICAAnimation.AnimationEvents]: For animations attached to SceneKit objects, a list of events attached to an animation.
//   - [ICAAnimation.SetAnimationEvents]
//
// # Instance Properties
//
//   - [ICAAnimation.PreferredFrameRateRange]
//   - [ICAAnimation.SetPreferredFrameRateRange]
//
// See: https://developer.apple.com/documentation/QuartzCore/CAAnimation
type ICAAnimation interface {
	objectivec.IObject
	CAAction
	CAMediaTiming

	// Topic: Animation Attributes

	// Determines if the animation is removed from the target layer’s animations upon completion.
	RemovedOnCompletion() bool
	SetRemovedOnCompletion(value bool)
	// An optional timing function defining the pacing of the animation.
	TimingFunction() ICAMediaTimingFunction
	SetTimingFunction(value ICAMediaTimingFunction)

	// Topic: Designating a Delegate

	// Specifies the receiver’s delegate object.
	Delegate() CAAnimationDelegate
	SetDelegate(value CAAnimationDelegate)

	// Topic: Archiving Properties

	// Specifies whether the value of the property for a given key is archived.
	ShouldArchiveValueForKey(key string) bool

	// Topic: Controlling SceneKit Animation Timing

	// For animations attached to SceneKit objects, a Boolean value that determines whether the animation is evaluated using the scene time or the system time.
	UsesSceneTimeBase() bool
	SetUsesSceneTimeBase(value bool)

	// Topic: Fading between SceneKit Animations

	// For animations attached to SceneKit objects, the duration for transitioning into the animation’s effect as it begins.
	FadeInDuration() float64
	SetFadeInDuration(value float64)
	// For animations attached to SceneKit objects, the duration for transitioning out of the animation’s effect as it ends.
	FadeOutDuration() float64
	SetFadeOutDuration(value float64)

	// Topic: Attaching SceneKit Animation Events

	// For animations attached to SceneKit objects, a list of events attached to an animation.
	AnimationEvents() []objectivec.IObject
	SetAnimationEvents(value []objectivec.IObject)

	// Topic: Instance Properties

	PreferredFrameRateRange() CAFrameRateRange
	SetPreferredFrameRateRange(value CAFrameRateRange)

	EncodeWithCoder(coder foundation.INSCoder)
	// Sets the value of the property identified by the given key.
	SetValueForKey(value objectivec.IObject, key string)
	// Returns the value of the property identified by the given key.
	ValueForKey(key string) objectivec.IObject
}

// Init initializes the instance.
func (a CAAnimation) Init() CAAnimation {
	rv := objc.Send[CAAnimation](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a CAAnimation) Autorelease() CAAnimation {
	rv := objc.Send[CAAnimation](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewCAAnimation creates a new CAAnimation instance.
func NewCAAnimation() CAAnimation {
	class := getCAAnimationClass()
	rv := objc.Send[CAAnimation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an animation from a SceneKit animation.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAAnimation/init(SCNAnimation:)
// animation is a [scenekit.SCNAnimation].
func NewAnimationWithSCNAnimation(animation objectivec.IObject) CAAnimation {
	rv := objc.Send[objc.ID](objc.ID(getCAAnimationClass().class), objc.Sel("animationWithSCNAnimation:"), animation)
	return CAAnimationFromID(rv)
}

// Specifies whether the value of the property for a given key is archived.
//
// key: The name of one of the receiver’s properties.
//
// # Return Value
//
// true if the specified property should be archived, otherwise false.
//
// # Discussion
//
// Called by the object’s implementation of “. The object must implement
// keyed archiving.
//
// The default implementation returns true.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAAnimation/shouldArchiveValue(forKey:)
func (a CAAnimation) ShouldArchiveValueForKey(key string) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("shouldArchiveValueForKey:"), objc.String(key))
	return rv
}

// Called to trigger the action specified by the identifier.
//
// event: The identifier of the action. The identifier may be a key or key path
// relative to `anObject`, an arbitrary external action, or one of the action
// identifiers defined in [CALayer].
//
// anObject: The layer on which the action should occur.
//
// dict: A dictionary containing parameters associated with this event. May be
// `nil`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAAction/run(forKey:object:arguments:)
func (a CAAnimation) RunActionForKeyObjectArguments(event string, anObject objectivec.IObject, dict foundation.INSDictionary) {
	objc.Send[objc.ID](a.ID, objc.Sel("runActionForKey:object:arguments:"), objc.String(event), anObject, dict)
}
func (a CAAnimation) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](a.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Sets the value of the property identified by the given key. [Full Topic]
func (a CAAnimation) SetValueForKey(value objectivec.IObject, key string) {
	objc.Send[objc.ID](a.ID, objc.Sel("setValue:forKey:"), value, objc.String(key))
}

// Returns the value of the property identified by the given key. [Full Topic]
func (a CAAnimation) ValueForKey(key string) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("valueForKey:"), objc.String(key))
	return objectivec.Object{ID: rv}
}

// Specifies the default value of the property with the specified key.
//
// key: The name of one of the receiver’s properties.
//
// # Return Value
//
// The default value for the named property. Returns `nil` if no default value
// has been set.
//
// # Discussion
//
// If this method returns `nil` a suitable “zero” default value for the
// property is provided, based on the declared type of the `key`. For example,
// if `key` is a [CGSize] object, a size of (0.0,0.0) is returned. For a
// [CGRect] an empty rectangle is returned. For [CGAffineTransform] and
// [CATransform3D], the appropriate identity matrix is returned.
//
// # Special Considerations
//
// If `key` is not a known for property of the class, the result of the method
// is undefined.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAAnimation/defaultValue(forKey:)
func (_CAAnimationClass CAAnimationClass) DefaultValueForKey(key string) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_CAAnimationClass.class), objc.Sel("defaultValueForKey:"), objc.String(key))
	return objectivec.Object{ID: rv}
}

// Creates and returns a new [CAAnimation] instance.
//
// # Return Value
//
// An [CAAnimation] object whose input values are initialized.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAAnimation/animation
func (_CAAnimationClass CAAnimationClass) Animation() CAAnimation {
	rv := objc.Send[objc.ID](objc.ID(_CAAnimationClass.class), objc.Sel("animation"))
	return CAAnimationFromID(rv)
}

// Determines if the animation is removed from the target layer’s animations
// upon completion.
//
// # Discussion
//
// When true, the animation is removed from the target layer’s animations
// once its active duration has passed. Defaults to true.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAAnimation/isRemovedOnCompletion
func (a CAAnimation) RemovedOnCompletion() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isRemovedOnCompletion"))
	return rv
}
func (a CAAnimation) SetRemovedOnCompletion(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setRemovedOnCompletion:"), value)
}

// An optional timing function defining the pacing of the animation.
//
// # Discussion
//
// Defaults to `nil`, indicating linear pacing.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAAnimation/timingFunction
func (a CAAnimation) TimingFunction() ICAMediaTimingFunction {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("timingFunction"))
	return CAMediaTimingFunctionFromID(objc.ID(rv))
}
func (a CAAnimation) SetTimingFunction(value ICAMediaTimingFunction) {
	objc.Send[struct{}](a.ID, objc.Sel("setTimingFunction:"), value)
}

// Specifies the receiver’s delegate object.
//
// # Discussion
//
// Defaults to `nil`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAAnimation/delegate
func (a CAAnimation) Delegate() CAAnimationDelegate {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("delegate"))
	return CAAnimationDelegateObjectFromID(rv)
}
func (a CAAnimation) SetDelegate(value CAAnimationDelegate) {
	objc.Send[struct{}](a.ID, objc.Sel("setDelegate:"), value)
}

// For animations attached to SceneKit objects, a Boolean value that
// determines whether the animation is evaluated using the scene time or the
// system time.
//
// # Discussion
//
// If the value of this property is true, animation timing is governed by the
// currentTime property of the view, layer, or custom renderer responsible for
// drawing the scene. The default value is false.
//
// To attach animations to SceneKit objects, see [SCNAnimatable].
//
// See: https://developer.apple.com/documentation/QuartzCore/CAAnimation/usesSceneTimeBase
//
// [SCNAnimatable]: https://developer.apple.com/documentation/SceneKit/SCNAnimatable
func (a CAAnimation) UsesSceneTimeBase() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("usesSceneTimeBase"))
	return rv
}
func (a CAAnimation) SetUsesSceneTimeBase(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setUsesSceneTimeBase:"), value)
}

// For animations attached to SceneKit objects, the duration for transitioning
// into the animation’s effect as it begins.
//
// # Discussion
//
// Use this property to create smooth transitions between the effects of
// multiple animations. These transitions are especially useful for geometry
// animations created with external 3D authoring tools.
//
// For example, the geometry loaded from a scene file for a game character may
// have associated animations for player actions such as walking and jumping.
// When the player jumps, if the fade duration is zero, SceneKit abruptly
// switches from the current frame of the walk animation to the first frame of
// the jump animation. If the fade duration is greater than zero, SceneKit
// plays both animations at once during that duration and interpolates vertex
// positions from one animation to the other, creating a smooth transition.
//
// To attach animations to SceneKit objects, see [SCNAnimatable].
//
// See: https://developer.apple.com/documentation/QuartzCore/CAAnimation/fadeInDuration
//
// [SCNAnimatable]: https://developer.apple.com/documentation/SceneKit/SCNAnimatable
func (a CAAnimation) FadeInDuration() float64 {
	rv := objc.Send[float64](a.ID, objc.Sel("fadeInDuration"))
	return rv
}
func (a CAAnimation) SetFadeInDuration(value float64) {
	objc.Send[struct{}](a.ID, objc.Sel("setFadeInDuration:"), value)
}

// For animations attached to SceneKit objects, the duration for transitioning
// out of the animation’s effect as it ends.
//
// # Discussion
//
// Use this property to create smooth transitions between the effects of
// multiple animations. These transitions are especially useful for geometry
// animations created with external 3D authoring tools.
//
// For example, the geometry loaded from a scene file for a game character may
// have associated animations for player actions such as walking and jumping.
// When the player jumps, if the fade duration is zero, SceneKit abruptly
// switches from the current frame of the walk animation to the first frame of
// the jump animation. If the fade duration is greater than zero, SceneKit
// plays both animations at once during that duration and interpolates vertex
// positions from one animation to the other, creating a smooth transition.
//
// To attach animations to SceneKit objects, see [SCNAnimatable].
//
// See: https://developer.apple.com/documentation/QuartzCore/CAAnimation/fadeOutDuration
//
// [SCNAnimatable]: https://developer.apple.com/documentation/SceneKit/SCNAnimatable
func (a CAAnimation) FadeOutDuration() float64 {
	rv := objc.Send[float64](a.ID, objc.Sel("fadeOutDuration"))
	return rv
}
func (a CAAnimation) SetFadeOutDuration(value float64) {
	objc.Send[struct{}](a.ID, objc.Sel("setFadeOutDuration:"), value)
}

// For animations attached to SceneKit objects, a list of events attached to
// an animation.
//
// # Discussion
//
// An array of [SCNAnimationEvent] objects, each of which adds a timed action
// to the animation.
//
// For example, you can create animation events that play sound effects timed
// to match the footsteps of an animated game character or that add new nodes
// to the scene when an animation completes.
//
// To attach animations to SceneKit objects, see [SCNAnimatable].
//
// See: https://developer.apple.com/documentation/QuartzCore/CAAnimation/animationEvents
//
// [SCNAnimatable]: https://developer.apple.com/documentation/SceneKit/SCNAnimatable
// [SCNAnimationEvent]: https://developer.apple.com/documentation/SceneKit/SCNAnimationEvent
func (a CAAnimation) AnimationEvents() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("animationEvents"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
func (a CAAnimation) SetAnimationEvents(value []objectivec.IObject) {
	objc.Send[struct{}](a.ID, objc.Sel("setAnimationEvents:"), objectivec.IObjectSliceToNSArray(value))
}

// Determines if the receiver plays in the reverse upon completion.
//
// # Discussion
//
// When true, the receiver plays backwards after playing forwards. Defaults to
// false.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/autoreverses
func (a CAAnimation) Autoreverses() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("autoreverses"))
	return rv
}
func (a CAAnimation) SetAutoreverses(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setAutoreverses:"), value)
}

// Specifies the begin time of the receiver in relation to its parent object,
// if applicable.
//
// # Discussion
//
// Defaults to 0.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/beginTime
func (a CAAnimation) BeginTime() float64 {
	rv := objc.Send[float64](a.ID, objc.Sel("beginTime"))
	return rv
}
func (a CAAnimation) SetBeginTime(value float64) {
	objc.Send[struct{}](a.ID, objc.Sel("setBeginTime:"), value)
}

// Specifies the basic duration of the animation, in seconds.
//
// # Discussion
//
// Defaults to 0.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/duration
func (a CAAnimation) Duration() float64 {
	rv := objc.Send[float64](a.ID, objc.Sel("duration"))
	return rv
}
func (a CAAnimation) SetDuration(value float64) {
	objc.Send[struct{}](a.ID, objc.Sel("setDuration:"), value)
}

// Determines if the receiver’s presentation is frozen or removed once its
// active duration has completed.
//
// # Discussion
//
// The possible values are described in [Fill Modes]. The default is
// [removed].
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/fillMode
//
// [Fill Modes]: https://developer.apple.com/documentation/QuartzCore/fill-modes
// [removed]: https://developer.apple.com/documentation/QuartzCore/CAMediaTimingFillMode/removed
func (a CAAnimation) FillMode() CAMediaTimingFillMode {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("fillMode"))
	return CAMediaTimingFillMode(foundation.NSStringFromID(rv).String())
}
func (a CAAnimation) SetFillMode(value CAMediaTimingFillMode) {
	objc.Send[struct{}](a.ID, objc.Sel("setFillMode:"), objc.String(string(value)))
}

// See: https://developer.apple.com/documentation/QuartzCore/CAAnimation/preferredFrameRateRange
func (a CAAnimation) PreferredFrameRateRange() CAFrameRateRange {
	rv := objc.Send[CAFrameRateRange](a.ID, objc.Sel("preferredFrameRateRange"))
	return CAFrameRateRange(rv)
}
func (a CAAnimation) SetPreferredFrameRateRange(value CAFrameRateRange) {
	objc.Send[struct{}](a.ID, objc.Sel("setPreferredFrameRateRange:"), value)
}

// Determines the number of times the animation will repeat.
//
// # Discussion
//
// May be fractional. If the `repeatCount` is 0, it is ignored. Defaults to 0.
// If both [RepeatDuration] and [RepeatCount] are specified the behavior is
// undefined.
//
// Setting this property to [greatestFiniteMagnitude] will cause the animation
// to repeat forever.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/repeatCount
//
// [greatestFiniteMagnitude]: https://developer.apple.com/documentation/Swift/Float/greatestFiniteMagnitude
func (a CAAnimation) RepeatCount() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("repeatCount"))
	return rv
}
func (a CAAnimation) SetRepeatCount(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setRepeatCount:"), value)
}

// Determines how many seconds the animation will repeat for.
//
// # Discussion
//
// Defaults to 0. If the `repeatDuration` is 0, it is ignored. If both
// [RepeatDuration] and [RepeatCount] are specified the behavior is undefined.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/repeatDuration
func (a CAAnimation) RepeatDuration() float64 {
	rv := objc.Send[float64](a.ID, objc.Sel("repeatDuration"))
	return rv
}
func (a CAAnimation) SetRepeatDuration(value float64) {
	objc.Send[struct{}](a.ID, objc.Sel("setRepeatDuration:"), value)
}

// Specifies how time is mapped to receiver’s time space from the parent
// time space.
//
// # Discussion
//
// For example, if `speed` is 2.0 local time progresses twice as fast as
// parent time. Defaults to 1.0.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/speed
func (a CAAnimation) Speed() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("speed"))
	return rv
}
func (a CAAnimation) SetSpeed(value float32) {
	objc.Send[struct{}](a.ID, objc.Sel("setSpeed:"), value)
}

// Specifies an additional time offset in active local time.
//
// # Discussion
//
// Defaults to 0. .
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/timeOffset
func (a CAAnimation) TimeOffset() float64 {
	rv := objc.Send[float64](a.ID, objc.Sel("timeOffset"))
	return rv
}
func (a CAAnimation) SetTimeOffset(value float64) {
	objc.Send[struct{}](a.ID, objc.Sel("setTimeOffset:"), value)
}

// Protocol methods for CAAction

// Protocol methods for CAMediaTiming
