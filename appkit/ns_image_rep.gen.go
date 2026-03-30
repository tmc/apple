// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSImageRep] class.
var (
	_NSImageRepClass     NSImageRepClass
	_NSImageRepClassOnce sync.Once
)

func getNSImageRepClass() NSImageRepClass {
	_NSImageRepClassOnce.Do(func() {
		_NSImageRepClass = NSImageRepClass{class: objc.GetClass("NSImageRep")}
	})
	return _NSImageRepClass
}

// GetNSImageRepClass returns the class object for NSImageRep.
func GetNSImageRepClass() NSImageRepClass {
	return getNSImageRepClass()
}

type NSImageRepClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSImageRepClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSImageRepClass) Alloc() NSImageRep {
	rv := objc.Send[NSImageRep](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A semiabstract superclass that provides subclasses that you use to draw an
// image from a particular type of source data.
//
// # Overview
//
// The [NSImageRep] class is called “semiabstract” because it has some
// instance variables and implementation of its own, in addition to defining
// subclasses. Although an [NSImageRep] subclass can be used directly, it is
// typically accessed through an [NSImage] object, which manages a group of
// image representations, choosing the best one for the current output device.
//
// # Creating Representations of Images
//
//   - [NSImageRep.InitWithCoder]: Creates and returns an image representation object from data in an unarchiver.
//
// # Accessing Size of Images
//
//   - [NSImageRep.Size]: The size of the image representation, measured in points in the user coordinate space.
//   - [NSImageRep.SetSize]
//
// # Specifying Information About the Representation
//
//   - [NSImageRep.BitsPerSample]: The number of bits per sample in the object (if the object is a planar image, this property contains the number of bits per sample per plane).
//   - [NSImageRep.SetBitsPerSample]
//   - [NSImageRep.ColorSpaceName]: The name of the color space used by the image data.
//   - [NSImageRep.SetColorSpaceName]
//   - [NSImageRep.Alpha]: A Boolean value that indicates whether the image data has an alpha channel.
//   - [NSImageRep.SetAlpha]
//   - [NSImageRep.Opaque]: A Boolean value that indicates whether the image is opaque.
//   - [NSImageRep.SetOpaque]
//   - [NSImageRep.PixelsHigh]: The height of the image, measured in pixels.
//   - [NSImageRep.SetPixelsHigh]
//   - [NSImageRep.PixelsWide]: The width of the image, measured in pixels.
//   - [NSImageRep.SetPixelsWide]
//   - [NSImageRep.LayoutDirection]: The layout direction for the image.
//   - [NSImageRep.SetLayoutDirection]
//
// # Getting Core Graphics Images
//
//   - [NSImageRep.CGImageForProposedRectContextHints]: Returns a Core Graphics image object that captures the drawing of the image.
//
// # Drawing Images
//
//   - [NSImageRep.Draw]: Implemented by subclasses to draw the image in the current coordinate system.
//   - [NSImageRep.DrawAtPoint]: Draws the image representation’s image data at the specified point in the current coordinate system.
//   - [NSImageRep.DrawInRect]: Draws the image, scaling it (as needed) to fit the specified rectangle.
//   - [NSImageRep.DrawInRectFromRectOperationFractionRespectFlippedHints]: Draws all or part of the image in the specified rectangle in the current coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSImageRep
type NSImageRep struct {
	objectivec.Object
}

// NSImageRepFromID constructs a [NSImageRep] from an objc.ID.
//
// A semiabstract superclass that provides subclasses that you use to draw an
// image from a particular type of source data.
func NSImageRepFromID(id objc.ID) NSImageRep {
	return NSImageRep{objectivec.Object{ID: id}}
}

// NOTE: NSImageRep adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSImageRep] class.
//
// # Creating Representations of Images
//
//   - [INSImageRep.InitWithCoder]: Creates and returns an image representation object from data in an unarchiver.
//
// # Accessing Size of Images
//
//   - [INSImageRep.Size]: The size of the image representation, measured in points in the user coordinate space.
//   - [INSImageRep.SetSize]
//
// # Specifying Information About the Representation
//
//   - [INSImageRep.BitsPerSample]: The number of bits per sample in the object (if the object is a planar image, this property contains the number of bits per sample per plane).
//   - [INSImageRep.SetBitsPerSample]
//   - [INSImageRep.ColorSpaceName]: The name of the color space used by the image data.
//   - [INSImageRep.SetColorSpaceName]
//   - [INSImageRep.Alpha]: A Boolean value that indicates whether the image data has an alpha channel.
//   - [INSImageRep.SetAlpha]
//   - [INSImageRep.Opaque]: A Boolean value that indicates whether the image is opaque.
//   - [INSImageRep.SetOpaque]
//   - [INSImageRep.PixelsHigh]: The height of the image, measured in pixels.
//   - [INSImageRep.SetPixelsHigh]
//   - [INSImageRep.PixelsWide]: The width of the image, measured in pixels.
//   - [INSImageRep.SetPixelsWide]
//   - [INSImageRep.LayoutDirection]: The layout direction for the image.
//   - [INSImageRep.SetLayoutDirection]
//
// # Getting Core Graphics Images
//
//   - [INSImageRep.CGImageForProposedRectContextHints]: Returns a Core Graphics image object that captures the drawing of the image.
//
// # Drawing Images
//
//   - [INSImageRep.Draw]: Implemented by subclasses to draw the image in the current coordinate system.
//   - [INSImageRep.DrawAtPoint]: Draws the image representation’s image data at the specified point in the current coordinate system.
//   - [INSImageRep.DrawInRect]: Draws the image, scaling it (as needed) to fit the specified rectangle.
//   - [INSImageRep.DrawInRectFromRectOperationFractionRespectFlippedHints]: Draws all or part of the image in the specified rectangle in the current coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSImageRep
type INSImageRep interface {
	objectivec.IObject

	// Topic: Creating Representations of Images

	// Creates and returns an image representation object from data in an unarchiver.
	InitWithCoder(coder foundation.INSCoder) NSImageRep

	// Topic: Accessing Size of Images

	// The size of the image representation, measured in points in the user coordinate space.
	Size() corefoundation.CGSize
	SetSize(value corefoundation.CGSize)

	// Topic: Specifying Information About the Representation

	// The number of bits per sample in the object (if the object is a planar image, this property contains the number of bits per sample per plane).
	BitsPerSample() int
	SetBitsPerSample(value int)
	// The name of the color space used by the image data.
	ColorSpaceName() NSColorSpaceName
	SetColorSpaceName(value NSColorSpaceName)
	// A Boolean value that indicates whether the image data has an alpha channel.
	Alpha() bool
	SetAlpha(value bool)
	// A Boolean value that indicates whether the image is opaque.
	Opaque() bool
	SetOpaque(value bool)
	// The height of the image, measured in pixels.
	PixelsHigh() int
	SetPixelsHigh(value int)
	// The width of the image, measured in pixels.
	PixelsWide() int
	SetPixelsWide(value int)
	// The layout direction for the image.
	LayoutDirection() NSImageLayoutDirection
	SetLayoutDirection(value NSImageLayoutDirection)

	// Topic: Getting Core Graphics Images

	// Returns a Core Graphics image object that captures the drawing of the image.
	CGImageForProposedRectContextHints(proposedDestRect *corefoundation.CGRect, context INSGraphicsContext, hints foundation.INSDictionary) coregraphics.CGImageRef

	// Topic: Drawing Images

	// Implemented by subclasses to draw the image in the current coordinate system.
	Draw() bool
	// Draws the image representation’s image data at the specified point in the current coordinate system.
	DrawAtPoint(point corefoundation.CGPoint) bool
	// Draws the image, scaling it (as needed) to fit the specified rectangle.
	DrawInRect(rect corefoundation.CGRect) bool
	// Draws all or part of the image in the specified rectangle in the current coordinate system.
	DrawInRectFromRectOperationFractionRespectFlippedHints(dstSpacePortionRect corefoundation.CGRect, srcSpacePortionRect corefoundation.CGRect, op NSCompositingOperation, requestedAlpha float64, respectContextIsFlipped bool, hints foundation.INSDictionary) bool

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (i NSImageRep) Init() NSImageRep {
	rv := objc.Send[NSImageRep](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i NSImageRep) Autorelease() NSImageRep {
	rv := objc.Send[NSImageRep](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSImageRep creates a new NSImageRep instance.
func NewNSImageRep() NSImageRep {
	class := getNSImageRepClass()
	rv := objc.Send[NSImageRep](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates and returns an image representation object from data in an
// unarchiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSImageRep/init(coder:)
func NewImageRepWithCoder(coder foundation.INSCoder) NSImageRep {
	instance := getNSImageRepClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSImageRepFromID(rv)
}

// Creates and returns an image representation object using the contents of
// the specified file.
//
// filename: A full or relative pathname specifying the file to open. This string should
// include the filename extension.
//
// # Return Value
//
// An initialized instance of an [NSImageRep] subclass, or `nil` if the image
// data could not be read.
//
// # Discussion
//
// If sent to the [NSImageRep] class object, this method returns a newly
// allocated instance of a subclass of [NSImageRep] (chosen through the use of
// [class(forFileType:)]) initialized with the contents of the specified file.
// If sent to a subclass of [NSImageRep] that recognizes the type of data in
// the file, it returns an instance of that subclass initialized with the
// contents of the file.
//
// This method returns `nil` in any of the following cases:
//
// - The message is sent to the [NSImageRep] class object and there are no
// subclasses in the [NSImageRep] class registry that handle the type of data
// in the specified file. - The message is sent to a subclass of [NSImageRep]
// and that subclass cannot handle the type of data in the specified file. -
// The [NSImageRep] subclass is unable to initialize itself with the contents
// of the specified file.
//
// The [NSImageRep] subclass is initialized by creating an [NSData] object
// based on the contents of the file and passing it to the “ method. By
// default, the files handled include those with the extensions “`tiff`”,
// “`gif`”, “`jpg`”, “`pict`”, “`pdf`”, and “`eps`”.
//
// See: https://developer.apple.com/documentation/AppKit/NSImageRep/init(contentsOfFile:)
//
// [class(forFileType:)]: https://developer.apple.com/documentation/AppKit/NSImageRep/class(forFileType:)
func NewImageRepWithContentsOfFile(filename string) NSImageRep {
	rv := objc.Send[objc.ID](objc.ID(getNSImageRepClass().class), objc.Sel("imageRepWithContentsOfFile:"), objc.String(filename))
	return NSImageRepFromID(rv)
}

// Creates and returns an image representation object using the data at the
// specified URL.
//
// url: The URL pointing to the image data.
//
// # Return Value
//
// An initialized instance of an [NSImageRep] subclass, or `nil` if the image
// data could not be read.
//
// # Discussion
//
// If sent to the [NSImageRep] class object, this method returns a newly
// allocated instance of a subclass of [NSImageRep] initialized with the
// contents of the specified URL. If sent to a subclass of [NSImageRep] that
// recognizes the data contained in the URL, it returns an instance of that
// subclass initialized with the data in the URL.
//
// This method returns `nil` in any of the following cases:
//
// - The message is sent to the [NSImageRep] class object and there are no
// subclasses in the [NSImageRep] class registry that handle the data
// contained in the specified URL. - The message is sent to a subclass of
// [NSImageRep] and that subclass cannot handle the data contained in the
// specified URL. - The [NSImageRep] subclass is unable to initialize itself
// with the contents of the specified URL.
//
// The [NSImageRep] subclass is initialized by creating an [NSData] object
// based on the contents of the file, then passing it to the “ method.
//
// See: https://developer.apple.com/documentation/AppKit/NSImageRep/init(contentsOf:)
func NewImageRepWithContentsOfURL(url foundation.INSURL) NSImageRep {
	rv := objc.Send[objc.ID](objc.ID(getNSImageRepClass().class), objc.Sel("imageRepWithContentsOfURL:"), url)
	return NSImageRepFromID(rv)
}

// Creates and returns an image representation object using the contents of
// the specified pasteboard.
//
// pasteboard: The pasteboard containing the image data.
//
// # Return Value
//
// An initialized instance of an [NSImageRep] subclass, or `nil` if the image
// data could not be read.
//
// # Discussion
//
// If sent to the [NSImageRep] class object, this method returns a newly
// allocated instance of a subclass of [NSImageRep] initialized with the data
// in the specified pasteboard. If sent to a subclass of [NSImageRep] that
// recognizes the data on the pasteboard, it returns an instance of that
// subclass initialized with that data.
//
// This method returns `nil` in any of the following cases:
//
// - The message is sent to the [NSImageRep] class object and there are no
// subclasses in the [NSImageRep] class registry that handle data of the type
// contained in the specified pasteboard. - The message is sent to a subclass
// of [NSImageRep] and that subclass cannot handle data of the type contained
// in the specified pasteboard. - The [NSImageRep] subclass is unable to
// initialize itself with the contents of the pasteboard.
//
// The [NSImageRep] subclass is initialized by creating an [NSData] object
// based on the data the specified pasteboard and passing it to the “ method.
//
// See: https://developer.apple.com/documentation/AppKit/NSImageRep/init(pasteboard:)
func NewImageRepWithPasteboard(pasteboard INSPasteboard) NSImageRep {
	rv := objc.Send[objc.ID](objc.ID(getNSImageRepClass().class), objc.Sel("imageRepWithPasteboard:"), pasteboard)
	return NSImageRepFromID(rv)
}

// Creates and returns an image representation object from data in an
// unarchiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSImageRep/init(coder:)
func (i NSImageRep) InitWithCoder(coder foundation.INSCoder) NSImageRep {
	rv := objc.Send[NSImageRep](i.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// Returns a Core Graphics image object that captures the drawing of the
// image.
//
// proposedDestRect: On input, the proposed destination rectangle for drawing the image. If
// `nil`, it defaults to the smallest pixel-integral rectangle containing
// `{{0,0}, self.Size()}`. The `proposedDestRect` is in user space in the
// reference context.
//
// On output, the `proposedDestRect` may have been altered. This is because a
// [CGImage] is necessarily pixel-integral, while an [NSImage] is not. In
// order to produce a [CGImage] for rect `(0.5, 0.5, 4.0, 4.0)` without
// distortion or double-antialiasing, we may have to produce a 5x5 [CGImage],
// and also inflate the `proposedDestRect`. Drawing the [CGImage] in the
// out-value `proposedDestRect` is the same as drawing the [NSImage] in the
// in-value of proposed rect.
//
// context: A graphics context. Can be `nil`.
//
// hints: An optional dictionary of hints that provide more context for selecting or
// generating the image. See [NSImageHintKey] for a summary of the possible
// key-value pairs.
//
// # Return Value
//
// A [CGImage]. This may be an existing [CGImage] if one is available. If not,
// a new [CGImage] is created.
//
// # Discussion
//
// An [NSImage] is potentially resolution independent, and may have
// representations that allow it to draw well in many contexts. A [CGImage] is
// more like a single pixel-based representation. This method produces a
// snapshot of how the [NSImage] would draw if it was asked to draw in the
// proposed rectangle in the graphics context.
//
// All input parameters are optional. They provide hints for how to choose
// among existing [CGImage] objects, or how to create one if there isn’t
// already a [CGImage] available. The parameters are only hints.
//
// This method is intended as an override point for image representation
// subclasses that naturally have a [CGImage] available. For example,
// [NSBitmapImageRep] overrides it to return the [CGImage] that naturally
// backs the representation. You don’t need to override the method except
// possibly for performance, though. The [NSImageRep]-level implementation
// will produce a [CGImage] by making a buffer and calling [Draw]. That’s
// likely to be the best possible implementation for reps that aren’t
// naturally [CGImage]-backed. The [Draw] method remains the only method of
// [NSImageRep] that a subclasser really needs to override.
//
// See: https://developer.apple.com/documentation/AppKit/NSImageRep/cgImage(forProposedRect:context:hints:)
//
// [CGImage]: https://developer.apple.com/documentation/CoreGraphics/CGImage
//
// [CGImage]: https://developer.apple.com/documentation/CoreGraphics/CGImage
// [CGImage]: https://developer.apple.com/documentation/CoreGraphics/CGImage
func (i NSImageRep) CGImageForProposedRectContextHints(proposedDestRect *corefoundation.CGRect, context INSGraphicsContext, hints foundation.INSDictionary) coregraphics.CGImageRef {
	rv := objc.Send[coregraphics.CGImageRef](i.ID, objc.Sel("CGImageForProposedRect:context:hints:"), proposedDestRect, context, hints)
	return coregraphics.CGImageRef(rv)
}

// Implemented by subclasses to draw the image in the current coordinate
// system.
//
// # Return Value
//
// true if the image was successfully drawn; otherwise, false if there was a
// problem. The default version of this method simply returns true.
//
// # Discussion
//
// Subclass override this method to draw the image using the image data. By
// the time this method is called, the graphics state is already configured
// for you to draw the image at location (0.0, 0.0) in the current coordinate
// system.
//
// The standard Application Kit subclasses all draw the image using the
// [NSCompositeCopy] composite operation defined in the [Constants] section of
// [NSImage]. Using the copy operator, the image data overwrites the
// destination without any blending effects. Transparent (alpha) regions in
// the source image appear black. To use other composite operations, you must
// place the representation into an [NSImage] object and use its
// [DrawAtPointFromRectOperationFraction] or
// [DrawInRectFromRectOperationFraction] methods.
//
// See: https://developer.apple.com/documentation/AppKit/NSImageRep/draw()
func (i NSImageRep) Draw() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("draw"))
	return rv
}

// Draws the image representation’s image data at the specified point in the
// current coordinate system.
//
// point: The point in the current coordinate system at which to draw the image.
//
// # Return Value
//
// true if the image was successfully drawn; otherwise, false. If the size of
// the image has not yet been set, this method returns false immediately
//
// # Discussion
//
// This method sets the origin of the current coordinate system to the
// specified point and then invokes the receiver’s `draw` method to draw the
// image at that point. Upon completion, it restores the current coordinates
// to their original setting. If `aPoint` is (0.0, 0.0), this method simply
// invokes the [Draw] method.
//
// See: https://developer.apple.com/documentation/AppKit/NSImageRep/draw(at:)
func (i NSImageRep) DrawAtPoint(point corefoundation.CGPoint) bool {
	rv := objc.Send[bool](i.ID, objc.Sel("drawAtPoint:"), point)
	return rv
}

// Draws the image, scaling it (as needed) to fit the specified rectangle.
//
// rect: The rectangle in the current coordinate system in which to draw the image.
//
// # Return Value
//
// true if the image was successfully drawn; otherwise, false. If the size of
// the image has not yet been set, this method returns false immediately.
//
// # Discussion
//
// This method sets the origin of the current coordinate system to the origin
// of the specified rectangle before invoking the receiver’s [Draw] method.
// If the rectangle size is different from the image’s native size, this
// method adjusts the coordinate transform, causing the image to be scaled
// appropriately. After the `draw` method returns, the coordinate system
// changes are undone, restoring the original graphics state.
//
// See: https://developer.apple.com/documentation/AppKit/NSImageRep/draw(in:)
func (i NSImageRep) DrawInRect(rect corefoundation.CGRect) bool {
	rv := objc.Send[bool](i.ID, objc.Sel("drawInRect:"), rect)
	return rv
}

// Draws all or part of the image in the specified rectangle in the current
// coordinate system.
//
// dstSpacePortionRect: The rectangle in which to draw the image, specified in the current
// coordinate system.
//
// srcSpacePortionRect: The source rectangle specifying the portion of the image you want to draw.
// The coordinates of this rectangle must be specified using the image’s own
// coordinate system. If you pass in [NSZeroRect], the entire image is drawn.
//
// op: The compositing operation to use when drawing the image. See the
// [NSCompositingOperation] constants.
//
// requestedAlpha: The opacity of the image, specified as a value from 0.0 to 1.0. Specifying
// a value of 0.0 draws the image as fully transparent while a value of 1.0
// draws the image as fully opaque. Values greater than 1.0 are interpreted as
// 1.0.
//
// respectContextIsFlipped: true if the flipped context of the receiver should be respected, otherwise
// false.
//
// hints: An optional dictionary of hints that provide more context for selecting or
// generating the image. See `Image Hint Dictionary Keys` for possible values.
//
// # Return Value
//
// true if the image was successfully drawn; otherwise, false.
//
// # Discussion
//
// If the `srcSpacePortionRect` and `dstSpacePortionRect` rectangles have
// different sizes, the source portion of the image is scaled to fit the
// specified destination rectangle.
//
// See: https://developer.apple.com/documentation/AppKit/NSImageRep/draw(in:from:operation:fraction:respectFlipped:hints:)
//
// [NSCompositingOperation]: https://developer.apple.com/documentation/AppKit/NSCompositingOperation
func (i NSImageRep) DrawInRectFromRectOperationFractionRespectFlippedHints(dstSpacePortionRect corefoundation.CGRect, srcSpacePortionRect corefoundation.CGRect, op NSCompositingOperation, requestedAlpha float64, respectContextIsFlipped bool, hints foundation.INSDictionary) bool {
	rv := objc.Send[bool](i.ID, objc.Sel("drawInRect:fromRect:operation:fraction:respectFlipped:hints:"), dstSpacePortionRect, srcSpacePortionRect, op, requestedAlpha, respectContextIsFlipped, hints)
	return rv
}
func (i NSImageRep) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](i.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Creates and returns an array of image representation objects initialized
// using the contents of the specified file.
//
// filename: A full or relative pathname specifying the file to open. This string should
// include the filename extension.
//
// # Return Value
//
// An array of image representation objects. The array contains one object for
// each image in the specified file.
//
// # Discussion
//
// If sent to the [NSImageRep] class object, this method returns an array of
// objects (all newly allocated instances of a subclass of [NSImageRep],
// chosen through the use of [class(forFileType:)]) that have been initialized
// with the contents of the file. If sent to a subclass of [NSImageRep] that
// recognizes the file type, this method returns an array of objects (all
// instances of that subclass) that have been initialized with the contents of
// the file.
//
// This method returns `nil` in any of the following cases:
//
// - The message is sent to the [NSImageRep] class object and there are no
// subclasses in the [NSImageRep] class registry that handle the data in the
// file. - The message is sent to a subclass of [NSImageRep] and that subclass
// cannot handle the data in the file. - The [NSImageRep] subclass is unable
// to initialize itself with the contents of `filename`.
//
// The [NSImageRep] subclass is initialized by creating an [NSData] object
// based on the contents of the file and passing it to the “ method of the
// subclass. By default, the files handled include those with the extensions
// “`tiff`”, “`gif`”, “`jpg`”, “`pict`”, “`pdf`”, and
// “`eps`”.
//
// See: https://developer.apple.com/documentation/AppKit/NSImageRep/imageReps(withContentsOfFile:)
//
// [class(forFileType:)]: https://developer.apple.com/documentation/AppKit/NSImageRep/class(forFileType:)
func (_NSImageRepClass NSImageRepClass) ImageRepsWithContentsOfFile(filename string) []NSImageRep {
	rv := objc.Send[[]objc.ID](objc.ID(_NSImageRepClass.class), objc.Sel("imageRepsWithContentsOfFile:"), objc.String(filename))
	return objc.ConvertSlice(rv, func(id objc.ID) NSImageRep {
		return NSImageRepFromID(id)
	})
}

// Creates and returns an array of image representation objects initialized
// using the contents of the pasteboard.
//
// pasteboard: The pasteboard containing the image data.
//
// # Return Value
//
// An array of image representation objects. The array contains one object for
// each image in the specified pasteboard.
//
// # Discussion
//
// If sent to the [NSImageRep] class object, this method returns an array of
// objects (all newly-allocated instances of a subclass of [NSImageRep]) that
// have been initialized with the data in the specified pasteboard. If sent to
// a subclass of [NSImageRep] that recognizes the pasteboard data, it returns
// an array of objects (all instances of that subclass) initialized with the
// pasteboard data.
//
// This method returns `nil` in any of the following cases:
//
// - The message is sent to the [NSImageRep] class object and there are no
// subclasses in the [NSImageRep] class registry that handle the pasteboard
// data. - The message is sent to a subclass of [NSImageRep] and that subclass
// cannot handle the pasteboard data. - The [NSImageRep] subclass is unable to
// initialize itself with the contents the pasteboard.
//
// The [NSImageRep] subclass is initialized by creating an [NSData] object
// based on the data in `pasteboard` and passing it to the “ method.
//
// See: https://developer.apple.com/documentation/AppKit/NSImageRep/imageReps(with:)
func (_NSImageRepClass NSImageRepClass) ImageRepsWithPasteboard(pasteboard INSPasteboard) []NSImageRep {
	rv := objc.Send[[]objc.ID](objc.ID(_NSImageRepClass.class), objc.Sel("imageRepsWithPasteboard:"), pasteboard)
	return objc.ConvertSlice(rv, func(id objc.ID) NSImageRep {
		return NSImageRepFromID(id)
	})
}

// Creates and returns an array of image representation objects initialized
// using the contents of the specified URL.
//
// url: The URL pointing to the image data.
//
// # Return Value
//
// An array of image representation objects. The array contains one object for
// each image in the data at the specified URL.
//
// # Discussion
//
// If sent to the [NSImageRep] class object, this method returns an array of
// objects (all newly allocated instances of a subclass of [NSImageRep]) that
// have been initialized with the contents of the specified URL. If sent to a
// subclass of [NSImageRep] that recognizes the data at the specified URL, it
// returns an array of objects (all instances of that subclass) that have been
// initialized with the contents of that URL.
//
// This method returns `nil` in any of the following cases:
//
// - The message is sent to the [NSImageRep] class object and there are no
// subclasses in the [NSImageRep] class registry that handle data in the
// specified URL. - The message is sent to a subclass of [NSImageRep] and that
// subclass cannot handle data in the specified URL. - The [NSImageRep]
// subclass is unable to initialize itself with the contents of the specified
// URL.
//
// The [NSImageRep] subclass is initialized by creating an [NSData] object
// based on the contents of the specified URL and passing it to the “ method.
//
// See: https://developer.apple.com/documentation/AppKit/NSImageRep/imageReps(withContentsOf:)
func (_NSImageRepClass NSImageRepClass) ImageRepsWithContentsOfURL(url foundation.INSURL) []NSImageRep {
	rv := objc.Send[[]objc.ID](objc.ID(_NSImageRepClass.class), objc.Sel("imageRepsWithContentsOfURL:"), url)
	return objc.ConvertSlice(rv, func(id objc.ID) NSImageRep {
		return NSImageRepFromID(id)
	})
}

// Returns a Boolean value that indicates whether the image representation can
// initialize itself from the specified data.
//
// data: The image data.
//
// # Return Value
//
// true if the receiver understands the format of the specified data and can
// use it to initialize itself; otherwise, false.
//
// # Discussion
//
// This method should be overridden by subclasses. Note that this method does
// not need to do a comprehensive check of the image data; it should return
// false only if it knows it cannot initialize itself from the data.
//
// See: https://developer.apple.com/documentation/AppKit/NSImageRep/canInit(with:)-6zv56
func (_NSImageRepClass NSImageRepClass) CanInitWithData(data foundation.INSData) bool {
	rv := objc.Send[bool](objc.ID(_NSImageRepClass.class), objc.Sel("canInitWithData:"), data)
	return rv
}

// Returns a Boolean value that indicates whether the receiver can initialize
// itself from the data on the specified pasteboard.
//
// pasteboard: The pasteboard containing the image data.
//
// # Return Value
//
// true if the receiver understands the format of the specified data and can
// use it to initialize itself; otherwise, false.
//
// # Discussion
//
// This method invokes the [imageUnfilteredPasteboardTypes()] class method and
// checks the list of types returned by that method against the data types in
// `pasteboard`. If it finds a match, it returns true. When creating a
// subclass of [NSImageRep] that accepts image data from a non-default
// pasteboard type, override the [imageUnfilteredPasteboardTypes()] method to
// assure this method returns the correct response.
//
// See: https://developer.apple.com/documentation/AppKit/NSImageRep/canInit(with:)-56pum
//
// [imageUnfilteredPasteboardTypes()]: https://developer.apple.com/documentation/AppKit/NSImageRep/imageUnfilteredPasteboardTypes()
func (_NSImageRepClass NSImageRepClass) CanInitWithPasteboard(pasteboard INSPasteboard) bool {
	rv := objc.Send[bool](objc.ID(_NSImageRepClass.class), objc.Sel("canInitWithPasteboard:"), pasteboard)
	return rv
}

// Returns the image representation subclass that handles image data for the
// specified UTI.
//
// type: The UTI string identifying the desired image type. Some sample
// image-related UTI strings include “`public.Image()`”,
// “`public.Jpeg()`”, and “`public.Tiff()`”. For a list of supported
// types, see `UTCoreTypes.H()`.
//
// # Return Value
//
// A [Class] object for the image representation that can handle the UTI, or
// `nil` if no image representation could handle the data.
//
// See: https://developer.apple.com/documentation/AppKit/NSImageRep/class(forType:)
func (_NSImageRepClass NSImageRepClass) ImageRepClassForType(type_ string) objc.Class {
	rv := objc.Send[objc.Class](objc.ID(_NSImageRepClass.class), objc.Sel("imageRepClassForType:"), objc.String(type_))
	return rv
}

// Returns the image representation subclass that handles the specified type
// of data.
//
// data: The image data.
//
// # Return Value
//
// A [Class] object for the image representation that can handle the data, or
// `nil` if no image representation could handle the data.
//
// See: https://developer.apple.com/documentation/AppKit/NSImageRep/class(for:)
func (_NSImageRepClass NSImageRepClass) ImageRepClassForData(data foundation.INSData) objc.Class {
	rv := objc.Send[objc.Class](objc.ID(_NSImageRepClass.class), objc.Sel("imageRepClassForData:"), data)
	return rv
}

// Adds the specified class to the registry of available image representation
// subclasses.
//
// imageRepClass: The [Class] object for an [NSImageRep] subclass.
//
// # Discussion
//
// This method posts an [registryDidChangeNotification], along with the
// receiving object, to the default notification center.
//
// A good place to add image representation classes to the registry is in the
// `load` class method.
//
// See: https://developer.apple.com/documentation/AppKit/NSImageRep/registerClass(_:)
//
// [registryDidChangeNotification]: https://developer.apple.com/documentation/AppKit/NSImageRep/registryDidChangeNotification
func (_NSImageRepClass NSImageRepClass) RegisterImageRepClass(imageRepClass objc.Class) {
	objc.Send[objc.ID](objc.ID(_NSImageRepClass.class), objc.Sel("registerImageRepClass:"), imageRepClass)
}

// Removes the specified image representation subclass from the registry of
// available image representations.
//
// imageRepClass: The [Class] object for an [NSImageRep] subclass.
//
// # Discussion
//
// This method posts the [registryDidChangeNotification], along with the
// receiving object, to the default notification center.
//
// See: https://developer.apple.com/documentation/AppKit/NSImageRep/unregisterClass(_:)
//
// [registryDidChangeNotification]: https://developer.apple.com/documentation/AppKit/NSImageRep/registryDidChangeNotification
func (_NSImageRepClass NSImageRepClass) UnregisterImageRepClass(imageRepClass objc.Class) {
	objc.Send[objc.ID](objc.ID(_NSImageRepClass.class), objc.Sel("unregisterImageRepClass:"), imageRepClass)
}

// The size of the image representation, measured in points in the user
// coordinate space.
//
// # Discussion
//
// This size is the size of the image representation when it’s rendered. It
// is not necessarily the same as the width and height of the image in pixels
// as specified by the image data, nor must it be equal to the size set for
// the [NSImage] object that wraps this image representation.
//
// The size of an image representation combined with the physical dimensions
// of the image data determine the resolution of the image.
//
// See: https://developer.apple.com/documentation/AppKit/NSImageRep/size
func (i NSImageRep) Size() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](i.ID, objc.Sel("size"))
	return corefoundation.CGSize(rv)
}
func (i NSImageRep) SetSize(value corefoundation.CGSize) {
	objc.Send[struct{}](i.ID, objc.Sel("setSize:"), value)
}

// The number of bits per sample in the object (if the object is a planar
// image, this property contains the number of bits per sample per plane).
//
// # Discussion
//
// The number of bits used to specify each component of data in a single pixel
// (for example, a value of 8 for an RGBA image means that each pixel is
// comprised of four 8-bit values) or [NSImageRepMatchesDevice].
//
// A subclass can set this property when loading image data to notify the
// parent class of how many bits each sample uses. Specifying a value that
// differs from the actual image data does not change the bit depth of the
// image.
//
// See: https://developer.apple.com/documentation/AppKit/NSImageRep/bitsPerSample
func (i NSImageRep) BitsPerSample() int {
	rv := objc.Send[int](i.ID, objc.Sel("bitsPerSample"))
	return rv
}
func (i NSImageRep) SetBitsPerSample(value int) {
	objc.Send[struct{}](i.ID, objc.Sel("setBitsPerSample:"), value)
}

// The name of the color space used by the image data.
//
// # Discussion
//
// By default, an [NSImageRep] object’s color space name is
// [NSCalibratedRGBColorSpace]. Color space names are defined as part of the
// [NSColor] class, in `NSGraphics.H()`. The following are valid color space
// names:
//
// - [NSCalibratedWhiteColorSpace] - [NSCalibratedBlackColorSpace] -
// [NSCalibratedRGBColorSpace] - [NSDeviceWhiteColorSpace] -
// [NSDeviceBlackColorSpace] - [NSDeviceRGBColorSpace] -
// [NSDeviceCMYKColorSpace] - [NSNamedColorSpace] - [NSCustomColorSpace]
//
// See: https://developer.apple.com/documentation/AppKit/NSImageRep/colorSpaceName
func (i NSImageRep) ColorSpaceName() NSColorSpaceName {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("colorSpaceName"))
	return NSColorSpaceName(foundation.NSStringFromID(rv).String())
}
func (i NSImageRep) SetColorSpaceName(value NSColorSpaceName) {
	objc.Send[struct{}](i.ID, objc.Sel("setColorSpaceName:"), objc.String(string(value)))
}

// A Boolean value that indicates whether the image data has an alpha channel.
//
// # Discussion
//
// The value of this property is true if the receiver has a known alpha
// channel; otherwise, false.
//
// Subclasses can set this property when loading image data to notify the
// parent class whether that data contains an alpha component. Specifying a
// value of true does not add an alpha channel to the image data itself; it
// merely records the fact that the data has an alpha channel.
//
// See: https://developer.apple.com/documentation/AppKit/NSImageRep/hasAlpha
func (i NSImageRep) Alpha() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("hasAlpha"))
	return rv
}
func (i NSImageRep) SetAlpha(value bool) {
	objc.Send[struct{}](i.ID, objc.Sel("setAlpha:"), value)
}

// A Boolean value that indicates whether the image is opaque.
//
// # Discussion
//
// The value of this property is true if the image should be treated as fully
// opaque; otherwise, false to indicate the image may include some transparent
// regions.
//
// Use this property to test whether an image representation completely covers
// the area within the rectangle given by the [Size] property.
//
// The property value does not indicate whether the image has an alpha channel
// or if there is partial or complete transparency when drawing the image rep.
// Use the [Alpha] property to determine if the image has an alpha channel.
//
// See: https://developer.apple.com/documentation/AppKit/NSImageRep/isOpaque
func (i NSImageRep) Opaque() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("isOpaque"))
	return rv
}
func (i NSImageRep) SetOpaque(value bool) {
	objc.Send[struct{}](i.ID, objc.Sel("setOpaque:"), value)
}

// The height of the image, measured in pixels.
//
// # Discussion
//
// The value of this property is the height of the image, measured in the
// units of the device coordinate space. This value is usually derived from
// the image data itself.
//
// Subclasses can use this property when loading image data to notify the
// parent class of the image height. Setting this property does not change the
// actual number of pixels in the image.
//
// See: https://developer.apple.com/documentation/AppKit/NSImageRep/pixelsHigh
func (i NSImageRep) PixelsHigh() int {
	rv := objc.Send[int](i.ID, objc.Sel("pixelsHigh"))
	return rv
}
func (i NSImageRep) SetPixelsHigh(value int) {
	objc.Send[struct{}](i.ID, objc.Sel("setPixelsHigh:"), value)
}

// The width of the image, measured in pixels.
//
// # Discussion
//
// The value of this property is the width of the image, measured in the units
// of the device coordinate space. This value is usually derived from the
// image data itself.
//
// Subclasses can use this property when loading image data to notify the
// parent class of the image width. Setting this property does not change the
// actual number of pixels in the image.
//
// See: https://developer.apple.com/documentation/AppKit/NSImageRep/pixelsWide
func (i NSImageRep) PixelsWide() int {
	rv := objc.Send[int](i.ID, objc.Sel("pixelsWide"))
	return rv
}
func (i NSImageRep) SetPixelsWide(value int) {
	objc.Send[struct{}](i.ID, objc.Sel("setPixelsWide:"), value)
}

// The layout direction for the image.
//
// # Discussion
//
// For possible values, see [NSImage.LayoutDirection]. The default value for
// new image representation objects is [NSImageLayoutDirectionUnspecified].
//
// See: https://developer.apple.com/documentation/AppKit/NSImageRep/layoutDirection
//
// [NSImage.LayoutDirection]: https://developer.apple.com/documentation/AppKit/NSImage/LayoutDirection
func (i NSImageRep) LayoutDirection() NSImageLayoutDirection {
	rv := objc.Send[NSImageLayoutDirection](i.ID, objc.Sel("layoutDirection"))
	return NSImageLayoutDirection(rv)
}
func (i NSImageRep) SetLayoutDirection(value NSImageLayoutDirection) {
	objc.Send[struct{}](i.ID, objc.Sel("setLayoutDirection:"), value)
}

// Returns an array of UTI strings identifying the image types supported by
// the image representation, either directly or through a user-installed
// filter service.
//
// # Return Value
//
// An array of [NSString] objects, each of which contains a UTI identifying a
// supported image type. Some sample image-related UTI strings include
// “`public.Image()`”, “`public.Jpeg()`”, and “`public.Tiff()`”.
// For a list of supported types, see `UTCoreTypes.H()`.
//
// # Discussion
//
// The returned list includes UTIs all file types supported by this image
// representation object plus those that can be opened by this image
// representation after being converted by a user-installed filter service.
// You can use the returned UTI strings with any method that supports UTIs.
//
// See: https://developer.apple.com/documentation/AppKit/NSImageRep/imageTypes
func (_NSImageRepClass NSImageRepClass) ImageTypes() []string {
	rv := objc.Send[[]objc.ID](objc.ID(_NSImageRepClass.class), objc.Sel("imageTypes"))
	return objc.ConvertSliceToStrings(rv)
}

// Returns an array of UTI strings identifying the image types supported
// directly by the ime representation.
//
// # Return Value
//
// An array of [NSString] objects, each of which contains a UTI identifying a
// supported image type. Some sample image-related UTI strings include
// “`public.Image()`”, “`public.Jpeg()`”, and “`public.Tiff()`”.
// For a list of supported types, see `UTCoreTypes.H()`.
//
// # Discussion
//
// The returned list includes UTI strings only for those file types that are
// supported directly by the receiver. It does not include types that are
// supported through user-installed filter services. You can use the returned
// UTI strings with any method that supports UTIs.
//
// See: https://developer.apple.com/documentation/AppKit/NSImageRep/imageUnfilteredTypes
func (_NSImageRepClass NSImageRepClass) ImageUnfilteredTypes() []string {
	rv := objc.Send[[]objc.ID](objc.ID(_NSImageRepClass.class), objc.Sel("imageUnfilteredTypes"))
	return objc.ConvertSliceToStrings(rv)
}

// Returns an array containing the registered image representation classes.
//
// # Return Value
//
// An array of [Class] objects identifying the registered [NSImageRep]
// subclasses.
//
// See: https://developer.apple.com/documentation/AppKit/NSImageRep/registeredClasses
func (_NSImageRepClass NSImageRepClass) RegisteredImageRepClasses() []objc.Class {
	rv := objc.Send[[]objc.ID](objc.ID(_NSImageRepClass.class), objc.Sel("registeredImageRepClasses"))
	return objc.ConvertSlice(rv, func(id objc.ID) objc.Class {
		return objc.Class(id)
	})
}
