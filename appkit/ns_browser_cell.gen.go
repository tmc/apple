// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSBrowserCell] class.
var (
	_NSBrowserCellClass     NSBrowserCellClass
	_NSBrowserCellClassOnce sync.Once
)

func getNSBrowserCellClass() NSBrowserCellClass {
	_NSBrowserCellClassOnce.Do(func() {
		_NSBrowserCellClass = NSBrowserCellClass{class: objc.GetClass("NSBrowserCell")}
	})
	return _NSBrowserCellClass
}

// GetNSBrowserCellClass returns the class object for NSBrowserCell.
func GetNSBrowserCellClass() NSBrowserCellClass {
	return getNSBrowserCellClass()
}

type NSBrowserCellClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSBrowserCellClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSBrowserCellClass) Alloc() NSBrowserCell {
	rv := objc.Send[NSBrowserCell](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The user interface of a browser.
//
// # Overview
//
// The [NSBrowserCell] class is the subclass of [NSCell] used by default to
// display data in the columns of an [NSBrowser] object. (Each column contains
// an [NSMatrix] object filled with [NSBrowserCell] objects.)
//
// # Configuring Browser Cells
//
//   - [NSBrowserCell.AlternateImage]: The browser cell’s image for the highlighted state.
//   - [NSBrowserCell.SetAlternateImage]
//
// # Managing Browser Cell State
//
//   - [NSBrowserCell.Reset]: Unhighlights the receiver and unsets its state.
//   - [NSBrowserCell.Set]: Highlights the receiver and sets its state.
//   - [NSBrowserCell.Leaf]: A Boolean that indicates whether the browser cell is a leaf or a branch cell.
//   - [NSBrowserCell.SetLeaf]
//   - [NSBrowserCell.Loaded]: A Boolean that indicates whether the cell is ready to display.
//   - [NSBrowserCell.SetLoaded]
//   - [NSBrowserCell.HighlightColorInView]: Returns the highlight color that the receiver wants to display.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowserCell
type NSBrowserCell struct {
	NSCell
}

// NSBrowserCellFromID constructs a [NSBrowserCell] from an objc.ID.
//
// The user interface of a browser.
func NSBrowserCellFromID(id objc.ID) NSBrowserCell {
	return NSBrowserCell{NSCell: NSCellFromID(id)}
}

// NOTE: NSBrowserCell adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSBrowserCell] class.
//
// # Configuring Browser Cells
//
//   - [INSBrowserCell.AlternateImage]: The browser cell’s image for the highlighted state.
//   - [INSBrowserCell.SetAlternateImage]
//
// # Managing Browser Cell State
//
//   - [INSBrowserCell.Reset]: Unhighlights the receiver and unsets its state.
//   - [INSBrowserCell.Set]: Highlights the receiver and sets its state.
//   - [INSBrowserCell.Leaf]: A Boolean that indicates whether the browser cell is a leaf or a branch cell.
//   - [INSBrowserCell.SetLeaf]
//   - [INSBrowserCell.Loaded]: A Boolean that indicates whether the cell is ready to display.
//   - [INSBrowserCell.SetLoaded]
//   - [INSBrowserCell.HighlightColorInView]: Returns the highlight color that the receiver wants to display.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowserCell
type INSBrowserCell interface {
	INSCell

	// Topic: Configuring Browser Cells

	// The browser cell’s image for the highlighted state.
	AlternateImage() INSImage
	SetAlternateImage(value INSImage)

	// Topic: Managing Browser Cell State

	// Unhighlights the receiver and unsets its state.
	Reset()
	// Highlights the receiver and sets its state.
	Set()
	// A Boolean that indicates whether the browser cell is a leaf or a branch cell.
	Leaf() bool
	SetLeaf(value bool)
	// A Boolean that indicates whether the cell is ready to display.
	Loaded() bool
	SetLoaded(value bool)
	// Returns the highlight color that the receiver wants to display.
	HighlightColorInView(controlView INSView) INSColor
}

// Init initializes the instance.
func (b NSBrowserCell) Init() NSBrowserCell {
	rv := objc.Send[NSBrowserCell](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b NSBrowserCell) Autorelease() NSBrowserCell {
	rv := objc.Send[NSBrowserCell](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSBrowserCell creates a new NSBrowserCell instance.
func NewNSBrowserCell() NSBrowserCell {
	class := getNSBrowserCellClass()
	rv := objc.Send[NSBrowserCell](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AppKit/NSBrowserCell/init(imageCell:)
func NewBrowserCellImageCell(image INSImage) NSBrowserCell {
	instance := getNSBrowserCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initImageCell:"), image)
	return NSBrowserCellFromID(rv)
}

// See: https://developer.apple.com/documentation/AppKit/NSBrowserCell/init(textCell:)
func NewBrowserCellTextCell(string_ string) NSBrowserCell {
	instance := getNSBrowserCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initTextCell:"), objc.String(string_))
	return NSBrowserCellFromID(rv)
}

// See: https://developer.apple.com/documentation/AppKit/NSBrowserCell/init(coder:)
func NewBrowserCellWithCoder(coder foundation.INSCoder) NSBrowserCell {
	instance := getNSBrowserCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSBrowserCellFromID(rv)
}

// Unhighlights the receiver and unsets its state.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowserCell/reset()
func (b NSBrowserCell) Reset() {
	objc.Send[objc.ID](b.ID, objc.Sel("reset"))
}

// Highlights the receiver and sets its state.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowserCell/set()
func (b NSBrowserCell) Set() {
	objc.Send[objc.ID](b.ID, objc.Sel("set"))
}

// Returns the highlight color that the receiver wants to display.
//
// controlView: The view for which to return the highlight color.
//
// # Return Value
//
// The highlight color.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowserCell/highlightColor(in:)
func (b NSBrowserCell) HighlightColorInView(controlView INSView) INSColor {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("highlightColorInView:"), controlView)
	return NSColorFromID(rv)
}

// The browser cell’s image for the highlighted state.
//
// # Discussion
//
// The image is drawn vertically centered on the left edge of the browser
// cell.
//
// Note that the image is drawn at the given size of the image.
// [NSBrowserCell] does not set the size of the image, nor does it clip the
// drawing of the image. Make sure this image is the correct size for drawing
// in the browser cell.
//
// When the value of this property is `nil`, it removes the alternate image
// for the browser cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowserCell/alternateImage
func (b NSBrowserCell) AlternateImage() INSImage {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("alternateImage"))
	return NSImageFromID(objc.ID(rv))
}
func (b NSBrowserCell) SetAlternateImage(value INSImage) {
	objc.Send[struct{}](b.ID, objc.Sel("setAlternateImage:"), value)
}

// A Boolean that indicates whether the browser cell is a leaf or a branch
// cell.
//
// # Discussion
//
// When the value of this property is true, the browser cell is a leaf cell.
//
// A branch [NSBrowserCell] has an image near its right edge indicating that
// more, hierarchically related information is available; when the user
// selects the cell, the [NSBrowser] displays a new column of [NSBrowserCell]
// objects. A leaf [NSBrowserCell] has no image, indicating that the user has
// reached a terminal piece of information; it doesn’t point to additional
// information.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowserCell/isLeaf
func (b NSBrowserCell) Leaf() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("isLeaf"))
	return rv
}
func (b NSBrowserCell) SetLeaf(value bool) {
	objc.Send[struct{}](b.ID, objc.Sel("setLeaf:"), value)
}

// A Boolean that indicates whether the cell is ready to display.
//
// # Discussion
//
// When the value of this property is true, the browser cell’s state has
// been set and the cell is ready to display.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowserCell/isLoaded
func (b NSBrowserCell) Loaded() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("isLoaded"))
	return rv
}
func (b NSBrowserCell) SetLoaded(value bool) {
	objc.Send[struct{}](b.ID, objc.Sel("setLoaded:"), value)
}

// Returns the default image for branch cells in a browser.
//
// # Return Value
//
// The default image used for branch [NSBrowserCell] objects. The default
// image is a right-pointing triangle.
//
// # Discussion
//
// Override this method if you want a different image. To have a branch
// [NSBrowserCell] with no image (and no space reserved for an image),
// override this method to return `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowserCell/branchImage
func (_NSBrowserCellClass NSBrowserCellClass) BranchImage() NSImage {
	rv := objc.Send[objc.ID](objc.ID(_NSBrowserCellClass.class), objc.Sel("branchImage"))
	return NSImageFromID(objc.ID(rv))
}

// Returns the default image for branch browser cells that are highlighted.
//
// # Return Value
//
// The default image used for branch [NSBrowserCell] objects that are
// highlighted. This is a lighter version of the image returned by
// [BranchImage].
//
// # Discussion
//
// Override this method if you want a different image.
//
// See: https://developer.apple.com/documentation/AppKit/NSBrowserCell/highlightedBranchImage
func (_NSBrowserCellClass NSBrowserCellClass) HighlightedBranchImage() NSImage {
	rv := objc.Send[objc.ID](objc.ID(_NSBrowserCellClass.class), objc.Sel("highlightedBranchImage"))
	return NSImageFromID(objc.ID(rv))
}
