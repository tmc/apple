// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSCustomImageRep] class.
var (
	_NSCustomImageRepClass     NSCustomImageRepClass
	_NSCustomImageRepClassOnce sync.Once
)

func getNSCustomImageRepClass() NSCustomImageRepClass {
	_NSCustomImageRepClassOnce.Do(func() {
		_NSCustomImageRepClass = NSCustomImageRepClass{class: objc.GetClass("NSCustomImageRep")}
	})
	return _NSCustomImageRepClass
}

// GetNSCustomImageRepClass returns the class object for NSCustomImageRep.
func GetNSCustomImageRepClass() NSCustomImageRepClass {
	return getNSCustomImageRepClass()
}

type NSCustomImageRepClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSCustomImageRepClass) Alloc() NSCustomImageRep {
	rv := objc.Send[NSCustomImageRep](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// An object that uses a delegate object to render an image from a custom
// format.
//
// # Overview
// 
// When called upon to produce an image, an [NSCustomImageRep] sends a message
// to its delegate to do the actual drawing. You can use this class to support
// custom image formats without going to the trouble of subclassing
// [NSImageRep] directly.
//
// # Creating Representations of Images in Custom Formats
//
//   - [NSCustomImageRep.InitWithDrawSelectorDelegate]: Returns a representation of an image initialized with the specified delegate information.
//   - [NSCustomImageRep.InitWithSizeFlippedDrawingHandler]: Initializes a representation of an image of the specified size and flipped status, using a block to draw its content.
//
// # Getting Drawing Handlers
//
//   - [NSCustomImageRep.DrawingHandler]: The destination rectangle of the drawing handler block.
//
// # Getting Information About Images
//
//   - [NSCustomImageRep.Delegate]: The delegate object that renders the image for the image representation.
//   - [NSCustomImageRep.DrawSelector]: The selector for the delegate’s drawing method.
//
// See: https://developer.apple.com/documentation/AppKit/NSCustomImageRep
type NSCustomImageRep struct {
	NSImageRep
}

// NSCustomImageRepFromID constructs a [NSCustomImageRep] from an objc.ID.
//
// An object that uses a delegate object to render an image from a custom
// format.
func NSCustomImageRepFromID(id objc.ID) NSCustomImageRep {
	return NSCustomImageRep{NSImageRep: NSImageRepFromID(id)}
}
// NOTE: NSCustomImageRep adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSCustomImageRep] class.
//
// # Creating Representations of Images in Custom Formats
//
//   - [INSCustomImageRep.InitWithDrawSelectorDelegate]: Returns a representation of an image initialized with the specified delegate information.
//   - [INSCustomImageRep.InitWithSizeFlippedDrawingHandler]: Initializes a representation of an image of the specified size and flipped status, using a block to draw its content.
//
// # Getting Drawing Handlers
//
//   - [INSCustomImageRep.DrawingHandler]: The destination rectangle of the drawing handler block.
//
// # Getting Information About Images
//
//   - [INSCustomImageRep.Delegate]: The delegate object that renders the image for the image representation.
//   - [INSCustomImageRep.DrawSelector]: The selector for the delegate’s drawing method.
//
// See: https://developer.apple.com/documentation/AppKit/NSCustomImageRep
type INSCustomImageRep interface {
	INSImageRep

	// Topic: Creating Representations of Images in Custom Formats

	// Returns a representation of an image initialized with the specified delegate information.
	InitWithDrawSelectorDelegate(selector objc.SEL, delegate objectivec.IObject) NSCustomImageRep
	// Initializes a representation of an image of the specified size and flipped status, using a block to draw its content.
	InitWithSizeFlippedDrawingHandler(size corefoundation.CGSize, drawingHandlerShouldBeCalledWithFlippedContext ErrorHandler, drawingHandler ErrorHandler) NSCustomImageRep

	// Topic: Getting Drawing Handlers

	// The destination rectangle of the drawing handler block.
	DrawingHandler() structCGRectHandler

	// Topic: Getting Information About Images

	// The delegate object that renders the image for the image representation.
	Delegate() objectivec.IObject
	// The selector for the delegate’s drawing method.
	DrawSelector() objc.SEL

	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (c NSCustomImageRep) Init() NSCustomImageRep {
	rv := objc.Send[NSCustomImageRep](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSCustomImageRep) Autorelease() NSCustomImageRep {
	rv := objc.Send[NSCustomImageRep](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSCustomImageRep creates a new NSCustomImageRep instance.
func NewNSCustomImageRep() NSCustomImageRep {
	class := getNSCustomImageRepClass()
	rv := objc.Send[NSCustomImageRep](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Creates and returns an image representation object from data in an
// unarchiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSImageRep/init(coder:)
func NewCustomImageRepWithCoder(coder foundation.INSCoder) NSCustomImageRep {
	instance := getNSCustomImageRepClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSCustomImageRepFromID(rv)
}


// Returns a representation of an image initialized with the specified
// delegate information.
//
// selector: The selector to call when it is time to draw the image. The method should
// take a single parameter of type `id` that represents the [NSCustomImageRep]
// object that initiated drawing. The method must draw the image starting at
// the point (0, 0) in the current coordinate system.
//
// delegate: The delegate object that responds to the selector in `aMethod`.
//
// # Return Value
// 
// An initialized [NSCustomImageRep] object, or `nil` if the object could not
// be initialized.
//
// # Discussion
// 
// When the receiver is asked to draw the image, it sends the specified
// message to the selector, passing itself as a parameter to the delegate
// method. The delegate’s drawing method should have the following form:
//
// See: https://developer.apple.com/documentation/AppKit/NSCustomImageRep/init(draw:delegate:)
func NewCustomImageRepWithDrawSelectorDelegate(selector objc.SEL, delegate objectivec.IObject) NSCustomImageRep {
	instance := getNSCustomImageRepClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDrawSelector:delegate:"), selector, delegate)
	return NSCustomImageRepFromID(rv)
}


// Initializes a representation of an image of the specified size and flipped
// status, using a block to draw its content.
//
// size: The size of the image.
//
// drawingHandlerShouldBeCalledWithFlippedContext: [true] if the drawing handler should be called with a flipped graphics
// context; otherwise [false].
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// drawingHandler: A block that draws the image representation content in the provided
// graphics context.
// 
// The block may be invoked whenever and on whatever thread the image itself
// is drawn on. Care should be taken to ensure that all state accessed within
// the drawingHandler block is done so in a thread safe manner.
// 
// This Block replaces the [LockFocus] and [UnlockFocus] technique of creating
// drawing content. The block is invoked at draw time, the drawing can be
// adjusted to suit the destination’s pixel density, color space, and other
// properties.
//
// # Return Value
// 
// An initialized [NSCustomImageRep] object, or `nil` if the object could not
// be initialized.
//
// # Discussion
// 
// Using the this method ensures you’ll get correct results under standard
// and high resolution.
// 
// Like other non-bitmap image rep types, drawing is cached as appropriate for
// the destination context. Practically speaking, the `drawingHandler` block
// will be invoked the first time the image is drawn to a particular type of
// destination (1x or 2x screen, for example). Subsequent drawing operations
// to the same type of destination will reuse the previously generated bitmap.
//
// See: https://developer.apple.com/documentation/AppKit/NSCustomImageRep/init(size:flipped:drawingHandler:)
func NewCustomImageRepWithSizeFlippedDrawingHandler(size corefoundation.CGSize, drawingHandlerShouldBeCalledWithFlippedContext bool, drawingHandler corefoundation.CGRect) NSCustomImageRep {
	instance := getNSCustomImageRepClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSize:flipped:drawingHandler:"), size, drawingHandlerShouldBeCalledWithFlippedContext, drawingHandler)
	return NSCustomImageRepFromID(rv)
}







// Returns a representation of an image initialized with the specified
// delegate information.
//
// selector: The selector to call when it is time to draw the image. The method should
// take a single parameter of type `id` that represents the [NSCustomImageRep]
// object that initiated drawing. The method must draw the image starting at
// the point (0, 0) in the current coordinate system.
//
// delegate: The delegate object that responds to the selector in `aMethod`.
//
// # Return Value
// 
// An initialized [NSCustomImageRep] object, or `nil` if the object could not
// be initialized.
//
// # Discussion
// 
// When the receiver is asked to draw the image, it sends the specified
// message to the selector, passing itself as a parameter to the delegate
// method. The delegate’s drawing method should have the following form:
//
// See: https://developer.apple.com/documentation/AppKit/NSCustomImageRep/init(draw:delegate:)
func (c NSCustomImageRep) InitWithDrawSelectorDelegate(selector objc.SEL, delegate objectivec.IObject) NSCustomImageRep {
	rv := objc.Send[NSCustomImageRep](c.ID, objc.Sel("initWithDrawSelector:delegate:"), selector, delegate)
	return rv
}

// Initializes a representation of an image of the specified size and flipped
// status, using a block to draw its content.
//
// size: The size of the image.
//
// drawingHandlerShouldBeCalledWithFlippedContext: [true] if the drawing handler should be called with a flipped graphics
// context; otherwise [false].
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// drawingHandler: A block that draws the image representation content in the provided
// graphics context.
// 
// The block may be invoked whenever and on whatever thread the image itself
// is drawn on. Care should be taken to ensure that all state accessed within
// the drawingHandler block is done so in a thread safe manner.
// 
// This Block replaces the [LockFocus] and [UnlockFocus] technique of creating
// drawing content. The block is invoked at draw time, the drawing can be
// adjusted to suit the destination’s pixel density, color space, and other
// properties.
//
// # Return Value
// 
// An initialized [NSCustomImageRep] object, or `nil` if the object could not
// be initialized.
//
// # Discussion
// 
// Using the this method ensures you’ll get correct results under standard
// and high resolution.
// 
// Like other non-bitmap image rep types, drawing is cached as appropriate for
// the destination context. Practically speaking, the `drawingHandler` block
// will be invoked the first time the image is drawn to a particular type of
// destination (1x or 2x screen, for example). Subsequent drawing operations
// to the same type of destination will reuse the previously generated bitmap.
//
// See: https://developer.apple.com/documentation/AppKit/NSCustomImageRep/init(size:flipped:drawingHandler:)
func (c NSCustomImageRep) InitWithSizeFlippedDrawingHandler(size corefoundation.CGSize, drawingHandlerShouldBeCalledWithFlippedContext ErrorHandler, drawingHandler ErrorHandler) NSCustomImageRep {
		_block1, _cleanup1 := NewErrorBlock(drawingHandlerShouldBeCalledWithFlippedContext)
	defer _cleanup1()
	_block2, _cleanup2 := NewErrorBlock(drawingHandler)
	defer _cleanup2()
		rv := objc.Send[objc.ID](c.ID, objc.Sel("initWithSize:flipped:drawingHandler:"), size, _block1, _block2)
	return NSCustomImageRepFromID(rv)
}
func (c NSCustomImageRep) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}












// The destination rectangle of the drawing handler block.
//
// See: https://developer.apple.com/documentation/AppKit/NSCustomImageRep/drawingHandler
func (c NSCustomImageRep) DrawingHandler() structCGRectHandler {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("drawingHandler"))
	_ = rv
	return nil
}



// The delegate object that renders the image for the image representation.
//
// See: https://developer.apple.com/documentation/AppKit/NSCustomImageRep/delegate
func (c NSCustomImageRep) Delegate() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("delegate"))
	return objectivec.Object{ID: rv}
}



// The selector for the delegate’s drawing method.
//
// See: https://developer.apple.com/documentation/AppKit/NSCustomImageRep/drawSelector
func (c NSCustomImageRep) DrawSelector() objc.SEL {
	rv := objc.Send[objc.SEL](c.ID, objc.Sel("drawSelector"))
	return rv
}



























