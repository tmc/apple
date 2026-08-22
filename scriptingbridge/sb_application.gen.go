// Code generated from Apple documentation for ScriptingBridge. DO NOT EDIT.

package scriptingbridge

import (
	"sync"

	"github.com/tmc/apple/coreservices"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SBApplication] class.
var (
	_SBApplicationClass     SBApplicationClass
	_SBApplicationClassOnce sync.Once
)

func getSBApplicationClass() SBApplicationClass {
	_SBApplicationClassOnce.Do(func() {
		_SBApplicationClass = SBApplicationClass{class: objc.GetClass("SBApplication")}
	})
	return _SBApplicationClass
}

// GetSBApplicationClass returns the class object for SBApplication.
func GetSBApplicationClass() SBApplicationClass {
	return getSBApplicationClass()
}

type SBApplicationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SBApplicationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SBApplicationClass) Alloc() SBApplication {
	rv := objc.Send[SBApplication](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// The [SBApplication] class provides a mechanism enabling an Objective-C
// program to send Apple events to a scriptable application and receive Apple
// events in response. It thereby makes it possible for that program to
// control the application and exchange data with it. Scripting Bridge works
// by bridging data types between Apple event descriptors and Cocoa objects.
//
// # Overview
//
// Although [SBApplication] includes methods that manually send and process
// Apple events, you should never have to call these methods directly.
// Instead, subclasses of [SBApplication] implement application-specific
// methods that handle the sending of Apple events automatically.
//
// For example, if you wanted to get the current iTunes track, you can simply
// use the `currentTrack` method of the dynamically defined subclass for the
// iTunes application—which handles the details of sending the Apple event
// for you—rather than figuring out the more complicated, low-level
// alternative:
//
// If you do need to send Apple events manually, consider using the
// [NSAppleEventDescriptor] class.
//
// # Subclassing Notes
//
// You rarely instantiate [SBApplication] objects directly. Instead, you get
// the shared instance of a application-specific subclass typically by calling
// one of the `applicationWith...` class methods, using a bundle identifier,
// process identifier, or URL to identify the application.
//
// # Initializing a Scriptable Application Object
//
//   - [SBApplication.InitWithBundleIdentifier]: Returns an instance of an [SBApplication] subclass that represents the target application identified by the given bundle identifier.
//   - [SBApplication.InitWithProcessIdentifier]: Returns an instance of an [SBApplication] subclass that represents the target application identified by the given process identifier.
//   - [SBApplication.InitWithURL]: Returns an instance of an [SBApplication] subclass that represents the target application identified by the given URL.
//
// # Creating a Scripting Class
//
//   - [SBApplication.ClassForScriptingClass]: Returns a class object that represents a particular class in the target application.
//
// # Controlling the Application
//
//   - [SBApplication.Activate]: Moves the target application to the foreground immediately.
//   - [SBApplication.IsRunning]: A Boolean that indicates whether the target application represented by the receiver is running.
//   - [SBApplication.LaunchFlags]: The launch flags for the application represented by the receiver.
//   - [SBApplication.SetLaunchFlags]
//   - [SBApplication.SendMode]: The mode for sending Apple events to the target application.
//   - [SBApplication.SetSendMode]
//   - [SBApplication.Timeout]: The period the application will wait to receive reply Apple events.
//   - [SBApplication.SetTimeout]
//
// # Managing the Delegate
//
//   - [SBApplication.Delegate]: The error-handling delegate of the receiver.
//   - [SBApplication.SetDelegate]
//
// See: https://developer.apple.com/documentation/ScriptingBridge/SBApplication
type SBApplication struct {
	SBObject
}

// SBApplicationFromID constructs a [SBApplication] from an objc.ID.
//
// The [SBApplication] class provides a mechanism enabling an Objective-C
// program to send Apple events to a scriptable application and receive Apple
// events in response. It thereby makes it possible for that program to
// control the application and exchange data with it. Scripting Bridge works
// by bridging data types between Apple event descriptors and Cocoa objects.
func SBApplicationFromID(id objc.ID) SBApplication {
	return SBApplication{SBObject: SBObjectFromID(id)}
}

// NOTE: SBApplication adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SBApplication] class.
//
// # Initializing a Scriptable Application Object
//
//   - [ISBApplication.InitWithBundleIdentifier]: Returns an instance of an [SBApplication] subclass that represents the target application identified by the given bundle identifier.
//   - [ISBApplication.InitWithProcessIdentifier]: Returns an instance of an [SBApplication] subclass that represents the target application identified by the given process identifier.
//   - [ISBApplication.InitWithURL]: Returns an instance of an [SBApplication] subclass that represents the target application identified by the given URL.
//
// # Creating a Scripting Class
//
//   - [ISBApplication.ClassForScriptingClass]: Returns a class object that represents a particular class in the target application.
//
// # Controlling the Application
//
//   - [ISBApplication.Activate]: Moves the target application to the foreground immediately.
//   - [ISBApplication.IsRunning]: A Boolean that indicates whether the target application represented by the receiver is running.
//   - [ISBApplication.LaunchFlags]: The launch flags for the application represented by the receiver.
//   - [ISBApplication.SetLaunchFlags]
//   - [ISBApplication.SendMode]: The mode for sending Apple events to the target application.
//   - [ISBApplication.SetSendMode]
//   - [ISBApplication.Timeout]: The period the application will wait to receive reply Apple events.
//   - [ISBApplication.SetTimeout]
//
// # Managing the Delegate
//
//   - [ISBApplication.Delegate]: The error-handling delegate of the receiver.
//   - [ISBApplication.SetDelegate]
//
// See: https://developer.apple.com/documentation/ScriptingBridge/SBApplication
type ISBApplication interface {
	ISBObject

	// Topic: Initializing a Scriptable Application Object

	// Returns an instance of an [SBApplication] subclass that represents the target application identified by the given bundle identifier.
	InitWithBundleIdentifier(ident string) SBApplication
	// Returns an instance of an [SBApplication] subclass that represents the target application identified by the given process identifier.
	InitWithProcessIdentifier(pid int32) SBApplication
	// Returns an instance of an [SBApplication] subclass that represents the target application identified by the given URL.
	InitWithURL(url foundation.NSURL) SBApplication

	// Topic: Creating a Scripting Class

	// Returns a class object that represents a particular class in the target application.
	ClassForScriptingClass(className string) objectivec.Class

	// Topic: Controlling the Application

	// Moves the target application to the foreground immediately.
	Activate()
	// A Boolean that indicates whether the target application represented by the receiver is running.
	IsRunning() bool
	// The launch flags for the application represented by the receiver.
	LaunchFlags() uint32
	SetLaunchFlags(value uint32)
	// The mode for sending Apple events to the target application.
	SendMode() coreservices.AESendMode
	SetSendMode(value coreservices.AESendMode)
	// The period the application will wait to receive reply Apple events.
	Timeout() int
	SetTimeout(value int)

	// Topic: Managing the Delegate

	// The error-handling delegate of the receiver.
	Delegate() SBApplicationDelegate
	SetDelegate(value SBApplicationDelegate)
}

// Init initializes the instance.
func (s SBApplication) Init() SBApplication {
	rv := objc.Send[SBApplication](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SBApplication) Autorelease() SBApplication {
	rv := objc.Send[SBApplication](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSBApplication creates a new SBApplication instance.
func NewSBApplication() SBApplication {
	class := getSBApplicationClass()
	rv := objc.Send[SBApplication](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns an instance of an [SBApplication] subclass that represents the
// target application identified by the given bundle identifier.
//
// ident: A bundle identifier specifying an application that is OSA-compliant.
//
// # Return Value
//
// An initialized shared instance of an [SBApplication] subclass that
// represents a target application with the bundle identifier of `ident`.
// Returns `nil` if no such application can be found or if the application
// does not have a scripting interface.
//
// # Discussion
//
// If you must initialize an [SBApplication] object explictly, you should use
// this initializer if possible; unlike
// [SBApplication.InitWithProcessIdentifier] and [SBApplication.InitWithURL],
// this method is not dependent on changeable factors such as the target
// application’s path or process ID. Even so, you should rarely have to
// initialize an [SBApplication] object yourself; instead, you should
// initialize an application-specific subclass such as `iTunesApplication`.
//
// Note that this method does not check whether an application with the given
// bundle identifier actually exists.
//
// See: https://developer.apple.com/documentation/ScriptingBridge/SBApplication/init(bundleIdentifier:)
func NewSBApplicationWithBundleIdentifier(ident string) SBApplication {
	instance := getSBApplicationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithBundleIdentifier:"), objc.String(ident))
	return SBApplicationFromID(rv)
}

// Returns an instance of an [SBObject] subclass initialized with the given
// data.
//
// data: An object containing data for the new [SBObject] object. The data varies
// according to the type of scripting object to be created.
//
// # Return Value
//
// An [SBObject] object or `nil` if the object could not be initialized.
//
// # Discussion
//
// Scripting Bridge does not actually create an object in the target
// application until you add the object returned from this method to an
// element array ([SBElementArray]).
//
// See: https://developer.apple.com/documentation/ScriptingBridge/SBObject/init(data:)
func NewSBApplicationWithData(data objectivec.IObject) SBApplication {
	instance := getSBApplicationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:"), data)
	return SBApplicationFromID(rv)
}

// Returns an instance of an [SBObject] subclass initialized with the
// specified properties and data and added to the designated element array.
//
// code: A four-character code used to identify an element in the target
// application’s scripting interface. See [Apple Event Manager] for details.
//
// properties: A dictionary with [NSNumber] keys specifying the four-character codes of
// properties (that is, attributes or to-one relationships) and the values for
// those properties. Pass `nil` if you are initializing the object by `data`
// only.
//
// data: An object containing data for the new [SBObject] object. The data varies
// according to the type of scripting object to be created. Pass `nil` if you
// initializing the object by `properties` only.
//
// # Return Value
//
// An [SBObject] object or `nil` if the object could not be initialized.
//
// # Discussion
//
// Unlike the other initializers of this class, this method not only
// initializes the [SBObject] object but adds it to a specified element array.
// This method is the designated initializer.
//
// See: https://developer.apple.com/documentation/ScriptingBridge/SBObject/init(elementCode:properties:data:)
//
// [Apple Event Manager]: https://developer.apple.com/documentation/applicationservices/apple_event_manager
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func NewSBApplicationWithElementCodePropertiesData(code coreservices.DescType, properties foundation.INSDictionary, data objectivec.IObject) SBApplication {
	instance := getSBApplicationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithElementCode:properties:data:"), code, properties, data)
	return SBApplicationFromID(rv)
}

// Returns an instance of an [SBApplication] subclass that represents the
// target application identified by the given process identifier.
//
// pid: A BSD process ID specifying an application that is OSA-compliant. Often you
// can get the process ID of a process using the [processIdentifier] method of
// [NSTask].
//
// # Return Value
//
// An initialized [SBApplication] that you can use to communicate with the
// target application specified by the process ID. Returns `nil` if no such
// application can be found or if the application does not have a scripting
// interface.
//
// # Discussion
//
// You should avoid using this method unless you know nothing about an
// external application but its PID. In most cases, it is better to use
// [SBApplication.InitWithBundleIdentifier], which will dynamically locate the
// external application’s path at runtime, or [SBApplication.InitWithURL],
// which is not dependent on the external application being open at the time
// the method is called.
//
// See: https://developer.apple.com/documentation/ScriptingBridge/SBApplication/init(processIdentifier:)
//
// [processIdentifier]: https://developer.apple.com/documentation/Foundation/Process/processIdentifier
func NewSBApplicationWithProcessIdentifier(pid int32) SBApplication {
	instance := getSBApplicationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithProcessIdentifier:"), pid)
	return SBApplicationFromID(rv)
}

// Returns an instance of an [SBObject] subclass initialized with the
// specified properties.
//
// properties: A dictionary with keys specifying the names of properties (that is,
// attributes or to-one relationships) and the values for those properties.
//
// # Return Value
//
// An [SBObject] object or `nil` if the object could not be initialized.
//
// # Discussion
//
// Scripting Bridge does not actually create an object in the target
// application until you add the object returned from this method to an
// element array ([SBElementArray]).
//
// See: https://developer.apple.com/documentation/ScriptingBridge/SBObject/init(properties:)
func NewSBApplicationWithProperties(properties foundation.INSDictionary) SBApplication {
	instance := getSBApplicationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithProperties:"), properties)
	return SBApplicationFromID(rv)
}

// Returns an instance of an [SBApplication] subclass that represents the
// target application identified by the given URL.
//
// url: A Universal Resource Locator (URL) specifying an application that is
// OSA-compliant.
//
// # Return Value
//
// An initialized [SBApplication] that you can use to communicate with the
// target application specified by the process ID. Returns `nil` if an
// application could not be found or if the application does not have a
// scripting interface.
//
// # Discussion
//
// This approach to initializing [SBApplication] objects should be used only
// if you know for certain the URL of the target application. In most cases,
// it is better to use [applicationWithBundleIdentifier:] which dynamically
// locates the target application at runtime. Even so, you should rarely have
// to initialize an [SBApplication] yourself.
//
// This method currently supports file URLs (“) and remote application URLs
// (“). It checks whether a file exists at the specified path, but it does
// not check whether an application identified via “ exists.
//
// See: https://developer.apple.com/documentation/ScriptingBridge/SBApplication/init(url:)
//
// [applicationWithBundleIdentifier:]: https://developer.apple.com/documentation/ScriptingBridge/SBApplication/applicationWithBundleIdentifier:
func NewSBApplicationWithURL(url foundation.NSURL) SBApplication {
	instance := getSBApplicationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:"), url)
	return SBApplicationFromID(rv)
}

// Returns an instance of an [SBApplication] subclass that represents the
// target application identified by the given bundle identifier.
//
// ident: A bundle identifier specifying an application that is OSA-compliant.
//
// # Return Value
//
// An initialized shared instance of an [SBApplication] subclass that
// represents a target application with the bundle identifier of `ident`.
// Returns `nil` if no such application can be found or if the application
// does not have a scripting interface.
//
// # Discussion
//
// If you must initialize an [SBApplication] object explictly, you should use
// this initializer if possible; unlike
// [SBApplication.InitWithProcessIdentifier] and [SBApplication.InitWithURL],
// this method is not dependent on changeable factors such as the target
// application’s path or process ID. Even so, you should rarely have to
// initialize an [SBApplication] object yourself; instead, you should
// initialize an application-specific subclass such as `iTunesApplication`.
//
// Note that this method does not check whether an application with the given
// bundle identifier actually exists.
//
// See: https://developer.apple.com/documentation/ScriptingBridge/SBApplication/init(bundleIdentifier:)
func (s SBApplication) InitWithBundleIdentifier(ident string) SBApplication {
	rv := objc.Send[SBApplication](s.ID, objc.Sel("initWithBundleIdentifier:"), objc.String(ident))
	return rv
}

// Returns an instance of an [SBApplication] subclass that represents the
// target application identified by the given process identifier.
//
// pid: A BSD process ID specifying an application that is OSA-compliant. Often you
// can get the process ID of a process using the [processIdentifier] method of
// [NSTask].
//
// # Return Value
//
// An initialized [SBApplication] that you can use to communicate with the
// target application specified by the process ID. Returns `nil` if no such
// application can be found or if the application does not have a scripting
// interface.
//
// # Discussion
//
// You should avoid using this method unless you know nothing about an
// external application but its PID. In most cases, it is better to use
// [SBApplication.InitWithBundleIdentifier], which will dynamically locate the
// external application’s path at runtime, or [SBApplication.InitWithURL],
// which is not dependent on the external application being open at the time
// the method is called.
//
// See: https://developer.apple.com/documentation/ScriptingBridge/SBApplication/init(processIdentifier:)
//
// [processIdentifier]: https://developer.apple.com/documentation/Foundation/Process/processIdentifier
func (s SBApplication) InitWithProcessIdentifier(pid int32) SBApplication {
	rv := objc.Send[SBApplication](s.ID, objc.Sel("initWithProcessIdentifier:"), pid)
	return rv
}

// Returns an instance of an [SBApplication] subclass that represents the
// target application identified by the given URL.
//
// url: A Universal Resource Locator (URL) specifying an application that is
// OSA-compliant.
//
// # Return Value
//
// An initialized [SBApplication] that you can use to communicate with the
// target application specified by the process ID. Returns `nil` if an
// application could not be found or if the application does not have a
// scripting interface.
//
// # Discussion
//
// This approach to initializing [SBApplication] objects should be used only
// if you know for certain the URL of the target application. In most cases,
// it is better to use [applicationWithBundleIdentifier:] which dynamically
// locates the target application at runtime. Even so, you should rarely have
// to initialize an [SBApplication] yourself.
//
// This method currently supports file URLs (“) and remote application URLs
// (“). It checks whether a file exists at the specified path, but it does
// not check whether an application identified via “ exists.
//
// See: https://developer.apple.com/documentation/ScriptingBridge/SBApplication/init(url:)
//
// [applicationWithBundleIdentifier:]: https://developer.apple.com/documentation/ScriptingBridge/SBApplication/applicationWithBundleIdentifier:
func (s SBApplication) InitWithURL(url foundation.NSURL) SBApplication {
	rv := objc.Send[SBApplication](s.ID, objc.Sel("initWithURL:"), url)
	return rv
}

// Returns a class object that represents a particular class in the target
// application.
//
// className: The name of the scripting class, as it appears in the scripting interface.
// For example, “document”.
//
// # Return Value
//
// A [Class] object representing the scripting class.
//
// # Discussion
//
// You invoke this method on an instance of a scriptable application. Once you
// have the class object, you may allocate an instance of the class and
// appropriately the raw instance. Or you may use it in a call to
// [isKind(of:)] to determine the class type of an object.
//
// See: https://developer.apple.com/documentation/ScriptingBridge/SBApplication/class(forScriptingClass:)
//
// [isKind(of:)]: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/isKind(of:)
func (s SBApplication) ClassForScriptingClass(className string) objectivec.Class {
	rv := objc.Send[objectivec.Class](s.ID, objc.Sel("classForScriptingClass:"), objc.String(className))
	return objectivec.Class(rv)
}

// Moves the target application to the foreground immediately.
//
// # Discussion
//
// If the target application is not already running, this method launches it.
//
// See: https://developer.apple.com/documentation/ScriptingBridge/SBApplication/activate()
func (s SBApplication) Activate() {
	objc.Send[objc.ID](s.ID, objc.Sel("activate"))
}

// A Boolean that indicates whether the target application represented by the
// receiver is running.
//
// # Discussion
//
// true if the application is running, false otherwise.
//
// This may be true for instances initialized with a bundle identifier or URL
// because [SBApplication] launches the application only when it’s necessary
// to send it an event.
//
// See: https://developer.apple.com/documentation/ScriptingBridge/SBApplication/isRunning
func (s SBApplication) IsRunning() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isRunning"))
	return rv
}

// The launch flags for the application represented by the receiver.
//
// # Discussion
//
// For more information, see [Launch Services].
//
// See: https://developer.apple.com/documentation/ScriptingBridge/SBApplication/launchFlags
//
// [Launch Services]: https://developer.apple.com/documentation/coreservices/launch_services
func (s SBApplication) LaunchFlags() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("launchFlags"))
	return rv
}
func (s SBApplication) SetLaunchFlags(value uint32) {
	objc.Send[struct{}](s.ID, objc.Sel("setLaunchFlags:"), value)
}

// The mode for sending Apple events to the target application.
//
// # Discussion
//
// For more information, see [Apple Event Manager].
//
// The default send mode is [kAEWaitReply]. If the send mode is something
// other than `kAEWaitReply`, the receiver might not correctly handle reply
// events from the target application.
//
// See: https://developer.apple.com/documentation/ScriptingBridge/SBApplication/sendMode
//
// [Apple Event Manager]: https://developer.apple.com/documentation/applicationservices/apple_event_manager
// [kAEWaitReply]: https://developer.apple.com/documentation/coreservices/1542914-anonymous/kaewaitreply
func (s SBApplication) SendMode() coreservices.AESendMode {
	rv := objc.Send[coreservices.AESendMode](s.ID, objc.Sel("sendMode"))
	return coreservices.AESendMode(rv)
}
func (s SBApplication) SetSendMode(value coreservices.AESendMode) {
	objc.Send[struct{}](s.ID, objc.Sel("setSendMode:"), value)
}

// The period the application will wait to receive reply Apple events.
//
// # Discussion
//
// For more information, see [Apple Event Manager].
//
// The default timeout value is [kAEDefaultTimeout], which is about a minute.
// If you want the receiver to wait indefinitely for reply Apple events, use
// [kNoTimeOut]. For more information, see [Apple Event Manager].
//
// See: https://developer.apple.com/documentation/ScriptingBridge/SBApplication/timeout
//
// [Apple Event Manager]: https://developer.apple.com/documentation/applicationservices/apple_event_manager
// [kAEDefaultTimeout]: https://developer.apple.com/documentation/coreservices/1542814-timeout_constants/kaedefaulttimeout
// [kNoTimeOut]: https://developer.apple.com/documentation/coreservices/1542814-timeout_constants/knotimeout
func (s SBApplication) Timeout() int {
	rv := objc.Send[int](s.ID, objc.Sel("timeout"))
	return rv
}
func (s SBApplication) SetTimeout(value int) {
	objc.Send[struct{}](s.ID, objc.Sel("setTimeout:"), value)
}

// The error-handling delegate of the receiver.
//
// # Discussion
//
// The delegate should implement the [EventDidFailWithError] method of the
// [SBApplicationDelegate] informal protocol.
//
// See: https://developer.apple.com/documentation/ScriptingBridge/SBApplication/delegate
func (s SBApplication) Delegate() SBApplicationDelegate {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("delegate"))
	return SBApplicationDelegateObjectFromID(rv)
}
func (s SBApplication) SetDelegate(value SBApplicationDelegate) {
	objc.Send[struct{}](s.ID, objc.Sel("setDelegate:"), value)
}
