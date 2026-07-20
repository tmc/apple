// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"context"
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSExtensionContext] class.
var (
	_NSExtensionContextClass     NSExtensionContextClass
	_NSExtensionContextClassOnce sync.Once
)

func getNSExtensionContextClass() NSExtensionContextClass {
	_NSExtensionContextClassOnce.Do(func() {
		_NSExtensionContextClass = NSExtensionContextClass{class: objc.GetClass("NSExtensionContext")}
	})
	return _NSExtensionContextClass
}

// GetNSExtensionContextClass returns the class object for NSExtensionContext.
func GetNSExtensionContextClass() NSExtensionContextClass {
	return getNSExtensionContextClass()
}

type NSExtensionContextClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSExtensionContextClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSExtensionContextClass) Alloc() NSExtensionContext {
	rv := objc.Send[NSExtensionContext](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The host app context from which an app extension is invoked.
//
// # Overview
//
// When a host app sends a request to an app extension, it provides an
// extension context. For many app extensions, the most important part of the
// context is the data the user wants to work with, which is contained in the
// [NSExtensionContext.InputItems] property.
//
// # Handling requests
//
//   - [NSExtensionContext.CompleteRequestReturningItemsCompletionHandler]: Tells the host app to complete the app extension request with an array of result items.
//   - [NSExtensionContext.CancelRequestWithError]: Tells the host app to cancel the app extension request, with a supplied error.
//
// # Opening URLs
//
//   - [NSExtensionContext.OpenURLCompletionHandler]: Asks the system to open a URL on behalf of the currently running app extension.
//
// # Storing extension items
//
//   - [NSExtensionContext.InputItems]: The list of input [NSExtensionItem](<doc://com.apple.foundation/documentation/Foundation/NSExtensionItem>) objects associated with the context.
//
// # Controlling media playback in notification content extensions
//
//   - [NSExtensionContext.MediaPlayingStarted]: Tells the system that the Notification Content app extension began playing a media file.
//   - [NSExtensionContext.MediaPlayingPaused]: Tells the system that the Notification Content app extension stopped playing a media file.
//
// # Supporting broadcasting
//
//   - [NSExtensionContext.LoadBroadcastingApplicationInfoWithCompletion]
//   - [NSExtensionContext.CompleteRequestWithBroadcastURLSetupInfo]
//
// # Handling notification actions
//
//   - [NSExtensionContext.NotificationActions]
//   - [NSExtensionContext.SetNotificationActions]
//   - [NSExtensionContext.PerformNotificationDefaultAction]
//   - [NSExtensionContext.DismissNotificationContentExtension]
//
// See: https://developer.apple.com/documentation/Foundation/NSExtensionContext
type NSExtensionContext struct {
	objectivec.Object
}

// NSExtensionContextFromID constructs a [NSExtensionContext] from an objc.ID.
//
// The host app context from which an app extension is invoked.
func NSExtensionContextFromID(id objc.ID) NSExtensionContext {
	return NSExtensionContext{objectivec.Object{ID: id}}
}

// NOTE: NSExtensionContext adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSExtensionContext] class.
//
// # Handling requests
//
//   - [INSExtensionContext.CompleteRequestReturningItemsCompletionHandler]: Tells the host app to complete the app extension request with an array of result items.
//   - [INSExtensionContext.CancelRequestWithError]: Tells the host app to cancel the app extension request, with a supplied error.
//
// # Opening URLs
//
//   - [INSExtensionContext.OpenURLCompletionHandler]: Asks the system to open a URL on behalf of the currently running app extension.
//
// # Storing extension items
//
//   - [INSExtensionContext.InputItems]: The list of input [NSExtensionItem](<doc://com.apple.foundation/documentation/Foundation/NSExtensionItem>) objects associated with the context.
//
// # Controlling media playback in notification content extensions
//
//   - [INSExtensionContext.MediaPlayingStarted]: Tells the system that the Notification Content app extension began playing a media file.
//   - [INSExtensionContext.MediaPlayingPaused]: Tells the system that the Notification Content app extension stopped playing a media file.
//
// # Supporting broadcasting
//
//   - [INSExtensionContext.LoadBroadcastingApplicationInfoWithCompletion]
//   - [INSExtensionContext.CompleteRequestWithBroadcastURLSetupInfo]
//
// # Handling notification actions
//
//   - [INSExtensionContext.NotificationActions]
//   - [INSExtensionContext.SetNotificationActions]
//   - [INSExtensionContext.PerformNotificationDefaultAction]
//   - [INSExtensionContext.DismissNotificationContentExtension]
//
// See: https://developer.apple.com/documentation/Foundation/NSExtensionContext
type INSExtensionContext interface {
	objectivec.IObject

	// Topic: Handling requests

	// Tells the host app to complete the app extension request with an array of result items.
	CompleteRequestReturningItemsCompletionHandler(items INSArray, completionHandler BoolHandler)
	// Tells the host app to cancel the app extension request, with a supplied error.
	CancelRequestWithError(error_ INSError)

	// Topic: Opening URLs

	// Asks the system to open a URL on behalf of the currently running app extension.
	OpenURLCompletionHandler(URL INSURL, completionHandler BoolHandler)

	// Topic: Storing extension items

	// The list of input [NSExtensionItem](<doc://com.apple.foundation/documentation/Foundation/NSExtensionItem>) objects associated with the context.
	InputItems() INSArray

	// Topic: Controlling media playback in notification content extensions

	// Tells the system that the Notification Content app extension began playing a media file.
	MediaPlayingStarted()
	// Tells the system that the Notification Content app extension stopped playing a media file.
	MediaPlayingPaused()

	// Topic: Supporting broadcasting

	LoadBroadcastingApplicationInfoWithCompletion(handler StringStringImageHandler)
	CompleteRequestWithBroadcastURLSetupInfo(broadcastURL INSURL, setupInfo INSDictionary)

	// Topic: Handling notification actions

	NotificationActions() []objectivec.IObject
	SetNotificationActions(value []objectivec.IObject)
	PerformNotificationDefaultAction()
	DismissNotificationContentExtension()
}

// Init initializes the instance.
func (e NSExtensionContext) Init() NSExtensionContext {
	rv := objc.Send[NSExtensionContext](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e NSExtensionContext) Autorelease() NSExtensionContext {
	rv := objc.Send[NSExtensionContext](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSExtensionContext creates a new NSExtensionContext instance.
func NewNSExtensionContext() NSExtensionContext {
	class := getNSExtensionContextClass()
	rv := objc.Send[NSExtensionContext](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Tells the host app to complete the app extension request with an array of
// result items.
//
// items: An array of result items, each an [NSExtensionItem] object, to return to
// the host app.
//
// completionHandler: An optional block to be called when the request completes, performed as a
// background priority task.
//
// The block takes the following parameter:
//
// expired: A Boolean value that indicates whether the system is terminating a
// previous invocation of the `completionHandler` block.
//
// This parameter is true when the system prematurely terminates a
// `completionHandler` block that was previously invoked and had not otherwise
// expired.
//
// # Discussion
//
// Calling this method eventually dismisses the app extension’s view
// controller.
//
// See: https://developer.apple.com/documentation/Foundation/NSExtensionContext/completeRequest(returningItems:completionHandler:)
func (e NSExtensionContext) CompleteRequestReturningItemsCompletionHandler(items INSArray, completionHandler BoolHandler) {
	_block1, _ := NewBoolBlock(completionHandler)
	objc.Send[objc.ID](e.ID, objc.Sel("completeRequestReturningItems:completionHandler:"), items, _block1)
}

// Tells the host app to cancel the app extension request, with a supplied
// error.
//
// error: The error object to return. It must be non-`nil`.
//
// # Discussion
//
// On return, the `userInfo` dictionary of the [NSError] object contains a key
// named [NSExtensionItemsAndErrorsKey] which has as its value a dictionary of
// [NSExtensionItem] objects and associated [NSError] instances.
//
// See: https://developer.apple.com/documentation/Foundation/NSExtensionContext/cancelRequest(withError:)
//
// [NSExtensionItemsAndErrorsKey]: https://developer.apple.com/documentation/Foundation/NSExtensionItemsAndErrorsKey
func (e NSExtensionContext) CancelRequestWithError(error_ INSError) {
	objc.Send[objc.ID](e.ID, objc.Sel("cancelRequestWithError:"), error_)
}

// Asks the system to open a URL on behalf of the currently running app
// extension.
//
// URL: The URL to open.
//
// completionHandler: A block/closure to be called when the URL has opened. The closure takes a
// single boolean parameter indicating whether the operation was successful.
//
// # Discussion
//
// Each extension point determines whether to support this method, or under
// which conditions to support this method. In iOS, the Today and iMessage app
// extension points support this method. An iMessage app extension can use
// this method only to open its parent app, and only if the parent app is
// shown on the iOS home screen.
//
// See: https://developer.apple.com/documentation/Foundation/NSExtensionContext/open(_:completionHandler:)
func (e NSExtensionContext) OpenURLCompletionHandler(URL INSURL, completionHandler BoolHandler) {
	_block1, _ := NewBoolBlock(completionHandler)
	objc.Send[objc.ID](e.ID, objc.Sel("openURL:completionHandler:"), URL, _block1)
}

// Tells the system that the Notification Content app extension began playing
// a media file.
//
// # Discussion
//
// In your Notification Content app extension code, call this method when you
// programmatically begin playing a media file. When called, the system
// updates the appearance of the media playback button displayed in the
// notification content extension’s interface. For more information about
// implementing a notification content extension, see
// [UNNotificationContentExtension].
//
// See: https://developer.apple.com/documentation/Foundation/NSExtensionContext/mediaPlayingStarted()
//
// [UNNotificationContentExtension]: https://developer.apple.com/documentation/UserNotificationsUI/UNNotificationContentExtension
func (e NSExtensionContext) MediaPlayingStarted() {
	objc.Send[objc.ID](e.ID, objc.Sel("mediaPlayingStarted"))
}

// Tells the system that the Notification Content app extension stopped
// playing a media file.
//
// # Discussion
//
// In your Notification Content app extension code, call this method when you
// programmatically stop playing a media file. When called, the system updates
// the appearance of the media playback button displayed in the notification
// content extension’s interface. For more information about implementing a
// notification content extension, see [UNNotificationContentExtension].
//
// See: https://developer.apple.com/documentation/Foundation/NSExtensionContext/mediaPlayingPaused()
//
// [UNNotificationContentExtension]: https://developer.apple.com/documentation/UserNotificationsUI/UNNotificationContentExtension
func (e NSExtensionContext) MediaPlayingPaused() {
	objc.Send[objc.ID](e.ID, objc.Sel("mediaPlayingPaused"))
}

// # Discussion
//
// See: https://developer.apple.com/documentation/Foundation/NSExtensionContext/loadBroadcastingApplicationInfo(completion:)
func (e NSExtensionContext) LoadBroadcastingApplicationInfoWithCompletion(handler StringStringImageHandler) {
	_block0, _ := NewStringStringImageBlock(handler)
	objc.Send[objc.ID](e.ID, objc.Sel("loadBroadcastingApplicationInfoWithCompletion:"), _block0)
}

// See: https://developer.apple.com/documentation/Foundation/NSExtensionContext/completeRequest(withBroadcast:setupInfo:)
func (e NSExtensionContext) CompleteRequestWithBroadcastURLSetupInfo(broadcastURL INSURL, setupInfo INSDictionary) {
	objc.Send[objc.ID](e.ID, objc.Sel("completeRequestWithBroadcastURL:setupInfo:"), broadcastURL, setupInfo)
}

// See: https://developer.apple.com/documentation/Foundation/NSExtensionContext/performNotificationDefaultAction()
func (e NSExtensionContext) PerformNotificationDefaultAction() {
	objc.Send[objc.ID](e.ID, objc.Sel("performNotificationDefaultAction"))
}

// See: https://developer.apple.com/documentation/Foundation/NSExtensionContext/dismissNotificationContentExtension()
func (e NSExtensionContext) DismissNotificationContentExtension() {
	objc.Send[objc.ID](e.ID, objc.Sel("dismissNotificationContentExtension"))
}

// The list of input [NSExtensionItem] objects associated with the context.
//
// # Discussion
//
// If the context has no input items, this array is empty.
//
// See: https://developer.apple.com/documentation/Foundation/NSExtensionContext/inputItems
func (e NSExtensionContext) InputItems() INSArray {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("inputItems"))
	return NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/Foundation/NSExtensionContext/notificationActions
func (e NSExtensionContext) NotificationActions() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](e.ID, objc.Sel("notificationActions"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
func (e NSExtensionContext) SetNotificationActions(value []objectivec.IObject) {
	objc.Send[struct{}](e.ID, objc.Sel("setNotificationActions:"), objectivec.IObjectSliceToNSArray(value))
}

// CompleteRequestReturningItems is a synchronous wrapper around [NSExtensionContext.CompleteRequestReturningItemsCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (e NSExtensionContext) CompleteRequestReturningItems(ctx context.Context, items INSArray) (bool, error) {
	done := make(chan bool, 1)
	e.CompleteRequestReturningItemsCompletionHandler(items, func(val bool) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return false, ctx.Err()
	}
}

// OpenURL is a synchronous wrapper around [NSExtensionContext.OpenURLCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (e NSExtensionContext) OpenURL(ctx context.Context, URL INSURL) (bool, error) {
	done := make(chan bool, 1)
	e.OpenURLCompletionHandler(URL, func(val bool) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return false, ctx.Err()
	}
}
