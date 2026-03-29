// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CAValueFunction] class.
var (
	_CAValueFunctionClass     CAValueFunctionClass
	_CAValueFunctionClassOnce sync.Once
)

func getCAValueFunctionClass() CAValueFunctionClass {
	_CAValueFunctionClassOnce.Do(func() {
		_CAValueFunctionClass = CAValueFunctionClass{class: objc.GetClass("CAValueFunction")}
	})
	return _CAValueFunctionClass
}

// GetCAValueFunctionClass returns the class object for CAValueFunction.
func GetCAValueFunctionClass() CAValueFunctionClass {
	return getCAValueFunctionClass()
}

type CAValueFunctionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CAValueFunctionClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CAValueFunctionClass) Alloc() CAValueFunction {
	rv := objc.Send[CAValueFunction](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object that provides a flexible method of defining animated
// transformations.
//
// # Overview
// 
// You can use a value function to specify the individual components of an
// animated transform.
// 
// For example, to create a basic animation that rotates a layer from 0° to
// 180° around its z-axis, you would create a [CABasicAnimation] object with
// a [CAValueFunction.FromValue] of `0`, a [CAValueFunction.ToValue] of [CAValueFunction.Pi], and a [ValueFunction] of a
// [CAValueFunction] with a function name of [rotateZ].
// 
// The following code shows how you would create such a rotation and apply it
// to a [CALayer] named `rotatingLayer`.
// 
// The value functions [scale] and [translate] require 3 values, for the
// individual `x`, `y` and `z` components. When working with these value
// functions, you specify the animation’s [CAValueFunction.FromValue] and [CAValueFunction.ToValue] as
// arrays.
// 
// The following code shows how you could animate a layer’s scale from `0`
// to `1` using a value function.
//
// [rotateZ]: https://developer.apple.com/documentation/QuartzCore/CAValueFunctionName/rotateZ
// [scale]: https://developer.apple.com/documentation/QuartzCore/CAValueFunctionName/scale
// [translate]: https://developer.apple.com/documentation/QuartzCore/CAValueFunctionName/translate
//
// # Getting Value Function Properties
//
//   - [CAValueFunction.Name]: Returns the name of the value function.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAValueFunction
type CAValueFunction struct {
	objectivec.Object
}

// CAValueFunctionFromID constructs a [CAValueFunction] from an objc.ID.
//
// An object that provides a flexible method of defining animated
// transformations.
func CAValueFunctionFromID(id objc.ID) CAValueFunction {
	return CAValueFunction{objectivec.Object{ID: id}}
}
// NOTE: CAValueFunction adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CAValueFunction] class.
//
// # Getting Value Function Properties
//
//   - [ICAValueFunction.Name]: Returns the name of the value function.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAValueFunction
type ICAValueFunction interface {
	objectivec.IObject

	// Topic: Getting Value Function Properties

	// Returns the name of the value function.
	Name() CAValueFunctionName

	// Defines the value the receiver uses to start interpolation.
	FromValue() objectivec.IObject
	SetFromValue(value objectivec.IObject)
	// The mathematical constant pi (π), approximately equal to 3.14159.
	Pi() float32
	SetPi(value float32)
	// A value function that rotates by the input value, in radians, around the z-axis. This value function expects a single input value.
	RotateZ() CAValueFunctionName
	// A value function scales by the input value along all three axis. Animations using this value transform function must provide animation values in an 
	Scale() CAValueFunctionName
	// Defines the value the receiver uses to end interpolation.
	ToValue() objectivec.IObject
	SetToValue(value objectivec.IObject)
	// A value function that translates by the input values along all three axis. Animations using this value transform function must provide animation values in an 
	Translate() CAValueFunctionName
	// An optional value function that is applied to interpolated values.
	ValueFunction() ICAValueFunction
	SetValueFunction(value ICAValueFunction)
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (v CAValueFunction) Init() CAValueFunction {
	rv := objc.Send[CAValueFunction](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v CAValueFunction) Autorelease() CAValueFunction {
	rv := objc.Send[CAValueFunction](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewCAValueFunction creates a new CAValueFunction instance.
func NewCAValueFunction() CAValueFunction {
	class := getCAValueFunctionClass()
	rv := objc.Send[CAValueFunction](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the value function object identified by the name.
//
// name: The name of the value function.
//
// # Return Value
// 
// A new [CAValueFunction] instance with the value function specified by the
// name.
//
// # Discussion
// 
// The possible values for `name` are specified in [Rotate Value Functions],
// [Scale Value Functions], and [Translate Functions].
//
// [Rotate Value Functions]: https://developer.apple.com/documentation/QuartzCore/rotate-value-functions
// [Scale Value Functions]: https://developer.apple.com/documentation/QuartzCore/scale-value-functions
// [Translate Functions]: https://developer.apple.com/documentation/QuartzCore/translate-functions
//
// See: https://developer.apple.com/documentation/QuartzCore/CAValueFunction/init(name:)
func NewValueFunctionWithName(name CAValueFunctionName) CAValueFunction {
	rv := objc.Send[objc.ID](objc.ID(getCAValueFunctionClass().class), objc.Sel("functionWithName:"), objc.String(string(name)))
	return CAValueFunctionFromID(rv)
}

func (v CAValueFunction) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](v.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Returns the name of the value function.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAValueFunction/name
func (v CAValueFunction) Name() CAValueFunctionName {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("name"))
	return CAValueFunctionName(foundation.NSStringFromID(rv).String())
}
// Defines the value the receiver uses to start interpolation.
//
// See: https://developer.apple.com/documentation/quartzcore/cabasicanimation/fromvalue
func (v CAValueFunction) FromValue() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("fromValue"))
	return objectivec.Object{ID: rv}
}
func (v CAValueFunction) SetFromValue(value objectivec.IObject) {
	objc.Send[struct{}](v.ID, objc.Sel("setFromValue:"), value)
}
// The mathematical constant pi (π), approximately equal to 3.14159.
//
// See: https://developer.apple.com/documentation/Swift/Float/pi
func (v CAValueFunction) Pi() float32 {
	rv := objc.Send[float32](v.ID, objc.Sel("pi"))
	return rv
}
func (v CAValueFunction) SetPi(value float32) {
	objc.Send[struct{}](v.ID, objc.Sel("setPi:"), value)
}
// A value function that rotates by the input value, in radians, around the
// z-axis. This value function expects a single input value.
//
// See: https://developer.apple.com/documentation/quartzcore/cavaluefunctionname/rotatez
func (v CAValueFunction) RotateZ() CAValueFunctionName {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("kCAValueFunctionRotateZ"))
	return CAValueFunctionName(foundation.NSStringFromID(rv).String())
}
// A value function scales by the input value along all three axis. Animations
// using this value transform function must provide animation values in an
//
// See: https://developer.apple.com/documentation/quartzcore/cavaluefunctionname/scale
func (v CAValueFunction) Scale() CAValueFunctionName {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("kCAValueFunctionScale"))
	return CAValueFunctionName(foundation.NSStringFromID(rv).String())
}
// Defines the value the receiver uses to end interpolation.
//
// See: https://developer.apple.com/documentation/quartzcore/cabasicanimation/tovalue
func (v CAValueFunction) ToValue() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("toValue"))
	return objectivec.Object{ID: rv}
}
func (v CAValueFunction) SetToValue(value objectivec.IObject) {
	objc.Send[struct{}](v.ID, objc.Sel("setToValue:"), value)
}
// A value function that translates by the input values along all three axis.
// Animations using this value transform function must provide animation
// values in an
//
// See: https://developer.apple.com/documentation/quartzcore/cavaluefunctionname/translate
func (v CAValueFunction) Translate() CAValueFunctionName {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("kCAValueFunctionTranslate"))
	return CAValueFunctionName(foundation.NSStringFromID(rv).String())
}
// An optional value function that is applied to interpolated values.
//
// See: https://developer.apple.com/documentation/quartzcore/capropertyanimation/valuefunction
func (v CAValueFunction) ValueFunction() ICAValueFunction {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("valueFunction"))
	return CAValueFunctionFromID(objc.ID(rv))
}
func (v CAValueFunction) SetValueFunction(value ICAValueFunction) {
	objc.Send[struct{}](v.ID, objc.Sel("setValueFunction:"), value)
}

