// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSAlert] class.
var (
	_NSAlertClass     NSAlertClass
	_NSAlertClassOnce sync.Once
)

func getNSAlertClass() NSAlertClass {
	_NSAlertClassOnce.Do(func() {
		_NSAlertClass = NSAlertClass{class: objc.GetClass("NSAlert")}
	})
	return _NSAlertClass
}

// GetNSAlertClass returns the class object for NSAlert.
func GetNSAlertClass() NSAlertClass {
	return getNSAlertClass()
}

type NSAlertClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSAlertClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSAlertClass) Alloc() NSAlert {
	rv := objc.Send[NSAlert](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A modal dialog or sheet attached to a document window.
//
// # Overview
//
// The methods of the [NSAlert] class allow you to specify alert level, alert
// text, button titles, and a custom icon should you require it. The class
// also lets your alerts display a help button and provides ways for apps to
// offer help specific to an alert.
//
// To display an alert as a sheet, call the
// [NSAlert.BeginSheetModalForWindowCompletionHandler] method; to display one as an
// app-modal dialog, use the [NSAlert.RunModal] method.
//
// By design, an [NSAlert] object is intended for a single alert—that is, an
// alert with a unique combination of title, buttons, and so on—that is
// displayed upon a particular condition. You should create an [NSAlert]
// object for each alert dialog, creating it only when you need to display an
// alert, and release it when you are done. If you have a particular alert
// dialog that you need to show repeatedly, you can retain and reuse an
// instance of [NSAlert] for this dialog.
//
// After creating an alert using one of the alert creation methods, you can
// customize it further prior to displaying it by customizing its attributes.
// See [NSAlert].
//
// Unless you must maintain compatibility with existing alert-processing code
// that uses the function-based API, you should allocate (`alloc`) and
// initialize (`init`) the alert object, and then set its attributes using the
// appropriate methods of the [NSAlert] class.
//
// # Instance Attributes
//
// [NSAlert] objects have the following attributes:
//
// - Type An alert’s type helps convey the importance or gravity of its
// message to the user. Specified with the [NSAlert.AlertStyle] property. - Message
// text The main message of the alert. Specified with [NSAlert.MessageText]. -
// Informative text Additional information about the alert. Specified with
// [NSAlert.InformativeText]. - Icon An optional, custom icon to display in the alert,
// which is used instead of the default app icon. Specified with [NSAlert.Icon]. -
// Help Alerts can let the user get help about them. Use [NSAlert.HelpAnchor] and
// [NSAlert.ShowsHelp]. - Response buttons By default an alert has one response
// button: the OK button. You can add more response buttons using the
// [NSAlert.AddButtonWithTitle] method. - Suppression checkbox A suppression checkbox
// allows the user to suppress the display of a particular alert in subsequent
// occurrences of the event that triggers it. Use [NSAlert.ShowsSuppressionButton]. -
// Accessory view An accessory view lets you add additional information to an
// alert; for example, a text field with contact information. Use
// [NSAlert.AccessoryView], [NSAlert.Layout].
//
// # Subclassing Notes
//
// The [NSAlert] class is not designed for subclassing.
//
// # Configuring Alerts
//
//   - [NSAlert.Layout]: Specifies that the alert must do immediate layout instead of lazily just before display.
//   - [NSAlert.AlertStyle]: Indicates the alert’s severity level.
//   - [NSAlert.SetAlertStyle]
//   - [NSAlert.AccessoryView]: The alert’s accessory view.
//   - [NSAlert.SetAccessoryView]
//   - [NSAlert.ShowsHelp]: Specifies whether the alert has a help button.
//   - [NSAlert.SetShowsHelp]
//   - [NSAlert.HelpAnchor]: The alert’s HTML help anchor.
//   - [NSAlert.SetHelpAnchor]
//   - [NSAlert.Delegate]: The alert’s delegate.
//   - [NSAlert.SetDelegate]
//
// # Displaying Alerts
//
//   - [NSAlert.RunModal]: Runs the alert as an app-modal dialog and returns the constant that identifies the button clicked.
//   - [NSAlert.BeginSheetModalForWindowCompletionHandler]: Runs the alert modally as a sheet attached to the specified window.
//   - [NSAlert.SuppressionButton]: The alert’s suppression checkbox.
//   - [NSAlert.ShowsSuppressionButton]: Specifies whether the alert includes a suppression checkbox, which you can employ to allow a user to opt out of seeing the alert again.
//   - [NSAlert.SetShowsSuppressionButton]
//
// # Accessing Alert Text
//
//   - [NSAlert.InformativeText]: The alert’s informative text.
//   - [NSAlert.SetInformativeText]
//   - [NSAlert.MessageText]: The alert’s message text or title.
//   - [NSAlert.SetMessageText]
//
// # Accessing a Custom Alert Icon
//
//   - [NSAlert.Icon]: The custom icon displayed in the alert.
//   - [NSAlert.SetIcon]
//
// # Accessing Alert Response Buttons
//
//   - [NSAlert.Buttons]: The array of response buttons for the alert.
//   - [NSAlert.AddButtonWithTitle]: Adds a button with a given title to the alert.
//
// # Getting Alert Windows
//
//   - [NSAlert.Window]: The app-modal panel or document-modal sheet that corresponds to the alert.
//
// See: https://developer.apple.com/documentation/AppKit/NSAlert
type NSAlert struct {
	objectivec.Object
}

// NSAlertFromID constructs a [NSAlert] from an objc.ID.
//
// A modal dialog or sheet attached to a document window.
func NSAlertFromID(id objc.ID) NSAlert {
	return NSAlert{objectivec.Object{ID: id}}
}

// NOTE: NSAlert adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSAlert] class.
//
// # Configuring Alerts
//
//   - [INSAlert.Layout]: Specifies that the alert must do immediate layout instead of lazily just before display.
//   - [INSAlert.AlertStyle]: Indicates the alert’s severity level.
//   - [INSAlert.SetAlertStyle]
//   - [INSAlert.AccessoryView]: The alert’s accessory view.
//   - [INSAlert.SetAccessoryView]
//   - [INSAlert.ShowsHelp]: Specifies whether the alert has a help button.
//   - [INSAlert.SetShowsHelp]
//   - [INSAlert.HelpAnchor]: The alert’s HTML help anchor.
//   - [INSAlert.SetHelpAnchor]
//   - [INSAlert.Delegate]: The alert’s delegate.
//   - [INSAlert.SetDelegate]
//
// # Displaying Alerts
//
//   - [INSAlert.RunModal]: Runs the alert as an app-modal dialog and returns the constant that identifies the button clicked.
//   - [INSAlert.BeginSheetModalForWindowCompletionHandler]: Runs the alert modally as a sheet attached to the specified window.
//   - [INSAlert.SuppressionButton]: The alert’s suppression checkbox.
//   - [INSAlert.ShowsSuppressionButton]: Specifies whether the alert includes a suppression checkbox, which you can employ to allow a user to opt out of seeing the alert again.
//   - [INSAlert.SetShowsSuppressionButton]
//
// # Accessing Alert Text
//
//   - [INSAlert.InformativeText]: The alert’s informative text.
//   - [INSAlert.SetInformativeText]
//   - [INSAlert.MessageText]: The alert’s message text or title.
//   - [INSAlert.SetMessageText]
//
// # Accessing a Custom Alert Icon
//
//   - [INSAlert.Icon]: The custom icon displayed in the alert.
//   - [INSAlert.SetIcon]
//
// # Accessing Alert Response Buttons
//
//   - [INSAlert.Buttons]: The array of response buttons for the alert.
//   - [INSAlert.AddButtonWithTitle]: Adds a button with a given title to the alert.
//
// # Getting Alert Windows
//
//   - [INSAlert.Window]: The app-modal panel or document-modal sheet that corresponds to the alert.
//
// See: https://developer.apple.com/documentation/AppKit/NSAlert
type INSAlert interface {
	objectivec.IObject

	// Topic: Configuring Alerts

	// Specifies that the alert must do immediate layout instead of lazily just before display.
	Layout()
	// Indicates the alert’s severity level.
	AlertStyle() NSAlertStyle
	SetAlertStyle(value NSAlertStyle)
	// The alert’s accessory view.
	AccessoryView() INSView
	SetAccessoryView(value INSView)
	// Specifies whether the alert has a help button.
	ShowsHelp() bool
	SetShowsHelp(value bool)
	// The alert’s HTML help anchor.
	HelpAnchor() NSHelpAnchorName
	SetHelpAnchor(value NSHelpAnchorName)
	// The alert’s delegate.
	Delegate() NSAlertDelegate
	SetDelegate(value NSAlertDelegate)

	// Topic: Displaying Alerts

	// Runs the alert as an app-modal dialog and returns the constant that identifies the button clicked.
	RunModal() NSModalResponse
	// Runs the alert modally as a sheet attached to the specified window.
	BeginSheetModalForWindowCompletionHandler(sheetWindow INSWindow, handler ModalResponseHandler)
	// The alert’s suppression checkbox.
	SuppressionButton() INSButton
	// Specifies whether the alert includes a suppression checkbox, which you can employ to allow a user to opt out of seeing the alert again.
	ShowsSuppressionButton() bool
	SetShowsSuppressionButton(value bool)

	// Topic: Accessing Alert Text

	// The alert’s informative text.
	InformativeText() string
	SetInformativeText(value string)
	// The alert’s message text or title.
	MessageText() string
	SetMessageText(value string)

	// Topic: Accessing a Custom Alert Icon

	// The custom icon displayed in the alert.
	Icon() INSImage
	SetIcon(value INSImage)

	// Topic: Accessing Alert Response Buttons

	// The array of response buttons for the alert.
	Buttons() []NSButton
	// Adds a button with a given title to the alert.
	AddButtonWithTitle(title string) INSButton

	// Topic: Getting Alert Windows

	// The app-modal panel or document-modal sheet that corresponds to the alert.
	Window() INSWindow
}

// Init initializes the instance.
func (a NSAlert) Init() NSAlert {
	rv := objc.Send[NSAlert](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a NSAlert) Autorelease() NSAlert {
	rv := objc.Send[NSAlert](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSAlert creates a new NSAlert instance.
func NewNSAlert() NSAlert {
	class := getNSAlertClass()
	rv := objc.Send[NSAlert](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns an alert initialized from information in an error object.
//
// error: Error information to display.
//
// # Return Value
//
// An initialized alert.
//
// # Discussion
//
// The [NSAlert] class extracts the localized error description, recovery
// suggestion, and recovery options from the `error` parameter and uses them
// as the alert’s message text, informative text, and button titles,
// respectively.
//
// See: https://developer.apple.com/documentation/AppKit/NSAlert/init(error:)
func NewAlertWithError(error_ foundation.INSError) NSAlert {
	rv := objc.Send[objc.ID](objc.ID(getNSAlertClass().class), objc.Sel("alertWithError:"), error_)
	return NSAlertFromID(rv)
}

// Specifies that the alert must do immediate layout instead of lazily just
// before display.
//
// # Discussion
//
// You need to call this method only when you need to customize the alert’s
// layout. Call this method after all the alert’s attributes have been
// customized, including the suppression checkbox and the accessory layout.
// After the method returns, you can make the necessary layout changes; for
// example, adjusting the frame of the accessory view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAlert/layout()
func (a NSAlert) Layout() {
	objc.Send[objc.ID](a.ID, objc.Sel("layout"))
}

// Runs the alert as an app-modal dialog and returns the constant that
// identifies the button clicked.
//
// # Return Value
//
// A response to the alert. See this method’s “Special Considerations”
// section for details.
//
// # Discussion
//
// You can create the alert either through the standard allocation and
// initialization procedure or, if necessary in your app, by using the
// deprecated compatibility method
// [alertWithMessageText:defaultButton:alternateButton:otherButton:informativeTextWithFormat:].
//
// # Special Considerations
//
// This method can return values other than those specific to the alert
// buttons ([alertFirstButtonReturn], [alertSecondButtonReturn], and so on) if
// the alert is canceled programatically.
//
// If you use “ to create an alert, the [NSAlertDefaultReturn],
// [NSAlertAlternateReturn], and [NSAlertOtherReturn] constants identify the
// button used to dismiss the alert. Otherwise, the constants used are the
// ones described in [AddButtonWithTitle].
//
// See: https://developer.apple.com/documentation/AppKit/NSAlert/runModal()
//
// [alertFirstButtonReturn]: https://developer.apple.com/documentation/AppKit/NSApplication/ModalResponse/alertFirstButtonReturn
// [alertSecondButtonReturn]: https://developer.apple.com/documentation/AppKit/NSApplication/ModalResponse/alertSecondButtonReturn
// [alertWithMessageText:defaultButton:alternateButton:otherButton:informativeTextWithFormat:]: https://developer.apple.com/documentation/AppKit/NSAlert/alertWithMessageText:defaultButton:alternateButton:otherButton:informativeTextWithFormat:
func (a NSAlert) RunModal() NSModalResponse {
	rv := objc.Send[NSModalResponse](a.ID, objc.Sel("runModal"))
	return NSModalResponse(rv)
}

// Runs the alert modally as a sheet attached to the specified window.
//
// sheetWindow: The window on which to display the sheet.
//
// handler: The completion handler that gets called when the sheet’s modal session
// ends.
//
// # Discussion
//
// This method uses the [NSWindow] sheet methods to display the alert (for
// more information, see Managing Sheets). If the alert has an alert style of
// [NSCriticalAlertStyle], it is presented as a critical sheet, which means
// that it can display on top of other sheets that might already be attached
// to the window. Otherwise, it is presented—or queued for presentation—as
// a standard sheet.
//
// Note that [OrderOut] no longer needs to be called in the completion
// handler. If you don’t dismiss the alert, it will be done for you after
// the completion handler finishes.
//
// See: https://developer.apple.com/documentation/AppKit/NSAlert/beginSheetModal(for:completionHandler:)
//
// [NSCriticalAlertStyle]: https://developer.apple.com/documentation/AppKit/NSCriticalAlertStyle
func (a NSAlert) BeginSheetModalForWindowCompletionHandler(sheetWindow INSWindow, handler ModalResponseHandler) {
	_block1, _ := NewModalResponseBlock(handler)
	objc.Send[objc.ID](a.ID, objc.Sel("beginSheetModalForWindow:completionHandler:"), sheetWindow, _block1)
}

// Adds a button with a given title to the alert.
//
// title: Title of the button to add to the alert. Must not be `nil`.
//
// # Return Value
//
// Button added to the alert.
//
// # Discussion
//
// Buttons are placed starting near the right side of the alert and going
// toward the left side (for languages that read left to right). The first
// three buttons are identified positionally as [NSAlertFirstButtonReturn],
// [NSAlertSecondButtonReturn], [NSAlertThirdButtonReturn] in the return-code
// parameter evaluated by the modal delegate. Subsequent buttons are
// identified as [NSAlertThirdButtonReturn] +`n`, where `n` is an integer
//
// By default, the first button has a key equivalent of Return, any button
// with a title of “Cancel” has a key equivalent of Escape, and any button
// with the title “Don’t Save” has a key equivalent of Command-D (but
// only if it’s the first button). You can also assign different key
// equivalents for the buttons using the [KeyEquivalent] method of the
// [NSButton] class. In addition, you can use the [Tag] method of the
// [NSButton] class to set the return value.
//
// See: https://developer.apple.com/documentation/AppKit/NSAlert/addButton(withTitle:)
func (a NSAlert) AddButtonWithTitle(title string) INSButton {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("addButtonWithTitle:"), objc.String(title))
	return NSButtonFromID(rv)
}

// Indicates the alert’s severity level.
//
// # Discussion
//
// See the [NSAlert.Style] enumeration for the list of alert style constants.
//
// See: https://developer.apple.com/documentation/AppKit/NSAlert/alertStyle
//
// [NSAlert.Style]: https://developer.apple.com/documentation/AppKit/NSAlert/Style
func (a NSAlert) AlertStyle() NSAlertStyle {
	rv := objc.Send[NSAlertStyle](a.ID, objc.Sel("alertStyle"))
	return NSAlertStyle(rv)
}
func (a NSAlert) SetAlertStyle(value NSAlertStyle) {
	objc.Send[struct{}](a.ID, objc.Sel("setAlertStyle:"), value)
}

// The alert’s accessory view.
//
// # Discussion
//
// The [NSAlert] class places the accessory view between the informative text
// or suppression checkbox (if present) and the response buttons. Before you
// change the location of the accessory view, first call the [Layout] method.
//
// [AlertStyle] shows an example of adding an accessory view to an alert.
// [Buttons] shows the alert generated.
//
// Listing 1. Adding an accessory view to an alert
//
// [media-1965585]
//
// See: https://developer.apple.com/documentation/AppKit/NSAlert/accessoryView
func (a NSAlert) AccessoryView() INSView {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessoryView"))
	return NSViewFromID(objc.ID(rv))
}
func (a NSAlert) SetAccessoryView(value INSView) {
	objc.Send[struct{}](a.ID, objc.Sel("setAccessoryView:"), value)
}

// Specifies whether the alert has a help button.
//
// # Discussion
//
// Set this property’s value to true to specify that the alert has a help
// button, or false to specify it does not.
//
// When a user clicks an alert’s help button, the alert delegate
// ([Delegate]) receives an [AlertShowHelp] message. The delegate is
// responsible for displaying the help information related to this particular
// alert.
//
// Clicking an alert’s help button can alternately cause the
// [OpenHelpAnchorInBook] message to be sent to the app’s help manager with
// a `nil` book and the anchor specified by the [HelpAnchor] property, if any
// of the following conditions are true:
//
// - There is no alert delegate. - The alert delegate does not implement
// [AlertShowHelp]. - The alert delegate implements [AlertShowHelp] but
// returns false. When this is the case, an exception is raised if no help
// anchor is set.
//
// See: https://developer.apple.com/documentation/AppKit/NSAlert/showsHelp
func (a NSAlert) ShowsHelp() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("showsHelp"))
	return rv
}
func (a NSAlert) SetShowsHelp(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setShowsHelp:"), value)
}

// The alert’s HTML help anchor.
//
// # Discussion
//
// To provide a help anchor for the alert, set this property to the
// appropriate string value. To remove the help anchor, set this property’s
// value to `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSAlert/helpAnchor
func (a NSAlert) HelpAnchor() NSHelpAnchorName {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("helpAnchor"))
	return NSHelpAnchorName(foundation.NSStringFromID(rv).String())
}
func (a NSAlert) SetHelpAnchor(value NSHelpAnchorName) {
	objc.Send[struct{}](a.ID, objc.Sel("setHelpAnchor:"), objc.String(string(value)))
}

// The alert’s delegate.
//
// # Discussion
//
// To set a delegate for the alert, provide an object conforming to the
// [NSAlertDelegate]protocol to this property. To remove the delegate, set
// this property’s value to `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSAlert/delegate
func (a NSAlert) Delegate() NSAlertDelegate {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("delegate"))
	return NSAlertDelegateObjectFromID(rv)
}
func (a NSAlert) SetDelegate(value NSAlertDelegate) {
	objc.Send[struct{}](a.ID, objc.Sel("setDelegate:"), value)
}

// The alert’s suppression checkbox.
//
// # Discussion
//
// If you want to customize an alert’s suppression checkbox, access it via
// this property and then use the methods of the [NSButton] class. For
// example, you can do this to change the suppression checkbox’s default
// message, or to change its initial selection state (which is
// “unselected” by default). For a code example, see the
// [ShowsSuppressionButton] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSAlert/suppressionButton
func (a NSAlert) SuppressionButton() INSButton {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("suppressionButton"))
	return NSButtonFromID(objc.ID(rv))
}

// Specifies whether the alert includes a suppression checkbox, which you can
// employ to allow a user to opt out of seeing the alert again.
//
// # Discussion
//
// The default value of this property is false, which specifies the absence of
// a suppression checkbox in the alert. Set the value to true to show a
// suppression checkbox in the alert.
//
// By default, a suppression checkbox has the title, “Do not show this
// message again.” In macOS 11.0 and later, if the alert displays multiple
// buttons that prompt the user to make a choice, the title is “Do not ask
// again.” To customize it, use the checkbox’s `title` property, as
// follows:
//
// To create an alert that responds to the selection state of the suppression
// checkbox, use code like that shown in Listing 1 to produce the alert shown
// below.
//
// Listing 1. Creating an alert with a suppression checkbox
//
// [media-3686600]
//
// See: https://developer.apple.com/documentation/AppKit/NSAlert/showsSuppressionButton
func (a NSAlert) ShowsSuppressionButton() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("showsSuppressionButton"))
	return rv
}
func (a NSAlert) SetShowsSuppressionButton(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setShowsSuppressionButton:"), value)
}

// The alert’s informative text.
//
// # Discussion
//
// The value of this property must not be `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSAlert/informativeText
func (a NSAlert) InformativeText() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("informativeText"))
	return foundation.NSStringFromID(rv).String()
}
func (a NSAlert) SetInformativeText(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setInformativeText:"), objc.String(value))
}

// The alert’s message text or title.
//
// See: https://developer.apple.com/documentation/AppKit/NSAlert/messageText
func (a NSAlert) MessageText() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("messageText"))
	return foundation.NSStringFromID(rv).String()
}
func (a NSAlert) SetMessageText(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setMessageText:"), objc.String(value))
}

// The custom icon displayed in the alert.
//
// # Discussion
//
// By default, the image used in an alert is the app icon. If you set this
// property’s value, your specified custom image is used in place of the app
// icon.
//
// If you’ve set a custom alert icon, you can clear it by setting this
// property’s value to `nil`, which restores use of the app icon for the
// alert.
//
// See: https://developer.apple.com/documentation/AppKit/NSAlert/icon
func (a NSAlert) Icon() INSImage {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("icon"))
	return NSImageFromID(objc.ID(rv))
}
func (a NSAlert) SetIcon(value INSImage) {
	objc.Send[struct{}](a.ID, objc.Sel("setIcon:"), value)
}

// The array of response buttons for the alert.
//
// # Discussion
//
// The button displayed rightmost in the alert (for a left-to-right language)
// corresponds to the button at index `0` in this property’s array, and is
// considered the default button. A user can invoke this button by pressing
// the Return key.
//
// Any button with a title of “Cancel” has a key equivalent of Escape, and
// any button with the title “Don’t Save” has a key equivalent of
// Command-D (but only if it is not the first button). You can also assign
// different key equivalents for the buttons using the [KeyEquivalent] method
// of the [NSButton] class.
//
// See: https://developer.apple.com/documentation/AppKit/NSAlert/buttons
func (a NSAlert) Buttons() []NSButton {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("buttons"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSButton {
		return NSButtonFromID(id)
	})
}

// The app-modal panel or document-modal sheet that corresponds to the alert.
//
// # Discussion
//
// The alert’s window is of type [NSPanel]. Use this property when you want
// to dismiss an alert created with the
// [BeginSheetModalForWindowCompletionHandler] method within that method’s
// completion handler block.
//
// See: https://developer.apple.com/documentation/AppKit/NSAlert/window
func (a NSAlert) Window() INSWindow {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("window"))
	return NSWindowFromID(objc.ID(rv))
}

// BeginSheetModalForWindow is a synchronous wrapper around [NSAlert.BeginSheetModalForWindowCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (a NSAlert) BeginSheetModalForWindow(ctx context.Context, sheetWindow INSWindow) (NSModalResponse, error) {
	done := make(chan NSModalResponse, 1)
	a.BeginSheetModalForWindowCompletionHandler(sheetWindow, func(val NSModalResponse) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}
