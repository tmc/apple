// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NSTextAttachmentCell] class.
var (
	_NSTextAttachmentCellClass     NSTextAttachmentCellClass
	_NSTextAttachmentCellClassOnce sync.Once
)

func getNSTextAttachmentCellClass() NSTextAttachmentCellClass {
	_NSTextAttachmentCellClassOnce.Do(func() {
		_NSTextAttachmentCellClass = NSTextAttachmentCellClass{class: objc.GetClass("NSTextAttachmentCell")}
	})
	return _NSTextAttachmentCellClass
}

// GetNSTextAttachmentCellClass returns the class object for NSTextAttachmentCell.
func GetNSTextAttachmentCellClass() NSTextAttachmentCellClass {
	return getNSTextAttachmentCellClass()
}

type NSTextAttachmentCellClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTextAttachmentCellClass) Alloc() NSTextAttachmentCell {
	rv := objc.Send[NSTextAttachmentCell](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that implements the functionality of the text attachment cell
// protocol.
//
// # Overview
// 
// This specification describes only those methods whose implementations have
// features that are particular to this class. For a general discussion of the
// protocol’s methods, see [NSTextAttachmentCellProtocol].
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentCell-swift.class
type NSTextAttachmentCell struct {
	NSCell
}

// NSTextAttachmentCellFromID constructs a [NSTextAttachmentCell] from an objc.ID.
//
// An object that implements the functionality of the text attachment cell
// protocol.
func NSTextAttachmentCellFromID(id objc.ID) NSTextAttachmentCell {
	return NSTextAttachmentCell{NSCell: NSCellFromID(id)}
}
// NOTE: NSTextAttachmentCell adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSTextAttachmentCell] class.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentCell-swift.class
type INSTextAttachmentCell interface {
	INSCell
	NSTextAttachmentCellProtocol
}

// Init initializes the instance.
func (t NSTextAttachmentCell) Init() NSTextAttachmentCell {
	rv := objc.Send[NSTextAttachmentCell](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTextAttachmentCell) Autorelease() NSTextAttachmentCell {
	rv := objc.Send[NSTextAttachmentCell](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTextAttachmentCell creates a new NSTextAttachmentCell instance.
func NewNSTextAttachmentCell() NSTextAttachmentCell {
	class := getNSTextAttachmentCellClass()
	rv := objc.Send[NSTextAttachmentCell](objc.ID(class.class), objc.Sel("new"))
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
func NewTextAttachmentCellImageCell(image INSImage) NSTextAttachmentCell {
	instance := getNSTextAttachmentCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initImageCell:"), image)
	return NSTextAttachmentCellFromID(rv)
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
func NewTextAttachmentCellTextCell(string_ string) NSTextAttachmentCell {
	instance := getNSTextAttachmentCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initTextCell:"), objc.String(string_))
	return NSTextAttachmentCellFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSCell/init(coder:)
func NewTextAttachmentCellWithCoder(coder foundation.INSCoder) NSTextAttachmentCell {
	instance := getNSTextAttachmentCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSTextAttachmentCellFromID(rv)
}

			// Protocol methods for NSTextAttachmentCellProtocol
			
// Returns the text attachment object that owns the cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentCellProtocol/attachment
func (o NSTextAttachmentCell) Attachment() INSTextAttachment {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("attachment"))
	return NSTextAttachmentFromID(rv)
	}
// Draws the cell’s image in the specified rectangle of the currently
// focused view.
//
// cellFrame: The frame rectangle in which to draw.
//
// controlView: The view in which to draw.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentCellProtocol/draw(withFrame:in:)
func (o NSTextAttachmentCell) DrawWithFrameInView(cellFrame corefoundation.CGRect, controlView INSView) {
	objc.Send[struct{}](o.ID, objc.Sel("drawWithFrame:inView:"), cellFrame, controlView)
	}
// Draws the cell’s image at the specified index point in the view.
//
// cellFrame: The frame rectangle in which to draw.
//
// controlView: The view in which to draw.
//
// charIndex: The index of the attachment character within the text.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentCellProtocol/draw(withFrame:in:characterIndex:)
func (o NSTextAttachmentCell) DrawWithFrameInViewCharacterIndex(cellFrame corefoundation.CGRect, controlView INSView, charIndex uint) {
	objc.Send[struct{}](o.ID, objc.Sel("drawWithFrame:inView:characterIndex:"), cellFrame, controlView, charIndex)
	}
// Draws the cell’s image using the specified layout manager.
//
// cellFrame: The frame rectangle in which to draw.
//
// controlView: The view in which to draw.
//
// charIndex: The index of the attachment character within the text.
//
// layoutManager: The layout manager for the text.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentCellProtocol/draw(withFrame:in:characterIndex:layoutManager:)
func (o NSTextAttachmentCell) DrawWithFrameInViewCharacterIndexLayoutManager(cellFrame corefoundation.CGRect, controlView INSView, charIndex uint, layoutManager INSLayoutManager) {
	objc.Send[struct{}](o.ID, objc.Sel("drawWithFrame:inView:characterIndex:layoutManager:"), cellFrame, controlView, charIndex, layoutManager)
	}
// Draws the receiver’s image with optional highlighting.
//
// flag: A Boolean value that indicates whether to highlight the image. Add a
// highlight if the value is [true].
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// cellFrame: The frame rectangle in which to draw.
//
// controlView: The view in which to draw.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentCellProtocol/highlight(_:withFrame:in:)
func (o NSTextAttachmentCell) HighlightWithFrameInView(flag bool, cellFrame corefoundation.CGRect, controlView INSView) {
	objc.Send[struct{}](o.ID, objc.Sel("highlight:withFrame:inView:"), flag, cellFrame, controlView)
	}
// Returns the size of the attachment’s icon.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentCellProtocol/cellSize()
func (o NSTextAttachmentCell) CellSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](o.ID, objc.Sel("cellSize"))
	return rv
	}
// Returns the text position where you draw the attachment cell’s image,
// relative to the current point established in the glyph layout.
//
// # Discussion
// 
// The image should be drawn so its lower-left corner lies on this point.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentCellProtocol/cellBaselineOffset()
func (o NSTextAttachmentCell) CellBaselineOffset() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("cellBaselineOffset"))
	return rv
	}
// Returns the frame of the cell to draw at the specified position in a text
// container.
//
// textContainer: The text container that contains the glyph.
//
// lineFrag: The line fragment that contains the glyph.
//
// position: The position of the glyph in the text container.
//
// charIndex: The index of the character.
//
// # Discussion
// 
// The proposed line fragment is specified by `lineFrag`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentCellProtocol/cellFrame(for:proposedLineFragment:glyphPosition:characterIndex:)
func (o NSTextAttachmentCell) CellFrameForTextContainerProposedLineFragmentGlyphPositionCharacterIndex(textContainer INSTextContainer, lineFrag corefoundation.CGRect, position corefoundation.CGPoint, charIndex uint) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("cellFrameForTextContainer:proposedLineFragment:glyphPosition:characterIndex:"), textContainer, lineFrag, position, charIndex)
	return rv
	}
// Returns a Boolean value that indicates whether the attachment handles mouse
// events occurring over its image.
//
// # Discussion
// 
// The default implementation of this method returns [true]. The [NSView]
// containing the cell should invoke this method before sending a
// [TrackMouseInRectOfViewUntilMouseUp] message.
// 
// For an attachment in an attributed string, if the attachment cell returns
// [false], its attachment character should be selected rather than the cell
// being asked to track the mouse. This results in the attachment icon
// behaving as any regular glyph in text.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentCellProtocol/wantsToTrackMouse()
func (o NSTextAttachmentCell) WantsToTrackMouse() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("wantsToTrackMouse"))
	return rv
	}
// Allows an attachment to specify the events for which it tracks the mouse.
//
// # Discussion
// 
// `theEvent` is the event in question that occurred in `cellFrame` inside
// `controlView`. `charIndex` is the index of the attachment character within
// the text. If [WantsToTrackMouse] returns [true], this method allows the
// attachment to decide whether it wishes to do so for particular events.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentCellProtocol/wantsToTrackMouse(for:in:of:atCharacterIndex:)
func (o NSTextAttachmentCell) WantsToTrackMouseForEventInRectOfViewAtCharacterIndex(theEvent INSEvent, cellFrame corefoundation.CGRect, controlView INSView, charIndex uint) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("wantsToTrackMouseForEvent:inRect:ofView:atCharacterIndex:"), theEvent, cellFrame, controlView, charIndex)
	return rv
	}
// Handles a mouse-down event on the cell’s image, and optionally waits
// until a mouse-up event
//
// theEvent: The mouse-down event.
//
// cellFrame: The region of an [NSTextView] in which to track further mouse events.
//
// controlView: The view that received the event. Typically, this view is an [NSTextView]
// object and is focused.
//
// flag: A Boolean value that indicates whether to track the mouse until a mouse-up
// event occurs. If this parameter is [false], stop tracking when a
// mouse-dragged event occurs outside of `cellFrame`.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
//
// # Return Value
// 
// [true] if the cell successfully finished tracking the mouse, or [false] if
// tracking was unsuccessful.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The [NSTextAttachmentCell] implementation of this method calls upon
// `aTextView`’s delegate to handle the event. If `theEvent` is a mouse-up
// event for a double click, the text attachment cell calls the
// [TextViewDoubleClickedOnCellInRectAtIndex] method of its delegate and
// returns [true]. Otherwise, depending on whether the user clicks or drags
// the cell, it sends the delegate a [TextViewClickedOnCellInRectAtIndex]: or
// a [TextViewDraggedCellInRectEventAtIndex] message and returns [true]. The
// [NSTextAttachmentCell] implementation returns [false] only if `flag` is
// [false] and the mouse is dragged outside of `cellFrame`. The delegate
// methods are invoked only if the delegate responds.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentCellProtocol/trackMouse(with:in:of:untilMouseUp:)
func (o NSTextAttachmentCell) TrackMouseInRectOfViewUntilMouseUp(theEvent INSEvent, cellFrame corefoundation.CGRect, controlView INSView, flag bool) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("trackMouse:inRect:ofView:untilMouseUp:"), theEvent, cellFrame, controlView, flag)
	return rv
	}
// Handles a mouse-down event on the image at the specified character
// position.
//
// theEvent: The mouse-down event.
//
// cellFrame: The region of an [NSTextView] in which to track further mouse events.
//
// controlView: The view that received the event. Typically, this view is an [NSTextView]
// object and is focused.
//
// charIndex: The position in the text at which the attachment appears.
//
// flag: A Boolean value that indicates whether to track the mouse until a mouse-up
// event occurs. If this parameter is [false], stop tracking when a
// mouse-dragged event occurs outside of `cellFrame`.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
//
// # Return Value
// 
// [true] if the cell successfully finished tracking the mouse, or [false] if
// tracking was unsuccessful.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The [NSTextAttachmentCell] implementation of this method calls upon
// `aTextView`’s delegate to handle the event. If `theEvent` is a mouse-up
// event for a double click, the text attachment cell calls the
// [TextViewDoubleClickedOnCellInRectAtIndex] method of its delegate and
// returns [true]. Otherwise, depending on whether the user clicks or drags
// the cell, it sends the delegate a [TextViewClickedOnCellInRectAtIndex]: or
// a [TextViewDraggedCellInRectEventAtIndex] message and returns [true]. The
// [NSTextAttachmentCell] implementation returns [false] only if `flag` is
// [false] and the mouse is dragged outside of `cellFrame`. The delegate
// methods are invoked only if the delegate responds.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentCellProtocol/trackMouse(with:in:of:atCharacterIndex:untilMouseUp:)
func (o NSTextAttachmentCell) TrackMouseInRectOfViewAtCharacterIndexUntilMouseUp(theEvent INSEvent, cellFrame corefoundation.CGRect, controlView INSView, charIndex uint, flag bool) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("trackMouse:inRect:ofView:atCharacterIndex:untilMouseUp:"), theEvent, cellFrame, controlView, charIndex, flag)
	return rv
	}

func (o NSTextAttachmentCell) SetAttachment(value INSTextAttachment) {
	objc.Send[struct{}](o.ID, objc.Sel("setAttachment:"), value)
}

