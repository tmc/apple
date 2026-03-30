// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [UNTextInputNotificationAction] class.
var (
	_UNTextInputNotificationActionClass     UNTextInputNotificationActionClass
	_UNTextInputNotificationActionClassOnce sync.Once
)

func getUNTextInputNotificationActionClass() UNTextInputNotificationActionClass {
	_UNTextInputNotificationActionClassOnce.Do(func() {
		_UNTextInputNotificationActionClass = UNTextInputNotificationActionClass{class: objc.GetClass("UNTextInputNotificationAction")}
	})
	return _UNTextInputNotificationActionClass
}

// GetUNTextInputNotificationActionClass returns the class object for UNTextInputNotificationAction.
func GetUNTextInputNotificationActionClass() UNTextInputNotificationActionClass {
	return getUNTextInputNotificationActionClass()
}

type UNTextInputNotificationActionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (uc UNTextInputNotificationActionClass) Class() objc.Class {
	return uc.class
}

// Alloc allocates memory for a new instance of the class.
func (uc UNTextInputNotificationActionClass) Alloc() UNTextInputNotificationAction {
	rv := objc.Send[UNTextInputNotificationAction](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// An action that accepts user-typed text.
//
// # Overview
//
// Use [UNTextInputNotificationAction] objects to define an action that allows
// the user to provide a custom text-based response. When the user selects an
// action of this type, the system displays controls for the user to enter or
// dictate the text content. That text is then included in the response object
// that’s delivered to your app.
//
// For information on how to define actions and categories, see [Declaring
// your actionable notification types].
//
// # Getting Information
//
//   - [UNTextInputNotificationAction.TextInputButtonTitle]: The localized title of the text input button that the system displays to the user.
//   - [UNTextInputNotificationAction.TextInputPlaceholder]: The placeholder text that the system localizes and displays in the text input field.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNTextInputNotificationAction
//
// [Declaring your actionable notification types]: https://developer.apple.com/documentation/UserNotifications/declaring-your-actionable-notification-types
type UNTextInputNotificationAction struct {
	UNNotificationAction
}

// UNTextInputNotificationActionFromID constructs a [UNTextInputNotificationAction] from an objc.ID.
//
// An action that accepts user-typed text.
func UNTextInputNotificationActionFromID(id objc.ID) UNTextInputNotificationAction {
	return UNTextInputNotificationAction{UNNotificationAction: UNNotificationActionFromID(id)}
}

// NOTE: UNTextInputNotificationAction adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [UNTextInputNotificationAction] class.
//
// # Getting Information
//
//   - [IUNTextInputNotificationAction.TextInputButtonTitle]: The localized title of the text input button that the system displays to the user.
//   - [IUNTextInputNotificationAction.TextInputPlaceholder]: The placeholder text that the system localizes and displays in the text input field.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNTextInputNotificationAction
type IUNTextInputNotificationAction interface {
	IUNNotificationAction

	// Topic: Getting Information

	// The localized title of the text input button that the system displays to the user.
	TextInputButtonTitle() string
	// The placeholder text that the system localizes and displays in the text input field.
	TextInputPlaceholder() string
}

// Init initializes the instance.
func (u UNTextInputNotificationAction) Init() UNTextInputNotificationAction {
	rv := objc.Send[UNTextInputNotificationAction](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u UNTextInputNotificationAction) Autorelease() UNTextInputNotificationAction {
	rv := objc.Send[UNTextInputNotificationAction](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewUNTextInputNotificationAction creates a new UNTextInputNotificationAction instance.
func NewUNTextInputNotificationAction() UNTextInputNotificationAction {
	class := getUNTextInputNotificationActionClass()
	rv := objc.Send[UNTextInputNotificationAction](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an action object by using the specified title and options.
//
// identifier: The string that you use internally to identify the action. This string must
// be unique among your app’s supported actions. When the user selects the
// action, the system passes this string to your app and asks the user to
// perform the related task. This parameter must not be `nil` or an empty
// string.
//
// title: The localized string the system displays to the user. The system displays
// this string as the title of a button, which the system adds to the
// notification interface. This parameter must not be `nil`.
//
// options: Additional options that describe how the action behaves. Include options
// when you need the related behavior. For a list of possible values, see
// [UNNotificationActionOptions].
//
// # Return Value
//
// An action object that the system initializes.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationAction/init(identifier:title:options:)
//
// [UNNotificationActionOptions]: https://developer.apple.com/documentation/UserNotifications/UNNotificationActionOptions
func NewUNTextInputNotificationActionWithIdentifierTitleOptions(identifier string, title string, options UNNotificationActionOptions) UNTextInputNotificationAction {
	rv := objc.Send[objc.ID](objc.ID(getUNTextInputNotificationActionClass().class), objc.Sel("actionWithIdentifier:title:options:"), objc.String(identifier), objc.String(title), options)
	return UNTextInputNotificationActionFromID(rv)
}

// Creates an action object by using the specified title, options, and icon.
//
// identifier: The string that you use internally to identify the action. This string must
// be unique among your app’s supported actions. When the user selects the
// action, the system passes this string to your app and asks the user to
// perform the related task. This parameter must not be `nil` or an empty
// string.
//
// title: The localized string the system displays to the user. The system displays
// this string as the title of a button, which the system adds to the
// notification interface. This parameter must not be `nil`.
//
// options: Additional options that describe how the action behaves. Include options
// when you need the related behavior. For a list of possible values, see
// [UNNotificationActionOptions].
//
// icon: The icon that the system displays to the user.
//
// # Return Value
//
// An action object that the system initializes.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationAction/init(identifier:title:options:icon:)
//
// [UNNotificationActionOptions]: https://developer.apple.com/documentation/UserNotifications/UNNotificationActionOptions
func NewUNTextInputNotificationActionWithIdentifierTitleOptionsIcon(identifier string, title string, options UNNotificationActionOptions, icon IUNNotificationActionIcon) UNTextInputNotificationAction {
	rv := objc.Send[objc.ID](objc.ID(getUNTextInputNotificationActionClass().class), objc.Sel("actionWithIdentifier:title:options:icon:"), objc.String(identifier), objc.String(title), options, icon)
	return UNTextInputNotificationActionFromID(rv)
}

// Creates an action object with an icon that accepts text input from the
// user.
//
// identifier: The string that you use internally to identify the action. This string must
// be unique among all of your app’s supported actions. When the user
// selects the action, the system passes this string to your app and asks you
// to perform the related task. This parameter must not be `nil` or an empty
// string.
//
// title: The localized string the system displays to the user. The system displays
// this string as the title of a button, which the system adds to the
// notification interface. This parameter must not be `nil`.
//
// options: Additional options describing how the action behaves. Include options when
// you need the related behavior. For a list of possible values, see
// [UNNotificationActionOptions].
//
// icon: The icon to display to the user.
//
// textInputButtonTitle: The localized title of the text input button that’s displayed to the
// user.
//
// textInputPlaceholder: The localized placeholder text to display in the text input field.
//
// # Return Value
//
// A new text input action object.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNTextInputNotificationAction/init(identifier:title:options:icon:textInputButtonTitle:textInputPlaceholder:)
//
// [UNNotificationActionOptions]: https://developer.apple.com/documentation/UserNotifications/UNNotificationActionOptions
func NewUNTextInputNotificationActionWithIdentifierTitleOptionsIconTextInputButtonTitleTextInputPlaceholder(identifier string, title string, options UNNotificationActionOptions, icon IUNNotificationActionIcon, textInputButtonTitle string, textInputPlaceholder string) UNTextInputNotificationAction {
	rv := objc.Send[objc.ID](objc.ID(getUNTextInputNotificationActionClass().class), objc.Sel("actionWithIdentifier:title:options:icon:textInputButtonTitle:textInputPlaceholder:"), objc.String(identifier), objc.String(title), options, icon, objc.String(textInputButtonTitle), objc.String(textInputPlaceholder))
	return UNTextInputNotificationActionFromID(rv)
}

// Creates an action object that accepts text input from the user.
//
// identifier: The string that you use internally to identify the action. This string must
// be unique among all of your app’s supported actions. When the user
// selects the action, the system passes this string to your app and asks you
// to perform the related task. This parameter must not be `nil` or an empty
// string.
//
// title: The localized string the system displays to the user. The system displays
// this string as the title of a button, which the system adds to the
// notification interface. This parameter must not be `nil`.
//
// options: Additional options describing how the action behaves. Include options when
// you need the related behavior. For a list of possible values, see
// [UNNotificationActionOptions].
//
// textInputButtonTitle: The localized title of the text input button that’s displayed to the
// user.
//
// textInputPlaceholder: The localized placeholder text the system displays in the text input field.
//
// # Return Value
//
// A new text input action object.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNTextInputNotificationAction/init(identifier:title:options:textInputButtonTitle:textInputPlaceholder:)
//
// [UNNotificationActionOptions]: https://developer.apple.com/documentation/UserNotifications/UNNotificationActionOptions
func NewUNTextInputNotificationActionWithIdentifierTitleOptionsTextInputButtonTitleTextInputPlaceholder(identifier string, title string, options UNNotificationActionOptions, textInputButtonTitle string, textInputPlaceholder string) UNTextInputNotificationAction {
	rv := objc.Send[objc.ID](objc.ID(getUNTextInputNotificationActionClass().class), objc.Sel("actionWithIdentifier:title:options:textInputButtonTitle:textInputPlaceholder:"), objc.String(identifier), objc.String(title), options, objc.String(textInputButtonTitle), objc.String(textInputPlaceholder))
	return UNTextInputNotificationActionFromID(rv)
}

// The localized title of the text input button that the system displays to
// the user.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNTextInputNotificationAction/textInputButtonTitle
func (u UNTextInputNotificationAction) TextInputButtonTitle() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("textInputButtonTitle"))
	return foundation.NSStringFromID(rv).String()
}

// The placeholder text that the system localizes and displays in the text
// input field.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNTextInputNotificationAction/textInputPlaceholder
func (u UNTextInputNotificationAction) TextInputPlaceholder() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("textInputPlaceholder"))
	return foundation.NSStringFromID(rv).String()
}
