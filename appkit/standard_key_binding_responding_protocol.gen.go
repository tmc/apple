// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// Methods that responder subclasses implement to support key binding commands, such as inserting tabs and newlines, or moving the insertion point.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding
type NSStandardKeyBindingResponding interface {
	objectivec.IObject
}

// NSStandardKeyBindingRespondingObject wraps an existing Objective-C object that conforms to the NSStandardKeyBindingResponding protocol.
type NSStandardKeyBindingRespondingObject struct {
	objectivec.Object
}

func (o NSStandardKeyBindingRespondingObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSStandardKeyBindingRespondingObjectFromID constructs a [NSStandardKeyBindingRespondingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSStandardKeyBindingRespondingObjectFromID(id objc.ID) NSStandardKeyBindingRespondingObject {
	return NSStandardKeyBindingRespondingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Performs the given selector if possible.
//
// selector: The selector to perform.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/doCommand(by:)
func (o NSStandardKeyBindingRespondingObject) DoCommandBySelector(selector objc.SEL) {
	objc.Send[struct{}](o.ID, objc.Sel("doCommandBySelector:"), selector)
}

// Inserts a backtab character.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/insertBacktab(_:)
func (o NSStandardKeyBindingRespondingObject) InsertBacktab(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("insertBacktab:"), sender)
}

// Inserts a container break, such as a new page break.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/insertContainerBreak(_:)
func (o NSStandardKeyBindingRespondingObject) InsertContainerBreak(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("insertContainerBreak:"), sender)
}

// Inserts a double quotation mark without substituting a curly quotation
// mark.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/insertDoubleQuoteIgnoringSubstitution(_:)
func (o NSStandardKeyBindingRespondingObject) InsertDoubleQuoteIgnoringSubstitution(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("insertDoubleQuoteIgnoringSubstitution:"), sender)
}

// Inserts a line break character.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/insertLineBreak(_:)
func (o NSStandardKeyBindingRespondingObject) InsertLineBreak(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("insertLineBreak:"), sender)
}

// Inserts a newline character.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/insertNewline(_:)
func (o NSStandardKeyBindingRespondingObject) InsertNewline(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("insertNewline:"), sender)
}

// Inserts a newline character without invoking the field editor’s normal
// handling to end editing.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/insertNewlineIgnoringFieldEditor(_:)
func (o NSStandardKeyBindingRespondingObject) InsertNewlineIgnoringFieldEditor(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("insertNewlineIgnoringFieldEditor:"), sender)
}

// Inserts a paragraph separator.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/insertParagraphSeparator(_:)
func (o NSStandardKeyBindingRespondingObject) InsertParagraphSeparator(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("insertParagraphSeparator:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/insertSingleQuoteIgnoringSubstitution(_:)
func (o NSStandardKeyBindingRespondingObject) InsertSingleQuoteIgnoringSubstitution(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("insertSingleQuoteIgnoringSubstitution:"), sender)
}

// Inserts a tab character.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/insertTab(_:)
func (o NSStandardKeyBindingRespondingObject) InsertTab(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("insertTab:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/insertTabIgnoringFieldEditor(_:)
func (o NSStandardKeyBindingRespondingObject) InsertTabIgnoringFieldEditor(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("insertTabIgnoringFieldEditor:"), sender)
}

// Inserts the text you specify.
//
// insertString: The string to insert in the responder.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/insertText(_:)
func (o NSStandardKeyBindingRespondingObject) InsertText(insertString objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("insertText:"), insertString)
}

// Deletes content moving backward from the current insertion point.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/deleteBackward(_:)
func (o NSStandardKeyBindingRespondingObject) DeleteBackward(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("deleteBackward:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/deleteBackwardByDecomposingPreviousCharacter(_:)
func (o NSStandardKeyBindingRespondingObject) DeleteBackwardByDecomposingPreviousCharacter(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("deleteBackwardByDecomposingPreviousCharacter:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/deleteForward(_:)
func (o NSStandardKeyBindingRespondingObject) DeleteForward(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("deleteForward:"), sender)
}

// Deletes content from the insertion point to the beginning of the current
// line.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/deleteToBeginningOfLine(_:)
func (o NSStandardKeyBindingRespondingObject) DeleteToBeginningOfLine(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("deleteToBeginningOfLine:"), sender)
}

// Deletes content from the insertion point to the beginning of the current
// paragraph.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/deleteToBeginningOfParagraph(_:)
func (o NSStandardKeyBindingRespondingObject) DeleteToBeginningOfParagraph(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("deleteToBeginningOfParagraph:"), sender)
}

// Deletes content from the insertion point to the end of the current line.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/deleteToEndOfLine(_:)
func (o NSStandardKeyBindingRespondingObject) DeleteToEndOfLine(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("deleteToEndOfLine:"), sender)
}

// Deletes content from the insertion point to the end of the current
// paragraph.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/deleteToEndOfParagraph(_:)
func (o NSStandardKeyBindingRespondingObject) DeleteToEndOfParagraph(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("deleteToEndOfParagraph:"), sender)
}

// Deletes the word preceding the current insertion point.
//
// # Discussion
//
// If the insertion point is in the middle of a word, this method deletes only
// the portion of the word preceding the insertion point.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/deleteWordBackward(_:)
func (o NSStandardKeyBindingRespondingObject) DeleteWordBackward(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("deleteWordBackward:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/deleteWordForward(_:)
func (o NSStandardKeyBindingRespondingObject) DeleteWordForward(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("deleteWordForward:"), sender)
}

// Deletes the current selection, placing it in a temporary buffer, such as
// the Clipboard.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/yank(_:)
func (o NSStandardKeyBindingRespondingObject) Yank(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("yank:"), sender)
}

// Moves the insertion pointer backward in the current content.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveBackward(_:)
func (o NSStandardKeyBindingRespondingObject) MoveBackward(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("moveBackward:"), sender)
}

// Moves the insertion pointer down in the current content.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveDown(_:)
func (o NSStandardKeyBindingRespondingObject) MoveDown(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("moveDown:"), sender)
}

// Moves the insertion pointer forward in the current content.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveForward(_:)
func (o NSStandardKeyBindingRespondingObject) MoveForward(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("moveForward:"), sender)
}

// Moves the insertion pointer left in the current content.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveLeft(_:)
func (o NSStandardKeyBindingRespondingObject) MoveLeft(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("moveLeft:"), sender)
}

// Moves the insertion pointer right in the current content.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveRight(_:)
func (o NSStandardKeyBindingRespondingObject) MoveRight(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("moveRight:"), sender)
}

// Moves the insertion pointer up in the current content.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveUp(_:)
func (o NSStandardKeyBindingRespondingObject) MoveUp(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("moveUp:"), sender)
}

// Extends the selection to include the content before the current selection.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveBackwardAndModifySelection(_:)
func (o NSStandardKeyBindingRespondingObject) MoveBackwardAndModifySelection(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("moveBackwardAndModifySelection:"), sender)
}

// Extends the selection to include the content below the current selection.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveDownAndModifySelection(_:)
func (o NSStandardKeyBindingRespondingObject) MoveDownAndModifySelection(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("moveDownAndModifySelection:"), sender)
}

// Extends the selection to include the content after the current selection.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveForwardAndModifySelection(_:)
func (o NSStandardKeyBindingRespondingObject) MoveForwardAndModifySelection(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("moveForwardAndModifySelection:"), sender)
}

// Extends the selection to include the content to the left of the current
// selection.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveLeftAndModifySelection(_:)
func (o NSStandardKeyBindingRespondingObject) MoveLeftAndModifySelection(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("moveLeftAndModifySelection:"), sender)
}

// Extends the selection to include the content to the right of the current
// selection.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveRightAndModifySelection(_:)
func (o NSStandardKeyBindingRespondingObject) MoveRightAndModifySelection(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("moveRightAndModifySelection:"), sender)
}

// Extends the selection to include the content above the current selection.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveUpAndModifySelection(_:)
func (o NSStandardKeyBindingRespondingObject) MoveUpAndModifySelection(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("moveUpAndModifySelection:"), sender)
}

// Scrolls the content down by a page.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/scrollPageDown(_:)
func (o NSStandardKeyBindingRespondingObject) ScrollPageDown(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("scrollPageDown:"), sender)
}

// Scrolls the content up by a page.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/scrollPageUp(_:)
func (o NSStandardKeyBindingRespondingObject) ScrollPageUp(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("scrollPageUp:"), sender)
}

// Scrolls the content down by a line.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/scrollLineDown(_:)
func (o NSStandardKeyBindingRespondingObject) ScrollLineDown(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("scrollLineDown:"), sender)
}

// Scrolls the content up by a line.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/scrollLineUp(_:)
func (o NSStandardKeyBindingRespondingObject) ScrollLineUp(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("scrollLineUp:"), sender)
}

// Scrolls the content to the beginning of the document.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/scrollToBeginningOfDocument(_:)
func (o NSStandardKeyBindingRespondingObject) ScrollToBeginningOfDocument(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("scrollToBeginningOfDocument:"), sender)
}

// Scrolls the content to the end of the document.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/scrollToEndOfDocument(_:)
func (o NSStandardKeyBindingRespondingObject) ScrollToEndOfDocument(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("scrollToEndOfDocument:"), sender)
}

// Moves the visible content region down by a page.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/pageDown(_:)
func (o NSStandardKeyBindingRespondingObject) PageDown(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("pageDown:"), sender)
}

// Moves the visible content region up by a page.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/pageUp(_:)
func (o NSStandardKeyBindingRespondingObject) PageUp(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("pageUp:"), sender)
}

// Moves the visible content region down by a page, and extends the current
// selection.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/pageDownAndModifySelection(_:)
func (o NSStandardKeyBindingRespondingObject) PageDownAndModifySelection(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("pageDownAndModifySelection:"), sender)
}

// Moves the visible content region up by a page, and extends the current
// selection.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/pageUpAndModifySelection(_:)
func (o NSStandardKeyBindingRespondingObject) PageUpAndModifySelection(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("pageUpAndModifySelection:"), sender)
}

// Moves the visible content region so the current selection is visually
// centered.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/centerSelectionInVisibleArea(_:)
func (o NSStandardKeyBindingRespondingObject) CenterSelectionInVisibleArea(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("centerSelectionInVisibleArea:"), sender)
}

// Transposes the content around the current selection.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/transpose(_:)
func (o NSStandardKeyBindingRespondingObject) Transpose(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("transpose:"), sender)
}

// Transposes the words around the current selection.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/transposeWords(_:)
func (o NSStandardKeyBindingRespondingObject) TransposeWords(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("transposeWords:"), sender)
}

// Indents the content at the current selection.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/indent(_:)
func (o NSStandardKeyBindingRespondingObject) Indent(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("indent:"), sender)
}

// Cancels the current operation.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/cancelOperation(_:)
func (o NSStandardKeyBindingRespondingObject) CancelOperation(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("cancelOperation:"), sender)
}

// Invokes QuickLook to preview the current selection.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/quickLookPreviewItems(_:)
func (o NSStandardKeyBindingRespondingObject) QuickLookPreviewItems(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("quickLookPreviewItems:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/makeBaseWritingDirectionLeftToRight(_:)
func (o NSStandardKeyBindingRespondingObject) MakeBaseWritingDirectionLeftToRight(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("makeBaseWritingDirectionLeftToRight:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/makeBaseWritingDirectionNatural(_:)
func (o NSStandardKeyBindingRespondingObject) MakeBaseWritingDirectionNatural(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("makeBaseWritingDirectionNatural:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/makeBaseWritingDirectionRightToLeft(_:)
func (o NSStandardKeyBindingRespondingObject) MakeBaseWritingDirectionRightToLeft(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("makeBaseWritingDirectionRightToLeft:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/makeTextWritingDirectionLeftToRight(_:)
func (o NSStandardKeyBindingRespondingObject) MakeTextWritingDirectionLeftToRight(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("makeTextWritingDirectionLeftToRight:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/makeTextWritingDirectionNatural(_:)
func (o NSStandardKeyBindingRespondingObject) MakeTextWritingDirectionNatural(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("makeTextWritingDirectionNatural:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/makeTextWritingDirectionRightToLeft(_:)
func (o NSStandardKeyBindingRespondingObject) MakeTextWritingDirectionRightToLeft(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("makeTextWritingDirectionRightToLeft:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/capitalizeWord(_:)
func (o NSStandardKeyBindingRespondingObject) CapitalizeWord(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("capitalizeWord:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/changeCaseOfLetter(_:)
func (o NSStandardKeyBindingRespondingObject) ChangeCaseOfLetter(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("changeCaseOfLetter:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/lowercaseWord(_:)
func (o NSStandardKeyBindingRespondingObject) LowercaseWord(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("lowercaseWord:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/uppercaseWord(_:)
func (o NSStandardKeyBindingRespondingObject) UppercaseWord(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("uppercaseWord:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveToBeginningOfDocument(_:)
func (o NSStandardKeyBindingRespondingObject) MoveToBeginningOfDocument(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("moveToBeginningOfDocument:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveToBeginningOfDocumentAndModifySelection(_:)
func (o NSStandardKeyBindingRespondingObject) MoveToBeginningOfDocumentAndModifySelection(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("moveToBeginningOfDocumentAndModifySelection:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveToEndOfDocument(_:)
func (o NSStandardKeyBindingRespondingObject) MoveToEndOfDocument(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("moveToEndOfDocument:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveToEndOfDocumentAndModifySelection(_:)
func (o NSStandardKeyBindingRespondingObject) MoveToEndOfDocumentAndModifySelection(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("moveToEndOfDocumentAndModifySelection:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveParagraphBackwardAndModifySelection(_:)
func (o NSStandardKeyBindingRespondingObject) MoveParagraphBackwardAndModifySelection(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("moveParagraphBackwardAndModifySelection:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveParagraphForwardAndModifySelection(_:)
func (o NSStandardKeyBindingRespondingObject) MoveParagraphForwardAndModifySelection(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("moveParagraphForwardAndModifySelection:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveToBeginningOfParagraph(_:)
func (o NSStandardKeyBindingRespondingObject) MoveToBeginningOfParagraph(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("moveToBeginningOfParagraph:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveToBeginningOfParagraphAndModifySelection(_:)
func (o NSStandardKeyBindingRespondingObject) MoveToBeginningOfParagraphAndModifySelection(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("moveToBeginningOfParagraphAndModifySelection:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveToEndOfParagraph(_:)
func (o NSStandardKeyBindingRespondingObject) MoveToEndOfParagraph(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("moveToEndOfParagraph:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveToEndOfParagraphAndModifySelection(_:)
func (o NSStandardKeyBindingRespondingObject) MoveToEndOfParagraphAndModifySelection(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("moveToEndOfParagraphAndModifySelection:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveToBeginningOfLine(_:)
func (o NSStandardKeyBindingRespondingObject) MoveToBeginningOfLine(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("moveToBeginningOfLine:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveToBeginningOfLineAndModifySelection(_:)
func (o NSStandardKeyBindingRespondingObject) MoveToBeginningOfLineAndModifySelection(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("moveToBeginningOfLineAndModifySelection:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveToEndOfLine(_:)
func (o NSStandardKeyBindingRespondingObject) MoveToEndOfLine(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("moveToEndOfLine:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveToEndOfLineAndModifySelection(_:)
func (o NSStandardKeyBindingRespondingObject) MoveToEndOfLineAndModifySelection(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("moveToEndOfLineAndModifySelection:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveToLeftEndOfLine(_:)
func (o NSStandardKeyBindingRespondingObject) MoveToLeftEndOfLine(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("moveToLeftEndOfLine:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveToLeftEndOfLineAndModifySelection(_:)
func (o NSStandardKeyBindingRespondingObject) MoveToLeftEndOfLineAndModifySelection(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("moveToLeftEndOfLineAndModifySelection:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveToRightEndOfLine(_:)
func (o NSStandardKeyBindingRespondingObject) MoveToRightEndOfLine(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("moveToRightEndOfLine:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveToRightEndOfLineAndModifySelection(_:)
func (o NSStandardKeyBindingRespondingObject) MoveToRightEndOfLineAndModifySelection(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("moveToRightEndOfLineAndModifySelection:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/selectAll(_:)
func (o NSStandardKeyBindingRespondingObject) SelectAll(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("selectAll:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/selectLine(_:)
func (o NSStandardKeyBindingRespondingObject) SelectLine(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("selectLine:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/selectParagraph(_:)
func (o NSStandardKeyBindingRespondingObject) SelectParagraph(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("selectParagraph:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/selectWord(_:)
func (o NSStandardKeyBindingRespondingObject) SelectWord(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("selectWord:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/setMark(_:)
func (o NSStandardKeyBindingRespondingObject) SetMark(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setMark:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/selectToMark(_:)
func (o NSStandardKeyBindingRespondingObject) SelectToMark(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("selectToMark:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/deleteToMark(_:)
func (o NSStandardKeyBindingRespondingObject) DeleteToMark(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("deleteToMark:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/swapWithMark(_:)
func (o NSStandardKeyBindingRespondingObject) SwapWithMark(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("swapWithMark:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/complete(_:)
func (o NSStandardKeyBindingRespondingObject) Complete(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("complete:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveWordBackward(_:)
func (o NSStandardKeyBindingRespondingObject) MoveWordBackward(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("moveWordBackward:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveWordBackwardAndModifySelection(_:)
func (o NSStandardKeyBindingRespondingObject) MoveWordBackwardAndModifySelection(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("moveWordBackwardAndModifySelection:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveWordForward(_:)
func (o NSStandardKeyBindingRespondingObject) MoveWordForward(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("moveWordForward:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveWordForwardAndModifySelection(_:)
func (o NSStandardKeyBindingRespondingObject) MoveWordForwardAndModifySelection(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("moveWordForwardAndModifySelection:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveWordLeft(_:)
func (o NSStandardKeyBindingRespondingObject) MoveWordLeft(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("moveWordLeft:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveWordLeftAndModifySelection(_:)
func (o NSStandardKeyBindingRespondingObject) MoveWordLeftAndModifySelection(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("moveWordLeftAndModifySelection:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveWordRight(_:)
func (o NSStandardKeyBindingRespondingObject) MoveWordRight(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("moveWordRight:"), sender)
}

// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/moveWordRightAndModifySelection(_:)
func (o NSStandardKeyBindingRespondingObject) MoveWordRightAndModifySelection(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("moveWordRightAndModifySelection:"), sender)
}

// sender: The object that originated the display of the context menu.
//
// # Discussion
//
// Present a context menu at the text cursor position, selection, or wherever
// is appropriate for your responder.
//
// NSView has a default implementation of this method. For any view that
// returns a non-nil value from `-`, the default implementation will display
// that menu automatically. The NSView implementation uses the
// `selectionAnchorRect` method in the [NSViewContentSelectionInfo] protocol
// to determine the location of the selection and of the menu. The NSView
// implementation determines the menu to display by calling “ with a
// right-mouse-down event that is centered on the anchor rect. If
// `selectionAnchorRect` is not implemented, then the NSView implementation
// calls `menuForEvent` with a right-mouse-down event that is centered on the
// view’s bounds, and also displays the menu at that location.
//
// You should only override this method in a custom view if you need full
// control over the display of a context menu from the keyboard or
// Accessibility, beyond what is provided by default by NSView.
//
// If the view does not support a context menu, then you should call `[[self
// nextResponder] _cmd sender]` to pass the request up the responder chain.
//
// See: https://developer.apple.com/documentation/AppKit/NSStandardKeyBindingResponding/showContextMenuForSelection(_:)
func (o NSStandardKeyBindingRespondingObject) ShowContextMenuForSelection(sender objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("showContextMenuForSelection:"), sender)
}
