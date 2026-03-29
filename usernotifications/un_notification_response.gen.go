// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [UNNotificationResponse] class.
var (
	_UNNotificationResponseClass     UNNotificationResponseClass
	_UNNotificationResponseClassOnce sync.Once
)

func getUNNotificationResponseClass() UNNotificationResponseClass {
	_UNNotificationResponseClassOnce.Do(func() {
		_UNNotificationResponseClass = UNNotificationResponseClass{class: objc.GetClass("UNNotificationResponse")}
	})
	return _UNNotificationResponseClass
}

// GetUNNotificationResponseClass returns the class object for UNNotificationResponse.
func GetUNNotificationResponseClass() UNNotificationResponseClass {
	return getUNNotificationResponseClass()
}

type UNNotificationResponseClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (uc UNNotificationResponseClass) Class() objc.Class {
	return uc.class
}

// Alloc allocates memory for a new instance of the class.
func (uc UNNotificationResponseClass) Alloc() UNNotificationResponse {
	rv := objc.Send[UNNotificationResponse](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// The user’s response to an actionable notification.
//
// # Overview
// 
// When the user interacts with a delivered notification, the system delivers
// a [UNNotificationResponse] object to your app so that you can process the
// response. Users can interact with delivered notifications in many ways. If
// the notification’s category had associated action buttons, they can
// select one of those buttons. Users can also dismiss the notification
// without selecting one of your actions and they can open your app. A
// response object tells you which option the user selected.
// 
// You don’t create [UNNotificationResponse] objects yourself. Instead, the
// shared user notification center object creates them and delivers them to
// the
// [UserNotificationCenterDidReceiveNotificationResponseWithCompletionHandler]
// method of its delegate object. Use that method to extract any needed
// information from the response object and take appropriate action.
// 
// For more information about responding to actions, see [Handling
// notifications and notification-related actions].
//
// [Handling notifications and notification-related actions]: https://developer.apple.com/documentation/UserNotifications/handling-notifications-and-notification-related-actions
//
// # Getting the Response Information
//
//   - [UNNotificationResponse.ActionIdentifier]: The identifier string of the action that the user selected.
//   - [UNNotificationResponse.Notification]: The notification to which the user responded.
//   - [UNNotificationResponse.UNNotificationDefaultActionIdentifier]: An action that indicates the user opened the app from the notification interface.
//   - [UNNotificationResponse.UNNotificationDismissActionIdentifier]: The action that indicates the user explicitly dismissed the notification interface.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationResponse
type UNNotificationResponse struct {
	objectivec.Object
}

// UNNotificationResponseFromID constructs a [UNNotificationResponse] from an objc.ID.
//
// The user’s response to an actionable notification.
func UNNotificationResponseFromID(id objc.ID) UNNotificationResponse {
	return UNNotificationResponse{objectivec.Object{ID: id}}
}
// NOTE: UNNotificationResponse adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [UNNotificationResponse] class.
//
// # Getting the Response Information
//
//   - [IUNNotificationResponse.ActionIdentifier]: The identifier string of the action that the user selected.
//   - [IUNNotificationResponse.Notification]: The notification to which the user responded.
//   - [IUNNotificationResponse.UNNotificationDefaultActionIdentifier]: An action that indicates the user opened the app from the notification interface.
//   - [IUNNotificationResponse.UNNotificationDismissActionIdentifier]: The action that indicates the user explicitly dismissed the notification interface.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationResponse
type IUNNotificationResponse interface {
	objectivec.IObject

	// Topic: Getting the Response Information

	// The identifier string of the action that the user selected.
	ActionIdentifier() string
	// The notification to which the user responded.
	Notification() IUNNotification
	// An action that indicates the user opened the app from the notification interface.
	UNNotificationDefaultActionIdentifier() string
	// The action that indicates the user explicitly dismissed the notification interface.
	UNNotificationDismissActionIdentifier() string

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (u UNNotificationResponse) Init() UNNotificationResponse {
	rv := objc.Send[UNNotificationResponse](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u UNNotificationResponse) Autorelease() UNNotificationResponse {
	rv := objc.Send[UNNotificationResponse](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewUNNotificationResponse creates a new UNNotificationResponse instance.
func NewUNNotificationResponse() UNNotificationResponse {
	class := getUNNotificationResponseClass()
	rv := objc.Send[UNNotificationResponse](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (u UNNotificationResponse) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](u.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The identifier string of the action that the user selected.
//
// # Discussion
// 
// This parameter may contain one the identifier of one of your
// [UNNotificationAction] objects or it may contain a system-defined
// identifier. The system defined identifiers are
// [UNNotificationDefaultActionIdentifier] and
// [UNNotificationDismissActionIdentifier], which indicate that the user
// opened the app or dismissed the notification without any further actions.
// 
// For more information about defining custom actions, see [Declaring your
// actionable notification types].
//
// [Declaring your actionable notification types]: https://developer.apple.com/documentation/UserNotifications/declaring-your-actionable-notification-types
// [UNNotificationDefaultActionIdentifier]: https://developer.apple.com/documentation/UserNotifications/UNNotificationDefaultActionIdentifier
// [UNNotificationDismissActionIdentifier]: https://developer.apple.com/documentation/UserNotifications/UNNotificationDismissActionIdentifier
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationResponse/actionIdentifier
func (u UNNotificationResponse) ActionIdentifier() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("actionIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
// The notification to which the user responded.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationResponse/notification
func (u UNNotificationResponse) Notification() IUNNotification {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("notification"))
	return UNNotificationFromID(objc.ID(rv))
}
// An action that indicates the user opened the app from the notification
// interface.
//
// See: https://developer.apple.com/documentation/usernotifications/unnotificationdefaultactionidentifier
func (u UNNotificationResponse) UNNotificationDefaultActionIdentifier() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("UNNotificationDefaultActionIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
// The action that indicates the user explicitly dismissed the notification
// interface.
//
// See: https://developer.apple.com/documentation/usernotifications/unnotificationdismissactionidentifier
func (u UNNotificationResponse) UNNotificationDismissActionIdentifier() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("UNNotificationDismissActionIdentifier"))
	return foundation.NSStringFromID(rv).String()
}

