// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [UNTextInputNotificationResponse] class.
var (
	_UNTextInputNotificationResponseClass     UNTextInputNotificationResponseClass
	_UNTextInputNotificationResponseClassOnce sync.Once
)

func getUNTextInputNotificationResponseClass() UNTextInputNotificationResponseClass {
	_UNTextInputNotificationResponseClassOnce.Do(func() {
		_UNTextInputNotificationResponseClass = UNTextInputNotificationResponseClass{class: objc.GetClass("UNTextInputNotificationResponse")}
	})
	return _UNTextInputNotificationResponseClass
}

// GetUNTextInputNotificationResponseClass returns the class object for UNTextInputNotificationResponse.
func GetUNTextInputNotificationResponseClass() UNTextInputNotificationResponseClass {
	return getUNTextInputNotificationResponseClass()
}

type UNTextInputNotificationResponseClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (uc UNTextInputNotificationResponseClass) Class() objc.Class {
	return uc.class
}

// Alloc allocates memory for a new instance of the class.
func (uc UNTextInputNotificationResponseClass) Alloc() UNTextInputNotificationResponse {
	rv := objc.Send[UNTextInputNotificationResponse](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// The user’s response to an actionable notification, including any custom
// text that the user typed or dictated.
//
// # Overview
//
// The system delivers a [UNTextInputNotificationResponse] object to your app
// so that you can process user-provided text content. When defining your
// categories, you can specify an [UNTextInputNotificationAction] object
// instead of an [UNNotificationAction] object for your action. If you do, the
// system creates an [UNTextInputNotificationResponse] object when the user
// selects the accompanying action, and it fills the [UNTextInputNotificationResponse.UserText] property with
// any user-entered text.
//
// You don’t create [UNTextInputNotificationResponse] objects yourself.
// Instead, the shared user notification center object creates them and
// delivers them to the
// [UserNotificationCenterDidReceiveNotificationResponseWithCompletionHandler]
// method of its delegate object. Use that method to extract any needed
// information from the response object and take appropriate action.
//
// For more information about responding to actions, see [Handling
// notifications and notification-related actions].
//
// # Getting the Text Response
//
//   - [UNTextInputNotificationResponse.UserText]: The text response provided by the user.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNTextInputNotificationResponse
//
// [Handling notifications and notification-related actions]: https://developer.apple.com/documentation/UserNotifications/handling-notifications-and-notification-related-actions
type UNTextInputNotificationResponse struct {
	UNNotificationResponse
}

// UNTextInputNotificationResponseFromID constructs a [UNTextInputNotificationResponse] from an objc.ID.
//
// The user’s response to an actionable notification, including any custom
// text that the user typed or dictated.
func UNTextInputNotificationResponseFromID(id objc.ID) UNTextInputNotificationResponse {
	return UNTextInputNotificationResponse{UNNotificationResponse: UNNotificationResponseFromID(id)}
}

// NOTE: UNTextInputNotificationResponse adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [UNTextInputNotificationResponse] class.
//
// # Getting the Text Response
//
//   - [IUNTextInputNotificationResponse.UserText]: The text response provided by the user.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNTextInputNotificationResponse
type IUNTextInputNotificationResponse interface {
	IUNNotificationResponse

	// Topic: Getting the Text Response

	// The text response provided by the user.
	UserText() string
}

// Init initializes the instance.
func (u UNTextInputNotificationResponse) Init() UNTextInputNotificationResponse {
	rv := objc.Send[UNTextInputNotificationResponse](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u UNTextInputNotificationResponse) Autorelease() UNTextInputNotificationResponse {
	rv := objc.Send[UNTextInputNotificationResponse](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewUNTextInputNotificationResponse creates a new UNTextInputNotificationResponse instance.
func NewUNTextInputNotificationResponse() UNTextInputNotificationResponse {
	class := getUNTextInputNotificationResponseClass()
	rv := objc.Send[UNTextInputNotificationResponse](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The text response provided by the user.
//
// # Discussion
//
// If the user does not specify any text, this property contains an empty
// string.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNTextInputNotificationResponse/userText
func (u UNTextInputNotificationResponse) UserText() string {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("userText"))
	return foundation.NSStringFromID(rv).String()
}
