// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"context"
	"sync"

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
// [NSApplication.SharedApplication] class method. After creating the application object,
// the `main()` function should load your app’s main nib file and then start
// the event loop by sending the application object a [NSApplication.Run] message. If you
// create an Application project in Xcode, this `main()` function is created
// for you. The `main()` function Xcode creates begins by calling a function
// named `NSApplicationMain()`, which is functionally similar to the
// following:
//
// The [NSApplication.SharedApplication] class method initializes the display environment
// and connects your program to the window server and the display server. The
// [NSApplication] object maintains a list of all the [NSWindow] objects the
// app uses, so it can retrieve any of the app’s [NSView] objects. The
// [NSApplication.SharedApplication] method also initializes the global variable [NSApplication.NSApp],
// which you use to retrieve the [NSApplication] instance. [NSApplication.SharedApplication]
// only performs the initialization once. If you invoke it more than once, it
// returns the application object it created previously.
//
// The shared [NSApplication] object performs the important task of receiving
// events from the window server and distributing them to the proper
// [NSResponder] objects. [NSApplication.NSApp] translates an event into an [NSEvent]
// object, then forwards the event object to the affected [NSWindow] object.
// All keyboard and mouse events go directly to the [NSWindow] object
// associated with the event. The only exception to this rule is if the
// Command key is pressed when a key-down event occurs; in this case, every
// [NSWindow] object has an opportunity to respond to the event. When a window
// object receives an [NSEvent] object from [NSApplication.NSApp], it distributes it to the
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
// initialization (or [NSApplication.SharedApplication]) and [NSApplication.Run] methods. Similarly, the
// methods AppKit adds to [Bundle] employ `@autorelease` blocks during the
// loading of nib files. These `@autorelease` blocks aren’t accessible
// outside the scope of the respective [NSApplication] and [Bundle] methods.
// Typically, an app creates objects either while the event loop is running or
// by loading objects from nib files, so this lack of access usually isn’t a
// problem. However, if you do need to use Cocoa classes within the `main()`
// function itself (other than to load nib files or to instantiate
// [NSApplication]), you should create an `@autorelease` block to contain the
// code using the classes.
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
// to subclass [NSApplication].
//
// To use a custom subclass of [NSApplication], send [NSApplication.SharedApplication] to
// your subclass rather than directly to [NSApplication]. If you create your
// app in Xcode, you can accomplish this by setting your custom app class to
// be the principal class. In Xcode, double-click the app target in the Groups
// and Files list to open the Info window for the target. Then display the
// Properties pane of the window and replace “NSApplication” in the
// Principal Class field with the name of your custom class. The
// [NSApplicationMain] function sends [NSApplication.SharedApplication] to the principal
// class to obtain the global app instance ([NSApplication.NSApp])—which in this case will
// be an instance of your custom subclass of [NSApplication].
//
// # Methods to override
//
// Generally, you subclass [NSApplication] to provide your own special
// responses to messages that are routinely sent to the global app object
// ([NSApplication.NSApp]). [NSApplication] doesn’t have primitive methods in the sense of
// methods that you must override in your subclass. Here are four methods that
// are possible candidates for overriding:
//
// - Override [NSApplication.Run] if you want the app to manage the main event loop
// differently than it does by default. (This a critical and complex task,
// however, that you should only attempt with good reason). - Override
// [NSApplication.SendEvent] if you want to change how events are dispatched or perform some
// special event processing. - Override [NSApplication.RequestUserAttention] if you want to
// modify how your app attracts the attention of the user (for example,
// offering an alternative to the bouncing app icon in the Dock). - Override
// [NSApplication.TargetForAction] to substitute another object for the target of an action
// message.
//
// # Special considerations
//
// The global app object uses `@autorelease` blocks in its [NSApplication.Run] method; if
// you override this method, you’ll need to create your own `@autorelease`
// blocks.
//
// Do not override [NSApplication.SharedApplication]. The default implementation, which is
// essential to app behavior, is too complex to duplicate on your own.
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
// # Getting the shared app object
//
//   - [NSApplication.NSApp]: The global variable for the shared app instance.
//   - [NSApplication.SetNSApp]
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
// # Getting the shared app object
//
//   - [INSApplication.NSApp]: The global variable for the shared app instance.
//   - [INSApplication.SetNSApp]
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
// # Scripting your app
//
//   - [INSApplication.OrderedDocuments]: An array of document objects arranged according to the front-to-back ordering of their associated windows.
//   - [INSApplication.OrderedWindows]: An array of window objects arranged according to their front-to-back ordering on the screen.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication
type INSApplication interface {
	INSResponder
	NSAppearanceCustomization

	// Topic: Getting the shared app object

	// The global variable for the shared app instance.
	NSApp() INSApplication
	SetNSApp(value INSApplication)

	// Topic: Managing the app’s behavior

	// The app delegate object.
	Delegate() NSApplicationDelegate
	SetDelegate(value NSApplicationDelegate)

	// Topic: Managing the event loop

	// Returns the next event matching a given mask, or `nil` if no such event is found before a specified expiration date.
	NextEventMatchingMaskUntilDateInModeDequeue(mask NSEventMask, expiration foundation.NSDate, mode foundation.NSString, deqFlag bool) INSEvent
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

	// Topic: Scripting your app

	// An array of document objects arranged according to the front-to-back ordering of their associated windows.
	OrderedDocuments() []NSDocument
	// An array of window objects arranged according to their front-to-back ordering on the screen.
	OrderedWindows() []NSWindow

	// Indicates the activation policy of the application.
	ActivationPolicy() NSApplicationActivationPolicy
	SetActivationPolicy(value NSApplicationActivationPolicy)
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
	// Aborts the event loop started by [runModal(for:)](<doc://com.apple.appkit/documentation/AppKit/NSApplication/runModal(for:)>) or [runModalSession(_:)](<doc://com.apple.appkit/documentation/AppKit/NSApplication/runModalSession(_:)>).
	AbortModal()
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
	EnumerateWindowsWithOptionsUsingBlock(options NSWindowListOptions, block WindowHandler)
	// Allows an app to extend its state restoration period.
	ExtendStateRestoration()
	// Hides all the receiver’s windows, and the next app in line is activated.
	Hide(sender objectivec.IObject)
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
	// Runs a given modal session, as defined in a previous invocation of [beginModalSession(for:)](<doc://com.apple.appkit/documentation/AppKit/NSApplication/beginModalSession(for:)>).
	RunModalSession(session NSModalSession) NSModalResponse
	// Displays the receiver’s page layout panel, an instance of [NSPageLayout].
	RunPageLayout(sender objectivec.IObject)
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
	// Sends an [update()](<doc://com.apple.appkit/documentation/AppKit/NSWindow/update()>) message to each onscreen window.
	UpdateWindows()
	// Updates the Window menu item for a given window to reflect the edited status of that window.
	UpdateWindowsItem(win INSWindow)
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
// together to create this mask. The [DiscardEventsMatchingMaskBeforeEvent]
// method also lists several of these constants.
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
func (a NSApplication) NextEventMatchingMaskUntilDateInModeDequeue(mask NSEventMask, expiration foundation.NSDate, mode foundation.NSString, deqFlag bool) INSEvent {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("nextEventMatchingMask:untilDate:inMode:dequeue:"), mask, expiration, mode, deqFlag)
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
// The loop continues until a [Stop] or [Terminate] message is received. Upon
// each iteration through the loop, the next available event from the window
// server is stored and then dispatched by sending it to [NSApp] using
// [SendEvent].
//
// After creating the [NSApplication] object, the `main` function should load
// your app’s main nib file and then start the event loop by sending the
// [NSApplication] object a [Run] message. If you create an Cocoa app project
// in Xcode, this `main` function is implemented for you.
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
// The [Run] method calls this method before it starts the event loop. When
// this method begins, it posts an [willFinishLaunchingNotification] to the
// default notification center. If you override [FinishLaunching], the
// subclass method should invoke the superclass method.
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
// loop, the app object exits out of the [Run] method, thereby returning
// control to the `main()` function. If you call this method from within a
// modal event loop, it will exit the modal loop instead of the main event
// loop.
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
// You rarely invoke [SendEvent] directly, although you might want to override
// this method to perform some action on every event. [SendEvent] messages are
// sent from the main event loop (the [Run] method). [SendEvent] is the method
// that dispatches events to the appropriate responders—[NSApp] handles app
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
// If `aTarget` is `nil`, [SharedApplication] looks for an object that can
// respond to the message—that is, an object that implements a method
// matching `anAction`. It begins with the first responder of the key window.
// If the first responder can’t respond, it tries the first responder’s
// next responder and continues following next responder links up the
// responder chain. If none of the objects in the key window’s responder
// chain can handle the message, [SharedApplication] attempts to send the
// message to the key window’s delegate.
//
// If the delegate doesn’t respond and the main window is different from the
// key window, [SharedApplication] begins again with the first responder in
// the main window. If objects in the main window can’t respond,
// [SharedApplication] attempts to send the message to the main window’s
// delegate. If still no object has responded, [SharedApplication] tries to
// handle the message itself. If [SharedApplication] can’t respond, it
// attempts to send the message to its own delegate.
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
// [ReplyToApplicationShouldTerminate] method is called with the value true or
// false. If the [ApplicationShouldTerminate] method returns [NSTerminateNow],
// this method posts a [willTerminateNotification] notification to the default
// notification center.
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
// call [YieldActivationToApplication] or equivalent before the target app
// invokes [Activate].
//
// Invoking [Activate] on an already-active application cancels any pending
// activation yields by the receiver.
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
// must request activation in the future by calling [Activate] or equivalent.
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
// must request activation in the future by calling [Activate] or equivalent.
//
// Use this method to yield activation to apps that aren’t running at the
// time the method invokes. If it’s known that the target application is
// running, use [YieldActivationToApplication] instead.
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
// this method once, and never pair it with an [EnableRelaunchOnLogin] method.
//
// If your app shouldn’t be relaunched because it triggers a restart, for
// example an installer, then the recommended usage is to invoke this method
// immediately before you attempt to trigger a restart, and
// [EnableRelaunchOnLogin] immediately after. This is because the user may
// cancel restarting; if the user later restarts for another reason, then your
// app should be brought back.
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
// request later using the [CancelUserAttentionRequest] method.
//
// # Discussion
//
// Activating the app cancels the user attention request. A spoken
// notification will occur if spoken notifications are enabled. Sending
// [RequestUserAttention] to an app that is already active has no effect.
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
// request: The request identifier returned by the [RequestUserAttention] method.
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
	rv := objc.Send[bool](a.ID, objc.Sel("searchString:inUserInterfaceItemString:searchRange:foundRange:"), objc.String(searchString), objc.String(stringToSearch), searchRange, foundRange)
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

// Aborts the event loop started by [RunModalForWindow] or [RunModalSession].
//
// # Discussion
//
// When stopped with this method, [RunModalForWindow] and [RunModalSession]
// return [NSModalResponseAbort].
//
// [AbortModal] must be used instead of [StopModal] or [StopModalWithCode]
// when you need to stop a modal event loop from anywhere other than a callout
// from that event loop. In other words, if you want to stop the loop in
// response to a user’s actions within the modal window, use [StopModal];
// otherwise, use [AbortModal]. For example, use [AbortModal] when running in
// a different thread from AppKit’s main thread or when responding to an
// [NSTimer] that you have added to the [NSModalPanelRunLoopMode] mode of the
// default [NSRunLoop].
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/abortModal()
func (a NSApplication) AbortModal() {
	objc.Send[objc.ID](a.ID, objc.Sel("abortModal"))
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
// path (the way the [NSWindow] method [SetTitleWithRepresentedFilename] shows
// a title)
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
// onscreen using the [NSWindow] method [Center].
//
// The [BeginModalSessionForWindow] method only sets up the modal session. To
// actually run the session, use [RunModalSession].
// [BeginModalSessionForWindow] should be balanced by [EndModalSession]. Make
// sure these two messages are sent within the same exception-handling scope.
// That is, if you send [BeginModalSessionForWindow] inside an `NS_DURING`
// construct, you must send [EndModalSession] before `NS_ENDHANDLER`.
//
// If an exception is raised, [BeginModalSessionForWindow] arranges for proper
// cleanup. Do not use `NS_DURING` constructs to send an [EndModalSession]
// message in the event of an exception.
//
// A loop using these methods is similar to a modal event loop run with
// [RunModalForWindow], except the app can continue processing between method
// invocations.
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
// path (the way the [NSWindow] method [SetTitleWithRepresentedFilename] shows
// a title)
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
// You call [ExtendStateRestoration] within your implementation of
// [RestoreWindowWithIdentifierStateCompletionHandler]. You would then call
// `completeStateRestoration` some time after the window is fully restored. If
// the app crashes in the interim, then it may offer to discard restorable
// state on the next launch.
//
// The [ExtendStateRestoration] and `completeStateRestoration` method act as a
// counter. Each call to [ExtendStateRestoration]increments the counter, and
// must be matched with a corresponding call to `completeStateRestoration`
// which decrements it. When the counter reaches zero, the app is considered
// to have been fully restored, and any further calls are silently ignored.
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
// [BeginModalSessionForWindow].
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
func (a NSApplication) EnumerateWindowsWithOptionsUsingBlock(options NSWindowListOptions, block WindowHandler) {
	_block1, _ := NewWindowBlock(block)
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
// web page, you may use this method and methods to [CompleteStateRestoration]
// to extend the period of this crash protection beyond the default.
//
// You call `extendStateRestoration` within your implementation of
// [RestoreWindowWithIdentifierStateCompletionHandler]. You would then call
// [CompleteStateRestoration] some time after the window is fully restored. If
// the app crashes in the interim, then it may offer to discard restorable
// state on the next launch.
//
// The `extendStateRestoration` and [CompleteStateRestoration] methods act as
// a counter. Each call to `extendStateRestoration` increments the counter,
// and must be matched with a corresponding call to [CompleteStateRestoration]
// which decrements it. When the counter reaches zero, the app is considered
// to have been fully restored, and any further calls are silently ignored.
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
// This method calls [OrderFrontStandardAboutPanelWithOptions] with a `nil`
// argument. See [OrderFrontStandardAboutPanelWithOptions] for a description
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
// again. Use the [ExcludedFromWindowsMenu] method of [NSWindow] if you want
// the item to remain excluded from the Window menu.
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
// window is centered on the screen using the value in its [Center] method and
// made visible and key. If it is already visible, it is simply made key.
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
// You can exit the modal loop by calling the [StopModal],
// [StopModalWithCode], or [AbortModal] methods from your modal window code.
// If you use the [StopModalWithCode] method to stop the modal event loop,
// this method returns the argument passed to [StopModalWithCode]. If you use
// [StopModal] instead, this method returns the constant [stop]. If you use
// [AbortModal], this method returns the constant [abort].
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
// [BeginModalSessionForWindow].
//
// session: A pointer to the modal session structure returned by the
// [BeginModalSessionForWindow] method for the window to be displayed.
//
// # Return Value
//
// An integer indicating the reason that this method returned. See the
// discussion for a description of possible return values.
//
// # Discussion
//
// A loop that uses this method is similar in some ways to a modal event loop
// run with [RunModalForWindow], except with this method your code can do some
// additional work between method invocations. When you invoke this method,
// events for the [NSWindow] object of this session are dispatched as normal.
// This method returns when there are no more events. You must invoke this
// method frequently enough in your loop that the window remains responsive to
// events. However, you should not invoke this method in a tight loop because
// it returns immediately if there are no events, and consequently you could
// end up polling for events rather than blocking.
//
// Typically, you use this method in situations where you want to do some
// additional processing on the current thread while the modal loop runs. For
// example, while processing a large data set, you might want to use a modal
// dialog to display progress and give the user a chance to cancel the
// operation. If you want to display a modal dialog and do not need to do any
// additional work in parallel, use [RunModalForWindow] instead. When there
// are no pending events, that method waits idly instead of consuming CPU
// time.
//
// The following code shows a sample loop you can use in your code:
//
// If the modal session was not stopped, this method returns
// [NSModalResponseContinue]. At this point, your app can do some work before
// the next invocation of [RunModalSession] (as indicated in the example’s
// `doSomeWork` call). If [StopModal] was invoked as the result of event
// processing, [RunModalSession] returns [NSModalResponseStop]. If
// [StopModalWithCode] was invoked, this method returns the value passed to
// [StopModalWithCode]. If [AbortModal] was invoked, this method returns
// [NSModalResponseAbort].
//
// The window is placed on the screen and made key as a result of the
// [RunModalSession] message. Do not send a separate [KeyAndOrderFront]
// message.
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
// [RunModalForWindow] or [BeginModalSessionForWindow]. When
// [RunModalForWindow] is stopped with this method, it returns
// [NSModalResponseStop]. In macOS 10.9 and later, you can use this method to
// stop a [RunModalForWindow] loop outside of an event callback, such as from
// within a method repeatedly invoked by an [Timer] object or a method running
// in a different thread.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/stopModal()
//
// [Timer]: https://developer.apple.com/documentation/Foundation/Timer
func (a NSApplication) StopModal() {
	objc.Send[objc.ID](a.ID, objc.Sel("stopModal"))
}

// Stops a modal event loop, allowing you to return a custom result code.
//
// returnCode: The result code you want returned from the [RunModalForWindow] or
// [RunModalSession] method. The meaning of this result code is up to you.
//
// # Discussion
//
// This method should always be paired with a previous invocation of
// [RunModalForWindow] or [BeginModalSessionForWindow]. When
// [RunModalForWindow] is stopped with this method, it returns the given
// `returnCode`. In macOS 10.9 and later, you can use this method to stop a
// [RunModalForWindow] loop outside of an event callback, such as from within
// a method repeatedly invoked by an [Timer] object or a method running in a
// different thread.
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
// Invokes [UnhideWithoutActivation].
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

// Sends an [Update] message to each onscreen window.
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
// [WindowWithWindowNumber] may return `nil` for window numbers found using
// [WindowNumbersWithOptions] if there is no corresponding window object owned
// by your app—for example, the menu bar.
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

// The global variable for the shared app instance.
//
// See: https://developer.apple.com/documentation/appkit/nsapp
func (a NSApplication) NSApp() INSApplication {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("NSApp"))
	return NSApplicationFromID(objc.ID(rv))
}
func (a NSApplication) SetNSApp(value INSApplication) {
	objc.Send[struct{}](a.ID, objc.Sel("setNSApp:"), value)
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
// false when it’s not. Calling the [Stop] method sets the value to false.
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
// directly. Call the [RegisterForRemoteNotificationTypes] method to register
// your app with Apple Push Notification Service and request the notification
// types your app supports. macOS delivers only notifications of types that
// the app supports. For a list of possible values, see
// [NSApplication.RemoteNotificationType].
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
// to the [Appearance] property, the app inherits the system’s effective
// appearance.
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
// [PresentationOptions] property or the [SetSystemUIMode] function. - Another
// app is active and makes presentation changes of its own. - Another app
// becomes active and causes the active set of presentation options to change.
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
// invalid combination of option flags raises an [InvalidArgumentException]
// exception. See the constants for a description of the valid combinations.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/presentationOptions-swift.property
//
// [NSApplication.PresentationOptions]: https://developer.apple.com/documentation/AppKit/NSApplication/PresentationOptions-swift.struct
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

// Indicates the activation policy of the application.
//
// See: https://developer.apple.com/documentation/appkit/nsrunningapplication/activationpolicy
func (a NSApplication) ActivationPolicy() NSApplicationActivationPolicy {
	rv := objc.Send[NSApplicationActivationPolicy](a.ID, objc.Sel("activationPolicy"))
	return NSApplicationActivationPolicy(rv)
}
func (a NSApplication) SetActivationPolicy(value NSApplicationActivationPolicy) {
	objc.Send[struct{}](a.ID, objc.Sel("setActivationPolicy:"), value)
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
// [AutomaticCustomizeTouchBarMenuItemEnabled] to false, and create the menu
// item with an action of [ToggleTouchBarCustomizationPalette].
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
// are attached to other windows. To retrieve a sheet, use the [AttachedSheet]
// method of [NSWindow].
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
func (_NSApplicationClass NSApplicationClass) SharedApplication() NSApplication {
	rv := objc.Send[objc.ID](objc.ID(_NSApplicationClass.class), objc.Sel("sharedApplication"))
	return NSApplicationFromID(objc.ID(rv))
}

// The priority at which windows are displayed.
//
// See: https://developer.apple.com/documentation/appkit/nsapplication/displaywindowrunloopordering
func (_NSApplicationClass NSApplicationClass) DisplayWindowRunLoopOrdering() int {
	rv := objc.Send[int](objc.ID(_NSApplicationClass.class), objc.Sel("NSDisplayWindowRunLoopOrdering"))
	return rv
}
func (_NSApplicationClass NSApplicationClass) SetNSDisplayWindowRunLoopOrdering(value int) {
	objc.Send[struct{}](objc.ID(_NSApplicationClass.class), objc.Sel("setNSDisplayWindowRunLoopOrdering:"), value)
}

// The priority at which cursor rects are reset.
//
// See: https://developer.apple.com/documentation/appkit/nsapplication/resetcursorrectsrunloopordering
func (_NSApplicationClass NSApplicationClass) ResetCursorRectsRunLoopOrdering() int {
	rv := objc.Send[int](objc.ID(_NSApplicationClass.class), objc.Sel("NSResetCursorRectsRunLoopOrdering"))
	return rv
}
func (_NSApplicationClass NSApplicationClass) SetNSResetCursorRectsRunLoopOrdering(value int) {
	objc.Send[struct{}](objc.ID(_NSApplicationClass.class), objc.Sel("setNSResetCursorRectsRunLoopOrdering:"), value)
}

// Run-loop message priority for handling window updates.
//
// See: https://developer.apple.com/documentation/appkit/nsapplication/updatewindowsrunloopordering
func (_NSApplicationClass NSApplicationClass) UpdateWindowsRunLoopOrdering() int {
	rv := objc.Send[int](objc.ID(_NSApplicationClass.class), objc.Sel("NSUpdateWindowsRunLoopOrdering"))
	return rv
}
func (_NSApplicationClass NSApplicationClass) SetNSUpdateWindowsRunLoopOrdering(value int) {
	objc.Send[struct{}](objc.ID(_NSApplicationClass.class), objc.Sel("setNSUpdateWindowsRunLoopOrdering:"), value)
}

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
