// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"fmt"
	"unsafe"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)
var _ = fmt.Sprintf

// A set of optional methods that text view delegates can use to manage selection, set text attributes, work with the spell checker, and more.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate
type NSTextViewDelegate interface {
	objectivec.IObject
	NSTextDelegate
}

// NSTextViewDelegateObject wraps an existing Objective-C object that conforms to the NSTextViewDelegate protocol.
type NSTextViewDelegateObject struct {
	objectivec.Object
}
func (o NSTextViewDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSTextViewDelegateObjectFromID constructs a [NSTextViewDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSTextViewDelegateObjectFromID(id objc.ID) NSTextViewDelegateObject {
	return NSTextViewDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Returns the undo manager for the specified text view.
//
// view: The text view whose undo manager should be returned.
//
// # Return Value
// 
// The undo manager for `view`.
//
// # Discussion
// 
// This method provides the flexibility to return a custom undo manager for
// the text view. Although [NSTextView] implements undo and redo for changes
// to text, applications may need a custom undo manager to handle interactions
// between changes to text and changes to other items in the application.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/undoManager(for:)

func (o NSTextViewDelegateObject) UndoManagerForTextView(view INSTextView) foundation.NSUndoManager {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("undoManagerForTextView:"), view)
	return foundation.NSUndoManagerFromID(rv)
	}

// Returns the actual tooltip to display.
//
// textView: The text view sending the message.
//
// tooltip: The proposed tooltip to display.
//
// characterIndex: The location in `textView`.
//
// # Return Value
// 
// The actual tooltip to display, or `nil` to suppress display of the tooltip.
//
// # Discussion
// 
// The tooltip string is the value of the [toolTip] attribute at
// `characterIndex`.
//
// [toolTip]: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key/toolTip
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textView(_:willDisplayToolTip:forCharacterAt:)

func (o NSTextViewDelegateObject) TextViewWillDisplayToolTipForCharacterAtIndex(textView INSTextView, tooltip string, characterIndex uint) string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("textView:willDisplayToolTip:forCharacterAtIndex:"), textView, objc.String(tooltip), characterIndex)
	return foundation.NSStringFromID(rv).String()
	}

// Returns a URL representing the document contents for a text attachment.
//
// textView: The text view sending the message.
//
// textAttachment: The text attachment object containing an [NSFileWrapper] object that holds
// the contents of the attached file.
//
// charIndex: The character index of the text attachment.
//
// # Return Value
// 
// The absolute URL for the document contents represented by `textAttachment`.
//
// # Discussion
// 
// The returned [NSURL] object is used by the text view to provide default
// behaviors involving text attachments such as Quick Look and
// double-clicking. For example, the [NSTextView] method
// [QuickLookPreviewableItemsInRanges] uses this method for mapping text
// attachments to their corresponding document URLs, and [NSTextView] invokes
// the [NSWorkspace] method [OpenURL] with the URL returned from this method
// when the delegate has no [TextViewDoubleClickedOnCellInRectAtIndex]
// implementation.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textView(_:urlForContentsOf:at:)

func (o NSTextViewDelegateObject) TextViewURLForContentsOfTextAttachmentAtIndex(textView INSTextView, textAttachment INSTextAttachment, charIndex uint) foundation.INSURL {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("textView:URLForContentsOfTextAttachment:atIndex:"), textView, textAttachment, charIndex)
	return foundation.NSURLFromID(rv)
	}

// Returns the actual range to select.
//
// textView: The text view sending the message. This is the first text view in a series
// shared by a layout manager, not necessarily the text view displaying the
// selected text.
//
// oldSelectedCharRange: The original range of the selection.
//
// newSelectedCharRange: The proposed character range for the new selection.
//
// # Return Value
// 
// The actual character range for the new selection.
//
// # Discussion
// 
// This method is invoked before a text view finishes changing the
// selection—that is, when the last argument to a
// [SetSelectedRangeAffinityStillSelecting] message is [false].
// 
// Non-selectable text views do not process any mouse events. If for some
// reason it is necessary to disallow user selection change in a text view
// that handles mouse events, this can be achieved by making the text view
// selectable but implementing this delegate method to disallow selection
// changes.
// 
// # Special Considerations
// 
// In macOS 10.4 and later, if a delegate implements this delegate method and
// not its multiple-selection replacement,
// [TextViewWillChangeSelectionFromCharacterRangesToCharacterRanges], then
// multiple selection is effectively disallowed; attempts to set the selected
// ranges call the old delegate method with the first subrange, and afterwards
// only a single selected range is set.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textView(_:willChangeSelectionFromCharacterRange:toCharacterRange:)

func (o NSTextViewDelegateObject) TextViewWillChangeSelectionFromCharacterRangeToCharacterRange(textView INSTextView, oldSelectedCharRange foundation.NSRange, newSelectedCharRange foundation.NSRange) foundation.NSRange {
	
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("textView:willChangeSelectionFromCharacterRange:toCharacterRange:"), textView, oldSelectedCharRange, newSelectedCharRange)
	return rv
	}

// Returns the actual character ranges to select.
//
// textView: The text view sending the message. This is the first text view in a series
// shared by a layout manager, not necessarily the text view displaying the
// selected text.
//
// oldSelectedCharRanges: An array containing the original ranges of the selection. This must be a
// non-`nil`, non-empty array of objects responding to the [NSValue] method
// `rangeValue`, and in addition its elements must be sorted, non-overlapping,
// non-contiguous, and (except for the case of a single range) have
// non-zero-length.
//
// newSelectedCharRanges: An array containing the proposed character ranges for the new selection.
// This must be a non-`nil`, non-empty array of objects responding to the
// [NSValue] method `rangeValue`, and in addition its elements must be sorted,
// non-overlapping, non-contiguous, and (except for the case of a single
// range) have non-zero-length.
//
// # Return Value
// 
// An array containing the actual character ranges for the new selection.
//
// # Discussion
// 
// Invoked before an [NSTextView] object finishes changing the
// selection—that is, when the last argument to a
// [SetSelectedRangeAffinityStillSelecting] or
// [SetSelectedRangesAffinityStillSelecting] message is [false].
// 
// Non-selectable text views do not process any mouse events. If for some
// reason it is necessary to disallow user selection change in a text view
// that handles mouse events, this can be achieved by making the text view
// selectable but implementing this delegate method to disallow selection
// changes.
// 
// If a delegate implements both this method and
// [TextViewWillChangeSelectionFromCharacterRangeToCharacterRange], then the
// latter is ignored.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textView(_:willChangeSelectionFromCharacterRanges:toCharacterRanges:)

func (o NSTextViewDelegateObject) TextViewWillChangeSelectionFromCharacterRangesToCharacterRanges(textView INSTextView, oldSelectedCharRanges []foundation.NSValue, newSelectedCharRanges []foundation.NSValue) []foundation.NSValue {
	
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("textView:willChangeSelectionFromCharacterRanges:toCharacterRanges:"), textView, objectivec.IObjectSliceToNSArray(oldSelectedCharRanges), objectivec.IObjectSliceToNSArray(newSelectedCharRanges))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSValue {
		return foundation.NSValueFromID(id)
	})
	}

// Sent when the selection changes in the text view.
//
// notification: A notification named [didChangeSelectionNotification].
// //
// [didChangeSelectionNotification]: https://developer.apple.com/documentation/AppKit/NSTextView/didChangeSelectionNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textViewDidChangeSelection(_:)

func (o NSTextViewDelegateObject) TextViewDidChangeSelection(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("textViewDidChangeSelection:"), notification)
	}

// Returns an array of objects that represent the elements of a selection.
//
// # Return Value
// 
// An array of objects that represent the selection.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textView(_:candidatesForSelectedRange:)

func (o NSTextViewDelegateObject) TextViewCandidatesForSelectedRange(textView INSTextView, selectedRange foundation.NSRange) foundation.INSArray {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("textView:candidatesForSelectedRange:"), textView, selectedRange)
	return foundation.NSArrayFromID(rv)
	}

// Returns a Boolean value that indicates whether to select the text object at
// the index.
//
// textView: The text view that sent the message.
//
// index: The index that represents the start of the candidate text to evaluate.
//
// # Return Value
// 
// Returns [true] if the framework selects the text.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textView(_:shouldSelectCandidateAt:)

func (o NSTextViewDelegateObject) TextViewShouldSelectCandidateAtIndex(textView INSTextView, index uint) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("textView:shouldSelectCandidateAtIndex:"), textView, index)
	return rv
	}

// Returns and array of touch bar elements for the framework to update.
//
// textView: The text view that sent the message.
//
// identifiers: An array of touch bar identifiers to evaluate.
//
// # Return Value
// 
// Returns an array of [NSTouchBarItemIdentifier] elements for framework to
// update.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textView(_:shouldUpdateTouchBarItemIdentifiers:)

func (o NSTextViewDelegateObject) TextViewShouldUpdateTouchBarItemIdentifiers(textView INSTextView, identifiers []string) []string {
	
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("textView:shouldUpdateTouchBarItemIdentifiers:"), textView, objectivec.StringSliceToNSArray(identifiers))
	return objc.ConvertSliceToStrings(rv)
	}

// Returns the writable pasteboard types for a given cell.
//
// view: The text view sending the message.
//
// cell: The cell in question.
//
// charIndex: The character index in the text view that was clicked.
//
// # Return Value
// 
// An array of types that can be written to the pasteboard for `cell`.
//
// # Discussion
// 
// This method is invoked after the user clicks `cell` at the specified
// `charIndex` location in `aTextView`. If the
// [TextViewDraggedCellInRectEventAtIndex] is not used, this method and
// [TextViewWriteCellAtIndexToPasteboardType] allow `aTextView` to take care
// of attachment dragging and pasting, with the delegate responsible only for
// writing the attachment to the pasteboard.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textView(_:writablePasteboardTypesFor:at:)

func (o NSTextViewDelegateObject) TextViewWritablePasteboardTypesForCellAtIndex(view INSTextView, cell NSTextAttachmentCell, charIndex uint) []string {
	
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("textView:writablePasteboardTypesForCell:atIndex:"), view, cell, charIndex)
	return objc.ConvertSliceToStrings(rv)
	}

// Returns whether data of the specified type for the given cell could be
// written to the specified pasteboard.
//
// view: The text view sending the message.
//
// cell: The cell whose contents should be written to the pasteboard.
//
// charIndex: The index at which the cell was accessed.
//
// pboard: The pasteboard to which the cell’s contents should be written.
//
// type: The type of data that should be written.
//
// # Return Value
// 
// [true] if the write succeeded, [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The receiver should attempt to write the `cell` to `pboard` with the given
// `type`, and return success or failure.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textView(_:write:at:to:type:)

func (o NSTextViewDelegateObject) TextViewWriteCellAtIndexToPasteboardType(view INSTextView, cell NSTextAttachmentCell, charIndex uint, pboard INSPasteboard, type_ NSPasteboardType) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("textView:writeCell:atIndex:toPasteboard:type:"), view, cell, charIndex, pboard, objc.String(string(type_)))
	return rv
	}

// Sent when a text view needs to determine if text in a specified range
// should be changed.
//
// textView: The text view sending the message. This is the first text view in a series
// shared by a layout manager, not necessarily the text view displaying the
// selected text.
//
// affectedCharRange: The range of characters to be replaced.
//
// replacementString: The characters that will replace the characters in `affectedCharRange`;
// `nil` if only text attributes are being changed.
//
// # Return Value
// 
// [true] to allow the replacement, or [false] to reject the change.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// If a delegate implements this method and not its multiple-selection
// replacement, [TextViewShouldChangeTextInRangesReplacementStrings], it is
// called with an appropriate range and string. If a delegate implements the
// new method, then this one is ignored.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textView(_:shouldChangeTextIn:replacementString:)

func (o NSTextViewDelegateObject) TextViewShouldChangeTextInRangeReplacementString(textView INSTextView, affectedCharRange foundation.NSRange, replacementString string) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("textView:shouldChangeTextInRange:replacementString:"), textView, affectedCharRange, objc.String(replacementString))
	return rv
	}

// Sent when a text view needs to determine if text in an array of specified
// ranges should be changed.
//
// textView: The text view sending the message. This is the first text view in a series
// shared by a layout manager, not necessarily the text view displaying the
// selected text.
//
// affectedRanges: The array of ranges of characters to be replaced. This array must be a
// non-nil, non-empty array of objects responding to the NSValue `rangeValue`
// method, and in addition its elements must be sorted, non-overlapping,
// non-contiguous, and (except for the case of a single range) have
// non-zero-length.
//
// replacementStrings: The array of strings that will replace the characters in `affectedRanges`,
// one string for each range; `nil` if only text attributes are being changed.
//
// # Return Value
// 
// [true] to allow the replacement, or [false] to reject the change.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textView(_:shouldChangeTextInRanges:replacementStrings:)

func (o NSTextViewDelegateObject) TextViewShouldChangeTextInRangesReplacementStrings(textView INSTextView, affectedRanges []foundation.NSValue, replacementStrings []string) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("textView:shouldChangeTextInRanges:replacementStrings:"), textView, objectivec.IObjectSliceToNSArray(affectedRanges), objectivec.StringSliceToNSArray(replacementStrings))
	return rv
	}

// Sent when the typing attributes are changed.
//
// textView: The text view sending the message.
//
// oldTypingAttributes: The old typing attributes.
//
// newTypingAttributes: The proposed typing attributes.
//
// # Return Value
// 
// The actual new typing attributes.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textView(_:shouldChangeTypingAttributes:toAttributes:)

func (o NSTextViewDelegateObject) TextViewShouldChangeTypingAttributesToAttributes(textView INSTextView, oldTypingAttributes foundation.INSDictionary, newTypingAttributes foundation.INSDictionary) foundation.INSDictionary {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("textView:shouldChangeTypingAttributes:toAttributes:"), textView, oldTypingAttributes, newTypingAttributes)
	return foundation.NSDictionaryFromID(rv)
	}

// Sent when a text view’s typing attributes change.
//
// notification: A notification named [didChangeTypingAttributesNotification].
// //
// [didChangeTypingAttributesNotification]: https://developer.apple.com/documentation/AppKit/NSTextView/didChangeTypingAttributesNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textViewDidChangeTypingAttributes(_:)

func (o NSTextViewDelegateObject) TextViewDidChangeTypingAttributes(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("textViewDidChangeTypingAttributes:"), notification)
	}

// Sent when the user clicks a cell.
//
// textView: The text view sending the message.
//
// cell: The cell clicked by the user.
//
// cellFrame: The frame of the clicked cell.
//
// charIndex: The character index of the clicked cell.
//
// # Discussion
// 
// The delegate can use this message as its cue to perform an action or select
// the attachment cell’s character. `aTextView` is the first text view in a
// series shared by a layout manager, not necessarily the one that draws
// `cell`.
// 
// The delegate may subsequently receive a
// [TextViewDoubleClickedOnCellInRectAtIndex] message if the user continues to
// perform a double click.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textView(_:clickedOn:in:at:)

func (o NSTextViewDelegateObject) TextViewClickedOnCellInRectAtIndex(textView INSTextView, cell NSTextAttachmentCell, cellFrame corefoundation.CGRect, charIndex uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("textView:clickedOnCell:inRect:atIndex:"), textView, cell, cellFrame, charIndex)
	}

// Sent when the user double-clicks a cell.
//
// textView: The text view sending the message.
//
// cell: The cell double-clicked by the user.
//
// cellFrame: The frame of the double-clicked cell.
//
// charIndex: The character index of the double-clicked cell.
//
// # Discussion
// 
// The delegate can use this message as its cue to perform an action, such as
// opening the file represented by the attachment. `aTextView` is the first
// text view in a series shared by a layout manager, not necessarily the one
// that draws `cell`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textView(_:doubleClickedOn:in:at:)

func (o NSTextViewDelegateObject) TextViewDoubleClickedOnCellInRectAtIndex(textView INSTextView, cell NSTextAttachmentCell, cellFrame corefoundation.CGRect, charIndex uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("textView:doubleClickedOnCell:inRect:atIndex:"), textView, cell, cellFrame, charIndex)
	}

// Sent after the user clicks a link.
//
// textView: The text view sending the message.
//
// link: The link that was clicked; the value of [link].
// //
// [link]: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key/link
//
// charIndex: The character index where the click occurred, indexed within the text
// storage.
//
// # Return Value
// 
// [true] if the click was handled; otherwise, [false] to allow the next
// responder to handle it.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The delegate can use this method to handle the click on the link. It is
// invoked by [ClickedOnLinkAtIndex].
// 
// The `charIndex` parameter is a character index somewhere in the range of
// the link attribute. If the user actually physically clicked the link, then
// it should be the character that was originally clicked. In some cases a
// link may be opened indirectly or programmatically, in which case a
// character index somewhere in the range of the link attribute is supplied.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textView(_:clickedOnLink:at:)

func (o NSTextViewDelegateObject) TextViewClickedOnLinkAtIndex(textView INSTextView, link objectivec.IObject, charIndex uint) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("textView:clickedOnLink:atIndex:"), textView, link, charIndex)
	return rv
	}

// Sent when the spelling state is changed.
//
// textView: The text view sending the message.
//
// value: The proposed spelling state value to set. Possible values, for the
// temporary attribute on the layout manager using the key
// NSSpellingStateAttributeName, are:
// 
// - [NSSpellingStateSpellingFlag] to highlight spelling issues. -
// [NSSpellingStateGrammarFlag] to highlight grammar issues.
// //
// [NSSpellingStateGrammarFlag]: https://developer.apple.com/documentation/AppKit/NSSpellingState/NSSpellingStateGrammarFlag
// [NSSpellingStateSpellingFlag]: https://developer.apple.com/documentation/AppKit/NSSpellingState/NSSpellingStateSpellingFlag
//
// affectedCharRange: The character range over which to set the given spelling state.
//
// # Return Value
// 
// The actual spelling state to set.
//
// # Discussion
// 
// Delegate only. Allows delegate to control the setting of spelling and
// grammar indicators.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textView(_:shouldSetSpellingState:range:)

func (o NSTextViewDelegateObject) TextViewShouldSetSpellingStateRange(textView INSTextView, value int, affectedCharRange foundation.NSRange) int {
	
	rv := objc.Send[int](o.ID, objc.Sel("textView:shouldSetSpellingState:range:"), textView, value, affectedCharRange)
	return rv
	}

// Invoked to allow the delegate to modify the text checking process before it
// occurs.
//
// view: The text view sending the message.
//
// range: The range to be checked.
//
// options: A dictionary of values used during the checking process to perform. See
// Spell Checking Option Dictionary Keys for the supported values.
//
// checkingTypes: The type of checking to be performed, passed by-reference. The possible
// constants are listed in [NSTextCheckingTypes] and can be combined using the
// C bit-wise [OR] operator to perform multiple checks at the same time.
// 
// You can change this parameter to alter the types of checking to be
// performed.
// //
// [NSTextCheckingTypes]: https://developer.apple.com/documentation/Foundation/NSTextCheckingTypes
//
// # Return Value
// 
// A dictionary containing an alternative to the options `dictionary`.
//
// # Discussion
// 
// Invoked by [CheckTextInRangeTypesOptions], this method allows control over
// text checking `options`s (via the return value) or types (by modifying the
// flags pointed to by the inout parameter `checkingTypes`)
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textView(_:willCheckTextIn:options:types:)

func (o NSTextViewDelegateObject) TextViewWillCheckTextInRangeOptionsTypes(view INSTextView, range_ foundation.NSRange, options foundation.INSDictionary, checkingTypes unsafe.Pointer) foundation.INSDictionary {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("textView:willCheckTextInRange:options:types:"), view, range_, options, checkingTypes)
	return foundation.NSDictionaryFromID(rv)
	}

// Invoked to allow the delegate to modify the text checking results after
// checking has occurred.
//
// view: The text view sending the message.
//
// range: The range that was checked.
//
// checkingTypes: The type of checking that was performed. The possible constants are listed
// in [NSTextCheckingTypes] and can be combined using the C bit-wise [OR]
// operator to perform multiple checks at the same time.
// //
// [NSTextCheckingTypes]: https://developer.apple.com/documentation/Foundation/NSTextCheckingTypes
//
// options: A dictionary of values used during the checking process to perform. See
// Spell Checking Option Dictionary Keys for the supported values.
//
// results: An array of [NSTextCheckingResult] instances.
// //
// [NSTextCheckingResult]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult
//
// orthography: The orthography of the text.
//
// wordCount: The number of words checked.
//
// # Return Value
// 
// An array of [NSTextCheckingResult] instances. You can return the results
// array as is, or an altered array of [NSTextCheckingResult] objects.
//
// [NSTextCheckingResult]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult
//
// # Discussion
// 
// Invoked by
// [HandleTextCheckingResultsForRangeTypesOptionsOrthographyWordCount], this
// method allows observation of text checking, or modification of the results
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textView(_:didCheckTextIn:types:options:results:orthography:wordCount:)

func (o NSTextViewDelegateObject) TextViewDidCheckTextInRangeTypesOptionsResultsOrthographyWordCount(view INSTextView, range_ foundation.NSRange, checkingTypes uint64, options foundation.INSDictionary, results []foundation.NSTextCheckingResult, orthography foundation.NSOrthography, wordCount int) []foundation.NSTextCheckingResult {
	
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("textView:didCheckTextInRange:types:options:results:orthography:wordCount:"), view, range_, checkingTypes, options, objectivec.IObjectSliceToNSArray(results), orthography, wordCount)
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSTextCheckingResult {
		return foundation.NSTextCheckingResultFromID(id)
	})
	}

//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textViewWritingToolsWillBegin(_:)

func (o NSTextViewDelegateObject) TextViewWritingToolsWillBegin(textView INSTextView) {
	
	objc.Send[struct{}](o.ID, objc.Sel("textViewWritingToolsWillBegin:"), textView)
	}

//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textViewWritingToolsDidEnd(_:)

func (o NSTextViewDelegateObject) TextViewWritingToolsDidEnd(textView INSTextView) {
	
	objc.Send[struct{}](o.ID, objc.Sel("textViewWritingToolsDidEnd:"), textView)
	}

//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textView(_:writingToolsIgnoredRangesInEnclosingRange:)

func (o NSTextViewDelegateObject) TextViewWritingToolsIgnoredRangesInEnclosingRange(textView INSTextView, enclosingRange foundation.NSRange) []foundation.NSValue {
	
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("textView:writingToolsIgnoredRangesInEnclosingRange:"), textView, enclosingRange)
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSValue {
		return foundation.NSValueFromID(id)
	})
	}

// Sent when the user attempts to drag a cell.
//
// view: The text view sending the message.
//
// cell: The cell being dragged.
//
// rect: The rectangle from which the cell was dragged.
//
// event: The mouse-down event that preceded the mouse-dragged event.
//
// charIndex: The character position where the mouse button was clicked.
//
// # Discussion
// 
// The delegate can use this message as its cue to initiate a dragging
// operation.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textView(_:draggedCell:in:event:at:)

func (o NSTextViewDelegateObject) TextViewDraggedCellInRectEventAtIndex(view INSTextView, cell NSTextAttachmentCell, rect corefoundation.CGRect, event INSEvent, charIndex uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("textView:draggedCell:inRect:event:atIndex:"), view, cell, rect, event, charIndex)
	}

// Returns the actual completions for a partial word.
//
// textView: The text view sending the message.
//
// words: The proposed array of completions.
//
// charRange: The range of characters to be completed.
//
// index: On return, the index of the initially selected completion. The default is
// 0, and –1 indicates no selection.
//
// # Return Value
// 
// The actual array of completions that will be presented for the partial word
// at the given range. Returning `nil` or a zero-length array suppresses
// completion.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textView(_:completions:forPartialWordRange:indexOfSelectedItem:)

func (o NSTextViewDelegateObject) TextViewCompletionsForPartialWordRangeIndexOfSelectedItem(textView INSTextView, words []string, charRange foundation.NSRange, index unsafe.Pointer) []string {
	
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("textView:completions:forPartialWordRange:indexOfSelectedItem:"), textView, objectivec.StringSliceToNSArray(words), charRange, index)
	return objc.ConvertSliceToStrings(rv)
	}

// Returns a sharing service picker for the current selection.
//
// textView: The text view.
//
// servicePicker: The service picker.
//
// items: The ranges of the items to share.
//
// # Return Value
// 
// An [NSSharingServicePicker] instance. The original sharing picker or a new
// sharing picker instance can be returned.
//
// # Discussion
// 
// Returns a sharing service picker created for items right before shown to
// the screen when the `` method. Return `nil` to remove the Share item from
// the menu.
// 
// The delegate is specified as the delegate for the [NSSharingServicePicker]
// instance.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textView(_:willShow:forItems:)

func (o NSTextViewDelegateObject) TextViewWillShowSharingServicePickerForItems(textView INSTextView, servicePicker INSSharingServicePicker, items foundation.INSArray) INSSharingServicePicker {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("textView:willShowSharingServicePicker:forItems:"), textView, servicePicker, items)
	return NSSharingServicePickerFromID(rv)
	}

// Sent to allow the delegate to perform the command for the text view.
//
// textView: The text view sending the message. This is the first text view in a series
// shared by a layout manager.
//
// commandSelector: The selector.
//
// # Return Value
// 
// [true] indicates that the delegate handled the command and the text view
// will not attempt to perform it; [false] indicates that the delegate did not
// handle the command the text view will attempt to perform it.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method is invoked by [NSTextView]’s `doCommand()` method.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textView(_:doCommandBy:)

func (o NSTextViewDelegateObject) TextViewDoCommandBySelector(textView INSTextView, commandSelector objc.SEL) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("textView:doCommandBySelector:"), textView, commandSelector)
	return rv
	}

// Allows delegate to control the context menu returned by the text view.
//
// view: The text view sending the message.
//
// menu: The proposed contextual menu.
//
// event: The mouse-down event that initiated the contextual menu’s display.
//
// charIndex: The character position where the mouse button was clicked.
//
// # Return Value
// 
// A menu to use as the contextual menu. You can return `menu` unaltered, or
// you can return a customized menu.
//
// # Discussion
// 
// This method allows the delegate to control the context menu returned by
// [MenuForEvent].
//
// See: https://developer.apple.com/documentation/AppKit/NSTextViewDelegate/textView(_:menu:for:at:)

func (o NSTextViewDelegateObject) TextViewMenuForEventAtIndex(view INSTextView, menu INSMenu, event INSEvent, charIndex uint) INSMenu {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("textView:menu:forEvent:atIndex:"), view, menu, event, charIndex)
	return NSMenuFromID(rv)
	}

// Informs the delegate that the text object has changed its characters or
// formatting attributes.
//
// # Discussion
// 
// The name of `aNotification` is [didChangeNotification].
//
// [didChangeNotification]: https://developer.apple.com/documentation/AppKit/NSText/didChangeNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSTextDelegate/textDidChange(_:)

func (o NSTextViewDelegateObject) TextDidChange(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("textDidChange:"), notification)
	}

// Invoked when a text object begins to change its text, this method requests
// permission for `aTextObject` to begin editing.
//
// # Discussion
// 
// If the delegate returns [true], the text object proceeds to make changes.
// If the delegate returns [false], the text object abandons the editing
// operation. This method is also invoked when the user drags and drops a file
// onto the text object.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextDelegate/textShouldBeginEditing(_:)

func (o NSTextViewDelegateObject) TextShouldBeginEditing(textObject INSText) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("textShouldBeginEditing:"), textObject)
	return rv
	}

// Informs the delegate that the text object has begun editing (that the user
// has begun changing it).
//
// # Discussion
// 
// The name of `aNotification` is [didBeginEditingNotification].
//
// [didBeginEditingNotification]: https://developer.apple.com/documentation/AppKit/NSText/didBeginEditingNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSTextDelegate/textDidBeginEditing(_:)

func (o NSTextViewDelegateObject) TextDidBeginEditing(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("textDidBeginEditing:"), notification)
	}

// Invoked from a text object’s implementation of [ResignFirstResponder],
// this method requests permission for `aTextObject` to end editing.
//
// # Discussion
// 
// If the delegate returns [true], the text object proceeds to finish editing
// and resign first responder status. If the delegate returns [false], the
// text object selects all of its text and remains the first responder.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextDelegate/textShouldEndEditing(_:)

func (o NSTextViewDelegateObject) TextShouldEndEditing(textObject INSText) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("textShouldEndEditing:"), textObject)
	return rv
	}

// Informs the delegate that the text object has finished editing (that it has
// resigned first responder status).
//
// # Discussion
// 
// The name of `aNotification` is [didEndEditingNotification].
//
// [didEndEditingNotification]: https://developer.apple.com/documentation/AppKit/NSText/didEndEditingNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSTextDelegate/textDidEndEditing(_:)

func (o NSTextViewDelegateObject) TextDidEndEditing(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("textDidEndEditing:"), notification)
	}

// NSTextViewDelegateConfig holds optional typed callbacks for [NSTextViewDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nstextviewdelegate
type NSTextViewDelegateConfig struct {

	// Managing the Selection
	// WillChangeSelectionFromCharacterRangeToCharacterRange — Returns the actual range to select.
	WillChangeSelectionFromCharacterRangeToCharacterRange func(textView NSTextView, oldSelectedCharRange foundation.NSRange, newSelectedCharRange foundation.NSRange) foundation.NSRange
	// DidChangeSelection — Sent when the selection changes in the text view.
	DidChangeSelection func(notification foundation.NSNotification)
	// CandidatesForSelectedRange — Returns an array of objects that represent the elements of a selection.
	CandidatesForSelectedRange func(textView NSTextView, selectedRange foundation.NSRange) foundation.INSArray

	// Setting Text Attributes
	// ShouldChangeTypingAttributesToAttributes — Sent when the typing attributes are changed.
	ShouldChangeTypingAttributesToAttributes func(textView NSTextView, oldTypingAttributes foundation.INSDictionary, newTypingAttributes foundation.INSDictionary) foundation.INSDictionary
	// DidChangeTypingAttributes — Sent when a text view’s typing attributes change.
	DidChangeTypingAttributes func(notification foundation.NSNotification)

	// Working With the Spelling Checker
	// ShouldSetSpellingStateRange — Sent when the spelling state is changed.
	ShouldSetSpellingStateRange func(textView NSTextView, value int, affectedCharRange foundation.NSRange) int

	// Responding to writing tools interactions
	WritingToolsWillBegin func(textView NSTextView)
	WritingToolsDidEnd func(textView NSTextView)

	// Other Methods
	// UndoManagerForTextView — Returns the undo manager for the specified text view.
	UndoManagerForTextView func(view NSTextView) foundation.NSUndoManager
	// URLForContentsOfTextAttachmentAtIndex — Returns a URL representing the document contents for a text attachment.
	URLForContentsOfTextAttachmentAtIndex func(textView NSTextView, textAttachment NSTextAttachment, charIndex uint) foundation.NSURL
	// ShouldSelectCandidateAtIndex — Returns a Boolean value that indicates whether to select the text object at the index.
	ShouldSelectCandidateAtIndex func(textView NSTextView, index uint) bool
	// WillCheckTextInRangeOptionsTypes — Invoked to allow the delegate to modify the text checking process before it occurs.
	WillCheckTextInRangeOptionsTypes func(view NSTextView, range_ foundation.NSRange, options foundation.INSDictionary, checkingTypes *uint64) foundation.INSDictionary
	// WillShowSharingServicePickerForItems — Returns a sharing service picker for the current selection.
	WillShowSharingServicePickerForItems func(textView NSTextView, servicePicker NSSharingServicePicker, items foundation.INSArray) NSSharingServicePicker
	// MenuForEventAtIndex — Allows delegate to control the context menu returned by the text view.
	MenuForEventAtIndex func(view NSTextView, menu NSMenu, event NSEvent, charIndex uint) NSMenu
}

// NewNSTextViewDelegate creates an Objective-C object implementing the [NSTextViewDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSTextViewDelegateObject] satisfies the [NSTextViewDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nstextviewdelegate
func NewNSTextViewDelegate(config NSTextViewDelegateConfig) NSTextViewDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSTextViewDelegate_%d", n)

	var methods []objc.MethodDef

	if config.UndoManagerForTextView != nil {
		fn := config.UndoManagerForTextView
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("undoManagerForTextView:"),
			Fn: func(self objc.ID, _cmd objc.SEL, viewID objc.ID) objc.ID {
				view := NSTextViewFromID(viewID)
				return fn(view).GetID()
			},
		})
	}

	if config.URLForContentsOfTextAttachmentAtIndex != nil {
		fn := config.URLForContentsOfTextAttachmentAtIndex
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("textView:URLForContentsOfTextAttachment:atIndex:"),
			Fn: func(self objc.ID, _cmd objc.SEL, textViewID objc.ID, textAttachmentID objc.ID, charIndex uint) objc.ID {
				textView := NSTextViewFromID(textViewID)
				textAttachment := NSTextAttachmentFromID(textAttachmentID)
				return fn(textView, textAttachment, charIndex).GetID()
			},
		})
	}

	if config.WillChangeSelectionFromCharacterRangeToCharacterRange != nil {
		fn := config.WillChangeSelectionFromCharacterRangeToCharacterRange
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("textView:willChangeSelectionFromCharacterRange:toCharacterRange:"),
			Fn: func(self objc.ID, _cmd objc.SEL, textViewID objc.ID, oldSelectedCharRange foundation.NSRange, newSelectedCharRange foundation.NSRange) foundation.NSRange {
				textView := NSTextViewFromID(textViewID)
				return fn(textView, oldSelectedCharRange, newSelectedCharRange)
			},
		})
	}

	if config.DidChangeSelection != nil {
		fn := config.DidChangeSelection
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("textViewDidChangeSelection:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.CandidatesForSelectedRange != nil {
		fn := config.CandidatesForSelectedRange
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("textView:candidatesForSelectedRange:"),
			Fn: func(self objc.ID, _cmd objc.SEL, textViewID objc.ID, selectedRange foundation.NSRange) objc.ID {
				textView := NSTextViewFromID(textViewID)
				return fn(textView, selectedRange).GetID()
			},
		})
	}

	if config.ShouldSelectCandidateAtIndex != nil {
		fn := config.ShouldSelectCandidateAtIndex
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("textView:shouldSelectCandidateAtIndex:"),
			Fn: func(self objc.ID, _cmd objc.SEL, textViewID objc.ID, index uint) bool {
				textView := NSTextViewFromID(textViewID)
				return fn(textView, index)
			},
		})
	}

	if config.ShouldChangeTypingAttributesToAttributes != nil {
		fn := config.ShouldChangeTypingAttributesToAttributes
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("textView:shouldChangeTypingAttributes:toAttributes:"),
			Fn: func(self objc.ID, _cmd objc.SEL, textViewID objc.ID, oldTypingAttributesID objc.ID, newTypingAttributesID objc.ID) objc.ID {
				textView := NSTextViewFromID(textViewID)
				oldTypingAttributes := foundation.NSDictionaryFromID(oldTypingAttributesID)
				newTypingAttributes := foundation.NSDictionaryFromID(newTypingAttributesID)
				return fn(textView, oldTypingAttributes, newTypingAttributes).GetID()
			},
		})
	}

	if config.DidChangeTypingAttributes != nil {
		fn := config.DidChangeTypingAttributes
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("textViewDidChangeTypingAttributes:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.ShouldSetSpellingStateRange != nil {
		fn := config.ShouldSetSpellingStateRange
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("textView:shouldSetSpellingState:range:"),
			Fn: func(self objc.ID, _cmd objc.SEL, textViewID objc.ID, value int, affectedCharRange foundation.NSRange) int {
				textView := NSTextViewFromID(textViewID)
				return fn(textView, value, affectedCharRange)
			},
		})
	}

	if config.WillCheckTextInRangeOptionsTypes != nil {
		fn := config.WillCheckTextInRangeOptionsTypes
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("textView:willCheckTextInRange:options:types:"),
			Fn: func(self objc.ID, _cmd objc.SEL, viewID objc.ID, range_ foundation.NSRange, optionsID objc.ID, checkingTypes *uint64) objc.ID {
				view := NSTextViewFromID(viewID)
				options := foundation.NSDictionaryFromID(optionsID)
				return fn(view, range_, options, checkingTypes).GetID()
			},
		})
	}

	if config.WritingToolsWillBegin != nil {
		fn := config.WritingToolsWillBegin
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("textViewWritingToolsWillBegin:"),
			Fn: func(self objc.ID, _cmd objc.SEL, textViewID objc.ID) {
				textView := NSTextViewFromID(textViewID)
				fn(textView)
			},
		})
	}

	if config.WritingToolsDidEnd != nil {
		fn := config.WritingToolsDidEnd
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("textViewWritingToolsDidEnd:"),
			Fn: func(self objc.ID, _cmd objc.SEL, textViewID objc.ID) {
				textView := NSTextViewFromID(textViewID)
				fn(textView)
			},
		})
	}

	if config.WillShowSharingServicePickerForItems != nil {
		fn := config.WillShowSharingServicePickerForItems
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("textView:willShowSharingServicePicker:forItems:"),
			Fn: func(self objc.ID, _cmd objc.SEL, textViewID objc.ID, servicePickerID objc.ID, itemsID objc.ID) objc.ID {
				textView := NSTextViewFromID(textViewID)
				servicePicker := NSSharingServicePickerFromID(servicePickerID)
				items := foundation.NSArrayFromID(itemsID)
				return fn(textView, servicePicker, items).GetID()
			},
		})
	}

	if config.MenuForEventAtIndex != nil {
		fn := config.MenuForEventAtIndex
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("textView:menu:forEvent:atIndex:"),
			Fn: func(self objc.ID, _cmd objc.SEL, viewID objc.ID, menuID objc.ID, eventID objc.ID, charIndex uint) objc.ID {
				view := NSTextViewFromID(viewID)
				menu := NSMenuFromID(menuID)
				event := NSEventFromID(eventID)
				return fn(view, menu, event, charIndex).GetID()
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSTextViewDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSTextViewDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSTextViewDelegateObjectFromID(instance)
}

