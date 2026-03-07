// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// An interface that enables customizing the behavior of the default notification center.
//
// See: https://developer.apple.com/documentation/Foundation/NSUserNotificationCenterDelegate
type NSUserNotificationCenterDelegate interface {
	objectivec.IObject
}



// NSUserNotificationCenterDelegateObject wraps an existing Objective-C object that conforms to the NSUserNotificationCenterDelegate protocol.
type NSUserNotificationCenterDelegateObject struct {
	objectivec.Object
}
func (o NSUserNotificationCenterDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}



// NSUserNotificationCenterDelegateObjectFromID constructs a [NSUserNotificationCenterDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSUserNotificationCenterDelegateObjectFromID(id objc.ID) NSUserNotificationCenterDelegateObject {
	return NSUserNotificationCenterDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}










