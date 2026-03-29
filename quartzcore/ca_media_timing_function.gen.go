// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CAMediaTimingFunction] class.
var (
	_CAMediaTimingFunctionClass     CAMediaTimingFunctionClass
	_CAMediaTimingFunctionClassOnce sync.Once
)

func getCAMediaTimingFunctionClass() CAMediaTimingFunctionClass {
	_CAMediaTimingFunctionClassOnce.Do(func() {
		_CAMediaTimingFunctionClass = CAMediaTimingFunctionClass{class: objc.GetClass("CAMediaTimingFunction")}
	})
	return _CAMediaTimingFunctionClass
}

// GetCAMediaTimingFunctionClass returns the class object for CAMediaTimingFunction.
func GetCAMediaTimingFunctionClass() CAMediaTimingFunctionClass {
	return getCAMediaTimingFunctionClass()
}

type CAMediaTimingFunctionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CAMediaTimingFunctionClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CAMediaTimingFunctionClass) Alloc() CAMediaTimingFunction {
	rv := objc.Send[CAMediaTimingFunction](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A function that defines the pacing of an animation as a timing curve.
//
// # Overview
// 
// [CAMediaTimingFunction] represents one segment of a function that defines
// the pacing of an animation as a timing curve. The function maps an input
// time normalized to the range `[0,1]` to an output time also in the range
// `[0,1]`.
// 
// You can create a media timing function by supplying your own cubic Bézier
// curve control points using the [CAMediaTimingFunction.InitWithControlPoints] method or by using
// one of the predefined timing functions.
//
// # Creating Timing Functions
//
//   - [CAMediaTimingFunction.InitWithControlPoints]: Returns an initialized timing function modeled as a cubic Bézier curve using the specified control points.
//
// # Accessing the Control Points
//
//   - [CAMediaTimingFunction.GetControlPointAtIndexValues]: Returns the control point for the specified index.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTimingFunction
type CAMediaTimingFunction struct {
	objectivec.Object
}

// CAMediaTimingFunctionFromID constructs a [CAMediaTimingFunction] from an objc.ID.
//
// A function that defines the pacing of an animation as a timing curve.
func CAMediaTimingFunctionFromID(id objc.ID) CAMediaTimingFunction {
	return CAMediaTimingFunction{objectivec.Object{ID: id}}
}
// NOTE: CAMediaTimingFunction adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CAMediaTimingFunction] class.
//
// # Creating Timing Functions
//
//   - [ICAMediaTimingFunction.InitWithControlPoints]: Returns an initialized timing function modeled as a cubic Bézier curve using the specified control points.
//
// # Accessing the Control Points
//
//   - [ICAMediaTimingFunction.GetControlPointAtIndexValues]: Returns the control point for the specified index.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTimingFunction
type ICAMediaTimingFunction interface {
	objectivec.IObject

	// Topic: Creating Timing Functions

	// Returns an initialized timing function modeled as a cubic Bézier curve using the specified control points.
	InitWithControlPoints(c1x float32, c1y float32, c2x float32, c2y float32) CAMediaTimingFunction

	// Topic: Accessing the Control Points

	// Returns the control point for the specified index.
	GetControlPointAtIndexValues(idx uintptr, ptr unsafe.Pointer)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (m CAMediaTimingFunction) Init() CAMediaTimingFunction {
	rv := objc.Send[CAMediaTimingFunction](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m CAMediaTimingFunction) Autorelease() CAMediaTimingFunction {
	rv := objc.Send[CAMediaTimingFunction](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewCAMediaTimingFunction creates a new CAMediaTimingFunction instance.
func NewCAMediaTimingFunction() CAMediaTimingFunction {
	class := getCAMediaTimingFunctionClass()
	rv := objc.Send[CAMediaTimingFunction](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns an initialized timing function modeled as a cubic Bézier curve
// using the specified control points.
//
// c1x: A floating point number representing the x position of the c1 control
// point.
//
// c1y: A floating point number representing the y position of the c1 control
// point.
//
// c2x: A floating point number representing the x position of the c2 control
// point.
//
// c2y: A floating point number representing the y position of the c2 control
// point.
//
// # Return Value
// 
// An instance of [CAMediaTimingFunction] with the timing function specified
// by the provided control points.
//
// # Discussion
// 
// The end points of the Bézier curve are automatically set to (0.0,0.0) and
// (1.0,1.0). The control points defining the Bézier curve are: [(0.0,0.0),
// (`c1x`,`c1y`), (`c2x`,`c2y`), (1.0,1.0)].
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTimingFunction/init(controlPoints:_:_:_:)
func NewMediaTimingFunctionWithControlPoints(c1x float32, c1y float32, c2x float32, c2y float32) CAMediaTimingFunction {
	instance := getCAMediaTimingFunctionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithControlPoints::::"), c1x, c1y, c2x, c2y)
	return CAMediaTimingFunctionFromID(rv)
}

// Creates and returns a new instance of [CAMediaTimingFunction] configured
// with the predefined timing function specified by `name`.
//
// name: The timing function to use as specified in [Predefined Timing Functions].
// //
// [Predefined Timing Functions]: https://developer.apple.com/documentation/QuartzCore/predefined-timing-functions
//
// # Return Value
// 
// A new instance of [CAMediaTimingFunction] with the timing function
// specified by `name`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTimingFunction/init(name:)
func NewMediaTimingFunctionWithName(name CAMediaTimingFunctionName) CAMediaTimingFunction {
	rv := objc.Send[objc.ID](objc.ID(getCAMediaTimingFunctionClass().class), objc.Sel("functionWithName:"), objc.String(string(name)))
	return CAMediaTimingFunctionFromID(rv)
}

// Returns an initialized timing function modeled as a cubic Bézier curve
// using the specified control points.
//
// c1x: A floating point number representing the x position of the c1 control
// point.
//
// c1y: A floating point number representing the y position of the c1 control
// point.
//
// c2x: A floating point number representing the x position of the c2 control
// point.
//
// c2y: A floating point number representing the y position of the c2 control
// point.
//
// # Return Value
// 
// An instance of [CAMediaTimingFunction] with the timing function specified
// by the provided control points.
//
// # Discussion
// 
// The end points of the Bézier curve are automatically set to (0.0,0.0) and
// (1.0,1.0). The control points defining the Bézier curve are: [(0.0,0.0),
// (`c1x`,`c1y`), (`c2x`,`c2y`), (1.0,1.0)].
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTimingFunction/init(controlPoints:_:_:_:)
func (m CAMediaTimingFunction) InitWithControlPoints(c1x float32, c1y float32, c2x float32, c2y float32) CAMediaTimingFunction {
	rv := objc.Send[CAMediaTimingFunction](m.ID, objc.Sel("initWithControlPoints::::"), c1x, c1y, c2x, c2y)
	return rv
}
// Returns the control point for the specified index.
//
// idx: An integer specifying the index of the control point to return.
//
// ptr: A pointer to an array that, upon return, will contain the x and y values of
// the specified point.
//
// # Discussion
// 
// The value of `index` must be between 0 and 3.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTimingFunction/getControlPoint(at:values:)
func (m CAMediaTimingFunction) GetControlPointAtIndexValues(idx uintptr, ptr unsafe.Pointer) {
	objc.Send[objc.ID](m.ID, objc.Sel("getControlPointAtIndex:values:"), idx, ptr)
}
func (m CAMediaTimingFunction) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeWithCoder:"), coder)
}

