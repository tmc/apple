// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// A set of methods that declares the interface for objects that draw text attachment icons and handle mouse events on their icons.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentCellProtocol
type NSTextAttachmentCellProtocol interface {
	objectivec.IObject

	// Returns the text attachment object that owns the cell.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentCellProtocol/attachment
	Attachment() INSTextAttachment

	// Draws the cell’s image in the specified rectangle of the currently focused view.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentCellProtocol/draw(withFrame:in:)
	DrawWithFrameInView(cellFrame corefoundation.CGRect, controlView INSView)

	// Draws the cell’s image at the specified index point in the view.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentCellProtocol/draw(withFrame:in:characterIndex:)
	DrawWithFrameInViewCharacterIndex(cellFrame corefoundation.CGRect, controlView INSView, charIndex uint)

	// Draws the cell’s image using the specified layout manager.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentCellProtocol/draw(withFrame:in:characterIndex:layoutManager:)
	DrawWithFrameInViewCharacterIndexLayoutManager(cellFrame corefoundation.CGRect, controlView INSView, charIndex uint, layoutManager INSLayoutManager)

	// Draws the receiver’s image with optional highlighting.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentCellProtocol/highlight(_:withFrame:in:)
	HighlightWithFrameInView(flag bool, cellFrame corefoundation.CGRect, controlView INSView)

	// Returns the size of the attachment’s icon.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentCellProtocol/cellSize()
	CellSize() corefoundation.CGSize

	// Returns the text position where you draw the attachment cell’s image, relative to the current point established in the glyph layout.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentCellProtocol/cellBaselineOffset()
	CellBaselineOffset() corefoundation.CGPoint

	// Returns the frame of the cell to draw at the specified position in a text container.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentCellProtocol/cellFrame(for:proposedLineFragment:glyphPosition:characterIndex:)
	CellFrameForTextContainerProposedLineFragmentGlyphPositionCharacterIndex(textContainer INSTextContainer, lineFrag corefoundation.CGRect, position corefoundation.CGPoint, charIndex uint) corefoundation.CGRect

	// Returns a Boolean value that indicates whether the attachment handles mouse events occurring over its image.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentCellProtocol/wantsToTrackMouse()
	WantsToTrackMouse() bool

	// Allows an attachment to specify the events for which it tracks the mouse.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentCellProtocol/wantsToTrackMouse(for:in:of:atCharacterIndex:)
	WantsToTrackMouseForEventInRectOfViewAtCharacterIndex(theEvent INSEvent, cellFrame corefoundation.CGRect, controlView INSView, charIndex uint) bool

	// Handles a mouse-down event on the cell’s image, and optionally waits until a mouse-up event
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentCellProtocol/trackMouse(with:in:of:untilMouseUp:)
	TrackMouseInRectOfViewUntilMouseUp(theEvent INSEvent, cellFrame corefoundation.CGRect, controlView INSView, flag bool) bool

	// Handles a mouse-down event on the image at the specified character position.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentCellProtocol/trackMouse(with:in:of:atCharacterIndex:untilMouseUp:)
	TrackMouseInRectOfViewAtCharacterIndexUntilMouseUp(theEvent INSEvent, cellFrame corefoundation.CGRect, controlView INSView, charIndex uint, flag bool) bool

	// Returns the text attachment object that owns the cell.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentCellProtocol/attachment
	SetAttachment(value INSTextAttachment)
}

// NSTextAttachmentCellProtocolObject wraps an existing Objective-C object that conforms to the NSTextAttachmentCellProtocol protocol.
type NSTextAttachmentCellProtocolObject struct {
	objectivec.Object
}
func (o NSTextAttachmentCellProtocolObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSTextAttachmentCellProtocolObjectFromID constructs a [NSTextAttachmentCellProtocolObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSTextAttachmentCellProtocolObjectFromID(id objc.ID) NSTextAttachmentCellProtocolObject {
	return NSTextAttachmentCellProtocolObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Returns the text attachment object that owns the cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentCellProtocol/attachment
func (o NSTextAttachmentCellProtocolObject) Attachment() INSTextAttachment {
	
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
func (o NSTextAttachmentCellProtocolObject) DrawWithFrameInView(cellFrame corefoundation.CGRect, controlView INSView) {
	
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
func (o NSTextAttachmentCellProtocolObject) DrawWithFrameInViewCharacterIndex(cellFrame corefoundation.CGRect, controlView INSView, charIndex uint) {
	
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
func (o NSTextAttachmentCellProtocolObject) DrawWithFrameInViewCharacterIndexLayoutManager(cellFrame corefoundation.CGRect, controlView INSView, charIndex uint, layoutManager INSLayoutManager) {
	
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
func (o NSTextAttachmentCellProtocolObject) HighlightWithFrameInView(flag bool, cellFrame corefoundation.CGRect, controlView INSView) {
	
	objc.Send[struct{}](o.ID, objc.Sel("highlight:withFrame:inView:"), flag, cellFrame, controlView)
	}
// Returns the size of the attachment’s icon.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextAttachmentCellProtocol/cellSize()
func (o NSTextAttachmentCellProtocolObject) CellSize() corefoundation.CGSize {
	
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
func (o NSTextAttachmentCellProtocolObject) CellBaselineOffset() corefoundation.CGPoint {
	
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
func (o NSTextAttachmentCellProtocolObject) CellFrameForTextContainerProposedLineFragmentGlyphPositionCharacterIndex(textContainer INSTextContainer, lineFrag corefoundation.CGRect, position corefoundation.CGPoint, charIndex uint) corefoundation.CGRect {
	
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
func (o NSTextAttachmentCellProtocolObject) WantsToTrackMouse() bool {
	
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
func (o NSTextAttachmentCellProtocolObject) WantsToTrackMouseForEventInRectOfViewAtCharacterIndex(theEvent INSEvent, cellFrame corefoundation.CGRect, controlView INSView, charIndex uint) bool {
	
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
func (o NSTextAttachmentCellProtocolObject) TrackMouseInRectOfViewUntilMouseUp(theEvent INSEvent, cellFrame corefoundation.CGRect, controlView INSView, flag bool) bool {
	
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
func (o NSTextAttachmentCellProtocolObject) TrackMouseInRectOfViewAtCharacterIndexUntilMouseUp(theEvent INSEvent, cellFrame corefoundation.CGRect, controlView INSView, charIndex uint, flag bool) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("trackMouse:inRect:ofView:atCharacterIndex:untilMouseUp:"), theEvent, cellFrame, controlView, charIndex, flag)
	return rv
	}

func (o NSTextAttachmentCellProtocolObject) SetAttachment(value INSTextAttachment) {
	objc.Send[struct{}](o.ID, objc.Sel("setAttachment:"), value)
}

