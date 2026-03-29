// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [UNNotificationTrigger] class.
var (
	_UNNotificationTriggerClass     UNNotificationTriggerClass
	_UNNotificationTriggerClassOnce sync.Once
)

func getUNNotificationTriggerClass() UNNotificationTriggerClass {
	_UNNotificationTriggerClassOnce.Do(func() {
		_UNNotificationTriggerClass = UNNotificationTriggerClass{class: objc.GetClass("UNNotificationTrigger")}
	})
	return _UNNotificationTriggerClass
}

// GetUNNotificationTriggerClass returns the class object for UNNotificationTrigger.
func GetUNNotificationTriggerClass() UNNotificationTriggerClass {
	return getUNNotificationTriggerClass()
}

type UNNotificationTriggerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (uc UNNotificationTriggerClass) Class() objc.Class {
	return uc.class
}

// Alloc allocates memory for a new instance of the class.
func (uc UNNotificationTriggerClass) Alloc() UNNotificationTrigger {
	rv := objc.Send[UNNotificationTrigger](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// The common behavior for subclasses that trigger the delivery of a local or
// remote notification.
//
// # Overview
// 
// The [UNNotificationTrigger] class is an abstract class for representing an
// event that triggers the delivery of a notification. You don’t create
// instances of this class directly. Instead, you instantiate the concrete
// subclass that defines the trigger condition you want for your notification.
// You then assign the resulting object to the [UNNotificationRequest] object
// that you use to schedule your notification.
// 
// Concrete trigger classes include the following:
// 
// - [UNTimeIntervalNotificationTrigger] - [UNCalendarNotificationTrigger] -
// [UNLocationNotificationTrigger] - [UNPushNotificationTrigger]
//
// [UNLocationNotificationTrigger]: https://developer.apple.com/documentation/UserNotifications/UNLocationNotificationTrigger
//
// # Configuring the Trigger’s Behavior
//
//   - [UNNotificationTrigger.Repeats]: A Boolean value indicating whether the system reschedules the notification after it’s delivered.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationTrigger
type UNNotificationTrigger struct {
	objectivec.Object
}

// UNNotificationTriggerFromID constructs a [UNNotificationTrigger] from an objc.ID.
//
// The common behavior for subclasses that trigger the delivery of a local or
// remote notification.
func UNNotificationTriggerFromID(id objc.ID) UNNotificationTrigger {
	return UNNotificationTrigger{objectivec.Object{ID: id}}
}
// NOTE: UNNotificationTrigger adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [UNNotificationTrigger] class.
//
// # Configuring the Trigger’s Behavior
//
//   - [IUNNotificationTrigger.Repeats]: A Boolean value indicating whether the system reschedules the notification after it’s delivered.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationTrigger
type IUNNotificationTrigger interface {
	objectivec.IObject

	// Topic: Configuring the Trigger’s Behavior

	// A Boolean value indicating whether the system reschedules the notification after it’s delivered.
	Repeats() bool

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (u UNNotificationTrigger) Init() UNNotificationTrigger {
	rv := objc.Send[UNNotificationTrigger](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u UNNotificationTrigger) Autorelease() UNNotificationTrigger {
	rv := objc.Send[UNNotificationTrigger](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewUNNotificationTrigger creates a new UNNotificationTrigger instance.
func NewUNNotificationTrigger() UNNotificationTrigger {
	class := getUNNotificationTriggerClass()
	rv := objc.Send[UNNotificationTrigger](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (u UNNotificationTrigger) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](u.ID, objc.Sel("encodeWithCoder:"), coder)
}

// A Boolean value indicating whether the system reschedules the notification
// after it’s delivered.
//
// # Discussion
// 
// When this property is [false], the system delivers the notification only
// once. When this property is [true], the system reschedules the notification
// request automatically, resulting in the system delivering the notification
// each time the trigger condition is met. To unschedule the notification
// request, use the methods of the [UNUserNotificationCenter] to remove the
// notification request.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationTrigger/repeats
func (u UNNotificationTrigger) Repeats() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("repeats"))
	return rv
}

