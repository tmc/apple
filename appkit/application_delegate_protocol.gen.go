// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf


// A set of methods that manage your app’s life cycle and its interaction with common system services.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplicationDelegate
type NSApplicationDelegate interface {
	objectivec.IObject
}



// NSApplicationDelegateObject wraps an existing Objective-C object that conforms to the NSApplicationDelegate protocol.
type NSApplicationDelegateObject struct {
	objectivec.Object
}
func (o NSApplicationDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}



// NSApplicationDelegateObjectFromID constructs a [NSApplicationDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSApplicationDelegateObjectFromID(id objc.ID) NSApplicationDelegateObject {
	return NSApplicationDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}




// Tells the delegate that the app’s initialization is about to complete.
//
// notification: A notification named [willFinishLaunchingNotification]. Calling the
// [object] method of this notification returns the [NSApplication] object
// itself.
// //
// [object]: https://developer.apple.com/documentation/Foundation/NSNotification/object
// [willFinishLaunchingNotification]: https://developer.apple.com/documentation/AppKit/NSApplication/willFinishLaunchingNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSApplicationDelegate/applicationWillFinishLaunching(_:)

func (o NSApplicationDelegateObject) ApplicationWillFinishLaunching(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("applicationWillFinishLaunching:"), notification)
	}

// Tells the delegate that the app’s initialization is complete but it
// hasn’t received its first event.
//
// notification: A notification named [didFinishLaunchingNotification]. Calling the [object]
// method of this notification returns the [NSApplication] object itself.
// //
// [didFinishLaunchingNotification]: https://developer.apple.com/documentation/AppKit/NSApplication/didFinishLaunchingNotification
// [object]: https://developer.apple.com/documentation/Foundation/NSNotification/object
//
// # Discussion
// 
// Delegates can implement this method to perform further initialization. This
// method is called after the application’s main run loop has been started
// but before it has processed any events. If the application was launched by
// the user opening a file, the delegate’s [ApplicationOpenFile] method is
// called before this method. If you want to perform initialization before any
// files are opened, implement the [ApplicationWillFinishLaunching] method in
// your delegate, which is called before [ApplicationOpenFile].)
//
// See: https://developer.apple.com/documentation/AppKit/NSApplicationDelegate/applicationDidFinishLaunching(_:)

func (o NSApplicationDelegateObject) ApplicationDidFinishLaunching(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("applicationDidFinishLaunching:"), notification)
	}

// Tells the delegate that the app is about to become active.
//
// notification: A notification named [willBecomeActiveNotification]. Calling the [object]
// method of this notification returns the [NSApplication] object itself.
// //
// [object]: https://developer.apple.com/documentation/Foundation/NSNotification/object
// [willBecomeActiveNotification]: https://developer.apple.com/documentation/AppKit/NSApplication/willBecomeActiveNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSApplicationDelegate/applicationWillBecomeActive(_:)

func (o NSApplicationDelegateObject) ApplicationWillBecomeActive(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("applicationWillBecomeActive:"), notification)
	}

// Tells the delegate that the app is now active.
//
// notification: A notification named [didBecomeActiveNotification]. Calling the [object]
// method of this notification returns the [NSApplication] object itself.
// //
// [didBecomeActiveNotification]: https://developer.apple.com/documentation/AppKit/NSApplication/didBecomeActiveNotification
// [object]: https://developer.apple.com/documentation/Foundation/NSNotification/object
//
// See: https://developer.apple.com/documentation/AppKit/NSApplicationDelegate/applicationDidBecomeActive(_:)

func (o NSApplicationDelegateObject) ApplicationDidBecomeActive(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("applicationDidBecomeActive:"), notification)
	}

// Tells the delegate that the app is about to become inactive and will lose
// focus.
//
// notification: A notification named [willResignActiveNotification]. Calling the [object]
// method of this notification returns the [NSApplication] object itself.
// //
// [object]: https://developer.apple.com/documentation/Foundation/NSNotification/object
// [willResignActiveNotification]: https://developer.apple.com/documentation/AppKit/NSApplication/willResignActiveNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSApplicationDelegate/applicationWillResignActive(_:)

func (o NSApplicationDelegateObject) ApplicationWillResignActive(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("applicationWillResignActive:"), notification)
	}

// Tells the delegate that the app is no longer active and doesn’t have
// focus.
//
// notification: A notification named [didResignActiveNotification]. Calling the [object]
// method of this notification returns the [NSApplication] object itself.
// //
// [didResignActiveNotification]: https://developer.apple.com/documentation/AppKit/NSApplication/didResignActiveNotification
// [object]: https://developer.apple.com/documentation/Foundation/NSNotification/object
//
// See: https://developer.apple.com/documentation/AppKit/NSApplicationDelegate/applicationDidResignActive(_:)

func (o NSApplicationDelegateObject) ApplicationDidResignActive(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("applicationDidResignActive:"), notification)
	}

// Returns a value that indicates if the app should terminate.
//
// sender: The application object that is about to be terminated.
//
// # Return Value
// 
// One of the values defined in [NSApplication.TerminateReply] constants
// indicating whether the application should terminate. For compatibility
// reasons, a return value of [false] is equivalent to
// [NSApplication.TerminateReply.terminateCancel], and a return value of
// [true] is equivalent to [NSApplication.TerminateReply.terminateNow].
//
// [NSApplication.TerminateReply.terminateCancel]: https://developer.apple.com/documentation/AppKit/NSApplication/TerminateReply/terminateCancel
// [NSApplication.TerminateReply.terminateNow]: https://developer.apple.com/documentation/AppKit/NSApplication/TerminateReply/terminateNow
// [NSApplication.TerminateReply]: https://developer.apple.com/documentation/AppKit/NSApplication/TerminateReply
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method is called after the application’s Quit menu item has been
// selected, or after the [Terminate] method has been called. Generally, you
// should return [NSApplication.TerminateReply.terminateNow] to allow the
// termination to complete, but you can cancel the termination process or
// delay it somewhat as needed. For example, you might delay termination to
// finish processing some critical data but then terminate the application as
// soon as you are done by calling the [ReplyToApplicationShouldTerminate]
// method.
//
// [NSApplication.TerminateReply.terminateNow]: https://developer.apple.com/documentation/AppKit/NSApplication/TerminateReply/terminateNow
//
// See: https://developer.apple.com/documentation/AppKit/NSApplicationDelegate/applicationShouldTerminate(_:)

func (o NSApplicationDelegateObject) ApplicationShouldTerminate(sender INSApplication) NSApplicationTerminateReply {
	
	rv := objc.Send[NSApplicationTerminateReply](o.ID, objc.Sel("applicationShouldTerminate:"), sender)
	return rv
	}

// Returns a Boolean value that indicates if the app terminates once the last
// window closes.
//
// sender: The application object whose last window was closed.
//
// # Return Value
// 
// [false] if the application should not be terminated when its last window is
// closed; otherwise, [true] to terminate the application.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The application sends this message to your delegate when the
// application’s last window is closed. It sends this message regardless of
// whether there are still panels open. (A panel in this case is defined as
// being an instance of [NSPanel] or one of its subclasses.)
// 
// If your implementation returns [false], control returns to the main event
// loop and the application is not terminated. If you return [true], your
// delegate’s [ApplicationShouldTerminate] method is subsequently invoked to
// confirm that the application should be terminated.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSApplicationDelegate/applicationShouldTerminateAfterLastWindowClosed(_:)

func (o NSApplicationDelegateObject) ApplicationShouldTerminateAfterLastWindowClosed(sender INSApplication) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("applicationShouldTerminateAfterLastWindowClosed:"), sender)
	return rv
	}

// Tells the delegate that the app is about to terminate.
//
// notification: A notification named [willTerminateNotification]. Calling the [object]
// method of this notification returns the [NSApplication] object itself.
// //
// [object]: https://developer.apple.com/documentation/Foundation/NSNotification/object
// [willTerminateNotification]: https://developer.apple.com/documentation/AppKit/NSApplication/willTerminateNotification
//
// # Discussion
// 
// Your delegate can use this method to perform any final cleanup before the
// app terminates. The app will terminate after this method returns.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplicationDelegate/applicationWillTerminate(_:)

func (o NSApplicationDelegateObject) ApplicationWillTerminate(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("applicationWillTerminate:"), notification)
	}

// Tells the delegate that the app is about to be hidden.
//
// notification: A notification named [willHideNotification]. Calling the [object] method of
// this notification returns the [NSApplication] object itself.
// //
// [object]: https://developer.apple.com/documentation/Foundation/NSNotification/object
// [willHideNotification]: https://developer.apple.com/documentation/AppKit/NSApplication/willHideNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSApplicationDelegate/applicationWillHide(_:)

func (o NSApplicationDelegateObject) ApplicationWillHide(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("applicationWillHide:"), notification)
	}

// Tells the delegate that the app is now hidden.
//
// notification: A notification named [didHideNotification]. Calling the [object] method of
// this notification returns the [NSApplication] object itself.
// //
// [didHideNotification]: https://developer.apple.com/documentation/AppKit/NSApplication/didHideNotification
// [object]: https://developer.apple.com/documentation/Foundation/NSNotification/object
//
// See: https://developer.apple.com/documentation/AppKit/NSApplicationDelegate/applicationDidHide(_:)

func (o NSApplicationDelegateObject) ApplicationDidHide(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("applicationDidHide:"), notification)
	}

// Tells the delegate that the app is about to become visible.
//
// notification: A notification named [willUnhideNotification]. Calling the [object] method
// of this notification returns the [NSApplication] object itself.
// //
// [object]: https://developer.apple.com/documentation/Foundation/NSNotification/object
// [willUnhideNotification]: https://developer.apple.com/documentation/AppKit/NSApplication/willUnhideNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSApplicationDelegate/applicationWillUnhide(_:)

func (o NSApplicationDelegateObject) ApplicationWillUnhide(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("applicationWillUnhide:"), notification)
	}

// Tells the delegate that the app is now visible.
//
// notification: A notification named [didUnhideNotification]. Calling the [object] method
// of this notification returns the [NSApplication] object itself.
// //
// [didUnhideNotification]: https://developer.apple.com/documentation/AppKit/NSApplication/didUnhideNotification
// [object]: https://developer.apple.com/documentation/Foundation/NSNotification/object
//
// See: https://developer.apple.com/documentation/AppKit/NSApplicationDelegate/applicationDidUnhide(_:)

func (o NSApplicationDelegateObject) ApplicationDidUnhide(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("applicationDidUnhide:"), notification)
	}

// Tells the delegate that the app is about to update its windows.
//
// notification: A notification named [willUpdateNotification]. Calling the [object] method
// of this notification returns the [NSApplication] object itself.
// //
// [object]: https://developer.apple.com/documentation/Foundation/NSNotification/object
// [willUpdateNotification]: https://developer.apple.com/documentation/AppKit/NSApplication/willUpdateNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSApplicationDelegate/applicationWillUpdate(_:)

func (o NSApplicationDelegateObject) ApplicationWillUpdate(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("applicationWillUpdate:"), notification)
	}

// Tells the delegate that the app’s windows did update.
//
// notification: A notification named [didUpdateNotification]. Calling the [object] method
// of this notification returns the [NSApplication] object itself.
// //
// [didUpdateNotification]: https://developer.apple.com/documentation/AppKit/NSApplication/didUpdateNotification
// [object]: https://developer.apple.com/documentation/Foundation/NSNotification/object
//
// See: https://developer.apple.com/documentation/AppKit/NSApplicationDelegate/applicationDidUpdate(_:)

func (o NSApplicationDelegateObject) ApplicationDidUpdate(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("applicationDidUpdate:"), notification)
	}

// Returns a Boolean value that indicates if the app responds to reopen
// AppleEvents.
//
// sender: The application object.
//
// hasVisibleWindows: Indicates whether the [NSApplication] object found any visible windows in
// your application. You can use this value as an indication of whether the
// application would do anything if you return [true].
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// [true] if you want the application to perform its normal tasks or [false]
// if you want the application to do nothing.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// These events are sent whenever the Finder reactivates an already running
// application because someone double-clicked it again or used the dock to
// activate it.
// 
// By default the Application Kit will handle this event by checking whether
// there are any visible [NSWindow] (not [NSPanel]) objects, and, if there are
// none, it goes through the standard untitled document creation (the same as
// it does if `theApplication` is launched without any document to open). For
// most document-based applications, an untitled document will be created.
// 
// The application delegate will also get a chance to respond to the normal
// untitled document delegate methods. If you implement this method in your
// application delegate, it will be called before any of the default behavior
// happens. If you return [true], then [NSApplication] will proceed as normal.
// If you return [false], then [NSApplication] will do nothing. So, you can
// either implement this method with a version that does nothing, and return
// [false] if you do not want anything to happen at all (not recommended), or
// you can implement this method, handle the event yourself in some custom
// way, and return [false].
// 
// Miniaturized windows, windows in the Dock, are considered visible by this
// method, and cause `flag` to return [true], despite the fact that
// miniaturized windows return [false] when sent an [Visible] message.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSApplicationDelegate/applicationShouldHandleReopen(_:hasVisibleWindows:)

func (o NSApplicationDelegateObject) ApplicationShouldHandleReopenHasVisibleWindows(sender INSApplication, hasVisibleWindows bool) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("applicationShouldHandleReopen:hasVisibleWindows:"), sender, hasVisibleWindows)
	return rv
	}

// Returns the app’s dock menu.
//
// sender: The application object associated with the delegate.
//
// # Return Value
// 
// The menu to display in the dock.
//
// # Discussion
// 
// You can also connect a menu in Interface Builder to the `dockMenu` outlet.
// A third way for your application to specify a dock menu is to provide an
// [NSMenu] in a nib.
// 
// If this method returns a menu, this menu takes precedence over the
// `dockMenu` in the nib.
// 
// The target and action for each menu item are passed to the dock. On
// selection of the menu item the dock messages your application, which should
// invoke `[NSApp selector target nil]`.
// 
// To specify an [NSMenu] in a nib, you add the nib name to the
// `info.Plist()`, using the key [AppleDockMenu]. The nib name is specified
// without an extension. You then create a connection from the file’s owner
// object (which by default is [NSApplication]) to the menu. Connect the menu
// to the `dockMenu` outlet of [NSApplication]. The menu is in its own nib
// file so it can be loaded lazily when the `dockMenu` is requested, rather
// than at launch time.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplicationDelegate/applicationDockMenu(_:)

func (o NSApplicationDelegateObject) ApplicationDockMenu(sender INSApplication) INSMenu {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("applicationDockMenu:"), sender)
	return NSMenuFromID(rv)
	}

// Returns a Boolean value that tells the system whether to remap menu
// shortcuts to support localized keyboards.
//
// application: The app object associated with the delegate.
//
// # Return Value
// 
// [true] to enable automatic shortcut localization for all app-specific
// menus, or [false] to handle shortcut localization yourself.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// A keyboard shortcut you specify in one language might be difficult or
// impossible to reproduce on a keyboard with a different character set or
// layout. Localized keyboards sometimes rearrange punctuation marks or
// replace them altogether to make room for a language’s required
// characters. The new locations of those keys might make it difficult to use
// your current shortcuts. To ensure your shortcuts are always usable, the
// system can automatically remap shortcuts, as needed, to accommodate the
// connected keyboard.
// 
// If you don’t implement this method, or if you implement it and return
// [true], the system automatically remaps all app-specific menu shortcuts
// that are unreachable on the current keyboard. The system doesn’t remap
// shortcuts when the input keys have identical positions on both keyboards,
// or when a shortcut is still easily reachable on the current keyboard. The
// remapping is transparent to your app.
// 
// If you already localize your app’s shortcuts for different languages, or
// if you allow someone to customize your app’s shortcuts, you can return
// [false] to disable the automatic remapping behavior. When you return
// [false], the system doesn’t change your app’s shortcuts. Instead, you
// are responsible to make any required changes to support localized
// keyboards.
// 
// During the final activation of your app at launch time, the app object’s
// [FinishLaunching] method calls this method once to record your app’s
// response.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSApplicationDelegate/applicationShouldAutomaticallyLocalizeKeyEquivalents(_:)

func (o NSApplicationDelegateObject) ApplicationShouldAutomaticallyLocalizeKeyEquivalents(application INSApplication) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("applicationShouldAutomaticallyLocalizeKeyEquivalents:"), application)
	return rv
	}

// Returns an error for the app to display to the user.
//
// application: The application object associated with the delegate.
//
// error: The error object that is used to construct the error message. Your
// implementation of this method can return a new [NSError] object or the same
// one in this parameter.
//
// # Return Value
// 
// The error object to display.
//
// # Discussion
// 
// You can implement this delegate method to customize the presentation of any
// error presented by your application, as long as no code in your application
// overrides either of the [NSResponder] methods `` or `` in a way that
// prevents errors from being passed down the responder chain to the
// application object.
// 
// Your implementation of this delegate method can examine `error` and, if its
// localized description or recovery information is unhelpfully generic,
// return an error object with specific localized text that is more suitable
// for presentation in alert sheets and dialogs. If you do this, always use
// the domain and error code of the [NSError] object to distinguish between
// errors whose presentation you want to customize and those you do not.
// Don’t make decisions based on the localized description, recovery
// suggestion, or recovery options because parsing localized text is
// problematic. If you decide not to customize the error presentation, just
// return the passed-in error object.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplicationDelegate/application(_:willPresentError:)

func (o NSApplicationDelegateObject) ApplicationWillPresentError(application INSApplication, error_ foundation.INSError) foundation.INSError {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("application:willPresentError:"), application, error_)
	return foundation.NSErrorFromID(rv)
	}

// Tells the delegate about changes to the configuration of any attached
// displays.
//
// notification: A notification named [didChangeScreenParametersNotification]. Calling the
// [object] method of this notification returns the [NSApplication] object
// itself.
// //
// [didChangeScreenParametersNotification]: https://developer.apple.com/documentation/AppKit/NSApplication/didChangeScreenParametersNotification
// [object]: https://developer.apple.com/documentation/Foundation/NSNotification/object
//
// See: https://developer.apple.com/documentation/AppKit/NSApplicationDelegate/applicationDidChangeScreenParameters(_:)

func (o NSApplicationDelegateObject) ApplicationDidChangeScreenParameters(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("applicationDidChangeScreenParameters:"), notification)
	}

// Returns a Boolean value that indicates if the app can continue the
// specified activity.
//
// application: The app continuing the user activity.
//
// userActivityType: The type of activity to be continued.
//
// # Return Value
// 
// [true] if you notify the user that your app is about to continue the
// activity or [false] if you want AppKit to notify the user.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use this method to provide immediate feedback to the user that an activity
// is about to continue on this device. The app calls this method as soon as
// the user confirms that an activity should be continued but possibly before
// the data associated with that activity is available.
// 
// This method is called on the main thread as soon as the user indicates they
// want to continue an activity in your app. The [NSUserActivity] object may
// not be available instantly, so use this as an opportunity to show the user
// that an activity will be continued shortly and return [true]. If you leave
// this method unimplemented or return [false], AppKit displays a default
// indication.
// 
// For each invocation of this method, the app delegate is guaranteed to get
// exactly one invocation of
// [ApplicationContinueUserActivityRestorationHandler] on success, or
// [ApplicationDidFailToContinueUserActivityWithTypeError] if an error was
// encountered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSApplicationDelegate/application(_:willContinueUserActivityWithType:)

func (o NSApplicationDelegateObject) ApplicationWillContinueUserActivityWithType(application INSApplication, userActivityType string) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("application:willContinueUserActivityWithType:"), application, objc.String(userActivityType))
	return rv
	}

// Returns a Boolean value that indicates if the app successfully recreates
// the specified activity.
//
// application: The app continuing the user activity.
//
// userActivity: The activity object containing the data associated with the task the user
// was performing. Use the data in this object to recreate what the user was
// doing.
//
// restorationHandler: A block to execute if your app creates or fetches objects to perform the
// task. Calling this block is optional and is only needed when specific
// objects are capable of continuing the activity. You can copy this block and
// call it at a later time. When calling a saved copy of the block, you must
// call it from the app’s main thread. This block has no return value and
// takes the following parameter:
// 
// `restorableObjects`: An array of [NSResponder] or [NSDocument] objects that
// you created or fetched in order to perform the operation. The system calls
// the [RestoreUserActivityState] method of each object in the array to
// perform the operation.
//
// # Return Value
// 
// [true] if this method handled continuing the activity; [false] to have
// AppKit attempt to continue the activity.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The app calls this method when it receives the data associated with the
// user activity. Use the data stored in the [NSUserActivity] object to
// re-create the user’s activity. This method is your opportunity to update
// your app so that it can perform the associated task.
// 
// If this user activity object was created automatically by having
// [NSUbiquitousDocumentUserActivityType] in a [CFBundleDocumentTypes] entry,
// AppKit can automatically restore the [NSUserActivity] in macOS if this
// method returns [false], or if it is unimplemented. It does this by creating
// a document of the appropriate type using the URL stored in the [userInfo]
// dictionary under the [NSUserActivityDocumentURLKey]. The system calls the
// [NSDocument] method [RestoreUserActivityState] on the document.
//
// [NSUserActivity]: https://developer.apple.com/documentation/Foundation/NSUserActivity
// [false]: https://developer.apple.com/documentation/Swift/false
// [userInfo]: https://developer.apple.com/documentation/Foundation/NSUserActivity/userInfo
//
// See: https://developer.apple.com/documentation/AppKit/NSApplicationDelegate/application(_:continue:restorationHandler:)

func (o NSApplicationDelegateObject) ApplicationContinueUserActivityRestorationHandler(application INSApplication, userActivity foundation.NSUserActivity, restorationHandler VoidHandler) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("application:continueUserActivity:restorationHandler:"), application, userActivity, restorationHandler)
	return rv
	}

// Tells the delegate that the app couldn’t continue the specified activity.
//
// application: The app that attempted to continue the activity.
//
// userActivityType: The activity type that was attempted.
//
// error: An error object indicating the reason for the failure.
//
// # Discussion
// 
// Use this method to let the user know that the specified activity could not
// be continued. If you do not implement this method, AppKit displays an error
// to the user with an appropriate message about the reason for the failure.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplicationDelegate/application(_:didFailToContinueUserActivityWithType:error:)

func (o NSApplicationDelegateObject) ApplicationDidFailToContinueUserActivityWithTypeError(application INSApplication, userActivityType string, error_ foundation.INSError) {
	
	objc.Send[struct{}](o.ID, objc.Sel("application:didFailToContinueUserActivityWithType:error:"), application, objc.String(userActivityType), error_)
	}

// Tells the delegate that there are changes to the specified activity.
//
// application: The shared app object.
//
// userActivity: The user activity object that was updated.
//
// # Discussion
// 
// This method is called when any user activity managed by AppKit has been
// updated. Use this as a last chance to add additional data to the user
// activity object.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplicationDelegate/application(_:didUpdate:)

func (o NSApplicationDelegateObject) ApplicationDidUpdateUserActivity(application INSApplication, userActivity foundation.NSUserActivity) {
	
	objc.Send[struct{}](o.ID, objc.Sel("application:didUpdateUserActivity:"), application, userActivity)
	}

// Tells the delegate that the app registered for Apple Push Services.
//
// application: The application that initiated the remote-notification registration
// process.
//
// deviceToken: A token that identifies the device to Apple Push Notification Service
// (APNS). The token is an opaque data type because that is the form that the
// provider needs to submit to the APNS servers when it sends a notification
// to a device. The APNS servers require a binary format for performance
// reasons.
// 
// The size of a device token is 32 bytes.
//
// # Discussion
// 
// The delegate receives this message after the
// [RegisterForRemoteNotificationTypes]method of [NSApplication] is invoked
// and there is no error in the registration process. After receiving the
// device token, the application should connect with its provider and give the
// token to it. APNS only pushes notifications to the application’s computer
// that are accompanied with this token.
// 
// For more information about how to implement push notifications in your
// application, see [Local and Remote Notification Programming Guide].
//
// [Local and Remote Notification Programming Guide]: https://developer.apple.com/library/archive/documentation/NetworkingInternet/Conceptual/RemoteNotificationsPG/index.html#//apple_ref/doc/uid/TP40008194
//
// See: https://developer.apple.com/documentation/AppKit/NSApplicationDelegate/application(_:didRegisterForRemoteNotificationsWithDeviceToken:)

func (o NSApplicationDelegateObject) ApplicationDidRegisterForRemoteNotificationsWithDeviceToken(application INSApplication, deviceToken foundation.INSData) {
	
	objc.Send[struct{}](o.ID, objc.Sel("application:didRegisterForRemoteNotificationsWithDeviceToken:"), application, deviceToken)
	}

// Tells the delegate that the app was unable to register for Apple Push
// Services.
//
// application: The application that initiated the remote-notification registration
// process.
//
// error: An NSError object that encapsulates information why registration did not
// succeed. The application can display this information to the user.
//
// # Discussion
// 
// The delegate receives this message after the
// [RegisterForRemoteNotificationTypes] method of [NSApplication] is invoked
// and there is an error in the registration process.
// 
// For more information about how to implement push notifications in your
// application, see [Local and Remote Notification Programming Guide].
//
// [Local and Remote Notification Programming Guide]: https://developer.apple.com/library/archive/documentation/NetworkingInternet/Conceptual/RemoteNotificationsPG/index.html#//apple_ref/doc/uid/TP40008194
//
// See: https://developer.apple.com/documentation/AppKit/NSApplicationDelegate/application(_:didFailToRegisterForRemoteNotificationsWithError:)

func (o NSApplicationDelegateObject) ApplicationDidFailToRegisterForRemoteNotificationsWithError(application INSApplication, error_ foundation.INSError) {
	
	objc.Send[struct{}](o.ID, objc.Sel("application:didFailToRegisterForRemoteNotificationsWithError:"), application, error_)
	}

// Tells the delegate when the app receives a remote notification.
//
// application: The application that received the remote notification.
//
// userInfo: A dictionary that contains information related to the remote notification,
// specifically a badge number for the application icon, a notification
// identifier, and possibly custom data. The provider originates it as a
// JSON-defined dictionary that AppKit converts to an [NSDictionary] object;
// the dictionary may contain only property-list objects plus [NSNull].
// //
// [NSDictionary]: https://developer.apple.com/documentation/Foundation/NSDictionary
// [NSNull]: https://developer.apple.com/documentation/Foundation/NSNull
//
// # Discussion
// 
// The delegate receives this message when the application is running and a
// remote notification arrives for it. In response, the application typically
// connects with its provider and downloads the data waiting for it. It may
// also process the notification in any other way it deems useful.
// 
// The `userInfo` dictionary contains another dictionary that you can obtain
// using the `aps` key. You can access the contents of the `aps` dictionary
// using the following keys:
// 
// - `badge`—A number indicating the quantity of data items to obtain from
// the provider. This number is to be displayed on the application icon. The
// absence of a `badge` property indicates that any number currently badging
// the icon should be removed.
// 
// Icon badging is the only notification type supported for non-running
// applications.
// 
// - `alert`—The value may either be a string for the alert message or a
// dictionary with two keys: `body` and `show-view`. The value of the former
// is the alert message and the latter is a Boolean (`false` or `true`). You
// may ignore the second key. - `sound`—The name of a sound file in the
// application bundle to play as an alert sound. If “default” is
// specified, the default sound should be played.
// 
// The `userInfo` dictionary may also have custom data defined by the provider
// according to the JSON schema. The properties for custom data should be
// specified at the same level as the `aps` dictionary. However,
// custom-defined properties should not be used for mass data transport
// because there is a strict size limit per notification (256 bytes) and
// delivery is not guaranteed.
// 
// If you implement [ApplicationDidFinishLaunching] and a push notification
// for the application has recently arrived, this method is not invoked for
// that push notification. In this case, you can access the JSON data in the
// [userInfo] dictionary of the passed-in [NSNotification] object.
// 
// For more information about how to implement push notifications in your
// application, see [Local and Remote Notification Programming Guide].
//
// [Local and Remote Notification Programming Guide]: https://developer.apple.com/library/archive/documentation/NetworkingInternet/Conceptual/RemoteNotificationsPG/index.html#//apple_ref/doc/uid/TP40008194
// [NSNotification]: https://developer.apple.com/documentation/Foundation/NSNotification
// [userInfo]: https://developer.apple.com/documentation/Foundation/NSNotification/userInfo
//
// See: https://developer.apple.com/documentation/AppKit/NSApplicationDelegate/application(_:didReceiveRemoteNotification:)

func (o NSApplicationDelegateObject) ApplicationDidReceiveRemoteNotification(application INSApplication, userInfo foundation.INSDictionary) {
	
	objc.Send[struct{}](o.ID, objc.Sel("application:didReceiveRemoteNotification:"), application, userInfo)
	}

// Tells the delegate when the user accepts a CloudKit sharing invitation.
//
// application: The shared app object.
//
// metadata: The metadata associated with the invitation. Use the URL of the
// metadata’s [CKShare] object and the [containerIdentifier] property to
// schedule a [CKAcceptSharesOperation] object.
// //
// [CKAcceptSharesOperation]: https://developer.apple.com/documentation/CloudKit/CKAcceptSharesOperation
// [CKShare]: https://developer.apple.com/documentation/CloudKit/CKShare
// [containerIdentifier]: https://developer.apple.com/documentation/CloudKit/CKShare/Metadata/containerIdentifier
//
// # Discussion
// 
// Use the provided metadata to begin sharing the specified content with the
// current user. For more information, see [CloudKit].
//
// [CloudKit]: https://developer.apple.com/documentation/CloudKit
//
// See: https://developer.apple.com/documentation/AppKit/NSApplicationDelegate/application(_:userDidAcceptCloudKitShareWith:)

func (o NSApplicationDelegateObject) ApplicationUserDidAcceptCloudKitShareWithMetadata(application INSApplication, metadata objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("application:userDidAcceptCloudKitShareWithMetadata:"), application, metadata)
	}

// Returns an intent handler that’s capable of handling the specified
// intent.
//
// application: The shared app object.
//
// intent: The intent object that represents the request coming from the system.
//
// # Return Value
// 
// An instance of a type capable of handling the specified intent, or `nil` if
// your app doesn’t handle the intent. Return an instance of a type that
// conforms to the protocol for handling intents of the same type as the
// provided `intent`.
//
// # Discussion
// 
// Siri invokes this method on the main queue when it wants to process one of
// your app’s supported intents. To indicate the intents that your app
// supports, populate the [INIntentsSupported] array in your app target’s
// `Info.Plist()` file.
// 
// In your delegate’s implementation of this method, check the `intent`
// parameter’s type and return a custom object that adopts the corresponding
// handling protocol. For example, if `intent` is an instance of
// [INPlayMediaIntent], return an object that adopts
// [INPlayMediaIntentHandling]. Only use the provided intent to determine
// which handler to return; don’t use it to initialize the handler and
// don’t store a reference to it. SiriKit updates the intent throughout the
// request to incorporate information the user provides. For more information,
// see [Dispatching intents to handlers].
// 
// For information about handling intents, see [Resolving and Handling
// Intents].
//
// [Dispatching intents to handlers]: https://developer.apple.com/documentation/SiriKit/dispatching-intents-to-handlers
// [INIntentsSupported]: https://developer.apple.com/documentation/BundleResources/Information-Property-List/INIntentsSupported
// [INPlayMediaIntentHandling]: https://developer.apple.com/documentation/Intents/INPlayMediaIntentHandling
// [INPlayMediaIntent]: https://developer.apple.com/documentation/Intents/INPlayMediaIntent
// [Resolving and Handling Intents]: https://developer.apple.com/documentation/SiriKit/resolving-and-handling-intents
//
// See: https://developer.apple.com/documentation/AppKit/NSApplicationDelegate/application(_:handlerFor:)

func (o NSApplicationDelegateObject) ApplicationHandlerForIntent(application INSApplication, intent objectivec.IObject) objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("application:handlerForIntent:"), application, intent)
	return objectivec.Object{ID: rv}
	}

// Tells the delegate to open the resource at the specified URL.
//
// application: Your singleton app object.
//
// urls: An array of URLs to open. The list does not include URLs for which your app
// has a defined document type.
//
// # Discussion
// 
// AppKit calls this method when your app is asked to open one or more
// URL-based resources. You must declare the URL types that your app supports
// in your `Info.Plist()` file using the [CFBundleURLTypes] key. The list can
// also include URLs for documents for which your app does not have an
// associated [NSDocument] class. You configure document types using Xcode, or
// by adding the [CFBundleDocumentTypes] key to your `Info.Plist()` file.
// 
// If your delegate implements this method, AppKit does not call the
// [ApplicationOpenFile] or [ApplicationOpenFiles] methods.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplicationDelegate/application(_:open:)

func (o NSApplicationDelegateObject) ApplicationOpenURLs(application INSApplication, urls []foundation.NSURL) {
	
	objc.Send[struct{}](o.ID, objc.Sel("application:openURLs:"), application, objectivec.IObjectSliceToNSArray(urls))
	}

// Returns a Boolean value that indicates if the app opens the specified file.
//
// sender: The application object associated with the delegate.
//
// filename: The name of the file to open.
//
// # Return Value
// 
// [true] if the file was successfully opened or [false] if it was not.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Sent directly by `theApplication` to the delegate. The method should open
// the file `filename`, returning [true] if the file is successfully opened,
// and [false] otherwise. If the user started up the application by
// double-clicking a file, the delegate receives the [ApplicationOpenFile]
// message before receiving [ApplicationDidFinishLaunching].
// ([ApplicationWillFinishLaunching] is sent before [ApplicationOpenFile].)
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSApplicationDelegate/application(_:openFile:)

func (o NSApplicationDelegateObject) ApplicationOpenFile(sender INSApplication, filename string) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("application:openFile:"), sender, objc.String(filename))
	return rv
	}

// Returns a Boolean value that indicates if the app opens the specified file
// without showing its user interface.
//
// sender: The object that sent the command.
//
// filename: The name of the file to open.
//
// # Return Value
// 
// [true] if the file was successfully opened or [false] if it was not.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Sent directly by `sender` to the delegate to request that the file
// `filename` be opened as a linked file. The method should open the file
// without bringing up its application’s user interface—that is, work with
// the file is under programmatic control of `sender`, rather than under
// keyboard control of the user.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplicationDelegate/application(_:openFileWithoutUI:)

func (o NSApplicationDelegateObject) ApplicationOpenFileWithoutUI(sender objectivec.IObject, filename string) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("application:openFileWithoutUI:"), sender, objc.String(filename))
	return rv
	}

// Returns a Boolean value that indicates if the app opens the specified
// temporary file.
//
// sender: The application object associated with the delegate.
//
// filename: The name of the temporary file to open.
//
// # Return Value
// 
// [true] if the file was successfully opened or [false] if it was not.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Sent directly by `theApplication` to the delegate. The method should
// attempt to open the file `filename`, returning [true] if the file is
// successfully opened, and [false] otherwise.
// 
// By design, a file opened through this method is assumed to be
// temporary—it’s the application’s responsibility to remove the file at
// the appropriate time.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSApplicationDelegate/application(_:openTempFile:)

func (o NSApplicationDelegateObject) ApplicationOpenTempFile(sender INSApplication, filename string) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("application:openTempFile:"), sender, objc.String(filename))
	return rv
	}

// Tells the delegate to open the specified files.
//
// sender: The application object associated with the delegate.
//
// filenames: An array of [NSString] objects containing the names of the files to open..
//
// # Discussion
// 
// Identical to [ApplicationOpenFile] except that the receiver opens multiple
// files corresponding to the file names in the `filenames` array. Delegates
// should invoke the [ReplyToOpenOrPrint] method upon success or failure, or
// when the user cancels the operation.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplicationDelegate/application(_:openFiles:)

func (o NSApplicationDelegateObject) ApplicationOpenFiles(sender INSApplication, filenames []string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("application:openFiles:"), sender, objectivec.StringSliceToNSArray(filenames))
	}

// Returns a Boolean value that indicates if the app can open an untitled
// file.
//
// sender: The application object associated with the delegate.
//
// # Return Value
// 
// [true] if the application should open a new untitled file or [false] if it
// should not.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use this method to decide whether the application should open a new,
// untitled file. Note that [ApplicationOpenUntitledFile] is invoked if this
// method returns [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSApplicationDelegate/applicationShouldOpenUntitledFile(_:)

func (o NSApplicationDelegateObject) ApplicationShouldOpenUntitledFile(sender INSApplication) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("applicationShouldOpenUntitledFile:"), sender)
	return rv
	}

// Returns a Boolean value that indicates if the app opens an untitled file.
//
// sender: The application object associated with the delegate.
//
// # Return Value
// 
// [true] if the file was successfully opened or [false] if it was not.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Sent directly by `theApplication` to the delegate to request that a new,
// untitled file be opened.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplicationDelegate/applicationOpenUntitledFile(_:)

func (o NSApplicationDelegateObject) ApplicationOpenUntitledFile(sender INSApplication) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("applicationOpenUntitledFile:"), sender)
	return rv
	}

// Returns a Boolean value that indicates if the app prints the specified file
// in its entirety.
//
// sender: The application object that is handling the printing.
//
// filename: The name of the file to print.
//
// # Return Value
// 
// [true] if the file was successfully printed or [false] if it was not.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This message is sent directly by `theApplication` to the delegate. The
// application terminates (using the [Terminate] method) after this method
// returns.
// 
// If at all possible, this method should print the file without displaying
// the user interface. For example, if you pass the `-NSPrint` option to the
// TextEdit application, TextEdit assumes you want to print the entire
// contents of the specified file. However, if the application opens more
// complex documents, you may want to display a panel that lets the user
// choose exactly what they want to print.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplicationDelegate/application(_:printFile:)

func (o NSApplicationDelegateObject) ApplicationPrintFile(sender INSApplication, filename string) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("application:printFile:"), sender, objc.String(filename))
	return rv
	}

// Returns a value that indicates if the app prints the specified files.
//
// application: The application object that is handling the printing.
//
// fileNames: An array of [NSString] objects, each of which contains the name of a file
// to print.
// //
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
//
// printSettings: A dictionary containing [NSPrintInfo]-compatible print job attributes.
//
// showPrintPanels: A Boolean that specifies whether the print panel should be displayed for
// each file printed. Print progress indicators will be presented even if this
// value is [false].
// //
// [false]: https://developer.apple.com/documentation/Swift/false
//
// # Return Value
// 
// A constant indicating whether printing was successful. For a list of
// possible values, see [NSApplication.PrintReply].
//
// [NSApplication.PrintReply]: https://developer.apple.com/documentation/AppKit/NSApplication/PrintReply
//
// # Discussion
// 
// Return [NSPrintingReplyLater] if the result of printing cannot be returned
// immediately, for example, if printing will cause the presentation of a
// sheet. If your method returns [NSPrintingReplyLater] it must always invoke
// the [NSApplication] method [ReplyToOpenOrPrint]] when the entire print
// operation has been completed, successfully or not.
// 
// This delegate method replaces ``, which is now deprecated. If your
// application delegate only implements the deprecated method, it is still
// invoked, and [NSApplication] uses private functionality to arrange for the
// print settings to take effect.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplicationDelegate/application(_:printFiles:withSettings:showPrintPanels:)

func (o NSApplicationDelegateObject) ApplicationPrintFilesWithSettingsShowPrintPanels(application INSApplication, fileNames []string, printSettings foundation.INSDictionary, showPrintPanels bool) NSApplicationPrintReply {
	
	rv := objc.Send[NSApplicationPrintReply](o.ID, objc.Sel("application:printFiles:withSettings:showPrintPanels:"), application, objectivec.StringSliceToNSArray(fileNames), printSettings, showPrintPanels)
	return rv
	}

// Returns a Boolean value that indicates if the app supports secure state
// restoration.
//
// app: The app object associated with the delegate.
//
// # Return Value
// 
// `true` when the app supports secure state restoration; otherwise, `false`.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplicationDelegate/applicationSupportsSecureRestorableState(_:)

func (o NSApplicationDelegateObject) ApplicationSupportsSecureRestorableState(app INSApplication) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("applicationSupportsSecureRestorableState:"), app)
	return rv
	}

// Tells the delegate that protected data is now available.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplicationDelegate/applicationProtectedDataDidBecomeAvailable(_:)

func (o NSApplicationDelegateObject) ApplicationProtectedDataDidBecomeAvailable(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("applicationProtectedDataDidBecomeAvailable:"), notification)
	}

// Tells the delegate that protected data is about to become unavailable.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplicationDelegate/applicationProtectedDataWillBecomeUnavailable(_:)

func (o NSApplicationDelegateObject) ApplicationProtectedDataWillBecomeUnavailable(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("applicationProtectedDataWillBecomeUnavailable:"), notification)
	}

// Tells the delegate that the app is about to encode its restorable state.
//
// app: The application.
//
// coder: The coder extracting the archive.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplicationDelegate/application(_:willEncodeRestorableState:)

func (o NSApplicationDelegateObject) ApplicationWillEncodeRestorableState(app INSApplication, coder foundation.INSCoder) {
	
	objc.Send[struct{}](o.ID, objc.Sel("application:willEncodeRestorableState:"), app, coder)
	}

// Tells the delegate when the app finished decoding its restorable state.
//
// app: The application.
//
// coder: The coder extracting the archive.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplicationDelegate/application(_:didDecodeRestorableState:)

func (o NSApplicationDelegateObject) ApplicationDidDecodeRestorableState(app INSApplication, coder foundation.INSCoder) {
	
	objc.Send[struct{}](o.ID, objc.Sel("application:didDecodeRestorableState:"), app, coder)
	}

// Tells the delegate about changes to the app’s occlusion state.
//
// notification: A notification named [didChangeOcclusionStateNotification]. Calling the
// [object] method of this notification returns the [NSApplication] object
// itself.
// //
// [didChangeOcclusionStateNotification]: https://developer.apple.com/documentation/AppKit/NSApplication/didChangeOcclusionStateNotification
// [object]: https://developer.apple.com/documentation/Foundation/NSNotification/object
//
// # Discussion
// 
// Upon receiving this method, you can query the application for its occlusion
// state. Note that this only notifies about changes in the state of the
// occlusion, not when the occlusion region changes. You can use this method
// to increase responsiveness and save power by halting any expensive
// calculations that the user can not see.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplicationDelegate/applicationDidChangeOcclusionState(_:)

func (o NSApplicationDelegateObject) ApplicationDidChangeOcclusionState(notification foundation.NSNotification) {
	
	objc.Send[struct{}](o.ID, objc.Sel("applicationDidChangeOcclusionState:"), notification)
	}

// Returns a Boolean value that indicates if the app supports the specified
// scripting key.
//
// sender: The app object associated with the delegate.
//
// key: The key to be handled.
//
// # Return Value
// 
// [true] if your delegate handles the key or [false] if it does not.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Implement this method and return [true] if your delegate handles the
// specified `key`, Your delegate must be able get or set the scriptable
// property or element that corresponds to that key. You need to handle only
// your app’s custom commands. The app already implements methods for each
// of the keys that it handles, where the method name matches the key.
// 
// For example, a scriptable app that doesn’t use Cocoa’s document-based
// app architecture can implement this method to supply its own document
// ordering. You want to do this because the standard app delegate expects to
// work with a document-based app. The TextEdit app (whose source is
// distributed with macOS developer tools) provides the following
// implementation:
// 
// TextEdit then implements the `orderedDocuments` method in its controller
// class to return an ordered list of documents. An app with its own window
// ordering might add a test for the key `orderedWindows` so that its delegate
// can provide its own version of `orderedWindows`.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSApplicationDelegate/application(_:delegateHandlesKey:)

func (o NSApplicationDelegateObject) ApplicationDelegateHandlesKey(sender INSApplication, key string) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("application:delegateHandlesKey:"), sender, objc.String(key))
	return rv
	}





// NSApplicationDelegateConfig holds optional typed callbacks for [NSApplicationDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate
type NSApplicationDelegateConfig struct {

	// Launching Applications
	// WillFinishLaunching — Tells the delegate that the app’s initialization is about to complete.
	WillFinishLaunching func(notification foundation.NSNotification)
	// DidFinishLaunching — Tells the delegate that the app’s initialization is complete but it hasn’t received its first event.
	DidFinishLaunching func(notification foundation.NSNotification)

	// Managing Active Status
	// WillBecomeActive — Tells the delegate that the app is about to become active.
	WillBecomeActive func(notification foundation.NSNotification)
	// DidBecomeActive — Tells the delegate that the app is now active.
	DidBecomeActive func(notification foundation.NSNotification)
	// WillResignActive — Tells the delegate that the app is about to become inactive and will lose focus.
	WillResignActive func(notification foundation.NSNotification)
	// DidResignActive — Tells the delegate that the app is no longer active and doesn’t have focus.
	DidResignActive func(notification foundation.NSNotification)

	// Terminating Applications
	// ShouldTerminate — Returns a value that indicates if the app should terminate.
	ShouldTerminate func(sender NSApplication) NSApplicationTerminateReply
	// ShouldTerminateAfterLastWindowClosed — Returns a Boolean value that indicates if the app terminates once the last window closes.
	ShouldTerminateAfterLastWindowClosed func(sender NSApplication) bool
	// WillTerminate — Tells the delegate that the app is about to terminate.
	WillTerminate func(notification foundation.NSNotification)

	// Hiding Applications
	// WillHide — Tells the delegate that the app is about to be hidden.
	WillHide func(notification foundation.NSNotification)
	// DidHide — Tells the delegate that the app is now hidden.
	DidHide func(notification foundation.NSNotification)
	// WillUnhide — Tells the delegate that the app is about to become visible.
	WillUnhide func(notification foundation.NSNotification)
	// DidUnhide — Tells the delegate that the app is now visible.
	DidUnhide func(notification foundation.NSNotification)

	// Managing Windows
	// WillUpdate — Tells the delegate that the app is about to update its windows.
	WillUpdate func(notification foundation.NSNotification)
	// DidUpdate — Tells the delegate that the app’s windows did update.
	DidUpdate func(notification foundation.NSNotification)
	// ShouldHandleReopenHasVisibleWindows — Returns a Boolean value that indicates if the app responds to reopen AppleEvents.
	ShouldHandleReopenHasVisibleWindows func(sender NSApplication, hasVisibleWindows bool) bool

	// Managing the Dock Menu
	// DockMenu — Returns the app’s dock menu.
	DockMenu func(sender NSApplication) NSMenu

	// Localizing Keyboard Shortcuts
	// ShouldAutomaticallyLocalizeKeyEquivalents — Returns a Boolean value that tells the system whether to remap menu shortcuts to support localized keyboards.
	ShouldAutomaticallyLocalizeKeyEquivalents func(application NSApplication) bool

	// Displaying Errors
	// WillPresentError — Returns an error for the app to display to the user.
	WillPresentError func(application NSApplication, error_ foundation.NSError) foundation.NSError

	// Managing the Screen
	// DidChangeScreenParameters — Tells the delegate about changes to the configuration of any attached displays.
	DidChangeScreenParameters func(notification foundation.NSNotification)

	// Handling Push Notifications
	// DidRegisterForRemoteNotificationsWithDeviceToken — Tells the delegate that the app registered for Apple Push Services.
	DidRegisterForRemoteNotificationsWithDeviceToken func(application NSApplication, deviceToken foundation.NSData)
	// DidFailToRegisterForRemoteNotificationsWithError — Tells the delegate that the app was unable to register for Apple Push Services.
	DidFailToRegisterForRemoteNotificationsWithError func(application NSApplication, error_ foundation.NSError)
	// DidReceiveRemoteNotification — Tells the delegate when the app receives a remote notification.
	DidReceiveRemoteNotification func(application NSApplication, userInfo foundation.INSDictionary)

	// Opening Files
	// ShouldOpenUntitledFile — Returns a Boolean value that indicates if the app can open an untitled file.
	ShouldOpenUntitledFile func(sender NSApplication) bool
	// OpenUntitledFile — Returns a Boolean value that indicates if the app opens an untitled file.
	OpenUntitledFile func(sender NSApplication) bool

	// Restoring Application State
	// SupportsSecureRestorableState — Returns a Boolean value that indicates if the app supports secure state restoration.
	SupportsSecureRestorableState func(app NSApplication) bool
	// ProtectedDataDidBecomeAvailable — Tells the delegate that protected data is now available.
	ProtectedDataDidBecomeAvailable func(notification foundation.NSNotification)
	// ProtectedDataWillBecomeUnavailable — Tells the delegate that protected data is about to become unavailable.
	ProtectedDataWillBecomeUnavailable func(notification foundation.NSNotification)
	// WillEncodeRestorableState — Tells the delegate that the app is about to encode its restorable state.
	WillEncodeRestorableState func(app NSApplication, coder foundation.INSCoder)
	// DidDecodeRestorableState — Tells the delegate when the app finished decoding its restorable state.
	DidDecodeRestorableState func(app NSApplication, coder foundation.INSCoder)

	// Handling Changes to the Occlusion State
	// DidChangeOcclusionState — Tells the delegate about changes to the app’s occlusion state.
	DidChangeOcclusionState func(notification foundation.NSNotification)

	// Other Methods
	// ContinueUserActivityRestorationHandler — Returns a Boolean value that indicates if the app successfully recreates the specified activity.
	ContinueUserActivityRestorationHandler func(application NSApplication, userActivity foundation.NSUserActivity, restorationHandler objc.ID) bool
	// DidUpdateUserActivity — Tells the delegate that there are changes to the specified activity.
	DidUpdateUserActivity func(application NSApplication, userActivity foundation.NSUserActivity)
}

// NewNSApplicationDelegate creates an Objective-C object implementing the [NSApplicationDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [NSApplicationDelegateObject] satisfies the [NSApplicationDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/appkit/nsapplicationdelegate
func NewNSApplicationDelegate(config NSApplicationDelegateConfig) NSApplicationDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoNSApplicationDelegate_%d", n)

	var methods []objc.MethodDef

	if config.WillFinishLaunching != nil {
		fn := config.WillFinishLaunching
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("applicationWillFinishLaunching:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.DidFinishLaunching != nil {
		fn := config.DidFinishLaunching
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("applicationDidFinishLaunching:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.WillBecomeActive != nil {
		fn := config.WillBecomeActive
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("applicationWillBecomeActive:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.DidBecomeActive != nil {
		fn := config.DidBecomeActive
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("applicationDidBecomeActive:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.WillResignActive != nil {
		fn := config.WillResignActive
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("applicationWillResignActive:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.DidResignActive != nil {
		fn := config.DidResignActive
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("applicationDidResignActive:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.ShouldTerminate != nil {
		fn := config.ShouldTerminate
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("applicationShouldTerminate:"),
			Fn: func(self objc.ID, _cmd objc.SEL, senderID objc.ID) NSApplicationTerminateReply {
				sender := NSApplicationFromID(senderID)
				return fn(sender)
			},
		})
	}

	if config.ShouldTerminateAfterLastWindowClosed != nil {
		fn := config.ShouldTerminateAfterLastWindowClosed
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("applicationShouldTerminateAfterLastWindowClosed:"),
			Fn: func(self objc.ID, _cmd objc.SEL, senderID objc.ID) bool {
				sender := NSApplicationFromID(senderID)
				return fn(sender)
			},
		})
	}

	if config.WillTerminate != nil {
		fn := config.WillTerminate
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("applicationWillTerminate:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.WillHide != nil {
		fn := config.WillHide
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("applicationWillHide:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.DidHide != nil {
		fn := config.DidHide
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("applicationDidHide:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.WillUnhide != nil {
		fn := config.WillUnhide
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("applicationWillUnhide:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.DidUnhide != nil {
		fn := config.DidUnhide
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("applicationDidUnhide:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.WillUpdate != nil {
		fn := config.WillUpdate
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("applicationWillUpdate:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.DidUpdate != nil {
		fn := config.DidUpdate
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("applicationDidUpdate:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.ShouldHandleReopenHasVisibleWindows != nil {
		fn := config.ShouldHandleReopenHasVisibleWindows
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("applicationShouldHandleReopen:hasVisibleWindows:"),
			Fn: func(self objc.ID, _cmd objc.SEL, senderID objc.ID, hasVisibleWindows bool) bool {
				sender := NSApplicationFromID(senderID)
				return fn(sender, hasVisibleWindows)
			},
		})
	}

	if config.DockMenu != nil {
		fn := config.DockMenu
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("applicationDockMenu:"),
			Fn: func(self objc.ID, _cmd objc.SEL, senderID objc.ID) objc.ID {
				sender := NSApplicationFromID(senderID)
				return fn(sender).GetID()
			},
		})
	}

	if config.ShouldAutomaticallyLocalizeKeyEquivalents != nil {
		fn := config.ShouldAutomaticallyLocalizeKeyEquivalents
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("applicationShouldAutomaticallyLocalizeKeyEquivalents:"),
			Fn: func(self objc.ID, _cmd objc.SEL, applicationID objc.ID) bool {
				application := NSApplicationFromID(applicationID)
				return fn(application)
			},
		})
	}

	if config.WillPresentError != nil {
		fn := config.WillPresentError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("application:willPresentError:"),
			Fn: func(self objc.ID, _cmd objc.SEL, applicationID objc.ID, error_ID objc.ID) objc.ID {
				application := NSApplicationFromID(applicationID)
				error_ := foundation.NSErrorFromID(error_ID)
				return fn(application, error_).GetID()
			},
		})
	}

	if config.DidChangeScreenParameters != nil {
		fn := config.DidChangeScreenParameters
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("applicationDidChangeScreenParameters:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.ContinueUserActivityRestorationHandler != nil {
		fn := config.ContinueUserActivityRestorationHandler
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("application:continueUserActivity:restorationHandler:"),
			Fn: func(self objc.ID, _cmd objc.SEL, applicationID objc.ID, userActivityID objc.ID, restorationHandlerID objc.ID) bool {
				application := NSApplicationFromID(applicationID)
				userActivity := foundation.NSUserActivityFromID(userActivityID)
				restorationHandler := restorationHandlerID
				return fn(application, userActivity, restorationHandler)
			},
		})
	}

	if config.DidUpdateUserActivity != nil {
		fn := config.DidUpdateUserActivity
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("application:didUpdateUserActivity:"),
			Fn: func(self objc.ID, _cmd objc.SEL, applicationID objc.ID, userActivityID objc.ID) {
				application := NSApplicationFromID(applicationID)
				userActivity := foundation.NSUserActivityFromID(userActivityID)
				fn(application, userActivity)
			},
		})
	}

	if config.DidRegisterForRemoteNotificationsWithDeviceToken != nil {
		fn := config.DidRegisterForRemoteNotificationsWithDeviceToken
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("application:didRegisterForRemoteNotificationsWithDeviceToken:"),
			Fn: func(self objc.ID, _cmd objc.SEL, applicationID objc.ID, deviceTokenID objc.ID) {
				application := NSApplicationFromID(applicationID)
				deviceToken := foundation.NSDataFromID(deviceTokenID)
				fn(application, deviceToken)
			},
		})
	}

	if config.DidFailToRegisterForRemoteNotificationsWithError != nil {
		fn := config.DidFailToRegisterForRemoteNotificationsWithError
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("application:didFailToRegisterForRemoteNotificationsWithError:"),
			Fn: func(self objc.ID, _cmd objc.SEL, applicationID objc.ID, error_ID objc.ID) {
				application := NSApplicationFromID(applicationID)
				error_ := foundation.NSErrorFromID(error_ID)
				fn(application, error_)
			},
		})
	}

	if config.DidReceiveRemoteNotification != nil {
		fn := config.DidReceiveRemoteNotification
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("application:didReceiveRemoteNotification:"),
			Fn: func(self objc.ID, _cmd objc.SEL, applicationID objc.ID, userInfoID objc.ID) {
				application := NSApplicationFromID(applicationID)
				userInfo := foundation.NSDictionaryFromID(userInfoID)
				fn(application, userInfo)
			},
		})
	}

	if config.ShouldOpenUntitledFile != nil {
		fn := config.ShouldOpenUntitledFile
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("applicationShouldOpenUntitledFile:"),
			Fn: func(self objc.ID, _cmd objc.SEL, senderID objc.ID) bool {
				sender := NSApplicationFromID(senderID)
				return fn(sender)
			},
		})
	}

	if config.OpenUntitledFile != nil {
		fn := config.OpenUntitledFile
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("applicationOpenUntitledFile:"),
			Fn: func(self objc.ID, _cmd objc.SEL, senderID objc.ID) bool {
				sender := NSApplicationFromID(senderID)
				return fn(sender)
			},
		})
	}

	if config.SupportsSecureRestorableState != nil {
		fn := config.SupportsSecureRestorableState
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("applicationSupportsSecureRestorableState:"),
			Fn: func(self objc.ID, _cmd objc.SEL, appID objc.ID) bool {
				app := NSApplicationFromID(appID)
				return fn(app)
			},
		})
	}

	if config.ProtectedDataDidBecomeAvailable != nil {
		fn := config.ProtectedDataDidBecomeAvailable
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("applicationProtectedDataDidBecomeAvailable:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.ProtectedDataWillBecomeUnavailable != nil {
		fn := config.ProtectedDataWillBecomeUnavailable
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("applicationProtectedDataWillBecomeUnavailable:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	if config.WillEncodeRestorableState != nil {
		fn := config.WillEncodeRestorableState
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("application:willEncodeRestorableState:"),
			Fn: func(self objc.ID, _cmd objc.SEL, appID objc.ID, coderID objc.ID) {
				app := NSApplicationFromID(appID)
				coder := foundation.NSCoderFromID(coderID)
				fn(app, coder)
			},
		})
	}

	if config.DidDecodeRestorableState != nil {
		fn := config.DidDecodeRestorableState
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("application:didDecodeRestorableState:"),
			Fn: func(self objc.ID, _cmd objc.SEL, appID objc.ID, coderID objc.ID) {
				app := NSApplicationFromID(appID)
				coder := foundation.NSCoderFromID(coderID)
				fn(app, coder)
			},
		})
	}

	if config.DidChangeOcclusionState != nil {
		fn := config.DidChangeOcclusionState
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("applicationDidChangeOcclusionState:"),
			Fn: func(self objc.ID, _cmd objc.SEL, notificationID objc.ID) {
				notification := foundation.NSNotificationFromID(notificationID)
				fn(notification)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("NSApplicationDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewNSApplicationDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return NSApplicationDelegateObjectFromID(instance)
}





