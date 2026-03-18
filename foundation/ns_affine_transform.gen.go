// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSAffineTransform] class.
var (
	_NSAffineTransformClass     NSAffineTransformClass
	_NSAffineTransformClassOnce sync.Once
)

func getNSAffineTransformClass() NSAffineTransformClass {
	_NSAffineTransformClassOnce.Do(func() {
		_NSAffineTransformClass = NSAffineTransformClass{class: objc.GetClass("NSAffineTransform")}
	})
	return _NSAffineTransformClass
}

// GetNSAffineTransformClass returns the class object for NSAffineTransform.
func GetNSAffineTransformClass() NSAffineTransformClass {
	return getNSAffineTransformClass()
}

type NSAffineTransformClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSAffineTransformClass) Alloc() NSAffineTransform {
	rv := objc.Send[NSAffineTransform](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A graphics coordinate transformation.
//
// # Overview
// 
// In Swift, this object bridges to [AffineTransform]; use [NSAffineTransform]
// when you need reference semantics or other Foundation-specific behavior.
// 
// A transformation specifies how points in one coordinate system are
// transformed to points in another coordinate system. An affine
// transformation is a special type of transformation that preserves parallel
// lines in a path but does not necessarily preserve lengths or angles.
// Scaling, rotation, and translation are the most commonly used manipulations
// supported by affine transforms, but shearing is also possible.
// 
// Methods for applying affine transformations to the current graphics context
// and a method for applying an affine transformation to an [NSBezierPath]
// object are described in NSAffineTransform Additions Reference in the
// Application Kit.
//
// [AffineTransform]: https://developer.apple.com/documentation/Foundation/AffineTransform
// [NSBezierPath]: https://developer.apple.com/documentation/AppKit/NSBezierPath
//
// # Creating an Affine Transform
//
//   - [NSAffineTransform.InitWithTransform]: Initializes the receiver’s matrix using another transform object.
//
// # Accumulating Transformations
//
//   - [NSAffineTransform.RotateByDegrees]: Applies a rotation factor (measured in degrees) to the receiver’s transformation matrix.
//   - [NSAffineTransform.RotateByRadians]: Applies a rotation factor (measured in radians) to the receiver’s transformation matrix.
//   - [NSAffineTransform.ScaleBy]: Applies the specified scaling factor along both x and y axes to the receiver’s transformation matrix.
//   - [NSAffineTransform.ScaleXByYBy]: Applies scaling factors to each axis of the receiver’s transformation matrix.
//   - [NSAffineTransform.TranslateXByYBy]: Applies the specified translation factors to the receiver’s transformation matrix.
//   - [NSAffineTransform.AppendTransform]: Appends the specified matrix to the receiver’s matrix.
//   - [NSAffineTransform.PrependTransform]: Prepends the specified matrix to the receiver’s matrix.
//   - [NSAffineTransform.Invert]: Replaces the receiver’s matrix with its inverse matrix.
//
// # Transforming Data and Objects
//
//   - [NSAffineTransform.TransformPoint]: Applies the receiver’s transform to the specified point and returns the result.
//   - [NSAffineTransform.TransformSize]: Applies the receiver’s transform to the specified size and returns the results.
//   - [NSAffineTransform.TransformBezierPath]: Creates and returns a new Bézier path object with each point in the given path transformed by the receiver.
//
// # Accessing the Transformation Matrix
//
//   - [NSAffineTransform.TransformStruct]: The matrix coefficients stored as the transformation matrix.
//   - [NSAffineTransform.SetTransformStruct]
//
// # Setting and Building the Current Transformation Matrix
//
//   - [NSAffineTransform.Set]: Sets the current transformation matrix to the receiver’s transformation matrix.
//   - [NSAffineTransform.Concat]: Appends the receiver’s matrix to the current transformation matrix stored in the current graphics context, replacing the current transformation matrix with the result.
//
// See: https://developer.apple.com/documentation/Foundation/NSAffineTransform
type NSAffineTransform struct {
	objectivec.Object
}

// NSAffineTransformFromID constructs a [NSAffineTransform] from an objc.ID.
//
// A graphics coordinate transformation.
func NSAffineTransformFromID(id objc.ID) NSAffineTransform {
	return NSAffineTransform{objectivec.Object{ID: id}}
}
// NOTE: NSAffineTransform adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSAffineTransform] class.
//
// # Creating an Affine Transform
//
//   - [INSAffineTransform.InitWithTransform]: Initializes the receiver’s matrix using another transform object.
//
// # Accumulating Transformations
//
//   - [INSAffineTransform.RotateByDegrees]: Applies a rotation factor (measured in degrees) to the receiver’s transformation matrix.
//   - [INSAffineTransform.RotateByRadians]: Applies a rotation factor (measured in radians) to the receiver’s transformation matrix.
//   - [INSAffineTransform.ScaleBy]: Applies the specified scaling factor along both x and y axes to the receiver’s transformation matrix.
//   - [INSAffineTransform.ScaleXByYBy]: Applies scaling factors to each axis of the receiver’s transformation matrix.
//   - [INSAffineTransform.TranslateXByYBy]: Applies the specified translation factors to the receiver’s transformation matrix.
//   - [INSAffineTransform.AppendTransform]: Appends the specified matrix to the receiver’s matrix.
//   - [INSAffineTransform.PrependTransform]: Prepends the specified matrix to the receiver’s matrix.
//   - [INSAffineTransform.Invert]: Replaces the receiver’s matrix with its inverse matrix.
//
// # Transforming Data and Objects
//
//   - [INSAffineTransform.TransformPoint]: Applies the receiver’s transform to the specified point and returns the result.
//   - [INSAffineTransform.TransformSize]: Applies the receiver’s transform to the specified size and returns the results.
//   - [INSAffineTransform.TransformBezierPath]: Creates and returns a new Bézier path object with each point in the given path transformed by the receiver.
//
// # Accessing the Transformation Matrix
//
//   - [INSAffineTransform.TransformStruct]: The matrix coefficients stored as the transformation matrix.
//   - [INSAffineTransform.SetTransformStruct]
//
// # Setting and Building the Current Transformation Matrix
//
//   - [INSAffineTransform.Set]: Sets the current transformation matrix to the receiver’s transformation matrix.
//   - [INSAffineTransform.Concat]: Appends the receiver’s matrix to the current transformation matrix stored in the current graphics context, replacing the current transformation matrix with the result.
//
// See: https://developer.apple.com/documentation/Foundation/NSAffineTransform
type INSAffineTransform interface {
	objectivec.IObject
	NSCoding
	NSCopying
	NSSecureCoding

	// Topic: Creating an Affine Transform

	// Initializes the receiver’s matrix using another transform object.
	InitWithTransform(transform INSAffineTransform) NSAffineTransform

	// Topic: Accumulating Transformations

	// Applies a rotation factor (measured in degrees) to the receiver’s transformation matrix.
	RotateByDegrees(angle float64)
	// Applies a rotation factor (measured in radians) to the receiver’s transformation matrix.
	RotateByRadians(angle float64)
	// Applies the specified scaling factor along both x and y axes to the receiver’s transformation matrix.
	ScaleBy(scale float64)
	// Applies scaling factors to each axis of the receiver’s transformation matrix.
	ScaleXByYBy(scaleX float64, scaleY float64)
	// Applies the specified translation factors to the receiver’s transformation matrix.
	TranslateXByYBy(deltaX float64, deltaY float64)
	// Appends the specified matrix to the receiver’s matrix.
	AppendTransform(transform INSAffineTransform)
	// Prepends the specified matrix to the receiver’s matrix.
	PrependTransform(transform INSAffineTransform)
	// Replaces the receiver’s matrix with its inverse matrix.
	Invert()

	// Topic: Transforming Data and Objects

	// Applies the receiver’s transform to the specified point and returns the result.
	TransformPoint(aPoint corefoundation.CGPoint) NSPoint
	// Applies the receiver’s transform to the specified size and returns the results.
	TransformSize(aSize corefoundation.CGSize) NSSize
	// Creates and returns a new Bézier path object with each point in the given path transformed by the receiver.
	TransformBezierPath(path objectivec.IObject) objectivec.IObject

	// Topic: Accessing the Transformation Matrix

	// The matrix coefficients stored as the transformation matrix.
	TransformStruct() NSAffineTransformStruct
	SetTransformStruct(value NSAffineTransformStruct)

	// Topic: Setting and Building the Current Transformation Matrix

	// Sets the current transformation matrix to the receiver’s transformation matrix.
	Set()
	// Appends the receiver’s matrix to the current transformation matrix stored in the current graphics context, replacing the current transformation matrix with the result.
	Concat()
}

// Init initializes the instance.
func (a NSAffineTransform) Init() NSAffineTransform {
	rv := objc.Send[NSAffineTransform](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a NSAffineTransform) Autorelease() NSAffineTransform {
	rv := objc.Send[NSAffineTransform](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSAffineTransform creates a new NSAffineTransform instance.
func NewNSAffineTransform() NSAffineTransform {
	class := getNSAffineTransformClass()
	rv := objc.Send[NSAffineTransform](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewAffineTransformWithCoder(coder INSCoder) NSAffineTransform {
	instance := getNSAffineTransformClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSAffineTransformFromID(rv)
}

// Initializes the receiver’s matrix using another transform object.
//
// transform: The transform object whose matrix values should be copied to this object.
//
// # Return Value
// 
// A new transform object initialized with the matrix values of `aTransform`.
//
// See: https://developer.apple.com/documentation/Foundation/NSAffineTransform/init(transform:)
func NewAffineTransformWithTransform(transform INSAffineTransform) NSAffineTransform {
	instance := getNSAffineTransformClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTransform:"), transform)
	return NSAffineTransformFromID(rv)
}

// Initializes the receiver’s matrix using another transform object.
//
// transform: The transform object whose matrix values should be copied to this object.
//
// # Return Value
// 
// A new transform object initialized with the matrix values of `aTransform`.
//
// See: https://developer.apple.com/documentation/Foundation/NSAffineTransform/init(transform:)
func (a NSAffineTransform) InitWithTransform(transform INSAffineTransform) NSAffineTransform {
	rv := objc.Send[NSAffineTransform](a.ID, objc.Sel("initWithTransform:"), transform)
	return rv
}

// Applies a rotation factor (measured in degrees) to the receiver’s
// transformation matrix.
//
// angle: The rotation angle, measured in degrees.
//
// # Discussion
// 
// After invoking this method, applying the receiver’s matrix turns the axes
// counterclockwise about the current origin by `angle` degrees, in addition
// to performing all previous transformations.
//
// See: https://developer.apple.com/documentation/Foundation/NSAffineTransform/rotate(byDegrees:)
func (a NSAffineTransform) RotateByDegrees(angle float64) {
	objc.Send[objc.ID](a.ID, objc.Sel("rotateByDegrees:"), angle)
}

// Applies a rotation factor (measured in radians) to the receiver’s
// transformation matrix.
//
// angle: The rotation angle, measured in radians.
//
// # Discussion
// 
// After invoking this method, applying the receiver’s matrix turns the axes
// counterclockwise about the current origin by `angle` radians, in addition
// to performing all previous transformations.
//
// See: https://developer.apple.com/documentation/Foundation/NSAffineTransform/rotate(byRadians:)
func (a NSAffineTransform) RotateByRadians(angle float64) {
	objc.Send[objc.ID](a.ID, objc.Sel("rotateByRadians:"), angle)
}

// Applies the specified scaling factor along both x and y axes to the
// receiver’s transformation matrix.
//
// scale: The scaling factor to apply to both axes. Specifying a negative value has
// the effect of inverting the direction of the axes in addition to scaling
// them. A scaling factor of 1.0 scales the content to exactly the same size.
//
// # Discussion
// 
// After invoking this method, applying the receiver’s matrix modifies the
// unit lengths along the current x and y axes by a factor of `scale`, in
// addition to performing all previous transformations.
//
// See: https://developer.apple.com/documentation/Foundation/NSAffineTransform/scale(by:)
func (a NSAffineTransform) ScaleBy(scale float64) {
	objc.Send[objc.ID](a.ID, objc.Sel("scaleBy:"), scale)
}

// Applies scaling factors to each axis of the receiver’s transformation
// matrix.
//
// scaleX: The scaling factor to apply to the x axis.
//
// scaleY: The scaling factor to apply to the y axis.
//
// # Discussion
// 
// After invoking this method, applying the receiver’s matrix modifies the
// unit length on the x axis by a factor of `scaleX` and the y axis by a
// factor of `scaleY`, in addition to performing all previous transformations.
// A value of 1.0 for either axis scales the content on that axis to the same
// size.
//
// See: https://developer.apple.com/documentation/Foundation/NSAffineTransform/scaleX(by:yBy:)
func (a NSAffineTransform) ScaleXByYBy(scaleX float64, scaleY float64) {
	objc.Send[objc.ID](a.ID, objc.Sel("scaleXBy:yBy:"), scaleX, scaleY)
}

// Applies the specified translation factors to the receiver’s
// transformation matrix.
//
// deltaX: The number of units to move along the x axis.
//
// deltaY: The number of units to move along the y axis.
//
// # Discussion
// 
// Subsequent transformations cause coordinates to be shifted by `deltaX`
// units along the x axis and by `deltaY` units along the y axis. Translation
// factors do not affect [NSSize] values, which specify a differential between
// points.
//
// See: https://developer.apple.com/documentation/Foundation/NSAffineTransform/translateX(by:yBy:)
func (a NSAffineTransform) TranslateXByYBy(deltaX float64, deltaY float64) {
	objc.Send[objc.ID](a.ID, objc.Sel("translateXBy:yBy:"), deltaX, deltaY)
}

// Appends the specified matrix to the receiver’s matrix.
//
// transform: The matrix to append to the receiver.
//
// # Discussion
// 
// This method multiplies the receiver’s matrix by the matrix in
// `aTransform` and replaces the receiver’s matrix with the results. This
// type of operation is the same as applying the transformations in the
// receiver followed by the transformations in `aTransform`.
//
// See: https://developer.apple.com/documentation/Foundation/NSAffineTransform/append(_:)
func (a NSAffineTransform) AppendTransform(transform INSAffineTransform) {
	objc.Send[objc.ID](a.ID, objc.Sel("appendTransform:"), transform)
}

// Prepends the specified matrix to the receiver’s matrix.
//
// transform: The matrix to prepend to the receiver.
//
// # Discussion
// 
// This method multiplies the matrix in `transform` by the receiver’s matrix
// and replaces the receiver’s matrix with the result. This type of
// operation is the same as applying the transformations in `transform`
// followed by the transformations in the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSAffineTransform/prepend(_:)
func (a NSAffineTransform) PrependTransform(transform INSAffineTransform) {
	objc.Send[objc.ID](a.ID, objc.Sel("prependTransform:"), transform)
}

// Replaces the receiver’s matrix with its inverse matrix.
//
// # Discussion
// 
// Inverse matrices are useful for undoing the effects of a matrix. If a
// previous point (x,y) was transformed to (x’,y’), inverting the matrix
// and applying it to point (x’,y’) yields the point (x,y).
// 
// You can also use inverse matrices in conjunction with the [Concat] method
// to remove the effects of concatenating the matrix to the current
// transformation matrix of the current graphic context.
//
// See: https://developer.apple.com/documentation/Foundation/NSAffineTransform/invert()
func (a NSAffineTransform) Invert() {
	objc.Send[objc.ID](a.ID, objc.Sel("invert"))
}

// Applies the receiver’s transform to the specified point and returns the
// result.
//
// aPoint: The point in the current coordinate system to which you want to apply the
// matrix.
//
// # Return Value
// 
// The resulting point after applying the receiver’s transformations.
//
// See: https://developer.apple.com/documentation/Foundation/NSAffineTransform/transform(_:)-41p16
func (a NSAffineTransform) TransformPoint(aPoint corefoundation.CGPoint) NSPoint {
	rv := objc.Send[NSPoint](a.ID, objc.Sel("transformPoint:"), aPoint)
	return NSPoint(rv)
}

// Applies the receiver’s transform to the specified size and returns the
// results.
//
// aSize: The size data to which you want to apply the matrix.
//
// # Return Value
// 
// The resulting size after applying the receiver’s transformations.
//
// # Discussion
// 
// This method applies the current rotation and scaling factors to `aSize`; it
// does not apply translation factors. You can think of this method as
// transforming a vector whose origin is (0, 0) and whose end point is
// specified by the value in `aSize`. After the rotation and scaling factors
// are applied, this method effectively returns the end point of the new
// vector.
// 
// This method is useful for transforming delta or distance values when you
// need to take scaling and rotation factors into account.
//
// See: https://developer.apple.com/documentation/Foundation/NSAffineTransform/transform(_:)-5r6ol
func (a NSAffineTransform) TransformSize(aSize corefoundation.CGSize) NSSize {
	rv := objc.Send[NSSize](a.ID, objc.Sel("transformSize:"), aSize)
	return NSSize(rv)
}

// Creates and returns a new Bézier path object with each point in the given
// path transformed by the receiver.
//
// path: An object representing the bezier path to be used in the transformation.
//
// path is a [appkit.NSBezierPath].
//
// # Discussion
// 
// The original [NSBezierPath] object is not modified.
//
// See: https://developer.apple.com/documentation/Foundation/NSAffineTransform/transform(_:)-6z1xo
func (a NSAffineTransform) TransformBezierPath(path objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("transformBezierPath:"), path)
	return objectivec.Object{ID: rv}
}

// Sets the current transformation matrix to the receiver’s transformation
// matrix.
//
// # Discussion
// 
// The current transformation is stored in the current graphics context and is
// applied to subsequent drawing operations. You should use this method
// sparingly because it removes the existing transformation matrix, which is
// an accumulation of transformation matrices for the screen, window, and any
// superviews. Instead use the [Concat] method to add this transformation
// matrix to the current transformation matrix.
//
// See: https://developer.apple.com/documentation/Foundation/NSAffineTransform/set()
func (a NSAffineTransform) Set() {
	objc.Send[objc.ID](a.ID, objc.Sel("set"))
}

// Appends the receiver’s matrix to the current transformation matrix stored
// in the current graphics context, replacing the current transformation
// matrix with the result.
//
// # Discussion
// 
// Concatenation is performed by matrix multiplication—see Manipulating
// Transform Values.
// 
// If this method is invoked from within an [NSView][draw(_:)] method, then
// the current transformation matrix is an accumulation of the screen, window,
// and any superview’s transformation matrices. Invoking this method defines
// a new user coordinate system whose coordinates are mapped into the former
// coordinate system according to the receiver’s transformation matrix. To
// undo the concatenation, you must invert the receiver’s matrix and invoke
// this method again.
//
// [draw(_:)]: https://developer.apple.com/documentation/AppKit/NSView/draw(_:)
//
// See: https://developer.apple.com/documentation/Foundation/NSAffineTransform/concat()
func (a NSAffineTransform) Concat() {
	objc.Send[objc.ID](a.ID, objc.Sel("concat"))
}

// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (a NSAffineTransform) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](a.ID, objc.Sel("encodeWithCoder:"), coder)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func (a NSAffineTransform) InitWithCoder(coder INSCoder) NSAffineTransform {
	rv := objc.Send[NSAffineTransform](a.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// Creates a new affine transform initialized to the identity matrix.
//
// # Return Value
// 
// A new identity transform object. This matrix transforms any point to the
// same point.
//
// See: https://developer.apple.com/documentation/Foundation/NSAffineTransform/transform
func (_NSAffineTransformClass NSAffineTransformClass) Transform() NSAffineTransform {
	rv := objc.Send[objc.ID](objc.ID(_NSAffineTransformClass.class), objc.Sel("transform"))
	return NSAffineTransformFromID(rv)
}

// The matrix coefficients stored as the transformation matrix.
//
// # Discussion
// 
// The matrix is of the form shown in [Transform Mathematics], and the
// six-element structure defined by the [NSAffineTransformStruct] structure is
// of the form:
// 
// The [NSAffineTransformStruct] structure is an alternate representation of a
// transformation matrix that can be used to specify matrix values directly.
//
// [NSAffineTransformStruct]: https://developer.apple.com/documentation/Foundation/NSAffineTransformStruct
// [Transform Mathematics]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/CocoaDrawingGuide/Transforms/Transforms.html#//apple_ref/doc/uid/TP40003290-CH204-BCIIICJI
//
// See: https://developer.apple.com/documentation/Foundation/NSAffineTransform/transformStruct
func (a NSAffineTransform) TransformStruct() NSAffineTransformStruct {
	rv := objc.Send[NSAffineTransformStruct](a.ID, objc.Sel("transformStruct"))
	return NSAffineTransformStruct(rv)
}
func (a NSAffineTransform) SetTransformStruct(value NSAffineTransformStruct) {
	objc.Send[struct{}](a.ID, objc.Sel("setTransformStruct:"), value)
}

			// Protocol methods for NSCopying
			

			// Protocol methods for NSSecureCoding
			

