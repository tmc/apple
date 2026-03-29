// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol the system uses to provide context relevant to user notifications.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationContentProviding
type UNNotificationContentProviding interface {
	objectivec.IObject
}

// UNNotificationContentProvidingObject wraps an existing Objective-C object that conforms to the UNNotificationContentProviding protocol.
type UNNotificationContentProvidingObject struct {
	objectivec.Object
}
func (o UNNotificationContentProvidingObject) BaseObject() objectivec.Object {
	return o.Object
}

// UNNotificationContentProvidingObjectFromID constructs a [UNNotificationContentProvidingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func UNNotificationContentProvidingObjectFromID(id objc.ID) UNNotificationContentProvidingObject {
	return UNNotificationContentProvidingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

