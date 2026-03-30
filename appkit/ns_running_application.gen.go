// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSRunningApplication] class.
var (
	_NSRunningApplicationClass     NSRunningApplicationClass
	_NSRunningApplicationClassOnce sync.Once
)

func getNSRunningApplicationClass() NSRunningApplicationClass {
	_NSRunningApplicationClassOnce.Do(func() {
		_NSRunningApplicationClass = NSRunningApplicationClass{class: objc.GetClass("NSRunningApplication")}
	})
	return _NSRunningApplicationClass
}

// GetNSRunningApplicationClass returns the class object for NSRunningApplication.
func GetNSRunningApplicationClass() NSRunningApplicationClass {
	return getNSRunningApplicationClass()
}

type NSRunningApplicationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSRunningApplicationClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSRunningApplicationClass) Alloc() NSRunningApplication {
	rv := objc.Send[NSRunningApplication](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that can manipulate and provide information for a single instance
// of an app.
//
// # Overview
//
// Some properties of an app are fixed, such as the bundle identifier. Other
// properties may vary over time, such as whether the app is hidden.
// Properties that vary can be observed with key-value observing, in which
// case the description comment for the method notes this capability.
//
// Properties that vary over time are inherently race-prone. For example, a
// hidden app may unhide itself at any time. To ameliorate this, properties
// persist until the next turn of the main run loop in a common mode. For
// example, if you repeatedly poll an unhidden app for its hidden property
// without allowing the run loop to run, it will continue to return false,
// even if the app hides, until the next turn of the run loop.
//
// [NSRunningApplication] is thread safe, in that its properties are returned
// atomically. However, it is still subject to the main run loop policy
// described above. If you access an instance of [NSRunningApplication] from a
// background thread, be aware that its time-varying properties may change
// from under you as the main run loop runs (or not).
//
// An [NSRunningApplication] instance remains valid after the app exits.
// However, most properties lose their significance, and some properties may
// not be available on a terminated application.
//
// To access the list of all running apps, use the [NSRunningApplication.RunningApplications]
// method in [NSWorkspace].
//
// # Activating applications
//
//   - [NSRunningApplication.ActivateWithOptions]: Attempts to activate the application using the specified options.
//   - [NSRunningApplication.ActivateFromApplicationOptions]: Attempts to activate the application using the specified options.
//   - [NSRunningApplication.Active]: Indicates whether the application is currently frontmost.
//   - [NSRunningApplication.ActivationPolicy]: Indicates the activation policy of the application.
//
// # Hiding and unhiding applications
//
//   - [NSRunningApplication.Hide]: Attempts to hide or the application.
//   - [NSRunningApplication.Unhide]: Attempts to unhide or the application.
//   - [NSRunningApplication.Hidden]: Indicates whether the application is currently hidden.
//
// # Application information
//
//   - [NSRunningApplication.LocalizedName]: Indicates the localized name of the application.
//   - [NSRunningApplication.Icon]: Returns the icon for the receiver’s application.
//   - [NSRunningApplication.BundleIdentifier]: Indicates the [CFBundleIdentifier] of the application.
//   - [NSRunningApplication.BundleURL]: Indicates the URL to the application’s bundle.
//   - [NSRunningApplication.ExecutableArchitecture]: Indicates the executing processor architecture for the application.
//   - [NSRunningApplication.ExecutableURL]: Indicates the URL to the application’s executable.
//   - [NSRunningApplication.LaunchDate]: Indicates the date when the application was launched.
//   - [NSRunningApplication.FinishedLaunching]: A Boolean value that determines whether the receiver’s process has finished launching.
//   - [NSRunningApplication.ProcessIdentifier]: Indicates the process identifier (pid) of the application.
//   - [NSRunningApplication.OwnsMenuBar]: Returns whether the application owns the current menu bar.
//
// # Terminating applications
//
//   - [NSRunningApplication.ForceTerminate]: Attempts to force the receiver to quit.
//   - [NSRunningApplication.Terminate]: Attempts to quit the receiver normally.
//   - [NSRunningApplication.Terminated]: Indicates that the receiver’s application has terminated.
//
// See: https://developer.apple.com/documentation/AppKit/NSRunningApplication
type NSRunningApplication struct {
	objectivec.Object
}

// NSRunningApplicationFromID constructs a [NSRunningApplication] from an objc.ID.
//
// An object that can manipulate and provide information for a single instance
// of an app.
func NSRunningApplicationFromID(id objc.ID) NSRunningApplication {
	return NSRunningApplication{objectivec.Object{ID: id}}
}

// NOTE: NSRunningApplication adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSRunningApplication] class.
//
// # Activating applications
//
//   - [INSRunningApplication.ActivateWithOptions]: Attempts to activate the application using the specified options.
//   - [INSRunningApplication.ActivateFromApplicationOptions]: Attempts to activate the application using the specified options.
//   - [INSRunningApplication.Active]: Indicates whether the application is currently frontmost.
//   - [INSRunningApplication.ActivationPolicy]: Indicates the activation policy of the application.
//
// # Hiding and unhiding applications
//
//   - [INSRunningApplication.Hide]: Attempts to hide or the application.
//   - [INSRunningApplication.Unhide]: Attempts to unhide or the application.
//   - [INSRunningApplication.Hidden]: Indicates whether the application is currently hidden.
//
// # Application information
//
//   - [INSRunningApplication.LocalizedName]: Indicates the localized name of the application.
//   - [INSRunningApplication.Icon]: Returns the icon for the receiver’s application.
//   - [INSRunningApplication.BundleIdentifier]: Indicates the [CFBundleIdentifier] of the application.
//   - [INSRunningApplication.BundleURL]: Indicates the URL to the application’s bundle.
//   - [INSRunningApplication.ExecutableArchitecture]: Indicates the executing processor architecture for the application.
//   - [INSRunningApplication.ExecutableURL]: Indicates the URL to the application’s executable.
//   - [INSRunningApplication.LaunchDate]: Indicates the date when the application was launched.
//   - [INSRunningApplication.FinishedLaunching]: A Boolean value that determines whether the receiver’s process has finished launching.
//   - [INSRunningApplication.ProcessIdentifier]: Indicates the process identifier (pid) of the application.
//   - [INSRunningApplication.OwnsMenuBar]: Returns whether the application owns the current menu bar.
//
// # Terminating applications
//
//   - [INSRunningApplication.ForceTerminate]: Attempts to force the receiver to quit.
//   - [INSRunningApplication.Terminate]: Attempts to quit the receiver normally.
//   - [INSRunningApplication.Terminated]: Indicates that the receiver’s application has terminated.
//
// See: https://developer.apple.com/documentation/AppKit/NSRunningApplication
type INSRunningApplication interface {
	objectivec.IObject

	// Topic: Activating applications

	// Attempts to activate the application using the specified options.
	ActivateWithOptions(options NSApplicationActivationOptions) bool
	// Attempts to activate the application using the specified options.
	ActivateFromApplicationOptions(application INSRunningApplication, options NSApplicationActivationOptions) bool
	// Indicates whether the application is currently frontmost.
	Active() bool
	// Indicates the activation policy of the application.
	ActivationPolicy() NSApplicationActivationPolicy

	// Topic: Hiding and unhiding applications

	// Attempts to hide or the application.
	Hide() bool
	// Attempts to unhide or the application.
	Unhide() bool
	// Indicates whether the application is currently hidden.
	Hidden() bool

	// Topic: Application information

	// Indicates the localized name of the application.
	LocalizedName() string
	// Returns the icon for the receiver’s application.
	Icon() INSImage
	// Indicates the [CFBundleIdentifier] of the application.
	BundleIdentifier() string
	// Indicates the URL to the application’s bundle.
	BundleURL() foundation.INSURL
	// Indicates the executing processor architecture for the application.
	ExecutableArchitecture() int
	// Indicates the URL to the application’s executable.
	ExecutableURL() foundation.INSURL
	// Indicates the date when the application was launched.
	LaunchDate() foundation.INSDate
	// A Boolean value that determines whether the receiver’s process has finished launching.
	FinishedLaunching() bool
	// Indicates the process identifier (pid) of the application.
	ProcessIdentifier() int32
	// Returns whether the application owns the current menu bar.
	OwnsMenuBar() bool

	// Topic: Terminating applications

	// Attempts to force the receiver to quit.
	ForceTerminate() bool
	// Attempts to quit the receiver normally.
	Terminate() bool
	// Indicates that the receiver’s application has terminated.
	Terminated() bool

	// Returns an array of running apps.
	RunningApplications() INSRunningApplication
	SetRunningApplications(value INSRunningApplication)
}

// Init initializes the instance.
func (r NSRunningApplication) Init() NSRunningApplication {
	rv := objc.Send[NSRunningApplication](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r NSRunningApplication) Autorelease() NSRunningApplication {
	rv := objc.Send[NSRunningApplication](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSRunningApplication creates a new NSRunningApplication instance.
func NewNSRunningApplication() NSRunningApplication {
	class := getNSRunningApplicationClass()
	rv := objc.Send[NSRunningApplication](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the running application with the given process identifier, or nil
// if no application has that pid.
//
// pid: The process identifier.
//
// # Return Value
//
// An instance of [NSRunningApplication] for the specified `pid`, or nil if
// the application has no process identifier.
//
// # Discussion
//
// Applications that do not have [PIDs] cannot be returned from this method.
//
// See: https://developer.apple.com/documentation/AppKit/NSRunningApplication/init(processIdentifier:)
func NewRunningApplicationWithProcessIdentifier(pid int32) NSRunningApplication {
	rv := objc.Send[objc.ID](objc.ID(getNSRunningApplicationClass().class), objc.Sel("runningApplicationWithProcessIdentifier:"), pid)
	return NSRunningApplicationFromID(rv)
}

// Attempts to activate the application using the specified options.
//
// options: The options to use when activating the application. See
// [NSApplication.ActivationOptions] for the possible values.
//
// # Return Value
//
// true if the application was activated successfully, otherwise false.
//
// # Discussion
//
// This method will return false if the application has quit, or is not a type
// of application than can be activated.
//
// See: https://developer.apple.com/documentation/AppKit/NSRunningApplication/activate(options:)
//
// [NSApplication.ActivationOptions]: https://developer.apple.com/documentation/AppKit/NSApplication/ActivationOptions
func (r NSRunningApplication) ActivateWithOptions(options NSApplicationActivationOptions) bool {
	rv := objc.Send[bool](r.ID, objc.Sel("activateWithOptions:"), options)
	return rv
}

// Attempts to activate the application using the specified options.
//
// application: The application to activate.
//
// options: The options to use during activation.
//
// # Return Value
//
// Returns true if the request is allowed by the system, otherwise false.
//
// # Discussion
//
// Use this method to request app activation. Calling this method doesn’t
// guarantee app activation. For cooperative activation, the other application
// should call [YieldActivationToApplication] or equivalent prior to the
// target application invoking [Activate].
//
// See: https://developer.apple.com/documentation/AppKit/NSRunningApplication/activate(from:options:)
func (r NSRunningApplication) ActivateFromApplicationOptions(application INSRunningApplication, options NSApplicationActivationOptions) bool {
	rv := objc.Send[bool](r.ID, objc.Sel("activateFromApplication:options:"), application, options)
	return rv
}

// Attempts to hide or the application.
//
// # Return Value
//
// true if the application was successfully hidden, otherwise false.
//
// # Discussion
//
// The property of this value will be false if the application has already
// quit, or if of a type that is unable to be hidden.
//
// See: https://developer.apple.com/documentation/AppKit/NSRunningApplication/hide()
func (r NSRunningApplication) Hide() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("hide"))
	return rv
}

// Attempts to unhide or the application.
//
// # Return Value
//
// true if the application was successfully shown, otherwise false.
//
// # Discussion
//
// The property of this value will be false if the application has already
// quit, or if of a type that is unable to be hidden.
//
// See: https://developer.apple.com/documentation/AppKit/NSRunningApplication/unhide()
func (r NSRunningApplication) Unhide() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("unhide"))
	return rv
}

// Attempts to force the receiver to quit.
//
// # Return Value
//
// Returns true if the application successfully terminated, otherwise false.
//
// # Discussion
//
// This method will return false if the application is no longer running when
// the `forceTerminate` message is sent to the receiver.
//
// This method may return before the receiver exits; you should observe the
// terminated property to determine when the application terminates.
//
// Sandboxed applications can’t use this method to terminate other
// applciations. This method returns false when called from a sandboxed
// application.
//
// See: https://developer.apple.com/documentation/AppKit/NSRunningApplication/forceTerminate()
func (r NSRunningApplication) ForceTerminate() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("forceTerminate"))
	return rv
}

// Attempts to quit the receiver normally.
//
// # Return Value
//
// Returns true if the request was successfully sent, otherwise false.
//
// # Discussion
//
// This method will return false if the application is no longer running when
// the terminate message is sent to the receiver.
//
// This method may return before the receiver exits; you should observe the
// terminated property to determine when the application terminates.
//
// Sandboxed applications can’t use this method to terminate other
// applications. This method returns false when called from a sandboxed
// application.
//
// See: https://developer.apple.com/documentation/AppKit/NSRunningApplication/terminate()
func (r NSRunningApplication) Terminate() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("terminate"))
	return rv
}

// Returns an array of currently running applications with the specified
// bundle identifier.
//
// bundleIdentifier: The bundle identifier.
//
// # Return Value
//
// An array of [NSRunningApplications], or an empty array if no applications
// match the bundle identifier.
//
// See: https://developer.apple.com/documentation/AppKit/NSRunningApplication/runningApplications(withBundleIdentifier:)
func (_NSRunningApplicationClass NSRunningApplicationClass) RunningApplicationsWithBundleIdentifier(bundleIdentifier string) []NSRunningApplication {
	rv := objc.Send[[]objc.ID](objc.ID(_NSRunningApplicationClass.class), objc.Sel("runningApplicationsWithBundleIdentifier:"), objc.String(bundleIdentifier))
	return objc.ConvertSlice(rv, func(id objc.ID) NSRunningApplication {
		return NSRunningApplicationFromID(id)
	})
}

// Terminates invisibly running applications as if triggered by system memory
// pressure.
//
// # Discussion
//
// This method is intended for installer applications and similar applications
// to ensure that nothing is unexpectedly relying on the files they’re
// replacing.
//
// See: https://developer.apple.com/documentation/AppKit/NSRunningApplication/terminateAutomaticallyTerminableApplications()
func (_NSRunningApplicationClass NSRunningApplicationClass) TerminateAutomaticallyTerminableApplications() {
	objc.Send[objc.ID](objc.ID(_NSRunningApplicationClass.class), objc.Sel("terminateAutomaticallyTerminableApplications"))
}

// Indicates whether the application is currently frontmost.
//
// # Discussion
//
// This property is observable using key-value observing.
//
// See: https://developer.apple.com/documentation/AppKit/NSRunningApplication/isActive
func (r NSRunningApplication) Active() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("isActive"))
	return rv
}

// Indicates the activation policy of the application.
//
// # Discussion
//
// The value returned by this property is usually fixed, but it may change
// through a call to [ActivateWithOptions].
//
// This property is observable using key-value observing.
//
// See: https://developer.apple.com/documentation/AppKit/NSRunningApplication/activationPolicy
func (r NSRunningApplication) ActivationPolicy() NSApplicationActivationPolicy {
	rv := objc.Send[NSApplicationActivationPolicy](r.ID, objc.Sel("activationPolicy"))
	return NSApplicationActivationPolicy(rv)
}

// Indicates whether the application is currently hidden.
//
// # Discussion
//
// This property is observable using key-value observing.
//
// See: https://developer.apple.com/documentation/AppKit/NSRunningApplication/isHidden
func (r NSRunningApplication) Hidden() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("isHidden"))
	return rv
}

// Indicates the localized name of the application.
//
// # Discussion
//
// The value of this property is dependent on the current localization of the
// application and is suitable for presentation to the user.
//
// See: https://developer.apple.com/documentation/AppKit/NSRunningApplication/localizedName
func (r NSRunningApplication) LocalizedName() string {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("localizedName"))
	return foundation.NSStringFromID(rv).String()
}

// Returns the icon for the receiver’s application.
//
// See: https://developer.apple.com/documentation/AppKit/NSRunningApplication/icon
func (r NSRunningApplication) Icon() INSImage {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("icon"))
	return NSImageFromID(objc.ID(rv))
}

// Indicates the [CFBundleIdentifier] of the application.
//
// # Discussion
//
// The value of this property will be `nil` if the application does not have
// an Info.plist.
//
// See: https://developer.apple.com/documentation/AppKit/NSRunningApplication/bundleIdentifier
func (r NSRunningApplication) BundleIdentifier() string {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("bundleIdentifier"))
	return foundation.NSStringFromID(rv).String()
}

// Indicates the URL to the application’s bundle.
//
// # Discussion
//
// The value of this property is `nil` of the application does not have a
// bundle structure.
//
// See: https://developer.apple.com/documentation/AppKit/NSRunningApplication/bundleURL
func (r NSRunningApplication) BundleURL() foundation.INSURL {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("bundleURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}

// Indicates the executing processor architecture for the application.
//
// # Discussion
//
// The returned value will be one of the constants in Mach-O Architecture in
// [Bundle].
//
// See: https://developer.apple.com/documentation/AppKit/NSRunningApplication/executableArchitecture
//
// [Bundle]: https://developer.apple.com/documentation/Foundation/Bundle
func (r NSRunningApplication) ExecutableArchitecture() int {
	rv := objc.Send[int](r.ID, objc.Sel("executableArchitecture"))
	return rv
}

// Indicates the URL to the application’s executable.
//
// See: https://developer.apple.com/documentation/AppKit/NSRunningApplication/executableURL
func (r NSRunningApplication) ExecutableURL() foundation.INSURL {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("executableURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}

// Indicates the date when the application was launched.
//
// # Discussion
//
// This property is only available for applications that were launched by
// LaunchServices.
//
// See: https://developer.apple.com/documentation/AppKit/NSRunningApplication/launchDate
func (r NSRunningApplication) LaunchDate() foundation.INSDate {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("launchDate"))
	return foundation.NSDateFromID(objc.ID(rv))
}

// A Boolean value that determines whether the receiver’s process has
// finished launching.
//
// # Discussion
//
// The value of this property corresponds to the running application having
// received an [didFinishLaunchingNotification] notification internally. Some
// applications do not post this notification (applications that do not rely
// on [NSApplication]) and so are never reported as finished launching.
//
// This property is observable using key-value observing.
//
// See: https://developer.apple.com/documentation/AppKit/NSRunningApplication/isFinishedLaunching
//
// [didFinishLaunchingNotification]: https://developer.apple.com/documentation/AppKit/NSApplication/didFinishLaunchingNotification
func (r NSRunningApplication) FinishedLaunching() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("isFinishedLaunching"))
	return rv
}

// Indicates the process identifier (pid) of the application.
//
// # Discussion
//
// Not all applications have a pid. Applications without a pid return a value
// of -1.
//
// Do not rely on this for comparing processes, instead compare
// NSRunningApplication instances using [isEqual(_:)].
//
// See: https://developer.apple.com/documentation/AppKit/NSRunningApplication/processIdentifier
//
// [isEqual(_:)]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/isEqual(_:)
func (r NSRunningApplication) ProcessIdentifier() int32 {
	rv := objc.Send[int32](r.ID, objc.Sel("processIdentifier"))
	return rv
}

// Returns whether the application owns the current menu bar.
//
// # Discussion
//
// This property is observable using key-value observing.
//
// See: https://developer.apple.com/documentation/AppKit/NSRunningApplication/ownsMenuBar
func (r NSRunningApplication) OwnsMenuBar() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("ownsMenuBar"))
	return rv
}

// Indicates that the receiver’s application has terminated.
//
// # Discussion
//
// The value of terminated is true if the receiver’s application has
// terminated, otherwise false.
//
// This property is observable using key-value observing.
//
// See: https://developer.apple.com/documentation/AppKit/NSRunningApplication/isTerminated
func (r NSRunningApplication) Terminated() bool {
	rv := objc.Send[bool](r.ID, objc.Sel("isTerminated"))
	return rv
}

// Returns an array of running apps.
//
// See: https://developer.apple.com/documentation/appkit/nsworkspace/runningapplications
func (r NSRunningApplication) RunningApplications() INSRunningApplication {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("runningApplications"))
	return NSRunningApplicationFromID(objc.ID(rv))
}
func (r NSRunningApplication) SetRunningApplications(value INSRunningApplication) {
	objc.Send[struct{}](r.ID, objc.Sel("setRunningApplications:"), value)
}

// Returns an [NSRunningApplication] representing this application.
//
// # Return Value
//
// An [NSRunningApplication] instance for the current application.
//
// See: https://developer.apple.com/documentation/AppKit/NSRunningApplication/current
func (_NSRunningApplicationClass NSRunningApplicationClass) CurrentApplication() NSRunningApplication {
	rv := objc.Send[objc.ID](objc.ID(_NSRunningApplicationClass.class), objc.Sel("currentApplication"))
	return NSRunningApplicationFromID(objc.ID(rv))
}
