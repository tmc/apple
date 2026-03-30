// Code generated from Apple documentation. DO NOT EDIT.

package usernotifications

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// BoolErrorHandler handles The block to execute asynchronously with the results.
//   - granted: A Boolean value indicating whether the person grants authorization. The value of this parameter is [true](<doc://com.apple.documentation/documentation/Swift/true>) when the person grants authorization for one or more options. The value is [false](<doc://com.apple.documentation/documentation/Swift/false>) when the person denies authorization or authorization is  undetermined. Use [getNotificationSettings(completionHandler:)](<doc://com.apple.usernotifications/documentation/UserNotifications/UNUserNotificationCenter/getNotificationSettings(completionHandler:)>) to check the authorization status.
//   - error: An object containing error information or `nil` if no error occurs.
//
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [UNUserNotificationCenter.RequestAuthorizationWithOptionsCompletionHandler]
type BoolErrorHandler = func(bool, error)

// NewBoolErrorBlock wraps a Go [BoolErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [UNUserNotificationCenter.RequestAuthorizationWithOptionsCompletionHandler]
func NewBoolErrorBlock(handler BoolErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, primitiveVal bool, errID objc.ID) {
		handler(primitiveVal, foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// ErrorHandler handles The handler to execute after the update finishes.
//   - error: An error object indicating whether a problem occurred. If the notification was scheduled successfully, this parameter is `nil`; otherwise, it is set to an error object indicating the reason for the failure.
//
// The error can be type-asserted to *foundation.NSError for Domain, Code, and UserInfo.
//
// Used by:
//   - [UNUserNotificationCenter.AddNotificationRequestWithCompletionHandler]
//   - [UNUserNotificationCenter.SetBadgeCountWithCompletionHandler]
type ErrorHandler = func(error)

// NewErrorBlock wraps a Go [ErrorHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [UNUserNotificationCenter.AddNotificationRequestWithCompletionHandler]
//   - [UNUserNotificationCenter.SetBadgeCountWithCompletionHandler]
func NewErrorBlock(handler ErrorHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, errID objc.ID) {
		handler(foundation.SafeErrorFrom(errID))
	})
	return objc.ID(block), func() { block.Release() }
}

// UNNotificationArrayHandler handles The block to execute with the results.
//   - notifications: An array of [UNNotification](<doc://com.apple.usernotifications/documentation/UserNotifications/UNNotification>) objects representing the local and remote notifications of your app that have been delivered and are still visible in Notification Center. If none of your app’s notifications are visible in Notification Center, the array is empty.
//
// Used by:
//   - [UNUserNotificationCenter.GetDeliveredNotificationsWithCompletionHandler]
type UNNotificationArrayHandler = func(*[]UNNotification)

// UNNotificationCategorySetHandler handles The block to execute asynchronously with the results.
//   - categories: The set of [UNNotificationCategory](<doc://com.apple.usernotifications/documentation/UserNotifications/UNNotificationCategory>) objects containing your registered notification types. If your app has not yet registered any categories, this parameter is an empty set.
//
// Used by:
//   - [UNUserNotificationCenter.GetNotificationCategoriesWithCompletionHandler]
type UNNotificationCategorySetHandler = func(*foundation.INSSet)

// UNNotificationContentHandler handles The block to execute with the modified content.
//   - contentToDeliver: A [UNNotificationContent](<doc://com.apple.usernotifications/documentation/UserNotifications/UNNotificationContent>) object with the content the system displays to the user.
//
// Used by:
//   - [UNNotificationServiceExtension.DidReceiveNotificationRequestWithContentHandler]
type UNNotificationContentHandler = func(*UNNotificationContent)

// NewUNNotificationContentBlock wraps a Go [UNNotificationContentHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [UNNotificationServiceExtension.DidReceiveNotificationRequestWithContentHandler]
func NewUNNotificationContentBlock(handler UNNotificationContentHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *UNNotificationContent
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := UNNotificationContentFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// UNNotificationPresentationOptionsHandler handles The block to execute with the presentation option for the notification.
//   - options: The option for notifying the user. Specify [UNNotificationPresentationOptionNone](<doc://com.apple.usernotifications/documentation/UserNotifications/UNNotificationPresentationOptionNone>) to silence the notification completely. Specify other values to interact with the user. For a list of possible options, see [UNNotificationPresentationOptions](<doc://com.apple.usernotifications/documentation/UserNotifications/UNNotificationPresentationOptions>).
//
// Used by:
//   - [UNUserNotificationCenterDelegate.UserNotificationCenterWillPresentNotificationWithCompletionHandler]
type UNNotificationPresentationOptionsHandler = func(UNNotificationPresentationOptions)

// NewUNNotificationPresentationOptionsBlock wraps a Go [UNNotificationPresentationOptionsHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [UNUserNotificationCenterDelegate.UserNotificationCenterWillPresentNotificationWithCompletionHandler]
func NewUNNotificationPresentationOptionsBlock(handler UNNotificationPresentationOptionsHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, primitiveVal UNNotificationPresentationOptions) {
		handler(primitiveVal)
	})
	return objc.ID(block), func() { block.Release() }
}

// UNNotificationRequestArrayHandler handles A block for processing notification requests.
//   - requests: An array of [UNNotificationRequest](<doc://com.apple.usernotifications/documentation/UserNotifications/UNNotificationRequest>) objects representing the scheduled notification requests. If there are no scheduled requests, this array is empty.
//
// Used by:
//   - [UNUserNotificationCenter.GetPendingNotificationRequestsWithCompletionHandler]
type UNNotificationRequestArrayHandler = func(*[]UNNotificationRequest)

// UNNotificationSettingsHandler handles The block to execute asynchronously with the results.
//   - settings: The [UNNotificationSettings](<doc://com.apple.usernotifications/documentation/UserNotifications/UNNotificationSettings>) object containing the current authorization settings for your app.
//
// Used by:
//   - [UNUserNotificationCenter.GetNotificationSettingsWithCompletionHandler]
type UNNotificationSettingsHandler = func(*UNNotificationSettings)

// NewUNNotificationSettingsBlock wraps a Go [UNNotificationSettingsHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [UNUserNotificationCenter.GetNotificationSettingsWithCompletionHandler]
func NewUNNotificationSettingsBlock(handler UNNotificationSettingsHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID) {
		var result *UNNotificationSettings
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := UNNotificationSettingsFromID(resultID)
			result = &v
		}
		handler(result)
	})
	return objc.ID(block), func() { block.Release() }
}

// VoidHandler handles The block to execute when you have finished processing the user’s response.
//
// Used by:
//   - [UNUserNotificationCenterDelegate.UserNotificationCenterDidReceiveNotificationResponseWithCompletionHandler]
type VoidHandler = func()

// NewVoidBlock wraps a Go [VoidHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [UNUserNotificationCenterDelegate.UserNotificationCenterDidReceiveNotificationResponseWithCompletionHandler]
func NewVoidBlock(handler VoidHandler) (objc.ID, func()) {
	block := objc.NewBlock(func(b objc.Block) {
		handler()
	})
	return objc.ID(block), func() { block.Release() }
}
