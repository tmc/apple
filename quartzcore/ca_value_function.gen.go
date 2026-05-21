// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
// a [CABasicAnimation.FromValue] of `0`, a [CABasicAnimation.ToValue] of
// [pi], and a [CAPropertyAnimation.ValueFunction] of a [CAValueFunction] with
// a function name of [rotateZ].
//
// The following code shows how you would create such a rotation and apply it
// to a [CALayer] named `rotatingLayer`.
//
// The value functions [scale] and [translate] require 3 values, for the
// individual `x`, `y` and `z` components. When working with these value
// functions, you specify the animation’s [CABasicAnimation.FromValue] and
// [CABasicAnimation.ToValue] as arrays.
//
// The following code shows how you could animate a layer’s scale from `0`
// to `1` using a value function.
//
// # Getting Value Function Properties
//
//   - [CAValueFunction.Name]: Returns the name of the value function.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAValueFunction
//
// [pi]: https://developer.apple.com/documentation/Swift/Float/pi
// [rotateZ]: https://developer.apple.com/documentation/QuartzCore/CAValueFunctionName/rotateZ
// [scale]: https://developer.apple.com/documentation/QuartzCore/CAValueFunctionName/scale
// [translate]: https://developer.apple.com/documentation/QuartzCore/CAValueFunctionName/translate
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
// See: https://developer.apple.com/documentation/QuartzCore/CAValueFunction/init(name:)
//
// [Rotate Value Functions]: https://developer.apple.com/documentation/QuartzCore/rotate-value-functions
// [Scale Value Functions]: https://developer.apple.com/documentation/QuartzCore/scale-value-functions
// [Translate Functions]: https://developer.apple.com/documentation/QuartzCore/translate-functions
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
