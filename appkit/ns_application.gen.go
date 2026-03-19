// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"context"
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
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
// [Bundle]: https://developer.apple.com/documentation/Foundation/Bundle
// [Cocoa Scripting Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ScriptableCocoaApplications/SApps_intro/SAppsIntro.html#//apple_ref/doc/uid/TP40002164
// [Delegation]: https://developer.apple.com/library/archive/documentation/General/Conceptual/DevPedia-CocoaCore/Delegation.html#//apple_ref/doc/uid/TP40008195-CH14
// [How Cocoa Applications Handle Apple Events]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ScriptableCocoaApplications/SApps_handle_AEs/SAppsHandleAEs.html#//apple_ref/doc/uid/20001239
// [NSAppleEventManager]: https://developer.apple.com/documentation/Foundation/NSAppleEventManager
// [addObserver(_:selector:name:object:)]: https://developer.apple.com/documentation/Foundation/NotificationCenter/addObserver(_:selector:name:object:)
// [didFinishLaunchingNotification]: https://developer.apple.com/documentation/AppKit/NSApplication/didFinishLaunchingNotification
// [willFinishLaunchingNotification]: https://developer.apple.com/documentation/AppKit/NSApplication/willFinishLaunchingNotification
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
//   - [NSApplication.Running]: A Boolean value indicating whether the main event loop is running.
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
//   - [NSApplication.Active]: A Boolean value indicating whether this is the active app.
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
//   - [NSApplication.RegisteredForRemoteNotifications]: A Boolean value indicating whether the app is registered with Apple Push Notification service (APNs).
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
//   - [NSApplication.FullKeyboardAccessEnabled]: A Boolean value indicating whether Full Keyboard Access is enabled in the Keyboard preference pane.
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
//   - [INSApplication.Running]: A Boolean value indicating whether the main event loop is running.
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
//   - [INSApplication.Active]: A Boolean value indicating whether this is the active app.
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
//   - [INSApplication.RegisteredForRemoteNotifications]: A Boolean value indicating whether the app is registered with Apple Push Notification service (APNs).
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
//   - [INSApplication.FullKeyboardAccessEnabled]: A Boolean value indicating whether Full Keyboard Access is enabled in the Keyboard preference pane.
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
	NSAccessibilityElementProtocol
	NSAccessibilityProtocol
	NSAppearanceCustomization
	NSMenuItemValidation
	NSUserInterfaceValidations

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
	NextEventMatchingMaskUntilDateInModeDequeue(mask NSEventMask, expiration foundation.INSDate, mode foundation.NSString, deqFlag bool) INSEvent
	// Removes all events matching the given mask and generated before the specified event.
	DiscardEventsMatchingMaskBeforeEvent(mask NSEventMask, lastEvent INSEvent)
	// The last event object that the app retrieved from the event queue.
	CurrentEvent() INSEvent
	// A Boolean value indicating whether the main event loop is running.
	Running() bool
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
	Active() bool
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
	RegisteredForRemoteNotifications() bool

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
	RegisterUserInterfaceItemSearchHandler(handler ErrorHandler)
	// Searches for the string in the user interface.
	SearchStringInUserInterfaceItemStringSearchRangeFoundRange(searchString string, stringToSearch string, searchRange foundation.NSRange, foundRange *foundation.NSRange) bool
	// Unregister an object that provides help data to your app.
	UnregisterUserInterfaceItemSearchHandler(handler ErrorHandler)
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
	FullKeyboardAccessEnabled() bool

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
	// The graphics context associated with the app.
	Context() INSGraphicsContext
	SetContext(value INSGraphicsContext)
	// A Boolean value indicating whether the main menu contains an item for customizing the contents of the Touch Bar.
	IsAutomaticCustomizeTouchBarMenuItemEnabled() bool
	SetIsAutomaticCustomizeTouchBarMenuItemEnabled(value bool)
	// A Boolean value indicating whether the app is hidden.
	IsHidden() bool
	SetIsHidden(value bool)
	IsProtectedDataAvailable() bool
	SetIsProtectedDataAvailable(value bool)
	// The window that currently receives keyboard events.
	KeyWindow() INSWindow
	SetKeyWindow(value INSWindow)
	// The app’s main menu bar.
	MainMenu() INSMenu
	SetMainMenu(value INSMenu)
	// The app’s main window.
	MainWindow() INSWindow
	SetMainWindow(value INSWindow)
	// The modal window displayed by the app.
	ModalWindow() INSWindow
	SetModalWindow(value INSWindow)
	// The occlusion state of the app.
	OcclusionState() objectivec.IObject
	SetOcclusionState(value objectivec.IObject)
	// The app’s Services menu.
	ServicesMenu() INSMenu
	SetServicesMenu(value INSMenu)
	// An array of the app’s window objects.
	Windows() INSWindow
	SetWindows(value INSWindow)
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
	BeginModalSessionForWindow(window INSWindow) objectivec.IObject
	// Changes the item for a given window in the Window menu to a given string.
	ChangeWindowsItemTitleFilename(win INSWindow, string_ string, isFilename bool)
	// Completes the extended state restoration.
	CompleteStateRestoration()
	// Finishes a modal session.
	EndModalSession(session objectivec.IObject)
	// Executes a block for each of the app’s windows.
	EnumerateWindowsWithOptionsUsingBlock(options uint, block WindowHandler)
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
	RunModalSession(session objectivec.IObject) NSModalResponse
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
// //
// [distantPast]: https://developer.apple.com/documentation/Foundation/NSDate/distantPast
//
// mode: The run loop mode in which to run while looking for events. The mode you
// specify also determines which timers and run-loop observers may fire while
// the app waits for the event.
//
// deqFlag: Specify [true] if you want the event removed from the queue.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
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
func (a NSApplication) NextEventMatchingMaskUntilDateInModeDequeue(mask NSEventMask, expiration foundation.INSDate, mode foundation.NSString, deqFlag bool) INSEvent {
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
// [willFinishLaunchingNotification]: https://developer.apple.com/documentation/AppKit/NSApplication/willFinishLaunchingNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/finishLaunching()
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
// atStart: Specify [true] to add the event to the front of the queue; otherwise,
// specify [false] to add the event to the back of the queue.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// [true] if the action was successfully sent; otherwise [false]. This method
// also returns [false] if `anAction` is `nil`.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// [ApplicationShouldTerminate] returns
// [NSApplication.TerminateReply.terminateCancel], the termination process is
// aborted and control is handed back to the main event loop. If the method
// returns [NSApplication.TerminateReply.terminateLater], the app runs its run
// loop in the [NSModalPanelRunLoopMode] mode until the
// [ReplyToApplicationShouldTerminate] method is called with the value [true]
// or [false]. If the [ApplicationShouldTerminate] method returns
// [NSApplication.TerminateReply.terminateNow], this method posts a
// [willTerminateNotification] notification to the default notification
// center.
// 
// Don’t bother to put final cleanup code in your app’s `main()`
// function—it will never be executed. If cleanup is necessary, perform that
// cleanup in the delegate’s [ApplicationWillTerminate] method.
//
// [NSApplication.TerminateReply.terminateCancel]: https://developer.apple.com/documentation/AppKit/NSApplication/TerminateReply/terminateCancel
// [NSApplication.TerminateReply.terminateLater]: https://developer.apple.com/documentation/AppKit/NSApplication/TerminateReply/terminateLater
// [NSApplication.TerminateReply.terminateNow]: https://developer.apple.com/documentation/AppKit/NSApplication/TerminateReply/terminateNow
// [NSModalPanelRunLoopMode]: https://developer.apple.com/documentation/AppKit/NSModalPanelRunLoopMode
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
// [willTerminateNotification]: https://developer.apple.com/documentation/AppKit/NSApplication/willTerminateNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/terminate(_:)
func (a NSApplication) Terminate(sender objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("terminate:"), sender)
}
// Responds to [NSTerminateLater] once the app knows whether it can terminate.
//
// shouldTerminate: Specify [true] if you want the app to terminate; otherwise, specify
// [false].
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// [Registering your app with APNs]: https://developer.apple.com/documentation/UserNotifications/registering-your-app-with-apns
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/registerForRemoteNotifications()
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
// //
// [NSApplication.RemoteNotificationType]: https://developer.apple.com/documentation/AppKit/NSApplication/RemoteNotificationType
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
// //
// [NSApplication.RequestUserAttentionType]: https://developer.apple.com/documentation/AppKit/NSApplication/RequestUserAttentionType
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
// //
// [NSApplication.DelegateReply]: https://developer.apple.com/documentation/AppKit/NSApplication/DelegateReply
//
// # Discussion
// 
// Delegates should invoke this method if an error is encountered in the
// [ApplicationOpenFiles] or
// [ApplicationPrintFilesWithSettingsShowPrintPanels] delegate methods.
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/reply(toOpenOrPrint:)
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
func (a NSApplication) RegisterUserInterfaceItemSearchHandler(handler ErrorHandler) {
_block0, _cleanup0 := NewErrorBlock(handler)
	defer _cleanup0()
	objc.Send[objc.ID](a.ID, objc.Sel("registerUserInterfaceItemSearchHandler:"), _block0)
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
// [true] if the searchString is matched, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
func (a NSApplication) UnregisterUserInterfaceItemSearchHandler(handler ErrorHandler) {
_block0, _cleanup0 := NewErrorBlock(handler)
	defer _cleanup0()
	objc.Send[objc.ID](a.ID, objc.Sel("unregisterUserInterfaceItemSearchHandler:"), _block0)
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
// [Specifying the Comprehensive Help File]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/OnlineHelp/Tasks/SpecifyHelpFile.html#//apple_ref/doc/uid/20000020
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/showHelp(_:)
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
// isFilename: If [false], `aString` appears literally in the menu; otherwise, `aString`
// is assumed to be a converted pathname with the name of the file preceding
// the path (the way the [NSWindow] method [SetTitleWithRepresentedFilename]
// shows a title)
// //
// [false]: https://developer.apple.com/documentation/Swift/false
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
func (a NSApplication) BeginModalSessionForWindow(window INSWindow) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("beginModalSessionForWindow:"), window)
	return objectivec.Object{ID: rv}
}
// Changes the item for a given window in the Window menu to a given string.
//
// win: The window whose title you want to change in the Window menu. If `aWindow`
// is not in the Window menu, this method adds it.
//
// string: The string to display for the window’s menu item. How the string is
// interpreted is dependent on the value in the `isFilename` parameter.
//
// isFilename: If [false], `aString` appears literally in the menu; otherwise, `aString`
// is assumed to be a converted pathname with the name of the file preceding
// the path (the way the [NSWindow] method [SetTitleWithRepresentedFilename]
// shows a title)
// //
// [false]: https://developer.apple.com/documentation/Swift/false
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
func (a NSApplication) EndModalSession(session objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("endModalSession:"), session)
}
// Executes a block for each of the app’s windows.
//
// options: A constant that indicates window ordering. See
// [NSApplication.WindowListOptions] for possible values.
// //
// [NSApplication.WindowListOptions]: https://developer.apple.com/documentation/AppKit/NSApplication/WindowListOptions
//
// block: The block to execute for each window. The block takes the following
// parameters:
// 
// window: The window for which to execute the block. stop: A Boolean value
// that stops the enumeration early when set to [true] (the default value is
// [false]).
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/enumerateWindows(options:using:)
func (a NSApplication) EnumerateWindowsWithOptionsUsingBlock(options uint, block WindowHandler) {
_block1, _cleanup1 := NewWindowBlock(block)
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
// [didHideNotification]: https://developer.apple.com/documentation/AppKit/NSApplication/didHideNotification
// [willHideNotification]: https://developer.apple.com/documentation/AppKit/NSApplication/willHideNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/hide(_:)
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
// of keys, see [NSApplication.AboutPanelOptionKey].
// //
// [NSApplication.AboutPanelOptionKey]: https://developer.apple.com/documentation/AppKit/NSApplication/AboutPanelOptionKey
//
// # Discussion
// 
// In addition to the keys in AboutPanelOptionKey, you may also include the
// following key in `optionsDictionary`:
// 
// - `@``"Copyright"`: An [NSString] object with a line of copyright
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
// [true] if the window was restored; otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
_block2, _cleanup2 := NewWindowErrorBlock(completionHandler)
	defer _cleanup2()
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
// [abort]: https://developer.apple.com/documentation/AppKit/NSApplication/ModalResponse/abort
// [stop]: https://developer.apple.com/documentation/AppKit/NSApplication/ModalResponse/stop
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/runModal(for:)
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
func (a NSApplication) RunModalSession(session objectivec.IObject) NSModalResponse {
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
// needUpdate: If [true], the receiver’s windows are updated after an event is
// processed.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
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
// [Timer]: https://developer.apple.com/documentation/Foundation/Timer
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/stopModal()
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
// [Timer]: https://developer.apple.com/documentation/Foundation/Timer
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/stopModal(withCode:)
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
// [didUnhideNotification]: https://developer.apple.com/documentation/AppKit/NSApplication/didUnhideNotification
// [willUnhideNotification]: https://developer.apple.com/documentation/AppKit/NSApplication/willUnhideNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/unhideWithoutActivation()
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
// [didUpdateNotification]: https://developer.apple.com/documentation/AppKit/NSApplication/didUpdateNotification
// [willUpdateNotification]: https://developer.apple.com/documentation/AppKit/NSApplication/willUpdateNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/updateWindows()
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
// [true] to enable `menuItem`, [false] to disable it.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// [true] if the user interface item should be enabled, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// [Thread]: https://developer.apple.com/documentation/Foundation/Thread
// [detachNewThreadSelector(_:toTarget:with:)]: https://developer.apple.com/documentation/Foundation/Thread/detachNewThreadSelector(_:toTarget:with:)
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/detachDrawingThread(_:toTarget:with:)
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
// The value of this property is [true] when the main event loop is running or
// [false] when it’s not. Calling the [Stop] method sets the value to
// [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/isRunning
func (a NSApplication) Running() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isRunning"))
	return rv
}
// A Boolean value indicating whether this is the active app.
//
// # Discussion
// 
// The value of this property is [true] if the app is active or [false] if
// it’s not.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/isActive
func (a NSApplication) Active() bool {
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
// [NSApplication.RemoteNotificationType]: https://developer.apple.com/documentation/AppKit/NSApplication/RemoteNotificationType
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/enabledRemoteNotificationTypes
func (a NSApplication) EnabledRemoteNotificationTypes() NSRemoteNotificationType {
	rv := objc.Send[NSRemoteNotificationType](a.ID, objc.Sel("enabledRemoteNotificationTypes"))
	return NSRemoteNotificationType(rv)
}
// A Boolean value indicating whether the app is registered with Apple Push
// Notification service (APNs).
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/isRegisteredForRemoteNotifications
func (a NSApplication) RegisteredForRemoteNotifications() bool {
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
// [NSApplication.PresentationOptions]: https://developer.apple.com/documentation/AppKit/NSApplication/PresentationOptions-swift.struct
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/currentSystemPresentationOptions
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
// [NSApplication.PresentationOptions]: https://developer.apple.com/documentation/AppKit/NSApplication/PresentationOptions-swift.struct
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/presentationOptions-swift.property
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
// [NSUserInterfaceLayoutDirection]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceLayoutDirection
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/userInterfaceLayoutDirection
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
// The value of this property is [true] if Full Keyboard Access is enabled or
// [false] if it’s not. You might use this value to implement your own key
// loop or to implement in-control tabbing behavior similar to [NSTableView].
// Because of the nature of the preference storage, you won’t be notified of
// changes to this property if you attempt to observe it through key-value
// observing; however, accessing this property is fairly inexpensive, so you
// can access it directly rather than caching it.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/isFullKeyboardAccessEnabled
func (a NSApplication) FullKeyboardAccessEnabled() bool {
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
// [application:delegateHandlesKey:]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/application:delegateHandlesKey:
//
// See: https://developer.apple.com/documentation/AppKit/NSApplication/orderedWindows
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
// The graphics context associated with the app.
//
// See: https://developer.apple.com/documentation/appkit/nsapplication/context
func (a NSApplication) Context() INSGraphicsContext {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("context"))
	return NSGraphicsContextFromID(objc.ID(rv))
}
func (a NSApplication) SetContext(value INSGraphicsContext) {
	objc.Send[struct{}](a.ID, objc.Sel("setContext:"), value)
}
// A Boolean value indicating whether the main menu contains an item for
// customizing the contents of the Touch Bar.
//
// See: https://developer.apple.com/documentation/appkit/nsapplication/isautomaticcustomizetouchbarmenuitemenabled
func (a NSApplication) IsAutomaticCustomizeTouchBarMenuItemEnabled() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("automaticCustomizeTouchBarMenuItemEnabled"))
	return rv
}
func (a NSApplication) SetIsAutomaticCustomizeTouchBarMenuItemEnabled(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setAutomaticCustomizeTouchBarMenuItemEnabled:"), value)
}
// A Boolean value indicating whether the app is hidden.
//
// See: https://developer.apple.com/documentation/appkit/nsapplication/ishidden
func (a NSApplication) IsHidden() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("hidden"))
	return rv
}
func (a NSApplication) SetIsHidden(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setHidden:"), value)
}
// See: https://developer.apple.com/documentation/appkit/nsapplication/isprotecteddataavailable
func (a NSApplication) IsProtectedDataAvailable() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("protectedDataAvailable"))
	return rv
}
func (a NSApplication) SetIsProtectedDataAvailable(value bool) {
	objc.Send[struct{}](a.ID, objc.Sel("setProtectedDataAvailable:"), value)
}
// The window that currently receives keyboard events.
//
// See: https://developer.apple.com/documentation/appkit/nsapplication/keywindow
func (a NSApplication) KeyWindow() INSWindow {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("keyWindow"))
	return NSWindowFromID(objc.ID(rv))
}
func (a NSApplication) SetKeyWindow(value INSWindow) {
	objc.Send[struct{}](a.ID, objc.Sel("setKeyWindow:"), value)
}
// The app’s main menu bar.
//
// See: https://developer.apple.com/documentation/appkit/nsapplication/mainmenu
func (a NSApplication) MainMenu() INSMenu {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("mainMenu"))
	return NSMenuFromID(objc.ID(rv))
}
func (a NSApplication) SetMainMenu(value INSMenu) {
	objc.Send[struct{}](a.ID, objc.Sel("setMainMenu:"), value)
}
// The app’s main window.
//
// See: https://developer.apple.com/documentation/appkit/nsapplication/mainwindow
func (a NSApplication) MainWindow() INSWindow {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("mainWindow"))
	return NSWindowFromID(objc.ID(rv))
}
func (a NSApplication) SetMainWindow(value INSWindow) {
	objc.Send[struct{}](a.ID, objc.Sel("setMainWindow:"), value)
}
// The modal window displayed by the app.
//
// See: https://developer.apple.com/documentation/appkit/nsapplication/modalwindow
func (a NSApplication) ModalWindow() INSWindow {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("modalWindow"))
	return NSWindowFromID(objc.ID(rv))
}
func (a NSApplication) SetModalWindow(value INSWindow) {
	objc.Send[struct{}](a.ID, objc.Sel("setModalWindow:"), value)
}
// The occlusion state of the app.
//
// See: https://developer.apple.com/documentation/appkit/nsapplication/occlusionstate-swift.property
func (a NSApplication) OcclusionState() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("occlusionState"))
	return objectivec.Object{ID: rv}
}
func (a NSApplication) SetOcclusionState(value objectivec.IObject) {
	objc.Send[struct{}](a.ID, objc.Sel("setOcclusionState:"), value)
}
// The app’s Services menu.
//
// See: https://developer.apple.com/documentation/appkit/nsapplication/servicesmenu
func (a NSApplication) ServicesMenu() INSMenu {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("servicesMenu"))
	return NSMenuFromID(objc.ID(rv))
}
func (a NSApplication) SetServicesMenu(value INSMenu) {
	objc.Send[struct{}](a.ID, objc.Sel("setServicesMenu:"), value)
}
// An array of the app’s window objects.
//
// See: https://developer.apple.com/documentation/appkit/nsapplication/windows
func (a NSApplication) Windows() INSWindow {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("windows"))
	return NSWindowFromID(objc.ID(rv))
}
func (a NSApplication) SetWindows(value INSWindow) {
	objc.Send[struct{}](a.ID, objc.Sel("setWindows:"), value)
}
// The Window menu of the app.
//
// See: https://developer.apple.com/documentation/appkit/nsapplication/windowsmenu
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
// [NSApp]: https://developer.apple.com/documentation/AppKit/NSApp
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
func (_NSApplicationClass NSApplicationClass) SetDisplayWindowRunLoopOrdering(value int) {
	objc.Send[struct{}](objc.ID(_NSApplicationClass.class), objc.Sel("setNSDisplayWindowRunLoopOrdering:"), value)
}
// The priority at which cursor rects are reset.
//
// See: https://developer.apple.com/documentation/appkit/nsapplication/resetcursorrectsrunloopordering
func (_NSApplicationClass NSApplicationClass) ResetCursorRectsRunLoopOrdering() int {
	rv := objc.Send[int](objc.ID(_NSApplicationClass.class), objc.Sel("NSResetCursorRectsRunLoopOrdering"))
	return rv
}
func (_NSApplicationClass NSApplicationClass) SetResetCursorRectsRunLoopOrdering(value int) {
	objc.Send[struct{}](objc.ID(_NSApplicationClass.class), objc.Sel("setNSResetCursorRectsRunLoopOrdering:"), value)
}
// Run-loop message priority for handling window updates.
//
// See: https://developer.apple.com/documentation/appkit/nsapplication/updatewindowsrunloopordering
func (_NSApplicationClass NSApplicationClass) UpdateWindowsRunLoopOrdering() int {
	rv := objc.Send[int](objc.ID(_NSApplicationClass.class), objc.Sel("NSUpdateWindowsRunLoopOrdering"))
	return rv
}
func (_NSApplicationClass NSApplicationClass) SetUpdateWindowsRunLoopOrdering(value int) {
	objc.Send[struct{}](objc.ID(_NSApplicationClass.class), objc.Sel("setNSUpdateWindowsRunLoopOrdering:"), value)
}

			// Protocol methods for NSAccessibilityElementProtocol
			
// Returns the accessibility element’s frame in screen coordinates.
//
// # Return Value
// 
// The element’s frame in screen coordinates.
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityFrame] property. This method is called whenever accessibility
// clients request the [size] or [position] attributes.
//
// [accessibilityFrame]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFrame
// [position]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/position
// [size]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/size
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityFrame()
func (o NSApplication) AccessibilityFrame() corefoundation.CGRect {
	
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("accessibilityFrame"))
	return rv
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
// [accessibilityParent] property.
//
// [accessibilityParent]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityParent
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityParent()
func (o NSApplication) AccessibilityParent() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityParent"))
	return objectivec.Object{ID: rv}
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
// [accessibilityIdentifier] property.
//
// [accessibilityIdentifier]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityIdentifier
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityIdentifier()
func (o NSApplication) AccessibilityIdentifier() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityIdentifier"))
	return foundation.NSStringFromID(rv).String()
	}
// Returns a Boolean value that indicates whether the accessibility element
// has the keyboard focus.
//
// # Return Value
// 
// [true] if this element has the keyboard focus; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityFocused] property.
//
// [accessibilityFocused]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFocused
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/isAccessibilityFocused()
func (o NSApplication) IsAccessibilityFocused() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFocused"))
	return rv
	}

			// Protocol methods for NSAccessibilityProtocol
			
// Returns a Boolean value that determines whether the accessibility element
// participates in the accessibility hierarchy.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityelement()
func (o NSApplication) IsAccessibilityElement() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityElement"))
	return rv
	}
// Sets a Boolean value that determines whether the accessibility element
// participates in the accessibility hierarchy.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityelement(_:)
func (o NSApplication) SetAccessibilityElement(accessibilityElement bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityElement:"), accessibilityElement)
	}
// Returns a Boolean value that determines whether the accessibility element
// responds to user events.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityenabled()
func (o NSApplication) IsAccessibilityEnabled() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityEnabled"))
	return rv
	}
// Sets a Boolean value that determines whether the accessibility element
// responds to user events.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityenabled(_:)
func (o NSApplication) SetAccessibilityEnabled(accessibilityEnabled bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityEnabled:"), accessibilityEnabled)
	}
// Sets the accessibility element’s frame in screen coordinates.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityframe(_:)
func (o NSApplication) SetAccessibilityFrame(accessibilityFrame corefoundation.CGRect) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFrame:"), accessibilityFrame)
	}
// Returns the help text for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityhelp()
func (o NSApplication) AccessibilityHelp() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHelp"))
	return foundation.NSStringFromID(rv).String()
	}
// Sets the help text for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityhelp(_:)
func (o NSApplication) SetAccessibilityHelp(accessibilityHelp string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHelp:"), objc.String(accessibilityHelp))
	}
// Returns a short description of the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitylabel()
func (o NSApplication) AccessibilityLabel() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityLabel"))
	return foundation.NSStringFromID(rv).String()
	}
// Sets a short description of the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitylabel(_:)
func (o NSApplication) SetAccessibilityLabel(accessibilityLabel string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityLabel:"), objc.String(accessibilityLabel))
	}
// Returns the title of the accessibility element—for example, a button’s
// visible text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitytitle()
func (o NSApplication) AccessibilityTitle() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTitle"))
	return foundation.NSStringFromID(rv).String()
	}
// Sets the title of the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitytitle(_:)
func (o NSApplication) SetAccessibilityTitle(accessibilityTitle string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityTitle:"), objc.String(accessibilityTitle))
	}
// Returns the accessibility element’s value.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityvalue()
func (o NSApplication) AccessibilityValue() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityValue"))
	return objectivec.Object{ID: rv}
	}
// Sets the accessibility element’s value.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityvalue(_:)
func (o NSApplication) SetAccessibilityValue(accessibilityValue objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityValue:"), accessibilityValue)
	}
// Returns a Boolean value that indicates whether assistive apps can invoke
// the specified selector on the accessibility element.
//
// selector: The selector to check.
//
// # Return Value
// 
// [true], if accessibility clients can call the selector; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilitySelectorAllowed(_:)
func (o NSApplication) IsAccessibilitySelectorAllowed(selector objc.SEL) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilitySelectorAllowed:"), selector)
	return rv
	}
// Returns the contents of the current accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycontents()
func (o NSApplication) AccessibilityContents() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityContents"))
	return objectivec.Object{ID: rv}
	}
// Sets the contents of the current accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycontents(_:)
func (o NSApplication) SetAccessibilityContents(accessibilityContents foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityContents:"), accessibilityContents)
	}
// Returns the critical value for the level indicator.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycriticalvalue()
func (o NSApplication) AccessibilityCriticalValue() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCriticalValue"))
	return objectivec.Object{ID: rv}
	}
// Sets the critical value for the level indicator.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycriticalvalue(_:)
func (o NSApplication) SetAccessibilityCriticalValue(accessibilityCriticalValue objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCriticalValue:"), accessibilityCriticalValue)
	}
// Sets the accessibility element’s identity.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityidentifier(_:)
func (o NSApplication) SetAccessibilityIdentifier(accessibilityIdentifier string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityIdentifier:"), objc.String(accessibilityIdentifier))
	}
// Returns the maximum value for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitymaxvalue()
func (o NSApplication) AccessibilityMaxValue() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMaxValue"))
	return objectivec.Object{ID: rv}
	}
// Sets the maximum value for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymaxvalue(_:)
func (o NSApplication) SetAccessibilityMaxValue(accessibilityMaxValue objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMaxValue:"), accessibilityMaxValue)
	}
// Returns the minimum value for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityminvalue()
func (o NSApplication) AccessibilityMinValue() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMinValue"))
	return objectivec.Object{ID: rv}
	}
// Sets the minimum value for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityminvalue(_:)
func (o NSApplication) SetAccessibilityMinValue(accessibilityMinValue objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMinValue:"), accessibilityMinValue)
	}
// Returns the orientation of the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityorientation()
func (o NSApplication) AccessibilityOrientation() NSAccessibilityOrientation {
	
	rv := objc.Send[NSAccessibilityOrientation](o.ID, objc.Sel("accessibilityOrientation"))
	return rv
	}
// Sets the orientation of the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityorientation(_:)
func (o NSApplication) SetAccessibilityOrientation(accessibilityOrientation NSAccessibilityOrientation) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityOrientation:"), accessibilityOrientation)
	}
// Returns a Boolean value that determines whether the accessibility element
// contains protected content.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityprotectedcontent()
func (o NSApplication) IsAccessibilityProtectedContent() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityProtectedContent"))
	return rv
	}
// Sets a Boolean value that determines whether the accessibility element
// contains protected content.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityprotectedcontent(_:)
func (o NSApplication) SetAccessibilityProtectedContent(accessibilityProtectedContent bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityProtectedContent:"), accessibilityProtectedContent)
	}
// Returns a Boolean value that determines whether the accessibility element
// is currently in a selected state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityselected()
func (o NSApplication) IsAccessibilitySelected() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilitySelected"))
	return rv
	}
// Sets a Boolean value that determines whether the accessibility element is
// currently in a selected state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselected(_:)
func (o NSApplication) SetAccessibilitySelected(accessibilitySelected bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelected:"), accessibilitySelected)
	}
// Returns the URL for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityurl()
func (o NSApplication) AccessibilityURL() foundation.INSURL {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityURL"))
	return foundation.NSURLFromID(rv)
	}
// Sets the URL for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityurl(_:)
func (o NSApplication) SetAccessibilityURL(accessibilityURL foundation.INSURL) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityURL:"), accessibilityURL)
	}
// Returns the human-readable description of the accessibility element’s
// value.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityvaluedescription()
func (o NSApplication) AccessibilityValueDescription() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityValueDescription"))
	return foundation.NSStringFromID(rv).String()
	}
// Sets the human-readable description of the accessibility element’s value.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityvaluedescription(_:)
func (o NSApplication) SetAccessibilityValueDescription(accessibilityValueDescription string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityValueDescription:"), objc.String(accessibilityValueDescription))
	}
// Returns the warning value for the level indicator.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitywarningvalue()
func (o NSApplication) AccessibilityWarningValue() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityWarningValue"))
	return objectivec.Object{ID: rv}
	}
// Sets the warning value for the level indicator.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitywarningvalue(_:)
func (o NSApplication) SetAccessibilityWarningValue(accessibilityWarningValue objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityWarningValue:"), accessibilityWarningValue)
	}
// Returns the child accessibility elements in the accessibility hierarchy.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitychildren()
func (o NSApplication) AccessibilityChildren() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityChildren"))
	return objectivec.Object{ID: rv}
	}
// Sets the child accessibility elements in the accessibility hierarchy.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitychildren(_:)
func (o NSApplication) SetAccessibilityChildren(accessibilityChildren foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityChildren:"), accessibilityChildren)
	}
// Returns the array of child accessibility elements in order for linear
// navigation.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitychildreninnavigationorder()
func (o NSApplication) AccessibilityChildrenInNavigationOrder() unsafe.Pointer {
	
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("accessibilityChildrenInNavigationOrder"))
	return rv
	}
// Sets the array of child accessibility elements in order for linear
// navigation.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitychildreninnavigationorder(_:)
func (o NSApplication) SetAccessibilityChildrenInNavigationOrder(accessibilityChildrenInNavigationOrder foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityChildrenInNavigationOrder:"), accessibilityChildrenInNavigationOrder)
	}
// Sets the accessibility element’s parent in the accessibility hierarchy.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityparent(_:)
func (o NSApplication) SetAccessibilityParent(accessibilityParent objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityParent:"), accessibilityParent)
	}
// Returns the accessibility element’s currently selected children.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityselectedchildren()
func (o NSApplication) AccessibilitySelectedChildren() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedChildren"))
	return objectivec.Object{ID: rv}
	}
// Sets the accessibility element’s currently selected children.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselectedchildren(_:)
func (o NSApplication) SetAccessibilitySelectedChildren(accessibilitySelectedChildren foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedChildren:"), accessibilitySelectedChildren)
	}
// Returns the top-level element that contains the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitytopleveluielement()
func (o NSApplication) AccessibilityTopLevelUIElement() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTopLevelUIElement"))
	return objectivec.Object{ID: rv}
	}
// Sets the top-level element that contains the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitytopleveluielement(_:)
func (o NSApplication) SetAccessibilityTopLevelUIElement(accessibilityTopLevelUIElement objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityTopLevelUIElement:"), accessibilityTopLevelUIElement)
	}
// Returns the accessibility element’s visible child accessibility elements.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityvisiblechildren()
func (o NSApplication) AccessibilityVisibleChildren() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleChildren"))
	return objectivec.Object{ID: rv}
	}
// Sets the accessibility element’s visible child accessibility elements.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityvisiblechildren(_:)
func (o NSApplication) SetAccessibilityVisibleChildren(accessibilityVisibleChildren foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleChildren:"), accessibilityVisibleChildren)
	}
// Returns the child accessibility element with the current focus.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityapplicationfocuseduielement()
func (o NSApplication) AccessibilityApplicationFocusedUIElement() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityApplicationFocusedUIElement"))
	return objectivec.Object{ID: rv}
	}
// Sets the child accessibility element with the current focus.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityapplicationfocuseduielement(_:)
func (o NSApplication) SetAccessibilityApplicationFocusedUIElement(accessibilityApplicationFocusedUIElement objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityApplicationFocusedUIElement:"), accessibilityApplicationFocusedUIElement)
	}
// Sets a Boolean value that determines whether the accessibility element has
// the keyboard focus.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityfocused(_:)
func (o NSApplication) SetAccessibilityFocused(accessibilityFocused bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFocused:"), accessibilityFocused)
	}
// Returns the child window with the current focus.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityfocusedwindow()
func (o NSApplication) AccessibilityFocusedWindow() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityFocusedWindow"))
	return objectivec.Object{ID: rv}
	}
// Sets the child window with the current focus.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityfocusedwindow(_:)
func (o NSApplication) SetAccessibilityFocusedWindow(accessibilityFocusedWindow objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFocusedWindow:"), accessibilityFocusedWindow)
	}
// Returns the array of elements that shares the keyboard focus with the
// accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysharedfocuselements()
func (o NSApplication) AccessibilitySharedFocusElements() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySharedFocusElements"))
	return objectivec.Object{ID: rv}
	}
// Sets the array of elements that shares the keyboard focus with the
// accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysharedfocuselements(_:)
func (o NSApplication) SetAccessibilitySharedFocusElements(accessibilitySharedFocusElements foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySharedFocusElements:"), accessibilitySharedFocusElements)
	}
// Returns a Boolean value that determines whether the accessibility element
// must have content for successful submission of a form.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityrequired()
func (o NSApplication) IsAccessibilityRequired() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityRequired"))
	return rv
	}
// Sets a Boolean value that determines whether the accessibility element must
// have content for successful submission of a form.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityrequired(_:)
func (o NSApplication) SetAccessibilityRequired(accessibilityRequired bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRequired:"), accessibilityRequired)
	}
// Returns the type of interface element that the accessibility element
// represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityrole()
func (o NSApplication) AccessibilityRole() NSAccessibilityRole {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRole"))
	return NSAccessibilityRole(foundation.NSStringFromID(rv).String())
	}
// Sets the type of interface element that the accessibility element
// represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityrole(_:)
func (o NSApplication) SetAccessibilityRole(accessibilityRole NSAccessibilityRole) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRole:"), objc.String(string(accessibilityRole)))
	}
// Returns a localized, human-intelligible description of the accessibility
// element’s role, such as
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityroledescription()
func (o NSApplication) AccessibilityRoleDescription() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRoleDescription"))
	return foundation.NSStringFromID(rv).String()
	}
// Sets the localized, human-intelligible description of the accessibility
// element’s role, such as
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityroledescription(_:)
func (o NSApplication) SetAccessibilityRoleDescription(accessibilityRoleDescription string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRoleDescription:"), objc.String(accessibilityRoleDescription))
	}
// Returns the specialized interface element type that the accessibility
// element represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysubrole()
func (o NSApplication) AccessibilitySubrole() NSAccessibilitySubrole {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySubrole"))
	return NSAccessibilitySubrole(foundation.NSStringFromID(rv).String())
	}
// Sets the specialized interface element type that the accessibility element
// represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysubrole(_:)
func (o NSApplication) SetAccessibilitySubrole(accessibilitySubrole NSAccessibilitySubrole) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySubrole:"), objc.String(string(accessibilitySubrole)))
	}
// Returns the custom actions of the current accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycustomactions()
func (o NSApplication) AccessibilityCustomActions() INSAccessibilityCustomAction {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCustomActions"))
	return NSAccessibilityCustomActionFromID(rv)
	}
// Sets the custom actions of the current accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycustomactions(_:)
func (o NSApplication) SetAccessibilityCustomActions(accessibilityCustomActions foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCustomActions:"), accessibilityCustomActions)
	}
// Returns the custom rotors of the current accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycustomrotors()
func (o NSApplication) AccessibilityCustomRotors() INSAccessibilityCustomRotor {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCustomRotors"))
	return NSAccessibilityCustomRotorFromID(rv)
	}
// Sets the custom rotors of the current accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycustomrotors(_:)
func (o NSApplication) SetAccessibilityCustomRotors(accessibilityCustomRotors foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCustomRotors:"), accessibilityCustomRotors)
	}
// Returns the line number that contains the insertion point.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityinsertionpointlinenumber()
func (o NSApplication) AccessibilityInsertionPointLineNumber() int {
	
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityInsertionPointLineNumber"))
	return rv
	}
// Sets the line number that contains the insertion point.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityinsertionpointlinenumber(_:)
func (o NSApplication) SetAccessibilityInsertionPointLineNumber(accessibilityInsertionPointLineNumber int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityInsertionPointLineNumber:"), accessibilityInsertionPointLineNumber)
	}
// Returns the number of characters in the text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitynumberofcharacters()
func (o NSApplication) AccessibilityNumberOfCharacters() int {
	
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityNumberOfCharacters"))
	return rv
	}
// Sets the number of characters in the text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitynumberofcharacters(_:)
func (o NSApplication) SetAccessibilityNumberOfCharacters(accessibilityNumberOfCharacters int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityNumberOfCharacters:"), accessibilityNumberOfCharacters)
	}
// Returns the placeholder value for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityplaceholdervalue()
func (o NSApplication) AccessibilityPlaceholderValue() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityPlaceholderValue"))
	return foundation.NSStringFromID(rv).String()
	}
// Sets the placeholder value for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityplaceholdervalue(_:)
func (o NSApplication) SetAccessibilityPlaceholderValue(accessibilityPlaceholderValue string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityPlaceholderValue:"), objc.String(accessibilityPlaceholderValue))
	}
// Returns the currently selected text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityselectedtext()
func (o NSApplication) AccessibilitySelectedText() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedText"))
	return foundation.NSStringFromID(rv).String()
	}
// Sets the currently selected text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselectedtext(_:)
func (o NSApplication) SetAccessibilitySelectedText(accessibilitySelectedText string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedText:"), objc.String(accessibilitySelectedText))
	}
// Returns the range of the currently selected text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityselectedtextrange()
func (o NSApplication) AccessibilitySelectedTextRange() foundation.NSRange {
	
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilitySelectedTextRange"))
	return rv
	}
// Sets the range of the currently selected text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselectedtextrange(_:)
func (o NSApplication) SetAccessibilitySelectedTextRange(accessibilitySelectedTextRange foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedTextRange:"), accessibilitySelectedTextRange)
	}
// Returns an array of ranges for the currently selected text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityselectedtextranges()
func (o NSApplication) AccessibilitySelectedTextRanges() foundation.NSValue {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedTextRanges"))
	return foundation.NSValueFromID(rv)
	}
// Sets an array of ranges for the currently selected text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselectedtextranges(_:)
func (o NSApplication) SetAccessibilitySelectedTextRanges(accessibilitySelectedTextRanges foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedTextRanges:"), accessibilitySelectedTextRanges)
	}
// Returns the range of characters that the accessibility element displays.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysharedcharacterrange()
func (o NSApplication) AccessibilitySharedCharacterRange() foundation.NSRange {
	
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilitySharedCharacterRange"))
	return rv
	}
// Sets the range of characters that the accessibility element displays.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysharedcharacterrange(_:)
func (o NSApplication) SetAccessibilitySharedCharacterRange(accessibilitySharedCharacterRange foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySharedCharacterRange:"), accessibilitySharedCharacterRange)
	}
// Returns the other elements that share text with the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysharedtextuielements()
func (o NSApplication) AccessibilitySharedTextUIElements() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySharedTextUIElements"))
	return objectivec.Object{ID: rv}
	}
// Sets the other elements that share text with the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysharedtextuielements(_:)
func (o NSApplication) SetAccessibilitySharedTextUIElements(accessibilitySharedTextUIElements foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySharedTextUIElements:"), accessibilitySharedTextUIElements)
	}
// Returns the range of visible characters in the document.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityvisiblecharacterrange()
func (o NSApplication) AccessibilityVisibleCharacterRange() foundation.NSRange {
	
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityVisibleCharacterRange"))
	return rv
	}
// Sets the range of visible characters in the document.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityvisiblecharacterrange(_:)
func (o NSApplication) SetAccessibilityVisibleCharacterRange(accessibilityVisibleCharacterRange foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleCharacterRange:"), accessibilityVisibleCharacterRange)
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
func (o NSApplication) AccessibilityStringForRange(range_ foundation.NSRange) string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityStringForRange:"), range_)
	return foundation.NSStringFromID(rv).String()
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
func (o NSApplication) AccessibilityAttributedStringForRange(range_ foundation.NSRange) foundation.NSAttributedString {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityAttributedStringForRange:"), range_)
	return foundation.NSAttributedStringFromID(rv)
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
func (o NSApplication) AccessibilityRTFForRange(range_ foundation.NSRange) foundation.INSData {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRTFForRange:"), range_)
	return foundation.NSDataFromID(rv)
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
func (o NSApplication) AccessibilityFrameForRange(range_ foundation.NSRange) corefoundation.CGRect {
	
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("accessibilityFrameForRange:"), range_)
	return rv
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
func (o NSApplication) AccessibilityLineForIndex(index int) int {
	
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityLineForIndex:"), index)
	return rv
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
func (o NSApplication) AccessibilityRangeForIndex(index int) foundation.NSRange {
	
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityRangeForIndex:"), index)
	return rv
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
func (o NSApplication) AccessibilityStyleRangeForIndex(index int) foundation.NSRange {
	
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityStyleRangeForIndex:"), index)
	return rv
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
func (o NSApplication) AccessibilityRangeForLine(line int) foundation.NSRange {
	
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityRangeForLine:"), line)
	return rv
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
func (o NSApplication) AccessibilityRangeForPosition(point corefoundation.CGPoint) foundation.NSRange {
	
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityRangeForPosition:"), point)
	return rv
	}
// Returns the activation point for the user interface element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityactivationpoint()
func (o NSApplication) AccessibilityActivationPoint() corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("accessibilityActivationPoint"))
	return rv
	}
// Sets the activation point for the user interface element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityactivationpoint(_:)
func (o NSApplication) SetAccessibilityActivationPoint(accessibilityActivationPoint corefoundation.CGPoint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityActivationPoint:"), accessibilityActivationPoint)
	}
// Returns the Boolean value that determines whether the accessibility
// element’s alternative UI is currently visible.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityalternateuivisible()
func (o NSApplication) IsAccessibilityAlternateUIVisible() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityAlternateUIVisible"))
	return rv
	}
// Sets the Boolean value that determines whether the accessibility
// element’s alternative UI is currently visible.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityalternateuivisible(_:)
func (o NSApplication) SetAccessibilityAlternateUIVisible(accessibilityAlternateUIVisible bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityAlternateUIVisible:"), accessibilityAlternateUIVisible)
	}
// Returns the child accessibility element that represents the window’s
// cancel button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycancelbutton()
func (o NSApplication) AccessibilityCancelButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCancelButton"))
	return objectivec.Object{ID: rv}
	}
// Sets the child accessibility element that represents the window’s cancel
// button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycancelbutton(_:)
func (o NSApplication) SetAccessibilityCancelButton(accessibilityCancelButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCancelButton:"), accessibilityCancelButton)
	}
// Returns the child accessibility element that represents the window’s
// close button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityclosebutton()
func (o NSApplication) AccessibilityCloseButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCloseButton"))
	return objectivec.Object{ID: rv}
	}
// Sets the child accessibility element that represents the window’s close
// button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityclosebutton(_:)
func (o NSApplication) SetAccessibilityCloseButton(accessibilityCloseButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCloseButton:"), accessibilityCloseButton)
	}
// Returns the child accessibility element that represents the window’s
// default button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitydefaultbutton()
func (o NSApplication) AccessibilityDefaultButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDefaultButton"))
	return objectivec.Object{ID: rv}
	}
// Sets the child accessibility element that represents the window’s default
// button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitydefaultbutton(_:)
func (o NSApplication) SetAccessibilityDefaultButton(accessibilityDefaultButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDefaultButton:"), accessibilityDefaultButton)
	}
// Returns the child accessibility element that represents the window’s
// full-screen button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityfullscreenbutton()
func (o NSApplication) AccessibilityFullScreenButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityFullScreenButton"))
	return objectivec.Object{ID: rv}
	}
// Sets the child accessibility element that represents the window’s
// full-screen button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityfullscreenbutton(_:)
func (o NSApplication) SetAccessibilityFullScreenButton(accessibilityFullScreenButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFullScreenButton:"), accessibilityFullScreenButton)
	}
// Returns the child accessibility element that represents the window’s grow
// area.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitygrowarea()
func (o NSApplication) AccessibilityGrowArea() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityGrowArea"))
	return objectivec.Object{ID: rv}
	}
// Sets the child accessibility element that represents the window’s grow
// area.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitygrowarea(_:)
func (o NSApplication) SetAccessibilityGrowArea(accessibilityGrowArea objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityGrowArea:"), accessibilityGrowArea)
	}
// Returns a Boolean value that determines whether the window is the app’s
// main window.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilitymain()
func (o NSApplication) IsAccessibilityMain() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityMain"))
	return rv
	}
// Sets a Boolean value that determines whether the window is the app’s main
// window.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymain(_:)
func (o NSApplication) SetAccessibilityMain(accessibilityMain bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMain:"), accessibilityMain)
	}
// Returns the child accessibility element that represents the window’s
// minimize button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityminimizebutton()
func (o NSApplication) AccessibilityMinimizeButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMinimizeButton"))
	return objectivec.Object{ID: rv}
	}
// Sets the child accessibility element that represents the window’s
// minimize button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityminimizebutton(_:)
func (o NSApplication) SetAccessibilityMinimizeButton(accessibilityMinimizeButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMinimizeButton:"), accessibilityMinimizeButton)
	}
// Returns the Boolean value that determines whether the window is in a
// minimized state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityminimized()
func (o NSApplication) IsAccessibilityMinimized() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityMinimized"))
	return rv
	}
// Sets the Boolean value that determines whether the window is in a minimized
// state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityminimized(_:)
func (o NSApplication) SetAccessibilityMinimized(accessibilityMinimized bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMinimized:"), accessibilityMinimized)
	}
// Returns a Boolean value that determines whether the window is modal.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilitymodal()
func (o NSApplication) IsAccessibilityModal() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityModal"))
	return rv
	}
// Sets a Boolean value that determines whether the window is modal.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymodal(_:)
func (o NSApplication) SetAccessibilityModal(accessibilityModal bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityModal:"), accessibilityModal)
	}
// Returns the child accessibility element that represents the window’s
// proxy icon.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityproxy()
func (o NSApplication) AccessibilityProxy() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityProxy"))
	return objectivec.Object{ID: rv}
	}
// Sets the child accessibility element that represents the window’s proxy
// icon.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityproxy(_:)
func (o NSApplication) SetAccessibilityProxy(accessibilityProxy objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityProxy:"), accessibilityProxy)
	}
// Returns the menu currently displaying for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityshownmenu()
func (o NSApplication) AccessibilityShownMenu() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityShownMenu"))
	return objectivec.Object{ID: rv}
	}
// Sets the menu currently displaying for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityshownmenu(_:)
func (o NSApplication) SetAccessibilityShownMenu(accessibilityShownMenu objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityShownMenu:"), accessibilityShownMenu)
	}
// Returns the child accessibility element that represents the window’s
// toolbar button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitytoolbarbutton()
func (o NSApplication) AccessibilityToolbarButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityToolbarButton"))
	return objectivec.Object{ID: rv}
	}
// Sets the child accessibility element that represents the window’s toolbar
// button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitytoolbarbutton(_:)
func (o NSApplication) SetAccessibilityToolbarButton(accessibilityToolbarButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityToolbarButton:"), accessibilityToolbarButton)
	}
// Returns the window that contains the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitywindow()
func (o NSApplication) AccessibilityWindow() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityWindow"))
	return objectivec.Object{ID: rv}
	}
// Sets the window that contains the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitywindow(_:)
func (o NSApplication) SetAccessibilityWindow(accessibilityWindow objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityWindow:"), accessibilityWindow)
	}
// Returns the child accessibility element that represents the window’s zoom
// button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityzoombutton()
func (o NSApplication) AccessibilityZoomButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityZoomButton"))
	return objectivec.Object{ID: rv}
	}
// Sets the child accessibility element that represents the window’s zoom
// button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityzoombutton(_:)
func (o NSApplication) SetAccessibilityZoomButton(accessibilityZoomButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityZoomButton:"), accessibilityZoomButton)
	}
// Returns the icon for the app’s menu bar extra.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityextrasmenubar()
func (o NSApplication) AccessibilityExtrasMenuBar() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityExtrasMenuBar"))
	return objectivec.Object{ID: rv}
	}
// Sets the icon for the app’s menu bar extra.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityextrasmenubar(_:)
func (o NSApplication) SetAccessibilityExtrasMenuBar(accessibilityExtrasMenuBar objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityExtrasMenuBar:"), accessibilityExtrasMenuBar)
	}
// Returns a Boolean value that determines whether the app is the frontmost
// app.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityfrontmost()
func (o NSApplication) IsAccessibilityFrontmost() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFrontmost"))
	return rv
	}
// Sets a Boolean value that determines whether the app is the frontmost app.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityfrontmost(_:)
func (o NSApplication) SetAccessibilityFrontmost(accessibilityFrontmost bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFrontmost:"), accessibilityFrontmost)
	}
// Returns a Boolean value that determines whether the app is in a hidden
// state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityhidden()
func (o NSApplication) IsAccessibilityHidden() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityHidden"))
	return rv
	}
// Sets a Boolean value that determines whether the app is in a hidden state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityhidden(_:)
func (o NSApplication) SetAccessibilityHidden(accessibilityHidden bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHidden:"), accessibilityHidden)
	}
// Returns the app’s main window.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitymainwindow()
func (o NSApplication) AccessibilityMainWindow() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMainWindow"))
	return objectivec.Object{ID: rv}
	}
// Sets the app’s main window.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymainwindow(_:)
func (o NSApplication) SetAccessibilityMainWindow(accessibilityMainWindow objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMainWindow:"), accessibilityMainWindow)
	}
// Returns the app’s menu bar.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitymenubar()
func (o NSApplication) AccessibilityMenuBar() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMenuBar"))
	return objectivec.Object{ID: rv}
	}
// Sets the app’s menu bar.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymenubar(_:)
func (o NSApplication) SetAccessibilityMenuBar(accessibilityMenuBar objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMenuBar:"), accessibilityMenuBar)
	}
// Returns an array that contains all the app’s windows.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitywindows()
func (o NSApplication) AccessibilityWindows() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityWindows"))
	return objectivec.Object{ID: rv}
	}
// Sets the array that contains all the app’s windows.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitywindows(_:)
func (o NSApplication) SetAccessibilityWindows(accessibilityWindows foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityWindows:"), accessibilityWindows)
	}
// Returns the number of columns in the accessibility element’s grid.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycolumncount()
func (o NSApplication) AccessibilityColumnCount() int {
	
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityColumnCount"))
	return rv
	}
// Sets the number of columns in the accessibility element’s grid.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycolumncount(_:)
func (o NSApplication) SetAccessibilityColumnCount(accessibilityColumnCount int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumnCount:"), accessibilityColumnCount)
	}
// Returns a Boolean value that determines whether the accessibility
// element’s grid is in row major order or in column major order.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityorderedbyrow()
func (o NSApplication) IsAccessibilityOrderedByRow() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityOrderedByRow"))
	return rv
	}
// Sets a Boolean value that determines whether the element’s grid is in row
// major order or in column major order.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityorderedbyrow(_:)
func (o NSApplication) SetAccessibilityOrderedByRow(accessibilityOrderedByRow bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityOrderedByRow:"), accessibilityOrderedByRow)
	}
// Returns the number of rows in the accessibility element’s grid.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityrowcount()
func (o NSApplication) AccessibilityRowCount() int {
	
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityRowCount"))
	return rv
	}
// Sets the number of rows in the accessibility element’s grid.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityrowcount(_:)
func (o NSApplication) SetAccessibilityRowCount(accessibilityRowCount int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRowCount:"), accessibilityRowCount)
	}
// Returns the horizontal scroll bar for the scroll view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityhorizontalscrollbar()
func (o NSApplication) AccessibilityHorizontalScrollBar() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHorizontalScrollBar"))
	return objectivec.Object{ID: rv}
	}
// Sets the horizontal scroll bar for the scroll view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityhorizontalscrollbar(_:)
func (o NSApplication) SetAccessibilityHorizontalScrollBar(accessibilityHorizontalScrollBar objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHorizontalScrollBar:"), accessibilityHorizontalScrollBar)
	}
// Returns the vertical scroll bar for the scroll view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityverticalscrollbar()
func (o NSApplication) AccessibilityVerticalScrollBar() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVerticalScrollBar"))
	return objectivec.Object{ID: rv}
	}
// Sets the vertical scroll bar for the scroll view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityverticalscrollbar(_:)
func (o NSApplication) SetAccessibilityVerticalScrollBar(accessibilityVerticalScrollBar objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVerticalScrollBar:"), accessibilityVerticalScrollBar)
	}
// Returns the column header accessibility elements for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycolumnheaderuielements()
func (o NSApplication) AccessibilityColumnHeaderUIElements() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityColumnHeaderUIElements"))
	return objectivec.Object{ID: rv}
	}
// Sets the column header accessibility elements for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycolumnheaderuielements(_:)
func (o NSApplication) SetAccessibilityColumnHeaderUIElements(accessibilityColumnHeaderUIElements foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumnHeaderUIElements:"), accessibilityColumnHeaderUIElements)
	}
// Returns the column accessibility elements for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycolumns()
func (o NSApplication) AccessibilityColumns() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityColumns"))
	return objectivec.Object{ID: rv}
	}
// Sets the column accessibility elements for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycolumns(_:)
func (o NSApplication) SetAccessibilityColumns(accessibilityColumns foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumns:"), accessibilityColumns)
	}
// Returns the column titles for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycolumntitles()
func (o NSApplication) AccessibilityColumnTitles() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityColumnTitles"))
	return objectivec.Object{ID: rv}
	}
// Sets the column titles for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycolumntitles(_:)
func (o NSApplication) SetAccessibilityColumnTitles(accessibilityColumnTitles foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumnTitles:"), accessibilityColumnTitles)
	}
// Returns a Boolean value that determines whether the accessibility element
// is in an expanded state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityexpanded()
func (o NSApplication) IsAccessibilityExpanded() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityExpanded"))
	return rv
	}
// Sets a Boolean value that determines whether accessibility element is in an
// expanded state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityexpanded(_:)
func (o NSApplication) SetAccessibilityExpanded(accessibilityExpanded bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityExpanded:"), accessibilityExpanded)
	}
// Returns the header for the table view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityheader()
func (o NSApplication) AccessibilityHeader() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHeader"))
	return objectivec.Object{ID: rv}
	}
// Sets the header for the table view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityheader(_:)
func (o NSApplication) SetAccessibilityHeader(accessibilityHeader objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHeader:"), accessibilityHeader)
	}
// Returns the index of the row or column that the accessibility element
// represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityindex()
func (o NSApplication) AccessibilityIndex() int {
	
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityIndex"))
	return rv
	}
// Sets the index of the row or column that the accessibility element
// represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityindex(_:)
func (o NSApplication) SetAccessibilityIndex(accessibilityIndex int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityIndex:"), accessibilityIndex)
	}
// Returns the row header accessibility elements for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityrowheaderuielements()
func (o NSApplication) AccessibilityRowHeaderUIElements() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRowHeaderUIElements"))
	return objectivec.Object{ID: rv}
	}
// Sets the row header accessibility elements for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityrowheaderuielements(_:)
func (o NSApplication) SetAccessibilityRowHeaderUIElements(accessibilityRowHeaderUIElements foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRowHeaderUIElements:"), accessibilityRowHeaderUIElements)
	}
// Returns the row accessibility elements for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityrows()
func (o NSApplication) AccessibilityRows() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRows"))
	return objectivec.Object{ID: rv}
	}
// Sets the row accessibility elements for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityrows(_:)
func (o NSApplication) SetAccessibilityRows(accessibilityRows foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRows:"), accessibilityRows)
	}
// Returns the currently selected columns for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityselectedcolumns()
func (o NSApplication) AccessibilitySelectedColumns() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedColumns"))
	return objectivec.Object{ID: rv}
	}
// Sets the currently selected columns for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselectedcolumns(_:)
func (o NSApplication) SetAccessibilitySelectedColumns(accessibilitySelectedColumns foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedColumns:"), accessibilitySelectedColumns)
	}
// Returns the currently selected rows for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityselectedrows()
func (o NSApplication) AccessibilitySelectedRows() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedRows"))
	return objectivec.Object{ID: rv}
	}
// Sets the currently selected rows for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselectedrows(_:)
func (o NSApplication) SetAccessibilitySelectedRows(accessibilitySelectedRows foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedRows:"), accessibilitySelectedRows)
	}
// Returns the accessibility element’s sort direction.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysortdirection()
func (o NSApplication) AccessibilitySortDirection() NSAccessibilitySortDirection {
	
	rv := objc.Send[NSAccessibilitySortDirection](o.ID, objc.Sel("accessibilitySortDirection"))
	return rv
	}
// Sets the accessibility element’s sort direction.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysortdirection(_:)
func (o NSApplication) SetAccessibilitySortDirection(accessibilitySortDirection NSAccessibilitySortDirection) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySortDirection:"), accessibilitySortDirection)
	}
// Returns the visible columns for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityvisiblecolumns()
func (o NSApplication) AccessibilityVisibleColumns() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleColumns"))
	return objectivec.Object{ID: rv}
	}
// Sets the visible columns for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityvisiblecolumns(_:)
func (o NSApplication) SetAccessibilityVisibleColumns(accessibilityVisibleColumns foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleColumns:"), accessibilityVisibleColumns)
	}
// Returns the visible rows for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityvisiblerows()
func (o NSApplication) AccessibilityVisibleRows() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleRows"))
	return objectivec.Object{ID: rv}
	}
// Sets the visible rows for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityvisiblerows(_:)
func (o NSApplication) SetAccessibilityVisibleRows(accessibilityVisibleRows foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleRows:"), accessibilityVisibleRows)
	}
// Returns a Boolean value that determines whether the row is disclosing other
// rows.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilitydisclosed()
func (o NSApplication) IsAccessibilityDisclosed() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityDisclosed"))
	return rv
	}
// Sets a Boolean value that determines whether the row is disclosing other
// rows.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitydisclosed(_:)
func (o NSApplication) SetAccessibilityDisclosed(accessibilityDisclosed bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosed:"), accessibilityDisclosed)
	}
// Returns the row disclosing the current row.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitydisclosedbyrow()
func (o NSApplication) AccessibilityDisclosedByRow() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDisclosedByRow"))
	return objectivec.Object{ID: rv}
	}
// Sets the row disclosing the current row.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitydisclosedbyrow(_:)
func (o NSApplication) SetAccessibilityDisclosedByRow(accessibilityDisclosedByRow objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosedByRow:"), accessibilityDisclosedByRow)
	}
// Returns the rows that the current row discloses.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitydisclosedrows()
func (o NSApplication) AccessibilityDisclosedRows() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDisclosedRows"))
	return objectivec.Object{ID: rv}
	}
// Sets the rows that the current row discloses.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitydisclosedrows(_:)
func (o NSApplication) SetAccessibilityDisclosedRows(accessibilityDisclosedRows objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosedRows:"), accessibilityDisclosedRows)
	}
// Returns the indention level for the row.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitydisclosurelevel()
func (o NSApplication) AccessibilityDisclosureLevel() int {
	
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityDisclosureLevel"))
	return rv
	}
// Sets the indention level for the row.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitydisclosurelevel(_:)
func (o NSApplication) SetAccessibilityDisclosureLevel(accessibilityDisclosureLevel int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosureLevel:"), accessibilityDisclosureLevel)
	}
// Returns the column index range of the cell.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycolumnindexrange()
func (o NSApplication) AccessibilityColumnIndexRange() foundation.NSRange {
	
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityColumnIndexRange"))
	return rv
	}
// Sets the column index range of the cell.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycolumnindexrange(_:)
func (o NSApplication) SetAccessibilityColumnIndexRange(accessibilityColumnIndexRange foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumnIndexRange:"), accessibilityColumnIndexRange)
	}
// Returns the row index range of the cell.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityrowindexrange()
func (o NSApplication) AccessibilityRowIndexRange() foundation.NSRange {
	
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityRowIndexRange"))
	return rv
	}
// Sets the row index range of the cell.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityrowindexrange(_:)
func (o NSApplication) SetAccessibilityRowIndexRange(accessibilityRowIndexRange foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRowIndexRange:"), accessibilityRowIndexRange)
	}
// Returns the currently selected cells for the table.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityselectedcells()
func (o NSApplication) AccessibilitySelectedCells() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedCells"))
	return objectivec.Object{ID: rv}
	}
// Sets the currently selected cells for the table.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselectedcells(_:)
func (o NSApplication) SetAccessibilitySelectedCells(accessibilitySelectedCells foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedCells:"), accessibilitySelectedCells)
	}
// Returns the visible cells for the table.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityvisiblecells()
func (o NSApplication) AccessibilityVisibleCells() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleCells"))
	return objectivec.Object{ID: rv}
	}
// Sets the visible cells for the table.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityvisiblecells(_:)
func (o NSApplication) SetAccessibilityVisibleCells(accessibilityVisibleCells foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleCells:"), accessibilityVisibleCells)
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
func (o NSApplication) AccessibilityCellForColumnRow(column int, row int) objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCellForColumn:row:"), column, row)
	return objectivec.Object{ID: rv}
	}
// Returns the drag handle elements for the layout item element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityhandles()
func (o NSApplication) AccessibilityHandles() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHandles"))
	return objectivec.Object{ID: rv}
	}
// Sets the drag handle accessibility elements for the layout item element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityhandles(_:)
func (o NSApplication) SetAccessibilityHandles(accessibilityHandles foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHandles:"), accessibilityHandles)
	}
// Returns the units that the layout area uses for horizontal values.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityhorizontalunits()
func (o NSApplication) AccessibilityHorizontalUnits() NSAccessibilityUnits {
	
	rv := objc.Send[NSAccessibilityUnits](o.ID, objc.Sel("accessibilityHorizontalUnits"))
	return rv
	}
// Sets the units that the layout area uses for horizontal values.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityhorizontalunits(_:)
func (o NSApplication) SetAccessibilityHorizontalUnits(accessibilityHorizontalUnits NSAccessibilityUnits) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHorizontalUnits:"), accessibilityHorizontalUnits)
	}
// Returns the description of the layout area’s horizontal units.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityhorizontalunitdescription()
func (o NSApplication) AccessibilityHorizontalUnitDescription() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHorizontalUnitDescription"))
	return foundation.NSStringFromID(rv).String()
	}
// Sets the description of the layout area’s horizontal units.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityhorizontalunitdescription(_:)
func (o NSApplication) SetAccessibilityHorizontalUnitDescription(accessibilityHorizontalUnitDescription string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHorizontalUnitDescription:"), objc.String(accessibilityHorizontalUnitDescription))
	}
// Returns the units that the layout area uses for vertical values.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityverticalunits()
func (o NSApplication) AccessibilityVerticalUnits() NSAccessibilityUnits {
	
	rv := objc.Send[NSAccessibilityUnits](o.ID, objc.Sel("accessibilityVerticalUnits"))
	return rv
	}
// Sets the units that the layout area uses for vertical values.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityverticalunits(_:)
func (o NSApplication) SetAccessibilityVerticalUnits(accessibilityVerticalUnits NSAccessibilityUnits) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVerticalUnits:"), accessibilityVerticalUnits)
	}
// Returns the description of the layout area’s vertical units.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityverticalunitdescription()
func (o NSApplication) AccessibilityVerticalUnitDescription() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVerticalUnitDescription"))
	return foundation.NSStringFromID(rv).String()
	}
// Sets the description of the layout area’s vertical units.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityverticalunitdescription(_:)
func (o NSApplication) SetAccessibilityVerticalUnitDescription(accessibilityVerticalUnitDescription string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVerticalUnitDescription:"), objc.String(accessibilityVerticalUnitDescription))
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
func (o NSApplication) AccessibilityLayoutPointForScreenPoint(point corefoundation.CGPoint) corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("accessibilityLayoutPointForScreenPoint:"), point)
	return rv
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
func (o NSApplication) AccessibilityLayoutSizeForScreenSize(size corefoundation.CGSize) corefoundation.CGSize {
	
	rv := objc.Send[corefoundation.CGSize](o.ID, objc.Sel("accessibilityLayoutSizeForScreenSize:"), size)
	return rv
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
func (o NSApplication) AccessibilityScreenPointForLayoutPoint(point corefoundation.CGPoint) corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("accessibilityScreenPointForLayoutPoint:"), point)
	return rv
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
func (o NSApplication) AccessibilityScreenSizeForLayoutSize(size corefoundation.CGSize) corefoundation.CGSize {
	
	rv := objc.Send[corefoundation.CGSize](o.ID, objc.Sel("accessibilityScreenSizeForLayoutSize:"), size)
	return rv
	}
// Returns the allowed values for the slider accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityallowedvalues()
func (o NSApplication) AccessibilityAllowedValues() foundation.NSNumber {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityAllowedValues"))
	return foundation.NSNumberFromID(rv)
	}
// Sets the allowed values for the slider accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityallowedvalues(_:)
func (o NSApplication) SetAccessibilityAllowedValues(accessibilityAllowedValues foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityAllowedValues:"), accessibilityAllowedValues)
	}
// Returns the child label elements for the slider accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitylabeluielements()
func (o NSApplication) AccessibilityLabelUIElements() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityLabelUIElements"))
	return objectivec.Object{ID: rv}
	}
// Sets the child label elements for the slider accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitylabeluielements(_:)
func (o NSApplication) SetAccessibilityLabelUIElements(accessibilityLabelUIElements foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityLabelUIElements:"), accessibilityLabelUIElements)
	}
// Returns the value of the label accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitylabelvalue()
func (o NSApplication) AccessibilityLabelValue() float64 {
	
	rv := objc.Send[float64](o.ID, objc.Sel("accessibilityLabelValue"))
	return rv
	}
// Sets the value of the label accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitylabelvalue(_:)
func (o NSApplication) SetAccessibilityLabelValue(accessibilityLabelValue float64) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityLabelValue:"), accessibilityLabelValue)
	}
// Returns the contents that follow the divider accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitynextcontents()
func (o NSApplication) AccessibilityNextContents() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityNextContents"))
	return objectivec.Object{ID: rv}
	}
// Sets the contents that follow the divider accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitynextcontents(_:)
func (o NSApplication) SetAccessibilityNextContents(accessibilityNextContents foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityNextContents:"), accessibilityNextContents)
	}
// Returns the contents that precede the divider accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitypreviouscontents()
func (o NSApplication) AccessibilityPreviousContents() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityPreviousContents"))
	return objectivec.Object{ID: rv}
	}
// Sets the contents that precede the divider accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitypreviouscontents(_:)
func (o NSApplication) SetAccessibilityPreviousContents(accessibilityPreviousContents foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityPreviousContents:"), accessibilityPreviousContents)
	}
// Returns an array that contains the views and splitter bar from the split
// view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysplitters()
func (o NSApplication) AccessibilitySplitters() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySplitters"))
	return objectivec.Object{ID: rv}
	}
// Sets the array that contains the views and splitter bar from the split
// view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysplitters(_:)
func (o NSApplication) SetAccessibilitySplitters(accessibilitySplitters foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySplitters:"), accessibilitySplitters)
	}
// Returns the overflow button for the toolbar.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityoverflowbutton()
func (o NSApplication) AccessibilityOverflowButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityOverflowButton"))
	return objectivec.Object{ID: rv}
	}
// Sets the overflow button for the toolbar.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityoverflowbutton(_:)
func (o NSApplication) SetAccessibilityOverflowButton(accessibilityOverflowButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityOverflowButton:"), accessibilityOverflowButton)
	}
// Returns the tab accessibility elements for the tab view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitytabs()
func (o NSApplication) AccessibilityTabs() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTabs"))
	return objectivec.Object{ID: rv}
	}
// Sets the tab accessibility elements for the tab view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitytabs(_:)
func (o NSApplication) SetAccessibilityTabs(accessibilityTabs foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityTabs:"), accessibilityTabs)
	}
// Returns the user interface element that functions as a marker group for the
// ruler accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitymarkergroupuielement()
func (o NSApplication) AccessibilityMarkerGroupUIElement() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerGroupUIElement"))
	return objectivec.Object{ID: rv}
	}
// Sets the user interface element that functions as a marker group for the
// ruler accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymarkergroupuielement(_:)
func (o NSApplication) SetAccessibilityMarkerGroupUIElement(accessibilityMarkerGroupUIElement objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMarkerGroupUIElement:"), accessibilityMarkerGroupUIElement)
	}
// Returns the human-readable description of the marker type.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitymarkertypedescription()
func (o NSApplication) AccessibilityMarkerTypeDescription() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerTypeDescription"))
	return foundation.NSStringFromID(rv).String()
	}
// Sets the human-readable description of the marker type.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymarkertypedescription(_:)
func (o NSApplication) SetAccessibilityMarkerTypeDescription(accessibilityMarkerTypeDescription string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMarkerTypeDescription:"), objc.String(accessibilityMarkerTypeDescription))
	}
// Returns the array of marker accessibility elements for the ruler.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitymarkeruielements()
func (o NSApplication) AccessibilityMarkerUIElements() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerUIElements"))
	return objectivec.Object{ID: rv}
	}
// Sets the array of marker accessibility elements for the ruler.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymarkeruielements(_:)
func (o NSApplication) SetAccessibilityMarkerUIElements(accessibilityMarkerUIElements foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMarkerUIElements:"), accessibilityMarkerUIElements)
	}
// Returns the marker values for the ruler.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitymarkervalues()
func (o NSApplication) AccessibilityMarkerValues() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerValues"))
	return objectivec.Object{ID: rv}
	}
// Sets the marker values for the ruler.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymarkervalues(_:)
func (o NSApplication) SetAccessibilityMarkerValues(accessibilityMarkerValues objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMarkerValues:"), accessibilityMarkerValues)
	}
// Returns the type of markers for the ruler.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityrulermarkertype()
func (o NSApplication) AccessibilityRulerMarkerType() NSAccessibilityRulerMarkerType {
	
	rv := objc.Send[NSAccessibilityRulerMarkerType](o.ID, objc.Sel("accessibilityRulerMarkerType"))
	return rv
	}
// Sets the type of markers for the ruler.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityrulermarkertype(_:)
func (o NSApplication) SetAccessibilityRulerMarkerType(accessibilityRulerMarkerType NSAccessibilityRulerMarkerType) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRulerMarkerType:"), accessibilityRulerMarkerType)
	}
// Returns the units for the ruler.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityunits()
func (o NSApplication) AccessibilityUnits() NSAccessibilityUnits {
	
	rv := objc.Send[NSAccessibilityUnits](o.ID, objc.Sel("accessibilityUnits"))
	return rv
	}
// Sets the units used for the ruler.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityunits(_:)
func (o NSApplication) SetAccessibilityUnits(accessibilityUnits NSAccessibilityUnits) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityUnits:"), accessibilityUnits)
	}
// Returns the human-readable description of the ruler’s units.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityunitdescription()
func (o NSApplication) AccessibilityUnitDescription() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityUnitDescription"))
	return foundation.NSStringFromID(rv).String()
	}
// Sets the human-readable description of the ruler’s units.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityunitdescription(_:)
func (o NSApplication) SetAccessibilityUnitDescription(accessibilityUnitDescription string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityUnitDescription:"), objc.String(accessibilityUnitDescription))
	}
// Returns the URL for the file that the accessibility element represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitydocument()
func (o NSApplication) AccessibilityDocument() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDocument"))
	return foundation.NSStringFromID(rv).String()
	}
// Sets the URL for the file that the accessibility element represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitydocument(_:)
func (o NSApplication) SetAccessibilityDocument(accessibilityDocument string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDocument:"), objc.String(accessibilityDocument))
	}
// Returns a Boolean value that indicates whether the accessibility element is
// in an edited state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityedited()
func (o NSApplication) IsAccessibilityEdited() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityEdited"))
	return rv
	}
// Sets a Boolean value that indicates whether the accessibility element is in
// an edited state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityedited(_:)
func (o NSApplication) SetAccessibilityEdited(accessibilityEdited bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityEdited:"), accessibilityEdited)
	}
// Returns the filename for the file that the accessibility element
// represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityfilename()
func (o NSApplication) AccessibilityFilename() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityFilename"))
	return foundation.NSStringFromID(rv).String()
	}
// Sets the filename for the file that the accessibility element represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityfilename(_:)
func (o NSApplication) SetAccessibilityFilename(accessibilityFilename string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFilename:"), objc.String(accessibilityFilename))
	}
// Returns the elements that have links with the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitylinkeduielements()
func (o NSApplication) AccessibilityLinkedUIElements() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityLinkedUIElements"))
	return objectivec.Object{ID: rv}
	}
// Sets the elements that have links with the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitylinkeduielements(_:)
func (o NSApplication) SetAccessibilityLinkedUIElements(accessibilityLinkedUIElements foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityLinkedUIElements:"), accessibilityLinkedUIElements)
	}
// Returns the list of elements that the accessibility element is a title for.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityservesastitleforuielements()
func (o NSApplication) AccessibilityServesAsTitleForUIElements() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityServesAsTitleForUIElements"))
	return objectivec.Object{ID: rv}
	}
// Sets the list of elements that the accessibility element is a title for.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityservesastitleforuielements(_:)
func (o NSApplication) SetAccessibilityServesAsTitleForUIElements(accessibilityServesAsTitleForUIElements foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityServesAsTitleForUIElements:"), accessibilityServesAsTitleForUIElements)
	}
// Returns the static text element that represents the accessibility
// element’s title.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitytitleuielement()
func (o NSApplication) AccessibilityTitleUIElement() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTitleUIElement"))
	return objectivec.Object{ID: rv}
	}
// Sets the static text element that represents the accessibility element’s
// title.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitytitleuielement(_:)
func (o NSApplication) SetAccessibilityTitleUIElement(accessibilityTitleUIElement objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityTitleUIElement:"), accessibilityTitleUIElement)
	}
// Returns the clear button for the search field.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityclearbutton()
func (o NSApplication) AccessibilityClearButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityClearButton"))
	return objectivec.Object{ID: rv}
	}
// Sets the clear button for the search field.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityclearbutton(_:)
func (o NSApplication) SetAccessibilityClearButton(accessibilityClearButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityClearButton:"), accessibilityClearButton)
	}
// Returns the search button for the search field.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysearchbutton()
func (o NSApplication) AccessibilitySearchButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySearchButton"))
	return objectivec.Object{ID: rv}
	}
// Sets the search button for the search field.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysearchbutton(_:)
func (o NSApplication) SetAccessibilitySearchButton(accessibilitySearchButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySearchButton:"), accessibilitySearchButton)
	}
// Returns the search menu for the search field.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysearchmenu()
func (o NSApplication) AccessibilitySearchMenu() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySearchMenu"))
	return objectivec.Object{ID: rv}
	}
// Sets the search menu for the search field.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysearchmenu(_:)
func (o NSApplication) SetAccessibilitySearchMenu(accessibilitySearchMenu objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySearchMenu:"), accessibilitySearchMenu)
	}
// Cancels the current operation.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformCancel()
func (o NSApplication) AccessibilityPerformCancel() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformCancel"))
	return rv
	}
// Simulates pressing Return in the accessibility element.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use this method on elements that take keyboard input, such as a text field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformConfirm()
func (o NSApplication) AccessibilityPerformConfirm() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformConfirm"))
	return rv
	}
// Selects the accessibility element.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use this method on selectable elements, such as a menu item.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformPick()
func (o NSApplication) AccessibilityPerformPick() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformPick"))
	return rv
	}
// Simulates clicking the accessibility element.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use this method on elements that behave like buttons.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformPress()
func (o NSApplication) AccessibilityPerformPress() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformPress"))
	return rv
	}
// Displays the accessibility element’s alternative UI.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use this method to trigger changes to the UI due to a mouse-hover or
// similar event.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformShowAlternateUI()
func (o NSApplication) AccessibilityPerformShowAlternateUI() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformShowAlternateUI"))
	return rv
	}
// Returns to the accessibility element’s original UI.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Call this method after successfully calling
// [AccessibilityPerformShowAlternateUI] to return to the original UI.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformShowDefaultUI()
func (o NSApplication) AccessibilityPerformShowDefaultUI() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformShowDefaultUI"))
	return rv
	}
// Displays the menu accessibility element.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use this method to display the contextual menu for the element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformShowMenu()
func (o NSApplication) AccessibilityPerformShowMenu() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformShowMenu"))
	return rv
	}
// Brings the window to the front.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The window behaves as if you had clicked on the window’s title bar.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformRaise()
func (o NSApplication) AccessibilityPerformRaise() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformRaise"))
	return rv
	}
// Returns the increment button for the stepper accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityincrementbutton()
func (o NSApplication) AccessibilityIncrementButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityIncrementButton"))
	return objectivec.Object{ID: rv}
	}
// Sets the increment button for the stepper accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityincrementbutton(_:)
func (o NSApplication) SetAccessibilityIncrementButton(accessibilityIncrementButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityIncrementButton:"), accessibilityIncrementButton)
	}
// Returns the decrement button for the stepper accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitydecrementbutton()
func (o NSApplication) AccessibilityDecrementButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDecrementButton"))
	return objectivec.Object{ID: rv}
	}
// Sets the decrement button for the stepper accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitydecrementbutton(_:)
func (o NSApplication) SetAccessibilityDecrementButton(accessibilityDecrementButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDecrementButton:"), accessibilityDecrementButton)
	}
// Increments the accessibility element’s value.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use this method on elements that have an adjustable [accessibilityValue]
// property.
//
// [accessibilityValue]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityValue
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformIncrement()
func (o NSApplication) AccessibilityPerformIncrement() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformIncrement"))
	return rv
	}
// Decrements the accessibility element’s value.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use this method on elements that have an adjustable [accessibilityValue]
// property.
//
// [accessibilityValue]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityValue
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformDecrement()
func (o NSApplication) AccessibilityPerformDecrement() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformDecrement"))
	return rv
	}
// Deletes the accessibility element’s value.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use this method on elements with values.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformDelete()
func (o NSApplication) AccessibilityPerformDelete() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformDelete"))
	return rv
	}
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityattributeduserinputlabels()
func (o NSApplication) AccessibilityAttributedUserInputLabels() foundation.NSAttributedString {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityAttributedUserInputLabels"))
	return foundation.NSAttributedStringFromID(rv)
	}
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityuserinputlabels()
func (o NSApplication) AccessibilityUserInputLabels() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityUserInputLabels"))
	return foundation.NSStringFromID(rv).String()
	}
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityattributeduserinputlabels(_:)
func (o NSApplication) SetAccessibilityAttributedUserInputLabels(accessibilityAttributedUserInputLabels foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityAttributedUserInputLabels:"), accessibilityAttributedUserInputLabels)
	}
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityuserinputlabels(_:)
func (o NSApplication) SetAccessibilityUserInputLabels(accessibilityUserInputLabels foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityUserInputLabels:"), accessibilityUserInputLabels)
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

