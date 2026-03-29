// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.
// +build ios

package usernotifications

import (
"github.com/tmc/apple/objc"
"github.com/tmc/apple/foundation"
)

// The name of the image or storyboard to use when your app launches because
// of the notification.
//
// # Discussion
// 
// If you specify a value for this property, the system displays the specified
// image or storyboard when the system launches your app. The string in this
// property must match the name of an image file or storyboard in your app’s
// bundle.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationContent/launchImageName
func (u UNNotificationContent) LaunchImageName() string {
rv := objc.Send[objc.ID](u.ID, objc.Sel("launchImageName"))
return foundation.NSStringFromID(rv).String()
}

