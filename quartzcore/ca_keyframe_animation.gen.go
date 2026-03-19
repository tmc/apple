// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CAKeyframeAnimation] class.
var (
	_CAKeyframeAnimationClass     CAKeyframeAnimationClass
	_CAKeyframeAnimationClassOnce sync.Once
)

func getCAKeyframeAnimationClass() CAKeyframeAnimationClass {
	_CAKeyframeAnimationClassOnce.Do(func() {
		_CAKeyframeAnimationClass = CAKeyframeAnimationClass{class: objc.GetClass("CAKeyframeAnimation")}
	})
	return _CAKeyframeAnimationClass
}

// GetCAKeyframeAnimationClass returns the class object for CAKeyframeAnimation.
func GetCAKeyframeAnimationClass() CAKeyframeAnimationClass {
	return getCAKeyframeAnimationClass()
}

type CAKeyframeAnimationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (cc CAKeyframeAnimationClass) Alloc() CAKeyframeAnimation {
	rv := objc.Send[CAKeyframeAnimation](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object that provides keyframe animation capabilities for a layer object.
//
// # Overview
// 
// You create a [CAKeyframeAnimation] object using the inherited
// [CAKeyframeAnimation.AnimationWithKeyPath] method, specifying the key path of the property that
// you want to animate on the layer. You can then specify the keyframe values
// to use to control the timing and animation behavior.
// 
// For most types of animations, you specify the keyframe values using the
// [CAKeyframeAnimation.Values] and [CAKeyframeAnimation.KeyTimes] properties. During the animation, Core Animation
// generates intermediate values by interpolating between the values you
// provide. When animating a value that is a coordinate point, such as the
// layer’s position, you can specify a [CAKeyframeAnimation.Path] for that point to follow
// instead of individual values. The pacing of the animation is controlled by
// the timing information you provide.
// 
// The following code shows how to create a keyframe animation that animates a
// layer’s background color from red to green to blue over a two second
// duration.
//
// # Providing keyframe values
//
//   - [CAKeyframeAnimation.Values]: An array of objects that specify the keyframe values to use for the animation.
//   - [CAKeyframeAnimation.SetValues]
//   - [CAKeyframeAnimation.Path]: The path for a point-based property to follow.
//   - [CAKeyframeAnimation.SetPath]
//
// # Keyframe timing
//
//   - [CAKeyframeAnimation.KeyTimes]: An optional array of [NSNumber] objects that define the time at which to apply a given keyframe segment.
//   - [CAKeyframeAnimation.SetKeyTimes]
//   - [CAKeyframeAnimation.TimingFunctions]: An optional array of [CAMediaTimingFunction] objects that define the pacing for each keyframe segment.
//   - [CAKeyframeAnimation.SetTimingFunctions]
//   - [CAKeyframeAnimation.CalculationMode]: Specifies how intermediate keyframe values are calculated by the receiver.
//   - [CAKeyframeAnimation.SetCalculationMode]
//
// # Rotation Mode Attribute
//
//   - [CAKeyframeAnimation.RotationMode]: Determines whether objects animating along the path rotate to match the path tangent.
//   - [CAKeyframeAnimation.SetRotationMode]
//
// # Cubic Mode Attributes
//
//   - [CAKeyframeAnimation.TensionValues]: An array of numbers that define the tightness of the curve.
//   - [CAKeyframeAnimation.SetTensionValues]
//   - [CAKeyframeAnimation.ContinuityValues]: An array of numbers that define the sharpness of the timing curve’s corners.
//   - [CAKeyframeAnimation.SetContinuityValues]
//   - [CAKeyframeAnimation.BiasValues]: An array of numbers that define the position of the curve relative to a control point.
//   - [CAKeyframeAnimation.SetBiasValues]
//
// See: https://developer.apple.com/documentation/QuartzCore/CAKeyframeAnimation
type CAKeyframeAnimation struct {
	CAPropertyAnimation
}

// CAKeyframeAnimationFromID constructs a [CAKeyframeAnimation] from an objc.ID.
//
// An object that provides keyframe animation capabilities for a layer object.
func CAKeyframeAnimationFromID(id objc.ID) CAKeyframeAnimation {
	return CAKeyframeAnimation{CAPropertyAnimation: CAPropertyAnimationFromID(id)}
}
// NOTE: CAKeyframeAnimation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CAKeyframeAnimation] class.
//
// # Providing keyframe values
//
//   - [ICAKeyframeAnimation.Values]: An array of objects that specify the keyframe values to use for the animation.
//   - [ICAKeyframeAnimation.SetValues]
//   - [ICAKeyframeAnimation.Path]: The path for a point-based property to follow.
//   - [ICAKeyframeAnimation.SetPath]
//
// # Keyframe timing
//
//   - [ICAKeyframeAnimation.KeyTimes]: An optional array of [NSNumber] objects that define the time at which to apply a given keyframe segment.
//   - [ICAKeyframeAnimation.SetKeyTimes]
//   - [ICAKeyframeAnimation.TimingFunctions]: An optional array of [CAMediaTimingFunction] objects that define the pacing for each keyframe segment.
//   - [ICAKeyframeAnimation.SetTimingFunctions]
//   - [ICAKeyframeAnimation.CalculationMode]: Specifies how intermediate keyframe values are calculated by the receiver.
//   - [ICAKeyframeAnimation.SetCalculationMode]
//
// # Rotation Mode Attribute
//
//   - [ICAKeyframeAnimation.RotationMode]: Determines whether objects animating along the path rotate to match the path tangent.
//   - [ICAKeyframeAnimation.SetRotationMode]
//
// # Cubic Mode Attributes
//
//   - [ICAKeyframeAnimation.TensionValues]: An array of numbers that define the tightness of the curve.
//   - [ICAKeyframeAnimation.SetTensionValues]
//   - [ICAKeyframeAnimation.ContinuityValues]: An array of numbers that define the sharpness of the timing curve’s corners.
//   - [ICAKeyframeAnimation.SetContinuityValues]
//   - [ICAKeyframeAnimation.BiasValues]: An array of numbers that define the position of the curve relative to a control point.
//   - [ICAKeyframeAnimation.SetBiasValues]
//
// See: https://developer.apple.com/documentation/QuartzCore/CAKeyframeAnimation
type ICAKeyframeAnimation interface {
	ICAPropertyAnimation

	// Topic: Providing keyframe values

	// An array of objects that specify the keyframe values to use for the animation.
	Values() foundation.INSArray
	SetValues(value foundation.INSArray)
	// The path for a point-based property to follow.
	Path() coregraphics.CGPathRef
	SetPath(value coregraphics.CGPathRef)

	// Topic: Keyframe timing

	// An optional array of [NSNumber] objects that define the time at which to apply a given keyframe segment.
	KeyTimes() []foundation.NSNumber
	SetKeyTimes(value []foundation.NSNumber)
	// An optional array of [CAMediaTimingFunction] objects that define the pacing for each keyframe segment.
	TimingFunctions() []CAMediaTimingFunction
	SetTimingFunctions(value []CAMediaTimingFunction)
	// Specifies how intermediate keyframe values are calculated by the receiver.
	CalculationMode() CAAnimationCalculationMode
	SetCalculationMode(value CAAnimationCalculationMode)

	// Topic: Rotation Mode Attribute

	// Determines whether objects animating along the path rotate to match the path tangent.
	RotationMode() CAAnimationRotationMode
	SetRotationMode(value CAAnimationRotationMode)

	// Topic: Cubic Mode Attributes

	// An array of numbers that define the tightness of the curve.
	TensionValues() []foundation.NSNumber
	SetTensionValues(value []foundation.NSNumber)
	// An array of numbers that define the sharpness of the timing curve’s corners.
	ContinuityValues() []foundation.NSNumber
	SetContinuityValues(value []foundation.NSNumber)
	// An array of numbers that define the position of the curve relative to a control point.
	BiasValues() []foundation.NSNumber
	SetBiasValues(value []foundation.NSNumber)
}

// Init initializes the instance.
func (k CAKeyframeAnimation) Init() CAKeyframeAnimation {
	rv := objc.Send[CAKeyframeAnimation](k.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (k CAKeyframeAnimation) Autorelease() CAKeyframeAnimation {
	rv := objc.Send[CAKeyframeAnimation](k.ID, objc.Sel("autorelease"))
	return rv
}

// NewCAKeyframeAnimation creates a new CAKeyframeAnimation instance.
func NewCAKeyframeAnimation() CAKeyframeAnimation {
	class := getCAKeyframeAnimationClass()
	rv := objc.Send[CAKeyframeAnimation](objc.ID(class.class), objc.Sel("new"))
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
func NewKeyframeAnimationWithKeyPath(path string) CAKeyframeAnimation {
	rv := objc.Send[objc.ID](objc.ID(getCAKeyframeAnimationClass().class), objc.Sel("animationWithKeyPath:"), objc.String(path))
	return CAKeyframeAnimationFromID(rv)
}

// An array of objects that specify the keyframe values to use for the
// animation.
//
// # Discussion
// 
// The keyframe values represent the values through which the animation must
// proceed. The time at which a given keyframe value is applied to the layer
// depends on the animation timing, which is controlled by the
// [CalculationMode], [KeyTimes], and [TimingFunctions] properties. Values
// between keyframes are created using interpolation, unless the calculation
// mode is set to [discrete].
// 
// Depending on the type of the property, you may need to wrap the values in
// this array with an [NSNumber] of [NSValue] object. For some Core Graphics
// data types, you may also need to cast them to `id` before adding them to
// the array.
// 
// The values in this property are used only if the value in the [Path]
// property is `nil`.
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
// [NSValue]: https://developer.apple.com/documentation/Foundation/NSValue
// [discrete]: https://developer.apple.com/documentation/QuartzCore/CAAnimationCalculationMode/discrete
//
// See: https://developer.apple.com/documentation/QuartzCore/CAKeyframeAnimation/values
func (k CAKeyframeAnimation) Values() foundation.INSArray {
	rv := objc.Send[objc.ID](k.ID, objc.Sel("values"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (k CAKeyframeAnimation) SetValues(value foundation.INSArray) {
	objc.Send[struct{}](k.ID, objc.Sel("setValues:"), value)
}
// The path for a point-based property to follow.
//
// # Discussion
// 
// For layer properties that contain a [CGPoint] data type, the path object
// you assign to this property defines the values for that property over the
// length of the animation. If you specify a value for this property, any data
// in the [Values] property is ignored.
// 
// Any timing values you specify for the animation are applied to the points
// used to create the path. Paths can contain points defining move-to,
// line-to, or curve-to segments. The end point of a line-to or curve-to
// segment defines the keyframe value. All other points between that end value
// and the previous value are then interpolated. Move-to segments do not
// define separate keyframe values.
// 
// How the animation proceeds along the path is dependent on the value in the
// [CalculationMode] property. To achieve a smooth, constant velocity
// animation along the path, set the [CalculationMode] property to [paced] or
// [cubicPaced]. To create an animation where the location value jumps from
// keyframe point to keyframe point (without interpolation in between), use
// the [discrete] value. To animate along the path by interpolating values
// between points, use the [linear] value.
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
// [cubicPaced]: https://developer.apple.com/documentation/QuartzCore/CAAnimationCalculationMode/cubicPaced
// [discrete]: https://developer.apple.com/documentation/QuartzCore/CAAnimationCalculationMode/discrete
// [linear]: https://developer.apple.com/documentation/QuartzCore/CAAnimationCalculationMode/linear
// [paced]: https://developer.apple.com/documentation/QuartzCore/CAAnimationCalculationMode/paced
//
// See: https://developer.apple.com/documentation/QuartzCore/CAKeyframeAnimation/path
func (k CAKeyframeAnimation) Path() coregraphics.CGPathRef {
	rv := objc.Send[coregraphics.CGPathRef](k.ID, objc.Sel("path"))
	return coregraphics.CGPathRef(rv)
}
func (k CAKeyframeAnimation) SetPath(value coregraphics.CGPathRef) {
	objc.Send[struct{}](k.ID, objc.Sel("setPath:"), value)
}
// An optional array of [NSNumber] objects that define the time at which to
// apply a given keyframe segment.
//
// # Discussion
// 
// Each value in the array is a floating point number between `0.0` and `1.0`
// that defines the time point (specified as a fraction of the animation’s
// total duration) at which to apply the corresponding keyframe value. Each
// successive value in the array must be greater than, or equal to, the
// previous value. Usually, the number of elements in the array should match
// the number of elements in the [Values] property or the number of control
// points in the [Path] property. If they do not, the timing of your animation
// might not be what you expect.
// 
// The appropriate values to include in the array are dependent on the
// [CalculationMode] property.
// 
// - If the [CalculationMode] is set to [linear] or [cubic], the first value
// in the array must be `0.0` and the last value must be `1.0`. All
// intermediate values represent time points between the start and end times.
// - If the [CalculationMode] is set to [discrete], the first value in the
// array must be `0.0` and the last value must be `1.0`. The array should have
// one more entry than appears in the values array. For example, if there are
// two values, there should be three key times. - If the [CalculationMode] is
// set to [paced] or [cubicPaced], the values in this property are ignored.
// 
// If the values in this array are invalid or inappropriate for the current
// calculation mode, they are ignored.
//
// [cubicPaced]: https://developer.apple.com/documentation/QuartzCore/CAAnimationCalculationMode/cubicPaced
// [cubic]: https://developer.apple.com/documentation/QuartzCore/CAAnimationCalculationMode/cubic
// [discrete]: https://developer.apple.com/documentation/QuartzCore/CAAnimationCalculationMode/discrete
// [linear]: https://developer.apple.com/documentation/QuartzCore/CAAnimationCalculationMode/linear
// [paced]: https://developer.apple.com/documentation/QuartzCore/CAAnimationCalculationMode/paced
//
// See: https://developer.apple.com/documentation/QuartzCore/CAKeyframeAnimation/keyTimes
func (k CAKeyframeAnimation) KeyTimes() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](k.ID, objc.Sel("keyTimes"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}
func (k CAKeyframeAnimation) SetKeyTimes(value []foundation.NSNumber) {
	objc.Send[struct{}](k.ID, objc.Sel("setKeyTimes:"), objectivec.IObjectSliceToNSArray(value))
}
// An optional array of [CAMediaTimingFunction] objects that define the pacing
// for each keyframe segment.
//
// # Discussion
// 
// You can use this array to apply ease-in, ease-out, or custom timing curves
// to the points that lie between two keyframe values. If the number of
// keyframes in the values property is , then this property should contain
// `-1` objects.
// 
// If you provide timing information in the [KeyTimes] property, the timing
// functions you specify using this property further modify the timing between
// those values. If you do not assign a value to the [KeyTimes] property, the
// timing functions modify the default timing provided by the animation
// object.
// 
// If you also specify a timing function in the animation object’s
// [TimingFunction] property, that function is applied first followed by the
// timing function for the specific keyframe segment.
// 
// For information on how to create a timing function, see
// [CAMediaTimingFunction].
//
// See: https://developer.apple.com/documentation/QuartzCore/CAKeyframeAnimation/timingFunctions
func (k CAKeyframeAnimation) TimingFunctions() []CAMediaTimingFunction {
	rv := objc.Send[[]objc.ID](k.ID, objc.Sel("timingFunctions"))
	return objc.ConvertSlice(rv, func(id objc.ID) CAMediaTimingFunction {
		return CAMediaTimingFunctionFromID(id)
	})
}
func (k CAKeyframeAnimation) SetTimingFunctions(value []CAMediaTimingFunction) {
	objc.Send[struct{}](k.ID, objc.Sel("setTimingFunctions:"), objectivec.IObjectSliceToNSArray(value))
}
// Specifies how intermediate keyframe values are calculated by the receiver.
//
// # Discussion
// 
// The possible values are described in [Value calculation modes]. The default
// value of this property is [linear].
//
// [Value calculation modes]: https://developer.apple.com/documentation/QuartzCore/value-calculation-modes
// [linear]: https://developer.apple.com/documentation/QuartzCore/CAAnimationCalculationMode/linear
//
// See: https://developer.apple.com/documentation/QuartzCore/CAKeyframeAnimation/calculationMode
func (k CAKeyframeAnimation) CalculationMode() CAAnimationCalculationMode {
	rv := objc.Send[objc.ID](k.ID, objc.Sel("calculationMode"))
	return CAAnimationCalculationMode(foundation.NSStringFromID(rv).String())
}
func (k CAKeyframeAnimation) SetCalculationMode(value CAAnimationCalculationMode) {
	objc.Send[struct{}](k.ID, objc.Sel("setCalculationMode:"), objc.String(string(value)))
}
// Determines whether objects animating along the path rotate to match the
// path tangent.
//
// # Discussion
// 
// The possible values for this property are described in [Rotation Mode
// Values]. The default value of this property is `nil`, which indicates that
// objects should not rotate to follow the path.
// 
// The effect of setting this property to a non-`nil` value when no path
// object is supplied is undefined.
//
// [Rotation Mode Values]: https://developer.apple.com/documentation/QuartzCore/rotation-mode-values
//
// See: https://developer.apple.com/documentation/QuartzCore/CAKeyframeAnimation/rotationMode
func (k CAKeyframeAnimation) RotationMode() CAAnimationRotationMode {
	rv := objc.Send[objc.ID](k.ID, objc.Sel("rotationMode"))
	return CAAnimationRotationMode(foundation.NSStringFromID(rv).String())
}
func (k CAKeyframeAnimation) SetRotationMode(value CAAnimationRotationMode) {
	objc.Send[struct{}](k.ID, objc.Sel("setRotationMode:"), objc.String(string(value)))
}
// An array of numbers that define the tightness of the curve.
//
// # Discussion
// 
// This property is an array of [NSNumber] objects, used only for the cubic
// calculation modes. Positive values indicate a tighter curve while negative
// values indicate a rounder curve. The first value defines the behavior of
// the tangent to the first control point, the second value controls the
// second point’s tangents, and so on. If you do not specify a value for a
// given control point, the value `0` is used.
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
//
// See: https://developer.apple.com/documentation/QuartzCore/CAKeyframeAnimation/tensionValues
func (k CAKeyframeAnimation) TensionValues() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](k.ID, objc.Sel("tensionValues"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}
func (k CAKeyframeAnimation) SetTensionValues(value []foundation.NSNumber) {
	objc.Send[struct{}](k.ID, objc.Sel("setTensionValues:"), objectivec.IObjectSliceToNSArray(value))
}
// An array of numbers that define the sharpness of the timing curve’s
// corners.
//
// # Discussion
// 
// This property is an array of [NSNumber] objects, used only for the cubic
// calculation modes. Positive values result in sharper corners while negative
// values create inverted corners. The first value defines the behavior of the
// tangent to the first control point, the second value controls the second
// point’s tangents, and so on. If you do not specify a value for a given
// control point, the value `0` is used.
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
//
// See: https://developer.apple.com/documentation/QuartzCore/CAKeyframeAnimation/continuityValues
func (k CAKeyframeAnimation) ContinuityValues() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](k.ID, objc.Sel("continuityValues"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}
func (k CAKeyframeAnimation) SetContinuityValues(value []foundation.NSNumber) {
	objc.Send[struct{}](k.ID, objc.Sel("setContinuityValues:"), objectivec.IObjectSliceToNSArray(value))
}
// An array of numbers that define the position of the curve relative to a
// control point.
//
// # Discussion
// 
// This property is an array of [NSNumber] objects, used only for the cubic
// calculation modes. Positive values move the curve before the control point
// while negative values move it after the control point. The first value
// defines the behavior of the tangent to the first control point, the second
// value controls the second point’s tangents, and so on. If you do not
// specify a value for a given control point, the value `0` is used.
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
//
// See: https://developer.apple.com/documentation/QuartzCore/CAKeyframeAnimation/biasValues
func (k CAKeyframeAnimation) BiasValues() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](k.ID, objc.Sel("biasValues"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}
func (k CAKeyframeAnimation) SetBiasValues(value []foundation.NSNumber) {
	objc.Send[struct{}](k.ID, objc.Sel("setBiasValues:"), objectivec.IObjectSliceToNSArray(value))
}

