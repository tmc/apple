// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSStringDrawingContext] class.
var (
	_NSStringDrawingContextClass     NSStringDrawingContextClass
	_NSStringDrawingContextClassOnce sync.Once
)

func getNSStringDrawingContextClass() NSStringDrawingContextClass {
	_NSStringDrawingContextClassOnce.Do(func() {
		_NSStringDrawingContextClass = NSStringDrawingContextClass{class: objc.GetClass("NSStringDrawingContext")}
	})
	return _NSStringDrawingContextClass
}

// GetNSStringDrawingContextClass returns the class object for NSStringDrawingContext.
func GetNSStringDrawingContextClass() NSStringDrawingContextClass {
	return getNSStringDrawingContextClass()
}

type NSStringDrawingContextClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSStringDrawingContextClass) Alloc() NSStringDrawingContext {
	rv := objc.Send[NSStringDrawingContext](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// An object that manages metrics for drawing attributed strings.
//
// # Overview
// 
// Prior to drawing, you can create an instance of this class and use it to
// specify the minimum scale factor and tracking adjustments for a string.
// After drawing, you can retrieve the actual values that were used during
// drawing.
// 
// To use this class, allocate and initialize a new instance, set the minimum
// values, and pass your object to one of the corresponding
// [NSAttributedString] methods that take the context object as a parameter.
// Upon completion of drawing, you can use the actual drawing values to make
// adjustments or record where the string was actually drawn.
//
// [NSAttributedString]: https://developer.apple.com/documentation/Foundation/NSAttributedString
//
// # Accessing the scale factors
//
//   - [NSStringDrawingContext.MinimumScaleFactor]: The scale factor that determines the smallest font size to use during drawing.
//   - [NSStringDrawingContext.SetMinimumScaleFactor]
//   - [NSStringDrawingContext.ActualScaleFactor]: The actual scale factor that the system applied to the font during drawing.
//
// # Getting the drawing bounds
//
//   - [NSStringDrawingContext.TotalBounds]: The most recent bounding rectangle that the system used to draw the string.
//
// See: https://developer.apple.com/documentation/AppKit/NSStringDrawingContext
type NSStringDrawingContext struct {
	objectivec.Object
}

// NSStringDrawingContextFromID constructs a [NSStringDrawingContext] from an objc.ID.
//
// An object that manages metrics for drawing attributed strings.
func NSStringDrawingContextFromID(id objc.ID) NSStringDrawingContext {
	return NSStringDrawingContext{objectivec.Object{id}}
}
// NOTE: NSStringDrawingContext adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSStringDrawingContext] class.
//
// # Accessing the scale factors
//
//   - [INSStringDrawingContext.MinimumScaleFactor]: The scale factor that determines the smallest font size to use during drawing.
//   - [INSStringDrawingContext.SetMinimumScaleFactor]
//   - [INSStringDrawingContext.ActualScaleFactor]: The actual scale factor that the system applied to the font during drawing.
//
// # Getting the drawing bounds
//
//   - [INSStringDrawingContext.TotalBounds]: The most recent bounding rectangle that the system used to draw the string.
//
// See: https://developer.apple.com/documentation/AppKit/NSStringDrawingContext
type INSStringDrawingContext interface {
	objectivec.IObject

	// Topic: Accessing the scale factors

	// The scale factor that determines the smallest font size to use during drawing.
	MinimumScaleFactor() float64
	SetMinimumScaleFactor(value float64)
	// The actual scale factor that the system applied to the font during drawing.
	ActualScaleFactor() float64

	// Topic: Getting the drawing bounds

	// The most recent bounding rectangle that the system used to draw the string.
	TotalBounds() corefoundation.CGRect
}





// Init initializes the instance.
func (s NSStringDrawingContext) Init() NSStringDrawingContext {
	rv := objc.Send[NSStringDrawingContext](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSStringDrawingContext) Autorelease() NSStringDrawingContext {
	rv := objc.Send[NSStringDrawingContext](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSStringDrawingContext creates a new NSStringDrawingContext instance.
func NewNSStringDrawingContext() NSStringDrawingContext {
	class := getNSStringDrawingContextClass()
	rv := objc.Send[NSStringDrawingContext](objc.ID(class.class), objc.Sel("new"))
	return rv
}





















// The scale factor that determines the smallest font size to use during
// drawing.
//
// # Discussion
// 
// A value of `0.0` corresponds to a scale factor of `1.0`. Any value greater
// than `0.0` is multiplied by the font point size to get the smallest font
// size that is permissible to use. For example, 0.5 indicates a font that is
// half the size of the actual font, 0.75 is three-quarters of the font size,
// and so on. Typically, you specify a value between 0.0 and 1.0 to indicate
// how much the font can be shrunk during drawing.
// 
// The default value of this property is `0.0`.
//
// See: https://developer.apple.com/documentation/AppKit/NSStringDrawingContext/minimumScaleFactor
func (s NSStringDrawingContext) MinimumScaleFactor() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("minimumScaleFactor"))
	return rv
}
func (s NSStringDrawingContext) SetMinimumScaleFactor(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setMinimumScaleFactor:"), value)
}



// The actual scale factor that the system applied to the font during drawing.
//
// # Discussion
// 
// If you specified a custom value in the [MinimumScaleFactor] property, when
// drawing is complete, this property contains the actual scale factor value
// that was used to draw the string.
//
// See: https://developer.apple.com/documentation/AppKit/NSStringDrawingContext/actualScaleFactor
func (s NSStringDrawingContext) ActualScaleFactor() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("actualScaleFactor"))
	return rv
}



// The most recent bounding rectangle that the system used to draw the string.
//
// # Discussion
// 
// This property contains the bounding rectangle that was last used when
// calling the [draw(with:options:context:)] method. The rectangle is
// specified in the coordinate system of the drawn string. (The origin of the
// bounds corresponds to neither a view the string might have been drawn into
// nor the origin of a possible [draw(in:)] call.)
//
// [draw(in:)]: https://developer.apple.com/documentation/Foundation/NSAttributedString/draw(in:)
// [draw(with:options:context:)]: https://developer.apple.com/documentation/Foundation/NSAttributedString/draw(with:options:context:)
//
// See: https://developer.apple.com/documentation/AppKit/NSStringDrawingContext/totalBounds
func (s NSStringDrawingContext) TotalBounds() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s.ID, objc.Sel("totalBounds"))
	return corefoundation.CGRect(rv)
}























