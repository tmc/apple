// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [UNNotificationAction] class.
var (
	_UNNotificationActionClass     UNNotificationActionClass
	_UNNotificationActionClassOnce sync.Once
)

func getUNNotificationActionClass() UNNotificationActionClass {
	_UNNotificationActionClassOnce.Do(func() {
		_UNNotificationActionClass = UNNotificationActionClass{class: objc.GetClass("UNNotificationAction")}
	})
	return _UNNotificationActionClass
}

// GetUNNotificationActionClass returns the class object for UNNotificationAction.
func GetUNNotificationActionClass() UNNotificationActionClass {
	return getUNNotificationActionClass()
}

type UNNotificationActionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (uc UNNotificationActionClass) Class() objc.Class {
	return uc.class
}

// Alloc allocates memory for a new instance of the class.
func (uc UNNotificationActionClass) Alloc() UNNotificationAction {
	rv := objc.Send[UNNotificationAction](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// A task your app performs in response to a notification that the system
// delivers.
//
// # Overview
//
// Use [UNNotificationAction] objects to define the actions that your app can
// perform in response to a delivered notification. You define the actions
// that your app supports. For example, a meeting app might define actions for
// accepting or rejecting a meeting invitation. The action object itself
// contains the title to display in an action button and the button’s
// appearance. After creating action objects, add them to a
// [UNNotificationCategory] object and register your categories with the
// system.
//
// For information on how to define actions and categories, see [Declaring
// your actionable notification types].
//
// # Responding to the Selection of Actions
//
// When the user selects one of your actions in response to a notification,
// the system notifies the delegate of the shared [UNUserNotificationCenter]
// object. Specifically, the system calls the
// [UserNotificationCenterDidReceiveNotificationResponseWithCompletionHandler]
// method of your delegate object. The response object passed to your delegate
// includes the [UNNotificationAction.Identifier] string of the action the user selects, which you
// can use to perform the corresponding task.
//
// For information on how to handle actions, see [Handling notifications and
// notification-related actions].
//
// # Getting Information
//
//   - [UNNotificationAction.Identifier]: The unique string that your app uses to identify the action.
//   - [UNNotificationAction.Title]: The localized string to use as the title of the action.
//   - [UNNotificationAction.Icon]: The icon associated with the action.
//
// # Getting Options
//
//   - [UNNotificationAction.Options]: The behaviors associated with the action.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationAction
//
// [Declaring your actionable notification types]: https://developer.apple.com/documentation/UserNotifications/declaring-your-actionable-notification-types
// [Handling notifications and notification-related actions]: https://developer.apple.com/documentation/UserNotifications/handling-notifications-and-notification-related-actions
type UNNotificationAction struct {
	objectivec.Object
}

// UNNotificationActionFromID constructs a [UNNotificationAction] from an objc.ID.
//
// A task your app performs in response to a notification that the system
// delivers.
func UNNotificationActionFromID(id objc.ID) UNNotificationAction {
	return UNNotificationAction{objectivec.Object{ID: id}}
}

// NOTE: UNNotificationAction adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [UNNotificationAction] class.
//
// # Getting Information
//
//   - [IUNNotificationAction.Identifier]: The unique string that your app uses to identify the action.
//   - [IUNNotificationAction.Title]: The localized string to use as the title of the action.
//   - [IUNNotificationAction.Icon]: The icon associated with the action.
//
// # Getting Options
//
//   - [IUNNotificationAction.Options]: The behaviors associated with the action.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationAction
type IUNNotificationAction interface {
	objectivec.IObject

	// Topic: Getting Information

	// The unique string that your app uses to identify the action.
	Identifier() string
	// The localized string to use as the title of the action.
	Title() string
	// The icon associated with the action.
	Icon() IUNNotificationActionIcon

	// Topic: Getting Options

	// The behaviors associated with the action.
	Options() UNNotificationActionOptions

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (u UNNotificationAction) Init() UNNotificationAction {
	rv := objc.Send[UNNotificationAction](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u UNNotificationAction) Autorelease() UNNotificationAction {
	rv := objc.Send[UNNotificationAction](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewUNNotificationAction creates a new UNNotificationAction instance.
func NewUNNotificationAction() UNNotificationAction {
	class := getUNNotificationActionClass()
	rv := objc.Send[UNNotificationAction](objc.ID(class.class), objc.Sel("new"))
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
func NewUNNotificationActionWithIdentifierTitleOptions(identifier string, title string, options UNNotificationActionOptions) UNNotificationAction {
	rv := objc.Send[objc.ID](objc.ID(getUNNotificationActionClass().class), objc.Sel("actionWithIdentifier:title:options:"), objc.String(identifier), objc.String(title), options)
	return UNNotificationActionFromID(rv)
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
func NewUNNotificationActionWithIdentifierTitleOptionsIcon(identifier string, title string, options UNNotificationActionOptions, icon IUNNotificationActionIcon) UNNotificationAction {
	rv := objc.Send[objc.ID](objc.ID(getUNNotificationActionClass().class), objc.Sel("actionWithIdentifier:title:options:icon:"), objc.String(identifier), objc.String(title), options, icon)
	return UNNotificationActionFromID(rv)
}

func (u UNNotificationAction) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](u.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The unique string that your app uses to identify the action.
//
// # Discussion
//
// When the user selects an action, the system reports the value of this
// string to your app. Because your app handles all actions by using a single
// delegate method, the identifier strings for all of your app’s actions
// must be unique.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationAction/identifier
func (u UNNotificationAction) Identifier() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("identifier"))
	return foundation.NSStringFromID(rv).String()
}

// The localized string to use as the title of the action.
//
// # Discussion
//
// The system displays this string as the title of the button that the user
// taps or selects in the notification interface.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationAction/title
func (u UNNotificationAction) Title() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("title"))
	return foundation.NSStringFromID(rv).String()
}

// The icon associated with the action.
//
// # Discussion
//
// The system displays this icon in the notification interface to help the
// user identify the app associated with the action.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationAction/icon
func (u UNNotificationAction) Icon() IUNNotificationActionIcon {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("icon"))
	return UNNotificationActionIconFromID(objc.ID(rv))
}

// The behaviors associated with the action.
//
// # Discussion
//
// You app should define options for an action when your app requires the
// corresponding behavior.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationAction/options
func (u UNNotificationAction) Options() UNNotificationActionOptions {
	rv := objc.Send[UNNotificationActionOptions](u.ID, objc.Sel("options"))
	return UNNotificationActionOptions(rv)
}
