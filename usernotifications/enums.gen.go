// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/UserNotifications/UNAlertStyle
type UNAlertStyle int

const (
	// UNAlertStyleAlert: Modal alerts.
	UNAlertStyleAlert UNAlertStyle = 2
	// UNAlertStyleBanner: Banner alerts.
	UNAlertStyleBanner UNAlertStyle = 1
	// UNAlertStyleNone: No alert.
	UNAlertStyleNone UNAlertStyle = 0
)

func (e UNAlertStyle) String() string {
	switch e {
	case UNAlertStyleAlert:
		return "UNAlertStyleAlert"
	case UNAlertStyleBanner:
		return "UNAlertStyleBanner"
	case UNAlertStyleNone:
		return "UNAlertStyleNone"
	default:
		return fmt.Sprintf("UNAlertStyle(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/UserNotifications/UNAuthorizationOptions
type UNAuthorizationOptions uint

const (
	// UNAuthorizationOptionAlert: The ability to display alerts.
	UNAuthorizationOptionAlert UNAuthorizationOptions = 4
	// UNAuthorizationOptionBadge: The ability to update the app’s badge.
	UNAuthorizationOptionBadge UNAuthorizationOptions = 1
	// UNAuthorizationOptionCarPlay: The ability to display notifications in a CarPlay environment.
	UNAuthorizationOptionCarPlay UNAuthorizationOptions = 8
	// UNAuthorizationOptionCriticalAlert: The ability to play sounds for critical alerts.
	UNAuthorizationOptionCriticalAlert UNAuthorizationOptions = 16
	// UNAuthorizationOptionProvidesAppNotificationSettings: An option indicating the system should display a button for in-app notification settings.
	UNAuthorizationOptionProvidesAppNotificationSettings UNAuthorizationOptions = 32
	// UNAuthorizationOptionProvisional: The ability to post noninterrupting notifications provisionally to the Notification Center.
	UNAuthorizationOptionProvisional UNAuthorizationOptions = 64
	// UNAuthorizationOptionSound: The ability to play sounds.
	UNAuthorizationOptionSound UNAuthorizationOptions = 2
	// Deprecated.
	UNAuthorizationOptionAnnouncement UNAuthorizationOptions = 128
	// Deprecated.
	UNAuthorizationOptionTimeSensitive UNAuthorizationOptions = 256
)

func (e UNAuthorizationOptions) String() string {
	switch e {
	case UNAuthorizationOptionAlert:
		return "UNAuthorizationOptionAlert"
	case UNAuthorizationOptionBadge:
		return "UNAuthorizationOptionBadge"
	case UNAuthorizationOptionCarPlay:
		return "UNAuthorizationOptionCarPlay"
	case UNAuthorizationOptionCriticalAlert:
		return "UNAuthorizationOptionCriticalAlert"
	case UNAuthorizationOptionProvidesAppNotificationSettings:
		return "UNAuthorizationOptionProvidesAppNotificationSettings"
	case UNAuthorizationOptionProvisional:
		return "UNAuthorizationOptionProvisional"
	case UNAuthorizationOptionSound:
		return "UNAuthorizationOptionSound"
	case UNAuthorizationOptionAnnouncement:
		return "UNAuthorizationOptionAnnouncement"
	case UNAuthorizationOptionTimeSensitive:
		return "UNAuthorizationOptionTimeSensitive"
	default:
		return fmt.Sprintf("UNAuthorizationOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/UserNotifications/UNAuthorizationStatus
type UNAuthorizationStatus int

const (
	// UNAuthorizationStatusAuthorized: The app is authorized to schedule or receive notifications.
	UNAuthorizationStatusAuthorized UNAuthorizationStatus = 2
	// UNAuthorizationStatusDenied: The app isn’t authorized to schedule or receive notifications.
	UNAuthorizationStatusDenied UNAuthorizationStatus = 1
	// UNAuthorizationStatusEphemeral: The app is authorized to schedule or receive notifications for a limited amount of time.
	UNAuthorizationStatusEphemeral UNAuthorizationStatus = 4
	// UNAuthorizationStatusNotDetermined: The user hasn’t yet made a choice about whether the app is allowed to schedule notifications.
	UNAuthorizationStatusNotDetermined UNAuthorizationStatus = 0
	// UNAuthorizationStatusProvisional: The application is provisionally authorized to post noninterruptive user notifications.
	UNAuthorizationStatusProvisional UNAuthorizationStatus = 3
)

func (e UNAuthorizationStatus) String() string {
	switch e {
	case UNAuthorizationStatusAuthorized:
		return "UNAuthorizationStatusAuthorized"
	case UNAuthorizationStatusDenied:
		return "UNAuthorizationStatusDenied"
	case UNAuthorizationStatusEphemeral:
		return "UNAuthorizationStatusEphemeral"
	case UNAuthorizationStatusNotDetermined:
		return "UNAuthorizationStatusNotDetermined"
	case UNAuthorizationStatusProvisional:
		return "UNAuthorizationStatusProvisional"
	default:
		return fmt.Sprintf("UNAuthorizationStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/UserNotifications/UNError/Code
type UNErrorCode int

const (
	// UNErrorCodeAttachmentCorrupt: The file for an attachment is corrupt.
	UNErrorCodeAttachmentCorrupt UNErrorCode = 105
	// UNErrorCodeAttachmentInvalidFileSize: An attachment is too large.
	UNErrorCodeAttachmentInvalidFileSize UNErrorCode = 102
	// UNErrorCodeAttachmentInvalidURL: The URL for an attachment was invalid.
	UNErrorCodeAttachmentInvalidURL UNErrorCode = 100
	// UNErrorCodeAttachmentMoveIntoDataStoreFailed: An error occurred when trying to move an attachment to the system data store.
	UNErrorCodeAttachmentMoveIntoDataStoreFailed UNErrorCode = 104
	// UNErrorCodeAttachmentNotInDataStore: The specified attachment isn’t in the system data store.
	UNErrorCodeAttachmentNotInDataStore UNErrorCode = 103
	// UNErrorCodeAttachmentUnrecognizedType: The file type of an attachment isn’t supported.
	UNErrorCodeAttachmentUnrecognizedType UNErrorCode = 101
	// UNErrorCodeNotificationInvalidNoContent: The notification has no user-facing content, but should.
	UNErrorCodeNotificationInvalidNoContent UNErrorCode = 1401
	// UNErrorCodeNotificationInvalidNoDate: The notification doesn’t have an associated date, but should.
	UNErrorCodeNotificationInvalidNoDate UNErrorCode = 1400
	// UNErrorCodeNotificationsNotAllowed: Notifications aren’t allowed.
	UNErrorCodeNotificationsNotAllowed UNErrorCode = 1
)

func (e UNErrorCode) String() string {
	switch e {
	case UNErrorCodeAttachmentCorrupt:
		return "UNErrorCodeAttachmentCorrupt"
	case UNErrorCodeAttachmentInvalidFileSize:
		return "UNErrorCodeAttachmentInvalidFileSize"
	case UNErrorCodeAttachmentInvalidURL:
		return "UNErrorCodeAttachmentInvalidURL"
	case UNErrorCodeAttachmentMoveIntoDataStoreFailed:
		return "UNErrorCodeAttachmentMoveIntoDataStoreFailed"
	case UNErrorCodeAttachmentNotInDataStore:
		return "UNErrorCodeAttachmentNotInDataStore"
	case UNErrorCodeAttachmentUnrecognizedType:
		return "UNErrorCodeAttachmentUnrecognizedType"
	case UNErrorCodeNotificationInvalidNoContent:
		return "UNErrorCodeNotificationInvalidNoContent"
	case UNErrorCodeNotificationInvalidNoDate:
		return "UNErrorCodeNotificationInvalidNoDate"
	case UNErrorCodeNotificationsNotAllowed:
		return "UNErrorCodeNotificationsNotAllowed"
	default:
		return fmt.Sprintf("UNErrorCode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationActionOptions
type UNNotificationActionOptions uint

const (
	// UNNotificationActionOptionAuthenticationRequired: The action can be performed only on an unlocked device.
	UNNotificationActionOptionAuthenticationRequired UNNotificationActionOptions = 1
	// UNNotificationActionOptionDestructive: The action performs a destructive task.
	UNNotificationActionOptionDestructive UNNotificationActionOptions = 2
	// UNNotificationActionOptionForeground: The action causes the app to launch in the foreground.
	UNNotificationActionOptionForeground UNNotificationActionOptions = 4
)

func (e UNNotificationActionOptions) String() string {
	switch e {
	case UNNotificationActionOptionAuthenticationRequired:
		return "UNNotificationActionOptionAuthenticationRequired"
	case UNNotificationActionOptionDestructive:
		return "UNNotificationActionOptionDestructive"
	case UNNotificationActionOptionForeground:
		return "UNNotificationActionOptionForeground"
	default:
		return fmt.Sprintf("UNNotificationActionOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationCategoryOptions
type UNNotificationCategoryOptions uint

const (
	// UNNotificationCategoryOptionAllowInCarPlay: Allow CarPlay to display notifications of this type.
	UNNotificationCategoryOptionAllowInCarPlay UNNotificationCategoryOptions = 2
	// UNNotificationCategoryOptionCustomDismissAction: Send dismiss actions to the [UNUserNotificationCenter] object’s delegate for handling.
	UNNotificationCategoryOptionCustomDismissAction UNNotificationCategoryOptions = 1
	// UNNotificationCategoryOptionHiddenPreviewsShowSubtitle: Show the notification’s subtitle, even if the user has disabled notification previews for the app.
	UNNotificationCategoryOptionHiddenPreviewsShowSubtitle UNNotificationCategoryOptions = 8
	// UNNotificationCategoryOptionHiddenPreviewsShowTitle: Show the notification’s title, even if the user has disabled notification previews for the app.
	UNNotificationCategoryOptionHiddenPreviewsShowTitle UNNotificationCategoryOptions = 4
	// Deprecated.
	UNNotificationCategoryOptionAllowAnnouncement UNNotificationCategoryOptions = 16
)

func (e UNNotificationCategoryOptions) String() string {
	switch e {
	case UNNotificationCategoryOptionAllowInCarPlay:
		return "UNNotificationCategoryOptionAllowInCarPlay"
	case UNNotificationCategoryOptionCustomDismissAction:
		return "UNNotificationCategoryOptionCustomDismissAction"
	case UNNotificationCategoryOptionHiddenPreviewsShowSubtitle:
		return "UNNotificationCategoryOptionHiddenPreviewsShowSubtitle"
	case UNNotificationCategoryOptionHiddenPreviewsShowTitle:
		return "UNNotificationCategoryOptionHiddenPreviewsShowTitle"
	case UNNotificationCategoryOptionAllowAnnouncement:
		return "UNNotificationCategoryOptionAllowAnnouncement"
	default:
		return fmt.Sprintf("UNNotificationCategoryOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationInterruptionLevel
type UNNotificationInterruptionLevel uint

const (
	// UNNotificationInterruptionLevelActive: The system presents the notification immediately, lights up the screen, and can play a sound.
	UNNotificationInterruptionLevelActive UNNotificationInterruptionLevel = 1
	// UNNotificationInterruptionLevelCritical: The system presents the notification immediately, lights up the screen, and bypasses the mute switch to play a sound.
	UNNotificationInterruptionLevelCritical UNNotificationInterruptionLevel = 3
	// UNNotificationInterruptionLevelPassive: The system adds the notification to the notification list without lighting up the screen or playing a sound.
	UNNotificationInterruptionLevelPassive UNNotificationInterruptionLevel = 0
	// UNNotificationInterruptionLevelTimeSensitive: The system presents the notification immediately, lights up the screen, can play a sound, and breaks through system notification controls.
	UNNotificationInterruptionLevelTimeSensitive UNNotificationInterruptionLevel = 2
)

func (e UNNotificationInterruptionLevel) String() string {
	switch e {
	case UNNotificationInterruptionLevelActive:
		return "UNNotificationInterruptionLevelActive"
	case UNNotificationInterruptionLevelCritical:
		return "UNNotificationInterruptionLevelCritical"
	case UNNotificationInterruptionLevelPassive:
		return "UNNotificationInterruptionLevelPassive"
	case UNNotificationInterruptionLevelTimeSensitive:
		return "UNNotificationInterruptionLevelTimeSensitive"
	default:
		return fmt.Sprintf("UNNotificationInterruptionLevel(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationPresentationOptions
type UNNotificationPresentationOptions uint

const (
	// UNNotificationPresentationOptionBadge: Apply the notification’s badge value to the app’s icon.
	UNNotificationPresentationOptionBadge UNNotificationPresentationOptions = 1
	// UNNotificationPresentationOptionBanner: Present the notification as a banner.
	UNNotificationPresentationOptionBanner UNNotificationPresentationOptions = 16
	// UNNotificationPresentationOptionList: Show the notification in Notification Center.
	UNNotificationPresentationOptionList UNNotificationPresentationOptions = 8
	// UNNotificationPresentationOptionSound: Play the sound associated with the notification.
	UNNotificationPresentationOptionSound UNNotificationPresentationOptions = 2
	// Deprecated.
	UNNotificationPresentationOptionAlert UNNotificationPresentationOptions = 4
)

func (e UNNotificationPresentationOptions) String() string {
	switch e {
	case UNNotificationPresentationOptionBadge:
		return "UNNotificationPresentationOptionBadge"
	case UNNotificationPresentationOptionBanner:
		return "UNNotificationPresentationOptionBanner"
	case UNNotificationPresentationOptionList:
		return "UNNotificationPresentationOptionList"
	case UNNotificationPresentationOptionSound:
		return "UNNotificationPresentationOptionSound"
	case UNNotificationPresentationOptionAlert:
		return "UNNotificationPresentationOptionAlert"
	default:
		return fmt.Sprintf("UNNotificationPresentationOptions(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationSetting
type UNNotificationSetting int

const (
	// UNNotificationSettingDisabled: The setting is disabled.
	UNNotificationSettingDisabled UNNotificationSetting = 1
	// UNNotificationSettingEnabled: The setting is enabled.
	UNNotificationSettingEnabled UNNotificationSetting = 2
	// UNNotificationSettingNotSupported: The setting is not available to your app.
	UNNotificationSettingNotSupported UNNotificationSetting = 0
)

func (e UNNotificationSetting) String() string {
	switch e {
	case UNNotificationSettingDisabled:
		return "UNNotificationSettingDisabled"
	case UNNotificationSettingEnabled:
		return "UNNotificationSettingEnabled"
	case UNNotificationSettingNotSupported:
		return "UNNotificationSettingNotSupported"
	default:
		return fmt.Sprintf("UNNotificationSetting(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/UserNotifications/UNShowPreviewsSetting
type UNShowPreviewsSetting int

const (
	// UNShowPreviewsSettingAlways: The notification’s content is always shown, even when the device is locked.
	UNShowPreviewsSettingAlways UNShowPreviewsSetting = 0
	// UNShowPreviewsSettingNever: The notification’s content is never shown, even when the device is unlocked
	UNShowPreviewsSettingNever UNShowPreviewsSetting = 2
	// UNShowPreviewsSettingWhenAuthenticated: The notification’s content is shown only when the device is unlocked.
	UNShowPreviewsSettingWhenAuthenticated UNShowPreviewsSetting = 1
)

func (e UNShowPreviewsSetting) String() string {
	switch e {
	case UNShowPreviewsSettingAlways:
		return "UNShowPreviewsSettingAlways"
	case UNShowPreviewsSettingNever:
		return "UNShowPreviewsSettingNever"
	case UNShowPreviewsSettingWhenAuthenticated:
		return "UNShowPreviewsSettingWhenAuthenticated"
	default:
		return fmt.Sprintf("UNShowPreviewsSetting(%d)", e)
	}
}

