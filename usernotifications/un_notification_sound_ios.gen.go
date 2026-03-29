// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.
// +build ios

package usernotifications

import (
"github.com/tmc/apple/objc"
)

//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationSound/ringtoneSoundNamed(_:)
func (_UNNotificationSoundClass UNNotificationSoundClass) RingtoneSoundNamed(name UNNotificationSoundName) UNNotificationSound {
rv := objc.Send[objc.ID](objc.ID(_UNNotificationSoundClass.class), objc.Sel("ringtoneSoundNamed:"), objc.String(string(name)))
return UNNotificationSoundFromID(rv)
}

// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationSound/defaultRingtone
func (_UNNotificationSoundClass UNNotificationSoundClass) DefaultRingtoneSound() UNNotificationSound {
rv := objc.Send[objc.ID](objc.ID(_UNNotificationSoundClass.class), objc.Sel("defaultRingtoneSound"))
return UNNotificationSoundFromID(rv)
}

