// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"context"
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSImage] class.
var (
	_NSImageClass     NSImageClass
	_NSImageClassOnce sync.Once
)

func getNSImageClass() NSImageClass {
	_NSImageClassOnce.Do(func() {
		_NSImageClass = NSImageClass{class: objc.GetClass("NSImage")}
	})
	return _NSImageClass
}

// GetNSImageClass returns the class object for NSImage.
func GetNSImageClass() NSImageClass {
	return getNSImageClass()
}

type NSImageClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSImageClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSImageClass) Alloc() NSImage {
	rv := objc.Send[NSImage](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A high-level interface for manipulating image data.
//
// # Overview
//
// You use instances of [NSImage] to load existing images, create new images,
// and draw the resulting image data into your views. Although you use this
// class predominantly for image-related operations, the class itself knows
// little about the underlying image data. Instead, it works in conjunction
// with one or more image representation objects (subclasses of [NSImageRep])
// to manage and render the image data. For the most part, these interactions
// are transparent.
//
// The class serves many purposes, providing support for the following tasks:
//
// - Loading images stored on disk or at a specified URL. - Drawing images
// into a view or graphics context. - Providing the contents of a [CALayer]
// object. - Creating new images based on a series of captured drawing
// commands. - Producing versions of the image in a different format.
//
// The [NSImage] class itself is capable of managing image data in a variety
// of formats. The specific list of formats is dependent on the version of the
// operating system but includes many standard formats such as TIFF, JPEG,
// GIF, PNG, and PDF among others. AppKit manages each format using a specific
// type of image representation object, whose job is to manage the actual
// image data. You can get a list of supported formats using the methods
// described in Determining Supported Types of Images.
//
// For more information about how to use image objects in your app, see [Cocoa
// Drawing Guide].
//
// # Using Images with Core Animation Layers
//
// Although you can assign an [NSImage] object directly to the [NSImage.Contents]
// property of a [CALayer] object, doing so may not always yield the best
// results. Instead of using your image object, you can use the
// [NSImage.LayerContentsForContentsScale] method to obtain an object that you can use
// for your layer’s contents. The image created by that method serves as the
// contents of a layer, which also supports all of the layer’s gravity
// modes. By contrast, the [NSImage] class supports only the [NSImage.Resize],
// [NSImage.ResizeAspect], and [NSImage.ResizeAspectFill] modes.
//
// Before calling the [NSImage.LayerContentsForContentsScale] method, use the
// [NSImage.RecommendedLayerContentsScale] method to get the recommended scale factor
// for the resulting image. The code listing below shows a typical example
// that uses the scale factor of a window’s backing store as the desired
// scale factor. From that scale factor, the code gets the scale factor for
// the specified image object and creates an object that you assign to the
// layer. You might use this code for images that fit the layer bounds
// precisely or for which you rely on the [NSImage.ContentsGravity] property of the
// layer to position or scale the image.
//
// Listing 1. Assigning an image to a layer
//
// # Creating Images by Name
//
//   - [NSImage.SetName]: Registers the image object under the specified name.
//   - [NSImage.Name]: Returns the name associated with the image, if any.
//
// # Creating Images from Resource Files
//
//   - [NSImage.InitByReferencingFile]: Initializes and returns an image object using the specified file.
//   - [NSImage.InitByReferencingURL]: Initializes and returns an image object using the specified URL.
//   - [NSImage.InitWithContentsOfFile]: Initializes and returns an image object with the contents of the specified file.
//   - [NSImage.InitWithContentsOfURL]: Initializes and returns an image object with the contents of the specified URL.
//
// # Creating Images from Existing Data
//
//   - [NSImage.InitWithData]: Initializes and returns an image object using the provided image data.
//   - [NSImage.InitWithDataIgnoringOrientation]: Initializes and returns an image object using the provided image data and ignoring the EXIF orientation tags.
//   - [NSImage.InitWithCGImageSize]: Creates a new image using the contents of the provided image.
//   - [NSImage.InitWithPasteboard]: Initializes and returns an image object with data from the specified pasteboard.
//   - [NSImage.InitWithCoder]: Initializes and returns an image object from data in an unarchiver.
//
// # Creating Empty Images
//
//   - [NSImage.InitWithSize]: Initializes and returns an image object with the specified dimensions.
//
// # Creating Symbol Images
//
//   - [NSImage.ImageWithSymbolConfiguration]: Creates a new symbol image with the specified configuration.
//
// # Getting the Symbol Image Configuration
//
//   - [NSImage.SymbolConfiguration]: The configuration details for a symbol image.
//
// # Managing Loading and Drawing of Images
//
//   - [NSImage.Delegate]: The image’s delegate object.
//   - [NSImage.SetDelegate]
//
// # Setting Attributes of Images
//
//   - [NSImage.Size]: The size of the image.
//   - [NSImage.SetSize]
//   - [NSImage.Template]: A Boolean value that determines whether the image represents a template image.
//   - [NSImage.SetTemplate]
//
// # Working with Representations of Images
//
//   - [NSImage.AddRepresentation]: Adds the specified image representation object to the image.
//   - [NSImage.AddRepresentations]: Adds an array of image representation objects to the image.
//   - [NSImage.Representations]: An array containing all of the image object’s image representations.
//   - [NSImage.RemoveRepresentation]: Removes and releases the specified image representation.
//   - [NSImage.BestRepresentationForRectContextHints]: Returns the best representation of the image for the specified rectangle using the provided hints.
//
// # Setting the Representation Selection Criteria for Images
//
//   - [NSImage.PrefersColorMatch]: A Boolean value that indicates whether the image prefers to choose image representations using color-matching or resolution-matching.
//   - [NSImage.SetPrefersColorMatch]
//   - [NSImage.UsesEPSOnResolutionMismatch]: A Boolean value that indicates whether EPS representations are preferred when no other representations match the resolution of the device.
//   - [NSImage.SetUsesEPSOnResolutionMismatch]
//   - [NSImage.MatchesOnMultipleResolution]: A Boolean value that indicates whether image representations whose resolution is an integral multiple of the device resolution are a match.
//   - [NSImage.SetMatchesOnMultipleResolution]
//
// # Drawing Images
//
//   - [NSImage.DrawInRect]: Draws the image in the specified rectangle.
//   - [NSImage.DrawAtPointFromRectOperationFraction]: Draws all or part of the image at the specified point in the current coordinate system.
//   - [NSImage.DrawInRectFromRectOperationFraction]: Draws all or part of the image in the specified rectangle in the current coordinate system.
//   - [NSImage.DrawInRectFromRectOperationFractionRespectFlippedHints]: Draws all or part of the image in the specified rectangle respecting the hints and the orientation of the current coordinate system.
//   - [NSImage.DrawRepresentationInRect]: Draws the image using the specified image representation object.
//
// # Managing Drawing Options
//
//   - [NSImage.Valid]: A Boolean value that indicates whether it is possible to draw an image representation.
//   - [NSImage.BackgroundColor]: The background color for the image.
//   - [NSImage.SetBackgroundColor]
//   - [NSImage.CapInsets]: The cap insets for the image.
//   - [NSImage.SetCapInsets]
//   - [NSImage.ResizingMode]: The resizing mode for the image.
//   - [NSImage.SetResizingMode]
//
// # Working with Alignment Metadata
//
//   - [NSImage.AlignmentRect]: A rectangle that you can use to position the image during layout.
//   - [NSImage.SetAlignmentRect]
//
// # Managing Caching Options
//
//   - [NSImage.CacheMode]: The image’s caching mode.
//   - [NSImage.SetCacheMode]
//   - [NSImage.Recache]: Invalidates and frees offscreen caches of all image representations.
//
// # Producing TIFF Data for Images
//
//   - [NSImage.TIFFRepresentation]: A data object containing TIFF data for all of the image representations in the image.
//   - [NSImage.TIFFRepresentationUsingCompressionFactor]: Returns a data object that contains TIFF data with the specified compression settings for all of the image representations in the image.
//
// # Producing Core Graphics Images
//
//   - [NSImage.CGImageForProposedRectContextHints]: Returns a Core Graphics image based on the contents of the current image object.
//
// # Hit-Testing Images
//
//   - [NSImage.HitTestRectWithImageDestinationRectContextHintsFlipped]: Returns whether the destination rectangle would intersect a non-transparent portion of the image.
//
// # Managing Image Accessibility
//
//   - [NSImage.AccessibilityDescription]: The image’s accessibility description.
//   - [NSImage.SetAccessibilityDescription]
//
// # Using Images with Core Animation
//
//   - [NSImage.LayerContentsForContentsScale]: Returns an object that may be used as the contents of a layer.
//   - [NSImage.RecommendedLayerContentsScale]: Returns the recommended layer contents scale for this image.
//
// # Managing Axis Matching
//
//   - [NSImage.MatchesOnlyOnBestFittingAxis]: A Boolean value that indicates whether the image matches only on the best fitting axis.
//   - [NSImage.SetMatchesOnlyOnBestFittingAxis]
//
// # Localizing Images
//
//   - [NSImage.ImageWithLocale]
//   - [NSImage.Locale]
//
// See: https://developer.apple.com/documentation/AppKit/NSImage
//
// [CALayer]: https://developer.apple.com/documentation/QuartzCore/CALayer
// [Cocoa Drawing Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/CocoaDrawingGuide/Introduction/Introduction.html#//apple_ref/doc/uid/TP40003290
type NSImage struct {
	objectivec.Object
}

// NSImageFromID constructs a [NSImage] from an objc.ID.
//
// A high-level interface for manipulating image data.
func NSImageFromID(id objc.ID) NSImage {
	return NSImage{objectivec.Object{ID: id}}
}

// NOTE: NSImage adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSImage] class.
//
// # Creating Images by Name
//
//   - [INSImage.SetName]: Registers the image object under the specified name.
//   - [INSImage.Name]: Returns the name associated with the image, if any.
//
// # Creating Images from Resource Files
//
//   - [INSImage.InitByReferencingFile]: Initializes and returns an image object using the specified file.
//   - [INSImage.InitByReferencingURL]: Initializes and returns an image object using the specified URL.
//   - [INSImage.InitWithContentsOfFile]: Initializes and returns an image object with the contents of the specified file.
//   - [INSImage.InitWithContentsOfURL]: Initializes and returns an image object with the contents of the specified URL.
//
// # Creating Images from Existing Data
//
//   - [INSImage.InitWithData]: Initializes and returns an image object using the provided image data.
//   - [INSImage.InitWithDataIgnoringOrientation]: Initializes and returns an image object using the provided image data and ignoring the EXIF orientation tags.
//   - [INSImage.InitWithCGImageSize]: Creates a new image using the contents of the provided image.
//   - [INSImage.InitWithPasteboard]: Initializes and returns an image object with data from the specified pasteboard.
//   - [INSImage.InitWithCoder]: Initializes and returns an image object from data in an unarchiver.
//
// # Creating Empty Images
//
//   - [INSImage.InitWithSize]: Initializes and returns an image object with the specified dimensions.
//
// # Creating Symbol Images
//
//   - [INSImage.ImageWithSymbolConfiguration]: Creates a new symbol image with the specified configuration.
//
// # Getting the Symbol Image Configuration
//
//   - [INSImage.SymbolConfiguration]: The configuration details for a symbol image.
//
// # Managing Loading and Drawing of Images
//
//   - [INSImage.Delegate]: The image’s delegate object.
//   - [INSImage.SetDelegate]
//
// # Setting Attributes of Images
//
//   - [INSImage.Size]: The size of the image.
//   - [INSImage.SetSize]
//   - [INSImage.Template]: A Boolean value that determines whether the image represents a template image.
//   - [INSImage.SetTemplate]
//
// # Working with Representations of Images
//
//   - [INSImage.AddRepresentation]: Adds the specified image representation object to the image.
//   - [INSImage.AddRepresentations]: Adds an array of image representation objects to the image.
//   - [INSImage.Representations]: An array containing all of the image object’s image representations.
//   - [INSImage.RemoveRepresentation]: Removes and releases the specified image representation.
//   - [INSImage.BestRepresentationForRectContextHints]: Returns the best representation of the image for the specified rectangle using the provided hints.
//
// # Setting the Representation Selection Criteria for Images
//
//   - [INSImage.PrefersColorMatch]: A Boolean value that indicates whether the image prefers to choose image representations using color-matching or resolution-matching.
//   - [INSImage.SetPrefersColorMatch]
//   - [INSImage.UsesEPSOnResolutionMismatch]: A Boolean value that indicates whether EPS representations are preferred when no other representations match the resolution of the device.
//   - [INSImage.SetUsesEPSOnResolutionMismatch]
//   - [INSImage.MatchesOnMultipleResolution]: A Boolean value that indicates whether image representations whose resolution is an integral multiple of the device resolution are a match.
//   - [INSImage.SetMatchesOnMultipleResolution]
//
// # Drawing Images
//
//   - [INSImage.DrawInRect]: Draws the image in the specified rectangle.
//   - [INSImage.DrawAtPointFromRectOperationFraction]: Draws all or part of the image at the specified point in the current coordinate system.
//   - [INSImage.DrawInRectFromRectOperationFraction]: Draws all or part of the image in the specified rectangle in the current coordinate system.
//   - [INSImage.DrawInRectFromRectOperationFractionRespectFlippedHints]: Draws all or part of the image in the specified rectangle respecting the hints and the orientation of the current coordinate system.
//   - [INSImage.DrawRepresentationInRect]: Draws the image using the specified image representation object.
//
// # Managing Drawing Options
//
//   - [INSImage.Valid]: A Boolean value that indicates whether it is possible to draw an image representation.
//   - [INSImage.BackgroundColor]: The background color for the image.
//   - [INSImage.SetBackgroundColor]
//   - [INSImage.CapInsets]: The cap insets for the image.
//   - [INSImage.SetCapInsets]
//   - [INSImage.ResizingMode]: The resizing mode for the image.
//   - [INSImage.SetResizingMode]
//
// # Working with Alignment Metadata
//
//   - [INSImage.AlignmentRect]: A rectangle that you can use to position the image during layout.
//   - [INSImage.SetAlignmentRect]
//
// # Managing Caching Options
//
//   - [INSImage.CacheMode]: The image’s caching mode.
//   - [INSImage.SetCacheMode]
//   - [INSImage.Recache]: Invalidates and frees offscreen caches of all image representations.
//
// # Producing TIFF Data for Images
//
//   - [INSImage.TIFFRepresentation]: A data object containing TIFF data for all of the image representations in the image.
//   - [INSImage.TIFFRepresentationUsingCompressionFactor]: Returns a data object that contains TIFF data with the specified compression settings for all of the image representations in the image.
//
// # Producing Core Graphics Images
//
//   - [INSImage.CGImageForProposedRectContextHints]: Returns a Core Graphics image based on the contents of the current image object.
//
// # Hit-Testing Images
//
//   - [INSImage.HitTestRectWithImageDestinationRectContextHintsFlipped]: Returns whether the destination rectangle would intersect a non-transparent portion of the image.
//
// # Managing Image Accessibility
//
//   - [INSImage.AccessibilityDescription]: The image’s accessibility description.
//   - [INSImage.SetAccessibilityDescription]
//
// # Using Images with Core Animation
//
//   - [INSImage.LayerContentsForContentsScale]: Returns an object that may be used as the contents of a layer.
//   - [INSImage.RecommendedLayerContentsScale]: Returns the recommended layer contents scale for this image.
//
// # Managing Axis Matching
//
//   - [INSImage.MatchesOnlyOnBestFittingAxis]: A Boolean value that indicates whether the image matches only on the best fitting axis.
//   - [INSImage.SetMatchesOnlyOnBestFittingAxis]
//
// # Localizing Images
//
//   - [INSImage.ImageWithLocale]
//   - [INSImage.Locale]
//
// See: https://developer.apple.com/documentation/AppKit/NSImage
type INSImage interface {
	objectivec.IObject
	NSPasteboardWriting

	// Topic: Creating Images by Name

	// Registers the image object under the specified name.
	SetName(string_ NSImageName) bool
	// Returns the name associated with the image, if any.
	Name() NSImageName

	// Topic: Creating Images from Resource Files

	// Initializes and returns an image object using the specified file.
	InitByReferencingFile(fileName string) NSImage
	// Initializes and returns an image object using the specified URL.
	InitByReferencingURL(url foundation.INSURL) NSImage
	// Initializes and returns an image object with the contents of the specified file.
	InitWithContentsOfFile(fileName string) NSImage
	// Initializes and returns an image object with the contents of the specified URL.
	InitWithContentsOfURL(url foundation.INSURL) NSImage

	// Topic: Creating Images from Existing Data

	// Initializes and returns an image object using the provided image data.
	InitWithData(data foundation.INSData) NSImage
	// Initializes and returns an image object using the provided image data and ignoring the EXIF orientation tags.
	InitWithDataIgnoringOrientation(data foundation.INSData) NSImage
	// Creates a new image using the contents of the provided image.
	InitWithCGImageSize(cgImage coregraphics.CGImageRef, size corefoundation.CGSize) NSImage
	// Initializes and returns an image object with data from the specified pasteboard.
	InitWithPasteboard(pasteboard INSPasteboard) NSImage
	// Initializes and returns an image object from data in an unarchiver.
	InitWithCoder(coder foundation.INSCoder) NSImage

	// Topic: Creating Empty Images

	// Initializes and returns an image object with the specified dimensions.
	InitWithSize(size corefoundation.CGSize) NSImage

	// Topic: Creating Symbol Images

	// Creates a new symbol image with the specified configuration.
	ImageWithSymbolConfiguration(configuration INSImageSymbolConfiguration) INSImage

	// Topic: Getting the Symbol Image Configuration

	// The configuration details for a symbol image.
	SymbolConfiguration() INSImageSymbolConfiguration

	// Topic: Managing Loading and Drawing of Images

	// The image’s delegate object.
	Delegate() NSImageDelegate
	SetDelegate(value NSImageDelegate)

	// Topic: Setting Attributes of Images

	// The size of the image.
	Size() corefoundation.CGSize
	SetSize(value corefoundation.CGSize)
	// A Boolean value that determines whether the image represents a template image.
	Template() bool
	SetTemplate(value bool)

	// Topic: Working with Representations of Images

	// Adds the specified image representation object to the image.
	AddRepresentation(imageRep INSImageRep)
	// Adds an array of image representation objects to the image.
	AddRepresentations(imageReps []NSImageRep)
	// An array containing all of the image object’s image representations.
	Representations() []NSImageRep
	// Removes and releases the specified image representation.
	RemoveRepresentation(imageRep INSImageRep)
	// Returns the best representation of the image for the specified rectangle using the provided hints.
	BestRepresentationForRectContextHints(rect corefoundation.CGRect, referenceContext INSGraphicsContext, hints foundation.INSDictionary) INSImageRep

	// Topic: Setting the Representation Selection Criteria for Images

	// A Boolean value that indicates whether the image prefers to choose image representations using color-matching or resolution-matching.
	PrefersColorMatch() bool
	SetPrefersColorMatch(value bool)
	// A Boolean value that indicates whether EPS representations are preferred when no other representations match the resolution of the device.
	UsesEPSOnResolutionMismatch() bool
	SetUsesEPSOnResolutionMismatch(value bool)
	// A Boolean value that indicates whether image representations whose resolution is an integral multiple of the device resolution are a match.
	MatchesOnMultipleResolution() bool
	SetMatchesOnMultipleResolution(value bool)

	// Topic: Drawing Images

	// Draws the image in the specified rectangle.
	DrawInRect(rect corefoundation.CGRect)
	// Draws all or part of the image at the specified point in the current coordinate system.
	DrawAtPointFromRectOperationFraction(point corefoundation.CGPoint, fromRect corefoundation.CGRect, op NSCompositingOperation, delta float64)
	// Draws all or part of the image in the specified rectangle in the current coordinate system.
	DrawInRectFromRectOperationFraction(rect corefoundation.CGRect, fromRect corefoundation.CGRect, op NSCompositingOperation, delta float64)
	// Draws all or part of the image in the specified rectangle respecting the hints and the orientation of the current coordinate system.
	DrawInRectFromRectOperationFractionRespectFlippedHints(dstSpacePortionRect corefoundation.CGRect, srcSpacePortionRect corefoundation.CGRect, op NSCompositingOperation, requestedAlpha float64, respectContextIsFlipped bool, hints foundation.INSDictionary)
	// Draws the image using the specified image representation object.
	DrawRepresentationInRect(imageRep INSImageRep, rect corefoundation.CGRect) bool

	// Topic: Managing Drawing Options

	// A Boolean value that indicates whether it is possible to draw an image representation.
	Valid() bool
	// The background color for the image.
	BackgroundColor() INSColor
	SetBackgroundColor(value INSColor)
	// The cap insets for the image.
	CapInsets() foundation.NSEdgeInsets
	SetCapInsets(value foundation.NSEdgeInsets)
	// The resizing mode for the image.
	ResizingMode() NSImageResizingMode
	SetResizingMode(value NSImageResizingMode)

	// Topic: Working with Alignment Metadata

	// A rectangle that you can use to position the image during layout.
	AlignmentRect() corefoundation.CGRect
	SetAlignmentRect(value corefoundation.CGRect)

	// Topic: Managing Caching Options

	// The image’s caching mode.
	CacheMode() NSImageCacheMode
	SetCacheMode(value NSImageCacheMode)
	// Invalidates and frees offscreen caches of all image representations.
	Recache()

	// Topic: Producing TIFF Data for Images

	// A data object containing TIFF data for all of the image representations in the image.
	TIFFRepresentation() foundation.INSData
	// Returns a data object that contains TIFF data with the specified compression settings for all of the image representations in the image.
	TIFFRepresentationUsingCompressionFactor(comp NSTIFFCompression, factor float32) foundation.INSData

	// Topic: Producing Core Graphics Images

	// Returns a Core Graphics image based on the contents of the current image object.
	CGImageForProposedRectContextHints(proposedDestRect *corefoundation.CGRect, referenceContext INSGraphicsContext, hints foundation.INSDictionary) coregraphics.CGImageRef

	// Topic: Hit-Testing Images

	// Returns whether the destination rectangle would intersect a non-transparent portion of the image.
	HitTestRectWithImageDestinationRectContextHintsFlipped(testRectDestSpace corefoundation.CGRect, imageRectDestSpace corefoundation.CGRect, context INSGraphicsContext, hints foundation.INSDictionary, flipped bool) bool

	// Topic: Managing Image Accessibility

	// The image’s accessibility description.
	AccessibilityDescription() string
	SetAccessibilityDescription(value string)

	// Topic: Using Images with Core Animation

	// Returns an object that may be used as the contents of a layer.
	LayerContentsForContentsScale(layerContentsScale float64) objectivec.IObject
	// Returns the recommended layer contents scale for this image.
	RecommendedLayerContentsScale(preferredContentsScale float64) float64

	// Topic: Managing Axis Matching

	// A Boolean value that indicates whether the image matches only on the best fitting axis.
	MatchesOnlyOnBestFittingAxis() bool
	SetMatchesOnlyOnBestFittingAxis(value bool)

	// Topic: Localizing Images

	ImageWithLocale(locale foundation.NSLocale) INSImage
	Locale() foundation.NSLocale

	// An object that provides the contents of the layer. Animatable.
	Contents() objectivec.IObject
	SetContents(value objectivec.IObject)
	// A constant that specifies how the layer’s contents are positioned or scaled within its bounds.
	ContentsGravity() foundation.NSString
	SetContentsGravity(value foundation.NSString)
	// The content is resized to fit the entire bounds rectangle.
	Resize() foundation.NSString
	// The content is resized to fit the bounds rectangle, preserving the aspect of the content. If the content does not completely fill the bounds rectangle, the content is centered in the partial axis.
	ResizeAspect() foundation.NSString
	// The content is resized to completely fill the bounds rectangle, while still preserving the aspect of the content. The content is centered in the axis it exceeds.
	ResizeAspectFill() foundation.NSString
	// Initializes an instance with a property list object and a type string.
	InitWithPasteboardPropertyListOfType(propertyList objectivec.IObject, type_ NSPasteboardType) NSImage
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (i NSImage) Init() NSImage {
	rv := objc.Send[NSImage](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i NSImage) Autorelease() NSImage {
	rv := objc.Send[NSImage](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSImage creates a new NSImage instance.
func NewNSImage() NSImage {
	class := getNSImageClass()
	rv := objc.Send[NSImage](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes and returns an image object using the specified file.
//
// fileName: A full or relative path name specifying the file with the desired image
// data. Relative paths must be relative to the current working directory.
//
// # Return Value
//
// An initialized [NSImage] object or `nil` if the new object cannot be
// initialized.
//
// # Discussion
//
// This method initializes the image object lazily. It does not actually open
// the specified file or create any image representations from its data until
// an app attempts to draw the image or request information about it.
//
// The `filename` parameter should include the file extension that identifies
// the type of the image data. The mechanism that actually creates the image
// representation for `filename` looks for an [NSImageRep] subclass that
// handles that data type from among those registered with [NSImage].
//
// Because this method doesn’t actually create image representations for the
// image data, your app should do error checking before attempting to use the
// image; one way to do so is by accessing the [Valid] property to check
// whether the image can be drawn.
//
// This method invokes [setDataRetained:] with an argument of true, thus
// enabling it to hold onto its filename. When archiving an image created with
// this method, only the image’s filename is written to the archive.
//
// If the cached version of the image uses less memory than the original image
// data, AppKit deletes the original data and uses the cached image. (This can
// occur for images whose resolution is greater than 72 dpi.) If you resize
// the image by less than 50%, AppKit loads the data in again from the file.
// If you expect to delete the file or change its contents, use
// [InitWithContentsOfFile] instead.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/init(byReferencingFile:)
//
// [setDataRetained:]: https://developer.apple.com/documentation/AppKit/NSImage/setDataRetained:
func NewImageByReferencingFile(fileName string) NSImage {
	instance := getNSImageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initByReferencingFile:"), objc.String(fileName))
	return NSImageFromID(rv)
}

// Initializes and returns an image object using the specified URL.
//
// url: The URL identifying the image.
//
// # Return Value
//
// An initialized [NSImage] object.
//
// # Discussion
//
// This method initializes the image object lazily. It does not attempt to
// retrieve the data from the specified URL or create any image
// representations from that data until an app attempts to draw the image or
// request information about it.
//
// The `url` parameter should include a file extension that identifies the
// type of the image data. The mechanism that actually creates the image
// representation looks for an [NSImageRep] subclass that handles that data
// type from among those registered with [NSImage].
//
// Because this method doesn’t actually create image representations for the
// image data, your app should do error checking before attempting to use the
// image; one way to do so is by accessing the [Valid] property to check
// whether the image can be drawn.
//
// This method invokes [setDataRetained:] with an argument of true, thus
// enabling it to hold onto its URL. When archiving an image created with this
// method, only the image’s URL is written to the archive.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/init(byReferencing:)
//
// [setDataRetained:]: https://developer.apple.com/documentation/AppKit/NSImage/setDataRetained:
func NewImageByReferencingURL(url foundation.INSURL) NSImage {
	instance := getNSImageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initByReferencingURL:"), url)
	return NSImageFromID(rv)
}

// Returns the image object associated with the specified name.
//
// name: The name associated with the desired image. This can be a name you assigned
// to the image or the name of an image file in your app bundle.
//
// # Return Value
//
// The [NSImage] object associated with the specified name or `nil` if no such
// image was found.
//
// # Discussion
//
// This method searches for named images in several places, returning the
// first image it finds matching the given name. The order of the search is as
// follows:
//
// - Search for an object whose name was set explicitly using the [SetName]
// method and currently resides in the image cache. - Search the app’s main
// bundle for a file whose name matches the specified string. (For information
// on how the bundle is searched, see “[Accessing a Bundle’s Contents]”
// in [Bundle Programming Guide].) - Search the Application Kit framework for
// a shared image with the specified name.
//
// When looking for files in the app bundle, it is better (but not required)
// to include the filename extension in the `name` parameter. When naming an
// image with the [SetName] method, it is convention not to include filename
// extensions in the names you specify. That way, you can easily distinguish
// between images you have named explicitly and those you want to load from
// the app’s bundle.
//
// One particularly useful image you can retrieve is your app’s icon. This
// image is set by Cocoa automatically and accessible using the
// [applicationIconName] constant. Icons for other apps can be obtained
// through the use of methods declared in the [NSWorkspace] class. You can
// also retrieve many of the standard system images using Cocoa defined
// constants; for more information, see the `Image Template Constants`,
// `Sharing Permissions Named Images`, `System Entity Images`, `Toolbar Named
// Images`, and `View Type Template Images` sections for applicable constants.
//
// If an app is linked on macOS 10.5 or later, images requested using this
// method and whose name ends in the word “Template” are automatically
// marked as template images.
//
// The [NSImage] class may cache a reference to the returned image object for
// performance in some cases. However, the class holds onto cached objects
// only while the object exists. If all strong references to the image are
// subsequently removed, the object may be quietly removed from the cache.
// Thus, if you plan to hold onto a returned image object, you must maintain a
// strong reference to it like you would any Cocoa object. You can clear an
// image object from the cache explicitly by calling the object’s [SetName]
// method and specifying `nil` for the image name.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/init(named:)
//
// [Accessing a Bundle’s Contents]: https://developer.apple.com/library/archive/documentation/CoreFoundation/Conceptual/CFBundles/AccessingaBundlesContents/AccessingaBundlesContents.html#//apple_ref/doc/uid/10000123i-CH104
// [Bundle Programming Guide]: https://developer.apple.com/library/archive/documentation/CoreFoundation/Conceptual/CFBundles/Introduction/Introduction.html#//apple_ref/doc/uid/10000123i
// [applicationIconName]: https://developer.apple.com/documentation/AppKit/NSImage/applicationIconName
func NewImageNamed(name NSImageName) NSImage {
	rv := objc.Send[objc.ID](objc.ID(getNSImageClass().class), objc.Sel("imageNamed:"), objc.String(string(name)))
	return NSImageFromID(rv)
}

// Creates a new image using the contents of the provided image.
//
// cgImage: The source image.
//
// size: The size of the new image. Use [zero], or [NSZeroSize] in Objective-C, to
// have the new image adopt the pixel dimensions of the source image.
//
// # Discussion
//
// Don’t assume anything about the image, other than drawing it is
// equivalent to drawing the source image.
//
// This is not a designated initializer.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/init(cgImage:size:)
//
// [zero]: https://developer.apple.com/documentation/CoreFoundation/CGSize/zero
func NewImageWithCGImageSize(cgImage coregraphics.CGImageRef, size corefoundation.CGSize) NSImage {
	instance := getNSImageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCGImage:size:"), cgImage, size)
	return NSImageFromID(rv)
}

// Initializes and returns an image object from data in an unarchiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/init(coder:)
func NewImageWithCoder(coder foundation.INSCoder) NSImage {
	instance := getNSImageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSImageFromID(rv)
}

// Initializes and returns an image object with the contents of the specified
// file.
//
// fileName: A full or relative path name specifying the file with the desired image
// data. Relative paths must be relative to the current working directory.
//
// # Return Value
//
// An initialized [NSImage] object or `nil` if the method cannot create an
// image representation from the contents of the specified file.
//
// # Discussion
//
// Unlike [InitByReferencingFile], which initializes an [NSImage] object
// lazily, this method immediately opens the specified file and creates one or
// more image representations from its data.
//
// The `filename` parameter should include the file extension that identifies
// the type of the image data. This method looks for an [NSImageRep] subclass
// that handles that data type from among those registered with [NSImage].
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/init(contentsOfFile:)
func NewImageWithContentsOfFile(fileName string) NSImage {
	instance := getNSImageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfFile:"), objc.String(fileName))
	return NSImageFromID(rv)
}

// Initializes and returns an image object with the contents of the specified
// URL.
//
// url: The URL identifying the image.
//
// # Return Value
//
// An initialized [NSImage] object or `nil` if the method cannot create an
// image representation from the contents of the specified URL.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/init(contentsOf:)
func NewImageWithContentsOfURL(url foundation.INSURL) NSImage {
	instance := getNSImageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfURL:"), url)
	return NSImageFromID(rv)
}

// Initializes and returns an image object using the provided image data.
//
// data: The data object containing the image data. The data can be in any format
// that macOS supports, including PDF, PICT, EPS, or any number of bitmap data
// formats.
//
// # Return Value
//
// An initialized [NSImage] object or `nil` if the method cannot create an
// image representation from the contents of the specified data object.
//
// # Discussion
//
// Use this method in cases where you already have image data in a supported
// format and want to obtain an [NSImage] object that represents that data.
// This method initializes the object with an image representation that is
// most appropriate for the type of data you provided.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/init(data:)
func NewImageWithData(data foundation.INSData) NSImage {
	instance := getNSImageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:"), data)
	return NSImageFromID(rv)
}

// Initializes and returns an image object using the provided image data and
// ignoring the EXIF orientation tags.
//
// data: The data object containing the image data. The data can be in any format
// that macOS supports, including PDF, PICT, EPS, or any number of bitmap data
// formats.
//
// # Return Value
//
// An initialized [NSImage] object or `nil` if the method cannot create an
// image representation from the contents of the specified data object.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/init(dataIgnoringOrientation:)
func NewImageWithDataIgnoringOrientation(data foundation.INSData) NSImage {
	instance := getNSImageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDataIgnoringOrientation:"), data)
	return NSImageFromID(rv)
}

// Initializes and returns an image object with data from the specified
// pasteboard.
//
// pasteboard: The pasteboard containing the image data. The data on the pasteboard can be
// in any format that macOS supports, including PDF, PICT, EPS, or any number
// of bitmap data formats.
//
// # Return Value
//
// An initialized [NSImage] object or `nil` if the method cannot create an
// image from the contents of the pasteboard.
//
// # Discussion
//
// The specified pasteboard should contain a type supported by one of the
// registered [NSImageRep] subclasses. The table below lists the default
// pasteboard types and file extensions for several [NSImageRep] subclasses.
//
// [Table data omitted]
//
// If the specified pasteboard contains the value [NSFilenamesPboardType],
// each filename on the pasteboard should have an extension supported by one
// of the registered [NSImageRep] subclasses. You can use the
// [imageUnfilteredFileTypes()] method of a given subclass to obtain the list
// of supported types for that class.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/init(pasteboard:)
//
// [imageUnfilteredFileTypes()]: https://developer.apple.com/documentation/AppKit/NSImageRep/imageUnfilteredFileTypes()
func NewImageWithPasteboard(pasteboard INSPasteboard) NSImage {
	instance := getNSImageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPasteboard:"), pasteboard)
	return NSImageFromID(rv)
}

// Initializes an instance with a property list object and a type string.
//
// propertyList: A property list containing data to initialize the receiver.
//
// By default, the property list object is an instance of [NSData]. If you
// implement [ReadingOptionsForTypePasteboard] and specify an option other
// than [NSPasteboardReadingAsData], the `propertyList` may be any other
// property list object.
//
// type: A UTI supported by the receiver for reading (one of the types returned by
// [ReadableTypesForPasteboard]).
//
// # Return Value
//
// An object initialized using the data in `propertyList`.
//
// # Discussion
//
// This method is considered optional because, if [ReadableTypesForPasteboard]
// returns just a single type, and that type uses the
// [NSPasteboardReadingAsKeyedArchive] reading option, then instances are
// initialized using [init(coder:)] instead of this method.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardReading/init(pasteboardPropertyList:ofType:)
//
// [init(coder:)]: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewImageWithPasteboardPropertyListOfType(propertyList objectivec.IObject, type_ NSPasteboardType) NSImage {
	instance := getNSImageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPasteboardPropertyList:ofType:"), propertyList, objc.String(string(type_)))
	return NSImageFromID(rv)
}

// Initializes and returns an image object with the specified dimensions.
//
// size: The size of the image, measured in points.
//
// # Return Value
//
// An initialized [NSImage] object with no rendered content.
//
// # Discussion
//
// This method does not add any image representations to the image object. It
// is permissible to initialize the image object by passing a size of `(0.0,
// 0.0)`; however, you must set the size to a non-zero value before using it
// or an exception will be raised.
//
// After using this method to initialize an image object, you are expected to
// provide the image contents before trying to draw the image. You might lock
// focus on the image and draw to the image or you might explicitly add an
// image representation that you created.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/init(size:)
func NewImageWithSize(size corefoundation.CGSize) NSImage {
	instance := getNSImageClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSize:"), size)
	return NSImageFromID(rv)
}

// See: https://developer.apple.com/documentation/AppKit/NSImage/init(symbolName:bundle:variableValue:)
func NewImageWithSymbolNameBundleVariableValue(name string, bundle foundation.NSBundle, value float64) NSImage {
	rv := objc.Send[objc.ID](objc.ID(getNSImageClass().class), objc.Sel("imageWithSymbolName:bundle:variableValue:"), objc.String(name), bundle, value)
	return NSImageFromID(rv)
}

// Creates a symbol image with the symbol name and variable value you specify.
//
// name: The name of the symbol image.
//
// value: The value the system uses to customize the symbol’s content, between `0`
// and `1`.
//
// # Discussion
//
// The `value` parameter is valid for symbols that support variable rendering.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/init(symbolName:variableValue:)
func NewImageWithSymbolNameVariableValue(name string, value float64) NSImage {
	rv := objc.Send[objc.ID](objc.ID(getNSImageClass().class), objc.Sel("imageWithSymbolName:variableValue:"), objc.String(name), value)
	return NSImageFromID(rv)
}

// Creates a symbol image with the system symbol name and accessibility
// description you specify.
//
// name: The name of the system symbol image.
//
// description: The accessibility description for the symbol image, if any.
//
// # Return Value
//
// A symbol image based on the name you specify; otherwise `nil` if the method
// couldn’t find a suitable image.
//
// # Discussion
//
// To look up the names of system symbol images, download the SF Symbols app
// from [Apple Design Resources].
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/init(systemSymbolName:accessibilityDescription:)
//
// [Apple Design Resources]: https://developer.apple.com/design/resources/
func NewImageWithSystemSymbolNameAccessibilityDescription(name string, description string) NSImage {
	rv := objc.Send[objc.ID](objc.ID(getNSImageClass().class), objc.Sel("imageWithSystemSymbolName:accessibilityDescription:"), objc.String(name), objc.String(description))
	return NSImageFromID(rv)
}

// Creates a symbol image with the system symbol name and variable value you
// specify.
//
// name: The name of the system symbol image.
//
// value: The value the system uses to customize the symbol’s content, between `0`
// and `1`.
//
// description: The accessibility description for the symbol image, if any.
//
// # Discussion
//
// The `value` parameter is valid for symbols that support variable rendering.
//
// To look up the names of system symbol images, download the SF Symbols app
// from [Apple Design Resources].
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/init(systemSymbolName:variableValue:accessibilityDescription:)
//
// [Apple Design Resources]: https://developer.apple.com/design/resources/
func NewImageWithSystemSymbolNameVariableValueAccessibilityDescription(name string, value float64, description string) NSImage {
	rv := objc.Send[objc.ID](objc.ID(getNSImageClass().class), objc.Sel("imageWithSystemSymbolName:variableValue:accessibilityDescription:"), objc.String(name), value, objc.String(description))
	return NSImageFromID(rv)
}

// Registers the image object under the specified name.
//
// string: The name to associate with the receiver. Specify `nil` if you want to
// remove the image from the image cache.
//
// # Return Value
//
// true if the receiver was successfully registered with the given name;
// otherwise, false.
//
// # Discussion
//
// If the receiver is already registered under a different name, this method
// unregisters the other name. If a different image is already registered
// under the name specified in `aString`, this method does nothing and returns
// false.
//
// When naming an image using this method, it is convention not to include
// filename extensions in the names you specify. That way, you can easily
// distinguish between images you have named explicitly and those you want to
// load from the app’s bundle. For information about the rules used to
// search for images, and for information about the ownership policy of named
// images, see the [ImageNamed] method.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/setName(_:)
func (i NSImage) SetName(string_ NSImageName) bool {
	rv := objc.Send[bool](i.ID, objc.Sel("setName:"), objc.String(string(string_)))
	return rv
}

// Returns the name associated with the image, if any.
//
// # Return Value
//
// The name associated with the image, or `nil` if it has no associated name.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/name()
func (i NSImage) Name() NSImageName {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("name"))
	return NSImageName(foundation.NSStringFromID(rv).String())
}

// Initializes and returns an image object using the specified file.
//
// fileName: A full or relative path name specifying the file with the desired image
// data. Relative paths must be relative to the current working directory.
//
// # Return Value
//
// An initialized [NSImage] object or `nil` if the new object cannot be
// initialized.
//
// # Discussion
//
// This method initializes the image object lazily. It does not actually open
// the specified file or create any image representations from its data until
// an app attempts to draw the image or request information about it.
//
// The `filename` parameter should include the file extension that identifies
// the type of the image data. The mechanism that actually creates the image
// representation for `filename` looks for an [NSImageRep] subclass that
// handles that data type from among those registered with [NSImage].
//
// Because this method doesn’t actually create image representations for the
// image data, your app should do error checking before attempting to use the
// image; one way to do so is by accessing the [Valid] property to check
// whether the image can be drawn.
//
// This method invokes [setDataRetained:] with an argument of true, thus
// enabling it to hold onto its filename. When archiving an image created with
// this method, only the image’s filename is written to the archive.
//
// If the cached version of the image uses less memory than the original image
// data, AppKit deletes the original data and uses the cached image. (This can
// occur for images whose resolution is greater than 72 dpi.) If you resize
// the image by less than 50%, AppKit loads the data in again from the file.
// If you expect to delete the file or change its contents, use
// [InitWithContentsOfFile] instead.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/init(byReferencingFile:)
//
// [setDataRetained:]: https://developer.apple.com/documentation/AppKit/NSImage/setDataRetained:
func (i NSImage) InitByReferencingFile(fileName string) NSImage {
	rv := objc.Send[NSImage](i.ID, objc.Sel("initByReferencingFile:"), objc.String(fileName))
	return rv
}

// Initializes and returns an image object using the specified URL.
//
// url: The URL identifying the image.
//
// # Return Value
//
// An initialized [NSImage] object.
//
// # Discussion
//
// This method initializes the image object lazily. It does not attempt to
// retrieve the data from the specified URL or create any image
// representations from that data until an app attempts to draw the image or
// request information about it.
//
// The `url` parameter should include a file extension that identifies the
// type of the image data. The mechanism that actually creates the image
// representation looks for an [NSImageRep] subclass that handles that data
// type from among those registered with [NSImage].
//
// Because this method doesn’t actually create image representations for the
// image data, your app should do error checking before attempting to use the
// image; one way to do so is by accessing the [Valid] property to check
// whether the image can be drawn.
//
// This method invokes [setDataRetained:] with an argument of true, thus
// enabling it to hold onto its URL. When archiving an image created with this
// method, only the image’s URL is written to the archive.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/init(byReferencing:)
//
// [setDataRetained:]: https://developer.apple.com/documentation/AppKit/NSImage/setDataRetained:
func (i NSImage) InitByReferencingURL(url foundation.INSURL) NSImage {
	rv := objc.Send[NSImage](i.ID, objc.Sel("initByReferencingURL:"), url)
	return rv
}

// Initializes and returns an image object with the contents of the specified
// file.
//
// fileName: A full or relative path name specifying the file with the desired image
// data. Relative paths must be relative to the current working directory.
//
// # Return Value
//
// An initialized [NSImage] object or `nil` if the method cannot create an
// image representation from the contents of the specified file.
//
// # Discussion
//
// Unlike [InitByReferencingFile], which initializes an [NSImage] object
// lazily, this method immediately opens the specified file and creates one or
// more image representations from its data.
//
// The `filename` parameter should include the file extension that identifies
// the type of the image data. This method looks for an [NSImageRep] subclass
// that handles that data type from among those registered with [NSImage].
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/init(contentsOfFile:)
func (i NSImage) InitWithContentsOfFile(fileName string) NSImage {
	rv := objc.Send[NSImage](i.ID, objc.Sel("initWithContentsOfFile:"), objc.String(fileName))
	return rv
}

// Initializes and returns an image object with the contents of the specified
// URL.
//
// url: The URL identifying the image.
//
// # Return Value
//
// An initialized [NSImage] object or `nil` if the method cannot create an
// image representation from the contents of the specified URL.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/init(contentsOf:)
func (i NSImage) InitWithContentsOfURL(url foundation.INSURL) NSImage {
	rv := objc.Send[NSImage](i.ID, objc.Sel("initWithContentsOfURL:"), url)
	return rv
}

// Initializes and returns an image object using the provided image data.
//
// data: The data object containing the image data. The data can be in any format
// that macOS supports, including PDF, PICT, EPS, or any number of bitmap data
// formats.
//
// # Return Value
//
// An initialized [NSImage] object or `nil` if the method cannot create an
// image representation from the contents of the specified data object.
//
// # Discussion
//
// Use this method in cases where you already have image data in a supported
// format and want to obtain an [NSImage] object that represents that data.
// This method initializes the object with an image representation that is
// most appropriate for the type of data you provided.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/init(data:)
func (i NSImage) InitWithData(data foundation.INSData) NSImage {
	rv := objc.Send[NSImage](i.ID, objc.Sel("initWithData:"), data)
	return rv
}

// Initializes and returns an image object using the provided image data and
// ignoring the EXIF orientation tags.
//
// data: The data object containing the image data. The data can be in any format
// that macOS supports, including PDF, PICT, EPS, or any number of bitmap data
// formats.
//
// # Return Value
//
// An initialized [NSImage] object or `nil` if the method cannot create an
// image representation from the contents of the specified data object.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/init(dataIgnoringOrientation:)
func (i NSImage) InitWithDataIgnoringOrientation(data foundation.INSData) NSImage {
	rv := objc.Send[NSImage](i.ID, objc.Sel("initWithDataIgnoringOrientation:"), data)
	return rv
}

// Creates a new image using the contents of the provided image.
//
// cgImage: The source image.
//
// size: The size of the new image. Use [zero], or [NSZeroSize] in Objective-C, to
// have the new image adopt the pixel dimensions of the source image.
//
// # Discussion
//
// Don’t assume anything about the image, other than drawing it is
// equivalent to drawing the source image.
//
// This is not a designated initializer.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/init(cgImage:size:)
//
// [zero]: https://developer.apple.com/documentation/CoreFoundation/CGSize/zero
func (i NSImage) InitWithCGImageSize(cgImage coregraphics.CGImageRef, size corefoundation.CGSize) NSImage {
	rv := objc.Send[NSImage](i.ID, objc.Sel("initWithCGImage:size:"), cgImage, size)
	return rv
}

// Initializes and returns an image object with data from the specified
// pasteboard.
//
// pasteboard: The pasteboard containing the image data. The data on the pasteboard can be
// in any format that macOS supports, including PDF, PICT, EPS, or any number
// of bitmap data formats.
//
// # Return Value
//
// An initialized [NSImage] object or `nil` if the method cannot create an
// image from the contents of the pasteboard.
//
// # Discussion
//
// The specified pasteboard should contain a type supported by one of the
// registered [NSImageRep] subclasses. The table below lists the default
// pasteboard types and file extensions for several [NSImageRep] subclasses.
//
// [Table data omitted]
//
// If the specified pasteboard contains the value [NSFilenamesPboardType],
// each filename on the pasteboard should have an extension supported by one
// of the registered [NSImageRep] subclasses. You can use the
// [imageUnfilteredFileTypes()] method of a given subclass to obtain the list
// of supported types for that class.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/init(pasteboard:)
//
// [imageUnfilteredFileTypes()]: https://developer.apple.com/documentation/AppKit/NSImageRep/imageUnfilteredFileTypes()
func (i NSImage) InitWithPasteboard(pasteboard INSPasteboard) NSImage {
	rv := objc.Send[NSImage](i.ID, objc.Sel("initWithPasteboard:"), pasteboard)
	return rv
}

// Initializes and returns an image object from data in an unarchiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/init(coder:)
func (i NSImage) InitWithCoder(coder foundation.INSCoder) NSImage {
	rv := objc.Send[NSImage](i.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// Initializes and returns an image object with the specified dimensions.
//
// size: The size of the image, measured in points.
//
// # Return Value
//
// An initialized [NSImage] object with no rendered content.
//
// # Discussion
//
// This method does not add any image representations to the image object. It
// is permissible to initialize the image object by passing a size of `(0.0,
// 0.0)`; however, you must set the size to a non-zero value before using it
// or an exception will be raised.
//
// After using this method to initialize an image object, you are expected to
// provide the image contents before trying to draw the image. You might lock
// focus on the image and draw to the image or you might explicitly add an
// image representation that you created.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/init(size:)
func (i NSImage) InitWithSize(size corefoundation.CGSize) NSImage {
	rv := objc.Send[NSImage](i.ID, objc.Sel("initWithSize:"), size)
	return rv
}

// Creates a new symbol image with the specified configuration.
//
// configuration: The configuration details to apply.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/withSymbolConfiguration(_:)
func (i NSImage) ImageWithSymbolConfiguration(configuration INSImageSymbolConfiguration) INSImage {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("imageWithSymbolConfiguration:"), configuration)
	return NSImageFromID(rv)
}

// Adds the specified image representation object to the image.
//
// imageRep: The image representation to add.
//
// # Discussion
//
// After invoking this method, you may need to explicitly set features of the
// new image representation, such as the size, number of colors, and so on.
// This is true particularly when the [NSImage] object has multiple image
// representations to choose from. See [NSImageRep] and its subclasses for the
// methods you use to complete initialization.
//
// Any representation added by this method is retained by the receiver. Image
// representations cannot be shared among multiple [NSImage] objects.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/addRepresentation(_:)
func (i NSImage) AddRepresentation(imageRep INSImageRep) {
	objc.Send[objc.ID](i.ID, objc.Sel("addRepresentation:"), imageRep)
}

// Adds an array of image representation objects to the image.
//
// imageReps: An array of [NSImageRep] objects.
//
// # Discussion
//
// After invoking this method, you may need to explicitly set features of the
// new image representations, such as their size, number of colors, and so on.
// This is true particularly when the [NSImage] object has multiple image
// representations to choose from. See [NSImageRep] and its subclasses for the
// methods you use to complete initialization.
//
// Representations added by this method are retained by the receiver. Image
// representations cannot be shared among multiple [NSImage] objects.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/addRepresentations(_:)
func (i NSImage) AddRepresentations(imageReps []NSImageRep) {
	objc.Send[objc.ID](i.ID, objc.Sel("addRepresentations:"), objectivec.IObjectSliceToNSArray(imageReps))
}

// Removes and releases the specified image representation.
//
// imageRep: The image representation object you want to remove.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/removeRepresentation(_:)
func (i NSImage) RemoveRepresentation(imageRep INSImageRep) {
	objc.Send[objc.ID](i.ID, objc.Sel("removeRepresentation:"), imageRep)
}

// Returns the best representation of the image for the specified rectangle
// using the provided hints.
//
// rect: The area of the image to return.
//
// referenceContext: A graphics context. This value can be `nil`.
//
// hints: An optional dictionary of hints that provide more context for selecting or
// generating a [CGImage], and may override properties of the
// `referenceContext`. See `Image Hint Dictionary Keys` for a summary of the
// possible key-value pairs.
//
// # Return Value
//
// The image representation that most closely matches the specified criteria.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/bestRepresentation(for:context:hints:)
func (i NSImage) BestRepresentationForRectContextHints(rect corefoundation.CGRect, referenceContext INSGraphicsContext, hints foundation.INSDictionary) INSImageRep {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("bestRepresentationForRect:context:hints:"), rect, referenceContext, hints)
	return NSImageRepFromID(rv)
}

// Draws the image in the specified rectangle.
//
// rect: The rectangle in which to draw the image, specified in the current
// coordinate system.
//
// # Discussion
//
// This method draws the entire image in the specified rectangle, scaling the
// image as needed. The method composites the image using the
// [NSCompositeSourceOver] operation
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/draw(in:)
//
// [NSCompositeSourceOver]: https://developer.apple.com/documentation/AppKit/NSCompositeSourceOver
func (i NSImage) DrawInRect(rect corefoundation.CGRect) {
	objc.Send[objc.ID](i.ID, objc.Sel("drawInRect:"), rect)
}

// Draws all or part of the image at the specified point in the current
// coordinate system.
//
// point: The location in the current coordinate system at which to draw the image.
//
// fromRect: The source rectangle specifying the portion of the image you want to draw.
// The coordinates of this rectangle are specified in the image’s own
// coordinate system. If you pass in [NSZeroRect], the entire image is drawn.
//
// op: The compositing operation to use when drawing the image. See the
// [NSCompositingOperation] constants.
//
// delta: The opacity of the image, specified as a value from 0.0 to 1.0. Specifying
// a value of 0.0 draws the image as fully transparent while a value of 1.0
// draws the image as fully opaque. Values greater than 1.0 are interpreted as
// 1.0.
//
// # Discussion
//
// The image content is drawn at its current resolution and is not scaled
// unless the CTM of the current coordinate system itself contains a scaling
// factor. The image is otherwise positioned and oriented using the current
// coordinate system.
//
// Unlike the [compositeToPoint:fromRect:operation:] and
// [compositeToPoint:fromRect:operation:fraction:] methods, this method checks
// the rectangle you pass to the `srcRect` parameter and makes sure it does
// not lie outside the image bounds.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/draw(at:from:operation:fraction:)
//
// [NSCompositingOperation]: https://developer.apple.com/documentation/AppKit/NSCompositingOperation
// [compositeToPoint:fromRect:operation:]: https://developer.apple.com/documentation/AppKit/NSImage/compositeToPoint:fromRect:operation:
// [compositeToPoint:fromRect:operation:fraction:]: https://developer.apple.com/documentation/AppKit/NSImage/compositeToPoint:fromRect:operation:fraction:
func (i NSImage) DrawAtPointFromRectOperationFraction(point corefoundation.CGPoint, fromRect corefoundation.CGRect, op NSCompositingOperation, delta float64) {
	objc.Send[objc.ID](i.ID, objc.Sel("drawAtPoint:fromRect:operation:fraction:"), point, fromRect, op, delta)
}

// Draws all or part of the image in the specified rectangle in the current
// coordinate system.
//
// rect: The rectangle in which to draw the image, specified in the current
// coordinate system.
//
// fromRect: The source rectangle specifying the portion of the image you want to draw.
// The coordinates of this rectangle must be specified using the image’s own
// coordinate system. If you pass in [NSZeroRect], the entire image is drawn.
//
// op: The compositing operation to use when drawing the image. See the
// [NSCompositingOperation] constants.
//
// delta: The opacity of the image, specified as a value from 0.0 to 1.0. Specifying
// a value of 0.0 draws the image as fully transparent while a value of 1.0
// draws the image as fully opaque. Values greater than 1.0 are interpreted as
// 1.0.
//
// # Discussion
//
// If the `srcRect` and `dstRect` rectangles have different sizes, the source
// portion of the image is scaled to fit the specified destination rectangle.
// The image is otherwise positioned and oriented using the current coordinate
// system.
//
// Unlike the [compositeToPoint:fromRect:operation:] and
// [compositeToPoint:fromRect:operation:fraction:] methods, this method checks
// the rectangle you pass to the `srcRect` parameter and makes sure it does
// not lie outside the image bounds.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/draw(in:from:operation:fraction:)
//
// [NSCompositingOperation]: https://developer.apple.com/documentation/AppKit/NSCompositingOperation
// [compositeToPoint:fromRect:operation:]: https://developer.apple.com/documentation/AppKit/NSImage/compositeToPoint:fromRect:operation:
// [compositeToPoint:fromRect:operation:fraction:]: https://developer.apple.com/documentation/AppKit/NSImage/compositeToPoint:fromRect:operation:fraction:
func (i NSImage) DrawInRectFromRectOperationFraction(rect corefoundation.CGRect, fromRect corefoundation.CGRect, op NSCompositingOperation, delta float64) {
	objc.Send[objc.ID](i.ID, objc.Sel("drawInRect:fromRect:operation:fraction:"), rect, fromRect, op, delta)
}

// Draws all or part of the image in the specified rectangle respecting the
// hints and the orientation of the current coordinate system.
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
// requestedAlpha: The alpha of the image, specified as a value from 0.0 to 1.0. Specifying a
// value of 0.0 draws the image as fully transparent while a value of 1.0
// draws the image as fully opaque. Values greater than 1.0 are interpreted as
// 1.0.
//
// respectContextIsFlipped: true if the drawing should respect the context flipped state, otherwise
// false.
//
// hints: An optional dictionary of hints that provide more context for selecting or
// generating the image. See `Image Hint Dictionary Keys` for a summary of the
// possible key-value pairs.
//
// # Discussion
//
// If the `srcSpacePortionRect` and `dstSpacePortionRect` rectangles have
// different sizes, the source portion of the image is scaled to fit the
// specified destination rectangle.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/draw(in:from:operation:fraction:respectFlipped:hints:)
//
// [NSCompositingOperation]: https://developer.apple.com/documentation/AppKit/NSCompositingOperation
func (i NSImage) DrawInRectFromRectOperationFractionRespectFlippedHints(dstSpacePortionRect corefoundation.CGRect, srcSpacePortionRect corefoundation.CGRect, op NSCompositingOperation, requestedAlpha float64, respectContextIsFlipped bool, hints foundation.INSDictionary) {
	objc.Send[objc.ID](i.ID, objc.Sel("drawInRect:fromRect:operation:fraction:respectFlipped:hints:"), dstSpacePortionRect, srcSpacePortionRect, op, requestedAlpha, respectContextIsFlipped, hints)
}

// Draws the image using the specified image representation object.
//
// imageRep: The image representation object to be drawn.
//
// rect: The rectangle in which to draw the image representation, specified in the
// current coordinate system.
//
// # Return Value
//
// true if the image was successfully drawn; otherwise, returns false.
//
// # Discussion
//
// This method fills the specified rectangle with the image’s current
// background color and then sends a message to the specified image
// representation asking if to draw itself. If the image supports the ability
// to scale itself when it is resized, this method sends a [DrawInRect]
// message; otherwise, it sends a [DrawAtPoint] message.
//
// You should not call this method directly; an [NSImage] object uses it to
// cache and print its image representations. You can override this method to
// change the way images are rendered into their caches and onto the printed
// page. For example, you could scale or rotate the coordinate system before
// sending this message to `super` to continue rendering the image
// representation.
//
// If the background color is fully transparent and the image data is not
// being cached, the specified rectangle is not to be filled before the
// representation draws.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/drawRepresentation(_:in:)
func (i NSImage) DrawRepresentationInRect(imageRep INSImageRep, rect corefoundation.CGRect) bool {
	rv := objc.Send[bool](i.ID, objc.Sel("drawRepresentation:inRect:"), imageRep, rect)
	return rv
}

// Invalidates and frees offscreen caches of all image representations.
//
// # Discussion
//
// If you modify an image representation, you must send a [Recache] message to
// the corresponding image object to force the changes to be recached. The
// next time any image representation is drawn, it is asked to recreate its
// cached image. If you do not send this message, the image representation may
// use the old cache data. This method simply clears the cached image data; it
// does not delete the [NSCachedImageRep] objects associated with any image
// representations.
//
// If you do not plan to use an image again right away, you can free its
// caches to reduce the amount of memory consumed by your program.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/recache()
func (i NSImage) Recache() {
	objc.Send[objc.ID](i.ID, objc.Sel("recache"))
}

// Returns a data object that contains TIFF data with the specified
// compression settings for all of the image representations in the image.
//
// comp: The type of compression to use. For a list of values, see the constants in
// [NSBitmapImageRep].
//
// factor: Provides a hint for compression types that implement variable compression
// ratios. Currently, only JPEG compression uses a compression factor.
//
// # Return Value
//
// A data object containing the TIFF data, or `nil` if the TIFF data could not
// be created.
//
// # Discussion
//
// You can use the returned data object to write the TIFF data to a file. If
// the specified compression isn’t applicable, no compression is used. If a
// problem is encountered during generation of the TIFF data, this method may
// raise an exception.
//
// If one of the receiver’s image representations does not support the
// creation of TIFF data natively (PDF and EPS images, for example), this
// method creates the TIFF data from that representation’s cached content.
//
// Additional image formats can be saved by using the [NSBitmapImageRep]
// method [RepresentationUsingTypeProperties].
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/tiffRepresentation(using:factor:)
func (i NSImage) TIFFRepresentationUsingCompressionFactor(comp NSTIFFCompression, factor float32) foundation.INSData {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("TIFFRepresentationUsingCompression:factor:"), comp, factor)
	return foundation.NSDataFromID(rv)
}

// Returns a Core Graphics image based on the contents of the current image
// object.
//
// proposedDestRect: On input, the proposed destination rectangle for drawing the image. If
// [NULL], it defaults to the smallest pixel-integral rectangle containing
// {{0,0}, [self size]}. The `proposedDestRect` is in user space in the
// reference context.
//
// referenceContext: A graphics context.
//
// hints: A dictionary of hints that provide more context for selecting or generating
// a [CGImage], and may override properties of the `referenceContext`.
//
// # Return Value
//
// A [CGImageRef]. This may be an existing [CGImage] if one is available. If
// not, a new [CGImage] is created.
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
// This method is typically called, not overridden.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/cgImage(forProposedRect:context:hints:)
//
// [CGImage]: https://developer.apple.com/documentation/CoreGraphics/CGImage
func (i NSImage) CGImageForProposedRectContextHints(proposedDestRect *corefoundation.CGRect, referenceContext INSGraphicsContext, hints foundation.INSDictionary) coregraphics.CGImageRef {
	rv := objc.Send[coregraphics.CGImageRef](i.ID, objc.Sel("CGImageForProposedRect:context:hints:"), proposedDestRect, referenceContext, hints)
	return coregraphics.CGImageRef(rv)
}

// Returns whether the destination rectangle would intersect a non-transparent
// portion of the image.
//
// testRectDestSpace: The rectangle to hit test.
//
// imageRectDestSpace: A rectangle representing the drawn size of the image.
//
// context: A graphics context. This value can be `nil`.
//
// hints: An optional dictionary of hints that provide more context for selecting or
// generating a [CGImage], and may override properties of the
// `referenceContext`. See `Image Hint Dictionary Keys` for a summary of the
// possible key-value pairs.
//
// flipped: true if the image is flipped, otherwise false.
//
// # Return Value
//
// YES if the `testRectDestSpace` intersects with non-transparent content
// within the `imageRectDestSpace`, otherwise NO.
//
// # Discussion
//
// This method simulates the results of hit-testing the test rectangle as if
// the image was drawn in the graphics context using the provided hints and
// respecting the specified flippedness.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/hitTest(_:withDestinationRect:context:hints:flipped:)
func (i NSImage) HitTestRectWithImageDestinationRectContextHintsFlipped(testRectDestSpace corefoundation.CGRect, imageRectDestSpace corefoundation.CGRect, context INSGraphicsContext, hints foundation.INSDictionary, flipped bool) bool {
	rv := objc.Send[bool](i.ID, objc.Sel("hitTestRect:withImageDestinationRect:context:hints:flipped:"), testRectDestSpace, imageRectDestSpace, context, hints, flipped)
	return rv
}

// Returns an object that may be used as the contents of a layer.
//
// layerContentsScale: The scale factor for the resulting image. Obtain the value for this
// parameter by calling the [RecommendedLayerContentsScale] method.
//
// # Return Value
//
// A object that you can assign to the [Contents] property of a [CALayer]
// object. This object contains the image data from the current image
// optimized for the specified scale factor.
//
// # Discussion
//
// Use this method in situations where you want to use the image as the
// contents of a layer. This method provides the image data wrapped in an
// object that correctly respects all of the possible content gravities
// supported by the layer. Use of the returned object as the layer’s
// contents is recommended over the use of the [NSImage] object itself.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/layerContents(forContentsScale:)
//
// [CALayer]: https://developer.apple.com/documentation/QuartzCore/CALayer
func (i NSImage) LayerContentsForContentsScale(layerContentsScale float64) objectivec.IObject {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("layerContentsForContentsScale:"), layerContentsScale)
	return objectivec.Object{ID: rv}
}

// Returns the recommended layer contents scale for this image.
//
// preferredContentsScale: The preferred layer contents scale. Don’t use a higher scale factor if
// the image can’t provide it. If the image is resolution independent the
// return value will be the same as the input. If you specify `0.0` for this
// parameter, the method uses the scale factor for the default screen.
//
// # Return Value
//
// The recommended layer contents scale. This scale factor may be different
// than the one in the `preferredContentsScale` parameter.
//
// # Discussion
//
// Use this method to obtain the scale factor value that you pass to the
// [LayerContentsForContentsScale] method. This method uses the image data to
// determine the scale factor that is best suited for creating an image that
// looks good in the layer.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/recommendedLayerContentsScale(_:)
func (i NSImage) RecommendedLayerContentsScale(preferredContentsScale float64) float64 {
	rv := objc.Send[float64](i.ID, objc.Sel("recommendedLayerContentsScale:"), preferredContentsScale)
	return rv
}

// See: https://developer.apple.com/documentation/AppKit/NSImage/withLocale(_:)
func (i NSImage) ImageWithLocale(locale foundation.NSLocale) INSImage {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("imageWithLocale:"), locale)
	return NSImageFromID(rv)
}

// Initializes an instance with a property list object and a type string.
//
// propertyList: A property list containing data to initialize the receiver.
//
// By default, the property list object is an instance of [NSData]. If you
// implement [ReadingOptionsForTypePasteboard] and specify an option other
// than [NSPasteboardReadingAsData], the `propertyList` may be any other
// property list object.
//
// type: A UTI supported by the receiver for reading (one of the types returned by
// [ReadableTypesForPasteboard]).
//
// # Return Value
//
// An object initialized using the data in `propertyList`.
//
// # Discussion
//
// This method is considered optional because, if [ReadableTypesForPasteboard]
// returns just a single type, and that type uses the
// [NSPasteboardReadingAsKeyedArchive] reading option, then instances are
// initialized using [init(coder:)] instead of this method.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardReading/init(pasteboardPropertyList:ofType:)
//
// [init(coder:)]: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func (i NSImage) InitWithPasteboardPropertyListOfType(propertyList objectivec.IObject, type_ NSPasteboardType) NSImage {
	rv := objc.Send[NSImage](i.ID, objc.Sel("initWithPasteboardPropertyList:ofType:"), propertyList, objc.String(string(type_)))
	return rv
}

// Returns a property list object to represent the receiver on a pasteboard as
// an object of a specified type.
//
// type: One of the types the receiver supports for writing (one of the UTIs
// returned by its implementation of [WritableTypesForPasteboard]).
//
// # Return Value
//
// A property list object to represent the receiver on a pasteboard as an
// object of type `type`.
//
// # Discussion
//
// The returned value will commonly be the [NSData] object for the specified
// data type. However, if this method returns either a string, or any other
// property-list type, the pasteboard will automatically convert these items
// to the correct data format required for the pasteboard.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardWriting/pasteboardPropertyList(forType:)
func (i NSImage) PasteboardPropertyListForType(type_ NSPasteboardType) objectivec.IObject {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("pasteboardPropertyListForType:"), objc.String(string(type_)))
	return objectivec.Object{ID: rv}
}

// Returns an array of UTI strings of data types the receiver can write to a
// given pasteboard.
//
// pasteboard: A pasteboard.
//
// You can use this argument to provide different options based on the
// pasteboard name, if you need to.
//
// # Return Value
//
// An array of UTI strings of data types the receiver can write to
// `pasteboard`.
//
// # Discussion
//
// By default, data for the first returned type is put onto the pasteboard
// immediately, with the remaining types being promised.
//
// To change the default behavior, implement
// -writingOptionsForType:pasteboard: and return [NSPasteboardWritingPromised]
// to lazily provide data for types, return no option to provide the data for
// that type immediately. Use the pasteboard argument to provide different
// types based on the pasteboard name, if desired. Do not perform other
// pasteboard operations in the method implementation.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardWriting/writableTypes(for:)
func (i NSImage) WritableTypesForPasteboard(pasteboard INSPasteboard) []string {
	rv := objc.Send[[]objc.ID](i.ID, objc.Sel("writableTypesForPasteboard:"), pasteboard)
	return objc.ConvertSliceToStrings(rv)
}

// Returns options for writing data of a specified type to a given pasteboard.
//
// type: One of the types the receiver supports for writing (one of the UTIs
// returned by its implementation of [WritableTypesForPasteboard]).
//
// pasteboard: A pasteboard.
//
// You can use this argument to provide different options based on the
// pasteboard name, if you need to.
//
// # Return Value
//
// Options for writing data of type type to `pasteboard`. Return `0` for no
// options, or a value given in [Pasteboard Writing Options].
//
// # Discussion
//
// Do not perform other pasteboard operations in the method implementation.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardWriting/writingOptions(forType:pasteboard:)
//
// [Pasteboard Writing Options]: https://developer.apple.com/documentation/AppKit/pasteboard-writing-options
func (i NSImage) WritingOptionsForTypePasteboard(type_ NSPasteboardType, pasteboard INSPasteboard) NSPasteboardWritingOptions {
	rv := objc.Send[NSPasteboardWritingOptions](i.ID, objc.Sel("writingOptionsForType:pasteboard:"), objc.String(string(type_)), pasteboard)
	return NSPasteboardWritingOptions(rv)
}
func (i NSImage) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](i.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Creates and returns an image object whose contents are drawn using the
// specified block.
//
// size: The size of the image, measured in points.
//
// drawingHandlerShouldBeCalledWithFlippedContext: true to apply a flip transformation to the graphics context before drawing
// or false to draw using the default Cocoa coordinate system orientation.
//
// drawingHandler: A block that draws the contents of the image representation. The image
// representation copies the block and stores it for later use. It is not
// executed until you draw the image. Because the block is executed later,
// AppKit executes it on the same thread on which you draw the image itself,
// which can be any thread of your app. Therefore, the block must be safe to
// call from any thread. The block takes the following parameter:
//
// dstRect: The destination rectangle in which to draw. The coordinates of
// this rectangle are specified in points.
//
// The block returns a Boolean that indicates whether it drew the image
// successfully. Return true from your block if it successfully drew the
// contents or false if it did not.
//
// # Return Value
//
// An initialized [NSImage] object.
//
// # Discussion
//
// Use this method to generate an image that is correct at any resolution.
// This method creates an image object with a single [NSCustomImageRep] object
// to manage drawing. The image representation uses the block in the
// `drawingHandler` parameter to do the actual drawing.
//
// When you draw the image for the first time, the underlying image
// representation executes the `drawingHandler` block. The image object caches
// the results according to its usual caching policies; see the [CacheMode]
// property. As long as the configuration of the destination graphics context
// does not change in a significant way, subsequent attempts to draw the image
// reuse the cached results whenever possible. If the pixel density or color
// space of the destination graphics context changes, though, the image
// representation throws away any caches and executes the block again to
// obtain a new version of the image. For example, if you drew the image on a
// standard resolution display but then draw it on a Retina display, AppKit
// executes the block again to obtain an image at the new resolution.
//
// The most typical use for this method is to create an image based on
// vector-based content. In that case, your `drawingHandler` block would
// redraw its existing path objects when asked. If you draw a mixture of
// vectors and images, you need to do more work to ensure that your images are
// the appropriate resolution for the destination graphics context.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/init(size:flipped:drawingHandler:)
func (_NSImageClass NSImageClass) ImageWithSizeFlippedDrawingHandler(size corefoundation.CGSize, drawingHandlerShouldBeCalledWithFlippedContext ErrorHandler, drawingHandler RectHandler) NSImage {
	_block1, _ := NewErrorBlock(drawingHandlerShouldBeCalledWithFlippedContext)
	_block2, _ := NewRectBlock(drawingHandler)
	rv := objc.Send[objc.ID](objc.ID(_NSImageClass.class), objc.Sel("imageWithSize:flipped:drawingHandler:"), size, _block1, _block2)
	return NSImageFromID(rv)
}

// Tests whether the image can create an instance of itself using pasteboard
// data.
//
// pasteboard: The pasteboard containing the image data.
//
// # Return Value
//
// true if the receiver knows how to handle the data on the pasteboard;
// otherwise, false.
//
// # Discussion
//
// This method uses the [NSImageRep] class method
// [imageUnfilteredPasteboardTypes()] to find a class that can handle the data
// in the specified pasteboard. If you create your own [NSImageRep]
// subclasses, override the [imageUnfilteredPasteboardTypes()] method to
// notify [NSImage] of the pasteboard types your class supports.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/canInit(with:)
//
// [imageUnfilteredPasteboardTypes()]: https://developer.apple.com/documentation/AppKit/NSImageRep/imageUnfilteredPasteboardTypes()
func (_NSImageClass NSImageClass) CanInitWithPasteboard(pasteboard INSPasteboard) bool {
	rv := objc.Send[bool](objc.ID(_NSImageClass.class), objc.Sel("canInitWithPasteboard:"), pasteboard)
	return rv
}

// Returns an array of uniform type identifier strings of data types the
// receiver can read from the pasteboard and initialize from.
//
// pasteboard: A pasteboard. You can use the pasteboard argument to provide different
// types based on the pasteboard name, if you need to.
//
// # Return Value
//
// An array of uniform type identifier strings of data types instances that
// the receiver can read from the pasteboard and initialize from.
//
// # Discussion
//
// By default, the system provides the data for a type to
// [InitWithPasteboardPropertyListOfType] as an instance of [NSData]. If you
// implement [ReadingOptionsForTypePasteboard] and specify a different option,
// the system converts the [NSData] object for a type to an [NSString] object
// or any other property list object.
//
// # Special Considerations
//
// Don’t perform other pasteboard operations in the method implementation.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardReading/readableTypes(for:)
func (_NSImageClass NSImageClass) ReadableTypesForPasteboard(pasteboard INSPasteboard) []string {
	rv := objc.Send[[]objc.ID](objc.ID(_NSImageClass.class), objc.Sel("readableTypesForPasteboard:"), pasteboard)
	return objc.ConvertSliceToStrings(rv)
}

// Returns options for reading data of a specified type from a given
// pasteboard.
//
// type: A UTI supported by instances of the receiver for reading (one of the types
// returned by [ReadableTypesForPasteboard]).
//
// pasteboard: A pasteboard.
//
// You can use the pasteboard argument to provide return different based on
// the pasteboard name, should you need to do so.
//
// # Return Value
//
// Options for reading data of `type` from `pasteboard`. For a list of valid
// values, see [NSPasteboard.ReadingOptions].
//
// # Discussion
//
// Do not perform other pasteboard operations in this method implementation.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardReading/readingOptions(forType:pasteboard:)
//
// [NSPasteboard.ReadingOptions]: https://developer.apple.com/documentation/AppKit/NSPasteboard/ReadingOptions
func (_NSImageClass NSImageClass) ReadingOptionsForTypePasteboard(type_ NSPasteboardType, pasteboard INSPasteboard) NSPasteboardReadingOptions {
	rv := objc.Send[NSPasteboardReadingOptions](objc.ID(_NSImageClass.class), objc.Sel("readingOptionsForType:pasteboard:"), objc.String(string(type_)), pasteboard)
	return NSPasteboardReadingOptions(rv)
}

// The configuration details for a symbol image.
//
// # Discussion
//
// Use this property to access the traits and rendering attributes the system
// uses with the symbol image. These details determine which variant of the
// image to load and draw and how to render it, falling back on the current
// environment for values that you don’t specify. For symbol images, the
// default value of this property is a symbol image configuration object with
// unspecified values.
//
// You can’t modify this property directly, but you can use
// [ImageWithSymbolConfiguration] when you want to create a new image object
// with a specific set of traits.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/symbolConfiguration-swift.property
func (i NSImage) SymbolConfiguration() INSImageSymbolConfiguration {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("symbolConfiguration"))
	return NSImageSymbolConfigurationFromID(objc.ID(rv))
}

// The image’s delegate object.
//
// # Discussion
//
// By default, this property contains `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/delegate
func (i NSImage) Delegate() NSImageDelegate {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("delegate"))
	return NSImageDelegateObjectFromID(rv)
}
func (i NSImage) SetDelegate(value NSImageDelegate) {
	objc.Send[struct{}](i.ID, objc.Sel("setDelegate:"), value)
}

// The size of the image.
//
// # Discussion
//
// Defaults to `{0.0, 0.0}` if no size has been set and the size cannot be
// determined from any of the receiver’s image representations. If the size
// of the image hasn’t already been set when an image representation is
// added, the size is taken from the image representation’s data. For EPS
// images, the size is taken from the image’s bounding box. For TIFF images,
// the size is taken from the [ImageLength] and [ImageWidth] attributes.
//
// Changing the size of an [NSImage] after it has been used effectively
// resizes the image. Changing the size invalidates all its caches and frees
// them. When the image is next composited, the selected representation will
// draw itself in an offscreen window to recreate the cache.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/size
func (i NSImage) Size() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](i.ID, objc.Sel("size"))
	return corefoundation.CGSize(rv)
}
func (i NSImage) SetSize(value corefoundation.CGSize) {
	objc.Send[struct{}](i.ID, objc.Sel("setSize:"), value)
}

// A Boolean value that determines whether the image represents a template
// image.
//
// # Discussion
//
// Images you mark as template images should consist of only black and clear
// colors. You can use the alpha channel in the image to adjust the opacity of
// black content, however.
//
// Template images are not intended to be used as standalone images. They are
// always mixed with other content and processed to create the desired
// appearance. You can mark an image as a “template image” to notify
// clients who care that the image contains only black and clear content. The
// most common use for template images is in image cells. For example, you
// might use a template image to provide the content for a button or segmented
// control. Cocoa cells take advantage of the nature of template images—that
// is, their simplified color scheme and use of transparency—to improve the
// appearance of the corresponding control in each of its supported states.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/isTemplate
func (i NSImage) Template() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("isTemplate"))
	return rv
}
func (i NSImage) SetTemplate(value bool) {
	objc.Send[struct{}](i.ID, objc.Sel("setTemplate:"), value)
}

// An array containing all of the image object’s image representations.
//
// # Discussion
//
// This property can contain zero or more [NSImageRep] objects.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/representations
func (i NSImage) Representations() []NSImageRep {
	rv := objc.Send[[]objc.ID](i.ID, objc.Sel("representations"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSImageRep {
		return NSImageRepFromID(id)
	})
}

// A Boolean value that indicates whether the image prefers to choose image
// representations using color-matching or resolution-matching.
//
// # Discussion
//
// When the value of this property is true, the image attempts to match the
// color capabilities of the rendering device first. When it is false, the
// image prefers resolution-matching first. The default value is true. Both
// color-matching and resolution-matching may influence the choice of an image
// representation. You use this method to choose which technique should be
// used first during the selection process.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/prefersColorMatch
func (i NSImage) PrefersColorMatch() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("prefersColorMatch"))
	return rv
}
func (i NSImage) SetPrefersColorMatch(value bool) {
	objc.Send[struct{}](i.ID, objc.Sel("setPrefersColorMatch:"), value)
}

// A Boolean value that indicates whether EPS representations are preferred
// when no other representations match the resolution of the device.
//
// # Discussion
//
// The default value is false.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/usesEPSOnResolutionMismatch
func (i NSImage) UsesEPSOnResolutionMismatch() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("usesEPSOnResolutionMismatch"))
	return rv
}
func (i NSImage) SetUsesEPSOnResolutionMismatch(value bool) {
	objc.Send[struct{}](i.ID, objc.Sel("setUsesEPSOnResolutionMismatch:"), value)
}

// A Boolean value that indicates whether image representations whose
// resolution is an integral multiple of the device resolution are a match.
//
// # Discussion
//
// When this property is set to false, only image representations whose
// resolution is exactly the same as the device resolution are matches. If the
// property is set to true and multiple image representations fit this
// criteria, the one whose resolution is closest to the device resolution is
// chosen.
//
// The default value is true.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/matchesOnMultipleResolution
func (i NSImage) MatchesOnMultipleResolution() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("matchesOnMultipleResolution"))
	return rv
}
func (i NSImage) SetMatchesOnMultipleResolution(value bool) {
	objc.Send[struct{}](i.ID, objc.Sel("setMatchesOnMultipleResolution:"), value)
}

// A Boolean value that indicates whether it is possible to draw an image
// representation.
//
// # Discussion
//
// If you created the image with an existing image file, but the corresponding
// image data is not yet loaded into memory, this method loads the data and
// expands it as needed. If the receiver contains no image representations and
// no associated image file, this method creates a valid cached image
// representation and initializes it to the default bit depth. If the file or
// URL from which the image was initialized is nonexistent, or the data in an
// existing file is invalid, this method returns false.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/isValid
func (i NSImage) Valid() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("isValid"))
	return rv
}

// The background color for the image.
//
// # Discussion
//
// The background color is visible only if the drawn image representation does
// not completely cover all of the pixels available for the image’s current
// size. The background color is ignored for cached image representations;
// such caches are always created with a white background. Assigning a new
// background color does not cause the receiver to recache itself.
//
// The default color is transparent, as returned by the [ClearColor] method of
// [NSColor].
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/backgroundColor
func (i NSImage) BackgroundColor() INSColor {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("backgroundColor"))
	return NSColorFromID(objc.ID(rv))
}
func (i NSImage) SetBackgroundColor(value INSColor) {
	objc.Send[struct{}](i.ID, objc.Sel("setBackgroundColor:"), value)
}

// The cap insets for the image.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/capInsets
func (i NSImage) CapInsets() foundation.NSEdgeInsets {
	rv := objc.Send[foundation.NSEdgeInsets](i.ID, objc.Sel("capInsets"))
	return foundation.NSEdgeInsets(rv)
}
func (i NSImage) SetCapInsets(value foundation.NSEdgeInsets) {
	objc.Send[struct{}](i.ID, objc.Sel("setCapInsets:"), value)
}

// The resizing mode for the image.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/resizingMode-swift.property
func (i NSImage) ResizingMode() NSImageResizingMode {
	rv := objc.Send[NSImageResizingMode](i.ID, objc.Sel("resizingMode"))
	return NSImageResizingMode(rv)
}
func (i NSImage) SetResizingMode(value NSImageResizingMode) {
	objc.Send[struct{}](i.ID, objc.Sel("setResizingMode:"), value)
}

// A rectangle that you can use to position the image during layout.
//
// # Discussion
//
// Alignment rectangles specify baselines that you can use to position the
// content of an image more accurately. These baselines are merely hints that
// your own code can use to determine positioning. The [NSImage] class does
// not use this rectangle during drawing; however, instances of [NSCell]
// typically use this information when laying out images within their
// boundaries.
//
// For example, if you have a 20 x 20 pixel icon that includes a glow effect,
// you might set the alignment rectangle to `{{2, 2}, {16, 16}}` to indicate
// the position of the underlying icon without the glow effect. This property
// defaults to a rectangle with an origin of `{0, 0}` and a size that matches
// the size of the image.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/alignmentRect
func (i NSImage) AlignmentRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](i.ID, objc.Sel("alignmentRect"))
	return corefoundation.CGRect(rv)
}
func (i NSImage) SetAlignmentRect(value corefoundation.CGRect) {
	objc.Send[struct{}](i.ID, objc.Sel("setAlignmentRect:"), value)
}

// The image’s caching mode.
//
// # Discussion
//
// The caching mode determines when the image representations use offscreen
// caches. Offscreen caches speed up rendering time but do so by using extra
// memory. In the default caching mode ([NSImageCacheDefault]), each image
// representation chooses the caching technique that produces the fastest
// drawing times. For example, in the default mode, the [NSPDFImageRep] and
// [NSEPSImageRep] classes use the [NSImageCacheAlways] mode but the
// [NSBitmapImageRep] class uses the [NSImageCacheBySize] mode. For a list of
// possible values, see [NSImage.CacheMode]. This value is set to
// [NSImageCacheDefault] by default.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/cacheMode-swift.property
//
// [NSImage.CacheMode]: https://developer.apple.com/documentation/AppKit/NSImage/CacheMode-swift.enum
func (i NSImage) CacheMode() NSImageCacheMode {
	rv := objc.Send[NSImageCacheMode](i.ID, objc.Sel("cacheMode"))
	return NSImageCacheMode(rv)
}
func (i NSImage) SetCacheMode(value NSImageCacheMode) {
	objc.Send[struct{}](i.ID, objc.Sel("setCacheMode:"), value)
}

// A data object containing TIFF data for all of the image representations in
// the image.
//
// # Discussion
//
// Use the value of this property to write the TIFF data to a file. For each
// image representation, this property uses the TIFF compression option
// associated with that representation or [NSTIFFCompressionNone], if no
// option is set.
//
// If one of the receiver’s image representations does not support the
// creation of TIFF data natively (PDF and EPS images, for example), this
// property creates the TIFF data from that representation’s cached content.
// This property contains `nil` if the TIFF data cannot be created.
//
// Additional image formats can be saved by using the [NSBitmapImageRep]
// method [RepresentationUsingTypeProperties].
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/tiffRepresentation
func (i NSImage) TIFFRepresentation() foundation.INSData {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("TIFFRepresentation"))
	return foundation.NSDataFromID(objc.ID(rv))
}

// The image’s accessibility description.
//
// # Discussion
//
// This description is used automatically by interface elements that display
// images. Like all accessibility descriptions, use a short localized string
// that does not include the name of the interface element. For instance,
// “delete” rather than “delete button”.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/accessibilityDescription
func (i NSImage) AccessibilityDescription() string {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("accessibilityDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (i NSImage) SetAccessibilityDescription(value string) {
	objc.Send[struct{}](i.ID, objc.Sel("setAccessibilityDescription:"), objc.String(value))
}

// A Boolean value that indicates whether the image matches only on the best
// fitting axis.
//
// # Discussion
//
// true if the image is drawn only on the best fitting axis; otherwise, false.
// This property defaults to false.
//
// NSImage has always tried to use a representation with at least as many
// pixels as the destination rectangle. Many apps try to implement banners and
// 3 part / 9 part images by stretching an NSImage over a much larger area
// (usually only on a single axis).
//
// With the addition of 2x assets these apps are finding this policy displays
// the 2x image rep when they would prefer the 1x rep. This behavior can be
// changed by using this method.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/matchesOnlyOnBestFittingAxis
func (i NSImage) MatchesOnlyOnBestFittingAxis() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("matchesOnlyOnBestFittingAxis"))
	return rv
}
func (i NSImage) SetMatchesOnlyOnBestFittingAxis(value bool) {
	objc.Send[struct{}](i.ID, objc.Sel("setMatchesOnlyOnBestFittingAxis:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSImage/locale
func (i NSImage) Locale() foundation.NSLocale {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("locale"))
	return foundation.NSLocaleFromID(objc.ID(rv))
}

// An object that provides the contents of the layer. Animatable.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/contents
func (i NSImage) Contents() objectivec.IObject {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("contents"))
	return objectivec.Object{ID: rv}
}
func (i NSImage) SetContents(value objectivec.IObject) {
	objc.Send[struct{}](i.ID, objc.Sel("setContents:"), value)
}

// A constant that specifies how the layer’s contents are positioned or
// scaled within its bounds.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/contentsGravity
func (i NSImage) ContentsGravity() foundation.NSString {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("contentsGravity"))
	return foundation.NSStringFromID(objc.ID(rv))
}
func (i NSImage) SetContentsGravity(value foundation.NSString) {
	objc.Send[struct{}](i.ID, objc.Sel("setContentsGravity:"), value)
}

// The content is resized to fit the entire bounds rectangle.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayerContentsGravity/resize
func (i NSImage) Resize() foundation.NSString {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("resize"))
	return foundation.NSStringFromID(objc.ID(rv))
}

// The content is resized to fit the bounds rectangle, preserving the aspect
// of the content. If the content does not completely fill the bounds
// rectangle, the content is centered in the partial axis.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayerContentsGravity/resizeAspect
func (i NSImage) ResizeAspect() foundation.NSString {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("resizeAspect"))
	return foundation.NSStringFromID(objc.ID(rv))
}

// The content is resized to completely fill the bounds rectangle, while still
// preserving the aspect of the content. The content is centered in the axis
// it exceeds.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayerContentsGravity/resizeAspectFill
func (i NSImage) ResizeAspectFill() foundation.NSString {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("resizeAspectFill"))
	return foundation.NSStringFromID(objc.ID(rv))
}

// Returns an array of UTI strings identifying the image types supported by
// the registered image representation objects, either directly or through a
// user-installed filter service.
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
// The returned list includes UTIs all file types supported by registered
// subclasses of [NSImageRep] plus those that can be converted to a supported
// type by a user-installed filter service. You can use the returned UTI
// strings with any method that supports UTIs.
//
// Do not override this method directly. If your app supports custom image
// types, create and register an [NSImageRep] subclass that handles those
// types.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/imageTypes
func (_NSImageClass NSImageClass) ImageTypes() []string {
	rv := objc.Send[[]objc.ID](objc.ID(_NSImageClass.class), objc.Sel("imageTypes"))
	return objc.ConvertSliceToStrings(rv)
}

// Returns an array of UTI strings identifying the image types supported
// directly by the registered image representation objects.
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
// supported directly by registered subclasses of [NSImageRep]. It does not
// include types that are supported through user-installed filter services.
// You can use the returned UTI strings with any method that supports UTIs.
//
// Do not override this method directly. If your app supports custom image
// types, create and register an [NSImageRep] subclass that handles those
// types.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/imageUnfilteredTypes
func (_NSImageClass NSImageClass) ImageUnfilteredTypes() []string {
	rv := objc.Send[[]objc.ID](objc.ID(_NSImageClass.class), objc.Sel("imageUnfilteredTypes"))
	return objc.ConvertSliceToStrings(rv)
}

// Protocol methods for NSPasteboardWriting

// ImageWithSizeFlippedDrawingHandlerSync is a synchronous wrapper around [NSImage.ImageWithSizeFlippedDrawingHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (ic NSImageClass) ImageWithSizeFlippedDrawingHandlerSync(ctx context.Context, size corefoundation.CGSize, drawingHandlerShouldBeCalledWithFlippedContext ErrorHandler) (corefoundation.CGRect, error) {
	done := make(chan corefoundation.CGRect, 1)
	ic.ImageWithSizeFlippedDrawingHandler(size, drawingHandlerShouldBeCalledWithFlippedContext, func(val corefoundation.CGRect) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return corefoundation.CGRect{}, ctx.Err()
	}
}
