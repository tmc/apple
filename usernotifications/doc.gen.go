
// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

// Package usernotifications provides Go bindings for the UserNotifications framework.
//
// Push user-facing notifications to the user’s device from a server, or generate them locally from your app.
//
// User-facing notifications communicate important information to users of your app, regardless of whether your app is running on the user’s device. For example, a sports app can let the user know when their favorite team scores. Notifications can also tell your app to download information and update its interface. Notifications can display an alert, play a sound, or badge the app’s icon.
//
// # Essentials
//
//   - User Notifications updates: Learn about important changes in User Notifications.
//   - Asking permission to use notifications: Request permission to display alerts, play sounds, or badge the app’s icon in response to a notification.
//
// # Notification management
//
//   - UNUserNotificationCenter: The central object for managing notification-related activities for your app or app extension. ([UNAuthorizationOptions], [UNUserNotificationCenterDelegate])
//   - UNUserNotificationCenterDelegate: An interface for processing incoming notifications and responding to notification actions. ([UNNotificationPresentationOptions])
//   - UNNotificationSettings: The object for managing notification-related settings and the authorization status of your app. ([UNAuthorizationStatus], [UNNotificationSetting], [UNAlertStyle], [UNShowPreviewsSetting])
//
// # Remote notifications
//
//   - Setting up a remote notification server: Generate notifications and push them to user devices.
//   - Sending push notifications using command-line tools: Use basic macOS command-line tools to send push notifications to Apple Push Notification service (APNs).
//   - Testing notifications using the Push Notification Console: Send test notifications and access delivery logs to test your app’s integration with Apple Push Notification service (APNs).
//
// # Notification requests
//
//   - Scheduling a notification locally from your app: Create and schedule notifications from your app when you want to get the user’s attention.
//   - UNNotificationRequest: A request to schedule a local notification, which includes the content of the notification and the trigger conditions for delivery.
//   - UNNotification: The data for a local or remote notification the system delivers to your app.
//
// # Push notifications in safari
//
//   - Sending web push notifications in web apps and browsers: Update your web server and website to send push notifications that work in Safari, other browsers, and web apps, following cross-browser standards.
//
// # Notification content
//
//   - Implementing communication notifications: Configure and display your app’s communication notifications by using intents.
//   - UNNotificationContentProviding: A protocol the system uses to provide context relevant to user notifications.
//   - UNNotificationActionIcon: An icon associated with an action.
//   - UNMutableNotificationContent: The editable content for a notification. ([UNNotificationInterruptionLevel])
//   - UNNotificationContent: The uneditable content of a notification. ([UNNotificationInterruptionLevel])
//   - UNNotificationAttachment: A media file associated with a notification.
//   - UNNotificationSound: The sound played upon delivery of a notification.
//   - UNNotificationSoundName: A string providing the name of a sound file.
//
// # Triggers
//
//   - UNCalendarNotificationTrigger: A trigger condition that causes a notification the system delivers at a specific date and time.
//   - UNTimeIntervalNotificationTrigger: A trigger condition that causes the system to deliver a notification after the amount of time you specify elapses.
//   - UNLocationNotificationTrigger: A trigger condition that causes the system to deliver a notification when the user’s device enters or exits a geographic region you specify.
//   - UNPushNotificationTrigger: A trigger condition that indicates Apple Push Notification Service (APNs) has sent the notification.
//   - UNNotificationTrigger: The common behavior for subclasses that trigger the delivery of a local or remote notification.
//
// # Notification categories and user actions
//
//   - Declaring your actionable notification types: Differentiate your notifications and add action buttons to the notification interface.
//   - UNNotificationCategory: A type of notification your app supports and the custom actions that the system displays. ([UNNotificationCategoryOptions])
//   - UNNotificationAction: A task your app performs in response to a notification that the system delivers. ([UNNotificationActionOptions])
//   - UNTextInputNotificationAction: An action that accepts user-typed text.
//
// # Notification responses
//
//   - Handling notifications and notification-related actions: Respond to user interactions with the system’s notification interfaces, including handling your app’s custom actions.
//   - UNNotificationResponse: The user’s response to an actionable notification.
//   - UNTextInputNotificationResponse: The user’s response to an actionable notification, including any custom text that the user typed or dictated.
//
// # Notification service app extension
//
//   - Modifying content in newly delivered notifications: Modify the payload of a remote notification before it’s displayed on the user’s iOS device.
//   - UNNotificationServiceExtension: An object that modifies the content of a remote notification before it’s delivered to the user.
//
// # Entitlements
//
//   - APS Environment Entitlement: The environment for push notifications.
//   - APS Environment (macOS) Entitlement: The environment for push notifications in macOS apps.
//
// # Sample code
//
//   - Handling Communication Notifications and Focus Status Updates: Create a richer calling and messaging experience in your app by implementing communication notifications and Focus status updates.
//   - Implementing Alert Push Notifications: Add visible alert notifications to your app by using the UserNotifications framework.
//   - Implementing Background Push Notifications: Add background notifications to your app by using the UserNotifications framework.
//
// # Classes
//
//   - UNNotificationAttributedMessageContext
//
// # Key Types
//
//   - [UNNotificationContent] - The uneditable content of a notification.
//   - [UNMutableNotificationContent] - The editable content for a notification.
//   - [UNUserNotificationCenter] - The central object for managing notification-related activities for your app or app extension.
//   - [UNNotificationSettings] - The object for managing notification-related settings and the authorization status of your app.
//   - [UNNotificationCategory] - A type of notification your app supports and the custom actions that the system displays.
//   - [UNNotificationSound] - The sound played upon delivery of a notification.
//   - [UNNotificationAttachment] - A media file associated with a notification.
//   - [UNNotificationAction] - A task your app performs in response to a notification that the system delivers.
//   - [UNTextInputNotificationAction] - An action that accepts user-typed text.
//   - [UNNotificationResponse] - The user’s response to an actionable notification.
//
// [UserNotifications Documentation]: https://developer.apple.com/documentation/UserNotifications
package usernotifications

import (
	"fmt"
	"os"
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/UserNotifications.framework/UserNotifications"
// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	var err error
	frameworkHandle, err = purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: UserNotifications: failed to load %s: %v\n", frameworkPath, err)
		return 
	}
}

