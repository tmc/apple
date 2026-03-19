// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NSScrubberImageItemView] class.
var (
	_NSScrubberImageItemViewClass     NSScrubberImageItemViewClass
	_NSScrubberImageItemViewClassOnce sync.Once
)

func getNSScrubberImageItemViewClass() NSScrubberImageItemViewClass {
	_NSScrubberImageItemViewClassOnce.Do(func() {
		_NSScrubberImageItemViewClass = NSScrubberImageItemViewClass{class: objc.GetClass("NSScrubberImageItemView")}
	})
	return _NSScrubberImageItemViewClass
}

// GetNSScrubberImageItemViewClass returns the class object for NSScrubberImageItemView.
func GetNSScrubberImageItemViewClass() NSScrubberImageItemViewClass {
	return getNSScrubberImageItemViewClass()
}

type NSScrubberImageItemViewClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSScrubberImageItemViewClass) Alloc() NSScrubberImageItemView {
	rv := objc.Send[NSScrubberImageItemView](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A concrete view subclass for displaying images in a scrubber items.
//
// # Overview
// 
// Provide the image you want to display in the scrubber item to the [Image]
// property. If you want finer control over the appearance of the image, you
// can access the underlying image view using the [ImageView] property.
// 
// The image is scaled proportionally to fit the view’s frame. Use the
// [NSScrubberImageItemView.ImageAlignment] property to determine how the scaled image is cropped
// within that frame.
//
// # Providing image content
//
//   - [NSScrubberImageItemView.Image]: The image displayed by the scrubber item.
//   - [NSScrubberImageItemView.SetImage]
//   - [NSScrubberImageItemView.ImageView]: The image view that the scrubber item uses to display its image.
//
// # Configuring the appearance
//
//   - [NSScrubberImageItemView.ImageAlignment]: The alignment of the image within the scrubber item.
//   - [NSScrubberImageItemView.SetImageAlignment]
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberImageItemView
type NSScrubberImageItemView struct {
	NSScrubberItemView
}

// NSScrubberImageItemViewFromID constructs a [NSScrubberImageItemView] from an objc.ID.
//
// A concrete view subclass for displaying images in a scrubber items.
func NSScrubberImageItemViewFromID(id objc.ID) NSScrubberImageItemView {
	return NSScrubberImageItemView{NSScrubberItemView: NSScrubberItemViewFromID(id)}
}
// NOTE: NSScrubberImageItemView adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSScrubberImageItemView] class.
//
// # Providing image content
//
//   - [INSScrubberImageItemView.Image]: The image displayed by the scrubber item.
//   - [INSScrubberImageItemView.SetImage]
//   - [INSScrubberImageItemView.ImageView]: The image view that the scrubber item uses to display its image.
//
// # Configuring the appearance
//
//   - [INSScrubberImageItemView.ImageAlignment]: The alignment of the image within the scrubber item.
//   - [INSScrubberImageItemView.SetImageAlignment]
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberImageItemView
type INSScrubberImageItemView interface {
	INSScrubberItemView

	// Topic: Providing image content

	// The image displayed by the scrubber item.
	Image() INSImage
	SetImage(value INSImage)
	// The image view that the scrubber item uses to display its image.
	ImageView() INSImageView

	// Topic: Configuring the appearance

	// The alignment of the image within the scrubber item.
	ImageAlignment() NSImageAlignment
	SetImageAlignment(value NSImageAlignment)
}

// Init initializes the instance.
func (s NSScrubberImageItemView) Init() NSScrubberImageItemView {
	rv := objc.Send[NSScrubberImageItemView](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSScrubberImageItemView) Autorelease() NSScrubberImageItemView {
	rv := objc.Send[NSScrubberImageItemView](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSScrubberImageItemView creates a new NSScrubberImageItemView instance.
func NewNSScrubberImageItemView() NSScrubberImageItemView {
	class := getNSScrubberImageItemViewClass()
	rv := objc.Send[NSScrubberImageItemView](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a view using from data in the specified coder object.
//
// coder: The coder object that contains the view’s configuration details.
//
// # Return Value
// 
// An initialized view or `nil` if AppKit couldn’t create the object.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/init(coder:)
func NewScrubberImageItemViewWithCoder(coder foundation.INSCoder) NSScrubberImageItemView {
	instance := getNSScrubberImageItemViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSScrubberImageItemViewFromID(rv)
}

// Initializes and returns a newly allocated [NSView] object with a specified
// frame rectangle.
//
// frameRect: The frame rectangle for the created view object.
//
// # Return Value
// 
// An initialized view or `nil` if AppKit couldn’t create the object.
//
// # Discussion
// 
// Insert the view into your window’s view hieararchy before you can do
// anything with it. This method is the designated initializer for the
// [NSView] class.
//
// See: https://developer.apple.com/documentation/AppKit/NSView/init(frame:)
func NewScrubberImageItemViewWithFrame(frameRect corefoundation.CGRect) NSScrubberImageItemView {
	instance := getNSScrubberImageItemViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSScrubberImageItemViewFromID(rv)
}

// The image displayed by the scrubber item.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberImageItemView/image
func (s NSScrubberImageItemView) Image() INSImage {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("image"))
	return NSImageFromID(objc.ID(rv))
}
func (s NSScrubberImageItemView) SetImage(value INSImage) {
	objc.Send[struct{}](s.ID, objc.Sel("setImage:"), value)
}
// The image view that the scrubber item uses to display its image.
//
// # Discussion
// 
// Use this property to access and configure the underlying image view used to
// display the image in the [Image] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberImageItemView/imageView
func (s NSScrubberImageItemView) ImageView() INSImageView {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("imageView"))
	return NSImageViewFromID(objc.ID(rv))
}
// The alignment of the image within the scrubber item.
//
// # Discussion
// 
// The image is scaled proportionally to fit the frame of the item. This
// property determines how the image is cropped within that frame.
// 
// For possible values, see [NSImageAlignment].
//
// [NSImageAlignment]: https://developer.apple.com/documentation/AppKit/NSImageAlignment
//
// See: https://developer.apple.com/documentation/AppKit/NSScrubberImageItemView/imageAlignment
func (s NSScrubberImageItemView) ImageAlignment() NSImageAlignment {
	rv := objc.Send[NSImageAlignment](s.ID, objc.Sel("imageAlignment"))
	return NSImageAlignment(rv)
}
func (s NSScrubberImageItemView) SetImageAlignment(value NSImageAlignment) {
	objc.Send[struct{}](s.ID, objc.Sel("setImageAlignment:"), value)
}

