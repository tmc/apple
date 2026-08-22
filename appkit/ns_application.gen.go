// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSApplication] class.
var (
	_NSApplicationClass     NSApplicationClass
	_NSApplicationClassOnce sync.Once
)

func getNSApplicationClass() NSApplicationClass {
	_NSApplicationClassOnce.Do(func() {
		_NSApplicationClass = NSApplicationClass{class: objc.GetClass("NSApplication")}
	})
	return _NSApplicationClass
}

// GetNSApplicationClass returns the class object for NSApplication.
func GetNSApplicationClass() NSApplicationClass {
	return getNSApplicationClass()
}

type NSApplicationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSApplicationClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSApplicationClass) Alloc() NSApplication {
	rv := objc.Send[NSApplication](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that manages an app’s main event loop and resources used by all
// of that app’s objects.
//
// # Overview
//
// Every app uses a single instance of [NSApplication] to control the main
// event loop, keep track of the app’s windows and menus, distribute events
// to the appropriate objects (that’s, itself or one of its windows), set up
// autorelease pools, and receive notification of app-level events. An
// [NSApplication] object has a delegate (an object that you assign) that’s
// notified when the app starts or terminates, is hidden or activated, should
// open a file selected by the user, and so forth. By setting the delegate and
// implementing the delegate methods, you customize the behavior of your app
// without having to subclass [NSApplication]. In your app’s `main()`
// function, create the [NSApplication] instance by calling the
// [NSApplicationClass.SharedApplication] class method. After creating the
// application object, the `main()` function should load your app’s main nib
// file and then start the event loop by sending the application object a
// [NSApplication.Run] message. If you create an Application project in Xcode,
// this `main()` function is created for you. The `main()` function Xcode
// creates begins by calling a function named `NSApplicationMain()`, which is
// functionally similar to the following:
//
// The [NSApplicationClass.SharedApplication] class method initializes the
// display environment and connects your program to the window server and the
// display server. The [NSApplication] object maintains a list of all the
// [NSWindow] objects the app uses, so it can retrieve any of the app’s
// [NSView] objects. The [NSApplicationClass.SharedApplication] method also
// initializes the global variable [NSApp], which you use to retrieve the
// [NSApplication] instance. [NSApplicationClass.SharedApplication] only
// performs the initialization once. If you invoke it more than once, it
// returns the application object it created previously.
//
// The shared [NSApplication] object performs the important task of receiving
// events from the window server and distributing them to the proper
// [NSResponder] objects. [NSApp] translates an event into an [NSEvent]
// object, then forwards the event object to the affected [NSWindow] object.
// All keyboard and mouse events go directly to the [NSWindow] object
// associated with the event. The only exception to this rule is if the
// Command key is pressed when a key-down event occurs; in this case, every
// [NSWindow] object has an opportunity to respond to the event. When a window
// object receives an [NSEvent] object from [NSApp], it distributes it to the
// objects in its view hierarchy.
//
// [NSApplication] is also responsible for dispatching certain Apple events
// received by the app. For example, macOS sends Apple events to your app at
// various times, such as when the app is launched or reopened.
// [NSApplication] installs Apple event handlers to handle these events by
// sending a message to the appropriate object. You can also use the
// [NSAppleEventManager] class to register your own Apple event handlers. The
// [ApplicationWillFinishLaunching] method is generally the best place to do
// so. For more information on how events are handled and how you can modify
// the default behavior, including information on working with Apple events in
// scriptable apps, see [How Cocoa Applications Handle Apple Events] in [Cocoa
// Scripting Guide].
//
// The [NSApplication] class sets up `@autorelease` block during
// initialization and inside the event loop—specifically, within its
// initialization (or [NSApplicationClass.SharedApplication]) and
// [NSApplication.Run] methods. Similarly, the methods AppKit adds to [Bundle]
// employ `@autorelease` blocks during the loading of nib files. These
// `@autorelease` blocks aren’t accessible outside the scope of the
// respective [NSApplication] and [Bundle] methods. Typically, an app creates
// objects either while the event loop is running or by loading objects from
// nib files, so this lack of access usually isn’t a problem. However, if
// you do need to use Cocoa classes within the `main()` function itself (other
// than to load nib files or to instantiate [NSApplication]), you should
// create an `@autorelease` block to contain the code using the classes.
//
// # The delegate and notifications
//
// You can assign a delegate to your [NSApplication] object. The delegate
// responds to certain messages on behalf of the object. Some of these
// messages, such as [ApplicationOpenFile], ask the delegate to perform an
// action. Another message, [ApplicationShouldTerminate], lets the delegate
// determine whether the app should be allowed to quit. The [NSApplication]
// class sends these messages directly to its delegate.
//
// [NSApplication] also posts notifications to the app’s default
// notification center. Any object may register to receive one or more of the
// notifications posted by [NSApplication] by sending the message
// [addObserver(_:selector:name:object:)] to the default notification center
// (an instance of the [NSNotificationCenter] class). The delegate of
// [NSApplication] is automatically registered to receive these notifications
// if it implements certain delegate methods. For example, [NSApplication]
// posts notifications when it’s about to be done launching the app and when
// it’s done launching the app ([willFinishLaunchingNotification] and
// [didFinishLaunchingNotification]). The delegate has an opportunity to
// respond to these notifications by implementing the methods
// [ApplicationWillFinishLaunching] and [ApplicationDidFinishLaunching]. If
// the delegate wants to be informed of both events, it implements both
// methods. If it needs to know only when the app is finished launching, it
// implements only [ApplicationDidFinishLaunching].
//
// # System services
//
// [NSApplication] interacts with the system services architecture to provide
// services to your app through the Services menu.
//
// # Subclassing notes
//
// You rarely should find a real need to create a custom [NSApplication]
// subclass. Unlike some object-oriented libraries, Cocoa doesn’t require
// you to subclass [NSApplication] to customize app behavior. Instead it gives
// you many other ways to customize an app. This section discusses both some
// of the possible reasons to subclass [NSApplication] and some of the reasons
// not to subclass [NSApplication].
//
// To use a custom subclass of [NSApplication], send
// [NSApplicationClass.SharedApplication] to your subclass rather than
// directly to [NSApplication]. If you create your app in Xcode, you can
// accomplish this by setting your custom app class to be the principal class.
// In Xcode, double-click the app target in the Groups and Files list to open
// the Info window for the target. Then display the Properties pane of the
// window and replace “NSApplication” in the Principal Class field with
// the name of your custom class. The [NSApplicationMain] function sends
// [NSApplicationClass.SharedApplication] to the principal class to obtain the
// global app instance ([NSApp])—which in this case will be an instance of
// your custom subclass of [NSApplication].
//
// # Methods to override
//
// Generally, you subclass [NSApplication] to provide your own special
// responses to messages that are routinely sent to the global app object
// ([NSApp]). [NSApplication] doesn’t have primitive methods in the sense of
// methods that you must override in your subclass. Here are four methods that
// are possible candidates for overriding:
//
// - Override [NSApplication.Run] if you want the app to manage the main event
// loop differently than it does by default. (This a critical and complex
// task, however, that you should only attempt with good reason). - Override
// [NSApplication.SendEvent] if you want to change how events are dispatched
// or perform some special event processing. - Override
// [NSApplication.RequestUserAttention] if you want to modify how your app
// attracts the attention of the user (for example, offering an alternative to
// the bouncing app icon in the Dock). - Override
// [NSApplication.TargetForAction] to substitute another object for the target
// of an action message.
//
// # Special considerations
//
// The global app object uses `@autorelease` blocks in its [NSApplication.Run]
// method; if you override this method, you’ll need to create your own
// `@autorelease` blocks.
//
// Do not override [NSApplicationClass.SharedApplication]. The default
// implementation, which is essential to app behavior, is too complex to
// duplicate on your own.
//
// # Alternatives to subclassing
//
// [NSApplication] defines numerous [Delegation] methods that offer
// opportunities for modifying specific aspects of app behavior. Instead of
// making a custom subclass of [NSApplication], your app delegate may be able
// to implement one or more of these methods to accomplish your design goals.
// In general, a better design than subclassing [NSApplication] is to put the
// code that expresses your app’s special behavior into one or more custom
// objects called controllers. Methods defined in your controllers can be
// invoked from a small dispatcher object without being closely tied to the
// global app object.
//
// # Managing the app’s behavior
//
//   - [NSApplication.Delegate]: The app delegate object.
//   - [NSApplication.SetDelegate]
//
// # Managing the event loop
//
//   - [NSApplication.NextEventMatchingMaskUntilDateInModeDequeue]: Returns the next event matching a given mask, or `nil` if no such event is found before a specified expiration date.
//   - [NSApplication.DiscardEventsMatchingMaskBeforeEvent]: Removes all events matching the given mask and generated before the specified event.
//   - [NSApplication.CurrentEvent]: The last event object that the app retrieved from the event queue.
//   - [NSApplication.IsRunning]: A Boolean value indicating whether the main event loop is running.
//   - [NSApplication.Run]: Starts the main event loop.
//   - [NSApplication.FinishLaunching]: Activates the app, opens any files specified by the [NSOpen] user default, and unhighlights the app’s icon.
//   - [NSApplication.Stop]: Stops the main event loop.
//   - [NSApplication.SendEvent]: Dispatches an event to other objects.
//   - [NSApplication.PostEventAtStart]: Adds a given event to the receiver’s event queue.
//
// # Posting actions
//
//   - [NSApplication.SendActionToFrom]: Sends the given action message to the given target.
//   - [NSApplication.TargetForAction]: Returns the object that receives the action message specified by the given selector.
//   - [NSApplication.TargetForActionToFrom]: Searches for an object that can receive the message specified by the given selector.
//
// # Terminating the app
//
//   - [NSApplication.Terminate]: Terminates the receiver.
//   - [NSApplication.ReplyToApplicationShouldTerminate]: Responds to [NSTerminateLater] once the app knows whether it can terminate.
//
// # Activating and deactivating the app
//
//   - [NSApplication.Activate]: Activates the receiver app, if appropriate.
//   - [NSApplication.Deactivate]: Deactivates the receiver.
//   - [NSApplication.IsActive]: A Boolean value indicating whether this is the active app.
//   - [NSApplication.YieldActivationToApplication]: Explicitly allows another app to make itself active.
//   - [NSApplication.YieldActivationToApplicationWithBundleIdentifier]: Explicitly allows another app to make itself active.
//
// # Managing relaunch on login
//
//   - [NSApplication.DisableRelaunchOnLogin]: Disables relaunching the app on login.
//   - [NSApplication.EnableRelaunchOnLogin]: Enables relaunching the app on login.
//
// # Managing remote notifications
//
//   - [NSApplication.RegisterForRemoteNotifications]: Register for notifications sent by Apple Push Notification service (APNs).
//   - [NSApplication.UnregisterForRemoteNotifications]: Unregister for notifications received from Apple Push Notification service.
//   - [NSApplication.EnabledRemoteNotificationTypes]: The types of push notifications that the app accepts.
//   - [NSApplication.RegisterForRemoteNotificationTypes]: Register to receive notifications of the specified types from a provider through the Apple Push Notification service.
//   - [NSApplication.IsRegisteredForRemoteNotifications]: A Boolean value indicating whether the app is registered with Apple Push Notification service (APNs).
//
// # Managing the app’s appearance
//
//   - [NSApplication.CurrentSystemPresentationOptions]: The set of app presentation options that are currently in effect for the system.
//   - [NSApplication.PresentationOptions]: The presentation options that should be in effect for the system when this app is active.
//   - [NSApplication.SetPresentationOptions]
//   - [NSApplication.ApplicationShouldSuppressHighDynamicRangeContent]: A boolean value indicating whether your application should suppress HDR content based on established policy. Built-in AppKit components such as NSImageView will automatically behave correctly with HDR content. You should use this value in conjunction with notifications ([NSApplicationShouldBeginSuppressingHighDynamicRangeContentNotification] and [NSApplicationShouldEndSuppressingHighDynamicRangeContentNotification]) to suppress HDR content in your application when signaled to do so.
//
// # User interface layout direction
//
//   - [NSApplication.UserInterfaceLayoutDirection]: The layout direction of the user interface.
//
// # Accessing the dock tile
//
//   - [NSApplication.DockTile]: The app’s Dock tile.
//   - [NSApplication.ApplicationIconImage]: The image used for the app’s icon.
//   - [NSApplication.SetApplicationIconImage]
//
// # Customizing the Touch Bar
//
//   - [NSApplication.ToggleTouchBarCustomizationPalette]: Show or hides the interface for customizing the Touch Bar.
//
// # Managing user attention requests
//
//   - [NSApplication.RequestUserAttention]: Starts a user attention request.
//   - [NSApplication.CancelUserAttentionRequest]: Cancels a previous user attention request.
//   - [NSApplication.ReplyToOpenOrPrint]: Handles errors that might occur when the user attempts to open or print files.
//
// # Providing help information
//
//   - [NSApplication.RegisterUserInterfaceItemSearchHandler]: Register an object that provides help data to your app.
//   - [NSApplication.SearchStringInUserInterfaceItemStringSearchRangeFoundRange]: Searches for the string in the user interface.
//   - [NSApplication.UnregisterUserInterfaceItemSearchHandler]: Unregister an object that provides help data to your app.
//   - [NSApplication.ShowHelp]: If your project is properly registered, and the necessary keys have been set in the property list, this method launches Help Viewer and displays the first page of your app’s help book.
//   - [NSApplication.ActivateContextHelpMode]: Places the receiver in context-sensitive help mode.
//   - [NSApplication.HelpMenu]: The help menu used by the app.
//   - [NSApplication.SetHelpMenu]
//
// # Providing services
//
//   - [NSApplication.ServicesProvider]: The object that provides the services the current app advertises in the Services menu of other apps.
//   - [NSApplication.SetServicesProvider]
//
// # Determining access to the keyboard
//
//   - [NSApplication.IsFullKeyboardAccessEnabled]: A Boolean value indicating whether Full Keyboard Access is enabled in the Keyboard preference pane.
//
// # Hiding apps
//
//   - [NSApplication.HideOtherApplications]: Hides all apps, except the receiver.
//   - [NSApplication.UnhideAllApplications]: Unhides all apps, including the receiver.
//
// # Logging exceptions
//
//   - [NSApplication.ReportException]: Logs a given exception by calling `NSLog()`.
//
// # Configuring the activation policy
//
//   - [NSApplication.ActivationPolicy]: Returns the app’s activation policy.
//   - [NSApplication.SetActivationPolicy]: Attempts to modify the app’s activation policy.
//
// # Scripting your app
//
//   - [NSApplication.OrderedDocuments]: An array of document objects arranged according to the front-to-back ordering of their associated windows.
//   - [NSApplication.OrderedWindows]: An array of window objects arranged according to their front-to-back ordering on the screen.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication
//
// [Bundle]: https://developer.apple.com/documentation/Foundation/Bundle
// [Cocoa Scripting Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ScriptableCocoaApplications/SApps_intro/SAppsIntro.html#//apple_ref/doc/uid/TP40002164
// [Delegation]: https://developer.apple.com/library/archive/documentation/General/Conceptual/DevPedia-CocoaCore/Delegation.html#//apple_ref/doc/uid/TP40008195-CH14
// [How Cocoa Applications Handle Apple Events]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ScriptableCocoaApplications/SApps_handle_AEs/SAppsHandleAEs.html#//apple_ref/doc/uid/20001239
// [NSAppleEventManager]: https://developer.apple.com/documentation/Foundation/NSAppleEventManager
// [addObserver(_:selector:name:object:)]: https://developer.apple.com/documentation/Foundation/NotificationCenter/addObserver(_:selector:name:object:)
// [didFinishLaunchingNotification]: https://developer.apple.com/documentation/AppKit/NSApplication/didFinishLaunchingNotification
// [willFinishLaunchingNotification]: https://developer.apple.com/documentation/AppKit/NSApplication/willFinishLaunchingNotification
type NSApplication struct {
	NSResponder
}

// NSApplicationFromID constructs a [NSApplication] from an objc.ID.
//
// An object that manages an app’s main event loop and resources used by all
// of that app’s objects.
func NSApplicationFromID(id objc.ID) NSApplication {
	return NSApplication{NSResponder: NSResponderFromID(id)}
}

// NOTE: NSApplication adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSApplication] class.
//
// # Managing the app’s behavior
//
//   - [INSApplication.Delegate]: The app delegate object.
//   - [INSApplication.SetDelegate]
//
// # Managing the event loop
//
//   - [INSApplication.NextEventMatchingMaskUntilDateInModeDequeue]: Returns the next event matching a given mask, or `nil` if no such event is found before a specified expiration date.
//   - [INSApplication.DiscardEventsMatchingMaskBeforeEvent]: Removes all events matching the given mask and generated before the specified event.
//   - [INSApplication.CurrentEvent]: The last event object that the app retrieved from the event queue.
//   - [INSApplication.IsRunning]: A Boolean value indicating whether the main event loop is running.
//   - [INSApplication.Run]: Starts the main event loop.
//   - [INSApplication.FinishLaunching]: Activates the app, opens any files specified by the [NSOpen] user default, and unhighlights the app’s icon.
//   - [INSApplication.Stop]: Stops the main event loop.
//   - [INSApplication.SendEvent]: Dispatches an event to other objects.
//   - [INSApplication.PostEventAtStart]: Adds a given event to the receiver’s event queue.
//
// # Posting actions
//
//   - [INSApplication.SendActionToFrom]: Sends the given action message to the given target.
//   - [INSApplication.TargetForAction]: Returns the object that receives the action message specified by the given selector.
//   - [INSApplication.TargetForActionToFrom]: Searches for an object that can receive the message specified by the given selector.
//
// # Terminating the app
//
//   - [INSApplication.Terminate]: Terminates the receiver.
//   - [INSApplication.ReplyToApplicationShouldTerminate]: Responds to [NSTerminateLater] once the app knows whether it can terminate.
//
// # Activating and deactivating the app
//
//   - [INSApplication.Activate]: Activates the receiver app, if appropriate.
//   - [INSApplication.Deactivate]: Deactivates the receiver.
//   - [INSApplication.IsActive]: A Boolean value indicating whether this is the active app.
//   - [INSApplication.YieldActivationToApplication]: Explicitly allows another app to make itself active.
//   - [INSApplication.YieldActivationToApplicationWithBundleIdentifier]: Explicitly allows another app to make itself active.
//
// # Managing relaunch on login
//
//   - [INSApplication.DisableRelaunchOnLogin]: Disables relaunching the app on login.
//   - [INSApplication.EnableRelaunchOnLogin]: Enables relaunching the app on login.
//
// # Managing remote notifications
//
//   - [INSApplication.RegisterForRemoteNotifications]: Register for notifications sent by Apple Push Notification service (APNs).
//   - [INSApplication.UnregisterForRemoteNotifications]: Unregister for notifications received from Apple Push Notification service.
//   - [INSApplication.EnabledRemoteNotificationTypes]: The types of push notifications that the app accepts.
//   - [INSApplication.RegisterForRemoteNotificationTypes]: Register to receive notifications of the specified types from a provider through the Apple Push Notification service.
//   - [INSApplication.IsRegisteredForRemoteNotifications]: A Boolean value indicating whether the app is registered with Apple Push Notification service (APNs).
//
// # Managing the app’s appearance
//
//   - [INSApplication.CurrentSystemPresentationOptions]: The set of app presentation options that are currently in effect for the system.
//   - [INSApplication.PresentationOptions]: The presentation options that should be in effect for the system when this app is active.
//   - [INSApplication.SetPresentationOptions]
//   - [INSApplication.ApplicationShouldSuppressHighDynamicRangeContent]: A boolean value indicating whether your application should suppress HDR content based on established policy. Built-in AppKit components such as NSImageView will automatically behave correctly with HDR content. You should use this value in conjunction with notifications ([NSApplicationShouldBeginSuppressingHighDynamicRangeContentNotification] and [NSApplicationShouldEndSuppressingHighDynamicRangeContentNotification]) to suppress HDR content in your application when signaled to do so.
//
// # User interface layout direction
//
//   - [INSApplication.UserInterfaceLayoutDirection]: The layout direction of the user interface.
//
// # Accessing the dock tile
//
//   - [INSApplication.DockTile]: The app’s Dock tile.
//   - [INSApplication.ApplicationIconImage]: The image used for the app’s icon.
//   - [INSApplication.SetApplicationIconImage]
//
// # Customizing the Touch Bar
//
//   - [INSApplication.ToggleTouchBarCustomizationPalette]: Show or hides the interface for customizing the Touch Bar.
//
// # Managing user attention requests
//
//   - [INSApplication.RequestUserAttention]: Starts a user attention request.
//   - [INSApplication.CancelUserAttentionRequest]: Cancels a previous user attention request.
//   - [INSApplication.ReplyToOpenOrPrint]: Handles errors that might occur when the user attempts to open or print files.
//
// # Providing help information
//
//   - [INSApplication.RegisterUserInterfaceItemSearchHandler]: Register an object that provides help data to your app.
//   - [INSApplication.SearchStringInUserInterfaceItemStringSearchRangeFoundRange]: Searches for the string in the user interface.
//   - [INSApplication.UnregisterUserInterfaceItemSearchHandler]: Unregister an object that provides help data to your app.
//   - [INSApplication.ShowHelp]: If your project is properly registered, and the necessary keys have been set in the property list, this method launches Help Viewer and displays the first page of your app’s help book.
//   - [INSApplication.ActivateContextHelpMode]: Places the receiver in context-sensitive help mode.
//   - [INSApplication.HelpMenu]: The help menu used by the app.
//   - [INSApplication.SetHelpMenu]
//
// # Providing services
//
//   - [INSApplication.ServicesProvider]: The object that provides the services the current app advertises in the Services menu of other apps.
//   - [INSApplication.SetServicesProvider]
//
// # Determining access to the keyboard
//
//   - [INSApplication.IsFullKeyboardAccessEnabled]: A Boolean value indicating whether Full Keyboard Access is enabled in the Keyboard preference pane.
//
// # Hiding apps
//
//   - [INSApplication.HideOtherApplications]: Hides all apps, except the receiver.
//   - [INSApplication.UnhideAllApplications]: Unhides all apps, including the receiver.
//
// # Logging exceptions
//
//   - [INSApplication.ReportException]: Logs a given exception by calling `NSLog()`.
//
// # Configuring the activation policy
//
//   - [INSApplication.ActivationPolicy]: Returns the app’s activation policy.
//   - [INSApplication.SetActivationPolicy]: Attempts to modify the app’s activation policy.
//
// # Scripting your app
//
//   - [INSApplication.OrderedDocuments]: An array of document objects arranged according to the front-to-back ordering of their associated windows.
//   - [INSApplication.OrderedWindows]: An array of window objects arranged according to their front-to-back ordering on the screen.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication
type INSApplication interface {
	INSResponder
	NSAppearanceCustomization

	// Topic: Managing the app’s behavior

	// The app delegate object.
	Delegate() NSApplicationDelegate
	SetDelegate(value NSApplicationDelegate)

	// Topic: Managing the event loop

	// Returns the next event matching a given mask, or `nil` if no such event is found before a specified expiration date.
	NextEventMatchingMaskUntilDateInModeDequeue(mask NSEventMask, expiration foundation.NSDate, mode foundation.NSRunLoopMode, deqFlag bool) INSEvent
	// Removes all events matching the given mask and generated before the specified event.
	DiscardEventsMatchingMaskBeforeEvent(mask NSEventMask, lastEvent INSEvent)
	// The last event object that the app retrieved from the event queue.
	CurrentEvent() INSEvent
	// A Boolean value indicating whether the main event loop is running.
	IsRunning() bool
	// Starts the main event loop.
	Run()
	// Activates the app, opens any files specified by the [NSOpen] user default, and unhighlights the app’s icon.
	FinishLaunching()
	// Stops the main event loop.
	Stop(sender objectivec.IObject)
	// Dispatches an event to other objects.
	SendEvent(event INSEvent)
	// Adds a given event to the receiver’s event queue.
	PostEventAtStart(event INSEvent, atStart bool)

	// Topic: Posting actions

	// Sends the given action message to the given target.
	SendActionToFrom(action objc.SEL, target objectivec.IObject, sender objectivec.IObject) bool
	// Returns the object that receives the action message specified by the given selector.
	TargetForAction(action objc.SEL) objectivec.IObject
	// Searches for an object that can receive the message specified by the given selector.
	TargetForActionToFrom(action objc.SEL, target objectivec.IObject, sender objectivec.IObject) objectivec.IObject

	// Topic: Terminating the app

	// Terminates the receiver.
	Terminate(sender objectivec.IObject)
	// Responds to [NSTerminateLater] once the app knows whether it can terminate.
	ReplyToApplicationShouldTerminate(shouldTerminate bool)

	// Topic: Activating and deactivating the app

	// Activates the receiver app, if appropriate.
	Activate()
	// Deactivates the receiver.
	Deactivate()
	// A Boolean value indicating whether this is the active app.
	IsActive() bool
	// Explicitly allows another app to make itself active.
	YieldActivationToApplication(application INSRunningApplication)
	// Explicitly allows another app to make itself active.
	YieldActivationToApplicationWithBundleIdentifier(bundleIdentifier string)

	// Topic: Managing relaunch on login

	// Disables relaunching the app on login.
	DisableRelaunchOnLogin()
	// Enables relaunching the app on login.
	EnableRelaunchOnLogin()

	// Topic: Managing remote notifications

	// Register for notifications sent by Apple Push Notification service (APNs).
	RegisterForRemoteNotifications()
	// Unregister for notifications received from Apple Push Notification service.
	UnregisterForRemoteNotifications()
	// The types of push notifications that the app accepts.
	EnabledRemoteNotificationTypes() NSRemoteNotificationType
	// Register to receive notifications of the specified types from a provider through the Apple Push Notification service.
	RegisterForRemoteNotificationTypes(types NSRemoteNotificationType)
	// A Boolean value indicating whether the app is registered with Apple Push Notification service (APNs).
	IsRegisteredForRemoteNotifications() bool

	// Topic: Managing the app’s appearance

	// The set of app presentation options that are currently in effect for the system.
	CurrentSystemPresentationOptions() NSApplicationPresentationOptions
	// The presentation options that should be in effect for the system when this app is active.
	PresentationOptions() NSApplicationPresentationOptions
	SetPresentationOptions(value NSApplicationPresentationOptions)
	// A boolean value indicating whether your application should suppress HDR content based on established policy. Built-in AppKit components such as NSImageView will automatically behave correctly with HDR content. You should use this value in conjunction with notifications ([NSApplicationShouldBeginSuppressingHighDynamicRangeContentNotification] and [NSApplicationShouldEndSuppressingHighDynamicRangeContentNotification]) to suppress HDR content in your application when signaled to do so.
	ApplicationShouldSuppressHighDynamicRangeContent() bool

	// Topic: User interface layout direction

	// The layout direction of the user interface.
	UserInterfaceLayoutDirection() NSUserInterfaceLayoutDirection

	// Topic: Accessing the dock tile

	// The app’s Dock tile.
	DockTile() INSDockTile
	// The image used for the app’s icon.
	ApplicationIconImage() INSImage
	SetApplicationIconImage(value INSImage)

	// Topic: Customizing the Touch Bar

	// Show or hides the interface for customizing the Touch Bar.
	ToggleTouchBarCustomizationPalette(sender objectivec.IObject)

	// Topic: Managing user attention requests

	// Starts a user attention request.
	RequestUserAttention(requestType NSRequestUserAttentionType) int
	// Cancels a previous user attention request.
	CancelUserAttentionRequest(request int)
	// Handles errors that might occur when the user attempts to open or print files.
	ReplyToOpenOrPrint(reply NSApplicationDelegateReply)

	// Topic: Providing help information

	// Register an object that provides help data to your app.
	RegisterUserInterfaceItemSearchHandler(handler NSUserInterfaceItemSearching)
	// Searches for the string in the user interface.
	SearchStringInUserInterfaceItemStringSearchRangeFoundRange(searchString string, stringToSearch string, searchRange foundation.NSRange, foundRange *foundation.NSRange) bool
	// Unregister an object that provides help data to your app.
	UnregisterUserInterfaceItemSearchHandler(handler NSUserInterfaceItemSearching)
	// If your project is properly registered, and the necessary keys have been set in the property list, this method launches Help Viewer and displays the first page of your app’s help book.
	ShowHelp(sender objectivec.IObject)
	// Places the receiver in context-sensitive help mode.
	ActivateContextHelpMode(sender objectivec.IObject)
	// The help menu used by the app.
	HelpMenu() INSMenu
	SetHelpMenu(value INSMenu)

	// Topic: Providing services

	// The object that provides the services the current app advertises in the Services menu of other apps.
	ServicesProvider() objectivec.IObject
	SetServicesProvider(value objectivec.IObject)

	// Topic: Determining access to the keyboard

	// A Boolean value indicating whether Full Keyboard Access is enabled in the Keyboard preference pane.
	IsFullKeyboardAccessEnabled() bool

	// Topic: Hiding apps

	// Hides all apps, except the receiver.
	HideOtherApplications(sender objectivec.IObject)
	// Unhides all apps, including the receiver.
	UnhideAllApplications(sender objectivec.IObject)

	// Topic: Logging exceptions

	// Logs a given exception by calling `NSLog()`.
	ReportException(exception foundation.NSException)

	// Topic: Configuring the activation policy

	// Returns the app’s activation policy.
	ActivationPolicy() NSApplicationActivationPolicy
	// Attempts to modify the app’s activation policy.
	SetActivationPolicy(activationPolicy NSApplicationActivationPolicy) bool

	// Topic: Scripting your app

	// An array of document objects arranged according to the front-to-back ordering of their associated windows.
	OrderedDocuments() []NSDocument
	// An array of window objects arranged according to their front-to-back ordering on the screen.
	OrderedWindows() []NSWindow

	// A Boolean value indicating whether the main menu contains an item for customizing the contents of the Touch Bar.
	IsAutomaticCustomizeTouchBarMenuItemEnabled() bool
	SetAutomaticCustomizeTouchBarMenuItemEnabled(value bool)
	// A Boolean value indicating whether the app is hidden.
	IsHidden() bool
	// The window that currently receives keyboard events.
	KeyWindow() INSWindow
	// The app’s main menu bar.
	MainMenu() INSMenu
	SetMainMenu(value INSMenu)
	// The app’s main window.
	MainWindow() INSWindow
	// The modal window displayed by the app.
	ModalWindow() INSWindow
	// The occlusion state of the app.
	OcclusionState() NSApplicationOcclusionState
	IsProtectedDataAvailable() bool
	// The app’s Services menu.
	ServicesMenu() INSMenu
	SetServicesMenu(value INSMenu)
	// An array of the app’s window objects.
	Windows() []NSWindow
	// The Window menu of the app.
	WindowsMenu() INSMenu
	SetWindowsMenu(value INSMenu)
	// Aborts the event loop started by [runModal(for:)](<https://developer.apple.com/documentation/AppKit/NSApplication/runModal(for:)>) or [runModalSession(_:)](<https://developer.apple.com/documentation/AppKit/NSApplication/runModalSession(_:)>).
	AbortModal()
	// Returns the activation point for the user interface element.
	AccessibilityActivationPoint() corefoundation.CGPoint
	// Returns the allowed values for the slider accessibility element.
	AccessibilityAllowedValues() []foundation.NSNumber
	// Returns the child accessibility element with the current focus.
	AccessibilityApplicationFocusedUIElement() objectivec.IObject
	// Returns the attributed substring for the specified range of characters.
	AccessibilityAttributedStringForRange(range_ foundation.NSRange) foundation.NSAttributedString
	AccessibilityAttributedUserInputLabels() []foundation.NSAttributedString
	// Returns the child accessibility element that represents the window’s cancel button.
	AccessibilityCancelButton() objectivec.IObject
	// Returns the cell at the specified column and row.
	AccessibilityCellForColumnRow(column int, row int) objectivec.IObject
	// Returns the child accessibility elements in the accessibility hierarchy.
	AccessibilityChildren() foundation.INSArray
	// Returns the array of child accessibility elements in order for linear navigation.
	AccessibilityChildrenInNavigationOrder() []objectivec.IObject
	// Returns the clear button for the search field.
	AccessibilityClearButton() objectivec.IObject
	// Returns the child accessibility element that represents the window’s close button.
	AccessibilityCloseButton() objectivec.IObject
	// Returns the number of columns in the accessibility element’s grid.
	AccessibilityColumnCount() int
	// Returns the column header accessibility elements for the table or outline.
	AccessibilityColumnHeaderUIElements() foundation.INSArray
	// Returns the column index range of the cell.
	AccessibilityColumnIndexRange() foundation.NSRange
	// Returns the column titles for the accessibility element.
	AccessibilityColumnTitles() foundation.INSArray
	// Returns the column accessibility elements for the table or outline.
	AccessibilityColumns() foundation.INSArray
	// Returns the contents of the current accessibility element.
	AccessibilityContents() foundation.INSArray
	// Returns the critical value for the level indicator.
	AccessibilityCriticalValue() objectivec.IObject
	// Returns the custom actions of the current accessibility element.
	AccessibilityCustomActions() []NSAccessibilityCustomAction
	// Returns the custom rotors of the current accessibility element.
	AccessibilityCustomRotors() []NSAccessibilityCustomRotor
	// Returns the decrement button for the stepper accessibility element.
	AccessibilityDecrementButton() objectivec.IObject
	// Returns the child accessibility element that represents the window’s default button.
	AccessibilityDefaultButton() objectivec.IObject
	// Returns the row disclosing the current row.
	AccessibilityDisclosedByRow() objectivec.IObject
	// Returns the rows that the current row discloses.
	AccessibilityDisclosedRows() objectivec.IObject
	// Returns the indention level for the row.
	AccessibilityDisclosureLevel() int
	// Returns the URL for the file that the accessibility element represents.
	AccessibilityDocument() string
	// Returns the icon for the app’s menu bar extra.
	AccessibilityExtrasMenuBar() objectivec.IObject
	// Returns the filename for the file that the accessibility element represents.
	AccessibilityFilename() string
	// Returns the child window with the current focus.
	AccessibilityFocusedWindow() objectivec.IObject
	// Returns the accessibility element’s frame in screen coordinates.
	AccessibilityFrame() corefoundation.CGRect
	// Returns the rectangle that encloses the specified range of characters.
	AccessibilityFrameForRange(range_ foundation.NSRange) corefoundation.CGRect
	// Returns the child accessibility element that represents the window’s full-screen button.
	AccessibilityFullScreenButton() objectivec.IObject
	// Returns the child accessibility element that represents the window’s grow area.
	AccessibilityGrowArea() objectivec.IObject
	// Returns the drag handle elements for the layout item element.
	AccessibilityHandles() foundation.INSArray
	// Returns the header for the table view.
	AccessibilityHeader() objectivec.IObject
	// Returns the help text for the accessibility element.
	AccessibilityHelp() string
	// Returns the horizontal scroll bar for the scroll view.
	AccessibilityHorizontalScrollBar() objectivec.IObject
	// Returns the description of the layout area’s horizontal units.
	AccessibilityHorizontalUnitDescription() string
	// Returns the units that the layout area uses for horizontal values.
	AccessibilityHorizontalUnits() NSAccessibilityUnits
	// Returns the accessibility element’s identity.
	AccessibilityIdentifier() string
	// Returns the increment button for the stepper accessibility element.
	AccessibilityIncrementButton() objectivec.IObject
	// Returns the index of the row or column that the accessibility element represents.
	AccessibilityIndex() int
	// Returns the line number that contains the insertion point.
	AccessibilityInsertionPointLineNumber() int
	// Returns a short description of the accessibility element.
	AccessibilityLabel() string
	// Returns the child label elements for the slider accessibility element.
	AccessibilityLabelUIElements() foundation.INSArray
	// Returns the value of the label accessibility element.
	AccessibilityLabelValue() float32
	// Converts the provided point in screen coordinates to a point in the layout area’s coordinate system.
	AccessibilityLayoutPointForScreenPoint(point corefoundation.CGPoint) corefoundation.CGPoint
	// Converts the provided size in screen coordinates to a size in the layout area’s coordinate system.
	AccessibilityLayoutSizeForScreenSize(size corefoundation.CGSize) corefoundation.CGSize
	// Returns the line number for the line that contains the specified character index.
	AccessibilityLineForIndex(index int) int
	// Returns the elements that have links with the accessibility element.
	AccessibilityLinkedUIElements() foundation.INSArray
	// Returns the app’s main window.
	AccessibilityMainWindow() objectivec.IObject
	// Returns the user interface element that functions as a marker group for the ruler accessibility element.
	AccessibilityMarkerGroupUIElement() objectivec.IObject
	// Returns the human-readable description of the marker type.
	AccessibilityMarkerTypeDescription() string
	// Returns the array of marker accessibility elements for the ruler.
	AccessibilityMarkerUIElements() foundation.INSArray
	// Returns the marker values for the ruler.
	AccessibilityMarkerValues() objectivec.IObject
	// Returns the maximum value for the accessibility element.
	AccessibilityMaxValue() objectivec.IObject
	// Returns the app’s menu bar.
	AccessibilityMenuBar() objectivec.IObject
	// Returns the minimum value for the accessibility element.
	AccessibilityMinValue() objectivec.IObject
	// Returns the child accessibility element that represents the window’s minimize button.
	AccessibilityMinimizeButton() objectivec.IObject
	// Returns the contents that follow the divider accessibility element.
	AccessibilityNextContents() foundation.INSArray
	// Returns the number of characters in the text.
	AccessibilityNumberOfCharacters() int
	// Returns the orientation of the accessibility element.
	AccessibilityOrientation() NSAccessibilityOrientation
	// Returns the overflow button for the toolbar.
	AccessibilityOverflowButton() objectivec.IObject
	// Returns the accessibility element’s parent in the accessibility hierarchy.
	AccessibilityParent() objectivec.IObject
	// Cancels the current operation.
	AccessibilityPerformCancel() bool
	// Simulates pressing Return in the accessibility element.
	AccessibilityPerformConfirm() bool
	// Decrements the accessibility element’s value.
	AccessibilityPerformDecrement() bool
	// Deletes the accessibility element’s value.
	AccessibilityPerformDelete() bool
	// Increments the accessibility element’s value.
	AccessibilityPerformIncrement() bool
	// Selects the accessibility element.
	AccessibilityPerformPick() bool
	// Simulates clicking the accessibility element.
	AccessibilityPerformPress() bool
	// Brings the window to the front.
	AccessibilityPerformRaise() bool
	// Displays the accessibility element’s alternative UI.
	AccessibilityPerformShowAlternateUI() bool
	// Returns to the accessibility element’s original UI.
	AccessibilityPerformShowDefaultUI() bool
	// Displays the menu accessibility element.
	AccessibilityPerformShowMenu() bool
	// Returns the placeholder value for the accessibility element.
	AccessibilityPlaceholderValue() string
	// Returns the contents that precede the divider accessibility element.
	AccessibilityPreviousContents() foundation.INSArray
	// Returns the child accessibility element that represents the window’s proxy icon.
	AccessibilityProxy() objectivec.IObject
	// Returns the rich text format (RTF) data that describes the specified range of characters.
	AccessibilityRTFForRange(range_ foundation.NSRange) foundation.NSData
	// Returns the range of characters for the glyph that includes the specified character.
	AccessibilityRangeForIndex(index int) foundation.NSRange
	// Returns the range of characters in the specified line.
	AccessibilityRangeForLine(line int) foundation.NSRange
	// Returns the range of characters for the glyph at the specified point.
	AccessibilityRangeForPosition(point corefoundation.CGPoint) foundation.NSRange
	// Returns the type of interface element that the accessibility element represents.
	AccessibilityRole() NSAccessibilityRole
	// Returns a localized, human-intelligible description of the accessibility element’s role, such as radio button.
	AccessibilityRoleDescription() string
	// Returns the number of rows in the accessibility element’s grid.
	AccessibilityRowCount() int
	// Returns the row header accessibility elements for the table or outline.
	AccessibilityRowHeaderUIElements() foundation.INSArray
	// Returns the row index range of the cell.
	AccessibilityRowIndexRange() foundation.NSRange
	// Returns the row accessibility elements for the table or outline.
	AccessibilityRows() foundation.INSArray
	// Returns the type of markers for the ruler.
	AccessibilityRulerMarkerType() NSAccessibilityRulerMarkerType
	// Converts the provided point in the layout area’s coordinates to a point in the screen’s coordinate system.
	AccessibilityScreenPointForLayoutPoint(point corefoundation.CGPoint) corefoundation.CGPoint
	// Converts the provided size in the layout area’s coordinates to a size in the screen’s coordinate system.
	AccessibilityScreenSizeForLayoutSize(size corefoundation.CGSize) corefoundation.CGSize
	// Returns the search button for the search field.
	AccessibilitySearchButton() objectivec.IObject
	// Returns the search menu for the search field.
	AccessibilitySearchMenu() objectivec.IObject
	// Returns the currently selected cells for the table.
	AccessibilitySelectedCells() foundation.INSArray
	// Returns the accessibility element’s currently selected children.
	AccessibilitySelectedChildren() foundation.INSArray
	// Returns the currently selected columns for the table or outline.
	AccessibilitySelectedColumns() foundation.INSArray
	// Returns the currently selected rows for the table or outline.
	AccessibilitySelectedRows() foundation.INSArray
	// Returns the currently selected text.
	AccessibilitySelectedText() string
	// Returns the range of the currently selected text.
	AccessibilitySelectedTextRange() foundation.NSRange
	// Returns an array of ranges for the currently selected text.
	AccessibilitySelectedTextRanges() []foundation.NSValue
	// Returns the list of elements that the accessibility element is a title for.
	AccessibilityServesAsTitleForUIElements() foundation.INSArray
	// Returns the range of characters that the accessibility element displays.
	AccessibilitySharedCharacterRange() foundation.NSRange
	// Returns the array of elements that shares the keyboard focus with the accessibility element.
	AccessibilitySharedFocusElements() foundation.INSArray
	// Returns the other elements that share text with the accessibility element.
	AccessibilitySharedTextUIElements() foundation.INSArray
	// Returns the menu currently displaying for the accessibility element.
	AccessibilityShownMenu() objectivec.IObject
	// Returns the accessibility element’s sort direction.
	AccessibilitySortDirection() NSAccessibilitySortDirection
	// Returns an array that contains the views and splitter bar from the split view.
	AccessibilitySplitters() foundation.INSArray
	// Returns the substring for the specified range.
	AccessibilityStringForRange(range_ foundation.NSRange) string
	// Returns a range of characters that all have the same style as the specified character.
	AccessibilityStyleRangeForIndex(index int) foundation.NSRange
	// Returns the specialized interface element type that the accessibility element represents.
	AccessibilitySubrole() NSAccessibilitySubrole
	// Returns the tab accessibility elements for the tab view.
	AccessibilityTabs() foundation.INSArray
	// Returns the title of the accessibility element—for example, a button’s visible text.
	AccessibilityTitle() string
	// Returns the static text element that represents the accessibility element’s title.
	AccessibilityTitleUIElement() objectivec.IObject
	// Returns the child accessibility element that represents the window’s toolbar button.
	AccessibilityToolbarButton() objectivec.IObject
	// Returns the top-level element that contains the accessibility element.
	AccessibilityTopLevelUIElement() objectivec.IObject
	// Returns the URL for the accessibility element.
	AccessibilityURL() foundation.NSURL
	// Returns the human-readable description of the ruler’s units.
	AccessibilityUnitDescription() string
	// Returns the units for the ruler.
	AccessibilityUnits() NSAccessibilityUnits
	AccessibilityUserInputLabels() []string
	// Returns the accessibility element’s value.
	AccessibilityValue() objectivec.IObject
	// Returns the human-readable description of the accessibility element’s value.
	AccessibilityValueDescription() string
	// Returns the vertical scroll bar for the scroll view.
	AccessibilityVerticalScrollBar() objectivec.IObject
	// Returns the description of the layout area’s vertical units.
	AccessibilityVerticalUnitDescription() string
	// Returns the units that the layout area uses for vertical values.
	AccessibilityVerticalUnits() NSAccessibilityUnits
	// Returns the visible cells for the table.
	AccessibilityVisibleCells() foundation.INSArray
	// Returns the range of visible characters in the document.
	AccessibilityVisibleCharacterRange() foundation.NSRange
	// Returns the accessibility element’s visible child accessibility elements.
	AccessibilityVisibleChildren() foundation.INSArray
	// Returns the visible columns for the table or outline.
	AccessibilityVisibleColumns() foundation.INSArray
	// Returns the visible rows for the table or outline.
	AccessibilityVisibleRows() foundation.INSArray
	// Returns the warning value for the level indicator.
	AccessibilityWarningValue() objectivec.IObject
	// Returns the window that contains the accessibility element.
	AccessibilityWindow() objectivec.IObject
	// Returns an array that contains all the app’s windows.
	AccessibilityWindows() foundation.INSArray
	// Returns the child accessibility element that represents the window’s zoom button.
	AccessibilityZoomButton() objectivec.IObject
	// Adds an item to the Window menu for a given window.
	AddWindowsItemTitleFilename(win INSWindow, string_ string, isFilename bool)
	// Arranges windows listed in the Window menu in front of all other windows.
	ArrangeInFront(sender objectivec.IObject)
	// Sets up a modal session with the given window and returns a pointer to the [NSModalSession] structure representing the session.
	BeginModalSessionForWindow(window INSWindow) NSModalSession
	// Changes the item for a given window in the Window menu to a given string.
	ChangeWindowsItemTitleFilename(win INSWindow, string_ string, isFilename bool)
	// Completes the extended state restoration.
	CompleteStateRestoration()
	// Finishes a modal session.
	EndModalSession(session NSModalSession)
	// Executes a block for each of the app’s windows.
	EnumerateWindowsWithOptionsUsingBlock(options NSWindowListOptions, block WindowBoolHandler)
	// Allows an app to extend its state restoration period.
	ExtendStateRestoration()
	// Hides all the receiver’s windows, and the next app in line is activated.
	Hide(sender objectivec.IObject)
	// Returns the Boolean value that determines whether the accessibility element’s alternative UI is currently visible.
	IsAccessibilityAlternateUIVisible() bool
	// Returns a Boolean value that determines whether the row is disclosing other rows.
	IsAccessibilityDisclosed() bool
	// Returns a Boolean value that indicates whether the accessibility element is in an edited state.
	IsAccessibilityEdited() bool
	// Returns a Boolean value that determines whether the accessibility element participates in the accessibility hierarchy.
	IsAccessibilityElement() bool
	// Returns a Boolean value that determines whether the accessibility element responds to user events.
	IsAccessibilityEnabled() bool
	// Returns a Boolean value that determines whether the accessibility element is in an expanded state.
	IsAccessibilityExpanded() bool
	// Returns a Boolean value that indicates whether the accessibility element has the keyboard focus.
	IsAccessibilityFocused() bool
	// Returns a Boolean value that determines whether the app is the frontmost app.
	IsAccessibilityFrontmost() bool
	// Returns a Boolean value that determines whether the app is in a hidden state.
	IsAccessibilityHidden() bool
	// Returns a Boolean value that determines whether the window is the app’s main window.
	IsAccessibilityMain() bool
	// Returns the Boolean value that determines whether the window is in a minimized state.
	IsAccessibilityMinimized() bool
	// Returns a Boolean value that determines whether the window is modal.
	IsAccessibilityModal() bool
	// Returns a Boolean value that determines whether the accessibility element’s grid is in row major order or in column major order.
	IsAccessibilityOrderedByRow() bool
	// Returns a Boolean value that determines whether the accessibility element contains protected content.
	IsAccessibilityProtectedContent() bool
	// Returns a Boolean value that determines whether the accessibility element must have content for successful submission of a form.
	IsAccessibilityRequired() bool
	// Returns a Boolean value that determines whether the accessibility element is currently in a selected state.
	IsAccessibilitySelected() bool
	// Returns a Boolean value that indicates whether assistive apps can invoke the specified selector on the accessibility element.
	IsAccessibilitySelectorAllowed(selector objc.SEL) bool
	// Miniaturizes all the receiver’s windows.
	MiniaturizeAll(sender objectivec.IObject)
	// Opens the character palette.
	OrderFrontCharacterPalette(sender objectivec.IObject)
	// Brings up the color panel, an instance of [NSColorPanel].
	OrderFrontColorPanel(sender objectivec.IObject)
	// Displays a standard About window.
	OrderFrontStandardAboutPanel(sender objectivec.IObject)
	// Displays a standard About window with information from a given options dictionary.
	OrderFrontStandardAboutPanelWithOptions(optionsDictionary foundation.INSDictionary)
	// Suppresses the usual window ordering in handling the most recent mouse-down event.
	PreventWindowOrdering()
	// Registers the pasteboard types the receiver can send and receive in response to service requests.
	RegisterServicesMenuSendTypesReturnTypes(sendTypes []string, returnTypes []string)
	// Removes the Window menu item for a given window.
	RemoveWindowsItem(win INSWindow)
	// Invoked to request that a window be restored.
	RestoreWindowWithIdentifierStateCompletionHandler(identifier NSUserInterfaceItemIdentifier, state foundation.INSCoder, completionHandler WindowErrorHandler) bool
	// Starts a modal event loop for the specified window.
	RunModalForWindow(window INSWindow) NSModalResponse
	// Runs a given modal session, as defined in a previous invocation of [beginModalSession(for:)](<https://developer.apple.com/documentation/AppKit/NSApplication/beginModalSession(for:)>).
	RunModalSession(session NSModalSession) NSModalResponse
	// Displays the receiver’s page layout panel, an instance of [NSPageLayout].
	RunPageLayout(sender objectivec.IObject)
	// Sets the activation point for the user interface element.
	SetAccessibilityActivationPoint(accessibilityActivationPoint corefoundation.CGPoint)
	// Sets the allowed values for the slider accessibility element.
	SetAccessibilityAllowedValues(accessibilityAllowedValues []foundation.NSNumber)
	// Sets the Boolean value that determines whether the accessibility element’s alternative UI is currently visible.
	SetAccessibilityAlternateUIVisible(accessibilityAlternateUIVisible bool)
	// Sets the child accessibility element with the current focus.
	SetAccessibilityApplicationFocusedUIElement(accessibilityApplicationFocusedUIElement objectivec.IObject)
	SetAccessibilityAttributedUserInputLabels(accessibilityAttributedUserInputLabels []foundation.NSAttributedString)
	// Sets the child accessibility element that represents the window’s cancel button.
	SetAccessibilityCancelButton(accessibilityCancelButton objectivec.IObject)
	// Sets the child accessibility elements in the accessibility hierarchy.
	SetAccessibilityChildren(accessibilityChildren foundation.INSArray)
	// Sets the array of child accessibility elements in order for linear navigation.
	SetAccessibilityChildrenInNavigationOrder(accessibilityChildrenInNavigationOrder []objectivec.IObject)
	// Sets the clear button for the search field.
	SetAccessibilityClearButton(accessibilityClearButton objectivec.IObject)
	// Sets the child accessibility element that represents the window’s close button.
	SetAccessibilityCloseButton(accessibilityCloseButton objectivec.IObject)
	// Sets the number of columns in the accessibility element’s grid.
	SetAccessibilityColumnCount(accessibilityColumnCount int)
	// Sets the column header accessibility elements for the table or outline.
	SetAccessibilityColumnHeaderUIElements(accessibilityColumnHeaderUIElements foundation.INSArray)
	// Sets the column index range of the cell.
	SetAccessibilityColumnIndexRange(accessibilityColumnIndexRange foundation.NSRange)
	// Sets the column titles for the accessibility element.
	SetAccessibilityColumnTitles(accessibilityColumnTitles foundation.INSArray)
	// Sets the column accessibility elements for the table or outline.
	SetAccessibilityColumns(accessibilityColumns foundation.INSArray)
	// Sets the contents of the current accessibility element.
	SetAccessibilityContents(accessibilityContents foundation.INSArray)
	// Sets the critical value for the level indicator.
	SetAccessibilityCriticalValue(accessibilityCriticalValue objectivec.IObject)
	// Sets the custom actions of the current accessibility element.
	SetAccessibilityCustomActions(accessibilityCustomActions []NSAccessibilityCustomAction)
	// Sets the custom rotors of the current accessibility element.
	SetAccessibilityCustomRotors(accessibilityCustomRotors []NSAccessibilityCustomRotor)
	// Sets the decrement button for the stepper accessibility element.
	SetAccessibilityDecrementButton(accessibilityDecrementButton objectivec.IObject)
	// Sets the child accessibility element that represents the window’s default button.
	SetAccessibilityDefaultButton(accessibilityDefaultButton objectivec.IObject)
	// Sets a Boolean value that determines whether the row is disclosing other rows.
	SetAccessibilityDisclosed(accessibilityDisclosed bool)
	// Sets the row disclosing the current row.
	SetAccessibilityDisclosedByRow(accessibilityDisclosedByRow objectivec.IObject)
	// Sets the rows that the current row discloses.
	SetAccessibilityDisclosedRows(accessibilityDisclosedRows objectivec.IObject)
	// Sets the indention level for the row.
	SetAccessibilityDisclosureLevel(accessibilityDisclosureLevel int)
	// Sets the URL for the file that the accessibility element represents.
	SetAccessibilityDocument(accessibilityDocument string)
	// Sets a Boolean value that indicates whether the accessibility element is in an edited state.
	SetAccessibilityEdited(accessibilityEdited bool)
	// Sets a Boolean value that determines whether the accessibility element participates in the accessibility hierarchy.
	SetAccessibilityElement(accessibilityElement bool)
	// Sets a Boolean value that determines whether the accessibility element responds to user events.
	SetAccessibilityEnabled(accessibilityEnabled bool)
	// Sets a Boolean value that determines whether accessibility element is in an expanded state.
	SetAccessibilityExpanded(accessibilityExpanded bool)
	// Sets the icon for the app’s menu bar extra.
	SetAccessibilityExtrasMenuBar(accessibilityExtrasMenuBar objectivec.IObject)
	// Sets the filename for the file that the accessibility element represents.
	SetAccessibilityFilename(accessibilityFilename string)
	// Sets a Boolean value that determines whether the accessibility element has the keyboard focus.
	SetAccessibilityFocused(accessibilityFocused bool)
	// Sets the child window with the current focus.
	SetAccessibilityFocusedWindow(accessibilityFocusedWindow objectivec.IObject)
	// Sets the accessibility element’s frame in screen coordinates.
	SetAccessibilityFrame(accessibilityFrame corefoundation.CGRect)
	// Sets a Boolean value that determines whether the app is the frontmost app.
	SetAccessibilityFrontmost(accessibilityFrontmost bool)
	// Sets the child accessibility element that represents the window’s full-screen button.
	SetAccessibilityFullScreenButton(accessibilityFullScreenButton objectivec.IObject)
	// Sets the child accessibility element that represents the window’s grow area.
	SetAccessibilityGrowArea(accessibilityGrowArea objectivec.IObject)
	// Sets the drag handle accessibility elements for the layout item element.
	SetAccessibilityHandles(accessibilityHandles foundation.INSArray)
	// Sets the header for the table view.
	SetAccessibilityHeader(accessibilityHeader objectivec.IObject)
	// Sets the help text for the accessibility element.
	SetAccessibilityHelp(accessibilityHelp string)
	// Sets a Boolean value that determines whether the app is in a hidden state.
	SetAccessibilityHidden(accessibilityHidden bool)
	// Sets the horizontal scroll bar for the scroll view.
	SetAccessibilityHorizontalScrollBar(accessibilityHorizontalScrollBar objectivec.IObject)
	// Sets the description of the layout area’s horizontal units.
	SetAccessibilityHorizontalUnitDescription(accessibilityHorizontalUnitDescription string)
	// Sets the units that the layout area uses for horizontal values.
	SetAccessibilityHorizontalUnits(accessibilityHorizontalUnits NSAccessibilityUnits)
	// Sets the accessibility element’s identity.
	SetAccessibilityIdentifier(accessibilityIdentifier string)
	// Sets the increment button for the stepper accessibility element.
	SetAccessibilityIncrementButton(accessibilityIncrementButton objectivec.IObject)
	// Sets the index of the row or column that the accessibility element represents.
	SetAccessibilityIndex(accessibilityIndex int)
	// Sets the line number that contains the insertion point.
	SetAccessibilityInsertionPointLineNumber(accessibilityInsertionPointLineNumber int)
	// Sets a short description of the accessibility element.
	SetAccessibilityLabel(accessibilityLabel string)
	// Sets the child label elements for the slider accessibility element.
	SetAccessibilityLabelUIElements(accessibilityLabelUIElements foundation.INSArray)
	// Sets the value of the label accessibility element.
	SetAccessibilityLabelValue(accessibilityLabelValue float32)
	// Sets the elements that have links with the accessibility element.
	SetAccessibilityLinkedUIElements(accessibilityLinkedUIElements foundation.INSArray)
	// Sets a Boolean value that determines whether the window is the app’s main window.
	SetAccessibilityMain(accessibilityMain bool)
	// Sets the app’s main window.
	SetAccessibilityMainWindow(accessibilityMainWindow objectivec.IObject)
	// Sets the user interface element that functions as a marker group for the ruler accessibility element.
	SetAccessibilityMarkerGroupUIElement(accessibilityMarkerGroupUIElement objectivec.IObject)
	// Sets the human-readable description of the marker type.
	SetAccessibilityMarkerTypeDescription(accessibilityMarkerTypeDescription string)
	// Sets the array of marker accessibility elements for the ruler.
	SetAccessibilityMarkerUIElements(accessibilityMarkerUIElements foundation.INSArray)
	// Sets the marker values for the ruler.
	SetAccessibilityMarkerValues(accessibilityMarkerValues objectivec.IObject)
	// Sets the maximum value for the accessibility element.
	SetAccessibilityMaxValue(accessibilityMaxValue objectivec.IObject)
	// Sets the app’s menu bar.
	SetAccessibilityMenuBar(accessibilityMenuBar objectivec.IObject)
	// Sets the minimum value for the accessibility element.
	SetAccessibilityMinValue(accessibilityMinValue objectivec.IObject)
	// Sets the child accessibility element that represents the window’s minimize button.
	SetAccessibilityMinimizeButton(accessibilityMinimizeButton objectivec.IObject)
	// Sets the Boolean value that determines whether the window is in a minimized state.
	SetAccessibilityMinimized(accessibilityMinimized bool)
	// Sets a Boolean value that determines whether the window is modal.
	SetAccessibilityModal(accessibilityModal bool)
	// Sets the contents that follow the divider accessibility element.
	SetAccessibilityNextContents(accessibilityNextContents foundation.INSArray)
	// Sets the number of characters in the text.
	SetAccessibilityNumberOfCharacters(accessibilityNumberOfCharacters int)
	// Sets a Boolean value that determines whether the element’s grid is in row major order or in column major order.
	SetAccessibilityOrderedByRow(accessibilityOrderedByRow bool)
	// Sets the orientation of the accessibility element.
	SetAccessibilityOrientation(accessibilityOrientation NSAccessibilityOrientation)
	// Sets the overflow button for the toolbar.
	SetAccessibilityOverflowButton(accessibilityOverflowButton objectivec.IObject)
	// Sets the accessibility element’s parent in the accessibility hierarchy.
	SetAccessibilityParent(accessibilityParent objectivec.IObject)
	// Sets the placeholder value for the accessibility element.
	SetAccessibilityPlaceholderValue(accessibilityPlaceholderValue string)
	// Sets the contents that precede the divider accessibility element.
	SetAccessibilityPreviousContents(accessibilityPreviousContents foundation.INSArray)
	// Sets a Boolean value that determines whether the accessibility element contains protected content.
	SetAccessibilityProtectedContent(accessibilityProtectedContent bool)
	// Sets the child accessibility element that represents the window’s proxy icon.
	SetAccessibilityProxy(accessibilityProxy objectivec.IObject)
	// Sets a Boolean value that determines whether the accessibility element must have content for successful submission of a form.
	SetAccessibilityRequired(accessibilityRequired bool)
	// Sets the type of interface element that the accessibility element represents.
	SetAccessibilityRole(accessibilityRole NSAccessibilityRole)
	// Sets the localized, human-intelligible description of the accessibility element’s role, such as radio button.
	SetAccessibilityRoleDescription(accessibilityRoleDescription string)
	// Sets the number of rows in the accessibility element’s grid.
	SetAccessibilityRowCount(accessibilityRowCount int)
	// Sets the row header accessibility elements for the table or outline.
	SetAccessibilityRowHeaderUIElements(accessibilityRowHeaderUIElements foundation.INSArray)
	// Sets the row index range of the cell.
	SetAccessibilityRowIndexRange(accessibilityRowIndexRange foundation.NSRange)
	// Sets the row accessibility elements for the table or outline.
	SetAccessibilityRows(accessibilityRows foundation.INSArray)
	// Sets the type of markers for the ruler.
	SetAccessibilityRulerMarkerType(accessibilityRulerMarkerType NSAccessibilityRulerMarkerType)
	// Sets the search button for the search field.
	SetAccessibilitySearchButton(accessibilitySearchButton objectivec.IObject)
	// Sets the search menu for the search field.
	SetAccessibilitySearchMenu(accessibilitySearchMenu objectivec.IObject)
	// Sets a Boolean value that determines whether the accessibility element is currently in a selected state.
	SetAccessibilitySelected(accessibilitySelected bool)
	// Sets the currently selected cells for the table.
	SetAccessibilitySelectedCells(accessibilitySelectedCells foundation.INSArray)
	// Sets the accessibility element’s currently selected children.
	SetAccessibilitySelectedChildren(accessibilitySelectedChildren foundation.INSArray)
	// Sets the currently selected columns for the table or outline.
	SetAccessibilitySelectedColumns(accessibilitySelectedColumns foundation.INSArray)
	// Sets the currently selected rows for the table or outline.
	SetAccessibilitySelectedRows(accessibilitySelectedRows foundation.INSArray)
	// Sets the currently selected text.
	SetAccessibilitySelectedText(accessibilitySelectedText string)
	// Sets the range of the currently selected text.
	SetAccessibilitySelectedTextRange(accessibilitySelectedTextRange foundation.NSRange)
	// Sets an array of ranges for the currently selected text.
	SetAccessibilitySelectedTextRanges(accessibilitySelectedTextRanges []foundation.NSValue)
	// Sets the list of elements that the accessibility element is a title for.
	SetAccessibilityServesAsTitleForUIElements(accessibilityServesAsTitleForUIElements foundation.INSArray)
	// Sets the range of characters that the accessibility element displays.
	SetAccessibilitySharedCharacterRange(accessibilitySharedCharacterRange foundation.NSRange)
	// Sets the array of elements that shares the keyboard focus with the accessibility element.
	SetAccessibilitySharedFocusElements(accessibilitySharedFocusElements foundation.INSArray)
	// Sets the other elements that share text with the accessibility element.
	SetAccessibilitySharedTextUIElements(accessibilitySharedTextUIElements foundation.INSArray)
	// Sets the menu currently displaying for the accessibility element.
	SetAccessibilityShownMenu(accessibilityShownMenu objectivec.IObject)
	// Sets the accessibility element’s sort direction.
	SetAccessibilitySortDirection(accessibilitySortDirection NSAccessibilitySortDirection)
	// Sets the array that contains the views and splitter bar from the split view.
	SetAccessibilitySplitters(accessibilitySplitters foundation.INSArray)
	// Sets the specialized interface element type that the accessibility element represents.
	SetAccessibilitySubrole(accessibilitySubrole NSAccessibilitySubrole)
	// Sets the tab accessibility elements for the tab view.
	SetAccessibilityTabs(accessibilityTabs foundation.INSArray)
	// Sets the title of the accessibility element.
	SetAccessibilityTitle(accessibilityTitle string)
	// Sets the static text element that represents the accessibility element’s title.
	SetAccessibilityTitleUIElement(accessibilityTitleUIElement objectivec.IObject)
	// Sets the child accessibility element that represents the window’s toolbar button.
	SetAccessibilityToolbarButton(accessibilityToolbarButton objectivec.IObject)
	// Sets the top-level element that contains the accessibility element.
	SetAccessibilityTopLevelUIElement(accessibilityTopLevelUIElement objectivec.IObject)
	// Sets the URL for the accessibility element.
	SetAccessibilityURL(accessibilityURL foundation.NSURL)
	// Sets the human-readable description of the ruler’s units.
	SetAccessibilityUnitDescription(accessibilityUnitDescription string)
	// Sets the units used for the ruler.
	SetAccessibilityUnits(accessibilityUnits NSAccessibilityUnits)
	SetAccessibilityUserInputLabels(accessibilityUserInputLabels []string)
	// Sets the accessibility element’s value.
	SetAccessibilityValue(accessibilityValue objectivec.IObject)
	// Sets the human-readable description of the accessibility element’s value.
	SetAccessibilityValueDescription(accessibilityValueDescription string)
	// Sets the vertical scroll bar for the scroll view.
	SetAccessibilityVerticalScrollBar(accessibilityVerticalScrollBar objectivec.IObject)
	// Sets the description of the layout area’s vertical units.
	SetAccessibilityVerticalUnitDescription(accessibilityVerticalUnitDescription string)
	// Sets the units that the layout area uses for vertical values.
	SetAccessibilityVerticalUnits(accessibilityVerticalUnits NSAccessibilityUnits)
	// Sets the visible cells for the table.
	SetAccessibilityVisibleCells(accessibilityVisibleCells foundation.INSArray)
	// Sets the range of visible characters in the document.
	SetAccessibilityVisibleCharacterRange(accessibilityVisibleCharacterRange foundation.NSRange)
	// Sets the accessibility element’s visible child accessibility elements.
	SetAccessibilityVisibleChildren(accessibilityVisibleChildren foundation.INSArray)
	// Sets the visible columns for the table or outline.
	SetAccessibilityVisibleColumns(accessibilityVisibleColumns foundation.INSArray)
	// Sets the visible rows for the table or outline.
	SetAccessibilityVisibleRows(accessibilityVisibleRows foundation.INSArray)
	// Sets the warning value for the level indicator.
	SetAccessibilityWarningValue(accessibilityWarningValue objectivec.IObject)
	// Sets the window that contains the accessibility element.
	SetAccessibilityWindow(accessibilityWindow objectivec.IObject)
	// Sets the array that contains all the app’s windows.
	SetAccessibilityWindows(accessibilityWindows foundation.INSArray)
	// Sets the child accessibility element that represents the window’s zoom button.
	SetAccessibilityZoomButton(accessibilityZoomButton objectivec.IObject)
	// Sets whether the receiver’s windows need updating when the receiver has finished processing the current event.
	SetWindowsNeedUpdate(needUpdate bool)
	// Stops a modal event loop.
	StopModal()
	// Stops a modal event loop, allowing you to return a custom result code.
	StopModalWithCode(returnCode NSModalResponse)
	// Restores hidden windows to the screen and makes the receiver active.
	Unhide(sender objectivec.IObject)
	// Restores hidden windows without activating their owner (the receiver).
	UnhideWithoutActivation()
	// Sends an [update()](<https://developer.apple.com/documentation/AppKit/NSWindow/update()>) message to each onscreen window.
	UpdateWindows()
	// Updates the Window menu item for a given window to reflect the edited status of that window.
	UpdateWindowsItem(win INSWindow)
	// Implemented to override the default action of enabling or disabling a specific menu item.
	ValidateMenuItem(menuItem INSMenuItem) bool
	// Returns a Boolean value that indicates whether the sender should be enabled.
	ValidateUserInterfaceItem(item NSValidatedUserInterfaceItem) bool
	// Returns the window corresponding to the specified window number.
	WindowWithWindowNumber(windowNum int) INSWindow
}

// Init initializes the instance.
func (a NSApplication) Init() NSApplication {
	rv := objc.Send[NSApplication](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a NSApplication) Autorelease() NSApplication {
	rv := objc.Send[NSApplication](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSApplication creates a new NSApplication instance.
func NewNSApplication() NSApplication {
	class := getNSApplicationClass()
	rv := objc.Send[NSApplication](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a new responder object with data in an unarchiver.
//
// coder: An unarchiver object.
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/init(coder:)
func NewApplicationWithCoder(coder foundation.INSCoder) NSApplication {
	instance := getNSApplicationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSApplicationFromID(rv)
}

// Returns the next event matching a given mask, or `nil` if no such event is
// found before a specified expiration date.
//
// mask: Contains one or more flags indicating the types of events to return. The
// constants section of the [NSEvent] class defines the constants you can add
// together to create this mask. The
// [NSApplication.DiscardEventsMatchingMaskBeforeEvent] method also lists
// several of these constants.
//
// expiration: The expiration date for the current event request. Specifying nil for this
// parameter is equivalent to returning a date object using the [distantPast]
// method.
//
// mode: The run loop mode in which to run while looking for events. The mode you
// specify also determines which timers and run-loop observers may fire while
// the app waits for the event.
//
// deqFlag: Specify true if you want the event removed from the queue.
//
// # Return Value
//
// The event object whose type matches one of the event types specified by the
// `mask` parameter.
//
// # Discussion
//
// You can use this method to short circuit normal event dispatching and get
// your own events. For example, you may want to do this in response to a
// mouse-down event in order to track the mouse while its button is down. (In
// such an example, you’d pass the appropriate event types for mouse-dragged
// and mouse-up events to the `mask` parameter and specify the
// [NSEventTrackingRunLoopMode] run loop mode). Events that don’t match one
// of the specified event types are left in the queue.
//
// You can specify one of the run loop modes defined by AppKit or a custom run
// loop mode used specifically by your app. AppKit defines the following
// run-loop modes:
//
// - [NSDefaultRunLoopMode] - [NSEventTrackingRunLoopMode] -
// [NSModalPanelRunLoopMode] - [NSConnectionReplyMode]
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/nextEvent(matching:until:inMode:dequeue:)
//
// [distantPast]: https://developer.apple.com/documentation/Foundation/NSDate/distantPast
func (a NSApplication) NextEventMatchingMaskUntilDateInModeDequeue(mask NSEventMask, expiration foundation.NSDate, mode foundation.NSRunLoopMode, deqFlag bool) INSEvent {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("nextEventMatchingMask:untilDate:inMode:dequeue:"), mask, expiration, objc.String(string(mode)), deqFlag)
	return NSEventFromID(rv)
}

// Removes all events matching the given mask and generated before the
// specified event.
//
// mask: Contains one or more flags indicating the types of events to discard. The
// constants section of the [NSEvent] class defines the constants you can add
// together to create this mask. The discussion section also lists some of the
// constants that are typically used.
//
// lastEvent: A marker event that you use to indicate which events should be discarded.
// Events that occurred before this event are discarded but those that
// occurred after it are not.
//
// # Discussion
//
// Use this method to ignore any events that occurred before a specific event.
// For example, suppose your app has a tracking loop that you exit when the
// user releases the mouse button. You could use this method, specifying
// [NSAnyEventMask] as the mask argument and the ending mouse-up event as the
// `lastEvent` argument, to discard all events that occurred while you were
// tracking mouse movements in your loop. Passing the mouse-up event as
// `lastEvent` ensures that any events that might have occurred after the
// mouse-up event (that is, that appear in the queue after the mouse-up event)
// aren’t discarded.
//
// For the `mask` parameter, you can add together event type constants such as
// the following:
//
// - [NSLeftMouseDownMask] - [NSLeftMouseUpMask] - [NSRightMouseDownMask] -
// [NSRightMouseUpMask] - [NSMouseMovedMask] - [NSLeftMouseDraggedMask] -
// [NSRightMouseDraggedMask] - [NSMouseEnteredMask] - [NSMouseExitedMask] -
// [NSKeyDownMask] - [NSKeyUpMask] - [NSFlagsChangedMask] - [NSPeriodicMask] -
// [NSCursorUpdateMask] - [NSAnyEventMask]
//
// This method can also be called in subthreads. Events posted in subthreads
// bubble up in the main thread event queue.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/discardEvents(matching:before:)
func (a NSApplication) DiscardEventsMatchingMaskBeforeEvent(mask NSEventMask, lastEvent INSEvent) {
	objc.Send[objc.ID](a.ID, objc.Sel("discardEventsMatchingMask:beforeEvent:"), mask, lastEvent)
}

// Starts the main event loop.
//
// # Discussion
//
// The loop continues until a [NSApplication.Stop] or
// [NSApplication.Terminate] message is received. Upon each iteration through
// the loop, the next available event from the window server is stored and
// then dispatched by sending it to [NSApp] using [NSApplication.SendEvent].
//
// After creating the [NSApplication] object, the `main` function should load
// your app’s main nib file and then start the event loop by sending the
// [NSApplication] object a [NSApplication.Run] message. If you create an
// Cocoa app project in Xcode, this `main` function is implemented for you.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/run()
func (a NSApplication) Run() {
	objc.Send[objc.ID](a.ID, objc.Sel("run"))
}

// Activates the app, opens any files specified by the [NSOpen] user default,
// and unhighlights the app’s icon.
//
// # Discussion
//
// The [NSApplication.Run] method calls this method before it starts the event
// loop. When this method begins, it posts an
// [willFinishLaunchingNotification] to the default notification center. If
// you override [NSApplication.FinishLaunching], the subclass method should
// invoke the superclass method.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/finishLaunching()
//
// [willFinishLaunchingNotification]: https://developer.apple.com/documentation/AppKit/NSApplication/willFinishLaunchingNotification
func (a NSApplication) FinishLaunching() {
	objc.Send[objc.ID](a.ID, objc.Sel("finishLaunching"))
}

// Stops the main event loop.
//
// sender: The object that sent this message.
//
// # Discussion
//
// This method notifies the app that you want to exit the current run loop as
// soon as it finishes processing the current [NSEvent] object. This method
// doesn’t forcibly exit the current run loop. Instead it sets a flag that
// the app checks only after it finishes dispatching an actual event object.
// For example, you could call this method from an action method responding to
// a button click or from one of the many methods defined by the [NSResponder]
// class. However, calling this method from a timer or run-loop observer
// routine wouldn’t stop the run loop because they don’t result in the
// posting of an [NSEvent] object.
//
// If you call this method from an event handler running in your main run
// loop, the app object exits out of the [NSApplication.Run] method, thereby
// returning control to the `main()` function. If you call this method from
// within a modal event loop, it will exit the modal loop instead of the main
// event loop.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/stop(_:)
func (a NSApplication) Stop(sender objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("stop:"), sender)
}

// Dispatches an event to other objects.
//
// event: The event object to dispatch.
//
// # Discussion
//
// You rarely invoke [NSApplication.SendEvent] directly, although you might
// want to override this method to perform some action on every event.
// [NSApplication.SendEvent] messages are sent from the main event loop (the
// [NSApplication.Run] method). [NSApplication.SendEvent] is the method that
// dispatches events to the appropriate responders—[NSApp] handles app
// events, the [NSWindow] object indicated in the event record handles
// window-related events, and mouse and key events are forwarded to the
// appropriate [NSWindow] object for further dispatching.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/sendEvent(_:)
func (a NSApplication) SendEvent(event INSEvent) {
	objc.Send[objc.ID](a.ID, objc.Sel("sendEvent:"), event)
}

// Adds a given event to the receiver’s event queue.
//
// event: The event object to post to the queue.
//
// atStart: Specify true to add the event to the front of the queue; otherwise, specify
// false to add the event to the back of the queue.
//
// # Discussion
//
// This method can also be called in subthreads. Events posted in subthreads
// bubble up in the main thread event queue.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/postEvent(_:atStart:)
func (a NSApplication) PostEventAtStart(event INSEvent, atStart bool) {
	objc.Send[objc.ID](a.ID, objc.Sel("postEvent:atStart:"), event, atStart)
}

// Sends the given action message to the given target.
//
// action: The action message you want to send.
//
// target: The target object that defines the specified action message.
//
// sender: The object to pass for the action message’s parameter.
//
// # Return Value
//
// true if the action was successfully sent; otherwise false. This method also
// returns false if `anAction` is `nil`.
//
// # Discussion
//
// If `aTarget` is `nil`, [NSApplicationClass.SharedApplication] looks for an
// object that can respond to the message—that is, an object that implements
// a method matching `anAction`. It begins with the first responder of the key
// window. If the first responder can’t respond, it tries the first
// responder’s next responder and continues following next responder links
// up the responder chain. If none of the objects in the key window’s
// responder chain can handle the message,
// [NSApplicationClass.SharedApplication] attempts to send the message to the
// key window’s delegate.
//
// If the delegate doesn’t respond and the main window is different from the
// key window, [NSApplicationClass.SharedApplication] begins again with the
// first responder in the main window. If objects in the main window can’t
// respond, [NSApplicationClass.SharedApplication] attempts to send the
// message to the main window’s delegate. If still no object has responded,
// [NSApplicationClass.SharedApplication] tries to handle the message itself.
// If [NSApplicationClass.SharedApplication] can’t respond, it attempts to
// send the message to its own delegate.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/sendAction(_:to:from:)
func (a NSApplication) SendActionToFrom(action objc.SEL, target objectivec.IObject, sender objectivec.IObject) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("sendAction:to:from:"), action, target, sender)
	return rv
}

// Returns the object that receives the action message specified by the given
// selector.
//
// action: The desired action message.
//
// # Return Value
//
// The object that would receive the specified action message or `nil` if no
// target object would receive the message. This method also returns `nil` if
// `aSelector` is `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/target(forAction:)
func (a NSApplication) TargetForAction(action objc.SEL) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("targetForAction:"), action)
	return objectivec.Object{ID: rv}
}

// Searches for an object that can receive the message specified by the given
// selector.
//
// action: The desired action message. May be `nil`, in which case this method will
// return `nil`.
//
// target: The target object to check. Specify `nil` if you want to search the
// responder chain starting with the current first responder.
//
// sender: The potential sender for the action message.
//
// # Return Value
//
// The object that can accept the specified action message or `nil` if no
// target object can receive the message from the specified `sender`. Returns
// `nil` if `anAction` is `nil`.
//
// # Discussion
//
// The system looks for an object that implements a method matching
// `anAction`.
//
// If `aTarget` is specified, the system verifies that it’s a valid target
// for the provided action and sender, returning `aTarget` if valid, `nil`
// otherwise.
//
// If the provided target is `nil`, the search begins with the first responder
// of the key window. The system follows the responder object looking for
// targets. If no object capable of handling the message is found in the
// responder chain, the system returns `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/target(forAction:to:from:)
func (a NSApplication) TargetForActionToFrom(action objc.SEL, target objectivec.IObject, sender objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("targetForAction:to:from:"), action, target, sender)
	return objectivec.Object{ID: rv}
}

// Terminates the receiver.
//
// sender: Typically, this parameter contains the object that initiated the
// termination request.
//
// # Discussion
//
// This method is typically invoked when the user chooses Quit or Exit from
// the app’s menu.
//
// When invoked, this method performs several steps to process the termination
// request. First, it asks the app’s document controller (if one exists) to
// save any unsaved changes in its documents. During this process, the
// document controller can cancel termination in response to input from the
// user. If the document controller doesn’t cancel the operation, this
// method then calls the delegate’s [ApplicationShouldTerminate] method. If
// [ApplicationShouldTerminate] returns [NSTerminateCancel], the termination
// process is aborted and control is handed back to the main event loop. If
// the method returns [NSTerminateLater], the app runs its run loop in the
// [NSModalPanelRunLoopMode] mode until the
// [NSApplication.ReplyToApplicationShouldTerminate] method is called with the
// value true or false. If the [ApplicationShouldTerminate] method returns
// [NSTerminateNow], this method posts a [willTerminateNotification]
// notification to the default notification center.
//
// Don’t bother to put final cleanup code in your app’s `main()`
// function—it will never be executed. If cleanup is necessary, perform that
// cleanup in the delegate’s [ApplicationWillTerminate] method.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/terminate(_:)
//
// [NSModalPanelRunLoopMode]: https://developer.apple.com/documentation/AppKit/NSModalPanelRunLoopMode
// [willTerminateNotification]: https://developer.apple.com/documentation/AppKit/NSApplication/willTerminateNotification
func (a NSApplication) Terminate(sender objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("terminate:"), sender)
}

// Responds to [NSTerminateLater] once the app knows whether it can terminate.
//
// shouldTerminate: Specify true if you want the app to terminate; otherwise, specify false.
//
// # Discussion
//
// If your app delegate returns [NSTerminateLater] from its
// [ApplicationShouldTerminate] method, your code must subsequently call this
// method to let the [NSApplication] object know whether it can actually
// terminate itself.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/reply(toApplicationShouldTerminate:)
func (a NSApplication) ReplyToApplicationShouldTerminate(shouldTerminate bool) {
	objc.Send[objc.ID](a.ID, objc.Sel("replyToApplicationShouldTerminate:"), shouldTerminate)
}

// Activates the receiver app, if appropriate.
//
// # Discussion
//
// Use this method to request app activation; calling this method doesn’t
// guarantee app activation. For cooperative activation, the other app should
// call [NSApplication.YieldActivationToApplication] or equivalent before the
// target app invokes [NSApplication.Activate].
//
// Invoking [NSApplication.Activate] on an already-active application cancels
// any pending activation yields by the receiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/activate()
func (a NSApplication) Activate() {
	objc.Send[objc.ID](a.ID, objc.Sel("activate"))
}

// Deactivates the receiver.
//
// # Discussion
//
// Normally, you shouldn’t invoke this method—AppKit is responsible for
// proper deactivation.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/deactivate()
func (a NSApplication) Deactivate() {
	objc.Send[objc.ID](a.ID, objc.Sel("deactivate"))
}

// Explicitly allows another app to make itself active.
//
// application: The app to yield activation state to.
//
// # Discussion
//
// Calling this method doesn’t deactivate the yielding app, nor does it
// activate the app you yield to. For cooperative activation, the other app
// must request activation in the future by calling [NSApplication.Activate]
// or equivalent.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/yieldActivation(to:)
func (a NSApplication) YieldActivationToApplication(application INSRunningApplication) {
	objc.Send[objc.ID](a.ID, objc.Sel("yieldActivationToApplication:"), application)
}

// Explicitly allows another app to make itself active.
//
// bundleIdentifier: The bundle identifier to yield activation state to.
//
// # Discussion
//
// Calling this method doesn’t deactivate the yielding app, nor does it
// activate the app you yield to. For cooperative activation, the other app
// must request activation in the future by calling [NSApplication.Activate]
// or equivalent.
//
// Use this method to yield activation to apps that aren’t running at the
// time the method invokes. If it’s known that the target application is
// running, use [NSApplication.YieldActivationToApplication] instead.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/yieldActivation(toApplicationWithBundleIdentifier:)
func (a NSApplication) YieldActivationToApplicationWithBundleIdentifier(bundleIdentifier string) {
	objc.Send[objc.ID](a.ID, objc.Sel("yieldActivationToApplicationWithBundleIdentifier:"), objc.String(bundleIdentifier))
}

// Disables relaunching the app on login.
//
// # Discussion
//
// Invoking this method will prevent the app from relaunching when the user
// next logs in to their account.
//
// If your app shouldn’t be relaunched because it launches via some other
// mechanism (for example, `launchd`), then the recommended usage is to call
// this method once, and never pair it with an
// [NSApplication.EnableRelaunchOnLogin] method.
//
// If your app shouldn’t be relaunched because it triggers a restart, for
// example an installer, then the recommended usage is to invoke this method
// immediately before you attempt to trigger a restart, and
// [NSApplication.EnableRelaunchOnLogin] immediately after. This is because
// the user may cancel restarting; if the user later restarts for another
// reason, then your app should be brought back.
//
// This methods is thread safe.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/disableRelaunchOnLogin()
func (a NSApplication) DisableRelaunchOnLogin() {
	objc.Send[objc.ID](a.ID, objc.Sel("disableRelaunchOnLogin"))
}

// Enables relaunching the app on login.
//
// # Discussion
//
// Invoking this method will cause the app to relaunch when the user next logs
// in to their account.
//
// This methods is thread safe.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/enableRelaunchOnLogin()
func (a NSApplication) EnableRelaunchOnLogin() {
	objc.Send[objc.ID](a.ID, objc.Sel("enableRelaunchOnLogin"))
}

// Register for notifications sent by Apple Push Notification service (APNs).
//
// # Discussion
//
// Call this method to register your app with APNs. When a valid connection is
// established, APNs sends a device token to your app delegate. Forward that
// token to your company’s provider server.
//
// For more information about how to register with APNs, see [Registering your
// app with APNs].
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/registerForRemoteNotifications()
//
// [Registering your app with APNs]: https://developer.apple.com/documentation/UserNotifications/registering-your-app-with-apns
func (a NSApplication) RegisterForRemoteNotifications() {
	objc.Send[objc.ID](a.ID, objc.Sel("registerForRemoteNotifications"))
}

// Unregister for notifications received from Apple Push Notification service.
//
// # Discussion
//
// You should only call this method in rare circumstances, such as when a new
// version of the app drops support for remote notifications. Apps
// unregistered through this method can always reregister.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/unregisterForRemoteNotifications()
func (a NSApplication) UnregisterForRemoteNotifications() {
	objc.Send[objc.ID](a.ID, objc.Sel("unregisterForRemoteNotifications"))
}

// Register to receive notifications of the specified types from a provider
// through the Apple Push Notification service.
//
// types: A bit mask specifying the types of notifications the app accepts. See
// [NSApplication.RemoteNotificationType] for valid bit-mask values.
//
// # Discussion
//
// When you send this message, the device initiates the registration process
// with Apple Push Notification Service. If it succeeds, the app delegate
// receives a device token in the
// [ApplicationDidRegisterForRemoteNotificationsWithDeviceToken] method; if
// registration doesn’t succeed, the delegate is informed via the
// [ApplicationDidFailToRegisterForRemoteNotificationsWithError] method. If
// the app delegate receives a device token, it should connect with its
// provider and pass it the token.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/registerForRemoteNotifications(matching:)
//
// [NSApplication.RemoteNotificationType]: https://developer.apple.com/documentation/AppKit/NSApplication/RemoteNotificationType
func (a NSApplication) RegisterForRemoteNotificationTypes(types NSRemoteNotificationType) {
	objc.Send[objc.ID](a.ID, objc.Sel("registerForRemoteNotificationTypes:"), types)
}

// Show or hides the interface for customizing the Touch Bar.
//
// # Discussion
//
// You can call this method yourself when you want to show or hide the Touch
// Bar customization interface.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/toggleTouchBarCustomizationPalette(_:)
func (a NSApplication) ToggleTouchBarCustomizationPalette(sender objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("toggleTouchBarCustomizationPalette:"), sender)
}

// Starts a user attention request.
//
// requestType: The severity of the request. For a list of possible values, see
// [NSApplication.RequestUserAttentionType].
//
// # Return Value
//
// The identifier for the request. You can use this value to cancel the
// request later using the [NSApplication.CancelUserAttentionRequest] method.
//
// # Discussion
//
// Activating the app cancels the user attention request. A spoken
// notification will occur if spoken notifications are enabled. Sending
// [NSApplication.RequestUserAttention] to an app that is already active has
// no effect.
//
// If the inactive app presents a modal panel, this method will be invoked
// with [NSCriticalRequest] automatically. The modal panel is not brought to
// the front for an inactive app.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/requestUserAttention(_:)
//
// [NSApplication.RequestUserAttentionType]: https://developer.apple.com/documentation/AppKit/NSApplication/RequestUserAttentionType
func (a NSApplication) RequestUserAttention(requestType NSRequestUserAttentionType) int {
	rv := objc.Send[int](a.ID, objc.Sel("requestUserAttention:"), requestType)
	return rv
}

// Cancels a previous user attention request.
//
// request: The request identifier returned by the [NSApplication.RequestUserAttention]
// method.
//
// # Discussion
//
// A request is also canceled automatically by user activation of the app.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/cancelUserAttentionRequest(_:)
func (a NSApplication) CancelUserAttentionRequest(request int) {
	objc.Send[objc.ID](a.ID, objc.Sel("cancelUserAttentionRequest:"), request)
}

// Handles errors that might occur when the user attempts to open or print
// files.
//
// reply: The error that occurred. For a list of possible values, see
// [NSApplication.DelegateReply].
//
// # Discussion
//
// Delegates should invoke this method if an error is encountered in the
// [ApplicationOpenFiles] or
// [ApplicationPrintFilesWithSettingsShowPrintPanels] delegate methods.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/reply(toOpenOrPrint:)
//
// [NSApplication.DelegateReply]: https://developer.apple.com/documentation/AppKit/NSApplication/DelegateReply
func (a NSApplication) ReplyToOpenOrPrint(reply NSApplicationDelegateReply) {
	objc.Send[objc.ID](a.ID, objc.Sel("replyToOpenOrPrint:"), reply)
}

// Register an object that provides help data to your app.
//
// handler: The class instance that conforms to [NSUserInterfaceItemSearching] and
// provides help content.
//
// # Discussion
//
// You can register as many search handlers as you like. If you register the
// same instance more than once the subsequent registrations are ignored.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/registerUserInterfaceItemSearchHandler(_:)
func (a NSApplication) RegisterUserInterfaceItemSearchHandler(handler NSUserInterfaceItemSearching) {
	objc.Send[objc.ID](a.ID, objc.Sel("registerUserInterfaceItemSearchHandler:"), handler)
}

// Searches for the string in the user interface.
//
// searchString: The search string.
//
// stringToSearch: The string to search.
//
// searchRange: The subrange of the `stringToSearch` to restrict the search to.
//
// foundRange: Returns, by-reference, the range of the `searchString` within
// `stringToSearch`.
//
// # Return Value
//
// true if the searchString is matched, otherwise false.
//
// # Discussion
//
// The search uses the default matching rules for Spotlight for Help.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/searchString(_:inUserInterfaceItemString:range:found:)
func (a NSApplication) SearchStringInUserInterfaceItemStringSearchRangeFoundRange(searchString string, stringToSearch string, searchRange foundation.NSRange, foundRange *foundation.NSRange) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("searchString:inUserInterfaceItemString:searchRange:foundRange:"), objc.String(searchString), objc.String(stringToSearch), searchRange, unsafe.Pointer(foundRange))
	return rv
}

// Unregister an object that provides help data to your app.
//
// handler: The class instance that conforms to [NSUserInterfaceItemSearching] and
// provides help content.
//
// # Discussion
//
// If you unregister the same instance more than once the subsequent
// invocations are ignored. Unregistering an instance that was never
// registered is ignored.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/unregisterUserInterfaceItemSearchHandler(_:)
func (a NSApplication) UnregisterUserInterfaceItemSearchHandler(handler NSUserInterfaceItemSearching) {
	objc.Send[objc.ID](a.ID, objc.Sel("unregisterUserInterfaceItemSearchHandler:"), handler)
}

// If your project is properly registered, and the necessary keys have been
// set in the property list, this method launches Help Viewer and displays the
// first page of your app’s help book.
//
// sender: The object that sent the command.
//
// # Discussion
//
// For information on how to set up your project to take advantage of having
// Help Viewer display your help book, see [Specifying the Comprehensive Help
// File].
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/showHelp(_:)
//
// [Specifying the Comprehensive Help File]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/OnlineHelp/Tasks/SpecifyHelpFile.html#//apple_ref/doc/uid/20000020
func (a NSApplication) ShowHelp(sender objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("showHelp:"), sender)
}

// Places the receiver in context-sensitive help mode.
//
// sender: The object that sent the command.
//
// # Discussion
//
// In this mode, the cursor becomes a question mark, and help appears for any
// user interface item the user clicks.
//
// Most apps don’t use this method. Instead, apps enter context-sensitive
// mode when the user presses the Help key. Apps exit context-sensitive help
// mode upon the first event after a help window is displayed.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/activateContextHelpMode(_:)
func (a NSApplication) ActivateContextHelpMode(sender objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("activateContextHelpMode:"), sender)
}

// Hides all apps, except the receiver.
//
// sender: The object that sent this message.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/hideOtherApplications(_:)
func (a NSApplication) HideOtherApplications(sender objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("hideOtherApplications:"), sender)
}

// Unhides all apps, including the receiver.
//
// sender: The object that sent this message.
//
// # Discussion
//
// This action causes each app to order its windows to the front, which could
// obscure the currently active window in the active app.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/unhideAllApplications(_:)
func (a NSApplication) UnhideAllApplications(sender objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("unhideAllApplications:"), sender)
}

// Logs a given exception by calling `NSLog()`.
//
// exception: The exception whose contents you want to write to the log file.
//
// # Discussion
//
// This method doesn’t raise `anException`. Use it inside of an exception
// handler to record that the exception occurred.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/reportException(_:)
func (a NSApplication) ReportException(exception foundation.NSException) {
	objc.Send[objc.ID](a.ID, objc.Sel("reportException:"), exception)
}

// Returns the app’s activation policy.
//
// # Return Value
//
// The app’s current activation policy.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/activationPolicy()
func (a NSApplication) ActivationPolicy() NSApplicationActivationPolicy {
	rv := objc.Send[NSApplicationActivationPolicy](a.ID, objc.Sel("activationPolicy"))
	return NSApplicationActivationPolicy(rv)
}

// Attempts to modify the app’s activation policy.
//
// activationPolicy: The desired activation policy.
//
// # Return Value
//
// true if the policy switch succeded; otherwise, false.
//
// # Discussion
//
// You can set any activation policy in macOS 10.9 and later; in macOS 10.8
// and earlier, you can only set the activation policy to
// [NSApplicationActivationPolicyProhibited] or
// [NSApplicationActivationPolicyRegular].
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/setActivationPolicy(_:)
func (a NSApplication) SetActivationPolicy(activationPolicy NSApplicationActivationPolicy) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("setActivationPolicy:"), activationPolicy)
	return rv
}

// Aborts the event loop started by [NSApplication.RunModalForWindow] or
// [NSApplication.RunModalSession].
//
// # Discussion
//
// When stopped with this method, [NSApplication.RunModalForWindow] and
// [NSApplication.RunModalSession] return [NSModalResponseAbort].
//
// [NSApplication.AbortModal] must be used instead of
// [NSApplication.StopModal] or [NSApplication.StopModalWithCode] when you
// need to stop a modal event loop from anywhere other than a callout from
// that event loop. In other words, if you want to stop the loop in response
// to a user’s actions within the modal window, use
// [NSApplication.StopModal]; otherwise, use [NSApplication.AbortModal]. For
// example, use [NSApplication.AbortModal] when running in a different thread
// from AppKit’s main thread or when responding to an [NSTimer] that you
// have added to the [NSModalPanelRunLoopMode] mode of the default
// [NSRunLoop].
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/abortModal()
func (a NSApplication) AbortModal() {
	objc.Send[objc.ID](a.ID, objc.Sel("abortModal"))
}

// Returns the activation point for the user interface element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityActivationPoint()
func (a NSApplication) AccessibilityActivationPoint() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](a.ID, objc.Sel("accessibilityActivationPoint"))
	return corefoundation.CGPoint(rv)
}

// Returns the allowed values for the slider accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityAllowedValues()
func (a NSApplication) AccessibilityAllowedValues() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("accessibilityAllowedValues"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}

// Returns the child accessibility element with the current focus.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityApplicationFocusedUIElement()
func (a NSApplication) AccessibilityApplicationFocusedUIElement() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityApplicationFocusedUIElement"))
	return objectivec.Object{ID: rv}
}

// Returns the attributed substring for the specified range of characters.
//
// range: The range of characters.
//
// # Return Value
//
// An attributed string representing the specified characters.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityAttributedString(for:)
func (a NSApplication) AccessibilityAttributedStringForRange(range_ foundation.NSRange) foundation.NSAttributedString {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityAttributedStringForRange:"), range_)
	return foundation.NSAttributedStringFromID(rv)
}

// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityAttributedUserInputLabels()
func (a NSApplication) AccessibilityAttributedUserInputLabels() []foundation.NSAttributedString {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("accessibilityAttributedUserInputLabels"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSAttributedString {
		return foundation.NSAttributedStringFromID(id)
	})
}

// Returns the child accessibility element that represents the window’s
// cancel button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityCancelButton()
func (a NSApplication) AccessibilityCancelButton() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityCancelButton"))
	return objectivec.Object{ID: rv}
}

// Returns the cell at the specified column and row.
//
// column: The column index.
//
// row: The row index.
//
// # Return Value
//
// The cell specified by the column and row indexes.
//
// # Discussion
//
// This property is required for all elements that function as cell-based
// tables.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityCell(forColumn:row:)
func (a NSApplication) AccessibilityCellForColumnRow(column int, row int) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityCellForColumn:row:"), column, row)
	return objectivec.Object{ID: rv}
}

// Returns the child accessibility elements in the accessibility hierarchy.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityChildren()
func (a NSApplication) AccessibilityChildren() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityChildren"))
	return foundation.NSArrayFromID(rv)
}

// Returns the array of child accessibility elements in order for linear
// navigation.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityChildrenInNavigationOrder()
func (a NSApplication) AccessibilityChildrenInNavigationOrder() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("accessibilityChildrenInNavigationOrder"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// Returns the clear button for the search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityClearButton()
func (a NSApplication) AccessibilityClearButton() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityClearButton"))
	return objectivec.Object{ID: rv}
}

// Returns the child accessibility element that represents the window’s
// close button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityCloseButton()
func (a NSApplication) AccessibilityCloseButton() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityCloseButton"))
	return objectivec.Object{ID: rv}
}

// Returns the number of columns in the accessibility element’s grid.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityColumnCount()
func (a NSApplication) AccessibilityColumnCount() int {
	rv := objc.Send[int](a.ID, objc.Sel("accessibilityColumnCount"))
	return rv
}

// Returns the column header accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityColumnHeaderUIElements()
func (a NSApplication) AccessibilityColumnHeaderUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityColumnHeaderUIElements"))
	return foundation.NSArrayFromID(rv)
}

// Returns the column index range of the cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityColumnIndexRange()
func (a NSApplication) AccessibilityColumnIndexRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](a.ID, objc.Sel("accessibilityColumnIndexRange"))
	return foundation.NSRange(rv)
}

// Returns the column titles for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityColumnTitles()
func (a NSApplication) AccessibilityColumnTitles() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityColumnTitles"))
	return foundation.NSArrayFromID(rv)
}

// Returns the column accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityColumns()
func (a NSApplication) AccessibilityColumns() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityColumns"))
	return foundation.NSArrayFromID(rv)
}

// Returns the contents of the current accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityContents()
func (a NSApplication) AccessibilityContents() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityContents"))
	return foundation.NSArrayFromID(rv)
}

// Returns the critical value for the level indicator.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityCriticalValue()
func (a NSApplication) AccessibilityCriticalValue() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityCriticalValue"))
	return objectivec.Object{ID: rv}
}

// Returns the custom actions of the current accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityCustomActions()
func (a NSApplication) AccessibilityCustomActions() []NSAccessibilityCustomAction {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("accessibilityCustomActions"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSAccessibilityCustomAction {
		return NSAccessibilityCustomActionFromID(id)
	})
}

// Returns the custom rotors of the current accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityCustomRotors()
func (a NSApplication) AccessibilityCustomRotors() []NSAccessibilityCustomRotor {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("accessibilityCustomRotors"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSAccessibilityCustomRotor {
		return NSAccessibilityCustomRotorFromID(id)
	})
}

// Returns the decrement button for the stepper accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityDecrementButton()
func (a NSApplication) AccessibilityDecrementButton() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityDecrementButton"))
	return objectivec.Object{ID: rv}
}

// Returns the child accessibility element that represents the window’s
// default button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityDefaultButton()
func (a NSApplication) AccessibilityDefaultButton() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityDefaultButton"))
	return objectivec.Object{ID: rv}
}

// Returns the row disclosing the current row.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityDisclosedByRow()
func (a NSApplication) AccessibilityDisclosedByRow() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityDisclosedByRow"))
	return objectivec.Object{ID: rv}
}

// Returns the rows that the current row discloses.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityDisclosedRows()
func (a NSApplication) AccessibilityDisclosedRows() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityDisclosedRows"))
	return objectivec.Object{ID: rv}
}

// Returns the indention level for the row.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityDisclosureLevel()
func (a NSApplication) AccessibilityDisclosureLevel() int {
	rv := objc.Send[int](a.ID, objc.Sel("accessibilityDisclosureLevel"))
	return rv
}

// Returns the URL for the file that the accessibility element represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityDocument()
func (a NSApplication) AccessibilityDocument() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityDocument"))
	return foundation.NSStringFromID(rv).String()
}

// Returns the icon for the app’s menu bar extra.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityExtrasMenuBar()
func (a NSApplication) AccessibilityExtrasMenuBar() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityExtrasMenuBar"))
	return objectivec.Object{ID: rv}
}

// Returns the filename for the file that the accessibility element
// represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityFilename()
func (a NSApplication) AccessibilityFilename() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityFilename"))
	return foundation.NSStringFromID(rv).String()
}

// Returns the child window with the current focus.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityFocusedWindow()
func (a NSApplication) AccessibilityFocusedWindow() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityFocusedWindow"))
	return objectivec.Object{ID: rv}
}

// Returns the accessibility element’s frame in screen coordinates.
//
// # Return Value
//
// The element’s frame in screen coordinates.
//
// # Discussion
//
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [NSWindow.AccessibilityFrame] property. This method is called whenever
// accessibility clients request the [size] or [position] attributes.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityFrame()
//
// [position]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/position
// [size]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/size
func (a NSApplication) AccessibilityFrame() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](a.ID, objc.Sel("accessibilityFrame"))
	return corefoundation.CGRect(rv)
}

// Returns the rectangle that encloses the specified range of characters.
//
// range: The range of characters.
//
// # Return Value
//
// The rectangle that encloses the specified characters.
//
// # Discussion
//
// If the range crosses a line boundary, the returned rectangle fully encloses
// all the lines of characters.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityFrame(for:)
func (a NSApplication) AccessibilityFrameForRange(range_ foundation.NSRange) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](a.ID, objc.Sel("accessibilityFrameForRange:"), range_)
	return corefoundation.CGRect(rv)
}

// Returns the child accessibility element that represents the window’s
// full-screen button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityFullScreenButton()
func (a NSApplication) AccessibilityFullScreenButton() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityFullScreenButton"))
	return objectivec.Object{ID: rv}
}

// Returns the child accessibility element that represents the window’s grow
// area.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityGrowArea()
func (a NSApplication) AccessibilityGrowArea() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityGrowArea"))
	return objectivec.Object{ID: rv}
}

// Returns the drag handle elements for the layout item element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityHandles()
func (a NSApplication) AccessibilityHandles() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityHandles"))
	return foundation.NSArrayFromID(rv)
}

// Returns the header for the table view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityHeader()
func (a NSApplication) AccessibilityHeader() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityHeader"))
	return objectivec.Object{ID: rv}
}

// Returns the help text for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityHelp()
func (a NSApplication) AccessibilityHelp() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityHelp"))
	return foundation.NSStringFromID(rv).String()
}

// Returns the horizontal scroll bar for the scroll view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityHorizontalScrollBar()
func (a NSApplication) AccessibilityHorizontalScrollBar() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityHorizontalScrollBar"))
	return objectivec.Object{ID: rv}
}

// Returns the description of the layout area’s horizontal units.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityHorizontalUnitDescription()
func (a NSApplication) AccessibilityHorizontalUnitDescription() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityHorizontalUnitDescription"))
	return foundation.NSStringFromID(rv).String()
}

// Returns the units that the layout area uses for horizontal values.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityHorizontalUnits()
func (a NSApplication) AccessibilityHorizontalUnits() NSAccessibilityUnits {
	rv := objc.Send[NSAccessibilityUnits](a.ID, objc.Sel("accessibilityHorizontalUnits"))
	return NSAccessibilityUnits(rv)
}

// Returns the accessibility element’s identity.
//
// # Return Value
//
// Returns the unique ID for the accessibility element. It is often used in
// automated testing.
//
// # Discussion
//
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [NSWindow.AccessibilityIdentifier] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityIdentifier()
func (a NSApplication) AccessibilityIdentifier() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityIdentifier"))
	return foundation.NSStringFromID(rv).String()
}

// Returns the increment button for the stepper accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityIncrementButton()
func (a NSApplication) AccessibilityIncrementButton() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityIncrementButton"))
	return objectivec.Object{ID: rv}
}

// Returns the index of the row or column that the accessibility element
// represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityIndex()
func (a NSApplication) AccessibilityIndex() int {
	rv := objc.Send[int](a.ID, objc.Sel("accessibilityIndex"))
	return rv
}

// Returns the line number that contains the insertion point.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityInsertionPointLineNumber()
func (a NSApplication) AccessibilityInsertionPointLineNumber() int {
	rv := objc.Send[int](a.ID, objc.Sel("accessibilityInsertionPointLineNumber"))
	return rv
}

// Returns a short description of the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityLabel()
func (a NSApplication) AccessibilityLabel() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityLabel"))
	return foundation.NSStringFromID(rv).String()
}

// Returns the child label elements for the slider accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityLabelUIElements()
func (a NSApplication) AccessibilityLabelUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityLabelUIElements"))
	return foundation.NSArrayFromID(rv)
}

// Returns the value of the label accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityLabelValue()
func (a NSApplication) AccessibilityLabelValue() float32 {
	rv := objc.Send[float32](a.ID, objc.Sel("accessibilityLabelValue"))
	return rv
}

// Converts the provided point in screen coordinates to a point in the layout
// area’s coordinate system.
//
// point: A point in the screen’s coordinate system.
//
// # Return Value
//
// A point in the layout area’s coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityLayoutPoint(forScreenPoint:)
func (a NSApplication) AccessibilityLayoutPointForScreenPoint(point corefoundation.CGPoint) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](a.ID, objc.Sel("accessibilityLayoutPointForScreenPoint:"), point)
	return corefoundation.CGPoint(rv)
}

// Converts the provided size in screen coordinates to a size in the layout
// area’s coordinate system.
//
// size: A size in the screen’s coordinate system.
//
// # Return Value
//
// A size in the layout area’s coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityLayoutSize(forScreenSize:)
func (a NSApplication) AccessibilityLayoutSizeForScreenSize(size corefoundation.CGSize) corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](a.ID, objc.Sel("accessibilityLayoutSizeForScreenSize:"), size)
	return corefoundation.CGSize(rv)
}

// Returns the line number for the line that contains the specified character
// index.
//
// index: The index for a character.
//
// # Return Value
//
// The line number for the line holding the specified character index.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityLine(for:)
func (a NSApplication) AccessibilityLineForIndex(index int) int {
	rv := objc.Send[int](a.ID, objc.Sel("accessibilityLineForIndex:"), index)
	return rv
}

// Returns the elements that have links with the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityLinkedUIElements()
func (a NSApplication) AccessibilityLinkedUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityLinkedUIElements"))
	return foundation.NSArrayFromID(rv)
}

// Returns the app’s main window.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMainWindow()
func (a NSApplication) AccessibilityMainWindow() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityMainWindow"))
	return objectivec.Object{ID: rv}
}

// Returns the user interface element that functions as a marker group for the
// ruler accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMarkerGroupUIElement()
func (a NSApplication) AccessibilityMarkerGroupUIElement() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityMarkerGroupUIElement"))
	return objectivec.Object{ID: rv}
}

// Returns the human-readable description of the marker type.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMarkerTypeDescription()
func (a NSApplication) AccessibilityMarkerTypeDescription() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityMarkerTypeDescription"))
	return foundation.NSStringFromID(rv).String()
}

// Returns the array of marker accessibility elements for the ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMarkerUIElements()
func (a NSApplication) AccessibilityMarkerUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityMarkerUIElements"))
	return foundation.NSArrayFromID(rv)
}

// Returns the marker values for the ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMarkerValues()
func (a NSApplication) AccessibilityMarkerValues() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityMarkerValues"))
	return objectivec.Object{ID: rv}
}

// Returns the maximum value for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMaxValue()
func (a NSApplication) AccessibilityMaxValue() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityMaxValue"))
	return objectivec.Object{ID: rv}
}

// Returns the app’s menu bar.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMenuBar()
func (a NSApplication) AccessibilityMenuBar() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityMenuBar"))
	return objectivec.Object{ID: rv}
}

// Returns the minimum value for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMinValue()
func (a NSApplication) AccessibilityMinValue() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityMinValue"))
	return objectivec.Object{ID: rv}
}

// Returns the child accessibility element that represents the window’s
// minimize button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMinimizeButton()
func (a NSApplication) AccessibilityMinimizeButton() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityMinimizeButton"))
	return objectivec.Object{ID: rv}
}

// Returns the contents that follow the divider accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityNextContents()
func (a NSApplication) AccessibilityNextContents() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityNextContents"))
	return foundation.NSArrayFromID(rv)
}

// Returns the number of characters in the text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityNumberOfCharacters()
func (a NSApplication) AccessibilityNumberOfCharacters() int {
	rv := objc.Send[int](a.ID, objc.Sel("accessibilityNumberOfCharacters"))
	return rv
}

// Returns the orientation of the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityOrientation()
func (a NSApplication) AccessibilityOrientation() NSAccessibilityOrientation {
	rv := objc.Send[NSAccessibilityOrientation](a.ID, objc.Sel("accessibilityOrientation"))
	return NSAccessibilityOrientation(rv)
}

// Returns the overflow button for the toolbar.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityOverflowButton()
func (a NSApplication) AccessibilityOverflowButton() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityOverflowButton"))
	return objectivec.Object{ID: rv}
}

// Returns the accessibility element’s parent in the accessibility
// hierarchy.
//
// # Return Value
//
// The element’s parent in the accessibility hierarchy.
//
// # Discussion
//
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [NSWindow.AccessibilityParent] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityParent()
func (a NSApplication) AccessibilityParent() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityParent"))
	return objectivec.Object{ID: rv}
}

// Cancels the current operation.
//
// # Return Value
//
// true if the action was successfully triggered; otherwise, false. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformCancel()
func (a NSApplication) AccessibilityPerformCancel() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("accessibilityPerformCancel"))
	return rv
}

// Simulates pressing Return in the accessibility element.
//
// # Return Value
//
// true if the action was successfully triggered; otherwise, false. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// # Discussion
//
// Use this method on elements that take keyboard input, such as a text field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformConfirm()
func (a NSApplication) AccessibilityPerformConfirm() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("accessibilityPerformConfirm"))
	return rv
}

// Decrements the accessibility element’s value.
//
// # Return Value
//
// true if the action was successfully triggered; otherwise, false. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// # Discussion
//
// Use this method on elements that have an adjustable
// [NSWindow.AccessibilityValue] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformDecrement()
func (a NSApplication) AccessibilityPerformDecrement() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("accessibilityPerformDecrement"))
	return rv
}

// Deletes the accessibility element’s value.
//
// # Return Value
//
// true if the action was successfully triggered; otherwise, false. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// # Discussion
//
// Use this method on elements with values.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformDelete()
func (a NSApplication) AccessibilityPerformDelete() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("accessibilityPerformDelete"))
	return rv
}

// Increments the accessibility element’s value.
//
// # Return Value
//
// true if the action was successfully triggered; otherwise, false. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// # Discussion
//
// Use this method on elements that have an adjustable
// [NSWindow.AccessibilityValue] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformIncrement()
func (a NSApplication) AccessibilityPerformIncrement() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("accessibilityPerformIncrement"))
	return rv
}

// Selects the accessibility element.
//
// # Return Value
//
// true if the action was successfully triggered; otherwise, false. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// # Discussion
//
// Use this method on selectable elements, such as a menu item.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformPick()
func (a NSApplication) AccessibilityPerformPick() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("accessibilityPerformPick"))
	return rv
}

// Simulates clicking the accessibility element.
//
// # Return Value
//
// true if the action was successfully triggered; otherwise, false. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// # Discussion
//
// Use this method on elements that behave like buttons.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformPress()
func (a NSApplication) AccessibilityPerformPress() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("accessibilityPerformPress"))
	return rv
}

// Brings the window to the front.
//
// # Return Value
//
// true if the action was successfully triggered; otherwise, false. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// # Discussion
//
// The window behaves as if you had clicked on the window’s title bar.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformRaise()
func (a NSApplication) AccessibilityPerformRaise() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("accessibilityPerformRaise"))
	return rv
}

// Displays the accessibility element’s alternative UI.
//
// # Return Value
//
// true if the action was successfully triggered; otherwise, false. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// # Discussion
//
// Use this method to trigger changes to the UI due to a mouse-hover or
// similar event.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformShowAlternateUI()
func (a NSApplication) AccessibilityPerformShowAlternateUI() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("accessibilityPerformShowAlternateUI"))
	return rv
}

// Returns to the accessibility element’s original UI.
//
// # Return Value
//
// true if the action was successfully triggered; otherwise, false. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// # Discussion
//
// Call this method after successfully calling
// [AccessibilityPerformShowAlternateUI] to return to the original UI.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformShowDefaultUI()
func (a NSApplication) AccessibilityPerformShowDefaultUI() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("accessibilityPerformShowDefaultUI"))
	return rv
}

// Displays the menu accessibility element.
//
// # Return Value
//
// true if the action was successfully triggered; otherwise, false. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// # Discussion
//
// Use this method to display the contextual menu for the element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformShowMenu()
func (a NSApplication) AccessibilityPerformShowMenu() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("accessibilityPerformShowMenu"))
	return rv
}

// Returns the placeholder value for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPlaceholderValue()
func (a NSApplication) AccessibilityPlaceholderValue() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityPlaceholderValue"))
	return foundation.NSStringFromID(rv).String()
}

// Returns the contents that precede the divider accessibility element.
//
// # Return Value
//
// Sets the contents preceding this divider element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPreviousContents()
func (a NSApplication) AccessibilityPreviousContents() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityPreviousContents"))
	return foundation.NSArrayFromID(rv)
}

// Returns the child accessibility element that represents the window’s
// proxy icon.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityProxy()
func (a NSApplication) AccessibilityProxy() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityProxy"))
	return objectivec.Object{ID: rv}
}

// Returns the rich text format (RTF) data that describes the specified range
// of characters.
//
// range: The range of characters.
//
// # Return Value
//
// A data object containing an RTF representation of the specified characters.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRTF(for:)
func (a NSApplication) AccessibilityRTFForRange(range_ foundation.NSRange) foundation.NSData {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityRTFForRange:"), range_)
	return foundation.NSDataFromID(rv)
}

// Returns the range of characters for the glyph that includes the specified
// character.
//
// index: The specified character.
//
// # Return Value
//
// The range of characters for the glyph.
//
// # Discussion
//
// This value always includes the specified character but may include
// additional characters if that character is part of a multicharacter glyph.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRange(for:)-6kv3
func (a NSApplication) AccessibilityRangeForIndex(index int) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](a.ID, objc.Sel("accessibilityRangeForIndex:"), index)
	return foundation.NSRange(rv)
}

// Returns the range of characters in the specified line.
//
// line: The line number to be examined.
//
// # Return Value
//
// The range of characters for the specified line number. If the line ends
// with a newline character, including the newline is preferred.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRange(forLine:)
func (a NSApplication) AccessibilityRangeForLine(line int) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](a.ID, objc.Sel("accessibilityRangeForLine:"), line)
	return foundation.NSRange(rv)
}

// Returns the range of characters for the glyph at the specified point.
//
// point: A point in screen coordinates.
//
// # Return Value
//
// The range of characters that make up the glyph at the given point.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRange(for:)-1iudm
func (a NSApplication) AccessibilityRangeForPosition(point corefoundation.CGPoint) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](a.ID, objc.Sel("accessibilityRangeForPosition:"), point)
	return foundation.NSRange(rv)
}

// Returns the type of interface element that the accessibility element
// represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRole()
func (a NSApplication) AccessibilityRole() NSAccessibilityRole {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityRole"))
	return NSAccessibilityRole(foundation.NSStringFromID(rv).String())
}

// Returns a localized, human-intelligible description of the accessibility
// element’s role, such as radio button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRoleDescription()
func (a NSApplication) AccessibilityRoleDescription() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityRoleDescription"))
	return foundation.NSStringFromID(rv).String()
}

// Returns the number of rows in the accessibility element’s grid.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRowCount()
func (a NSApplication) AccessibilityRowCount() int {
	rv := objc.Send[int](a.ID, objc.Sel("accessibilityRowCount"))
	return rv
}

// Returns the row header accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRowHeaderUIElements()
func (a NSApplication) AccessibilityRowHeaderUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityRowHeaderUIElements"))
	return foundation.NSArrayFromID(rv)
}

// Returns the row index range of the cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRowIndexRange()
func (a NSApplication) AccessibilityRowIndexRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](a.ID, objc.Sel("accessibilityRowIndexRange"))
	return foundation.NSRange(rv)
}

// Returns the row accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRows()
func (a NSApplication) AccessibilityRows() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityRows"))
	return foundation.NSArrayFromID(rv)
}

// Returns the type of markers for the ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRulerMarkerType()
func (a NSApplication) AccessibilityRulerMarkerType() NSAccessibilityRulerMarkerType {
	rv := objc.Send[NSAccessibilityRulerMarkerType](a.ID, objc.Sel("accessibilityRulerMarkerType"))
	return NSAccessibilityRulerMarkerType(rv)
}

// Converts the provided point in the layout area’s coordinates to a point
// in the screen’s coordinate system.
//
// point: A point in the layout area’s coordinate system.
//
// # Return Value
//
// A point in the screen’s coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityScreenPoint(forLayoutPoint:)
func (a NSApplication) AccessibilityScreenPointForLayoutPoint(point corefoundation.CGPoint) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](a.ID, objc.Sel("accessibilityScreenPointForLayoutPoint:"), point)
	return corefoundation.CGPoint(rv)
}

// Converts the provided size in the layout area’s coordinates to a size in
// the screen’s coordinate system.
//
// size: A size in the layout area’s coordinate system.
//
// # Return Value
//
// A size in the screen’s coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityScreenSize(forLayoutSize:)
func (a NSApplication) AccessibilityScreenSizeForLayoutSize(size corefoundation.CGSize) corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](a.ID, objc.Sel("accessibilityScreenSizeForLayoutSize:"), size)
	return corefoundation.CGSize(rv)
}

// Returns the search button for the search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySearchButton()
func (a NSApplication) AccessibilitySearchButton() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilitySearchButton"))
	return objectivec.Object{ID: rv}
}

// Returns the search menu for the search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySearchMenu()
func (a NSApplication) AccessibilitySearchMenu() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilitySearchMenu"))
	return objectivec.Object{ID: rv}
}

// Returns the currently selected cells for the table.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySelectedCells()
func (a NSApplication) AccessibilitySelectedCells() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilitySelectedCells"))
	return foundation.NSArrayFromID(rv)
}

// Returns the accessibility element’s currently selected children.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySelectedChildren()
func (a NSApplication) AccessibilitySelectedChildren() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilitySelectedChildren"))
	return foundation.NSArrayFromID(rv)
}

// Returns the currently selected columns for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySelectedColumns()
func (a NSApplication) AccessibilitySelectedColumns() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilitySelectedColumns"))
	return foundation.NSArrayFromID(rv)
}

// Returns the currently selected rows for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySelectedRows()
func (a NSApplication) AccessibilitySelectedRows() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilitySelectedRows"))
	return foundation.NSArrayFromID(rv)
}

// Returns the currently selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySelectedText()
func (a NSApplication) AccessibilitySelectedText() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilitySelectedText"))
	return foundation.NSStringFromID(rv).String()
}

// Returns the range of the currently selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySelectedTextRange()
func (a NSApplication) AccessibilitySelectedTextRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](a.ID, objc.Sel("accessibilitySelectedTextRange"))
	return foundation.NSRange(rv)
}

// Returns an array of ranges for the currently selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySelectedTextRanges()
func (a NSApplication) AccessibilitySelectedTextRanges() []foundation.NSValue {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("accessibilitySelectedTextRanges"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSValue {
		return foundation.NSValueFromID(id)
	})
}

// Returns the list of elements that the accessibility element is a title for.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityServesAsTitleForUIElements()
func (a NSApplication) AccessibilityServesAsTitleForUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityServesAsTitleForUIElements"))
	return foundation.NSArrayFromID(rv)
}

// Returns the range of characters that the accessibility element displays.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySharedCharacterRange()
func (a NSApplication) AccessibilitySharedCharacterRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](a.ID, objc.Sel("accessibilitySharedCharacterRange"))
	return foundation.NSRange(rv)
}

// Returns the array of elements that shares the keyboard focus with the
// accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySharedFocusElements()
func (a NSApplication) AccessibilitySharedFocusElements() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilitySharedFocusElements"))
	return foundation.NSArrayFromID(rv)
}

// Returns the other elements that share text with the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySharedTextUIElements()
func (a NSApplication) AccessibilitySharedTextUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilitySharedTextUIElements"))
	return foundation.NSArrayFromID(rv)
}

// Returns the menu currently displaying for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityShownMenu()
func (a NSApplication) AccessibilityShownMenu() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityShownMenu"))
	return objectivec.Object{ID: rv}
}

// Returns the accessibility element’s sort direction.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySortDirection()
func (a NSApplication) AccessibilitySortDirection() NSAccessibilitySortDirection {
	rv := objc.Send[NSAccessibilitySortDirection](a.ID, objc.Sel("accessibilitySortDirection"))
	return NSAccessibilitySortDirection(rv)
}

// Returns an array that contains the views and splitter bar from the split
// view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySplitters()
func (a NSApplication) AccessibilitySplitters() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilitySplitters"))
	return foundation.NSArrayFromID(rv)
}

// Returns the substring for the specified range.
//
// range: A range of characters contained by the element.
//
// # Return Value
//
// The substring specified by the given range.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityString(for:)
func (a NSApplication) AccessibilityStringForRange(range_ foundation.NSRange) string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityStringForRange:"), range_)
	return foundation.NSStringFromID(rv).String()
}

// Returns a range of characters that all have the same style as the specified
// character.
//
// index: The index of the specified character.
//
// # Return Value
//
// A range of characters with the same style as the specified character.
//
// # Discussion
//
// This method returns a range of characters that meet two conditions: The
// range must include the specified character, and all the other characters in
// the range must match the specified character’s style. If none of the
// adjacent characters match the specified character’s style, the method
// returns only the specified character.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityStyleRange(for:)
func (a NSApplication) AccessibilityStyleRangeForIndex(index int) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](a.ID, objc.Sel("accessibilityStyleRangeForIndex:"), index)
	return foundation.NSRange(rv)
}

// Returns the specialized interface element type that the accessibility
// element represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySubrole()
func (a NSApplication) AccessibilitySubrole() NSAccessibilitySubrole {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilitySubrole"))
	return NSAccessibilitySubrole(foundation.NSStringFromID(rv).String())
}

// Returns the tab accessibility elements for the tab view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityTabs()
func (a NSApplication) AccessibilityTabs() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityTabs"))
	return foundation.NSArrayFromID(rv)
}

// Returns the title of the accessibility element—for example, a button’s
// visible text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityTitle()
func (a NSApplication) AccessibilityTitle() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityTitle"))
	return foundation.NSStringFromID(rv).String()
}

// Returns the static text element that represents the accessibility
// element’s title.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityTitleUIElement()
func (a NSApplication) AccessibilityTitleUIElement() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityTitleUIElement"))
	return objectivec.Object{ID: rv}
}

// Returns the child accessibility element that represents the window’s
// toolbar button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityToolbarButton()
func (a NSApplication) AccessibilityToolbarButton() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityToolbarButton"))
	return objectivec.Object{ID: rv}
}

// Returns the top-level element that contains the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityTopLevelUIElement()
func (a NSApplication) AccessibilityTopLevelUIElement() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityTopLevelUIElement"))
	return objectivec.Object{ID: rv}
}

// Returns the URL for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityURL()
func (a NSApplication) AccessibilityURL() foundation.NSURL {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityURL"))
	return foundation.NSURLFromID(rv)
}

// Returns the human-readable description of the ruler’s units.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityUnitDescription()
func (a NSApplication) AccessibilityUnitDescription() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityUnitDescription"))
	return foundation.NSStringFromID(rv).String()
}

// Returns the units for the ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityUnits()
func (a NSApplication) AccessibilityUnits() NSAccessibilityUnits {
	rv := objc.Send[NSAccessibilityUnits](a.ID, objc.Sel("accessibilityUnits"))
	return NSAccessibilityUnits(rv)
}

// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityUserInputLabels()
func (a NSApplication) AccessibilityUserInputLabels() []string {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("accessibilityUserInputLabels"))
	return objc.ConvertSliceToStrings(rv)
}

// Returns the accessibility element’s value.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityValue()
func (a NSApplication) AccessibilityValue() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityValue"))
	return objectivec.Object{ID: rv}
}

// Returns the human-readable description of the accessibility element’s
// value.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityValueDescription()
func (a NSApplication) AccessibilityValueDescription() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityValueDescription"))
	return foundation.NSStringFromID(rv).String()
}

// Returns the vertical scroll bar for the scroll view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityVerticalScrollBar()
func (a NSApplication) AccessibilityVerticalScrollBar() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityVerticalScrollBar"))
	return objectivec.Object{ID: rv}
}

// Returns the description of the layout area’s vertical units.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityVerticalUnitDescription()
func (a NSApplication) AccessibilityVerticalUnitDescription() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityVerticalUnitDescription"))
	return foundation.NSStringFromID(rv).String()
}

// Returns the units that the layout area uses for vertical values.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityVerticalUnits()
func (a NSApplication) AccessibilityVerticalUnits() NSAccessibilityUnits {
	rv := objc.Send[NSAccessibilityUnits](a.ID, objc.Sel("accessibilityVerticalUnits"))
	return NSAccessibilityUnits(rv)
}

// Returns the visible cells for the table.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityVisibleCells()
func (a NSApplication) AccessibilityVisibleCells() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityVisibleCells"))
	return foundation.NSArrayFromID(rv)
}

// Returns the range of visible characters in the document.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityVisibleCharacterRange()
func (a NSApplication) AccessibilityVisibleCharacterRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](a.ID, objc.Sel("accessibilityVisibleCharacterRange"))
	return foundation.NSRange(rv)
}

// Returns the accessibility element’s visible child accessibility elements.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityVisibleChildren()
func (a NSApplication) AccessibilityVisibleChildren() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityVisibleChildren"))
	return foundation.NSArrayFromID(rv)
}

// Returns the visible columns for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityVisibleColumns()
func (a NSApplication) AccessibilityVisibleColumns() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityVisibleColumns"))
	return foundation.NSArrayFromID(rv)
}

// Returns the visible rows for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityVisibleRows()
func (a NSApplication) AccessibilityVisibleRows() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityVisibleRows"))
	return foundation.NSArrayFromID(rv)
}

// Returns the warning value for the level indicator.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityWarningValue()
func (a NSApplication) AccessibilityWarningValue() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityWarningValue"))
	return objectivec.Object{ID: rv}
}

// Returns the window that contains the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityWindow()
func (a NSApplication) AccessibilityWindow() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityWindow"))
	return objectivec.Object{ID: rv}
}

// Returns an array that contains all the app’s windows.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityWindows()
func (a NSApplication) AccessibilityWindows() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityWindows"))
	return foundation.NSArrayFromID(rv)
}

// Returns the child accessibility element that represents the window’s zoom
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityZoomButton()
func (a NSApplication) AccessibilityZoomButton() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("accessibilityZoomButton"))
	return objectivec.Object{ID: rv}
}

// Adds an item to the Window menu for a given window.
//
// win: The window being added to the menu. If this window object already exists in
// the Window menu, this method has no effect.
//
// string: The string to display for the window’s menu item. How the string is
// interpreted is dependent on the value in the `isFilename` parameter.
//
// isFilename: If false, `aString` appears literally in the menu; otherwise, `aString` is
// assumed to be a converted pathname with the name of the file preceding the
// path (the way the [NSWindow] method
// [NSWindow.SetTitleWithRepresentedFilename] shows a title)
//
// # Discussion
//
// You rarely need to invoke this method directly because Cocoa places an item
// in the Window menu automatically whenever you set the title of an
// [NSWindow] object.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/addWindowsItem(_:title:filename:)
func (a NSApplication) AddWindowsItemTitleFilename(win INSWindow, string_ string, isFilename bool) {
	objc.Send[objc.ID](a.ID, objc.Sel("addWindowsItem:title:filename:"), win, objc.String(string_), isFilename)
}

// Arranges windows listed in the Window menu in front of all other windows.
//
// sender: The object that sent the command.
//
// # Discussion
//
// Windows associated with the app but not listed in the Window menu are not
// ordered to the front.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/arrangeInFront(_:)
func (a NSApplication) ArrangeInFront(sender objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("arrangeInFront:"), sender)
}

// Sets up a modal session with the given window and returns a pointer to the
// [NSModalSession] structure representing the session.
//
// window: The window for the session.
//
// # Return Value
//
// A pointer to the [NSModalSession] structure that represents the session.
//
// # Discussion
//
// In a modal session, the app receives mouse events only if they occur in
// `aWindow`. The window is made key, and if not already visible is placed
// onscreen using the [NSWindow] method [NSWindow.Center].
//
// The [NSApplication.BeginModalSessionForWindow] method only sets up the
// modal session. To actually run the session, use
// [NSApplication.RunModalSession]. [NSApplication.BeginModalSessionForWindow]
// should be balanced by [NSApplication.EndModalSession]. Make sure these two
// messages are sent within the same exception-handling scope. That is, if you
// send [NSApplication.BeginModalSessionForWindow] inside an `NS_DURING`
// construct, you must send [NSApplication.EndModalSession] before
// `NS_ENDHANDLER`.
//
// If an exception is raised, [NSApplication.BeginModalSessionForWindow]
// arranges for proper cleanup. Do not use `NS_DURING` constructs to send an
// [NSApplication.EndModalSession] message in the event of an exception.
//
// A loop using these methods is similar to a modal event loop run with
// [NSApplication.RunModalForWindow], except the app can continue processing
// between method invocations.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/beginModalSession(for:)
func (a NSApplication) BeginModalSessionForWindow(window INSWindow) NSModalSession {
	rv := objc.Send[NSModalSession](a.ID, objc.Sel("beginModalSessionForWindow:"), window)
	return NSModalSession(rv)
}

// Changes the item for a given window in the Window menu to a given string.
//
// win: The window whose title you want to change in the Window menu. If `aWindow`
// is not in the Window menu, this method adds it.
//
// string: The string to display for the window’s menu item. How the string is
// interpreted is dependent on the value in the `isFilename` parameter.
//
// isFilename: If false, `aString` appears literally in the menu; otherwise, `aString` is
// assumed to be a converted pathname with the name of the file preceding the
// path (the way the [NSWindow] method
// [NSWindow.SetTitleWithRepresentedFilename] shows a title)
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/changeWindowsItem(_:title:filename:)
func (a NSApplication) ChangeWindowsItemTitleFilename(win INSWindow, string_ string, isFilename bool) {
	objc.Send[objc.ID](a.ID, objc.Sel("changeWindowsItem:title:filename:"), win, objc.String(string_), isFilename)
}

// Completes the extended state restoration.
//
// # Discussion
//
// This method informs the app that the extended state restoration is
// completed for the balancing .
//
// If a window has some state that may take a long time to restore, such as a
// web page, you may use this method and methods to `completeStateRestoration`
// to extend the period of this crash protection beyond the default.
//
// You call [NSApplication.ExtendStateRestoration] within your implementation
// of [NSApplication.RestoreWindowWithIdentifierStateCompletionHandler]. You
// would then call `completeStateRestoration` some time after the window is
// fully restored. If the app crashes in the interim, then it may offer to
// discard restorable state on the next launch.
//
// The [NSApplication.ExtendStateRestoration] and `completeStateRestoration`
// method act as a counter. Each call to
// [NSApplication.ExtendStateRestoration]increments the counter, and must be
// matched with a corresponding call to `completeStateRestoration` which
// decrements it. When the counter reaches zero, the app is considered to have
// been fully restored, and any further calls are silently ignored.
//
// This method is thread safe.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/completeStateRestoration()
func (a NSApplication) CompleteStateRestoration() {
	objc.Send[objc.ID](a.ID, objc.Sel("completeStateRestoration"))
}

// Finishes a modal session.
//
// session: A modal session structure returned by a previous invocation of
// [NSApplication.BeginModalSessionForWindow].
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/endModalSession(_:)
func (a NSApplication) EndModalSession(session NSModalSession) {
	objc.Send[objc.ID](a.ID, objc.Sel("endModalSession:"), session)
}

// Executes a block for each of the app’s windows.
//
// options: A constant that indicates window ordering. See
// [NSApplication.WindowListOptions] for possible values.
//
// block: The block to execute for each window. The block takes the following
// parameters:
//
// window: The window for which to execute the block. stop: A Boolean value
// that stops the enumeration early when set to true (the default value is
// false).
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/enumerateWindows(options:using:)
//
// [NSApplication.WindowListOptions]: https://developer.apple.com/documentation/AppKit/NSApplication/WindowListOptions
func (a NSApplication) EnumerateWindowsWithOptionsUsingBlock(options NSWindowListOptions, block WindowBoolHandler) {
	_block1, _cleanup1 := NewWindowBoolBlock(block)
	defer _cleanup1()
	objc.Send[objc.ID](a.ID, objc.Sel("enumerateWindowsWithOptions:usingBlock:"), options, _block1)
}

// Allows an app to extend its state restoration period.
//
// # Discussion
//
// This method allows an app to extend the state restoration period beyond the
// usual. For example, the app crashes before state restoration is complete,
// then it may offer to discard restorable state on the next launch.
//
// If a window has some state that may take a long time to restore, such as a
// web page, you may use this method and methods to
// [NSApplication.CompleteStateRestoration] to extend the period of this crash
// protection beyond the default.
//
// You call `extendStateRestoration` within your implementation of
// [NSApplication.RestoreWindowWithIdentifierStateCompletionHandler]. You
// would then call [NSApplication.CompleteStateRestoration] some time after
// the window is fully restored. If the app crashes in the interim, then it
// may offer to discard restorable state on the next launch.
//
// The `extendStateRestoration` and [NSApplication.CompleteStateRestoration]
// methods act as a counter. Each call to `extendStateRestoration` increments
// the counter, and must be matched with a corresponding call to
// [NSApplication.CompleteStateRestoration] which decrements it. When the
// counter reaches zero, the app is considered to have been fully restored,
// and any further calls are silently ignored.
//
// This method is thread safe.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/extendStateRestoration()
func (a NSApplication) ExtendStateRestoration() {
	objc.Send[objc.ID](a.ID, objc.Sel("extendStateRestoration"))
}

// Hides all the receiver’s windows, and the next app in line is activated.
//
// sender: The object that sent the command.
//
// # Discussion
//
// This method is usually invoked when the user chooses Hide in the app’s
// main menu. When this method begins, it posts an [willHideNotification] to
// the default notification center. When it completes successfully, it posts
// an [didHideNotification].
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/hide(_:)
//
// [didHideNotification]: https://developer.apple.com/documentation/AppKit/NSApplication/didHideNotification
// [willHideNotification]: https://developer.apple.com/documentation/AppKit/NSApplication/willHideNotification
func (a NSApplication) Hide(sender objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("hide:"), sender)
}

// Returns the Boolean value that determines whether the accessibility
// element’s alternative UI is currently visible.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityAlternateUIVisible()
func (a NSApplication) IsAccessibilityAlternateUIVisible() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isAccessibilityAlternateUIVisible"))
	return rv
}

// Returns a Boolean value that determines whether the row is disclosing other
// rows.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityDisclosed()
func (a NSApplication) IsAccessibilityDisclosed() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isAccessibilityDisclosed"))
	return rv
}

// Returns a Boolean value that indicates whether the accessibility element is
// in an edited state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityEdited()
func (a NSApplication) IsAccessibilityEdited() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isAccessibilityEdited"))
	return rv
}

// Returns a Boolean value that determines whether the accessibility element
// participates in the accessibility hierarchy.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityElement()
func (a NSApplication) IsAccessibilityElement() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isAccessibilityElement"))
	return rv
}

// Returns a Boolean value that determines whether the accessibility element
// responds to user events.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityEnabled()
func (a NSApplication) IsAccessibilityEnabled() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isAccessibilityEnabled"))
	return rv
}

// Returns a Boolean value that determines whether the accessibility element
// is in an expanded state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityExpanded()
func (a NSApplication) IsAccessibilityExpanded() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isAccessibilityExpanded"))
	return rv
}

// Returns a Boolean value that indicates whether the accessibility element
// has the keyboard focus.
//
// # Return Value
//
// true if this element has the keyboard focus; otherwise, false.
//
// # Discussion
//
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [NSWindow.AccessibilityFocused] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/isAccessibilityFocused()
func (a NSApplication) IsAccessibilityFocused() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isAccessibilityFocused"))
	return rv
}

// Returns a Boolean value that determines whether the app is the frontmost
// app.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityFrontmost()
func (a NSApplication) IsAccessibilityFrontmost() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isAccessibilityFrontmost"))
	return rv
}

// Returns a Boolean value that determines whether the app is in a hidden
// state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityHidden()
func (a NSApplication) IsAccessibilityHidden() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isAccessibilityHidden"))
	return rv
}

// Returns a Boolean value that determines whether the window is the app’s
// main window.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityMain()
func (a NSApplication) IsAccessibilityMain() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isAccessibilityMain"))
	return rv
}

// Returns the Boolean value that determines whether the window is in a
// minimized state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityMinimized()
func (a NSApplication) IsAccessibilityMinimized() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isAccessibilityMinimized"))
	return rv
}

// Returns a Boolean value that determines whether the window is modal.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityModal()
func (a NSApplication) IsAccessibilityModal() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isAccessibilityModal"))
	return rv
}

// Returns a Boolean value that determines whether the accessibility
// element’s grid is in row major order or in column major order.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityOrderedByRow()
func (a NSApplication) IsAccessibilityOrderedByRow() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isAccessibilityOrderedByRow"))
	return rv
}

// Returns a Boolean value that determines whether the accessibility element
// contains protected content.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityProtectedContent()
func (a NSApplication) IsAccessibilityProtectedContent() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isAccessibilityProtectedContent"))
	return rv
}

// Returns a Boolean value that determines whether the accessibility element
// must have content for successful submission of a form.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityRequired()
func (a NSApplication) IsAccessibilityRequired() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isAccessibilityRequired"))
	return rv
}

// Returns a Boolean value that determines whether the accessibility element
// is currently in a selected state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilitySelected()
func (a NSApplication) IsAccessibilitySelected() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isAccessibilitySelected"))
	return rv
}

// Returns a Boolean value that indicates whether assistive apps can invoke
// the specified selector on the accessibility element.
//
// selector: The selector to check.
//
// # Return Value
//
// true, if accessibility clients can call the selector; otherwise, false.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilitySelectorAllowed(_:)
func (a NSApplication) IsAccessibilitySelectorAllowed(selector objc.SEL) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isAccessibilitySelectorAllowed:"), selector)
	return rv
}

// Miniaturizes all the receiver’s windows.
//
// sender: The object that sent the command.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/miniaturizeAll(_:)
func (a NSApplication) MiniaturizeAll(sender objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("miniaturizeAll:"), sender)
}

// Opens the character palette.
//
// sender: The object that sent the command.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/orderFrontCharacterPalette(_:)
func (a NSApplication) OrderFrontCharacterPalette(sender objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("orderFrontCharacterPalette:"), sender)
}

// Brings up the color panel, an instance of [NSColorPanel].
//
// sender: The object that sent the command.
//
// # Discussion
//
// If the [NSColorPanel] object does not exist yet, this method creates one.
// This method is typically invoked when the user chooses Colors from a menu.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/orderFrontColorPanel(_:)
func (a NSApplication) OrderFrontColorPanel(sender objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("orderFrontColorPanel:"), sender)
}

// Displays a standard About window.
//
// sender: The object that sent the command.
//
// # Discussion
//
// This method calls [NSApplication.OrderFrontStandardAboutPanelWithOptions]
// with a `nil` argument. See
// [NSApplication.OrderFrontStandardAboutPanelWithOptions] for a description
// of what’s displayed.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/orderFrontStandardAboutPanel(_:)
func (a NSApplication) OrderFrontStandardAboutPanel(sender objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("orderFrontStandardAboutPanel:"), sender)
}

// Displays a standard About window with information from a given options
// dictionary.
//
// optionsDictionary: A dictionary whose keys define the contents of the About window. For a list
// of keys, see [NSAboutPanelOptionKey].
//
// # Discussion
//
// In addition to the keys in AboutPanelOptionKey, you may also include the
// following key in `optionsDictionary`:
//
// - `@“"Copyright"`: An [NSString] object with a line of copyright
// information. If not specified, this method then looks for the value of
// [NSHumanReadableCopyright] in the localized version of the app’s
// `Info.Plist()` file. If neither is available, this method leaves the space
// blank.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/orderFrontStandardAboutPanel(options:)
func (a NSApplication) OrderFrontStandardAboutPanelWithOptions(optionsDictionary foundation.INSDictionary) {
	objc.Send[objc.ID](a.ID, objc.Sel("orderFrontStandardAboutPanelWithOptions:"), optionsDictionary)
}

// Suppresses the usual window ordering in handling the most recent mouse-down
// event.
//
// # Discussion
//
// This method is only useful for mouse-down events when you want to prevent
// the window that receives the event from being ordered to the front.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/preventWindowOrdering()
func (a NSApplication) PreventWindowOrdering() {
	objc.Send[objc.ID](a.ID, objc.Sel("preventWindowOrdering"))
}

// Registers the pasteboard types the receiver can send and receive in
// response to service requests.
//
// sendTypes: An array of [NSString] objects, each of which corresponds to a particular
// pasteboard type that the app can send.
//
// returnTypes: An array of [NSString] objects, each of which corresponds to a particular
// pasteboard type that the app can receive.
//
// # Discussion
//
// If the receiver has a Services menu, a menu item is added for each service
// provider that can accept one of the specified `sendTypes` or return one of
// the specified `returnTypes`. You should typically invoke this method at app
// startup time or when an object that can use services is created. You can
// invoke it more than once—its purpose is to ensure there is a menu item
// for every service the app can use. The event-handling mechanism will
// dynamically enable the individual items to indicate which services are
// currently appropriate. All the [NSResponder] objects in your app (typically
// [NSView] objects) should register every possible type they can send and
// receive by sending this message to [NSApp].
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/registerServicesMenuSendTypes(_:returnTypes:)
func (a NSApplication) RegisterServicesMenuSendTypesReturnTypes(sendTypes []string, returnTypes []string) {
	objc.Send[objc.ID](a.ID, objc.Sel("registerServicesMenuSendTypes:returnTypes:"), objectivec.StringSliceToNSArray(sendTypes), objectivec.StringSliceToNSArray(returnTypes))
}

// Removes the Window menu item for a given window.
//
// win: The window whose menu item is to be removed.
//
// # Discussion
//
// This method doesn’t prevent the item from being automatically added
// again. Use the [NSWindow.ExcludedFromWindowsMenu] method of [NSWindow] if
// you want the item to remain excluded from the Window menu.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/removeWindowsItem(_:)
func (a NSApplication) RemoveWindowsItem(win INSWindow) {
	objc.Send[objc.ID](a.ID, objc.Sel("removeWindowsItem:"), win)
}

// Invoked to request that a window be restored.
//
// identifier: The unique interface item identifier string that was previously associated
// with the window. Use this string to determine which window to create.
//
// state: A coder object containing the window state information. This coder object
// contains the combined restorable state of the window, which can include the
// state of the window, its delegate, window controller, and document object.
// You can use this state to determine which window to create.
//
// completionHandler: A Block object to execute with the results of creating the window. You must
// execute this block at some point but may do so after the method returns if
// needed. This block takes the following parameters:
//
// - The window that was created or nil if the window could not be created. -
// An error object if the window was not recognized or could not be created
// for whatever reason; otherwise, specify `nil`. In OS X v10.7, the error
// parameter is ignored.
//
// # Return Value
//
// true if the window was restored; otherwise false.
//
// # Discussion
//
// If the receiver knows how to restore the identified window, it should
// invoke the completion handler with the window, possibly creating it. It is
// acceptable to use a pre-existing window, though you should not pass the
// same window to more than one completion handler. If the receiver cannot
// restore the identified window (for example, the window referenced a
// document that has been deleted), it should invoke the completion handler
// with a nil window.
//
// The receiver is app is passed the identifier of the window, which allows it
// to quickly check for known windows. For example, you might give your
// preferences window an identifier of “preferences” in the nib, and then
// check for that identifier in your implementation. The receiver is also
// passed the [NSCoder] instance containing the combined restorable state of
// the window, its delegate, the window controller, and any document. The
// receiver may decode information previously stored in the coder to determine
// what window to restore.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/restoreWindow(withIdentifier:state:completionHandler:)
func (a NSApplication) RestoreWindowWithIdentifierStateCompletionHandler(identifier NSUserInterfaceItemIdentifier, state foundation.INSCoder, completionHandler WindowErrorHandler) bool {
	_block2, _ := NewWindowErrorBlock(completionHandler)
	rv := objc.Send[bool](a.ID, objc.Sel("restoreWindowWithIdentifier:state:completionHandler:"), identifier, state, _block2)
	return rv
}

// Starts a modal event loop for the specified window.
//
// window: The window to be displayed modally. If it is not already visible, the
// window is centered on the screen using the value in its [NSWindow.Center]
// method and made visible and key. If it is already visible, it is simply
// made key.
//
// # Return Value
//
// An integer indicating the reason that this method returned. See
// [NSModalResponse] possible return values.
//
// # Discussion
//
// This method runs a modal event loop for the specified window synchronously.
// It displays the specified window, makes it key, starts the run loop, and
// processes events for that window. (You do not need to show the window
// yourself.) While the app is in that loop, it does not respond to any other
// events (including mouse, keyboard, or window-close events) unless they are
// associated with the window. It also does not perform any tasks (such as
// firing timers) that are not associated with the modal run loop. In other
// words, this method consumes only enough CPU time to process events and
// dispatch them to the action methods associated with the modal window.
//
// You can exit the modal loop by calling the [NSApplication.StopModal],
// [NSApplication.StopModalWithCode], or [NSApplication.AbortModal] methods
// from your modal window code. If you use the
// [NSApplication.StopModalWithCode] method to stop the modal event loop, this
// method returns the argument passed to [NSApplication.StopModalWithCode]. If
// you use [NSApplication.StopModal] instead, this method returns the constant
// [stop]. If you use [NSApplication.AbortModal], this method returns the
// constant [abort].
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/runModal(for:)
//
// [abort]: https://developer.apple.com/documentation/AppKit/NSApplication/ModalResponse/abort
// [stop]: https://developer.apple.com/documentation/AppKit/NSApplication/ModalResponse/stop
func (a NSApplication) RunModalForWindow(window INSWindow) NSModalResponse {
	rv := objc.Send[NSModalResponse](a.ID, objc.Sel("runModalForWindow:"), window)
	return NSModalResponse(rv)
}

// Runs a given modal session, as defined in a previous invocation of
// [NSApplication.BeginModalSessionForWindow].
//
// session: A pointer to the modal session structure returned by the
// [NSApplication.BeginModalSessionForWindow] method for the window to be
// displayed.
//
// # Return Value
//
// An integer indicating the reason that this method returned. See the
// discussion for a description of possible return values.
//
// # Discussion
//
// A loop that uses this method is similar in some ways to a modal event loop
// run with [NSApplication.RunModalForWindow], except with this method your
// code can do some additional work between method invocations. When you
// invoke this method, events for the [NSWindow] object of this session are
// dispatched as normal. This method returns when there are no more events.
// You must invoke this method frequently enough in your loop that the window
// remains responsive to events. However, you should not invoke this method in
// a tight loop because it returns immediately if there are no events, and
// consequently you could end up polling for events rather than blocking.
//
// Typically, you use this method in situations where you want to do some
// additional processing on the current thread while the modal loop runs. For
// example, while processing a large data set, you might want to use a modal
// dialog to display progress and give the user a chance to cancel the
// operation. If you want to display a modal dialog and do not need to do any
// additional work in parallel, use [NSApplication.RunModalForWindow] instead.
// When there are no pending events, that method waits idly instead of
// consuming CPU time.
//
// The following code shows a sample loop you can use in your code:
//
// If the modal session was not stopped, this method returns
// [NSModalResponseContinue]. At this point, your app can do some work before
// the next invocation of [NSApplication.RunModalSession] (as indicated in the
// example’s `doSomeWork` call). If [NSApplication.StopModal] was invoked as
// the result of event processing, [NSApplication.RunModalSession] returns
// [NSModalResponseStop]. If [NSApplication.StopModalWithCode] was invoked,
// this method returns the value passed to [NSApplication.StopModalWithCode].
// If [NSApplication.AbortModal] was invoked, this method returns
// [NSModalResponseAbort].
//
// The window is placed on the screen and made key as a result of the
// [NSApplication.RunModalSession] message. Do not send a separate
// [NSWindow.KeyAndOrderFront] message.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/runModalSession(_:)
func (a NSApplication) RunModalSession(session NSModalSession) NSModalResponse {
	rv := objc.Send[NSModalResponse](a.ID, objc.Sel("runModalSession:"), session)
	return NSModalResponse(rv)
}

// Displays the receiver’s page layout panel, an instance of [NSPageLayout].
//
// sender: The object that sent the command.
//
// # Discussion
//
// If the [NSPageLayout] instance does not exist, this method creates one.
// This method is typically invoked when the user chooses Page Setup from the
// app’s File menu.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/runPageLayout(_:)
func (a NSApplication) RunPageLayout(sender objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("runPageLayout:"), sender)
}

// Sets the activation point for the user interface element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityActivationPoint(_:)
func (a NSApplication) SetAccessibilityActivationPoint(accessibilityActivationPoint corefoundation.CGPoint) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityActivationPoint:"), accessibilityActivationPoint)
}

// Sets the allowed values for the slider accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityAllowedValues(_:)
func (a NSApplication) SetAccessibilityAllowedValues(accessibilityAllowedValues []foundation.NSNumber) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityAllowedValues:"), objectivec.IObjectSliceToNSArray(accessibilityAllowedValues))
}

// Sets the Boolean value that determines whether the accessibility
// element’s alternative UI is currently visible.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityAlternateUIVisible(_:)
func (a NSApplication) SetAccessibilityAlternateUIVisible(accessibilityAlternateUIVisible bool) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityAlternateUIVisible:"), accessibilityAlternateUIVisible)
}

// Sets the child accessibility element with the current focus.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityApplicationFocusedUIElement(_:)
func (a NSApplication) SetAccessibilityApplicationFocusedUIElement(accessibilityApplicationFocusedUIElement objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityApplicationFocusedUIElement:"), accessibilityApplicationFocusedUIElement)
}

// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityAttributedUserInputLabels(_:)
func (a NSApplication) SetAccessibilityAttributedUserInputLabels(accessibilityAttributedUserInputLabels []foundation.NSAttributedString) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityAttributedUserInputLabels:"), objectivec.IObjectSliceToNSArray(accessibilityAttributedUserInputLabels))
}

// Sets the child accessibility element that represents the window’s cancel
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityCancelButton(_:)
func (a NSApplication) SetAccessibilityCancelButton(accessibilityCancelButton objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityCancelButton:"), accessibilityCancelButton)
}

// Sets the child accessibility elements in the accessibility hierarchy.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityChildren(_:)
func (a NSApplication) SetAccessibilityChildren(accessibilityChildren foundation.INSArray) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityChildren:"), accessibilityChildren)
}

// Sets the array of child accessibility elements in order for linear
// navigation.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityChildrenInNavigationOrder(_:)
func (a NSApplication) SetAccessibilityChildrenInNavigationOrder(accessibilityChildrenInNavigationOrder []objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityChildrenInNavigationOrder:"), objectivec.IObjectSliceToNSArray(accessibilityChildrenInNavigationOrder))
}

// Sets the clear button for the search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityClearButton(_:)
func (a NSApplication) SetAccessibilityClearButton(accessibilityClearButton objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityClearButton:"), accessibilityClearButton)
}

// Sets the child accessibility element that represents the window’s close
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityCloseButton(_:)
func (a NSApplication) SetAccessibilityCloseButton(accessibilityCloseButton objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityCloseButton:"), accessibilityCloseButton)
}

// Sets the number of columns in the accessibility element’s grid.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityColumnCount(_:)
func (a NSApplication) SetAccessibilityColumnCount(accessibilityColumnCount int) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityColumnCount:"), accessibilityColumnCount)
}

// Sets the column header accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityColumnHeaderUIElements(_:)
func (a NSApplication) SetAccessibilityColumnHeaderUIElements(accessibilityColumnHeaderUIElements foundation.INSArray) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityColumnHeaderUIElements:"), accessibilityColumnHeaderUIElements)
}

// Sets the column index range of the cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityColumnIndexRange(_:)
func (a NSApplication) SetAccessibilityColumnIndexRange(accessibilityColumnIndexRange foundation.NSRange) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityColumnIndexRange:"), accessibilityColumnIndexRange)
}

// Sets the column titles for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityColumnTitles(_:)
func (a NSApplication) SetAccessibilityColumnTitles(accessibilityColumnTitles foundation.INSArray) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityColumnTitles:"), accessibilityColumnTitles)
}

// Sets the column accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityColumns(_:)
func (a NSApplication) SetAccessibilityColumns(accessibilityColumns foundation.INSArray) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityColumns:"), accessibilityColumns)
}

// Sets the contents of the current accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityContents(_:)
func (a NSApplication) SetAccessibilityContents(accessibilityContents foundation.INSArray) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityContents:"), accessibilityContents)
}

// Sets the critical value for the level indicator.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityCriticalValue(_:)
func (a NSApplication) SetAccessibilityCriticalValue(accessibilityCriticalValue objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityCriticalValue:"), accessibilityCriticalValue)
}

// Sets the custom actions of the current accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityCustomActions(_:)
func (a NSApplication) SetAccessibilityCustomActions(accessibilityCustomActions []NSAccessibilityCustomAction) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityCustomActions:"), objectivec.IObjectSliceToNSArray(accessibilityCustomActions))
}

// Sets the custom rotors of the current accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityCustomRotors(_:)
func (a NSApplication) SetAccessibilityCustomRotors(accessibilityCustomRotors []NSAccessibilityCustomRotor) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityCustomRotors:"), objectivec.IObjectSliceToNSArray(accessibilityCustomRotors))
}

// Sets the decrement button for the stepper accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityDecrementButton(_:)
func (a NSApplication) SetAccessibilityDecrementButton(accessibilityDecrementButton objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityDecrementButton:"), accessibilityDecrementButton)
}

// Sets the child accessibility element that represents the window’s default
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityDefaultButton(_:)
func (a NSApplication) SetAccessibilityDefaultButton(accessibilityDefaultButton objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityDefaultButton:"), accessibilityDefaultButton)
}

// Sets a Boolean value that determines whether the row is disclosing other
// rows.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityDisclosed(_:)
func (a NSApplication) SetAccessibilityDisclosed(accessibilityDisclosed bool) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityDisclosed:"), accessibilityDisclosed)
}

// Sets the row disclosing the current row.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityDisclosedByRow(_:)
func (a NSApplication) SetAccessibilityDisclosedByRow(accessibilityDisclosedByRow objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityDisclosedByRow:"), accessibilityDisclosedByRow)
}

// Sets the rows that the current row discloses.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityDisclosedRows(_:)
func (a NSApplication) SetAccessibilityDisclosedRows(accessibilityDisclosedRows objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityDisclosedRows:"), accessibilityDisclosedRows)
}

// Sets the indention level for the row.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityDisclosureLevel(_:)
func (a NSApplication) SetAccessibilityDisclosureLevel(accessibilityDisclosureLevel int) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityDisclosureLevel:"), accessibilityDisclosureLevel)
}

// Sets the URL for the file that the accessibility element represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityDocument(_:)
func (a NSApplication) SetAccessibilityDocument(accessibilityDocument string) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityDocument:"), objc.String(accessibilityDocument))
}

// Sets a Boolean value that indicates whether the accessibility element is in
// an edited state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityEdited(_:)
func (a NSApplication) SetAccessibilityEdited(accessibilityEdited bool) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityEdited:"), accessibilityEdited)
}

// Sets a Boolean value that determines whether the accessibility element
// participates in the accessibility hierarchy.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityElement(_:)
func (a NSApplication) SetAccessibilityElement(accessibilityElement bool) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityElement:"), accessibilityElement)
}

// Sets a Boolean value that determines whether the accessibility element
// responds to user events.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityEnabled(_:)
func (a NSApplication) SetAccessibilityEnabled(accessibilityEnabled bool) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityEnabled:"), accessibilityEnabled)
}

// Sets a Boolean value that determines whether accessibility element is in an
// expanded state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityExpanded(_:)
func (a NSApplication) SetAccessibilityExpanded(accessibilityExpanded bool) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityExpanded:"), accessibilityExpanded)
}

// Sets the icon for the app’s menu bar extra.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityExtrasMenuBar(_:)
func (a NSApplication) SetAccessibilityExtrasMenuBar(accessibilityExtrasMenuBar objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityExtrasMenuBar:"), accessibilityExtrasMenuBar)
}

// Sets the filename for the file that the accessibility element represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityFilename(_:)
func (a NSApplication) SetAccessibilityFilename(accessibilityFilename string) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityFilename:"), objc.String(accessibilityFilename))
}

// Sets a Boolean value that determines whether the accessibility element has
// the keyboard focus.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityFocused(_:)
func (a NSApplication) SetAccessibilityFocused(accessibilityFocused bool) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityFocused:"), accessibilityFocused)
}

// Sets the child window with the current focus.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityFocusedWindow(_:)
func (a NSApplication) SetAccessibilityFocusedWindow(accessibilityFocusedWindow objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityFocusedWindow:"), accessibilityFocusedWindow)
}

// Sets the accessibility element’s frame in screen coordinates.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityFrame(_:)
func (a NSApplication) SetAccessibilityFrame(accessibilityFrame corefoundation.CGRect) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityFrame:"), accessibilityFrame)
}

// Sets a Boolean value that determines whether the app is the frontmost app.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityFrontmost(_:)
func (a NSApplication) SetAccessibilityFrontmost(accessibilityFrontmost bool) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityFrontmost:"), accessibilityFrontmost)
}

// Sets the child accessibility element that represents the window’s
// full-screen button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityFullScreenButton(_:)
func (a NSApplication) SetAccessibilityFullScreenButton(accessibilityFullScreenButton objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityFullScreenButton:"), accessibilityFullScreenButton)
}

// Sets the child accessibility element that represents the window’s grow
// area.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityGrowArea(_:)
func (a NSApplication) SetAccessibilityGrowArea(accessibilityGrowArea objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityGrowArea:"), accessibilityGrowArea)
}

// Sets the drag handle accessibility elements for the layout item element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityHandles(_:)
func (a NSApplication) SetAccessibilityHandles(accessibilityHandles foundation.INSArray) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityHandles:"), accessibilityHandles)
}

// Sets the header for the table view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityHeader(_:)
func (a NSApplication) SetAccessibilityHeader(accessibilityHeader objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityHeader:"), accessibilityHeader)
}

// Sets the help text for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityHelp(_:)
func (a NSApplication) SetAccessibilityHelp(accessibilityHelp string) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityHelp:"), objc.String(accessibilityHelp))
}

// Sets a Boolean value that determines whether the app is in a hidden state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityHidden(_:)
func (a NSApplication) SetAccessibilityHidden(accessibilityHidden bool) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityHidden:"), accessibilityHidden)
}

// Sets the horizontal scroll bar for the scroll view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityHorizontalScrollBar(_:)
func (a NSApplication) SetAccessibilityHorizontalScrollBar(accessibilityHorizontalScrollBar objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityHorizontalScrollBar:"), accessibilityHorizontalScrollBar)
}

// Sets the description of the layout area’s horizontal units.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityHorizontalUnitDescription(_:)
func (a NSApplication) SetAccessibilityHorizontalUnitDescription(accessibilityHorizontalUnitDescription string) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityHorizontalUnitDescription:"), objc.String(accessibilityHorizontalUnitDescription))
}

// Sets the units that the layout area uses for horizontal values.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityHorizontalUnits(_:)
func (a NSApplication) SetAccessibilityHorizontalUnits(accessibilityHorizontalUnits NSAccessibilityUnits) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityHorizontalUnits:"), accessibilityHorizontalUnits)
}

// Sets the accessibility element’s identity.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityIdentifier(_:)
func (a NSApplication) SetAccessibilityIdentifier(accessibilityIdentifier string) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityIdentifier:"), objc.String(accessibilityIdentifier))
}

// Sets the increment button for the stepper accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityIncrementButton(_:)
func (a NSApplication) SetAccessibilityIncrementButton(accessibilityIncrementButton objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityIncrementButton:"), accessibilityIncrementButton)
}

// Sets the index of the row or column that the accessibility element
// represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityIndex(_:)
func (a NSApplication) SetAccessibilityIndex(accessibilityIndex int) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityIndex:"), accessibilityIndex)
}

// Sets the line number that contains the insertion point.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityInsertionPointLineNumber(_:)
func (a NSApplication) SetAccessibilityInsertionPointLineNumber(accessibilityInsertionPointLineNumber int) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityInsertionPointLineNumber:"), accessibilityInsertionPointLineNumber)
}

// Sets a short description of the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityLabel(_:)
func (a NSApplication) SetAccessibilityLabel(accessibilityLabel string) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityLabel:"), objc.String(accessibilityLabel))
}

// Sets the child label elements for the slider accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityLabelUIElements(_:)
func (a NSApplication) SetAccessibilityLabelUIElements(accessibilityLabelUIElements foundation.INSArray) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityLabelUIElements:"), accessibilityLabelUIElements)
}

// Sets the value of the label accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityLabelValue(_:)
func (a NSApplication) SetAccessibilityLabelValue(accessibilityLabelValue float32) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityLabelValue:"), accessibilityLabelValue)
}

// Sets the elements that have links with the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityLinkedUIElements(_:)
func (a NSApplication) SetAccessibilityLinkedUIElements(accessibilityLinkedUIElements foundation.INSArray) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityLinkedUIElements:"), accessibilityLinkedUIElements)
}

// Sets a Boolean value that determines whether the window is the app’s main
// window.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMain(_:)
func (a NSApplication) SetAccessibilityMain(accessibilityMain bool) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityMain:"), accessibilityMain)
}

// Sets the app’s main window.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMainWindow(_:)
func (a NSApplication) SetAccessibilityMainWindow(accessibilityMainWindow objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityMainWindow:"), accessibilityMainWindow)
}

// Sets the user interface element that functions as a marker group for the
// ruler accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMarkerGroupUIElement(_:)
func (a NSApplication) SetAccessibilityMarkerGroupUIElement(accessibilityMarkerGroupUIElement objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityMarkerGroupUIElement:"), accessibilityMarkerGroupUIElement)
}

// Sets the human-readable description of the marker type.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMarkerTypeDescription(_:)
func (a NSApplication) SetAccessibilityMarkerTypeDescription(accessibilityMarkerTypeDescription string) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityMarkerTypeDescription:"), objc.String(accessibilityMarkerTypeDescription))
}

// Sets the array of marker accessibility elements for the ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMarkerUIElements(_:)
func (a NSApplication) SetAccessibilityMarkerUIElements(accessibilityMarkerUIElements foundation.INSArray) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityMarkerUIElements:"), accessibilityMarkerUIElements)
}

// Sets the marker values for the ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMarkerValues(_:)
func (a NSApplication) SetAccessibilityMarkerValues(accessibilityMarkerValues objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityMarkerValues:"), accessibilityMarkerValues)
}

// Sets the maximum value for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMaxValue(_:)
func (a NSApplication) SetAccessibilityMaxValue(accessibilityMaxValue objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityMaxValue:"), accessibilityMaxValue)
}

// Sets the app’s menu bar.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMenuBar(_:)
func (a NSApplication) SetAccessibilityMenuBar(accessibilityMenuBar objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityMenuBar:"), accessibilityMenuBar)
}

// Sets the minimum value for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMinValue(_:)
func (a NSApplication) SetAccessibilityMinValue(accessibilityMinValue objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityMinValue:"), accessibilityMinValue)
}

// Sets the child accessibility element that represents the window’s
// minimize button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMinimizeButton(_:)
func (a NSApplication) SetAccessibilityMinimizeButton(accessibilityMinimizeButton objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityMinimizeButton:"), accessibilityMinimizeButton)
}

// Sets the Boolean value that determines whether the window is in a minimized
// state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMinimized(_:)
func (a NSApplication) SetAccessibilityMinimized(accessibilityMinimized bool) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityMinimized:"), accessibilityMinimized)
}

// Sets a Boolean value that determines whether the window is modal.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityModal(_:)
func (a NSApplication) SetAccessibilityModal(accessibilityModal bool) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityModal:"), accessibilityModal)
}

// Sets the contents that follow the divider accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityNextContents(_:)
func (a NSApplication) SetAccessibilityNextContents(accessibilityNextContents foundation.INSArray) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityNextContents:"), accessibilityNextContents)
}

// Sets the number of characters in the text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityNumberOfCharacters(_:)
func (a NSApplication) SetAccessibilityNumberOfCharacters(accessibilityNumberOfCharacters int) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityNumberOfCharacters:"), accessibilityNumberOfCharacters)
}

// Sets a Boolean value that determines whether the element’s grid is in row
// major order or in column major order.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityOrderedByRow(_:)
func (a NSApplication) SetAccessibilityOrderedByRow(accessibilityOrderedByRow bool) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityOrderedByRow:"), accessibilityOrderedByRow)
}

// Sets the orientation of the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityOrientation(_:)
func (a NSApplication) SetAccessibilityOrientation(accessibilityOrientation NSAccessibilityOrientation) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityOrientation:"), accessibilityOrientation)
}

// Sets the overflow button for the toolbar.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityOverflowButton(_:)
func (a NSApplication) SetAccessibilityOverflowButton(accessibilityOverflowButton objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityOverflowButton:"), accessibilityOverflowButton)
}

// Sets the accessibility element’s parent in the accessibility hierarchy.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityParent(_:)
func (a NSApplication) SetAccessibilityParent(accessibilityParent objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityParent:"), accessibilityParent)
}

// Sets the placeholder value for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityPlaceholderValue(_:)
func (a NSApplication) SetAccessibilityPlaceholderValue(accessibilityPlaceholderValue string) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityPlaceholderValue:"), objc.String(accessibilityPlaceholderValue))
}

// Sets the contents that precede the divider accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityPreviousContents(_:)
func (a NSApplication) SetAccessibilityPreviousContents(accessibilityPreviousContents foundation.INSArray) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityPreviousContents:"), accessibilityPreviousContents)
}

// Sets a Boolean value that determines whether the accessibility element
// contains protected content.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityProtectedContent(_:)
func (a NSApplication) SetAccessibilityProtectedContent(accessibilityProtectedContent bool) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityProtectedContent:"), accessibilityProtectedContent)
}

// Sets the child accessibility element that represents the window’s proxy
// icon.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityProxy(_:)
func (a NSApplication) SetAccessibilityProxy(accessibilityProxy objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityProxy:"), accessibilityProxy)
}

// Sets a Boolean value that determines whether the accessibility element must
// have content for successful submission of a form.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityRequired(_:)
func (a NSApplication) SetAccessibilityRequired(accessibilityRequired bool) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityRequired:"), accessibilityRequired)
}

// Sets the type of interface element that the accessibility element
// represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityRole(_:)
func (a NSApplication) SetAccessibilityRole(accessibilityRole NSAccessibilityRole) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityRole:"), objc.String(string(accessibilityRole)))
}

// Sets the localized, human-intelligible description of the accessibility
// element’s role, such as radio button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityRoleDescription(_:)
func (a NSApplication) SetAccessibilityRoleDescription(accessibilityRoleDescription string) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityRoleDescription:"), objc.String(accessibilityRoleDescription))
}

// Sets the number of rows in the accessibility element’s grid.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityRowCount(_:)
func (a NSApplication) SetAccessibilityRowCount(accessibilityRowCount int) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityRowCount:"), accessibilityRowCount)
}

// Sets the row header accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityRowHeaderUIElements(_:)
func (a NSApplication) SetAccessibilityRowHeaderUIElements(accessibilityRowHeaderUIElements foundation.INSArray) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityRowHeaderUIElements:"), accessibilityRowHeaderUIElements)
}

// Sets the row index range of the cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityRowIndexRange(_:)
func (a NSApplication) SetAccessibilityRowIndexRange(accessibilityRowIndexRange foundation.NSRange) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityRowIndexRange:"), accessibilityRowIndexRange)
}

// Sets the row accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityRows(_:)
func (a NSApplication) SetAccessibilityRows(accessibilityRows foundation.INSArray) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityRows:"), accessibilityRows)
}

// Sets the type of markers for the ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityRulerMarkerType(_:)
func (a NSApplication) SetAccessibilityRulerMarkerType(accessibilityRulerMarkerType NSAccessibilityRulerMarkerType) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityRulerMarkerType:"), accessibilityRulerMarkerType)
}

// Sets the search button for the search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySearchButton(_:)
func (a NSApplication) SetAccessibilitySearchButton(accessibilitySearchButton objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilitySearchButton:"), accessibilitySearchButton)
}

// Sets the search menu for the search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySearchMenu(_:)
func (a NSApplication) SetAccessibilitySearchMenu(accessibilitySearchMenu objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilitySearchMenu:"), accessibilitySearchMenu)
}

// Sets a Boolean value that determines whether the accessibility element is
// currently in a selected state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySelected(_:)
func (a NSApplication) SetAccessibilitySelected(accessibilitySelected bool) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilitySelected:"), accessibilitySelected)
}

// Sets the currently selected cells for the table.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySelectedCells(_:)
func (a NSApplication) SetAccessibilitySelectedCells(accessibilitySelectedCells foundation.INSArray) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilitySelectedCells:"), accessibilitySelectedCells)
}

// Sets the accessibility element’s currently selected children.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySelectedChildren(_:)
func (a NSApplication) SetAccessibilitySelectedChildren(accessibilitySelectedChildren foundation.INSArray) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilitySelectedChildren:"), accessibilitySelectedChildren)
}

// Sets the currently selected columns for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySelectedColumns(_:)
func (a NSApplication) SetAccessibilitySelectedColumns(accessibilitySelectedColumns foundation.INSArray) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilitySelectedColumns:"), accessibilitySelectedColumns)
}

// Sets the currently selected rows for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySelectedRows(_:)
func (a NSApplication) SetAccessibilitySelectedRows(accessibilitySelectedRows foundation.INSArray) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilitySelectedRows:"), accessibilitySelectedRows)
}

// Sets the currently selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySelectedText(_:)
func (a NSApplication) SetAccessibilitySelectedText(accessibilitySelectedText string) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilitySelectedText:"), objc.String(accessibilitySelectedText))
}

// Sets the range of the currently selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySelectedTextRange(_:)
func (a NSApplication) SetAccessibilitySelectedTextRange(accessibilitySelectedTextRange foundation.NSRange) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilitySelectedTextRange:"), accessibilitySelectedTextRange)
}

// Sets an array of ranges for the currently selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySelectedTextRanges(_:)
func (a NSApplication) SetAccessibilitySelectedTextRanges(accessibilitySelectedTextRanges []foundation.NSValue) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilitySelectedTextRanges:"), objectivec.IObjectSliceToNSArray(accessibilitySelectedTextRanges))
}

// Sets the list of elements that the accessibility element is a title for.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityServesAsTitleForUIElements(_:)
func (a NSApplication) SetAccessibilityServesAsTitleForUIElements(accessibilityServesAsTitleForUIElements foundation.INSArray) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityServesAsTitleForUIElements:"), accessibilityServesAsTitleForUIElements)
}

// Sets the range of characters that the accessibility element displays.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySharedCharacterRange(_:)
func (a NSApplication) SetAccessibilitySharedCharacterRange(accessibilitySharedCharacterRange foundation.NSRange) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilitySharedCharacterRange:"), accessibilitySharedCharacterRange)
}

// Sets the array of elements that shares the keyboard focus with the
// accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySharedFocusElements(_:)
func (a NSApplication) SetAccessibilitySharedFocusElements(accessibilitySharedFocusElements foundation.INSArray) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilitySharedFocusElements:"), accessibilitySharedFocusElements)
}

// Sets the other elements that share text with the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySharedTextUIElements(_:)
func (a NSApplication) SetAccessibilitySharedTextUIElements(accessibilitySharedTextUIElements foundation.INSArray) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilitySharedTextUIElements:"), accessibilitySharedTextUIElements)
}

// Sets the menu currently displaying for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityShownMenu(_:)
func (a NSApplication) SetAccessibilityShownMenu(accessibilityShownMenu objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityShownMenu:"), accessibilityShownMenu)
}

// Sets the accessibility element’s sort direction.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySortDirection(_:)
func (a NSApplication) SetAccessibilitySortDirection(accessibilitySortDirection NSAccessibilitySortDirection) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilitySortDirection:"), accessibilitySortDirection)
}

// Sets the array that contains the views and splitter bar from the split
// view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySplitters(_:)
func (a NSApplication) SetAccessibilitySplitters(accessibilitySplitters foundation.INSArray) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilitySplitters:"), accessibilitySplitters)
}

// Sets the specialized interface element type that the accessibility element
// represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySubrole(_:)
func (a NSApplication) SetAccessibilitySubrole(accessibilitySubrole NSAccessibilitySubrole) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilitySubrole:"), objc.String(string(accessibilitySubrole)))
}

// Sets the tab accessibility elements for the tab view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityTabs(_:)
func (a NSApplication) SetAccessibilityTabs(accessibilityTabs foundation.INSArray) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityTabs:"), accessibilityTabs)
}

// Sets the title of the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityTitle(_:)
func (a NSApplication) SetAccessibilityTitle(accessibilityTitle string) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityTitle:"), objc.String(accessibilityTitle))
}

// Sets the static text element that represents the accessibility element’s
// title.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityTitleUIElement(_:)
func (a NSApplication) SetAccessibilityTitleUIElement(accessibilityTitleUIElement objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityTitleUIElement:"), accessibilityTitleUIElement)
}

// Sets the child accessibility element that represents the window’s toolbar
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityToolbarButton(_:)
func (a NSApplication) SetAccessibilityToolbarButton(accessibilityToolbarButton objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityToolbarButton:"), accessibilityToolbarButton)
}

// Sets the top-level element that contains the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityTopLevelUIElement(_:)
func (a NSApplication) SetAccessibilityTopLevelUIElement(accessibilityTopLevelUIElement objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityTopLevelUIElement:"), accessibilityTopLevelUIElement)
}

// Sets the URL for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityURL(_:)
func (a NSApplication) SetAccessibilityURL(accessibilityURL foundation.NSURL) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityURL:"), accessibilityURL)
}

// Sets the human-readable description of the ruler’s units.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityUnitDescription(_:)
func (a NSApplication) SetAccessibilityUnitDescription(accessibilityUnitDescription string) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityUnitDescription:"), objc.String(accessibilityUnitDescription))
}

// Sets the units used for the ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityUnits(_:)
func (a NSApplication) SetAccessibilityUnits(accessibilityUnits NSAccessibilityUnits) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityUnits:"), accessibilityUnits)
}

// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityUserInputLabels(_:)
func (a NSApplication) SetAccessibilityUserInputLabels(accessibilityUserInputLabels []string) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityUserInputLabels:"), objectivec.StringSliceToNSArray(accessibilityUserInputLabels))
}

// Sets the accessibility element’s value.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityValue(_:)
func (a NSApplication) SetAccessibilityValue(accessibilityValue objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityValue:"), accessibilityValue)
}

// Sets the human-readable description of the accessibility element’s value.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityValueDescription(_:)
func (a NSApplication) SetAccessibilityValueDescription(accessibilityValueDescription string) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityValueDescription:"), objc.String(accessibilityValueDescription))
}

// Sets the vertical scroll bar for the scroll view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityVerticalScrollBar(_:)
func (a NSApplication) SetAccessibilityVerticalScrollBar(accessibilityVerticalScrollBar objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityVerticalScrollBar:"), accessibilityVerticalScrollBar)
}

// Sets the description of the layout area’s vertical units.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityVerticalUnitDescription(_:)
func (a NSApplication) SetAccessibilityVerticalUnitDescription(accessibilityVerticalUnitDescription string) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityVerticalUnitDescription:"), objc.String(accessibilityVerticalUnitDescription))
}

// Sets the units that the layout area uses for vertical values.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityVerticalUnits(_:)
func (a NSApplication) SetAccessibilityVerticalUnits(accessibilityVerticalUnits NSAccessibilityUnits) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityVerticalUnits:"), accessibilityVerticalUnits)
}

// Sets the visible cells for the table.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityVisibleCells(_:)
func (a NSApplication) SetAccessibilityVisibleCells(accessibilityVisibleCells foundation.INSArray) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityVisibleCells:"), accessibilityVisibleCells)
}

// Sets the range of visible characters in the document.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityVisibleCharacterRange(_:)
func (a NSApplication) SetAccessibilityVisibleCharacterRange(accessibilityVisibleCharacterRange foundation.NSRange) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityVisibleCharacterRange:"), accessibilityVisibleCharacterRange)
}

// Sets the accessibility element’s visible child accessibility elements.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityVisibleChildren(_:)
func (a NSApplication) SetAccessibilityVisibleChildren(accessibilityVisibleChildren foundation.INSArray) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityVisibleChildren:"), accessibilityVisibleChildren)
}

// Sets the visible columns for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityVisibleColumns(_:)
func (a NSApplication) SetAccessibilityVisibleColumns(accessibilityVisibleColumns foundation.INSArray) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityVisibleColumns:"), accessibilityVisibleColumns)
}

// Sets the visible rows for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityVisibleRows(_:)
func (a NSApplication) SetAccessibilityVisibleRows(accessibilityVisibleRows foundation.INSArray) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityVisibleRows:"), accessibilityVisibleRows)
}

// Sets the warning value for the level indicator.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityWarningValue(_:)
func (a NSApplication) SetAccessibilityWarningValue(accessibilityWarningValue objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityWarningValue:"), accessibilityWarningValue)
}

// Sets the window that contains the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityWindow(_:)
func (a NSApplication) SetAccessibilityWindow(accessibilityWindow objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityWindow:"), accessibilityWindow)
}

// Sets the array that contains all the app’s windows.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityWindows(_:)
func (a NSApplication) SetAccessibilityWindows(accessibilityWindows foundation.INSArray) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityWindows:"), accessibilityWindows)
}

// Sets the child accessibility element that represents the window’s zoom
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityZoomButton(_:)
func (a NSApplication) SetAccessibilityZoomButton(accessibilityZoomButton objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("setAccessibilityZoomButton:"), accessibilityZoomButton)
}

// Sets whether the receiver’s windows need updating when the receiver has
// finished processing the current event.
//
// needUpdate: If true, the receiver’s windows are updated after an event is processed.
//
// # Discussion
//
// This method is especially useful for making sure menus are updated to
// reflect changes not initiated by user actions, such as messages received
// from remote objects.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/setWindowsNeedUpdate(_:)
func (a NSApplication) SetWindowsNeedUpdate(needUpdate bool) {
	objc.Send[objc.ID](a.ID, objc.Sel("setWindowsNeedUpdate:"), needUpdate)
}

// Stops a modal event loop.
//
// # Discussion
//
// This method should always be paired with a previous invocation of
// [NSApplication.RunModalForWindow] or
// [NSApplication.BeginModalSessionForWindow]. When
// [NSApplication.RunModalForWindow] is stopped with this method, it returns
// [NSModalResponseStop]. In macOS 10.9 and later, you can use this method to
// stop a [NSApplication.RunModalForWindow] loop outside of an event callback,
// such as from within a method repeatedly invoked by an [Timer] object or a
// method running in a different thread.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/stopModal()
//
// [Timer]: https://developer.apple.com/documentation/Foundation/Timer
func (a NSApplication) StopModal() {
	objc.Send[objc.ID](a.ID, objc.Sel("stopModal"))
}

// Stops a modal event loop, allowing you to return a custom result code.
//
// returnCode: The result code you want returned from the
// [NSApplication.RunModalForWindow] or [NSApplication.RunModalSession]
// method. The meaning of this result code is up to you.
//
// # Discussion
//
// This method should always be paired with a previous invocation of
// [NSApplication.RunModalForWindow] or
// [NSApplication.BeginModalSessionForWindow]. When
// [NSApplication.RunModalForWindow] is stopped with this method, it returns
// the given `returnCode`. In macOS 10.9 and later, you can use this method to
// stop a [NSApplication.RunModalForWindow] loop outside of an event callback,
// such as from within a method repeatedly invoked by an [Timer] object or a
// method running in a different thread.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/stopModal(withCode:)
//
// [Timer]: https://developer.apple.com/documentation/Foundation/Timer
func (a NSApplication) StopModalWithCode(returnCode NSModalResponse) {
	objc.Send[objc.ID](a.ID, objc.Sel("stopModalWithCode:"), returnCode)
}

// Restores hidden windows to the screen and makes the receiver active.
//
// sender: The object that sent the command.
//
// # Discussion
//
// Invokes [NSApplication.UnhideWithoutActivation].
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/unhide(_:)
func (a NSApplication) Unhide(sender objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("unhide:"), sender)
}

// Restores hidden windows without activating their owner (the receiver).
//
// # Discussion
//
// When this method begins, it posts an [willUnhideNotification] to the
// default notification center. If it completes successfully, it posts an
// [didUnhideNotification].
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/unhideWithoutActivation()
//
// [didUnhideNotification]: https://developer.apple.com/documentation/AppKit/NSApplication/didUnhideNotification
// [willUnhideNotification]: https://developer.apple.com/documentation/AppKit/NSApplication/willUnhideNotification
func (a NSApplication) UnhideWithoutActivation() {
	objc.Send[objc.ID](a.ID, objc.Sel("unhideWithoutActivation"))
}

// Sends an [NSWindow.Update] message to each onscreen window.
//
// # Discussion
//
// This method is invoked automatically in the main event loop after each
// event when running in [NSDefaultRunLoopMode] or [NSModalRunLoopMode]. This
// method is not invoked automatically when running in
// [NSEventTrackingRunLoopMode].
//
// When this method begins, it posts an [willUpdateNotification] to the
// default notification center. When it successfully completes, it posts an
// [didUpdateNotification].
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/updateWindows()
//
// [didUpdateNotification]: https://developer.apple.com/documentation/AppKit/NSApplication/didUpdateNotification
// [willUpdateNotification]: https://developer.apple.com/documentation/AppKit/NSApplication/willUpdateNotification
func (a NSApplication) UpdateWindows() {
	objc.Send[objc.ID](a.ID, objc.Sel("updateWindows"))
}

// Updates the Window menu item for a given window to reflect the edited
// status of that window.
//
// win: The window whose menu item is to be updated.
//
// # Discussion
//
// You rarely need to invoke this method because it is invoked automatically
// when the edit status of an [NSWindow] object is set.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/updateWindowsItem(_:)
func (a NSApplication) UpdateWindowsItem(win INSWindow) {
	objc.Send[objc.ID](a.ID, objc.Sel("updateWindowsItem:"), win)
}

// Implemented to override the default action of enabling or disabling a
// specific menu item.
//
// menuItem: An [NSMenuItem] object that represents the menu item.
//
// # Return Value
//
// true to enable `menuItem`, false to disable it.
//
// # Discussion
//
// The object implementing this method must be the target of `menuItem`. You
// can determine which menu item `menuItem` is by querying it for its tag or
// action.
//
// The following example disables the menu item associated with the
// `nextRecord` action method when the selected line in a table view is the
// last one; conversely, it disables the menu item with `priorRecord` as its
// action method when the selected row is the first one in the table view.
// (The `countryOrRegionKeys` array contains names that appear in the table
// view.)
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItemValidation/validateMenuItem(_:)
func (a NSApplication) ValidateMenuItem(menuItem INSMenuItem) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("validateMenuItem:"), menuItem)
	return rv
}

// Returns a Boolean value that indicates whether the sender should be
// enabled.
//
// item: The user interface item to validate. You can send `anItem` the [Action] and
// [Tag] messages.
//
// # Return Value
//
// true if the user interface item should be enabled, otherwise false.
//
// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceValidations/validateUserInterfaceItem(_:)
func (a NSApplication) ValidateUserInterfaceItem(item NSValidatedUserInterfaceItem) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("validateUserInterfaceItem:"), item)
	return rv
}

// Returns the window corresponding to the specified window number.
//
// windowNum: The unique window number associated with the desired [NSWindow] object.
//
// # Return Value
//
// The desired window object or `nil` if the window could not be found.
//
// # Discussion
//
// [NSApplication.WindowWithWindowNumber] may return `nil` for window numbers
// found using [NSWindowClass.WindowNumbersWithOptions] if there is no
// corresponding window object owned by your app—for example, the menu bar.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/window(withWindowNumber:)
func (a NSApplication) WindowWithWindowNumber(windowNum int) INSWindow {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("windowWithWindowNumber:"), windowNum)
	return NSWindowFromID(rv)
}

// Creates and executes a new thread based on the specified target and
// selector.
//
// selector: The selector whose code you want to execute in the new thread.
//
// target: The object that defines the specified selector.
//
// argument: An optional argument you want to pass to the selector.
//
// # Discussion
//
// This method is a convenience wrapper for the
// [detachNewThreadSelector(_:toTarget:with:)] method of [Thread]. This method
// automatically creates an `@autoreleasepool` block for the new thread before
// invoking `selector`.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/detachDrawingThread(_:toTarget:with:)
//
// [Thread]: https://developer.apple.com/documentation/Foundation/Thread
// [detachNewThreadSelector(_:toTarget:with:)]: https://developer.apple.com/documentation/Foundation/Thread/detachNewThreadSelector(_:toTarget:with:)
func (_NSApplicationClass NSApplicationClass) DetachDrawingThreadToTargetWithObject(selector objc.SEL, target objectivec.IObject, argument objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_NSApplicationClass.class), objc.Sel("detachDrawingThread:toTarget:withObject:"), selector, target, argument)
}

// The app delegate object.
//
// # Discussion
//
// The app object and app delegate work in tandem to manage the app’s
// overall behavior. Typically, the delegate is configured automatically by
// the Xcode project templates.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/delegate
func (a NSApplication) Delegate() NSApplicationDelegate {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("delegate"))
	return NSApplicationDelegateObjectFromID(rv)
}
func (a NSApplication) SetDelegate(value NSApplicationDelegate) {
	objc.Send[struct{}](a.ID, objc.Sel("setDelegate:"), value)
}

// The last event object that the app retrieved from the event queue.
//
// # Discussion
//
// The shared app object receives events and forwards them to the affected
// [NSWindow] objects, which then distribute them to the objects in its view
// hierarchy. Use this property to get the event that was last handled by the
// app.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/currentEvent
func (a NSApplication) CurrentEvent() INSEvent {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("currentEvent"))
	return NSEventFromID(objc.ID(rv))
}

// A Boolean value indicating whether the main event loop is running.
//
// # Discussion
//
// The value of this property is true when the main event loop is running or
// false when it’s not. Calling the [NSApplication.Stop] method sets the
// value to false.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/isRunning
func (a NSApplication) IsRunning() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isRunning"))
	return rv
}

// A Boolean value indicating whether this is the active app.
//
// # Discussion
//
// The value of this property is true if the app is active or false if it’s
// not.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/isActive
func (a NSApplication) IsActive() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isActive"))
	return rv
}

// The types of push notifications that the app accepts.
//
// # Return Value
//
// A bit mask whose values indicate the types of notifications the user has
// requested for the app. See [NSApplication.RemoteNotificationType] for valid
// bit-mask values.
//
// # Discussion
//
// This property contains a bitmask whose values indicate the types of push
// notifications that the app requested. You don’t set this property
// directly. Call the [NSApplication.RegisterForRemoteNotificationTypes]
// method to register your app with Apple Push Notification Service and
// request the notification types your app supports. macOS delivers only
// notifications of types that the app supports. For a list of possible
// values, see [NSApplication.RemoteNotificationType].
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/enabledRemoteNotificationTypes
//
// [NSApplication.RemoteNotificationType]: https://developer.apple.com/documentation/AppKit/NSApplication/RemoteNotificationType
func (a NSApplication) EnabledRemoteNotificationTypes() NSRemoteNotificationType {
	rv := objc.Send[NSRemoteNotificationType](a.ID, objc.Sel("enabledRemoteNotificationTypes"))
	return NSRemoteNotificationType(rv)
}

// A Boolean value indicating whether the app is registered with Apple Push
// Notification service (APNs).
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/isRegisteredForRemoteNotifications
func (a NSApplication) IsRegisteredForRemoteNotifications() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isRegisteredForRemoteNotifications"))
	return rv
}

// The appearance associated with the app’s windows.
//
// # Discussion
//
// When the value of this property is `nil` (the default), AppKit applies the
// current system appearance to the app’s user interface elements, including
// its windows, views, panels, and popovers. Assigning an [NSAppearance]
// object to this property causes the app’s interface elements to adopt the
// specified appearance instead.
//
// Individual windows and views may still override the app’s appearance to
// customize their own appearance.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/appearance
func (a NSApplication) Appearance() INSAppearance {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("appearance"))
	return NSAppearanceFromID(objc.ID(rv))
}
func (a NSApplication) SetAppearance(value INSAppearance) {
	objc.Send[struct{}](a.ID, objc.Sel("setAppearance:"), value)
}

// The appearance that AppKit uses to draw the app’s interface.
//
// # Discussion
//
// This property always contains an [NSAppearance] object representing the
// appearance to use during drawing. If you don’t explicitly assign a value
// to the [NSApplication.Appearance] property, the app inherits the system’s
// effective appearance.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/effectiveAppearance
func (a NSApplication) EffectiveAppearance() INSAppearance {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("effectiveAppearance"))
	return NSAppearanceFromID(objc.ID(rv))
}

// The set of app presentation options that are currently in effect for the
// system.
//
// # Return Value
//
// The presentation options. The constants are listed in
// [NSApplication.PresentationOptions] and can combined using a C bitwise OR
// operator.
//
// # Discussion
//
// This property contains the presentation options that have been put into
// effect by the currently active app. You can use key-value observing on this
// property to receive notifications when:
//
// - The client is the active app and makes a change itself using either the
// [NSApplication.PresentationOptions] property or the [SetSystemUIMode]
// function. - Another app is active and makes presentation changes of its
// own. - Another app becomes active and causes the active set of presentation
// options to change.
//
// Key-value observing notifications aren’t sent when one of the above
// conditions occur, but has the same set of presentation options as the
// previously active app.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/currentSystemPresentationOptions
//
// [NSApplication.PresentationOptions]: https://developer.apple.com/documentation/AppKit/NSApplication/PresentationOptions-swift.struct
func (a NSApplication) CurrentSystemPresentationOptions() NSApplicationPresentationOptions {
	rv := objc.Send[NSApplicationPresentationOptions](a.ID, objc.Sel("currentSystemPresentationOptions"))
	return NSApplicationPresentationOptions(rv)
}

// The presentation options that should be in effect for the system when this
// app is active.
//
// # Discussion
//
// This value contains a bitwise OR of the constants listed in
// [NSApplication.PresentationOptions]. Trying to set the property to an
// invalid combination of option flags raises an [invalidArgumentException]
// exception. See the constants for a description of the valid combinations.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/presentationOptions-swift.property
//
// [NSApplication.PresentationOptions]: https://developer.apple.com/documentation/AppKit/NSApplication/PresentationOptions-swift.struct
// [invalidArgumentException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/invalidArgumentException
func (a NSApplication) PresentationOptions() NSApplicationPresentationOptions {
	rv := objc.Send[NSApplicationPresentationOptions](a.ID, objc.Sel("presentationOptions"))
	return NSApplicationPresentationOptions(rv)
}
func (a NSApplication) SetPresentationOptions(value NSApplicationPresentationOptions) {
	objc.Send[struct{}](a.ID, objc.Sel("setPresentationOptions:"), value)
}

// The layout direction of the user interface.
//
// # Discussion
//
// This property contains the general user interface layout flow directions.
// For a list of possible values, see [NSUserInterfaceLayoutDirection].
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/userInterfaceLayoutDirection
//
// [NSUserInterfaceLayoutDirection]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceLayoutDirection
func (a NSApplication) UserInterfaceLayoutDirection() NSUserInterfaceLayoutDirection {
	rv := objc.Send[NSUserInterfaceLayoutDirection](a.ID, objc.Sel("userInterfaceLayoutDirection"))
	return NSUserInterfaceLayoutDirection(rv)
}

// The app’s Dock tile.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/dockTile
func (a NSApplication) DockTile() INSDockTile {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("dockTile"))
	return NSDockTileFromID(objc.ID(rv))
}

// The image used for the app’s icon.
//
// # Discussion
//
// Assign an image to this property when you want to temporarily change the
// app icon in the dock app tile. The image you provide is scaled as needed so
// that it fits in the tile. To restore your app’s original icon, set this
// property to `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/applicationIconImage
func (a NSApplication) ApplicationIconImage() INSImage {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("applicationIconImage"))
	return NSImageFromID(objc.ID(rv))
}
func (a NSApplication) SetApplicationIconImage(value INSImage) {
	objc.Send[struct{}](a.ID, objc.Sel("setApplicationIconImage:"), value)
}

// The help menu used by the app.
//
// # Discussion
//
// Use this property to specify your app’s Help menu. When this property
// contains a valid menu, the system installs its Spotlight-related menu items
// on that menu. When the value is `nil`, AppKit installs Spotlight menu items
// on the menu of its choosing. To suppress Spotlight help items altogether,
// specify a menu that doesn’t appear on the menu bar.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/helpMenu
func (a NSApplication) HelpMenu() INSMenu {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("helpMenu"))
	return NSMenuFromID(objc.ID(rv))
}
func (a NSApplication) SetHelpMenu(value INSMenu) {
	objc.Send[struct{}](a.ID, objc.Sel("setHelpMenu:"), value)
}

// The object that provides the services the current app advertises in the
// Services menu of other apps.
//
// # Return Value
//
// The app’s service provider object.
//
// # Discussion
//
// The service provider performs all advertised services for the app. When
// another app requests a service from the current app, the app object
// forwards the request to its service provider. Service requests can arrive
// immediately after the service provider is set, so assign an object to this
// property only when your app is ready to receive requests.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/servicesProvider
func (a NSApplication) ServicesProvider() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("servicesProvider"))
	return objectivec.Object{ID: rv}
}
func (a NSApplication) SetServicesProvider(value objectivec.IObject) {
	objc.Send[struct{}](a.ID, objc.Sel("setServicesProvider:"), value)
}

// A Boolean value indicating whether Full Keyboard Access is enabled in the
// Keyboard preference pane.
//
// # Discussion
//
// The value of this property is true if Full Keyboard Access is enabled or
// false if it’s not. You might use this value to implement your own key
// loop or to implement in-control tabbing behavior similar to [NSTableView].
// Because of the nature of the preference storage, you won’t be notified of
// changes to this property if you attempt to observe it through key-value
// observing; however, accessing this property is fairly inexpensive, so you
// can access it directly rather than caching it.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/isFullKeyboardAccessEnabled
func (a NSApplication) IsFullKeyboardAccessEnabled() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isFullKeyboardAccessEnabled"))
	return rv
}

// An array of document objects arranged according to the front-to-back
// ordering of their associated windows.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/orderedDocuments
func (a NSApplication) OrderedDocuments() []NSDocument {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("orderedDocuments"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSDocument {
		return NSDocumentFromID(id)
	})
}

// An array of window objects arranged according to their front-to-back
// ordering on the screen.
//
// # Discussion
//
// Only windows that are typically scriptable are included in the array. For
// example, panels are not included. This property is accessed during script
// command evaluation—for example, while finding the window in the script
// statement `close the second window`. For information on how your app can
// return its own array of ordered windows, see
// [application:delegateHandlesKey:].
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/orderedWindows
//
// [application:delegateHandlesKey:]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/application:delegateHandlesKey:
func (a NSApplication) OrderedWindows() []NSWindow {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("orderedWindows"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSWindow {
		return NSWindowFromID(id)
	})
}

// A boolean value indicating whether your application should suppress HDR
// content based on established policy. Built-in AppKit components such as
// NSImageView will automatically behave correctly with HDR content. You
// should use this value in conjunction with notifications
// ([NSApplicationShouldBeginSuppressingHighDynamicRangeContentNotification]
// and [NSApplicationShouldEndSuppressingHighDynamicRangeContentNotification])
// to suppress HDR content in your application when signaled to do so.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/applicationShouldSuppressHighDynamicRangeContent
func (a NSApplication) ApplicationShouldSuppressHighDynamicRangeContent() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("applicationShouldSuppressHighDynamicRangeContent"))
	return rv
}

// A Boolean value indicating whether the main menu contains an item for
// customizing the contents of the Touch Bar.
//
// # Discussion
//
// When the value of this property is true, AppKit adds a standard item to the
// app’s View menu that users can select to customize the Touch Bar
// contents, but only if a Touch Bar is present. If the View menu is
// unavailable, AppKit adds the item to either the Windows or App menu.
//
// If you prefer to provide a customize menu item, set
// [NSApplication.AutomaticCustomizeTouchBarMenuItemEnabled] to false, and
// create the menu item with an action of
// [NSApplication.ToggleTouchBarCustomizationPalette].
//
// The default value of this property is false.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/isAutomaticCustomizeTouchBarMenuItemEnabled
func (a NSApplication) IsAutomaticCustomizeTouchBarMenuItemEnabled() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isAutomaticCustomizeTouchBarMenuItemEnabled"))
	return rv
}
func (a NSApplication) SetAutomaticCustomizeTouchBarMenuItemEnabled(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setAutomaticCustomizeTouchBarMenuItemEnabled:"), value)
}

// A Boolean value indicating whether the app is hidden.
//
// # Discussion
//
// The value of this property is true if the app is hidden or false if it is
// not.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/isHidden
func (a NSApplication) IsHidden() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isHidden"))
	return rv
}

// The window that currently receives keyboard events.
//
// # Discussion
//
// The value of this property is `nil` when there is no window receiving
// keyboard events. The property might be `nil` because the app’s storyboard
// file has not yet finished loading or when the receiver is not active.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/keyWindow
func (a NSApplication) KeyWindow() INSWindow {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("keyWindow"))
	return NSWindowFromID(objc.ID(rv))
}

// The app’s main menu bar.
//
// # Discussion
//
// Use this property to assign a new menu bar for your app or to access the
// current menu bar.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/mainMenu
func (a NSApplication) MainMenu() INSMenu {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("mainMenu"))
	return NSMenuFromID(objc.ID(rv))
}
func (a NSApplication) SetMainMenu(value INSMenu) {
	objc.Send[struct{}](a.ID, objc.Sel("setMainMenu:"), value)
}

// The app’s main window.
//
// # Discussion
//
// The value in this property is `nil` when the app’s storyboard or nib file
// has not yet finished loading. It might also be `nil` when the app is
// inactive or hidden.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/mainWindow
func (a NSApplication) MainWindow() INSWindow {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("mainWindow"))
	return NSWindowFromID(objc.ID(rv))
}

// The modal window displayed by the app.
//
// # Discussion
//
// This property contains the current standalone modal window or `nil` if no
// modal window is being displayed. This property does not contain sheets that
// are attached to other windows. To retrieve a sheet, use the
// [NSWindow.AttachedSheet] method of [NSWindow].
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/modalWindow
func (a NSApplication) ModalWindow() INSWindow {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("modalWindow"))
	return NSWindowFromID(objc.ID(rv))
}

// The occlusion state of the app.
//
// # Discussion
//
// The value of this property reflects whether any part of the app’s windows
// are visible to the user. Use this information to disable expensive screen
// updates when your app is not visible.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/occlusionState-swift.property
func (a NSApplication) OcclusionState() NSApplicationOcclusionState {
	rv := objc.Send[NSApplicationOcclusionState](a.ID, objc.Sel("occlusionState"))
	return NSApplicationOcclusionState(rv)
}

// See: https://developer.apple.com/documentation/AppKit/NSApplication/isProtectedDataAvailable
func (a NSApplication) IsProtectedDataAvailable() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isProtectedDataAvailable"))
	return rv
}

// The app’s Services menu.
//
// # Discussion
//
// This property contains the app’s Services menu or `nil` if that menu has
// not been created. You can assign a new value to the property to set the
// Services menu for your app.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/servicesMenu
func (a NSApplication) ServicesMenu() INSMenu {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("servicesMenu"))
	return NSMenuFromID(objc.ID(rv))
}
func (a NSApplication) SetServicesMenu(value INSMenu) {
	objc.Send[struct{}](a.ID, objc.Sel("setServicesMenu:"), value)
}

// An array of the app’s window objects.
//
// # Discussion
//
// This property contains an array of [NSWindow] objects corresponding to all
// currently existing windows for the app. The array includes all onscreen and
// offscreen windows, whether or not they are visible on any space. There is
// no guarantee of the order of the windows in the array.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/windows
func (a NSApplication) Windows() []NSWindow {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("windows"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSWindow {
		return NSWindowFromID(id)
	})
}

// The Window menu of the app.
//
// # Return Value
//
// The window menu or `nil` if such a menu does not exist or has not yet been
// created.
//
// # Discussion
//
// This property contains the app’s Window menu or `nil` if such a menu does
// not yet exist or has not yet been created. You can use this property to
// specify the Window menu for your app.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/windowsMenu
func (a NSApplication) WindowsMenu() INSMenu {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("windowsMenu"))
	return NSMenuFromID(objc.ID(rv))
}
func (a NSApplication) SetWindowsMenu(value INSMenu) {
	objc.Send[struct{}](a.ID, objc.Sel("setWindowsMenu:"), value)
}

// Returns the application instance, creating it if it doesn’t exist yet.
//
// # Return Value
//
// The shared application object.
//
// # Discussion
//
// This method also makes a connection to the window server and completes
// other initialization. Your program should invoke this method as one of the
// first statements in `main()`; this invoking is done for you if you create
// your application with Xcode. To retrieve the [NSApplication] instance after
// it has been created, use the global variable [NSApp] or invoke this method.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/shared
//
// [NSApp]: https://developer.apple.com/documentation/AppKit/NSApp
func (_NSApplicationClass NSApplicationClass) SharedApplication() NSApplication {
	rv := objc.Send[objc.ID](objc.ID(_NSApplicationClass.class), objc.Sel("sharedApplication"))
	return NSApplicationFromID(objc.ID(rv))
}

// Protocol methods for NSAccessibilityElementProtocol

// Protocol methods for NSAccessibilityProtocol

// Protocol methods for NSAppearanceCustomization

// Protocol methods for NSMenuItemValidation

// Protocol methods for NSUserInterfaceValidations

// RestoreWindowWithIdentifierState is a synchronous wrapper around [NSApplication.RestoreWindowWithIdentifierStateCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (a NSApplication) RestoreWindowWithIdentifierState(ctx context.Context, identifier NSUserInterfaceItemIdentifier, state foundation.INSCoder) (*NSWindow, error) {
	type result struct {
		val *NSWindow
		err error
	}
	done := make(chan result, 1)
	a.RestoreWindowWithIdentifierStateCompletionHandler(identifier, state, func(val *NSWindow, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
