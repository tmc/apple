// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [UNNotificationRequest] class.
var (
	_UNNotificationRequestClass     UNNotificationRequestClass
	_UNNotificationRequestClassOnce sync.Once
)

func getUNNotificationRequestClass() UNNotificationRequestClass {
	_UNNotificationRequestClassOnce.Do(func() {
		_UNNotificationRequestClass = UNNotificationRequestClass{class: objc.GetClass("UNNotificationRequest")}
	})
	return _UNNotificationRequestClass
}

// GetUNNotificationRequestClass returns the class object for UNNotificationRequest.
func GetUNNotificationRequestClass() UNNotificationRequestClass {
	return getUNNotificationRequestClass()
}

type UNNotificationRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (uc UNNotificationRequestClass) Class() objc.Class {
	return uc.class
}

// Alloc allocates memory for a new instance of the class.
func (uc UNNotificationRequestClass) Alloc() UNNotificationRequest {
	rv := objc.Send[UNNotificationRequest](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// A request to schedule a local notification, which includes the content of
// the notification and the trigger conditions for delivery.
//
// # Overview
//
// Create a [UNNotificationRequest] object when you want to schedule the
// delivery of a local notification. A notification request object contains a
// [UNNotificationContent] object with the payload and the
// [UNNotificationTrigger] object with the conditions that trigger the
// delivery of the notification. To schedule the delivery of your
// notification, pass your request object to the
// [AddNotificationRequestWithCompletionHandler] method of the shared user
// notification center object.
//
// After scheduling a request, you interact with [UNNotificationRequest]
// objects in the following ways:
//
// - View your app’s pending notifications by calling the
// [GetPendingNotificationRequestsWithCompletionHandler] method of your shared
// user notification center object. - When the system delivers a notification
// to your app, the provided [UNNotification] object contains a
// [UNNotificationRequest] object that you can inspect to get the notification
// details. - Use the request’s [UNNotificationRequest.Identifier] to remove delivered
// notifications from Notification Center.
//
// When receiving a local or remote notification, use the provided
// [UNNotificationRequest] object to fetch details about the notification.
//
// # Getting the Request Details
//
//   - [UNNotificationRequest.Identifier]: The unique identifier for this notification request.
//   - [UNNotificationRequest.Content]: The content associated with the notification.
//   - [UNNotificationRequest.Trigger]: The conditions that trigger the delivery of the notification.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationRequest
type UNNotificationRequest struct {
	objectivec.Object
}

// UNNotificationRequestFromID constructs a [UNNotificationRequest] from an objc.ID.
//
// A request to schedule a local notification, which includes the content of
// the notification and the trigger conditions for delivery.
func UNNotificationRequestFromID(id objc.ID) UNNotificationRequest {
	return UNNotificationRequest{objectivec.Object{ID: id}}
}

// NOTE: UNNotificationRequest adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [UNNotificationRequest] class.
//
// # Getting the Request Details
//
//   - [IUNNotificationRequest.Identifier]: The unique identifier for this notification request.
//   - [IUNNotificationRequest.Content]: The content associated with the notification.
//   - [IUNNotificationRequest.Trigger]: The conditions that trigger the delivery of the notification.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationRequest
type IUNNotificationRequest interface {
	objectivec.IObject

	// Topic: Getting the Request Details

	// The unique identifier for this notification request.
	Identifier() string
	// The content associated with the notification.
	Content() IUNNotificationContent
	// The conditions that trigger the delivery of the notification.
	Trigger() IUNNotificationTrigger

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (u UNNotificationRequest) Init() UNNotificationRequest {
	rv := objc.Send[UNNotificationRequest](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u UNNotificationRequest) Autorelease() UNNotificationRequest {
	rv := objc.Send[UNNotificationRequest](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewUNNotificationRequest creates a new UNNotificationRequest instance.
func NewUNNotificationRequest() UNNotificationRequest {
	class := getUNNotificationRequestClass()
	rv := objc.Send[UNNotificationRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a notification request object that you use to schedule a
// notification.
//
// identifier: An identifier for the request; this parameter must not be `nil`. You can
// use this identifier to cancel the request if it’s still pending (see the
// [RemovePendingNotificationRequestsWithIdentifiers] method).
//
// content: The content of the notification. This parameter must not be `nil`.
//
// trigger: The condition that causes the system to deliver the notification. Specify
// `nil` to deliver the notification right away.
//
// # Return Value
//
// A new notification request object.
//
// # Discussion
//
// Use this method when you want to schedule the delivery of a local
// notification. This method creates the request object that you subsequently
// pass to the [AddNotificationRequestWithCompletionHandler] method.
//
// The system uses the `identifier` parameter to determine how to handle the
// request:
//
// - the system creates a new notification. - the system alerts the user
// again, replaces the old notification with the new one, and places the new
// notification at the top of the list. - the new request replaces the pending
// request.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationRequest/init(identifier:content:trigger:)
func NewUNNotificationRequestWithIdentifierContentTrigger(identifier string, content IUNNotificationContent, trigger IUNNotificationTrigger) UNNotificationRequest {
	rv := objc.Send[objc.ID](objc.ID(getUNNotificationRequestClass().class), objc.Sel("requestWithIdentifier:content:trigger:"), objc.String(identifier), content, trigger)
	return UNNotificationRequestFromID(rv)
}

func (u UNNotificationRequest) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](u.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The unique identifier for this notification request.
//
// # Discussion
//
// Use this string to identify notifications in your app. For example, you can
// pass this string to the [RemovePendingNotificationRequestsWithIdentifiers]
// method to cancel a previously scheduled notification.
//
// If you use the same identifier when scheduling a new notification, the
// system removes the previously scheduled notification with that identifier
// and replaces it with the new one.
//
// For local notifications, the system sets this property to the value passed
// to the request’s initializer (see the
// [RequestWithIdentifierContentTrigger] method). For remote notifications,
// the system sets this property to the value of the `apns-collapse-id` key
// that you specified in the APNs request header when generating the remote
// notification. If your app doesn’t set a value, the system automatically
// assigns an identifier.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationRequest/identifier
func (u UNNotificationRequest) Identifier() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("identifier"))
	return foundation.NSStringFromID(rv).String()
}

// The content associated with the notification.
//
// # Discussion
//
// Use this property to access the contents of the notification.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationRequest/content
func (u UNNotificationRequest) Content() IUNNotificationContent {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("content"))
	return UNNotificationContentFromID(objc.ID(rv))
}

// The conditions that trigger the delivery of the notification.
//
// # Discussion
//
// For notifications that the system has delivered, use this property to
// determine what caused the delivery to occur. For remote notifications, this
// property contains a [UNPushNotificationTrigger] object. For other
// notifications, the system sets this type using the trigger condition
// specified in the original request.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationRequest/trigger
func (u UNNotificationRequest) Trigger() IUNNotificationTrigger {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("trigger"))
	return UNNotificationTriggerFromID(objc.ID(rv))
}
