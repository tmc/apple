// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// A set of methods that text views need to implement to interact properly with the text input management system.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextInputClient
type NSTextInputClient interface {
	objectivec.IObject

	// Returns a Boolean value indicating whether the receiver has marked text.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputClient/hasMarkedText()
	HasMarkedText() bool

	// Returns the range of the marked text.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputClient/markedRange()
	MarkedRange() foundation.NSRange

	// Returns the range of selected text.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputClient/selectedRange()
	SelectedRange() foundation.NSRange

	// Unmarks the marked text.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputClient/unmarkText()
	UnmarkText()

	// Returns an array of attribute names recognized by the receiver.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputClient/validAttributesForMarkedText()
	ValidAttributesForMarkedText() []string

	// Returns an attributed string derived from the given range in the receiver’s text storage.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputClient/attributedSubstring(forProposedRange:actualRange:)
	AttributedSubstringForProposedRangeActualRange(range_ foundation.NSRange, actualRange foundation.NSRange) foundation.NSAttributedString

	// Returns the index of the character whose bounding rectangle includes the given point.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputClient/characterIndex(for:)
	CharacterIndexForPoint(point corefoundation.CGPoint) uint

	// Returns the first logical boundary rectangle for characters in the given range.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputClient/firstRect(forCharacterRange:actualRange:)
	FirstRectForCharacterRangeActualRange(range_ foundation.NSRange, actualRange foundation.NSRange) corefoundation.CGRect

	// DocumentVisibleRect protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputClient/documentVisibleRect
	DocumentVisibleRect() corefoundation.CGRect

	// UnionRectInVisibleSelectedRange protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputClient/unionRectInVisibleSelectedRange
	UnionRectInVisibleSelectedRange() corefoundation.CGRect

	// Invokes the action specified by the given selector.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputClient/doCommand(by:)
	DoCommandBySelector(selector objc.SEL)

	// A Boolean value that indicates whether the document supports adaptive images in the input.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSTextInputClient/supportsAdaptiveImageGlyph
	SupportsAdaptiveImageGlyph() bool
}

// NSTextInputClientObject wraps an existing Objective-C object that conforms to the NSTextInputClient protocol.
type NSTextInputClientObject struct {
	objectivec.Object
}
func (o NSTextInputClientObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSTextInputClientObjectFromID constructs a [NSTextInputClientObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSTextInputClientObjectFromID(id objc.ID) NSTextInputClientObject {
	return NSTextInputClientObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Returns a Boolean value indicating whether the receiver has marked text.
//
// # Return Value
// 
// [true] if the receiver has marked text; otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The text view itself may call this method to determine whether there
// currently is marked text. [NSTextView], for example, disables the Edit >
// Copy menu item when this method returns [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextInputClient/hasMarkedText()
func (o NSTextInputClientObject) HasMarkedText() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("hasMarkedText"))
	return rv
	}
// Returns the range of the marked text.
//
// # Return Value
// 
// The range of marked text or `{NSNotFound, 0}` if there is no marked range.
//
// # Discussion
// 
// The returned range measures from the start of the receiver’s text
// storage. The return value’s `location` is [NSNotFound] and its `length`
// is 0 if and only if [HasMarkedText] returns [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AppKit/NSTextInputClient/markedRange()
func (o NSTextInputClientObject) MarkedRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("markedRange"))
	return rv
	}
// Returns the range of selected text.
//
// # Return Value
// 
// The range of selected text or `{NSNotFound, 0}` if there is no selection.
//
// # Discussion
// 
// The returned range measures from the start of the receiver’s text
// storage, that is, from 0 to the document length.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextInputClient/selectedRange()
func (o NSTextInputClientObject) SelectedRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("selectedRange"))
	return rv
	}
// Replaces a specified range in the receiver’s text storage with the given
// string and sets the selection.
//
// string: The string to insert. Can be either an [NSString] or [NSAttributedString]
// instance.
//
// selectedRange: The range to set as the selection, computed from the beginning of the
// inserted string.
//
// replacementRange: The range to replace, computed from the beginning of the marked text.
//
// # Discussion
// 
// If there is no marked text, the current selection is replaced. If there is
// no selection, the string is inserted at the insertion point.
// 
// When `aString` is an [NSString] object, the receiver is expected to render
// the marked text with distinguishing appearance (for example, [NSTextView]
// renders with [MarkedTextAttributes]).
//
// See: https://developer.apple.com/documentation/AppKit/NSTextInputClient/setMarkedText(_:selectedRange:replacementRange:)
func (o NSTextInputClientObject) SetMarkedTextSelectedRangeReplacementRange(string_ objectivec.IObject, selectedRange foundation.NSRange, replacementRange foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setMarkedText:selectedRange:replacementRange:"), string_, selectedRange, replacementRange)
	}
// Unmarks the marked text.
//
// # Discussion
// 
// The receiver removes any marking from pending input text and disposes of
// the marked text as it wishes. The text view should accept the marked text
// as if it had been inserted normally. If there is no marked text, the
// invocation of this method has no effect.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextInputClient/unmarkText()
func (o NSTextInputClientObject) UnmarkText() {
	objc.Send[struct{}](o.ID, objc.Sel("unmarkText"))
	}
// Returns an array of attribute names recognized by the receiver.
//
// # Return Value
// 
// An array of [NSString] objects representing names for the supported
// attributes.
//
// # Discussion
// 
// Returns an empty array if no attributes are supported. See
// NSAttributedString Application Kit Additions Reference for the set of
// string constants representing standard attributes.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextInputClient/validAttributesForMarkedText()
func (o NSTextInputClientObject) ValidAttributesForMarkedText() []string {
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("validAttributesForMarkedText"))
	return objc.ConvertSliceToStrings(rv)
	}
// Returns an attributed string derived from the given range in the
// receiver’s text storage.
//
// range: The range in the text storage from which to create the returned string.
//
// actualRange: The actual range of the returned string if it was adjusted, for example, to
// a grapheme cluster boundary or for performance or other reasons. [NULL] if
// range was not adjusted.
//
// # Return Value
// 
// The string created from the given range. May return `nil`.
//
// # Discussion
// 
// An implementation of this method should be prepared for `aRange` to be out
// of bounds. For example, the InkWell text input service can ask for the
// contents of the text input client that extends beyond the document’s
// range. In this case, you should return the intersection of the document’s
// range and `aRange`. If the location of `aRange` is completely outside of
// the document’s range, return `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextInputClient/attributedSubstring(forProposedRange:actualRange:)
func (o NSTextInputClientObject) AttributedSubstringForProposedRangeActualRange(range_ foundation.NSRange, actualRange foundation.NSRange) foundation.NSAttributedString {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("attributedSubstringForProposedRange:actualRange:"), range_, actualRange)
	return foundation.NSAttributedStringFromID(rv)
	}
// Inserts the given string into the receiver, replacing the specified
// content.
//
// string: The text to insert, either an [NSString] or [NSAttributedString] instance.
//
// replacementRange: The range of content to replace in the receiver’s text storage.
//
// # Discussion
// 
// This method is the entry point for inserting text typed by the user and is
// generally not suitable for other purposes. Programmatic modification of the
// text is best done by operating on the text storage directly. Because this
// method pertains to the actions of the user, the text view must be editable
// for the insertion to work.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextInputClient/insertText(_:replacementRange:)
func (o NSTextInputClientObject) InsertTextReplacementRange(string_ objectivec.IObject, replacementRange foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("insertText:replacementRange:"), string_, replacementRange)
	}
// Returns the index of the character whose bounding rectangle includes the
// given point.
//
// point: The point to test, in screen coordinates.
//
// # Return Value
// 
// The character index, measured from the start of the receiver’s text
// storage, of the character containing the given point. Returns [NSNotFound]
// if the cursor is not within a character’s bounding rectangle.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextInputClient/characterIndex(for:)
func (o NSTextInputClientObject) CharacterIndexForPoint(point corefoundation.CGPoint) uint {
	rv := objc.Send[uint](o.ID, objc.Sel("characterIndexForPoint:"), point)
	return rv
	}
// Returns the first logical boundary rectangle for characters in the given
// range.
//
// range: The character range whose boundary rectangle is returned.
//
// actualRange: If non-[NULL], contains the character range corresponding to the returned
// area if it was adjusted, for example, to a grapheme cluster boundary or
// characters in the first line fragment.
//
// # Return Value
// 
// The boundary rectangle for the given range of characters, in screen
// coordinates. The rectangle’s `size` value can be negative if the text
// flows to the left.
//
// # Discussion
// 
// If `aRange` spans multiple lines of text in the text view, the rectangle
// returned is the one surrounding the characters in the first line. In that
// case `actualRange` contains the range covered by the first rect, so you can
// query all line fragments by invoking this method repeatedly. If the length
// of `aRange` is 0 (as it would be if there is nothing selected at the
// insertion point), the rectangle coincides with the insertion point, and its
// width is 0.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextInputClient/firstRect(forCharacterRange:actualRange:)
func (o NSTextInputClientObject) FirstRectForCharacterRangeActualRange(range_ foundation.NSRange, actualRange foundation.NSRange) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("firstRectForCharacterRange:actualRange:"), range_, actualRange)
	return rv
	}
// See: https://developer.apple.com/documentation/AppKit/NSTextInputClient/documentVisibleRect
func (o NSTextInputClientObject) DocumentVisibleRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("documentVisibleRect"))
	return rv
	}
// See: https://developer.apple.com/documentation/AppKit/NSTextInputClient/unionRectInVisibleSelectedRange
func (o NSTextInputClientObject) UnionRectInVisibleSelectedRange() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("unionRectInVisibleSelectedRange"))
	return rv
	}
// Invokes the action specified by the given selector.
//
// selector: The selector to invoke.
//
// # Discussion
// 
// If `selector` cannot be invoked, then `` should not pass this message up
// the responder chain. [NSResponder] also implements this method, and it does
// forward uninvokable commands up the responder chain, but a text view should
// not. A text view implementing the [NSTextInputClient] protocol inherits
// from [NSView], which inherits from [NSResponder], so your implementation of
// this method will override the one in [NSResponder]. It should not call
// `super`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextInputClient/doCommand(by:)
func (o NSTextInputClientObject) DoCommandBySelector(selector objc.SEL) {
	objc.Send[struct{}](o.ID, objc.Sel("doCommandBySelector:"), selector)
	}
// A Boolean value that indicates whether the document supports adaptive
// images in the input.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextInputClient/supportsAdaptiveImageGlyph
func (o NSTextInputClientObject) SupportsAdaptiveImageGlyph() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("supportsAdaptiveImageGlyph"))
	return rv
	}
// Returns an attributed string representing the receiver’s text storage.
//
// # Return Value
// 
// The attributed string of the receiver’s text storage.
//
// # Discussion
// 
// Implementation of this method is optional. A class adopting the
// [NSTextInputClient] protocol can implement this interface if it can be done
// efficiently to enable callers of this interface to access arbitrary
// portions of the receiver’s content more efficiently.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextInputClient/attributedString()
func (o NSTextInputClientObject) AttributedString() foundation.NSAttributedString {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("attributedString"))
	return foundation.NSAttributedStringFromID(rv)
	}
// Returns the baseline position of a given character relative to the origin
// of rectangle returned by [FirstRectForCharacterRangeActualRange].
//
// anIndex: Index of the character whose baseline is tested.
//
// # Return Value
// 
// The vertical distance, in points, between the baseline of the character at
// `anIndex` and the rectangle origin.
//
// # Discussion
// 
// Implementation of this method is optional. This information allows the
// caller to determine finer-grained character positioning within the text
// storage of the text view adopting [NSTextInputClient].
//
// See: https://developer.apple.com/documentation/AppKit/NSTextInputClient/baselineDeltaForCharacter(at:)
func (o NSTextInputClientObject) BaselineDeltaForCharacterAtIndex(anIndex uint) float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("baselineDeltaForCharacterAtIndex:"), anIndex)
	return rv
	}
// Informs the text input management system whether the protocol-conforming
// client renders the character at the given index vertically.
//
// charIndex: The index of the character to test.
//
// # Return Value
// 
// [true] if the character is rendered vertically; otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextInputClient/drawsVerticallyForCharacter(at:)
func (o NSTextInputClientObject) DrawsVerticallyForCharacterAtIndex(charIndex uint) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("drawsVerticallyForCharacterAtIndex:"), charIndex)
	return rv
	}
// Returns the fraction of the distance from the left side of the character to
// the right side that a given point lies.
//
// point: The point to test.
//
// # Return Value
// 
// The fraction of the distance `aPoint` is through the glyph in which it
// lies. May be 0 or 1 if `aPoint` is not within the bounding rectangle of a
// glyph (0 if the point is to the left or above the glyph; 1 if it’s to the
// right or below).
//
// # Discussion
// 
// Implementation of this method is optional. This allows caller to perform
// precise selection handling.
// 
// For purposes such as dragging out a selection or placing the insertion
// point, a partial percentage less than or equal to 0.5 indicates that
// `aPoint` should be considered as falling before the glyph; a partial
// percentage greater than 0.5 indicates that it should be considered as
// falling after the glyph. If the nearest glyph doesn’t lie under `aPoint`
// at all (for example, if `aPoint` is beyond the beginning or end of a line),
// this ratio is 0 or 1.
// 
// For example, if the glyph stream contains the glyphs “A” and “b”,
// with the width of “A” being 13 points, and `aPoint` is 8 points from
// the left side of “A”, then the fraction of the distance is 8/13, or
// 0.615. In this case, the `aPoint` should be considered as falling between
// “A” and “b” for purposes such as dragging out a selection or
// placing the insertion point.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextInputClient/fractionOfDistanceThroughGlyph(for:)
func (o NSTextInputClientObject) FractionOfDistanceThroughGlyphForPoint(point corefoundation.CGPoint) float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("fractionOfDistanceThroughGlyphForPoint:"), point)
	return rv
	}
// See: https://developer.apple.com/documentation/AppKit/NSTextInputClient/preferredTextAccessoryPlacement()
func (o NSTextInputClientObject) PreferredTextAccessoryPlacement() NSTextCursorAccessoryPlacement {
	rv := objc.Send[NSTextCursorAccessoryPlacement](o.ID, objc.Sel("preferredTextAccessoryPlacement"))
	return rv
	}
// Returns the window level of the receiver.
//
// # Return Value
// 
// The window level of the receiver.
//
// # Discussion
// 
// Implementation of this method is optional. A class adopting
// [NSTextInputClient] can implement this interface to specify its window
// level if it is higher than [NSFloatingWindowLevel].
//
// See: https://developer.apple.com/documentation/AppKit/NSTextInputClient/windowLevel()
func (o NSTextInputClientObject) WindowLevel() int {
	rv := objc.Send[int](o.ID, objc.Sel("windowLevel"))
	return rv
	}
// Inserts an adaptive image into the text at the specifed location.
//
// adaptiveImageGlyph: The adaptive image to add to the text.
//
// replacementRange: The text range at which to insert the image.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextInputClient/insert(_:replacementRange:)
func (o NSTextInputClientObject) InsertAdaptiveImageGlyphReplacementRange(adaptiveImageGlyph INSAdaptiveImageGlyph, replacementRange foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("insertAdaptiveImageGlyph:replacementRange:"), adaptiveImageGlyph, replacementRange)
	}

