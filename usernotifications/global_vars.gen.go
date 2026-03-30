// Code generated from Apple documentation. DO NOT EDIT.

package usernotifications

import (
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
)

// UNAuthorizationOptions values.
const (

	// UNAuthorizationOptionNone is no authorization options.
	//
	// See: https://developer.apple.com/documentation/UserNotifications/UNAuthorizationOptionNone
	UNAuthorizationOptionNone UNAuthorizationOptions = 0
)

// UNNotificationActionOptions values.
const (

	// UNNotificationActionOptionNone is the action has the default behavior.
	//
	// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationActionOptionNone
	UNNotificationActionOptionNone UNNotificationActionOptions = 0
)

// UNNotificationCategoryOptions values.
const (

	// UNNotificationCategoryOptionNone is no options.
	//
	// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationCategoryOptionNone
	UNNotificationCategoryOptionNone UNNotificationCategoryOptions = 0
)

// UNNotificationPresentationOptions values.
const (

	// UNNotificationPresentationOptionNone is no alert.
	//
	// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationPresentationOptionNone
	UNNotificationPresentationOptionNone UNNotificationPresentationOptions = 0
)

var (
	// UNErrorDomain is the error domain for notifications.
	//
	// See: https://developer.apple.com/documentation/UserNotifications/UNErrorDomain
	UNErrorDomain string
	// UNNotificationAttachmentOptionsThumbnailClippingRectKey is the clipping rectangle for a thumbnail image.
	//
	// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationAttachmentOptionsThumbnailClippingRectKey
	UNNotificationAttachmentOptionsThumbnailClippingRectKey string
	// UNNotificationAttachmentOptionsThumbnailHiddenKey is a Boolean value indicating whether the system hides the attachment’s thumbnail.
	//
	// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationAttachmentOptionsThumbnailHiddenKey
	UNNotificationAttachmentOptionsThumbnailHiddenKey string
	// UNNotificationAttachmentOptionsThumbnailTimeKey is the frame number of an animation to use as a thumbnail image.
	//
	// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationAttachmentOptionsThumbnailTimeKey
	UNNotificationAttachmentOptionsThumbnailTimeKey string
	// UNNotificationAttachmentOptionsTypeHintKey is a hint about an attachment’s file type.
	//
	// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationAttachmentOptionsTypeHintKey
	UNNotificationAttachmentOptionsTypeHintKey string
	// UNNotificationDefaultActionIdentifier is an action that indicates the user opened the app from the notification interface.
	//
	// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationDefaultActionIdentifier
	UNNotificationDefaultActionIdentifier string
	// UNNotificationDismissActionIdentifier is the action that indicates the user explicitly dismissed the notification interface.
	//
	// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationDismissActionIdentifier
	UNNotificationDismissActionIdentifier string
)

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UNErrorDomain"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				UNErrorDomain = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UNNotificationAttachmentOptionsThumbnailClippingRectKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				UNNotificationAttachmentOptionsThumbnailClippingRectKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UNNotificationAttachmentOptionsThumbnailHiddenKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				UNNotificationAttachmentOptionsThumbnailHiddenKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UNNotificationAttachmentOptionsThumbnailTimeKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				UNNotificationAttachmentOptionsThumbnailTimeKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UNNotificationAttachmentOptionsTypeHintKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				UNNotificationAttachmentOptionsTypeHintKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UNNotificationDefaultActionIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				UNNotificationDefaultActionIdentifier = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "UNNotificationDismissActionIdentifier"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				UNNotificationDismissActionIdentifier = objc.GoString(cstr)
			}
		}
	}

}
