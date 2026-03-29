// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [CAPropertyAnimation] class.
var (
	_CAPropertyAnimationClass     CAPropertyAnimationClass
	_CAPropertyAnimationClassOnce sync.Once
)

func getCAPropertyAnimationClass() CAPropertyAnimationClass {
	_CAPropertyAnimationClassOnce.Do(func() {
		_CAPropertyAnimationClass = CAPropertyAnimationClass{class: objc.GetClass("CAPropertyAnimation")}
	})
	return _CAPropertyAnimationClass
}

// GetCAPropertyAnimationClass returns the class object for CAPropertyAnimation.
func GetCAPropertyAnimationClass() CAPropertyAnimationClass {
	return getCAPropertyAnimationClass()
}

type CAPropertyAnimationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CAPropertyAnimationClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CAPropertyAnimationClass) Alloc() CAPropertyAnimation {
	rv := objc.Send[CAPropertyAnimation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An abstract subclass for creating animations that manipulate the value of
// layer properties.
//
// # Overview
// 
// The property to animate is specified using a key path that is relative to
// the layer using the animation.
// 
// You do not create instances of [CAPropertyAnimation]: to animate the
// properties of a Core Animation layer, create instance of the concrete
// subclasses [CABasicAnimation] or [CAKeyframeAnimation].
//
// # Animated Key Path
//
//   - [CAPropertyAnimation.KeyPath]: Specifies the key path the receiver animates.
//   - [CAPropertyAnimation.SetKeyPath]
//
// # Property Value Calculation Behavior
//
//   - [CAPropertyAnimation.Cumulative]: Determines if the value of the property is the value at the end of the previous repeat cycle, plus the value of the current repeat cycle.
//   - [CAPropertyAnimation.SetCumulative]
//   - [CAPropertyAnimation.Additive]: Determines if the value specified by the animation is added to the current render tree value to produce the new render tree value.
//   - [CAPropertyAnimation.SetAdditive]
//   - [CAPropertyAnimation.ValueFunction]: An optional value function that is applied to interpolated values.
//   - [CAPropertyAnimation.SetValueFunction]
//
// See: https://developer.apple.com/documentation/QuartzCore/CAPropertyAnimation
type CAPropertyAnimation struct {
	CAAnimation
}

// CAPropertyAnimationFromID constructs a [CAPropertyAnimation] from an objc.ID.
//
// An abstract subclass for creating animations that manipulate the value of
// layer properties.
func CAPropertyAnimationFromID(id objc.ID) CAPropertyAnimation {
	return CAPropertyAnimation{CAAnimation: CAAnimationFromID(id)}
}
// NOTE: CAPropertyAnimation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CAPropertyAnimation] class.
//
// # Animated Key Path
//
//   - [ICAPropertyAnimation.KeyPath]: Specifies the key path the receiver animates.
//   - [ICAPropertyAnimation.SetKeyPath]
//
// # Property Value Calculation Behavior
//
//   - [ICAPropertyAnimation.Cumulative]: Determines if the value of the property is the value at the end of the previous repeat cycle, plus the value of the current repeat cycle.
//   - [ICAPropertyAnimation.SetCumulative]
//   - [ICAPropertyAnimation.Additive]: Determines if the value specified by the animation is added to the current render tree value to produce the new render tree value.
//   - [ICAPropertyAnimation.SetAdditive]
//   - [ICAPropertyAnimation.ValueFunction]: An optional value function that is applied to interpolated values.
//   - [ICAPropertyAnimation.SetValueFunction]
//
// See: https://developer.apple.com/documentation/QuartzCore/CAPropertyAnimation
type ICAPropertyAnimation interface {
	ICAAnimation

	// Topic: Animated Key Path

	// Specifies the key path the receiver animates.
	KeyPath() string
	SetKeyPath(value string)

	// Topic: Property Value Calculation Behavior

	// Determines if the value of the property is the value at the end of the previous repeat cycle, plus the value of the current repeat cycle.
	Cumulative() bool
	SetCumulative(value bool)
	// Determines if the value specified by the animation is added to the current render tree value to produce the new render tree value.
	Additive() bool
	SetAdditive(value bool)
	// An optional value function that is applied to interpolated values.
	ValueFunction() ICAValueFunction
	SetValueFunction(value ICAValueFunction)
}

// Init initializes the instance.
func (p CAPropertyAnimation) Init() CAPropertyAnimation {
	rv := objc.Send[CAPropertyAnimation](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p CAPropertyAnimation) Autorelease() CAPropertyAnimation {
	rv := objc.Send[CAPropertyAnimation](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewCAPropertyAnimation creates a new CAPropertyAnimation instance.
func NewCAPropertyAnimation() CAPropertyAnimation {
	class := getCAPropertyAnimationClass()
	rv := objc.Send[CAPropertyAnimation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates and returns an [CAPropertyAnimation] instance for the specified key
// path.
//
// path: The key path of the property to be animated.
//
// # Return Value
// 
// A new instance of [CAPropertyAnimation] with the key path set to `keyPath`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAPropertyAnimation/init(keyPath:)
func NewPropertyAnimationWithKeyPath(path string) CAPropertyAnimation {
	rv := objc.Send[objc.ID](objc.ID(getCAPropertyAnimationClass().class), objc.Sel("animationWithKeyPath:"), objc.String(path))
	return CAPropertyAnimationFromID(rv)
}

// Specifies the key path the receiver animates.
//
// # Discussion
// 
// The key path is relative to the layer the receiver is attached to.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAPropertyAnimation/keyPath
func (p CAPropertyAnimation) KeyPath() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("keyPath"))
	return foundation.NSStringFromID(rv).String()
}
func (p CAPropertyAnimation) SetKeyPath(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setKeyPath:"), objc.String(value))
}
// Determines if the value of the property is the value at the end of the
// previous repeat cycle, plus the value of the current repeat cycle.
//
// # Discussion
// 
// If [true], then the value of the property is the value at the end of the
// previous repeat cycle, plus the value of the current repeat cycle. If
// [false], the value of the property is simply the value calculated for the
// current repeat cycle. The default is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/QuartzCore/CAPropertyAnimation/isCumulative
func (p CAPropertyAnimation) Cumulative() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isCumulative"))
	return rv
}
func (p CAPropertyAnimation) SetCumulative(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setCumulative:"), value)
}
// Determines if the value specified by the animation is added to the current
// render tree value to produce the new render tree value.
//
// # Discussion
// 
// If [true], the value specified by the animation will be added to the
// current render tree value of the property to produce the new render tree
// value. The addition function is type-dependent, e.g. for affine transforms
// the two matrices are concatenated. The default is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/QuartzCore/CAPropertyAnimation/isAdditive
func (p CAPropertyAnimation) Additive() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isAdditive"))
	return rv
}
func (p CAPropertyAnimation) SetAdditive(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setAdditive:"), value)
}
// An optional value function that is applied to interpolated values.
//
// # Discussion
// 
// If the `valueFunction` property is not `nil`, the function is applied to
// the values interpolated by the animation as they are applied to the
// presentation layer. Defaults to `nil`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAPropertyAnimation/valueFunction
func (p CAPropertyAnimation) ValueFunction() ICAValueFunction {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("valueFunction"))
	return CAValueFunctionFromID(objc.ID(rv))
}
func (p CAPropertyAnimation) SetValueFunction(value ICAValueFunction) {
	objc.Send[struct{}](p.ID, objc.Sel("setValueFunction:"), value)
}

