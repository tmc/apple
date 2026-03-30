// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSMenuItemCell] class.
var (
	_NSMenuItemCellClass     NSMenuItemCellClass
	_NSMenuItemCellClassOnce sync.Once
)

func getNSMenuItemCellClass() NSMenuItemCellClass {
	_NSMenuItemCellClassOnce.Do(func() {
		_NSMenuItemCellClass = NSMenuItemCellClass{class: objc.GetClass("NSMenuItemCell")}
	})
	return _NSMenuItemCellClass
}

// GetNSMenuItemCellClass returns the class object for NSMenuItemCell.
func GetNSMenuItemCellClass() NSMenuItemCellClass {
	return getNSMenuItemCellClass()
}

type NSMenuItemCellClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSMenuItemCellClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSMenuItemCellClass) Alloc() NSMenuItemCell {
	rv := objc.Send[NSMenuItemCell](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that handles the measurement and display of a single menu item in
// its encompassing frame.
//
// # Overview
//
// # Configuring Menu-Item Attributes
//
//   - [NSMenuItemCell.MenuItem]: The menu item object associated with the cell.
//   - [NSMenuItemCell.SetMenuItem]
//
// # Calculating the Size of a Menu Item
//
//   - [NSMenuItemCell.CalcSize]: Calculates the minimum required width and height of the receiver’s menu item.
//   - [NSMenuItemCell.NeedsSizing]: A Boolean value indicating whether the size of the menu needs to be calculated.
//   - [NSMenuItemCell.SetNeedsSizing]
//   - [NSMenuItemCell.ImageWidth]: The width of the image associated with the menu item.
//   - [NSMenuItemCell.TitleWidth]: The width of the menu item’s text, measured in points.
//   - [NSMenuItemCell.KeyEquivalentWidth]: The width of the menu item’s key equivalent string.
//   - [NSMenuItemCell.StateImageWidth]: The width of the image used to indicate the state of the menu item.
//
// # Getting the Menu Item’s Drawing Rectangle
//
//   - [NSMenuItemCell.KeyEquivalentRectForBounds]: Returns the rectangle into which the menu item’s key equivalent should be drawn.
//   - [NSMenuItemCell.StateImageRectForBounds]: Returns the rectangle into which the menu item’s state image should be drawn.
//
// # Drawing the Menu Item
//
//   - [NSMenuItemCell.DrawBorderAndBackgroundWithFrameInView]: Draws the borders and background associated with the receiver’s menu item (if any).
//   - [NSMenuItemCell.DrawKeyEquivalentWithFrameInView]: Draws the key equivalent associated with the menu item.
//   - [NSMenuItemCell.DrawSeparatorItemWithFrameInView]: Draws a menu item separator.
//   - [NSMenuItemCell.DrawStateImageWithFrameInView]: Draws the state image associated with the menu item.
//   - [NSMenuItemCell.NeedsDisplay]: A Boolean value indicating whether the menu item needs to be displayed.
//   - [NSMenuItemCell.SetNeedsDisplay]
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItemCell
type NSMenuItemCell struct {
	NSButtonCell
}

// NSMenuItemCellFromID constructs a [NSMenuItemCell] from an objc.ID.
//
// An object that handles the measurement and display of a single menu item in
// its encompassing frame.
func NSMenuItemCellFromID(id objc.ID) NSMenuItemCell {
	return NSMenuItemCell{NSButtonCell: NSButtonCellFromID(id)}
}

// NOTE: NSMenuItemCell adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSMenuItemCell] class.
//
// # Configuring Menu-Item Attributes
//
//   - [INSMenuItemCell.MenuItem]: The menu item object associated with the cell.
//   - [INSMenuItemCell.SetMenuItem]
//
// # Calculating the Size of a Menu Item
//
//   - [INSMenuItemCell.CalcSize]: Calculates the minimum required width and height of the receiver’s menu item.
//   - [INSMenuItemCell.NeedsSizing]: A Boolean value indicating whether the size of the menu needs to be calculated.
//   - [INSMenuItemCell.SetNeedsSizing]
//   - [INSMenuItemCell.ImageWidth]: The width of the image associated with the menu item.
//   - [INSMenuItemCell.TitleWidth]: The width of the menu item’s text, measured in points.
//   - [INSMenuItemCell.KeyEquivalentWidth]: The width of the menu item’s key equivalent string.
//   - [INSMenuItemCell.StateImageWidth]: The width of the image used to indicate the state of the menu item.
//
// # Getting the Menu Item’s Drawing Rectangle
//
//   - [INSMenuItemCell.KeyEquivalentRectForBounds]: Returns the rectangle into which the menu item’s key equivalent should be drawn.
//   - [INSMenuItemCell.StateImageRectForBounds]: Returns the rectangle into which the menu item’s state image should be drawn.
//
// # Drawing the Menu Item
//
//   - [INSMenuItemCell.DrawBorderAndBackgroundWithFrameInView]: Draws the borders and background associated with the receiver’s menu item (if any).
//   - [INSMenuItemCell.DrawKeyEquivalentWithFrameInView]: Draws the key equivalent associated with the menu item.
//   - [INSMenuItemCell.DrawSeparatorItemWithFrameInView]: Draws a menu item separator.
//   - [INSMenuItemCell.DrawStateImageWithFrameInView]: Draws the state image associated with the menu item.
//   - [INSMenuItemCell.NeedsDisplay]: A Boolean value indicating whether the menu item needs to be displayed.
//   - [INSMenuItemCell.SetNeedsDisplay]
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItemCell
type INSMenuItemCell interface {
	INSButtonCell

	// Topic: Configuring Menu-Item Attributes

	// The menu item object associated with the cell.
	MenuItem() INSMenuItem
	SetMenuItem(value INSMenuItem)

	// Topic: Calculating the Size of a Menu Item

	// Calculates the minimum required width and height of the receiver’s menu item.
	CalcSize()
	// A Boolean value indicating whether the size of the menu needs to be calculated.
	NeedsSizing() bool
	SetNeedsSizing(value bool)
	// The width of the image associated with the menu item.
	ImageWidth() float64
	// The width of the menu item’s text, measured in points.
	TitleWidth() float64
	// The width of the menu item’s key equivalent string.
	KeyEquivalentWidth() float64
	// The width of the image used to indicate the state of the menu item.
	StateImageWidth() float64

	// Topic: Getting the Menu Item’s Drawing Rectangle

	// Returns the rectangle into which the menu item’s key equivalent should be drawn.
	KeyEquivalentRectForBounds(cellFrame corefoundation.CGRect) corefoundation.CGRect
	// Returns the rectangle into which the menu item’s state image should be drawn.
	StateImageRectForBounds(cellFrame corefoundation.CGRect) corefoundation.CGRect

	// Topic: Drawing the Menu Item

	// Draws the borders and background associated with the receiver’s menu item (if any).
	DrawBorderAndBackgroundWithFrameInView(cellFrame corefoundation.CGRect, controlView INSView)
	// Draws the key equivalent associated with the menu item.
	DrawKeyEquivalentWithFrameInView(cellFrame corefoundation.CGRect, controlView INSView)
	// Draws a menu item separator.
	DrawSeparatorItemWithFrameInView(cellFrame corefoundation.CGRect, controlView INSView)
	// Draws the state image associated with the menu item.
	DrawStateImageWithFrameInView(cellFrame corefoundation.CGRect, controlView INSView)
	// A Boolean value indicating whether the menu item needs to be displayed.
	NeedsDisplay() bool
	SetNeedsDisplay(value bool)
}

// Init initializes the instance.
func (m NSMenuItemCell) Init() NSMenuItemCell {
	rv := objc.Send[NSMenuItemCell](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m NSMenuItemCell) Autorelease() NSMenuItemCell {
	rv := objc.Send[NSMenuItemCell](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSMenuItemCell creates a new NSMenuItemCell instance.
func NewNSMenuItemCell() NSMenuItemCell {
	class := getNSMenuItemCellClass()
	rv := objc.Send[NSMenuItemCell](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AppKit/NSButtonCell/init(imageCell:)
func NewMenuItemCellImageCell(image INSImage) NSMenuItemCell {
	instance := getNSMenuItemCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initImageCell:"), image)
	return NSMenuItemCellFromID(rv)
}

// See: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/init(textCell:)
func NewMenuItemCellTextCell(string_ string) NSMenuItemCell {
	instance := getNSMenuItemCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initTextCell:"), objc.String(string_))
	return NSMenuItemCellFromID(rv)
}

// See: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/init(coder:)
func NewMenuItemCellWithCoder(coder foundation.INSCoder) NSMenuItemCell {
	instance := getNSMenuItemCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSMenuItemCellFromID(rv)
}

// Calculates the minimum required width and height of the receiver’s menu
// item.
//
// # Discussion
//
// The calculated values are cached for future use. This method also
// calculates the sizes of individual components of the cell’s menu item and
// caches those values.
//
// This method is invoked automatically when necessary. You should not need to
// invoke it directly.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/calcSize()
func (m NSMenuItemCell) CalcSize() {
	objc.Send[objc.ID](m.ID, objc.Sel("calcSize"))
}

// Returns the rectangle into which the menu item’s key equivalent should be
// drawn.
//
// cellFrame: A rectangle that defines the bounds of the receiver.
//
// # Return Value
//
// The returned rectangle is based on `cellFrame` but encompasses only the
// area to be occupied by the key equivalent.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/keyEquivalentRect(forBounds:)
func (m NSMenuItemCell) KeyEquivalentRectForBounds(cellFrame corefoundation.CGRect) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](m.ID, objc.Sel("keyEquivalentRectForBounds:"), cellFrame)
	return corefoundation.CGRect(rv)
}

// Returns the rectangle into which the menu item’s state image should be
// drawn.
//
// cellFrame: A rectangle that defines the bounds of the receiver.
//
// # Return Value
//
// The returned rectangle is based on `cellFrame` but encompasses only the
// area to be occupied by the menu item’s state image.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/stateImageRect(forBounds:)
func (m NSMenuItemCell) StateImageRectForBounds(cellFrame corefoundation.CGRect) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](m.ID, objc.Sel("stateImageRectForBounds:"), cellFrame)
	return corefoundation.CGRect(rv)
}

// Draws the borders and background associated with the receiver’s menu item
// (if any).
//
// cellFrame: A rectangle defining the receiver’s frame area.
//
// controlView: The view object that contains this cell (usually an [NSControl] object).
//
// # Discussion
//
// This method invokes the [NSCell] method [ImageRectForBounds], passing it
// `cellFrame`, to calculate the rectangle in which to draw the image. The
// cell invokes this method before invoking the methods to draw the other menu
// item components.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/drawBorderAndBackground(withFrame:in:)
func (m NSMenuItemCell) DrawBorderAndBackgroundWithFrameInView(cellFrame corefoundation.CGRect, controlView INSView) {
	objc.Send[objc.ID](m.ID, objc.Sel("drawBorderAndBackgroundWithFrame:inView:"), cellFrame, controlView)
}

// Draws the key equivalent associated with the menu item.
//
// cellFrame: A rectangle defining the receiver’s frame area.
//
// controlView: The view object that contains this cell (usually an [NSControl] object).
//
// # Discussion
//
// This method invokes [KeyEquivalentRectForBounds], passing it `cellFrame`,
// to calculate the rectangle in which to draw the key equivalent. This method
// is invoked by the cell’s “ method. You should not need to invoke it
// directly. Subclasses may override this method to control the drawing of the
// key equivalent.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/drawKeyEquivalent(withFrame:in:)
func (m NSMenuItemCell) DrawKeyEquivalentWithFrameInView(cellFrame corefoundation.CGRect, controlView INSView) {
	objc.Send[objc.ID](m.ID, objc.Sel("drawKeyEquivalentWithFrame:inView:"), cellFrame, controlView)
}

// Draws a menu item separator.
//
// cellFrame: A rectangle defining the receiver’s frame area.
//
// controlView: The view object that contains this cell (usually an [NSControl] object).
//
// # Discussion
//
// This method uses the `cellFrame` parameter to calculate the rectangle in
// which to draw the menu item separator. This method uses the `controlView`
// to determine whether the separator item should be drawn normally or
// flipped.
//
// You should not need to invoke this method directly. Subclasses may override
// this method to control the drawing of the separator.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/drawSeparatorItem(withFrame:in:)
func (m NSMenuItemCell) DrawSeparatorItemWithFrameInView(cellFrame corefoundation.CGRect, controlView INSView) {
	objc.Send[objc.ID](m.ID, objc.Sel("drawSeparatorItemWithFrame:inView:"), cellFrame, controlView)
}

// Draws the state image associated with the menu item.
//
// cellFrame: A rectangle defining the receiver’s frame area.
//
// controlView: The view object that contains this cell (usually an [NSControl] object).
//
// # Discussion
//
// This method invokes [StateImageRectForBounds], passing it `cellFrame`, to
// calculate the rectangle in which to draw the state image. This method is
// invoked by the cell’s “ method. You should not need to invoke it
// directly. Subclasses may override this method to control the drawing of the
// state image.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/drawStateImage(withFrame:in:)
func (m NSMenuItemCell) DrawStateImageWithFrameInView(cellFrame corefoundation.CGRect, controlView INSView) {
	objc.Send[objc.ID](m.ID, objc.Sel("drawStateImageWithFrame:inView:"), cellFrame, controlView)
}

// The menu item object associated with the cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/menuItem
func (m NSMenuItemCell) MenuItem() INSMenuItem {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("menuItem"))
	return NSMenuItemFromID(objc.ID(rv))
}
func (m NSMenuItemCell) SetMenuItem(value INSMenuItem) {
	objc.Send[struct{}](m.ID, objc.Sel("setMenuItem:"), value)
}

// A Boolean value indicating whether the size of the menu needs to be
// calculated.
//
// # Discussion
//
// When the value of this property is true, the next attempt to obtain size
// information about the menu cause the [CalcSize] method to be called. When
// the value of the property is false, the size information is obtained from
// the currently cached values.
//
// Subclasses that drastically change the way a menu item is drawn can change
// the value of this property to update the menu item information. Other parts
// of your application should not need to change this property directly. The
// cell checks this value of this property as necessary when the content of
// its menu item changes.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/needsSizing
func (m NSMenuItemCell) NeedsSizing() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("needsSizing"))
	return rv
}
func (m NSMenuItemCell) SetNeedsSizing(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setNeedsSizing:"), value)
}

// The width of the image associated with the menu item.
//
// # Discussion
//
// The width of the image is measured in points. You can associate an image
// with a menu item using the setImage: method of [NSMenuItem].
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/imageWidth
func (m NSMenuItemCell) ImageWidth() float64 {
	rv := objc.Send[float64](m.ID, objc.Sel("imageWidth"))
	return rv
}

// The width of the menu item’s text, measured in points.
//
// # Discussion
//
// To set the menu item’s text, use the setTitle: method of [NSMenuItem].
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/titleWidth
func (m NSMenuItemCell) TitleWidth() float64 {
	rv := objc.Send[float64](m.ID, objc.Sel("titleWidth"))
	return rv
}

// The width of the menu item’s key equivalent string.
//
// # Discussion
//
// To set the menu item’s key equivalent, use the [KeyEquivalent] property
// of [NSMenuItem].
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/keyEquivalentWidth
func (m NSMenuItemCell) KeyEquivalentWidth() float64 {
	rv := objc.Send[float64](m.ID, objc.Sel("keyEquivalentWidth"))
	return rv
}

// The width of the image used to indicate the state of the menu item.
//
// # Discussion
//
// If the menu item has multiple images associated with it (to indicate any of
// the available states: on, off, or mixed), this property contains the width
// of the largest image. You can set the state images for a menu item using
// the setOnStateImage:, setOffStateImage:, and setMixedStateImage: methods of
// [NSMenuItem].
//
// To change the state of the cell’s menu item, use the setState: method of
// [NSMenuItem].
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/stateImageWidth
func (m NSMenuItemCell) StateImageWidth() float64 {
	rv := objc.Send[float64](m.ID, objc.Sel("stateImageWidth"))
	return rv
}

// A Boolean value indicating whether the menu item needs to be displayed.
//
// # Discussion
//
// Set this property to true when you want the menu item to be drawn.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/needsDisplay
func (m NSMenuItemCell) NeedsDisplay() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("needsDisplay"))
	return rv
}
func (m NSMenuItemCell) SetNeedsDisplay(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setNeedsDisplay:"), value)
}
