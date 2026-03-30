// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [UNPushNotificationTrigger] class.
var (
	_UNPushNotificationTriggerClass     UNPushNotificationTriggerClass
	_UNPushNotificationTriggerClassOnce sync.Once
)

func getUNPushNotificationTriggerClass() UNPushNotificationTriggerClass {
	_UNPushNotificationTriggerClassOnce.Do(func() {
		_UNPushNotificationTriggerClass = UNPushNotificationTriggerClass{class: objc.GetClass("UNPushNotificationTrigger")}
	})
	return _UNPushNotificationTriggerClass
}

// GetUNPushNotificationTriggerClass returns the class object for UNPushNotificationTrigger.
func GetUNPushNotificationTriggerClass() UNPushNotificationTriggerClass {
	return getUNPushNotificationTriggerClass()
}

type UNPushNotificationTriggerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (uc UNPushNotificationTriggerClass) Class() objc.Class {
	return uc.class
}

// Alloc allocates memory for a new instance of the class.
func (uc UNPushNotificationTriggerClass) Alloc() UNPushNotificationTrigger {
	rv := objc.Send[UNPushNotificationTrigger](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// A trigger condition that indicates Apple Push Notification Service (APNs)
// has sent the notification.
//
// # Overview
//
// You don’t create instances of this class yourself. The system creates
// [UNPushNotificationTrigger] objects and associates them with requests that
// originated from Apple Push Notification service. You encounter instances of
// this class when managing your app’s delivered notification requests,
// which store an object of this type in their [UNPushNotificationTrigger.Trigger] property.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNPushNotificationTrigger
type UNPushNotificationTrigger struct {
	UNNotificationTrigger
}

// UNPushNotificationTriggerFromID constructs a [UNPushNotificationTrigger] from an objc.ID.
//
// A trigger condition that indicates Apple Push Notification Service (APNs)
// has sent the notification.
func UNPushNotificationTriggerFromID(id objc.ID) UNPushNotificationTrigger {
	return UNPushNotificationTrigger{UNNotificationTrigger: UNNotificationTriggerFromID(id)}
}

// NOTE: UNPushNotificationTrigger adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [UNPushNotificationTrigger] class.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNPushNotificationTrigger
type IUNPushNotificationTrigger interface {
	IUNNotificationTrigger

	// The conditions that trigger the delivery of the notification.
	Trigger() IUNNotificationTrigger
	SetTrigger(value IUNNotificationTrigger)
}

// Init initializes the instance.
func (u UNPushNotificationTrigger) Init() UNPushNotificationTrigger {
	rv := objc.Send[UNPushNotificationTrigger](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u UNPushNotificationTrigger) Autorelease() UNPushNotificationTrigger {
	rv := objc.Send[UNPushNotificationTrigger](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewUNPushNotificationTrigger creates a new UNPushNotificationTrigger instance.
func NewUNPushNotificationTrigger() UNPushNotificationTrigger {
	class := getUNPushNotificationTriggerClass()
	rv := objc.Send[UNPushNotificationTrigger](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The conditions that trigger the delivery of the notification.
//
// See: https://developer.apple.com/documentation/usernotifications/unnotificationrequest/trigger
func (u UNPushNotificationTrigger) Trigger() IUNNotificationTrigger {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("trigger"))
	return UNNotificationTriggerFromID(objc.ID(rv))
}
func (u UNPushNotificationTrigger) SetTrigger(value IUNNotificationTrigger) {
	objc.Send[struct{}](u.ID, objc.Sel("setTrigger:"), value)
}
