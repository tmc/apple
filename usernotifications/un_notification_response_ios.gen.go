// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.
// +build ios

package usernotifications

import (
"github.com/tmc/apple/objc"
"github.com/tmc/apple/objectivec"
)

// The scene where the system reflects the user’s response to a
// notification.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationResponse/targetScene
func (u UNNotificationResponse) TargetScene() objectivec.IObject {
rv := objc.Send[objc.ID](u.ID, objc.Sel("targetScene"))
return objectivec.Object{ID: rv}
}

