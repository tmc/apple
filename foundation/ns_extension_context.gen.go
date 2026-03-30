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
//   - [NSExtensionContext.NSExtensionItemsAndErrorsKey]: The extension items and errors key.
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
// # Working with notifications
//
//   - [NSExtensionContext.NSExtensionHostDidBecomeActive]: Posted when the extension’s host app moves from the inactive to the active state.
//   - [NSExtensionContext.NSExtensionHostWillResignActive]: Posted when the extension’s host app moves from the active to the inactive state.
//   - [NSExtensionContext.NSExtensionHostDidEnterBackground]: Posted when the extension’s host app begins running in the background.
//   - [NSExtensionContext.NSExtensionHostWillEnterForeground]: Posted when the extension’s host app begins running in the foreground.
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
//   - [INSExtensionContext.NSExtensionItemsAndErrorsKey]: The extension items and errors key.
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
// # Working with notifications
//
//   - [INSExtensionContext.NSExtensionHostDidBecomeActive]: Posted when the extension’s host app moves from the inactive to the active state.
//   - [INSExtensionContext.NSExtensionHostWillResignActive]: Posted when the extension’s host app moves from the active to the inactive state.
//   - [INSExtensionContext.NSExtensionHostDidEnterBackground]: Posted when the extension’s host app begins running in the background.
//   - [INSExtensionContext.NSExtensionHostWillEnterForeground]: Posted when the extension’s host app begins running in the foreground.
//
// See: https://developer.apple.com/documentation/Foundation/NSExtensionContext
type INSExtensionContext interface {
	objectivec.IObject

	// Topic: Handling requests

	// Tells the host app to complete the app extension request with an array of result items.
	CompleteRequestReturningItemsCompletionHandler(items INSArray, completionHandler BoolHandler)
	// Tells the host app to cancel the app extension request, with a supplied error.
	CancelRequestWithError(error_ INSError)
	// The extension items and errors key.
	NSExtensionItemsAndErrorsKey() string

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

	LoadBroadcastingApplicationInfoWithCompletion(handler VoidHandler)
	CompleteRequestWithBroadcastURLSetupInfo(broadcastURL INSURL, setupInfo INSDictionary)

	// Topic: Handling notification actions

	NotificationActions() []objectivec.IObject
	SetNotificationActions(value []objectivec.IObject)
	PerformNotificationDefaultAction()
	DismissNotificationContentExtension()

	// Topic: Working with notifications

	// Posted when the extension’s host app moves from the inactive to the active state.
	NSExtensionHostDidBecomeActive() NSNotificationName
	// Posted when the extension’s host app moves from the active to the inactive state.
	NSExtensionHostWillResignActive() NSNotificationName
	// Posted when the extension’s host app begins running in the background.
	NSExtensionHostDidEnterBackground() NSNotificationName
	// Posted when the extension’s host app begins running in the foreground.
	NSExtensionHostWillEnterForeground() NSNotificationName
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
func (e NSExtensionContext) LoadBroadcastingApplicationInfoWithCompletion(handler VoidHandler) {
	_block0, _ := NewVoidBlock(handler)
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

// The extension items and errors key.
//
// See: https://developer.apple.com/documentation/foundation/nsextensionitemsanderrorskey
func (e NSExtensionContext) NSExtensionItemsAndErrorsKey() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("NSExtensionItemsAndErrorsKey"))
	return NSStringFromID(rv).String()
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

// Posted when the extension’s host app moves from the inactive to the
// active state.
//
// See: https://developer.apple.com/documentation/foundation/nsnotification/name-swift.struct/nsextensionhostdidbecomeactive
func (e NSExtensionContext) NSExtensionHostDidBecomeActive() NSNotificationName {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("NSExtensionHostDidBecomeActiveNotification"))
	return NSNotificationName(NSStringFromID(rv).String())
}

// Posted when the extension’s host app moves from the active to the
// inactive state.
//
// See: https://developer.apple.com/documentation/foundation/nsnotification/name-swift.struct/nsextensionhostwillresignactive
func (e NSExtensionContext) NSExtensionHostWillResignActive() NSNotificationName {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("NSExtensionHostWillResignActiveNotification"))
	return NSNotificationName(NSStringFromID(rv).String())
}

// Posted when the extension’s host app begins running in the background.
//
// See: https://developer.apple.com/documentation/foundation/nsnotification/name-swift.struct/nsextensionhostdidenterbackground
func (e NSExtensionContext) NSExtensionHostDidEnterBackground() NSNotificationName {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("NSExtensionHostDidEnterBackgroundNotification"))
	return NSNotificationName(NSStringFromID(rv).String())
}

// Posted when the extension’s host app begins running in the foreground.
//
// See: https://developer.apple.com/documentation/foundation/nsnotification/name-swift.struct/nsextensionhostwillenterforeground
func (e NSExtensionContext) NSExtensionHostWillEnterForeground() NSNotificationName {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("NSExtensionHostWillEnterForegroundNotification"))
	return NSNotificationName(NSStringFromID(rv).String())
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

// LoadBroadcastingApplicationInfo is a synchronous wrapper around [NSExtensionContext.LoadBroadcastingApplicationInfoWithCompletion].
// It blocks until the completion handler fires or the context is cancelled.
func (e NSExtensionContext) LoadBroadcastingApplicationInfo(ctx context.Context) error {
	done := make(chan struct{}, 1)
	e.LoadBroadcastingApplicationInfoWithCompletion(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
