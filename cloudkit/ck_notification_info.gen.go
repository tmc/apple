// Code generated from Apple documentation for CloudKit. DO NOT EDIT.

package cloudkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CKNotificationInfo] class.
var (
	_CKNotificationInfoClass     CKNotificationInfoClass
	_CKNotificationInfoClassOnce sync.Once
)

func getCKNotificationInfoClass() CKNotificationInfoClass {
	_CKNotificationInfoClassOnce.Do(func() {
		_CKNotificationInfoClass = CKNotificationInfoClass{class: objc.GetClass("CKNotificationInfo")}
	})
	return _CKNotificationInfoClass
}

// GetCKNotificationInfoClass returns the class object for CKNotificationInfo.
func GetCKNotificationInfoClass() CKNotificationInfoClass {
	return getCKNotificationInfoClass()
}

type CKNotificationInfoClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CKNotificationInfoClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CKNotificationInfoClass) Alloc() CKNotificationInfo {
	rv := objc.Send[CKNotificationInfo](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// An object that describes the configuration of a subscription’s push
// notifications.
//
// # Overview
//
// When configuring a subscription, use this class to specify the type of push
// notifications you want to generate when conditions meet the
// subscription’s trigger. You can provide content that the system displays
// to the user, describe the sounds to play, and indicate whether the app’s
// icon has a badge. You can request that the notification include information
// about the record that triggers it.
//
// When your app receives a push notification that a subscription generates,
// instantiate an instance of [CKNotification] using the
// [CKRecordZoneNotificationClass.NotificationFromRemoteNotificationDictionary]
// method and pass the notification’s payload. The object that the method
// returns contains the data you specify when configuring the subscription.
//
// For more information about push notification alerts and how they display to
// the user, see [Apple Push Notification Service] in [Local and Remote
// Notification Programming Guide].
//
// # Grouping Notifications
//
//   - [CKNotificationInfo.Category]: The name of the action group that corresponds to this notification.
//   - [CKNotificationInfo.SetCategory]
//   - [CKNotificationInfo.CollapseIDKey]: A value that the system uses to coalesce unseen push notifications.
//   - [CKNotificationInfo.SetCollapseIDKey]
//
// # Displaying Badges
//
//   - [CKNotificationInfo.ShouldBadge]: A Boolean value that determines whether an app’s icon badge increments its value.
//   - [CKNotificationInfo.SetShouldBadge]
//
// # Accessing the Notification Alert
//
//   - [CKNotificationInfo.AlertBody]: The text for the notification’s alert.
//   - [CKNotificationInfo.SetAlertBody]
//   - [CKNotificationInfo.AlertLocalizationKey]: The key that identifies the localized string for the notification’s alert.
//   - [CKNotificationInfo.SetAlertLocalizationKey]
//   - [CKNotificationInfo.AlertActionLocalizationKey]: The key that identifies the localized string for the notification’s action.
//   - [CKNotificationInfo.SetAlertActionLocalizationKey]
//   - [CKNotificationInfo.AlertLaunchImage]: The filename of an image to use as a launch image.
//   - [CKNotificationInfo.SetAlertLaunchImage]
//   - [CKNotificationInfo.SoundName]: The filename of the sound file to play when a notification arrives.
//   - [CKNotificationInfo.SetSoundName]
//
// # Accessing the Notification Info
//
//   - [CKNotificationInfo.ShouldSendContentAvailable]: A Boolean value that indicates whether the push notification includes the content available flag.
//   - [CKNotificationInfo.SetShouldSendContentAvailable]
//   - [CKNotificationInfo.ShouldSendMutableContent]: A Boolean value that indicates whether the push notification sets the mutable content flag.
//   - [CKNotificationInfo.SetShouldSendMutableContent]
//
// # Accessing the Notification Title
//
//   - [CKNotificationInfo.Title]: The notification’s title.
//   - [CKNotificationInfo.SetTitle]
//   - [CKNotificationInfo.TitleLocalizationKey]: The key that identifies the localized string for the notification’s title.
//   - [CKNotificationInfo.SetTitleLocalizationKey]
//
// # Accessing the Notification Subtitle
//
//   - [CKNotificationInfo.Subtitle]: The notification’s subtitle.
//   - [CKNotificationInfo.SetSubtitle]
//   - [CKNotificationInfo.SubtitleLocalizationKey]: The key that identifies the localized string for the notification’s subtitle.
//   - [CKNotificationInfo.SetSubtitleLocalizationKey]
//
// See: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class
//
// [Apple Push Notification Service]: https://developer.apple.com/library/archive/documentation/NetworkingInternet/Conceptual/RemoteNotificationsPG/Chapters/ApplePushService.html#//apple_ref/doc/uid/TP40008194-CH100
// [Local and Remote Notification Programming Guide]: https://developer.apple.com/library/archive/documentation/NetworkingInternet/Conceptual/RemoteNotificationsPG/index.html#//apple_ref/doc/uid/TP40008194
type CKNotificationInfo struct {
	objectivec.Object
}

// CKNotificationInfoFromID constructs a [CKNotificationInfo] from an objc.ID.
//
// An object that describes the configuration of a subscription’s push
// notifications.
func CKNotificationInfoFromID(id objc.ID) CKNotificationInfo {
	return CKNotificationInfo{objectivec.Object{ID: id}}
}

// NOTE: CKNotificationInfo adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CKNotificationInfo] class.
//
// # Grouping Notifications
//
//   - [ICKNotificationInfo.Category]: The name of the action group that corresponds to this notification.
//   - [ICKNotificationInfo.SetCategory]
//   - [ICKNotificationInfo.CollapseIDKey]: A value that the system uses to coalesce unseen push notifications.
//   - [ICKNotificationInfo.SetCollapseIDKey]
//
// # Displaying Badges
//
//   - [ICKNotificationInfo.ShouldBadge]: A Boolean value that determines whether an app’s icon badge increments its value.
//   - [ICKNotificationInfo.SetShouldBadge]
//
// # Accessing the Notification Alert
//
//   - [ICKNotificationInfo.AlertBody]: The text for the notification’s alert.
//   - [ICKNotificationInfo.SetAlertBody]
//   - [ICKNotificationInfo.AlertLocalizationKey]: The key that identifies the localized string for the notification’s alert.
//   - [ICKNotificationInfo.SetAlertLocalizationKey]
//   - [ICKNotificationInfo.AlertActionLocalizationKey]: The key that identifies the localized string for the notification’s action.
//   - [ICKNotificationInfo.SetAlertActionLocalizationKey]
//   - [ICKNotificationInfo.AlertLaunchImage]: The filename of an image to use as a launch image.
//   - [ICKNotificationInfo.SetAlertLaunchImage]
//   - [ICKNotificationInfo.SoundName]: The filename of the sound file to play when a notification arrives.
//   - [ICKNotificationInfo.SetSoundName]
//
// # Accessing the Notification Info
//
//   - [ICKNotificationInfo.ShouldSendContentAvailable]: A Boolean value that indicates whether the push notification includes the content available flag.
//   - [ICKNotificationInfo.SetShouldSendContentAvailable]
//   - [ICKNotificationInfo.ShouldSendMutableContent]: A Boolean value that indicates whether the push notification sets the mutable content flag.
//   - [ICKNotificationInfo.SetShouldSendMutableContent]
//
// # Accessing the Notification Title
//
//   - [ICKNotificationInfo.Title]: The notification’s title.
//   - [ICKNotificationInfo.SetTitle]
//   - [ICKNotificationInfo.TitleLocalizationKey]: The key that identifies the localized string for the notification’s title.
//   - [ICKNotificationInfo.SetTitleLocalizationKey]
//
// # Accessing the Notification Subtitle
//
//   - [ICKNotificationInfo.Subtitle]: The notification’s subtitle.
//   - [ICKNotificationInfo.SetSubtitle]
//   - [ICKNotificationInfo.SubtitleLocalizationKey]: The key that identifies the localized string for the notification’s subtitle.
//   - [ICKNotificationInfo.SetSubtitleLocalizationKey]
//
// See: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class
type ICKNotificationInfo interface {
	objectivec.IObject

	// Topic: Grouping Notifications

	// The name of the action group that corresponds to this notification.
	Category() string
	SetCategory(value string)
	// A value that the system uses to coalesce unseen push notifications.
	CollapseIDKey() string
	SetCollapseIDKey(value string)

	// Topic: Displaying Badges

	// A Boolean value that determines whether an app’s icon badge increments its value.
	ShouldBadge() bool
	SetShouldBadge(value bool)

	// Topic: Accessing the Notification Alert

	// The text for the notification’s alert.
	AlertBody() string
	SetAlertBody(value string)
	// The key that identifies the localized string for the notification’s alert.
	AlertLocalizationKey() string
	SetAlertLocalizationKey(value string)
	// The key that identifies the localized string for the notification’s action.
	AlertActionLocalizationKey() string
	SetAlertActionLocalizationKey(value string)
	// The filename of an image to use as a launch image.
	AlertLaunchImage() string
	SetAlertLaunchImage(value string)
	// The filename of the sound file to play when a notification arrives.
	SoundName() string
	SetSoundName(value string)

	// Topic: Accessing the Notification Info

	// A Boolean value that indicates whether the push notification includes the content available flag.
	ShouldSendContentAvailable() bool
	SetShouldSendContentAvailable(value bool)
	// A Boolean value that indicates whether the push notification sets the mutable content flag.
	ShouldSendMutableContent() bool
	SetShouldSendMutableContent(value bool)

	// Topic: Accessing the Notification Title

	// The notification’s title.
	Title() string
	SetTitle(value string)
	// The key that identifies the localized string for the notification’s title.
	TitleLocalizationKey() string
	SetTitleLocalizationKey(value string)

	// Topic: Accessing the Notification Subtitle

	// The notification’s subtitle.
	Subtitle() string
	SetSubtitle(value string)
	// The key that identifies the localized string for the notification’s subtitle.
	SubtitleLocalizationKey() string
	SetSubtitleLocalizationKey(value string)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (c CKNotificationInfo) Init() CKNotificationInfo {
	rv := objc.Send[CKNotificationInfo](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CKNotificationInfo) Autorelease() CKNotificationInfo {
	rv := objc.Send[CKNotificationInfo](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCKNotificationInfo creates a new CKNotificationInfo instance.
func NewCKNotificationInfo() CKNotificationInfo {
	class := getCKNotificationInfoClass()
	rv := objc.Send[CKNotificationInfo](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (c CKNotificationInfo) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The name of the action group that corresponds to this notification.
//
// # Discussion
//
// Categories allow you to present custom actions to the user on your push
// notifications. For more information, see
// [UIMutableUserNotificationCategory].
//
// See: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/category
//
// [UIMutableUserNotificationCategory]: https://developer.apple.com/documentation/UIKit/UIMutableUserNotificationCategory
func (c CKNotificationInfo) Category() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("category"))
	return foundation.NSStringFromID(rv).String()
}
func (c CKNotificationInfo) SetCategory(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setCategory:"), objc.String(value))
}

// A value that the system uses to coalesce unseen push notifications.
//
// # Discussion
//
// When CloudKit generates a push notification, it sets the notification’s
// `apns-collapse-id` header to this property’s value. The system uses this
// header to coalesce unseen notifications.
//
// See [Sending notification requests to APNs] for more information about
// sending notifications using the Apple Push Notification service.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/collapseIDKey
//
// [Sending notification requests to APNs]: https://developer.apple.com/documentation/UserNotifications/sending-notification-requests-to-apns
func (c CKNotificationInfo) CollapseIDKey() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("collapseIDKey"))
	return foundation.NSStringFromID(rv).String()
}
func (c CKNotificationInfo) SetCollapseIDKey(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setCollapseIDKey:"), objc.String(value))
}

// A Boolean value that determines whether an app’s icon badge increments
// its value.
//
// # Discussion
//
// The default value of this property is false. Set it to true to cause the
// system to increment the badge value whenever it receives the corresponding
// push notification.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/shouldBadge
func (c CKNotificationInfo) ShouldBadge() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("shouldBadge"))
	return rv
}
func (c CKNotificationInfo) SetShouldBadge(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setShouldBadge:"), value)
}

// The text for the notification’s alert.
//
// # Discussion
//
// Set this property’s value to have the system display the specified string
// when it receives the corresponding push notification. If you localize your
// app’s content, use the [CKNotificationInfo.AlertLocalizationKey] property
// instead.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/alertBody
func (c CKNotificationInfo) AlertBody() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("alertBody"))
	return foundation.NSStringFromID(rv).String()
}
func (c CKNotificationInfo) SetAlertBody(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setAlertBody:"), objc.String(value))
}

// The key that identifies the localized string for the notification’s
// alert.
//
// # Discussion
//
// Set this property’s value to have the system display a localized string
// when it receives the corresponding push notification. The system uses the
// key to find the matching string in your app’s `Localizable.String()`
// file. If you specify a value for this property, CloudKit ignores the
// [CKNotificationInfo.AlertBody] property’s value.
//
// For information about localizing string resources, see
// [Internationalization and Localization Guide].
//
// See: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/alertLocalizationKey
//
// [Internationalization and Localization Guide]: https://developer.apple.com/library/archive/documentation/MacOSX/Conceptual/BPInternational/Introduction/Introduction.html#//apple_ref/doc/uid/10000171i
func (c CKNotificationInfo) AlertLocalizationKey() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("alertLocalizationKey"))
	return foundation.NSStringFromID(rv).String()
}
func (c CKNotificationInfo) SetAlertLocalizationKey(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setAlertLocalizationKey:"), objc.String(value))
}

// The key that identifies the localized string for the notification’s
// action.
//
// # Discussion
//
// Set this property’s value to have the system use a localized string for
// the text of the notification’s button that opens your app. The system
// uses the key to find the matching string in your app’s
// `Localizable.String()` file.
//
// If this property’s value is `nil`, the system displays a single button to
// dismiss the alert.
//
// For information about localizing string resources, see
// [Internationalization and Localization Guide].
//
// See: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/alertActionLocalizationKey
//
// [Internationalization and Localization Guide]: https://developer.apple.com/library/archive/documentation/MacOSX/Conceptual/BPInternational/Introduction/Introduction.html#//apple_ref/doc/uid/10000171i
func (c CKNotificationInfo) AlertActionLocalizationKey() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("alertActionLocalizationKey"))
	return foundation.NSStringFromID(rv).String()
}
func (c CKNotificationInfo) SetAlertActionLocalizationKey(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setAlertActionLocalizationKey:"), objc.String(value))
}

// The filename of an image to use as a launch image.
//
// # Discussion
//
// If you specify a value, the system uses it to locate an image in the
// app’s bundle, and displays it as a launch image when the user launches
// the app after receiving a push notification.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/alertLaunchImage
func (c CKNotificationInfo) AlertLaunchImage() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("alertLaunchImage"))
	return foundation.NSStringFromID(rv).String()
}
func (c CKNotificationInfo) SetAlertLaunchImage(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setAlertLaunchImage:"), objc.String(value))
}

// The filename of the sound file to play when a notification arrives.
//
// # Discussion
//
// If you specify a value, the system uses it to locate a sound file in the
// app’s bundle. The sound plays when the system receives a push
// notification. If the system can’t find the specified file, or if you use
// the string `default`, the system plays the default sound.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/soundName
func (c CKNotificationInfo) SoundName() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("soundName"))
	return foundation.NSStringFromID(rv).String()
}
func (c CKNotificationInfo) SetSoundName(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setSoundName:"), objc.String(value))
}

// A Boolean value that indicates whether the push notification includes the
// content available flag.
//
// # Discussion
//
// When this property is true, the server includes the `content-available`
// flag in the push notification’s payload. That flag causes the system to
// wake or launch an app that isn’t currently running. The app then receives
// background execution time to download any data for the push notification,
// such as the set of changed records. If the app is already running in the
// foreground, the inclusion of this flag has no additional effect and the
// system delivers the notification to the app delegate for processing as
// usual.
//
// The default value of this property is false.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/shouldSendContentAvailable
func (c CKNotificationInfo) ShouldSendContentAvailable() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("shouldSendContentAvailable"))
	return rv
}
func (c CKNotificationInfo) SetShouldSendContentAvailable(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setShouldSendContentAvailable:"), value)
}

// A Boolean value that indicates whether the push notification sets the
// mutable content flag.
//
// # Discussion
//
// When this property is true, the server includes the `mutable-content` flag
// with a value of `1` in the push notification’s payload. When the value is
// `1`, the system passes the notification to your app extension for
// modification before delivery.
//
// See [Generating a remote notification] for more information about the
// `mutable-content` flag, and [Modifying content in newly delivered
// notifications] for information about how to modify push notifiction content
// in your app extension prior to delivery.
//
// The default value of this property is false.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/shouldSendMutableContent
//
// [Generating a remote notification]: https://developer.apple.com/documentation/UserNotifications/generating-a-remote-notification
// [Modifying content in newly delivered notifications]: https://developer.apple.com/documentation/UserNotifications/modifying-content-in-newly-delivered-notifications
func (c CKNotificationInfo) ShouldSendMutableContent() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("shouldSendMutableContent"))
	return rv
}
func (c CKNotificationInfo) SetShouldSendMutableContent(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setShouldSendMutableContent:"), value)
}

// The notification’s title.
//
// # Discussion
//
// CloudKit uses this value to set the `title` push notification property.
//
// See [Generating a remote notification] for more detail about push
// notification properties.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/title
//
// [Generating a remote notification]: https://developer.apple.com/documentation/UserNotifications/generating-a-remote-notification
func (c CKNotificationInfo) Title() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("title"))
	return foundation.NSStringFromID(rv).String()
}
func (c CKNotificationInfo) SetTitle(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setTitle:"), objc.String(value))
}

// The key that identifies the localized string for the notification’s
// title.
//
// # Discussion
//
// CloudKit uses this value to set the `title-loc-key` push notification
// property.
//
// See [Generating a remote notification] for more details about push
// notification properties.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/titleLocalizationKey
//
// [Generating a remote notification]: https://developer.apple.com/documentation/UserNotifications/generating-a-remote-notification
func (c CKNotificationInfo) TitleLocalizationKey() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("titleLocalizationKey"))
	return foundation.NSStringFromID(rv).String()
}
func (c CKNotificationInfo) SetTitleLocalizationKey(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setTitleLocalizationKey:"), objc.String(value))
}

// The notification’s subtitle.
//
// # Discussion
//
// CloudKit uses this value to set the `subtitle` push notification property.
// If you set [CKNotificationInfo.SubtitleLocalizationKey], CloudKit ignores
// this value.
//
// See [Generating a remote notification] for more details about push
// notification properties.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/subtitle
//
// [Generating a remote notification]: https://developer.apple.com/documentation/UserNotifications/generating-a-remote-notification
func (c CKNotificationInfo) Subtitle() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("subtitle"))
	return foundation.NSStringFromID(rv).String()
}
func (c CKNotificationInfo) SetSubtitle(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setSubtitle:"), objc.String(value))
}

// The key that identifies the localized string for the notification’s
// subtitle.
//
// # Discussion
//
// CloudKit uses this value to set the `subtitle-loc-key` push notification
// property. Setting this property overrides any value in
// [CKNotificationInfo.Subtitle].
//
// See [Generating a remote notification] for more details about push
// notification properties.
//
// See: https://developer.apple.com/documentation/CloudKit/CKSubscription/NotificationInfo-swift.class/subtitleLocalizationKey
//
// [Generating a remote notification]: https://developer.apple.com/documentation/UserNotifications/generating-a-remote-notification
func (c CKNotificationInfo) SubtitleLocalizationKey() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("subtitleLocalizationKey"))
	return foundation.NSStringFromID(rv).String()
}
func (c CKNotificationInfo) SetSubtitleLocalizationKey(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setSubtitleLocalizationKey:"), objc.String(value))
}
