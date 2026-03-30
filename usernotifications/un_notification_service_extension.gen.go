// Code generated from Apple documentation for UserNotifications. DO NOT EDIT.

package usernotifications

import (
	"context"
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [UNNotificationServiceExtension] class.
var (
	_UNNotificationServiceExtensionClass     UNNotificationServiceExtensionClass
	_UNNotificationServiceExtensionClassOnce sync.Once
)

func getUNNotificationServiceExtensionClass() UNNotificationServiceExtensionClass {
	_UNNotificationServiceExtensionClassOnce.Do(func() {
		_UNNotificationServiceExtensionClass = UNNotificationServiceExtensionClass{class: objc.GetClass("UNNotificationServiceExtension")}
	})
	return _UNNotificationServiceExtensionClass
}

// GetUNNotificationServiceExtensionClass returns the class object for UNNotificationServiceExtension.
func GetUNNotificationServiceExtensionClass() UNNotificationServiceExtensionClass {
	return getUNNotificationServiceExtensionClass()
}

type UNNotificationServiceExtensionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (uc UNNotificationServiceExtensionClass) Class() objc.Class {
	return uc.class
}

// Alloc allocates memory for a new instance of the class.
func (uc UNNotificationServiceExtensionClass) Alloc() UNNotificationServiceExtension {
	rv := objc.Send[UNNotificationServiceExtension](objc.ID(uc.class), objc.Sel("alloc"))
	return rv
}

// An object that modifies the content of a remote notification before it’s
// delivered to the user.
//
// # Overview
//
// A [UNNotificationServiceExtension] object provides the entry point for a
// notification service app extension. This object lets you customize the
// content of a remote notification before the system delivers it to the user.
// A notification service app extension doesn’t present any UI of its own.
// Instead, it’s launched on demand when the system delivers a notification
// of the appropriate type to the user’s device. You use this extension to
// modify the notification’s content or download content related to the
// extension. For example, you could use the extension to decrypt an encrypted
// data block or to download images associated with the notification.
//
// You don’t create [UNNotificationServiceExtension] objects yourself.
// Instead, the Xcode template for a notification service extension target
// contains a subclass for you to modify. Use the methods of that subclass to
// implement your app extension’s behavior. When your app receives a remote
// notification for your app, the system loads your extension and calls its
// [UNNotificationServiceExtension.DidReceiveNotificationRequestWithContentHandler] method given the
// following conditions:
//
// - Your app has configured the remote notification to display an alert. -
// The remote notification’s `aps` dictionary includes the `mutable-content`
// key with the value set to `1`.
//
// The [UNNotificationServiceExtension.DidReceiveNotificationRequestWithContentHandler] method performs the
// main work of your extension. You use that method to make any changes to the
// notification’s content. That method has a limited amount of time to
// perform its task and execute the provided completion block. If your method
// doesn’t finish in time, the system calls the
// [UNNotificationServiceExtension.ServiceExtensionTimeWillExpire] method to give you one last chance to
// submit your changes. If you don’t update the notification content before
// time expires, the system displays the original content.
//
// As for any app extension, you deliver a notification service app extension
// class as a bundle inside your app. The template that Xcode provides
// configures the `Info.Plist()` file automatically for this app extension
// type. Specifically, it sets the value of the [NSExtensionPointIdentifier]
// key to `com.AppleXCUIElementTypeUsernotificationsXCUIElementTypeService()`
// and sets the value of the [NSExtensionPrincipalClass] key to the name of
// your [UNNotificationServiceExtension] subclass.
//
// For information about how to set up and send remote notifications, see
// [Setting up a remote notification server].
//
// # Subclassing Notes
//
// The Xcode templates provide a subclass of [UNNotificationServiceExtension]
// for you to modify. You must implement the
// [UNNotificationServiceExtension.DidReceiveNotificationRequestWithContentHandler] method and use it to
// process incoming notifications. It’s also strongly recommended that you
// override the [UNNotificationServiceExtension.ServiceExtensionTimeWillExpire] method.
//
// # Processing Notifications
//
//   - [UNNotificationServiceExtension.DidReceiveNotificationRequestWithContentHandler]: Asks you to make any needed changes to the notification and notify the system when you’re done.
//   - [UNNotificationServiceExtension.ServiceExtensionTimeWillExpire]: Tells you that the system is terminating your extension.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationServiceExtension
//
// [NSExtensionPointIdentifier]: https://developer.apple.com/library/archive/documentation/General/Reference/InfoPlistKeyReference/Articles/AppExtensionKeys.html#//apple_ref/doc/uid/TP40014212-SW15
// [NSExtensionPrincipalClass]: https://developer.apple.com/library/archive/documentation/General/Reference/InfoPlistKeyReference/Articles/AppExtensionKeys.html#//apple_ref/doc/uid/TP40014212-SW16
// [Setting up a remote notification server]: https://developer.apple.com/documentation/UserNotifications/setting-up-a-remote-notification-server
type UNNotificationServiceExtension struct {
	objectivec.Object
}

// UNNotificationServiceExtensionFromID constructs a [UNNotificationServiceExtension] from an objc.ID.
//
// An object that modifies the content of a remote notification before it’s
// delivered to the user.
func UNNotificationServiceExtensionFromID(id objc.ID) UNNotificationServiceExtension {
	return UNNotificationServiceExtension{objectivec.Object{ID: id}}
}

// NOTE: UNNotificationServiceExtension adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [UNNotificationServiceExtension] class.
//
// # Processing Notifications
//
//   - [IUNNotificationServiceExtension.DidReceiveNotificationRequestWithContentHandler]: Asks you to make any needed changes to the notification and notify the system when you’re done.
//   - [IUNNotificationServiceExtension.ServiceExtensionTimeWillExpire]: Tells you that the system is terminating your extension.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationServiceExtension
type IUNNotificationServiceExtension interface {
	objectivec.IObject

	// Topic: Processing Notifications

	// Asks you to make any needed changes to the notification and notify the system when you’re done.
	DidReceiveNotificationRequestWithContentHandler(request IUNNotificationRequest, contentHandler UNNotificationContentHandler)
	// Tells you that the system is terminating your extension.
	ServiceExtensionTimeWillExpire()
}

// Init initializes the instance.
func (u UNNotificationServiceExtension) Init() UNNotificationServiceExtension {
	rv := objc.Send[UNNotificationServiceExtension](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u UNNotificationServiceExtension) Autorelease() UNNotificationServiceExtension {
	rv := objc.Send[UNNotificationServiceExtension](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewUNNotificationServiceExtension creates a new UNNotificationServiceExtension instance.
func NewUNNotificationServiceExtension() UNNotificationServiceExtension {
	class := getUNNotificationServiceExtensionClass()
	rv := objc.Send[UNNotificationServiceExtension](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Asks you to make any needed changes to the notification and notify the
// system when you’re done.
//
// request: The original notification request. Use this object to get the original
// content of the notification.
//
// contentHandler: The block to execute with the modified content. This block has no return
// value and takes the following parameter:
//
// contentToDeliver: A [UNNotificationContent] object with the content the
// system displays to the user.
//
// # Discussion
//
// Override this method and use it to modify the [UNNotificationContent]
// object that the system delivers with the notification. At some point during
// your implementation, execute the `contentHandler` block and pass it your
// modified content. If you decide not to modify the content, call the
// `contentHandler` block with the original content from the `request`
// parameter.
//
// You can modify any of the content from the original request. You might
// customize the content for the current user or replace it altogether. You
// can use this method to download images or movies and add them as
// attachments to the content. You may also modify the alert text as long as
// you don’t remove it. If the content object doesn’t contain any alert
// text, the system ignores your modifications and delivers the original
// notification content.
//
// Your extension has a limited amount of time (no more than 30 seconds) to
// modify the content and execute the `contentHandler` block. If you don’t
// execute that block in a timely manner, the system calls your extension’s
// [ServiceExtensionTimeWillExpire] method to give you one last chance to
// execute the block. If you don’t, the system presents the notification’s
// original content to the user.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationServiceExtension/didReceive(_:withContentHandler:)
func (u UNNotificationServiceExtension) DidReceiveNotificationRequestWithContentHandler(request IUNNotificationRequest, contentHandler UNNotificationContentHandler) {
	_block1, _ := NewUNNotificationContentBlock(contentHandler)
	objc.Send[objc.ID](u.ID, objc.Sel("didReceiveNotificationRequest:withContentHandler:"), request, _block1)
}

// Tells you that the system is terminating your extension.
//
// # Discussion
//
// If your [DidReceiveNotificationRequestWithContentHandler] method takes too
// long to execute its completion block, the system calls this method on a
// separate thread to give you one last chance to execute the block. Use this
// method to execute the block as quickly as possible. Doing so might mean
// providing some fallback content. For example, if your extension is still
// downloading an image file with the intent of attaching it to the
// notification’s content, update the notification’s alert text to
// indicate that an image download is in progress. If you fail to execute the
// completion block from the [DidReceiveNotificationRequestWithContentHandler]
// method in time, the system displays the notification’s original content.
//
// See: https://developer.apple.com/documentation/UserNotifications/UNNotificationServiceExtension/serviceExtensionTimeWillExpire()
func (u UNNotificationServiceExtension) ServiceExtensionTimeWillExpire() {
	objc.Send[objc.ID](u.ID, objc.Sel("serviceExtensionTimeWillExpire"))
}

// DidReceiveNotificationRequestWithContentHandlerSync is a synchronous wrapper around [UNNotificationServiceExtension.DidReceiveNotificationRequestWithContentHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (u UNNotificationServiceExtension) DidReceiveNotificationRequestWithContentHandlerSync(ctx context.Context, request IUNNotificationRequest) (*UNNotificationContent, error) {
	done := make(chan *UNNotificationContent, 1)
	u.DidReceiveNotificationRequestWithContentHandler(request, func(val *UNNotificationContent) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
