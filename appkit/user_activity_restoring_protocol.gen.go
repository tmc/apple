// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol that marks classes to restore the state of your app to continue a user activity.
//
// See: https://developer.apple.com/documentation/AppKit/NSUserActivityRestoring
type NSUserActivityRestoring interface {
	objectivec.IObject

	// Restores the state necessary to continue the specified user activity.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSUserActivityRestoring/restoreUserActivityState(_:)
	RestoreUserActivityState(userActivity foundation.NSUserActivity)
}

// NSUserActivityRestoringObject wraps an existing Objective-C object that conforms to the NSUserActivityRestoring protocol.
type NSUserActivityRestoringObject struct {
	objectivec.Object
}

func (o NSUserActivityRestoringObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSUserActivityRestoringObjectFromID constructs a [NSUserActivityRestoringObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSUserActivityRestoringObjectFromID(id objc.ID) NSUserActivityRestoringObject {
	return NSUserActivityRestoringObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Restores the state necessary to continue the specified user activity.
//
// userActivity: The user activity to continue.
//
// # Discussion
//
// Implement this method to restore an object’s state using the specified
// user activity. The system calls this method on any responders or documents
// passed to the `restorationHandler` in
// [ApplicationContinueUserActivityRestorationHandler]. The system calls this
// method on the main thread. Your implementation should use the state data
// contained in the specified user activity’s [userInfo] dictionary to
// restore the object.
//
// On macOS, the system can automatically restore activities managed by
// [NSDocument] if you don’t implement
// [ApplicationContinueUserActivityRestorationHandler], or if you return
// false. When this occurs, the system opens the document using
// [OpenDocumentWithContentsOfURLDisplayCompletionHandler], and calls
// `restoreUserActivityState` on it.
//
// See: https://developer.apple.com/documentation/AppKit/NSUserActivityRestoring/restoreUserActivityState(_:)
//
// [userInfo]: https://developer.apple.com/documentation/Foundation/NSUserActivity/userInfo
func (o NSUserActivityRestoringObject) RestoreUserActivityState(userActivity foundation.NSUserActivity) {
	objc.Send[struct{}](o.ID, objc.Sel("restoreUserActivityState:"), userActivity)
}
