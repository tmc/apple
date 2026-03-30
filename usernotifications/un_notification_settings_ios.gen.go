// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.
//go:build ios
// +build ios

package usernotifications

import (
	"github.com/tmc/apple/objc"
)

// The setting that indicates whether your app’s notifications appear in
// CarPlay.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationSettings/carPlaySetting
func (u UNNotificationSettings) CarPlaySetting() UNNotificationSetting {
	rv := objc.Send[UNNotificationSetting](u.ID, objc.Sel("carPlaySetting"))
	return UNNotificationSetting(rv)
}

// The setting that indicates whether Siri can announce your app’s
// notifications.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationSettings/announcementSetting
func (u UNNotificationSettings) AnnouncementSetting() UNNotificationSetting {
	rv := objc.Send[UNNotificationSetting](u.ID, objc.Sel("announcementSetting"))
	return UNNotificationSetting(rv)
}
