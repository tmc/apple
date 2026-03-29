// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSTextView] class.
var (
	_NSTextViewClass     NSTextViewClass
	_NSTextViewClassOnce sync.Once
)

func getNSTextViewClass() NSTextViewClass {
	_NSTextViewClassOnce.Do(func() {
		_NSTextViewClass = NSTextViewClass{class: objc.GetClass("NSTextView")}
	})
	return _NSTextViewClass
}

// GetNSTextViewClass returns the class object for NSTextView.
func GetNSTextViewClass() NSTextViewClass {
	return getNSTextViewClass()
}

type NSTextViewClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSTextViewClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTextViewClass) Alloc() NSTextView {
	rv := objc.Send[NSTextView](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A view that draws text and handles user interactions with that text.
//
// # Overview
// 
// The [NSTextView] class is the front-end class to the AppKit text system.
// The class draws the text managed by the back-end components and handles
// user events to select and modify its text, in addition to supporting rich
// text, attachments, input management, and key binding, and marked text
// attributes.
// 
// [NSTextView] is the principal means to obtain a text object that caters to
// almost all needs for displaying and managing text at the user interface
// level. While [NSTextView] is a subclass of the [NSText] class — which
// declares the most general Cocoa interface to the text system —
// [NSTextView] adds major features beyond the capabilities of [NSText]. You
// can also do more powerful and more creative text manipulation (such as
// displaying text in a circle) using [NSTextStorage], [NSTextLayoutManager],
// [NSTextContainer], and related classes.
// 
// You’re more likely to use the [NSTextView] class than [NSText]. It’s
// also important to remember that [NSTextView] conforms to a large number of
// protocols, the methods of which are available to instances of the
// [NSTextView] class.
// 
// [NSTextView] communicates with its delegate through methods declared both
// by the [NSTextViewDelegate] and by its superclass’s protocol,
// [NSTextDelegate]. All delegation messages come from the first text view.
// 
// In macOS 12 and later, if you explicitly call the `layoutManager` property
// on a text view or text container, the framework reverts to a compatibility
// mode that uses [NSLayoutManager]. The text view also switches to this
// compatibility mode when it encounters text content that’s not yet
// supported, such as [NSTextTable].
// 
// # About Delegate Methods
// 
// The [NSTextView] class communicates with its delegate through methods
// declared both by the [NSTextViewDelegate] and by its superclass’s
// protocol, [NSTextDelegate]. All delegation messages come from the first
// text view.
// 
// # Becoming the first responder
// 
// When the system invokes [BecomeFirstResponder] on a text view, if the
// previous first responder was not a text view on the same layout manager as
// the receiving text view, the receiving text view draws the selection and
// updates the insertion point if necessary.
// 
// To make a text view the first responder, call the containing window’s
// [FirstResponder] method. Never invoke a text view’s
// [BecomeFirstResponder] method directly.
// 
// # Resigning as first responder
// 
// When the system invokes [ResignFirstResponder] on a text view, if the
// object that will become the new first responder is a text view attached to
// the same layout manager as the receiver, the receiving text view returns
// [true] with no further action. Otherwise, it sends a [TextShouldEndEditing]
// message to its delegate (if any). If the delegate returns [false], the text
// view returns [false]. If the delegate returns [true], the text view hides
// the selection highlighting and posts an [didEndEditingNotification] to the
// default notification center and then returns [true].
//
// [didEndEditingNotification]: https://developer.apple.com/documentation/AppKit/NSText/didEndEditingNotification
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Creating a text view
//
//   - [NSTextView.InitWithFrameTextContainer]: Initializes a text view.
//   - [NSTextView.InitUsingTextLayoutManager]
//
// # Accessing text system objects
//
//   - [NSTextView.TextContainer]: The receiver’s text container.
//   - [NSTextView.SetTextContainer]
//   - [NSTextView.ReplaceTextContainer]: Replaces the text container for the group of text system objects containing the receiver, keeping the association between the receiver and its layout manager intact.
//   - [NSTextView.TextContainerInset]: The empty space the receiver leaves around its associated text container.
//   - [NSTextView.SetTextContainerInset]
//   - [NSTextView.TextContainerOrigin]: The origin of the receiver’s text container.
//   - [NSTextView.InvalidateTextContainerOrigin]: Invalidates the calculated origin of the text container.
//   - [NSTextView.TextLayoutManager]: The manager that lays out text for the receiver’s text container.
//   - [NSTextView.LayoutManager]: The layout manager that lays out text for the receiver’s text container.
//   - [NSTextView.TextContentStorage]: The receiver’s text storage object.
//   - [NSTextView.TextStorage]: The receiver’s text storage object.
//
// # Setting graphics attributes
//
//   - [NSTextView.AllowsDocumentBackgroundColorChange]: A Boolean value that indicates whether the receiver allows its background color to change.
//   - [NSTextView.SetAllowsDocumentBackgroundColorChange]
//   - [NSTextView.ChangeDocumentBackgroundColor]: An action method used to set the background color.
//
// # Controlling text display
//
//   - [NSTextView.SetNeedsDisplayInRectAvoidAdditionalLayout]: Marks the receiver as requiring display.
//   - [NSTextView.ShouldDrawInsertionPoint]: A Boolean value that determines whether the receiver should draw its insertion point.
//   - [NSTextView.DrawInsertionPointInRectColorTurnedOn]: Draws or erases the insertion point.
//   - [NSTextView.DrawViewBackgroundInRect]: Draws the background of the text view.
//   - [NSTextView.SetConstrainedFrameSize]: Attempts to set the frame size as if by user action.
//   - [NSTextView.CleanUpAfterDragOperation]: Releases the drag information still existing after the dragging session has completed.
//   - [NSTextView.ShowFindIndicatorForRange]: Causes a temporary highlighting effect to appear around the visible portion (or portions) of the specified range.
//
// # Inserting text
//
//   - [NSTextView.AllowedInputSourceLocales]: An array of locale identifiers representing input sources that are allowed to be enabled when the receiver has the keyboard focus.
//   - [NSTextView.SetAllowedInputSourceLocales]
//
// # Setting behavioral attributes
//
//   - [NSTextView.AllowsUndo]: A Boolean value that indicates whether the receiver allows undo.
//   - [NSTextView.SetAllowsUndo]
//   - [NSTextView.SetBaseWritingDirectionRange]: Sets the base writing direction of a range of text.
//   - [NSTextView.DefaultParagraphStyle]: The receiver’s default paragraph style.
//   - [NSTextView.SetDefaultParagraphStyle]
//   - [NSTextView.Outline]: Adds the outline attribute to the selected text attributes if absent; removes the attribute if present.
//   - [NSTextView.AllowsImageEditing]: Indicates whether image attachments should permit editing of their images.
//   - [NSTextView.SetAllowsImageEditing]
//   - [NSTextView.AutomaticQuoteSubstitutionEnabled]: A Boolean value that enables and disables automatic quotation mark substitution.
//   - [NSTextView.SetAutomaticQuoteSubstitutionEnabled]
//   - [NSTextView.ToggleAutomaticQuoteSubstitution]: Changes the state of automatic quotation mark substitution from enabled to disabled and vice versa.
//   - [NSTextView.AutomaticLinkDetectionEnabled]: A Boolean value that enables or disables automatic link detection.
//   - [NSTextView.SetAutomaticLinkDetectionEnabled]
//   - [NSTextView.ToggleAutomaticLinkDetection]: Changes the state of automatic link detection from enabled to disabled and vice versa.
//   - [NSTextView.DisplaysLinkToolTips]: A Boolean value that indicates whether the text view automatically supplies the destination of a link as a tooltip for text that has a link attribute.
//   - [NSTextView.SetDisplaysLinkToolTips]
//   - [NSTextView.AutomaticTextCompletionEnabled]: A Boolean value that indicates whether the text view supplies autocompletion suggestions as the user types.
//   - [NSTextView.SetAutomaticTextCompletionEnabled]
//   - [NSTextView.ToggleAutomaticTextCompletion]
//   - [NSTextView.UsesAdaptiveColorMappingForDarkAppearance]: A Boolean value that indicates whether the framework should use adaptive color mapping for dark appearance.
//   - [NSTextView.SetUsesAdaptiveColorMappingForDarkAppearance]
//   - [NSTextView.UsesRolloverButtonForSelection]
//   - [NSTextView.SetUsesRolloverButtonForSelection]
//
// # Using text formatting controls
//
//   - [NSTextView.UsesRuler]: A Boolean value that controls whether the text views sharing the receiver’s layout manager use a ruler.
//   - [NSTextView.SetUsesRuler]
//   - [NSTextView.UsesInspectorBar]: A Boolean value that indicates whether this text view uses the inspector bar.
//   - [NSTextView.SetUsesInspectorBar]
//
// # Managing the selection
//
//   - [NSTextView.SelectedRanges]: An array containing the ranges of characters selected in the receiver’s layout manager.
//   - [NSTextView.SetSelectedRanges]
//   - [NSTextView.SetSelectedRangeAffinityStillSelecting]: Sets the selection to a range of characters in response to user action.
//   - [NSTextView.SetSelectedRangesAffinityStillSelecting]: Sets the selection to the characters in an array of ranges in response to user action.
//   - [NSTextView.SelectionAffinity]: The preferred direction of selection.
//   - [NSTextView.SelectionGranularity]: The selection granularity for subsequent extension of a selection.
//   - [NSTextView.SetSelectionGranularity]
//   - [NSTextView.InsertionPointColor]: The color of the insertion point.
//   - [NSTextView.SetInsertionPointColor]
//   - [NSTextView.UpdateInsertionPointStateAndRestartTimer]: Updates the insertion point’s location and optionally restarts the blinking cursor timer.
//   - [NSTextView.SelectedTextAttributes]: The attributes used to indicate the selection.
//   - [NSTextView.SetSelectedTextAttributes]
//   - [NSTextView.MarkedTextAttributes]: The attributes used to draw marked text.
//   - [NSTextView.SetMarkedTextAttributes]
//   - [NSTextView.LinkTextAttributes]: The attributes used to draw the onscreen presentation of link text.
//   - [NSTextView.SetLinkTextAttributes]
//   - [NSTextView.CharacterIndexForInsertionAtPoint]: Returns a character index appropriate for placing a zero-length selection for an insertion point associated with the mouse at the given point.
//   - [NSTextView.UpdateCandidates]
//
// # Managing the pasteboard
//
//   - [NSTextView.PreferredPasteboardTypeFromArrayRestrictedToTypesFromArray]: Returns whatever type on the pasteboard would be most preferred for copying data.
//   - [NSTextView.ReadSelectionFromPasteboard]: Reads the text view’s preferred type of data from the specified pasteboard.
//   - [NSTextView.ReadSelectionFromPasteboardType]: Reads data of the given type from the specified pasteboard.
//   - [NSTextView.ReadablePasteboardTypes]: The types this text view can read immediately from the pasteboard.
//   - [NSTextView.WritablePasteboardTypes]: The pasteboard types that can be provided from the current selection.
//   - [NSTextView.WriteSelectionToPasteboardType]: Writes the current selection to the specified pasteboard using the given type.
//   - [NSTextView.WriteSelectionToPasteboardTypes]: Writes the current selection to the specified pasteboard under each given type.
//
// # Setting text attributes
//
//   - [NSTextView.AlignJustified]: Applies full justification to selected paragraphs (or all text, if the receiver is a plain text object).
//   - [NSTextView.ChangeAttributes]: Changes the attributes of the current selection.
//   - [NSTextView.ChangeColor]: Sets the color of the selected text.
//   - [NSTextView.SetAlignmentRange]: Sets the alignment of the paragraphs containing characters in the specified range.
//   - [NSTextView.TypingAttributes]: The receiver’s typing attributes.
//   - [NSTextView.SetTypingAttributes]
//   - [NSTextView.UseStandardKerning]: Set the receiver to use pair kerning data for the glyphs in its selection, or for all glyphs if the receiver is a plain text view.
//   - [NSTextView.LowerBaseline]: Lowers the baseline offset of selected text by 1 point, or of all text if the receiver is a plain text view.
//   - [NSTextView.RaiseBaseline]: Raises the baseline offset of selected text by 1 point, or of all text if the receiver is a plain text view.
//   - [NSTextView.TurnOffKerning]: Sets the receiver to use nominal glyph spacing for the glyphs in its selection, or for all glyphs if the receiver is a plain text view.
//   - [NSTextView.LoosenKerning]: Increases the space between glyphs in the receiver’s selection, or in all text if the receiver is a plain text view.
//   - [NSTextView.TightenKerning]: Decreases the space between glyphs in the receiver’s selection, or for all glyphs if the receiver is a plain text view.
//   - [NSTextView.UseStandardLigatures]: Sets the receiver to use the standard ligatures available for the fonts and languages used when setting text, for the glyphs in the selection if the receiver is a rich text view, or for all glyphs if it’s a plain text view.
//   - [NSTextView.TurnOffLigatures]: Sets the receiver to use only required ligatures when setting text, for the glyphs in the selection if the receiver is a rich text view, or for all glyphs if it’s a plain text view.
//   - [NSTextView.UseAllLigatures]: Sets the receiver to use all ligatures available for the fonts and languages used when setting text, for the glyphs in the selection if the receiver is a rich text view, or for all glyphs if it’s a plain text view.
//
// # Clicking and pasting
//
//   - [NSTextView.ClickedOnLinkAtIndex]: Causes the text view to act as if the user clicked on some text with the given link as the value of a link attribute associated with the text.
//   - [NSTextView.PasteAsPlainText]: Inserts the contents of the pasteboard into the receiver’s text as plain text.
//   - [NSTextView.PasteAsRichText]: This action method inserts the contents of the pasteboard into the receiver’s text as rich text, maintaining its attributes.
//
// # Supporting undo
//
//   - [NSTextView.BreakUndoCoalescing]: Informs the receiver that it should begin coalescing successive typing operations in a new undo grouping.
//   - [NSTextView.CoalescingUndo]: A Boolean value that indicates whether undo coalescing is in progress.
//
// # Customizing subclass behaviors
//
//   - [NSTextView.UpdateFontPanel]: Updates the Font panel to contain the font attributes of the selection.
//   - [NSTextView.UpdateRuler]: Updates the ruler view in the receiver’s enclosing scroll view to reflect the selection’s paragraph and marker attributes.
//   - [NSTextView.AcceptableDragTypes]: The data types that the receiver accepts as the destination view of a dragging operation.
//   - [NSTextView.UpdateDragTypeRegistration]: Updates the acceptable drag types of all text views associated with the receiver’s layout manager.
//   - [NSTextView.SelectionRangeForProposedRangeGranularity]: Returns an adjusted selected range based on the selection granularity.
//   - [NSTextView.RangeForUserCharacterAttributeChange]: The range of characters affected by an action method that changes character (not paragraph) attributes.
//   - [NSTextView.RangesForUserCharacterAttributeChange]: An array containing the ranges of characters affected by an action method that changes character (not paragraph) attributes.
//   - [NSTextView.RangeForUserParagraphAttributeChange]: The range of characters affected by an action method that changes paragraph (not character) attributes.
//   - [NSTextView.RangesForUserParagraphAttributeChange]: An array containing the ranges of characters affected by a method that changes paragraph (not character) attributes.
//   - [NSTextView.RangeForUserTextChange]: The range of characters affected by a method that changes characters (as opposed to attributes).
//   - [NSTextView.RangesForUserTextChange]: An array containing the ranges of characters affected by a method that changes characters (as opposed to attributes).
//   - [NSTextView.ShouldChangeTextInRangeReplacementString]: Initiates a series of delegate messages (and general notifications) to determine whether modifications can be made to the characters and attributes of the receiver’s text.
//   - [NSTextView.ShouldChangeTextInRangesReplacementStrings]: Initiates a series of delegate messages (and general notifications) to determine whether modifications can be made to the characters and attributes of the receiver’s text.
//   - [NSTextView.DidChangeText]: Sends out necessary notifications when a text change completes.
//   - [NSTextView.SmartInsertDeleteEnabled]: A Boolean value that controls whether the receiver inserts or deletes space around selected words so as to preserve proper spacing and punctuation.
//   - [NSTextView.SetSmartInsertDeleteEnabled]
//   - [NSTextView.SmartDeleteRangeForProposedRange]: Returns an extended range that includes adjacent whitespace that should be deleted along with the proposed range in order to preserve proper spacing and punctuation.
//   - [NSTextView.SmartInsertAfterStringForStringReplacingRange]: Returns any whitespace that needs to be added after the string to preserve proper spacing and punctuation when the string replaces the characters in the specified range.
//   - [NSTextView.SmartInsertBeforeStringForStringReplacingRange]: Returns any whitespace that needs to be added before the string to preserve proper spacing and punctuation when the string replaces the characters in the specified range.
//   - [NSTextView.SmartInsertForStringReplacingRangeBeforeStringAfterString]: Determines whether whitespace needs to be added around the string to preserve proper spacing and punctuation when it replaces the characters in the specified range.
//   - [NSTextView.ToggleSmartInsertDelete]: Changes the state of smart insert and delete from enabled to disabled and vice versa.
//
// # Working with the spelling checker
//
//   - [NSTextView.ContinuousSpellCheckingEnabled]: A Boolean value that indicates whether the receiver has continuous spell checking enabled.
//   - [NSTextView.SetContinuousSpellCheckingEnabled]
//   - [NSTextView.SpellCheckerDocumentTag]: A tag identifying the text view’s text as a document for the spell checker server.
//   - [NSTextView.ToggleContinuousSpellChecking]: Toggles whether continuous spell checking is enabled for the receiver.
//   - [NSTextView.GrammarCheckingEnabled]: Enables and disables grammar checking.
//   - [NSTextView.SetGrammarCheckingEnabled]
//   - [NSTextView.ToggleGrammarChecking]: Changes the state of grammar checking from enabled to disabled and vice versa.
//   - [NSTextView.SetSpellingStateRange]: Sets the spelling state, which controls the display of the spelling and grammar indicators on the given text range.
//
// # Working with the sharing service picker
//
//   - [NSTextView.OrderFrontSharingServicePicker]: Creates and displays a new instance of the sharing service picker.
//
// # Dragging
//
//   - [NSTextView.DragImageForSelectionWithEventOrigin]: Returns an appropriate drag image for the drag initiated by the specified event.
//   - [NSTextView.DragOperationForDraggingInfoType]: Returns the type of drag operation that should be performed if the image were released now.
//   - [NSTextView.DragSelectionWithEventOffsetSlideBack]: Begins dragging the current selected text range.
//   - [NSTextView.AcceptsGlyphInfo]: A Boolean value that indicates whether the receiver accepts the glyph info attribute.
//   - [NSTextView.SetAcceptsGlyphInfo]
//
// # Speaking text
//
//   - [NSTextView.StartSpeaking]: Speaks the selected text, or all text if no selection.
//   - [NSTextView.StopSpeaking]: Stops the speaking of text.
//
// # Working with panels
//
//   - [NSTextView.UsesFindPanel]: A Boolean value that indicates whether the receiver allows for a find panel.
//   - [NSTextView.SetUsesFindPanel]
//   - [NSTextView.PerformFindPanelAction]: Performs a find panel action specified by the sender’s tag.
//   - [NSTextView.OrderFrontLinkPanel]: Brings forward a panel allowing the user to manipulate links in the text view.
//   - [NSTextView.OrderFrontListPanel]: Brings forward a panel allowing the user to manipulate text lists in the text view.
//   - [NSTextView.OrderFrontSpacingPanel]: Brings forward a panel allowing the user to manipulate text line heights, interline spacing, and paragraph spacing, in the text view.
//   - [NSTextView.OrderFrontTablePanel]: Brings forward a panel allowing the user to manipulate text tables in the text view.
//   - [NSTextView.OrderFrontSubstitutionsPanel]: Brings forward a panel allowing the user to specify string substitutions in the text view.
//
// # Performing text completion
//
//   - [NSTextView.CompletionsForPartialWordRangeIndexOfSelectedItem]: Returns an array of potential completions, in the order to be presented, representing possible word completions available from a partial word.
//   - [NSTextView.InsertCompletionForPartialWordRangeMovementIsFinal]: Inserts the selected completion into the text at the appropriate location.
//   - [NSTextView.RangeForUserCompletion]: The partial range from the most recent beginning of a word up to the insertion point.
//
// # Checking and substituting text
//
//   - [NSTextView.CheckTextInDocument]: Performs the default text checking on the entire document.
//   - [NSTextView.CheckTextInSelection]: Performs the default text checking on the current selection.
//   - [NSTextView.CheckTextInRangeTypesOptions]: Check and replace the text in the range using the specified checking types and options.
//   - [NSTextView.HandleTextCheckingResultsForRangeTypesOptionsOrthographyWordCount]: Handles the text checking results returned by the text view
//   - [NSTextView.EnabledTextCheckingTypes]: The default text checking types.
//   - [NSTextView.SetEnabledTextCheckingTypes]
//   - [NSTextView.AutomaticDashSubstitutionEnabled]: A Boolean value that indicates whether automatic dash substitution is enabled.
//   - [NSTextView.SetAutomaticDashSubstitutionEnabled]
//   - [NSTextView.ToggleAutomaticDashSubstitution]: Toggles the state of the automatic dash substitution.
//   - [NSTextView.AutomaticDataDetectionEnabled]: A Boolean value that indicates whether automatic data detection is enabled.
//   - [NSTextView.SetAutomaticDataDetectionEnabled]
//   - [NSTextView.ToggleAutomaticDataDetection]: Toggles the state of the automatic data detection.
//   - [NSTextView.AutomaticSpellingCorrectionEnabled]: A Boolean value that indicates whether automatic spelling correction is enabled.
//   - [NSTextView.SetAutomaticSpellingCorrectionEnabled]
//   - [NSTextView.ToggleAutomaticSpellingCorrection]: Toggles the state of the automatic spelling correction.
//   - [NSTextView.AutomaticTextReplacementEnabled]: A Boolean value that indicates whether automatic text replacement is enabled.
//   - [NSTextView.SetAutomaticTextReplacementEnabled]
//   - [NSTextView.ToggleAutomaticTextReplacement]: Toggles the state of the automatic text replacement.
//   - [NSTextView.PerformValidatedReplacementInRangeWithAttributedString]: Replaces text in the range you specify with the attributed string you provide.
//
// # Getting the writing tools status
//
//   - [NSTextView.WritingToolsActive]
//
// # Supporting QuickLook
//
//   - [NSTextView.UpdateQuickLookPreviewPanel]: Notifies the QuickLook panel that an update may be required.
//   - [NSTextView.ToggleQuickLookPreviewPanel]: An action message that toggles the visibility state of the Quick Look preview panel.
//   - [NSTextView.QuickLookPreviewableItemsInRanges]: Returns an array of URLs for items that can be displayed by QuickLook in the specified ranges.
//
// # Changing layout orientation
//
//   - [NSTextView.ChangeLayoutOrientation]: An action method that sets the layout orientation of the text.
//
// # Using the Find Bar
//
//   - [NSTextView.UsesFindBar]: A Boolean value that indicates whether to use the find bar for this text view.
//   - [NSTextView.SetUsesFindBar]
//   - [NSTextView.IncrementalSearchingEnabled]: A Boolean value that indicates whether incremental searching is enabled.
//   - [NSTextView.SetIncrementalSearchingEnabled]
//
// # Interacting with the Touch Bar
//
//   - [NSTextView.AllowsCharacterPickerTouchBarItem]
//   - [NSTextView.SetAllowsCharacterPickerTouchBarItem]
//   - [NSTextView.UpdateTextTouchBarItems]
//   - [NSTextView.UpdateTouchBarItemIdentifiers]
//
// # Instance Properties
//
//   - [NSTextView.AllowedWritingToolsResultOptions]
//   - [NSTextView.SetAllowedWritingToolsResultOptions]
//   - [NSTextView.InlinePredictionType]
//   - [NSTextView.SetInlinePredictionType]
//   - [NSTextView.MathExpressionCompletionType]
//   - [NSTextView.SetMathExpressionCompletionType]
//   - [NSTextView.TextHighlightAttributes]: ************************* Text Highlight  support **************************
//   - [NSTextView.SetTextHighlightAttributes]
//   - [NSTextView.WritingToolsBehavior]
//   - [NSTextView.SetWritingToolsBehavior]
//
// # Instance Methods
//
//   - [NSTextView.DrawTextHighlightBackgroundForTextRangeOrigin]
//   - [NSTextView.Highlight]: An action for toggling [NSTextHighlightStyleAttributeName] in the receiver’s selected range. The sender should be a menu item with a `representedObject` of type ([NSTextHighlightColorScheme]).
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView
type NSTextView struct {
	NSText
}

// NSTextViewFromID constructs a [NSTextView] from an objc.ID.
//
// A view that draws text and handles user interactions with that text.
func NSTextViewFromID(id objc.ID) NSTextView {
	return NSTextView{NSText: NSTextFromID(id)}
}
// NOTE: NSTextView adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSTextView] class.
//
// # Creating a text view
//
//   - [INSTextView.InitWithFrameTextContainer]: Initializes a text view.
//   - [INSTextView.InitUsingTextLayoutManager]
//
// # Accessing text system objects
//
//   - [INSTextView.TextContainer]: The receiver’s text container.
//   - [INSTextView.SetTextContainer]
//   - [INSTextView.ReplaceTextContainer]: Replaces the text container for the group of text system objects containing the receiver, keeping the association between the receiver and its layout manager intact.
//   - [INSTextView.TextContainerInset]: The empty space the receiver leaves around its associated text container.
//   - [INSTextView.SetTextContainerInset]
//   - [INSTextView.TextContainerOrigin]: The origin of the receiver’s text container.
//   - [INSTextView.InvalidateTextContainerOrigin]: Invalidates the calculated origin of the text container.
//   - [INSTextView.TextLayoutManager]: The manager that lays out text for the receiver’s text container.
//   - [INSTextView.LayoutManager]: The layout manager that lays out text for the receiver’s text container.
//   - [INSTextView.TextContentStorage]: The receiver’s text storage object.
//   - [INSTextView.TextStorage]: The receiver’s text storage object.
//
// # Setting graphics attributes
//
//   - [INSTextView.AllowsDocumentBackgroundColorChange]: A Boolean value that indicates whether the receiver allows its background color to change.
//   - [INSTextView.SetAllowsDocumentBackgroundColorChange]
//   - [INSTextView.ChangeDocumentBackgroundColor]: An action method used to set the background color.
//
// # Controlling text display
//
//   - [INSTextView.SetNeedsDisplayInRectAvoidAdditionalLayout]: Marks the receiver as requiring display.
//   - [INSTextView.ShouldDrawInsertionPoint]: A Boolean value that determines whether the receiver should draw its insertion point.
//   - [INSTextView.DrawInsertionPointInRectColorTurnedOn]: Draws or erases the insertion point.
//   - [INSTextView.DrawViewBackgroundInRect]: Draws the background of the text view.
//   - [INSTextView.SetConstrainedFrameSize]: Attempts to set the frame size as if by user action.
//   - [INSTextView.CleanUpAfterDragOperation]: Releases the drag information still existing after the dragging session has completed.
//   - [INSTextView.ShowFindIndicatorForRange]: Causes a temporary highlighting effect to appear around the visible portion (or portions) of the specified range.
//
// # Inserting text
//
//   - [INSTextView.AllowedInputSourceLocales]: An array of locale identifiers representing input sources that are allowed to be enabled when the receiver has the keyboard focus.
//   - [INSTextView.SetAllowedInputSourceLocales]
//
// # Setting behavioral attributes
//
//   - [INSTextView.AllowsUndo]: A Boolean value that indicates whether the receiver allows undo.
//   - [INSTextView.SetAllowsUndo]
//   - [INSTextView.SetBaseWritingDirectionRange]: Sets the base writing direction of a range of text.
//   - [INSTextView.DefaultParagraphStyle]: The receiver’s default paragraph style.
//   - [INSTextView.SetDefaultParagraphStyle]
//   - [INSTextView.Outline]: Adds the outline attribute to the selected text attributes if absent; removes the attribute if present.
//   - [INSTextView.AllowsImageEditing]: Indicates whether image attachments should permit editing of their images.
//   - [INSTextView.SetAllowsImageEditing]
//   - [INSTextView.AutomaticQuoteSubstitutionEnabled]: A Boolean value that enables and disables automatic quotation mark substitution.
//   - [INSTextView.SetAutomaticQuoteSubstitutionEnabled]
//   - [INSTextView.ToggleAutomaticQuoteSubstitution]: Changes the state of automatic quotation mark substitution from enabled to disabled and vice versa.
//   - [INSTextView.AutomaticLinkDetectionEnabled]: A Boolean value that enables or disables automatic link detection.
//   - [INSTextView.SetAutomaticLinkDetectionEnabled]
//   - [INSTextView.ToggleAutomaticLinkDetection]: Changes the state of automatic link detection from enabled to disabled and vice versa.
//   - [INSTextView.DisplaysLinkToolTips]: A Boolean value that indicates whether the text view automatically supplies the destination of a link as a tooltip for text that has a link attribute.
//   - [INSTextView.SetDisplaysLinkToolTips]
//   - [INSTextView.AutomaticTextCompletionEnabled]: A Boolean value that indicates whether the text view supplies autocompletion suggestions as the user types.
//   - [INSTextView.SetAutomaticTextCompletionEnabled]
//   - [INSTextView.ToggleAutomaticTextCompletion]
//   - [INSTextView.UsesAdaptiveColorMappingForDarkAppearance]: A Boolean value that indicates whether the framework should use adaptive color mapping for dark appearance.
//   - [INSTextView.SetUsesAdaptiveColorMappingForDarkAppearance]
//   - [INSTextView.UsesRolloverButtonForSelection]
//   - [INSTextView.SetUsesRolloverButtonForSelection]
//
// # Using text formatting controls
//
//   - [INSTextView.UsesRuler]: A Boolean value that controls whether the text views sharing the receiver’s layout manager use a ruler.
//   - [INSTextView.SetUsesRuler]
//   - [INSTextView.UsesInspectorBar]: A Boolean value that indicates whether this text view uses the inspector bar.
//   - [INSTextView.SetUsesInspectorBar]
//
// # Managing the selection
//
//   - [INSTextView.SelectedRanges]: An array containing the ranges of characters selected in the receiver’s layout manager.
//   - [INSTextView.SetSelectedRanges]
//   - [INSTextView.SetSelectedRangeAffinityStillSelecting]: Sets the selection to a range of characters in response to user action.
//   - [INSTextView.SetSelectedRangesAffinityStillSelecting]: Sets the selection to the characters in an array of ranges in response to user action.
//   - [INSTextView.SelectionAffinity]: The preferred direction of selection.
//   - [INSTextView.SelectionGranularity]: The selection granularity for subsequent extension of a selection.
//   - [INSTextView.SetSelectionGranularity]
//   - [INSTextView.InsertionPointColor]: The color of the insertion point.
//   - [INSTextView.SetInsertionPointColor]
//   - [INSTextView.UpdateInsertionPointStateAndRestartTimer]: Updates the insertion point’s location and optionally restarts the blinking cursor timer.
//   - [INSTextView.SelectedTextAttributes]: The attributes used to indicate the selection.
//   - [INSTextView.SetSelectedTextAttributes]
//   - [INSTextView.MarkedTextAttributes]: The attributes used to draw marked text.
//   - [INSTextView.SetMarkedTextAttributes]
//   - [INSTextView.LinkTextAttributes]: The attributes used to draw the onscreen presentation of link text.
//   - [INSTextView.SetLinkTextAttributes]
//   - [INSTextView.CharacterIndexForInsertionAtPoint]: Returns a character index appropriate for placing a zero-length selection for an insertion point associated with the mouse at the given point.
//   - [INSTextView.UpdateCandidates]
//
// # Managing the pasteboard
//
//   - [INSTextView.PreferredPasteboardTypeFromArrayRestrictedToTypesFromArray]: Returns whatever type on the pasteboard would be most preferred for copying data.
//   - [INSTextView.ReadSelectionFromPasteboard]: Reads the text view’s preferred type of data from the specified pasteboard.
//   - [INSTextView.ReadSelectionFromPasteboardType]: Reads data of the given type from the specified pasteboard.
//   - [INSTextView.ReadablePasteboardTypes]: The types this text view can read immediately from the pasteboard.
//   - [INSTextView.WritablePasteboardTypes]: The pasteboard types that can be provided from the current selection.
//   - [INSTextView.WriteSelectionToPasteboardType]: Writes the current selection to the specified pasteboard using the given type.
//   - [INSTextView.WriteSelectionToPasteboardTypes]: Writes the current selection to the specified pasteboard under each given type.
//
// # Setting text attributes
//
//   - [INSTextView.AlignJustified]: Applies full justification to selected paragraphs (or all text, if the receiver is a plain text object).
//   - [INSTextView.ChangeAttributes]: Changes the attributes of the current selection.
//   - [INSTextView.ChangeColor]: Sets the color of the selected text.
//   - [INSTextView.SetAlignmentRange]: Sets the alignment of the paragraphs containing characters in the specified range.
//   - [INSTextView.TypingAttributes]: The receiver’s typing attributes.
//   - [INSTextView.SetTypingAttributes]
//   - [INSTextView.UseStandardKerning]: Set the receiver to use pair kerning data for the glyphs in its selection, or for all glyphs if the receiver is a plain text view.
//   - [INSTextView.LowerBaseline]: Lowers the baseline offset of selected text by 1 point, or of all text if the receiver is a plain text view.
//   - [INSTextView.RaiseBaseline]: Raises the baseline offset of selected text by 1 point, or of all text if the receiver is a plain text view.
//   - [INSTextView.TurnOffKerning]: Sets the receiver to use nominal glyph spacing for the glyphs in its selection, or for all glyphs if the receiver is a plain text view.
//   - [INSTextView.LoosenKerning]: Increases the space between glyphs in the receiver’s selection, or in all text if the receiver is a plain text view.
//   - [INSTextView.TightenKerning]: Decreases the space between glyphs in the receiver’s selection, or for all glyphs if the receiver is a plain text view.
//   - [INSTextView.UseStandardLigatures]: Sets the receiver to use the standard ligatures available for the fonts and languages used when setting text, for the glyphs in the selection if the receiver is a rich text view, or for all glyphs if it’s a plain text view.
//   - [INSTextView.TurnOffLigatures]: Sets the receiver to use only required ligatures when setting text, for the glyphs in the selection if the receiver is a rich text view, or for all glyphs if it’s a plain text view.
//   - [INSTextView.UseAllLigatures]: Sets the receiver to use all ligatures available for the fonts and languages used when setting text, for the glyphs in the selection if the receiver is a rich text view, or for all glyphs if it’s a plain text view.
//
// # Clicking and pasting
//
//   - [INSTextView.ClickedOnLinkAtIndex]: Causes the text view to act as if the user clicked on some text with the given link as the value of a link attribute associated with the text.
//   - [INSTextView.PasteAsPlainText]: Inserts the contents of the pasteboard into the receiver’s text as plain text.
//   - [INSTextView.PasteAsRichText]: This action method inserts the contents of the pasteboard into the receiver’s text as rich text, maintaining its attributes.
//
// # Supporting undo
//
//   - [INSTextView.BreakUndoCoalescing]: Informs the receiver that it should begin coalescing successive typing operations in a new undo grouping.
//   - [INSTextView.CoalescingUndo]: A Boolean value that indicates whether undo coalescing is in progress.
//
// # Customizing subclass behaviors
//
//   - [INSTextView.UpdateFontPanel]: Updates the Font panel to contain the font attributes of the selection.
//   - [INSTextView.UpdateRuler]: Updates the ruler view in the receiver’s enclosing scroll view to reflect the selection’s paragraph and marker attributes.
//   - [INSTextView.AcceptableDragTypes]: The data types that the receiver accepts as the destination view of a dragging operation.
//   - [INSTextView.UpdateDragTypeRegistration]: Updates the acceptable drag types of all text views associated with the receiver’s layout manager.
//   - [INSTextView.SelectionRangeForProposedRangeGranularity]: Returns an adjusted selected range based on the selection granularity.
//   - [INSTextView.RangeForUserCharacterAttributeChange]: The range of characters affected by an action method that changes character (not paragraph) attributes.
//   - [INSTextView.RangesForUserCharacterAttributeChange]: An array containing the ranges of characters affected by an action method that changes character (not paragraph) attributes.
//   - [INSTextView.RangeForUserParagraphAttributeChange]: The range of characters affected by an action method that changes paragraph (not character) attributes.
//   - [INSTextView.RangesForUserParagraphAttributeChange]: An array containing the ranges of characters affected by a method that changes paragraph (not character) attributes.
//   - [INSTextView.RangeForUserTextChange]: The range of characters affected by a method that changes characters (as opposed to attributes).
//   - [INSTextView.RangesForUserTextChange]: An array containing the ranges of characters affected by a method that changes characters (as opposed to attributes).
//   - [INSTextView.ShouldChangeTextInRangeReplacementString]: Initiates a series of delegate messages (and general notifications) to determine whether modifications can be made to the characters and attributes of the receiver’s text.
//   - [INSTextView.ShouldChangeTextInRangesReplacementStrings]: Initiates a series of delegate messages (and general notifications) to determine whether modifications can be made to the characters and attributes of the receiver’s text.
//   - [INSTextView.DidChangeText]: Sends out necessary notifications when a text change completes.
//   - [INSTextView.SmartInsertDeleteEnabled]: A Boolean value that controls whether the receiver inserts or deletes space around selected words so as to preserve proper spacing and punctuation.
//   - [INSTextView.SetSmartInsertDeleteEnabled]
//   - [INSTextView.SmartDeleteRangeForProposedRange]: Returns an extended range that includes adjacent whitespace that should be deleted along with the proposed range in order to preserve proper spacing and punctuation.
//   - [INSTextView.SmartInsertAfterStringForStringReplacingRange]: Returns any whitespace that needs to be added after the string to preserve proper spacing and punctuation when the string replaces the characters in the specified range.
//   - [INSTextView.SmartInsertBeforeStringForStringReplacingRange]: Returns any whitespace that needs to be added before the string to preserve proper spacing and punctuation when the string replaces the characters in the specified range.
//   - [INSTextView.SmartInsertForStringReplacingRangeBeforeStringAfterString]: Determines whether whitespace needs to be added around the string to preserve proper spacing and punctuation when it replaces the characters in the specified range.
//   - [INSTextView.ToggleSmartInsertDelete]: Changes the state of smart insert and delete from enabled to disabled and vice versa.
//
// # Working with the spelling checker
//
//   - [INSTextView.ContinuousSpellCheckingEnabled]: A Boolean value that indicates whether the receiver has continuous spell checking enabled.
//   - [INSTextView.SetContinuousSpellCheckingEnabled]
//   - [INSTextView.SpellCheckerDocumentTag]: A tag identifying the text view’s text as a document for the spell checker server.
//   - [INSTextView.ToggleContinuousSpellChecking]: Toggles whether continuous spell checking is enabled for the receiver.
//   - [INSTextView.GrammarCheckingEnabled]: Enables and disables grammar checking.
//   - [INSTextView.SetGrammarCheckingEnabled]
//   - [INSTextView.ToggleGrammarChecking]: Changes the state of grammar checking from enabled to disabled and vice versa.
//   - [INSTextView.SetSpellingStateRange]: Sets the spelling state, which controls the display of the spelling and grammar indicators on the given text range.
//
// # Working with the sharing service picker
//
//   - [INSTextView.OrderFrontSharingServicePicker]: Creates and displays a new instance of the sharing service picker.
//
// # Dragging
//
//   - [INSTextView.DragImageForSelectionWithEventOrigin]: Returns an appropriate drag image for the drag initiated by the specified event.
//   - [INSTextView.DragOperationForDraggingInfoType]: Returns the type of drag operation that should be performed if the image were released now.
//   - [INSTextView.DragSelectionWithEventOffsetSlideBack]: Begins dragging the current selected text range.
//   - [INSTextView.AcceptsGlyphInfo]: A Boolean value that indicates whether the receiver accepts the glyph info attribute.
//   - [INSTextView.SetAcceptsGlyphInfo]
//
// # Speaking text
//
//   - [INSTextView.StartSpeaking]: Speaks the selected text, or all text if no selection.
//   - [INSTextView.StopSpeaking]: Stops the speaking of text.
//
// # Working with panels
//
//   - [INSTextView.UsesFindPanel]: A Boolean value that indicates whether the receiver allows for a find panel.
//   - [INSTextView.SetUsesFindPanel]
//   - [INSTextView.PerformFindPanelAction]: Performs a find panel action specified by the sender’s tag.
//   - [INSTextView.OrderFrontLinkPanel]: Brings forward a panel allowing the user to manipulate links in the text view.
//   - [INSTextView.OrderFrontListPanel]: Brings forward a panel allowing the user to manipulate text lists in the text view.
//   - [INSTextView.OrderFrontSpacingPanel]: Brings forward a panel allowing the user to manipulate text line heights, interline spacing, and paragraph spacing, in the text view.
//   - [INSTextView.OrderFrontTablePanel]: Brings forward a panel allowing the user to manipulate text tables in the text view.
//   - [INSTextView.OrderFrontSubstitutionsPanel]: Brings forward a panel allowing the user to specify string substitutions in the text view.
//
// # Performing text completion
//
//   - [INSTextView.CompletionsForPartialWordRangeIndexOfSelectedItem]: Returns an array of potential completions, in the order to be presented, representing possible word completions available from a partial word.
//   - [INSTextView.InsertCompletionForPartialWordRangeMovementIsFinal]: Inserts the selected completion into the text at the appropriate location.
//   - [INSTextView.RangeForUserCompletion]: The partial range from the most recent beginning of a word up to the insertion point.
//
// # Checking and substituting text
//
//   - [INSTextView.CheckTextInDocument]: Performs the default text checking on the entire document.
//   - [INSTextView.CheckTextInSelection]: Performs the default text checking on the current selection.
//   - [INSTextView.CheckTextInRangeTypesOptions]: Check and replace the text in the range using the specified checking types and options.
//   - [INSTextView.HandleTextCheckingResultsForRangeTypesOptionsOrthographyWordCount]: Handles the text checking results returned by the text view
//   - [INSTextView.EnabledTextCheckingTypes]: The default text checking types.
//   - [INSTextView.SetEnabledTextCheckingTypes]
//   - [INSTextView.AutomaticDashSubstitutionEnabled]: A Boolean value that indicates whether automatic dash substitution is enabled.
//   - [INSTextView.SetAutomaticDashSubstitutionEnabled]
//   - [INSTextView.ToggleAutomaticDashSubstitution]: Toggles the state of the automatic dash substitution.
//   - [INSTextView.AutomaticDataDetectionEnabled]: A Boolean value that indicates whether automatic data detection is enabled.
//   - [INSTextView.SetAutomaticDataDetectionEnabled]
//   - [INSTextView.ToggleAutomaticDataDetection]: Toggles the state of the automatic data detection.
//   - [INSTextView.AutomaticSpellingCorrectionEnabled]: A Boolean value that indicates whether automatic spelling correction is enabled.
//   - [INSTextView.SetAutomaticSpellingCorrectionEnabled]
//   - [INSTextView.ToggleAutomaticSpellingCorrection]: Toggles the state of the automatic spelling correction.
//   - [INSTextView.AutomaticTextReplacementEnabled]: A Boolean value that indicates whether automatic text replacement is enabled.
//   - [INSTextView.SetAutomaticTextReplacementEnabled]
//   - [INSTextView.ToggleAutomaticTextReplacement]: Toggles the state of the automatic text replacement.
//   - [INSTextView.PerformValidatedReplacementInRangeWithAttributedString]: Replaces text in the range you specify with the attributed string you provide.
//
// # Getting the writing tools status
//
//   - [INSTextView.WritingToolsActive]
//
// # Supporting QuickLook
//
//   - [INSTextView.UpdateQuickLookPreviewPanel]: Notifies the QuickLook panel that an update may be required.
//   - [INSTextView.ToggleQuickLookPreviewPanel]: An action message that toggles the visibility state of the Quick Look preview panel.
//   - [INSTextView.QuickLookPreviewableItemsInRanges]: Returns an array of URLs for items that can be displayed by QuickLook in the specified ranges.
//
// # Changing layout orientation
//
//   - [INSTextView.ChangeLayoutOrientation]: An action method that sets the layout orientation of the text.
//
// # Using the Find Bar
//
//   - [INSTextView.UsesFindBar]: A Boolean value that indicates whether to use the find bar for this text view.
//   - [INSTextView.SetUsesFindBar]
//   - [INSTextView.IncrementalSearchingEnabled]: A Boolean value that indicates whether incremental searching is enabled.
//   - [INSTextView.SetIncrementalSearchingEnabled]
//
// # Interacting with the Touch Bar
//
//   - [INSTextView.AllowsCharacterPickerTouchBarItem]
//   - [INSTextView.SetAllowsCharacterPickerTouchBarItem]
//   - [INSTextView.UpdateTextTouchBarItems]
//   - [INSTextView.UpdateTouchBarItemIdentifiers]
//
// # Instance Properties
//
//   - [INSTextView.AllowedWritingToolsResultOptions]
//   - [INSTextView.SetAllowedWritingToolsResultOptions]
//   - [INSTextView.InlinePredictionType]
//   - [INSTextView.SetInlinePredictionType]
//   - [INSTextView.MathExpressionCompletionType]
//   - [INSTextView.SetMathExpressionCompletionType]
//   - [INSTextView.TextHighlightAttributes]: ************************* Text Highlight  support **************************
//   - [INSTextView.SetTextHighlightAttributes]
//   - [INSTextView.WritingToolsBehavior]
//   - [INSTextView.SetWritingToolsBehavior]
//
// # Instance Methods
//
//   - [INSTextView.DrawTextHighlightBackgroundForTextRangeOrigin]
//   - [INSTextView.Highlight]: An action for toggling [NSTextHighlightStyleAttributeName] in the receiver’s selected range. The sender should be a menu item with a `representedObject` of type ([NSTextHighlightColorScheme]).
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView
type INSTextView interface {
	INSText
	NSAccessibilityNavigableStaticText
	NSAccessibilityStaticText
	NSCandidateListTouchBarItemDelegate
	NSDraggingSource
	NSMenuItemValidation
	NSTextContent
	NSTextInput
	NSTextInputClient
	NSTextLayoutOrientationProvider
	NSTouchBarDelegate
	NSUserInterfaceValidations

	// Topic: Creating a text view

	// Initializes a text view.
	InitWithFrameTextContainer(frameRect corefoundation.CGRect, container INSTextContainer) NSTextView
	InitUsingTextLayoutManager(usingTextLayoutManager bool) NSTextView

	// Topic: Accessing text system objects

	// The receiver’s text container.
	TextContainer() INSTextContainer
	SetTextContainer(value INSTextContainer)
	// Replaces the text container for the group of text system objects containing the receiver, keeping the association between the receiver and its layout manager intact.
	ReplaceTextContainer(newContainer INSTextContainer)
	// The empty space the receiver leaves around its associated text container.
	TextContainerInset() corefoundation.CGSize
	SetTextContainerInset(value corefoundation.CGSize)
	// The origin of the receiver’s text container.
	TextContainerOrigin() corefoundation.CGPoint
	// Invalidates the calculated origin of the text container.
	InvalidateTextContainerOrigin()
	// The manager that lays out text for the receiver’s text container.
	TextLayoutManager() INSTextLayoutManager
	// The layout manager that lays out text for the receiver’s text container.
	LayoutManager() INSLayoutManager
	// The receiver’s text storage object.
	TextContentStorage() INSTextContentStorage
	// The receiver’s text storage object.
	TextStorage() NSTextStorage

	// Topic: Setting graphics attributes

	// A Boolean value that indicates whether the receiver allows its background color to change.
	AllowsDocumentBackgroundColorChange() bool
	SetAllowsDocumentBackgroundColorChange(value bool)
	// An action method used to set the background color.
	ChangeDocumentBackgroundColor(sender objectivec.IObject)

	// Topic: Controlling text display

	// Marks the receiver as requiring display.
	SetNeedsDisplayInRectAvoidAdditionalLayout(rect corefoundation.CGRect, flag bool)
	// A Boolean value that determines whether the receiver should draw its insertion point.
	ShouldDrawInsertionPoint() bool
	// Draws or erases the insertion point.
	DrawInsertionPointInRectColorTurnedOn(rect corefoundation.CGRect, color INSColor, flag bool)
	// Draws the background of the text view.
	DrawViewBackgroundInRect(rect corefoundation.CGRect)
	// Attempts to set the frame size as if by user action.
	SetConstrainedFrameSize(desiredSize corefoundation.CGSize)
	// Releases the drag information still existing after the dragging session has completed.
	CleanUpAfterDragOperation()
	// Causes a temporary highlighting effect to appear around the visible portion (or portions) of the specified range.
	ShowFindIndicatorForRange(charRange foundation.NSRange)

	// Topic: Inserting text

	// An array of locale identifiers representing input sources that are allowed to be enabled when the receiver has the keyboard focus.
	AllowedInputSourceLocales() []string
	SetAllowedInputSourceLocales(value []string)

	// Topic: Setting behavioral attributes

	// A Boolean value that indicates whether the receiver allows undo.
	AllowsUndo() bool
	SetAllowsUndo(value bool)
	// Sets the base writing direction of a range of text.
	SetBaseWritingDirectionRange(writingDirection NSWritingDirection, range_ foundation.NSRange)
	// The receiver’s default paragraph style.
	DefaultParagraphStyle() INSParagraphStyle
	SetDefaultParagraphStyle(value INSParagraphStyle)
	// Adds the outline attribute to the selected text attributes if absent; removes the attribute if present.
	Outline(sender objectivec.IObject)
	// Indicates whether image attachments should permit editing of their images.
	AllowsImageEditing() bool
	SetAllowsImageEditing(value bool)
	// A Boolean value that enables and disables automatic quotation mark substitution.
	AutomaticQuoteSubstitutionEnabled() bool
	SetAutomaticQuoteSubstitutionEnabled(value bool)
	// Changes the state of automatic quotation mark substitution from enabled to disabled and vice versa.
	ToggleAutomaticQuoteSubstitution(sender objectivec.IObject)
	// A Boolean value that enables or disables automatic link detection.
	AutomaticLinkDetectionEnabled() bool
	SetAutomaticLinkDetectionEnabled(value bool)
	// Changes the state of automatic link detection from enabled to disabled and vice versa.
	ToggleAutomaticLinkDetection(sender objectivec.IObject)
	// A Boolean value that indicates whether the text view automatically supplies the destination of a link as a tooltip for text that has a link attribute.
	DisplaysLinkToolTips() bool
	SetDisplaysLinkToolTips(value bool)
	// A Boolean value that indicates whether the text view supplies autocompletion suggestions as the user types.
	AutomaticTextCompletionEnabled() bool
	SetAutomaticTextCompletionEnabled(value bool)
	ToggleAutomaticTextCompletion(sender objectivec.IObject)
	// A Boolean value that indicates whether the framework should use adaptive color mapping for dark appearance.
	UsesAdaptiveColorMappingForDarkAppearance() bool
	SetUsesAdaptiveColorMappingForDarkAppearance(value bool)
	UsesRolloverButtonForSelection() bool
	SetUsesRolloverButtonForSelection(value bool)

	// Topic: Using text formatting controls

	// A Boolean value that controls whether the text views sharing the receiver’s layout manager use a ruler.
	UsesRuler() bool
	SetUsesRuler(value bool)
	// A Boolean value that indicates whether this text view uses the inspector bar.
	UsesInspectorBar() bool
	SetUsesInspectorBar(value bool)

	// Topic: Managing the selection

	// An array containing the ranges of characters selected in the receiver’s layout manager.
	SelectedRanges() []foundation.NSValue
	SetSelectedRanges(value []foundation.NSValue)
	// Sets the selection to a range of characters in response to user action.
	SetSelectedRangeAffinityStillSelecting(charRange foundation.NSRange, affinity NSSelectionAffinity, stillSelectingFlag bool)
	// Sets the selection to the characters in an array of ranges in response to user action.
	SetSelectedRangesAffinityStillSelecting(ranges []foundation.NSValue, affinity NSSelectionAffinity, stillSelectingFlag bool)
	// The preferred direction of selection.
	SelectionAffinity() NSSelectionAffinity
	// The selection granularity for subsequent extension of a selection.
	SelectionGranularity() NSSelectionGranularity
	SetSelectionGranularity(value NSSelectionGranularity)
	// The color of the insertion point.
	InsertionPointColor() INSColor
	SetInsertionPointColor(value INSColor)
	// Updates the insertion point’s location and optionally restarts the blinking cursor timer.
	UpdateInsertionPointStateAndRestartTimer(restartFlag bool)
	// The attributes used to indicate the selection.
	SelectedTextAttributes() foundation.INSDictionary
	SetSelectedTextAttributes(value foundation.INSDictionary)
	// The attributes used to draw marked text.
	MarkedTextAttributes() foundation.INSDictionary
	SetMarkedTextAttributes(value foundation.INSDictionary)
	// The attributes used to draw the onscreen presentation of link text.
	LinkTextAttributes() foundation.INSDictionary
	SetLinkTextAttributes(value foundation.INSDictionary)
	// Returns a character index appropriate for placing a zero-length selection for an insertion point associated with the mouse at the given point.
	CharacterIndexForInsertionAtPoint(point corefoundation.CGPoint) uint
	UpdateCandidates()

	// Topic: Managing the pasteboard

	// Returns whatever type on the pasteboard would be most preferred for copying data.
	PreferredPasteboardTypeFromArrayRestrictedToTypesFromArray(availableTypes []string, allowedTypes []string) NSPasteboardType
	// Reads the text view’s preferred type of data from the specified pasteboard.
	ReadSelectionFromPasteboard(pboard INSPasteboard) bool
	// Reads data of the given type from the specified pasteboard.
	ReadSelectionFromPasteboardType(pboard INSPasteboard, type_ NSPasteboardType) bool
	// The types this text view can read immediately from the pasteboard.
	ReadablePasteboardTypes() []string
	// The pasteboard types that can be provided from the current selection.
	WritablePasteboardTypes() []string
	// Writes the current selection to the specified pasteboard using the given type.
	WriteSelectionToPasteboardType(pboard INSPasteboard, type_ NSPasteboardType) bool
	// Writes the current selection to the specified pasteboard under each given type.
	WriteSelectionToPasteboardTypes(pboard INSPasteboard, types []string) bool

	// Topic: Setting text attributes

	// Applies full justification to selected paragraphs (or all text, if the receiver is a plain text object).
	AlignJustified(sender objectivec.IObject)
	// Changes the attributes of the current selection.
	ChangeAttributes(sender objectivec.IObject)
	// Sets the color of the selected text.
	ChangeColor(sender objectivec.IObject)
	// Sets the alignment of the paragraphs containing characters in the specified range.
	SetAlignmentRange(alignment NSTextAlignment, range_ foundation.NSRange)
	// The receiver’s typing attributes.
	TypingAttributes() foundation.INSDictionary
	SetTypingAttributes(value foundation.INSDictionary)
	// Set the receiver to use pair kerning data for the glyphs in its selection, or for all glyphs if the receiver is a plain text view.
	UseStandardKerning(sender objectivec.IObject)
	// Lowers the baseline offset of selected text by 1 point, or of all text if the receiver is a plain text view.
	LowerBaseline(sender objectivec.IObject)
	// Raises the baseline offset of selected text by 1 point, or of all text if the receiver is a plain text view.
	RaiseBaseline(sender objectivec.IObject)
	// Sets the receiver to use nominal glyph spacing for the glyphs in its selection, or for all glyphs if the receiver is a plain text view.
	TurnOffKerning(sender objectivec.IObject)
	// Increases the space between glyphs in the receiver’s selection, or in all text if the receiver is a plain text view.
	LoosenKerning(sender objectivec.IObject)
	// Decreases the space between glyphs in the receiver’s selection, or for all glyphs if the receiver is a plain text view.
	TightenKerning(sender objectivec.IObject)
	// Sets the receiver to use the standard ligatures available for the fonts and languages used when setting text, for the glyphs in the selection if the receiver is a rich text view, or for all glyphs if it’s a plain text view.
	UseStandardLigatures(sender objectivec.IObject)
	// Sets the receiver to use only required ligatures when setting text, for the glyphs in the selection if the receiver is a rich text view, or for all glyphs if it’s a plain text view.
	TurnOffLigatures(sender objectivec.IObject)
	// Sets the receiver to use all ligatures available for the fonts and languages used when setting text, for the glyphs in the selection if the receiver is a rich text view, or for all glyphs if it’s a plain text view.
	UseAllLigatures(sender objectivec.IObject)

	// Topic: Clicking and pasting

	// Causes the text view to act as if the user clicked on some text with the given link as the value of a link attribute associated with the text.
	ClickedOnLinkAtIndex(link objectivec.IObject, charIndex uint)
	// Inserts the contents of the pasteboard into the receiver’s text as plain text.
	PasteAsPlainText(sender objectivec.IObject)
	// This action method inserts the contents of the pasteboard into the receiver’s text as rich text, maintaining its attributes.
	PasteAsRichText(sender objectivec.IObject)

	// Topic: Supporting undo

	// Informs the receiver that it should begin coalescing successive typing operations in a new undo grouping.
	BreakUndoCoalescing()
	// A Boolean value that indicates whether undo coalescing is in progress.
	CoalescingUndo() bool

	// Topic: Customizing subclass behaviors

	// Updates the Font panel to contain the font attributes of the selection.
	UpdateFontPanel()
	// Updates the ruler view in the receiver’s enclosing scroll view to reflect the selection’s paragraph and marker attributes.
	UpdateRuler()
	// The data types that the receiver accepts as the destination view of a dragging operation.
	AcceptableDragTypes() []string
	// Updates the acceptable drag types of all text views associated with the receiver’s layout manager.
	UpdateDragTypeRegistration()
	// Returns an adjusted selected range based on the selection granularity.
	SelectionRangeForProposedRangeGranularity(proposedCharRange foundation.NSRange, granularity NSSelectionGranularity) foundation.NSRange
	// The range of characters affected by an action method that changes character (not paragraph) attributes.
	RangeForUserCharacterAttributeChange() foundation.NSRange
	// An array containing the ranges of characters affected by an action method that changes character (not paragraph) attributes.
	RangesForUserCharacterAttributeChange() []foundation.NSValue
	// The range of characters affected by an action method that changes paragraph (not character) attributes.
	RangeForUserParagraphAttributeChange() foundation.NSRange
	// An array containing the ranges of characters affected by a method that changes paragraph (not character) attributes.
	RangesForUserParagraphAttributeChange() []foundation.NSValue
	// The range of characters affected by a method that changes characters (as opposed to attributes).
	RangeForUserTextChange() foundation.NSRange
	// An array containing the ranges of characters affected by a method that changes characters (as opposed to attributes).
	RangesForUserTextChange() []foundation.NSValue
	// Initiates a series of delegate messages (and general notifications) to determine whether modifications can be made to the characters and attributes of the receiver’s text.
	ShouldChangeTextInRangeReplacementString(affectedCharRange foundation.NSRange, replacementString string) bool
	// Initiates a series of delegate messages (and general notifications) to determine whether modifications can be made to the characters and attributes of the receiver’s text.
	ShouldChangeTextInRangesReplacementStrings(affectedRanges []foundation.NSValue, replacementStrings []string) bool
	// Sends out necessary notifications when a text change completes.
	DidChangeText()
	// A Boolean value that controls whether the receiver inserts or deletes space around selected words so as to preserve proper spacing and punctuation.
	SmartInsertDeleteEnabled() bool
	SetSmartInsertDeleteEnabled(value bool)
	// Returns an extended range that includes adjacent whitespace that should be deleted along with the proposed range in order to preserve proper spacing and punctuation.
	SmartDeleteRangeForProposedRange(proposedCharRange foundation.NSRange) foundation.NSRange
	// Returns any whitespace that needs to be added after the string to preserve proper spacing and punctuation when the string replaces the characters in the specified range.
	SmartInsertAfterStringForStringReplacingRange(pasteString string, charRangeToReplace foundation.NSRange) string
	// Returns any whitespace that needs to be added before the string to preserve proper spacing and punctuation when the string replaces the characters in the specified range.
	SmartInsertBeforeStringForStringReplacingRange(pasteString string, charRangeToReplace foundation.NSRange) string
	// Determines whether whitespace needs to be added around the string to preserve proper spacing and punctuation when it replaces the characters in the specified range.
	SmartInsertForStringReplacingRangeBeforeStringAfterString(pasteString string, charRangeToReplace foundation.NSRange, beforeString string, afterString string)
	// Changes the state of smart insert and delete from enabled to disabled and vice versa.
	ToggleSmartInsertDelete(sender objectivec.IObject)

	// Topic: Working with the spelling checker

	// A Boolean value that indicates whether the receiver has continuous spell checking enabled.
	ContinuousSpellCheckingEnabled() bool
	SetContinuousSpellCheckingEnabled(value bool)
	// A tag identifying the text view’s text as a document for the spell checker server.
	SpellCheckerDocumentTag() int
	// Toggles whether continuous spell checking is enabled for the receiver.
	ToggleContinuousSpellChecking(sender objectivec.IObject)
	// Enables and disables grammar checking.
	GrammarCheckingEnabled() bool
	SetGrammarCheckingEnabled(value bool)
	// Changes the state of grammar checking from enabled to disabled and vice versa.
	ToggleGrammarChecking(sender objectivec.IObject)
	// Sets the spelling state, which controls the display of the spelling and grammar indicators on the given text range.
	SetSpellingStateRange(value int, charRange foundation.NSRange)

	// Topic: Working with the sharing service picker

	// Creates and displays a new instance of the sharing service picker.
	OrderFrontSharingServicePicker(sender objectivec.IObject)

	// Topic: Dragging

	// Returns an appropriate drag image for the drag initiated by the specified event.
	DragImageForSelectionWithEventOrigin(event INSEvent, origin foundation.NSPoint) INSImage
	// Returns the type of drag operation that should be performed if the image were released now.
	DragOperationForDraggingInfoType(dragInfo NSDraggingInfo, type_ NSPasteboardType) NSDragOperation
	// Begins dragging the current selected text range.
	DragSelectionWithEventOffsetSlideBack(event INSEvent, mouseOffset corefoundation.CGSize, slideBack bool) bool
	// A Boolean value that indicates whether the receiver accepts the glyph info attribute.
	AcceptsGlyphInfo() bool
	SetAcceptsGlyphInfo(value bool)

	// Topic: Speaking text

	// Speaks the selected text, or all text if no selection.
	StartSpeaking(sender objectivec.IObject)
	// Stops the speaking of text.
	StopSpeaking(sender objectivec.IObject)

	// Topic: Working with panels

	// A Boolean value that indicates whether the receiver allows for a find panel.
	UsesFindPanel() bool
	SetUsesFindPanel(value bool)
	// Performs a find panel action specified by the sender’s tag.
	PerformFindPanelAction(sender objectivec.IObject)
	// Brings forward a panel allowing the user to manipulate links in the text view.
	OrderFrontLinkPanel(sender objectivec.IObject)
	// Brings forward a panel allowing the user to manipulate text lists in the text view.
	OrderFrontListPanel(sender objectivec.IObject)
	// Brings forward a panel allowing the user to manipulate text line heights, interline spacing, and paragraph spacing, in the text view.
	OrderFrontSpacingPanel(sender objectivec.IObject)
	// Brings forward a panel allowing the user to manipulate text tables in the text view.
	OrderFrontTablePanel(sender objectivec.IObject)
	// Brings forward a panel allowing the user to specify string substitutions in the text view.
	OrderFrontSubstitutionsPanel(sender objectivec.IObject)

	// Topic: Performing text completion

	// Returns an array of potential completions, in the order to be presented, representing possible word completions available from a partial word.
	CompletionsForPartialWordRangeIndexOfSelectedItem(charRange foundation.NSRange, index unsafe.Pointer) []string
	// Inserts the selected completion into the text at the appropriate location.
	InsertCompletionForPartialWordRangeMovementIsFinal(word string, charRange foundation.NSRange, movement int, flag bool)
	// The partial range from the most recent beginning of a word up to the insertion point.
	RangeForUserCompletion() foundation.NSRange

	// Topic: Checking and substituting text

	// Performs the default text checking on the entire document.
	CheckTextInDocument(sender objectivec.IObject)
	// Performs the default text checking on the current selection.
	CheckTextInSelection(sender objectivec.IObject)
	// Check and replace the text in the range using the specified checking types and options.
	CheckTextInRangeTypesOptions(range_ foundation.NSRange, checkingTypes uint64, options foundation.INSDictionary)
	// Handles the text checking results returned by the text view
	HandleTextCheckingResultsForRangeTypesOptionsOrthographyWordCount(results []foundation.NSTextCheckingResult, range_ foundation.NSRange, checkingTypes uint64, options foundation.INSDictionary, orthography foundation.NSOrthography, wordCount int)
	// The default text checking types.
	EnabledTextCheckingTypes() uint64
	SetEnabledTextCheckingTypes(value uint64)
	// A Boolean value that indicates whether automatic dash substitution is enabled.
	AutomaticDashSubstitutionEnabled() bool
	SetAutomaticDashSubstitutionEnabled(value bool)
	// Toggles the state of the automatic dash substitution.
	ToggleAutomaticDashSubstitution(sender objectivec.IObject)
	// A Boolean value that indicates whether automatic data detection is enabled.
	AutomaticDataDetectionEnabled() bool
	SetAutomaticDataDetectionEnabled(value bool)
	// Toggles the state of the automatic data detection.
	ToggleAutomaticDataDetection(sender objectivec.IObject)
	// A Boolean value that indicates whether automatic spelling correction is enabled.
	AutomaticSpellingCorrectionEnabled() bool
	SetAutomaticSpellingCorrectionEnabled(value bool)
	// Toggles the state of the automatic spelling correction.
	ToggleAutomaticSpellingCorrection(sender objectivec.IObject)
	// A Boolean value that indicates whether automatic text replacement is enabled.
	AutomaticTextReplacementEnabled() bool
	SetAutomaticTextReplacementEnabled(value bool)
	// Toggles the state of the automatic text replacement.
	ToggleAutomaticTextReplacement(sender objectivec.IObject)
	// Replaces text in the range you specify with the attributed string you provide.
	PerformValidatedReplacementInRangeWithAttributedString(range_ foundation.NSRange, attributedString foundation.NSAttributedString) bool

	// Topic: Getting the writing tools status

	WritingToolsActive() bool

	// Topic: Supporting QuickLook

	// Notifies the QuickLook panel that an update may be required.
	UpdateQuickLookPreviewPanel()
	// An action message that toggles the visibility state of the Quick Look preview panel.
	ToggleQuickLookPreviewPanel(sender objectivec.IObject)
	// Returns an array of URLs for items that can be displayed by QuickLook in the specified ranges.
	QuickLookPreviewableItemsInRanges(ranges []foundation.NSValue) []objectivec.IObject

	// Topic: Changing layout orientation

	// An action method that sets the layout orientation of the text.
	ChangeLayoutOrientation(sender objectivec.IObject)

	// Topic: Using the Find Bar

	// A Boolean value that indicates whether to use the find bar for this text view.
	UsesFindBar() bool
	SetUsesFindBar(value bool)
	// A Boolean value that indicates whether incremental searching is enabled.
	IncrementalSearchingEnabled() bool
	SetIncrementalSearchingEnabled(value bool)

	// Topic: Interacting with the Touch Bar

	AllowsCharacterPickerTouchBarItem() bool
	SetAllowsCharacterPickerTouchBarItem(value bool)
	UpdateTextTouchBarItems()
	UpdateTouchBarItemIdentifiers()

	// Topic: Instance Properties

	AllowedWritingToolsResultOptions() NSWritingToolsResultOptions
	SetAllowedWritingToolsResultOptions(value NSWritingToolsResultOptions)
	InlinePredictionType() NSTextInputTraitType
	SetInlinePredictionType(value NSTextInputTraitType)
	MathExpressionCompletionType() NSTextInputTraitType
	SetMathExpressionCompletionType(value NSTextInputTraitType)
	// ************************* Text Highlight  support **************************
	TextHighlightAttributes() foundation.INSDictionary
	SetTextHighlightAttributes(value foundation.INSDictionary)
	WritingToolsBehavior() NSWritingToolsBehavior
	SetWritingToolsBehavior(value NSWritingToolsBehavior)

	// Topic: Instance Methods

	DrawTextHighlightBackgroundForTextRangeOrigin(textRange INSTextRange, origin corefoundation.CGPoint)
	// An action for toggling [NSTextHighlightStyleAttributeName] in the receiver’s selected range. The sender should be a menu item with a `representedObject` of type ([NSTextHighlightColorScheme]).
	Highlight(sender objectivec.IObject)

	// Type for the find panel metadata property list.
	FindPanelSearchOptions() NSPasteboardType
}

// Init initializes the instance.
func (t NSTextView) Init() NSTextView {
	rv := objc.Send[NSTextView](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTextView) Autorelease() NSTextView {
	rv := objc.Send[NSTextView](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTextView creates a new NSTextView instance.
func NewNSTextView() NSTextView {
	class := getNSTextViewClass()
	rv := objc.Send[NSTextView](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/init(usingTextLayoutManager:)
func NewTextViewUsingTextLayoutManager(usingTextLayoutManager bool) NSTextView {
	instance := getNSTextViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initUsingTextLayoutManager:"), usingTextLayoutManager)
	return NSTextViewFromID(rv)
}

// Initializes a text view with data in an unarchiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/init(coder:)
func NewTextViewWithCoder(coder foundation.INSCoder) NSTextView {
	instance := getNSTextViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSTextViewFromID(rv)
}

// Initializes a text view.
//
// frameRect: The frame rectangle of the text view.
//
// # Return Value
// 
// An initialized text view.
//
// # Discussion
// 
// This method creates the entire collection of objects associated with a text
// view—its text container, layout manager, and text storage—and invokes
// [InitWithFrameTextContainer].
// 
// This method creates the text web in such a manner that the text view is the
// principal owner of the objects in the web.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/init(frame:)
func NewTextViewWithFrame(frameRect corefoundation.CGRect) NSTextView {
	instance := getNSTextViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSTextViewFromID(rv)
}

// Initializes a text view.
//
// frameRect: The frame rectangle of the text view.
//
// container: The text container of the text view.
//
// # Return Value
// 
// An initialized text view.
//
// # Discussion
// 
// This method is the designated initializer for [NSTextView] objects.
// 
// Unlike [InitWithFrame], which builds up an entire group of text-handling
// objects, you use this method after you’ve created the other components of
// the text-handling system — a text storage object, a layout manager, and a
// text container. Assembling the components in this fashion means that the
// text storage, not the text view, is the principal owner of the component
// objects.
// 
// The [InitWithFrame] initializer uses [NSLayoutManager] by default. When you
// use this initializer in macOS 12 and later, you have the option to use
// [NSTextLayoutManager] which gives you access to newer TextKit functionality
// and performance improvements. To use the new layout manager create
// instances of [NSTextLayoutManager], [NSTextContainer], and
// [NSTextContentStorage]; these manage the view’s text layout, text
// regions, and backingstore, respectively. The example below shows the order
// of creation and initialization of these objects, and how configure them to
// initialize an [NSTextView]:
// 
// In macOS 11 and earlier, you follow a similar pattern but using
// [NSLayoutManager] and [NSTextStorage] instead:
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/init(frame:textContainer:)
func NewTextViewWithFrameTextContainer(frameRect corefoundation.CGRect, container INSTextContainer) NSTextView {
	instance := getNSTextViewClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:textContainer:"), frameRect, container)
	return NSTextViewFromID(rv)
}

// Initializes a text view.
//
// frameRect: The frame rectangle of the text view.
//
// container: The text container of the text view.
//
// # Return Value
// 
// An initialized text view.
//
// # Discussion
// 
// This method is the designated initializer for [NSTextView] objects.
// 
// Unlike [InitWithFrame], which builds up an entire group of text-handling
// objects, you use this method after you’ve created the other components of
// the text-handling system — a text storage object, a layout manager, and a
// text container. Assembling the components in this fashion means that the
// text storage, not the text view, is the principal owner of the component
// objects.
// 
// The [InitWithFrame] initializer uses [NSLayoutManager] by default. When you
// use this initializer in macOS 12 and later, you have the option to use
// [NSTextLayoutManager] which gives you access to newer TextKit functionality
// and performance improvements. To use the new layout manager create
// instances of [NSTextLayoutManager], [NSTextContainer], and
// [NSTextContentStorage]; these manage the view’s text layout, text
// regions, and backingstore, respectively. The example below shows the order
// of creation and initialization of these objects, and how configure them to
// initialize an [NSTextView]:
// 
// In macOS 11 and earlier, you follow a similar pattern but using
// [NSLayoutManager] and [NSTextStorage] instead:
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/init(frame:textContainer:)
func (t NSTextView) InitWithFrameTextContainer(frameRect corefoundation.CGRect, container INSTextContainer) NSTextView {
	rv := objc.Send[NSTextView](t.ID, objc.Sel("initWithFrame:textContainer:"), frameRect, container)
	return rv
}
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/init(usingTextLayoutManager:)
func (t NSTextView) InitUsingTextLayoutManager(usingTextLayoutManager bool) NSTextView {
	rv := objc.Send[NSTextView](t.ID, objc.Sel("initUsingTextLayoutManager:"), usingTextLayoutManager)
	return rv
}
// Replaces the text container for the group of text system objects containing
// the receiver, keeping the association between the receiver and its layout
// manager intact.
//
// newContainer: The new text container. This method raises [NSInvalidArgumentException] if
// `aTextContainer` is `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/replaceTextContainer(_:)
func (t NSTextView) ReplaceTextContainer(newContainer INSTextContainer) {
	objc.Send[objc.ID](t.ID, objc.Sel("replaceTextContainer:"), newContainer)
}
// Invalidates the calculated origin of the text container.
//
// # Discussion
// 
// This method is invoked automatically; you should never need to invoke it
// directly. Usually called because the text view has been resized or the
// contents of the text container have changed.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/invalidateTextContainerOrigin()
func (t NSTextView) InvalidateTextContainerOrigin() {
	objc.Send[objc.ID](t.ID, objc.Sel("invalidateTextContainerOrigin"))
}
// An action method used to set the background color.
//
// sender: The control that wants to set the background color.
//
// # Discussion
// 
// This method gets the new color by sending a [Color] message to `sender`.
// 
// This will only set the background color if
// [AllowsDocumentBackgroundColorChange]returns [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/changeDocumentBackgroundColor(_:)
func (t NSTextView) ChangeDocumentBackgroundColor(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("changeDocumentBackgroundColor:"), sender)
}
// Marks the receiver as requiring display.
//
// rect: The rectangle in which display is required.
//
// flag: A value of [true] causes the receiver to not perform any layout, even if
// this means that portions of the text view remain empty. Otherwise the
// receiver performs at least as much layout as needed to display `aRect`.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// [NSTextView] overrides the [NSView] [SetNeedsDisplayInRect] method to
// invoke this method with a `flag` argument of [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/setNeedsDisplay(_:avoidAdditionalLayout:)
func (t NSTextView) SetNeedsDisplayInRectAvoidAdditionalLayout(rect corefoundation.CGRect, flag bool) {
	objc.Send[objc.ID](t.ID, objc.Sel("setNeedsDisplayInRect:avoidAdditionalLayout:"), rect, flag)
}
// Draws or erases the insertion point.
//
// rect: The rectangle in which to draw the insertion point.
//
// color: The color with which to draw the insertion point.
//
// flag: [true] to draw the insertion point, [false] to erase it.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The focus must be locked on the receiver when this method is invoked. You
// should not need to invoke this method directly.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/drawInsertionPoint(in:color:turnedOn:)
func (t NSTextView) DrawInsertionPointInRectColorTurnedOn(rect corefoundation.CGRect, color INSColor, flag bool) {
	objc.Send[objc.ID](t.ID, objc.Sel("drawInsertionPointInRect:color:turnedOn:"), rect, color, flag)
}
// Draws the background of the text view.
//
// rect: The rectangle in which to draw the background.
//
// # Discussion
// 
// Subclasses can override this method to perform additional drawing behind
// the text.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/drawBackground(in:)
func (t NSTextView) DrawViewBackgroundInRect(rect corefoundation.CGRect) {
	objc.Send[objc.ID](t.ID, objc.Sel("drawViewBackgroundInRect:"), rect)
}
// Attempts to set the frame size as if by user action.
//
// desiredSize: The new desired size.
//
// # Discussion
// 
// This method respects the receiver’s existing minimum and maximum sizes
// and by whether resizing is permitted.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/setConstrainedFrameSize(_:)
func (t NSTextView) SetConstrainedFrameSize(desiredSize corefoundation.CGSize) {
	objc.Send[objc.ID](t.ID, objc.Sel("setConstrainedFrameSize:"), desiredSize)
}
// Releases the drag information still existing after the dragging session has
// completed.
//
// # Discussion
// 
// Subclasses may override this method to clean up any additional data
// structures used for dragging. In your overridden method, be sure to invoke
// `super`’s implementation of this method.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/cleanUpAfterDragOperation()
func (t NSTextView) CleanUpAfterDragOperation() {
	objc.Send[objc.ID](t.ID, objc.Sel("cleanUpAfterDragOperation"))
}
// Causes a temporary highlighting effect to appear around the visible portion
// (or portions) of the specified range.
//
// charRange: The character range around which indicators appear.
//
// # Discussion
// 
// This method supports lozenge-style indication of find results. The
// indicators automatically disappear after a certain period of time, or when
// the method is called again, or when any of a number of changes occur to the
// view (such as changes to text, view size, or view position).
// 
// This method does not itself scroll the specified range to be visible; any
// desired scrolling should be done before this method is called, first,
// because the method acts only on the visible portion of the specified range,
// and, second, because scrolling causes the indicators to disappear. Calling
// this method with a zero-length range always removes any existing
// indicators.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/showFindIndicator(for:)
func (t NSTextView) ShowFindIndicatorForRange(charRange foundation.NSRange) {
	objc.Send[objc.ID](t.ID, objc.Sel("showFindIndicatorForRange:"), charRange)
}
// Sets the base writing direction of a range of text.
//
// writingDirection: The new writing direction for the text in `range`.
//
// range: The range of text that will have the new writing direction.
//
// # Discussion
// 
// Invoke this method to change the base writing direction from left-to-right
// to right-to-left for languages like Hebrew and Arabic, for example.
// 
// This method does not include undo support by default. Clients must invoke
// [ShouldChangeTextInRangesReplacementStrings] or
// [ShouldChangeTextInRangeReplacementString] to include this method in an
// undoable action.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/setBaseWritingDirection(_:range:)
func (t NSTextView) SetBaseWritingDirectionRange(writingDirection NSWritingDirection, range_ foundation.NSRange) {
	objc.Send[objc.ID](t.ID, objc.Sel("setBaseWritingDirection:range:"), writingDirection, range_)
}
// Adds the outline attribute to the selected text attributes if absent;
// removes the attribute if present.
//
// sender: The control that sent the message; may be `nil`.
//
// # Discussion
// 
// If there is a selection and the first character of the selected range has a
// non-zero stroke width, or if there is no selection and the typing
// attributes have a non-zero stroke width, then the stroke width is removed;
// otherwise the value of [NSStrokeWidthAttributeName] is set to the default
// value for outline (3.0).
// 
// Operates on the selected range if the receiver contains rich text. For
// plain text the range is the entire contents of the receiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/outline(_:)
func (t NSTextView) Outline(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("outline:"), sender)
}
// Changes the state of automatic quotation mark substitution from enabled to
// disabled and vice versa.
//
// sender: The control sending the message; may be `nil`.
//
// # Discussion
// 
// Automatic quote substitution causes ASCII quotation marks and apostrophes
// to be automatically replaced, on a context-dependent basis, with more
// typographically accurate symbols.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/toggleAutomaticQuoteSubstitution(_:)
func (t NSTextView) ToggleAutomaticQuoteSubstitution(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("toggleAutomaticQuoteSubstitution:"), sender)
}
// Changes the state of automatic link detection from enabled to disabled and
// vice versa.
//
// sender: The control sending the message; may be `nil`.
//
// # Discussion
// 
// Automatic link detection causes strings representing URLs typed in the view
// to be automatically made into links to those URLs.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/toggleAutomaticLinkDetection(_:)
func (t NSTextView) ToggleAutomaticLinkDetection(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("toggleAutomaticLinkDetection:"), sender)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/toggleAutomaticTextCompletion(_:)
func (t NSTextView) ToggleAutomaticTextCompletion(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("toggleAutomaticTextCompletion:"), sender)
}
// Sets the selection to a range of characters in response to user action.
//
// charRange: The range of characters to select. This range must begin and end on glyph
// boundaries and not split base glyphs and their nonspacing marks.
//
// affinity: The selection affinity for the selection. See [SelectionAffinity] for more
// information about how affinities work.
//
// stillSelectingFlag: [true] to behave appropriately for a continuing selection where the user is
// still dragging the mouse, [false] otherwise. If [true], the receiver
// doesn’t send notifications or remove the marking from its marked text. If
// [false], the receiver posts an [didChangeSelectionNotification] to the
// default notification center and removes the marking from marked text if the
// new selection is greater than the marked region.
// //
// [didChangeSelectionNotification]: https://developer.apple.com/documentation/AppKit/NSTextView/didChangeSelectionNotification
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method resets the selection granularity to [NSSelectByCharacter].
// 
// # Special Considerations
// 
// In macOS 10.4 and later, if there are multiple selections, this method acts
// on the first selected subrange.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/setSelectedRange(_:affinity:stillSelecting:)
func (t NSTextView) SetSelectedRangeAffinityStillSelecting(charRange foundation.NSRange, affinity NSSelectionAffinity, stillSelectingFlag bool) {
	objc.Send[objc.ID](t.ID, objc.Sel("setSelectedRange:affinity:stillSelecting:"), charRange, affinity, stillSelectingFlag)
}
// Sets the selection to the characters in an array of ranges in response to
// user action.
//
// ranges: A non-nil, non-empty array of objects responding to the NSValue
// `rangeValue` method. The ranges in the `ranges` array must begin and end on
// glyph boundaries and not split base glyphs and their nonspacing marks.
//
// affinity: The selection affinity for the selection. See [SelectionAffinity] for more
// information about how affinities work.
//
// stillSelectingFlag: [true] to behave appropriately for a continuing selection where the user is
// still dragging the mouse, [false] otherwise. If [true], the receiver
// doesn’t send notifications or remove the marking from its marked text. If
// [false], the receiver posts an [didChangeSelectionNotification] to the
// default notification center and removes the marking from marked text if the
// new selection is greater than the marked region.
// //
// [didChangeSelectionNotification]: https://developer.apple.com/documentation/AppKit/NSTextView/didChangeSelectionNotification
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method also resets the selection granularity to [NSSelectByCharacter].
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/setSelectedRanges(_:affinity:stillSelecting:)
func (t NSTextView) SetSelectedRangesAffinityStillSelecting(ranges []foundation.NSValue, affinity NSSelectionAffinity, stillSelectingFlag bool) {
	objc.Send[objc.ID](t.ID, objc.Sel("setSelectedRanges:affinity:stillSelecting:"), objectivec.IObjectSliceToNSArray(ranges), affinity, stillSelectingFlag)
}
// Updates the insertion point’s location and optionally restarts the
// blinking cursor timer.
//
// restartFlag: [true] to restart the blinking cursor timer, [false] otherwise.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method is invoked automatically whenever the insertion point needs to
// be moved; you should never need to invoke it directly, but you can override
// it to modify insertion point behavior.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/updateInsertionPointStateAndRestartTimer(_:)
func (t NSTextView) UpdateInsertionPointStateAndRestartTimer(restartFlag bool) {
	objc.Send[objc.ID](t.ID, objc.Sel("updateInsertionPointStateAndRestartTimer:"), restartFlag)
}
// Returns a character index appropriate for placing a zero-length selection
// for an insertion point associated with the mouse at the given point.
//
// point: The point for which to return an index, in view coordinates.
//
// # Return Value
// 
// The character index for the insertion point.
//
// # Discussion
// 
// This method should be used for insertion points associated with mouse
// clicks, drag events, and so forth. For other purposes, it is better to use
// [NSLayoutManager] methods.
// 
// The [NSTextInput] method [characterIndexForPoint:] is not suitable for this
// role; it is intended only for uses related to text input methods.
//
// [characterIndexForPoint:]: https://developer.apple.com/documentation/AppKit/NSTextInput/characterIndexForPoint:
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/characterIndexForInsertion(at:)
func (t NSTextView) CharacterIndexForInsertionAtPoint(point corefoundation.CGPoint) uint {
	rv := objc.Send[uint](t.ID, objc.Sel("characterIndexForInsertionAtPoint:"), point)
	return rv
}
// See: https://developer.apple.com/documentation/AppKit/NSTextView/updateCandidates()
func (t NSTextView) UpdateCandidates() {
	objc.Send[objc.ID](t.ID, objc.Sel("updateCandidates"))
}
// Returns whatever type on the pasteboard would be most preferred for copying
// data.
//
// availableTypes: The types currently available on the pasteboard.
//
// allowedTypes: Types allowed in the return value. If `nil`, any available type is allowed.
//
// # Return Value
// 
// The preferred type to provide given the available types and the allowed
// types.
//
// # Discussion
// 
// You should not need to override this method. You should also not need to
// invoke it unless you are implementing a new type of pasteboard to handle
// services other than copy/paste or dragging.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/preferredPasteboardType(from:restrictedToTypesFrom:)
func (t NSTextView) PreferredPasteboardTypeFromArrayRestrictedToTypesFromArray(availableTypes []string, allowedTypes []string) NSPasteboardType {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("preferredPasteboardTypeFromArray:restrictedToTypesFromArray:"), objectivec.StringSliceToNSArray(availableTypes), objectivec.StringSliceToNSArray(allowedTypes))
	return NSPasteboardType(foundation.NSStringFromID(rv).String())
}
// Reads the text view’s preferred type of data from the specified
// pasteboard.
//
// pboard: The pasteboard to read from.
//
// # Return Value
// 
// [true] if the data was successfully read, [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method invokes the
// [PreferredPasteboardTypeFromArrayRestrictedToTypesFromArray] method to
// determine the text view’s preferred type of data and then reads the data
// using the [ReadSelectionFromPasteboardType] method.
// 
// You should not need to override this method. You might need to invoke this
// method if you are implementing a new type of pasteboard to handle services
// other than copy/paste or dragging.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/readSelection(from:)
func (t NSTextView) ReadSelectionFromPasteboard(pboard INSPasteboard) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("readSelectionFromPasteboard:"), pboard)
	return rv
}
// Reads data of the given type from the specified pasteboard.
//
// pboard: The pasteboard to read from.
//
// type: The type of data to read.
//
// # Return Value
// 
// [true] if the data was successfully read, [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The new data is placed at the current insertion point, replacing the
// current selection if one exists.
// 
// You should override this method to read pasteboard types other than the
// default types. Use the [RangeForUserTextChange] method to obtain the range
// of characters (if any) to be replaced by the new data.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/readSelection(from:type:)
func (t NSTextView) ReadSelectionFromPasteboardType(pboard INSPasteboard, type_ NSPasteboardType) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("readSelectionFromPasteboard:type:"), pboard, objc.String(string(type_)))
	return rv
}
// Writes the current selection to the specified pasteboard using the given
// type.
//
// pboard: The pasteboard to write to.
//
// type: The type of data to write.
//
// # Return Value
// 
// [true] if the data was successfully written, [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The complete set of data types being written to `pboard` should be declared
// before invoking this method.
// 
// This method should be invoked only from [WriteSelectionToPasteboardTypes].
// You can override this method to add support for writing new types of data
// to the pasteboard. You should invoke `super`’s implementation of the
// method to handle any types of data your overridden version does not.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/writeSelection(to:type:)
func (t NSTextView) WriteSelectionToPasteboardType(pboard INSPasteboard, type_ NSPasteboardType) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("writeSelectionToPasteboard:type:"), pboard, objc.String(string(type_)))
	return rv
}
// Writes the current selection to the specified pasteboard under each given
// type.
//
// pboard: The pasteboard to write to.
//
// types: An array of strings describing the types of data to write.
//
// # Return Value
// 
// [true] if the data for any single type was successfully written, [false]
// otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method declares the data types on `pboard` and then invokes
// [WriteSelectionToPasteboardType] or the delegate method
// [TextViewWriteCellAtIndexToPasteboardType] for each type in the `types`
// array.
// 
// You should not need to override this method. You might need to invoke this
// method if you are implementing a new type of pasteboard to handle services
// other than copy/paste or dragging.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/writeSelection(to:types:)
func (t NSTextView) WriteSelectionToPasteboardTypes(pboard INSPasteboard, types []string) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("writeSelectionToPasteboard:types:"), pboard, objectivec.StringSliceToNSArray(types))
	return rv
}
// Applies full justification to selected paragraphs (or all text, if the
// receiver is a plain text object).
//
// sender: The control that sent the message; may be `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/alignJustified(_:)
func (t NSTextView) AlignJustified(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("alignJustified:"), sender)
}
// Changes the attributes of the current selection.
//
// sender: The control that sent the message. Must respond to ``.
//
// # Discussion
// 
// This method changes the attributes by invoking [ConvertAttributes] on
// `sender` and applying the returned attributes to the appropriate text. See
// [Font Handling] in [Cocoa Text Architecture Guide] for more information on
// attribute conversion.
//
// [Cocoa Text Architecture Guide]: https://developer.apple.com/library/archive/documentation/TextFonts/Conceptual/CocoaTextArchitecture/Introduction/Introduction.html#//apple_ref/doc/uid/TP40009459
// [Font Handling]: https://developer.apple.com/library/archive/documentation/TextFonts/Conceptual/CocoaTextArchitecture/FontHandling/FontHandling.html#//apple_ref/doc/uid/TP40009459-CH5
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/changeAttributes(_:)
func (t NSTextView) ChangeAttributes(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("changeAttributes:"), sender)
}
// Sets the color of the selected text.
//
// sender: The control that sent the message. [NSTextView]’s implementation sends a
// [Color] message to `sender` to get the new color.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/changeColor(_:)
func (t NSTextView) ChangeColor(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("changeColor:"), sender)
}
// Sets the alignment of the paragraphs containing characters in the specified
// range.
//
// alignment: The new alignment.
//
// range: The range of characters whose paragraphs will have their alignment set.
//
// # Discussion
// 
// This method does not include undo support by default. Clients must invoke
// [ShouldChangeTextInRangesReplacementStrings] or
// [ShouldChangeTextInRangeReplacementString] to include this method in an
// undoable action.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/setAlignment(_:range:)
func (t NSTextView) SetAlignmentRange(alignment NSTextAlignment, range_ foundation.NSRange) {
	objc.Send[objc.ID](t.ID, objc.Sel("setAlignment:range:"), alignment, range_)
}
// Set the receiver to use pair kerning data for the glyphs in its selection,
// or for all glyphs if the receiver is a plain text view.
//
// sender: The control that sent the message; may be `nil`.
//
// # Discussion
// 
// This data is taken from a font’s AFM file
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/useStandardKerning(_:)
func (t NSTextView) UseStandardKerning(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("useStandardKerning:"), sender)
}
// Lowers the baseline offset of selected text by 1 point, or of all text if
// the receiver is a plain text view.
//
// sender: The control that sent the message; may be `nil`.
//
// # Discussion
// 
// As such, this method defines a more primitive operation than subscripting.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/lowerBaseline(_:)
func (t NSTextView) LowerBaseline(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("lowerBaseline:"), sender)
}
// Raises the baseline offset of selected text by 1 point, or of all text if
// the receiver is a plain text view.
//
// sender: The control that sent the message; may be `nil`.
//
// # Discussion
// 
// As such, this method defines a more primitive operation than
// superscripting.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/raiseBaseline(_:)
func (t NSTextView) RaiseBaseline(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("raiseBaseline:"), sender)
}
// Sets the receiver to use nominal glyph spacing for the glyphs in its
// selection, or for all glyphs if the receiver is a plain text view.
//
// sender: The control that sent the message; may be `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/turnOffKerning(_:)
func (t NSTextView) TurnOffKerning(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("turnOffKerning:"), sender)
}
// Increases the space between glyphs in the receiver’s selection, or in all
// text if the receiver is a plain text view.
//
// sender: The control that sent the message; may be `nil`.
//
// # Discussion
// 
// Kerning values are determined by the point size of the fonts in the
// selection.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/loosenKerning(_:)
func (t NSTextView) LoosenKerning(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("loosenKerning:"), sender)
}
// Decreases the space between glyphs in the receiver’s selection, or for
// all glyphs if the receiver is a plain text view.
//
// sender: The control that sent the message; may be `nil`.
//
// # Discussion
// 
// Kerning values are determined by the point size of the fonts in the
// selection.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/tightenKerning(_:)
func (t NSTextView) TightenKerning(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("tightenKerning:"), sender)
}
// Sets the receiver to use the standard ligatures available for the fonts and
// languages used when setting text, for the glyphs in the selection if the
// receiver is a rich text view, or for all glyphs if it’s a plain text
// view.
//
// sender: The control that sent the message; may be `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/useStandardLigatures(_:)
func (t NSTextView) UseStandardLigatures(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("useStandardLigatures:"), sender)
}
// Sets the receiver to use only required ligatures when setting text, for the
// glyphs in the selection if the receiver is a rich text view, or for all
// glyphs if it’s a plain text view.
//
// sender: The control that sent the message; may be `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/turnOffLigatures(_:)
func (t NSTextView) TurnOffLigatures(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("turnOffLigatures:"), sender)
}
// Sets the receiver to use all ligatures available for the fonts and
// languages used when setting text, for the glyphs in the selection if the
// receiver is a rich text view, or for all glyphs if it’s a plain text
// view.
//
// sender: The control that sent the message; may be `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/useAllLigatures(_:)
func (t NSTextView) UseAllLigatures(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("useAllLigatures:"), sender)
}
// Causes the text view to act as if the user clicked on some text with the
// given link as the value of a link attribute associated with the text.
//
// link: The link that was clicked; the value of [link].
// //
// [link]: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key/link
//
// charIndex: The character index where the click occurred, indexed within the text
// storage.
//
// # Discussion
// 
// If, for instance, you have a special attachment cell that can follow links,
// you can use this method to ask the text view to follow a link once you
// decide it should. In addition, this method is invoked by the text view
// during mouse tracking if the user clicks a link.
// 
// The `charIndex` parameter is a character index somewhere in the range of
// the link attribute. If the user actually physically clicked the link, then
// it should be the character that was originally clicked. In some cases a
// link may be opened indirectly or programmatically, in which case a
// character index somewhere in the range of the link attribute is supplied.
// 
// This method sends the [TextViewClickedOnLinkAtIndex] delegate message if
// the delegate implements it, so that the delegate can handle the click.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/clicked(onLink:at:)
func (t NSTextView) ClickedOnLinkAtIndex(link objectivec.IObject, charIndex uint) {
	objc.Send[objc.ID](t.ID, objc.Sel("clickedOnLink:atIndex:"), link, charIndex)
}
// Inserts the contents of the pasteboard into the receiver’s text as plain
// text.
//
// sender: The control that sent the message; may be `nil`.
//
// # Discussion
// 
// This method behaves analogously to [insertText(_:)].
//
// [insertText(_:)]: https://developer.apple.com/documentation/AppKit/NSTextView/insertText(_:)
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/pasteAsPlainText(_:)
func (t NSTextView) PasteAsPlainText(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("pasteAsPlainText:"), sender)
}
// This action method inserts the contents of the pasteboard into the
// receiver’s text as rich text, maintaining its attributes.
//
// sender: The control that sent the message; may be `nil`.
//
// # Discussion
// 
// The text is inserted at the insertion point if there is one, otherwise
// replacing the selection.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/pasteAsRichText(_:)
func (t NSTextView) PasteAsRichText(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("pasteAsRichText:"), sender)
}
// Informs the receiver that it should begin coalescing successive typing
// operations in a new undo grouping.
//
// # Discussion
// 
// This method should be invoked when saving the receiver’s contents to
// preserve proper tracking of unsaved changes and the document’s dirty
// state.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/breakUndoCoalescing()
func (t NSTextView) BreakUndoCoalescing() {
	objc.Send[objc.ID](t.ID, objc.Sel("breakUndoCoalescing"))
}
// Updates the Font panel to contain the font attributes of the selection.
//
// # Discussion
// 
// Does nothing if the receiver doesn’t use the Font panel. You should never
// need to invoke this method directly, but you can override it if needed to
// handle additional font attributes.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/updateFontPanel()
func (t NSTextView) UpdateFontPanel() {
	objc.Send[objc.ID](t.ID, objc.Sel("updateFontPanel"))
}
// Updates the ruler view in the receiver’s enclosing scroll view to reflect
// the selection’s paragraph and marker attributes.
//
// # Discussion
// 
// Does nothing if the ruler isn’t visible or if the receiver doesn’t use
// the ruler. You should never need to invoke this method directly, but you
// can override this method if needed to handle additional ruler attributes.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/updateRuler()
func (t NSTextView) UpdateRuler() {
	objc.Send[objc.ID](t.ID, objc.Sel("updateRuler"))
}
// Updates the acceptable drag types of all text views associated with the
// receiver’s layout manager.
//
// # Discussion
// 
// If the receiver is editable and is a rich text view, causes all text views
// associated with the receiver’s layout manager to register their
// acceptable drag types. If the text view isn’t editable or isn’t rich
// text, causes those text views to unregister their dragged types.
// 
// Subclasses can override this method to change the conditions for
// registering and unregistering drag types, whether as a group or
// individually based on the current state of the text view. They should
// invoke this method when that state changes to perform the necessary update.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/updateDragTypeRegistration()
func (t NSTextView) UpdateDragTypeRegistration() {
	objc.Send[objc.ID](t.ID, objc.Sel("updateDragTypeRegistration"))
}
// Returns an adjusted selected range based on the selection granularity.
//
// proposedCharRange: The proposed selected range.
//
// granularity: The selection granularity.
//
// # Return Value
// 
// The adjusted selected range, taking into account the selection granularity.
//
// # Discussion
// 
// This method is invoked repeatedly during mouse tracking to modify the range
// of the selection. Override this method to specialize selection behavior.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/selectionRange(forProposedRange:granularity:)
func (t NSTextView) SelectionRangeForProposedRangeGranularity(proposedCharRange foundation.NSRange, granularity NSSelectionGranularity) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](t.ID, objc.Sel("selectionRangeForProposedRange:granularity:"), proposedCharRange, granularity)
	return foundation.NSRange(rv)
}
// Initiates a series of delegate messages (and general notifications) to
// determine whether modifications can be made to the characters and
// attributes of the receiver’s text.
//
// affectedCharRange: The range of characters affected by the proposed change.
//
// replacementString: The characters that will replace those in `affectedCharRange`. If only text
// attributes are being changed, `replacementString` is `nil`.
//
// # Return Value
// 
// [true] to allow the change, [false] to prohibit it.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method checks with the delegate as needed using
// [TextShouldBeginEditing] and
// [TextViewShouldChangeTextInRangeReplacementString].
// 
// This method must be invoked at the start of any sequence of user-initiated
// editing changes. If your subclass of [NSTextView] implements methods that
// modify the text, make sure to invoke this method to determine whether the
// change should be made. If the change is allowed, complete the change by
// invoking the [DidChangeText] method.
// 
// # Special Considerations
// 
// If you override this method, you must call `super` at the beginning of the
// override.
// 
// If the receiver is not editable, this method automatically returns [false].
// This result prevents instances in which a text view could be changed by
// user actions even though it had been set to be non-editable.
// 
// In macOS 10.4 and later, if there are multiple selections, this method acts
// on the first selected subrange.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/shouldChangeText(in:replacementString:)
func (t NSTextView) ShouldChangeTextInRangeReplacementString(affectedCharRange foundation.NSRange, replacementString string) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("shouldChangeTextInRange:replacementString:"), affectedCharRange, objc.String(replacementString))
	return rv
}
// Initiates a series of delegate messages (and general notifications) to
// determine whether modifications can be made to the characters and
// attributes of the receiver’s text.
//
// affectedRanges: An array of ranges to change.
//
// replacementStrings: An array of strings containing the characters that replace those in
// `affectedRanges`, one for each range. If only text attributes are being
// changed, `replacementStrings` is `nil`.
//
// # Return Value
// 
// [true] to allow the change, [false] to prohibit it.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method checks with the delegate as needed using
// [TextShouldBeginEditing] and
// [TextViewShouldChangeTextInRangesReplacementStrings].
// 
// This method must be invoked at the start of any sequence of user-initiated
// editing changes. If your subclass of [NSTextView] implements
// 
// that modify the text, make sure to invoke this method to determine whether
// the change should be made. If the change is allowed, complete the change by
// invoking the [DidChangeText] method. If you can’t determine the affected
// range or replacement string before beginning changes, pass `nil` for these
// values.
// 
// # Special Considerations
// 
// If you override this method, you must call `super` at the beginning of the
// override.
// 
// If the receiver is not editable, this method automatically returns [false].
// This result prevents instances in which a text view could be changed by
// user actions even though it had been set to be non-editable.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/shouldChangeText(inRanges:replacementStrings:)
func (t NSTextView) ShouldChangeTextInRangesReplacementStrings(affectedRanges []foundation.NSValue, replacementStrings []string) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("shouldChangeTextInRanges:replacementStrings:"), objectivec.IObjectSliceToNSArray(affectedRanges), objectivec.StringSliceToNSArray(replacementStrings))
	return rv
}
// Sends out necessary notifications when a text change completes.
//
// # Discussion
// 
// Invoked automatically at the end of a series of changes, this method posts
// an [didChangeNotification] to the default notification center, which also
// results in the delegate receiving an [NSText] delegate [TextDidChange]
// message.
// 
// Subclasses implementing methods that change their text should invoke this
// method at the end of those methods. See [Subclassing NSTextView] for more
// information.
//
// [Subclassing NSTextView]: https://developer.apple.com/library/archive/documentation/TextFonts/Conceptual/CocoaTextArchitecture/TextEditing/TextEditing.html#//apple_ref/doc/uid/TP40009459-CH3-SW16
// [didChangeNotification]: https://developer.apple.com/documentation/AppKit/NSText/didChangeNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/didChangeText()
func (t NSTextView) DidChangeText() {
	objc.Send[objc.ID](t.ID, objc.Sel("didChangeText"))
}
// Returns an extended range that includes adjacent whitespace that should be
// deleted along with the proposed range in order to preserve proper spacing
// and punctuation.
//
// proposedCharRange: The proposed character range for deleting.
//
// # Return Value
// 
// An extended range that includes adjacent whitespace that should be deleted
// along with the proposed range in order to preserve proper spacing and
// punctuation of the text surrounding the deletion.
//
// # Discussion
// 
// [NSTextView] uses this method as necessary; you can also use it in
// implementing your own methods that delete text, typically when the
// selection granularity is [NSSelectByWord]. To do so, invoke this method
// with the proposed range to delete, then actually delete the range returned.
// If placing text on the pasteboard, however, you should put only the
// characters from the proposed range onto the pasteboard.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/smartDeleteRange(forProposedRange:)
func (t NSTextView) SmartDeleteRangeForProposedRange(proposedCharRange foundation.NSRange) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](t.ID, objc.Sel("smartDeleteRangeForProposedRange:"), proposedCharRange)
	return foundation.NSRange(rv)
}
// Returns any whitespace that needs to be added after the string to preserve
// proper spacing and punctuation when the string replaces the characters in
// the specified range.
//
// pasteString: The string that is replacing the characters in `charRange`.
//
// charRangeToReplace: The range of characters which `aString` is replacing.
//
// # Return Value
// 
// Any whitespace that needs to be added after `aString` to preserve proper
// spacing and punctuation when the characters in `charRange` are replaced by
// `aString`. If `aString` is `nil` or if smart insertion and deletion are
// disabled, this method returns `nil`.
//
// # Discussion
// 
// Don’t invoke this method directly. Instead, use
// [SmartInsertForStringReplacingRangeBeforeStringAfterString], which calls
// this method as part of its implementation.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/smartInsert(afterStringFor:replacing:)
func (t NSTextView) SmartInsertAfterStringForStringReplacingRange(pasteString string, charRangeToReplace foundation.NSRange) string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("smartInsertAfterStringForString:replacingRange:"), objc.String(pasteString), charRangeToReplace)
	return foundation.NSStringFromID(rv).String()
}
// Returns any whitespace that needs to be added before the string to preserve
// proper spacing and punctuation when the string replaces the characters in
// the specified range.
//
// pasteString: The string that is replacing the characters in `charRange`.
//
// charRangeToReplace: The range of characters which `aString` is replacing.
//
// # Return Value
// 
// Any whitespace that needs to be added before `aString` to preserve proper
// spacing and punctuation when the characters in `charRange` are replaced by
// `aString`. If `aString` is `nil` or if smart insertion and deletion are
// disabled, this method returns `nil`.
//
// # Discussion
// 
// Don’t invoke this method directly. Instead, use
// [SmartInsertForStringReplacingRangeBeforeStringAfterString], which calls
// this method as part of its implementation.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/smartInsert(beforeStringFor:replacing:)
func (t NSTextView) SmartInsertBeforeStringForStringReplacingRange(pasteString string, charRangeToReplace foundation.NSRange) string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("smartInsertBeforeStringForString:replacingRange:"), objc.String(pasteString), charRangeToReplace)
	return foundation.NSStringFromID(rv).String()
}
// Determines whether whitespace needs to be added around the string to
// preserve proper spacing and punctuation when it replaces the characters in
// the specified range.
//
// pasteString: The string that is replacing the characters in `charRange`.
//
// charRangeToReplace: The range of characters which `aString` is replacing.
//
// beforeString: On return, a pointer to the string with the characters that should be added
// before `aString`; `nil` if there are no characters to add, if `aString` is
// `nil`, or if smart insertion and deletion are disabled.
//
// afterString: On return, a pointer to the string with the characters that should be added
// after `aString`; `nil` if there are no characters to add, if `aString` is
// `nil`, or if smart insertion and deletion are disabled.
//
// # Discussion
// 
// As part of its implementation, this method calls
// [SmartInsertAfterStringForStringReplacingRange] and
// [SmartInsertBeforeStringForStringReplacingRange]. To change this method’s
// behavior, override those two methods instead of this one.
// 
// [NSTextView] uses this method as necessary. You can also use it in
// implementing your own methods that insert text. To do so, invoke this
// method with the proper arguments, then insert `beforeString`, `aString`,
// and `afterString` in order over `charRange`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/smartInsert(for:replacing:before:after:)
func (t NSTextView) SmartInsertForStringReplacingRangeBeforeStringAfterString(pasteString string, charRangeToReplace foundation.NSRange, beforeString string, afterString string) {
	objc.Send[objc.ID](t.ID, objc.Sel("smartInsertForString:replacingRange:beforeString:afterString:"), objc.String(pasteString), charRangeToReplace, objc.String(beforeString), objc.String(afterString))
}
// Changes the state of smart insert and delete from enabled to disabled and
// vice versa.
//
// sender: The control sending the message; may be `nil`.
//
// # Discussion
// 
// Controls whether the receiver inserts or deletes space around selected
// words so as to preserve proper spacing and punctuation.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/toggleSmartInsertDelete(_:)
func (t NSTextView) ToggleSmartInsertDelete(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("toggleSmartInsertDelete:"), sender)
}
// Toggles whether continuous spell checking is enabled for the receiver.
//
// sender: The control sending the message; may be `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/toggleContinuousSpellChecking(_:)
func (t NSTextView) ToggleContinuousSpellChecking(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("toggleContinuousSpellChecking:"), sender)
}
// Changes the state of grammar checking from enabled to disabled and vice
// versa.
//
// sender: The control sending the message; may be `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/toggleGrammarChecking(_:)
func (t NSTextView) ToggleGrammarChecking(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("toggleGrammarChecking:"), sender)
}
// Sets the spelling state, which controls the display of the spelling and
// grammar indicators on the given text range.
//
// value: The spelling state value to set. Possible values, for the temporary
// attribute on the layout manager using the key NSSpellingStateAttributeName,
// are:
// 
// - [NSSpellingStateSpellingFlag] to highlight spelling issues. -
// [NSSpellingStateGrammarFlag] to highlight grammar issues.
// //
// [NSSpellingStateGrammarFlag]: https://developer.apple.com/documentation/AppKit/NSSpellingState/NSSpellingStateGrammarFlag
// [NSSpellingStateSpellingFlag]: https://developer.apple.com/documentation/AppKit/NSSpellingState/NSSpellingStateSpellingFlag
//
// charRange: The character range over which to set the given spelling state.
//
// # Discussion
// 
// May be called or overridden to control setting of spelling and grammar
// indicators on text, used to highlight portions of the text that are flagged
// for spelling or grammar issues.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/setSpellingState(_:range:)
func (t NSTextView) SetSpellingStateRange(value int, charRange foundation.NSRange) {
	objc.Send[objc.ID](t.ID, objc.Sel("setSpellingState:range:"), value, charRange)
}
// Creates and displays a new instance of the sharing service picker.
//
// sender: The sender.
//
// # Discussion
// 
// Creates a new instance of [NSSharingServicePicker] based on the current
// selection and shows to the screen. The items passed to the
// [NSSharingServicePicker] initializer are determined using the delegate
// method ``.
// 
// When the current selection is 0 length, the whole document is passed to the
// method.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/orderFrontSharingServicePicker(_:)
func (t NSTextView) OrderFrontSharingServicePicker(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("orderFrontSharingServicePicker:"), sender)
}
// Returns an appropriate drag image for the drag initiated by the specified
// event.
//
// event: The event that initiated the drag session.
//
// origin: On return, the lower-left point of the image in view coordinates.
//
// # Return Value
// 
// An appropriate drag image for the drag initiated by `event`. May be `nil`,
// in which case a default icon will be used.
//
// # Discussion
// 
// This method is used by [DragSelectionWithEventOffsetSlideBack]. It can be
// called by others who need such an image, or can be overridden by subclasses
// to return a different image.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/dragImageForSelection(with:origin:)
func (t NSTextView) DragImageForSelectionWithEventOrigin(event INSEvent, origin foundation.NSPoint) INSImage {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("dragImageForSelectionWithEvent:origin:"), event, origin)
	return NSImageFromID(rv)
}
// Returns the type of drag operation that should be performed if the image
// were released now.
//
// dragInfo: The drag information.
//
// type: The pasteboard type that will be read from the dragging pasteboard.
//
// # Return Value
// 
// The drag operation that should be performed if the image were released now.
//
// # Discussion
// 
// The returned value should be one of the following:
// 
// [Table data omitted]
// 
// If none of the operations is appropriate, this method should return
// [NSDragOperationNone].
// 
// This method is called repeatedly from [DraggingEntered] and
// [DraggingUpdated] as the user drags the image.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/dragOperation(for:type:)
func (t NSTextView) DragOperationForDraggingInfoType(dragInfo NSDraggingInfo, type_ NSPasteboardType) NSDragOperation {
	rv := objc.Send[NSDragOperation](t.ID, objc.Sel("dragOperationForDraggingInfo:type:"), dragInfo, objc.String(string(type_)))
	return NSDragOperation(rv)
}
// Begins dragging the current selected text range.
//
// event: The event that initiated dragging the selection.
//
// mouseOffset: The cursor’s current location relative to the mouse-down `event`.
//
// slideBack: [true] if the image being dragged should slide back to its original
// position if the drag does not succeed, [false] otherwise.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// [true] if the drag can be successfully initiated, [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Primarily for subclasses, who can override it to intervene at the beginning
// of a drag.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/dragSelection(with:offset:slideBack:)
func (t NSTextView) DragSelectionWithEventOffsetSlideBack(event INSEvent, mouseOffset corefoundation.CGSize, slideBack bool) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("dragSelectionWithEvent:offset:slideBack:"), event, mouseOffset, slideBack)
	return rv
}
// Speaks the selected text, or all text if no selection.
//
// sender: The control sending the message; can be `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/startSpeaking(_:)
func (t NSTextView) StartSpeaking(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("startSpeaking:"), sender)
}
// Stops the speaking of text.
//
// sender: The control sending the message; can be `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/stopSpeaking(_:)
func (t NSTextView) StopSpeaking(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("stopSpeaking:"), sender)
}
// Performs a find panel action specified by the sender’s tag.
//
// sender: The control sending the message. This method sends the [Tag] method to
// determine what operation to perform. The list of possible tags is provided
// in Constants.
//
// # Discussion
// 
// This is the generic action method for the find menu and find panel, and can
// be overridden to implement a custom find panel.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/performFindPanelAction(_:)
func (t NSTextView) PerformFindPanelAction(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("performFindPanelAction:"), sender)
}
// Brings forward a panel allowing the user to manipulate links in the text
// view.
//
// sender: The control sending the message. May be `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/orderFrontLinkPanel(_:)
func (t NSTextView) OrderFrontLinkPanel(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("orderFrontLinkPanel:"), sender)
}
// Brings forward a panel allowing the user to manipulate text lists in the
// text view.
//
// sender: The control sending the message. May be `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/orderFrontListPanel(_:)
func (t NSTextView) OrderFrontListPanel(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("orderFrontListPanel:"), sender)
}
// Brings forward a panel allowing the user to manipulate text line heights,
// interline spacing, and paragraph spacing, in the text view.
//
// sender: The control sending the message. May be `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/orderFrontSpacingPanel(_:)
func (t NSTextView) OrderFrontSpacingPanel(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("orderFrontSpacingPanel:"), sender)
}
// Brings forward a panel allowing the user to manipulate text tables in the
// text view.
//
// sender: The control sending the message. May be `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/orderFrontTablePanel(_:)
func (t NSTextView) OrderFrontTablePanel(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("orderFrontTablePanel:"), sender)
}
// Brings forward a panel allowing the user to specify string substitutions in
// the text view.
//
// sender: The control sending the message. May be `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/orderFrontSubstitutionsPanel(_:)
func (t NSTextView) OrderFrontSubstitutionsPanel(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("orderFrontSubstitutionsPanel:"), sender)
}
// Returns an array of potential completions, in the order to be presented,
// representing possible word completions available from a partial word.
//
// charRange: The range of characters of the matched partial word to be completed.
//
// index: On return, optionally set to the completion that should be initially
// selected. The default is 0, and –1 indicates no selection.
//
// # Return Value
// 
// An array of potential completions, in the order to be presented,
// representing possible word completions available from a partial word at
// `charRange`. Returning `nil` or a zero-length array suppresses completion.
//
// # Discussion
// 
// May be overridden by subclasses to modify or override the list of possible
// completions.
// 
// This method should call the delegate method
// [TextViewCompletionsForPartialWordRangeIndexOfSelectedItem] if the delegate
// implements such a method.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/completions(forPartialWordRange:indexOfSelectedItem:)
func (t NSTextView) CompletionsForPartialWordRangeIndexOfSelectedItem(charRange foundation.NSRange, index unsafe.Pointer) []string {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("completionsForPartialWordRange:indexOfSelectedItem:"), charRange, index)
	return objc.ConvertSliceToStrings(rv)
}
// Inserts the selected completion into the text at the appropriate location.
//
// word: The text to insert, including the matched partial word and its potential
// completion.
//
// charRange: The range of characters of the matched partial word to be completed.
//
// movement: The direction of movement. For possible values see the [NSText] Constants
// section. This value allows subclasses to distinguish between canceling
// completion and selection by arrow keys, by return, by tab, or by other
// means such as clicking.
//
// flag: [false] while the user navigates through the potential text completions,
// [true] when a completion is definitively selected or cancelled and the
// original value is reinserted.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method has two effects, text substitution and changing of the
// selection:
// 
// - It replaces the text between `charRange.Start()` and the current
// insertion point with `word`. - If `flag` is [false] it changes the
// selection to be the last characters of `word` where is equal to `[word
// length]` minus `charRange.Length()`, that is, the potential completion. -
// If `flag` is [true] it makes the selection empty and puts the insertion
// point just after `word`.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/insertCompletion(_:forPartialWordRange:movement:isFinal:)
func (t NSTextView) InsertCompletionForPartialWordRangeMovementIsFinal(word string, charRange foundation.NSRange, movement int, flag bool) {
	objc.Send[objc.ID](t.ID, objc.Sel("insertCompletion:forPartialWordRange:movement:isFinal:"), objc.String(word), charRange, movement, flag)
}
// Performs the default text checking on the entire document.
//
// sender: The control sending the message. May be `nil`.
//
// # Discussion
// 
// Immediately performs the text checking and replaces the document content
// with the checked content.
// 
// The checks performed are specified by [EnabledTextCheckingTypes];
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/checkTextInDocument(_:)
func (t NSTextView) CheckTextInDocument(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("checkTextInDocument:"), sender)
}
// Performs the default text checking on the current selection.
//
// sender: The control sending the message. May be `nil`.
//
// # Discussion
// 
// Immediately performs the text checking and replaces the selection with the
// checked content.
// 
// The checks performed are specified by [EnabledTextCheckingTypes];
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/checkTextInSelection(_:)
func (t NSTextView) CheckTextInSelection(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("checkTextInSelection:"), sender)
}
// Check and replace the text in the range using the specified checking types
// and options.
//
// range: The range to check.
//
// checkingTypes: The type of checking to be performed, passed by-reference. The possible
// constants are listed in [NSTextCheckingTypes] and can be combined using the
// C bit-wise [OR] operator to perform multiple checks at the same time.
// //
// [NSTextCheckingTypes]: https://developer.apple.com/documentation/Foundation/NSTextCheckingTypes
//
// options: A dictionary of values used during the checking process to perform. See
// Spell Checking Option Dictionary Keys for the supported values.
//
// # Discussion
// 
// This method calls the delegate method
// [TextViewWillCheckTextInRangeOptionsTypes] allowing you to modify the
// parameters before the checking occurs.
// 
// This method usually would not be called directly, since [NSTextView] itself
// will call it as needed, but it can be overridden.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/checkText(in:types:options:)
func (t NSTextView) CheckTextInRangeTypesOptions(range_ foundation.NSRange, checkingTypes uint64, options foundation.INSDictionary) {
	objc.Send[objc.ID](t.ID, objc.Sel("checkTextInRange:types:options:"), range_, checkingTypes, options)
}
// Handles the text checking results returned by the text view
//
// results: An array of [NSTextCheckingResult] objects.
// //
// [NSTextCheckingResult]: https://developer.apple.com/documentation/Foundation/NSTextCheckingResult
//
// range: The range of text that was checked.
//
// checkingTypes: The type of checking performed. The possible constants are listed in
// [NSTextCheckingTypes] and can be combined using the C bit-wise [OR]
// operator to perform multiple checks at the same time.
// //
// [NSTextCheckingTypes]: https://developer.apple.com/documentation/Foundation/NSTextCheckingTypes
//
// options: The dictionary of values used during the checking process to perform. See
// Spell Checking Option Dictionary Keys for the supported values.
//
// orthography: The orthography of the checked text.
//
// wordCount: The number of words.
//
// # Discussion
// 
// The [NSTextViewDelegate] offers a method,
// [TextViewDidCheckTextInRangeTypesOptionsResultsOrthographyWordCount] that
// is called after the checking is performed, allowing you to modify the
// results.
// 
// This method usually would not be called directly, since [NSTextView] itself
// will call it as needed, but it can be overridden.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/handleTextCheckingResults(_:forRange:types:options:orthography:wordCount:)
func (t NSTextView) HandleTextCheckingResultsForRangeTypesOptionsOrthographyWordCount(results []foundation.NSTextCheckingResult, range_ foundation.NSRange, checkingTypes uint64, options foundation.INSDictionary, orthography foundation.NSOrthography, wordCount int) {
	objc.Send[objc.ID](t.ID, objc.Sel("handleTextCheckingResults:forRange:types:options:orthography:wordCount:"), objectivec.IObjectSliceToNSArray(results), range_, checkingTypes, options, orthography, wordCount)
}
// Toggles the state of the automatic dash substitution.
//
// sender: The control sending the message. May be `nil`.
//
// # Discussion
// 
// Turning on automatic dash substitution enables automatic conversion of
// sequences of ASCII hyphen (`-`) characters to typographic dashes.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/toggleAutomaticDashSubstitution(_:)
func (t NSTextView) ToggleAutomaticDashSubstitution(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("toggleAutomaticDashSubstitution:"), sender)
}
// Toggles the state of the automatic data detection.
//
// sender: The control sending the message. May be `nil`.
//
// # Discussion
// 
// Automatic data detection enables detection of dates, addresses, and phone
// numbers.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/toggleAutomaticDataDetection(_:)
func (t NSTextView) ToggleAutomaticDataDetection(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("toggleAutomaticDataDetection:"), sender)
}
// Toggles the state of the automatic spelling correction.
//
// sender: The control sending the message. May be `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/toggleAutomaticSpellingCorrection(_:)
func (t NSTextView) ToggleAutomaticSpellingCorrection(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("toggleAutomaticSpellingCorrection:"), sender)
}
// Toggles the state of the automatic text replacement.
//
// sender: The control sending the message. May be `nil`.
//
// # Discussion
// 
// Turning on automatic text replacement enables automatic substitution of a
// variety of static text items based on user preferences.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/toggleAutomaticTextReplacement(_:)
func (t NSTextView) ToggleAutomaticTextReplacement(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("toggleAutomaticTextReplacement:"), sender)
}
// Replaces text in the range you specify with the attributed string you
// provide.
//
// range: The range of the replacement.
//
// attributedString: The attributed string to use as the replacement text.
//
// # Return Value
// 
// Retuns `true` if the replacement was successful, `false` otherwise.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/performValidatedReplacement(in:with:)
func (t NSTextView) PerformValidatedReplacementInRangeWithAttributedString(range_ foundation.NSRange, attributedString foundation.NSAttributedString) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("performValidatedReplacementInRange:withAttributedString:"), range_, attributedString)
	return rv
}
// Notifies the QuickLook panel that an update may be required.
//
// # Discussion
// 
// Notifies the [QLPreviewPanel] class for possible status changes with the
// data source or controller. Typically invoked in response to selection
// changes.
//
// [QLPreviewPanel]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewPanel
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/updateQuickLookPreviewPanel()
func (t NSTextView) UpdateQuickLookPreviewPanel() {
	objc.Send[objc.ID](t.ID, objc.Sel("updateQuickLookPreviewPanel"))
}
// An action message that toggles the visibility state of the Quick Look
// preview panel.
//
// sender: The message sender.
//
// # Discussion
// 
// This action message toggles the visibility state of the Quick Look preview
// panel if the receiver is the current Quick Look controller.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/toggleQuickLookPreviewPanel(_:)
func (t NSTextView) ToggleQuickLookPreviewPanel(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("toggleQuickLookPreviewPanel:"), sender)
}
// Returns an array of URLs for items that can be displayed by QuickLook in
// the specified ranges.
//
// ranges: An array of ranges.
//
// # Return Value
// 
// Returns an array of document URLs for text attachment content, if
// available.
//
// # Discussion
// 
// Each preview item must conform to the [QLPreviewItem] protocol.
//
// [QLPreviewItem]: https://developer.apple.com/documentation/QuickLookUI/QLPreviewItem
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/quickLookPreviewableItems(inRanges:)
func (t NSTextView) QuickLookPreviewableItemsInRanges(ranges []foundation.NSValue) []objectivec.IObject {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("quickLookPreviewableItemsInRanges:"), objectivec.IObjectSliceToNSArray(ranges))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
// An action method that sets the layout orientation of the text.
//
// sender: The sender.
//
// # Discussion
// 
// Calls [SetLayoutOrientation] with the sender’s tag as the orientation.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/changeLayoutOrientation(_:)
func (t NSTextView) ChangeLayoutOrientation(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("changeLayoutOrientation:"), sender)
}
// Changes the receiver’s layout orientation and invalidates the contents.
//
// orientation: The text layout orientation.
//
// # Discussion
// 
// Unlike other [NSTextView] properties, this is not shared by sibling views.
// It also rotates the bounds 90 degrees, swaps horizontal and vertical bits
// of the [autoresizingMask] mask, and reconfigures [HorizontallyResizable]
// and [VerticallyResizable] properties accordingly. Also, if
// [enclosingScrollView] returns non-`nil`, it reconfigures the horizontal and
// vertical ruler views, the horizontal and vertical scrollers, and the frame.
//
// [autoresizingMask]: https://developer.apple.com/documentation/AppKit/NSView/autoresizingMask-swift.property
// [enclosingScrollView]: https://developer.apple.com/documentation/AppKit/NSView/enclosingScrollView
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/setLayoutOrientation(_:)
func (t NSTextView) SetLayoutOrientation(orientation NSTextLayoutOrientation) {
	objc.Send[objc.ID](t.ID, objc.Sel("setLayoutOrientation:"), orientation)
}
// See: https://developer.apple.com/documentation/AppKit/NSTextView/updateTextTouchBarItems()
func (t NSTextView) UpdateTextTouchBarItems() {
	objc.Send[objc.ID](t.ID, objc.Sel("updateTextTouchBarItems"))
}
// See: https://developer.apple.com/documentation/AppKit/NSTextView/updateTouchBarItemIdentifiers()
func (t NSTextView) UpdateTouchBarItemIdentifiers() {
	objc.Send[objc.ID](t.ID, objc.Sel("updateTouchBarItemIdentifiers"))
}
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/drawTextHighlightBackground(for:origin:)
func (t NSTextView) DrawTextHighlightBackgroundForTextRangeOrigin(textRange INSTextRange, origin corefoundation.CGPoint) {
	objc.Send[objc.ID](t.ID, objc.Sel("drawTextHighlightBackgroundForTextRange:origin:"), textRange, origin)
}
// An action for toggling [NSTextHighlightStyleAttributeName] in the
// receiver’s selected range. The sender should be a menu item with a
// `representedObject` of type ([NSTextHighlightColorScheme]).
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/highlight(_:)
func (t NSTextView) Highlight(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("highlight:"), sender)
}
// Returns the attributed substring for the specified range of characters.
//
// range: The range of characters.
//
// # Return Value
// 
// An attributed string representing the specified characters.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityStaticText/accessibilityAttributedString(for:)
func (t NSTextView) AccessibilityAttributedStringForRange(range_ foundation.NSRange) foundation.NSAttributedString {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("accessibilityAttributedStringForRange:"), range_)
	return foundation.NSAttributedStringFromID(rv)
}
// Returns the rectangle that encloses the specified range of characters.
//
// range: The range of characters.
//
// # Return Value
// 
// The rectangle that encloses the specified characters.
//
// # Discussion
// 
// If the range crosses a line boundary, the returned rectangle will fully
// enclose all the lines of characters.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityNavigableStaticText/accessibilityFrame(for:)
func (t NSTextView) AccessibilityFrameForRange(range_ foundation.NSRange) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](t.ID, objc.Sel("accessibilityFrameForRange:"), range_)
	return corefoundation.CGRect(rv)
}
// Returns the line number for the line that contains the specified character
// index.
//
// index: The index for a character.
//
// # Return Value
// 
// The line number for the line holding the specified character index.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityNavigableStaticText/accessibilityLine(for:)
func (t NSTextView) AccessibilityLineForIndex(index int) int {
	rv := objc.Send[int](t.ID, objc.Sel("accessibilityLineForIndex:"), index)
	return rv
}
// Returns the range of characters in the specified line.
//
// lineNumber: The line number to be examined.
//
// # Return Value
// 
// The range of characters for the specified line number. If the line ends
// with a newline character, including the newline is preferred.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityNavigableStaticText/accessibilityRange(forLine:)
func (t NSTextView) AccessibilityRangeForLine(lineNumber int) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](t.ID, objc.Sel("accessibilityRangeForLine:"), lineNumber)
	return foundation.NSRange(rv)
}
// Returns the substring for the specified range.
//
// range: A range of characters contained by this element.
//
// # Return Value
// 
// The substring specified by the given range.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityNavigableStaticText/accessibilityString(for:)
func (t NSTextView) AccessibilityStringForRange(range_ foundation.NSRange) string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("accessibilityStringForRange:"), range_)
	return foundation.NSStringFromID(rv).String()
}
// Returns the text that the accessibility element displays.
//
// # Return Value
// 
// The text displayed by the element.
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityValue] property.
//
// [accessibilityValue]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityValue
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityStaticText/accessibilityValue()
func (t NSTextView) AccessibilityValue() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("accessibilityValue"))
	return foundation.NSStringFromID(rv).String()
}
// Returns the range of visible characters in the document.
//
// # Return Value
// 
// The range of the visible characters in the document. This method should
// return the range for entire lines. Characters that are horizontally clipped
// are included in this range.
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityVisibleCharacterRange] property.
//
// [accessibilityVisibleCharacterRange]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityVisibleCharacterRange
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityStaticText/accessibilityVisibleCharacterRange()
func (t NSTextView) AccessibilityVisibleCharacterRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](t.ID, objc.Sel("accessibilityVisibleCharacterRange"))
	return foundation.NSRange(rv)
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
func (t NSTextView) AttributedString() foundation.NSAttributedString {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("attributedString"))
	return foundation.NSAttributedStringFromID(rv)
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
func (t NSTextView) AttributedSubstringForProposedRangeActualRange(range_ foundation.NSRange, actualRange foundation.NSRange) foundation.NSAttributedString {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("attributedSubstringForProposedRange:actualRange:"), range_, actualRange)
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
func (t NSTextView) BaselineDeltaForCharacterAtIndex(anIndex uint) float64 {
	rv := objc.Send[float64](t.ID, objc.Sel("baselineDeltaForCharacterAtIndex:"), anIndex)
	return rv
}
// Tells the delegate that the user has started touching one of the candidates
// in the candidate list item.
//
// anItem: The candidate list bar item that the user is interacting with.
//
// index: The index of the candidate that the user is currently touching.
//
// See: https://developer.apple.com/documentation/AppKit/NSCandidateListTouchBarItemDelegate/candidateListTouchBarItem(_:beginSelectingCandidateAt:)
func (t NSTextView) CandidateListTouchBarItemBeginSelectingCandidateAtIndex(anItem INSCandidateListTouchBarItem, index int) {
	objc.Send[objc.ID](t.ID, objc.Sel("candidateListTouchBarItem:beginSelectingCandidateAtIndex:"), anItem, index)
}
// Tells the delegate that user has moved from touching one candidate in the
// candidate list item to another.
//
// anItem: The candidate list item that the user is interacting with.
//
// previousIndex: The index of the candidate that the user was previously touching.
//
// index: The index of the candidate that the user is currently touching.
//
// See: https://developer.apple.com/documentation/AppKit/NSCandidateListTouchBarItemDelegate/candidateListTouchBarItem(_:changeSelectionFromCandidateAt:to:)
func (t NSTextView) CandidateListTouchBarItemChangeSelectionFromCandidateAtIndexToIndex(anItem INSCandidateListTouchBarItem, previousIndex int, index int) {
	objc.Send[objc.ID](t.ID, objc.Sel("candidateListTouchBarItem:changeSelectionFromCandidateAtIndex:toIndex:"), anItem, previousIndex, index)
}
// Tells the delegate that the visibility of the candidate list has changed.
//
// anItem: The candidate list item whose candidate list’s visibility has changed.
//
// isVisible: A Boolean value that specifies whether or not the candidate list is
// visible. If [true] then the candidate list is visible, [false] otherwise.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSCandidateListTouchBarItemDelegate/candidateListTouchBarItem(_:changedCandidateListVisibility:)
func (t NSTextView) CandidateListTouchBarItemChangedCandidateListVisibility(anItem INSCandidateListTouchBarItem, isVisible bool) {
	objc.Send[objc.ID](t.ID, objc.Sel("candidateListTouchBarItem:changedCandidateListVisibility:"), anItem, isVisible)
}
// Tells the delegate that a user has stopped touching candidates in the
// candidate list item.
//
// anItem: The candidate list item that the user is interacting with.
//
// index: The index of the candidate that the user was touching when they lifted
// their finger.
//
// # Discussion
// 
// If `index` is equal to [NSNotFound] then the user didn’t select a
// candidate.
//
// See: https://developer.apple.com/documentation/AppKit/NSCandidateListTouchBarItemDelegate/candidateListTouchBarItem(_:endSelectingCandidateAt:)
func (t NSTextView) CandidateListTouchBarItemEndSelectingCandidateAtIndex(anItem INSCandidateListTouchBarItem, index int) {
	objc.Send[objc.ID](t.ID, objc.Sel("candidateListTouchBarItem:endSelectingCandidateAtIndex:"), anItem, index)
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
func (t NSTextView) CharacterIndexForPoint(point corefoundation.CGPoint) uint {
	rv := objc.Send[uint](t.ID, objc.Sel("characterIndexForPoint:"), point)
	return rv
}
// Invoked when the dragging session has completed.
//
// session: The dragging session.
//
// screenPoint: The point where the drag ended, in screen coordinates.
//
// operation: The drag operation. See [NSDragOperation] for drag operation types.
// //
// [NSDragOperation]: https://developer.apple.com/documentation/AppKit/NSDragOperation
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingSource/draggingSession(_:endedAt:operation:)
func (t NSTextView) DraggingSessionEndedAtPointOperation(session INSDraggingSession, screenPoint corefoundation.CGPoint, operation NSDragOperation) {
	objc.Send[objc.ID](t.ID, objc.Sel("draggingSession:endedAtPoint:operation:"), session, screenPoint, operation)
}
// Invoked when the drag moves on the screen.
//
// session: The dragging session.
//
// screenPoint: The point where the drag moved to, in screen coordinates.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingSource/draggingSession(_:movedTo:)
func (t NSTextView) DraggingSessionMovedToPoint(session INSDraggingSession, screenPoint corefoundation.CGPoint) {
	objc.Send[objc.ID](t.ID, objc.Sel("draggingSession:movedToPoint:"), session, screenPoint)
}
// Declares the types of operations the source allows to be performed.
//
// session: The dragging session.
//
// context: The dragging context. See [NSDraggingContext] for the supported values.
// //
// [NSDraggingContext]: https://developer.apple.com/documentation/AppKit/NSDraggingContext
//
// # Return Value
// 
// The appropriate dragging operation as defined in
//
// # Discussion
// 
// In the future Apple may provide more specific “within” values in the
// future. To account for this, for unrecognized localities, return the
// operation mask for the most specific context that you are concerned with.
// The following code is an example of how to implement this functionality:
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingSource/draggingSession(_:sourceOperationMaskFor:)
func (t NSTextView) DraggingSessionSourceOperationMaskForDraggingContext(session INSDraggingSession, context NSDraggingContext) NSDragOperation {
	rv := objc.Send[NSDragOperation](t.ID, objc.Sel("draggingSession:sourceOperationMaskForDraggingContext:"), session, context)
	return NSDragOperation(rv)
}
// Invoked when the drag will begin.
//
// session: The dragging session.
//
// screenPoint: The point where the drag will begin, in screen coordinates.
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingSource/draggingSession(_:willBeginAt:)
func (t NSTextView) DraggingSessionWillBeginAtPoint(session INSDraggingSession, screenPoint corefoundation.CGPoint) {
	objc.Send[objc.ID](t.ID, objc.Sel("draggingSession:willBeginAtPoint:"), session, screenPoint)
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
func (t NSTextView) DrawsVerticallyForCharacterAtIndex(charIndex uint) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("drawsVerticallyForCharacterAtIndex:"), charIndex)
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
func (t NSTextView) FirstRectForCharacterRangeActualRange(range_ foundation.NSRange, actualRange foundation.NSRange) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](t.ID, objc.Sel("firstRectForCharacterRange:actualRange:"), range_, actualRange)
	return corefoundation.CGRect(rv)
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
func (t NSTextView) FractionOfDistanceThroughGlyphForPoint(point corefoundation.CGPoint) float64 {
	rv := objc.Send[float64](t.ID, objc.Sel("fractionOfDistanceThroughGlyphForPoint:"), point)
	return rv
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
func (t NSTextView) HasMarkedText() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("hasMarkedText"))
	return rv
}
// Returns whether the modifier keys will be ignored for this dragging
// session.
//
// session: The dragging session.
//
// # Return Value
// 
// [true] if the modifier keys will be ignored, [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSDraggingSource/ignoreModifierKeys(for:)
func (t NSTextView) IgnoreModifierKeysForDraggingSession(session INSDraggingSession) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("ignoreModifierKeysForDraggingSession:"), session)
	return rv
}
// Inserts an adaptive image into the text at the specifed location.
//
// adaptiveImageGlyph: The adaptive image to add to the text.
//
// replacementRange: The text range at which to insert the image.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextInputClient/insert(_:replacementRange:)
func (t NSTextView) InsertAdaptiveImageGlyphReplacementRange(adaptiveImageGlyph INSAdaptiveImageGlyph, replacementRange foundation.NSRange) {
	objc.Send[objc.ID](t.ID, objc.Sel("insertAdaptiveImageGlyph:replacementRange:"), adaptiveImageGlyph, replacementRange)
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
func (t NSTextView) InsertTextReplacementRange(string_ objectivec.IObject, replacementRange foundation.NSRange) {
	objc.Send[objc.ID](t.ID, objc.Sel("insertText:replacementRange:"), string_, replacementRange)
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
func (t NSTextView) MarkedRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](t.ID, objc.Sel("markedRange"))
	return foundation.NSRange(rv)
}
// See: https://developer.apple.com/documentation/AppKit/NSTextInputClient/preferredTextAccessoryPlacement()
func (t NSTextView) PreferredTextAccessoryPlacement() NSTextCursorAccessoryPlacement {
	rv := objc.Send[NSTextCursorAccessoryPlacement](t.ID, objc.Sel("preferredTextAccessoryPlacement"))
	return NSTextCursorAccessoryPlacement(rv)
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
func (t NSTextView) SetMarkedTextSelectedRangeReplacementRange(string_ objectivec.IObject, selectedRange foundation.NSRange, replacementRange foundation.NSRange) {
	objc.Send[objc.ID](t.ID, objc.Sel("setMarkedText:selectedRange:replacementRange:"), string_, selectedRange, replacementRange)
}
// Asks the delegate object for the bar item for the specified bar and item
// identifier.
//
// touchBar: The bar that’s requesting the bar item.
//
// identifier: The item identifier associated with the item being requested.
//
// # Return Value
// 
// A fully initialized bar item for the specified bar and identifier.
//
// # Discussion
// 
// When the system needs to populate a bar’s items array, the system calls
// this delegate method to retrieve an item if that item can’t be found in
// the bar’s private array or in the bar’s [TemplateItems] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBarDelegate/touchBar(_:makeItemForIdentifier:)
func (t NSTextView) TouchBarMakeItemForIdentifier(touchBar INSTouchBar, identifier NSTouchBarItemIdentifier) INSTouchBarItem {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("touchBar:makeItemForIdentifier:"), touchBar, objc.String(string(identifier)))
	return NSTouchBarItemFromID(rv)
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
func (t NSTextView) UnmarkText() {
	objc.Send[objc.ID](t.ID, objc.Sel("unmarkText"))
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
func (t NSTextView) ValidAttributesForMarkedText() []string {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("validAttributesForMarkedText"))
	return objc.ConvertSliceToStrings(rv)
}
// Implemented to override the default action of enabling or disabling a
// specific menu item.
//
// menuItem: An [NSMenuItem] object that represents the menu item.
//
// # Return Value
// 
// [true] to enable `menuItem`, [false] to disable it.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The object implementing this method must be the target of `menuItem`. You
// can determine which menu item `menuItem` is by querying it for its tag or
// action.
// 
// The following example disables the menu item associated with the
// `nextRecord` action method when the selected line in a table view is the
// last one; conversely, it disables the menu item with `priorRecord` as its
// action method when the selected row is the first one in the table view.
// (The `countryOrRegionKeys` array contains names that appear in the table
// view.)
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItemValidation/validateMenuItem(_:)
func (t NSTextView) ValidateMenuItem(menuItem INSMenuItem) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("validateMenuItem:"), menuItem)
	return rv
}
// Returns a Boolean value that indicates whether the sender should be
// enabled.
//
// item: The user interface item to validate. You can send `anItem` the [Action] and
// [Tag] messages.
//
// # Return Value
// 
// [true] if the user interface item should be enabled, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceValidations/validateUserInterfaceItem(_:)
func (t NSTextView) ValidateUserInterfaceItem(item NSValidatedUserInterfaceItem) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("validateUserInterfaceItem:"), item)
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
func (t NSTextView) WindowLevel() int {
	rv := objc.Send[int](t.ID, objc.Sel("windowLevel"))
	return rv
}

// Registers send and return types for the Services facility.
//
// # Discussion
// 
// This method is invoked automatically when the first instance of a text view
// is created; you should never need to invoke it directly.
// 
// Subclasses of [NSTextView] that wish to add support for new service types
// should override [RegisterForServices] to call `super` and then register
// their own new types.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/registerForServices()
func (_NSTextViewClass NSTextViewClass) RegisterForServices() {
	objc.Send[objc.ID](objc.ID(_NSTextViewClass.class), objc.Sel("registerForServices"))
}
// See: https://developer.apple.com/documentation/AppKit/NSTextView/scrollableDocumentContentTextView()
func (_NSTextViewClass NSTextViewClass) ScrollableDocumentContentTextView() NSScrollView {
	rv := objc.Send[objc.ID](objc.ID(_NSTextViewClass.class), objc.Sel("scrollableDocumentContentTextView"))
	return NSScrollViewFromID(rv)
}
// See: https://developer.apple.com/documentation/AppKit/NSTextView/scrollablePlainDocumentContentTextView()
func (_NSTextViewClass NSTextViewClass) ScrollablePlainDocumentContentTextView() NSScrollView {
	rv := objc.Send[objc.ID](objc.ID(_NSTextViewClass.class), objc.Sel("scrollablePlainDocumentContentTextView"))
	return NSScrollViewFromID(rv)
}
// See: https://developer.apple.com/documentation/AppKit/NSTextView/scrollableTextView()
func (_NSTextViewClass NSTextViewClass) ScrollableTextView() NSScrollView {
	rv := objc.Send[objc.ID](objc.ID(_NSTextViewClass.class), objc.Sel("scrollableTextView"))
	return NSScrollViewFromID(rv)
}

// The receiver’s text container.
//
// # Discussion
// 
// The receiver uses the layout manager and text storage of `aTextContainer`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/textContainer
func (t NSTextView) TextContainer() INSTextContainer {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("textContainer"))
	return NSTextContainerFromID(objc.ID(rv))
}
func (t NSTextView) SetTextContainer(value INSTextContainer) {
	objc.Send[struct{}](t.ID, objc.Sel("setTextContainer:"), value)
}
// The empty space the receiver leaves around its associated text container.
//
// # Discussion
// 
// It is possible to set the text container and view sizes and resizing
// behavior so that the inset cannot be maintained exactly, although the text
// system tries to maintain the inset wherever possible. In any case, the
// [TextContainerOrigin] and size of the text container are authoritative as
// to the location of the text container within the view.
// 
// The text itself can have an additional inset, inside the text container,
// specified by the [LineFragmentPadding] method of [NSTextContainer].
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/textContainerInset
func (t NSTextView) TextContainerInset() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](t.ID, objc.Sel("textContainerInset"))
	return corefoundation.CGSize(rv)
}
func (t NSTextView) SetTextContainerInset(value corefoundation.CGSize) {
	objc.Send[struct{}](t.ID, objc.Sel("setTextContainerInset:"), value)
}
// The origin of the receiver’s text container.
//
// # Discussion
// 
// Calculated from the receiver’s bounds rectangle, container inset, and the
// container’s used rect.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/textContainerOrigin
func (t NSTextView) TextContainerOrigin() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](t.ID, objc.Sel("textContainerOrigin"))
	return corefoundation.CGPoint(rv)
}
// The manager that lays out text for the receiver’s text container.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/textLayoutManager
func (t NSTextView) TextLayoutManager() INSTextLayoutManager {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("textLayoutManager"))
	return NSTextLayoutManagerFromID(objc.ID(rv))
}
// The layout manager that lays out text for the receiver’s text container.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/layoutManager
func (t NSTextView) LayoutManager() INSLayoutManager {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("layoutManager"))
	return NSLayoutManagerFromID(objc.ID(rv))
}
// The receiver’s text storage object.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/textContentStorage
func (t NSTextView) TextContentStorage() INSTextContentStorage {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("textContentStorage"))
	return NSTextContentStorageFromID(objc.ID(rv))
}
// The receiver’s text storage object.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/textStorage
func (t NSTextView) TextStorage() NSTextStorage {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("textStorage"))
	return NSTextStorageFromID(objc.ID(rv))
}
// A Boolean value that indicates whether the receiver allows its background
// color to change.
//
// # Discussion
// 
// [true] if the receiver allows the background color to change, otherwise
// [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/allowsDocumentBackgroundColorChange
func (t NSTextView) AllowsDocumentBackgroundColorChange() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("allowsDocumentBackgroundColorChange"))
	return rv
}
func (t NSTextView) SetAllowsDocumentBackgroundColorChange(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setAllowsDocumentBackgroundColorChange:"), value)
}
// A Boolean value that determines whether the receiver should draw its
// insertion point.
//
// # Discussion
// 
// [true] if the receiver should draw its insertion point, [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/shouldDrawInsertionPoint
func (t NSTextView) ShouldDrawInsertionPoint() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("shouldDrawInsertionPoint"))
	return rv
}
// An array of locale identifiers representing input sources that are allowed
// to be enabled when the receiver has the keyboard focus.
//
// # Discussion
// 
// You can use the meta-locale identifier,
// [NSAllRomanInputSourcesLocaleIdentifier], to specify input sources that are
// limited for Roman script editing.
//
// [NSAllRomanInputSourcesLocaleIdentifier]: https://developer.apple.com/documentation/AppKit/NSAllRomanInputSourcesLocaleIdentifier
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/allowedInputSourceLocales
func (t NSTextView) AllowedInputSourceLocales() []string {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("allowedInputSourceLocales"))
	return objc.ConvertSliceToStrings(rv)
}
func (t NSTextView) SetAllowedInputSourceLocales(value []string) {
	objc.Send[struct{}](t.ID, objc.Sel("setAllowedInputSourceLocales:"), objectivec.StringSliceToNSArray(value))
}
// A Boolean value that indicates whether the receiver allows undo.
//
// # Discussion
// 
// [true] if the receiver allows undo, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/allowsUndo
func (t NSTextView) AllowsUndo() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("allowsUndo"))
	return rv
}
func (t NSTextView) SetAllowsUndo(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setAllowsUndo:"), value)
}
// The receiver’s default paragraph style.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/defaultParagraphStyle
func (t NSTextView) DefaultParagraphStyle() INSParagraphStyle {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("defaultParagraphStyle"))
	return NSParagraphStyleFromID(objc.ID(rv))
}
func (t NSTextView) SetDefaultParagraphStyle(value INSParagraphStyle) {
	objc.Send[struct{}](t.ID, objc.Sel("setDefaultParagraphStyle:"), value)
}
// Indicates whether image attachments should permit editing of their images.
//
// # Discussion
// 
// [true] if image editing is allowed; otherwise, [false].
// 
// For image editing to be allowed, the text view must be editable and the
// text attachment cell must support image editing.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/allowsImageEditing
func (t NSTextView) AllowsImageEditing() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("allowsImageEditing"))
	return rv
}
func (t NSTextView) SetAllowsImageEditing(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setAllowsImageEditing:"), value)
}
// A Boolean value that enables and disables automatic quotation mark
// substitution.
//
// # Discussion
// 
// If [true], automatic quotation mark substitution is enabled; if [false], it
// is disabled.
// 
// Automatic quote substitution causes ASCII quotation marks and apostrophes
// to be automatically replaced, on a context-dependent basis, with more
// typographically accurate symbols.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/isAutomaticQuoteSubstitutionEnabled
func (t NSTextView) AutomaticQuoteSubstitutionEnabled() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isAutomaticQuoteSubstitutionEnabled"))
	return rv
}
func (t NSTextView) SetAutomaticQuoteSubstitutionEnabled(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setAutomaticQuoteSubstitutionEnabled:"), value)
}
// A Boolean value that enables or disables automatic link detection.
//
// # Discussion
// 
// If [true], automatic link detection is enabled; if [false], it is disabled.
// 
// Automatic link detection causes strings representing URLs typed in the view
// to be automatically made into links to those URLs.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/isAutomaticLinkDetectionEnabled
func (t NSTextView) AutomaticLinkDetectionEnabled() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isAutomaticLinkDetectionEnabled"))
	return rv
}
func (t NSTextView) SetAutomaticLinkDetectionEnabled(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setAutomaticLinkDetectionEnabled:"), value)
}
// A Boolean value that indicates whether the text view automatically supplies
// the destination of a link as a tooltip for text that has a link attribute.
//
// # Discussion
// 
// [true] if link tooltips are automatically displayed; otherwise, [false].
// 
// The default value for this feature is [true]; clients who do not wish
// tooltips to be displayed automatically must explicitly disable it.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/displaysLinkToolTips
func (t NSTextView) DisplaysLinkToolTips() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("displaysLinkToolTips"))
	return rv
}
func (t NSTextView) SetDisplaysLinkToolTips(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setDisplaysLinkToolTips:"), value)
}
// A Boolean value that indicates whether the text view supplies
// autocompletion suggestions as the user types.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/isAutomaticTextCompletionEnabled
func (t NSTextView) AutomaticTextCompletionEnabled() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isAutomaticTextCompletionEnabled"))
	return rv
}
func (t NSTextView) SetAutomaticTextCompletionEnabled(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setAutomaticTextCompletionEnabled:"), value)
}
// A Boolean value that indicates whether the framework should use adaptive
// color mapping for dark appearance.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/usesAdaptiveColorMappingForDarkAppearance
func (t NSTextView) UsesAdaptiveColorMappingForDarkAppearance() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("usesAdaptiveColorMappingForDarkAppearance"))
	return rv
}
func (t NSTextView) SetUsesAdaptiveColorMappingForDarkAppearance(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setUsesAdaptiveColorMappingForDarkAppearance:"), value)
}
// See: https://developer.apple.com/documentation/AppKit/NSTextView/usesRolloverButtonForSelection
func (t NSTextView) UsesRolloverButtonForSelection() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("usesRolloverButtonForSelection"))
	return rv
}
func (t NSTextView) SetUsesRolloverButtonForSelection(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setUsesRolloverButtonForSelection:"), value)
}
// A Boolean value that controls whether the text views sharing the
// receiver’s layout manager use a ruler.
//
// # Discussion
// 
// [true] to cause text views sharing the receiver’s layout manager to
// respond to [NSRulerView] client messages and to paragraph-related menu
// actions, and update the ruler (when visible) as the selection changes with
// its paragraph and tab attributes, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/usesRuler
func (t NSTextView) UsesRuler() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("usesRuler"))
	return rv
}
func (t NSTextView) SetUsesRuler(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setUsesRuler:"), value)
}
// A Boolean value that indicates whether this text view uses the inspector
// bar.
//
// # Discussion
// 
// The inspector bar displays text formatting controls, much like those in
// iWork applications, which can be used in place of the formatting controls
// in the ruler accessory view.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/usesInspectorBar
func (t NSTextView) UsesInspectorBar() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("usesInspectorBar"))
	return rv
}
func (t NSTextView) SetUsesInspectorBar(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setUsesInspectorBar:"), value)
}
// An array containing the ranges of characters selected in the receiver’s
// layout manager.
//
// # Discussion
// 
// The objects in the array are sorted, non-overlapping, non-contiguous, and
// (except for the case of a single range) have non-zero-length
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/selectedRanges
func (t NSTextView) SelectedRanges() []foundation.NSValue {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("selectedRanges"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSValue {
		return foundation.NSValueFromID(id)
	})
}
func (t NSTextView) SetSelectedRanges(value []foundation.NSValue) {
	objc.Send[struct{}](t.ID, objc.Sel("setSelectedRanges:"), objectivec.IObjectSliceToNSArray(value))
}
// The preferred direction of selection.
//
// # Discussion
// 
// Selection affinity determines whether, for example, the insertion point
// appears after the last character on a line or before the first character on
// the following line in cases where text wraps across line boundaries.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/selectionAffinity
func (t NSTextView) SelectionAffinity() NSSelectionAffinity {
	rv := objc.Send[NSSelectionAffinity](t.ID, objc.Sel("selectionAffinity"))
	return NSSelectionAffinity(rv)
}
// The selection granularity for subsequent extension of a selection.
//
// # Discussion
// 
// Selection granularity is used to determine how the selection is modified
// when the user Shift-clicks or drags the mouse after a double or triple
// click. For example, if the user selects a word by double-clicking, the
// selection granularity is set to [NSSelectByWord]. Subsequent Shift-clicks
// then extend the selection by words.
// 
// Selection granularity is reset to [NSSelectByCharacter] whenever the
// selection is set. You should always set the selection granularity after
// setting the selection.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/selectionGranularity
func (t NSTextView) SelectionGranularity() NSSelectionGranularity {
	rv := objc.Send[NSSelectionGranularity](t.ID, objc.Sel("selectionGranularity"))
	return NSSelectionGranularity(rv)
}
func (t NSTextView) SetSelectionGranularity(value NSSelectionGranularity) {
	objc.Send[struct{}](t.ID, objc.Sel("setSelectionGranularity:"), value)
}
// The color of the insertion point.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/insertionPointColor
func (t NSTextView) InsertionPointColor() INSColor {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("insertionPointColor"))
	return NSColorFromID(objc.ID(rv))
}
func (t NSTextView) SetInsertionPointColor(value INSColor) {
	objc.Send[struct{}](t.ID, objc.Sel("setInsertionPointColor:"), value)
}
// The attributes used to indicate the selection.
//
// # Discussion
// 
// Text color, background color, and underline are the only supported
// attributes for selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/selectedTextAttributes
func (t NSTextView) SelectedTextAttributes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("selectedTextAttributes"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (t NSTextView) SetSelectedTextAttributes(value foundation.INSDictionary) {
	objc.Send[struct{}](t.ID, objc.Sel("setSelectedTextAttributes:"), value)
}
// The attributes used to draw marked text.
//
// # Discussion
// 
// Text color, background color, and underline are the only supported
// attributes for marked text.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/markedTextAttributes
func (t NSTextView) MarkedTextAttributes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("markedTextAttributes"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (t NSTextView) SetMarkedTextAttributes(value foundation.INSDictionary) {
	objc.Send[struct{}](t.ID, objc.Sel("setMarkedTextAttributes:"), value)
}
// The attributes used to draw the onscreen presentation of link text.
//
// # Discussion
// 
// Link text attributes are applied as temporary attributes to any text with a
// link attribute. Candidates include those attributes that do not affect
// layout.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/linkTextAttributes
func (t NSTextView) LinkTextAttributes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("linkTextAttributes"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (t NSTextView) SetLinkTextAttributes(value foundation.INSDictionary) {
	objc.Send[struct{}](t.ID, objc.Sel("setLinkTextAttributes:"), value)
}
// The types this text view can read immediately from the pasteboard.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/readablePasteboardTypes
func (t NSTextView) ReadablePasteboardTypes() []string {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("readablePasteboardTypes"))
	return objc.ConvertSliceToStrings(rv)
}
// The pasteboard types that can be provided from the current selection.
//
// # Discussion
// 
// An array of strings describing the types that can be written to the
// pasteboard immediately, or an array with no members if the text view has no
// text or no selection.
// 
// Overriders can copy the result from super and add their own new types.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/writablePasteboardTypes
func (t NSTextView) WritablePasteboardTypes() []string {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("writablePasteboardTypes"))
	return objc.ConvertSliceToStrings(rv)
}
// The receiver’s typing attributes.
//
// # Discussion
// 
// Typing attributes are reset automatically whenever the selection changes.
// However, if you add any user actions that change text attributes, the
// action should use this method to apply those attributes afterwards. User
// actions that change attributes should always set the typing attributes
// because there might not be a subsequent change in selection before the next
// typing.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/typingAttributes
func (t NSTextView) TypingAttributes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("typingAttributes"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (t NSTextView) SetTypingAttributes(value foundation.INSDictionary) {
	objc.Send[struct{}](t.ID, objc.Sel("setTypingAttributes:"), value)
}
// A Boolean value that indicates whether undo coalescing is in progress.
//
// # Discussion
// 
// [true] if undo coalescing is in progress, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/isCoalescingUndo
func (t NSTextView) CoalescingUndo() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isCoalescingUndo"))
	return rv
}
// The data types that the receiver accepts as the destination view of a
// dragging operation.
//
// # Discussion
// 
// These types are automatically registered as necessary by the text view.
// Subclasses should override this value as necessary to add their own types
// to those returned by [NSTextView]‘s implementation. They must then also
// override the appropriate methods of the [NSDraggingDestination] protocol to
// support import of those types. See that protocol’s specification for more
// information.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/acceptableDragTypes
func (t NSTextView) AcceptableDragTypes() []string {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("acceptableDragTypes"))
	return objc.ConvertSliceToStrings(rv)
}
// The range of characters affected by an action method that changes character
// (not paragraph) attributes.
//
// # Discussion
// 
// The range of characters affected by an action method that changes character
// (not paragraph) attributes, such as the NSText action method [ChangeFont].
// For rich text this range is typically the range of the selection. For plain
// text this range is the entire contents of the receiver. If the receiver
// isn’t editable or doesn’t use the Font panel, the range has a location
// of [NSNotFound].
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/rangeForUserCharacterAttributeChange
func (t NSTextView) RangeForUserCharacterAttributeChange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](t.ID, objc.Sel("rangeForUserCharacterAttributeChange"))
	return foundation.NSRange(rv)
}
// An array containing the ranges of characters affected by an action method
// that changes character (not paragraph) attributes.
//
// # Discussion
// 
// An array containing the ranges of characters affected by an action method
// that changes character (not paragraph) attributes, such as the NSText
// action method [ChangeFont]. For rich text these ranges are typically the
// ranges of the selections. For plain text the range is the entire contents
// of the receiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/rangesForUserCharacterAttributeChange
func (t NSTextView) RangesForUserCharacterAttributeChange() []foundation.NSValue {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("rangesForUserCharacterAttributeChange"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSValue {
		return foundation.NSValueFromID(id)
	})
}
// The range of characters affected by an action method that changes paragraph
// (not character) attributes.
//
// # Discussion
// 
// The range of characters affected by an action method that changes paragraph
// (not character) attributes, such as the NSText action method [AlignLeft].
// For rich text this range is typically calculated by extending the range of
// the selection to paragraph boundaries. For plain text this range is the
// entire contents of the receiver. If the receiver isn’t editable or
// doesn’t use the Font panel, the range has a location of [NSNotFound].
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/rangeForUserParagraphAttributeChange
func (t NSTextView) RangeForUserParagraphAttributeChange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](t.ID, objc.Sel("rangeForUserParagraphAttributeChange"))
	return foundation.NSRange(rv)
}
// An array containing the ranges of characters affected by a method that
// changes paragraph (not character) attributes.
//
// # Discussion
// 
// For rich text these ranges are typically calculated by extending the range
// of the selection to paragraph boundaries. For plain text the range is the
// entire contents of the receiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/rangesForUserParagraphAttributeChange
func (t NSTextView) RangesForUserParagraphAttributeChange() []foundation.NSValue {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("rangesForUserParagraphAttributeChange"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSValue {
		return foundation.NSValueFromID(id)
	})
}
// The range of characters affected by a method that changes characters (as
// opposed to attributes).
//
// # Discussion
// 
// If the receiver isn’t editable the range has a location of [NSNotFound].
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/rangeForUserTextChange
func (t NSTextView) RangeForUserTextChange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](t.ID, objc.Sel("rangeForUserTextChange"))
	return foundation.NSRange(rv)
}
// An array containing the ranges of characters affected by a method that
// changes characters (as opposed to attributes).
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/rangesForUserTextChange
func (t NSTextView) RangesForUserTextChange() []foundation.NSValue {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("rangesForUserTextChange"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSValue {
		return foundation.NSValueFromID(id)
	})
}
// A Boolean value that controls whether the receiver inserts or deletes space
// around selected words so as to preserve proper spacing and punctuation.
//
// # Discussion
// 
// [true] if the receiver should insert or delete space around selected words
// so as to preserve proper spacing and punctuation, [false] if it should
// insert and delete exactly what’s selected.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/smartInsertDeleteEnabled
func (t NSTextView) SmartInsertDeleteEnabled() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("smartInsertDeleteEnabled"))
	return rv
}
func (t NSTextView) SetSmartInsertDeleteEnabled(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setSmartInsertDeleteEnabled:"), value)
}
// A Boolean value that indicates whether the receiver has continuous spell
// checking enabled.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/isContinuousSpellCheckingEnabled
func (t NSTextView) ContinuousSpellCheckingEnabled() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isContinuousSpellCheckingEnabled"))
	return rv
}
func (t NSTextView) SetContinuousSpellCheckingEnabled(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setContinuousSpellCheckingEnabled:"), value)
}
// A tag identifying the text view’s text as a document for the spell
// checker server.
//
// # Discussion
// 
// The document tag is obtained by sending a [UniqueSpellDocumentTag] message
// to the spell server the first time this method is invoked for a particular
// group of text views. See the [NSSpellChecker]and [NSSpellServer]class
// specifications for more information on how this tag is used.
//
// [NSSpellServer]: https://developer.apple.com/documentation/Foundation/NSSpellServer
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/spellCheckerDocumentTag
func (t NSTextView) SpellCheckerDocumentTag() int {
	rv := objc.Send[int](t.ID, objc.Sel("spellCheckerDocumentTag"))
	return rv
}
// Enables and disables grammar checking.
//
// # Discussion
// 
// If [true], grammar checking is enabled; if [false], it is disabled.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/isGrammarCheckingEnabled
func (t NSTextView) GrammarCheckingEnabled() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isGrammarCheckingEnabled"))
	return rv
}
func (t NSTextView) SetGrammarCheckingEnabled(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setGrammarCheckingEnabled:"), value)
}
// A Boolean value that indicates whether the receiver accepts the glyph info
// attribute.
//
// # Discussion
// 
// [true] if the receiver accepts the [NSGlyphInfoAttributeName] attribute
// from text input sources such as input methods and the pasteboard, [false]
// otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/acceptsGlyphInfo
func (t NSTextView) AcceptsGlyphInfo() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("acceptsGlyphInfo"))
	return rv
}
func (t NSTextView) SetAcceptsGlyphInfo(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setAcceptsGlyphInfo:"), value)
}
// A Boolean value that indicates whether the receiver allows for a find
// panel.
//
// # Discussion
// 
// [true] to allow the use of a find panel, [false] otherwise. A text view can
// use either a find panel or a find bar. If [UsesFindPanel] is set to [true],
// [UsesFindBar] is set to [false] and vice versa.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/usesFindPanel
func (t NSTextView) UsesFindPanel() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("usesFindPanel"))
	return rv
}
func (t NSTextView) SetUsesFindPanel(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setUsesFindPanel:"), value)
}
// The partial range from the most recent beginning of a word up to the
// insertion point.
//
// # Discussion
// 
// This value is intended to be used for the range argument in the text
// completion methods such as
// [CompletionsForPartialWordRangeIndexOfSelectedItem].
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/rangeForUserCompletion
func (t NSTextView) RangeForUserCompletion() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](t.ID, objc.Sel("rangeForUserCompletion"))
	return foundation.NSRange(rv)
}
// The default text checking types.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/enabledTextCheckingTypes
func (t NSTextView) EnabledTextCheckingTypes() uint64 {
	rv := objc.Send[uint64](t.ID, objc.Sel("enabledTextCheckingTypes"))
	return rv
}
func (t NSTextView) SetEnabledTextCheckingTypes(value uint64) {
	objc.Send[struct{}](t.ID, objc.Sel("setEnabledTextCheckingTypes:"), value)
}
// A Boolean value that indicates whether automatic dash substitution is
// enabled.
//
// # Discussion
// 
// [true] if it is enabled, otherwise [false].
// 
// Turning on automatic dash substitution enables automatic conversion of a
// two ASCII hyphen (`-`) characters into an em-dash.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/isAutomaticDashSubstitutionEnabled
func (t NSTextView) AutomaticDashSubstitutionEnabled() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isAutomaticDashSubstitutionEnabled"))
	return rv
}
func (t NSTextView) SetAutomaticDashSubstitutionEnabled(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setAutomaticDashSubstitutionEnabled:"), value)
}
// A Boolean value that indicates whether automatic data detection is enabled.
//
// # Discussion
// 
// Automatic data detection enables detection of dates, addresses, and phone
// numbers.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/isAutomaticDataDetectionEnabled
func (t NSTextView) AutomaticDataDetectionEnabled() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isAutomaticDataDetectionEnabled"))
	return rv
}
func (t NSTextView) SetAutomaticDataDetectionEnabled(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setAutomaticDataDetectionEnabled:"), value)
}
// A Boolean value that indicates whether automatic spelling correction is
// enabled.
//
// # Discussion
// 
// [true] if it is enabled, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/isAutomaticSpellingCorrectionEnabled
func (t NSTextView) AutomaticSpellingCorrectionEnabled() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isAutomaticSpellingCorrectionEnabled"))
	return rv
}
func (t NSTextView) SetAutomaticSpellingCorrectionEnabled(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setAutomaticSpellingCorrectionEnabled:"), value)
}
// A Boolean value that indicates whether automatic text replacement is
// enabled.
//
// # Discussion
// 
// [true] if it is enabled, otherwise [false].
// 
// Turning on automatic text replacement enables automatic substitution of a
// variety of static text items based on user preferences.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/isAutomaticTextReplacementEnabled
func (t NSTextView) AutomaticTextReplacementEnabled() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isAutomaticTextReplacementEnabled"))
	return rv
}
func (t NSTextView) SetAutomaticTextReplacementEnabled(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setAutomaticTextReplacementEnabled:"), value)
}
// See: https://developer.apple.com/documentation/AppKit/NSTextView/isWritingToolsActive
func (t NSTextView) WritingToolsActive() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isWritingToolsActive"))
	return rv
}
// A Boolean value that indicates whether to use the find bar for this text
// view.
//
// # Discussion
// 
// The value of this property is [true] if the find bar is used for this text
// view; otherwise [false]. See [NSTextFinder] for information about the find
// bar.
// 
// A text view can use either a find panel or a find bar. If [UsesFindBar] is
// set to [true], [UsesFindPanel] is set to [false] and vice versa.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/usesFindBar
func (t NSTextView) UsesFindBar() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("usesFindBar"))
	return rv
}
func (t NSTextView) SetUsesFindBar(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setUsesFindBar:"), value)
}
// A Boolean value that indicates whether incremental searching is enabled.
//
// # Discussion
// 
// See [NSTextFinder] for information about the find bar.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/isIncrementalSearchingEnabled
func (t NSTextView) IncrementalSearchingEnabled() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isIncrementalSearchingEnabled"))
	return rv
}
func (t NSTextView) SetIncrementalSearchingEnabled(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setIncrementalSearchingEnabled:"), value)
}
// See: https://developer.apple.com/documentation/AppKit/NSTextView/allowsCharacterPickerTouchBarItem
func (t NSTextView) AllowsCharacterPickerTouchBarItem() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("allowsCharacterPickerTouchBarItem"))
	return rv
}
func (t NSTextView) SetAllowsCharacterPickerTouchBarItem(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setAllowsCharacterPickerTouchBarItem:"), value)
}
// See: https://developer.apple.com/documentation/AppKit/NSTextView/allowedWritingToolsResultOptions
func (t NSTextView) AllowedWritingToolsResultOptions() NSWritingToolsResultOptions {
	rv := objc.Send[NSWritingToolsResultOptions](t.ID, objc.Sel("allowedWritingToolsResultOptions"))
	return NSWritingToolsResultOptions(rv)
}
func (t NSTextView) SetAllowedWritingToolsResultOptions(value NSWritingToolsResultOptions) {
	objc.Send[struct{}](t.ID, objc.Sel("setAllowedWritingToolsResultOptions:"), value)
}
// See: https://developer.apple.com/documentation/AppKit/NSTextView/inlinePredictionType
func (t NSTextView) InlinePredictionType() NSTextInputTraitType {
	rv := objc.Send[NSTextInputTraitType](t.ID, objc.Sel("inlinePredictionType"))
	return NSTextInputTraitType(rv)
}
func (t NSTextView) SetInlinePredictionType(value NSTextInputTraitType) {
	objc.Send[struct{}](t.ID, objc.Sel("setInlinePredictionType:"), value)
}
// See: https://developer.apple.com/documentation/AppKit/NSTextView/mathExpressionCompletionType
func (t NSTextView) MathExpressionCompletionType() NSTextInputTraitType {
	rv := objc.Send[NSTextInputTraitType](t.ID, objc.Sel("mathExpressionCompletionType"))
	return NSTextInputTraitType(rv)
}
func (t NSTextView) SetMathExpressionCompletionType(value NSTextInputTraitType) {
	objc.Send[struct{}](t.ID, objc.Sel("setMathExpressionCompletionType:"), value)
}
// ************************* Text Highlight support **************************
//
// See: https://developer.apple.com/documentation/AppKit/NSTextView/textHighlightAttributes
func (t NSTextView) TextHighlightAttributes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("textHighlightAttributes"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (t NSTextView) SetTextHighlightAttributes(value foundation.INSDictionary) {
	objc.Send[struct{}](t.ID, objc.Sel("setTextHighlightAttributes:"), value)
}
// See: https://developer.apple.com/documentation/AppKit/NSTextView/writingToolsBehavior
func (t NSTextView) WritingToolsBehavior() NSWritingToolsBehavior {
	rv := objc.Send[NSWritingToolsBehavior](t.ID, objc.Sel("writingToolsBehavior"))
	return NSWritingToolsBehavior(rv)
}
func (t NSTextView) SetWritingToolsBehavior(value NSWritingToolsBehavior) {
	objc.Send[struct{}](t.ID, objc.Sel("setWritingToolsBehavior:"), value)
}
// The semantic meaning for a text input area.
//
// # Discussion
// 
// Use this property to give the system information about the expected
// semantic meaning for the content that people enter. For example, you might
// specify [emailAddress] for a text field that people fill in to receive an
// email confirmation.
// 
// For possible values you can use, see [NSTextContentType]; by default, the
// value of this property is `nil`.
//
// [emailAddress]: https://developer.apple.com/documentation/AppKit/NSTextContentType/emailAddress
//
// See: https://developer.apple.com/documentation/AppKit/NSTextContent/contentType
func (t NSTextView) ContentType() NSTextContentType {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("contentType"))
	return NSTextContentType(foundation.NSStringFromID(rv).String())
}
func (t NSTextView) SetContentType(value NSTextContentType) {
	objc.Send[struct{}](t.ID, objc.Sel("setContentType:"), objc.String(string(value)))
}
// See: https://developer.apple.com/documentation/AppKit/NSTextInputClient/documentVisibleRect
func (t NSTextView) DocumentVisibleRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](t.ID, objc.Sel("documentVisibleRect"))
	return corefoundation.CGRect(rv)
}
// Type for the find panel metadata property list.
//
// See: https://developer.apple.com/documentation/appkit/nspasteboard/pasteboardtype/findpanelsearchoptions
func (t NSTextView) FindPanelSearchOptions() NSPasteboardType {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("NSFindPanelSearchOptionsPboardType"))
	return NSPasteboardType(foundation.NSStringFromID(rv).String())
}
// The default layout orientation.
//
// # Discussion
// 
// This property contains the default layout orientation for text in the
// object that adopts the protocol. If the text contains an explicit
// [verticalGlyphForm] attribute, that attribute overrides the value in this
// property. When rendering, TextKit assumes the coordinate system is
// appropriately rotated.
//
// [verticalGlyphForm]: https://developer.apple.com/documentation/Foundation/NSAttributedString/Key/verticalGlyphForm
//
// See: https://developer.apple.com/documentation/AppKit/NSTextLayoutOrientationProvider/layoutOrientation
func (t NSTextView) LayoutOrientation() NSTextLayoutOrientation {
	rv := objc.Send[NSTextLayoutOrientation](t.ID, objc.Sel("layoutOrientation"))
	return NSTextLayoutOrientation(rv)
}
// A Boolean value that indicates whether the document supports adaptive
// images in the input.
//
// # Discussion
// 
// When this property is [false], the input system doesn’t allow the text
// input to contain adaptive images. Set the value of this property to [true]
// only if your document supports adaptive images and handles them properly.
// For more information, see [NSAdaptiveImageGlyph].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextInputClient/supportsAdaptiveImageGlyph
func (t NSTextView) SupportsAdaptiveImageGlyph() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("supportsAdaptiveImageGlyph"))
	return rv
}
// See: https://developer.apple.com/documentation/AppKit/NSTextInputClient/unionRectInVisibleSelectedRange
func (t NSTextView) UnionRectInVisibleSelectedRange() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](t.ID, objc.Sel("unionRectInVisibleSelectedRange"))
	return corefoundation.CGRect(rv)
}

// See: https://developer.apple.com/documentation/AppKit/NSTextView/stronglyReferencesTextStorage
func (_NSTextViewClass NSTextViewClass) StronglyReferencesTextStorage() bool {
	rv := objc.Send[bool](objc.ID(_NSTextViewClass.class), objc.Sel("stronglyReferencesTextStorage"))
	return rv
}
// Posted when focus leaves an
//
// See: https://developer.apple.com/documentation/appkit/nstext/didendeditingnotification
func (_NSTextViewClass NSTextViewClass) DidEndEditingNotification() foundation.NSString {
	rv := objc.Send[objc.ID](objc.ID(_NSTextViewClass.class), objc.Sel("NSTextDidEndEditingNotification"))
	return foundation.NSStringFromID(objc.ID(rv))
}

			// Protocol methods for NSAccessibilityNavigableStaticText
			
// Returns the accessibility element’s frame in screen coordinates.
//
// # Return Value
// 
// The element’s frame in screen coordinates.
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityFrame] property. This method is called whenever accessibility
// clients request the [size] or [position] attributes.
//
// [accessibilityFrame]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFrame
// [position]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/position
// [size]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/size
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityFrame()
func (o NSTextView) AccessibilityFrame() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("accessibilityFrame"))
	return rv
	}
// Returns the accessibility element’s parent in the accessibility
// hierarchy.
//
// # Return Value
// 
// The element’s parent in the accessibility hierarchy.
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityParent] property.
//
// [accessibilityParent]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityParent
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityParent()
func (o NSTextView) AccessibilityParent() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityParent"))
	return objectivec.Object{ID: rv}
	}
// Returns the accessibility element’s identity.
//
// # Return Value
// 
// Returns the unique ID for the accessibility element. It is often used in
// automated testing.
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityIdentifier] property.
//
// [accessibilityIdentifier]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityIdentifier
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityIdentifier()
func (o NSTextView) AccessibilityIdentifier() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityIdentifier"))
	return foundation.NSStringFromID(rv).String()
	}
// Returns a Boolean value that indicates whether the accessibility element
// has the keyboard focus.
//
// # Return Value
// 
// [true] if this element has the keyboard focus; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityFocused] property.
//
// [accessibilityFocused]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFocused
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/isAccessibilityFocused()
func (o NSTextView) IsAccessibilityFocused() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFocused"))
	return rv
	}

			// Protocol methods for NSCandidateListTouchBarItemDelegate
			

			// Protocol methods for NSDraggingSource
			

			// Protocol methods for NSMenuItemValidation
			

			// Protocol methods for NSTextContent
			

			// Protocol methods for NSTextInput
			

			// Protocol methods for NSTextInputClient
			

			// Protocol methods for NSTextLayoutOrientationProvider
			

			// Protocol methods for NSTouchBarDelegate
			

			// Protocol methods for NSUserInterfaceValidations
			

