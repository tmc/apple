// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [UNNotificationSettings] class.
var (
	_UNNotificationSettingsClass     UNNotificationSettingsClass
	_UNNotificationSettingsClassOnce sync.Once
)

func getUNNotificationSettingsClass() UNNotificationSettingsClass {
	_UNNotificationSettingsClassOnce.Do(func() {
		_UNNotificationSettingsClass = UNNotificationSettingsClass{class: objc.GetClass("UNNotificationSettings")}
	})
	return _UNNotificationSettingsClass
}

// GetUNNotificationSettingsClass returns the class object for UNNotificationSettings.
func GetUNNotificationSettingsClass() UNNotificationSettingsClass {
	return getUNNotificationSettingsClass()
}

type UNNotificationSettingsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (uc UNNotificationSettingsClass) Class() objc.Class {
	return uc.class
}

// Alloc allocates memory for a new instance of the class.
func (uc UNNotificationSettingsClass) Alloc() UNNotificationSettings {
	rv := objc.Send[UNNotificationSettings](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// The object for managing notification-related settings and the authorization
// status of your app.
//
// # Overview
//
// A [UNNotificationSettings] object contains the current authorization status
// and notification-related settings for your app. Apps must receive
// authorization to schedule notifications and to interact with the user. Apps
// that run in CarPlay must similarly receive authorization to do so. Use this
// object to determine what notification-related actions your app can perform.
// You might then use that information to enable, disable, or adjust your
// app’s notification-related behaviors. Regardless of whether you take
// action, the system enforces your app’s settings by preventing denied
// interactions from occurring.
//
// You don’t create instances of this class directly. Instead, call the
// [GetNotificationSettingsWithCompletionHandler] method of your app’s
// [UNUserNotificationCenter] object to get the current settings.
//
// For more information about requesting authorization for user interactions,
// see [UNUserNotificationCenter].
//
// # Getting the Authorization Status
//
//   - [UNNotificationSettings.AuthorizationStatus]: The app’s ability to schedule and receive local and remote notifications.
//
// # Getting Device-Specific Settings
//
//   - [UNNotificationSettings.NotificationCenterSetting]: The setting that indicates whether your app’s notifications appear in Notification Center.
//   - [UNNotificationSettings.LockScreenSetting]: The setting that indicates whether your app’s notifications appear on a device’s Lock screen.
//   - [UNNotificationSettings.AlertSetting]: The authorization status for displaying alerts.
//   - [UNNotificationSettings.BadgeSetting]: The setting that indicates whether badges appear on your app’s icon.
//   - [UNNotificationSettings.SoundSetting]: The authorization status for playing sounds for incoming notifications.
//   - [UNNotificationSettings.CriticalAlertSetting]: The authorization status for playing sounds for critical alerts.
//   - [UNNotificationSettings.ScheduledDeliverySetting]: The setting that indicates the system schedules the notification.
//   - [UNNotificationSettings.TimeSensitiveSetting]: The setting that indicates the system treats the notification as time-sensitive.
//
// # Getting Interface Settings
//
//   - [UNNotificationSettings.AlertStyle]: The type of alert that the app may display when the device is unlocked.
//   - [UNNotificationSettings.ShowPreviewsSetting]: The setting that indicates whether the app shows a preview of the notification’s content.
//   - [UNNotificationSettings.ProvidesAppNotificationSettings]: A Boolean value indicating the system displays a button for in-app notification settings.
//
// # Instance Properties
//
//   - [UNNotificationSettings.DirectMessagesSetting]
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationSettings
type UNNotificationSettings struct {
	objectivec.Object
}

// UNNotificationSettingsFromID constructs a [UNNotificationSettings] from an objc.ID.
//
// The object for managing notification-related settings and the authorization
// status of your app.
func UNNotificationSettingsFromID(id objc.ID) UNNotificationSettings {
	return UNNotificationSettings{objectivec.Object{ID: id}}
}

// NOTE: UNNotificationSettings adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [UNNotificationSettings] class.
//
// # Getting the Authorization Status
//
//   - [IUNNotificationSettings.AuthorizationStatus]: The app’s ability to schedule and receive local and remote notifications.
//
// # Getting Device-Specific Settings
//
//   - [IUNNotificationSettings.NotificationCenterSetting]: The setting that indicates whether your app’s notifications appear in Notification Center.
//   - [IUNNotificationSettings.LockScreenSetting]: The setting that indicates whether your app’s notifications appear on a device’s Lock screen.
//   - [IUNNotificationSettings.AlertSetting]: The authorization status for displaying alerts.
//   - [IUNNotificationSettings.BadgeSetting]: The setting that indicates whether badges appear on your app’s icon.
//   - [IUNNotificationSettings.SoundSetting]: The authorization status for playing sounds for incoming notifications.
//   - [IUNNotificationSettings.CriticalAlertSetting]: The authorization status for playing sounds for critical alerts.
//   - [IUNNotificationSettings.ScheduledDeliverySetting]: The setting that indicates the system schedules the notification.
//   - [IUNNotificationSettings.TimeSensitiveSetting]: The setting that indicates the system treats the notification as time-sensitive.
//
// # Getting Interface Settings
//
//   - [IUNNotificationSettings.AlertStyle]: The type of alert that the app may display when the device is unlocked.
//   - [IUNNotificationSettings.ShowPreviewsSetting]: The setting that indicates whether the app shows a preview of the notification’s content.
//   - [IUNNotificationSettings.ProvidesAppNotificationSettings]: A Boolean value indicating the system displays a button for in-app notification settings.
//
// # Instance Properties
//
//   - [IUNNotificationSettings.DirectMessagesSetting]
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationSettings
type IUNNotificationSettings interface {
	objectivec.IObject

	// Topic: Getting the Authorization Status

	// The app’s ability to schedule and receive local and remote notifications.
	AuthorizationStatus() UNAuthorizationStatus

	// Topic: Getting Device-Specific Settings

	// The setting that indicates whether your app’s notifications appear in Notification Center.
	NotificationCenterSetting() UNNotificationSetting
	// The setting that indicates whether your app’s notifications appear on a device’s Lock screen.
	LockScreenSetting() UNNotificationSetting
	// The authorization status for displaying alerts.
	AlertSetting() UNNotificationSetting
	// The setting that indicates whether badges appear on your app’s icon.
	BadgeSetting() UNNotificationSetting
	// The authorization status for playing sounds for incoming notifications.
	SoundSetting() UNNotificationSetting
	// The authorization status for playing sounds for critical alerts.
	CriticalAlertSetting() UNNotificationSetting
	// The setting that indicates the system schedules the notification.
	ScheduledDeliverySetting() UNNotificationSetting
	// The setting that indicates the system treats the notification as time-sensitive.
	TimeSensitiveSetting() UNNotificationSetting

	// Topic: Getting Interface Settings

	// The type of alert that the app may display when the device is unlocked.
	AlertStyle() UNAlertStyle
	// The setting that indicates whether the app shows a preview of the notification’s content.
	ShowPreviewsSetting() UNShowPreviewsSetting
	// A Boolean value indicating the system displays a button for in-app notification settings.
	ProvidesAppNotificationSettings() bool

	// Topic: Instance Properties

	DirectMessagesSetting() UNNotificationSetting

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (u UNNotificationSettings) Init() UNNotificationSettings {
	rv := objc.Send[UNNotificationSettings](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u UNNotificationSettings) Autorelease() UNNotificationSettings {
	rv := objc.Send[UNNotificationSettings](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewUNNotificationSettings creates a new UNNotificationSettings instance.
func NewUNNotificationSettings() UNNotificationSettings {
	class := getUNNotificationSettingsClass()
	rv := objc.Send[UNNotificationSettings](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (u UNNotificationSettings) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](u.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The app’s ability to schedule and receive local and remote notifications.
//
// # Discussion
//
// When the value of this property is [UNAuthorizationStatusAuthorized], your
// app is allowed to schedule and receive local and remote notifications. When
// authorized, use the [AlertSetting], [BadgeSetting], and [SoundSetting]
// properties to specify which types of interactions are allowed. When the
// value of the property is [UNAuthorizationStatusDenied], the system
// doesn’t deliver notifications to your app, and the system ignores any
// attempts to schedule local notifications.
//
// The value of this property is [UNAuthorizationStatusNotDetermined] if your
// app has never requested authorization using the
// [RequestAuthorizationWithOptionsCompletionHandler] method.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationSettings/authorizationStatus
func (u UNNotificationSettings) AuthorizationStatus() UNAuthorizationStatus {
	rv := objc.Send[UNAuthorizationStatus](u.ID, objc.Sel("authorizationStatus"))
	return UNAuthorizationStatus(rv)
}

// The setting that indicates whether your app’s notifications appear in
// Notification Center.
//
// # Discussion
//
// The default value of this property is [UNNotificationSettingEnabled].
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationSettings/notificationCenterSetting
func (u UNNotificationSettings) NotificationCenterSetting() UNNotificationSetting {
	rv := objc.Send[UNNotificationSetting](u.ID, objc.Sel("notificationCenterSetting"))
	return UNNotificationSetting(rv)
}

// The setting that indicates whether your app’s notifications appear on a
// device’s Lock screen.
//
// # Discussion
//
// Even if the user disables lock screen notifications, your notifications may
// still appear onscreen when the device is unlocked.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationSettings/lockScreenSetting
func (u UNNotificationSettings) LockScreenSetting() UNNotificationSetting {
	rv := objc.Send[UNNotificationSetting](u.ID, objc.Sel("lockScreenSetting"))
	return UNNotificationSetting(rv)
}

// The authorization status for displaying alerts.
//
// # Discussion
//
// When the value of this property is [UNNotificationSettingEnabled], the app
// is authorized to display alerts. Authorization does not guarantee that
// alerts always appear on the user’s screen. When a device is unlocked, the
// [AlertStyle] property determines the presentation style for the alert,
// which can include not displaying the alert at all.
//
// The system tries to display an alert when the [Title], [Subtitle], or
// [Body] properties of a [UNNotificationContent] object contain values, or
// when the `aps` dictionary in a remote notification contains the `alert`
// key.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationSettings/alertSetting
func (u UNNotificationSettings) AlertSetting() UNNotificationSetting {
	rv := objc.Send[UNNotificationSetting](u.ID, objc.Sel("alertSetting"))
	return UNNotificationSetting(rv)
}

// The setting that indicates whether badges appear on your app’s icon.
//
// # Discussion
//
// When the value of this property is [UNNotificationSettingEnabled], the app
// is authorized to badge its icon. The system tries to badge your app’s
// icon when the [Badge] property of a [UNNotificationContent] object contain
// a value, or when the `aps` dictionary in a remote notification contains the
// `badge` key.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationSettings/badgeSetting
func (u UNNotificationSettings) BadgeSetting() UNNotificationSetting {
	rv := objc.Send[UNNotificationSetting](u.ID, objc.Sel("badgeSetting"))
	return UNNotificationSetting(rv)
}

// The authorization status for playing sounds for incoming notifications.
//
// # Discussion
//
// When the value of this property is [UNNotificationSettingEnabled], the
// system authorizes the app to play sounds. The system tries to play a sound
// when the [Sound] property of the [UNNotificationContent] object contains a
// value, or when the `aps` dictionary in a remote notification contains the
// `sound` key.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationSettings/soundSetting
func (u UNNotificationSettings) SoundSetting() UNNotificationSetting {
	rv := objc.Send[UNNotificationSetting](u.ID, objc.Sel("soundSetting"))
	return UNNotificationSetting(rv)
}

// The authorization status for playing sounds for critical alerts.
//
// # Discussion
//
// When [UNNotificationSettingEnabled], this property authorizes the app to
// play critical sounds that ignore Do Not Disturb and the device’s mute
// switch.
//
// For local notifications, the system attempts to play a critical sound when
// the [Sound] property of the [UNNotificationContent] object contains an
// object returned by the [DefaultCriticalSound] property, the
// [CriticalSoundNamed] method, or a related method.
//
// For remote notifications, the system attempts to play a critical sound when
// the notification’s payload contains a `sound` directory that contains the
// `critical` key.
//
// Critical alerts require a special entitlement issued by Apple.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationSettings/criticalAlertSetting
func (u UNNotificationSettings) CriticalAlertSetting() UNNotificationSetting {
	rv := objc.Send[UNNotificationSetting](u.ID, objc.Sel("criticalAlertSetting"))
	return UNNotificationSetting(rv)
}

// The setting that indicates the system schedules the notification.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationSettings/scheduledDeliverySetting
func (u UNNotificationSettings) ScheduledDeliverySetting() UNNotificationSetting {
	rv := objc.Send[UNNotificationSetting](u.ID, objc.Sel("scheduledDeliverySetting"))
	return UNNotificationSetting(rv)
}

// The setting that indicates the system treats the notification as
// time-sensitive.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationSettings/timeSensitiveSetting
func (u UNNotificationSettings) TimeSensitiveSetting() UNNotificationSetting {
	rv := objc.Send[UNNotificationSetting](u.ID, objc.Sel("timeSensitiveSetting"))
	return UNNotificationSetting(rv)
}

// The type of alert that the app may display when the device is unlocked.
//
// # Discussion
//
// When alerts are authorized, this property specifies the presentation style
// for alerts when the device is unlocked. The user may choose to display
// alerts as automatically disappearing banners or as modal windows that
// require explicit dismissal. The user may also choose not to display alerts
// at all.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationSettings/alertStyle
func (u UNNotificationSettings) AlertStyle() UNAlertStyle {
	rv := objc.Send[UNAlertStyle](u.ID, objc.Sel("alertStyle"))
	return UNAlertStyle(rv)
}

// The setting that indicates whether the app shows a preview of the
// notification’s content.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationSettings/showPreviewsSetting
func (u UNNotificationSettings) ShowPreviewsSetting() UNShowPreviewsSetting {
	rv := objc.Send[UNShowPreviewsSetting](u.ID, objc.Sel("showPreviewsSetting"))
	return UNShowPreviewsSetting(rv)
}

// A Boolean value indicating the system displays a button for in-app
// notification settings.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationSettings/providesAppNotificationSettings
func (u UNNotificationSettings) ProvidesAppNotificationSettings() bool {
	rv := objc.Send[bool](u.ID, objc.Sel("providesAppNotificationSettings"))
	return rv
}

// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationSettings/directMessagesSetting
func (u UNNotificationSettings) DirectMessagesSetting() UNNotificationSetting {
	rv := objc.Send[UNNotificationSetting](u.ID, objc.Sel("directMessagesSetting"))
	return UNNotificationSetting(rv)
}
