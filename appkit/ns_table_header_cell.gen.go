// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NSTableHeaderCell] class.
var (
	_NSTableHeaderCellClass     NSTableHeaderCellClass
	_NSTableHeaderCellClassOnce sync.Once
)

func getNSTableHeaderCellClass() NSTableHeaderCellClass {
	_NSTableHeaderCellClassOnce.Do(func() {
		_NSTableHeaderCellClass = NSTableHeaderCellClass{class: objc.GetClass("NSTableHeaderCell")}
	})
	return _NSTableHeaderCellClass
}

// GetNSTableHeaderCellClass returns the class object for NSTableHeaderCell.
func GetNSTableHeaderCellClass() NSTableHeaderCellClass {
	return getNSTableHeaderCellClass()
}

type NSTableHeaderCellClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSTableHeaderCellClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTableHeaderCellClass) Alloc() NSTableHeaderCell {
	rv := objc.Send[NSTableHeaderCell](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that a table header view uses to draw the content of the column
// headers.
//
// # Overview
// 
// Subclasses of the [NSTableHeaderCell] class can override the
// [DrawInteriorWithFrameInView], [EditWithFrameInViewEditorDelegateEvent],
// and [HighlightWithFrameInView] methods to change the way headers appear.
// This specific subclass is responsible for drawing the sort indicators. See
// the [NSCell] class specification for information on overriding these
// methods.
// 
// See the [NSTableView] and [NSTableHeaderCell] for more information.
//
// # Drawing Sorting Indicators
//
//   - [NSTableHeaderCell.DrawSortIndicatorWithFrameInViewAscendingPriority]: Draws a sorting indicator given a cell frame contained inside a view.
//   - [NSTableHeaderCell.SortIndicatorRectForBounds]: Returns the location to display the sorting indicator given `theRect`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableHeaderCell
type NSTableHeaderCell struct {
	NSTextFieldCell
}

// NSTableHeaderCellFromID constructs a [NSTableHeaderCell] from an objc.ID.
//
// An object that a table header view uses to draw the content of the column
// headers.
func NSTableHeaderCellFromID(id objc.ID) NSTableHeaderCell {
	return NSTableHeaderCell{NSTextFieldCell: NSTextFieldCellFromID(id)}
}
// NOTE: NSTableHeaderCell adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSTableHeaderCell] class.
//
// # Drawing Sorting Indicators
//
//   - [INSTableHeaderCell.DrawSortIndicatorWithFrameInViewAscendingPriority]: Draws a sorting indicator given a cell frame contained inside a view.
//   - [INSTableHeaderCell.SortIndicatorRectForBounds]: Returns the location to display the sorting indicator given `theRect`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableHeaderCell
type INSTableHeaderCell interface {
	INSTextFieldCell

	// Topic: Drawing Sorting Indicators

	// Draws a sorting indicator given a cell frame contained inside a view.
	DrawSortIndicatorWithFrameInViewAscendingPriority(cellFrame corefoundation.CGRect, controlView INSView, ascending bool, priority int)
	// Returns the location to display the sorting indicator given `theRect`.
	SortIndicatorRectForBounds(rect corefoundation.CGRect) corefoundation.CGRect
}

// Init initializes the instance.
func (t NSTableHeaderCell) Init() NSTableHeaderCell {
	rv := objc.Send[NSTableHeaderCell](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTableHeaderCell) Autorelease() NSTableHeaderCell {
	rv := objc.Send[NSTableHeaderCell](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTableHeaderCell creates a new NSTableHeaderCell instance.
func NewNSTableHeaderCell() NSTableHeaderCell {
	class := getNSTableHeaderCellClass()
	rv := objc.Send[NSTableHeaderCell](objc.ID(class.class), objc.Sel("new"))
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
func NewTableHeaderCellImageCell(image INSImage) NSTableHeaderCell {
	instance := getNSTableHeaderCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initImageCell:"), image)
	return NSTableHeaderCellFromID(rv)
}

// Initializes a text field cell that displays the specified string.
//
// string: The string that the text field cell displays.
//
// # Return Value
// 
// A text field cell that displays a string.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/init(textCell:)
func NewTableHeaderCellTextCell(string_ string) NSTableHeaderCell {
	instance := getNSTableHeaderCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initTextCell:"), objc.String(string_))
	return NSTableHeaderCellFromID(rv)
}

// Initializes a text field cell from data in the provided unarchiver.
//
// coder: An unarchiver object.
//
// # Return Value
// 
// A text field cell that displays a string.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/init(coder:)
func NewTableHeaderCellWithCoder(coder foundation.INSCoder) NSTableHeaderCell {
	instance := getNSTableHeaderCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSTableHeaderCellFromID(rv)
}

// Draws a sorting indicator given a cell frame contained inside a view.
//
// cellFrame: The cell frame.
//
// controlView: The control view.
//
// ascending: If YES the sort indicator is drawn as ascending; otherwise it is drawn as
// descending.
//
// priority: If `priority` is 0, this is the primary sort indicator.
//
// # Discussion
// 
// Override this method to customize the sorting user interface.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableHeaderCell/drawSortIndicator(withFrame:in:ascending:priority:)
func (t NSTableHeaderCell) DrawSortIndicatorWithFrameInViewAscendingPriority(cellFrame corefoundation.CGRect, controlView INSView, ascending bool, priority int) {
	objc.Send[objc.ID](t.ID, objc.Sel("drawSortIndicatorWithFrame:inView:ascending:priority:"), cellFrame, controlView, ascending, priority)
}
// Returns the location to display the sorting indicator given `theRect`.
//
// rect: A cell rectangle.
//
// # Return Value
// 
// The rectangle within `theRect` that should contain the sorting indicator.
//
// See: https://developer.apple.com/documentation/AppKit/NSTableHeaderCell/sortIndicatorRect(forBounds:)
func (t NSTableHeaderCell) SortIndicatorRectForBounds(rect corefoundation.CGRect) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](t.ID, objc.Sel("sortIndicatorRectForBounds:"), rect)
	return corefoundation.CGRect(rv)
}

