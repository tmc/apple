// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [UNNotificationCategory] class.
var (
	_UNNotificationCategoryClass     UNNotificationCategoryClass
	_UNNotificationCategoryClassOnce sync.Once
)

func getUNNotificationCategoryClass() UNNotificationCategoryClass {
	_UNNotificationCategoryClassOnce.Do(func() {
		_UNNotificationCategoryClass = UNNotificationCategoryClass{class: objc.GetClass("UNNotificationCategory")}
	})
	return _UNNotificationCategoryClass
}

// GetUNNotificationCategoryClass returns the class object for UNNotificationCategory.
func GetUNNotificationCategoryClass() UNNotificationCategoryClass {
	return getUNNotificationCategoryClass()
}

type UNNotificationCategoryClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (uc UNNotificationCategoryClass) Class() objc.Class {
	return uc.class
}

// Alloc allocates memory for a new instance of the class.
func (uc UNNotificationCategoryClass) Alloc() UNNotificationCategory {
	rv := objc.Send[UNNotificationCategory](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// A type of notification your app supports and the custom actions that the
// system displays.
//
// # Overview
//
// A [UNNotificationCategory] object defines a type of notification that your
// executable can receive. You create category objects to define your app’s
// — notifications that have action buttons the user can select in response
// to the notification. Each category object you create stores the actions and
// other behaviors associated with a specific type of notification. Register
// your category objects using the [SetNotificationCategories] method of
// [UNUserNotificationCenter]. You can register as many category objects as
// you need.
//
// To apply category objects to your notifications, include the category’s
// identifier string in the payload of any notifications you create. For local
// notifications, put this string in the [UNNotificationCategory.CategoryIdentifier] property of the
// [UNMutableNotificationContent] object that you use to specify the
// notification’s content. For remote notifications, use this string as the
// value of the `category` key in the `aps` dictionary of your payload.
//
// Categories can have associated actions, which define custom buttons the
// system displays for notifications of that category. When the system has
// unlimited space, the system displays up to 10 actions. When the system has
// limited space, the system displays at most two actions.
//
// # Getting the Information
//
//   - [UNNotificationCategory.Identifier]: The unique string assigned to the category.
//   - [UNNotificationCategory.Actions]: The actions to display when the system delivers notifications of this type.
//   - [UNNotificationCategory.IntentIdentifiers]: The intents related to notifications of this category.
//   - [UNNotificationCategory.HiddenPreviewsBodyPlaceholder]: The placeholder text to display when the system disables notification previews for the app.
//   - [UNNotificationCategory.CategorySummaryFormat]: A format string for the summary description used when the system groups the category’s notifications.
//
// # Getting the Options
//
//   - [UNNotificationCategory.Options]: Options for how to handle notifications of this type.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationCategory
type UNNotificationCategory struct {
	objectivec.Object
}

// UNNotificationCategoryFromID constructs a [UNNotificationCategory] from an objc.ID.
//
// A type of notification your app supports and the custom actions that the
// system displays.
func UNNotificationCategoryFromID(id objc.ID) UNNotificationCategory {
	return UNNotificationCategory{objectivec.Object{ID: id}}
}

// NOTE: UNNotificationCategory adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [UNNotificationCategory] class.
//
// # Getting the Information
//
//   - [IUNNotificationCategory.Identifier]: The unique string assigned to the category.
//   - [IUNNotificationCategory.Actions]: The actions to display when the system delivers notifications of this type.
//   - [IUNNotificationCategory.IntentIdentifiers]: The intents related to notifications of this category.
//   - [IUNNotificationCategory.HiddenPreviewsBodyPlaceholder]: The placeholder text to display when the system disables notification previews for the app.
//   - [IUNNotificationCategory.CategorySummaryFormat]: A format string for the summary description used when the system groups the category’s notifications.
//
// # Getting the Options
//
//   - [IUNNotificationCategory.Options]: Options for how to handle notifications of this type.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationCategory
type IUNNotificationCategory interface {
	objectivec.IObject

	// Topic: Getting the Information

	// The unique string assigned to the category.
	Identifier() string
	// The actions to display when the system delivers notifications of this type.
	Actions() []UNNotificationAction
	// The intents related to notifications of this category.
	IntentIdentifiers() []string
	// The placeholder text to display when the system disables notification previews for the app.
	HiddenPreviewsBodyPlaceholder() string
	// A format string for the summary description used when the system groups the category’s notifications.
	CategorySummaryFormat() string

	// Topic: Getting the Options

	// Options for how to handle notifications of this type.
	Options() UNNotificationCategoryOptions

	// The identifier of the notification’s category.
	CategoryIdentifier() string
	SetCategoryIdentifier(value string)
	// The action performs a destructive task.
	Destructive() UNNotificationActionOptions
	SetDestructive(value UNNotificationActionOptions)
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (u UNNotificationCategory) Init() UNNotificationCategory {
	rv := objc.Send[UNNotificationCategory](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u UNNotificationCategory) Autorelease() UNNotificationCategory {
	rv := objc.Send[UNNotificationCategory](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewUNNotificationCategory creates a new UNNotificationCategory instance.
func NewUNNotificationCategory() UNNotificationCategory {
	class := getUNNotificationCategoryClass()
	rv := objc.Send[UNNotificationCategory](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a category object containing the specified actions, options,
// placeholder text used when previews aren’t shown, and summary format
// string.
//
// identifier: The unique identifier for the category. Each category that your app uses
// must have a unique identifier. Don’t specify an empty string.
//
// actions: The actions to display when the system delivers notifications of this type.
// When minimal space is available, the system displays only the first two
// actions in the array. You may specify an empty array for this parameter if
// you don’t want to display custom actions.
//
// intentIdentifiers: The intent identifier strings that you want to associate with notifications
// of this type. The Intents framework defines constants for each type of
// intent that you can associate with your notifications.
//
// hiddenPreviewsBodyPlaceholder: A placeholder string to display when the user has disabled notification
// previews for the app. Include the characters `%u` (the only supported
// formatting characters) in the string to represent the number of
// notifications with the same thread identifier. For example, the string
// “`%u Messages`” becomes “`2 Messages`” when there are two messages.
//
// To specify different strings for the singular and plural cases, use the
// [localizedUserNotificationString(forKey:arguments:)] method of [NSString]
// to specify the value for this parameter. The key passed to that method
// contains the identifier of an entry in a `XCUIElementTypeStringsdict`
// property list of your project. A strings dictionary lets you specify
// different formatted strings based on the language rules, and is as
// described in [Internationalization and Localization Guide].
//
// categorySummaryFormat: A format string for the summary description used when the system groups the
// category’s notifications.
//
// options: Additional options for handling notifications of this type. For a list of
// possible values, see [UNNotificationCategoryOptions].
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationCategory/init(identifier:actions:intentIdentifiers:hiddenPreviewsBodyPlaceholder:categorySummaryFormat:options:)
//
// [Internationalization and Localization Guide]: https://developer.apple.com/library/archive/documentation/MacOSX/Conceptual/BPInternational/Introduction/Introduction.html#//apple_ref/doc/uid/10000171i
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
// [localizedUserNotificationString(forKey:arguments:)]: https://developer.apple.com/documentation/Foundation/NSString/localizedUserNotificationString(forKey:arguments:)
// [UNNotificationCategoryOptions]: https://developer.apple.com/documentation/UserNotifications/UNNotificationCategoryOptions
func NewUNNotificationCategoryWithIdentifierActionsIntentIdentifiersHiddenPreviewsBodyPlaceholderCategorySummaryFormatOptions(identifier string, actions []UNNotificationAction, intentIdentifiers []string, hiddenPreviewsBodyPlaceholder string, categorySummaryFormat string, options UNNotificationCategoryOptions) UNNotificationCategory {
	rv := objc.Send[objc.ID](objc.ID(getUNNotificationCategoryClass().class), objc.Sel("categoryWithIdentifier:actions:intentIdentifiers:hiddenPreviewsBodyPlaceholder:categorySummaryFormat:options:"), objc.String(identifier), objectivec.IObjectSliceToNSArray(actions), objectivec.StringSliceToNSArray(intentIdentifiers), objc.String(hiddenPreviewsBodyPlaceholder), objc.String(categorySummaryFormat), options)
	return UNNotificationCategoryFromID(rv)
}

// Creates a category object containing the specified actions, options, and
// placeholder text used when previews aren’t shown.
//
// identifier: The unique identifier for the category. Each category that your app uses
// must have a unique identifier. Don’t specify an empty string.
//
// actions: The actions to display when the system delivers notifications of this type.
// When minimal space is available, the system displays only the first two
// actions in the array. You may specify an empty array for this parameter if
// you don’t want to display custom actions.
//
// intentIdentifiers: The intent identifier strings that you want to associate with notifications
// of this type. The Intents framework defines constants for each type of
// intent that you can associate with your notifications.
//
// hiddenPreviewsBodyPlaceholder: A placeholder string to display when the user has disabled notification
// previews for the app. Include the characters `%u` (the only supported
// formatting characters) in the string to represent the number of
// notifications with the same thread identifier. For example, the string
// “`%u Messages`” becomes “`2 Messages`” when there are two messages.
//
// To specify different strings for the singular and plural cases, use the
// [localizedUserNotificationString(forKey:arguments:)] method of [NSString]
// to specify the value for this parameter. The key passed to that method
// contains the identifier of an entry in a `XCUIElementTypeStringsdict`
// property list of your project. A strings dictionary lets you specify
// different formatted strings based on the language rules, and is as
// described in [Internationalization and Localization Guide].
//
// options: Additional options for handling notifications of this type. For a list of
// possible values, see [UNNotificationCategoryOptions].
//
// # Return Value
//
// An initialized category object.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationCategory/init(identifier:actions:intentIdentifiers:hiddenPreviewsBodyPlaceholder:options:)
//
// [Internationalization and Localization Guide]: https://developer.apple.com/library/archive/documentation/MacOSX/Conceptual/BPInternational/Introduction/Introduction.html#//apple_ref/doc/uid/10000171i
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
// [localizedUserNotificationString(forKey:arguments:)]: https://developer.apple.com/documentation/Foundation/NSString/localizedUserNotificationString(forKey:arguments:)
// [UNNotificationCategoryOptions]: https://developer.apple.com/documentation/UserNotifications/UNNotificationCategoryOptions
func NewUNNotificationCategoryWithIdentifierActionsIntentIdentifiersHiddenPreviewsBodyPlaceholderOptions(identifier string, actions []UNNotificationAction, intentIdentifiers []string, hiddenPreviewsBodyPlaceholder string, options UNNotificationCategoryOptions) UNNotificationCategory {
	rv := objc.Send[objc.ID](objc.ID(getUNNotificationCategoryClass().class), objc.Sel("categoryWithIdentifier:actions:intentIdentifiers:hiddenPreviewsBodyPlaceholder:options:"), objc.String(identifier), objectivec.IObjectSliceToNSArray(actions), objectivec.StringSliceToNSArray(intentIdentifiers), objc.String(hiddenPreviewsBodyPlaceholder), options)
	return UNNotificationCategoryFromID(rv)
}

// Creates a category object containing the specified actions and options.
//
// identifier: The unique identifier for the category. Each category that your app uses
// must have a unique identifier. Don’t specify an empty string.
//
// actions: The actions to display when the system delivers notifications of this type.
// When minimal space is available, the system displays only the first two
// actions in the array. You may specify an empty array for this parameter if
// you don’t want to display custom actions.
//
// intentIdentifiers: The intent identifier strings that you want to associate with notifications
// of this type. The Intents framework defines constants for each type of
// intent that you can associate with your notifications.
//
// options: Additional options for handling notifications of this type. For a list of
// possible values, see [UNNotificationCategoryOptions].
//
// # Return Value
//
// An initialized category object.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationCategory/init(identifier:actions:intentIdentifiers:options:)
//
// [UNNotificationCategoryOptions]: https://developer.apple.com/documentation/UserNotifications/UNNotificationCategoryOptions
func NewUNNotificationCategoryWithIdentifierActionsIntentIdentifiersOptions(identifier string, actions []UNNotificationAction, intentIdentifiers []string, options UNNotificationCategoryOptions) UNNotificationCategory {
	rv := objc.Send[objc.ID](objc.ID(getUNNotificationCategoryClass().class), objc.Sel("categoryWithIdentifier:actions:intentIdentifiers:options:"), objc.String(identifier), objectivec.IObjectSliceToNSArray(actions), objectivec.StringSliceToNSArray(intentIdentifiers), options)
	return UNNotificationCategoryFromID(rv)
}

func (u UNNotificationCategory) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](u.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The unique string assigned to the category.
//
// # Discussion
//
// Use this string to differentiate the different types of notifications that
// your app can send. To assign a category to a local notification, assign
// this string to the [CategoryIdentifier] property of the content object. To
// assign a category to a remote notification, use the string as the value of
// the `category` key in the notification payload `aps` dictionary.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationCategory/identifier
func (u UNNotificationCategory) Identifier() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("identifier"))
	return foundation.NSStringFromID(rv).String()
}

// The actions to display when the system delivers notifications of this type.
//
// # Discussion
//
// When displaying a notification assigned to this category, the system adds a
// button to the notification interface for each action in this property. The
// system displays these buttons after the notification’s content but before
// the Dismiss button.
//
// When displaying banner notifications, the system displays only the first
// two actions.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationCategory/actions
func (u UNNotificationCategory) Actions() []UNNotificationAction {
	rv := objc.Send[[]objc.ID](u.ID, objc.Sel("actions"))
	return objc.ConvertSlice(rv, func(id objc.ID) UNNotificationAction {
		return UNNotificationActionFromID(id)
	})
}

// The intents related to notifications of this category.
//
// # Discussion
//
// When the system delivers a notification, the presence of an intent
// identifier lets the system know that the notification is potentially
// related to the handling of a request made through Siri.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationCategory/intentIdentifiers
func (u UNNotificationCategory) IntentIdentifiers() []string {
	rv := objc.Send[[]objc.ID](u.ID, objc.Sel("intentIdentifiers"))
	return objc.ConvertSliceToStrings(rv)
}

// The placeholder text to display when the system disables notification
// previews for the app.
//
// # Discussion
//
// The string in this property may contain the special characters `%u` as a
// placeholder for the number of messages with the same thread identifier. If
// your app declares this string in a `XCUIElementTypeStringsdict` property
// list, the system formats the preview message using the information in that
// file. For more information about specifying a `XCUIElementTypeStringsdict`
// property file, see [Internationalization and Localization Guide].
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationCategory/hiddenPreviewsBodyPlaceholder
//
// [Internationalization and Localization Guide]: https://developer.apple.com/library/archive/documentation/MacOSX/Conceptual/BPInternational/Introduction/Introduction.html#//apple_ref/doc/uid/10000171i
func (u UNNotificationCategory) HiddenPreviewsBodyPlaceholder() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("hiddenPreviewsBodyPlaceholder"))
	return foundation.NSStringFromID(rv).String()
}

// A format string for the summary description used when the system groups the
// category’s notifications.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationCategory/categorySummaryFormat
func (u UNNotificationCategory) CategorySummaryFormat() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("categorySummaryFormat"))
	return foundation.NSStringFromID(rv).String()
}

// Options for how to handle notifications of this type.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationCategory/options
func (u UNNotificationCategory) Options() UNNotificationCategoryOptions {
	rv := objc.Send[UNNotificationCategoryOptions](u.ID, objc.Sel("options"))
	return UNNotificationCategoryOptions(rv)
}

// The identifier of the notification’s category.
//
// See: https://developer.apple.com/documentation/usernotifications/unmutablenotificationcontent/categoryidentifier
func (u UNNotificationCategory) CategoryIdentifier() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("categoryIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (u UNNotificationCategory) SetCategoryIdentifier(value string) {
	objc.Send[struct{}](u.ID, objc.Sel("setCategoryIdentifier:"), objc.String(value))
}

// The action performs a destructive task.
//
// See: https://developer.apple.com/documentation/usernotifications/unnotificationactionoptions/destructive
func (u UNNotificationCategory) Destructive() UNNotificationActionOptions {
	rv := objc.Send[UNNotificationActionOptions](u.ID, objc.Sel("UNNotificationActionOptionDestructive"))
	return UNNotificationActionOptions(rv)
}
func (u UNNotificationCategory) SetDestructive(value UNNotificationActionOptions) {
	objc.Send[struct{}](u.ID, objc.Sel("setUNNotificationActionOptionDestructive:"), value)
}
