// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSTextPreview] class.
var (
	_NSTextPreviewClass     NSTextPreviewClass
	_NSTextPreviewClassOnce sync.Once
)

func getNSTextPreviewClass() NSTextPreviewClass {
	_NSTextPreviewClassOnce.Do(func() {
		_NSTextPreviewClass = NSTextPreviewClass{class: objc.GetClass("NSTextPreview")}
	})
	return _NSTextPreviewClass
}

// GetNSTextPreviewClass returns the class object for NSTextPreview.
func GetNSTextPreviewClass() NSTextPreviewClass {
	return getNSTextPreviewClass()
}

type NSTextPreviewClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSTextPreviewClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTextPreviewClass) Alloc() NSTextPreview {
	rv := objc.Send[NSTextPreview](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A snapshot of the text in your view, which the system uses to create
// user-visible effects.
//
// # Overview
// 
// An [NSTextPreview] object provides a static image of your view’s text
// content that the system can use to create animations. You provide preview
// objects in response to system requests, such as ones from Writing Tools. In
// addition to creating an image of your view’s text, you also specify the
// location of that text in your view’s frame rectangle. When creating
// animations, the system places the image on top of your view’s content and
// animates changes to the image instead of to your view.
// 
// Create an [NSTextPreview] object in response to specific system requests.
// Create an image with a transparent background and render your view’s text
// into the image using the current text attributes. Construct your
// [NSTextPreview] object with both the image and the frame rectangle that
// represents the location of the rendered text in your view’s coordinate
// system. To highlight specific portions of text, instead of all the text in
// the image, provide a set of candidate rectangles with the locations of the
// text you want to highlight.
//
// # Creating a text preview
//
//   - [NSTextPreview.InitWithSnapshotImagePresentationFrame]: Creates a text preview using the specified image.
//   - [NSTextPreview.InitWithSnapshotImagePresentationFrameCandidateRects]: Creates a text preview using the specified image and rectangles that indicate the portions of text to highlight.
//
// # Getting the preview details
//
//   - [NSTextPreview.PreviewImage]: The image that contains the requested text from your view.
//   - [NSTextPreview.PresentationFrame]: The frame rectangle that places the preview image directly over the matching text.
//   - [NSTextPreview.CandidateRects]: Rectangles that define the specific portions of text to highlight.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextPreview
type NSTextPreview struct {
	objectivec.Object
}

// NSTextPreviewFromID constructs a [NSTextPreview] from an objc.ID.
//
// A snapshot of the text in your view, which the system uses to create
// user-visible effects.
func NSTextPreviewFromID(id objc.ID) NSTextPreview {
	return NSTextPreview{objectivec.Object{ID: id}}
}
// NOTE: NSTextPreview adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSTextPreview] class.
//
// # Creating a text preview
//
//   - [INSTextPreview.InitWithSnapshotImagePresentationFrame]: Creates a text preview using the specified image.
//   - [INSTextPreview.InitWithSnapshotImagePresentationFrameCandidateRects]: Creates a text preview using the specified image and rectangles that indicate the portions of text to highlight.
//
// # Getting the preview details
//
//   - [INSTextPreview.PreviewImage]: The image that contains the requested text from your view.
//   - [INSTextPreview.PresentationFrame]: The frame rectangle that places the preview image directly over the matching text.
//   - [INSTextPreview.CandidateRects]: Rectangles that define the specific portions of text to highlight.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextPreview
type INSTextPreview interface {
	objectivec.IObject

	// Topic: Creating a text preview

	// Creates a text preview using the specified image.
	InitWithSnapshotImagePresentationFrame(snapshotImage coregraphics.CGImageRef, presentationFrame corefoundation.CGRect) NSTextPreview
	// Creates a text preview using the specified image and rectangles that indicate the portions of text to highlight.
	InitWithSnapshotImagePresentationFrameCandidateRects(snapshotImage coregraphics.CGImageRef, presentationFrame corefoundation.CGRect, candidateRects []foundation.NSValue) NSTextPreview

	// Topic: Getting the preview details

	// The image that contains the requested text from your view.
	PreviewImage() coregraphics.CGImageRef
	// The frame rectangle that places the preview image directly over the matching text.
	PresentationFrame() corefoundation.CGRect
	// Rectangles that define the specific portions of text to highlight.
	CandidateRects() []foundation.NSValue
}

// Init initializes the instance.
func (t NSTextPreview) Init() NSTextPreview {
	rv := objc.Send[NSTextPreview](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTextPreview) Autorelease() NSTextPreview {
	rv := objc.Send[NSTextPreview](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTextPreview creates a new NSTextPreview instance.
func NewNSTextPreview() NSTextPreview {
	class := getNSTextPreviewClass()
	rv := objc.Send[NSTextPreview](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a text preview using the specified image.
//
// snapshotImage: An image that contains the requested text from your view. Create the image
// using a transparent background and the current rendering attributes for
// your text.
//
// presentationFrame: A rectangle in your frame’s coordinate space. The system uses this
// rectangle to place your image precisely over your view’s actual text. Set
// its size to the size of your snapshot image, and set its origin to the
// point that allows the system to place your image directly over the text.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextPreview/init(snapshotImage:presentationFrame:)
func NewTextPreviewWithSnapshotImagePresentationFrame(snapshotImage coregraphics.CGImageRef, presentationFrame corefoundation.CGRect) NSTextPreview {
	instance := getNSTextPreviewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSnapshotImage:presentationFrame:"), snapshotImage, presentationFrame)
	return NSTextPreviewFromID(rv)
}

// Creates a text preview using the specified image and rectangles that
// indicate the portions of text to highlight.
//
// snapshotImage: An image that contains the requested text from your view. Create the image
// using a transparent background and the current rendering attributes for
// your text.
//
// presentationFrame: A rectangle in the coordinate space of your text view. The system uses this
// rectangle to place your image precisely over your view’s actual text. Set
// its size to the size of your snapshot image, and set its origin to the
// point that allows the system to place your image directly over the text.
//
// candidateRects: An array of [NSValue] objects, each of which contains an [NSRect] in the
// coordinate space of your text view. Each rectangle contains a bounding
// rectangle for text that is part of the preview. When applying visual
// effects, the system adds highlights only to the text in the specified
// rectangles.
// //
// [NSRect]: https://developer.apple.com/documentation/Foundation/NSRect
// [NSValue]: https://developer.apple.com/documentation/Foundation/NSValue
//
// See: https://developer.apple.com/documentation/AppKit/NSTextPreview/init(snapshotImage:presentationFrame:candidateRects:)
func NewTextPreviewWithSnapshotImagePresentationFrameCandidateRects(snapshotImage coregraphics.CGImageRef, presentationFrame corefoundation.CGRect, candidateRects []foundation.NSValue) NSTextPreview {
	instance := getNSTextPreviewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSnapshotImage:presentationFrame:candidateRects:"), snapshotImage, presentationFrame, objectivec.IObjectSliceToNSArray(candidateRects))
	return NSTextPreviewFromID(rv)
}

// Creates a text preview using the specified image.
//
// snapshotImage: An image that contains the requested text from your view. Create the image
// using a transparent background and the current rendering attributes for
// your text.
//
// presentationFrame: A rectangle in your frame’s coordinate space. The system uses this
// rectangle to place your image precisely over your view’s actual text. Set
// its size to the size of your snapshot image, and set its origin to the
// point that allows the system to place your image directly over the text.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextPreview/init(snapshotImage:presentationFrame:)
func (t NSTextPreview) InitWithSnapshotImagePresentationFrame(snapshotImage coregraphics.CGImageRef, presentationFrame corefoundation.CGRect) NSTextPreview {
	rv := objc.Send[NSTextPreview](t.ID, objc.Sel("initWithSnapshotImage:presentationFrame:"), snapshotImage, presentationFrame)
	return rv
}
// Creates a text preview using the specified image and rectangles that
// indicate the portions of text to highlight.
//
// snapshotImage: An image that contains the requested text from your view. Create the image
// using a transparent background and the current rendering attributes for
// your text.
//
// presentationFrame: A rectangle in the coordinate space of your text view. The system uses this
// rectangle to place your image precisely over your view’s actual text. Set
// its size to the size of your snapshot image, and set its origin to the
// point that allows the system to place your image directly over the text.
//
// candidateRects: An array of [NSValue] objects, each of which contains an [NSRect] in the
// coordinate space of your text view. Each rectangle contains a bounding
// rectangle for text that is part of the preview. When applying visual
// effects, the system adds highlights only to the text in the specified
// rectangles.
// //
// [NSRect]: https://developer.apple.com/documentation/Foundation/NSRect
// [NSValue]: https://developer.apple.com/documentation/Foundation/NSValue
//
// See: https://developer.apple.com/documentation/AppKit/NSTextPreview/init(snapshotImage:presentationFrame:candidateRects:)
func (t NSTextPreview) InitWithSnapshotImagePresentationFrameCandidateRects(snapshotImage coregraphics.CGImageRef, presentationFrame corefoundation.CGRect, candidateRects []foundation.NSValue) NSTextPreview {
	rv := objc.Send[NSTextPreview](t.ID, objc.Sel("initWithSnapshotImage:presentationFrame:candidateRects:"), snapshotImage, presentationFrame, objectivec.IObjectSliceToNSArray(candidateRects))
	return rv
}

// The image that contains the requested text from your view.
//
// # Discussion
// 
// You specify this image at initialization time. The system uses it to
// implement any visual effects involving your view’s text. Create the image
// with your text on a transparent background.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextPreview/previewImage
func (t NSTextPreview) PreviewImage() coregraphics.CGImageRef {
	rv := objc.Send[coregraphics.CGImageRef](t.ID, objc.Sel("previewImage"))
	return coregraphics.CGImageRef(rv)
}
// The frame rectangle that places the preview image directly over the
// matching text.
//
// # Discussion
// 
// You specify this value at initialization time. The system uses it to
// position the preview image over the text in your view. Make sure the frame
// rectangle is in your view’s coordinate space.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextPreview/presentationFrame
func (t NSTextPreview) PresentationFrame() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](t.ID, objc.Sel("presentationFrame"))
	return corefoundation.CGRect(rv)
}
// Rectangles that define the specific portions of text to highlight.
//
// # Discussion
// 
// At initialization time, you set this property to an array of [NSValue]
// objects, each of which contains an [NSRect] in the coordinate space of the
// target view. Each rectangle contains a bounding rectangle for text that is
// part of the preview. When applying visual effects, the system adds
// highlights only to the text in the specified rectangles.
//
// [NSRect]: https://developer.apple.com/documentation/Foundation/NSRect
// [NSValue]: https://developer.apple.com/documentation/Foundation/NSValue
//
// See: https://developer.apple.com/documentation/AppKit/NSTextPreview/candidateRects
func (t NSTextPreview) CandidateRects() []foundation.NSValue {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("candidateRects"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSValue {
		return foundation.NSValueFromID(id)
	})
}

