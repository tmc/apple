// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [UNUserNotificationCenter] class.
var (
	_UNUserNotificationCenterClass     UNUserNotificationCenterClass
	_UNUserNotificationCenterClassOnce sync.Once
)

func getUNUserNotificationCenterClass() UNUserNotificationCenterClass {
	_UNUserNotificationCenterClassOnce.Do(func() {
		_UNUserNotificationCenterClass = UNUserNotificationCenterClass{class: objc.GetClass("UNUserNotificationCenter")}
	})
	return _UNUserNotificationCenterClass
}

// GetUNUserNotificationCenterClass returns the class object for UNUserNotificationCenter.
func GetUNUserNotificationCenterClass() UNUserNotificationCenterClass {
	return getUNUserNotificationCenterClass()
}

type UNUserNotificationCenterClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (uc UNUserNotificationCenterClass) Class() objc.Class {
	return uc.class
}

// Alloc allocates memory for a new instance of the class.
func (uc UNUserNotificationCenterClass) Alloc() UNUserNotificationCenter {
	rv := objc.Send[UNUserNotificationCenter](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// The central object for managing notification-related activities for your
// app or app extension.
//
// # Overview
// 
// Use the shared [UNUserNotificationCenter] object to manage all
// notification-related behaviors in your app or app extension. Specifically,
// use this object to do the following:
// 
// - Request authorization to interact with the user through alerts, sounds,
// and icon badges. See [Asking permission to use notifications]. - Declare
// the notification types that your app supports and the custom actions the
// user may perform when the system delivers those notifications. See
// [Declaring your actionable notification types]. - Schedule the delivery of
// notifications from your app. See [Scheduling a notification locally from
// your app]. - Process the payloads from remote notifications the system
// delivers by Apple Push Notification service (APNs). See [Handling
// notifications and notification-related actions]. - Manage the already
// delivered notifications the system displays in Notification Center. See
// Managing Delivered Notifications. - Handle user-selected actions associated
// with your custom notification types. See [Handling notifications and
// notification-related actions]. - Get the notification-related settings for
// your app. See Managing Settings and Authorization.
// 
// To handle incoming notifications and notification-related actions, create
// an object that adopts the [UNUserNotificationCenterDelegate] protocol and
// assign it to the [UNUserNotificationCenter.Delegate] property. Always assign an object to the
// [UNUserNotificationCenter.Delegate] property before performing any tasks that might interact with
// that delegate.
// 
// You may use the shared user notification center object simultaneously from
// any of your app’s threads. The object processes requests serially in the
// order that the system initiates them.
//
// [Asking permission to use notifications]: https://developer.apple.com/documentation/UserNotifications/asking-permission-to-use-notifications
// [Declaring your actionable notification types]: https://developer.apple.com/documentation/UserNotifications/declaring-your-actionable-notification-types
// [Handling notifications and notification-related actions]: https://developer.apple.com/documentation/UserNotifications/handling-notifications-and-notification-related-actions
// [Scheduling a notification locally from your app]: https://developer.apple.com/documentation/UserNotifications/scheduling-a-notification-locally-from-your-app
//
// # Managing the notification center
//
//   - [UNUserNotificationCenter.GetNotificationSettingsWithCompletionHandler]: Retrieves the authorization and feature-related settings for your app.
//   - [UNUserNotificationCenter.SetBadgeCountWithCompletionHandler]: Updates the badge count for your app’s icon.
//
// # Requesting authorization
//
//   - [UNUserNotificationCenter.RequestAuthorizationWithOptionsCompletionHandler]: Requests a person’s authorization to allow local and remote notifications for your app.
//
// # Processing received notifications
//
//   - [UNUserNotificationCenter.Delegate]: The notification center’s delegate.
//   - [UNUserNotificationCenter.SetDelegate]
//   - [UNUserNotificationCenter.SupportsContentExtensions]: A Boolean value that indicates whether the device supports notification content extensions.
//
// # Scheduling notifications
//
//   - [UNUserNotificationCenter.AddNotificationRequestWithCompletionHandler]: Schedules the delivery of a local notification.
//   - [UNUserNotificationCenter.RemovePendingNotificationRequestsWithIdentifiers]: Removes your app’s local notifications that are pending and match the specified identifiers.
//   - [UNUserNotificationCenter.RemoveAllPendingNotificationRequests]: Removes all of your app’s pending local notifications.
//
// # Removing delivered notifications
//
//   - [UNUserNotificationCenter.RemoveDeliveredNotificationsWithIdentifiers]: Removes your app’s notifications from Notification Center that match the specified identifiers.
//   - [UNUserNotificationCenter.RemoveAllDeliveredNotifications]: Removes all of your app’s delivered notifications from Notification Center.
//
// # Managing notification categories
//
//   - [UNUserNotificationCenter.SetNotificationCategories]: Registers the notification categories that your app supports.
//
// # Handling errors
//
//   - [UNUserNotificationCenter.UNErrorDomain]: The error domain for notifications.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenter
type UNUserNotificationCenter struct {
	objectivec.Object
}

// UNUserNotificationCenterFromID constructs a [UNUserNotificationCenter] from an objc.ID.
//
// The central object for managing notification-related activities for your
// app or app extension.
func UNUserNotificationCenterFromID(id objc.ID) UNUserNotificationCenter {
	return UNUserNotificationCenter{objectivec.Object{ID: id}}
}
// NOTE: UNUserNotificationCenter adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [UNUserNotificationCenter] class.
//
// # Managing the notification center
//
//   - [IUNUserNotificationCenter.GetNotificationSettingsWithCompletionHandler]: Retrieves the authorization and feature-related settings for your app.
//   - [IUNUserNotificationCenter.SetBadgeCountWithCompletionHandler]: Updates the badge count for your app’s icon.
//
// # Requesting authorization
//
//   - [IUNUserNotificationCenter.RequestAuthorizationWithOptionsCompletionHandler]: Requests a person’s authorization to allow local and remote notifications for your app.
//
// # Processing received notifications
//
//   - [IUNUserNotificationCenter.Delegate]: The notification center’s delegate.
//   - [IUNUserNotificationCenter.SetDelegate]
//   - [IUNUserNotificationCenter.SupportsContentExtensions]: A Boolean value that indicates whether the device supports notification content extensions.
//
// # Scheduling notifications
//
//   - [IUNUserNotificationCenter.AddNotificationRequestWithCompletionHandler]: Schedules the delivery of a local notification.
//   - [IUNUserNotificationCenter.RemovePendingNotificationRequestsWithIdentifiers]: Removes your app’s local notifications that are pending and match the specified identifiers.
//   - [IUNUserNotificationCenter.RemoveAllPendingNotificationRequests]: Removes all of your app’s pending local notifications.
//
// # Removing delivered notifications
//
//   - [IUNUserNotificationCenter.RemoveDeliveredNotificationsWithIdentifiers]: Removes your app’s notifications from Notification Center that match the specified identifiers.
//   - [IUNUserNotificationCenter.RemoveAllDeliveredNotifications]: Removes all of your app’s delivered notifications from Notification Center.
//
// # Managing notification categories
//
//   - [IUNUserNotificationCenter.SetNotificationCategories]: Registers the notification categories that your app supports.
//
// # Handling errors
//
//   - [IUNUserNotificationCenter.UNErrorDomain]: The error domain for notifications.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenter
type IUNUserNotificationCenter interface {
	objectivec.IObject

	// Topic: Managing the notification center

	// Retrieves the authorization and feature-related settings for your app.
	GetNotificationSettingsWithCompletionHandler(completionHandler UNNotificationSettingsHandler)
	// Updates the badge count for your app’s icon.
	SetBadgeCountWithCompletionHandler(newBadgeCount int, completionHandler ErrorHandler)

	// Topic: Requesting authorization

	// Requests a person’s authorization to allow local and remote notifications for your app.
	RequestAuthorizationWithOptionsCompletionHandler(options UNAuthorizationOptions, completionHandler BoolErrorHandler)

	// Topic: Processing received notifications

	// The notification center’s delegate.
	Delegate() UNUserNotificationCenterDelegate
	SetDelegate(value UNUserNotificationCenterDelegate)
	// A Boolean value that indicates whether the device supports notification content extensions.
	SupportsContentExtensions() bool

	// Topic: Scheduling notifications

	// Schedules the delivery of a local notification.
	AddNotificationRequestWithCompletionHandler(request IUNNotificationRequest, completionHandler ErrorHandler)
	// Removes your app’s local notifications that are pending and match the specified identifiers.
	RemovePendingNotificationRequestsWithIdentifiers(identifiers []string)
	// Removes all of your app’s pending local notifications.
	RemoveAllPendingNotificationRequests()

	// Topic: Removing delivered notifications

	// Removes your app’s notifications from Notification Center that match the specified identifiers.
	RemoveDeliveredNotificationsWithIdentifiers(identifiers []string)
	// Removes all of your app’s delivered notifications from Notification Center.
	RemoveAllDeliveredNotifications()

	// Topic: Managing notification categories

	// Registers the notification categories that your app supports.
	SetNotificationCategories(categories foundation.INSSet)

	// Topic: Handling errors

	// The error domain for notifications.
	UNErrorDomain() string
}

// Init initializes the instance.
func (u UNUserNotificationCenter) Init() UNUserNotificationCenter {
	rv := objc.Send[UNUserNotificationCenter](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u UNUserNotificationCenter) Autorelease() UNUserNotificationCenter {
	rv := objc.Send[UNUserNotificationCenter](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewUNUserNotificationCenter creates a new UNUserNotificationCenter instance.
func NewUNUserNotificationCenter() UNUserNotificationCenter {
	class := getUNUserNotificationCenterClass()
	rv := objc.Send[UNUserNotificationCenter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Retrieves the authorization and feature-related settings for your app.
//
// completionHandler: The block to execute asynchronously with the results. Your app may execute
// this block on a background thread. The block has no return value and takes
// the following parameter:
// 
// settings: The [UNNotificationSettings] object containing the current
// authorization settings for your app.
//
// # Discussion
// 
// Use this method to determine the user interactions and notification-related
// features that the system authorizes your app to use. You might then use
// this information to enable or disable specific notification-related
// features of your app.
// 
// When the user initially grants authorization to your app, the system gives
// your app a set of default notification-related settings. The user may
// change those settings at any time to enable or disable specific
// capabilities. For example, the user might disable the playing of sounds
// when a notification arrives.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenter/getNotificationSettings(completionHandler:)
func (u UNUserNotificationCenter) GetNotificationSettingsWithCompletionHandler(completionHandler UNNotificationSettingsHandler) {
_block0, _ := NewUNNotificationSettingsBlock(completionHandler)
	objc.Send[objc.ID](u.ID, objc.Sel("getNotificationSettingsWithCompletionHandler:"), _block0)
}
// Updates the badge count for your app’s icon.
//
// newBadgeCount: The new value to display.
//
// completionHandler: The handler to execute after the update finishes. If the update fails, the
// system provides an error that contains additional information about the
// failure.
//
// # Discussion
// 
// Here’s an example that sets the badge count to a specific number.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenter/setBadgeCount(_:withCompletionHandler:)
func (u UNUserNotificationCenter) SetBadgeCountWithCompletionHandler(newBadgeCount int, completionHandler ErrorHandler) {
_block1, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](u.ID, objc.Sel("setBadgeCount:withCompletionHandler:"), newBadgeCount, _block1)
}
// Requests a person’s authorization to allow local and remote notifications
// for your app.
//
// options: The authorization options your app is requesting. You may combine the
// available constants to request authorization for multiple items. Request
// only the authorization options that you plan to use. For a list of possible
// values, see [UNAuthorizationOptions].
// //
// [UNAuthorizationOptions]: https://developer.apple.com/documentation/UserNotifications/UNAuthorizationOptions
//
// completionHandler: The block to execute asynchronously with the results. This block may
// execute on a background thread. The block has no return value and has the
// following parameters:
// 
// granted: A Boolean value indicating whether the person grants
// authorization. The value of this parameter is [true] when the person grants
// authorization for one or more options. The value is [false] when the person
// denies authorization or authorization is undetermined. Use
// [GetNotificationSettingsWithCompletionHandler] to check the authorization
// status. error: An object containing error information or `nil` if no error
// occurs.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// If your app’s local or remote notifications involve user interactions,
// you must request authorization for the system to perform those interactions
// on your app’s behalf. Interactions include displaying an alert, playing a
// sound, or badging the app’s icon.
// 
// The first time your app calls the method, the system prompts the person to
// authorize the requested interactions. The person may grant or deny
// authorization, and the system stores the person’s response. Subsequent
// calls to this method don’t prompt the person again. After determining the
// authorization status, the user notification center object executes the
// block in the `completionHandler` parameter. Use that block to make any
// adjustments to your app’s behavior. For example, if the person denied
// authorization, you might notify a remote notification server not to send
// notifications to the user’s device.
// 
// The person may change the interactions they allow at any time in system
// settings. Use the [GetNotificationSettingsWithCompletionHandler] method to
// determine what interactions are allowed for your app.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenter/requestAuthorization(options:completionHandler:)
func (u UNUserNotificationCenter) RequestAuthorizationWithOptionsCompletionHandler(options UNAuthorizationOptions, completionHandler BoolErrorHandler) {
_block1, _ := NewBoolErrorBlock(completionHandler)
	objc.Send[objc.ID](u.ID, objc.Sel("requestAuthorizationWithOptions:completionHandler:"), options, _block1)
}
// Schedules the delivery of a local notification.
//
// request: The request object containing the notification payload and trigger
// information. This parameter must not be `nil`.
//
// completionHandler: The block to execute with the results. This block may be executed on a
// background thread. The block has no return value and takes the following
// parameter:
// 
// error: An error object indicating whether a problem occurred. If the
// notification was scheduled successfully, this parameter is `nil`;
// otherwise, it is set to an error object indicating the reason for the
// failure.
//
// # Discussion
// 
// This method schedules local notifications only; you cannot use it to
// schedule the delivery of remote notifications. Upon calling this method,
// the system begins tracking the trigger conditions associated with your
// request. When the trigger condition is met, the system delivers your
// notification. If the request does not contain a [UNNotificationTrigger]
// object, the notification is delivered right away.
// 
// You may call this method from any thread of your app.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenter/add(_:withCompletionHandler:)
func (u UNUserNotificationCenter) AddNotificationRequestWithCompletionHandler(request IUNNotificationRequest, completionHandler ErrorHandler) {
_block1, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](u.ID, objc.Sel("addNotificationRequest:withCompletionHandler:"), request, _block1)
}
// Removes your app’s local notifications that are pending and match the
// specified identifiers.
//
// identifiers: An array of [NSString] objects, each of which contains the [Identifier] of
// an active [UNNotificationRequest] object. If the identifier belongs to a
// non repeating request, and the trigger condition for that request has
// already been met, this method ignores the identifier.
// //
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
//
// # Discussion
// 
// This method executes asynchronously, removing the pending notification
// requests on a secondary thread.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenter/removePendingNotificationRequests(withIdentifiers:)
func (u UNUserNotificationCenter) RemovePendingNotificationRequestsWithIdentifiers(identifiers []string) {
	objc.Send[objc.ID](u.ID, objc.Sel("removePendingNotificationRequestsWithIdentifiers:"), objectivec.StringSliceToNSArray(identifiers))
}
// Removes all of your app’s pending local notifications.
//
// # Discussion
// 
// This method executes asynchronously, removing all pending notification
// requests on a secondary thread.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenter/removeAllPendingNotificationRequests()
func (u UNUserNotificationCenter) RemoveAllPendingNotificationRequests() {
	objc.Send[objc.ID](u.ID, objc.Sel("removeAllPendingNotificationRequests"))
}
// Removes your app’s notifications from Notification Center that match the
// specified identifiers.
//
// identifiers: An array of [NSString] objects, each of which corresponds to a value in the
// [Identifier] property of a [UNNotificationRequest] object. This method
// ignores the identifiers of requests whose notifications are not currently
// displayed in Notification Center.
// //
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
//
// # Discussion
// 
// Use this method to selectively remove notifications that you no longer want
// displayed in Notification Center. The method executes asynchronously,
// returning immediately and removing the specified notifications on a
// background thread.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenter/removeDeliveredNotifications(withIdentifiers:)
func (u UNUserNotificationCenter) RemoveDeliveredNotificationsWithIdentifiers(identifiers []string) {
	objc.Send[objc.ID](u.ID, objc.Sel("removeDeliveredNotificationsWithIdentifiers:"), objectivec.StringSliceToNSArray(identifiers))
}
// Removes all of your app’s delivered notifications from Notification
// Center.
//
// # Discussion
// 
// Use this method to remove all of your app’s delivered notifications from
// Notification Center. The method executes asynchronously, returning
// immediately and removing the identifiers on a background thread. This
// method does not affect any notification requests that are scheduled, but
// have not yet been delivered.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenter/removeAllDeliveredNotifications()
func (u UNUserNotificationCenter) RemoveAllDeliveredNotifications() {
	objc.Send[objc.ID](u.ID, objc.Sel("removeAllDeliveredNotifications"))
}
// Registers the notification categories that your app supports.
//
// categories: A set of [UNNotificationCategory] objects, each of which contains the
// actions that are displayed with the notification interface. This parameter
// must contain all of your app’s supported categories.
//
// # Discussion
// 
// Call this method at launch time to register your app’s actionable
// notification types. This method registers all of your categories at once,
// replacing any previously registered categories with the new ones in the
// `categories` parameter. Typically, you call this method only once.
// 
// Each object in the `categories` parameter contains a string for identifying
// the notification’s type. It also contains one or more custom actions that
// the user may perform in response to notifications of that type. When the
// system displays an alert for a notification, it looks in the notification
// payload for one of the identifier strings from your category objects. If it
// finds one, it adds user-selectable buttons for each action associated with
// that category object. Tapping a button notifies your app of the selected
// action, without bringing your app to the foreground.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenter/setNotificationCategories(_:)
func (u UNUserNotificationCenter) SetNotificationCategories(categories foundation.INSSet) {
	objc.Send[objc.ID](u.ID, objc.Sel("setNotificationCategories:"), categories)
}

// Returns your app’s notification center.
//
// # Return Value
// 
// The notification center object to use.
//
// # Discussion
// 
// Always use this method to retrieve the shared notification center object
// for your app. Do not try to create instances of the
// [UNUserNotificationCenter] class directly.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenter/current()
func (_UNUserNotificationCenterClass UNUserNotificationCenterClass) CurrentNotificationCenter() UNUserNotificationCenter {
	rv := objc.Send[objc.ID](objc.ID(_UNUserNotificationCenterClass.class), objc.Sel("currentNotificationCenter"))
	return UNUserNotificationCenterFromID(rv)
}

// The notification center’s delegate.
//
// # Discussion
// 
// Use the delegate object to respond to user-selected actions and to process
// incoming notifications when your app is in the foreground. For example, you
// might use your delegate to silence notifications when your app is in the
// foreground.
// 
// To guarantee that your app responds to all actionable notifications, you
// must set the value of this property before your app finishes launching. For
// an iOS app, this means updating this property in the
// [application(_:willFinishLaunchingWithOptions:)] or
// [application(_:didFinishLaunchingWithOptions:)] method of the app delegate.
// Notifications that cause your app to be launched or delivered shortly after
// these methods finish executing.
// 
// For more information about implementing the delegate methods, see
// [UNUserNotificationCenterDelegate].
//
// [application(_:didFinishLaunchingWithOptions:)]: https://developer.apple.com/documentation/UIKit/UIApplicationDelegate/application(_:didFinishLaunchingWithOptions:)
// [application(_:willFinishLaunchingWithOptions:)]: https://developer.apple.com/documentation/UIKit/UIApplicationDelegate/application(_:willFinishLaunchingWithOptions:)
//
// See: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenter/delegate
func (u UNUserNotificationCenter) Delegate() UNUserNotificationCenterDelegate {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("delegate"))
	return UNUserNotificationCenterDelegateObjectFromID(rv)
}
func (u UNUserNotificationCenter) SetDelegate(value UNUserNotificationCenterDelegate) {
	objc.Send[struct{}](u.ID, objc.Sel("setDelegate:"), value)
}
// A Boolean value that indicates whether the device supports notification
// content extensions.
//
// # Discussion
// 
// Notification content extensions let you customize the appearance of the
// alerts displayed for your app’s notifications. The value of this property
// is [true] for devices that support notification content extensions and
// [false] for devices that do not support them. For information about how to
// implement a notification content extension, see [Customizing the Appearance
// of Notifications].
//
// [Customizing the Appearance of Notifications]: https://developer.apple.com/documentation/UserNotificationsUI/customizing-the-appearance-of-notifications
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenter/supportsContentExtensions
func (u UNUserNotificationCenter) SupportsContentExtensions() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("supportsContentExtensions"))
	return rv
}
// The error domain for notifications.
//
// See: https://developer.apple.com/documentation/usernotifications/unerrordomain
func (u UNUserNotificationCenter) UNErrorDomain() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("UNErrorDomain"))
	return foundation.NSStringFromID(rv).String()
}

// GetNotificationSettings is a synchronous wrapper around [UNUserNotificationCenter.GetNotificationSettingsWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (u UNUserNotificationCenter) GetNotificationSettings(ctx context.Context) (*UNNotificationSettings, error) {
	done := make(chan *UNNotificationSettings, 1)
	u.GetNotificationSettingsWithCompletionHandler(func(val *UNNotificationSettings) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// SetBadgeCount is a synchronous wrapper around [UNUserNotificationCenter.SetBadgeCountWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (u UNUserNotificationCenter) SetBadgeCount(ctx context.Context, newBadgeCount int) error {
	done := make(chan error, 1)
	u.SetBadgeCountWithCompletionHandler(newBadgeCount, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// RequestAuthorizationWithOptions is a synchronous wrapper around [UNUserNotificationCenter.RequestAuthorizationWithOptionsCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (u UNUserNotificationCenter) RequestAuthorizationWithOptions(ctx context.Context, options UNAuthorizationOptions) (bool, error) {
	type result struct {
		val bool
		err error
	}
	done := make(chan result, 1)
	u.RequestAuthorizationWithOptionsCompletionHandler(options, func(val bool, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return false, ctx.Err()
	}
}

// AddNotificationRequest is a synchronous wrapper around [UNUserNotificationCenter.AddNotificationRequestWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (u UNUserNotificationCenter) AddNotificationRequest(ctx context.Context, request IUNNotificationRequest) error {
	done := make(chan error, 1)
	u.AddNotificationRequestWithCompletionHandler(request, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

