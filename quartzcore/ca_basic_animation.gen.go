// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CABasicAnimation] class.
var (
	_CABasicAnimationClass     CABasicAnimationClass
	_CABasicAnimationClassOnce sync.Once
)

func getCABasicAnimationClass() CABasicAnimationClass {
	_CABasicAnimationClassOnce.Do(func() {
		_CABasicAnimationClass = CABasicAnimationClass{class: objc.GetClass("CABasicAnimation")}
	})
	return _CABasicAnimationClass
}

// GetCABasicAnimationClass returns the class object for CABasicAnimation.
func GetCABasicAnimationClass() CABasicAnimationClass {
	return getCABasicAnimationClass()
}

type CABasicAnimationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CABasicAnimationClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CABasicAnimationClass) Alloc() CABasicAnimation {
	rv := objc.Send[CABasicAnimation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object that provides basic, single-keyframe animation capabilities for a
// layer property.
//
// # Overview
// 
// You create an instance of [CABasicAnimation] using the inherited
// [CABasicAnimation.AnimationWithKeyPath] method, specifying the key path of the property to
// be animated in the render tree.
// 
// For example, you can animate a layer’s scalar (i.e. containing a single
// value) properties such as its [CABasicAnimation.Opacity]. The following code fades in a
// layer by animating its opacity from `0` to `1`.
// 
// Non-scalar properties, such as [CABasicAnimation.BackgroundColor], can also be animated.
// Core Animation will interpolate between the [CABasicAnimation.FromValue] color and the
// [CABasicAnimation.ToValue] color. The animation created in the following code fades a
// layer’s background color from red to blue.
// 
// If you want to animate the individual components of a non-scalar property
// with different values, you pass the values to [CABasicAnimation.ToValue] and [CABasicAnimation.FromValue] as
// arrays. The following animation moves a layer from `(0, 0)` to `(100,
// 100)`.
// 
// The `keyPath` can access the individual components of a property. For
// example, the following animation stretches a layer by animating its
// [CABasicAnimation.Transform] object’s `x` from `1` to `2`.
// 
// # Setting Interpolation Values
// 
// The [CABasicAnimation.FromValue], [CABasicAnimation.ByValue] and [CABasicAnimation.ToValue] properties define the values being
// interpolated between. All are optional, and no more than two should be
// non-`nil`. The object type should match the type of the property being
// animated.
// 
// The interpolation values are used as follows:
// 
// - Both [CABasicAnimation.FromValue] and [CABasicAnimation.ToValue] are non-`nil`. Interpolates between
// [CABasicAnimation.FromValue] and [CABasicAnimation.ToValue]. - [CABasicAnimation.FromValue] and [CABasicAnimation.ByValue] are non-`nil`.
// Interpolates between [CABasicAnimation.FromValue] and ([CABasicAnimation.FromValue] + [CABasicAnimation.ByValue]). - [CABasicAnimation.ByValue]
// and [CABasicAnimation.ToValue] are non-`nil`. Interpolates between ([CABasicAnimation.ToValue] - [CABasicAnimation.ByValue])
// and [CABasicAnimation.ToValue]. - [CABasicAnimation.FromValue] is non-`nil`. Interpolates between [CABasicAnimation.FromValue]
// and the current presentation value of the property. - [CABasicAnimation.ToValue] is
// non-`nil`. Interpolates between the current value of `keyPath` in the
// target layer’s presentation layer and [CABasicAnimation.ToValue]. - [CABasicAnimation.ByValue] is
// non-`nil`. Interpolates between the current value of `keyPath` in the
// target layer’s presentation layer and that value plus [CABasicAnimation.ByValue]. - All
// properties are `nil`. Interpolates between the previous value of `keyPath`
// in the target layer’s presentation layer and the current value of
// `keyPath` in the target layer’s presentation layer.
//
// # Interpolation values
//
//   - [CABasicAnimation.FromValue]: Defines the value the receiver uses to start interpolation.
//   - [CABasicAnimation.SetFromValue]
//   - [CABasicAnimation.ToValue]: Defines the value the receiver uses to end interpolation.
//   - [CABasicAnimation.SetToValue]
//   - [CABasicAnimation.ByValue]: Defines the value the receiver uses to perform relative interpolation.
//   - [CABasicAnimation.SetByValue]
//
// See: https://developer.apple.com/documentation/QuartzCore/CABasicAnimation
type CABasicAnimation struct {
	CAPropertyAnimation
}

// CABasicAnimationFromID constructs a [CABasicAnimation] from an objc.ID.
//
// An object that provides basic, single-keyframe animation capabilities for a
// layer property.
func CABasicAnimationFromID(id objc.ID) CABasicAnimation {
	return CABasicAnimation{CAPropertyAnimation: CAPropertyAnimationFromID(id)}
}
// NOTE: CABasicAnimation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CABasicAnimation] class.
//
// # Interpolation values
//
//   - [ICABasicAnimation.FromValue]: Defines the value the receiver uses to start interpolation.
//   - [ICABasicAnimation.SetFromValue]
//   - [ICABasicAnimation.ToValue]: Defines the value the receiver uses to end interpolation.
//   - [ICABasicAnimation.SetToValue]
//   - [ICABasicAnimation.ByValue]: Defines the value the receiver uses to perform relative interpolation.
//   - [ICABasicAnimation.SetByValue]
//
// See: https://developer.apple.com/documentation/QuartzCore/CABasicAnimation
type ICABasicAnimation interface {
	ICAPropertyAnimation

	// Topic: Interpolation values

	// Defines the value the receiver uses to start interpolation.
	FromValue() objectivec.IObject
	SetFromValue(value objectivec.IObject)
	// Defines the value the receiver uses to end interpolation.
	ToValue() objectivec.IObject
	SetToValue(value objectivec.IObject)
	// Defines the value the receiver uses to perform relative interpolation.
	ByValue() objectivec.IObject
	SetByValue(value objectivec.IObject)

	// The background color of the receiver. Animatable.
	BackgroundColor() coregraphics.CGColorRef
	SetBackgroundColor(value coregraphics.CGColorRef)
	// The opacity of the receiver. Animatable.
	Opacity() float32
	SetOpacity(value float32)
	// The transform applied to the layer’s contents. Animatable.
	Transform() CATransform3D
	SetTransform(value CATransform3D)
}

// Init initializes the instance.
func (b CABasicAnimation) Init() CABasicAnimation {
	rv := objc.Send[CABasicAnimation](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b CABasicAnimation) Autorelease() CABasicAnimation {
	rv := objc.Send[CABasicAnimation](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewCABasicAnimation creates a new CABasicAnimation instance.
func NewCABasicAnimation() CABasicAnimation {
	class := getCABasicAnimationClass()
	rv := objc.Send[CABasicAnimation](objc.ID(class.class), objc.Sel("new"))
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
func NewBasicAnimationWithKeyPath(path string) CABasicAnimation {
	rv := objc.Send[objc.ID](objc.ID(getCABasicAnimationClass().class), objc.Sel("animationWithKeyPath:"), objc.String(path))
	return CABasicAnimationFromID(rv)
}

// Defines the value the receiver uses to start interpolation.
//
// # Discussion
// 
// See [CABasicAnimation] for details on how `fromValue` interacts with the
// other interpolation values.
//
// See: https://developer.apple.com/documentation/QuartzCore/CABasicAnimation/fromValue
func (b CABasicAnimation) FromValue() objectivec.IObject {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("fromValue"))
	return objectivec.Object{ID: rv}
}
func (b CABasicAnimation) SetFromValue(value objectivec.IObject) {
	objc.Send[struct{}](b.ID, objc.Sel("setFromValue:"), value)
}
// Defines the value the receiver uses to end interpolation.
//
// # Discussion
// 
// See [CABasicAnimation] for details on how `toValue` interacts with the
// other interpolation values.
//
// See: https://developer.apple.com/documentation/QuartzCore/CABasicAnimation/toValue
func (b CABasicAnimation) ToValue() objectivec.IObject {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("toValue"))
	return objectivec.Object{ID: rv}
}
func (b CABasicAnimation) SetToValue(value objectivec.IObject) {
	objc.Send[struct{}](b.ID, objc.Sel("setToValue:"), value)
}
// Defines the value the receiver uses to perform relative interpolation.
//
// # Discussion
// 
// See [CABasicAnimation] for details on how `byValue` interacts with the
// other interpolation values.
//
// See: https://developer.apple.com/documentation/QuartzCore/CABasicAnimation/byValue
func (b CABasicAnimation) ByValue() objectivec.IObject {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("byValue"))
	return objectivec.Object{ID: rv}
}
func (b CABasicAnimation) SetByValue(value objectivec.IObject) {
	objc.Send[struct{}](b.ID, objc.Sel("setByValue:"), value)
}
// The background color of the receiver. Animatable.
//
// See: https://developer.apple.com/documentation/quartzcore/calayer/backgroundcolor
func (b CABasicAnimation) BackgroundColor() coregraphics.CGColorRef {
	rv := objc.Send[coregraphics.CGColorRef](b.ID, objc.Sel("backgroundColor"))
	return coregraphics.CGColorRef(rv)
}
func (b CABasicAnimation) SetBackgroundColor(value coregraphics.CGColorRef) {
	objc.Send[struct{}](b.ID, objc.Sel("setBackgroundColor:"), value)
}
// The opacity of the receiver. Animatable.
//
// See: https://developer.apple.com/documentation/quartzcore/calayer/opacity
func (b CABasicAnimation) Opacity() float32 {
	rv := objc.Send[float32](b.ID, objc.Sel("opacity"))
	return rv
}
func (b CABasicAnimation) SetOpacity(value float32) {
	objc.Send[struct{}](b.ID, objc.Sel("setOpacity:"), value)
}
// The transform applied to the layer’s contents. Animatable.
//
// See: https://developer.apple.com/documentation/quartzcore/calayer/transform
func (b CABasicAnimation) Transform() CATransform3D {
	rv := objc.Send[CATransform3D](b.ID, objc.Sel("transform"))
	return CATransform3D(rv)
}
func (b CABasicAnimation) SetTransform(value CATransform3D) {
	objc.Send[struct{}](b.ID, objc.Sel("setTransform:"), value)
}

