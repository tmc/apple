// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)
var _ = fmt.Sprintf

// An interface for processing incoming notifications and responding to notification actions.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenterDelegate
type UNUserNotificationCenterDelegate interface {
	objectivec.IObject
}

// UNUserNotificationCenterDelegateObject wraps an existing Objective-C object that conforms to the UNUserNotificationCenterDelegate protocol.
type UNUserNotificationCenterDelegateObject struct {
	objectivec.Object
}
func (o UNUserNotificationCenterDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// UNUserNotificationCenterDelegateObjectFromID constructs a [UNUserNotificationCenterDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func UNUserNotificationCenterDelegateObjectFromID(id objc.ID) UNUserNotificationCenterDelegateObject {
	return UNUserNotificationCenterDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Asks the delegate to process the user’s response to a delivered
// notification.
//
// center: The shared user notification center object that received the notification.
//
// response: The user’s response to the notification. This object contains the
// original notification and the identifier string for the selected action. If
// the action allowed the user to provide a textual response, this parameter
// contains a [UNTextInputNotificationResponse] object.
//
// completionHandler: The block to execute when you have finished processing the user’s
// response. You must execute this block at some point after processing the
// user’s response to let the system know that you are done. The block has
// no return value or parameters.
//
// # Discussion
// 
// Use this method to process the user’s response to a notification. If the
// user selected one of your app’s custom actions, the `response` parameter
// contains the identifier for that action. (The response can also indicate
// that the user dismissed the notification interface, or launched your app,
// without selecting a custom action.) At the end of your implementation, call
// the `completionHandler` block to let the system know that you are done
// processing the user’s response. If you do not implement this method, your
// app never responds to custom actions.
// 
// You specify your app’s notification types at app launch using
// [UNNotificationCategory] objects, and you specify the custom actions for
// each type using [UNNotificationAction] objects. For information, see
// [Declaring your actionable notification types].
//
// [Declaring your actionable notification types]: https://developer.apple.com/documentation/UserNotifications/declaring-your-actionable-notification-types
//
// See: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenterDelegate/userNotificationCenter(_:didReceive:withCompletionHandler:)
func (o UNUserNotificationCenterDelegateObject) UserNotificationCenterDidReceiveNotificationResponseWithCompletionHandler(center IUNUserNotificationCenter, response IUNNotificationResponse, completionHandler VoidHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("userNotificationCenter:didReceiveNotificationResponse:withCompletionHandler:"), center, response, completionHandler)
	}
// Asks the delegate how to handle a notification that arrived while the app
// was running in the foreground.
//
// center: The shared user notification center object that received the notification.
//
// notification: The notification that is about to be delivered. Use the information in this
// object to determine an appropriate course of action. For example, you might
// use the information to update your app’s interface.
//
// completionHandler: The block to execute with the presentation option for the notification.
// Always execute this block at some point during your implementation of this
// method. Use the `options` parameter to specify how you want the system to
// alert the user, if at all. This block has no return value and takes the
// following parameter:
// 
// options: The option for notifying the user. Specify
// [UNNotificationPresentationOptionNone] to silence the notification
// completely. Specify other values to interact with the user. For a list of
// possible options, see [UNNotificationPresentationOptions].
// //
// [UNNotificationPresentationOptionNone]: https://developer.apple.com/documentation/UserNotifications/UNNotificationPresentationOptionNone
// [UNNotificationPresentationOptions]: https://developer.apple.com/documentation/UserNotifications/UNNotificationPresentationOptions
//
// # Discussion
// 
// If your app is in the foreground when a notification arrives, the shared
// user notification center calls this method to deliver the notification
// directly to your app. If you implement this method, you can take whatever
// actions are necessary to process the notification and update your app. When
// you finish, call the `completionHandler` block and specify how you want the
// system to alert the user, if at all.
// 
// If your delegate does not implement this method, the system behaves as if
// you had passed the [UNNotificationPresentationOptionNone] option to the
// `completionHandler` block. If you do not provide a delegate at all for the
// [UNUserNotificationCenter] object, the system uses the notification’s
// original options to alert the user.
//
// [UNNotificationPresentationOptionNone]: https://developer.apple.com/documentation/UserNotifications/UNNotificationPresentationOptionNone
//
// See: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenterDelegate/userNotificationCenter(_:willPresent:withCompletionHandler:)
func (o UNUserNotificationCenterDelegateObject) UserNotificationCenterWillPresentNotificationWithCompletionHandler(center IUNUserNotificationCenter, notification IUNNotification, completionHandler UNNotificationPresentationOptionsHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("userNotificationCenter:willPresentNotification:withCompletionHandler:"), center, notification, completionHandler)
	}
// Asks the delegate to display the in-app notification settings.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNUserNotificationCenterDelegate/userNotificationCenter(_:openSettingsFor:)
func (o UNUserNotificationCenterDelegateObject) UserNotificationCenterOpenSettingsForNotification(center IUNUserNotificationCenter, notification IUNNotification) {
	objc.Send[struct{}](o.ID, objc.Sel("userNotificationCenter:openSettingsForNotification:"), center, notification)
	}

// UNUserNotificationCenterDelegateConfig holds optional typed callbacks for [UNUserNotificationCenterDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/usernotifications/unusernotificationcenterdelegate
type UNUserNotificationCenterDelegateConfig struct {

	// Other Methods
	// UserNotificationCenterOpenSettingsForNotification — Asks the delegate to display the in-app notification settings.
	UserNotificationCenterOpenSettingsForNotification func(center UNUserNotificationCenter, notification UNNotification)
}

// NewUNUserNotificationCenterDelegate creates an Objective-C object implementing the [UNUserNotificationCenterDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [UNUserNotificationCenterDelegateObject] satisfies the [UNUserNotificationCenterDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/usernotifications/unusernotificationcenterdelegate
func NewUNUserNotificationCenterDelegate(config UNUserNotificationCenterDelegateConfig) UNUserNotificationCenterDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoUNUserNotificationCenterDelegate_%d", n)

	var methods []objc.MethodDef

	if config.UserNotificationCenterOpenSettingsForNotification != nil {
		fn := config.UserNotificationCenterOpenSettingsForNotification
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("userNotificationCenter:openSettingsForNotification:"),
			Fn: func(self objc.ID, _cmd objc.SEL, centerID objc.ID, notificationID objc.ID) {
				center := UNUserNotificationCenterFromID(centerID)
				notification := UNNotificationFromID(notificationID)
				fn(center, notification)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("UNUserNotificationCenterDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewUNUserNotificationCenterDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return UNUserNotificationCenterDelegateObjectFromID(instance)
}

