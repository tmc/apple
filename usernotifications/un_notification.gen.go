// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [UNNotification] class.
var (
	_UNNotificationClass     UNNotificationClass
	_UNNotificationClassOnce sync.Once
)

func getUNNotificationClass() UNNotificationClass {
	_UNNotificationClassOnce.Do(func() {
		_UNNotificationClass = UNNotificationClass{class: objc.GetClass("UNNotification")}
	})
	return _UNNotificationClass
}

// GetUNNotificationClass returns the class object for UNNotification.
func GetUNNotificationClass() UNNotificationClass {
	return getUNNotificationClass()
}

type UNNotificationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (uc UNNotificationClass) Class() objc.Class {
	return uc.class
}

// Alloc allocates memory for a new instance of the class.
func (uc UNNotificationClass) Alloc() UNNotification {
	rv := objc.Send[UNNotification](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// The data for a local or remote notification the system delivers to your
// app.
//
// # Overview
//
// A [UNNotification] object contains the initial notification request, which
// contains the notification’s payload, and the date that the system
// delivered the notification.
//
// Don’t create notification objects directly. When handling notifications,
// the system delivers notification objects to your
// [UNUserNotificationCenterDelegate] object. The [UNUserNotificationCenter]
// object also maintains the list of notifications that the system delivers,
// and you use the [GetDeliveredNotificationsWithCompletionHandler] method to
// retrieve those objects.
//
// # Getting the Notification Details
//
//   - [UNNotification.Request]: The notification request containing the payload and trigger condition for the notification.
//   - [UNNotification.Date]: The delivery date of the notification.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotification
type UNNotification struct {
	objectivec.Object
}

// UNNotificationFromID constructs a [UNNotification] from an objc.ID.
//
// The data for a local or remote notification the system delivers to your
// app.
func UNNotificationFromID(id objc.ID) UNNotification {
	return UNNotification{objectivec.Object{ID: id}}
}

// NOTE: UNNotification adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [UNNotification] class.
//
// # Getting the Notification Details
//
//   - [IUNNotification.Request]: The notification request containing the payload and trigger condition for the notification.
//   - [IUNNotification.Date]: The delivery date of the notification.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotification
type IUNNotification interface {
	objectivec.IObject

	// Topic: Getting the Notification Details

	// The notification request containing the payload and trigger condition for the notification.
	Request() IUNNotificationRequest
	// The delivery date of the notification.
	Date() foundation.INSDate

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (u UNNotification) Init() UNNotification {
	rv := objc.Send[UNNotification](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u UNNotification) Autorelease() UNNotification {
	rv := objc.Send[UNNotification](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewUNNotification creates a new UNNotification instance.
func NewUNNotification() UNNotification {
	class := getUNNotificationClass()
	rv := objc.Send[UNNotification](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (u UNNotification) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](u.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The notification request containing the payload and trigger condition for
// the notification.
//
// # Discussion
//
// For local notifications, the request object is a copy of the one you
// originally configured. For remote notifications, the system synthesizes the
// request object from information received from Apple Push Notification
// service.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotification/request
func (u UNNotification) Request() IUNNotificationRequest {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("request"))
	return UNNotificationRequestFromID(objc.ID(rv))
}

// The delivery date of the notification.
//
// # Discussion
//
// The system displays this date to the user in Notification Center.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotification/date
func (u UNNotification) Date() foundation.INSDate {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("date"))
	return foundation.NSDateFromID(objc.ID(rv))
}
