// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSWorkspaceOpenConfiguration] class.
var (
	_NSWorkspaceOpenConfigurationClass     NSWorkspaceOpenConfigurationClass
	_NSWorkspaceOpenConfigurationClassOnce sync.Once
)

func getNSWorkspaceOpenConfigurationClass() NSWorkspaceOpenConfigurationClass {
	_NSWorkspaceOpenConfigurationClassOnce.Do(func() {
		_NSWorkspaceOpenConfigurationClass = NSWorkspaceOpenConfigurationClass{class: objc.GetClass("NSWorkspaceOpenConfiguration")}
	})
	return _NSWorkspaceOpenConfigurationClass
}

// GetNSWorkspaceOpenConfigurationClass returns the class object for NSWorkspaceOpenConfiguration.
func GetNSWorkspaceOpenConfigurationClass() NSWorkspaceOpenConfigurationClass {
	return getNSWorkspaceOpenConfigurationClass()
}

type NSWorkspaceOpenConfigurationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSWorkspaceOpenConfigurationClass) Alloc() NSWorkspaceOpenConfiguration {
	rv := objc.Send[NSWorkspaceOpenConfiguration](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The configuration options for opening URLs or launching apps.
//
// # Overview
// 
// Create an [NSWorkspaceOpenConfiguration] object before launching an app or
// opening a URL using the shared [NSWorkspace] object. Use the properties of
// this object to customize the behavior of the launched app or the handling
// of the URLs. For example, you might tell the app to hide itself immediately
// after launch.
//
// # Handling URLs
//
//   - [NSWorkspaceOpenConfiguration.RequiresUniversalLinks]: A Boolean value indicating whether you require the URL to have an associated universal link.
//   - [NSWorkspaceOpenConfiguration.SetRequiresUniversalLinks]
//   - [NSWorkspaceOpenConfiguration.ForPrinting]: A Boolean value indicating whether you want to print the contents of documents and URLs instead of opening them.
//   - [NSWorkspaceOpenConfiguration.SetForPrinting]
//
// # Specifying app-related behaviors
//
//   - [NSWorkspaceOpenConfiguration.Activates]: A Boolean value indicating whether the system activates the app and brings it to the foreground.
//   - [NSWorkspaceOpenConfiguration.SetActivates]
//   - [NSWorkspaceOpenConfiguration.AddsToRecentItems]: A Boolean value indicating whether to add the app or documents to the Recent Items menu.
//   - [NSWorkspaceOpenConfiguration.SetAddsToRecentItems]
//   - [NSWorkspaceOpenConfiguration.AllowsRunningApplicationSubstitution]: A Boolean value that indicates whether to use a running instance of an application even if it’s at a different URL.
//   - [NSWorkspaceOpenConfiguration.SetAllowsRunningApplicationSubstitution]
//   - [NSWorkspaceOpenConfiguration.CreatesNewApplicationInstance]: A Boolean value indicating whether you want the system to launch a new instance of the app.
//   - [NSWorkspaceOpenConfiguration.SetCreatesNewApplicationInstance]
//   - [NSWorkspaceOpenConfiguration.Hides]: A Boolean value indicating whether you want the app to hide itself after it launches.
//   - [NSWorkspaceOpenConfiguration.SetHides]
//   - [NSWorkspaceOpenConfiguration.HidesOthers]: A Boolean value indicating whether you want to hide all apps except the one that launched.
//   - [NSWorkspaceOpenConfiguration.SetHidesOthers]
//
// # Prompting the user
//
//   - [NSWorkspaceOpenConfiguration.PromptsUserIfNeeded]: A Boolean value indicating whether to display errors, authentication requests, or other UI elements to the user.
//   - [NSWorkspaceOpenConfiguration.SetPromptsUserIfNeeded]
//
// # Specifying launch attributes
//
//   - [NSWorkspaceOpenConfiguration.AppleEvent]: The first Apple event to send to the new app.
//   - [NSWorkspaceOpenConfiguration.SetAppleEvent]
//   - [NSWorkspaceOpenConfiguration.Arguments]: The set of command-line arguments to pass to a new app instance at launch time.
//   - [NSWorkspaceOpenConfiguration.SetArguments]
//   - [NSWorkspaceOpenConfiguration.Environment]: The set of environment variables to set in a new app instance.
//   - [NSWorkspaceOpenConfiguration.SetEnvironment]
//   - [NSWorkspaceOpenConfiguration.Architecture]: The architecture version of the app to launch.
//   - [NSWorkspaceOpenConfiguration.SetArchitecture]
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration
type NSWorkspaceOpenConfiguration struct {
	objectivec.Object
}

// NSWorkspaceOpenConfigurationFromID constructs a [NSWorkspaceOpenConfiguration] from an objc.ID.
//
// The configuration options for opening URLs or launching apps.
func NSWorkspaceOpenConfigurationFromID(id objc.ID) NSWorkspaceOpenConfiguration {
	return NSWorkspaceOpenConfiguration{objectivec.Object{ID: id}}
}
// NOTE: NSWorkspaceOpenConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSWorkspaceOpenConfiguration] class.
//
// # Handling URLs
//
//   - [INSWorkspaceOpenConfiguration.RequiresUniversalLinks]: A Boolean value indicating whether you require the URL to have an associated universal link.
//   - [INSWorkspaceOpenConfiguration.SetRequiresUniversalLinks]
//   - [INSWorkspaceOpenConfiguration.ForPrinting]: A Boolean value indicating whether you want to print the contents of documents and URLs instead of opening them.
//   - [INSWorkspaceOpenConfiguration.SetForPrinting]
//
// # Specifying app-related behaviors
//
//   - [INSWorkspaceOpenConfiguration.Activates]: A Boolean value indicating whether the system activates the app and brings it to the foreground.
//   - [INSWorkspaceOpenConfiguration.SetActivates]
//   - [INSWorkspaceOpenConfiguration.AddsToRecentItems]: A Boolean value indicating whether to add the app or documents to the Recent Items menu.
//   - [INSWorkspaceOpenConfiguration.SetAddsToRecentItems]
//   - [INSWorkspaceOpenConfiguration.AllowsRunningApplicationSubstitution]: A Boolean value that indicates whether to use a running instance of an application even if it’s at a different URL.
//   - [INSWorkspaceOpenConfiguration.SetAllowsRunningApplicationSubstitution]
//   - [INSWorkspaceOpenConfiguration.CreatesNewApplicationInstance]: A Boolean value indicating whether you want the system to launch a new instance of the app.
//   - [INSWorkspaceOpenConfiguration.SetCreatesNewApplicationInstance]
//   - [INSWorkspaceOpenConfiguration.Hides]: A Boolean value indicating whether you want the app to hide itself after it launches.
//   - [INSWorkspaceOpenConfiguration.SetHides]
//   - [INSWorkspaceOpenConfiguration.HidesOthers]: A Boolean value indicating whether you want to hide all apps except the one that launched.
//   - [INSWorkspaceOpenConfiguration.SetHidesOthers]
//
// # Prompting the user
//
//   - [INSWorkspaceOpenConfiguration.PromptsUserIfNeeded]: A Boolean value indicating whether to display errors, authentication requests, or other UI elements to the user.
//   - [INSWorkspaceOpenConfiguration.SetPromptsUserIfNeeded]
//
// # Specifying launch attributes
//
//   - [INSWorkspaceOpenConfiguration.AppleEvent]: The first Apple event to send to the new app.
//   - [INSWorkspaceOpenConfiguration.SetAppleEvent]
//   - [INSWorkspaceOpenConfiguration.Arguments]: The set of command-line arguments to pass to a new app instance at launch time.
//   - [INSWorkspaceOpenConfiguration.SetArguments]
//   - [INSWorkspaceOpenConfiguration.Environment]: The set of environment variables to set in a new app instance.
//   - [INSWorkspaceOpenConfiguration.SetEnvironment]
//   - [INSWorkspaceOpenConfiguration.Architecture]: The architecture version of the app to launch.
//   - [INSWorkspaceOpenConfiguration.SetArchitecture]
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration
type INSWorkspaceOpenConfiguration interface {
	objectivec.IObject

	// Topic: Handling URLs

	// A Boolean value indicating whether you require the URL to have an associated universal link.
	RequiresUniversalLinks() bool
	SetRequiresUniversalLinks(value bool)
	// A Boolean value indicating whether you want to print the contents of documents and URLs instead of opening them.
	ForPrinting() bool
	SetForPrinting(value bool)

	// Topic: Specifying app-related behaviors

	// A Boolean value indicating whether the system activates the app and brings it to the foreground.
	Activates() bool
	SetActivates(value bool)
	// A Boolean value indicating whether to add the app or documents to the Recent Items menu.
	AddsToRecentItems() bool
	SetAddsToRecentItems(value bool)
	// A Boolean value that indicates whether to use a running instance of an application even if it’s at a different URL.
	AllowsRunningApplicationSubstitution() bool
	SetAllowsRunningApplicationSubstitution(value bool)
	// A Boolean value indicating whether you want the system to launch a new instance of the app.
	CreatesNewApplicationInstance() bool
	SetCreatesNewApplicationInstance(value bool)
	// A Boolean value indicating whether you want the app to hide itself after it launches.
	Hides() bool
	SetHides(value bool)
	// A Boolean value indicating whether you want to hide all apps except the one that launched.
	HidesOthers() bool
	SetHidesOthers(value bool)

	// Topic: Prompting the user

	// A Boolean value indicating whether to display errors, authentication requests, or other UI elements to the user.
	PromptsUserIfNeeded() bool
	SetPromptsUserIfNeeded(value bool)

	// Topic: Specifying launch attributes

	// The first Apple event to send to the new app.
	AppleEvent() foundation.NSAppleEventDescriptor
	SetAppleEvent(value foundation.NSAppleEventDescriptor)
	// The set of command-line arguments to pass to a new app instance at launch time.
	Arguments() []string
	SetArguments(value []string)
	// The set of environment variables to set in a new app instance.
	Environment() foundation.INSDictionary
	SetEnvironment(value foundation.INSDictionary)
	// The architecture version of the app to launch.
	Architecture() int32
	SetArchitecture(value int32)
}

// Init initializes the instance.
func (w NSWorkspaceOpenConfiguration) Init() NSWorkspaceOpenConfiguration {
	rv := objc.Send[NSWorkspaceOpenConfiguration](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w NSWorkspaceOpenConfiguration) Autorelease() NSWorkspaceOpenConfiguration {
	rv := objc.Send[NSWorkspaceOpenConfiguration](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSWorkspaceOpenConfiguration creates a new NSWorkspaceOpenConfiguration instance.
func NewNSWorkspaceOpenConfiguration() NSWorkspaceOpenConfiguration {
	class := getNSWorkspaceOpenConfigurationClass()
	rv := objc.Send[NSWorkspaceOpenConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A Boolean value indicating whether you require the URL to have an
// associated universal link.
//
// # Discussion
// 
// The default value of this property is [false], which tells the app to open
// any URL you provide. Set the value to [true] when you want the app to open
// only valid universal links.
// 
// The app must be specifically configured to open universal links, and
// attempts to open such links fail with an appropriate error if the app
// isn’t properly configured. Attempts may also fail with an error if the
// user disabled support for opening links with the specified app.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/requiresUniversalLinks
func (w NSWorkspaceOpenConfiguration) RequiresUniversalLinks() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("requiresUniversalLinks"))
	return rv
}
func (w NSWorkspaceOpenConfiguration) SetRequiresUniversalLinks(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setRequiresUniversalLinks:"), value)
}
// A Boolean value indicating whether you want to print the contents of
// documents and URLs instead of opening them.
//
// # Discussion
// 
// The default value of this property is [false], which causes the app to open
// documents and URLs. Set the value to [true] if you want the app to print
// the items instead.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/isForPrinting
func (w NSWorkspaceOpenConfiguration) ForPrinting() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("isForPrinting"))
	return rv
}
func (w NSWorkspaceOpenConfiguration) SetForPrinting(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setForPrinting:"), value)
}
// A Boolean value indicating whether the system activates the app and brings
// it to the foreground.
//
// # Discussion
// 
// The default value of this property is [true], which causes the system to
// bring the app to the foreground.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/activates
func (w NSWorkspaceOpenConfiguration) Activates() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("activates"))
	return rv
}
func (w NSWorkspaceOpenConfiguration) SetActivates(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setActivates:"), value)
}
// A Boolean value indicating whether to add the app or documents to the
// Recent Items menu.
//
// # Discussion
// 
// The default value of this property is [true], which causes AppKit to add
// the items to the Recent Items menu.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/addsToRecentItems
func (w NSWorkspaceOpenConfiguration) AddsToRecentItems() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("addsToRecentItems"))
	return rv
}
func (w NSWorkspaceOpenConfiguration) SetAddsToRecentItems(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setAddsToRecentItems:"), value)
}
// A Boolean value that indicates whether to use a running instance of an
// application even if it’s at a different URL.
//
// # Discussion
// 
// If an instance of an application is already running and is capable of
// opening the provided URLs, but the running instance is at a different URL,
// use the running application.
// 
// This property defaults to [true]. Set this to [false] if you let the user
// select between specific versions of an application, or let them choose a
// particular installation.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/allowsRunningApplicationSubstitution
func (w NSWorkspaceOpenConfiguration) AllowsRunningApplicationSubstitution() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("allowsRunningApplicationSubstitution"))
	return rv
}
func (w NSWorkspaceOpenConfiguration) SetAllowsRunningApplicationSubstitution(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setAllowsRunningApplicationSubstitution:"), value)
}
// A Boolean value indicating whether you want the system to launch a new
// instance of the app.
//
// # Discussion
// 
// When the value of this property is [true], the system always launches a new
// version of the app, even if an existing copy is already running. The
// default value of this property is [false], which causes the system to open
// the already running app when present.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/createsNewApplicationInstance
func (w NSWorkspaceOpenConfiguration) CreatesNewApplicationInstance() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("createsNewApplicationInstance"))
	return rv
}
func (w NSWorkspaceOpenConfiguration) SetCreatesNewApplicationInstance(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setCreatesNewApplicationInstance:"), value)
}
// A Boolean value indicating whether you want the app to hide itself after it
// launches.
//
// # Discussion
// 
// The default value of this property is [false], which leaves the app in its
// default state after launch. Setting the property to [true] causes the app
// to hide itself.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/hides
func (w NSWorkspaceOpenConfiguration) Hides() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("hides"))
	return rv
}
func (w NSWorkspaceOpenConfiguration) SetHides(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setHides:"), value)
}
// A Boolean value indicating whether you want to hide all apps except the one
// that launched.
//
// # Discussion
// 
// The default value of this property is [false], which leaves the visibility
// of other apps unchanged. Setting the property to [true] hides all apps
// except the one that launched.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/hidesOthers
func (w NSWorkspaceOpenConfiguration) HidesOthers() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("hidesOthers"))
	return rv
}
func (w NSWorkspaceOpenConfiguration) SetHidesOthers(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setHidesOthers:"), value)
}
// A Boolean value indicating whether to display errors, authentication
// requests, or other UI elements to the user.
//
// # Discussion
// 
// When the value of this property is [true], the system presents a user
// interface to request or display relevant information. The system waits
// until the user dismisses the UI before calling any relevant completion
// handlers. The default value of this property is [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/promptsUserIfNeeded
func (w NSWorkspaceOpenConfiguration) PromptsUserIfNeeded() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("promptsUserIfNeeded"))
	return rv
}
func (w NSWorkspaceOpenConfiguration) SetPromptsUserIfNeeded(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setPromptsUserIfNeeded:"), value)
}
// The first Apple event to send to the new app.
//
// # Discussion
// 
// The default value of this property is `nil`, which causes the system to
// send a default Apple event, as needed. The system sends the event only if
// an instance of the app is already running.
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/appleEvent
func (w NSWorkspaceOpenConfiguration) AppleEvent() foundation.NSAppleEventDescriptor {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("appleEvent"))
	return foundation.NSAppleEventDescriptorFromID(objc.ID(rv))
}
func (w NSWorkspaceOpenConfiguration) SetAppleEvent(value foundation.NSAppleEventDescriptor) {
	objc.Send[struct{}](w.ID, objc.Sel("setAppleEvent:"), value)
}
// The set of command-line arguments to pass to a new app instance at launch
// time.
//
// # Discussion
// 
// The default value of this property is an empty array. When launching a new
// instance of an app, use this property to specify any additional launch
// arguments. The system inserts the app’s path as the first element in the
// array.
// 
// If the calling process is sandboxed, the system ignores the value of this
// property.
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/arguments
func (w NSWorkspaceOpenConfiguration) Arguments() []string {
	rv := objc.Send[[]objc.ID](w.ID, objc.Sel("arguments"))
	return objc.ConvertSliceToStrings(rv)
}
func (w NSWorkspaceOpenConfiguration) SetArguments(value []string) {
	objc.Send[struct{}](w.ID, objc.Sel("setArguments:"), objectivec.StringSliceToNSArray(value))
}
// The set of environment variables to set in a new app instance.
//
// # Discussion
// 
// The default value of this property is an empty dictionary. When launching a
// new instance of an app, use this property to specify the key/value pairs
// for any environment variables.
// 
// If the calling process is sandboxed, the system ignores the value of this
// property.
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/environment
func (w NSWorkspaceOpenConfiguration) Environment() foundation.INSDictionary {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("environment"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (w NSWorkspaceOpenConfiguration) SetEnvironment(value foundation.INSDictionary) {
	objc.Send[struct{}](w.ID, objc.Sel("setEnvironment:"), value)
}
// The architecture version of the app to launch.
//
// # Discussion
// 
// The default value of this property is `CPU_TYPE_ANY`, which causes the
// system to launch the app’s preferred architecture. You may specify a
// different value to force the system to launch that architecture. For a list
// of possible types, see the `` header file.
//
// See: https://developer.apple.com/documentation/AppKit/NSWorkspace/OpenConfiguration/architecture
func (w NSWorkspaceOpenConfiguration) Architecture() int32 {
	rv := objc.Send[int32](w.ID, objc.Sel("architecture"))
	return rv
}
func (w NSWorkspaceOpenConfiguration) SetArchitecture(value int32) {
	objc.Send[struct{}](w.ID, objc.Sel("setArchitecture:"), value)
}

