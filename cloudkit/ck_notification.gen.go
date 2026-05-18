// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKNotification] class.
var (
	_CKNotificationClass     CKNotificationClass
	_CKNotificationClassOnce sync.Once
)

func getCKNotificationClass() CKNotificationClass {
	_CKNotificationClassOnce.Do(func() {
		_CKNotificationClass = CKNotificationClass{class: objc.GetClass("CKNotification")}
	})
	return _CKNotificationClass
}

// GetCKNotificationClass returns the class object for CKNotification.
func GetCKNotificationClass() CKNotificationClass {
	return getCKNotificationClass()
}

type CKNotificationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKNotificationClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKNotificationClass) Alloc() CKNotification {
	rv := objc.Send[CKNotification](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// The abstract base class for CloudKit notifications.
//
// # Overview
//
// Use subclasses of [CKNotification] to extract data from push notifications
// that the system receives, or to fetch a container’s previous push
// notifications. In both cases, the object indicates the changed data.
//
// [CKNotification] is an abstract class. When you create a notification from
// a payload dictionary, the [CKNotification.NotificationFromRemoteNotificationDictionary]
// method returns an instance of the appropriate subclass. Similarly, when you
// fetch notifications from a container, you receive instances of a concrete
// subclass. [CKNotification] provides information about the push notification
// and its method of delivery. Subclasses contain specific data that provides
// the changes.
//
// The system delivers notifications with alerts, badges, or sounds via the
// [UserNotifications] framework, in the form of a [UNNotification].
//
// Applications should use the [UserNotifications] framework to interact with
// the alert, badge, and sound properties of the notification.
//
// Applications may create a [CKNotification] from a [UNNotification] in their
// [UNUserNotificationCenterDelegate]:
//
// Notifications without alerts, badges, or sounds are delivered via an
// application delegate, in the form of a remote notification.
//
// For example: `UIApplicationDelegate.Application(_:) async`
//
// Applications may create a [CKNotification] from the remote notification in
// their [UIApplicationDelegate]:
//
// # Identifying the Notification
//
//   - [CKNotification.NotificationID]: The notification’s ID.
//   - [CKNotification.NotificationType]: The type of event that generates the notification.
//   - [CKNotification.ContainerIdentifier]: The ID of the container with the content that triggers the notification.
//
// # Getting the Notification’s Status
//
//   - [CKNotification.IsPruned]: A Boolean value that indicates whether the system removes some push notification content before delivery.
//
// # Accessing the Notification Info
//
//   - [CKNotification.AlertBody]: The notification’s alert body.
//   - [CKNotification.AlertLocalizationKey]: The key that identifies the localized text for the alert body.
//   - [CKNotification.AlertLocalizationArgs]: The fields for building a notification’s alert.
//   - [CKNotification.AlertActionLocalizationKey]: The key that identifies the localized string for the notification’s action.
//   - [CKNotification.AlertLaunchImage]: The filename of an image to use as a launch image.
//   - [CKNotification.SoundName]: The name of the sound file to play when a notification arrives.
//   - [CKNotification.Badge]: The value that the app icon’s badge displays.
//   - [CKNotification.Category]: The name of the action group that corresponds to this notification.
//   - [CKNotification.SubscriptionID]: The ID of the subscription that triggers the notification.
//   - [CKNotification.SetSubscriptionID]
//   - [CKNotification.SubscriptionOwnerUserRecordID]: The ID of the user record that creates the subscription that generates the push notification.
//   - [CKNotification.Title]: The notification’s title.
//   - [CKNotification.TitleLocalizationKey]: The key that identifies the localized string for the notification’s title.
//   - [CKNotification.TitleLocalizationArgs]: The fields for building a notification’s title.
//   - [CKNotification.Subtitle]: The notification’s subtitle.
//   - [CKNotification.SubtitleLocalizationKey]: The key that identifies the localized string for the notification’s subtitle.
//   - [CKNotification.SubtitleLocalizationArgs]: The fields for building a notification’s subtitle.
//
// See: https://developer.apple.com/documentation/CloudKit/CKNotification
type CKNotification struct {
	objectivec.Object
}

// CKNotificationFromID constructs a [CKNotification] from an objc.ID.
//
// The abstract base class for CloudKit notifications.
func CKNotificationFromID(id objc.ID) CKNotification {
	return CKNotification{objectivec.Object{ID: id}}
}

// NOTE: CKNotification adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKNotification] class.
//
// # Identifying the Notification
//
//   - [ICKNotification.NotificationID]: The notification’s ID.
//   - [ICKNotification.NotificationType]: The type of event that generates the notification.
//   - [ICKNotification.ContainerIdentifier]: The ID of the container with the content that triggers the notification.
//
// # Getting the Notification’s Status
//
//   - [ICKNotification.IsPruned]: A Boolean value that indicates whether the system removes some push notification content before delivery.
//
// # Accessing the Notification Info
//
//   - [ICKNotification.AlertBody]: The notification’s alert body.
//   - [ICKNotification.AlertLocalizationKey]: The key that identifies the localized text for the alert body.
//   - [ICKNotification.AlertLocalizationArgs]: The fields for building a notification’s alert.
//   - [ICKNotification.AlertActionLocalizationKey]: The key that identifies the localized string for the notification’s action.
//   - [ICKNotification.AlertLaunchImage]: The filename of an image to use as a launch image.
//   - [ICKNotification.SoundName]: The name of the sound file to play when a notification arrives.
//   - [ICKNotification.Badge]: The value that the app icon’s badge displays.
//   - [ICKNotification.Category]: The name of the action group that corresponds to this notification.
//   - [ICKNotification.SubscriptionID]: The ID of the subscription that triggers the notification.
//   - [ICKNotification.SetSubscriptionID]
//   - [ICKNotification.SubscriptionOwnerUserRecordID]: The ID of the user record that creates the subscription that generates the push notification.
//   - [ICKNotification.Title]: The notification’s title.
//   - [ICKNotification.TitleLocalizationKey]: The key that identifies the localized string for the notification’s title.
//   - [ICKNotification.TitleLocalizationArgs]: The fields for building a notification’s title.
//   - [ICKNotification.Subtitle]: The notification’s subtitle.
//   - [ICKNotification.SubtitleLocalizationKey]: The key that identifies the localized string for the notification’s subtitle.
//   - [ICKNotification.SubtitleLocalizationArgs]: The fields for building a notification’s subtitle.
//
// See: https://developer.apple.com/documentation/CloudKit/CKNotification
type ICKNotification interface {
	objectivec.IObject

	// Topic: Identifying the Notification

	// The notification’s ID.
	NotificationID() ICKNotificationID
	// The type of event that generates the notification.
	NotificationType() CKNotificationType
	// The ID of the container with the content that triggers the notification.
	ContainerIdentifier() string

	// Topic: Getting the Notification’s Status

	// A Boolean value that indicates whether the system removes some push notification content before delivery.
	IsPruned() bool

	// Topic: Accessing the Notification Info

	// The notification’s alert body.
	AlertBody() string
	// The key that identifies the localized text for the alert body.
	AlertLocalizationKey() string
	// The fields for building a notification’s alert.
	AlertLocalizationArgs() []string
	// The key that identifies the localized string for the notification’s action.
	AlertActionLocalizationKey() string
	// The filename of an image to use as a launch image.
	AlertLaunchImage() string
	// The name of the sound file to play when a notification arrives.
	SoundName() string
	// The value that the app icon’s badge displays.
	Badge() foundation.NSNumber
	// The name of the action group that corresponds to this notification.
	Category() string
	// The ID of the subscription that triggers the notification.
	SubscriptionID() string
	SetSubscriptionID(value string)
	// The ID of the user record that creates the subscription that generates the push notification.
	SubscriptionOwnerUserRecordID() ICKRecordID
	// The notification’s title.
	Title() string
	// The key that identifies the localized string for the notification’s title.
	TitleLocalizationKey() string
	// The fields for building a notification’s title.
	TitleLocalizationArgs() []string
	// The notification’s subtitle.
	Subtitle() string
	// The key that identifies the localized string for the notification’s subtitle.
	SubtitleLocalizationKey() string
	// The fields for building a notification’s subtitle.
	SubtitleLocalizationArgs() []string
}

// Init initializes the instance.
func (c CKNotification) Init() CKNotification {
	rv := objc.Send[CKNotification](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKNotification) Autorelease() CKNotification {
	rv := objc.Send[CKNotification](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKNotification creates a new CKNotification instance.
func NewCKNotification() CKNotification {
	class := getCKNotificationClass()
	rv := objc.Send[CKNotification](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new notification using the specified payload data.
//
// notificationDictionary: The push notification’s payload data. Use the dictionary that the system
// provides to your app delegate’s
// [application(_:didReceiveRemoteNotification:fetchCompletionHandler:)]
// method. This parameter must not be `nil`.
//
// See: https://developer.apple.com/documentation/CloudKit/CKNotification/init(fromRemoteNotificationDictionary:)
//
// [application(_:didReceiveRemoteNotification:fetchCompletionHandler:)]: https://developer.apple.com/documentation/UIKit/UIApplicationDelegate/application(_:didReceiveRemoteNotification:fetchCompletionHandler:)
func NewCKNotificationFromRemoteNotificationDictionary(notificationDictionary foundation.INSDictionary) CKNotification {
	rv := objc.Send[objc.ID](objc.ID(getCKNotificationClass().class), objc.Sel("notificationFromRemoteNotificationDictionary:"), notificationDictionary)
	return CKNotificationFromID(rv)
}

// The notification’s ID.
//
// # Discussion
//
// Use this property to differentiate notifications.
//
// See: https://developer.apple.com/documentation/CloudKit/CKNotification/notificationID
func (c CKNotification) NotificationID() ICKNotificationID {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("notificationID"))
	return CKNotificationIDFromID(objc.ID(rv))
}

// The type of event that generates the notification.
//
// # Discussion
//
// Different notification types correspond to different subclasses of
// [CKNotification], so you can use the value in this property to determine
// how to handle the notification data.
//
// See: https://developer.apple.com/documentation/CloudKit/CKNotification/notificationType-swift.property
func (c CKNotification) NotificationType() CKNotificationType {
	rv := objc.Send[CKNotificationType](c.ID, objc.Sel("notificationType"))
	return CKNotificationType(rv)
}

// The ID of the container with the content that triggers the notification.
//
// # Discussion
//
// Use this property to determine the location of the changed content.
//
// See: https://developer.apple.com/documentation/CloudKit/CKNotification/containerIdentifier
func (c CKNotification) ContainerIdentifier() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("containerIdentifier"))
	return foundation.NSStringFromID(rv).String()
}

// A Boolean value that indicates whether the system removes some push
// notification content before delivery.
//
// # Discussion
//
// The server may truncate the payload data of a push notification if the size
// of that data exceeds the allowed maximum. For notifications you create
// using a payload dictionary, the value of this property is true if the
// payload data doesn’t contain all information regarding the change. The
// value is false if the payload data is complete.
//
// For notifications you fetch from the database using a
// [CKFetchNotificationChangesOperation] operation, this property’s value is
// always true.
//
// When CloudKit must remove payload data, it removes it in a specific order.
// This class’s properties are among the last that CloudKit removes because
// they define information about how to deliver the push notification. The
// following list shows the properties that CloudKit removes, and the order
// for removing them:
//
// - [ContainerIdentifier] - Keys that subclasses of [CKNotification] define.
// - [SoundName] - [AlertLaunchImage] - [AlertActionLocalizationKey] -
// [AlertBody] - [AlertLocalizationArgs] - [AlertLocalizationKey] - [Badge] -
// [NotificationID]
//
// See: https://developer.apple.com/documentation/CloudKit/CKNotification/isPruned
func (c CKNotification) IsPruned() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isPruned"))
	return rv
}

// The notification’s alert body.
//
// # Discussion
//
// This property contains the nonlocalized text that the notification’s
// alert displays.
//
// See: https://developer.apple.com/documentation/CloudKit/CKNotification/alertBody
func (c CKNotification) AlertBody() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("alertBody"))
	return foundation.NSStringFromID(rv).String()
}

// The key that identifies the localized text for the alert body.
//
// # Discussion
//
// When the system delivers a push notification to your app, it gets the text
// for the alert body by looking up the specified key in your app’s
// `Localizable.Strings()` file. CloudKit ignores the value in [AlertBody] if
// you set this property.
//
// .
//
// See: https://developer.apple.com/documentation/CloudKit/CKNotification/alertLocalizationKey
func (c CKNotification) AlertLocalizationKey() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("alertLocalizationKey"))
	return foundation.NSStringFromID(rv).String()
}

// The fields for building a notification’s alert.
//
// # Discussion
//
// This property is an array of field names that CloudKit uses to extract the
// corresponding values from the record that triggers the push notification.
// The values are strings, numbers, or dates. CloudKit may truncate strings
// with a length greater than 100 characters when it adds them to a
// notification’s payload.
//
// If you use `%@` for your substitution variables, CloudKit replaces those
// variables by traversing the array in order. If you use variables of the
// form `%n$@`, where `n` is an integer, `n` represents the index (starting at
// 1) of the item in the array to use. So, the first item in the array
// replaces the variable `%1$@`, the second item replaces the variable `%2$@`,
// and so on. You can use indexed substitution variables to change the order
// of items in the resulting string, which might be necessary when you
// localize your app’s content.
//
// See: https://developer.apple.com/documentation/CloudKit/CKNotification/alertLocalizationArgs
func (c CKNotification) AlertLocalizationArgs() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("alertLocalizationArgs"))
	return objc.ConvertSliceToStrings(rv)
}

// The key that identifies the localized string for the notification’s
// action.
//
// # Discussion
//
// The system uses this property’s value to find the matching string in your
// app’s `Localizable.Strings()` file. It uses the string as the text of the
// button that opens your app, which the notification alert displays.
//
// If this property’s value is `nil`, the system displays a single button to
// dismiss the alert.
//
// See: https://developer.apple.com/documentation/CloudKit/CKNotification/alertActionLocalizationKey
func (c CKNotification) AlertActionLocalizationKey() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("alertActionLocalizationKey"))
	return foundation.NSStringFromID(rv).String()
}

// The filename of an image to use as a launch image.
//
// # Discussion
//
// The system uses this property’s value to locate an image in the app’s
// bundle, and displays it as a launch image when the user launches the app
// after receiving a push notification.
//
// See: https://developer.apple.com/documentation/CloudKit/CKNotification/alertLaunchImage
func (c CKNotification) AlertLaunchImage() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("alertLaunchImage"))
	return foundation.NSStringFromID(rv).String()
}

// The name of the sound file to play when a notification arrives.
//
// # Discussion
//
// The system uses this property’s value to locate a sound file in the
// app’s bundle. The sound plays when the system receives a push
// notification. If the system can’t find the specified file, or if the
// property’s value is the string `default`, the system plays the default
// sound.
//
// See: https://developer.apple.com/documentation/CloudKit/CKNotification/soundName
func (c CKNotification) SoundName() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("soundName"))
	return foundation.NSStringFromID(rv).String()
}

// The value that the app icon’s badge displays.
//
// See: https://developer.apple.com/documentation/CloudKit/CKNotification/badge
func (c CKNotification) Badge() foundation.NSNumber {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("badge"))
	return foundation.NSNumberFromID(objc.ID(rv))
}

// The name of the action group that corresponds to this notification.
//
// # Discussion
//
// Categories allow you to present custom actions to the user on your push
// notifications. For more information, see
// [UIMutableUserNotificationCategory].
//
// See: https://developer.apple.com/documentation/CloudKit/CKNotification/category
//
// [UIMutableUserNotificationCategory]: https://developer.apple.com/documentation/UIKit/UIMutableUserNotificationCategory
func (c CKNotification) Category() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("category"))
	return foundation.NSStringFromID(rv).String()
}

// The ID of the subscription that triggers the notification.
//
// See: https://developer.apple.com/documentation/cloudkit/cknotification/subscriptionid-16ygj
func (c CKNotification) SubscriptionID() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("subscriptionID"))
	return foundation.NSStringFromID(rv).String()
}
func (c CKNotification) SetSubscriptionID(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setSubscriptionID:"), objc.String(value))
}

// The ID of the user record that creates the subscription that generates the
// push notification.
//
// # Discussion
//
// On a system that supports multiple users, such as tvOS, use this identifier
// to check whether the pending content is for the current user. If your app
// always fetches data from CloudKit on launch, you may improve efficiency by
// disregarding notifications for other users.
//
// For more information about supporting a multiuser environment, see
// [Personalizing Your App for Each User on Apple TV].
//
// See: https://developer.apple.com/documentation/CloudKit/CKNotification/subscriptionOwnerUserRecordID
//
// [Personalizing Your App for Each User on Apple TV]: https://developer.apple.com/documentation/TVServices/personalizing-your-app-for-each-user-on-apple-tv
func (c CKNotification) SubscriptionOwnerUserRecordID() ICKRecordID {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("subscriptionOwnerUserRecordID"))
	return CKRecordIDFromID(objc.ID(rv))
}

// The notification’s title.
//
// # Discussion
//
// The system ignores this property if [TitleLocalizationKey] has a value.
//
// See: https://developer.apple.com/documentation/CloudKit/CKNotification/title
func (c CKNotification) Title() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("title"))
	return foundation.NSStringFromID(rv).String()
}

// The key that identifies the localized string for the notification’s
// title.
//
// # Discussion
//
// This property takes precedence over [Title].
//
// See: https://developer.apple.com/documentation/CloudKit/CKNotification/titleLocalizationKey
func (c CKNotification) TitleLocalizationKey() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("titleLocalizationKey"))
	return foundation.NSStringFromID(rv).String()
}

// The fields for building a notification’s title.
//
// # Discussion
//
// This property is an array of field names that CloudKit uses to extract the
// corresponding values from the record that triggers the push notification.
// The values are strings, numbers, or dates. CloudKit may truncate strings
// with a length greater than 100 characters when it adds them to a
// notification’s payload.
//
// If you use `%@` for your substitution variables, CloudKit replaces those
// variables by traversing the array in order. If you use variables of the
// form `%n$@`, where `n` is an integer, `n` represents the index (starting at
// 1) of the item in the array to use. So, the first item in the array
// replaces the variable `%1$@`, the second item replaces the variable `%2$@`,
// and so on. You can use indexed substitution variables to change the order
// of items in the resulting string, which might be necessary when you
// localize your app’s content.
//
// See: https://developer.apple.com/documentation/CloudKit/CKNotification/titleLocalizationArgs
func (c CKNotification) TitleLocalizationArgs() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("titleLocalizationArgs"))
	return objc.ConvertSliceToStrings(rv)
}

// The notification’s subtitle.
//
// # Discussion
//
// The system ignores this property if [SubtitleLocalizationKey] has a value.
//
// See: https://developer.apple.com/documentation/CloudKit/CKNotification/subtitle
func (c CKNotification) Subtitle() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("subtitle"))
	return foundation.NSStringFromID(rv).String()
}

// The key that identifies the localized string for the notification’s
// subtitle.
//
// # Discussion
//
// This property takes precedence over [Subtitle].
//
// See: https://developer.apple.com/documentation/CloudKit/CKNotification/subtitleLocalizationKey
func (c CKNotification) SubtitleLocalizationKey() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("subtitleLocalizationKey"))
	return foundation.NSStringFromID(rv).String()
}

// The fields for building a notification’s subtitle.
//
// # Discussion
//
// This property is an array of field names that CloudKit uses to extract the
// corresponding values from the record that triggers the push notification.
// The values are strings, numbers, or dates. CloudKit may truncate strings
// with a length greater than 100 characters when it adds them to a
// notification’s payload.
//
// If you use `%@` for your substitution variables, CloudKit replaces those
// variables by traversing the array in order. If you use variables of the
// form `%n$@`, where `n` is an integer, `n` represents the index (starting at
// 1) of the item in the array to use. So, the first item in the array
// replaces the variable `%1$@`, the second item replaces the variable `%2$@`,
// and so on. You can use indexed substitution variables to change the order
// of items in the resulting string, which might be necessary when you
// localize your app’s content.
//
// See: https://developer.apple.com/documentation/CloudKit/CKNotification/subtitleLocalizationArgs
func (c CKNotification) SubtitleLocalizationArgs() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("subtitleLocalizationArgs"))
	return objc.ConvertSliceToStrings(rv)
}
