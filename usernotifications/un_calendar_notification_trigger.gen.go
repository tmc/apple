// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [UNCalendarNotificationTrigger] class.
var (
	_UNCalendarNotificationTriggerClass     UNCalendarNotificationTriggerClass
	_UNCalendarNotificationTriggerClassOnce sync.Once
)

func getUNCalendarNotificationTriggerClass() UNCalendarNotificationTriggerClass {
	_UNCalendarNotificationTriggerClassOnce.Do(func() {
		_UNCalendarNotificationTriggerClass = UNCalendarNotificationTriggerClass{class: objc.GetClass("UNCalendarNotificationTrigger")}
	})
	return _UNCalendarNotificationTriggerClass
}

// GetUNCalendarNotificationTriggerClass returns the class object for UNCalendarNotificationTrigger.
func GetUNCalendarNotificationTriggerClass() UNCalendarNotificationTriggerClass {
	return getUNCalendarNotificationTriggerClass()
}

type UNCalendarNotificationTriggerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (uc UNCalendarNotificationTriggerClass) Class() objc.Class {
	return uc.class
}

// Alloc allocates memory for a new instance of the class.
func (uc UNCalendarNotificationTriggerClass) Alloc() UNCalendarNotificationTrigger {
	rv := objc.Send[UNCalendarNotificationTrigger](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// A trigger condition that causes a notification the system delivers at a
// specific date and time.
//
// # Overview
//
// Create a [UNCalendarNotificationTrigger] object when you want to schedule
// the delivery of a local notification at the date and time you specify. You
// use an [NSDateComponents] object to specify only the time values that you
// want the system to use to determine the matching date and time.
//
// Listing 1 creates a trigger that delivers its notification every morning at
// 8:30. The repeating behavior is achieved by specifying `true` for the
// `repeats` parameter when creating the trigger.
//
// Listing 1. Creating a trigger that repeats at a specific time
//
// # Getting the Trigger Information
//
//   - [UNCalendarNotificationTrigger.NextTriggerDate]: The next date at which the trigger conditions are met.
//   - [UNCalendarNotificationTrigger.DateComponents]: The date components to construct this object.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNCalendarNotificationTrigger
//
// [NSDateComponents]: https://developer.apple.com/documentation/Foundation/NSDateComponents
type UNCalendarNotificationTrigger struct {
	UNNotificationTrigger
}

// UNCalendarNotificationTriggerFromID constructs a [UNCalendarNotificationTrigger] from an objc.ID.
//
// A trigger condition that causes a notification the system delivers at a
// specific date and time.
func UNCalendarNotificationTriggerFromID(id objc.ID) UNCalendarNotificationTrigger {
	return UNCalendarNotificationTrigger{UNNotificationTrigger: UNNotificationTriggerFromID(id)}
}

// NOTE: UNCalendarNotificationTrigger adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [UNCalendarNotificationTrigger] class.
//
// # Getting the Trigger Information
//
//   - [IUNCalendarNotificationTrigger.NextTriggerDate]: The next date at which the trigger conditions are met.
//   - [IUNCalendarNotificationTrigger.DateComponents]: The date components to construct this object.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNCalendarNotificationTrigger
type IUNCalendarNotificationTrigger interface {
	IUNNotificationTrigger

	// Topic: Getting the Trigger Information

	// The next date at which the trigger conditions are met.
	NextTriggerDate() foundation.INSDate
	// The date components to construct this object.
	DateComponents() foundation.NSDateComponents
}

// Init initializes the instance.
func (u UNCalendarNotificationTrigger) Init() UNCalendarNotificationTrigger {
	rv := objc.Send[UNCalendarNotificationTrigger](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u UNCalendarNotificationTrigger) Autorelease() UNCalendarNotificationTrigger {
	rv := objc.Send[UNCalendarNotificationTrigger](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewUNCalendarNotificationTrigger creates a new UNCalendarNotificationTrigger instance.
func NewUNCalendarNotificationTrigger() UNCalendarNotificationTrigger {
	class := getUNCalendarNotificationTriggerClass()
	rv := objc.Send[UNCalendarNotificationTrigger](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a calendar trigger using the date components parameter.
//
// dateComponents: The temporal information to use when constructing the trigger. Provide only
// the date components that are relevant for your trigger.
//
// repeats: Specify false to deliver the notification one time. Specify true to
// reschedule the notification request each time the system delivers the
// notification.
//
// # Return Value
//
// A new calendar trigger based on the specified temporal information.
//
// # Discussion
//
// If you specify `true` for the `repeats` parameter, you must explicitly
// remove the notification request to stop the delivery of the associated
// notification. Use the methods of [UNUserNotificationCenter] to remove
// notification requests that are no longer needed.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNCalendarNotificationTrigger/init(dateMatching:repeats:)
func NewUNCalendarNotificationTriggerWithDateMatchingComponentsRepeats(dateComponents foundation.NSDateComponents, repeats bool) UNCalendarNotificationTrigger {
	rv := objc.Send[objc.ID](objc.ID(getUNCalendarNotificationTriggerClass().class), objc.Sel("triggerWithDateMatchingComponents:repeats:"), dateComponents, repeats)
	return UNCalendarNotificationTriggerFromID(rv)
}

// The next date at which the trigger conditions are met.
//
// # Return Value
//
// The next trigger date.
//
// # Discussion
//
// Use this property to find out when the system will deliver a notification
// associated with this trigger.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNCalendarNotificationTrigger/nextTriggerDate()
func (u UNCalendarNotificationTrigger) NextTriggerDate() foundation.INSDate {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("nextTriggerDate"))
	return foundation.NSDateFromID(rv)
}

// The date components to construct this object.
//
// # Discussion
//
// Use this property to review the date components associated with this
// trigger.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNCalendarNotificationTrigger/dateComponents
func (u UNCalendarNotificationTrigger) DateComponents() foundation.NSDateComponents {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("dateComponents"))
	return foundation.NSDateComponentsFromID(objc.ID(rv))
}
