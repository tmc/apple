// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NSTextTable] class.
var (
	_NSTextTableClass     NSTextTableClass
	_NSTextTableClassOnce sync.Once
)

func getNSTextTableClass() NSTextTableClass {
	_NSTextTableClassOnce.Do(func() {
		_NSTextTableClass = NSTextTableClass{class: objc.GetClass("NSTextTable")}
	})
	return _NSTextTableClass
}

// GetNSTextTableClass returns the class object for NSTextTable.
func GetNSTextTableClass() NSTextTableClass {
	return getNSTextTableClass()
}

type NSTextTableClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTextTableClass) Alloc() NSTextTable {
	rv := objc.Send[NSTextTable](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// An object that represents a text table as a whole.
//
// # Overview
// 
// A text table is responsible for laying out and drawing the text table
// blocks it contains, and it maintains the basic parameters of the table.
//
// # Getting and setting number of columns
//
//   - [NSTextTable.NumberOfColumns]: The number of columns in the text table.
//   - [NSTextTable.SetNumberOfColumns]
//
// # Getting and setting layout algorithm
//
//   - [NSTextTable.LayoutAlgorithm]: The text table layout algorithm.
//   - [NSTextTable.SetLayoutAlgorithm]
//
// # Collapsing borders
//
//   - [NSTextTable.CollapsesBorders]: A Boolean value indicating whether the text table borders are collapsible.
//   - [NSTextTable.SetCollapsesBorders]
//
// # Hiding empty cells
//
//   - [NSTextTable.HidesEmptyCells]: A Boolean value indicating whether the text table hides empty cells.
//   - [NSTextTable.SetHidesEmptyCells]
//
// # Determining layout rectangles
//
//   - [NSTextTable.RectForBlockLayoutAtPointInRectTextContainerCharacterRange]: Returns the rectangle within which glyphs should be laid out for a text table block.
//   - [NSTextTable.BoundsRectForBlockContentRectInRectTextContainerCharacterRange]: Returns the rectangle the text table block actually occupies, including padding, borders, and margins.
//
// # Drawing the table
//
//   - [NSTextTable.DrawBackgroundForBlockWithFrameInViewCharacterRangeLayoutManager]: Draws any colors and other decorations for a text table block.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextTable
type NSTextTable struct {
	NSTextBlock
}

// NSTextTableFromID constructs a [NSTextTable] from an objc.ID.
//
// An object that represents a text table as a whole.
func NSTextTableFromID(id objc.ID) NSTextTable {
	return NSTextTable{NSTextBlock: NSTextBlockFromID(id)}
}
// NOTE: NSTextTable adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSTextTable] class.
//
// # Getting and setting number of columns
//
//   - [INSTextTable.NumberOfColumns]: The number of columns in the text table.
//   - [INSTextTable.SetNumberOfColumns]
//
// # Getting and setting layout algorithm
//
//   - [INSTextTable.LayoutAlgorithm]: The text table layout algorithm.
//   - [INSTextTable.SetLayoutAlgorithm]
//
// # Collapsing borders
//
//   - [INSTextTable.CollapsesBorders]: A Boolean value indicating whether the text table borders are collapsible.
//   - [INSTextTable.SetCollapsesBorders]
//
// # Hiding empty cells
//
//   - [INSTextTable.HidesEmptyCells]: A Boolean value indicating whether the text table hides empty cells.
//   - [INSTextTable.SetHidesEmptyCells]
//
// # Determining layout rectangles
//
//   - [INSTextTable.RectForBlockLayoutAtPointInRectTextContainerCharacterRange]: Returns the rectangle within which glyphs should be laid out for a text table block.
//   - [INSTextTable.BoundsRectForBlockContentRectInRectTextContainerCharacterRange]: Returns the rectangle the text table block actually occupies, including padding, borders, and margins.
//
// # Drawing the table
//
//   - [INSTextTable.DrawBackgroundForBlockWithFrameInViewCharacterRangeLayoutManager]: Draws any colors and other decorations for a text table block.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextTable
type INSTextTable interface {
	INSTextBlock

	// Topic: Getting and setting number of columns

	// The number of columns in the text table.
	NumberOfColumns() uint
	SetNumberOfColumns(value uint)

	// Topic: Getting and setting layout algorithm

	// The text table layout algorithm.
	LayoutAlgorithm() NSTextTableLayoutAlgorithm
	SetLayoutAlgorithm(value NSTextTableLayoutAlgorithm)

	// Topic: Collapsing borders

	// A Boolean value indicating whether the text table borders are collapsible.
	CollapsesBorders() bool
	SetCollapsesBorders(value bool)

	// Topic: Hiding empty cells

	// A Boolean value indicating whether the text table hides empty cells.
	HidesEmptyCells() bool
	SetHidesEmptyCells(value bool)

	// Topic: Determining layout rectangles

	// Returns the rectangle within which glyphs should be laid out for a text table block.
	RectForBlockLayoutAtPointInRectTextContainerCharacterRange(block INSTextTableBlock, startingPoint corefoundation.CGPoint, rect corefoundation.CGRect, textContainer INSTextContainer, charRange foundation.NSRange) corefoundation.CGRect
	// Returns the rectangle the text table block actually occupies, including padding, borders, and margins.
	BoundsRectForBlockContentRectInRectTextContainerCharacterRange(block INSTextTableBlock, contentRect corefoundation.CGRect, rect corefoundation.CGRect, textContainer INSTextContainer, charRange foundation.NSRange) corefoundation.CGRect

	// Topic: Drawing the table

	// Draws any colors and other decorations for a text table block.
	DrawBackgroundForBlockWithFrameInViewCharacterRangeLayoutManager(block INSTextTableBlock, frameRect corefoundation.CGRect, controlView INSView, charRange foundation.NSRange, layoutManager INSLayoutManager)

	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (t NSTextTable) Init() NSTextTable {
	rv := objc.Send[NSTextTable](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTextTable) Autorelease() NSTextTable {
	rv := objc.Send[NSTextTable](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTextTable creates a new NSTextTable instance.
func NewNSTextTable() NSTextTable {
	class := getNSTextTableClass()
	rv := objc.Send[NSTextTable](objc.ID(class.class), objc.Sel("new"))
	return rv
}











// Returns the rectangle within which glyphs should be laid out for a text
// table block.
//
// block: The text table block that wants to determine where to layout its glyphs.
//
// startingPoint: The location, in container coordinates, where layout begins.
//
// rect: The rectangle in which the block is constrained to lie. For top-level
// blocks, this is the container rectangle of `textContainer`; for nested
// blocks, this is the layout rectangle of the enclosing block.
//
// textContainer: The text container being used for the layout.
//
// charRange: The range of the characters whose glyphs are to be drawn.
//
// # Return Value
// 
// The rectangle within which glyphs should be laid out.
//
// # Discussion
// 
// This method is called by the text table block `block` to determine the
// rectangle within which glyphs should be laid out for the text table block.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextTable/rect(for:layoutAt:in:textContainer:characterRange:)
func (t NSTextTable) RectForBlockLayoutAtPointInRectTextContainerCharacterRange(block INSTextTableBlock, startingPoint corefoundation.CGPoint, rect corefoundation.CGRect, textContainer INSTextContainer, charRange foundation.NSRange) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](t.ID, objc.Sel("rectForBlock:layoutAtPoint:inRect:textContainer:characterRange:"), block, startingPoint, rect, textContainer, charRange)
	return corefoundation.CGRect(rv)
}

// Returns the rectangle the text table block actually occupies, including
// padding, borders, and margins.
//
// block: The text table block that wants to determine where to layout its glyphs.
//
// contentRect: The actual rectangle in which the text was laid out, as determined by
// [RectForLayoutAtPointInRectTextContainerCharacterRange].
//
// rect: The initial rectangle in `textContainer` proposed by the typesetter.
//
// textContainer: The text container being used for the layout.
//
// charRange: The range of the characters whose glyphs are to be drawn.
//
// # Return Value
// 
// The rectangle the text table block actually occupies, including padding,
// borders, and margins.
//
// # Discussion
// 
// This method is called by the text table block `block` after it is laid out
// to determine the rectangle the text table block actually occupies,
// including padding, borders, and margins.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextTable/boundsRect(for:contentRect:in:textContainer:characterRange:)
func (t NSTextTable) BoundsRectForBlockContentRectInRectTextContainerCharacterRange(block INSTextTableBlock, contentRect corefoundation.CGRect, rect corefoundation.CGRect, textContainer INSTextContainer, charRange foundation.NSRange) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](t.ID, objc.Sel("boundsRectForBlock:contentRect:inRect:textContainer:characterRange:"), block, contentRect, rect, textContainer, charRange)
	return corefoundation.CGRect(rv)
}

// Draws any colors and other decorations for a text table block.
//
// block: The text table block that wants to draw its background.
//
// frameRect: The area in which drawing occurs.
//
// controlView: The view controlling the drawing.
//
// charRange: The range of the characters whose glyphs are to be drawn.
//
// layoutManager: The layout manager controlling the typesetting.
//
// # Discussion
// 
// This methods is called by the text table block `block` to draw any colors
// and other decorations before the text is drawn.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextTable/drawBackground(for:withFrame:in:characterRange:layoutManager:)
func (t NSTextTable) DrawBackgroundForBlockWithFrameInViewCharacterRangeLayoutManager(block INSTextTableBlock, frameRect corefoundation.CGRect, controlView INSView, charRange foundation.NSRange, layoutManager INSLayoutManager) {
	objc.Send[objc.ID](t.ID, objc.Sel("drawBackgroundForBlock:withFrame:inView:characterRange:layoutManager:"), block, frameRect, controlView, charRange, layoutManager)
}
func (t NSTextTable) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](t.ID, objc.Sel("encodeWithCoder:"), coder)
}












// The number of columns in the text table.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextTable/numberOfColumns
func (t NSTextTable) NumberOfColumns() uint {
	rv := objc.Send[uint](t.ID, objc.Sel("numberOfColumns"))
	return rv
}
func (t NSTextTable) SetNumberOfColumns(value uint) {
	objc.Send[struct{}](t.ID, objc.Sel("setNumberOfColumns:"), value)
}



// The text table layout algorithm.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextTable/layoutAlgorithm-swift.property
func (t NSTextTable) LayoutAlgorithm() NSTextTableLayoutAlgorithm {
	rv := objc.Send[NSTextTableLayoutAlgorithm](t.ID, objc.Sel("layoutAlgorithm"))
	return NSTextTableLayoutAlgorithm(rv)
}
func (t NSTextTable) SetLayoutAlgorithm(value NSTextTableLayoutAlgorithm) {
	objc.Send[struct{}](t.ID, objc.Sel("setLayoutAlgorithm:"), value)
}



// A Boolean value indicating whether the text table borders are collapsible.
//
// # Discussion
// 
// The value of this property is [true] when the text table borders are
// collapsible.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextTable/collapsesBorders
func (t NSTextTable) CollapsesBorders() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("collapsesBorders"))
	return rv
}
func (t NSTextTable) SetCollapsesBorders(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setCollapsesBorders:"), value)
}



// A Boolean value indicating whether the text table hides empty cells.
//
// # Discussion
// 
// The value of this property is [true] when the text table hides empty cells.
// If empty cells are hidden, locations with empty cells allow the background
// of the enclosing block or text container to show through.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextTable/hidesEmptyCells
func (t NSTextTable) HidesEmptyCells() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("hidesEmptyCells"))
	return rv
}
func (t NSTextTable) SetHidesEmptyCells(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setHidesEmptyCells:"), value)
}




























