// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [UNMutableNotificationContent] class.
var (
	_UNMutableNotificationContentClass     UNMutableNotificationContentClass
	_UNMutableNotificationContentClassOnce sync.Once
)

func getUNMutableNotificationContentClass() UNMutableNotificationContentClass {
	_UNMutableNotificationContentClassOnce.Do(func() {
		_UNMutableNotificationContentClass = UNMutableNotificationContentClass{class: objc.GetClass("UNMutableNotificationContent")}
	})
	return _UNMutableNotificationContentClass
}

// GetUNMutableNotificationContentClass returns the class object for UNMutableNotificationContent.
func GetUNMutableNotificationContentClass() UNMutableNotificationContentClass {
	return getUNMutableNotificationContentClass()
}

type UNMutableNotificationContentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (uc UNMutableNotificationContentClass) Class() objc.Class {
	return uc.class
}

// Alloc allocates memory for a new instance of the class.
func (uc UNMutableNotificationContentClass) Alloc() UNMutableNotificationContent {
	rv := objc.Send[UNMutableNotificationContent](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// The editable content for a notification.
//
// # Overview
//
// Create a [UNMutableNotificationContent] object when you want to specify the
// payload for a local notification. Specifically, use this object to specify
// the title and message for an alert, the sound to play, or the value to
// assign to your app’s badge. You might also provide details about how the
// system handles the notification. For example, you can specify a custom
// launch image and a thread identifier for visually grouping related
// notifications.
//
// After creating your content object, assign it to a [UNNotificationRequest]
// object, add a trigger condition, and schedule your notification. The
// trigger condition defines when the system delivers the notification to the
// user. Listing 1 shows the scheduling of a local notification that displays
// an alert and plays a sound after a delay of five seconds. Store the strings
// for the alert’s title and body in the app’s `Localizable.Strings()`
// file.
//
// Listing 1. Creating the content for a local notification
//
// # Localizing the Alert Strings
//
// Localize the strings you display in a notification alert for the current
// user. Although you can use the [NSLocalizedString] macros to load strings
// from your app’s resource files, a better option is to specify your string
// using the [localizedUserNotificationString(forKey:arguments:)] method of
// [NSString]. The [localizedUserNotificationString(forKey:arguments:)] method
// delays the loading of the localized string until the system delivers the
// notification. If the user changes the language setting before the system
// delivers a notification, the system updates the alert text to the user’s
// current language instead of the language in use when the system scheduled
// the notification.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent
//
// [NSLocalizedString]: https://developer.apple.com/documentation/Foundation/NSLocalizedString
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
// [localizedUserNotificationString(forKey:arguments:)]: https://developer.apple.com/documentation/Foundation/NSString/localizedUserNotificationString(forKey:arguments:)
type UNMutableNotificationContent struct {
	UNNotificationContent
}

// UNMutableNotificationContentFromID constructs a [UNMutableNotificationContent] from an objc.ID.
//
// The editable content for a notification.
func UNMutableNotificationContentFromID(id objc.ID) UNMutableNotificationContent {
	return UNMutableNotificationContent{UNNotificationContent: UNNotificationContentFromID(id)}
}

// NOTE: UNMutableNotificationContent adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [UNMutableNotificationContent] class.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNMutableNotificationContent
type IUNMutableNotificationContent interface {
	IUNNotificationContent
}

// Init initializes the instance.
func (u UNMutableNotificationContent) Init() UNMutableNotificationContent {
	rv := objc.Send[UNMutableNotificationContent](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u UNMutableNotificationContent) Autorelease() UNMutableNotificationContent {
	rv := objc.Send[UNMutableNotificationContent](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewUNMutableNotificationContent creates a new UNMutableNotificationContent instance.
func NewUNMutableNotificationContent() UNMutableNotificationContent {
	class := getUNMutableNotificationContentClass()
	rv := objc.Send[UNMutableNotificationContent](objc.ID(class.class), objc.Sel("new"))
	return rv
}
