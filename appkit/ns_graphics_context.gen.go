// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/coreimage"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSGraphicsContext] class.
var (
	_NSGraphicsContextClass     NSGraphicsContextClass
	_NSGraphicsContextClassOnce sync.Once
)

func getNSGraphicsContextClass() NSGraphicsContextClass {
	_NSGraphicsContextClassOnce.Do(func() {
		_NSGraphicsContextClass = NSGraphicsContextClass{class: objc.GetClass("NSGraphicsContext")}
	})
	return _NSGraphicsContextClass
}

// GetNSGraphicsContextClass returns the class object for NSGraphicsContext.
func GetNSGraphicsContextClass() NSGraphicsContextClass {
	return getNSGraphicsContextClass()
}

type NSGraphicsContextClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSGraphicsContextClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSGraphicsContextClass) Alloc() NSGraphicsContext {
	rv := objc.Send[NSGraphicsContext](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a graphics context.
//
// # Overview
//
// You can think of a graphics context as a destination to which drawing and
// graphics state operations are sent for execution. Each graphics context
// contains its own graphics environment and state.
//
// The [NSGraphicsContext] class is an abstract superclass for
// destination-specific graphics contexts. You obtain instances of concrete
// subclasses with the class methods [NSGraphicsContext.CurrentContext],
// [NSGraphicsContext.GraphicsContextWithAttributes], [NSGraphicsContext.GraphicsContextWithBitmapImageRep],
// [NSGraphicsContext.GraphicsContextWithCGContextFlipped], and [init(window:)].
//
// At any time there is the notion of the current context. The current context
// for the current thread may be set using [NSGraphicsContext.CurrentContext].
//
// Graphics contexts are maintained on a stack. You push a graphics context
// onto the stack by sending it a [NSGraphicsContext.SaveGraphicsState] message, and pop it off
// the stack by sending it a [NSGraphicsContext.RestoreGraphicsState] message. By sending
// [NSGraphicsContext.RestoreGraphicsState] to a graphics context object you remove it from the
// stack, and the next graphics context on the stack becomes the current
// graphics context.
//
// # Managing the Current Context
//
//   - [NSGraphicsContext.CGContext]: The Core Graphics context, which is a low-level, platform-specific graphics context.
//
// # Testing the Drawing Destination
//
//   - [NSGraphicsContext.DrawingToScreen]: A Boolean value that indicates whether the drawing destination is the screen.
//
// # Getting Information About the Context
//
//   - [NSGraphicsContext.Attributes]: The attributes used to create this instance.
//   - [NSGraphicsContext.Flipped]: A Boolean value that indicates the graphics context’s flipped state.
//
// # Flushing Graphics to the Context
//
//   - [NSGraphicsContext.FlushGraphics]: Forces any buffered operations or data to be sent to the graphics context’s destination.
//
// # Configuring Rendering Options
//
//   - [NSGraphicsContext.CompositingOperation]: The graphics context’s global compositing operation setting.
//   - [NSGraphicsContext.SetCompositingOperation]
//   - [NSGraphicsContext.ImageInterpolation]: A constant that specifies the graphics context’s interpolation, or image smoothing, behavior.
//   - [NSGraphicsContext.SetImageInterpolation]
//   - [NSGraphicsContext.ShouldAntialias]: A Boolean value that indicates whether the graphics context uses antialiasing.
//   - [NSGraphicsContext.SetShouldAntialias]
//   - [NSGraphicsContext.PatternPhase]: The amount to offset the pattern color when filling the graphics context.
//   - [NSGraphicsContext.SetPatternPhase]
//
// # Getting the Context for Rendering Core Image Objects
//
//   - [NSGraphicsContext.CIContext]: A context for Core Image objects that you can use to render into the graphics context.
//
// # Managing Color Rendering
//
//   - [NSGraphicsContext.ColorRenderingIntent]: The color rendering intent in the graphics context’s graphics state.
//   - [NSGraphicsContext.SetColorRenderingIntent]
//
// See: https://developer.apple.com/documentation/AppKit/NSGraphicsContext
//
// [init(window:)]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/init(window:)
type NSGraphicsContext struct {
	objectivec.Object
}

// NSGraphicsContextFromID constructs a [NSGraphicsContext] from an objc.ID.
//
// An object that represents a graphics context.
func NSGraphicsContextFromID(id objc.ID) NSGraphicsContext {
	return NSGraphicsContext{objectivec.Object{ID: id}}
}

// NOTE: NSGraphicsContext adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSGraphicsContext] class.
//
// # Managing the Current Context
//
//   - [INSGraphicsContext.CGContext]: The Core Graphics context, which is a low-level, platform-specific graphics context.
//
// # Testing the Drawing Destination
//
//   - [INSGraphicsContext.DrawingToScreen]: A Boolean value that indicates whether the drawing destination is the screen.
//
// # Getting Information About the Context
//
//   - [INSGraphicsContext.Attributes]: The attributes used to create this instance.
//   - [INSGraphicsContext.Flipped]: A Boolean value that indicates the graphics context’s flipped state.
//
// # Flushing Graphics to the Context
//
//   - [INSGraphicsContext.FlushGraphics]: Forces any buffered operations or data to be sent to the graphics context’s destination.
//
// # Configuring Rendering Options
//
//   - [INSGraphicsContext.CompositingOperation]: The graphics context’s global compositing operation setting.
//   - [INSGraphicsContext.SetCompositingOperation]
//   - [INSGraphicsContext.ImageInterpolation]: A constant that specifies the graphics context’s interpolation, or image smoothing, behavior.
//   - [INSGraphicsContext.SetImageInterpolation]
//   - [INSGraphicsContext.ShouldAntialias]: A Boolean value that indicates whether the graphics context uses antialiasing.
//   - [INSGraphicsContext.SetShouldAntialias]
//   - [INSGraphicsContext.PatternPhase]: The amount to offset the pattern color when filling the graphics context.
//   - [INSGraphicsContext.SetPatternPhase]
//
// # Getting the Context for Rendering Core Image Objects
//
//   - [INSGraphicsContext.CIContext]: A context for Core Image objects that you can use to render into the graphics context.
//
// # Managing Color Rendering
//
//   - [INSGraphicsContext.ColorRenderingIntent]: The color rendering intent in the graphics context’s graphics state.
//   - [INSGraphicsContext.SetColorRenderingIntent]
//
// See: https://developer.apple.com/documentation/AppKit/NSGraphicsContext
type INSGraphicsContext interface {
	objectivec.IObject

	// Topic: Managing the Current Context

	// The Core Graphics context, which is a low-level, platform-specific graphics context.
	CGContext() coregraphics.CGContextRef

	// Topic: Testing the Drawing Destination

	// A Boolean value that indicates whether the drawing destination is the screen.
	DrawingToScreen() bool

	// Topic: Getting Information About the Context

	// The attributes used to create this instance.
	Attributes() foundation.INSDictionary
	// A Boolean value that indicates the graphics context’s flipped state.
	Flipped() bool

	// Topic: Flushing Graphics to the Context

	// Forces any buffered operations or data to be sent to the graphics context’s destination.
	FlushGraphics()

	// Topic: Configuring Rendering Options

	// The graphics context’s global compositing operation setting.
	CompositingOperation() NSCompositingOperation
	SetCompositingOperation(value NSCompositingOperation)
	// A constant that specifies the graphics context’s interpolation, or image smoothing, behavior.
	ImageInterpolation() NSImageInterpolation
	SetImageInterpolation(value NSImageInterpolation)
	// A Boolean value that indicates whether the graphics context uses antialiasing.
	ShouldAntialias() bool
	SetShouldAntialias(value bool)
	// The amount to offset the pattern color when filling the graphics context.
	PatternPhase() corefoundation.CGPoint
	SetPatternPhase(value corefoundation.CGPoint)

	// Topic: Getting the Context for Rendering Core Image Objects

	// A context for Core Image objects that you can use to render into the graphics context.
	CIContext() coreimage.CIContext

	// Topic: Managing Color Rendering

	// The color rendering intent in the graphics context’s graphics state.
	ColorRenderingIntent() NSColorRenderingIntent
	SetColorRenderingIntent(value NSColorRenderingIntent)
}

// Init initializes the instance.
func (g NSGraphicsContext) Init() NSGraphicsContext {
	rv := objc.Send[NSGraphicsContext](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g NSGraphicsContext) Autorelease() NSGraphicsContext {
	rv := objc.Send[NSGraphicsContext](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSGraphicsContext creates a new NSGraphicsContext instance.
func NewNSGraphicsContext() NSGraphicsContext {
	class := getNSGraphicsContextClass()
	rv := objc.Send[NSGraphicsContext](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a graphics context using the specified attributes.
//
// attributes: A dictionary of values associated with the keys described in
// [NSGraphicsContextAttributeKey]. The attributes specify such things as
// representation format and destination.
//
// # Return Value
//
// A new [NSGraphicsContext] object, or `nil` if the object could not be
// created.
//
// # Discussion
//
// Use this method to create a graphics context for a window or bitmap
// destination. If you want to create a graphics context for a PDF or
// PostScript destination, do not use this method; instead, use the
// [NSPrintOperation] class to set up the printing environment needed to
// generate that type of information.
//
// See: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/init(attributes:)
func NewGraphicsContextWithAttributes(attributes foundation.INSDictionary) NSGraphicsContext {
	rv := objc.Send[objc.ID](objc.ID(getNSGraphicsContextClass().class), objc.Sel("graphicsContextWithAttributes:"), attributes)
	return NSGraphicsContextFromID(rv)
}

// Creates a new graphics context using the specified bitmap image
// representation object as the context destination.
//
// bitmapRep: The [NSBitmapImageRep] object to use as the destination.
//
// # Return Value
//
// The created [NSGraphicsContext] object, or `nil` if the object could not be
// created.
//
// # Discussion
//
// This method accepts only single plane [NSBitmapImageRep] instances. It is
// the equivalent of using [GraphicsContextWithAttributes] and passing
// `bitmapRep` as the value for the dictionary’s [destination] key.
//
// See: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/init(bitmapImageRep:)
//
// [destination]: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/AttributeKey/destination
func NewGraphicsContextWithBitmapImageRep(bitmapRep INSBitmapImageRep) NSGraphicsContext {
	rv := objc.Send[objc.ID](objc.ID(getNSGraphicsContextClass().class), objc.Sel("graphicsContextWithBitmapImageRep:"), bitmapRep)
	return NSGraphicsContextFromID(rv)
}

// Creates a new graphics context from the specified Core Graphics context and
// the initial flipped state.
//
// graphicsPort: The graphics port used to create the graphics-context object, as a
// [CGContext] (opaque type) object.
//
// initialFlippedState: Specifies the receiver’s initial flipped state. This is the value
// returned by [Flipped] when no view has focus.
//
// # Return Value
//
// The created [NSGraphicsContext] object, or `nil` if the object could not be
// created.
//
// See: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/init(cgContext:flipped:)
//
// [CGContext]: https://developer.apple.com/documentation/CoreGraphics/CGContext
func NewGraphicsContextWithCGContextFlipped(graphicsPort coregraphics.CGContextRef, initialFlippedState bool) NSGraphicsContext {
	rv := objc.Send[objc.ID](objc.ID(getNSGraphicsContextClass().class), objc.Sel("graphicsContextWithCGContext:flipped:"), graphicsPort, initialFlippedState)
	return NSGraphicsContextFromID(rv)
}

// Forces any buffered operations or data to be sent to the graphics
// context’s destination.
//
// # Discussion
//
// Graphics contexts use buffers to queue pending operations but for
// efficiency reasons may not always empty those buffers immediately. This
// method forces the buffers to be emptied.
//
// See: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/flushGraphics()
func (g NSGraphicsContext) FlushGraphics() {
	objc.Send[objc.ID](g.ID, objc.Sel("flushGraphics"))
}

// Pops a graphics context from the per-thread stack, makes it current, and
// sends the context a restore graphics state message.
//
// See: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/restoreGraphicsState()-swift.type.method
func (_NSGraphicsContextClass NSGraphicsContextClass) RestoreGraphicsState() {
	objc.Send[objc.ID](objc.ID(_NSGraphicsContextClass.class), objc.Sel("restoreGraphicsState"))
}

// Saves the graphics state of the current graphics context.
//
// # Discussion
//
// This method sends the current graphics context a [SaveGraphicsState]
// message and pushes the context onto the per-thread stack.
//
// See: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/saveGraphicsState()-swift.type.method
func (_NSGraphicsContextClass NSGraphicsContextClass) SaveGraphicsState() {
	objc.Send[objc.ID](objc.ID(_NSGraphicsContextClass.class), objc.Sel("saveGraphicsState"))
}

// Returns a Boolean value that indicates whether the current context is
// drawing to the screen.
//
// # Return Value
//
// true if the current context is drawing to the screen, otherwise false.
//
// # Discussion
//
// This convenience method is equivalent to sending [DrawingToScreen] to the
// result of [CurrentContext].
//
// See: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/currentContextDrawingToScreen()
func (_NSGraphicsContextClass NSGraphicsContextClass) CurrentContextDrawingToScreen() bool {
	rv := objc.Send[bool](objc.ID(_NSGraphicsContextClass.class), objc.Sel("currentContextDrawingToScreen"))
	return rv
}

// The Core Graphics context, which is a low-level, platform-specific graphics
// context.
//
// See: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/cgContext
func (g NSGraphicsContext) CGContext() coregraphics.CGContextRef {
	rv := objc.Send[coregraphics.CGContextRef](g.ID, objc.Sel("CGContext"))
	return coregraphics.CGContextRef(rv)
}

// A Boolean value that indicates whether the drawing destination is the
// screen.
//
// # Discussion
//
// true if the drawing destination is the screen. If the value of the property
// is false may mean that the drawing destination is a printer, but the
// destination may also be a PDF or EPS file. You can call [Attributes] to see
// if additional information is available about the drawing destination.
//
// See: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/isDrawingToScreen
func (g NSGraphicsContext) DrawingToScreen() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("isDrawingToScreen"))
	return rv
}

// The attributes used to create this instance.
//
// # Discussion
//
// Screen-based graphics contexts do not store attributes, even if you create
// them using [GraphicsContextWithAttributes].
//
// See: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/attributes
func (g NSGraphicsContext) Attributes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("attributes"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// A Boolean value that indicates the graphics context’s flipped state.
//
// # Discussion
//
// The state is determined by sending `flipped` to the receiver’s view that
// has focus. If no view has focus, returns false unless the receiver is
// instantiated using [GraphicsContextWithCGContextFlipped] specifying true as
// the `flipped` parameter.
//
// See: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/isFlipped
func (g NSGraphicsContext) Flipped() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("isFlipped"))
	return rv
}

// The graphics context’s global compositing operation setting.
//
// # Discussion
//
// The compositing operation is a global attribute of the graphics context and
// affects drawing operations that do not take an explicit compositing
// operation parameter. For methods that do take an explicit compositing
// operation parameter, the value of that parameter supersedes the global
// value. The compositing operations are related to (but different from) the
// blend mode settings used in Quartz. Only the default compositing operation
// ([NSCompositeCopy]) is supported when rendering PDF or PostScript content.
// See [NSCompositingOperation] for valid constants.
//
// See: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/compositingOperation
//
// [NSCompositeCopy]: https://developer.apple.com/documentation/AppKit/NSCompositeCopy
// [NSCompositingOperation]: https://developer.apple.com/documentation/AppKit/NSCompositingOperation
func (g NSGraphicsContext) CompositingOperation() NSCompositingOperation {
	rv := objc.Send[NSCompositingOperation](g.ID, objc.Sel("compositingOperation"))
	return NSCompositingOperation(rv)
}
func (g NSGraphicsContext) SetCompositingOperation(value NSCompositingOperation) {
	objc.Send[struct{}](g.ID, objc.Sel("setCompositingOperation:"), value)
}

// A constant that specifies the graphics context’s interpolation, or image
// smoothing, behavior.
//
// # Discussion
//
// Note that this value is not part of the graphics state, so it cannot be
// reset using [RestoreGraphicsState].
//
// See: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/imageInterpolation
func (g NSGraphicsContext) ImageInterpolation() NSImageInterpolation {
	rv := objc.Send[NSImageInterpolation](g.ID, objc.Sel("imageInterpolation"))
	return NSImageInterpolation(rv)
}
func (g NSGraphicsContext) SetImageInterpolation(value NSImageInterpolation) {
	objc.Send[struct{}](g.ID, objc.Sel("setImageInterpolation:"), value)
}

// A Boolean value that indicates whether the graphics context uses
// antialiasing.
//
// # Discussion
//
// true if the receiver uses antialiasing. This value is part of the graphics
// state and is restored by [RestoreGraphicsState].
//
// See: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/shouldAntialias
func (g NSGraphicsContext) ShouldAntialias() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("shouldAntialias"))
	return rv
}
func (g NSGraphicsContext) SetShouldAntialias(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setShouldAntialias:"), value)
}

// The amount to offset the pattern color when filling the graphics context.
//
// # Discussion
//
// Use this property when you need to line up the pattern color with another
// pattern, such as the pattern in a superview. The pattern phase is a
// translation (width, height) applied before a pattern is drawn in the
// current context and is part of the saved graphics state of the context. The
// default pattern phase is (0,0). Setting the pattern phase has the effect of
// temporarily changing the pattern matrix of any pattern you decide to draw.
// For example, setting the pattern phase to (2,3) has the effect of moving
// the start of pattern cell tiling to the point (2,3) in default user space.
//
// See: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/patternPhase
func (g NSGraphicsContext) PatternPhase() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](g.ID, objc.Sel("patternPhase"))
	return corefoundation.CGPoint(rv)
}
func (g NSGraphicsContext) SetPatternPhase(value corefoundation.CGPoint) {
	objc.Send[struct{}](g.ID, objc.Sel("setPatternPhase:"), value)
}

// A context for Core Image objects that you can use to render into the
// graphics context.
//
// # Discussion
//
// The [CIContext] object is created on demand and remains in existence for
// the lifetime of its owning [NSGraphicsContext] object. A [CIContext] object
// is an evaluation context for rendering a [CIImage] object through Quartz 2D
// or OpenGL. You use [CIContext]` `objects in conjunction with [CIFilter],
// [CIImage], [CIVector], and [CIColor] objects to take advantage of the
// built-in Core Image filters when processing images.
//
// For more on [CIContext] objects and related Core Image objects, see [Core
// Image Programming Guide].
//
// See: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/ciContext
//
// [CIColor]: https://developer.apple.com/documentation/CoreImage/CIColor
// [CIContext]: https://developer.apple.com/documentation/CoreImage/CIContext
// [CIFilter]: https://developer.apple.com/documentation/CoreImage/CIFilter-swift.class
// [CIImage]: https://developer.apple.com/documentation/CoreImage/CIImage
// [CIVector]: https://developer.apple.com/documentation/CoreImage/CIVector
// [Core Image Programming Guide]: https://developer.apple.com/library/archive/documentation/GraphicsImaging/Conceptual/CoreImaging/ci_intro/ci_intro.html#//apple_ref/doc/uid/TP30001185
func (g NSGraphicsContext) CIContext() coreimage.CIContext {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("CIContext"))
	return coreimage.CIContextFromID(objc.ID(rv))
}

// The color rendering intent in the graphics context’s graphics state.
//
// # Discussion
//
// A value that specifies the rendering intent currently used by the graphics
// context. For possible values, see [NSColorRenderingIntent]. The rendering
// intent specifies how Cocoa should handle colors that are not located within
// the gamut of the destination color space of a graphics context. If you do
// not explicitly set the rendering intent, and sampled images are being
// drawn, [NSGraphicsContext] uses perceptual rendering intent. Otherwise,
// [NSGraphicsContext] uses relative colorimetric rendering intent.
//
// See: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/colorRenderingIntent
//
// [NSColorRenderingIntent]: https://developer.apple.com/documentation/AppKit/NSColorRenderingIntent
func (g NSGraphicsContext) ColorRenderingIntent() NSColorRenderingIntent {
	rv := objc.Send[NSColorRenderingIntent](g.ID, objc.Sel("colorRenderingIntent"))
	return NSColorRenderingIntent(rv)
}
func (g NSGraphicsContext) SetColorRenderingIntent(value NSColorRenderingIntent) {
	objc.Send[struct{}](g.ID, objc.Sel("setColorRenderingIntent:"), value)
}

// Returns the current graphics context of the current thread.
//
// # Return Value
//
// The current graphics context of the current thread.
//
// # Discussion
//
// Returns an instance of a concrete subclass of [NSGraphicsContext].
//
// See: https://developer.apple.com/documentation/AppKit/NSGraphicsContext/current
func (_NSGraphicsContextClass NSGraphicsContextClass) CurrentContext() NSGraphicsContext {
	rv := objc.Send[objc.ID](objc.ID(_NSGraphicsContextClass.class), objc.Sel("currentContext"))
	return NSGraphicsContextFromID(objc.ID(rv))
}
func (_NSGraphicsContextClass NSGraphicsContextClass) SetCurrentContext(value NSGraphicsContext) {
	objc.Send[struct{}](objc.ID(_NSGraphicsContextClass.class), objc.Sel("setCurrentContext:"), value)
}
