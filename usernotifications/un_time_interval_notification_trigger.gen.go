// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [UNTimeIntervalNotificationTrigger] class.
var (
	_UNTimeIntervalNotificationTriggerClass     UNTimeIntervalNotificationTriggerClass
	_UNTimeIntervalNotificationTriggerClassOnce sync.Once
)

func getUNTimeIntervalNotificationTriggerClass() UNTimeIntervalNotificationTriggerClass {
	_UNTimeIntervalNotificationTriggerClassOnce.Do(func() {
		_UNTimeIntervalNotificationTriggerClass = UNTimeIntervalNotificationTriggerClass{class: objc.GetClass("UNTimeIntervalNotificationTrigger")}
	})
	return _UNTimeIntervalNotificationTriggerClass
}

// GetUNTimeIntervalNotificationTriggerClass returns the class object for UNTimeIntervalNotificationTrigger.
func GetUNTimeIntervalNotificationTriggerClass() UNTimeIntervalNotificationTriggerClass {
	return getUNTimeIntervalNotificationTriggerClass()
}

type UNTimeIntervalNotificationTriggerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (uc UNTimeIntervalNotificationTriggerClass) Class() objc.Class {
	return uc.class
}

// Alloc allocates memory for a new instance of the class.
func (uc UNTimeIntervalNotificationTriggerClass) Alloc() UNTimeIntervalNotificationTrigger {
	rv := objc.Send[UNTimeIntervalNotificationTrigger](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// A trigger condition that causes the system to deliver a notification after
// the amount of time you specify elapses.
//
// # Overview
// 
// Create a [UNTimeIntervalNotificationTrigger] object when you want to
// schedule the delivery of a local notification after the number of seconds
// you specify elapses. You use this type of trigger to implement timers.
// 
// Listing 1 creates a trigger that delivers its notification one time after
// 30 minutes have elapsed.
// 
// Listing 1. Creating a trigger that fires in 30 minutes
//
// # Getting the Trigger Information
//
//   - [UNTimeIntervalNotificationTrigger.NextTriggerDate]: The next date at which the trigger conditions are met.
//   - [UNTimeIntervalNotificationTrigger.TimeInterval]: The time interval to create the trigger.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNTimeIntervalNotificationTrigger
type UNTimeIntervalNotificationTrigger struct {
	UNNotificationTrigger
}

// UNTimeIntervalNotificationTriggerFromID constructs a [UNTimeIntervalNotificationTrigger] from an objc.ID.
//
// A trigger condition that causes the system to deliver a notification after
// the amount of time you specify elapses.
func UNTimeIntervalNotificationTriggerFromID(id objc.ID) UNTimeIntervalNotificationTrigger {
	return UNTimeIntervalNotificationTrigger{UNNotificationTrigger: UNNotificationTriggerFromID(id)}
}
// NOTE: UNTimeIntervalNotificationTrigger adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [UNTimeIntervalNotificationTrigger] class.
//
// # Getting the Trigger Information
//
//   - [IUNTimeIntervalNotificationTrigger.NextTriggerDate]: The next date at which the trigger conditions are met.
//   - [IUNTimeIntervalNotificationTrigger.TimeInterval]: The time interval to create the trigger.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNTimeIntervalNotificationTrigger
type IUNTimeIntervalNotificationTrigger interface {
	IUNNotificationTrigger

	// Topic: Getting the Trigger Information

	// The next date at which the trigger conditions are met.
	NextTriggerDate() foundation.INSDate
	// The time interval to create the trigger.
	TimeInterval() float64
}

// Init initializes the instance.
func (u UNTimeIntervalNotificationTrigger) Init() UNTimeIntervalNotificationTrigger {
	rv := objc.Send[UNTimeIntervalNotificationTrigger](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u UNTimeIntervalNotificationTrigger) Autorelease() UNTimeIntervalNotificationTrigger {
	rv := objc.Send[UNTimeIntervalNotificationTrigger](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewUNTimeIntervalNotificationTrigger creates a new UNTimeIntervalNotificationTrigger instance.
func NewUNTimeIntervalNotificationTrigger() UNTimeIntervalNotificationTrigger {
	class := getUNTimeIntervalNotificationTriggerClass()
	rv := objc.Send[UNTimeIntervalNotificationTrigger](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a time interval trigger using the time value parameter.
//
// timeInterval: The time (in seconds) that must elapse from the current time before the
// trigger fires. This value must be greater than zero.
//
// repeats: Specify [false] to deliver the notification one time. Specify [true] to
// reschedule the notification request each time the system delivers the
// notification. If this parameter is [true], the value in the `timeInterval`
// parameter must be 60 seconds or greater.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// A new time interval trigger based on the specified temporal information.
//
// # Discussion
// 
// If you specify `true` for the `repeats` parameter, you must explicitly
// remove the notification request to stop the delivery of the associated
// notification. Use the methods of [UNUserNotificationCenter] to remove
// notification requests that are no longer needed.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNTimeIntervalNotificationTrigger/init(timeInterval:repeats:)
func NewUNTimeIntervalNotificationTriggerWithTimeIntervalRepeats(timeInterval float64, repeats bool) UNTimeIntervalNotificationTrigger {
	rv := objc.Send[objc.ID](objc.ID(getUNTimeIntervalNotificationTriggerClass().class), objc.Sel("triggerWithTimeInterval:repeats:"), timeInterval, repeats)
	return UNTimeIntervalNotificationTriggerFromID(rv)
}

// The next date at which the trigger conditions are met.
//
// # Return Value
// 
// The next trigger date.
//
// # Discussion
// 
// Use this property to find out when a notification associated with this
// trigger will next be delivered.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNTimeIntervalNotificationTrigger/nextTriggerDate()
func (u UNTimeIntervalNotificationTrigger) NextTriggerDate() foundation.INSDate {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("nextTriggerDate"))
	return foundation.NSDateFromID(rv)
}

// The time interval to create the trigger.
//
// # Discussion
// 
// This property contains the original time interval that you specified when
// creating the trigger object. The value in this property isn’t updated as
// time counts down. To find out when the trigger will fire next, call the
// [NextTriggerDate] method.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNTimeIntervalNotificationTrigger/timeInterval
func (u UNTimeIntervalNotificationTrigger) TimeInterval() float64 {
	rv := objc.Send[float64](u.ID, objc.Sel("timeInterval"))
	return rv
}

