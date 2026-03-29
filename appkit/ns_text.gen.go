// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSText] class.
var (
	_NSTextClass     NSTextClass
	_NSTextClassOnce sync.Once
)

func getNSTextClass() NSTextClass {
	_NSTextClassOnce.Do(func() {
		_NSTextClass = NSTextClass{class: objc.GetClass("NSText")}
	})
	return _NSTextClass
}

// GetNSTextClass returns the class object for NSText.
func GetNSTextClass() NSTextClass {
	return getNSTextClass()
}

type NSTextClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSTextClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTextClass) Alloc() NSText {
	rv := objc.Send[NSText](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The most general programmatic interface for objects that manage text.
//
// # Overview
// 
// [NSText] draws text for user interface objects, provides text editing
// capabilities, and controls text attributes such as type size, font, and
// color.
// 
// [NSText] initialization creates an instance of a concrete subclass, such as
// [NSTextView] (generically called a text object). In general, you’re more
// likely to use the [NSTextView] subclass, because it extends the interface
// declared by [NSText] and provides much more sophisticated functionality
// than that declared in [NSText].
// 
// AppKit uses text objects wherever text appears in interface objects. For
// example, a text object draws the title of a window, the commands in a menu,
// the title of a button, and the items in a browser. Your app can also create
// text objects for its own purposes.
//
// # Setting graphics attributes
//
//   - [NSText.BackgroundColor]: The receiver’s background color to a given color.
//   - [NSText.SetBackgroundColor]
//   - [NSText.DrawsBackground]: A Boolean that controls whether the receiver draws its background.
//   - [NSText.SetDrawsBackground]
//
// # Setting behavioral attributes
//
//   - [NSText.Editable]: A Boolean that controls whether the receiver allows the user to edit its text.
//   - [NSText.SetEditable]
//   - [NSText.Selectable]: A Boolean that controls whether the receiver allows the user to select its text.
//   - [NSText.SetSelectable]
//   - [NSText.FieldEditor]: A Boolean that controls whether the receiver interprets Tab, Shift-Tab, and Return (Enter) as cues to end editing and possibly to change the first responder.
//   - [NSText.SetFieldEditor]
//   - [NSText.RichText]: A Boolean that controls whether the receiver allows the user to apply attributes to specific ranges of the text.
//   - [NSText.SetRichText]
//   - [NSText.ImportsGraphics]: A Boolean that controls whether the receiver allows the user to import files by dragging.
//   - [NSText.SetImportsGraphics]
//
// # Using the Font panel and menu
//
//   - [NSText.UsesFontPanel]: A Boolean that controls whether the receiver uses the Font panel and Font menu.
//   - [NSText.SetUsesFontPanel]
//
// # Using the ruler
//
//   - [NSText.ToggleRuler]: This action method shows or hides the ruler, if the receiver is enclosed in a scroll view.
//   - [NSText.RulerVisible]: A Boolean value that indicates whether the receiver’s enclosing scroll view shows its ruler.
//
// # Changing the selection
//
//   - [NSText.SelectedRange]: The receiver’s characters within `aRange`.
//   - [NSText.SetSelectedRange]
//
// # Replacing text
//
//   - [NSText.ReplaceCharactersInRangeWithRTF]: Replaces the characters in the given range with RTF text interpreted from the given RTF data.
//   - [NSText.ReplaceCharactersInRangeWithRTFD]: Replaces the characters in the given range with RTFD text interpreted from the given RTFD data.
//   - [NSText.ReplaceCharactersInRangeWithString]: Replaces the characters in the given range with those in the given string.
//
// # Action methods for editing
//
//   - [NSText.Copy]: This action method copies the selected text onto the general pasteboard, in as many formats as the receiver supports.
//   - [NSText.Cut]: This action method deletes the selected text and places it onto the general pasteboard, in as many formats as the receiver supports.
//   - [NSText.Paste]: This action method pastes text from the general pasteboard at the insertion point or over the selection.
//   - [NSText.CopyFont]: This action method copies the font information for the first character of the selection (or for the insertion point) onto the font pasteboard, as [NSFontPboardType].
//   - [NSText.PasteFont]: This action method pastes font information from the font pasteboard onto the selected text or insertion point of a rich text object, or over all text of a plain text object.
//   - [NSText.CopyRuler]: This action method copies the paragraph style information for first selected paragraph onto the ruler pasteboard, as [NSRulerPboardType], and expands the selection to paragraph boundaries.
//   - [NSText.PasteRuler]: This action method pastes paragraph style information from the ruler pasteboard onto the selected paragraphs of a rich text object.
//   - [NSText.Delete]: This action method deletes the selected text.
//
// # Changing the font
//
//   - [NSText.ChangeFont]: This action method changes the font of the selection for a rich text object, or of all text for a plain text object.
//   - [NSText.Font]: The font of all the receiver’s text.
//   - [NSText.SetFont]
//   - [NSText.SetFontRange]: Sets the font of characters within `aRange` to `aFont`.
//
// # Setting text alignment
//
//   - [NSText.Alignment]: The alignment of all the receiver’s text.
//   - [NSText.SetAlignment]
//   - [NSText.AlignCenter]: This action method applies center alignment to selected paragraphs (or all text if the receiver is a plain text object).
//   - [NSText.AlignLeft]: This action method applies left alignment to selected paragraphs (or all text if the receiver is a plain text object).
//   - [NSText.AlignRight]: This action method applies right alignment to selected paragraphs (or all text if the receiver is a plain text object).
//
// # Setting text color
//
//   - [NSText.TextColor]: The text color of all characters in the receiver.
//   - [NSText.SetTextColor]
//   - [NSText.SetTextColorRange]: Sets the text color of characters within the specified range to the specified color.
//
// # Writing direction
//
//   - [NSText.BaseWritingDirection]: The initial writing direction used to determine the actual writing direction for text.
//   - [NSText.SetBaseWritingDirection]
//
// # Setting superscripting and subscripting
//
//   - [NSText.Superscript]: This action method applies a superscript attribute to selected text (or all text if the receiver is a plain text object), raising its baseline offset by a predefined amount.
//   - [NSText.Subscript]: This action method applies a subscript attribute to selected text (or all text if the receiver is a plain text object), lowering its baseline offset by a predefined amount.
//   - [NSText.Unscript]: This action method removes any superscripting or subscripting from selected text (or all text if the receiver is a plain text object).
//
// # Underlining text
//
//   - [NSText.Underline]: Adds the underline attribute to the selected text attributes if absent; removes the attribute if present.
//
// # Reading and writing RTF files
//
//   - [NSText.ReadRTFDFromFile]: Attempts to read the RTFD file at the specified path.
//   - [NSText.WriteRTFDToFileAtomically]: Writes the receiver’s text as RTF with attachments to a file or directory at `path`.
//   - [NSText.RTFDFromRange]: Returns an NSData object that contains an RTFD stream corresponding to the characters and attributes within `aRange`.
//   - [NSText.RTFFromRange]: Returns an NSData object that contains an RTF stream corresponding to the characters and attributes within `aRange`, omitting any attachment characters and attributes.
//
// # Checking spelling
//
//   - [NSText.CheckSpelling]: This action method searches for a misspelled word in the receiver’s text.
//   - [NSText.ShowGuessPanel]: This action method opens the Spelling panel, allowing the user to make a correction during spell checking.
//
// # Constraining size
//
//   - [NSText.MaxSize]: The receiver’s maximum size.
//   - [NSText.SetMaxSize]
//   - [NSText.MinSize]: The receiver’s minimum size.
//   - [NSText.SetMinSize]
//   - [NSText.VerticallyResizable]: A Boolean that controls whether the receiver changes its height to fit the height of its text.
//   - [NSText.SetVerticallyResizable]
//   - [NSText.HorizontallyResizable]: A Boolean that controls whether the receiver changes its width to fit the width of its text.
//   - [NSText.SetHorizontallyResizable]
//   - [NSText.SizeToFit]: Resizes the receiver to fit its text.
//
// # Scrolling
//
//   - [NSText.ScrollRangeToVisible]: Scrolls the receiver in its enclosing scroll view so the first characters of `aRange` are visible.
//
// # Setting the delegate
//
//   - [NSText.Delegate]: The receiver’s delegate.
//   - [NSText.SetDelegate]
//
// See: https://developer.apple.com/documentation/AppKit/NSText
type NSText struct {
	NSView
}

// NSTextFromID constructs a [NSText] from an objc.ID.
//
// The most general programmatic interface for objects that manage text.
func NSTextFromID(id objc.ID) NSText {
	return NSText{NSView: NSViewFromID(id)}
}
// NOTE: NSText adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSText] class.
//
// # Setting graphics attributes
//
//   - [INSText.BackgroundColor]: The receiver’s background color to a given color.
//   - [INSText.SetBackgroundColor]
//   - [INSText.DrawsBackground]: A Boolean that controls whether the receiver draws its background.
//   - [INSText.SetDrawsBackground]
//
// # Setting behavioral attributes
//
//   - [INSText.Editable]: A Boolean that controls whether the receiver allows the user to edit its text.
//   - [INSText.SetEditable]
//   - [INSText.Selectable]: A Boolean that controls whether the receiver allows the user to select its text.
//   - [INSText.SetSelectable]
//   - [INSText.FieldEditor]: A Boolean that controls whether the receiver interprets Tab, Shift-Tab, and Return (Enter) as cues to end editing and possibly to change the first responder.
//   - [INSText.SetFieldEditor]
//   - [INSText.RichText]: A Boolean that controls whether the receiver allows the user to apply attributes to specific ranges of the text.
//   - [INSText.SetRichText]
//   - [INSText.ImportsGraphics]: A Boolean that controls whether the receiver allows the user to import files by dragging.
//   - [INSText.SetImportsGraphics]
//
// # Using the Font panel and menu
//
//   - [INSText.UsesFontPanel]: A Boolean that controls whether the receiver uses the Font panel and Font menu.
//   - [INSText.SetUsesFontPanel]
//
// # Using the ruler
//
//   - [INSText.ToggleRuler]: This action method shows or hides the ruler, if the receiver is enclosed in a scroll view.
//   - [INSText.RulerVisible]: A Boolean value that indicates whether the receiver’s enclosing scroll view shows its ruler.
//
// # Changing the selection
//
//   - [INSText.SelectedRange]: The receiver’s characters within `aRange`.
//   - [INSText.SetSelectedRange]
//
// # Replacing text
//
//   - [INSText.ReplaceCharactersInRangeWithRTF]: Replaces the characters in the given range with RTF text interpreted from the given RTF data.
//   - [INSText.ReplaceCharactersInRangeWithRTFD]: Replaces the characters in the given range with RTFD text interpreted from the given RTFD data.
//   - [INSText.ReplaceCharactersInRangeWithString]: Replaces the characters in the given range with those in the given string.
//
// # Action methods for editing
//
//   - [INSText.Copy]: This action method copies the selected text onto the general pasteboard, in as many formats as the receiver supports.
//   - [INSText.Cut]: This action method deletes the selected text and places it onto the general pasteboard, in as many formats as the receiver supports.
//   - [INSText.Paste]: This action method pastes text from the general pasteboard at the insertion point or over the selection.
//   - [INSText.CopyFont]: This action method copies the font information for the first character of the selection (or for the insertion point) onto the font pasteboard, as [NSFontPboardType].
//   - [INSText.PasteFont]: This action method pastes font information from the font pasteboard onto the selected text or insertion point of a rich text object, or over all text of a plain text object.
//   - [INSText.CopyRuler]: This action method copies the paragraph style information for first selected paragraph onto the ruler pasteboard, as [NSRulerPboardType], and expands the selection to paragraph boundaries.
//   - [INSText.PasteRuler]: This action method pastes paragraph style information from the ruler pasteboard onto the selected paragraphs of a rich text object.
//   - [INSText.Delete]: This action method deletes the selected text.
//
// # Changing the font
//
//   - [INSText.ChangeFont]: This action method changes the font of the selection for a rich text object, or of all text for a plain text object.
//   - [INSText.Font]: The font of all the receiver’s text.
//   - [INSText.SetFont]
//   - [INSText.SetFontRange]: Sets the font of characters within `aRange` to `aFont`.
//
// # Setting text alignment
//
//   - [INSText.Alignment]: The alignment of all the receiver’s text.
//   - [INSText.SetAlignment]
//   - [INSText.AlignCenter]: This action method applies center alignment to selected paragraphs (or all text if the receiver is a plain text object).
//   - [INSText.AlignLeft]: This action method applies left alignment to selected paragraphs (or all text if the receiver is a plain text object).
//   - [INSText.AlignRight]: This action method applies right alignment to selected paragraphs (or all text if the receiver is a plain text object).
//
// # Setting text color
//
//   - [INSText.TextColor]: The text color of all characters in the receiver.
//   - [INSText.SetTextColor]
//   - [INSText.SetTextColorRange]: Sets the text color of characters within the specified range to the specified color.
//
// # Writing direction
//
//   - [INSText.BaseWritingDirection]: The initial writing direction used to determine the actual writing direction for text.
//   - [INSText.SetBaseWritingDirection]
//
// # Setting superscripting and subscripting
//
//   - [INSText.Superscript]: This action method applies a superscript attribute to selected text (or all text if the receiver is a plain text object), raising its baseline offset by a predefined amount.
//   - [INSText.Subscript]: This action method applies a subscript attribute to selected text (or all text if the receiver is a plain text object), lowering its baseline offset by a predefined amount.
//   - [INSText.Unscript]: This action method removes any superscripting or subscripting from selected text (or all text if the receiver is a plain text object).
//
// # Underlining text
//
//   - [INSText.Underline]: Adds the underline attribute to the selected text attributes if absent; removes the attribute if present.
//
// # Reading and writing RTF files
//
//   - [INSText.ReadRTFDFromFile]: Attempts to read the RTFD file at the specified path.
//   - [INSText.WriteRTFDToFileAtomically]: Writes the receiver’s text as RTF with attachments to a file or directory at `path`.
//   - [INSText.RTFDFromRange]: Returns an NSData object that contains an RTFD stream corresponding to the characters and attributes within `aRange`.
//   - [INSText.RTFFromRange]: Returns an NSData object that contains an RTF stream corresponding to the characters and attributes within `aRange`, omitting any attachment characters and attributes.
//
// # Checking spelling
//
//   - [INSText.CheckSpelling]: This action method searches for a misspelled word in the receiver’s text.
//   - [INSText.ShowGuessPanel]: This action method opens the Spelling panel, allowing the user to make a correction during spell checking.
//
// # Constraining size
//
//   - [INSText.MaxSize]: The receiver’s maximum size.
//   - [INSText.SetMaxSize]
//   - [INSText.MinSize]: The receiver’s minimum size.
//   - [INSText.SetMinSize]
//   - [INSText.VerticallyResizable]: A Boolean that controls whether the receiver changes its height to fit the height of its text.
//   - [INSText.SetVerticallyResizable]
//   - [INSText.HorizontallyResizable]: A Boolean that controls whether the receiver changes its width to fit the width of its text.
//   - [INSText.SetHorizontallyResizable]
//   - [INSText.SizeToFit]: Resizes the receiver to fit its text.
//
// # Scrolling
//
//   - [INSText.ScrollRangeToVisible]: Scrolls the receiver in its enclosing scroll view so the first characters of `aRange` are visible.
//
// # Setting the delegate
//
//   - [INSText.Delegate]: The receiver’s delegate.
//   - [INSText.SetDelegate]
//
// See: https://developer.apple.com/documentation/AppKit/NSText
type INSText interface {
	INSView
	NSChangeSpelling
	NSIgnoreMisspelledWords

	// Topic: Setting graphics attributes

	// The receiver’s background color to a given color.
	BackgroundColor() INSColor
	SetBackgroundColor(value INSColor)
	// A Boolean that controls whether the receiver draws its background.
	DrawsBackground() bool
	SetDrawsBackground(value bool)

	// Topic: Setting behavioral attributes

	// A Boolean that controls whether the receiver allows the user to edit its text.
	Editable() bool
	SetEditable(value bool)
	// A Boolean that controls whether the receiver allows the user to select its text.
	Selectable() bool
	SetSelectable(value bool)
	// A Boolean that controls whether the receiver interprets Tab, Shift-Tab, and Return (Enter) as cues to end editing and possibly to change the first responder.
	FieldEditor() bool
	SetFieldEditor(value bool)
	// A Boolean that controls whether the receiver allows the user to apply attributes to specific ranges of the text.
	RichText() bool
	SetRichText(value bool)
	// A Boolean that controls whether the receiver allows the user to import files by dragging.
	ImportsGraphics() bool
	SetImportsGraphics(value bool)

	// Topic: Using the Font panel and menu

	// A Boolean that controls whether the receiver uses the Font panel and Font menu.
	UsesFontPanel() bool
	SetUsesFontPanel(value bool)

	// Topic: Using the ruler

	// This action method shows or hides the ruler, if the receiver is enclosed in a scroll view.
	ToggleRuler(sender objectivec.IObject)
	// A Boolean value that indicates whether the receiver’s enclosing scroll view shows its ruler.
	RulerVisible() bool

	// Topic: Changing the selection

	// The receiver’s characters within `aRange`.
	SelectedRange() foundation.NSRange
	SetSelectedRange(value foundation.NSRange)

	// Topic: Replacing text

	// Replaces the characters in the given range with RTF text interpreted from the given RTF data.
	ReplaceCharactersInRangeWithRTF(range_ foundation.NSRange, rtfData foundation.INSData)
	// Replaces the characters in the given range with RTFD text interpreted from the given RTFD data.
	ReplaceCharactersInRangeWithRTFD(range_ foundation.NSRange, rtfdData foundation.INSData)
	// Replaces the characters in the given range with those in the given string.
	ReplaceCharactersInRangeWithString(range_ foundation.NSRange, string_ string)

	// Topic: Action methods for editing

	// This action method copies the selected text onto the general pasteboard, in as many formats as the receiver supports.
	Copy(sender objectivec.IObject)
	// This action method deletes the selected text and places it onto the general pasteboard, in as many formats as the receiver supports.
	Cut(sender objectivec.IObject)
	// This action method pastes text from the general pasteboard at the insertion point or over the selection.
	Paste(sender objectivec.IObject)
	// This action method copies the font information for the first character of the selection (or for the insertion point) onto the font pasteboard, as [NSFontPboardType].
	CopyFont(sender objectivec.IObject)
	// This action method pastes font information from the font pasteboard onto the selected text or insertion point of a rich text object, or over all text of a plain text object.
	PasteFont(sender objectivec.IObject)
	// This action method copies the paragraph style information for first selected paragraph onto the ruler pasteboard, as [NSRulerPboardType], and expands the selection to paragraph boundaries.
	CopyRuler(sender objectivec.IObject)
	// This action method pastes paragraph style information from the ruler pasteboard onto the selected paragraphs of a rich text object.
	PasteRuler(sender objectivec.IObject)
	// This action method deletes the selected text.
	Delete(sender objectivec.IObject)

	// Topic: Changing the font

	// This action method changes the font of the selection for a rich text object, or of all text for a plain text object.
	ChangeFont(sender objectivec.IObject)
	// The font of all the receiver’s text.
	Font() NSFont
	SetFont(value NSFont)
	// Sets the font of characters within `aRange` to `aFont`.
	SetFontRange(font NSFont, range_ foundation.NSRange)

	// Topic: Setting text alignment

	// The alignment of all the receiver’s text.
	Alignment() NSTextAlignment
	SetAlignment(value NSTextAlignment)
	// This action method applies center alignment to selected paragraphs (or all text if the receiver is a plain text object).
	AlignCenter(sender objectivec.IObject)
	// This action method applies left alignment to selected paragraphs (or all text if the receiver is a plain text object).
	AlignLeft(sender objectivec.IObject)
	// This action method applies right alignment to selected paragraphs (or all text if the receiver is a plain text object).
	AlignRight(sender objectivec.IObject)

	// Topic: Setting text color

	// The text color of all characters in the receiver.
	TextColor() INSColor
	SetTextColor(value INSColor)
	// Sets the text color of characters within the specified range to the specified color.
	SetTextColorRange(color INSColor, range_ foundation.NSRange)

	// Topic: Writing direction

	// The initial writing direction used to determine the actual writing direction for text.
	BaseWritingDirection() NSWritingDirection
	SetBaseWritingDirection(value NSWritingDirection)

	// Topic: Setting superscripting and subscripting

	// This action method applies a superscript attribute to selected text (or all text if the receiver is a plain text object), raising its baseline offset by a predefined amount.
	Superscript(sender objectivec.IObject)
	// This action method applies a subscript attribute to selected text (or all text if the receiver is a plain text object), lowering its baseline offset by a predefined amount.
	Subscript(sender objectivec.IObject)
	// This action method removes any superscripting or subscripting from selected text (or all text if the receiver is a plain text object).
	Unscript(sender objectivec.IObject)

	// Topic: Underlining text

	// Adds the underline attribute to the selected text attributes if absent; removes the attribute if present.
	Underline(sender objectivec.IObject)

	// Topic: Reading and writing RTF files

	// Attempts to read the RTFD file at the specified path.
	ReadRTFDFromFile(path string) bool
	// Writes the receiver’s text as RTF with attachments to a file or directory at `path`.
	WriteRTFDToFileAtomically(path string, flag bool) bool
	// Returns an NSData object that contains an RTFD stream corresponding to the characters and attributes within `aRange`.
	RTFDFromRange(range_ foundation.NSRange) foundation.INSData
	// Returns an NSData object that contains an RTF stream corresponding to the characters and attributes within `aRange`, omitting any attachment characters and attributes.
	RTFFromRange(range_ foundation.NSRange) foundation.INSData

	// Topic: Checking spelling

	// This action method searches for a misspelled word in the receiver’s text.
	CheckSpelling(sender objectivec.IObject)
	// This action method opens the Spelling panel, allowing the user to make a correction during spell checking.
	ShowGuessPanel(sender objectivec.IObject)

	// Topic: Constraining size

	// The receiver’s maximum size.
	MaxSize() corefoundation.CGSize
	SetMaxSize(value corefoundation.CGSize)
	// The receiver’s minimum size.
	MinSize() corefoundation.CGSize
	SetMinSize(value corefoundation.CGSize)
	// A Boolean that controls whether the receiver changes its height to fit the height of its text.
	VerticallyResizable() bool
	SetVerticallyResizable(value bool)
	// A Boolean that controls whether the receiver changes its width to fit the width of its text.
	HorizontallyResizable() bool
	SetHorizontallyResizable(value bool)
	// Resizes the receiver to fit its text.
	SizeToFit()

	// Topic: Scrolling

	// Scrolls the receiver in its enclosing scroll view so the first characters of `aRange` are visible.
	ScrollRangeToVisible(range_ foundation.NSRange)

	// Topic: Setting the delegate

	// The receiver’s delegate.
	Delegate() NSTextDelegate
	SetDelegate(value NSTextDelegate)
}

// Init initializes the instance.
func (t NSText) Init() NSText {
	rv := objc.Send[NSText](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSText) Autorelease() NSText {
	rv := objc.Send[NSText](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSText creates a new NSText instance.
func NewNSText() NSText {
	class := getNSTextClass()
	rv := objc.Send[NSText](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppKit/NSText/init(coder:)
func NewTextWithCoder(coder foundation.INSCoder) NSText {
	instance := getNSTextClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSTextFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSText/init(frame:)
func NewTextWithFrame(frameRect corefoundation.CGRect) NSText {
	instance := getNSTextClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSTextFromID(rv)
}

// This action method shows or hides the ruler, if the receiver is enclosed in
// a scroll view.
//
// See: https://developer.apple.com/documentation/AppKit/NSText/toggleRuler(_:)
func (t NSText) ToggleRuler(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("toggleRuler:"), sender)
}
// Replaces the characters in the given range with RTF text interpreted from
// the given RTF data.
//
// range: The range of characters to be replaced.
//
// rtfData: The RTF data from which to derive the replacement string.
//
// # Discussion
// 
// This method applies only to rich text objects.
// 
// This method does not include undo support by default. Clients must invoke
// [ShouldChangeTextInRangesReplacementStrings] or
// [ShouldChangeTextInRangeReplacementString] to include this method in an
// undoable action.
// 
// This method is designed for transferring text from out-of-process sources
// such as the pasteboard. In most cases, programmatic modification of the
// text is best done by operating on the text storage directly, using the
// general methods of [NSMutableAttributedString].
//
// [NSMutableAttributedString]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString
//
// See: https://developer.apple.com/documentation/AppKit/NSText/replaceCharacters(in:withRTF:)
func (t NSText) ReplaceCharactersInRangeWithRTF(range_ foundation.NSRange, rtfData foundation.INSData) {
	objc.Send[objc.ID](t.ID, objc.Sel("replaceCharactersInRange:withRTF:"), range_, rtfData)
}
// Replaces the characters in the given range with RTFD text interpreted from
// the given RTFD data.
//
// range: The range of characters to be replaced.
//
// rtfdData: The RTFD data from which to derive the replacement string.
//
// # Discussion
// 
// This method applies only to rich text objects.
// 
// This method does not include undo support by default. Clients must invoke
// [ShouldChangeTextInRangesReplacementStrings] or
// [ShouldChangeTextInRangeReplacementString] to include this method in an
// undoable action.
// 
// This method is designed for transferring text from out-of-process sources
// such as the pasteboard. In most cases, programmatic modification of the
// text is best done by operating on the text storage directly, using the
// general methods of [NSMutableAttributedString].
//
// [NSMutableAttributedString]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString
//
// See: https://developer.apple.com/documentation/AppKit/NSText/replaceCharacters(in:withRTFD:)
func (t NSText) ReplaceCharactersInRangeWithRTFD(range_ foundation.NSRange, rtfdData foundation.INSData) {
	objc.Send[objc.ID](t.ID, objc.Sel("replaceCharactersInRange:withRTFD:"), range_, rtfdData)
}
// Replaces the characters in the given range with those in the given string.
//
// range: The range of characters to be replaced.
//
// string: The replacement string.
//
// # Discussion
// 
// For a rich text object, the text of `aString` is assigned the formatting
// attributes of the first character of the text it replaces, or of the
// character immediately before `aRange` if the range’s length is 0. If the
// range’s location is 0, the formatting attributes of the first character
// in the receiver are used.
// 
// This method does not include undo support by default. Clients must invoke
// [ShouldChangeTextInRangesReplacementStrings] or
// [ShouldChangeTextInRangeReplacementString] to include this method in an
// undoable action.
// 
// In most cases, programmatic modification of the text is best done by
// operating on the text storage directly, using the general methods of
// [NSMutableAttributedString].
//
// [NSMutableAttributedString]: https://developer.apple.com/documentation/Foundation/NSMutableAttributedString
//
// See: https://developer.apple.com/documentation/AppKit/NSText/replaceCharacters(in:with:)
func (t NSText) ReplaceCharactersInRangeWithString(range_ foundation.NSRange, string_ string) {
	objc.Send[objc.ID](t.ID, objc.Sel("replaceCharactersInRange:withString:"), range_, objc.String(string_))
}
// This action method copies the selected text onto the general pasteboard, in
// as many formats as the receiver supports.
//
// # Discussion
// 
// A plain text object uses [NSStringPboardType] for plain text, and a rich
// text object also uses [NSRTFPboardType].
//
// See: https://developer.apple.com/documentation/AppKit/NSText/copy(_:)
func (t NSText) Copy(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("copy:"), sender)
}
// This action method deletes the selected text and places it onto the general
// pasteboard, in as many formats as the receiver supports.
//
// # Discussion
// 
// A plain text object uses [NSStringPboardType] for plain text, and a rich
// text object also uses [NSRTFPboardType].
//
// See: https://developer.apple.com/documentation/AppKit/NSText/cut(_:)
func (t NSText) Cut(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("cut:"), sender)
}
// This action method pastes text from the general pasteboard at the insertion
// point or over the selection.
//
// See: https://developer.apple.com/documentation/AppKit/NSText/paste(_:)
func (t NSText) Paste(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("paste:"), sender)
}
// This action method copies the font information for the first character of
// the selection (or for the insertion point) onto the font pasteboard, as
// [NSFontPboardType].
//
// See: https://developer.apple.com/documentation/AppKit/NSText/copyFont(_:)
func (t NSText) CopyFont(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("copyFont:"), sender)
}
// This action method pastes font information from the font pasteboard onto
// the selected text or insertion point of a rich text object, or over all
// text of a plain text object.
//
// See: https://developer.apple.com/documentation/AppKit/NSText/pasteFont(_:)
func (t NSText) PasteFont(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("pasteFont:"), sender)
}
// This action method copies the paragraph style information for first
// selected paragraph onto the ruler pasteboard, as [NSRulerPboardType], and
// expands the selection to paragraph boundaries.
//
// See: https://developer.apple.com/documentation/AppKit/NSText/copyRuler(_:)
func (t NSText) CopyRuler(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("copyRuler:"), sender)
}
// This action method pastes paragraph style information from the ruler
// pasteboard onto the selected paragraphs of a rich text object.
//
// # Discussion
// 
// It doesn’t apply to a plain text object.
//
// See: https://developer.apple.com/documentation/AppKit/NSText/pasteRuler(_:)
func (t NSText) PasteRuler(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("pasteRuler:"), sender)
}
// This action method deletes the selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSText/delete(_:)
func (t NSText) Delete(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("delete:"), sender)
}
// This action method changes the font of the selection for a rich text
// object, or of all text for a plain text object.
//
// # Discussion
// 
// If the receiver doesn’t use the Font panel, this method does nothing.
// 
// This method changes the font by sending a [ConvertFont] message to the
// shared NSFontManager and applying each NSFont returned to the appropriate
// text. See the [NSFontManager] class specification for more information on
// font conversion.
//
// See: https://developer.apple.com/documentation/AppKit/NSText/changeFont(_:)
func (t NSText) ChangeFont(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("changeFont:"), sender)
}
// Sets the font of characters within `aRange` to `aFont`.
//
// # Discussion
// 
// This method applies only to a rich text object.
// 
// This method does not include undo support by default. Clients must invoke
// [ShouldChangeTextInRangesReplacementStrings] or
// [ShouldChangeTextInRangeReplacementString] to include this method in an
// undoable action.
//
// See: https://developer.apple.com/documentation/AppKit/NSText/setFont(_:range:)
func (t NSText) SetFontRange(font NSFont, range_ foundation.NSRange) {
	objc.Send[objc.ID](t.ID, objc.Sel("setFont:range:"), font, range_)
}
// This action method applies center alignment to selected paragraphs (or all
// text if the receiver is a plain text object).
//
// See: https://developer.apple.com/documentation/AppKit/NSText/alignCenter(_:)
func (t NSText) AlignCenter(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("alignCenter:"), sender)
}
// This action method applies left alignment to selected paragraphs (or all
// text if the receiver is a plain text object).
//
// See: https://developer.apple.com/documentation/AppKit/NSText/alignLeft(_:)
func (t NSText) AlignLeft(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("alignLeft:"), sender)
}
// This action method applies right alignment to selected paragraphs (or all
// text if the receiver is a plain text object).
//
// See: https://developer.apple.com/documentation/AppKit/NSText/alignRight(_:)
func (t NSText) AlignRight(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("alignRight:"), sender)
}
// Sets the text color of characters within the specified range to the
// specified color.
//
// # Discussion
// 
// Removes the text color attribute if `color` is `nil`. This method applies
// only to rich text objects.
// 
// This method does not include undo support by default. Clients must invoke
// [ShouldChangeTextInRangesReplacementStrings] or
// [ShouldChangeTextInRangeReplacementString] to include this method in an
// undoable action.
//
// See: https://developer.apple.com/documentation/AppKit/NSText/setTextColor(_:range:)
func (t NSText) SetTextColorRange(color INSColor, range_ foundation.NSRange) {
	objc.Send[objc.ID](t.ID, objc.Sel("setTextColor:range:"), color, range_)
}
// This action method applies a superscript attribute to selected text (or all
// text if the receiver is a plain text object), raising its baseline offset
// by a predefined amount.
//
// See: https://developer.apple.com/documentation/AppKit/NSText/superscript(_:)
func (t NSText) Superscript(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("superscript:"), sender)
}
// This action method applies a subscript attribute to selected text (or all
// text if the receiver is a plain text object), lowering its baseline offset
// by a predefined amount.
//
// See: https://developer.apple.com/documentation/AppKit/NSText/subscript(_:)
func (t NSText) Subscript(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("subscript:"), sender)
}
// This action method removes any superscripting or subscripting from selected
// text (or all text if the receiver is a plain text object).
//
// See: https://developer.apple.com/documentation/AppKit/NSText/unscript(_:)
func (t NSText) Unscript(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("unscript:"), sender)
}
// Adds the underline attribute to the selected text attributes if absent;
// removes the attribute if present.
//
// # Discussion
// 
// If there is a selection and the first character of the selected range has
// any form of underline on it, or if there is no selection and the typing
// attributes have any form of underline, then underline is removed; otherwise
// a single simple underline is added.
// 
// Operates on the selected range if the receiver contains rich text. For
// plain text the range is the entire contents of the receiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSText/underline(_:)
func (t NSText) Underline(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("underline:"), sender)
}
// Attempts to read the RTFD file at the specified path.
//
// # Return Value
// 
// [true] if successful; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// `path` should be the path for an `XCUIElementTypeRtf` file or an
// `XCUIElementTypeRtfd` file wrapper, not for the RTF file within an
// `XCUIElementTypeRtfd` file wrapper.
//
// See: https://developer.apple.com/documentation/AppKit/NSText/readRTFD(fromFile:)
func (t NSText) ReadRTFDFromFile(path string) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("readRTFDFromFile:"), objc.String(path))
	return rv
}
// Writes the receiver’s text as RTF with attachments to a file or directory
// at `path`.
//
// # Discussion
// 
// Returns [true] on success and [false] on failure. If `atomicFlag` is
// [true], attempts to write the file safely so that an existing file at
// `path` is not overwritten, nor does a new file at `path` actually get
// created, unless the write is successful.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSText/writeRTFD(toFile:atomically:)
func (t NSText) WriteRTFDToFileAtomically(path string, flag bool) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("writeRTFDToFile:atomically:"), objc.String(path), flag)
	return rv
}
// Returns an NSData object that contains an RTFD stream corresponding to the
// characters and attributes within `aRange`.
//
// # Discussion
// 
// Raises an [NSRangeException] if any part of `aRange` lies beyond the end of
// the receiver’s characters.
// 
// When writing data to the pasteboard, you can use the NSData object as the
// first argument to [NSPasteboard]’s [SetDataForType] method, with a second
// argument of [NSRTFDPboardType].
//
// See: https://developer.apple.com/documentation/AppKit/NSText/rtfd(from:)
func (t NSText) RTFDFromRange(range_ foundation.NSRange) foundation.INSData {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("RTFDFromRange:"), range_)
	return foundation.NSDataFromID(rv)
}
// Returns an NSData object that contains an RTF stream corresponding to the
// characters and attributes within `aRange`, omitting any attachment
// characters and attributes.
//
// # Discussion
// 
// Raises an [NSRangeException] if any part of `aRange` lies beyond the end of
// the receiver’s characters.
// 
// When writing data to the pasteboard, you can use the NSData object as the
// first argument to [NSPasteboard]’s [SetDataForType] method, with a second
// argument of [NSRTFPboardType].
//
// See: https://developer.apple.com/documentation/AppKit/NSText/rtf(from:)
func (t NSText) RTFFromRange(range_ foundation.NSRange) foundation.INSData {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("RTFFromRange:"), range_)
	return foundation.NSDataFromID(rv)
}
// This action method searches for a misspelled word in the receiver’s text.
//
// # Discussion
// 
// The search starts at the end of the selection and continues until it
// reaches a word suspected of being misspelled or the end of the text. If a
// word isn’t recognized by the spelling server, a [ShowGuessPanel] message
// then opens the Guess panel and allows the user to make a correction or add
// the word to the local dictionary.
//
// See: https://developer.apple.com/documentation/AppKit/NSText/checkSpelling(_:)
func (t NSText) CheckSpelling(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("checkSpelling:"), sender)
}
// This action method opens the Spelling panel, allowing the user to make a
// correction during spell checking.
//
// See: https://developer.apple.com/documentation/AppKit/NSText/showGuessPanel(_:)
func (t NSText) ShowGuessPanel(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("showGuessPanel:"), sender)
}
// Resizes the receiver to fit its text.
//
// # Discussion
// 
// The text view will not be sized any smaller than its minimum size, however.
//
// See: https://developer.apple.com/documentation/AppKit/NSText/sizeToFit()
func (t NSText) SizeToFit() {
	objc.Send[objc.ID](t.ID, objc.Sel("sizeToFit"))
}
// Scrolls the receiver in its enclosing scroll view so the first characters
// of `aRange` are visible.
//
// See: https://developer.apple.com/documentation/AppKit/NSText/scrollRangeToVisible(_:)
func (t NSText) ScrollRangeToVisible(range_ foundation.NSRange) {
	objc.Send[objc.ID](t.ID, objc.Sel("scrollRangeToVisible:"), range_)
}
// Replaces the selected word in the receiver with a corrected version from
// the Spelling panel.
//
// # Discussion
// 
// This message is sent by the [NSSpellChecker] to the object whose text is
// being checked. To get the corrected spelling, ask `sender` for the string
// value of its selected cell (visible to the user as the text field in the
// Spelling panel). This method should replace the selected portion of the
// text with the string that it gets from the NSSpellChecker.
//
// See: https://developer.apple.com/documentation/AppKit/NSChangeSpelling/changeSpelling(_:)
func (t NSText) ChangeSpelling(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("changeSpelling:"), sender)
}
//
// # Discussion
// 
// Implement this action method to allow an application to ignore misspelled
// words on a document-by-document basis. This message is sent by the
// NSSpellChecker instance to the object whose text is being checked.
// 
// Implement this method by using the code shown in the protocol description.
//
// See: https://developer.apple.com/documentation/AppKit/NSIgnoreMisspelledWords/ignoreSpelling(_:)
func (t NSText) IgnoreSpelling(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("ignoreSpelling:"), sender)
}

// The receiver’s background color to a given color.
//
// See: https://developer.apple.com/documentation/AppKit/NSText/backgroundColor
func (t NSText) BackgroundColor() INSColor {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("backgroundColor"))
	return NSColorFromID(objc.ID(rv))
}
func (t NSText) SetBackgroundColor(value INSColor) {
	objc.Send[struct{}](t.ID, objc.Sel("setBackgroundColor:"), value)
}
// A Boolean that controls whether the receiver draws its background.
//
// # Discussion
// 
// If `flag` is [true], the receiver fills its background with the background
// color, if `flag` is [false], it doesn’t.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSText/drawsBackground
func (t NSText) DrawsBackground() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("drawsBackground"))
	return rv
}
func (t NSText) SetDrawsBackground(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setDrawsBackground:"), value)
}
// A Boolean that controls whether the receiver allows the user to edit its
// text.
//
// # Discussion
// 
// If `flag` is [true], the receiver allows the user to edit text and
// attributes; if `flag` is [false], it doesn’t.
// 
// You can change the receiver’s text programmatically regardless of this
// setting. If the receiver is made editable, it’s also made selectable.
// [NSText] objects are by default editable.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSText/isEditable
func (t NSText) Editable() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isEditable"))
	return rv
}
func (t NSText) SetEditable(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setEditable:"), value)
}
// A Boolean that controls whether the receiver allows the user to select its
// text.
//
// # Discussion
// 
// If `flag` is [true], the receiver allows the user to select text; if `flag`
// is [false], it doesn’t.
// 
// You can set selections programmatically regardless of this setting. If the
// receiver is made not selectable, it’s also made not editable. [NSText]
// objects are by default editable and selectable.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSText/isSelectable
func (t NSText) Selectable() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isSelectable"))
	return rv
}
func (t NSText) SetSelectable(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setSelectable:"), value)
}
// A Boolean that controls whether the receiver interprets Tab, Shift-Tab, and
// Return (Enter) as cues to end editing and possibly to change the first
// responder.
//
// # Discussion
// 
// If `flag` is [true], the receiver interprets Tab, Shift-Tab, and Return
// (Enter) as cues to end editing and possibly to change the first responder;
// if `flag` is [false], it doesn’t, instead accepting these characters as
// text input.
// 
// See the [NSWindow] class specification for more information on field
// editors. By default, [NSText] objects don’t behave as field editors.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSText/isFieldEditor
func (t NSText) FieldEditor() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isFieldEditor"))
	return rv
}
func (t NSText) SetFieldEditor(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setFieldEditor:"), value)
}
// A Boolean that controls whether the receiver allows the user to apply
// attributes to specific ranges of the text.
//
// # Discussion
// 
// If `flag` is [true] the receiver allows the user to apply attributes to
// specific ranges of the text; if `flag` is [false] it doesn’t.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSText/isRichText
func (t NSText) RichText() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isRichText"))
	return rv
}
func (t NSText) SetRichText(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setRichText:"), value)
}
// A Boolean that controls whether the receiver allows the user to import
// files by dragging.
//
// # Discussion
// 
// If `flag` is [true], the receiver allows the user to import files by
// dragging; if `flag` is [false], it doesn’t.
// 
// If the receiver is set to accept dragged files, it’s also made a rich
// text object. Subclasses may or may not accept dragged files by default.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSText/importsGraphics
func (t NSText) ImportsGraphics() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("importsGraphics"))
	return rv
}
func (t NSText) SetImportsGraphics(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setImportsGraphics:"), value)
}
// A Boolean that controls whether the receiver uses the Font panel and Font
// menu.
//
// # Discussion
// 
// If `flag` is [true], the receiver responds to messages from the Font panel
// and from the Font menu and updates the Font panel with the selection font
// whenever it changes. If `flag` is [false] the receiver doesn’t do any of
// these actions.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSText/usesFontPanel
func (t NSText) UsesFontPanel() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("usesFontPanel"))
	return rv
}
func (t NSText) SetUsesFontPanel(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setUsesFontPanel:"), value)
}
// A Boolean value that indicates whether the receiver’s enclosing scroll
// view shows its ruler.
//
// # Discussion
// 
// [true] if the receiver’s enclosing scroll view shows its ruler, otherwise
// [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSText/isRulerVisible
func (t NSText) RulerVisible() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isRulerVisible"))
	return rv
}
// The receiver’s characters within `aRange`.
//
// See: https://developer.apple.com/documentation/AppKit/NSText/selectedRange
func (t NSText) SelectedRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](t.ID, objc.Sel("selectedRange"))
	return foundation.NSRange(rv)
}
func (t NSText) SetSelectedRange(value foundation.NSRange) {
	objc.Send[struct{}](t.ID, objc.Sel("setSelectedRange:"), value)
}
// The font of all the receiver’s text.
//
// # Discussion
// 
// When the specified font doesn’t include a character, the text system uses
// an alternate font that contains the character. The substituted font may not
// have compatible metrics.
//
// See: https://developer.apple.com/documentation/AppKit/NSText/font
func (t NSText) Font() NSFont {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("font"))
	return NSFontFromID(objc.ID(rv))
}
func (t NSText) SetFont(value NSFont) {
	objc.Send[struct{}](t.ID, objc.Sel("setFont:"), value)
}
// The alignment of all the receiver’s text.
//
// # Discussion
// 
// The value of `mode` must be one of the alignments described in
// [NSTextAlignment].
// 
// Text using [NSNaturalTextAlignment] is actually displayed using one of the
// other alignments, depending on the natural alignment of the text’s
// script.
//
// [NSTextAlignment]: https://developer.apple.com/documentation/AppKit/NSTextAlignment
//
// See: https://developer.apple.com/documentation/AppKit/NSText/alignment
func (t NSText) Alignment() NSTextAlignment {
	rv := objc.Send[NSTextAlignment](t.ID, objc.Sel("alignment"))
	return NSTextAlignment(rv)
}
func (t NSText) SetAlignment(value NSTextAlignment) {
	objc.Send[struct{}](t.ID, objc.Sel("setAlignment:"), value)
}
// The text color of all characters in the receiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSText/textColor
func (t NSText) TextColor() INSColor {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("textColor"))
	return NSColorFromID(objc.ID(rv))
}
func (t NSText) SetTextColor(value INSColor) {
	objc.Send[struct{}](t.ID, objc.Sel("setTextColor:"), value)
}
// The initial writing direction used to determine the actual writing
// direction for text.
//
// # Discussion
// 
// The Text system uses this value as a hint for calculating the actual
// direction for displaying Unicode characters. If no writing direction is
// set, returns [NSWritingDirectionNatural].
//
// See: https://developer.apple.com/documentation/AppKit/NSText/baseWritingDirection
func (t NSText) BaseWritingDirection() NSWritingDirection {
	rv := objc.Send[NSWritingDirection](t.ID, objc.Sel("baseWritingDirection"))
	return NSWritingDirection(rv)
}
func (t NSText) SetBaseWritingDirection(value NSWritingDirection) {
	objc.Send[struct{}](t.ID, objc.Sel("setBaseWritingDirection:"), value)
}
// The receiver’s maximum size.
//
// See: https://developer.apple.com/documentation/AppKit/NSText/maxSize
func (t NSText) MaxSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](t.ID, objc.Sel("maxSize"))
	return corefoundation.CGSize(rv)
}
func (t NSText) SetMaxSize(value corefoundation.CGSize) {
	objc.Send[struct{}](t.ID, objc.Sel("setMaxSize:"), value)
}
// The receiver’s minimum size.
//
// See: https://developer.apple.com/documentation/AppKit/NSText/minSize
func (t NSText) MinSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](t.ID, objc.Sel("minSize"))
	return corefoundation.CGSize(rv)
}
func (t NSText) SetMinSize(value corefoundation.CGSize) {
	objc.Send[struct{}](t.ID, objc.Sel("setMinSize:"), value)
}
// A Boolean that controls whether the receiver changes its height to fit the
// height of its text.
//
// # Discussion
// 
// If `flag` is [true] it does; if `flag` is [false] it doesn’t.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSText/isVerticallyResizable
func (t NSText) VerticallyResizable() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isVerticallyResizable"))
	return rv
}
func (t NSText) SetVerticallyResizable(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setVerticallyResizable:"), value)
}
// A Boolean that controls whether the receiver changes its width to fit the
// width of its text.
//
// # Discussion
// 
// If `flag` is [true] it does; if `flag` is [false] it doesn’t
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSText/isHorizontallyResizable
func (t NSText) HorizontallyResizable() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isHorizontallyResizable"))
	return rv
}
func (t NSText) SetHorizontallyResizable(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setHorizontallyResizable:"), value)
}
// The receiver’s delegate.
//
// See: https://developer.apple.com/documentation/AppKit/NSText/delegate
func (t NSText) Delegate() NSTextDelegate {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("delegate"))
	return NSTextDelegateObjectFromID(rv)
}
func (t NSText) SetDelegate(value NSTextDelegate) {
	objc.Send[struct{}](t.ID, objc.Sel("setDelegate:"), value)
}

			// Protocol methods for NSChangeSpelling
			

			// Protocol methods for NSIgnoreMisspelledWords
			

