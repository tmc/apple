// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NSImageCell] class.
var (
	_NSImageCellClass     NSImageCellClass
	_NSImageCellClassOnce sync.Once
)

func getNSImageCellClass() NSImageCellClass {
	_NSImageCellClassOnce.Do(func() {
		_NSImageCellClass = NSImageCellClass{class: objc.GetClass("NSImageCell")}
	})
	return _NSImageCellClass
}

// GetNSImageCellClass returns the class object for NSImageCell.
func GetNSImageCellClass() NSImageCellClass {
	return getNSImageCellClass()
}

type NSImageCellClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSImageCellClass) Alloc() NSImageCell {
	rv := objc.Send[NSImageCell](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An [NSImageCell] object displays a single image (encapsulated in an
// [NSImage] object) in a frame. This class provides methods for choosing the
// frame and for aligning and scaling the image to fit the frame.
//
// # Overview
// 
// The object value of an [NSImageCell] object must be an [NSImage] object, so
// if you use the [NSImageCell.ObjectValue] method of [NSCell], be sure to supply an
// [NSImage] object as an argument. Because an [NSImage] object does not need
// to be converted for display, do not use the [NSCell] methods relating to
// formatters.
// 
// An [NSImageCell] object is usually associated with some kind of control
// object. For example, an [NSMatrix] or an [NSTableView].
// 
// # Designated Initializers
// 
// When subclassing [NSImageCell] you must implement all of the designated
// initializers. Those methods are: init, [NSImageCell.InitWithCoder], [NSImageCell.InitTextCell], and
// [NSImageCell.InitImageCell].
//
// # Aligning and Scaling the Image
//
//   - [NSImageCell.ImageAlignment]: The alignment of the receiver’s image relative to its frame.
//   - [NSImageCell.SetImageAlignment]
//   - [NSImageCell.ImageScaling]: The scaling mode used to fit the receiver’s image into the frame.
//   - [NSImageCell.SetImageScaling]
//
// # Choosing the Frame
//
//   - [NSImageCell.ImageFrameStyle]: The style of the frame that borders the image.
//   - [NSImageCell.SetImageFrameStyle]
//
// See: https://developer.apple.com/documentation/AppKit/NSImageCell
type NSImageCell struct {
	NSCell
}

// NSImageCellFromID constructs a [NSImageCell] from an objc.ID.
//
// An [NSImageCell] object displays a single image (encapsulated in an
// [NSImage] object) in a frame. This class provides methods for choosing the
// frame and for aligning and scaling the image to fit the frame.
func NSImageCellFromID(id objc.ID) NSImageCell {
	return NSImageCell{NSCell: NSCellFromID(id)}
}
// NOTE: NSImageCell adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSImageCell] class.
//
// # Aligning and Scaling the Image
//
//   - [INSImageCell.ImageAlignment]: The alignment of the receiver’s image relative to its frame.
//   - [INSImageCell.SetImageAlignment]
//   - [INSImageCell.ImageScaling]: The scaling mode used to fit the receiver’s image into the frame.
//   - [INSImageCell.SetImageScaling]
//
// # Choosing the Frame
//
//   - [INSImageCell.ImageFrameStyle]: The style of the frame that borders the image.
//   - [INSImageCell.SetImageFrameStyle]
//
// See: https://developer.apple.com/documentation/AppKit/NSImageCell
type INSImageCell interface {
	INSCell

	// Topic: Aligning and Scaling the Image

	// The alignment of the receiver’s image relative to its frame.
	ImageAlignment() NSImageAlignment
	SetImageAlignment(value NSImageAlignment)
	// The scaling mode used to fit the receiver’s image into the frame.
	ImageScaling() NSImageScaling
	SetImageScaling(value NSImageScaling)

	// Topic: Choosing the Frame

	// The style of the frame that borders the image.
	ImageFrameStyle() NSImageFrameStyle
	SetImageFrameStyle(value NSImageFrameStyle)
}

// Init initializes the instance.
func (i NSImageCell) Init() NSImageCell {
	rv := objc.Send[NSImageCell](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i NSImageCell) Autorelease() NSImageCell {
	rv := objc.Send[NSImageCell](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSImageCell creates a new NSImageCell instance.
func NewNSImageCell() NSImageCell {
	class := getNSImageCellClass()
	rv := objc.Send[NSImageCell](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns an [NSCell] object initialized with the specified image and set to
// have the cell’s default menu.
//
// image: The image to use for the cell. If this parameter is `nil`, no image is set.
//
// # Return Value
// 
// An initialized [NSCell] object, or `nil` if the cell could not be
// initialized.
//
// # Discussion
// 
// This is one of four designated initializers you must implement when
// subclassing. See [NSCell] for the complete list.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/init(imageCell:)
func NewImageCellImageCell(image INSImage) NSImageCell {
	instance := getNSImageCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initImageCell:"), image)
	return NSImageCellFromID(rv)
}

// Returns an NSCell object initialized with the specified string and set to
// have the cell’s default menu.
//
// string: The initial string to use for the cell.
//
// # Return Value
// 
// An initialized [NSCell] object, or `nil` if the cell could not be
// initialized.
//
// # Discussion
// 
// If no field editor (a shared [NSText] object) has been created for all
// [NSCell] objects, one is created.
// 
// This is one of four designated initializers you must implement when
// subclassing. See [NSCell] for the complete list.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/init(textCell:)
func NewImageCellTextCell(string_ string) NSImageCell {
	instance := getNSImageCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initTextCell:"), objc.String(string_))
	return NSImageCellFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSCell/init(coder:)
func NewImageCellWithCoder(coder foundation.INSCoder) NSImageCell {
	instance := getNSImageCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSImageCellFromID(rv)
}

// The alignment of the receiver’s image relative to its frame.
//
// # Discussion
// 
// For a list of possible values, see [NSImageAlignment]. The default value is
// [NSImageAlignCenter].
//
// [NSImageAlignment]: https://developer.apple.com/documentation/AppKit/NSImageAlignment
//
// See: https://developer.apple.com/documentation/AppKit/NSImageCell/imageAlignment
func (i NSImageCell) ImageAlignment() NSImageAlignment {
	rv := objc.Send[NSImageAlignment](i.ID, objc.Sel("imageAlignment"))
	return NSImageAlignment(rv)
}
func (i NSImageCell) SetImageAlignment(value NSImageAlignment) {
	objc.Send[struct{}](i.ID, objc.Sel("setImageAlignment:"), value)
}
// The scaling mode used to fit the receiver’s image into the frame.
//
// # Discussion
// 
// For a list of possible values, see [NSImageScaling]. The default value is
// [NSImageScaling.scaleProportionallyDown].
//
// [NSImageScaling.scaleProportionallyDown]: https://developer.apple.com/documentation/AppKit/NSImageScaling/scaleProportionallyDown
// [NSImageScaling]: https://developer.apple.com/documentation/AppKit/NSImageScaling
//
// See: https://developer.apple.com/documentation/AppKit/NSImageCell/imageScaling
func (i NSImageCell) ImageScaling() NSImageScaling {
	rv := objc.Send[NSImageScaling](i.ID, objc.Sel("imageScaling"))
	return NSImageScaling(rv)
}
func (i NSImageCell) SetImageScaling(value NSImageScaling) {
	objc.Send[struct{}](i.ID, objc.Sel("setImageScaling:"), value)
}
// The style of the frame that borders the image.
//
// # Discussion
// 
// For a list of frame styles, see [NSImageView.FrameStyle]. The default value
// is [NSImageFrameNone].
//
// [NSImageView.FrameStyle]: https://developer.apple.com/documentation/AppKit/NSImageView/FrameStyle
//
// See: https://developer.apple.com/documentation/AppKit/NSImageCell/imageFrameStyle
func (i NSImageCell) ImageFrameStyle() NSImageFrameStyle {
	rv := objc.Send[NSImageFrameStyle](i.ID, objc.Sel("imageFrameStyle"))
	return NSImageFrameStyle(rv)
}
func (i NSImageCell) SetImageFrameStyle(value NSImageFrameStyle) {
	objc.Send[struct{}](i.ID, objc.Sel("setImageFrameStyle:"), value)
}

