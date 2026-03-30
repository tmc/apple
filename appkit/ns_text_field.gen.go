// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSTextField] class.
var (
	_NSTextFieldClass     NSTextFieldClass
	_NSTextFieldClassOnce sync.Once
)

func getNSTextFieldClass() NSTextFieldClass {
	_NSTextFieldClassOnce.Do(func() {
		_NSTextFieldClass = NSTextFieldClass{class: objc.GetClass("NSTextField")}
	})
	return _NSTextFieldClass
}

// GetNSTextFieldClass returns the class object for NSTextField.
func GetNSTextFieldClass() NSTextFieldClass {
	return getNSTextFieldClass()
}

type NSTextFieldClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSTextFieldClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTextFieldClass) Alloc() NSTextField {
	rv := objc.Send[NSTextField](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// Text the user can select or edit to send an action message to a target when
// the user presses the Return key.
//
// # Overview
//
// The [NSTextField] class uses the [NSTextFieldCell] class to implement its
// user interface. Text fields display text either as a static label or as an
// editable input field. The content of a text field is either plain text or a
// rich-text attributed string. Text fields also support line wrapping to
// display multiline text, and a variety of truncation styles if the content
// doesn’t fit the available space.
//
// The parent class, [NSControl], provides the methods for setting the values
// of the text field, such as [NSTextField.StringValue] and [NSTextField.DoubleValue]. There are
// corresponding methods to retrieve values.
//
// In macOS 12 and later, if you explicitly call the `layoutManager` property
// on your text field, the framework will revert to a compatibility mode that
// uses [NSLayoutManager]. The text view also switches to this compatibility
// mode when it encounters text content that’s not yet supported.
//
// # Controlling Selection and Editing
//
//   - [NSTextField.Selectable]: A Boolean value that determines whether the user can select the content of the text field.
//   - [NSTextField.SetSelectable]
//   - [NSTextField.Editable]: A Boolean value that controls whether the user can edit the value in the text field.
//   - [NSTextField.SetEditable]
//
// # Controlling Rich Text Behavior
//
//   - [NSTextField.AllowsEditingTextAttributes]: A Boolean value that controls whether the user can change font attributes of the text field’s string.
//   - [NSTextField.SetAllowsEditingTextAttributes]
//   - [NSTextField.ImportsGraphics]: A Boolean value that controls whether the user can drag image files into the text field.
//   - [NSTextField.SetImportsGraphics]
//
// # Setting Placeholder Text
//
//   - [NSTextField.PlaceholderString]: The string the text field displays when empty to help the user understand the text field’s purpose.
//   - [NSTextField.SetPlaceholderString]
//   - [NSTextField.PlaceholderAttributedString]: The attributed string the text field displays when empty to help the user understand the text field’s purpose.
//   - [NSTextField.SetPlaceholderAttributedString]
//
// # Configuring Line Wrapping
//
//   - [NSTextField.LineBreakStrategy]: The strategy that the system uses to break lines when laying out multiple lines of text.
//   - [NSTextField.SetLineBreakStrategy]
//   - [NSTextField.AllowsDefaultTighteningForTruncation]: A Boolean value that controls whether single-line text fields tighten intercharacter spacing before truncating the text.
//   - [NSTextField.SetAllowsDefaultTighteningForTruncation]
//   - [NSTextField.MaximumNumberOfLines]: The maximum number of lines a wrapping text field displays before clipping or truncating the text.
//   - [NSTextField.SetMaximumNumberOfLines]
//
// # Sizing with Auto Layout
//
//   - [NSTextField.PreferredMaxLayoutWidth]: The maximum width of the text field’s intrinsic content size.
//   - [NSTextField.SetPreferredMaxLayoutWidth]
//
// # Setting the Text Color
//
//   - [NSTextField.TextColor]: The color of the text field’s content.
//   - [NSTextField.SetTextColor]
//
// # Controlling the Background
//
//   - [NSTextField.BackgroundColor]: The color of the background the text field’s cell draws behind the text.
//   - [NSTextField.SetBackgroundColor]
//   - [NSTextField.DrawsBackground]: A Boolean value that controls whether the text field’s cell draws a background color behind the text.
//   - [NSTextField.SetDrawsBackground]
//   - [NSTextField.Bezeled]: A Boolean value that controls whether the text field draws a bezeled background around its contents.
//   - [NSTextField.SetBezeled]
//   - [NSTextField.BezelStyle]: The text field’s bezel style, square or rounded.
//   - [NSTextField.SetBezelStyle]
//
// # Setting a Border
//
//   - [NSTextField.Bordered]: A Boolean value that controls whether the text field draws a solid black border around its contents.
//   - [NSTextField.SetBordered]
//
// # Selecting the Text
//
//   - [NSTextField.SelectText]: Ends editing in the text field and, if it’s selectable, selects the entire text content.
//
// # Using Keyboard Interface Control
//
//   - [NSTextField.AllowsCharacterPickerTouchBarItem]: A Boolean value that controls whether the Touch Bar displays the character picker item for rich text fields.
//   - [NSTextField.SetAllowsCharacterPickerTouchBarItem]
//
// # Supporting Text Completion and Suggestions
//
//   - [NSTextField.AutomaticTextCompletionEnabled]: A Boolean value that indicates whether the text field automatically completes text as the user types.
//   - [NSTextField.SetAutomaticTextCompletionEnabled]
//
// # Setting the Delegate
//
//   - [NSTextField.Delegate]: The text field’s delegate.
//   - [NSTextField.SetDelegate]
//
// # Implementing Delegate Methods
//
//   - [NSTextField.TextShouldBeginEditing]: Requests permission to begin editing a text object.
//   - [NSTextField.TextDidBeginEditing]: Posts a notification to the default notification center that the text is about to go into edit mode.
//   - [NSTextField.TextDidChange]: Posts a notification when the text changes, and forwards the message to the text field’s cell if it responds.
//   - [NSTextField.TextShouldEndEditing]: Performs validation on the text field’s new value.
//   - [NSTextField.TextDidEndEditing]: Posts a notification when the text is no longer in edit mode.
//
// # Instance Properties
//
//   - [NSTextField.AllowsWritingTools]
//   - [NSTextField.SetAllowsWritingTools]
//   - [NSTextField.AllowsWritingToolsAffordance]
//   - [NSTextField.SetAllowsWritingToolsAffordance]
//   - [NSTextField.PlaceholderAttributedStrings]
//   - [NSTextField.SetPlaceholderAttributedStrings]
//   - [NSTextField.PlaceholderStrings]
//   - [NSTextField.SetPlaceholderStrings]
//   - [NSTextField.ResolvesNaturalAlignmentWithBaseWritingDirection]: Specifies the behavior for resolving `NSTextAlignment.Natural()` to the visual alignment.
//   - [NSTextField.SetResolvesNaturalAlignmentWithBaseWritingDirection]
//   - [NSTextField.SuggestionsDelegate]: The delegate that provides text suggestions for the receiving text field and responds to the user highlighting and selecting items.
//   - [NSTextField.SetSuggestionsDelegate]
//
// See: https://developer.apple.com/documentation/AppKit/NSTextField
type NSTextField struct {
	NSControl
}

// NSTextFieldFromID constructs a [NSTextField] from an objc.ID.
//
// Text the user can select or edit to send an action message to a target when
// the user presses the Return key.
func NSTextFieldFromID(id objc.ID) NSTextField {
	return NSTextField{NSControl: NSControlFromID(id)}
}

// NOTE: NSTextField adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSTextField] class.
//
// # Controlling Selection and Editing
//
//   - [INSTextField.Selectable]: A Boolean value that determines whether the user can select the content of the text field.
//   - [INSTextField.SetSelectable]
//   - [INSTextField.Editable]: A Boolean value that controls whether the user can edit the value in the text field.
//   - [INSTextField.SetEditable]
//
// # Controlling Rich Text Behavior
//
//   - [INSTextField.AllowsEditingTextAttributes]: A Boolean value that controls whether the user can change font attributes of the text field’s string.
//   - [INSTextField.SetAllowsEditingTextAttributes]
//   - [INSTextField.ImportsGraphics]: A Boolean value that controls whether the user can drag image files into the text field.
//   - [INSTextField.SetImportsGraphics]
//
// # Setting Placeholder Text
//
//   - [INSTextField.PlaceholderString]: The string the text field displays when empty to help the user understand the text field’s purpose.
//   - [INSTextField.SetPlaceholderString]
//   - [INSTextField.PlaceholderAttributedString]: The attributed string the text field displays when empty to help the user understand the text field’s purpose.
//   - [INSTextField.SetPlaceholderAttributedString]
//
// # Configuring Line Wrapping
//
//   - [INSTextField.LineBreakStrategy]: The strategy that the system uses to break lines when laying out multiple lines of text.
//   - [INSTextField.SetLineBreakStrategy]
//   - [INSTextField.AllowsDefaultTighteningForTruncation]: A Boolean value that controls whether single-line text fields tighten intercharacter spacing before truncating the text.
//   - [INSTextField.SetAllowsDefaultTighteningForTruncation]
//   - [INSTextField.MaximumNumberOfLines]: The maximum number of lines a wrapping text field displays before clipping or truncating the text.
//   - [INSTextField.SetMaximumNumberOfLines]
//
// # Sizing with Auto Layout
//
//   - [INSTextField.PreferredMaxLayoutWidth]: The maximum width of the text field’s intrinsic content size.
//   - [INSTextField.SetPreferredMaxLayoutWidth]
//
// # Setting the Text Color
//
//   - [INSTextField.TextColor]: The color of the text field’s content.
//   - [INSTextField.SetTextColor]
//
// # Controlling the Background
//
//   - [INSTextField.BackgroundColor]: The color of the background the text field’s cell draws behind the text.
//   - [INSTextField.SetBackgroundColor]
//   - [INSTextField.DrawsBackground]: A Boolean value that controls whether the text field’s cell draws a background color behind the text.
//   - [INSTextField.SetDrawsBackground]
//   - [INSTextField.Bezeled]: A Boolean value that controls whether the text field draws a bezeled background around its contents.
//   - [INSTextField.SetBezeled]
//   - [INSTextField.BezelStyle]: The text field’s bezel style, square or rounded.
//   - [INSTextField.SetBezelStyle]
//
// # Setting a Border
//
//   - [INSTextField.Bordered]: A Boolean value that controls whether the text field draws a solid black border around its contents.
//   - [INSTextField.SetBordered]
//
// # Selecting the Text
//
//   - [INSTextField.SelectText]: Ends editing in the text field and, if it’s selectable, selects the entire text content.
//
// # Using Keyboard Interface Control
//
//   - [INSTextField.AllowsCharacterPickerTouchBarItem]: A Boolean value that controls whether the Touch Bar displays the character picker item for rich text fields.
//   - [INSTextField.SetAllowsCharacterPickerTouchBarItem]
//
// # Supporting Text Completion and Suggestions
//
//   - [INSTextField.AutomaticTextCompletionEnabled]: A Boolean value that indicates whether the text field automatically completes text as the user types.
//   - [INSTextField.SetAutomaticTextCompletionEnabled]
//
// # Setting the Delegate
//
//   - [INSTextField.Delegate]: The text field’s delegate.
//   - [INSTextField.SetDelegate]
//
// # Implementing Delegate Methods
//
//   - [INSTextField.TextShouldBeginEditing]: Requests permission to begin editing a text object.
//   - [INSTextField.TextDidBeginEditing]: Posts a notification to the default notification center that the text is about to go into edit mode.
//   - [INSTextField.TextDidChange]: Posts a notification when the text changes, and forwards the message to the text field’s cell if it responds.
//   - [INSTextField.TextShouldEndEditing]: Performs validation on the text field’s new value.
//   - [INSTextField.TextDidEndEditing]: Posts a notification when the text is no longer in edit mode.
//
// # Instance Properties
//
//   - [INSTextField.AllowsWritingTools]
//   - [INSTextField.SetAllowsWritingTools]
//   - [INSTextField.AllowsWritingToolsAffordance]
//   - [INSTextField.SetAllowsWritingToolsAffordance]
//   - [INSTextField.PlaceholderAttributedStrings]
//   - [INSTextField.SetPlaceholderAttributedStrings]
//   - [INSTextField.PlaceholderStrings]
//   - [INSTextField.SetPlaceholderStrings]
//   - [INSTextField.ResolvesNaturalAlignmentWithBaseWritingDirection]: Specifies the behavior for resolving `NSTextAlignment.Natural()` to the visual alignment.
//   - [INSTextField.SetResolvesNaturalAlignmentWithBaseWritingDirection]
//   - [INSTextField.SuggestionsDelegate]: The delegate that provides text suggestions for the receiving text field and responds to the user highlighting and selecting items.
//   - [INSTextField.SetSuggestionsDelegate]
//
// See: https://developer.apple.com/documentation/AppKit/NSTextField
type INSTextField interface {
	INSControl
	NSAccessibilityNavigableStaticText
	NSAccessibilityStaticText
	NSTextContent
	NSUserInterfaceValidations

	// Topic: Controlling Selection and Editing

	// A Boolean value that determines whether the user can select the content of the text field.
	Selectable() bool
	SetSelectable(value bool)
	// A Boolean value that controls whether the user can edit the value in the text field.
	Editable() bool
	SetEditable(value bool)

	// Topic: Controlling Rich Text Behavior

	// A Boolean value that controls whether the user can change font attributes of the text field’s string.
	AllowsEditingTextAttributes() bool
	SetAllowsEditingTextAttributes(value bool)
	// A Boolean value that controls whether the user can drag image files into the text field.
	ImportsGraphics() bool
	SetImportsGraphics(value bool)

	// Topic: Setting Placeholder Text

	// The string the text field displays when empty to help the user understand the text field’s purpose.
	PlaceholderString() string
	SetPlaceholderString(value string)
	// The attributed string the text field displays when empty to help the user understand the text field’s purpose.
	PlaceholderAttributedString() foundation.NSAttributedString
	SetPlaceholderAttributedString(value foundation.NSAttributedString)

	// Topic: Configuring Line Wrapping

	// The strategy that the system uses to break lines when laying out multiple lines of text.
	LineBreakStrategy() NSLineBreakStrategy
	SetLineBreakStrategy(value NSLineBreakStrategy)
	// A Boolean value that controls whether single-line text fields tighten intercharacter spacing before truncating the text.
	AllowsDefaultTighteningForTruncation() bool
	SetAllowsDefaultTighteningForTruncation(value bool)
	// The maximum number of lines a wrapping text field displays before clipping or truncating the text.
	MaximumNumberOfLines() int
	SetMaximumNumberOfLines(value int)

	// Topic: Sizing with Auto Layout

	// The maximum width of the text field’s intrinsic content size.
	PreferredMaxLayoutWidth() float64
	SetPreferredMaxLayoutWidth(value float64)

	// Topic: Setting the Text Color

	// The color of the text field’s content.
	TextColor() INSColor
	SetTextColor(value INSColor)

	// Topic: Controlling the Background

	// The color of the background the text field’s cell draws behind the text.
	BackgroundColor() INSColor
	SetBackgroundColor(value INSColor)
	// A Boolean value that controls whether the text field’s cell draws a background color behind the text.
	DrawsBackground() bool
	SetDrawsBackground(value bool)
	// A Boolean value that controls whether the text field draws a bezeled background around its contents.
	Bezeled() bool
	SetBezeled(value bool)
	// The text field’s bezel style, square or rounded.
	BezelStyle() NSTextFieldBezelStyle
	SetBezelStyle(value NSTextFieldBezelStyle)

	// Topic: Setting a Border

	// A Boolean value that controls whether the text field draws a solid black border around its contents.
	Bordered() bool
	SetBordered(value bool)

	// Topic: Selecting the Text

	// Ends editing in the text field and, if it’s selectable, selects the entire text content.
	SelectText(sender objectivec.IObject)

	// Topic: Using Keyboard Interface Control

	// A Boolean value that controls whether the Touch Bar displays the character picker item for rich text fields.
	AllowsCharacterPickerTouchBarItem() bool
	SetAllowsCharacterPickerTouchBarItem(value bool)

	// Topic: Supporting Text Completion and Suggestions

	// A Boolean value that indicates whether the text field automatically completes text as the user types.
	AutomaticTextCompletionEnabled() bool
	SetAutomaticTextCompletionEnabled(value bool)

	// Topic: Setting the Delegate

	// The text field’s delegate.
	Delegate() NSTextFieldDelegate
	SetDelegate(value NSTextFieldDelegate)

	// Topic: Implementing Delegate Methods

	// Requests permission to begin editing a text object.
	TextShouldBeginEditing(textObject INSText) bool
	// Posts a notification to the default notification center that the text is about to go into edit mode.
	TextDidBeginEditing(notification foundation.NSNotification)
	// Posts a notification when the text changes, and forwards the message to the text field’s cell if it responds.
	TextDidChange(notification foundation.NSNotification)
	// Performs validation on the text field’s new value.
	TextShouldEndEditing(textObject INSText) bool
	// Posts a notification when the text is no longer in edit mode.
	TextDidEndEditing(notification foundation.NSNotification)

	// Topic: Instance Properties

	AllowsWritingTools() bool
	SetAllowsWritingTools(value bool)
	AllowsWritingToolsAffordance() bool
	SetAllowsWritingToolsAffordance(value bool)
	PlaceholderAttributedStrings() []foundation.NSAttributedString
	SetPlaceholderAttributedStrings(value []foundation.NSAttributedString)
	PlaceholderStrings() []string
	SetPlaceholderStrings(value []string)
	// Specifies the behavior for resolving `NSTextAlignment.Natural()` to the visual alignment.
	ResolvesNaturalAlignmentWithBaseWritingDirection() bool
	SetResolvesNaturalAlignmentWithBaseWritingDirection(value bool)
	// The delegate that provides text suggestions for the receiving text field and responds to the user highlighting and selecting items.
	SuggestionsDelegate() objectivec.IObject
	SetSuggestionsDelegate(value objectivec.IObject)
}

// Init initializes the instance.
func (t NSTextField) Init() NSTextField {
	rv := objc.Send[NSTextField](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTextField) Autorelease() NSTextField {
	rv := objc.Send[NSTextField](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTextField creates a new NSTextField instance.
func NewNSTextField() NSTextField {
	class := getNSTextFieldClass()
	rv := objc.Send[NSTextField](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a text field for use as a static label that displays styled text,
// doesn’t wrap, and doesn’t have selectable text.
//
// attributedStringValue: An attributed string to use as the content of the label.
//
// # Return Value
//
// A text field that displays the specified attributed string as a static
// label.
//
// # Discussion
//
// The text field determines its line-break mode by inspecting the paragraph
// style attributes in the attributed string.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextField/init(labelWithAttributedString:)
func NewTextFieldLabelWithAttributedString(attributedStringValue foundation.NSAttributedString) NSTextField {
	rv := objc.Send[objc.ID](objc.ID(getNSTextFieldClass().class), objc.Sel("labelWithAttributedString:"), attributedStringValue)
	return NSTextFieldFromID(rv)
}

// Initializes a text field for use as a static label that uses the system
// default font, doesn’t wrap, and doesn’t have selectable text.
//
// stringValue: A string to use as the content of the label.
//
// # Return Value
//
// A text field that displays the specified string as a static label.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextField/init(labelWithString:)
func NewTextFieldLabelWithString(stringValue string) NSTextField {
	rv := objc.Send[objc.ID](objc.ID(getNSTextFieldClass().class), objc.Sel("labelWithString:"), objc.String(stringValue))
	return NSTextFieldFromID(rv)
}

// Initializes a control with data in an unarchiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/init(coder:)
func NewTextFieldWithCoder(coder foundation.INSCoder) NSTextField {
	instance := getNSTextFieldClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSTextFieldFromID(rv)
}

// Initializes a control with the specified frame rectangle.
//
// frameRect: The rectangle of the control, specified in points in the coordinate space
// of the enclosing view.
//
// # Return Value
//
// An initialized control object, or `nil` if the object couldn’t be
// initialized.
//
// # Discussion
//
// If a cell has been specified for controls of this type, this method also
// creates an instance of the cell. Because [NSControl] is an abstract class,
// invocations of this method should appear only in the designated
// initializers of subclasses; that is, there should always be a more specific
// designated initializer for the subclass, as this method is the designated
// initializer for [NSControl].
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/init(frame:)
func NewTextFieldWithFrame(frameRect corefoundation.CGRect) NSTextField {
	instance := getNSTextFieldClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSTextFieldFromID(rv)
}

// Initializes a single-line editable text field for user input using the
// system default font and standard visual appearance.
//
// stringValue: A string to use as the initial content of the editable text field.
//
// # Return Value
//
// A single-line editable text field that displays the specified string.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextField/init(string:)
func NewTextFieldWithString(stringValue string) NSTextField {
	rv := objc.Send[objc.ID](objc.ID(getNSTextFieldClass().class), objc.Sel("textFieldWithString:"), objc.String(stringValue))
	return NSTextFieldFromID(rv)
}

// Initializes a text field for use as a multiline static label with
// selectable text that uses the system default font.
//
// stringValue: A string to use as the initial content of the editable text field.
//
// # Return Value
//
// A multiline text field that displays the specified string.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextField/init(wrappingLabelWithString:)
func NewTextFieldWrappingLabelWithString(stringValue string) NSTextField {
	rv := objc.Send[objc.ID](objc.ID(getNSTextFieldClass().class), objc.Sel("wrappingLabelWithString:"), objc.String(stringValue))
	return NSTextFieldFromID(rv)
}

// Ends editing in the text field and, if it’s selectable, selects the
// entire text content.
//
// sender: The sender of the message.
//
// # Discussion
//
// If the text field isn’t in a window’s view hierarchy, this method has
// no effect.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextField/selectText(_:)
func (t NSTextField) SelectText(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("selectText:"), sender)
}

// Requests permission to begin editing a text object.
//
// textObject: The object to begin editing.
//
// # Return Value
//
// true if editing can begin; otherwise, false.
//
// # Discussion
//
// If the text field isn’t editable, this method returns false immediately.
// If the text field is editable and its delegate responds to
// [ControlTextShouldBeginEditing], this method invokes that method and
// returns the result. Otherwise, it returns true to allow editing to occur.
// See [NSControl] for more information about the text delegate method.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextField/textShouldBeginEditing(_:)
func (t NSTextField) TextShouldBeginEditing(textObject INSText) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("textShouldBeginEditing:"), textObject)
	return rv
}

// Posts a notification to the default notification center that the text is
// about to go into edit mode.
//
// notification: The [TextDidBeginEditingNotification] notification to post to the default
// notification center.
//
// # Discussion
//
// This action causes the text field’s delegate to receive a
// [controlTextDidBeginEditing:] message. See [NSControl] for more information
// about the text delegate method.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextField/textDidBeginEditing(_:)
//
// [controlTextDidBeginEditing:]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/controlTextDidBeginEditing:
func (t NSTextField) TextDidBeginEditing(notification foundation.NSNotification) {
	objc.Send[objc.ID](t.ID, objc.Sel("textDidBeginEditing:"), notification)
}

// Posts a notification when the text changes, and forwards the message to the
// text field’s cell if it responds.
//
// notification: The [TextDidChangeNotification] notification to post to the default
// notification center.
//
// # Discussion
//
// This method causes the text field’s delegate to receive a
// [controlTextDidChange:] message. See the [NSControl] class specification
// for more information about the text delegate method.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextField/textDidChange(_:)
//
// [controlTextDidChange:]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/controlTextDidChange:
func (t NSTextField) TextDidChange(notification foundation.NSNotification) {
	objc.Send[objc.ID](t.ID, objc.Sel("textDidChange:"), notification)
}

// Performs validation on the text field’s new value.
//
// textObject: The text object that requests permission to end editing.
//
// # Return Value
//
// true if the new value is valid; otherwise, false.
//
// # Discussion
//
// This method validates the text field’s new value using the [NSCell]
// method [isEntryAcceptable:]. If the new value is valid and the delegate
// responds to [ControlTextShouldEndEditing], this method invokes that method
// and returns the result. If the delegate returns false, the system beeps to
// indicate that the text field can’t validate the text. See [NSControl] for
// more information about the text delegate method.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextField/textShouldEndEditing(_:)
//
// [isEntryAcceptable:]: https://developer.apple.com/documentation/AppKit/NSCell/isEntryAcceptable:
func (t NSTextField) TextShouldEndEditing(textObject INSText) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("textShouldEndEditing:"), textObject)
	return rv
}

// Posts a notification when the text is no longer in edit mode.
//
// notification: The [textDidEndEditingNotification] to post to the default notification
// center.
//
// # Discussion
//
// After validating the new value, this method posts a
// [textDidEndEditingNotification] to the default notification center, which
// causes the text field’s delegate to receive a [controlTextDidEndEditing:]
// message. This method then sends [EndEditing] to the text field’s cell and
// handles the key that causes editing to end as follows:
//
// - If the user ends editing by pressing Return, this method tries to send
// the text field’s action to its target. If unsuccessful, it sends
// [PerformKeyEquivalent] to its [NSView], for example, to handle the default
// button on a panel. If that also fails, the text field selects its text. -
// If the user ends editing by pressing Tab or Shift-Tab, the text field tries
// to have its [NSWindow] object select its next or previous key view, using
// the [NSWindow] method [SelectKeyViewFollowingView] or
// [SelectKeyViewPrecedingView]. If unsuccessful, the text field selects its
// text.
//
// See [NSControl] for more information about the text delegate method.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextField/textDidEndEditing(_:)
//
// [textDidEndEditingNotification]: https://developer.apple.com/documentation/AppKit/NSControl/textDidEndEditingNotification
// [controlTextDidEndEditing:]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/controlTextDidEndEditing:
//
// [textDidEndEditingNotification]: https://developer.apple.com/documentation/AppKit/NSControl/textDidEndEditingNotification
func (t NSTextField) TextDidEndEditing(notification foundation.NSNotification) {
	objc.Send[objc.ID](t.ID, objc.Sel("textDidEndEditing:"), notification)
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
func (t NSTextField) AccessibilityAttributedStringForRange(range_ foundation.NSRange) foundation.NSAttributedString {
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
func (t NSTextField) AccessibilityFrameForRange(range_ foundation.NSRange) corefoundation.CGRect {
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
func (t NSTextField) AccessibilityLineForIndex(index int) int {
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
func (t NSTextField) AccessibilityRangeForLine(lineNumber int) foundation.NSRange {
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
func (t NSTextField) AccessibilityStringForRange(range_ foundation.NSRange) string {
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
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityStaticText/accessibilityValue()
//
// [accessibilityValue]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityValue
func (t NSTextField) AccessibilityValue() string {
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
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityStaticText/accessibilityVisibleCharacterRange()
//
// [accessibilityVisibleCharacterRange]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityVisibleCharacterRange
func (t NSTextField) AccessibilityVisibleCharacterRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](t.ID, objc.Sel("accessibilityVisibleCharacterRange"))
	return foundation.NSRange(rv)
}

// Returns a Boolean value that indicates whether the sender should be
// enabled.
//
// item: The user interface item to validate. You can send `anItem` the [Action] and
// [Tag] messages.
//
// # Return Value
//
// true if the user interface item should be enabled, otherwise false.
//
// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceValidations/validateUserInterfaceItem(_:)
func (t NSTextField) ValidateUserInterfaceItem(item NSValidatedUserInterfaceItem) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("validateUserInterfaceItem:"), item)
	return rv
}

// A Boolean value that determines whether the user can select the content of
// the text field.
//
// # Discussion
//
// If true, the text field becomes selectable but not editable. Use [Editable]
// to make the text field selectable and editable. If false, the text is
// neither editable nor selectable.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextField/isSelectable
func (t NSTextField) Selectable() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isSelectable"))
	return rv
}
func (t NSTextField) SetSelectable(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setSelectable:"), value)
}

// A Boolean value that controls whether the user can edit the value in the
// text field.
//
// # Discussion
//
// If true, the user can select and edit text. If false, the user can’t edit
// text, and the ability to select the text field’s content is dependent on
// the value of [Selectable].
//
// For example, if an [NSTextField] object is selectable but uneditable,
// becomes editable for a time, and then becomes uneditable again, it remains
// selectable. To ensure that text is neither editable nor selectable, use
// [Selectable] to disable text selection.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextField/isEditable
func (t NSTextField) Editable() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isEditable"))
	return rv
}
func (t NSTextField) SetEditable(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setEditable:"), value)
}

// A Boolean value that controls whether the user can change font attributes
// of the text field’s string.
//
// # Discussion
//
// If true, and the text value is an attributed string, the text field’s
// content displays using the attributed string’s visual settings. The user
// can modify the text field’s style attributes in the font panel.
//
// If false, and the text is an attributed string, the text field ignores
// style attributes, such as font and color. The text field’s content
// displays according to the text field’s settings. The text field ignores
// changes to the attributed string’s attributes when displaying the string
// and when the text field is in edit mode.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextField/allowsEditingTextAttributes
func (t NSTextField) AllowsEditingTextAttributes() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("allowsEditingTextAttributes"))
	return rv
}
func (t NSTextField) SetAllowsEditingTextAttributes(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setAllowsEditingTextAttributes:"), value)
}

// A Boolean value that controls whether the user can drag image files into
// the text field.
//
// # Discussion
//
// If true, the text field accepts dragged images; if false, it doesn’t. You
// can add images programmatically regardless of this setting.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextField/importsGraphics
func (t NSTextField) ImportsGraphics() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("importsGraphics"))
	return rv
}
func (t NSTextField) SetImportsGraphics(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setImportsGraphics:"), value)
}

// The string the text field displays when empty to help the user understand
// the text field’s purpose.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextField/placeholderString
func (t NSTextField) PlaceholderString() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("placeholderString"))
	return foundation.NSStringFromID(rv).String()
}
func (t NSTextField) SetPlaceholderString(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setPlaceholderString:"), objc.String(value))
}

// The attributed string the text field displays when empty to help the user
// understand the text field’s purpose.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextField/placeholderAttributedString
func (t NSTextField) PlaceholderAttributedString() foundation.NSAttributedString {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("placeholderAttributedString"))
	return foundation.NSAttributedStringFromID(objc.ID(rv))
}
func (t NSTextField) SetPlaceholderAttributedString(value foundation.NSAttributedString) {
	objc.Send[struct{}](t.ID, objc.Sel("setPlaceholderAttributedString:"), value)
}

// The strategy that the system uses to break lines when laying out multiple
// lines of text.
//
// # Discussion
//
// The default value for editable text fields is [NSLineBreakStrategyNone] to
// match the field editor’s behavior. The default value for selectable,
// uneditable text fields is [NSLineBreakStrategyStandard].
//
// See: https://developer.apple.com/documentation/AppKit/NSTextField/lineBreakStrategy
func (t NSTextField) LineBreakStrategy() NSLineBreakStrategy {
	rv := objc.Send[NSLineBreakStrategy](t.ID, objc.Sel("lineBreakStrategy"))
	return NSLineBreakStrategy(rv)
}
func (t NSTextField) SetLineBreakStrategy(value NSLineBreakStrategy) {
	objc.Send[struct{}](t.ID, objc.Sel("setLineBreakStrategy:"), value)
}

// A Boolean value that controls whether single-line text fields tighten
// intercharacter spacing before truncating the text.
//
// # Discussion
//
// The text field ignores this property when its value is an attributed
// string.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextField/allowsDefaultTighteningForTruncation
func (t NSTextField) AllowsDefaultTighteningForTruncation() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("allowsDefaultTighteningForTruncation"))
	return rv
}
func (t NSTextField) SetAllowsDefaultTighteningForTruncation(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setAllowsDefaultTighteningForTruncation:"), value)
}

// The maximum number of lines a wrapping text field displays before clipping
// or truncating the text.
//
// # Discussion
//
// The default value of `0` indicates no limit to the number of lines, and the
// text fills the bounds of the text field cell.
//
// If the text field reaches the maximum number of lines, or if the height of
// the container can’t accommodate the number of lines, the text field clips
// or truncates the text, depending on the cell’s [TruncatesLastVisibleLine]
// setting.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextField/maximumNumberOfLines
func (t NSTextField) MaximumNumberOfLines() int {
	rv := objc.Send[int](t.ID, objc.Sel("maximumNumberOfLines"))
	return rv
}
func (t NSTextField) SetMaximumNumberOfLines(value int) {
	objc.Send[struct{}](t.ID, objc.Sel("setMaximumNumberOfLines:"), value)
}

// The maximum width of the text field’s intrinsic content size.
//
// # Discussion
//
// If the text field wraps, the intrinsic height is large enough to show the
// entire text contents at that width.
//
// The default value is `0`, indicating no maximum preferred width.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextField/preferredMaxLayoutWidth
func (t NSTextField) PreferredMaxLayoutWidth() float64 {
	rv := objc.Send[float64](t.ID, objc.Sel("preferredMaxLayoutWidth"))
	return rv
}
func (t NSTextField) SetPreferredMaxLayoutWidth(value float64) {
	objc.Send[struct{}](t.ID, objc.Sel("setPreferredMaxLayoutWidth:"), value)
}

// The color of the text field’s content.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/AppKit/NSTextField/textColor
func (t NSTextField) TextColor() INSColor {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("textColor"))
	return NSColorFromID(objc.ID(rv))
}
func (t NSTextField) SetTextColor(value INSColor) {
	objc.Send[struct{}](t.ID, objc.Sel("setTextColor:"), value)
}

// The color of the background the text field’s cell draws behind the text.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextField/backgroundColor
func (t NSTextField) BackgroundColor() INSColor {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("backgroundColor"))
	return NSColorFromID(objc.ID(rv))
}
func (t NSTextField) SetBackgroundColor(value INSColor) {
	objc.Send[struct{}](t.ID, objc.Sel("setBackgroundColor:"), value)
}

// A Boolean value that controls whether the text field’s cell draws a
// background color behind the text.
//
// # Discussion
//
// If true, the text field’s cell draws a background; if false, it draws
// nothing behind the text.
//
// To prevent inconsistent rendering, [NSTextField] disables background color
// rendering for text fields with rounded bezels.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextField/drawsBackground
func (t NSTextField) DrawsBackground() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("drawsBackground"))
	return rv
}
func (t NSTextField) SetDrawsBackground(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setDrawsBackground:"), value)
}

// A Boolean value that controls whether the text field draws a bezeled
// background around its contents.
//
// # Discussion
//
// If true, the text field draws a bezel and sets [DrawsBackground] to false;
// if false, it doesn’t draw a bezeled background.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextField/isBezeled
func (t NSTextField) Bezeled() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isBezeled"))
	return rv
}
func (t NSTextField) SetBezeled(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setBezeled:"), value)
}

// The text field’s bezel style, square or rounded.
//
// # Discussion
//
// To enable a bezel for a text field, set [Bezeled] to true, then set the
// bezel style. See [NSTextField.BezelStyle] for available bezel styles.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextField/bezelStyle-swift.property
//
// [NSTextField.BezelStyle]: https://developer.apple.com/documentation/AppKit/NSTextField/BezelStyle-swift.enum
func (t NSTextField) BezelStyle() NSTextFieldBezelStyle {
	rv := objc.Send[NSTextFieldBezelStyle](t.ID, objc.Sel("bezelStyle"))
	return NSTextFieldBezelStyle(rv)
}
func (t NSTextField) SetBezelStyle(value NSTextFieldBezelStyle) {
	objc.Send[struct{}](t.ID, objc.Sel("setBezelStyle:"), value)
}

// A Boolean value that controls whether the text field draws a solid black
// border around its contents.
//
// # Discussion
//
// If true, the text field draws a border; if false, it doesn’t draw a
// border.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextField/isBordered
func (t NSTextField) Bordered() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isBordered"))
	return rv
}
func (t NSTextField) SetBordered(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setBordered:"), value)
}

// A Boolean value that indicates whether the text field automatically
// completes text as the user types.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextField/isAutomaticTextCompletionEnabled
func (t NSTextField) AutomaticTextCompletionEnabled() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isAutomaticTextCompletionEnabled"))
	return rv
}
func (t NSTextField) SetAutomaticTextCompletionEnabled(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setAutomaticTextCompletionEnabled:"), value)
}

// The text field’s delegate.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextField/delegate
func (t NSTextField) Delegate() NSTextFieldDelegate {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("delegate"))
	return NSTextFieldDelegateObjectFromID(rv)
}
func (t NSTextField) SetDelegate(value NSTextFieldDelegate) {
	objc.Send[struct{}](t.ID, objc.Sel("setDelegate:"), value)
}

// A Boolean value that controls whether the Touch Bar displays the character
// picker item for rich text fields.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextField/allowsCharacterPickerTouchBarItem
func (t NSTextField) AllowsCharacterPickerTouchBarItem() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("allowsCharacterPickerTouchBarItem"))
	return rv
}
func (t NSTextField) SetAllowsCharacterPickerTouchBarItem(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setAllowsCharacterPickerTouchBarItem:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSTextField/allowsWritingTools
func (t NSTextField) AllowsWritingTools() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("allowsWritingTools"))
	return rv
}
func (t NSTextField) SetAllowsWritingTools(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setAllowsWritingTools:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSTextField/allowsWritingToolsAffordance
func (t NSTextField) AllowsWritingToolsAffordance() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("allowsWritingToolsAffordance"))
	return rv
}
func (t NSTextField) SetAllowsWritingToolsAffordance(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setAllowsWritingToolsAffordance:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSTextField/placeholderAttributedStrings
func (t NSTextField) PlaceholderAttributedStrings() []foundation.NSAttributedString {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("placeholderAttributedStrings"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSAttributedString {
		return foundation.NSAttributedStringFromID(id)
	})
}
func (t NSTextField) SetPlaceholderAttributedStrings(value []foundation.NSAttributedString) {
	objc.Send[struct{}](t.ID, objc.Sel("setPlaceholderAttributedStrings:"), objectivec.IObjectSliceToNSArray(value))
}

// See: https://developer.apple.com/documentation/AppKit/NSTextField/placeholderStrings
func (t NSTextField) PlaceholderStrings() []string {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("placeholderStrings"))
	return objc.ConvertSliceToStrings(rv)
}
func (t NSTextField) SetPlaceholderStrings(value []string) {
	objc.Send[struct{}](t.ID, objc.Sel("setPlaceholderStrings:"), objectivec.StringSliceToNSArray(value))
}

// Specifies the behavior for resolving `NSTextAlignment.Natural()` to the
// visual alignment.
//
// # Discussion
//
// When set to `true`, the resolved visual alignment is determined by the
// resolved base writing direction; otherwise, it is using the user’s
// preferred language. The default value is `false`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextField/resolvesNaturalAlignmentWithBaseWritingDirection
func (t NSTextField) ResolvesNaturalAlignmentWithBaseWritingDirection() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("resolvesNaturalAlignmentWithBaseWritingDirection"))
	return rv
}
func (t NSTextField) SetResolvesNaturalAlignmentWithBaseWritingDirection(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setResolvesNaturalAlignmentWithBaseWritingDirection:"), value)
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
// See: https://developer.apple.com/documentation/AppKit/NSTextContent/contentType
//
// [emailAddress]: https://developer.apple.com/documentation/AppKit/NSTextContentType/emailAddress
func (t NSTextField) ContentType() NSTextContentType {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("contentType"))
	return NSTextContentType(foundation.NSStringFromID(rv).String())
}
func (t NSTextField) SetContentType(value NSTextContentType) {
	objc.Send[struct{}](t.ID, objc.Sel("setContentType:"), objc.String(string(value)))
}

// The delegate that provides text suggestions for the receiving text field
// and responds to the user highlighting and selecting items.
//
// See: https://developer.apple.com/documentation/appkit/nstextfield/suggestionsdelegate
func (t NSTextField) SuggestionsDelegate() objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("suggestionsDelegate"))
	return objectivec.Object{ID: rv}
}
func (t NSTextField) SetSuggestionsDelegate(value objectivec.IObject) {
	objc.Send[struct{}](t.ID, objc.Sel("setSuggestionsDelegate:"), value)
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
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityFrame()
//
// [accessibilityFrame]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFrame
// [position]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/position
// [size]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/size
func (o NSTextField) AccessibilityFrame() corefoundation.CGRect {
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
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityParent()
//
// [accessibilityParent]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityParent
func (o NSTextField) AccessibilityParent() objectivec.IObject {
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
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityIdentifier()
//
// [accessibilityIdentifier]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityIdentifier
func (o NSTextField) AccessibilityIdentifier() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityIdentifier"))
	return foundation.NSStringFromID(rv).String()
}

// Returns a Boolean value that indicates whether the accessibility element
// has the keyboard focus.
//
// # Return Value
//
// true if this element has the keyboard focus; otherwise, false.
//
// # Discussion
//
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityFocused] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/isAccessibilityFocused()
//
// [accessibilityFocused]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFocused
func (o NSTextField) IsAccessibilityFocused() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFocused"))
	return rv
}

// Protocol methods for NSTextContent

// Protocol methods for NSUserInterfaceValidations
